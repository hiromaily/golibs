# golibs

[![Build Status](https://travis-ci.org/hiromaily/golibs.svg?branch=master)](https://travis-ci.org/hiromaily/golibs)
[![Coverage Status](https://coveralls.io/repos/github/hiromaily/golibs/badge.svg?branch=master)](https://coveralls.io/github/hiromaily/golibs?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hiromaily/golibs)](https://goreportcard.com/report/github.com/hiromaily/golibs)
[![codebeat badge](https://codebeat.co/badges/233e0a28-9066-450e-acd6-8fdeb58d993a)](https://codebeat.co/projects/github-com-hiromaily-golibs-master)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](https://raw.githubusercontent.com/hiromaily/golibs/master/LICENSE)

Golang libraries 

## requirement
golang version is limited some packages.  
* [log] only version 1.5+

## libraries

### ■ algorithm
* For indivisual study. (It works in progress.)

----
### ■ auth
#### jwt
* authentication by JWT (Json Web TOken) 

----
### ■ cast 
* for now, cast byte to string mutually.

----
### ■ cipher
#### encryption
* For encryption.

#### hash
* For hash.

----
### ■ compress
* For compress string data.

----
### ■ config
* For config settings from toml file.

----
### ■ db
#### boltdb
* For check how to use boltdb package.

#### cache
* Cache result of query on MySQL to Redis. (It works in progress.)

#### cassandra
* Handling Cassandra.

#### gorm
* Handling RDB usign gorm package. For check.

#### gorp
* Handling RDB usign gorp package. For check.

#### mongodb
* Handling MongoDB.

#### mysql
* Handling MySQL.

#### postgresql
* Handling PostgreSQL. (It's not started to develop yet.)

#### redis
* Handling Redis.

#### textindex
* For check how to use textindex package. (It's not started to develop yet.)

----
### ■ files
* For finding specific file in directories.

----
### ■ goroutine
* It's to control how many goroutine can be run at a time.

----
### ■ heroku
* Heroku related library.

----
### ■ log
* logger. It can control loglevel in files and stdout respectively
* only version 1.5 or later.

----
### ■ mail
* Send mail. (It works in progress.)

----
### ■ messaging
#### kafka
* Handling Apache kafka

#### nats
* Handling NATS

#### rabbitmq
* Handling RabbitMQ.

----
### ■ os
* For util feature using "os" package.

----
### ■ reflects
* For check how to use "reflect" package.

----
### ■ regexp
* For search string to use "regexp" package.

----
### ■ runtimes
* It's just check how to use "runtime" package.

----
### ■ serial
* Serialize each Type to binary data.

----
### ■ signal
* Control signal to check where goroutine stop.

----
### ■ testutil
* Utility for test packages.

----
### ■ time
* For time related library.

----
### ■ tmpl
* It's just check how to use "template/text" package.

----
### ■ utils
* For useful funcs like handling slice, interface, type and so on.

----
### ■ validator
* To check variable.

----
### ■ web
#### ■ context
* For context. (It's not started to develop yet.)

----
#### ■ session
* Control session on the web. (It works in progress.)

----

## libraries (example)  
### ■ example
* These packages are for just check

#### ■ defaultdata
* For check how nil is treated in each Type through Test.
* And how each type is treated on Interface{} type. 

----
#### ■ draw
* For drawing library. (It's not started to develop yet.)

----
#### ■ exec
* For check how to use "os/exec" package.

----
#### ■ flag
* For check how to use "flag" package.

----
#### ■ http
* It's just check how to use "net/http" package. (It works in progress.)

----
#### ■ join
* For check which way is most efficient to join strings on testing package.

----
#### ■ json
* For check how to use "json" package.

----
#### ■ xml
* For handling xml.

