package download_bai

import (
	"fmt"
	"github.com/vishaltelangre/gaana-dl/scraper"
	"github.com/vishaltelangre/gaana-dl/voodoo"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type Purse struct {
	DestPath    string
	PlaylistURL string
}

func (p *Purse) DownloadTrack(trackID3Meta *scraper.TrackID3Meta) error {
	if trackID3Meta == nil {
		return nil
	}

	trackStreamMeta, err := voodoo.FetchTrackStreamMeta(trackID3Meta.Id)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("==> Downloading: %s\n", trackID3Meta.Title)

	if trackStreamMeta.TrackFormat == "mp4_aac" {
		err := downloadHDSStream(
			p.DestPath,
			trackStreamMeta.StreamPath,
			trackID3Meta.BaseName,
		)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	// NOTE: RTMP stream is allowed to fail and exit here, not supporting
	// downloading it.
	// MPEG audio files gets downloaded, BTW.
	err = downloadOtherStream(
		p.DestPath,
		trackStreamMeta.StreamPath,
		trackID3Meta.BaseName,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func downloadHDSStream(destPath string, trackStreamPath string, trackBaseName string) error {
	hsdUtility := "../vendor/AdobeHDS.php"
	manifestOpt := fmt.Sprintf(
		"--manifest \"%s&g=EBCSRDTPHGVB&hdcore=3.4.0&plugin=aasp-3.4.0.132.66\"",
		trackStreamPath,
	)
	miscOpts := "--quality high --delete"
	hdsCmd := fmt.Sprintf(
		"php %s %s %s --outfile /tmp/%s.flv",
		hsdUtility,
		manifestOpt,
		miscOpts,
		trackBaseName,
	)

	hdsOutfile := "/tmp/" + trackBaseName + ".flv"
	defer os.Remove(hdsOutfile)

	_, err := exec.Command("sh", "-c", hdsCmd).Output()
	if err != nil {
		return err
	}

	// HDS fragments are combined into a single 'hdsOutfile' which has FLV
	// format, so we need to extract audio off of it
	flvToMp3Cmd := fmt.Sprintf(
		"ffmpeg -i \"%s\" -acodec libmp3lame -b:a 192K -vn %s/%s.mp3",
		hdsOutfile,
		destPath,
		trackBaseName,
	)

	_, err = exec.Command("sh", "-c", flvToMp3Cmd).Output()
	if err != nil {
		return err
	}

	return nil
}

func downloadOtherStream(destPath string, trackStreamPath string, trackBaseName string) error {
	trackFile, _ := os.Create(destPath + "/" + trackBaseName + ".mp3")
	defer trackFile.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", trackStreamPath, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	io.Copy(trackFile, res.Body)

	return nil
}
