package pay

import (
	"crypto/hmac"
	"crypto/sha256"
)

// BuildSignature
//
//	@Description: 生成签名
func BuildSignature(joytalkId, orderId, amount, uData string, secret []byte) string {
	str := joytalkId + orderId + amount + uData
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(str))
	signature := h.Sum(nil)
	return string(signature)
}

// BuildRefundSignature
//
//	@Description: 生成退款签名
func BuildRefundSignature(joytalkId, orderId, coinAmount, refundAmount string, secret []byte) string {
	str := joytalkId + orderId + coinAmount + refundAmount
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(str))
	signature := h.Sum(nil)
	return string(signature)
}
