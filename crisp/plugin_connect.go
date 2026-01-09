// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "time"
  "net/url"
)


// PluginConnectAccountData mapping
type PluginConnectAccountData struct {
  Data  *PluginConnectAccount  `json:"data,omitempty"`
}

// PluginConnectAccount mapping
type PluginConnectAccount struct {
  PluginID  *string  `json:"plugin_id,omitempty"`
}

// PluginConnectAllWebsitesData mapping
type PluginConnectAllWebsitesData struct {
  Data  *[]PluginConnectAllWebsites  `json:"data,omitempty"`
}

// PluginConnectAllWebsites mapping
type PluginConnectAllWebsites struct {
  WebsiteID  *string                        `json:"website_id,omitempty"`
  Token      *string                        `json:"token,omitempty"`
  Settings   *interface{}                   `json:"settings,omitempty"`
  Plan       *PluginConnectAllWebsitesPlan  `json:"plan,omitempty"`
}

// PluginConnectAllWebsitesPlan mapping
type PluginConnectAllWebsitesPlan struct {
  ID        *string  `json:"id,omitempty"`
  Name      *string  `json:"name,omitempty"`
  Trialing  *bool    `json:"trialing,omitempty"`
}

// PluginConnectWebsitesSinceData mapping
type PluginConnectWebsitesSinceData struct {
  Data  *[]PluginConnectWebsitesSince  `json:"data,omitempty"`
}

// PluginConnectWebsitesSince mapping
type PluginConnectWebsitesSince struct {
  PluginConnectAllWebsites
  Difference  *string  `json:"difference,omitempty"`
}

// PluginConnectEndpointsData mapping
type PluginConnectEndpointsData struct {
  Data  *PluginConnectEndpoints  `json:"data,omitempty"`
}

// PluginConnectEndpoints mapping
type PluginConnectEndpoints struct {
  Socket  *PluginConnectEndpointsSocket  `json:"socket,omitempty"`
  Rescue  *PluginConnectEndpointsRescue  `json:"rescue,omitempty"`
}

// PluginConnectEndpointsSocket mapping
type PluginConnectEndpointsSocket struct {
  App  *string  `json:"app,omitempty"`
}

// PluginConnectEndpointsRescue mapping
type PluginConnectEndpointsRescue struct {
  Socket  *PluginConnectEndpointsSocket  `json:"app,omitempty"`
}


// String returns the string representation of PluginConnectAccount
func (instance PluginConnectAccount) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginConnectAllWebsites
func (instance PluginConnectAllWebsites) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginConnectWebsitesSince
func (instance PluginConnectWebsitesSince) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginConnectEndpoints
func (instance PluginConnectEndpoints) String() string {
  return Stringify(instance)
}


// GetConnectAccount resolves the current plugin account information.
func (service *PluginService) GetConnectAccount() (*PluginConnectAccount, *Response, error) {
  url := "plugin/connect/account"
  req, _ := service.client.NewRequest("GET", url, nil)

  account := new(PluginConnectAccountData)
  resp, err := service.client.Do(req, account)
  if err != nil {
    return nil, resp, err
  }

  return account.Data, resp, err
}


// CheckConnectSessionValidity checks whether the connected plugin session is valid or not.
func (service *PluginService) CheckConnectSessionValidity() (*Response, error) {
  url := "plugin/connect/session"
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ListAllConnectWebsites lists all websites linked to connected plugin.
func (service *PluginService) ListAllConnectWebsites(pageNumber uint, filterConfigured bool, includePlan bool) (*[]PluginConnectAllWebsites, *Response, error) {
  var (
    filterConfiguredValue string
    includePlanValue string
  )

  if filterConfigured == true {
    filterConfiguredValue = "1"
  } else {
    filterConfiguredValue = "0"
  }

  if includePlan == true {
    includePlanValue = "1"
  } else {
    includePlanValue = "0"
  }

  url := fmt.Sprintf("plugin/connect/websites/all/%d?filter_configured=%s&include_plan=%s", pageNumber, url.QueryEscape(filterConfiguredValue), url.QueryEscape(includePlanValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  websites := new(PluginConnectAllWebsitesData)
  resp, err := service.client.Do(req, websites)
  if err != nil {
    return nil, resp, err
  }

  return websites.Data, resp, err
}


// ListConnectWebsitesSince lists the websites linked or unlinked or updated for connected plugin, since given date.
func (service *PluginService) ListConnectWebsitesSince(dateSince time.Time, filterConfigured bool, includePlan bool) (*[]PluginConnectWebsitesSince, *Response, error) {
  dateSinceFormat, err := dateSince.UTC().MarshalText()
  if err != nil {
    return nil, nil, err
  }

  var (
    filterConfiguredValue string
    includePlanValue string
  )

  if filterConfigured == true {
    filterConfiguredValue = "1"
  } else {
    filterConfiguredValue = "0"
  }

  if includePlan == true {
    includePlanValue = "1"
  } else {
    includePlanValue = "0"
  }

  url := fmt.Sprintf("plugin/connect/websites/since?date_since=%s&filter_configured=%s&include_plan=%s", url.QueryEscape(string(dateSinceFormat[:])), url.QueryEscape(filterConfiguredValue), url.QueryEscape(includePlanValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  websites := new(PluginConnectWebsitesSinceData)
  resp, err := service.client.Do(req, websites)
  if err != nil {
    return nil, resp, err
  }

  return websites.Data, resp, err
}


// GetConnectEndpoints resolves the current plugin endpoints information.
func (service *PluginService) GetConnectEndpoints() (*PluginConnectEndpoints, *Response, error) {
  url := "plugin/connect/endpoints"
  req, _ := service.client.NewRequest("GET", url, nil)

  endpoints := new(PluginConnectEndpointsData)
  resp, err := service.client.Do(req, endpoints)
  if err != nil {
    return nil, resp, err
  }

  return endpoints.Data, resp, err
}
