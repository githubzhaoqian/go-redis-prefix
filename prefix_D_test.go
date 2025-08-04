package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestDBSize(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/dbsize/
	// Syntax: DBSIZE
	cmd := client.DBSize(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("DBSIZE")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestDecr(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/decr/
	// Syntax: DECR key
	cmd := client.Decr(testCtx, "DecrKey1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("DECR", prefixHook.formatKey("DecrKey1"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestDecrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/decrby/
	// Syntax: DECRBY key decrement
	cmd := client.DecrBy(testCtx, "DecrByKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("DECRBY", prefixHook.formatKey("DecrByKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/del/
	// Syntax: DEL key [key ...]
	cmd := client.Del(testCtx, "DelKey1", "DelKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("DEL", prefixHook.formatKey("DelKey1"), prefixHook.formatKey("DelKey2"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

// TODO
func TestDISCARD(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/discard/
	// Syntax: DISCARD
	t.Errorf("go-redis v9 not supported DISCARD")
}

func TestDump(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/dump/
	// Syntax: DUMP key
	cmd := client.Dump(testCtx, "DumpKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("DUMP", prefixHook.formatKey("DumpKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}
