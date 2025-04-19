package hasher

import (
	"crypto/sha256"
)

func HashIt(data string) []byte {
	h := sha256.New()
	h.Write([]byte(data))
	return h.Sum(nil)
}