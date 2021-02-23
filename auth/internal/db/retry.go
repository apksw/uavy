package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	retry struct {
		Client *mongo.Client
		Error  error
	}
)

func (c *Client) retryConnection() (r chan retry) {
	bo := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), c.maxRetries)

	go func() {
		defer close(r)

		url := c.URL()

		for i := 0; i <= int(c.maxRetries); i++ {
			log.Println("dialing to Mongo", "host", url)

			opts := options.Client().ApplyURI(url)

			client, err := mongo.Connect(context.TODO(), opts)
			if err != nil {
				log.Printf("mongo connection error: %s", err.Error())
			}

			err = client.Ping(context.TODO(), nil)

			if err != nil {
				log.Printf("mongo ping error: %s", err.Error())

				// Backoff
				next := bo.NextBackOff()
				if next == backoff.Stop {
					err := errors.New("max number of Mongo connection attempts reached")

					r <- retry{nil, err}

					bo.Reset()
					return
				}

				fmt.Printf("connection attempt to Mongo failed: %s", err.Error())
				fmt.Printf("retrying connection to Mongo in %s seconds", next.String(), "unit", "seconds")

				time.Sleep(next)

			} else {
				r <- retry{client, nil}
				return
			}
		}
	}()

	return r
}
