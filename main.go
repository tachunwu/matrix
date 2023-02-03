package main

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	for {
		// Here is matrix!
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		CellFn(ctx)
		time.Sleep(10 * time.Second)
	}

}

func CellFn(ctx context.Context) {
	/////////////////////////////////////////////
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go srv.ListenAndServe()
	/////////////////////////////////////////////
	<-ctx.Done()
	srv.Shutdown(ctx)

}
