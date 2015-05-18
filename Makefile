export PKG_CONFIG_PATH=$(shell pwd)/lib/pkgconfig

install: deps

deps: fetch lib

fetch:
	git submodule sync
	git submodule update --init

libsass-build:
	# generate configure scripts
	cd libsass-src; make clean && autoreconf -fvi
	- rm -rf libsass-build/
	mkdir -p libsass-build
	# configure and install libsass
	cd libsass-build && \
		../libsass-src/configure --disable-shared --prefix=$(shell pwd) --disable-silent-rules --disable-dependency-tracking

lib: libsass-build
	cd libsass-build && make install

.PHONY: test
test:
	go test -race .

clean:
	rm -rf libsass-build lib include
