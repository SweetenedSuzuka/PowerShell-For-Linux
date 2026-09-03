# PowerShell Command Categories

Commands of the three core modules, categorized by implementation approach:
Category notes:
- **Built-in**: implemented natively in Go (object-pipeline semantics preserved).
- **Mapped Linux**: delegated to the corresponding Linux command/tool.
- **Skipped**: Windows-only / remote sessions / GUI / module management / debugging and the like — platform-bound or out of scope.

> Implementation status: the three core modules hold roughly 245 commands in total; this program builds 115 of them in;
> unimplemented commands that can pass through (mkdir, ping, curl, etc.) run as external commands found on PATH.
> Items marked "skipped" are platform-bound or out of scope. See the batch-by-batch commits in git history.
>
> Detailed documents:
> - [CommandReference.en.md](CommandReference.en.md) — For users: mapping, version, distro, and two examples per command.
> - [CommandDetails.en.md](CommandDetails.en.md) — Development docs: complete parameter-by-parameter, field-by-field descriptions.
> - [UnimplementedCommands.en.md](UnimplementedCommands.en.md) — Every unimplemented command with reasons.
> - [test/核验.ps1](../test/核验.ps1) — Runs each example against the code for verification.

## Microsoft.PowerShell.Core

| cmdlet | Handling | Linux counterpart |
|---|---|---|
| ForEach-Object | Built-in | awk / xargs |
| Where-Object | Built-in | grep / awk |
| Get-Command | Built-in | type / which |
| Get-Help | Built-in | man |
| Get-History | Built-in | history |
| Add-History | Built-in | fc |
| Clear-History | Built-in | history -c |
| Invoke-History | Built-in | fc / !! |
| Clear-Host | Built-in | clear |
| Out-Default (no such command; default output is handled by the pipeline endpoint) | Skipped | — |
| Out-Host | Built-in | — |
| Out-Null | Built-in | /dev/null |
| Set-StrictMode | Skipped | — |
| Set-PSDebug | Skipped | — |
| Get-Module / Import-Module / Remove-Module / New-Module etc. | Skipped | — |
| All PSSession / PSRemoting / Job / ExperimentalFeature | Skipped | — |

## Microsoft.PowerShell.Management

| cmdlet | Handling | Linux counterpart |
|---|---|---|
| Get-ChildItem | Built-in | ls / dir |
| Get-Item | Built-in | stat |
| Get-ItemProperty | Built-in | stat |
| Set-ItemProperty | Built-in | touch |
| Get-Content | Built-in | cat / tail / head |
| Set-Content | Built-in | echo > / tee |
| Add-Content | Built-in | echo >> |
| Clear-Content | Built-in | : > file |
| Set-Item / Clear-Item | Built-in | echo > / : > file |
| Set-Location | Built-in | cd |
| Get-Location | Built-in | pwd |
| Push-Location / Pop-Location | Built-in | pushd / popd |
| Get-PSDrive | Built-in | df / mount |
| New-Item | Built-in | mkdir / touch |
| Remove-Item | Built-in | rm |
| Copy-Item | Built-in | cp |
| Move-Item | Built-in | mv |
| Rename-Item | Built-in | mv / rename |
| Invoke-Item | Built-in | xdg-open |
| Test-Path | Built-in | test -e |
| Resolve-Path | Built-in | realpath |
| Convert-Path | Built-in | realpath |
| Split-Path | Built-in | dirname / basename |
| Join-Path | Built-in | — |
| Get-Process | Built-in | ps |
| Stop-Process | Built-in | kill |
| Start-Process | Built-in | nohup / & |
| Wait-Process | Built-in | wait |
| Debug-Process | Skipped | — |
| Get-ComputerInfo | Built-in | uname -a / lsb_release |
| Get-TimeZone | Built-in | cat /etc/timezone |
| Set-TimeZone | Mapped Linux | timedatectl |
| Get-Culture | Built-in | locale |
| Get-Service | Mapped Linux | systemctl |
| Start/Stop/Restart/Resume/Set-Service | Mapped Linux | systemctl start/stop/restart |
| Test-Connection | Mapped Linux | ping |
| Get-Clipboard / Set-Clipboard | Mapped Linux | xclip / xsel |
| Set-Date | Mapped Linux | date -s |
| Get-HotFix / Get-RecycleBin / registry ItemProperty family (Clear/Copy/Move/New/Remove/Rename-ItemProperty) | Skipped | — |
| Restart-Computer / Stop-Computer / Rename-Computer | Mapped Linux | sudo reboot / shutdown / hostnamectl |
| New/Remove-PSDrive, Get-PSProvider | Skipped | — |
| New/Remove-Service | Skipped | — |

## Microsoft.PowerShell.Utility

| cmdlet | Handling | Linux counterpart |
|---|---|---|
| Write-Output / Write-Host / Write-Error | Built-in | echo |
| Write-Verbose / Write-Warning / Write-Information / Write-Debug | Built-in | echo |
| Format-Table / Format-List / Format-Wide | Built-in | column |
| Format-Hex | Built-in | xxd / od |
| Select-Object | Built-in | cut / awk |
| Select-String | Built-in | grep |
| Sort-Object | Built-in | sort |
| Group-Object | Built-in | uniq -c |
| Measure-Object | Built-in | wc |
| Measure-Command | Built-in | time |
| Get-Member | Built-in | — |
| Get-Unique | Built-in | uniq |
| Compare-Object | Built-in | diff |
| Tee-Object | Built-in | tee |
| Out-String | Built-in | — |
| Out-File | Built-in | > |
| Read-Host | Built-in | read |
| Get-Date | Built-in | date |
| Get-Uptime | Built-in | uptime |
| Get-Random | Built-in | shuf / $RANDOM |
| New-Guid | Built-in | uuidgen |
| Get-Alias / Set-Alias / New-Alias / Remove-Alias / Export-Alias / Import-Alias | Built-in | alias |
| Get-Variable / Set-Variable / New-Variable / Remove-Variable / Clear-Variable | Built-in | — |
| Get-Host | Built-in | — |
| ConvertTo-Csv / ConvertFrom-Csv | Built-in | — |
| ConvertTo-Json / ConvertFrom-Json | Built-in | jq |
| ConvertFrom-StringData | Built-in | — |
| Test-Json | Built-in | jq |
| Join-String | Built-in | paste |
| New-TimeSpan | Built-in | — |
| New-TemporaryFile | Built-in | mktemp |
| Start-Sleep | Built-in | sleep |
| Invoke-Expression | Built-in | eval |
| Invoke-WebRequest / Invoke-RestMethod | Built-in | curl / wget |
| Get-FileHash | Built-in | sha256sum |
| Add-Member | Built-in | — |
| New-Object | Built-in | — |
| Format-Custom | Skipped | — |
| Get-Error / all Event / Breakpoint / Runspace / Trace / TypeData | Skipped | — |
| Out-GridView / Out-Printer / Show-Command / Show-Markdown / Write-Progress | Skipped | GUI/printing |
| CliXml / Html / Xml / Markdown conversion | Skipped | — |
| Send-MailMessage (deprecated) | Skipped | — |
| Add-Type / Get-Verb / Import-LocalizedData | Skipped | — |

## Aliases

The two styles' alias sets are largely identical; style 5 adds `curl`, `wget`, and `sc`.

| Alias | Target | Applicability |
| :--- | :--- | :--- |
| `dir` / `ls` / `gci` | Get-ChildItem | 5.1 and 7 |
| `cd` / `sl` / `chdir` | Set-Location | 5.1 and 7 |
| `pwd` / `gl` | Get-Location | 5.1 and 7 |
| `cat` / `type` / `gc` | Get-Content | 5.1 and 7 |
| `echo` / `write` | Write-Output | 5.1 and 7 |
| `cls` / `clear` | Clear-Host | 5.1 and 7 |
| `ps` / `gps` | Get-Process | 5.1 and 7 |
| `rm` / `del` / `erase` / `ri` / `rd` / `rmdir` | Remove-Item | 5.1 and 7 |
| `cp` / `copy` / `cpi` | Copy-Item | 5.1 and 7 |
| `mv` / `move` / `mi` | Move-Item | 5.1 and 7 |
| `ren` / `rni` | Rename-Item | 5.1 and 7 |
| `ni` | New-Item | 5.1 and 7 |
| `gi` | Get-Item | 5.1 and 7 |
| `ii` | Invoke-Item | 5.1 and 7 |
| `?` / `where` | Where-Object | 5.1 and 7 |
| `%` | ForEach-Object | 5.1 and 7 |
| `sort` | Sort-Object | 5.1 and 7 |
| `select` | Select-Object | 5.1 and 7 |
| `measure` | Measure-Object | 5.1 and 7 |
| `group` | Group-Object | 5.1 and 7 |
| `fl` / `ft` / `fw` | Format-List / Format-Table / Format-Wide | 5.1 and 7 |
| `gcm` | Get-Command | 5.1 and 7 |
| `gal` | Get-Alias | 5.1 and 7 |
| `gv` | Get-Variable | 5.1 and 7 |
| `help` / `man` / `gh` (gh being this program's extension) | Get-Help | 5.1 and 7 |
| `date` | Get-Date | 5.1 and 7 |
| `gdr` | Get-PSDrive | 5.1 and 7 |
| `history` / `ghy` | Get-History | 5.1 and 7 |
| `ihy` | Invoke-History | 5.1 and 7 |
| `sls` | Select-String | 5.1 and 7 |
| `switchstyle`（this program's extension） | Set-PSVersion | 5.1 and 7 |
| `gp` | Get-ItemProperty | 5.1 and 7 |
| `sp` | Set-ItemProperty | 5.1 and 7 |
| `si` | Set-Item | 5.1 and 7 |
| `cli` | Clear-Item | 5.1 and 7 |
| `sv` / `nv` / `rv` / `clv` | Set-Variable / New-Variable / Remove-Variable / Clear-Variable | 5.1 and 7 |
| `sa` / `na` | Set-Alias / New-Alias | 5.1 and 7 |
| `gm` | Get-Member | 5.1 and 7 |
| `gu` | Get-Unique | 5.1 and 7 |
| `gsv` | Get-Service | 5.1 and 7 |
| `sasv` / `spsv` | Start-Service / Stop-Service | 5.1 and 7 |
| `iex` | Invoke-Expression | 5.1 and 7 |
| `iwr` / `irm` | Invoke-WebRequest / Invoke-RestMethod | 5.1 and 7 |
| `sleep` | Start-Sleep | 5.1 and 7 |
| `tee` | Tee-Object | 5.1 and 7 |
| `curl` / `wget` | Invoke-WebRequest | 5.1 only |
| `sc` | Set-Content | 5.1 only |

## Actual differences between 5.X and 7.X

- `$PSVersionTable`: 5.X → PSVersion=5.1, PSEdition=Desktop; 7.X → PSVersion=7, PSEdition=Core, GitCommitId, OS, Platform.
- Automatic variables: 7.X has `$IsLinux/$IsWindows/$IsMacOS`; 5.X doesn't (references yield $null).
- Default alias-set differences (above).
- Syntax and commands are not gated by style: `??` null coalescing, ternary `?:`, `-f` formatting, `$i++` increments, `-in`/`-contains` and other 7.X syntax, plus commands new in 7 such as Test-Json, Join-String, Get-Uptime all work equally under style 5.X in this program.

## Interactive shell behavior

- Typing `powershell` enters the shell process, working directory the Linux root `/`.
- At startup the profile script `~/.config/powershell/profile.ps1` loads by default (skipped when absent or under `-NoProfile`; a broken script reports its error and startup continues).
- The prompt is fixed at the Windows style: `PS C:\>` (`C:\` corresponding to `/`).
- Commands typed map straight onto Linux commands for execution; where privilege escalation is needed, sudo runs them PowerShell-style.
- The default command format is PowerShell 7; `Set-PSVersion 5.1` (alias `switchstyle`) moves to 5.
