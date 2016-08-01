// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// PlanInformationData mapping
type PlanInformationData struct {
  Data  *PlanInformation  `json:"data,omitempty"`
}

// PlanInformation mapping
type PlanInformation struct {
  ID           *string    `json:"id,omitempty"`
  Name         *string    `json:"name,omitempty"`
  Price        *uint      `json:"price,omitempty"`
  Since        *string    `json:"since,omitempty"`
}


// String returns the string representation of PlanInformation
func (instance PlanInformation) String() string {
  return Stringify(instance)
}


// GetPlanInformation resolves plan information.
func (service *PlanService) GetPlanInformation(planID string) (*PlanInformation, *Response, error) {
  url := fmt.Sprintf("plan/%s", planID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plan := new(PlanInformationData)
  resp, err := service.client.Do(req, plan)
  if err != nil {
    return nil, resp, err
  }

  return plan.Data, resp, err
}
