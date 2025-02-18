#!/bin/bash

function run_backend() {
    chmod +x ./tools/executables/scripts/backend.sh
    ./tools/executables/scripts/backend.sh
}

function run_mockgen() {
    chmod +x ./tools/executables/scripts/mockgen.sh
    ./tools/executables/scripts/mockgen.sh "$@"
}

function run_sqlc() {
    chmod +x ./tools/executables/scripts/sqlc.sh
    ./tools/executables/scripts/sqlc.sh
}

case "$1" in
    -backend)
        run_backend
        ;;
    -mockgen)
        shift
        run_mockgen "$@"
        ;;
    -sqlc)
        run_sqlc
        ;;
    *)
        echo "Invalid option. Usage: ./run.sh <-backend|-sqlc>"
        ;;
esac
