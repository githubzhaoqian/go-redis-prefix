package redisprefix

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

type KeyPrefixHook struct {
	prefix string
}

var _ redis.Hook = (*KeyPrefixHook)(nil) // nolint:typecheck

// NewKeyPrefixHook create key prefix hook
func NewKeyPrefixHook(prefix string) *KeyPrefixHook {
	return &KeyPrefixHook{prefix}
}

// DialHook is used to intercept and modify the behavior of the connection establishment process.
// When it's called: This hook is invoked whenever a new connection to the Redis server is established.
// No need in this case, just skip
// nolint:typecheck
func (k *KeyPrefixHook) DialHook(next redis.DialHook) redis.DialHook {
	return next
}

// ProcessHook is used to intercept and modify the behavior of individual commands sent to the Redis server.
// When it's called: This hook is invoked before any command is processed.
// nolint:typecheck
func (k *KeyPrefixHook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		k.keyWithPrefix(cmd)
		return next(ctx, cmd)
	}
}

// ProcessPipelineHook is used to intercept and modify the behavior of commands sent to the Redis server as part of a pipeline.
// When it's called: This hook is invoked before a set of commands in a pipeline is processed.
// nolint:typecheck
func (k *KeyPrefixHook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmd []redis.Cmder) error {
		for _, cmd := range cmd {
			k.keyWithPrefix(cmd)
		}
		return next(ctx, cmd)
	}
}

// keyWithPrefix key with prefix
// nolint:typecheck
func (k *KeyPrefixHook) keyWithPrefix(cmd redis.Cmder) {
	fullName := strings.ToUpper(cmd.FullName())
	prefixType, ok := CommandPrefixType[fullName]
	if prefixType == PrefixJoinNextArg {
		if len(cmd.Args()) < 2 {
			return
		}
		nextArg := cmd.Args()[1]
		fullName = fmt.Sprintf("%s %s", fullName, strings.ToUpper(nextArg.(string)))
		prefixType, ok = CommandPrefixType[fullName]
	}
	if !ok || prefixType == PrefixNone {
		return
	}
	switch prefixType {
	case PrefixAll:
		for i := 1; i <= len(cmd.Args())-1; i++ {
			cmd.Args()[i] = k.formatKey(cmd.Args()[i].(string))
		}
	case PrefixExceptSecondAll:
		if len(cmd.Args()) < 3 {
			return
		}
		for i := 2; i <= len(cmd.Args())-1; i++ {
			cmd.Args()[i] = k.formatKey(cmd.Args()[i].(string))
		}
	case PrefixOdd:
		for i := 1; i <= len(cmd.Args())-1; i += 2 {
			cmd.Args()[i] = k.formatKey(cmd.Args()[i].(string))
		}
	case PrefixFirstEveryThree:
		for i := 1; i <= len(cmd.Args())-1; i += 3 {
			cmd.Args()[i] = k.formatKey(cmd.Args()[i].(string))
		}
	case PrefixExceptLast:
		for i := 1; i <= len(cmd.Args())-2; i++ {
			cmd.Args()[i] = k.formatKey(cmd.Args()[i].(string))
		}
	case PrefixSecond:
		if len(cmd.Args()) < 2 {
			return
		}
		cmd.Args()[1] = k.formatKey(cmd.Args()[1].(string))
	case PrefixThird:
		if len(cmd.Args()) < 3 {
			return
		}
		cmd.Args()[1] = k.formatKey(cmd.Args()[2].(string))
	case PrefixSecondAndThird:
		if len(cmd.Args()) < 3 {
			return
		}
		cmd.Args()[1] = k.formatKey(cmd.Args()[1].(string))
		cmd.Args()[2] = k.formatKey(cmd.Args()[2].(string))
	case PrefixNumkeys, PrefixNumkeysSecond:
		offset := 0
		if prefixType == PrefixNumkeysSecond {
			offset = 1
		}
		if len(cmd.Args()) < 3+offset {
			return
		}
		numKeys := cmd.Args()[1+offset]
		num := Int(numKeys)
		if len(cmd.Args()) < 2+offset+num {
			return
		}
		for i := 0; i < num; i++ {
			cmd.Args()[2+offset+i] = k.formatKey(cmd.Args()[2+offset+i].(string))
		}
	case PrefixSTREAMS:
		// Find the location of the keyword "STREAMS"
		streamsIdx := -1
		for i, v := range cmd.Args() {
			if s, ok := v.(string); ok && strings.ToUpper(s) == "STREAMS" {
				streamsIdx = i
				break
			}
		}
		if streamsIdx == -1 || streamsIdx+1 >= len(cmd.Args()) {
			return // Parameter mismatch
		}
		total := len(cmd.Args()) - streamsIdx - 1
		if total%2 != 0 {
			return // Parameter mismatch
		}
		numKeys := total / 2
		for i := 0; i < numKeys; i++ {
			cmd.Args()[streamsIdx+1+i] = k.formatKey(cmd.Args()[streamsIdx+1+i].(string))
		}
	case PrefixNumkeysStore:
		if len(cmd.Args()) < 4 {
			return
		}
		cmd.Args()[1] = k.formatKey(cmd.Args()[1].(string))
		numKeys := cmd.Args()[2]
		num := Int(numKeys)
		if len(cmd.Args()) < 3+num {
			return
		}
		for i := 0; i < num; i++ {
			cmd.Args()[3+i] = k.formatKey(cmd.Args()[3+i].(string))
		}
	}
}

// formatKey format key
func (k *KeyPrefixHook) formatKey(key string) string {
	return fmt.Sprintf("%s:%s", k.prefix, key)
}

// FilterPrefixNone filter Prefixnone clear the memory
func FilterPrefixNone() {
	oldMap := CommandPrefixType
	CommandPrefixType = make(map[string]PrefixType, len(oldMap))
	for k, v := range oldMap {
		if v == PrefixNone {
			continue
		}
		CommandPrefixType[k] = v
	}
}
