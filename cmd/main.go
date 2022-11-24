package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"rss_reader_go_v2/pkg/rss_reader"
)

var feedURLs = []string{
	"https://blog.golang.org/feed.atom",
	"http://feeds.kottke.org/main",
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/feeds", GetRssData)

	// Start server
	e.Logger.Fatal(e.Start(":8070"))
	return
}

// Handler
func hello(c echo.Context) error {
	rss_reader.FetchFeeds(feedURLs)
	return nil
}
