package redisprefix

import (
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestPersist(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/persist/
	// Syntax: PERSIST key
	cmd := client.Persist(testCtx, "PersistKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PERSIST", prefixHook.formatKey("PersistKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPExpire(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pexpire/
	// Syntax: PEXPIRE key milliseconds [NX | XX | GT | LT]
	cmd := client.PExpire(testCtx, "PExpireKey", testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PEXPIRE", prefixHook.formatKey("PExpireKey"), testFormatMs())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPExpireAt(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pexpireat/
	// Syntax: PEXPIREAT key unix-time-milliseconds [NX | XX | GT | LT]
	tn := time.Now()
	cmd := client.PExpireAt(testCtx, "PExpireAtKey", tn)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PEXPIREAT", prefixHook.formatKey("PExpireAtKey"), tn.UnixMilli())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPExpireTime(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pexpiretime/
	// Syntax: PEXPIRETIME key
	cmd := client.PExpireTime(testCtx, "PExpireTimeKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PEXPIRETIME", prefixHook.formatKey("PExpireTimeKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPFAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pfadd/
	// Syntax: PFADD key [element [element ...]]
	cmd := client.PFAdd(testCtx, "PFAddKey", "element1", "element2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PFADD", prefixHook.formatKey("PFAddKey"), "element1", "element2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPFCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pfcount/
	// Syntax: PFCOUNT key [key ...]
	cmd := client.PFCount(testCtx, "PFCountKey1", "PFCountKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PFCOUNT", prefixHook.formatKey("PFCountKey1"), prefixHook.formatKey("PFCountKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPFDEBUG(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pfdebug/
	// Syntax: PFDEBUG subcommand key
	t.Errorf("go-redis v9 not supported PFDEBUG")
}

func TestPFMerge(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pfmerge/
	// Syntax: PFMERGE destkey [sourcekey [sourcekey ...]]
	cmd := client.PFMerge(testCtx, "PFMergeDestkey", "PFMergeSourceKey1", "PFMergeSourceKey2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PFMERGE", prefixHook.formatKey("PFMergeDestkey"), prefixHook.formatKey("PFMergeSourceKey1"), prefixHook.formatKey("PFMergeSourceKey2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPFSELFTEST(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pfselftest/
	// Syntax: PFSELFTEST
	t.Errorf("go-redis v9 not supported PFSELFTEST")
}

func TestPing(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/ping/
	// Syntax: PING [message]
	cmd := client.Ping(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PING")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPSETEX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/psetex/
	// Syntax: PSETEX key milliseconds value
	t.Errorf("go-redis v9 not supported PSETEX")
}

func TestPSubscribe(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/psubscribe/
	// Syntax: PSUBSCRIBE pattern [pattern ...]
	pubSub := client.PSubscribe(testCtx, "*")
	defer pubSub.Close()
	r, err := pubSub.ReceiveTimeout(testCtx, time.Second)
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	subscription := r.(*redis.Subscription)
	argsStr := JoinInterfaceSliceBySpan("psubscribe", subscription.Channel)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PSUBSCRIBE", "*")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPSYNC(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/psync/
	// Syntax: PSYNC replicationid offset
	t.Errorf("go-redis v9 not supported PSYNC")
}

func TestPTTL(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pttl/
	// Syntax: PTTL key
	cmd := client.PTTL(testCtx, "PTTLKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PTTL", prefixHook.formatKey("PTTLKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPublish(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/publish/
	// Syntax: PUBLISH channel message
	cmd := client.Publish(testCtx, "Publish", "message")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PUBLISH", "Publish", "message")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPubSubChannels(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pubsub-channels/
	// Syntax: PUBSUB CHANNELS [pattern]
	cmd := client.PubSubChannels(testCtx, "PubSubChannelsKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PUBSUB CHANNELS", "PubSubChannelsKey")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPubSubNumPat(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pubsub-numpat/
	// Syntax: PUBSUB NUMPAT
	cmd := client.PubSubNumPat(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PUBSUB NUMPAT")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPubSubNumSub(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pubsub-numsub/
	// Syntax: PUBSUB NUMSUB [channel [channel ...]]
	cmd := client.PubSubNumSub(testCtx, "PubSubNumSubChannel1", "PubSubNumSubChannel2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PUBSUB NUMSUB", "PubSubNumSubChannel1", "PubSubNumSubChannel2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPubSubShardChannels(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pubsub-shardchannels/
	// Syntax: PUBSUB SHARDCHANNELS [pattern]
	cmd := client.PubSubShardChannels(testCtx, "*K")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PUBSUB SHARDCHANNELS", "*K")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPubSubShardNumSub(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/pubsub-shardnumsub/
	// Syntax: PUBSUB SHARDNUMSUB [shardchannel [shardchannel ...]]
	cmd := client.PubSubShardNumSub(testCtx, "PubSubShardNumSubShardchannel1", "PubSubShardNumSubShardchannel2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("PUBSUB SHARDNUMSUB", "PubSubShardNumSubShardchannel1", "PubSubShardNumSubShardchannel2")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestPUNSUBSCRIBE(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/punsubscribe/
	// Syntax: PUNSUBSCRIBE [pattern [pattern ...]]
	t.Errorf("go-redis v9 not supported PUNSUBSCRIBE")
}
