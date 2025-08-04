package redisprefix

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestBFAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.add/
	// Syntax: BF.ADD key item
	cmd := client.BFAdd(testCtx, "BFAddKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.add", prefixHook.formatKey("BFAddKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFCard(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.card/
	// Syntax: BF.CARD key
	cmd := client.BFCard(testCtx, "BFCardKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.card", prefixHook.formatKey("BFCardKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.exists/
	// Syntax: BF.EXISTS key item
	cmd := client.BFExists(testCtx, "BFExistsKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.exists", prefixHook.formatKey("BFExistsKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.info/
	// Syntax: BF.INFO key [CAPACITY | SIZE | FILTERS | ITEMS | EXPANSION]
	client.BFAdd(testCtx, "BFInfoKey", "item1")
	cmd := client.BFInfo(testCtx, "BFInfoKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.info", prefixHook.formatKey("BFInfoKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFInsert(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.insert/
	// Syntax: BF.INSERT key [CAPACITY capacity] [ERROR error]  [EXPANSION expansion] [NOCREATE] [NONSCALING] ITEMS item [item  ...]
	cmd := client.BFInsert(testCtx, "BFInsertKey", nil, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.insert", prefixHook.formatKey("BFInsertKey"), "ITEMS", 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFLoadChunk(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.loadchunk/
	// Syntax: BF.LOADCHUNK key iterator data
	client.BFAdd(testCtx, "BFLoadChunkTest", "item1")
	r, err := client.BFScanDump(testCtx, "BFLoadChunkTest", 0).Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd BFScanDump %v", err)
	}
	bFLoadChunkKey := fmt.Sprintf("BFLoadChunkKey:%d", time.Now().Unix())
	cmd := client.BFLoadChunk(testCtx, bFLoadChunkKey, 1, r.Data)
	_, err = cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.loadchunk", prefixHook.formatKey(bFLoadChunkKey), 1, r.Data)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFMAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.madd/
	// Syntax: BF.MADD key item [item ...]
	cmd := client.BFMAdd(testCtx, "BFMAddKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.madd", prefixHook.formatKey("BFMAddKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFMExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.mexists/
	// Syntax: BF.MEXISTS key item [item ...]
	cmd := client.BFMExists(testCtx, "BFMExistsKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bf.mexists", prefixHook.formatKey("BFMExistsKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFReserve(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.reserve/
	// Syntax: BF.RESERVE key error_rate capacity [EXPANSION expansion]  [NONSCALING]
	cmd := client.BFReserve(testCtx, "BFReserveKey1", 0.01, 100)
	_, err := cmd.Result()
	args := cmd.Args()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v, %v", err, args)
	}
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.RESERVE", prefixHook.formatKey("BFReserveKey1"), 0.01, 1, 100)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFScanDump(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bf.scandump/
	// Syntax: BF.SCANDUMP key iterator
	client.BFAdd(testCtx, "BFScanDumpKey", "item1")
	cmd := client.BFScanDump(testCtx, "BFScanDumpKey", 0)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.SCANDUMP", prefixHook.formatKey("BFScanDumpKey"), 0)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBgRewriteAOF(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bgrewriteaof/
	// Syntax: BGREWRITEAOF
	cmd := client.BgRewriteAOF(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bgrewriteaof")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBgSave(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bgsave/
	// Syntax: BGSAVE [SCHEDULE]
	cmd := client.BgSave(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bgsave")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBitCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bitcount/
	// Syntax: BITCOUNT key [start end [BYTE | BIT]]
	cmd := client.BitCount(testCtx, "BitCountKey", &redis.BitCount{Start: 1, End: 2})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bitcount", prefixHook.formatKey("BitCountKey"), 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBitField(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bitfield/
	// Syntax: BITFIELD key [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>]  <SET encoding offset value | INCRBY encoding offset increment>  [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>]  <SET encoding offset value | INCRBY encoding offset increment>  ...]]
	cmd := client.BitField(testCtx, "BitFieldKey", "SET", "i8", "#0", "100", "SET", "i8", "#1", "200")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bitfield", prefixHook.formatKey("BitFieldKey"), "SET", "i8", "#0", "100", "SET", "i8", "#1", "200")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBITFIELD_RO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bitfield_ro/
	// Syntax: BITFIELD_RO key [GET encoding offset [GET encoding offset ...]]
	t.Errorf("go-redis v9 not supported BITFIELD_RO")
}

func TestBITOP(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bitop/
	// Syntax: BITOP <AND | OR | XOR | NOT> destkey key [key ...]
	t.Errorf("go-redis v9 not supported BITOP")
}

func TestBitPos(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bitpos/
	// Syntax: BITPOS key bit [start [end [BYTE | BIT]]]
	cmd := client.BitPos(testCtx, "BitPos", 0, []int64{1}...)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bitpos", prefixHook.formatKey("BitPos"), 0, 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBLMove(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/blmove/
	// Syntax: BLMOVE source destination <LEFT | RIGHT> <LEFT | RIGHT> timeout
	cmd := client.BLMove(testCtx, "BLMoveSource", "BLMoveDestination", "LEFT", "RIGHT", testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("blmove", prefixHook.formatKey("BLMoveSource"), prefixHook.formatKey("BLMoveDestination"), "LEFT", "RIGHT", testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBLMPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/blmpop/
	// Syntax: BLMPOP timeout numkeys key [key ...] <LEFT | RIGHT> [COUNT count]
	cmd := client.BLMPop(testCtx, testTimeout, "LEFT", 1, "BLMPopKey1", "BLMPopKey2")
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("blmpop", 1, 2, prefixHook.formatKey("BLMPopKey1"), prefixHook.formatKey("BLMPopKey2"), "left", "count", testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBLPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/blpop/
	// Syntax: BLPOP key [key ...] timeout
	cmd := client.BLPop(testCtx, testTimeout, "BLPopKey1", "BLPopKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("blpop", prefixHook.formatKey("BLPopKey1"), prefixHook.formatKey("BLPopKey2"), testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBRPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/brpop/
	// Syntax: BRPOP key [key ...] timeout
	cmd := client.BRPop(testCtx, testTimeout, "BRPopKey1", "BRPopKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("brpop", prefixHook.formatKey("BRPopKey1"), prefixHook.formatKey("BRPopKey2"), testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBRPopLPush(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/brpoplpush/
	// Syntax: BRPOPLPUSH source destination timeout
	cmd := client.BRPopLPush(testCtx, "BRPopLPushSource", "BRPopLPushDestination", testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("brpoplpush", prefixHook.formatKey("BRPopLPushSource"), prefixHook.formatKey("BRPopLPushDestination"), testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBZMPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bzmpop/
	// Syntax: BZMPOP timeout numkeys key [key ...] <MIN | MAX> [COUNT count]
	cmd := client.BZMPop(testCtx, testTimeout, "max", 2, "BZMPopKey1", "BZMPopKey2")
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bzmpop", testFormatSec(), 2, prefixHook.formatKey("BZMPopKey1"), prefixHook.formatKey("BZMPopKey2"), "max", "count", 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBZPopMax(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bzpopmax/
	// Syntax: BZPOPMAX key [key ...] timeout
	cmd := client.BZPopMax(testCtx, testTimeout, "BZPopMax1", "BZPopMax2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bzpopmax", prefixHook.formatKey("BZPopMax1"), prefixHook.formatKey("BZPopMax2"), testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBZPopMin(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/bzpopmin/
	// Syntax: BZPOPMIN key [key ...] timeout
	cmd := client.BZPopMin(testCtx, testTimeout, "BZPopMin1", "BZPopMin2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("bzpopmin", prefixHook.formatKey("BZPopMin1"), prefixHook.formatKey("BZPopMin2"), testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}
