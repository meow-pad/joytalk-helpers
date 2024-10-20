package order

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestOrder_BuildServerKey(t *testing.T) {
	var serverId = BuildServerKey("n6gzdxritkjw")
	t.Logf("serverId: %s, len:%d", serverId, len(serverId))
	serverId = BuildServerKey("54321")
	t.Logf("serverId: %s, len:%d", serverId, len(serverId))
}

func TestOrder_BuildOrderId(t *testing.T) {
	should := require.New(t)
	now := time.Now()
	var serverId = BuildServerKey("n6gzdxritkjw")
	for i := 0; i < 100; i += 2 {
		orderId, err := BuildOrderId(now, serverId, int32(i))
		should.Nil(err)
		should.Equal(true, maxOrderIDLength == len(orderId))
		t.Logf("%d.orderId: %s, len:%d", i, orderId, len(orderId))
	}
	_, err := BuildOrderId(now, serverId, 83949)
	should.NotNil(err)
}
