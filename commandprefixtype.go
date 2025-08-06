package redisprefix

// Redis命令参数前缀处理分类
type PrefixType uint8

const (
	PrefixNone PrefixType = iota
	PrefixAll
	PrefixExceptSecondAll
	PrefixOdd
	PrefixExceptLast
	PrefixSecond
	PrefixThird
	PrefixSecondAndThird
	PrefixFirstEveryThree
	PrefixNumkeys       // Numkeys
	PrefixNumkeysSecond // Numkeys
	PrefixNumkeysStore  // Numkeys store to other key
	PrefixSTREAMS       // STREAMS
	PrefixThirdNonEmpty // STREAMS
	PrefixJoinNextArg   // next arg
)

var CommandPrefixType = map[string]PrefixType{
	"ACL":                           PrefixJoinNextArg,
	"CLIENT":                        PrefixJoinNextArg,
	"CLUSTER":                       PrefixJoinNextArg,
	"FUNCTION":                      PrefixJoinNextArg,
	"LATENCY":                       PrefixJoinNextArg,
	"MEMORY":                        PrefixJoinNextArg,
	"MODULE":                        PrefixJoinNextArg,
	"OBJECT":                        PrefixJoinNextArg,
	"PUBSUB":                        PrefixJoinNextArg,
	"SCRIPT":                        PrefixJoinNextArg,
	"SLOWLOG":                       PrefixJoinNextArg,
	"XGROUP":                        PrefixJoinNextArg,
	"JSON.DEBUG":                    PrefixJoinNextArg,     // JSON.DEBUG
	"ACL CAT":                       PrefixNone,            // AACL CAT [category]
	"ACL DELUSER":                   PrefixNone,            // ACL DELUSER username [username ...]
	"ACL DRYRUN":                    PrefixNone,            // ACL DRYRUN username command [arg [arg ...]]
	"ACL GENPASS":                   PrefixNone,            // ACL GENPASS [bits]
	"ACL GETUSER":                   PrefixNone,            // ACL GETUSER username
	"ACL LIST":                      PrefixNone,            // ACL LIST
	"ACL LOAD":                      PrefixNone,            // ACL LOAD
	"ACL LOG":                       PrefixNone,            // ACL LOG [count | RESET]
	"ACL SAVE":                      PrefixNone,            // ACL SAVE
	"ACL SETUSER":                   PrefixNone,            // ACL SETUSER username [rule [rule ...]]
	"ACL USERS":                     PrefixNone,            // ACL USERS
	"ACL WHOAMI":                    PrefixNone,            // ACL WHOAMI
	"APPEND":                        PrefixSecond,          // APPEND key value
	"ASKING":                        PrefixNone,            // ASKING
	"AUTH":                          PrefixNone,            // AUTH [username] password
	"BF.ADD":                        PrefixSecond,          // BF.ADD key item
	"BF.CARD":                       PrefixSecond,          // BF.CARD key
	"BF.EXISTS":                     PrefixSecond,          // BF.EXISTS key item
	"BF.INFO":                       PrefixSecond,          // BF.INFO key [CAPACITY | SIZE | FILTERS | ITEMS | EXPANSION]
	"BF.INSERT":                     PrefixSecond,          // BF.INSERT key [CAPACITY capacity] [ERROR error]  [EXPANSION expansion] [NOCREATE] [NONSCALING] ITEMS item [item  ...]
	"BF.LOADCHUNK":                  PrefixSecond,          // BF.LOADCHUNK key iterator data
	"BF.MADD":                       PrefixSecond,          // BF.MADD key item [item ...]
	"BF.MEXISTS":                    PrefixSecond,          // BF.MEXISTS key item [item ...]
	"BF.RESERVE":                    PrefixSecond,          // BF.RESERVE key error_rate capacity [EXPANSION expansion]  [NONSCALING]
	"BF.SCANDUMP":                   PrefixSecond,          // BF.SCANDUMP key iterator
	"BGREWRITEAOF":                  PrefixNone,            // BGREWRITEAOF
	"BGSAVE":                        PrefixNone,            // BGSAVE [SCHEDULE]
	"BITCOUNT":                      PrefixSecond,          // BITCOUNT key [start end [BYTE | BIT]]
	"BITFIELD":                      PrefixSecond,          // BITFIELD key [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>]  <SET encoding offset value | INCRBY encoding offset increment>  [GET encoding offset | [OVERFLOW <WRAP | SAT | FAIL>]  <SET encoding offset value | INCRBY encoding offset increment>  ...]]
	"BITFIELD_RO":                   PrefixSecond,          // BITFIELD_RO key [GET encoding offset [GET encoding offset ...]]
	"BITOP":                         PrefixExceptSecondAll, // BITOP <AND | OR | XOR | NOT> destkey key [key ...]
	"BITPOS":                        PrefixSecond,          // BITPOS key bit [start [end [BYTE | BIT]]]
	"BLMOVE":                        PrefixSecondAndThird,  // BLMOVE source destination <LEFT | RIGHT> <LEFT | RIGHT> timeout
	"BLMPOP":                        PrefixNumkeysSecond,   // BLMPOP timeout numkeys key [key ...] <LEFT | RIGHT> [COUNT count]
	"BLPOP":                         PrefixExceptLast,      // BLPOP key [key ...] timeout
	"BRPOP":                         PrefixExceptLast,      // BRPOP key [key ...] timeout
	"BRPOPLPUSH":                    PrefixSecondAndThird,  // BRPOPLPUSH source destination timeout
	"BZMPOP":                        PrefixNumkeysSecond,   // BZMPOP timeout numkeys key [key ...] <MIN | MAX> [COUNT count]
	"BZPOPMAX":                      PrefixExceptLast,      // BZPOPMAX key [key ...] timeout
	"BZPOPMIN":                      PrefixExceptLast,      // BZPOPMIN key [key ...] timeout
	"CF.ADD":                        PrefixSecond,          // CF.ADD key item
	"CF.ADDNX":                      PrefixSecond,          // CF.ADDNX key item
	"CF.COUNT":                      PrefixSecond,          // CF.COUNT key item
	"CF.DEL":                        PrefixSecond,          // CF.DEL key item
	"CF.EXISTS":                     PrefixSecond,          // CF.EXISTS key item
	"CF.INFO":                       PrefixSecond,          // CF.INFO key
	"CF.INSERT":                     PrefixSecond,          // CF.INSERT key [CAPACITY capacity] [NOCREATE] ITEMS item [item ...]
	"CF.INSERTNX":                   PrefixSecond,          // CF.INSERTNX key [CAPACITY capacity] [NOCREATE] ITEMS item [item ...]
	"CF.LOADCHUNK":                  PrefixSecond,          // CF.LOADCHUNK key iterator data
	"CF.MEXISTS":                    PrefixSecond,          // CF.MEXISTS key item [item ...]
	"CF.RESERVE":                    PrefixSecond,          // CF.RESERVE key capacity [BUCKETSIZE bucketsize]  [MAXITERATIONS maxiterations] [EXPANSION expansion]
	"CF.SCANDUMP":                   PrefixSecond,          // CF.SCANDUMP key iterator
	"CLIENT CACHING":                PrefixNone,            // CLIENT CACHING <YES | NO>
	"CLIENT GETNAME":                PrefixNone,            // CLIENT GETNAME
	"CLIENT GETREDIR":               PrefixNone,            // CLIENT GETREDIR
	"CLIENT ID":                     PrefixNone,            // CLIENT ID
	"CLIENT INFO":                   PrefixNone,            // CLIENT INFO
	"CLIENT KILL":                   PrefixNone,            // CLIENT KILL <ip:port | <[ID client-id] | [TYPE <NORMAL | MASTER |  SLAVE | REPLICA | PUBSUB>] | [USER username] | [ADDR ip:port] |  [LADDR ip:port] | [SKIPME <YES | NO>] | [MAXAGE maxage]  [[ID client-id] | [TYPE <NORMAL | MASTER | SLAVE | REPLICA |  PUBSUB>] | [USER username] | [ADDR ip:port] | [LADDR ip:port] |  [SKIPME <YES | NO>] | [MAXAGE maxage] ...]>>
	"CLIENT LIST":                   PrefixNone,            // CLIENT LIST [TYPE <NORMAL | MASTER | REPLICA | PUBSUB>]  [ID client-id [client-id ...]]
	"CLIENT NO-EVICT":               PrefixNone,            // CLIENT NO-EVICT <ON | OFF>
	"CLIENT NO-TOUCH":               PrefixNone,            // CLIENT NO-TOUCH <ON | OFF>
	"CLIENT PAUSE":                  PrefixNone,            // CLIENT PAUSE timeout [WRITE | ALL]
	"CLIENT REPLY":                  PrefixNone,            // CLIENT REPLY <ON | OFF | SKIP>
	"CLIENT SETINFO":                PrefixNone,            // CLIENT SETINFO <LIB-NAME libname | LIB-VER libver>
	"CLIENT SETNAME":                PrefixNone,            // CLIENT SETNAME connection-name
	"CLIENT TRACKING":               PrefixNone,            // CLIENT TRACKING <ON | OFF> [REDIRECT client-id] [PREFIX prefix  [PREFIX prefix ...]] [BCAST] [OPTIN] [OPTOUT] [NOLOOP]
	"CLIENT TRACKINGINFO":           PrefixNone,            // CLIENT TRACKINGINFO
	"CLIENT UNBLOCK":                PrefixNone,            // CLIENT UNBLOCK client-id [TIMEOUT | ERROR]
	"CLIENT UNPAUSE":                PrefixNone,            // CLIENT UNPAUSE
	"CLUSTER ADDSLOTS":              PrefixNone,            // CLUSTER ADDSLOTS slot [slot ...]
	"CLUSTER ADDSLOTSRANGE":         PrefixNone,            // CLUSTER ADDSLOTSRANGE start-slot end-slot [start-slot end-slot ...]
	"CLUSTER BUMPEPOCH":             PrefixNone,            // CLUSTER BUMPEPOCH
	"CLUSTER COUNT-FAILURE-REPORTS": PrefixNone,            // CLUSTER COUNT-FAILURE-REPORTS node-id
	"CLUSTER COUNTKEYSINSLOT":       PrefixNone,            // CLUSTER COUNTKEYSINSLOT slot
	"CLUSTER DELSLOTS":              PrefixNone,            // CLUSTER DELSLOTS slot [slot ...]
	"CLUSTER DELSLOTSRANGE":         PrefixNone,            // CLUSTER DELSLOTSRANGE start-slot end-slot [start-slot end-slot ...]
	"CLUSTER FAILOVER":              PrefixNone,            // CLUSTER FAILOVER [FORCE | TAKEOVER]
	"CLUSTER FLUSHSLOTS":            PrefixNone,            // CLUSTER FLUSHSLOTS
	"CLUSTER FORGET":                PrefixNone,            // CLUSTER FORGET node-id
	"CLUSTER GETKEYSINSLOT":         PrefixNone,            // CLUSTER GETKEYSINSLOT slot count
	"CLUSTER INFO":                  PrefixNone,            // CLUSTER INFO
	"CLUSTER KEYSLOT":               PrefixThird,           // CLUSTER KEYSLOT key
	"CLUSTER LINKS":                 PrefixNone,            // CLUSTER LINKS
	"CLUSTER MEET":                  PrefixNone,            // CLUSTER MEET ip port [cluster-bus-port]
	"CLUSTER MYID":                  PrefixNone,            // CLUSTER MYID
	"CLUSTER MYSHARDID":             PrefixNone,            // CLUSTER MYSHARDID
	"CLUSTER NODES":                 PrefixNone,            // CLUSTER NODES
	"CLUSTER REPLICAS":              PrefixNone,            // CLUSTER REPLICAS node-id
	"CLUSTER REPLICATE":             PrefixNone,            // CLUSTER REPLICATE node-id
	"CLUSTER RESET":                 PrefixNone,            // CLUSTER RESET [HARD | SOFT]
	"CLUSTER SAVECONFIG":            PrefixNone,            // CLUSTER SAVECONFIG
	"CLUSTER SET-CONFIG-EPOCH":      PrefixNone,            // CLUSTER SET-CONFIG-EPOCH config-epoch
	"CLUSTER SETSLOT":               PrefixNone,            // CLUSTER SETSLOT slot <IMPORTING node-id | MIGRATING node-id |  NODE node-id | STABLE>
	"CLUSTER SHARDS":                PrefixNone,            // CLUSTER SHARDS
	"CLUSTER SLAVES":                PrefixNone,            // CLUSTER SLAVES node-id
	"CLUSTER SLOTS":                 PrefixNone,            // CLUSTER SLOTS
	"CMS.INCRBY":                    PrefixSecond,          // CMS.INCRBY key item increment [item increment ...]
	"CMS.INFO":                      PrefixSecond,          // CMS.INFO key
	"CMS.INITBYDIM":                 PrefixSecond,          // CMS.INITBYDIM key width depth
	"CMS.INITBYPROB":                PrefixNone,            // CMS.INITBYPROB key error probability
	"CMS.MERGE":                     PrefixNumkeysStore,    // CMS.MERGE destination numKeys source [source ...] [WEIGHTS weight  [weight ...]]
	"CMS.QUERY":                     PrefixSecond,          // CMS.QUERY key item [item ...]
	"COMMAND":                       PrefixNone,            // COMMAND
	"COMMAND COUNT":                 PrefixNone,            // COMMAND COUNT
	"COMMAND DOCS":                  PrefixNone,            // COMMAND DOCS [command-name [command-name ...]]
	"COMMAND GETKEYS":               PrefixNone,            // COMMAND GETKEYS command [arg [arg ...]]
	"COMMAND GETKEYSANDFLAGS":       PrefixNone,            // COMMAND GETKEYSANDFLAGS command [arg [arg ...]]
	"COMMAND INFO":                  PrefixNone,            // COMMAND INFO [command-name [command-name ...]]
	"COMMAND LIST":                  PrefixNone,            // COMMAND LIST [FILTERBY <MODULE module-name | ACLCAT category |  PATTERN pattern>]
	"CONFIG GET":                    PrefixNone,            // CONFIG GET parameter [parameter ...]
	"CONFIG RESETSTAT":              PrefixNone,            // CONFIG RESETSTAT
	"CONFIG REWRITE":                PrefixNone,            // CONFIG REWRITE
	"CONFIG SET":                    PrefixNone,            // CONFIG SET parameter value [parameter value ...]
	"COPY":                          PrefixSecondAndThird,  // COPY source destination [DB destination-db] [REPLACE]
	"DBSIZE":                        PrefixNone,            // DBSIZE
	"DECR":                          PrefixSecond,          // DECR key
	"DECRBY":                        PrefixSecond,          // DECRBY key decrement
	"DEL":                           PrefixAll,             // DEL key [key ...]
	"DISCARD":                       PrefixNone,            // DISCARD
	"DUMP":                          PrefixSecond,          // DUMP key
	"ECHO":                          PrefixNone,            // ECHO message
	"EVAL":                          PrefixNone,            // EVAL script numkeys [key [key ...]] [arg [arg ...]]
	"EVAL_RO":                       PrefixNone,            // EVAL_RO script numkeys [key [key ...]] [arg [arg ...]]
	"EVALSHA":                       PrefixNone,            // EVALSHA sha1 numkeys [key [key ...]] [arg [arg ...]]
	"EVALSHA_RO":                    PrefixNone,            // EVALSHA_RO sha1 numkeys [key [key ...]] [arg [arg ...]]
	"EXEC":                          PrefixNone,            // EXEC
	"EXISTS":                        PrefixAll,             // EXISTS key [key ...]
	"EXPIRE":                        PrefixSecond,          // EXPIRE key seconds [NX | XX | GT | LT]
	"EXPIREAT":                      PrefixSecond,          // EXPIREAT key unix-time-seconds [NX | XX | GT | LT]
	"EXPIRETIME":                    PrefixSecond,          // EXPIRETIME key
	"FAILOVER":                      PrefixNone,            // FAILOVER [TO host port [FORCE]] [ABORT] [TIMEOUT milliseconds]
	"FCALL":                         PrefixNumkeysSecond,   // FCALL function numkeys [key [key ...]] [arg [arg ...]]
	"FCALL_RO":                      PrefixNumkeysSecond,   // FCALL_RO function numkeys [key [key ...]] [arg [arg ...]]
	"FLUSHALL":                      PrefixNone,            // FLUSHALL [ASYNC | SYNC]
	"FLUSHDB":                       PrefixNone,            // FLUSHDB [ASYNC | SYNC]
	"FT._LIST":                      PrefixNone,            // FT._LIST
	"FT.AGGREGATE":                  PrefixNone,            // FT.AGGREGATE index query   [VERBATIM]   [LOAD count field [field ...]]   [TIMEOUT timeout]   [ GROUPBY nargs property [property ...] [ REDUCE function nargs arg [arg ...] [AS name] [ REDUCE function nargs arg [arg ...] [AS name] ...]] ...]]   [ SORTBY nargs [ property ASC | DESC [ property ASC | DESC ...]] [MAX num] [WITHCOUNT]   [ APPLY expression AS name [ APPLY expression AS name ...]]   [ LIMIT offset num]   [FILTER filter]   [ WITHCURSOR [COUNT read_size] [MAXIDLE idle_time]]   [ PARAMS nargs name value [ name value ...]]   [SCORER scorer]  [ADDSCORES]   [DIALECT dialect]
	"FT.ALIASADD":                   PrefixNone,            // FT.ALIASADD alias index
	"FT.ALIASDEL":                   PrefixNone,            // FT.ALIASDEL alias
	"FT.ALIASUPDATE":                PrefixNone,            // FT.ALIASUPDATE alias index
	"FT.ALTER":                      PrefixNone,            // FT.ALTER {index} [SKIPINITIALSCAN] SCHEMA ADD {attribute} {options} ...
	"FT.CONFIG GET":                 PrefixNone,            // FT.CONFIG GET option
	"FT.CONFIG SET":                 PrefixNone,            // FT.CONFIG SET option value
	"FT.CREATE":                     PrefixNone,            // FT.CREATE index   [ON HASH | JSON]   [PREFIX count prefix [prefix ...]]   [FILTER {filter}]  [LANGUAGE default_lang]   [LANGUAGE_FIELD lang_attribute]   [SCORE default_score]   [SCORE_FIELD score_attribute]   [PAYLOAD_FIELD payload_attribute]   [MAXTEXTFIELDS]   [TEMPORARY seconds]   [NOOFFSETS]   [NOHL]   [NOFIELDS]   [NOFREQS]   [STOPWORDS count [stopword ...]]   [SKIPINITIALSCAN]  SCHEMA field_name [AS alias] TEXT | TAG | NUMERIC | GEO | VECTOR | GEOSHAPE [ SORTABLE [UNF]]   [NOINDEX] [ field_name [AS alias] TEXT | TAG | NUMERIC | GEO | VECTOR | GEOSHAPE [ SORTABLE [UNF]] [NOINDEX] ...]
	"FT.CURSOR DEL":                 PrefixNone,            // FT.CURSOR DEL index cursor_id
	"FT.CURSOR READ":                PrefixNone,            // FT.CURSOR READ index cursor_id [COUNT read_size]
	"FT.DICTADD":                    PrefixNone,            // FT.DICTADD dict term [term ...]
	"FT.DICTDEL":                    PrefixNone,            // FT.DICTDEL dict term [term ...]
	"FT.DICTDUMP":                   PrefixNone,            // FT.DICTDUMP dict
	"FT.DROPINDEX":                  PrefixNone,            // FT.DROPINDEX index   [DD]
	"FT.EXPLAIN":                    PrefixNone,            // FT.EXPLAIN index query   [DIALECT dialect]
	"FT.EXPLAINCLI":                 PrefixNone,            // FT.EXPLAINCLI index query   [DIALECT dialect]
	"FT.INFO":                       PrefixNone,            // FT.INFO index
	"FT.PROFILE":                    PrefixNone,            // FT.PROFILE index SEARCH | AGGREGATE [LIMITED] QUERY query
	"FT.SEARCH":                     PrefixNone,            // FT.SEARCH index query   [NOCONTENT]   [VERBATIM]   [NOSTOPWORDS]   [WITHSCORES]   [WITHPAYLOADS]   [WITHSORTKEYS]   [FILTER numeric_field min max [ FILTER numeric_field min max ...]]   [GEOFILTER geo_field lon lat radius m | km | mi | ft [ GEOFILTER geo_field lon lat radius m | km | mi | ft ...]]   [INKEYS count key [key ...]]   [INFIELDS count field [field ...]]   [RETURN count identifier [AS property] [ identifier [AS property] ...]]   [SUMMARIZE [ FIELDS count field [field ...]] [FRAGS num] [LEN fragsize] [SEPARATOR separator]]   [HIGHLIGHT [ FIELDS count field [field ...]] [ TAGS open close]]   [SLOP slop]   [TIMEOUT timeout]   [INORDER]   [LANGUAGE language]   [EXPANDER expander]   [SCORER scorer]   [EXPLAINSCORE]   [PAYLOAD payload]   [SORTBY sortby [ ASC | DESC] [WITHCOUNT]]   [LIMIT offset num]   [PARAMS nargs name value [ name value ...]]   [DIALECT dialect]
	"FT.SPELLCHECK":                 PrefixNone,            // FT.SPELLCHECK index query   [DISTANCE distance]   [TERMS INCLUDE | EXCLUDE dictionary [terms [terms ...]]]   [DIALECT dialect]
	"FT.SUGADD":                     PrefixSecond,          // FT.SUGADD key string score   [INCR]   [PAYLOAD payload]
	"FT.SUGDEL":                     PrefixSecond,          // FT.SUGDEL key string
	"FT.SUGGET":                     PrefixSecond,          // FT.SUGGET key prefix   [FUZZY]   [WITHSCORES]   [WITHPAYLOADS]   [MAX max]
	"FT.SUGLEN":                     PrefixSecond,          // FT.SUGLEN key
	"FT.SYNDUMP":                    PrefixNone,            // FT.SYNDUMP index
	"FT.SYNUPDATE":                  PrefixNone,            // FT.SYNUPDATE index synonym_group_id   [SKIPINITIALSCAN] term [term ...]
	"FT.TAGVALS":                    PrefixNone,            // FT.TAGVALS index field_name
	"FUNCTION DELETE":               PrefixNone,            // FUNCTION DELETE library-name
	"FUNCTION DUMP":                 PrefixNone,            // FUNCTION DUMP
	"FUNCTION FLUSH":                PrefixNone,            // FUNCTION FLUSH [ASYNC | SYNC]
	"FUNCTION KILL":                 PrefixNone,            // FUNCTION KILL
	"FUNCTION LIST":                 PrefixNone,            // FUNCTION LIST [LIBRARYNAME library-name-pattern] [WITHCODE]
	"FUNCTION LOAD":                 PrefixNone,            // FUNCTION LOAD [REPLACE] function-code
	"FUNCTION RESTORE":              PrefixNone,            // FUNCTION RESTORE serialized-value [FLUSH | APPEND | REPLACE]
	"FUNCTION STATS":                PrefixNone,            // FUNCTION STATS
	"GEOADD":                        PrefixSecond,          // GEOADD key [NX | XX] [CH] longitude latitude member [longitude  latitude member ...]
	"GEODIST":                       PrefixSecond,          // GEODIST key member1 member2 [M | KM | FT | MI]
	"GEOHASH":                       PrefixSecond,          // GEOHASH key [member [member ...]]
	"GEOPOS":                        PrefixSecond,          // GEOPOS key [member [member ...]]
	"GEORADIUS":                     PrefixSecond,          // GEORADIUS key longitude latitude radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]  [STORE key | STOREDIST key]
	"GEORADIUS_RO":                  PrefixSecond,          // GEORADIUS_RO key longitude latitude radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]
	"GEORADIUSBYMEMBER":             PrefixSecond,          // GEORADIUSBYMEMBER key member radius <M | KM | FT | MI> [WITHCOORD]  [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC] [STORE key  | STOREDIST key]
	"GEORADIUSBYMEMBER_RO":          PrefixSecond,          // GEORADIUSBYMEMBER_RO key member radius <M | KM | FT | MI>  [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count [ANY]] [ASC | DESC]
	"GEOSEARCH":                     PrefixSecond,          // GEOSEARCH key <FROMMEMBER member | FROMLONLAT longitude latitude>  <BYRADIUS radius <M | KM | FT | MI> | BYBOX width height <M | KM |  FT | MI>> [ASC | DESC] [COUNT count [ANY]] [WITHCOORD] [WITHDIST]  [WITHHASH]
	"GEOSEARCHSTORE":                PrefixSecondAndThird,  // GEOSEARCHSTORE destination source <FROMMEMBER member |  FROMLONLAT longitude latitude> <BYRADIUS radius <M | KM | FT | MI>  | BYBOX width height <M | KM | FT | MI>> [ASC | DESC] [COUNT count  [ANY]] [STOREDIST]
	"GET":                           PrefixSecond,          // GET key
	"GETBIT":                        PrefixSecond,          // GETBIT key offset
	"GETDEL":                        PrefixSecond,          // GETDEL key
	"GETEX":                         PrefixSecond,          // GETEX key [EX seconds | PX milliseconds | EXAT unix-time-seconds |  PXAT unix-time-milliseconds | PERSIST]
	"GETRANGE":                      PrefixSecond,          // GETRANGE key start end
	"GETSET":                        PrefixSecond,          // GETSET key value
	"HDEL":                          PrefixSecond,          // HDEL key field [field ...]
	"HELLO":                         PrefixNone,            // HELLO [protover [AUTH username password] [SETNAME clientname]]
	"HEXISTS":                       PrefixSecond,          // HEXISTS key field
	"HEXPIRE":                       PrefixSecond,          // HEXPIRE key seconds [NX | XX | GT | LT] FIELDS numfields field  [field ...]
	"HEXPIREAT":                     PrefixSecond,          // HEXPIREAT key unix-time-seconds [NX | XX | GT | LT] FIELDS numfields  field [field ...]
	"HEXPIRETIME":                   PrefixSecond,          // HEXPIRETIME key FIELDS numfields field [field ...]
	"HGET":                          PrefixSecond,          // HGET key field
	"HGETALL":                       PrefixSecond,          // HGETALL key
	"HGETDEL":                       PrefixSecond,          // HGETDEL key FIELDS numfields field [field ...]
	"HGETEX":                        PrefixSecond,          // HGETEX key [EX seconds | PX milliseconds | EXAT unix-time-seconds |  PXAT unix-time-milliseconds | PERSIST] FIELDS numfields field  [field ...]
	"HINCRBY":                       PrefixSecond,          // HINCRBY key field increment
	"HINCRBYFLOAT":                  PrefixSecond,          // HINCRBYFLOAT key field increment
	"HKEYS":                         PrefixSecond,          // HKEYS key
	"HLEN":                          PrefixSecond,          // HLEN key
	"HMGET":                         PrefixSecond,          // HMGET key field [field ...]
	"HMSET":                         PrefixSecond,          // HMSET key field value [field value ...]
	"HPERSIST":                      PrefixSecond,          // HPERSIST key FIELDS numfields field [field ...]
	"HPEXPIRE":                      PrefixSecond,          // HPEXPIRE key milliseconds [NX | XX | GT | LT] FIELDS numfields field  [field ...]
	"HPEXPIREAT":                    PrefixSecond,          // HPEXPIREAT key unix-time-milliseconds [NX | XX | GT | LT]  FIELDS numfields field [field ...]
	"HPEXPIRETIME":                  PrefixSecond,          // HPEXPIRETIME key FIELDS numfields field [field ...]
	"HPTTL":                         PrefixSecond,          // HPTTL key FIELDS numfields field [field ...]
	"HRANDFIELD":                    PrefixSecond,          // HRANDFIELD key [count [WITHVALUES]]
	"HSCAN":                         PrefixSecond,          // HSCAN key cursor [MATCH pattern] [COUNT count] [NOVALUES]
	"HSET":                          PrefixSecond,          // HSET key field value [field value ...]
	"HSETEX":                        PrefixSecond,          // HSETEX key [FNX | FXX] [EX seconds | PX milliseconds |  EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]  FIELDS numfields field value [field value ...]
	"HSETNX":                        PrefixSecond,          // HSETNX key field value
	"HSTRLEN":                       PrefixSecond,          // HSTRLEN key field
	"HTTL":                          PrefixSecond,          // HTTL key FIELDS numfields field [field ...]
	"HVALS":                         PrefixSecond,          // HVALS key
	"INCR":                          PrefixSecond,          // INCR key
	"INCRBY":                        PrefixSecond,          // INCRBY key increment
	"INCRBYFLOAT":                   PrefixSecond,          // INCRBYFLOAT key increment
	"INFO":                          PrefixNone,            // INFO [section [section ...]]
	"JSON.ARRAPPEND":                PrefixSecond,          // JSON.ARRAPPEND key path value [value ...]
	"JSON.ARRINDEX":                 PrefixSecond,          // JSON.ARRINDEX key path value [start [stop]]
	"JSON.ARRINSERT":                PrefixSecond,          // JSON.ARRINSERT key path index value [value ...]
	"JSON.ARRLEN":                   PrefixSecond,          // JSON.ARRLEN key [path]
	"JSON.ARRPOP":                   PrefixSecond,          // JSON.ARRPOP key [path [index]]
	"JSON.ARRTRIM":                  PrefixSecond,          // JSON.ARRTRIM key path start stop
	"JSON.CLEAR":                    PrefixSecond,          // JSON.CLEAR key [path]
	"JSON.DEBUG MEMORY":             PrefixThird,           // JSON.DEBUG MEMORY key [path]
	"JSON.DEL":                      PrefixSecond,          // JSON.DEL key [path]
	"JSON.FORGET":                   PrefixSecond,          // JSON.FORGET key [path]
	"JSON.GET":                      PrefixSecond,          // JSON.GET key [INDENT indent] [NEWLINE newline] [SPACE space] [path  [path ...]]
	"JSON.MERGE":                    PrefixSecond,          // JSON.MERGE key path value
	"JSON.MGET":                     PrefixExceptLast,      // JSON.MGET key [key ...] path
	"JSON.MSET":                     PrefixFirstEveryThree, // JSON.MSET key path value [key path value ...]
	"JSON.NUMINCRBY":                PrefixSecond,          // JSON.NUMINCRBY key path value
	"JSON.NUMMULTBY":                PrefixSecond,          // JSON.NUMMULTBY key path value
	"JSON.OBJKEYS":                  PrefixSecond,          // JSON.OBJKEYS key [path]
	"JSON.OBJLEN":                   PrefixSecond,          // JSON.OBJLEN key [path]
	"JSON.RESP":                     PrefixSecond,          // JSON.RESP key [path]
	"JSON.SET":                      PrefixSecond,          // JSON.SET key path value [NX | XX]
	"JSON.STRAPPEND":                PrefixSecond,          // JSON.STRAPPEND key [path] value
	"JSON.STRLEN":                   PrefixSecond,          // JSON.STRLEN key [path]
	"JSON.TOGGLE":                   PrefixSecond,          // JSON.TOGGLE key path
	"JSON.TYPE":                     PrefixSecond,          // JSON.TYPE key [path]
	"KEYS":                          PrefixNone,            // KEYS pattern
	"LASTSAVE":                      PrefixNone,            // LASTSAVE
	"LATENCY DOCTOR":                PrefixNone,            // LATENCY DOCTOR
	"LATENCY GRAPH":                 PrefixNone,            // LATENCY GRAPH event
	"LATENCY HISTOGRAM":             PrefixNone,            // LATENCY HISTOGRAM [command [command ...]]
	"LATENCY HISTORY":               PrefixNone,            // LATENCY HISTORY event
	"LATENCY LATEST":                PrefixNone,            // LATENCY LATEST
	"LATENCY RESET":                 PrefixNone,            // LATENCY RESET [event [event ...]]
	"LCS":                           PrefixSecondAndThird,  // LCS key1 key2 [LEN] [IDX] [MINMATCHLEN min-match-len] [WITHMATCHLEN]
	"LINDEX":                        PrefixSecond,          // LINDEX key index
	"LINSERT":                       PrefixSecond,          // LINSERT key <BEFORE | AFTER> pivot element
	"LLEN":                          PrefixSecond,          // LLEN key
	"LMOVE":                         PrefixSecondAndThird,  // LMOVE source destination <LEFT | RIGHT> <LEFT | RIGHT>
	"LMPOP":                         PrefixNumkeys,         // LMPOP numkeys key [key ...] <LEFT | RIGHT> [COUNT count]
	"LOLWUT":                        PrefixNone,            // LOLWUT [VERSION version]
	"LPOP":                          PrefixSecond,          // LPOP key [count]
	"LPOS":                          PrefixSecond,          // LPOS key element [RANK rank] [COUNT num-matches] [MAXLEN len]
	"LPUSH":                         PrefixSecond,          // LPUSH key element [element ...]
	"LPUSHX":                        PrefixSecond,          // LPUSHX key element [element ...]
	"LRANGE":                        PrefixSecond,          // LRANGE key start stop
	"LREM":                          PrefixSecond,          // LREM key count element
	"LSET":                          PrefixSecond,          // LSET key index element
	"LTRIM":                         PrefixSecond,          // LTRIM key start stop
	"MEMORY DOCTOR":                 PrefixNone,            // MEMORY DOCTOR
	"MEMORY MALLOC-STATS":           PrefixNone,            // MEMORY MALLOC-STATS
	"MEMORY PURGE":                  PrefixNone,            // MEMORY PURGE
	"MEMORY STATS":                  PrefixNone,            // MEMORY STATS
	"MEMORY USAGE":                  PrefixThird,           // MEMORY USAGE key [SAMPLES count]
	"MGET":                          PrefixAll,             // MGET key [key ...]
	"MIGRATE":                       PrefixThirdNonEmpty,   // MIGRATE host port <key | ""> destination-db timeout [COPY] [REPLACE]  [AUTH password | AUTH2 username password] [KEYS key [key ...]]
	"MODULE LIST":                   PrefixNone,            // MODULE LIST
	"MODULE LOAD":                   PrefixNone,            // MODULE LOAD path [arg [arg ...]]
	"MODULE LOADEX":                 PrefixNone,            // MODULE LOADEX path [CONFIG name value [CONFIG name value ...]]  [ARGS args [args ...]]
	"MODULE UNLOAD":                 PrefixNone,            // MODULE UNLOAD name
	"MONITOR":                       PrefixNone,            // MONITOR
	"MOVE":                          PrefixSecond,          // MOVE key db
	"MSET":                          PrefixOdd,             // MSET key value [key value ...]
	"MSETNX":                        PrefixOdd,             // MSETNX key value [key value ...]
	"MULTI":                         PrefixNone,            // MULTI
	"OBJECT ENCODING":               PrefixThird,           // OBJECT ENCODING key
	"OBJECT FREQ":                   PrefixThird,           // OBJECT FREQ key
	"OBJECT IDLETIME":               PrefixThird,           // OBJECT IDLETIME key
	"OBJECT REFCOUNT":               PrefixThird,           // OBJECT REFCOUNT key
	"PERSIST":                       PrefixSecond,          // PERSIST key
	"PEXPIRE":                       PrefixSecond,          // PEXPIRE key milliseconds [NX | XX | GT | LT]
	"PEXPIREAT":                     PrefixSecond,          // PEXPIREAT key unix-time-milliseconds [NX | XX | GT | LT]
	"PEXPIRETIME":                   PrefixSecond,          // PEXPIRETIME key
	"PFADD":                         PrefixSecond,          // PFADD key [element [element ...]]
	"PFCOUNT":                       PrefixAll,             // PFCOUNT key [key ...]
	"PFDEBUG":                       PrefixThird,           // PFDEBUG subcommand key
	"PFMERGE":                       PrefixAll,             // PFMERGE destkey [sourcekey [sourcekey ...]]
	"PFSELFTEST":                    PrefixNone,            // PFSELFTEST
	"PING":                          PrefixNone,            // PING [message]
	"PSETEX":                        PrefixSecond,          // PSETEX key milliseconds value
	"PSUBSCRIBE":                    PrefixNone,            // PSUBSCRIBE pattern [pattern ...]
	"PSYNC":                         PrefixNone,            // PSYNC replicationid offset
	"PTTL":                          PrefixSecond,          // PTTL key
	"PUBLISH":                       PrefixNone,            // PUBLISH channel message
	"PUBSUB CHANNELS":               PrefixNone,            // PUBSUB CHANNELS [pattern]
	"PUBSUB NUMPAT":                 PrefixNone,            // PUBSUB NUMPAT
	"PUBSUB NUMSUB":                 PrefixNone,            // PUBSUB NUMSUB [channel [channel ...]]
	"PUBSUB SHARDCHANNELS":          PrefixNone,            // PUBSUB SHARDCHANNELS [pattern]
	"PUBSUB SHARDNUMSUB":            PrefixNone,            // PUBSUB SHARDNUMSUB [shardchannel [shardchannel ...]]
	"PUNSUBSCRIBE":                  PrefixNone,            // PUNSUBSCRIBE [pattern [pattern ...]]
	"QUIT":                          PrefixNone,            // QUIT
	"RANDOMKEY":                     PrefixNone,            // RANDOMKEY
	"READONLY":                      PrefixNone,            // READONLY
	"READWRITE":                     PrefixNone,            // READWRITE
	"RENAME":                        PrefixAll,             // RENAME key newkey
	"RENAMENX":                      PrefixAll,             // RENAMENX key newkey
	"REPLCONF":                      PrefixNone,            // REPLCONF
	"REPLICAOF":                     PrefixNone,            // REPLICAOF <host port | NO ONE>
	"RESET":                         PrefixNone,            // RESET
	"RESTORE":                       PrefixSecond,          // RESTORE key ttl serialized-value [REPLACE] [ABSTTL]  [IDLETIME seconds] [FREQ frequency]
	"RESTORE-ASKING":                PrefixSecond,          // RESTORE-ASKING key ttl serialized-value [REPLACE] [ABSTTL]  [IDLETIME seconds] [FREQ frequency]
	"ROLE":                          PrefixNone,            // ROLE
	"RPOP":                          PrefixSecond,          // RPOP key [count]
	"RPOPLPUSH":                     PrefixSecondAndThird,  // RPOPLPUSH source destination
	"RPUSH":                         PrefixSecond,          // RPUSH key element [element ...]
	"RPUSHX":                        PrefixSecond,          // RPUSHX key element [element ...]
	"SADD":                          PrefixSecond,          // SADD key member [member ...]
	"SAVE":                          PrefixNone,            // SAVE
	"SCAN":                          PrefixNone,            // SCAN cursor [MATCH pattern] [COUNT count] [TYPE type]
	"SCARD":                         PrefixSecond,          // SCARD key
	"SCRIPT DEBUG":                  PrefixNone,            // SCRIPT DEBUG <YES | SYNC | NO>
	"SCRIPT EXISTS":                 PrefixNone,            // SCRIPT EXISTS sha1 [sha1 ...]
	"SCRIPT FLUSH":                  PrefixNone,            // SCRIPT FLUSH [ASYNC | SYNC]
	"SCRIPT KILL":                   PrefixNone,            // SCRIPT KILL
	"SCRIPT LOAD":                   PrefixNone,            // SCRIPT LOAD script
	"SDIFF":                         PrefixAll,             // SDIFF key [key ...]
	"SDIFFSTORE":                    PrefixAll,             // SDIFFSTORE destination key [key ...]
	"SELECT":                        PrefixNone,            // SELECT index
	"SET":                           PrefixSecond,          // SET key value [NX | XX] [GET] [EX seconds | PX milliseconds |  EXAT unix-time-seconds | PXAT unix-time-milliseconds | KEEPTTL]
	"SETBIT":                        PrefixSecond,          // SETBIT key offset value
	"SETEX":                         PrefixSecond,          // SETEX key seconds value
	"SETNX":                         PrefixSecond,          // SETNX key value
	"SETRANGE":                      PrefixSecond,          // SETRANGE key offset value
	"SHUTDOWN":                      PrefixNone,            // SHUTDOWN [NOSAVE | SAVE] [NOW] [FORCE] [ABORT]
	"SINTER":                        PrefixAll,             // SINTER key [key ...]
	"SINTERCARD":                    PrefixNumkeys,         // SINTERCARD numkeys key [key ...] [LIMIT limit]
	"SINTERSTORE":                   PrefixAll,             // SINTERSTORE destination key [key ...]
	"SISMEMBER":                     PrefixSecond,          // SISMEMBER key member
	"SLAVEOF":                       PrefixNone,            // SLAVEOF <host port | NO ONE>
	"SLOWLOG GET":                   PrefixNone,            // SLOWLOG GET [count]
	"SLOWLOG LEN":                   PrefixNone,            // SLOWLOG LEN
	"SLOWLOG RESET":                 PrefixNone,            // SLOWLOG RESET
	"SMEMBERS":                      PrefixSecond,          // SMEMBERS key
	"SMISMEMBER":                    PrefixSecond,          // SMISMEMBER key member [member ...]
	"SMOVE":                         PrefixSecondAndThird,  // SMOVE source destination member
	"SORT":                          PrefixSecond,          // SORT key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern  ...]] [ASC | DESC] [ALPHA] [STORE destination]
	"SORT_RO":                       PrefixSecond,          // SORT_RO key [BY pattern] [LIMIT offset count] [GET pattern [GET  pattern ...]] [ASC | DESC] [ALPHA]
	"SPOP":                          PrefixSecond,          // SPOP key [count]
	"SPUBLISH":                      PrefixNone,            // SPUBLISH shardchannel message
	"SRANDMEMBER":                   PrefixSecond,          // SRANDMEMBER key [count]
	"SREM":                          PrefixSecond,          // SREM key member [member ...]
	"SSCAN":                         PrefixSecond,          // SSCAN key cursor [MATCH pattern] [COUNT count]
	"SSUBSCRIBE":                    PrefixNone,            // SSUBSCRIBE shardchannel [shardchannel ...]
	"STRLEN":                        PrefixSecond,          // STRLEN key
	"SUBSCRIBE":                     PrefixNone,            // SUBSCRIBE channel [channel ...]
	"SUBSTR":                        PrefixSecond,          // SUBSTR key start end
	"SUNION":                        PrefixAll,             // SUNION key [key ...]
	"SUNIONSTORE":                   PrefixAll,             // SUNIONSTORE destination key [key ...]
	"SUNSUBSCRIBE":                  PrefixNone,            // SUNSUBSCRIBE [shardchannel [shardchannel ...]]
	"SWAPDB":                        PrefixNone,            // SWAPDB index1 index2
	"SYNC":                          PrefixNone,            // SYNC
	"TDIGEST.ADD":                   PrefixSecond,          // TDIGEST.ADD key value [value ...]
	"TDIGEST.BYRANK":                PrefixSecond,          // TDIGEST.BYRANK key rank [rank ...]
	"TDIGEST.BYREVRANK":             PrefixSecond,          // TDIGEST.BYREVRANK key reverse_rank [reverse_rank ...]
	"TDIGEST.CDF":                   PrefixSecond,          // TDIGEST.CDF key value [value ...]
	"TDIGEST.CREATE":                PrefixSecond,          // TDIGEST.CREATE key [COMPRESSION compression]
	"TDIGEST.INFO":                  PrefixSecond,          // TDIGEST.INFO key
	"TDIGEST.MAX":                   PrefixSecond,          // TDIGEST.MAX key
	"TDIGEST.MERGE":                 PrefixNumkeysStore,    // TDIGEST.MERGE destination-key numkeys source-key [source-key ...]  [COMPRESSION compression] [OVERRIDE]
	"TDIGEST.MIN":                   PrefixSecond,          // TDIGEST.MIN key
	"TDIGEST.QUANTILE":              PrefixSecond,          // TDIGEST.QUANTILE key quantile [quantile ...]
	"TDIGEST.RANK":                  PrefixSecond,          // TDIGEST.RANK key value [value ...]
	"TDIGEST.RESET":                 PrefixSecond,          // TDIGEST.RESET key
	"TDIGEST.REVRANK":               PrefixSecond,          // TDIGEST.REVRANK key value [value ...]
	"TDIGEST.TRIMMED_MEAN":          PrefixSecond,          // TDIGEST.TRIMMED_MEAN key low_cut_quantile high_cut_quantile
	"TIME":                          PrefixNone,            // TIME
	"TOPK.ADD":                      PrefixSecond,          // TOPK.ADD key items [items ...]
	"TOPK.COUNT":                    PrefixSecond,          // TOPK.COUNT key item [item ...]
	"TOPK.INCRBY":                   PrefixSecond,          // TOPK.INCRBY key item increment [item increment ...]
	"TOPK.INFO":                     PrefixSecond,          // TOPK.INFO key
	"TOPK.LIST":                     PrefixSecond,          // TOPK.LIST key [WITHCOUNT]
	"TOPK.QUERY":                    PrefixSecond,          // TOPK.QUERY key item [item ...]
	"TOPK.RESERVE":                  PrefixSecond,          // TOPK.RESERVE key topk [width depth decay]
	"TOUCH":                         PrefixAll,             // TOUCH key [key ...]
	"TS.ADD":                        PrefixSecond,          // TS.ADD key timestamp value   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [ON_DUPLICATE policy_ovr]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]   [LABELS [label value ...]]
	"TS.ALTER":                      PrefixSecond,          // TS.ALTER key   [RETENTION retentionPeriod]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]   [LABELS [label value ...]]
	"TS.CREATE":                     PrefixSecond,          // TS.CREATE key   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]   [LABELS [label value ...]]
	"TS.CREATERULE":                 PrefixSecondAndThird,  // TS.CREATERULE sourceKey destKey   AGGREGATION aggregator bucketDuration   [alignTimestamp]
	"TS.DECRBY":                     PrefixSecond,          // TS.DECRBY key subtrahend   [TIMESTAMP timestamp]   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]    [LABELS [label value ...]]
	"TS.DEL":                        PrefixSecond,          // TS.DEL key fromTimestamp toTimestamp
	"TS.DELETERULE":                 PrefixAll,             // TS.DELETERULE sourceKey destKey
	"TS.GET":                        PrefixSecond,          // TS.GET key   [LATEST]
	"TS.INCRBY":                     PrefixSecond,          // TS.INCRBY key addend   [TIMESTAMP timestamp]   [RETENTION retentionPeriod]   [ENCODING <COMPRESSED|UNCOMPRESSED>]   [CHUNK_SIZE size]   [DUPLICATE_POLICY policy]   [IGNORE ignoreMaxTimediff ignoreMaxValDiff]     [LABELS [label value ...]]
	"TS.INFO":                       PrefixSecond,          // TS.INFO key   [DEBUG]
	"TS.MADD":                       PrefixFirstEveryThree, // TS.MADD {key timestamp value}...
	"TS.MGET":                       PrefixNone,            // TS.MGET [LATEST] [WITHLABELS | <SELECTED_LABELS label...>] FILTER filterExpr...
	"TS.MRANGE":                     PrefixNone,            // TS.MRANGE fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [WITHLABELS | <SELECTED_LABELS label...>]  [COUNT count]  [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]  FILTER filterExpr...  [GROUPBY label REDUCE reducer]
	"TS.MREVRANGE":                  PrefixNone,            // TS.MREVRANGE fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [WITHLABELS | <SELECTED_LABELS label...>]  [COUNT count]  [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]  FILTER filterExpr...  [GROUPBY label REDUCE reducer]
	"TS.QUERYINDEX":                 PrefixNone,            // TS.QUERYINDEX filterExpr...
	"TS.RANGE":                      PrefixSecond,          // TS.RANGE key fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [COUNT count]   [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]
	"TS.REVRANGE":                   PrefixSecond,          // TS.REVRANGE key fromTimestamp toTimestamp  [LATEST]  [FILTER_BY_TS ts...]  [FILTER_BY_VALUE min max]  [COUNT count]  [[ALIGN align] AGGREGATION aggregator bucketDuration [BUCKETTIMESTAMP bt] [EMPTY]]
	"TTL":                           PrefixSecond,          // TTL key
	"TYPE":                          PrefixSecond,          // TYPE key
	"UNLINK":                        PrefixAll,             // UNLINK key [key ...]
	"UNSUBSCRIBE":                   PrefixNone,            // UNSUBSCRIBE [channel [channel ...]]
	"UNWATCH":                       PrefixNone,            // UNWATCH
	"VADD":                          PrefixSecond,          // VADD key [REDUCE dim] (FP32 | VALUES num) vector element [CAS] [NOQUANT | Q8 | BIN]  [EF build-exploration-factor] [SETATTR attributes] [M numlinks]
	"VCARD":                         PrefixSecond,          // VCARD key
	"VDIM":                          PrefixSecond,          // VDIM key
	"VEMB":                          PrefixSecond,          // VEMB key element [RAW]
	"VGETATTR":                      PrefixSecond,          // VGETATTR key element
	"VINFO":                         PrefixSecond,          // VINFO key
	"VLINKS":                        PrefixSecond,          // VLINKS key element [WITHSCORES]
	"VRANDMEMBER":                   PrefixSecond,          // VRANDMEMBER key [count]
	"VREM":                          PrefixSecond,          // VREM key element
	"VSETATTR":                      PrefixSecond,          // VSETATTR key element "{ JSON obj }"
	"VSIM":                          PrefixSecond,          // VSIM key (ELE | FP32 | VALUES num) (vector | element) [WITHSCORES] [COUNT num] [EF search-exploration-factor]  [FILTER expression] [FILTER-EF max-filtering-effort] [TRUTH] [NOTHREAD]
	"WAIT":                          PrefixNone,            // WAIT numreplicas timeout
	"WAITAOF":                       PrefixNone,            // WAITAOF numlocal numreplicas timeout
	"WATCH":                         PrefixAll,             // WATCH key [key ...]
	"XACK":                          PrefixNone,            // XACK key group id [id ...]
	"XADD":                          PrefixNone,            // XADD key [NOMKSTREAM] [<MAXLEN | MINID> [= | ~] threshold  [LIMIT count]] <* | id> field value [field value ...]
	"XAUTOCLAIM":                    PrefixNone,            // XAUTOCLAIM key group consumer min-idle-time start [COUNT count]  [JUSTID]
	"XCLAIM":                        PrefixNone,            // XCLAIM key group consumer min-idle-time id [id ...] [IDLE ms]  [TIME unix-time-milliseconds] [RETRYCOUNT count] [FORCE] [JUSTID]  [LASTID lastid]
	"XDEL":                          PrefixNone,            // XDEL key id [id ...]
	"XGROUP CREATE":                 PrefixNone,            // XGROUP CREATE key group <id | $> [MKSTREAM]  [ENTRIESREAD entries-read]
	"XGROUP CREATECONSUMER":         PrefixNone,            // XGROUP CREATECONSUMER key group consumer
	"XGROUP DELCONSUMER":            PrefixNone,            // XGROUP DELCONSUMER key group consumer
	"XGROUP DESTROY":                PrefixNone,            // XGROUP DESTROY key group
	"XGROUP SETID":                  PrefixNone,            // XGROUP SETID key group <id | $> [ENTRIESREAD entries-read]
	"XINFO CONSUMERS":               PrefixNone,            // XINFO CONSUMERS key group
	"XINFO GROUPS":                  PrefixNone,            // XINFO GROUPS key
	"XINFO STREAM":                  PrefixNone,            // XINFO STREAM key [FULL [COUNT count]]
	"XLEN":                          PrefixNone,            // XLEN key
	"XPENDING":                      PrefixNone,            // XPENDING key group [[IDLE min-idle-time] start end count [consumer]]
	"XRANGE":                        PrefixNone,            // XRANGE key start end [COUNT count]
	"XREAD":                         PrefixNone,            // XREAD [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] id  [id ...]
	"XREADGROUP":                    PrefixNone,            // XREADGROUP GROUP group consumer [COUNT count] [BLOCK milliseconds]  [NOACK] STREAMS key [key ...] id [id ...]
	"XREVRANGE":                     PrefixNone,            // XREVRANGE key end start [COUNT count]
	"XSETID":                        PrefixNone,            // XSETID key last-id [ENTRIESADDED entries-added]  [MAXDELETEDID max-deleted-id]
	"XTRIM":                         PrefixNone,            // XTRIM key <MAXLEN | MINID> [= | ~] threshold [LIMIT count]
	"ZADD":                          PrefixSecond,          // ZADD key [NX | XX] [GT | LT] [CH] [INCR] score member [score member  ...]
	"ZCARD":                         PrefixSecond,          // ZCARD key
	"ZCOUNT":                        PrefixSecond,          // ZCOUNT key min max
	"ZDIFF":                         PrefixNumkeys,         // ZDIFF numkeys key [key ...] [WITHSCORES]
	"ZDIFFSTORE":                    PrefixNumkeysStore,    // ZDIFFSTORE destination numkeys key [key ...]
	"ZINCRBY":                       PrefixSecond,          // ZINCRBY key increment member
	"ZINTER":                        PrefixNumkeys,         // ZINTER numkeys key [key ...] [WEIGHTS weight [weight ...]]  [AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
	"ZINTERCARD":                    PrefixNumkeys,         // ZINTERCARD numkeys key [key ...] [LIMIT limit]
	"ZINTERSTORE":                   PrefixNumkeysStore,    // ZINTERSTORE destination numkeys key [key ...] [WEIGHTS weight  [weight ...]] [AGGREGATE <SUM | MIN | MAX>]
	"ZLEXCOUNT":                     PrefixSecond,          // ZLEXCOUNT key min max
	"ZMPOP":                         PrefixNumkeys,         // ZMPOP numkeys key [key ...] <MIN | MAX> [COUNT count]
	"ZMSCORE":                       PrefixSecond,          // ZMSCORE key member [member ...]
	"ZPOPMAX":                       PrefixSecond,          // ZPOPMAX key [count]
	"ZPOPMIN":                       PrefixSecond,          // ZPOPMIN key [count]
	"ZRANDMEMBER":                   PrefixSecond,          // ZRANDMEMBER key [count [WITHSCORES]]
	"ZRANGE":                        PrefixSecond,          // ZRANGE key start stop [BYSCORE | BYLEX] [REV] [LIMIT offset count]  [WITHSCORES]
	"ZRANGEBYLEX":                   PrefixSecond,          // ZRANGEBYLEX key min max [LIMIT offset count]
	"ZRANGEBYSCORE":                 PrefixSecond,          // ZRANGEBYSCORE key min max [WITHSCORES] [LIMIT offset count]
	"ZRANGESTORE":                   PrefixSecondAndThird,  // ZRANGESTORE dst src min max [BYSCORE | BYLEX] [REV] [LIMIT offset  count]
	"ZRANK":                         PrefixSecond,          // ZRANK key member [WITHSCORE]
	"ZREM":                          PrefixSecond,          // ZREM key member [member ...]
	"ZREMRANGEBYLEX":                PrefixSecond,          // ZREMRANGEBYLEX key min max
	"ZREMRANGEBYRANK":               PrefixSecond,          // ZREMRANGEBYRANK key start stop
	"ZREMRANGEBYSCORE":              PrefixSecond,          // ZREMRANGEBYSCORE key min max
	"ZREVRANGE":                     PrefixSecond,          // ZREVRANGE key start stop [WITHSCORES]
	"ZREVRANGEBYLEX":                PrefixSecond,          // ZREVRANGEBYLEX key max min [LIMIT offset count]
	"ZREVRANGEBYSCORE":              PrefixSecond,          // ZREVRANGEBYSCORE key max min [WITHSCORES] [LIMIT offset count]
	"ZREVRANK":                      PrefixSecond,          // ZREVRANK key member [WITHSCORE]
	"ZSCAN":                         PrefixSecond,          // ZSCAN key cursor [MATCH pattern] [COUNT count]
	"ZSCORE":                        PrefixSecond,          // ZSCORE key member
	"ZUNION":                        PrefixNumkeys,         // ZUNION numkeys key [key ...] [WEIGHTS weight [weight ...]]  [AGGREGATE <SUM | MIN | MAX>] [WITHSCORES]
	"ZUNIONSTORE":                   PrefixNumkeysStore,    // ZUNIONSTORE destination numkeys key [key ...] [WEIGHTS weight  [weight ...]] [AGGREGATE <SUM | MIN | MAX>]
}
