#!/bin/bash

die () {
    echo >&2 "$@"
    exit 1
}

[ "$#" -eq 4 ] || die "Usage: ./build.sh [go | cpp] [euler] [maarten | harro] [challenge #]"

BASEDIR=$(dirname $0)
LANGUAGE="$1"
WORKDIR="${BASEDIR}/${LANGUAGE}/$2"
PERSON="$3"
CHALLENGE="$4"

if [ ${LANGUAGE} = "go" ]; then
  go build -o "${WORKDIR}/${CHALLENGE}/${PERSON}.bin" ${WORKDIR}/${CHALLENGE}/${PERSON}/*
fi
