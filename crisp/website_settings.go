// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteSettingsData mapping
type WebsiteSettingsData struct {
  Data  *WebsiteSettings  `json:"data,omitempty"`
}

// WebsiteSettings mapping
type WebsiteSettings struct {
  WebsiteID  *string                  `json:"website_id,omitempty"`
  Name       *string                  `json:"name,omitempty"`
  Domain     *string                  `json:"domain,omitempty"`
  Logo       *string                  `json:"logo,omitempty"`
  Contact    *WebsiteSettingsContact  `json:"contact,omitempty"`
  Chatbox    *WebsiteSettingsChatbox  `json:"chatbox,omitempty"`
}

// WebsiteSettingsContact mapping
type WebsiteSettingsContact struct {
  Email  *string  `json:"email,omitempty"`
  Phone  *string  `json:"phone,omitempty"`
}

// WebsiteSettingsChatbox mapping
type WebsiteSettingsChatbox struct {
  Tile                 *string    `json:"tile,omitempty"`
  LastOperatorFace     *bool      `json:"last_operator_face,omitempty"`
  ActivityMetrics      *bool      `json:"activity_metrics,omitempty"`
  AvailabilityTooltip  *bool      `json:"availability_tooltip,omitempty"`
  HideOnAway           *bool      `json:"hide_on_away,omitempty"`
  PositionReverse      *bool      `json:"position_reverse,omitempty"`
  EmailVisitors        *bool      `json:"email_visitors,omitempty"`
  ForceIdentify        *bool      `json:"force_identify,omitempty"`
  CheckDomain          *bool      `json:"check_domain,omitempty"`
  ColorTheme           *string    `json:"color_theme,omitempty"`
  TextTheme            *string    `json:"text_theme,omitempty"`
  WelcomeMessage       *string    `json:"welcome_message,omitempty"`
  Locale               *string    `json:"locale,omitempty"`
  BlockedPages         *[]string  `json:"blocked_pages,omitempty"`
  BlockedLocales       *[]string  `json:"blocked_locales,omitempty"`
}

// WebsiteSettingsUpdate mapping
type WebsiteSettingsUpdate struct {
  WebsiteID  string                        `json:"website_id,omitempty"`
  Name       string                        `json:"name,omitempty"`
  Domain     string                        `json:"domain,omitempty"`
  Logo       string                        `json:"logo,omitempty"`
  Contact    WebsiteSettingsUpdateContact  `json:"contact,omitempty"`
  Chatbox    WebsiteSettingsUpdateChatbox  `json:"chatbox,omitempty"`
}

// WebsiteSettingsUpdateContact mapping
type WebsiteSettingsUpdateContact struct {
  Email  string  `json:"email,omitempty"`
  Phone  string  `json:"phone,omitempty"`
}

// WebsiteSettingsUpdateChatbox mapping
type WebsiteSettingsUpdateChatbox struct {
  Tile                 string     `json:"tile,omitempty"`
  LastOperatorFace     bool       `json:"last_operator_face,omitempty"`
  ActivityMetrics      bool       `json:"activity_metrics,omitempty"`
  AvailabilityTooltip  bool       `json:"availability_tooltip,omitempty"`
  HideOnAway           bool       `json:"hide_on_away,omitempty"`
  PositionReverse      bool       `json:"position_reverse,omitempty"`
  EmailVisitors        bool       `json:"email_visitors,omitempty"`
  ForceIdentify        bool       `json:"force_identify,omitempty"`
  CheckDomain          bool       `json:"check_domain,omitempty"`
  ColorTheme           string     `json:"color_theme,omitempty"`
  TextTheme            string     `json:"text_theme,omitempty"`
  WelcomeMessage       string     `json:"welcome_message,omitempty"`
  Locale               string     `json:"locale,omitempty"`
  BlockedPages         *[]string  `json:"blocked_pages,omitempty"`
  BlockedLocales       *[]string  `json:"blocked_locales,omitempty"`
}


// String returns the string representation of WebsiteSettings
func (instance WebsiteSettings) String() string {
  return Stringify(instance)
}


// GetWebsiteSettings resolves the current settings for a website.
func (service *WebsiteService) GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error) {
  url := fmt.Sprintf("website/%s/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(WebsiteSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// UpdateWebsiteSettings updates the current settings for a website.
func (service *WebsiteService) UpdateWebsiteSettings(websiteID string, settings WebsiteSettingsUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}
