#!/bin/bash

# Usage: ./build.sh <GO_VERSION>

# Check if the GO_VERSION argument is provided
if [ $# -ne 1 ]; then
  echo "Usage: $0 <GO_VERSION>"
  exit 1
fi

# Extract CLI name from the parent directory name
CliName="$(basename "$(dirname "$PWD")")"

# Parse the GO_VERSION argument
GO_VERSION="$1"

WorkDir="$PWD/.."
BinDir="$WorkDir/bin"
Platforms="darwin linux windows"
Architectures="386 amd64"

echo "Selected Go version: $GO_VERSION"
echo "Supported platforms: $Platforms"
echo "Supported architectures: $Architectures"

# Define the logic for Windows and non-Windows platforms
windows_logic="
if [ \$GOOS == \"windows\" ]; then
  echo \"Building \$GOOS-\$GOARCH.exe\"
  go build -o \"bin/cli-\$FinalCliName.exe\" \"cmd/$CliName/*.go\"
else
  echo \"Building \$GOOS-\$GOARCH\"
  go build -o \"bin/cli-\$FinalCliName\" \"cmd/$CliName/*.go\"
fi
"

non_windows_logic="
echo \"Building \$GOOS-\$GOARCH\"
go build -o \"bin/cli-\$FinalCliName\" \"cmd/$CliName/*.go\"
"

# Define the command to be executed inside the Docker container
build_command="
rm -rf bin && \
mkdir bin && \
cp -R assets bin && \
cp -R configs bin && \
for GOOS in $Platforms; do
  for GOARCH in $Architectures; do
    export GOOS=\$GOOS
    export GOARCH=\$GOARCH

    # Construct the output CLI name
    FinalCliName=\"${CliName}-\$GOOS-\$GOARCH\"

    if [ \$GOOS == \"windows\" ]; then
      $windows_logic
    else
      $non_windows_logic
    fi
  done
done
"

echo ""
echo " ---- [Start] deploy for platforms ($Platforms [$Architectures]) [Start]-----"
echo ""
echo "Work dir     : $WorkDir"
echo "Binaries dir : $BinDir"

# Run the Docker container with the specified Go version and build command
docker run --rm -it -v "$WorkDir":/usr/src/myapp -v "$GOPATH":/go -w /usr/src/myapp "golang:$GO_VERSION" bash -c "$build_command" &&
echo "Build complete" &&
echo "" &&
echo "Permission adding:" &&
echo "chown -R root:root $BinDir" &&
echo "chmod -R 777 $BinDir" &&
echo "" &&
chown -R root:root "$BinDir" &&
chmod -R 777 "$BinDir" &&
echo "" &&
echo "ls -la $BinDir:" &&
ls -la "$BinDir" &&
echo "" &&
echo "EnvPath" &&
export PATH=$PATH:"$BinDir" &&
echo "" &&
echo $PATH &&
echo "" &&
echo "$\"Path\":" &&
echo "export PATH=\$PATH:\"$BinDir\"" &&
echo "running : ${BinDir}/${FinalCliName}" &&
"$BinDir"/"$FinalCliName" &&
echo "" &&
echo " ---- [End] deploy for all platforms ($Platforms [$Architectures]) [end]-----" &&
echo ""
