/*
fastly is a simple CLI tool to interact with the Fastly CDN

Usage

```
lappy:~ root$ export FASTLYAPIKEY=adkjsdfiousdksdfoiujsdflkjsdfjk
lappy:~ root$ ./fastly service domains -service test.com
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
lappy:~ root$ ./fastly purge -service test.com -o http://test.com/js/main.js
Service test.com successfully purged
```

```
lappy:~ root$ ./fastly purgeall -service test.com
Service test.com successfully purged
```

```
lappy:~ root$ fastly -h
NAME:
   fastly - Manage Fastly CDN Services via the cli

USAGE:
   fastly [global options] command [command options] [arguments...]

VERSION:
   0.9.0

COMMANDS:
   service	Get Service Details
   purge	Purge objects from the CDN
   purgeall	Purge all objects from the CDN
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```
*/

package main
