#!/bin/bash
set -e

if [ "$#" -eq 0 ]; then
    echo "Error: No command provided."
    exit 1
fi

exec "$@"


