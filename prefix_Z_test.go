package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestZAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zadd/
	// Syntax: ZADD key [NX | XX] [GT | LT] [CH] [INCR] score member [score member  ...]
	cmd := client.ZAdd(testCtx, "ZAddKey", redis.Z{
		Score:  1,
		Member: "Member1",
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZADD", prefixHook.formatKey("ZAddKey"), 1, "Member1")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZCard(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zcard/
	// Syntax: ZCARD key
	cmd := client.ZCard(testCtx, "ZCardKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZCARD", prefixHook.formatKey("ZCardKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zcount/
	// Syntax: ZCOUNT key min max
	cmd := client.ZCount(testCtx, "ZCountKey", "1", "2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZCOUNT", prefixHook.formatKey("ZCountKey"), "1", "2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZDiff(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zdiff/
	// Syntax: ZDIFF numkeys key [key ...] [WITHSCORES]
	cmd := client.ZDiff(testCtx, "ZDiffKey1", "ZDiffKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZDIFF", 2, prefixHook.formatKey("ZDiffKey1"), prefixHook.formatKey("ZDiffKey2"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZDiffStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zdiffstore/
	// Syntax: ZDIFFSTORE destination numkeys key [key ...]
	cmd := client.ZDiffStore(testCtx, "ZDiffStoreDestination", "ZDiffStoreKey1", "ZDiffStoreKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZDIFFSTORE", prefixHook.formatKey("ZDiffStoreDestination"), 2, prefixHook.formatKey("ZDiffStoreKey1"), prefixHook.formatKey("ZDiffStoreKey2"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zincrby/
	// Syntax: ZINCRBY key increment member
	cmd := client.ZIncrBy(testCtx, "ZIncrByKey", 1, "member")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZINCRBY", prefixHook.formatKey("ZIncrByKey"), 1, "member")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZInter(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zinter/
	// Syntax: ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]]  [AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
	cmd := client.ZInter(testCtx, &redis.ZStore{
		Keys:    []string{"ZInterKey1", "ZInterKey2"},
		Weights: []float64{1, 2},
		// Can be SUM, MIN or MAX.
		Aggregate: "SUM",
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZINTER", 2, prefixHook.formatKey("ZInterKey1"), prefixHook.formatKey("ZInterKey2"), "WEIGHTS", 1, 2, "AGGREGATE", "SUM")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZInterCard(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zintercard/
	// Syntax: ZINTERCARD numkeys key [key ...] [LIMIT limit]
	cmd := client.ZInterCard(testCtx, 2, "ZInterCardKey1", "ZInterCardKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZINTERCARD", 2, prefixHook.formatKey("ZInterCardKey1"), prefixHook.formatKey("ZInterCardKey2"), "LIMIT", 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZInterStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zinterstore/
	// Syntax: ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight  [weight ...]] [AGGREGATE <SUM | MIN | MAX>]
	cmd := client.ZInterStore(testCtx, "ZInterStoreDestination", &redis.ZStore{
		Keys:    []string{"ZInterStoreKey1", "ZInterStoreKey2"},
		Weights: []float64{1, 2},
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZINTERSTORE", prefixHook.formatKey("ZInterStoreDestination"), 2, prefixHook.formatKey("ZInterStoreKey1"), prefixHook.formatKey("ZInterStoreKey2"), "weights", 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZLexCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zlexcount/
	// Syntax: ZLEXCOUNT key min max
	cmd := client.ZLexCount(testCtx, "ZLexCountKey", "-", "+")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZLEXCOUNT", prefixHook.formatKey("ZLexCountKey"), "-", "+")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZMPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zmpop/
	// Syntax: ZMPOP numkeys key [key ...] <MIN | MAX> [COUNT count]
	cmd := client.ZMPop(testCtx, "MIN", 1, "ZMPopKey1", "ZMPopKey2")
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZMPOP", 2, prefixHook.formatKey("ZMPopKey1"), prefixHook.formatKey("ZMPopKey2"), "min", "count", 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZMScore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zmscore/
	// Syntax: ZMSCORE key member [member ...]
	cmd := client.ZMScore(testCtx, "ZMScoreKey", "member1", "member2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZMSCORE", prefixHook.formatKey("ZMScoreKey"), "member1", "member2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZPopMax(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zpopmax/
	// Syntax: ZPOPMAX key [count]
	cmd := client.ZPopMax(testCtx, "ZPopMaxKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZPOPMAX", prefixHook.formatKey("ZPopMaxKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZPopMin(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zpopmin/
	// Syntax: ZPOPMIN key [count]
	cmd := client.ZPopMin(testCtx, "ZPopMinKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZPOPMIN", prefixHook.formatKey("ZPopMinKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRandMember(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrandmember/
	// Syntax: ZRANDMEMBER key [count [WITHSCORES]]
	cmd := client.ZRandMember(testCtx, "ZRandMemberKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZRANDMEMBER", prefixHook.formatKey("ZRandMemberKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrange/
	// Syntax: ZRANGE key start stop [BYSCORE | BYLEX] [REV] [LIMIT offset count]  [WITHSCORES]
	cmd := client.ZRange(testCtx, "ZRangeKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZRANGE", prefixHook.formatKey("ZRangeKey"), 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRangeByLex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrangebylex/
	// Syntax: ZRANGEBYLEX key min max [LIMIT offset count]
	cmd := client.ZRangeByLex(testCtx, "ZRangeByLexKey", &redis.ZRangeBy{
		Min: "-", Max: "+",
		Offset: 1, Count: 1,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZRANGEBYLEX", prefixHook.formatKey("ZRangeByLexKey"), "-", "+", "limit", 1, 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRangeByScore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrangebyscore/
	// Syntax: ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
	cmd := client.ZRangeByScore(testCtx, "ZRangeByScoreKey", &redis.ZRangeBy{
		Min: "1", Max: "2",
		Offset: 1, Count: 1,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZRANGEBYSCORE", prefixHook.formatKey("ZRangeByScoreKey"), 1, 2, "LIMIT", 1, 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRangeStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrangestore/
	// Syntax: ZRANGESTORE dst src min max [BYSCORE | BYLEX] [REV] [LIMIT offset  count]
	cmd := client.ZRangeStore(testCtx, "ZRangeStoreDst", redis.ZRangeArgs{
		Key:     "ZRangeStoreSrc",
		Start:   "(3",
		Stop:    8,
		ByScore: true,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZRANGESTORE", prefixHook.formatKey("ZRangeStoreDst"), prefixHook.formatKey("ZRangeStoreSrc"), "(3", 8, "byscore")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrank/
	// Syntax: ZRANK key member [WITHSCORE]
	cmd := client.ZRank(testCtx, "ZRankKey", "member")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZRANK", prefixHook.formatKey("ZRankKey"), "member")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRem(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrem/
	// Syntax: ZREM key member [member ...]
	cmd := client.ZRem(testCtx, "ZRemKey", "member1", "member2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREM", prefixHook.formatKey("ZRemKey"), "member1", "member2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRemRangeByLex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zremrangebylex/
	// Syntax: ZREMRANGEBYLEX key min max
	cmd := client.ZRemRangeByLex(testCtx, "ZRemRangeByLexKey", "-", "+")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREMRANGEBYLEX", prefixHook.formatKey("ZRemRangeByLexKey"), "-", "+")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRemRangeByRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zremrangebyrank/
	// Syntax: ZREMRANGEBYRANK key start stop
	cmd := client.ZRemRangeByRank(testCtx, "ZRemRangeByRankKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREMRANGEBYRANK", prefixHook.formatKey("ZRemRangeByRankKey"), 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRemRangeByScore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zremrangebyscore/
	// Syntax: ZREMRANGEBYSCORE key min max
	cmd := client.ZRemRangeByScore(testCtx, "ZRemRangeByScoreKey", "1", "2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREMRANGEBYSCORE", prefixHook.formatKey("ZRemRangeByScoreKey"), "1", "2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRevRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrevrange/
	// Syntax: ZREVRANGE key start stop [WITHSCORES]
	cmd := client.ZRevRange(testCtx, "ZRevRangeKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREVRANGE", prefixHook.formatKey("ZRevRangeKey"), 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRevRangeByLex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrevrangebylex/
	// Syntax: ZREVRANGEBYLEX key max min [LIMIT offset count]
	cmd := client.ZRevRangeByLex(testCtx, "ZRevRangeByLexKey", &redis.ZRangeBy{
		Min: "-", Max: "+",
		Offset: 1, Count: 1,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREVRANGEBYLEX", prefixHook.formatKey("ZRevRangeByLexKey"), "+", "-", "limit", 1, 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRevRangeByScore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrevrangebyscore/
	// Syntax: ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
	cmd := client.ZRevRangeByScore(testCtx, "ZRevRangeByScoreKey", &redis.ZRangeBy{
		Min: "1", Max: "2",
		Offset: 1, Count: 1,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREVRANGEBYSCORE", prefixHook.formatKey("ZRevRangeByScoreKey"), 2, 1, "limit", 1, 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZRevRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zrevrank/
	// Syntax: ZREVRANK key member [WITHSCORE]
	cmd := client.ZRevRank(testCtx, "ZRevRankKey", "member")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZREVRANK", prefixHook.formatKey("ZRevRankKey"), "member")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZScan(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zscan/
	// Syntax: ZSCAN key cursor [MATCH pattern] [COUNT count]
	cmd := client.ZScan(testCtx, "ZScanKey", 1, "*", 1)
	_, _, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZSCAN", prefixHook.formatKey("ZScanKey"), 1, "match", "*", "count", 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZScore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zscore/
	// Syntax: ZSCORE key member
	cmd := client.ZScore(testCtx, "ZScoreKey", "member")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZSCORE", prefixHook.formatKey("ZScoreKey"), "member")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZUnion(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zunion/
	// Syntax: ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]]  [AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
	cmd := client.ZUnion(testCtx, redis.ZStore{
		Keys:    []string{"ZUnionKey1", "ZUnionKey2"},
		Weights: []float64{1, 2},
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZUNION", 2, prefixHook.formatKey("ZUnionKey1"), prefixHook.formatKey("ZUnionKey2"), "WEIGHTS", 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestZUnionStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/zunionstore/
	// Syntax: ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight  [weight ...]] [AGGREGATE <SUM | MIN | MAX>]
	cmd := client.ZUnionStore(testCtx, "ZUnionStoreDestination", &redis.ZStore{
		Keys:    []string{"ZUnionStoreKey1", "ZUnionStore2"},
		Weights: []float64{1, 2},
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("ZUNIONSTORE", prefixHook.formatKey("ZUnionStoreDestination"), 2, prefixHook.formatKey("ZUnionStoreKey1"), prefixHook.formatKey("ZUnionStore2"), "WEIGHTS", 1, 2)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}
