package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestSAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sadd/
	// Syntax: SADD key member [member ...]
	cmd := client.SAdd(testCtx, "SAddKey", "member1", "member2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SADD", prefixHook.formatKey("SAddKey"), "member1", "member2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSave(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/save/
	// Syntax: SAVE
	cmd := client.Save(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SAVE")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestScan(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/scan/
	// Syntax: SCAN cursor [MATCH pattern] [COUNT count] [TYPE type]
	cmd := client.Scan(testCtx, 1, "*", 1)
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SCAN", 1, "match", "*", "count", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSCard(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/scard/
	// Syntax: SCARD key
	cmd := client.SCard(testCtx, "SCardKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SCARD", prefixHook.formatKey("SCardKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSCRIPTDEBUG(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/script-debug/
	// Syntax: SCRIPT DEBUG <YES | SYNC | NO>
	t.Errorf("go-redis v9 not supported SCRIPT DEBUG")
}

func TestScriptExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/script-exists/
	// Syntax: SCRIPT EXISTS sha1 [sha1 ...]
	cmd := client.ScriptExists(testCtx, "sha1", "sha2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SCRIPT EXISTS", "sha1", "sha2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestScriptFlush(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/script-flush/
	// Syntax: SCRIPT FLUSH [ASYNC | SYNC]
	cmd := client.ScriptFlush(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SCRIPT FLUSH")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestScriptKill(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/script-kill/
	// Syntax: SCRIPT KILL
	cmd := client.ScriptKill(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SCRIPT KILL")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestScriptLoad(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/script-load/
	// Syntax: SCRIPT LOAD script
	cmd := client.ScriptLoad(testCtx, "return redis.call('set', KEYS[1], ARGV[1])")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SCRIPT LOAD", "return redis.call('set', KEYS[1], ARGV[1])")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSDiff(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sdiff/
	// Syntax: SDIFF key [key ...]
	cmd := client.SDiff(testCtx, "SDiffKey1", "SDiffKey2", "SDiffKey3")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SDIFF", prefixHook.formatKey("SDiffKey1"), prefixHook.formatKey("SDiffKey2"), prefixHook.formatKey("SDiffKey3"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSDiffStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sdiffstore/
	// Syntax: SDIFFSTORE destination key [key ...]
	cmd := client.SDiffStore(testCtx, "SDiffStoreDestination", "SDiffStoreKey1", "SDiffStoreKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SDIFFSTORE", prefixHook.formatKey("SDiffStoreDestination"), prefixHook.formatKey("SDiffStoreKey1"), prefixHook.formatKey("SDiffStoreKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSELECT(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/select/
	// Syntax: SELECT index
	t.Errorf("go-redis v9 not supported SELECT")
}

func TestSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/set/
	// Syntax: SET key value [NX | XX] [GET] [EX seconds | PX milliseconds |  EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]
	cmd := client.Set(testCtx, "SetKey", 1, testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SET", prefixHook.formatKey("SetKey"), 1, "ex", testFormatSec())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSetBit(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/setbit/
	// Syntax: SETBIT key offset value
	cmd := client.SetBit(testCtx, "SetBitKey", 1, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SETBIT", prefixHook.formatKey("SetBitKey"), 1, 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSetEx(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/setex/
	// Syntax: SETEX key seconds value
	cmd := client.SetEx(testCtx, "SetExKey", 1, testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SETEX", prefixHook.formatKey("SetExKey"), 1, testFormatSec())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSetNX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/setnx/
	// Syntax: SETNX key value
	cmd := client.SetNX(testCtx, "SetNXKey", 1, testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SET", prefixHook.formatKey("SetNXKey"), 1, "ex", testFormatSec(), "nx")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSetRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/setrange/
	// Syntax: SETRANGE key offset value
	cmd := client.SetRange(testCtx, "SetRangeKey", 1, "value")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SETRANGE", prefixHook.formatKey("SetRangeKey"), 1, "value")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestShutdown(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/shutdown/
	// Syntax: SHUTDOWN [NOSAVE | SAVE] [NOW] [FORCE] [ABORT]
	//cmd := client.Shutdown(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("SHUTDOWN")
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("SHUTDOWN test not enabled")
}

func TestSInter(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sinter/
	// Syntax: SINTER key [key ...]
	cmd := client.SInter(testCtx, "SInterKey1", "SInterKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SINTER", prefixHook.formatKey("SInterKey1"), prefixHook.formatKey("SInterKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSInterCard(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sintercard/
	// Syntax: SINTERCARD numkeys key [key ...] [LIMIT limit]
	cmd := client.SInterCard(testCtx, 2, "SInterCardKey1", "SInterCardKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SINTERCARD", 2, prefixHook.formatKey("SInterCardKey1"), prefixHook.formatKey("SInterCardKey2"), "limit", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSInterStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sinterstore/
	// Syntax: SINTERSTORE destination key [key ...]
	cmd := client.SInterStore(testCtx, "SInterStoreDestination", "SInterStoreKey1", "SInterStoreKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SINTERSTORE", prefixHook.formatKey("SInterStoreDestination"), prefixHook.formatKey("SInterStoreKey1"), prefixHook.formatKey("SInterStoreKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSIsMember(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sismember/
	// Syntax: SISMEMBER key member
	cmd := client.SIsMember(testCtx, "SIsMemberKey", "member")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SISMEMBER", prefixHook.formatKey("SIsMemberKey"), "member")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSlaveOf(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/slaveof/
	// Syntax: SLAVEOF <host port | NO ONE>
	cmd := client.SlaveOf(testCtx, "127.0.0.1", "6379")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SLAVEOF", "127.0.0.1", "6379")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSlowLogGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/slowlog-get/
	// Syntax: SLOWLOG GET [count]
	cmd := client.SlowLogGet(testCtx, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SLOWLOG GET", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSLOWLOGLEN(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/slowlog-len/
	// Syntax: SLOWLOG LEN
	t.Errorf("go-redis v9 not supported SLOWLOG LEN")
}

func TestSLOWLOGRESET(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/slowlog-reset/
	// Syntax: SLOWLOG RESET
	t.Errorf("go-redis v9 not supported SLOWLOG RESET")
}

func TestSMembers(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/smembers/
	// Syntax: SMEMBERS key
	cmd := client.SMembers(testCtx, "SMembersKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SMEMBERS", prefixHook.formatKey("SMembersKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSMIsMember(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/smismember/
	// Syntax: SMISMEMBER key member [member ...]
	cmd := client.SMIsMember(testCtx, "SMIsMemberKey", "member1", "member2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SMISMEMBER", prefixHook.formatKey("SMIsMemberKey"), "member1", "member2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSMove(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/smove/
	// Syntax: SMOVE source destination member
	cmd := client.SMove(testCtx, "SMoveSource", "SMovedestination", "member")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SMOVE", prefixHook.formatKey("SMoveSource"), prefixHook.formatKey("SMovedestination"), "member")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSort(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sort/
	// Syntax: SORT key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern  ...]] [ASC | DESC] [ALPHA] [STORE destination]
	cmd := client.Sort(testCtx, "SortKey", &redis.Sort{Order: "DESC"})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SORT", prefixHook.formatKey("SortKey"), "DESC")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSortRO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sort_ro/
	// Syntax: SORT_RO key [BY pattern] [LIMIT offset count] [GET pattern [GET  pattern ...]] [ASC | DESC] [ALPHA]
	cmd := client.SortRO(testCtx, "SortROKey", &redis.Sort{Order: "DESC"})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SORT_RO", prefixHook.formatKey("SortROKey"), "DESC")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/spop/
	// Syntax: SPOP key [count]
	cmd := client.SPop(testCtx, "SPopKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SPOP", prefixHook.formatKey("SPopKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSPublish(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/spublish/
	// Syntax: SPUBLISH shardchannel message
	cmd := client.SPublish(testCtx, "SPublishKey", "message")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SPUBLISH", "SPublishKey", "message")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSRandMember(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/srandmember/
	// Syntax: SRANDMEMBER key [count]
	cmd := client.SRandMember(testCtx, "SRandMemberKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SRANDMEMBER", prefixHook.formatKey("SRandMemberKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSRem(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/srem/
	// Syntax: SREM key member [member ...]
	cmd := client.SRem(testCtx, "SRemKey", "member1", "member2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SREM", prefixHook.formatKey("SRemKey"), "member1", "member2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSScan(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sscan/
	// Syntax: SSCAN key cursor [MATCH pattern] [COUNT count]
	cmd := client.SScan(testCtx, "SScanKey", 1, "*", 2)
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SSCAN", prefixHook.formatKey("SScanKey"), 1, "match", "*", "count", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSSubscribe(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ssubscribe/
	// Syntax: SSUBSCRIBE shardchannel [shardchannel ...]
	pubsub := client.SSubscribe(testCtx, "SSubscribeShardchannel1")
	defer pubsub.Close()
	msgi, err := pubsub.ReceiveTimeout(testCtx, testTimeout)
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	subscr := msgi.(*redis.Subscription)
	argsStr := JoinInterfaceSliceBySpan("ssubscribe", subscr.Channel)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SSUBSCRIBE", "SSubscribeShardchannel1")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestStrLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/strlen/
	// Syntax: STRLEN key
	cmd := client.StrLen(testCtx, "StrLenKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("STRLEN", prefixHook.formatKey("StrLenKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSubscribe(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/subscribe/
	// Syntax: SUBSCRIBE channel [channel ...]
	pubsub := client.Subscribe(testCtx, "SubscribeChannel")
	defer pubsub.Close()
	msgi, err := pubsub.ReceiveTimeout(testCtx, testTimeout)
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	subscr := msgi.(*redis.Subscription)
	argsStr := JoinInterfaceSliceBySpan("ssubscribe", subscr.Channel)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SSUBSCRIBE", "SubscribeChannel")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSUBSTR(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/substr/
	// Syntax: SUBSTR key start end
	t.Errorf("go-redis v9 not supported SUBSTR")
}

func TestSUnion(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sunion/
	// Syntax: SUNION key [key ...]
	cmd := client.SUnion(testCtx, "SUnionKey1", "SUnionKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SUNION", prefixHook.formatKey("SUnionKey1"), prefixHook.formatKey("SUnionKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSUnionStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sunionstore/
	// Syntax: SUNIONSTORE destination key [key ...]
	cmd := client.SUnionStore(testCtx, "SUnionStoreDestination", "SUnionStoreKey1", "SUnionStoreKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SUNIONSTORE", prefixHook.formatKey("SUnionStoreDestination"), prefixHook.formatKey("SUnionStoreKey1"), prefixHook.formatKey("SUnionStoreKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestSUNSUBSCRIBE(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sunsubscribe/
	// Syntax: SUNSUBSCRIBE [shardchannel [shardchannel ...]]
	t.Errorf("go-redis v9 not supported SUNSUBSCRIBE")
}

func TestSWAPDB(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/swapdb/
	// Syntax: SWAPDB index1 index2
	t.Errorf("go-redis v9 not supported SWAPDB")
}

func TestSync(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/sync/
	// Syntax: SYNC
	t.Log("Success")
}
