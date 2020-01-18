package hash

import (
	"crypto"
	_ "crypto/sha256"
	"encoding/hex"
)

func CalculateSha256AndHex(data []byte) (hexStr string) {
	var sha256Hash = crypto.SHA256.New()
	sha256Hash.Write(data)
	hexStr = hex.EncodeToString(sha256Hash.Sum(nil))
	return
}
