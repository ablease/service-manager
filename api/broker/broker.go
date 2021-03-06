/*
 *    Copyright 2018 The Service Manager Authors
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

// Package broker contains logic for building the Service Manager Broker Management API
package broker

import (
	"net/http"

	"github.com/Peripli/service-manager/pkg/web"
)

// Routes returns slice of routes which handle broker operations
func (c *Controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   web.BrokersURL,
			},
			Handler: c.createBroker,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.BrokersURL + "/{broker_id}",
			},
			Handler: c.getBroker,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   web.BrokersURL,
			},
			Handler: c.listBrokers,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   web.BrokersURL,
			},
			Handler: c.deleteBrokers,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   web.BrokersURL + "/{broker_id}",
			},
			Handler: c.deleteBroker,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPatch,
				Path:   web.BrokersURL + "/{broker_id}",
			},
			Handler: c.patchBroker,
		},
	}
}
