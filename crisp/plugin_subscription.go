// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// PluginSubscriptionListData mapping
type PluginSubscriptionListData struct {
  Data  *[]PluginSubscription  `json:"data,omitempty"`
}

// PluginSubscriptionData mapping
type PluginSubscriptionData struct {
  Data  *PluginSubscription  `json:"data,omitempty"`
}

// PluginSubscription mapping
type PluginSubscription struct {
  ID            *string                   `json:"id,omitempty"`
  URN           *string                   `json:"urn,omitempty"`
  Type          *string                   `json:"type,omitempty"`
  Name          *string                   `json:"name,omitempty"`
  Description   *string                   `json:"description,omitempty"`
  Price         *uint                     `json:"price,omitempty"`
  Plans         *[]PluginInformationPlan  `json:"plans,omitempty"`
  Icon          *string                   `json:"icon,omitempty"`
  Configurable  *bool                     `json:"configurable,omitempty"`
  Since         *string                   `json:"since,omitempty"`
  Active        *bool                     `json:"active,omitempty"`
  WebsiteID     *string                   `json:"website_id,omitempty"`
  CardID        *string                   `json:"card_id,omitempty"`
}

// PluginSubscriptionCreate mapping
type PluginSubscriptionCreate struct {
  PluginID  *string  `json:"plugin_id,omitempty"`
}

// PluginSubscriptionSettingsData mapping
type PluginSubscriptionSettingsData struct {
  Data  *PluginSubscriptionSettings  `json:"data,omitempty"`
}

// PluginSubscriptionSettings mapping
type PluginSubscriptionSettings struct {
  PluginID          *string       `json:"plugin_id,omitempty"`
  WebsiteID         *string       `json:"website_id,omitempty"`
  Token             *string       `json:"token,omitempty"`
  Schema            *interface{}  `json:"schema,omitempty"`
  Settings          *interface{}  `json:"settings,omitempty"`
  SettingsFormURL   *string       `json:"settings_form_url,omitempty"`
  CallbackURL       *string       `json:"callback_url,omitempty"`
}


// String returns the string representation of PluginSubscription
func (instance PluginSubscription) String() string {
  return Stringify(instance)
}


// String returns the string representation of PluginSubscriptionSettings
func (instance PluginSubscriptionSettings) String() string {
  return Stringify(instance)
}


// ListAllActiveSubscriptions lists all active plugin subscriptions on all websites, linked to payment methods owned by the user, or from websites the user is member of.
func (service *PluginService) ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error) {
  url := "plugins/subscription"
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// ListSubscriptionsForWebsite lists plugin subscriptions for given website.
func (service *PluginService) ListSubscriptionsForWebsite(websiteID string) (*[]PluginSubscription, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// GetSubscriptionDetails resolves details on a given subscription.
func (service *PluginService) GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s", websiteID, pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// SubscribeWebsiteToPlugin subscribes a given website to a given plugin.
func (service *PluginService) SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s", websiteID)
  req, _ := service.client.NewRequest("POST", url, PluginSubscriptionCreate{PluginID: &pluginID})

  return service.client.Do(req, nil)
}


// UnsubscribePluginFromWebsite unsubscribes a given plugin from a given website.
func (service *PluginService) UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s", websiteID, pluginID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// GetSubscriptionSettings resolves plugin subscription settings. Used to read plugin configuration on a given website.
func (service *PluginService) GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/settings", websiteID, pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginSubscriptionSettingsData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// SaveSubscriptionSettings saves plugin subscription settings (overwrites existing settings). Used to configure a given plugin on a given website.
func (service *PluginService) SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/settings", websiteID, pluginID)
  req, _ := service.client.NewRequest("PUT", url, settings)

  return service.client.Do(req, nil)
}


// UpdateSubscriptionSettings updates plugin subscription settings (merges with existing settings). Used to configure a given plugin on a given website.
func (service *PluginService) UpdateSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/settings", websiteID, pluginID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}
