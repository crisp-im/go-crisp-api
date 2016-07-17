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


// GetNotificationSettings resolves the user notification settings.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-notification-get
func (service *UserService) GetNotificationSettings() (*UserNotificationSettings, *Response, error) {
  url := "user/account/notification"
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(UserNotificationSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// UpdateNotificationSettings updates the user notification settings.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-notification-patch
func (service *EmailService) UpdateNotificationSettings(disabled bool, messagesOnline bool, messagesOffline bool, messagesTranscript bool, sounds bool) (*Response, error) {
  url := "user/account/notification"
  req, _ := service.client.NewRequest("POST", url, UserNotificationSettings{Disabled: &disabled, MessagesOnline: &messagesOnline, MessagesOffline: &messagesOffline, MessagesTranscript: &messagesTranscript, Sounds: &sounds})

  return service.client.Do(req, nil)
}
