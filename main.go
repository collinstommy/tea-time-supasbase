package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"tea-time/components"
	"tea-time/db"
	"tea-time/types"

	auth "tea-time/pkg"
	"tea-time/template"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func getTestUserUUID() pgtype.UUID {
	uuidStr := "30cb0d13-ecf8-4dc0-aae2-e8fe0530cfac"
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	var id pgtype.UUID
	copy(id.Bytes[:], parsedUUID[:])
	id.Valid = true
	return id
}

func DBMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	connectionString := os.Getenv("SUPABASE_CONNECTION_STRING")

	return func(c echo.Context) error {
		ctx := context.Background()
		conn, err := pgx.Connect(ctx, connectionString)
		if err != nil {
			return c.String(http.StatusInternalServerError, "failed to connect to db")
		}
		defer conn.Close(ctx)

		queries := db.New(conn)

		c.Set("db", queries)

		return next(c)
	}
}

func GetArtistString(items []types.Item) string {
	var artistString string
	var count int
	for _, track := range items {
		for _, artist := range track.Track.Artists {
			artistString += artist.Name + ", "
			count++
			if count == 5 {
				return strings.TrimSuffix(artistString, ", ")
			}
		}
	}
	return strings.TrimSuffix(artistString, ", ")
}

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	auth.GetSpotifyToken()

	e := echo.New()
	template.NewTemplateRenderer(e)

	e.Use(DBMiddleware)
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		db := c.Get("db").(*db.Queries)
		playlists, err := db.ListPlaylists(context.Background())

		if err != nil {
			return err
		}

		component := components.Index(playlists)

		return template.Html(c, component)
	})

	e.POST("/playlists", func(c echo.Context) error {
		data := c.Get("db").(*db.Queries)
		spotifyID := c.FormValue("spotifyId")

		// ToDo: move spotify out of auth
		playlist, err := auth.GetPlaylist(spotifyID)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(http.StatusOK)
		}

		params := db.CreatePlaylistParams{
			SpotifyID:  spotifyID,
			UserID:     getTestUserUUID(),
			Artists:    GetArtistString(playlist.Tracks.Items),
			Name:       playlist.Name,
			TrackCount: playlist.Tracks.Total,
		}

		newPlaylist, err := data.CreatePlaylist(context.Background(), params)
		if err != nil {
			fmt.Println(err)
			return c.NoContent(http.StatusOK)
		}
		return template.Html(c, components.SinglePlaylist(newPlaylist))
	})
	e.Logger.Fatal(e.Start(":1323"))
}
