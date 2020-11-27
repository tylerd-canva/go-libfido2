package libfido2

/*
#cgo darwin LDFLAGS: -framework CoreFoundation -framework IOKit ${SRCDIR}/darwin/lib/libfido2.a ${SRCDIR}/darwin/lib/libcrypto.a ${SRCDIR}/darwin/lib/libcbor.a
#cgo darwin CFLAGS: -I${SRCDIR}/include
*/
import "C"
