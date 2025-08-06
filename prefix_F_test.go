package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestFAILOVER(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/failover/
	// Syntax: FAILOVER [TO host port [FORCE]] [ABORT] [TIMEOUT milliseconds]
	t.Errorf("go-redis v9 not supported FAILOVER")
}

func TestFCall(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/fcall/
	// Syntax: FCALL function numkeys [key [key ...]] [arg [arg ...]]
	cmd := client.FCall(testCtx, lib1.Functions[0].Name, []string{"my_hash1", "my_hash2"}, "a", 1, "b", 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("FCALL", lib1.Functions[0].Name, 2, prefixHook.formatKey("my_hash1"), prefixHook.formatKey("my_hash2"), "a", 1, "b", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestFCallRo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/fcall_ro/
	// Syntax: FCALL_RO function numkeys [key [key ...]] [arg [arg ...]]
	cmd := client.FCallRo(testCtx, lib1.Functions[0].Name, []string{"my_hash1", "my_hash2"}, "a", 1, "b", 2)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("FCall_Ro", lib1.Functions[0].Name, 2, prefixHook.formatKey("my_hash1"), prefixHook.formatKey("my_hash2"), "a", 1, "b", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestFTSUGADD(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ft.sugadd/
	// Syntax: FT.SUGADD key string score   [INCR]   [PAYLOAD payload]
	t.Errorf("go-redis v9 not supported FT.SUGADD")
}

func TestFTSUGDEL(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ft.sugdel/
	// Syntax: FT.SUGDEL key string
	t.Errorf("go-redis v9 not supported FT.SUGDEL")
}

func TestFTSUGGET(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ft.sugget/
	// Syntax: FT.SUGGET key prefix   [FUZZY]   [WITHSCORES]   [WITHPAYLOADS]   [MAX max]
	t.Errorf("go-redis v9 not supported FT.SUGGET")
}

func TestFTSUGLEN(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ft.suglen/
	// Syntax: FT.SUGLEN key
	t.Errorf("go-redis v9 not supported FT.SUGLEN")
}
