#!/usr/bin/env python3
"""Generate a daily Amazon and Microsoft LeetCode practice pair."""

from __future__ import annotations

import argparse
import csv
import html
import json
import os
import re
import shutil
import subprocess
import sys
import tempfile
import urllib.error
import urllib.request
from dataclasses import dataclass, field
from datetime import date, datetime, time, timedelta
from pathlib import Path
from typing import Any, Callable
from urllib.parse import urlparse
from zoneinfo import ZoneInfo


TIMEZONE = ZoneInfo("Africa/Cairo")
COMPANY_FILES = {
    "Amazon": "amazon_6months.csv",
    "Microsoft": "microsoft_6months.csv",
}
GRAPHQL_URL = "https://leetcode.com/graphql/"
GRAPHQL_QUERY = """
query questionData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionFrontendId
    title
    titleSlug
    content
    difficulty
    isPaidOnly
    codeSnippets { langSlug code }
  }
}
"""


class GeneratorError(RuntimeError):
    """A fatal generator error."""


class CandidateUnavailable(RuntimeError):
    """A candidate cannot produce the requested full local scaffold."""


@dataclass
class Problem:
    problem_id: str
    title: str
    acceptance: str
    difficulty: str
    link: str
    frequencies: dict[str, float] = field(default_factory=dict)

    @property
    def companies(self) -> list[str]:
        return [name for name in COMPANY_FILES if name in self.frequencies]

    @property
    def slug(self) -> str:
        parts = [part for part in urlparse(self.link.strip()).path.split("/") if part]
        if len(parts) >= 2 and parts[0] == "problems":
            return parts[1]
        raise GeneratorError(f"Invalid LeetCode link for problem {self.problem_id}: {self.link}")


@dataclass
class Details:
    title: str
    content: str
    difficulty: str
    go_code: str


@dataclass
class Selected:
    company: str
    problem: Problem
    details: Details
    solution: str


def load_catalog(repo_root: Path) -> dict[str, Problem]:
    catalog: dict[str, Problem] = {}
    for company, filename in COMPANY_FILES.items():
        csv_path = repo_root / filename
        if not csv_path.is_file():
            raise GeneratorError(f"Missing source file: {csv_path}")
        with csv_path.open(newline="", encoding="utf-8-sig") as handle:
            reader = csv.DictReader(handle)
            required = {
                "ID",
                "Title",
                "Acceptance",
                "Difficulty",
                "Frequency",
                "Leetcode Question Link",
            }
            if not reader.fieldnames or not required.issubset(reader.fieldnames):
                raise GeneratorError(f"Unexpected CSV columns in {csv_path}")
            for row_number, row in enumerate(reader, start=2):
                problem_id = row["ID"].strip()
                title = row["Title"].strip()
                link = row["Leetcode Question Link"].strip()
                try:
                    frequency = float(row["Frequency"].strip())
                except ValueError as exc:
                    raise GeneratorError(
                        f"Invalid frequency in {csv_path}:{row_number}"
                    ) from exc
                if not problem_id or not title or not link:
                    raise GeneratorError(f"Incomplete row in {csv_path}:{row_number}")
                current = catalog.get(problem_id)
                if current is None:
                    current = Problem(
                        problem_id=problem_id,
                        title=title,
                        acceptance=row["Acceptance"].strip(),
                        difficulty=row["Difficulty"].strip(),
                        link=link,
                    )
                    catalog[problem_id] = current
                current.frequencies[company] = frequency
    return catalog


def discover_existing_ids(repo_root: Path) -> set[str]:
    existing: set[str] = set()
    year_pattern = re.compile(r"^20\d{2}$")
    folder_pattern = re.compile(r"^(\d+)\.")
    for year_dir in repo_root.iterdir():
        if not year_dir.is_dir() or not year_pattern.match(year_dir.name):
            continue
        for root, directories, _ in os.walk(year_dir):
            directories[:] = [name for name in directories if not name.startswith(".")]
            for name in directories:
                match = folder_pattern.match(name)
                if match:
                    existing.add(match.group(1))
    return existing


def fetch_details(problem: Problem, timeout: int = 20) -> Details:
    payload = json.dumps(
        {"query": GRAPHQL_QUERY, "variables": {"titleSlug": problem.slug}}
    ).encode("utf-8")
    request = urllib.request.Request(
        GRAPHQL_URL,
        data=payload,
        headers={
            "Content-Type": "application/json",
            "User-Agent": "solvendo-daily-problem/1.0",
            "Referer": problem.link,
        },
    )
    try:
        with urllib.request.urlopen(request, timeout=timeout) as response:
            result = json.load(response)
    except (urllib.error.URLError, TimeoutError, json.JSONDecodeError) as exc:
        raise GeneratorError(f"LeetCode request failed for {problem.problem_id}: {exc}") from exc
    if result.get("errors"):
        raise GeneratorError(f"LeetCode returned GraphQL errors for {problem.problem_id}")
    question = result.get("data", {}).get("question")
    if not question:
        raise CandidateUnavailable("problem metadata is unavailable")
    if question.get("isPaidOnly"):
        raise CandidateUnavailable("problem is paid-only")
    content = question.get("content")
    snippets = question.get("codeSnippets") or []
    go_code = next(
        (snippet.get("code") for snippet in snippets if snippet.get("langSlug") == "golang"),
        None,
    )
    if not content or not go_code:
        raise CandidateUnavailable("full statement or Go starter is unavailable")
    return Details(
        title=question.get("title") or problem.title,
        content=content,
        difficulty=question.get("difficulty") or problem.difficulty,
        go_code=go_code,
    )


def _inject_todo_bodies(go_code: str) -> str:
    """Keep LeetCode signatures intact while making empty starters compile locally."""
    pattern = re.compile(r"(?ms)(^func[^\{]*\{)([ \t\r\n]*)(\})")
    return pattern.sub(r'\1\n\tpanic("TODO")\n\3', go_code)


def _local_support_types(go_code: str) -> str:
    without_comments = re.sub(r"/\*.*?\*/", "", go_code, flags=re.DOTALL)
    without_comments = re.sub(r"//.*", "", without_comments)
    declared = set(re.findall(r"\btype\s+([A-Z]\w*)", without_comments))
    used: set[str] = set()
    signature_pattern = re.compile(
        r"(?ms)^func(?:\s*\([^)]*\))?\s+\w+\s*\(([^)]*)\)\s*([^\{]*)\{"
    )
    for match in signature_pattern.finditer(without_comments):
        used.update(re.findall(r"\b[A-Z]\w*\b", " ".join(match.groups())))
    missing = used - declared
    definitions: list[str] = []
    if "ListNode" in missing:
        definitions.append("type ListNode struct {\n\tVal int\n\tNext *ListNode\n}")
        missing.remove("ListNode")
    if "TreeNode" in missing:
        definitions.append(
            "type TreeNode struct {\n\tVal int\n\tLeft *TreeNode\n\tRight *TreeNode\n}"
        )
        missing.remove("TreeNode")
    if "Node" in missing:
        definitions.append(
            "type Node struct {\n"
            "\tVal int\n\tNext, Random, Left, Right, Prev, Child *Node\n"
            "\tNeighbors, Children []*Node\n}"
        )
        missing.remove("Node")
    for type_name in sorted(missing):
        definitions.append(f"type {type_name} struct{{}}")
    if not definitions:
        return ""
    return (
        "// LOCAL SUPPORT TYPES BEGIN - do not copy these to LeetCode.\n"
        + "\n\n".join(definitions)
        + "\n// LOCAL SUPPORT TYPES END\n\n"
    )


def _find_tool(name: str) -> str | None:
    found = shutil.which(name)
    if found:
        return found
    for prefix in (Path("/usr/local/go/bin"), Path("/opt/homebrew/bin")):
        candidate = prefix / name
        if candidate.is_file():
            return str(candidate)
    return None


def build_solution(go_code: str, validate: bool = True) -> str:
    scaffold = _inject_todo_bodies(go_code.strip())
    source = (
        "package main\n\n"
        + _local_support_types(scaffold)
        + "// LEETCODE SOLUTION BEGIN - copy this section to LeetCode.\n"
        + scaffold
        + "\n// LEETCODE SOLUTION END\n\n"
        + "func main() {\n\t// TODO: add local test cases.\n}\n"
    )
    gofmt = _find_tool("gofmt")
    if not gofmt:
        raise GeneratorError("gofmt was not found")
    formatted = subprocess.run(
        [gofmt], input=source, text=True, capture_output=True, check=False
    )
    if formatted.returncode:
        raise CandidateUnavailable(f"Go starter cannot be formatted: {formatted.stderr.strip()}")
    if validate:
        go = _find_tool("go")
        if not go:
            raise GeneratorError("go was not found")
        with tempfile.TemporaryDirectory(prefix="solvendo-go-check-") as temp_dir:
            solution_path = Path(temp_dir) / "solution.go"
            solution_path.write_text(formatted.stdout, encoding="utf-8")
            checked = subprocess.run(
                [go, "test", "solution.go"],
                cwd=temp_dir,
                text=True,
                capture_output=True,
                check=False,
                env={**os.environ, "GOCACHE": "/tmp/solvendo-go-cache"},
            )
            if checked.returncode:
                reason = checked.stderr.strip() or checked.stdout.strip()
                raise CandidateUnavailable(f"Go starter does not compile locally: {reason}")
    return formatted.stdout


def select_problem(
    catalog: dict[str, Problem],
    company: str,
    excluded: set[str],
    fetcher: Callable[[Problem], Details] = fetch_details,
    validate_go: bool = True,
) -> Selected:
    candidates = sorted(
        (
            problem
            for problem in catalog.values()
            if company in problem.frequencies and problem.problem_id not in excluded
        ),
        key=lambda problem: (-problem.frequencies[company], int(problem.problem_id)),
    )
    rejected: list[str] = []
    for problem in candidates:
        try:
            details = fetcher(problem)
            solution = build_solution(details.go_code, validate=validate_go)
        except CandidateUnavailable as exc:
            rejected.append(f"{problem.problem_id} ({exc})")
            continue
        return Selected(company=company, problem=problem, details=details, solution=solution)
    suffix = f" Rejected: {', '.join(rejected[:5])}" if rejected else ""
    raise GeneratorError(f"No eligible unsolved free {company} problem remains.{suffix}")


def safe_folder_title(title: str) -> str:
    cleaned = re.sub(r"[/:\x00-\x1f]", " - ", title)
    cleaned = re.sub(r"\s+", " ", cleaned).strip(" .")
    return cleaned or "Untitled Problem"


def folder_for(repo_root: Path, run_date: date, selected: Selected) -> Path:
    title = safe_folder_title(selected.details.title)
    return repo_root / str(run_date.year) / f"{selected.problem.problem_id}. {title}"


def folder_from_state(repo_root: Path, item: dict[str, Any], run_date: date, selected: Selected) -> Path:
    stored_folder = item.get("folder")
    if isinstance(stored_folder, str) and stored_folder:
        return repo_root / stored_folder
    return folder_for(repo_root, run_date, selected)


def render_readme(selected: Selected, run_date: date) -> str:
    problem = selected.problem
    company_tags = ", ".join(problem.companies)
    frequency_rows = "\n".join(
        f"- **{company} frequency:** {problem.frequencies[company]:.6g}"
        for company in problem.companies
    )
    return f"""<!-- daily-problem: date={run_date.isoformat()} company={selected.company} id={problem.problem_id} -->
# {problem.problem_id}. {selected.details.title}

- **Difficulty:** {selected.details.difficulty}
- **Acceptance:** {problem.acceptance}
- **Company lists:** {company_tags}
- **Selected for:** {selected.company}
- **Generated:** {run_date.isoformat()} (Africa/Cairo)
{frequency_rows}

## LeetCode

[{html.escape(problem.link)}]({html.escape(problem.link)})

## Problem statement

{selected.details.content.strip()}
"""


def load_state(path: Path) -> dict[str, Any]:
    if not path.exists():
        return {"version": 1, "runs": {}}
    try:
        state = json.loads(path.read_text(encoding="utf-8"))
    except (OSError, json.JSONDecodeError) as exc:
        raise GeneratorError(f"Cannot read state file {path}: {exc}") from exc
    if state.get("version") != 1 or not isinstance(state.get("runs"), dict):
        raise GeneratorError(f"Unsupported state format in {path}")
    return state


def save_state(path: Path, state: dict[str, Any]) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    descriptor, temp_name = tempfile.mkstemp(prefix=f".{path.name}.", dir=path.parent)
    try:
        with os.fdopen(descriptor, "w", encoding="utf-8") as handle:
            json.dump(state, handle, indent=2, sort_keys=True)
            handle.write("\n")
        os.replace(temp_name, path)
    finally:
        if os.path.exists(temp_name):
            os.unlink(temp_name)
        for sidecar in path.parent.glob(f"._.{path.name}.*"):
            sidecar.unlink(missing_ok=True)
        (path.parent / f"._{path.name}").unlink(missing_ok=True)


def remove_generated_sidecars(target: Path) -> None:
    """Remove AppleDouble files created for this generated folder on external disks."""
    for sidecar in target.glob("._*"):
        sidecar.unlink(missing_ok=True)
    (target.parent / f"._{target.name}").unlink(missing_ok=True)


def launchd_poll_interval_seconds() -> int:
    return 60 * 60


def latest_scheduled_date(now: datetime, schedule_hour: int) -> date:
    scheduled_today = datetime.combine(now.date(), time(schedule_hour), tzinfo=TIMEZONE)
    return now.date() if now >= scheduled_today else now.date() - timedelta(days=1)


def scheduled_run_is_due(state: dict[str, Any], now: datetime, schedule_hour: int) -> bool:
    today_key = now.date().isoformat()
    if state["runs"].get(today_key, {}).get("status") == "complete":
        return False
    latest_key = latest_scheduled_date(now, schedule_hour).isoformat()
    return state["runs"].get(latest_key, {}).get("status") != "complete"


def _selected_from_plan(
    item: dict[str, Any],
    catalog: dict[str, Problem],
    fetcher: Callable[[Problem], Details],
    validate_go: bool,
) -> Selected:
    problem_id = str(item["id"])
    problem = catalog.get(problem_id)
    if not problem:
        raise GeneratorError(f"Planned problem {problem_id} is no longer in the CSV files")
    details = fetcher(problem)
    return Selected(
        company=item["company"],
        problem=problem,
        details=details,
        solution=build_solution(details.go_code, validate=validate_go),
    )


def generate_pair(
    repo_root: Path,
    run_date: date,
    state_path: Path,
    *,
    dry_run: bool = False,
    fetcher: Callable[[Problem], Details] = fetch_details,
    validate_go: bool = True,
) -> list[Selected]:
    catalog = load_catalog(repo_root)
    state = load_state(state_path)
    date_key = run_date.isoformat()
    run = state["runs"].get(date_key)
    if run and run.get("status") == "complete":
        print(f"Daily pair already generated for {date_key}.")
        return []

    existing = discover_existing_ids(repo_root)
    if run and run.get("status") == "planned":
        selected = [
            _selected_from_plan(item, catalog, fetcher, validate_go)
            for item in run["problems"]
        ]
    else:
        amazon = select_problem(
            catalog, "Amazon", existing, fetcher=fetcher, validate_go=validate_go
        )
        microsoft = select_problem(
            catalog,
            "Microsoft",
            existing | {amazon.problem.problem_id},
            fetcher=fetcher,
            validate_go=validate_go,
        )
        selected = [amazon, microsoft]
        run = {
            "status": "planned",
            "problems": [
                {
                    "company": item.company,
                    "id": item.problem.problem_id,
                    "folder": str(folder_for(repo_root, run_date, item).relative_to(repo_root)),
                    "generated": False,
                }
                for item in selected
            ],
        }

    if dry_run:
        for item in selected:
            print(
                f"{item.company}: {item.problem.problem_id}. {item.details.title} "
                f"({item.problem.link})"
            )
        return selected

    state["runs"][date_key] = run
    save_state(state_path, state)
    for index, item in enumerate(selected):
        target = folder_from_state(repo_root, run["problems"][index], run_date, item)
        marker = (
            f"daily-problem: date={date_key} company={item.company} "
            f"id={item.problem.problem_id}"
        )
        if target.exists():
            readme = target / "README.md"
            if not readme.is_file() or marker not in readme.read_text(encoding="utf-8"):
                raise GeneratorError(f"Refusing to overwrite existing folder: {target}")
        else:
            target.parent.mkdir(parents=True, exist_ok=True)
            with tempfile.TemporaryDirectory(prefix="solvendo-problem-") as temp_dir:
                staged = Path(temp_dir) / target.name
                staged.mkdir()
                (staged / "README.md").write_text(
                    render_readme(item, run_date), encoding="utf-8"
                )
                (staged / "solution.go").write_text(item.solution, encoding="utf-8")
                shutil.move(str(staged), str(target))
        remove_generated_sidecars(target)
        run["problems"][index]["generated"] = True
        save_state(state_path, state)
        print(f"Created {target.relative_to(repo_root)}")
    run["status"] = "complete"
    run["completed_at"] = datetime.now(TIMEZONE).isoformat(timespec="seconds")
    save_state(state_path, state)
    return selected


def parse_now(value: str | None) -> datetime:
    if value is None:
        return datetime.now(TIMEZONE)
    parsed = datetime.fromisoformat(value)
    if parsed.tzinfo is None:
        parsed = parsed.replace(tzinfo=TIMEZONE)
    return parsed.astimezone(TIMEZONE)


def build_parser() -> argparse.ArgumentParser:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--repo-root",
        type=Path,
        default=Path(__file__).resolve().parent.parent,
        help="Repository root (defaults to the parent of automation/).",
    )
    parser.add_argument("--dry-run", action="store_true", help="Select and validate without writing.")
    parser.add_argument(
        "--scheduled",
        action="store_true",
        help="Apply launch-agent missed-run and idempotency rules.",
    )
    parser.add_argument("--schedule-hour", type=int, default=9)
    parser.add_argument(
        "--now",
        help="Override the current ISO date/time for deterministic testing.",
    )
    parser.add_argument("--skip-go-check", action="store_true", help=argparse.SUPPRESS)
    return parser


def main(argv: list[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    if not 0 <= args.schedule_hour <= 23:
        raise GeneratorError("--schedule-hour must be between 0 and 23")
    repo_root = args.repo_root.resolve()
    state_path = repo_root / "automation" / "state.json"
    now = parse_now(args.now)
    state = load_state(state_path)
    if args.scheduled and not scheduled_run_is_due(state, now, args.schedule_hour):
        print("No missed or current daily run is due.")
        return 0
    generate_pair(
        repo_root,
        now.date(),
        state_path,
        dry_run=args.dry_run,
        validate_go=not args.skip_go_check,
    )
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except GeneratorError as error:
        print(f"error: {error}", file=sys.stderr)
        raise SystemExit(1)
