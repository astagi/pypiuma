# PyPiuma
Experimental porting of Piuma code (written in Go) in Python 3

## Install Python dependencies

    pip install -r requirements.txt

## Install Go dependencies

    go get github.com/nfnt/resize

## Install go-libjpeg using libjpeg-turbo

    export CGO_LDFLAGS="-L/usr/local/Cellar/jpeg-turbo/2.0.0/lib/" CGO_CPPFLAGS="-I/usr/local/Cellar/jpeg-turbo/2.0.0/include/" LD_LIBRARY_PATH="/usr/local/Cellar/jpeg-turbo/2.0.0/lib/"

    go install github.com/pixiv/go-libjpeg/jpeg

## Build the library

    cd pypiuma && ./build.sh

## Try yourself

If you want to try the pure Python implementation, run

    python purepiuma.py

For the Go implementation

    python piuma.py

export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/Library/Frameworks/Python.framework/Versions/3.6/lib/pkgconfig/
