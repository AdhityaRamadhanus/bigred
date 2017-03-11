# BIGRED

Cache server with BigCache engine and redis (resp) protocol

<p>
  <a href="#Installation">Installation |</a>
  <a href="#Usage">Usage</a> |
  <a href="#licenses">License</a>
  <br><br>
  <blockquote>
	cache server with BigCache engine and redis (resp) protocol. You can use any redis client to use this server since it comply with resp protocol. Unfortunately, due to the design of BigCache and some constraint, only some of redis command is implemented. The list of the commands can be found in <a href="#Usage">Usage</a>

    This project still in progress

    Tested with redis-cli
  </blockquote>
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
    5. PING // of course
```

License
----

MIT © [Adhitya Ramadhanus]

