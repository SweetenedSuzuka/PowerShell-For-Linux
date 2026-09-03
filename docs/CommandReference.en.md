# PowerShell Command Reference

This document is aimed at users and lists the commands this program supports. Commands come in two kinds:

- **Go implementation** — behavior is reproduced in Go inside this program, with no external commands called. Most match Windows PowerShell; differences are noted.
- **Mapped Linux commands** — native system tools are called directly (`systemctl`, `ping`, `xclip`, etc.).

Each command gets two examples — one basic, one a bit more involved — with an explanation of what each does. Full parameter and field details live in the matching section of [Command Details](CommandDetails.en.md) (each entry links there).

General notes:

- Directories are displayed Windows-style (the root directory shows as `C:\`), but the real Linux filesystem is what gets operated on.
- `C:` stands for the system drive, which simply is the root directory: `cd C:\tmp` equals `cd /tmp`, and `C:` or `C:\` alone is the root. Other drive letters (`D:`, `E:`, etc.) don't exist on Linux and produce an error.
- The PowerShell 7 command format is the default; use `Set-PSVersion 5.1` to move to the 5 format. Style only affects the banner, `$PSVersionTable`, aliases, and automatic variables: in style 5 the aliases `curl`, `wget`, and `sc` exist while they don't in 7; `$PSVersionTable` shows PSVersion=5.1 with PSEdition=Desktop under 5, versus PSVersion=7 with PSEdition=Core, GitCommitId, OS, and Platform under 7, and variables like `$IsLinux` exist only under 7.
- "Version" records where the command belongs in real PowerShell: "5.1 and 7" means both have it; "7" means only 7 does. This program does not gate commands by style — both work under either style (Test-Json, Join-String, Get-Uptime, for instance, are available under style 5 too).
- "Distro: any" means every distribution can use it. Where a concrete dependency is named (systemd, X11, ping, etc.), only environments meeting that condition can.

---

## Getting in and switching

### Startup
Run `powershell` to enter the shell. You start in the Linux root directory with the prompt `PS C:\>` and the PowerShell 7 command format by default.

### Set-PSVersion (switchstyle)
- Type: Go implementation (this program's own extension).
- Version: extension of this program. Distro: any.
- Function: switch between the 5.X / 7.X command formats; the banner, `$PSVersionTable`, aliases, and automatic variables follow along.
Examples:
- `Set-PSVersion 5.1` — switch to the 5 format.
- `Set-PSVersion 7; $PSVersionTable.PSEdition` — switch back to 7 and show the format name (prints Core).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-PSVersion](CommandDetails.en.md#set-psversion).

### Get-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: returns the host object. InstanceId is a random UUID for this session; the locale follows the interface language (zh-CN for unregistered languages).
Examples:
- `Get-Host` — shows the host object.
- `Get-Host | Select-Object Name,InstanceId` — just the name and session identifier.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Host](CommandDetails.en.md#get-host).

---

## Directory navigation

### Get-Location (pwd, gl)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows the current directory, bash's `pwd`.
- Difference from Windows PowerShell: prints the path string directly (PowerShell renders a PathInfo object as a Path table).
Examples:
- `pwd` — shows the current directory (bash's `pwd`).
- `Set-Location /tmp; Get-Location` — moves to /tmp first, then shows it (bash's `cd /tmp; pwd`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Location](CommandDetails.en.md#get-location).

### Set-Location (cd, sl, chdir)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: changes directory, bash's `cd`.
- Same as Windows PowerShell.
Examples:
- `cd /etc` — moves to /etc (bash's `cd /etc`).
- `Set-Location $HOME; Get-Location` — moves home and confirms (bash's `cd ~; pwd`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Location](CommandDetails.en.md#set-location).

### Push-Location / Pop-Location
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: remembers the current directory on a stack before switching; Pop returns to it. Bash's `pushd` / `popd`.
- Difference from Windows PowerShell: the `pushd` / `popd` aliases are not implemented; use the full names.
Examples:
- `Push-Location /etc; Pop-Location` — notes the current directory, moves to /etc, then goes back (bash's `pushd /etc; popd`).
- `Push-Location /usr; Push-Location /bin; Pop-Location; Get-Location` — pushes two directories in a row then pops one, landing back at /usr (bash's `pushd /usr; pushd /bin; popd; pwd`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Push-Location](CommandDetails.en.md#push-location), [Pop-Location](CommandDetails.en.md#pop-location).

---

## Looking at files and directories

### Get-ChildItem (dir, ls, gci)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: lists directory contents, bash's `ls`; `-Recurse` amounts to `find`.
- Difference from Windows PowerShell: `-Force` has no effect on hidden files.
Examples:
- `dir` — lists the current directory (bash's `ls`).
- `Get-ChildItem /etc -Filter "*.conf" -Recurse | Where-Object Length -gt 1000` — recursively finds .conf files over 1 KB under /etc (roughly bash's `find /etc -name '*.conf' -size +1000c`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-ChildItem](CommandDetails.en.md#get-childitem).

### Get-Item (gi)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: inspects a single file or directory, bash's `stat`.
- Same as Windows PowerShell.
Examples:
- `Get-Item /etc/hostname` — shows hostname's info (bash's `stat /etc/hostname`).
- `(Get-Item /etc/passwd).Length` — gets passwd's size in bytes (bash's `stat -c %s /etc/passwd`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Item](CommandDetails.en.md#get-item).

### Get-ItemProperty (gp)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows a path's properties (name, size, modification time, permissions), the full `stat` output in bash terms.
- Difference from Windows PowerShell: outputs just five fields.
Examples:
- `gp /etc/hostname` — shows hostname's Name/FullName/Length/LastWriteTime/Mode.
- `Get-ItemProperty /etc | Format-List *` — lists all properties of /etc line by line.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-ItemProperty](CommandDetails.en.md#get-itemproperty).

### Set-ItemProperty (sp)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: changes file attributes; currently only the modification time can be changed, bash's `touch`.
- Difference from Windows PowerShell: only the modification time is supported; other attributes are ignored.
Examples:
- `Set-ItemProperty a.txt -Name LastWriteTime -Value $null` — sets a.txt's modification time to now (bash's `touch a.txt`).
- `Set-ItemProperty a.txt -Name LastWriteTime -Value $null; (Get-Item a.txt).LastWriteTime` — change it, then look at the time to confirm.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-ItemProperty](CommandDetails.en.md#set-itemproperty).

### New-Item (ni)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: creates files or directories, bash's `touch` and `mkdir`.
- Same as Windows PowerShell.
Examples:
- `New-Item -ItemType Directory newdir` — creates directory newdir (bash's `mkdir newdir`).
- `New-Item -ItemType Directory -Force a/b/c; New-Item a/b/c/f.txt` — creates nested directories and a file inside them (bash's `mkdir -p a/b/c && touch a/b/c/f.txt`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [New-Item](CommandDetails.en.md#new-item).

### Remove-Item (rm, del, erase, rd, rmdir, ri)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: deletes files or directories, bash's `rm`.
- Difference from Windows PowerShell: nonexistent paths are silently ignored (PowerShell raises a "path not found" error).
Examples:
- `rm old.txt` — deletes a file (bash's `rm old.txt`).
- `Remove-Item -Recurse -Force build/` — force-deletes a directory together with everything inside (bash's `rm -rf build/`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Remove-Item](CommandDetails.en.md#remove-item).

### Copy-Item (cp, copy, cpi)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: copies, bash's `cp`.
- Same as Windows PowerShell.
Examples:
- `cp a.txt b.txt` — copies a.txt to b.txt (bash's `cp a.txt b.txt`).
- `Copy-Item -Path src/ -Destination dst/ -Recurse` — copies a whole directory (bash's `cp -r src/ dst/`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Copy-Item](CommandDetails.en.md#copy-item).

### Move-Item (mv, move, mi)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: moves or renames, bash's `mv`.
- Difference from Windows PowerShell: moving a directory requires `-Recurse`.
Examples:
- `mv a.txt b.txt` — renames a.txt to b.txt (bash's `mv a.txt b.txt`).
- `Move-Item -Path *.log -Destination logs/ -Recurse` — moves a pile of logs into logs/ (bash's `mv *.log logs/`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Move-Item](CommandDetails.en.md#move-item).

### Rename-Item (ren, rni)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: renames, bash's `mv` (rename within one directory).
- Same as Windows PowerShell.
Examples:
- `ren old.txt new.txt` — renames old.txt to new.txt (bash's `mv old.txt new.txt`).
- `Get-ChildItem *.tmp | ForEach-Object { Rename-Item $_ -NewName ($_.Name -replace ".tmp",".txt") }` — renames every .tmp file to .txt.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Rename-Item](CommandDetails.en.md#rename-item).

### Invoke-Item (ii)
- Type: Go implementation (placeholder).
- Version: 5.1 and 7. Distro: any.
- Function: opens a file with its default application, `xdg-open` territory.
- Difference from Windows PowerShell: currently only prints the path rather than actually opening anything.
Examples:
- `ii README.md` — for now merely prints the path README.md.
- `Invoke-Item -Path ./report.pdf` — prints report.pdf's path.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Invoke-Item](CommandDetails.en.md#invoke-item).

### Get-PSDrive (gdr)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: lists drives (the root directory shows as `C`; style 7 shows `/`), semantically close to bash's `df`.
- Difference from Windows PowerShell: only two exist — the root drive and Env.
Examples:
- `gdr` — lists drives.
- `Get-PSDrive | Select-Object Name,Root,CurrentLocation` — just name, root, and current location.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-PSDrive](CommandDetails.en.md#get-psdrive).

---

## Reading and writing file contents

### Get-Content (cat, type, gc)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: reads file contents, emitting string objects line by line, bash's `cat`; `-Tail` amounts to `tail`, `-TotalCount` to `head`.
- Same as Windows PowerShell.
Examples:
- `cat /etc/hostname` — prints hostname's contents (bash's `cat /etc/hostname`).
- `Get-Content -Path /var/log/syslog -Tail 20` — reads only the last 20 lines (bash's `tail -n 20 /var/log/syslog`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Content](CommandDetails.en.md#get-content).

### Set-Content (sc)
- Type: Go implementation.
- Version: 5.1 and 7 (the `sc` alias exists only in 5). Distro: any.
- Function: writes content into a file (overwrite), bash's `echo ... > file`. `-Encoding` picks the encoding (utf8 / utf8BOM / utf8NoBOM / ascii / unicode / bigendianunicode / utf32 / bigendianutf32; default utf8 without BOM).
- Same as Windows PowerShell.
Examples:
- `Set-Content out.txt "hello"` — writes hello into out.txt (bash's `echo hello > out.txt`).
- `Get-ChildItem -Name | Set-Content filelist.txt` — writes the list of file names into a file (bash's `ls > filelist.txt`).
- `Set-Content utf8.txt "你好" -Encoding utf8BOM` — UTF-8 with BOM, which Windows Notepad recognizes right away.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Content](CommandDetails.en.md#set-content).

### Add-Content
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: appends to the end of a file, bash's `echo ... >> file`.
- Difference from Windows PowerShell: the `ac` alias is not implemented; use the full name.
Examples:
- `Add-Content log.txt "a new line"` — appends a line at the end of log.txt (bash's `echo a new line >> log.txt`).
- `1..10 | Add-Content numbers.txt` — appends 1 through 10, one per line (bash's `seq 1 10 >> numbers.txt`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Add-Content](CommandDetails.en.md#add-content).

### Clear-Content
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: empties a file, bash's `: > file`.
- Difference from Windows PowerShell: silently creates an empty file when the path doesn't exist (PowerShell raises a "path not found" error).
Examples:
- `Clear-Content data.txt` — empties data.txt (bash's `: > data.txt`).
- `Get-ChildItem *.log | Clear-Content` — empties every .log file in the current directory.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Clear-Content](CommandDetails.en.md#clear-content).

### Set-Item (si) / Clear-Item (cli)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: Set-Item writes a file or sets an environment variable; Clear-Item empties a file.
- Difference from Windows PowerShell: Set-Item accepts the `env:` prefix for setting environment variables.
Examples:
- `Set-Item env:MY_VAR "value"` — sets an environment variable (bash's `export MY_VAR=value`).
- `Set-Item -Path config.txt -Value "mode=fast"` — writes content into config.txt (bash's `echo mode=fast > config.txt`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Item](CommandDetails.en.md#set-item).

### Get-FileHash
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any (the program computes hashes itself, no external commands involved).
- Function: computes file hashes, bash's `sha256sum` and `md5sum`.
- Difference from Windows PowerShell: supports fewer algorithms.
Examples:
- `Get-FileHash /etc/hostname -Algorithm MD5` — computes MD5 (bash's `md5sum /etc/hostname`).
- `Get-ChildItem *.iso | Get-FileHash -Algorithm SHA256 | Select-Object Hash,Path` — batch-computes SHA256 (bash's `sha256sum *.iso`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-FileHash](CommandDetails.en.md#get-filehash).

### Select-String (sls)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: finds matching lines in text, bash's `grep`.
- Difference from Windows PowerShell: invalid regexes don't raise errors (they simply fail to match).
Examples:
- `Select-String -Path /etc/os-release "NAME"` — finds lines containing NAME in os-release (bash's `grep NAME /etc/os-release`).
- `Get-Content server.log | Select-String -Pattern "ERROR|WARN" | Select-Object LineNumber,Line` — finds ERROR or WARN lines with line numbers (bash's `grep -nE 'ERROR|WARN' server.log`).
- `Get-Content server.log | Select-String -Pattern "ERROR" -Quiet` — asks only whether any ERROR line exists (bash's `grep -q ERROR server.log`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Select-String](CommandDetails.en.md#select-string).

---

## Path handling

### Test-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: checks whether a path exists, returning True / False, bash's `test -e`. `-PathType` filters by kind (Leaf for files / Container for directories).
Examples:
- `Test-Path /etc/passwd` — checks whether the file is there (bash's `test -e /etc/passwd`).
- `Test-Path /etc/passwd -PathType Container` — checks whether it's a directory (bash's `test -d /etc/passwd`).
- `if (Test-Path config.json) { Get-Content config.json } else { "missing configuration file" }` — read it if present, say so otherwise (bash's `[ -e config.json ] && cat config.json || echo missing configuration file`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Test-Path](CommandDetails.en.md#test-path).

### Resolve-Path / Convert-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: turns paths absolute, bash's `realpath`.
- Difference from Windows PowerShell: Resolve-Path additionally resolves symbolic links.
Examples:
- `Resolve-Path .` — shows the current directory's absolute path (bash's `realpath .`).
- `Resolve-Path ./a/../b` — straightens out the .. into the final path (bash's `realpath ./a/../b`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Resolve-Path](CommandDetails.en.md#resolve-path).

### Split-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: splits paths; `-Leaf` amounts to `basename`, `-Parent` to `dirname`.
- Same as Windows PowerShell.
Examples:
- `Split-Path /etc/hostname -Leaf` — takes the last segment, hostname (bash's `basename /etc/hostname`).
- `Get-ChildItem -Name | ForEach-Object { Split-Path $_ -Parent }` — takes the containing directory of each file in turn.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Split-Path](CommandDetails.en.md#split-path).

### Join-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: joins two path segments.
- Same as Windows PowerShell.
Examples:
- `Join-Path /tmp "test.txt"` — yields `/tmp/test.txt`.
- `Join-Path $HOME (Join-Path ".config" "app")` — yields `~/.config/app`.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Join-Path](CommandDetails.en.md#join-path).

---

## Processes and services

### Get-Process (ps, gps)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any (reads /proc on Linux).
- Function: lists processes, bash's `ps`.
- Difference from Windows PowerShell: the memory field is called Memory (PowerShell calls it WS); its value is physical memory in bytes, matching PowerShell's WS semantics. `-Name` matches by substring (PowerShell uses exact names or wildcards).
Examples:
- `ps | Select-Object -First 5` — lists the first 5 processes (roughly bash's `ps -ef | head -5`).
- `ps | Select-Object -Skip 1 -First 5` — skips 1 and lists the next 5 processes (roughly bash's `ps -ef | tail -n +2 | head -5`).
- `Get-Process | Where-Object ProcessName -like "*ssh*"` — picks out processes whose names contain ssh (roughly bash's `ps -ef | grep ssh`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Process](CommandDetails.en.md#get-process).

### Stop-Process
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: ends processes, bash's `kill`.
- Same as Windows PowerShell.
Examples:
- `Stop-Process 1234` — ends PID 1234 (bash's `kill 1234`).
- `Get-Process | Where-Object ProcessName -eq "loop" | Stop-Process` — kills every process named loop.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Stop-Process](CommandDetails.en.md#stop-process).

### Start-Process
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: launches programs in the background, bash's `nohup program &`.
- Difference from Windows PowerShell: doesn't take over the new process's input or output.
Examples:
- `Start-Process sleep 10` — starts `sleep 10` in the background.
- `Start-Process -FilePath /usr/bin/python3 -ArgumentList "server.py"` — runs server.py with python in the background (bash's `nohup python3 server.py &`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Start-Process](CommandDetails.en.md#start-process).

### Wait-Process
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: waits for a process to finish, the `wait` of bash scripts.
- Same as Windows PowerShell.
Examples:
- `Wait-Process 1234` — waits for PID 1234 to finish.
- `Start-Process sleep 2; Wait-Process sleep; "sleep finished"` — sleeps 2 seconds in the background and carries on once it's done.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Wait-Process](CommandDetails.en.md#wait-process).

### Start-Sleep (sleep)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: pauses a while, bash's `sleep`.
- Same as Windows PowerShell.
Examples:
- `sleep 1` — pauses for 1 second (bash's `sleep 1`).
- `Start-Sleep -Milliseconds 500; "continue"` — half a second, then onward.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Start-Sleep](CommandDetails.en.md#start-sleep).

### Get-Service (gsv)
- Type: mapped Linux command (`systemctl`).
- Version: 5.1 and 7. Distro: systemd-based distributions (Debian/Ubuntu/Fedora/RHEL/Arch, etc.).
- Function: lists system services; actually calls `systemctl list-units --type=service --all --no-pager`.
Examples:
- `Get-Service | Select-Object -First 5` — lists the first 5 services (like `systemctl list-units --type=service | head -5`).
- `Get-Service | Where-Object Status -eq "active" | Select-Object Name` — lists only running services.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Service](CommandDetails.en.md#get-service).

### Start-Service / Stop-Service / Restart-Service / Resume-Service (sasv / spsv)
- Type: mapped Linux command (`systemctl`; falls back to sudo automatically when permissions fall short).
- Version: 5.1 and 7. Distro: systemd-based + sudo.
- Function: start / stop / restart services.
Examples:
- `Restart-Service sshd` — restarts sshd (like `sudo systemctl restart sshd`).
- `Get-Service nginx | Where-Object Status -ne "active" | ForEach-Object { Start-Service $_.Name }` — starts nginx if it isn't running.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Start-Service](CommandDetails.en.md#start-service).

### Set-Service
- Type: mapped Linux command (`systemctl`, sudo where needed).
- Version: 5.1 and 7. Distro: systemd-based + sudo.
- Function: changes service state and startup behavior.
Examples:
- `Set-Service nginx -Status running` — starts nginx (like `sudo systemctl start nginx`).
- `Set-Service nginx -StartupType automatic` — makes it start at boot (like `sudo systemctl enable nginx`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Service](CommandDetails.en.md#set-service).

### Test-Connection
- Type: mapped Linux command (`ping`).
- Version: 5.1 and 7. Distro: needs ping (bundled with most distributions).
- Function: tests network reachability.
Examples:
- `Test-Connection localhost` — pings localhost 4 times (bash's `ping -c 4 localhost`).
- `Test-Connection -Count 1 -TargetName 8.8.8.8` — pings just once (bash's `ping -c 1 8.8.8.8`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Test-Connection](CommandDetails.en.md#test-connection).

---

## System information

### Get-Date (date)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows the current time, bash's `date`. `-Format` takes a .NET format string (yyyy year, MM month, dd day, HH hour).
Examples:
- `Get-Date` — shows the current time (bash's `date`).
- `Get-Date -Format "yyyy-MM-dd HH:mm:ss"` — displays in a chosen format (bash's `date +"%Y-%m-%d %H:%M:%S"`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Date](CommandDetails.en.md#get-date).

### Set-Date
- Type: mapped Linux command (`sudo date -s`).
- Version: 5.1 and 7. Distro: needs sudo.
- Function: sets the system time.
Examples:
- `Set-Date "2026-01-01 00:00:00"` — sets the system clock to the given value (like `sudo date -s "2026-01-01 00:00:00"`).
- `Set-Date (Get-Date).AddHours(1)` — moves the clock forward an hour.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Date](CommandDetails.en.md#set-date).

### Get-Uptime
- Type: Go implementation.
- Version: 7. Distro: any (reads /proc/uptime on Linux).
- Function: shows how long the system has been up, bash's `uptime`.
- Difference from Windows PowerShell: always 0 on other platforms.
Examples:
- `Get-Uptime` — shows uptime duration (roughly `uptime -p`).
- `Get-Uptime | Select-Object Days,Hours,Minutes` — just days, hours, minutes.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Uptime](CommandDetails.en.md#get-uptime).

### Get-ComputerInfo
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any (reads /etc/os-release and /proc).
- Function: gathers system information (host name, distro, architecture, memory, CPU), roughly bash's `uname -a` plus distro information.
- Difference from Windows PowerShell: trimmed-down fields.
Examples:
- `Get-ComputerInfo | Select-Object OsName,CsName` — OS name and host name.
- `Get-ComputerInfo | Format-List OsName,OsVersion,CsTotalPhysicalMemory` — lists the key items in full.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-ComputerInfo](CommandDetails.en.md#get-computerinfo).

### Get-TimeZone
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows the current time zone, `/etc/timezone` in essence.
- Difference from Windows PowerShell: reads /etc/timezone.
Examples:
- `Get-TimeZone` — shows the current time zone (bash's `cat /etc/timezone`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-TimeZone](CommandDetails.en.md#get-timezone).

### Set-TimeZone
- Type: mapped Linux command (`sudo timedatectl`).
- Version: 5.1 and 7. Distro: needs systemd's timedatectl + sudo.
- Function: changes the time zone.
Examples:
- `Set-TimeZone Asia/Shanghai; Get-TimeZone` — switches to Shanghai time and confirms (like `sudo timedatectl set-timezone Asia/Shanghai; cat /etc/timezone`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-TimeZone](CommandDetails.en.md#set-timezone).

### Get-Culture
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows the current locale, bash's `locale`.
- Difference from Windows PowerShell: the locale follows the interface language, falling back to zh-CN when the language has no registered locale.
Examples:
- `Get-Culture` — shows the locale.
- `(Get-Culture).Name` — just the locale name.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Culture](CommandDetails.en.md#get-culture).

### Get-Clipboard / Set-Clipboard
- Type: mapped Linux command (`xclip` / `xsel`).
- Version: 5.1 and 7. Distro: needs X11 plus xclip or xsel installed.
- Function: reads and writes the clipboard.
Examples:
- `Set-Clipboard "text"; Get-Clipboard` — puts text on the clipboard, then reads it back.
- `Get-Content note.txt | Set-Clipboard` — puts a file's contents on the clipboard (like `xclip -selection clipboard < note.txt`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Clipboard](CommandDetails.en.md#get-clipboard).

### Restart-Computer / Stop-Computer / Rename-Computer
- Type: mapped Linux command (`sudo reboot` / `sudo shutdown -h now` / `sudo hostnamectl set-hostname`).
- Version: 5.1 and 7. Distro: needs sudo.
- Function: reboot, shut down, change host name.
Examples:
- `Rename-Computer myhost` — changes the host name (like `sudo hostnamectl set-hostname myhost`).
- `Restart-Computer` — reboots the machine (save your work first).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Restart-Computer](CommandDetails.en.md#restart-computer).

---

## Output and formatting

### Write-Output (echo, write)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: emits content, bash's `echo`. The difference: what comes out are objects, ready for further piping.
- Same as Windows PowerShell.
Examples:
- `echo "hello"` — prints hello (bash's `echo hello`).
- `Write-Output (1..5) | Where-Object { $_ % 2 -eq 0 }` — prints 1..5 then picks the even numbers (bash's `printf '1\n2\n3\n4\n5\n' | awk '$1%2==0'`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Write-Output](CommandDetails.en.md#write-output).

### Write-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: puts text straight onto the screen without entering the pipeline, bash's `echo` (to the terminal).
- Same as Windows PowerShell.
Examples:
- `Write-Host "50% done"` — shows "50% done" right on screen (like `echo 50% done`).
- `Write-Host -NoNewline "working"; Start-Sleep 1; Write-Host " done"` — shows progress without a newline, appending "done" after a second (bash's `echo -n working; sleep 1; echo ' done'`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Write-Host](CommandDetails.en.md#write-host).

### Clear-Host (cls, clear)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: clears the screen, bash's `clear`.
- Same as Windows PowerShell.
Examples:
- `clear` — clears the screen (bash's `clear`).
- `cls` — same thing via the Windows-style alias.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Clear-Host](CommandDetails.en.md#clear-host).

### Write-Error / Write-Warning / Write-Verbose / Write-Debug / Write-Information
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: emits leveled notices onto the error stream, bash's `echo ... 1>&2`.
- Difference from Windows PowerShell: prefixes are fixed English strings "ERROR/WARNING/VERBOSE/DEBUG".
Examples:
- `Write-Error "something went wrong"` — writes one line onto the error stream (like `echo something went wrong 1>&2`) and sets `$?` to False.
- `if (Test-Path x) { Write-Information "present" } else { Write-Warning "missing" }` — an informational message or a warning depending on the case.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Write-Error](CommandDetails.en.md#write-error).

### Out-File
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: writes output into a file, bash's `> file`; `-Append` corresponds to `>>`. `-Encoding` picks the encoding (same set as Set-Content).
- Same as Windows PowerShell.
Examples:
- `Get-ChildItem -Name > list.txt` — writes file names into list.txt (bash's `ls > list.txt`).
- `Get-Process | Sort-Object CPU -Descending | Out-File -Append proc.log` — appends to a log (roughly bash's `ps -eo pid,comm,%cpu --sort=-%cpu >> proc.log`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Out-File](CommandDetails.en.md#out-file).

### Out-Null
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: throws output away, bash's `> /dev/null`.
- Same as Windows PowerShell.
Examples:
- `Get-ChildItem -Recurse | Out-Null` — walks the tree listing everything while showing nothing (like `ls -R > /dev/null`).
- `1..100000 | Out-Null; "done"` — generates a heap of numbers and tosses them, showing only "done" (like `seq 1 100000 > /dev/null; echo done`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Out-Null](CommandDetails.en.md#out-null).

### Out-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: sends output to the screen (which is the default anyway), plain display in bash terms.
- Same as Windows PowerShell.
Examples:
- `Get-Date | Out-Host` — shows the date (like `date`).
- `Get-ChildItem | Out-Host` — shows the directory contents (like `ls`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Out-Host](CommandDetails.en.md#out-host).

### Out-String
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: formats objects into a single string.
- Same as Windows PowerShell.
Examples:
- `Get-Date | Out-String` — turns a date object into a string.
- `Get-Process | Out-String | Set-Content proc.txt` — saves the process list as text (bash's `ps -ef > proc.txt`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Out-String](CommandDetails.en.md#out-string).

### Format-Table (ft)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: renders as a table, the kind of alignment `column` does in bash.
- Difference from Windows PowerShell: the column-width auto-fit algorithm differs; no ANSI colors and no "directory" header; time columns use a fixed format (PowerShell localizes them).
Examples:
- `Get-ChildItem | ft -AutoSize` — tabulates the current directory.
- `Get-Process | ft ProcessName,Id -AutoSize` — just the process-name and ID columns.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Format-Table](CommandDetails.en.md#format-table).

### Format-List (fl)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: lists properties one per line.
- Same as Windows PowerShell.
Examples:
- `Get-Item /etc/hostname | fl` — lists hostname's properties line by line.
- `Get-Process | fl * | Select-Object -First 10` — lists every property of each process, first ten blocks only.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Format-List](CommandDetails.en.md#format-list).

### Format-Wide (fw)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: arranges values several columns per row.
- Same as Windows PowerShell.
Examples:
- `Get-ChildItem -Name | fw` — lays names out in multiple columns (like `ls | column`).
- `Get-Service -ErrorAction SilentlyContinue | fw Name` — lays service names out in multiple columns.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Format-Wide](CommandDetails.en.md#format-wide).

### Format-Hex
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: hex view of contents, bash's `xxd` / `od -x`. Prints a type label, a full 16-column layout, and an ASCII reference column.
Examples:
- `"abc" | Format-Hex` — shows abc's bytes in hex (like `echo -n abc | xxd`).
- `Format-Hex /var/log/boot.log` — views a file in hex.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Format-Hex](CommandDetails.en.md#format-hex).

---

## Pipeline processing

### Where-Object (?)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: picks objects out of a pipeline, the filtering that `grep` / `awk` do in bash. Takes a script block `{ $_ -gt 5 }` or a bare `property comparison value` form such as `Length -gt 100`.
- Same as Windows PowerShell.
Examples:
- `1..10 | Where-Object { $_ -gt 5 }` — keeps numbers over 5 (bash's `seq 1 10 | awk '$1>5'`).
- `Get-Process | Where-Object CPU -gt 1 | Select-Object ProcessName,CPU` — keeps processes using more than 1 second of CPU.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Where-Object](CommandDetails.en.md#where-object).

### ForEach-Object (%)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: acts on each object in the pipeline, bash's `xargs` or a `for` loop.
- Same as Windows PowerShell.
Examples:
- `1..3 | % { $_ * 2 }` — doubles each number (bash's `printf '1\n2\n3\n' | awk '{print $1*2}'`).
- `Get-ChildItem *.txt | % { Add-Content $_ "END" }` — appends END to every txt file (bash's `for f in *.txt; do echo END >> "$f"; done`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [ForEach-Object](CommandDetails.en.md#foreach-object).

### Select-Object (select)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: picks properties or takes the leading entries, bash's `cut` and `head`.
- Same as Windows PowerShell.
Examples:
- `1..10 | Select-Object -First 3` — the first 3 numbers (bash's `seq 1 10 | head -3`).
- `Get-ChildItem | Select-Object Name,Length | Sort-Object Length -Descending | Select-Object -First 5` — lists files and keeps the 5 largest.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Select-Object](CommandDetails.en.md#select-object).

### Sort-Object (sort)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: sorts, bash's `sort`; `-Descending` corresponds to `sort -r`.
- Difference from Windows PowerShell: stable sorting (PowerShell uses an unstable sort, leaving equal elements in no guaranteed order; this project preserves input order).
Examples:
- `3,1,2 | sort` — sorts into 1 2 3 (bash's `printf '3\n1\n2\n' | sort`).
- `Get-ChildItem | Sort-Object Length -Descending | Select-Object -First 3` — sorts by size, taking the 3 largest.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Sort-Object](CommandDetails.en.md#sort-object).

### Group-Object (group)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: groups by a property and counts, bash's `sort | uniq -c`.
- Same as Windows PowerShell.
Examples:
- `1,1,2,3,3,3 | Group-Object` — counts occurrences of each number (bash's `printf '1\n1\n2\n3\n3\n3\n' | sort | uniq -c`).
- `Get-Process | Group-Object ProcessName | Sort-Object Count -Descending | Select-Object -First 5` — counts each process type, keeping the 5 most numerous.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Group-Object](CommandDetails.en.md#group-object).

### Measure-Object (measure)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: counts, sums, averages, finds maxima and minima, the job of `wc` and awk statistics in bash.
- Difference from Windows PowerShell: Min/Max count numeric inputs only when numbers mix with non-numbers (PowerShell compares non-numbers as strings); hitting a non-number leaves the statistic empty instead of raising a non-terminating error (as PowerShell does).
Examples:
- `Get-Content file.txt | Measure-Object -Line` — counts lines (bash's `wc -l file.txt`).
- `Get-ChildItem | Measure-Object -Property Length -Sum -Average -Maximum` — sums, averages, and maximizes file sizes.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Measure-Object](CommandDetails.en.md#measure-object).

### Measure-Command
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: times a stretch of commands, bash's `time`.
- Same as Windows PowerShell.
Examples:
- `Measure-Command { Start-Sleep -Milliseconds 100 }` — measures how long sleeping 0.1 seconds takes (like `time sleep 0.1`).
- `(Measure-Command { Get-ChildItem -Recurse | Out-Null }).TotalSeconds` — how many seconds a recursive directory walk took.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Measure-Command](CommandDetails.en.md#measure-command).

### Get-Member (gm)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: reveals an object's properties and types, playing the role Python's `dir()` plays. `-MemberType` filters by kind (Property / TypeName, etc.).
- Difference from Windows PowerShell: method members aren't listed, only types and properties.
Examples:
- `Get-Date | Get-Member` — what members a date object has.
- `Get-ChildItem | Select-Object -First 1 | Get-Member` — what members a file object has.
- `Get-Item x.txt | Get-Member -MemberType Property` — properties only (in spirit like inspecting `stat` fields in bash).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Member](CommandDetails.en.md#get-member).

### Get-Unique (gu)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: removes duplicates, bash's `sort -u`.
- Difference from Windows PowerShell: deduplicates on object string form, not necessarily requiring adjacency.
Examples:
- `1,1,2,2,3 | Get-Unique` — deduplicates into 1 2 3 (bash's `printf '1\n1\n2\n2\n3\n' | sort -u`).
- `Get-Content words.txt | Sort-Object | Get-Unique | Select-Object -First 10` — deduplicated words, first ten.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Unique](CommandDetails.en.md#get-unique).

### Compare-Object
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: compares two sets of things for differences, bash's `diff`.
- Same as Windows PowerShell.
Examples:
- `Compare-Object (1,2,3) (2,3,4)` — shows 1 only on the left and 4 only on the right (like `diff <(printf '1\n2\n3\n') <(printf '2\n3\n4\n')`).
- `Compare-Object (Get-ChildItem a -Name) (Get-ChildItem b -Name)` — which file names differ between directories a and b.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Compare-Object](CommandDetails.en.md#compare-object).

### Tee-Object (tee)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: writes results into a file while passing them down the pipeline, bash's `tee`.
- Same as Windows PowerShell.
Examples:
- `Get-ChildItem -Name | Tee-Object list.txt` — file names go into list.txt while still being shown (bash's `ls | tee list.txt`).
- `Get-Process | Tee-Object -Append proc.log | Select-Object -First 5` — appends to a log and takes the first 5 entries (like `ps -ef | tee -a proc.log | head -5`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Tee-Object](CommandDetails.en.md#tee-object).

### Add-Member
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: tacks a temporary property onto an object. No direct bash counterpart.
- Difference from Windows PowerShell: only the NoteProperty type is supported.
Examples:
- `Get-Date | Add-Member -Name tag -Value test | Select-Object tag` — adds a "tag" property to a date object and reads it back.
- `Get-ChildItem | Add-Member -Name source -Value "local" | Select-Object -First 2` — gives each file object a "source" property.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Add-Member](CommandDetails.en.md#add-member).

### New-Object
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: constructs custom objects (`PSObject` / `pscustomobject`); a `-Property` hashtable becomes properties.
- Difference from Windows PowerShell: only `PSObject` / `PSCustomObject` are supported; other types (System.Collections.ArrayList, for example) get an "unsupported" error.
Examples:
- `New-Object PSObject -Property @{name="x"; n=1}` — constructs an object with two properties (equivalent to `[pscustomobject]@{...}`).
- `[pscustomobject]@{ a = 1; b = "x" }` — type literal constructing a custom object directly.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [New-Object](CommandDetails.en.md#new-object).

---

## Data conversion

### ConvertTo-Json / ConvertFrom-Json
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: turns objects into JSON and JSON into objects, bash's `jq`. `-Depth` caps nesting depth (default 2).
- Difference from Windows PowerShell: output uses a fixed 2-space indent.
Examples:
- `ConvertTo-Json @{name="x"; n=1}` — produces (key order preserved, 2-space indent):
  ```
  {
    "name": "x",
    "n": 1
  }
  ```
- `ConvertFrom-Json '{"a":1,"b":[2,3]}' | Select-Object a` — parses JSON and takes a (bash's `echo '{"a":1,"b":[2,3]}' | jq -r .a`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [ConvertTo-Json](CommandDetails.en.md#convertto-json).

### ConvertTo-Csv / ConvertFrom-Csv
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: turns objects into CSV text and CSV text into objects.
- Same as Windows PowerShell.
Examples:
- `Get-ChildItem | ConvertTo-Csv` — turns a file listing into CSV.
- `Get-Process | Select-Object ProcessName,Id | ConvertTo-Csv | Set-Content proc.csv` — saves process names and IDs as a CSV file.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [ConvertTo-Csv](CommandDetails.en.md#convertto-csv).

### ConvertFrom-StringData
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: parses `key=value` text into a hashtable.
- Same as Windows PowerShell.
Examples:
- `ConvertFrom-StringData "a=1"` — yields `{a=1}`.
- `Get-Content app.conf | ConvertFrom-StringData` — reads a configuration file into a key-value table.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [ConvertFrom-StringData](CommandDetails.en.md#convertfrom-stringdata).

### Test-Json
- Type: Go implementation.
- Version: 7. Distro: any.
- Function: checks whether text is valid JSON, bash's `jq -e`.
- Same as Windows PowerShell.
Examples:
- `Test-Json '{"a":1}'` — returns True.
- `Get-Content data.json | Test-Json` — checks data.json's validity.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Test-Json](CommandDetails.en.md#test-json).

### Get-Random
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: rolls random numbers or samples from a range/list, bash's `shuf`.
- Same as Windows PowerShell.
Examples:
- `Get-Random -Minimum 1 -Maximum 100` — a random number from 1 to 99 (bash's `shuf -i 1-100 -n 1`).
- `Get-ChildItem -Name | Get-Random -Count 2` — two random file names (bash's `ls | shuf -n 2`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Random](CommandDetails.en.md#get-random).

### New-Guid
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: generates a UUID, bash's `uuidgen`.
- Same as Windows PowerShell.
Examples:
- `New-Guid` — generates a UUID (like `uuidgen`).
- `1..3 | % { New-Guid }` — generates three UUIDs.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [New-Guid](CommandDetails.en.md#new-guid).

### New-TimeSpan
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: constructs a time span.
- Same as Windows PowerShell.
Examples:
- `New-TimeSpan -Minutes 5` — a 5-minute span.
- `(New-TimeSpan -Hours 2 -Minutes 30).TotalMinutes` — works out that 2 hours 30 minutes is 150 minutes.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [New-TimeSpan](CommandDetails.en.md#new-timespan).

### New-TemporaryFile
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: creates a temporary file, bash's `mktemp`.
- Same as Windows PowerShell.
Examples:
- `New-TemporaryFile` — creates a temp file (like `mktemp`).
- `$f = New-TemporaryFile; Set-Content $f.FullName "tmp"; Remove-Item $f.FullName` — creates a temp file, puts something in it, deletes it.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [New-TemporaryFile](CommandDetails.en.md#new-temporaryfile).

### Join-String
- Type: Go implementation.
- Version: 7. Distro: any.
- Function: joins a series of strings, bash's `paste`.
- Difference from Windows PowerShell: without `-Separator` the default joiner is a space (PowerShell concatenates with nothing in between).
Examples:
- `1..3 | Join-String -Separator ", "` — yields `1, 2, 3`.
- `Get-ChildItem -Name | Join-String -Separator "`n"` — joins file names with line breaks.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Join-String](CommandDetails.en.md#join-string).

---

## Variables and environment

### Get-Variable (gv)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: lists variables (including automatic ones such as PWD, HOME, PID, PSVersionTable), roughly bash's `env`.
- Same as Windows PowerShell.
Examples:
- `gv HOME` — looks at the HOME variable.
- `Get-Variable P* | Select-Object Name,Value` — lists all variables starting with P.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Variable](CommandDetails.en.md#get-variable).

### Set-Variable / New-Variable / Remove-Variable / Clear-Variable (sv / nv / rv / clv)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: sets, creates, removes, empties variables — assignment and `unset` in bash terms.
- Difference from Windows PowerShell: assigning to read-only automatic variables (PID etc.) is refused.
Examples:
- `New-Variable data 42` — creates data=42 (like `data=42`).
- `Set-Variable -Name total -Value (1..100 | Measure-Object -Sum).Sum; Remove-Variable total` — sums 1 through 100 into total, then removes it.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Variable](CommandDetails.en.md#set-variable).

### $env: environment variables
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: reads and writes environment variables, bash's `$VAR` and `export`.
- Same as Windows PowerShell.
Examples:
- `$env:MY_VAR = "v"; $env:MY_VAR` — sets then reads (bash's `export MY_VAR=v; echo $MY_VAR`).
- `$env:PATH = "/usr/bin:$env:PATH"` — prepends to PATH (bash's `export PATH=/usr/bin:$PATH`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [$env: environment variables](CommandDetails.en.md#env).

---

## Aliases

### Get-Alias (gal)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: lists or looks up aliases, bash's `alias`. In style 5 you'll see sc, curl, and wget; not in 7.
- Difference from Windows PowerShell: looking up a nonexistent alias returns empty with `$?` staying True (PowerShell raises a "cannot find alias" error).
Examples:
- `gal ls` — who ls points to.
- `Get-Alias | Where-Object Definition -like "*Get-Content*"` — finds every alias pointing at Get-Content.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Alias](CommandDetails.en.md#get-alias).

### Set-Alias / New-Alias / Remove-Alias / Import-Alias / Export-Alias (sa / na)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: sets, creates, removes aliases, and imports/exports whole batches — bash's `alias` / `unalias`.
- Same as Windows PowerShell.
Examples:
- `Set-Alias ll "Get-ChildItem"` — points ll at Get-ChildItem (bash's `alias ll='ls'`).
- `New-Alias -Name untar -Value tar; untar xf a.tar.gz` — dubs tar "untar", then unpacks with it.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Set-Alias](CommandDetails.en.md#set-alias).

---

## History

### Get-History / Clear-History / Add-History / Invoke-History (history / ghy / ihy)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: view history, clear history, append an entry, replay one — bash's `history` and `!!`.
- Same as Windows PowerShell.
Examples:
- `history` — shows history (bash's `history`).
- `Get-History | Where-Object CommandLine -like "*git*" | Invoke-History` — digs out git-related commands and replays them.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-History](CommandDetails.en.md#get-history).

---

## Networking

### Invoke-WebRequest (iwr; curl and wget also point here in style 5)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: sends HTTP requests and shows the response body, bash's `curl`.
- Difference from Windows PowerShell: the response comes back as plain text (split into lines of string), with no StatusCode or similar attributes.
Examples:
- `iwr https://example.com` — fetches a page (like `curl https://example.com`).
- `Invoke-WebRequest -Uri https://api.example.com/data -Method POST -Body '{"k":1}'` — POSTs some JSON (like `curl -X POST -d '{"k":1}' https://api.example.com/data`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Invoke-WebRequest](CommandDetails.en.md#invoke-webrequest).

### Invoke-RestMethod (irm)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: sends HTTP requests and parses returned JSON into objects, `curl` paired with `jq` in bash terms.
- Difference from Windows PowerShell: non-JSON responses come back as plain text.
Examples:
- `irm https://api.github.com/zen` — fetches and parses one aphorism (like `curl -s https://api.github.com/zen`).
- `$r = Invoke-RestMethod https://jsonplaceholder.typicode.com/todos/1; $r.title` — gets JSON, then takes the title field (like `curl -s ... | jq -r .title`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Invoke-RestMethod](CommandDetails.en.md#invoke-restmethod).

---

## Commands and help

### Get-Command (gcm)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: surveys available commands and locates a given one, bash's `type` / `which`.
- Difference from Windows PowerShell: external commands aren't listed, only built-ins and aliases.
Examples:
- `gcm Get-Content` — what Get-Content is (bash's `type cat`).
- `Get-Command | Where-Object CommandType -eq "Cmdlet" | Select-Object -First 10` — lists the first 10 built-in commands.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Command](CommandDetails.en.md#get-command).

### Get-Help (help, man, gh)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows a command's usage, bash's `man`.
- Aliases: help, man, gh (gh is this program's own extension, absent from PowerShell).
- Difference from Windows PowerShell: shows name, syntax, and aliases only — no detailed explanation.
Examples:
- `help Get-Content` — Get-Content's syntax.
- `Get-Help Get-ChildItem | Out-String` — captures help text as a string.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Get-Help](CommandDetails.en.md#get-help).

### Invoke-Expression (iex)
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: executes a string as a command, bash's `eval`.
- Same as Windows PowerShell.
Examples:
- `iex "1 + 2"` — executes and prints 3.
- `Get-Content commands.txt | Invoke-Expression` — runs commands.txt's contents line by line.

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Invoke-Expression](CommandDetails.en.md#invoke-expression).

### Read-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: reads one line of keyboard input, bash's `read`.
- Same as Windows PowerShell.
Examples:
- `$name = Read-Host "your name"` — prompts, reads a line into $name (like `read -p "your name" name`).
- `if ((Read-Host "continue? (y/n)") -eq "y") { "continuing" } else { "quitting" }` — continue on y, quit otherwise (bash's `read -p "continue? (y/n)" a; [ "$a" = y ] && echo continuing || echo quitting`).

Full parameter and field descriptions in [Command Details](CommandDetails.en.md): [Read-Host](CommandDetails.en.md#read-host).

---

## Special usage

### sudo
`sudo` isn't built in; it rides external-command passthrough: an unknown command gets looked up on the system PATH and run for real. So `sudo apt install`, `sudo systemctl restart nginx`, and friends are written exactly as in bash. What's more, built-in commands needing root (services, shutdown, time zone, time) automatically try sudo on their own.
Examples:
- `sudo apt update` — exactly bash's `sudo apt update`.
- `sudo systemctl restart nginx && Get-Service nginx` — restart nginx under sudo, check service status on success.

### Pipeline chains && and ||
This is PowerShell 7 syntax: the left side succeeded (`$?` true) before the `&&` right side runs; failed before the `||` right side runs.
Examples:
- `Test-Path /etc/passwd && echo "present"` — prints only if it's there (bash's `[ -e /etc/passwd ] && echo present`).
- `Get-Content important.txt 2>$null || echo "read failed"` — prints only on failure (bash's `cat important.txt 2>/dev/null || echo read failed`).

### Redirection
- `>` writes output into a file (overwrite), `>>` appends, `2>` sends the error stream to a file, `2>$null` discards errors.
Examples:
- `Get-ChildItem -Name > list.txt` — writes output into list.txt (bash's `ls > list.txt`).
- `Get-Content a.txt > out.txt 2> err.txt` — stdout into out.txt, stderr into err.txt (bash's `cat a.txt > out.txt 2> err.txt`).

### Comparison operators
Comparison operators such as `-eq`, `-lt`, `-like`, `-match` work throughout filter scripts (for instance `Where-Object { $_ -lt 100 }`).
- When comparing, the right operand converts to the left operand's type (matching Windows PowerShell): a number on the left compares both sides numerically (`5 -lt "10"` is True); a string on the left compares both sides character by character, case-insensitively (`"5" -lt "10"` is False).

### Capture variable $Matches
After a successful `-match` / `-cmatch`, capture groups land in `$Matches`; `$Matches[1]` takes the first capture group, named groups via `$Matches.groupname`.
Examples:
- `"abc123" -match "(\d+)"; $Matches[1]` — yields 123.
- `"abc123" -match "(?<n>\d+)"; $Matches.n` — named groups go by group name, yielding 123.

### 7.X operators
These operators belong to genuine PowerShell 7, usable under both styles of this program:
- `??` — falls back to the right side only when the left is $null. Example: `$name ?? "anonymous"` (bash's `${name:-anonymous}`).
- `?:` — takes the former when the condition holds, the latter otherwise. Example: `$ok ? "pass" : "fail"`.
- `-f` — fills {N} placeholders in a template with later values. Example: `"{0} files total" -f (Get-ChildItem).Count` (spiritually bash's printf).
- `-in` / `-contains` — membership tests, ranges included. Example: `5 -in 1..10` is True, `1..10 -contains 5` is True.

### Hashtable members
Hashtables take `.keyname` or `["keyname"]`, keys case-insensitive; same-named keys outrank built-in properties. `.Count` is the number of entries, `.Keys` / `.Values` give arrays of keys and values respectively.
Examples:
- `@{a=1;b=2}.Count` — yields 2.
- `@{a=1;b=2}.Keys -join ","` — yields a,b.
