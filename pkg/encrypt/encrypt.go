package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5WithSalt(str, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(str))
	m5.Write([]byte(string(salt)))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
