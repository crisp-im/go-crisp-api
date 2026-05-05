// Copyright 2026 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteConnectEndpointsData mapping
type WebsiteConnectEndpointsData struct {
  Data  *WebsiteConnectEndpoints  `json:"data,omitempty"`
}

// WebsiteConnectEndpoints mapping
type WebsiteConnectEndpoints struct {
  Socket  *WebsiteConnectEndpointsSocket  `json:"socket,omitempty"`
  Rescue  *WebsiteConnectEndpointsRescue  `json:"rescue,omitempty"`
}

// WebsiteConnectEndpointsSocket mapping
type WebsiteConnectEndpointsSocket struct {
  App     *string  `json:"app,omitempty"`
  Stream  *string  `json:"stream,omitempty"`
}

// WebsiteConnectEndpointsRescue mapping
type WebsiteConnectEndpointsRescue struct {
  Socket  *WebsiteConnectEndpointsSocket  `json:"app,omitempty"`
}


// String returns the string representation of WebsiteConnectEndpoints
func (instance WebsiteConnectEndpoints) String() string {
  return Stringify(instance)
}


// GetConnectEndpoints resolves the current Website endpoints information.
func (service *WebsiteService) GetConnectEndpoints(websiteID string) (*WebsiteConnectEndpoints, *Response, error) {
  url := fmt.Sprintf("website/%s/connect/endpoints", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  endpoints := new(WebsiteConnectEndpointsData)
  resp, err := service.client.Do(req, endpoints)
  if err != nil {
    return nil, resp, err
  }

  return endpoints.Data, resp, err
}
