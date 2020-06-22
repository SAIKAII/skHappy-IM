package encrypto

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
)

// EnCryptoPassword 用盐加密密码，并返回base64结果
func EnCryptoPassword(pwd, salt string) string {
	pwdSalt := sha256.Sum256(bytes.NewBufferString(pwd + salt).Bytes())
	return base64.StdEncoding.EncodeToString(pwdSalt[:])
}
