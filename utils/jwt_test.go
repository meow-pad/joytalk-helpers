package utils

import "testing"

var (
	_jwtTestTime  = int64(1556188391)
	_jwtTestAppId = "xxxxxxxxxxxx"
	_jwtTestPass  = "123456"
)

func Test_BuildBase64JoytalkJWTHeader(t *testing.T) {
	header := BuildBase64JoytalkJWTHeader(NewJoytalkJWTHeader())
	if header != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" {
		t.Errorf("invalid jwt header:%s", header)
	}
}

func Test_BuildBase64JoytalkJWTPayload(t *testing.T) {
	payload := BuildBase64JoytalkJWTPayload(NewJoytalkJWTPayload(_jwtTestTime, _jwtTestAppId))
	if payload != "eyJpYXQiOjE1NTYxODgzOTEsImV4cCI6MTU1NjE4ODk5MSwiYXBwSWQiOiJ4eHh4eHh4eHh4eHgifQ" {
		t.Errorf("invalid jwt payload:%s", payload)
	}
}

func Test_BuildBase64JoytalkJWTSignature(t *testing.T) {
	header := BuildBase64JoytalkJWTHeader(NewJoytalkJWTHeader())
	payload := BuildBase64JoytalkJWTPayload(NewJoytalkJWTPayload(_jwtTestTime, _jwtTestAppId))
	signature := BuildBase64Signature(header, payload, []byte(_jwtTestPass))
	if signature != "DFdsMFu_VzPkUUS1eu_Kwyzvc6vSQ-x_HI3wEut72cU" {
		t.Errorf("invalid signature:%s", signature)
	}
}

func Test_BuildJoytalkToken(t *testing.T) {
	header := BuildBase64JoytalkJWTHeader(NewJoytalkJWTHeader())
	payload := BuildBase64JoytalkJWTPayload(NewJoytalkJWTPayload(_jwtTestTime, _jwtTestAppId))
	token := BuildJoytalkToken(header, payload, []byte(_jwtTestPass))
	if token != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NTYxODgzOTEsImV4cCI6MTU1NjE4ODk5MSwiYXBwSWQiOiJ4eHh4eHh4eHh4eHgifQ.DFdsMFu_VzPkUUS1eu_Kwyzvc6vSQ-x_HI3wEut72cU" {
		t.Errorf("invalid token:%s", token)
	}
}
