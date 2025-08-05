package redisprefix

import (
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestWait(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/wait/
	// Syntax: WAIT numreplicas timeout
	cmd := client.Wait(testCtx, 1, testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("WAIT", 1, testFormatMs())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestWaitAOF(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/waitaof/
	// Syntax: WAITAOF numlocal numreplicas timeout
	cmd := client.WaitAOF(testCtx, 1, 0, 0)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("WAITAOF", 1, 0, 0)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestWatch(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/watch/
	// Syntax: WATCH key [key ...]
	clientB := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6739",
		DB:   0,
	})
	err := clientB.Ping(testCtx).Err()
	if err != nil {
		panic(err)
	}
	clientB.AddHook(prefixHook)
	done := make(chan struct{})
	key := "WatchKey"
	// 客户端A开启WATCH
	go func() {
		defer close(done)
		err := client.Watch(testCtx, func(tx *redis.Tx) error {
			getCmd := tx.Get(testCtx, key)
			val, _ := getCmd.Int()
			getArgs := getCmd.Args()
			argsStr := JoinInterfaceSliceBySpan(getArgs...)
			t.Log("tx get cmd: ", argsStr)
			target := JoinInterfaceSliceBySpan("GET", prefixHook.formatKey(key))
			if !strings.EqualFold(argsStr, target) {
				t.Fatalf("waning! %s != %s", argsStr, target)
			}
			// 等待B修改
			time.Sleep(3 * time.Second)
			// 事务尝试加1
			_, err := tx.TxPipelined(testCtx, func(pipe redis.Pipeliner) error {
				pipe.Set(testCtx, key, val+1, 0)
				return nil
			})
			return err
		}, key)
		if err == redis.TxFailedErr {
			t.Log("Success conflict")
			return
		} else if err != nil && err != redis.Nil {
			t.Errorf("A Watch err: %v", err)
			return
		}
		t.Errorf("no conflict")
	}()
	// 客户端B在A的WATCH后修改key
	time.Sleep(1 * time.Second)
	setCmd := clientB.Set(testCtx, key, 100, 0)
	setArgs := setCmd.Args()
	argsStr := JoinInterfaceSliceBySpan(setArgs...)
	t.Log("set cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("SET", prefixHook.formatKey(key), 100)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	<-done
}
