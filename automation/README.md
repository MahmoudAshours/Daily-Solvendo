# Daily LeetCode problems

The generator creates one Amazon and one Microsoft problem each day under the
current year. It chooses the highest-frequency unsolved free problem in each CSV,
fetches the full statement and Go starter from LeetCode, and never overwrites an
existing numbered problem folder.

## Commands

Preview today's pair without writing files:

```sh
python3 automation/daily_problem.py --dry-run
```

Generate today's pair manually:

```sh
python3 automation/daily_problem.py
```

Install or refresh the macOS launch agent:

```sh
./automation/install_launch_agent.sh
```

The launch agent polls hourly and the generator only creates a pair after the
09:00 Africa/Cairo cutoff. `RunAtLoad` also invokes the generator after login;
its state file prevents duplicate runs and allows a missed scheduled run to be
created after startup without backlogging older days. Logs are written to
`~/Library/Logs/solvendo/`.

Each `solution.go` has a marked LeetCode section and a local `main()` harness.
Copy only the marked solution section into the LeetCode editor.
