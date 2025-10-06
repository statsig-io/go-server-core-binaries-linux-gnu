//go:build linux && gnu && x86_64

package go_server_core_binaries_linux_gnu

import (
	_ "embed"
)

//go:embed linux_gnu_x86_64.so
var binaryData []byte

//go:embed linux_gnu_x86_64.so.sig
var signatureData []byte

func GetBinaryData() []byte {
	return binaryData
}

func GetSignatureData() []byte {
	return signatureData
}
