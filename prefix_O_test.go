package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestObjectEncoding(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/object-encoding/
	// Syntax: OBJECT ENCODING key
	cmd := client.ObjectEncoding(testCtx, "ObjectEncodingKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("OBJECT ENCODING", prefixHook.formatKey("ObjectEncodingKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestObjectFreq(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/object-freq/
	// Syntax: OBJECT FREQ key
	cmd := client.ObjectFreq(testCtx, "ObjectFreqKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("OBJECT FREQ", prefixHook.formatKey("ObjectFreqKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestObjectIdleTime(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/object-idletime/
	// Syntax: OBJECT IDLETIME key
	cmd := client.ObjectIdleTime(testCtx, "ObjectIdleTimeKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("OBJECT IDLETIME", prefixHook.formatKey("ObjectIdleTimeKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestObjectRefCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/object-refcount/
	// Syntax: OBJECT REFCOUNT key
	cmd := client.ObjectRefCount(testCtx, "ObjectRefCountKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("OBJECT REFCOUNT", prefixHook.formatKey("ObjectRefCountKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
