// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserSessionParametersData mapping
type UserSessionParametersData struct {
  Data  *UserSessionParameters  `json:"data"`
}

// UserSessionParameters mapping
type UserSessionParameters struct {
  Identifier  *string  `json:"identifier"`
  Key         *string  `json:"key"`
}

// UserSessionLogin mapping
type UserSessionLogin struct {
  Email     *string  `json:"email"`
  Password  *string  `json:"password"`
}

// UserSessionRecover mapping
type UserSessionRecover struct {
  Email  *string  `json:"email"`
}


// CheckSessionValidity checks whether the user is logged in or not, and whether his session is valid or not.
// Reference: https://docs.crisp.im/api/v1/#user-user-session-head
func (service *UserService) CheckSessionValidity() (*Response, error) {
  url := "user/session"
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// CreateNewSession logins to user account and create a new session.
// Reference: https://docs.crisp.im/api/v1/#user-user-session-post
func (service *UserService) CreateNewSession(email string, password string) (*UserSessionParameters, *Response, error) {
  url := "user/session/login"
  req, _ := service.client.NewRequest("POST", url, UserSessionLogin{Email: &email, Password: &password})

  session := new(UserSessionParametersData)
  resp, err := service.client.Do(req, session)
  if err != nil {
    return nil, resp, err
  }

  return session.Data, resp, err
}


// DestroySession logouts from user account and destroys current session.
// Reference: https://docs.crisp.im/api/v1/#user-user-session-post-1
func (service *UserService) DestroySession() (*Response, error) {
  url := "user/session/logout"
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// RecoverSession recovers an user account from which we are locked out. A password recovery email is sent.
// Reference: https://docs.crisp.im/api/v1/#user-user-session-post-2
func (service *UserService) RecoverSession(email string) (*Response, error) {
  url := "user/session/recover"
  req, _ := service.client.NewRequest("POST", url, UserSessionRecover{Email: &email})

  return service.client.Do(req, nil)
}
