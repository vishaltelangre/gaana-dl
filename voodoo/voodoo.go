package voodoo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type TrackStreamMeta struct {
	Status       string `json:"status,Number"`
	StreamPath   string `json:"stream_path"`
	StreamServer string `json:"stream_server"`
	RequestType  string `json:"request_type"`
	TrackFormat  string `json:"track_format"`
	Protocol     string `json:"protocol"`
}

func GetTrackStreamMetaPath(trackId string) string {
	url := "https://gnr-w1.gaana.com/gs2.php"
	return fmt.Sprintf("%s?quality=high&type=web&track_id=%s", url, trackId)
}

func FetchTrackStreamMeta(trackId string) (*TrackStreamMeta, error) {
	trackStreamMetaPath := GetTrackStreamMetaPath(trackId)
	req, err := http.NewRequest("GET", trackStreamMetaPath, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var trackStreamMeta TrackStreamMeta
	err = json.Unmarshal(content, &trackStreamMeta)
	if err != nil {
		return nil, err
	}

	return &trackStreamMeta, nil
}

func GetTrackStreamPath(trackId string) (string, error) {
	trackStreamMeta, err := FetchTrackStreamMeta(trackId)
	if err != nil {
		return "", err
	}

	return trackStreamMeta.StreamPath, err
}
