#!/bin/bash

rm *.profile
rm *.grind

go run piuma.go piumawrapper.go profile.go
go tool pprof -callgrind -output=piumago.grind ./piumago.profile

echo "Profile file 'piumago.grind' created!"
