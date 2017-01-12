#!/bin/bash

go install v2ray.com/core/tools/build

$GOPATH/bin/build --os=windows --arch=x86 --zip
$GOPATH/bin/build --os=windows --arch=x64 --zip
$GOPATH/bin/build --os=macos --arch=x64 --zip
$GOPATH/bin/build --os=linux --arch=x86 --zip
$GOPATH/bin/build --os=linux --arch=x64 --zip
$GOPATH/bin/build --os=linux --arch=arm --zip
$GOPATH/bin/build --os=linux --arch=arm64 --zip
$GOPATH/bin/build --os=linux --arch=mips64 --zip
$GOPATH/bin/build --os=freebsd --arch=x86 --zip
$GOPATH/bin/build --os=freebsd --arch=amd64 --zip
$GOPATH/bin/build --os=openbsd --arch=x86 --zip
$GOPATH/bin/build --os=openbsd --arch=amd64 --zip

INSTALL_DIR=_install

git clone "https://github.com/v2ray/install.git" ${INSTALL_DIR}

rm -rf ${INSTALL_DIR}/releases/
mkdir ${INSTALL_DIR}/releases/
cp $GOPATH/bin/metadata.txt ${INSTALL_DIR}/releases/
cp $GOPATH/bin/v2ray-*.zip ${INSTALL_DIR}/releases/

cp $GOPATH/bin/v2ray-${TRAVIS_TAG}-linux-64/v2ray ${INSTALL_DIR}/docker/official/

pushd ${INSTALL_DIR}
git config user.name "V2Ray Auto Build"
git config user.email "admin@v2ray.com"
git add -A
git commit -m "Update for ${TRAVIS_TAG}"
git push "https://${GIT_KEY_INSTALL}@github.com/v2ray/install.git" master
popd

DOCKER_HUB_API=https://registry.hub.docker.com/u/v2ray/official/trigger/${DOCKER_HUB_KEY}/
curl -H "Content-Type: application/json" --data '{"build": true}' -X POST "${DOCKER_HUB_API}"
