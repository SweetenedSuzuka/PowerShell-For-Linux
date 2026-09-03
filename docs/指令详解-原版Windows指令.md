# 指令详解-原版Windows指令

原版 PowerShell 中仅 Windows 可用的指令。
PowerShell For Linux默认不实现它们，但有时候会因为特殊原因实现一部分，会在下面标注。

实现状态说明：
- **Go实现**：行为在本程序内部用 Go 复现，不调用外部命令。
- **映射 Linux**：委托给对应的 Linux 命令/工具，括号内为所调工具。
- **未实现**：本程序尚未实现。
- **不实现**：超出范围，本程序不做，原因见备注。

## 指令列表

| 指令 | 模块 | 版本 | 差异 | 用途 | 实现状态 | 备注 |
|---|---|---|---|---|---|---|
| [`Add-AppProvisionedSharedPackageContainer`](#add-appprovisionedsharedpackagecontainer) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Add-AppSharedPackageContainer`](#add-appsharedpackagecontainer) | Appx | 都有 | 无 | 部署共享包容器定义。 | 不实现 |  |
| [`Add-AppvClientConnectionGroup`](#add-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 创建多个包的组合。 | 不实现 |  |
| [`Add-AppvClientPackage`](#add-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 向运行 App-V 客户端的计算机添加包。 | 不实现 |  |
| [`Add-AppvPublishingServer`](#add-appvpublishingserver) | AppvClient | 仅5.1 | 仅5.1提供 | 为运行 App-V 客户端的计算机添加发布服务器。 | 不实现 |  |
| [`Add-AppxPackage`](#add-appxpackage) | Appx | 都有 | 无 | 向用户帐户添加已签名的应用包。 | 不实现 |  |
| [`Add-AppxProvisionedPackage`](#add-appxprovisionedpackage) | Dism | 都有 | 语法不同 | 向 Windows 映像添加应用包 (.appx)，使其为每个新用户安装。 | 不实现 |  |
| [`Add-AppxVolume`](#add-appxvolume) | Appx | 都有 | 无 | 向包管理器添加 appx 卷。 | 不实现 |  |
| [`Add-BitsFile`](#add-bitsfile) | BitsTransfer | 都有 | 无 | 向现有 BITS 传输作业添加一个或多个文件。 | 不实现 |  |
| [`Add-CertificateEnrollmentPolicyServer`](#add-certificateenrollmentpolicyserver) | PKI | 都有 | 无 | 向当前用户或本地系统配置添加注册策略服务器。 | 不实现 |  |
| [`Add-Computer`](#add-computer) | Microsoft.PowerShell.Management | 都有 | 无 | 将本地计算机添加到域或工作组中。 | 不实现 |  |
| [`Add-JobTrigger`](#add-jobtrigger) | PSScheduledJob | 都有 | 无 | 将作业触发器添加到计划作业。 | 不实现 |  |
| [`Add-KdsRootKey`](#add-kdsrootkey) | Kds | 都有 | 无 | 为 AD 中的 KdsSvc 生成新根密钥。 | 不实现 |  |
| [`Add-LocalGroupMember`](#add-localgroupmember) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 将成员添加到本地组。 | 不实现 |  |
| [`Add-PSSnapin`](#add-pssnapin) | Microsoft.PowerShell.Core | 仅5.1 | 仅5.1提供 | 将一个或多个 Windows PowerShell 管理单元添加到当前会话。 | 不实现 |  |
| [`Add-SignerRule`](#add-signerrule) | ConfigCI | 都有 | 无 | 创建签名者规则并添加到策略。 | 不实现 |  |
| [`Add-WindowsCapability`](#add-windowscapability) | Dism | 都有 | 无 | 在指定操作系统映像安装 Windows 功能包。 | 不实现 |  |
| [`Add-WindowsDriver`](#add-windowsdriver) | Dism | 都有 | 无 | 向脱机 Windows 映像添加驱动程序。 | 不实现 |  |
| [`Add-WindowsImage`](#add-windowsimage) | Dism | 都有 | 无 | 向现有映像 (.wim) 文件添加附加映像。 | 不实现 |  |
| [`Add-WindowsPackage`](#add-windowspackage) | Dism | 都有 | 无 | 向 Windows 映像添加单个 .cab 或 .msu 文件。 | 不实现 |  |
| [`Checkpoint-Computer`](#checkpoint-computer) | Microsoft.PowerShell.Management | 都有 | 无 | 在本地计算机上创建系统还原点。 | 不实现 |  |
| [`Clear-EventLog`](#clear-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 清除本地或远程计算机上的指定事件日志中的所有条目。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Clear-KdsCache`](#clear-kdscache) | Kds | 都有 | 无 | 清除本地计算机的组密钥缓存。 | 不实现 |  |
| [`Clear-RecycleBin`](#clear-recyclebin) | Microsoft.PowerShell.Management | 仅7 | 仅7提供，5.1 中以 Clear-Recyclebin 形式提供（名称大小写不同） | 清除当前用户的回收站的内容。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Clear-Recyclebin`](#clear-recyclebin) | Microsoft.PowerShell.Management | 仅5.1 | 仅5.1提供，7 中以 Clear-RecycleBin 形式提供（名称大小写不同） | 清除当前用户的回收站的内容。 | 不实现 |  |
| [`Clear-ReFSDedupSchedule`](#clear-refsdedupschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 清除指定 ReFS 卷上的重复数据删除计划任务。 | 不实现 |  |
| [`Clear-ReFSDedupScrubSchedule`](#clear-refsdedupscrubschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 清除指定 ReFS 卷上的重复数据删除清理计划。 | 不实现 |  |
| [`Clear-Tpm`](#clear-tpm) | TrustedPlatformModule | 都有 | 无 | 将 TPM 重置为默认状态。 | 不实现 |  |
| [`Clear-UevAppxPackage`](#clear-uevappxpackage) | UEV | 仅5.1 | 仅5.1提供 | 清除注册表计算机或用户节中的设置。 | 不实现 |  |
| [`Clear-UevConfiguration`](#clear-uevconfiguration) | UEV | 仅5.1 | 仅5.1提供 | 清除 UE-V 配置设置。 | 不实现 |  |
| [`Clear-WindowsCorruptMountPoint`](#clear-windowscorruptmountpoint) | Dism | 都有 | 无 | 删除与已损坏的已装载映像关联的全部资源。 | 不实现 |  |
| [`Complete-BitsTransfer`](#complete-bitstransfer) | BitsTransfer | 都有 | 无 | 完成 BITS 传输作业。 | 不实现 |  |
| [`Complete-DtcDiagnosticTransaction`](#complete-dtcdiagnostictransaction) | MsDtc | 都有 | 无 | 指定的事务是根事务则调用提交，否则对事务对象调用完成方法。 | 不实现 |  |
| [`Complete-Transaction`](#complete-transaction) | Microsoft.PowerShell.Management | 都有 | 无 | 提交活动事务。 | 不实现 |  |
| [`Confirm-SecureBootUEFI`](#confirm-securebootuefi) | SecureBoot | 都有 | 无 | 检查本地计算机安全启动状态，确认安全启动已启用。 | 不实现 |  |
| [`Connect-PSSession`](#connect-pssession) | Microsoft.PowerShell.Core | 都有 | 无 | 重新连接到断开连接的会话。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Connect-WSMan`](#connect-wsman) | Microsoft.WSMan.Management | 都有 | 无 | 连接到远程计算机上的 WinRM 服务。 | 不实现 |  |
| [`Convert-String`](#convert-string) | Microsoft.PowerShell.Utility | 仅5.1 | 仅5.1提供 | 设置字符串的格式以匹配示例。 | 不实现 |  |
| [`ConvertFrom-CIPolicy`](#convertfrom-cipolicy) | ConfigCI | 都有 | 无 | 将包含代码完整性策略的 .xml 文件转换为二进制格式。 | 不实现 |  |
| [`ConvertFrom-SddlString`](#convertfrom-sddlstring) | Microsoft.PowerShell.Utility | 仅7 | 仅7提供 | 将 SDDL 字符串转换为自定义对象。 | 不实现 | 序列化 / 标记 / 格式（少用） |
| [`ConvertFrom-String`](#convertfrom-string) | Microsoft.PowerShell.Utility | 仅5.1 | 仅5.1提供 | 从字符串内容中提取和分析结构化属性。 | 不实现 |  |
| [`ConvertTo-ProcessMitigationPolicy`](#convertto-processmitigationpolicy) | ProcessMitigations | 都有 | 无 | 转换缓解策略文件格式。 | 不实现 |  |
| [`ConvertTo-TpmOwnerAuth`](#convertto-tpmownerauth) | TrustedPlatformModule | 都有 | 无 | 从给定字符串创建 TPM 所有者授权值。 | 不实现 |  |
| [`Copy-BcdEntry`](#copy-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Copy-UserInternationalSettingsToSystem`](#copy-userinternationalsettingstosystem) | International | 都有 | 无 | 将当前用户的区域设置复制到系统。 | 不实现 |  |
| [`Disable-AppBackgroundTaskDiagnosticLog`](#disable-appbackgroundtaskdiagnosticlog) | AppBackgroundTask | 都有 | 无 | 禁用事件查看器中的后台任务日志。 | 不实现 |  |
| [`Disable-Appv`](#disable-appv) | AppvClient | 仅5.1 | 仅5.1提供 | 禁用 App-V 服务。 | 不实现 |  |
| [`Disable-AppvClientConnectionGroup`](#disable-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 禁用运行 App-V 客户端的计算机上的连接组。 | 不实现 |  |
| [`Disable-BcdElementBootDebug`](#disable-bcdelementbootdebug) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Disable-BcdElementBootEms`](#disable-bcdelementbootems) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Disable-BcdElementDebug`](#disable-bcdelementdebug) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Disable-BcdElementEms`](#disable-bcdelementems) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Disable-BcdElementEventLogging`](#disable-bcdelementeventlogging) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Disable-BcdElementHypervisorDebug`](#disable-bcdelementhypervisordebug) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Disable-ComputerRestore`](#disable-computerrestore) | Microsoft.PowerShell.Management | 都有 | 无 | 禁用指定文件系统驱动器上的系统还原功能。 | 不实现 |  |
| [`Disable-JobTrigger`](#disable-jobtrigger) | PSScheduledJob | 都有 | 无 | 禁用计划作业的作业触发器。 | 不实现 |  |
| [`Disable-LocalUser`](#disable-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 禁用本地用户帐户。 | 不实现 |  |
| [`Disable-PSRemoting`](#disable-psremoting) | Microsoft.PowerShell.Core | 都有 | 无 | 阻止 PowerShell 终结点接收远程连接。 | 不实现 |  |
| [`Disable-PSSessionConfiguration`](#disable-pssessionconfiguration) | Microsoft.PowerShell.Core | 都有 | 无 | 在本地计算机上禁用会话配置。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Disable-ReFSDedup`](#disable-refsdedup) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 在指定 ReFS 卷禁用数据重复删除。 | 不实现 |  |
| [`Disable-ScheduledJob`](#disable-scheduledjob) | PSScheduledJob | 都有 | 无 | 禁用计划作业。 | 不实现 |  |
| [`Disable-TlsCipherSuite`](#disable-tlsciphersuite) | TLS | 都有 | 无 | 禁用 TLS 密码套件。 | 不实现 |  |
| [`Disable-TlsEccCurve`](#disable-tlsecccurve) | TLS | 都有 | 无 | 禁用计算机上 TLS 的 ECC 密码套件。 | 不实现 |  |
| [`Disable-TlsSessionTicketKey`](#disable-tlssessionticketkey) | TLS | 都有 | 无 | 禁用 TLS 会话票证密钥。 | 不实现 |  |
| [`Disable-TpmAutoProvisioning`](#disable-tpmautoprovisioning) | TrustedPlatformModule | 都有 | 无 | 禁用 TPM 自动预配。 | 不实现 |  |
| [`Disable-Uev`](#disable-uev) | UEV | 仅5.1 | 仅5.1提供 | 禁用 UE-V 服务。 | 不实现 |  |
| [`Disable-UevAppxPackage`](#disable-uevappxpackage) | UEV | 仅5.1 | 仅5.1提供 | 禁用 UE-V 对 Windows 8 应用的同步。 | 不实现 |  |
| [`Disable-UevTemplate`](#disable-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 禁用设置位置模板。 | 不实现 |  |
| [`Disable-WindowsErrorReporting`](#disable-windowserrorreporting) | WindowsErrorReporting | 都有 | 无 | 禁用 Windows 错误报告。 | 不实现 |  |
| [`Disable-WindowsOptionalFeature`](#disable-windowsoptionalfeature) | Dism | 都有 | 无 | 禁用 Windows 映像中的功能。 | 不实现 |  |
| [`Disable-WSManCredSSP`](#disable-wsmancredssp) | Microsoft.WSMan.Management | 都有 | 无 | 在计算机上禁用 CredSSP 身份验证。 | 不实现 |  |
| [`Disconnect-PSSession`](#disconnect-pssession) | Microsoft.PowerShell.Core | 都有 | 语法不同 | 断开与会话的连接。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Disconnect-WSMan`](#disconnect-wsman) | Microsoft.WSMan.Management | 都有 | 无 | 断开客户端与远程计算机上的 WinRM 服务的连接。 | 不实现 |  |
| [`Dismount-AppxVolume`](#dismount-appxvolume) | Appx | 都有 | 无 | 卸载 appx 卷。 | 不实现 |  |
| [`Dismount-WindowsImage`](#dismount-windowsimage) | Dism | 都有 | 无 | 从其映射目录卸载 Windows 映像。 | 不实现 |  |
| [`Edit-CIPolicyRule`](#edit-cipolicyrule) | ConfigCI | 都有 | 无 | 此 cmdlet 不受支持。 | 不实现 |  |
| [`Enable-AppBackgroundTaskDiagnosticLog`](#enable-appbackgroundtaskdiagnosticlog) | AppBackgroundTask | 都有 | 无 | 启用事件查看器中的后台任务日志。 | 不实现 |  |
| [`Enable-Appv`](#enable-appv) | AppvClient | 仅5.1 | 仅5.1提供 | 启用 App-V 服务。 | 不实现 |  |
| [`Enable-AppvClientConnectionGroup`](#enable-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 启用运行 App-V 客户端的计算机上正在运行的连接组。 | 不实现 |  |
| [`Enable-BcdElementBootDebug`](#enable-bcdelementbootdebug) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Enable-BcdElementBootEms`](#enable-bcdelementbootems) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Enable-BcdElementDebug`](#enable-bcdelementdebug) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Enable-BcdElementEms`](#enable-bcdelementems) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Enable-BcdElementEventLogging`](#enable-bcdelementeventlogging) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Enable-BcdElementHypervisorDebug`](#enable-bcdelementhypervisordebug) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Enable-ComputerRestore`](#enable-computerrestore) | Microsoft.PowerShell.Management | 都有 | 无 | 在指定的文件系统驱动器上启用系统还原功能。 | 不实现 |  |
| [`Enable-JobTrigger`](#enable-jobtrigger) | PSScheduledJob | 都有 | 无 | 启用计划作业的作业触发器。 | 不实现 |  |
| [`Enable-LocalUser`](#enable-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 启用本地用户帐户。 | 不实现 |  |
| [`Enable-PSRemoting`](#enable-psremoting) | Microsoft.PowerShell.Core | 都有 | 无 | 将计算机配置为接收远程命令。 | 不实现 |  |
| [`Enable-PSSessionConfiguration`](#enable-pssessionconfiguration) | Microsoft.PowerShell.Core | 都有 | 无 | 在本地计算机上启用会话配置。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Enable-ReFSDedup`](#enable-refsdedup) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 在指定 ReFS 卷启用数据重复删除。 | 不实现 |  |
| [`Enable-ScheduledJob`](#enable-scheduledjob) | PSScheduledJob | 都有 | 无 | 启用计划作业。 | 不实现 |  |
| [`Enable-TlsCipherSuite`](#enable-tlsciphersuite) | TLS | 都有 | 语法不同 | 启用 TLS 密码套件。 | 不实现 |  |
| [`Enable-TlsEccCurve`](#enable-tlsecccurve) | TLS | 都有 | 语法不同 | 启用 TLS 的 ECC 密码套件。 | 不实现 |  |
| [`Enable-TlsSessionTicketKey`](#enable-tlssessionticketkey) | TLS | 都有 | 无 | 用 TLS 会话票证密钥配置 TLS 服务器。 | 不实现 |  |
| [`Enable-TpmAutoProvisioning`](#enable-tpmautoprovisioning) | TrustedPlatformModule | 都有 | 无 | 启用 TPM 自动预配。 | 不实现 |  |
| [`Enable-Uev`](#enable-uev) | UEV | 仅5.1 | 仅5.1提供 | 启用 UE-V 服务。 | 不实现 |  |
| [`Enable-UevAppxPackage`](#enable-uevappxpackage) | UEV | 仅5.1 | 仅5.1提供 | 启用 UE-V 对 Windows 8 应用的同步。 | 不实现 |  |
| [`Enable-UevTemplate`](#enable-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 启用设置位置模板。 | 不实现 |  |
| [`Enable-WindowsErrorReporting`](#enable-windowserrorreporting) | WindowsErrorReporting | 都有 | 无 | 启用 Windows 错误报告。 | 不实现 |  |
| [`Enable-WindowsOptionalFeature`](#enable-windowsoptionalfeature) | Dism | 都有 | 无 | 启用 Windows 映像中的功能。 | 不实现 |  |
| [`Enable-WSManCredSSP`](#enable-wsmancredssp) | Microsoft.WSMan.Management | 都有 | 无 | 在计算机上启用凭据安全支持提供程序（CredSSP）身份验证。 | 不实现 |  |
| [`Expand-OsImage`](#expand-osimage) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Expand-WindowsCustomDataImage`](#expand-windowscustomdataimage) | Dism | 都有 | 无 | 展开自定义数据映像。 | 不实现 |  |
| [`Expand-WindowsImage`](#expand-windowsimage) | Dism | 都有 | 语法不同 | 将映像应用到指定位置。 | 不实现 |  |
| [`Export-BcdStore`](#export-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Export-BinaryMiLog`](#export-binarymilog) | CimCmdlets | 都有 | 无 | 创建对象或对象的二进制编码表示形式，并将其存储在文件中。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Export-Certificate`](#export-certificate) | PKI | 都有 | 无 | 从证书存储导出证书到文件。 | 不实现 |  |
| [`Export-Console`](#export-console) | Microsoft.PowerShell.Core | 仅5.1 | 仅5.1提供 | 将当前会话中管理单元的名称导出到控制台文件。 | 不实现 |  |
| [`Export-Counter`](#export-counter) | Microsoft.PowerShell.Diagnostics | 都有 | 无 | 将性能计数器数据导出到日志文件。 | 不实现 |  |
| [`Export-OsImage`](#export-osimage) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Export-PfxCertificate`](#export-pfxcertificate) | PKI | 都有 | 无 | 将证书或 PFXData 对象导出到 PFX 文件。 | 不实现 |  |
| [`Export-ProvisioningPackage`](#export-provisioningpackage) | Provisioning | 都有 | 无 | 提取预配包内容。 | 不实现 |  |
| [`Export-StartLayout`](#export-startlayout) | StartLayout | 都有 | 无 | 导出开始屏幕布局。 | 不实现 |  |
| [`Export-StartLayoutEdgeAssets`](#export-startlayoutedgeassets) | StartLayout | 都有 | 无 | 导出显示自定义图像的 Edge 辅助磁贴。 | 不实现 |  |
| [`Export-TlsSessionTicketKey`](#export-tlssessionticketkey) | TLS | 都有 | 无 | 导出 TLS 会话票证密钥。 | 不实现 |  |
| [`Export-Trace`](#export-trace) | Provisioning | 都有 | 无 | 为预配导出事件跟踪日志文件。 | 不实现 |  |
| [`Export-UevConfiguration`](#export-uevconfiguration) | UEV | 仅5.1 | 仅5.1提供 | 导出 UE-V 配置。 | 不实现 |  |
| [`Export-UevPackage`](#export-uevpackage) | UEV | 仅5.1 | 仅5.1提供 | 导出设置包中存储的设置。 | 不实现 |  |
| [`Export-WindowsCapabilitySource`](#export-windowscapabilitysource) | Dism | 都有 | 无 | 创建自定义 FOD 存储库，收录支持指定功能安装的包。 | 不实现 |  |
| [`Export-WindowsDriver`](#export-windowsdriver) | Dism | 都有 | 无 | 将 Windows 映像中的第三方驱动程序导出到目标文件夹。 | 不实现 |  |
| [`Export-WindowsImage`](#export-windowsimage) | Dism | 都有 | 语法不同 | 将指定映像复制到另一个映像文件。 | 不实现 |  |
| [`Find-LapsADExtendedRights`](#find-lapsadextendedrights) | LAPS | 都有 | 无 | 在 AD 中查询被授予读取 LAPS 密码属性权限的主体。 | 不实现 |  |
| [`Format-SecureBootUEFI`](#format-securebootuefi) | SecureBoot | 都有 | 无 | 将证书或哈希格式化为返回的内容对象，并创建待签名的文件。 | 不实现 |  |
| [`Get-Acl`](#get-acl) | Microsoft.PowerShell.Security | 都有 | 语法不同 | 获取资源（如文件或注册表项）的安全描述符。 | 不实现 |  |
| [`Get-AppLockerFileInformation`](#get-applockerfileinformation) | AppLocker | 都有 | 无 | 获取从文件列表或事件日志创建 AppLocker 规则所需的文件信息。 | 不实现 |  |
| [`Get-AppLockerPolicy`](#get-applockerpolicy) | AppLocker | 都有 | 无 | 获取本地、有效或域 AppLocker 策略。 | 不实现 |  |
| [`Get-AppProvisionedSharedPackageContainer`](#get-appprovisionedsharedpackagecontainer) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-AppSharedPackageContainer`](#get-appsharedpackagecontainer) | Appx | 都有 | 无 | 获取共享包容器的信息。 | 不实现 |  |
| [`Get-AppvClientApplication`](#get-appvclientapplication) | AppvClient | 仅5.1 | 仅5.1提供 | 返回 App-V 客户端包中的应用。 | 不实现 |  |
| [`Get-AppvClientConfiguration`](#get-appvclientconfiguration) | AppvClient | 仅5.1 | 仅5.1提供 | 返回 App-V 客户端配置。 | 不实现 |  |
| [`Get-AppvClientConnectionGroup`](#get-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 返回 App-V 连接组对象。 | 不实现 |  |
| [`Get-AppvClientMode`](#get-appvclientmode) | AppvClient | 仅5.1 | 仅5.1提供 | 显示 App-V 客户端模式。 | 不实现 |  |
| [`Get-AppvClientPackage`](#get-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 返回 App-V 客户端包。 | 不实现 |  |
| [`Get-AppvPublishingServer`](#get-appvpublishingserver) | AppvClient | 仅5.1 | 仅5.1提供 | 返回 App-V 服务器对象。 | 不实现 |  |
| [`Get-AppvStatus`](#get-appvstatus) | AppvClient | 仅5.1 | 仅5.1提供 | 获取 App-V 服务状态。 | 不实现 |  |
| [`Get-AppxDefaultVolume`](#get-appxdefaultvolume) | Appx | 都有 | 无 | 获取默认 appx 卷。 | 不实现 |  |
| [`Get-AppxPackage`](#get-appxpackage) | Appx | 都有 | 无 | 获取用户配置文件中安装的应用包列表。 | 不实现 |  |
| [`Get-AppxPackageAutoUpdateSettings`](#get-appxpackageautoupdatesettings) | Appx | 都有 | 无 | 显示特定 Windows 应用的配置设置。 | 不实现 |  |
| [`Get-AppxPackageManifest`](#get-appxpackagemanifest) | Appx | 都有 | 无 | 获取应用包的清单。 | 不实现 |  |
| [`Get-AppxProvisionedPackage`](#get-appxprovisionedpackage) | Dism | 都有 | 无 | 获取映像中为每个新用户安装的应用包 (.appx) 信息。 | 不实现 |  |
| [`Get-AppxVolume`](#get-appxvolume) | Appx | 都有 | 无 | 获取计算机的 appx 卷。 | 不实现 |  |
| [`Get-AuthenticodeSignature`](#get-authenticodesignature) | Microsoft.PowerShell.Security | 都有 | 无 | 获取有关文件的验证码签名的信息。 | 不实现 |  |
| [`Get-BcdEntry`](#get-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-BcdEntryDebugSettings`](#get-bcdentrydebugsettings) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-BcdEntryHypervisorSettings`](#get-bcdentryhypervisorsettings) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-BcdStore`](#get-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-BitsTransfer`](#get-bitstransfer) | BitsTransfer | 都有 | 无 | 获取现有 BITS 传输作业的关联 BitsJob 对象。 | 不实现 |  |
| [`Get-Certificate`](#get-certificate) | PKI | 都有 | 无 | 向注册服务器提交证书申请并安装响应，或取回已提交申请的证书。 | 不实现 |  |
| [`Get-CertificateAutoEnrollmentPolicy`](#get-certificateautoenrollmentpolicy) | PKI | 都有 | 无 | 检索证书自动注册策略设置。 | 不实现 |  |
| [`Get-CertificateEnrollmentPolicyServer`](#get-certificateenrollmentpolicyserver) | PKI | 都有 | 无 | 返回全部证书注册策略服务器 URL 配置。 | 不实现 |  |
| [`Get-CertificateNotificationTask`](#get-certificatenotificationtask) | PKI | 都有 | 无 | 返回全部已注册证书通知任务。 | 不实现 |  |
| [`Get-CimAssociatedInstance`](#get-cimassociatedinstance) | CimCmdlets | 都有 | 语法不同 | 通过关联检索连接到特定 CIM 实例的 CIM 实例。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Get-CimClass`](#get-cimclass) | CimCmdlets | 都有 | 语法不同 | 获取特定命名空间中的 CIM 类的列表。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Get-CimInstance`](#get-ciminstance) | CimCmdlets | 都有 | 语法不同 | 从 CIM 服务器获取类的 CIM 实例。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Get-CimSession`](#get-cimsession) | CimCmdlets | 都有 | 语法不同 | 从当前会话中获取 CIM 会话对象。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Get-CIPolicy`](#get-cipolicy) | ConfigCI | 都有 | 无 | 获取代码完整性策略中的规则。 | 不实现 |  |
| [`Get-CIPolicyIdInfo`](#get-cipolicyidinfo) | ConfigCI | 都有 | 无 | 显示代码完整性策略信息。 | 不实现 |  |
| [`Get-CIPolicyInfo`](#get-cipolicyinfo) | ConfigCI | 都有 | 无 | 此 cmdlet 不受支持。 | 不实现 |  |
| [`Get-ComputerInfo`](#get-computerinfo) | Microsoft.PowerShell.Management | 都有 | 无 | 获取系统和作系统属性的合并对象。 | Go实现 | 字段精简。 |
| [`Get-ComputerRestorePoint`](#get-computerrestorepoint) | Microsoft.PowerShell.Management | 都有 | 无 | 获取本地计算机上的还原点。 | 不实现 |  |
| [`Get-ControlPanelItem`](#get-controlpanelitem) | Microsoft.PowerShell.Management | 都有 | 无 | 获取控制面板项。 | 不实现 |  |
| [`Get-Counter`](#get-counter) | Microsoft.PowerShell.Diagnostics | 都有 | 无 | 从本地和远程计算机获取性能计数器数据。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Get-DAPolicyChange`](#get-dapolicychange) | NetSecurity | 都有 | 无 | 获取需增删的 IP 地址列表以更新 IPsec 规则，并生成更新规则的脚本。 | 不实现 |  |
| [`Get-DeliveryOptimizationLog`](#get-deliveryoptimizationlog) | DeliveryOptimization | 都有 | 语法不同 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-DeliveryOptimizationLogAnalysis`](#get-deliveryoptimizationloganalysis) | DeliveryOptimization | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-EventLog`](#get-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 获取本地计算机或远程计算机上事件日志中的事件或事件日志列表。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Get-HotFix`](#get-hotfix) | Microsoft.PowerShell.Management | 都有 | 无 | 获取在本地或远程计算机上安装的修补程序。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Get-InstalledLanguage`](#get-installedlanguage) | LanguagePackManagement | 都有 | 无 | 返回设备上已安装语言的信息。 | 不实现 |  |
| [`Get-JobTrigger`](#get-jobtrigger) | PSScheduledJob | 都有 | 无 | 获取计划作业的作业触发器。 | 不实现 |  |
| [`Get-KdsConfiguration`](#get-kdsconfiguration) | Kds | 都有 | 无 | 从 AD 检索 KdsSvc 的当前配置。 | 不实现 |  |
| [`Get-KdsRootKey`](#get-kdsrootkey) | Kds | 都有 | 无 | 检索 KdsSvc 存储的根密钥值列表。 | 不实现 |  |
| [`Get-LapsADPassword`](#get-lapsadpassword) | LAPS | 都有 | 无 | 从 AD 指定计算机或域控制器对象查询 LAPS 凭据。 | 不实现 |  |
| [`Get-LocalGroup`](#get-localgroup) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 获取本地安全组。 | 不实现 |  |
| [`Get-LocalGroupMember`](#get-localgroupmember) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 从本地组获取成员。 | 不实现 |  |
| [`Get-LocalUser`](#get-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 获取本地用户帐户。 | 不实现 |  |
| [`Get-NonRemovableAppsPolicy`](#get-nonremovableappspolicy) | Dism | 都有 | 无 | 返回已安装且配置为不可移除的应用包列表。 | 不实现 |  |
| [`Get-OSConfiguration`](#get-osconfiguration) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-OsConfigurationDocument`](#get-osconfigurationdocument) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-OsConfigurationDocumentContent`](#get-osconfigurationdocumentcontent) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-OsConfigurationDocumentResult`](#get-osconfigurationdocumentresult) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-OsConfigurationProperty`](#get-osconfigurationproperty) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-OSConfigurationScenarioDefinition`](#get-osconfigurationscenariodefinition) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-OSConfigurationScenarioDefinitionInfo`](#get-osconfigurationscenariodefinitioninfo) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-PfxData`](#get-pfxdata) | PKI | 都有 | 无 | 将 PFX 文件内容提取为结构，不导入证书存储。 | 不实现 |  |
| [`Get-PmemDedicatedMemory`](#get-pmemdedicatedmemory) | PersistentMemory | 都有 | 无 | 获取专用持久内存。 | 不实现 |  |
| [`Get-PmemDisk`](#get-pmemdisk) | PersistentMemory | 都有 | 无 | 获取持久内存磁盘。 | 不实现 |  |
| [`Get-PmemPhysicalDevice`](#get-pmemphysicaldevice) | PersistentMemory | 都有 | 无 | 获取与持久内存关联的物理设备。 | 不实现 |  |
| [`Get-PmemUnusedRegion`](#get-pmemunusedregion) | PersistentMemory | 都有 | 无 | 获取持久内存中未使用的区域。 | 不实现 |  |
| [`Get-ProcessMitigation`](#get-processmitigation) | ProcessMitigations | 都有 | 无 | 从注册表或运行中的进程获取当前进程缓解设置，或全部保存到 XML。 | 不实现 |  |
| [`Get-ProvisioningPackage`](#get-provisioningpackage) | Provisioning | 都有 | 无 | 获取已安装预配包的信息。 | 不实现 |  |
| [`Get-PSSessionCapability`](#get-pssessioncapability) | Microsoft.PowerShell.Core | 都有 | 简介不同 | 5.1：获取受约束会话配置上特定用户的功能。 / 7：在特定会话配置中获取特定用户的权限。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Get-PSSessionConfiguration`](#get-pssessionconfiguration) | Microsoft.PowerShell.Core | 都有 | 简介不同 | 5.1：获取计算机上的已注册会话配置。 / 7：获取计算机上已注册的会话配置。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Get-PSSnapin`](#get-pssnapin) | Microsoft.PowerShell.Core | 仅5.1 | 仅5.1提供 | 获取计算机上的 Windows PowerShell 管理单元。 | 不实现 |  |
| [`Get-RecoveryManagementPluginAltitude`](#get-recoverymanagementpluginaltitude) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-RecoveryManagementPluginInfo`](#get-recoverymanagementplugininfo) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-RecoveryManagementPlugins`](#get-recoverymanagementplugins) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-RecoveryRemoteManagementStatus`](#get-recoveryremotemanagementstatus) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-ReFSDedupSchedule`](#get-refsdedupschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 检索指定 ReFS 卷上的重复数据删除计划。 | 不实现 |  |
| [`Get-ReFSDedupScrubSchedule`](#get-refsdedupscrubschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 检索指定 ReFS 卷上的重复数据删除清理计划。 | 不实现 |  |
| [`Get-ReFSDedupStatus`](#get-refsdedupstatus) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 检索指定 ReFS 卷上数据重复删除的状态。 | 不实现 |  |
| [`Get-ScheduledJob`](#get-scheduledjob) | PSScheduledJob | 都有 | 无 | 获取本地计算机上的计划作业。 | 不实现 |  |
| [`Get-ScheduledJobOption`](#get-scheduledjoboption) | PSScheduledJob | 都有 | 无 | 获取计划作业的作业选项。 | 不实现 |  |
| [`Get-SecureBootPolicy`](#get-securebootpolicy) | SecureBoot | 都有 | 无 | 获取安全启动配置策略的发布者 GUID 和策略版本。 | 不实现 |  |
| [`Get-SecureBootSVN`](#get-securebootsvn) | SecureBoot | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Get-SecureBootUEFI`](#get-securebootuefi) | SecureBoot | 都有 | 无 | 获取与安全启动相关的 UEFI 变量值。 | 不实现 |  |
| [`Get-Service`](#get-service) | Microsoft.PowerShell.Management | 都有 | 简介不同；语法不同 | 5.1：获取本地或远程计算机上的服务。 / 7：获取计算机上的服务。 | 映射 Linux（systemctl） |  |
| [`Get-SystemDriver`](#get-systemdriver) | ConfigCI | 都有 | 无 | 扫描系统驱动程序。 | 不实现 |  |
| [`Get-SystemPreferredUILanguage`](#get-systempreferreduilanguage) | LanguagePackManagement | 都有 | 无 | 返回当前系统首选语言。 | 不实现 |  |
| [`Get-TlsCipherSuite`](#get-tlsciphersuite) | TLS | 都有 | 无 | 获取计算机的 TLS 密码套件。 | 不实现 |  |
| [`Get-TlsEccCurve`](#get-tlsecccurve) | TLS | 都有 | 无 | 获取计算机上 TLS 可用的 ECC 密码套件列表。 | 不实现 |  |
| [`Get-Tpm`](#get-tpm) | TrustedPlatformModule | 都有 | 无 | 获取包含 TPM 信息的对象。 | 不实现 |  |
| [`Get-TpmEndorsementKeyInfo`](#get-tpmendorsementkeyinfo) | TrustedPlatformModule | 都有 | 无 | 获取 TPM 背书密钥和证书的信息。 | 不实现 |  |
| [`Get-TpmSupportedFeature`](#get-tpmsupportedfeature) | TrustedPlatformModule | 都有 | 无 | 验证 TPM 是否支持指定功能。 | 不实现 |  |
| [`Get-Transaction`](#get-transaction) | Microsoft.PowerShell.Management | 都有 | 无 | 获取当前（活动）事务。 | 不实现 |  |
| [`Get-TroubleshootingPack`](#get-troubleshootingpack) | TroubleshootingPack | 都有 | 无 | 获取疑难解答包或生成应答文件。 | 不实现 |  |
| [`Get-TrustedProvisioningCertificate`](#get-trustedprovisioningcertificate) | Provisioning | 都有 | 无 | 列出已安装的受信任预配证书。 | 不实现 |  |
| [`Get-UevAppxPackage`](#get-uevappxpackage) | UEV | 仅5.1 | 仅5.1提供 | 获取 Windows 8 应用及同步状态列表。 | 不实现 |  |
| [`Get-UevConfiguration`](#get-uevconfiguration) | UEV | 仅5.1 | 仅5.1提供 | 获取 UE-V 配置设置。 | 不实现 |  |
| [`Get-UevStatus`](#get-uevstatus) | UEV | 仅5.1 | 仅5.1提供 | 获取 UE-V 服务状态。 | 不实现 |  |
| [`Get-UevTemplate`](#get-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 获取 UE-V 设置位置模板。 | 不实现 |  |
| [`Get-UevTemplateProgram`](#get-uevtemplateprogram) | UEV | 仅5.1 | 仅5.1提供 | 获取设置位置模板定义的程序信息。 | 不实现 |  |
| [`Get-WheaMemoryPolicy`](#get-wheamemorypolicy) | Whea | 都有 | 无 | 获取计算机的 WHEA 内存策略。 | 不实现 |  |
| [`Get-WIMBootEntry`](#get-wimbootentry) | Dism | 都有 | 无 | 显示指定磁盘卷的 WIMBoot 配置项。 | 不实现 |  |
| [`Get-WinAcceptLanguageFromLanguageListOptOut`](#get-winacceptlanguagefromlanguagelistoptout) | International | 都有 | 无 | 获取当前用户帐户语言列表退出设置的 HTTP 接受语言。 | 不实现 |  |
| [`Get-WinCultureFromLanguageListOptOut`](#get-winculturefromlanguagelistoptout) | International | 都有 | 无 | 获取当前用户帐户语言列表退出设置的区域。 | 不实现 |  |
| [`Get-WinDefaultInputMethodOverride`](#get-windefaultinputmethodoverride) | International | 都有 | 无 | 获取当前用户帐户的默认输入法替代设置。 | 不实现 |  |
| [`Get-WindowsCapability`](#get-windowscapability) | Dism | 都有 | 无 | 获取映像或正在运行的操作系统的 Windows 功能。 | 不实现 |  |
| [`Get-WindowsDeveloperLicense`](#get-windowsdeveloperlicense) | WindowsDeveloperLicense | 都有 | 无 | 提供当前计算机开发者模式的信息。 | 不实现 |  |
| [`Get-WindowsDriver`](#get-windowsdriver) | Dism | 都有 | 无 | 显示 Windows 映像中的驱动程序信息。 | 不实现 |  |
| [`Get-WindowsEdition`](#get-windowsedition) | Dism | 都有 | 无 | 获取 Windows 映像的版本信息。 | 不实现 |  |
| [`Get-WindowsErrorReporting`](#get-windowserrorreporting) | WindowsErrorReporting | 都有 | 无 | 检索 Windows 错误报告状态。 | 不实现 |  |
| [`Get-WindowsImage`](#get-windowsimage) | Dism | 都有 | 语法不同 | 获取 WIM 或 VHD 文件中 Windows 映像的信息。 | 不实现 |  |
| [`Get-WindowsImageContent`](#get-windowsimagecontent) | Dism | 都有 | 语法不同 | 显示指定映像中的文件和文件夹列表。 | 不实现 |  |
| [`Get-WindowsOptionalFeature`](#get-windowsoptionalfeature) | Dism | 都有 | 无 | 获取 Windows 映像中可选功能的信息。 | 不实现 |  |
| [`Get-WindowsPackage`](#get-windowspackage) | Dism | 都有 | 无 | 获取 Windows 映像中包的信息。 | 不实现 |  |
| [`Get-WindowsReservedStorageState`](#get-windowsreservedstoragestate) | Dism | 都有 | 无 | 获取映像的保留存储状态。 | 不实现 |  |
| [`Get-WindowsSearchSetting`](#get-windowssearchsetting) | WindowsSearch | 都有 | 无 | 获取 Windows 搜索的设置值。 | 不实现 |  |
| [`Get-WinEvent`](#get-winevent) | Microsoft.PowerShell.Diagnostics | 都有 | 简介不同 | 5.1：从本地和远程计算机上的事件日志和事件跟踪日志文件中获取事件。 / 7：获取本地和远程计算机上的事件日志和事件跟踪日志文件中的事件。 | 不实现 |  |
| [`Get-WinHomeLocation`](#get-winhomelocation) | International | 都有 | 无 | 获取当前用户帐户的 Windows GeoID 主位置设置。 | 不实现 |  |
| [`Get-WinLanguageBarOption`](#get-winlanguagebaroption) | International | 都有 | 无 | 获取当前用户帐户的语言栏模式和类型。 | 不实现 |  |
| [`Get-WinSystemLocale`](#get-winsystemlocale) | International | 都有 | 无 | 获取当前计算机的系统区域设置。 | 不实现 |  |
| [`Get-WinUILanguageOverride`](#get-winuilanguageoverride) | International | 都有 | 无 | 获取当前用户帐户的 Windows UI 语言替代设置。 | 不实现 |  |
| [`Get-WinUserLanguageList`](#get-winuserlanguagelist) | International | 都有 | 无 | 获取当前用户帐户的语言列表。 | 不实现 |  |
| [`Get-WmiObject`](#get-wmiobject) | Microsoft.PowerShell.Management | 都有 | 无 | 获取 Windows Management Instrumentation （WMI） 类的实例或有关可用类的信息。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Get-WSManCredSSP`](#get-wsmancredssp) | Microsoft.WSMan.Management | 都有 | 无 | 获取客户端的凭据安全支持提供程序相关配置。 | 不实现 |  |
| [`Get-WSManInstance`](#get-wsmaninstance) | Microsoft.WSMan.Management | 都有 | 无 | 显示资源 URI 指定的资源实例的管理信息。 | 不实现 |  |
| [`Import-BcdStore`](#import-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Import-BinaryMiLog`](#import-binarymilog) | CimCmdlets | 都有 | 无 | 用于根据导出文件的内容重新创建已保存的对象。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Import-Certificate`](#import-certificate) | PKI | 都有 | 无 | 将一个或多个证书导入证书存储。 | 不实现 |  |
| [`Import-Counter`](#import-counter) | Microsoft.PowerShell.Diagnostics | 都有 | 无 | 导入性能计数器日志文件，并创建表示日志中每个计数器示例的对象。 | 不实现 |  |
| [`Import-PfxCertificate`](#import-pfxcertificate) | PKI | 都有 | 无 | 将 PFX 文件中的证书和私钥导入目标存储。 | 不实现 |  |
| [`Import-StartLayout`](#import-startlayout) | StartLayout | 都有 | 无 | 将开始布局导入已装载的 Windows 映像。 | 不实现 |  |
| [`Import-TpmOwnerAuth`](#import-tpmownerauth) | TrustedPlatformModule | 都有 | 无 | 将 TPM 所有者授权值导入注册表。 | 不实现 |  |
| [`Import-UevConfiguration`](#import-uevconfiguration) | UEV | 仅5.1 | 仅5.1提供 | 导入 UE-V 配置。 | 不实现 |  |
| [`Initialize-PmemPhysicalDevice`](#initialize-pmemphysicaldevice) | PersistentMemory | 都有 | 无 | 初始化物理持久内存设备的标签存储区。 | 不实现 |  |
| [`Initialize-Tpm`](#initialize-tpm) | TrustedPlatformModule | 都有 | 无 | 执行 TPM 预配流程的一部分。 | 不实现 |  |
| [`Install-Language`](#install-language) | LanguagePackManagement | 都有 | 无 | 在设备上安装语言。 | 不实现 |  |
| [`Install-ProvisioningPackage`](#install-provisioningpackage) | Provisioning | 都有 | 无 | 在本地计算机安装 .PPKG 包。 | 不实现 |  |
| [`Install-TrustedProvisioningCertificate`](#install-trustedprovisioningcertificate) | Provisioning | 都有 | 无 | 向受信任证书存储添加证书。 | 不实现 |  |
| [`Invoke-CimMethod`](#invoke-cimmethod) | CimCmdlets | 都有 | 语法不同 | 调用 CIM 类的方法。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Invoke-CommandInDesktopPackage`](#invoke-commandindesktoppackage) | Appx | 都有 | 无 | 在打包应用的上下文中创建新进程的调试工具。 | 不实现 |  |
| [`Invoke-DscResource`](#invoke-dscresource) | PSDesiredStateConfiguration | 都有 | 无 | 运行指定的 PowerShell Desired State Configuration （DSC） 资源的方法。 | 不实现 |  |
| [`Invoke-LapsPolicyProcessing`](#invoke-lapspolicyprocessing) | LAPS | 都有 | 无 | 使 LAPS 处理当前配置的策略。 | 不实现 |  |
| [`Invoke-TroubleshootingPack`](#invoke-troubleshootingpack) | TroubleshootingPack | 都有 | 无 | 运行疑难解答包。 | 不实现 |  |
| [`Invoke-WmiMethod`](#invoke-wmimethod) | Microsoft.PowerShell.Management | 都有 | 无 | 调用 WMI 方法。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Invoke-WSManAction`](#invoke-wsmanaction) | Microsoft.WSMan.Management | 都有 | 无 | 对资源 URI 和选择器指定的对象调用作。 | 不实现 |  |
| [`Join-DtcDiagnosticResourceManager`](#join-dtcdiagnosticresourcemanager) | MsDtc | 都有 | 无 | 为事务对象登记诊断资源管理器。 | 不实现 |  |
| [`Limit-EventLog`](#limit-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 设置用于限制事件日志大小及其条目期限的事件日志属性。 | 不实现 |  |
| [`Merge-CIPolicy`](#merge-cipolicy) | ConfigCI | 都有 | 无 | 合并多个代码完整性策略文件中的规则。 | 不实现 |  |
| [`Mount-AppvClientConnectionGroup`](#mount-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 将包内容流式传输到本地磁盘。 | 不实现 |  |
| [`Mount-AppvClientPackage`](#mount-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 将包加载到 App-V 缓存。 | 不实现 |  |
| [`Mount-AppxVolume`](#mount-appxvolume) | Appx | 都有 | 无 | 装载 appx 卷。 | 不实现 |  |
| [`Mount-WindowsImage`](#mount-windowsimage) | Dism | 都有 | 语法不同 | 将 WIM 或 VHD 文件中的 Windows 映像装载到本地目录。 | 不实现 |  |
| [`Move-AppxPackage`](#move-appxpackage) | Appx | 都有 | 无 | 将包从当前位置移到另一个 appx 卷。 | 不实现 |  |
| [`New-AppLockerPolicy`](#new-applockerpolicy) | AppLocker | 都有 | 无 | 从文件信息列表创建新的 AppLocker 策略。 | 不实现 |  |
| [`New-BcdEntry`](#new-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`New-BcdStore`](#new-bcdstore) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`New-CertificateNotificationTask`](#new-certificatenotificationtask) | PKI | 都有 | 无 | 在证书被替换、过期或即将过期时，在任务计划程序中创建新任务。 | 不实现 |  |
| [`New-CimInstance`](#new-ciminstance) | CimCmdlets | 都有 | 语法不同 | 创建 CIM 实例。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`New-CimSession`](#new-cimsession) | CimCmdlets | 都有 | 语法不同 | 创建 CIM 会话。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`New-CimSessionOption`](#new-cimsessionoption) | CimCmdlets | 都有 | 语法不同 | 指定 New-CimSession cmdlet 的高级选项。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`New-CIPolicy`](#new-cipolicy) | ConfigCI | 都有 | 无 | 创建代码完整性策略 .xml 文件。 | 不实现 |  |
| [`New-CIPolicyRule`](#new-cipolicyrule) | ConfigCI | 都有 | 无 | 为用户模式代码和驱动程序生成代码完整性策略规则。 | 不实现 |  |
| [`New-DtcDiagnosticTransaction`](#new-dtcdiagnostictransaction) | MsDtc | 都有 | 无 | 在本地计算机的事务管理器中创建新事务。 | 不实现 |  |
| [`New-EventLog`](#new-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 在本地或远程计算机上创建新的事件日志和新事件源。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`New-FileCatalog`](#new-filecatalog) | Microsoft.PowerShell.Security | 都有 | 无 | 创建一个Windows目录文件，其中包含指定路径中文件和文件夹的加密哈希。 | 不实现 |  |
| [`New-JobTrigger`](#new-jobtrigger) | PSScheduledJob | 都有 | 无 | 为计划作业创建作业触发器。 | 不实现 |  |
| [`New-LocalGroup`](#new-localgroup) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 创建本地安全组。 | 不实现 |  |
| [`New-LocalUser`](#new-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 创建本地用户帐户。 | 不实现 |  |
| [`New-NetIPsecAuthProposal`](#new-netipsecauthproposal) | NetSecurity | 都有 | 无 | 创建主模式身份验证建议，指定 IPsec 主模式协商中提供的身份验证协议套件。 | 不实现 |  |
| [`New-NetIPsecMainModeCryptoProposal`](#new-netipsecmainmodecryptoproposal) | NetSecurity | 都有 | 无 | 创建主模式加密建议，指定 IPsec 主模式协商中提供的加密协议套件。 | 不实现 |  |
| [`New-NetIPsecQuickModeCryptoProposal`](#new-netipsecquickmodecryptoproposal) | NetSecurity | 都有 | 语法不同 | 创建快速模式加密建议，指定 IPsec 快速模式协商中提供的加密协议套件。 | 不实现 |  |
| [`New-PmemDedicatedMemory`](#new-pmemdedicatedmemory) | PersistentMemory | 都有 | 无 | 在指定区域中创建专用持久内存。 | 不实现 |  |
| [`New-PmemDisk`](#new-pmemdisk) | PersistentMemory | 都有 | 无 | 在未使用的持久内存区域中创建持久内存磁盘。 | 不实现 |  |
| [`New-ProvisioningRepro`](#new-provisioningrepro) | Provisioning | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`New-PSWorkflowExecutionOption`](#new-psworkflowexecutionoption) | PSWorkflow | 都有 | 无 | 创建一个对象，其中包含工作流会话的会话配置选项。 | 不实现 |  |
| [`New-ScheduledJobOption`](#new-scheduledjoboption) | PSScheduledJob | 都有 | 无 | 创建一个对象，该对象包含计划作业的高级选项。 | 不实现 |  |
| [`New-SelfSignedCertificate`](#new-selfsignedcertificate) | PKI | 都有 | 无 | 创建用于测试的新自签名证书。 | 不实现 |  |
| [`New-Service`](#new-service) | Microsoft.PowerShell.Management | 都有 | 语法不同 | 创建新的 Windows 服务。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`New-TlsSessionTicketKey`](#new-tlssessionticketkey) | TLS | 都有 | 无 | 创建 TLS 会话票证密钥配置文件。 | 不实现 |  |
| [`New-WebServiceProxy`](#new-webserviceproxy) | Microsoft.PowerShell.Management | 都有 | 无 | 创建一个 Web 服务代理对象，该对象允许你在 PowerShell 中使用和管理 Web 服务。 | 不实现 |  |
| [`New-WindowsCustomImage`](#new-windowscustomimage) | Dism | 都有 | 无 | 在配置了 WIMBoot 的设备上捕获定制或维护过的 Windows 组件映像。 | 不实现 |  |
| [`New-WindowsImage`](#new-windowsimage) | Dism | 都有 | 无 | 将驱动器映像捕获到新 WIM 文件。 | 不实现 |  |
| [`New-WinEvent`](#new-winevent) | Microsoft.PowerShell.Diagnostics | 都有 | 无 | 为指定的事件提供程序创建新的 Windows 事件。 | 不实现 |  |
| [`New-WinUserLanguageList`](#new-winuserlanguagelist) | International | 都有 | 无 | 实例化新语言列表对象。 | 不实现 |  |
| [`New-WSManInstance`](#new-wsmaninstance) | Microsoft.WSMan.Management | 都有 | 无 | 创建管理资源的新实例。 | 不实现 |  |
| [`New-WSManSessionOption`](#new-wsmansessionoption) | Microsoft.WSMan.Management | 都有 | 无 | 创建会话选项哈希表，用作 WS-Management cmdlet 的输入参数。 | 不实现 |  |
| [`Optimize-AppxProvisionedPackages`](#optimize-appxprovisionedpackages) | Dism | 都有 | 无 | 用硬链接替换相同文件，优化映像中预配包的总大小。 | 不实现 |  |
| [`Optimize-WindowsImage`](#optimize-windowsimage) | Dism | 都有 | 无 | 按指定优化配置 Windows 映像。 | 不实现 |  |
| [`Out-GridView`](#out-gridview) | Microsoft.PowerShell.Utility | 都有 | 无 | 将输出发送到单独窗口中的交互式表中。 | 不实现 | GUI 与打印 |
| [`Out-Printer`](#out-printer) | Microsoft.PowerShell.Utility | 都有 | 无 | 将输出发送到打印机。 | 不实现 | GUI 与打印 |
| [`Publish-AppvClientPackage`](#publish-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 发布 App-V 包。 | 不实现 |  |
| [`Publish-DscConfiguration`](#publish-dscconfiguration) | PSDesiredStateConfiguration | 都有 | 无 | 将 DSC 配置发布到一组计算机。 | 不实现 |  |
| [`Receive-DtcDiagnosticTransaction`](#receive-dtcdiagnostictransaction) | MsDtc | 都有 | 无 | 从给定诊断资源管理器传播事务。 | 不实现 |  |
| [`Receive-PSSession`](#receive-pssession) | Microsoft.PowerShell.Core | 都有 | 语法不同 | 获取断开连接的会话中的命令的结果 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Register-CimIndicationEvent`](#register-cimindicationevent) | CimCmdlets | 都有 | 语法不同 | 使用筛选器表达式或查询表达式订阅指示。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Register-PSSessionConfiguration`](#register-pssessionconfiguration) | Microsoft.PowerShell.Core | 都有 | 语法不同 | 创建并注册新的会话配置。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Register-RecoveryManagementPlugin`](#register-recoverymanagementplugin) | Dism | 都有 | 语法不同 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Register-ScheduledJob`](#register-scheduledjob) | PSScheduledJob | 都有 | 无 | 创建计划作业。 | 不实现 |  |
| [`Register-UevTemplate`](#register-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 向 UE-V 注册设置位置模板。 | 不实现 |  |
| [`Register-WmiEvent`](#register-wmievent) | Microsoft.PowerShell.Management | 都有 | 无 | 订阅 Windows Management Instrumentation （WMI） 事件。 | 不实现 |  |
| [`Remove-AppProvisionedSharedPackageContainer`](#remove-appprovisionedsharedpackagecontainer) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Remove-AppSharedPackageContainer`](#remove-appsharedpackagecontainer) | Appx | 都有 | 无 | 删除共享包容器。 | 不实现 |  |
| [`Remove-AppvClientConnectionGroup`](#remove-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 删除客户端上的 App-V 连接组。 | 不实现 |  |
| [`Remove-AppvClientPackage`](#remove-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 从计算机删除包。 | 不实现 |  |
| [`Remove-AppvPublishingServer`](#remove-appvpublishingserver) | AppvClient | 仅5.1 | 仅5.1提供 | 删除 App-V 发布服务器。 | 不实现 |  |
| [`Remove-AppxPackage`](#remove-appxpackage) | Appx | 都有 | 无 | 从一个或多个用户帐户删除应用包。 | 不实现 |  |
| [`Remove-AppxPackageAutoUpdateSettings`](#remove-appxpackageautoupdatesettings) | Appx | 都有 | 无 | 删除特定 Windows 应用的设置。 | 不实现 |  |
| [`Remove-AppxProvisionedPackage`](#remove-appxprovisionedpackage) | Dism | 都有 | 无 | 从 Windows 映像删除应用包 (.appx)。 | 不实现 |  |
| [`Remove-AppxVolume`](#remove-appxvolume) | Appx | 都有 | 无 | 删除 appx 卷。 | 不实现 |  |
| [`Remove-BcdElement`](#remove-bcdelement) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Remove-BcdEntry`](#remove-bcdentry) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Remove-BitsTransfer`](#remove-bitstransfer) | BitsTransfer | 都有 | 无 | 取消 BITS 传输作业。 | 不实现 |  |
| [`Remove-CertificateEnrollmentPolicyServer`](#remove-certificateenrollmentpolicyserver) | PKI | 都有 | 无 | 从当前用户或本地计算机配置删除注册策略服务器及其 URL。 | 不实现 |  |
| [`Remove-CertificateNotificationTask`](#remove-certificatenotificationtask) | PKI | 都有 | 无 | 从任务计划程序删除证书通知任务。 | 不实现 |  |
| [`Remove-CimInstance`](#remove-ciminstance) | CimCmdlets | 都有 | 语法不同 | 从计算机中删除 CIM 实例。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Remove-CimSession`](#remove-cimsession) | CimCmdlets | 都有 | 语法不同 | 删除一个或多个 CIM 会话。 | 不实现 | Windows 专属（CIM/WMI 体系） |
| [`Remove-CIPolicyRule`](#remove-cipolicyrule) | ConfigCI | 都有 | 无 | 此 cmdlet 不受支持。 | 不实现 |  |
| [`Remove-Computer`](#remove-computer) | Microsoft.PowerShell.Management | 都有 | 无 | 从其域中删除本地计算机。 | 不实现 |  |
| [`Remove-EventLog`](#remove-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 删除事件日志或取消注册事件源。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Remove-JobTrigger`](#remove-jobtrigger) | PSScheduledJob | 都有 | 无 | 从计划作业中删除作业触发器。 | 不实现 |  |
| [`Remove-LocalGroup`](#remove-localgroup) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 删除本地安全组。 | 不实现 |  |
| [`Remove-LocalGroupMember`](#remove-localgroupmember) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 从本地组中删除成员。 | 不实现 |  |
| [`Remove-LocalUser`](#remove-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 删除本地用户帐户。 | 不实现 |  |
| [`Remove-OsConfigurationDocument`](#remove-osconfigurationdocument) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Remove-OSConfigurationScenarioDefinition`](#remove-osconfigurationscenariodefinition) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Remove-PmemDedicatedMemory`](#remove-pmemdedicatedmemory) | PersistentMemory | 都有 | 无 | 获取专用持久内存。 | 不实现 |  |
| [`Remove-PmemDisk`](#remove-pmemdisk) | PersistentMemory | 都有 | 无 | 删除持久内存磁盘。 | 不实现 |  |
| [`Remove-PSSnapin`](#remove-pssnapin) | Microsoft.PowerShell.Core | 仅5.1 | 仅5.1提供 | 从当前会话中删除 Windows PowerShell 管理单元。 | 不实现 |  |
| [`Remove-RecoveryManagementPluginAltitude`](#remove-recoverymanagementpluginaltitude) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Remove-Service`](#remove-service) | Microsoft.PowerShell.Management | 仅7 | 仅7提供 | 删除 Windows 服务。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Remove-WindowsCapability`](#remove-windowscapability) | Dism | 都有 | 无 | 从映像卸载 Windows 功能包。 | 不实现 |  |
| [`Remove-WindowsDriver`](#remove-windowsdriver) | Dism | 都有 | 无 | 从脱机 Windows 映像删除驱动程序。 | 不实现 |  |
| [`Remove-WindowsImage`](#remove-windowsimage) | Dism | 都有 | 语法不同 | 从包含多个卷映像的 WIM 文件删除指定卷映像。 | 不实现 |  |
| [`Remove-WindowsPackage`](#remove-windowspackage) | Dism | 都有 | 无 | 从 Windows 映像删除包。 | 不实现 |  |
| [`Remove-WmiObject`](#remove-wmiobject) | Microsoft.PowerShell.Management | 都有 | 无 | 删除现有 Windows Management Instrumentation （WMI） 类的实例。 | 不实现 |  |
| [`Remove-WSManInstance`](#remove-wsmaninstance) | Microsoft.WSMan.Management | 都有 | 无 | 删除管理资源实例。 | 不实现 |  |
| [`Rename-Computer`](#rename-computer) | Microsoft.PowerShell.Management | 都有 | 语法不同 | 重命名计算机。 | 映射 Linux（sudo reboot / shutdown / hostnamectl） |  |
| [`Rename-LocalGroup`](#rename-localgroup) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 重命名本地安全组。 | 不实现 |  |
| [`Rename-LocalUser`](#rename-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 重命名本地用户帐户。 | 不实现 |  |
| [`Repair-AppvClientConnectionGroup`](#repair-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 重置连接组的用户包设置。 | 不实现 |  |
| [`Repair-AppvClientPackage`](#repair-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 重置包的用户设置。 | 不实现 |  |
| [`Repair-UevTemplateIndex`](#repair-uevtemplateindex) | UEV | 仅5.1 | 仅5.1提供 | 修复损坏的 UE-V 模板索引。 | 不实现 |  |
| [`Repair-WindowsImage`](#repair-windowsimage) | Dism | 都有 | 无 | 修复 WIM 或 VHD 文件中的 Windows 映像。 | 不实现 |  |
| [`Reset-AppSharedPackageContainer`](#reset-appsharedpackagecontainer) | Appx | 都有 | 无 | 销毁容器的全部应用数据。 | 不实现 |  |
| [`Reset-AppxPackage`](#reset-appxpackage) | Appx | 都有 | 无 | 将 Windows 应用恢复到初始配置。 | 不实现 |  |
| [`Reset-ComputerMachinePassword`](#reset-computermachinepassword) | Microsoft.PowerShell.Management | 都有 | 无 | 重置计算机的计算机帐户密码。 | 不实现 |  |
| [`Reset-LapsPassword`](#reset-lapspassword) | LAPS | 都有 | 无 | 使 LAPS 立即轮换当前受管本地帐户的密码。 | 不实现 |  |
| [`Resolve-DnsName`](#resolve-dnsname) | DnsClient | 都有 | 无 | 对指定名称执行 DNS 解析。 | 不实现 |  |
| [`Restart-Service`](#restart-service) | Microsoft.PowerShell.Management | 都有 | 无 | 停止并接着启动一个或更多服务。 | 映射 Linux（systemctl start/stop/restart） |  |
| [`Restore-Computer`](#restore-computer) | Microsoft.PowerShell.Management | 都有 | 无 | 在本地计算机上启动系统还原。 | 不实现 |  |
| [`Restore-UevBackup`](#restore-uevbackup) | UEV | 仅5.1 | 仅5.1提供 | 将另一台计算机备份的设置应用到本机。 | 不实现 |  |
| [`Restore-UevUserSetting`](#restore-uevusersetting) | UEV | 仅5.1 | 仅5.1提供 | 为用户设置设置还原标志。 | 不实现 |  |
| [`Resume-BitsTransfer`](#resume-bitstransfer) | BitsTransfer | 都有 | 无 | 恢复 BITS 传输作业。 | 不实现 |  |
| [`Resume-Job`](#resume-job) | Microsoft.PowerShell.Core | 仅5.1 | 仅5.1提供 | 重启挂起的作业。 | 不实现 |  |
| [`Resume-ProvisioningSession`](#resume-provisioningsession) | Provisioning | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Resume-ReFSDedupSchedule`](#resume-refsdedupschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 恢复指定 ReFS 卷上的重复数据删除计划。 | 不实现 |  |
| [`Resume-Service`](#resume-service) | Microsoft.PowerShell.Management | 都有 | 无 | 恢复一个或多个挂起（已暂停）服务。 | 映射 Linux（systemctl start/stop/restart） |  |
| [`Save-OsImage`](#save-osimage) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Save-SoftwareInventory`](#save-softwareinventory) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Save-WindowsImage`](#save-windowsimage) | Dism | 都有 | 无 | 将对已装载映像的更改保存到其 WIM 或 VHD 文件。 | 不实现 |  |
| [`Send-AppvClientReport`](#send-appvclientreport) | AppvClient | 仅5.1 | 仅5.1提供 | 从客户端发送报告数据。 | 不实现 |  |
| [`Send-DtcDiagnosticTransaction`](#send-dtcdiagnostictransaction) | MsDtc | 都有 | 无 | 向指定诊断资源管理器传播事务。 | 不实现 |  |
| [`Set-Acl`](#set-acl) | Microsoft.PowerShell.Security | 都有 | 语法不同 | 更改指定项（如文件或注册表项）的安全描述符。 | 不实现 |  |
| [`Set-AppBackgroundTaskResourcePolicy`](#set-appbackgroundtaskresourcepolicy) | AppBackgroundTask | 都有 | 无 | 配置后台任务对全局池的使用。 | 不实现 |  |
| [`Set-AppLockerPolicy`](#set-applockerpolicy) | AppLocker | 都有 | 无 | 为指定 GPO 设置 AppLocker 策略。 | 不实现 |  |
| [`Set-AppvClientConfiguration`](#set-appvclientconfiguration) | AppvClient | 仅5.1 | 仅5.1提供 | 配置 App-V 客户端的设置。 | 不实现 |  |
| [`Set-AppvClientMode`](#set-appvclientmode) | AppvClient | 仅5.1 | 仅5.1提供 | 设置客户端运行模式。 | 不实现 |  |
| [`Set-AppvClientPackage`](#set-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 配置 App-V 客户端包。 | 不实现 |  |
| [`Set-AppvPublishingServer`](#set-appvpublishingserver) | AppvClient | 仅5.1 | 仅5.1提供 | 修改 App-V 发布服务器属性。 | 不实现 |  |
| [`Set-AppxDefaultVolume`](#set-appxdefaultvolume) | Appx | 都有 | 无 | 指定默认 appx 卷。 | 不实现 |  |
| [`Set-AppxPackageAutoUpdateSettings`](#set-appxpackageautoupdatesettings) | Appx | 都有 | 语法不同 | 配置指定 Windows 应用的自动更新和修复设置。 | 不实现 |  |
| [`Set-AppXProvisionedDataFile`](#set-appxprovisioneddatafile) | Dism | 都有 | 无 | 向 Windows 映像中已预配的指定应用 (.appx) 包添加自定义数据。 | 不实现 |  |
| [`Set-AuthenticodeSignature`](#set-authenticodesignature) | Microsoft.PowerShell.Security | 都有 | 无 | 将验证码签名添加到 PowerShell 脚本或其他文件。 | 不实现 |  |
| [`Set-BcdBootDefault`](#set-bcdbootdefault) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdBootDisplayOrder`](#set-bcdbootdisplayorder) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdBootSequence`](#set-bcdbootsequence) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdBootTimeout`](#set-bcdboottimeout) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdBootToolsDisplayOrder`](#set-bcdboottoolsdisplayorder) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdDebugSettings`](#set-bcddebugsettings) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdElement`](#set-bcdelement) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BcdHypervisorSettings`](#set-bcdhypervisorsettings) | Microsoft.Windows.Bcd.Cmdlets | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-BitsTransfer`](#set-bitstransfer) | BitsTransfer | 都有 | 无 | 修改现有 BITS 传输作业的属性。 | 不实现 |  |
| [`Set-CertificateAutoEnrollmentPolicy`](#set-certificateautoenrollmentpolicy) | PKI | 都有 | 无 | 设置本地证书自动注册策略。 | 不实现 |  |
| [`Set-CimInstance`](#set-ciminstance) | CimCmdlets | 都有 | 语法不同 | 通过调用 CIM 类的 ModifyInstance 方法修改 CIM 服务器上的 CIM 实例。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Set-CIPolicyIdInfo`](#set-cipolicyidinfo) | ConfigCI | 都有 | 无 | 修改代码完整性策略的名称和 ID。 | 不实现 |  |
| [`Set-CIPolicySetting`](#set-cipolicysetting) | ConfigCI | 都有 | 无 | 修改代码完整性策略中的 SecureSettings。 | 不实现 |  |
| [`Set-CIPolicyVersion`](#set-cipolicyversion) | ConfigCI | 都有 | 无 | 更新策略版本号。 | 不实现 |  |
| [`Set-Culture`](#set-culture) | International | 都有 | 无 | 设置当前用户帐户的用户区域。 | 不实现 |  |
| [`Set-DscLocalConfigurationManager`](#set-dsclocalconfigurationmanager) | PSDesiredStateConfiguration | 都有 | 无 | 将本地配置管理器（LCM）设置应用于节点。 | 不实现 |  |
| [`Set-HVCIOptions`](#set-hvcioptions) | ConfigCI | 都有 | 无 | 修改策略的 Hypervisor 代码完整性选项。 | 不实现 |  |
| [`Set-JobTrigger`](#set-jobtrigger) | PSScheduledJob | 都有 | 无 | 更改计划作业的作业触发器。 | 不实现 |  |
| [`Set-KdsConfiguration`](#set-kdsconfiguration) | Kds | 都有 | 无 | 设置 KdsSvc 的配置。 | 不实现 |  |
| [`Set-LapsADAuditing`](#set-lapsadauditing) | LAPS | 都有 | 无 | 配置 AD 组织单位，启用对 LAPS 密码架构属性的审核。 | 不实现 |  |
| [`Set-LapsADComputerSelfPermission`](#set-lapsadcomputerselfpermission) | LAPS | 都有 | 无 | 配置 AD 组织单位权限，允许其中计算机更新 LAPS 密码。 | 不实现 |  |
| [`Set-LapsADPasswordExpirationTime`](#set-lapsadpasswordexpirationtime) | LAPS | 都有 | 无 | 在 AD 计算机或域控制器对象上设置 LAPS 密码过期时间戳。 | 不实现 |  |
| [`Set-LapsADReadPasswordPermission`](#set-lapsadreadpasswordpermission) | LAPS | 都有 | 无 | 配置 AD 组织单位安全，授予指定用户或组查询 LAPS 密码的权限。 | 不实现 |  |
| [`Set-LapsADResetPasswordPermission`](#set-lapsadresetpasswordpermission) | LAPS | 都有 | 无 | 配置 AD 组织单位安全，授予指定用户或组设置 LAPS 密码过期时间的权限。 | 不实现 |  |
| [`Set-LocalGroup`](#set-localgroup) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 更改本地安全组。 | 不实现 |  |
| [`Set-LocalUser`](#set-localuser) | Microsoft.PowerShell.LocalAccounts | 都有 | 无 | 修改本地用户帐户。 | 不实现 |  |
| [`Set-NonRemovableAppsPolicy`](#set-nonremovableappspolicy) | Dism | 都有 | 无 | 将应用包设为不可移除（不可卸载）。 | 不实现 |  |
| [`Set-OsConfigurationDocument`](#set-osconfigurationdocument) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-OsConfigurationProperty`](#set-osconfigurationproperty) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-OSConfigurationScenarioDefinition`](#set-osconfigurationscenariodefinition) | OsConfiguration | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-ProcessMitigation`](#set-processmitigation) | ProcessMitigations | 都有 | 无 | 启用或禁用进程缓解措施，或从 XML 文件批量设置。 | 不实现 |  |
| [`Set-PSSessionConfiguration`](#set-pssessionconfiguration) | Microsoft.PowerShell.Core | 都有 | 无 | 更改已注册会话配置的属性。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Set-RecoveryManagementPluginAltitude`](#set-recoverymanagementpluginaltitude) | Dism | 都有 | 语法不同 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-RecoveryRemoteManagementStatus`](#set-recoveryremotemanagementstatus) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Set-ReFSDedupSchedule`](#set-refsdedupschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 设置指定 ReFS 卷上的重复数据删除计划。 | 不实现 |  |
| [`Set-ReFSDedupScrubSchedule`](#set-refsdedupscrubschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 设置指定 ReFS 卷上的重复数据删除清理计划。 | 不实现 |  |
| [`Set-RuleOption`](#set-ruleoption) | ConfigCI | 都有 | 无 | 修改代码完整性策略中的规则选项。 | 不实现 |  |
| [`Set-ScheduledJob`](#set-scheduledjob) | PSScheduledJob | 都有 | 无 | 更改计划作业。 | 不实现 |  |
| [`Set-ScheduledJobOption`](#set-scheduledjoboption) | PSScheduledJob | 都有 | 无 | 更改计划作业的作业选项。 | 不实现 |  |
| [`Set-SecureBootUEFI`](#set-securebootuefi) | SecureBoot | 都有 | 无 | 设置与安全启动相关的 UEFI 变量。 | 不实现 |  |
| [`Set-Service`](#set-service) | Microsoft.PowerShell.Management | 都有 | 简介不同；语法不同 | 5.1：启动、停止和挂起服务，并更改其属性。 / 7：启动、停止和暂停服务，并更改其属性。 | 映射 Linux（systemctl start/stop/restart） |  |
| [`Set-SystemPreferredUILanguage`](#set-systempreferreduilanguage) | LanguagePackManagement | 都有 | 无 | 将给定语言设为系统首选 UI 语言。 | 不实现 |  |
| [`Set-TimeZone`](#set-timezone) | Microsoft.PowerShell.Management | 都有 | 无 | 将系统时区设置为指定的时区。 | 映射 Linux（timedatectl） |  |
| [`Set-TpmOwnerAuth`](#set-tpmownerauth) | TrustedPlatformModule | 都有 | 语法不同 | 更改 TPM 所有者授权值。 | 不实现 |  |
| [`Set-UevConfiguration`](#set-uevconfiguration) | UEV | 仅5.1 | 仅5.1提供 | 修改 UE-V 配置设置。 | 不实现 |  |
| [`Set-UevTemplateProfile`](#set-uevtemplateprofile) | UEV | 仅5.1 | 仅5.1提供 | 修改模板关联的配置文件。 | 不实现 |  |
| [`Set-WheaMemoryPolicy`](#set-wheamemorypolicy) | Whea | 都有 | 语法不同 | 设置计算机的 WHEA 内存策略。 | 不实现 |  |
| [`Set-WinAcceptLanguageFromLanguageListOptOut`](#set-winacceptlanguagefromlanguagelistoptout) | International | 都有 | 无 | 设置当前用户帐户语言列表退出设置的 HTTP 接受语言。 | 不实现 |  |
| [`Set-WinCultureFromLanguageListOptOut`](#set-winculturefromlanguagelistoptout) | International | 都有 | 无 | 设置当前用户帐户语言列表退出设置的区域。 | 不实现 |  |
| [`Set-WinDefaultInputMethodOverride`](#set-windefaultinputmethodoverride) | International | 都有 | 无 | 设置当前用户帐户的默认输入法替代。 | 不实现 |  |
| [`Set-WindowsEdition`](#set-windowsedition) | Dism | 都有 | 无 | 将 Windows 映像更改为更高版本。 | 不实现 |  |
| [`Set-WindowsProductKey`](#set-windowsproductkey) | Dism | 都有 | 无 | 设置 Windows 映像的产品密钥。 | 不实现 |  |
| [`Set-WindowsReservedStorageState`](#set-windowsreservedstoragestate) | Dism | 都有 | 无 | 设置映像的保留存储状态。 | 不实现 |  |
| [`Set-WindowsSearchSetting`](#set-windowssearchsetting) | WindowsSearch | 都有 | 无 | 修改控制 Windows 搜索的值。 | 不实现 |  |
| [`Set-WinHomeLocation`](#set-winhomelocation) | International | 都有 | 无 | 设置当前用户帐户的主位置。 | 不实现 |  |
| [`Set-WinLanguageBarOption`](#set-winlanguagebaroption) | International | 都有 | 无 | 设置当前用户帐户的语言栏类型和模式。 | 不实现 |  |
| [`Set-WinSystemLocale`](#set-winsystemlocale) | International | 都有 | 无 | 设置当前计算机的系统区域。 | 不实现 |  |
| [`Set-WinUILanguageOverride`](#set-winuilanguageoverride) | International | 都有 | 无 | 设置当前用户帐户的 Windows UI 语言替代。 | 不实现 |  |
| [`Set-WinUserLanguageList`](#set-winuserlanguagelist) | International | 都有 | 无 | 设置当前用户帐户的语言列表及关联属性。 | 不实现 |  |
| [`Set-WmiInstance`](#set-wmiinstance) | Microsoft.PowerShell.Management | 都有 | 无 | 创建或更新现有 Windows Management Instrumentation （WMI） 类的实例。 | 不实现 |  |
| [`Set-WSManInstance`](#set-wsmaninstance) | Microsoft.WSMan.Management | 都有 | 无 | 修改与资源相关的管理信息。 | 不实现 |  |
| [`Set-WSManQuickConfig`](#set-wsmanquickconfig) | Microsoft.WSMan.Management | 都有 | 无 | 配置用于远程管理的本地计算机。 | 不实现 |  |
| [`Show-Command`](#show-command) | Microsoft.PowerShell.Utility | 都有 | 无 | 在图形窗口中显示 PowerShell 命令信息。 | 不实现 | GUI 与打印 |
| [`Show-ControlPanelItem`](#show-controlpanelitem) | Microsoft.PowerShell.Management | 都有 | 无 | 打开控制面板项。 | 不实现 |  |
| [`Show-EventLog`](#show-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 在事件查看器中显示本地或远程计算机的事件日志。 | 不实现 |  |
| [`Show-WindowsDeveloperLicenseRegistration`](#show-windowsdeveloperlicenseregistration) | WindowsDeveloperLicense | 都有 | 无 | 提供如何启用设备进行开发的信息。 | 不实现 |  |
| [`Split-WindowsImage`](#split-windowsimage) | Dism | 都有 | 语法不同 | 将现有 .wim 文件拆分为多个只读拆分 .wim 文件。 | 不实现 |  |
| [`Start-BitsTransfer`](#start-bitstransfer) | BitsTransfer | 都有 | 无 | 创建 BITS 传输作业。 | 不实现 |  |
| [`Start-DscConfiguration`](#start-dscconfiguration) | PSDesiredStateConfiguration | 都有 | 无 | 将配置应用于节点。 | 不实现 |  |
| [`Start-DtcDiagnosticResourceManager`](#start-dtcdiagnosticresourcemanager) | MsDtc | 都有 | 无 | 启动诊断资源管理器。 | 不实现 |  |
| [`Start-OSUninstall`](#start-osuninstall) | Dism | 都有 | 无 | Windows 允许用户卸载并回滚到先前版本，可用 DISM 发起卸载。 | 不实现 |  |
| [`Start-ReFSDedupJob`](#start-refsdedupjob) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 在指定 ReFS 卷启动重复数据删除作业。 | 不实现 |  |
| [`Start-Service`](#start-service) | Microsoft.PowerShell.Management | 都有 | 无 | 启动一个或多个已停止的服务。 | 映射 Linux（systemctl start/stop/restart） |  |
| [`Start-Transaction`](#start-transaction) | Microsoft.PowerShell.Management | 都有 | 无 | 启动事务。 | 不实现 |  |
| [`Stop-AppvClientConnectionGroup`](#stop-appvclientconnectiongroup) | AppvClient | 仅5.1 | 仅5.1提供 | 关闭连接组的共享虚拟环境。 | 不实现 |  |
| [`Stop-AppvClientPackage`](#stop-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 关闭指定包的虚拟环境。 | 不实现 |  |
| [`Stop-DtcDiagnosticResourceManager`](#stop-dtcdiagnosticresourcemanager) | MsDtc | 都有 | 无 | 停止并删除诊断资源管理器作业。 | 不实现 |  |
| [`Stop-ReFSDedupJob`](#stop-refsdedupjob) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 停止指定 ReFS 卷上正在运行的重复数据删除作业。 | 不实现 |  |
| [`Stop-Service`](#stop-service) | Microsoft.PowerShell.Management | 都有 | 无 | 停止一个或多个正在运行的服务。 | 映射 Linux（systemctl start/stop/restart） |  |
| [`Suspend-BitsTransfer`](#suspend-bitstransfer) | BitsTransfer | 都有 | 无 | 挂起 BITS 传输作业。 | 不实现 |  |
| [`Suspend-Job`](#suspend-job) | Microsoft.PowerShell.Core | 仅5.1 | 仅5.1提供 | 暂时停止工作流作业。 | 不实现 |  |
| [`Suspend-ReFSDedupSchedule`](#suspend-refsdedupschedule) | Microsoft.ReFsDedup.Commands | 都有 | 无 | 挂起指定 ReFS 卷上的重复数据删除计划。 | 不实现 |  |
| [`Suspend-Service`](#suspend-service) | Microsoft.PowerShell.Management | 都有 | 无 | 挂起（暂停）一个或多个正在运行的服务。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |
| [`Switch-Certificate`](#switch-certificate) | PKI | 都有 | 无 | 将一个证书标记为被另一证书替换。 | 不实现 |  |
| [`Sync-AppvPublishingServer`](#sync-appvpublishingserver) | AppvClient | 仅5.1 | 仅5.1提供 | 启动 App-V 发布刷新操作。 | 不实现 |  |
| [`Test-AppLockerPolicy`](#test-applockerpolicy) | AppLocker | 都有 | 无 | 指定 AppLocker 策略，判断输入文件是否允许指定用户运行。 | 不实现 |  |
| [`Test-Certificate`](#test-certificate) | PKI | 都有 | 无 | 按输入参数验证证书。 | 不实现 |  |
| [`Test-ComputerSecureChannel`](#test-computersecurechannel) | Microsoft.PowerShell.Management | 都有 | 无 | 测试和修复本地计算机与其域之间的安全通道。 | 不实现 |  |
| [`Test-DscConfiguration`](#test-dscconfiguration) | PSDesiredStateConfiguration | 都有 | 无 | 测试节点上的实际配置是否与所需配置匹配。 | 不实现 |  |
| [`Test-FileCatalog`](#test-filecatalog) | Microsoft.PowerShell.Security | 都有 | 无 | Test-FileCatalog 验证目录文件 （.cat） 中包含的哈希是否与实际文件的哈希匹配，以验证其真实性。 | 不实现 |  |
| [`Test-KdsRootKey`](#test-kdsrootkey) | Kds | 都有 | 无 | 测试根密钥配置。 | 不实现 |  |
| [`Test-PSSessionConfigurationFile`](#test-pssessionconfigurationfile) | Microsoft.PowerShell.Core | 都有 | 无 | 验证会话配置文件中的密钥和值。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Test-UevTemplate`](#test-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 验证模板是否符合 UE-V 架构。 | 不实现 |  |
| [`Test-WSMan`](#test-wsman) | Microsoft.WSMan.Management | 都有 | 无 | 测试 WinRM 服务是在本地还是远程计算机上运行。 | 不实现 |  |
| [`Unblock-Tpm`](#unblock-tpm) | TrustedPlatformModule | 都有 | 无 | 重置 TPM 锁定。 | 不实现 |  |
| [`Undo-DtcDiagnosticTransaction`](#undo-dtcdiagnostictransaction) | MsDtc | 都有 | 无 | 对指定事务调用中止流程。 | 不实现 |  |
| [`Undo-Transaction`](#undo-transaction) | Microsoft.PowerShell.Management | 都有 | 无 | 回滚活动事务。 | 不实现 |  |
| [`Uninstall-Language`](#uninstall-language) | LanguagePackManagement | 都有 | 无 | 从设备卸载语言。 | 不实现 |  |
| [`Uninstall-ProvisioningPackage`](#uninstall-provisioningpackage) | Provisioning | 都有 | 无 | 从本地计算机卸载 .PPKG 包。 | 不实现 |  |
| [`Uninstall-TrustedProvisioningCertificate`](#uninstall-trustedprovisioningcertificate) | Provisioning | 都有 | 无 | 删除以前安装的预配证书。 | 不实现 |  |
| [`Unpublish-AppvClientPackage`](#unpublish-appvclientpackage) | AppvClient | 仅5.1 | 仅5.1提供 | 删除包的扩展点。 | 不实现 |  |
| [`Unregister-PSSessionConfiguration`](#unregister-pssessionconfiguration) | Microsoft.PowerShell.Core | 都有 | 无 | 从计算机中删除已注册的会话配置。 | 不实现 | 远程会话，远程操作不在实现范围内 |
| [`Unregister-RecoveryManagementPlugin`](#unregister-recoverymanagementplugin) | Dism | 都有 | 无 | - | 不实现 | 没有文档准确说明指令的作用 |
| [`Unregister-ScheduledJob`](#unregister-scheduledjob) | PSScheduledJob | 都有 | 无 | 删除本地计算机上的计划作业。 | 不实现 |  |
| [`Unregister-UevTemplate`](#unregister-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 从 UE-V 注销设置位置模板。 | 不实现 |  |
| [`Unregister-WindowsDeveloperLicense`](#unregister-windowsdeveloperlicense) | WindowsDeveloperLicense | 都有 | 无 | 在当前计算机禁用开发者模式。 | 不实现 |  |
| [`Update-LapsADSchema`](#update-lapsadschema) | LAPS | 都有 | 无 | 用 LAPS 架构属性扩展 AD 架构。 | 不实现 |  |
| [`Update-UevTemplate`](#update-uevtemplate) | UEV | 仅5.1 | 仅5.1提供 | 更新 UE-V 中的设置位置模板。 | 不实现 |  |
| [`Update-WIMBootEntry`](#update-wimbootentry) | Dism | 都有 | 无 | 更新与数据源 ID、重命名映像路径或移动映像路径关联的 WIMBoot 配置项。 | 不实现 |  |
| [`Use-Transaction`](#use-transaction) | Microsoft.PowerShell.Management | 都有 | 无 | 将脚本块添加到活动事务。 | 不实现 |  |
| [`Use-WindowsUnattend`](#use-windowsunattend) | Dism | 都有 | 无 | 将无人参与应答文件应用到 Windows 映像。 | 不实现 |  |
| [`Write-EventLog`](#write-eventlog) | Microsoft.PowerShell.Management | 都有 | 无 | 将事件写入事件日志。 | 不实现 | Windows 专属（注册表 / 服务 / 回收站 / 补丁 / 事件） |

## 指令详细说明

### Add-AppProvisionedSharedPackageContainer

版本：都有

模块：Dism

语法：

```powershell
Add-AppProvisionedSharedPackageContainer -DefinitionFile <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-AppProvisionedSharedPackageContainer -DefinitionFile <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Add-AppSharedPackageContainer

版本：都有

模块：Appx

语法：

```powershell
Add-AppSharedPackageContainer [-Path] <string> [-ForceApplicationShutdown] [-Merge] [-RequirePackagesPresent] [-Force] [<CommonParameters>]
```

示例：

```powershell
Add-AppSharedPackageContainer -Path C:\MyFolder\ContosoTestContainer.xml
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/add-appsharedpackagecontainer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/add-appsharedpackagecontainer?view=powershell-7.5)

### Add-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Add-AppvClientConnectionGroup [-Path] <string> [<CommonParameters>]
```

示例：添加连接组

```powershell
PS C:\> Add-AppvClientConnectionGroup -Path "C:\MyApps\MyGroup.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/add-appvclientconnectiongroup?view=powershell-5.1)

### Add-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Add-AppvClientPackage [-Path] <string> [[-DynamicDeploymentConfiguration] <string>] [<CommonParameters>]
```

示例：向客户端添加包

```powershell
PS C:\> Add-AppvClientPackage -Path "http://MyServer/content/package.APPV"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/add-appvclientpackage?view=powershell-5.1)

### Add-AppvPublishingServer

版本：仅5.1

模块：AppvClient

语法：

```powershell
Add-AppvPublishingServer [-Name] <string> [-URL] <string> [[-GlobalRefreshEnabled] <bool>] [[-GlobalRefreshOnLogon] <bool>] [[-GlobalRefreshInterval] <uint32>] [[-GlobalRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [[-UserRefreshEnabled] <bool>] [[-UserRefreshOnLogon] <bool>] [[-UserRefreshInterval] <uint32>] [[-UserRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/add-appvpublishingserver?view=powershell-5.1)

### Add-AppxPackage

版本：都有

模块：Appx

语法：

```powershell
Add-AppxPackage [-Path] <string> [-DependencyPath <string[]>] [-RequiredContentGroupOnly] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-RetainFilesOnFailure] [-InstallAllResources] [-Volume <AppxVolume>] [-ExternalPackages <string[]>] [-OptionalPackages <string[]>] [-RelatedPackages <string[]>] [-ExternalLocation <string>] [-DeferRegistrationWhenPackagesAreInUse] [-StubPackageOption <StubPackageOption>] [-AllowUnsigned] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -AppInstallerFile [-RequiredContentGroupOnly] [-ForceTargetApplicationShutdown] [-InstallAllResources] [-LimitToExistingPackages] [-Volume <AppxVolume>] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -Register [-DependencyPath <string[]>] [-DisableDevelopmentMode] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-InstallAllResources] [-ExternalLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -Update [-DependencyPath <string[]>] [-RequiredContentGroupOnly] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-RetainFilesOnFailure] [-InstallAllResources] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage [-Path] <string> -Stage [-DependencyPath <string[]>] [-RequiredContentGroupOnly] [-ForceUpdateFromAnyVersion] [-Volume <AppxVolume>] [-ExternalPackages <string[]>] [-OptionalPackages <string[]>] [-RelatedPackages <string[]>] [-ExternalLocation <string>] [-StubPackageOption <StubPackageOption>] [-AllowUnsigned] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage -MainPackage <string> [-Register] [-DependencyPackages <string[]>] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-ForceUpdateFromAnyVersion] [-InstallAllResources] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-AppxPackage -RegisterByFamilyName -MainPackage <string> [-DependencyPackages <string[]>] [-ForceApplicationShutdown] [-ForceTargetApplicationShutdown] [-InstallAllResources] [-OptionalPackages <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：添加应用包

```powershell
Add-AppxPackage -Path '.\MyApp.msix' -DependencyPath '.\winjs.msix'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/add-appxpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/add-appxpackage?view=powershell-7.5)

### Add-AppxProvisionedPackage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Add-AppxProvisionedPackage -Path <string> [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint32>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-AppxProvisionedPackage -Online [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint32>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Add-AppxProvisionedPackage -Path <string> [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-AppxProvisionedPackage -Online [-FolderPath <string>] [-PackagePath <string>] [-DependencyPackagePath <string[]>] [-OptionalPackagePath <string[]>] [-LicensePath <string[]>] [-SkipLicense] [-CustomDataPath <string>] [-Regions <string>] [-StubPackageOption <StubPackageOption>] [-FeatureID <uint>] [-ExternalLocationPath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：向正在运行的操作系统添加应用包

```powershell
PS C:\> Add-AppxProvisionedPackage -Online -FolderPath "c:\Appx"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-appxprovisionedpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-appxprovisionedpackage?view=powershell-7.5)

### Add-AppxVolume

版本：都有

模块：Appx

语法：

```powershell
Add-AppxVolume [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：添加卷

```powershell
Add-AppxVolume -Path "E:\WindowsApps"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/add-appxvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/add-appxvolume?view=powershell-7.5)

### Add-BitsFile

版本：都有

模块：BitsTransfer

语法：

```powershell
Add-BitsFile [-BitsJob] <BitsJob[]> [-Source] <string[]> [[-Destination] <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：向现有 BITS 传输作业的传输队列追加文件

```powershell
PS C:\> Get-BitsTransfer -JobId 10778CFA-C1D7-4A82-8A9D-80B19224879C | Add-BitsFile -Source http://server01/servertestdir/testfile1.txt -Destination "c:\clienttestdir\testfile1.txt"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/add-bitsfile?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/add-bitsfile?view=powershell-7.5)

### Add-CertificateEnrollmentPolicyServer

版本：都有

模块：PKI

语法：

```powershell
Add-CertificateEnrollmentPolicyServer -Url <uri> -context <Context> [-NoClobber] [-RequireStrongValidation] [-Credential <PkiCredential>] [-AutoEnrollmentEnabled] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Add-CertificateEnrollmentPolicyServer -Url $url -Context Machine
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/add-certificateenrollmentpolicyserver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/add-certificateenrollmentpolicyserver?view=powershell-7.5)

### Add-Computer

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Add-Computer [-DomainName] <string> -Credential <pscredential> [-ComputerName <string[]>] [-LocalCredential <pscredential>] [-UnjoinDomainCredential <pscredential>] [-OUPath <string>] [-Server <string>] [-Unsecure] [-Options <JoinOptions>] [-Restart] [-PassThru] [-NewName <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-Computer [-WorkgroupName] <string> [-ComputerName <string[]>] [-LocalCredential <pscredential>] [-Credential <pscredential>] [-Restart] [-PassThru] [-NewName <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将本地计算机添加到域，然后重新启动计算机

```powershell
Add-Computer -DomainName Domain01 -Restart
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/add-computer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/add-computer?view=powershell-7.5)

### Add-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
Add-JobTrigger [-InputObject] <ScheduledJobDefinition[]> [-Trigger] <ScheduledJobTrigger[]> [<CommonParameters>]
Add-JobTrigger [-Id] <int[]> [-Trigger] <ScheduledJobTrigger[]> [<CommonParameters>]
Add-JobTrigger [-Name] <string[]> [-Trigger] <ScheduledJobTrigger[]> [<CommonParameters>]
```

示例：将作业触发器添加到计划作业

```powershell
$Daily = New-JobTrigger -Daily -At 3AMPS
Add-JobTrigger -Trigger $Daily -Name "TestJob"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/add-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/add-jobtrigger?view=powershell-7.5)

### Add-KdsRootKey

版本：都有

模块：Kds

语法：

```powershell
Add-KdsRootKey [[-EffectiveTime] <datetime>] [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
Add-KdsRootKey -EffectiveImmediately [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：生成新根密钥

```powershell
PS C:\> Add-KdsRootKey
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/kds/add-kdsrootkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/kds/add-kdsrootkey?view=powershell-7.5)

### Add-LocalGroupMember

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Add-LocalGroupMember [-Group] <LocalGroup> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Add-LocalGroupMember [-Name] <string> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Add-LocalGroupMember [-SID] <SecurityIdentifier> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将成员添加到管理员组

```powershell
Add-LocalGroupMember -Group "Administrators" -Member "Admin02", "MicrosoftAccount\username@Outlook.com", "AzureAD\DavidChew@contoso.com", "CONTOSO\Domain Admins"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/add-localgroupmember?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/add-localgroupmember?view=powershell-7.5)

### Add-PSSnapin

版本：仅5.1

模块：Microsoft.PowerShell.Core

语法：

```powershell
Add-PSSnapin [-Name] <string[]> [-PassThru] [<CommonParameters>]
```

示例：添加管理单元

```powershell
PS C:\> Add-PSSnapin -Name Microsoft.Exchange, Microsoft.Windows.AD
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/add-pssnapin?view=powershell-5.1)

### Add-SignerRule

版本：都有

模块：ConfigCI

语法：

```powershell
Add-SignerRule -FilePath <string> -CertificatePath <string> [-Kernel] [-User] [-Update] [-Supplemental] [-Deny] [<CommonParameters>]
Add-SignerRule -FilePath <string> -CertStorePath <string> [-Kernel] [-User] [-Update] [-Supplemental] [-Deny] [<CommonParameters>]
```

示例：为用户模式创建并添加签名者规则

```powershell
PS C:\> Add-SignerRule -FilePath '.\Policy.xml' -CertificatePath '.\certificate07.cer' -User
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/add-signerrule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/add-signerrule?view=powershell-7.5)

### Add-WindowsCapability

版本：都有

模块：Dism

语法：

```powershell
Add-WindowsCapability -Name <string> -Online [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-WindowsCapability -Name <string> -Path <string> [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：通过 Windows 更新客户端向正在运行的操作系统添加 Windows 功能包

```powershell
PS C:\> Add-WindowsCapability -Online -Name "Msix.PackagingTool.Driver~~~~0.0.1.0"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowscapability?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowscapability?view=powershell-7.5)

### Add-WindowsDriver

版本：都有

模块：Dism

语法：

```powershell
Add-WindowsDriver -Path <string> [-Recurse] [-ForceUnsigned] [-Driver <string>] [-BasicDriverObject <BasicDriverObject>] [-AdvancedDriverObject <AdvancedDriverObject>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：向映像添加驱动程序

```powershell
PS C:\> Add-WindowsDriver -Path "c:\offline" -Driver "c:\test\drivers" -Recurse
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowsdriver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowsdriver?view=powershell-7.5)

### Add-WindowsImage

版本：都有

模块：Dism

语法：

```powershell
Add-WindowsImage -ImagePath <string> -CapturePath <string> [-ConfigFilePath <string>] [-Description <string>] [-Name <string>] [-CheckIntegrity] [-NoRpFix] [-Setbootable] [-Verify] [-WIMBoot] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：向映像添加文件

```powershell
PS C:\> Add-WindowsImage -ImagePath "C:\imagestore\custom.wim" -CapturePath d:\ -Name "Drive D"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowsimage?view=powershell-7.5)

### Add-WindowsPackage

版本：都有

模块：Dism

语法：

```powershell
Add-WindowsPackage -PackagePath <string> -Online [-IgnoreCheck] [-PreventPending] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Add-WindowsPackage -PackagePath <string> -Path <string> [-IgnoreCheck] [-PreventPending] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：向联机映像添加包

```powershell
PS C:\> Add-WindowsPackage -Online -PackagePath "c:\packages\package.cab"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowspackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/add-windowspackage?view=powershell-7.5)

### Checkpoint-Computer

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Checkpoint-Computer [-Description] <string> [[-RestorePointType] <string>] [<CommonParameters>]
```

示例：创建系统还原点

```powershell
Checkpoint-Computer -Description "Install MyApp"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/checkpoint-computer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/checkpoint-computer?view=powershell-7.5)

### Clear-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Clear-EventLog [-LogName] <string[]> [[-ComputerName] <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：从本地计算机清除特定事件日志类型

```powershell
Clear-EventLog "Windows PowerShell"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/clear-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/clear-eventlog?view=powershell-7.5)

### Clear-KdsCache

版本：都有

模块：Kds

语法：

```powershell
Clear-KdsCache [-CacheOwnerSid <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：清除组密钥缓存

```powershell
PS C:\> Clear-KdsCache
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/kds/clear-kdscache?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/kds/clear-kdscache?view=powershell-7.5)

### Clear-RecycleBin

版本：仅7

模块：Microsoft.PowerShell.Management

语法：

```powershell
Clear-RecycleBin [[-DriveLetter] <string[]>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：清除所有回收站

```powershell
Clear-RecycleBin
```

出处：[官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/clear-recyclebin?view=powershell-7.5)

### Clear-Recyclebin

版本：仅5.1

模块：Microsoft.PowerShell.Management

语法：

```powershell
Clear-RecycleBin [[-DriveLetter] <string[]>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：清除所有回收站

```powershell
Clear-RecycleBin
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/clear-recyclebin?view=powershell-5.1)

### Clear-ReFSDedupSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Clear-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Clear-ReFSDedupSchedule -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/clear-refsdedupschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/clear-refsdedupschedule?view=powershell-7.5)

### Clear-ReFSDedupScrubSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Clear-ReFSDedupScrubSchedule [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Clear-ReFSDedupScrubSchedule -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/clear-refsdedupscrubschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/clear-refsdedupscrubschedule?view=powershell-7.5)

### Clear-Tpm

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Clear-Tpm [[-OwnerAuthorization] <string>] [-UsePPI] [<CommonParameters>]
Clear-Tpm -File <string> [<CommonParameters>]
```

示例：重置 TPM

```powershell
Clear-Tpm
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/clear-tpm?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/clear-tpm?view=powershell-7.5)

### Clear-UevAppxPackage

版本：仅5.1

模块：UEV

语法：

```powershell
Clear-UevAppxPackage [-PackageFamilyName] <string[]> [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevAppxPackage [-PackageFamilyName] <string[]> -Computer [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevAppxPackage -Computer -All [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevAppxPackage -All [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除 Windows 8 应用

```powershell
PS C:\>Clear-UevAppxPackage -Computer -All
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/clear-uevappxpackage?view=powershell-5.1)

### Clear-UevConfiguration

版本：仅5.1

模块：UEV

语法：

```powershell
Clear-UevConfiguration [-CurrentComputerUser] [-MaxPackageSizeInBytes] [-SettingsStoragePath] [-SyncProviderPingEnabled] [-SyncTimeoutInMilliseconds] [-SyncMethod] [-SyncEnabled] [-SyncOverMeteredNetwork] [-SyncOverMeteredNetworkWhenRoaming] [-SettingsImportNotifyEnabled] [-SettingsImportNotifyDelayInSeconds] [-DontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds] [-WaitForSyncOnApplicationStart] [-WaitForSyncOnLogon] [-SyncUnlistedWindows8Apps] [-VdiCollectionName] [-WhatIf] [-Confirm] [<CommonParameters>]
Clear-UevConfiguration [-Computer] [-MaxPackageSizeInBytes] [-SettingsStoragePath] [-SettingsTemplateCatalogPath] [-SyncProviderPingEnabled] [-SyncTimeoutInMilliseconds] [-SyncMethod] [-SyncEnabled] [-SyncOverMeteredNetwork] [-SyncOverMeteredNetworkWhenRoaming] [-SettingsImportNotifyEnabled] [-SettingsImportNotifyDelayInSeconds] [-ContactITUrl] [-ContactITDescription] [-TrayIconEnabled] [-FirstUseNotificationEnabled] [-DontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds] [-WaitForSyncOnApplicationStart] [-WaitForSyncOnLogon] [-SyncUnlistedWindows8Apps] [-VdiCollectionName] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：清除全部用户的包大小上限设置

```powershell
PS C:\> Clear-UevConfiguration -Computer -MaxPackageSizeInBytes
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/clear-uevconfiguration?view=powershell-5.1)

### Clear-WindowsCorruptMountPoint

版本：都有

模块：Dism

语法：

```powershell
Clear-WindowsCorruptMountPoint [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：删除与已装载映像关联的全部资源

```powershell
PS C:\> Clear-WindowsCorruptMountPoint
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/clear-windowscorruptmountpoint?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/clear-windowscorruptmountpoint?view=powershell-7.5)

### Complete-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Complete-BitsTransfer [-BitsJob] <BitsJob[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：完成当前用户拥有的全部 BITS 传输作业

```powershell
C:\PS>Get-BitsTransfer | Complete-BitsTransfer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/complete-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/complete-bitstransfer?view=powershell-7.5)

### Complete-DtcDiagnosticTransaction

版本：都有

模块：MsDtc

语法：

```powershell
Complete-DtcDiagnosticTransaction [-Transaction] <DtcDiagnosticTransaction> [<CommonParameters>]
```

示例：完成 DTC 诊断事务

```powershell
PS C:\> $Tx = New-DtcDiagnosticTransaction
PS C:\> Complete-DtcDiagnosticTransaction -Transaction $Tx
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/complete-dtcdiagnostictransaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/complete-dtcdiagnostictransaction?view=powershell-7.5)

### Complete-Transaction

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Complete-Transaction [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：提交事务

```powershell
Set-Location HKCU:\software
Start-Transaction
New-Item MyCompany -UseTransaction
Get-ChildItem m*
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/complete-transaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/complete-transaction?view=powershell-7.5)

### Confirm-SecureBootUEFI

版本：都有

模块：SecureBoot

语法：

```powershell
Confirm-SecureBootUEFI [<CommonParameters>]
```

示例：确认安全启动

```powershell
PS C:\> Confirm-SecureBootUEFI
True
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/confirm-securebootuefi?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/confirm-securebootuefi?view=powershell-7.5)

### Connect-PSSession

版本：都有

模块：Microsoft.PowerShell.Core

语法：

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

示例：重新连接到会话

```powershell
Connect-PSSession -ComputerName Server01 -Name ITTask
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/connect-pssession?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/connect-pssession?view=powershell-7.5)

### Connect-WSMan

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Connect-WSMan [[-ComputerName] <string>] [-ApplicationName <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Connect-WSMan [-ConnectionURI <uri>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：连接到远程计算机

```powershell
PS C:\> Connect-WSMan -ComputerName "server01"
PS C:\> cd WSMan:
PS WSMan:\>
PS WSMan:\> dir
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/connect-wsman?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/connect-wsman?view=powershell-7.5)

### Convert-String

版本：仅5.1

模块：Microsoft.PowerShell.Utility

语法：

```powershell
Convert-String -InputObject <string> [-Example <List[psobject]>] [<CommonParameters>]
```

示例：转换字符串的格式

```powershell
"Mu Han", "Jim Hance", "David Ahs", "Kim Akers" | Convert-String -Example "Ed Wilson=Wilson, E."
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/convert-string?view=powershell-5.1)

### ConvertFrom-CIPolicy

版本：都有

模块：ConfigCI

语法：

```powershell
ConvertFrom-CIPolicy [-XmlFilePath] <string> [-BinaryFilePath] <string> [<CommonParameters>]
```

示例：转换策略

```powershell
PS C:\> ConvertFrom-CIPolicy -XmlFilePath ".\Policy03.xml" -BinaryFilePath "Policy03.bin"
C:\Policies\Policy03.bin
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/convertfrom-cipolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/convertfrom-cipolicy?view=powershell-7.5)

### ConvertFrom-SddlString

版本：仅7

模块：Microsoft.PowerShell.Utility

语法：

```powershell
ConvertFrom-SddlString [-Sddl] <string> [-Type <ConvertFromSddlStringCommand+AccessRightTypeNames>] [<CommonParameters>]
```

示例：将文件系统访问权限 SDDL 转换为 PSCustomObject

```powershell
$acl = Get-Acl -Path C:\Windows
ConvertFrom-SddlString -Sddl $acl.Sddl
```

出处：[官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/convertfrom-sddlstring?view=powershell-7.5)

### ConvertFrom-String

版本：仅5.1

模块：Microsoft.PowerShell.Utility

语法：

```powershell
ConvertFrom-String [-InputObject] <string> [-Delimiter <string>] [-PropertyNames <string[]>] [<CommonParameters>]
ConvertFrom-String [-InputObject] <string> [-TemplateFile <string[]>] [-TemplateContent <string[]>] [-IncludeExtent] [-UpdateTemplate] [<CommonParameters>]
```

示例：生成具有默认属性名称的对象

```powershell
"Hello World" | ConvertFrom-String
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/convertfrom-string?view=powershell-5.1)

### ConvertTo-ProcessMitigationPolicy

版本：都有

模块：ProcessMitigations

语法：

```powershell
ConvertTo-ProcessMitigationPolicy [-EMETFilePath] <string> [-OutputFilePath] <string> [<CommonParameters>]
```

示例：

```powershell
PS C:\> ConvertTo-ProcessMitigationPolicy -EMETFilePath policy.xml -OutputFilePath result.xml
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/processmitigations/convertto-processmitigationpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/processmitigations/convertto-processmitigationpolicy?view=powershell-7.5)

### ConvertTo-TpmOwnerAuth

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
ConvertTo-TpmOwnerAuth [-PassPhrase] <string> [<CommonParameters>]
```

示例：转换为所有者授权值

```powershell
PS C:\> ConvertTo-TpmOwnerAuth -PassPhrase "Saturn1977&&"
puJvGK4O6Qvl0loP8r1bIxipDVo=
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/convertto-tpmownerauth?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/convertto-tpmownerauth?view=powershell-7.5)

### Copy-BcdEntry

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Copy-BcdEntry [-SourceEntryId] <string> -TargetStore <BcdStoreInfo[]> [-Description <string>] [-SourceStore <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Copy-BcdEntry [-SourceEntry] <BcdEntryInfo> -TargetStore <BcdStoreInfo[]> [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/copy-bcdentry?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/copy-bcdentry?view=powershell-7.5)

### Copy-UserInternationalSettingsToSystem

版本：都有

模块：International

语法：

```powershell
Copy-UserInternationalSettingsToSystem [-NewUser] <bool> [<CommonParameters>]
```

示例：将设置复制到欢迎屏幕、系统帐户和新用户帐户

```powershell
PS C:\> Copy-UserInternationalSettingsToSystem -WelcomeScreen $True -NewUser $True
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/copy-userinternationalsettingstosystem?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/copy-userinternationalsettingstosystem?view=powershell-7.5)

### Disable-AppBackgroundTaskDiagnosticLog

版本：都有

模块：AppBackgroundTask

语法：

```powershell
Disable-AppBackgroundTaskDiagnosticLog [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用后台任务日志

```powershell
PS C:\> Disable-AppBackgroundTaskDiagnosticLog
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appbackgroundtask/disable-appbackgroundtaskdiagnosticlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appbackgroundtask/disable-appbackgroundtaskdiagnosticlog?view=powershell-7.5)

### Disable-Appv

版本：仅5.1

模块：AppvClient

语法：

```powershell
Disable-Appv [<CommonParameters>]
```

示例：禁用 App-V 服务

```powershell
PS C:\> Disable-Appv
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/disable-appv?view=powershell-5.1)

### Disable-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Disable-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [-UserSID <string>] [<CommonParameters>]
Disable-AppvClientConnectionGroup [-Name] <string> [-Global] [-UserSID <string>] [<CommonParameters>]
Disable-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [-UserSID <string>] [<CommonParameters>]
```

示例：按名称禁用连接组

```powershell
PS C:\> Disable-AppvClientConnectionGroup -Name "MyGroup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/disable-appvclientconnectiongroup?view=powershell-5.1)

### Disable-BcdElementBootDebug

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Disable-BcdElementBootDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementBootDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementbootdebug?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementbootdebug?view=powershell-7.5)

### Disable-BcdElementBootEms

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Disable-BcdElementBootEms [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementBootEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementBootEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementbootems?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementbootems?view=powershell-7.5)

### Disable-BcdElementDebug

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Disable-BcdElementDebug [[-Id] <string>] [-Store <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementdebug?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementdebug?view=powershell-7.5)

### Disable-BcdElementEms

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Disable-BcdElementEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementems?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementems?view=powershell-7.5)

### Disable-BcdElementEventLogging

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Disable-BcdElementEventLogging [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementEventLogging [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementEventLogging [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementeventlogging?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementeventlogging?view=powershell-7.5)

### Disable-BcdElementHypervisorDebug

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Disable-BcdElementHypervisorDebug [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementHypervisorDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-BcdElementHypervisorDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementhypervisordebug?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/disable-bcdelementhypervisordebug?view=powershell-7.5)

### Disable-ComputerRestore

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Disable-ComputerRestore [-Drive] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用指定驱动器上的系统还原

```powershell
Disable-ComputerRestore -Drive "C:\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/disable-computerrestore?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/disable-computerrestore?view=powershell-7.5)

### Disable-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
Disable-JobTrigger [-InputObject] <ScheduledJobTrigger[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用作业触发器

```powershell
PS C:\> Get-JobTrigger -Name "Backup-Archives" -TriggerId 1 | Disable-JobTrigger
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/disable-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/disable-jobtrigger?view=powershell-7.5)

### Disable-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Disable-LocalUser [-InputObject] <LocalUser[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-LocalUser [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-LocalUser [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：通过指定名称禁用帐户

```powershell
Disable-LocalUser -Name "Admin02"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/disable-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/disable-localuser?view=powershell-7.5)

### Disable-PSRemoting

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Disable-PSRemoting [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例（5.1）：阻止远程访问所有会话配置

```powershell
Disable-PSRemoting
```

示例（7）：阻止远程访问所有 PowerShell 会话配置

```powershell
Disable-PSRemoting
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/disable-psremoting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/disable-psremoting?view=powershell-7.5)

### Disable-PSSessionConfiguration

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Disable-PSSessionConfiguration [[-Name] <string[]>] [-Force] [-NoServiceRestart] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用默认配置

```powershell
Disable-PSSessionConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/disable-pssessionconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/disable-pssessionconfiguration?view=powershell-7.5)

### Disable-ReFSDedup

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Disable-ReFSDedup [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Disable-ReFSDedup -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/disable-refsdedup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/disable-refsdedup?view=powershell-7.5)

### Disable-ScheduledJob

版本：都有

模块：PSScheduledJob

语法：

```powershell
Disable-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-ScheduledJob [-Id] <int> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-ScheduledJob [-Name] <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用计划作业

```powershell
Disable-ScheduledJob -Id 2 -PassThru
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/disable-scheduledjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/disable-scheduledjob?view=powershell-7.5)

### Disable-TlsCipherSuite

版本：都有

模块：TLS

语法：

```powershell
Disable-TlsCipherSuite [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用密码套件

```powershell
Disable-TlsCipherSuite -Name 'TLS_RSA_WITH_3DES_EDE_CBC_SHA'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/disable-tlsciphersuite?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/disable-tlsciphersuite?view=powershell-7.5)

### Disable-TlsEccCurve

版本：都有

模块：TLS

语法：

```powershell
Disable-TlsEccCurve [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Disable-TlsEccCurve -Name curve25519
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/disable-tlsecccurve?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/disable-tlsecccurve?view=powershell-7.5)

### Disable-TlsSessionTicketKey

版本：都有

模块：TLS

语法：

```powershell
Disable-TlsSessionTicketKey [-ServiceAccountName] <NTAccount> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用 TLS 会话票证密钥

```powershell
Disable-TlsSessionTicketKey -ServiceAccountName NetworkService
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/disable-tlssessionticketkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/disable-tlssessionticketkey?view=powershell-7.5)

### Disable-TpmAutoProvisioning

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Disable-TpmAutoProvisioning [-OnlyForNextRestart] [<CommonParameters>]
```

示例：禁用自动预配

```powershell
PS C:\> Disable-TpmAutoProvisioning
TpmReady : False
TpmPresent : True
ManagedAuthLevel : Full
OwnerAuth : OwnerClearDisabled : True
AutoProvisioning : Disabled
LockedOut : False
SelfTest : {191, 191, 245, 191...}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/disable-tpmautoprovisioning?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/disable-tpmautoprovisioning?view=powershell-7.5)

### Disable-Uev

版本：仅5.1

模块：UEV

语法：

```powershell
Disable-Uev [<CommonParameters>]
```

示例：禁用 UE-V 服务

```powershell
PS C:\>Disable-Uev
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/disable-uev?view=powershell-5.1)

### Disable-UevAppxPackage

版本：仅5.1

模块：UEV

语法：

```powershell
Disable-UevAppxPackage [-PackageFamilyName] <string[]> [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
Disable-UevAppxPackage [-PackageFamilyName] <string[]> -Computer [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用 Windows 8 应用的同步

```powershell
PS C:\>Disable-UevAppxPackage -Computer -PackageFamilyName "Microsoft.BingFinance"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/disable-uevappxpackage?view=powershell-5.1)

### Disable-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Disable-UevTemplate [-ID] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用指定模板

```powershell
PS C:\> Disable-UevTemplate -ID "MicrosoftCalculator6"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/disable-uevtemplate?view=powershell-5.1)

### Disable-WindowsErrorReporting

版本：都有

模块：WindowsErrorReporting

语法：

```powershell
Disable-WindowsErrorReporting [<CommonParameters>]
```

示例：禁用 Windows 错误报告

```powershell
PS C:\> Disable-WindowsErrorReporting
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowserrorreporting/disable-windowserrorreporting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowserrorreporting/disable-windowserrorreporting?view=powershell-7.5)

### Disable-WindowsOptionalFeature

版本：都有

模块：Dism

语法：

```powershell
Disable-WindowsOptionalFeature -FeatureName <string[]> -Online [-PackageName <string>] [-Remove] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Disable-WindowsOptionalFeature -FeatureName <string[]> -Path <string> [-PackageName <string>] [-Remove] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：禁用可选功能

```powershell
PS C:\> Disable-WindowsOptionalFeature -Online -FeatureName "Hearts"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/disable-windowsoptionalfeature?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/disable-windowsoptionalfeature?view=powershell-7.5)

### Disable-WSManCredSSP

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Disable-WSManCredSSP [-Role] <string> [<CommonParameters>]
```

示例：在客户端上禁用 CredSSP

```powershell
Disable-WSManCredSSP -Role Client
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/disable-wsmancredssp?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/disable-wsmancredssp?view=powershell-7.5)

### Disconnect-PSSession

版本：都有

模块：Microsoft.PowerShell.Core

语法（5.1）：

```powershell
Disconnect-PSSession [-Session] <PSSession[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -InstanceId <guid[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -Name <string[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession [-Id] <int[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Disconnect-PSSession [-Session] <PSSession[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -Name <string[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession [-Id] <int[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Disconnect-PSSession -InstanceId <guid[]> [-IdleTimeoutSec <int>] [-OutputBufferingMode <OutputBufferingMode>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按名称断开会话的连接

```powershell
PS> Disconnect-PSSession -Name UpdateSession
Id Name ComputerName State ConfigurationName Availability
-- ---- ------------ ----- ----------------- ------------
1 UpdateSession Server01 Disconnected Microsoft.PowerShell None
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/disconnect-pssession?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/disconnect-pssession?view=powershell-7.5)

### Disconnect-WSMan

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Disconnect-WSMan [[-ComputerName] <string>] [<CommonParameters>]
```

示例：删除与远程计算机的连接

```powershell
PS C:\> Disconnect-WSMan -Computer server01
PS C:\> cd WSMan:
PS WSMan:\> dir
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/disconnect-wsman?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/disconnect-wsman?view=powershell-7.5)

### Dismount-AppxVolume

版本：都有

模块：Appx

语法：

```powershell
Dismount-AppxVolume [-Volume] <AppxVolume[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按路径卸载卷

```powershell
Dismount-AppxVolume -Volume E:\
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/dismount-appxvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/dismount-appxvolume?view=powershell-7.5)

### Dismount-WindowsImage

版本：都有

模块：Dism

语法：

```powershell
Dismount-WindowsImage -Path <string> -Discard [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Dismount-WindowsImage -Path <string> -Save [-CheckIntegrity] [-Append] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：卸载操作系统映像

```powershell
PS C:\> Dismount-WindowsImage -Path "c:\offline" -Save
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/dismount-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/dismount-windowsimage?view=powershell-7.5)

### Edit-CIPolicyRule

版本：都有

模块：ConfigCI

语法：

```powershell
Edit-CIPolicyRule [-Id] <string> -FilePath <string> [-Name <string>] [-RType <string>] [-FileName <string>] [-Version <string>] [-HashPath <string>] [<CommonParameters>]
Edit-CIPolicyRule [-Id] <string> -FilePath <string> [-Name <string>] [-RType <string>] [-Root <string>] [-AddEkus <string[]>] [-RemoveEkus <string[]>] [-Issuer <string>] [-Publisher <string>] [-OemId <string>] [-AddExceptions <string[]>] [-RemoveExceptions <string[]>] [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/edit-cipolicyrule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/edit-cipolicyrule?view=powershell-7.5)

### Enable-AppBackgroundTaskDiagnosticLog

版本：都有

模块：AppBackgroundTask

语法：

```powershell
Enable-AppBackgroundTaskDiagnosticLog [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启用后台任务日志

```powershell
PS C:\> Enable-AppBackgroundTaskDiagnosticLog
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appbackgroundtask/enable-appbackgroundtaskdiagnosticlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appbackgroundtask/enable-appbackgroundtaskdiagnosticlog?view=powershell-7.5)

### Enable-Appv

版本：仅5.1

模块：AppvClient

语法：

```powershell
Enable-Appv [<CommonParameters>]
```

示例：启用服务

```powershell
PS C:\> Enable-Appv
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/enable-appv?view=powershell-5.1)

### Enable-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Enable-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [-UserSID <string>] [<CommonParameters>]
Enable-AppvClientConnectionGroup [-Name] <string> [-Global] [-UserSID <string>] [<CommonParameters>]
Enable-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [-UserSID <string>] [<CommonParameters>]
```

示例：按名称启用连接组

```powershell
PS C:\> Enable-AppvClientConnectionGroup -Name "MyGroup" -Global
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/enable-appvclientconnectiongroup?view=powershell-5.1)

### Enable-BcdElementBootDebug

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Enable-BcdElementBootDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementBootDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementbootdebug?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementbootdebug?view=powershell-7.5)

### Enable-BcdElementBootEms

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Enable-BcdElementBootEms [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementBootEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementBootEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementbootems?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementbootems?view=powershell-7.5)

### Enable-BcdElementDebug

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Enable-BcdElementDebug [[-Id] <string>] [-Store <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementdebug?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementdebug?view=powershell-7.5)

### Enable-BcdElementEms

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Enable-BcdElementEms [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementEms [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementems?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementems?view=powershell-7.5)

### Enable-BcdElementEventLogging

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Enable-BcdElementEventLogging [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementEventLogging [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementEventLogging [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementeventlogging?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementeventlogging?view=powershell-7.5)

### Enable-BcdElementHypervisorDebug

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Enable-BcdElementHypervisorDebug [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementHypervisorDebug [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-BcdElementHypervisorDebug [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementhypervisordebug?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/enable-bcdelementhypervisordebug?view=powershell-7.5)

### Enable-ComputerRestore

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Enable-ComputerRestore [-Drive] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：在指定的驱动器上启用系统还原

```powershell
Enable-ComputerRestore -Drive "C:\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/enable-computerrestore?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/enable-computerrestore?view=powershell-7.5)

### Enable-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
Enable-JobTrigger [-InputObject] <ScheduledJobTrigger[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启用作业触发器

```powershell
Get-JobTrigger -Name Backup-Archives -TriggerId 1 | Enable-JobTrigger
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/enable-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/enable-jobtrigger?view=powershell-7.5)

### Enable-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Enable-LocalUser [-InputObject] <LocalUser[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-LocalUser [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-LocalUser [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：通过指定名称启用帐户

```powershell
Enable-LocalUser -Name "Admin02"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/enable-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/enable-localuser?view=powershell-7.5)

### Enable-PSRemoting

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Enable-PSRemoting [-Force] [-SkipNetworkProfileCheck] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将计算机配置为接收远程命令

```powershell
Enable-PSRemoting
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/enable-psremoting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/enable-psremoting?view=powershell-7.5)

### Enable-PSSessionConfiguration

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Enable-PSSessionConfiguration [[-Name] <string[]>] [-Force] [-SecurityDescriptorSddl <string>] [-SkipNetworkProfileCheck] [-NoServiceRestart] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：重新启用默认会话

```powershell
Enable-PSSessionConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/enable-pssessionconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/enable-pssessionconfiguration?view=powershell-7.5)

### Enable-ReFSDedup

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Enable-ReFSDedup [-Volume] <string> [-Type] <DedupVolumeType> [<CommonParameters>]
```

示例：

```powershell
Enable-ReFSDedup -Volume "D:" -Type DedupAndCompress
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/enable-refsdedup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/enable-refsdedup?view=powershell-7.5)

### Enable-ScheduledJob

版本：都有

模块：PSScheduledJob

语法：

```powershell
Enable-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-ScheduledJob [-Id] <int> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-ScheduledJob [-Name] <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启用计划作业

```powershell
Enable-ScheduledJob -Id 2 -PassThru
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/enable-scheduledjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/enable-scheduledjob?view=powershell-7.5)

### Enable-TlsCipherSuite

版本：都有

模块：TLS

语法（5.1）：

```powershell
Enable-TlsCipherSuite [-Name] <string> [[-Position] <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Enable-TlsCipherSuite [-Name] <string> [[-Position] <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启用密码套件

```powershell
Enable-TlsCipherSuite -Name TLS_DHE_DSS_WITH_AES_256_CBC_SHA
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/enable-tlsciphersuite?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/enable-tlsciphersuite?view=powershell-7.5)

### Enable-TlsEccCurve

版本：都有

模块：TLS

语法（5.1）：

```powershell
Enable-TlsEccCurve [-Name] <string> [[-Position] <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Enable-TlsEccCurve [-Name] <string> [[-Position] <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Enable-TlsEccCurve 'NistP384' -Position 0
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/enable-tlsecccurve?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/enable-tlsecccurve?view=powershell-7.5)

### Enable-TlsSessionTicketKey

版本：都有

模块：TLS

语法：

```powershell
Enable-TlsSessionTicketKey [-Password] <securestring> [-Path] <string> [-ServiceAccountName] <NTAccount> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：为 NetworkService 帐户配置 TLS 会话票证密钥

```powershell
$Password = Read-Host -AsSecureString
Enable-TlsSessionTicketKey -Password $Password -Path 'C:\KeyConfig\TlsSessionTicketKey.config' -ServiceAccountName NetworkService
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/enable-tlssessionticketkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/enable-tlssessionticketkey?view=powershell-7.5)

### Enable-TpmAutoProvisioning

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Enable-TpmAutoProvisioning [<CommonParameters>]
```

示例：启用自动预配

```powershell
PS C:\> Enable-TpmAutoProvisioning
TpmReady : False
TpmPresent : True
ManagedAuthLevel : Full
OwnerAuth : OwnerClearDisabled : True
AutoProvisioning : Enabled
LockedOut : False
SelfTest : {191, 191, 245, 191...}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/enable-tpmautoprovisioning?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/enable-tpmautoprovisioning?view=powershell-7.5)

### Enable-Uev

版本：仅5.1

模块：UEV

语法：

```powershell
Enable-Uev [<CommonParameters>]
```

示例：启用 UE-V 服务

```powershell
PS C:\>Enable-Uev
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/enable-uev?view=powershell-5.1)

### Enable-UevAppxPackage

版本：仅5.1

模块：UEV

语法：

```powershell
Enable-UevAppxPackage [-PackageFamilyName] <string[]> [-CurrentComputerUser] [-WhatIf] [-Confirm] [<CommonParameters>]
Enable-UevAppxPackage [-PackageFamilyName] <string[]> -Computer [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启用 Windows 8 应用的同步

```powershell
PS C:\>Enable-UevAppxPackage -PackageFamilyName "Microsoft.BingTravel"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/enable-uevappxpackage?view=powershell-5.1)

### Enable-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Enable-UevTemplate [-ID] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启用指定模板

```powershell
PS C:\> Enable-UevTemplate -ID "MicrosoftCalculator6"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/enable-uevtemplate?view=powershell-5.1)

### Enable-WindowsErrorReporting

版本：都有

模块：WindowsErrorReporting

语法：

```powershell
Enable-WindowsErrorReporting [<CommonParameters>]
```

示例：启用 Windows 错误报告

```powershell
PS C:\> Enable-WindowsErrorReporting
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowserrorreporting/enable-windowserrorreporting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowserrorreporting/enable-windowserrorreporting?view=powershell-7.5)

### Enable-WindowsOptionalFeature

版本：都有

模块：Dism

语法：

```powershell
Enable-WindowsOptionalFeature -FeatureName <string[]> -Online [-PackageName <string>] [-All] [-LimitAccess] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Enable-WindowsOptionalFeature -FeatureName <string[]> -Path <string> [-PackageName <string>] [-All] [-LimitAccess] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：在正在运行的操作系统中启用可选功能

```powershell
PS C:\> Enable-WindowsOptionalFeature -Online -FeatureName "Hearts" -All
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/enable-windowsoptionalfeature?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/enable-windowsoptionalfeature?view=powershell-7.5)

### Enable-WSManCredSSP

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Enable-WSManCredSSP [-Role] <string> [[-DelegateComputer] <string[]>] [-Force] [<CommonParameters>]
```

示例：委托客户端凭据

```powershell
Enable-WSManCredSSP -Role "Client" -DelegateComputer "Server02.fabrikam.com"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/enable-wsmancredssp?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/enable-wsmancredssp?view=powershell-7.5)

### Expand-OsImage

版本：都有

模块：Dism

语法：

```powershell
Expand-OsImage -ImagePath <string> -ApplyPath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Expand-WindowsCustomDataImage

版本：都有

模块：Dism

语法：

```powershell
Expand-WindowsCustomDataImage -ImagePath <string> -CustomDataImage <string> -SingleInstance [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：展开自定义数据映像

```powershell
PS C:\> Expand-WindowsCustomDataImage -CustomDataImage "C:\oem.ppkg" -ImagePath "C:\" -SingleInstance
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/expand-windowscustomdataimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/expand-windowscustomdataimage?view=powershell-7.5)

### Expand-WindowsImage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Expand-WindowsImage -ImagePath <string> -Name <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -Index <uint32> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Expand-WindowsImage -ImagePath <string> -Name <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -Index <uint> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Expand-WindowsImage -ImagePath <string> -ApplyPath <string> [-SplitImageFilePattern <string>] [-CheckIntegrity] [-ConfirmTrustedFile] [-NoRpFix] [-Verify] [-WIMBoot] [-Compact] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：将文件中的映像应用到分区

```powershell
PS C:\> Expand-WindowsImage -ImagePath "c:\imagestore\custom.wim" -ApplyPath "d:\" -Index 1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/expand-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/expand-windowsimage?view=powershell-7.5)

### Export-BcdStore

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Export-BcdStore [-Path] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/export-bcdstore?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/export-bcdstore?view=powershell-7.5)

### Export-BinaryMiLog

版本：都有

模块：CimCmdlets

语法：

```powershell
Export-BinaryMiLog [-Path] <string> [-InputObject <ciminstance>] [<CommonParameters>]
```

示例：创建 CimInstances 的二进制表示形式

```powershell
Get-CimInstance Win32_Process | Export-BinaryMiLog -Path "Processes.bmil"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/export-binarymilog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/export-binarymilog?view=powershell-7.5)

### Export-Certificate

版本：都有

模块：PKI

语法：

```powershell
Export-Certificate -FilePath <string> -Cert <Certificate> [-Type <CertType>] [-NoClobber] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$cert = Get-ChildItem -Path Cert:\CurrentUser\My\EEDEF61D4FF6EDBAAD538BB08CCAADDC3EE28FF

Export-Certificate -Cert $cert -FilePath C:\Certs\user.sst -Type SST
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/export-certificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/export-certificate?view=powershell-7.5)

### Export-Console

版本：仅5.1

模块：Microsoft.PowerShell.Core

语法：

```powershell
Export-Console [[-Path] <string>] [-Force] [-NoClobber] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：导出当前会话中管理单元的名称

```powershell
PS C:\> Export-Console -Path $PSHOME\Consoles\ConsoleS1.psc1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/export-console?view=powershell-5.1)

### Export-Counter

版本：都有

模块：Microsoft.PowerShell.Diagnostics

语法：

```powershell
Export-Counter [-Path] <string> -InputObject <PerformanceCounterSampleSet[]> [-FileFormat <string>] [-MaxSize <uint32>] [-Force] [-Circular] [<CommonParameters>]
```

示例：将计数器数据导出到文件

```powershell
Get-Counter "\Processor(*)\% Processor Time" | Export-Counter -Path $HOME\Counters.blg
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/export-counter?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/export-counter?view=powershell-7.5)

### Export-OsImage

版本：都有

模块：Dism

语法：

```powershell
Export-OsImage -SrcImagePath <string> -DestImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Export-PfxCertificate

版本：都有

模块：PKI

语法：

```powershell
Export-PfxCertificate [-PFXData] <PfxData> [-FilePath] <string> [-NoProperties] [-NoClobber] [-Force] [-CryptoAlgorithmOption <CryptoAlgorithmOptions>] [-ChainOption <ExportChainOption>] [-ProtectTo <string[]>] [-Password <securestring>] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-PfxCertificate [-Cert] <Certificate> [-FilePath] <string> [-NoProperties] [-NoClobber] [-Force] [-CryptoAlgorithmOption <CryptoAlgorithmOptions>] [-ChainOption <ExportChainOption>] [-ProtectTo <string[]>] [-Password <securestring>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$mypwd = ConvertTo-SecureString -String '1234' -Force -AsPlainText

Get-ChildItem -Path Cert:\LocalMachine\My\5F98EBBFE735CDDAE00E33E0FD69050EF9220254 |
 Export-PfxCertificate -FilePath C:\mypfx.pfx -Password $mypwd
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/export-pfxcertificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/export-pfxcertificate?view=powershell-7.5)

### Export-ProvisioningPackage

版本：都有

模块：Provisioning

语法：

```powershell
Export-ProvisioningPackage [-OutputFolder] <string> -PackageId <string> [-AllowClobber] [-AnswerFileOnly] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Export-ProvisioningPackage [-PackagePath] <string> [-OutputFolder] <string> [-AllowClobber] [-AnswerFileOnly] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Export-ProvisioningPackage [-RuntimeMetadata] <RuntimeProvPackageMetadata> [-OutputFolder] <string> [-AllowClobber] [-AnswerFileOnly] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Export-ProvisioningPackage -PackageId {e2ea11f5-d8b0-4db9-bf96-8c909dc2fed5} -OutputFolder D:\Package
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/export-provisioningpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/export-provisioningpackage?view=powershell-7.5)

### Export-StartLayout

版本：都有

模块：StartLayout

语法：

```powershell
Export-StartLayout [-Path] <string> [-UseDesktopApplicationID] [-WhatIf] [-Confirm] [<CommonParameters>]
Export-StartLayout -LiteralPath <string> [-UseDesktopApplicationID] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：导出布局

```powershell
PS C:\> Export-StartLayout -Path "C:\Layouts\Marketing.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/startlayout/export-startlayout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/startlayout/export-startlayout?view=powershell-7.5)

### Export-StartLayoutEdgeAssets

版本：都有

模块：StartLayout

语法：

```powershell
Export-StartLayoutEdgeAssets [-Path] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Export-StartLayoutEdgeAssets -LiteralPath <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：导出资源

```powershell
Export-StartLayoutEdgeAssets -Path "C:\Layouts\assets.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/startlayout/export-startlayoutedgeassets?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/startlayout/export-startlayoutedgeassets?view=powershell-7.5)

### Export-TlsSessionTicketKey

版本：都有

模块：TLS

语法：

```powershell
Export-TlsSessionTicketKey [-Password] <securestring> [[-Path] <string>] [-ServiceAccountName] <NTAccount> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：导出 TLS 会话票证密钥

```powershell
$Password = Read-Host -AsSecureString
Export-TlsSessionTicketKey -Password $Password -Path 'C:\KeyConfig\TlsSessionTicketKey.config' -ServiceAccountName NetworkService
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/export-tlssessionticketkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/export-tlssessionticketkey?view=powershell-7.5)

### Export-Trace

版本：都有

模块：Provisioning

语法：

```powershell
Export-Trace [-ETLFile] <string> [-Overwrite] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：从 ETL 文件导出跟踪事件

```powershell
PS C:\> Export-Trace -ETLFile C:\Windows\Logs\WindowsUpdate\WindowsUpdate.20211013.074054.819.1.etl -LogsDirectoryPath C:\ETL\Logs
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/export-trace?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/export-trace?view=powershell-7.5)

### Export-UevConfiguration

版本：仅5.1

模块：UEV

语法：

```powershell
Export-UevConfiguration [-Path] <string> [<CommonParameters>]
```

示例：导出 UE-V 配置

```powershell
PS C:\> Export-UevConfiguration -Path "ContosoUev.uev"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/export-uevconfiguration?view=powershell-5.1)

### Export-UevPackage

版本：仅5.1

模块：UEV

语法：

```powershell
Export-UevPackage [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Export-UevPackage -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：导出 UE-V 包

```powershell
PS C:\> Export-UevPackage -Path "MicrosoftCalculator6.pkgx"
<SettingsDocument>
<registry>
<Setting Type="VT_BINARY" Name="registry://HKCU\Software\Microsoft\Calc\Window_Placement" Action="Update">LAAAAAAAAAABAAAA/////////////////////60AAABQAAAAVAIAANQBAAA=</Setting>
<Setting Type="VT_DWORD" Name="registry://HKCU\Software\Microsoft\Calc\layout" Action="Update">2</Setting>
</registry>
</SettingsDocument>
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/export-uevpackage?view=powershell-5.1)

### Export-WindowsCapabilitySource

版本：都有

模块：Dism

语法：

```powershell
Export-WindowsCapabilitySource -Name <string> -Source <string> -Target <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：导出功能包的存储库

```powershell
Export-WindowsCapabilitySource -Path c:\mount\windows -Source D:\ -Target C:\repository -Name App.StepsRecorder~~~~0.0.1.0
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/export-windowscapabilitysource?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/export-windowscapabilitysource?view=powershell-7.5)

### Export-WindowsDriver

版本：都有

模块：Dism

语法：

```powershell
Export-WindowsDriver -Path <string> [-Destination <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsDriver -Online [-Destination <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：从正在运行的操作系统导出驱动程序

```powershell
PS C:\> Export-WindowsDriver -Online -Destination d:\drivers
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/export-windowsdriver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/export-windowsdriver?view=powershell-7.5)

### Export-WindowsImage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceName <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceIndex <uint32> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceName <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> -SourceIndex <uint> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Export-WindowsImage -DestinationImagePath <string> -SourceImagePath <string> [-CheckIntegrity] [-CompressionType <string>] [-DestinationName <string>] [-Setbootable] [-SplitImageFilePattern <string>] [-WIMBoot] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：导出映像

```powershell
PS C:\> Export-WindowsImage -SourceImagePath C:\imagestore\custom.wim -SourceIndex 1 -DestinationImagePath c:\imagestore\export.wim -DestinationName "Exported Image"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/export-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/export-windowsimage?view=powershell-7.5)

### Find-LapsADExtendedRights

版本：都有

模块：LAPS

语法：

```powershell
Find-LapsADExtendedRights -Identity <string[]> [-Credential <pscredential>] [-Domain <string>] [-DomainController <string>] [-IncludeComputers] [<CommonParameters>]
```

示例：

```powershell
Find-LapsADExtendedRights -Identity LapsTestOU
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/find-lapsadextendedrights?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/find-lapsadextendedrights?view=powershell-7.5)

### Format-SecureBootUEFI

版本：都有

模块：SecureBoot

语法：

```powershell
Format-SecureBootUEFI -Name <string> -SignatureOwner <guid> -CertificateFilePath <string[]> [-FormatWithCert] [-SignableFilePath <string>] [-Time <string>] [-AppendWrite] [-ContentFilePath <string>] [<CommonParameters>]
Format-SecureBootUEFI -Name <string> -SignatureOwner <guid> -Hash <string[]> -Algorithm <string> [-SignableFilePath <string>] [-Time <string>] [-AppendWrite] [-ContentFilePath <string>] [<CommonParameters>]
Format-SecureBootUEFI -Name <string> -Delete [-SignableFilePath <string>] [-Time <string>] [<CommonParameters>]
```

示例：格式化私钥

```powershell
PS C:\> Format-SecureBootUefi -Name PK -SignatureOwner 12345678-1234-1234-1234-123456789abc -CertificateFilePath PK.cer -SignableFilePath GeneratedFileToSign.bin -Time 2011-11-01T13:30:00Z | Format-List
Name : PK
Time : 2011-11-01T13:30:00Z
AppendWrite : False
Content : {232, 102, 87, 60...}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/format-securebootuefi?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/format-securebootuefi?view=powershell-7.5)

### Get-Acl

版本：都有

模块：Microsoft.PowerShell.Security

语法（5.1）：

```powershell
Get-Acl [[-Path] <string[]>] [-Audit] [-AllCentralAccessPolicies] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-UseTransaction] [<CommonParameters>]
Get-Acl -InputObject <psobject> [-Audit] [-AllCentralAccessPolicies] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-UseTransaction] [<CommonParameters>]
Get-Acl [-LiteralPath <string[]>] [-Audit] [-AllCentralAccessPolicies] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-UseTransaction] [<CommonParameters>]
```

语法（7）：

```powershell
Get-Acl [[-Path] <string[]>] [-Audit] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Acl -InputObject <psobject> [-Audit] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Acl [-LiteralPath <string[]>] [-Audit] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
```

示例：获取文件夹的 ACL

```powershell
Get-Acl C:\Windows
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/get-acl?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/get-acl?view=powershell-7.5)

### Get-AppLockerFileInformation

版本：都有

模块：AppLocker

语法：

```powershell
Get-AppLockerFileInformation [[-Path] <List[string]>] [<CommonParameters>]
Get-AppLockerFileInformation [[-Packages] <List[AppxPackage]>] [<CommonParameters>]
Get-AppLockerFileInformation -Directory <string> [-FileType <List[AppLockerFileType]>] [-Recurse] [<CommonParameters>]
Get-AppLockerFileInformation -EventLog [-LogPath <string>] [-EventType <List[AppLockerEventType]>] [-Statistics] [<CommonParameters>]
```

示例：获取 .exe 文件和脚本的文件信息

```powershell
PS C:\> Get-AppLockerFileInformation -Directory C:\Windows\system32\ -Recurse -FileType exe, script
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/get-applockerfileinformation?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/get-applockerfileinformation?view=powershell-7.5)

### Get-AppLockerPolicy

版本：都有

模块：AppLocker

语法：

```powershell
Get-AppLockerPolicy -Local [-Xml] [<CommonParameters>]
Get-AppLockerPolicy -Domain -Ldap <string> [-Xml] [<CommonParameters>]
Get-AppLockerPolicy -Effective [-Xml] [<CommonParameters>]
```

示例：获取 AppLocker 策略

```powershell
PS C:\> Get-AppLockerPolicy -Local
 Version RuleCollections RuleCollectionTypes
 ------- --------------- -------------------
 1 {} {}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/get-applockerpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/get-applockerpolicy?view=powershell-7.5)

### Get-AppProvisionedSharedPackageContainer

版本：都有

模块：Dism

语法：

```powershell
Get-AppProvisionedSharedPackageContainer -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-AppProvisionedSharedPackageContainer -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Get-AppSharedPackageContainer

版本：都有

模块：Appx

语法：

```powershell
Get-AppSharedPackageContainer [[-Name] <string>] [[-Id] <string>] [<CommonParameters>]
```

示例：

```powershell
Get-AppSharedPackageContainer -Name Contoso*
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appsharedpackagecontainer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appsharedpackagecontainer?view=powershell-7.5)

### Get-AppvClientApplication

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvClientApplication [[-Name] <string>] [[-Version] <string>] [-All] [<CommonParameters>]
```

示例：获取当前用户的应用版本

```powershell
PS C:\> Get-AppvClientApplication -Name "AppName" -Version 1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvclientapplication?view=powershell-5.1)

### Get-AppvClientConfiguration

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvClientConfiguration [[-Name] <string>] [<CommonParameters>]
```

示例：显示全部配置设置

```powershell
PS C:\> Get-AppvClientConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvclientconfiguration?view=powershell-5.1)

### Get-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvClientConnectionGroup [[-Name] <string>] [-All] [<CommonParameters>]
Get-AppvClientConnectionGroup [-GroupId] <guid> [[-VersionId] <guid>] [-All] [<CommonParameters>]
```

示例：按名称获取组的全部版本

```powershell
PS C:\> Get-AppvClientConnectionGroup -Name "MyConnectionGroup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvclientconnectiongroup?view=powershell-5.1)

### Get-AppvClientMode

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvClientMode [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvclientmode?view=powershell-5.1)

### Get-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvClientPackage [[-Name] <string>] [[-Version] <string>] [-All] [<CommonParameters>]
Get-AppvClientPackage [-PackageId] <guid> [[-VersionId] <guid>] [-All] [<CommonParameters>]
```

示例：获取名称匹配字符串的包

```powershell
PS C:\> Get-AppvClientPackage -Name "MyApp*" -All
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvclientpackage?view=powershell-5.1)

### Get-AppvPublishingServer

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvPublishingServer [[-ServerId] <uint32>] [<CommonParameters>]
Get-AppvPublishingServer [[-Name] <string>] [[-URL] <string>] [<CommonParameters>]
```

示例：按友好名称获取服务器

```powershell
PS C:\> Get-AppvPublishingServer -Name "Server*"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvpublishingserver?view=powershell-5.1)

### Get-AppvStatus

版本：仅5.1

模块：AppvClient

语法：

```powershell
Get-AppvStatus [<CommonParameters>]
```

示例：获取状态

```powershell
PS C:\> Get-AppvStatus
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/get-appvstatus?view=powershell-5.1)

### Get-AppxDefaultVolume

版本：都有

模块：Appx

语法：

```powershell
Get-AppxDefaultVolume [<CommonParameters>]
```

示例：获取默认卷

```powershell
Get-AppxDefaultVolume
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxdefaultvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxdefaultvolume?view=powershell-7.5)

### Get-AppxPackage

版本：都有

模块：Appx

语法：

```powershell
Get-AppxPackage [[-Name] <string>] [[-Publisher] <string>] [-AllUsers] [-PackageTypeFilter <PackageTypes>] [-User <string>] [-Volume <AppxVolume>] [<CommonParameters>]
```

示例：获取每个用户帐户的全部应用包

```powershell
Get-AppxPackage -AllUsers
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxpackage?view=powershell-7.5)

### Get-AppxPackageAutoUpdateSettings

版本：都有

模块：Appx

语法：

```powershell
Get-AppxPackageAutoUpdateSettings [[-PackageFamilyName] <string>] [-ShowUpdateAvailability] [-AllUsers] [<CommonParameters>]
```

示例：获取全部应用包自动更新设置

```powershell
Get-AppxPackageAutoUpdateSettings
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxpackageautoupdatesettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxpackageautoupdatesettings?view=powershell-7.5)

### Get-AppxPackageManifest

版本：都有

模块：Appx

语法：

```powershell
Get-AppxPackageManifest [-Package] <string> [[-User] <string>] [<CommonParameters>]
```

示例：获取应用包的清单

```powershell
Get-AppxPackageManifest -Package "package1_1.0.0.0_neutral__8wekyb3d8bbwe"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxpackagemanifest?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxpackagemanifest?view=powershell-7.5)

### Get-AppxProvisionedPackage

版本：都有

模块：Dism

语法：

```powershell
Get-AppxProvisionedPackage -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-AppxProvisionedPackage -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：列出已装载映像中为每个帐户安装的应用包

```powershell
PS C:\> Get-AppxProvisionedPackage -Path "c:\offline"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-appxprovisionedpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-appxprovisionedpackage?view=powershell-7.5)

### Get-AppxVolume

版本：都有

模块：Appx

语法：

```powershell
Get-AppxVolume [[-Path] <string>] [<CommonParameters>]
Get-AppxVolume -Online [-Path <string>] [<CommonParameters>]
Get-AppxVolume -Offline [-Path <string>] [<CommonParameters>]
```

示例：获取全部卷

```powershell
Get-AppxVolume
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/get-appxvolume?view=powershell-7.5)

### Get-AuthenticodeSignature

版本：都有

模块：Microsoft.PowerShell.Security

语法：

```powershell
Get-AuthenticodeSignature [-FilePath] <string[]> [<CommonParameters>]
Get-AuthenticodeSignature -LiteralPath <string[]> [<CommonParameters>]
Get-AuthenticodeSignature -SourcePathOrExtension <string[]> -Content <byte[]> [<CommonParameters>]
```

示例：获取文件的验证码签名

```powershell
Get-AuthenticodeSignature -FilePath "C:\Test\NewScript.ps1"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/get-authenticodesignature?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/get-authenticodesignature?view=powershell-7.5)

### Get-BcdEntry

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Get-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdentry?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdentry?view=powershell-7.5)

### Get-BcdEntryDebugSettings

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Get-BcdEntryDebugSettings [[-Store] <BcdStoreInfo>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdentrydebugsettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdentrydebugsettings?view=powershell-7.5)

### Get-BcdEntryHypervisorSettings

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Get-BcdEntryHypervisorSettings [[-Store] <BcdStoreInfo>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdentryhypervisorsettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdentryhypervisorsettings?view=powershell-7.5)

### Get-BcdStore

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Get-BcdStore [[-Path] <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdstore?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/get-bcdstore?view=powershell-7.5)

### Get-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Get-BitsTransfer [[-Name] <string[]>] [-AllUsers] [<CommonParameters>]
Get-BitsTransfer [-JobId] <guid[]> [<CommonParameters>]
```

示例：获取当前用户拥有的全部 BitsJob 对象

```powershell
PS C:\> Get-BitsTransfer

JobId DisplayName TransferType JobState OwnerAccount
----- ----------- ------------ -------- ------------
07acbe90-7d25-4d05-a... TestJob2 Download Suspended DOMAIN01\user01
c0dd3d8c-c3a2-4562-8... TestJob1 Download Transferred DOMAIN01\user01
1ef8c549-7a92-4173-b... BitsJobTransfer Download Transferred DOMAIN01\user01
2c8302d5-3f44-4981-8... BitsJobTransfer Download Transferred DOMAIN01\user01
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/get-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/get-bitstransfer?view=powershell-7.5)

### Get-Certificate

版本：都有

模块：PKI

语法：

```powershell
Get-Certificate -Template <string> [-Url <uri>] [-SubjectName <string>] [-DnsName <string[]>] [-Credential <PkiCredential>] [-CertStoreLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Get-Certificate -Request <Certificate> [-Credential <PkiCredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

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

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificate?view=powershell-7.5)

### Get-CertificateAutoEnrollmentPolicy

版本：都有

模块：PKI

语法：

```powershell
Get-CertificateAutoEnrollmentPolicy -Scope <AutoEnrollmentPolicyScope> -context <Context> [<CommonParameters>]
```

示例：

```powershell
Get-CertificateAutoEnrollmentPolicy -Scope Local -Context User
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificateautoenrollmentpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificateautoenrollmentpolicy?view=powershell-7.5)

### Get-CertificateEnrollmentPolicyServer

版本：都有

模块：PKI

语法：

```powershell
Get-CertificateEnrollmentPolicyServer -Scope <EnrollmentPolicyServerScope> -context <Context> [-Url <uri>] [<CommonParameters>]
```

示例：

```powershell
Get-CertificateEnrollmentPolicyServer -Scope All -Context User
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificateenrollmentpolicyserver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificateenrollmentpolicyserver?view=powershell-7.5)

### Get-CertificateNotificationTask

版本：都有

模块：PKI

语法：

```powershell
Get-CertificateNotificationTask [<CommonParameters>]
```

示例：

```powershell
Get-CertificateNotificationTask
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificatenotificationtask?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-certificatenotificationtask?view=powershell-7.5)

### Get-CimAssociatedInstance

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ResourceUri <uri>] [-ComputerName <string[]>] [-KeyOnly] [<CommonParameters>]
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] -CimSession <CimSession[]> [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ResourceUri <uri>] [-KeyOnly] [<CommonParameters>]
```

语法（7）：

```powershell
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ResourceUri <uri>] [-ComputerName <string[]>] [-KeyOnly] [<CommonParameters>]
Get-CimAssociatedInstance [-InputObject] <ciminstance> [[-Association] <string>] -CimSession <CimSession[]> [-ResultClassName <string>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ResourceUri <uri>] [-KeyOnly] [<CommonParameters>]
```

示例：获取特定实例的所有关联实例

```powershell
$disk = Get-CimInstance -ClassName Win32_LogicalDisk -KeyOnly
Get-CimAssociatedInstance -InputObject $disk[1]
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-cimassociatedinstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-cimassociatedinstance?view=powershell-7.5)

### Get-CimClass

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
```

语法（7）：

```powershell
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] [-Amended] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
Get-CimClass [[-ClassName] <string>] [[-Namespace] <string>] -CimSession <CimSession[]> [-Amended] [-OperationTimeoutSec <uint>] [-MethodName <string>] [-PropertyName <string>] [-QualifierName <string>] [<CommonParameters>]
```

示例：获取所有类定义

```powershell
Get-CimClass
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-cimclass?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-cimclass?view=powershell-7.5)

### Get-CimInstance

版本：都有

模块：CimCmdlets

语法（5.1）：

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

语法（7）：

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

示例：获取指定类的 CIM 实例

```powershell
Get-CimInstance -ClassName Win32_Process
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-ciminstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-ciminstance?view=powershell-7.5)

### Get-CimSession

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Get-CimSession [[-ComputerName] <string[]>] [<CommonParameters>]
Get-CimSession [-Id] <uint32[]> [<CommonParameters>]
Get-CimSession -InstanceId <guid[]> [<CommonParameters>]
Get-CimSession -Name <string[]> [<CommonParameters>]
```

语法（7）：

```powershell
Get-CimSession [[-ComputerName] <string[]>] [<CommonParameters>]
Get-CimSession [-Id] <uint[]> [<CommonParameters>]
Get-CimSession -InstanceId <guid[]> [<CommonParameters>]
Get-CimSession -Name <string[]> [<CommonParameters>]
```

示例：从当前 PowerShell 会话获取 CIM 会话

```powershell
New-CimSession -ComputerName Server01, Server02
Get-CimSession
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-cimsession?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/get-cimsession?view=powershell-7.5)

### Get-CIPolicy

版本：都有

模块：ConfigCI

语法：

```powershell
Get-CIPolicy [-FilePath] <string> [<CommonParameters>]
```

示例：从策略获取规则

```powershell
PS C:\> Get-CIPolicy -FilePath '.\Policy.xml'
Name : MSIT Test CodeSign CA 3
Id : ID_SIGNER_S_17
TypeId : Allow
Root : FA6B9A2230CE08BCA81D096B28CF495672401D3A43A0D285CF352464A6C9C7FD
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : VeriSign Class 3 Code Signing 2010 CA
Id : ID_SIGNER_S_1D
TypeId : Allow
Root : 4843A82ED3B1F2BFBEE9671960E1940C942F688D
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : Microsoft Windows Third Party Component CA 2012
Id : ID_SIGNER_S_1E
TypeId : Allow
Root : CEC1AFD0E310C55C1DCC601AB8E172917706AA32FB5EAF826813547FDF02DD46
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-cipolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-cipolicy?view=powershell-7.5)

### Get-CIPolicyIdInfo

版本：都有

模块：ConfigCI

语法：

```powershell
Get-CIPolicyIdInfo [-FilePath] <string> [<CommonParameters>]
```

示例：显示代码完整性策略信息

```powershell
PS C:\> Get-CIPolicyIdInfo -FilePath ".\Policy03.xml"

Provider : ConfigCIPolicy
Key : PolicyInfo
ValueName : Name
ValueType : String
Value : CIPolicy03

Provider : ConfigCIPolicy
Key : PolicyInfo
ValueName : PolicyId
ValueType : String
Value : CIP077
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-cipolicyidinfo?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-cipolicyidinfo?view=powershell-7.5)

### Get-CIPolicyInfo

版本：都有

模块：ConfigCI

语法：

```powershell
Get-CIPolicyInfo [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-cipolicyinfo?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-cipolicyinfo?view=powershell-7.5)

### Get-ComputerInfo

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-ComputerInfo [[-Property] <string[]>] [<CommonParameters>]
```

示例：获取所有计算机属性

```powershell
Get-ComputerInfo
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-computerinfo?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-computerinfo?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 与原版差异：字段精简。

- 类型：Go实现。
- 功能：汇总系统信息。对应 bash `uname -a` + 发行版信息。
- 参数：无。
- 实现：读 /etc/os-release（NAME、VERSION_ID）与 /proc/meminfo（MemTotal）、runtime 的 GOOS/GOARCH、`os.Hostname()`、`runtime.NumCPU()`。
- 输出：ComputerInfo 对象，字段 CsName、OsName、OsVersion、OsArchitecture、OsPlatform、CsTotalPhysicalMemory、CsProcessors。

### Get-ComputerRestorePoint

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-ComputerRestorePoint [[-RestorePoint] <int[]>] [<CommonParameters>]
Get-ComputerRestorePoint -LastStatus [<CommonParameters>]
```

示例：获取所有系统还原点

```powershell
Get-ComputerRestorePoint
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-computerrestorepoint?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-computerrestorepoint?view=powershell-7.5)

### Get-ControlPanelItem

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-ControlPanelItem [[-Name] <string[]>] [-Category <string[]>] [<CommonParameters>]
Get-ControlPanelItem -CanonicalName <string[]> [-Category <string[]>] [<CommonParameters>]
```

示例：获取所有控制面板项

```powershell
Get-ControlPanelItem
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-controlpanelitem?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-controlpanelitem?view=powershell-7.5)

### Get-Counter

版本：都有

模块：Microsoft.PowerShell.Diagnostics

语法：

```powershell
Get-Counter [[-Counter] <string[]>] [-SampleInterval <int>] [-MaxSamples <long>] [-Continuous] [-ComputerName <string[]>] [<CommonParameters>]
Get-Counter [-ListSet] <string[]> [-ComputerName <string[]>] [<CommonParameters>]
```

示例：获取计数器集列表

```powershell
Get-Counter -ListSet *
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/get-counter?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/get-counter?view=powershell-7.5)

### Get-DAPolicyChange

版本：都有

模块：NetSecurity

语法：

```powershell
Get-DAPolicyChange [[-Servers] <string[]>] [[-Domains] <string[]>] [-DisplayName] <string> [[-PolicyStore] <string>] [-PSLocation] <string> [-EndpointType] <string> [[-DnsServers] <string[]>] [<CommonParameters>]
```

示例：

```powershell
PS C:\>Get-DAPolicyChange -DisplayName "TunnelPolicy1" -EndpointType Endpoint1 -PSLocation "C:\Update.ps1" -Servers "server1.corp.contoso.com", "server2.corp.contoso.com", "server3.corp.contoso.com"
IPsec Rule name : TunnelPolicy1
Action : Add
IPv6addresses : 2001:4829:3243::100:1
 : 2001:4829:3243::100:1
GPO : contoso\DAClientPolicy

IPsec Rule name : TunnelPolicy1
Action : Delete
IPv6addresses : 2001:4829:3243::100:3
 : 2001:4829:3243::100:4
GPO : contoso\DAClientPolicy

FQDN's that did not resolve into IP address:
server1.corp.contoso.com
server3.corp.contoso.com
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/get-dapolicychange?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/get-dapolicychange?view=powershell-7.5)

### Get-DeliveryOptimizationLog

版本：都有

模块：DeliveryOptimization

语法（5.1）：

```powershell
Get-DeliveryOptimizationLog [[-Path] <string[]>] [-LevelFilter <uint32>] [-Provider <ProviderType>] [-Flush] [<CommonParameters>]
```

语法（7）：

```powershell
Get-DeliveryOptimizationLog [[-Path] <string[]>] [-LevelFilter <uint>] [-Provider <ProviderType>] [-Flush] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/deliveryoptimization/get-deliveryoptimizationlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/deliveryoptimization/get-deliveryoptimizationlog?view=powershell-7.5)

### Get-DeliveryOptimizationLogAnalysis

版本：都有

模块：DeliveryOptimization

语法：

```powershell
Get-DeliveryOptimizationLogAnalysis [[-Path] <string[]>] [-ListConnections] [-Flush] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/deliveryoptimization/get-deliveryoptimizationloganalysis?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/deliveryoptimization/get-deliveryoptimizationloganalysis?view=powershell-7.5)

### Get-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-EventLog [-LogName] <string> [[-InstanceId] <long[]>] [-ComputerName <string[]>] [-Newest <int>] [-After <datetime>] [-Before <datetime>] [-UserName <string[]>] [-Index <int[]>] [-EntryType <string[]>] [-Source <string[]>] [-Message <string>] [-AsBaseObject] [<CommonParameters>]
Get-EventLog [-ComputerName <string[]>] [-List] [-AsString] [<CommonParameters>]
```

示例：获取本地计算机上的事件日志

```powershell
Get-EventLog -List
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-eventlog?view=powershell-7.5)

### Get-HotFix

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-HotFix [[-Id] <string[]>] [-ComputerName <string[]>] [-Credential <pscredential>] [<CommonParameters>]
Get-HotFix [-Description <string[]>] [-ComputerName <string[]>] [-Credential <pscredential>] [<CommonParameters>]
```

示例：获取本地计算机上的所有修补程序

```powershell
Get-HotFix
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-hotfix?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-hotfix?view=powershell-7.5)

### Get-InstalledLanguage

版本：都有

模块：LanguagePackManagement

语法：

```powershell
Get-InstalledLanguage [[-Language] <string>] [<CommonParameters>]
```

示例：查看设备上安装的语言

```powershell
Get-InstalledLanguage
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/get-installedlanguage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/get-installedlanguage?view=powershell-7.5)

### Get-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
Get-JobTrigger [-InputObject] <ScheduledJobDefinition> [[-TriggerId] <int[]>] [<CommonParameters>]
Get-JobTrigger [-Id] <int> [[-TriggerId] <int[]>] [<CommonParameters>]
Get-JobTrigger [-Name] <string> [[-TriggerId] <int[]>] [<CommonParameters>]
```

示例：按计划作业名称获取作业触发器

```powershell
Get-JobTrigger -Name "BackupJob"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/get-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/get-jobtrigger?view=powershell-7.5)

### Get-KdsConfiguration

版本：都有

模块：Kds

语法：

```powershell
Get-KdsConfiguration [<CommonParameters>]
```

示例：检索当前 KDS 配置

```powershell
PS C:\> Get-KdsConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/kds/get-kdsconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/kds/get-kdsconfiguration?view=powershell-7.5)

### Get-KdsRootKey

版本：都有

模块：Kds

语法：

```powershell
Get-KdsRootKey [<CommonParameters>]
```

示例：检索根密钥值列表

```powershell
PS C:\> Get-KdsRootKey
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/kds/get-kdsrootkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/kds/get-kdsrootkey?view=powershell-7.5)

### Get-LapsADPassword

版本：都有

模块：LAPS

语法：

```powershell
Get-LapsADPassword [-Identity] <string[]> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -Domain <string> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -DomainController <string> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -Port <int> [-Credential <pscredential>] [-DecryptionCredential <pscredential>] [-IncludeHistory] [-AsPlainText] [-DomainController <string>] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -RecoveryMode [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
Get-LapsADPassword [-Identity] <string[]> -RecoveryMode -Port <int> [-IncludeHistory] [-AsPlainText] [<CommonParameters>]
```

示例：

```powershell
Get-LapsADPassword LAPSCLIENT
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/get-lapsadpassword?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/get-lapsadpassword?view=powershell-7.5)

### Get-LocalGroup

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Get-LocalGroup [[-Name] <string[]>] [<CommonParameters>]
Get-LocalGroup [[-SID] <SecurityIdentifier[]>] [<CommonParameters>]
```

示例：获取管理员组

```powershell
Get-LocalGroup -Name "Administrators"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/get-localgroup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/get-localgroup?view=powershell-7.5)

### Get-LocalGroupMember

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Get-LocalGroupMember [-Name] <string> [[-Member] <string>] [<CommonParameters>]
Get-LocalGroupMember [-Group] <LocalGroup> [[-Member] <string>] [<CommonParameters>]
Get-LocalGroupMember [-SID] <SecurityIdentifier> [[-Member] <string>] [<CommonParameters>]
```

示例：获取管理员组的所有成员

```powershell
Get-LocalGroupMember -Group "Administrators"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/get-localgroupmember?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/get-localgroupmember?view=powershell-7.5)

### Get-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Get-LocalUser [[-Name] <string[]>] [<CommonParameters>]
Get-LocalUser [[-SID] <SecurityIdentifier[]>] [<CommonParameters>]
```

示例：使用帐户名称获取帐户

```powershell
Get-LocalUser -Name "AdminContoso02"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/get-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/get-localuser?view=powershell-7.5)

### Get-NonRemovableAppsPolicy

版本：都有

模块：Dism

语法：

```powershell
Get-NonRemovableAppsPolicy -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-NonRemovableAppsPolicy -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：获取全部已安装不可移除应用包

```powershell
PS> Get-NonRemovableAppsPolicy -Online
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-nonremovableappspolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-nonremovableappspolicy?view=powershell-7.5)

### Get-OSConfiguration

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OSConfiguration [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfiguration?view=powershell-7.5)

### Get-OsConfigurationDocument

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OsConfigurationDocument [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
Get-OsConfigurationDocument [[-Id] <string>] [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationdocument?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationdocument?view=powershell-7.5)

### Get-OsConfigurationDocumentContent

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OsConfigurationDocumentContent [-Id] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationdocumentcontent?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationdocumentcontent?view=powershell-7.5)

### Get-OsConfigurationDocumentResult

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OsConfigurationDocumentResult [-Id] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [[-VerboseOption] <string>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationdocumentresult?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationdocumentresult?view=powershell-7.5)

### Get-OsConfigurationProperty

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OsConfigurationProperty [-Name] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [[-Id] <string>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationproperty?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/get-osconfigurationproperty?view=powershell-7.5)

### Get-OSConfigurationScenarioDefinition

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OsConfigurationScenarioDefinition [-Name] <string> [-Version] <string> [-SchemaVersion] <string> [<CommonParameters>]
```

示例：暂无

出处：[OsConfiguration 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration)（没有单独介绍）。

### Get-OSConfigurationScenarioDefinitionInfo

版本：都有

模块：OsConfiguration

语法：

```powershell
Get-OsConfigurationScenarioDefinitionInfo [[-Name] <string>] [[-Version] <string>] [[-SchemaVersion] <string>] [<CommonParameters>]
```

示例：暂无

出处：[OsConfiguration 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration)（没有单独介绍）。

### Get-PfxData

版本：都有

模块：PKI

语法：

```powershell
Get-PfxData [-FilePath] <string> [-Password <securestring>] [<CommonParameters>]
```

示例：

```powershell
$mypwd = ConvertTo-SecureString -String '1234' -Force -AsPlainText

$mypfx = Get-PfxData -FilePath C:\mypfx.pfx -Password $mypwd
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-pfxdata?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/get-pfxdata?view=powershell-7.5)

### Get-PmemDedicatedMemory

版本：都有

模块：PersistentMemory

语法：

```powershell
Get-PmemDedicatedMemory [<CommonParameters>]
Get-PmemDedicatedMemory [[-DeviceNumber] <uint32[]>] [<CommonParameters>]
```

示例：获取全部专用持久内存

```powershell
Get-PmemDedicatedMemory
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemdedicatedmemory?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemdedicatedmemory?view=powershell-7.5)

### Get-PmemDisk

版本：都有

模块：PersistentMemory

语法：

```powershell
Get-PmemDisk [<CommonParameters>]
Get-PmemDisk [[-DiskNumber] <uint32[]>] [<CommonParameters>]
Get-PmemDisk [-PhysicalDevice <PmemPhysicalDevice>] [<CommonParameters>]
Get-PmemDisk [-PhysicalDeviceId <string[]>] [<CommonParameters>]
Get-PmemDisk [-InputObject <ciminstance>] [<CommonParameters>]
```

示例：获取持久内存磁盘

```powershell
Get-PmemDisk
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemdisk?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemdisk?view=powershell-7.5)

### Get-PmemPhysicalDevice

版本：都有

模块：PersistentMemory

语法：

```powershell
Get-PmemPhysicalDevice [<CommonParameters>]
Get-PmemPhysicalDevice [[-DeviceId] <string[]>] [<CommonParameters>]
Get-PmemPhysicalDevice [-LogicalDisk <PmemDisk>] [<CommonParameters>]
Get-PmemPhysicalDevice [-DiskNumber <uint32>] [<CommonParameters>]
Get-PmemPhysicalDevice [-InputObject <ciminstance>] [<CommonParameters>]
```

示例：获取有持久内存的物理设备

```powershell
Get-PmemPhysicalDevice
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemphysicaldevice?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemphysicaldevice?view=powershell-7.5)

### Get-PmemUnusedRegion

版本：都有

模块：PersistentMemory

语法：

```powershell
Get-PmemUnusedRegion [[-RegionId] <uint32[]>] [<CommonParameters>]
```

示例：获取未使用区域

```powershell
Get-PmemUnusedRegion
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemunusedregion?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/get-pmemunusedregion?view=powershell-7.5)

### Get-ProcessMitigation

版本：都有

模块：ProcessMitigations

语法：

```powershell
Get-ProcessMitigation [-FullPolicy] [<CommonParameters>]
Get-ProcessMitigation [-Name] <string> [-RunningProcesses] [<CommonParameters>]
Get-ProcessMitigation [-Id] <int[]> [<CommonParameters>]
Get-ProcessMitigation [-RegistryConfigFilePath <string>] [<CommonParameters>]
Get-ProcessMitigation [-System] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Get-ProcessMitigation -Name notepad.exe -RunningProcess
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/processmitigations/get-processmitigation?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/processmitigations/get-processmitigation?view=powershell-7.5)

### Get-ProvisioningPackage

版本：都有

模块：Provisioning

语法：

```powershell
Get-ProvisioningPackage [-PackageId] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Get-ProvisioningPackage [-PackagePath] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Get-ProvisioningPackage [-AllInstalledPackages] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Get-ProvisioningPackage -PackagePath c:\test\testppkg.ppkg
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/get-provisioningpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/get-provisioningpackage?view=powershell-7.5)

### Get-PSSessionCapability

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Get-PSSessionCapability [-ConfigurationName] <string> [-Username] <string> [-Full] [<CommonParameters>]
```

示例：获取用户可用的命令

```powershell
Get-PSSessionCapability -ConfigurationName Endpoint1 -Username 'CONTOSO\User'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/get-pssessioncapability?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/get-pssessioncapability?view=powershell-7.5)

### Get-PSSessionConfiguration

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Get-PSSessionConfiguration [[-Name] <string[]>] [-Force] [<CommonParameters>]
```

示例：在本地计算机上获取会话配置

```powershell
Get-PSSessionConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/get-pssessionconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/get-pssessionconfiguration?view=powershell-7.5)

### Get-PSSnapin

版本：仅5.1

模块：Microsoft.PowerShell.Core

语法：

```powershell
Get-PSSnapin [[-Name] <string[]>] [-Registered] [<CommonParameters>]
```

示例：获取当前已加载的管理单元

```powershell
PS C:\> Get-PSSnapin
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/get-pssnapin?view=powershell-5.1)

### Get-RecoveryManagementPluginAltitude

版本：都有

模块：Dism

语法：

```powershell
Get-RecoveryManagementPluginAltitude -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryManagementPluginAltitude -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Get-RecoveryManagementPluginInfo

版本：都有

模块：Dism

语法：

```powershell
Get-RecoveryManagementPluginInfo -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryManagementPluginInfo -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Get-RecoveryManagementPlugins

版本：都有

模块：Dism

语法：

```powershell
Get-RecoveryManagementPlugins -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryManagementPlugins -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Get-RecoveryRemoteManagementStatus

版本：都有

模块：Dism

语法：

```powershell
Get-RecoveryRemoteManagementStatus -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-RecoveryRemoteManagementStatus -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Get-ReFSDedupSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Get-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Get-ReFSDedupSchedule -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/get-refsdedupschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/get-refsdedupschedule?view=powershell-7.5)

### Get-ReFSDedupScrubSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Get-ReFSDedupScrubSchedule [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Get-ReFSDedupScrubSchedule -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/get-refsdedupscrubschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/get-refsdedupscrubschedule?view=powershell-7.5)

### Get-ReFSDedupStatus

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Get-ReFSDedupStatus [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Get-ReFSDedupStatus -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/get-refsdedupstatus?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/get-refsdedupstatus?view=powershell-7.5)

### Get-ScheduledJob

版本：都有

模块：PSScheduledJob

语法：

```powershell
Get-ScheduledJob [[-Id] <int[]>] [<CommonParameters>]
Get-ScheduledJob [-Name] <string[]> [<CommonParameters>]
```

示例：获取所有计划作业

```powershell
Get-ScheduledJob
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/get-scheduledjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/get-scheduledjob?view=powershell-7.5)

### Get-ScheduledJobOption

版本：都有

模块：PSScheduledJob

语法：

```powershell
Get-ScheduledJobOption [-InputObject] <ScheduledJobDefinition> [<CommonParameters>]
Get-ScheduledJobOption [-Id] <int> [<CommonParameters>]
Get-ScheduledJobOption [-Name] <string> [<CommonParameters>]
```

示例：获取作业选项

```powershell
Get-ScheduledJobOption -Name "*Backup*"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/get-scheduledjoboption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/get-scheduledjoboption?view=powershell-7.5)

### Get-SecureBootPolicy

版本：都有

模块：SecureBoot

语法：

```powershell
Get-SecureBootPolicy [<CommonParameters>]
```

示例：获取安全启动策略

```powershell
PS C:\> Get-SecureBootPolicy | Format-List
Publisher: 77fa9abd-0359-4d32-bd60-28f4e78f784b
Version : 1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/get-securebootpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/get-securebootpolicy?view=powershell-7.5)

### Get-SecureBootSVN

版本：都有

模块：SecureBoot

语法：

```powershell
Get-SecureBootSVN [-BootManagerPath <string>] [<CommonParameters>]
```

示例：暂无

出处：[SecureBoot 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/secureboot)（没有单独介绍）。

### Get-SecureBootUEFI

版本：都有

模块：SecureBoot

语法：

```powershell
Get-SecureBootUEFI [-Name] <string> [-OutputFilePath <string>] [-Decoded] [<CommonParameters>]
```

示例：获取 PK 信息

```powershell
PS C:\>Get-SecureBootUefi -Name PK | Format-List
Name : PK
Bytes : {161, 89, 192, 165...}
Attributes : NON VOLATILE
 BOOTSERVICE ACCESS
 RUNTIME ACCESS
 TIME BASED AUTHENTICATED WRITE ACCESS
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/get-securebootuefi?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/get-securebootuefi?view=powershell-7.5)

### Get-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法（5.1）：

```powershell
Get-Service [[-Name] <string[]>] [-ComputerName <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service -DisplayName <string[]> [-ComputerName <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service [-ComputerName <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [-InputObject <ServiceController[]>] [<CommonParameters>]
```

语法（7）：

```powershell
Get-Service [[-Name] <string[]>] [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service -DisplayName <string[]> [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [<CommonParameters>]
Get-Service [-DependentServices] [-RequiredServices] [-Include <string[]>] [-Exclude <string[]>] [-InputObject <ServiceController[]>] [<CommonParameters>]
```

示例：获取计算机上的所有服务

```powershell
Get-Service
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-service?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`systemctl`）。
- 发行版：systemd 系。
- 功能：列服务。映射到 `systemctl list-units --type=service --all --no-pager`。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 按服务名过滤（支持通配），如 `Get-Service -Name ssh*` |

- 实现：调用外部 `systemctl`（权限不足时自动加 `sudo`），解析输出第 0/2 列（单元名、状态）。
- 输出：ServiceController 对象，字段 Name（去掉 .service 后缀）、Status（active/inactive/...）、DisplayName（同 Name）。表格列 Status/Name/DisplayName。
- 无 systemctl → 报错。

### Get-SystemDriver

版本：都有

模块：ConfigCI

语法：

```powershell
Get-SystemDriver [-Audit] [-ScanPath <string>] [-UserPEs] [-NoScript] [-NoShadowCopy] [-OmitPaths <string[]>] [-PathToCatroot <string>] [-ScriptFileNames] [<CommonParameters>]
```

示例：扫描文件夹中的驱动程序

```powershell
PS C:\> Get-SystemDriver -ScanPath '.\temp' -UserPEs

FilePath : \\?\GLOBALROOT\Device\HarddiskVolumeShadowCopy9\cmdlets\temp\ConfigCI.psd1
FriendlyName : \\?\E:\cmdlets\temp\ConfigCI.psd1
FileName :
Loaded : False
FileVersion :
Hash : 1844B4531711EC9170A9D33277CE1D4FF7626C54
Hash256 : 60311157F6685727F42CC04717FEF6F905EC2A317C3B8381CDD9A79D0B184483
PageHash :
PageHash256 :
UserMode : True
OpusInfos : {}
Signers : {}

FilePath : \\?\GLOBALROOT\Device\HarddiskVolumeShadowCopy9\cmdlets\temp\Microsoft.ConfigCI.Commands.dll
FriendlyName : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll
FileName : Microsoft.ConfigCI.Commands.dll
Loaded : False
FileVersion : 10.0.10543.1000
Hash : BE0777F5AF88628D4555A875036648DF1AD19BBE
Hash256 : 6FA5AF724499C338A77FEEAD90F55DDF5F23D081C6DCE8E9DF486E95C6A9B310
PageHash : D41570F2E6E7E6245CF342131D4706C944562B1E
PageHash256 : F714D9784E15B88F56180C8EE2B40C769CC83428954585A1DCF9A260FE967CDD
UserMode : False
OpusInfos : {}
Signers : {}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-systemdriver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/get-systemdriver?view=powershell-7.5)

### Get-SystemPreferredUILanguage

版本：都有

模块：LanguagePackManagement

语法：

```powershell
Get-SystemPreferredUILanguage [<CommonParameters>]
```

示例：

```powershell
Get-SystemPreferredUILanguage
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/get-systempreferreduilanguage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/get-systempreferreduilanguage?view=powershell-7.5)

### Get-TlsCipherSuite

版本：都有

模块：TLS

语法：

```powershell
Get-TlsCipherSuite [[-Name] <string>] [<CommonParameters>]
```

示例：获取全部密码套件

```powershell
Get-TlsCipherSuite
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/get-tlsciphersuite?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/get-tlsciphersuite?view=powershell-7.5)

### Get-TlsEccCurve

版本：都有

模块：TLS

语法：

```powershell
Get-TlsEccCurve [[-Name] <string>] [<CommonParameters>]
```

示例：获取全部 ECC 曲线

```powershell
Get-TlsEccCurve
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/get-tlsecccurve?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/get-tlsecccurve?view=powershell-7.5)

### Get-Tpm

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Get-Tpm [<CommonParameters>]
```

示例：显示 TPM 信息

```powershell
PS C:\> Get-Tpm
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/get-tpm?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/get-tpm?view=powershell-7.5)

### Get-TpmEndorsementKeyInfo

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Get-TpmEndorsementKeyInfo [[-HashAlgorithm] <string>] [<CommonParameters>]
```

示例：获取背书密钥信息

```powershell
PS C:\> Get-TpmEndorsementKeyInfo -Hash "Sha256"
IsPresent : True
PublicKey : System.Security.Cryptography.AsnEncodedData
PublicKeyHash : 70769c52b6e24ef683693c2a0208da68d77e94192e1f4080ae7c9b97c6caa681
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
AdditionalCertificates : {}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/get-tpmendorsementkeyinfo?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/get-tpmendorsementkeyinfo?view=powershell-7.5)

### Get-TpmSupportedFeature

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Get-TpmSupportedFeature [[-FeatureList] <StringCollection>] [<CommonParameters>]
```

示例：验证密钥证明支持

```powershell
PS C:\> Get-TpmSupportedFeature -FeatureList "Key Attestation"
key attestation
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/get-tpmsupportedfeature?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/get-tpmsupportedfeature?view=powershell-7.5)

### Get-Transaction

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-Transaction [<CommonParameters>]
```

示例：获取当前事务

```powershell
Start-Transaction
Get-Transaction
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-transaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-transaction?view=powershell-7.5)

### Get-TroubleshootingPack

版本：都有

模块：TroubleshootingPack

语法：

```powershell
Get-TroubleshootingPack [-Path] <string> [-AnswerFile <string>] [<CommonParameters>]
```

示例：获取疑难解答包

```powershell
PS C:\> Get-TroubleshootingPack -Path "C:\Windows\Diagnostics\System\Audio"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/troubleshootingpack/get-troubleshootingpack?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/troubleshootingpack/get-troubleshootingpack?view=powershell-7.5)

### Get-TrustedProvisioningCertificate

版本：都有

模块：Provisioning

语法：

```powershell
Get-TrustedProvisioningCertificate [[-Thumbprint] <string>] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：列出已安装受信任预配证书

```powershell
PS C:\> Get-TrustedProvisioningCertificate
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/get-trustedprovisioningcertificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/get-trustedprovisioningcertificate?view=powershell-7.5)

### Get-UevAppxPackage

版本：仅5.1

模块：UEV

语法：

```powershell
Get-UevAppxPackage [<CommonParameters>]
Get-UevAppxPackage -Computer [<CommonParameters>]
Get-UevAppxPackage -CurrentComputerUser [<CommonParameters>]
```

示例：获取 Windows 8 应用列表

```powershell
PS C:\>Get-UevAppxPackage -CurrentComputerUser
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/get-uevappxpackage?view=powershell-5.1)

### Get-UevConfiguration

版本：仅5.1

模块：UEV

语法：

```powershell
Get-UevConfiguration [<CommonParameters>]
Get-UevConfiguration -Computer [<CommonParameters>]
Get-UevConfiguration -CurrentComputerUser [<CommonParameters>]
Get-UevConfiguration -Details [<CommonParameters>]
```

示例：获取 uev_tla 配置

```powershell
PS C:\> Get-UevConfiguration

Key Value
--- -----
MaxPackageSizeInBytes 700000
SettingsImportNotifyDelayInSeconds 10
SettingsImportNotifyEnabled False
SettingsStoragePath \\ServerName\Path\To\CentralStore
SettingsTemplateCatalogPath
SyncEnabled True
SyncMethod OfflineFiles
SyncFromRepositoryTimeoutInMilliseconds 2000
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/get-uevconfiguration?view=powershell-5.1)

### Get-UevStatus

版本：仅5.1

模块：UEV

语法：

```powershell
Get-UevStatus [<CommonParameters>]
```

示例：获取 UE-V 服务状态

```powershell
PS C:\>Get-UevStatus
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/get-uevstatus?view=powershell-5.1)

### Get-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Get-UevTemplate [<CommonParameters>]
Get-UevTemplate -Application <string> [<CommonParameters>]
Get-UevTemplate -TemplateID <string> [<CommonParameters>]
Get-UevTemplate -Profile <string> [<CommonParameters>]
Get-UevTemplate [-ApplicationOrTemplateID] <string> [<CommonParameters>]
```

示例：获取全部已注册模板

```powershell
PS C:\> Get-UevTemplate | Format-Table -AutoSize
TemplateId TemplateName TemplateVersion PackageVersion TemplateType Enabled EnableStateLocation TemplateDescription
---------- ------------ --------------- -------------- ------------ ------- ------------------- -------------------
DesktopSettings Desktop Settings 1 N/A OS False LocalMachine
MicrosoftNotepad6 Microsoft Notepad 0 N/A Application True NotSet
MicrosoftCalculator6 Microsoft Calculator 0 N/A Application True NotSet
MicrosoftCommunicator2007 Microsoft Communicator 2007 7 N/A Application True NotSet
MicrosoftOffice2010Win64 Microsoft Office 2010 (64-bit) 18 N/A Application True NotSet
MicrosoftOffice2010Win64.common Common Settings 18 N/A Application True NotSet
MicrosoftOffice2010Win64.Access Microsoft Access 2010 (64-bit) 18 N/A Application True NotSet
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/get-uevtemplate?view=powershell-5.1)

### Get-UevTemplateProgram

版本：仅5.1

模块：UEV

语法：

```powershell
Get-UevTemplateProgram [-ID] <string> [<CommonParameters>]
```

示例：获取全部已定义程序

```powershell
PS C:\> Get-UevTemplate | Get-UevTemplateProgram | Format-Table -AutoSize

TemplateId ProgramName ProductVersionRange FileVersionRange
---------- ----------- ------------------- ----------------
MicrosoftCalculator6 CALC.EXE 6-6
MicrosoftNotepad6 NOTEPAD.EXE 6-6
MicrosoftOffice2010.OneNote ONENOTE.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Word WINWORD.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Excel EXCEL.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.PowerPoint POWERPNT.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Outlook OUTLOOK.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.InfoPath INFOPATH.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Visio VISIO.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Groove Groove.exe 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Access MSACCESS.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Project WINPROJ.EXE 14.0-14.0 14.0-14.0
MicrosoftOffice2010.Publisher MSPUB.EXE 14.0-14.0 14.0-14.0
MicrosoftWordpad6 WORDPAD.EXE 6-6
MicrosoftInternetExplorer.Version8 iexplore.exe 8.0-8.0 8.0-8.0
MicrosoftInternetExplorer.Version9 iexplore.exe 9.0-9.0 9.0-9.0
MicrosoftInternetExplorer.Version10 iexplore.exe 10.0-10.0 10.0-10.0
MicrosoftLync2010 communicator.exe 4.0-4.0 4.0-4.0
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/get-uevtemplateprogram?view=powershell-5.1)

### Get-WheaMemoryPolicy

版本：都有

模块：Whea

语法：

```powershell
Get-WheaMemoryPolicy [-ComputerName <string>] [<CommonParameters>]
```

示例：从本地计算机获取 WHEA 内存策略设置

```powershell
PS C:\> Get-WHEAMemoryPolicy
DisableOffline : False
DisablePFA : False
PersistMemoryOffline : True
PFAPageCount : 64
PFAErrorThreshold : 16
PFATimeOut : 86400
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/whea/get-wheamemorypolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/whea/get-wheamemorypolicy?view=powershell-7.5)

### Get-WIMBootEntry

版本：都有

模块：Dism

语法：

```powershell
Get-WIMBootEntry -Path <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：显示驱动器的 WIMBoot 配置

```powershell
PS C:\> Get-WIMBootEntry -Path "C:\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-wimbootentry?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-wimbootentry?view=powershell-7.5)

### Get-WinAcceptLanguageFromLanguageListOptOut

版本：都有

模块：International

语法：

```powershell
Get-WinAcceptLanguageFromLanguageListOptOut [<CommonParameters>]
```

示例：获取该设置的状态

```powershell
PS C:\> Get-WinAcceptLanguageFromLanguageListOptOut
TRUE
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winacceptlanguagefromlanguagelistoptout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winacceptlanguagefromlanguagelistoptout?view=powershell-7.5)

### Get-WinCultureFromLanguageListOptOut

版本：都有

模块：International

语法：

```powershell
Get-WinCultureFromLanguageListOptOut [<CommonParameters>]
```

示例：获取区域替代设置

```powershell
PS C:\> Get-WinCultureFromLanguageListOptOut
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winculturefromlanguagelistoptout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winculturefromlanguagelistoptout?view=powershell-7.5)

### Get-WinDefaultInputMethodOverride

版本：都有

模块：International

语法：

```powershell
Get-WinDefaultInputMethodOverride [<CommonParameters>]
```

示例：显示默认输入法

```powershell
PS C:\> Get-WinDefaultInputMethodOverride
InputMethodTip Keyboard name
--------------- -------------
0409:00000409 English (United States) - US
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-windefaultinputmethodoverride?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-windefaultinputmethodoverride?view=powershell-7.5)

### Get-WindowsCapability

版本：都有

模块：Dism

语法：

```powershell
Get-WindowsCapability -Path <string> [-Name <string>] [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsCapability -Online [-Name <string>] [-LimitAccess] [-Source <string[]>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：获取映像的 Windows 功能

```powershell
PS C:\> Get-WindowsCapability -Path "C:\offline" -Name "Language.TextToSpeech~~~fr-FR~0.0.1.0"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowscapability?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowscapability?view=powershell-7.5)

### Get-WindowsDeveloperLicense

版本：都有

模块：WindowsDeveloperLicense

语法：

```powershell
Get-WindowsDeveloperLicense [<CommonParameters>]
```

示例：检查开发者模式 DM 状态

```powershell
PS C:\> Get-WindowsDeveloperLicense
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowsdeveloperlicense/get-windowsdeveloperlicense?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowsdeveloperlicense/get-windowsdeveloperlicense?view=powershell-7.5)

### Get-WindowsDriver

版本：都有

模块：Dism

语法：

```powershell
Get-WindowsDriver -Path <string> [-All] [-Driver <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsDriver -Online [-All] [-Driver <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：获取联机映像中的全部驱动程序

```powershell
PS C:\> Get-WindowsDriver -Online -All
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsdriver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsdriver?view=powershell-7.5)

### Get-WindowsEdition

版本：都有

模块：Dism

语法：

```powershell
Get-WindowsEdition -Path <string> [-Target] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsEdition -Online [-Target] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：显示操作系统当前版本

```powershell
PS C:\> Get-WindowsEdition -Online
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsedition?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsedition?view=powershell-7.5)

### Get-WindowsErrorReporting

版本：都有

模块：WindowsErrorReporting

语法：

```powershell
Get-WindowsErrorReporting [<CommonParameters>]
```

示例：获取 Windows 错误报告状态

```powershell
PS C:\> Get-WindowsErrorReporting
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowserrorreporting/get-windowserrorreporting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowserrorreporting/get-windowserrorreporting?view=powershell-7.5)

### Get-WindowsImage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Get-WindowsImage -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Index <uint32> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -Mounted [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Get-WindowsImage -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -ImagePath <string> -Index <uint> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImage -Mounted [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：获取全部已装载映像的信息

```powershell
PS C:\> Get-WindowsImage -Mounted
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsimage?view=powershell-7.5)

### Get-WindowsImageContent

版本：都有

模块：Dism

语法（5.1）：

```powershell
Get-WindowsImageContent -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> -Index <uint32> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Get-WindowsImageContent -ImagePath <string> -Name <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> -Index <uint> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsImageContent -ImagePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：列出映像的文件和文件夹

```powershell
PS C:\> Get-WindowsImageContent -ImagePath "c:\imagestore\install.wim" -Index 1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsimagecontent?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsimagecontent?view=powershell-7.5)

### Get-WindowsOptionalFeature

版本：都有

模块：Dism

语法：

```powershell
Get-WindowsOptionalFeature -Path <string> [-FeatureName <string>] [-PackageName <string>] [-PackagePath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsOptionalFeature -Online [-FeatureName <string>] [-PackageName <string>] [-PackagePath <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：列出正在运行的操作系统中的可选功能

```powershell
PS C:\> Get-WindowsOptionalFeature -Online
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsoptionalfeature?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsoptionalfeature?view=powershell-7.5)

### Get-WindowsPackage

版本：都有

模块：Dism

语法：

```powershell
Get-WindowsPackage -Path <string> [-PackagePath <string>] [-PackageName <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Get-WindowsPackage -Online [-PackagePath <string>] [-PackageName <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：列出已装载映像中的包

```powershell
PS C:\> Get-WindowsPackage -Path "c:\offline"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowspackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowspackage?view=powershell-7.5)

### Get-WindowsReservedStorageState

版本：都有

模块：Dism

语法：

```powershell
Get-WindowsReservedStorageState [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Get-WindowsReservedStorageState
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsreservedstoragestate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/get-windowsreservedstoragestate?view=powershell-7.5)

### Get-WindowsSearchSetting

版本：都有

模块：WindowsSearch

语法：

```powershell
Get-WindowsSearchSetting [<CommonParameters>]
```

示例：获取 Windows 搜索设置

```powershell
PS C:\> Get-WindowsSearchSetting
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowssearch/get-windowssearchsetting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowssearch/get-windowssearchsetting?view=powershell-7.5)

### Get-WinEvent

版本：都有

模块：Microsoft.PowerShell.Diagnostics

语法：

```powershell
Get-WinEvent [[-LogName] <string[]>] [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-FilterXPath <string>] [-Force] [-Oldest] [<CommonParameters>]
Get-WinEvent [-ListLog] <string[]> [-ComputerName <string>] [-Credential <pscredential>] [-Force] [<CommonParameters>]
Get-WinEvent [-ListProvider] <string[]> [-ComputerName <string>] [-Credential <pscredential>] [<CommonParameters>]
Get-WinEvent [-ProviderName] <string[]> [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-FilterXPath <string>] [-Force] [-Oldest] [<CommonParameters>]
Get-WinEvent [-Path] <string[]> [-MaxEvents <long>] [-Credential <pscredential>] [-FilterXPath <string>] [-Oldest] [<CommonParameters>]
Get-WinEvent [-FilterHashtable] <hashtable[]> [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-Force] [-Oldest] [<CommonParameters>]
Get-WinEvent [-FilterXml] <xml> [-MaxEvents <long>] [-ComputerName <string>] [-Credential <pscredential>] [-Oldest] [<CommonParameters>]
```

示例：从本地计算机获取所有日志

```powershell
Get-WinEvent -ListLog *
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/get-winevent?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/get-winevent?view=powershell-7.5)

### Get-WinHomeLocation

版本：都有

模块：International

语法：

```powershell
Get-WinHomeLocation [<CommonParameters>]
```

示例：显示当前帐户的 GeoID

```powershell
PS C:\> Get-WinHomeLocation
HomeLocation Description
---- -----------
244 United States
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winhomelocation?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winhomelocation?view=powershell-7.5)

### Get-WinLanguageBarOption

版本：都有

模块：International

语法：

```powershell
Get-WinLanguageBarOption [<CommonParameters>]
```

示例：获取语言栏设置

```powershell
PS C:\> Get-WinLanguageBarOption
IsLegacyLanguageBar IsLegacySwitchingMode
------------------- ---------------------
False False
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winlanguagebaroption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winlanguagebaroption?view=powershell-7.5)

### Get-WinSystemLocale

版本：都有

模块：International

语法：

```powershell
Get-WinSystemLocale [<CommonParameters>]
```

示例：获取系统区域

```powershell
PS C:\> GET-WinSystemLocale
LCID Name DisplayName
---- ---- -----------
1033 en-US English (United States)
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winsystemlocale?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winsystemlocale?view=powershell-7.5)

### Get-WinUILanguageOverride

版本：都有

模块：International

语法：

```powershell
Get-WinUILanguageOverride [<CommonParameters>]
```

示例：显示语言替代设置

```powershell
PS C:\> Get-WinUILanguageOverride
LCID Name DisplayName
---- ---- -----------
1033 en-US English (United States)
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winuilanguageoverride?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winuilanguageoverride?view=powershell-7.5)

### Get-WinUserLanguageList

版本：都有

模块：International

语法：

```powershell
Get-WinUserLanguageList [<CommonParameters>]
```

示例：获取当前帐户的语言列表

```powershell
PS C:\> Get-WinUserLanguageList
LanguageTag : en-US
Autonym : English (United States)
EnglishName : English (United States)
LocalizedName : English (United States)
ScriptName : Latin
InputMethodTips : {0409:00000409}
Handwriting : False
LanguageTag : fr-FR
Autonym : français (France)
EnglishName : French (France)
LocalizedName : French (France)
ScriptName : Latin
InputMethodTips : {040c:0000040c}
Handwriting : False
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winuserlanguagelist?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/get-winuserlanguagelist?view=powershell-7.5)

### Get-WmiObject

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Get-WmiObject [-Class] <string> [[-Property] <string[]>] [-Filter <string>] [-Amended] [-DirectRead] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject [[-Class] <string>] [-Recurse] [-Amended] [-List] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject -Query <string> [-Amended] [-DirectRead] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject [-Amended] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
Get-WmiObject [-Amended] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [<CommonParameters>]
```

示例：获取本地计算机上的进程

```powershell
Get-WmiObject -Class Win32_Process
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-wmiobject?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/get-wmiobject?view=powershell-7.5)

### Get-WSManCredSSP

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Get-WSManCredSSP [<CommonParameters>]
```

示例：显示 CredSSP 配置

```powershell
Get-WSManCredSSP
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/get-wsmancredssp?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/get-wsmancredssp?view=powershell-7.5)

### Get-WSManInstance

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Get-WSManInstance [-ResourceURI] <uri> [-ApplicationName <string>] [-ComputerName <string>] [-ConnectionURI <uri>] [-Dialect <uri>] [-Fragment <string>] [-OptionSet <hashtable>] [-Port <int>] [-SelectorSet <hashtable>] [-SessionOption <SessionOption>] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Get-WSManInstance [-ResourceURI] <uri> -Enumerate [-ApplicationName <string>] [-BasePropertiesOnly] [-ComputerName <string>] [-ConnectionURI <uri>] [-Dialect <uri>] [-Filter <string>] [-OptionSet <hashtable>] [-Port <int>] [-Associations] [-ReturnType <string>] [-SessionOption <SessionOption>] [-Shallow] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：从 WMI 获取所有信息

```powershell
Get-WSManInstance -ResourceURI wmicimv2/Win32_Service -SelectorSet @{name="winrm"} -ComputerName "Server01"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/get-wsmaninstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/get-wsmaninstance?view=powershell-7.5)

### Import-BcdStore

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Import-BcdStore [-Path] <string> [-NoClean] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/import-bcdstore?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/import-bcdstore?view=powershell-7.5)

### Import-BinaryMiLog

版本：都有

模块：CimCmdlets

语法：

```powershell
Import-BinaryMiLog [-Path] <string> [<CommonParameters>]
```

示例：还原导出到文件的对象

```powershell
Import-BinaryMiLog -Path "Processes.bmil"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/import-binarymilog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/import-binarymilog?view=powershell-7.5)

### Import-Certificate

版本：都有

模块：PKI

语法：

```powershell
Import-Certificate [-FilePath] <string> [-CertStoreLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 FilePath = 'C:\Users\Xyz\Desktop\BackupCert.cer'
 CertStoreLocation = 'Cert:\CurrentUser\Root'
}
Import-Certificate @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/import-certificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/import-certificate?view=powershell-7.5)

### Import-Counter

版本：都有

模块：Microsoft.PowerShell.Diagnostics

语法：

```powershell
Import-Counter [-Path] <string[]> [-StartTime <datetime>] [-EndTime <datetime>] [-Counter <string[]>] [-MaxSamples <long>] [<CommonParameters>]
Import-Counter [-Path] <string[]> -ListSet <string[]> [<CommonParameters>]
Import-Counter [-Path] <string[]> [-Summary] [<CommonParameters>]
```

示例：从文件导入所有计数器数据

```powershell
$data = Import-Counter -Path ProcessorData.csv
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/import-counter?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/import-counter?view=powershell-7.5)

### Import-PfxCertificate

版本：都有

模块：PKI

语法：

```powershell
Import-PfxCertificate [-FilePath] <string> [[-CertStoreLocation] <string>] [-Exportable] [-ProtectPrivateKey <ProtectPrivateKeyType>] [-Password <securestring>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$mypwd = Get-Credential -UserName 'Enter password below' -Message 'Enter password below'

$params = @{
 FilePath = 'C:\mypfx.pfx'
 CertStoreLocation = 'Cert:\LocalMachine\My'
 Password = $mypwd.Password
}
Import-PfxCertificate @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/import-pfxcertificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/import-pfxcertificate?view=powershell-7.5)

### Import-StartLayout

版本：都有

模块：StartLayout

语法：

```powershell
Import-StartLayout [-LayoutPath] <string> [-MountPath] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Import-StartLayout -LayoutLiteralPath <string> -MountLiteralPath <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将布局导入 Windows 映像

```powershell
PS C:\> Import-StartLayout -LayoutPath "Layout.xml" -MountPath "C:\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/startlayout/import-startlayout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/startlayout/import-startlayout?view=powershell-7.5)

### Import-TpmOwnerAuth

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Import-TpmOwnerAuth -File <string> [<CommonParameters>]
Import-TpmOwnerAuth [-OwnerAuthorization] <string> [<CommonParameters>]
```

示例：导入所有者授权值

```powershell
PS C:\> Import-TpmOwnerAuth -OwnerAuthorization "Qn2sdCFQmvjf+tBtSWH4GT87sQs="
TpmReady : False
TpmPresent : True
ManagedAuthLevel : Full
OwnerAuth : Qn2sdCFQmvjf+tBtSWH4GT87sQs=
OwnerClearDisabled : True
AutoProvisioning : DisabledForNextBoot
LockedOut : False
SelfTest : {191, 191, 245, 191...}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/import-tpmownerauth?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/import-tpmownerauth?view=powershell-7.5)

### Import-UevConfiguration

版本：仅5.1

模块：UEV

语法：

```powershell
Import-UevConfiguration [-Path] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：导入 UE-V 配置

```powershell
PS C:\> Import-UevConfiguration -Path "ContosoUev.uev"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/import-uevconfiguration?view=powershell-5.1)

### Initialize-PmemPhysicalDevice

版本：都有

模块：PersistentMemory

语法：

```powershell
Initialize-PmemPhysicalDevice -DeviceId <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：初始化物理设备

```powershell
Get-PmemPhysicalDevice | Initialize-PmemPhysicalDevice
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/initialize-pmemphysicaldevice?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/initialize-pmemphysicaldevice?view=powershell-7.5)

### Initialize-Tpm

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Initialize-Tpm [-AllowClear] [-AllowPhysicalPresence] [<CommonParameters>]
```

示例：初始化 TPM

```powershell
PS C:\> Initialize-Tpm -AllowClear -AllowPhysicalPresence
TpmReady : False
RestartRequired : True
ShutdownRequired : False
ClearRequired : True
PhysicalPresenceRequired : True
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/initialize-tpm?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/initialize-tpm?view=powershell-7.5)

### Install-Language

版本：都有

模块：LanguagePackManagement

语法：

```powershell
Install-Language [-Language] <string> [-CopyToSettings] [-ExcludeFeatures] [-AsJob] [<CommonParameters>]
```

示例：向设备添加语言

```powershell
Install-Language ja-JP
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/install-language?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/install-language?view=powershell-7.5)

### Install-ProvisioningPackage

版本：都有

模块：Provisioning

语法：

```powershell
Install-ProvisioningPackage [-PackagePath] <string> [-ForceInstall] [-QuietInstall] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Install-ProvisioningPackage -PackagePath C:\mypackage.ppkg -QuietInstall
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/install-provisioningpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/install-provisioningpackage?view=powershell-7.5)

### Install-TrustedProvisioningCertificate

版本：都有

模块：Provisioning

语法：

```powershell
Install-TrustedProvisioningCertificate [-CertificatePath] <string> [-ForceInstall] [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：安装受信任预配证书

```powershell
PS C:\> Install-TrustedProvisioningCertificate -CertificatePath trustedCert.cer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/install-trustedprovisioningcertificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/install-trustedprovisioningcertificate?view=powershell-7.5)

### Invoke-CimMethod

版本：都有

模块：CimCmdlets

语法（5.1）：

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

语法（7）：

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

示例：调用方法

```powershell
$method = @{
 Query = 'select * from Win32_Process where name like "notepad%"'
 MethodName = "Terminate"
}
Invoke-CimMethod @method
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/invoke-cimmethod?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/invoke-cimmethod?view=powershell-7.5)

### Invoke-CommandInDesktopPackage

版本：都有

模块：Appx

语法：

```powershell
Invoke-CommandInDesktopPackage [-PackageFamilyName] <string> [[-AppId] <string>] [-Command] <string> [[-Args] <string>] [-PreventBreakaway] [<CommonParameters>]
```

示例：调用记事本读取虚拟化文件

```powershell
$params = @{
 AppId = 'ContosoApp'
 PackageFamilyName = 'Contoso.MyApp_abcdefgh23456'
 Command = 'notepad.exe'
}
Invoke-CommandInDesktopPackage @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/invoke-commandindesktoppackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/invoke-commandindesktoppackage?view=powershell-7.5)

### Invoke-DscResource

版本：都有

模块：PSDesiredStateConfiguration

语法：

```powershell
Invoke-DscResource [-Name] <string> [-Method] <string> -ModuleName <ModuleSpecification> -Property <hashtable> [<CommonParameters>]
```

示例：通过指定资源的必需属性调用 Set 方法

```powershell
Invoke-DscResource -Name WindowsProcess -Method Set -ModuleName PSDesiredStateConfiguration -Property @{
 Path = 'C:\Windows\System32\WindowsPowerShell\v1.0\powershell.exe'
 Arguments = ''
}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/invoke-dscresource?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/invoke-dscresource?view=powershell-7.5)

### Invoke-LapsPolicyProcessing

版本：都有

模块：LAPS

语法：

```powershell
Invoke-LapsPolicyProcessing [<CommonParameters>]
```

示例：

```powershell
Invoke-LapsPolicyProcessing
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/invoke-lapspolicyprocessing?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/invoke-lapspolicyprocessing?view=powershell-7.5)

### Invoke-TroubleshootingPack

版本：都有

模块：TroubleshootingPack

语法：

```powershell
Invoke-TroubleshootingPack [-Pack] <DiagPack> [-AnswerFile <string>] [-Result <string>] [-Unattended] [<CommonParameters>]
```

示例：运行疑难解答包

```powershell
PS C:\> Get-TroubleshootingPack -Path "C:\Windows\Diagnostics\System\Audio" | Invoke-TroubleshootingPack
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/troubleshootingpack/invoke-troubleshootingpack?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/troubleshootingpack/invoke-troubleshootingpack?view=powershell-7.5)

### Invoke-WmiMethod

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Invoke-WmiMethod [-Class] <string> [-Name] <string> [[-ArgumentList] <Object[]>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> -InputObject <wmi> [-ArgumentList <Object[]>] [-AsJob] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> -Path <string> [-ArgumentList <Object[]>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Invoke-WmiMethod [-Name] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：列出 WMI 方法参数的必需顺序

```powershell
Get-WmiObject Win32_Volume |
 Get-Member -MemberType Method -Name Format |
 Select-Object -ExpandProperty Definition
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/invoke-wmimethod?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/invoke-wmimethod?view=powershell-7.5)

### Invoke-WSManAction

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Invoke-WSManAction [-ResourceURI] <uri> [-Action] <string> [[-SelectorSet] <hashtable>] [-ConnectionURI <uri>] [-FilePath <string>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Invoke-WSManAction [-ResourceURI] <uri> [-Action] <string> [[-SelectorSet] <hashtable>] [-ApplicationName <string>] [-ComputerName <string>] [-FilePath <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：调用方法

```powershell
$params = @{
 Action = 'StartService'
 ResourceURI = 'wmicimv2/Win32_Service'
 SelectorSet = @{name = 'spooler'}
 Authentication = 'Default'
}
Invoke-WSManAction @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/invoke-wsmanaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/invoke-wsmanaction?view=powershell-7.5)

### Join-DtcDiagnosticResourceManager

版本：都有

模块：MsDtc

语法：

```powershell
Join-DtcDiagnosticResourceManager [-Transaction] <DtcDiagnosticTransaction> [[-ComputerName] <string>] [[-Port] <int>] [-Volatile] [<CommonParameters>]
```

示例：登记新的诊断事务

```powershell
PS C:\> $Transaction = New-DtcDiagnosticTransaction
PS C:\> Join-DtcDiagnosticResourceManager -Transaction $Transaction
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/join-dtcdiagnosticresourcemanager?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/join-dtcdiagnosticresourcemanager?view=powershell-7.5)

### Limit-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Limit-EventLog [-LogName] <string[]> [-ComputerName <string[]>] [-RetentionDays <int>] [-OverflowAction <OverflowAction>] [-MaximumSize <long>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：增加事件日志的大小

```powershell
Limit-EventLog -LogName "Windows PowerShell" -MaximumSize 20KB
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/limit-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/limit-eventlog?view=powershell-7.5)

### Merge-CIPolicy

版本：都有

模块：ConfigCI

语法：

```powershell
Merge-CIPolicy [-OutputFilePath] <string> [-PolicyPaths] <string[]> [-Rules <Rule[]>] [-AppIdTaggingPolicy] [<CommonParameters>]
```

示例：合并策略

```powershell
PS C:\> Merge-CIPolicy -PolicyPaths '.\Policy.xml','.\Policy02.xml' -OutputFilePath '.\MergedPolicy.xml'

Name : MSIT Test CodeSign CA 3
Id : ID_SIGNER_S_17_0
TypeId : Allow
Root : FA6B9A2230CE08BCA81D096B28CF495672401D3A43A0D285CF352464A6C9C7FD
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : VeriSign Class 3 Code Signing 2010 CA
Id : ID_SIGNER_S_1D_0
TypeId : Allow
Root : 4843A82ED3B1F2BFBEE9671960E1940C942F688D
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : Microsoft Windows Third Party Component CA 2012
Id : ID_SIGNER_S_1E_0
TypeId : Allow
Root : CEC1AFD0E310C55C1DCC601AB8E172917706AA32FB5EAF826813547FDF02DD46
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll Hash Sha1
Id : ID_ALLOW_A_49_1
TypeId : Allow
Root :
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/merge-cipolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/merge-cipolicy?view=powershell-7.5)

### Mount-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Mount-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [<CommonParameters>]
Mount-AppvClientConnectionGroup [-Name] <string> [<CommonParameters>]
Mount-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [<CommonParameters>]
```

示例：为指定组下载包

```powershell
PS C:\> Mount-AppvClientConnectionGroup -Name "MyGroup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/mount-appvclientconnectiongroup?view=powershell-5.1)

### Mount-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Mount-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Cancel] [<CommonParameters>]
Mount-AppvClientPackage [-Package] <AppvClientPackage> [-Cancel] [<CommonParameters>]
Mount-AppvClientPackage [-Name] <string> [[-Version] <string>] [<CommonParameters>]
```

示例：获取包的指定版本

```powershell
PS C:\> Mount-AppvClientPackage -Name "MyApp" -Version 2
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/mount-appvclientpackage?view=powershell-5.1)

### Mount-AppxVolume

版本：都有

模块：Appx

语法：

```powershell
Mount-AppxVolume [-Volume] <AppxVolume[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按路径装载卷

```powershell
Mount-AppxVolume -Volume E:\
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/mount-appxvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/mount-appxvolume?view=powershell-7.5)

### Mount-WindowsImage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Mount-WindowsImage -Path <string> -ImagePath <string> -Index <uint32> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> -Name <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -Remount [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Mount-WindowsImage -Path <string> -ImagePath <string> -Index <uint> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> -Name <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -ImagePath <string> [-ReadOnly] [-Optimize] [-CheckIntegrity] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Mount-WindowsImage -Path <string> -Remount [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：将 install.vhd 文件中的映像装载到目录

```powershell
PS C:\> Mount-WindowsImage -ImagePath "c:\imagestore\install.vhd" -Index 1 -Path "c:\offline"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/mount-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/mount-windowsimage?view=powershell-7.5)

### Move-AppxPackage

版本：都有

模块：Appx

语法：

```powershell
Move-AppxPackage [-Package] <string[]> [-Volume] <AppxVolume> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将包移到按路径指定的卷

```powershell
Move-AppxPackage -Package "package1_1.0.0.0_neutral__8wekyb3d8bbwe" -Volume F:\
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/move-appxpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/move-appxpackage?view=powershell-7.5)

### New-AppLockerPolicy

版本：都有

模块：AppLocker

语法：

```powershell
New-AppLockerPolicy [-FileInformation] <List[FileInformation]> [-AllowWindows] [-RuleType <List[RuleType]>] [-RuleNamePrefix <string>] [-User <string>] [-Optimize] [-IgnoreMissingFileInformation] [-Xml] [-ServiceEnforcement <ServiceEnforcementMode>] [<CommonParameters>]
New-AppLockerPolicy -AllowWindows [-RuleType <List[RuleType]>] [-RuleNamePrefix <string>] [-User <string>] [-Optimize] [-IgnoreMissingFileInformation] [-Xml] [-ServiceEnforcement <ServiceEnforcementMode>] [<CommonParameters>]
```

示例：创建带允许规则的 AppLocker 策略

```powershell
C:\PS>Get-ChildItem C:\Windows\System32\*.exe | Get-AppLockerFileInformation | New-AppLockerPolicy -RuleType Publisher, Hash -User Everyone -RuleNamePrefix System32

 Version RuleCollections RuleCollectionTypes
 ------- --------------- -------------------
 1 {Microsoft.Security.ApplicationId.Po... {Exe}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/new-applockerpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/new-applockerpolicy?view=powershell-7.5)

### New-BcdEntry

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
New-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Application <string> [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
New-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Device [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
New-BcdEntry [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Inherit <string> [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
New-BcdEntry [-Id] <string> [[-Store] <BcdStoreInfo>] [-Description <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/new-bcdentry?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/new-bcdentry?view=powershell-7.5)

### New-BcdStore

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
New-BcdStore [-Path] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/new-bcdstore?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/new-bcdstore?view=powershell-7.5)

### New-CertificateNotificationTask

版本：都有

模块：PKI

语法：

```powershell
New-CertificateNotificationTask -Type <CertificateNotificationType> -PSScript <string> -Name <string> -Channel <NotificationChannel> [-RunTaskForExistingCertificates] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 PSScript = 'C:\myscript.ps1'
 Channel = 'System'
 Type = 'Replace'
 Name = 'My System Certificate Task'
}
New-CertificateNotificationTask @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/new-certificatenotificationtask?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/new-certificatenotificationtask?view=powershell-7.5)

### New-CimInstance

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] [-OperationTimeoutSec <uint32>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-ClassName] <string> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> -CimSession <CimSession[]> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [[-Property] <IDictionary>] -ResourceUri <uri> [-Key <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
New-CimInstance [-CimClass] <cimclass> [[-Property] <IDictionary>] [-OperationTimeoutSec <uint>] [-ComputerName <string[]>] [-ClientOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建 CIM 类的实例

```powershell
$prop = @{
 Name = "testvar"
 VariableValue = "testvalue"
 UserName = "domain\user"
}
New-CimInstance -ClassName Win32_Environment -Property $prop
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/new-ciminstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/new-ciminstance?view=powershell-7.5)

### New-CimSession

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
New-CimSession [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-Authentication <PasswordAuthenticationMechanism>] [-Name <string>] [-OperationTimeoutSec <uint32>] [-SkipTestConnection] [-Port <uint32>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
New-CimSession [[-ComputerName] <string[]>] [-CertificateThumbprint <string>] [-Name <string>] [-OperationTimeoutSec <uint32>] [-SkipTestConnection] [-Port <uint32>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
```

语法（7）：

```powershell
New-CimSession [[-ComputerName] <string[]>] [[-Credential] <pscredential>] [-Authentication <PasswordAuthenticationMechanism>] [-Name <string>] [-OperationTimeoutSec <uint>] [-SkipTestConnection] [-Port <uint>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
New-CimSession [[-ComputerName] <string[]>] [-CertificateThumbprint <string>] [-Name <string>] [-OperationTimeoutSec <uint>] [-SkipTestConnection] [-Port <uint>] [-SessionOption <CimSessionOptions>] [<CommonParameters>]
```

示例：使用默认选项创建 CIM 会话

```powershell
New-CimSession
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/new-cimsession?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/new-cimsession?view=powershell-7.5)

### New-CimSessionOption

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
New-CimSessionOption [-Protocol] <ProtocolType> [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-NoEncryption] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-EncodePortInServicePrincipalName] [-Encoding <PacketEncoding>] [-HttpPrefix <uri>] [-MaxEnvelopeSizeKB <uint32>] [-ProxyAuthentication <PasswordAuthenticationMechanism>] [-ProxyCertificateThumbprint <string>] [-ProxyCredential <pscredential>] [-ProxyType <ProxyType>] [-UseSsl] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-Impersonation <ImpersonationType>] [-PacketIntegrity] [-PacketPrivacy] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
```

语法（7）：

```powershell
New-CimSessionOption [-Protocol] <ProtocolType> [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-NoEncryption] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-EncodePortInServicePrincipalName] [-Encoding <PacketEncoding>] [-HttpPrefix <uri>] [-MaxEnvelopeSizeKB <uint>] [-ProxyAuthentication <PasswordAuthenticationMechanism>] [-ProxyCertificateThumbprint <string>] [-ProxyCredential <pscredential>] [-ProxyType <ProxyType>] [-UseSsl] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
New-CimSessionOption [-Impersonation <ImpersonationType>] [-PacketIntegrity] [-PacketPrivacy] [-UICulture <cultureinfo>] [-Culture <cultureinfo>] [<CommonParameters>]
```

示例：为 DCOM 创建 CIM 会话选项对象

```powershell
$so = New-CimSessionOption -Protocol Dcom
New-CimSession -ComputerName Server01 -SessionOption $so
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/new-cimsessionoption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/new-cimsessionoption?view=powershell-7.5)

### New-CIPolicy

版本：都有

模块：ConfigCI

语法：

```powershell
New-CIPolicy [-FilePath] <string> -Level <RuleLevel> [-DriverFiles <DriverFile[]>] [-Fallback <RuleLevel[]>] [-Audit] [-ScanPath <string>] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-UserPEs] [-NoScript] [-Deny] [-NoShadowCopy] [-MultiplePolicyFormat] [-OmitPaths <string[]>] [-PathToCatroot <string>] [-AppIdTaggingPolicy] [-AppIdTaggingKey <string[]>] [-AppIdTaggingValue <string[]>] [<CommonParameters>]
New-CIPolicy [-FilePath] <string> -Rules <Rule[]> [-Audit] [-ScanPath <string>] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-UserPEs] [-NoScript] [-Deny] [-NoShadowCopy] [-MultiplePolicyFormat] [-OmitPaths <string[]>] [-PathToCatroot <string>] [-AppIdTaggingPolicy] [-AppIdTaggingKey <string[]>] [-AppIdTaggingValue <string[]>] [<CommonParameters>]
```

示例：创建多种策略格式的策略

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

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/new-cipolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/new-cipolicy?view=powershell-7.5)

### New-CIPolicyRule

版本：都有

模块：ConfigCI

语法：

```powershell
New-CIPolicyRule -Level <RuleLevel> [-DriverFiles <DriverFile[]>] [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [<CommonParameters>]
New-CIPolicyRule -DriverFilePath <string[]> -Level <RuleLevel> [-AppID <string>] [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [<CommonParameters>]
New-CIPolicyRule [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-Package <AppxPackage>] [<CommonParameters>]
New-CIPolicyRule [-Fallback <RuleLevel[]>] [-Deny] [-ScriptFileNames] [-AllowFileNameFallbacks] [-SpecificFileNameLevel <FileNameLevel>] [-UserWriteablePaths] [-FilePathRule <string>] [<CommonParameters>]
```

示例：为驱动程序创建策略规则

```powershell
PS C:\> $DriverFiles = Get-SystemDriver -ScanPath '.\temp\' -UserPEs -OmitPaths '.\temp\ConfigCITestBinaries' -NoScript
PS C:\> New-CIPolicyRule -Level FileName -DriverFiles $DriverFiles
Scan completed successfully

Name : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.dll FileRule
Id : ID_ALLOW_A_1
TypeId : Allow
Root :
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : \\?\E:\cmdlets\temp\Microsoft.ConfigCI.Commands.Tests.dll FileRule
Id : ID_ALLOW_A_3
TypeId : Allow
Root :
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False

Name : \\?\E:\cmdlets\temp\Microsoft.PackageInspector.Tests.dll FileRule
Id : ID_ALLOW_A_5
TypeId : Allow
Root :
FileVersionRef :
Wellknown : False
Ekus :
Exceptions :
FileAttributes :
FileException : False
UserMode : False
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/new-cipolicyrule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/new-cipolicyrule?view=powershell-7.5)

### New-DtcDiagnosticTransaction

版本：都有

模块：MsDtc

语法：

```powershell
New-DtcDiagnosticTransaction [[-Timeout] <int>] [[-IsolationLevel] <IsolationLevel>] [<CommonParameters>]
```

示例：创建诊断事务

```powershell
PS C:\> New-DtcDiagnosticTransaction -Timeout 60 -IsolationLevel Serializable
Id
--
4625a5a3-af35-465d-a331-f224d79e4c85
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/new-dtcdiagnostictransaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/new-dtcdiagnostictransaction?view=powershell-7.5)

### New-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
New-EventLog [-LogName] <string> [-Source] <string[]> [[-ComputerName] <string[]>] [-CategoryResourceFile <string>] [-MessageResourceFile <string>] [-ParameterResourceFile <string>] [<CommonParameters>]
```

示例：创建新的事件日志

```powershell
New-EventLog -Source TestApp -LogName TestLog -MessageResourceFile C:\Test\TestApp.dll
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/new-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/new-eventlog?view=powershell-7.5)

### New-FileCatalog

版本：都有

模块：Microsoft.PowerShell.Security

语法：

```powershell
New-FileCatalog [-CatalogFilePath] <string> [[-Path] <string[]>] [-CatalogVersion <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：为“Microsoft.PowerShell.Utility”创建文件目录

```powershell
$newFileCatalogSplat = @{
 Path = "$PSHOME\Modules\Microsoft.PowerShell.Utility"
 CatalogFilePath = '\temp\Microsoft.PowerShell.Utility.cat'
 CatalogVersion = 2.0
}
New-FileCatalog @newFileCatalogSplat
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/new-filecatalog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/new-filecatalog?view=powershell-7.5)

### New-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
New-JobTrigger -Once -At <datetime> [-RandomDelay <timespan>] [-RepetitionInterval <timespan>] [-RepetitionDuration <timespan>] [-RepeatIndefinitely] [<CommonParameters>]
New-JobTrigger -Daily -At <datetime> [-DaysInterval <int>] [-RandomDelay <timespan>] [<CommonParameters>]
New-JobTrigger -Weekly -At <datetime> -DaysOfWeek <DayOfWeek[]> [-WeeksInterval <int>] [-RandomDelay <timespan>] [<CommonParameters>]
New-JobTrigger -AtStartup [-RandomDelay <timespan>] [<CommonParameters>]
New-JobTrigger -AtLogOn [-RandomDelay <timespan>] [-User <string>] [<CommonParameters>]
```

示例：一次计划

```powershell
New-JobTrigger -Once -At "1/20/2012 3:00 AM"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/new-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/new-jobtrigger?view=powershell-7.5)

### New-LocalGroup

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
New-LocalGroup [-Name] <string> [-Description <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建安全组

```powershell
New-LocalGroup -Name "SecurityGroup04"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/new-localgroup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/new-localgroup?view=powershell-7.5)

### New-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
New-LocalUser [-Name] <string> -Password <securestring> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-Disabled] [-FullName <string>] [-PasswordNeverExpires] [-UserMayNotChangePassword] [-WhatIf] [-Confirm] [<CommonParameters>]
New-LocalUser [-Name] <string> -NoPassword [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-Disabled] [-FullName <string>] [-UserMayNotChangePassword] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建用户帐户

```powershell
New-LocalUser -Name 'User02' -Description 'Description of this account.' -NoPassword
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/new-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/new-localuser?view=powershell-7.5)

### New-NetIPsecAuthProposal

版本：都有

模块：NetSecurity

语法：

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

示例：

```powershell
PS C:\>$cert1Proposal = New-NetIPsecAuthProposal -Machine -Cert -Authority "C=US,O=MSFT,CN=ꞌMicrosoft Root Authorityꞌ" -AuthorityType Root

PS C:\>$cert2Proposal = New-NetIPsecAuthProposal -Machine -Cert -Authority "C=US,O=MYORG,CN='My Organizations Root Certificate'" -AuthorityType Root

PS C:\>$certAuthSet = New-NetIPsecPhase1AuthSet -DisplayName "Computer Certificate Auth Set" -Proposal $cert1Proposal,$cert2Proposal

PS C:\>New-NetIPSecRule -DisplayName "Authenticate with Certificates Rule" -InboundSecurity Require -OutboundSecurity Request -Phase2AuthSet $certAuthSet.Name
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/new-netipsecauthproposal?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/new-netipsecauthproposal?view=powershell-7.5)

### New-NetIPsecMainModeCryptoProposal

版本：都有

模块：NetSecurity

语法：

```powershell
New-NetIPsecMainModeCryptoProposal [-Encryption <EncryptionAlgorithm>] [-KeyExchange <DiffieHellmanGroup>] [-Hash <HashAlgorithm>] [<CommonParameters>]
```

示例：

```powershell
PS C:\>$proposal1 = (New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH1)

PS C:\>$proposal2 = (New-NetIPsecMainModeCryptoProposal -Encryption AES192 -Hash MD5 -KeyExchange DH14)

PS C:\>$proposal3 = (New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH19)

PS C:\>$mMCryptoSet= (New-NetIPsecMainModeCryptoSet -DisplayName "Main Mode Crypto Set" -Proposal $proposal1,$proposal2,$proposal3)

This cmdlet shows an alternative method of accomplishing the previous steps.
PS C:\>$mMCryptoSet = New-NetIPsecMainModeCryptoSet -DisplayName "Main Mode Crypto Set" -Proposal (New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH1),(New-NetIPsecMainModeCryptoProposal -Encryption AES192 -Hash MD5 -KeyExchange DH14),(New-NetIPsecMainModeCryptoProposal -Encryption DES3 -Hash MD5 -KeyExchange DH19)

PS C:\>New-NetIPsecMainModeRule -DisplayName "Main Mode Rule" -MainModeCryptoSet $mMCryptoSet.Name
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/new-netipsecmainmodecryptoproposal?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/new-netipsecmainmodecryptoproposal?view=powershell-7.5)

### New-NetIPsecQuickModeCryptoProposal

版本：都有

模块：NetSecurity

语法（5.1）：

```powershell
New-NetIPsecQuickModeCryptoProposal [-Encryption <EncryptionAlgorithm>] [-AHHash <HashAlgorithm>] [-ESPHash <HashAlgorithm>] [-MaxKiloBytes <uint64>] [-MaxMinutes <uint64>] [-Encapsulation <IPsecEncapsulation>] [<CommonParameters>]
```

语法（7）：

```powershell
New-NetIPsecQuickModeCryptoProposal [-Encryption <EncryptionAlgorithm>] [-AHHash <HashAlgorithm>] [-ESPHash <HashAlgorithm>] [-MaxKiloBytes <ulong>] [-MaxMinutes <ulong>] [-Encapsulation <IPsecEncapsulation>] [<CommonParameters>]
```

示例：

```powershell
PS C:\>$QMProposal = New-NetIPsecQuickModeCryptoProposal -Encapsulation ESP -ESPHash SHA1 -Encryption AES128

PS C:\>$QMCryptoSet = New-NetIPsecQuickModeCryptoSet -DisplayName "esp:sha1-des3" -Proposal $QMProposal

PS C:\>New-NetIPSecRule -DisplayName "Tunnel from HQ to Dallas Branch" -Mode Tunnel -LocalAddress 192.168.0.0/16 -RemoteAddress 192.157.0.0/16 -LocalTunnelEndpoint 1.1.1.1 -RemoteTunnelEndpoint 2.2.2.2 -InboundSecurity Require -OutboundSecurity Require -QuickModeCryptoSet $QMCryptoSet.Name
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/new-netipsecquickmodecryptoproposal?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/netsecurity/new-netipsecquickmodecryptoproposal?view=powershell-7.5)

### New-PmemDedicatedMemory

版本：都有

模块：PersistentMemory

语法：

```powershell
New-PmemDedicatedMemory -RegionId <uint32[]> [-FriendlyName <string[]>] [-SizeInBytes <uint64[]>] [<CommonParameters>]
```

示例：创建专用持久内存

```powershell
New-PmemDedicatedMemory -RegionId 1 -SizeInBytes 270582939648
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/new-pmemdedicatedmemory?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/new-pmemdedicatedmemory?view=powershell-7.5)

### New-PmemDisk

版本：都有

模块：PersistentMemory

语法：

```powershell
New-PmemDisk -RegionId <uint32[]> [-FriendlyName <string[]>] [-DiskSizeInBytes <uint64[]>] [-AtomicityType <NAMESPACE_ATOMICITY_TYPE[]>] [<CommonParameters>]
New-PmemDisk -DiskSizeInBytes <uint64[]> -Simulated [-AtomicityType <NAMESPACE_ATOMICITY_TYPE[]>] [<CommonParameters>]
```

示例：创建磁盘

```powershell
New-PmemDisk -RegionId 1 -AtomicityType BlockTranslationTable
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/new-pmemdisk?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/new-pmemdisk?view=powershell-7.5)

### New-ProvisioningRepro

版本：都有

模块：Provisioning

语法：暂无

示例：暂无

出处：[Provisioning 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/provisioning)（没有单独介绍）。

### New-PSWorkflowExecutionOption

版本：都有

模块：PSWorkflow

语法：

```powershell
New-PSWorkflowExecutionOption [-PersistencePath <string>] [-MaxPersistenceStoreSizeGB <long>] [-PersistWithEncryption] [-MaxRunningWorkflows <int>] [-AllowedActivity <string[]>] [-OutOfProcessActivity <string[]>] [-EnableValidation] [-MaxDisconnectedSessions <int>] [-MaxConnectedSessions <int>] [-MaxSessionsPerWorkflow <int>] [-MaxSessionsPerRemoteNode <int>] [-MaxActivityProcesses <int>] [-ActivityProcessIdleTimeoutSec <int>] [-RemoteNodeSessionIdleTimeoutSec <int>] [-SessionThrottleLimit <int>] [-WorkflowShutdownTimeoutMSec <int>] [<CommonParameters>]
```

示例：创建工作流选项对象

```powershell
New-PSWorkflowExecutionOption -MaxSessionsPerWorkflow 10 -MaxDisconnectedSessions 200
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psworkflow/new-psworkflowexecutionoption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psworkflow/new-psworkflowexecutionoption?view=powershell-7.5)

### New-ScheduledJobOption

版本：都有

模块：PSScheduledJob

语法：

```powershell
New-ScheduledJobOption [-RunElevated] [-HideInTaskScheduler] [-RestartOnIdleResume] [-MultipleInstancePolicy <TaskMultipleInstancePolicy>] [-DoNotAllowDemandStart] [-RequireNetwork] [-StopIfGoingOffIdle] [-WakeToRun] [-ContinueIfGoingOnBattery] [-StartIfOnBattery] [-IdleTimeout <timespan>] [-IdleDuration <timespan>] [-StartIfIdle] [<CommonParameters>]
```

示例：创建具有默认值的计划作业选项对象

```powershell
New-ScheduledJobOption
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/new-scheduledjoboption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/new-scheduledjoboption?view=powershell-7.5)

### New-SelfSignedCertificate

版本：都有

模块：PKI

语法：

```powershell
New-SelfSignedCertificate [-SecurityDescriptor <FileSecurity>] [-TextExtension <string[]>] [-Extension <X509Extension[]>] [-HardwareKeyUsage <HardwareKeyUsage[]>] [-KeyUsageProperty <KeyUsageProperty[]>] [-KeyUsage <KeyUsage[]>] [-KeyProtection <KeyProtection[]>] [-KeyExportPolicy <KeyExportPolicy[]>] [-KeyLength <int>] [-KeyAlgorithm <string>] [-SmimeCapabilities] [-ExistingKey] [-KeyLocation <string>] [-SignerReader <string>] [-Reader <string>] [-SignerPin <securestring>] [-Pin <securestring>] [-KeyDescription <string>] [-KeyFriendlyName <string>] [-Container <string>] [-Provider <string>] [-CurveExport <CurveParametersExportType>] [-KeySpec <KeySpec>] [-Type <CertificateType>] [-FriendlyName <string>] [-NotAfter <datetime>] [-NotBefore <datetime>] [-SerialNumber <string>] [-Subject <string>] [-DnsName <string[]>] [-SuppressOid <string[]>] [-HashAlgorithm <string>] [-AlternateSignatureAlgorithm] [-TestRoot] [-Signer <Certificate>] [-CloneCert <Certificate>] [-CertStoreLocation <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 DnsName = 'www.fabrikam.com', 'www.contoso.com'
 CertStoreLocation = 'Cert:\LocalMachine\My'
}
New-SelfSignedCertificate @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/new-selfsignedcertificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/new-selfsignedcertificate?view=powershell-7.5)

### New-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法（5.1）：

```powershell
New-Service [-Name] <string> [-BinaryPathName] <string> [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartMode>] [-Credential <pscredential>] [-DependsOn <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
New-Service [-Name] <string> [-BinaryPathName] <string> [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartupType>] [-Credential <pscredential>] [-SecurityDescriptorSddl <string>] [-DependsOn <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建服务

```powershell
New-Service -Name "TestService" -BinaryPathName 'C:\WINDOWS\System32\svchost.exe -k netsvcs'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/new-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/new-service?view=powershell-7.5)

### New-TlsSessionTicketKey

版本：都有

模块：TLS

语法：

```powershell
New-TlsSessionTicketKey [-Password] <securestring> [[-Path] <string>] [<CommonParameters>]
```

示例：创建 TLS 会话票证密钥

```powershell
$Password = Read-Host -AsSecureString
New-TlsSessionTicketKey -Password $Password -Path 'C:\KeyConfig\TlsSessionTicketKey.config'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/tls/new-tlssessionticketkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/tls/new-tlssessionticketkey?view=powershell-7.5)

### New-WebServiceProxy

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
New-WebServiceProxy [-Uri] <uri> [[-Class] <string>] [[-Namespace] <string>] [<CommonParameters>]
New-WebServiceProxy [-Uri] <uri> [[-Class] <string>] [[-Namespace] <string>] [-Credential <pscredential>] [<CommonParameters>]
New-WebServiceProxy [-Uri] <uri> [[-Class] <string>] [[-Namespace] <string>] [-UseDefaultCredential] [<CommonParameters>]
```

示例：为 Web 服务创建代理

```powershell
$calc = New-WebServiceProxy -Uri "http://www.dneonline.com/calculator.asmx?wsdl"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/new-webserviceproxy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/new-webserviceproxy?view=powershell-7.5)

### New-WindowsCustomImage

版本：都有

模块：Dism

语法：

```powershell
New-WindowsCustomImage -CapturePath <string> [-ConfigFilePath <string>] [-CheckIntegrity] [-Verify] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：捕获映像定制文件

```powershell
PS C:\> New-WindowsCustomImage -CapturePath "c:\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/new-windowscustomimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/new-windowscustomimage?view=powershell-7.5)

### New-WindowsImage

版本：都有

模块：Dism

语法：

```powershell
New-WindowsImage -ImagePath <string> -CapturePath <string> [-CompressionType <string>] [-ConfigFilePath <string>] [-Description <string>] [-Name <string>] [-CheckIntegrity] [-NoRpFix] [-Setbootable] [-Verify] [-WIMBoot] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：为 WIM 文件捕获驱动器映像

```powershell
PS C:\> New-WindowsImage -ImagePath "c:\imagestore\custom.wim" -CapturePath "d:\" -Name "Drive D"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/new-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/new-windowsimage?view=powershell-7.5)

### New-WinEvent

版本：都有

模块：Microsoft.PowerShell.Diagnostics

语法：

```powershell
New-WinEvent [-ProviderName] <string> [-Id] <int> [[-Payload] <Object[]>] [-Version <byte>] [<CommonParameters>]
```

示例：创建新事件

```powershell
New-WinEvent -ProviderName Microsoft-Windows-PowerShell -Id 45090 -Payload @("Workflow", "Running")
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/new-winevent?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.diagnostics/new-winevent?view=powershell-7.5)

### New-WinUserLanguageList

版本：都有

模块：International

语法：

```powershell
New-WinUserLanguageList [-Language] <string> [<CommonParameters>]
```

示例：创建并设置语言列表

```powershell
PS C:\> $UserLanguageList = New-WinUserLanguageList -Language "en-US"
PS C:\> $UserLanguageList.Add("fr-FR")
PS C:\> Set-WinUserLanguageList -LanguageList $UserLanguageList
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/new-winuserlanguagelist?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/new-winuserlanguagelist?view=powershell-7.5)

### New-WSManInstance

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
New-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ApplicationName <string>] [-ComputerName <string>] [-FilePath <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
New-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ConnectionURI <uri>] [-FilePath <string>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：创建 HTTPS 侦听器

```powershell
New-WSManInstance winrm/config/Listener -SelectorSet @{Transport='HTTPS'; Address='*'} -ValueSet @{Hostname="HOST";CertificateThumbprint="XXXXXXXXXX"}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/new-wsmaninstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/new-wsmaninstance?view=powershell-7.5)

### New-WSManSessionOption

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
New-WSManSessionOption [-ProxyAccessType <ProxyAccessType>] [-ProxyAuthentication <ProxyAuthentication>] [-ProxyCredential <pscredential>] [-SkipCACheck] [-SkipCNCheck] [-SkipRevocationCheck] [-SPNPort <int>] [-OperationTimeout <int>] [-NoEncryption] [-UseUTF16] [<CommonParameters>]
```

示例：创建使用连接选项的连接

```powershell
PS C:\> $a = New-WSManSessionOption -OperationTimeout 30000
PS C:\> Connect-WSMan -ComputerName "server01" -SessionOption $a
PS C:\> cd WSMan:
PS WSMan:\> dir
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/new-wsmansessionoption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/new-wsmansessionoption?view=powershell-7.5)

### Optimize-AppxProvisionedPackages

版本：都有

模块：Dism

语法：

```powershell
Optimize-AppXProvisionedPackages -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>] Optimize-AppXProvisionedPackages -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：优化脱机 Windows 映像中的预配包

```powershell
PS> Optimize-AppXProvisionedPackages -Path ".\wim\image.wim"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/optimize-appxprovisionedpackages?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/optimize-appxprovisionedpackages?view=powershell-7.5)

### Optimize-WindowsImage

版本：都有

模块：Dism

语法：

```powershell
Optimize-WindowsImage -OptimizationTarget <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：将映像配置为 WIMBoot 系统

```powershell
PS C:\> Optimize-WindowsImage -Path "c:\" -OptimizationTarget "WIMBoot"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/optimize-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/optimize-windowsimage?view=powershell-7.5)

### Out-GridView

版本：都有

模块：Microsoft.PowerShell.Utility

语法：

```powershell
Out-GridView [-InputObject <psobject>] [-Title <string>] [-PassThru] [<CommonParameters>]
Out-GridView [-InputObject <psobject>] [-Title <string>] [-Wait] [<CommonParameters>]
Out-GridView [-InputObject <psobject>] [-Title <string>] [-OutputMode <OutputModeOption>] [<CommonParameters>]
```

示例：将进程输出到网格视图

```powershell
Get-Process | Out-GridView
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/out-gridview?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/out-gridview?view=powershell-7.5)

### Out-Printer

版本：都有

模块：Microsoft.PowerShell.Utility

语法：

```powershell
Out-Printer [[-Name] <string>] [-InputObject <psobject>] [<CommonParameters>]
```

示例：发送要打印在默认打印机上的文件

```powershell
Get-Content -Path ./readme.txt | Out-Printer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/out-printer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/out-printer?view=powershell-7.5)

### Publish-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Publish-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [[-DynamicUserConfigurationPath] <string>] [-Global] [-UserSID <string>] [-DynamicUserConfigurationType <DynamicUserConfiguration>] [<CommonParameters>]
Publish-AppvClientPackage [-Package] <AppvClientPackage> [[-DynamicUserConfigurationPath] <string>] [-Global] [-UserSID <string>] [-DynamicUserConfigurationType <DynamicUserConfiguration>] [<CommonParameters>]
Publish-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [-UserSID <string>] [-DynamicUserConfigurationPath <string>] [-DynamicUserConfigurationType <DynamicUserConfiguration>] [<CommonParameters>]
```

示例：向全部用户发布包版本

```powershell
PS C:\> Publish-AppvClientPackage -Name "MyApp" -Version 1 -Global -DynamicUserConfiguration "C:\content\policies\MyApp.policy"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/publish-appvclientpackage?view=powershell-5.1)

### Publish-DscConfiguration

版本：都有

模块：PSDesiredStateConfiguration

语法：

```powershell
Publish-DscConfiguration [-Path] <string> [[-ComputerName] <string[]>] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Publish-DscConfiguration [-Path] <string> -CimSession <CimSession[]> [-Force] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将配置发布到远程计算机

```powershell
Publish-DscConfiguration -Path '$home\WebServer' -ComputerName "ContosoWebServer" -Credential (get-credential Contoso\webadministrator)
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/publish-dscconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/publish-dscconfiguration?view=powershell-7.5)

### Receive-DtcDiagnosticTransaction

版本：都有

模块：MsDtc

语法：

```powershell
Receive-DtcDiagnosticTransaction [[-ComputerName] <string>] [[-Port] <int>] [[-PropagationMethod] <DtcTransactionPropagation>] [<CommonParameters>]
```

示例：接收诊断事务

```powershell
PS C:\> Receive-DtcDiagnosticTransaction -ComputerName "Host1" -Port 17123 -PropagationMethod Pull
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/receive-dtcdiagnostictransaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/receive-dtcdiagnostictransaction?view=powershell-7.5)

### Receive-PSSession

版本：都有

模块：Microsoft.PowerShell.Core

语法（5.1）：

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

语法（7）：

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

示例：连接到 PSSession

```powershell
Receive-PSSession -ComputerName Server01 -Name ITTask
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/receive-pssession?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/receive-pssession?view=powershell-7.5)

### Register-CimIndicationEvent

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint32>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint32>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

语法（7）：

```powershell
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-ClassName] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-OperationTimeoutSec <uint>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] -CimSession <CimSession> [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-CimIndicationEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-QueryDialect <string>] [-OperationTimeoutSec <uint>] [-ComputerName <string>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

示例：注册类生成的事件

```powershell
$event = @{
 ClassName = 'Win32_ProcessStartTrace'
 SourceIdentifier = 'ProcessStarted'
}
Register-CimIndicationEvent @event
Get-Event -SourceIdentifier "ProcessStarted"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/register-cimindicationevent?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/register-cimindicationevent?view=powershell-7.5)

### Register-PSSessionConfiguration

版本：都有

模块：Microsoft.PowerShell.Core

语法（5.1）：

```powershell
Register-PSSessionConfiguration [-Name] <string> [-ProcessorArchitecture <string>] [-SessionType <PSSessionType>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> [-AssemblyName] <string> [-ConfigurationTypeName] <string> [-ProcessorArchitecture <string>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> -Path <string> [-ProcessorArchitecture <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-TransportOption <PSTransportOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Register-PSSessionConfiguration [-Name] <string> [-ProcessorArchitecture <string>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> [-AssemblyName] <string> [-ConfigurationTypeName] <string> [-ProcessorArchitecture <string>] [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-PSSessionConfiguration [-Name] <string> -Path <string> [-ProcessorArchitecture <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-TransportOption <PSTransportOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：注册 NewShell 会话配置

```powershell
$sessionConfiguration = @{
 Name='NewShell'
 ApplicationBase='C:\MyShells\'
 AssemblyName='MyShell.dll'
 ConfigurationTypeName='MyClass'
}
Register-PSSessionConfiguration @sessionConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/register-pssessionconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/register-pssessionconfiguration?view=powershell-7.5)

### Register-RecoveryManagementPlugin

版本：都有

模块：Dism

语法（5.1）：

```powershell
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint32> -Path <string> [-CapabilitiesDesired <uint32>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint32> -Online [-CapabilitiesDesired <uint32>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint> -Path <string> [-CapabilitiesDesired <uint>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Register-RecoveryManagementPlugin -BinaryLocation <string> -ClassID <string> -CapabilitiesRequired <uint> -Online [-CapabilitiesDesired <uint>] [-ThreadingModel <string>] [-ExceptionHandling <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Register-ScheduledJob

版本：都有

模块：PSScheduledJob

语法：

```powershell
Register-ScheduledJob [-Name] <string> [-ScriptBlock] <scriptblock> [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-ArgumentList <Object[]>] [-MaxResultCount <int>] [-RunNow] [-RunEvery <timespan>] [-WhatIf] [-Confirm] [<CommonParameters>]
Register-ScheduledJob [-Name] <string> [-FilePath] <string> [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-ArgumentList <Object[]>] [-MaxResultCount <int>] [-RunNow] [-RunEvery <timespan>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建计划作业

```powershell
Register-ScheduledJob -Name "Archive-Scripts" -ScriptBlock {
 Get-ChildItem $HOME\*.ps1 -Recurse |
 Copy-Item -Destination "\\Server\Share\PSScriptArchive"
}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/register-scheduledjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/register-scheduledjob?view=powershell-7.5)

### Register-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Register-UevTemplate [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Register-UevTemplate -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：注册模板

```powershell
PS C:\> Register-UevTemplate -Path "MicrosoftCalculator.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/register-uevtemplate?view=powershell-5.1)

### Register-WmiEvent

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Register-WmiEvent [-Class] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-Credential <pscredential>] [-ComputerName <string>] [-Timeout <long>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
Register-WmiEvent [-Query] <string> [[-SourceIdentifier] <string>] [[-Action] <scriptblock>] [-Namespace <string>] [-Credential <pscredential>] [-ComputerName <string>] [-Timeout <long>] [-MessageData <psobject>] [-SupportEvent] [-Forward] [-MaxTriggerCount <int>] [<CommonParameters>]
```

示例：订阅类生成的事件

```powershell
Register-WmiEvent -Class 'Win32_ProcessStartTrace' -SourceIdentifier "ProcessStarted"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/register-wmievent?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/register-wmievent?view=powershell-7.5)

### Remove-AppProvisionedSharedPackageContainer

版本：都有

模块：Dism

语法：

```powershell
Remove-AppProvisionedSharedPackageContainer -Name <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-AppProvisionedSharedPackageContainer -Name <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Remove-AppSharedPackageContainer

版本：都有

模块：Appx

语法：

```powershell
Remove-AppSharedPackageContainer [-Name] <string> [-ForceApplicationShutdown] [<CommonParameters>]
```

示例：

```powershell
Remove-AppSharedPackageContainer -Name ContosoTestContainer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appsharedpackagecontainer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appsharedpackagecontainer?view=powershell-7.5)

### Remove-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Remove-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [<CommonParameters>]
Remove-AppvClientConnectionGroup [-Name] <string> [<CommonParameters>]
Remove-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [<CommonParameters>]
```

示例：删除指定连接组

```powershell
PS C:\> Remove-AppvClientConnectionGroup -Name "MyGroup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/remove-appvclientconnectiongroup?view=powershell-5.1)

### Remove-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Remove-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [<CommonParameters>]
Remove-AppvClientPackage [-Package] <AppvClientPackage> [<CommonParameters>]
Remove-AppvClientPackage [-Name] <string> [[-Version] <string>] [<CommonParameters>]
```

示例：用管道符删除包版本

```powershell
PS C:\> Get-AppvPackage -Name "MyPackage" -Version 1 | Remove-Package
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/remove-appvclientpackage?view=powershell-5.1)

### Remove-AppvPublishingServer

版本：仅5.1

模块：AppvClient

语法：

```powershell
Remove-AppvPublishingServer [-ServerId] <uint32> [<CommonParameters>]
Remove-AppvPublishingServer [-Server] <AppvPublishingServer> [<CommonParameters>]
Remove-AppvPublishingServer [[-Name] <string>] [[-URL] <string>] [<CommonParameters>]
```

示例：删除发布服务器

```powershell
PS C:\> Remove-AppvPublishingServer -Name "Server01"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/remove-appvpublishingserver?view=powershell-5.1)

### Remove-AppxPackage

版本：都有

模块：Appx

语法：

```powershell
Remove-AppxPackage [-Package] <string> [-PreserveApplicationData] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-AppxPackage [-Package] <string> [-PreserveRoamableApplicationData] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-AppxPackage [-Package] <string> [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-AppxPackage [-Package] <string> -User <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除应用包

```powershell
Remove-AppxPackage -Package 'package1_1.0.0.0_neutral__8wekyb3d8bbwe'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appxpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appxpackage?view=powershell-7.5)

### Remove-AppxPackageAutoUpdateSettings

版本：都有

模块：Appx

语法：

```powershell
Remove-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> [-UseSystemPolicySource] [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Remove-AppxPackageAutoUpdateSettings -PackageFullName publisher.package1_1.0.0.0_neutral__8wekyb3d8bbwe
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appxpackageautoupdatesettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appxpackageautoupdatesettings?view=powershell-7.5)

### Remove-AppxProvisionedPackage

版本：都有

模块：Dism

语法：

```powershell
Remove-AppxProvisionedPackage -PackageName <string> -Online [-AllUsers] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-AppxProvisionedPackage -PackageName <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：从映像删除应用包

```powershell
PS C:\> Remove-AppxProvisionedPackage -Path c:\offline -PackageName MyAppxPkg
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-appxprovisionedpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-appxprovisionedpackage?view=powershell-7.5)

### Remove-AppxVolume

版本：都有

模块：Appx

语法：

```powershell
Remove-AppxVolume [-Volume] <AppxVolume[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按 ID 删除卷

```powershell
Remove-AppxVolume -Volume {984786d3-0cae-49de-a68f-8bedb0ca260b}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appxvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/remove-appxvolume?view=powershell-7.5)

### Remove-BcdElement

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Remove-BcdElement [-Element] <string> [[-Id] <string>] [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-BcdElement [-Element] <string> [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/remove-bcdelement?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/remove-bcdelement?view=powershell-7.5)

### Remove-BcdEntry

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Remove-BcdEntry [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-BcdEntry [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/remove-bcdentry?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/remove-bcdentry?view=powershell-7.5)

### Remove-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Remove-BitsTransfer [-BitsJob] <BitsJob[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：取消当前用户拥有的全部 BITS 传输作业

```powershell
PS C:\> Get-BitsTransfer | Remove-BitsTransfer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/remove-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/remove-bitstransfer?view=powershell-7.5)

### Remove-CertificateEnrollmentPolicyServer

版本：都有

模块：PKI

语法：

```powershell
Remove-CertificateEnrollmentPolicyServer [-Url] <uri> -context <Context> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 Url = 'https://www.contoso.com/policy/service.svc'
 Context = 'User'
}
Remove-CertificateEnrollmentPolicyServer @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/remove-certificateenrollmentpolicyserver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/remove-certificateenrollmentpolicyserver?view=powershell-7.5)

### Remove-CertificateNotificationTask

版本：都有

模块：PKI

语法：

```powershell
Remove-CertificateNotificationTask [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Remove-CertificateNotificationTask -Name "My Task"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/remove-certificatenotificationtask?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/remove-certificatenotificationtask?view=powershell-7.5)

### Remove-CimInstance

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Remove-CimInstance [-InputObject] <ciminstance> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Remove-CimInstance [-InputObject] <ciminstance> [-ResourceUri <uri>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] -CimSession <CimSession[]> [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimInstance [-Query] <string> [[-Namespace] <string>] [-ComputerName <string[]>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除 CIM 实例

```powershell
Remove-CimInstance -Query 'Select * from Win32_Environment where name LIKE "testvar%"'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/remove-ciminstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/remove-ciminstance?view=powershell-7.5)

### Remove-CimSession

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Remove-CimSession [-CimSession] <CimSession[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-ComputerName] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-Id] <uint32[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -InstanceId <guid[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -Name <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Remove-CimSession [-CimSession] <CimSession[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-ComputerName] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession [-Id] <uint[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -InstanceId <guid[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-CimSession -Name <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除所有 CIM 会话

```powershell
Get-CimSession | Remove-CimSession
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/remove-cimsession?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/remove-cimsession?view=powershell-7.5)

### Remove-CIPolicyRule

版本：都有

模块：ConfigCI

语法：

```powershell
Remove-CIPolicyRule [-Id] <string> -FilePath <string> [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/remove-cipolicyrule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/remove-cipolicyrule?view=powershell-7.5)

### Remove-Computer

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Remove-Computer [[-UnjoinDomainCredential] <pscredential>] [-Restart] [-Force] [-PassThru] [-WorkgroupName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Computer -UnjoinDomainCredential <pscredential> [-LocalCredential <pscredential>] [-Restart] [-ComputerName <string[]>] [-Force] [-PassThru] [-WorkgroupName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：从其域中删除本地计算机

```powershell
Remove-Computer -UnjoinDomaincredential Domain01\Admin01 -PassThru -Verbose -Restart
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-computer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-computer?view=powershell-7.5)

### Remove-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Remove-EventLog [-LogName] <string[]> [[-ComputerName] <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-EventLog [[-ComputerName] <string[]>] [-Source <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：从本地计算机中删除事件日志

```powershell
Remove-EventLog -LogName "MyLog"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-eventlog?view=powershell-7.5)

### Remove-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
Remove-JobTrigger [-InputObject] <ScheduledJobDefinition[]> [-TriggerId <int[]>] [<CommonParameters>]
Remove-JobTrigger [-Id] <int[]> [-TriggerId <int[]>] [<CommonParameters>]
Remove-JobTrigger [-Name] <string[]> [-TriggerId <int[]>] [<CommonParameters>]
```

示例：删除所有作业触发器

```powershell
Remove-JobTrigger -Name "Test*"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/remove-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/remove-jobtrigger?view=powershell-7.5)

### Remove-LocalGroup

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Remove-LocalGroup [-InputObject] <LocalGroup[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroup [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroup [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除安全组

```powershell
Remove-LocalGroup -Name "SecurityGroup04"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/remove-localgroup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/remove-localgroup?view=powershell-7.5)

### Remove-LocalGroupMember

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Remove-LocalGroupMember [-Group] <LocalGroup> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroupMember [-Name] <string> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalGroupMember [-SID] <SecurityIdentifier> [-Member] <LocalPrincipal[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：从管理员组中删除成员

```powershell
$members = "Admin02", "MicrosoftAccount\username@Outlook.com", "AzureAD\DavidChew@contoso.com", "CONTOSO\Domain Admins"
Remove-LocalGroupMember -Group "Administrators" -Member $members
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/remove-localgroupmember?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/remove-localgroupmember?view=powershell-7.5)

### Remove-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Remove-LocalUser [-InputObject] <LocalUser[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalUser [-Name] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-LocalUser [-SID] <SecurityIdentifier[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除用户帐户

```powershell
Remove-LocalUser -Name "AdminContoso02"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/remove-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/remove-localuser?view=powershell-7.5)

### Remove-OsConfigurationDocument

版本：都有

模块：OsConfiguration

语法：

```powershell
Remove-OsConfigurationDocument [-Id] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [-Wait] [[-TimeoutInSeconds] <int>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/remove-osconfigurationdocument?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/remove-osconfigurationdocument?view=powershell-7.5)

### Remove-OSConfigurationScenarioDefinition

版本：都有

模块：OsConfiguration

语法：

```powershell
Remove-OsConfigurationScenarioDefinition [-Name] <string> [-Version] <string> [-SchemaVersion] <string> [<CommonParameters>]
```

示例：暂无

出处：[OsConfiguration 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration)（没有单独介绍）。

### Remove-PmemDedicatedMemory

版本：都有

模块：PersistentMemory

语法：

```powershell
Remove-PmemDedicatedMemory -DeviceNumber <uint32> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除专用持久内存

```powershell
Remove-PmemDedicatedMemory -DeviceNumber 1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/remove-pmemdedicatedmemory?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/remove-pmemdedicatedmemory?view=powershell-7.5)

### Remove-PmemDisk

版本：都有

模块：PersistentMemory

语法：

```powershell
Remove-PmemDisk -DiskNumber <uint32> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-PmemDisk -Simulated [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除持久内存磁盘

```powershell
Remove-PmemDisk -DiskNumber 2
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/remove-pmemdisk?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/persistentmemory/remove-pmemdisk?view=powershell-7.5)

### Remove-PSSnapin

版本：仅5.1

模块：Microsoft.PowerShell.Core

语法：

```powershell
Remove-PSSnapin [-Name] <string[]> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除管理单元

```powershell
Remove-PSSnapin -Name Microsoft.Exchange
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/remove-pssnapin?view=powershell-5.1)

### Remove-RecoveryManagementPluginAltitude

版本：都有

模块：Dism

语法：

```powershell
Remove-RecoveryManagementPluginAltitude -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-RecoveryManagementPluginAltitude -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Remove-Service

版本：仅7

模块：Microsoft.PowerShell.Management

语法：

```powershell
Remove-Service [-Name] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-Service [-InputObject <ServiceController>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除服务

```powershell
Remove-Service -Name "TestService"
```

出处：[官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-service?view=powershell-7.5)

### Remove-WindowsCapability

版本：都有

模块：Dism

语法：

```powershell
Remove-WindowsCapability -Name <string> -Online [-DelayExecutionIfPending] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsCapability -Name <string> -Path <string> [-DelayExecutionIfPending] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：删除映像的 Windows 功能

```powershell
PS C:\> Remove-WindowsCapability -Name "Language.TextToSpeech~~~fr-FR~0.0.1.0" -Path "C:\offline"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowscapability?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowscapability?view=powershell-7.5)

### Remove-WindowsDriver

版本：都有

模块：Dism

语法：

```powershell
Remove-WindowsDriver -Driver <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：从映像删除驱动程序

```powershell
PS C:\> Remove-WindowsDriver -Path "c:\offline" -Driver "OEM1.inf"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowsdriver?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowsdriver?view=powershell-7.5)

### Remove-WindowsImage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Remove-WindowsImage -ImagePath <string> -Name <string> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsImage -ImagePath <string> -Index <uint32> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Remove-WindowsImage -ImagePath <string> -Name <string> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsImage -ImagePath <string> -Index <uint> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：从 WIM 文件删除第一个映像

```powershell
PS C:\> Remove-WindowsImage -ImagePath "c:\imagestore\custom.wim" -Index 1 -CheckIntegrity
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowsimage?view=powershell-7.5)

### Remove-WindowsPackage

版本：都有

模块：Dism

语法：

```powershell
Remove-WindowsPackage -Path <string> [-PackagePath <string>] [-PackageName <string>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Remove-WindowsPackage -Online [-PackagePath <string>] [-PackageName <string>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：从正在运行的操作系统映像删除包

```powershell
PS C:\> Remove-WindowsPackage -Online -PackageName "Microsoft-Windows-Backup-Package~31bf3856ad364e35~x86~~6.1.7601.16525"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowspackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/remove-windowspackage?view=powershell-7.5)

### Remove-WmiObject

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Remove-WmiObject [-Class] <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject -InputObject <wmi> [-AsJob] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject -Path <string> [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Remove-WmiObject [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：关闭 Win32 进程的所有实例

```powershell
notepad
$np = Get-WmiObject -Query "select * from Win32_Process where name='notepad.exe'"
$np | Remove-WmiObject
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-wmiobject?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/remove-wmiobject?view=powershell-7.5)

### Remove-WSManInstance

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Remove-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ApplicationName <string>] [-ComputerName <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Remove-WSManInstance [-ResourceURI] <uri> [-SelectorSet] <hashtable> [-ConnectionURI <uri>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：删除侦听器

```powershell
Remove-WSManInstance -ResourceUri winrm/config/Listener -SelectorSet @{
 Address = 'test.fabrikam.com'
 Transport = 'http'
}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/remove-wsmaninstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/remove-wsmaninstance?view=powershell-7.5)

### Rename-Computer

版本：都有

模块：Microsoft.PowerShell.Management

语法（5.1）：

```powershell
Rename-Computer [-NewName] <string> [-ComputerName <string>] [-PassThru] [-DomainCredential <pscredential>] [-LocalCredential <pscredential>] [-Force] [-Restart] [-WsmanAuthentication <string>] [-Protocol <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Rename-Computer [-NewName] <string> [-ComputerName <string>] [-PassThru] [-DomainCredential <pscredential>] [-LocalCredential <pscredential>] [-Force] [-Restart] [-WsmanAuthentication <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：重命名本地计算机

```powershell
Rename-Computer -NewName "Server044" -DomainCredential Domain01\Admin01 -Restart
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/rename-computer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/rename-computer?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`sudo hostnamectl set-hostname`）。
- 发行版：需要 sudo。
- 同组：Restart-Computer、Stop-Computer。
- 功能：改主机名。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-NewName`（位置 0） | string | 新主机名；传给 `sudo hostnamectl set-hostname` |

### Rename-LocalGroup

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Rename-LocalGroup [-InputObject] <LocalGroup> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalGroup [-Name] <string> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalGroup [-SID] <SecurityIdentifier> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：更改组的名称

```powershell
PS C:\> Rename-LocalGroup -Name "SecurityGroup" -NewName "SecurityGroup04"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/rename-localgroup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/rename-localgroup?view=powershell-7.5)

### Rename-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Rename-LocalUser [-InputObject] <LocalUser> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalUser [-Name] <string> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Rename-LocalUser [-SID] <SecurityIdentifier> [-NewName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：重命名用户帐户

```powershell
Rename-LocalUser -Name "Admin02" -NewName "AdminContoso02"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/rename-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/rename-localuser?view=powershell-7.5)

### Repair-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Repair-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientConnectionGroup [-Name] <string> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
```

示例：修复指定连接组

```powershell
PS C:\> Repair-AppvClientConnectionGroup -Name "MyGroup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/repair-appvclientconnectiongroup?view=powershell-5.1)

### Repair-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Repair-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientPackage [-Package] <AppvClientPackage> [-Global] [-UserState] [-Extensions] [<CommonParameters>]
Repair-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [-UserState] [-Extensions] [<CommonParameters>]
```

示例：删除包某版本的用户状态

```powershell
PS C:\> Repair-AppvClientPackage -Name "MyApp" -Version 3
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/repair-appvclientpackage?view=powershell-5.1)

### Repair-UevTemplateIndex

版本：仅5.1

模块：UEV

语法：

```powershell
Repair-UevTemplateIndex [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：修复模板索引

```powershell
PS C:\> Repair-UevTemplateIndex
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/repair-uevtemplateindex?view=powershell-5.1)

### Repair-WindowsImage

版本：都有

模块：Dism

语法：

```powershell
Repair-WindowsImage -Path <string> [-CheckHealth] [-ScanHealth] [-RestoreHealth] [-StartComponentCleanup] [-LimitAccess] [-ResetBase] [-Defer] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Repair-WindowsImage -Online [-CheckHealth] [-ScanHealth] [-RestoreHealth] [-StartComponentCleanup] [-LimitAccess] [-ResetBase] [-Defer] [-Source <string[]>] [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：扫描映像损坏

```powershell
PS C:\> Repair-WindowsImage -Path "C:\offline\Mount" -ScanHealth
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/repair-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/repair-windowsimage?view=powershell-7.5)

### Reset-AppSharedPackageContainer

版本：都有

模块：Appx

语法：

```powershell
Reset-AppSharedPackageContainer [-Name] <string> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Reset-AppSharedPackageContainer -Name ContosoTestContainer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/reset-appsharedpackagecontainer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/reset-appsharedpackagecontainer?view=powershell-7.5)

### Reset-AppxPackage

版本：都有

模块：Appx

语法：

```powershell
Reset-AppxPackage [-Package] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：重置应用包

```powershell
Reset-AppxPackage -Package publisher.package1_1.0.0.0_neutral__8wekyb3d8bbwe
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/reset-appxpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/reset-appxpackage?view=powershell-7.5)

### Reset-ComputerMachinePassword

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Reset-ComputerMachinePassword [-Server <string>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：重置本地计算机的密码

```powershell
Reset-ComputerMachinePassword
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/reset-computermachinepassword?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/reset-computermachinepassword?view=powershell-7.5)

### Reset-LapsPassword

版本：都有

模块：LAPS

语法：

```powershell
Reset-LapsPassword [<CommonParameters>]
```

示例：

```powershell
Reset-LapsPassword
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/reset-lapspassword?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/reset-lapspassword?view=powershell-7.5)

### Resolve-DnsName

版本：都有

模块：DnsClient

语法：

```powershell
Resolve-DnsName [-Name] <string> [[-Type] <RecordType>] [-Server <string[]>] [-DohServer <string[][]>] [-DotServer <string[][]>] [-DnsOnly] [-CacheOnly] [-DnssecOk] [-DnssecCd] [-NoHostsFile] [-LlmnrNetbiosOnly] [-LlmnrFallback] [-LlmnrOnly] [-NetbiosFallback] [-NoIdn] [-NoRecursion] [-QuickTimeout] [-TcpOnly] [-CheckCache] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Resolve-DnsName -Name www.bing.com
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dnsclient/resolve-dnsname?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dnsclient/resolve-dnsname?view=powershell-7.5)

### Restart-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Restart-Service [-InputObject] <ServiceController[]> [-Force] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Restart-Service [-Name] <string[]> [-Force] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Restart-Service -DisplayName <string[]> [-Force] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：在本地计算机上重启服务

```powershell
PS C:\> Restart-Service -Name winmgmt
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/restart-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/restart-service?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`systemctl` + sudo）。
- 发行版：systemd 系 + sudo。
- 同组：Stop-Service、Restart-Service、Resume-Service。
- 功能：启动/停止/重启服务。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 服务名；传给 `systemctl <动作> <单元>`（自动补 .service 后缀） |

- 实现：`systemctl start/stop/restart <单元>`；普通权限失败自动用 `sudo` 重试。对应 `sudo systemctl start/stop/restart`。
- Stop-Service、Restart-Service、Resume-Service 参数同 Start-Service，动作为 stop/restart/start。

### Restore-Computer

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Restore-Computer [-RestorePoint] <int> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：还原本地计算机

```powershell
Restore-Computer -RestorePoint 253
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/restore-computer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/restore-computer?view=powershell-7.5)

### Restore-UevBackup

版本：仅5.1

模块：UEV

语法：

```powershell
Restore-UevBackup [-ComputerName] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：从另一台计算机还原备份设置

```powershell
PS C:\>Restore-UevBackup -ComputerName "PattiFullerDevice03@Contoso.Com"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/restore-uevbackup?view=powershell-5.1)

### Restore-UevUserSetting

版本：仅5.1

模块：UEV

语法：

```powershell
Restore-UevUserSetting -Application <string> [-Force] [-LastKnownGood] [-WhatIf] [-Confirm] [<CommonParameters>]
Restore-UevUserSetting [-TemplateId] <string> [-LastKnownGood] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：还原指定模板的用户设置

```powershell
PS C:\> Restore-UevUserSetting -TemplateId "MicrosoftCalculator6"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/restore-uevusersetting?view=powershell-5.1)

### Resume-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Resume-BitsTransfer [-BitsJob] <BitsJob[]> [-Asynchronous] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：恢复当前用户拥有的全部 BITS 传输作业

```powershell
PS C:\> Get-BitsTransfer | Resume-BitsTransfer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/resume-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/resume-bitstransfer?view=powershell-7.5)

### Resume-Job

版本：仅5.1

模块：Microsoft.PowerShell.Core

语法：

```powershell
Resume-Job [-Id] <int[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-Job] <Job[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-Name] <string[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-InstanceId] <guid[]> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-State] <JobState> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Job [-Filter] <hashtable> [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按 ID 恢复作业

```powershell
PS C:\> Get-Job EventJob
Id Name PSJobTypeName State HasMoreData Location Command
-- ---- ------------- ----- ----------- -------- -------
4 EventJob PSWorkflowJob Suspended True Server01 \\Script\Share\Event.ps1

PS C:\> Resume-Job -Id 4
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/resume-job?view=powershell-5.1)

### Resume-ProvisioningSession

版本：都有

模块：Provisioning

语法：暂无

示例：暂无

出处：[Provisioning 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/provisioning)（没有单独介绍）。

### Resume-ReFSDedupSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Resume-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Resume-ReFSDedupSchedule -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/resume-refsdedupschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/resume-refsdedupschedule?view=powershell-7.5)

### Resume-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Resume-Service [-InputObject] <ServiceController[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Service [-Name] <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Resume-Service -DisplayName <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：恢复本地计算机上的服务

```powershell
PS C:\> Resume-Service "sens"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/resume-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/resume-service?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`systemctl` + sudo）。
- 发行版：systemd 系 + sudo。
- 同组：Stop-Service、Restart-Service、Resume-Service。
- 功能：启动/停止/重启服务。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 服务名；传给 `systemctl <动作> <单元>`（自动补 .service 后缀） |

- 实现：`systemctl start/stop/restart <单元>`；普通权限失败自动用 `sudo` 重试。对应 `sudo systemctl start/stop/restart`。
- Stop-Service、Restart-Service、Resume-Service 参数同 Start-Service，动作为 stop/restart/start。

### Save-OsImage

版本：都有

模块：Dism

语法：

```powershell
Save-OsImage -ImagePath <string> -CapturePath <string> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Save-SoftwareInventory

版本：都有

模块：Dism

语法：

```powershell
Save-SoftwareInventory -PartitioningScript <string> -ResetConfigXml <string> -Path <string> [-DevicesInf <string>] [-OutputDirectory <string>] [-CSRConfigFile <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Save-SoftwareInventory -PartitioningScript <string> -ResetConfigXml <string> -Online [-DevicesInf <string>] [-OutputDirectory <string>] [-CSRConfigFile <string>] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Save-WindowsImage

版本：都有

模块：Dism

语法：

```powershell
Save-WindowsImage -Path <string> [-CheckIntegrity] [-Append] [-SupportEa] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：保存对已装载映像的维护更改

```powershell
PS C:\> Save-WindowsImage -Path "c:\offline"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/save-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/save-windowsimage?view=powershell-7.5)

### Send-AppvClientReport

版本：仅5.1

模块：AppvClient

语法：

```powershell
Send-AppvClientReport [[-URL] <string>] [-NetworkCostAware] [-DeleteOnSuccess] [<CommonParameters>]
```

示例：向先前配置的位置发送数据

```powershell
PS C:\> Send-AppVClientReport
The Application Virtualization Client Report was sent successfully
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/send-appvclientreport?view=powershell-5.1)

### Send-DtcDiagnosticTransaction

版本：都有

模块：MsDtc

语法：

```powershell
Send-DtcDiagnosticTransaction [-Transaction] <DtcDiagnosticTransaction> [[-ComputerName] <string>] [[-Port] <int>] [[-PropagationMethod] <DtcTransactionPropagation>] [<CommonParameters>]
```

示例：发送 DTC 诊断事务

```powershell
PS C:\> $Tx = New-DtcDiagnosticTransaction
PS C:\> Send-DtcDiagnosticTransaction -Transaction $Tx -ComputerName "Host1" -PropagationMethod Push
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/send-dtcdiagnostictransaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/send-dtcdiagnostictransaction?view=powershell-7.5)

### Set-Acl

版本：都有

模块：Microsoft.PowerShell.Security

语法（5.1）：

```powershell
Set-Acl [-Path] <string[]> [-AclObject] <Object> [[-CentralAccessPolicy] <string>] [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-Acl [-InputObject] <psobject> [-AclObject] <Object> [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
Set-Acl [-AclObject] <Object> [[-CentralAccessPolicy] <string>] -LiteralPath <string[]> [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [-UseTransaction] [<CommonParameters>]
```

语法（7）：

```powershell
Set-Acl [-Path] <string[]> [-AclObject] <Object> [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Acl [-InputObject] <psobject> [-AclObject] <Object> [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Acl [-AclObject] <Object> -LiteralPath <string[]> [-ClearCentralAccessPolicy] [-Passthru] [-Filter <string>] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将安全描述符从一个文件复制到另一个文件

```powershell
$DogACL = Get-Acl -Path "C:\Dog.txt"
Set-Acl -Path "C:\Cat.txt" -AclObject $DogACL
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/set-acl?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/set-acl?view=powershell-7.5)

### Set-AppBackgroundTaskResourcePolicy

版本：都有

模块：AppBackgroundTask

语法：

```powershell
Set-AppBackgroundTaskResourcePolicy -Mode <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将全局资源策略设为保守模式

```powershell
PS C:\> Set-AppBackgroundTaskResourcePolicy -Mode Conservative
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appbackgroundtask/set-appbackgroundtaskresourcepolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appbackgroundtask/set-appbackgroundtaskresourcepolicy?view=powershell-7.5)

### Set-AppLockerPolicy

版本：都有

模块：AppLocker

语法：

```powershell
Set-AppLockerPolicy [-XmlPolicy] <string> [-Ldap <string>] [-Merge] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppLockerPolicy [-PolicyObject] <AppLockerPolicy> [-Ldap <string>] [-Merge] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：设置本地 AppLocker 策略

```powershell
PS C:\> Set-AppLockerPolicy -XMLPolicy C:\Policy.xml
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/set-applockerpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/set-applockerpolicy?view=powershell-7.5)

### Set-AppvClientConfiguration

版本：仅5.1

模块：AppvClient

语法：

```powershell
Set-AppvClientConfiguration [-AllowHighCostLaunch <bool>] [-AutoLoad <uint32>] [-AutoCleanupEnabled <bool>] [-CertFilterForClientSsl <string>] [-EnablePackageScripts <bool>] [-EnablePublishingRefreshUI <bool>] [-IntegrationRootGlobal <string>] [-IntegrationRootUser <string>] [-LocationProvider <string>] [-MigrationMode <bool>] [-PackageInstallationRoot <string>] [-PackageSourceRoot <string>] [-RequirePublishAsAdmin <bool>] [-ReestablishmentInterval <uint32>] [-ReestablishmentRetries <uint32>] [-ReportingDataBlockSize <uint32>] [-ReportingDataCacheLimit <uint32>] [-ReportingEnabled <bool>] [-ReportingInterval <uint32>] [-ReportingRandomDelay <uint32>] [-ReportingServerURL <string>] [-ReportingStartTime <uint32>] [-RoamingFileExclusions <string>] [-RoamingRegistryExclusions <string>] [-SharedContentStoreMode <bool>] [-VerifyCertificateRevocationList <bool>] [-ExperienceImprovementOptIn <bool>] [-ProcessesUsingVirtualComponents <string[]>] [-EnableDynamicVirtualization <bool>] [-IgnoreLocationProvider <bool>] [-SupportBranchCache <bool>] [-SyncOnBatteriesEnabled <bool>] [<CommonParameters>]
```

示例：设置客户端配置参数

```powershell
PS C:\> Set-AppvClientConfiguration -parameter1 "parameterVal1"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/set-appvclientconfiguration?view=powershell-5.1)

### Set-AppvClientMode

版本：仅5.1

模块：AppvClient

语法：

```powershell
Set-AppvClientMode -Normal [<CommonParameters>]
Set-AppvClientMode -Uninstall [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/set-appvclientmode?view=powershell-5.1)

### Set-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Set-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Path <string>] [-DynamicDeploymentConfiguration <string>] [-UseNoConfiguration] [<CommonParameters>]
Set-AppvClientPackage [-Package] <AppvClientPackage> [-Path <string>] [-DynamicDeploymentConfiguration <string>] [-UseNoConfiguration] [<CommonParameters>]
Set-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Path <string>] [-DynamicDeploymentConfiguration <string>] [-UseNoConfiguration] [<CommonParameters>]
```

示例：为包设置部署配置

```powershell
PS C:\> Set-AppvClientPackage -Name "MyApp" -Version 1 -DynamicDeploymentConfiguration "C:\policies\MyApp.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/set-appvclientpackage?view=powershell-5.1)

### Set-AppvPublishingServer

版本：仅5.1

模块：AppvClient

语法：

```powershell
Set-AppvPublishingServer [-ServerId] <uint32> [[-GlobalRefreshEnabled] <bool>] [[-GlobalRefreshOnLogon] <bool>] [[-GlobalRefreshInterval] <uint32>] [[-GlobalRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [[-UserRefreshEnabled] <bool>] [[-UserRefreshOnLogon] <bool>] [[-UserRefreshInterval] <uint32>] [[-UserRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [<CommonParameters>]
Set-AppvPublishingServer [-Server] <AppvPublishingServer> [[-GlobalRefreshEnabled] <bool>] [[-GlobalRefreshOnLogon] <bool>] [[-GlobalRefreshInterval] <uint32>] [[-GlobalRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [[-UserRefreshEnabled] <bool>] [[-UserRefreshOnLogon] <bool>] [[-UserRefreshInterval] <uint32>] [[-UserRefreshIntervalUnit] <IPublishingServer+IntervalUnit>] [<CommonParameters>]
```

示例：暂无

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/set-appvpublishingserver?view=powershell-5.1)

### Set-AppxDefaultVolume

版本：都有

模块：Appx

语法：

```powershell
Set-AppxDefaultVolume [-Volume] <AppxVolume> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按路径设置默认卷

```powershell
Set-AppxDefaultVolume -Volume F:\
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/set-appxdefaultvolume?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/set-appxdefaultvolume?view=powershell-7.5)

### Set-AppxPackageAutoUpdateSettings

版本：都有

模块：Appx

语法（5.1）：

```powershell
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> [-AppInstallerUri <string>] [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint32>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -AppInstallerUri <string> -ClearPreviousSettings [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint32>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -PauseUpdates -HoursToPause <uint32> [-UseSystemPolicySource] [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> [-AppInstallerUri <string>] [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -AppInstallerUri <string> -ClearPreviousSettings [-UpdateUris <string[]>] [-RepairUris <string[]>] [-OptionalPackages <string[]>] [-DependencyPackages <string[]>] [-EnableAutomaticBackgroundTask <bool>] [-ForceUpdateFromAnyVersion <bool>] [-DisableAutoRepairs <bool>] [-CheckOnLaunch <bool>] [-ShowPrompt <bool>] [-UpdateBlocksActivation <bool>] [-UseSystemPolicySource] [-AllUsers] [-HoursBetweenUpdateChecks <uint>] [-Version <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AppxPackageAutoUpdateSettings [-PackageFamilyName] <string> -PauseUpdates -HoursToPause <uint> [-UseSystemPolicySource] [-AllUsers] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：更新应用的自动更新设置

```powershell
$params = @{
 AppInstallerUri = 'https://website.com/PackageName.appinstaller '
 PackageFamilyName = 'PackageName_8h66172c634n0 '
 CheckOnLaunch = $true
 ForceUpdateFromAnyVersion = $true
 HoursBetweenUpdateChecks = 2
 ShowPrompt = $true
 UpdateUris = 'file://ComputerName/Share/PackageName_x64.appinstaller'
}
Set-AppxPackageAutoUpdateSettings @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appx/set-appxpackageautoupdatesettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/appx/set-appxpackageautoupdatesettings?view=powershell-7.5)

### Set-AppXProvisionedDataFile

版本：都有

模块：Dism

语法：

```powershell
Set-AppXProvisionedDataFile -PackageName <string> -CustomDataPath <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-AppXProvisionedDataFile -PackageName <string> -CustomDataPath <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：为正在运行的操作系统向应用包添加自定义数据文件

```powershell
PS C:\> Set-AppXProvisionedDataFile -Online -PackageName "MyAppxPkg" -CustomDataPath "c:\Appx\myCustomData.dat"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-appxprovisioneddatafile?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-appxprovisioneddatafile?view=powershell-7.5)

### Set-AuthenticodeSignature

版本：都有

模块：Microsoft.PowerShell.Security

语法：

```powershell
Set-AuthenticodeSignature [-FilePath] <string[]> [-Certificate] <X509Certificate2> [-IncludeChain <string>] [-TimestampServer <string>] [-HashAlgorithm <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AuthenticodeSignature [-Certificate] <X509Certificate2> -LiteralPath <string[]> [-IncludeChain <string>] [-TimestampServer <string>] [-HashAlgorithm <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-AuthenticodeSignature [-Certificate] <X509Certificate2> -SourcePathOrExtension <string[]> -Content <byte[]> [-IncludeChain <string>] [-TimestampServer <string>] [-HashAlgorithm <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例（5.1）：使用本地证书存储中的证书对脚本进行签名

```powershell
$cert=Get-ChildItem -Path Cert:\CurrentUser\My -CodeSigningCert
$signingParameters = @{
 FilePath = 'PsTestInternet2.ps1'
 Certificate = $cert
 HashAlgorithm = 'SHA256'
}
Set-AuthenticodeSignature @signingParameters
```

示例（7）：使用本地证书存储中的证书对脚本进行签名

```powershell
$cert = Get-ChildItem -Path Cert:\CurrentUser\My -CodeSigningCert
Set-AuthenticodeSignature -FilePath PsTestInternet2.ps1 -Certificate $cert
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/set-authenticodesignature?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/set-authenticodesignature?view=powershell-7.5)

### Set-BcdBootDefault

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Set-BcdBootDefault [-Id] <string> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdBootDefault [-Entry] <BcdEntryInfo> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdbootdefault?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdbootdefault?view=powershell-7.5)

### Set-BcdBootDisplayOrder

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

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

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdbootdisplayorder?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdbootdisplayorder?view=powershell-7.5)

### Set-BcdBootSequence

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

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

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdbootsequence?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdbootsequence?view=powershell-7.5)

### Set-BcdBootTimeout

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Set-BcdBootTimeout [-Value] <long> [[-Store] <BcdStoreInfo>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdboottimeout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdboottimeout?view=powershell-7.5)

### Set-BcdBootToolsDisplayOrder

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

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

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdboottoolsdisplayorder?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdboottoolsdisplayorder?view=powershell-7.5)

### Set-BcdDebugSettings

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -DebugPort <long> -Serial [-Baudrate <long>] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Serial [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Port <long> -HostIp <string> -Net -Key <string> [-NoDhcp] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Port <long> -HostIp <string> -Net [-NewKey] [-NoDhcp] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Channel <long> -Ieee1394 [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Usb [-TargetName <string>] [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
Set-BcdDebugSettings [[-Store] <BcdStoreInfo>] -Local [-StartPolicy <StartPolicy>] [-NoUserModeExceptions] [-Force] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcddebugsettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcddebugsettings?view=powershell-7.5)

### Set-BcdElement

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Set-BcdElement [-Element] <string> [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Type <SetBcdElementCommand+ElementType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdElement [-Element] <string> [[-Id] <string>] [[-Store] <BcdStoreInfo>] -Device <SetBcdElementCommand+DeviceType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdElement [-Element] <string> [-Entry] <BcdEntryInfo> -Type <SetBcdElementCommand+ElementType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdElement [-Element] <string> [-Entry] <BcdEntryInfo> -Device <SetBcdElementCommand+DeviceType> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdelement?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdelement?view=powershell-7.5)

### Set-BcdHypervisorSettings

版本：都有

模块：Microsoft.Windows.Bcd.Cmdlets

语法：

```powershell
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -DebugPort <long> -Serial [-Baudrate <long>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -Ieee1394 [-Channel <long>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -HostIp <string> -Port <long> -Net [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-BcdHypervisorSettings [[-Store] <BcdStoreInfo>] -Serial [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdhypervisorsettings?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.windows.bcd.cmdlets/set-bcdhypervisorsettings?view=powershell-7.5)

### Set-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Set-BitsTransfer [-BitsJob] <BitsJob[]> [-DisplayName <string>] [-Priority <string>] [-Description <string>] [-Dynamic <switch>] [-CustomHeadersWriteOnly] [-HttpMethod <string>] [-ProxyAuthentication <string>] [-RetryInterval <int>] [-RetryTimeout <int>] [-MaxDownloadTime <int>] [-TransferPolicy <CostStates>] [-ACLFlags <ACLFlagValue>] [-SecurityFlags <SecurityFlagValue>] [-UseStoredCredential <AuthenticationTargetValue>] [-Credential <pscredential>] [-ProxyCredential <pscredential>] [-Authentication <string>] [-SetOwnerToCurrentUser] [-ProxyUsage <string>] [-ProxyList <uri[]>] [-ProxyBypass <string[]>] [-CustomHeaders <string[]>] [-NotifyFlags <NotifyFlagValue>] [-NotifyCmdLine <string[]>] [-CertStoreLocation <CertStoreLocationValue>] [-CertStoreName <string>] [-CertHash <byte[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：修改 BITS 传输作业的优先级

```powershell
PS C:\> $Bits = Get-BitsTransfer -JobId 10778CFA-C1D7-4A82-8A9D-80B19224879C
PS C:\> Set-BitsTransfer -BitsJob $Bits -Priority High
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/set-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/set-bitstransfer?view=powershell-7.5)

### Set-CertificateAutoEnrollmentPolicy

版本：都有

模块：PKI

语法：

```powershell
Set-CertificateAutoEnrollmentPolicy -PolicyState <PolicySetting> -context <Context> [-StoreName <string[]>] [-ExpirationPercentage <int>] [-EnableTemplateCheck] [-EnableMyStoreManagement] [-EnableBalloonNotifications] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CertificateAutoEnrollmentPolicy -EnableAll -context <Context> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 PolicyState = 'Enabled'
 EnableMyStoreManagement = $true
 EnableTemplateCheck = $true
 Context = 'User'
}
Set-CertificateAutoEnrollmentPolicy @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/set-certificateautoenrollmentpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/set-certificateautoenrollmentpolicy?view=powershell-7.5)

### Set-CimInstance

版本：都有

模块：CimCmdlets

语法（5.1）：

```powershell
Set-CimInstance [-InputObject] <ciminstance> [-ComputerName <string[]>] [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -CimSession <CimSession[]> -Property <IDictionary> [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint32>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -Property <IDictionary> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint32>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Set-CimInstance [-InputObject] <ciminstance> [-ComputerName <string[]>] [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-InputObject] <ciminstance> -CimSession <CimSession[]> [-ResourceUri <uri>] [-OperationTimeoutSec <uint>] [-Property <IDictionary>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -CimSession <CimSession[]> -Property <IDictionary> [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-CimInstance [-Query] <string> -Property <IDictionary> [-ComputerName <string[]>] [-Namespace <string>] [-OperationTimeoutSec <uint>] [-QueryDialect <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：设置 CIM 实例

```powershell
$instance = @ {
 Query = 'Select * from Win32_Environment where name LIKE "testvar%"'
 Property = @{VariableValue="abcd"}
}
Set-CimInstance @instance
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/set-ciminstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/cimcmdlets/set-ciminstance?view=powershell-7.5)

### Set-CIPolicyIdInfo

版本：都有

模块：ConfigCI

语法：

```powershell
Set-CIPolicyIdInfo [-FilePath] <string> [-PolicyName <string>] [-SupplementsBasePolicyID <guid>] [-BasePolicyToSupplementPath <string>] [-ResetPolicyID] [-PolicyId <string>] [-AppIdTaggingPolicy] [-AppIdTaggingKey <string[]>] [-AppIdTaggingValue <string[]>] [<CommonParameters>]
```

示例：修改策略的 ID 和名称

```powershell
PS C:\> Set-CIPolicyIdInfo -FilePath ".\Policy03.xml" -PolicyId "CIP077" -PolicyName "CIPolicy03"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-cipolicyidinfo?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-cipolicyidinfo?view=powershell-7.5)

### Set-CIPolicySetting

版本：都有

模块：ConfigCI

语法：

```powershell
Set-CIPolicySetting [-FilePath] <string> -Provider <string> -Key <string> -ValueName <string> -ValueType <string> -Value <string> [<CommonParameters>]
Set-CIPolicySetting [-FilePath] <string> -Provider <string> -Key <string> -ValueName <string> -Delete [<CommonParameters>]
```

示例：设置代码完整性策略

```powershell
Set-CIPolicySetting -FilePath C:\Policies\WDAC_policy.xml -Key "{12345678-9abc-def0-1234-56789abcdef0}" -Provider WSH -Value $True -ValueName EnterpriseDefinedClsId -ValueType Boolean
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-cipolicysetting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-cipolicysetting?view=powershell-7.5)

### Set-CIPolicyVersion

版本：都有

模块：ConfigCI

语法：

```powershell
Set-CIPolicyVersion -FilePath <string> -Version <string> [<CommonParameters>]
```

示例：更新策略版本号

```powershell
PS C:\> Set-CIPolicyVersion -FilePath '.\Policy.xml' -Version '11.1.0.2'
PS C:\> Get-Content -Path '.Policy.xml'
<?xml version="1.0" encoding="utf-8"?>
<SiPolicy xmlns="urn:schemas-microsoft-com:sipolicy">
 <VersionEx>11.1.0.2</VersionEx>
 <PolicyTypeID>{A244370E-44C9-4C06-B551-F6016E563076}</PolicyTypeID>
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-cipolicyversion?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-cipolicyversion?view=powershell-7.5)

### Set-Culture

版本：都有

模块：International

语法：

```powershell
Set-Culture [-CultureInfo] <cultureinfo> [<CommonParameters>]
```

示例：设置区域

```powershell
PS C:\> Set-Culture -CultureInfo de-DE
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-culture?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-culture?view=powershell-7.5)

### Set-DscLocalConfigurationManager

版本：都有

模块：PSDesiredStateConfiguration

语法：

```powershell
Set-DscLocalConfigurationManager [-Path] <string> [[-ComputerName] <string[]>] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-DscLocalConfigurationManager [-Path] <string> -CimSession <CimSession[]> [-Force] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：应用 LCM 设置

```powershell
Set-DscLocalConfigurationManager -Path "C:\DSC\Configurations\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/set-dsclocalconfigurationmanager?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/set-dsclocalconfigurationmanager?view=powershell-7.5)

### Set-HVCIOptions

版本：都有

模块：ConfigCI

语法：

```powershell
Set-HVCIOptions [-FilePath] <string> [-Enabled] [-Strict] [-DebugMode] [-DisableAllowed] [<CommonParameters>]
Set-HVCIOptions [-FilePath] <string> [-None] [<CommonParameters>]
```

示例：分配 Strict 选项

```powershell
PS C:\> Set-HVCIOptions -Strict -FilePath '.\Policy.xml'
PS C:\> Get-Content -Path '.Policy.xml'
 <CiSigner SignerId="ID_SIGNER_S_21" />
 </CiSigners>
 <HvciOptions>2</HvciOptions>
</SiPolicy>
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-hvcioptions?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-hvcioptions?view=powershell-7.5)

### Set-JobTrigger

版本：都有

模块：PSScheduledJob

语法：

```powershell
Set-JobTrigger [-InputObject] <ScheduledJobTrigger[]> [-DaysInterval <int>] [-WeeksInterval <int>] [-RandomDelay <timespan>] [-At <datetime>] [-User <string>] [-DaysOfWeek <DayOfWeek[]>] [-AtStartup] [-AtLogOn] [-Once] [-RepetitionInterval <timespan>] [-RepetitionDuration <timespan>] [-RepeatIndefinitely] [-Daily] [-Weekly] [-PassThru] [<CommonParameters>]
```

示例：更改作业触发器中的天数

```powershell
Get-JobTrigger -Name "DeployPackage"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/set-jobtrigger?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/set-jobtrigger?view=powershell-7.5)

### Set-KdsConfiguration

版本：都有

模块：Kds

语法：

```powershell
Set-KdsConfiguration [-LocalTestOnly] [-SecretAgreementPublicKeyLength <int>] [-SecretAgreementPrivateKeyLength <int>] [-SecretAgreementParameters <byte[]>] [-SecretAgreementAlgorithm <string>] [-KdfParameters <byte[]>] [-KdfAlgorithm <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-KdsConfiguration -RevertToDefault [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-KdsConfiguration [-InputObject] <KdsServerConfiguration> [-LocalTestOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：设置 KDS 配置

```powershell
PS C:\> $config = Get-KdsConfiguration
PS C:\> Set-KdsConfiguration -InputObject $config
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/kds/set-kdsconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/kds/set-kdsconfiguration?view=powershell-7.5)

### Set-LapsADAuditing

版本：都有

模块：LAPS

语法：

```powershell
Set-LapsADAuditing -Identity <string[]> -AuditedPrincipals <string[]> [-Credential <pscredential>] [-AuditType <AuditFlags>] [-Domain <string>] [-DomainController <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Set-LapsADAuditing -Identity LapsTestOU -AuditedPrincipals "laps.com\LapsAdmin" -AuditType Success
OU=LapsTestOU,DC=laps,DC=com
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadauditing?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadauditing?view=powershell-7.5)

### Set-LapsADComputerSelfPermission

版本：都有

模块：LAPS

语法：

```powershell
Set-LapsADComputerSelfPermission -Identity <string[]> [-Domain <string>] [-DomainController <string>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Set-LapsADComputerSelfPermission -Identity LapsTestOU
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadcomputerselfpermission?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadcomputerselfpermission?view=powershell-7.5)

### Set-LapsADPasswordExpirationTime

版本：都有

模块：LAPS

语法：

```powershell
Set-LapsADPasswordExpirationTime -Identity <string[]> [-Credential <pscredential>] [-WhenEffective <datetime>] [-Domain <string>] [-DomainController <string>] [<CommonParameters>]
```

示例：

```powershell
Set-LapsADPasswordExpirationTime -Identity lapsClient
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadpasswordexpirationtime?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadpasswordexpirationtime?view=powershell-7.5)

### Set-LapsADReadPasswordPermission

版本：都有

模块：LAPS

语法：

```powershell
Set-LapsADReadPasswordPermission -Identity <string[]> -AllowedPrincipals <string[]> [-Credential <pscredential>] [-Domain <string>] [-DomainController <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Set-LapsADReadPasswordPermission -Identity LapsTestOU -AllowedPrincipals "Domain Admins"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadreadpasswordpermission?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadreadpasswordpermission?view=powershell-7.5)

### Set-LapsADResetPasswordPermission

版本：都有

模块：LAPS

语法：

```powershell
Set-LapsADResetPasswordPermission -Identity <string[]> -AllowedPrincipals <string[]> [-Credential <pscredential>] [-Domain <string>] [-DomainController <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Set-LapsADResetPasswordPermission -Identity LapsTestOU -AllowedPrincipals "Domain Admins"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadresetpasswordpermission?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/set-lapsadresetpasswordpermission?view=powershell-7.5)

### Set-LocalGroup

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Set-LocalGroup [-InputObject] <LocalGroup> -Description <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalGroup [-Name] <string> -Description <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalGroup [-SID] <SecurityIdentifier> -Description <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：更改组说明

```powershell
Set-LocalGroup -Name "SecurityGroup04" -Description "This is a sample description."
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/set-localgroup?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/set-localgroup?view=powershell-7.5)

### Set-LocalUser

版本：都有

模块：Microsoft.PowerShell.LocalAccounts

语法：

```powershell
Set-LocalUser [-Name] <string> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-FullName <string>] [-Password <securestring>] [-PasswordNeverExpires <bool>] [-UserMayChangePassword <bool>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalUser [-InputObject] <LocalUser> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-FullName <string>] [-Password <securestring>] [-PasswordNeverExpires <bool>] [-UserMayChangePassword <bool>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-LocalUser [-SID] <SecurityIdentifier> [-AccountExpires <datetime>] [-AccountNeverExpires] [-Description <string>] [-FullName <string>] [-Password <securestring>] [-PasswordNeverExpires <bool>] [-UserMayChangePassword <bool>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：更改用户帐户的说明

```powershell
Set-LocalUser -Name "Admin07" -Description "Description of this account."
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/set-localuser?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.localaccounts/set-localuser?view=powershell-7.5)

### Set-NonRemovableAppsPolicy

版本：都有

模块：Dism

语法：

```powershell
Set-NonRemovableAppsPolicy -PackageFamilyName <string> -NonRemovable <int> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-NonRemovableAppsPolicy -PackageFamilyName <string> -NonRemovable <int> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：将应用包 Application1 设为不可移除

```powershell
PS> Set-NonRemovableAppsPolicy -Online -PackageFamilyName Application1_1.0.0.0+x64__ms7gsqeatfeb6 -NonRemovable 1
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-nonremovableappspolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-nonremovableappspolicy?view=powershell-7.5)

### Set-OsConfigurationDocument

版本：都有

模块：OsConfiguration

语法：

```powershell
Set-OsConfigurationDocument [-Content] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [-Wait] [[-TimeoutInSeconds] <int>] [-Update] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/set-osconfigurationdocument?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/set-osconfigurationdocument?view=powershell-7.5)

### Set-OsConfigurationProperty

版本：都有

模块：OsConfiguration

语法：

```powershell
Set-OsConfigurationProperty [-Name] <string> [-Value] <string> [[-SourceId] <string>] [[-Id] <string>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> {{ Add example code here }}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/set-osconfigurationproperty?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration/set-osconfigurationproperty?view=powershell-7.5)

### Set-OSConfigurationScenarioDefinition

版本：都有

模块：OsConfiguration

语法：

```powershell
Set-OsConfigurationScenarioDefinition [-Content] <string> [[-SourceId] <string>] [[-FriendlyName] <string>] [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：暂无

出处：[OsConfiguration 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/osconfiguration)（没有单独介绍）。

### Set-ProcessMitigation

版本：都有

模块：ProcessMitigations

语法：

```powershell
Set-ProcessMitigation [[-Name] <string>] [-Disable <string[]>] [-Enable <string[]>] [-EAFModules <string[]>] [-Force <string>] [-Reset] [-Remove] [<CommonParameters>]
Set-ProcessMitigation -PolicyFilePath <string> [-IsValid] [<CommonParameters>]
Set-ProcessMitigation [-Disable <string[]>] [-Enable <string[]>] [-EAFModules <string[]>] [-System] [-Force <string>] [-Reset] [-Remove] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Set-ProcessMitigation -Name Notepad.exe -Enable SEHOP -Disable ForceRelocateImages
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/processmitigations/set-processmitigation?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/processmitigations/set-processmitigation?view=powershell-7.5)

### Set-PSSessionConfiguration

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Set-PSSessionConfiguration [-Name] <string> [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PSSessionConfiguration [-Name] <string> [-AssemblyName] <string> [-ConfigurationTypeName] <string> [-ApplicationBase <string>] [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-PSVersion <version>] [-SessionTypeOption <PSSessionTypeOption>] [-TransportOption <PSTransportOption>] [-ModulesToImport <Object[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-PSSessionConfiguration [-Name] <string> -Path <string> [-RunAsCredential <pscredential>] [-ThreadApartmentState <ApartmentState>] [-ThreadOptions <PSThreadOptions>] [-AccessMode <PSSessionConfigurationAccessMode>] [-UseSharedProcess] [-StartupScript <string>] [-MaximumReceivedDataSizePerCommandMB <double>] [-MaximumReceivedObjectSizeMB <double>] [-SecurityDescriptorSddl <string>] [-ShowSecurityDescriptorUI] [-Force] [-NoServiceRestart] [-TransportOption <PSTransportOption>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例（5.1）：更改线程单元状态

```powershell
PS C:\> Set-PSSessionConfiguration -Name "MaintenanceShell" -ThreadApartmentState STA
```

示例（7）：创建和更改会话配置

```powershell
Register-PSSessionConfiguration -Name "AdminShell" -AssemblyName "C:\Shells\AdminShell.dll" -ConfigurationTypeName "AdminClass"
Set-PSSessionConfiguration -Name "AdminShell" -StartupScript "AdminConfig.ps1"
Set-PSSessionConfiguration -Name "AdminShell" -StartupScript $null
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/set-pssessionconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/set-pssessionconfiguration?view=powershell-7.5)

### Set-RecoveryManagementPluginAltitude

版本：都有

模块：Dism

语法（5.1）：

```powershell
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint32> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint32> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-RecoveryManagementPluginAltitude -ClassID <string> -Altitude <uint> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Set-RecoveryRemoteManagementStatus

版本：都有

模块：Dism

语法：

```powershell
Set-RecoveryRemoteManagementStatus -Enabled <bool> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Set-RecoveryRemoteManagementStatus -Enabled <bool> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Set-ReFSDedupSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Set-ReFSDedupSchedule [-Volume] <string> -Start <datetime> -Days <DaysOfWeek> [-Duration <timespan>] [-CpuPercentage <uint32>] [-ConcurrentOpenFiles <uint32>] [-MinimumLastModifiedTimeHours <int>] [-ExcludeFileExtension <string[]>] [-ExcludeFolder <string[]>] [-CompressionFormat <Format>] [-CompressionLevel <uint16>] [-CompressionChunkSize <uint32>] [-CompressionTuning <uint32>] [-RecompressionTuning <uint32>] [-DecompressionTuning <uint32>] [<CommonParameters>]
```

示例：

```powershell
Set-ReFSDedupSchedule -Volume "D:" -Start "10:00 PM" -Days Monday,Wednesday,Friday -Duration 4:00:00
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/set-refsdedupschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/set-refsdedupschedule?view=powershell-7.5)

### Set-ReFSDedupScrubSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Set-ReFSDedupScrubSchedule [-Volume] <string> -Start <datetime> -Days <DaysOfWeek> -WeeksInterval <uint16> [-DedupDataOnly <bool>] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 Volume = "D:"
 Start = "12/01/2024 8:00 AM"
 Days = "Monday,Thursday"
 WeeksInterval = 2
}
Set-ReFSDedupScrubSchedule @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/set-refsdedupscrubschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/set-refsdedupscrubschedule?view=powershell-7.5)

### Set-RuleOption

版本：都有

模块：ConfigCI

语法：

```powershell
Set-RuleOption [-FilePath] <string> [-Option] <int> [-Delete] [<CommonParameters>]
Set-RuleOption -Help [<CommonParameters>]
```

示例：删除规则选项

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

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-ruleoption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/configci/set-ruleoption?view=powershell-7.5)

### Set-ScheduledJob

版本：都有

模块：PSScheduledJob

语法：

```powershell
Set-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-Name <string>] [-ScriptBlock <scriptblock>] [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-MaxResultCount <int>] [-PassThru] [-ArgumentList <Object[]>] [-RunNow] [-RunEvery <timespan>] [<CommonParameters>]
Set-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-Name <string>] [-FilePath <string>] [-Trigger <ScheduledJobTrigger[]>] [-InitializationScript <scriptblock>] [-RunAs32] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-ScheduledJobOption <ScheduledJobOptions>] [-MaxResultCount <int>] [-PassThru] [-ArgumentList <Object[]>] [-RunNow] [-RunEvery <timespan>] [<CommonParameters>]
Set-ScheduledJob [-InputObject] <ScheduledJobDefinition> [-ClearExecutionHistory] [-PassThru] [<CommonParameters>]
```

示例：更改作业运行的脚本

```powershell
Get-ScheduledJob -Name "Inventory"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/set-scheduledjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/set-scheduledjob?view=powershell-7.5)

### Set-ScheduledJobOption

版本：都有

模块：PSScheduledJob

语法：

```powershell
Set-ScheduledJobOption [-InputObject] <ScheduledJobOptions> [-PassThru] [-RunElevated] [-HideInTaskScheduler] [-RestartOnIdleResume] [-MultipleInstancePolicy <TaskMultipleInstancePolicy>] [-DoNotAllowDemandStart] [-RequireNetwork] [-StopIfGoingOffIdle] [-WakeToRun] [-ContinueIfGoingOnBattery] [-StartIfOnBattery] [-IdleTimeout <timespan>] [-IdleDuration <timespan>] [-StartIfIdle] [<CommonParameters>]
```

示例：更改作业选项

```powershell
Get-ScheduledJobOption -Name "DeployPackage"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/set-scheduledjoboption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/set-scheduledjoboption?view=powershell-7.5)

### Set-SecureBootUEFI

版本：都有

模块：SecureBoot

语法：

```powershell
Set-SecureBootUEFI -Name <string> -Time <string> [-ContentFilePath <string>] [-SignedFilePath <string>] [-AppendWrite] [-OutputFilePath <string>] [<CommonParameters>]
Set-SecureBootUEFI -Name <string> -Time <string> [-Content <byte[]>] [-SignedFilePath <string>] [-AppendWrite] [-OutputFilePath <string>] [<CommonParameters>]
```

示例：设置 DBX UEFI 变量

```powershell
PS C:\> $ObjectFromFormat = ( Format-SecureBootUEFI -Name DBX -SignatureOwner 12345678-1234-1234-1234-123456789abc -Algorithm SHA256 -Hash 0011223344556677889900112233445566778899001122334455667788990011 -SignableFilePath GeneratedFileToSign.bin -Time 2011-11-01T13:30:00Z -AppendWrite )
PS C:\> .\signtool.exe sign /fd sha256 /p7 .\ /p7co 1.2.840.113549.1.7.1 /p7ce DetachedSignedData /a /f PrivateKey.pfx GeneratedFileToSign.bin
PS C:\> $ObjectFromFormat | Set-SecureBootUEFI -SignedFilePath GeneratedFileToSign.bin.p7
Name : dbx
Bytes : {161, 89, 192, 165...}
Attributes : NON VOLATILE
 BOOTSERVICE ACCESS
 RUNTIME ACCESS
 TIME BASED AUTHENTICATED WRITE ACCESS
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/set-securebootuefi?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/secureboot/set-securebootuefi?view=powershell-7.5)

### Set-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法（5.1）：

```powershell
Set-Service [-Name] <string> [-ComputerName <string[]>] [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartMode>] [-Status <string>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Service [-ComputerName <string[]>] [-DisplayName <string>] [-Description <string>] [-StartupType <ServiceStartMode>] [-Status <string>] [-InputObject <ServiceController>] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

语法（7）：

```powershell
Set-Service [-Name] <string> [-DisplayName <string>] [-Credential <pscredential>] [-Description <string>] [-StartupType <ServiceStartupType>] [-SecurityDescriptorSddl <string>] [-Status <string>] [-Force] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-Service [-InputObject] <ServiceController> [-DisplayName <string>] [-Credential <pscredential>] [-Description <string>] [-StartupType <ServiceStartupType>] [-SecurityDescriptorSddl <string>] [-Status <string>] [-Force] [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：更改显示名称

```powershell
Set-Service -Name LanmanWorkstation -DisplayName "LanMan Workstation"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/set-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/set-service?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`systemctl` + sudo）。
- 发行版：systemd 系 + sudo。
- 功能：改服务状态与自启。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 服务名（自动补 .service 后缀） |
| `-Status` | string | `running`/`started` → `systemctl start`；`stopped` → `systemctl stop` |
| `-StartupType` | string | `automatic`/`auto` → `systemctl enable`；`disabled` → `systemctl disable` |

- 实现：映射到 `systemctl start/stop/enable/disable`（必要时 sudo）。

### Set-SystemPreferredUILanguage

版本：都有

模块：LanguagePackManagement

语法：

```powershell
Set-SystemPreferredUILanguage [-Language] <string> [-PassThru] [<CommonParameters>]
```

示例：在 Windows 安装上设置系统首选 UI 语言

```powershell
Set-SystemPreferredUILanguage ja-JP
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/set-systempreferreduilanguage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/set-systempreferreduilanguage?view=powershell-7.5)

### Set-TimeZone

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Set-TimeZone [-Name] <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-TimeZone -Id <string> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-TimeZone [-InputObject] <TimeZoneInfo> [-PassThru] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按 ID 设置时区

```powershell
Set-TimeZone -Id "UTC"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/set-timezone?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/set-timezone?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`sudo timedatectl`）。
- 发行版：需要 systemd 的 timedatectl + sudo。
- 同组：Get-TimeZone。
- 功能：改时区。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 时区名（如 Asia/Shanghai）；传给 `sudo timedatectl set-timezone` |

### Set-TpmOwnerAuth

版本：都有

模块：TrustedPlatformModule

语法（5.1）：

```powershell
Set-TpmOwnerAuth -File <string> -NewFile <string> [<CommonParameters>]
Set-TpmOwnerAuth -File <string> -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewFile <string> [<CommonParameters>]
```

语法（7）：

```powershell
Set-TpmOwnerAuth -File <string> -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth -File <string> -NewFile <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewOwnerAuthorization <string> [<CommonParameters>]
Set-TpmOwnerAuth [[-OwnerAuthorization] <string>] -NewFile <string> [<CommonParameters>]
```

示例：替换导入的所有者授权值

```powershell
PS C:\> Set-TpmOwnerAuth -NewOwnerAuthorization "h4FCmNeWVNp5IMHxRfFL9QEq4vM="
TpmReady : True
TpmPresent : True
ManagedAuthLevel : Full
OwnerAuth : h4FCmNeWVNp5IMHxRfFL9QEq4vM=
OwnerClearDisabled : True
AutoProvisioning : DisabledForNextBoot
LockedOut : False
SelfTest : {191, 191, 245, 191...}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/set-tpmownerauth?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/set-tpmownerauth?view=powershell-7.5)

### Set-UevConfiguration

版本：仅5.1

模块：UEV

语法：

```powershell
Set-UevConfiguration [-CurrentComputerUser] [-MaxPackageSizeInBytes <int>] [-SettingsStoragePath <string>] [-EnableSyncProviderPing] [-DisableSyncProviderPing] [-SyncTimeoutInMilliseconds <int>] [-SyncMethod <string>] [-EnableSync] [-DisableSync] [-EnableSyncOverMeteredNetwork] [-DisableSyncOverMeteredNetwork] [-EnableSyncOverMeteredNetworkWhenRoaming] [-DisableSyncOverMeteredNetworkWhenRoaming] [-EnableSettingsImportNotify] [-DisableSettingsImportNotify] [-SettingsImportNotifyDelayInSeconds <int>] [-EnableDontSyncWindows8AppSettings] [-DisableDontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds <int>] [-EnableWaitForSyncOnApplicationStart] [-DisableWaitForSyncOnApplicationStart] [-EnableWaitForSyncOnLogon] [-DisableWaitForSyncOnLogon] [-EnableSyncUnlistedWindows8Apps] [-DisableSyncUnlistedWindows8Apps] [-VdiCollectionName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-UevConfiguration [-Computer] [-MaxPackageSizeInBytes <int>] [-SettingsStoragePath <string>] [-SettingsTemplateCatalogPath <string>] [-EnableSyncProviderPing] [-DisableSyncProviderPing] [-SyncTimeoutInMilliseconds <int>] [-SyncMethod <string>] [-EnableSync] [-DisableSync] [-EnableSyncOverMeteredNetwork] [-DisableSyncOverMeteredNetwork] [-EnableSyncOverMeteredNetworkWhenRoaming] [-DisableSyncOverMeteredNetworkWhenRoaming] [-EnableSettingsImportNotify] [-DisableSettingsImportNotify] [-SettingsImportNotifyDelayInSeconds <int>] [-ContactITUrl <string>] [-ContactITDescription <string>] [-EnableTrayIcon] [-DisableTrayIcon] [-EnableFirstUseNotification] [-DisableFirstUseNotification] [-EnableDontSyncWindows8AppSettings] [-DisableDontSyncWindows8AppSettings] [-WaitForSyncTimeoutInMilliseconds <int>] [-EnableWaitForSyncOnApplicationStart] [-DisableWaitForSyncOnApplicationStart] [-EnableWaitForSyncOnLogon] [-DisableWaitForSyncOnLogon] [-EnableSyncUnlistedWindows8Apps] [-DisableSyncUnlistedWindows8Apps] [-VdiCollectionName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：修改全部用户的同步超时设置

```powershell
PS C:\> Set-UevConfiguration -Computer -SyncTimeoutInMilliseconds 3000
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/set-uevconfiguration?view=powershell-5.1)

### Set-UevTemplateProfile

版本：仅5.1

模块：UEV

语法：

```powershell
Set-UevTemplateProfile -ID <string> -Profile <string> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：将模板与 Backup 配置文件关联

```powershell
PS C:\>Set-UevTemplateProfile -ID "MicrosoftCalculator6" -Profile "Backup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/set-uevtemplateprofile?view=powershell-5.1)

### Set-WheaMemoryPolicy

版本：都有

模块：Whea

语法（5.1）：

```powershell
Set-WheaMemoryPolicy [-ComputerName <string>] [-DisableOffline <bool>] [-DisablePFA <bool>] [-PersistMemoryOffline <bool>] [-PFAPageCount <uint32>] [-PFAErrorThreshold <uint32>] [-PFATimeout <uint32>] [<CommonParameters>]
```

语法（7）：

```powershell
Set-WheaMemoryPolicy [-ComputerName <string>] [-DisableOffline <bool>] [-DisablePFA <bool>] [-PersistMemoryOffline <bool>] [-PFAPageCount <uint>] [-PFAErrorThreshold <uint>] [-PFATimeout <uint>] [<CommonParameters>]
```

示例：启用 WHEA 预测性故障分析

```powershell
PS C:\> Set-WheaMemoryPolicy -DisablePFA $False
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/whea/set-wheamemorypolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/whea/set-wheamemorypolicy?view=powershell-7.5)

### Set-WinAcceptLanguageFromLanguageListOptOut

版本：都有

模块：International

语法：

```powershell
Set-WinAcceptLanguageFromLanguageListOptOut [-OptOut] <bool> [<CommonParameters>]
```

示例：更新注册表项

```powershell
PS C:\> Set-WinAcceptLanguageFromLanguageListOptOut -OptOut $True
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winacceptlanguagefromlanguagelistoptout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winacceptlanguagefromlanguagelistoptout?view=powershell-7.5)

### Set-WinCultureFromLanguageListOptOut

版本：都有

模块：International

语法：

```powershell
Set-WinCultureFromLanguageListOptOut [-OptOut] <bool> [<CommonParameters>]
```

示例：阻止动态设置

```powershell
PS C:\> Set-WinCultureFromLanguageListOptOut -OptOut $True
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winculturefromlanguagelistoptout?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winculturefromlanguagelistoptout?view=powershell-7.5)

### Set-WinDefaultInputMethodOverride

版本：都有

模块：International

语法：

```powershell
Set-WinDefaultInputMethodOverride [[-InputTip] <string>] [<CommonParameters>]
```

示例：设置默认输入法替代

```powershell
PS C:\> Set-WinDefaultInputMethodOverride -InputTip "0409:00000409"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-windefaultinputmethodoverride?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-windefaultinputmethodoverride?view=powershell-7.5)

### Set-WindowsEdition

版本：都有

模块：Dism

语法：

```powershell
Set-WindowsEdition -Edition <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：更改映像版本

```powershell
PS C:\> Set-WindowsEdition -Path "c:\offline" -Edition "Ultimate"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-windowsedition?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-windowsedition?view=powershell-7.5)

### Set-WindowsProductKey

版本：都有

模块：Dism

语法：

```powershell
Set-WindowsProductKey -ProductKey <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：在已装载映像上设置产品密钥

```powershell
PS C:\> Set-WindowsProductKey -Path "c:\offline" -ProductKey "12345-12345-12345-12345-12345"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-windowsproductkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-windowsproductkey?view=powershell-7.5)

### Set-WindowsReservedStorageState

版本：都有

模块：Dism

语法：

```powershell
Set-WindowsReservedStorageState -State <ReservedStorageState> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：

```powershell
PS C:\> Set-WindowsReservedStorageState -State Enabled -Online
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-windowsreservedstoragestate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/set-windowsreservedstoragestate?view=powershell-7.5)

### Set-WindowsSearchSetting

版本：都有

模块：WindowsSearch

语法：

```powershell
Set-WindowsSearchSetting [-EnableWebResultsSetting <bool>] [-EnableMeteredWebResultsSetting <bool>] [-SearchExperienceSetting <string>] [-SafeSearchSetting <string>] [<CommonParameters>]
```

示例：个性化 Windows 搜索

```powershell
PS C:\> Set-WindowsSearchSetting -SearchExperienceSetting "Personalized"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowssearch/set-windowssearchsetting?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowssearch/set-windowssearchsetting?view=powershell-7.5)

### Set-WinHomeLocation

版本：都有

模块：International

语法：

```powershell
Set-WinHomeLocation [-GeoId] <int> [<CommonParameters>]
```

示例：设置主位置

```powershell
PS C:\> Set-WinHomeLocation -GeoId 0xF4
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winhomelocation?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winhomelocation?view=powershell-7.5)

### Set-WinLanguageBarOption

版本：都有

模块：International

语法：

```powershell
Set-WinLanguageBarOption [-UseLegacySwitchMode] [-UseLegacyLanguageBar] [<CommonParameters>]
```

示例：设置语言栏选项

```powershell
PS C:\> Set-WinLanguageBarOption -UseLegacySwitchMode -UseLegacyLanguageBar
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winlanguagebaroption?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winlanguagebaroption?view=powershell-7.5)

### Set-WinSystemLocale

版本：都有

模块：International

语法：

```powershell
Set-WinSystemLocale [-SystemLocale] <cultureinfo> [<CommonParameters>]
```

示例：设置系统区域

```powershell
PS C:\> Set-WinSystemLocale -SystemLocale ja-JP
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winsystemlocale?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winsystemlocale?view=powershell-7.5)

### Set-WinUILanguageOverride

版本：都有

模块：International

语法：

```powershell
Set-WinUILanguageOverride [[-Language] <cultureinfo>] [<CommonParameters>]
```

示例：设置语言替代

```powershell
PS C:\> Set-WinUILanguageOverride -Language de-DE
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winuilanguageoverride?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winuilanguageoverride?view=powershell-7.5)

### Set-WinUserLanguageList

版本：都有

模块：International

语法：

```powershell
Set-WinUserLanguageList [-LanguageList] <List[WinUserLanguage]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：向现有语言列表添加语言

```powershell
PS C:\> $OldList = Get-WinUserLanguageList
PS C:\> $OldList.Add("fr-FR")
PS C:\> Set-WinUserLanguageList -LanguageList $OldList
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winuserlanguagelist?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/international/set-winuserlanguagelist?view=powershell-7.5)

### Set-WmiInstance

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Set-WmiInstance [-Class] <string> [[-Arguments] <hashtable>] [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance -InputObject <wmi> [-Arguments <hashtable>] [-PutType <PutType>] [-AsJob] [-ThrottleLimit <int>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance -Path <string> [-Arguments <hashtable>] [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Set-WmiInstance [-PutType <PutType>] [-AsJob] [-Impersonation <ImpersonationLevel>] [-Authentication <AuthenticationLevel>] [-Locale <string>] [-EnableAllPrivileges] [-Authority <string>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-ComputerName <string[]>] [-Namespace <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：设置 WMI 日志记录级别

```powershell
Set-WmiInstance -Class Win32_WMISetting -Arguments @{LoggingLevel=2}
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/set-wmiinstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/set-wmiinstance?view=powershell-7.5)

### Set-WSManInstance

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Set-WSManInstance [-ResourceURI] <uri> [[-SelectorSet] <hashtable>] [-ApplicationName <string>] [-ComputerName <string>] [-Dialect <uri>] [-FilePath <string>] [-Fragment <string>] [-OptionSet <hashtable>] [-Port <int>] [-SessionOption <SessionOption>] [-UseSSL] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
Set-WSManInstance [-ResourceURI] <uri> [[-SelectorSet] <hashtable>] [-ConnectionURI <uri>] [-Dialect <uri>] [-FilePath <string>] [-Fragment <string>] [-OptionSet <hashtable>] [-SessionOption <SessionOption>] [-ValueSet <hashtable>] [-Credential <pscredential>] [-Authentication <AuthenticationMechanism>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：在本地计算机上禁用侦听器

```powershell
$params = @{
 ResourceURI = 'winrm/config/listener'
 SelectorSet = @{address = '*'; transport = 'https'}
 ValueSet = @{Enabled = 'false'}
}
Set-WSManInstance @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/set-wsmaninstance?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/set-wsmaninstance?view=powershell-7.5)

### Set-WSManQuickConfig

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Set-WSManQuickConfig [-UseSSL] [-Force] [-SkipNetworkProfileCheck] [<CommonParameters>]
```

示例：通过 HTTP 启用本地计算机的远程管理

```powershell
Set-WSManQuickConfig
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/set-wsmanquickconfig?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/set-wsmanquickconfig?view=powershell-7.5)

### Show-Command

版本：都有

模块：Microsoft.PowerShell.Utility

语法：

```powershell
Show-Command [[-Name] <string>] [-Height <double>] [-Width <double>] [-NoCommonParameter] [-ErrorPopup] [-PassThru] [<CommonParameters>]
```

示例：打开“命令”窗口

```powershell
Show-Command
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/show-command?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.utility/show-command?view=powershell-7.5)

### Show-ControlPanelItem

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Show-ControlPanelItem [-Name] <string[]> [<CommonParameters>]
Show-ControlPanelItem -CanonicalName <string[]> [<CommonParameters>]
Show-ControlPanelItem [[-InputObject] <ControlPanelItem[]>] [<CommonParameters>]
```

示例：显示控制面板项

```powershell
Show-ControlPanelItem -Name "AutoPlay"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/show-controlpanelitem?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/show-controlpanelitem?view=powershell-7.5)

### Show-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Show-EventLog [[-ComputerName] <string>] [<CommonParameters>]
```

示例：显示本地计算机的事件日志

```powershell
Show-EventLog
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/show-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/show-eventlog?view=powershell-7.5)

### Show-WindowsDeveloperLicenseRegistration

版本：都有

模块：WindowsDeveloperLicense

语法：

```powershell
Show-WindowsDeveloperLicenseRegistration [<CommonParameters>]
```

示例：将设备设为可开发

```powershell
PS C:\> Show-WindowsDeveloperLicenseRegistration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowsdeveloperlicense/show-windowsdeveloperlicenseregistration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowsdeveloperlicense/show-windowsdeveloperlicenseregistration?view=powershell-7.5)

### Split-WindowsImage

版本：都有

模块：Dism

语法（5.1）：

```powershell
Split-WindowsImage -ImagePath <string> -SplitImagePath <string> -FileSize <uint64> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

语法（7）：

```powershell
Split-WindowsImage -ImagePath <string> -SplitImagePath <string> -FileSize <ulong> [-CheckIntegrity] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：拆分 .wim 文件

```powershell
PS C:\> Split-WindowsImage -ImagePath "c:\imagestore\install.wim" -SplitImagePath "c:\imagestore\splitfiles\split.swm" -FileSize 1024 -CheckIntegrity
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/split-windowsimage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/split-windowsimage?view=powershell-7.5)

### Start-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Start-BitsTransfer [-Source] <string[]> [[-Destination] <string[]>] [-Asynchronous] [-Dynamic] [-CustomHeadersWriteOnly] [-Authentication <string>] [-Credential <pscredential>] [-Description <string>] [-HttpMethod <string>] [-DisplayName <string>] [-Priority <string>] [-TransferPolicy <CostStates>] [-ACLFlags <ACLFlagValue>] [-SecurityFlags <SecurityFlagValue>] [-UseStoredCredential <AuthenticationTargetValue>] [-ProxyAuthentication <string>] [-ProxyBypass <string[]>] [-ProxyCredential <pscredential>] [-ProxyList <uri[]>] [-ProxyUsage <string>] [-RetryInterval <int>] [-RetryTimeout <int>] [-MaxDownloadTime <int>] [-Suspended] [-TransferType <string>] [-CustomHeaders <string[]>] [-NotifyFlags <NotifyFlagValue>] [-NotifyCmdLine <string[]>] [-CertStoreLocation <CertStoreLocationValue>] [-CertStoreName <string>] [-CertHash <byte[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建下载文件的 BITS 传输作业

```powershell
PS C:\> Start-BitsTransfer -Source "http://server01/servertestdir/testfile1.txt" -Destination "c:\clienttestdir\testfile1.txt"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/start-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/start-bitstransfer?view=powershell-7.5)

### Start-DscConfiguration

版本：都有

模块：PSDesiredStateConfiguration

语法：

```powershell
Start-DscConfiguration [[-Path] <string>] [[-ComputerName] <string[]>] [-Wait] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-DscConfiguration [[-Path] <string>] -CimSession <CimSession[]> [-Wait] [-Force] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-DscConfiguration [[-ComputerName] <string[]>] -UseExisting [-Wait] [-Force] [-Credential <pscredential>] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-DscConfiguration -CimSession <CimSession[]> -UseExisting [-Wait] [-Force] [-ThrottleLimit <int>] [-JobName <string>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：应用配置设置

```powershell
Start-DscConfiguration -Path "C:\DSC\Configurations\"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/start-dscconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/start-dscconfiguration?view=powershell-7.5)

### Start-DtcDiagnosticResourceManager

版本：都有

模块：MsDtc

语法：

```powershell
Start-DtcDiagnosticResourceManager [[-Port] <int>] [[-Name] <string>] [<CommonParameters>]
```

示例：启动诊断资源管理器

```powershell
PS C:\> Start-DtcDiagnosticResourceManager -Port 17124 -Name "testRM"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/start-dtcdiagnosticresourcemanager?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/start-dtcdiagnosticresourcemanager?view=powershell-7.5)

### Start-OSUninstall

版本：都有

模块：Dism

语法：

```powershell
Start-OSUninstall -Path <string> [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Start-OSUninstall -Online [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：卸载操作系统

```powershell
Start-OSUninstall -Online
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/start-osuninstall?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/start-osuninstall?view=powershell-7.5)

### Start-ReFSDedupJob

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Start-ReFSDedupJob [-Volume] <string> [-Duration <timespan>] [-FullRun] [-CpuPercentage <uint32>] [-ConcurrentOpenFiles <uint32>] [-MinimumLastModifiedTimeHours <int>] [-ExcludeFileExtension <string[]>] [-ExcludeFolder <string[]>] [-CompressionFormat <Format>] [-CompressionLevel <uint16>] [-CompressionChunkSize <uint32>] [-CompressionTuning <uint32>] [-RecompressionTuning <uint32>] [-DecompressionTuning <uint32>] [<CommonParameters>]
```

示例：

```powershell
Start-ReFSDedupJob -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/start-refsdedupjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/start-refsdedupjob?view=powershell-7.5)

### Start-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Start-Service [-InputObject] <ServiceController[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Service [-Name] <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Start-Service -DisplayName <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：使用服务名称启动服务

```powershell
Start-Service -Name "eventlog"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/start-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/start-service?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`systemctl` + sudo）。
- 发行版：systemd 系 + sudo。
- 同组：Stop-Service、Restart-Service、Resume-Service。
- 功能：启动/停止/重启服务。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 服务名；传给 `systemctl <动作> <单元>`（自动补 .service 后缀） |

- 实现：`systemctl start/stop/restart <单元>`；普通权限失败自动用 `sudo` 重试。对应 `sudo systemctl start/stop/restart`。
- Stop-Service、Restart-Service、Resume-Service 参数同 Start-Service，动作为 stop/restart/start。

### Start-Transaction

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Start-Transaction [-Timeout <int>] [-Independent] [-RollbackPreference <RollbackSeverity>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：启动和回滚事务

```powershell
Set-Location HKCU:\software
Start-Transaction
New-Item "ContosoCompany" -UseTransaction
New-ItemProperty "ContosoCompany" -Name "MyKey" -Value 123 -UseTransaction
Undo-Transaction
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/start-transaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/start-transaction?view=powershell-7.5)

### Stop-AppvClientConnectionGroup

版本：仅5.1

模块：AppvClient

语法：

```powershell
Stop-AppvClientConnectionGroup [-GroupId] <guid> [-VersionId] <guid> [-Global] [<CommonParameters>]
Stop-AppvClientConnectionGroup [-Name] <string> [-Global] [<CommonParameters>]
Stop-AppvClientConnectionGroup [-ConnectionGroup] <AppvClientConnectionGroup> [-Global] [<CommonParameters>]
```

示例：停止指定组的虚拟环境

```powershell
PS C:\> Stop-AppvClientConnectionGroup -Name "MyGroup"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/stop-appvclientconnectiongroup?view=powershell-5.1)

### Stop-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Stop-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Global] [<CommonParameters>]
Stop-AppvClientPackage [-Package] <AppvClientPackage> [-Global] [<CommonParameters>]
Stop-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [<CommonParameters>]
```

示例：关闭包某版本的虚拟环境

```powershell
PS C:\> Stop-AppvClientPackage -Name "MyPackage" -Version 2
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/stop-appvclientpackage?view=powershell-5.1)

### Stop-DtcDiagnosticResourceManager

版本：都有

模块：MsDtc

语法：

```powershell
Stop-DtcDiagnosticResourceManager [[-Job] <DtcDiagnosticResourceManagerJob>] [<CommonParameters>]
Stop-DtcDiagnosticResourceManager [[-Name] <string>] [<CommonParameters>]
Stop-DtcDiagnosticResourceManager [[-InstanceId] <guid>] [<CommonParameters>]
```

示例：停止诊断资源管理器

```powershell
PS C:\> Stop-DtcDiagnosticResourceManager -Name "testRM"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/stop-dtcdiagnosticresourcemanager?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/stop-dtcdiagnosticresourcemanager?view=powershell-7.5)

### Stop-ReFSDedupJob

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Stop-ReFSDedupJob [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Stop-ReFSDedupJob -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/stop-refsdedupjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/stop-refsdedupjob?view=powershell-7.5)

### Stop-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Stop-Service [-InputObject] <ServiceController[]> [-Force] [-NoWait] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Service [-Name] <string[]> [-Force] [-NoWait] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Stop-Service -DisplayName <string[]> [-Force] [-NoWait] [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：停止本地计算机上的服务

```powershell
Stop-Service -Name "sysmonlog"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/stop-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/stop-service?view=powershell-7.5)

#### PowerShell For Linux中的实现：

- 类型：映射 Linux 命令（`systemctl` + sudo）。
- 发行版：systemd 系 + sudo。
- 同组：Stop-Service、Restart-Service、Resume-Service。
- 功能：启动/停止/重启服务。

| 参数 | 类型 | 映射 / 说明 |
| :--- | :--- | :--- |
| `-Name`（位置 0） | string | 服务名；传给 `systemctl <动作> <单元>`（自动补 .service 后缀） |

- 实现：`systemctl start/stop/restart <单元>`；普通权限失败自动用 `sudo` 重试。对应 `sudo systemctl start/stop/restart`。
- Stop-Service、Restart-Service、Resume-Service 参数同 Start-Service，动作为 stop/restart/start。

### Suspend-BitsTransfer

版本：都有

模块：BitsTransfer

语法：

```powershell
Suspend-BitsTransfer [-BitsJob] <BitsJob[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：挂起当前用户拥有的全部 BITS 传输作业

```powershell
PS C:\> Get-BitsTransfer | Suspend-BitsTransfer
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/suspend-bitstransfer?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/bitstransfer/suspend-bitstransfer?view=powershell-7.5)

### Suspend-Job

版本：仅5.1

模块：Microsoft.PowerShell.Core

语法：

```powershell
Suspend-Job [-Id] <int[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-Job] <Job[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-InstanceId] <guid[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-Name] <string[]> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-Filter] <hashtable> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Job [-State] <JobState> [-Force] [-Wait] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：按名称挂起工作流作业

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

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/suspend-job?view=powershell-5.1)

### Suspend-ReFSDedupSchedule

版本：都有

模块：Microsoft.ReFsDedup.Commands

语法：

```powershell
Suspend-ReFSDedupSchedule [-Volume] <string> [<CommonParameters>]
```

示例：

```powershell
Suspend-ReFSDedupSchedule -Volume "D:"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/suspend-refsdedupschedule?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.refsdedup.commands/suspend-refsdedupschedule?view=powershell-7.5)

### Suspend-Service

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Suspend-Service [-InputObject] <ServiceController[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Service [-Name] <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
Suspend-Service -DisplayName <string[]> [-PassThru] [-Include <string[]>] [-Exclude <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：暂停服务

```powershell
Suspend-Service -DisplayName "Telnet"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/suspend-service?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/suspend-service?view=powershell-7.5)

### Switch-Certificate

版本：都有

模块：PKI

语法：

```powershell
Switch-Certificate [-OldCert] <Certificate> [-NewCert] <Certificate> [-NotifyOnly] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
$params = @{
 OldCert = 'Cert:\LocalMachine\My\E42DBC3B3F2771990A9B3E35D0C3C422779DACD7'
 NewCert = 'Cert:\LocalMachine\My\4A346B4385F139CA843912D358D765AB8DEE9FD4'
}
Switch-Certificate @params
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/switch-certificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/switch-certificate?view=powershell-7.5)

### Sync-AppvPublishingServer

版本：仅5.1

模块：AppvClient

语法：

```powershell
Sync-AppvPublishingServer [-ServerId] <uint32> [-Global] [-Force] [-NetworkCostAware] [-HidePublishingRefreshUI] [<CommonParameters>]
Sync-AppvPublishingServer [-Server] <AppvPublishingServer> [-Global] [-Force] [-NetworkCostAware] [-HidePublishingRefreshUI] [<CommonParameters>]
Sync-AppvPublishingServer [[-Name] <string>] [[-URL] <string>] [-Global] [-Force] [-NetworkCostAware] [-HidePublishingRefreshUI] [<CommonParameters>]
Sync-AppvPublishingServer [-PublishFromXML] [-Global] [-NetworkCostAware] [<CommonParameters>]
```

示例：启动发布刷新

```powershell
PS C:\> Sync-AppvPublishingServer -Name "MyServer"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/sync-appvpublishingserver?view=powershell-5.1)

### Test-AppLockerPolicy

版本：都有

模块：AppLocker

语法：

```powershell
Test-AppLockerPolicy [-XmlPolicy] <string> -Path <List[string]> [-User <string>] [-Filter <List[PolicyDecision]>] [<CommonParameters>]
Test-AppLockerPolicy [-XmlPolicy] <string> -Packages <List[AppxPackage]> [-User <string>] [-Filter <List[PolicyDecision]>] [<CommonParameters>]
Test-AppLockerPolicy [-PolicyObject] <AppLockerPolicy> -Path <List[string]> [-User <string>] [-Filter <List[PolicyDecision]>] [<CommonParameters>]
```

示例：报告程序是否允许运行

```powershell
PS C:\> Test-AppLockerPolicy -XMLPolicy C:\Policy.xml -Path c:\windows\system32\calc.exe, C:\windows\system32\notepad.exe -User Everyone
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/test-applockerpolicy?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/applocker/test-applockerpolicy?view=powershell-7.5)

### Test-Certificate

版本：都有

模块：PKI

语法：

```powershell
Test-Certificate [-Cert] <Certificate> [-Policy <TestCertificatePolicy>] [-User] [-EKU <string[]>] [-DNSName <string>] [-AllowUntrustedRoot] [<CommonParameters>]
```

示例：

```powershell
Get-ChildItem -Path Cert:\LocalMachine\My |
 Test-Certificate -Policy SSL -DNSName 'dns=contoso.com'
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/pki/test-certificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/pki/test-certificate?view=powershell-7.5)

### Test-ComputerSecureChannel

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Test-ComputerSecureChannel [-Repair] [-Server <string>] [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：测试本地计算机与其域之间的通道

```powershell
Test-ComputerSecureChannel
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/test-computersecurechannel?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/test-computersecurechannel?view=powershell-7.5)

### Test-DscConfiguration

版本：都有

模块：PSDesiredStateConfiguration

语法：

```powershell
Test-DscConfiguration [[-ComputerName] <string[]>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-AsJob] [-Detailed] [<CommonParameters>]
Test-DscConfiguration [-Path] <string> [[-ComputerName] <string[]>] [-Credential <pscredential>] [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration [[-ComputerName] <string[]>] -ReferenceConfiguration <string> [-Credential <pscredential>] [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration [-Path] <string> -CimSession <CimSession[]> [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration -CimSession <CimSession[]> -ReferenceConfiguration <string> [-ThrottleLimit <int>] [-AsJob] [<CommonParameters>]
Test-DscConfiguration -CimSession <CimSession[]> [-ThrottleLimit <int>] [-AsJob] [-Detailed] [<CommonParameters>]
```

示例：本地计算机的测试配置

```powershell
Test-DscConfiguration
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/test-dscconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psdesiredstateconfiguration/test-dscconfiguration?view=powershell-7.5)

### Test-FileCatalog

版本：都有

模块：Microsoft.PowerShell.Security

语法：

```powershell
Test-FileCatalog [-CatalogFilePath] <string> [[-Path] <string[]>] [-Detailed] [-FilesToSkip <string[]>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：创建和验证文件目录

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

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/test-filecatalog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.security/test-filecatalog?view=powershell-7.5)

### Test-KdsRootKey

版本：都有

模块：Kds

语法：

```powershell
Test-KdsRootKey [-KeyId] <guid> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：测试根密钥配置

```powershell
PS C:\> Test-KdsRootKey -KeyId 4A3615F1-5A90-22E4-0B1D-1416F93D4412
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/kds/test-kdsrootkey?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/kds/test-kdsrootkey?view=powershell-7.5)

### Test-PSSessionConfigurationFile

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Test-PSSessionConfigurationFile [-Path] <string> [<CommonParameters>]
```

示例：测试会话配置文件

```powershell
Test-PSSessionConfigurationFile -Path "FullLanguage.pssc"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/test-pssessionconfigurationfile?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/test-pssessionconfigurationfile?view=powershell-7.5)

### Test-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Test-UevTemplate [-Path] <string[]> [<CommonParameters>]
Test-UevTemplate -LiteralPath <string[]> [<CommonParameters>]
```

示例：测试文件

```powershell
PS C:\> Test-UevTemplate -Path "MicrosoftWordpad.xml" | Format-Table -AutoSize
Path Status Message
---- ------ -------
C:\Program Files\Microsoft User Experience Virtualization\Templates\MicrosoftWordpad.xml Valid The template is valid.
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/test-uevtemplate?view=powershell-5.1)

### Test-WSMan

版本：都有

模块：Microsoft.WSMan.Management

语法：

```powershell
Test-WSMan [[-ComputerName] <string>] [-Authentication <AuthenticationMechanism>] [-Port <int>] [-UseSSL] [-ApplicationName <string>] [-Credential <pscredential>] [-CertificateThumbprint <string>] [<CommonParameters>]
```

示例：确定 WinRM 服务的状态

```powershell
Test-WSMan
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/test-wsman?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.wsman.management/test-wsman?view=powershell-7.5)

### Unblock-Tpm

版本：都有

模块：TrustedPlatformModule

语法：

```powershell
Unblock-Tpm [[-OwnerAuthorization] <string>] [<CommonParameters>]
Unblock-Tpm -File <string> [<CommonParameters>]
```

示例：重置锁定

```powershell
PS C:\>Unblock-Tpm -OwnerAuthorization "vjnuW6rToM41os3xxEpjLdIW2gA="
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/unblock-tpm?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/trustedplatformmodule/unblock-tpm?view=powershell-7.5)

### Undo-DtcDiagnosticTransaction

版本：都有

模块：MsDtc

语法：

```powershell
Undo-DtcDiagnosticTransaction [-Transaction] <DtcDiagnosticTransaction> [<CommonParameters>]
```

示例：撤销 DTC 诊断事务

```powershell
PS C:\> $Tx = New-DtcDiagnosticTransaction
PS C:\> Undo-DtcDiagnosticTransaction -Transaction $Tx
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/undo-dtcdiagnostictransaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/msdtc/undo-dtcdiagnostictransaction?view=powershell-7.5)

### Undo-Transaction

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Undo-Transaction [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：回滚当前事务

```powershell
Undo-Transaction
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/undo-transaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/undo-transaction?view=powershell-7.5)

### Uninstall-Language

版本：都有

模块：LanguagePackManagement

语法：

```powershell
Uninstall-Language [-Language] <string> [-PassThru] [<CommonParameters>]
```

示例：从设备删除语言

```powershell
Uninstall-Language ja-jp
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/uninstall-language?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/languagepackmanagement/uninstall-language?view=powershell-7.5)

### Uninstall-ProvisioningPackage

版本：都有

模块：Provisioning

语法：

```powershell
Uninstall-ProvisioningPackage [-PackageId] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Uninstall-ProvisioningPackage -PackagePath <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Uninstall-ProvisioningPackage -AllInstalledPackages [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
Uninstall-ProvisioningPackage [-RuntimeMetadata] <RuntimeProvPackageMetadata> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：卸载全部预配包

```powershell
PS C:\> Uninstall-ProvisioningPackage -AllInstalledPackages
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/uninstall-provisioningpackage?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/uninstall-provisioningpackage?view=powershell-7.5)

### Uninstall-TrustedProvisioningCertificate

版本：都有

模块：Provisioning

语法：

```powershell
Uninstall-TrustedProvisioningCertificate [-Thumbprint] <string> [-LogsDirectoryPath <string>] [-WprpFile <string>] [-ConnectedDevice] [<CommonParameters>]
```

示例：卸载受信任预配证书

```powershell
PS C:\> Uninstall-TrustedProvisioningCertificate -Thumbprint fedd995b45e633d4ef30fcbc8f3a48b627e9a28b
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/uninstall-trustedprovisioningcertificate?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/provisioning/uninstall-trustedprovisioningcertificate?view=powershell-7.5)

### Unpublish-AppvClientPackage

版本：仅5.1

模块：AppvClient

语法：

```powershell
Unpublish-AppvClientPackage [-PackageId] <guid> [-VersionId] <guid> [-Global] [-UserSID <string>] [<CommonParameters>]
Unpublish-AppvClientPackage [-Package] <AppvClientPackage> [-Global] [-UserSID <string>] [<CommonParameters>]
Unpublish-AppvClientPackage [-Name] <string> [[-Version] <string>] [-Global] [-UserSID <string>] [<CommonParameters>]
```

示例：取消发布包版本

```powershell
PS C:\> Unpublish-AppvClientPackage -Name "MyApp" -Version 3
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/appvclient/unpublish-appvclientpackage?view=powershell-5.1)

### Unregister-PSSessionConfiguration

版本：都有

模块：Microsoft.PowerShell.Core

语法：

```powershell
Unregister-PSSessionConfiguration [-Name] <string> [-Force] [-NoServiceRestart] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除会话配置

```powershell
Unregister-PSSessionConfiguration -Name "MaintenanceShell"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/unregister-pssessionconfiguration?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.core/unregister-pssessionconfiguration?view=powershell-7.5)

### Unregister-RecoveryManagementPlugin

版本：都有

模块：Dism

语法：

```powershell
Unregister-RecoveryManagementPlugin -ClassID <string> -Path <string> [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Unregister-RecoveryManagementPlugin -ClassID <string> -Online [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：暂无

出处：[Dism 模块文档](https://learn.microsoft.com/zh-cn/powershell/module/dism)（没有单独介绍）。

### Unregister-ScheduledJob

版本：都有

模块：PSScheduledJob

语法：

```powershell
Unregister-ScheduledJob [-InputObject] <ScheduledJobDefinition[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-ScheduledJob [-Id] <int[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-ScheduledJob [-Name] <string[]> [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：删除计划作业

```powershell
Unregister-ScheduledJob TestJob
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/unregister-scheduledjob?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/psscheduledjob/unregister-scheduledjob?view=powershell-7.5)

### Unregister-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Unregister-UevTemplate [-ID] <string> [-WhatIf] [-Confirm] [<CommonParameters>]
Unregister-UevTemplate -All [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：注销模板

```powershell
PS C:\> Unregister-UevTemplate -TemplateId "MicrosoftCalculator6"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/unregister-uevtemplate?view=powershell-5.1)

### Unregister-WindowsDeveloperLicense

版本：都有

模块：WindowsDeveloperLicense

语法：

```powershell
Unregister-WindowsDeveloperLicense [-Force] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：禁用开发者模式

```powershell
PS C:\> Unregister-WindowsDeveloperLicense
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/windowsdeveloperlicense/unregister-windowsdeveloperlicense?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/windowsdeveloperlicense/unregister-windowsdeveloperlicense?view=powershell-7.5)

### Update-LapsADSchema

版本：都有

模块：LAPS

语法：

```powershell
Update-LapsADSchema [-Credential <pscredential>] [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：

```powershell
Update-LapsADSchema
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/laps/update-lapsadschema?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/laps/update-lapsadschema?view=powershell-7.5)

### Update-UevTemplate

版本：仅5.1

模块：UEV

语法：

```powershell
Update-UevTemplate [-Path] <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
Update-UevTemplate -LiteralPath <string[]> [-WhatIf] [-Confirm] [<CommonParameters>]
```

示例：更新模板

```powershell
PS C:\> Update-UevTemplate -Path "MicrosoftCalculator.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/uev/update-uevtemplate?view=powershell-5.1)

### Update-WIMBootEntry

版本：都有

模块：Dism

语法：

```powershell
Update-WIMBootEntry -Path <string> -ImagePath <string> -DataSourceID <long> [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：更新 .wim 文件的配置项

```powershell
PS C:\> Update-WIMBootEntry -Path "C:\" -DataSourceID 0 -ImagePath "D:\Windows Images\install.wim"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/update-wimbootentry?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/update-wimbootentry?view=powershell-7.5)

### Use-Transaction

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Use-Transaction [-TransactedScript] <scriptblock> [-UseTransaction] [<CommonParameters>]
```

示例：使用启用了事务的对象编写脚本

```powershell
Start-Transaction
$transactedString = New-Object Microsoft.PowerShell.Commands.Management.TransactedString
$transactedString.Append("Hello")
Use-Transaction -TransactedScript { $transactedString.Append(", World") } -UseTransaction
$transactedString.ToString()
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/use-transaction?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/use-transaction?view=powershell-7.5)

### Use-WindowsUnattend

版本：都有

模块：Dism

语法：

```powershell
Use-WindowsUnattend -UnattendPath <string> -Path <string> [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
Use-WindowsUnattend -UnattendPath <string> -Online [-NoRestart] [-WindowsDirectory <string>] [-SystemDrive <string>] [-LogPath <string>] [-ScratchDirectory <string>] [-LogLevel <LogLevel>] [<CommonParameters>]
```

示例：将应答文件应用到已装载映像

```powershell
PS C:\> Use-WindowsUnattend -Path "c:\offline" -UnattendPath "c:\test\answerfiles\myunattend.xml"
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/dism/use-windowsunattend?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/dism/use-windowsunattend?view=powershell-7.5)

### Write-EventLog

版本：都有

模块：Microsoft.PowerShell.Management

语法：

```powershell
Write-EventLog [-LogName] <string> [-Source] <string> [-EventId] <int> [[-EntryType] <EventLogEntryType>] [-Message] <string> [-Category <int16>] [-RawData <byte[]>] [-ComputerName <string>] [<CommonParameters>]
```

示例：将事件写入应用程序事件日志

```powershell
PS C:\> Write-EventLog -LogName "Application" -Source "MyApp" -EventID 3001 -EntryType Information -Message "MyApp added a user-requested feature to the display." -Category 1 -RawData 10,20
```

出处：[官方中文文档（5.1）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/write-eventlog?view=powershell-5.1) / [官方中文文档（7）](https://learn.microsoft.com/zh-cn/powershell/module/microsoft.powershell.management/write-eventlog?view=powershell-7.5)

