// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


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
  MessagesUnread      *bool    `json:"messages_unread,omitempty"`
  MessagesTranscript  *bool    `json:"messages_transcript,omitempty"`
  MessagesRating      *bool    `json:"messages_rating,omitempty"`
  BillingInvoice      *bool    `json:"billing_invoice,omitempty"`
  Sounds              *bool    `json:"sounds,omitempty"`
}

// UserNotificationSettingsUpdate mapping
type UserNotificationSettingsUpdate struct {
  Disabled            bool  `json:"disabled,omitempty"`
  MessagesOnline      bool  `json:"messages_online,omitempty"`
  MessagesOffline     bool  `json:"messages_offline,omitempty"`
  MessagesUnread      bool  `json:"messages_unread,omitempty"`
  MessagesTranscript  bool  `json:"messages_transcript,omitempty"`
  MessagesRating      bool  `json:"messages_rating,omitempty"`
  BillingInvoice      bool  `json:"billing_invoice,omitempty"`
  Sounds              bool  `json:"sounds,omitempty"`
}

// UserNotificationProviderAdd mapping
type UserNotificationProviderAdd struct {
  NotificationID  *string  `json:"notification_id,omitempty"`
}


// String returns the string representation of UserNotificationSettings
func (instance UserNotificationSettings) String() string {
  return Stringify(instance)
}


// GetNotificationSettings resolves the user notification settings.
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
func (service *UserService) UpdateNotificationSettings(notifications UserNotificationSettingsUpdate) (*Response, error) {
  url := "user/account/notification"
  req, _ := service.client.NewRequest("PATCH", url, notifications)

  return service.client.Do(req, nil)
}


// AddNotificationProvider adds a notification provider.
func (service *UserService) AddNotificationProvider(notificationID string) (*Response, error) {
  url := "user/account/notification/provider"
  req, _ := service.client.NewRequest("POST", url, UserNotificationProviderAdd{&notificationID})

  return service.client.Do(req, nil)
}


// DeleteNotificationProvider deletes a notification provider.
func (service *UserService) DeleteNotificationProvider(notificationID string) (*Response, error) {
  url := fmt.Sprintf("user/account/notification/provider/%s", notificationID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
