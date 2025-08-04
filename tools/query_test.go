package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"

	redisprefix "github.com/githubzhaoqian/go-redis-prefix"
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
	replacer = strings.NewReplacer("\n", "", "&nbsp;", " ", " ", " ")
	ctx      = context.Background()
)

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
		_, commandExist := redisprefix.CommandPrefixType[commond]
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
