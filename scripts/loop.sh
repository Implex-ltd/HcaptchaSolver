#!/bin/bash

EXECUTABLE="."

run_program() {
    while true; do
        echo "Starting the API..."
        go run "$EXECUTABLE"
        echo "API crashed. Restarting in 1 second..."
        sleep 1
    done
}

run_program