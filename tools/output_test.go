package tools

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"

	redisprefix "github.com/githubzhaoqian/go-redis-prefix"
)

var (
	methodReplacer = strings.NewReplacer(".", "", " ", "", "-", "", "_", "")
	version        = "v9"
	tmplStr        = `

// TODO
func Test{{.MethodName}}(t *testing.T) {
	// Url: {{.Url}}
	// Syntax: {{.Syntax}}
	{{ if .Signature -}}
	//cmd := client.{{.Signature}}
	//_, err := cmd.Result()
	//if err != nil && err != redis.Nil {
	//	t.Fatalf("cmd %v", err)
	//}
	//args := cmd.Args()
	//argsStr := JoinInterfaceSliceBySpan(args...)
	//t.Log("cmd: ", argsStr)
	//target := JoinInterfaceSliceBySpan("{{.Name}}"{{.JoinSignature}})
	//if !strings.EqualFold(argsStr, target) {
	//	t.Fatalf("waning! %s != %s", argsStr, target)
	//}
	//t.Log("Success")
	t.Errorf("TODO")
	{{- else -}}
	t.Errorf("go-redis {{.Version}} not supported {{.Name}}")
	{{- end}}
}
`
)

func TestOutput(t *testing.T) {
	// redis:8.0.3
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6739",
		DB:   0,
	})
	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	clientType := reflect.TypeOf(client)

	fmt.Printf("Total methods: %d\n", clientType.NumMethod())

	methodMap := make(map[string]reflect.Method)
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		methodName := strings.ToUpper(method.Name)
		methodMap[methodName] = method
	}
	tmpl, err := template.New("example").Parse(tmplStr)
	if err != nil {
		panic(err)
	}
	nameList := make([]string, 0, len(testCommandPrefixType))
	for name := range testCommandPrefixType {
		nameList = append(nameList, name)
	}
	sort.Slice(nameList, func(i, j int) bool {
		return nameList[i] <= nameList[j]
	})
	for _, name := range nameList {
		if len(name) == 0 {
			continue
		}
		cData := testCommandPrefixType[name]
		methodName := methodReplacer.Replace(name)
		method, exist := methodMap[methodName]
		if !exist {
			t.Logf("Warning %s method %s not exist!", name, methodName)
		}
		var (
			signature, joinSignature string
		)
		firstLetter := name[0]
		methodType := method.Type
		if methodType != nil {
			signature = fmt.Sprintf("%s(", method.Name)
			for j := 0; j < methodType.NumIn(); j++ {
				if j > 1 {
					signature += ", "
					joinSignature += ", "
				}
				argType := methodType.In(j)
				arg := argType.String()
				if arg == "*redis.Client" {
					continue
				}
				if arg == "context.Context" {
					arg = "testCtx"
				}
				signature += arg
				if arg != "testCtx" {
					joinSignature += arg
				}
			}
			signature += ") "
		}
		methodTitle := method.Name
		if methodTitle == "" {
			methodTitle = methodName
		}
		var buf bytes.Buffer
		data := struct {
			MethodName    string
			Name          string
			Signature     string
			Method        string
			Version       string
			JoinSignature string
			Url           string
			Syntax        template.HTML
		}{
			MethodName:    methodTitle,
			Name:          name,
			Signature:     signature,
			JoinSignature: joinSignature,
			Version:       version,
			Url:           cData.Url,
			Syntax:        template.HTML(cData.Syntax),
		}
		err = tmpl.Execute(&buf, data)
		if err != nil {
			t.Fatalf("%s method template err %v", name, err)
		}
		fileName := fmt.Sprintf("output/prefix_%s_test.go", string(firstLetter))
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			t.Fatalf("%s method OpenFile %s err %v", name, fileName, err)
		}
		_, err = file.Write(buf.Bytes())
		if err != nil {
			t.Fatalf("%s method WriteFile %s err %v", name, fileName, err)
		}
		file.Close()
	}
}

func TestOutputMdTable(t *testing.T) {
	nameList := make([]string, 0, len(testCommandPrefixType))
	for name := range testCommandPrefixType {
		nameList = append(nameList, name)
	}
	sort.Slice(nameList, func(i, j int) bool {
		return nameList[i] <= nameList[j]
	})
	var mdBuf bytes.Buffer
	mdBuf.WriteString("| Command | support | explain |\n|---------|---------|----|\n")
	for _, name := range nameList {
		if len(name) == 0 {
			continue
		}
		prefixType := redisprefix.CommandPrefixType[name]
		td := formatMdTable(name, prefixType)
		mdBuf.WriteString(td)
	}
	fileName := "output/command_table.md"
	file, err := os.OpenFile("output/command_table.md", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		t.Fatalf("OpenFile %s err %v", fileName, err)
	}
	_, err = file.Write(mdBuf.Bytes())
	if err != nil {
		t.Fatalf("WriteFile %s err %v", fileName, err)
	}
}

func formatMdTable(command string, prefixType redisprefix.PrefixType) string {
	support := "Yes"
	if prefixType == redisprefix.PrefixNone {
		support = "No"
	}
	return fmt.Sprintf("| %s | %s | |\n", command, support)
}
