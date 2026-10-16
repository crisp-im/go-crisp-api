// Copyright 2026 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
)

// PlanSubscriptionData mapping
type PlanSubscriptionData struct {
	Data *PlanSubscription `json:"data,omitempty"`
}

// PlanSubscription mapping
type PlanSubscription struct {
	ID             *string                  `json:"id,omitempty"`
	Name           *string                  `json:"name,omitempty"`
	Price          *uint16                  `json:"price,omitempty"`
	Since          *string                  `json:"since,omitempty"`
	Trialing       *bool                    `json:"trialing,omitempty"`
	TrialEnd       *string                  `json:"trial_end,omitempty"`
	TrialEndDate   *string                  `json:"trial_end_date,omitempty"`
	BillPeriod     *string                  `json:"bill_period,omitempty"`
	BillValidUntil *string                  `json:"bill_valid_until,omitempty"`
	Active         *bool                    `json:"active,omitempty"`
	Sandbox        *bool                    `json:"sandbox,omitempty"`
	Website        *PlanSubscriptionWebsite `json:"website,omitempty"`
	CouponRedeemed *bool                    `json:"coupon_redeemed,omitempty"`
	CardID         *string                  `json:"card_id,omitempty"`
	Owner          *PlanSubscriptionOwner   `json:"owner,omitempty"`
}

// PlanSubscriptionWebsite mapping
type PlanSubscriptionWebsite struct {
	ID     *string `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Logo   *string `json:"logo,omitempty"`
}

// PlanSubscriptionOwner mapping
type PlanSubscriptionOwner struct {
	UserID    *string `json:"user_id,omitempty"`
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

// String returns the string representation of PlanSubscription
func (instance PlanSubscription) String() string {
	return Stringify(instance)
}

// GetPlanSubscriptionForWebsite resolves plan subscription for given website.
func (service *PlanService) GetPlanSubscriptionForWebsite(websiteID string) (*PlanSubscription, *Response, error) {
	url := fmt.Sprintf("plans/subscription/%s", websiteID)
	req, _ := service.client.NewRequest("GET", url, nil)

	plans := new(PlanSubscriptionData)
	resp, err := service.client.Do(req, plans)
	if err != nil {
		return nil, resp, err
	}

	return plans.Data, resp, err
}
