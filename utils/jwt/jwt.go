package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/meow-pad/persian/utils/json"
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
	Iat   int64  `json:"iat"`
	Exp   int64  `json:"exp"`
	AppId string `json:"appId"`
}

func NewJoytalkJWTPayload(iat int64, appId string) *JoyTalkJWTPayload {
	return &JoyTalkJWTPayload{
		Iat:   iat,
		Exp:   iat + 600,
		AppId: appId,
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

func BuildSignature(base64Header, base64Payload string, secret []byte) string {
	str := base64Header + "." + base64Payload
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(str))
	signature := h.Sum(nil)
	return string(signature)
}

func BuildBase64Signature(base64Header, base64Payload string, secret []byte) string {
	signature := BuildSignature(base64Header, base64Payload, secret)
	return base64.RawURLEncoding.EncodeToString([]byte(signature))
}

func BuildJoytalkToken(base64Header, base64Payload string, secret []byte) string {
	signature := BuildBase64Signature(base64Header, base64Payload, secret)
	return strings.Join([]string{base64Header, base64Payload, signature}, ".")
}

func CheckToken(appId string, secret []byte, iat, exp int64, srcToken string) bool {
	header := BuildBase64JoytalkJWTHeader(NewJoytalkJWTHeader())
	payload := BuildBase64JoytalkJWTPayload(&JoyTalkJWTPayload{
		Iat:   iat,
		Exp:   exp,
		AppId: appId,
	})
	token := BuildJoytalkToken(header, payload, secret)
	return srcToken == token
}
