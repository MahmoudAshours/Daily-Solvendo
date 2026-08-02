#!/bin/zsh
set -euo pipefail

script_dir=${0:A:h}
repo_root=${script_dir:h}
template="$script_dir/com.solvendo.daily-problems.plist.template"
label="com.solvendo.daily-problems"
launch_agents="$HOME/Library/LaunchAgents"
logs_dir="$HOME/Library/Logs/solvendo"
target="$launch_agents/$label.plist"
python_bin=$(command -v python3)

mkdir -p "$launch_agents" "$logs_dir"

"$python_bin" - "$template" "$target" "$python_bin" "$script_dir/daily_problem.py" "$repo_root" "$logs_dir" <<'PY'
from pathlib import Path
import html
import sys

template, target, python_bin, script, repo_root, logs_dir = sys.argv[1:]
start_interval = 60 * 60
values = {
    "__PYTHON__": python_bin,
    "__SCRIPT__": script,
    "__REPO_ROOT__": repo_root,
    "__START_INTERVAL__": str(start_interval),
    "__STDOUT_LOG__": str(Path(logs_dir) / "daily-problems.out.log"),
    "__STDERR_LOG__": str(Path(logs_dir) / "daily-problems.err.log"),
}
rendered = Path(template).read_text(encoding="utf-8")
for marker, value in values.items():
    rendered = rendered.replace(marker, html.escape(value, quote=True))
Path(target).write_text(rendered, encoding="utf-8")
PY

plutil -lint "$target"
domain="gui/$(id -u)"
launchctl bootout "$domain/$label" 2>/dev/null || true
launchctl bootstrap "$domain" "$target"
launchctl enable "$domain/$label"
launchctl print "$domain/$label"

echo "Installed $target"
echo "The job polls hourly, generates only after 09:00 Africa/Cairo, and uses RunAtLoad for missed runs."
