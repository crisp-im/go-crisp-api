// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteVisitorCountData mapping
type WebsiteVisitorCountData struct {
  Data  *WebsiteVisitorCount  `json:"data,omitempty"`
}

// WebsiteVisitorCount mapping
type WebsiteVisitorCount struct {
  Count   *uint  `json:"count,omitempty"`
  Active  *uint  `json:"active,omitempty"`
}

// WebsiteVisitorListData mapping
type WebsiteVisitorListData struct {
  Data  *[]WebsiteVisitor  `json:"data,omitempty"`
}

// WebsiteVisitor mapping
type WebsiteVisitor struct {
  SessionID    *string                     `json:"session_id,omitempty"`
  Nickname     *string                     `json:"nickname,omitempty"`
  Email        *string                     `json:"email,omitempty"`
  Avatar       *string                     `json:"avatar,omitempty"`
  Useragent    *string                     `json:"useragent,omitempty"`
  Initiated    *bool                       `json:"initiated,omitempty"`
  Active       *bool                       `json:"active,omitempty"`
  LastPage     *WebsiteVisitorLastPage     `json:"last_page,omitempty"`
  Geolocation  *WebsiteVisitorGeolocation  `json:"geolocation,omitempty"`
  Timezone     *int16                      `json:"timezone,omitempty"`
  Locales      *[]string                   `json:"locales,omitempty"`
}

// WebsiteVisitorLastPage mapping
type WebsiteVisitorLastPage struct {
  PageTitle  *string  `json:"page_title,omitempty"`
  PageURL    *string  `json:"page_url,omitempty"`
}

// WebsiteVisitorGeolocation mapping
type WebsiteVisitorGeolocation struct {
  Coordinates  *WebsiteVisitorGeolocationCoordinates  `json:"coordinates,omitempty"`
  City         *string                                `json:"city,omitempty"`
  Region       *string                                `json:"region,omitempty"`
  Country      *string                                `json:"country,omitempty"`
}

// WebsiteVisitorGeolocationCoordinates mapping
type WebsiteVisitorGeolocationCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}


// String returns the string representation of WebsiteVisitorCount
func (instance WebsiteVisitorCount) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteVisitor
func (instance WebsiteVisitor) String() string {
  return Stringify(instance)
}


// CountVisitors counts visitors currently on website.
func (service *WebsiteService) CountVisitors(websiteID string) (*WebsiteVisitorCount, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/count", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  visitors := new(WebsiteVisitorCountData)
  resp, err := service.client.Do(req, visitors)
  if err != nil {
    return nil, resp, err
  }

  return visitors.Data, resp, err
}


// ListVisitors lists visitors currently on website.
func (service *WebsiteService) ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/list/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  visitors := new(WebsiteVisitorListData)
  resp, err := service.client.Do(req, visitors)
  if err != nil {
    return nil, resp, err
  }

  return visitors.Data, resp, err
}


// RequestVisitorMapWide requests a map of visitors to be generated, in a geographical area (wide variant).
func (service *WebsiteService) RequestVisitorMapWide(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/visitors/map", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  resp, err := service.client.Do(req, nil)
  return resp, err
}


// RequestVisitorMapArea requests a map of visitors to be generated, in a geographical area (area variant).
func (service *WebsiteService) RequestVisitorMapArea(websiteID string, centerLongitude float32, centerLatitude float32, centerRadius uint) (*Response, error) {
  url := fmt.Sprintf("website/%s/visitors/map?center_longitude=%f&center_latitude=%f&center_radius=%d", websiteID, centerLongitude, centerLatitude, centerRadius)
  req, _ := service.client.NewRequest("GET", url, nil)

  resp, err := service.client.Do(req, nil)
  return resp, err
}
