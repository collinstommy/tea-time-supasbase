package types

type Playlist struct {
	Collaborative bool   `json:"collaborative"`
	Description   string `json:"description"`
	ExternalUrls  struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name         string      `json:"name"`
	Owner        Owner       `json:"owner"`
	PrimaryColor interface{} `json:"primary_color"`
	Public       bool        `json:"public"`
	SnapshotID   string      `json:"snapshot_id"`
	Tracks       Tracks      `json:"tracks"`
	Type         string      `json:"type"`
	Uri          string      `json:"uri"`
}

type Owner struct {
	DisplayName  string `json:"display_name"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href string `json:"href"`
	ID   string `json:"id"`
	Type string `json:"type"`
	Uri  string `json:"uri"`
}

type Tracks struct {
	Href     string      `json:"href"`
	Items    []Item      `json:"items"`
	Limit    int         `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int         `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int         `json:"total"`
}

type Item struct {
	AddedAt string `json:"added_at"`
	AddedBy Owner  `json:"added_by"`
	IsLocal bool   `json:"is_local"`
	Track   Track  `json:"track"`
}

type Track struct {
	Album   Album `json:"album"`
	Artists []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	} `json:"artists"`
	DiscNumber  int  `json:"disc_number"`
	DurationMs  int  `json:"duration_ms"`
	Episode     bool `json:"episode"`
	Explicit    bool `json:"explicit"`
	ExternalIds struct {
		Isrc string `json:"isrc"`
	} `json:"external_ids"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href        string      `json:"href"`
	ID          string      `json:"id"`
	IsLocal     bool        `json:"is_local"`
	Name        string      `json:"name"`
	Popularity  int         `json:"popularity"`
	PreviewUrl  interface{} `json:"preview_url"`
	Track       bool        `json:"track"`
	TrackNumber int         `json:"track_number"`
	Type        string      `json:"type"`
	Uri         string      `json:"uri"`
}

type Album struct {
	AlbumType string `json:"album_type"`
	Artists   []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		ID   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
		Uri  string `json:"uri"`
	} `json:"artists"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	ID     string `json:"id"`
	Images []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int    `json:"total_tracks"`
	Type                 string `json:"type"`
	Uri                  string `json:"uri"`
}
