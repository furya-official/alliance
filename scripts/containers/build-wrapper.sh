#!/usr/bin/env sh
set -x

export PATH=$PATH:/furyad/furyad
BINARY=/furyad/furyad
ID=${ID:-0}
LOG=${LOG:-furyad.log}

if ! [ -f "${BINARY}" ]; then
	echo "The binary $(basename "${BINARY}") cannot be found."
	exit 1
fi

export FURYADHOME="/furyad/data/node${ID}/furyad"

if [ -d "$(dirname "${FURYADHOME}"/"${LOG}")" ]; then
  "${BINARY}" --home "${FURYADHOME}" "$@" | tee "${FURYADHOME}/${LOG}"
else
  "${BINARY}" --home "${FURYADHOME}" "$@"
fi
