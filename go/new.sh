#!/bin/bash

die () {
    echo >&2 "$@"
    exit 1
}

[ "$#" -eq 2 ] || die "Usage: ./new.sh [euler] [challenge #]"

BASEDIR=$(dirname $0)
WORKDIR="${BASEDIR}/$1"
CHALLENGE="$2"

[ ! -d "${WORKDIR}/${CHALLENGE}" ] || die "Directory already exists: ${WORKDIR}/${CHALLENGE}"

mkdir -p "${WORKDIR}/${CHALLENGE}/harro"
mkdir -p "${WORKDIR}/${CHALLENGE}/maarten"
touch "${WORKDIR}/${CHALLENGE}/exercise.txt"
touch "${WORKDIR}/${CHALLENGE}/harro/main.go"
touch "${WORKDIR}/${CHALLENGE}/maarten/main.go"

echo "Directory contents created: ${WORKDIR}/${CHALLENGE}"
