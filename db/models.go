// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Game struct {
	GameID    pgtype.UUID
	Url       string
	Name      string
	CreatedBy pgtype.UUID
}

type Like struct {
	LikeID      pgtype.UUID
	UserID      pgtype.UUID
	ContentType string
	ContentID   pgtype.UUID
}

type Playlist struct {
	PlaylistID pgtype.UUID
	SpotifyID  string
	UserID     pgtype.UUID
	Artists    string
	Name       string
	TrackCount int
}

type Post struct {
	PostID pgtype.UUID
	Url    string
	UserID pgtype.UUID
	Type   string
}

type Session struct {
	SessionID pgtype.UUID
	GameID    pgtype.UUID
	StartTime pgtype.Timestamp
}

type SessionParticipant struct {
	SessionID pgtype.UUID
	UserID    pgtype.UUID
	Result    string
}
