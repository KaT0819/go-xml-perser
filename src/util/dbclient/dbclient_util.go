package dbclient

import (
	ctx "context"
	"errors"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/spanner"
)

var client *spanner.Client

// InitClient google spannerのクライアントを生成する
func initialize(projectID, instance, db string) (c *spanner.Client, err error) {

	dbPath := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectID, instance, db)

	client, err := spanner.NewClientWithConfig(ctx.Background(), dbPath, spanner.ClientConfig{
		SessionPoolConfig: spanner.SessionPoolConfig{
			MinOpened: 5,
			MaxOpened: 50,
			MaxIdle:   10,
			MaxBurst:  50,
		},
	})
	if err != nil {
		log.Println(err)
		return nil, errors.New("failed to Create Spanner Client")
	}

	return client, nil
}

func InitClient() error {
	projectId := os.Getenv("SPANNER_PROJECT_ID")
	instance := os.Getenv("SPANNER_INSTANCE_ID")
	db := os.Getenv("SPANNER_DATABASE_ID")

	var err error
	client, err = initialize(projectId, instance, db)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetClient() *spanner.Client {
	return client
}
