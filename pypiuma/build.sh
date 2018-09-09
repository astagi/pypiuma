#!/bin/bash

rm *.so
rm *.a

# go build -buildmode=c-shared -o piuma.so *.go

# python pygen.py > piuma.c

# gcc piuma.c -dynamiclib piuma.so -o pypiuma.so `pkg-config --cflags --libs python3`


go build -buildmode=c-archive -o libpiuma.a main.go piuma.go piumawrapper.go

python pygen.py > piuma.c

gcc piuma.c -shared -o pypiuma.so `pkg-config --cflags --libs python3` -framework CoreFoundation -framework Security -L . -lpiuma
