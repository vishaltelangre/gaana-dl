package id3

import (
	taggr "github.com/mikkyang/id3-go"
	"github.com/vishaltelangre/gaana-dl/scraper"
)

func Set(trackPath string, meta *scraper.TrackID3Meta) error {
	mp3File, err := taggr.Open(trackPath)

	defer mp3File.Close()

	if err != nil {
		return err
	}

	mp3File.SetTitle(meta.Title)
	mp3File.SetArtist(meta.Artist)
	mp3File.SetAlbum(meta.AlbumTitle)

	return nil
}
