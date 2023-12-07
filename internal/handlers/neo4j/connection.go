package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SetupNeo4jDB() (neo4j.DriverWithContext, context.Context) {
	ctx := context.Background()

	dbUri := "neo4j://localhost"
	dbUsername := ""
	dbPassword := ""
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUsername, dbPassword, ""))
	// defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	return driver, ctx
}
