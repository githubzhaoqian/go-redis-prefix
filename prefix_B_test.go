package redisprefix

import (
	"testing"
)

var (
	bfKey = "BFKey"
	bfEle = "bfItem"
)

func TestBFADD(t *testing.T) {
	aclCatCmd := client.BFAdd(ctx, bfKey, bfEle)
	aclCatArgs := aclCatCmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(aclCatArgs...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.ADD", prefixHook.formatKey(bfKey), bfEle)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFCard(t *testing.T) {
	cmd := client.BFCard(ctx, bfKey)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.CARD", prefixHook.formatKey(bfKey))
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFExists(t *testing.T) {
	cmd := client.BFExists(ctx, bfKey, bfEle)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.EXISTS", prefixHook.formatKey(bfKey), bfEle)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFInfo(t *testing.T) {
	cmd := client.BFInfo(ctx, bfKey)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.INFO", prefixHook.formatKey(bfKey))
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFInsert(t *testing.T) {
	cmd := client.BFInsert(ctx, bfKey, nil, bfEle)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.INSERT", prefixHook.formatKey(bfKey), "ITEMS", bfEle)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFLoadChunk(t *testing.T) {
	cmd := client.BFLoadChunk(ctx, bfKey, 1, bfEle)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.LOADCHUNK", prefixHook.formatKey(bfKey), 1, bfEle)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFMAdd(t *testing.T) {
	cmd := client.BFMAdd(ctx, bfKey, bfEle)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.MADD", prefixHook.formatKey(bfKey), bfEle)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFMExists(t *testing.T) {
	cmd := client.BFMExists(ctx, bfKey, bfEle, bfEle)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.MEXISTS", prefixHook.formatKey(bfKey), bfEle, bfEle)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFReserve(t *testing.T) {
	cmd := client.BFReserve(ctx, bfKey, 1, 2)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.RESERVE", prefixHook.formatKey(bfKey), 1, 2)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestBFScanDump(t *testing.T) {
	cmd := client.BFScanDump(ctx, bfKey, 1)
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("BF.SCANDUMP", prefixHook.formatKey(bfKey), 1, 2)
	if argsSrt != target {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}
