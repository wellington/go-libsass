deps:
	cd libsass-src; autoreconf -fvi && \
		./configure --disable-shared --prefix=$(shell pwd) --disable-silent-rules --disable-dependency-tracking && \
		make install
