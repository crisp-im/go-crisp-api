// Copyright 2018 Crisp IM SAS All rights reserved.
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
  ID              *string                    `json:"id,omitempty"`
  URN             *string                    `json:"urn,omitempty"`
  Type            *string                    `json:"type,omitempty"`
  Category        *string                    `json:"category,omitempty"`
  Name            *string                    `json:"name,omitempty"`
  Summary         *string                    `json:"summary,omitempty"`
  Price           *uint16                    `json:"price,omitempty"`
  Plans           *[]PluginSubscriptionPlan  `json:"plans,omitempty"`
  Icon            *string                    `json:"icon,omitempty"`
  WebsiteURL      *string                    `json:"website_url,omitempty"`
  ContactURL      *string                    `json:"contact_url,omitempty"`
  TermsURL        *string                    `json:"terms_url,omitempty"`
  PrivacyURL      *string                    `json:"privacy_url,omitempty"`
  HelpURL         *string                    `json:"help_url,omitempty"`
  VideoURL        *string                    `json:"video_url,omitempty"`
  Configurable    *bool                      `json:"configurable,omitempty"`
  Since           *string                    `json:"since,omitempty"`
  Active          *bool                      `json:"active,omitempty"`
  BillPeriod      *string                    `json:"bill_period,omitempty"`
  BillValidUntil  *string                    `json:"bill_valid_until,omitempty"`
  WebsiteID       *string                    `json:"website_id,omitempty"`
  CardID          *string                    `json:"card_id,omitempty"`
}

// PluginSubscriptionPlan mapping
type PluginSubscriptionPlan struct {
  ID     *string  `json:"id,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Price  *uint16  `json:"price,omitempty"`
}

// PluginSubscriptionCreate mapping
type PluginSubscriptionCreate struct {
  PluginID  *string  `json:"plugin_id,omitempty"`
}

// PluginSubscriptionBillUsageReport mapping
type PluginSubscriptionBillUsageReport struct {
  Name   *string   `json:"name,omitempty"`
  Units  *uint32   `json:"units,omitempty"`
  Price  *float32  `json:"price,omitempty"`
}

// PluginSubscriptionChannelForward mapping
type PluginSubscriptionChannelForward struct {
  Namespace   *string       `json:"namespace,omitempty"`
  Identifier  *string       `json:"identifier,omitempty"`
  Payload     *interface{}  `json:"payload,omitempty"`
}

// PluginSubscriptionEventDispatch mapping
type PluginSubscriptionEventDispatch struct {
  Name  *string       `json:"name,omitempty"`
  Data  *interface{}  `json:"data,omitempty"`
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

// PluginBillUsageData mapping
type PluginBillUsageData struct {
  Data  *[]PluginBillUsage  `json:"data,omitempty"`
}

// PluginBillUsage mapping
type PluginBillUsage struct {
  Name   *string   `json:"name,omitempty"`
  Units  *uint32   `json:"units,omitempty"`
  Price  *float32  `json:"price,omitempty"`
}

// PluginAttestProvenanceData mapping
type PluginAttestProvenanceData struct {
  Data  *PluginAttestProvenance  `json:"data,omitempty"`
}

// PluginAttestProvenance mapping
type PluginAttestProvenance struct {
  Payload   *PluginAttestProvenancePayload  `json:"payload,omitempty"`
  Signed    *string                         `json:"signed,omitempty"`
  ExpireAt  *uint64                         `json:"expire_at,omitempty"`
}

// PluginAttestProvenancePayload mapping
type PluginAttestProvenancePayload struct {
  Website  *PluginAttestProvenancePayloadWebsite  `json:"website,omitempty"`
  User     *PluginAttestProvenancePayloadUser     `json:"user,omitempty"`
  Plugin   *PluginAttestProvenancePayloadPlugin   `json:"plugin,omitempty"`
}

// PluginAttestProvenancePayloadWebsite mapping
type PluginAttestProvenancePayloadWebsite struct {
  WebsiteID  *string  `json:"website_id,omitempty"`
}

// PluginAttestProvenancePayloadUser mapping
type PluginAttestProvenancePayloadUser struct {
  UserID  *string  `json:"user_id,omitempty"`
  Role    *string  `json:"role,omitempty"`
}

// PluginAttestProvenancePayloadPlugin mapping
type PluginAttestProvenancePayloadPlugin struct {
  PluginID  *string  `json:"plugin_id,omitempty"`
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


// GetPluginUsageBills acquires all non-settled (ie. open) usage bills for a subscribed plugin.
func (service *PluginService) GetPluginUsageBills(websiteID string, pluginID string) (*[]PluginBillUsage, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/bill/usage", websiteID, pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  bills := new(PluginBillUsageData)
  resp, err := service.client.Do(req, bills)
  if err != nil {
    return nil, resp, err
  }

  return bills.Data, resp, err
}


// ReportPluginUsageToBill reports a billable usage for a website using a subscribed plugin.
func (service *PluginService) ReportPluginUsageToBill(websiteID string, pluginID string, usage PluginSubscriptionBillUsageReport) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/bill/usage", websiteID, pluginID)
  req, _ := service.client.NewRequest("POST", url, usage)

  return service.client.Do(req, nil)
}


// GetPluginAttestProvenance attests for plugin user provenance, using the plugin signature secret
func (service *PluginService) GetPluginAttestProvenance(websiteID string, pluginID string) (*PluginAttestProvenance, *Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/attest/provenance", websiteID, pluginID)
  req, _ := service.client.NewRequest("GET", url, nil)

  provenance := new(PluginAttestProvenanceData)
  resp, err := service.client.Do(req, provenance)
  if err != nil {
    return nil, resp, err
  }

  return provenance.Data, resp, err
}


// ForwardPluginPayloadToChannel forwards generic payload given generic namespace to plugin channel.
func (service *PluginService) ForwardPluginPayloadToChannel(websiteID string, pluginID string, payload PluginSubscriptionChannelForward) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/channel", websiteID, pluginID)
  req, _ := service.client.NewRequest("POST", url, payload)

  return service.client.Do(req, nil)
}


// DispatchPluginEvent dispatches a generic data event for plugin.
func (service *PluginService) DispatchPluginEvent(websiteID string, pluginID string, payload PluginSubscriptionEventDispatch) (*Response, error) {
  url := fmt.Sprintf("plugins/subscription/%s/%s/event", websiteID, pluginID)
  req, _ := service.client.NewRequest("POST", url, payload)

  return service.client.Do(req, nil)
}
