#!/usr/bin/env bash
set -e

VERSION=$(git describe --tags --always)
echo "当前目录为:$(pwd)"

make build PROJECT="$PROJECT" IMAGE_REPO="$IMAGE_REPO" VERSION="$VERSION"
