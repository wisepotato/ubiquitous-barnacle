#!/bin/bash

die () {
    echo >&2 "$@"
    exit 1
}

[ "$#" -eq 4 ] || die "Usage: ./run.sh [go | cpp] [euler] [maarten | harro] [challenge #]"

BASEDIR=$(dirname $0)
LANGUAGE="$1"
WORKDIR="${BASEDIR}/${LANGUAGE}/$2"
PERSON="$3"
CHALLENGE="$4"

./${WORKDIR}/${CHALLENGE}/${PERSON}.bin
