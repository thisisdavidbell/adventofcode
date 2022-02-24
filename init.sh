# script which will setup a new day

if [ $# -ne 3 ]; then
    echo "Usage: init.sh YEAR DAY NAME"
    echo "  e.g. init.sh 2021 1 ocean"
    echo
    return 1
fi

YEAR=$1
DAY=$2
NAME=$3

AOCDIR=~/go/src/github.com/thisisdavidbell/adventofcode
NEWDIR=$AOCDIR/${YEAR}/day${DAY}

cp -rf $AOCDIR/template $NEWDIR

cd ${NEWDIR}

mv template.go ${NAME}.go
mv template_test.go ${NAME}_test.go

# Visual Studio Code prefers a single go.mod at top of repo, so using that instead
# go mod init
# go mod edit -replace github.com/thisisdavidbell/adventofcode/utils=../../utils
