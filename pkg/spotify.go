package spotify // should be a service

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"tea-time/types"
)

type SpotifyToken struct {
	expires int // seconds
	token   string
}

type tokenResponse struct {
	Token   string `json:"access_token"`
	Expires int    `json:"expires_in"`
}

var instance *SpotifyToken

func GetSpotifyToken() string {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")

	encoded := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))

	if instance == nil {
		data := url.Values{}
		data.Set("grant_type", "client_credentials")

		req, err := http.NewRequest(http.MethodPost, "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
		if err != nil {
			return ""
		}

		req.Header.Set("Authorization", "Basic "+encoded)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return ""
		}
		defer resp.Body.Close()

		var tr tokenResponse
		if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
			return ""
		}

		instance = &SpotifyToken{
			token:   tr.Token,
			expires: tr.Expires,
		}
	}

	return instance.token
}

func GetPlaylist(playlistId string) (*types.Playlist, error) {
	url := "https://api.spotify.com/v1/playlists/" + playlistId

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+GetSpotifyToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get playlist")
	}

	defer resp.Body.Close()

	// Create a file to write the response
	outFile, err := os.Create("response.json")
	if err != nil {
		return nil, err
	}
	defer outFile.Close()

	// Write the response body to file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return nil, err
	}

	// ToDo remove this file logging stuff

	// Seek to the beginning of the file for reading
	_, err = outFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	// Decode the JSON from the file
	var playlist types.Playlist
	if err := json.NewDecoder(outFile).Decode(&playlist); err != nil {
		return nil, err
	}

	return &playlist, nil
}
