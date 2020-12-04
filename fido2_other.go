package libfido2

/*
#cgo linux LDFLAGS: -lfido2 -lcbor -lcrypto -ludev
#cgo windows LDFLAGS: -L${SRCDIR}/windows/lib -lfido2
#cgo windows CFLAGS: -I${SRCDIR}/include
*/
import "C"
