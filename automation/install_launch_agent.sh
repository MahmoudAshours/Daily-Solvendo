#!/bin/zsh
set -euo pipefail

script_dir=${0:A:h}
repo_root=${script_dir:h}
template="$script_dir/com.solvendo.daily-problems.plist.template"
label="com.solvendo.daily-problems"
launch_agents="$HOME/Library/LaunchAgents"
logs_dir="$HOME/Library/Logs/solvendo"
target="$launch_agents/$label.plist"
cli_dir="$HOME/.local/bin"
cli_target="$cli_dir/daily-problem"
cli_source="$script_dir/daily_problem.py"
python_bin=$(command -v python3)

if [[ -e "$cli_target" || -L "$cli_target" ]]; then
    existing_target=$(readlink "$cli_target" 2>/dev/null || true)
    if [[ "$existing_target" != "$cli_source" ]]; then
        echo "Refusing to replace unrelated command: $cli_target" >&2
        exit 1
    fi
fi

mkdir -p "$launch_agents" "$logs_dir" "$cli_dir"
ln -sfn "$cli_source" "$cli_target"

"$python_bin" - "$template" "$target" "$python_bin" "$script_dir/daily_problem.py" "$repo_root" "$logs_dir" <<'PY'
from pathlib import Path
import html
import sys

template, target, python_bin, script, repo_root, logs_dir = sys.argv[1:]
values = {
    "__PYTHON__": python_bin,
    "__SCRIPT__": script,
    "__REPO_ROOT__": repo_root,
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
echo "Installed $cli_target"
echo "The job polls hourly, generates only after 09:00 Africa/Cairo, and uses RunAtLoad for missed runs."
