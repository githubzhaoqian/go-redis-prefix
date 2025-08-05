package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestMEMORYDOCTOR(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/memory-doctor/
	// Syntax: MEMORY DOCTOR
	t.Errorf("go-redis v9 not supported MEMORY DOCTOR")
}

func TestMEMORYMALLOCSTATS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/memory-malloc-stats/
	// Syntax: MEMORY MALLOC-STATS
	t.Errorf("go-redis v9 not supported MEMORY MALLOC-STATS")
}

func TestMEMORYPURGE(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/memory-purge/
	// Syntax: MEMORY PURGE
	t.Errorf("go-redis v9 not supported MEMORY PURGE")
}

func TestMEMORYSTATS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/memory-stats/
	// Syntax: MEMORY STATS
	t.Errorf("go-redis v9 not supported MEMORY STATS")
}

func TestMemoryUsage(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/memory-usage/
	// Syntax: MEMORY USAGE key [SAMPLES count]
	cmd := client.MemoryUsage(testCtx, "MemoryUsageKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MEMORY USAGE", prefixHook.formatKey("MemoryUsageKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/mget/
	// Syntax: MGET key [key ...]
	cmd := client.MGet(testCtx, "MGetKey1", "MGetKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MGET", prefixHook.formatKey("MGetKey1"), prefixHook.formatKey("MGetKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMigrate(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/migrate/
	// Syntax: MIGRATE host port <key | ""> destination-db timeout [COPY] [REPLACE]  [AUTH password | AUTH2 username password] [KEYS key [key ...]]
	cmd := client.Migrate(testCtx, "127.0.0.1", "6379", "MigrateKey", 0, testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MIGRATE", "127.0.0.1", "6379", prefixHook.formatKey("MigrateKey"), 0, testFormatMs())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMODULELIST(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/module-list/
	// Syntax: MODULE LIST
	t.Errorf("go-redis v9 not supported MODULE LIST")
}

func TestMODULELOAD(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/module-load/
	// Syntax: MODULE LOAD path [arg [arg ...]]
	t.Errorf("go-redis v9 not supported MODULE LOAD")
}

func TestModuleLoadex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/module-loadex/
	// Syntax: MODULE LOADEX path [CONFIG name value [CONFIG name value ...]]  [ARGS args [args ...]]
	cmd := client.ModuleLoadex(testCtx, &redis.ModuleLoadexConfig{
		Path: "",
		Conf: map[string]interface{}{},
		Args: []interface{}{},
	})
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MODULE LOADEX")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMODULEUNLOAD(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/module-unload/
	// Syntax: MODULE UNLOAD name
	t.Errorf("go-redis v9 not supported MODULE UNLOAD")
}

func TestMonitor(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/monitor/
	// Syntax: MONITOR
	ress := make(chan string)
	cmd := client.Monitor(testCtx, ress)
	args := cmd.Args()
	close(ress)
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MONITOR")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMove(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/move/
	// Syntax: MOVE key db
	cmd := client.Move(testCtx, "MoveKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MOVE", prefixHook.formatKey("MoveKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/mset/
	// Syntax: MSET key value [key value ...]
	cmd := client.MSet(testCtx, "MSetKey1", "val1", "MSetKey2", "val2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MSET", prefixHook.formatKey("MSetKey1"), "val1", prefixHook.formatKey("MSetKey2"), "val2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMSetNX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/msetnx/
	// Syntax: MSETNX key value [key value ...]
	cmd := client.MSetNX(testCtx, "MSetNXKey1", "val1", "MSetNXKey2", "val2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("MSETNX", prefixHook.formatKey("MSetNXKey1"), "val1", prefixHook.formatKey("MSetNXKey2"), "val2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestMULTI(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/multi/
	// Syntax: MULTI
	t.Errorf("go-redis v9 not supported MULTI")
}
