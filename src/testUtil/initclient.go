package testUtil

import (
	"log"
	"os"

	"go-xml-parser/src/util/dbclient"

	"cloud.google.com/go/spanner"
)

func DbClient() *spanner.Client {
	//spanner接続用
	os.Setenv("SPANNER_EMULATOR_HOST", "localhost:9010")
	err := dbclient.InitClient()
	if err != nil {
		log.Println(err)
	}
	return dbclient.GetClient()

}

func ConvertInt64(v int) *int64 {
	t := int64(v)
	return &t
}

func String(v string) *string {
	return &v
}
