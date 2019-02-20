#!/bin/bash

die () {
    echo >&2 "$@"
    exit 1
}

[ "$#" -eq 3 ] || die "Usage: ./new.sh [go | cpp] [euler] [challenge #]"

BASEDIR=$(dirname $0)
LANGUAGE="$1"
WORKDIR="${BASEDIR}/${LANGUAGE}/$2"
CHALLENGE="$3"

[ ! -d "${WORKDIR}/${CHALLENGE}" ] || die "Directory already exists: ${WORKDIR}/${CHALLENGE}"

mkdir -p "${WORKDIR}/${CHALLENGE}/harro"
mkdir -p "${WORKDIR}/${CHALLENGE}/maarten"
touch "${WORKDIR}/${CHALLENGE}/README.md"
touch "${WORKDIR}/${CHALLENGE}/harro/main.go"
touch "${WORKDIR}/${CHALLENGE}/maarten/main.go"

echo "Directory contents created: ${WORKDIR}/${CHALLENGE}"
