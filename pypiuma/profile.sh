#!/bin/bash

go run *.go
go tool pprof -callgrind -output=piumago.grind ./piumago.profile

echo "Profile file 'piumago.grind' created!"
