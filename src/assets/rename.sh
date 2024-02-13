#!/bin/bash

counter=1
directory="restaurant"

for file in "$directory"/*; do
    if [ -f "$file" ]; then
        new_name="restau_$counter.jpg"
        mv "$file" "$directory/$new_name"
        ((counter++))
    fi
done
