package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/lottery/internal/service"
)
var svc *service.Svc

func New(port string, s *service.Svc) *http.Server {
	svc = s
	router := gin.Default()
	initRouter(router)
	srv := &http.Server{Addr: port, Handler: router}
	serveAsync(srv)
	return srv
}

func initRouter(e *gin.Engine) {
	g := e.Group("lottery")
	g.Use(accessJsMiddleware())
	g.GET("/list", prizeList)
	g.GET("/draw", rateLimitMiddleware(), drawPrize)
	g.GET("/export", drawRecords)
	g.POST("/participate", participate)
	g.GET("/users", allUsers)

}

func serveAsync(daemon *http.Server) {
	// it won't block the graceful shutdown handling
	go func() {
		if err := daemon.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

