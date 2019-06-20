#!/bin/bash

set -e

# Check params
if [ $# -ne 1 ]
	then
		echo Usage: $0 version;
    echo E.g: $0 0.1.0
		echo Version is MAJOR.MINOR.BUGFIX
		echo Latest version:
		git describe --abbrev=0
		exit 1;
fi

# Set environment
export LANG="C.UTF-8"
export VERSION=$1

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

echo ----[ Compile ]----
./gradlew goClean goBuild

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
cp -rv .gogradle/sendmail-to-msmtp-linux-amd64 $DEB_PATH/usr/sbin/sendmail

cd $DEB_PATH/..
dpkg-deb --no-uniform-compression --build sendmail-to-msmtp
mv sendmail-to-msmtp.deb $DEB_FILE

echo ----[ Upload to Bintray ]----
cd $RUN_PATH
curl -T $DEB_PATH/../$DEB_FILE -u$BINTRAY_USER:$BINTRAY_KEY "https://api.bintray.com/content/foilen/debian/sendmail-to-msmtp/$VERSION/$DEB_FILE;deb_distribution=stable;deb_component=main;deb_architecture=amd64;publish=1"

echo ----[ Git Tag ]==----
git tag -a -m $VERSION $VERSION

echo ----[ Operation completed successfully ]==----

echo
echo You can send the tag: git push --tags
