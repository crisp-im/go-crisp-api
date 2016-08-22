// Copyright 2016 Crisp IM, Inc. All rights reserved.
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


// String returns the string representation of UserAvailability
func (instance UserAvailability) String() string {
  return Stringify(instance)
}


// String returns the string representation of UserAvailabilityStatus
func (instance UserAvailabilityStatus) String() string {
  return Stringify(instance)
}


// GetUserAvailability resolves the current user availability. Useful to check if a Crisp app is currently connected to the account.
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
func (service *UserService) UpdateUserAvailability(availabilityType string, timeFor uint) (*Response, error) {
  url := "user/availability"
  req, _ := service.client.NewRequest("PATCH", url, UserAvailability{Type: &availabilityType, Time: &UserAvailabilityTime{For: &timeFor}})

  return service.client.Do(req, nil)
}


// GetUserAvailabilityStatus resolves the current user availability status. It differs from the raw user availability (the raw availability only tells if the user is connected on a Crisp app).
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
