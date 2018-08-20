# PyPiuma
Experimental porting of Piuma code (written in Go) in Python 3

## Install Python dependencies

    pip install -r requirements.txt

## Install Go dependencies

    go get github.com/nfnt/resize

## Build the library

    ./build.sh

## Try yourself

If you want to try the pure Python implementation, run

    python purepiuma.py

For the Go implementation

    python piuma.py

export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/Library/Frameworks/Python.framework/Versions/3.6/lib/pkgconfig/
