// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserAvailabilityData mapping
type UserAvailabilityData struct {
  Data  *UserAvailability  `json:"data,omitempty"`
}

// UserAvailability mapping
type UserAvailability struct {
  UserID  *string                `json:"user_id,omitempty"`
  Type    *string                `json:"type,omitempty"`
  Time    *UserAvailabilityTime  `json:"time,omitempty"`
}

// UserAvailabilityTime mapping
type UserAvailabilityTime struct {
  For    *uint  `json:"for,omitempty"`
  Since  *uint  `json:"since,omitempty"`
}

// UserAvailabilityStatusData mapping
type UserAvailabilityStatusData struct {
  Data  *UserAvailabilityStatus  `json:"data,omitempty"`
}

// UserAvailabilityStatus mapping
type UserAvailabilityStatus struct {
  UserID  *string  `json:"user_id,omitempty"`
  Status  *string  `json:"status,omitempty"`
}


// GetUserAvailability resolves the current user availability. Useful to check if a Crisp app is currently connected to the account.
// Reference: https://docs.crisp.im/api/v1/#user-user-availability
func (service *UserService) GetUserAvailability() (*UserAvailability, *Response, error) {
  url := "user/availability"
  req, _ := service.client.NewRequest("GET", url, nil)

  availability := new(UserAvailabilityData)
  resp, err := service.client.Do(req, availability)
  if err != nil {
    return nil, resp, err
  }

  return availability.Data, resp, err
}


// UpdateUserAvailability updates the advertised user availability, for a defined period of time after which to automatically expire.
// Reference: https://docs.crisp.im/api/v1/#user-user-availability-patch
func (service *UserService) UpdateUserAvailability(availabilityType string, timeFor uint) (*Response, error) {
  url := "user/availability"
  req, _ := service.client.NewRequest("PATCH", url, UserAvailability{Type: &availabilityType, Time: &UserAvailabilityTime{For: &timeFor}})

  return service.client.Do(req, nil)
}


// GetUserAvailabilityStatus resolves the current user availability status. It differs from the raw user availability (the raw availability only tells if the user is connected on a Crisp app).
// Reference: https://docs.crisp.im/api/v1/#user-user-availability-get-1
func (service *UserService) GetUserAvailabilityStatus() (*UserAvailabilityStatus, *Response, error) {
  url := "user/availability/status"
  req, _ := service.client.NewRequest("GET", url, nil)

  status := new(UserAvailabilityStatusData)
  resp, err := service.client.Do(req, status)
  if err != nil {
    return nil, resp, err
  }

  return status.Data, resp, err
}
