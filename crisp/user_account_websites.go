// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserWebsiteData mapping
type UserWebsiteData struct {
  Data  *[]UserWebsite  `json:"data,omitempty"`
}

// UserWebsite mapping
type UserWebsite struct {
  ID       *string    `json:"id,omitempty"`
  Name     *string    `json:"name,omitempty"`
  Domain   *string    `json:"domain,omitempty"`
  Logo     *string    `json:"logo,omitempty"`
  Members  *[]string  `json:"members,omitempty"`
  Count    *uint      `json:"count,omitempty"`
}


// ListWebsites lists the websites linked to user.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-websites-get
func (service *UserService) ListWebsites() (*[]UserWebsite, *Response, error) {
  url := "user/account/websites"
  req, _ := service.client.NewRequest("GET", url, nil)

  websites := new(UserWebsiteData)
  resp, err := service.client.Do(req, websites)
  if err != nil {
    return nil, resp, err
  }

  return websites.Data, resp, err
}
