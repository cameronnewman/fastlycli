# fastlycli
[![Build Status](https://travis-ci.org/cameronnewman/fastlycli.svg?branch=master)](https://travis-ci.org/cameronnewman/fastlycli)

a simple CLI tool to interact with the Fastly CDN


Usage

```
lappy:~ root$ export FASTLYAPIKEY=adkjsdfiousdksdfoiujsdflkjsdfjk
lappy:~ root$ ./fastlycli service domains -service test.com
[
	{
		"comment": "",
		"name": "css.test.com"
	},
	{
		"comment": "",
		"name": "js.test.com"
	},
	{
		"comment": "",
		"name": "img.test.com"
	}
]
```
```
lappy:~ root$ ./fastlycli purge -service example.com -o http://www.example.com/js/main.js
Service test.com successfully purged
```

```
lappy:~ root$ ./fastlycli purgeall -service example.com
Service test.com successfully purged
```

```
lappy:~ root$ fastlycli -h
NAME:
   fastlycli - Manage Fastly CDN Services via the cli

USAGE:
   fastlycli [global options] command [command options] [arguments...]

VERSION:
   0.8.0

COMMANDS:
   service	Get Service Details
   purge	Purge objects from the CDN
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```
