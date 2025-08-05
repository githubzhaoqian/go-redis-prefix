package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestVAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vadd/
	// Syntax: VADD key [REDUCE dim] (FP32 | VALUES num) vector element [CAS] [NOQUANT | Q8 | BIN]  [EF build-exploration-factor] [SETATTR attributes] [M numlinks]
	cmd := client.VAdd(testCtx, "VAddKey", "element", &redis.VectorValues{Val: []float64{1, 2}})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VADD", prefixHook.formatKey("VAddKey"), "VALUES", 2, 1, 2, "element")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVCard(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vcard/
	// Syntax: VCARD key
	cmd := client.VCard(testCtx, "VCardKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VCARD", prefixHook.formatKey("VCardKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVDim(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vdim/
	// Syntax: VDIM key
	client.VAdd(testCtx, "VDimKey", "element", &redis.VectorValues{Val: []float64{1, 2}})
	cmd := client.VDim(testCtx, "VDimKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VDIM", prefixHook.formatKey("VDimKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVEmb(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vemb/
	// Syntax: VEMB key element [RAW]
	cmd := client.VEmb(testCtx, "VEmbKey", "element", true)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VEMB", prefixHook.formatKey("VEmbKey"), "element", "raw")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVGetAttr(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vgetattr/
	// Syntax: VGETATTR key element
	cmd := client.VGetAttr(testCtx, "VGetAttrKey", "element")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VGETATTR", prefixHook.formatKey("VGetAttrKey"), "element")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVInfo(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vinfo/
	// Syntax: VINFO key
	cmd := client.VInfo(testCtx, "VInfoKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VINFO", prefixHook.formatKey("VInfoKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVLinks(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vlinks/
	// Syntax: VLINKS key element [WITHSCORES]
	cmd := client.VLinks(testCtx, "VLinksKey", "element")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VLINKS", prefixHook.formatKey("VLinksKey"), "element")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVRandMember(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vrandmember/
	// Syntax: VRANDMEMBER key [count]
	cmd := client.VRandMember(testCtx, "VRandMemberKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VRANDMEMBER", prefixHook.formatKey("VRandMemberKey"))
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVRem(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vrem/
	// Syntax: VREM key element
	cmd := client.VRem(testCtx, "VRemKey", "element")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VREM", prefixHook.formatKey("VRemKey"), "element")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVSetAttr(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vsetattr/
	// Syntax: VSETATTR key element "{ JSON obj }"
	cmd := client.VSetAttr(testCtx, "VSetAttrKey", "element", "{ JSON obj }")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VSETATTR", prefixHook.formatKey("VSetAttrKey"), "element", "{ JSON obj }")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestVSim(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/vsim/
	// Syntax: VSIM key (ELE | FP32 | VALUES num) (vector | element) [WITHSCORES] [COUNT num] [EF search-exploration-factor]  [FILTER expression] [FILTER-EF max-filtering-effort] [TRUTH] [NOTHREAD]
	cmd := client.VSim(testCtx, "VSimKey", &redis.VectorValues{Val: []float64{1, 2}})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("VSIM", prefixHook.formatKey("VSimKey"), "VALUES", 2, 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
