package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/lucaseg/go-bookstore-oauth-api/src/domain/access_token"
	"github.com/lucaseg/go-bookstore-oauth-api/src/utils/errors"
)

type client struct {
}

func NewClient() *client {
	return &client{}
}

func (c *client) GetById(id string) (*access_token.AccessToken, *errors.RestError) {
	return nil, nil
}

var (
	cluster *gocql.ClusterConfig
)

func init() {
	//connnect to the cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	return session, err
}
