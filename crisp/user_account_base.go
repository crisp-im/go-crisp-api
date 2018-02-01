// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserAccountData mapping
type UserAccountData struct {
  Data  *UserAccount  `json:"data,omitempty"`
}

// UserAccount mapping
type UserAccount struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// UserAccountCreate mapping
type UserAccountCreate struct {
  Email         string   `json:"email,omitempty"`
  Password      string   `json:"password,omitempty"`
  FirstName     string   `json:"first_name,omitempty"`
  LastName      string   `json:"last_name,omitempty"`
  AffiliatesID  *string  `json:"affiliates_id,omitempty"`
}


// String returns the string representation of UserAccount
func (instance UserAccount) String() string {
  return Stringify(instance)
}


// GetUserAccount resolves the current user account information.
func (service *UserService) GetUserAccount() (*UserAccount, *Response, error) {
  url := "user/account"
  req, _ := service.client.NewRequest("GET", url, nil)

  account := new(UserAccountData)
  resp, err := service.client.Do(req, account)
  if err != nil {
    return nil, resp, err
  }

  return account.Data, resp, err
}


// CreateUserAccount creates a new Crisp user account (operator account).
func (service *UserService) CreateUserAccount(user UserAccountCreate) (*Response, error) {
  url := "user/account"
  req, _ := service.client.NewRequest("POST", url, user)

  return service.client.Do(req, nil)
}
