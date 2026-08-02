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


@dataclass(frozen=True)
class ProblemSummary:
    company: str
    problem_id: str
    title: str
    difficulty: str
    link: str
    readme_path: Path
    solution_path: Path
    generated: bool


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
        target = (repo_root / stored_folder).resolve()
        try:
            target.relative_to(repo_root.resolve())
        except ValueError as exc:
            raise GeneratorError(f"State folder escapes the repository: {stored_folder}") from exc
        return target
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


def ensure_generated_files(target: Path, selected: Selected, run_date: date) -> None:
    readme = target / "README.md"
    solution = target / "solution.go"
    marker = (
        f"daily-problem: date={run_date.isoformat()} company={selected.company} "
        f"id={selected.problem.problem_id}"
    )
    if target.exists():
        if not target.is_dir():
            raise GeneratorError(f"Refusing to overwrite existing folder: {target}")
        if readme.exists():
            if not readme.is_file() or marker not in readme.read_text(encoding="utf-8"):
                raise GeneratorError(f"Refusing to overwrite existing folder: {target}")
        else:
            readme.write_text(render_readme(selected, run_date), encoding="utf-8")
        if not solution.exists():
            solution.write_text(selected.solution, encoding="utf-8")
        elif not solution.is_file():
            raise GeneratorError(f"Refusing to overwrite existing folder: {target}")
        return

    target.parent.mkdir(parents=True, exist_ok=True)
    with tempfile.TemporaryDirectory(prefix="solvendo-problem-") as temp_dir:
        staged = Path(temp_dir) / target.name
        staged.mkdir()
        (staged / "README.md").write_text(render_readme(selected, run_date), encoding="utf-8")
        (staged / "solution.go").write_text(selected.solution, encoding="utf-8")
        shutil.move(str(staged), str(target))


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
    repo_root = repo_root.resolve()
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
        return selected

    state["runs"][date_key] = run
    save_state(state_path, state)
    for index, item in enumerate(selected):
        target = folder_from_state(repo_root, run["problems"][index], run_date, item)
        ensure_generated_files(target, item, run_date)
        remove_generated_sidecars(target)
        run["problems"][index]["generated"] = True
        save_state(state_path, state)
        print(f"Created {target.relative_to(repo_root)}")
    run["status"] = "complete"
    run["completed_at"] = datetime.now(TIMEZONE).isoformat(timespec="seconds")
    save_state(state_path, state)
    return selected


def summaries_for_date(
    repo_root: Path, state_path: Path, run_date: date
) -> tuple[str, list[ProblemSummary]]:
    repo_root = repo_root.resolve()
    state = load_state(state_path)
    run = state["runs"].get(run_date.isoformat())
    if not run:
        return "missing", []
    catalog = load_catalog(repo_root)
    summaries: list[ProblemSummary] = []
    for item in run.get("problems", []):
        problem_id = str(item.get("id", ""))
        problem = catalog.get(problem_id)
        if not problem:
            raise GeneratorError(f"State problem {problem_id!r} is no longer in the CSV files")
        stored_folder = item.get("folder")
        if not isinstance(stored_folder, str) or not stored_folder:
            raise GeneratorError(f"State problem {problem_id} has no folder")
        folder = (repo_root / stored_folder).resolve()
        try:
            folder.relative_to(repo_root.resolve())
        except ValueError as exc:
            raise GeneratorError(f"State folder escapes the repository: {stored_folder}") from exc
        readme_path = folder / "README.md"
        solution_path = folder / "solution.go"
        summaries.append(
            ProblemSummary(
                company=str(item.get("company", "Unknown")),
                problem_id=problem.problem_id,
                title=problem.title,
                difficulty=problem.difficulty,
                link=problem.link,
                readme_path=readme_path,
                solution_path=solution_path,
                generated=bool(item.get("generated"))
                and readme_path.is_file()
                and solution_path.is_file(),
            )
        )
    return str(run.get("status", "unknown")), summaries


def selected_summaries(
    repo_root: Path, run_date: date, selected: list[Selected]
) -> list[ProblemSummary]:
    return [
        ProblemSummary(
            company=item.company,
            problem_id=item.problem.problem_id,
            title=item.details.title,
            difficulty=item.details.difficulty,
            link=item.problem.link,
            readme_path=folder_for(repo_root, run_date, item) / "README.md",
            solution_path=folder_for(repo_root, run_date, item) / "solution.go",
            generated=False,
        )
        for item in selected
    ]


def print_summary(run_date: date, status: str, summaries: list[ProblemSummary]) -> None:
    print(f"Daily problems for {run_date.isoformat()} ({status})")
    if not summaries:
        print("No daily pair has been generated.")
        return
    for summary in summaries:
        print(f"\n{summary.company}: {summary.problem_id}. {summary.title}")
        print(f"  Difficulty: {summary.difficulty}")
        print(f"  LeetCode: {summary.link}")
        print(f"  README: {summary.readme_path}")
        print(f"  Solution: {summary.solution_path}")
        if status != "complete" or not summary.generated:
            print(f"  Files ready: {'yes' if summary.generated else 'no'}")


def require_complete_pair(
    run_date: date, status: str, summaries: list[ProblemSummary]
) -> None:
    if (
        status != "complete"
        or len(summaries) != len(COMPANY_FILES)
        or not all(item.generated for item in summaries)
    ):
        raise GeneratorError(f"Daily pair for {run_date.isoformat()} is incomplete")


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
    parser.add_argument(
        "--now",
        help="Override the current ISO date/time for deterministic testing.",
    )
    parser.add_argument("--skip-go-check", action="store_true", help=argparse.SUPPRESS)
    parser.add_argument("--dry-run", dest="legacy_dry_run", action="store_true", help=argparse.SUPPRESS)
    parser.add_argument("--scheduled", dest="legacy_scheduled", action="store_true", help=argparse.SUPPRESS)
    parser.add_argument("--schedule-hour", dest="legacy_schedule_hour", type=int, help=argparse.SUPPRESS)

    commands = parser.add_subparsers(dest="command")
    get_parser = commands.add_parser(
        "get", help="Generate today's pair if missing, then display it."
    )
    get_parser.add_argument(
        "--dry-run", dest="get_dry_run", action="store_true", help="Preview without writing files."
    )
    commands.add_parser("status", help="Display today's saved pair and state.")
    scheduled_parser = commands.add_parser(
        "run-scheduled", help="Run once using scheduler cutoff and catch-up rules."
    )
    scheduled_parser.add_argument("--schedule-hour", type=int, default=9)
    return parser


def main(argv: list[str] | None = None) -> int:
    args = build_parser().parse_args(argv)
    command = args.command
    if command is None:
        command = "run-scheduled" if args.legacy_scheduled else "get"
    repo_root = args.repo_root.resolve()
    state_path = repo_root / "automation" / "state.json"
    now = parse_now(args.now)

    if command == "status":
        status, summaries = summaries_for_date(repo_root, state_path, now.date())
        print_summary(now.date(), status, summaries)
        return 0

    if command == "get":
        dry_run = bool(args.legacy_dry_run or getattr(args, "get_dry_run", False))
        existing_status, existing_summaries = summaries_for_date(
            repo_root, state_path, now.date()
        )
        if dry_run and existing_status == "complete":
            print_summary(now.date(), existing_status, existing_summaries)
            return 0
        selected = generate_pair(
            repo_root,
            now.date(),
            state_path,
            dry_run=dry_run,
            validate_go=not args.skip_go_check,
        )
        if dry_run:
            print_summary(now.date(), "preview", selected_summaries(repo_root, now.date(), selected))
        else:
            status, summaries = summaries_for_date(repo_root, state_path, now.date())
            require_complete_pair(now.date(), status, summaries)
            print_summary(now.date(), status, summaries)
        return 0

    schedule_hour = getattr(args, "schedule_hour", None)
    if schedule_hour is None:
        schedule_hour = args.legacy_schedule_hour if args.legacy_schedule_hour is not None else 9
    if not 0 <= schedule_hour <= 23:
        raise GeneratorError("--schedule-hour must be between 0 and 23")
    state = load_state(state_path)
    if not scheduled_run_is_due(state, now, schedule_hour):
        print("No missed or current daily run is due.")
        return 0
    generate_pair(
        repo_root,
        now.date(),
        state_path,
        validate_go=not args.skip_go_check,
    )
    status, summaries = summaries_for_date(repo_root, state_path, now.date())
    require_complete_pair(now.date(), status, summaries)
    print_summary(now.date(), status, summaries)
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except GeneratorError as error:
        print(f"error: {error}", file=sys.stderr)
        raise SystemExit(1)
