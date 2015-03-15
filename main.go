/*
 * gaana-dl - The one and only free gaana.com tracks downloader!
 *
 * Author : Vishal Telangre
 * Source : http://github.com/vishaltelangre/gaana-dl
 * License: MIT
 *
 */

package main

import (
	"flag"
	"fmt"
	"github.com/vishaltelangre/gaana-dl/download_bai"
	"github.com/vishaltelangre/gaana-dl/id3"
	"github.com/vishaltelangre/gaana-dl/scraper"
	"log"
	"os"
	"path/filepath"
)

const VERSION = "0.0.3"

var (
	showVersion        *bool   = flag.Bool("v", false, "Show version info")
	showHelp           *bool   = flag.Bool("h", false, "Show help and usage of command")
	customDownloadDir  *string = flag.String("d", ".", "Destination directory path")
	playlistURL        *string = flag.String("u", "", "Playlist URL")
	adobeHDSScriptPath *string = flag.String("a", "", "Absolute path to AdobeHDS.php script")
)

func printUsage() {
	fmt.Println(`
Dude, let's bang gaana.com!

Usage:
	gaana-dl [OPTIONS]

The OPTIONS are:
	-u 		Playlist URL (Required).
	-a 		Absolute path to AdobeHDS.php script (Required if HDS_SCRIPT_PATH environment vairable is not defined).
	-d 		Destination directory path where all the tracks will be downloaded.
			By Default, it will download in the current directory only.
	-h 		Show this usage help.
	-v 		Display version.
`)
}

func ignite(dbp *download_bai.Purse) {
	trackID3MetaArr := scraper.Scrape(dbp.PlaylistURL)

	for _, trackID3Meta := range trackID3MetaArr {
		trackLocation, err := dbp.DownloadTrack(trackID3Meta)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if trackLocation == "" {
			continue
		}

		err = id3.Set(trackLocation, trackID3Meta)
		if err != nil {
			fmt.Printf("Error while setting ID3 tags: %s\n", err)
		}

		fmt.Println("[SUCCESS]")
	}
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("gaana-dl - v%s\n", VERSION)
		return
	}

	if *showHelp || *playlistURL == "" {
		printUsage()
		return
	}

	if *customDownloadDir != "" {
		if _, err := os.Stat(*customDownloadDir); os.IsNotExist(err) {
			log.Fatalln("Download directory \"" +
				*customDownloadDir +
				"\" seems doesn't exists!")
		}
	}

	hdsScriptPath := *adobeHDSScriptPath
	if _, err := os.Stat(hdsScriptPath); os.IsNotExist(err) {
		hdsScriptPath = os.Getenv("HDS_SCRIPT_PATH")

		if _, err = os.Stat(hdsScriptPath); os.IsNotExist(err) {
			log.Fatalf(
				"AdobeHDS.php script at \"%s\" seems doesn't exists!\n",
				hdsScriptPath)
		}
	}

	var dbp download_bai.Purse
	hdsScriptPath, _ = filepath.Abs(hdsScriptPath)
	dbp.AdobeHDSScriptPath = hdsScriptPath
	dbp.DestPath = *customDownloadDir
	dbp.PlaylistURL = *playlistURL

	ignite(&dbp)
}
