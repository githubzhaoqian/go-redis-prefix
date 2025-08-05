package redisprefix

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestTDigestAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.add/
	// Syntax: TDIGEST.ADD key value [value ...]
	client.TDigestCreate(testCtx, "TDigestAddKey")
	cmd := client.TDigestAdd(testCtx, "TDigestAddKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.ADD", prefixHook.formatKey("TDigestAddKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestByRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.byrank/
	// Syntax: TDIGEST.BYRANK key rank [rank ...]
	client.TDigestCreate(testCtx, "TDigestByRankKey")
	cmd := client.TDigestByRank(testCtx, "TDigestByRankKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.BYRANK", prefixHook.formatKey("TDigestByRankKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestByRevRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.byrevrank/
	// Syntax: TDIGEST.BYREVRANK key reverse_rank [reverse_rank ...]
	client.TDigestCreate(testCtx, "TDigestByRevRankKey")
	cmd := client.TDigestByRevRank(testCtx, "TDigestByRevRankKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.BYREVRANK", prefixHook.formatKey("TDigestByRevRankKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestCDF(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.cdf/
	// Syntax: TDIGEST.CDF key value [value ...]
	client.TDigestCreate(testCtx, "TDigestCDFKey")
	cmd := client.TDigestCDF(testCtx, "TDigestCDFKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.CDF", prefixHook.formatKey("TDigestCDFKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestCreate(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.create/
	// Syntax: TDIGEST.CREATE key [COMPRESSION compression]
	cmd := client.TDigestCreate(testCtx, "TDigestCreateKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.CREATE", prefixHook.formatKey("TDigestCreateKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.info/
	// Syntax: TDIGEST.INFO key
	client.TDigestCreate(testCtx, "TDigestInfoKey")
	cmd := client.TDigestInfo(testCtx, "TDigestInfoKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.INFO", prefixHook.formatKey("TDigestInfoKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestMax(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.max/
	// Syntax: TDIGEST.MAX key
	client.TDigestCreate(testCtx, "TDigestMaxKey")
	cmd := client.TDigestMax(testCtx, "TDigestMaxKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.MAX", prefixHook.formatKey("TDigestMaxKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestMerge(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.merge/
	// Syntax: TDIGEST.MERGE destination-key numkeys source-key [source-key ...]  [COMPRESSION compression] [OVERRIDE]
	client.TDigestCreate(testCtx, "TDigestMergeSourceKey1")
	client.TDigestCreate(testCtx, "TDigestMergeSourceKey2")
	cmd := client.TDigestMerge(testCtx, "TDigestMergeDedstKey", &redis.TDigestMergeOptions{}, "TDigestMergeSourceKey1", "TDigestMergeSourceKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.MERGE",
		prefixHook.formatKey("TDigestMergeDedstKey"),
		2,
		prefixHook.formatKey("TDigestMergeSourceKey1"),
		prefixHook.formatKey("TDigestMergeSourceKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestMin(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.min/
	// Syntax: TDIGEST.MIN key
	client.TDigestCreate(testCtx, "TDigestMinKey")
	cmd := client.TDigestMin(testCtx, "TDigestMinKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.MIN", prefixHook.formatKey("TDigestMinKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestQuantile(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.quantile/
	// Syntax: TDIGEST.QUANTILE key quantile [quantile ...]
	client.TDigestCreate(testCtx, "TDigestQuantileKey")
	cmd := client.TDigestQuantile(testCtx, "TDigestQuantileKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.QUANTILE", prefixHook.formatKey("TDigestQuantileKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.rank/
	// Syntax: TDIGEST.RANK key value [value ...]
	client.TDigestCreate(testCtx, "TDigestRankKey")
	cmd := client.TDigestRank(testCtx, "TDigestRankKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.RANK", prefixHook.formatKey("TDigestRankKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestReset(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.reset/
	// Syntax: TDIGEST.RESET key
	client.TDigestCreate(testCtx, "TDigestResetKey")
	cmd := client.TDigestReset(testCtx, "TDigestResetKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.RESET", prefixHook.formatKey("TDigestResetKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestRevRank(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.revrank/
	// Syntax: TDIGEST.REVRANK key value [value ...]
	client.TDigestCreate(testCtx, "TDigestRevRankKey")
	cmd := client.TDigestRevRank(testCtx, "TDigestRevRankKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.REVRANK", prefixHook.formatKey("TDigestRevRankKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTDigestTrimmedMean(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/tdigest.trimmed_mean/
	// Syntax: TDIGEST.TRIMMED_MEAN key low_cut_quantile high_cut_quantile
	client.TDigestCreate(testCtx, "TDigestTrimmedMeanKey1")
	cmd := client.TDigestTrimmedMean(testCtx, "TDigestTrimmedMeanKey1", 0.1, 0.6)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TDIGEST.TRIMMED_MEAN", prefixHook.formatKey("TDigestTrimmedMeanKey1"), 0.1, 0.6)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTime(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/time/
	// Syntax: TIME
	cmd := client.Time(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TIME")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.add/
	// Syntax: TOPK.ADD key items [items ...]
	client.TopKReserve(testCtx, "TopKAddKey", 10)
	cmd := client.TopKAdd(testCtx, "TopKAddKey", "items1", "items2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.ADD", prefixHook.formatKey("TopKAddKey"), "items1", "items2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.count/
	// Syntax: TOPK.COUNT key item [item ...]
	client.TopKReserve(testCtx, "TopKCountKey", 10)
	cmd := client.TopKCount(testCtx, "TopKCountKey", "item1", "item2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.COUNT", prefixHook.formatKey("TopKCountKey"), "item1", "item2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.incrby/
	// Syntax: TOPK.INCRBY key item increment [item increment ...]
	client.TopKReserve(testCtx, "TopKIncrByKey", 10)
	cmd := client.TopKIncrBy(testCtx, "TopKIncrByKey", "item1", 1, "item2", 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.INCRBY", prefixHook.formatKey("TopKIncrByKey"), "item1", 1, "item2", 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.info/
	// Syntax: TOPK.INFO key
	client.TopKReserve(testCtx, "TopKInfoKey", 10)
	cmd := client.TopKInfo(testCtx, "TopKInfoKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.INFO", prefixHook.formatKey("TopKInfoKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKList(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.list/
	// Syntax: TOPK.LIST key [WITHCOUNT]
	client.TopKReserve(testCtx, "TopKListKey", 10)
	cmd := client.TopKList(testCtx, "TopKListKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.LIST", prefixHook.formatKey("TopKListKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKQuery(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.query/
	// Syntax: TOPK.QUERY key item [item ...]
	client.TopKReserve(testCtx, "TopKQueryKey", 10)
	cmd := client.TopKQuery(testCtx, "TopKQueryKey", "item1", "item2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.QUERY", prefixHook.formatKey("TopKQueryKey"), "item1", "item2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTopKReserve(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/topk.reserve/
	// Syntax: TOPK.RESERVE key topk [width depth decay]
	key := fmt.Sprintf("TopKReserveKey%d", time.Now().Unix())
	cmd := client.TopKReserve(testCtx, key, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOPK.RESERVE", prefixHook.formatKey(key), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTouch(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/touch/
	// Syntax: TOUCH key [key ...]
	cmd := client.Touch(testCtx, "TouchKey1", "TouchKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TOUCH", prefixHook.formatKey("TouchKey1"), prefixHook.formatKey("TouchKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.add/
	// Syntax: TS.ADD key timestamp value   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [ON_DUPLICATE policy_ovr]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]   [LABELS [label value ...]]
	cmd := client.TSAdd(testCtx, "TSAddKey", 1000, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.ADD", prefixHook.formatKey("TSAddKey"), 1000, 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSAlter(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.alter/
	// Syntax: TS.ALTER key   [RETENTION retentionPeriod]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]   [LABELS [label value ...]]
	client.TSAdd(testCtx, "TSAlterKey", 1000, 1)
	cmd := client.TSAlter(testCtx, "TSAlterKey", &redis.TSAlterOptions{})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.ALTER", prefixHook.formatKey("TSAlterKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSCreate(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.create/
	// Syntax: TS.CREATE key   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]   [LABELS [label value ...]]
	cmd := client.TSCreate(testCtx, "TSCreateKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.CREATE", prefixHook.formatKey("TSCreateKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSCreateRule(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.createrule/
	// Syntax: TS.CREATERULE sourceKey destKey   AGGREGATION aggregator bucketDuration   [alignTimestamp]
	client.Del(testCtx, "TSCreateRuleDestKey")
	client.TSCreate(testCtx, "TSCreateRuleSourceKey")
	client.TSCreate(testCtx, "TSCreateRuleDestKey")
	cmd := client.TSCreateRule(testCtx, "TSCreateRuleSourceKey", "TSCreateRuleDestKey", 1, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.CREATERULE", prefixHook.formatKey("TSCreateRuleSourceKey"), prefixHook.formatKey("TSCreateRuleDestKey"), "AGGREGATION", "AVG", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}

	t.Log("Success")
}

func TestTSDecrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.decrby/
	// Syntax: TS.DECRBY key subtrahend   [TIMESTAMP timestamp]   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]    [LABELS [label value ...]]
	client.TSCreate(testCtx, "TSDecrByKey")
	cmd := client.TSDecrBy(testCtx, "TSDecrByKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.DECRBY", prefixHook.formatKey("TSDecrByKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.del/
	// Syntax: TS.DEL key fromTimestamp toTimestamp
	client.TSCreate(testCtx, "TSDelKey")
	cmd := client.TSDel(testCtx, "TSDelKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.DEL", prefixHook.formatKey("TSDelKey"), 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSDeleteRule(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.deleterule/
	// Syntax: TS.DELETERULE sourceKey destKey
	client.Del(testCtx, "TSDeleteRuleDestKey")
	client.TSCreate(testCtx, "TSDeleteRuleSourceKey")
	client.TSCreate(testCtx, "TSDeleteRuleDestKey")
	client.TSCreateRule(testCtx, "TSDeleteRuleSourceKey", "TSDeleteRuleDestKey", 1, 1)

	cmd := client.TSDeleteRule(testCtx, "TSDeleteRuleSourceKey", "TSDeleteRuleDestKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.DELETERULE", prefixHook.formatKey("TSDeleteRuleSourceKey"), prefixHook.formatKey("TSDeleteRuleDestKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.get/
	// Syntax: TS.GET key   [LATEST]
	client.TSCreate(testCtx, "TSGetKey")
	cmd := client.TSGet(testCtx, "TSGetKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.GET", prefixHook.formatKey("TSGetKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.incrby/
	// Syntax: TS.INCRBY key addend   [TIMESTAMP timestamp]   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]     [LABELS [label value ...]]
	client.TSCreate(testCtx, "TSIncrByKey")
	cmd := client.TSIncrBy(testCtx, "TSIncrByKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.INCRBY", prefixHook.formatKey("TSIncrByKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.info/
	// Syntax: TS.INFO key   [DEBUG]
	client.TSCreate(testCtx, "TSInfoKey")
	cmd := client.TSInfo(testCtx, "TSInfoKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.INFO", prefixHook.formatKey("TSInfoKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSMAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.madd/
	// Syntax: TS.MADD {key timestamp value}...
	client.Del(testCtx, "TSMAddKey")
	client.TSCreate(testCtx, "TSMAddKey")
	ktvSlices := make([][]interface{}, 2)
	for i := 0; i < 2; i++ {
		ktvSlices[i] = make([]interface{}, 3)
		ktvSlices[i][0] = "TSMAddKey"
		for j := 1; j < 3; j++ {
			ktvSlices[i][j] = (i + j) * j
		}
	}
	cmd := client.TSMAdd(testCtx, ktvSlices)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.MADD", prefixHook.formatKey("TSMAddKey"), 1, 4, prefixHook.formatKey("TSMAddKey"), 2, 6)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSMGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.mget/
	// Syntax: TS.MGET [LATEST] [WITHLABELS | <SELECTED_LABELS label...>] FILTER filterExpr...
	cmd := client.TSMGet(testCtx, []string{"Test=This"})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.MGET", "FILTER", "Test=This")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSMRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.mrange/
	// Syntax: TS.MRANGE fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [WITHLABELS | <SELECTED_LABELS label...>]  [COUNT count]  [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]  FILTER filterExpr...  [GROUPBY label REDUCE reducer]
	cmd := client.TSMRange(testCtx, 1, 2, []string{"Test=This"})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.MRANGE", 1, 2, "FILTER", "Test=This")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSMRevRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.mrevrange/
	// Syntax: TS.MREVRANGE fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [WITHLABELS | <SELECTED_LABELS label...>]  [COUNT count]  [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]  FILTER filterExpr...  [GROUPBY label REDUCE reducer]
	cmd := client.TSMRevRange(testCtx, 1, 2, []string{"Test=This"})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.MREVRANGE", 1, 2, "FILTER", "Test=This")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSQueryIndex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.queryindex/
	// Syntax: TS.QUERYINDEX filterExpr...
	cmd := client.TSQueryIndex(testCtx, []string{"Test=This"})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.QUERYINDEX", "Test=This")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.range/
	// Syntax: TS.RANGE key fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [COUNT count]   [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]
	client.TSCreate(testCtx, "TSRangeKey")
	cmd := client.TSRange(testCtx, "TSRangeKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.RANGE", prefixHook.formatKey("TSRangeKey"), 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTSRevRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ts.revrange/
	// Syntax: TS.REVRANGE key fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [COUNT count]  [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]
	client.TSCreate(testCtx, "TSRevRangeKey")
	cmd := client.TSRevRange(testCtx, "TSRevRangeKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TS.REVRANGE", prefixHook.formatKey("TSRevRangeKey"), 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestTTL(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ttl/
	// Syntax: TTL key
	cmd := client.TTL(testCtx, "TTLKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TTL", prefixHook.formatKey("TTLKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestType(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/type/
	// Syntax: TYPE key
	cmd := client.Type(testCtx, "TypeKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("TYPE", prefixHook.formatKey("TypeKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
