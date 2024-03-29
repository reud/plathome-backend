// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// DeleteDeviceHandlerFunc turns a function with the right signature into a delete device handler
type DeleteDeviceHandlerFunc func(DeleteDeviceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteDeviceHandlerFunc) Handle(params DeleteDeviceParams) middleware.Responder {
	return fn(params)
}

// DeleteDeviceHandler interface for that can handle valid delete device params
type DeleteDeviceHandler interface {
	Handle(DeleteDeviceParams) middleware.Responder
}

// NewDeleteDevice creates a new http.Handler for the delete device operation
func NewDeleteDevice(ctx *middleware.Context, handler DeleteDeviceHandler) *DeleteDevice {
	return &DeleteDevice{Context: ctx, Handler: handler}
}

/*DeleteDevice swagger:route DELETE /device deleteDevice

Delete Devices from DB by IP(converted)

*/
type DeleteDevice struct {
	Context *middleware.Context
	Handler DeleteDeviceHandler
}

func (o *DeleteDevice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteDeviceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeleteDeviceNotFoundBody delete device not found body
// swagger:model DeleteDeviceNotFoundBody
type DeleteDeviceNotFoundBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete device not found body
func (o *DeleteDeviceNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// DeleteDeviceOKBody delete device o k body
// swagger:model DeleteDeviceOKBody
type DeleteDeviceOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete device o k body
func (o *DeleteDeviceOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteDeviceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteDeviceOKBody) UnmarshalBinary(b []byte) error {
	var res DeleteDeviceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
