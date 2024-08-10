package notice

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReader(t *testing.T) {
	should := require.New(t)
	closeChan := make(chan struct{})
	reader, err := NewReader(
		WithKafkaBrokers([]string{"192.168.91.130:19092"}),
		WithKafkaGroupId("TG1"),
		WithKafkaTopic("kTest"),
		WithHandler(func(msgData []byte) error {
			t.Logf("reseived msg: %s", string(msgData))
			closeChan <- struct{}{}
			return nil
		}),
	)
	should.Nil(err)
	err = reader.Start()
	should.Nil(err)
	<-closeChan
	err = reader.Close()
	should.Nil(err)
}
