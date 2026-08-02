from __future__ import annotations

import csv
import json
import sys
import tempfile
import unittest
from datetime import datetime
from pathlib import Path


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
        go_code="func solve(value int) int {\n    \n}",
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

    def test_schedule_handles_daily_and_missed_runs(self) -> None:
        complete_yesterday = {
            "version": 1,
            "runs": {"2026-08-01": {"status": "complete"}},
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
                validate_go=False,
            )
            self.assertEqual([item.problem.problem_id for item in selected], ["2", "3"])
            amazon_dir = fixture.root / "2026" / "2. Amazon Pick"
            microsoft_dir = fixture.root / "2026" / "3. Shared"
            self.assertTrue((amazon_dir / "README.md").is_file())
            self.assertTrue((microsoft_dir / "solution.go").is_file())
            readme = (microsoft_dir / "README.md").read_text(encoding="utf-8")
            self.assertIn("Amazon, Microsoft", readme)
            self.assertIn("<p>Shared statement</p>", readme)
            self.assertEqual(
                daily.generate_pair(
                    fixture.root,
                    daily.date(2026, 8, 2),
                    state_path,
                    fetcher=fetch,
                    validate_go=False,
                ),
                [],
            )
            state = json.loads(state_path.read_text(encoding="utf-8"))
            self.assertEqual(state["runs"]["2026-08-02"]["status"], "complete")


if __name__ == "__main__":
    unittest.main()
