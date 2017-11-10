// Copyright 2017 Crisp IM SARL All rights reserved.
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
  ID        *string    `json:"id,omitempty"`
  Name      *string    `json:"name,omitempty"`
  Domain    *string    `json:"domain,omitempty"`
  Logo      *string    `json:"logo,omitempty"`
  Members   *[]string  `json:"members,omitempty"`
  Count     *uint      `json:"count,omitempty"`
  Assigned  *uint      `json:"assigned,omitempty"`
}


// String returns the string representation of UserWebsite
func (instance UserWebsite) String() string {
  return Stringify(instance)
}


// ListWebsites lists the websites linked to user.
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
