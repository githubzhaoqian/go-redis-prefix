package redisprefix

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestCFAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.add/
	// Syntax: CF.ADD key item
	cmd := client.CFAdd(testCtx, "CFAddKey", "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.ADD", prefixHook.formatKey("CFAddKey"), "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFAddNX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.addnx/
	// Syntax: CF.ADDNX key item
	client.CFInsert(testCtx, "CFAddKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	cmd := client.CFAddNX(testCtx, "CFAddKey", "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.ADDNX", prefixHook.formatKey("CFAddKey"), "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFCount(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.count/
	// Syntax: CF.COUNT key item
	client.CFInsert(testCtx, "CFInsertKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	cmd := client.CFCount(testCtx, "CFInsertKey", "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.COUNT", prefixHook.formatKey("CFInsertKey"), "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.del/
	// Syntax: CF.DEL key item
	client.CFInsert(testCtx, "CFInsertKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	cmd := client.CFDel(testCtx, "CFInsertKey", "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.DEL", prefixHook.formatKey("CFInsertKey"), "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.exists/
	// Syntax: CF.EXISTS key item
	cmd := client.CFExists(testCtx, "CFExistsKey", "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.EXISTS", prefixHook.formatKey("CFExistsKey"), "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.info/
	// Syntax: CF.INFO key
	client.CFInsert(testCtx, "CFInsertKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	cmd := client.CFInfo(testCtx, "CFInsertKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.INFO", prefixHook.formatKey("CFInsertKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFInsert(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.insert/
	// Syntax: CF.INSERT key [CAPACITY capacity] [NOCREATE] ITEMS item [item ...]
	cmd := client.CFInsert(testCtx, "CFInsertKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.INSERT", prefixHook.formatKey("CFInsertKey"), "CAPACITY", 5, "NOCREATE", "ITEMS", "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFInsertNX(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.insertnx/
	// Syntax: CF.INSERTNX key [CAPACITY capacity] [NOCREATE] ITEMS item [item ...]
	cmd := client.CFInsertNX(testCtx, "CFInsertKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.INSERTNX", prefixHook.formatKey("CFInsertKey"), "CAPACITY", 5, "NOCREATE", "ITEMS", "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFLoadChunk(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.loadchunk/
	// Syntax: CF.LOADCHUNK key iterator data
	key := fmt.Sprintf("testcf1%d", time.Now().Unix())
	client.Del(testCtx, key)
	client.Del(testCtx, "CFLCKey")
	cmdInsert := client.CFInsert(testCtx, "CFInsertKey", &redis.CFInsertOptions{
		Capacity: 5,
		NoCreate: true,
	}, "item")
	_, err := cmdInsert.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd cmdInsert %v", err)
	}
	r, err := client.CFScanDump(testCtx, "CFInsertKey", 0).Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd CFScanDump %v", err)
	}
	fd := []redis.ScanDump{r}
	for _, e := range fd {
		cmd := client.CFLoadChunk(testCtx, "CFLCKey", e.Iter, e.Data)
		_, err = cmd.Result()
		if err != nil && err != redis.Nil {
			t.Fatalf("cmd %v", err)
		}
		args := cmd.Args()
		argsStr := JoinInterfaceSliceBySpan(args...)
		t.Log("cmd: ", argsStr)
		target := JoinInterfaceSliceBySpan("CF.LOADCHUNK", prefixHook.formatKey("CFLCKey"), 1, r.Data)
		if !strings.EqualFold(argsStr, target) {
			t.Fatalf("waning! %s != %s", argsStr, target)
		}
	}
	t.Log("Success")
}

func TestCFMExists(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.mexists/
	// Syntax: CF.MEXISTS key item [item ...]
	client.CFInsert(testCtx, "CFMExistsKey", &redis.CFInsertOptions{Capacity: 1000, NoCreate: true}, "item")
	cmd := client.CFMExists(testCtx, "CFMExistsKey", "item")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.MEXISTS", prefixHook.formatKey("CFMExistsKey"), "item")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFReserve(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.reserve/
	// Syntax: CF.RESERVE key capacity [BUCKETSIZE bucketsize]  [MAXITERATIONS maxiterations] [EXPANSION expansion]
	key := fmt.Sprintf("CFReserveKey:%d", time.Now().Unix())
	client.CFInsert(testCtx, key, &redis.CFInsertOptions{Capacity: 1000, NoCreate: true}, "item")
	cmd := client.CFReserve(testCtx, key, 10)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.RESERVE", prefixHook.formatKey(key), 10)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCFScanDump(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cf.scandump/
	// Syntax: CF.SCANDUMP key iterator
	client.CFInsert(testCtx, "CFScanDumpKey", &redis.CFInsertOptions{}, "item")
	cmd := client.CFScanDump(testCtx, "CFScanDumpKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CF.SCANDUMP", prefixHook.formatKey("CFScanDumpKey"), 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLIENTCACHING(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-caching/
	// Syntax: CLIENT CACHING <YES | NO>
	t.Errorf("go-redis v9 not supported CLIENT CACHING")
}

func TestClientGetName(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-getname/
	// Syntax: CLIENT GETNAME
	cmd := client.ClientGetName(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT GETNAME")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLIENTGETREDIR(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-getredir/
	// Syntax: CLIENT GETREDIR
	t.Errorf("go-redis v9 not supported CLIENT GETREDIR")
}

func TestClientID(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-id/
	// Syntax: CLIENT ID
	cmd := client.ClientID(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT ID")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClientInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-info/
	// Syntax: CLIENT INFO
	cmd := client.ClientInfo(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT INFO")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClientKill(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-kill/
	// Syntax: CLIENT KILL <ip:port | <[ID client-id] | [TYPE <NORMAL | MASTER |  SLAVE | REPLICA | PUBSUB>] | [USER username] | [ADDR ip:port] |  [LADDR ip:port] | [SKIPME <YES | NO>] | [MAXAGE maxage]  [[ID client-id] | [TYPE <NORMAL | MASTER | SLAVE | REPLICA |  PUBSUB>] | [USER username] | [ADDR ip:port] | [LADDR ip:port] |  [SKIPME <YES | NO>] | [MAXAGE maxage] ...]>>
	cmd := client.ClientKill(testCtx, "string")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT KILL", "string")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClientList(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-list/
	// Syntax: CLIENT LIST [TYPE <NORMAL | MASTER | REPLICA | PUBSUB>]  [ID client-id [client-id ...]]
	cmd := client.ClientList(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT LIST")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLIENTNOEVICT(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-no-evict/
	// Syntax: CLIENT NO-EVICT <ON | OFF>
	t.Errorf("go-redis v9 not supported CLIENT NO-EVICT")
}

func TestCLIENTNOTOUCH(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-no-touch/
	// Syntax: CLIENT NO-TOUCH <ON | OFF>
	t.Errorf("go-redis v9 not supported CLIENT NO-TOUCH")
}

func TestClientPause(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-pause/
	// Syntax: CLIENT PAUSE timeout [WRITE | ALL]
	cmd := client.ClientPause(testCtx, testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT PAUSE", testFormatMs())
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLIENTREPLY(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-reply/
	// Syntax: CLIENT REPLY <ON | OFF | SKIP>
	t.Errorf("go-redis v9 not supported CLIENT REPLY")
}

func TestCLIENTSETINFO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-setinfo/
	// Syntax: CLIENT SETINFO <LIB-NAME libname | LIB-VER libver>
	t.Errorf("go-redis v9 not supported CLIENT SETINFO")
}

func TestCLIENTSETNAME(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-setname/
	// Syntax: CLIENT SETNAME connection-name
	t.Errorf("go-redis v9 not supported CLIENT SETNAME")
}

func TestCLIENTTRACKING(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-tracking/
	// Syntax: CLIENT TRACKING <ON | OFF> [REDIRECT client-id] [PREFIX prefix  [PREFIX prefix ...]] [BCAST] [OPTIN] [OPTOUT] [NOLOOP]
	t.Errorf("go-redis v9 not supported CLIENT TRACKING")
}

func TestCLIENTTRACKINGINFO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-trackinginfo/
	// Syntax: CLIENT TRACKINGINFO
	t.Errorf("go-redis v9 not supported CLIENT TRACKINGINFO")
}

func TestClientUnblock(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-unblock/
	// Syntax: CLIENT UNBLOCK client-id [TIMEOUT | ERROR]
	cmd := client.ClientUnblock(testCtx, 1)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT UNBLOCK", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClientUnpause(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/client-unpause/
	// Syntax: CLIENT UNPAUSE
	cmd := client.ClientUnpause(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLIENT UNPAUSE")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterAddSlots(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-addslots/
	// Syntax: CLUSTER ADDSLOTS slot [slot ...]
	//cmd := client.ClusterAddSlots(testCtx, 1, 3)
	////_, err := cmd.Result()
	////if err != nil && err != redis.Nil {
	////	t.Fatalf("cmd %v", err)
	////}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("CLUSTER ADDSLOTS", 1, 3)
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("CLUSTER ADDSLOTS test not enabled")
}

func TestClusterAddSlotsRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-addslotsrange/
	// Syntax: CLUSTER ADDSLOTSRANGE start-slot end-slot [start-slot end-slot ...]
	//cmd := client.ClusterAddSlotsRange(testCtx, 1, 3)
	////_, err := cmd.Result()
	////if err != nil && err != redis.Nil {
	////	t.Fatalf("cmd %v", err)
	////}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("CLUSTER ADDSLOTS", 1, 2, 3)
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("CLUSTER ADDSLOTSRANGE test not enabled")
}

func TestCLUSTERBUMPEPOCH(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-bumpepoch/
	// Syntax: CLUSTER BUMPEPOCH
	t.Errorf("go-redis v9 not supported CLUSTER BUMPEPOCH")
}

func TestClusterCountFailureReports(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-count-failure-reports/
	// Syntax: CLUSTER COUNT-FAILURE-REPORTS node-id
	cmd := client.ClusterCountFailureReports(testCtx, "node-id")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER COUNT-FAILURE-REPORTS", "node-id")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterCountKeysInSlot(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-countkeysinslot/
	// Syntax: CLUSTER COUNTKEYSINSLOT slot
	cmd := client.ClusterCountKeysInSlot(testCtx, 1)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER COUNTKEYSINSLOT", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterDelSlots(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-delslots/
	// Syntax: CLUSTER DELSLOTS slot [slot ...]
	//cmd := client.ClusterDelSlots(testCtx, 1, 3)
	////_, err := cmd.Result()
	////if err != nil && err != redis.Nil {
	////	t.Fatalf("cmd %v", err)
	////}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("CLUSTER DELSLOTS", 1, 3)
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("CLUSTER DELSLOTS test not enabled")
}

func TestClusterDelSlotsRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-delslotsrange/
	// Syntax: CLUSTER DELSLOTSRANGE start-slot end-slot [start-slot end-slot ...]
	//cmd := client.ClusterDelSlotsRange(testCtx, 1, 3)
	////_, err := cmd.Result()
	////if err != nil && err != redis.Nil {
	////	t.Fatalf("cmd %v", err)
	////}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("CLUSTER DELSLOTS", 1, 2, 3)
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("CLUSTER DELSLOTSRANGE test not enabled")
}

func TestClusterFailover(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-failover/
	// Syntax: CLUSTER FAILOVER [FORCE | TAKEOVER]
	cmd := client.ClusterFailover(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER FAILOVER")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLUSTERFLUSHSLOTS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-flushslots/
	// Syntax: CLUSTER FLUSHSLOTS
	t.Errorf("go-redis v9 not supported CLUSTER FLUSHSLOTS")
}

func TestClusterForget(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-forget/
	// Syntax: CLUSTER FORGET node-id
	cmd := client.ClusterForget(testCtx, "node-id")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER FORGET", "node-id")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterGetKeysInSlot(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-getkeysinslot/
	// Syntax: CLUSTER GETKEYSINSLOT slot count
	cmd := client.ClusterGetKeysInSlot(testCtx, 1, 10)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER GETKEYSINSLOT", 1, 10)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-info/
	// Syntax: CLUSTER INFO
	cmd := client.ClusterInfo(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER INFO")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterKeySlot(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-keyslot/
	// Syntax: CLUSTER KEYSLOT key
	cmd := client.ClusterKeySlot(testCtx, "ClusterKeySlotKey")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER KEYSLOT", prefixHook.formatKey("ClusterKeySlotKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterLinks(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-links/
	// Syntax: CLUSTER LINKS
	cmd := client.ClusterLinks(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER LINKS")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterMeet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-meet/
	// Syntax: CLUSTER MEET ip port [cluster-bus-port]
	cmd := client.ClusterMeet(testCtx, "127.0.0.1", "6379")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER MEET", "127.0.0.1", "6379")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterMyID(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-myid/
	// Syntax: CLUSTER MYID
	cmd := client.ClusterMyID(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER MYID")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterMyShardID(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-myshardid/
	// Syntax: CLUSTER MYSHARDID
	cmd := client.ClusterMyShardID(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER MYSHARDID")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterNodes(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-nodes/
	// Syntax: CLUSTER NODES
	cmd := client.ClusterNodes(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER NODES")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLUSTERREPLICAS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-replicas/
	// Syntax: CLUSTER REPLICAS node-id
	t.Errorf("go-redis v9 not supported CLUSTER REPLICAS")
}

func TestClusterReplicate(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-replicate/
	// Syntax: CLUSTER REPLICATE node-id
	cmd := client.ClusterReplicate(testCtx, "node-id")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER REPLICATE", "node-id")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLUSTERRESET(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-reset/
	// Syntax: CLUSTER RESET [HARD | SOFT]
	//cmd := client.ClusterResetHard(testCtx)
	////_, err := cmd.Result()
	////if err != nil && err != redis.Nil {
	////	t.Fatalf("cmd %v", err)
	////}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("CLUSTER RESET", "HARD")
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("CLUSTER RESET test not enabled")
}

func TestClusterSaveConfig(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-saveconfig/
	// Syntax: CLUSTER SAVECONFIG
	cmd := client.ClusterSaveConfig(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER SAVECONFIG")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCLUSTERSETCONFIGEPOCH(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-set-config-epoch/
	// Syntax: CLUSTER SET-CONFIG-EPOCH config-epoch
	t.Errorf("go-redis v9 not supported CLUSTER SET-CONFIG-EPOCH")
}

func TestCLUSTERSETSLOT(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-setslot/
	// Syntax: CLUSTER SETSLOT slot <IMPORTING node-id | MIGRATING node-id |  NODE node-id | STABLE>
	t.Errorf("go-redis v9 not supported CLUSTER SETSLOT")
}

func TestClusterShards(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-shards/
	// Syntax: CLUSTER SHARDS
	cmd := client.ClusterShards(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER SHARDS")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterSlaves(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-slaves/
	// Syntax: CLUSTER SLAVES node-id
	cmd := client.ClusterSlaves(testCtx, "node-id")
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER SLAVES", "node-id")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestClusterSlots(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cluster-slots/
	// Syntax: CLUSTER SLOTS
	cmd := client.ClusterSlots(testCtx)
	//cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CLUSTER SLOTS")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCMSIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cms.incrby/
	// Syntax: CMS.INCRBY key item increment [item increment ...]
	cmd := client.CMSIncrBy(testCtx, "CMSKey", "foo", 10, "bar", 42)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CMS.INCRBY", prefixHook.formatKey("CMSKey"), "foo", 10, "bar", 42)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCMSInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cms.info/
	// Syntax: CMS.INFO key
	client.CMSInitByDim(testCtx, "CMSInfoKey", 1, 2)
	cmd := client.CMSInfo(testCtx, "CMSInfoKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CMS.INFO", prefixHook.formatKey("CMSInfoKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCMSInitByDim(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cms.initbydim/
	// Syntax: CMS.INITBYDIM key width depth
	cmd := client.CMSInitByDim(testCtx, "CMSInitByDimKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CMS.INITBYDIM", prefixHook.formatKey("CMSInitByDimKey"), 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCMSInitByProb(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cms.initbyprob/
	// Syntax: CMS.INITBYPROB key error probability
	client.CMSInitByDim(testCtx, "CMSInitByProbKey", 1, 2)
	cmd := client.CMSInitByProb(testCtx, "CMSInitByProbKey", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CMS.INITBYPROB", prefixHook.formatKey("CMSInitByProbKey"), 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCMSMerge(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cms.merge/
	// Syntax: CMS.MERGE destination numKeys source [source ...] [WEIGHTS weight  [weight ...]]
	client.CMSInitByDim(testCtx, "CMSMerge", 1, 2)
	client.CMSInitByDim(testCtx, "CMSMerge1", 1, 2)
	client.CMSInitByDim(testCtx, "CMSMerge2", 1, 2)
	cmd := client.CMSMerge(testCtx, "CMSMerge", "CMSMerge1", "CMSMerge2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CMS.MERGE", prefixHook.formatKey("CMSMerge"), 2, prefixHook.formatKey("CMSMerge1"), prefixHook.formatKey("CMSMerge2"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCMSQuery(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/cms.query/
	// Syntax: CMS.QUERY key item [item ...]
	client.CMSInitByDim(testCtx, "CMSQueryKey", 1, 2)
	cmd := client.CMSQuery(testCtx, "CMSQueryKey", "foo", "bar")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CMS.QUERY", prefixHook.formatKey("CMSQueryKey"), "foo", "bar")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCommand(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command/
	// Syntax: COMMAND
	cmd := client.Command(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("COMMAND")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCOMMANDCOUNT(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command-count/
	// Syntax: COMMAND COUNT
	t.Errorf("go-redis v9 not supported COMMAND COUNT")
}

func TestCOMMANDDOCS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command-docs/
	// Syntax: COMMAND DOCS [command-name [command-name ...]]
	t.Errorf("go-redis v9 not supported COMMAND DOCS")
}

func TestCommandGetKeys(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command-getkeys/
	// Syntax: COMMAND GETKEYS command [arg [arg ...]]
	cmd := client.CommandGetKeys(testCtx, "SORT", "mylist", "ALPHA", "STORE", "outlist")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("COMMAND GETKEYS", "SORT", "mylist", "ALPHA", "STORE", "outlist")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCommandGetKeysAndFlags(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command-getkeysandflags/
	// Syntax: COMMAND GETKEYSANDFLAGS command [arg [arg ...]]
	cmd := client.CommandGetKeysAndFlags(testCtx, "LMOVE", "mylist1", "mylist2", "left", "left")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("COMMAND GETKEYSANDFLAGS", "LMOVE", "mylist1", "mylist2", "left", "left")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCOMMANDINFO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command-info/
	// Syntax: COMMAND INFO [command-name [command-name ...]]
	t.Errorf("go-redis v9 not supported COMMAND INFO")
}

func TestCommandList(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/command-list/
	// Syntax: COMMAND LIST [FILTERBY <MODULE module-name | ACLCAT category |  PATTERN pattern>]
	cmd := client.CommandList(testCtx, &redis.FilterBy{
		Module:  "",
		ACLCat:  "ACLCAT category",
		Pattern: "",
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("COMMAND LIST", "filterby", "aclcat ACLCAT category")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestConfigGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/config-get/
	// Syntax: CONFIG GET parameter [parameter ...]
	cmd := client.ConfigGet(testCtx, "appendonly")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CONFIG GET", "appendonly")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestConfigResetStat(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/config-resetstat/
	// Syntax: CONFIG RESETSTAT
	cmd := client.ConfigResetStat(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CONFIG RESETSTAT")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestConfigRewrite(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/config-rewrite/
	// Syntax: CONFIG REWRITE
	cmd := client.ConfigRewrite(testCtx)
	//cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CONFIG REWRITE")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestConfigSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/config-set/
	// Syntax: CONFIG SET parameter value [parameter value ...]
	cmd := client.ConfigSet(testCtx, "appendonly", "yes")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("CONFIG SET", "appendonly", "yes")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestCopy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/copy/
	// Syntax: COPY source destination [DB destination-db] [REPLACE]
	cmd := client.Copy(testCtx, "CopyKey1", "CopyKey2", 0, true)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("COPY", prefixHook.formatKey("CopyKey1"), prefixHook.formatKey("CopyKey2"), "DB", "0", "REPLACE")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
