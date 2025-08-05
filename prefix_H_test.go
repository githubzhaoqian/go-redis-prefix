package redisprefix

import (
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestHDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hdel/
	// Syntax: HDEL key field [field ...]
	cmd := client.HDel(testCtx, "HDelKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HDEL", prefixHook.formatKey("HDelKey"), "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHELLO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hello/
	// Syntax: HELLO [protover [AUTH username password] [SETNAME clientname]]
	t.Errorf("go-redis v9 not supported HELLO")
}

func TestHExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hexists/
	// Syntax: HEXISTS key field
	cmd := client.HExists(testCtx, "HExistsKey", "field1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HEXISTS", prefixHook.formatKey("HExistsKey"), "field1")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHExpire(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hexpire/
	// Syntax: HEXPIRE key seconds [NX | XX | GT | LT] FIELDS numfields field  [field ...]
	cmd := client.HExpire(testCtx, "HExpireKey", testTimeout, "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HEXPIRE", prefixHook.formatKey("HExpireKey"), testFormatSec(), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHExpireAt(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hexpireat/
	// Syntax: HEXPIREAT key unix-time-seconds [NX | XX | GT | LT] FIELDS numfields  field [field ...]
	tn := time.Now()
	cmd := client.HExpireAt(testCtx, "HExpireAtKey", tn, "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HEXPIREAT", prefixHook.formatKey("HExpireAtKey"), tn.Unix(), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHExpireTime(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hexpiretime/
	// Syntax: HEXPIRETIME key FIELDS numfields field [field ...]
	cmd := client.HExpireTime(testCtx, "HExpireTimeKKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HEXPIRETIME", prefixHook.formatKey("HExpireTimeKKey"), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hget/
	// Syntax: HGET key field
	cmd := client.HGet(testCtx, "HGetKey", "field1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HGET", prefixHook.formatKey("HGetKey"), "field1")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHGetAll(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hgetall/
	// Syntax: HGETALL key
	cmd := client.HGetAll(testCtx, "HGetAllKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HGETALL", prefixHook.formatKey("HGetAllKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHGetDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hgetdel/
	// Syntax: HGETDEL key FIELDS numfields field [field ...]
	cmd := client.HGetDel(testCtx, "HGetDel", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HGETDEL", prefixHook.formatKey("HGetDel"), "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHGetEX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hgetex/
	// Syntax: HGETEX key [EX seconds | PX milliseconds | EXAT unix-time-seconds |  PXAT unix-time-milliseconds | PERSIST] FIELDS numfields field  [field ...]
	cmd := client.HGetEX(testCtx, "HGetEXKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HGETEX", prefixHook.formatKey("HGetEXKey"), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hincrby/
	// Syntax: HINCRBY key field increment
	cmd := client.HIncrBy(testCtx, "HIncrByKey", "field1", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HINCRBY", prefixHook.formatKey("HIncrByKey"), "field1", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHIncrByFloat(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hincrbyfloat/
	// Syntax: HINCRBYFLOAT key field increment
	cmd := client.HIncrByFloat(testCtx, "HIncrByFloatKey", "field1", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HINCRBYFLOAT", prefixHook.formatKey("HIncrByFloatKey"), "field1", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHKeys(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hkeys/
	// Syntax: HKEYS key
	cmd := client.HKeys(testCtx, "HKeysKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HKEYS", prefixHook.formatKey("HKeysKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hlen/
	// Syntax: HLEN key
	cmd := client.HLen(testCtx, "HLenKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HLEN", prefixHook.formatKey("HLenKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHMGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hmget/
	// Syntax: HMGET key field [field ...]
	cmd := client.HMGet(testCtx, "HMGetKey", "field1", "field1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HMGET", prefixHook.formatKey("HMGetKey"), "field1", "field1")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHMSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hmset/
	// Syntax: HMSET key field value [field value ...]
	cmd := client.HMSet(testCtx, "HMSetKey", "field1", 1, "field2", 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HMSET", prefixHook.formatKey("HMSetKey"), "field1", 1, "field2", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHPersist(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hpersist/
	// Syntax: HPERSIST key FIELDS numfields field [field ...]
	cmd := client.HPersist(testCtx, "HPersistKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HPERSIST", prefixHook.formatKey("HPersistKey"), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHPExpire(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hpexpire/
	// Syntax: HPEXPIRE key milliseconds [NX | XX | GT | LT] FIELDS numfields field  [field ...]
	cmd := client.HPExpire(testCtx, "HPExpireKey", testTimeout, "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HPEXPIRE", prefixHook.formatKey("HPExpireKey"), testFormatMs(), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHPExpireAt(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hpexpireat/
	// Syntax: HPEXPIREAT key unix-time-milliseconds [NX | XX | GT | LT]  FIELDS numfields field [field ...]
	tn := time.Now()
	cmd := client.HPExpireAt(testCtx, "HPExpireAtKey", tn, "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HPEXPIREAT", prefixHook.formatKey("HPExpireAtKey"), tn.UnixMilli(), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHPExpireTime(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hpexpiretime/
	// Syntax: HPEXPIRETIME key FIELDS numfields field [field ...]
	cmd := client.HPExpireTime(testCtx, "HPExpireTimeKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HPEXPIRETIME", prefixHook.formatKey("HPExpireTimeKey"), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHPTTL(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hpttl/
	// Syntax: HPTTL key FIELDS numfields field [field ...]
	cmd := client.HPTTL(testCtx, "HPTTLKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HPTTL", prefixHook.formatKey("HPTTLKey"), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHRandField(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hrandfield/
	// Syntax: HRANDFIELD key [count [WITHVALUES]]
	cmd := client.HRandField(testCtx, "HRandFieldKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HRANDFIELD", prefixHook.formatKey("HRandFieldKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHScan(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hscan/
	// Syntax: HSCAN key cursor [MATCH pattern] [COUNT count] [NOVALUES]
	cmd := client.HScan(testCtx, "HScanKey", 1, "*", 1)
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HSCAN", prefixHook.formatKey("HScanKey"), 1, "match", "*", "count", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hset/
	// Syntax: HSET key field value [field value ...]
	cmd := client.HSet(testCtx, "HSetKey", "field1", 1, "field2", 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HSET", prefixHook.formatKey("HSetKey"), "field1", 1, "field2", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHSetEX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hsetex/
	// Syntax: HSETEX key [FNX | FXX] [EX seconds | PX milliseconds |  EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]  FIELDS numfields field value [field value ...]
	cmd := client.HSetEX(testCtx, "HSetEXKey", "field1", "1", "field2", "2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HSETEX", prefixHook.formatKey("HSetEXKey"), "FIELDS", 2, "field1", "1", "field2", "2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHSetNX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hsetnx/
	// Syntax: HSETNX key field value
	cmd := client.HSetNX(testCtx, "HSetNXKey", "field1", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HSETNX", prefixHook.formatKey("HSetNXKey"), "field1", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHStrLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hstrlen/
	// Syntax: HSTRLEN key field
	cmd := client.HStrLen(testCtx, "HStrLenKey", "field1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HSTRLEN", prefixHook.formatKey("HStrLenKey"), "field1")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHTTL(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/httl/
	// Syntax: HTTL key FIELDS numfields field [field ...]
	cmd := client.HTTL(testCtx, "HTTLKey", "field1", "field2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HTTL", prefixHook.formatKey("HTTLKey"), "FIELDS", 2, "field1", "field2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestHVals(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/hvals/
	// Syntax: HVALS key
	cmd := client.HVals(testCtx, "HValsKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("HVALS", prefixHook.formatKey("HValsKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
