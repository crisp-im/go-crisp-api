// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteOperatorListData mapping
type WebsiteOperatorListData struct {
  Data  *[]WebsiteOperatorListOne  `json:"data,omitempty"`
}

// WebsiteOperatorListOne mapping
type WebsiteOperatorListOne struct {
  Type     *string           `json:"type,omitempty"`
  Details  *WebsiteOperator  `json:"details,omitempty"`
}

// WebsiteOperatorData mapping
type WebsiteOperatorData struct {
  Data  *WebsiteOperator  `json:"data,omitempty"`
}

// WebsiteOperator mapping
type WebsiteOperator struct {
  UserID        *string  `json:"user_id,omitempty"`
  Email         *string  `json:"email,omitempty"`
  Avatar        *string  `json:"avatar,omitempty"`
  FirstName     *string  `json:"first_name,omitempty"`
  LastName      *string  `json:"last_name,omitempty"`
  Role          *string  `json:"role,omitempty"`
  Availability  *string  `json:"availability,omitempty"`
  HasToken      *bool    `json:"has_token,omitempty"`
}

// WebsiteOperatorsLastActiveListData mapping
type WebsiteOperatorsLastActiveListData struct {
  Data  *[]WebsiteOperatorsLastActiveListOne  `json:"data,omitempty"`
}

// WebsiteOperatorsLastActiveListOne mapping
type WebsiteOperatorsLastActiveListOne struct {
  UserID     *string  `json:"user_id,omitempty"`
  Avatar     *string  `json:"avatar,omitempty"`
  Nickname   *string  `json:"nickname,omitempty"`
  Timestamp  *uint64  `json:"timestamp,omitempty"`
}

// WebsiteOperatorEmail mapping
type WebsiteOperatorEmail struct {
  Recipient  *string                      `json:"recipient,omitempty"`
  UserID     *string                      `json:"user_id,omitempty"`
  Subject    *string                      `json:"subject,omitempty"`
  Message    *string                      `json:"message,omitempty"`
  Target     *WebsiteOperatorEmailTarget  `json:"target,omitempty"`
}

// WebsiteOperatorEmailTarget mapping
type WebsiteOperatorEmailTarget struct {
  Label  *string  `json:"label,omitempty"`
  URL    *string  `json:"url,omitempty"`
}

// WebsiteOperatorInvite mapping
type WebsiteOperatorInvite struct {
  Email   *string  `json:"email,omitempty"`
  Role    *string  `json:"role,omitempty"`
  Verify  *string  `json:"verify,omitempty"`
}

// WebsiteOperatorEdit mapping
type WebsiteOperatorEdit struct {
  Role  *string  `json:"role,omitempty"`
}


// String returns the string representation of WebsiteOperatorListOne
func (instance WebsiteOperatorListOne) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteOperator
func (instance WebsiteOperator) String() string {
  return Stringify(instance)
}


// ListWebsiteOperators lists all operator members of website.
func (service *WebsiteService) ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error) {
  url := fmt.Sprintf("website/%s/operators/list", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  operators := new(WebsiteOperatorListData)
  resp, err := service.client.Do(req, operators)
  if err != nil {
    return nil, resp, err
  }

  return operators.Data, resp, err
}


// ListLastActiveWebsiteOperators lists last active website operators.
func (service *WebsiteService) ListLastActiveWebsiteOperators(websiteID string) (*[]WebsiteOperatorsLastActiveListOne, *Response, error) {
  url := fmt.Sprintf("website/%s/operators/active", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  operators := new(WebsiteOperatorsLastActiveListData)
  resp, err := service.client.Do(req, operators)
  if err != nil {
    return nil, resp, err
  }

  return operators.Data, resp, err
}


// FlushLastActiveWebsiteOperators flushes the list of last active website operators.
func (service *WebsiteService) FlushLastActiveWebsiteOperators(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operators/active", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// SendEmailToWebsiteOperators sends an email to target website operators.
func (service *WebsiteService) SendEmailToWebsiteOperators(websiteID string, email WebsiteOperatorEmail) (*Response, error) {
  url := fmt.Sprintf("website/%s/operators/email", websiteID)
  req, _ := service.client.NewRequest("POST", url, email)

  return service.client.Do(req, nil)
}


// GetWebsiteOperator resolves a given website operator.
func (service *WebsiteService) GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error) {
  url := fmt.Sprintf("website/%s/operator/%s", websiteID, userID)
  req, _ := service.client.NewRequest("GET", url, nil)

  operator := new(WebsiteOperatorData)
  resp, err := service.client.Do(req, operator)
  if err != nil {
    return nil, resp, err
  }

  return operator.Data, resp, err
}


// InviteWebsiteOperator invites an email to join website as operator.
func (service *WebsiteService) InviteWebsiteOperator(websiteID string, email string, role string, verify string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operator", websiteID)
  req, _ := service.client.NewRequest("POST", url, WebsiteOperatorInvite{Email: &email, Role: &role, Verify: &verify})

  return service.client.Do(req, nil)
}


// ChangeOperatorRole changes the role of an existing operator. Useful to downgrade or upgrade an operator from/to owner role.
func (service *WebsiteService) ChangeOperatorRole(websiteID string, userID string, role string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operator/%s", websiteID, userID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteOperatorEdit{Role: &role})

  return service.client.Do(req, nil)
}


// UnlinkOperatorFromWebsite unlinks given operator from website. Note that the last operator in the website cannot be unlinked.
func (service *WebsiteService) UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operator/%s", websiteID, userID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
