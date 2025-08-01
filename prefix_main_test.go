package redisprefix

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"

	"github.com/redis/go-redis/v9"
)

const (
	domain = "https://redis.io"
)

type commandData struct {
	Command        string
	Url            string
	Syntax         string
	AvailableSince string
}

var (
	replacer   = strings.NewReplacer("\n", "", "&nbsp;", " ", " ", " ")
	client     *redis.Client
	prefix     = "go-prefix"
	prefixHook = NewKeyPrefixHook(prefix)
	ctx        = context.Background()
)

func TestMain(m *testing.M) {
	// redis:8.0.3
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6739",
		DB:   0,
	})
	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	client.AddHook(prefixHook)
	os.Exit(m.Run())
}

//
//func TestACLCat(t *testing.T) {
//	aclCatCmd := client.ACLCat(ctx)
//	aclCatArgs := aclCatCmd.Args()
//	t.Log("ACLCat Args: ", aclCatArgs)
//	aclCatArgsSrt := JoinInterfaceSliceBySpan(aclCatArgs)
//	if aclCatArgsSrt != "acl cat" {
//		t.Fatal("ACLCat waning!")
//	}
//	t.Log("Success")
//}
//
//func TestACLDelUser(t *testing.T) {
//	ACLDelUserCmd := client.ACLDelUser(ctx, "haah")
//	ACLDelUserCmdArgs := ACLDelUserCmd.Args()
//	t.Log("ACLDelUserCmd Args: ", ACLDelUserCmdArgs)
//	aclCatArgsSrt := JoinInterfaceSliceBySpan(ACLDelUserCmdArgs)
//	if aclCatArgsSrt != "acl deluser haah" {
//		t.Fatal("ACLDelUser waning!")
//	}
//	t.Log("Success")
//}
//
//func TestRedisGet(t *testing.T) {
//	setKey := "test"
//	setVal := "setKey"
//	setCmd := client.Set(ctx, setKey, setVal, 60*time.Second)
//	setArgs := setCmd.Args()
//	setArgsSrt := JoinInterfaceSliceBySpan(setArgs)
//	t.Log("Set Args: ", setArgsSrt)
//	if setArgsSrt != JoinInterfaceSliceBySpan([]interface{}{setCmd.FullName(), prefixHook.formatKey(setKey), setVal, "ex", "60"}) {
//		t.Fatal("set prefix waning!")
//	}
//
//	t.Log("Get Args: ", client.Get(ctx, "test").Args())
//	t.Log("Get Val: ", client.Get(ctx, "test").Val())
//
//	t.Log("MGet Args: ", client.MGet(ctx, "test", "test1").Args())
//	t.Log("MGet Val: ", client.MGet(ctx, "test", "test1").Val())
//}
//
//func TestRedisSet(t *testing.T) {
//	t.Log("Set Args: ", client.Set(ctx, "SetTest", "Set hello redis prefix!", 60*time.Second).Args())
//
//	t.Log("MSet Args: ", client.MSet(ctx, "MSetTest", "MSet hello redis prefix!", "MSetTest1", "MSet 1 hello redis prefix!").Args())
//
//	t.Log("MGet Args: ", client.MGet(ctx, "SetTest", "MSetTest", "MSetTest1").Args())
//	t.Log("MGet Val: ", client.MGet(ctx, "SetTest", "MSetTest", "MSetTest1").Val())
//}
//
//func TestRedisHSet(t *testing.T) {
//	t.Log("HSet Args: ", client.HSet(ctx, "hTest", "key", "HSet hello redis prefix!").Args())
//	t.Log("HMGet Args: ", client.HMGet(ctx, "hTest", "key", "key1").Args())
//	t.Log("HMGet Args: ", client.HMGet(ctx, "hTest", "key", "key1").Val())
//}

func TestFilterPrefixNone(t *testing.T) {
	FilterPrefixNone()
	for k, v := range commandPrefixType {
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

// TestQuery query redis command
func TestQuery(t *testing.T) {
	resp, err := http.Get(domain + "/docs/latest/commands/")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	selection := doc.Find("#commands-grid article")
	if selection == nil {
		t.Fatal("article is empty")
	}
	warning := make(map[string]commandData)
	warningList := make([]string, 0)
	for _, node := range selection.Nodes {
		nodeDoc := goquery.NewDocumentFromNode(node)
		if err != nil {
			t.Fatal(err)
		}
		commond := nodeDoc.Find("h1").First().Text()
		commond = strings.TrimSpace(commond)
		if commond == "Commands" {
			continue
		}
		href, attrExist := nodeDoc.Find("a").First().Attr("href")
		if attrExist == false {
			t.Fatalf("%s href is empty", commond)
		}
		commondUrl := domain + href
		commondResp, err := http.Get(commondUrl)
		if err != nil {
			t.Fatal(err)
		}
		defer commondResp.Body.Close()
		commondDoc, err := goquery.NewDocumentFromReader(commondResp.Body)
		if err != nil {
			t.Fatal(err)
		}
		syntax := replacer.Replace(commondDoc.Find(".command-syntax").First().Text())
		availableSince := strings.TrimSpace(commondDoc.Find(".command-syntax").First().Next().Find("dd").First().Text())
		data := commandData{
			Command:        commond,
			Url:            commondUrl,
			Syntax:         syntax,
			AvailableSince: availableSince,
		}
		existData, exist := testCommandPrefixType[commond]
		_, commandExist := commandPrefixType[commond]
		if !exist || !commandExist || existData != data {
			warning[commond] = data
			warningList = append(warningList, fmt.Sprintf("\"%s\" : PrefixUnKnow, // %s", commond, syntax))
			t.Logf("warning: %s", commond)
			continue
		}
		t.Logf("command: %s OK!", commond)
	}
	if len(warning) > 0 {
		var warningStr string
		for _, w := range warningList {
			warningStr += w + "\n"
		}
		t.Logf("warning commandPrefixType: %s", warningStr)
		wBs, _ := json.Marshal(warning)
		t.Fatalf("warning map json: %s", string(wBs))
	}
	t.Log("END")
}
