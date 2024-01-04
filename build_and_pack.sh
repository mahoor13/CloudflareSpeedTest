#!/bin/bash

# clean the old Releases folder
if [ -d Releases ]; then
    rm -r Releases
fi

mkdir Releases

version="v2.3.6"

# Platforms to build
platforms=("darwin/amd64" "darwin/arm64" "linux/386" "linux/amd64" "linux/arm64" "linux/arm/v5" "linux/arm/v6" "linux/arm/v7" "linux/mips" "linux/mips64")

# Build and pack for each platform
for platform in "${platforms[@]}"
do
    IFS='/' read -r -a platform_info <<< "$platform"
    GOOS="${platform_info[0]}"
    if [ -n "${platform_info[2]}" ]; then
        GOARCH="${platform_info[1]}"
        ARCHNAME="${platform_info[1]}${platform_info[2]}"
        GOARM="${platform_info[2]:1}"
    else
        GOARCH="${platform_info[1]}"
        ARCHNAME="${platform_info[1]}"
        GOARM=""
    fi

    echo "Building ${ARCHNAME} for ${GOOS}"
    # Build binary
    GOOS=$GOOS GOARCH=$GOARCH GOARM=$GOARM go build -o CloudflareST -ldflags "-s -w -X main.version=$version"

    # Pack binary
    if [ "$GOOS" == "windows" ]; then
        mv CloudflareST CloudflareST.exe
        echo "Packing into Releases/CloudflareST_${GOOS}_${ARCHNAME}.zip"
        zip -m Releases/CloudflareST_"$GOOS"_"$ARCHNAME".zip CloudflareST.exe
        rm CloudflareST.exe
    else
        echo "Packing into Releases/CloudflareST_${GOOS}_${ARCHNAME}.tar.gz"
        tar -czf Releases/CloudflareST_"$GOOS"_"$ARCHNAME".tar.gz CloudflareST
        rm CloudflareST
    fi

    # Clean up binary
done

echo "Build and pack completed."
