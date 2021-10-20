package main

import (
	"os"
	"time"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"

	"go-xml-parser/src/gen/restapi"
	"go-xml-parser/src/gen/restapi/operations"
	"go-xml-parser/src/util/dbclient"

	_ "github.com/lib/pq" // postgres
	"go.uber.org/zap"
)

const location = "Asia/Tokyo"

// set JST timezone
func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

func main() {
	logger, _ := zap.NewDevelopment()

	logger.Info("swagger定義の組込み")
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Info("routing設定")
	api := operations.NewRoutingAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer func() {
		err := server.Shutdown()
		if err != nil {
			logger.Fatal(err.Error())
		}
	}()

	logger.Info("spannerクライアント初期化")
	err = dbclient.InitClient()
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer dbclient.GetClient().Close()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "xml-parser-api"
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	if _, err = parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()
	logger.Info("app will serve.")
	if err := server.Serve(); err != nil {
		logger.Fatal(err.Error())
	}
}
