// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteChannelEmailData mapping
type WebsiteChannelEmailData struct {
  Data  *WebsiteChannelEmail  `json:"data,omitempty"`
}

// WebsiteChannelEmail mapping
type WebsiteChannelEmail struct {
  Email  *string  `json:"email,omitempty"`
}


// String returns the string representation of WebsiteChannelEmail
func (instance WebsiteChannelEmail) String() string {
  return Stringify(instance)
}


// GetWebsiteEmailChannel resolves the website email channel value.
func (service *WebsiteService) GetWebsiteEmailChannel(websiteID string) (*WebsiteChannelEmail, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  channel := new(WebsiteChannelEmailData)
  resp, err := service.client.Do(req, channel)
  if err != nil {
    return nil, resp, err
  }

  return channel.Data, resp, err
}
