#!/bin/bash

rm *.so
rm piuma.h
rm piuma.c

go build -buildmode=c-shared -o piuma.so *.go

python pygen.py > piuma.c

gcc piuma.c -dynamiclib piuma.so -o pypiuma.so `pkg-config --cflags --libs python3` -I /Library/Frameworks/Python.framework/Versions/3.6/include/python3.6m/

