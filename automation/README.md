# Daily LeetCode problems

The generator creates one Amazon and one Microsoft problem each day under the
current year. It chooses two distinct highest-frequency unsolved free problems,
one from each CSV, skips existing or unavailable entries, fetches the full
statement and exact Go starter from LeetCode, and never overwrites an existing
numbered problem folder.

## Commands

Install or refresh both the CLI and macOS launch agent:

```sh
./automation/install_launch_agent.sh
```

The installer links `daily-problem` into `~/.local/bin` and refuses to replace
an unrelated existing command there.

Generate today's pair if it is missing, or display the existing pair:

```sh
daily-problem get
```

Display today's saved state without contacting LeetCode:

```sh
daily-problem status
```

Preview today's pair without writing files:

```sh
daily-problem get --dry-run
```

The launch agent polls hourly and the generator only creates a pair after the
09:00 Africa/Cairo cutoff. `RunAtLoad` also invokes the generator after login;
its state file prevents duplicate runs and allows a missed scheduled run to be
created after startup without backlogging older days. Logs are written to
`~/Library/Logs/solvendo/`.

Each `solution.go` preserves the exact LeetCode function signature, fills empty
bodies with `panic("TODO")`, and includes a marked LeetCode section plus a
local `main()` harness. Copy only the marked solution section into the LeetCode
editor. The automation scaffolds problems only; it does not solve or submit
them.

For compatibility, running `python3 automation/daily_problem.py` without a
subcommand behaves like `daily-problem get`.
