package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mailgun/groupcache/v2"
)

const (
	PORT     = "8084"
	BASE_URL = "http://localhost:" + PORT
)

func routes(e *echo.Echo, cache *groupcache.Group, pool *groupcache.HTTPPool) {
	e.GET("/", func(c echo.Context) error {
		key := c.QueryParam("key")
		if key == "" {
			return c.String(http.StatusOK, "Server 4 :: Query param not provided")
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()

		var data []byte
		err := cache.Get(ctx, key, groupcache.AllocatingByteSliceSink(&data))
		if err != nil {
			response := fmt.Sprintf("Server 4 :: Error getting data from cache, %v", err)
			return c.String(http.StatusInternalServerError, response)
		}

		response := fmt.Sprintf("Server 4 :: Response %s", data)
		return c.String(http.StatusOK, response)
	})

	// group cache use following path to communicate with other peers
	// so add '/_groupcache/' path in our server
	e.GET("/_groupcache/*path", func(c echo.Context) error {
		pool.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}

func main() {
	cache, pool := InitializeCache(BASE_URL)

	e := echo.New()
	routes(e, cache, pool)

	if err := e.Start(":" + PORT); err != nil {
		log.Fatalf("Error starting Server 4: %v", err)
	}
}
