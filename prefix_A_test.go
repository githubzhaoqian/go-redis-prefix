package redisprefix

import (
	"testing"
)

func TestACLCat(t *testing.T) {
	aclCatCmd := client.ACLCat(ctx)
	aclCatArgs := aclCatCmd.Args()
	t.Log("ACLCat Args: ", aclCatArgs)
	aclCatArgsSrt := JoinInterfaceSliceBySpan(aclCatArgs)
	if aclCatArgsSrt != "acl cat" {
		t.Fatal("ACLCat waning!")
	}
	t.Log("Success")
}

func TestACLDelUser(t *testing.T) {
	ACLDelUserCmd := client.ACLDelUser(ctx, "haah")
	ACLDelUserCmdArgs := ACLDelUserCmd.Args()
	t.Log("ACLDelUserCmd Args: ", ACLDelUserCmdArgs)
	aclCatArgsSrt := JoinInterfaceSliceBySpan(ACLDelUserCmdArgs)
	if aclCatArgsSrt != "acl deluser haah" {
		t.Fatal("ACLDelUser waning!")
	}
	t.Log("Success")
}

func TestACLDryRun(t *testing.T) {
	cmd := client.ACLDryRun(ctx, "virginia", "set", "foo", "bar")
	cmdArgs := cmd.Args()
	t.Log("ACLDryRun Args: ", cmdArgs)
	cmdArgsSrt := JoinInterfaceSliceBySpan(cmdArgs)
	if cmdArgsSrt != "acl dryrun virginia set foo bar" {
		t.Fatal("ACLDryRun waning!")
	}
	t.Log("Success")
}

func TestACLGENPASS(t *testing.T) {
	t.Errorf("go-redis not supported ACL GENPASS")
}

func TestACLGETUSER(t *testing.T) {
	t.Errorf("go-redis not supported ACL GETUSER")
}

func TestACLHELP(t *testing.T) {
	t.Errorf("go-redis not supported ACL HELP")
}

func TestACLLIST(t *testing.T) {
	cmd := client.ACLList(ctx)
	cmdArgs := cmd.Args()
	t.Log("ACLDryRun Args: ", cmdArgs)
	cmdArgsSrt := JoinInterfaceSliceBySpan(cmdArgs)
	if cmdArgsSrt != "acl list" {
		t.Fatal("ACL LIST waning!")
	}
	t.Log("Success")
}

func TestACLLOAD(t *testing.T) {
	t.Errorf("go-redis not supported ACL LOAD")
}

func TestACLLog(t *testing.T) {
	cmd := client.ACLLog(ctx, 1)
	cmdArgs := cmd.Args()
	t.Log("ACLLog Args: ", cmdArgs)
	cmdArgsSrt := JoinInterfaceSliceBySpan(cmdArgs)
	if cmdArgsSrt != "acl log 1" {
		t.Fatal("ACL Log waning!")
	}
	t.Log("Success")
}

func TestACLSAVE(t *testing.T) {
	t.Errorf("go-redis not supported ACL SAVE")
}

func TestACLSETUSER(t *testing.T) {
	cmd := client.ACLSetUser(ctx, "guguda", "run")
	cmdArgs := cmd.Args()
	t.Log("ACLLog Args: ", cmdArgs)
	cmdArgsSrt := JoinInterfaceSliceBySpan(cmdArgs)
	if cmdArgsSrt != "acl setuser guguda run" {
		t.Fatal("ACL SetUser waning!")
	}
	t.Log("Success")
}

func TestACLWHOAMI(t *testing.T) {
	t.Errorf("go-redis not supported ACL WHOAMI")
}

func TestAppend(t *testing.T) {
	cmd := client.Append(ctx, "append", "append1")
	cmdArgs := cmd.Args()
	t.Log("Append Args: ", cmdArgs)
	cmdArgsSrt := JoinInterfaceSliceBySpan(cmdArgs)
	if cmdArgsSrt != JoinInterfaceSliceBySpan([]interface{}{"append", prefixHook.formatKey("append"), "append1"}) {
		t.Fatal("Append waning!")
	}
	t.Log("Success")
}

func TestASKING(t *testing.T) {
	t.Errorf("go-redis not supported ASKING")
}

func TestAUTH(t *testing.T) {
	t.Errorf("go-redis not supported AUTH")
}
