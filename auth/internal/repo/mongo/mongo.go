// package mongo provides a Mongo based implementation of UserRepo
// interface
package mongo

import (
	"errors"

	"gitlab.com/adrianpk/uavy/auth/internal/db"
	"gitlab.com/adrianpk/uavy/auth/pkg/base"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repo struct {
		*base.BaseWorker
		conn       *db.Client
		collection string
	}

	Config struct {
		TracingLevel string
	}
)

func NewRepo(name string, conn *db.Client, collection string, cfg Config) *Repo {
	return &Repo{
		BaseWorker: base.NewWorker(name, cfg.TracingLevel),
		conn:       conn,
		collection: collection,
	}
}

func (r *Repo) Conn() (c *db.Client) {
	return r.conn
}

func (r *Repo) Client() (c *mongo.Client, err error) {
	if r.conn.Client == nil {
		return c, errors.New("no MongoDB client")
	}

	return r.conn.Client, nil
}

func (r *Repo) Session() (s mongo.Session, err error) {
	c, err := r.Client()
	if err != nil {
		return s, err
	}

	return c.StartSession()
}

func (r *Repo) Collection() (coll *mongo.Collection, err error) {
	c, err := r.Client()
	if err != nil {
		return coll, err
	}

	return c.Database(r.conn.Db()).Collection(r.collection), nil
}
