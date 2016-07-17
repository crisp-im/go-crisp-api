// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
  "fmt"
)

// SubscriptionStatusData mapping
type SubscriptionStatusData struct {
  Data  *SubscriptionStatus  `json:"data"`
}

// SubscriptionStatus mapping
type SubscriptionStatus struct {
  Subscribed  *bool  `json:"subscribed"`
}


// GetSubscriptionStatus resolves current subscription status (subscribed or unsubscribed).
// Reference: https://docs.crisp.im/api/v1/#email-email-subscription-get
func (service *EmailService) GetSubscriptionStatus(emailHash string, key string) (*SubscriptionStatus, *Response, error) {
  url := fmt.Sprintf("email/%s/subscription/%s", emailHash, key)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugin := new(SubscriptionStatusData)
  resp, err := service.client.Do(req, plugin)
  if err != nil {
    return nil, resp, err
  }

  return plugin.Data, resp, err
}


// UpdateSubscriptionStatus updates current subscription status (subscribe or unsubscribe).
// Reference: https://docs.crisp.im/api/v1/#email-email-subscription-patch
func (service *EmailService) UpdateSubscriptionStatus(emailHash string, key string, subscribed bool) (*Response, error) {
  url := fmt.Sprintf("email/%s/subscription/%s", emailHash, key)
  req, _ := service.client.NewRequest("PATCH", url, SubscriptionStatus{&subscribed})

  return service.client.Do(req, nil)
}
