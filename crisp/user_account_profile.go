// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserProfileData mapping
type UserProfileData struct {
  Data  *UserProfile  `json:"data"`
}

// UserProfile mapping
type UserProfile struct {
  UserID     *string  `json:"user_id"`
  Avatar     *string  `json:"avatar"`
  Email      *string  `json:"email"`
  FirstName  *string  `json:"first_name"`
  LastName   *string  `json:"last_name"`
}

// UserProfileSave mapping
type UserProfileSave struct {
  FirstName  *string  `json:"first_name"`
  LastName   *string  `json:"last_name"`
  Email      *string  `json:"email"`
  Password   *string  `json:"password"`
  Avatar     *string  `json:"avatar"`
}


// GetProfile resolves user profile data.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-profile-get
func (service *UserService) GetProfile() (*UserProfile, *Response, error) {
  url := "user/account/profile"
  req, _ := service.client.NewRequest("GET", url, nil)

  account := new(UserProfileData)
  resp, err := service.client.Do(req, account)
  if err != nil {
    return nil, resp, err
  }

  return account.Data, resp, err
}


// UpdateProfile updates user profile data.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-profile-patch
func (service *EmailService) UpdateProfile(firstName string, lastName string, email string, password string, avatar string) (*Response, error) {
  url := "user/account/profile"
  req, _ := service.client.NewRequest("PATCH", url, UserProfileSave{FirstName: &firstName, LastName: &lastName, Email: &email, Password: &password, Avatar: &avatar})

  return service.client.Do(req, nil)
}
