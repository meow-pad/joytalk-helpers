package diamond

import "fmt"

const (
	PayTypeInvalid = 0
	PayTypeDiamond = 2012 // 钻石
	PayTypeCoin    = 1108 // 游戏币
)

func PayAmountToDiamond(payType int, amount int64) (int64, error) {
	switch payType {
	case PayTypeDiamond:
		//1:1
		return amount, nil
	case PayTypeCoin:
		//10:1
		return amount / 10, nil
	default:
		return 0, fmt.Errorf("unknown pay type %d", payType)
	}
}
