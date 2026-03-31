# ─────────────────────────────────────────────────────────────────────────────
# DashboardBoxPrimitives.psm1 — Progress bar + box-drawing primitives
#
# Dependencies: DashboardTheme.psm1 (script-scope color variables)
# ─────────────────────────────────────────────────────────────────────────────

function Get-ProgressBar {
    <# .SYNOPSIS Returns a colored progress bar string using ANSI + Unicode block chars. #>
    [CmdletBinding()]
    [OutputType([string])]
    param([Parameter(Mandatory)][int]$Score, [int]$MaxScore = 100, [int]$BarWidth = 15)

    if ($MaxScore -le 0) { $MaxScore = 100 }
    $percentage  = [math]::Min(1.0, [math]::Max(0.0, $Score / $MaxScore))
    $filledCount = [math]::Round($percentage * $BarWidth)
    $emptyCount  = $BarWidth - $filledCount
    $filled = if ($filledCount -gt 0) { [string]::new([char]0x2588, $filledCount) } else { "" }
    $empty  = if ($emptyCount  -gt 0) { [string]::new([char]0x2592, $emptyCount)  } else { "" }
    return "$($script:cLime)$filled$($script:cBarE)$empty$($script:cReset)"
}

function Write-BoxTop {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)╔$("═" * $Width)╗$($script:cReset)"
}

function Write-BoxBottom {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)╚$("═" * $Width)╝$($script:cReset)"
}

function Write-BoxDivider {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)╠$("═" * $Width)╣$($script:cReset)"
}

function Write-BoxEmptyLine {
    [CmdletBinding()]
    param([int]$Width = $script:BoxWidth)
    Write-Host "$($script:cBorder)║$($script:cReset)$(" " * $Width)$($script:cBorder)║$($script:cReset)"
}

function Write-BoxLine {
    [CmdletBinding()]
    param([string]$Content, [int]$Width = $script:BoxWidth, [int]$VisualLength = -1)
    if ($VisualLength -ge 0) {
        $rightPad = [math]::Max(0, $Width - $VisualLength - 1)
        Write-Host "$($script:cBorder)║$($script:cReset) $Content$(" " * $rightPad)$($script:cBorder)║$($script:cReset)"
    } else {
        Write-Host "$($script:cBorder)║$($script:cReset) $Content"
    }
}

function Write-BoxLineCenter {
    [CmdletBinding()]
    param([string]$Text, [int]$Width = $script:BoxWidth, [string]$Color = "")
    if (-not $Color) { $Color = $script:cWhite }
    $textLen = $Text.Length
    $leftPad = [math]::Max(0, [math]::Floor(($Width - $textLen) / 2))
    $rightPad = [math]::Max(0, $Width - $textLen - $leftPad)
    $line = (" " * $leftPad) + $Text + (" " * $rightPad)
    Write-Host "$($script:cBorder)║$($script:cReset)$Color$($script:cBold)$line$($script:cReset)$($script:cBorder)║$($script:cReset)"
}

Export-ModuleMember -Function @(
    'Get-ProgressBar', 'Write-BoxTop', 'Write-BoxBottom', 'Write-BoxDivider',
    'Write-BoxEmptyLine', 'Write-BoxLine', 'Write-BoxLineCenter'
)
