package authenticator

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lmindwarel/authentificator/models"
	"github.com/lmindwarel/authentificator/utils"
	"github.com/pkg/errors"
)

// Login is used to log account in
func (a *Authenticator) Login(c *gin.Context) {
	var err error

	var creds models.Login
	err = c.ShouldBindJSON(&creds)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, found := a.ds.GetLoginByEmail(creds.Email)
	if !found {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if !utils.HashCorrespond(login.Password, creds.Password) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	access, err := a.GrantAccess(login)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, access)
}

// Register was used to create new account
func (a *Authenticator) Register(c *gin.Context) {
	var err error

	var creds models.Login
	err = c.ShouldBindJSON(&creds)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, found := a.ds.GetLoginByEmail(creds.Email)
	if found {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	hash, err := utils.HashFromString(creds.Password)

	login := models.Login{
		Email:    creds.Email,
		Password: hash,
	}

	err = a.ds.UpsertLogin(login)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

// Logout is used to log account out
func (a *Authenticator) Logout(c *gin.Context) {
	var err error

	token := c.Param("token")

	access, found := a.ds.GetAccessByToken(token)
	if !found {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	access.EndDate = time.Now()

	err = a.ds.UpsertAccess(access)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (a *Authenticator) GetAccess(c *gin.Context) {

	token := c.Param("token")

	access, found := a.ds.GetAccessByToken(token)
	if !found {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// check access validity
	// if login.EndDate

	// Refresh token

	c.JSON(http.StatusOK, access)
}

// GrantAccess generate a token and an create an access for login
func (a *Authenticator) GrantAccess(login models.Login) (models.Access, error) {
	var access models.Access
	var token string

	token, err := utils.HashFromString(login.Email)
	if err != nil {
		return access, errors.Wrap(err, "failed to generate access token")
	}

	now := time.Now()

	access = models.Access{
		ID:        utils.NewUUID(),
		Token:     token,
		StartDate: now,
		EndDate:   now.Add(a.config.AccessDuration),
	}

	err = a.ds.UpsertAccess(access)
	if err != nil {
		return access, errors.Wrap(err, "failed to upsert access")
	}

	return access, nil
}
