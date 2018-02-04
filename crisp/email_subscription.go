// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// SubscriptionStatusData mapping
type SubscriptionStatusData struct {
  Data  *SubscriptionStatus  `json:"data,omitempty"`
}

// SubscriptionStatus mapping
type SubscriptionStatus struct {
  Subscribed  *bool  `json:"subscribed,omitempty"`
}


// String returns the string representation of SubscriptionStatus
func (instance SubscriptionStatus) String() string {
  return Stringify(instance)
}


// GetSubscriptionStatus resolves current subscription status (subscribed or unsubscribed).
func (service *EmailService) GetSubscriptionStatus(emailHash string, key string, websiteID string) (*SubscriptionStatus, *Response, error) {
  url := fmt.Sprintf("email/%s/subscription/%s?website_id=%s", emailHash, key, url.QueryEscape(websiteID))
  req, _ := service.client.NewRequest("GET", url, nil)

  subscription := new(SubscriptionStatusData)
  resp, err := service.client.Do(req, subscription)
  if err != nil {
    return nil, resp, err
  }

  return subscription.Data, resp, err
}


// UpdateSubscriptionStatus updates current subscription status (subscribe or unsubscribe).
func (service *EmailService) UpdateSubscriptionStatus(emailHash string, key string, websiteID string, subscribed bool) (*Response, error) {
  url := fmt.Sprintf("email/%s/subscription/%s?website_id=%s", emailHash, key, url.QueryEscape(websiteID))
  req, _ := service.client.NewRequest("PATCH", url, SubscriptionStatus{&subscribed})

  return service.client.Do(req, nil)
}
