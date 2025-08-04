package redisprefix

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

type commandData struct {
	Command        string
	Url            string
	Syntax         string
	AvailableSince string
}

var (
	client      *redis.Client
	prefix      = "go-prefix"
	prefixHook  = NewKeyPrefixHook(prefix)
	testCtx     = context.Background()
	testTimeout = 1 * time.Second
)

func TestMain(m *testing.M) {
	// redis:8.0.3
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6739",
		DB:   0,
	})
	err := client.Ping(testCtx).Err()
	if err != nil {
		panic(err)
	}
	client.AddHook(prefixHook)
	os.Exit(m.Run())
}

func TestFilterPrefixNone(t *testing.T) {
	FilterPrefixNone()
	for k, v := range CommandPrefixType {
		if v == PrefixNone {
			t.Fatalf("command: %s is PrefixNone", k)
		}
	}
	t.Log("END")
}

// TestJoinInterfaceSlice join interfaceSlice
func TestJoinInterfaceSlice(t *testing.T) {
	data := []interface{}{"1", "2", nil, 4}
	t.Log(JoinInterfaceSliceBySpan(data))
}

func testFormatSec() int64 {
	return formatSec(testTimeout)
}
