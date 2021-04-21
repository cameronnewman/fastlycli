# fastlycli

[![Build][1]][2]
[![GoDoc][3]][4]
[![Go Report Card][5]][6]

[1]: https://github.com/cameronnewman/fastlycli/workflows/Continuous%20Integration/badge.svg
[2]: https://github.com/cameronnewman/fastlycli/actions
[3]: https://godoc.org/github.com/cameronnewman/fastlycli?status.svg
[4]: http://godoc.org/github.com/cameronnewman/fastlycli
[5]: https://goreportcard.com/badge/github.com/cameronnewman/fastlycli
[6]: https://goreportcard.com/report/github.com/cameronnewman/fastlycli

## Purpose

a simple CLI tool to interact with the Fastly CDN

## Usage

```bash
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

```bash
lappy:~ root$ ./fastly purge -service example.com -o http://www.example.com/js/main.js
Service test.com successfully purged
```

```bash
lappy:~ root$ ./fastly purgeall -service example.com
Service test.com successfully purged
```

```bash
lappy:~ root$ fastly -h
NAME:
   fastly - Manage Fastly CDN Services via the cli

USAGE:
   fastly [global options] command [command options] [arguments...]

VERSION:
   0.9.0

COMMANDS:
   service   Get Service Details
   purge     Purge objects from the CDN
   purgeall  Purge all objects from the CDN
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h        show help
   --version, -v     print the version
```

## Issues

* None

## License

MIT Licensed
