# cmdconsts — Command Constants

Package `cmdconsts` provides string constants for common shell commands, flags, and service management operations used in deployment scripts and system automation.

## Categories

| Category | Examples |
|----------|---------|
| **Package Managers** | `Apt`, `AptGet`, `AptInstallYes`, `AptGetUpdate` |
| **Service Management** | `SystemCtlStart`, `SystemCtlStop`, `SystemCtlRestart`, `SystemCtlEnable` |
| **File Operations** | `Chown`, `ChmodCommand`, `RmRf`, `Ln`, `LnHyphenS`, `Touch` |
| **User/Group Management** | `UserAdd`, `UserMod`, `GroupAdd`, `GroupDelete`, `DelUser` |
| **Networking** | `IPCommand`, `Ping`, `Hostname`, `NetPlan`, `NMCli` |
| **Shell** | `BashDefaultPath`, `Sudo`, `Echo`, `Grep`, `Xargs` |
| **Flags** | `HyphenA` through `HyphenZ`, `FlagYes`, `FlagPurge`, `FlagSystem` |

## Usage

```go
import "github.com/alimtvnetwork/core/cmdconsts"

cmd := cmdconsts.SystemCtlRestart + " nginx"
// "systemctl restart nginx"

script := cmdconsts.AptGetUpdateYes + cmdconsts.SingleLineScriptsJoiner + cmdconsts.AptGetInstallYes + " curl"
// "apt-get update -y && apt-get install -y curl"
```

## Related Docs

- [Coding Guidelines](/spec/01-app/17-coding-guidelines.md)
