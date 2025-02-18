#!/bin/bash

function secondary_mockgen(){
    local actor="$1"

    case "$actor" in
        Account | Auth | Session | task)
            mockgen -source="./src/core/interfaces/secondary/${actor^}Loader.go" -destination="./src/core/services/mocks/${actor^}RepositoryMock.go"
            ;;
        *)
            echo "Invalid actor. Available actors: task"
            return 1
            ;;
    esac
}

function primary_mockgen(){
    local actor="$1"

    case "$actor" in
        Account | Auth | task)
            mockgen -source="./src/core/interfaces/primary/${actor^}Manager.go" -destination="./src/api/handlers/mocks/${actor^}ServicesMock.go"
            ;;
        *)
            echo "Invalid actor. Available actors: task"
            return 1
            ;;
    esac
}

function generate_secondary_mocks(){
    local actors=()
    while [[ $# -gt 0 ]]; do
        case $1 in
            -Account | -Auth | -Session | -task)
                actors+=("${1#-}")
                shift
                ;;
            *)
                echo "Invalid actor."
                return 1
                ;;
        esac
    done

    for actor in "${actors[@]}"; do
        secondary_mockgen "$actor"
    done
}

function generate_primary_mocks(){
    local actors=()
    while [[ $# -gt 0 ]]; do
        case $1 in
            -Account | -Auth | -task)
                actors+=("${1#-}")
                shift
                ;;
            *)
                echo "Invalid actor."
                return 1
                ;;
        esac
    done

    for actor in "${actors[@]}"; do
        primary_mockgen "$actor"
    done
}

case "$1" in
    -secondary)
        shift
        if [[ $# -eq 0 ]]; then
            generate_secondary_mocks -Account -Auth -Session -task
        else
            generate_secondary_mocks "$@"
        fi
        ;;
    -primary)
        shift
        if [[ $# -eq 0 ]]; then
            generate_primary_mocks -Account -Auth -task
        else
            generate_primary_mocks "$@"
        fi
        ;;
    -all)
        generate_primary_mocks -Account -Auth -task
        generate_secondary_mocks -Account -Auth -Session -task
        ;;
    *)
        echo "Invalid option for mock generation. Usage: -mockgen <-primary|-secondary|-all>"
        ;;
esac
