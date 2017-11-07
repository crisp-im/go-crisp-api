// Copyright 2017 Crisp IM SARL All rights reserved.
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
  Since   *uint    `json:"since,omitempty"`
}


// String returns the string representation of WebsiteAvailabilityStatus
func (instance WebsiteAvailabilityStatus) String() string {
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
