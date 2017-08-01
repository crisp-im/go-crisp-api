// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserScheduleSettingsData mapping
type UserScheduleSettingsData struct {
  Data  *UserScheduleSettings  `json:"data,omitempty"`
}

// UserScheduleSettings mapping
type UserScheduleSettings struct {
  UserID    *string                       `json:"user_id,omitempty"`
  Enabled   *bool                         `json:"enabled,omitempty"`
  Presence  *bool                         `json:"presence,omitempty"`
  Stealth   *bool                         `json:"stealth,omitempty"`
  Offset    *int                          `json:"offset,omitempty"`
  Days      *UserScheduleSettingsDays     `json:"days,omitempty"`
  Hours     *[]UserScheduleSettingsHours  `json:"hours,omitempty"`
}

// UserScheduleSettingsDays mapping
type UserScheduleSettingsDays struct {
  Monday     *bool  `json:"monday,omitempty"`
  Tuesday    *bool  `json:"tuesday,omitempty"`
  Wednesday  *bool  `json:"wednesday,omitempty"`
  Thursday   *bool  `json:"thursday,omitempty"`
  Friday     *bool  `json:"friday,omitempty"`
  Saturday   *bool  `json:"saturday,omitempty"`
  Sunday     *bool  `json:"sunday,omitempty"`
}

// UserScheduleSettingsHours mapping
type UserScheduleSettingsHours struct {
  From  *string  `json:"from,omitempty"`
  To    *string  `json:"to,omitempty"`
}

// UserScheduleSettingsUpdate mapping
type UserScheduleSettingsUpdate struct {
  Enabled   bool                             `json:"enabled,omitempty"`
  Presence  bool                             `json:"presence,omitempty"`
  Stealth   bool                             `json:"stealth,omitempty"`
  Offset    int                              `json:"offset,omitempty"`
  Days      UserScheduleSettingsUpdateDays   `json:"days,omitempty"`
  Hours     UserScheduleSettingsUpdateHours  `json:"hours,omitempty"`
}

// UserScheduleSettingsUpdateDays mapping
type UserScheduleSettingsUpdateDays struct {
  Monday     bool  `json:"monday,omitempty"`
  Tuesday    bool  `json:"tuesday,omitempty"`
  Wednesday  bool  `json:"wednesday,omitempty"`
  Thursday   bool  `json:"thursday,omitempty"`
  Friday     bool  `json:"friday,omitempty"`
  Saturday   bool  `json:"saturday,omitempty"`
  Sunday     bool  `json:"sunday,omitempty"`
}

// UserScheduleSettingsUpdateHours mapping
type UserScheduleSettingsUpdateHours struct {
  From  string  `json:"from,omitempty"`
  To    string  `json:"to,omitempty"`
}


// String returns the string representation of UserScheduleSettings
func (instance UserScheduleSettings) String() string {
  return Stringify(instance)
}


// GetScheduleSettings gets user schedule settings. Those settings are used by Crisp to automatically schedule when user will be seen online or offline by website visitors.
func (service *UserService) GetScheduleSettings() (*UserScheduleSettings, *Response, error) {
  url := "user/account/schedule"
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(UserScheduleSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// UpdateScheduleSettings updates user schedule settings.
func (service *UserService) UpdateScheduleSettings(schedule UserScheduleSettingsUpdate) (*Response, error) {
  url := "user/account/schedule"
  req, _ := service.client.NewRequest("PATCH", url, schedule)

  return service.client.Do(req, nil)
}
