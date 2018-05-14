// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// UserTokenGenerateData mapping
type UserTokenGenerateData struct {
  Data  *UserTokenGenerate  `json:"data,omitempty"`
}

// UserTokenGenerate mapping
type UserTokenGenerate struct {
  Secret  *string  `json:"secret,omitempty"`
}

// UserTokenConfigure mapping
type UserTokenConfigure struct {
  Secret  string  `json:"secret,omitempty"`
}


// String returns the string representation of UserTokenGenerate
func (instance UserTokenGenerate) String() string {
  return Stringify(instance)
}


// CheckAccountTokenConfigured checks if account token is configured.
func (service *UserService) CheckAccountTokenConfigured() (*Response, error) {
  url := "user/account/token"
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ConfigureAccountToken configures account token, with the provided token secret key.
func (service *UserService) ConfigureAccountToken(secret string) (*Response, error) {
  url := "user/account/token"
  req, _ := service.client.NewRequest("PUT", url, UserTokenConfigure{Secret: secret})

  return service.client.Do(req, nil)
}


// UnconfigureAccountToken unconfigures account token.
func (service *UserService) UnconfigureAccountToken() (*Response, error) {
  url := "user/account/token"
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// GenerateAccountToken generates a new account token secret key, to be sent upon configuration.
func (service *UserService) GenerateAccountToken() (*UserTokenGenerate, *Response, error) {
  url := "user/account/token/generate"
  req, _ := service.client.NewRequest("POST", url, nil)

  token := new(UserTokenGenerateData)
  resp, err := service.client.Do(req, token)
  if err != nil {
    return nil, resp, err
  }

  return token.Data, resp, err
}


// VerifyAccountToken verifies provided token against provided secret key.
func (service *UserService) VerifyAccountToken(secret string, token string) (*Response, error) {
  url := fmt.Sprintf("user/account/token/verify?secret=%s&token=%s", url.QueryEscape(secret), url.QueryEscape(token))
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}
