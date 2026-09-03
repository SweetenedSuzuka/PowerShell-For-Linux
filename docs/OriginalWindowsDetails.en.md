# Original Windows-Only Command Details

Commands in the original PowerShell that are Windows-only.
PowerShell For Linux does not implement them by default, but occasionally implements some for special reasons, as marked below.

Status legend:
- **Go implementation** — behavior is reproduced in Go inside this program, with no external commands called.
- **Mapped Linux** — native Linux commands/tools are called; the tool follows in parentheses.
- **Not implemented** — not implemented in this program.
- **Out of scope** — Windows-only; this program does not do it, see notes for reasons.

## Command List

| Command | Module | Version | Differences | Purpose | Status | Notes |
|---|---|---|---|---|---|---|
| [`Add-AppProvisionedSharedPackageContainer`](#add-appprovisionedsharedpackagecontainer) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Add-AppSharedPackageContainer`](#add-appsharedpackagecontainer) | Appx | Both | None | Deploys the shared package container definition. | Out of scope |  |
| [`Add-AppvClientConnectionGroup`](#add-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Creates a composition of multiple packages. | Out of scope |  |
| [`Add-AppvClientPackage`](#add-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Adds a package to a computer running the App-V client. | Out of scope |  |
| [`Add-AppvPublishingServer`](#add-appvpublishingserver) | AppvClient | 5.1 only | 5.1 only | Adds a publishing server for the computer that runs the App-V client. | Out of scope |  |
| [`Add-AppxPackage`](#add-appxpackage) | Appx | Both | None | Adds a signed app package to a user account. | Out of scope |  |
| [`Add-AppxProvisionedPackage`](#add-appxprovisionedpackage) | Dism | Both | Syntax differs | Adds an app package (.appx) that will install for each new user to a Windows image. | Out of scope |  |
| [`Add-AppxVolume`](#add-appxvolume) | Appx | Both | None | Adds an appx volume to the Package Manager. | Out of scope |  |
| [`Add-BitsFile`](#add-bitsfile) | BitsTransfer | Both | None | Adds one or more files to an existing BITS transfer job. | Out of scope |  |
| [`Add-CertificateEnrollmentPolicyServer`](#add-certificateenrollmentpolicyserver) | PKI | Both | None | Adds an enrollment policy server to the current user or local system configuration. | Out of scope |  |
| [`Add-Computer`](#add-computer) | Microsoft.PowerShell.Management | Both | None | Add the local computer to a domain or workgroup. | Out of scope |  |
| [`Add-JobTrigger`](#add-jobtrigger) | PSScheduledJob | Both | None | Adds job triggers to scheduled jobs. | Out of scope |  |
| [`Add-KdsRootKey`](#add-kdsrootkey) | Kds | Both | None | Generates a new root key for the Microsoft Group KdsSvc within Active Directory. | Out of scope |  |
| [`Add-LocalGroupMember`](#add-localgroupmember) | Microsoft.PowerShell.LocalAccounts | Both | None | Adds members to a local group. | Out of scope |  |
| [`Add-PSSnapin`](#add-pssnapin) | Microsoft.PowerShell.Core | 5.1 only | 5.1 only | Adds one or more Windows PowerShell snap-ins to the current session. | Out of scope |  |
| [`Add-SignerRule`](#add-signerrule) | ConfigCI | Both | None | Creates a signer rule and adds it to a policy. | Out of scope |  |
| [`Add-WindowsCapability`](#add-windowscapability) | Dism | Both | None | Installs a Windows capability package on the specified operating system image. | Out of scope |  |
| [`Add-WindowsDriver`](#add-windowsdriver) | Dism | Both | None | Adds a driver to an offline Windows image. | Out of scope |  |
| [`Add-WindowsImage`](#add-windowsimage) | Dism | Both | None | Adds an additional image to an existing image (.wim) file. | Out of scope |  |
| [`Add-WindowsPackage`](#add-windowspackage) | Dism | Both | None | Adds a single .cab or .msu file to a Windows image. | Out of scope |  |
| [`Checkpoint-Computer`](#checkpoint-computer) | Microsoft.PowerShell.Management | Both | None | Creates a system restore point on the local computer. | Out of scope |  |
| [`Clear-EventLog`](#clear-eventlog) | Microsoft.PowerShell.Management | Both | None | Clears all entries from specified event logs on the local or remote computers. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Clear-KdsCache`](#clear-kdscache) | Kds | Both | None | Clears the group key cache of the local computer. | Out of scope |  |
| [`Clear-RecycleBin`](#clear-recyclebin) | Microsoft.PowerShell.Management | 7 only | 7 only, spelled Clear-Recyclebin in 5.1 (letter case differs) | Clears the contents of the current user's recycle bin. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Clear-Recyclebin`](#clear-recyclebin) | Microsoft.PowerShell.Management | 5.1 only | 5.1 only, spelled Clear-RecycleBin in 7 (letter case differs) | Clears the contents of the current user's recycle bin. | Out of scope |  |
| [`Clear-ReFSDedupSchedule`](#clear-refsdedupschedule) | Microsoft.ReFsDedup.Commands | Both | None | Clears the scheduled task for deduplication on a specified ReFS volume. | Out of scope |  |
| [`Clear-ReFSDedupScrubSchedule`](#clear-refsdedupscrubschedule) | Microsoft.ReFsDedup.Commands | Both | None | Clears the deduplication scrub schedule on a specified ReFS volume. | Out of scope |  |
| [`Clear-Tpm`](#clear-tpm) | TrustedPlatformModule | Both | None | Resets a TPM to its default state. | Out of scope |  |
| [`Clear-UevAppxPackage`](#clear-uevappxpackage) | UEV | 5.1 only | 5.1 only | Clears a setting in the computer or user sections of the registry. | Out of scope |  |
| [`Clear-UevConfiguration`](#clear-uevconfiguration) | UEV | 5.1 only | 5.1 only | Clears UE-V configuration settings. | Out of scope |  |
| [`Clear-WindowsCorruptMountPoint`](#clear-windowscorruptmountpoint) | Dism | Both | None | Deletes all of the resources associated with a mounted image that has been corrupted. | Out of scope |  |
| [`Complete-BitsTransfer`](#complete-bitstransfer) | BitsTransfer | Both | None | Completes a BITS transfer job. | Out of scope |  |
| [`Complete-DtcDiagnosticTransaction`](#complete-dtcdiagnostictransaction) | MsDtc | Both | None | Invokes the Commit process if the specified transaction is the root transaction; otherwise, invokes the Complete method on a transaction object. | Out of scope |  |
| [`Complete-Transaction`](#complete-transaction) | Microsoft.PowerShell.Management | Both | None | Commits the active transaction. | Out of scope |  |
| [`Confirm-SecureBootUEFI`](#confirm-securebootuefi) | SecureBoot | Both | None | Confirms that Secure Boot is enabled by checking the Secure Boot status on the local computer. | Out of scope |  |
| [`Connect-PSSession`](#connect-pssession) | Microsoft.PowerShell.Core | Both | None | Reconnects to disconnected sessions. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Connect-WSMan`](#connect-wsman) | Microsoft.WSMan.Management | Both | None | Connects to the WinRM service on a remote computer. | Out of scope |  |
| [`Convert-String`](#convert-string) | Microsoft.PowerShell.Utility | 5.1 only | 5.1 only | Formats a string to match examples. | Out of scope |  |
| [`ConvertFrom-CIPolicy`](#convertfrom-cipolicy) | ConfigCI | Both | None | Converts an .xml file that contains a Code Integrity policy into binary format. | Out of scope |  |
| [`ConvertFrom-SddlString`](#convertfrom-sddlstring) | Microsoft.PowerShell.Utility | 7 only | 7 only | Converts a SDDL string to a custom object. | Out of scope | Serialization / markup / formatting (rarely used) |
| [`ConvertFrom-String`](#convertfrom-string) | Microsoft.PowerShell.Utility | 5.1 only | 5.1 only | Extracts and parses structured properties from string content. | Out of scope |  |
| [`ConvertTo-ProcessMitigationPolicy`](#convertto-processmitigationpolicy) | ProcessMitigations | Both | None | Converts an mitigation policy file formats. | Out of scope |  |
| [`ConvertTo-TpmOwnerAuth`](#convertto-tpmownerauth) | TrustedPlatformModule | Both | None | Creates a TPM owner authorization value from a supplied string. | Out of scope |  |
| [`Copy-BcdEntry`](#copy-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Copy-UserInternationalSettingsToSystem`](#copy-userinternationalsettingstosystem) | International | Both | None | Copies the current user's international settings (Windows Display language, Input language, Regional Format/locale, and Location/GeoID) to one or both of the following: * Welcome screen and system accounts | Out of scope |  |
| [`Disable-AppBackgroundTaskDiagnosticLog`](#disable-appbackgroundtaskdiagnosticlog) | AppBackgroundTask | Both | None | Disables background task logging in Event Viewer. | Out of scope |  |
| [`Disable-Appv`](#disable-appv) | AppvClient | 5.1 only | 5.1 only | Disables the App-V service. | Out of scope |  |
| [`Disable-AppvClientConnectionGroup`](#disable-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Disables a connection group on the computer running the App-V client. | Out of scope |  |
| [`Disable-BcdElementBootDebug`](#disable-bcdelementbootdebug) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Disable-BcdElementBootEms`](#disable-bcdelementbootems) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Disable-BcdElementDebug`](#disable-bcdelementdebug) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Disable-BcdElementEms`](#disable-bcdelementems) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Disable-BcdElementEventLogging`](#disable-bcdelementeventlogging) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Disable-BcdElementHypervisorDebug`](#disable-bcdelementhypervisordebug) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Disable-ComputerRestore`](#disable-computerrestore) | Microsoft.PowerShell.Management | Both | None | Disables the System Restore feature on the specified file system drive. | Out of scope |  |
| [`Disable-JobTrigger`](#disable-jobtrigger) | PSScheduledJob | Both | None | Disables the job triggers of scheduled jobs. | Out of scope |  |
| [`Disable-LocalUser`](#disable-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Disables a local user account. | Out of scope |  |
| [`Disable-PSRemoting`](#disable-psremoting) | Microsoft.PowerShell.Core | Both | None | Prevents PowerShell endpoints from receiving remote connections. | Out of scope |  |
| [`Disable-PSSessionConfiguration`](#disable-pssessionconfiguration) | Microsoft.PowerShell.Core | Both | None | Disables session configurations on the local computer. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Disable-ReFSDedup`](#disable-refsdedup) | Microsoft.ReFsDedup.Commands | Both | None | Disables data deduplication on a specified ReFS volume. | Out of scope |  |
| [`Disable-ScheduledJob`](#disable-scheduledjob) | PSScheduledJob | Both | None | Disables a scheduled job. | Out of scope |  |
| [`Disable-TlsCipherSuite`](#disable-tlsciphersuite) | TLS | Both | None | Disables a TLS cipher suite. | Out of scope |  |
| [`Disable-TlsEccCurve`](#disable-tlsecccurve) | TLS | Both | None | Disables the Elliptic Curve Cryptography (ECC) cipher suites available for TLS(Transport Layer Security) for a computer. | Out of scope |  |
| [`Disable-TlsSessionTicketKey`](#disable-tlssessionticketkey) | TLS | Both | None | Disables a TLS session ticket key. | Out of scope |  |
| [`Disable-TpmAutoProvisioning`](#disable-tpmautoprovisioning) | TrustedPlatformModule | Both | None | Disables TPM auto-provisioning. | Out of scope |  |
| [`Disable-Uev`](#disable-uev) | UEV | 5.1 only | 5.1 only | Disables the UE-V service. | Out of scope |  |
| [`Disable-UevAppxPackage`](#disable-uevappxpackage) | UEV | 5.1 only | 5.1 only | Disables UE-V synchronization of Windows 8 apps. | Out of scope |  |
| [`Disable-UevTemplate`](#disable-uevtemplate) | UEV | 5.1 only | 5.1 only | Disables a settings location template. | Out of scope |  |
| [`Disable-WindowsErrorReporting`](#disable-windowserrorreporting) | WindowsErrorReporting | Both | None | Disables Windows Error Reporting. | Out of scope |  |
| [`Disable-WindowsOptionalFeature`](#disable-windowsoptionalfeature) | Dism | Both | None | Disables a feature in a Windows image. | Out of scope |  |
| [`Disable-WSManCredSSP`](#disable-wsmancredssp) | Microsoft.WSMan.Management | Both | None | Disables CredSSP authentication on a computer. | Out of scope |  |
| [`Disconnect-PSSession`](#disconnect-pssession) | Microsoft.PowerShell.Core | Both | Syntax differs | Disconnects from a session. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Disconnect-WSMan`](#disconnect-wsman) | Microsoft.WSMan.Management | Both | None | Disconnects the client from the WinRM service on a remote computer. | Out of scope |  |
| [`Dismount-AppxVolume`](#dismount-appxvolume) | Appx | Both | None | Dismounts an appx volume. | Out of scope |  |
| [`Dismount-WindowsImage`](#dismount-windowsimage) | Dism | Both | None | Dismounts a Windows image from the directory it is mapped to. | Out of scope |  |
| [`Edit-CIPolicyRule`](#edit-cipolicyrule) | ConfigCI | Both | None | This cmdlet is not supported. | Out of scope |  |
| [`Enable-AppBackgroundTaskDiagnosticLog`](#enable-appbackgroundtaskdiagnosticlog) | AppBackgroundTask | Both | None | Enables background task logging in Event Viewer. | Out of scope |  |
| [`Enable-Appv`](#enable-appv) | AppvClient | 5.1 only | 5.1 only | Enables the App-V service. | Out of scope |  |
| [`Enable-AppvClientConnectionGroup`](#enable-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Enables a running connection group on the computer running the App-V client. | Out of scope |  |
| [`Enable-BcdElementBootDebug`](#enable-bcdelementbootdebug) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Enable-BcdElementBootEms`](#enable-bcdelementbootems) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Enable-BcdElementDebug`](#enable-bcdelementdebug) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Enable-BcdElementEms`](#enable-bcdelementems) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Enable-BcdElementEventLogging`](#enable-bcdelementeventlogging) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Enable-BcdElementHypervisorDebug`](#enable-bcdelementhypervisordebug) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Enable-ComputerRestore`](#enable-computerrestore) | Microsoft.PowerShell.Management | Both | None | Enables the System Restore feature on the specified file system drive. | Out of scope |  |
| [`Enable-JobTrigger`](#enable-jobtrigger) | PSScheduledJob | Both | None | Enables the job triggers of scheduled jobs. | Out of scope |  |
| [`Enable-LocalUser`](#enable-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Enables a local user account. | Out of scope |  |
| [`Enable-PSRemoting`](#enable-psremoting) | Microsoft.PowerShell.Core | Both | None | Configures the computer to receive remote commands. | Out of scope |  |
| [`Enable-PSSessionConfiguration`](#enable-pssessionconfiguration) | Microsoft.PowerShell.Core | Both | None | Enables the session configurations on the local computer. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Enable-ReFSDedup`](#enable-refsdedup) | Microsoft.ReFsDedup.Commands | Both | None | Enables data deduplication on a specified ReFS volume. | Out of scope |  |
| [`Enable-ScheduledJob`](#enable-scheduledjob) | PSScheduledJob | Both | None | Enables a scheduled job. | Out of scope |  |
| [`Enable-TlsCipherSuite`](#enable-tlsciphersuite) | TLS | Both | Syntax differs | Enables a TLS cipher suite. | Out of scope |  |
| [`Enable-TlsEccCurve`](#enable-tlsecccurve) | TLS | Both | Syntax differs | Enables Elliptic Curve Cryptography (ECC) cipher suites available for TLS. | Out of scope |  |
| [`Enable-TlsSessionTicketKey`](#enable-tlssessionticketkey) | TLS | Both | None | Configures a TLS server with a TLS session ticket key. | Out of scope |  |
| [`Enable-TpmAutoProvisioning`](#enable-tpmautoprovisioning) | TrustedPlatformModule | Both | None | Enables TPM auto-provisioning. | Out of scope |  |
| [`Enable-Uev`](#enable-uev) | UEV | 5.1 only | 5.1 only | Enables the UE-V service. | Out of scope |  |
| [`Enable-UevAppxPackage`](#enable-uevappxpackage) | UEV | 5.1 only | 5.1 only | Enables UE-V synchronization of Windows 8 apps. | Out of scope |  |
| [`Enable-UevTemplate`](#enable-uevtemplate) | UEV | 5.1 only | 5.1 only | Enables a settings location template. | Out of scope |  |
| [`Enable-WindowsErrorReporting`](#enable-windowserrorreporting) | WindowsErrorReporting | Both | None | Enables Windows Error Reporting. | Out of scope |  |
| [`Enable-WindowsOptionalFeature`](#enable-windowsoptionalfeature) | Dism | Both | None | Enables a feature in a Windows image. | Out of scope |  |
| [`Enable-WSManCredSSP`](#enable-wsmancredssp) | Microsoft.WSMan.Management | Both | None | Enables Credential Security Support Provider (CredSSP) authentication on a computer. | Out of scope |  |
| [`Expand-OsImage`](#expand-osimage) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Expand-WindowsCustomDataImage`](#expand-windowscustomdataimage) | Dism | Both | None | Expands a custom data image. | Out of scope |  |
| [`Expand-WindowsImage`](#expand-windowsimage) | Dism | Both | Syntax differs | Applies an image to a specified location. | Out of scope |  |
| [`Export-BcdStore`](#export-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Export-BinaryMiLog`](#export-binarymilog) | CimCmdlets | Both | None | Creates a binary encoded representation of an object or objects and stores it in a file. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Export-Certificate`](#export-certificate) | PKI | Both | None | Exports a certificate from a certificate store into a file. | Out of scope |  |
| [`Export-Console`](#export-console) | Microsoft.PowerShell.Core | 5.1 only | 5.1 only | Exports the names of snap-ins in the current session to a console file. | Out of scope |  |
| [`Export-Counter`](#export-counter) | Microsoft.PowerShell.Diagnostics | Both | None | Exports performance counter data to log files. | Out of scope |  |
| [`Export-OsImage`](#export-osimage) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Export-PfxCertificate`](#export-pfxcertificate) | PKI | Both | None | Exports a certificate or a PFXData object to a Personal Information Exchange (PFX) file. | Out of scope |  |
| [`Export-ProvisioningPackage`](#export-provisioningpackage) | Provisioning | Both | None | Extract the contents of a provisioning package. | Out of scope |  |
| [`Export-StartLayout`](#export-startlayout) | StartLayout | Both | None | Exports the layout of the Start screen. | Out of scope |  |
| [`Export-StartLayoutEdgeAssets`](#export-startlayoutedgeassets) | StartLayout | Both | None | Exports secondary tiles for Microsoft Edge that display a custom image. | Out of scope |  |
| [`Export-TlsSessionTicketKey`](#export-tlssessionticketkey) | TLS | Both | None | Exports a TLS session ticket key. | Out of scope |  |
| [`Export-Trace`](#export-trace) | Provisioning | Both | None | Exports an event trace log file for provisioning. | Out of scope |  |
| [`Export-UevConfiguration`](#export-uevconfiguration) | UEV | 5.1 only | 5.1 only | Exports the UE-V configuration. | Out of scope |  |
| [`Export-UevPackage`](#export-uevpackage) | UEV | 5.1 only | 5.1 only | Exports the settings stored in a settings package. | Out of scope |  |
| [`Export-WindowsCapabilitySource`](#export-windowscapabilitysource) | Dism | Both | None | Creates a custom FOD repository that includes packages that support the installation of the specified capabilities. See | Out of scope |  |
| [`Export-WindowsDriver`](#export-windowsdriver) | Dism | Both | None | Exports all third-party drivers from a Windows image to a destination folder. | Out of scope |  |
| [`Export-WindowsImage`](#export-windowsimage) | Dism | Both | Syntax differs | Exports a copy of the specified image to another image file. | Out of scope |  |
| [`Find-LapsADExtendedRights`](#find-lapsadextendedrights) | LAPS | Both | None | Queries Active Directory (AD) to find principals that have been granted permission to read Windows Local Administrator Password Solution (LAPS) password attributes. | Out of scope |  |
| [`Format-SecureBootUEFI`](#format-securebootuefi) | SecureBoot | Both | None | Formats certificates or hashes into a content object that is returned and creates a file that is ready to be signed. | Out of scope |  |
| [`Get-Acl`](#get-acl) | Microsoft.PowerShell.Security | Both | Syntax differs | Gets the security descriptor for a resource, such as a file or registry key. | Out of scope |  |
| [`Get-AppLockerFileInformation`](#get-applockerfileinformation) | AppLocker | Both | None | Gets the file information necessary to create AppLocker rules from a list of files or an event log. | Out of scope |  |
| [`Get-AppLockerPolicy`](#get-applockerpolicy) | AppLocker | Both | None | Gets the local, the effective, or a domain AppLocker policy. | Out of scope |  |
| [`Get-AppProvisionedSharedPackageContainer`](#get-appprovisionedsharedpackagecontainer) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-AppSharedPackageContainer`](#get-appsharedpackagecontainer) | Appx | Both | None | Gets information about the shared package container. | Out of scope |  |
| [`Get-AppvClientApplication`](#get-appvclientapplication) | AppvClient | 5.1 only | 5.1 only | Returns applications that are part of App-V Client Packages. | Out of scope |  |
| [`Get-AppvClientConfiguration`](#get-appvclientconfiguration) | AppvClient | 5.1 only | 5.1 only | Returns the configuration for the App-V client. | Out of scope |  |
| [`Get-AppvClientConnectionGroup`](#get-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Returns an App-V connection group object. | Out of scope |  |
| [`Get-AppvClientMode`](#get-appvclientmode) | AppvClient | 5.1 only | 5.1 only | Displays the mode for the App-V Client. | Out of scope |  |
| [`Get-AppvClientPackage`](#get-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Returns App-V Client Packages. | Out of scope |  |
| [`Get-AppvPublishingServer`](#get-appvpublishingserver) | AppvClient | 5.1 only | 5.1 only | Returns App-V Server objects. | Out of scope |  |
| [`Get-AppvStatus`](#get-appvstatus) | AppvClient | 5.1 only | 5.1 only | Gets the status of the App-V service. | Out of scope |  |
| [`Get-AppxDefaultVolume`](#get-appxdefaultvolume) | Appx | Both | None | Gets the default appx volume. | Out of scope |  |
| [`Get-AppxPackage`](#get-appxpackage) | Appx | Both | None | Gets a list of the app packages that are installed in a user profile. | Out of scope |  |
| [`Get-AppxPackageAutoUpdateSettings`](#get-appxpackageautoupdatesettings) | Appx | Both | None | Provides visibility to the settings configured for a particular Windows App. | Out of scope |  |
| [`Get-AppxPackageManifest`](#get-appxpackagemanifest) | Appx | Both | None | Gets the manifest of an app package. | Out of scope |  |
| [`Get-AppxProvisionedPackage`](#get-appxprovisionedpackage) | Dism | Both | None | Gets information about app packages (.appx) in an image that will be installed for each new user. | Out of scope |  |
| [`Get-AppxVolume`](#get-appxvolume) | Appx | Both | None | Gets appx volumes for the computer. | Out of scope |  |
| [`Get-AuthenticodeSignature`](#get-authenticodesignature) | Microsoft.PowerShell.Security | Both | None | Gets information about the Authenticode signature for a file. | Out of scope |  |
| [`Get-BcdEntry`](#get-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-BcdEntryDebugSettings`](#get-bcdentrydebugsettings) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-BcdEntryHypervisorSettings`](#get-bcdentryhypervisorsettings) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-BcdStore`](#get-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-BitsTransfer`](#get-bitstransfer) | BitsTransfer | Both | None | Gets the associated BitsJob object for an existing BITS transfer job. | Out of scope |  |
| [`Get-Certificate`](#get-certificate) | PKI | Both | None | Submits a certificate request to an enrollment server and installs the response or retrieves a certificate for a previously submitted request. | Out of scope |  |
| [`Get-CertificateAutoEnrollmentPolicy`](#get-certificateautoenrollmentpolicy) | PKI | Both | None | Retrieves certificate auto-enrollment policy settings. | Out of scope |  |
| [`Get-CertificateEnrollmentPolicyServer`](#get-certificateenrollmentpolicyserver) | PKI | Both | None | Returns all of the certificate enrollment policy server URL configurations. | Out of scope |  |
| [`Get-CertificateNotificationTask`](#get-certificatenotificationtask) | PKI | Both | None | Returns all registered certificate notification tasks. | Out of scope |  |
| [`Get-CimAssociatedInstance`](#get-cimassociatedinstance) | CimCmdlets | Both | Syntax differs | Retrieves the CIM instances that are connected to a specific CIM instance by an association. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Get-CimClass`](#get-cimclass) | CimCmdlets | Both | Syntax differs | Gets a list of CIM classes in a specific namespace. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Get-CimInstance`](#get-ciminstance) | CimCmdlets | Both | Syntax differs | Gets the CIM instances of a class from a CIM server. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Get-CimSession`](#get-cimsession) | CimCmdlets | Both | Syntax differs | Gets the CIM session objects from the current session. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Get-CIPolicy`](#get-cipolicy) | ConfigCI | Both | None | Gets the rules in a Code Integrity policy. | Out of scope |  |
| [`Get-CIPolicyIdInfo`](#get-cipolicyidinfo) | ConfigCI | Both | None | Displays Code Integrity policy information. | Out of scope |  |
| [`Get-CIPolicyInfo`](#get-cipolicyinfo) | ConfigCI | Both | None | This cmdlet is not supported. | Out of scope |  |
| [`Get-ComputerInfo`](#get-computerinfo) | Microsoft.PowerShell.Management | Both | None | Gets a consolidated object of system and operating system properties. | Go implementation | Reduced field set. |
| [`Get-ComputerRestorePoint`](#get-computerrestorepoint) | Microsoft.PowerShell.Management | Both | None | Gets the restore points on the local computer. | Out of scope |  |
| [`Get-ControlPanelItem`](#get-controlpanelitem) | Microsoft.PowerShell.Management | Both | None | Gets control panel items. | Out of scope |  |
| [`Get-Counter`](#get-counter) | Microsoft.PowerShell.Diagnostics | Both | None | Gets performance counter data from local and remote computers. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Get-DAPolicyChange`](#get-dapolicychange) | NetSecurity | Both | None | Gets a list of IP addresses that need to be added and deleted to an IPsec rule based on the differences detected between the IP addresses for the existing rule and the IP addresses derived from the input parameters, and creates a Windows PowerShell® script (.ps1) that updates the IPsec rule in the appropriate policy stores. | Out of scope |  |
| [`Get-DeliveryOptimizationLog`](#get-deliveryoptimizationlog) | DeliveryOptimization | Both | Syntax differs | - | Out of scope | No documentation accurately describes this command. |
| [`Get-DeliveryOptimizationLogAnalysis`](#get-deliveryoptimizationloganalysis) | DeliveryOptimization | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-EventLog`](#get-eventlog) | Microsoft.PowerShell.Management | Both | None | Gets the events in an event log, or a list of the event logs, on the local computer or remote computers. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Get-HotFix`](#get-hotfix) | Microsoft.PowerShell.Management | Both | None | Gets the hotfixes that are installed on local or remote computers. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Get-InstalledLanguage`](#get-installedlanguage) | LanguagePackManagement | Both | None | Returns information about the installed languages on a device. | Out of scope |  |
| [`Get-JobTrigger`](#get-jobtrigger) | PSScheduledJob | Both | None | Gets the job triggers of scheduled jobs. | Out of scope |  |
| [`Get-KdsConfiguration`](#get-kdsconfiguration) | Kds | Both | None | Retrieves the current configuration of the Microsoft Group KdsSvc from Active Directory. | Out of scope |  |
| [`Get-KdsRootKey`](#get-kdsrootkey) | Kds | Both | None | Retrieves a list of root key values stored by the Microsoft Group KdsSvc. | Out of scope |  |
| [`Get-LapsADPassword`](#get-lapsadpassword) | LAPS | Both | None | Queries Windows Local Administrator Password Solution (LAPS) credentials from Active Directory (AD) on a specified AD computer or domain controller object. | Out of scope |  |
| [`Get-LocalGroup`](#get-localgroup) | Microsoft.PowerShell.LocalAccounts | Both | None | Gets the local security groups. | Out of scope |  |
| [`Get-LocalGroupMember`](#get-localgroupmember) | Microsoft.PowerShell.LocalAccounts | Both | None | Gets members from a local group. | Out of scope |  |
| [`Get-LocalUser`](#get-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Gets local user accounts. | Out of scope |  |
| [`Get-NonRemovableAppsPolicy`](#get-nonremovableappspolicy) | Dism | Both | None | Returns a list of the app packages that are installed and configured as non-removable apps. | Out of scope |  |
| [`Get-OSConfiguration`](#get-osconfiguration) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-OsConfigurationDocument`](#get-osconfigurationdocument) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-OsConfigurationDocumentContent`](#get-osconfigurationdocumentcontent) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-OsConfigurationDocumentResult`](#get-osconfigurationdocumentresult) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-OsConfigurationProperty`](#get-osconfigurationproperty) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-OSConfigurationScenarioDefinition`](#get-osconfigurationscenariodefinition) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-OSConfigurationScenarioDefinitionInfo`](#get-osconfigurationscenariodefinitioninfo) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-PfxData`](#get-pfxdata) | PKI | Both | None | Extracts the content of a Personal Information Exchange (PFX) file into a structure without importing it to certificate store. | Out of scope |  |
| [`Get-PmemDedicatedMemory`](#get-pmemdedicatedmemory) | PersistentMemory | Both | None | Gets dedicated persistent memory. | Out of scope |  |
| [`Get-PmemDisk`](#get-pmemdisk) | PersistentMemory | Both | None | Gets persistent memory disks. | Out of scope |  |
| [`Get-PmemPhysicalDevice`](#get-pmemphysicaldevice) | PersistentMemory | Both | None | Gets the physical devices associated with persistent memory. | Out of scope |  |
| [`Get-PmemUnusedRegion`](#get-pmemunusedregion) | PersistentMemory | Both | None | Gets unused regions in persistent memory. | Out of scope |  |
| [`Get-ProcessMitigation`](#get-processmitigation) | ProcessMitigations | Both | None | Gets the current process mitigation settings, either from the registry, from a running process, or saves all to a XML. | Out of scope |  |
| [`Get-ProvisioningPackage`](#get-provisioningpackage) | Provisioning | Both | None | Gets information about the installed provisioning package. | Out of scope |  |
| [`Get-PSSessionCapability`](#get-pssessioncapability) | Microsoft.PowerShell.Core | Both | None | Gets the capabilities of a specific user on a constrained session configuration. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Get-PSSessionConfiguration`](#get-pssessionconfiguration) | Microsoft.PowerShell.Core | Both | None | Gets the registered session configurations on the computer. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Get-PSSnapin`](#get-pssnapin) | Microsoft.PowerShell.Core | 5.1 only | 5.1 only | Gets the Windows PowerShell snap-ins on the computer. | Out of scope |  |
| [`Get-RecoveryManagementPluginAltitude`](#get-recoverymanagementpluginaltitude) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-RecoveryManagementPluginInfo`](#get-recoverymanagementplugininfo) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-RecoveryManagementPlugins`](#get-recoverymanagementplugins) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-RecoveryRemoteManagementStatus`](#get-recoveryremotemanagementstatus) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-ReFSDedupSchedule`](#get-refsdedupschedule) | Microsoft.ReFsDedup.Commands | Both | None | Retrieves the deduplication schedule on a specified ReFS volume. | Out of scope |  |
| [`Get-ReFSDedupScrubSchedule`](#get-refsdedupscrubschedule) | Microsoft.ReFsDedup.Commands | Both | None | Retrieves the deduplication scrub schedule on the specified ReFS volume. | Out of scope |  |
| [`Get-ReFSDedupStatus`](#get-refsdedupstatus) | Microsoft.ReFsDedup.Commands | Both | None | Retrieves the status of data deduplication on a specified ReFS volume. | Out of scope |  |
| [`Get-ScheduledJob`](#get-scheduledjob) | PSScheduledJob | Both | None | Gets scheduled jobs on the local computer. | Out of scope |  |
| [`Get-ScheduledJobOption`](#get-scheduledjoboption) | PSScheduledJob | Both | None | Gets the job options of scheduled jobs. | Out of scope |  |
| [`Get-SecureBootPolicy`](#get-securebootpolicy) | SecureBoot | Both | None | Gets the publisher GUID and the policy version of the Secure Boot configuration policy. | Out of scope |  |
| [`Get-SecureBootSVN`](#get-securebootsvn) | SecureBoot | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Get-SecureBootUEFI`](#get-securebootuefi) | SecureBoot | Both | None | Gets the UEFI variable values related to Secure Boot. | Out of scope |  |
| [`Get-Service`](#get-service) | Microsoft.PowerShell.Management | Both | Description differs; Syntax differs | 5.1: Gets the services on a local or remote computer. / 7: Gets the services on the computer. | Mapped Linux (systemctl) |  |
| [`Get-SystemDriver`](#get-systemdriver) | ConfigCI | Both | None | Scans for drivers on the system. | Out of scope |  |
| [`Get-SystemPreferredUILanguage`](#get-systempreferreduilanguage) | LanguagePackManagement | Both | None | Returns the current System Preferred Language. | Out of scope |  |
| [`Get-TlsCipherSuite`](#get-tlsciphersuite) | TLS | Both | None | Gets the TLS cipher suites for a computer. | Out of scope |  |
| [`Get-TlsEccCurve`](#get-tlsecccurve) | TLS | Both | None | Gets the list of Elliptic Curve Cryptography (ECC) cipher suites available for TLS for a computer. | Out of scope |  |
| [`Get-Tpm`](#get-tpm) | TrustedPlatformModule | Both | None | Gets an object that contains information about a TPM. | Out of scope |  |
| [`Get-TpmEndorsementKeyInfo`](#get-tpmendorsementkeyinfo) | TrustedPlatformModule | Both | None | Gets information about the endorsement key and certificates of the TPM. | Out of scope |  |
| [`Get-TpmSupportedFeature`](#get-tpmsupportedfeature) | TrustedPlatformModule | Both | None | Verifies whether a TPM supports specified features. | Out of scope |  |
| [`Get-Transaction`](#get-transaction) | Microsoft.PowerShell.Management | Both | None | Gets the current (active) transaction. | Out of scope |  |
| [`Get-TroubleshootingPack`](#get-troubleshootingpack) | TroubleshootingPack | Both | None | Gets a troubleshooting pack or generates an answer file. | Out of scope |  |
| [`Get-TrustedProvisioningCertificate`](#get-trustedprovisioningcertificate) | Provisioning | Both | None | Lists all installed trusted provisioning certificates. | Out of scope |  |
| [`Get-UevAppxPackage`](#get-uevappxpackage) | UEV | 5.1 only | 5.1 only | Gets a list of Windows 8 apps and synchronization status. | Out of scope |  |
| [`Get-UevConfiguration`](#get-uevconfiguration) | UEV | 5.1 only | 5.1 only | Gets the UE-V configuration settings. | Out of scope |  |
| [`Get-UevStatus`](#get-uevstatus) | UEV | 5.1 only | 5.1 only | Gets the status of the UE-V service. | Out of scope |  |
| [`Get-UevTemplate`](#get-uevtemplate) | UEV | 5.1 only | 5.1 only | Gets settings location templates for UE-V. | Out of scope |  |
| [`Get-UevTemplateProgram`](#get-uevtemplateprogram) | UEV | 5.1 only | 5.1 only | Gets the information about programs defined by a settings location template. | Out of scope |  |
| [`Get-WheaMemoryPolicy`](#get-wheamemorypolicy) | Whea | Both | None | Gets the WHEA memory policies for a computer. | Out of scope |  |
| [`Get-WIMBootEntry`](#get-wimbootentry) | Dism | Both | None | Displays the Windows image file boot (WIMBoot) configuration entries for a specified disk volume. | Out of scope |  |
| [`Get-WinAcceptLanguageFromLanguageListOptOut`](#get-winacceptlanguagefromlanguagelistoptout) | International | Both | None | Gets the HTTP Accept Language from the Language List opt-out setting for the current user account. | Out of scope |  |
| [`Get-WinCultureFromLanguageListOptOut`](#get-winculturefromlanguagelistoptout) | International | Both | None | Gets the Culture from the language list opt-out setting for the current user account. | Out of scope |  |
| [`Get-WinDefaultInputMethodOverride`](#get-windefaultinputmethodoverride) | International | Both | None | Gets the default input method override setting for the current user account. | Out of scope |  |
| [`Get-WindowsCapability`](#get-windowscapability) | Dism | Both | None | Gets Windows capabilities for an image or a running operating system. | Out of scope |  |
| [`Get-WindowsDeveloperLicense`](#get-windowsdeveloperlicense) | WindowsDeveloperLicense | Both | None | Provides information about Developer Mode for the current computer. | Out of scope |  |
| [`Get-WindowsDriver`](#get-windowsdriver) | Dism | Both | None | Displays information about drivers in a Windows image. | Out of scope |  |
| [`Get-WindowsEdition`](#get-windowsedition) | Dism | Both | None | Gets edition information about a Windows image. | Out of scope |  |
| [`Get-WindowsErrorReporting`](#get-windowserrorreporting) | WindowsErrorReporting | Both | None | Retrieves the Windows Error Reporting status. | Out of scope |  |
| [`Get-WindowsImage`](#get-windowsimage) | Dism | Both | Syntax differs | Gets information about a Windows image in a WIM or VHD file. | Out of scope |  |
| [`Get-WindowsImageContent`](#get-windowsimagecontent) | Dism | Both | Syntax differs | Displays a list of the files and folders in a specified image. | Out of scope |  |
| [`Get-WindowsOptionalFeature`](#get-windowsoptionalfeature) | Dism | Both | None | Gets information about optional features in a Windows image. | Out of scope |  |
| [`Get-WindowsPackage`](#get-windowspackage) | Dism | Both | None | Gets information about packages in a Windows image. | Out of scope |  |
| [`Get-WindowsReservedStorageState`](#get-windowsreservedstoragestate) | Dism | Both | None | Gets the reserved storage state of the image. | Out of scope |  |
| [`Get-WindowsSearchSetting`](#get-windowssearchsetting) | WindowsSearch | Both | None | Gets the values of settings for Windows Search. | Out of scope |  |
| [`Get-WinEvent`](#get-winevent) | Microsoft.PowerShell.Diagnostics | Both | None | Gets events from event logs and event tracing log files on local and remote computers. | Out of scope |  |
| [`Get-WinHomeLocation`](#get-winhomelocation) | International | Both | None | Gets the Windows GeoID home location setting for the current user account. | Out of scope |  |
| [`Get-WinLanguageBarOption`](#get-winlanguagebaroption) | International | Both | None | Gets the language bar mode and language bar type for the current user account. | Out of scope |  |
| [`Get-WinSystemLocale`](#get-winsystemlocale) | International | Both | None | Gets the System-locale setting for the current computer. | Out of scope |  |
| [`Get-WinUILanguageOverride`](#get-winuilanguageoverride) | International | Both | None | Gets the Windows UI language override setting for the current user account. | Out of scope |  |
| [`Get-WinUserLanguageList`](#get-winuserlanguagelist) | International | Both | None | Gets the language list for the current user account. | Out of scope |  |
| [`Get-WmiObject`](#get-wmiobject) | Microsoft.PowerShell.Management | Both | None | Gets instances of Windows Management Instrumentation (WMI) classes or information about the available classes. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Get-WSManCredSSP`](#get-wsmancredssp) | Microsoft.WSMan.Management | Both | None | Gets the Credential Security Support Provider-related configuration for the client. | Out of scope |  |
| [`Get-WSManInstance`](#get-wsmaninstance) | Microsoft.WSMan.Management | Both | None | Displays management information for a resource instance specified by a Resource URI. | Out of scope |  |
| [`Import-BcdStore`](#import-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Import-BinaryMiLog`](#import-binarymilog) | CimCmdlets | Both | None | Used to re-create the saved objects based on the contents of an export file. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Import-Certificate`](#import-certificate) | PKI | Both | None | Imports one or more certificates into a certificate store. | Out of scope |  |
| [`Import-Counter`](#import-counter) | Microsoft.PowerShell.Diagnostics | Both | None | Imports performance counter log files and creates the objects that represent each counter sample in the log. | Out of scope |  |
| [`Import-PfxCertificate`](#import-pfxcertificate) | PKI | Both | None | Imports certificates and private keys from a Personal Information Exchange (PFX) file to the destination store. | Out of scope |  |
| [`Import-StartLayout`](#import-startlayout) | StartLayout | Both | None | Imports the layout of the Start into a mounted Windows image. | Out of scope |  |
| [`Import-TpmOwnerAuth`](#import-tpmownerauth) | TrustedPlatformModule | Both | None | Imports a TPM owner authorization value to the registry. | Out of scope |  |
| [`Import-UevConfiguration`](#import-uevconfiguration) | UEV | 5.1 only | 5.1 only | Imports the UE-V configuration. | Out of scope |  |
| [`Initialize-PmemPhysicalDevice`](#initialize-pmemphysicaldevice) | PersistentMemory | Both | None | Initializes the label storage area on a physical persistent memory device. | Out of scope |  |
| [`Initialize-Tpm`](#initialize-tpm) | TrustedPlatformModule | Both | None | Performs part of the provisioning process for a TPM. | Out of scope |  |
| [`Install-Language`](#install-language) | LanguagePackManagement | Both | None | Installs a language onto a device. | Out of scope |  |
| [`Install-ProvisioningPackage`](#install-provisioningpackage) | Provisioning | Both | None | Install .PPKG package onto the local machine. | Out of scope |  |
| [`Install-TrustedProvisioningCertificate`](#install-trustedprovisioningcertificate) | Provisioning | Both | None | Adds a certificate to the Trusted Certificate Store. | Out of scope |  |
| [`Invoke-CimMethod`](#invoke-cimmethod) | CimCmdlets | Both | Syntax differs | Invokes a method of a CIM class. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Invoke-CommandInDesktopPackage`](#invoke-commandindesktoppackage) | Appx | Both | None | A debugging tool that creates a new process in the context of a packaged app. | Out of scope |  |
| [`Invoke-DscResource`](#invoke-dscresource) | PSDesiredStateConfiguration | Both | None | Runs a method of a specified PowerShell Desired State Configuration (DSC) resource. | Out of scope |  |
| [`Invoke-LapsPolicyProcessing`](#invoke-lapspolicyprocessing) | LAPS | Both | None | Causes Windows Local Administrator Password Solution (LAPS) to process the currently configured policy. | Out of scope |  |
| [`Invoke-TroubleshootingPack`](#invoke-troubleshootingpack) | TroubleshootingPack | Both | None | Runs a troubleshooting pack. | Out of scope |  |
| [`Invoke-WmiMethod`](#invoke-wmimethod) | Microsoft.PowerShell.Management | Both | None | Calls WMI methods. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Invoke-WSManAction`](#invoke-wsmanaction) | Microsoft.WSMan.Management | Both | None | Invokes an action on the object that is specified by the Resource URI and by the selectors. | Out of scope |  |
| [`Join-DtcDiagnosticResourceManager`](#join-dtcdiagnosticresourcemanager) | MsDtc | Both | None | Enlists a diagnostic Resource Manager for a transaction object. | Out of scope |  |
| [`Limit-EventLog`](#limit-eventlog) | Microsoft.PowerShell.Management | Both | None | Sets the event log properties that limit the size of the event log and the age of its entries. | Out of scope |  |
| [`Merge-CIPolicy`](#merge-cipolicy) | ConfigCI | Both | None | Combines the rules in several Code Integrity policy files. | Out of scope |  |
| [`Mount-AppvClientConnectionGroup`](#mount-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Streams the contents of packages to the local disk. | Out of scope |  |
| [`Mount-AppvClientPackage`](#mount-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Loads a package into the App-V cache. | Out of scope |  |
| [`Mount-AppxVolume`](#mount-appxvolume) | Appx | Both | None | Mounts an appx volume. | Out of scope |  |
| [`Mount-WindowsImage`](#mount-windowsimage) | Dism | Both | Syntax differs | Mounts a Windows image in a WIM or VHD file to a directory on the local computer. | Out of scope |  |
| [`Move-AppxPackage`](#move-appxpackage) | Appx | Both | None | Moves a package from its current location to another appx volume. | Out of scope |  |
| [`New-AppLockerPolicy`](#new-applockerpolicy) | AppLocker | Both | None | Creates a new AppLocker policy from a list of file information and other rule creation options. | Out of scope |  |
| [`New-BcdEntry`](#new-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`New-BcdStore`](#new-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`New-CertificateNotificationTask`](#new-certificatenotificationtask) | PKI | Both | None | Creates a new task in the Task Scheduler that will be triggered when a certificate is replaced, expired, or about to expired. | Out of scope |  |
| [`New-CimInstance`](#new-ciminstance) | CimCmdlets | Both | Syntax differs | Creates a CIM instance. | Out of scope | Windows-only (CIM/WMI stack) |
| [`New-CimSession`](#new-cimsession) | CimCmdlets | Both | Syntax differs | Creates a CIM session. | Out of scope | Windows-only (CIM/WMI stack) |
| [`New-CimSessionOption`](#new-cimsessionoption) | CimCmdlets | Both | Syntax differs | Specifies advanced options for the New-CimSession cmdlet. | Out of scope | Windows-only (CIM/WMI stack) |
| [`New-CIPolicy`](#new-cipolicy) | ConfigCI | Both | None | Creates a Code Integrity policy as an .xml file. | Out of scope |  |
| [`New-CIPolicyRule`](#new-cipolicyrule) | ConfigCI | Both | None | Generates Code Integrity policy rules for user mode code and drivers. | Out of scope |  |
| [`New-DtcDiagnosticTransaction`](#new-dtcdiagnostictransaction) | MsDtc | Both | None | Creates a new transaction in a Transaction Manager on the local computer. | Out of scope |  |
| [`New-EventLog`](#new-eventlog) | Microsoft.PowerShell.Management | Both | None | Creates a new event log and a new event source on a local or remote computer. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`New-FileCatalog`](#new-filecatalog) | Microsoft.PowerShell.Security | Both | None | Creates a Windows catalog file containing cryptographic hashes for files and folders in the specified paths. | Out of scope |  |
| [`New-JobTrigger`](#new-jobtrigger) | PSScheduledJob | Both | None | Creates a job trigger for a scheduled job. | Out of scope |  |
| [`New-LocalGroup`](#new-localgroup) | Microsoft.PowerShell.LocalAccounts | Both | None | Creates a local security group. | Out of scope |  |
| [`New-LocalUser`](#new-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Creates a local user account. | Out of scope |  |
| [`New-NetIPsecAuthProposal`](#new-netipsecauthproposal) | NetSecurity | Both | None | Creates a main mode authentication proposal that specifies a suite of authentication protocols to offer in IPsec main mode negotiations with other computers. | Out of scope |  |
| [`New-NetIPsecMainModeCryptoProposal`](#new-netipsecmainmodecryptoproposal) | NetSecurity | Both | None | Creates a main mode cryptographic proposal that specifies a suite of cryptographic protocols to offer in IPsec main mode negotiations with other computers. | Out of scope |  |
| [`New-NetIPsecQuickModeCryptoProposal`](#new-netipsecquickmodecryptoproposal) | NetSecurity | Both | Syntax differs | Creates a quick mode cryptographic proposal that specifies a suite of cryptographic protocols to offer in IPsec quick mode negotiations with other computers. | Out of scope |  |
| [`New-PmemDedicatedMemory`](#new-pmemdedicatedmemory) | PersistentMemory | Both | None | Creates dedicated persistent memory in the specified region. | Out of scope |  |
| [`New-PmemDisk`](#new-pmemdisk) | PersistentMemory | Both | None | Creates a persistent memory disk in an unused persistent memory region. | Out of scope |  |
| [`New-ProvisioningRepro`](#new-provisioningrepro) | Provisioning | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`New-PSWorkflowExecutionOption`](#new-psworkflowexecutionoption) | PSWorkflow | Both | None | Creates an object that contains session configuration options for workflow sessions. | Out of scope |  |
| [`New-ScheduledJobOption`](#new-scheduledjoboption) | PSScheduledJob | Both | None | Creates an object that contains advanced options for a scheduled job. | Out of scope |  |
| [`New-SelfSignedCertificate`](#new-selfsignedcertificate) | PKI | Both | None | Creates a new self-signed certificate for testing purposes. | Out of scope |  |
| [`New-Service`](#new-service) | Microsoft.PowerShell.Management | Both | Syntax differs | Creates a new Windows service. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`New-TlsSessionTicketKey`](#new-tlssessionticketkey) | TLS | Both | None | Creates a TLS session ticket key configuration file. | Out of scope |  |
| [`New-WebServiceProxy`](#new-webserviceproxy) | Microsoft.PowerShell.Management | Both | None | Creates a Web service proxy object that lets you use and manage the Web service in PowerShell. | Out of scope |  |
| [`New-WindowsCustomImage`](#new-windowscustomimage) | Dism | Both | None | Captures an image of customized or serviced Windows components on a Windows Image File Boot (WIMBoot) configured device. | Out of scope |  |
| [`New-WindowsImage`](#new-windowsimage) | Dism | Both | None | Captures an image of a drive to a new WIM file. | Out of scope |  |
| [`New-WinEvent`](#new-winevent) | Microsoft.PowerShell.Diagnostics | Both | None | Creates a new Windows event for the specified event provider. | Out of scope |  |
| [`New-WinUserLanguageList`](#new-winuserlanguagelist) | International | Both | None | Instantiates a new language list object. | Out of scope |  |
| [`New-WSManInstance`](#new-wsmaninstance) | Microsoft.WSMan.Management | Both | None | Creates a new instance of a management resource. | Out of scope |  |
| [`New-WSManSessionOption`](#new-wsmansessionoption) | Microsoft.WSMan.Management | Both | None | Creates session option hash table to use as input parameters for WS-Management cmdlets. | Out of scope |  |
| [`Optimize-AppxProvisionedPackages`](#optimize-appxprovisionedpackages) | Dism | Both | None | Optimizes the total file size of provisioned packages on the image by replacing identical files with hard links. | Out of scope |  |
| [`Optimize-WindowsImage`](#optimize-windowsimage) | Dism | Both | None | Configures a Windows image with specified optimizations. | Out of scope |  |
| [`Out-GridView`](#out-gridview) | Microsoft.PowerShell.Utility | Both | None | Sends output to an interactive table in a separate window. | Out of scope | GUI and printing |
| [`Out-Printer`](#out-printer) | Microsoft.PowerShell.Utility | Both | None | Sends output to a printer. | Out of scope | GUI and printing |
| [`Publish-AppvClientPackage`](#publish-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Publishes the App-V package. | Out of scope |  |
| [`Publish-DscConfiguration`](#publish-dscconfiguration) | PSDesiredStateConfiguration | Both | None | Publishes a DSC configuration to a set of computers. | Out of scope |  |
| [`Receive-DtcDiagnosticTransaction`](#receive-dtcdiagnostictransaction) | MsDtc | Both | None | Propagates a transaction from a given diagnostic Resource Manager. | Out of scope |  |
| [`Receive-PSSession`](#receive-pssession) | Microsoft.PowerShell.Core | Both | Syntax differs | Gets results of commands in disconnected sessions | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Register-CimIndicationEvent`](#register-cimindicationevent) | CimCmdlets | Both | Syntax differs | Subscribes to indications using a filter expression or a query expression. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Register-PSSessionConfiguration`](#register-pssessionconfiguration) | Microsoft.PowerShell.Core | Both | Syntax differs | Creates and registers a new session configuration. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Register-RecoveryManagementPlugin`](#register-recoverymanagementplugin) | Dism | Both | Syntax differs | - | Out of scope | No documentation accurately describes this command. |
| [`Register-ScheduledJob`](#register-scheduledjob) | PSScheduledJob | Both | None | Creates a scheduled job. | Out of scope |  |
| [`Register-UevTemplate`](#register-uevtemplate) | UEV | 5.1 only | 5.1 only | Registers a settings location template with UE-V. | Out of scope |  |
| [`Register-WmiEvent`](#register-wmievent) | Microsoft.PowerShell.Management | Both | None | Subscribes to a Windows Management Instrumentation (WMI) event. | Out of scope |  |
| [`Remove-AppProvisionedSharedPackageContainer`](#remove-appprovisionedsharedpackagecontainer) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Remove-AppSharedPackageContainer`](#remove-appsharedpackagecontainer) | Appx | Both | None | Removes the shared package container. | Out of scope |  |
| [`Remove-AppvClientConnectionGroup`](#remove-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Deletes an App-V connection group on the client. | Out of scope |  |
| [`Remove-AppvClientPackage`](#remove-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Removes the package from a computer. | Out of scope |  |
| [`Remove-AppvPublishingServer`](#remove-appvpublishingserver) | AppvClient | 5.1 only | 5.1 only | Removes an App-V publishing server. | Out of scope |  |
| [`Remove-AppxPackage`](#remove-appxpackage) | Appx | Both | None | Removes an app package from one or more user accounts. | Out of scope |  |
| [`Remove-AppxPackageAutoUpdateSettings`](#remove-appxpackageautoupdatesettings) | Appx | Both | None | Removes settings configured for a particular Windows app. | Out of scope |  |
| [`Remove-AppxProvisionedPackage`](#remove-appxprovisionedpackage) | Dism | Both | None | Removes an app package (.appx) from a Windows image. | Out of scope |  |
| [`Remove-AppxVolume`](#remove-appxvolume) | Appx | Both | None | Removes an appx volume. | Out of scope |  |
| [`Remove-BcdElement`](#remove-bcdelement) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Remove-BcdEntry`](#remove-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Remove-BitsTransfer`](#remove-bitstransfer) | BitsTransfer | Both | None | Cancels a BITS transfer job. | Out of scope |  |
| [`Remove-CertificateEnrollmentPolicyServer`](#remove-certificateenrollmentpolicyserver) | PKI | Both | None | Removes an enrollment policy server and the URL of the enrollment policy server from the current user or local computer configuration. | Out of scope |  |
| [`Remove-CertificateNotificationTask`](#remove-certificatenotificationtask) | PKI | Both | None | Removes a certificate notification task from Task Scheduler. | Out of scope |  |
| [`Remove-CimInstance`](#remove-ciminstance) | CimCmdlets | Both | Syntax differs | Removes a CIM instance from a computer. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Remove-CimSession`](#remove-cimsession) | CimCmdlets | Both | Syntax differs | Removes one or more CIM sessions. | Out of scope | Windows-only (CIM/WMI stack) |
| [`Remove-CIPolicyRule`](#remove-cipolicyrule) | ConfigCI | Both | None | This cmdlet is not supported. | Out of scope |  |
| [`Remove-Computer`](#remove-computer) | Microsoft.PowerShell.Management | Both | None | Removes the local computer from its domain. | Out of scope |  |
| [`Remove-EventLog`](#remove-eventlog) | Microsoft.PowerShell.Management | Both | None | Deletes an event log or unregisters an event source. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Remove-JobTrigger`](#remove-jobtrigger) | PSScheduledJob | Both | None | Delete job triggers from scheduled jobs. | Out of scope |  |
| [`Remove-LocalGroup`](#remove-localgroup) | Microsoft.PowerShell.LocalAccounts | Both | None | Deletes local security groups. | Out of scope |  |
| [`Remove-LocalGroupMember`](#remove-localgroupmember) | Microsoft.PowerShell.LocalAccounts | Both | None | Removes members from a local group. | Out of scope |  |
| [`Remove-LocalUser`](#remove-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Deletes local user accounts. | Out of scope |  |
| [`Remove-OsConfigurationDocument`](#remove-osconfigurationdocument) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Remove-OSConfigurationScenarioDefinition`](#remove-osconfigurationscenariodefinition) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Remove-PmemDedicatedMemory`](#remove-pmemdedicatedmemory) | PersistentMemory | Both | None | Gets dedicated persistent memory. | Out of scope |  |
| [`Remove-PmemDisk`](#remove-pmemdisk) | PersistentMemory | Both | None | Removes persistent memory disks. | Out of scope |  |
| [`Remove-PSSnapin`](#remove-pssnapin) | Microsoft.PowerShell.Core | 5.1 only | 5.1 only | Removes Windows PowerShell snap-ins from the current session. | Out of scope |  |
| [`Remove-RecoveryManagementPluginAltitude`](#remove-recoverymanagementpluginaltitude) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Remove-Service`](#remove-service) | Microsoft.PowerShell.Management | 7 only | 7 only | Removes a Windows service. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Remove-WindowsCapability`](#remove-windowscapability) | Dism | Both | None | Uninstalls a Windows capability package from an image. | Out of scope |  |
| [`Remove-WindowsDriver`](#remove-windowsdriver) | Dism | Both | None | Removes a driver from an offline Windows image. | Out of scope |  |
| [`Remove-WindowsImage`](#remove-windowsimage) | Dism | Both | Syntax differs | Deletes the specified volume image from a WIM file that has multiple volume images. | Out of scope |  |
| [`Remove-WindowsPackage`](#remove-windowspackage) | Dism | Both | None | Removes a package from a Windows image. | Out of scope |  |
| [`Remove-WmiObject`](#remove-wmiobject) | Microsoft.PowerShell.Management | Both | None | Deletes an instance of an existing Windows Management Instrumentation (WMI) class. | Out of scope |  |
| [`Remove-WSManInstance`](#remove-wsmaninstance) | Microsoft.WSMan.Management | Both | None | Deletes a management resource instance. | Out of scope |  |
| [`Rename-Computer`](#rename-computer) | Microsoft.PowerShell.Management | Both | Syntax differs | Renames a computer. | Mapped Linux (sudo reboot / shutdown / hostnamectl) |  |
| [`Rename-LocalGroup`](#rename-localgroup) | Microsoft.PowerShell.LocalAccounts | Both | None | Renames a local security group. | Out of scope |  |
| [`Rename-LocalUser`](#rename-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Renames a local user account. | Out of scope |  |
| [`Repair-AppvClientConnectionGroup`](#repair-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Resets the user package settings for the connection group. | Out of scope |  |
| [`Repair-AppvClientPackage`](#repair-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Resets the user settings of a package. | Out of scope |  |
| [`Repair-UevTemplateIndex`](#repair-uevtemplateindex) | UEV | 5.1 only | 5.1 only | Repairs a corrupted UE-V template index. | Out of scope |  |
| [`Repair-WindowsImage`](#repair-windowsimage) | Dism | Both | None | Repairs a Windows image in a WIM or VHD file. | Out of scope |  |
| [`Reset-AppSharedPackageContainer`](#reset-appsharedpackagecontainer) | Appx | Both | None | Destroys all the application data of the container. | Out of scope |  |
| [`Reset-AppxPackage`](#reset-appxpackage) | Appx | Both | None | Restores the Windows app to its initial configuration. | Out of scope |  |
| [`Reset-ComputerMachinePassword`](#reset-computermachinepassword) | Microsoft.PowerShell.Management | Both | None | Resets the machine account password for the computer. | Out of scope |  |
| [`Reset-LapsPassword`](#reset-lapspassword) | LAPS | Both | None | Causes Windows Local Administrator Password Solution (LAPS) to immediately rotate the password for the currently managed local account. | Out of scope |  |
| [`Resolve-DnsName`](#resolve-dnsname) | DnsClient | Both | None | Performs a DNS name query resolution for the specified name. | Out of scope |  |
| [`Restart-Service`](#restart-service) | Microsoft.PowerShell.Management | Both | None | Stops and then starts one or more services. | Mapped Linux (systemctl start/stop/restart) |  |
| [`Restore-Computer`](#restore-computer) | Microsoft.PowerShell.Management | Both | None | Starts a system restore on the local computer. | Out of scope |  |
| [`Restore-UevBackup`](#restore-uevbackup) | UEV | 5.1 only | 5.1 only | Applies backed up settings from another computer to this computer. | Out of scope |  |
| [`Restore-UevUserSetting`](#restore-uevusersetting) | UEV | 5.1 only | 5.1 only | Sets a restore flag for the user settings. | Out of scope |  |
| [`Resume-BitsTransfer`](#resume-bitstransfer) | BitsTransfer | Both | None | Resumes a BITS transfer job. | Out of scope |  |
| [`Resume-Job`](#resume-job) | Microsoft.PowerShell.Core | 5.1 only | 5.1 only | Restarts a suspended job. | Out of scope |  |
| [`Resume-ProvisioningSession`](#resume-provisioningsession) | Provisioning | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Resume-ReFSDedupSchedule`](#resume-refsdedupschedule) | Microsoft.ReFsDedup.Commands | Both | None | Resumes the deduplication schedule on a specified ReFS volume. | Out of scope |  |
| [`Resume-Service`](#resume-service) | Microsoft.PowerShell.Management | Both | None | Resumes one or more suspended (paused) services. | Mapped Linux (systemctl start/stop/restart) |  |
| [`Save-OsImage`](#save-osimage) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Save-SoftwareInventory`](#save-softwareinventory) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Save-WindowsImage`](#save-windowsimage) | Dism | Both | None | Applies changes made to a mounted image to its WIM or VHD file. | Out of scope |  |
| [`Send-AppvClientReport`](#send-appvclientreport) | AppvClient | 5.1 only | 5.1 only | Sends reporting data from the client. | Out of scope |  |
| [`Send-DtcDiagnosticTransaction`](#send-dtcdiagnostictransaction) | MsDtc | Both | None | Propagates a transaction to a specified diagnostic Resource Manager. | Out of scope |  |
| [`Set-Acl`](#set-acl) | Microsoft.PowerShell.Security | Both | Syntax differs | Changes the security descriptor of a specified item, such as a file or a registry key. | Out of scope |  |
| [`Set-AppBackgroundTaskResourcePolicy`](#set-appbackgroundtaskresourcepolicy) | AppBackgroundTask | Both | None | Configures the use of the global pool by background tasks. | Out of scope |  |
| [`Set-AppLockerPolicy`](#set-applockerpolicy) | AppLocker | Both | None | Sets the AppLocker policy for the specified GPO. | Out of scope |  |
| [`Set-AppvClientConfiguration`](#set-appvclientconfiguration) | AppvClient | 5.1 only | 5.1 only | Applies configuration settings to the App-V Client. | Out of scope |  |
| [`Set-AppvClientMode`](#set-appvclientmode) | AppvClient | 5.1 only | 5.1 only | Sets the mode in which the client runs. | Out of scope |  |
| [`Set-AppvClientPackage`](#set-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Configures an App-V Client Package. | Out of scope |  |
| [`Set-AppvPublishingServer`](#set-appvpublishingserver) | AppvClient | 5.1 only | 5.1 only | Modifies properties of an App-V Publishing Server. | Out of scope |  |
| [`Set-AppxDefaultVolume`](#set-appxdefaultvolume) | Appx | Both | None | Specifies a default appx volume. | Out of scope |  |
| [`Set-AppxPackageAutoUpdateSettings`](#set-appxpackageautoupdatesettings) | Appx | Both | Syntax differs | Configures a specific Windows App's Auto Update and Repair settings. | Out of scope |  |
| [`Set-AppXProvisionedDataFile`](#set-appxprovisioneddatafile) | Dism | Both | None | Adds custom data into the specified app (.appx) package that has been provisioned in a Windows image. | Out of scope |  |
| [`Set-AuthenticodeSignature`](#set-authenticodesignature) | Microsoft.PowerShell.Security | Both | None | Adds an Authenticode signature to a PowerShell script or other file. | Out of scope |  |
| [`Set-BcdBootDefault`](#set-bcdbootdefault) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdBootDisplayOrder`](#set-bcdbootdisplayorder) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdBootSequence`](#set-bcdbootsequence) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdBootTimeout`](#set-bcdboottimeout) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdBootToolsDisplayOrder`](#set-bcdboottoolsdisplayorder) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdDebugSettings`](#set-bcddebugsettings) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdElement`](#set-bcdelement) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BcdHypervisorSettings`](#set-bcdhypervisorsettings) | Microsoft.Windows.Bcd.Cmdlets | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-BitsTransfer`](#set-bitstransfer) | BitsTransfer | Both | None | Modifies the properties of an existing BITS transfer job. | Out of scope |  |
| [`Set-CertificateAutoEnrollmentPolicy`](#set-certificateautoenrollmentpolicy) | PKI | Both | None | Sets local certificate auto-enrollment policy. | Out of scope |  |
| [`Set-CimInstance`](#set-ciminstance) | CimCmdlets | Both | Syntax differs | Modifies a CIM instance on a CIM server by calling the ModifyInstance method of the CIM class. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Set-CIPolicyIdInfo`](#set-cipolicyidinfo) | ConfigCI | Both | None | Modifies the name and ID of a Code Integrity policy. | Out of scope |  |
| [`Set-CIPolicySetting`](#set-cipolicysetting) | ConfigCI | Both | None | Modifies the SecureSettings within the Code Integrity policy. | Out of scope |  |
| [`Set-CIPolicyVersion`](#set-cipolicyversion) | ConfigCI | Both | None | Updates the version number of the policy. | Out of scope |  |
| [`Set-Culture`](#set-culture) | International | Both | None | Sets the user culture for the current user account. | Out of scope |  |
| [`Set-DscLocalConfigurationManager`](#set-dsclocalconfigurationmanager) | PSDesiredStateConfiguration | Both | None | Applies Local Configuration Manager (LCM) settings to nodes. | Out of scope |  |
| [`Set-HVCIOptions`](#set-hvcioptions) | ConfigCI | Both | None | Modifies hypervisor Code Integrity options for a policy. | Out of scope |  |
| [`Set-JobTrigger`](#set-jobtrigger) | PSScheduledJob | Both | None | Changes the job trigger of a scheduled job. | Out of scope |  |
| [`Set-KdsConfiguration`](#set-kdsconfiguration) | Kds | Both | None | Sets the configuration of Microsoft Group KdsSvc. | Out of scope |  |
| [`Set-LapsADAuditing`](#set-lapsadauditing) | LAPS | Both | None | Configures an Active Directory (AD) Organizational Unit (OU) to enable auditing on the Windows Local Administrator Password Solution (LAPS) password schema attributes. | Out of scope |  |
| [`Set-LapsADComputerSelfPermission`](#set-lapsadcomputerselfpermission) | LAPS | Both | None | Configures permissions on an Active Directory (AD) Organizational Unit (OU) to enable computers in that OU to update their Windows Local Administrator Password Solution (LAPS) passwords. | Out of scope |  |
| [`Set-LapsADPasswordExpirationTime`](#set-lapsadpasswordexpirationtime) | LAPS | Both | None | Sets the Windows Local Administrator Password Solution (LAPS) password expiration timestamp on an Active Directory (AD) computer or domain controller object. | Out of scope |  |
| [`Set-LapsADReadPasswordPermission`](#set-lapsadreadpasswordpermission) | LAPS | Both | None | Configures security on an Active Directory (AD) Organizational Unit (OU) to grant specific users or groups permission to query Windows Local Administrator Password Solution (LAPS) passwords. | Out of scope |  |
| [`Set-LapsADResetPasswordPermission`](#set-lapsadresetpasswordpermission) | LAPS | Both | None | Configures security on an Active Directory (AD) Organizational Unit (OU) to grant specific users or groups permission to set the Windows Local Administrator Password Solution (LAPS) password | Out of scope |  |
| [`Set-LocalGroup`](#set-localgroup) | Microsoft.PowerShell.LocalAccounts | Both | None | Changes a local security group. | Out of scope |  |
| [`Set-LocalUser`](#set-localuser) | Microsoft.PowerShell.LocalAccounts | Both | None | Modifies a local user account. | Out of scope |  |
| [`Set-NonRemovableAppsPolicy`](#set-nonremovableappspolicy) | Dism | Both | None | Sets an app package as non-removable (can not be uninstalled). | Out of scope |  |
| [`Set-OsConfigurationDocument`](#set-osconfigurationdocument) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-OsConfigurationProperty`](#set-osconfigurationproperty) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-OSConfigurationScenarioDefinition`](#set-osconfigurationscenariodefinition) | OsConfiguration | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-ProcessMitigation`](#set-processmitigation) | ProcessMitigations | Both | None | Commands to enable and disable process mitigations or set them in bulk from an XML file. | Out of scope |  |
| [`Set-PSSessionConfiguration`](#set-pssessionconfiguration) | Microsoft.PowerShell.Core | Both | None | Changes the properties of a registered session configuration. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Set-RecoveryManagementPluginAltitude`](#set-recoverymanagementpluginaltitude) | Dism | Both | Syntax differs | - | Out of scope | No documentation accurately describes this command. |
| [`Set-RecoveryRemoteManagementStatus`](#set-recoveryremotemanagementstatus) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Set-ReFSDedupSchedule`](#set-refsdedupschedule) | Microsoft.ReFsDedup.Commands | Both | None | Sets the deduplication schedule on a specified ReFS volume. | Out of scope |  |
| [`Set-ReFSDedupScrubSchedule`](#set-refsdedupscrubschedule) | Microsoft.ReFsDedup.Commands | Both | None | Sets the deduplication scrub schedule on the specified ReFS volume. | Out of scope |  |
| [`Set-RuleOption`](#set-ruleoption) | ConfigCI | Both | None | Modifies rule options in a Code Integrity policy. | Out of scope |  |
| [`Set-ScheduledJob`](#set-scheduledjob) | PSScheduledJob | Both | None | Changes scheduled jobs. | Out of scope |  |
| [`Set-ScheduledJobOption`](#set-scheduledjoboption) | PSScheduledJob | Both | None | Changes the job options of a scheduled job. | Out of scope |  |
| [`Set-SecureBootUEFI`](#set-securebootuefi) | SecureBoot | Both | None | Sets the Secure Boot-related UEFI variables. | Out of scope |  |
| [`Set-Service`](#set-service) | Microsoft.PowerShell.Management | Both | Syntax differs | Starts, stops, and suspends a service, and changes its properties. | Mapped Linux (systemctl start/stop/restart) |  |
| [`Set-SystemPreferredUILanguage`](#set-systempreferreduilanguage) | LanguagePackManagement | Both | None | Sets the provided language as the System Preferred UI Language. | Out of scope |  |
| [`Set-TimeZone`](#set-timezone) | Microsoft.PowerShell.Management | Both | None | Sets the system time zone to a specified time zone. | Mapped Linux (timedatectl) |  |
| [`Set-TpmOwnerAuth`](#set-tpmownerauth) | TrustedPlatformModule | Both | Syntax differs | Changes the TPM owner authorization value. | Out of scope |  |
| [`Set-UevConfiguration`](#set-uevconfiguration) | UEV | 5.1 only | 5.1 only | Modifies UE-V configuration settings. | Out of scope |  |
| [`Set-UevTemplateProfile`](#set-uevtemplateprofile) | UEV | 5.1 only | 5.1 only | Modifies which profile to associate with an individual template. | Out of scope |  |
| [`Set-WheaMemoryPolicy`](#set-wheamemorypolicy) | Whea | Both | Syntax differs | Sets the WHEA memory policy for a computer. | Out of scope |  |
| [`Set-WinAcceptLanguageFromLanguageListOptOut`](#set-winacceptlanguagefromlanguagelistoptout) | International | Both | None | Sets the HTTP Accept Language from the Language List opt-out setting for the current user account. | Out of scope |  |
| [`Set-WinCultureFromLanguageListOptOut`](#set-winculturefromlanguagelistoptout) | International | Both | None | Sets the Culture from language list opt out setting for the current user account. | Out of scope |  |
| [`Set-WinDefaultInputMethodOverride`](#set-windefaultinputmethodoverride) | International | Both | None | Sets the default input method override for the current user account. | Out of scope |  |
| [`Set-WindowsEdition`](#set-windowsedition) | Dism | Both | None | Changes a Windows image to a higher edition. | Out of scope |  |
| [`Set-WindowsProductKey`](#set-windowsproductkey) | Dism | Both | None | Sets the product key for the Windows image. | Out of scope |  |
| [`Set-WindowsReservedStorageState`](#set-windowsreservedstoragestate) | Dism | Both | None | Sets the reserved storage state of the image. | Out of scope |  |
| [`Set-WindowsSearchSetting`](#set-windowssearchsetting) | WindowsSearch | Both | None | Modifies values that control Windows Search. | Out of scope |  |
| [`Set-WinHomeLocation`](#set-winhomelocation) | International | Both | None | Sets the home location setting for the current user account. | Out of scope |  |
| [`Set-WinLanguageBarOption`](#set-winlanguagebaroption) | International | Both | None | Sets the language bar type and mode for the current user account. | Out of scope |  |
| [`Set-WinSystemLocale`](#set-winsystemlocale) | International | Both | None | Sets the system locale for the current computer. | Out of scope |  |
| [`Set-WinUILanguageOverride`](#set-winuilanguageoverride) | International | Both | None | Sets the Windows UI language override setting for the current user account. | Out of scope |  |
| [`Set-WinUserLanguageList`](#set-winuserlanguagelist) | International | Both | None | Sets the language list and associated properties for the current user account. | Out of scope |  |
| [`Set-WmiInstance`](#set-wmiinstance) | Microsoft.PowerShell.Management | Both | None | Creates or updates an instance of an existing Windows Management Instrumentation (WMI) class. | Out of scope |  |
| [`Set-WSManInstance`](#set-wsmaninstance) | Microsoft.WSMan.Management | Both | None | Modifies the management information that is related to a resource. | Out of scope |  |
| [`Set-WSManQuickConfig`](#set-wsmanquickconfig) | Microsoft.WSMan.Management | Both | None | Configures the local computer for remote management. | Out of scope |  |
| [`Show-Command`](#show-command) | Microsoft.PowerShell.Utility | Both | None | Displays PowerShell command information in a graphical window. | Out of scope | GUI and printing |
| [`Show-ControlPanelItem`](#show-controlpanelitem) | Microsoft.PowerShell.Management | Both | None | Opens control panel items. | Out of scope |  |
| [`Show-EventLog`](#show-eventlog) | Microsoft.PowerShell.Management | Both | None | Displays the event logs of the local or a remote computer in Event Viewer. | Out of scope |  |
| [`Show-WindowsDeveloperLicenseRegistration`](#show-windowsdeveloperlicenseregistration) | WindowsDeveloperLicense | Both | None | Provides information about how to enable a device for development. | Out of scope |  |
| [`Split-WindowsImage`](#split-windowsimage) | Dism | Both | Syntax differs | Splits an existing .wim file into multiple read-only split .wim files. | Out of scope |  |
| [`Start-BitsTransfer`](#start-bitstransfer) | BitsTransfer | Both | None | Creates a BITS transfer job. | Out of scope |  |
| [`Start-DscConfiguration`](#start-dscconfiguration) | PSDesiredStateConfiguration | Both | None | Applies configuration to nodes. | Out of scope |  |
| [`Start-DtcDiagnosticResourceManager`](#start-dtcdiagnosticresourcemanager) | MsDtc | Both | None | Starts a diagnostic Resource Manager. | Out of scope |  |
| [`Start-OSUninstall`](#start-osuninstall) | Dism | Both | None | Windows gives a user the ability to uninstall and roll back to a previous version of Windows. You can use DISM to initiate an uninstall. | Out of scope |  |
| [`Start-ReFSDedupJob`](#start-refsdedupjob) | Microsoft.ReFsDedup.Commands | Both | None | Starts a deduplication job on the specified ReFS volume. | Out of scope |  |
| [`Start-Service`](#start-service) | Microsoft.PowerShell.Management | Both | None | Starts one or more stopped services. | Mapped Linux (systemctl start/stop/restart) |  |
| [`Start-Transaction`](#start-transaction) | Microsoft.PowerShell.Management | Both | None | Starts a transaction. | Out of scope |  |
| [`Stop-AppvClientConnectionGroup`](#stop-appvclientconnectiongroup) | AppvClient | 5.1 only | 5.1 only | Shuts down the shared virtual environment of a connection group. | Out of scope |  |
| [`Stop-AppvClientPackage`](#stop-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Shuts down virtual environments for specified packages. | Out of scope |  |
| [`Stop-DtcDiagnosticResourceManager`](#stop-dtcdiagnosticresourcemanager) | MsDtc | Both | None | Stops and removes a diagnostic Resource Manager job. | Out of scope |  |
| [`Stop-ReFSDedupJob`](#stop-refsdedupjob) | Microsoft.ReFsDedup.Commands | Both | None | Stops a running deduplication job on a specified ReFS volume. | Out of scope |  |
| [`Stop-Service`](#stop-service) | Microsoft.PowerShell.Management | Both | None | Stops one or more running services. | Mapped Linux (systemctl start/stop/restart) |  |
| [`Suspend-BitsTransfer`](#suspend-bitstransfer) | BitsTransfer | Both | None | Suspends a BITS transfer job. | Out of scope |  |
| [`Suspend-Job`](#suspend-job) | Microsoft.PowerShell.Core | 5.1 only | 5.1 only | Temporarily stops workflow jobs. | Out of scope |  |
| [`Suspend-ReFSDedupSchedule`](#suspend-refsdedupschedule) | Microsoft.ReFsDedup.Commands | Both | None | Suspends the deduplication schedule on a specified ReFS volume. | Out of scope |  |
| [`Suspend-Service`](#suspend-service) | Microsoft.PowerShell.Management | Both | None | Suspends (pauses) one or more running services. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |
| [`Switch-Certificate`](#switch-certificate) | PKI | Both | None | Marks one certificate as having been replaced by another certificate. | Out of scope |  |
| [`Sync-AppvPublishingServer`](#sync-appvpublishingserver) | AppvClient | 5.1 only | 5.1 only | Initiates the App-V Publishing Refresh operation. | Out of scope |  |
| [`Test-AppLockerPolicy`](#test-applockerpolicy) | AppLocker | Both | None | Specifies the AppLocker policy to determine whether the input files will be allowed to run for a given user. | Out of scope |  |
| [`Test-Certificate`](#test-certificate) | PKI | Both | None | Verifies a certificate according to the input parameters. | Out of scope |  |
| [`Test-ComputerSecureChannel`](#test-computersecurechannel) | Microsoft.PowerShell.Management | Both | None | Tests and repairs the secure channel between the local computer and its domain. | Out of scope |  |
| [`Test-DscConfiguration`](#test-dscconfiguration) | PSDesiredStateConfiguration | Both | None | Tests whether the actual configuration on the nodes matches the desired configuration. | Out of scope |  |
| [`Test-FileCatalog`](#test-filecatalog) | Microsoft.PowerShell.Security | Both | None | `Test-FileCatalog` validates whether the hashes contained in a catalog file (.cat) matches the hashes of the actual files in order to validate their authenticity. | Out of scope |  |
| [`Test-KdsRootKey`](#test-kdsrootkey) | Kds | Both | None | Tests the root key configuration. | Out of scope |  |
| [`Test-PSSessionConfigurationFile`](#test-pssessionconfigurationfile) | Microsoft.PowerShell.Core | Both | None | Verifies the keys and values in a session configuration file. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Test-UevTemplate`](#test-uevtemplate) | UEV | 5.1 only | 5.1 only | Verifies whether a template complies with the schema for UE-V. | Out of scope |  |
| [`Test-WSMan`](#test-wsman) | Microsoft.WSMan.Management | Both | None | Tests whether the WinRM service is running on a local or remote computer. | Out of scope |  |
| [`Unblock-Tpm`](#unblock-tpm) | TrustedPlatformModule | Both | None | Resets a TPM lockout. | Out of scope |  |
| [`Undo-DtcDiagnosticTransaction`](#undo-dtcdiagnostictransaction) | MsDtc | Both | None | Invokes the Abort process on the specified transaction. | Out of scope |  |
| [`Undo-Transaction`](#undo-transaction) | Microsoft.PowerShell.Management | Both | None | Rolls back the active transaction. | Out of scope |  |
| [`Uninstall-Language`](#uninstall-language) | LanguagePackManagement | Both | None | Uninstalls a language from a device. | Out of scope |  |
| [`Uninstall-ProvisioningPackage`](#uninstall-provisioningpackage) | Provisioning | Both | None | Uninstalls .PPKG package from the local machine. | Out of scope |  |
| [`Uninstall-TrustedProvisioningCertificate`](#uninstall-trustedprovisioningcertificate) | Provisioning | Both | None | Removes a previously installed provisioning certificate. | Out of scope |  |
| [`Unpublish-AppvClientPackage`](#unpublish-appvclientpackage) | AppvClient | 5.1 only | 5.1 only | Removes the extension points for packages. | Out of scope |  |
| [`Unregister-PSSessionConfiguration`](#unregister-pssessionconfiguration) | Microsoft.PowerShell.Core | Both | None | Deletes registered session configurations from the computer. | Out of scope | Remote sessions (requires a remote host; out of scope) |
| [`Unregister-RecoveryManagementPlugin`](#unregister-recoverymanagementplugin) | Dism | Both | None | - | Out of scope | No documentation accurately describes this command. |
| [`Unregister-ScheduledJob`](#unregister-scheduledjob) | PSScheduledJob | Both | None | Deletes scheduled jobs on the local computer. | Out of scope |  |
| [`Unregister-UevTemplate`](#unregister-uevtemplate) | UEV | 5.1 only | 5.1 only | Unregisters a settings location template from Microsoft User Experience Virtualization (UE-V). | Out of scope |  |
| [`Unregister-WindowsDeveloperLicense`](#unregister-windowsdeveloperlicense) | WindowsDeveloperLicense | Both | None | Disables Developer Mode on the current computer. | Out of scope |  |
| [`Update-LapsADSchema`](#update-lapsadschema) | LAPS | Both | None | Extends the Active Directory (AD) schema with the Windows Local Administrator Password Solution (LAPS) schema attributes. | Out of scope |  |
| [`Update-UevTemplate`](#update-uevtemplate) | UEV | 5.1 only | 5.1 only | Updates settings location templates in UE-V. | Out of scope |  |
| [`Update-WIMBootEntry`](#update-wimbootentry) | Dism | Both | None | Updates the Windows image file boot (WIMBoot) configuration entry, associated with either the specified data source ID, the renamed image file path or the moved image file path. | Out of scope |  |
| [`Use-Transaction`](#use-transaction) | Microsoft.PowerShell.Management | Both | None | Adds the script block to the active transaction. | Out of scope |  |
| [`Use-WindowsUnattend`](#use-windowsunattend) | Dism | Both | None | Applies an unattended answer file to a Windows image. | Out of scope |  |
| [`Write-EventLog`](#write-eventlog) | Microsoft.PowerShell.Management | Both | None | Writes an event to an event log. | Out of scope | Windows-only (registry / services / recycle bin / hotfixes / event log) |

## Command Details

### Add-AppProvisionedSharedPackageContainer

Version: Both

Module: Dism

Syntax:

```powershell
Add-AppProvisionedSharedPackageContainer -DefinitionFile <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-AppProvisionedSharedPackageContainer -DefinitionFile <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Add-AppSharedPackageContainer

Version: Both

Module: Appx

Syntax:

```powershell
Add-AppSharedPackageContainer [-Path] <string> [-ForceApplicationShutdown] [-Merge] [-RequirePackagesPresent] [-Force] [<CommonParameters>]
```

Example: 

```powershell
Add-AppSharedPackageContainer -Path C:\MyFolder\ContosoTestContainer.xml
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Add-AppSharedPackageContainer.md)


### Add-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Add-AppvClientConnectionGroup [-Path] <string> [<CommonParameters>]
```

Example: Add a connection group

```powershell
PS C:\> Add-AppvClientConnectionGroup -Path "C:\MyApps\MyGroup.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Add-AppvClientConnectionGroup.md)


### Add-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Add-AppvClientPackage [-Path] <string> [[-DynamicDeploymentConfiguration] <string>] [<CommonParameters>]
```

Example: Add a package to the client

```powershell
PS C:\> Add-AppvClientPackage -Path "http://MyServer/content/package.APPV"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Add-AppvClientPackage.md)


### Add-AppvPublishingServer

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Add-AppvPublishingServer [-Name] <string> [-URL] <string> [[-GlobalRefreshEnabled] <bool>] [[-GlobalRefreshOnLogon] <bool>] [[-GlobalRefreshInterval] <uint32>] [[-GlobalRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [[-UserRefreshEnabled] <bool>] [[-UserRefreshOnLogon] <bool>] [[-UserRefreshInterval] <uint32>] [[-UserRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Add-AppvPublishingServer.md)


### Add-AppxPackage

Version: Both

Module: Appx

Syntax:

```powershell
Add-AppxPackage [-Path] <string> [-DependencyPath <string[]>] [-RequiredContentGroupOnly] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-RetainFilesOnFailure] [-InstallAllResources] [-Volume <AppxVolume>] [-ExternalPackages <string[]>] [-OptionalPackages <string[]>] [-RelatedPackages <string[]>] [-ExternalLocation <string>] [-DeferRegistrationWhenPackagesAreInUse] [-StubPackageOption <StubPackageOption>] [-AllowUnsigned] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -AppInstallerFile [-RequiredContentGroupOnly] [-ForceTargetApplicationShutdown] [-InstallAllResources] [-LimitToExistingPackages] [-Volume <AppxVolume>] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -Register [-DependencyPath <string[]>] [-DisableDevelopmentMode] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-InstallAllResources] [-ExternalLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -Update [-DependencyPath <string[]>] [-RequiredContentGroupOnly] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-RetainFilesOnFailure] [-InstallAllResources] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -Stage [-DependencyPath <string[]>] [-RequiredContentGroupOnly] [-ForceUpdateFromAnyVersion] [-Volume <AppxVolume>] [-ExternalPackages <string[]>] [-OptionalPackages <string[]>] [-RelatedPackages <string[]>] [-ExternalLocation <string>] [-StubPackageOption <StubPackageOption>] [-AllowUnsigned] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage -MainPackage <string> [-Register] [-DependencyPackages <string[]>] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-InstallAllResources] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage -RegisterByFamilyName -MainPackage <string> [-DependencyPackages <string[]>] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-InstallAllResources] [-OptionalPackages <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Add an app package

```powershell
Add-AppxPackage -Path '.\MyApp.msix' -DependencyPath '.\winjs.msix'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Add-AppxPackage.md)


### Add-AppxProvisionedPackage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Add-AppxProvisionedPackage -Path <string> [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint32>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-AppxProvisionedPackage -Online [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint32>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Add-AppxProvisionedPackage -Path <string> [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-AppxProvisionedPackage -Online [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Add an app package to the running operating system

```powershell
PS C:\> Add-AppxProvisionedPackage -Online -FolderPath "c:\Appx"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Add-AppxProvisionedPackage.md)


### Add-AppxVolume

Version: Both

Module: Appx

Syntax:

```powershell
Add-AppxVolume [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Add a volume

```powershell
Add-AppxVolume -Path "E:\WindowsApps"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Add-AppxVolume.md)


### Add-BitsFile

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Add-BitsFile [-BitsJob] <BitsJob[]> [-Source] <string[]> [[-Destination] <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Append a file to the transfer queue of an existing BITS transfer job

```powershell
PS C:\> Get-BitsTransfer -JobId 10778CFA-C1D7-4A82-8A9D-80B19224879C | Add-BitsFile -Source http://server01/servertestdir/testfile1.txt -Destination "c:\clienttestdir\testfile1.txt"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Add-BitsFile.md)


### Add-CertificateEnrollmentPolicyServer

Version: Both

Module: PKI

Syntax:

```powershell
Add-CertificateEnrollmentPolicyServer -Url <uri> -context <Context> [-NoClobber] [-RequireStrongValidation] [-Credential <PkiCredential>] [-AutoEnrollmentEnabled] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Add-CertificateEnrollmentPolicyServer -Url $url -Context Machine
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Add-CertificateEnrollmentPolicyServer.md)


### Add-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Add-Computer [-DomainName] <string> -Credential <pscredential> [-ComputerName <string[]>] [-LocalCredential <pscredential>] [-UnjoinDomainCredential <pscredential>] [-OUPath <string>] [-Server <string>] [-Unsecure] [-Options <JoinOptions>] [-Restart] [-PassThru] [-NewName <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-Computer [-WorkgroupName] <string> [-ComputerName <string[]>] [-LocalCredential <pscredential>] [-Credential <pscredential>] [-Restart] [-PassThru] [-NewName <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Add a local computer to a domain then restart the computer

```powershell
Add-Computer -DomainName Domain01 -Restart
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Add-Computer.md)


### Add-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Add-JobTrigger [-InputObject] <ScheduledJobDefinition[]> [-Trigger] <ScheduledJobTrigger[]> [<CommonParameters>]
Add-JobTrigger [-Id] <int[]> [-Trigger] <ScheduledJobTrigger[]> [<CommonParameters>]
Add-JobTrigger [-Name] <string[]> [-Trigger] <ScheduledJobTrigger[]> [<CommonParameters>]
```

Example (5.1): Add a job trigger to a scheduled job

```powershell
$Daily = New-JobTrigger -Daily -At 3AMPS
Add-JobTrigger -Trigger $Daily -Name "TestJob"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Add-JobTrigger.md)


### Add-KdsRootKey

Version: Both

Module: Kds

Syntax:

```powershell
Add-KdsRootKey [[-EffectiveTime] <datetime>] [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-KdsRootKey -EffectiveImmediately [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Generate a new root key

```powershell
PS C:\> Add-KdsRootKey
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/kds/Add-KdsRootKey.md)


### Add-LocalGroupMember

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Add-LocalGroupMember [-Group] <LocalGroup> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Add-LocalGroupMember [-Name] <string> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Add-LocalGroupMember [-SID] <SecurityIdentifier> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Add members to the Administrators group

```powershell
Add-LocalGroupMember -Group "Administrators" -Member "Admin02", "MicrosoftAccount\username@Outlook.com", "AzureAD\DavidChew@contoso.com", "CONTOSO\Domain Admins"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Add-LocalGroupMember.md)


### Add-PSSnapin

Version: 5.1 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Add-PSSnapin [-Name] <string[]> [-PassThru] [<CommonParameters>]
```

Example: Add snap-ins

```powershell
PS C:\> Add-PSSnapin -Name Microsoft.Exchange, Microsoft.Windows.AD
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Add-PSSnapin.md)


### Add-SignerRule

Version: Both

Module: ConfigCI

Syntax:

```powershell
Add-SignerRule -FilePath <string> -CertificatePath <string> [-Kernel] [-User] [-Update] [-Supplemental] [-Deny] [<CommonParameters>]
Add-SignerRule -FilePath <string> -CertStorePath <string> [-Kernel] [-User] [-Update] [-Supplemental] [-Deny] [<CommonParameters>]
```

Example: Create and add a signer rule for User mode

```powershell
PS C:\> Add-SignerRule -FilePath '.\Policy.xml' -CertificatePath '.\certificate07.cer' -User
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Add-SignerRule.md)


### Add-WindowsCapability

Version: Both

Module: Dism

Syntax:

```powershell
Add-WindowsCapability -Name <string> -Online [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-WindowsCapability -Name <string> -Path <string> [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Add a Windows capability package to the running OS via the Windows Update client

```powershell
PS C:\> Add-WindowsCapability -Online -Name "Msix.PackagingTool.Driver~~~~0.0.1.0"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Add-WindowsCapability.md)


### Add-WindowsDriver

Version: Both

Module: Dism

Syntax:

```powershell
Add-WindowsDriver -Path <string> [-Recurse] [-ForceUnsigned] [-Driver <string>] [-BasicDriverObject <BasicDriverObject>] [-AdvancedDriverObject <AdvancedDriverObject>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Add drivers to an image

```powershell
PS C:\> Add-WindowsDriver -Path "c:\offline" -Driver "c:\test\drivers" -Recurse
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Add-WindowsDriver.md)


### Add-WindowsImage

Version: Both

Module: Dism

Syntax:

```powershell
Add-WindowsImage -ImagePath <string> -CapturePath <string> [-ConfigFilePath <string>] [-Description <string>] [-Name <string>] [-CheckIntegrity] [-NoRpFix] [-Setbootable] [-Verify] [-WIMBoot] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Add files to an image

```powershell
PS C:\> Add-WindowsImage -ImagePath "C:\imagestore\custom.wim" -CapturePath d:\ -Name "Drive D"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Add-WindowsImage.md)


### Add-WindowsPackage

Version: Both

Module: Dism

Syntax:

```powershell
Add-WindowsPackage -PackagePath <string> -Online [-IgnoreCheck] [-PreventPending] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-WindowsPackage -PackagePath <string> -Path <string> [-IgnoreCheck] [-PreventPending] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Add a package to an online image

```powershell
PS C:\> Add-WindowsPackage -Online -PackagePath "c:\packages\package.cab"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Add-WindowsPackage.md)


### Checkpoint-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Checkpoint-Computer [-Description] <string> [[-RestorePointType] <string>] [<CommonParameters>]
```

Example (5.1): Create a system restore point

```powershell
Checkpoint-Computer -Description "Install MyApp"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Checkpoint-Computer.md)


### Clear-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Clear-EventLog [-LogName] <string[]> [[-ComputerName] <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Clear specific event log types from the local computer

```powershell
Clear-EventLog "Windows PowerShell"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Clear-EventLog.md)


### Clear-KdsCache

Version: Both

Module: Kds

Syntax:

```powershell
Clear-KdsCache [-CacheOwnerSid <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Clear the group key cache

```powershell
PS C:\> Clear-KdsCache
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/kds/Clear-KdsCache.md)


### Clear-RecycleBin

Version: 7 only

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Clear-RecycleBin [[-DriveLetter] <string[]>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Clear all recycle bins

```powershell
Clear-RecycleBin
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Clear-RecycleBin.md)


### Clear-Recyclebin

Version: 5.1 only

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Clear-RecycleBin [[-DriveLetter] <string[]>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Clear all recycle bins

```powershell
Clear-RecycleBin
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Clear-RecycleBin.md)


### Clear-ReFSDedupSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Clear-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Clear-ReFSDedupSchedule -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Clear-ReFSDedupSchedule.md)


### Clear-ReFSDedupScrubSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Clear-ReFSDedupScrubSchedule [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Clear-ReFSDedupScrubSchedule -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Clear-ReFSDedupScrubSchedule.md)


### Clear-Tpm

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Clear-Tpm [[-OwnerAuthorization] <string>] [-UsePPI] [<CommonParameters>]
Clear-Tpm -File <string> [<CommonParameters>]
```

Example: Reset TPM

```powershell
Clear-Tpm
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Clear-Tpm.md)


### Clear-UevAppxPackage

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Clear-UevAppxPackage [-PackageFamilyName] <string[]> [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevAppxPackage [-PackageFamilyName] <string[]> -Computer [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevAppxPackage -Computer -All [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevAppxPackage -All [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove Windows 8 apps

```powershell
PS C:\>Clear-UevAppxPackage -Computer -All
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Clear-UevAppxPackage.md)


### Clear-UevConfiguration

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Clear-UevConfiguration [-CurrentComputerUser] [-MaxPackageSizeInBytes] [-SettingsStoragePath] [-SyncProviderPingEnabled] [-SyncTimeoutInMilliseconds] [-SyncMethod] [-SyncEnabled] [-SyncOverMeteredNetwork] [-SyncOverMeteredNetworkWhenRoaming] [-SettingsImportNotifyEnabled] [-SettingsImportNotifyDelayInSeconds] [-DontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds] [-WaitForSyncOnApplicationStart] [-WaitForSyncOnLogon] [-SyncUnlistedWindows8Apps] [-VdiCollectionName] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevConfiguration [-Computer] [-MaxPackageSizeInBytes] [-SettingsStoragePath] [-SettingsTemplateCatalogPath] [-SyncProviderPingEnabled] [-SyncTimeoutInMilliseconds] [-SyncMethod] [-SyncEnabled] [-SyncOverMeteredNetwork] [-SyncOverMeteredNetworkWhenRoaming] [-SettingsImportNotifyEnabled] [-SettingsImportNotifyDelayInSeconds] [-ContactITUrl] [-ContactITDescription] [-TrayIconEnabled] [-FirstUseNotificationEnabled] [-DontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds] [-WaitForSyncOnApplicationStart] [-WaitForSyncOnLogon] [-SyncUnlistedWindows8Apps] [-VdiCollectionName] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Clear the setting for maximum package size for all users

```powershell
PS C:\> Clear-UevConfiguration -Computer -MaxPackageSizeInBytes
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Clear-UevConfiguration.md)


### Clear-WindowsCorruptMountPoint

Version: Both

Module: Dism

Syntax:

```powershell
Clear-WindowsCorruptMountPoint [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: deletes all resources associated with a mounted image

```powershell
PS C:\> Clear-WindowsCorruptMountPoint
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Clear-WindowsCorruptMountPoint.md)


### Complete-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Complete-BitsTransfer [-BitsJob] <BitsJob[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Complete all BITS transfer jobs owned by the current user

```powershell
C:\PS>Get-BitsTransfer | Complete-BitsTransfer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Complete-BitsTransfer.md)


### Complete-DtcDiagnosticTransaction

Version: Both

Module: MsDtc

Syntax:

```powershell
Complete-DtcDiagnosticTransaction [-Transaction] <DtcDiagnosticTransaction> [<CommonParameters>]
```

Example: Complete a DTC diagnostic transaction

```powershell
PS C:\> $Tx = New-DtcDiagnosticTransaction
PS C:\> Complete-DtcDiagnosticTransaction -Transaction $Tx
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Complete-DtcDiagnosticTransaction.md)


### Complete-Transaction

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Complete-Transaction [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Commit a transaction

```powershell
Set-Location HKCU:\software
Start-Transaction
New-Item MyCompany -UseTransaction
Get-ChildItem m*
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Complete-Transaction.md)


### Confirm-SecureBootUEFI

Version: Both

Module: SecureBoot

Syntax:

```powershell
Confirm-SecureBootUEFI [<CommonParameters>]
```

Example: Confirm Secure Boot

```powershell
PS C:\> Confirm-SecureBootUEFI
True
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/SecureBoot/Confirm-SecureBootUEFI.md)


### Connect-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Connect-PSSession -Name <string[]> [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession [-Session] <PSSession[]> [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession [-ComputerName] <string[]> [-ApplicationName <string>] [-ConfigurationName <string>] [-Name <string[]>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-SessionOption <PSSessionOption>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession -ComputerName <string[]> -InstanceId <guid[]> [-ApplicationName <string>] [-ConfigurationName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-SessionOption <PSSessionOption>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession [-ConnectionUri] <uri[]> [-ConfigurationName <string>] [-AllowRedirection] [-Name <string[]>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-SessionOption <PSSessionOption>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession [-ConnectionUri] <uri[]> -InstanceId <guid[]> [-ConfigurationName <string>] [-AllowRedirection] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-SessionOption <PSSessionOption>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession -InstanceId <guid[]> [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Connect-PSSession [-Id] <int[]> [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Reconnect to a session

```powershell
Connect-PSSession -ComputerName Server01 -Name ITTask
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Connect-PSSession.md)


### Connect-WSMan

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Connect-WSMan [[-ComputerName] <string>] [-ApplicationName <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Connect-WSMan [-ConnectionURI <uri>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Connect to a remote computer

```powershell
PS C:\> Connect-WSMan -ComputerName "server01"
PS C:\> cd WSMan:
PS WSMan:\>
PS WSMan:\> dir
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Connect-WSMan.md)


### Convert-String

Version: 5.1 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Convert-String -InputObject <string> [-Example <List[psobject]>] [<CommonParameters>]
```

Example: Convert format of a string

```powershell
"Mu Han", "Jim Hance", "David Ahs", "Kim Akers" | Convert-String -Example "Ed Wilson=Wilson, E."
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/Convert-String.md)


### ConvertFrom-CIPolicy

Version: Both

Module: ConfigCI

Syntax:

```powershell
ConvertFrom-CIPolicy [-XmlFilePath] <string> [-BinaryFilePath] <string> [<CommonParameters>]
```

Example: Converts a policy

```powershell
PS C:\> ConvertFrom-CIPolicy -XmlFilePath ".\Policy03.xml" -BinaryFilePath "Policy03.bin"
C:\Policies\Policy03.bin
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/ConvertFrom-CIPolicy.md)


### ConvertFrom-SddlString

Version: 7 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertFrom-SddlString [-Sddl] <string> [-Type <ConvertFromSddlStringCommand+AccessRightTypeNames>] [<CommonParameters>]
```

Example: Convert file system access rights SDDL to a PSCustomObject

```powershell
$acl = Get-Acl -Path C:\Windows
ConvertFrom-SddlString -Sddl $acl.Sddl
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/ConvertFrom-SddlString.md)


### ConvertFrom-String

Version: 5.1 only

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
ConvertFrom-String [-InputObject] <string> [-Delimiter <string>] [-PropertyNames <string[]>] [<CommonParameters>]
ConvertFrom-String [-InputObject] <string> [-TemplateFile <string[]>] [-TemplateContent <string[]>] [-IncludeExtent] [-UpdateTemplate] [<CommonParameters>]
```

Example: Generate an object with default property names

```powershell
"Hello World" | ConvertFrom-String
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Utility/ConvertFrom-String.md)


### ConvertTo-ProcessMitigationPolicy

Version: Both

Module: ProcessMitigations

Syntax:

```powershell
ConvertTo-ProcessMitigationPolicy [-EMETFilePath] <string> [-OutputFilePath] <string> [<CommonParameters>]
```

Example: 

```powershell
PS C:\> ConvertTo-ProcessMitigationPolicy -EMETFilePath policy.xml -OutputFilePath result.xml
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ProcessMitigations/ConvertTo-ProcessMitigationPolicy.md)


### ConvertTo-TpmOwnerAuth

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
ConvertTo-TpmOwnerAuth [-PassPhrase] <string> [<CommonParameters>]
```

Example: Convert to owner authorization value

```powershell
PS C:\> ConvertTo-TpmOwnerAuth -PassPhrase "Saturn1977&&"
puJvGK4O6Qvl0loP8r1bIxipDVo=
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/ConvertTo-TpmOwnerAuth.md)


### Copy-BcdEntry

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Copy-BcdEntry [-SourceEntryId] <string> -TargetStore <BcdStoreInfo[]> [-Description <string>] [-SourceStore <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Copy-BcdEntry [-SourceEntry] <BcdEntryInfo> -TargetStore <BcdStoreInfo[]> [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Copy-BcdEntry.md)


### Copy-UserInternationalSettingsToSystem

Version: Both

Module: International

Syntax:

```powershell
Copy-UserInternationalSettingsToSystem [-NewUser] <bool> [<CommonParameters>]
```

Example: Copy settings into both the Welcome screen and system accounts, and new user accounts

```powershell
PS C:\> Copy-UserInternationalSettingsToSystem -WelcomeScreen $True -NewUser $True
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Copy-UserInternationalSettingsToSystem.md)


### Disable-AppBackgroundTaskDiagnosticLog

Version: Both

Module: AppBackgroundTask

Syntax:

```powershell
Disable-AppBackgroundTaskDiagnosticLog [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable background task logging

```powershell
PS C:\> Disable-AppBackgroundTaskDiagnosticLog
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppBackgroundTask/Disable-AppBackgroundTaskDiagnosticLog.md)


### Disable-Appv

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Disable-Appv [<CommonParameters>]
```

Example: Disable the App-V service

```powershell
PS C:\> Disable-Appv
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Disable-Appv.md)


### Disable-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Disable-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [-UserSID <string>] [<CommonParameters>]
Disable-AppvClientConnectionGroup [-Name] <string> [-Global] [-UserSID <string>] [<CommonParameters>]
Disable-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [-UserSID <string>] [<CommonParameters>]
```

Example: Disable a connection group by using its name

```powershell
PS C:\> Disable-AppvClientConnectionGroup -Name "MyGroup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Disable-AppvClientConnectionGroup.md)


### Disable-BcdElementBootDebug

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Disable-BcdElementBootDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementBootDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Disable-BcdElementBootDebug.md)


### Disable-BcdElementBootEms

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Disable-BcdElementBootEms [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementBootEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementBootEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Disable-BcdElementBootEms.md)


### Disable-BcdElementDebug

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Disable-BcdElementDebug [[-Id] <string>] [-Store <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Disable-BcdElementDebug.md)


### Disable-BcdElementEms

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Disable-BcdElementEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Disable-BcdElementEms.md)


### Disable-BcdElementEventLogging

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Disable-BcdElementEventLogging [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementEventLogging [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementEventLogging [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Disable-BcdElementEventLogging.md)


### Disable-BcdElementHypervisorDebug

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Disable-BcdElementHypervisorDebug [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementHypervisorDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementHypervisorDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Disable-BcdElementHypervisorDebug.md)


### Disable-ComputerRestore

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Disable-ComputerRestore [-Drive] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Disable System Restore on the specified drive

```powershell
Disable-ComputerRestore -Drive "C:\"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Disable-ComputerRestore.md)


### Disable-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Disable-JobTrigger [-InputObject] <ScheduledJobTrigger[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Disable a job trigger

```powershell
PS C:\> Get-JobTrigger -Name "Backup-Archives" -TriggerId 1 | Disable-JobTrigger
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Disable-JobTrigger.md)


### Disable-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Disable-LocalUser [-InputObject] <LocalUser[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-LocalUser [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-LocalUser [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Disable an account by specifying a name

```powershell
Disable-LocalUser -Name "Admin02"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Disable-LocalUser.md)


### Disable-PSRemoting

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Disable-PSRemoting [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Prevent remote access to all session configurations

```powershell
Disable-PSRemoting
```

Example (7): Prevent remote access to all PowerShell session configurations

```powershell
Disable-PSRemoting
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Disable-PSRemoting.md)


### Disable-PSSessionConfiguration

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Disable-PSSessionConfiguration [[-Name] <string[]>] [-Force] [-NoServiceRestart] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable the default configuration

```powershell
Disable-PSSessionConfiguration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Disable-PSSessionConfiguration.md)


### Disable-ReFSDedup

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Disable-ReFSDedup [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Disable-ReFSDedup -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Disable-ReFSDedup.md)


### Disable-ScheduledJob

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Disable-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-ScheduledJob [-Id] <int> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-ScheduledJob [-Name] <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Disable a scheduled job

```powershell
Disable-ScheduledJob -Id 2 -PassThru
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Disable-ScheduledJob.md)


### Disable-TlsCipherSuite

Version: Both

Module: TLS

Syntax:

```powershell
Disable-TlsCipherSuite [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable a cipher suite

```powershell
Disable-TlsCipherSuite -Name 'TLS_RSA_WITH_3DES_EDE_CBC_SHA'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Disable-TlsCipherSuite.md)


### Disable-TlsEccCurve

Version: Both

Module: TLS

Syntax:

```powershell
Disable-TlsEccCurve [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Disable-TlsEccCurve -Name curve25519
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Disable-TlsEccCurve.md)


### Disable-TlsSessionTicketKey

Version: Both

Module: TLS

Syntax:

```powershell
Disable-TlsSessionTicketKey [-ServiceAccountName] <NTAccount> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable a TLS session ticket key

```powershell
Disable-TlsSessionTicketKey -ServiceAccountName NetworkService
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Disable-TlsSessionTicketKey.md)


### Disable-TpmAutoProvisioning

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Disable-TpmAutoProvisioning [-OnlyForNextRestart] [<CommonParameters>]
```

Example: Disable auto-provisioning

```powershell
PS C:\> Disable-TpmAutoProvisioning
TpmReady           : False
TpmPresent         : True
ManagedAuthLevel   : Full
OwnerAuth          : OwnerClearDisabled : True
AutoProvisioning   : Disabled
LockedOut          : False
SelfTest           : {191, 191, 245, 191...}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Disable-TpmAutoProvisioning.md)


### Disable-Uev

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Disable-Uev [<CommonParameters>]
```

Example: Disable the UE-V service

```powershell
PS C:\>Disable-Uev
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Disable-Uev.md)


### Disable-UevAppxPackage

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Disable-UevAppxPackage [-PackageFamilyName] <string[]> [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-UevAppxPackage [-PackageFamilyName] <string[]> -Computer [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable synchronization of a Windows 8 app

```powershell
PS C:\>Disable-UevAppxPackage -Computer -PackageFamilyName "Microsoft.BingFinance"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Disable-UevAppxPackage.md)


### Disable-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Disable-UevTemplate [-ID] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable a specific template

```powershell
PS C:\> Disable-UevTemplate -ID "MicrosoftCalculator6"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Disable-UevTemplate.md)


### Disable-WindowsErrorReporting

Version: Both

Module: WindowsErrorReporting

Syntax:

```powershell
Disable-WindowsErrorReporting [<CommonParameters>]
```

Example: Disable Windows Error Reporting

```powershell
PS C:\> Disable-WindowsErrorReporting
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsErrorReporting/Disable-WindowsErrorReporting.md)


### Disable-WindowsOptionalFeature

Version: Both

Module: Dism

Syntax:

```powershell
Disable-WindowsOptionalFeature -FeatureName <string[]> -Online [-PackageName <string>] [-Remove] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Disable-WindowsOptionalFeature -FeatureName <string[]> -Path <string> [-PackageName <string>] [-Remove] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Disable an optional feature

```powershell
PS C:\> Disable-WindowsOptionalFeature -Online -FeatureName "Hearts"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Disable-WindowsOptionalFeature.md)


### Disable-WSManCredSSP

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Disable-WSManCredSSP [-Role] <string> [<CommonParameters>]
```

Example: Disable CredSSP on a client

```powershell
Disable-WSManCredSSP -Role Client
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Disable-WSManCredSSP.md)


### Disconnect-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Disconnect-PSSession [-Session] <PSSession[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -InstanceId <guid[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -Name <string[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession [-Id] <int[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Disconnect-PSSession [-Session] <PSSession[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -Name <string[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession [-Id] <int[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -InstanceId <guid[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disconnect a session by name

```powershell
PS> Disconnect-PSSession -Name UpdateSession
Id Name            ComputerName    State         ConfigurationName     Availability
-- ----            ------------    -----         -----------------     ------------
1  UpdateSession   Server01        Disconnected  Microsoft.PowerShell          None
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Disconnect-PSSession.md)


### Disconnect-WSMan

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Disconnect-WSMan [[-ComputerName] <string>] [<CommonParameters>]
```

Example: Delete a connection to a remote computer

```powershell
PS C:\> Disconnect-WSMan -Computer server01
PS C:\> cd WSMan:
PS WSMan:\> dir
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Disconnect-WSMan.md)


### Dismount-AppxVolume

Version: Both

Module: Appx

Syntax:

```powershell
Dismount-AppxVolume [-Volume] <AppxVolume[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Dismount a volume by using a path

```powershell
Dismount-AppxVolume -Volume E:\
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Dismount-AppxVolume.md)


### Dismount-WindowsImage

Version: Both

Module: Dism

Syntax:

```powershell
Dismount-WindowsImage -Path <string> -Discard [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Dismount-WindowsImage -Path <string> -Save [-CheckIntegrity] [-Append] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Dismount an operating system image

```powershell
PS C:\> Dismount-WindowsImage -Path "c:\offline" -Save
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Dismount-WindowsImage.md)


### Edit-CIPolicyRule

Version: Both

Module: ConfigCI

Syntax:

```powershell
Edit-CIPolicyRule [-Id] <string> -FilePath <string> [-Name <string>] [-RType <string>] [-FileName <string>] [-Version <string>] [-HashPath <string>] [<CommonParameters>]
Edit-CIPolicyRule [-Id] <string> -FilePath <string> [-Name <string>] [-RType <string>] [-Root <string>] [-AddEkus <string[]>] [-RemoveEkus <string[]>] [-Issuer <string>] [-Publisher <string>] [-OemId <string>] [-AddExceptions <string[]>] [-RemoveExceptions <string[]>] [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Edit-CIPolicyRule.md)


### Enable-AppBackgroundTaskDiagnosticLog

Version: Both

Module: AppBackgroundTask

Syntax:

```powershell
Enable-AppBackgroundTaskDiagnosticLog [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Enable background task logging

```powershell
PS C:\> Enable-AppBackgroundTaskDiagnosticLog
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppBackgroundTask/Enable-AppBackgroundTaskDiagnosticLog.md)


### Enable-Appv

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Enable-Appv [<CommonParameters>]
```

Example: Enable the service

```powershell
PS C:\> Enable-Appv
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Enable-Appv.md)


### Enable-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Enable-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [-UserSID <string>] [<CommonParameters>]
Enable-AppvClientConnectionGroup [-Name] <string> [-Global] [-UserSID <string>] [<CommonParameters>]
Enable-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [-UserSID <string>] [<CommonParameters>]
```

Example: Enable a connection group by using its name

```powershell
PS C:\> Enable-AppvClientConnectionGroup -Name "MyGroup" -Global
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Enable-AppvClientConnectionGroup.md)


### Enable-BcdElementBootDebug

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Enable-BcdElementBootDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementBootDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Enable-BcdElementBootDebug.md)


### Enable-BcdElementBootEms

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Enable-BcdElementBootEms [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementBootEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementBootEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Enable-BcdElementBootEms.md)


### Enable-BcdElementDebug

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Enable-BcdElementDebug [[-Id] <string>] [-Store <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Enable-BcdElementDebug.md)


### Enable-BcdElementEms

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Enable-BcdElementEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Enable-BcdElementEms.md)


### Enable-BcdElementEventLogging

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Enable-BcdElementEventLogging [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementEventLogging [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementEventLogging [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Enable-BcdElementEventLogging.md)


### Enable-BcdElementHypervisorDebug

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Enable-BcdElementHypervisorDebug [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementHypervisorDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementHypervisorDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Enable-BcdElementHypervisorDebug.md)


### Enable-ComputerRestore

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Enable-ComputerRestore [-Drive] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Enable System Restore on the specified drive

```powershell
Enable-ComputerRestore -Drive "C:\"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Enable-ComputerRestore.md)


### Enable-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Enable-JobTrigger [-InputObject] <ScheduledJobTrigger[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Enable a job trigger

```powershell
Get-JobTrigger -Name Backup-Archives -TriggerId 1 | Enable-JobTrigger
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Enable-JobTrigger.md)


### Enable-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Enable-LocalUser [-InputObject] <LocalUser[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-LocalUser [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-LocalUser [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Enable an account by specifying a name

```powershell
Enable-LocalUser -Name "Admin02"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Enable-LocalUser.md)


### Enable-PSRemoting

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Enable-PSRemoting [-Force] [-SkipNetworkProfileCheck] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Configure a computer to receive remote commands

```powershell
Enable-PSRemoting
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Enable-PSRemoting.md)


### Enable-PSSessionConfiguration

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Enable-PSSessionConfiguration [[-Name] <string[]>] [-Force] [-SecurityDescriptorSddl <string>] [-SkipNetworkProfileCheck] [-NoServiceRestart] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Re-enable the default session

```powershell
Enable-PSSessionConfiguration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Enable-PSSessionConfiguration.md)


### Enable-ReFSDedup

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Enable-ReFSDedup [-Volume] <string> [-Type] <DedupVolumeType> [<CommonParameters>]
```

Example: 

```powershell
Enable-ReFSDedup -Volume "D:" -Type DedupAndCompress
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Enable-ReFSDedup.md)


### Enable-ScheduledJob

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Enable-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-ScheduledJob [-Id] <int> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-ScheduledJob [-Name] <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Enable a scheduled job

```powershell
Enable-ScheduledJob -Id 2 -PassThru
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Enable-ScheduledJob.md)


### Enable-TlsCipherSuite

Version: Both

Module: TLS

Syntax (5.1):

```powershell
Enable-TlsCipherSuite [-Name] <string> [[-Position] <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Enable-TlsCipherSuite [-Name] <string> [[-Position] <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Enable a cipher suite

```powershell
Enable-TlsCipherSuite -Name TLS_DHE_DSS_WITH_AES_256_CBC_SHA
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Enable-TlsCipherSuite.md)


### Enable-TlsEccCurve

Version: Both

Module: TLS

Syntax (5.1):

```powershell
Enable-TlsEccCurve [-Name] <string> [[-Position] <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Enable-TlsEccCurve [-Name] <string> [[-Position] <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Enable-TlsEccCurve 'NistP384' -Position 0
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Enable-TlsEccCurve.md)


### Enable-TlsSessionTicketKey

Version: Both

Module: TLS

Syntax:

```powershell
Enable-TlsSessionTicketKey [-Password] <securestring> [-Path] <string> [-ServiceAccountName] <NTAccount> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Configure a TLS server with a TLS session ticket key for the NetworkService account

```powershell
$Password = Read-Host -AsSecureString
Enable-TlsSessionTicketKey -Password $Password -Path 'C:\KeyConfig\TlsSessionTicketKey.config' -ServiceAccountName NetworkService
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Enable-TlsSessionTicketKey.md)


### Enable-TpmAutoProvisioning

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Enable-TpmAutoProvisioning [<CommonParameters>]
```

Example: Enable auto-provisioning

```powershell
PS C:\> Enable-TpmAutoProvisioning
TpmReady           : False
TpmPresent         : True
ManagedAuthLevel   : Full
OwnerAuth          : OwnerClearDisabled : True
AutoProvisioning   : Enabled
LockedOut          : False
SelfTest           : {191, 191, 245, 191...}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Enable-TpmAutoProvisioning.md)


### Enable-Uev

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Enable-Uev [<CommonParameters>]
```

Example: Enable the UE-V service

```powershell
PS C:\>Enable-Uev
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Enable-Uev.md)


### Enable-UevAppxPackage

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Enable-UevAppxPackage [-PackageFamilyName] <string[]> [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-UevAppxPackage [-PackageFamilyName] <string[]> -Computer [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Enable synchronization of a Windows 8 app

```powershell
PS C:\>Enable-UevAppxPackage -PackageFamilyName "Microsoft.BingTravel"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Enable-UevAppxPackage.md)


### Enable-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Enable-UevTemplate [-ID] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Enable a specific template

```powershell
PS C:\> Enable-UevTemplate -ID "MicrosoftCalculator6"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Enable-UevTemplate.md)


### Enable-WindowsErrorReporting

Version: Both

Module: WindowsErrorReporting

Syntax:

```powershell
Enable-WindowsErrorReporting [<CommonParameters>]
```

Example: Enable Windows Error Reporting

```powershell
PS C:\> Enable-WindowsErrorReporting
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsErrorReporting/Enable-WindowsErrorReporting.md)


### Enable-WindowsOptionalFeature

Version: Both

Module: Dism

Syntax:

```powershell
Enable-WindowsOptionalFeature -FeatureName <string[]> -Online [-PackageName <string>] [-All] [-LimitAccess] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Enable-WindowsOptionalFeature -FeatureName <string[]> -Path <string> [-PackageName <string>] [-All] [-LimitAccess] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Enable an optional feature in the running operating system

```powershell
PS C:\> Enable-WindowsOptionalFeature -Online -FeatureName "Hearts" -All
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Enable-WindowsOptionalFeature.md)


### Enable-WSManCredSSP

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Enable-WSManCredSSP [-Role] <string> [[-DelegateComputer] <string[]>] [-Force] [<CommonParameters>]
```

Example: Delegate client credentials

```powershell
Enable-WSManCredSSP -Role "Client" -DelegateComputer "Server02.fabrikam.com"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Enable-WSManCredSSP.md)


### Expand-OsImage

Version: Both

Module: Dism

Syntax:

```powershell
Expand-OsImage -ImagePath <string> -ApplyPath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Expand-WindowsCustomDataImage

Version: Both

Module: Dism

Syntax:

```powershell
Expand-WindowsCustomDataImage -ImagePath <string> -CustomDataImage <string> -SingleInstance [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Expand a custom data image

```powershell
PS C:\> Expand-WindowsCustomDataImage -CustomDataImage "C:\oem.ppkg" -ImagePath "C:\" -SingleInstance
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Expand-WindowsCustomDataImage.md)


### Expand-WindowsImage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Expand-WindowsImage -ImagePath <string> -Name <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -Index <uint32> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Expand-WindowsImage -ImagePath <string> -Name <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -Index <uint> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Apply an image in a file to a partion

```powershell
PS C:\> Expand-WindowsImage -ImagePath "c:\imagestore\custom.wim" -ApplyPath "d:\" -Index 1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Expand-WindowsImage.md)


### Export-BcdStore

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Export-BcdStore [-Path] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Export-BcdStore.md)


### Export-BinaryMiLog

Version: Both

Module: CimCmdlets

Syntax:

```powershell
Export-BinaryMiLog [-Path] <string> [-InputObject <ciminstance>] [<CommonParameters>]
```

Example (5.1): Create a binary representation of CimInstances

```powershell
Get-CimInstance Win32_Process | Export-BinaryMiLog -Path "Processes.bmil"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/CimCmdlets/Export-BinaryMiLog.md)


### Export-Certificate

Version: Both

Module: PKI

Syntax:

```powershell
Export-Certificate -FilePath <string> -Cert <Certificate> [-Type <CertType>] [-NoClobber] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$cert = Get-ChildItem -Path Cert:\CurrentUser\My\EEDEF61D4FF6EDBAAD538BB08CCAADDC3EE28FF

Export-Certificate -Cert $cert -FilePath C:\Certs\user.sst -Type SST
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Export-Certificate.md)


### Export-Console

Version: 5.1 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Export-Console [[-Path] <string>] [-Force] [-NoClobber] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export the names of snap-ins in the current session

```powershell
PS C:\> Export-Console -Path $PSHOME\Consoles\ConsoleS1.psc1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Export-Console.md)


### Export-Counter

Version: Both

Module: Microsoft.PowerShell.Diagnostics

Syntax:

```powershell
Export-Counter [-Path] <string> -InputObject <PerformanceCounterSampleSet[]> [-FileFormat <string>] [-MaxSize <uint32>] [-Force] [-Circular] [<CommonParameters>]
```

Example (5.1): Export counter data to a file

```powershell
Get-Counter "\Processor(*)\% Processor Time" | Export-Counter -Path $HOME\Counters.blg
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Diagnostics/Export-Counter.md)


### Export-OsImage

Version: Both

Module: Dism

Syntax:

```powershell
Export-OsImage -SrcImagePath <string> -DestImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Export-PfxCertificate

Version: Both

Module: PKI

Syntax:

```powershell
Export-PfxCertificate [-PFXData] <PfxData> [-FilePath] <string> [-NoProperties] [-NoClobber] [-Force] [-CryptoAlgorithmOption <CryptoAlgorithmOptions>] [-ChainOption <ExportChainOption>] [-ProtectTo <string[]>] [-Password <securestring>] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-PfxCertificate [-Cert] <Certificate> [-FilePath] <string> [-NoProperties] [-NoClobber] [-Force] [-CryptoAlgorithmOption <CryptoAlgorithmOptions>] [-ChainOption <ExportChainOption>] [-ProtectTo <string[]>] [-Password <securestring>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$mypwd = ConvertTo-SecureString -String '1234' -Force -AsPlainText

Get-ChildItem -Path Cert:\LocalMachine\My\5F98EBBFE735CDDAE00E33E0FD69050EF9220254 |
    Export-PfxCertificate -FilePath C:\mypfx.pfx -Password $mypwd
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Export-PfxCertificate.md)


### Export-ProvisioningPackage

Version: Both

Module: Provisioning

Syntax:

```powershell
Export-ProvisioningPackage [-OutputFolder] <string> -PackageId <string> [-AllowClobber] [-AnswerFileOnly] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Export-ProvisioningPackage [-PackagePath] <string> [-OutputFolder] <string> [-AllowClobber] [-AnswerFileOnly] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Export-ProvisioningPackage [-RuntimeMetadata] <RuntimeProvPackageMetadata> [-OutputFolder] <string> [-AllowClobber] [-AnswerFileOnly] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Export-ProvisioningPackage -PackageId {e2ea11f5-d8b0-4db9-bf96-8c909dc2fed5} -OutputFolder D:\Package
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Export-ProvisioningPackage.md)


### Export-StartLayout

Version: Both

Module: StartLayout

Syntax:

```powershell
Export-StartLayout [-Path] <string> [-UseDesktopApplicationID] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-StartLayout -LiteralPath <string> [-UseDesktopApplicationID] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export the layout

```powershell
PS C:\> Export-StartLayout -Path "C:\Layouts\Marketing.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/StartLayout/Export-StartLayout.md)


### Export-StartLayoutEdgeAssets

Version: Both

Module: StartLayout

Syntax:

```powershell
Export-StartLayoutEdgeAssets [-Path] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Export-StartLayoutEdgeAssets -LiteralPath <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export assets

```powershell
Export-StartLayoutEdgeAssets -Path "C:\Layouts\assets.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/StartLayout/Export-StartLayoutEdgeAssets.md)


### Export-TlsSessionTicketKey

Version: Both

Module: TLS

Syntax:

```powershell
Export-TlsSessionTicketKey [-Password] <securestring> [[-Path] <string>] [-ServiceAccountName] <NTAccount> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export a TLS session ticket key

```powershell
$Password = Read-Host -AsSecureString
Export-TlsSessionTicketKey -Password $Password -Path 'C:\KeyConfig\TlsSessionTicketKey.config' -ServiceAccountName NetworkService
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Export-TlsSessionTicketKey.md)


### Export-Trace

Version: Both

Module: Provisioning

Syntax:

```powershell
Export-Trace [-ETLFile] <string> [-Overwrite] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: Export trace events from an ETL file

```powershell
PS C:\> Export-Trace -ETLFile C:\Windows\Logs\WindowsUpdate\WindowsUpdate.20211013.074054.819.1.etl -LogsDirectoryPath C:\ETL\Logs
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Export-Trace.md)


### Export-UevConfiguration

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Export-UevConfiguration [-Path] <string> [<CommonParameters>]
```

Example: Export the UE-V configuration

```powershell
PS C:\> Export-UevConfiguration -Path "ContosoUev.uev"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Export-UevConfiguration.md)


### Export-UevPackage

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Export-UevPackage [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Export-UevPackage -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Export a UE-V package

```powershell
PS C:\> Export-UevPackage -Path "MicrosoftCalculator6.pkgx"
<SettingsDocument>
<registry>
<Setting Type="VT_BINARY" Name="registry://HKCU\Software\Microsoft\Calc\Window_Placement" Action="Update">LAAAAAAAAAABAAAA/////////////////////60AAABQAAAAVAIAANQBAAA=</Setting>
<Setting Type="VT_DWORD" Name="registry://HKCU\Software\Microsoft\Calc\layout" Action="Update">2</Setting>
</registry>
</SettingsDocument>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Export-UevPackage.md)


### Export-WindowsCapabilitySource

Version: Both

Module: Dism

Syntax:

```powershell
Export-WindowsCapabilitySource -Name <string> -Source <string> -Target <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Export repository for capability

```powershell
Export-WindowsCapabilitySource -Path c:\mount\windows -Source D:\ -Target C:\repository -Name App.StepsRecorder~~~~0.0.1.0
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Export-WindowsCapabilitySource.md)


### Export-WindowsDriver

Version: Both

Module: Dism

Syntax:

```powershell
Export-WindowsDriver -Path <string> [-Destination <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsDriver -Online [-Destination <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Export drivers from the running operating system

```powershell
PS C:\> Export-WindowsDriver -Online -Destination d:\drivers
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Export-WindowsDriver.md)


### Export-WindowsImage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceName <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceIndex <uint32> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceName <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceIndex <uint> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Export an image

```powershell
PS C:\> Export-WindowsImage -SourceImagePath C:\imagestore\custom.wim -SourceIndex 1 -DestinationImagePath c:\imagestore\export.wim -DestinationName "Exported Image"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Export-WindowsImage.md)


### Find-LapsADExtendedRights

Version: Both

Module: LAPS

Syntax:

```powershell
Find-LapsADExtendedRights -Identity <string[]> [-Credential <pscredential>] [-Domain <string>] [-DomainController <string>] [-IncludeComputers] [<CommonParameters>]
```

Example: 

```powershell
Find-LapsADExtendedRights -Identity LapsTestOU
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Find-LapsADExtendedRights.md)


### Format-SecureBootUEFI

Version: Both

Module: SecureBoot

Syntax:

```powershell
Format-SecureBootUEFI -Name <string> -SignatureOwner <guid> -CertificateFilePath <string[]> [-FormatWithCert] [-SignableFilePath <string>] [-Time <string>] [-AppendWrite] [-ContentFilePath <string>] [<CommonParameters>]
Format-SecureBootUEFI -Name <string> -SignatureOwner <guid> -Hash <string[]> -Algorithm <string> [-SignableFilePath <string>] [-Time <string>] [-AppendWrite] [-ContentFilePath <string>] [<CommonParameters>]
Format-SecureBootUEFI -Name <string> -Delete [-SignableFilePath <string>] [-Time <string>] [<CommonParameters>]
```

Example: Format a private key

```powershell
PS C:\> Format-SecureBootUefi -Name PK -SignatureOwner 12345678-1234-1234-1234-123456789abc -CertificateFilePath PK.cer -SignableFilePath GeneratedFileToSign.bin -Time 2011-11-01T13:30:00Z | Format-List
Name        : PK
Time        : 2011-11-01T13:30:00Z
AppendWrite : False
Content     : {232, 102, 87, 60...}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/SecureBoot/Format-SecureBootUEFI.md)


### Get-Acl

Version: Both

Module: Microsoft.PowerShell.Security

Syntax (5.1):

```powershell
Get-Acl [[-Path] <string[]>] [-Audit] [-AllCentralAccessPolicies] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-UseTransaction] [<CommonParameters>]
Get-Acl -InputObject <psobject> [-Audit] [-AllCentralAccessPolicies] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-UseTransaction] [<CommonParameters>]
Get-Acl [-LiteralPath <string[]>] [-Audit] [-AllCentralAccessPolicies] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Acl [[-Path] <string[]>] [-Audit] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Acl -InputObject <psobject> [-Audit] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Acl [-LiteralPath <string[]>] [-Audit] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
```

Example: Get an ACL for a folder

```powershell
Get-Acl C:\Windows
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Get-Acl.md)


### Get-AppLockerFileInformation

Version: Both

Module: AppLocker

Syntax:

```powershell
Get-AppLockerFileInformation [[-Path] <List[string]>] [<CommonParameters>]
Get-AppLockerFileInformation [[-Packages] <List[AppxPackage]>] [<CommonParameters>]
Get-AppLockerFileInformation -Directory <string> [-FileType <List[AppLockerFileType]>] [-Recurse] [<CommonParameters>]
Get-AppLockerFileInformation -EventLog [-LogPath <string>] [-EventType <List[AppLockerEventType]>] [-Statistics] [<CommonParameters>]
```

Example: Get file information for .exe files and scripts

```powershell
PS C:\> Get-AppLockerFileInformation -Directory C:\Windows\system32\ -Recurse -FileType exe, script
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppLocker/Get-AppLockerFileInformation.md)


### Get-AppLockerPolicy

Version: Both

Module: AppLocker

Syntax:

```powershell
Get-AppLockerPolicy -Local [-Xml] [<CommonParameters>]
Get-AppLockerPolicy -Domain -Ldap <string> [-Xml] [<CommonParameters>]
Get-AppLockerPolicy -Effective [-Xml] [<CommonParameters>]
```

Example: Get an AppLocker policy

```powershell
PS C:\> Get-AppLockerPolicy -Local
                                Version RuleCollections                         RuleCollectionTypes
                                ------- ---------------                         -------------------
                                      1 {}                                      {}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppLocker/Get-AppLockerPolicy.md)


### Get-AppProvisionedSharedPackageContainer

Version: Both

Module: Dism

Syntax:

```powershell
Get-AppProvisionedSharedPackageContainer -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-AppProvisionedSharedPackageContainer -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Get-AppSharedPackageContainer

Version: Both

Module: Appx

Syntax:

```powershell
Get-AppSharedPackageContainer [[-Name] <string>] [[-Id] <string>] [<CommonParameters>]
```

Example: 

```powershell
Get-AppSharedPackageContainer -Name Contoso*
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Get-AppSharedPackageContainer.md)


### Get-AppvClientApplication

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvClientApplication [[-Name] <string>] [[-Version] <string>] [-All] [<CommonParameters>]
```

Example: Get a version of an application for the current user

```powershell
PS C:\> Get-AppvClientApplication -Name "AppName" -Version 1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvClientApplication.md)


### Get-AppvClientConfiguration

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvClientConfiguration [[-Name] <string>] [<CommonParameters>]
```

Example: Display all configuration settings

```powershell
PS C:\> Get-AppvClientConfiguration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvClientConfiguration.md)


### Get-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvClientConnectionGroup [[-Name] <string>] [-All] [<CommonParameters>]
Get-AppvClientConnectionGroup [-GroupId] <guid> [[-VersionId] <guid>] [-All] [<CommonParameters>]
```

Example: Get all versions of a group by name

```powershell
PS C:\> Get-AppvClientConnectionGroup -Name "MyConnectionGroup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvClientConnectionGroup.md)


### Get-AppvClientMode

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvClientMode [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvClientMode.md)


### Get-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvClientPackage [[-Name] <string>] [[-Version] <string>] [-All] [<CommonParameters>]
Get-AppvClientPackage [-PackageId] <guid> [[-VersionId] <guid>] [-All] [<CommonParameters>]
```

Example: Get packages that have names that match a string

```powershell
PS C:\> Get-AppvClientPackage -Name "MyApp*" -All
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvClientPackage.md)


### Get-AppvPublishingServer

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvPublishingServer [[-ServerId] <uint32>] [<CommonParameters>]
Get-AppvPublishingServer [[-Name] <string>] [[-URL] <string>] [<CommonParameters>]
```

Example: Get servers by friendly name

```powershell
PS C:\> Get-AppvPublishingServer -Name "Server*"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvPublishingServer.md)


### Get-AppvStatus

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Get-AppvStatus [<CommonParameters>]
```

Example: Get status

```powershell
PS C:\> Get-AppvStatus
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Get-AppvStatus.md)


### Get-AppxDefaultVolume

Version: Both

Module: Appx

Syntax:

```powershell
Get-AppxDefaultVolume [<CommonParameters>]
```

Example: Get the default volume

```powershell
Get-AppxDefaultVolume
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Get-AppxDefaultVolume.md)


### Get-AppxPackage

Version: Both

Module: Appx

Syntax:

```powershell
Get-AppxPackage [[-Name] <string>] [[-Publisher] <string>] [-AllUsers] [-PackageTypeFilter <PackageTypes>] [-User <string>] [-Volume <AppxVolume>] [<CommonParameters>]
```

Example: Get all app packages for every user account

```powershell
Get-AppxPackage -AllUsers
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Get-AppxPackage.md)


### Get-AppxPackageAutoUpdateSettings

Version: Both

Module: Appx

Syntax:

```powershell
Get-AppxPackageAutoUpdateSettings [[-PackageFamilyName] <string>] [-ShowUpdateAvailability] [-AllUsers] [<CommonParameters>]
```

Example: Get all App Package Auto Update settings

```powershell
Get-AppxPackageAutoUpdateSettings
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Get-AppxPackageAutoUpdateSettings.md)


### Get-AppxPackageManifest

Version: Both

Module: Appx

Syntax:

```powershell
Get-AppxPackageManifest [-Package] <string> [[-User] <string>] [<CommonParameters>]
```

Example: Get the manifest for an app package

```powershell
Get-AppxPackageManifest -Package "package1_1.0.0.0_neutral__8wekyb3d8bbwe"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Get-AppxPackageManifest.md)


### Get-AppxProvisionedPackage

Version: Both

Module: Dism

Syntax:

```powershell
Get-AppxProvisionedPackage -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-AppxProvisionedPackage -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: List the app packages the mounted image to install for each account

```powershell
PS C:\> Get-AppxProvisionedPackage -Path "c:\offline"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-AppxProvisionedPackage.md)


### Get-AppxVolume

Version: Both

Module: Appx

Syntax:

```powershell
Get-AppxVolume [[-Path] <string>] [<CommonParameters>]
Get-AppxVolume -Online [-Path <string>] [<CommonParameters>]
Get-AppxVolume -Offline [-Path <string>] [<CommonParameters>]
```

Example: Get all the volumes

```powershell
Get-AppxVolume
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Get-AppxVolume.md)


### Get-AuthenticodeSignature

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Get-AuthenticodeSignature [-FilePath] <string[]> [<CommonParameters>]
Get-AuthenticodeSignature -LiteralPath <string[]> [<CommonParameters>]
Get-AuthenticodeSignature -SourcePathOrExtension <string[]> -Content <byte[]> [<CommonParameters>]
```

Example: Get the Authenticode signature for a file

```powershell
Get-AuthenticodeSignature -FilePath "C:\Test\NewScript.ps1"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Get-AuthenticodeSignature.md)


### Get-BcdEntry

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Get-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Get-BcdEntry.md)


### Get-BcdEntryDebugSettings

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Get-BcdEntryDebugSettings [[-Store] <BcdStoreInfo>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Get-BcdEntryDebugSettings.md)


### Get-BcdEntryHypervisorSettings

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Get-BcdEntryHypervisorSettings [[-Store] <BcdStoreInfo>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Get-BcdEntryHypervisorSettings.md)


### Get-BcdStore

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Get-BcdStore [[-Path] <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Get-BcdStore.md)


### Get-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Get-BitsTransfer [[-Name] <string[]>] [-AllUsers] [<CommonParameters>]
Get-BitsTransfer [-JobId] <guid[]> [<CommonParameters>]
```

Example: Get all BitsJob objects owned by the current user

```powershell
PS C:\> Get-BitsTransfer

JobId                   DisplayName             TransferType            JobState                OwnerAccount
-----                   -----------             ------------            --------                ------------
07acbe90-7d25-4d05-a... TestJob2                Download                Suspended               DOMAIN01\user01
c0dd3d8c-c3a2-4562-8... TestJob1                Download                Transferred             DOMAIN01\user01
1ef8c549-7a92-4173-b... BitsJobTransfer         Download                Transferred             DOMAIN01\user01
2c8302d5-3f44-4981-8... BitsJobTransfer         Download                Transferred             DOMAIN01\user01
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Get-BitsTransfer.md)


### Get-Certificate

Version: Both

Module: PKI

Syntax:

```powershell
Get-Certificate -Template <string> [-Url <uri>] [-SubjectName <string>] [-DnsName <string[]>] [-Credential <PkiCredential>] [-CertStoreLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Get-Certificate -Request <Certificate> [-Credential <PkiCredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$cred = Get-Credential

$params = @{
    Template = 'SslWebServer'
    DnsName = 'www.contoso.com', 'www.fabrikam.com'
    Url = 'https://www.contoso.com/Policy/service.svc'
    Credential = $cred
    CertStoreLocation = 'Cert:\LocalMachine\My'
}
Get-Certificate @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Get-Certificate.md)


### Get-CertificateAutoEnrollmentPolicy

Version: Both

Module: PKI

Syntax:

```powershell
Get-CertificateAutoEnrollmentPolicy -Scope <AutoEnrollmentPolicyScope> -context <Context> [<CommonParameters>]
```

Example: 

```powershell
Get-CertificateAutoEnrollmentPolicy -Scope Local -Context User
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Get-CertificateAutoEnrollmentPolicy.md)


### Get-CertificateEnrollmentPolicyServer

Version: Both

Module: PKI

Syntax:

```powershell
Get-CertificateEnrollmentPolicyServer -Scope <EnrollmentPolicyServerScope> -context <Context> [-Url <uri>] [<CommonParameters>]
```

Example: 

```powershell
Get-CertificateEnrollmentPolicyServer -Scope All -Context User
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Get-CertificateEnrollmentPolicyServer.md)


### Get-CertificateNotificationTask

Version: Both

Module: PKI

Syntax:

```powershell
Get-CertificateNotificationTask [<CommonParameters>]
```

Example: 

```powershell
Get-CertificateNotificationTask
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Get-CertificateNotificationTask.md)


### Get-CimAssociatedInstance

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ResourceUri <uri>] [-ComputerName <string[]>] [-KeyOnly] [<CommonParameters>]
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] -CimSession <CimSession[]> [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ResourceUri <uri>] [-KeyOnly] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ResourceUri <uri>] [-ComputerName <string[]>] [-KeyOnly] [<CommonParameters>]
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] -CimSession <CimSession[]> [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ResourceUri <uri>] [-KeyOnly] [<CommonParameters>]
```

Example: Get all the associated instances of a specific instance

```powershell
$disk = Get-CimInstance -ClassName Win32_LogicalDisk -KeyOnly
Get-CimAssociatedInstance -InputObject $disk[1]
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Get-CimAssociatedInstance.md)


### Get-CimClass

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] [-Amended] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] -CimSession <CimSession[]> [-Amended] [-OperationTimeoutSec <uint>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
```

Example: Get all the class definitions

```powershell
Get-CimClass
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Get-CimClass.md)


### Get-CimInstance

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Get-CimInstance [-ClassName] <string> [-ComputerName <string[]>] [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance -CimSession <CimSession[]> -ResourceUri <uri> [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance -CimSession <CimSession[]> -Query <string> [-ResourceUri <uri>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-Shallow] [<CommonParameters>]
Get-CimInstance [-ClassName] <string> -CimSession <CimSession[]> [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [<CommonParameters>]
Get-CimInstance [-InputObject] <ciminstance> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [<CommonParameters>]
Get-CimInstance -ResourceUri <uri> [-ComputerName <string[]>] [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance -Query <string> [-ResourceUri <uri>] [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-Shallow] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-CimInstance [-ClassName] <string> [-ComputerName <string[]>] [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [<CommonParameters>]
Get-CimInstance -CimSession <CimSession[]> -Query <string> [-ResourceUri <uri>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-Shallow] [<CommonParameters>]
Get-CimInstance [-ClassName] <string> -CimSession <CimSession[]> [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance -CimSession <CimSession[]> -ResourceUri <uri> [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance -ResourceUri <uri> [-ComputerName <string[]>] [-KeyOnly] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-Shallow] [-Filter <string>] [-Property <string[]>] [<CommonParameters>]
Get-CimInstance [-InputObject] <ciminstance> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [<CommonParameters>]
Get-CimInstance -Query <string> [-ResourceUri <uri>] [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-Shallow] [<CommonParameters>]
```

Example: Get the CIM instances of a specified class

```powershell
Get-CimInstance -ClassName Win32_Process
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Get-CimInstance.md)


### Get-CimSession

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Get-CimSession [[-ComputerName] <string[]>] [<CommonParameters>]
Get-CimSession [-Id] <uint32[]> [<CommonParameters>]
Get-CimSession -InstanceId <guid[]> [<CommonParameters>]
Get-CimSession -Name <string[]> [<CommonParameters>]
```

Syntax (7):

```powershell
Get-CimSession [[-ComputerName] <string[]>] [<CommonParameters>]
Get-CimSession [-Id] <uint[]> [<CommonParameters>]
Get-CimSession -InstanceId <guid[]> [<CommonParameters>]
Get-CimSession -Name <string[]> [<CommonParameters>]
```

Example: Get CIM sessions from the current PowerShell session

```powershell
New-CimSession -ComputerName Server01, Server02
Get-CimSession
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Get-CimSession.md)


### Get-CIPolicy

Version: Both

Module: ConfigCI

Syntax:

```powershell
Get-CIPolicy [-FilePath] <string> [<CommonParameters>]
```

Example: Get rules from a policy

```powershell
PS C:\> Get-CIPolicy -FilePath '.\Policy.xml'
Name           : MSIT Test CodeSign CA 3
Id             : ID_SIGNER_S_17
TypeId         : Allow
Root           : FA6B9A2230CE08BCA81D096B28CF495672401D3A43A0D285CF352464A6C9C7FD
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : VeriSign Class 3 Code Signing 2010 CA
Id             : ID_SIGNER_S_1D
TypeId         : Allow
Root           : 4843A82ED3B1F2BFBEE9671960E1940C942F688D
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : Microsoft Windows Third Party Component CA 2012
Id             : ID_SIGNER_S_1E
TypeId         : Allow
Root           : CEC1AFD0E310C55C1DCC601AB8E172917706AA32FB5EAF826813547FDF02DD46
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Get-CIPolicy.md)


### Get-CIPolicyIdInfo

Version: Both

Module: ConfigCI

Syntax:

```powershell
Get-CIPolicyIdInfo [-FilePath] <string> [<CommonParameters>]
```

Example: Display Code Integrity policy information

```powershell
PS C:\> Get-CIPolicyIdInfo -FilePath ".\Policy03.xml"

Provider  : ConfigCIPolicy
Key       : PolicyInfo
ValueName : Name
ValueType : String
Value     : CIPolicy03

Provider  : ConfigCIPolicy
Key       : PolicyInfo
ValueName : PolicyId
ValueType : String
Value     : CIP077
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Get-CIPolicyIdInfo.md)


### Get-CIPolicyInfo

Version: Both

Module: ConfigCI

Syntax:

```powershell
Get-CIPolicyInfo [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Get-CIPolicyInfo.md)


### Get-ComputerInfo

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-ComputerInfo [[-Property] <string[]>] [<CommonParameters>]
```

Example: Get all computer properties

```powershell
Get-ComputerInfo
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-ComputerInfo.md)

#### Implementation in PowerShell For Linux:

- Differs from the original: Reduced field set.

- Type: Go implementation.
- Function: gathers system information. Bash's `uname -a` + distro info.
- Parameters: none.
- Implementation: reads /etc/os-release (NAME, VERSION_ID) and /proc/meminfo (MemTotal), plus runtime GOOS/GOARCH, `os.Hostname()`, and `runtime.NumCPU()`.
- Output: a ComputerInfo object with fields CsName, OsName, OsVersion, OsArchitecture, OsPlatform, CsTotalPhysicalMemory, CsProcessors.


### Get-ComputerRestorePoint

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-ComputerRestorePoint [[-RestorePoint] <int[]>] [<CommonParameters>]
Get-ComputerRestorePoint -LastStatus [<CommonParameters>]
```

Example (5.1): Get all system restore points

```powershell
Get-ComputerRestorePoint
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Get-ComputerRestorePoint.md)


### Get-ControlPanelItem

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-ControlPanelItem [[-Name] <string[]>] [-Category <string[]>] [<CommonParameters>]
Get-ControlPanelItem -CanonicalName <string[]> [-Category <string[]>] [<CommonParameters>]
```

Example (5.1): Get all control panel items

```powershell
Get-ControlPanelItem
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Get-ControlPanelItem.md)


### Get-Counter

Version: Both

Module: Microsoft.PowerShell.Diagnostics

Syntax:

```powershell
Get-Counter [[-Counter] <string[]>] [-SampleInterval <int>] [-MaxSamples <long>] [-Continuous] [-ComputerName <string[]>] [<CommonParameters>]
Get-Counter [-ListSet] <string[]> [-ComputerName <string[]>] [<CommonParameters>]
```

Example: Get the counter set list

```powershell
Get-Counter -ListSet *
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Diagnostics/Get-Counter.md)


### Get-DAPolicyChange

Version: Both

Module: NetSecurity

Syntax:

```powershell
Get-DAPolicyChange [[-Servers] <string[]>] [[-Domains] <string[]>] [-DisplayName] <string> [[-PolicyStore] <string>] [-PSLocation] <string> [-EndpointType] <string> [[-DnsServers] <string[]>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\>Get-DAPolicyChange -DisplayName "TunnelPolicy1" -EndpointType Endpoint1 -PSLocation "C:\Update.ps1" -Servers "server1.corp.contoso.com", "server2.corp.contoso.com", "server3.corp.contoso.com"
IPsec Rule name  : TunnelPolicy1
Action           : Add
IPv6addresses    : 2001:4829:3243::100:1
                 : 2001:4829:3243::100:1
GPO              : contoso\DAClientPolicy

IPsec Rule name  : TunnelPolicy1
Action           : Delete
IPv6addresses    : 2001:4829:3243::100:3
                 : 2001:4829:3243::100:4
GPO              : contoso\DAClientPolicy

FQDN's that did not resolve into IP address:
server1.corp.contoso.com
server3.corp.contoso.com
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/NetSecurity/Get-DAPolicyChange.md)


### Get-DeliveryOptimizationLog

Version: Both

Module: DeliveryOptimization

Syntax (5.1):

```powershell
Get-DeliveryOptimizationLog [[-Path] <string[]>] [-LevelFilter <uint32>] [-Provider <ProviderType>] [-Flush] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-DeliveryOptimizationLog [[-Path] <string[]>] [-LevelFilter <uint>] [-Provider <ProviderType>] [-Flush] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/DeliveryOptimization/Get-DeliveryOptimizationLog.md)


### Get-DeliveryOptimizationLogAnalysis

Version: Both

Module: DeliveryOptimization

Syntax:

```powershell
Get-DeliveryOptimizationLogAnalysis [[-Path] <string[]>] [-ListConnections] [-Flush] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/DeliveryOptimization/Get-DeliveryOptimizationLogAnalysis.md)


### Get-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-EventLog [-LogName] <string> [[-InstanceId] <long[]>] [-ComputerName <string[]>] [-Newest <int>] [-After <datetime>] [-Before <datetime>] [-UserName <string[]>] [-Index <int[]>] [-EntryType <string[]>] [-Source <string[]>] [-Message <string>] [-AsBaseObject] [<CommonParameters>]
Get-EventLog [-ComputerName <string[]>] [-List] [-AsString] [<CommonParameters>]
```

Example (5.1): Get event logs on the local computer

```powershell
Get-EventLog -List
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Get-EventLog.md)


### Get-HotFix

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-HotFix [[-Id] <string[]>] [-ComputerName <string[]>] [-Credential <pscredential>] [<CommonParameters>]
Get-HotFix [-Description <string[]>] [-ComputerName <string[]>] [-Credential <pscredential>] [<CommonParameters>]
```

Example: Get all hotfixes on the local computer

```powershell
Get-HotFix
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-HotFix.md)


### Get-InstalledLanguage

Version: Both

Module: LanguagePackManagement

Syntax:

```powershell
Get-InstalledLanguage [[-Language] <string>] [<CommonParameters>]
```

Example: See what languages are installed on a device

```powershell
Get-InstalledLanguage
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LanguagePackManagement/Get-InstalledLanguage.md)


### Get-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Get-JobTrigger [-InputObject] <ScheduledJobDefinition> [[-TriggerId] <int[]>] [<CommonParameters>]
Get-JobTrigger [-Id] <int> [[-TriggerId] <int[]>] [<CommonParameters>]
Get-JobTrigger [-Name] <string> [[-TriggerId] <int[]>] [<CommonParameters>]
```

Example (5.1): Get a job trigger by scheduled job name

```powershell
Get-JobTrigger -Name "BackupJob"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Get-JobTrigger.md)


### Get-KdsConfiguration

Version: Both

Module: Kds

Syntax:

```powershell
Get-KdsConfiguration [<CommonParameters>]
```

Example: Retrieve the current KDS configuration

```powershell
PS C:\> Get-KdsConfiguration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/kds/Get-KdsConfiguration.md)


### Get-KdsRootKey

Version: Both

Module: Kds

Syntax:

```powershell
Get-KdsRootKey [<CommonParameters>]
```

Example: Retrieve a list of root key values

```powershell
PS C:\> Get-KdsRootKey
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/kds/Get-KdsRootKey.md)


### Get-LapsADPassword

Version: Both

Module: LAPS

Syntax:

```powershell
Get-LapsADPassword [-Identity] <string[]> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -Domain <string> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -DomainController <string> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -Port <int> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [-DomainController <string>] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -RecoveryMode [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -RecoveryMode -Port <int> [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
```

Example: 

```powershell
Get-LapsADPassword LAPSCLIENT
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Get-LapsADPassword.md)


### Get-LocalGroup

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Get-LocalGroup [[-Name] <string[]>] [<CommonParameters>]
Get-LocalGroup [[-SID] <SecurityIdentifier[]>] [<CommonParameters>]
```

Example (5.1): Get the Administrators group

```powershell
Get-LocalGroup -Name "Administrators"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Get-LocalGroup.md)


### Get-LocalGroupMember

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Get-LocalGroupMember [-Name] <string> [[-Member] <string>] [<CommonParameters>]
Get-LocalGroupMember [-Group] <LocalGroup> [[-Member] <string>] [<CommonParameters>]
Get-LocalGroupMember [-SID] <SecurityIdentifier> [[-Member] <string>] [<CommonParameters>]
```

Example (5.1): Get all members of the Administrators group

```powershell
Get-LocalGroupMember -Group "Administrators"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Get-LocalGroupMember.md)


### Get-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Get-LocalUser [[-Name] <string[]>] [<CommonParameters>]
Get-LocalUser [[-SID] <SecurityIdentifier[]>] [<CommonParameters>]
```

Example (5.1): Get an account by using its name

```powershell
Get-LocalUser -Name "AdminContoso02"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Get-LocalUser.md)


### Get-NonRemovableAppsPolicy

Version: Both

Module: Dism

Syntax:

```powershell
Get-NonRemovableAppsPolicy -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-NonRemovableAppsPolicy -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Get all installed non-removable app packages

```powershell
PS> Get-NonRemovableAppsPolicy -Online
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-NonRemovableAppsPolicy.md)


### Get-OSConfiguration

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OSConfiguration [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Get-OSConfiguration.md)


### Get-OsConfigurationDocument

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OsConfigurationDocument [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
Get-OsConfigurationDocument [[-Id] <string>] [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Get-OsConfigurationDocument.md)


### Get-OsConfigurationDocumentContent

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OsConfigurationDocumentContent [-Id] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Get-OsConfigurationDocumentContent.md)


### Get-OsConfigurationDocumentResult

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OsConfigurationDocumentResult [-Id] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [[-VerboseOption] <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Get-OsConfigurationDocumentResult.md)


### Get-OsConfigurationProperty

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OsConfigurationProperty [-Name] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [[-Id] <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Get-OsConfigurationProperty.md)


### Get-OSConfigurationScenarioDefinition

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OsConfigurationScenarioDefinition [-Name] <string> [-Version] <string> [-SchemaVersion] <string> [<CommonParameters>]
```

Example: none

Source: [OsConfiguration module documentation](https://learn.microsoft.com/en-us/powershell/module/osconfiguration) (no dedicated page).


### Get-OSConfigurationScenarioDefinitionInfo

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Get-OsConfigurationScenarioDefinitionInfo [[-Name] <string>] [[-Version] <string>] [[-SchemaVersion] <string>] [<CommonParameters>]
```

Example: none

Source: [OsConfiguration module documentation](https://learn.microsoft.com/en-us/powershell/module/osconfiguration) (no dedicated page).


### Get-PfxData

Version: Both

Module: PKI

Syntax:

```powershell
Get-PfxData [-FilePath] <string> [-Password <securestring>] [<CommonParameters>]
```

Example: 

```powershell
$mypwd = ConvertTo-SecureString -String '1234' -Force -AsPlainText

$mypfx = Get-PfxData -FilePath C:\mypfx.pfx -Password $mypwd
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Get-PfxData.md)


### Get-PmemDedicatedMemory

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Get-PmemDedicatedMemory [<CommonParameters>]
Get-PmemDedicatedMemory [[-DeviceNumber] <uint32[]>] [<CommonParameters>]
```

Example: Get all dedicated persistent memory

```powershell
Get-PmemDedicatedMemory
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Get-PmemDedicatedMemory.md)


### Get-PmemDisk

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Get-PmemDisk [<CommonParameters>]
Get-PmemDisk [[-DiskNumber] <uint32[]>] [<CommonParameters>]
Get-PmemDisk [-PhysicalDevice <PmemPhysicalDevice>] [<CommonParameters>]
Get-PmemDisk [-PhysicalDeviceId <string[]>] [<CommonParameters>]
Get-PmemDisk [-InputObject <ciminstance>] [<CommonParameters>]
```

Example: Get persistent memory disks

```powershell
Get-PmemDisk
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Get-PmemDisk.md)


### Get-PmemPhysicalDevice

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Get-PmemPhysicalDevice [<CommonParameters>]
Get-PmemPhysicalDevice [[-DeviceId] <string[]>] [<CommonParameters>]
Get-PmemPhysicalDevice [-LogicalDisk <PmemDisk>] [<CommonParameters>]
Get-PmemPhysicalDevice [-DiskNumber <uint32>] [<CommonParameters>]
Get-PmemPhysicalDevice [-InputObject <ciminstance>] [<CommonParameters>]
```

Example: Get physical devices that have persistent memory

```powershell
Get-PmemPhysicalDevice
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Get-PmemPhysicalDevice.md)


### Get-PmemUnusedRegion

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Get-PmemUnusedRegion [[-RegionId] <uint32[]>] [<CommonParameters>]
```

Example: Get unused regions

```powershell
Get-PmemUnusedRegion
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Get-PmemUnusedRegion.md)


### Get-ProcessMitigation

Version: Both

Module: ProcessMitigations

Syntax:

```powershell
Get-ProcessMitigation [-FullPolicy] [<CommonParameters>]
Get-ProcessMitigation [-Name] <string> [-RunningProcesses] [<CommonParameters>]
Get-ProcessMitigation [-Id] <int[]> [<CommonParameters>]
Get-ProcessMitigation [-RegistryConfigFilePath <string>] [<CommonParameters>]
Get-ProcessMitigation [-System] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Get-ProcessMitigation -Name notepad.exe -RunningProcess
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ProcessMitigations/Get-ProcessMitigation.md)


### Get-ProvisioningPackage

Version: Both

Module: Provisioning

Syntax:

```powershell
Get-ProvisioningPackage [-PackageId] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Get-ProvisioningPackage [-PackagePath] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Get-ProvisioningPackage [-AllInstalledPackages] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Get-ProvisioningPackage -PackagePath c:\test\testppkg.ppkg
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Get-ProvisioningPackage.md)


### Get-PSSessionCapability

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-PSSessionCapability [-ConfigurationName] <string> [-Username] <string> [-Full] [<CommonParameters>]
```

Example: Get commands available for a user

```powershell
Get-PSSessionCapability -ConfigurationName Endpoint1 -Username 'CONTOSO\User'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-PSSessionCapability.md)


### Get-PSSessionConfiguration

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-PSSessionConfiguration [[-Name] <string[]>] [-Force] [<CommonParameters>]
```

Example: Get session configurations on the local computer

```powershell
Get-PSSessionConfiguration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Get-PSSessionConfiguration.md)


### Get-PSSnapin

Version: 5.1 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Get-PSSnapin [[-Name] <string[]>] [-Registered] [<CommonParameters>]
```

Example: Get snap-ins that are currently loaded

```powershell
PS C:\> Get-PSSnapin
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Get-PSSnapin.md)


### Get-RecoveryManagementPluginAltitude

Version: Both

Module: Dism

Syntax:

```powershell
Get-RecoveryManagementPluginAltitude -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryManagementPluginAltitude -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Get-RecoveryManagementPluginInfo

Version: Both

Module: Dism

Syntax:

```powershell
Get-RecoveryManagementPluginInfo -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryManagementPluginInfo -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Get-RecoveryManagementPlugins

Version: Both

Module: Dism

Syntax:

```powershell
Get-RecoveryManagementPlugins -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryManagementPlugins -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Get-RecoveryRemoteManagementStatus

Version: Both

Module: Dism

Syntax:

```powershell
Get-RecoveryRemoteManagementStatus -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryRemoteManagementStatus -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Get-ReFSDedupSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Get-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Get-ReFSDedupSchedule -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Get-ReFSDedupSchedule.md)


### Get-ReFSDedupScrubSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Get-ReFSDedupScrubSchedule [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Get-ReFSDedupScrubSchedule -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Get-ReFSDedupScrubSchedule.md)


### Get-ReFSDedupStatus

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Get-ReFSDedupStatus [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Get-ReFSDedupStatus -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Get-ReFSDedupStatus.md)


### Get-ScheduledJob

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Get-ScheduledJob [[-Id] <int[]>] [<CommonParameters>]
Get-ScheduledJob [-Name] <string[]> [<CommonParameters>]
```

Example (5.1): Get all scheduled jobs

```powershell
Get-ScheduledJob
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Get-ScheduledJob.md)


### Get-ScheduledJobOption

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Get-ScheduledJobOption [-InputObject] <ScheduledJobDefinition> [<CommonParameters>]
Get-ScheduledJobOption [-Id] <int> [<CommonParameters>]
Get-ScheduledJobOption [-Name] <string> [<CommonParameters>]
```

Example (5.1): Get job options

```powershell
Get-ScheduledJobOption -Name "*Backup*"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Get-ScheduledJobOption.md)


### Get-SecureBootPolicy

Version: Both

Module: SecureBoot

Syntax:

```powershell
Get-SecureBootPolicy [<CommonParameters>]
```

Example: Get Secure Boot policy

```powershell
PS C:\> Get-SecureBootPolicy | Format-List
Publisher: 77fa9abd-0359-4d32-bd60-28f4e78f784b
Version  : 1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/SecureBoot/Get-SecureBootPolicy.md)


### Get-SecureBootSVN

Version: Both

Module: SecureBoot

Syntax:

```powershell
Get-SecureBootSVN [-BootManagerPath <string>] [<CommonParameters>]
```

Example: none

Source: [SecureBoot module documentation](https://learn.microsoft.com/en-us/powershell/module/secureboot) (no dedicated page).


### Get-SecureBootUEFI

Version: Both

Module: SecureBoot

Syntax:

```powershell
Get-SecureBootUEFI [-Name] <string> [-OutputFilePath <string>] [-Decoded] [<CommonParameters>]
```

Example: Get information about PK

```powershell
PS C:\>Get-SecureBootUefi -Name PK | Format-List
Name       : PK
Bytes      : {161, 89, 192, 165...}
Attributes : NON VOLATILE
             BOOTSERVICE ACCESS
             RUNTIME ACCESS
             TIME BASED AUTHENTICATED WRITE ACCESS
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/SecureBoot/Get-SecureBootUEFI.md)


### Get-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Get-Service [[-Name] <string[]>] [-ComputerName <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service -DisplayName <string[]> [-ComputerName <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service [-ComputerName <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [-InputObject <ServiceController[]>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-Service [[-Name] <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service -DisplayName <string[]> [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [-InputObject <ServiceController[]>] [<CommonParameters>]
```

Example: Get all services on the computer

```powershell
Get-Service
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Get-Service.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`systemctl`).
- Distro: systemd-based.
- Function: lists services. Maps to `systemctl list-units --type=service --all --no-pager`.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Filters by service name (wildcards supported), such as `Get-Service -Name ssh*` |

- Implementation: calls external `systemctl` (adding `sudo` automatically when permissions fall short), parsing columns 0/2 of the output (unit name, status).
- Output: ServiceController objects with fields Name (with the .service suffix stripped), Status (active/inactive/...), DisplayName (same as Name). Table columns Status/Name/DisplayName.
- No systemctl → error.


### Get-SystemDriver

Version: Both

Module: ConfigCI

Syntax:

```powershell
Get-SystemDriver [-Audit] [-ScanPath <string>] [-UserPEs] [-NoScript] [-NoShadowCopy] [-OmitPaths <string[]>] [-PathToCatroot <string>] [-ScriptFileNames] [<CommonParameters>]
```

Example: Scan a folder for drivers

```powershell
PS C:\> Get-SystemDriver -ScanPath '.\temp' -UserPEs

FilePath     : \\?\GLOBALROOT\Device\HarddiskVolumeShadowCopy9\cmdlets\temp\ConfigCI.psd1
FriendlyName : \\?\E:\cmdlets\temp\ConfigCI.psd1
FileName     :
Loaded       : False
FileVersion  :
Hash         : 1844B4531711EC9170A9D33277CE1D4FF7626C54
Hash256      : 60311157F6685727F42CC04717FEF6F905EC2A317C3B8381CDD9A79D0B184483
PageHash     :
PageHash256  :
UserMode     : True
OpusInfos    : {}
Signers      : {}

FilePath     : \\?\GLOBALROOT\Device\HarddiskVolumeShadowCopy9\cmdlets\temp\Microsoft.ConfigCI.Commands.dll
FriendlyName : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll
FileName     : Microsoft.ConfigCI.Commands.dll
Loaded       : False
FileVersion  : 10.0.10543.1000
Hash         : BE0777F5AF88628D4555A875036648DF1AD19BBE
Hash256      : 6FA5AF724499C338A77FEEAD90F55DDF5F23D081C6DCE8E9DF486E95C6A9B310
PageHash     : D41570F2E6E7E6245CF342131D4706C944562B1E
PageHash256  : F714D9784E15B88F56180C8EE2B40C769CC83428954585A1DCF9A260FE967CDD
UserMode     : False
OpusInfos    : {}
Signers      : {}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Get-SystemDriver.md)


### Get-SystemPreferredUILanguage

Version: Both

Module: LanguagePackManagement

Syntax:

```powershell
Get-SystemPreferredUILanguage [<CommonParameters>]
```

Example: 

```powershell
Get-SystemPreferredUILanguage
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LanguagePackManagement/Get-SystemPreferredUILanguage.md)


### Get-TlsCipherSuite

Version: Both

Module: TLS

Syntax:

```powershell
Get-TlsCipherSuite [[-Name] <string>] [<CommonParameters>]
```

Example: Get all cipher suites

```powershell
Get-TlsCipherSuite
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Get-TlsCipherSuite.md)


### Get-TlsEccCurve

Version: Both

Module: TLS

Syntax:

```powershell
Get-TlsEccCurve [[-Name] <string>] [<CommonParameters>]
```

Example: Get all ECC curves

```powershell
Get-TlsEccCurve
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/Get-TlsEccCurve.md)


### Get-Tpm

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Get-Tpm [<CommonParameters>]
```

Example: Display TPM information

```powershell
PS C:\> Get-Tpm
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Get-Tpm.md)


### Get-TpmEndorsementKeyInfo

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Get-TpmEndorsementKeyInfo [[-HashAlgorithm] <string>] [<CommonParameters>]
```

Example: Get endorsement key information

```powershell
PS C:\> Get-TpmEndorsementKeyInfo -Hash "Sha256"
IsPresent                : True
PublicKey                : System.Security.Cryptography.AsnEncodedData
PublicKeyHash            : 70769c52b6e24ef683693c2a0208da68d77e94192e1f4080ae7c9b97c6caa681
ManufacturerCertificates : {[Subject]
OID.2.23.133.2.3=1.2,
OID.2.23.133.2.2=C4T8SOX3.5,
OID.2.23.133.2.1=id:782F345A

[Issuer]
CN=Contoso TPM CA1, OU=Contoso
Certification Authority, O=Contoso, C=KR

[Serial Number]
77A120A

[Not Before]
6/4/2012 6:35:58 PM

[Not After]
6/4/2022 6:35:57 PM

[Thumbprint]
77378D1480AB48FEA2D4E610B2C7EEF648FEA2
}
AdditionalCertificates   : {}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Get-TpmEndorsementKeyInfo.md)


### Get-TpmSupportedFeature

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Get-TpmSupportedFeature [[-FeatureList] <StringCollection>] [<CommonParameters>]
```

Example: Verify support for key attestation

```powershell
PS C:\> Get-TpmSupportedFeature -FeatureList "Key Attestation"
key attestation
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Get-TpmSupportedFeature.md)


### Get-Transaction

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-Transaction [<CommonParameters>]
```

Example (5.1): Get the current transaction

```powershell
Start-Transaction
Get-Transaction
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Get-Transaction.md)


### Get-TroubleshootingPack

Version: Both

Module: TroubleshootingPack

Syntax:

```powershell
Get-TroubleshootingPack [-Path] <string> [-AnswerFile <string>] [<CommonParameters>]
```

Example: Get a troubleshooting pack

```powershell
PS C:\> Get-TroubleshootingPack -Path "C:\Windows\Diagnostics\System\Audio"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TroubleshootingPack/Get-TroubleshootingPack.md)


### Get-TrustedProvisioningCertificate

Version: Both

Module: Provisioning

Syntax:

```powershell
Get-TrustedProvisioningCertificate [[-Thumbprint] <string>] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: List installed trusted provisioning certificates

```powershell
PS C:\> Get-TrustedProvisioningCertificate
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Get-TrustedProvisioningCertificate.md)


### Get-UevAppxPackage

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Get-UevAppxPackage [<CommonParameters>]
Get-UevAppxPackage -Computer [<CommonParameters>]
Get-UevAppxPackage -CurrentComputerUser [<CommonParameters>]
```

Example: Get the list of Windows 8 apps

```powershell
PS C:\>Get-UevAppxPackage -CurrentComputerUser
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Get-UevAppxPackage.md)


### Get-UevConfiguration

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Get-UevConfiguration [<CommonParameters>]
Get-UevConfiguration -Computer [<CommonParameters>]
Get-UevConfiguration -CurrentComputerUser [<CommonParameters>]
Get-UevConfiguration -Details [<CommonParameters>]
```

Example: Get the uev_tla configuration

```powershell
PS C:\> Get-UevConfiguration


Key                                     Value
---                                     -----
MaxPackageSizeInBytes                   700000
SettingsImportNotifyDelayInSeconds      10
SettingsImportNotifyEnabled             False
SettingsStoragePath                     \\ServerName\Path\To\CentralStore
SettingsTemplateCatalogPath
SyncEnabled                             True
SyncMethod                              OfflineFiles
SyncFromRepositoryTimeoutInMilliseconds 2000
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Get-UevConfiguration.md)


### Get-UevStatus

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Get-UevStatus [<CommonParameters>]
```

Example: Get status of the UE-V service

```powershell
PS C:\>Get-UevStatus
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Get-UevStatus.md)


### Get-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Get-UevTemplate [<CommonParameters>]
Get-UevTemplate -Application <string> [<CommonParameters>]
Get-UevTemplate -TemplateID <string> [<CommonParameters>]
Get-UevTemplate -Profile <string> [<CommonParameters>]
Get-UevTemplate [-ApplicationOrTemplateID] <string> [<CommonParameters>]
```

Example: Get all registered templates

```powershell
PS C:\> Get-UevTemplate | Format-Table -AutoSize
TemplateId                                  TemplateName                                TemplateVersion PackageVersion TemplateType Enabled EnableStateLocation TemplateDescription
----------                                  ------------                                --------------- -------------- ------------ ------- ------------------- -------------------
DesktopSettings                             Desktop Settings                                          1 N/A            OS             False LocalMachine
MicrosoftNotepad6                           Microsoft Notepad                                         0 N/A            Application     True NotSet
MicrosoftCalculator6                        Microsoft Calculator                                      0 N/A            Application     True NotSet
MicrosoftCommunicator2007                   Microsoft Communicator 2007                               7 N/A            Application     True NotSet
MicrosoftOffice2010Win64                    Microsoft Office 2010 (64-bit)                           18 N/A            Application     True NotSet
MicrosoftOffice2010Win64.common             Common Settings                                          18 N/A            Application     True NotSet
MicrosoftOffice2010Win64.Access             Microsoft Access 2010 (64-bit)                           18 N/A            Application     True NotSet
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Get-UevTemplate.md)


### Get-UevTemplateProgram

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Get-UevTemplateProgram [-ID] <string> [<CommonParameters>]
```

Example: Get all defined programs

```powershell
PS C:\> Get-UevTemplate | Get-UevTemplateProgram | Format-Table -AutoSize


TemplateId                          ProgramName      ProductVersionRange FileVersionRange
----------                          -----------      ------------------- ----------------
MicrosoftCalculator6                CALC.EXE         6-6
MicrosoftNotepad6                   NOTEPAD.EXE      6-6
MicrosoftOffice2010.OneNote         ONENOTE.EXE      14.0-14.0           14.0-14.0
MicrosoftOffice2010.Word            WINWORD.EXE      14.0-14.0           14.0-14.0
MicrosoftOffice2010.Excel           EXCEL.EXE        14.0-14.0           14.0-14.0
MicrosoftOffice2010.PowerPoint      POWERPNT.EXE     14.0-14.0           14.0-14.0
MicrosoftOffice2010.Outlook         OUTLOOK.EXE      14.0-14.0           14.0-14.0
MicrosoftOffice2010.InfoPath        INFOPATH.EXE     14.0-14.0           14.0-14.0
MicrosoftOffice2010.Visio           VISIO.EXE        14.0-14.0           14.0-14.0
MicrosoftOffice2010.Groove          Groove.exe       14.0-14.0           14.0-14.0
MicrosoftOffice2010.Access          MSACCESS.EXE     14.0-14.0           14.0-14.0
MicrosoftOffice2010.Project         WINPROJ.EXE      14.0-14.0           14.0-14.0
MicrosoftOffice2010.Publisher       MSPUB.EXE        14.0-14.0           14.0-14.0
MicrosoftWordpad6                   WORDPAD.EXE      6-6
MicrosoftInternetExplorer.Version8  iexplore.exe     8.0-8.0             8.0-8.0
MicrosoftInternetExplorer.Version9  iexplore.exe     9.0-9.0             9.0-9.0
MicrosoftInternetExplorer.Version10 iexplore.exe     10.0-10.0           10.0-10.0
MicrosoftLync2010                   communicator.exe 4.0-4.0             4.0-4.0
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Get-UevTemplateProgram.md)


### Get-WheaMemoryPolicy

Version: Both

Module: Whea

Syntax:

```powershell
Get-WheaMemoryPolicy [-ComputerName <string>] [<CommonParameters>]
```

Example: Get the WHEA memory policy settings from the local computer

```powershell
PS C:\> Get-WHEAMemoryPolicy
DisableOffline : False
DisablePFA : False
PersistMemoryOffline : True
PFAPageCount : 64
PFAErrorThreshold : 16
PFATimeOut : 86400
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Whea/Get-WheaMemoryPolicy.md)


### Get-WIMBootEntry

Version: Both

Module: Dism

Syntax:

```powershell
Get-WIMBootEntry -Path <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Displays the WIMBoot configuration for a drive

```powershell
PS C:\> Get-WIMBootEntry -Path "C:\"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WIMBootEntry.md)


### Get-WinAcceptLanguageFromLanguageListOptOut

Version: Both

Module: International

Syntax:

```powershell
Get-WinAcceptLanguageFromLanguageListOptOut [<CommonParameters>]
```

Example: Get the status of the setting

```powershell
PS C:\> Get-WinAcceptLanguageFromLanguageListOptOut
TRUE
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinAcceptLanguageFromLanguageListOptOut.md)


### Get-WinCultureFromLanguageListOptOut

Version: Both

Module: International

Syntax:

```powershell
Get-WinCultureFromLanguageListOptOut [<CommonParameters>]
```

Example: Get the Culture override setting

```powershell
PS C:\> Get-WinCultureFromLanguageListOptOut
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinCultureFromLanguageListOptOut.md)


### Get-WinDefaultInputMethodOverride

Version: Both

Module: International

Syntax:

```powershell
Get-WinDefaultInputMethodOverride [<CommonParameters>]
```

Example: Display default input method

```powershell
PS C:\> Get-WinDefaultInputMethodOverride
InputMethodTip      Keyboard name
---------------     -------------
0409:00000409       English (United States) - US
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinDefaultInputMethodOverride.md)


### Get-WindowsCapability

Version: Both

Module: Dism

Syntax:

```powershell
Get-WindowsCapability -Path <string> [-Name <string>] [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsCapability -Online [-Name <string>] [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Get the Windows capabilities for an image

```powershell
PS C:\> Get-WindowsCapability -Path "C:\offline" -Name "Language.TextToSpeech~~~fr-FR~0.0.1.0"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsCapability.md)


### Get-WindowsDeveloperLicense

Version: Both

Module: WindowsDeveloperLicense

Syntax:

```powershell
Get-WindowsDeveloperLicense [<CommonParameters>]
```

Example: Check the status of Developer Mode DM

```powershell
PS C:\> Get-WindowsDeveloperLicense
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsDeveloperLicense/Get-WindowsDeveloperLicense.md)


### Get-WindowsDriver

Version: Both

Module: Dism

Syntax:

```powershell
Get-WindowsDriver -Path <string> [-All] [-Driver <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsDriver -Online [-All] [-Driver <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Gets all drivers in an online image

```powershell
PS C:\> Get-WindowsDriver -Online -All
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsDriver.md)


### Get-WindowsEdition

Version: Both

Module: Dism

Syntax:

```powershell
Get-WindowsEdition -Path <string> [-Target] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsEdition -Online [-Target] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Display the current edition of the operating system

```powershell
PS C:\> Get-WindowsEdition -Online
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsEdition.md)


### Get-WindowsErrorReporting

Version: Both

Module: WindowsErrorReporting

Syntax:

```powershell
Get-WindowsErrorReporting [<CommonParameters>]
```

Example: Get the Windows Error Reporting status

```powershell
PS C:\> Get-WindowsErrorReporting
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsErrorReporting/Get-WindowsErrorReporting.md)


### Get-WindowsImage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Get-WindowsImage -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Index <uint32> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -Mounted [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-WindowsImage -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Index <uint> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -Mounted [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Get information about all mounted images

```powershell
PS C:\> Get-WindowsImage -Mounted
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsImage.md)


### Get-WindowsImageContent

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Get-WindowsImageContent -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> -Index <uint32> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Get-WindowsImageContent -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> -Index <uint> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: List files and folders for an image

```powershell
PS C:\> Get-WindowsImageContent -ImagePath "c:\imagestore\install.wim" -Index 1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsImageContent.md)


### Get-WindowsOptionalFeature

Version: Both

Module: Dism

Syntax:

```powershell
Get-WindowsOptionalFeature -Path <string> [-FeatureName <string>] [-PackageName <string>] [-PackagePath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsOptionalFeature -Online [-FeatureName <string>] [-PackageName <string>] [-PackagePath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Lists optional features in the running operating system

```powershell
PS C:\> Get-WindowsOptionalFeature -Online
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsOptionalFeature.md)


### Get-WindowsPackage

Version: Both

Module: Dism

Syntax:

```powershell
Get-WindowsPackage -Path <string> [-PackagePath <string>] [-PackageName <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsPackage -Online [-PackagePath <string>] [-PackageName <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Lists packages in a mounted image

```powershell
PS C:\> Get-WindowsPackage -Path "c:\offline"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsPackage.md)


### Get-WindowsReservedStorageState

Version: Both

Module: Dism

Syntax:

```powershell
Get-WindowsReservedStorageState [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Get-WindowsReservedStorageState
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Get-WindowsReservedStorageState.md)


### Get-WindowsSearchSetting

Version: Both

Module: WindowsSearch

Syntax:

```powershell
Get-WindowsSearchSetting [<CommonParameters>]
```

Example: Get Windows Search settings

```powershell
PS C:\> Get-WindowsSearchSetting
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsSearch/Get-WindowsSearchSetting.md)


### Get-WinEvent

Version: Both

Module: Microsoft.PowerShell.Diagnostics

Syntax:

```powershell
Get-WinEvent [[-LogName] <string[]>] [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-FilterXPath <string>] [-Force] [-Oldest] [<CommonParameters>]
Get-WinEvent [-ListLog] <string[]> [-ComputerName <string>] [-Credential <pscredential>] [-Force] [<CommonParameters>]
Get-WinEvent [-ListProvider] <string[]> [-ComputerName <string>] [-Credential <pscredential>] [<CommonParameters>]
Get-WinEvent [-ProviderName] <string[]> [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-FilterXPath <string>] [-Force] [-Oldest] [<CommonParameters>]
Get-WinEvent [-Path] <string[]> [-MaxEvents <long>] [-Credential <pscredential>] [-FilterXPath <string>] [-Oldest] [<CommonParameters>]
Get-WinEvent [-FilterHashtable] <hashtable[]> [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-Force] [-Oldest] [<CommonParameters>]
Get-WinEvent [-FilterXml] <xml> [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-Oldest] [<CommonParameters>]
```

Example: Get all the logs from a local computer

```powershell
Get-WinEvent -ListLog *
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Diagnostics/Get-WinEvent.md)


### Get-WinHomeLocation

Version: Both

Module: International

Syntax:

```powershell
Get-WinHomeLocation [<CommonParameters>]
```

Example: Display the GeoID for the current account

```powershell
PS C:\> Get-WinHomeLocation
HomeLocation     Description
----             -----------
244              United States
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinHomeLocation.md)


### Get-WinLanguageBarOption

Version: Both

Module: International

Syntax:

```powershell
Get-WinLanguageBarOption [<CommonParameters>]
```

Example: Get the settings for the language bar

```powershell
PS C:\> Get-WinLanguageBarOption
IsLegacyLanguageBar    IsLegacySwitchingMode
-------------------    ---------------------
False                  False
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinLanguageBarOption.md)


### Get-WinSystemLocale

Version: Both

Module: International

Syntax:

```powershell
Get-WinSystemLocale [<CommonParameters>]
```

Example: Get the system locale

```powershell
PS C:\> GET-WinSystemLocale
LCID             Name             DisplayName
----             ----             -----------
1033             en-US            English (United States)
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinSystemLocale.md)


### Get-WinUILanguageOverride

Version: Both

Module: International

Syntax:

```powershell
Get-WinUILanguageOverride [<CommonParameters>]
```

Example: Display the language override setting

```powershell
PS C:\> Get-WinUILanguageOverride
LCID             Name             DisplayName
----             ----             -----------
1033             en-US            English (United States)
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinUILanguageOverride.md)


### Get-WinUserLanguageList

Version: Both

Module: International

Syntax:

```powershell
Get-WinUserLanguageList [<CommonParameters>]
```

Example: Get language list for the current account

```powershell
PS C:\> Get-WinUserLanguageList
LanguageTag     : en-US
Autonym         : English (United States)
EnglishName     : English (United States)
LocalizedName   : English (United States)
ScriptName      : Latin
InputMethodTips : {0409:00000409}
Handwriting     : False
LanguageTag     : fr-FR
Autonym         : français (France)
EnglishName     : French (France)
LocalizedName   : French (France)
ScriptName      : Latin
InputMethodTips : {040c:0000040c}
Handwriting     : False
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Get-WinUserLanguageList.md)


### Get-WmiObject

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Get-WmiObject [-Class] <string> [[-Property] <string[]>] [-Filter <string>] [-Amended] [-DirectRead] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject [[-Class] <string>] [-Recurse] [-Amended] [-List] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject -Query <string> [-Amended] [-DirectRead] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject [-Amended] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject [-Amended] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
```

Example (5.1): Get processes on the local computer

```powershell
Get-WmiObject -Class Win32_Process
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Get-WmiObject.md)


### Get-WSManCredSSP

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Get-WSManCredSSP [<CommonParameters>]
```

Example: Display CredSSP configuration

```powershell
Get-WSManCredSSP
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Get-WSManCredSSP.md)


### Get-WSManInstance

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Get-WSManInstance [-ResourceURI] <uri> [-ApplicationName <string>] [-ComputerName <string>] [-ConnectionURI <uri>] [-Dialect <uri>] [-Fragment <string>] [-OptionSet <hashtable>] [-Port <int>] [-SelectorSet <hashtable>] [-SessionOption <SessionOption>] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Get-WSManInstance [-ResourceURI] <uri> -Enumerate [-ApplicationName <string>] [-BasePropertiesOnly] [-ComputerName <string>] [-ConnectionURI <uri>] [-Dialect <uri>] [-Filter <string>] [-OptionSet <hashtable>] [-Port <int>] [-Associations] [-ReturnType <string>] [-SessionOption <SessionOption>] [-Shallow] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Get all information from WMI

```powershell
Get-WSManInstance -ResourceURI wmicimv2/Win32_Service -SelectorSet @{name="winrm"} -ComputerName "Server01"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Get-WSManInstance.md)


### Import-BcdStore

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Import-BcdStore [-Path] <string> [-NoClean] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Import-BcdStore.md)


### Import-BinaryMiLog

Version: Both

Module: CimCmdlets

Syntax:

```powershell
Import-BinaryMiLog [-Path] <string> [<CommonParameters>]
```

Example (5.1): Restore objects exported to a file

```powershell
Import-BinaryMiLog -Path "Processes.bmil"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/CimCmdlets/Import-BinaryMiLog.md)


### Import-Certificate

Version: Both

Module: PKI

Syntax:

```powershell
Import-Certificate [-FilePath] <string> [-CertStoreLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    FilePath = 'C:\Users\Xyz\Desktop\BackupCert.cer'
    CertStoreLocation = 'Cert:\CurrentUser\Root'
}
Import-Certificate @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Import-Certificate.md)


### Import-Counter

Version: Both

Module: Microsoft.PowerShell.Diagnostics

Syntax:

```powershell
Import-Counter [-Path] <string[]> [-StartTime <datetime>] [-EndTime <datetime>] [-Counter <string[]>] [-MaxSamples <long>] [<CommonParameters>]
Import-Counter [-Path] <string[]> -ListSet <string[]> [<CommonParameters>]
Import-Counter [-Path] <string[]> [-Summary] [<CommonParameters>]
```

Example (5.1): Import all counter data from a file

```powershell
$data = Import-Counter -Path ProcessorData.csv
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Diagnostics/Import-Counter.md)


### Import-PfxCertificate

Version: Both

Module: PKI

Syntax:

```powershell
Import-PfxCertificate [-FilePath] <string> [[-CertStoreLocation] <string>] [-Exportable] [-ProtectPrivateKey <ProtectPrivateKeyType>] [-Password <securestring>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$mypwd = Get-Credential -UserName 'Enter password below' -Message 'Enter password below'

$params = @{
    FilePath = 'C:\mypfx.pfx'
    CertStoreLocation = 'Cert:\LocalMachine\My'
    Password = $mypwd.Password
}
Import-PfxCertificate @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Import-PfxCertificate.md)


### Import-StartLayout

Version: Both

Module: StartLayout

Syntax:

```powershell
Import-StartLayout [-LayoutPath] <string> [-MountPath] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Import-StartLayout -LayoutLiteralPath <string> -MountLiteralPath <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Import a layout into a Windows image

```powershell
PS C:\> Import-StartLayout -LayoutPath "Layout.xml" -MountPath "C:\"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/StartLayout/Import-StartLayout.md)


### Import-TpmOwnerAuth

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Import-TpmOwnerAuth -File <string> [<CommonParameters>]
Import-TpmOwnerAuth [-OwnerAuthorization] <string> [<CommonParameters>]
```

Example: Import an owner authorization value

```powershell
PS C:\> Import-TpmOwnerAuth -OwnerAuthorization "Qn2sdCFQmvjf+tBtSWH4GT87sQs="
TpmReady           : False
TpmPresent         : True
ManagedAuthLevel   : Full
OwnerAuth          : Qn2sdCFQmvjf+tBtSWH4GT87sQs=
OwnerClearDisabled : True
AutoProvisioning   : DisabledForNextBoot
LockedOut          : False
SelfTest           : {191, 191, 245, 191...}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Import-TpmOwnerAuth.md)


### Import-UevConfiguration

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Import-UevConfiguration [-Path] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Import the UE-V configuration

```powershell
PS C:\> Import-UevConfiguration -Path "ContosoUev.uev"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Import-UevConfiguration.md)


### Initialize-PmemPhysicalDevice

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Initialize-PmemPhysicalDevice -DeviceId <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Initialize physical devices

```powershell
Get-PmemPhysicalDevice | Initialize-PmemPhysicalDevice
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Initialize-PmemPhysicalDevice.md)


### Initialize-Tpm

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Initialize-Tpm [-AllowClear] [-AllowPhysicalPresence] [<CommonParameters>]
```

Example: Initialize a TPM

```powershell
PS C:\> Initialize-Tpm -AllowClear -AllowPhysicalPresence
TpmReady                 : False
RestartRequired          : True
ShutdownRequired         : False
ClearRequired            : True
PhysicalPresenceRequired : True
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Initialize-Tpm.md)


### Install-Language

Version: Both

Module: LanguagePackManagement

Syntax:

```powershell
Install-Language [-Language] <string> [-CopyToSettings] [-ExcludeFeatures] [-AsJob] [<CommonParameters>]
```

Example: Add a language to a device

```powershell
Install-Language ja-JP
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LanguagePackManagement/Install-Language.md)


### Install-ProvisioningPackage

Version: Both

Module: Provisioning

Syntax:

```powershell
Install-ProvisioningPackage [-PackagePath] <string> [-ForceInstall] [-QuietInstall] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Install-ProvisioningPackage -PackagePath C:\mypackage.ppkg -QuietInstall
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Install-ProvisioningPackage.md)


### Install-TrustedProvisioningCertificate

Version: Both

Module: Provisioning

Syntax:

```powershell
Install-TrustedProvisioningCertificate [-CertificatePath] <string> [-ForceInstall] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: Install Trusted Provisioning Certificate

```powershell
PS C:\> Install-TrustedProvisioningCertificate -CertificatePath trustedCert.cer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Install-TrustedProvisioningCertificate.md)


### Invoke-CimMethod

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Invoke-CimMethod [-ClassName] <string> [[-Arguments] <IDictionary>] [-MethodName] <string> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-ClassName] <string> [[-Arguments] <IDictionary>] [-MethodName] <string> -CimSession <CimSession[]> [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -ResourceUri <uri> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-InputObject] <ciminstance> [[-Arguments] <IDictionary>] [-MethodName] <string> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-InputObject] <ciminstance> [[-Arguments] <IDictionary>] [-MethodName] <string> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -ResourceUri <uri> -CimSession <CimSession[]> [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-CimClass] <cimclass> [[-Arguments] <IDictionary>] [-MethodName] <string> [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-CimClass] <cimclass> [[-Arguments] <IDictionary>] [-MethodName] <string> -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -Query <string> [-QueryDialect <string>] [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -Query <string> -CimSession <CimSession[]> [-QueryDialect <string>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Invoke-CimMethod [-ClassName] <string> [[-Arguments] <IDictionary>] [-MethodName] <string> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-ClassName] <string> [[-Arguments] <IDictionary>] [-MethodName] <string> -CimSession <CimSession[]> [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-InputObject] <ciminstance> [[-Arguments] <IDictionary>] [-MethodName] <string> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-InputObject] <ciminstance> [[-Arguments] <IDictionary>] [-MethodName] <string> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -ResourceUri <uri> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -ResourceUri <uri> -CimSession <CimSession[]> [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-CimClass] <cimclass> [[-Arguments] <IDictionary>] [-MethodName] <string> [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [-CimClass] <cimclass> [[-Arguments] <IDictionary>] [-MethodName] <string> -CimSession <CimSession[]> [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -Query <string> [-QueryDialect <string>] [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-CimMethod [[-Arguments] <IDictionary>] [-MethodName] <string> -Query <string> -CimSession <CimSession[]> [-QueryDialect <string>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Invoke a method

```powershell
$method = @{
  Query = 'select * from Win32_Process where name like "notepad%"'
  MethodName = "Terminate"
}
Invoke-CimMethod @method
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Invoke-CimMethod.md)


### Invoke-CommandInDesktopPackage

Version: Both

Module: Appx

Syntax:

```powershell
Invoke-CommandInDesktopPackage [-PackageFamilyName] <string> [[-AppId] <string>] [-Command] <string> [[-Args] <string>] [-PreventBreakaway] [<CommonParameters>]
```

Example: Invoke Notepad to read virtualized files

```powershell
$params = @{
    AppId             = 'ContosoApp'
    PackageFamilyName = 'Contoso.MyApp_abcdefgh23456'
    Command           = 'notepad.exe'
}
Invoke-CommandInDesktopPackage @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Invoke-CommandInDesktopPackage.md)


### Invoke-DscResource

Version: Both

Module: PSDesiredStateConfiguration

Syntax:

```powershell
Invoke-DscResource [-Name] <string> [-Method] <string> -ModuleName <ModuleSpecification> -Property <hashtable> [<CommonParameters>]
```

Example: Invoke the Set method of a resource by specifying its mandatory properties

```powershell
Invoke-DscResource -Name WindowsProcess -Method Set -ModuleName PSDesiredStateConfiguration -Property @{
    Path      = 'C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe'
    Arguments = ''
}
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/invoke-dscresource?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/invoke-dscresource?view=powershell-7.5)


### Invoke-LapsPolicyProcessing

Version: Both

Module: LAPS

Syntax:

```powershell
Invoke-LapsPolicyProcessing [<CommonParameters>]
```

Example: 

```powershell
Invoke-LapsPolicyProcessing
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Invoke-LapsPolicyProcessing.md)


### Invoke-TroubleshootingPack

Version: Both

Module: TroubleshootingPack

Syntax:

```powershell
Invoke-TroubleshootingPack [-Pack] <DiagPack> [-AnswerFile <string>] [-Result <string>] [-Unattended] [<CommonParameters>]
```

Example: Run a troubleshooting pack

```powershell
PS C:\> Get-TroubleshootingPack -Path "C:\Windows\Diagnostics\System\Audio" | Invoke-TroubleshootingPack
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TroubleshootingPack/Invoke-TroubleshootingPack.md)


### Invoke-WmiMethod

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Invoke-WmiMethod [-Class] <string> [-Name] <string> [[-ArgumentList] <Object[]>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> -InputObject <wmi> [-ArgumentList <Object[]>] [-AsJob] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> -Path <string> [-ArgumentList <Object[]>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): List the required order of WMI method parameters

```powershell
Get-WmiObject Win32_Volume |
    Get-Member -MemberType Method -Name Format |
    Select-Object -ExpandProperty Definition
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Invoke-WmiMethod.md)


### Invoke-WSManAction

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Invoke-WSManAction [-ResourceURI] <uri> [-Action] <string> [[-SelectorSet] <hashtable>] [-ConnectionURI <uri>] [-FilePath <string>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Invoke-WSManAction [-ResourceURI] <uri> [-Action] <string> [[-SelectorSet] <hashtable>] [-ApplicationName <string>] [-ComputerName <string>] [-FilePath <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Invoke a method

```powershell
$params = @{
    Action         = 'StartService'
    ResourceURI    = 'wmicimv2/Win32_Service'
    SelectorSet    = @{name = 'spooler'}
    Authentication = 'Default'
}
Invoke-WSManAction @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Invoke-WSManAction.md)


### Join-DtcDiagnosticResourceManager

Version: Both

Module: MsDtc

Syntax:

```powershell
Join-DtcDiagnosticResourceManager [-Transaction] <DtcDiagnosticTransaction> [[-ComputerName] <string>] [[-Port] <int>] [-Volatile] [<CommonParameters>]
```

Example: Enlist a new diagnostic transaction

```powershell
PS C:\> $Transaction = New-DtcDiagnosticTransaction
PS C:\> Join-DtcDiagnosticResourceManager -Transaction $Transaction
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Join-DtcDiagnosticResourceManager.md)


### Limit-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Limit-EventLog [-LogName] <string[]> [-ComputerName <string[]>] [-RetentionDays <int>] [-OverflowAction <OverflowAction>] [-MaximumSize <long>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Increase the size of an event log

```powershell
Limit-EventLog -LogName "Windows PowerShell" -MaximumSize 20KB
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Limit-EventLog.md)


### Merge-CIPolicy

Version: Both

Module: ConfigCI

Syntax:

```powershell
Merge-CIPolicy [-OutputFilePath] <string> [-PolicyPaths] <string[]> [-Rules <Rule[]>] [-AppIdTaggingPolicy] [<CommonParameters>]
```

Example: Merge policies

```powershell
PS C:\> Merge-CIPolicy -PolicyPaths '.\Policy.xml','.\Policy02.xml' -OutputFilePath '.\MergedPolicy.xml'

Name           : MSIT Test CodeSign CA 3
Id             : ID_SIGNER_S_17_0
TypeId         : Allow
Root           : FA6B9A2230CE08BCA81D096B28CF495672401D3A43A0D285CF352464A6C9C7FD
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : VeriSign Class 3 Code Signing 2010 CA
Id             : ID_SIGNER_S_1D_0
TypeId         : Allow
Root           : 4843A82ED3B1F2BFBEE9671960E1940C942F688D
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : Microsoft Windows Third Party Component CA 2012
Id             : ID_SIGNER_S_1E_0
TypeId         : Allow
Root           : CEC1AFD0E310C55C1DCC601AB8E172917706AA32FB5EAF826813547FDF02DD46
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Sha1
Id             : ID_ALLOW_A_49_1
TypeId         : Allow
Root           :
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Merge-CIPolicy.md)


### Mount-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Mount-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [<CommonParameters>]
Mount-AppvClientConnectionGroup [-Name] <string> [<CommonParameters>]
Mount-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [<CommonParameters>]
```

Example: Download packages for a named group

```powershell
PS C:\> Mount-AppvClientConnectionGroup -Name "MyGroup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Mount-AppvClientConnectionGroup.md)


### Mount-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Mount-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Cancel] [<CommonParameters>]
Mount-AppvClientPackage [-Package] <AppvClientPackage> [-Cancel] [<CommonParameters>]
Mount-AppvClientPackage [-Name] <string> [[-Version] <string>] [<CommonParameters>]
```

Example: Get a specific version of a package

```powershell
PS C:\> Mount-AppvClientPackage -Name "MyApp" -Version 2
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Mount-AppvClientPackage.md)


### Mount-AppxVolume

Version: Both

Module: Appx

Syntax:

```powershell
Mount-AppxVolume [-Volume] <AppxVolume[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Mount a volume by using a path

```powershell
Mount-AppxVolume -Volume E:\
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Mount-AppxVolume.md)


### Mount-WindowsImage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Mount-WindowsImage -Path <string> -ImagePath <string> -Index <uint32> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> -Name <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -Remount [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Mount-WindowsImage -Path <string> -ImagePath <string> -Index <uint> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> -Name <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -Remount [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Mount an image from the install.vhd file to a directory

```powershell
PS C:\> Mount-WindowsImage -ImagePath "c:\imagestore\install.vhd" -Index 1 -Path "c:\offline"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Mount-WindowsImage.md)


### Move-AppxPackage

Version: Both

Module: Appx

Syntax:

```powershell
Move-AppxPackage [-Package] <string[]> [-Volume] <AppxVolume> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Move a package to a volume specified by a path

```powershell
Move-AppxPackage -Package "package1_1.0.0.0_neutral__8wekyb3d8bbwe" -Volume F:\
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Move-AppxPackage.md)


### New-AppLockerPolicy

Version: Both

Module: AppLocker

Syntax:

```powershell
New-AppLockerPolicy [-FileInformation] <List[FileInformation]> [-AllowWindows] [-RuleType <List[RuleType]>] [-RuleNamePrefix <string>] [-User <string>] [-Optimize] [-IgnoreMissingFileInformation] [-Xml] [-ServiceEnforcement <ServiceEnforcementMode>] [<CommonParameters>]
New-AppLockerPolicy -AllowWindows [-RuleType <List[RuleType]>] [-RuleNamePrefix <string>] [-User <string>] [-Optimize] [-IgnoreMissingFileInformation] [-Xml] [-ServiceEnforcement <ServiceEnforcementMode>] [<CommonParameters>]
```

Example: Create an AppLocker policy with allow rules

```powershell
C:\PS>Get-ChildItem C:\Windows\System32\*.exe | Get-AppLockerFileInformation | New-AppLockerPolicy -RuleType Publisher, Hash -User Everyone -RuleNamePrefix System32

                                Version RuleCollections                         RuleCollectionTypes
                                ------- ---------------                         -------------------
                                      1 {Microsoft.Security.ApplicationId.Po... {Exe}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppLocker/New-AppLockerPolicy.md)


### New-BcdEntry

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
New-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Application <string> [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
New-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Device [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
New-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Inherit <string> [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
New-BcdEntry [-Id] <string> [[-Store] <BcdStoreInfo>] [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/New-BcdEntry.md)


### New-BcdStore

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
New-BcdStore [-Path] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/New-BcdStore.md)


### New-CertificateNotificationTask

Version: Both

Module: PKI

Syntax:

```powershell
New-CertificateNotificationTask -Type <CertificateNotificationType> -PSScript <string> -Name <string> -Channel <NotificationChannel> [-RunTaskForExistingCertificates] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    PSScript = 'C:\myscript.ps1'
    Channel = 'System'
    Type = 'Replace'
    Name = 'My System Certificate Task'
}
New-CertificateNotificationTask @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/New-CertificateNotificationTask.md)


### New-CimInstance

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create an instance of a CIM class

```powershell
$prop = @{
    Name = "testvar"
    VariableValue = "testvalue"
    UserName = "domain\user"
}
New-CimInstance -ClassName Win32_Environment -Property $prop
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/New-CimInstance.md)


### New-CimSession

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
New-CimSession [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-Authentication <PasswordAuthenticationMechanism>] [-Name <string>] [-OperationTimeoutSec <uint32>] [-SkipTestConnection] [-Port <uint32>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
New-CimSession [[-ComputerName] <string[]>] [-CertificateThumbprint <string>] [-Name <string>] [-OperationTimeoutSec <uint32>] [-SkipTestConnection] [-Port <uint32>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
```

Syntax (7):

```powershell
New-CimSession [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-Authentication <PasswordAuthenticationMechanism>] [-Name <string>] [-OperationTimeoutSec <uint>] [-SkipTestConnection] [-Port <uint>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
New-CimSession [[-ComputerName] <string[]>] [-CertificateThumbprint <string>] [-Name <string>] [-OperationTimeoutSec <uint>] [-SkipTestConnection] [-Port <uint>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
```

Example: Create a CIM session with default options

```powershell
New-CimSession
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/New-CimSession.md)


### New-CimSessionOption

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
New-CimSessionOption [-Protocol] <ProtocolType> [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-NoEncryption] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-EncodePortInServicePrincipalName] [-Encoding <PacketEncoding>] [-HttpPrefix <uri>] [-MaxEnvelopeSizeKB <uint32>] [-ProxyAuthentication <PasswordAuthenticationMechanism>] [-ProxyCertificateThumbprint <string>] [-ProxyCredential <pscredential>] [-ProxyType <ProxyType>] [-UseSsl] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-Impersonation <ImpersonationType>] [-PacketIntegrity] [-PacketPrivacy] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
```

Syntax (7):

```powershell
New-CimSessionOption [-Protocol] <ProtocolType> [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-NoEncryption] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-EncodePortInServicePrincipalName] [-Encoding <PacketEncoding>] [-HttpPrefix <uri>] [-MaxEnvelopeSizeKB <uint>] [-ProxyAuthentication <PasswordAuthenticationMechanism>] [-ProxyCertificateThumbprint <string>] [-ProxyCredential <pscredential>] [-ProxyType <ProxyType>] [-UseSsl] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-Impersonation <ImpersonationType>] [-PacketIntegrity] [-PacketPrivacy] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
```

Example: Create a CIM session options object for DCOM

```powershell
$so = New-CimSessionOption -Protocol Dcom
New-CimSession -ComputerName Server01 -SessionOption $so
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/New-CimSessionOption.md)


### New-CIPolicy

Version: Both

Module: ConfigCI

Syntax:

```powershell
New-CIPolicy [-FilePath] <string> -Level <RuleLevel> [-DriverFiles <DriverFile[]>] [-Fallback <RuleLevel[]>] [-Audit] [-ScanPath <string>] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-UserPEs] [-NoScript] [-Deny] [-NoShadowCopy] [-MultiplePolicyFormat] [-OmitPaths <string[]>] [-PathToCatroot <string>] [-AppIdTaggingPolicy] [-AppIdTaggingKey <string[]>] [-AppIdTaggingValue <string[]>] [<CommonParameters>]
New-CIPolicy [-FilePath] <string> -Rules <Rule[]> [-Audit] [-ScanPath <string>] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-UserPEs] [-NoScript] [-Deny] [-NoShadowCopy] [-MultiplePolicyFormat] [-OmitPaths <string[]>] [-PathToCatroot <string>] [-AppIdTaggingPolicy] [-AppIdTaggingKey <string[]>] [-AppIdTaggingValue <string[]>] [<CommonParameters>]
```

Example: Create a policy in multiple policy format

```powershell
PS C:\> New-CIPolicy -ScanPath '.\temp\' -UserPEs -OmitPaths '.\temp\ConfigCITestBinaries' -NoScript -FilePath '.\Policy.xml' -Level Publisher -MultiplePolicyFormat
Scan completed successfully

The second command displays the contents of the policy.
PS C:\> Get-Content -Path '.\policy.xml'
<?xml version="1.0" encoding="utf-8"?>
<SiPolicy xmlns="urn:schemas-microsoft-com:sipolicy" PolicyType="Base Policy">
  <VersionEx>10.0.0.0</VersionEx>
  <BasePolicyID>{BB9EC112-DD85-41AD-9778-22680D3D8A22}</BasePolicyID>
  <PolicyID>{BB9EC112-DD85-41AD-9778-22680D3D8A22}</PolicyID>
  <PlatformID>{2E07F7E4-194C-4D20-B7C9-6F44A6C5A234}</PlatformID>
  <Rules>
    <Rule>
      <Option>Enabled:Unsigned System Integrity Policy</Option>
    </Rule>
    <Rule>
      <Option>Enabled:Audit Mode</Option>
    </Rule>
    <Rule>
      <Option>Enabled:Advanced Boot Options Menu</Option>
    </Rule>
    <Rule>
      <Option>Enabled:UMCI</Option>
    </Rule>
    <Rule>
      <Option>Disabled:Script Enforcement</Option>
    </Rule>
  </Rules>
  <!--EKUS-->
  <EKUs />
  <!--File Rules-->
  <FileRules>
    <Allow ID="ID_ALLOW_A_2F" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Sha1" Hash="BE0777
F5AF88628D4555A875036648DF1AD19BBE" />
    <Allow ID="ID_ALLOW_A_30" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Sha256" Hash="6FA5
AF724499C338A77FEEAD90F55DDF5F23D081C6DCE8E9DF486E95C6A9B310" />
    <Allow ID="ID_ALLOW_A_31" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Page Sha1" Hash="D
41570F2E6E7E6245CF342131D4706C944562B1E" />
    <Allow ID="ID_ALLOW_A_32" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Page Sha256" Hash=
"F714D9784E15B88F56180C8EE2B40C769CC83428954585A1DCF9A260FE967CDD" />
    <Allow ID="ID_ALLOW_A_37" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Sha1" Ha
sh="FD58E1BFA1E661C809F8A2437932B0F0308A99F8" />
    <Allow ID="ID_ALLOW_A_38" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Sha256"
Hash="A1C9FA473C2D79D0F68DF6EC72E31847F0FDA283D3A9E6B1405B0DF5929CCCB8" />
    <Allow ID="ID_ALLOW_A_39" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Page Sha
1" Hash="6D3764B75C6502634234911B8F224FC9568217C9" />
    <Allow ID="ID_ALLOW_A_3A" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Page Sha
256" Hash="2196AD3A00A59F4C35EEF7FE843FA3D6F80D5EFB3C674ADC087396B77AB35768" />
    <Allow ID="ID_ALLOW_A_3F" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Sha1" Ha
sh="28FAEFE1B18A979F9FF55744B22C6E5EA2949959" />
    <Allow ID="ID_ALLOW_A_40" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Sha256"
Hash="DA737C142A51A73D82E6AD677474C8031486FDEF018A6FE9D178564F83AB284B" />
    <Allow ID="ID_ALLOW_A_41" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Page Sha
1" Hash="029606A9B560F4921EC1122AA73C19C9B97F7764" />
    <Allow ID="ID_ALLOW_A_42" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Page Sha
256" Hash="2A5D6BCBFA55DB0F0487F45F4A6986AFC2C4783820EDA48DE9E0560E51D8DB56" />
    <Allow ID="ID_ALLOW_A_33" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Sha1" Hash="BE0777F5AF88628D4555A875036648DF1AD19BBE" />
    <Allow ID="ID_ALLOW_A_34" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Sha256" Hash="6FA5
AF724499C338A77FEEAD90F55DDF5F23D081C6DCE8E9DF486E95C6A9B310" />
    <Allow ID="ID_ALLOW_A_35" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Page Sha1" Hash="D
41570F2E6E7E6245CF342131D4706C944562B1E" />
    <Allow ID="ID_ALLOW_A_36" FriendlyName="\\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Page Sha256" Hash=
"F714D9784E15B88F56180C8EE2B40C769CC83428954585A1DCF9A260FE967CDD" />
    <Allow ID="ID_ALLOW_A_3B" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Sha1" Hash="FD58E1BFA1E661C809F8A2437932B0F0308A99F8" />
    <Allow ID="ID_ALLOW_A_3C" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Sha256"
Hash="A1C9FA473C2D79D0F68DF6EC72E31847F0FDA283D3A9E6B1405B0DF5929CCCB8" />
    <Allow ID="ID_ALLOW_A_3D" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Page Sha
1" Hash="6D3764B75C6502634234911B8F224FC9568217C9" />
    <Allow ID="ID_ALLOW_A_3E" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\ntoskrnl.exe Hash Page Sha
256" Hash="2196AD3A00A59F4C35EEF7FE843FA3D6F80D5EFB3C674ADC087396B77AB35768" />
    <Allow ID="ID_ALLOW_A_43" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Sha1" Ha
sh="28FAEFE1B18A979F9FF55744B22C6E5EA2949959" />
    <Allow ID="ID_ALLOW_A_44" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Sha256"
Hash="DA737C142A51A73D82E6AD677474C8031486FDEF018A6FE9D178564F83AB284B" />
    <Allow ID="ID_ALLOW_A_45" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Page Sha
1" Hash="029606A9B560F4921EC1122AA73C19C9B97F7764" />
    <Allow ID="ID_ALLOW_A_46" FriendlyName="\\?\E:\cmdlets\temp\PackageInspectorTestBinaries\storahci.sys Hash Page Sha
256" Hash="2A5D6BCBFA55DB0F0487F45F4A6986AFC2C4783820EDA48DE9E0560E51D8DB56" />
  </FileRules>
  <!--Signers-->
  <Signers>
    <Signer ID="ID_SIGNER_S_D" Name="MSIT Test CodeSign CA 3">
      <CertRoot Type="TBS" Value="FA6B9A2230CE08BCA81D096B28CF495672401D3A43A0D285CF352464A6C9C7FD" />
      <CertPublisher Value="Microsoft Windows" />
    </Signer>
    <Signer ID="ID_SIGNER_S_E" Name="MSIT Test CodeSign CA 3">
      <CertRoot Type="TBS" Value="FA6B9A2230CE08BCA81D096B28CF495672401D3A43A0D285CF352464A6C9C7FD" />
      <CertPublisher Value="Microsoft Windows" />
    </Signer>
    <Signer ID="ID_SIGNER_S_13" Name="VeriSign Class 3 Code Signing 2010 CA">
      <CertRoot Type="TBS" Value="4843A82ED3B1F2BFBEE9671960E1940C942F688D" />
      <CertPublisher Value="NVIDIA Corporation" />
    </Signer>
    <Signer ID="ID_SIGNER_S_14" Name="Microsoft Windows Third Party Component CA 2012">
      <CertRoot Type="TBS" Value="CEC1AFD0E310C55C1DCC601AB8E172917706AA32FB5EAF826813547FDF02DD46" />
      <CertPublisher Value="Microsoft Windows Hardware Compatibility Publisher" />
    </Signer>
    <Signer ID="ID_SIGNER_S_15" Name="VeriSign Class 3 Code Signing 2010 CA">
      <CertRoot Type="TBS" Value="4843A82ED3B1F2BFBEE9671960E1940C942F688D" />
      <CertPublisher Value="NVIDIA Corporation" />
    </Signer>
    <Signer ID="ID_SIGNER_S_16" Name="Microsoft Windows Third Party Component CA 2012">
      <CertRoot Type="TBS" Value="CEC1AFD0E310C55C1DCC601AB8E172917706AA32FB5EAF826813547FDF02DD46" />
      <CertPublisher Value="Microsoft Windows Hardware Compatibility Publisher" />
    </Signer>
  </Signers>
  <!--Driver Signing Scenarios-->
  <SigningScenarios>
    <SigningScenario Value="131" ID="ID_SIGNINGSCENARIO_DRIVERS_1" FriendlyName="Auto generated policy on 11-13-2015">
      <ProductSigners>
        <FileRulesRef>
          <FileRuleRef RuleID="ID_ALLOW_A_2F" />
          <FileRuleRef RuleID="ID_ALLOW_A_30" />
          <FileRuleRef RuleID="ID_ALLOW_A_31" />
          <FileRuleRef RuleID="ID_ALLOW_A_32" />
          <FileRuleRef RuleID="ID_ALLOW_A_37" />
          <FileRuleRef RuleID="ID_ALLOW_A_38" />
          <FileRuleRef RuleID="ID_ALLOW_A_39" />
          <FileRuleRef RuleID="ID_ALLOW_A_3A" />
          <FileRuleRef RuleID="ID_ALLOW_A_3F" />
          <FileRuleRef RuleID="ID_ALLOW_A_40" />
          <FileRuleRef RuleID="ID_ALLOW_A_41" />
          <FileRuleRef RuleID="ID_ALLOW_A_42" />
        </FileRulesRef>
        <AllowedSigners>
          <AllowedSigner SignerId="ID_SIGNER_S_D" />
          <AllowedSigner SignerId="ID_SIGNER_S_13" />
          <AllowedSigner SignerId="ID_SIGNER_S_14" />
        </AllowedSigners>
      </ProductSigners>
    </SigningScenario>
    <SigningScenario Value="12" ID="ID_SIGNINGSCENARIO_WINDOWS" FriendlyName="Auto generated policy on 11-13-2015">
      <ProductSigners>
        <FileRulesRef>
          <FileRuleRef RuleID="ID_ALLOW_A_33" />
          <FileRuleRef RuleID="ID_ALLOW_A_34" />
          <FileRuleRef RuleID="ID_ALLOW_A_35" />
          <FileRuleRef RuleID="ID_ALLOW_A_36" />
          <FileRuleRef RuleID="ID_ALLOW_A_3B" />
          <FileRuleRef RuleID="ID_ALLOW_A_3C" />
          <FileRuleRef RuleID="ID_ALLOW_A_3D" />
          <FileRuleRef RuleID="ID_ALLOW_A_3E" />
          <FileRuleRef RuleID="ID_ALLOW_A_43" />
          <FileRuleRef RuleID="ID_ALLOW_A_44" />
          <FileRuleRef RuleID="ID_ALLOW_A_45" />
          <FileRuleRef RuleID="ID_ALLOW_A_46" />
        </FileRulesRef>
        <AllowedSigners>
          <AllowedSigner SignerId="ID_SIGNER_S_E" />
          <AllowedSigner SignerId="ID_SIGNER_S_15" />
          <AllowedSigner SignerId="ID_SIGNER_S_16" />
        </AllowedSigners>
      </ProductSigners>
    </SigningScenario>
  </SigningScenarios>
  <UpdatePolicySigners />
  <CiSigners>
    <CiSigner SignerId="ID_SIGNER_S_E" />
    <CiSigner SignerId="ID_SIGNER_S_15" />
    <CiSigner SignerId="ID_SIGNER_S_16" />
  </CiSigners>
  <HvciOptions>0</HvciOptions>
</SiPolicy>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/New-CIPolicy.md)


### New-CIPolicyRule

Version: Both

Module: ConfigCI

Syntax:

```powershell
New-CIPolicyRule -Level <RuleLevel> [-DriverFiles <DriverFile[]>] [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [<CommonParameters>]
New-CIPolicyRule -DriverFilePath <string[]> -Level <RuleLevel> [-AppID <string>] [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [<CommonParameters>]
New-CIPolicyRule [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-Package <AppxPackage>] [<CommonParameters>]
New-CIPolicyRule [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-FilePathRule <string>] [<CommonParameters>]
```

Example: Create policy rules for drivers

```powershell
PS C:\> $DriverFiles = Get-SystemDriver -ScanPath '.\temp\' -UserPEs -OmitPaths '.\temp\ConfigCITestBinaries' -NoScript
PS C:\> New-CIPolicyRule -Level FileName -DriverFiles $DriverFiles
Scan completed successfully


Name           : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll FileRule
Id             : ID_ALLOW_A_1
TypeId         : Allow
Root           :
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.Tests.dll FileRule
Id             : ID_ALLOW_A_3
TypeId         : Allow
Root           :
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False

Name           : \\?\E:\cmdlets\temp\Microsoft.PackageInspector.Tests.dll FileRule
Id             : ID_ALLOW_A_5
TypeId         : Allow
Root           :
FileVersionRef :
Wellknown      : False
Ekus           :
Exceptions     :
FileAttributes :
FileException  : False
UserMode       : False
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/New-CIPolicyRule.md)


### New-DtcDiagnosticTransaction

Version: Both

Module: MsDtc

Syntax:

```powershell
New-DtcDiagnosticTransaction [[-Timeout] <int>] [[-IsolationLevel] <IsolationLevel>] [<CommonParameters>]
```

Example: Create a diagnostic transaction

```powershell
PS C:\> New-DtcDiagnosticTransaction -Timeout 60 -IsolationLevel Serializable
Id
--
4625a5a3-af35-465d-a331-f224d79e4c85
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/New-DtcDiagnosticTransaction.md)


### New-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
New-EventLog [-LogName] <string> [-Source] <string[]> [[-ComputerName] <string[]>] [-CategoryResourceFile <string>] [-MessageResourceFile <string>] [-ParameterResourceFile <string>] [<CommonParameters>]
```

Example (5.1): create a new event log

```powershell
New-EventLog -Source TestApp -LogName TestLog -MessageResourceFile C:\Test\TestApp.dll
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/New-EventLog.md)


### New-FileCatalog

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
New-FileCatalog [-CatalogFilePath] <string> [[-Path] <string[]>] [-CatalogVersion <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a file catalog for `Microsoft.PowerShell.Utility`

```powershell
$newFileCatalogSplat = @{
    Path = "$PSHOME\Modules\Microsoft.PowerShell.Utility"
    CatalogFilePath = '\temp\Microsoft.PowerShell.Utility.cat'
    CatalogVersion = 2.0
}
New-FileCatalog @newFileCatalogSplat
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/New-FileCatalog.md)


### New-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
New-JobTrigger -Once -At <datetime> [-RandomDelay <timespan>] [-RepetitionInterval <timespan>] [-RepetitionDuration <timespan>] [-RepeatIndefinitely] [<CommonParameters>]
New-JobTrigger -Daily -At <datetime> [-DaysInterval <int>] [-RandomDelay <timespan>] [<CommonParameters>]
New-JobTrigger -Weekly -At <datetime> -DaysOfWeek <DayOfWeek[]> [-WeeksInterval <int>] [-RandomDelay <timespan>] [<CommonParameters>]
New-JobTrigger -AtStartup [-RandomDelay <timespan>] [<CommonParameters>]
New-JobTrigger -AtLogOn [-RandomDelay <timespan>] [-User <string>] [<CommonParameters>]
```

Example (5.1): Once Schedule

```powershell
New-JobTrigger -Once -At "1/20/2012 3:00 AM"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/New-JobTrigger.md)


### New-LocalGroup

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
New-LocalGroup [-Name] <string> [-Description <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Create a security group

```powershell
New-LocalGroup -Name "SecurityGroup04"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/New-LocalGroup.md)


### New-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
New-LocalUser [-Name] <string> -Password <securestring> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-Disabled] [-FullName <string>] [-PasswordNeverExpires] [-UserMayNotChangePassword] [-WhatIf] [-Confirm] [<CommonParameters>]
New-LocalUser [-Name] <string> -NoPassword [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-Disabled] [-FullName <string>] [-UserMayNotChangePassword] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Create a user account

```powershell
New-LocalUser -Name 'User02' -Description 'Description of this account.' -NoPassword
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/New-LocalUser.md)


### New-NetIPsecAuthProposal

Version: Both

Module: NetSecurity

Syntax:

```powershell
New-NetIPsecAuthProposal -Machine [-Health] -Cert -Authority <string> [-AccountMapping] [-AuthorityType <CertificateAuthorityType>] [-ExtendedKeyUsage <string[]>] [-ExcludeCAName] [-FollowRenewal] [-SelectionCriteria] [-Signing <CertificateSigningAlgorithm>] [-SubjectName <string>] [-SubjectNameType <CertificateSubjectType>] [-Thumbprint <string>] [-ValidationCriteria] [<CommonParameters>]
New-NetIPsecAuthProposal -User -Cert -Authority <string> [-AccountMapping] [-AuthorityType <CertificateAuthorityType>] [-ExtendedKeyUsage <string[]>] [-ExcludeCAName] [-FollowRenewal] [-SelectionCriteria] [-Signing <CertificateSigningAlgorithm>] [-SubjectName <string>] [-SubjectNameType <CertificateSubjectType>] [-Thumbprint <string>] [-ValidationCriteria] [<CommonParameters>]
New-NetIPsecAuthProposal -Anonymous [<CommonParameters>]
New-NetIPsecAuthProposal -Machine -Kerberos [-Proxy <string>] [<CommonParameters>]
New-NetIPsecAuthProposal -User -Kerberos [-Proxy <string>] [<CommonParameters>]
New-NetIPsecAuthProposal -Machine -Ntlm [<CommonParameters>]
New-NetIPsecAuthProposal -Machine [-PreSharedKey] <string> [<CommonParameters>]
New-NetIPsecAuthProposal -User -Ntlm [<CommonParameters>]
```

Example: 

```powershell
PS C:\>$cert1Proposal = New-NetIPsecAuthProposal -Machine -Cert -Authority "C=US,O=MSFT,CN=ꞌMicrosoft Root Authorityꞌ" -AuthorityType Root



PS C:\>$cert2Proposal = New-NetIPsecAuthProposal -Machine -Cert -Authority "C=US,O=MYORG,CN='My Organizations Root Certificate'" -AuthorityType Root



PS C:\>$certAuthSet = New-NetIPsecPhase1AuthSet -DisplayName "Computer Certificate Auth Set" -Proposal $cert1Proposal,$cert2Proposal



PS C:\>New-NetIPSecRule -DisplayName "Authenticate with Certificates Rule" -InboundSecurity Require -OutboundSecurity Request -Phase2AuthSet $certAuthSet.Name
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/NetSecurity/New-NetIPsecAuthProposal.md)


### New-NetIPsecMainModeCryptoProposal

Version: Both

Module: NetSecurity

Syntax:

```powershell
New-NetIPsecMainModeCryptoProposal [-Encryption <EncryptionAlgorithm>] [-KeyExchange <DiffieHellmanGroup>] [-Hash <HashAlgorithm>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\>$proposal1 = (New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH1)



PS C:\>$proposal2 = (New-NetIPsecMainModeCryptoProposal -Encryption AES192 -Hash MD5 -KeyExchange DH14)



PS C:\>$proposal3 = (New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH19)



PS C:\>$mMCryptoSet= (New-NetIPsecMainModeCryptoSet -DisplayName "Main Mode Crypto Set" -Proposal $proposal1,$proposal2,$proposal3)


This cmdlet shows an alternative method of accomplishing the previous steps.
PS C:\>$mMCryptoSet = New-NetIPsecMainModeCryptoSet -DisplayName "Main Mode Crypto Set" -Proposal (New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH1),(New-NetIPsecMainModeCryptoProposal -Encryption AES192 -Hash MD5 -KeyExchange DH14),(New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH19)



PS C:\>New-NetIPsecMainModeRule -DisplayName "Main Mode Rule" -MainModeCryptoSet $mMCryptoSet.Name
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/NetSecurity/New-NetIPsecMainModeCryptoProposal.md)


### New-NetIPsecQuickModeCryptoProposal

Version: Both

Module: NetSecurity

Syntax (5.1):

```powershell
New-NetIPsecQuickModeCryptoProposal [-Encryption <EncryptionAlgorithm>] [-AHHash <HashAlgorithm>] [-ESPHash <HashAlgorithm>] [-MaxKiloBytes <uint64>] [-MaxMinutes <uint64>] [-Encapsulation <IPsecEncapsulation>] [<CommonParameters>]
```

Syntax (7):

```powershell
New-NetIPsecQuickModeCryptoProposal [-Encryption <EncryptionAlgorithm>] [-AHHash <HashAlgorithm>] [-ESPHash <HashAlgorithm>] [-MaxKiloBytes <ulong>] [-MaxMinutes <ulong>] [-Encapsulation <IPsecEncapsulation>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\>$QMProposal = New-NetIPsecQuickModeCryptoProposal -Encapsulation ESP -ESPHash SHA1 -Encryption AES128



PS C:\>$QMCryptoSet = New-NetIPsecQuickModeCryptoSet -DisplayName "esp:sha1-des3" -Proposal $QMProposal



PS C:\>New-NetIPSecRule -DisplayName "Tunnel from HQ to Dallas Branch" -Mode Tunnel -LocalAddress 192.168.0.0/16 -RemoteAddress 192.157.0.0/16 -LocalTunnelEndpoint 1.1.1.1 -RemoteTunnelEndpoint 2.2.2.2 -InboundSecurity Require -OutboundSecurity Require -QuickModeCryptoSet $QMCryptoSet.Name
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/NetSecurity/New-NetIPsecQuickModeCryptoProposal.md)


### New-PmemDedicatedMemory

Version: Both

Module: PersistentMemory

Syntax:

```powershell
New-PmemDedicatedMemory -RegionId <uint32[]> [-FriendlyName <string[]>] [-SizeInBytes <uint64[]>] [<CommonParameters>]
```

Example: Create dedicated persistent memory

```powershell
New-PmemDedicatedMemory -RegionId 1 -SizeInBytes 270582939648
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/New-PmemDedicatedMemory.md)


### New-PmemDisk

Version: Both

Module: PersistentMemory

Syntax:

```powershell
New-PmemDisk -RegionId <uint32[]> [-FriendlyName <string[]>] [-DiskSizeInBytes <uint64[]>] [-AtomicityType <NAMESPACE_ATOMICITY_TYPE[]>] [<CommonParameters>]
New-PmemDisk -DiskSizeInBytes <uint64[]> -Simulated [-AtomicityType <NAMESPACE_ATOMICITY_TYPE[]>] [<CommonParameters>]
```

Example: Create a disk

```powershell
New-PmemDisk -RegionId 1 -AtomicityType BlockTranslationTable
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/New-PmemDisk.md)


### New-ProvisioningRepro

Version: Both

Module: Provisioning

Syntax: none

Example: none

Source: [Provisioning module documentation](https://learn.microsoft.com/en-us/powershell/module/provisioning) (no dedicated page).


### New-PSWorkflowExecutionOption

Version: Both

Module: PSWorkflow

Syntax:

```powershell
New-PSWorkflowExecutionOption [-PersistencePath <string>] [-MaxPersistenceStoreSizeGB <long>] [-PersistWithEncryption] [-MaxRunningWorkflows <int>] [-AllowedActivity <string[]>] [-OutOfProcessActivity <string[]>] [-EnableValidation] [-MaxDisconnectedSessions <int>] [-MaxConnectedSessions <int>] [-MaxSessionsPerWorkflow <int>] [-MaxSessionsPerRemoteNode <int>] [-MaxActivityProcesses <int>] [-ActivityProcessIdleTimeoutSec <int>] [-RemoteNodeSessionIdleTimeoutSec <int>] [-SessionThrottleLimit <int>] [-WorkflowShutdownTimeoutMSec <int>] [<CommonParameters>]
```

Example (5.1): Create a Workflow Options Object

```powershell
New-PSWorkflowExecutionOption -MaxSessionsPerWorkflow 10 -MaxDisconnectedSessions 200
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSWorkflow/New-PSWorkflowExecutionOption.md)


### New-ScheduledJobOption

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
New-ScheduledJobOption [-RunElevated] [-HideInTaskScheduler] [-RestartOnIdleResume] [-MultipleInstancePolicy <TaskMultipleInstancePolicy>] [-DoNotAllowDemandStart] [-RequireNetwork] [-StopIfGoingOffIdle] [-WakeToRun] [-ContinueIfGoingOnBattery] [-StartIfOnBattery] [-IdleTimeout <timespan>] [-IdleDuration <timespan>] [-StartIfIdle] [<CommonParameters>]
```

Example (5.1): Create a scheduled job option object with default values

```powershell
New-ScheduledJobOption
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/New-ScheduledJobOption.md)


### New-SelfSignedCertificate

Version: Both

Module: PKI

Syntax:

```powershell
New-SelfSignedCertificate [-SecurityDescriptor <FileSecurity>] [-TextExtension <string[]>] [-Extension <X509Extension[]>] [-HardwareKeyUsage <HardwareKeyUsage[]>] [-KeyUsageProperty <KeyUsageProperty[]>] [-KeyUsage <KeyUsage[]>] [-KeyProtection <KeyProtection[]>] [-KeyExportPolicy <KeyExportPolicy[]>] [-KeyLength <int>] [-KeyAlgorithm <string>] [-SmimeCapabilities] [-ExistingKey] [-KeyLocation <string>] [-SignerReader <string>] [-Reader <string>] [-SignerPin <securestring>] [-Pin <securestring>] [-KeyDescription <string>] [-KeyFriendlyName <string>] [-Container <string>] [-Provider <string>] [-CurveExport <CurveParametersExportType>] [-KeySpec <KeySpec>] [-Type <CertificateType>] [-FriendlyName <string>] [-NotAfter <datetime>] [-NotBefore <datetime>] [-SerialNumber <string>] [-Subject <string>] [-DnsName <string[]>] [-SuppressOid <string[]>] [-HashAlgorithm <string>] [-AlternateSignatureAlgorithm] [-TestRoot] [-Signer <Certificate>] [-CloneCert <Certificate>] [-CertStoreLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    DnsName = 'www.fabrikam.com', 'www.contoso.com'
    CertStoreLocation = 'Cert:\LocalMachine\My'
}
New-SelfSignedCertificate @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/New-SelfSignedCertificate.md)


### New-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
New-Service [-Name] <string> [-BinaryPathName] <string> [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartMode>] [-Credential <pscredential>] [-DependsOn <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
New-Service [-Name] <string> [-BinaryPathName] <string> [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartupType>] [-Credential <pscredential>] [-SecurityDescriptorSddl <string>] [-DependsOn <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a service

```powershell
New-Service -Name "TestService" -BinaryPathName 'C:\WINDOWS\System32\svchost.exe -k netsvcs'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/New-Service.md)


### New-TlsSessionTicketKey

Version: Both

Module: TLS

Syntax:

```powershell
New-TlsSessionTicketKey [-Password] <securestring> [[-Path] <string>] [<CommonParameters>]
```

Example: Create a TLS session ticket key

```powershell
$Password = Read-Host -AsSecureString
New-TlsSessionTicketKey -Password $Password -Path 'C:\KeyConfig\TlsSessionTicketKey.config'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TLS/New-TlsSessionTicketKey.md)


### New-WebServiceProxy

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
New-WebServiceProxy [-Uri] <uri> [[-Class] <string>] [[-Namespace] <string>] [<CommonParameters>]
New-WebServiceProxy [-Uri] <uri> [[-Class] <string>] [[-Namespace] <string>] [-Credential <pscredential>] [<CommonParameters>]
New-WebServiceProxy [-Uri] <uri> [[-Class] <string>] [[-Namespace] <string>] [-UseDefaultCredential] [<CommonParameters>]
```

Example (5.1): Create a proxy for a Web service

```powershell
$calc = New-WebServiceProxy -Uri "http://www.dneonline.com/calculator.asmx?wsdl"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/New-WebServiceProxy.md)


### New-WindowsCustomImage

Version: Both

Module: Dism

Syntax:

```powershell
New-WindowsCustomImage -CapturePath <string> [-ConfigFilePath <string>] [-CheckIntegrity] [-Verify] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Capture image customization files

```powershell
PS C:\> New-WindowsCustomImage -CapturePath "c:\"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/New-WindowsCustomImage.md)


### New-WindowsImage

Version: Both

Module: Dism

Syntax:

```powershell
New-WindowsImage -ImagePath <string> -CapturePath <string> [-CompressionType <string>] [-ConfigFilePath <string>] [-Description <string>] [-Name <string>] [-CheckIntegrity] [-NoRpFix] [-Setbootable] [-Verify] [-WIMBoot] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Capture an image of a drive for a WIM file

```powershell
PS C:\> New-WindowsImage -ImagePath "c:\imagestore\custom.wim" -CapturePath "d:\" -Name "Drive D"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/New-WindowsImage.md)


### New-WinEvent

Version: Both

Module: Microsoft.PowerShell.Diagnostics

Syntax:

```powershell
New-WinEvent [-ProviderName] <string> [-Id] <int> [[-Payload] <Object[]>] [-Version <byte>] [<CommonParameters>]
```

Example: Create a new event

```powershell
New-WinEvent -ProviderName Microsoft-Windows-PowerShell -Id 45090 -Payload @("Workflow", "Running")
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Diagnostics/New-WinEvent.md)


### New-WinUserLanguageList

Version: Both

Module: International

Syntax:

```powershell
New-WinUserLanguageList [-Language] <string> [<CommonParameters>]
```

Example: Create and set a language list

```powershell
PS C:\> $UserLanguageList = New-WinUserLanguageList -Language "en-US"
PS C:\> $UserLanguageList.Add("fr-FR")
PS C:\> Set-WinUserLanguageList -LanguageList $UserLanguageList
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/New-WinUserLanguageList.md)


### New-WSManInstance

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
New-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ApplicationName <string>] [-ComputerName <string>] [-FilePath <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
New-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ConnectionURI <uri>] [-FilePath <string>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Create a HTTPS listener

```powershell
New-WSManInstance winrm/config/Listener -SelectorSet @{Transport='HTTPS'; Address='*'} -ValueSet @{Hostname="HOST";CertificateThumbprint="XXXXXXXXXX"}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/New-WSManInstance.md)


### New-WSManSessionOption

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
New-WSManSessionOption [-ProxyAccessType <ProxyAccessType>] [-ProxyAuthentication <ProxyAuthentication>] [-ProxyCredential <pscredential>] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-SPNPort <int>] [-OperationTimeout <int>] [-NoEncryption] [-UseUTF16] [<CommonParameters>]
```

Example: Create a connection that uses connection options

```powershell
PS C:\> $a = New-WSManSessionOption -OperationTimeout 30000
PS C:\> Connect-WSMan -ComputerName "server01" -SessionOption $a
PS C:\> cd WSMan:
PS WSMan:\> dir
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/New-WSManSessionOption.md)


### Optimize-AppxProvisionedPackages

Version: Both

Module: Dism

Syntax:

```powershell
Optimize-AppXProvisionedPackages -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>] Optimize-AppXProvisionedPackages -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Optimize-WindowsImage

Version: Both

Module: Dism

Syntax:

```powershell
Optimize-WindowsImage -OptimizationTarget <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Configure an image as a WIMBoot system

```powershell
PS C:\> Optimize-WindowsImage -Path "c:\" -OptimizationTarget "WIMBoot"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Optimize-WindowsImage.md)


### Out-GridView

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Out-GridView [-InputObject <psobject>] [-Title <string>] [-PassThru] [<CommonParameters>]
Out-GridView [-InputObject <psobject>] [-Title <string>] [-Wait] [<CommonParameters>]
Out-GridView [-InputObject <psobject>] [-Title <string>] [-OutputMode <OutputModeOption>] [<CommonParameters>]
```

Example: Output processes to a grid view

```powershell
Get-Process | Out-GridView
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Out-GridView.md)


### Out-Printer

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Out-Printer [[-Name] <string>] [-InputObject <psobject>] [<CommonParameters>]
```

Example: Send a file to be printed on the default printer

```powershell
Get-Content -Path ./readme.txt | Out-Printer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Out-Printer.md)


### Publish-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Publish-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [[-DynamicUserConfigurationPath] <string>] [-Global] [-UserSID <string>] [-DynamicUserConfigurationType <DynamicUserConfiguration>] [<CommonParameters>]
Publish-AppvClientPackage [-Package] <AppvClientPackage> [[-DynamicUserConfigurationPath] <string>] [-Global] [-UserSID <string>] [-DynamicUserConfigurationType <DynamicUserConfiguration>] [<CommonParameters>]
Publish-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [-UserSID <string>] [-DynamicUserConfigurationPath <string>] [-DynamicUserConfigurationType <DynamicUserConfiguration>] [<CommonParameters>]
```

Example: Publish a version of a package to all users

```powershell
PS C:\> Publish-AppvClientPackage -Name "MyApp" -Version 1 -Global -DynamicUserConfiguration "C:\content\policies\MyApp.policy"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Publish-AppvClientPackage.md)


### Publish-DscConfiguration

Version: Both

Module: PSDesiredStateConfiguration

Syntax:

```powershell
Publish-DscConfiguration [-Path] <string> [[-ComputerName] <string[]>] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Publish-DscConfiguration [-Path] <string> -CimSession <CimSession[]> [-Force] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Publish a configuration to a remote computer

```powershell
Publish-DscConfiguration -Path '$home\WebServer' -ComputerName "ContosoWebServer" -Credential (get-credential Contoso\webadministrator)
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/publish-dscconfiguration?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/publish-dscconfiguration?view=powershell-7.5)


### Receive-DtcDiagnosticTransaction

Version: Both

Module: MsDtc

Syntax:

```powershell
Receive-DtcDiagnosticTransaction [[-ComputerName] <string>] [[-Port] <int>] [[-PropagationMethod] <DtcTransactionPropagation>] [<CommonParameters>]
```

Example: Receive a diagnostic transaction

```powershell
PS C:\> Receive-DtcDiagnosticTransaction -ComputerName "Host1" -Port 17123 -PropagationMethod Pull
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Receive-DtcDiagnosticTransaction.md)


### Receive-PSSession

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Receive-PSSession [-Session] <PSSession> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-Id] <int> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ComputerName] <string> -InstanceId <guid> [-ApplicationName <string>] [-ConfigurationName <string>] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ComputerName] <string> -Name <string> [-ApplicationName <string>] [-ConfigurationName <string>] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ConnectionUri] <uri> -Name <string> [-ConfigurationName <string>] [-AllowRedirection] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ConnectionUri] <uri> -InstanceId <guid> [-ConfigurationName <string>] [-AllowRedirection] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-InstanceId] <guid> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-Name] <string> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Receive-PSSession [-Session] <PSSession> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-Id] <int> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ComputerName] <string> -Name <string> [-ApplicationName <string>] [-ConfigurationName <string>] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ComputerName] <string> -InstanceId <guid> [-ApplicationName <string>] [-ConfigurationName <string>] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-Port <int>] [-UseSSL] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ConnectionUri] <uri> -Name <string> [-ConfigurationName <string>] [-AllowRedirection] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-ConnectionUri] <uri> -InstanceId <guid> [-ConfigurationName <string>] [-AllowRedirection] [-OutTarget <OutTarget>] [-JobName <string>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [-SessionOption <PSSessionOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-InstanceId] <guid> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Receive-PSSession [-Name] <string> [-OutTarget <OutTarget>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Connect to a PSSession

```powershell
Receive-PSSession -ComputerName Server01 -Name ITTask
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Receive-PSSession.md)


### Register-CimIndicationEvent

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint32>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

Syntax (7):

```powershell
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-OperationTimeoutSec <uint>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

Example: Register the events generated by a class

```powershell
$event = @{
    ClassName = 'Win32_ProcessStartTrace'
    SourceIdentifier = 'ProcessStarted'
}
Register-CimIndicationEvent @event
Get-Event -SourceIdentifier "ProcessStarted"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Register-CimIndicationEvent.md)


### Register-PSSessionConfiguration

Version: Both

Module: Microsoft.PowerShell.Core

Syntax (5.1):

```powershell
Register-PSSessionConfiguration [-Name] <string> [-ProcessorArchitecture <string>] [-SessionType <PSSessionType>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> [-AssemblyName] <string> [-ConfigurationTypeName] <string> [-ProcessorArchitecture <string>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> -Path <string> [-ProcessorArchitecture <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-TransportOption <PSTransportOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Register-PSSessionConfiguration [-Name] <string> [-ProcessorArchitecture <string>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> [-AssemblyName] <string> [-ConfigurationTypeName] <string> [-ProcessorArchitecture <string>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> -Path <string> [-ProcessorArchitecture <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-TransportOption <PSTransportOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Register a NewShell session configuration

```powershell
$sessionConfiguration = @{
    Name='NewShell'
    ApplicationBase='C:\MyShells\'
    AssemblyName='MyShell.dll'
    ConfigurationTypeName='MyClass'
}
Register-PSSessionConfiguration @sessionConfiguration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Register-PSSessionConfiguration.md)


### Register-RecoveryManagementPlugin

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint32> -Path <string> [-CapabilitiesDesired <uint32>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint32> -Online [-CapabilitiesDesired <uint32>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint> -Path <string> [-CapabilitiesDesired <uint>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint> -Online [-CapabilitiesDesired <uint>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Register-ScheduledJob

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Register-ScheduledJob [-Name] <string> [-ScriptBlock] <scriptblock> [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-ArgumentList <Object[]>] [-MaxResultCount <int>] [-RunNow] [-RunEvery <timespan>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-ScheduledJob [-Name] <string> [-FilePath] <string> [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-ArgumentList <Object[]>] [-MaxResultCount <int>] [-RunNow] [-RunEvery <timespan>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Create a scheduled job

```powershell
Register-ScheduledJob -Name "Archive-Scripts" -ScriptBlock {
  Get-ChildItem $HOME\*.ps1 -Recurse |
  Copy-Item -Destination "\\Server\Share\PSScriptArchive"
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Register-ScheduledJob.md)


### Register-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Register-UevTemplate [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Register-UevTemplate -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Register a template

```powershell
PS C:\> Register-UevTemplate -Path "MicrosoftCalculator.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Register-UevTemplate.md)


### Register-WmiEvent

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Register-WmiEvent [-Class] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-Credential <pscredential>] [-ComputerName <string>] [-Timeout <long>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-WmiEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-Credential <pscredential>] [-ComputerName <string>] [-Timeout <long>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

Example (5.1): Subscribe to events generated by a class

```powershell
Register-WmiEvent -Class 'Win32_ProcessStartTrace' -SourceIdentifier "ProcessStarted"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Register-WmiEvent.md)


### Remove-AppProvisionedSharedPackageContainer

Version: Both

Module: Dism

Syntax:

```powershell
Remove-AppProvisionedSharedPackageContainer -Name <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-AppProvisionedSharedPackageContainer -Name <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Remove-AppSharedPackageContainer

Version: Both

Module: Appx

Syntax:

```powershell
Remove-AppSharedPackageContainer [-Name] <string> [-ForceApplicationShutdown] [<CommonParameters>]
```

Example: 

```powershell
Remove-AppSharedPackageContainer -Name ContosoTestContainer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Remove-AppSharedPackageContainer.md)


### Remove-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Remove-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [<CommonParameters>]
Remove-AppvClientConnectionGroup [-Name] <string> [<CommonParameters>]
Remove-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [<CommonParameters>]
```

Example: Remove a named connection group

```powershell
PS C:\> Remove-AppvClientConnectionGroup -Name "MyGroup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Remove-AppvClientConnectionGroup.md)


### Remove-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Remove-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [<CommonParameters>]
Remove-AppvClientPackage [-Package] <AppvClientPackage> [<CommonParameters>]
Remove-AppvClientPackage [-Name] <string> [[-Version] <string>] [<CommonParameters>]
```

Example: Remove a version of a package by using the pipeline operator

```powershell
PS C:\> Get-AppvPackage -Name "MyPackage" -Version 1 | Remove-Package
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Remove-AppvClientPackage.md)


### Remove-AppvPublishingServer

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Remove-AppvPublishingServer [-ServerId] <uint32> [<CommonParameters>]
Remove-AppvPublishingServer [-Server] <AppvPublishingServer> [<CommonParameters>]
Remove-AppvPublishingServer [[-Name] <string>] [[-URL] <string>] [<CommonParameters>]
```

Example: Remove a publishing server

```powershell
PS C:\> Remove-AppvPublishingServer -Name "Server01"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Remove-AppvPublishingServer.md)


### Remove-AppxPackage

Version: Both

Module: Appx

Syntax:

```powershell
Remove-AppxPackage [-Package] <string> [-PreserveApplicationData] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-AppxPackage [-Package] <string> [-PreserveRoamableApplicationData] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-AppxPackage [-Package] <string> [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-AppxPackage [-Package] <string> -User <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove an app package

```powershell
Remove-AppxPackage -Package 'package1_1.0.0.0_neutral__8wekyb3d8bbwe'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Remove-AppxPackage.md)


### Remove-AppxPackageAutoUpdateSettings

Version: Both

Module: Appx

Syntax:

```powershell
Remove-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> [-UseSystemPolicySource] [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Remove-AppxPackageAutoUpdateSettings -PackageFullName publisher.package1_1.0.0.0_neutral__8wekyb3d8bbwe
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Remove-AppxPackageAutoUpdateSettings.md)


### Remove-AppxProvisionedPackage

Version: Both

Module: Dism

Syntax:

```powershell
Remove-AppxProvisionedPackage -PackageName <string> -Online [-AllUsers] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-AppxProvisionedPackage -PackageName <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Remove an app package from an image

```powershell
PS C:\> Remove-AppxProvisionedPackage -Path c:\offline -PackageName MyAppxPkg
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Remove-AppxProvisionedPackage.md)


### Remove-AppxVolume

Version: Both

Module: Appx

Syntax:

```powershell
Remove-AppxVolume [-Volume] <AppxVolume[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a volume by using an ID

```powershell
Remove-AppxVolume -Volume {984786d3-0cae-49de-a68f-8bedb0ca260b}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Remove-AppxVolume.md)


### Remove-BcdElement

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Remove-BcdElement [-Element] <string> [[-Id] <string>] [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-BcdElement [-Element] <string> [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Remove-BcdElement.md)


### Remove-BcdEntry

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Remove-BcdEntry [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-BcdEntry [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Remove-BcdEntry.md)


### Remove-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Remove-BitsTransfer [-BitsJob] <BitsJob[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Cancel all BITS transfer jobs owned by the current user

```powershell
PS C:\> Get-BitsTransfer | Remove-BitsTransfer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Remove-BitsTransfer.md)


### Remove-CertificateEnrollmentPolicyServer

Version: Both

Module: PKI

Syntax:

```powershell
Remove-CertificateEnrollmentPolicyServer [-Url] <uri> -context <Context> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    Url = 'https://www.contoso.com/policy/service.svc'
    Context = 'User'
}
Remove-CertificateEnrollmentPolicyServer @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Remove-CertificateEnrollmentPolicyServer.md)


### Remove-CertificateNotificationTask

Version: Both

Module: PKI

Syntax:

```powershell
Remove-CertificateNotificationTask [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Remove-CertificateNotificationTask -Name "My Task"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Remove-CertificateNotificationTask.md)


### Remove-CimInstance

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Remove-CimInstance [-InputObject] <ciminstance> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-CimInstance [-InputObject] <ciminstance> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove the CIM instance

```powershell
Remove-CimInstance -Query 'Select * from Win32_Environment where name LIKE "testvar%"'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Remove-CimInstance.md)


### Remove-CimSession

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Remove-CimSession [-CimSession] <CimSession[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-ComputerName] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-Id] <uint32[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -InstanceId <guid[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -Name <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-CimSession [-CimSession] <CimSession[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-ComputerName] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-Id] <uint[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -InstanceId <guid[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -Name <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove all the CIM sessions

```powershell
Get-CimSession | Remove-CimSession
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Remove-CimSession.md)


### Remove-CIPolicyRule

Version: Both

Module: ConfigCI

Syntax:

```powershell
Remove-CIPolicyRule [-Id] <string> -FilePath <string> [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Remove-CIPolicyRule.md)


### Remove-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Remove-Computer [[-UnjoinDomainCredential] <pscredential>] [-Restart] [-Force] [-PassThru] [-WorkgroupName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Computer -UnjoinDomainCredential <pscredential> [-LocalCredential <pscredential>] [-Restart] [-ComputerName <string[]>] [-Force] [-PassThru] [-WorkgroupName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Remove the local computer from its domain

```powershell
Remove-Computer -UnjoinDomaincredential Domain01\Admin01 -PassThru -Verbose -Restart
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Remove-Computer.md)


### Remove-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Remove-EventLog [-LogName] <string[]> [[-ComputerName] <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-EventLog [[-ComputerName] <string[]>] [-Source <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Remove an event log from the local computer

```powershell
Remove-EventLog -LogName "MyLog"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Remove-EventLog.md)


### Remove-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Remove-JobTrigger [-InputObject] <ScheduledJobDefinition[]> [-TriggerId <int[]>] [<CommonParameters>]
Remove-JobTrigger [-Id] <int[]> [-TriggerId <int[]>] [<CommonParameters>]
Remove-JobTrigger [-Name] <string[]> [-TriggerId <int[]>] [<CommonParameters>]
```

Example (5.1): Delete all job triggers

```powershell
Remove-JobTrigger -Name "Test*"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Remove-JobTrigger.md)


### Remove-LocalGroup

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Remove-LocalGroup [-InputObject] <LocalGroup[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroup [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroup [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Delete a security group

```powershell
Remove-LocalGroup -Name "SecurityGroup04"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Remove-LocalGroup.md)


### Remove-LocalGroupMember

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Remove-LocalGroupMember [-Group] <LocalGroup> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroupMember [-Name] <string> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroupMember [-SID] <SecurityIdentifier> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Remove members from the Administrators group

```powershell
$members = "Admin02", "MicrosoftAccount\username@Outlook.com", "AzureAD\DavidChew@contoso.com", "CONTOSO\Domain Admins"
Remove-LocalGroupMember -Group "Administrators" -Member $members
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Remove-LocalGroupMember.md)


### Remove-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Remove-LocalUser [-InputObject] <LocalUser[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalUser [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalUser [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Delete a user account

```powershell
Remove-LocalUser -Name "AdminContoso02"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Remove-LocalUser.md)


### Remove-OsConfigurationDocument

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Remove-OsConfigurationDocument [-Id] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [-Wait] [[-TimeoutInSeconds] <int>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Remove-OsConfigurationDocument.md)


### Remove-OSConfigurationScenarioDefinition

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Remove-OsConfigurationScenarioDefinition [-Name] <string> [-Version] <string> [-SchemaVersion] <string> [<CommonParameters>]
```

Example: none

Source: [OsConfiguration module documentation](https://learn.microsoft.com/en-us/powershell/module/osconfiguration) (no dedicated page).


### Remove-PmemDedicatedMemory

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Remove-PmemDedicatedMemory -DeviceNumber <uint32> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove dedicated persistent memory

```powershell
Remove-PmemDedicatedMemory -DeviceNumber 1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Remove-PmemDedicatedMemory.md)


### Remove-PmemDisk

Version: Both

Module: PersistentMemory

Syntax:

```powershell
Remove-PmemDisk -DiskNumber <uint32> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PmemDisk -Simulated [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a persistent memory disk

```powershell
Remove-PmemDisk -DiskNumber 2
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/PersistentMemory/Remove-PmemDisk.md)


### Remove-PSSnapin

Version: 5.1 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Remove-PSSnapin [-Name] <string[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a snap-in

```powershell
Remove-PSSnapin -Name Microsoft.Exchange
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Remove-PSSnapin.md)


### Remove-RecoveryManagementPluginAltitude

Version: Both

Module: Dism

Syntax:

```powershell
Remove-RecoveryManagementPluginAltitude -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-RecoveryManagementPluginAltitude -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Remove-Service

Version: 7 only

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Remove-Service [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Service [-InputObject <ServiceController>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Remove a service

```powershell
Remove-Service -Name "TestService"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Remove-Service.md)


### Remove-WindowsCapability

Version: Both

Module: Dism

Syntax:

```powershell
Remove-WindowsCapability -Name <string> -Online [-DelayExecutionIfPending] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsCapability -Name <string> -Path <string> [-DelayExecutionIfPending] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Remove Windows capabilities for an image

```powershell
PS C:\> Remove-WindowsCapability -Name "Language.TextToSpeech~~~fr-FR~0.0.1.0" -Path "C:\offline"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Remove-WindowsCapability.md)


### Remove-WindowsDriver

Version: Both

Module: Dism

Syntax:

```powershell
Remove-WindowsDriver -Driver <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Remove a driver from an image

```powershell
PS C:\> Remove-WindowsDriver -Path "c:\offline" -Driver "OEM1.inf"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Remove-WindowsDriver.md)


### Remove-WindowsImage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Remove-WindowsImage -ImagePath <string> -Name <string> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsImage -ImagePath <string> -Index <uint32> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Remove-WindowsImage -ImagePath <string> -Name <string> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsImage -ImagePath <string> -Index <uint> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Remove the first image from a WIM file

```powershell
PS C:\> Remove-WindowsImage -ImagePath "c:\imagestore\custom.wim" -Index 1 -CheckIntegrity
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Remove-WindowsImage.md)


### Remove-WindowsPackage

Version: Both

Module: Dism

Syntax:

```powershell
Remove-WindowsPackage -Path <string> [-PackagePath <string>] [-PackageName <string>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsPackage -Online [-PackagePath <string>] [-PackageName <string>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Remove a package from a running operating system image

```powershell
PS C:\> Remove-WindowsPackage -Online -PackageName "Microsoft-Windows-Backup-Package~31bf3856ad364e35~x86~~6.1.7601.16525"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Remove-WindowsPackage.md)


### Remove-WmiObject

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Remove-WmiObject [-Class] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject -InputObject <wmi> [-AsJob] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject -Path <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Close all instances of a Win32 process

```powershell
notepad
$np = Get-WmiObject -Query "select * from Win32_Process where name='notepad.exe'"
$np | Remove-WmiObject
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Remove-WmiObject.md)


### Remove-WSManInstance

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Remove-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ApplicationName <string>] [-ComputerName <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Remove-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ConnectionURI <uri>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Delete a listener

```powershell
Remove-WSManInstance -ResourceUri winrm/config/Listener -SelectorSet @{
    Address   = 'test.fabrikam.com'
    Transport = 'http'
}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Remove-WSManInstance.md)


### Rename-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Rename-Computer [-NewName] <string> [-ComputerName <string>] [-PassThru] [-DomainCredential <pscredential>] [-LocalCredential <pscredential>] [-Force] [-Restart] [-WsmanAuthentication <string>] [-Protocol <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Rename-Computer [-NewName] <string> [-ComputerName <string>] [-PassThru] [-DomainCredential <pscredential>] [-LocalCredential <pscredential>] [-Force] [-Restart] [-WsmanAuthentication <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Rename the local computer

```powershell
Rename-Computer -NewName "Server044" -DomainCredential Domain01\Admin01 -Restart
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Rename-Computer.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`sudo hostnamectl set-hostname`).
- Companions: Restart-Computer, Stop-Computer.
- Distro: needs sudo.
- Function: changes the host name.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-NewName` (position 0) | string | New host name; handed to `sudo hostnamectl set-hostname` |


### Rename-LocalGroup

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Rename-LocalGroup [-InputObject] <LocalGroup> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalGroup [-Name] <string> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalGroup [-SID] <SecurityIdentifier> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Change the name of a group

```powershell
PS C:\> Rename-LocalGroup -Name "SecurityGroup" -NewName "SecurityGroup04"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Rename-LocalGroup.md)


### Rename-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Rename-LocalUser [-InputObject] <LocalUser> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalUser [-Name] <string> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalUser [-SID] <SecurityIdentifier> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Rename a user account

```powershell
Rename-LocalUser -Name "Admin02" -NewName "AdminContoso02"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Rename-LocalUser.md)


### Repair-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Repair-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientConnectionGroup [-Name] <string> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
```

Example: Repair a named connection group

```powershell
PS C:\> Repair-AppvClientConnectionGroup -Name "MyGroup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Repair-AppvClientConnectionGroup.md)


### Repair-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Repair-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientPackage [-Package] <AppvClientPackage> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [-UserState] [-Extensions] [<CommonParameters>]
```

Example: Delete user state for a version of a package

```powershell
PS C:\> Repair-AppvClientPackage -Name "MyApp" -Version 3
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Repair-AppvClientPackage.md)


### Repair-UevTemplateIndex

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Repair-UevTemplateIndex [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Repair the template index

```powershell
PS C:\> Repair-UevTemplateIndex
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Repair-UevTemplateIndex.md)


### Repair-WindowsImage

Version: Both

Module: Dism

Syntax:

```powershell
Repair-WindowsImage -Path <string> [-CheckHealth] [-ScanHealth] [-RestoreHealth] [-StartComponentCleanup] [-LimitAccess] [-ResetBase] [-Defer] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Repair-WindowsImage -Online [-CheckHealth] [-ScanHealth] [-RestoreHealth] [-StartComponentCleanup] [-LimitAccess] [-ResetBase] [-Defer] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Scan an image for corruption

```powershell
PS C:\> Repair-WindowsImage -Path "C:\offline\Mount" -ScanHealth
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Repair-WindowsImage.md)


### Reset-AppSharedPackageContainer

Version: Both

Module: Appx

Syntax:

```powershell
Reset-AppSharedPackageContainer [-Name] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Reset-AppSharedPackageContainer -Name ContosoTestContainer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Reset-AppSharedPackageContainer.md)


### Reset-AppxPackage

Version: Both

Module: Appx

Syntax:

```powershell
Reset-AppxPackage [-Package] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Reset app package

```powershell
Reset-AppxPackage -Package publisher.package1_1.0.0.0_neutral__8wekyb3d8bbwe
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Reset-AppxPackage.md)


### Reset-ComputerMachinePassword

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Reset-ComputerMachinePassword [-Server <string>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Reset the password for the local computer

```powershell
Reset-ComputerMachinePassword
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Reset-ComputerMachinePassword.md)


### Reset-LapsPassword

Version: Both

Module: LAPS

Syntax:

```powershell
Reset-LapsPassword [<CommonParameters>]
```

Example: 

```powershell
Reset-LapsPassword
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Reset-LapsPassword.md)


### Resolve-DnsName

Version: Both

Module: DnsClient

Syntax:

```powershell
Resolve-DnsName [-Name] <string> [[-Type] <RecordType>] [-Server <string[]>] [-DohServer <string[][]>] [-DotServer <string[][]>] [-DnsOnly] [-CacheOnly] [-DnssecOk] [-DnssecCd] [-NoHostsFile] [-LlmnrNetbiosOnly] [-LlmnrFallback] [-LlmnrOnly] [-NetbiosFallback] [-NoIdn] [-NoRecursion] [-QuickTimeout] [-TcpOnly] [-CheckCache] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Resolve-DnsName -Name www.bing.com
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/DnsClient/Resolve-DnsName.md)


### Restart-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Restart-Service [-InputObject] <ServiceController[]> [-Force] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Restart-Service [-Name] <string[]> [-Force] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Restart-Service -DisplayName <string[]> [-Force] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Restart a service on the local computer

```powershell
PS C:\> Restart-Service -Name winmgmt
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Restart-Service.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`systemctl` + sudo).
- Companions: Stop-Service, Restart-Service, Resume-Service.
- Distro: systemd-based + sudo.
- Function: starts/stops/restarts services.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name; handed to `systemctl <action> <unit>` (.service suffix added automatically) |

- Implementation: `systemctl start/stop/restart <unit>`; a failure under ordinary permissions retries automatically with `sudo`. Counterpart: `sudo systemctl start/stop/restart`.
- Stop-Service, Restart-Service, and Resume-Service share Start-Service's parameters, their actions being stop/restart/start.


### Restore-Computer

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Restore-Computer [-RestorePoint] <int> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Restore the local computer

```powershell
Restore-Computer -RestorePoint 253
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Restore-Computer.md)


### Restore-UevBackup

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Restore-UevBackup [-ComputerName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Restore backed up settings from another computer

```powershell
PS C:\>Restore-UevBackup -ComputerName "PattiFullerDevice03@Contoso.Com"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Restore-UevBackup.md)


### Restore-UevUserSetting

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Restore-UevUserSetting -Application <string> [-Force] [-LastKnownGood] [-WhatIf] [-Confirm] [<CommonParameters>]
Restore-UevUserSetting [-TemplateId] <string> [-LastKnownGood] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Restore user settings for a specific template

```powershell
PS C:\> Restore-UevUserSetting -TemplateId "MicrosoftCalculator6"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Restore-UevUserSetting.md)


### Resume-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Resume-BitsTransfer [-BitsJob] <BitsJob[]> [-Asynchronous] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Resume all BITS transfer jobs owned by the current user

```powershell
PS C:\> Get-BitsTransfer | Resume-BitsTransfer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Resume-BitsTransfer.md)


### Resume-Job

Version: 5.1 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Resume-Job [-Id] <int[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-Job] <Job[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-Name] <string[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-InstanceId] <guid[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-State] <JobState> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-Filter] <hashtable> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Resume a job by ID

```powershell
PS C:\> Get-Job EventJob
Id     Name            PSJobTypeName   State         HasMoreData     Location   Command
--     ----            -------------   -----         -----------     --------   -------
4      EventJob        PSWorkflowJob   Suspended     True            Server01   \\Script\Share\Event.ps1

PS C:\> Resume-Job -Id 4
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Resume-Job.md)


### Resume-ProvisioningSession

Version: Both

Module: Provisioning

Syntax: none

Example: none

Source: [Provisioning module documentation](https://learn.microsoft.com/en-us/powershell/module/provisioning) (no dedicated page).


### Resume-ReFSDedupSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Resume-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Resume-ReFSDedupSchedule -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Resume-ReFSDedupSchedule.md)


### Resume-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Resume-Service [-InputObject] <ServiceController[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Service [-Name] <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Service -DisplayName <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Resume a service on the local computer

```powershell
PS C:\> Resume-Service "sens"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Resume-Service.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`systemctl` + sudo).
- Companions: Stop-Service, Restart-Service, Resume-Service.
- Distro: systemd-based + sudo.
- Function: starts/stops/restarts services.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name; handed to `systemctl <action> <unit>` (.service suffix added automatically) |

- Implementation: `systemctl start/stop/restart <unit>`; a failure under ordinary permissions retries automatically with `sudo`. Counterpart: `sudo systemctl start/stop/restart`.
- Stop-Service, Restart-Service, and Resume-Service share Start-Service's parameters, their actions being stop/restart/start.


### Save-OsImage

Version: Both

Module: Dism

Syntax:

```powershell
Save-OsImage -ImagePath <string> -CapturePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Save-SoftwareInventory

Version: Both

Module: Dism

Syntax:

```powershell
Save-SoftwareInventory -PartitioningScript <string> -ResetConfigXml <string> -Path <string> [-DevicesInf <string>] [-OutputDirectory <string>] [-CSRConfigFile <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Save-SoftwareInventory -PartitioningScript <string> -ResetConfigXml <string> -Online [-DevicesInf <string>] [-OutputDirectory <string>] [-CSRConfigFile <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Save-WindowsImage

Version: Both

Module: Dism

Syntax:

```powershell
Save-WindowsImage -Path <string> [-CheckIntegrity] [-Append] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Save servicing changes made to a mounted image

```powershell
PS C:\> Save-WindowsImage -Path "c:\offline"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Save-WindowsImage.md)


### Send-AppvClientReport

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Send-AppvClientReport [[-URL] <string>] [-NetworkCostAware] [-DeleteOnSuccess] [<CommonParameters>]
```

Example: Send data to previously configured location

```powershell
PS C:\> Send-AppVClientReport
The Application Virtualization Client Report was sent successfully
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Send-AppvClientReport.md)


### Send-DtcDiagnosticTransaction

Version: Both

Module: MsDtc

Syntax:

```powershell
Send-DtcDiagnosticTransaction [-Transaction] <DtcDiagnosticTransaction> [[-ComputerName] <string>] [[-Port] <int>] [[-PropagationMethod] <DtcTransactionPropagation>] [<CommonParameters>]
```

Example: Send a DTC diagnostic transaction

```powershell
PS C:\> $Tx = New-DtcDiagnosticTransaction
PS C:\> Send-DtcDiagnosticTransaction -Transaction $Tx -ComputerName "Host1" -PropagationMethod Push
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Send-DtcDiagnosticTransaction.md)


### Set-Acl

Version: Both

Module: Microsoft.PowerShell.Security

Syntax (5.1):

```powershell
Set-Acl [-Path] <string[]> [-AclObject] <Object> [[-CentralAccessPolicy] <string>] [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-Acl [-InputObject] <psobject> [-AclObject] <Object> [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-Acl [-AclObject] <Object> [[-CentralAccessPolicy] <string>] -LiteralPath <string[]> [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Acl [-Path] <string[]> [-AclObject] <Object> [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Acl [-InputObject] <psobject> [-AclObject] <Object> [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Acl [-AclObject] <Object> -LiteralPath <string[]> [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Copy a security descriptor from one file to another

```powershell
$DogACL = Get-Acl -Path "C:\Dog.txt"
Set-Acl -Path "C:\Cat.txt" -AclObject $DogACL
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Set-Acl.md)


### Set-AppBackgroundTaskResourcePolicy

Version: Both

Module: AppBackgroundTask

Syntax:

```powershell
Set-AppBackgroundTaskResourcePolicy -Mode <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set global resource policy to Conservative mode

```powershell
PS C:\> Set-AppBackgroundTaskResourcePolicy -Mode Conservative
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppBackgroundTask/Set-AppBackgroundTaskResourcePolicy.md)


### Set-AppLockerPolicy

Version: Both

Module: AppLocker

Syntax:

```powershell
Set-AppLockerPolicy [-XmlPolicy] <string> [-Ldap <string>] [-Merge] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppLockerPolicy [-PolicyObject] <AppLockerPolicy> [-Ldap <string>] [-Merge] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set the local AppLocker policy

```powershell
PS C:\> Set-AppLockerPolicy -XMLPolicy C:\Policy.xml
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppLocker/Set-AppLockerPolicy.md)


### Set-AppvClientConfiguration

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Set-AppvClientConfiguration [-AllowHighCostLaunch <bool>] [-AutoLoad <uint32>] [-AutoCleanupEnabled <bool>] [-CertFilterForClientSsl <string>] [-EnablePackageScripts <bool>] [-EnablePublishingRefreshUI <bool>] [-IntegrationRootGlobal <string>] [-IntegrationRootUser <string>] [-LocationProvider <string>] [-MigrationMode <bool>] [-PackageInstallationRoot <string>] [-PackageSourceRoot <string>] [-RequirePublishAsAdmin <bool>] [-ReestablishmentInterval <uint32>] [-ReestablishmentRetries <uint32>] [-ReportingDataBlockSize <uint32>] [-ReportingDataCacheLimit <uint32>] [-ReportingEnabled <bool>] [-ReportingInterval <uint32>] [-ReportingRandomDelay <uint32>] [-ReportingServerURL <string>] [-ReportingStartTime <uint32>] [-RoamingFileExclusions <string>] [-RoamingRegistryExclusions <string>] [-SharedContentStoreMode <bool>] [-VerifyCertificateRevocationList <bool>] [-ExperienceImprovementOptIn <bool>] [-ProcessesUsingVirtualComponents <string[]>] [-EnableDynamicVirtualization <bool>] [-IgnoreLocationProvider <bool>] [-SupportBranchCache <bool>] [-SyncOnBatteriesEnabled <bool>] [<CommonParameters>]
```

Example: Set a client configuration parameter

```powershell
PS C:\> Set-AppvClientConfiguration -parameter1 "parameterVal1"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Set-AppvClientConfiguration.md)


### Set-AppvClientMode

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Set-AppvClientMode -Normal [<CommonParameters>]
Set-AppvClientMode -Uninstall [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Set-AppvClientMode.md)


### Set-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Set-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Path <string>] [-DynamicDeploymentConfiguration <string>] [-UseNoConfiguration] [<CommonParameters>]
Set-AppvClientPackage [-Package] <AppvClientPackage> [-Path <string>] [-DynamicDeploymentConfiguration <string>] [-UseNoConfiguration] [<CommonParameters>]
Set-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Path <string>] [-DynamicDeploymentConfiguration <string>] [-UseNoConfiguration] [<CommonParameters>]
```

Example: Set a deployment configuration for a package

```powershell
PS C:\> Set-AppvClientPackage -Name "MyApp" -Version 1 -DynamicDeploymentConfiguration "C:\policies\MyApp.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Set-AppvClientPackage.md)


### Set-AppvPublishingServer

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Set-AppvPublishingServer [-ServerId] <uint32> [[-GlobalRefreshEnabled] <bool>] [[-GlobalRefreshOnLogon] <bool>] [[-GlobalRefreshInterval] <uint32>] [[-GlobalRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [[-UserRefreshEnabled] <bool>] [[-UserRefreshOnLogon] <bool>] [[-UserRefreshInterval] <uint32>] [[-UserRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [<CommonParameters>]
Set-AppvPublishingServer [-Server] <AppvPublishingServer> [[-GlobalRefreshEnabled] <bool>] [[-GlobalRefreshOnLogon] <bool>] [[-GlobalRefreshInterval] <uint32>] [[-GlobalRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [[-UserRefreshEnabled] <bool>] [[-UserRefreshOnLogon] <bool>] [[-UserRefreshInterval] <uint32>] [[-UserRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [<CommonParameters>]
```

Example: none

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Set-AppvPublishingServer.md)


### Set-AppxDefaultVolume

Version: Both

Module: Appx

Syntax:

```powershell
Set-AppxDefaultVolume [-Volume] <AppxVolume> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set a default volume by using a path

```powershell
Set-AppxDefaultVolume -Volume F:\
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Set-AppxDefaultVolume.md)


### Set-AppxPackageAutoUpdateSettings

Version: Both

Module: Appx

Syntax (5.1):

```powershell
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> [-AppInstallerUri <string>] [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint32>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -AppInstallerUri <string> -ClearPreviousSettings [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint32>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -PauseUpdates -HoursToPause <uint32> [-UseSystemPolicySource] [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> [-AppInstallerUri <string>] [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -AppInstallerUri <string> -ClearPreviousSettings [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -PauseUpdates -HoursToPause <uint> [-UseSystemPolicySource] [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Update the Auto Update settings for an App

```powershell
$params = @{
    AppInstallerUri           =  'https://website.com/PackageName.appinstaller '
    PackageFamilyName         =  'PackageName_8h66172c634n0 '
    CheckOnLaunch             =  $true
    ForceUpdateFromAnyVersion =  $true
    HoursBetweenUpdateChecks  =  2
    ShowPrompt                =  $true
    UpdateUris                =  'file://ComputerName/Share/PackageName_x64.appinstaller'
}
Set-AppxPackageAutoUpdateSettings @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Appx/Set-AppxPackageAutoUpdateSettings.md)


### Set-AppXProvisionedDataFile

Version: Both

Module: Dism

Syntax:

```powershell
Set-AppXProvisionedDataFile -PackageName <string> -CustomDataPath <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-AppXProvisionedDataFile -PackageName <string> -CustomDataPath <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Add a custom data file to an app package for the running operating system

```powershell
PS C:\> Set-AppXProvisionedDataFile -Online -PackageName "MyAppxPkg" -CustomDataPath "c:\Appx\myCustomData.dat"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Set-AppXProvisionedDataFile.md)


### Set-AuthenticodeSignature

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Set-AuthenticodeSignature [-FilePath] <string[]> [-Certificate] <X509Certificate2> [-IncludeChain <string>] [-TimestampServer <string>] [-HashAlgorithm <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AuthenticodeSignature [-Certificate] <X509Certificate2> -LiteralPath <string[]> [-IncludeChain <string>] [-TimestampServer <string>] [-HashAlgorithm <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AuthenticodeSignature [-Certificate] <X509Certificate2> -SourcePathOrExtension <string[]> -Content <byte[]> [-IncludeChain <string>] [-TimestampServer <string>] [-HashAlgorithm <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Sign a script using a certificate from the local certificate store

```powershell
$cert=Get-ChildItem -Path Cert:\CurrentUser\My -CodeSigningCert
$signingParameters = @{
    FilePath      = 'PsTestInternet2.ps1'
    Certificate   = $cert
    HashAlgorithm = 'SHA256'
}
Set-AuthenticodeSignature @signingParameters
```

Example (7): Sign a script using a certificate from the local certificate store

```powershell
$cert = Get-ChildItem -Path Cert:\CurrentUser\My -CodeSigningCert
Set-AuthenticodeSignature -FilePath PsTestInternet2.ps1 -Certificate $cert
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Set-AuthenticodeSignature.md)


### Set-BcdBootDefault

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdBootDefault [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDefault [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdBootDefault.md)


### Set-BcdBootDisplayOrder

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdBootDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] -AddFirst [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Entry] <BcdEntryInfo[]> -AddFirst [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] -AddLast [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Entry] <BcdEntryInfo[]> -AddLast [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] -Remove [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Entry] <BcdEntryInfo[]> -Remove [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDisplayOrder [-Entry] <BcdEntryInfo[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdBootDisplayOrder.md)


### Set-BcdBootSequence

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdBootSequence [-Id] <string[]> [[-Store] <BcdStoreInfo>] -AddFirst [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Entry] <BcdEntryInfo[]> -AddFirst [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Id] <string[]> [[-Store] <BcdStoreInfo>] -AddLast [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Entry] <BcdEntryInfo[]> -AddLast [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Id] <string[]> [[-Store] <BcdStoreInfo>] -Remove [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Entry] <BcdEntryInfo[]> -Remove [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Id] <string[]> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootSequence [-Entry] <BcdEntryInfo[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdBootSequence.md)


### Set-BcdBootTimeout

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdBootTimeout [-Value] <long> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdBootTimeout.md)


### Set-BcdBootToolsDisplayOrder

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdBootToolsDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] -AddFirst [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Entry] <BcdEntryInfo[]> -AddFirst [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] -AddLast [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Entry] <BcdEntryInfo[]> -AddLast [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] -Remove [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Entry] <BcdEntryInfo[]> -Remove [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Id] <string[]> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootToolsDisplayOrder [-Entry] <BcdEntryInfo[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdBootToolsDisplayOrder.md)


### Set-BcdDebugSettings

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -DebugPort <long> -Serial [-Baudrate <long>] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Serial [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Port <long> -HostIp <string> -Net -Key <string> [-NoDhcp] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Port <long> -HostIp <string> -Net [-NewKey] [-NoDhcp] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Channel <long> -Ieee1394 [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Usb [-TargetName <string>] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Local [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdDebugSettings.md)


### Set-BcdElement

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdElement [-Element] <string> [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Type <SetBcdElementCommand+ElementType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdElement [-Element] <string> [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Device <SetBcdElementCommand+DeviceType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdElement [-Element] <string> [-Entry] <BcdEntryInfo> -Type <SetBcdElementCommand+ElementType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdElement [-Element] <string> [-Entry] <BcdEntryInfo> -Device <SetBcdElementCommand+DeviceType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdElement.md)


### Set-BcdHypervisorSettings

Version: Both

Module: Microsoft.Windows.Bcd.Cmdlets

Syntax:

```powershell
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -DebugPort <long> -Serial [-Baudrate <long>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -Ieee1394 [-Channel <long>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -HostIp <string> -Port <long> -Net [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -Serial [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.Windows.Bcd.Cmdlets/Set-BcdHypervisorSettings.md)


### Set-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Set-BitsTransfer [-BitsJob] <BitsJob[]> [-DisplayName <string>] [-Priority <string>] [-Description <string>] [-Dynamic <switch>] [-CustomHeadersWriteOnly] [-HttpMethod <string>] [-ProxyAuthentication <string>] [-RetryInterval <int>] [-RetryTimeout <int>] [-MaxDownloadTime <int>] [-TransferPolicy <CostStates>] [-ACLFlags <ACLFlagValue>] [-SecurityFlags <SecurityFlagValue>] [-UseStoredCredential <AuthenticationTargetValue>] [-Credential <pscredential>] [-ProxyCredential <pscredential>] [-Authentication <string>] [-SetOwnerToCurrentUser] [-ProxyUsage <string>] [-ProxyList <uri[]>] [-ProxyBypass <string[]>] [-CustomHeaders <string[]>] [-NotifyFlags <NotifyFlagValue>] [-NotifyCmdLine <string[]>] [-CertStoreLocation <CertStoreLocationValue>] [-CertStoreName <string>] [-CertHash <byte[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Modify the priority of a BITS transfer job

```powershell
PS C:\> $Bits = Get-BitsTransfer -JobId 10778CFA-C1D7-4A82-8A9D-80B19224879C
PS C:\> Set-BitsTransfer -BitsJob $Bits -Priority High
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Set-BitsTransfer.md)


### Set-CertificateAutoEnrollmentPolicy

Version: Both

Module: PKI

Syntax:

```powershell
Set-CertificateAutoEnrollmentPolicy -PolicyState <PolicySetting> -context <Context> [-StoreName <string[]>] [-ExpirationPercentage <int>] [-EnableTemplateCheck] [-EnableMyStoreManagement] [-EnableBalloonNotifications] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CertificateAutoEnrollmentPolicy -EnableAll -context <Context> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    PolicyState = 'Enabled'
    EnableMyStoreManagement = $true
    EnableTemplateCheck = $true
    Context = 'User'
}
Set-CertificateAutoEnrollmentPolicy @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Set-CertificateAutoEnrollmentPolicy.md)


### Set-CimInstance

Version: Both

Module: CimCmdlets

Syntax (5.1):

```powershell
Set-CimInstance [-InputObject] <ciminstance> [-ComputerName <string[]>] [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -CimSession <CimSession[]> -Property <IDictionary> [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -Property <IDictionary> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-CimInstance [-InputObject] <ciminstance> [-ComputerName <string[]>] [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -CimSession <CimSession[]> -Property <IDictionary> [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -Property <IDictionary> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set the CIM instance

```powershell
$instance = @ {
    Query = 'Select * from Win32_Environment where name LIKE "testvar%"'
    Property = @{VariableValue="abcd"}
}
Set-CimInstance @instance
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/CimCmdlets/Set-CimInstance.md)


### Set-CIPolicyIdInfo

Version: Both

Module: ConfigCI

Syntax:

```powershell
Set-CIPolicyIdInfo [-FilePath] <string> [-PolicyName <string>] [-SupplementsBasePolicyID <guid>] [-BasePolicyToSupplementPath <string>] [-ResetPolicyID] [-PolicyId <string>] [-AppIdTaggingPolicy] [-AppIdTaggingKey <string[]>] [-AppIdTaggingValue <string[]>] [<CommonParameters>]
```

Example: Modify the ID and name of a policy

```powershell
PS C:\> Set-CIPolicyIdInfo -FilePath ".\Policy03.xml" -PolicyId "CIP077" -PolicyName "CIPolicy03"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Set-CIPolicyIdInfo.md)


### Set-CIPolicySetting

Version: Both

Module: ConfigCI

Syntax:

```powershell
Set-CIPolicySetting [-FilePath] <string> -Provider <string> -Key <string> -ValueName <string> -ValueType <string> -Value <string> [<CommonParameters>]
Set-CIPolicySetting [-FilePath] <string> -Provider <string> -Key <string> -ValueName <string> -Delete [<CommonParameters>]
```

Example: Set the Code Integrity policy

```powershell
Set-CIPolicySetting -FilePath C:\Policies\WDAC_policy.xml -Key "{12345678-9abc-def0-1234-56789abcdef0}" -Provider WSH -Value $True -ValueName EnterpriseDefinedClsId -ValueType Boolean
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Set-CIPolicySetting.md)


### Set-CIPolicyVersion

Version: Both

Module: ConfigCI

Syntax:

```powershell
Set-CIPolicyVersion -FilePath <string> -Version <string> [<CommonParameters>]
```

Example: Update the version number of a policy

```powershell
PS C:\> Set-CIPolicyVersion -FilePath '.\Policy.xml' -Version '11.1.0.2'
PS C:\> Get-Content -Path '.Policy.xml'
<?xml version="1.0" encoding="utf-8"?>
<SiPolicy xmlns="urn:schemas-microsoft-com:sipolicy">
  <VersionEx>11.1.0.2</VersionEx>
  <PolicyTypeID>{A244370E-44C9-4C06-B551-F6016E563076}</PolicyTypeID>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Set-CIPolicyVersion.md)


### Set-Culture

Version: Both

Module: International

Syntax:

```powershell
Set-Culture [-CultureInfo] <cultureinfo> [<CommonParameters>]
```

Example: Set the culture

```powershell
PS C:\> Set-Culture -CultureInfo de-DE
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-Culture.md)


### Set-DscLocalConfigurationManager

Version: Both

Module: PSDesiredStateConfiguration

Syntax:

```powershell
Set-DscLocalConfigurationManager [-Path] <string> [[-ComputerName] <string[]>] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-DscLocalConfigurationManager [-Path] <string> -CimSession <CimSession[]> [-Force] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Apply LCM settings

```powershell
Set-DscLocalConfigurationManager -Path "C:\DSC\Configurations\"
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/set-dsclocalconfigurationmanager?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/set-dsclocalconfigurationmanager?view=powershell-7.5)


### Set-HVCIOptions

Version: Both

Module: ConfigCI

Syntax:

```powershell
Set-HVCIOptions [-FilePath] <string> [-Enabled] [-Strict] [-DebugMode] [-DisableAllowed] [<CommonParameters>]
Set-HVCIOptions [-FilePath] <string> [-None] [<CommonParameters>]
```

Example: Assign the Strict option

```powershell
PS C:\> Set-HVCIOptions -Strict -FilePath '.\Policy.xml'
PS C:\> Get-Content -Path '.Policy.xml'
    <CiSigner SignerId="ID_SIGNER_S_21" />
  </CiSigners>
  <HvciOptions>2</HvciOptions>
</SiPolicy>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Set-HVCIOptions.md)


### Set-JobTrigger

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Set-JobTrigger [-InputObject] <ScheduledJobTrigger[]> [-DaysInterval <int>] [-WeeksInterval <int>] [-RandomDelay <timespan>] [-At <datetime>] [-User <string>] [-DaysOfWeek <DayOfWeek[]>] [-AtStartup] [-AtLogOn] [-Once] [-RepetitionInterval <timespan>] [-RepetitionDuration <timespan>] [-RepeatIndefinitely] [-Daily] [-Weekly] [-PassThru] [<CommonParameters>]
```

Example (5.1): Change the days in a job trigger

```powershell
Get-JobTrigger -Name "DeployPackage"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Set-JobTrigger.md)


### Set-KdsConfiguration

Version: Both

Module: Kds

Syntax:

```powershell
Set-KdsConfiguration [-LocalTestOnly] [-SecretAgreementPublicKeyLength <int>] [-SecretAgreementPrivateKeyLength <int>] [-SecretAgreementParameters <byte[]>] [-SecretAgreementAlgorithm <string>] [-KdfParameters <byte[]>] [-KdfAlgorithm <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-KdsConfiguration -RevertToDefault [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-KdsConfiguration [-InputObject] <KdsServerConfiguration> [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set the configuration of Microsoft Group Key Distribution Service

```powershell
PS C:\> $config = Get-KdsConfiguration
PS C:\> Set-KdsConfiguration -InputObject $config
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/kds/Set-KdsConfiguration.md)


### Set-LapsADAuditing

Version: Both

Module: LAPS

Syntax:

```powershell
Set-LapsADAuditing -Identity <string[]> -AuditedPrincipals <string[]> [-Credential <pscredential>] [-AuditType <AuditFlags>] [-Domain <string>] [-DomainController <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Set-LapsADAuditing -Identity LapsTestOU -AuditedPrincipals "laps.com\LapsAdmin" -AuditType Success
OU=LapsTestOU,DC=laps,DC=com
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Set-LapsADAuditing.md)


### Set-LapsADComputerSelfPermission

Version: Both

Module: LAPS

Syntax:

```powershell
Set-LapsADComputerSelfPermission -Identity <string[]> [-Domain <string>] [-DomainController <string>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Set-LapsADComputerSelfPermission -Identity LapsTestOU
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Set-LapsADComputerSelfPermission.md)


### Set-LapsADPasswordExpirationTime

Version: Both

Module: LAPS

Syntax:

```powershell
Set-LapsADPasswordExpirationTime -Identity <string[]> [-Credential <pscredential>] [-WhenEffective <datetime>] [-Domain <string>] [-DomainController <string>] [<CommonParameters>]
```

Example: 

```powershell
Set-LapsADPasswordExpirationTime -Identity lapsClient
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Set-LapsADPasswordExpirationTime.md)


### Set-LapsADReadPasswordPermission

Version: Both

Module: LAPS

Syntax:

```powershell
Set-LapsADReadPasswordPermission -Identity <string[]> -AllowedPrincipals <string[]> [-Credential <pscredential>] [-Domain <string>] [-DomainController <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Set-LapsADReadPasswordPermission -Identity LapsTestOU -AllowedPrincipals "Domain Admins"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Set-LapsADReadPasswordPermission.md)


### Set-LapsADResetPasswordPermission

Version: Both

Module: LAPS

Syntax:

```powershell
Set-LapsADResetPasswordPermission -Identity <string[]> -AllowedPrincipals <string[]> [-Credential <pscredential>] [-Domain <string>] [-DomainController <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Set-LapsADResetPasswordPermission -Identity LapsTestOU -AllowedPrincipals "Domain Admins"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Set-LapsADResetPasswordPermission.md)


### Set-LocalGroup

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Set-LocalGroup [-InputObject] <LocalGroup> -Description <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalGroup [-Name] <string> -Description <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalGroup [-SID] <SecurityIdentifier> -Description <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Change a group description

```powershell
Set-LocalGroup -Name "SecurityGroup04" -Description "This is a sample description."
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Set-LocalGroup.md)


### Set-LocalUser

Version: Both

Module: Microsoft.PowerShell.LocalAccounts

Syntax:

```powershell
Set-LocalUser [-Name] <string> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-FullName <string>] [-Password <securestring>] [-PasswordNeverExpires <bool>] [-UserMayChangePassword <bool>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalUser [-InputObject] <LocalUser> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-FullName <string>] [-Password <securestring>] [-PasswordNeverExpires <bool>] [-UserMayChangePassword <bool>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalUser [-SID] <SecurityIdentifier> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-FullName <string>] [-Password <securestring>] [-PasswordNeverExpires <bool>] [-UserMayChangePassword <bool>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Change a description of a user account

```powershell
Set-LocalUser -Name "Admin07" -Description "Description of this account."
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.LocalAccounts/Set-LocalUser.md)


### Set-NonRemovableAppsPolicy

Version: Both

Module: Dism

Syntax:

```powershell
Set-NonRemovableAppsPolicy -PackageFamilyName <string> -NonRemovable <int> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-NonRemovableAppsPolicy -PackageFamilyName <string> -NonRemovable <int> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Set the app package Application1 as non-removable

```powershell
PS> Set-NonRemovableAppsPolicy -Online -PackageFamilyName Application1_1.0.0.0+x64__ms7gsqeatfeb6 -NonRemovable 1
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Set-NonRemovableAppsPolicy.md)


### Set-OsConfigurationDocument

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Set-OsConfigurationDocument [-Content] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [-Wait] [[-TimeoutInSeconds] <int>] [-Update] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Set-OsConfigurationDocument.md)


### Set-OsConfigurationProperty

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Set-OsConfigurationProperty [-Name] <string> [-Value] <string> [[-SourceId] <string>] [[-Id] <string>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> {{ Add example code here }}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/OsConfiguration/Set-OsConfigurationProperty.md)


### Set-OSConfigurationScenarioDefinition

Version: Both

Module: OsConfiguration

Syntax:

```powershell
Set-OsConfigurationScenarioDefinition [-Content] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: none

Source: [OsConfiguration module documentation](https://learn.microsoft.com/en-us/powershell/module/osconfiguration) (no dedicated page).


### Set-ProcessMitigation

Version: Both

Module: ProcessMitigations

Syntax:

```powershell
Set-ProcessMitigation [[-Name] <string>] [-Disable <string[]>] [-Enable <string[]>] [-EAFModules <string[]>] [-Force <string>] [-Reset] [-Remove] [<CommonParameters>]
Set-ProcessMitigation -PolicyFilePath <string> [-IsValid] [<CommonParameters>]
Set-ProcessMitigation [-Disable <string[]>] [-Enable <string[]>] [-EAFModules <string[]>] [-System] [-Force <string>] [-Reset] [-Remove] [<CommonParameters>]
```

Example: 

```powershell
PS C:\>  Set-ProcessMitigation -Name Notepad.exe -Enable SEHOP -Disable ForceRelocateImages
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ProcessMitigations/Set-ProcessMitigation.md)


### Set-PSSessionConfiguration

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Set-PSSessionConfiguration [-Name] <string> [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PSSessionConfiguration [-Name] <string> [-AssemblyName] <string> [-ConfigurationTypeName] <string> [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PSSessionConfiguration [-Name] <string> -Path <string> [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-TransportOption <PSTransportOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Change the thread apartment state

```powershell
PS C:\> Set-PSSessionConfiguration -Name "MaintenanceShell" -ThreadApartmentState STA
```

Example (7): Create and change a session configuration

```powershell
Register-PSSessionConfiguration -Name "AdminShell" -AssemblyName "C:\Shells\AdminShell.dll" -ConfigurationTypeName "AdminClass"
Set-PSSessionConfiguration -Name "AdminShell" -StartupScript "AdminConfig.ps1"
Set-PSSessionConfiguration -Name "AdminShell" -StartupScript $null
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Set-PSSessionConfiguration.md)


### Set-RecoveryManagementPluginAltitude

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint32> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint32> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Set-RecoveryRemoteManagementStatus

Version: Both

Module: Dism

Syntax:

```powershell
Set-RecoveryRemoteManagementStatus -Enabled <bool> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-RecoveryRemoteManagementStatus -Enabled <bool> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Set-ReFSDedupSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Set-ReFSDedupSchedule [-Volume] <string> -Start <datetime> -Days <DaysOfWeek> [-Duration <timespan>] [-CpuPercentage <uint32>] [-ConcurrentOpenFiles <uint32>] [-MinimumLastModifiedTimeHours <int>] [-ExcludeFileExtension <string[]>] [-ExcludeFolder <string[]>] [-CompressionFormat <Format>] [-CompressionLevel <uint16>] [-CompressionChunkSize <uint32>] [-CompressionTuning <uint32>] [-RecompressionTuning <uint32>] [-DecompressionTuning <uint32>] [<CommonParameters>]
```

Example: 

```powershell
Set-ReFSDedupSchedule -Volume "D:" -Start "10:00 PM" -Days Monday,Wednesday,Friday -Duration 4:00:00
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Set-ReFSDedupSchedule.md)


### Set-ReFSDedupScrubSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Set-ReFSDedupScrubSchedule [-Volume] <string> -Start <datetime> -Days <DaysOfWeek> -WeeksInterval <uint16> [-DedupDataOnly <bool>] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    Volume        = "D:"
    Start         = "12/01/2024 8:00 AM"
    Days          = "Monday,Thursday"
    WeeksInterval = 2
}
Set-ReFSDedupScrubSchedule @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Set-ReFSDedupScrubSchedule.md)


### Set-RuleOption

Version: Both

Module: ConfigCI

Syntax:

```powershell
Set-RuleOption [-FilePath] <string> [-Option] <int> [-Delete] [<CommonParameters>]
Set-RuleOption -Help [<CommonParameters>]
```

Example: Remove a rule option

```powershell
The first command displays the contents of the policy. This example shows only the first few lines of the policy, which include the **Rules** property. One of the options displayed is Enabled:Audit Mode.
PS C:\> Get-Content -Path '.Policy.xml'
<?xml version="1.0" encoding="utf-8"?>
<SiPolicy xmlns="urn:schemas-microsoft-com:sipolicy">
  <VersionEx>10.0.0.0</VersionEx>
  <PolicyTypeID>{A244370E-44C9-4C06-B551-F6016E563076}</PolicyTypeID>
  <PlatformID>{2E07F7E4-194C-4D20-B7C9-6F44A6C5A234}</PlatformID>
  <Rules>
    <Rule>
      <Option>Enabled:Unsigned System Integrity Policy</Option>
    </Rule>
    <Rule>
      <Option>Enabled:Audit Mode</Option>
    </Rule>
    <Rule>
      <Option>Enabled:Advanced Boot Options Menu</Option>
    </Rule>
    <Rule>
      <Option>Enabled:UMCI</Option>
    </Rule>
  </Rules>

The second command removes the Enabled:Audit Mode from Policy.xml.The final command displays the contents of the policy again. Enabled:Audit Mode is no longer part of the policy.
PS C:\> Set-RuleOption -FilePath '.\Policy.xml' -Option 3 -Delete
PS C:\> Get-Content -Path '.Policy.xml'
<?xml version="1.0" encoding="utf-8"?>
<SiPolicy xmlns="urn:schemas-microsoft-com:sipolicy">
  <VersionEx>10.0.0.0</VersionEx>
  <PolicyTypeID>{A244370E-44C9-4C06-B551-F6016E563076}</PolicyTypeID>
  <PlatformID>{2E07F7E4-194C-4D20-B7C9-6F44A6C5A234}</PlatformID>
  <Rules>
    <Rule>
      <Option>Enabled:Unsigned System Integrity Policy</Option>
    </Rule>
    <Rule>
      <Option>Enabled:Advanced Boot Options Menu</Option>
    </Rule>
    <Rule>
      <Option>Enabled:UMCI</Option>
    </Rule>
  </Rules>
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/ConfigCI/Set-RuleOption.md)


### Set-ScheduledJob

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Set-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-Name <string>] [-ScriptBlock <scriptblock>] [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-MaxResultCount <int>] [-PassThru] [-ArgumentList <Object[]>] [-RunNow] [-RunEvery <timespan>] [<CommonParameters>]
Set-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-Name <string>] [-FilePath <string>] [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-MaxResultCount <int>] [-PassThru] [-ArgumentList <Object[]>] [-RunNow] [-RunEvery <timespan>] [<CommonParameters>]
Set-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-ClearExecutionHistory] [-PassThru] [<CommonParameters>]
```

Example (5.1): Change the script that a job runs

```powershell
Get-ScheduledJob -Name "Inventory"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Set-ScheduledJob.md)


### Set-ScheduledJobOption

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Set-ScheduledJobOption [-InputObject] <ScheduledJobOptions> [-PassThru] [-RunElevated] [-HideInTaskScheduler] [-RestartOnIdleResume] [-MultipleInstancePolicy <TaskMultipleInstancePolicy>] [-DoNotAllowDemandStart] [-RequireNetwork] [-StopIfGoingOffIdle] [-WakeToRun] [-ContinueIfGoingOnBattery] [-StartIfOnBattery] [-IdleTimeout <timespan>] [-IdleDuration <timespan>] [-StartIfIdle] [<CommonParameters>]
```

Example (5.1): Change job options

```powershell
Get-ScheduledJobOption -Name "DeployPackage"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Set-ScheduledJobOption.md)


### Set-SecureBootUEFI

Version: Both

Module: SecureBoot

Syntax:

```powershell
Set-SecureBootUEFI -Name <string> -Time <string> [-ContentFilePath <string>] [-SignedFilePath <string>] [-AppendWrite] [-OutputFilePath <string>] [<CommonParameters>]
Set-SecureBootUEFI -Name <string> -Time <string> [-Content <byte[]>] [-SignedFilePath <string>] [-AppendWrite] [-OutputFilePath <string>] [<CommonParameters>]
```

Example: Set the DBX UEFI variable

```powershell
PS C:\> $ObjectFromFormat = ( Format-SecureBootUEFI -Name DBX -SignatureOwner 12345678-1234-1234-1234-123456789abc -Algorithm SHA256 -Hash 0011223344556677889900112233445566778899001122334455667788990011 -SignableFilePath GeneratedFileToSign.bin -Time 2011-11-01T13:30:00Z -AppendWrite )
PS C:\> .\signtool.exe sign /fd sha256 /p7 .\ /p7co 1.2.840.113549.1.7.1 /p7ce DetachedSignedData /a /f PrivateKey.pfx GeneratedFileToSign.bin
PS C:\> $ObjectFromFormat | Set-SecureBootUEFI -SignedFilePath GeneratedFileToSign.bin.p7
Name       : dbx
Bytes      : {161, 89, 192, 165...}
Attributes : NON VOLATILE
             BOOTSERVICE ACCESS
             RUNTIME ACCESS
             TIME BASED AUTHENTICATED WRITE ACCESS
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/SecureBoot/Set-SecureBootUEFI.md)


### Set-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax (5.1):

```powershell
Set-Service [-Name] <string> [-ComputerName <string[]>] [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartMode>] [-Status <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Service [-ComputerName <string[]>] [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartMode>] [-Status <string>] [-InputObject <ServiceController>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-Service [-Name] <string> [-DisplayName <string>] [-Credential <pscredential>] [-Description <string>] [-StartupType <ServiceStartupType>] [-SecurityDescriptorSddl <string>] [-Status <string>] [-Force] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Service [-InputObject] <ServiceController> [-DisplayName <string>] [-Credential <pscredential>] [-Description <string>] [-StartupType <ServiceStartupType>] [-SecurityDescriptorSddl <string>] [-Status <string>] [-Force] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Change a display name

```powershell
Set-Service -Name LanmanWorkstation -DisplayName "LanMan Workstation"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-Service.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`systemctl` + sudo).
- Distro: systemd-based + sudo.
- Function: changes service state and startup behavior.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name (.service suffix added automatically) |
| `-Status` | string | `running`/`started` → `systemctl start`; `stopped` → `systemctl stop` |
| `-StartupType` | string | `automatic`/`auto` → `systemctl enable`; `disabled` → `systemctl disable` |

- Implementation: maps to `systemctl start/stop/enable/disable` (sudo where needed).


### Set-SystemPreferredUILanguage

Version: Both

Module: LanguagePackManagement

Syntax:

```powershell
Set-SystemPreferredUILanguage [-Language] <string> [-PassThru] [<CommonParameters>]
```

Example: Set the System Preferred UI Language on a Windows installation

```powershell
Set-SystemPreferredUILanguage ja-JP
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LanguagePackManagement/Set-SystemPreferredUILanguage.md)


### Set-TimeZone

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Set-TimeZone [-Name] <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-TimeZone -Id <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-TimeZone [-InputObject] <TimeZoneInfo> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Set the time zone by Id

```powershell
Set-TimeZone -Id "UTC"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Set-TimeZone.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`sudo timedatectl`).
- Companion: Get-TimeZone.
- Distro: needs systemd's timedatectl + sudo.
- Function: changes the time zone.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Time-zone name (Asia/Shanghai, for instance); handed to `sudo timedatectl set-timezone` |


### Set-TpmOwnerAuth

Version: Both

Module: TrustedPlatformModule

Syntax (5.1):

```powershell
Set-TpmOwnerAuth -File <string> -NewFile <string> [<CommonParameters>]
Set-TpmOwnerAuth -File <string> -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewFile <string> [<CommonParameters>]
```

Syntax (7):

```powershell
Set-TpmOwnerAuth -File <string> -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth -File <string> -NewFile <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewFile <string> [<CommonParameters>]
```

Example: Replace imported owner authorization value

```powershell
PS C:\> Set-TpmOwnerAuth -NewOwnerAuthorization "h4FCmNeWVNp5IMHxRfFL9QEq4vM="
TpmReady           : True
TpmPresent         : True
ManagedAuthLevel   : Full
OwnerAuth          : h4FCmNeWVNp5IMHxRfFL9QEq4vM=
OwnerClearDisabled : True
AutoProvisioning   : DisabledForNextBoot
LockedOut          : False
SelfTest           : {191, 191, 245, 191...}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Set-TpmOwnerAuth.md)


### Set-UevConfiguration

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Set-UevConfiguration [-CurrentComputerUser] [-MaxPackageSizeInBytes <int>] [-SettingsStoragePath <string>] [-EnableSyncProviderPing] [-DisableSyncProviderPing] [-SyncTimeoutInMilliseconds <int>] [-SyncMethod <string>] [-EnableSync] [-DisableSync] [-EnableSyncOverMeteredNetwork] [-DisableSyncOverMeteredNetwork] [-EnableSyncOverMeteredNetworkWhenRoaming] [-DisableSyncOverMeteredNetworkWhenRoaming] [-EnableSettingsImportNotify] [-DisableSettingsImportNotify] [-SettingsImportNotifyDelayInSeconds <int>] [-EnableDontSyncWindows8AppSettings] [-DisableDontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds <int>] [-EnableWaitForSyncOnApplicationStart] [-DisableWaitForSyncOnApplicationStart] [-EnableWaitForSyncOnLogon] [-DisableWaitForSyncOnLogon] [-EnableSyncUnlistedWindows8Apps] [-DisableSyncUnlistedWindows8Apps] [-VdiCollectionName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-UevConfiguration [-Computer] [-MaxPackageSizeInBytes <int>] [-SettingsStoragePath <string>] [-SettingsTemplateCatalogPath <string>] [-EnableSyncProviderPing] [-DisableSyncProviderPing] [-SyncTimeoutInMilliseconds <int>] [-SyncMethod <string>] [-EnableSync] [-DisableSync] [-EnableSyncOverMeteredNetwork] [-DisableSyncOverMeteredNetwork] [-EnableSyncOverMeteredNetworkWhenRoaming] [-DisableSyncOverMeteredNetworkWhenRoaming] [-EnableSettingsImportNotify] [-DisableSettingsImportNotify] [-SettingsImportNotifyDelayInSeconds <int>] [-ContactITUrl <string>] [-ContactITDescription <string>] [-EnableTrayIcon] [-DisableTrayIcon] [-EnableFirstUseNotification] [-DisableFirstUseNotification] [-EnableDontSyncWindows8AppSettings] [-DisableDontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds <int>] [-EnableWaitForSyncOnApplicationStart] [-DisableWaitForSyncOnApplicationStart] [-EnableWaitForSyncOnLogon] [-DisableWaitForSyncOnLogon] [-EnableSyncUnlistedWindows8Apps] [-DisableSyncUnlistedWindows8Apps] [-VdiCollectionName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Modify the synchronization timeout setting for all users

```powershell
PS C:\> Set-UevConfiguration -Computer -SyncTimeoutInMilliseconds 3000
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Set-UevConfiguration.md)


### Set-UevTemplateProfile

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Set-UevTemplateProfile -ID <string> -Profile <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Associate a template with the Backup profile

```powershell
PS C:\>Set-UevTemplateProfile -ID "MicrosoftCalculator6" -Profile "Backup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Set-UevTemplateProfile.md)


### Set-WheaMemoryPolicy

Version: Both

Module: Whea

Syntax (5.1):

```powershell
Set-WheaMemoryPolicy [-ComputerName <string>] [-DisableOffline <bool>] [-DisablePFA <bool>] [-PersistMemoryOffline <bool>] [-PFAPageCount <uint32>] [-PFAErrorThreshold <uint32>] [-PFATimeout <uint32>] [<CommonParameters>]
```

Syntax (7):

```powershell
Set-WheaMemoryPolicy [-ComputerName <string>] [-DisableOffline <bool>] [-DisablePFA <bool>] [-PersistMemoryOffline <bool>] [-PFAPageCount <uint>] [-PFAErrorThreshold <uint>] [-PFATimeout <uint>] [<CommonParameters>]
```

Example: Enable WHEA predictive failure analysis

```powershell
PS C:\> Set-WheaMemoryPolicy -DisablePFA $False
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Whea/Set-WheaMemoryPolicy.md)


### Set-WinAcceptLanguageFromLanguageListOptOut

Version: Both

Module: International

Syntax:

```powershell
Set-WinAcceptLanguageFromLanguageListOptOut [-OptOut] <bool> [<CommonParameters>]
```

Example: Update the registry key

```powershell
PS C:\> Set-WinAcceptLanguageFromLanguageListOptOut -OptOut $True
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinAcceptLanguageFromLanguageListOptOut.md)


### Set-WinCultureFromLanguageListOptOut

Version: Both

Module: International

Syntax:

```powershell
Set-WinCultureFromLanguageListOptOut [-OptOut] <bool> [<CommonParameters>]
```

Example: Block dynamic setting

```powershell
PS C:\> Set-WinCultureFromLanguageListOptOut -OptOut $True
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinCultureFromLanguageListOptOut.md)


### Set-WinDefaultInputMethodOverride

Version: Both

Module: International

Syntax:

```powershell
Set-WinDefaultInputMethodOverride [[-InputTip] <string>] [<CommonParameters>]
```

Example: Set the default input method override

```powershell
PS C:\> Set-WinDefaultInputMethodOverride -InputTip "0409:00000409"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinDefaultInputMethodOverride.md)


### Set-WindowsEdition

Version: Both

Module: Dism

Syntax:

```powershell
Set-WindowsEdition -Edition <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Change the edition of an image

```powershell
PS C:\> Set-WindowsEdition -Path "c:\offline" -Edition "Ultimate"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Set-WindowsEdition.md)


### Set-WindowsProductKey

Version: Both

Module: Dism

Syntax:

```powershell
Set-WindowsProductKey -ProductKey <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Set the product key on a mounted image

```powershell
PS C:\> Set-WindowsProductKey -Path "c:\offline" -ProductKey "12345-12345-12345-12345-12345"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Set-WindowsProductKey.md)


### Set-WindowsReservedStorageState

Version: Both

Module: Dism

Syntax:

```powershell
Set-WindowsReservedStorageState -State <ReservedStorageState> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: 

```powershell
PS C:\> Set-WindowsReservedStorageState -State Enabled -Online
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Set-WindowsReservedStorageState.md)


### Set-WindowsSearchSetting

Version: Both

Module: WindowsSearch

Syntax:

```powershell
Set-WindowsSearchSetting [-EnableWebResultsSetting <bool>] [-EnableMeteredWebResultsSetting <bool>] [-SearchExperienceSetting <string>] [-SafeSearchSetting <string>] [<CommonParameters>]
```

Example: Personalize Windows Search

```powershell
PS C:\> Set-WindowsSearchSetting -SearchExperienceSetting "Personalized"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsSearch/Set-WindowsSearchSetting.md)


### Set-WinHomeLocation

Version: Both

Module: International

Syntax:

```powershell
Set-WinHomeLocation [-GeoId] <int> [<CommonParameters>]
```

Example: Set the home location

```powershell
PS C:\> Set-WinHomeLocation -GeoId 0xF4
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinHomeLocation.md)


### Set-WinLanguageBarOption

Version: Both

Module: International

Syntax:

```powershell
Set-WinLanguageBarOption [-UseLegacySwitchMode] [-UseLegacyLanguageBar] [<CommonParameters>]
```

Example: Set language bar options

```powershell
PS C:\> Set-WinLanguageBarOption -UseLegacySwitchMode -UseLegacyLanguageBar
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinLanguageBarOption.md)


### Set-WinSystemLocale

Version: Both

Module: International

Syntax:

```powershell
Set-WinSystemLocale [-SystemLocale] <cultureinfo> [<CommonParameters>]
```

Example: Set the system locale

```powershell
PS C:\> Set-WinSystemLocale -SystemLocale ja-JP
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinSystemLocale.md)


### Set-WinUILanguageOverride

Version: Both

Module: International

Syntax:

```powershell
Set-WinUILanguageOverride [[-Language] <cultureinfo>] [<CommonParameters>]
```

Example: Set a language override

```powershell
PS C:\> Set-WinUILanguageOverride -Language de-DE
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinUILanguageOverride.md)


### Set-WinUserLanguageList

Version: Both

Module: International

Syntax:

```powershell
Set-WinUserLanguageList [-LanguageList] <List[WinUserLanguage]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Add a language to an existing language list

```powershell
PS C:\> $OldList = Get-WinUserLanguageList
PS C:\> $OldList.Add("fr-FR")
PS C:\> Set-WinUserLanguageList -LanguageList $OldList
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/International/Set-WinUserLanguageList.md)


### Set-WmiInstance

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Set-WmiInstance [-Class] <string> [[-Arguments] <hashtable>] [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance -InputObject <wmi> [-Arguments <hashtable>] [-PutType <PutType>] [-AsJob] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance -Path <string> [-Arguments <hashtable>] [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Set WMI logging level

```powershell
Set-WmiInstance -Class Win32_WMISetting -Arguments @{LoggingLevel=2}
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Set-WmiInstance.md)


### Set-WSManInstance

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Set-WSManInstance [-ResourceURI] <uri> [[-SelectorSet] <hashtable>] [-ApplicationName <string>] [-ComputerName <string>] [-Dialect <uri>] [-FilePath <string>] [-Fragment <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Set-WSManInstance [-ResourceURI] <uri> [[-SelectorSet] <hashtable>] [-ConnectionURI <uri>] [-Dialect <uri>] [-FilePath <string>] [-Fragment <string>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Disable a listener on the local computer

```powershell
$params = @{
    ResourceURI = 'winrm/config/listener'
    SelectorSet = @{address = '*'; transport = 'https'}
    ValueSet    = @{Enabled = 'false'}
}
Set-WSManInstance @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Set-WSManInstance.md)


### Set-WSManQuickConfig

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Set-WSManQuickConfig [-UseSSL] [-Force] [-SkipNetworkProfileCheck] [<CommonParameters>]
```

Example: Enable remote management of the local computer over HTTP

```powershell
Set-WSManQuickConfig
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Set-WSManQuickConfig.md)


### Show-Command

Version: Both

Module: Microsoft.PowerShell.Utility

Syntax:

```powershell
Show-Command [[-Name] <string>] [-Height <double>] [-Width <double>] [-NoCommonParameter] [-ErrorPopup] [-PassThru] [<CommonParameters>]
```

Example: Open the Commands window

```powershell
Show-Command
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Utility/Show-Command.md)


### Show-ControlPanelItem

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Show-ControlPanelItem [-Name] <string[]> [<CommonParameters>]
Show-ControlPanelItem -CanonicalName <string[]> [<CommonParameters>]
Show-ControlPanelItem [[-InputObject] <ControlPanelItem[]>] [<CommonParameters>]
```

Example (5.1): Show a control panel item

```powershell
Show-ControlPanelItem -Name "AutoPlay"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Show-ControlPanelItem.md)


### Show-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Show-EventLog [[-ComputerName] <string>] [<CommonParameters>]
```

Example (5.1): Display event logs for the local computer

```powershell
Show-EventLog
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Show-EventLog.md)


### Show-WindowsDeveloperLicenseRegistration

Version: Both

Module: WindowsDeveloperLicense

Syntax:

```powershell
Show-WindowsDeveloperLicenseRegistration [<CommonParameters>]
```

Example: Enable your device for development

```powershell
PS C:\> Show-WindowsDeveloperLicenseRegistration
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsDeveloperLicense/Show-WindowsDeveloperLicenseRegistration.md)


### Split-WindowsImage

Version: Both

Module: Dism

Syntax (5.1):

```powershell
Split-WindowsImage -ImagePath <string> -SplitImagePath <string> -FileSize <uint64> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Syntax (7):

```powershell
Split-WindowsImage -ImagePath <string> -SplitImagePath <string> -FileSize <ulong> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Split a .wim file

```powershell
PS C:\> Split-WindowsImage -ImagePath "c:\imagestore\install.wim" -SplitImagePath "c:\imagestore\splitfiles\split.swm" -FileSize 1024 -CheckIntegrity
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Split-WindowsImage.md)


### Start-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Start-BitsTransfer [-Source] <string[]> [[-Destination] <string[]>] [-Asynchronous] [-Dynamic] [-CustomHeadersWriteOnly] [-Authentication <string>] [-Credential <pscredential>] [-Description <string>] [-HttpMethod <string>] [-DisplayName <string>] [-Priority <string>] [-TransferPolicy <CostStates>] [-ACLFlags <ACLFlagValue>] [-SecurityFlags <SecurityFlagValue>] [-UseStoredCredential <AuthenticationTargetValue>] [-ProxyAuthentication <string>] [-ProxyBypass <string[]>] [-ProxyCredential <pscredential>] [-ProxyList <uri[]>] [-ProxyUsage <string>] [-RetryInterval <int>] [-RetryTimeout <int>] [-MaxDownloadTime <int>] [-Suspended] [-TransferType <string>] [-CustomHeaders <string[]>] [-NotifyFlags <NotifyFlagValue>] [-NotifyCmdLine <string[]>] [-CertStoreLocation <CertStoreLocationValue>] [-CertStoreName <string>] [-CertHash <byte[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create a BITS transfer job that downloads a file

```powershell
PS C:\> Start-BitsTransfer -Source "http://server01/servertestdir/testfile1.txt" -Destination "c:\clienttestdir\testfile1.txt"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Start-BitsTransfer.md)


### Start-DscConfiguration

Version: Both

Module: PSDesiredStateConfiguration

Syntax:

```powershell
Start-DscConfiguration [[-Path] <string>] [[-ComputerName] <string[]>] [-Wait] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-DscConfiguration [[-Path] <string>] -CimSession <CimSession[]> [-Wait] [-Force] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-DscConfiguration [[-ComputerName] <string[]>] -UseExisting [-Wait] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-DscConfiguration -CimSession <CimSession[]> -UseExisting [-Wait] [-Force] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Apply configuration settings

```powershell
Start-DscConfiguration -Path "C:\DSC\Configurations\"
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/start-dscconfiguration?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/start-dscconfiguration?view=powershell-7.5)


### Start-DtcDiagnosticResourceManager

Version: Both

Module: MsDtc

Syntax:

```powershell
Start-DtcDiagnosticResourceManager [[-Port] <int>] [[-Name] <string>] [<CommonParameters>]
```

Example: Start a diagnostic resource manager

```powershell
PS C:\> Start-DtcDiagnosticResourceManager -Port 17124 -Name "testRM"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Start-DtcDiagnosticResourceManager.md)


### Start-OSUninstall

Version: Both

Module: Dism

Syntax:

```powershell
Start-OSUninstall -Path <string> [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Start-OSUninstall -Online [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Uninstall OS

```powershell
Start-OSUninstall -Online
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Start-OSUninstall.md)


### Start-ReFSDedupJob

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Start-ReFSDedupJob [-Volume] <string> [-Duration <timespan>] [-FullRun] [-CpuPercentage <uint32>] [-ConcurrentOpenFiles <uint32>] [-MinimumLastModifiedTimeHours <int>] [-ExcludeFileExtension <string[]>] [-ExcludeFolder <string[]>] [-CompressionFormat <Format>] [-CompressionLevel <uint16>] [-CompressionChunkSize <uint32>] [-CompressionTuning <uint32>] [-RecompressionTuning <uint32>] [-DecompressionTuning <uint32>] [<CommonParameters>]
```

Example: 

```powershell
Start-ReFSDedupJob -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Start-ReFSDedupJob.md)


### Start-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Start-Service [-InputObject] <ServiceController[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Service [-Name] <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Service -DisplayName <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Start a service by using its name

```powershell
Start-Service -Name "eventlog"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Start-Service.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`systemctl` + sudo).
- Companions: Stop-Service, Restart-Service, Resume-Service.
- Distro: systemd-based + sudo.
- Function: starts/stops/restarts services.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name; handed to `systemctl <action> <unit>` (.service suffix added automatically) |

- Implementation: `systemctl start/stop/restart <unit>`; a failure under ordinary permissions retries automatically with `sudo`. Counterpart: `sudo systemctl start/stop/restart`.
- Stop-Service, Restart-Service, and Resume-Service share Start-Service's parameters, their actions being stop/restart/start.


### Start-Transaction

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Start-Transaction [-Timeout <int>] [-Independent] [-RollbackPreference <RollbackSeverity>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Start and roll back a transaction

```powershell
Set-Location HKCU:\software
Start-Transaction
New-Item "ContosoCompany" -UseTransaction
New-ItemProperty "ContosoCompany" -Name "MyKey" -Value 123 -UseTransaction
Undo-Transaction
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Start-Transaction.md)


### Stop-AppvClientConnectionGroup

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Stop-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [<CommonParameters>]
Stop-AppvClientConnectionGroup [-Name] <string> [-Global] [<CommonParameters>]
Stop-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [<CommonParameters>]
```

Example: Stop a virtual environment for a named group

```powershell
PS C:\> Stop-AppvClientConnectionGroup -Name "MyGroup"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Stop-AppvClientConnectionGroup.md)


### Stop-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Stop-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Global] [<CommonParameters>]
Stop-AppvClientPackage [-Package] <AppvClientPackage> [-Global] [<CommonParameters>]
Stop-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [<CommonParameters>]
```

Example: Shut down a virtual environment for a version of a package

```powershell
PS C:\> Stop-AppvClientPackage -Name "MyPackage" -Version 2
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Stop-AppvClientPackage.md)


### Stop-DtcDiagnosticResourceManager

Version: Both

Module: MsDtc

Syntax:

```powershell
Stop-DtcDiagnosticResourceManager [[-Job] <DtcDiagnosticResourceManagerJob>] [<CommonParameters>]
Stop-DtcDiagnosticResourceManager [[-Name] <string>] [<CommonParameters>]
Stop-DtcDiagnosticResourceManager [[-InstanceId] <guid>] [<CommonParameters>]
```

Example: Stop a diagnostic resource manager

```powershell
PS C:\> Stop-DtcDiagnosticResourceManager -Name "testRM"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Stop-DtcDiagnosticResourceManager.md)


### Stop-ReFSDedupJob

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Stop-ReFSDedupJob [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Stop-ReFSDedupJob -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Stop-ReFSDedupJob.md)


### Stop-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Stop-Service [-InputObject] <ServiceController[]> [-Force] [-NoWait] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Service [-Name] <string[]> [-Force] [-NoWait] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Service -DisplayName <string[]> [-Force] [-NoWait] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Stop a service on the local computer

```powershell
Stop-Service -Name "sysmonlog"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Stop-Service.md)

#### Implementation in PowerShell For Linux:

- Type: mapped Linux command (`systemctl` + sudo).
- Companions: Stop-Service, Restart-Service, Resume-Service.
- Distro: systemd-based + sudo.
- Function: starts/stops/restarts services.

| Parameter | Type | Mapping / notes |
| :--- | :--- | :--- |
| `-Name` (position 0) | string | Service name; handed to `systemctl <action> <unit>` (.service suffix added automatically) |

- Implementation: `systemctl start/stop/restart <unit>`; a failure under ordinary permissions retries automatically with `sudo`. Counterpart: `sudo systemctl start/stop/restart`.
- Stop-Service, Restart-Service, and Resume-Service share Start-Service's parameters, their actions being stop/restart/start.


### Suspend-BitsTransfer

Version: Both

Module: BitsTransfer

Syntax:

```powershell
Suspend-BitsTransfer [-BitsJob] <BitsJob[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Suspend all BITS transfer jobs owned by the current user

```powershell
PS C:\> Get-BitsTransfer | Suspend-BitsTransfer
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/BitsTransfer/Suspend-BitsTransfer.md)


### Suspend-Job

Version: 5.1 only

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Suspend-Job [-Id] <int[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-Job] <Job[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-InstanceId] <guid[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-Name] <string[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-Filter] <hashtable> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-State] <JobState> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Suspend a workflow job by name

```powershell
#Sample Workflow
workflow Get-SystemLog
{
    $Events = Get-WinEvent -LogName System
    CheckPoint-Workflow
    inlinescript {\\Server01\Scripts\Analyze-SystemEvents.ps1 -Events $Events}
}
Get-SystemLog -AsJob -JobName "LogflowJob"
Get-Job -Name LogflowJob
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Core/Suspend-Job.md)


### Suspend-ReFSDedupSchedule

Version: Both

Module: Microsoft.ReFsDedup.Commands

Syntax:

```powershell
Suspend-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

Example: 

```powershell
Suspend-ReFSDedupSchedule -Volume "D:"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Microsoft.ReFsDedup.Commands/Suspend-ReFSDedupSchedule.md)


### Suspend-Service

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Suspend-Service [-InputObject] <ServiceController[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Service [-Name] <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Service -DisplayName <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Suspend a service

```powershell
Suspend-Service -DisplayName "Telnet"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Management/Suspend-Service.md)


### Switch-Certificate

Version: Both

Module: PKI

Syntax:

```powershell
Switch-Certificate [-OldCert] <Certificate> [-NewCert] <Certificate> [-NotifyOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
$params = @{
    OldCert = 'Cert:\LocalMachine\My\E42DBC3B3F2771990A9B3E35D0C3C422779DACD7'
    NewCert = 'Cert:\LocalMachine\My\4A346B4385F139CA843912D358D765AB8DEE9FD4'
}
Switch-Certificate @params
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Switch-Certificate.md)


### Sync-AppvPublishingServer

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Sync-AppvPublishingServer [-ServerId] <uint32> [-Global] [-Force] [-NetworkCostAware] [-HidePublishingRefreshUI] [<CommonParameters>]
Sync-AppvPublishingServer [-Server] <AppvPublishingServer> [-Global] [-Force] [-NetworkCostAware] [-HidePublishingRefreshUI] [<CommonParameters>]
Sync-AppvPublishingServer [[-Name] <string>] [[-URL] <string>] [-Global] [-Force] [-NetworkCostAware] [-HidePublishingRefreshUI] [<CommonParameters>]
Sync-AppvPublishingServer [-PublishFromXML] [-Global] [-NetworkCostAware] [<CommonParameters>]
```

Example: Start publishing refresh

```powershell
PS C:\> Sync-AppvPublishingServer -Name "MyServer"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Sync-AppvPublishingServer.md)


### Test-AppLockerPolicy

Version: Both

Module: AppLocker

Syntax:

```powershell
Test-AppLockerPolicy [-XmlPolicy] <string> -Path <List[string]> [-User <string>] [-Filter <List[PolicyDecision]>] [<CommonParameters>]
Test-AppLockerPolicy [-XmlPolicy] <string> -Packages <List[AppxPackage]> [-User <string>] [-Filter <List[PolicyDecision]>] [<CommonParameters>]
Test-AppLockerPolicy [-PolicyObject] <AppLockerPolicy> -Path <List[string]> [-User <string>] [-Filter <List[PolicyDecision]>] [<CommonParameters>]
```

Example: Report if programs are allowed to run

```powershell
PS C:\> Test-AppLockerPolicy -XMLPolicy C:\Policy.xml -Path c:\windows\system32\calc.exe, C:\windows\system32\notepad.exe -User Everyone
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppLocker/Test-AppLockerPolicy.md)


### Test-Certificate

Version: Both

Module: PKI

Syntax:

```powershell
Test-Certificate [-Cert] <Certificate> [-Policy <TestCertificatePolicy>] [-User] [-EKU <string[]>] [-DNSName <string>] [-AllowUntrustedRoot] [<CommonParameters>]
```

Example: 

```powershell
Get-ChildItem -Path Cert:\LocalMachine\My |
    Test-Certificate -Policy SSL -DNSName 'dns=contoso.com'
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/pki/Test-Certificate.md)


### Test-ComputerSecureChannel

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Test-ComputerSecureChannel [-Repair] [-Server <string>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Test a channel between the local computer and its domain

```powershell
Test-ComputerSecureChannel
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Test-ComputerSecureChannel.md)


### Test-DscConfiguration

Version: Both

Module: PSDesiredStateConfiguration

Syntax:

```powershell
Test-DscConfiguration [[-ComputerName] <string[]>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-AsJob] [-Detailed] [<CommonParameters>]
Test-DscConfiguration [-Path] <string> [[-ComputerName] <string[]>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration [[-ComputerName] <string[]>] -ReferenceConfiguration <string> [-Credential <pscredential>] [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration [-Path] <string> -CimSession <CimSession[]> [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration -CimSession <CimSession[]> -ReferenceConfiguration <string> [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration -CimSession <CimSession[]> [-ThrottleLimit <int>] [-AsJob] [-Detailed] [<CommonParameters>]
```

Example: Test configuration for the local computer

```powershell
Test-DscConfiguration
```

Source: [Official English docs (5.1)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/test-dscconfiguration?view=powershell-5.1) / [Official English docs (7)](https://learn.microsoft.com/en-us/powershell/module/psdesiredstateconfiguration/test-dscconfiguration?view=powershell-7.5)


### Test-FileCatalog

Version: Both

Module: Microsoft.PowerShell.Security

Syntax:

```powershell
Test-FileCatalog [-CatalogFilePath] <string> [[-Path] <string[]>] [-Detailed] [-FilesToSkip <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Create and validate a file catalog

```powershell
$NewFileCatalogParams = @{
    Path = "$PSHOME\Modules\Microsoft.PowerShell.Utility"
    CatalogFilePath = "\temp\Microsoft.PowerShell.Utility.cat"
    CatalogVersion = 2.0
}
New-FileCatalog @NewFileCatalogParams

$TestFileCatalogParams = @{
    CatalogFilePath = "\temp\Microsoft.PowerShell.Utility.cat"
    Path = "$PSHOME\Modules\Microsoft.PowerShell.Utility\"
}
Test-FileCatalog @TestFileCatalogParams
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Security/Test-FileCatalog.md)


### Test-KdsRootKey

Version: Both

Module: Kds

Syntax:

```powershell
Test-KdsRootKey [-KeyId] <guid> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Test the root key configuration

```powershell
PS C:\> Test-KdsRootKey -KeyId 4A3615F1-5A90-22E4-0B1D-1416F93D4412
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/kds/Test-KdsRootKey.md)


### Test-PSSessionConfigurationFile

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Test-PSSessionConfigurationFile [-Path] <string> [<CommonParameters>]
```

Example: Test a session configuration file

```powershell
Test-PSSessionConfigurationFile -Path "FullLanguage.pssc"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Test-PSSessionConfigurationFile.md)


### Test-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Test-UevTemplate [-Path] <string[]> [<CommonParameters>]
Test-UevTemplate -LiteralPath <string[]> [<CommonParameters>]
```

Example: Test a file

```powershell
PS C:\> Test-UevTemplate -Path "MicrosoftWordpad.xml" | Format-Table -AutoSize
Path                                                                                     Status Message
----                                                                                     ------ -------
C:\Program Files\Microsoft User Experience Virtualization\Templates\MicrosoftWordpad.xml Valid  The template is valid.
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Test-UevTemplate.md)


### Test-WSMan

Version: Both

Module: Microsoft.WSMan.Management

Syntax:

```powershell
Test-WSMan [[-ComputerName] <string>] [-Authentication <AuthenticationMechanism>] [-Port <int>] [-UseSSL] [-ApplicationName <string>] [-Credential <pscredential>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

Example: Determine the status of the WinRM service

```powershell
Test-WSMan
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.WSMan.Management/Test-WSMan.md)


### Unblock-Tpm

Version: Both

Module: TrustedPlatformModule

Syntax:

```powershell
Unblock-Tpm [[-OwnerAuthorization] <string>] [<CommonParameters>]
Unblock-Tpm -File <string> [<CommonParameters>]
```

Example: Reset a lockout

```powershell
PS C:\>Unblock-Tpm -OwnerAuthorization "vjnuW6rToM41os3xxEpjLdIW2gA="
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/TrustedPlatformModule/Unblock-Tpm.md)


### Undo-DtcDiagnosticTransaction

Version: Both

Module: MsDtc

Syntax:

```powershell
Undo-DtcDiagnosticTransaction [-Transaction] <DtcDiagnosticTransaction> [<CommonParameters>]
```

Example: Undo a DTC diagnostic transaction

```powershell
PS C:\> $Tx = New-DtcDiagnosticTransaction
PS C:\> Undo-DtcDiagnosticTransaction -Transaction $Tx
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/MsDtc/Undo-DtcDiagnosticTransaction.md)


### Undo-Transaction

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Undo-Transaction [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Roll back the current transaction

```powershell
Undo-Transaction
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Undo-Transaction.md)


### Uninstall-Language

Version: Both

Module: LanguagePackManagement

Syntax:

```powershell
Uninstall-Language [-Language] <string> [-PassThru] [<CommonParameters>]
```

Example: Remove a language from a device

```powershell
Uninstall-Language ja-jp
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LanguagePackManagement/Uninstall-Language.md)


### Uninstall-ProvisioningPackage

Version: Both

Module: Provisioning

Syntax:

```powershell
Uninstall-ProvisioningPackage [-PackageId] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Uninstall-ProvisioningPackage -PackagePath <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Uninstall-ProvisioningPackage -AllInstalledPackages [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Uninstall-ProvisioningPackage [-RuntimeMetadata] <RuntimeProvPackageMetadata> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: Uninstall all provisioning packages

```powershell
PS C:\> Uninstall-ProvisioningPackage -AllInstalledPackages
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Uninstall-ProvisioningPackage.md)


### Uninstall-TrustedProvisioningCertificate

Version: Both

Module: Provisioning

Syntax:

```powershell
Uninstall-TrustedProvisioningCertificate [-Thumbprint] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

Example: Uninstall a trusted provisioning certificate

```powershell
PS C:\> Uninstall-TrustedProvisioningCertificate -Thumbprint fedd995b45e633d4ef30fcbc8f3a48b627e9a28b
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Provisioning/Uninstall-TrustedProvisioningCertificate.md)


### Unpublish-AppvClientPackage

Version: 5.1 only

Module: AppvClient

Syntax:

```powershell
Unpublish-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Global] [-UserSID <string>] [<CommonParameters>]
Unpublish-AppvClientPackage [-Package] <AppvClientPackage> [-Global] [-UserSID <string>] [<CommonParameters>]
Unpublish-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [-UserSID <string>] [<CommonParameters>]
```

Example: Unpublish a version of a package

```powershell
PS C:\> Unpublish-AppvClientPackage -Name "MyApp" -Version 3
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/AppvClient/Unpublish-AppvClientPackage.md)


### Unregister-PSSessionConfiguration

Version: Both

Module: Microsoft.PowerShell.Core

Syntax:

```powershell
Unregister-PSSessionConfiguration [-Name] <string> [-Force] [-NoServiceRestart] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Delete a session configuration

```powershell
Unregister-PSSessionConfiguration -Name "MaintenanceShell"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/7.5/Microsoft.PowerShell.Core/Unregister-PSSessionConfiguration.md)


### Unregister-RecoveryManagementPlugin

Version: Both

Module: Dism

Syntax:

```powershell
Unregister-RecoveryManagementPlugin -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Unregister-RecoveryManagementPlugin -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: none

Source: [Dism module documentation](https://learn.microsoft.com/en-us/powershell/module/dism) (no dedicated page).


### Unregister-ScheduledJob

Version: Both

Module: PSScheduledJob

Syntax:

```powershell
Unregister-ScheduledJob [-InputObject] <ScheduledJobDefinition[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-ScheduledJob [-Id] <int[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-ScheduledJob [-Name] <string[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example (5.1): Delete a scheduled job

```powershell
Unregister-ScheduledJob TestJob
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/PSScheduledJob/Unregister-ScheduledJob.md)


### Unregister-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Unregister-UevTemplate [-ID] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-UevTemplate -All [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Unregister a template

```powershell
PS C:\> Unregister-UevTemplate -TemplateId "MicrosoftCalculator6"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Unregister-UevTemplate.md)


### Unregister-WindowsDeveloperLicense

Version: Both

Module: WindowsDeveloperLicense

Syntax:

```powershell
Unregister-WindowsDeveloperLicense [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Disable Developer Mode

```powershell
PS C:\> Unregister-WindowsDeveloperLicense
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/WindowsDeveloperLicense/Unregister-WindowsDeveloperLicense.md)


### Update-LapsADSchema

Version: Both

Module: LAPS

Syntax:

```powershell
Update-LapsADSchema [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: 

```powershell
Update-LapsADSchema
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/LAPS/Update-LapsADSchema.md)


### Update-UevTemplate

Version: 5.1 only

Module: UEV

Syntax:

```powershell
Update-UevTemplate [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Update-UevTemplate -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

Example: Update a template

```powershell
PS C:\> Update-UevTemplate -Path "MicrosoftCalculator.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/UEV/Update-UevTemplate.md)


### Update-WIMBootEntry

Version: Both

Module: Dism

Syntax:

```powershell
Update-WIMBootEntry -Path <string> -ImagePath <string> -DataSourceID <long> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Update the configuration entry for a .wim file

```powershell
PS C:\> Update-WIMBootEntry -Path "C:\" -DataSourceID 0 -ImagePath "D:\Windows Images\install.wim"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Update-WIMBootEntry.md)


### Use-Transaction

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Use-Transaction [-TransactedScript] <scriptblock> [-UseTransaction] [<CommonParameters>]
```

Example (5.1): Script by using a transaction-enabled object

```powershell
Start-Transaction
$transactedString = New-Object Microsoft.PowerShell.Commands.Management.TransactedString
$transactedString.Append("Hello")
Use-Transaction -TransactedScript { $transactedString.Append(", World") } -UseTransaction
$transactedString.ToString()
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Use-Transaction.md)


### Use-WindowsUnattend

Version: Both

Module: Dism

Syntax:

```powershell
Use-WindowsUnattend -UnattendPath <string> -Path <string> [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Use-WindowsUnattend -UnattendPath <string> -Online [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

Example: Apply an answer file to a mounted image

```powershell
PS C:\> Use-WindowsUnattend -Path "c:\offline" -UnattendPath "c:\test\answerfiles\myunattend.xml"
```

Source: [Official reference source](https://github.com/MicrosoftDocs/windows-powershell-docs/blob/main/docset/winserver2025-ps/Dism/Use-WindowsUnattend.md)


### Write-EventLog

Version: Both

Module: Microsoft.PowerShell.Management

Syntax:

```powershell
Write-EventLog [-LogName] <string> [-Source] <string> [-EventId] <int> [[-EntryType] <EventLogEntryType>] [-Message] <string> [-Category <int16>] [-RawData <byte[]>] [-ComputerName <string>] [<CommonParameters>]
```

Example (5.1): Write an event to the Application event log

```powershell
PS C:\> Write-EventLog -LogName "Application" -Source "MyApp" -EventID 3001 -EntryType Information -Message "MyApp added a user-requested feature to the display." -Category 1 -RawData 10,20
```

Source: [Official reference source](https://github.com/MicrosoftDocs/PowerShell-Docs/blob/main/reference/5.1/Microsoft.PowerShell.Management/Write-EventLog.md)


