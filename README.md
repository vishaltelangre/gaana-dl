# gaana-dl

One and only **free** gaana.com tracks downloader!

![gaana-dl screenshot](https://raw.github.com/vishaltelangre/gaana-dl/master/preview.png)

## Prerequsites
- `php`, `ffmpeg` commands

## Installation

- Before getting started, make sure that you have [installed Go](http://golang.org/doc/install) and have set workspace (`$GOPATH`, etc.), or [RTFM](http://golang.org/doc/code.html) yourself how to do it!

```
go get -u github.com/vishaltelangre/gaana-dl
```

- Test whether `gaana-dl` command works fine by checking version of it:

```
gaana-dl -v
```

**NOTE:** Also, look for `gaana-dl` executable binary in this repository built using `go build` on Mac OSX.

**NOTE:** Downloading some songs don't work yet, such as RTMP streams for example.

## Usage

```
Usage:
  gaana-dl [OPTIONS]

The OPTIONS are:
  -u    Playlist URL (Required).
  -d    Destination directory path where all the tracks will be downloaded.
      By Default, it will download in the current directory only.
  -h    Show this usage help.
  -v    Display version.
```

Examples:

```
gaana-dl -u http://gaana.com/album/ek-paheli-leela -d /Users/vishal/Music/leela
gaana-dl -u http://gaana.com/playlist/gaana-dj-us-top-50 -d ./dj-us-top-50
```

## TODO

- ID3 tag (track meta details) injection support

## Thankings

I have used [this](https://github.com/K-S-V/Scripts/blob/master/AdobeHDS.php) script to download fragmented Adobe HDS streams, thanks for folks who contributed to this script.

## Copyright and License

Copyright (c) 2015, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).