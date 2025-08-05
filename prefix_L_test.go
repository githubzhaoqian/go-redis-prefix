package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestLastSave(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lastsave/
	// Syntax: LASTSAVE
	cmd := client.LastSave(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LASTSAVE")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLATENCYDOCTOR(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/latency-doctor/
	// Syntax: LATENCY DOCTOR
	t.Errorf("go-redis v9 not supported LATENCY DOCTOR")
}

func TestLATENCYGRAPH(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/latency-graph/
	// Syntax: LATENCY GRAPH event
	t.Errorf("go-redis v9 not supported LATENCY GRAPH")
}

func TestLATENCYHISTOGRAM(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/latency-histogram/
	// Syntax: LATENCY HISTOGRAM [command [command ...]]
	t.Errorf("go-redis v9 not supported LATENCY HISTOGRAM")
}

func TestLATENCYHISTORY(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/latency-history/
	// Syntax: LATENCY HISTORY event
	t.Errorf("go-redis v9 not supported LATENCY HISTORY")
}

func TestLATENCYLATEST(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/latency-latest/
	// Syntax: LATENCY LATEST
	t.Errorf("go-redis v9 not supported LATENCY LATEST")
}

func TestLATENCYRESET(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/latency-reset/
	// Syntax: LATENCY RESET [event [event ...]]
	t.Errorf("go-redis v9 not supported LATENCY RESET")
}

func TestLCS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lcs/
	// Syntax: LCS key1 key2 [LEN] [IDX] [MINMATCHLEN min-match-len] [WITHMATCHLEN]
	cmd := client.LCS(testCtx, &redis.LCSQuery{
		Key1:         "LCSKey1",
		Key2:         "LCSKey2",
		Len:          true,
		Idx:          true,
		MinMatchLen:  1,
		WithMatchLen: true,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LCS", prefixHook.formatKey("LCSKey1"), prefixHook.formatKey("LCSKey2"), "len")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLIndex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lindex/
	// Syntax: LINDEX key index
	cmd := client.LIndex(testCtx, "LIndexKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LINDEX", prefixHook.formatKey("LIndexKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLInsert(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/linsert/
	// Syntax: LINSERT key <BEFORE | AFTER> pivot element
	cmd := client.LInsert(testCtx, "LInsertKey", "AFTER", "pivot", "element")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LINSERT", prefixHook.formatKey("LInsertKey"), "AFTER", "pivot", "element")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/llen/
	// Syntax: LLEN key
	cmd := client.LLen(testCtx, "LLenKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LLEN", prefixHook.formatKey("LLenKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLMove(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lmove/
	// Syntax: LMOVE source destination <LEFT | RIGHT> <LEFT | RIGHT>
	cmd := client.LMove(testCtx, "LMoveSource", "LMoveDestination", "RIGHT", "LEFT")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LMOVE", prefixHook.formatKey("LMoveSource"), prefixHook.formatKey("LMoveDestination"), "RIGHT", "LEFT")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLMPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lmpop/
	// Syntax: LMPOP numkeys key [key ...] <LEFT | RIGHT> [COUNT count]
	cmd := client.LMPop(testCtx, "RIGHT", 2, "LMPopKey1", "LMPopKey2")
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LMPOP", 2, prefixHook.formatKey("LMPopKey1"), prefixHook.formatKey("LMPopKey2"), "right", "count", 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLOLWUT(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lolwut/
	// Syntax: LOLWUT [VERSION version]
	t.Errorf("go-redis v9 not supported LOLWUT")
}

func TestLPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lpop/
	// Syntax: LPOP key [count]
	cmd := client.LPop(testCtx, "LPopKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LPOP", prefixHook.formatKey("LPopKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLPos(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lpos/
	// Syntax: LPOS key element [RANK rank] [COUNT num-matches] [MAXLEN len]
	cmd := client.LPos(testCtx, "LPosKey", "val", redis.LPosArgs{
		Rank: 1, MaxLen: 1,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LPOS", prefixHook.formatKey("LPosKey"), "val", "rank", 1, "maxlen", 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLPush(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lpush/
	// Syntax: LPUSH key element [element ...]
	cmd := client.LPush(testCtx, "LPushKey", "element1", "element2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LPUSH", prefixHook.formatKey("LPushKey"), "element1", "element2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLPushX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lpushx/
	// Syntax: LPUSHX key element [element ...]
	cmd := client.LPushX(testCtx, "LPushXKey", "element1", "element2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LPUSHX", prefixHook.formatKey("LPushXKey"), "element1", "element2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lrange/
	// Syntax: LRANGE key start stop
	cmd := client.LRange(testCtx, "LRangeKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LRANGE", prefixHook.formatKey("LRangeKey"), 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLRem(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lrem/
	// Syntax: LREM key count element
	cmd := client.LRem(testCtx, "LRemKey", 1, "element")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LREM", prefixHook.formatKey("LRemKey"), 1, "element")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/lset/
	// Syntax: LSET key index element
	client.LPush(testCtx, "LSetKey", "element1", "element2")
	cmd := client.LSet(testCtx, "LSetKey", 1, "setElement")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LSET", prefixHook.formatKey("LSetKey"), 1, "setElement")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestLTrim(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ltrim/
	// Syntax: LTRIM key start stop
	cmd := client.LTrim(testCtx, "LTrimKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("LTRIM", prefixHook.formatKey("LTrimKey"), 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}
