package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestIncr(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/incr/
	// Syntax: INCR key
	cmd := client.Incr(testCtx, "INCRKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("INCR", prefixHook.formatKey("INCRKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/incrby/
	// Syntax: INCRBY key increment
	cmd := client.IncrBy(testCtx, "IncrByKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("INCRBY", prefixHook.formatKey("IncrByKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestIncrByFloat(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/incrbyfloat/
	// Syntax: INCRBYFLOAT key increment
	cmd := client.IncrByFloat(testCtx, "IncrByFloatKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("INCRBYFLOAT", prefixHook.formatKey("IncrByFloatKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/info/
	// Syntax: INFO [section [section ...]]
	cmd := client.Info(testCtx, "redis_version", "redis_git_sha1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("INFO", "redis_version", "redis_git_sha1")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
