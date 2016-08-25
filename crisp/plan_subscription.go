// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// PlanSubscriptionListData mapping
type PlanSubscriptionListData struct {
  Data  *[]PlanSubscription  `json:"data,omitempty"`
}

// PlanSubscriptionData mapping
type PlanSubscriptionData struct {
  Data  *PlanSubscription  `json:"data,omitempty"`
}

// PlanSubscription mapping
type PlanSubscription struct {
  ID         *string                   `json:"id,omitempty"`
  Name       *string                   `json:"name,omitempty"`
  Price      *uint                     `json:"price,omitempty"`
  TrialDays  *uint                     `json:"trial_days,omitempty"`
  Since      *string                   `json:"since,omitempty"`
  TrialEnd   *string                   `json:"trial_end,omitempty"`
  Active     *bool                     `json:"active,omitempty"`
  Website    *PlanSubscriptionWebsite  `json:"website,omitempty"`
  CardID     *string                   `json:"card_id,omitempty"`
}

// PlanSubscriptionWebsite mapping
type PlanSubscriptionWebsite struct {
  ID      *string  `json:"id,omitempty"`
  Name    *string  `json:"name,omitempty"`
  Domain  *string  `json:"domain,omitempty"`
  Logo    *string  `json:"logo,omitempty"`
}

// PlanSubscriptionCreate mapping
type PlanSubscriptionCreate struct {
  PlanID  *string  `json:"plan_id,omitempty"`
}


// String returns the string representation of PlanSubscription
func (instance PlanSubscription) String() string {
  return Stringify(instance)
}


// ListAllActiveSubscriptions lists all active plan subscriptions on all websites, linked to payment methods owned by the user, or from websites the user is member of.
func (service *PlanService) ListAllActiveSubscriptions() (*[]PlanSubscription, *Response, error) {
  url := "plans/subscription"
  req, _ := service.client.NewRequest("GET", url, nil)

  plans := new(PlanSubscriptionListData)
  resp, err := service.client.Do(req, plans)
  if err != nil {
    return nil, resp, err
  }

  return plans.Data, resp, err
}


// GetSubscriptionForWebsite resolves plan subscription for given website.
func (service *PlanService) GetSubscriptionForWebsite(websiteID string) (*PlanSubscription, *Response, error) {
  url := fmt.Sprintf("plans/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plans := new(PlanSubscriptionData)
  resp, err := service.client.Do(req, plans)
  if err != nil {
    return nil, resp, err
  }

  return plans.Data, resp, err
}


// SubscribeWebsiteToPlan subscribes a given website to a given plan.
func (service *PlanService) SubscribeWebsiteToPlan(websiteID string, planID string) (*Response, error) {
  url := fmt.Sprintf("plans/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("POST", url, PlanSubscriptionCreate{PlanID: &planID})

  return service.client.Do(req, nil)
}


// UnsubscribePlanFromWebsite unsubscribes a given plan from a given website.
func (service *PlanService) UnsubscribePlanFromWebsite(websiteID string) (*Response, error) {
  url := fmt.Sprintf("plans/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
