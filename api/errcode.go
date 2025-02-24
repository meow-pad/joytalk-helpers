package api

import (
	"errors"
	"fmt"
)

const (
	ErrCodeSuccess             = 0
	ErrCodeUnknownError        = 1
	ErrCodeInvalidParam        = 6
	ErrCodeLessAuth            = 109 // 权限不足
	ErrCodeAuthFailed          = 110 // 鉴权失败
	ErrCodeInsufficientBalance = 201 // 用户余额不足
	ErrCodeMerchantCantAward   = 202 // 商户可发奖余额不足
	ErrCodeInPaymentBlacklist  = 203 // 用户在支付黑名单中
)

var (
	ErrRespUnknownError        = errors.New("unknown error")
	ErrRespInvalidParam        = errors.New("invalid param")
	ErrRespLessAuth            = errors.New("less auth")
	ErrRespAuthFailed          = errors.New("auth failed")
	ErrRespInsufficientBalance = errors.New("insufficient balance")
	ErrRespMerchantCantAward   = errors.New("merchant cant award anymore")
	ErrRespInPaymentBlacklist  = errors.New("in payment blacklist")
)

func GetRespErr(errCode int32, errMsg string) error {
	switch errCode {
	case ErrCodeSuccess:
		return nil
	case ErrCodeUnknownError:
		return ErrRespUnknownError
	case ErrCodeInvalidParam:
		return ErrRespInvalidParam
	case ErrCodeLessAuth:
		return ErrRespLessAuth
	case ErrCodeAuthFailed:
		return ErrRespAuthFailed
	case ErrCodeInsufficientBalance:
		return ErrRespInsufficientBalance
	case ErrCodeMerchantCantAward:
		return ErrRespMerchantCantAward
	case ErrCodeInPaymentBlacklist:
		return ErrRespInPaymentBlacklist
	default:
		return fmt.Errorf("request failed, bizcode: %d, errMsg:%s", errCode, errMsg)
	}
}
