#!/bin/bash

die () {
    echo >&2 "$@"
    exit 1
}

[ "$#" -eq 3 ] || die "Usage: ./run.sh [euler] [maarten | harro] [challenge #]"

BASEDIR=$(dirname $0)
WORKDIR="${BASEDIR}/$1"
PERSON="$2"
CHALLENGE="$3"

./${WORKDIR}/${CHALLENGE}/${PERSON}.bin
