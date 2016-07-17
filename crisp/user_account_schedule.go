// Copyright 2016 Crisp IM. All rights reserved.
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
  UserID    *string                     `json:"user_id,omitempty"`
  Enabled   *bool                       `json:"enabled,omitempty"`
  Presence  *bool                       `json:"presence,omitempty"`
  Offset    *int                        `json:"offset,omitempty"`
  Days      *UserScheduleSettingsDays   `json:"days,omitempty"`
  Hours     *UserScheduleSettingsHours  `json:"hours,omitempty"`
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


// GetScheduleSettings gets user schedule settings. Those settings are used by Crisp to automatically schedule when user will be seen online or offline by website visitors.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-schedule-get
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
// Reference: https://docs.crisp.im/api/v1/#user-user-account-schedule-patch
func (service *UserService) UpdateScheduleSettings(enabled bool, presence bool, offset int, dayMonday bool, dayTuesday bool, dayWednesday bool, dayThursday bool, dayFriday bool, daySaturday bool, daySunday bool, hourFrom string, hourTo string) (*Response, error) {
  url := "user/account/schedule"
  req, _ := service.client.NewRequest("PATCH", url, UserScheduleSettings{Enabled: &enabled, Presence: &presence, Offset: &offset, Days: &UserScheduleSettingsDays{Monday: &dayMonday, Tuesday: &dayTuesday, Wednesday: &dayWednesday, Thursday: &dayThursday, Friday: &dayFriday, Saturday: &daySaturday, Sunday: &daySunday}, Hours: &UserScheduleSettingsHours{From: &hourFrom, To: &hourTo}})

  return service.client.Do(req, nil)
}
