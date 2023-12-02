package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func SetupNeo4jDB() (neo4j.DriverWithContext, context.Context) {
	ctx := context.Background()

	dbUri := "neo4j://localhost"
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth("", "", ""))
	// defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	// result, _ := neo4j.ExecuteQuery(ctx, driver,
	// 	"MATCH (f:Film)-[:HAS_GENDER]->(g:Gender) RETURN f,g LIMIT 25;",
	// 	map[string]any{
	// 		"id": "1",
	// 	}, neo4j.EagerResultTransformer,
	// 	neo4j.ExecuteQueryWithDatabase("neo4j"))

	// for _, record := range result.Records {
	// 	fmt.Println(record.AsMap())
	// }

	// fmt.Printf("The query `%v` returned %v records in %+v.\n",
	// 	result.Summary.Query().Text(), len(result.Records),
	// 	result.Summary.ResultAvailableAfter())

	return driver, ctx
}
