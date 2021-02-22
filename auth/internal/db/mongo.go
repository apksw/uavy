// Package mongo is a thin wrapper over official Mongo client including support
// for exponential backoff connection retries and base.Service interface
// implementation.
// TODO: Replace logging by tracing data.
package mongo

import (
	"fmt"

	"gitlab.com/adrianpk/uavy/auth/pkg/base"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Client struct {
		*base.BaseService
		*mongo.Client
		maxRetries uint64
	}
)

func NewClient(name string, maxRetries uint64) (*Client, error) {
	return &Client{
		BaseService: base.NewService("mongo-client"),
		maxRetries:  maxRetries,
	}, nil
}

func (c *Client) Init() (ok chan bool) {
	ok = make(chan bool)

	go func() {
		defer close(ok)

		err := c.connect()
		if err != nil {
			ok <- false
			return
		}

		fmt.Printf("%s service initializated", c.Name())

		ok <- true
	}()

	return ok
}

func (c *Client) Start() error {
	panic("not implemented")
}

func (c *Client) connect() error {
	r := <-c.retryConnection()

	if r.Error != nil {
		return r.Error
	}

	c.Client = r.Client

	return nil
}

func (c *Client) URL() string {
	// FIX: Get value from configuration
	return "mongodb://localhost:27016"
}
