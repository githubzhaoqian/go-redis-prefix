package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestRandomKey(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/randomkey/
	// Syntax: RANDOMKEY
	cmd := client.RandomKey(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RANDOMKEY")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestReadOnly(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/readonly/
	// Syntax: READONLY
	cmd := client.ReadOnly(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("READONLY")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestReadWrite(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/readwrite/
	// Syntax: READWRITE
	cmd := client.ReadWrite(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("READWRITE")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestRename(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/rename/
	// Syntax: RENAME key newkey
	client.Set(testCtx, "RenameKey1", 1, 0)
	cmd := client.Rename(testCtx, "RenameKey1", "RenameKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RENAME", prefixHook.formatKey("RenameKey1"), prefixHook.formatKey("RenameKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestRenameNX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/renamenx/
	// Syntax: RENAMENX key newkey
	client.Set(testCtx, "RenameNXKey1", 1, 0)
	cmd := client.RenameNX(testCtx, "RenameNXKey1", "RenameNXKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RENAMENX", prefixHook.formatKey("RenameNXKey1"), prefixHook.formatKey("RenameNXKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestREPLCONF(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/replconf/
	// Syntax: REPLCONF
	t.Errorf("go-redis v9 not supported REPLCONF")
}

func TestREPLICAOF(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/replicaof/
	// Syntax: REPLICAOF <host port | NO ONE>
	t.Errorf("go-redis v9 not supported REPLICAOF")
}

func TestRESET(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/reset/
	// Syntax: RESET
	t.Errorf("go-redis v9 not supported RESET")
}

func TestRestore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/restore/
	// Syntax: RESTORE key ttl serialized-value [REPLACE] [ABSTTL]  [IDLETIME seconds] [FREQ frequency]
	key := "RestoreKey"
	client.Del(testCtx, key)
	client.Set(testCtx, "DumpKeySet", 1, 0)
	r, _ := client.Dump(testCtx, "DumpKeySet").Result()
	cmd := client.Restore(testCtx, key, testTimeout, r)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RESTORE", prefixHook.formatKey(key), testFormatMs(), r)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestRESTOREASKING(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/restore-asking/
	// Syntax: RESTORE-ASKING key ttl serialized-value [REPLACE] [ABSTTL]  [IDLETIME seconds] [FREQ frequency]
	t.Errorf("go-redis v9 not supported RESTORE-ASKING")
}

func TestROLE(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/role/
	// Syntax: ROLE
	t.Errorf("go-redis v9 not supported ROLE")
}

func TestRPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/rpop/
	// Syntax: RPOP key [count]
	cmd := client.RPop(testCtx, "RPopKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RPOP", prefixHook.formatKey("RPopKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestRPopLPush(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/rpoplpush/
	// Syntax: RPOPLPUSH source destination
	cmd := client.RPopLPush(testCtx, "RPopLPushSource", "RPopLPushDestination")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RPOPLPUSH", prefixHook.formatKey("RPopLPushSource"), prefixHook.formatKey("RPopLPushDestination"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestRPush(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/rpush/
	// Syntax: RPUSH key element [element ...]
	cmd := client.RPush(testCtx, "RPushKey", "element1", "element2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RPUSH", prefixHook.formatKey("RPushKey"), "element1", "element2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestRPushX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/rpushx/
	// Syntax: RPUSHX key element [element ...]
	cmd := client.RPushX(testCtx, "RPushXKey", "element1", "element2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("RPUSHX", prefixHook.formatKey("RPushXKey"), "element1", "element2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
