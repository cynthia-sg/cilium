// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2017-2021 Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteFqdnCacheHandlerFunc turns a function with the right signature into a delete fqdn cache handler
type DeleteFqdnCacheHandlerFunc func(DeleteFqdnCacheParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteFqdnCacheHandlerFunc) Handle(params DeleteFqdnCacheParams) middleware.Responder {
	return fn(params)
}

// DeleteFqdnCacheHandler interface for that can handle valid delete fqdn cache params
type DeleteFqdnCacheHandler interface {
	Handle(DeleteFqdnCacheParams) middleware.Responder
}

// NewDeleteFqdnCache creates a new http.Handler for the delete fqdn cache operation
func NewDeleteFqdnCache(ctx *middleware.Context, handler DeleteFqdnCacheHandler) *DeleteFqdnCache {
	return &DeleteFqdnCache{Context: ctx, Handler: handler}
}

/*DeleteFqdnCache swagger:route DELETE /fqdn/cache policy deleteFqdnCache

Deletes matching DNS lookups from the policy-generation cache.

Deletes matching DNS lookups from the cache, optionally restricted by
DNS name. The removed IP data will no longer be used in generated
policies.


*/
type DeleteFqdnCache struct {
	Context *middleware.Context
	Handler DeleteFqdnCacheHandler
}

func (o *DeleteFqdnCache) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteFqdnCacheParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
