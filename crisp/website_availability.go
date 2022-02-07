// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteAvailabilityStatusData mapping
type WebsiteAvailabilityStatusData struct {
  Data  *WebsiteAvailabilityStatus  `json:"data,omitempty"`
}

// WebsiteAvailabilityStatus mapping
type WebsiteAvailabilityStatus struct {
  Status  *string  `json:"status,omitempty"`
  Since   *uint64  `json:"since,omitempty"`
}

// WebsiteAvailabilityOperatorsData mapping
type WebsiteAvailabilityOperatorsData struct {
  Data  *[]WebsiteAvailabilityOperator  `json:"data,omitempty"`
}

// WebsiteAvailabilityOperator mapping
type WebsiteAvailabilityOperator struct {
  UserID  *string                           `json:"user_id,omitempty"`
  Type    *string                           `json:"type,omitempty"`
  Time    *WebsiteAvailabilityOperatorTime  `json:"time,omitempty"`
}

// WebsiteAvailabilityOperatorTime mapping
type WebsiteAvailabilityOperatorTime struct {
  For    *uint32  `json:"for,omitempty"`
  Since  *uint64  `json:"since,omitempty"`
}


// String returns the string representation of WebsiteAvailabilityStatus
func (instance WebsiteAvailabilityStatus) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteAvailabilityOperator
func (instance WebsiteAvailabilityOperator) String() string {
  return Stringify(instance)
}


// GetWebsiteAvailabilityStatus resolves the website availability status. This tells whether the chatbox is seen as online or away by visitors.
func (service *WebsiteService) GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error) {
  url := fmt.Sprintf("website/%s/availability/status", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  availability := new(WebsiteAvailabilityStatusData)
  resp, err := service.client.Do(req, availability)
  if err != nil {
    return nil, resp, err
  }

  return availability.Data, resp, err
}


// ListWebsiteOperatorAvailabilities lists the availabilities for website operators. This maps the availability of each operator in the website.
func (service *WebsiteService) ListWebsiteOperatorAvailabilities(websiteID string) (*[]WebsiteAvailabilityOperator, *Response, error) {
  url := fmt.Sprintf("website/%s/availability/operators", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  operators := new(WebsiteAvailabilityOperatorsData)
  resp, err := service.client.Do(req, operators)
  if err != nil {
    return nil, resp, err
  }

  return operators.Data, resp, err
}
