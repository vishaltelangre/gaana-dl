package main

import (
	"flag"
	"fmt"
	"github.com/vishaltelangre/gaana-dl/download_bai"
	"github.com/vishaltelangre/gaana-dl/scraper"
	"log"
	"os"
)

const VERSION = "0.0.1"

var (
	showVersion       *bool   = flag.Bool("v", false, "Show version info")
	showHelp          *bool   = flag.Bool("h", false, "Show help and usage of command")
	customDownloadDir *string = flag.String("d", ".", "Destination directory path")
	playlistURL       *string = flag.String("u", "", "Playlist URL")
)

func printUsage() {
	fmt.Println(`
Dude, let's bang gaana.com!

Usage:
	gaana-dl [OPTIONS]

The OPTIONS are:
	-u 		Playlist URL (Required).
	-d 		Destination directory path where all the tracks will be downloaded.
			By Default, it will download in the current directory only.
	-h 		Show this usage help.
	-v 		Display version.
`)
}

func ignite(dbp *download_bai.Purse) {
	trackID3MetaArr := scraper.Scrape(dbp.PlaylistURL)
	for _, trackID3Meta := range trackID3MetaArr {
		dbp.DownloadTrack(trackID3Meta)
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("gaana-dl - v%s\n", VERSION)
		return
	}

	if *customDownloadDir != "" {
		if _, err := os.Stat(*customDownloadDir); os.IsNotExist(err) {
			log.Fatalln("Download directory \"" +
				*customDownloadDir +
				"\" seems doesn't exists!")
			return
		}
	}

	if *showHelp || *playlistURL == "" {
		printUsage()
		return
	}

	var dbp download_bai.Purse
	dbp.DestPath = *customDownloadDir
	dbp.PlaylistURL = *playlistURL

	ignite(&dbp)
}
