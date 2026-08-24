# Unimplemented Commands

The following commands from the three core modules of PowerShell 5.1 / 7 are **not built in**, grouped by reason.
Every command not listed here is either built in (see the [Command Reference](CommandReference.en.md)) or runs via external-command passthrough.
Out-Default is not a standalone command — default output is handled by the pipeline's endpoint, so nothing needs implementing.

## Remote sessions (need remote hosts, out of scope)

- Connect-PSSession, Disconnect-PSSession, Enter-PSSession, Exit-PSSession
- Receive-PSSession, New-PSSession, Remove-PSSession, Get-PSSession
- New-PSSessionOption, New-PSSessionConfigurationFile, Register-PSSessionConfiguration, Set-PSSessionConfiguration
- Get-PSSessionConfiguration, Get-PSSessionCapability, Disable-PSSessionConfiguration
- Enable-PSSessionConfiguration, Test-PSSessionConfigurationFile, Unregister-PSSessionConfiguration
- Export-PSSession, Import-PSSession, Enter-PSHostProcess, Exit-PSHostProcess, Get-PSHostProcessInfo

## Jobs and runspaces (need the job facility, out of scope)

- Start-Job, Stop-Job, Wait-Job, Receive-Job, Get-Job, Remove-Job, Debug-Job
- Wait-Debugger, Debug-Runspace, Get-Runspace, Get-RunspaceDebug, Enable-RunspaceDebug, Disable-RunspaceDebug

## Modules and assemblies (don't fit a single-file interpreter)

- Import-Module, Export-ModuleMember, Get-Module, Remove-Module, New-Module, New-ModuleManifest
- Test-ModuleManifest, Import-PowerShellDataFile, Import-LocalizedData, Add-Type, Get-Verb

## Windows-only (registry / services / recycle bin / patches / events)

- Clear-ItemProperty, Copy-ItemProperty, Move-ItemProperty, New-ItemProperty, Remove-ItemProperty, Rename-ItemProperty
- New-Service, Remove-Service, Suspend-Service
- Get-RecycleBin, Clear-RecycleBin, Get-HotFix, Get-Counter
- Get-WmiObject, Invoke-WmiMethod and the rest of the CIM/WMI family (including Set-CimInstance)
- Get-EventLog, Clear-EventLog, New-EventLog, Remove-EventLog, Write-EventLog

## GUI and printing

- Out-GridView, Show-Command, Show-Markdown, Out-Printer, Write-Progress (no progress bar in terminals)

## Events / breakpoints / tracing (debugging facilities)

- Register-ObjectEvent, Register-EngineEvent, Unregister-Event, Get-Event, Get-EventSubscriber
- New-Event, Remove-Event, Wait-Event, Set-PSBreakpoint, Get-PSBreakpoint, Remove-PSBreakpoint
- Disable-PSBreakpoint, Enable-PSBreakpoint, Get-PSCallStack, Trace-Command, Get-TraceSource, Set-TraceSource, Get-Error

## Serialization / markers / formats (rarely used)

- ConvertTo-Xml, ConvertFrom-CliXml, ConvertTo-CliXml, Export-Clixml, Import-Clixml
- ConvertTo-Html, ConvertFrom-Markdown, ConvertFrom-SddlString, Show-Markdown
- Get-MarkdownOption, Set-MarkdownOption, Format-Custom
- Get-FormatData, Export-FormatData, Update-FormatData
- Update-TypeData, Get-TypeData, Remove-TypeData, Update-List, Select-Xml

## Platform / miscellany

- Get-SecureRandom (7; use Get-Random instead), Unblock-File (Windows marker)
- Send-MailMessage (deprecated), Set-PSDebug, Set-StrictMode, Save-Help, Update-Help
- Register-ArgumentCompleter, Disable-ExperimentalFeature, Enable-ExperimentalFeature
- Get-ExperimentalFeature, Get-PSSubsystem, Switch-Process (new in 7.x)

> Note: if typed in anyway, these commands are looked up on PATH as **external commands**.
> When not found, "The term 'X' is not recognized as the name of a cmdlet, function, script file, or runnable program." appears, asking the user to run it as a single command.
> Anything demanding complex combination or nesting is rejected outright with a message to that effect.
