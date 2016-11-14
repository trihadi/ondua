<img src="https://s14.postimg.org/97kllozpd/Demo.png" alt="screenshot"/>

**kosim Torrent** is an a self-hosted remote torrent client, written in Go (golang). You start torrents remotely, which are downloaded as sets of files on the local disk of the server, which are then retrievable or streamable via HTTP.

### Features

* Single binary
* Cross platform
* Embedded torrent search
* Real-time updates
* Mobile-friendly
* Fast [content server](http://golang.org/pkg/net/http/#ServeContent)
* [*More features coming soon*](https://github.com/kosim21/kosim-torrent/labels/core-feature)

### Install

**Binaries**

See [the latest release](https://github.com/kosim21/kosim-torrent/releases/latest)

**Source**

*[Go](https://golang.org/dl/) is required to install from source*

``` sh
$ go get -v github.com/kosim21/kosim-torrent
```

**Docker**

``` sh
$ docker run -d -p 3000:3000 -v /path/to/my/downloads:/downloads kosim21/kosim-torrent
```

**Heroku**

Click this button to...

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Usage

```
$ kosim-torrent --help

  Usage: iTS-torrent [options]

  Options:
  --title, -t        Title of this instance (default kosim Torrent, env TITLE)
  --port, -p         Listening port (default 3000, env PORT)
  --host, -h         Listening interface (default all)
  --auth, -a         Optional basic auth in form 'user:password' (env AUTH)
  --config-path, -c  Configuration file path (default kosim-torrent.json)
  --key-path, -k     TLS Key file path
  --cert-path, -r    TLS Certicate file path
  --log, -l          Enable request logging
  --open, -o         Open now with your default browser
  --help
  --version, -v

  Version:
    0.8.8

  Read more:
    https://github.com/kosim21/kosim-torrent

```

#### License

Copyright (c) 2016 kosim21

[Creative Commons Legal Code - Attribution-NonCommercial 3.0 Unported](LICENSE)
