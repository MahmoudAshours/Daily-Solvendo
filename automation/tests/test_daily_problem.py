from __future__ import annotations

import csv
import io
import json
import sys
import tempfile
import unittest
from contextlib import redirect_stdout
from datetime import datetime
from pathlib import Path
from unittest import mock


AUTOMATION_DIR = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(AUTOMATION_DIR))

import daily_problem as daily  # noqa: E402


HEADERS = [
    "ID",
    "Title",
    "Acceptance",
    "Difficulty",
    "Frequency",
    "Leetcode Question Link",
]


def details(title: str) -> daily.Details:
    return daily.Details(
        title=title,
        content=f"<p>{title} statement</p>",
        difficulty="Medium",
        go_code="func solve(value int) int {\n\n}",
    )


def daily_challenge(
    run_date: daily.date, validate_go: bool = True
) -> daily.Selected:
    problem = daily.Problem(
        problem_id="5",
        title="Daily Challenge",
        acceptance="55.0%",
        difficulty="Hard",
        link="https://leetcode.com/problems/daily-challenge",
    )
    problem_details = details(problem.title)
    return daily.Selected(
        company=daily.DAILY_SOURCE,
        problem=problem,
        details=problem_details,
        solution=daily.build_solution(problem_details.go_code, validate=validate_go),
    )


class RepositoryFixture:
    def __init__(self, root: Path) -> None:
        self.root = root
        (root / "automation").mkdir()
        (root / "automation" / "state.json").write_text(
            '{"runs": {}, "version": 1}\n', encoding="utf-8"
        )

    def write_company(self, filename: str, rows: list[list[str]]) -> None:
        with (self.root / filename).open("w", newline="", encoding="utf-8") as handle:
            writer = csv.writer(handle)
            writer.writerow(HEADERS)
            writer.writerows(rows)

    def populate(self) -> None:
        self.write_company(
            "amazon_6months.csv",
            [
                ["1", "Paid First", "50%", "Medium", "9", "https://leetcode.com/problems/paid-first"],
                ["2", "Amazon Pick", "60%", "Easy", "8", "https://leetcode.com/problems/amazon-pick"],
                ["3", "Shared", "40%", "Hard", "7", "https://leetcode.com/problems/shared"],
            ],
        )
        self.write_company(
            "microsoft_6months.csv",
            [
                ["3", "Shared", "40%", "Hard", "10", "https://leetcode.com/problems/shared"],
                ["4", "Microsoft Pick", "70%", "Easy", "8", "https://leetcode.com/problems/microsoft-pick"],
            ],
        )


class DailyProblemTests(unittest.TestCase):
    def test_catalog_merges_company_membership_and_trims_links(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            catalog = daily.load_catalog(fixture.root)
            self.assertEqual(catalog["3"].companies, ["Amazon", "Microsoft"])
            self.assertEqual(catalog["2"].slug, "amazon-pick")

    def test_existing_numbered_folders_are_discovered_recursively(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            root = Path(temp_dir)
            (root / "2024" / "IVQs" / "2. Amazon Pick").mkdir(parents=True)
            (root / "notes" / "3. Not a problem year").mkdir(parents=True)
            self.assertEqual(daily.discover_existing_ids(root), {"2"})

    def test_selection_skips_unavailable_and_uses_frequency_order(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            catalog = daily.load_catalog(fixture.root)

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            selected = daily.select_problem(
                catalog, "Amazon", set(), fetcher=fetch, validate_go=False
            )
            self.assertEqual(selected.problem.problem_id, "2")

    def test_solution_has_local_main_markers_and_compiles(self) -> None:
        solution = daily.build_solution(
            """/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

}"""
        )
        self.assertIn("package main", solution)
        self.assertIn("type ListNode struct", solution)
        self.assertIn("panic(\"TODO\")", solution)
        self.assertIn("func main()", solution)
        self.assertIn("LEETCODE SOLUTION BEGIN", solution)

    def test_solution_preserves_signature_and_adapts_empty_body(self) -> None:
        go_code = "func solve(value int) int {\n\n}"
        solution = daily.build_solution(go_code, validate=False)
        self.assertIn("func solve(value int) int", solution)
        self.assertIn("panic(\"TODO\")", solution)

    def test_schedule_handles_daily_and_missed_runs(self) -> None:
        complete_yesterday = {
            "version": 1,
            "runs": {
                "2026-08-01": {
                    "status": "complete",
                    "problems": [
                        {"company": source} for source in daily.EXPECTED_SOURCES
                    ],
                }
            },
        }
        before = datetime(2026, 8, 2, 7, 0, tzinfo=daily.TIMEZONE)
        after = datetime(2026, 8, 2, 10, 0, tzinfo=daily.TIMEZONE)
        self.assertFalse(daily.scheduled_run_is_due(complete_yesterday, before, 9))
        self.assertTrue(daily.scheduled_run_is_due(complete_yesterday, after, 9))
        self.assertTrue(
            daily.scheduled_run_is_due(
                {"version": 1, "runs": {}}, before, 9
            )
        )

    def test_generated_appledouble_sidecars_are_removed(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            parent = Path(temp_dir)
            target = parent / "937. Problem"
            target.mkdir()
            (target / "README.md").write_text("readme", encoding="utf-8")
            (target / "._README.md").write_text("metadata", encoding="utf-8")
            (parent / "._937. Problem").write_text("metadata", encoding="utf-8")
            daily.remove_generated_sidecars(target)
            self.assertFalse((target / "._README.md").exists())
            self.assertFalse((parent / "._937. Problem").exists())

    def test_launchd_uses_hourly_polling_for_cairo_schedule(self) -> None:
        self.assertEqual(daily.launchd_poll_interval_seconds(), 3600)

    def test_cli_get_generates_missing_pair_and_prints_summary(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            original_generate = daily.generate_pair

            def generate_with_fake_fetch(*args: object, **kwargs: object) -> list[daily.Selected]:
                kwargs["fetcher"] = fetch
                kwargs["daily_fetcher"] = daily_challenge
                kwargs["validate_go"] = False
                return original_generate(*args, **kwargs)

            output = io.StringIO()
            with mock.patch.object(daily, "generate_pair", side_effect=generate_with_fake_fetch):
                with redirect_stdout(output):
                    result = daily.main(
                        [
                            "--repo-root",
                            str(fixture.root),
                            "--now",
                            "2026-08-02T10:00:00+03:00",
                            "get",
                        ]
                    )
            rendered = output.getvalue()
            self.assertEqual(result, 0)
            self.assertIn("Daily problems for 2026-08-02 (complete)", rendered)
            self.assertIn("Amazon: 2. Amazon Pick", rendered)
            self.assertIn("Microsoft: 3. Shared", rendered)
            self.assertIn("LeetCode Daily: 5. Daily Challenge", rendered)
            self.assertIn("Difficulty:", rendered)
            self.assertIn("LeetCode:", rendered)
            self.assertIn("README:", rendered)
            self.assertIn("Solution:", rendered)

    def test_cli_get_shows_existing_pair_without_network(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            state_path = fixture.root / "automation" / "state.json"

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )
            output = io.StringIO()
            with mock.patch.object(
                daily.urllib.request,
                "urlopen",
                side_effect=AssertionError("existing get must not contact LeetCode"),
            ):
                with redirect_stdout(output):
                    result = daily.main(
                        [
                            "--repo-root",
                            str(fixture.root),
                            "--now",
                            "2026-08-02T11:00:00+03:00",
                            "get",
                        ]
                    )
            self.assertEqual(result, 0)
            self.assertIn("Daily set already generated", output.getvalue())
            self.assertIn("Daily problems for 2026-08-02 (complete)", output.getvalue())

    def test_cli_status_reports_missing_without_network(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            output = io.StringIO()
            with redirect_stdout(output):
                result = daily.main(
                    [
                        "--repo-root",
                        str(fixture.root),
                        "--now",
                        "2026-08-02T11:00:00+03:00",
                        "status",
                    ]
                )
            self.assertEqual(result, 0)
            self.assertIn("Daily problems for 2026-08-02 (missing)", output.getvalue())
            self.assertIn("No daily set has been generated", output.getvalue())

    def test_legacy_scheduled_flags_still_parse(self) -> None:
        parser = daily.build_parser()
        args = parser.parse_args(["--scheduled", "--schedule-hour", "9"])
        self.assertTrue(args.legacy_scheduled)
        self.assertEqual(args.legacy_schedule_hour, 9)

    def test_pair_generation_is_distinct_idempotent_and_recoverable(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            state_path = fixture.root / "automation" / "state.json"

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            selected = daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )
            self.assertEqual(
                [item.problem.problem_id for item in selected], ["2", "3", "5"]
            )
            amazon_dir = fixture.root / "2026" / "2. Amazon Pick"
            microsoft_dir = fixture.root / "2026" / "3. Shared"
            self.assertTrue((amazon_dir / "README.md").is_file())
            self.assertTrue((microsoft_dir / "solution.go").is_file())
            self.assertTrue(
                (fixture.root / "2026" / "5. Daily Challenge" / "README.md").is_file()
            )
            readme = (microsoft_dir / "README.md").read_text(encoding="utf-8")
            self.assertIn("Amazon, Microsoft", readme)
            self.assertIn("<p>Shared statement</p>", readme)
            self.assertEqual(
                daily.generate_pair(
                    fixture.root,
                    daily.date(2026, 8, 2),
                    state_path,
                    fetcher=fetch,
                    daily_fetcher=daily_challenge,
                    validate_go=False,
                ),
                [],
            )
            state = json.loads(state_path.read_text(encoding="utf-8"))
            self.assertEqual(state["runs"]["2026-08-02"]["status"], "complete")

    def test_existing_two_problem_run_is_upgraded_with_daily_challenge(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            state_path = fixture.root / "automation" / "state.json"

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )
            state = json.loads(state_path.read_text(encoding="utf-8"))
            run = state["runs"]["2026-08-02"]
            daily_item = run["problems"].pop()
            daily_folder = fixture.root / daily_item["folder"]
            (daily_folder / "README.md").unlink()
            (daily_folder / "solution.go").unlink()
            daily_folder.rmdir()
            state_path.write_text(json.dumps(state, indent=2) + "\n", encoding="utf-8")

            selected = daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )

            self.assertEqual(len(selected), 3)
            state = json.loads(state_path.read_text(encoding="utf-8"))
            run = state["runs"]["2026-08-02"]
            self.assertTrue(daily.state_run_is_complete(run))
            self.assertTrue(
                (fixture.root / "2026" / "5. Daily Challenge" / "README.md").is_file()
            )

    def test_daily_challenge_uses_a_source_suffix_on_folder_collision(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            existing = fixture.root / "2026" / "5. Daily Challenge"
            existing.mkdir(parents=True)
            (existing / "notes.txt").write_text("keep", encoding="utf-8")

            selected = daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                fixture.root / "automation" / "state.json",
                fetcher=lambda problem: details(problem.title),
                daily_fetcher=daily_challenge,
                validate_go=False,
            )

            self.assertEqual(len(selected), 3)
            daily_folder = fixture.root / "2026" / "5. Daily Challenge [LeetCode Daily]"
            self.assertTrue((daily_folder / "README.md").is_file())
            self.assertEqual((existing / "notes.txt").read_text(encoding="utf-8"), "keep")

    def test_recovery_uses_planned_folder_when_title_changes(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            state_path = fixture.root / "automation" / "state.json"
            original_title = "Amazon Pick"
            renamed_title = "Amazon Pick Renamed"
            microsoft_title = "Shared"
            calls: dict[str, int] = {"1": 0, "2": 0, "3": 0, "5": 0}

            def fetch(problem: daily.Problem) -> daily.Details:
                calls[problem.problem_id] += 1
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                if problem.problem_id == "2":
                    title = original_title if calls["2"] == 1 else renamed_title
                    return details(title)
                if problem.problem_id == "3":
                    return details(microsoft_title)
                return details("Daily Challenge")

            selected = daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )
            self.assertEqual(
                [item.problem.problem_id for item in selected], ["2", "3", "5"]
            )

            state = json.loads(state_path.read_text(encoding="utf-8"))
            run = state["runs"]["2026-08-02"]
            run["status"] = "planned"
            run["completed_at"] = None
            run["problems"][0]["generated"] = False
            amazon_folder = run["problems"][0]["folder"]
            readme_path = fixture.root / amazon_folder / "README.md"
            solution_path = fixture.root / amazon_folder / "solution.go"
            readme_path.unlink()
            solution_path.unlink()
            (fixture.root / amazon_folder).rmdir()
            state_path.write_text(json.dumps(state, indent=2) + "\n", encoding="utf-8")

            recovered = daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )
            self.assertEqual(
                [item.details.title for item in recovered],
                [renamed_title, microsoft_title, "Daily Challenge"],
            )
            self.assertTrue((fixture.root / amazon_folder / "README.md").is_file())
            self.assertFalse((fixture.root / "2026" / f"2. {renamed_title}").exists())

    def test_recovery_recreates_missing_generated_files_in_existing_folder(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            state_path = fixture.root / "automation" / "state.json"

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )

            state = json.loads(state_path.read_text(encoding="utf-8"))
            run = state["runs"]["2026-08-02"]
            run["status"] = "planned"
            run["completed_at"] = None
            run["problems"][0]["generated"] = False
            amazon_folder = fixture.root / run["problems"][0]["folder"]
            readme_path = amazon_folder / "README.md"
            solution_path = amazon_folder / "solution.go"
            readme_contents = readme_path.read_text(encoding="utf-8")
            solution_path.unlink()
            readme_path.unlink()
            state_path.write_text(json.dumps(state, indent=2) + "\n", encoding="utf-8")

            daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )

            self.assertEqual(readme_path.read_text(encoding="utf-8"), readme_contents)
            self.assertTrue(solution_path.is_file())

    def test_recovery_rejects_unvalidated_existing_solution_without_readme(self) -> None:
        with tempfile.TemporaryDirectory() as temp_dir:
            fixture = RepositoryFixture(Path(temp_dir))
            fixture.populate()
            state_path = fixture.root / "automation" / "state.json"

            def fetch(problem: daily.Problem) -> daily.Details:
                if problem.problem_id == "1":
                    raise daily.CandidateUnavailable("paid-only")
                return details(problem.title)

            daily.generate_pair(
                fixture.root,
                daily.date(2026, 8, 2),
                state_path,
                fetcher=fetch,
                daily_fetcher=daily_challenge,
                validate_go=False,
            )

            state = json.loads(state_path.read_text(encoding="utf-8"))
            run = state["runs"]["2026-08-02"]
            run["status"] = "planned"
            run["completed_at"] = None
            run["problems"][0]["generated"] = False
            amazon_folder = fixture.root / run["problems"][0]["folder"]
            readme_path = amazon_folder / "README.md"
            solution_path = amazon_folder / "solution.go"
            readme_path.unlink()
            original_solution = solution_path.read_text(encoding="utf-8")
            state_path.write_text(json.dumps(state, indent=2) + "\n", encoding="utf-8")

            with self.assertRaisesRegex(daily.GeneratorError, "Refusing to overwrite existing folder"):
                daily.generate_pair(
                    fixture.root,
                    daily.date(2026, 8, 2),
                    state_path,
                    fetcher=fetch,
                    daily_fetcher=daily_challenge,
                    validate_go=False,
                )

            self.assertFalse(readme_path.exists())
            self.assertEqual(solution_path.read_text(encoding="utf-8"), original_solution)


if __name__ == "__main__":
    unittest.main()
