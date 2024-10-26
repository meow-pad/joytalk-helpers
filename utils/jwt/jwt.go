package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/meow-pad/persian/utils/json"
	"hash"
	"strings"
)

type JoyTalkJWTHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func NewJoytalkJWTHeader() *JoyTalkJWTHeader {
	return &JoyTalkJWTHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
}

type JoyTalkJWTPayload struct {
	Iat    int64  `json:"iat"`
	Exp    int64  `json:"exp"`
	AppId  string `json:"appId"`
	Digest string `json:"digest,omitempty"`
}

func NewJoytalkJWTPayload(iat int64, appId string, digest string) *JoyTalkJWTPayload {
	return &JoyTalkJWTPayload{
		Iat:    iat,
		Exp:    iat + 600,
		AppId:  appId,
		Digest: digest,
	}
}

func BuildBase64JoytalkJWTHeader(header *JoyTalkJWTHeader) string {
	hStr := json.ToString(header)
	base64Str := base64.RawURLEncoding.EncodeToString([]byte(hStr))
	return base64Str
}

func BuildBase64JoytalkJWTPayload(payload *JoyTalkJWTPayload) string {
	pStr := json.ToString(payload)
	base64Str := base64.RawURLEncoding.EncodeToString([]byte(pStr))
	return base64Str
}

func BuildSha256Hash(secret []byte) hash.Hash {
	return hmac.New(sha256.New, secret)
}

func BuildSignature(base64Header, base64Payload string, hash hash.Hash) string {
	str := base64Header + "." + base64Payload
	hash.Reset()
	hash.Write([]byte(str))
	signature := hash.Sum(nil)
	return string(signature)
}

func BuildBase64Signature(base64Header, base64Payload string, hash hash.Hash) string {
	signature := BuildSignature(base64Header, base64Payload, hash)
	return base64.RawURLEncoding.EncodeToString([]byte(signature))
}

func BuildJoytalkToken(base64Header, base64Payload string, hash hash.Hash) string {
	signature := BuildBase64Signature(base64Header, base64Payload, hash)
	return strings.Join([]string{base64Header, base64Payload, signature}, ".")
}

func CheckToken(appId string, h hash.Hash, iat, exp int64, digest string, srcToken string) bool {
	header := BuildBase64JoytalkJWTHeader(NewJoytalkJWTHeader())
	payload := BuildBase64JoytalkJWTPayload(&JoyTalkJWTPayload{
		Iat:    iat,
		Exp:    exp,
		AppId:  appId,
		Digest: digest,
	})
	token := BuildJoytalkToken(header, payload, h)
	return srcToken == token
}
