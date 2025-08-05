package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestUnlink(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/unlink/
	// Syntax: UNLINK key [key ...]
	cmd := client.Unlink(testCtx, "UnlinkKey1", "UnlinkKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("UNLINK", prefixHook.formatKey("UnlinkKey1"), prefixHook.formatKey("UnlinkKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestUNSUBSCRIBE(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/unsubscribe/
	// Syntax: UNSUBSCRIBE [channel [channel ...]]
	t.Errorf("go-redis v9 not supported UNSUBSCRIBE")
}

func TestUNWATCH(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/unwatch/
	// Syntax: UNWATCH
	t.Errorf("go-redis v9 not supported UNWATCH")
}
