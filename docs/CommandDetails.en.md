# PowerShell Command Details

This document gives every command a **complete, parameter-by-parameter, field-by-field** description. Commands fall into two kinds:

- **Go implementation** — behavior is reproduced in Go inside this program, with no external commands called. Parameter tables explain what each parameter means and how it's used.
- **Mapped Linux commands** — native system tools are called directly (`systemctl`, `ping`, `xclip`, etc.). Parameter tables say what each parameter maps to and which ones are ignored.

The user-facing companion is the [Command Reference](CommandReference.en.md); each command there points to its section here.

Notation:

- Types: `switch` (a flag, true when present), `string`, `int`, `object` (any object), `object[]` (object array), `string[]` (string array), `path`, `scriptblock`.
- Positional parameters: parameters marked `（position N）` in their tables receive positional arguments in N order; unmarked ones accept only named form and occupy no position slot. Named arguments claim their own slots first, with remaining positional arguments falling into unclaimed slots in turn; arguments beyond the declared range are kept as extras, which a few commands treat as data when no pipeline feeds them (such as `Select-Object a,b` or `Get-Unique a b c`). Script-block parameters (the FilterScript of `Where-Object { ... }`, for example) are stored before being evaluated, evaluated only when needed.
- Common parameters (-ErrorAction / -ErrorVariable / -Verbose / -Debug / -WhatIf / -Confirm / -OutVariable / -OutBuffer / -WarningAction / -WarningVariable / -InformationAction / -InformationVariable / -PipelineVariable / -ProgressAction / -InputObject / -Encoding): **accepted but ignored by every command**, without a "parameter cannot be found" error. For switch-like ones among these (-Verbose / -Debug / -WhatIf / -Confirm), a value that follows still gets treated as a positional argument; values of the value-taking kind are ignored. The tables below list only each command's own parameters.
- Missing required parameters: a command lacking its necessary parameters (like -Path or -Name) silently outputs nothing rather than raising an error (PowerShell raises a "missing argument" error).
- `$?`: whether the previous command succeeded. Errors raised inside a command set it to False (equivalent to a nonzero exit code in bash).
- Paths: relative paths resolve against the current directory; they display Windows-style (root directory `C:\`) while operating on the real Linux filesystem.
- Drive letters: Linux has no concept of drive letters — `C:` is merely how the system drive is written. Input like `C:\tmp` or `C:/tmp` equals `/tmp`; `C:` or `C:\` alone is the root. Any other letter (`D:`, `E:`, etc.) errors out with a note that on Linux only drive C stands for the system drive.

---

## I. Getting in and switching styles

### Set-PSVersion
- Type: Go implementation (this program's own extension, absent officially).
- Alias: switchstyle.
- Version: extension of this program. Distro: any.
- Function: switch between the 5.X / 7.X command formats.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Version` (position 0) | string | Target format: values starting `5`/`5.1` → style 5; starting `7` → style 7; anything else errors |

- Positional arguments act as -Version.
- Side effects: rewrites `$PSVersionTable` (5 shows PSVersion 5.1 with PSEdition Desktop; 7 shows PSVersion 7 with PSEdition Core, GitCommitId, OS, Platform), the banner, the alias table (5 gains sc/curl/wget), and automatic variables ($IsLinux/$IsWindows/$IsMacOS exist only under 7; referencing them under 5 yields $null).
- Output: none. Prints "Style switched to %s." on success.

### Get-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: returns the host object.
- Parameters: none.
- Output: a `System.Management.Automation.Internal.Host.InternalHost` object with fields:

| Field | Value |
| :--- | :--- |
| `Name` | "ConsoleHost" |
| `InstanceId` | A random UUID for this session (changes each start) |
| `UI` | Host UI object carrying `SupportsVirtualTerminal` (always True) and a RawUI type name |
| `CurrentCulture` | CultureInfo object (LCID/Name/DisplayName), returned from the registry table of the interface language, zh-CN when that language has no registered locale |
| `CurrentUICulture` | Same as CurrentCulture |

### Get-Location
- Type: Go implementation.
- Aliases: pwd, gl.
- Version: 5.1 and 7. Distro: any.
- Function: shows the current directory.
- Parameters: none.
- Output: a `String` — the Windows-style display of the current directory (root `C:\`), bash's `pwd`. Difference from PowerShell: prints the path string directly, not wrapped in a PathInfo table.

### Set-Location
- Type: Go implementation.
- Aliases: cd, sl, chdir.
- Version: 5.1 and 7. Distro: any.
- Function: changes directory.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target directory. Accepts relative paths, `..`, `~`/`~/...` (home). Bash's `cd` equivalent |

- Behavior: nonexistent or non-directory target → error onto stderr, $?=false. On success the current directory changes.

### Push-Location
- Type: Go implementation.
- Companion: Pop-Location.
- Version: 5.1 and 7. Distro: any.
- Function: pushes the current directory onto a stack and switches.
- Difference from Windows: the `pushd` alias is not implemented; use the full name.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Directory to switch to; omitted, only the push happens with no switch |

- Behavior: nonexistent directory → error. Bash's `pushd` counterpart.

### Pop-Location
- Type: Go implementation.
- Companion: Push-Location.
- Version: 5.1 and 7. Distro: any.
- Function: pops the stack back to the previous directory.
- Difference from Windows: the `popd` alias is not implemented; use the full name.
- Parameters: none.
- Behavior: empty stack means nothing happens. Bash's `popd` counterpart.

---

## II. Looking at files and directories

### Get-ChildItem
- Type: Go implementation.
- Aliases: dir, ls, gci.
- Version: 5.1 and 7. Distro: any.
- Function: lists directory contents. Bash's `ls` / `find`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Starting path, wildcards supported (`*`, `?`, `[...]`). Defaults to "." |
| `-Name` | switch | Output name strings only, no full objects (`ls` without `-l`) |
| `-Recurse` | switch | Recurse into subdirectories (`find` / `ls -R`) |
| `-Directory` | switch | Directories only |
| `-File` | switch | Files only |
| `-Filter` | string | Filter file names by wildcard |
| `-Force` | switch | Accepted, no behavioral change (no hidden-file special-casing) |

- Output objects: `System.IO.FileInfo` / `System.IO.DirectoryInfo` with fields:

| Field | Meaning |
| :--- | :--- |
| `Name` | File name |
| `FullName` | Absolute path |
| `Length` | Size in bytes ($null for directories) |
| `LastWriteTime` | Modification time (DateTime) |
| `PSIsContainer` | Whether it's a directory |
| `Mode` | Permission string (such as `drwxr-xr-x`; d marks a directory) |

- Default table columns: Mode / LastWriteTime / Length / Name (the permission, time, size, and name columns of bash `ls -l`).
- Behavior: nonexistent path → skipped silently (matching `ls`, whose error for missing paths gets ignored).

### Get-Item
- Type: Go implementation.
- Alias: gi.
- Version: 5.1 and 7. Distro: any.
- Function: inspects a single file/directory. Bash's `stat`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path, wildcards supported; several allowed (`Get-Item a.txt,b.txt`) |

- Output: a single FileInfo/DirectoryInfo object, fields same as Get-ChildItem.
- Behavior: nonexistent path → error, $?=false.

### Get-ItemProperty
- Type: Go implementation.
- Alias: gp.
- Version: 5.1 and 7. Distro: any.
- Function: emits an object of the path's properties. The full `stat` picture in bash terms.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |
| `-Name` (position 1) | string | Keep only that property (omit to get all five), e.g. `Get-ItemProperty x -Name Length` |

- Output: a `System.Management.Automation.PSCustomObject` with fields Name / FullName / Length / LastWriteTime / Mode.

### Set-ItemProperty
- Type: Go implementation.
- Alias: sp.
- Version: 5.1 and 7. Distro: any.
- Function: modifies file attributes. Currently supports modification time only.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |
| `-Name` (position 1) | string | Property name; `LastWriteTime` sets the modification time to now (bash's `touch`) |
| `-Value` (position 2) | object | Property value; currently ignored, other property names likewise ignored (without error) |

- Behavior: missing file → error.

### New-Item
- Type: Go implementation.
- Alias: ni.
- Version: 5.1 and 7. Distro: any.
- Function: creates files or directories.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Path to create |
| `-ItemType` | string | "Directory" makes a directory (`mkdir`); any other value makes a file (`touch`) |
| `-Force` | switch | Directories go through MkdirAll (`mkdir -p`); existing files don't raise an error |

- Output: the FileInfo/DirectoryInfo object of the newly created item.
- Behavior: existing file without -Force → error (consistent with PowerShell semantics, unlike `touch`'s silent success).

### Remove-Item
- Type: Go implementation.
- Aliases: rm, del, rd, rmdir, ri, erase.
- Version: 5.1 and 7. Distro: any.
- Function: deletes files/directories.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target paths, multiple/wildcards allowed |
| `-Recurse` | switch | Removes a whole directory tree, `rm -r` |
| `-Force` | switch | Forces removal, `rm -f` (files also go through RemoveAll) |

- Behavior: nonexistent path → ignored silently.

### Copy-Item
- Type: Go implementation.
- Aliases: cp, copy, cpi.
- Version: 5.1 and 7. Distro: any.
- Function: copies.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Source paths, multiple allowed |
| `-Destination` (position 1) | path | Destination path |
| `-Recurse` | switch | Copies a whole directory tree, `cp -r` |

- Behavior: copying a directory without -Recurse → error; an already-existing destination directory receives the source inside it (`cp` semantics); with multiple sources the destination must be an existing directory, otherwise error.

### Move-Item
- Type: Go implementation.
- Aliases: mv, move, mi.
- Version: 5.1 and 7. Distro: any.
- Function: moves/renames. Implemented as copy followed by deleting the source.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Source path |
| `-Destination` (position 1) | path | Destination path |
| `-Recurse` | switch | Moves a whole directory tree (same rule as Copy-Item, directories need it) |
| `-Force` | switch | Accepted, no extra effect at present |

- Behavior: same as Copy-Item (directories need -Recurse), with the source deleted once copying succeeds.

### Rename-Item
- Type: Go implementation.
- Aliases: ren, rni.
- Version: 5.1 and 7. Distro: any.
- Function: renames files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Old path |
| `-NewName` (position 1) | string | New name; only the leaf name is taken, full paths accepted |

- Implementation: `os.Rename(old path, new leaf name in the same directory)`, matching bash `mv old new-name-in-same-directory`.

### Invoke-Item
- Type: Go implementation (placeholder).
- Alias: ii.
- Version: 5.1 and 7. Distro: any.
- Function: opens a file with its default application (`xdg-open` territory). **Currently just prints the path without actually calling xdg-open — a placeholder.**

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File to open; for now its path is merely printed |

### Get-PSDrive
- Type: Go implementation.
- Alias: gdr.
- Version: 5.1 and 7. Distro: any.
- Function: lists drives. Semantically close to bash `df`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | List only drives of the given names (like `C`, `/`, `Env`) |

- Output: PSDriveInfo objects with fields:

| Field | Value |
| :--- | :--- |
| `Name` | Root directory shown as "C" (no colon; "/" under style 7); plus "Env" |
| `Root` | "/" |
| `CurrentLocation` | Windows-style display of the current directory |

- Table columns: Name/Root/CurrentLocation.

---

## III. Reading and writing file contents

### Get-Content
- Type: Go implementation.
- Aliases: cat, type, gc.
- Version: 5.1 and 7. Distro: any.
- Function: reads text files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File paths; several allowed (`Get-Content a.txt,b.txt` reads them one by one, with -TotalCount/-Tail applied per file) |
| `-Raw` | switch | Reads whole without splitting into lines (`cat`) |
| `-TotalCount` | int | Only the first N lines (`head -n N`) |
| `-Tail` | int | Only the last N lines (`tail -n N`) |

- Output: one `String` object per line (line endings stripped). Empty files produce no output.
- Behavior: nonexistent path → error, $?=false, matching `cat`'s complaint about missing files.

### Set-Content
- Type: Go implementation.
- Alias: sc (style 5 only).
- Version: 5.1 and 7. Distro: any.
- Function: overwrites files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target file |
| `-Value` (position 1) | object | Content to write; pipeline input takes precedence over it |
| `-Encoding` | string | Encoding: utf8 (default, no BOM) / utf8BOM / utf8NoBOM / ascii / unicode (UTF-16LE) / bigendianunicode / utf32 / bigendianutf32; unknown names treated as UTF-8 |

- Implementation: writes each object's string into the file line by line (overwrite), `echo ... > file` / `tee` territory. Line endings added automatically.

### Add-Content
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: appends to files.
- Difference from Windows: the `ac` alias is not implemented; use the full name.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target file |
| `-Value` (position 1) | object | Content to append; pipeline input takes precedence |

- Implementation: append mode (O_APPEND), `echo ... >> file` territory.

### Clear-Content
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: empties files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target file |

- Implementation: `os.WriteFile(path, empty, 0644)`, `: > file` / `truncate -s 0` territory.

### Set-Item
- Type: Go implementation.
- Alias: si.
- Companion: Clear-Item.
- Version: 5.1 and 7. Distro: any.
- Function: Set-Item writes files or sets environment variables; Clear-Item empties files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path; beginning with `env:` it's read as an environment variable name |
| `-Value` (position 1) | object | Content to write |

- Behavior: `env:NAME` → `os.Setenv(NAME, value)`, `export NAME=value` territory. Ordinary paths → writes the file.

### Clear-Item
- Type: Go implementation.
- Alias: cli.
- Version: 5.1 and 7. Distro: any.
- Function: empties files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path; directories are left alone |

### Get-FileHash
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any (computed by the program itself, no external commands involved).
- Function: computes file hashes.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File path, wildcards supported |
| `-Algorithm` | string | Algorithm: SHA256 (default) / SHA1 / MD5 / SHA512 (SHA2_256/SHA2_512 also recognized) |

- Implementation: pure Go (crypto standard library). Counterparts: `sha256sum` / `md5sum` / `sha1sum` / `sha512sum`.
- Output: a FileHash object with fields Algorithm (uppercase), Hash (lowercase hex), Path.

### Select-String
- Type: Go implementation.
- Alias: sls.
- Version: 5.1 and 7. Distro: any.
- Function: finds matching lines by regex. Bash's `grep -E`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Pattern` (position 0) | string | Regex pattern |
| `-Path` (position 1) | path | File to search |
| `-SimpleMatch` | switch | Literal matching, `grep -F` |
| `-CaseSensitive` | switch | Case-sensitive matching; insensitive by default (as with `grep -i` reversed) |

- Without -Path, searches pipeline input.
- Output: MatchInfo objects with fields LineNumber (from 1), Line, Path, Pattern.
- Behavior: case-insensitive by default; invalid regexes → no match (no error raised).
- LineNumber: file input scans line by line with blank lines counted; each pipeline input object participates in matching as a single line, its LineNumber being the object's ordinal in the stream — non-matching objects occupy numbers too.

---

## IV. Path handling

### Test-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: checks whether a path exists. Returns Bool, bash's `test -e`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target paths, wildcards supported; several allowed (`Test-Path a.txt,b.txt`) |
| `-PathType` | string | Kind filter: `Leaf` files only, `Container` directories only, `Any`/omitted unrestricted |

- Behavior: path exists (including any wildcard hit) → True, else False. `-PathType` filters further by kind atop existence. Nonexistence isn't an error ($? stays true).

### Resolve-Path
- Type: Go implementation.
- Companion: Convert-Path.
- Version: 5.1 and 7. Distro: any.
- Function: converts to absolute paths. Bash's `realpath`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |

- Implementation: Resolve-Path additionally runs `EvalSymlinks` (resolves symbolic links); Convert-Path only cleans.

### Convert-Path
- Type: Go implementation.
- Companion: Resolve-Path.
- Version: 5.1 and 7. Distro: any.
- Function: converts to absolute paths (symbolic links untouched).

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |

### Split-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: splits paths.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |
| `-Leaf` | switch | Last segment only (`basename`) |
| `-Parent` | switch | Directory part only (`dirname`) |
| `-Qualifier` | switch | Drive qualifier (only drive C exists here: absolute paths return `C:`, relative ones return empty) |

- Note: with no switch given the parent directory is output by default; both `-Leaf` and `-Parent` are switches.

### Join-Path
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: joins paths.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Base path |
| `-ChildPath` (position 1) | path | Child path to append |

- Implementation: `filepath.Join(base, child)`.

---

## V. Processes and services

### Get-Process
- Type: Go implementation.
- Aliases: ps, gps.
- Version: 5.1 and 7. Distro: any.
- Function: lists processes. Bash's `ps`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Filters by process name (substring containment, unlike PowerShell's exact-name/wildcard scheme) |
| `-Id` | int[] | Exact lookup by process ID, comma-separated multiples supported; nonexistent PIDs raise a non-terminating error and set $? to false |

- Implementation: on Linux reads /proc/<pid>/comm and /proc/<pid>/stat; other platforms fall back to the current process.
- Output: Process objects with fields:

| Field | Meaning |
| :--- | :--- |
| `Id` | Process ID (int) |
| `ProcessName` | Process name |
| `CPU` | Accumulated CPU time in seconds (utime+stime divided by clock ticks of 100, Double, same semantics as PowerShell's CPU(s)) |
| `Memory` | Physical memory in bytes (RSS pages times page size, semantically PowerShell's WS; PowerShell names that field WS) |

- Table columns: Id/ProcessName/CPU/Memory.

### Stop-Process
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: ends processes. Bash's `kill`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Process name (kills each match by name) |
| `-Id` | int | Process ID |

- Also accepts objects carrying an Id property via the pipeline (as in `Get-Process | Stop-Process`).
- Implementation: numeric arguments kill by PID via `os.FindProcess(pid).Kill()`; otherwise kills each process by name.

### Start-Process
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: launches in the background. Bash's `nohup program &`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilePath` (position 0) | path | Program to launch |
| `-ArgumentList` | string[] | Arguments passed to the program; leftover positional arguments also become arguments |

- Implementation: `exec.Command(...).Start()` (doesn't wait, doesn't touch IO).
- Output: the new process's Process object (Id, process name).

### Wait-Process
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: waits for a process to finish. Bash's `wait`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Process name |
| `-Id` | int | Process ID |

- Implementation: on Linux checks whether /proc/<pid> disappears; by name, polls the process list until it does.

### Start-Sleep
- Type: Go implementation.
- Alias: sleep.
- Version: 5.1 and 7. Distro: any.
- Function: sleeps. Bash's `sleep`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Seconds` (position 0) | int | How many seconds to sleep |
| `-Milliseconds` | int | How many milliseconds to sleep |

### Get-Service
- Type: mapped Linux command (`systemctl`).
- Alias: gsv.
- Version: 5.1 and 7. Distro: systemd-based.
- Function: lists services. Maps to `systemctl list-units --type=service --all --no-pager`.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Filters by service name (wildcards supported), such as `Get-Service -Name ssh*` |

- Implementation: calls external `systemctl` (adding `sudo` automatically when permissions fall short), parsing columns 0/2 of the output (unit name, status).
- Output: ServiceController objects with fields Name (with the .service suffix stripped), Status (active/inactive/...), DisplayName (same as Name). Table columns Status/Name/DisplayName.
- No systemctl → error.

### Start-Service
- Type: mapped Linux command (`systemctl` + sudo).
- Aliases: sasv (Start-Service), spsv (Stop-Service).
- Companions: Stop-Service, Restart-Service, Resume-Service.
- Version: 5.1 and 7. Distro: systemd-based + sudo.
- Function: starts/stops/restarts services.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name; handed to `systemctl <action> <unit>` (.service suffix added automatically) |

- Implementation: `systemctl start/stop/restart <unit>`; a failure under ordinary permissions retries automatically with `sudo`. Counterpart: `sudo systemctl start/stop/restart`.
- Stop-Service, Restart-Service, and Resume-Service share Start-Service's parameters, their actions being stop/restart/start.

### Set-Service
- Type: mapped Linux command (`systemctl` + sudo).
- Version: 5.1 and 7. Distro: systemd-based + sudo.
- Function: changes service state and startup behavior.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name (.service suffix added automatically) |
| `-Status` | string | `running`/`started` → `systemctl start`; `stopped` → `systemctl stop` |
| `-StartupType` | string | `automatic`/`auto` → `systemctl enable`; `disabled` → `systemctl disable` |

- Implementation: maps to `systemctl start/stop/enable/disable` (sudo where needed).

### Test-Connection
- Type: mapped Linux command (`ping`).
- Version: 5.1 and 7. Distro: needs ping.
- Function: network reachability test. Maps to `ping`.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-TargetName` (position 0) | string | Target host; handed to `ping -c <Count>` |
| `-Count` | int | Number of pings, default 4 |

- Implementation: calls external `ping -c <Count> <target>`, returning whether the exit code was 0.
- Output: Bool. No ping found → error.

---

## VI. System information

### Get-Date
- Type: Go implementation.
- Alias: date.
- Version: 5.1 and 7. Distro: any.
- Function: current time or a specified date-time. Bash's `date`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Date` | string | Specified date-time (parsed against the local time zone when none is written), accepting `2006-01-02 15:04:05`, `2006-01-02T15:04:05`, `2006-01-02`, `2006/1/2`, RFC3339 |
| `-Format` | string | .NET format string; conversion rules below |

- `-Format` conversion rules: `yyyy`→`2006`, `yy`→`06`; `M`→`1`, `MM`→`01`, `MMM`→`Jan`, `MMMM`→`January`; `d`→`2`, `dd`→`02`, `ddd`→`Mon`, `dddd`→`Monday`; `H`→`15`, `hh`→`03`, `m`→`4`, `mm`→`04`, `s`→`5`, `ss`→`05`, `tt`→`PM`, `zzz`→`-07:00`.
- Output: without -Format, a DateTime object (current time when -Date is absent); with -Format, a String. Without -Format the rendering follows the interface language — registered languages use their own formats (English looks like `Saturday, 15 August 2026 15:28:08`), unregistered languages fall back to the default language's Chinese format looking like `2026年8月15日星期六 15:28:08`.

### Set-Date
- Type: mapped Linux command (`sudo date -s`).
- Version: 5.1 and 7. Distro: needs sudo.
- Function: sets system time.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Date` (position 0) | string | Time string; handed to `sudo date -s`; omitted, the current time comes back |

### Get-Uptime
- Type: Go implementation.
- Version: 7. Distro: any.
- Function: uptime duration. Bash's `uptime -p`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| (none) | | |

- Implementation: on Linux reads seconds from /proc/uptime; zero elsewhere.
- Output: a TimeSpan object with fields Days/Hours/Minutes/Seconds/TotalSeconds (plus TotalMilliseconds/TotalMinutes/TotalHours). Table columns Days/Hours/Minutes/Seconds.

### Get-ComputerInfo
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: gathers system information. Bash's `uname -a` + distro info.
- Parameters: none.
- Implementation: reads /etc/os-release (NAME, VERSION_ID) and /proc/meminfo (MemTotal), plus runtime GOOS/GOARCH, `os.Hostname()`, and `runtime.NumCPU()`.
- Output: a ComputerInfo object with fields CsName, OsName, OsVersion, OsArchitecture, OsPlatform, CsTotalPhysicalMemory, CsProcessors.

### Get-TimeZone
- Type: Go implementation (reads /etc/timezone).
- Companion: Set-TimeZone.
- Version: 5.1 and 7. Distro: any.
- Function: shows the current time zone.
- Parameters: none.
- Implementation: reads /etc/timezone.
- Output: a TimeZoneInfo object with fields Id, DisplayName.

### Set-TimeZone
- Type: mapped Linux command (`sudo timedatectl`).
- Companion: Get-TimeZone.
- Version: 5.1 and 7. Distro: needs systemd's timedatectl + sudo.
- Function: changes the time zone.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Time-zone name (Asia/Shanghai, for instance); handed to `sudo timedatectl set-timezone` |

### Get-Culture
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: shows the locale. `locale` counterpart.
- Parameters: none.
- Implementation: returns a CultureInfo from the interface language's registry table, falling back to zh-CN when the language has no registered locale (consistent with `$Host.CurrentCulture`).
- Output: a CultureInfo object with fields LCID, Name, DisplayName.

### Get-Clipboard
- Type: mapped Linux command (`xclip` / `xsel`).
- Companion: Set-Clipboard.
- Version: 5.1 and 7. Distro: X11 + xclip/xsel.
- Function: reads the clipboard.
- Parameters: none.
- Implementation: xclip first (`xclip -o -selection clipboard`), otherwise xsel (`xsel -b`); neither present, returns an empty string.

### Set-Clipboard
- Type: mapped Linux command (`xclip` / `xsel`).
- Companion: Get-Clipboard.
- Version: 5.1 and 7. Distro: X11 + xclip/xsel.
- Function: writes the clipboard.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Value` (position 0) | object | Text to write; pipeline input takes precedence |

- Implementation: `xclip -selection clipboard` first, otherwise `xsel -b`; neither present, quietly does nothing.

### Restart-Computer
- Type: mapped Linux command (`sudo reboot`).
- Companions: Stop-Computer, Rename-Computer.
- Version: 5.1 and 7. Distro: needs sudo.
- Function: reboots the machine.
- Parameters: none.
- Implementation: `sudo reboot`.

### Stop-Computer
- Type: mapped Linux command (`sudo shutdown -h now`).
- Companions: Restart-Computer, Rename-Computer.
- Version: 5.1 and 7. Distro: needs sudo.
- Function: shuts down.
- Parameters: none.
- Implementation: `sudo shutdown -h now`.

### Rename-Computer
- Type: mapped Linux command (`sudo hostnamectl set-hostname`).
- Companions: Restart-Computer, Stop-Computer.
- Version: 5.1 and 7. Distro: needs sudo.
- Function: changes the host name.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-NewName` (position 0) | string | New host name; handed to `sudo hostnamectl set-hostname` |

---

## VII. Output and formatting

### Write-Output
- Type: Go implementation.
- Aliases: echo, write.
- Version: 5.1 and 7. Distro: any.
- Function: emits objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Objects to emit, multiple allowed; pipeline input takes precedence |

- Behavior: places objects into the pipeline unchanged — bash's `echo` counterpart that keeps object types.

### Write-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: writes straight to the screen, bypassing the pipeline. `echo` (to the terminal) territory.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Object` (position 0) | object | Objects to show, multiple allowed |
| `-NoNewline` | switch | No trailing line break (`echo -n`) |

- Implementation: joins each object's string with spaces and writes to stdout.

### Clear-Host
- Type: Go implementation.
- Aliases: cls, clear.
- Version: 5.1 and 7. Distro: any.
- Function: clears the screen. Bash's `clear`.
- Parameters: none.
- Implementation: prints the ANSI clear-screen escape sequence.

### Write-Error
- Type: Go implementation.
- Companions: Write-Warning, Write-Verbose, Write-Debug, Write-Information.
- Version: 5.1 and 7. Distro: any.
- Function: writes leveled messages onto stderr/stdout.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Message` (position 0) | string | Message text (-MessageData for Write-Information) |

- Implementation: Error/Warning/Verbose/Debug go to stderr (`echo ... 1>&2`), prefixed "ERROR/WARNING/VERBOSE/DEBUG"; Information goes to stdout. Error additionally sets $?=false.

### Out-File
- Type: Go implementation.
- Alias: the command form of `>`.
- Version: 5.1 and 7. Distro: any.
- Function: writes files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilePath` (position 0) | path | Target file |
| `-Append` | switch | Append mode (`>>`) |
| `-Encoding` | string | Encoding: same set as Set-Content (utf8 default, utf8BOM, ascii, unicode, etc.); no BOM in append mode |

- Implementation: formats input objects and writes them out (overwrite or append). Equivalent to `> file` / `>> file`.

### Out-Null
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: discards output. `> /dev/null` counterpart.
- Parameters: none.

### Out-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: outputs to the screen (the default behavior).
- Parameters: none.
- Implementation: formats input objects and writes to stdout.

### Out-String
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: formats objects into a string.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Stream` | switch | One line per object (`command | Out-String -Stream` ≈ verbatim text); otherwise one whole string |

- Implementation: reuses the object formatter writing into a buffer.

### Format-Table
- Type: Go implementation.
- Alias: ft.
- Version: 5.1 and 7. Distro: any.
- Function: table display. The alignment effect of `column`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Property columns to display |
| `-AutoSize` | switch | Auto-fit column widths (always auto-fits anyway; accepted) |

- Implementation: builds columns from the first object's properties (or its Props for objects without a table definition), right-aligns numeric columns, underlines headers with `----`.

### Format-List
- Type: Go implementation.
- Alias: fl.
- Version: 5.1 and 7. Distro: any.
- Function: one property per line.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Properties to display; `*` means all |

- Implementation: renders `property : value` lines with names right-aligned to the widest; objects separated by a blank line.

### Format-Wide
- Type: Go implementation.
- Alias: fw.
- Version: 5.1 and 7. Distro: any.
- Function: multi-column arrangement.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string | Which property of each object to display |
| `-Column` | int | Column width, default 40 |

- Implementation: lays object strings out across an 80-character width.

### Format-Hex
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: hexadecimal view. `xxd` / `od -x` counterparts.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File to view; emits a `Label: path` section |
| `-InputObject` | object | Piped-in objects, one section per object |

- Implementation: renders each input at 16 bytes per row — a 16-digit offset, a fixed-width byte area, an ASCII reference column on the right (unprintable characters shown as `.`), and a leading `Label:` line per section — the path for files, the type name plus content checksum for piped objects.

---

## VIII. Pipeline processing

### Where-Object
- Type: Go implementation.
- Alias: ?.
- Version: 5.1 and 7. Distro: any.
- Function: filters pipeline objects. The filtering done by bash `grep` / `awk`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilterScript` (position 0) | scriptblock | Script block where `$_` is the current object; kept when the block outputs non-empty with a truthy last value |
| `-Property` | string | Bare-property comparison (as in `-Property Length -gt 100`); a comparison expression as its value acts as the filter instead |
| `-Not` | switch | Negates |
| `-SimpleMatch` | switch | Accepted, no extra effect |

- Filter expression forms:
  - `Where-Object { $_.Length -gt 100 }` (script block).
  - `Where-Object Length -gt 100` (bare property + comparison operator — the parser merges `Length -gt 100` into one comparison expression, evaluated with `Length` resolving to `$_.Length`).

### ForEach-Object
- Type: Go implementation.
- Alias: %.
- Version: 5.1 and 7. Distro: any.
- Function: runs a script block per object. Bash's `xargs` / `for` loop.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Process` (position 0) | scriptblock | Block run against every input object, `$_` bound to it |
| `-MemberName` | string | Takes each object's value of that member (as in `-MemberName Length`) |
| `-Begin` / `-End` | scriptblock | Each runs once (aggregation style), `-Begin` before the loop, `-End` after |

- Implementation: for each input object, runs the block with `$_` bound to it and emits whatever objects the block produces.

### Select-Object
- Type: Go implementation.
- Alias: select.
- Version: 5.1 and 7. Distro: any.
- Function: picks properties or leading/trailing entries. Bash's `cut`, `head`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Keeps only these properties (producing new objects); `*` means all |
| `-ExpandProperty` | string | Outputs that property's own value without wrapping it in an object; array values flatten (as in `Select-Object -ExpandProperty Name`) |
| `-First` | int | First N entries (`head -n N`) |
| `-Last` | int | Last N entries (`tail -n N`) |
| `-Unique` | switch | Deduplicates (`sort -u`, by string, case-sensitive) |

- Behavior: no input with an array at position 0 → treated by array element. `-First 0` / `-Last 0` return empty (explicit 0 is distinct from unset).

### Sort-Object
- Type: Go implementation.
- Alias: sort.
- Version: 5.1 and 7. Distro: any.
- Function: sorts. Bash's `sort`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Sort keys, compared one attribute after another |
| `-Descending` | switch | Descending order (`sort -r`) |
| `-Unique` | switch | Deduplicates (`sort -u`, on sort keys) |
| `-CaseSensitive` | switch | Case-sensitive; insensitive by default |

- Implementation: stable sort; numbers compare by magnitude, strings case-insensitively. `-Unique` deduplicates folding case over sort keys by default, or on original values with `-CaseSensitive`.

### Group-Object
- Type: Go implementation.
- Alias: group.
- Version: 5.1 and 7. Distro: any.
- Function: groups and counts. `sort | uniq -c` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string | Property to group by; none given, groups by object string |
| `-CaseSensitive` | switch | Case-sensitive grouping; insensitive by default |

- Implementation: groups by property value (or object string when absent), case-insensitive by default (merged after folding), with Name taken from the first original value seen in each group.
- Output: GroupInfo objects with fields Name, Count. Table columns Count/Name.

### Measure-Object
- Type: Go implementation.
- Alias: measure.
- Version: 5.1 and 7. Distro: any.
- Function: statistics. `wc` / awk statistics counterparts.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string | Which property to measure |
| `-Sum` | switch | Sum |
| `-Average` | switch | Average |
| `-Minimum` | switch | Minimum |
| `-Maximum` | switch | Maximum |
| `-Line` | switch | Counts lines of input (`wc -l`) |

- Output: a MeasureInfo object whose fields Count, Average, Sum, Maximum, Minimum, StandardDeviation, Property are always present; unrequested statistics hold $null.
- Behavior: in `-Property` mode, Count tallies only objects having that property. Sum/Average become void ($null) upon hitting non-numeric input; Min/Max count numeric inputs only.

### Measure-Command
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: timed execution. Bash's `time`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Expression` (position 0) | scriptblock | The block to time |

- Output: a TimeSpan object with fields Days/Hours/Minutes/Seconds/TotalMilliseconds/TotalSeconds (plus TotalMinutes/TotalHours).

### Get-Member
- Type: Go implementation.
- Alias: gm.
- Version: 5.1 and 7. Distro: any.
- Function: inspects object members. Python's `dir()` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to inspect |
| `-MemberType` | string | Member-type filter (Property / TypeName etc., case-insensitive, all returned by default) |

- Output: PSMemberInfo objects with fields Name, MemberType (TypeName/Property), Definition (the property value as string). Each type lists its type row once, properties deduplicated.

### Get-Unique
- Type: Go implementation.
- Alias: gu.
- Version: 5.1 and 7. Distro: any.
- Function: deduplicates. `sort -u` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Input objects |
| `-AsString` | switch | Accepted, no extra effect |

- Implementation: keeps the first occurrence of each object string.

### Compare-Object
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: compares sets for differences. `diff` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-ReferenceObject` (position 0) | object | Reference set |
| `-DifferenceObject` (position 1) | object | Set to compare against it |
| `-CaseSensitive` | switch | Case-sensitive comparison; insensitive by default |
| `-IncludeEqual` | switch | Also outputs equal entries (`==`), differences only by default |

- Output: PSCustomObjects with fields InputObject, SideIndicator. `<=` reference-set only, `=>` comparison-set only, `==` both sides (needs -IncludeEqual). Case-insensitive by default; output order runs equals first, then right side (=>), then left side (<=). Equal entries display the reference set's values.

### Tee-Object
- Type: Go implementation.
- Alias: tee.
- Version: 5.1 and 7. Distro: any.
- Function: writes to a file while passing output along. `tee` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilePath` (position 0) | path | Output file |
| `-Append` | switch | Append mode (`tee -a`) |

- Implementation: writes every input object into the file while returning it unchanged to the pipeline.

### Add-Member
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: adds properties to objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to receive the property |
| `-MemberType` | string | NoteProperty only; other values ignored |
| `-Name` | string | New property name |
| `-Value` | object | Property value |
| `-Force` | switch | Overwrites even when the property already exists |

- Implementation: copies the object, appends the property, and returns the new object.

### New-Object
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: constructs custom objects. Counterpart of the `[pscustomobject]@{...}` type literal.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-TypeName` (position 0) | string | Type name: currently `PSObject` / `pscustomobject`; other types get an "unsupported" error |
| `-Property` | hashtable | Hashtable whose keys become properties in order |

- Implementation: both `PSObject` and `pscustomobject` construct a `System.Management.Automation.PSCustomObject`; without `-Property` an empty object results.
- The type literal `[pscustomobject]@{ a = 1; b = "x" }` equals `New-Object pscustomobject -Property @{ a = 1; b = "x" }`.

---

## IX. Data conversion

### ConvertTo-Json
- Type: Go implementation.
- Companion: ConvertFrom-Json.
- Version: 5.1 and 7. Distro: any.
- Function: turns objects into JSON. Bash's `jq` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to convert |
| `-Depth` | int | Nesting depth cap (default 2); beyond it arrays/objects print as empty shells (`[]`/`{}`) |

- Rules: single object → object JSON; several objects → array. Property-bearing objects → `{"property": value}`; Hashtable → key-value object; arrays recurse; scalars stay JSON scalars; $null → null. Pretty-printed with a 2-space indent.

### ConvertFrom-Json
- Type: Go implementation.
- Companion: ConvertTo-Json.
- Version: 5.1 and 7. Distro: any.
- Function: turns JSON into objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | JSON text; pipeline input takes precedence |

- Implementation: parses text into PSCustomObjects (properties), arrays, and scalars.

### ConvertTo-Csv
- Type: Go implementation.
- Companion: ConvertFrom-Csv.
- Version: 5.1 and 7. Distro: any.
- Function: turns objects into CSV text.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Input objects |
| `-Property` | string[] | Chooses columns; omitted, takes the first object's property names (single-column "Value" for property-less objects) |

### ConvertFrom-Csv
- Type: Go implementation.
- Companion: ConvertTo-Csv.
- Version: 5.1 and 7. Distro: any.
- Function: turns CSV text into objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | CSV text; pipeline input takes precedence |

- Implementation: first line is the header, remaining lines produce PSCustomObjects.

### ConvertFrom-StringData
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: parses `key=value` text into a Hashtable.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-StringData` (position 0) | string | `key=value` text; named/positional take precedence, pipeline input only as fallback |

- Implementation: goes line by line, skipping blank lines and `#` comments, splitting keys from values at `=` with TrimSpace.

### Test-Json
- Type: Go implementation.
- Version: 7. Distro: any.
- Function: validates JSON. `jq -e` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Json` (position 0) | string | Text to validate; named/positional take precedence, pipeline input only as fallback |

- Output: Bool.

### Get-Random
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: randomness. `shuf` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Samples randomly among these objects (1 by default) |
| `-Minimum` / `-Maximum` | int | Random within [Minimum, Maximum) |
| `-Count` | int | Draws N items randomly from the input |

- Implementation: crypto/rand.

### New-Guid
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: generates UUIDs. `uuidgen` counterpart.
- Parameters: none.
- Implementation: crypto/rand producing RFC 4122 version 4.

### New-TimeSpan
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: constructs time spans.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Seconds` / `-Minutes` / `-Hours` | int | Values per unit, combinable |

- Output: a TimeSpan object with fields Days/Hours/Minutes/Seconds/TotalMilliseconds/TotalSeconds (plus TotalMinutes/TotalHours).

### New-TemporaryFile
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: creates temporary files. `mktemp` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Extension` | string | File suffix |

- Implementation: `os.CreateTemp("", "tmp*"+suffix)`.
- Output: a FileInfo object.

### Join-String
- Type: Go implementation.
- Version: 7. Distro: any.
- Function: joins strings. `paste` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Strings to join |
| `-Separator` (position 1) | string | Separator, default the empty string |
| `-Property` | string | Takes that property's value converted to text before joining |
| `-FormatString` | string | Formats each object with it first (`{0}` is the current object; `{{`/`}}` literal braces) |
| `-OutputPrefix` | string | Text placed ahead of the output |
| `-OutputSuffix` | string | Text appended after the output |
| `-DoubleQuote` | switch | Wraps each object in double quotes |
| `-SingleQuote` | switch | Wraps each object in single quotes |

- Implementation: joins object strings with the separator; without -Separator they concatenate directly.

---

## X. Variables and environment

### Get-Variable
- Type: Go implementation.
- Alias: gv.
- Version: 5.1 and 7. Distro: any.
- Function: lists variables. Bash's `env` / `declare` counterparts.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable names, wildcards and arrays supported |

- Output: PSVariable objects with fields Name, Value. Automatic variables (PWD/HOME/PID/PSVersionTable/LASTEXITCODE/?/PSCommandPath/args/Host/PSEdition/IsLinux/IsWindows/IsMacOS/PSHOME/OFS) appear too.

### Set-Variable
- Type: Go implementation.
- Aliases: sv / nv / rv / clv.
- Companions: New-Variable, Remove-Variable, Clear-Variable.
- Version: 5.1 and 7. Distro: any.
- Function: sets/creates/removes/empties variables.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable name |
| `-Value` (position 1) | object | Value (New-Variable's `-Value` also takes position 1) |
| `-Force` | switch | Lets New-Variable overwrite an existing variable |

- Behavior: creating over an existing variable without -Force → error; assigning to read-only automatic variables (PID etc.) → error. Remove deletes straight from the map; Clear sets $null.

### $env:
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: reads and writes process environment variables.
- Behavior: `$env:NAME` reads via os.Getenv; `$env:NAME = value` writes via os.Setenv (`export` territory). `+=` appends.

---

## XI. Aliases

### Get-Alias
- Type: Go implementation.
- Alias: gal.
- Version: 5.1 and 7. Distro: any.
- Function: lists aliases. Bash's `alias`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Alias names, wildcards/arrays supported |

- Output: AliasInfo objects with fields Name, Definition (target cmdlet). Table columns Name/Definition.
- Note: style 5's alias table contains sc/curl/wget, style 7 doesn't. Looking up a nonexistent alias returns empty with $? staying True (PowerShell raises an error).

### Set-Alias
- Type: Go implementation.
- Aliases: sa / na.
- Companions: New-Alias, Remove-Alias, Import-Alias, Export-Alias.
- Version: 5.1 and 7. Distro: any.
- Function: manages aliases.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Alias name |
| `-Value` (position 1) | string | Target command |
| `-Force` | switch | Lets New-Alias overwrite an existing alias |

- Set-Alias overwrites; New-Alias errors on an existing alias without -Force; Remove-Alias deletes; Export-Alias writes one `name=target` per line; Import-Alias reads them back.

---

## XII. History

### Get-History
- Type: Go implementation.
- Aliases: history / ghy.
- Companions: Clear-History, Add-History, Invoke-History.
- Version: 5.1 and 7. Distro: any.
- Function: lists history. Bash's `history`.
- Parameters: none.
- Output: HistoryInfo objects with fields Id, CommandLine. Table columns Id/CommandLine.

### Clear-History
- Type: Go implementation.
- Companions: Get-History, Add-History, Invoke-History.
- Version: 5.1 and 7. Distro: any.
- Function: clears history. Bash's `history -c`.
- Parameters: none.

### Add-History
- Type: Go implementation.
- Companions: Get-History, Clear-History, Invoke-History.
- Version: 5.1 and 7. Distro: any.
- Function: appends a history entry.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Command text; pipeline input takes precedence |

### Invoke-History
- Type: Go implementation.
- Alias: ihy.
- Companions: Get-History, Clear-History, Add-History.
- Version: 5.1 and 7. Distro: any.
- Function: replays a history entry. Bash's `!!`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Id` | int | History number |
| `-InputObject` (position 0) | object | Command text (omitting it takes the last entry) |

- Behavior: Invoke-History echoes the command before replaying it (parsed and executed through RunSource).

---

## XIII. Networking

### Invoke-WebRequest
- Type: Go implementation.
- Aliases: iwr; curl/wget point here under style 5.
- Version: 5.1 and 7. Distro: any.
- Function: HTTP requests returning response text. `curl` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Uri` (position 0) | string | Request address |
| `-Method` | string | HTTP method, GET by default |
| `-Body` | object | Request body, used as such for POST/PUT/PATCH |

- Implementation: net/http with a 30-second timeout.
- Output: the response text split into Strings by line.

### Invoke-RestMethod
- Type: Go implementation.
- Alias: irm.
- Version: 5.1 and 7. Distro: any.
- Function: HTTP requests with JSON parsing. `curl | jq` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Uri` (position 0) | string | Request address |
| `-Method` | string | HTTP method, GET by default |
| `-Body` | object | Request body, used as such for POST/PUT/PATCH |

- Implementation: after the request, JSON parsing is attempted (reusing ConvertFrom-Json's object construction); non-JSON comes back as text.

---

## XIV. Commands and help

### Get-Command
- Type: Go implementation.
- Alias: gcm.
- Version: 5.1 and 7. Distro: any.
- Function: looks up commands. Bash's `type` / `which`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Command names, arrays supported |

- Output: CommandInfo objects (Name normalized-case, CommandType=Cmdlet/Function/Alias); Alias objects carry Definition besides. Table columns CommandType/Name.
- Difference from Windows: only built-in commands, functions, and aliases are listed, external commands aren't.

### Get-Help
- Type: Go implementation.
- Aliases: help, man (official); gh (this program's extension).
- Version: 5.1 and 7. Distro: any.
- Function: shows command syntax. Bash's `man`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Command name, wildcards supported |

- Output: name, syntax (`command [-param <type>]...`), aliases. No match found checks whether it's an alias and shows its target; failing that, reports not found.
- Difference from Windows: only name, syntax, and aliases are shown — no detailed explanation.

### Invoke-Expression
- Type: Go implementation.
- Alias: iex.
- Version: 5.1 and 7. Distro: any.
- Function: executes strings as commands. Bash's `eval`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Command` (position 0) | string | Source code to execute; named/positional take precedence, pipeline input only as fallback |

- Implementation: goes through `RunSource` (parse + statement-by-statement execution), returning output objects.

### Read-Host
- Type: Go implementation.
- Version: 5.1 and 7. Distro: any.
- Function: reads one line of input. Bash's `read`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Prompt` (position 0) | string | Prompt text |

- Implementation: prints the prompt then reads a line with bufio, trimming the trailing line break.

---

## Special mechanics

### External command passthrough
- Commands matching neither built-ins nor aliases nor functions get looked up on PATH and executed (`os/exec`), with stdin/stdout/stderr passed through.
- Exit codes land in `$LASTEXITCODE`; nonzero exit code → $?=false.
- External commands at the end of a pipeline connect straight to the terminal; mid-pipeline ones have their stdout converted line-by-line into String objects entering the pipeline, their stdin fed the left-side objects' text.
- Not found → stderr reports "The term 'X' is not recognized as the name of a cmdlet, function, script file, or runnable program.", $?=false.

### sudo
- `sudo` is an external command passed straight through to the system sudo. Built-in commands needing root (services/shutdown/time zone/time) call `sudo` internally on their own (trying without sudo first, retrying with `sudo` on failure).

### Pipeline chains && / ||
- PowerShell 7 syntax: the left command's $? being true gates the `&&` right side; false gates the `||` right side.
- Bash's `&&` / `||` counterparts.

### Redirection
- `> file`: overwrite write (Out-File equivalent); `>> file`: append (Out-File -Append).
- `2> file`: sends the error stream to a file; `2>$null`: discards errors (io.Discard).
- `> $null`: discards output.
- External-command redirection: files serve as that command's stdout/stderr.

### Bare operator arguments
- A command positional argument appearing alone as `/`, `-`, `+`, `*`, `%` followed by end-of-statement (as in `Write-Output /`, `cd -`) is taken as a string argument rather than an operator.
- Named parameters followed by these symbols likewise treat them as values (as in `Get-PSDrive -Name /`, `Set-Location -Path /`).

### Comparison operators
- In comparisons like `-eq`/`-ne`/`-lt`/`-gt`/`-le`/`-ge`: the right operand converts to the left operand's type — a numeric left compares both sides numerically (`5 -lt "10"` is True); a string left compares character by character case-insensitively (`"5" -lt "10"` is False); `$null` equals only `$null`.
- `-like`/`-match` are case-insensitive by default, the `-c*` prefix marking sensitive variants.
- The range operator `..` binds tighter than comparisons: `1..3 -gt 1` forms the range `1 2 3` first and filters after, giving `2 3`; `5 -in 1..10` and `1..10 -contains 5` both work.

### Capture variable $Matches
- After a scalar `-match` / `-cmatch` succeeds, capture groups go into the automatic variable `$Matches` (a hashtable).
- Key rules match genuine PowerShell: `0` is the whole match; named groups `(?<name>...)` use the group name as key; unnamed groups `(...)` use their ordinal among unnamed groups (from 1); groups not participating in the match aren't written.
- Failed matches, array left values, and `-notmatch` all leave `$Matches` untouched (preserving the previous value); never having matched, `$Matches` is `$null`.
- Differences: named groups support only `(?<name>...)` (internally converted to Go's `(?P<name>...)`); `(?'name'...)` and lookarounds like `(?<=...)` aren't supported by Go regexes.

### Format operator -f
- `"template {0} {1}" -f value1, value2`: fills placeholders from later values per .NET-style format strings, spiritually bash's printf.
- Supports `{N}`, `{N,width}` (space-aligned, negative width left-aligns), `{N:spec}`; spec accepts `D`/`Dk` (decimal, k digits zero-padded), `X`/`x` (hexadecimal), `Fk` (k decimal places), `Nk` (thousands separators + k decimal places); unknown specs degrade to the plain string.
- `{{` and `}}` escape literal braces; arguments may be ranges/arrays (flattened automatically). An out-of-range placeholder index → error with `$?=false` (matching PowerShell's throw).

### Null coalescing ?? and ternary ?:
- `$a ?? $b` (7): takes `$b` only when `$a` is `$null`, otherwise `$a`; short-circuiting, the right side evaluated only when the left is empty. Bash's `${a:-$b}` fallback.
- `$cond ? $true : $false` (7): true condition takes the true branch, else the false branch; right-associative (`$a ? $b : $c ? $d : $e` equals `$a ? $b : ($c ? $d : $e)`).
- Both are genuine PowerShell 7 syntax; this program sets no style gate, so they work equally under style 5.X.

### Type literals and casts
- `[int]"42"` converts the value to an integer; decimal or non-numeric strings raise an error (in bash, truncation is usually done with `$((10#$x))`).
- Supported types: `int`, `double`, `string`, `bool`, `datetime`, `hashtable`, `array`, `void`; convert array elements one by one with `int[]`.
- `[datetime]"2020-01-02"` produces a date object; `.Year` and other fields work afterwards.
- `[void](command or expression)` runs it and discards the output, like bash's `> /dev/null`.
- A type name can be stored on its own: `$t = [int]`, then `5 -is $t` tests the type; `"7" -as [double]` yields an empty value instead of an error when conversion fails.
- `[Type]::Member` calls static methods and properties (no bash equivalent): `[math]::Sqrt(4)` gives 2, `[math]::Floor(1.9)` gives 1, `[math]::Round(2.5)` gives 2 (banker's rounding), `[math]::Pow(2, 10)` gives 1024; `[string]::Join(",", 1, 2)` gives `1,2`, `[string]::Format("{0}-{1}", "a", "b")` gives `a-b`; `[datetime]::Now.Year` reads the current year; `[guid]::NewGuid()` generates a random GUID.
- Parameters can declare a type: `function f([int]$x)`, or `param([int]$n)` at the top of a script; positional arguments, named arguments, and default values all convert under the same rules as type literals, and an `[int[]]` annotation wraps a single value into a one-element array; when an argument cannot be converted an error is reported and the function body or script body does not run (no bash equivalent).

### Script blocks and the call operator &
- `{ ... }` is a script block: a statement sequence stored before it runs. Save it as a value (`$sb = { ... }`) or pass it to script-block parameters (such as `Where-Object { ... }`); it displays as `{ ... }`.
- `& target args...` runs the call target: with a script block as the target it executes with function semantics — a leading `param()` declares parameters, positional/named arguments and default values bind under the same rules as functions, extra arguments land in `$args`, and pipeline input enters through `$input`; with a command-name string it calls that command by name (no direct bash equivalent, conceptually close to `eval "cmd args"`). A missing call target after `&` is an error.
- `$sb = { param($x) $x * 2 }; & $sb 21` gives 42; `& { "a"; return "v"; "b" }` gives `a`, `v`.
- `$block.Invoke(args...)` also runs the script block and collects its output into an array: `{ 1; 2 }.Invoke()` gives an array holding 1 and 2.
- A script block sees the caller's variables (dynamic scoping consistent with PowerShell); throw inside the block can be caught by the caller's try/catch.

### Member access and hashtable properties
- `$x.property` fetches members, chaining `$x.a.b` supported alongside method calls `$x.M(...)` (a parenthesis hugging the last segment marks a method).
- Hashtables resolve members "keys before properties": same-named keys win, built-in properties come only after the key misses. Keys are case-insensitive.
- Hashtable built-in properties: `.Count` returns entry count; `.Keys` / `.Values` return insertion-ordered arrays of keys/values; `.Length` is undefined, falling through the scalar fallback to return 1.
- Differences from genuine PowerShell: real `@{...}` guarantees no key order while this program keeps insertion order (like `[ordered]@{}`); real `$h.Keys[0]` returns the whole collection because KeyCollection rejects integer indexing, this program returns the first key.
