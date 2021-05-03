#!/bin/bash

set -e

RUN_PATH="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $RUN_PATH

IPFS_ROOT=/com.foilen.deploy
IPFS_KEY_NAME=com.foilen.deploy

echo ----[ Add to IPFS ]----
DEB_FILE=sendmail-to-msmtp_${VERSION}_amd64.deb
DEB_PATH=$RUN_PATH/build/debian_out/sendmail-to-msmtp
IPFS_FILE_ID=$(ipfs add -q $DEB_PATH/../$DEB_FILE | tail -n1)
echo IPFS_FILE_ID: $IPFS_FILE_ID

echo ----[ Put to IPFS under $IPFS_ROOT/sendmail-to-msmtp/$DEB_FILE ]----
ipfs files cp /ipfs/$IPFS_FILE_ID $IPFS_ROOT/sendmail-to-msmtp/$DEB_FILE

echo ----[ Publish to IPFS $IPFS_ROOT ]----
IPFS_ROOT_DIR_ID=$(ipfs files stat $IPFS_ROOT | head -n 1)
ipfs name publish --key $IPFS_KEY_NAME $IPFS_ROOT_DIR_ID
