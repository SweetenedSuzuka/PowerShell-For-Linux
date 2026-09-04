# Original Cross-Platform Command Details

Commands in the original PowerShell that work beyond Windows.

Status legend:
- **Go implementation** — behavior is reproduced in Go inside this program, with no external commands called.
- **Mapped Linux** — native Linux commands/tools are called; the tool follows in parentheses.
- **Not implemented** — not implemented in this program.
- **Out of scope** — Windows-only; this program does not do it, see notes for reasons.

## Command List

| Command | Module | Version | Differences | Purpose | Status | Notes |
|---|---|---|---|---|---|---|
| [`Add-Content`](#add-content) | Microsoft.PowerShell.Management | Both | Syntax differs | Adds content to the specified items, such as adding words to a file. | Go implementation | The `ac` alias is not implemented; use the full name. |
| [`Add-History`](#add-history) | Microsoft.PowerShell.Core | Both | None | Appends entries to the session history. | Go implementation |  |
| [`Add-Member`](#add-member) | Microsoft.PowerShell.Utility | Both | Syntax differs | Adds custom properties and methods to an instance of a PowerShell object. | Go implementation | Only the NoteProperty type is supported. |
| [`Add-Type`](#add-type) | Microsoft.PowerShell.Utility | Both | Syntax differs | Adds a Microsoft .NET class to a PowerShell session. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Clear-Content`](#clear-content) | Microsoft.PowerShell.Management | Both | Syntax differs | Deletes the contents of an item, but does not delete the item. | Go implementation | Nonexistent paths silently create an empty file (PowerShell errors with path-not-found). |
| [`Clear-History`](#clear-history) | Microsoft.PowerShell.Core | Both | None | Deletes entries from the PowerShell session command history. | Go implementation |  |
| [`Clear-Item`](#clear-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Clears the contents of an item, but does not delete the item. | Go implementation |  |
| [`Clear-ItemProperty`](#clear-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Clears the value of a property but does not delete the property. | Not implemented | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Clear-Variable`](#clear-variable) | Microsoft.PowerShell.Utility | Both | None | Deletes the value of a variable. | Go implementation |  |
| [`Compare-Object`](#compare-object) | Microsoft.PowerShell.Utility | Both | None | Compares two sets of objects. | Go implementation | Behaves identically. |
| [`Compress-PSResource`](#compress-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Compresses a specified folder containing module or script resources into a .nupkg file. | Not implemented |  |
| [`Convert-Path`](#convert-path) | Microsoft.PowerShell.Management | Both | Syntax differs | Converts a path from a PowerShell path to a PowerShell provider path. | Go implementation |  |
| [`ConvertFrom-CliXml`](#convertfrom-clixml) | Microsoft.PowerShell.Utility | 7 only | 7 only | Converts a CliXml-formatted string to a custom **PSObject**. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`ConvertFrom-Csv`](#convertfrom-csv) | Microsoft.PowerShell.Utility | Both | None | Converts object properties in character-separated value (CSV) format into CSV versions of the original objects. | Go implementation |  |
| [`ConvertFrom-Json`](#convertfrom-json) | Microsoft.PowerShell.Utility | Both | Description differs; Syntax differs | 5.1: Converts a JSON-formatted string to a custom object. / 7: Converts a JSON-formatted string to a custom object or a hash table. | Go implementation |  |
| [`ConvertFrom-Markdown`](#convertfrom-markdown) | Microsoft.PowerShell.Utility | 7 only | 7 only | Convert the contents of a string or a file to a **MarkdownInfo** object. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`ConvertFrom-SecureString`](#convertfrom-securestring) | Microsoft.PowerShell.Security | Both | Syntax differs | Converts a secure string to an encrypted standard string. | Not implemented |  |
| [`ConvertFrom-StringData`](#convertfrom-stringdata) | Microsoft.PowerShell.Utility | Both | Syntax differs | Converts a string containing one or more key and value pairs to a hash table. | Go implementation |  |
| [`ConvertTo-CliXml`](#convertto-clixml) | Microsoft.PowerShell.Utility | 7 only | 7 only | Converts an object to a CliXml-formatted string. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`ConvertTo-Csv`](#convertto-csv) | Microsoft.PowerShell.Utility | Both | Syntax differs | Converts .NET objects into a series of character-separated value (CSV) strings. | Go implementation |  |
| [`ConvertTo-Html`](#convertto-html) | Microsoft.PowerShell.Utility | Both | Syntax differs | Converts .NET objects into HTML that can be displayed in a Web browser. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`ConvertTo-Json`](#convertto-json) | Microsoft.PowerShell.Utility | Both | Syntax differs | Converts an object to a JSON-formatted string. | Go implementation | Output uses a fixed 2-space indent. |
| [`ConvertTo-SecureString`](#convertto-securestring) | Microsoft.PowerShell.Security | Both | None | Converts plain text or encrypted strings to secure strings. | Not implemented |  |
| [`ConvertTo-Xml`](#convertto-xml) | Microsoft.PowerShell.Utility | Both | None | Creates an XML-based representation of an object. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Copy-Item`](#copy-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Copies an item from one location to another. | Go implementation |  |
| [`Copy-ItemProperty`](#copy-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Copies a property and value from a specified location to another location. | Not implemented | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Debug-Job`](#debug-job) | Microsoft.PowerShell.Core | Both | Description differs; Syntax differs | 5.1: Debugs a running background, remote, or Windows PowerShell Workflow job. / 7: Debugs a running background or remote job. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Debug-Process`](#debug-process) | Microsoft.PowerShell.Management | Both | None | Debugs one or more processes running on the local computer. | Not implemented |  |
| [`Debug-Runspace`](#debug-runspace) | Microsoft.PowerShell.Utility | Both | Syntax differs | Starts an interactive debugging session with a runspace. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Disable-ExperimentalFeature`](#disable-experimentalfeature) | Microsoft.PowerShell.Core | 7 only | 7 only | Disable an experimental feature on startup of new instance of PowerShell. | Not implemented | Platform / miscellaneous |
| [`Disable-PSBreakpoint`](#disable-psbreakpoint) | Microsoft.PowerShell.Utility | Both | Syntax differs | Disables the breakpoints in the current console. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Disable-RunspaceDebug`](#disable-runspacedebug) | Microsoft.PowerShell.Utility | Both | None | Disables debugging on one or more runspaces, and releases any pending debugger stop. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Enable-ExperimentalFeature`](#enable-experimentalfeature) | Microsoft.PowerShell.Core | 7 only | 7 only | Enable an experimental feature on startup of new instance of PowerShell. | Not implemented | Platform / miscellaneous |
| [`Enable-PSBreakpoint`](#enable-psbreakpoint) | Microsoft.PowerShell.Utility | Both | Syntax differs | Enables the breakpoints in the current console. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Enable-RunspaceDebug`](#enable-runspacedebug) | Microsoft.PowerShell.Utility | Both | None | Enables debugging on runspaces where any breakpoint is preserved until a debugger is attached. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Enter-PSHostProcess`](#enter-pshostprocess) | Microsoft.PowerShell.Core | Both | Syntax differs | Connects to and enters into an interactive session with a local process. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Enter-PSSession`](#enter-pssession) | Microsoft.PowerShell.Core | Both | Syntax differs | Starts an interactive session with a remote computer. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Exit-PSHostProcess`](#exit-pshostprocess) | Microsoft.PowerShell.Core | Both | None | Closes an interactive session with a local process. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Exit-PSSession`](#exit-pssession) | Microsoft.PowerShell.Core | Both | None | Ends an interactive session with a remote computer. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Export-Alias`](#export-alias) | Microsoft.PowerShell.Utility | Both | None | Exports information about currently defined aliases to a file. | Go implementation |  |
| [`Export-Clixml`](#export-clixml) | Microsoft.PowerShell.Utility | Both | Syntax differs | Creates an XML-based representation of an object or objects and stores it in a file. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Export-Csv`](#export-csv) | Microsoft.PowerShell.Utility | Both | Syntax differs | Converts objects into a series of character-separated value (CSV) strings and saves the strings to a file. | Not implemented |  |
| [`Export-FormatData`](#export-formatdata) | Microsoft.PowerShell.Utility | Both | None | Saves formatting data from the current session in a formatting file. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Export-ModuleMember`](#export-modulemember) | Microsoft.PowerShell.Core | Both | None | Specifies the module members that are exported. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Export-PSSession`](#export-pssession) | Microsoft.PowerShell.Utility | Both | Syntax differs | Exports commands from another session and saves them in a PowerShell module. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Find-Package`](#find-package) | PackageManagement | Both | Syntax differs | Finds software packages in available package sources. | Not implemented |  |
| [`Find-PackageProvider`](#find-packageprovider) | PackageManagement | Both | None | Returns a list of Package Management package providers available for installation. | Not implemented |  |
| [`Find-PSResource`](#find-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Searches for packages from a repository (local or remote), based on a name or other package properties. | Not implemented |  |
| [`ForEach-Object`](#foreach-object) | Microsoft.PowerShell.Core | Both | Syntax differs | Performs an operation against each item in a collection of input objects. | Go implementation |  |
| [`Format-Custom`](#format-custom) | Microsoft.PowerShell.Utility | Both | None | Uses a customized view to format the output. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Format-Hex`](#format-hex) | Microsoft.PowerShell.Utility | Both | Syntax differs | Displays a file or other input as hexadecimal. | Go implementation |  |
| [`Format-List`](#format-list) | Microsoft.PowerShell.Utility | Both | None | Formats the output as a list of properties in which each property appears on a new line. | Go implementation |  |
| [`Format-Table`](#format-table) | Microsoft.PowerShell.Utility | Both | None | Formats the output as a table. | Go implementation | Column-width fitting differs; no ANSI colors or directory headers; the time column uses a fixed format (PowerShell follows locale). |
| [`Format-Wide`](#format-wide) | Microsoft.PowerShell.Utility | Both | None | Formats objects as a wide table that displays only one property of each object. | Go implementation |  |
| [`Get-Alias`](#get-alias) | Microsoft.PowerShell.Utility | Both | None | Gets the aliases for the current session. | Go implementation | Looking up a nonexistent alias returns empty with `$?` staying True (PowerShell errors with alias-not-found). |
| [`Get-ChildItem`](#get-childitem) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets the items and child items in one or more specified locations. | Go implementation | `-Force` has no hidden-file effect. |
| [`Get-Clipboard`](#get-clipboard) | Microsoft.PowerShell.Management | Both | Description differs; Syntax differs | 5.1: Gets the current Windows clipboard entry. / 7: Gets the contents of the clipboard. | Mapped Linux (xclip / xsel) |  |
| [`Get-CmsMessage`](#get-cmsmessage) | Microsoft.PowerShell.Security | Both | None | Gets content that has been encrypted by using the Cryptographic Message Syntax format. | Not implemented |  |
| [`Get-Command`](#get-command) | Microsoft.PowerShell.Core | Both | Syntax differs | Gets all commands. | Go implementation | External commands are not listed, only built-ins and aliases. |
| [`Get-Content`](#get-content) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets the content of the item at the specified location. | Go implementation |  |
| [`Get-Credential`](#get-credential) | Microsoft.PowerShell.Security | Both | Syntax differs | Gets a credential object based on a user name and password. | Not implemented |  |
| [`Get-Culture`](#get-culture) | Microsoft.PowerShell.Utility | Both | Syntax differs | Gets the current culture set in the operating system. | Go implementation | Culture follows the UI language, falling back to zh-CN when the UI language has no registered culture. |
| [`Get-Date`](#get-date) | Microsoft.PowerShell.Utility | Both | Syntax differs | Gets the current date and time. | Go implementation |  |
| [`Get-Error`](#get-error) | Microsoft.PowerShell.Utility | 7 only | 7 only | Gets and displays the most recent error messages from the current session. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Get-Event`](#get-event) | Microsoft.PowerShell.Utility | Both | None | Gets the events in the event queue. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Get-EventSubscriber`](#get-eventsubscriber) | Microsoft.PowerShell.Utility | Both | None | Gets the event subscribers in the current session. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Get-ExecutionPolicy`](#get-executionpolicy) | Microsoft.PowerShell.Security | Both | None | Gets the execution policies for the current session. | Not implemented |  |
| [`Get-ExperimentalFeature`](#get-experimentalfeature) | Microsoft.PowerShell.Core | 7 only | 7 only | Gets experimental features. | Not implemented | Platform / miscellaneous |
| [`Get-FileHash`](#get-filehash) | Microsoft.PowerShell.Utility | Both | Syntax differs | Computes the hash value for a file by using a specified hash algorithm. | Go implementation | SHA384, MACTripleDES and RIPEMD160 are not supported. |
| [`Get-FormatData`](#get-formatdata) | Microsoft.PowerShell.Utility | Both | None | Gets the formatting data in the current session. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Get-Help`](#get-help) | Microsoft.PowerShell.Core | Both | Syntax differs | Displays information about PowerShell commands and concepts. | Go implementation | Shows only names, syntax and aliases without detailed descriptions. |
| [`Get-History`](#get-history) | Microsoft.PowerShell.Core | Both | None | Gets a list of the commands entered during the current session. | Go implementation |  |
| [`Get-Host`](#get-host) | Microsoft.PowerShell.Utility | Both | None | Gets an object that represents the current host program. | Go implementation |  |
| [`Get-InstalledPSResource`](#get-installedpsresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Returns modules and scripts installed on the machine via PowerShellGet. | Not implemented |  |
| [`Get-Item`](#get-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets the item at the specified location. | Go implementation |  |
| [`Get-ItemProperty`](#get-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets the properties of a specified item. | Go implementation | Only 5 fields are output. |
| [`Get-ItemPropertyValue`](#get-itempropertyvalue) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets the value for one or more properties of a specified item. | Not implemented |  |
| [`Get-Job`](#get-job) | Microsoft.PowerShell.Core | Both | None | Gets PowerShell background jobs that are running in the current session. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Get-Location`](#get-location) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets information about the current working location or a location stack. | Go implementation | Outputs path strings directly (PowerShell renders a Path table of PathInfo objects). |
| [`Get-MarkdownOption`](#get-markdownoption) | Microsoft.PowerShell.Utility | 7 only | 7 only | Returns the current colors and styles used for rendering Markdown content in the console. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Get-Member`](#get-member) | Microsoft.PowerShell.Utility | Both | None | Gets the properties and methods of objects. | Go implementation | Method members are not listed, only types and properties. |
| [`Get-Module`](#get-module) | Microsoft.PowerShell.Core | Both | Syntax differs | List the modules imported in the current session or that can be imported from the PSModulePath. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Get-Package`](#get-package) | PackageManagement | Both | Syntax differs | Returns a list of all software packages that were installed with PackageManagement. | Not implemented |  |
| [`Get-PackageProvider`](#get-packageprovider) | PackageManagement | Both | None | Returns a list of package providers that are connected to Package Management. | Not implemented |  |
| [`Get-PackageSource`](#get-packagesource) | PackageManagement | Both | Syntax differs | Gets a list of package sources that are registered for a package provider. | Not implemented |  |
| [`Get-PfxCertificate`](#get-pfxcertificate) | Microsoft.PowerShell.Security | Both | Syntax differs | Gets information about PFX certificate files on the computer. | Not implemented |  |
| [`Get-Process`](#get-process) | Microsoft.PowerShell.Management | Both | Description differs; Syntax differs | 5.1: Gets the processes that are running on the local computer or a remote computer. / 7: Gets the processes that are running on the local computer. | Go implementation | The memory field is named Memory (WS in PowerShell) holding physical memory bytes, matching the semantics of PowerShell's WS; `-Name` matches by substring (PowerShell uses exact names or wildcards). |
| [`Get-PSBreakpoint`](#get-psbreakpoint) | Microsoft.PowerShell.Utility | Both | Syntax differs | Gets the breakpoints that are set in the current session. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Get-PSCallStack`](#get-pscallstack) | Microsoft.PowerShell.Utility | Both | None | Displays the current call stack. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Get-PSDrive`](#get-psdrive) | Microsoft.PowerShell.Management | Both | Syntax differs | Gets drives in the current session. | Go implementation | Only the root drive and Env exist. |
| [`Get-PSHostProcessInfo`](#get-pshostprocessinfo) | Microsoft.PowerShell.Core | Both | None | Gets process information about the PowerShell host. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Get-PSProvider`](#get-psprovider) | Microsoft.PowerShell.Management | Both | None | Gets information about the specified PowerShell provider. | Not implemented |  |
| [`Get-PSReadLineKeyHandler`](#get-psreadlinekeyhandler) | PSReadLine | Both | Syntax differs | Gets the key bindings for the PSReadLine module. | Not implemented |  |
| [`Get-PSReadLineOption`](#get-psreadlineoption) | PSReadLine | Both | None | Gets values for the options that can be configured. | Not implemented |  |
| [`Get-PSResourceRepository`](#get-psresourcerepository) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Finds and returns registered repository information. | Not implemented |  |
| [`Get-PSScriptFileInfo`](#get-psscriptfileinfo) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Returns the metadata for a script. | Not implemented |  |
| [`Get-PSSession`](#get-pssession) | Microsoft.PowerShell.Core | Both | Syntax differs | Gets the PowerShell sessions on local and remote computers. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Get-PSSubsystem`](#get-pssubsystem) | Microsoft.PowerShell.Core | 7 only | 7 only | Retrieves information about the subsystems registered in PowerShell. | Not implemented | Platform / miscellaneous |
| [`Get-Random`](#get-random) | Microsoft.PowerShell.Utility | Both | Syntax differs | Gets a random number, or selects objects randomly from a collection. | Go implementation |  |
| [`Get-Runspace`](#get-runspace) | Microsoft.PowerShell.Utility | Both | None | Gets active runspaces within a PowerShell host process. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Get-RunspaceDebug`](#get-runspacedebug) | Microsoft.PowerShell.Utility | Both | None | Shows runspace debugging options. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Get-SecureRandom`](#get-securerandom) | Microsoft.PowerShell.Utility | 7 only | 7 only | Gets a random number, or selects objects randomly from a collection. | Not implemented | Platform / miscellaneous |
| [`Get-TimeZone`](#get-timezone) | Microsoft.PowerShell.Management | Both | None | Gets the current time zone or a list of available time zones. | Go implementation | Reads /etc/timezone. |
| [`Get-TraceSource`](#get-tracesource) | Microsoft.PowerShell.Utility | Both | None | Gets PowerShell components that are instrumented for tracing. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Get-TypeData`](#get-typedata) | Microsoft.PowerShell.Utility | Both | None | Gets the extended type data in the current session. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Get-UICulture`](#get-uiculture) | Microsoft.PowerShell.Utility | Both | None | Gets the current UI culture settings in the operating system. | Not implemented |  |
| [`Get-Unique`](#get-unique) | Microsoft.PowerShell.Utility | Both | Syntax differs | Returns unique items from a sorted list. | Go implementation | Deduplicates by object string; inputs need not be adjacent. |
| [`Get-Uptime`](#get-uptime) | Microsoft.PowerShell.Utility | 7 only | 7 only | Get the **TimeSpan** since last boot. | Go implementation | Always 0 on other platforms. |
| [`Get-Variable`](#get-variable) | Microsoft.PowerShell.Utility | Both | None | Gets the variables in the current console. | Go implementation |  |
| [`Get-Verb`](#get-verb) | 5.1 in Microsoft.PowerShell.Core<br>7 in Microsoft.PowerShell.Utility | Both | None | Gets approved PowerShell verbs. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Group-Object`](#group-object) | Microsoft.PowerShell.Utility | Both | None | Groups objects that contain the same value for specified properties. | Go implementation | Behaves identically. |
| [`Import-Alias`](#import-alias) | Microsoft.PowerShell.Utility | Both | None | Imports an alias list from a file. | Go implementation |  |
| [`Import-Clixml`](#import-clixml) | Microsoft.PowerShell.Utility | Both | Syntax differs | Imports a CLIXML file and creates corresponding objects in PowerShell. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Import-Csv`](#import-csv) | Microsoft.PowerShell.Utility | Both | Syntax differs | Creates table-like custom objects from the items in a character-separated value (CSV) file. | Not implemented |  |
| [`Import-LocalizedData`](#import-localizeddata) | Microsoft.PowerShell.Utility | Both | None | Imports language-specific data into scripts and functions based on the UI culture that's selected for the operating system. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Import-Module`](#import-module) | Microsoft.PowerShell.Core | Both | Syntax differs | Adds modules to the current session. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Import-PackageProvider`](#import-packageprovider) | PackageManagement | Both | None | Adds Package Management package providers to the current session. | Not implemented |  |
| [`Import-PowerShellDataFile`](#import-powershelldatafile) | Microsoft.PowerShell.Utility | Both | None | Imports values from a `.psd1` file without invoking its contents. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Import-PSSession`](#import-pssession) | Microsoft.PowerShell.Utility | Both | None | Imports commands from another session into the current session. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Install-Package`](#install-package) | PackageManagement | Both | Syntax differs | Installs one or more software packages. | Not implemented |  |
| [`Install-PackageProvider`](#install-packageprovider) | PackageManagement | Both | None | Installs one or more Package Management package providers. | Not implemented |  |
| [`Install-PSResource`](#install-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Installs resources from a registered repository. | Not implemented |  |
| [`Invoke-Command`](#invoke-command) | Microsoft.PowerShell.Core | Both | Syntax differs | Runs commands on local and remote computers. | Not implemented |  |
| [`Invoke-Expression`](#invoke-expression) | Microsoft.PowerShell.Utility | Both | None | Runs commands or expressions on the local computer. | Go implementation |  |
| [`Invoke-History`](#invoke-history) | Microsoft.PowerShell.Core | Both | None | Runs commands from the session history. | Go implementation |  |
| [`Invoke-Item`](#invoke-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Performs the default action on the specified item. | Go implementation | Currently only prints the path without opening it. |
| [`Invoke-RestMethod`](#invoke-restmethod) | Microsoft.PowerShell.Utility | Both | Syntax differs | Sends an HTTP or HTTPS request to a RESTful web service. | Go implementation | Non-JSON responses return plain text. |
| [`Invoke-WebRequest`](#invoke-webrequest) | Microsoft.PowerShell.Utility | Both | Syntax differs | Gets content from a web page on the internet. | Go implementation | Returns plain text (split into strings by line) with no properties like StatusCode. |
| [`Join-Path`](#join-path) | Microsoft.PowerShell.Management | Both | Syntax differs | Combines a path and a child path into a single path. | Go implementation |  |
| [`Join-String`](#join-string) | Microsoft.PowerShell.Utility | 7 only | 7 only | Combines objects from the pipeline into a single string. | Go implementation | Without `-Separator`, items join with a space by default (PowerShell joins with no separator). |
| [`Measure-Command`](#measure-command) | Microsoft.PowerShell.Utility | Both | None | Measures the time it takes to run scriptblocks and cmdlets. | Go implementation |  |
| [`Measure-Object`](#measure-object) | Microsoft.PowerShell.Utility | Both | Syntax differs | Calculates the numeric properties of objects, and the characters, words, and lines in string objects, such as files of text. | Go implementation | Min/Max over mixed numeric and non-numeric input only counts numbers (PowerShell compares non-numbers as strings); non-numeric input yields empty instead of erroring (PowerShell reports a non-terminating error). |
| [`Move-Item`](#move-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Moves an item from one location to another. | Go implementation | Moving directories requires `-Recurse`. |
| [`Move-ItemProperty`](#move-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Moves a property from one location to another. | Not implemented | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`New-Alias`](#new-alias) | Microsoft.PowerShell.Utility | Both | None | Creates a new alias. | Go implementation |  |
| [`New-Event`](#new-event) | Microsoft.PowerShell.Utility | Both | None | Creates a new event. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`New-Guid`](#new-guid) | Microsoft.PowerShell.Utility | Both | Syntax differs | Creates a GUID. | Go implementation |  |
| [`New-Item`](#new-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Creates a new item. | Go implementation |  |
| [`New-ItemProperty`](#new-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Creates a new property for an item and sets its value. | Not implemented | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`New-Module`](#new-module) | Microsoft.PowerShell.Core | Both | None | Creates a new dynamic module that exists only in memory. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`New-ModuleManifest`](#new-modulemanifest) | Microsoft.PowerShell.Core | Both | Syntax differs | Creates a new module manifest. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`New-Object`](#new-object) | Microsoft.PowerShell.Utility | Both | None | Creates an instance of a Microsoft .NET Framework or COM object. | Go implementation | Only `PSObject` / `PSCustomObject` are supported; other types (e.g. System.Collections.ArrayList) report "not supported". |
| [`New-PSDrive`](#new-psdrive) | Microsoft.PowerShell.Management | Both | Syntax differs | Creates temporary and persistent drives that are associated with a location in an item data store. | Not implemented |  |
| [`New-PSRoleCapabilityFile`](#new-psrolecapabilityfile) | Microsoft.PowerShell.Core | Both | None | Creates a file that defines a set of capabilities to be exposed through a session configuration. | Not implemented |  |
| [`New-PSScriptFileInfo`](#new-psscriptfileinfo) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | The cmdlet creates a new script file, including metadata about the script. | Not implemented |  |
| [`New-PSSession`](#new-pssession) | Microsoft.PowerShell.Core | Both | Syntax differs | Creates a persistent connection to a local or remote computer. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`New-PSSessionConfigurationFile`](#new-pssessionconfigurationfile) | Microsoft.PowerShell.Core | Both | None | Creates a file that defines a session configuration. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`New-PSSessionOption`](#new-pssessionoption) | Microsoft.PowerShell.Core | Both | None | Creates an object that contains advanced options for a PSSession. | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`New-PSTransportOption`](#new-pstransportoption) | Microsoft.PowerShell.Core | Both | None | Creates an object that contains advanced options for a session configuration. | Not implemented |  |
| [`New-TemporaryFile`](#new-temporaryfile) | Microsoft.PowerShell.Utility | Both | None | Creates a temporary file. | Go implementation |  |
| [`New-TimeSpan`](#new-timespan) | Microsoft.PowerShell.Utility | Both | Syntax differs | Creates a TimeSpan object. | Go implementation |  |
| [`New-Variable`](#new-variable) | Microsoft.PowerShell.Utility | Both | None | Creates a new variable. | Go implementation |  |
| [`Out-Default`](#out-default) | Microsoft.PowerShell.Core | Both | None | Sends the output to the default formatter and to the default output cmdlet. | Not implemented |  |
| [`Out-File`](#out-file) | Microsoft.PowerShell.Utility | Both | Syntax differs | Sends output to a file. | Go implementation |  |
| [`Out-Host`](#out-host) | Microsoft.PowerShell.Core | Both | None | Sends output to the command line. | Go implementation |  |
| [`Out-Null`](#out-null) | Microsoft.PowerShell.Core | Both | None | Hides the output instead of sending it down the pipeline or displaying it. | Go implementation |  |
| [`Out-String`](#out-string) | Microsoft.PowerShell.Utility | Both | Syntax differs | Outputs input objects as a string. | Go implementation |  |
| [`Pop-Location`](#pop-location) | Microsoft.PowerShell.Management | Both | Syntax differs | Changes the current location to the location most recently pushed onto the stack. | Go implementation | The `pushd` / `popd` aliases are not implemented; use the full names. |
| [`Protect-CmsMessage`](#protect-cmsmessage) | Microsoft.PowerShell.Security | Both | None | Encrypts content by using the Cryptographic Message Syntax format. | Not implemented |  |
| [`Publish-PSResource`](#publish-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Publishes a specified module from the local computer to PSResource repository. | Not implemented |  |
| [`Push-Location`](#push-location) | Microsoft.PowerShell.Management | Both | Syntax differs | Adds the current location to the top of a location stack. | Go implementation | The `pushd` / `popd` aliases are not implemented; use the full names. |
| [`Read-Host`](#read-host) | Microsoft.PowerShell.Utility | Both | Syntax differs | Reads a line of input from the console. | Go implementation |  |
| [`Receive-Job`](#receive-job) | Microsoft.PowerShell.Core | Both | None | Gets the results of the PowerShell background jobs in the current session. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Register-ArgumentCompleter`](#register-argumentcompleter) | Microsoft.PowerShell.Core | Both | Syntax differs | Registers a custom argument completer. | Not implemented | Platform / miscellaneous |
| [`Register-EngineEvent`](#register-engineevent) | Microsoft.PowerShell.Utility | Both | None | Subscribes to events that are generated by the PowerShell engine and by the `New-Event` cmdlet. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Register-ObjectEvent`](#register-objectevent) | Microsoft.PowerShell.Utility | Both | None | Subscribes to the events that are generated by a Microsoft .NET Framework object. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Register-PackageSource`](#register-packagesource) | PackageManagement | Both | Syntax differs | Adds a package source for a specified package provider. | Not implemented |  |
| [`Register-PSResourceRepository`](#register-psresourcerepository) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Registers a repository for PowerShell resources. | Not implemented |  |
| [`Remove-Alias`](#remove-alias) | Microsoft.PowerShell.Utility | 7 only | 7 only | Remove an alias from the current session. | Go implementation |  |
| [`Remove-Event`](#remove-event) | Microsoft.PowerShell.Utility | Both | None | Deletes events from the event queue. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Remove-Item`](#remove-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Deletes the specified items. | Go implementation | Nonexistent paths are silently ignored (PowerShell errors with path-not-found). |
| [`Remove-ItemProperty`](#remove-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Deletes the property and its value from an item. | Not implemented | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Remove-Job`](#remove-job) | Microsoft.PowerShell.Core | Both | Syntax differs | Deletes a PowerShell background job. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Remove-Module`](#remove-module) | Microsoft.PowerShell.Core | Both | None | Removes modules from the current session. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Remove-PSBreakpoint`](#remove-psbreakpoint) | Microsoft.PowerShell.Utility | Both | Syntax differs | Deletes breakpoints from the current console. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Remove-PSDrive`](#remove-psdrive) | Microsoft.PowerShell.Management | Both | Syntax differs | Deletes temporary PowerShell drives and disconnects mapped network drives. | Not implemented |  |
| [`Remove-PSReadLineKeyHandler`](#remove-psreadlinekeyhandler) | PSReadLine | Both | None | Removes a key binding. | Not implemented |  |
| [`Remove-PSSession`](#remove-pssession) | Microsoft.PowerShell.Core | Both | None | Closes one or more PowerShell sessions (PSSessions). | Not implemented | Remote sessions (requires a remote host; out of scope) |
| [`Remove-TypeData`](#remove-typedata) | Microsoft.PowerShell.Utility | Both | None | Deletes extended types from the current session. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Remove-Variable`](#remove-variable) | Microsoft.PowerShell.Utility | Both | None | Deletes a variable and its value. | Go implementation |  |
| [`Rename-Item`](#rename-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Renames an item in a PowerShell provider namespace. | Go implementation |  |
| [`Rename-ItemProperty`](#rename-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Renames a property of an item. | Not implemented | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Reset-PSResourceRepository`](#reset-psresourcerepository) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Creates a new default PSRepositories.xml file with preregistered repositories. | Not implemented |  |
| [`Resolve-Path`](#resolve-path) | Microsoft.PowerShell.Management | Both | Syntax differs | Resolves the wildcard characters in a path, and displays the path contents. | Go implementation | Resolve-Path additionally resolves symbolic links. |
| [`Restart-Computer`](#restart-computer) | Microsoft.PowerShell.Management | Both | Syntax differs | Restarts the operating system on local and remote computers. | Mapped Linux (sudo reboot / shutdown / hostnamectl) |  |
| [`Save-Help`](#save-help) | Microsoft.PowerShell.Core | Both | Syntax differs | Downloads and saves the newest help files to a file system directory. | Not implemented | Platform / miscellaneous |
| [`Save-Package`](#save-package) | PackageManagement | Both | Syntax differs | Saves packages to the local computer without installing them. | Not implemented |  |
| [`Save-PSResource`](#save-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Saves resources (modules and scripts) from a registered repository onto the machine. | Not implemented |  |
| [`Select-Object`](#select-object) | Microsoft.PowerShell.Utility | Both | Syntax differs | Selects objects or object properties. | Go implementation | Behaves identically. |
| [`Select-String`](#select-string) | Microsoft.PowerShell.Utility | Both | Syntax differs | Finds text in strings and files. | Go implementation | Invalid regex does not error (it simply does not match). |
| [`Select-Xml`](#select-xml) | Microsoft.PowerShell.Utility | Both | None | Finds text in an XML string or document. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Send-MailMessage`](#send-mailmessage) | Microsoft.PowerShell.Utility | Both | Syntax differs | Sends an email message. | Not implemented | Platform / miscellaneous |
| [`Set-Alias`](#set-alias) | Microsoft.PowerShell.Utility | Both | None | Creates or changes an alias for a cmdlet or other command in the current PowerShell session. | Go implementation |  |
| [`Set-Clipboard`](#set-clipboard) | Microsoft.PowerShell.Management | Both | Description differs; Syntax differs | 5.1: Sets the current Windows clipboard entry. / 7: Sets the contents of the clipboard. | Mapped Linux (xclip / xsel) |  |
| [`Set-Content`](#set-content) | Microsoft.PowerShell.Management | Both | Syntax differs | Writes new content or replaces existing content in a file. | Go implementation |  |
| [`Set-Date`](#set-date) | Microsoft.PowerShell.Utility | Both | None | Changes the system time on the computer to a time that you specify. | Mapped Linux (date -s) |  |
| [`Set-ExecutionPolicy`](#set-executionpolicy) | Microsoft.PowerShell.Security | Both | None | Sets the PowerShell execution policies for Windows computers. | Not implemented |  |
| [`Set-Item`](#set-item) | Microsoft.PowerShell.Management | Both | Syntax differs | Changes the value of an item to the value specified in the command. | Go implementation | Set-Item accepts the `env:` prefix to set environment variables. |
| [`Set-ItemProperty`](#set-itemproperty) | Microsoft.PowerShell.Management | Both | Syntax differs | Creates or changes the value of a property of an item. | Go implementation | Only modification time can be changed; other properties are ignored. |
| [`Set-Location`](#set-location) | Microsoft.PowerShell.Management | Both | Syntax differs | Sets the current working location to a specified location. | Go implementation |  |
| [`Set-MarkdownOption`](#set-markdownoption) | Microsoft.PowerShell.Utility | 7 only | 7 only | Sets the colors and styles used for rendering Markdown content in the console. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Set-PackageSource`](#set-packagesource) | PackageManagement | Both | Syntax differs | Replaces a package source for a specified package provider. | Not implemented |  |
| [`Set-PSBreakpoint`](#set-psbreakpoint) | Microsoft.PowerShell.Utility | Both | Syntax differs | Sets a breakpoint on a line, command, or variable. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Set-PSDebug`](#set-psdebug) | Microsoft.PowerShell.Core | Both | None | Turns script debugging features on and off, sets the trace level, and toggles strict mode. | Not implemented | Platform / miscellaneous |
| [`Set-PSReadLineKeyHandler`](#set-psreadlinekeyhandler) | PSReadLine | Both | None | Binds keys to user-defined or PSReadLine key handler functions. | Not implemented |  |
| [`Set-PSReadLineOption`](#set-psreadlineoption) | PSReadLine | Both | Syntax differs | Customizes the behavior of command line editing in **PSReadLine**. | Not implemented |  |
| [`Set-PSResourceRepository`](#set-psresourcerepository) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Sets information for a registered repository. | Not implemented |  |
| [`Set-StrictMode`](#set-strictmode) | Microsoft.PowerShell.Core | Both | None | Establishes and enforces coding rules in expressions, scripts, and scriptblocks. | Not implemented | Platform / miscellaneous |
| [`Set-TraceSource`](#set-tracesource) | Microsoft.PowerShell.Utility | Both | None | Configures, starts, and stops a trace of PowerShell components. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Set-Variable`](#set-variable) | Microsoft.PowerShell.Utility | Both | Syntax differs | Sets the value of a variable. Creates the variable if one with the requested name does not exist. | Go implementation | Assignment to read-only automatic variables (PID etc.) is rejected. |
| [`Show-Markdown`](#show-markdown) | Microsoft.PowerShell.Utility | 7 only | 7 only | Shows a Markdown file or string in the console in a friendly way using VT100 escape sequences or in a browser using HTML. | Not implemented | GUI and printing |
| [`Sort-Object`](#sort-object) | Microsoft.PowerShell.Utility | Both | Syntax differs | Sorts objects by property values. | Go implementation | Stable sort (PowerShell sorting is unstable with no order guarantee for equal elements; this program preserves input order). |
| [`Split-Path`](#split-path) | Microsoft.PowerShell.Management | Both | Syntax differs | Returns the specified part of a path. | Go implementation |  |
| [`Start-Job`](#start-job) | Microsoft.PowerShell.Core | Both | Syntax differs | Starts a PowerShell background job. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Start-Process`](#start-process) | Microsoft.PowerShell.Management | Both | Syntax differs | Starts one or more processes on the local computer. | Go implementation | Does not take over the new process's input/output. |
| [`Start-Sleep`](#start-sleep) | Microsoft.PowerShell.Utility | Both | Syntax differs | Suspends the activity in a script or session for the specified period of time. | Go implementation |  |
| [`Start-ThreadJob`](#start-threadjob) | Microsoft.PowerShell.ThreadJob | 7 only | 7 only | Creates background jobs similar to the `Start-Job` cmdlet. | Not implemented |  |
| [`Start-Transcript`](#start-transcript) | Microsoft.PowerShell.Host | Both | Syntax differs | Creates a record of all or part of a PowerShell session to a text file. | Not implemented |  |
| [`Stop-Computer`](#stop-computer) | Microsoft.PowerShell.Management | Both | Syntax differs | Stops (shuts down) local and remote computers. | Mapped Linux (sudo reboot / shutdown / hostnamectl) |  |
| [`Stop-Job`](#stop-job) | Microsoft.PowerShell.Core | Both | None | Stops a PowerShell background job. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Stop-Process`](#stop-process) | Microsoft.PowerShell.Management | Both | None | Stops one or more running processes. | Go implementation |  |
| [`Stop-Transcript`](#stop-transcript) | Microsoft.PowerShell.Host | Both | Syntax differs | Stops a transcript. | Not implemented |  |
| [`Switch-Process`](#switch-process) | Microsoft.PowerShell.Core | 7 only | 7 only on Linux/macOS (absent from Windows builds) | On Linux and macOS, the cmdlet calls the execv() function to provide similar behavior as POSIX
shells. | Not implemented | Platform / miscellaneous |
| [`Tee-Object`](#tee-object) | Microsoft.PowerShell.Utility | Both | Syntax differs | Saves command output in a file or variable and also sends it down the pipeline. | Go implementation |  |
| [`Test-Connection`](#test-connection) | Microsoft.PowerShell.Management | Both | Syntax differs | Sends ICMP echo request packets, or pings, to one or more computers. | Mapped Linux (ping) |  |
| [`Test-Json`](#test-json) | Microsoft.PowerShell.Utility | 7 only | 7 only | Tests whether a string is a valid JSON document | Go implementation |  |
| [`Test-ModuleManifest`](#test-modulemanifest) | Microsoft.PowerShell.Core | Both | None | Verifies that a module manifest file accurately describes the contents of a module. | Not implemented | Modules and assemblies (not applicable to a single-file interpreter) |
| [`Test-Path`](#test-path) | Microsoft.PowerShell.Management | Both | Syntax differs | Determines whether all elements of a path exist. | Go implementation |  |
| [`Test-PSScriptFileInfo`](#test-psscriptfileinfo) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Tests the comment-based metadata in a .ps1 file to ensure it's valid for publication. | Not implemented |  |
| [`Trace-Command`](#trace-command) | Microsoft.PowerShell.Utility | Both | None | Configures and starts a trace of the specified expression or command. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Unblock-File`](#unblock-file) | Microsoft.PowerShell.Utility | Both | None | Unblocks files that were downloaded from the internet. | Not implemented | Platform / miscellaneous |
| [`Uninstall-Package`](#uninstall-package) | PackageManagement | Both | Syntax differs | Uninstalls one or more software packages. | Not implemented |  |
| [`Uninstall-PSResource`](#uninstall-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Uninstalls a resource that was installed using PowerShellGet. | Not implemented |  |
| [`Unprotect-CmsMessage`](#unprotect-cmsmessage) | Microsoft.PowerShell.Security | Both | None | Decrypts content that has been encrypted by using the Cryptographic Message Syntax format. | Not implemented |  |
| [`Unregister-Event`](#unregister-event) | Microsoft.PowerShell.Utility | Both | None | Cancels an event subscription. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Unregister-PackageSource`](#unregister-packagesource) | PackageManagement | Both | Syntax differs | Removes a registered package source. | Not implemented |  |
| [`Unregister-PSResourceRepository`](#unregister-psresourcerepository) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Removes a registered repository from the local machine. | Not implemented |  |
| [`Update-FormatData`](#update-formatdata) | Microsoft.PowerShell.Utility | Both | None | Updates the formatting data in the current session. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Update-Help`](#update-help) | Microsoft.PowerShell.Core | Both | Syntax differs | Downloads and installs the newest help files on your computer. | Not implemented | Platform / miscellaneous |
| [`Update-List`](#update-list) | Microsoft.PowerShell.Utility | Both | None | Adds items to and removes items from a property value that contains a collection of objects. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Update-PSModuleManifest`](#update-psmodulemanifest) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Updates a module manifest file. | Not implemented |  |
| [`Update-PSResource`](#update-psresource) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | Downloads and installs the newest version of a package already installed on the local machine. | Not implemented |  |
| [`Update-PSScriptFileInfo`](#update-psscriptfileinfo) | Microsoft.PowerShell.PSResourceGet | 7 only | 7 only | This cmdlet updates the comment-based metadata in an existing script .ps1 file. | Not implemented |  |
| [`Update-TypeData`](#update-typedata) | Microsoft.PowerShell.Utility | Both | None | Updates the extended type data in the session. | Not implemented | Serialization / markup / formatting (rarely used) |
| [`Wait-Debugger`](#wait-debugger) | Microsoft.PowerShell.Utility | Both | None | Stops a script in the debugger before running the next statement in the script. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Wait-Event`](#wait-event) | Microsoft.PowerShell.Utility | Both | None | Waits until a particular event is raised before continuing to run. | Not implemented | Events / breakpoints / tracing (debugger facilities) |
| [`Wait-Job`](#wait-job) | Microsoft.PowerShell.Core | Both | None | Waits until one or all of the PowerShell jobs running in the session are in a terminating state. | Not implemented | Jobs and runspaces (requires job facilities; out of scope) |
| [`Wait-Process`](#wait-process) | Microsoft.PowerShell.Management | Both | Syntax differs | Waits for the processes to be stopped before accepting more input. | Go implementation |  |
| [`Where-Object`](#where-object) | Microsoft.PowerShell.Core | Both | Syntax differs | Selects objects from a collection based on their property values. | Go implementation |  |
| [`Write-Debug`](#write-debug) | Microsoft.PowerShell.Utility | Both | None | Writes a debug message to the console. | Go implementation |  |
| [`Write-Error`](#write-error) | Microsoft.PowerShell.Utility | Both | Syntax differs | Writes an object to the error stream. | Go implementation | Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG). |
| [`Write-Host`](#write-host) | Microsoft.PowerShell.Utility | Both | None | Writes customized output to a host. | Go implementation |  |
| [`Write-Information`](#write-information) | Microsoft.PowerShell.Utility | Both | None | Specifies how PowerShell handles information stream data for a command. | Go implementation |  |
| [`Write-Output`](#write-output) | Microsoft.PowerShell.Utility | Both | Syntax differs | Writes the specified objects to the pipeline. | Go implementation |  |
| [`Write-Progress`](#write-progress) | Microsoft.PowerShell.Utility | Both | Syntax differs | Displays a progress bar within a PowerShell command window. | Not implemented | GUI and printing |
| [`Write-Verbose`](#write-verbose) | Microsoft.PowerShell.Utility | Both | None | Writes text to the verbose message stream. | Go implementation |  |
| [`Write-Warning`](#write-warning) | Microsoft.PowerShell.Utility | Both | None | Writes a warning message. | Go implementation |  |

## Command Details

### Add-Content

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Add-Content [-Path] <string[]> [-Value] <Object[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-NoNewline] [-Encoding <FileSystemCmdletProviderEncoding>] [-Stream <string>] [<CommonParameters>]
Add-Content [-Value] <Object[]> -LiteralPath <string[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-NoNewline] [-Encoding <FileSystemCmdletProviderEncoding>] [-Stream <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Add-Content [-Path] <string[]> [-Value] <Object[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-NoNewline] [-Encoding <Encoding>] [-AsByteStream] [-Stream <string>] [<CommonParameters>]
Add-Content [-Value] <Object[]> -LiteralPath <string[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-NoNewline] [-Encoding <Encoding>] [-AsByteStream] [-Stream <string>] [<CommonParameters>]
```

Example: Add a string to all text files with an exception

```powershell
Add-Content -Path .\*.txt -Exclude help* -Value 'End of file'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Add-Content.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: The `ac` alias is not implemented; use the full name.

- Type: Go implementation.
- Function: appends to files.
- Difference from Windows: the `ac` alias is not implemented; use the full name.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target file |
| `-Value` (position 1) | object | Content to append; pipeline input takes precedence |

- Implementation: append mode (O_APPEND), `echo ... >> file` territory.


### Add-History

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Add-History [[-InputObject] <psobject[]>] [-Passthru] [<CommonParameters>]
```

Example: Add commands to the history of a different session

```powershell
Get-History | Export-Csv -Path C:\testing\history.csv -IncludeTypeInformation
Import-Csv -Path C:\testing\history.csv | Add-History
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Add-History.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Get-History, Clear-History, Invoke-History.
- Function: appends a history entry.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Command text; pipeline input takes precedence |


### Add-Member

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Add-Member -InputObject <psobject> -TypeName <string> [-PassThru] [<CommonParameters>]
Add-Member [-NotePropertyName] <string> [-NotePropertyValue] <Object> -InputObject <psobject> [-TypeName <string>] [-Force] [-PassThru] [<CommonParameters>]
Add-Member [-MemberType] <PSMemberTypes> [-Name] <string> [[-Value] <Object>] [[-SecondValue] <Object>] -InputObject <psobject> [-TypeName <string>] [-Force] [-PassThru] [<CommonParameters>]
Add-Member [-NotePropertyMembers] <IDictionary> -InputObject <psobject> [-TypeName <string>] [-Force] [-PassThru] [<CommonParameters>]
```

Syntax (7):

```powershell
Add-Member -InputObject <psobject> -TypeName <string> [-PassThru] [<CommonParameters>]
Add-Member [-MemberType] <PSMemberTypes> [-Name] <string> [[-Value] <Object>] [[-SecondValue] <Object>] -InputObject <psobject> [-TypeName <string>] [-Force] [-PassThru] [<CommonParameters>]
Add-Member [-NotePropertyName] <string> [-NotePropertyValue] <Object> -InputObject <psobject> [-TypeName <string>] [-Force] [-PassThru] [<CommonParameters>]
Add-Member [-NotePropertyMembers] <IDictionary> -InputObject <psobject> [-TypeName <string>] [-Force] [-PassThru] [<CommonParameters>]
```

Example: Add a note property to a PSObject

```powershell
$A = Get-ChildItem C:\ps-test\test.txt
$A | Add-Member -NotePropertyName Status -NotePropertyValue Done
$A.Status
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Add-Member.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Only the NoteProperty type is supported.

- Type: Go implementation.
- Function: adds properties to objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to receive the property |
| `-MemberType` | string | NoteProperty only; other values ignored |
| `-Name` | string | New property name |
| `-Value` | object | Property value |
| `-Force` | switch | Overwrites even when the property already exists |

- Implementation: copies the object, appends the property, and returns the new object.


### Add-Type

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Add-Type [-TypeDefinition] <string> [-CodeDomProvider <CodeDomProvider>] [-CompilerParameters <CompilerParameters>] [-Language <Language>] [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [<CommonParameters>]
Add-Type [-Name] <string> [-MemberDefinition] <string[]> [-CodeDomProvider <CodeDomProvider>] [-CompilerParameters <CompilerParameters>] [-Namespace <string>] [-UsingNamespace <string[]>] [-Language <Language>] [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [<CommonParameters>]
Add-Type [-Path] <string[]> [-CompilerParameters <CompilerParameters>] [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [<CommonParameters>]
Add-Type -LiteralPath <string[]> [-CompilerParameters <CompilerParameters>] [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [<CommonParameters>]
Add-Type -AssemblyName <string[]> [-PassThru] [-IgnoreWarnings] [<CommonParameters>]
```

Syntax (7):

```powershell
Add-Type [-TypeDefinition] <string> [-Language <Language>] [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [-CompilerOptions <string[]>] [<CommonParameters>]
Add-Type [-Name] <string> [-MemberDefinition] <string[]> [-Namespace <string>] [-UsingNamespace <string[]>] [-Language <Language>] [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [-CompilerOptions <string[]>] [<CommonParameters>]
Add-Type [-Path] <string[]> [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [-CompilerOptions <string[]>] [<CommonParameters>]
Add-Type -LiteralPath <string[]> [-ReferencedAssemblies <string[]>] [-OutputAssembly <string>] [-OutputType <OutputAssemblyType>] [-PassThru] [-IgnoreWarnings] [-CompilerOptions <string[]>] [<CommonParameters>]
Add-Type -AssemblyName <string[]> [-PassThru] [<CommonParameters>]
```

Example: Add a .NET type to a session

```powershell
$Source = @"
public class BasicTest
{
  public static int Add(int a, int b)
    {
        return (a + b);
    }
  public int Multiply(int a, int b)
    {
    return (a * b);
    }
}
"@

Add-Type -TypeDefinition $Source
[BasicTest]::Add(4, 3)
$BasicTestObject = New-Object BasicTest
$BasicTestObject.Multiply(5, 2)
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Add-Type.md)


### Clear-Content

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Clear-Content [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-Stream <string>] [<CommonParameters>]
Clear-Content -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-Stream <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Clear-Content [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-Stream <string>] [<CommonParameters>]
Clear-Content -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-Stream <string>] [<CommonParameters>]
```

Example: Delete all content from a directory

```powershell
Clear-Content "..\SmpUsers\*\init.txt"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Clear-Content.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: none. Only existing files are emptied; nonexistent paths raise an error.

- Type: Go implementation.
- Function: empties files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target file |

- Implementation: `os.WriteFile(path, empty, 0644)`, `: > file` / `truncate -s 0` territory.


### Clear-History

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Clear-History [[-Id] <int[]>] [[-Count] <int>] [-Newest] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-History [[-Count] <int>] [-CommandLine <string[]>] [-Newest] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Delete the command history from a PowerShell session

```powershell
Get-History
Clear-History
Get-History
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Clear-History.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Get-History, Add-History, Invoke-History.
- Function: clears history. Bash's `history -c`.
- Parameters: none.


### Clear-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Clear-Item [-Path] <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Clear-Item -LiteralPath <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Clear-Item [-Path] <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-Item -LiteralPath <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Clear the value of a variable

```powershell
Clear-Item Variable:TestVar1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Clear-Item.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Function: empties files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path; directories are left alone |


### Clear-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Clear-ItemProperty [-Path] <string[]> [-Name] <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Clear-ItemProperty [-Name] <string> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Clear-ItemProperty [-Path] <string[]> [-Name] <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-ItemProperty [-Name] <string> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Clear the value of registry key

```powershell
Clear-ItemProperty -Path "HKLM:\Software\MyCompany\MyApp" -Name "Options"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Clear-ItemProperty.md)


### Clear-Variable

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Clear-Variable [-Name] <string[]> [-Include <string[]>] [-Exclude <string[]>] [-Force] [-PassThru] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove the value of global variables that begin with a search string

```powershell
Clear-Variable my* -Scope Global
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Clear-Variable.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Variable, New-Variable, Remove-Variable, Clear-Variable.
- Function: sets/creates/removes/empties variables.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable name |
| `-Value` (position 1) | object | Value (New-Variable's `-Value` also takes position 1) |
| `-Force` | switch | Lets New-Variable overwrite an existing variable |

- Behavior: creating over an existing variable without -Force → error; assigning to read-only automatic variables (PID etc.) → error with Set/New, silently ignored with Clear. Remove deletes straight from the map; Clear sets $null.


### Compare-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Compare-Object [-ReferenceObject] <psobject[]> [-DifferenceObject] <psobject[]> [-SyncWindow <int>] [-Property <Object[]>] [-ExcludeDifferent] [-IncludeEqual] [-PassThru] [-Culture <string>] [-CaseSensitive] [<CommonParameters>]
```

Example: Compare the content of two text files

```powershell
$objects = @{
  ReferenceObject = (Get-Content -Path C:\Test\Testfile1.txt)
  DifferenceObject = (Get-Content -Path C:\Test\Testfile2.txt)
}
Compare-Object @objects
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Compare-Object.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original.

- Type: Go implementation.
- Function: compares sets for differences. `diff` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-ReferenceObject` (position 0) | object | Reference set |
| `-DifferenceObject` (position 1) | object | Set to compare against it |
| `-CaseSensitive` | switch | Case-sensitive comparison; insensitive by default |
| `-IncludeEqual` | switch | Also outputs equal entries (`==`), differences only by default |

- Output: PSCustomObjects with fields InputObject, SideIndicator. `<=` reference-set only, `=>` comparison-set only, `==` both sides (needs -IncludeEqual). Case-insensitive by default; output order runs equals first, then right side (=>), then left side (<=). Equal entries display the reference set's values.


### Compress-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Compress-PSResource [-Path] <string> [-DestinationPath] <string> [-PassThru] [-SkipModuleManifestValidate] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Compress-PSResource -Path C:\TestModule -DestinationPath C:\NupkgDestination
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/compress-psresource?view=powershell-7.5)


### Convert-Path

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Convert-Path [-Path] <string[]> [-UseTransaction] [<CommonParameters>]
Convert-Path -LiteralPath <string[]> [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Convert-Path [-Path] <string[]> [<CommonParameters>]
Convert-Path -LiteralPath <string[]> [<CommonParameters>]
```

Example: Convert the working directory to a standard file system path

```powershell
PS C:\> Convert-Path .
C:\
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Convert-Path.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companion: Resolve-Path.
- Function: converts to absolute paths (symbolic links untouched).

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |


### ConvertFrom-CliXml

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertFrom-CliXml [-InputObject] <string> [<CommonParameters>]
```

Example: Convert a process object to CliXml and back

```powershell
$process = Get-Process -Id $PID
$process.pstypenames
```

```powershell
$process | Get-Member | Group-Object MemberType | Select-Object Name, Count
```

```powershell
$xml = $process | ConvertTo-CliXml
$fromXML = ConvertFrom-CliXml $xml
$fromXML.pstypenames
```

```powershell
$fromXML | Get-Member | Group-Object MemberType | Select-Object Name, Count
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertFrom-CliXml.md)


### ConvertFrom-Csv

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertFrom-Csv [-InputObject] <psobject[]> [[-Delimiter] <char>] [-Header <string[]>] [<CommonParameters>]
ConvertFrom-Csv [-InputObject] <psobject[]> -UseCulture [-Header <string[]>] [<CommonParameters>]
```

Example: Convert processes on the local computer to CSV format

```powershell
$P = Get-Process | ConvertTo-Csv
$P | ConvertFrom-Csv
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertFrom-Csv.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companion: ConvertTo-Csv.
- Function: turns CSV text into objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | CSV text; pipeline input takes precedence |

- Implementation: first line is the header, remaining lines produce PSCustomObjects.


### ConvertFrom-Json

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
ConvertFrom-Json [-InputObject] <string> [<CommonParameters>]
```

Syntax (7):

```powershell
ConvertFrom-Json [-InputObject] <string> [-AsHashtable] [-Depth <int>] [-NoEnumerate] [-DateKind <JsonDateKind>] [<CommonParameters>]
```

Example: Convert a DateTime object to a JSON object

```powershell
Get-Date | Select-Object -Property * | ConvertTo-Json | ConvertFrom-Json
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertFrom-Json.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companion: ConvertTo-Json.
- Function: turns JSON into objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | JSON text; pipeline input takes precedence |

- Implementation: parses text into PSCustomObjects (properties), arrays, and scalars.


### ConvertFrom-Markdown

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertFrom-Markdown [-Path] <string[]> [-AsVT100EncodedString] [<CommonParameters>]
ConvertFrom-Markdown -LiteralPath <string[]> [-AsVT100EncodedString] [<CommonParameters>]
ConvertFrom-Markdown -InputObject <psobject> [-AsVT100EncodedString] [<CommonParameters>]
```

Example: Convert a file containing Markdown content to HTML

```powershell
ConvertFrom-Markdown -Path .\README.md
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertFrom-Markdown.md)


### ConvertFrom-SecureString

Version: Both

Module: Microsoft.PowerShell.Security

Syntax (5.1):

```powershell
ConvertFrom-SecureString [-SecureString] <securestring> [[-SecureKey] <securestring>] [<CommonParameters>]
ConvertFrom-SecureString [-SecureString] <securestring> [-Key <byte[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
ConvertFrom-SecureString [-SecureString] <securestring> [[-SecureKey] <securestring>] [<CommonParameters>]
ConvertFrom-SecureString [-SecureString] <securestring> [-AsPlainText] [<CommonParameters>]
ConvertFrom-SecureString [-SecureString] <securestring> [-Key <byte[]>] [<CommonParameters>]
```

Example: Create a secure string

```powershell
$SecureString = Read-Host -AsSecureString
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/ConvertFrom-SecureString.md)


### ConvertFrom-StringData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
ConvertFrom-StringData [-StringData] <string> [<CommonParameters>]
```

Syntax (7):

```powershell
ConvertFrom-StringData [-StringData] <string> [[-Delimiter] <char>] [<CommonParameters>]
```

Example: Convert a single-quoted here-string to a hash table

```powershell
$Here = @'
Msg1 = The string parameter is required.
Msg2 = Credentials are required for this command.
Msg3 = The specified variable doesn't exist.
'@
ConvertFrom-StringData -StringData $Here
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertFrom-StringData.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: parses `key=value` text into a Hashtable.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-StringData` (position 0) | string | `key=value` text; named/positional take precedence, pipeline input only as fallback |

- Implementation: goes line by line, skipping blank lines and `#` comments, splitting keys from values at `=` with TrimSpace.


### ConvertTo-CliXml

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertTo-CliXml [-InputObject] <psobject> [-Depth <int>] [<CommonParameters>]
```

Example: Convert a process object to CliXml and back

```powershell
$process = Get-Process -Id $PID
$process.pstypenames
```

```powershell
$process | Get-Member | Group-Object MemberType | Select-Object Name, Count
```

```powershell
$xml = $process | ConvertTo-CliXml
$fromXML = ConvertFrom-CliXml $xml
$fromXML.pstypenames
```

```powershell
$fromXML | Get-Member | Group-Object MemberType | Select-Object Name, Count
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertTo-CliXml.md)


### ConvertTo-Csv

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
ConvertTo-Csv [-InputObject] <psobject> [[-Delimiter] <char>] [-NoTypeInformation] [<CommonParameters>]
ConvertTo-Csv [-InputObject] <psobject> [-UseCulture] [-NoTypeInformation] [<CommonParameters>]
```

Syntax (7):

```powershell
ConvertTo-Csv [-InputObject] <psobject> [[-Delimiter] <char>] [-IncludeTypeInformation] [-NoTypeInformation] [-QuoteFields <string[]>] [-UseQuotes <QuoteKind>] [-NoHeader] [<CommonParameters>]
ConvertTo-Csv [-InputObject] <psobject> [-UseCulture] [-IncludeTypeInformation] [-NoTypeInformation] [-QuoteFields <string[]>] [-UseQuotes <QuoteKind>] [-NoHeader] [<CommonParameters>]
```

Example (5.1): Convert an object to CSV

```powershell
Get-Process -Name 'PowerShell' | ConvertTo-Csv -NoTypeInformation
```

Example (7): Convert an object to CSV

```powershell
Get-Process -Name pwsh | ConvertTo-Csv -NoTypeInformation
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertTo-Csv.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Companion: ConvertFrom-Csv.
- Function: turns objects into CSV text.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Input objects |
| `-Property` | string[] | Chooses columns; omitted, takes the first object's property names (single-column "Value" for property-less objects) |


### ConvertTo-Html

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
ConvertTo-Html [[-Property] <Object[]>] [[-Head] <string[]>] [[-Title] <string>] [[-Body] <string[]>] [-InputObject <psobject>] [-As <string>] [-CssUri <uri>] [-PostContent <string[]>] [-PreContent <string[]>] [<CommonParameters>]
ConvertTo-Html [[-Property] <Object[]>] [-InputObject <psobject>] [-As <string>] [-Fragment] [-PostContent <string[]>] [-PreContent <string[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
ConvertTo-Html [[-Property] <Object[]>] [[-Head] <string[]>] [[-Title] <string>] [[-Body] <string[]>] [-InputObject <psobject>] [-As <string>] [-CssUri <uri>] [-PostContent <string[]>] [-PreContent <string[]>] [-Meta <hashtable>] [-Charset <string>] [-Transitional] [<CommonParameters>]
ConvertTo-Html [[-Property] <Object[]>] [-InputObject <psobject>] [-As <string>] [-Fragment] [-PostContent <string[]>] [-PreContent <string[]>] [<CommonParameters>]
```

Example: Create a web page to display the date

```powershell
ConvertTo-Html -InputObject (Get-Date)
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertTo-Html.md)


### ConvertTo-Json

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
ConvertTo-Json [-InputObject] <Object> [-Depth <int>] [-Compress] [<CommonParameters>]
```

Syntax (7):

```powershell
ConvertTo-Json [-InputObject] <Object> [-Depth <int>] [-Compress] [-EnumsAsStrings] [-AsArray] [-EscapeHandling <StringEscapeHandling>] [<CommonParameters>]
```

Example: 

```powershell
(Get-UICulture).Calendar | ConvertTo-Json
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertTo-Json.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Output uses a fixed 2-space indent.

- Type: Go implementation.
- Companion: ConvertFrom-Json.
- Function: turns objects into JSON. Bash's `jq` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to convert |
| `-Depth` | int | Nesting depth cap (default 2); beyond it arrays/objects print as empty shells (`[]`/`{}`) |

- Rules: single object → object JSON; several objects → array. Property-bearing objects → `{"property": value}`; Hashtable → key-value object; arrays recurse; scalars stay JSON scalars; $null → null. Pretty-printed with a 2-space indent.


### ConvertTo-SecureString

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
ConvertTo-SecureString [-String] <string> [[-SecureKey] <securestring>] [<CommonParameters>]
ConvertTo-SecureString [-String] <string> [-AsPlainText] [-Force] [<CommonParameters>]
ConvertTo-SecureString [-String] <string> [-Key <byte[]>] [<CommonParameters>]
```

Example: Convert a secure string to an encrypted string

```powershell
PS C:\> $Secure = Read-Host -AsSecureString
PS C:\> $Secure
System.Security.SecureString
PS C:\> $Encrypted = ConvertFrom-SecureString -SecureString $Secure
PS C:\> $Encrypted
01000000d08c9ddf0115d1118c7a00c04fc297eb010000001a114d45b8dd3f4aa11ad7c0abdae98000000000
02000000000003660000a8000000100000005df63cea84bfb7d70bd6842e7efa79820000000004800000a000
000010000000f10cd0f4a99a8d5814d94e0687d7430b100000008bf11f1960158405b2779613e9352c6d1400
0000e6b7bf46a9d485ff211b9b2a2df3bd6eb67aae41
PS C:\> $Secure2 = ConvertTo-SecureString -String $Encrypted
PS C:\> $Secure2
System.Security.SecureString
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/ConvertTo-SecureString.md)


### ConvertTo-Xml

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertTo-Xml [-InputObject] <psobject> [-Depth <int>] [-NoTypeInformation] [-As <string>] [<CommonParameters>]
```

Example: Convert a date to XML

```powershell
Get-Date | ConvertTo-Xml
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertTo-Xml.md)


### Copy-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Copy-Item [-Path] <string[]> [[-Destination] <string>] [-Container] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-FromSession <PSSession>] [-ToSession <PSSession>] [<CommonParameters>]
Copy-Item [[-Destination] <string>] -LiteralPath <string[]> [-Container] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-FromSession <PSSession>] [-ToSession <PSSession>] [<CommonParameters>]
```

Syntax (7):

```powershell
Copy-Item [-Path] <string[]> [[-Destination] <string>] [-Container] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-FromSession <PSSession>] [-ToSession <PSSession>] [<CommonParameters>]
Copy-Item [[-Destination] <string>] -LiteralPath <string[]> [-Container] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-FromSession <PSSession>] [-ToSession <PSSession>] [<CommonParameters>]
```

Example: Copy a file to the specified directory

```powershell
Copy-Item "C:\Wabash\Logfiles\mar1604.log.txt" -Destination "C:\Presentation"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Copy-Item.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: copies.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Source paths, multiple allowed |
| `-Destination` (position 1) | path | Destination path |
| `-Recurse` | switch | Copies a whole directory tree, `cp -r` |

- Behavior: copying a directory without -Recurse → error; an already-existing destination directory receives the source inside it (`cp` semantics); with multiple sources the destination must be an existing directory, otherwise error.


### Copy-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Copy-ItemProperty [-Path] <string[]> [-Destination] <string> [-Name] <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Copy-ItemProperty [-Destination] <string> [-Name] <string> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Copy-ItemProperty [-Path] <string[]> [-Destination] <string> [-Name] <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Copy-ItemProperty [-Destination] <string> [-Name] <string> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Copy a property from a registry key to another registry key

```powershell
$params = @{
    Path        = 'MyApplication'
    Destination = 'HKLM:\Software\MyApplicationRev2'
    Name        = 'MyProperty'
}
Copy-ItemProperty @params
```

Example (7): Copy a property from a registry key to another registry key

```powershell
$copyParams = @{
    Path        = "MyApplication"
    Destination = "HKLM:\Software\MyApplicationRev2"
    Name        = "MyProperty"
}
Copy-ItemProperty @copyParams
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Copy-ItemProperty.md)


### Debug-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Debug-Job [-Job] <Job> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Job [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Job [-Id] <int> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Job [-InstanceId] <guid> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Debug-Job [-Job] <Job> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Job [-Name] <string> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Job [-Id] <int> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Job [-InstanceId] <guid> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Debug a job by job ID

```powershell
Debug-Job -Id 3
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Debug-Job.md)


### Debug-Process

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Debug-Process [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Process [-Id] <int[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Process -InputObject <Process[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Attach a debugger to a process on the computer

```powershell
Debug-Process -Name powershell
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Debug-Process.md)


### Debug-Runspace

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Debug-Runspace [-Runspace] <runspace> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Runspace [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Runspace [-Id] <int> [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Runspace [-InstanceId] <guid> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Debug-Runspace [-Runspace] <runspace> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Runspace [-Name] <string> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Runspace [-Id] <int> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
Debug-Runspace [-InstanceId] <guid> [-BreakAll] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Debug a remote runspace

```powershell
PS C:\> Get-Process -ComputerName "WS10TestServer" -Name "*powershell*"

Handles      WS(K)   VM(M)      CPU(s)    Id  ProcessName
-------      -----   -----      ------    --  -----------
    377      69912     63     2.09      2420  powershell
    399     123396    829     4.48      1152  powershell_ise

PS C:\> Enter-PSSession -ComputerName "WS10TestServer"
[WS10TestServer]:PS C:\> Enter-PSHostProcess -Id 1152
[WS10TestServer:][Process:1152]: PS C:\Users\Test\Documents> Get-Runspace

Id Name            ComputerName    Type          State         Availability
-- ----            ------------    ----          -----         ------------
 1 Runspace1       WS10TestServer  Remote        Opened        Available
 2 RemoteHost      WS10TestServer  Remote        Opened        Busy

[WS10TestServer][Process:1152]: PS C:\Users\Test\Documents> Debug-Runspace -Id 2

Hit Line breakpoint on 'C:\TestWFVar1.ps1:83'
At C:\TestWFVar1.ps1:83 char:1
+ $scriptVar = "Script Variable"
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
[Process:1152]: [RSDBG: 2]: PS C:\> >
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Debug-Runspace.md)


### Disable-ExperimentalFeature

Version: 7 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Disable-ExperimentalFeature [-Name] <string[]> [-Scope <ConfigScope>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable an experimental feature

```powershell
Disable-ExperimentalFeature -Name PSImplicitRemotingBatching
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Disable-ExperimentalFeature.md)


### Disable-PSBreakpoint

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Disable-PSBreakpoint [-Breakpoint] <Breakpoint[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-PSBreakpoint [-Id] <int[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Disable-PSBreakpoint [-Breakpoint] <Breakpoint[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-PSBreakpoint [-Id] <int[]> [-PassThru] [-Runspace <runspace>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set a breakpoint and disable it

```powershell
$B = Set-PSBreakpoint -Script "sample.ps1" -Variable "name"
$B | Disable-PSBreakpoint
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Disable-PSBreakpoint.md)


### Disable-RunspaceDebug

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Disable-RunspaceDebug [[-RunspaceName] <string[]>] [<CommonParameters>]
Disable-RunspaceDebug [-Runspace] <runspace[]> [<CommonParameters>]
Disable-RunspaceDebug [-RunspaceId] <int[]> [<CommonParameters>]
Disable-RunspaceDebug [-RunspaceInstanceId] <guid[]> [<CommonParameters>]
Disable-RunspaceDebug [[-ProcessName] <string>] [[-AppDomainName] <string[]>] [<CommonParameters>]
```

Example: Disable the default runspace debugger

```powershell
Disable-RunspaceDebug
Get-RunspaceDebug
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Disable-RunspaceDebug.md)


### Enable-ExperimentalFeature

Version: 7 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Enable-ExperimentalFeature [-Name] <string[]> [-Scope <ConfigScope>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Enable an experimental feature

```powershell
Enable-ExperimentalFeature PSImplicitRemotingBatching
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Enable-ExperimentalFeature.md)


### Enable-PSBreakpoint

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Enable-PSBreakpoint [-Id] <int[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-PSBreakpoint [-Breakpoint] <Breakpoint[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Enable-PSBreakpoint [-Breakpoint] <Breakpoint[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-PSBreakpoint [-Id] <int[]> [-PassThru] [-Runspace <runspace>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Enable all breakpoints

```powershell
Get-PSBreakpoint | Enable-PSBreakpoint
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Enable-PSBreakpoint.md)


### Enable-RunspaceDebug

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Enable-RunspaceDebug [[-RunspaceName] <string[]>] [-BreakAll] [<CommonParameters>]
Enable-RunspaceDebug [-Runspace] <runspace[]> [-BreakAll] [<CommonParameters>]
Enable-RunspaceDebug [-RunspaceId] <int[]> [-BreakAll] [<CommonParameters>]
Enable-RunspaceDebug [-RunspaceInstanceId] <guid[]> [<CommonParameters>]
Enable-RunspaceDebug [[-ProcessName] <string>] [[-AppDomainName] <string[]>] [<CommonParameters>]
```

Example: Enable the default runspace debugger

```powershell
Enable-RunspaceDebug
Get-RunspaceDebug
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Enable-RunspaceDebug.md)


### Enter-PSHostProcess

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Enter-PSHostProcess [-Id] <int> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess [-Process] <Process> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess [-Name] <string> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess [-HostProcessInfo] <PSHostProcessInfo> [[-AppDomainName] <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Enter-PSHostProcess [-Id] <int> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess [-Process] <Process> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess [-Name] <string> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess [-HostProcessInfo] <PSHostProcessInfo> [[-AppDomainName] <string>] [<CommonParameters>]
Enter-PSHostProcess -CustomPipeName <string> [<CommonParameters>]
```

Example: Example Part 1: Start debugging a runspace within the PowerShell ISE process

```powershell
PS C:\> Enter-PSHostProcess -Name powershell_ise
[Process:1520]: PS C:\>  Get-Runspace
Id    Name          InstanceId                               State           Availability
--    -------       -----------                              ------          -------------
1     Runspace1     2d91211d-9cce-42f0-ab0e-71ac258b32b5     Opened          Available
2     Runspace2     a3855043-cb16-424a-a616-685360c3763b     Opened          RemoteDebug
3     MyLocalRS     2236dbd8-2105-4dec-a15a-a27d0bfaacb5     Opened          LocalDebug
4     MyRunspace    771356e9-8c44-4b70-9de5-dd17cb41e48e     Opened          Busy
5     Runspace8     3e517382-a97a-49ba-9c3c-fd21f6664288     Broken          None
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Enter-PSHostProcess.md)


### Enter-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Enter-PSSession [-ComputerName] <string> [-EnableNetworkAccess] [-Credential <pscredential>] [-ConfigurationName <string>] [-Port <int>] [-UseSSL] [-ApplicationName <string>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Enter-PSSession [[-Session] <PSSession>] [<CommonParameters>]
Enter-PSSession [[-ConnectionUri] <uri>] [-EnableNetworkAccess] [-Credential <pscredential>] [-ConfigurationName <string>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Enter-PSSession [-InstanceId <guid>] [<CommonParameters>]
Enter-PSSession [[-Id] <int>] [<CommonParameters>]
Enter-PSSession [-Name <string>] [<CommonParameters>]
Enter-PSSession [-VMId] <guid> [-Credential] <pscredential> [-ConfigurationName <string>] [<CommonParameters>]
Enter-PSSession [-VMName] <string> [-Credential] <pscredential> [-ConfigurationName <string>] [<CommonParameters>]
Enter-PSSession [-ContainerId] <string> [-ConfigurationName <string>] [-RunAsAdministrator] [<CommonParameters>]
```

Syntax (7):

```powershell
Enter-PSSession [-ComputerName] <string> [-EnableNetworkAccess] [-Credential <pscredential>] [-ConfigurationName <string>] [-Port <int>] [-UseSSL] [-ApplicationName <string>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Enter-PSSession [-HostName] <string> [-Options <hashtable>] [-Port <int>] [-UserName <string>] [-KeyFilePath <string>] [-Subsystem <string>] [-ConnectingTimeout <int>] [-SSHTransport] [<CommonParameters>]
Enter-PSSession [[-Session] <PSSession>] [<CommonParameters>]
Enter-PSSession [[-ConnectionUri] <uri>] [-EnableNetworkAccess] [-Credential <pscredential>] [-ConfigurationName <string>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Enter-PSSession [-InstanceId <guid>] [<CommonParameters>]
Enter-PSSession [[-Id] <int>] [<CommonParameters>]
Enter-PSSession [-Name <string>] [<CommonParameters>]
Enter-PSSession [-VMId] <guid> [-Credential] <pscredential> [-ConfigurationName <string>] [<CommonParameters>]
Enter-PSSession [-VMName] <string> [-Credential] <pscredential> [-ConfigurationName <string>] [<CommonParameters>]
Enter-PSSession [-ContainerId] <string> [-ConfigurationName <string>] [-RunAsAdministrator] [<CommonParameters>]
```

Example (5.1): Start an interactive session

```powershell
PS C:\> Enter-PSSession
[localhost]: PS C:\>
```

Example (7): Start an interactive session

```powershell
PS> Enter-PSSession
[localhost]: PS>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Enter-PSSession.md)


### Exit-PSHostProcess

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Exit-PSHostProcess [<CommonParameters>]
```

Example (5.1): Exit a process

```powershell
PS C:\> [Process:1520]: PS C:\>  Exit-PSHostProcess
PS C:\>
```

Example (7): Exit a process

```powershell
[Process:1520]: PS>  Exit-PSHostProcess
PS>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Exit-PSHostProcess.md)


### Exit-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Exit-PSSession [<CommonParameters>]
```

Example (5.1): Start and stop an interactive session

```powershell
PS C:\> Enter-PSSession -ComputerName Server01
Server01\PS> Exit-PSSession
PS C:\>
```

Example (7): Start and stop an interactive session

```powershell
PS> Enter-PSSession -ComputerName Server01
Server01\PS> Exit-PSSession
PS>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Exit-PSSession.md)


### Export-Alias

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Export-Alias [-Path] <string> [[-Name] <string[]>] [-PassThru] [-As <ExportAliasFormat>] [-Append] [-Force] [-NoClobber] [-Description <string>] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-Alias [[-Name] <string[]>] -LiteralPath <string> [-PassThru] [-As <ExportAliasFormat>] [-Append] [-Force] [-NoClobber] [-Description <string>] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export an alias

```powershell
Export-Alias -Path Alias.csv
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Export-Alias.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Alias, New-Alias, Remove-Alias, Import-Alias, Export-Alias.
- Function: manages aliases.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Alias file path |

- Set-Alias overwrites; New-Alias errors on an existing alias without -Force; Remove-Alias deletes; Export-Alias writes one `name=target` per line; Import-Alias reads them back.


### Export-Clixml

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Export-Clixml [-Path] <string> -InputObject <psobject> [-Depth <int>] [-Force] [-NoClobber] [-Encoding <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-Clixml -LiteralPath <string> -InputObject <psobject> [-Depth <int>] [-Force] [-NoClobber] [-Encoding <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Export-Clixml [-Path] <string> -InputObject <psobject> [-Depth <int>] [-Force] [-NoClobber] [-Encoding <Encoding>] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-Clixml -LiteralPath <string> -InputObject <psobject> [-Depth <int>] [-Force] [-NoClobber] [-Encoding <Encoding>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export a string to an XML file

```powershell
"This is a test" | Export-Clixml -Path .\sample.xml
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Export-Clixml.md)


### Export-Csv

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Export-Csv [[-Path] <string>] [[-Delimiter] <char>] -InputObject <psobject> [-LiteralPath <string>] [-Force] [-NoClobber] [-Encoding <string>] [-Append] [-NoTypeInformation] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-Csv [[-Path] <string>] -InputObject <psobject> [-LiteralPath <string>] [-Force] [-NoClobber] [-Encoding <string>] [-Append] [-UseCulture] [-NoTypeInformation] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Export-Csv [[-Path] <string>] [[-Delimiter] <char>] -InputObject <psobject> [-LiteralPath <string>] [-Force] [-NoClobber] [-Encoding <Encoding>] [-Append] [-IncludeTypeInformation] [-NoTypeInformation] [-QuoteFields <string[]>] [-UseQuotes <QuoteKind>] [-NoHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-Csv [[-Path] <string>] -InputObject <psobject> [-LiteralPath <string>] [-Force] [-NoClobber] [-Encoding <Encoding>] [-Append] [-UseCulture] [-IncludeTypeInformation] [-NoTypeInformation] [-QuoteFields <string[]>] [-UseQuotes <QuoteKind>] [-NoHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Export process properties to a CSV file

```powershell
Get-Process -Name WmiPrvSE |
    Select-Object -Property BasePriority,Id,SessionId,WorkingSet |
    Export-Csv -Path .\WmiData.csv -NoTypeInformation
Import-Csv -Path .\WmiData.csv
```

Example (7): Export process properties to a CSV file

```powershell
Get-Process -Name WmiPrvSE |
    Select-Object -Property BasePriority, Id, SessionId, WorkingSet |
    Export-Csv -Path .\WmiData.csv -NoTypeInformation
Import-Csv -Path .\WmiData.csv
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Export-Csv.md)


### Export-FormatData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Export-FormatData -InputObject <ExtendedTypeDefinition[]> -Path <string> [-Force] [-NoClobber] [-IncludeScriptBlock] [<CommonParameters>]
Export-FormatData -InputObject <ExtendedTypeDefinition[]> -LiteralPath <string> [-Force] [-NoClobber] [-IncludeScriptBlock] [<CommonParameters>]
```

Example: Export session format data

```powershell
Get-FormatData -TypeName "*" |
    Export-FormatData -Path "AllFormat.ps1xml" -IncludeScriptBlock
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Export-FormatData.md)


### Export-ModuleMember

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Export-ModuleMember [[-Function] <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [<CommonParameters>]
```

Example: Export functions and aliases in a script module

```powershell
Export-ModuleMember -Function * -Alias *
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Export-ModuleMember.md)


### Export-PSSession

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Export-PSSession [-Session] <PSSession> [-OutputModule] <string> [[-CommandName] <string[]>] [[-FormatTypeName] <string[]>] [-Force] [-Encoding <string>] [-AllowClobber] [-ArgumentList <Object[]>] [-CommandType <CommandTypes>] [-Module <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Certificate <X509Certificate2>] [<CommonParameters>]
```

Syntax (7):

```powershell
Export-PSSession [-Session] <PSSession> [-OutputModule] <string> [[-CommandName] <string[]>] [[-FormatTypeName] <string[]>] [-Force] [-Encoding <Encoding>] [-AllowClobber] [-ArgumentList <Object[]>] [-CommandType <CommandTypes>] [-Module <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Certificate <X509Certificate2>] [<CommonParameters>]
```

Example: Export commands from a PSSession

```powershell
$S = New-PSSession -ComputerName Server01
Export-PSSession -Session $S -OutputModule Server01
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Export-PSSession.md)


### Find-Package

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Find-Package [[-Name] <string[]>] [-IncludeDependencies] [-AllVersions] [-Source <string[]>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Find-Package [[-Name] <string[]>] [-IncludeDependencies] [-AllVersions] [-Source <string[]>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-ConfigFile <string>] [-SkipValidate] [-Headers <string[]>] [-FilterOnTag <string[]>] [-Contains <string>] [-AllowPrereleaseVersions] [<CommonParameters>]
Find-Package [[-Name] <string[]>] [-IncludeDependencies] [-AllVersions] [-Source <string[]>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-AllowPrereleaseVersions] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-AcceptLicense] [<CommonParameters>]
```

Example: Find all available packages from a package provider

```powershell
Find-Package -ProviderName NuGet
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/find-package?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/find-package?view=powershell-7.5)


### Find-PackageProvider

Version: Both

Module: PackageManagement

Syntax:

```powershell
Find-PackageProvider [[-Name] <string[]>] [-AllVersions] [-Source <string[]>] [-IncludeDependencies] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Force] [-ForceBootstrap] [<CommonParameters>]
```

Example: Find all available package providers

```powershell
Find-PackageProvider
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/find-packageprovider?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/find-packageprovider?view=powershell-7.5)


### Find-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Find-PSResource [[-Name] <string[]>] [-Type <ResourceType>] [-Version <string>] [-Prerelease] [-Tag <string[]>] [-Repository <string[]>] [-Credential <pscredential>] [-IncludeDependencies] [<CommonParameters>]
Find-PSResource -CommandName <string[]> [-Prerelease] [-Repository <string[]>] [-Credential <pscredential>] [<CommonParameters>]
Find-PSResource -DscResourceName <string[]> [-Prerelease] [-Repository <string[]>] [-Credential <pscredential>] [<CommonParameters>]
```

Example: 

```powershell
Find-PSResource -Name PowerShellGet -Repository PSGallery
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/find-psresource?view=powershell-7.5)


### ForEach-Object

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
ForEach-Object [-Process] <scriptblock[]> [-InputObject <psobject>] [-Begin <scriptblock>] [-End <scriptblock>] [-RemainingScripts <scriptblock[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
ForEach-Object [-MemberName] <string> [-InputObject <psobject>] [-ArgumentList <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
ForEach-Object [-Process] <scriptblock[]> [-InputObject <psobject>] [-Begin <scriptblock>] [-End <scriptblock>] [-RemainingScripts <scriptblock[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
ForEach-Object [-MemberName] <string> [-InputObject <psobject>] [-ArgumentList <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
ForEach-Object -Parallel <scriptblock> [-InputObject <psobject>] [-ThrottleLimit <int>] [-TimeoutSeconds <int>] [-AsJob] [-UseNewRunspace] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Divide integers in an array

```powershell
30000, 56798, 12432 | ForEach-Object -Process {$_/1024}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/ForEach-Object.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: runs a script block per object. Bash's `xargs` / `for` loop.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Process` (position 0) | scriptblock | Block run against every input object, `$_` bound to it |
| `-MemberName` | string | Takes each object's value of that member (as in `-MemberName Length`) |
| `-Begin` / `-End` | scriptblock | Each runs once (aggregation style), `-Begin` before the loop, `-End` after |

- Implementation: for each input object, runs the block with `$_` bound to it and emits whatever objects the block produces.


### Format-Custom

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Format-Custom [[-Property] <Object[]>] [-Depth <int>] [-GroupBy <Object>] [-View <string>] [-ShowError] [-DisplayError] [-Force] [-Expand <string>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Format output with a custom view

```powershell
Get-Command Start-Transcript | Format-Custom -View MyView
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Format-Custom.md)


### Format-Hex

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Format-Hex [-Path] <string[]> [<CommonParameters>]
Format-Hex -LiteralPath <string[]> [<CommonParameters>]
Format-Hex -InputObject <Object> [-Encoding <string>] [-Raw] [<CommonParameters>]
```

Syntax (7):

```powershell
Format-Hex [-Path] <string[]> [-Count <Int64>] [-Offset <Int64>] [<CommonParameters>]
Format-Hex -LiteralPath <string[]> [-Count <Int64>] [-Offset <Int64>] [<CommonParameters>]
Format-Hex -InputObject <psobject> [-Encoding <Encoding>] [-Count <Int64>] [-Offset <Int64>] [-Raw] [<CommonParameters>]
```

Example: Get the hexadecimal representation of a string

```powershell
'Hello World' | Format-Hex
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Format-Hex.md) / [Official reference source (5.1)](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/Format-Hex.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Function: hexadecimal view. `xxd` / `od -x` counterparts.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File to view; emits a `Label: path` section |
| `-InputObject` | object | Piped-in objects, one section per object |

- Implementation: renders each input at 16 bytes per row — a 16-digit offset, a fixed-width byte area, an ASCII reference column on the right (unprintable characters shown as `.`), and a leading `Label:` line per section — the path for files, the type name plus content checksum for piped objects.


### Format-List

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Format-List [[-Property] <Object[]>] [-GroupBy <Object>] [-View <string>] [-ShowError] [-DisplayError] [-Force] [-Expand <string>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Format computer services

```powershell
Get-Service | Format-List
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Format-List.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: one property per line.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Properties to display; `*` means all |

- Implementation: renders `property : value` lines with names right-aligned to the widest; objects separated by a blank line.


### Format-Table

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Format-Table [[-Property] <Object[]>] [-AutoSize] [-RepeatHeader] [-HideTableHeaders] [-Wrap] [-GroupBy <Object>] [-View <string>] [-ShowError] [-DisplayError] [-Force] [-Expand <string>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Format PowerShell host

```powershell
Get-Host | Format-Table -AutoSize
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Format-Table.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Column-width fitting differs; no ANSI colors or directory headers; the time column uses a fixed format (PowerShell follows locale).

- Type: Go implementation.
- Function: table display. The alignment effect of `column`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Property columns to display |
| `-AutoSize` | switch | Auto-fit column widths (always auto-fits anyway; accepted) |

- Implementation: builds columns from the first object's properties (or its Props for objects without a table definition), right-aligns numeric columns, underlines headers with `----`. Scalars (strings/numbers/booleans/scriptblocks) have no tabulatable properties and print one per line.
- Custom-object streams merge into one table: consecutive same-shape objects share the first object's columns, later properties open no new columns, missing values stay blank. A lone object takes the list view.


### Format-Wide

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Format-Wide [[-Property] <Object>] [-AutoSize] [-Column <int>] [-GroupBy <Object>] [-View <string>] [-ShowError] [-DisplayError] [-Force] [-Expand <string>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Format names of files in the current directory

```powershell
Get-ChildItem | Format-Wide -Column 3
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Format-Wide.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: multi-column arrangement.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string | Which property of each object to display |
| `-Column` | int | Column width, default 40 |

- Implementation: lays object strings out across an 80-character width.


### Get-Alias

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Alias [[-Name] <string[]>] [-Exclude <string[]>] [-Scope <string>] [<CommonParameters>]
Get-Alias [-Exclude <string[]>] [-Scope <string>] [-Definition <string[]>] [<CommonParameters>]
```

Example: Get all aliases in the current session

```powershell
Get-Alias
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Alias.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Looking up a nonexistent alias returns empty with `$?` staying True (PowerShell errors with alias-not-found).

- Type: Go implementation.
- Function: lists aliases. Bash's `alias`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Alias names, wildcards/arrays supported |

- Output: AliasInfo objects with fields Name, Definition (target cmdlet). Table columns Name/Definition.
- Note: style 5's alias table contains sc/curl/wget, style 7 doesn't. Looking up a nonexistent alias returns empty with $? staying True (PowerShell raises an error).


### Get-ChildItem

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-ChildItem [[-Path] <string[]>] [[-Filter] <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Depth <uint32>] [-Force] [-Name] [-UseTransaction] [-Attributes <FlagsExpression[FileAttributes]>] [-FollowSymlink] [-Directory] [-File] [-Hidden] [-ReadOnly] [-System] [<CommonParameters>]
Get-ChildItem [[-Filter] <string>] -LiteralPath <string[]> [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Depth <uint>] [-Force] [-Name] [-UseTransaction] [-Attributes <FlagsExpression[FileAttributes]>] [-FollowSymlink] [-Directory] [-File] [-Hidden] [-ReadOnly] [-System] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-ChildItem [[-Path] <string[]>] [[-Filter] <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Depth <uint>] [-Force] [-Name] [-Attributes <FlagsExpression[FileAttributes]>] [-FollowSymlink] [-Directory] [-File] [-Hidden] [-ReadOnly] [-System] [<CommonParameters>]
Get-ChildItem [[-Filter] <string>] -LiteralPath <string[]> [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Depth <uint>] [-Force] [-Name] [-Attributes <FlagsExpression[FileAttributes]>] [-FollowSymlink] [-Directory] [-File] [-Hidden] [-ReadOnly] [-System] [<CommonParameters>]
```

Example: Get child items from a file system directory

```powershell
Get-ChildItem -Path C:\Test
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-ChildItem.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: `-Force` has no hidden-file effect.

- Type: Go implementation.
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


### Get-Clipboard

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-Clipboard [-Format <ClipboardFormat>] [-TextFormatType <TextDataFormat>] [-Raw] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Clipboard [-Raw] [<CommonParameters>]
```

Example: Get the content of the clipboard

```powershell
Set-Clipboard -Value 'hello world'
Get-Clipboard
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-Clipboard.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`xclip` / `xsel`).
- Companion: Set-Clipboard.
- Distro: X11 + xclip/xsel.
- Function: reads the clipboard.
- Parameters: none.
- Implementation: xclip first (`xclip -o -selection clipboard`), otherwise xsel (`xsel -b`); neither present, returns an empty string.


### Get-CmsMessage

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Get-CmsMessage [-Content] <string> [<CommonParameters>]
Get-CmsMessage [-Path] <string> [<CommonParameters>]
Get-CmsMessage [-LiteralPath] <string> [<CommonParameters>]
```

Example: Get encrypted content

```powershell
$Msg = Get-CmsMessage -Path "C:\Users\Test\Documents\PowerShell\Future_Plans.txt"
$Msg.Content
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Get-CmsMessage.md)


### Get-Command

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Get-Command [[-ArgumentList] <Object[]>] [-Verb <string[]>] [-Noun <string[]>] [-Module <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-TotalCount <int>] [-Syntax] [-ShowCommandInfo] [-All] [-ListImported] [-ParameterName <string[]>] [-ParameterType <PSTypeName[]>] [<CommonParameters>]
Get-Command [[-Name] <string[]>] [[-ArgumentList] <Object[]>] [-Module <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-CommandType <CommandTypes>] [-TotalCount <int>] [-Syntax] [-ShowCommandInfo] [-All] [-ListImported] [-ParameterName <string[]>] [-ParameterType <PSTypeName[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Command [[-ArgumentList] <Object[]>] [-Verb <string[]>] [-Noun <string[]>] [-Module <string[]>] [-ExcludeModule <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-TotalCount <int>] [-Syntax] [-ShowCommandInfo] [-All] [-ListImported] [-ParameterName <string[]>] [-ParameterType <PSTypeName[]>] [<CommonParameters>]
Get-Command [[-Name] <string[]>] [[-ArgumentList] <Object[]>] [-Module <string[]>] [-ExcludeModule <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-CommandType <CommandTypes>] [-TotalCount <int>] [-Syntax] [-ShowCommandInfo] [-All] [-ListImported] [-ParameterName <string[]>] [-ParameterType <PSTypeName[]>] [-UseFuzzyMatching] [-FuzzyMinimumDistance <uint>] [-UseAbbreviationExpansion] [<CommonParameters>]
```

Example: Get cmdlets, functions, and aliases

```powershell
Get-Command
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-Command.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: External commands are not listed, only built-ins and aliases.

- Type: Go implementation.
- Function: looks up commands. Bash's `type` / `which`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Command names, arrays supported |

- Output: CommandInfo objects (Name normalized-case, CommandType=Cmdlet/Function/Alias); Alias objects carry Definition besides. Table columns CommandType/Name.
- Difference from Windows: only built-in commands, functions, and aliases are listed, external commands aren't.


### Get-Content

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-Content [-Path] <string[]> [-ReadCount <long>] [-TotalCount <long>] [-Tail <int>] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-UseTransaction] [-Delimiter <string>] [-Wait] [-Raw] [-Encoding <FileSystemCmdletProviderEncoding>] [-Stream <string>] [<CommonParameters>]
Get-Content -LiteralPath <string[]> [-ReadCount <long>] [-TotalCount <long>] [-Tail <int>] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-UseTransaction] [-Delimiter <string>] [-Wait] [-Raw] [-Encoding <FileSystemCmdletProviderEncoding>] [-Stream <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Content [-Path] <string[]> [-ReadCount <long>] [-TotalCount <long>] [-Tail <int>] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-Delimiter <string>] [-Wait] [-Raw] [-Encoding <Encoding>] [-AsByteStream] [-Stream <string>] [<CommonParameters>]
Get-Content -LiteralPath <string[]> [-ReadCount <long>] [-TotalCount <long>] [-Tail <int>] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-Delimiter <string>] [-Wait] [-Raw] [-Encoding <Encoding>] [-AsByteStream] [-Stream <string>] [<CommonParameters>]
```

Example: Get the content of a text file

```powershell
1..100 | ForEach-Object {
    Add-Content -Path .\LineNumbers.txt -Value "This is line $_."
}
Get-Content -Path .\LineNumbers.txt
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-Content.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: reads text files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File paths; several allowed (`Get-Content a.txt,b.txt` reads them one by one, with -TotalCount/-Tail applied per file) |
| `-Raw` | switch | Reads whole without splitting into lines (`cat`) |
| `-TotalCount` | int | Only the first N lines (`head -n N`) |
| `-Tail` | int | Only the last N lines (`tail -n N`) |

- Output: one `String` object per line (line endings stripped). Empty files produce no output.
- Behavior: nonexistent path → error, $?=false, matching `cat`'s complaint about missing files.


### Get-Credential

Version: Both

Module: Microsoft.PowerShell.Security

Syntax (5.1):

```powershell
Get-Credential [-Credential] <pscredential> [<CommonParameters>]
Get-Credential [[-UserName] <string>] -Message <string> [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Credential [[-Credential] <pscredential>] [<CommonParameters>]
Get-Credential [[-UserName] <string>] [-Message <string>] [-Title <string>] [<CommonParameters>]
```

Example: 

```powershell
$c = Get-Credential
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Get-Credential.md)


### Get-Culture

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Get-Culture
```

Syntax (7):

```powershell
Get-Culture [-NoUserOverrides] [<CommonParameters>]
Get-Culture [-Name <string[]>] [-NoUserOverrides] [<CommonParameters>]
Get-Culture [-ListAvailable] [<CommonParameters>]
```

Example: Get culture settings

```powershell
Get-Culture
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Culture.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Culture follows the UI language, falling back to zh-CN when the UI language has no registered culture.

- Type: Go implementation.
- Function: shows the locale. `locale` counterpart.
- Parameters: none.
- Implementation: returns a CultureInfo from the interface language's registry table, falling back to zh-CN when the language has no registered locale (consistent with `$Host.CurrentCulture`).
- Output: a CultureInfo object with fields LCID, Name, DisplayName.


### Get-Date

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Get-Date [[-Date] <datetime>] [-Year <int>] [-Month <int>] [-Day <int>] [-Hour <int>] [-Minute <int>] [-Second <int>] [-Millisecond <int>] [-DisplayHint <DisplayHintType>] [-Format <string>] [<CommonParameters>]
Get-Date [[-Date] <datetime>] [-Year <int>] [-Month <int>] [-Day <int>] [-Hour <int>] [-Minute <int>] [-Second <int>] [-Millisecond <int>] [-DisplayHint <DisplayHintType>] [-UFormat <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Date [[-Date] <datetime>] [-Year <int>] [-Month <int>] [-Day <int>] [-Hour <int>] [-Minute <int>] [-Second <int>] [-Millisecond <int>] [-DisplayHint <DisplayHintType>] [-Format <string>] [-AsUTC] [<CommonParameters>]
Get-Date [[-Date] <datetime>] -UFormat <string> [-Year <int>] [-Month <int>] [-Day <int>] [-Hour <int>] [-Minute <int>] [-Second <int>] [-Millisecond <int>] [-DisplayHint <DisplayHintType>] [<CommonParameters>]
Get-Date -UnixTimeSeconds <long> [-Year <int>] [-Month <int>] [-Day <int>] [-Hour <int>] [-Minute <int>] [-Second <int>] [-Millisecond <int>] [-DisplayHint <DisplayHintType>] [-Format <string>] [-AsUTC] [<CommonParameters>]
Get-Date -UnixTimeSeconds <long> -UFormat <string> [-Year <int>] [-Month <int>] [-Day <int>] [-Hour <int>] [-Minute <int>] [-Second <int>] [-Millisecond <int>] [-DisplayHint <DisplayHintType>] [<CommonParameters>]
```

Example: Get the current date and time

```powershell
Get-Date
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Date.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Function: current time or a specified date-time. Bash's `date`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Date` | string | Specified date-time (parsed against the local time zone when none is written), accepting `2006-01-02 15:04:05`, `2006-01-02T15:04:05`, `2006-01-02`, `2006/1/2`, RFC3339 |
| `-Format` | string | .NET format string; conversion rules below |

- `-Format` conversion rules: `yyyy`→`2006`, `yy`→`06`; `M`→`1`, `MM`→`01`, `MMM`→`Jan`, `MMMM`→`January`; `d`→`2`, `dd`→`02`, `ddd`→`Mon`, `dddd`→`Monday`; `H`→`15`, `hh`→`03`, `m`→`4`, `mm`→`04`, `s`→`5`, `ss`→`05`, `tt`→`PM`, `zzz`→`-07:00`.
- Output: without -Format, a DateTime object (current time when -Date is absent); with -Format, a String. Without -Format the rendering follows the interface language — registered languages use their own formats (English looks like `Saturday, 15 August 2026 15:28:08`), unregistered languages fall back to the default language's Chinese format looking like `2026年8月15日星期六 15:28:08`.


### Get-Error

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Error [[-Newest] <int>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Get the most recent error details

```powershell
Get-ChildItem -Path /NoRealDirectory
Get-Error
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Error.md)


### Get-Event

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Event [[-SourceIdentifier] <string>] [<CommonParameters>]
Get-Event [-EventIdentifier] <int> [<CommonParameters>]
```

Example: Get all events

```powershell
PS C:\> Get-Event
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Event.md)


### Get-EventSubscriber

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-EventSubscriber [[-SourceIdentifier] <string>] [-Force] [<CommonParameters>]
Get-EventSubscriber [-SubscriptionId] <int> [-Force] [<CommonParameters>]
```

Example: Get the event subscriber for a timer event

```powershell
$Timer = New-Object Timers.Timer
$Timer | Get-Member -Type Event
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-EventSubscriber.md)


### Get-ExecutionPolicy

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Get-ExecutionPolicy [[-Scope] <ExecutionPolicyScope>] [-List] [<CommonParameters>]
```

Example: Get all execution policies

```powershell
Get-ExecutionPolicy -List
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Get-ExecutionPolicy.md)


### Get-ExperimentalFeature

Version: 7 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-ExperimentalFeature [[-Name] <string[]>] [<CommonParameters>]
```

Example: 

```powershell
Get-ExperimentalFeature
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-ExperimentalFeature.md)


### Get-FileHash

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Get-FileHash [-Path] <String[]> [-Algorithm <String>] [<CommonParameters>]
Get-FileHash -LiteralPath <String[]> [-Algorithm <String>] [<CommonParameters>]
Get-FileHash -InputStream <Stream> [-Algorithm <String>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-FileHash [-Path] <string[]> [[-Algorithm] <string>] [<CommonParameters>]
Get-FileHash [-LiteralPath] <string[]> [[-Algorithm] <string>] [<CommonParameters>]
Get-FileHash [-InputStream] <Stream> [[-Algorithm] <string>] [<CommonParameters>]
```

Example (5.1): Compute the hash value for a file

```powershell
Get-FileHash $PSHOME\powershell.exe | Format-List
```

Example (7): Compute the hash value for a file

```powershell
Get-FileHash /etc/apt/sources.list | Format-List
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-FileHash.md) / [Official reference source (5.1)](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/Get-FileHash.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: SHA384, MACTripleDES and RIPEMD160 are not supported.

- Type: Go implementation.
- Distro: any (computed by the program itself, no external commands involved).
- Function: computes file hashes.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File path, wildcards supported |
| `-Algorithm` | string | Algorithm: SHA256 (default) / SHA1 / MD5 / SHA512 (SHA2_256/SHA2_512 also recognized) |

- Implementation: pure Go (crypto standard library). Counterparts: `sha256sum` / `md5sum` / `sha1sum` / `sha512sum`.
- Output: a FileHash object with fields Algorithm (uppercase), Hash (lowercase hex), Path.


### Get-FormatData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-FormatData [[-TypeName] <string[]>] [-PowerShellVersion <version>] [<CommonParameters>]
```

Example: Get all formatting data

```powershell
Get-FormatData
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-FormatData.md)


### Get-Help

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Get-Help [[-Name] <string>] [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [-Full] [<CommonParameters>]
Get-Help [[-Name] <string>] -Detailed [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Examples [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Parameter <string> [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Online [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -ShowWindow [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Help [[-Name] <string>] [-Path <string>] [-Category <string[]>] [-Full] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Detailed [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Examples [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Parameter <string[]> [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -Online [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
Get-Help [[-Name] <string>] -ShowWindow [-Path <string>] [-Category <string[]>] [-Component <string[]>] [-Functionality <string[]>] [-Role <string[]>] [<CommonParameters>]
```

Example: Display basic help information about a cmdlet

```powershell
Get-Help Format-Table
Get-Help -Name Format-Table
Format-Table -?
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-Help.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Shows only names, syntax and aliases without detailed descriptions.

- Type: Go implementation.
- Function: shows command syntax. Bash's `man`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Command name, wildcards supported |

- Output: name, syntax (`command [-param <type>]...`), aliases. No match found checks whether it's an alias and shows its target; failing that, reports not found.
- Difference from Windows: only name, syntax, and aliases are shown — no detailed explanation.


### Get-History

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-History [[-Id] <long[]>] [[-Count] <int>] [<CommonParameters>]
```

Example: Get the session history

```powershell
Get-History
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-History.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Companions: Clear-History, Add-History, Invoke-History.
- Function: lists history. Bash's `history`.
- Parameters: none.
- Output: HistoryInfo objects with fields Id, CommandLine. Table columns Id/CommandLine.


### Get-Host

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Host
```

Example: Get information about the PowerShell console host

```powershell
Get-Host
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Host.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Function: returns the host object.
- Parameters: none.
- Output: a `System.Management.Automation.Internal.Host.InternalHost` object with fields:

| Field | Value |
| :--- | :--- |
| `Name` | "ConsoleHost" |
| `InstanceId` | A random UUID for this session (changes each start) |
| `UI` | Host UI object containing only `SupportsVirtualTerminal` (always True) |
| `CurrentCulture` | CultureInfo object (LCID/Name/DisplayName), returned from the registry table of the interface language, zh-CN when that language has no registered locale |
| `CurrentUICulture` | Same as CurrentCulture |


### Get-InstalledPSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Get-InstalledPSResource [[-Name] <string[]>] [-Version <string>] [-Path <string>] [-Scope <ScopeType>] [<CommonParameters>]
```

Example: 

```powershell
Get-InstalledPSResource
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/get-installedpsresource?view=powershell-7.5)


### Get-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-Item [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-UseTransaction] [-Stream <string[]>] [<CommonParameters>]
Get-Item -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-UseTransaction] [-Stream <string[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Item [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-Stream <string[]>] [<CommonParameters>]
Get-Item -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-Stream <string[]>] [<CommonParameters>]
```

Example: Get the current directory

```powershell
Get-Item .
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-Item.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: inspects a single file/directory. Bash's `stat`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path, wildcards supported; several allowed (`Get-Item a.txt,b.txt`) |

- Output: a single FileInfo/DirectoryInfo object, fields same as Get-ChildItem.
- Behavior: nonexistent path → error, $?=false.


### Get-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-ItemProperty [-Path] <string[]> [[-Name] <string[]>] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Get-ItemProperty [[-Name] <string[]>] -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-ItemProperty [-Path] <string[]> [[-Name] <string[]>] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [<CommonParameters>]
Get-ItemProperty [[-Name] <string[]>] -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Get information about a specific directory

```powershell
Get-ItemProperty C:\Windows
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-ItemProperty.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Only 5 fields are output.

- Type: Go implementation.
- Function: emits an object of the path's properties. The full `stat` picture in bash terms.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |
| `-Name` (position 1) | string | Keep only that property (omit to get all five), e.g. `Get-ItemProperty x -Name Length` |

- Output: a `System.Management.Automation.PSCustomObject` with fields Name / FullName / Length / LastWriteTime / Mode.


### Get-ItemPropertyValue

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-ItemPropertyValue [[-Path] <string[]>] [-Name] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Get-ItemPropertyValue [-Name] <string[]> -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-ItemPropertyValue [[-Path] <string[]>] [-Name] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [<CommonParameters>]
Get-ItemPropertyValue [-Name] <string[]> -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Get the value of the ProductID property

```powershell
Get-ItemPropertyValue 'HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion' -Name ProductID
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-ItemPropertyValue.md)


### Get-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-Job [[-Id] <int[]>] [-IncludeChildJob] [-ChildJobState <JobState>] [-HasMoreData <bool>] [-Before <datetime>] [-After <datetime>] [-Newest <int>] [<CommonParameters>]
Get-Job [-InstanceId] <guid[]> [-IncludeChildJob] [-ChildJobState <JobState>] [-HasMoreData <bool>] [-Before <datetime>] [-After <datetime>] [-Newest <int>] [<CommonParameters>]
Get-Job [-Name] <string[]> [-IncludeChildJob] [-ChildJobState <JobState>] [-HasMoreData <bool>] [-Before <datetime>] [-After <datetime>] [-Newest <int>] [<CommonParameters>]
Get-Job [-State] <JobState> [-IncludeChildJob] [-ChildJobState <JobState>] [-HasMoreData <bool>] [-Before <datetime>] [-After <datetime>] [-Newest <int>] [<CommonParameters>]
Get-Job [-IncludeChildJob] [-ChildJobState <JobState>] [-HasMoreData <bool>] [-Before <datetime>] [-After <datetime>] [-Newest <int>] [-Command <string[]>] [<CommonParameters>]
Get-Job [-Filter] <hashtable> [<CommonParameters>]
```

Example: Get all background jobs started in the current session

```powershell
Get-Job
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-Job.md)


### Get-Location

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-Location [-PSProvider <string[]>] [-PSDrive <string[]>] [-UseTransaction] [<CommonParameters>]
Get-Location [-Stack] [-StackName <string[]>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Location [-PSProvider <string[]>] [-PSDrive <string[]>] [<CommonParameters>]
Get-Location [-Stack] [-StackName <string[]>] [<CommonParameters>]
```

Example: Display your current drive location

```powershell
PS C:\Windows> Get-Location
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-Location.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Outputs path strings directly (PowerShell renders a Path table of PathInfo objects).

- Type: Go implementation.
- Function: shows the current directory.
- Parameters: none.
- Output: a `String` — the Windows-style display of the current directory (root `C:\`), bash's `pwd`. Difference from PowerShell: prints the path string directly, not wrapped in a PathInfo table.


### Get-MarkdownOption

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-MarkdownOption [<CommonParameters>]
```

Example: Get the current colors and style

```powershell
Get-MarkdownOption
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-MarkdownOption.md)


### Get-Member

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Member [[-Name] <string[]>] [-InputObject <psobject>] [-MemberType <PSMemberTypes>] [-View <PSMemberViewTypes>] [-Static] [-Force] [<CommonParameters>]
```

Example: Get the members of process objects

```powershell
Get-Service | Get-Member
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Member.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Method members are not listed, only types and properties.

- Type: Go implementation.
- Function: inspects object members. Python's `dir()` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to inspect |
| `-MemberType` | string | Member-type filter (Property / TypeName etc., case-insensitive, all returned by default) |

- Output: PSMemberInfo objects with fields Name, MemberType (TypeName/Property), Definition (the property value as string). Each type lists its type row once, properties deduplicated.


### Get-Module

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Get-Module [[-Name] <string[]>] [-FullyQualifiedName <ModuleSpecification[]>] [-All] [<CommonParameters>]
Get-Module [[-Name] <string[]>] -PSSession <PSSession> [-FullyQualifiedName <ModuleSpecification[]>] [-ListAvailable] [-PSEdition <string>] [-Refresh] [<CommonParameters>]
Get-Module [[-Name] <string[]>] -ListAvailable [-FullyQualifiedName <ModuleSpecification[]>] [-All] [-PSEdition <string>] [-Refresh] [<CommonParameters>]
Get-Module [[-Name] <string[]>] -CimSession <CimSession> [-FullyQualifiedName <ModuleSpecification[]>] [-ListAvailable] [-Refresh] [-CimResourceUri <uri>] [-CimNamespace <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Module [[-Name] <string[]>] [-FullyQualifiedName <ModuleSpecification[]>] [-All] [<CommonParameters>]
Get-Module [[-Name] <string[]>] -ListAvailable [-FullyQualifiedName <ModuleSpecification[]>] [-All] [-PSEdition <string>] [-SkipEditionCheck] [-Refresh] [<CommonParameters>]
Get-Module [[-Name] <string[]>] -PSSession <PSSession> [-FullyQualifiedName <ModuleSpecification[]>] [-ListAvailable] [-PSEdition <string>] [-SkipEditionCheck] [-Refresh] [<CommonParameters>]
Get-Module [[-Name] <string[]>] -CimSession <CimSession> [-FullyQualifiedName <ModuleSpecification[]>] [-ListAvailable] [-SkipEditionCheck] [-Refresh] [-CimResourceUri <uri>] [-CimNamespace <string>] [<CommonParameters>]
```

Example: Get modules imported into the current session

```powershell
Get-Module
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-Module.md)


### Get-Package

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Get-Package [[-Name] <string[]>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-AllVersions] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Get-Package [[-Name] <string[]>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-AllVersions] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-Scope <string>] [-PackageManagementProvider <string>] [-Type <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [-AllowPrereleaseVersions] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Package [[-Name] <string[]>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-AllVersions] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Get-Package [[-Name] <string[]>] [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-AllVersions] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-Scope <string>] [-PackageManagementProvider <string>] [-Type <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [-AllowPrereleaseVersions] [<CommonParameters>]
```

Example: Get all installed packages

```powershell
Get-Package
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/get-package?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/get-package?view=powershell-7.5)


### Get-PackageProvider

Version: Both

Module: PackageManagement

Syntax:

```powershell
Get-PackageProvider [[-Name] <string[]>] [-ListAvailable] [-Force] [-ForceBootstrap] [<CommonParameters>]
```

Example: Get all currently loaded package providers

```powershell
Get-PackageProvider
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/get-packageprovider?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/get-packageprovider?view=powershell-7.5)


### Get-PackageSource

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Get-PackageSource [[-Name] <string>] [-Location <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Get-PackageSource [[-Name] <string>] [-Location <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-PackageSource [[-Name] <string>] [-Location <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Get-PackageSource [[-Name] <string>] [-Location <string>] [-Force] [-ForceBootstrap] [-ProviderName <string[]>] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Example: Get all package sources

```powershell
Get-PackageSource
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/get-packagesource?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/get-packagesource?view=powershell-7.5)


### Get-PfxCertificate

Version: Both

Module: Microsoft.PowerShell.Security

Syntax (5.1):

```powershell
Get-PfxCertificate [-FilePath] <string[]> [<CommonParameters>]
Get-PfxCertificate -LiteralPath <string[]> [<CommonParameters>]
```

Syntax (7):

```powershell
Get-PfxCertificate [-FilePath] <string[]> [-Password <securestring>] [-NoPromptForPassword] [<CommonParameters>]
Get-PfxCertificate -LiteralPath <string[]> [-Password <securestring>] [-NoPromptForPassword] [<CommonParameters>]
```

Example: Get a PFX certificate

```powershell
Get-PfxCertificate -FilePath "C:\windows\system32\Test.pfx"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Get-PfxCertificate.md)


### Get-Process

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-Process [[-Name] <string[]>] [-ComputerName <string[]>] [-Module] [-FileVersionInfo] [<CommonParameters>]
Get-Process [[-Name] <string[]>] -IncludeUserName [<CommonParameters>]
Get-Process -Id <int[]> [-ComputerName <string[]>] [-Module] [-FileVersionInfo] [<CommonParameters>]
Get-Process -Id <int[]> -IncludeUserName [<CommonParameters>]
Get-Process -InputObject <Process[]> -IncludeUserName [<CommonParameters>]
Get-Process -InputObject <Process[]> [-ComputerName <string[]>] [-Module] [-FileVersionInfo] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Process [[-Name] <string[]>] [-Module] [-FileVersionInfo] [<CommonParameters>]
Get-Process [[-Name] <string[]>] -IncludeUserName [<CommonParameters>]
Get-Process -Id <int[]> [-Module] [-FileVersionInfo] [<CommonParameters>]
Get-Process -Id <int[]> -IncludeUserName [<CommonParameters>]
Get-Process -InputObject <Process[]> [-Module] [-FileVersionInfo] [<CommonParameters>]
Get-Process -InputObject <Process[]> -IncludeUserName [<CommonParameters>]
```

Example: Get a list of all running processes on the local computer

```powershell
Get-Process
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-Process.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: The memory field is named Memory (WS in PowerShell) holding physical memory bytes, matching the semantics of PowerShell's WS; `-Name` matches by substring (PowerShell uses exact names or wildcards).

- Type: Go implementation.
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


### Get-PSBreakpoint

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Get-PSBreakpoint [-Script] <string[]> [<CommonParameters>]
Get-PSBreakpoint -Variable <string[]> [-Script <string[]>] [<CommonParameters>]
Get-PSBreakpoint -Command <string[]> [-Script <string[]>] [<CommonParameters>]
Get-PSBreakpoint [-Type] <BreakpointType[]> [-Script <string[]>] [<CommonParameters>]
Get-PSBreakpoint [-Id] <int[]> [<CommonParameters>]
```

Syntax (7):

```powershell
Get-PSBreakpoint [[-Script] <string[]>] [-Runspace <runspace>] [<CommonParameters>]
Get-PSBreakpoint -Command <string[]> [-Script <string[]>] [-Runspace <runspace>] [<CommonParameters>]
Get-PSBreakpoint -Variable <string[]> [-Script <string[]>] [-Runspace <runspace>] [<CommonParameters>]
Get-PSBreakpoint [-Type] <BreakpointType[]> [-Script <string[]>] [-Runspace <runspace>] [<CommonParameters>]
Get-PSBreakpoint [-Id] <int[]> [-Runspace <runspace>] [<CommonParameters>]
```

Example: Get all breakpoints for all scripts and functions

```powershell
Get-PSBreakpoint
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-PSBreakpoint.md)


### Get-PSCallStack

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-PSCallStack [<CommonParameters>]
```

Example: Get the call stack for a function

```powershell
PS C:\> function My-Alias {
$p = $args[0]
Get-Alias | where {$_.Definition -like "*$p"} | Format-Table Definition, Name -Auto
}
PS C:\ps-test> Set-PSBreakpoint -Command My-Alias
Command    : My-Alias
Action     :
Enabled    : True
HitCount   : 0
Id         : 0
Script     : prompt PS C:\> My-Alias Get-Content

Entering debug mode. Use h or ? for help.
Hit Command breakpoint on 'prompt:My-Alias'
My-Alias Get-Content
[DBG]: PS C:\ps-test> s
$p = $args[0]
DEBUG: Stepped to ':    $p = $args[0]    '
[DBG]: PS C:\ps-test> s
Get-Alias | where {$_.Definition -like "*$p*"} | Format-Table Definition,
[DBG]: PS C:\ps-test>Get-PSCallStack

Name        CommandLineParameters         UnboundArguments              Location
----        ---------------------         ----------------              --------
prompt      {}                            {}                            prompt
My-Alias    {}                            {Get-Content}                 prompt
prompt      {}                            {}                            prompt

PS C:\> [DBG]: PS C:\ps-test> o
Definition  Name
----------  ----
Get-Content gc
Get-Content cat
Get-Content type
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-PSCallStack.md)


### Get-PSDrive

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-PSDrive [[-Name] <string[]>] [-Scope <string>] [-PSProvider <string[]>] [-UseTransaction] [<CommonParameters>]
Get-PSDrive [-LiteralName] <string[]> [-Scope <string>] [-PSProvider <string[]>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-PSDrive [[-Name] <string[]>] [-Scope <string>] [-PSProvider <string[]>] [<CommonParameters>]
Get-PSDrive [-LiteralName] <string[]> [-Scope <string>] [-PSProvider <string[]>] [<CommonParameters>]
```

Example: Get drives in the current session

```powershell
PS C:\> Get-PSDrive

Name           Used (GB)     Free (GB) Provider      Root
----           ---------     --------- --------      ----
Alias                                  Alias
C                 202.06      23718.91 FileSystem    C:\
Cert                                   Certificate   \
D                1211.06     123642.32 FileSystem    D:\
Env                                    Environment
Function                               Function
HKCU                                   Registry      HKEY_CURRENT_USER
HKLM                                   Registry      HKEY_LOCAL_MACHINE
Variable                               Variable
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-PSDrive.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Only the root drive and Env exist.

- Type: Go implementation.
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


### Get-PSHostProcessInfo

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-PSHostProcessInfo [[-Name] <string[]>] [<CommonParameters>]
Get-PSHostProcessInfo [-Process] <Process[]> [<CommonParameters>]
Get-PSHostProcessInfo [-Id] <int[]> [<CommonParameters>]
```

Example: Get a list of PowerShell hosts running on the system

```powershell
Get-PSHostProcessInfo
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-PSHostProcessInfo.md)


### Get-PSProvider

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-PSProvider [[-PSProvider] <string[]>] [<CommonParameters>]
```

Example: Display a list of all available providers

```powershell
Get-PSProvider
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-PSProvider.md)


### Get-PSReadLineKeyHandler

Version: Both

Module: PSReadLine

Syntax (5.1):

```powershell
Get-PSReadLineKeyHandler [-Bound] [-Unbound] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-PSReadLineKeyHandler [-Bound] [-Unbound] [<CommonParameters>]
Get-PSReadLineKeyHandler [-Chord] <string[]> [<CommonParameters>]
```

Example: Get all key mappings

```powershell
Get-PSReadLineKeyHandler -Bound -Unbound
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/PSReadLine/Get-PSReadLineKeyHandler.md)


### Get-PSReadLineOption

Version: Both

Module: PSReadLine

Syntax:

```powershell
Get-PSReadLineOption [<CommonParameters>]
```

Example: Get options and their values

```powershell
Get-PSReadLineOption
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/PSReadLine/Get-PSReadLineOption.md)


### Get-PSResourceRepository

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Get-PSResourceRepository [[-Name] <string[]>] [<CommonParameters>]
```

Example: 

```powershell
Get-PSResourceRepository
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/get-psresourcerepository?view=powershell-7.5)


### Get-PSScriptFileInfo

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Get-PSScriptFileInfo [-Path] <string> [<CommonParameters>]
```

Example: 

```powershell
Get-PSScriptFileInfo -Path '.\Scripts\MyScript.ps1'
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/get-psscriptfileinfo?view=powershell-7.5)


### Get-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Get-PSSession [-Name <string[]>] [<CommonParameters>]
Get-PSSession [-ComputerName] <string[]> -InstanceId <guid[]> [-ApplicationName <string>] [-ConfigurationName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession [-ComputerName] <string[]> [-ApplicationName <string>] [-ConfigurationName <string>] [-Name <string[]>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession [-ConnectionUri] <uri[]> -InstanceId <guid[]> [-ConfigurationName <string>] [-AllowRedirection] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession [-ConnectionUri] <uri[]> [-ConfigurationName <string>] [-AllowRedirection] [-Name <string[]>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession -ContainerId <string[]> [-ConfigurationName <string>] [-Name <string[]>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -InstanceId <guid[]> -ContainerId <string[]> [-ConfigurationName <string>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -VMId <guid[]> [-ConfigurationName <string>] [-Name <string[]>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -InstanceId <guid[]> -VMId <guid[]> [-ConfigurationName <string>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -VMName <string[]> [-ConfigurationName <string>] [-Name <string[]>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -InstanceId <guid[]> -VMName <string[]> [-ConfigurationName <string>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession [-InstanceId <guid[]>] [<CommonParameters>]
Get-PSSession [-Id] <int[]> [<CommonParameters>]
```

Syntax (7):

```powershell
Get-PSSession [-Name <string[]>] [<CommonParameters>]
Get-PSSession [-ComputerName] <string[]> [-ApplicationName <string>] [-ConfigurationName <string>] [-Name <string[]>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession [-ComputerName] <string[]> -InstanceId <guid[]> [-ApplicationName <string>] [-ConfigurationName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession [-ConnectionUri] <uri[]> [-ConfigurationName <string>] [-AllowRedirection] [-Name <string[]>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession [-ConnectionUri] <uri[]> -InstanceId <guid[]> [-ConfigurationName <string>] [-AllowRedirection] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-ThrottleLimit <int>] [-State <SessionFilterState>] [-SessionOption <PSSessionOption>] [<CommonParameters>]
Get-PSSession -ContainerId <string[]> [-ConfigurationName <string>] [-Name <string[]>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -InstanceId <guid[]> -ContainerId <string[]> [-ConfigurationName <string>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -VMId <guid[]> [-ConfigurationName <string>] [-Name <string[]>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -InstanceId <guid[]> -VMId <guid[]> [-ConfigurationName <string>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -VMName <string[]> [-ConfigurationName <string>] [-Name <string[]>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession -InstanceId <guid[]> -VMName <string[]> [-ConfigurationName <string>] [-State <SessionFilterState>] [<CommonParameters>]
Get-PSSession [-InstanceId <guid[]>] [<CommonParameters>]
Get-PSSession [-Id] <int[]> [<CommonParameters>]
```

Example: Get sessions created in the current session

```powershell
Get-PSSession
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-PSSession.md)


### Get-PSSubsystem

Version: 7 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-PSSubsystem [<CommonParameters>]
Get-PSSubsystem -Kind <SubsystemKind> [<CommonParameters>]
Get-PSSubsystem -SubsystemType <type> [<CommonParameters>]
```

Example: Display all available subsystems

```powershell
Get-PSSubsystem
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-PSSubsystem.md)


### Get-Random

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Get-Random [[-Maximum] <Object>] [-SetSeed <int>] [-Minimum <Object>] [-Count <int>] [<CommonParameters>]
Get-Random [-InputObject] <Object[]> [-SetSeed <int>] [-Count <int>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Random [[-Maximum] <Object>] [-SetSeed <int>] [-Minimum <Object>] [-Count <int>] [<CommonParameters>]
Get-Random [-InputObject] <Object[]> [-SetSeed <int>] [-Count <int>] [<CommonParameters>]
Get-Random [-InputObject] <Object[]> -Shuffle [-SetSeed <int>] [<CommonParameters>]
```

Example: Get a random integer

```powershell
Get-Random
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Random.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: randomness. `shuf` / `$RANDOM` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Samples randomly among these objects (1 by default) |
| `-Minimum` / `-Maximum` | int | Random within [Minimum, Maximum) |
| `-Count` | int | Draws N items randomly from the input |

- Implementation: crypto/rand.


### Get-Runspace

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Runspace [[-Name] <string[]>] [<CommonParameters>]
Get-Runspace [-Id] <int[]> [<CommonParameters>]
Get-Runspace [-InstanceId] <guid[]> [<CommonParameters>]
```

Example: Get runspaces

```powershell
Get-Runspace
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Runspace.md)


### Get-RunspaceDebug

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-RunspaceDebug [[-RunspaceName] <string[]>] [<CommonParameters>]
Get-RunspaceDebug [-Runspace] <runspace[]> [<CommonParameters>]
Get-RunspaceDebug [-RunspaceId] <int[]> [<CommonParameters>]
Get-RunspaceDebug [-RunspaceInstanceId] <guid[]> [<CommonParameters>]
Get-RunspaceDebug [[-ProcessName] <string>] [[-AppDomainName] <string[]>] [<CommonParameters>]
```

Example: Show the state of the default runspace debugger

```powershell
Get-RunspaceDebug
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-RunspaceDebug.md)


### Get-SecureRandom

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-SecureRandom [[-Maximum] <Object>] [-Minimum <Object>] [-Count <int>] [<CommonParameters>]
Get-SecureRandom [-InputObject] <Object[]> [-Count <int>] [<CommonParameters>]
Get-SecureRandom [-InputObject] <Object[]> -Shuffle [<CommonParameters>]
```

Example: Get a random integer

```powershell
Get-SecureRandom
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-SecureRandom.md)


### Get-TimeZone

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-TimeZone [[-Name] <string[]>] [<CommonParameters>]
Get-TimeZone -Id <string[]> [<CommonParameters>]
Get-TimeZone -ListAvailable [<CommonParameters>]
```

Example: Get the current time zone

```powershell
Get-TimeZone
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-TimeZone.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Reads /etc/timezone.

- Type: Go implementation (reads /etc/timezone).
- Companion: Set-TimeZone.
- Function: shows the current time zone.
- Parameters: none.
- Implementation: reads /etc/timezone.
- Output: a TimeZoneInfo object with fields Id, DisplayName.


### Get-TraceSource

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-TraceSource [[-Name] <string[]>] [<CommonParameters>]
```

Example: Get trace sources by name

```powershell
Get-TraceSource -Name "*Provider*"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-TraceSource.md)


### Get-TypeData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-TypeData [[-TypeName] <string[]>] [<CommonParameters>]
```

Example: Get all extended type data

```powershell
Get-TypeData
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-TypeData.md)


### Get-UICulture

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-UICulture [<CommonParameters>]
```

Example: Get the UI culture

```powershell
Get-UICulture
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-UICulture.md)


### Get-Unique

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Get-Unique [-InputObject <psobject>] [-AsString] [<CommonParameters>]
Get-Unique [-InputObject <psobject>] [-OnType] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Unique [-InputObject <psobject>] [-AsString] [-CaseInsensitive] [<CommonParameters>]
Get-Unique [-InputObject <psobject>] [-OnType] [-CaseInsensitive] [<CommonParameters>]
```

Example: Get unique words in a text file

```powershell
$A = $( foreach ($line in Get-Content C:\Test1\File1.txt) {
    $line.ToLower().Split(" ")
  }) | Sort-Object | Get-Unique
$A.Count
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Unique.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Deduplicates by object string; inputs need not be adjacent.

- Type: Go implementation.
- Function: deduplicates. `sort -u` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Input objects |
| `-AsString` | switch | Accepted, no extra effect |

- Implementation: keeps the first occurrence of each object string.


### Get-Uptime

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Uptime [<CommonParameters>]
Get-Uptime [-Since] [<CommonParameters>]
```

Example: Show time since last boot

```powershell
Get-Uptime
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Uptime.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Always 0 on other platforms.

- Type: Go implementation.
- Function: uptime duration. Bash's `uptime -p`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| (none) | | |

- Implementation: on Linux reads seconds from /proc/uptime; zero elsewhere.
- Output: a TimeSpan object with fields Days/Hours/Minutes/Seconds/TotalSeconds (plus TotalMilliseconds/TotalMinutes/TotalHours). Table columns Days/Hours/Minutes/Seconds.


### Get-Variable

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Variable [[-Name] <string[]>] [-ValueOnly] [-Include <string[]>] [-Exclude <string[]>] [-Scope <string>] [<CommonParameters>]
```

Example: Get variables by letter

```powershell
Get-Variable m*
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Variable.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: lists variables. Bash's `env` / `declare` counterparts.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable names, wildcards and arrays supported |

- Output: PSVariable objects with fields Name, Value. Automatic variables (PWD/HOME/PID/PSVersionTable/LASTEXITCODE/?/PSCommandPath/args/Host/PSEdition/IsLinux/IsWindows/IsMacOS/PSHOME/OFS) appear too.


### Get-Verb

Version: Both

Module: 5.1 in Microsoft.PowerShell.Core, 7 in Microsoft.PowerShell.Utility

Syntax:

```powershell
Get-Verb [[-Verb] <string[]>] [[-Group] <string[]>] [<CommonParameters>]
```

Example: Get a list of all verbs

```powershell
Get-Verb
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Get-Verb.md) / [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Get-Verb.md)


### Group-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Group-Object [[-Property] <Object[]>] [-NoElement] [-AsHashTable] [-AsString] [-InputObject <psobject>] [-Culture <string>] [-CaseSensitive] [<CommonParameters>]
```

Example: Group files by extension

```powershell
$files = Get-ChildItem -Path $PSHOME -Recurse
$files |
    Group-Object -Property Extension -NoElement |
    Sort-Object -Property Count -Descending
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Group-Object.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original.

- Type: Go implementation.
- Function: groups and counts. `sort | uniq -c` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string | Property to group by; none given, groups by object string |
| `-CaseSensitive` | switch | Case-sensitive grouping; insensitive by default |

- Implementation: groups by property value (or object string when absent), case-insensitive by default (merged after folding), with Name taken from the first original value seen in each group.
- Output: GroupInfo objects with fields Name, Count. Table columns Count/Name.


### Import-Alias

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Import-Alias [-Path] <string> [-Scope <string>] [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Import-Alias -LiteralPath <string> [-Scope <string>] [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Import aliases from a file

```powershell
Import-Alias test.txt
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Import-Alias.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Alias, New-Alias, Remove-Alias, Import-Alias, Export-Alias.
- Function: manages aliases.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Alias file path |

- Set-Alias overwrites; New-Alias errors on an existing alias without -Force; Remove-Alias deletes; Export-Alias writes one `name=target` per line; Import-Alias reads them back.


### Import-Clixml

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Import-Clixml [-Path] <string[]> [-IncludeTotalCount] [-Skip <uint64>] [-First <uint64>] [<CommonParameters>]
Import-Clixml -LiteralPath <string[]> [-IncludeTotalCount] [-Skip <uint64>] [-First <uint64>] [<CommonParameters>]
```

Syntax (7):

```powershell
Import-Clixml [-Path] <string[]> [-IncludeTotalCount] [-Skip <ulong>] [-First <ulong>] [<CommonParameters>]
Import-Clixml -LiteralPath <string[]> [-IncludeTotalCount] [-Skip <ulong>] [-First <ulong>] [<CommonParameters>]
```

Example: Import a serialized file and recreate an object

```powershell
Get-Process | Export-Clixml -Path .\pi.xml
$Processes = Import-Clixml -Path .\pi.xml
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Import-Clixml.md)


### Import-Csv

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Import-Csv [[-Path] <string[]>] [[-Delimiter] <char>] [-LiteralPath <string[]>] [-Header <string[]>] [-Encoding <string>] [<CommonParameters>]
Import-Csv [[-Path] <string[]>] -UseCulture [-LiteralPath <string[]>] [-Header <string[]>] [-Encoding <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Import-Csv [-Path] <string[]> [[-Delimiter] <char>] [-Header <string[]>] [-Encoding <Encoding>] [<CommonParameters>]
Import-Csv [[-Delimiter] <char>] -LiteralPath <string[]> [-Header <string[]>] [-Encoding <Encoding>] [<CommonParameters>]
Import-Csv [-Path] <string[]> -UseCulture [-Header <string[]>] [-Encoding <Encoding>] [<CommonParameters>]
Import-Csv -LiteralPath <string[]> -UseCulture [-Header <string[]>] [-Encoding <Encoding>] [<CommonParameters>]
```

Example: Import process objects

```powershell
Get-Process | Export-Csv -Path .\Processes.csv
$P = Import-Csv -Path .\Processes.csv
$P | Get-Member
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Import-Csv.md)


### Import-LocalizedData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Import-LocalizedData [[-BindingVariable] <string>] [[-UICulture] <string>] [-BaseDirectory <string>] [-FileName <string>] [-SupportedCommand <string[]>] [<CommonParameters>]
```

Example: Import text strings

```powershell
Import-LocalizedData -BindingVariable "Messages"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Import-LocalizedData.md)


### Import-Module

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Import-Module [-Name] <string[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-Name] <string[]> -PSSession <PSSession> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-Name] <string[]> -CimSession <CimSession> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [-CimResourceUri <uri>] [-CimNamespace <string>] [<CommonParameters>]
Import-Module [-FullyQualifiedName] <ModuleSpecification[]> -PSSession <PSSession> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-FullyQualifiedName] <ModuleSpecification[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-Assembly] <Assembly[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-ModuleInfo] <psmoduleinfo[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Import-Module [-Name] <string[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-Name] <string[]> -PSSession <PSSession> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-Name] <string[]> -CimSession <CimSession> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [-CimResourceUri <uri>] [-CimNamespace <string>] [<CommonParameters>]
Import-Module [-Name] <string[]> -UseWindowsPowerShell [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-MinimumVersion <version>] [-MaximumVersion <string>] [-RequiredVersion <version>] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-FullyQualifiedName] <ModuleSpecification[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-FullyQualifiedName] <ModuleSpecification[]> -PSSession <PSSession> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-FullyQualifiedName] <ModuleSpecification[]> -UseWindowsPowerShell [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-Assembly] <Assembly[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
Import-Module [-ModuleInfo] <psmoduleinfo[]> [-Global] [-Prefix <string>] [-Function <string[]>] [-Cmdlet <string[]>] [-Variable <string[]>] [-Alias <string[]>] [-Force] [-SkipEditionCheck] [-PassThru] [-AsCustomObject] [-ArgumentList <Object[]>] [-DisableNameChecking] [-NoClobber] [-Scope <string>] [<CommonParameters>]
```

Example: Import the members of a module into the current session

```powershell
Import-Module -Name PSDiagnostics
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Import-Module.md)


### Import-PackageProvider

Version: Both

Module: PackageManagement

Syntax:

```powershell
Import-PackageProvider [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Force] [-ForceBootstrap] [<CommonParameters>]
```

Example: Import a package provider from the local computer

```powershell
PS C:\> Import-PackageProvider -Name "Nuget"
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/import-packageprovider?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/import-packageprovider?view=powershell-7.5)


### Import-PowerShellDataFile

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Import-PowerShellDataFile [-Path] <string[]> [-SkipLimitCheck] [<CommonParameters>]
Import-PowerShellDataFile [-LiteralPath] <string[]> [-SkipLimitCheck] [<CommonParameters>]
```

Example: Retrieve values from PSD1

```powershell
Get-Content .\Configuration.psd1
$config = Import-PowerShellDataFile .\Configuration.psd1
$config.AllNodes
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/Import-PowerShellDataFile.md) / [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Import-PowerShellDataFile.md)


### Import-PSSession

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Import-PSSession [-Session] <PSSession> [[-CommandName] <string[]>] [[-FormatTypeName] <string[]>] [-Prefix <string>] [-DisableNameChecking] [-AllowClobber] [-ArgumentList <Object[]>] [-CommandType <CommandTypes>] [-Module <string[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Certificate <X509Certificate2>] [<CommonParameters>]
```

Example: Import all commands from a PSSession

```powershell
$S = New-PSSession -ComputerName Server01
Import-PSSession -Session $S
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Import-PSSession.md)


### Install-Package

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Install-Package [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Source <string[]>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string[]>] [<CommonParameters>]
Install-Package [-InputObject] <SoftwareIdentity[]> [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-IncludeWindowsInstaller] [-IncludeSystemComponent] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-IncludeWindowsInstaller] [-IncludeSystemComponent] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-AdditionalArguments <string[]>] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-AdditionalArguments <string[]>] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-Scope <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-Scope <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [<CommonParameters>]
```

Syntax (7):

```powershell
Install-Package [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Source <string[]>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string[]>] [<CommonParameters>]
Install-Package [-InputObject] <SoftwareIdentity[]> [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [-Headers <string[]>] [-FilterOnTag <string[]>] [-Contains <string>] [-AllowPrereleaseVersions] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [-Headers <string[]>] [-FilterOnTag <string[]>] [-Contains <string>] [-AllowPrereleaseVersions] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-AllowPrereleaseVersions] [-Scope <string>] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-AcceptLicense] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [<CommonParameters>]
Install-Package [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-AllowPrereleaseVersions] [-Scope <string>] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-AcceptLicense] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [<CommonParameters>]
```

Example: Install a package by package name

```powershell
PS> Install-Package -Name NuGet.Core -Source MyNuGet -Credential Contoso\TestUser
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/install-package?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/install-package?view=powershell-7.5)


### Install-PackageProvider

Version: Both

Module: PackageManagement

Syntax:

```powershell
Install-PackageProvider [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Credential <pscredential>] [-Scope <string>] [-Source <string[]>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Install-PackageProvider [-InputObject] <SoftwareIdentity[]> [-Scope <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Install a package provider from the PowerShell Gallery

```powershell
Install-PackageProvider -Name "GistProvider" -Verbose
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/install-packageprovider?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/install-packageprovider?view=powershell-7.5)


### Install-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Install-PSResource [-Name] <string[]> [-Version <string>] [-Prerelease] [-Repository <string[]>] [-Credential <pscredential>] [-Scope <ScopeType>] [-TemporaryPath <string>] [-TrustRepository] [-Reinstall] [-Quiet] [-AcceptLicense] [-NoClobber] [-SkipDependencyCheck] [-AuthenticodeCheck] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Install-PSResource [-InputObject] <PSResourceInfo[]> [-Repository <string[]>] [-Credential <pscredential>] [-Scope <ScopeType>] [-TemporaryPath <string>] [-TrustRepository] [-Reinstall] [-Quiet] [-AcceptLicense] [-NoClobber] [-SkipDependencyCheck] [-AuthenticodeCheck] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Install-PSResource -RequiredResourceFile <string> [-Credential <pscredential>] [-Scope <ScopeType>] [-TemporaryPath <string>] [-TrustRepository] [-Reinstall] [-Quiet] [-AcceptLicense] [-NoClobber] [-SkipDependencyCheck] [-AuthenticodeCheck] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Install-PSResource -RequiredResource <Object> [-Credential <pscredential>] [-Scope <ScopeType>] [-TemporaryPath <string>] [-TrustRepository] [-Reinstall] [-Quiet] [-AcceptLicense] [-NoClobber] [-SkipDependencyCheck] [-AuthenticodeCheck] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Install-PSResource Az -Repository PSGallery
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/install-psresource?view=powershell-7.5)


### Invoke-Command

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Invoke-Command [-ScriptBlock] <scriptblock> [-NoNewScope] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-Session] <PSSession[]>] [-ScriptBlock] <scriptblock> [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-Session] <PSSession[]>] [-FilePath] <string> [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-ComputerName] <string[]>] [-FilePath] <string> [-Credential <pscredential>] [-Port <int>] [-UseSSL] [-ConfigurationName <string>] [-ApplicationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-SessionName <string[]>] [-HideComputerName] [-JobName <string>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-ComputerName] <string[]>] [-ScriptBlock] <scriptblock> [-Credential <pscredential>] [-Port <int>] [-UseSSL] [-ConfigurationName <string>] [-ApplicationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-SessionName <string[]>] [-HideComputerName] [-JobName <string>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-CertificateThumbprint <string>] [<CommonParameters>]
Invoke-Command [-VMId] <guid[]> [-ScriptBlock] <scriptblock> -Credential <pscredential> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-ConnectionUri] <uri[]>] [-ScriptBlock] <scriptblock> [-Credential <pscredential>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-HideComputerName] [-JobName <string>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-CertificateThumbprint <string>] [<CommonParameters>]
Invoke-Command [[-ConnectionUri] <uri[]>] [-FilePath] <string> [-Credential <pscredential>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-HideComputerName] [-JobName <string>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-ScriptBlock] <scriptblock> -Credential <pscredential> -VMName <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-VMId] <guid[]> [-FilePath] <string> -Credential <pscredential> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-FilePath] <string> -Credential <pscredential> -VMName <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-ScriptBlock] <scriptblock> -ContainerId <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-RunAsAdministrator] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-FilePath] <string> -ContainerId <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-RunAsAdministrator] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Invoke-Command [-ScriptBlock] <scriptblock> [-NoNewScope] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-Session] <PSSession[]>] [-ScriptBlock] <scriptblock> [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-Session] <PSSession[]>] [-FilePath] <string> [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-ComputerName] <string[]>] [-ScriptBlock] <scriptblock> [-Credential <pscredential>] [-Port <int>] [-UseSSL] [-ConfigurationName <string>] [-ApplicationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-SessionName <string[]>] [-HideComputerName] [-JobName <string>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-CertificateThumbprint <string>] [<CommonParameters>]
Invoke-Command [[-ComputerName] <string[]>] [-FilePath] <string> [-Credential <pscredential>] [-Port <int>] [-UseSSL] [-ConfigurationName <string>] [-ApplicationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-SessionName <string[]>] [-HideComputerName] [-JobName <string>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [[-ConnectionUri] <uri[]>] [-ScriptBlock] <scriptblock> [-Credential <pscredential>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-HideComputerName] [-JobName <string>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-CertificateThumbprint <string>] [<CommonParameters>]
Invoke-Command [[-ConnectionUri] <uri[]>] [-FilePath] <string> [-Credential <pscredential>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-InDisconnectedSession] [-HideComputerName] [-JobName <string>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-EnableNetworkAccess] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-VMId] <guid[]> [-ScriptBlock] <scriptblock> -Credential <pscredential> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-ScriptBlock] <scriptblock> -Credential <pscredential> -VMName <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-VMId] <guid[]> [-FilePath] <string> -Credential <pscredential> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-FilePath] <string> -Credential <pscredential> -VMName <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-ScriptBlock] <scriptblock> -HostName <string[]> [-Port <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-UserName <string>] [-KeyFilePath <string>] [-Subsystem <string>] [-ConnectingTimeout <int>] [-SSHTransport] [-Options <hashtable>] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-ScriptBlock] <scriptblock> -ContainerId <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-RunAsAdministrator] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-FilePath] <string> -ContainerId <string[]> [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AsJob] [-HideComputerName] [-JobName <string>] [-RunAsAdministrator] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command [-ScriptBlock] <scriptblock> -SSHConnection <hashtable[]> [-AsJob] [-HideComputerName] [-JobName <string>] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command -FilePath <string> -HostName <string[]> [-AsJob] [-HideComputerName] [-UserName <string>] [-KeyFilePath <string>] [-Subsystem <string>] [-ConnectingTimeout <int>] [-SSHTransport] [-Options <hashtable>] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Invoke-Command -FilePath <string> -SSHConnection <hashtable[]> [-AsJob] [-HideComputerName] [-RemoteDebug] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
```

Example: Run a script on a server

```powershell
Invoke-Command -FilePath C:\scripts\test.ps1 -ComputerName Server01
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Invoke-Command.md)


### Invoke-Expression

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Invoke-Expression [-Command] <string> [<CommonParameters>]
```

Example: Evaluate an expression

```powershell
$Command = "Get-Process"
$Command
Invoke-Expression $Command
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Invoke-Expression.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: executes strings as commands. Bash's `eval`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Command` (position 0) | string | Source code to execute; named/positional take precedence, pipeline input only as fallback |

- Implementation: goes through `RunSource` (parse + statement-by-statement execution), returning output objects.


### Invoke-History

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Invoke-History [[-Id] <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Run the most recent command in the history

```powershell
Invoke-History
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Invoke-History.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Get-History, Clear-History, Add-History.
- Function: replays a history entry. Bash's `!!`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Id` | int | History number |
| `-InputObject` (position 0) | object | Command text (omitting it takes the last entry) |

- Behavior: Invoke-History echoes the command before replaying it (parsed and executed through RunSource).


### Invoke-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Invoke-Item [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Invoke-Item -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Invoke-Item [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-Item -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Open a file

```powershell
Invoke-Item "C:\Test\aliasApr04.doc"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Invoke-Item.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Currently only prints the path without opening it.

- Type: Go implementation (placeholder).
- Function: opens a file with its default application (`xdg-open` territory). **Currently just prints the path without actually calling xdg-open — a placeholder.**

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | File to open; for now its path is merely printed |


### Invoke-RestMethod

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Invoke-RestMethod [-Uri] <uri> [-Method <WebRequestMethod>] [-UseBasicParsing] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-UserAgent <string>] [-DisableKeepAlive] [-TimeoutSec <int>] [-Headers <IDictionary>] [-MaximumRedirection <int>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ProxyUseDefaultCredentials] [-Body <Object>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [<CommonParameters>]
```

Syntax (7):

```powershell
Invoke-RestMethod [-Uri] <uri> [-FollowRelLink] [-MaximumFollowRelLink <int>] [-ResponseHeadersVariable <string>] [-StatusCodeVariable <string>] [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-Method <WebRequestMethod>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ProxyUseDefaultCredentials] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
Invoke-RestMethod [-Uri] <uri> -NoProxy [-FollowRelLink] [-MaximumFollowRelLink <int>] [-ResponseHeadersVariable <string>] [-StatusCodeVariable <string>] [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-Method <WebRequestMethod>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
Invoke-RestMethod [-Uri] <uri> -CustomMethod <string> [-FollowRelLink] [-MaximumFollowRelLink <int>] [-ResponseHeadersVariable <string>] [-StatusCodeVariable <string>] [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ProxyUseDefaultCredentials] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
Invoke-RestMethod [-Uri] <uri> -CustomMethod <string> -NoProxy [-FollowRelLink] [-MaximumFollowRelLink <int>] [-ResponseHeadersVariable <string>] [-StatusCodeVariable <string>] [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
```

Example (5.1): Get the PowerShell RSS feed

```powershell
Invoke-RestMethod -Uri https://devblogs.microsoft.com/powershell/feed/ |
  Format-Table -Property Title, pubDate
```

Example (7): Get the PowerShell RSS feed

```powershell
Invoke-RestMethod -Uri https://blogs.msdn.microsoft.com/powershell/feed/ |
  Format-Table -Property Title, pubDate
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Invoke-RestMethod.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Non-JSON responses return plain text.

- Type: Go implementation.
- Function: HTTP requests with JSON parsing. `curl | jq` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Uri` (position 0) | string | Request address |
| `-Method` | string | HTTP method, GET by default |
| `-Body` | object | Request body, used as such for POST/PUT/PATCH |

- Implementation: after the request, JSON parsing is attempted (reusing ConvertFrom-Json's object construction); non-JSON comes back as text.


### Invoke-WebRequest

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Invoke-WebRequest [-Uri] <uri> [-UseBasicParsing] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-UserAgent <string>] [-DisableKeepAlive] [-TimeoutSec <int>] [-Headers <IDictionary>] [-MaximumRedirection <int>] [-Method <WebRequestMethod>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ProxyUseDefaultCredentials] [-Body <Object>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [<CommonParameters>]
```

Syntax (7):

```powershell
Invoke-WebRequest [-Uri] <uri> [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-Method <WebRequestMethod>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ProxyUseDefaultCredentials] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
Invoke-WebRequest [-Uri] <uri> -NoProxy [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-Method <WebRequestMethod>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
Invoke-WebRequest [-Uri] <uri> -CustomMethod <string> [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ProxyUseDefaultCredentials] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
Invoke-WebRequest [-Uri] <uri> -CustomMethod <string> -NoProxy [-UseBasicParsing] [-HttpVersion <version>] [-WebSession <WebRequestSession>] [-SessionVariable <string>] [-AllowUnencryptedAuthentication] [-Authentication <WebAuthenticationType>] [-Credential <pscredential>] [-UseDefaultCredentials] [-CertificateThumbprint <string>] [-Certificate <X509Certificate>] [-SkipCertificateCheck] [-SslProtocol <WebSslProtocol>] [-Token <securestring>] [-UserAgent <string>] [-DisableKeepAlive] [-ConnectionTimeoutSeconds <int>] [-OperationTimeoutSeconds <int>] [-Headers <IDictionary>] [-SkipHeaderValidation] [-AllowInsecureRedirect] [-MaximumRedirection <int>] [-MaximumRetryCount <int>] [-PreserveAuthorizationOnRedirect] [-RetryIntervalSec <int>] [-PreserveHttpMethodOnRedirect] [-UnixSocket <UnixDomainSocketEndPoint>] [-Body <Object>] [-Form <IDictionary>] [-ContentType <string>] [-TransferEncoding <string>] [-InFile <string>] [-OutFile <string>] [-PassThru] [-Resume] [-SkipHttpErrorCheck] [<CommonParameters>]
```

Example (5.1): Send a web request

```powershell
$Response = Invoke-WebRequest -UseBasicParsing -Uri https://www.bing.com?q=how+many+feet+in+a+mile
$Response.InputFields |
    Where-Object Name -Like "* Value" |
    Select-Object Name, Value
```

Example (7): Send a web request

```powershell
$Response = Invoke-WebRequest -Uri https://www.bing.com/search?q=how+many+feet+in+a+mile
$Response.InputFields | Where-Object {
    $_.Name -like "* Value*"
} | Select-Object Name, Value
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Invoke-WebRequest.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Returns plain text (split into strings by line) with no properties like StatusCode.

- Type: Go implementation.
- Function: HTTP requests returning response text. `curl` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Uri` (position 0) | string | Request address |
| `-Method` | string | HTTP method, GET by default |
| `-Body` | object | Request body, used as such for POST/PUT/PATCH |

- Implementation: net/http with a 30-second timeout.
- Output: the response text split into Strings by line.


### Join-Path

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Join-Path [-Path] <string[]> [-ChildPath] <string> [-Resolve] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Join-Path [-Path] <string[]> [-ChildPath] <string[]> [[-AdditionalChildPath] <string[]>] [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Combine a path with a child path

```powershell
Join-Path -Path "path" -ChildPath "childpath"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Join-Path.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: joins paths.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Base path |
| `-ChildPath` (position 1) | path | Child path to append |

- Implementation: `filepath.Join(base, child)`.


### Join-String

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Join-String [[-Property] <pspropertyexpression>] [[-Separator] <string>] [-OutputPrefix <string>] [-OutputSuffix <string>] [-UseCulture] [-InputObject <psobject[]>] [<CommonParameters>]
Join-String [[-Property] <pspropertyexpression>] [[-Separator] <string>] [-OutputPrefix <string>] [-OutputSuffix <string>] [-SingleQuote] [-UseCulture] [-InputObject <psobject[]>] [<CommonParameters>]
Join-String [[-Property] <pspropertyexpression>] [[-Separator] <string>] [-OutputPrefix <string>] [-OutputSuffix <string>] [-DoubleQuote] [-UseCulture] [-InputObject <psobject[]>] [<CommonParameters>]
Join-String [[-Property] <pspropertyexpression>] [[-Separator] <string>] [-OutputPrefix <string>] [-OutputSuffix <string>] [-FormatString <string>] [-UseCulture] [-InputObject <psobject[]>] [<CommonParameters>]
```

Example: Join directory names

```powershell
Get-ChildItem -Directory C:\ | Join-String -Property Name -DoubleQuote -Separator ', '
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Join-String.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Without `-Separator`, items join with a space by default (PowerShell joins with no separator).

- Type: Go implementation.
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


### Measure-Command

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Measure-Command [-Expression] <scriptblock> [-InputObject <psobject>] [<CommonParameters>]
```

Example: Measure a command

```powershell
Measure-Command { Get-EventLog "Windows PowerShell" }
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Measure-Command.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: timed execution. Bash's `time`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Expression` (position 0) | scriptblock | The block to time |

- Output: a TimeSpan object with fields Days/Hours/Minutes/Seconds/TotalMilliseconds/TotalSeconds (plus TotalMinutes/TotalHours).


### Measure-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Measure-Object [[-Property] <string[]>] [-InputObject <psobject>] [-Sum] [-Average] [-Maximum] [-Minimum] [<CommonParameters>]
Measure-Object [[-Property] <string[]>] [-InputObject <psobject>] [-Line] [-Word] [-Character] [-IgnoreWhiteSpace] [<CommonParameters>]
```

Syntax (7):

```powershell
Measure-Object [[-Property] <pspropertyexpression[]>] [-InputObject <psobject>] [-StandardDeviation] [-Sum] [-AllStats] [-Average] [-Maximum] [-Minimum] [<CommonParameters>]
Measure-Object [[-Property] <pspropertyexpression[]>] [-InputObject <psobject>] [-Line] [-Word] [-Character] [-IgnoreWhiteSpace] [<CommonParameters>]
```

Example: Count the files and folders in a directory

```powershell
Get-ChildItem | Measure-Object
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Measure-Object.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Min/Max over mixed numeric and non-numeric input only counts numbers (PowerShell compares non-numbers as strings); non-numeric input yields empty instead of erroring (PowerShell reports a non-terminating error).

- Type: Go implementation.
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


### Move-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Move-Item [-Path] <string[]> [[-Destination] <string>] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Move-Item [[-Destination] <string>] -LiteralPath <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Move-Item [-Path] <string[]> [[-Destination] <string>] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Move-Item [[-Destination] <string>] -LiteralPath <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Move a file to another directory and rename it

```powershell
Move-Item -Path C:\test.txt -Destination E:\Temp\tst.txt
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Move-Item.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Moving directories requires `-Recurse`.

- Type: Go implementation.
- Function: moves/renames. Implemented as copy followed by deleting the source.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Source path |
| `-Destination` (position 1) | path | Destination path |
| `-Recurse` | switch | Moves a whole directory tree (same rule as Copy-Item, directories need it) |
| `-Force` | switch | Accepted, no extra effect at present |

- Behavior: same as Copy-Item (directories need -Recurse), with the source deleted once copying succeeds.


### Move-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Move-ItemProperty [-Path] <string[]> [-Destination] <string> [-Name] <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Move-ItemProperty [-Destination] <string> [-Name] <string[]> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Move-ItemProperty [-Path] <string[]> [-Destination] <string> [-Name] <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Move-ItemProperty [-Destination] <string> [-Name] <string[]> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Move a registry value and its data to another key

```powershell
$params = @{
    Path        = 'HKLM:\Software\MyCompany\MyApp'
    Name        = 'Version'
    Destination = 'HKLM:\Software\MyCompany\NewApp'
}
Move-ItemProperty @params
```

Example (7): Move a registry value and its data to another key

```powershell
Move-ItemProperty "HKLM:\Software\MyCompany\MyApp" -Name "Version" -Destination "HKLM:\Software\MyCompany\NewApp"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Move-ItemProperty.md)


### New-Alias

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
New-Alias [-Name] <string> [-Value] <string> [-Description <string>] [-Option <ScopedItemOptions>] [-PassThru] [-Scope <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create an alias for a cmdlet

```powershell
New-Alias -Name "List" Get-ChildItem
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-Alias.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Alias, New-Alias, Remove-Alias, Import-Alias, Export-Alias.
- Function: manages aliases.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Alias name |
| `-Value` (position 1) | string | Target command |
| `-Force` | switch | Lets New-Alias overwrite an existing alias |

- Set-Alias overwrites; New-Alias errors on an existing alias without -Force; Remove-Alias deletes; Export-Alias writes one `name=target` per line; Import-Alias reads them back.


### New-Event

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
New-Event [-SourceIdentifier] <string> [[-Sender] <psobject>] [[-EventArguments] <psobject[]>] [[-MessageData] <psobject>] [<CommonParameters>]
```

Example: Create a new event in the event queue

```powershell
PS C:\> New-Event -SourceIdentifier Timer -Sender Windows.Timer -MessageData "Test"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-Event.md)


### New-Guid

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
New-Guid [<CommonParameters>]
```

Syntax (7):

```powershell
New-Guid [<CommonParameters>]
New-Guid [-Empty] [<CommonParameters>]
New-Guid [[-InputObject] <string>] [<CommonParameters>]
```

Example (5.1): Create a GUID

```powershell
New-Guid
```

Example (7): Create a new GUID

```powershell
New-Guid
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-Guid.md) / [Official reference source (5.1)](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/New-Guid.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 7.

- Type: Go implementation.
- Function: generates UUIDs. `uuidgen` counterpart.
- Parameters: none.
- Implementation: crypto/rand producing RFC 4122 version 4.


### New-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
New-Item [-Path] <string[]> [-ItemType <string>] [-Value <Object>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
New-Item [[-Path] <string[]>] -Name <string> [-ItemType <string>] [-Value <Object>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
New-Item [-Path] <string[]> [-ItemType <string>] [-Value <Object>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-Item [[-Path] <string[]>] -Name <string> [-ItemType <string>] [-Value <Object>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a file in the current directory

```powershell
New-Item -Path . -Name "testfile1.txt" -ItemType "File" -Value "This is a text string."
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/New-Item.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: creates files or directories.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Path to create |
| `-ItemType` | string | "Directory" makes a directory (`mkdir`); any other value makes a file (`touch`) |
| `-Force` | switch | Directories go through MkdirAll (`mkdir -p`); existing files don't raise an error |

- Output: the FileInfo/DirectoryInfo object of the newly created item.
- Behavior: existing file without -Force → error (consistent with PowerShell semantics, unlike `touch`'s silent success).


### New-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
New-ItemProperty [-Path] <string[]> [-Name] <string> [-PropertyType <string>] [-Value <Object>] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
New-ItemProperty [-Name] <string> -LiteralPath <string[]> [-PropertyType <string>] [-Value <Object>] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
New-ItemProperty [-Path] <string[]> [-Name] <string> [-PropertyType <string>] [-Value <Object>] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-ItemProperty [-Name] <string> -LiteralPath <string[]> [-PropertyType <string>] [-Value <Object>] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Add a registry entry

```powershell
New-ItemProperty -Path "HKLM:\Software\MyCompany" -Name "NoOfEmployees" -Value 822
Get-ItemProperty "HKLM:\Software\MyCompany"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/New-ItemProperty.md)


### New-Module

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
New-Module [-ScriptBlock] <scriptblock> [-Function <string[]>] [-Cmdlet <string[]>] [-ReturnResult] [-AsCustomObject] [-ArgumentList <Object[]>] [<CommonParameters>]
New-Module [-Name] <string> [-ScriptBlock] <scriptblock> [-Function <string[]>] [-Cmdlet <string[]>] [-ReturnResult] [-AsCustomObject] [-ArgumentList <Object[]>] [<CommonParameters>]
```

Example: Create a dynamic module

```powershell
New-Module -ScriptBlock {function Hello {"Hello!"}}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-Module.md)


### New-ModuleManifest

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
New-ModuleManifest [-Path] <string> [-NestedModules <Object[]>] [-Guid <guid>] [-Author <string>] [-CompanyName <string>] [-Copyright <string>] [-RootModule <string>] [-ModuleVersion <version>] [-Description <string>] [-ProcessorArchitecture <ProcessorArchitecture>] [-PowerShellVersion <version>] [-ClrVersion <version>] [-DotNetFrameworkVersion <version>] [-PowerShellHostName <string>] [-PowerShellHostVersion <version>] [-RequiredModules <Object[]>] [-TypesToProcess <string[]>] [-FormatsToProcess <string[]>] [-ScriptsToProcess <string[]>] [-RequiredAssemblies <string[]>] [-FileList <string[]>] [-ModuleList <Object[]>] [-FunctionsToExport <string[]>] [-AliasesToExport <string[]>] [-VariablesToExport <string[]>] [-CmdletsToExport <string[]>] [-DscResourcesToExport <string[]>] [-CompatiblePSEditions <string[]>] [-PrivateData <Object>] [-Tags <string[]>] [-ProjectUri <uri>] [-LicenseUri <uri>] [-IconUri <uri>] [-ReleaseNotes <string>] [-HelpInfoUri <string>] [-PassThru] [-DefaultCommandPrefix <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
New-ModuleManifest [-Path] <string> [-NestedModules <Object[]>] [-Guid <guid>] [-Author <string>] [-CompanyName <string>] [-Copyright <string>] [-RootModule <string>] [-ModuleVersion <version>] [-Description <string>] [-ProcessorArchitecture <ProcessorArchitecture>] [-PowerShellVersion <version>] [-ClrVersion <version>] [-DotNetFrameworkVersion <version>] [-PowerShellHostName <string>] [-PowerShellHostVersion <version>] [-RequiredModules <Object[]>] [-TypesToProcess <string[]>] [-FormatsToProcess <string[]>] [-ScriptsToProcess <string[]>] [-RequiredAssemblies <string[]>] [-FileList <string[]>] [-ModuleList <Object[]>] [-FunctionsToExport <string[]>] [-AliasesToExport <string[]>] [-VariablesToExport <string[]>] [-CmdletsToExport <string[]>] [-DscResourcesToExport <string[]>] [-CompatiblePSEditions <string[]>] [-PrivateData <Object>] [-Tags <string[]>] [-ProjectUri <uri>] [-LicenseUri <uri>] [-IconUri <uri>] [-ReleaseNotes <string>] [-Prerelease <string>] [-RequireLicenseAcceptance] [-ExternalModuleDependencies <string[]>] [-HelpInfoUri <string>] [-PassThru] [-DefaultCommandPrefix <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a new module manifest

```powershell
New-ModuleManifest -Path C:\ps-test\Test-Module\Test-Module.psd1 -PassThru
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-ModuleManifest.md)


### New-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
New-Object [-TypeName] <string> [[-ArgumentList] <Object[]>] [-Property <IDictionary>] [<CommonParameters>]
New-Object [-ComObject] <string> [-Strict] [-Property <IDictionary>] [<CommonParameters>]
```

Example: Create a System.Version object

```powershell
New-Object -TypeName System.Version -ArgumentList "1.2.3.4"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-Object.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Only `PSObject` / `PSCustomObject` are supported; other types (e.g. System.Collections.ArrayList) report "not supported".

- Type: Go implementation.
- Function: constructs custom objects. Counterpart of the `[pscustomobject]@{...}` type literal.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-TypeName` (position 0) | string | Type name: currently `PSObject` / `pscustomobject`; other types get an "unsupported" error |
| `-Property` | hashtable | Hashtable whose keys become properties in order |

- Implementation: both `PSObject` and `pscustomobject` construct a `System.Management.Automation.PSCustomObject`; without `-Property` an empty object results.
- The type literal `[pscustomobject]@{ a = 1; b = "x" }` equals `New-Object pscustomobject -Property @{ a = 1; b = "x" }`.


### New-PSDrive

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
New-PSDrive [-Name] <string> [-PSProvider] <string> [-Root] <string> [-Description <string>] [-Scope <string>] [-Persist] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
New-PSDrive [-Name] <string> [-PSProvider] <string> [-Root] <string> [-Description <string>] [-Scope <string>] [-Persist] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a temporary drive mapped to a network share

```powershell
New-PSDrive -Name "Public" -PSProvider "FileSystem" -Root "\\Server01\Public"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/New-PSDrive.md)


### New-PSRoleCapabilityFile

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
New-PSRoleCapabilityFile [-Path] <string> [-Guid <guid>] [-Author <string>] [-Description <string>] [-CompanyName <string>] [-Copyright <string>] [-ModulesToImport <Object[]>] [-VisibleAliases <string[]>] [-VisibleCmdlets <Object[]>] [-VisibleFunctions <Object[]>] [-VisibleExternalCommands <string[]>] [-VisibleProviders <string[]>] [-ScriptsToProcess <string[]>] [-AliasDefinitions <IDictionary[]>] [-FunctionDefinitions <IDictionary[]>] [-VariableDefinitions <Object>] [-EnvironmentVariables <IDictionary>] [-TypesToProcess <string[]>] [-FormatsToProcess <string[]>] [-AssembliesToLoad <string[]>] [<CommonParameters>]
```

Example: Create a blank role capability file

```powershell
New-PSRoleCapabilityFile -Path ".\ExampleFile.psrc"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-PSRoleCapabilityFile.md)


### New-PSScriptFileInfo

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
New-PSScriptFileInfo [-Path] <string> -Description <string> [-Version <string>] [-Author <string>] [-Guid <guid>] [-CompanyName <string>] [-Copyright <string>] [-RequiredModules <hashtable[]>] [-ExternalModuleDependencies <string[]>] [-RequiredScripts <string[]>] [-ExternalScriptDependencies <string[]>] [-Tags <string[]>] [-ProjectUri <string>] [-LicenseUri <string>] [-IconUri <string>] [-ReleaseNotes <string>] [-PrivateData <string>] [-Force] [<CommonParameters>]
```

Example: Creating an empty script with minimal information

```powershell
New-PSScriptFileInfo -Path ./test_script.ps1 -Description 'This is a test script.'
Get-Content ./test_script.ps1
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/new-psscriptfileinfo?view=powershell-7.5)


### New-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
New-PSSession [[-ComputerName] <string[]>] [-Credential <pscredential>] [-Name <string[]>] [-EnableNetworkAccess] [-ConfigurationName <string>] [-Port <int>] [-UseSSL] [-ApplicationName <string>] [-ThrottleLimit <int>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
New-PSSession [-VMId] <guid[]> -Credential <pscredential> [-Name <string[]>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession [-ConnectionUri] <uri[]> [-Credential <pscredential>] [-Name <string[]>] [-EnableNetworkAccess] [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
New-PSSession -Credential <pscredential> -VMName <string[]> [-Name <string[]>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession [[-Session] <PSSession[]>] [-Name <string[]>] [-EnableNetworkAccess] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession -ContainerId <string[]> [-Name <string[]>] [-ConfigurationName <string>] [-RunAsAdministrator] [-ThrottleLimit <int>] [<CommonParameters>]
```

Syntax (7):

```powershell
New-PSSession [[-ComputerName] <string[]>] [-Credential <pscredential>] [-Name <string[]>] [-EnableNetworkAccess] [-ConfigurationName <string>] [-Port <int>] [-UseSSL] [-ApplicationName <string>] [-ThrottleLimit <int>] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
New-PSSession [-ConnectionUri] <uri[]> [-Credential <pscredential>] [-Name <string[]>] [-EnableNetworkAccess] [-ConfigurationName <string>] [-ThrottleLimit <int>] [-AllowRedirection] [-SessionOption <PSSessionOption>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
New-PSSession [-VMId] <guid[]> -Credential <pscredential> [-Name <string[]>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession -Credential <pscredential> -VMName <string[]> [-Name <string[]>] [-ConfigurationName <string>] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession [[-Session] <PSSession[]>] [-Name <string[]>] [-EnableNetworkAccess] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession -ContainerId <string[]> [-Name <string[]>] [-ConfigurationName <string>] [-RunAsAdministrator] [-ThrottleLimit <int>] [<CommonParameters>]
New-PSSession -UseWindowsPowerShell [-Name <string[]>] [<CommonParameters>]
New-PSSession [-HostName] <string[]> [-Name <string[]>] [-Port <int>] [-UserName <string>] [-KeyFilePath <string>] [-Subsystem <string>] [-ConnectingTimeout <int>] [-SSHTransport] [-Options <hashtable>] [<CommonParameters>]
New-PSSession -SSHConnection <hashtable[]> [-Name <string[]>] [<CommonParameters>]
```

Example: Create a session on the local computer

```powershell
$s = New-PSSession
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-PSSession.md)


### New-PSSessionConfigurationFile

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
New-PSSessionConfigurationFile [-Path] <string> [-SchemaVersion <version>] [-Guid <guid>] [-Author <string>] [-Description <string>] [-CompanyName <string>] [-Copyright <string>] [-SessionType <SessionType>] [-TranscriptDirectory <string>] [-RunAsVirtualAccount] [-RunAsVirtualAccountGroups <string[]>] [-MountUserDrive] [-UserDriveMaximumSize <long>] [-GroupManagedServiceAccount <string>] [-ScriptsToProcess <string[]>] [-RoleDefinitions <IDictionary>] [-RequiredGroups <IDictionary>] [-LanguageMode <PSLanguageMode>] [-ExecutionPolicy <ExecutionPolicy>] [-PowerShellVersion <version>] [-ModulesToImport <Object[]>] [-VisibleAliases <string[]>] [-VisibleCmdlets <Object[]>] [-VisibleFunctions <Object[]>] [-VisibleExternalCommands <string[]>] [-VisibleProviders <string[]>] [-AliasDefinitions <IDictionary[]>] [-FunctionDefinitions <IDictionary[]>] [-VariableDefinitions <Object>] [-EnvironmentVariables <IDictionary>] [-TypesToProcess <string[]>] [-FormatsToProcess <string[]>] [-AssembliesToLoad <string[]>] [-Full] [<CommonParameters>]
```

Example: Creating and using a NoLanguage session

```powershell
New-PSSessionConfigurationFile -Path .\NoLanguage.pssc -LanguageMode NoLanguage
Register-PSSessionConfiguration -Path .\NoLanguage.pssc -Name NoLanguage -Force
$NoLanguageSession = New-PSSession -ComputerName Srv01 -ConfigurationName NoLanguage
Invoke-Command -Session $NoLanguageSession -ScriptBlock {
  if ((Get-Date) -lt '1January2099') {'Before'} else {'After'}
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-PSSessionConfigurationFile.md)


### New-PSSessionOption

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
New-PSSessionOption [-MaximumRedirection <int>] [-NoCompression] [-NoMachineProfile] [-Culture <cultureinfo>] [-UICulture <cultureinfo>] [-MaximumReceivedDataSizePerCommand <int>] [-MaximumReceivedObjectSize <int>] [-OutputBufferingMode <OutputBufferingMode>] [-MaxConnectionRetryCount <int>] [-ApplicationArguments <psprimitivedictionary>] [-OpenTimeout <int>] [-CancelTimeout <int>] [-IdleTimeout <int>] [-ProxyAccessType <ProxyAccessType>] [-ProxyAuthentication <AuthenticationMechanism>] [-ProxyCredential <pscredential>] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-OperationTimeout <int>] [-NoEncryption] [-UseUTF16] [-IncludePortInSPN] [<CommonParameters>]
```

Example: Create a default session option

```powershell
New-PSSessionOption
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-PSSessionOption.md)


### New-PSTransportOption

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
New-PSTransportOption [-MaxIdleTimeoutSec <int>] [-ProcessIdleTimeoutSec <int>] [-MaxSessions <int>] [-MaxConcurrentCommandsPerSession <int>] [-MaxSessionsPerUser <int>] [-MaxMemoryPerSessionMB <int>] [-MaxProcessesPerSession <int>] [-MaxConcurrentUsers <int>] [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [<CommonParameters>]
```

Example: Generate a default transport option

```powershell
New-PSTransportOption
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/New-PSTransportOption.md)


### New-TemporaryFile

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
New-TemporaryFile [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a temporary file

```powershell
$TempFile = New-TemporaryFile
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-TemporaryFile.md) / [Official reference source (5.1)](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/New-TemporaryFile.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 7.

- Type: Go implementation.
- Function: creates temporary files. `mktemp` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Extension` | string | File suffix |

- Implementation: `os.CreateTemp("", "tmp*"+suffix)`.
- Output: a FileInfo object.


### New-TimeSpan

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
New-TimeSpan [[-Start] <datetime>] [[-End] <datetime>] [<CommonParameters>]
New-TimeSpan [-Days <int>] [-Hours <int>] [-Minutes <int>] [-Seconds <int>] [<CommonParameters>]
```

Syntax (7):

```powershell
New-TimeSpan [[-Start] <datetime>] [[-End] <datetime>] [<CommonParameters>]
New-TimeSpan [-Days <int>] [-Hours <int>] [-Minutes <int>] [-Seconds <int>] [-Milliseconds <int>] [<CommonParameters>]
```

Example: Create a TimeSpan object for a specified duration

```powershell
$TimeSpan = New-TimeSpan -Hours 1 -Minutes 25
$TimeSpan
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-TimeSpan.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: constructs time spans.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Seconds` / `-Minutes` / `-Hours` | int | Values per unit, combinable |

- Output: a TimeSpan object with fields Days/Hours/Minutes/Seconds/TotalMilliseconds/TotalSeconds (plus TotalMinutes/TotalHours).


### New-Variable

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
New-Variable [-Name] <string> [[-Value] <Object>] [-Description <string>] [-Option <ScopedItemOptions>] [-Visibility <SessionStateEntryVisibility>] [-Force] [-PassThru] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a variable

```powershell
New-Variable days
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/New-Variable.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Variable, New-Variable, Remove-Variable, Clear-Variable.
- Function: sets/creates/removes/empties variables.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable name |
| `-Value` (position 1) | object | Value (New-Variable's `-Value` also takes position 1) |
| `-Force` | switch | Lets New-Variable overwrite an existing variable |

- Behavior: creating over an existing variable without -Force → error; assigning to read-only automatic variables (PID etc.) → error with Set/New, silently ignored with Clear. Remove deletes straight from the map; Clear sets $null.


### Out-Default

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Out-Default [-Transcript] [-InputObject <psobject>] [<CommonParameters>]
```

Example: 

```powershell
Get-Process | Select-Object -First 5 | Out-Default
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Out-Default.md)


### Out-File

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Out-File [-FilePath] <string> [[-Encoding] <string>] [-Append] [-Force] [-NoClobber] [-Width <int>] [-NoNewline] [-InputObject <psobject>] [-WhatIf] [-Confirm] [<CommonParameters>]
Out-File [[-Encoding] <string>] -LiteralPath <string> [-Append] [-Force] [-NoClobber] [-Width <int>] [-NoNewline] [-InputObject <psobject>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Out-File [-FilePath] <string> [[-Encoding] <Encoding>] [-Append] [-Force] [-NoClobber] [-Width <int>] [-NoNewline] [-InputObject <psobject>] [-WhatIf] [-Confirm] [<CommonParameters>]
Out-File [[-Encoding] <Encoding>] -LiteralPath <string> [-Append] [-Force] [-NoClobber] [-Width <int>] [-NoNewline] [-InputObject <psobject>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Send output and create a file

```powershell
Get-Process | Out-File -FilePath .\Process.txt
Get-Content -Path .\Process.txt
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Out-File.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: writes files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilePath` (position 0) | path | Target file |
| `-Append` | switch | Append mode (`>>`) |
| `-Encoding` | string | Encoding: same set as Set-Content (utf8 default, utf8BOM, ascii, unicode, etc.); no BOM in append mode |

- Implementation: formats input objects and writes them out (overwrite or append). Equivalent to `> file` / `>> file`.


### Out-Host

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Out-Host [-Paging] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Display output one page at a time

```powershell
Get-Process | Out-Host -Paging
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Out-Host.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: `-Paging` is not supported.

- Type: Go implementation.
- Function: outputs to the screen (the default behavior).

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to output |
- Implementation: formats input objects and writes to stdout.


### Out-Null

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Out-Null [-InputObject <psobject>] [<CommonParameters>]
```

Example: Delete output

```powershell
Get-ChildItem | Out-Null
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Out-Null.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: discards output. `> /dev/null` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Input object, discarded |


### Out-String

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Out-String [-Stream] [-Width <int>] [-InputObject <psobject>] [<CommonParameters>]
```

Syntax (7):

```powershell
Out-String [-Width <int>] [-NoNewline] [-InputObject <psobject>] [<CommonParameters>]
Out-String [-Stream] [-Width <int>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Get the current culture and convert the data to strings

```powershell
$C = Get-Culture | Select-Object -Property *
Out-String -InputObject $C -Width 100
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Out-String.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: `-Width` / `-NoNewline` are not supported.

- Type: Go implementation.
- Function: formats objects into a string.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Object to stringify |
| `-Stream` | switch | One line per object (`command | Out-String -Stream` ≈ verbatim text); otherwise one whole string |

- Implementation: reuses the object formatter writing into a buffer.


### Pop-Location

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Pop-Location [-PassThru] [-StackName <string>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Pop-Location [-PassThru] [-StackName <string>] [<CommonParameters>]
```

Example: Change to most recent location

```powershell
PS C:\> Pop-Location
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Pop-Location.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: The `pushd` / `popd` aliases are not implemented; use the full names.

- Type: Go implementation.
- Companion: Push-Location.
- Function: pops the stack back to the previous directory.
- Difference from Windows: the `popd` alias is not implemented; use the full name.
- Parameters: none.
- Behavior: empty stack means nothing happens. Bash's `popd` counterpart.


### Protect-CmsMessage

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Protect-CmsMessage [-To] <CmsMessageRecipient[]> [-Content] <psobject> [[-OutFile] <string>] [<CommonParameters>]
Protect-CmsMessage [-To] <CmsMessageRecipient[]> [-Path] <string> [[-OutFile] <string>] [<CommonParameters>]
Protect-CmsMessage [-To] <CmsMessageRecipient[]> [-LiteralPath] <string> [[-OutFile] <string>] [<CommonParameters>]
```

Example: Create a certificate for encrypting content

```powershell
# Create .INF file for certreq
{[Version]
Signature = "$Windows NT$"

[Strings]
szOID_ENHANCED_KEY_USAGE = "2.5.29.37"
szOID_DOCUMENT_ENCRYPTION = "1.3.6.1.4.1.311.80.1"

[NewRequest]
Subject = "cn=youralias@emailaddress.com"
MachineKeySet = false
KeyLength = 2048
KeySpec = AT_KEYEXCHANGE
HashAlgorithm = Sha1
Exportable = true
RequestType = Cert
KeyUsage = "CERT_KEY_ENCIPHERMENT_KEY_USAGE | CERT_DATA_ENCIPHERMENT_KEY_USAGE"
ValidityPeriod = "Years"
ValidityPeriodUnits = "1000"

[Extensions]
%szOID_ENHANCED_KEY_USAGE% = "{text}%szOID_DOCUMENT_ENCRYPTION%"
} | Out-File -FilePath DocumentEncryption.inf

# After you have created your certificate file, run the following command to add
# the certificate file to the certificate store. Now you are ready to encrypt and
# decrypt content with the next two examples.
certreq.exe -new DocumentEncryption.inf DocumentEncryption.cer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Protect-CmsMessage.md)


### Publish-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Publish-PSResource [-Path] <string> [-ApiKey <string>] [-Repository <string>] [-DestinationPath <string>] [-Credential <pscredential>] [-SkipDependenciesCheck] [-SkipModuleManifestValidate] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ModulePrefix <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Publish-PSResource -NupkgPath <string> [-ApiKey <string>] [-Repository <string>] [-DestinationPath <string>] [-Credential <pscredential>] [-SkipDependenciesCheck] [-SkipModuleManifestValidate] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-ModulePrefix <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Publish-PSResource -Path c:\TestModule
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/publish-psresource?view=powershell-7.5)


### Push-Location

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Push-Location [[-Path] <string>] [-PassThru] [-StackName <string>] [-UseTransaction] [<CommonParameters>]
Push-Location [-LiteralPath <string>] [-PassThru] [-StackName <string>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Push-Location [[-Path] <string>] [-PassThru] [-StackName <string>] [<CommonParameters>]
Push-Location [-LiteralPath <string>] [-PassThru] [-StackName <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Push-Location C:\Windows
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Push-Location.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: The `pushd` / `popd` aliases are not implemented; use the full names.

- Type: Go implementation.
- Companion: Pop-Location.
- Function: pushes the current directory onto a stack and switches.
- Difference from Windows: the `pushd` alias is not implemented; use the full name.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Directory to switch to; omitted, only the push happens with no switch |

- Behavior: nonexistent directory → error. Bash's `pushd` counterpart.


### Read-Host

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Read-Host [[-Prompt] <Object>] [-AsSecureString] [<CommonParameters>]
```

Syntax (7):

```powershell
Read-Host [[-Prompt] <Object>] [-MaskInput] [<CommonParameters>]
Read-Host [[-Prompt] <Object>] [-AsSecureString] [<CommonParameters>]
```

Example: Save console input to a variable

```powershell
$Age = Read-Host "Please enter your age"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Read-Host.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: reads one line of input. Bash's `read`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Prompt` (position 0) | string | Prompt text |

- Implementation: prints the prompt then reads a line with bufio, trimming the trailing line break.


### Receive-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Receive-Job [-Job] <Job[]> [[-Location] <string[]>] [-Keep] [-NoRecurse] [-Force] [-Wait] [-AutoRemoveJob] [-WriteEvents] [-WriteJobInResults] [<CommonParameters>]
Receive-Job [-Job] <Job[]> [[-ComputerName] <string[]>] [-Keep] [-NoRecurse] [-Force] [-Wait] [-AutoRemoveJob] [-WriteEvents] [-WriteJobInResults] [<CommonParameters>]
Receive-Job [-Job] <Job[]> [[-Session] <PSSession[]>] [-Keep] [-NoRecurse] [-Force] [-Wait] [-AutoRemoveJob] [-WriteEvents] [-WriteJobInResults] [<CommonParameters>]
Receive-Job [-Name] <string[]> [-Keep] [-NoRecurse] [-Force] [-Wait] [-AutoRemoveJob] [-WriteEvents] [-WriteJobInResults] [<CommonParameters>]
Receive-Job [-InstanceId] <guid[]> [-Keep] [-NoRecurse] [-Force] [-Wait] [-AutoRemoveJob] [-WriteEvents] [-WriteJobInResults] [<CommonParameters>]
Receive-Job [-Id] <int[]> [-Keep] [-NoRecurse] [-Force] [-Wait] [-AutoRemoveJob] [-WriteEvents] [-WriteJobInResults] [<CommonParameters>]
```

Example: Get results for a particular job

```powershell
$job = Start-Job -ScriptBlock {Get-Process}
Start-Sleep -Seconds 1
Receive-Job -Job $job
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Receive-Job.md)


### Register-ArgumentCompleter

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Register-ArgumentCompleter -ParameterName <string> -ScriptBlock <scriptblock> [-CommandName <string[]>] [<CommonParameters>]
Register-ArgumentCompleter -CommandName <string[]> -ScriptBlock <scriptblock> [-Native] [<CommonParameters>]
```

Syntax (7):

```powershell
Register-ArgumentCompleter -CommandName <string[]> -ScriptBlock <scriptblock> [-Native] [<CommonParameters>]
Register-ArgumentCompleter -ParameterName <string> -ScriptBlock <scriptblock> [-CommandName <string[]>] [<CommonParameters>]
```

Example: Register a custom argument completer

```powershell
$s = {
    param(
        $commandName,
        $parameterName,
        $wordToComplete,
        $commandAst,
        $fakeBoundParameters
    )

    (Get-TimeZone -ListAvailable).Id | Where-Object {
        $_ -like "$wordToComplete*"
    } | ForEach-Object {
        "'$_'"
    }
}

Register-ArgumentCompleter -CommandName Set-TimeZone -ParameterName Id -ScriptBlock $s
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Register-ArgumentCompleter.md)


### Register-EngineEvent

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Register-EngineEvent [-SourceIdentifier] <string> [[-Action] <scriptblock>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

Example: Register a PowerShell engine event on remote computers

```powershell
$S = New-PSSession -ComputerName "Server01, Server02"
Invoke-Command -Session $S {
  Register-EngineEvent -SourceIdentifier ([System.Management.Automation.PSEngineEvent]::Exiting) -Forward
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Register-EngineEvent.md)


### Register-ObjectEvent

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Register-ObjectEvent [-InputObject] <psobject> [-EventName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

Example: Subscribe to events when a new process starts

```powershell
$queryParameters = '__InstanceCreationEvent', (New-Object TimeSpan 0,0,1),
    "TargetInstance isa 'Win32_Process'"
$Query = New-Object System.Management.WqlEventQuery -ArgumentList $queryParameters
$ProcessWatcher = New-Object System.Management.ManagementEventWatcher $Query
Register-ObjectEvent -InputObject $ProcessWatcher -EventName "EventArrived"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Register-ObjectEvent.md)


### Register-PackageSource

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Register-PackageSource [[-Name] <string>] [[-Location] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string>] [<CommonParameters>]
Register-PackageSource [[-Name] <string>] [[-Location] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Register-PackageSource [[-Name] <string>] [[-Location] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string>] [<CommonParameters>]
Register-PackageSource [[-Name] <string>] [[-Location] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Register-PackageSource [[-Name] <string>] [[-Location] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Example: Register a package source for the NuGet provider

```powershell
Register-PackageSource -Name MyNuGet -Location https://www.nuget.org/api/v2 -ProviderName NuGet
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/register-packagesource?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/register-packagesource?view=powershell-7.5)


### Register-PSResourceRepository

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Register-PSResourceRepository [-Name] <string> [-Uri] <string> [-Trusted] [-Priority <int>] [-ApiVersion <PSRepositoryInfo+APIVersion>] [-CredentialInfo <PSCredentialInfo>] [-CredentialProvider <CredentialProvider>] [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSResourceRepository -PSGallery [-Trusted] [-Priority <int>] [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSResourceRepository -MicrosoftArtifactRegistry [-Trusted] [-Priority <int>] [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSResourceRepository -Repository <hashtable[]> [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Register-PSResourceRepository -Name PoshTestGallery -Uri 'https://www.poshtestgallery.com/api/v2'
Get-PSResourceRepository -Name PoshTestGallery
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/register-psresourcerepository?view=powershell-7.5)


### Remove-Alias

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Remove-Alias [-Name] <string[]> [-Scope <string>] [-Force] [<CommonParameters>]
```

Example: Remove an alias

```powershell
Remove-Alias -Name del
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Remove-Alias.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Alias, New-Alias, Remove-Alias, Import-Alias, Export-Alias.
- Function: manages aliases.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Alias name |

- Set-Alias overwrites; New-Alias errors on an existing alias without -Force; Remove-Alias deletes; Export-Alias writes one `name=target` per line; Import-Alias reads them back.


### Remove-Event

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Remove-Event [-SourceIdentifier] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Event [-EventIdentifier] <int> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove an event by source identifier

```powershell
PS C:\> Remove-Event -SourceIdentifier "ProcessStarted"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Remove-Event.md)


### Remove-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Remove-Item [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-Stream <string[]>] [<CommonParameters>]
Remove-Item -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-Stream <string[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-Item [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-Stream <string[]>] [<CommonParameters>]
Remove-Item -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Recurse] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-Stream <string[]>] [<CommonParameters>]
```

Example: Delete files that have any file extension

```powershell
Remove-Item C:\Test\*.*
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Remove-Item.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Nonexistent paths are silently ignored (PowerShell errors with path-not-found).

- Type: Go implementation.
- Function: deletes files/directories.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target paths, multiple/wildcards allowed |
| `-Recurse` | switch | Removes a whole directory tree, `rm -r` |
| `-Force` | switch | Forces removal, `rm -f` (files also go through RemoveAll) |

- Behavior: nonexistent path → ignored silently.


### Remove-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Remove-ItemProperty [-Path] <string[]> [-Name] <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Remove-ItemProperty [-Name] <string[]> -LiteralPath <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-ItemProperty [-Path] <string[]> [-Name] <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-ItemProperty [-Name] <string[]> -LiteralPath <string[]> [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Delete a registry value

```powershell
Remove-ItemProperty -Path "HKLM:\Software\SmpApplication" -Name "SmpProperty"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Remove-ItemProperty.md)


### Remove-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Remove-Job [-Id] <int[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Job] <Job[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Name] <string[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-InstanceId] <guid[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Filter] <hashtable> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-State] <JobState> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Command <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-Job [-Id] <int[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Job] <Job[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-InstanceId] <guid[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Name] <string[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Filter] <hashtable> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-State] <JobState> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Job [-Command <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Delete a job by using its name

```powershell
$batch = Get-Job -Name BatchJob
$batch | Remove-Job
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Remove-Job.md)


### Remove-Module

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Remove-Module [-Name] <string[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Module [-FullyQualifiedName] <ModuleSpecification[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Module [-ModuleInfo] <psmoduleinfo[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a module

```powershell
Remove-Module -Name "BitsTransfer"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Remove-Module.md)


### Remove-PSBreakpoint

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Remove-PSBreakpoint [-Breakpoint] <Breakpoint[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSBreakpoint [-Id] <int[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-PSBreakpoint [-Breakpoint] <Breakpoint[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSBreakpoint [-Id] <int[]> [-Runspace <runspace>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove all breakpoints

```powershell
Get-PSBreakpoint | Remove-PSBreakpoint
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Remove-PSBreakpoint.md)


### Remove-PSDrive

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Remove-PSDrive [-Name] <string[]> [-PSProvider <string[]>] [-Scope <string>] [-Force] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Remove-PSDrive [-LiteralName] <string[]> [-PSProvider <string[]>] [-Scope <string>] [-Force] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-PSDrive [-Name] <string[]> [-PSProvider <string[]>] [-Scope <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSDrive [-LiteralName] <string[]> [-PSProvider <string[]>] [-Scope <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a file system drive

```powershell
Remove-PSDrive -Name smp
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Remove-PSDrive.md)


### Remove-PSReadLineKeyHandler

Version: Both

Module: PSReadLine

Syntax:

```powershell
Remove-PSReadLineKeyHandler [-Chord] <string[]> [-ViMode <ViMode>] [<CommonParameters>]
```

Example: Remove a binding

```powershell
Remove-PSReadLineKeyHandler -Chord Ctrl+B
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/PSReadLine/Remove-PSReadLineKeyHandler.md)


### Remove-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Remove-PSSession [-Id] <int[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession [-Session] <PSSession[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession -ContainerId <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession -VMId <guid[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession -VMName <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession -InstanceId <guid[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession -Name <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PSSession [-ComputerName] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove sessions by ID

```powershell
Remove-PSSession -Id 1, 2
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Remove-PSSession.md)


### Remove-TypeData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Remove-TypeData -TypeData <TypeData> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-TypeData [-TypeName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-TypeData -Path <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove type data for a specified type

```powershell
Remove-TypeData -TypeName System.Array
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Remove-TypeData.md)


### Remove-Variable

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Remove-Variable [-Name] <string[]> [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a variable

```powershell
Remove-Variable Smp
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Remove-Variable.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Set-Variable, New-Variable, Remove-Variable, Clear-Variable.
- Function: sets/creates/removes/empties variables.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable name |
| `-Value` (position 1) | object | Value (New-Variable's `-Value` also takes position 1) |
| `-Force` | switch | Lets New-Variable overwrite an existing variable |

- Behavior: creating over an existing variable without -Force → error; assigning to read-only automatic variables (PID etc.) → error with Set/New, silently ignored with Clear. Remove deletes straight from the map; Clear sets $null.


### Rename-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Rename-Item [-Path] <string> [-NewName] <string> [-Force] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Rename-Item [-NewName] <string> -LiteralPath <string> [-Force] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Rename-Item [-Path] <string> [-NewName] <string> [-Force] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-Item [-NewName] <string> -LiteralPath <string> [-Force] [-PassThru] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Rename a file

```powershell
Rename-Item -Path "C:\logfiles\daily_file.txt" -NewName "monday_file.txt"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Rename-Item.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: renames files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Old path |
| `-NewName` (position 1) | string | New name; only the leaf name is taken, full paths accepted |

- Implementation: `os.Rename(old path, new leaf name in the same directory)`, matching bash `mv old new-name-in-same-directory`.


### Rename-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Rename-ItemProperty [-Path] <string> [-Name] <string> [-NewName] <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Rename-ItemProperty [-Name] <string> [-NewName] <string> -LiteralPath <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Rename-ItemProperty [-Path] <string> [-Name] <string> [-NewName] <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-ItemProperty [-Name] <string> [-NewName] <string> -LiteralPath <string> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Rename a registry entry

```powershell
Rename-ItemProperty -Path HKLM:\Software\SmpApplication -Name config -NewName oldconfig
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Rename-ItemProperty.md)


### Reset-PSResourceRepository

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Reset-PSResourceRepository [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Reset the repository store and display the results

```powershell
Reset-PSResourceRepository -PassThru
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/reset-psresourcerepository?view=powershell-7.5)


### Resolve-Path

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Resolve-Path [-Path] <string[]> [-Relative] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Resolve-Path -LiteralPath <string[]> [-Relative] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Resolve-Path [-Path] <string[]> [-Relative] [-RelativeBasePath <string>] [-Force] [-Credential <pscredential>] [<CommonParameters>]
Resolve-Path -LiteralPath <string[]> [-Relative] [-RelativeBasePath <string>] [-Force] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Resolve the home folder path

```powershell
Resolve-Path ~
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Resolve-Path.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Resolve-Path additionally resolves symbolic links.

- Type: Go implementation.
- Companion: Convert-Path.
- Function: converts to absolute paths. Bash's `realpath`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |

- Implementation: Resolve-Path additionally runs `EvalSymlinks` (resolves symbolic links); Convert-Path only cleans.


### Restart-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Restart-Computer [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-DcomAuthentication <AuthenticationLevel>] [-Impersonation <ImpersonationLevel>] [-WsmanAuthentication <string>] [-Protocol <string>] [-Force] [-Wait] [-Timeout <int>] [-For <WaitForServiceTypes>] [-Delay <int16>] [-WhatIf] [-Confirm] [<CommonParameters>]
Restart-Computer [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-AsJob] [-DcomAuthentication <AuthenticationLevel>] [-Impersonation <ImpersonationLevel>] [-Force] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Restart-Computer [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-WsmanAuthentication <string>] [-Force] [-Wait] [-Timeout <int>] [-For <WaitForServiceTypes>] [-Delay <short>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Restart the local computer

```powershell
Restart-Computer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Restart-Computer.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`sudo reboot`).
- Companions: Stop-Computer, Rename-Computer.
- Distro: needs sudo.
- Function: reboots the machine.
- Parameters: none.
- Implementation: `sudo reboot`.


### Save-Help

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Save-Help [-DestinationPath] <string[]> [[-Module] <psmoduleinfo[]>] [[-UICulture] <cultureinfo[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [<CommonParameters>]
Save-Help [[-Module] <psmoduleinfo[]>] [[-UICulture] <cultureinfo[]>] -LiteralPath <string[]> [-FullyQualifiedModule <ModuleSpecification[]>] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [<CommonParameters>]
```

Syntax (7):

```powershell
Save-Help [-DestinationPath] <string[]> [[-Module] <psmoduleinfo[]>] [[-UICulture] <cultureinfo[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [-Scope <UpdateHelpScope>] [<CommonParameters>]
Save-Help [[-Module] <psmoduleinfo[]>] [[-UICulture] <cultureinfo[]>] -LiteralPath <string[]> [-FullyQualifiedModule <ModuleSpecification[]>] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [-Scope <UpdateHelpScope>] [<CommonParameters>]
```

Example: Save the help for the DhcpServer module

```powershell
# Option 1:
# 1. Run Invoke-Command to get the PSModuleInfo object for the DhcpServer module,
# 2. Save-Help on the PSModuleInfo object to save the help files to a folder on
#    the local computer.

$mod = Invoke-Command -ComputerName RemoteServer -ScriptBlock {
    Get-Module -Name DhcpServer -ListAvailable
}
Save-Help -Module $mod -DestinationPath C:\SavedHelp


# Option 2:
# 1. Open a PSSession to the remote computer that's running the DhcpServer module
# 2. Get the PSModuleInfo object from the remote computer
# 3. Save-Help on the PSModuleInfo object

$session = New-PSSession -ComputerName "RemoteServer"
$mod = Get-Module -PSSession $session -Name "DhcpServer" -ListAvailable
Save-Help -Module $mod -DestinationPath C:\SavedHelp


# Option 3:
# 1. Open a CimSession to the remote computer that's running the DhcpServer module
# 2. Get the PSModuleInfo object from the remote computer
# 3. Save-Help on the PSModuleInfo object
$cimsession = New-CimSession -ComputerName "RemoteServer"
$mod = Get-Module -CimSession $cimsession -Name "DhcpServer" -ListAvailable
Save-Help -Module $mod -DestinationPath "C:\SavedHelp"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Save-Help.md)


### Save-Package

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Save-Package [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Source <string[]>] [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string[]>] [<CommonParameters>]
Save-Package -InputObject <SoftwareIdentity> [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Save-Package [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [<CommonParameters>]
Save-Package [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Save-Package [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-Source <string[]>] [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string[]>] [<CommonParameters>]
Save-Package -InputObject <SoftwareIdentity> [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Save-Package [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [-Headers <string[]>] [-FilterOnTag <string[]>] [-Contains <string>] [-AllowPrereleaseVersions] [<CommonParameters>]
Save-Package [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [-Headers <string[]>] [-FilterOnTag <string[]>] [-Contains <string>] [-AllowPrereleaseVersions] [<CommonParameters>]
Save-Package [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-AllowPrereleaseVersions] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-AcceptLicense] [<CommonParameters>]
Save-Package [-Path <string>] [-LiteralPath <string>] [-Credential <pscredential>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-AllowPrereleaseVersions] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [-Type <string>] [-Filter <string>] [-Tag <string[]>] [-Includes <string[]>] [-DscResource <string[]>] [-RoleCapability <string[]>] [-Command <string[]>] [-AcceptLicense] [<CommonParameters>]
```

Example: Save a package to the local computer

```powershell
PS> Save-Package -Name NuGet.Core -ProviderName NuGet -Path C:\LocalPkg
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/save-package?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/save-package?view=powershell-7.5)


### Save-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Save-PSResource [-Name] <string[]> [-Version <string>] [-Prerelease] [-Repository <string[]>] [-Credential <pscredential>] [-IncludeXml] [-Path <string>] [-TemporaryPath <string>] [-TrustRepository] [-PassThru] [-SkipDependencyCheck] [-AuthenticodeCheck] [-Quiet] [-AcceptLicense] [-WhatIf] [-Confirm] [<CommonParameters>]
Save-PSResource [-Name] <string[]> [-Version <string>] [-Prerelease] [-Repository <string[]>] [-Credential <pscredential>] [-AsNupkg] [-Path <string>] [-TemporaryPath <string>] [-TrustRepository] [-PassThru] [-SkipDependencyCheck] [-AuthenticodeCheck] [-Quiet] [-AcceptLicense] [-WhatIf] [-Confirm] [<CommonParameters>]
Save-PSResource [-InputObject] <PSResourceInfo[]> [-Repository <string[]>] [-Credential <pscredential>] [-AsNupkg] [-IncludeXml] [-Path <string>] [-TemporaryPath <string>] [-TrustRepository] [-PassThru] [-SkipDependencyCheck] [-AuthenticodeCheck] [-Quiet] [-AcceptLicense] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Save-PSResource -Name Az
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/save-psresource?view=powershell-7.5)


### Select-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Select-Object [[-Property] <Object[]>] [-InputObject <psobject>] [-ExcludeProperty <string[]>] [-ExpandProperty <string>] [-Unique] [-Last <int>] [-First <int>] [-Skip <int>] [-Wait] [<CommonParameters>]
Select-Object [[-Property] <Object[]>] [-InputObject <psobject>] [-ExcludeProperty <string[]>] [-ExpandProperty <string>] [-Unique] [-SkipLast <int>] [<CommonParameters>]
Select-Object [-InputObject <psobject>] [-Unique] [-Wait] [-Index <int[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Select-Object [[-Property] <Object[]>] [-InputObject <psobject>] [-ExcludeProperty <string[]>] [-ExpandProperty <string>] [-Unique] [-CaseInsensitive] [-Last <int>] [-First <int>] [-Skip <int>] [-Wait] [<CommonParameters>]
Select-Object [[-Property] <Object[]>] [-InputObject <psobject>] [-ExcludeProperty <string[]>] [-ExpandProperty <string>] [-Unique] [-CaseInsensitive] [-Skip <int>] [-SkipLast <int>] [<CommonParameters>]
Select-Object [-InputObject <psobject>] [-Unique] [-CaseInsensitive] [-Wait] [-Index <int[]>] [<CommonParameters>]
Select-Object [-InputObject <psobject>] [-Unique] [-CaseInsensitive] [-SkipIndex <int[]>] [<CommonParameters>]
```

Example: Select objects by property

```powershell
Get-Process | Select-Object -Property ProcessName, Id, WS
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Select-Object.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original.

- Type: Go implementation.
- Function: picks properties, leading/trailing entries, or skips entries. Bash's `cut`, `head`, `tail`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Keeps only these properties (producing new objects); `*` means all |
| `-ExpandProperty` | string | Outputs that property's own value without wrapping it in an object; array values flatten (as in `Select-Object -ExpandProperty Name`) |
| `-First` | int | First N entries (`head -n N`) |
| `-Last` | int | Last N entries (`tail -n N`) |
| `-Skip` | int | Skips N entries first (`tail -n +N+1`); skips from the end when `-Last` is present; negative values error out |
| `-Unique` | switch | Deduplicates (`sort -u`, by string, case-sensitive) |

- Behavior: no input with an array at position 0 → treated by array element. `-First 0` / `-Last 0` return empty (explicit 0 is distinct from unset). `-Unique` deduplicates after projection/expansion, by resulting values.


### Select-String

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Select-String [-Pattern] <string[]> [-Path] <string[]> [-SimpleMatch] [-CaseSensitive] [-Quiet] [-List] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <string>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> -InputObject <psobject> [-SimpleMatch] [-CaseSensitive] [-Quiet] [-List] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <string>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> -LiteralPath <string[]> [-SimpleMatch] [-CaseSensitive] [-Quiet] [-List] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <string>] [-Context <int[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Select-String [-Pattern] <string[]> [-Path] <string[]> [-Culture <string>] [-SimpleMatch] [-CaseSensitive] [-Quiet] [-List] [-NoEmphasis] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <Encoding>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> -InputObject <psobject> [-Culture <string>] [-SimpleMatch] [-CaseSensitive] [-Quiet] [-List] [-NoEmphasis] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <Encoding>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> -InputObject <psobject> -Raw [-Culture <string>] [-SimpleMatch] [-CaseSensitive] [-List] [-NoEmphasis] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <Encoding>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> [-Path] <string[]> -Raw [-Culture <string>] [-SimpleMatch] [-CaseSensitive] [-List] [-NoEmphasis] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <Encoding>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> -LiteralPath <string[]> [-Culture <string>] [-SimpleMatch] [-CaseSensitive] [-Quiet] [-List] [-NoEmphasis] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <Encoding>] [-Context <int[]>] [<CommonParameters>]
Select-String [-Pattern] <string[]> -LiteralPath <string[]> -Raw [-Culture <string>] [-SimpleMatch] [-CaseSensitive] [-List] [-NoEmphasis] [-Include <string[]>] [-Exclude <string[]>] [-NotMatch] [-AllMatches] [-Encoding <Encoding>] [-Context <int[]>] [<CommonParameters>]
```

Example: Find a case-sensitive match

```powershell
'Hello', 'HELLO' | Select-String -Pattern 'HELLO' -CaseSensitive -SimpleMatch
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Select-String.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Invalid regex does not error (it simply does not match).

- Type: Go implementation.
- Function: finds matching lines by regex. Bash's `grep -E`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Pattern` (position 0) | string | Regex pattern |
| `-Path` (position 1) | path | File to search |
| `-SimpleMatch` | switch | Literal matching, `grep -F` |
| `-CaseSensitive` | switch | Case-sensitive matching; insensitive by default (as with `grep -i` reversed) |
| `-Quiet` | switch | Reports only whether anything matched: a single True on match, no output otherwise (`grep -q`) |

- Without -Path, searches pipeline input.
- Output: MatchInfo objects with fields LineNumber (from 1), Line, Path, Pattern.
- Behavior: case-insensitive by default; invalid regexes → no match (no error raised).
- LineNumber: file input scans line by line with blank lines counted; each pipeline input object participates in matching as a single line, its LineNumber being the object's ordinal in the stream — non-matching objects occupy numbers too.


### Select-Xml

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Select-Xml [-XPath] <string> [-Xml] <XmlNode[]> [-Namespace <hashtable>] [<CommonParameters>]
Select-Xml [-XPath] <string> [-Path] <string[]> [-Namespace <hashtable>] [<CommonParameters>]
Select-Xml [-XPath] <string> -LiteralPath <string[]> [-Namespace <hashtable>] [<CommonParameters>]
Select-Xml [-XPath] <string> -Content <string[]> [-Namespace <hashtable>] [<CommonParameters>]
```

Example: Select AliasProperty nodes

```powershell
$Path = "$PSHOME\Types.ps1xml"
$XPath = "/Types/Type/Members/AliasProperty"
Select-Xml -Path $Path -XPath $Xpath | Select-Object -ExpandProperty Node
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Select-Xml.md)


### Send-MailMessage

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Send-MailMessage [-To] <string[]> [-Subject] <string> [[-Body] <string>] [[-SmtpServer] <string>] -From <string> [-Attachments <string[]>] [-Bcc <string[]>] [-BodyAsHtml] [-Encoding <Encoding>] [-Cc <string[]>] [-DeliveryNotificationOption <DeliveryNotificationOptions>] [-Priority <MailPriority>] [-Credential <pscredential>] [-UseSsl] [-Port <int>] [<CommonParameters>]
```

Syntax (7):

```powershell
Send-MailMessage [-To] <string[]> [[-Subject] <string>] [[-Body] <string>] [[-SmtpServer] <string>] -From <string> [-Attachments <string[]>] [-Bcc <string[]>] [-BodyAsHtml] [-Encoding <Encoding>] [-Cc <string[]>] [-DeliveryNotificationOption <DeliveryNotificationOptions>] [-Priority <MailPriority>] [-ReplyTo <string[]>] [-Credential <pscredential>] [-UseSsl] [-Port <int>] [<CommonParameters>]
```

Example: Send an email from one person to another person

```powershell
$sendMailMessageSplat = @{
    From = 'User01 <user01@fabrikam.com>'
    To = 'User02 <user02@fabrikam.com>'
    Subject = 'Test mail'
}
Send-MailMessage @sendMailMessageSplat
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Send-MailMessage.md)


### Set-Alias

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Set-Alias [-Name] <string> [-Value] <string> [-Description <string>] [-Option <ScopedItemOptions>] [-PassThru] [-Scope <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create an alias for a cmdlet

```powershell
PS> Set-Alias -Name list -Value Get-ChildItem

PS> Get-Alias -Name list

CommandType     Name
-----------     ----
Alias           list -> Get-ChildItem
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Set-Alias.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Companions: Set-Alias, New-Alias, Remove-Alias, Import-Alias, Export-Alias.
- Function: manages aliases.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Alias name |
| `-Value` (position 1) | string | Target command |

- Set-Alias overwrites; New-Alias errors on an existing alias without -Force; Remove-Alias deletes; Export-Alias writes one `name=target` per line; Import-Alias reads them back.


### Set-Clipboard

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Set-Clipboard [-Append] [-AsHtml] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Clipboard [-Value] <string[]> [-Append] [-AsHtml] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Clipboard -Path <string[]> [-Append] [-AsHtml] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Clipboard -LiteralPath <string[]> [-Append] [-AsHtml] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Clipboard [-Value] <string[]> [-Append] [-PassThru] [-AsOSC52] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Copy text to the clipboard

```powershell
Set-Clipboard -Value "This is a test string"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-Clipboard.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`xclip` / `xsel`).
- Companion: Get-Clipboard.
- Distro: X11 + xclip/xsel.
- Function: writes the clipboard.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Value` (position 0) | object | Text to write; pipeline input takes precedence |

- Implementation: `xclip -selection clipboard` first, otherwise `xsel -b`; neither present, quietly does nothing.


### Set-Content

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Set-Content [-Path] <string[]> [-Value] <Object[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-NoNewline] [-Encoding <FileSystemCmdletProviderEncoding>] [-Stream <string>] [<CommonParameters>]
Set-Content [-Value] <Object[]> -LiteralPath <string[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [-NoNewline] [-Encoding <FileSystemCmdletProviderEncoding>] [-Stream <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Content [-Path] <string[]> [-Value] <Object[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-NoNewline] [-Encoding <Encoding>] [-AsByteStream] [-Stream <string>] [<CommonParameters>]
Set-Content [-Value] <Object[]> -LiteralPath <string[]> [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Force] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-NoNewline] [-Encoding <Encoding>] [-AsByteStream] [-Stream <string>] [<CommonParameters>]
```

Example: Replace the contents of multiple files in a directory

```powershell
Get-ChildItem -Path .\Test*.txt
```

```powershell
Set-Content -Path .\Test*.txt -Value 'Hello, World'
Get-Content -Path .\Test*.txt
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-Content.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: overwrites files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target file |
| `-Value` (position 1) | object | Content to write; pipeline input takes precedence over it |
| `-Encoding` | string | Encoding: utf8 (default, no BOM) / utf8BOM / utf8NoBOM / ascii / unicode (UTF-16LE) / bigendianunicode / utf32 / bigendianutf32; unknown names treated as UTF-8 |

- Implementation: writes each object's string into the file line by line (overwrite), `echo ... > file` / `tee` territory. Line endings added automatically.


### Set-Date

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Set-Date [-Date] <datetime> [-DisplayHint <DisplayHintType>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Date [-Adjust] <timespan> [-DisplayHint <DisplayHintType>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Add three days to the system date

```powershell
Set-Date -Date (Get-Date).AddDays(3)
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Set-Date.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`sudo date -s`).
- Distro: needs sudo.
- Function: sets system time.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Date` (position 0) | string | Time string; handed to `sudo date -s`; omitted, the current time comes back |


### Set-ExecutionPolicy

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Set-ExecutionPolicy [-ExecutionPolicy] <ExecutionPolicy> [[-Scope] <ExecutionPolicyScope>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set an execution policy

```powershell
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope LocalMachine
Get-ExecutionPolicy -List
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Set-ExecutionPolicy.md)


### Set-Item

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Set-Item [-Path] <string[]> [[-Value] <Object>] [-Force] [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-Item [[-Value] <Object>] -LiteralPath <string[]> [-Force] [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Item [-Path] <string[]> [[-Value] <Object>] [-Force] [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Item [[-Value] <Object>] -LiteralPath <string[]> [-Force] [-PassThru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create an alias

```powershell
Set-Item -Path Alias:np -Value "C:\windows\notepad.exe"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-Item.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Set-Item accepts the `env:` prefix to set environment variables.

- Type: Go implementation.
- Companion: Clear-Item.
- Function: Set-Item writes files or sets environment variables; Clear-Item empties files.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path; beginning with `env:` it's read as an environment variable name |
| `-Value` (position 1) | object | Content to write |

- Behavior: `env:NAME` → `os.Setenv(NAME, value)`, `export NAME=value` territory. Ordinary paths → writes the file.


### Set-ItemProperty

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Set-ItemProperty [-Path] <string[]> [-Name] <string> [-Value] <Object> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-ItemProperty [-Path] <string[]> -InputObject <psobject> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-ItemProperty [-Name] <string> [-Value] <Object> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-ItemProperty -LiteralPath <string[]> -InputObject <psobject> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-ItemProperty [-Path] <string[]> [-Name] <string> [-Value] <Object> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-ItemProperty [-Path] <string[]> -InputObject <psobject> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-ItemProperty [-Name] <string> [-Value] <Object> -LiteralPath <string[]> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-ItemProperty -LiteralPath <string[]> -InputObject <psobject> [-PassThru] [-Force] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set a property of a file

```powershell
Set-ItemProperty -Path C:\GroupFiles\final.doc -Name IsReadOnly -Value $true
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-ItemProperty.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Only modification time can be changed; other properties are ignored.

- Type: Go implementation.
- Function: modifies file attributes. Currently supports modification time only.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |
| `-Name` (position 1) | string | Property name; `LastWriteTime` sets the modification time to now (bash's `touch`) |
| `-Value` (position 2) | object | Property value; currently ignored, other property names likewise ignored (without error) |

- Behavior: missing file → error.


### Set-Location

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Set-Location [[-Path] <string>] [-PassThru] [-UseTransaction] [<CommonParameters>]
Set-Location -LiteralPath <string> [-PassThru] [-UseTransaction] [<CommonParameters>]
Set-Location [-PassThru] [-StackName <string>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Location [[-Path] <string>] [-PassThru] [<CommonParameters>]
Set-Location -LiteralPath <string> [-PassThru] [<CommonParameters>]
Set-Location [-PassThru] [-StackName <string>] [<CommonParameters>]
```

Example: Set the current location

```powershell
PS C:\> Set-Location -Path "HKLM:\"
PS HKLM:\>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-Location.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: changes directory.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target directory. Accepts relative paths, `..`, `~`/`~/...` (home). Bash's `cd` equivalent |

- Behavior: nonexistent or non-directory target → error onto stderr, $?=false. On success the current directory changes; bare invocation returns to the home directory.


### Set-MarkdownOption

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Set-MarkdownOption [-Header1Color <string>] [-Header2Color <string>] [-Header3Color <string>] [-Header4Color <string>] [-Header5Color <string>] [-Header6Color <string>] [-Code <string>] [-ImageAltTextForegroundColor <string>] [-LinkForegroundColor <string>] [-ItalicsForegroundColor <string>] [-BoldForegroundColor <string>] [-PassThru] [<CommonParameters>]
Set-MarkdownOption -Theme <string> [-PassThru] [<CommonParameters>]
Set-MarkdownOption [-InputObject] <psobject> [-PassThru] [<CommonParameters>]
```

Example: Switch to the Light Theme

```powershell
Set-MarkdownOption -Theme Light -PassThru
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Set-MarkdownOption.md)


### Set-PackageSource

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Set-PackageSource [[-Name] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Location <string>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string>] [<CommonParameters>]
Set-PackageSource -InputObject <PackageSource> [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PackageSource [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
Set-PackageSource [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-PackageSource [[-Name] <string>] [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-Location <string>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string>] [<CommonParameters>]
Set-PackageSource -InputObject <PackageSource> [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PackageSource [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Set-PackageSource [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Set-PackageSource [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
Set-PackageSource [-Proxy <uri>] [-ProxyCredential <pscredential>] [-Credential <pscredential>] [-NewLocation <string>] [-NewName <string>] [-Trusted] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Example: Change a package source

```powershell
PS C:\> Set-PackageSource -Name MyNuget -NewName NewNuGet -Trusted -ProviderName NuGet
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/set-packagesource?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/set-packagesource?view=powershell-7.5)


### Set-PSBreakpoint

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Set-PSBreakpoint [-Script] <string[]> [-Line] <int[]> [[-Column] <int>] [-Action <scriptblock>] [<CommonParameters>]
Set-PSBreakpoint [[-Script] <string[]>] -Variable <string[]> [-Action <scriptblock>] [-Mode <VariableAccessMode>] [<CommonParameters>]
Set-PSBreakpoint [[-Script] <string[]>] -Command <string[]> [-Action <scriptblock>] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-PSBreakpoint [-Script] <string[]> [-Line] <int[]> [[-Column] <int>] [-Action <scriptblock>] [-Runspace <runspace>] [<CommonParameters>]
Set-PSBreakpoint [[-Script] <string[]>] -Command <string[]> [-Action <scriptblock>] [-Runspace <runspace>] [<CommonParameters>]
Set-PSBreakpoint [[-Script] <string[]>] -Variable <string[]> [-Action <scriptblock>] [-Mode <VariableAccessMode>] [-Runspace <runspace>] [<CommonParameters>]
```

Example: Set a breakpoint on a line

```powershell
Set-PSBreakpoint -Script "sample.ps1" -Line 5
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Set-PSBreakpoint.md)


### Set-PSDebug

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Set-PSDebug [-Trace <int>] [-Step] [-Strict] [<CommonParameters>]
Set-PSDebug [-Off] [<CommonParameters>]
```

Example: Set the trace level

```powershell
Set-PSDebug -Trace 2; foreach ($i in 1..3) {$i}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Set-PSDebug.md)


### Set-PSReadLineKeyHandler

Version: Both

Module: PSReadLine

Syntax:

```powershell
Set-PSReadLineKeyHandler [-Chord] <string[]> [-ScriptBlock] <scriptblock> [-BriefDescription <string>] [-Description <string>] [-ViMode <ViMode>] [<CommonParameters>]
Set-PSReadLineKeyHandler [-Chord] <string[]> [-Function] <string> [-ViMode <ViMode>] [<CommonParameters>]
```

Example: Bind the arrow key to a function

```powershell
Set-PSReadLineKeyHandler -Chord UpArrow -Function HistorySearchBackward
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/PSReadLine/Set-PSReadLineKeyHandler.md)


### Set-PSReadLineOption

Version: Both

Module: PSReadLine

Syntax (5.1):

```powershell
Set-PSReadLineOption [-EditMode <EditMode>] [-ContinuationPrompt <string>] [-HistoryNoDuplicates] [-AddToHistoryHandler <Func[string,Object]>] [-CommandValidationHandler <Action[CommandAst]>] [-HistorySearchCursorMovesToEnd] [-MaximumHistoryCount <int>] [-MaximumKillRingCount <int>] [-ShowToolTips] [-ExtraPromptLineCount <int>] [-DingTone <int>] [-DingDuration <int>] [-BellStyle <BellStyle>] [-CompletionQueryItems <int>] [-WordDelimiters <string>] [-HistorySearchCaseSensitive] [-HistorySaveStyle <HistorySaveStyle>] [-HistorySavePath <string>] [-AnsiEscapeTimeout <int>] [-PromptText <string[]>] [-ViModeIndicator <ViModeStyle>] [-ViModeChangeHandler <scriptblock>] [-Colors <hashtable>] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-PSReadLineOption [-EditMode <EditMode>] [-ContinuationPrompt <string>] [-HistoryNoDuplicates] [-AddToHistoryHandler <Func[string,Object]>] [-CommandValidationHandler <Action[CommandAst]>] [-HistorySearchCursorMovesToEnd] [-MaximumHistoryCount <int>] [-MaximumKillRingCount <int>] [-ShowToolTips] [-ExtraPromptLineCount <int>] [-DingTone <int>] [-DingDuration <int>] [-BellStyle <BellStyle>] [-CompletionQueryItems <int>] [-WordDelimiters <string>] [-HistorySearchCaseSensitive] [-HistorySaveStyle <HistorySaveStyle>] [-HistorySavePath <string>] [-AnsiEscapeTimeout <int>] [-PromptText <string[]>] [-ViModeIndicator <ViModeStyle>] [-ViModeChangeHandler <scriptblock>] [-PredictionSource <PredictionSource>] [-PredictionViewStyle <PredictionViewStyle>] [-Colors <hashtable>] [-TerminateOrphanedConsoleApps] [<CommonParameters>]
```

Example (5.1): Set foreground and background colors

```powershell
Set-PSReadLineOption -Colors @{ "Comment"="$([char]0x1b)[32;47m" }
```

Example (7): Set foreground and background colors

```powershell
Set-PSReadLineOption -Colors @{ "Comment"="`e[32;47m" }
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/PSReadLine/Set-PSReadLineOption.md)


### Set-PSResourceRepository

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Set-PSResourceRepository [-Name] <string> [-Uri <string>] [-Trusted] [-Priority <int>] [-ApiVersion <PSRepositoryInfo+APIVersion>] [-CredentialInfo <PSCredentialInfo>] [-CredentialProvider <CredentialProvider>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PSResourceRepository -Repository <hashtable[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Get-PSResourceRepository -Name "PoshTestGallery"
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/set-psresourcerepository?view=powershell-7.5)


### Set-StrictMode

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Set-StrictMode -Version <version> [<CommonParameters>]
Set-StrictMode [-Off] [<CommonParameters>]
```

Example: Turn on strict mode as version 1.0

```powershell
# Strict mode is off by default.
$a -gt 5
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Set-StrictMode.md)


### Set-TraceSource

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Set-TraceSource [-Name] <string[]> [[-Option] <PSTraceSourceOptions>] [-ListenerOption <TraceOptions>] [-FilePath <string>] [-Force] [-Debugger] [-PSHost] [-PassThru] [<CommonParameters>]
Set-TraceSource [-Name] <string[]> [-RemoveListener <string[]>] [<CommonParameters>]
Set-TraceSource [-Name] <string[]> [-RemoveFileListener <string[]>] [<CommonParameters>]
```

Example: Trace the ParameterBinding component

```powershell
Set-TraceSource -Name "ParameterBinding" -Option ExecutionFlow -PSHost -ListenerOption "ProcessId,TimeStamp"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Set-TraceSource.md)


### Set-Variable

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Set-Variable [-Name] <string[]> [[-Value] <Object>] [-Include <string[]>] [-Exclude <string[]>] [-Description <string>] [-Option <ScopedItemOptions>] [-Force] [-Visibility <SessionStateEntryVisibility>] [-PassThru] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Variable [-Name] <string[]> [[-Value] <Object>] [-Include <string[]>] [-Exclude <string[]>] [-Description <string>] [-Option <ScopedItemOptions>] [-Force] [-Visibility <SessionStateEntryVisibility>] [-PassThru] [-Append] [-Scope <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set a variable and get its value

```powershell
Set-Variable -Name "desc" -Value "A description"
Get-Variable -Name "desc"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Set-Variable.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Assignment to read-only automatic variables (PID etc.) is rejected.

- Type: Go implementation.
- Companions: Set-Variable, New-Variable, Remove-Variable, Clear-Variable.
- Function: sets/creates/removes/empties variables.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Variable name |
| `-Value` (position 1) | object | Value (New-Variable's `-Value` also takes position 1) |
| `-Force` | switch | Lets New-Variable overwrite an existing variable |

- Behavior: creating over an existing variable without -Force → error; assigning to read-only automatic variables (PID etc.) → error with Set/New, silently ignored with Clear. Remove deletes straight from the map; Clear sets $null.


### Show-Markdown

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Show-Markdown [-Path] <string[]> [-UseBrowser] [<CommonParameters>]
Show-Markdown -InputObject <psobject> [-UseBrowser] [<CommonParameters>]
Show-Markdown -LiteralPath <string[]> [-UseBrowser] [<CommonParameters>]
```

Example: Simple example specifying a path

```powershell
Show-Markdown -Path ./README.md
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Show-Markdown.md)


### Sort-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Sort-Object [[-Property] <Object[]>] [-Descending] [-Unique] [-InputObject <psobject>] [-Culture <string>] [-CaseSensitive] [<CommonParameters>]
```

Syntax (7):

```powershell
Sort-Object [[-Property] <Object[]>] [-Stable] [-Descending] [-Unique] [-InputObject <psobject>] [-Culture <string>] [-CaseSensitive] [<CommonParameters>]
Sort-Object [[-Property] <Object[]>] -Top <int> [-Descending] [-Unique] [-InputObject <psobject>] [-Culture <string>] [-CaseSensitive] [<CommonParameters>]
Sort-Object [[-Property] <Object[]>] -Bottom <int> [-Descending] [-Unique] [-InputObject <psobject>] [-Culture <string>] [-CaseSensitive] [<CommonParameters>]
```

Example: Sort the current directory by name

```powershell
Get-ChildItem -Path C:\Test | Sort-Object
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Sort-Object.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Stable sort (PowerShell sorting is unstable with no order guarantee for equal elements; this program preserves input order).

- Type: Go implementation.
- Function: sorts. Bash's `sort`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Property` (position 0) | string[] | Sort keys, compared one attribute after another |
| `-Descending` | switch | Descending order (`sort -r`) |
| `-Unique` | switch | Deduplicates (`sort -u`, on sort keys) |
| `-CaseSensitive` | switch | Case-sensitive; insensitive by default |

- Implementation: stable sort; numbers compare by magnitude, strings case-insensitively. `-Unique` deduplicates folding case over sort keys by default, or on original values with `-CaseSensitive`.


### Split-Path

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Split-Path [-Path] <string[]> [-Parent] [-Resolve] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Split-Path [-Path] <string[]> [-Leaf] [-Resolve] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Split-Path [-Path] <string[]> [-Qualifier] [-Resolve] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Split-Path [-Path] <string[]> [-NoQualifier] [-Resolve] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Split-Path [-Path] <string[]> [-Resolve] [-IsAbsolute] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Split-Path -LiteralPath <string[]> [-Resolve] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Split-Path [-Path] <string[]> [-Parent] [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path [-Path] <string[]> -Leaf [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path [-Path] <string[]> -LeafBase [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path [-Path] <string[]> -Extension [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path [-Path] <string[]> -Qualifier [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path [-Path] <string[]> -NoQualifier [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path [-Path] <string[]> -IsAbsolute [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
Split-Path -LiteralPath <string[]> [-Resolve] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Get the qualifier of a path

```powershell
Split-Path -Path "HKCU:\Software\Microsoft" -Qualifier
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Split-Path.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: splits paths.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target path |
| `-Leaf` | switch | Last segment only (`basename`) |
| `-Parent` | switch | Directory part only (`dirname`) |
| `-Qualifier` | switch | Drive qualifier (only drive C exists here: absolute paths return `C:`, relative ones return empty) |

- Note: with no switch given the parent directory is output by default; both `-Leaf` and `-Parent` are switches.


### Start-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Start-Job [-ScriptBlock] <scriptblock> [[-InitializationScript] <scriptblock>] [-Name <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-RunAs32] [-PSVersion <version>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Start-Job [-DefinitionName] <string> [[-DefinitionPath] <string>] [[-Type] <string>] [<CommonParameters>]
Start-Job [[-InitializationScript] <scriptblock>] -LiteralPath <string> [-Name <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-RunAs32] [-PSVersion <version>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Start-Job [-FilePath] <string> [[-InitializationScript] <scriptblock>] [-Name <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-RunAs32] [-PSVersion <version>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Start-Job [-ScriptBlock] <scriptblock> [[-InitializationScript] <scriptblock>] [-Name <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-WorkingDirectory <string>] [-RunAs32] [-PSVersion <version>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Start-Job [-DefinitionName] <string> [[-DefinitionPath] <string>] [[-Type] <string>] [-WorkingDirectory <string>] [<CommonParameters>]
Start-Job [-FilePath] <string> [[-InitializationScript] <scriptblock>] [-Name <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-WorkingDirectory <string>] [-RunAs32] [-PSVersion <version>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Start-Job [[-InitializationScript] <scriptblock>] -LiteralPath <string> [-Name <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-WorkingDirectory <string>] [-RunAs32] [-PSVersion <version>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [<CommonParameters>]
Start-Job [-WorkingDirectory <string>] [-ConnectingTimeout <int>] [-Options <hashtable>] [<CommonParameters>]
```

Example (5.1): Start a background job

```powershell
Start-Job -ScriptBlock { Get-Process -Name powershell }
```

Example (7): Start a background job

```powershell
Start-Job -ScriptBlock { Get-Process -Name pwsh }
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Start-Job.md)


### Start-Process

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Start-Process [-FilePath] <string> [[-ArgumentList] <string[]>] [-Credential <pscredential>] [-WorkingDirectory <string>] [-LoadUserProfile] [-NoNewWindow] [-PassThru] [-RedirectStandardError <string>] [-RedirectStandardInput <string>] [-RedirectStandardOutput <string>] [-WindowStyle <ProcessWindowStyle>] [-Wait] [-UseNewEnvironment] [<CommonParameters>]
Start-Process [-FilePath] <string> [[-ArgumentList] <string[]>] [-WorkingDirectory <string>] [-PassThru] [-Verb <string>] [-WindowStyle <ProcessWindowStyle>] [-Wait] [<CommonParameters>]
```

Syntax (7):

```powershell
Start-Process [-FilePath] <string> [[-ArgumentList] <string[]>] [-Credential <pscredential>] [-WorkingDirectory <string>] [-LoadUserProfile] [-NoNewWindow] [-PassThru] [-RedirectStandardError <string>] [-RedirectStandardInput <string>] [-RedirectStandardOutput <string>] [-WindowStyle <ProcessWindowStyle>] [-Wait] [-UseNewEnvironment] [-Environment <hashtable>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Process [-FilePath] <string> [[-ArgumentList] <string[]>] [-WorkingDirectory <string>] [-PassThru] [-Verb <string>] [-WindowStyle <ProcessWindowStyle>] [-Wait] [-Environment <hashtable>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Start a process that uses default values

```powershell
Start-Process -FilePath "sort.exe"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Start-Process.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Does not take over the new process's input/output.

- Type: Go implementation.
- Function: launches in the background. Bash's `nohup program &`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilePath` (position 0) | path | Program to launch |
| `-ArgumentList` | string[] | Arguments passed to the program; leftover positional arguments also become arguments |

- Implementation: `exec.Command(...).Start()` (doesn't wait, doesn't touch IO).
- Output: the new process's Process object (Id, process name).


### Start-Sleep

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Start-Sleep [-Seconds] <int> [<CommonParameters>]
Start-Sleep -Milliseconds <int> [<CommonParameters>]
```

Syntax (7):

```powershell
Start-Sleep [-Seconds] <double> [<CommonParameters>]
Start-Sleep -Milliseconds <int> [<CommonParameters>]
Start-Sleep [-Duration] <timespan> [<CommonParameters>]
```

Example (5.1): Pause execution for 1 second

```powershell
Start-Sleep -Seconds 1
```

Example (7): Pause execution for 1.5 seconds

```powershell
Start-Sleep -Seconds 1.5
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Start-Sleep.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: sleeps. Bash's `sleep`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Seconds` (position 0) | int | How many seconds to sleep |
| `-Milliseconds` | int | How many milliseconds to sleep |

- `-Seconds` and `-Milliseconds` are mutually exclusive; specifying both is an error.
- `-Seconds` takes only position 0; extra positional arguments are an error.


### Start-ThreadJob

Version: 7 only

Module: Microsoft.PowerShell.ThreadJob

Syntax:

```powershell
Start-ThreadJob [-ScriptBlock] <scriptblock> [-Name <string>] [-InitializationScript <scriptblock>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-ThrottleLimit <int>] [-StreamingHost <PSHost>] [<CommonParameters>]
Start-ThreadJob [-FilePath] <string> [-Name <string>] [-InitializationScript <scriptblock>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-ThrottleLimit <int>] [-StreamingHost <PSHost>] [<CommonParameters>]
```

Example: Create background jobs with a thread limit of 2

```powershell
Start-ThreadJob -ScriptBlock { 1..100 | % { sleep 1; "Output $_" } } -ThrottleLimit 2
Start-ThreadJob -ScriptBlock { 1..100 | % { sleep 1; "Output $_" } }
Start-ThreadJob -ScriptBlock { 1..100 | % { sleep 1; "Output $_" } }
Get-Job
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.6/Microsoft.PowerShell.ThreadJob/Start-ThreadJob.md)


### Start-Transcript

Version: Both

Module: Microsoft.PowerShell.Host

Syntax (5.1):

```powershell
Start-Transcript [[-Path] <string>] [-Append] [-Force] [-NoClobber] [-IncludeInvocationHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Transcript [[-LiteralPath] <string>] [-Append] [-Force] [-NoClobber] [-IncludeInvocationHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Transcript [[-OutputDirectory] <string>] [-Append] [-Force] [-NoClobber] [-IncludeInvocationHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Start-Transcript [[-Path] <string>] [-Append] [-Force] [-NoClobber] [-IncludeInvocationHeader] [-UseMinimalHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Transcript [[-LiteralPath] <string>] [-Append] [-Force] [-NoClobber] [-IncludeInvocationHeader] [-UseMinimalHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Transcript [[-OutputDirectory] <string>] [-Append] [-Force] [-NoClobber] [-IncludeInvocationHeader] [-UseMinimalHeader] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Start a transcript file with default settings

```powershell
Start-Transcript
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Host/Start-Transcript.md)


### Stop-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Stop-Computer [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-AsJob] [-DcomAuthentication <AuthenticationLevel>] [-WsmanAuthentication <string>] [-Protocol <string>] [-Impersonation <ImpersonationLevel>] [-ThrottleLimit <int>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Stop-Computer [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-WsmanAuthentication <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Shut down the local computer

```powershell
Stop-Computer -ComputerName localhost
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Stop-Computer.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`sudo shutdown -h now`).
- Companions: Restart-Computer, Rename-Computer.
- Distro: needs sudo.
- Function: shuts down.
- Parameters: none.
- Implementation: `sudo shutdown -h now`.


### Stop-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Stop-Job [-Id] <int[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Job [-Job] <Job[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Job [-Name] <string[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Job [-InstanceId] <guid[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Job [-State] <JobState> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Job [-Filter] <hashtable> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Stop a job on a remote computer with Invoke-Command

```powershell
$s = New-PSSession -ComputerName Server01 -Credential Domain01\Admin02
$j = Invoke-Command -Session $s -ScriptBlock {Start-Job -ScriptBlock {Get-EventLog -LogName System}}
Invoke-Command -Session $s -ScriptBlock { Stop-Job -Job $Using:j }
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Stop-Job.md)


### Stop-Process

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Stop-Process [-Id] <int[]> [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Process -Name <string[]> [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Process [-InputObject] <Process[]> [-PassThru] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Stop all instances of a process

```powershell
PS C:\> Stop-Process -Name "notepad"
```

Example (7): Stop all instances of a process

```powershell
Stop-Process -Name "notepad"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Stop-Process.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: ends processes. Bash's `kill`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Process name (kills each match by name) |
| `-Id` | int | Process ID |

- Also accepts objects carrying an Id property via the pipeline (as in `Get-Process | Stop-Process`).
- Implementation: numeric arguments kill by PID via `os.FindProcess(pid).Kill()`; otherwise kills each process by name.


### Stop-Transcript

Version: Both

Module: Microsoft.PowerShell.Host

Syntax (5.1):

```powershell
Stop-Transcript [<CommonParameters>]
```

Syntax (7):

```powershell
Stop-Transcript [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Stop all transcripts

```powershell
Stop-Transcript
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Host/Stop-Transcript.md)


### Switch-Process

Version: 7 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Switch-Process [[-WithCommand] <string[]>] [<CommonParameters>]
```

Example: Execute a command that depends on `exec`

```powershell
ssh-copy-id user@host
```


### Tee-Object

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Tee-Object [-FilePath] <string> [-InputObject <psobject>] [-Append] [<CommonParameters>]
Tee-Object -LiteralPath <string> [-InputObject <psobject>] [<CommonParameters>]
Tee-Object -Variable <string> [-InputObject <psobject>] [<CommonParameters>]
```

Syntax (7):

```powershell
Tee-Object [-FilePath] <string> [[-Encoding] <Encoding>] [-InputObject <psobject>] [-Append] [<CommonParameters>]
Tee-Object [[-Encoding] <Encoding>] -LiteralPath <string> [-InputObject <psobject>] [<CommonParameters>]
Tee-Object -Variable <string> [-InputObject <psobject>] [<CommonParameters>]
```

Example: Output processes to a file and to the console

```powershell
Get-Process | Tee-Object -FilePath "C:\Test1\testfile2.txt"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Tee-Object.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: writes to a file while passing output along. `tee` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-FilePath` (position 0) | path | Output file |
| `-Append` | switch | Append mode (`tee -a`) |

- Implementation: writes every input object into the file while returning it unchanged to the pipeline.


### Test-Connection

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Test-Connection [-ComputerName] <string[]> [-AsJob] [-DcomAuthentication <AuthenticationLevel>] [-WsmanAuthentication <string>] [-Protocol <string>] [-BufferSize <int>] [-Count <int>] [-Impersonation <ImpersonationLevel>] [-ThrottleLimit <int>] [-TimeToLive <int>] [-Delay <int>] [<CommonParameters>]
Test-Connection [-ComputerName] <string[]> [-Source] <string[]> [-AsJob] [-DcomAuthentication <AuthenticationLevel>] [-WsmanAuthentication <string>] [-Protocol <string>] [-BufferSize <int>] [-Count <int>] [-Credential <pscredential>] [-Impersonation <ImpersonationLevel>] [-ThrottleLimit <int>] [-TimeToLive <int>] [-Delay <int>] [<CommonParameters>]
Test-Connection [-ComputerName] <string[]> [-DcomAuthentication <AuthenticationLevel>] [-WsmanAuthentication <string>] [-Protocol <string>] [-BufferSize <int>] [-Count <int>] [-Impersonation <ImpersonationLevel>] [-TimeToLive <int>] [-Delay <int>] [-Quiet] [<CommonParameters>]
```

Syntax (7):

```powershell
Test-Connection [-TargetName] <string[]> [-Ping] [-IPv4] [-IPv6] [-ResolveDestination] [-Source <string>] [-MaxHops <int>] [-Count <int>] [-Delay <int>] [-BufferSize <int>] [-DontFragment] [-Quiet] [-TimeoutSeconds <int>] [<CommonParameters>]
Test-Connection [-TargetName] <string[]> -Repeat [-Ping] [-IPv4] [-IPv6] [-ResolveDestination] [-Source <string>] [-MaxHops <int>] [-Delay <int>] [-BufferSize <int>] [-DontFragment] [-Quiet] [-TimeoutSeconds <int>] [<CommonParameters>]
Test-Connection [-TargetName] <string[]> -Traceroute [-IPv4] [-IPv6] [-ResolveDestination] [-Source <string>] [-MaxHops <int>] [-Quiet] [-TimeoutSeconds <int>] [<CommonParameters>]
Test-Connection [-TargetName] <string[]> -MtuSize [-IPv4] [-IPv6] [-ResolveDestination] [-Quiet] [-TimeoutSeconds <int>] [<CommonParameters>]
Test-Connection [-TargetName] <string[]> -TcpPort <int> [-IPv4] [-IPv6] [-ResolveDestination] [-Source <string>] [-Count <int>] [-Delay <int>] [-Repeat] [-Quiet] [-TimeoutSeconds <int>] [-Detailed] [<CommonParameters>]
```

Example (5.1): Send echo requests to a remote computer

```powershell
Test-Connection -ComputerName Server01
```

Example (7): Send echo requests to a remote computer

```powershell
Test-Connection -TargetName Server01 -IPv4
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Test-Connection.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`ping`).
- Distro: needs ping.
- Function: network reachability test. Maps to `ping`.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-TargetName` (position 0) | string | Target host; handed to `ping -c <Count>` |
| `-Count` | int | Number of pings, default 4 |

- Implementation: calls external `ping -c <Count> <target>`, returning whether the exit code was 0.
- Output: Bool. No ping found → error.


### Test-Json

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Test-Json [-Json] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-Json] <string> [-Schema] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-Json] <string> [-SchemaFile] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-Path] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-Path] <string> [-Schema] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-Path] <string> [-SchemaFile] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-LiteralPath] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-LiteralPath] <string> [-Schema] <string> [-Options <string[]>] [<CommonParameters>]
Test-Json [-LiteralPath] <string> [-SchemaFile] <string> [-Options <string[]>] [<CommonParameters>]
```

Example: Test if an object is valid JSON

```powershell
'{"name": "Ashley", "age": 25}' | Test-Json
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Test-Json.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 7.

- Type: Go implementation.
- Function: validates JSON. `jq -e` counterpart.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Json` (position 0) | string | Text to validate; named/positional take precedence, pipeline input only as fallback |

- Output: Bool.


### Test-ModuleManifest

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Test-ModuleManifest [-Path] <string> [<CommonParameters>]
```

Example: Test a manifest

```powershell
Test-ModuleManifest -Path "$PSHOME\Modules\TestModule.psd1"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Test-ModuleManifest.md)


### Test-Path

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Test-Path [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [-UseTransaction] [-OlderThan <datetime>] [-NewerThan <datetime>] [<CommonParameters>]
Test-Path -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [-UseTransaction] [-OlderThan <datetime>] [-NewerThan <datetime>] [<CommonParameters>]
Test-Path [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
Test-Path -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Test-Path [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [-OlderThan <datetime>] [-NewerThan <datetime>] [<CommonParameters>]
Test-Path -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [-OlderThan <datetime>] [-NewerThan <datetime>] [<CommonParameters>]
Test-Path [-Path] <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [<CommonParameters>]
Test-Path -LiteralPath <string[]> [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-PathType <TestPathType>] [-IsValid] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Test a path

```powershell
Test-Path -Path "C:\Documents and Settings\DavidC"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Test-Path.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Function: checks whether a path exists. Returns Bool, bash's `test -e`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Path` (position 0) | path | Target paths, wildcards supported; several allowed (`Test-Path a.txt,b.txt`) |
| `-PathType` | string | Kind filter: `Leaf` files only, `Container` directories only, `Any`/omitted unrestricted |

- Behavior: path exists (including any wildcard hit) → True, else False. `-PathType` filters further by kind atop existence. Nonexistence isn't an error ($? stays true).


### Test-PSScriptFileInfo

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Test-PSScriptFileInfo [-Path] <string> [<CommonParameters>]
```

Example: Test a valid script

```powershell
New-PSScriptFileInfo -Path "C:\MyScripts\test_script.ps1" -Description "this is a test script"
Test-PSScriptFileInfo -Path "C:\MyScripts\test_script.ps1"
True
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/test-psscriptfileinfo?view=powershell-7.5)


### Trace-Command

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Trace-Command [-Name] <string[]> [-Expression] <scriptblock> [[-Option] <PSTraceSourceOptions>] [-InputObject <psobject>] [-ListenerOption <TraceOptions>] [-FilePath <string>] [-Force] [-Debugger] [-PSHost] [<CommonParameters>]
Trace-Command [-Name] <string[]> [-Command] <string> [[-Option] <PSTraceSourceOptions>] [-InputObject <psobject>] [-ArgumentList <Object[]>] [-ListenerOption <TraceOptions>] [-FilePath <string>] [-Force] [-Debugger] [-PSHost] [<CommonParameters>]
```

Example: Trace metadata processing, parameter binding, and an expression

```powershell
Trace-Command -Name Metadata, ParameterBinding, Cmdlet -Expression {Get-Process Notepad} -PSHost
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Trace-Command.md)


### Unblock-File

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Unblock-File [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Unblock-File -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Unblock a file

```powershell
PS C:\> Unblock-File -Path C:\Users\User01\Documents\Downloads\PowerShellTips.chm
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Unblock-File.md)


### Uninstall-Package

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Uninstall-Package [-InputObject] <SoftwareIdentity[]> [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Uninstall-Package [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string[]>] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Scope <string>] [-PackageManagementProvider <string>] [-Type <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [-AllowPrereleaseVersions] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Scope <string>] [-PackageManagementProvider <string>] [-Type <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [-AllowPrereleaseVersions] [<CommonParameters>]
```

Syntax (7):

```powershell
Uninstall-Package [-InputObject] <SoftwareIdentity[]> [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Uninstall-Package [-Name] <string[]> [-RequiredVersion <string>] [-MinimumVersion <string>] [-MaximumVersion <string>] [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string[]>] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Destination <string>] [-ExcludeVersion] [-Scope <string>] [-SkipDependencies] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Scope <string>] [-PackageManagementProvider <string>] [-Type <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [-AllowPrereleaseVersions] [<CommonParameters>]
Uninstall-Package [-AllVersions] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-Scope <string>] [-PackageManagementProvider <string>] [-Type <string>] [-AllowClobber] [-SkipPublisherCheck] [-InstallUpdate] [-NoPathUpdate] [-AllowPrereleaseVersions] [<CommonParameters>]
```

Example: Uninstall a package

```powershell
PS> Uninstall-Package -Name NuGet.Core
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/uninstall-package?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/uninstall-package?view=powershell-7.5)


### Uninstall-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Uninstall-PSResource [-Name] <string[]> [-Version <string>] [-Prerelease] [-SkipDependencyCheck] [-Scope <ScopeType>] [-WhatIf] [-Confirm] [<CommonParameters>]
Uninstall-PSResource [-InputObject] <PSResourceInfo[]> [-Prerelease] [-SkipDependencyCheck] [-Scope <ScopeType>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Uninstall-PSResource Az
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/uninstall-psresource?view=powershell-7.5)


### Unprotect-CmsMessage

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Unprotect-CmsMessage [-EventLogRecord] <EventLogRecord> [[-To] <CmsMessageRecipient[]>] [-IncludeContext] [<CommonParameters>]
Unprotect-CmsMessage [-Content] <string> [[-To] <CmsMessageRecipient[]>] [-IncludeContext] [<CommonParameters>]
Unprotect-CmsMessage [-Path] <string> [[-To] <CmsMessageRecipient[]>] [-IncludeContext] [<CommonParameters>]
Unprotect-CmsMessage [-LiteralPath] <string> [[-To] <CmsMessageRecipient[]>] [-IncludeContext] [<CommonParameters>]
```

Example: Decrypt a message

```powershell
$parameters = @{
  LiteralPath = "C:\Users\Test\Documents\PowerShell\Future_Plans.txt"
  To = '0f 8j b1 ab e0 ce 35 1d 67 d2 f2 6f a2 d2 00 cl 22 z9 m9 85'
}
Unprotect-CmsMessage -LiteralPath @parameters
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Unprotect-CmsMessage.md)


### Unregister-Event

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Unregister-Event [-SourceIdentifier] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-Event [-SubscriptionId] <int> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Cancel an event subscription by source identifier

```powershell
Unregister-Event -SourceIdentifier "ProcessStarted"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Unregister-Event.md)


### Unregister-PackageSource

Version: Both

Module: PackageManagement

Syntax (5.1):

```powershell
Unregister-PackageSource [[-Source] <string>] [-Location <string>] [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string>] [<CommonParameters>]
Unregister-PackageSource -InputObject <PackageSource[]> [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Unregister-PackageSource [[-Source] <string>] [-Location <string>] [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ProviderName <string>] [<CommonParameters>]
Unregister-PackageSource -InputObject <PackageSource[]> [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-ConfigFile <string>] [-SkipValidate] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
Unregister-PackageSource [-Credential <pscredential>] [-Force] [-ForceBootstrap] [-WhatIf] [-Confirm] [-PackageManagementProvider <string>] [-PublishLocation <string>] [-ScriptSourceLocation <string>] [-ScriptPublishLocation <string>] [<CommonParameters>]
```

Example: Unregister a package source for the NuGet provider

```powershell
PS> Unregister-PackageSource -Source MyNuGet
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/unregister-packagesource?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/packagemanagement/unregister-packagesource?view=powershell-7.5)


### Unregister-PSResourceRepository

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Unregister-PSResourceRepository [-Name] <string[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Get-PSResourceRepository
Unregister-PSResourceRepository -Name PSGv3
Get-PSResourceRepository
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/unregister-psresourcerepository?view=powershell-7.5)


### Update-FormatData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Update-FormatData [[-AppendPath] <string[]>] [-PrependPath <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Reload previously loaded formatting files

```powershell
Update-FormatData
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Update-FormatData.md)


### Update-Help

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Update-Help [[-Module] <string[]>] [[-SourcePath] <string[]>] [[-UICulture] <cultureinfo[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Recurse] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Update-Help [[-Module] <string[]>] [[-UICulture] <cultureinfo[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-LiteralPath <string[]>] [-Recurse] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Update-Help [[-Module] <string[]>] [[-SourcePath] <string[]>] [[-UICulture] <cultureinfo[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-Recurse] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [-Scope <UpdateHelpScope>] [-WhatIf] [-Confirm] [<CommonParameters>]
Update-Help [[-Module] <string[]>] [[-UICulture] <cultureinfo[]>] [-FullyQualifiedModule <ModuleSpecification[]>] [-LiteralPath <string[]>] [-Recurse] [-Credential <pscredential>] [-UseDefaultCredentials] [-Force] [-Scope <UpdateHelpScope>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Update help files for all modules

```powershell
Update-Help
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Update-Help.md)


### Update-List

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Update-List [[-Property] <string>] [-Add <Object[]>] [-Remove <Object[]>] [-InputObject <psobject>] [<CommonParameters>]
Update-List [[-Property] <string>] -Replace <Object[]> [-InputObject <psobject>] [<CommonParameters>]
```

Example (5.1): Add items to a property value

```powershell
class Cards {

    [System.Collections.Generic.List[string]]$Cards
    [string]$Name

    Cards([string]$_name) {
        $this.Name = $_name
        $this.Cards = [System.Collections.Generic.List[string]]::new()
    }

    NewDeck() {
        $_suits = [char]0x2663,[char]0x2666,[char]0x2665,[char]0x2660
        $_values = 'A',2,3,4,5,6,7,8,9,10,'J','Q','K'
        $_deck = foreach ($s in $_suits){ foreach ($v in $_values){ "$v$s"} }
        $this | Update-List -Property Cards -Add $_deck | Out-Null
    }

    Show() {
        Write-Host
        Write-Host $this.Name ": " $this.Cards[0..12]
        if ($this.Cards.Count -gt 13) {
            Write-Host (' ' * ($this.Name.Length+3)) $this.Cards[13..25]
        }
        if ($this.Cards.Count -gt 26) {
            Write-Host (' ' * ($this.Name.Length+3)) $this.Cards[26..38]
        }
        if ($this.Cards.Count -gt 39) {
            Write-Host (' ' * ($this.Name.Length+3)) $this.Cards[39..51]
        }
    }

    Shuffle() { $this.Cards = Get-Random -InputObject $this.Cards -Count 52 }

    Sort() { $this.Cards.Sort() }
}
```

Example (7): Add items to a property value

```powershell
class Cards {

    [System.Collections.Generic.List[string]]$Cards
    [string]$Name

    Cards([string]$_name) {
        $this.Name = $_name
        $this.Cards = [System.Collections.Generic.List[string]]::new()
    }

    NewDeck() {
        $_suits = "`u{2663}","`u{2666}","`u{2665}","`u{2660}"
        $_values = 'A',2,3,4,5,6,7,8,9,10,'J','Q','K'
        $_deck = foreach ($s in $_suits){ foreach ($v in $_values){ "$v$s"} }
        $this | Update-List -Property Cards -Add $_deck | Out-Null
    }

    Show() {
        Write-Host
        Write-Host $this.Name ": " $this.Cards[0..12]
        if ($this.Cards.Count -gt 13) {
            Write-Host (' ' * ($this.Name.Length+3)) $this.Cards[13..25]
        }
        if ($this.Cards.Count -gt 26) {
            Write-Host (' ' * ($this.Name.Length+3)) $this.Cards[26..38]
        }
        if ($this.Cards.Count -gt 39) {
            Write-Host (' ' * ($this.Name.Length+3)) $this.Cards[39..51]
        }
    }

    Shuffle() { $this.Cards = Get-Random -InputObject $this.Cards -Count 52 }

    Sort() { $this.Cards.Sort() }
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Update-List.md)


### Update-PSModuleManifest

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Update-PSModuleManifest [-Path] <string> [-NestedModules <Object[]>] [-Guid <guid>] [-Author <string>] [-CompanyName <string>] [-Copyright <string>] [-RootModule <string>] [-ModuleVersion <version>] [-Description <string>] [-ProcessorArchitecture <ProcessorArchitecture>] [-CompatiblePSEditions <string[]>] [-PowerShellVersion <version>] [-ClrVersion <version>] [-DotNetFrameworkVersion <version>] [-PowerShellHostName <string>] [-PowerShellHostVersion <version>] [-RequiredModules <Object[]>] [-TypesToProcess <string[]>] [-FormatsToProcess <string[]>] [-ScriptsToProcess <string[]>] [-RequiredAssemblies <string[]>] [-FileList <string[]>] [-ModuleList <Object[]>] [-FunctionsToExport <string[]>] [-AliasesToExport <string[]>] [-VariablesToExport <string[]>] [-CmdletsToExport <string[]>] [-DscResourcesToExport <string[]>] [-Tags <string[]>] [-ProjectUri <uri>] [-LicenseUri <uri>] [-IconUri <uri>] [-ReleaseNotes <string>] [-Prerelease <string>] [-HelpInfoUri <uri>] [-DefaultCommandPrefix <string>] [-ExternalModuleDependencies <string[]>] [-RequireLicenseAcceptance] [-PrivateData <hashtable>] [<CommonParameters>]
```

Example: 

```powershell
Update-PSModuleManifest -Path 'C:\MyModules\TestModule' -Author 'New Author'
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/update-psmodulemanifest?view=powershell-7.5)


### Update-PSResource

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Update-PSResource [[-Name] <string[]>] [-Version <string>] [-Prerelease] [-Repository <string[]>] [-Scope <ScopeType>] [-TemporaryPath <string>] [-TrustRepository] [-Credential <pscredential>] [-Quiet] [-AcceptLicense] [-Force] [-PassThru] [-SkipDependencyCheck] [-AuthenticodeCheck] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Get-PSResource -Name "TestModule"
Update-PSResource -Name "TestModule"
Get-PSResource -Name "TestModule"
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/update-psresource?view=powershell-7.5)


### Update-PSScriptFileInfo

Version: 7 only

Module: Microsoft.PowerShell.PSResourceGet

Syntax:

```powershell
Update-PSScriptFileInfo [-Path] <string> [-Author <string>] [-CompanyName <string>] [-Copyright <string>] [-Description <string>] [-ExternalModuleDependencies <string[]>] [-ExternalScriptDependencies <string[]>] [-Guid <guid>] [-IconUri <string>] [-LicenseUri <string>] [-PrivateData <string>] [-ProjectUri <string>] [-ReleaseNotes <string>] [-RemoveSignature] [-RequiredModules <hashtable[]>] [-RequiredScripts <string[]>] [-Tags <string[]>] [-Version <string>] [<CommonParameters>]
```

Example: Update the version of a script

```powershell
$parameters = @{
    FilePath = "C:\Users\johndoe\MyScripts\test_script.ps1"
    Version = "1.0.0.0"
    Description = "this is a test script"
}
New-PSScriptFileInfo @parameters
$parameters.Version = "2.0.0.0"
Update-PSScriptFileInfo @parameters
Get-Content $parameters.FilePath
```

Source: [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.psresourceget/update-psscriptfileinfo?view=powershell-7.5)


### Update-TypeData

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Update-TypeData [[-AppendPath] <string[]>] [-PrependPath <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Update-TypeData -TypeName <string> [-MemberType <PSMemberTypes>] [-MemberName <string>] [-Value <Object>] [-SecondValue <Object>] [-TypeConverter <type>] [-TypeAdapter <type>] [-SerializationMethod <string>] [-TargetTypeForDeserialization <type>] [-SerializationDepth <int>] [-DefaultDisplayProperty <string>] [-InheritPropertySerializationSet <Nullable`1>] [-StringSerializationSource <string>] [-DefaultDisplayPropertySet <string[]>] [-DefaultKeyPropertySet <string[]>] [-PropertySerializationSet <string[]>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Update-TypeData [-TypeData] <TypeData[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Update extended types

```powershell
Update-TypeData
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Update-TypeData.md)


### Wait-Debugger

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Wait-Debugger [<CommonParameters>]
```

Example: Insert breakpoint for debugging

```powershell
function Test-Condition {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory)]
        [string]$Name,
        [string]$Message = "Hello, $Name!"
    )

    if ($Name -eq $Env:USERNAME) {
        Write-Output "$Message"
    } else {
        # Remove after debugging
        Wait-Debugger

        Write-Output "$Name is not the current user."
    }
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Wait-Debugger.md)


### Wait-Event

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Wait-Event [[-SourceIdentifier] <string>] [-Timeout <int>] [<CommonParameters>]
```

Example: Wait for the next event

```powershell
Wait-Event
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Wait-Event.md)


### Wait-Job

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Wait-Job [-Id] <int[]> [-Any] [-Timeout <int>] [-Force] [<CommonParameters>]
Wait-Job [-Job] <Job[]> [-Any] [-Timeout <int>] [-Force] [<CommonParameters>]
Wait-Job [-Name] <string[]> [-Any] [-Timeout <int>] [-Force] [<CommonParameters>]
Wait-Job [-InstanceId] <guid[]> [-Any] [-Timeout <int>] [-Force] [<CommonParameters>]
Wait-Job [-State] <JobState> [-Any] [-Timeout <int>] [-Force] [<CommonParameters>]
Wait-Job [-Filter] <hashtable> [-Any] [-Timeout <int>] [-Force] [<CommonParameters>]
```

Example: Wait for all jobs

```powershell
Get-Job | Wait-Job
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Wait-Job.md)


### Wait-Process

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Wait-Process [-Name] <string[]> [[-Timeout] <int>] [<CommonParameters>]
Wait-Process [-Id] <int[]> [[-Timeout] <int>] [<CommonParameters>]
Wait-Process [[-Timeout] <int>] -InputObject <Process[]> [<CommonParameters>]
```

Syntax (7):

```powershell
Wait-Process [-Name] <string[]> [[-Timeout] <int>] [-Any] [-PassThru] [<CommonParameters>]
Wait-Process [-Id] <int[]> [[-Timeout] <int>] [-Any] [-PassThru] [<CommonParameters>]
Wait-Process [[-Timeout] <int>] -InputObject <Process[]> [-Any] [-PassThru] [<CommonParameters>]
```

Example (5.1): Stop a process and wait

```powershell
PS C:\> $nid = (Get-Process notepad).Id
PS C:\> Stop-Process -Id $nid
PS C:\> Wait-Process -Id $nid
```

Example (7): Stop a process and wait

```powershell
$nid = (Get-Process notepad).Id
Stop-Process -Id $nid
Wait-Process -Id $nid
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Wait-Process.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: waits for a process to finish. Bash's `wait`.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Process name |
| `-Id` | int | Process ID |

- Implementation: on Linux reads `/proc/<pid>/stat` to check whether the process is active (Z/X count as ended), polling every 100ms; by name, polls the process list until it disappears. A target already gone at entry errors immediately.


### Where-Object

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Where-Object [-Property] <string> [[-Value] <Object>] [-InputObject <psobject>] [-EQ] [<CommonParameters>]
Where-Object [-FilterScript] <scriptblock> [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CIn [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CEQ [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -GT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CGT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -LT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CLT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -GE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CGE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -LE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CLE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Like [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CLike [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotLike [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotLike [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Match [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CMatch [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotMatch [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotMatch [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Contains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CContains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotContains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotContains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -In [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotIn [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotIn [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Is [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -IsNot [-InputObject <psobject>] [<CommonParameters>]
```

Syntax (7):

```powershell
Where-Object [-Property] <string> [[-Value] <Object>] [-InputObject <psobject>] [-EQ] [<CommonParameters>]
Where-Object [-FilterScript] <scriptblock> [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CEQ [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -GT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CGT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -LT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CLT [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -GE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CGE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -LE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CLE [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Like [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CLike [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotLike [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotLike [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Match [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CMatch [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotMatch [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotMatch [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Contains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CContains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotContains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotContains [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -In [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CIn [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -NotIn [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -CNotIn [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -Is [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> [[-Value] <Object>] -IsNot [-InputObject <psobject>] [<CommonParameters>]
Where-Object [-Property] <string> -Not [-InputObject <psobject>] [<CommonParameters>]
```

Example: Get stopped services

```powershell
Get-Service | Where-Object { $_.Status -eq "Stopped" }
Get-Service | Where-Object Status -EQ "Stopped"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Where-Object.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
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


### Write-Debug

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Write-Debug [-Message] <string> [<CommonParameters>]
```

Example: Understand $DebugPreference

```powershell
Write-Debug "Cannot open file."
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Debug.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Write-Error, Write-Warning, Write-Verbose, Write-Debug, Write-Information.
- Function: writes leveled messages onto stderr/stdout.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Message` (position 0) | string | Message text (-MessageData for Write-Information) |

- Implementation: Error/Warning/Verbose/Debug go to stderr (`echo ... 1>&2`); Information goes to stdout. Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG). Error additionally sets $?=false.


### Write-Error

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Write-Error [-Message] <string> [-Category <ErrorCategory>] [-ErrorId <string>] [-TargetObject <Object>] [-RecommendedAction <string>] [-CategoryActivity <string>] [-CategoryReason <string>] [-CategoryTargetName <string>] [-CategoryTargetType <string>] [<CommonParameters>]
Write-Error -Exception <Exception> [-Message <string>] [-Category <ErrorCategory>] [-ErrorId <string>] [-TargetObject <Object>] [-RecommendedAction <string>] [-CategoryActivity <string>] [-CategoryReason <string>] [-CategoryTargetName <string>] [-CategoryTargetType <string>] [<CommonParameters>]
Write-Error -ErrorRecord <ErrorRecord> [-RecommendedAction <string>] [-CategoryActivity <string>] [-CategoryReason <string>] [-CategoryTargetName <string>] [-CategoryTargetType <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Write-Error [-Message] <string> [-Category <ErrorCategory>] [-ErrorId <string>] [-TargetObject <Object>] [-RecommendedAction <string>] [-CategoryActivity <string>] [-CategoryReason <string>] [-CategoryTargetName <string>] [-CategoryTargetType <string>] [<CommonParameters>]
Write-Error [-Exception] <Exception> [-Message <string>] [-Category <ErrorCategory>] [-ErrorId <string>] [-TargetObject <Object>] [-RecommendedAction <string>] [-CategoryActivity <string>] [-CategoryReason <string>] [-CategoryTargetName <string>] [-CategoryTargetType <string>] [<CommonParameters>]
Write-Error [-ErrorRecord] <ErrorRecord> [-RecommendedAction <string>] [-CategoryActivity <string>] [-CategoryReason <string>] [-CategoryTargetName <string>] [-CategoryTargetType <string>] [<CommonParameters>]
```

Example: Write an error for RegistryKey object

```powershell
Get-ChildItem | ForEach-Object {
    if ($_.GetType().ToString() -eq "Microsoft.Win32.RegistryKey")
    {
        Write-Error "Invalid object" -ErrorId B1 -TargetObject $_
    }
    else
    {
        $_
    }
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Error.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG).

- Type: Go implementation.
- Companions: Write-Error, Write-Warning, Write-Verbose, Write-Debug, Write-Information.
- Function: writes leveled messages onto stderr/stdout.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Message` (position 0) | string | Message text (-MessageData for Write-Information) |

- Implementation: Error/Warning/Verbose/Debug go to stderr (`echo ... 1>&2`); Information goes to stdout. Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG). Error additionally sets $?=false.


### Write-Host

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Write-Host [[-Object] <Object>] [-NoNewline] [-Separator <Object>] [-ForegroundColor <ConsoleColor>] [-BackgroundColor <ConsoleColor>] [<CommonParameters>]
```

Example: Write to the console without adding a new line

```powershell
Write-Host "no newline test " -NoNewline
Write-Host "second string"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Host.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: writes straight to the screen, bypassing the pipeline. `echo` (to the terminal) territory.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Object` (position 0) | object | Objects to show, multiple allowed |
| `-NoNewline` | switch | No trailing line break (`echo -n`) |

- Implementation: joins each object's string with spaces and writes to stdout.


### Write-Information

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Write-Information [-MessageData] <Object> [[-Tags] <string[]>] [<CommonParameters>]
```

Example: Write information for Get- results

```powershell
Write-Information -MessageData "Processes starting with 'P'" -InformationAction Continue
Get-Process -Name p*
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Information.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Write-Error, Write-Warning, Write-Verbose, Write-Debug, Write-Information.
- Function: writes leveled messages onto stderr/stdout.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Message` (position 0) | string | Message text (-MessageData for Write-Information) |

- Implementation: Error/Warning/Verbose/Debug go to stderr (`echo ... 1>&2`); Information goes to stdout. Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG). Error additionally sets $?=false.


### Write-Output

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Write-Output [-InputObject] <psobject[]> [-NoEnumerate] [<CommonParameters>]
```

Syntax (7):

```powershell
Write-Output [-InputObject] <PSObject[]> [-NoEnumerate] [<CommonParameters>]
```

Example: Get objects and write them to the console

```powershell
$P = Get-Process
Write-Output $P
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Output.md)

#### Implementation in PowerShell For Linux:

- Consistent with the original 5.1/7.

- Type: Go implementation.
- Function: emits objects.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-InputObject` (position 0) | object | Objects to emit, multiple allowed; pipeline input takes precedence |

- Behavior: places objects into the pipeline unchanged — bash's `echo` counterpart that keeps object types.


### Write-Progress

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax (5.1):

```powershell
Write-Progress [-Activity] <string> [[-Status] <string>] [[-Id] <int>] [-PercentComplete <int>] [-SecondsRemaining <int>] [-CurrentOperation <string>] [-ParentId <int>] [-Completed] [-SourceId <int>] [<CommonParameters>]
```

Syntax (7):

```powershell
Write-Progress [[-Activity] <string>] [[-Status] <string>] [[-Id] <int>] [-PercentComplete <int>] [-SecondsRemaining <int>] [-CurrentOperation <string>] [-ParentId <int>] [-Completed] [-SourceId <int>] [<CommonParameters>]
```

Example: Display the progress of a `for` loop

```powershell
for ($i = 1; $i -le 100; $i++ ) {
    Write-Progress -Activity "Search in Progress" -Status "$i% Complete:" -PercentComplete $i
    Start-Sleep -Milliseconds 250
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Progress.md)


### Write-Verbose

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Write-Verbose [-Message] <string> [<CommonParameters>]
```

Example: Write a status message

```powershell
Write-Verbose -Message "Searching the Application Event Log."
Write-Verbose -Message "Searching the Application Event Log." -Verbose
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Verbose.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Write-Error, Write-Warning, Write-Verbose, Write-Debug, Write-Information.
- Function: writes leveled messages onto stderr/stdout.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Message` (position 0) | string | Message text (-MessageData for Write-Information) |

- Implementation: Error/Warning/Verbose/Debug go to stderr (`echo ... 1>&2`); Information goes to stdout. Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG). Error additionally sets $?=false.


### Write-Warning

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Write-Warning [-Message] <string> [<CommonParameters>]
```

Example: Write a warning message

```powershell
Write-Warning "This is only a test warning."
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Write-Warning.md)

#### Implementation in PowerShell For Linux:

- Type: Go implementation.
- Companions: Write-Error, Write-Warning, Write-Verbose, Write-Debug, Write-Information.
- Function: writes leveled messages onto stderr/stdout.

| Parameter | Type | Meaning |
| :--- | :--- | :--- |
| `-Message` (position 0) | string | Message text (-MessageData for Write-Information) |

- Implementation: Error/Warning/Verbose/Debug go to stderr (`echo ... 1>&2`); Information goes to stdout. Prefixes follow the UI language (Chinese: 错误/警告/详细/调试; English: ERROR/WARNING/VERBOSE/DEBUG). Error additionally sets $?=false.


