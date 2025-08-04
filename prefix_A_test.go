package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestACLCat(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-cat/
	// Syntax: ACL CAT [category]
	cmd := client.ACLCat(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("acl cat")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestACLDelUser(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-deluser/
	// Syntax: ACL DELUSER username [username ...]
	cmd := client.ACLDelUser(testCtx, "DelUser")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("acl deluser", "DelUser")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestACLDryRun(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-dryrun/
	// Syntax: ACL DRYRUN username command [arg [arg ...]]
	client.ACLSetUser(testCtx, "DryRunKey", "+SET", "~*")
	cmd := client.ACLDryRun(testCtx, "DryRunKey", "set", "foo", "bar")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("acl dryrun", "DryRunKey", "set", "foo", "bar")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestACLGENPASS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-genpass/
	// Syntax: ACL GENPASS [bits]
	t.Errorf("go-redis v9 not supported ACL GENPASS")
}

func TestACLGETUSER(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-getuser/
	// Syntax: ACL GETUSER username
	t.Errorf("go-redis v9 not supported ACL GETUSER")
}

func TestACLList(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-list/
	// Syntax: ACL LIST
	cmd := client.ACLList(testCtx)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("acl list")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestACLLOAD(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-load/
	// Syntax: ACL LOAD
	t.Errorf("go-redis v9 not supported ACL LOAD")
}

func TestACLLog(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-log/
	// Syntax: ACL LOG [count | RESET]
	cmd := client.ACLLog(testCtx, 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("acl log", 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestACLSAVE(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-save/
	// Syntax: ACL SAVE
	t.Errorf("go-redis v9 not supported ACL SAVE")
}

func TestACLSetUser(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-setuser/
	// Syntax: ACL SETUSER username [rule [rule ...]]
	cmd := client.ACLSetUser(testCtx, "SetUser", "+SET", "~*")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("acl setuser", "SetUser", "+SET", "~*")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestACLUSERS(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-users/
	// Syntax: ACL USERS
	t.Errorf("go-redis v9 not supported ACL USERS")
}

func TestACLWHOAMI(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/acl-whoami/
	// Syntax: ACL WHOAMI
	t.Errorf("go-redis v9 not supported ACL WHOAMI")
}

func TestAppend(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/append/
	// Syntax: APPEND key value
	cmd := client.Append(testCtx, "append", "append1")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("append", prefixHook.formatKey("append"), "append1")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestASKING(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/asking/
	// Syntax: ASKING
	t.Errorf("go-redis v9 not supported ASKING")

}

func TestAUTH(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/auth/
	// Syntax: AUTH [username] password
	t.Errorf("go-redis v9 not supported AUTH")
}
