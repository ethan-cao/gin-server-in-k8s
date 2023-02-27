package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const (
	uri      = "neo4j://localhost:7687"
	username = "neo4j"
	password = "password"
)

func main() {
	r := gin.New()

	r.SetTrustedProxies([]string{"192.168.1.*"})

	message, _ := runWithNeo4j()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World, time is "+fmt.Sprint(time.Now().Unix()))
	})

	r.GET("/neo4j", func(c *gin.Context) {
		c.String(200, message)
	})

	r.Run(":3000")
}

func runWithNeo4j() (string, error) {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return "", err
	}
	defer driver.Close(ctx)

	session := driver.NewSession(ctx, neo4j.SessionConfig{})
	defer session.Close(ctx)

	greeting, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (any, error) {
		result, err := tx.Run(
			ctx,
			"CREATE (a:Greeting) SET a.message = $message RETURN a.message + ', from node ' + id(a)",
			map[string]any{"message": "hello, world"},
		)

		if err != nil {
			return nil, err
		}

		if result.Next(ctx) {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})

	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}
