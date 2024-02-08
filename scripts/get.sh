#!/bin/bash

set -e
set -x

# shellcheck disable=SC1090
. <(curl -fL "https://raw.githubusercontent.com/black-desk/get/master/get.sh") \
	black-desk busagent

$SUDO install -m755 -D "$TMP_DIR/busagent" "$PREFIX/bin/busagent"
