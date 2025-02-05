#!/usr/bin/env bash

# based on https://leg100.github.io/en/posts/building-bubbletea-programs/#3-live-reload-code-changes

# watch code changes, trigger re-build, and kill process 
while true; do
    make build && pkill -f 'bin/roverctl'
    inotifywait -e attrib $(find . -name '*.go') || exit
done