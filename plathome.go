package main

import (
	"log"
	"os"

	"PlatHome-Backend/gen/restapi"
	"PlatHome-Backend/gen/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	_ "github.com/lib/pq"
)

func closeServer(sv *restapi.Server) {
	err := sv.Shutdown()
	if err != nil {
		panic(err)
	}
}

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewPlatHomeAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer closeServer(server)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "PlatHome"
	parser.LongDescription = "Home Netework Watcher Service"

	server.ConfigureFlags()
	server.Port = 8080
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
