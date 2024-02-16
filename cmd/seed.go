package main

import (
	"context"
	"fmt"
	"os"
	"tea-time/db"
	auth "tea-time/pkg"
	spotify "tea-time/pkg"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	connectionString := os.Getenv("SUPABASE_CONNECTION_STRING")

	auth.GetSpotifyToken()

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connectionString)

	if err != nil {
		os.Exit(1)
	}

	defer conn.Close(ctx)

	queries := db.New(conn)
	playlists, err := queries.ListPlaylists(context.Background())

	if err != nil {
		fmt.Println("Failed to list playlists")
		os.Exit(1)
	}

	for _, playlist := range playlists {
		p, err := spotify.GetPlaylist(playlist.SpotifyID)
	}
}
