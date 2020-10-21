package authenticator

import (
	"time"

	"github.com/gin-gonic/gin"
	utils "github.com/lmindwarel/quizzbox-utils"
	"quizzbox.fr/authentificator/datastore"
)

type Config struct {
	ServerPort     string        `json:"serverPort"`
	AccessDuration time.Duration `json:"accessDuration"`
}

// Authenticator is the authenticator object
type Authenticator struct {
	config Config
	ds     *datastore.Datastore
	engine *gin.Engine
}

var log = utils.GetLogger("authenticator")

// New create new authenticator with the given config
func New(config Config, ds *datastore.Datastore) *Authenticator {
	e := gin.Default()
	config.AccessDuration = config.AccessDuration * time.Minute
	return &Authenticator{
		config: config,
		ds:     ds,
		engine: e,
	}
}
