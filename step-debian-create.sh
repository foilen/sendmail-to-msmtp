#!/bin/bash

set -e

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Create .deb ]----
DEB_FILE=sendmail-to-msmtp_${VERSION}_amd64.deb
DEB_PATH=$RUN_PATH/build/debian_out/sendmail-to-msmtp
rm -rf $DEB_PATH
mkdir -p $DEB_PATH $DEB_PATH/DEBIAN/ $DEB_PATH/usr/sbin/

cat > $DEB_PATH/DEBIAN/control << _EOF
Package: sendmail-to-msmtp
Version: $VERSION
Maintainer: Foilen
Architecture: amd64
Description: This is a bridge from sendmail to msmtp. The goal is to support all the different ways to provide arguments and translate them to msmtp.
_EOF

cp -rv DEBIAN $DEB_PATH/
cp -rv build/bin/* $DEB_PATH/usr/sbin/

cd $DEB_PATH/..
dpkg-deb --no-uniform-compression --build sendmail-to-msmtp
mv sendmail-to-msmtp.deb $DEB_FILE
