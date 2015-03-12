package scraper

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type TrackID3Meta struct {
	Id           string `json:"id,Number"`
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	AlbumTitle   string `json:"albumtitle"`
	AlbumArtwork string `json:"albumartwork"`
	BaseName     string
}

const (
	TRACK_NODE_CSS_PATH = `.content-container .track > .parentnode,
                         .content-container .album-horizantal-listing
                         		> .parentnode`
)

func parseDOM(playlistURL string) *goquery.Selection {
	doc, err := goquery.NewDocument(playlistURL)
	if err != nil {
		log.Fatal(err)
	}

	trackNodes := doc.Find(TRACK_NODE_CSS_PATH)
	return trackNodes
}

func unmarshalTrackMeta(rawTrackMeta string) (*TrackID3Meta, error) {
	var trackID3Meta TrackID3Meta
	err := json.Unmarshal([]byte(rawTrackMeta), &trackID3Meta)
	if err != nil {
		return nil, err
	}

	// Cleanup dirty Artist stuff!
	delimitedArtists := strings.Split(trackID3Meta.Artist, ",")
	artistsArr := make([]string, len(delimitedArtists))
	for _, artist := range delimitedArtists {
		artist := strings.Split(artist, "###")[0]
		artistsArr = append(artistsArr, artist)
	}
	trackID3Meta.Artist = strings.Join(artistsArr, "")

	// Cleanup dirty Title, also
	trackID3Meta.Title = strings.Replace(trackID3Meta.Title, "&nbsp;", "", -1)
	trackID3Meta.Title = strings.Replace(trackID3Meta.Title, "&#039;", "'", -1)

	// Set track's basename
	trackID3Meta.BaseName = strings.Replace(trackID3Meta.Title, " ", "_", -1)

	return &trackID3Meta, nil
}

func Scrape(playlistURL string) []*TrackID3Meta {
	trackNodes := parseDOM(playlistURL)

	// Make a slice for tracks having length of found trackNodes
	tracks := make([]*TrackID3Meta, trackNodes.Length())

	trackNodes.Each(func(i int, s *goquery.Selection) {
		trackID3Meta, err := unmarshalTrackMeta(s.Text())
		if err == nil {
			tracks = append(tracks, trackID3Meta)
		}
	})

	return tracks
}
