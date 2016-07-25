// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteVisitorListData mapping
type WebsiteVisitorListData struct {
  Data  *[]WebsiteVisitor  `json:"data,omitempty"`
}

// WebsiteVisitor mapping
type WebsiteVisitor struct {
  SessionID  *string                  `json:"session_id,omitempty"`
  Nickname   *string                  `json:"nickname,omitempty"`
  Email      *string                  `json:"email,omitempty"`
  Avatar     *string                  `json:"avatar,omitempty"`
  Cover      *string                  `json:"cover,omitempty"`
  IP         *string                  `json:"ip,omitempty"`
  Useragent  *string                  `json:"useragent,omitempty"`
  Initiated  *bool                    `json:"initiated,omitempty"`
  Location   *WebsiteVisitorLocation  `json:"location,omitempty"`
  Locales    *[]string                `json:"locales,omitempty"`
}

// WebsiteVisitorLocation mapping
type WebsiteVisitorLocation struct {
  City     *string  `json:"city,omitempty"`
  Country  *string  `json:"country,omitempty"`
}


// String returns the string representation of WebsiteVisitor
func (instance WebsiteVisitor) String() string {
  return Stringify(instance)
}


// ListVisitors lists visitors currently on website.
func (service *WebsiteService) ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  visitors := new(WebsiteVisitorListData)
  resp, err := service.client.Do(req, visitors)
  if err != nil {
    return nil, resp, err
  }

  return visitors.Data, resp, err
}
