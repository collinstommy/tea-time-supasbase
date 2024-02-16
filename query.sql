

-- name: ListPlaylists :many
SELECT * FROM playlists;

-- name: CreatePlaylist :one
INSERT INTO playlists (spotify_id, user_id, artists, name, track_count)
VALUES ($1, $2, $3, $4, $5) RETURNING *;
