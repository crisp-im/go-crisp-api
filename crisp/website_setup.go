// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteSetupData mapping
type WebsiteSetupData struct {
  Data  *WebsiteSetup  `json:"data,omitempty"`
}

// WebsiteSetup mapping
type WebsiteSetup struct {
  Disabled  *bool                `json:"disabled,omitempty"`
  States    *WebsiteSetupStates  `json:"states,omitempty"`
}

// WebsiteSetupStates mapping
type WebsiteSetupStates struct {
  Chatbox        *WebsiteSetupStatesItem  `json:"chatbox,omitempty"`
  Messenger      *WebsiteSetupStatesItem  `json:"messenger,omitempty"`
  Email          *WebsiteSetupStatesItem  `json:"email,omitempty"`
  Slack          *WebsiteSetupStatesItem  `json:"slack,omitempty"`
  AutoResponder  *WebsiteSetupStatesItem  `json:"autoresponder,omitempty"`
  Helpdesk       *WebsiteSetupStatesItem  `json:"helpdesk,omitempty"`
  Status         *WebsiteSetupStatesItem  `json:"status,omitempty"`
}

// WebsiteSetupStatesItem mapping
type WebsiteSetupStatesItem struct {
  Done   *bool    `json:"done,omitempty"`
  Trial  *uint16  `json:"trial,omitempty"`
}

// WebsiteSetupUpdate mapping
type WebsiteSetupUpdate struct {
  Disabled  bool  `json:"disabled,omitempty"`
}


// String returns the string representation of WebsiteSetup
func (instance WebsiteSetup) String() string {
  return Stringify(instance)
}


// GetWebsiteSetup resolves the current setup for a website.
func (service *WebsiteService) GetWebsiteSetup(websiteID string) (*WebsiteSetup, *Response, error) {
  url := fmt.Sprintf("website/%s/setup", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  setup := new(WebsiteSetupData)
  resp, err := service.client.Do(req, setup)
  if err != nil {
    return nil, resp, err
  }

  return setup.Data, resp, err
}


// UpdateWebsiteSetup updates the current setup for a website.
func (service *WebsiteService) UpdateWebsiteSetup(websiteID string, setup WebsiteSetupUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/setup", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, setup)

  return service.client.Do(req, nil)
}
