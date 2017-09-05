// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserProfileData mapping
type UserProfileData struct {
  Data  *UserProfile  `json:"data,omitempty"`
}

// UserProfile mapping
type UserProfile struct {
  UserID     *string  `json:"user_id,omitempty"`
  Avatar     *string  `json:"avatar,omitempty"`
  Email      *string  `json:"email,omitempty"`
  FirstName  *string  `json:"first_name,omitempty"`
  LastName   *string  `json:"last_name,omitempty"`
  Locale     *string  `json:"locale,omitempty"`
}

// UserProfileUpdate mapping
type UserProfileUpdate struct {
  FirstName  string  `json:"first_name,omitempty"`
  LastName   string  `json:"last_name,omitempty"`
  Email      string  `json:"email,omitempty"`
  Password   string  `json:"password,omitempty"`
  Avatar     string  `json:"avatar,omitempty"`
  Locale     string  `json:"locale,omitempty"`
}


// String returns the string representation of UserProfile
func (instance UserProfile) String() string {
  return Stringify(instance)
}


// GetProfile resolves user profile data.
func (service *UserService) GetProfile() (*UserProfile, *Response, error) {
  url := "user/account/profile"
  req, _ := service.client.NewRequest("GET", url, nil)

  profile := new(UserProfileData)
  resp, err := service.client.Do(req, profile)
  if err != nil {
    return nil, resp, err
  }

  return profile.Data, resp, err
}


// UpdateProfile updates user profile data.
func (service *UserService) UpdateProfile(profile UserProfileUpdate) (*Response, error) {
  url := "user/account/profile"
  req, _ := service.client.NewRequest("PATCH", url, profile)

  return service.client.Do(req, nil)
}
