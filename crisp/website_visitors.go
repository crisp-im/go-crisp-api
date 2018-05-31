// Copyright 2018 Crisp IM SARL All rights reserved.
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
  Count   *uint32  `json:"count,omitempty"`
  Active  *uint32  `json:"active,omitempty"`
}

// WebsiteVisitorListData mapping
type WebsiteVisitorListData struct {
  Data  *[]WebsiteVisitor  `json:"data,omitempty"`
}

// WebsiteVisitor mapping
type WebsiteVisitor struct {
  SessionID     *string                     `json:"session_id,omitempty"`
  Nickname      *string                     `json:"nickname,omitempty"`
  Email         *string                     `json:"email,omitempty"`
  Avatar        *string                     `json:"avatar,omitempty"`
  Useragent     *string                     `json:"useragent,omitempty"`
  Initiated     *bool                       `json:"initiated,omitempty"`
  Active        *bool                       `json:"active,omitempty"`
  LastPage      *WebsiteVisitorLastPage     `json:"last_page,omitempty"`
  Geolocation   *WebsiteVisitorGeolocation  `json:"geolocation,omitempty"`
  Timezone      *int16                      `json:"timezone,omitempty"`
  Capabilities  *[]string                   `json:"capabilities,omitempty"`
  Locales       *[]string                   `json:"locales,omitempty"`
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

// WebsiteVisitorsMapPointsData mapping
type WebsiteVisitorsMapPointsData struct {
  Data  *[]WebsiteVisitorsMapPoint  `json:"data,omitempty"`
}

// WebsiteVisitorsMapPoint mapping
type WebsiteVisitorsMapPoint struct {
  Visitors     *WebsiteVisitorsMapPointVisitors     `json:"visitors"`
  Geolocation  *WebsiteVisitorsMapPointGeolocation  `json:"geolocation"`
}

// WebsiteVisitorsMapPointGeolocation mapping
type WebsiteVisitorsMapPointGeolocation struct {
  Coordinates  *WebsiteVisitorsMapPointGeolocationCoordinates  `json:"coordinates,omitempty"`
  City         *string                                              `json:"city,omitempty"`
  Region       *string                                              `json:"region,omitempty"`
  Country      *string                                              `json:"country,omitempty"`
}

// WebsiteVisitorsMapPointGeolocationCoordinates mapping
type WebsiteVisitorsMapPointGeolocationCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}

// WebsiteVisitorsMapPointVisitors mapping
type WebsiteVisitorsMapPointVisitors struct {
  Count       *uint32    `json:"count,omitempty"`
  Threshold   *uint32    `json:"threshold,omitempty"`
  SessionIDs  *[]string  `json:"session_ids,omitempty"`
}

// WebsiteVisitorsBlockedAllData mapping
type WebsiteVisitorsBlockedAllData struct {
  Data  *[]WebsiteVisitorsBlocked  `json:"data,omitempty"`
}

// WebsiteVisitorsBlockedRuleData mapping
type WebsiteVisitorsBlockedRuleData struct {
  Data  *WebsiteVisitorsBlocked  `json:"data,omitempty"`
}

// WebsiteVisitorsBlocked mapping
type WebsiteVisitorsBlocked struct {
  Rule     *[]string  `json:"rule,omitempty"`
  Blocked  *uint32    `json:"blocked,omitempty"`
}


// String returns the string representation of WebsiteVisitorCount
func (instance WebsiteVisitorCount) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteVisitor
func (instance WebsiteVisitor) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteVisitorsMapPoint
func (instance WebsiteVisitorsMapPoint) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteVisitorsBlocked
func (instance WebsiteVisitorsBlocked) String() string {
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


// PinpointVisitorsOnMapWide maps visitors in a geographical area, given a geographical center and a map radius (wide variant).
func (service *WebsiteService) PinpointVisitorsOnMapWide(websiteID string) (*[]WebsiteVisitorsMapPoint, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/map", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  points := new(WebsiteVisitorsMapPointsData)
  resp, err := service.client.Do(req, points)
  if err != nil {
    return nil, resp, err
  }

  return points.Data, resp, err
}


// PinpointVisitorsOnMapArea maps visitors in a geographical area, given a geographical center and a map radius (area variant).
func (service *WebsiteService) PinpointVisitorsOnMapArea(websiteID string, centerLongitude float32, centerLatitude float32, centerRadius uint) (*[]WebsiteVisitorsMapPoint, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/map?center_longitude=%f&center_latitude=%f&center_radius=%d", websiteID, centerLongitude, centerLatitude, centerRadius)
  req, _ := service.client.NewRequest("GET", url, nil)

  points := new(WebsiteVisitorsMapPointsData)
  resp, err := service.client.Do(req, points)
  if err != nil {
    return nil, resp, err
  }

  return points.Data, resp, err
}


// CountBlockedVisitors counts all blocked visitors on website, by rule list.
func (service *WebsiteService) CountBlockedVisitors(websiteID string) (*[]WebsiteVisitorsBlocked, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/blocked", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  blocked := new(WebsiteVisitorsBlockedAllData)
  resp, err := service.client.Do(req, blocked)
  if err != nil {
    return nil, resp, err
  }

  return blocked.Data, resp, err
}


// CountBlockedVisitorsInRule counts all blocked visitors in rule list on website.
func (service *WebsiteService) CountBlockedVisitorsInRule(websiteID string, rule string) (*WebsiteVisitorsBlocked, *Response, error) {
  url := fmt.Sprintf("website/%s/visitors/blocked/%s", websiteID, rule)
  req, _ := service.client.NewRequest("GET", url, nil)

  blocked := new(WebsiteVisitorsBlockedRuleData)
  resp, err := service.client.Do(req, blocked)
  if err != nil {
    return nil, resp, err
  }

  return blocked.Data, resp, err
}


// ClearBlockedVisitorsInRule clears all blocked visitors in rule list on website.
func (service *PluginService) ClearBlockedVisitorsInRule(websiteID string, rule string) (*Response, error) {
  url := fmt.Sprintf("website/%s/visitors/blocked/%s", websiteID, rule)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
