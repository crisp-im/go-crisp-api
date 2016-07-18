// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserNotificationSettingsData mapping
type UserNotificationSettingsData struct {
  Data  *UserNotificationSettings  `json:"data,omitempty"`
}

// UserNotificationSettings mapping
type UserNotificationSettings struct {
  UserID              *string  `json:"user_id,omitempty"`
  Disabled            *bool    `json:"disabled,omitempty"`
  MessagesOnline      *bool    `json:"messages_online,omitempty"`
  MessagesOffline     *bool    `json:"messages_offline,omitempty"`
  MessagesTranscript  *bool    `json:"messages_transcript,omitempty"`
  Sounds              *bool    `json:"sounds,omitempty"`
}

// UserNotificationSettingsUpdate mapping
type UserNotificationSettingsUpdate struct {
  UserID              string  `json:"user_id,omitempty"`
  Disabled            bool    `json:"disabled,omitempty"`
  MessagesOnline      bool    `json:"messages_online,omitempty"`
  MessagesOffline     bool    `json:"messages_offline,omitempty"`
  MessagesTranscript  bool    `json:"messages_transcript,omitempty"`
  Sounds              bool    `json:"sounds,omitempty"`
}


// String returns the string representation of UserNotificationSettings
func (instance UserNotificationSettings) String() string {
  return Stringify(instance)
}


// GetNotificationSettings resolves the user notification settings.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-notification-get
func (service *UserService) GetNotificationSettings() (*UserNotificationSettings, *Response, error) {
  url := "user/account/notification"
  req, _ := service.client.NewRequest("GET", url, nil)

  notifications := new(UserNotificationSettingsData)
  resp, err := service.client.Do(req, notifications)
  if err != nil {
    return nil, resp, err
  }

  return notifications.Data, resp, err
}


// UpdateNotificationSettings updates the user notification settings.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-notification-patch
func (service *EmailService) UpdateNotificationSettings(notifications UserNotificationSettingsUpdate) (*Response, error) {
  url := "user/account/notification"
  req, _ := service.client.NewRequest("POST", url, notifications)

  return service.client.Do(req, nil)
}
