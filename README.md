# BIGRED

[![Go Report Card](https://goreportcard.com/badge/github.com/AdhityaRamadhanus/bigred)](https://goreportcard.com/report/github.com/AdhityaRamadhanus/bigred)

Cache server with [Big Cache](https://github.com/allegro/bigcache) engine and redis ([resp](https://redis.io/topics/protocol)) protocol

<p>
  <a href="#installation">Installation |</a>
  <a href="#usage">Usage</a> |
  <a href="#licenses">License</a>
  <br><br>
  <blockquote>
	cache server with BigCache engine and redis (resp) protocol. You can use any redis client to use this server since it comply with resp protocol. Unfortunately, due to the design of BigCache and some constraint, only some of redis command is implemented. The list of the commands can be found in <a href="#usage">Usage</a>
  </blockquote>
  This project still in progress<br>
  Tested with redis-cli
</p>

Installation
------------
* git clone
* go get -v
* make (coming soon)

Usage
------------
* Only these commands available now
```
    1. GET <key>
    2. SET <key> <value>
    3. DBSIZE //Currently only using one db
    4. DEL <key>
    5. FLUSHALL
    6. PING // of course
```

![screenshot](https://cloud.githubusercontent.com/assets/5761975/23824951/5b5e9c2a-06b3-11e7-8e9c-c68dc4ff45f4.png)


License
----

MIT © [Adhitya Ramadhanus]

