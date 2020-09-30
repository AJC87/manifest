#!/bin/bash
#
# Install the application

go build main.go

mv main /usr/local/bin/manifest

rm -f main
