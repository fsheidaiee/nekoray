#!/bin/bash
set -e

source libs/deploy_common.sh
[ "$GOOS" == "windows" ] && [ "$GOARCH" == "amd64" ] && DEST=$DEPLOYMENT/windows64 || true
[ "$GOOS" == "linux" ] && [ "$GOARCH" == "amd64" ] && DEST=$DEPLOYMENT/linux64 || true
[ "$GOOS" == "darwin" ] && [ "$GOARCH" == "amd64" ] && DEST=$DEPLOYMENT/macos-amd64 || true
[ "$GOOS" == "darwin" ] && [ "$GOARCH" == "arm64" ] && DEST=$DEPLOYMENT/macos-arm64 || true
if [ -z $DEST ]; then
  echo "Please set GOOS GOARCH"
  exit 1
fi
rm -rf $DEST
mkdir -p $DEST

export CGO_ENABLED=0

#### Go: updater ####
pushd updater
[ "$GOOS" == "darwin" ] || go build -o $DEST -trimpath -ldflags "-w -s"
[ "$GOOS" == "linux" ] && mv $DEST/updater $DEST/launcher || true
popd

#### Go: nekoray_core ####
pushd ../v2ray-core
version_v2ray=$(git log --pretty=format:'%h' -n 1)
popd
pushd go/cmd/nekoray_core
go build -v -o $DEST -trimpath -ldflags "-w -s -X neko/pkg/neko_common.Version_v2ray=$version_v2ray -X neko/pkg/neko_common.Version_neko=$version_standalone"
popd

#### Go: nekobox_core ####
pushd go/cmd/nekobox_core
go build -v -o $DEST -trimpath -ldflags "-w -s -X neko/pkg/neko_common.Version_neko=$version_standalone" -tags "with_gvisor,with_quic,with_wireguard,with_v2ray_api,with_ech,with_utls"
popd
