# gaana-dl

The one and only **free** gaana.com tracks downloader!

![gaana-dl screenshot](https://raw.github.com/vishaltelangre/gaana-dl/master/preview.png)

With automatic ID3 tags injection support:

![gaana-dl screenshot](https://raw.github.com/vishaltelangre/gaana-dl/master/id3_preview.png)

:trollface:

## Prerequsites
- `php`, `ffmpeg` commands

## Download standalone `gaana-dl` binaries:

Head onto Releases page [here](https://github.com/vishaltelangre/gaana-dl/releases).

## Hacker's way of installation

- Before getting started, make sure that you have [installed Go](http://golang.org/doc/install) and have set workspace (`$GOPATH`, etc.), or [RTFM](http://golang.org/doc/code.html) yourself how to do it!

```
go get -u github.com/vishaltelangre/gaana-dl
```

- Test whether `gaana-dl` command works fine by checking version of it:

```
gaana-dl -v
```

- **NOTE:** Downloading some songs don't work yet, such as RTMP streams for example.

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

### Examples:

```
gaana-dl -u http://gaana.com/album/ek-paheli-leela -d /Users/vishal/Music/leela
gaana-dl -u http://gaana.com/album/ek-villain -d ./ek-villain
 ```

## Changelog

### __v0.0.4__
- No need to provide in AdobeHDS script path now onwards (thanks [@sharang-d](http://github.com/sharang-d) for a good suggestion)

### __v0.0.3__
- [FIX] Generalize CSS selector to parse a provided listing url, for e.g. this one -- http://gaana.com/artist/hans-zimmer/songs :P

### __v0.0.2__
- Automatic support for setting ID3 tags (meta details) of tracks while downloading.
- [FIX] Abode HDS streams aren't downloading.

### __v0.0.1__
- Support to download most of the songs from gaana.com!


## Thankings

I am making use of [this](https://github.com/K-S-V/Scripts/blob/master/AdobeHDS.php) script to download fragmented Adobe HDS streams, thanks for folks who contributed to this script.

## Copyright and License

Copyright (c) 2015, Vishal Telangre. All Rights Reserved.

This project is licenced under the [MIT License](LICENSE.md).