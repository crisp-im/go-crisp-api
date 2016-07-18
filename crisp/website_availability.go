// Copyright 2016 Crisp IM. All rights reserved.
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
}


// String returns the string representation of WebsiteAvailabilityStatus
func (instance WebsiteAvailabilityStatus) String() string {
  return Stringify(instance)
}


// GetWebsiteAvailabilityStatus resolves the website availability status. This tells whether the chatbox is seen as online or away by visitors.
// Reference: https://docs.crisp.im/api/v1/#website-website-availability-get
func (service *WebsiteService) GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error) {
  url := fmt.Sprintf("website/%s/availability/status", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  billing := new(WebsiteAvailabilityStatusData)
  resp, err := service.client.Do(req, billing)
  if err != nil {
    return nil, resp, err
  }

  return billing.Data, resp, err
}
