#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

# Add the architecture for building image
platforms=("linux/amd64" "linux/arm64" "darwin/arm64" "darwin/amd64" "windows/amd64" "windows/arm64")
for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name="$package_name-$GOOS-$GOARCH"
    if [[ $GOOS == "windows" ]]; then
      output_name+=".exe" # Add the .exe extension for Windows binaries
    fi

    env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
