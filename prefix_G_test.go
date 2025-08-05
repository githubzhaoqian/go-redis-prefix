package redisprefix

import (
	"strings"
	"testing"

	"github.com/redis/go-redis/v9"
)

func TestGeoAdd(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/geoadd/
	// Syntax: GEOADD key [NX | XX] [CH] longitude latitude member [longitude  latitude member ...]
	cmd := client.GeoAdd(testCtx, "GeoAddKey", &redis.GeoLocation{
		Name:      "name",
		Longitude: 0.1,
		Latitude:  0.1,
		Dist:      0,
		GeoHash:   0,
	})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEOADD", prefixHook.formatKey("GeoAddKey"), 0.1, 0.1, "name")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGeoDist(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/geodist/
	// Syntax: GEODIST key member1 member2 [M | KM | FT | MI]
	cmd := client.GeoDist(testCtx, "GeoDistKey", "member1", "member2", "M")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEODIST", prefixHook.formatKey("GeoDistKey"), "member1", "member2", "M")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGeoHash(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/geohash/
	// Syntax: GEOHASH key [member [member ...]]
	cmd := client.GeoHash(testCtx, "GeoHashKey", "item1", "item2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEOHASH", prefixHook.formatKey("GeoHashKey"), "item1", "item2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGeoPos(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/geopos/
	// Syntax: GEOPOS key [member [member ...]]
	cmd := client.GeoPos(testCtx, "GeoPosKey", "item1", "item2")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEOPOS", prefixHook.formatKey("GeoPosKey"), "item1", "item2")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGeoRadius(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/georadius/
	// Syntax: GEORADIUS key longitude latitude radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]  [STORE key | STOREDIST key]
	cmd := client.GeoRadius(testCtx, "GeoRadiusKey", 0.1, 0.1, &redis.GeoRadiusQuery{})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("georadius_ro", prefixHook.formatKey("GeoRadiusKey"), 0.1, 0.1, 0, "km")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGeoRadiusByMember(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/georadiusbymember/
	// Syntax: GEORADIUSBYMEMBER_RO key member radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]
	cmd := client.GeoRadiusByMember(testCtx, "GeoRadiusByMemberKey", "member", &redis.GeoRadiusQuery{})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEORADIUSBYMEMBER_RO", prefixHook.formatKey("GeoRadiusByMemberKey"), "member", 0, "km")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGEORADIUSBYMEMBERRO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/georadiusbymember_ro/
	// Syntax: GEORADIUSBYMEMBER_RO key member radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]
	TestGeoRadiusByMember(t)
}

func TestGEORADIUSRO(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/georadius_ro/
	// Syntax: GEORADIUS_RO key longitude latitude radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]
	TestGeoRadius(t)
}

func TestGeoSearch(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/geosearch/
	// Syntax: GEOSEARCH key
	// <FROMMEMBER member | FROMLONLAT longitude latitude>
	// <BYRADIUS radius <M | KM | FT | MI> | BYBOX width height <M | KM |  FT | MI>>
	// [ASC | DESC] [COUNT count [ANY]] [WITHCOORD] [WITHDIST]  [WITHHASH]
	cmd := client.GeoSearch(testCtx, "GeoSearchKey", &redis.GeoSearchQuery{})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEOSEARCH", prefixHook.formatKey("GeoSearchKey"), "fromlonlat", 0, 0, "bybox", 0, 0, "km")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGeoSearchStore(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/geosearchstore/
	// Syntax: GEOSEARCHSTORE destination source
	// <FROMMEMBER member |  FROMLONLAT longitude latitude>
	// <BYRADIUS radius <M | KM | FT | MI>  | BYBOX width height <M | KM | FT | MI>>
	// [ASC | DESC] [COUNT count  [ANY]] [STOREDIST]
	cmd := client.GeoSearchStore(testCtx, "GeoSearchStoreKey", "GeoSearchStoreStore", &redis.GeoSearchStoreQuery{})
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GEOSEARCHSTORE",
		prefixHook.formatKey("GeoSearchStoreStore"), prefixHook.formatKey("GeoSearchStoreKey"),
		"fromlonlat", 0, 0, "bybox", 0, 0, "km")
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/get/
	// Syntax: GET key
	cmd := client.Get(testCtx, "GetKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GET", prefixHook.formatKey("GetKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGetBit(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/getbit/
	// Syntax: GETBIT key offset
	cmd := client.GetBit(testCtx, "GetBitKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GETBIT", prefixHook.formatKey("GetBitKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGetDel(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/getdel/
	// Syntax: GETDEL key
	cmd := client.GetDel(testCtx, "GetDelKey")
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GETDEL", prefixHook.formatKey("GetDelKey"))
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGetEx(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/getex/
	// Syntax: GETEX key [EX seconds | PX milliseconds | EXAT unix-time-seconds |  PXAT unix-time-milliseconds | PERSIST]
	cmd := client.GetEx(testCtx, "GetExKey", testTimeout)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GETEX", prefixHook.formatKey("GetExKey"), "ex", testFormatSec())
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGetRange(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/getrange/
	// Syntax: GETRANGE key start end
	cmd := client.GetRange(testCtx, "GetRangeKey", 1, 10)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GETRANGE", prefixHook.formatKey("GetRangeKey"), 1, 10)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}

func TestGetSet(t *testing.T) {
	// Url: https://redis.io/docs/latest/commands/getset/
	// Syntax: GETSET key value
	cmd := client.GetSet(testCtx, "GetSetKey", 1)
	_, err := cmd.Result()
	if err != nil && err != redis.Nil {
		t.Fatalf("cmd %v", err)
	}
	args := cmd.Args()
	argsSrt := JoinInterfaceSliceBySpan(args...)
	t.Log("cmd: ", argsSrt)
	target := JoinInterfaceSliceBySpan("GETSET", prefixHook.formatKey("GetSetKey"), 1)
	if !strings.EqualFold(argsSrt, target) {
		t.Fatalf("waning! %s != %s", argsSrt, target)
	}
	t.Log("Success")
}
