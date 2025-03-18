// Copyright 2018 Crisp IM SAS All rights reserved.
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
  Audit      *WebsiteSettingsAudit    `json:"audit,omitempty"`
  Contact    *WebsiteSettingsContact  `json:"contact,omitempty"`
  Inbox      *WebsiteSettingsInbox    `json:"inbox,omitempty"`
  Emails     *WebsiteSettingsEmails   `json:"emails,omitempty"`
  Chatbox    *WebsiteSettingsChatbox  `json:"chatbox,omitempty"`
}

// WebsiteSettingsAudit mapping
type WebsiteSettingsAudit struct {
  Log  *bool  `json:"log,omitempty"`
}

// WebsiteSettingsContact mapping
type WebsiteSettingsContact struct {
  Email      *string  `json:"email,omitempty"`
  Phone      *string  `json:"phone,omitempty"`
  Messenger  *string  `json:"messenger,omitempty"`
  Telegram   *string  `json:"telegram,omitempty"`
  Twitter    *string  `json:"twitter,omitempty"`
  WhatsApp   *string  `json:"whatsapp,omitempty"`
  Instagram  *string  `json:"instagram,omitempty"`
}

// WebsiteSettingsInbox mapping
type WebsiteSettingsInbox struct {
  LockRemoval         *bool    `json:"lock_removal,omitempty"`
  ForceOperatorToken  *bool    `json:"force_operator_token,omitempty"`
  Locale              *string  `json:"locale,omitempty"`
}

// WebsiteSettingsEmails mapping
type WebsiteSettingsEmails struct {
  Rating      *bool  `json:"rating,omitempty"`
  Transcript  *bool  `json:"transcript,omitempty"`
  Enrich      *bool  `json:"enrich,omitempty"`
  JunkFilter  *bool  `json:"junk_filter,omitempty"`
}

// WebsiteSettingsChatbox mapping
type WebsiteSettingsChatbox struct {
  Tile                 *string    `json:"tile,omitempty"`
  WaitGame             *bool      `json:"wait_game,omitempty"`
  WebsiteLogo          *bool      `json:"website_logo,omitempty"`
  LastOperatorFace     *bool      `json:"last_operator_face,omitempty"`
  OngoingOperatorFace  *bool      `json:"ongoing_operator_face,omitempty"`
  ActivityMetrics      *bool      `json:"activity_metrics,omitempty"`
  OperatorPrivacy      *bool      `json:"operator_privacy,omitempty"`
  VisitorPrivacy       *bool      `json:"visitor_privacy,omitempty"`
  AvailabilityTooltip  *bool      `json:"availability_tooltip,omitempty"`
  HideVacation         *bool      `json:"hide_vacation,omitempty"`
  HideOnAway           *bool      `json:"hide_on_away,omitempty"`
  HideOnMobile         *bool      `json:"hide_on_mobile,omitempty"`
  PositionReverse      *bool      `json:"position_reverse,omitempty"`
  EmailVisitors        *bool      `json:"email_visitors,omitempty"`
  PhoneVisitors        *bool      `json:"phone_visitors,omitempty"`
  ForceIdentify        *bool      `json:"force_identify,omitempty"`
  IgnorePrivacy        *bool      `json:"ignore_privacy,omitempty"`
  VisitorCompose       *bool      `json:"visitor_compose,omitempty"`
  FileTransfer         *bool      `json:"file_transfer,omitempty"`
  AudioRecord          *bool      `json:"audio_record,omitempty"`
  OverlaySearch        *bool      `json:"overlay_search,omitempty"`
  OverlayMode          *bool      `json:"overlay_mode,omitempty"`
  HelpdeskLink         *bool      `json:"helpdesk_link,omitempty"`
  HelpdeskOnly         *bool      `json:"helpdesk_only,omitempty"`
  StatusHealthDead     *bool      `json:"status_health_dead,omitempty"`
  CheckDomain          *bool      `json:"check_domain,omitempty"`
  ColorTheme           *string    `json:"color_theme,omitempty"`
  TextTheme            *string    `json:"text_theme,omitempty"`
  WelcomeMessage       *string    `json:"welcome_message,omitempty"`
  Locale               *string    `json:"locale,omitempty"`
  AllowedPages         *[]string  `json:"allowed_pages,omitempty"`
  BlockedPages         *[]string  `json:"blocked_pages,omitempty"`
  BlockedCountries     *[]string  `json:"blocked_countries,omitempty"`
  BlockedLocales       *[]string  `json:"blocked_locales,omitempty"`
}

// WebsiteSettingsUpdate mapping
type WebsiteSettingsUpdate struct {
  WebsiteID  string                         `json:"website_id,omitempty"`
  Name       string                         `json:"name,omitempty"`
  Domain     string                         `json:"domain,omitempty"`
  Logo       string                         `json:"logo,omitempty"`
  Audit      *WebsiteSettingsUpdateAudit    `json:"audit,omitempty"`
  Contact    *WebsiteSettingsUpdateContact  `json:"contact,omitempty"`
  Inbox      *WebsiteSettingsUpdateInbox    `json:"inbox,omitempty"`
  Emails     *WebsiteSettingsUpdateEmails   `json:"emails,omitempty"`
  Chatbox    *WebsiteSettingsUpdateChatbox  `json:"chatbox,omitempty"`
}

// WebsiteSettingsUpdateAudit mapping
type WebsiteSettingsUpdateAudit struct {
  Log  bool  `json:"log,omitempty"`
}

// WebsiteSettingsUpdateContact mapping
type WebsiteSettingsUpdateContact struct {
  Email      string  `json:"email,omitempty"`
  Phone      string  `json:"phone,omitempty"`
  Messenger  string  `json:"messenger,omitempty"`
  Telegram   string  `json:"telegram,omitempty"`
  Twitter    string  `json:"twitter,omitempty"`
  WhatsApp   string  `json:"whatsapp,omitempty"`
  Instagram  string  `json:"instagram,omitempty"`
}

// WebsiteSettingsUpdateInbox mapping
type WebsiteSettingsUpdateInbox struct {
  LockRemoval         bool    `json:"lock_removal,omitempty"`
  ForceOperatorToken  bool    `json:"force_operator_token,omitempty"`
  Locale              string  `json:"locale,omitempty"`
}

// WebsiteSettingsUpdateEmails mapping
type WebsiteSettingsUpdateEmails struct {
  Rating      bool  `json:"rating,omitempty"`
  Transcript  bool  `json:"transcript,omitempty"`
  Enrich      bool  `json:"enrich,omitempty"`
  JunkFilter  bool  `json:"junk_filter,omitempty"`
}

// WebsiteSettingsUpdateChatbox mapping
type WebsiteSettingsUpdateChatbox struct {
  Tile                 string     `json:"tile,omitempty"`
  WaitGame             bool       `json:"wait_game,omitempty"`
  WebsiteLogo          bool       `json:"website_logo,omitempty"`
  LastOperatorFace     bool       `json:"last_operator_face,omitempty"`
  OngoingOperatorFace  bool       `json:"ongoing_operator_face,omitempty"`
  ActivityMetrics      bool       `json:"activity_metrics,omitempty"`
  OperatorPrivacy      bool       `json:"operator_privacy,omitempty"`
  VisitorPrivacy       bool       `json:"visitor_privacy,omitempty"`
  AvailabilityTooltip  bool       `json:"availability_tooltip,omitempty"`
  HideVacation         bool       `json:"hide_vacation,omitempty"`
  HideOnAway           bool       `json:"hide_on_away,omitempty"`
  HideOnMobile         bool       `json:"hide_on_mobile,omitempty"`
  PositionReverse      bool       `json:"position_reverse,omitempty"`
  EmailVisitors        bool       `json:"email_visitors,omitempty"`
  PhoneVisitors        bool       `json:"phone_visitors,omitempty"`
  ForceIdentify        bool       `json:"force_identify,omitempty"`
  IgnorePrivacy        bool       `json:"ignore_privacy,omitempty"`
  VisitorCompose       bool       `json:"visitor_compose,omitempty"`
  FileTransfer         bool       `json:"file_transfer,omitempty"`
  AudioRecord          bool       `json:"audio_record,omitempty"`
  OverlaySearch        bool       `json:"overlay_search,omitempty"`
  OverlayMode          bool       `json:"overlay_mode,omitempty"`
  HelpdeskLink         bool       `json:"helpdesk_link,omitempty"`
  HelpdeskOnly         bool       `json:"helpdesk_only,omitempty"`
  StatusHealthDead     bool       `json:"status_health_dead,omitempty"`
  CheckDomain          bool       `json:"check_domain,omitempty"`
  ColorTheme           string     `json:"color_theme,omitempty"`
  TextTheme            string     `json:"text_theme,omitempty"`
  WelcomeMessage       string     `json:"welcome_message,omitempty"`
  Locale               string     `json:"locale,omitempty"`
  AllowedPages         *[]string  `json:"allowed_pages,omitempty"`
  BlockedPages         *[]string  `json:"blocked_pages,omitempty"`
  BlockedCountries     *[]string  `json:"blocked_countries,omitempty"`
  BlockedLocales       *[]string  `json:"blocked_locales,omitempty"`
  BlockedIPs           *[]string  `json:"blocked_ips,omitempty"`
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
