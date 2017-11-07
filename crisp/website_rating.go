// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteRatingSessionData mapping
type WebsiteRatingSessionData struct {
  Data  *WebsiteRatingSession  `json:"data,omitempty"`
}

// WebsiteRatingSession mapping
type WebsiteRatingSession struct {
  Stars      *uint8   `json:"stars,omitempty"`
  Comment    *string  `json:"comment,omitempty"`
  CreatedAt  *uint    `json:"created_at,omitempty"`
}

// WebsiteRatingCreateSession mapping
type WebsiteRatingCreateSession struct {
  Stars    *uint8   `json:"stars,omitempty"`
  Comment  *string  `json:"comment,omitempty"`
}


// String returns the string representation of WebsiteRatingSession
func (instance WebsiteRatingSession) String() string {
  return Stringify(instance)
}


// ResolveSessionRating gets a session rating for website. Used to retrieve rating details from a given session.
func (service *WebsiteService) ResolveSessionRating(websiteID string, sessionID string) (*WebsiteRatingSession, *Response, error) {
  url := fmt.Sprintf("website/%s/rating/session/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  rating := new(WebsiteRatingSessionData)
  resp, err := service.client.Do(req, rating)
  if err != nil {
    return nil, resp, err
  }

  return rating.Data, resp, err
}


// SubmitSessionRating submits a session rating for website. Used for session users to publish their own website rating.
func (service *WebsiteService) SubmitSessionRating(websiteID string, sessionID string, websiteRatingSession WebsiteRatingCreateSession) (*Response, error) {
  url := fmt.Sprintf("website/%s/rating/session/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("PUT", url, websiteRatingSession)

  return service.client.Do(req, nil)
}


// DeleteSessionRating deletes a session rating for website. Used for session users to revoke their own website rating.
func (service *WebsiteService) DeleteSessionRating(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/rating/session/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
