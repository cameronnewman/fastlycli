# fastlycli
[![Build Status](https://travis-ci.org/cameronnewman/fastlycli.svg?branch=master)](https://travis-ci.org/cameronnewman/fastlycli)

a simple CLI tool to interact with the Fastly CDN


Usage
```
lappy:~ root$ export FASTLYAPIKEY=adkjsdfiousdksdfoiujsdflkjsdfjk
lappy:~ root$ ./fastlycli purge -service test.com *
Getting test.com service details
Service test.com found. ID=xxxxxxxxxxx
Purging test.com service
Service test.com successfully purged
```

```
lappy:~ root$ ./fastlycli service -service test.com
{
  "id":"SU1Z0isxPaozGVKXdv0eY",
  "name":"test.com",
  "customer_id":"sjshjs9e",
  "versions": null
  ]
}
```

```
lappy:~ root$ fastlycli -h
NAME:
   fastlycli - Manage Fastly CDN Services via the cli

USAGE:
   fastlycli [global options] command [command options] [arguments...]

VERSION:
   0.5.0

COMMANDS:
   purge, p	Purge objects from the CDN
   service, s	Get Service Details
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
```
