#!/bin/bash

die () {
    echo >&2 "$@"
    exit 1
}

[ "$#" -eq 3 ] || die "Usage: ./build.sh [euler] [maarten | harro] [challenge #]"

BASEDIR=$(dirname $0)
WORKDIR="${BASEDIR}/$1"
PERSON="$2"
CHALLENGE="$3"

go build -o "${WORKDIR}/${CHALLENGE}/${PERSON}.bin" ${WORKDIR}/${CHALLENGE}/${PERSON}/*
