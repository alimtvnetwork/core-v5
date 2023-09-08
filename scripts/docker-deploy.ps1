param (
    [string]$GO_VERSION
)

# Check if the GO_VERSION argument is provided
if (-not $GO_VERSION) {
    Write-Host "Usage: ./build.ps1 -GO_VERSION <GO_VERSION>"
    exit 1
}

# Extract CLI name from the parent directory name
$CliName = (Get-Item $PSScriptRoot).Parent.Name

$WorkDir = (Get-Location).Path
$BinDir = Join-Path $WorkDir "bin"
$Platforms = "darwin", "linux", "windows"
$Architectures = "386", "amd64"

Write-Host "Selected Go version: $GO_VERSION"
Write-Host "Supported platforms: $Platforms"
Write-Host "Supported architectures: $Architectures"

# Define the Bash script for building in the Docker container
$build_command = @"
rm -rf bin
mkdir bin
cp -R assets bin
cp -R configs bin
$Platforms | ForEach-Object {
    GOOS=\$_
$Architectures | ForEach-Object {
        GOARCH=\$_

        # Construct the output CLI name
        FinalCliName="${CliName}-\$GOOS-\$GOARCH"

        if [ "\$GOOS" == "windows" ]; then
            echo "Building \$GOOS-\$GOARCH.exe"
            go build -o "bin/cli-\$FinalCliName.exe" "cmd/$CliName/*.go"
        else
            echo "Building \$GOOS-\$GOARCH"
            go build -o "bin/cli-\$FinalCliName" "cmd/$CliName/*.go"
        fi
    }
}
"@

Write-Host ""
Write-Host " ---- [Start] deploy for platforms ($($Platforms -join ", ") [$($Architectures -join ", ")]) [Start]-----"
Write-Host ""
Write-Host "Work dir     : $WorkDir"
Write-Host "Binaries dir : $BinDir"

# Run the Docker container with the specified Go version and build command
docker run --rm -it `
    -v "$WorkDir:/usr/src/myapp" `
    -v "${Env:GOPATH}:/go" `
    -w /usr/src/myapp `
    "golang:$GO_VERSION" `
    bash -c "$build_command"

Write-Host "Build complete"
