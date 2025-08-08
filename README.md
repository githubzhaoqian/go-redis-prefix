# go-redis-prefix
go-redis-prefix is a hook library that sets the prefix for github.com/redis/go-redis/v9 key

The version number corresponds to the go-redis version(major. minor). If you find that the versions do not match, you can contact me to create a branch version.
```golang
import (
    "github.com/redis/go-redis/v9"
    "github.com/githubzhaoqian/go-redis-prefix/v9"
)

func NewRedis() *redis.Client{
    // redisprefix.FilterPrefixNone()
    prefixHook  = redisprefix.NewKeyPrefixHook("redis-prefix")
    client = redis.NewClient(&redis.Options{
        Addr: "127.0.0.1:6739",
        DB:   0,
    })
    err := client.Ping(testCtx).Err()
    if err != nil {
        panic(err)
    }
    client.AddHook(prefixHook)
    return client
}
```

# Command details 

| Command | support | explain |
|---------|---------|----|
| ACL CAT | No | |
| ACL DELUSER | No | |
| ACL DRYRUN | No | |
| ACL GENPASS | No | |
| ACL GETUSER | No | |
| ACL LIST | No | |
| ACL LOAD | No | |
| ACL LOG | No | |
| ACL SAVE | No | |
| ACL SETUSER | No | |
| ACL USERS | No | |
| ACL WHOAMI | No | |
| APPEND | Yes | |
| ASKING | No | |
| AUTH | No | |
| BF.ADD | Yes | |
| BF.CARD | Yes | |
| BF.EXISTS | Yes | |
| BF.INFO | Yes | |
| BF.INSERT | Yes | |
| BF.LOADCHUNK | Yes | |
| BF.MADD | Yes | |
| BF.MEXISTS | Yes | |
| BF.RESERVE | Yes | |
| BF.SCANDUMP | Yes | |
| BGREWRITEAOF | No | |
| BGSAVE | No | |
| BITCOUNT | Yes | |
| BITFIELD | Yes | |
| BITFIELD_RO | Yes | |
| BITOP | Yes | |
| BITPOS | Yes | |
| BLMOVE | Yes | |
| BLMPOP | Yes | |
| BLPOP | Yes | |
| BRPOP | Yes | |
| BRPOPLPUSH | Yes | |
| BZMPOP | Yes | |
| BZPOPMAX | Yes | |
| BZPOPMIN | Yes | |
| CF.ADD | Yes | |
| CF.ADDNX | Yes | |
| CF.COUNT | Yes | |
| CF.DEL | Yes | |
| CF.EXISTS | Yes | |
| CF.INFO | Yes | |
| CF.INSERT | Yes | |
| CF.INSERTNX | Yes | |
| CF.LOADCHUNK | Yes | |
| CF.MEXISTS | Yes | |
| CF.RESERVE | Yes | |
| CF.SCANDUMP | Yes | |
| CLIENT CACHING | No | |
| CLIENT GETNAME | No | |
| CLIENT GETREDIR | No | |
| CLIENT ID | No | |
| CLIENT INFO | No | |
| CLIENT KILL | No | |
| CLIENT LIST | No | |
| CLIENT NO-EVICT | No | |
| CLIENT NO-TOUCH | No | |
| CLIENT PAUSE | No | |
| CLIENT REPLY | No | |
| CLIENT SETINFO | No | |
| CLIENT SETNAME | No | |
| CLIENT TRACKING | No | |
| CLIENT TRACKINGINFO | No | |
| CLIENT UNBLOCK | No | |
| CLIENT UNPAUSE | No | |
| CLUSTER ADDSLOTS | No | |
| CLUSTER ADDSLOTSRANGE | No | |
| CLUSTER BUMPEPOCH | No | |
| CLUSTER COUNT-FAILURE-REPORTS | No | |
| CLUSTER COUNTKEYSINSLOT | No | |
| CLUSTER DELSLOTS | No | |
| CLUSTER DELSLOTSRANGE | No | |
| CLUSTER FAILOVER | No | |
| CLUSTER FLUSHSLOTS | No | |
| CLUSTER FORGET | No | |
| CLUSTER GETKEYSINSLOT | No | |
| CLUSTER INFO | No | |
| CLUSTER KEYSLOT | Yes | |
| CLUSTER LINKS | No | |
| CLUSTER MEET | No | |
| CLUSTER MYID | No | |
| CLUSTER MYSHARDID | No | |
| CLUSTER NODES | No | |
| CLUSTER REPLICAS | No | |
| CLUSTER REPLICATE | No | |
| CLUSTER RESET | No | |
| CLUSTER SAVECONFIG | No | |
| CLUSTER SET-CONFIG-EPOCH | No | |
| CLUSTER SETSLOT | No | |
| CLUSTER SHARDS | No | |
| CLUSTER SLAVES | No | |
| CLUSTER SLOTS | No | |
| CMS.INCRBY | Yes | |
| CMS.INFO | Yes | |
| CMS.INITBYDIM | Yes | |
| CMS.INITBYPROB | No | |
| CMS.MERGE | Yes | |
| CMS.QUERY | Yes | |
| COMMAND | No | |
| COMMAND COUNT | No | |
| COMMAND DOCS | No | |
| COMMAND GETKEYS | No | |
| COMMAND GETKEYSANDFLAGS | No | |
| COMMAND INFO | No | |
| COMMAND LIST | No | |
| CONFIG GET | No | |
| CONFIG RESETSTAT | No | |
| CONFIG REWRITE | No | |
| CONFIG SET | No | |
| COPY | Yes | |
| DBSIZE | No | |
| DECR | Yes | |
| DECRBY | Yes | |
| DEL | Yes | |
| DISCARD | No | |
| DUMP | Yes | |
| ECHO | No | |
| EVAL | Yes | |
| EVALSHA | Yes | |
| EVALSHA_RO | Yes | |
| EVAL_RO | Yes | |
| EXEC | No | |
| EXISTS | Yes | |
| EXPIRE | Yes | |
| EXPIREAT | Yes | |
| EXPIRETIME | Yes | |
| FAILOVER | No | |
| FCALL | Yes | |
| FCALL_RO | Yes | |
| FLUSHALL | No | |
| FLUSHDB | No | |
| FT.AGGREGATE | No | |
| FT.ALIASADD | No | |
| FT.ALIASDEL | No | |
| FT.ALIASUPDATE | No | |
| FT.ALTER | No | |
| FT.CONFIG GET | No | |
| FT.CONFIG SET | No | |
| FT.CREATE | No | |
| FT.CURSOR DEL | No | |
| FT.CURSOR READ | No | |
| FT.DICTADD | No | |
| FT.DICTDEL | No | |
| FT.DICTDUMP | No | |
| FT.DROPINDEX | No | |
| FT.EXPLAIN | No | |
| FT.EXPLAINCLI | No | |
| FT.INFO | No | |
| FT.PROFILE | No | |
| FT.SEARCH | No | |
| FT.SPELLCHECK | No | |
| FT.SUGADD | Yes | |
| FT.SUGDEL | Yes | |
| FT.SUGGET | Yes | |
| FT.SUGLEN | Yes | |
| FT.SYNDUMP | No | |
| FT.SYNUPDATE | No | |
| FT.TAGVALS | No | |
| FT._LIST | No | |
| FUNCTION DELETE | No | |
| FUNCTION DUMP | No | |
| FUNCTION FLUSH | No | |
| FUNCTION KILL | No | |
| FUNCTION LIST | No | |
| FUNCTION LOAD | No | |
| FUNCTION RESTORE | No | |
| FUNCTION STATS | No | |
| GEOADD | Yes | |
| GEODIST | Yes | |
| GEOHASH | Yes | |
| GEOPOS | Yes | |
| GEORADIUS | Yes | |
| GEORADIUSBYMEMBER | Yes | |
| GEORADIUSBYMEMBER_RO | Yes | |
| GEORADIUS_RO | Yes | |
| GEOSEARCH | Yes | |
| GEOSEARCHSTORE | Yes | |
| GET | Yes | |
| GETBIT | Yes | |
| GETDEL | Yes | |
| GETEX | Yes | |
| GETRANGE | Yes | |
| GETSET | Yes | |
| HDEL | Yes | |
| HELLO | No | |
| HEXISTS | Yes | |
| HEXPIRE | Yes | |
| HEXPIREAT | Yes | |
| HEXPIRETIME | Yes | |
| HGET | Yes | |
| HGETALL | Yes | |
| HGETDEL | Yes | |
| HGETEX | Yes | |
| HINCRBY | Yes | |
| HINCRBYFLOAT | Yes | |
| HKEYS | Yes | |
| HLEN | Yes | |
| HMGET | Yes | |
| HMSET | Yes | |
| HPERSIST | Yes | |
| HPEXPIRE | Yes | |
| HPEXPIREAT | Yes | |
| HPEXPIRETIME | Yes | |
| HPTTL | Yes | |
| HRANDFIELD | Yes | |
| HSCAN | Yes | |
| HSET | Yes | |
| HSETEX | Yes | |
| HSETNX | Yes | |
| HSTRLEN | Yes | |
| HTTL | Yes | |
| HVALS | Yes | |
| INCR | Yes | |
| INCRBY | Yes | |
| INCRBYFLOAT | Yes | |
| INFO | No | |
| JSON.ARRAPPEND | Yes | |
| JSON.ARRINDEX | Yes | |
| JSON.ARRINSERT | Yes | |
| JSON.ARRLEN | Yes | |
| JSON.ARRPOP | Yes | |
| JSON.ARRTRIM | Yes | |
| JSON.CLEAR | Yes | |
| JSON.DEBUG | Yes | |
| JSON.DEBUG MEMORY | Yes | |
| JSON.DEL | Yes | |
| JSON.FORGET | Yes | |
| JSON.GET | Yes | |
| JSON.MERGE | Yes | |
| JSON.MGET | Yes | |
| JSON.MSET | Yes | |
| JSON.NUMINCRBY | Yes | |
| JSON.NUMMULTBY | Yes | |
| JSON.OBJKEYS | Yes | |
| JSON.OBJLEN | Yes | |
| JSON.RESP | Yes | |
| JSON.SET | Yes | |
| JSON.STRAPPEND | Yes | |
| JSON.STRLEN | Yes | |
| JSON.TOGGLE | Yes | |
| JSON.TYPE | Yes | |
| KEYS | No | |
| LASTSAVE | No | |
| LATENCY DOCTOR | No | |
| LATENCY GRAPH | No | |
| LATENCY HISTOGRAM | No | |
| LATENCY HISTORY | No | |
| LATENCY LATEST | No | |
| LATENCY RESET | No | |
| LCS | Yes | |
| LINDEX | Yes | |
| LINSERT | Yes | |
| LLEN | Yes | |
| LMOVE | Yes | |
| LMPOP | Yes | |
| LOLWUT | No | |
| LPOP | Yes | |
| LPOS | Yes | |
| LPUSH | Yes | |
| LPUSHX | Yes | |
| LRANGE | Yes | |
| LREM | Yes | |
| LSET | Yes | |
| LTRIM | Yes | |
| MEMORY DOCTOR | No | |
| MEMORY MALLOC-STATS | No | |
| MEMORY PURGE | No | |
| MEMORY STATS | No | |
| MEMORY USAGE | Yes | |
| MGET | Yes | |
| MIGRATE | Yes | |
| MODULE LIST | No | |
| MODULE LOAD | No | |
| MODULE LOADEX | No | |
| MODULE UNLOAD | No | |
| MONITOR | No | |
| MOVE | Yes | |
| MSET | Yes | |
| MSETNX | Yes | |
| MULTI | No | |
| OBJECT ENCODING | Yes | |
| OBJECT FREQ | Yes | |
| OBJECT IDLETIME | Yes | |
| OBJECT REFCOUNT | Yes | |
| PERSIST | Yes | |
| PEXPIRE | Yes | |
| PEXPIREAT | Yes | |
| PEXPIRETIME | Yes | |
| PFADD | Yes | |
| PFCOUNT | Yes | |
| PFDEBUG | Yes | |
| PFMERGE | Yes | |
| PFSELFTEST | No | |
| PING | No | |
| PSETEX | Yes | |
| PSUBSCRIBE | No | |
| PSYNC | No | |
| PTTL | Yes | |
| PUBLISH | No | |
| PUBSUB CHANNELS | No | |
| PUBSUB NUMPAT | No | |
| PUBSUB NUMSUB | No | |
| PUBSUB SHARDCHANNELS | No | |
| PUBSUB SHARDNUMSUB | No | |
| PUNSUBSCRIBE | No | |
| QUIT | No | |
| RANDOMKEY | No | |
| READONLY | No | |
| READWRITE | No | |
| RENAME | Yes | |
| RENAMENX | Yes | |
| REPLCONF | No | |
| REPLICAOF | No | |
| RESET | No | |
| RESTORE | Yes | |
| RESTORE-ASKING | Yes | |
| ROLE | No | |
| RPOP | Yes | |
| RPOPLPUSH | Yes | |
| RPUSH | Yes | |
| RPUSHX | Yes | |
| SADD | Yes | |
| SAVE | No | |
| SCAN | No | |
| SCARD | Yes | |
| SCRIPT DEBUG | No | |
| SCRIPT EXISTS | No | |
| SCRIPT FLUSH | No | |
| SCRIPT KILL | No | |
| SCRIPT LOAD | No | |
| SDIFF | Yes | |
| SDIFFSTORE | Yes | |
| SELECT | No | |
| SET | Yes | |
| SETBIT | Yes | |
| SETEX | Yes | |
| SETNX | Yes | |
| SETRANGE | Yes | |
| SHUTDOWN | No | |
| SINTER | Yes | |
| SINTERCARD | Yes | |
| SINTERSTORE | Yes | |
| SISMEMBER | Yes | |
| SLAVEOF | No | |
| SLOWLOG GET | No | |
| SLOWLOG LEN | No | |
| SLOWLOG RESET | No | |
| SMEMBERS | Yes | |
| SMISMEMBER | Yes | |
| SMOVE | Yes | |
| SORT | Yes | |
| SORT_RO | Yes | |
| SPOP | Yes | |
| SPUBLISH | No | |
| SRANDMEMBER | Yes | |
| SREM | Yes | |
| SSCAN | Yes | |
| SSUBSCRIBE | No | |
| STRLEN | Yes | |
| SUBSCRIBE | No | |
| SUBSTR | Yes | |
| SUNION | Yes | |
| SUNIONSTORE | Yes | |
| SUNSUBSCRIBE | No | |
| SWAPDB | No | |
| SYNC | No | |
| TDIGEST.ADD | Yes | |
| TDIGEST.BYRANK | Yes | |
| TDIGEST.BYREVRANK | Yes | |
| TDIGEST.CDF | Yes | |
| TDIGEST.CREATE | Yes | |
| TDIGEST.INFO | Yes | |
| TDIGEST.MAX | Yes | |
| TDIGEST.MERGE | Yes | |
| TDIGEST.MIN | Yes | |
| TDIGEST.QUANTILE | Yes | |
| TDIGEST.RANK | Yes | |
| TDIGEST.RESET | Yes | |
| TDIGEST.REVRANK | Yes | |
| TDIGEST.TRIMMED_MEAN | Yes | |
| TIME | No | |
| TOPK.ADD | Yes | |
| TOPK.COUNT | Yes | |
| TOPK.INCRBY | Yes | |
| TOPK.INFO | Yes | |
| TOPK.LIST | Yes | |
| TOPK.QUERY | Yes | |
| TOPK.RESERVE | Yes | |
| TOUCH | Yes | |
| TS.ADD | Yes | |
| TS.ALTER | Yes | |
| TS.CREATE | Yes | |
| TS.CREATERULE | Yes | |
| TS.DECRBY | Yes | |
| TS.DEL | Yes | |
| TS.DELETERULE | Yes | |
| TS.GET | Yes | |
| TS.INCRBY | Yes | |
| TS.INFO | Yes | |
| TS.MADD | Yes | |
| TS.MGET | No | |
| TS.MRANGE | No | |
| TS.MREVRANGE | No | |
| TS.QUERYINDEX | No | |
| TS.RANGE | Yes | |
| TS.REVRANGE | Yes | |
| TTL | Yes | |
| TYPE | Yes | |
| UNLINK | Yes | |
| UNSUBSCRIBE | No | |
| UNWATCH | No | |
| VADD | Yes | |
| VCARD | Yes | |
| VDIM | Yes | |
| VEMB | Yes | |
| VGETATTR | Yes | |
| VINFO | Yes | |
| VLINKS | Yes | |
| VRANDMEMBER | Yes | |
| VREM | Yes | |
| VSETATTR | Yes | |
| VSIM | Yes | |
| WAIT | No | |
| WAITAOF | No | |
| WATCH | Yes | |
| XACK | No | |
| XADD | No | |
| XAUTOCLAIM | No | |
| XCLAIM | No | |
| XDEL | No | |
| XGROUP CREATE | No | |
| XGROUP CREATECONSUMER | No | |
| XGROUP DELCONSUMER | No | |
| XGROUP DESTROY | No | |
| XGROUP SETID | No | |
| XINFO CONSUMERS | No | |
| XINFO GROUPS | No | |
| XINFO STREAM | No | |
| XLEN | No | |
| XPENDING | No | |
| XRANGE | No | |
| XREAD | No | |
| XREADGROUP | No | |
| XREVRANGE | No | |
| XSETID | No | |
| XTRIM | No | |
| ZADD | Yes | |
| ZCARD | Yes | |
| ZCOUNT | Yes | |
| ZDIFF | Yes | |
| ZDIFFSTORE | Yes | |
| ZINCRBY | Yes | |
| ZINTER | Yes | |
| ZINTERCARD | Yes | |
| ZINTERSTORE | Yes | |
| ZLEXCOUNT | Yes | |
| ZMPOP | Yes | |
| ZMSCORE | Yes | |
| ZPOPMAX | Yes | |
| ZPOPMIN | Yes | |
| ZRANDMEMBER | Yes | |
| ZRANGE | Yes | |
| ZRANGEBYLEX | Yes | |
| ZRANGEBYSCORE | Yes | |
| ZRANGESTORE | Yes | |
| ZRANK | Yes | |
| ZREM | Yes | |
| ZREMRANGEBYLEX | Yes | |
| ZREMRANGEBYRANK | Yes | |
| ZREMRANGEBYSCORE | Yes | |
| ZREVRANGE | Yes | |
| ZREVRANGEBYLEX | Yes | |
| ZREVRANGEBYSCORE | Yes | |
| ZREVRANK | Yes | |
| ZSCAN | Yes | |
| ZSCORE | Yes | |
| ZUNION | Yes | |
| ZUNIONSTORE | Yes | |