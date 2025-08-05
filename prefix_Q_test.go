package redisprefix

import (
	"testing"
)

func TestQuit(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/quit/
	// Syntax: QUIT
	//cmd := client.Quit(testCtx)
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("QUIT")
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	t.Errorf("go-redis v9 not implemented QUIT")
}
