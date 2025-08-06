package redisprefix

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

type commandData struct {
	Command        string
	Url            string
	Syntax         string
	AvailableSince string
}

var (
	client      *redis.Client
	prefix      = "go-prefix"
	prefixHook  = NewKeyPrefixHook(prefix)
	testCtx     = context.Background()
	testTimeout = 1 * time.Second
	lib1        = redis.Library{
		Name:   "mylib1",
		Engine: "LUA",
		Functions: []redis.Function{
			{
				Name:        "lib1_func1",
				Description: "This is the func-1 of lib 1",
				Flags:       []string{"allow-oom", "allow-stale"},
			},
		},
		Code: `#!lua name=%s
					
                     local function f1(keys, args)
                        local hash = keys[1]  -- Get the key name
                        local time = redis.call('TIME')[1]  -- Get the current time from the Redis server

                        -- Add the current timestamp to the arguments that the user passed to the function, stored in args
                        table.insert(args, '_updated_at')
                        table.insert(args, time)

                        -- Run HSET with the updated argument list
                        return redis.call('HSET', hash, unpack(args))
                     end

					redis.register_function{
						function_name='%s',
						description ='%s',
						callback=f1,
						flags={'%s', '%s'}
					}`,
	}
)

func TestMain(m *testing.M) {
	// redis:8.0.3
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6739",
		DB:   0,
	})
	err := client.Ping(testCtx).Err()
	if err != nil {
		panic(err)
	}
	client.AddHook(prefixHook)

	lib1Code := fmt.Sprintf(lib1.Code, lib1.Name, lib1.Functions[0].Name,
		lib1.Functions[0].Description, lib1.Functions[0].Flags[0], lib1.Functions[0].Flags[1])
	err = client.FunctionDelete(testCtx, lib1.Name).Err()
	if err != nil && err.Error() != "ERR Library not found" {
		panic(err)
	}
	err = client.FunctionLoad(testCtx, lib1Code).Err()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestFilterPrefixNone(t *testing.T) {
	FilterPrefixNone()
	for k, v := range CommandPrefixType {
		if v == PrefixNone {
			t.Fatalf("command: %s is PrefixNone", k)
		}
	}
	t.Log("END")
}

// TestJoinInterfaceSlice join interfaceSlice
func TestJoinInterfaceSlice(t *testing.T) {
	data := []interface{}{"1", "2", nil, 4}
	t.Log(JoinInterfaceSliceBySpan(data))
}

func testFormatSec() int64 {
	return formatSec(testTimeout)
}

func testFormatMs() int64 {
	return formatMs(testTimeout)
}
