// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
	"os"
	"plathome-backend/controller"
	"plathome-backend/gen/restapi/operations"
	handler2 "plathome-backend/handler"
	"plathome-backend/models"
	"plathome-backend/ping"
	"strings"
)

//go:generate swagger generate server --target ../../gen --name PlatHome --spec ../../swagger.yaml --exclude-main

func configureFlags(api *operations.PlatHomeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PlatHomeAPI) http.Handler {
	if os.Getenv("DBHOST") == "" {
		panic("DBHOST is required")
	}
	var (
		dialect  = "postgres"
		settings = "host=" + os.Getenv("DBHOST") + " user=postgres port=5432 sslmode=disable"
	)
	db := controller.NewDatabase(dialect, settings)
	api.ServeError = errors.ServeError

	go ping.StartPingTask(db)

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DeleteDeviceHandler = operations.DeleteDeviceHandlerFunc(func(params operations.DeleteDeviceParams) middleware.Responder {
		ip := strings.Replace(params.IP, "_", ".", -1)
		db.Delete(ip)
		return operations.NewDeleteDeviceOK().WithPayload(&operations.DeleteDeviceOKBody{Message: "Deleted"})
	})
	api.PutDeviceHandler = operations.PutDeviceHandlerFunc(func(params operations.PutDeviceParams) middleware.Responder {
		fmt.Println("calling...")
		fmt.Println(*params.Device.EzRequesterModels[0].Parameter)
		device := models.NewDevice(params.Device)
		db.Create(&device)
		return operations.NewPutDeviceOK().WithPayload(&operations.PutDeviceOKBody{Message: "Created"})
	})

	api.GetDeviceHandler = operations.GetDeviceHandlerFunc(func(params operations.GetDeviceParams) middleware.
		Responder {
		ms := db.FindAll()
		gms := models.ConvertDevices(*ms)
		return operations.NewGetDeviceOK().WithPayload(gms)
	})

	api.GetPingHandler = operations.GetPingHandlerFunc(func(params operations.GetPingParams) middleware.Responder {
		ip := strings.Replace(params.IP, "_", ".", -1)
		r := models.Device{}
		r.IP = ip
		db.First(&r)
		// record exist check
		if r.State == "" {
			return operations.NewGetPingNotFound().WithPayload(&operations.
				GetPingNotFoundBody{Result: "record not found"})
		}
		res := ping.Ping(ip)
		switch res {
		case "alive":
			return operations.NewGetPingOK().WithPayload(&operations.GetPingOKBody{Result: res})
		case "timeout":
			return operations.NewGetPingOK().WithPayload(&operations.GetPingOKBody{Result: res})
		default:
			return operations.NewGetPingInternalServerError().WithPayload(&operations.
				GetPingInternalServerErrorBody{Result: res})
		}

	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler2.CORShandler(handler)
}
