// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteInviteData mapping
type WebsiteInviteData struct {
  Data  *WebsiteInvite  `json:"data,omitempty"`
}

// WebsiteInvite mapping
type WebsiteInvite struct {
  Email  *string  `json:"email,omitempty"`
  Role   *string  `json:"role,omitempty"`
}


// GetInviteDetails gets details on a invite keypair. Useful to check validity of invite keypair.
// Reference: https://docs.crisp.im/api/v1/#website-website-invite-get
func (service *WebsiteService) GetInviteDetails(websiteID string, recoverIdentifier string, recoverKey string) (*WebsiteInvite, *Response, error) {
  url := fmt.Sprintf("website/%s/invite/%s/%s", websiteID, recoverIdentifier, recoverKey)
  req, _ := service.client.NewRequest("GET", url, nil)

  invite := new(WebsiteInviteData)
  resp, err := service.client.Do(req, invite)
  if err != nil {
    return nil, resp, err
  }

  return invite.Data, resp, err
}


// RedeemInvite redeems invite and join the website as operator.
// Reference: https://docs.crisp.im/api/v1/#website-website-invite-put
func (service *WebsiteService) RedeemInvite(websiteID string, recoverIdentifier string, recoverKey string) (*Response, error) {
  url := fmt.Sprintf("website/%s/invite/%s/%s", websiteID, recoverIdentifier, recoverKey)
  req, _ := service.client.NewRequest("PUT", url, nil)

  return service.client.Do(req, nil)
}


// DeleteInviteKeypair deletes an invite keypair. Useful to invalidate keys if you ignore invite and never use the keys to redeem invite.
// Reference: https://docs.crisp.im/api/v1/#website-website-invite-delete
func (service *WebsiteService) DeleteInviteKeypair(websiteID string, recoverIdentifier string, recoverKey string) (*Response, error) {
  url := fmt.Sprintf("website/%s/invite/%s/%s", websiteID, recoverIdentifier, recoverKey)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
