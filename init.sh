#!/bin/bash

# set up a new day

if [ $# -ne 2 ]; then
    echo "Usage: init.sh YEAR DAY"
    echo "  e.g. init.sh 2021 1"
    echo
    exit 1
fi

YEAR=$1
DAY=$2

AOCDIR="${HOME}/go/src/github.com/thisisdavidbell/adventofcode"
NEWDIR=${AOCDIR}/${YEAR}/day${DAY}

mkdir -p "${NEWDIR}"
cp -rf "${AOCDIR}"/template/* "${NEWDIR}"/.

cd "${NEWDIR}" || echo "${NEWDIR} missing" && exit