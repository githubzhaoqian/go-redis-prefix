package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestJSONArrAppend(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.arrappend/
	// Syntax: JSON.ARRAPPEND key path value [value ...]
	client.JSONSet(testCtx, "JSONArrAppendKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONArrAppend(testCtx, "JSONArrAppendKey", "$..a", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.ARRAPPEND", prefixHook.formatKey("JSONArrAppendKey"), "$..a", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONArrIndex(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.arrindex/
	// Syntax: JSON.ARRINDEX key path value [start [stop]]
	client.JSONSet(testCtx, "JSONArrIndexKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONArrIndex(testCtx, "JSONArrIndexKey", "$..a", 10)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.ARRINDEX", prefixHook.formatKey("JSONArrIndexKey"), "$..a", 10)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONArrInsert(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.arrinsert/
	// Syntax: JSON.ARRINSERT key path index value [value ...]
	client.JSONSet(testCtx, "JSONArrInsertKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONArrInsert(testCtx, "JSONArrInsertKey", "$..a", 1, 2)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.ARRINSERT", prefixHook.formatKey("JSONArrInsertKey"), "$..a", 1, 2)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONArrLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.arrlen/
	// Syntax: JSON.ARRLEN key [path]
	client.JSONSet(testCtx, "JSONArrLenKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONArrLen(testCtx, "JSONArrLenKey", "$..a")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.ARRLEN", prefixHook.formatKey("JSONArrLenKey"), "$..a")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONArrPop(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.arrpop/
	// Syntax: JSON.ARRPOP key [path [index]]
	client.JSONSet(testCtx, "JSONArrPopKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONArrPop(testCtx, "JSONArrPopKey", "$..a", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.ARRPOP", prefixHook.formatKey("JSONArrPopKey"), "$..a", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONArrTrim(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.arrtrim/
	// Syntax: JSON.ARRTRIM key path start stop
	client.JSONSet(testCtx, "JSONArrTrimKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	stop := 3
	cmd := client.JSONArrTrimWithArgs(testCtx, "JSONArrTrimKey", "$..a", &redis.JSONArrTrimArgs{Start: -1, Stop: &stop})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.ARRTRIM", prefixHook.formatKey("JSONArrTrimKey"), "$..a", -1, 3)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONClear(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.clear/
	// Syntax: JSON.CLEAR key [path]
	client.JSONSet(testCtx, "JSONClearKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONClear(testCtx, "JSONClearKey", "$..a")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.CLEAR", prefixHook.formatKey("JSONClearKey"), "$..a")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONDEBUG(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.debug/
	// Syntax: JSON.DEBUG
	t.Errorf("go-redis v9 not supported JSON.DEBUG")
}

func TestJSONDebugMemory(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.debug-memory/
	// Syntax: JSON.DEBUG MEMORY key [path]
	t.Errorf("go-redis v9 not implemented JSON.DEBUG")
}

func TestJSONDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.del/
	// Syntax: JSON.DEL key [path]
	client.JSONSet(testCtx, "JSONDelKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONDel(testCtx, "JSONDelKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.DEL", prefixHook.formatKey("JSONDelKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONForget(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.forget/
	// Syntax: JSON.FORGET key [path]
	client.JSONSet(testCtx, "JSONForgetKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONForget(testCtx, "JSONForgetKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.FORGET", prefixHook.formatKey("JSONForgetKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.get/
	// Syntax: JSON.GET key [INDENT indent] [NEWLINE newline] [SPACE space] [path  [path ...]]
	client.JSONSet(testCtx, "JSONGetKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONGet(testCtx, "JSONGetKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.GET", prefixHook.formatKey("JSONGetKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONMerge(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.merge/
	// Syntax: JSON.MERGE key path value
	client.JSONSet(testCtx, "JSONMergeKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONMerge(testCtx, "JSONMergeKey", "$..a", "111")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.MERGE", prefixHook.formatKey("JSONMergeKey"), "$..a", "111")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONMGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.mget/
	// Syntax: JSON.MGET key [key ...] path
	client.JSONSet(testCtx, "JSONMGetKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONMGet(testCtx, "$", "JSONMGetKey", "JSONMGetKey1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.MGET", prefixHook.formatKey("JSONMGetKey"), prefixHook.formatKey("JSONMGetKey1"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONMSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.mset/
	// Syntax: JSON.MSET key path value [key path value ...]
	client.JSONSet(testCtx, "JSONArrPopKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONMSet(testCtx, "JSONArrPopKey", "$.a", `[1]`, "JSONArrPopKey", "$..a", `[1]`)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.MSET", prefixHook.formatKey("JSONArrPopKey"), "$.a", `[1]`, prefixHook.formatKey("JSONArrPopKey"), "$..a", `[1]`)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONNumIncrBy(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.numincrby/
	// Syntax: JSON.NUMINCRBY key path value
	client.JSONSet(testCtx, "JSONNumIncrByKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONNumIncrBy(testCtx, "JSONNumIncrByKey", "$.c", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.NUMINCRBY", prefixHook.formatKey("JSONNumIncrByKey"), "$.c", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONNUMMULTBY(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.nummultby/
	// Syntax: JSON.NUMMULTBY key path value
	t.Errorf("go-redis v9 not supported JSON.NUMMULTBY")
}

func TestJSONObjKeys(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.objkeys/
	// Syntax: JSON.OBJKEYS key [path]
	client.JSONSet(testCtx, "JSONObjKeysKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONObjKeys(testCtx, "JSONObjKeysKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.OBJKEYS", prefixHook.formatKey("JSONObjKeysKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONObjLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.objlen/
	// Syntax: JSON.OBJLEN key [path]
	client.JSONSet(testCtx, "JSONObjLenKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONObjLen(testCtx, "JSONObjLenKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.OBJLEN", prefixHook.formatKey("JSONObjLenKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONRESP(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.resp/
	// Syntax: JSON.RESP key [path]
	t.Errorf("go-redis v9 not supported JSON.RESP")
}

func TestJSONSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.set/
	// Syntax: JSON.SET key path value [NX | XX]
	client.JSONSet(testCtx, "JSONSetKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONSet(testCtx, "JSONSetKey", "$.a", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.SET", prefixHook.formatKey("JSONSetKey"), "$.a", 1)
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONStrAppend(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.strappend/
	// Syntax: JSON.STRAPPEND key [path] value
	client.JSONSet(testCtx, "JSONStrAppendKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONStrAppend(testCtx, "JSONStrAppendKey", "$.c", "cc")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.STRAPPEND", prefixHook.formatKey("JSONStrAppendKey"), "$.c", "cc")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONStrLen(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.strlen/
	// Syntax: JSON.STRLEN key [path]
	client.JSONSet(testCtx, "JSONStrLenKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONStrLen(testCtx, "JSONStrLenKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.STRLEN", prefixHook.formatKey("JSONStrLenKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONToggle(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.toggle/
	// Syntax: JSON.TOGGLE key path
	client.JSONSet(testCtx, "JSONToggleKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONToggle(testCtx, "JSONToggleKey", "$")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.TOGGLE", prefixHook.formatKey("JSONToggleKey"), "$")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}

func TestJSONType(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/json.type/
	// Syntax: JSON.TYPE key [path]
	client.JSONSet(testCtx, "JSONTypeKey", "$", `{"a": [10], "b": {"a": [12, 13]}}`)
	cmd := client.JSONType(testCtx, "JSONTypeKey", "$.a")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsStr := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsStr)
	target := JoinInterfaceSliceBySpan("JSON.TYPE", prefixHook.formatKey("JSONTypeKey"), "$.a")
	if !strings.EqualFold(argsStr, target) {
		t.Fatalf("waning! %s != %s", argsStr, target)
	}
	t.Log("Success")
}
