#!/usr/bin/env sh
set -x

export PATH=$PATH:/kaijud/kaijud
BINARY=/kaijud/kaijud
ID=${ID:-0}
LOG=${LOG:-kaijud.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found."
	exit 1
fi

export KAIJUDHOME="/kaijud/data/node${ID}/kaijud"

if [ -d "$(dirname "${KAIJUDHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${KAIJUDHOME}" "$@" | tee "${KAIJUDHOME}/${LOG}"
else
  "${BINARY}" --home "${KAIJUDHOME}" "$@"
fi
