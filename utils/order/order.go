package order

import (
	"fmt"
	"strconv"
	"time"
)

const (
	maxOrderIDLength = 24 // 订单号长度

	timeFormatLayout = "20060102150405"

	MaxIdInSec = 46655 // 既是三个36进制字符 zzz ，为 35 * 36^2 + 35 * 36 + 35 = 46655
)

var (
	ErrIdInSecTooLarge = fmt.Errorf("id in second too large")
)

func BuildServerKey(serverId string) (serverKey [7]byte) {
	copy(serverKey[:], serverId)
	return serverKey
}

func BuildOrderId(orderTime time.Time, serverKey [7]byte, idInSec int32) (orderId string, err error) {
	if idInSec > MaxIdInSec {
		return "", ErrIdInSecTooLarge
	}
	idStr := strconv.FormatUint(uint64(idInSec), 36)
	if len(idStr) < 3 { // 如果不足3位，前面补0；这里的编号只有这么多空位了
		idStr = fmt.Sprintf("%03s", idStr)
	}
	orderId = fmt.Sprintf("%s%s%s", orderTime.Format(timeFormatLayout), serverKey, idStr)
	return orderId, nil
}
