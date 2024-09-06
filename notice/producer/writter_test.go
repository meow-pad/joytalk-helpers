package producer

import (
	"context"
	"github.com/meow-pad/joytalk-helpers/notice/producer/msg"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestWriter(t *testing.T) {
	should := require.New(t)
	writer, err := NewWriter(
		WithKafkaAddress([]string{"192.168.91.130:19092"}),
		WithAsyncWrite(false),
		WithAllowAutoTopicCreation(true),
	)
	should.Nil(err)
	defer func() {
		err = writer.Close()
		should.Nil(err)
	}()
	err = writer.WriteMessage(context.Background(), &msg.ClanStar{
		ClanId:    "123401",
		TimeSec:   time.Now().Unix(),
		StarIncr:  1,
		TotalStar: 10,
	})
	should.Nil(err)
	err = writer.WriteMessage(context.Background(), &msg.UserClanFunds{
		Uid:        "120001",
		TimeSec:    time.Now().Unix(),
		FundsIncr:  30,
		TotalFunds: 600,
	})
	should.Nil(err)
	err = writer.WriteMessage(context.Background(), &msg.ClanIsland{
		ClanId:      "123401",
		TimeSec:     time.Now().Unix(),
		IslandGrade: 4,
	})
	should.Nil(err)
}
