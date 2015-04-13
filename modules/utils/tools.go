package utils

import (
	"fmt"
	//"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	//"fmt"
	"hash"
	//"math/big"
	//"net/url"
	//"reflect"
	//"strconv"
	"time"

	"github.com/astaxie/beego"

	"train_system/setting"
)

// http://code.google.com/p/go/source/browse/pbkdf2/pbkdf2.go?repo=crypto
func PBKDF2(password, salt []byte, iter, keyLen int, h func() hash.Hash) []byte {
	prf := hmac.New(h, password)
	hashLen := prf.Size()
	numBlocks := (keyLen + hashLen - 1) / hashLen
	var buf [4]byte
	dk := make([]byte, 0, numBlocks*hashLen)
	U := make([]byte, hashLen)
	for block := 1; block <= numBlocks; block++ {
		// N.B.: || means concatenation, ^ means XOR
		// for each block T_i = U_1 ^ U_2 ^ ... ^ U_iter
		// U_1 = PRF(password, salt || uint(i))
		prf.Reset()
		prf.Write(salt)
		buf[0] = byte(block >> 24)
		buf[1] = byte(block >> 16)
		buf[2] = byte(block >> 8)
		buf[3] = byte(block)
		prf.Write(buf[:4])
		dk = prf.Sum(dk)
		T := dk[len(dk)-hashLen:]
		copy(U, T)
		// U_n = PRF(password, U_(n-1))
		for n := 2; n <= iter; n++ {
			prf.Reset()
			prf.Write(U)
			U = U[:0]
			U = prf.Sum(U)
			for x := range U {
				T[x] ^= U[x]
			}
		}
	}
	return dk[:keyLen]
}

// Encode string to md5 hex value
func EncodeMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}

// use pbkdf2 encode password
func EncodePassword(rawPwd string, salt string, save_len int) string {
	pwd := PBKDF2([]byte(rawPwd), []byte(salt), 10000, save_len, sha256.New)
	return hex.EncodeToString(pwd)
}

// Random generate string
func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// verify time limit code
func VerifyTimeLimitCode(data string, minutes int, code string) bool {
	if len(code) <= 18 {
		return false
	}

	// split code
	start := code[:12]
	lives := code[12:18]
	if d, err := StrTo(lives).Int(); err == nil {
		minutes = d
	}

	// right active code
	retCode := CreateTimeLimitCode(data, minutes, start)
	if retCode == code && minutes > 0 {
		// check time is expired or not
		before, _ := beego.DateParse(start, "YmdHi")
		now := time.Now()
		if before.Add(time.Minute*time.Duration(minutes)).Unix() > now.Unix() {
			return true
		}
	}
	return false
}

const TimeLimitCodeLength = 12 + 6 + 40

// create a time limit code
// code format: 12 length date time string + 6 minutes string + 40 sha1 encoded string
func CreateTimeLimitCode(data string, minutes int, startInf interface{}) string {
	format := "YmdHi"

	var start, end time.Time
	var startStr, endStr string

	if startInf == nil {
		// Use now time create code
		start = time.Now()
		startStr = beego.Date(start, format)
	} else {
		// use start string create code
		startStr = startInf.(string)
		start, _ = beego.DateParse(startStr, format)
		startStr = beego.Date(start, format)
	}

	end = start.Add(time.Minute * time.Duration(minutes))
	endStr = beego.Date(end, format)

	// create sha1 encode string
	sh := sha1.New()
	sh.Write([]byte(data + setting.SecretKey + startStr + endStr + ToStr(minutes)))
	encoded := hex.EncodeToString(sh.Sum(nil))

	code := fmt.Sprintf("%s%06d%s", startStr, minutes, encoded)
	return code
}
