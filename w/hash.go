package w

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	b64 "encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// date: 2021/12/10
// email: brach@lssin.com
var Hash *hash

func init() {
	Hash = &hash{}
}

func (h *hash) _pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (h *hash) _pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 加密
func (h *hash) AesEncrypt_CBC(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = h._pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (h *hash) AesDecrypt_CBC(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = h._pKCS7UnPadding(origData)
	return origData, nil
}

func (h *hash) HmacEncryption_GL(key_name string, unix int64, data map[string]interface{}) string {
	//t := time.Now().Unix()
	var data_key []string
	for _k, _ := range data {
		data_key = append(data_key, _k)
	}
	json_str := "["
	sort.Strings(data_key)
	for _, k := range data_key {
		k = fmt.Sprintf("[\"%v\", \"%v\"], ", k, data[k])
		json_str += k
	}
	sort_str := strings.Trim(json_str, ", ") + "]" + strconv.FormatInt(int64((unix)), 10)
	mac := hmac.New(sha1.New, []byte(key_name))
	mac.Write([]byte(sort_str))
	return hex.EncodeToString(mac.Sum(nil))
}

func (h *hash) UrlEncryption_GL(data map[string]interface{}) string {
	var str_list []string
	for k, v := range data {
		str_list = append(str_list, fmt.Sprintf("%v=%v", url.QueryEscape(k), url.QueryEscape(v.(string))))
	}
	return strings.Join(str_list, "&")
}

func (h *hash) Md5(content string) string {
	x := md5.Sum([]byte(content))
	return hex.EncodeToString(x[:])
}

func (h *hash) Sha1(content string) string {
	x := sha1.Sum([]byte(content))
	return hex.EncodeToString(x[:])
}

// 加密
func (h hash) Base64_Encode(content []byte) string {
	return b64.StdEncoding.EncodeToString(content)
}

func (h hash) Base64_Decode(content string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(content)
}
