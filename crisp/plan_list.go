// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// PlanListData mapping
type PlanListData struct {
  Data  *[]PlanInformation  `json:"data,omitempty"`
}


// ListPlans lists available plans
func (service *PlanService) ListPlans() (*[]PlanInformation, *Response, error) {
  url := "plans/list"
  req, _ := service.client.NewRequest("GET", url, nil)

  plans := new(PlanListData)
  resp, err := service.client.Do(req, plans)
  if err != nil {
    return nil, resp, err
  }

  return plans.Data, resp, err
}
