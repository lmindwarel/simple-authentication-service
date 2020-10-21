package datastore

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"quizzbox.fr/authentificator/models"
)

type Datastore struct {
	db *mgo.Database
}

// Config is the configuration for the datastore
type Config struct {
	Username string
	Password string
	Name     string
	Hosts    []string
}

// New create a datastore
func New(config Config) (*Datastore, error) {

	loginURL := ""
	if config.Username != "" {
		if config.Password != "" {
			loginURL = fmt.Sprintf("%s:%s@", config.Username, config.Password)
		} else {
			loginURL = fmt.Sprintf("%s@", config.Username)
		}
	}
	url := fmt.Sprintf("mongodb://%s%s/%s", loginURL, strings.Join(config.Hosts, ","), config.Name)

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, errors.New("failed to create mongo session: \n" + err.Error())
	}

	datastore := &Datastore{
		db: session.DB(config.Name),
	}

	return datastore, nil
}

// GetLoginByEmail return the login for given email
func (ds *Datastore) GetLoginByEmail(email string) (models.Login, bool) {
	var login models.Login
	err := ds.db.C(models.CollLogin).Find(bson.M{"email": email}).One(&login)
	return login, err == nil
}

// UpsertLogin add or update the given login in database
func (ds *Datastore) UpsertLogin(login models.Login) error {
	_, err := ds.db.C(models.CollLogin).UpsertId(login.ID, login)
	return err
}

// GetAccessByToken return the login for given email
func (ds *Datastore) GetAccessByToken(token string) (models.Access, bool) {
	var access models.Access
	err := ds.db.C(models.CollAccess).Find(bson.M{"token": token}).One(&access)
	return access, err == nil
}

// UpsertAccess add or update the given access in database
func (ds *Datastore) UpsertAccess(access models.Access) error {
	_, err := ds.db.C(models.CollAccess).UpsertId(access.ID, access)
	return err
}
