$WorkDir = (Get-Location).Path
$BinDir = Join-Path $WorkDir "bin"

Write-Host ""
Write-Host " ---- [Start] Running all in docker [Start]-----"
Write-Host ""
Write-Host "Work dir     : $WorkDir"
Write-Host "Binaries dir : $BinDir"

# docker run --rm -it -v "$PWD":/usr/src/myapp -v "$GOPATH":/go -w /usr/src/myapp golang:1.17.8

# Check if the 'results' directory exists, and create it if not
if (-Not (Test-Path -Path "$BinDir/results" -PathType Container)) {
    New-Item -Path "$BinDir/results" -ItemType Directory | Out-Null
}

Write-Host ""

# Run the Docker command and capture output
docker run --rm -it -v "$WorkDir":/usr/src/myapp -v "$Env:GOPATH":/go -w /usr/src/myapp golang:1.17.8 bash -c ' \
    ./bin/cli-linux-amd64 2>&1 | tee bin/results/linux-amd64.out; cat bin/results/linux-amd64.out \
'

Write-Host "Running complete"
Write-Host ""
Write-Host "Output"
Write-Host ""
Write-Host ""

# Display the contents of the 'results' directory
Write-Host "ls -la $BinDir/results"
Get-Content "$BinDir/results/linux-amd64.out"
(Get-ChildItem -Path "$BinDir/results" | Format-Table -Property Name, Length)

Write-Host "`$Path:"
Write-Host "export PATH=`$PATH:`"$BinDir`""
Write-Host "running : ${BinDir}/cli-linux-amd64"
Write-Host ""
Write-Host " ---- [End] Running in docker [end]-----"
Write-Host ""
