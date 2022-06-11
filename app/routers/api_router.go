package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jun00rbiter/go-blog/app/controllers"
)

// RegisterAPIRouter は、APIのルーティングを定義する
func RegisterAPIRouter(engine *gin.Engine) {
	// ルーティングしていないURLにアクセスが来た時にエラーJSONで返す
	engine.NoRoute(func(c *gin.Context) {
		c.AsciiJSON(http.StatusNotFound, gin.H{"error": gin.H{"message": "not found"}})
	})

	apiGroup := engine.Group("/api")
	apiGroup.GET("/health-check", controllers.GetHealthCheck) // ヘルスチェック
}
