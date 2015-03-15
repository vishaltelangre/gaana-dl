# gaana-dl

The one and only **free** gaana.com tracks downloader!

![gaana-dl screenshot](https://raw.github.com/vishaltelangre/gaana-dl/master/preview.png)

With automatic ID3 tags injection support:

![gaana-dl screenshot](https://raw.github.com/vishaltelangre/gaana-dl/master/id3_preview.png)

:trollface:

## Prerequsites
- `php`, `ffmpeg` commands
- Download this [AdobeHDS.php](https://github.com/vishaltelangre/gaana-dl/blob/master/vendor-scripts/AdobeHDS.php) script somewhere, which is needed to download Adobe HDS streams.

## Installation

- Before getting started, make sure that you have [installed Go](http://golang.org/doc/install) and have set workspace (`$GOPATH`, etc.), or [RTFM](http://golang.org/doc/code.html) yourself how to do it!

```
go get -u github.com/vishaltelangre/gaana-dl
```

- Test whether `gaana-dl` command works fine by checking version of it:

```
gaana-dl -v
```

- Define `HDS_SCRIPT_PATH` with the path to above downloaded `AdobeHDS.php` script in your `~/.bashrc`, or `~/.zshrc` file. Or you can export it while using `gaana-dl` command, for example:

```
export HDS_SCRIPT_PATH=/path/to/AdobeHDS.php
```

- **NOTE:** Also, look for stand-alone `gaana-dl` executable binary in this repository built using `go build` on Mac OSX.

- **NOTE:** Downloading some songs don't work yet, such as RTMP streams for example.

## Usage

```
Usage:
  gaana-dl [OPTIONS]

The OPTIONS are:
  -u    Playlist URL (Required).
  -a    Absolute path to AdobeHDS.php script (Required if HDS_SCRIPT_PATH environment vairable is not defined).
  -d    Destination directory path where all the tracks will be downloaded.
      By Default, it will download in the current directory only.
  -h    Show this usage help.
  -v    Display version.
```

### Examples:

```
# Here, providing path to AdbobeHDS.php with "-a" option
gaana-dl -u http://gaana.com/album/ek-paheli-leela -d /Users/vishal/Music/leela -a /path/to/AdobeHDS.php

# Or by exporting path while executing command
export HDS_SCRIPT_PATH=path/to/AdobeHDS.php
gaana-dl -u gaana-dl -u http://gaana.com/album/ek-villain -d ./ek-villain

# Or by assuming, HDS_SCRIPT_PATH is already set in your ~/.bashrc:
gaana-dl -u http://gaana.com/playlist/gaana-dj-us-top-50 -d ./dj-us-top-50
 ```

## Changelog

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