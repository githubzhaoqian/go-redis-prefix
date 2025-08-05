package redisprefix

import (
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestEcho(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/echo/
	// Syntax: ECHO message
	cmd := client.Echo(testCtx, "message")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("ECHO", "message")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestEval(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/eval/
	// Syntax: EVAL script numkeys [key [key ...]] [arg [arg ...]]
	cmd := client.Eval(testCtx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EVAL", "return {KEYS[1],ARGV[1]}", 1, "key", "hello")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestEvalSha(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/evalsha/
	// Syntax: EVALSHA sha1 numkeys [key [key ...]] [arg [arg ...]]
	sha1, err := client.ScriptLoad(testCtx, "return {KEYS[1],ARGV[1]}").Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd ScriptLoad %v", err)
	}
	cmd := client.EvalSha(testCtx, sha1, []string{"key"}, "hello")
	_, err = cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EVALSHA", sha1, 1, "key", "hello")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestEVALSHA_RO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/evalsha_ro/
	// Syntax: EVALSHA_RO sha1 numkeys [key [key ...]] [arg [arg ...]]
	cmd := client.EvalShaRO(testCtx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EVALSHA_RO", "return {KEYS[1],ARGV[1]}", 1, "key", "hello")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestEVAL_RO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/eval_ro/
	// Syntax: EVAL_RO script numkeys [key [key ...]] [arg [arg ...]]
	cmd := client.EvalRO(testCtx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EVAL_RO", "return {KEYS[1],ARGV[1]}", 1, "key", "hello")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestEXEC(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/exec/
	// Syntax: EXEC
	t.Errorf("go-redis v9 not supported EXEC")
}

func TestExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/exists/
	// Syntax: EXISTS key [key ...]
	cmd := client.Exists(testCtx, "ExistsKey1", "ExistsKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EXISTS", prefixHook.formatKey("ExistsKey1"), prefixHook.formatKey("ExistsKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestExpire(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/expire/
	// Syntax: EXPIRE key seconds [NX | XX | GT | LT]
	cmd := client.Expire(testCtx, "ExpireKey", testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EXPIRE", prefixHook.formatKey("ExpireKey"), testFormatSec())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestExpireAt(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/expireat/
	// Syntax: EXPIREAT key unix-time-seconds [NX | XX | GT | LT]
	tn := time.Now()
	cmd := client.ExpireAt(testCtx, "ExpireAtKey", tn)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EXPIREAT", prefixHook.formatKey("ExpireAtKey"), tn.Unix())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestExpireTime(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/expiretime/
	// Syntax: EXPIRETIME key
	cmd := client.ExpireTime(testCtx, "ExpireTimeKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("EXPIRETIME", prefixHook.formatKey("ExpireTimeKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
