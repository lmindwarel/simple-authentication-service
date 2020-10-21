package authenticator

import "github.com/gin-gonic/gin"

// StartServer start the server on core config port
func (a *Authenticator) StartServer() error {
	a.setupRoutes()
	a.engine.Use(gin.Logger())
	log.Info("Starting server on port %s", a.config.ServerPort)

	return a.engine.Run(":" + a.config.ServerPort)
}

func (a *Authenticator) setupRoutes() {
	a.engine.POST("/login", a.Register)
	a.engine.POST("/access", a.Login)
	a.engine.GET("/access/:token", a.GetAccess)
	a.engine.DELETE("/access/:token", a.Logout)
}
