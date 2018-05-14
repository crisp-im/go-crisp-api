// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteCampaignNewData mapping
type WebsiteCampaignNewData struct {
  Data  *WebsiteCampaignNew  `json:"data,omitempty"`
}

// WebsiteCampaignNew mapping
type WebsiteCampaignNew struct {
  CampaignID  *string  `json:"campaign_id,omitempty"`
}

// WebsiteCampaignNewItem mapping
type WebsiteCampaignNewItem struct {
  Type  string  `json:"type"`
  Name  string  `json:"name"`
}

// WebsiteCampaignExcerptsData mapping
type WebsiteCampaignExcerptsData struct {
  Data  *[]WebsiteCampaignExcerpt  `json:"data,omitempty"`
}

// WebsiteCampaignExcerpt mapping
type WebsiteCampaignExcerpt struct {
  CampaignID    *string  `json:"campaign_id,omitempty"`
  Type          *string  `json:"type,omitempty"`
  Name          *string  `json:"name,omitempty"`
  Ready         *bool    `json:"ready,omitempty"`
  Dispatched    *bool    `json:"dispatched,omitempty"`
  Running       *bool    `json:"running,omitempty"`
  Progress      *uint8   `json:"progress,omitempty"`
  Targets       *uint32  `json:"targets,omitempty"`
  Reached       *uint32  `json:"reached,omitempty"`
  CreatedAt     *uint    `json:"created_at,omitempty"`
  UpdatedAt     *uint    `json:"updated_at,omitempty"`
  DispatchedAt  *uint    `json:"dispatched_at,omitempty"`
}

// WebsiteCampaignItemData mapping
type WebsiteCampaignItemData struct {
  Data  *WebsiteCampaignItem  `json:"data,omitempty"`
}

// WebsiteCampaignItem mapping
type WebsiteCampaignItem struct {
  WebsiteCampaignExcerpt
  Sender      *WebsiteCampaignItemSender      `json:"sender,omitempty"`
  Recipients  *WebsiteCampaignItemRecipients  `json:"recipients,omitempty"`
  Flow        *WebsiteCampaignItemFlow        `json:"flow,omitempty"`
  Message     *string                         `json:"message,omitempty"`
  Options     *WebsiteCampaignItemOptions     `json:"options,omitempty"`
  Statistics  *WebsiteCampaignItemStatistics  `json:"statistics,omitempty"`
}

// WebsiteCampaignItemSender mapping
type WebsiteCampaignItemSender struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// WebsiteCampaignItemRecipients mapping
type WebsiteCampaignItemRecipients struct {
  Type      *string        `json:"type,omitempty"`
  Segments  *[]string      `json:"segments,omitempty"`
  People    *[]string      `json:"people,omitempty"`
  Filter    *[]UserFilter  `json:"filter,omitempty"`
}

// WebsiteCampaignItemFlow mapping
type WebsiteCampaignItemFlow struct {
  LaunchEvent   *string        `json:"launch_event,omitempty"`
  AssertFilter  *[]UserFilter  `json:"assert_filter,omitempty"`
  AssertDelay   *uint16        `json:"assert_delay,omitempty"`
  DeliverOnce   *bool          `json:"deliver_once,omitempty"`
  DeliverDelay  *uint16        `json:"deliver_delay,omitempty"`
}

// WebsiteCampaignItemOptions mapping
type WebsiteCampaignItemOptions struct {
  DeliverToChatbox  *bool  `json:"deliver_to_chatbox,omitempty"`
  DeliverToEmail    *bool  `json:"deliver_to_email,omitempty"`
}

// WebsiteCampaignItemStatistics mapping
type WebsiteCampaignItemStatistics struct {
  Opened   *uint  `json:"opened,omitempty"`
  Clicked  *uint  `json:"clicked,omitempty"`
}

// WebsiteCampaignStatisticsData mapping
type WebsiteCampaignStatisticsData struct {
  Data  *[]WebsiteCampaignStatistic  `json:"data,omitempty"`
}

// WebsiteCampaignStatistic mapping
type WebsiteCampaignStatistic struct {
  Profile    *WebsiteCampaignStatisticProfile  `json:"profile,omitempty"`
  Data       *interface{}                      `json:"data,omitempty"`
  CreatedAt  *uint                             `json:"created_at,omitempty"`
  UpdatedAt  *uint                             `json:"updated_at,omitempty"`
}

// WebsiteCampaignStatisticProfile mapping
type WebsiteCampaignStatisticProfile struct {
  PeopleID  *string                                 `json:"people_id,omitempty"`
  Email     *string                                 `json:"email,omitempty"`
  Person    *WebsiteCampaignStatisticProfilePerson  `json:"person,omitempty"`
}

// WebsiteCampaignStatisticProfilePerson mapping
type WebsiteCampaignStatisticProfilePerson struct {
  Nickname     *string                                            `json:"nickname,omitempty"`
  Avatar       *string                                            `json:"avatar,omitempty"`
  Geolocation  *WebsiteCampaignStatisticProfilePersonGeolocation  `json:"geolocation,omitempty"`
}

// WebsiteCampaignStatisticProfilePersonGeolocation mapping
type WebsiteCampaignStatisticProfilePersonGeolocation struct {
  Country      *string                                                       `json:"country,omitempty"`
  Region       *string                                                       `json:"region,omitempty"`
  City         *string                                                       `json:"city,omitempty"`
  Coordinates  *WebsiteCampaignStatisticProfilePersonGeolocationCoordinates  `json:"coordinates,omitempty"`
}

// WebsiteCampaignStatisticProfilePersonGeolocationCoordinates mapping
type WebsiteCampaignStatisticProfilePersonGeolocationCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}


// String returns the string representation of WebsiteCampaignNew
func (instance WebsiteCampaignNew) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteCampaignExcerpt
func (instance WebsiteCampaignExcerpt) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteCampaignItem
func (instance WebsiteCampaignItem) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteCampaignStatistic
func (instance WebsiteCampaignStatistic) String() string {
  return Stringify(instance)
}


// ListCampaigns lists campaigns for website.
func (service *WebsiteService) ListCampaigns(websiteID string, pageNumber uint) (*[]WebsiteCampaignExcerpt, *Response, error) {
  url := fmt.Sprintf("website/%s/campaigns/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  campaigns := new(WebsiteCampaignExcerptsData)
  resp, err := service.client.Do(req, campaigns)
  if err != nil {
    return nil, resp, err
  }

  return campaigns.Data, resp, err
}


// FilterCampaigns lists campaigns for website (filter variant).
func (service *WebsiteService) FilterCampaigns(websiteID string, pageNumber uint, searchName string, filterTypeOneShot bool, filterTypeAutomated bool, filterStatusNotConfigured bool, filterStatusReady bool, filterStatusPaused bool, filterStatusSending bool, filterStatusDone bool) (*[]WebsiteCampaignExcerpt, *Response, error) {
  var (
    filterTypeOneShotValue string
    filterTypeAutomatedValue string
    filterStatusNotConfiguredValue string
    filterStatusReadyValue string
    filterStatusPausedValue string
    filterStatusSendingValue string
    filterStatusDoneValue string
  )

  if filterTypeOneShot == true {
    filterTypeOneShotValue = "1"
  } else {
    filterTypeOneShotValue = "0"
  }

  if filterTypeAutomated == true {
    filterTypeAutomatedValue = "1"
  } else {
    filterTypeAutomatedValue = "0"
  }

  if filterStatusNotConfigured == true {
    filterStatusNotConfiguredValue = "1"
  } else {
    filterStatusNotConfiguredValue = "0"
  }

  if filterStatusReady == true {
    filterStatusReadyValue = "1"
  } else {
    filterStatusReadyValue = "0"
  }

  if filterStatusPaused == true {
    filterStatusPausedValue = "1"
  } else {
    filterStatusPausedValue = "0"
  }

  if filterStatusSending == true {
    filterStatusSendingValue = "1"
  } else {
    filterStatusSendingValue = "0"
  }

  if filterStatusDone == true {
    filterStatusDoneValue = "1"
  } else {
    filterStatusDoneValue = "0"
  }

  url := fmt.Sprintf("website/%s/campaigns/%d?search_name=%s&filter_type_one_shot=%s&filter_type_automated=%s&filter_status_not_configured=%s&filter_status_ready=%s&filter_status_paused=%s&filter_status_sending=%s&filter_status_done=%s", websiteID, pageNumber, url.QueryEscape(searchName), url.QueryEscape(filterTypeOneShotValue), url.QueryEscape(filterTypeAutomatedValue), url.QueryEscape(filterStatusNotConfiguredValue), url.QueryEscape(filterStatusReadyValue), url.QueryEscape(filterStatusPausedValue), url.QueryEscape(filterStatusSendingValue), url.QueryEscape(filterStatusDoneValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  campaigns := new(WebsiteCampaignExcerptsData)
  resp, err := service.client.Do(req, campaigns)
  if err != nil {
    return nil, resp, err
  }

  return campaigns.Data, resp, err
}


// CreateNewCampaign creates a new campaign.
func (service *WebsiteService) CreateNewCampaign(websiteID string, campaignType string, campaignName string) (*WebsiteCampaignNew, *Response, error) {
  url := fmt.Sprintf("website/%s/campaign", websiteID)
  req, _ := service.client.NewRequest("POST", url, WebsiteCampaignNewItem{Type: campaignType, Name: campaignName})

  campaign := new(WebsiteCampaignNewData)
  resp, err := service.client.Do(req, campaign)
  if err != nil {
    return nil, resp, err
  }

  return campaign.Data, resp, err
}


// CheckCampaignExists checks if given campaign exists.
func (service *WebsiteService) CheckCampaignExists(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s", websiteID, campaignID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// GetCampaign resolves campaign information.
func (service *WebsiteService) GetCampaign(websiteID string, campaignID string) (*WebsiteCampaignItem, *Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s", websiteID, campaignID)
  req, _ := service.client.NewRequest("GET", url, nil)

  campaign := new(WebsiteCampaignItemData)
  resp, err := service.client.Do(req, campaign)
  if err != nil {
    return nil, resp, err
  }

  return campaign.Data, resp, err
}


// SaveCampaign saves a campaign in website, and overwrite previous campaign information.
func (service *WebsiteService) SaveCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s", websiteID, campaignID)
  req, _ := service.client.NewRequest("PUT", url, websiteCampaignItem)

  return service.client.Do(req, nil)
}


// UpdateCampaign updates a campaign in website, and save only changed fields.
func (service *WebsiteService) UpdateCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s", websiteID, campaignID)
  req, _ := service.client.NewRequest("PATCH", url, websiteCampaignItem)

  return service.client.Do(req, nil)
}


// RemoveCampaign removes a campaign in website.
func (service *WebsiteService) RemoveCampaign(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s", websiteID, campaignID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// DispatchCampaign dispatches a ready campaign.
func (service *WebsiteService) DispatchCampaign(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/dispatch", websiteID, campaignID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ResumeCampaign resumes a dispatched campaign.
func (service *WebsiteService) ResumeCampaign(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/resume", websiteID, campaignID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// PauseCampaign pauses a dispatched campaign.
func (service *WebsiteService) PauseCampaign(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/pause", websiteID, campaignID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// TestCampaign tests a ready campaign.
func (service *WebsiteService) TestCampaign(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/test", websiteID, campaignID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ListCampaignStatistics lists campaigns statistics on action for website.
func (service *WebsiteService) ListCampaignStatistics(websiteID string, campaignID string, action string, pageNumber uint) (*[]WebsiteCampaignStatistic, *Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/statistics/%s/%d", websiteID, campaignID, action, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteCampaignStatisticsData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}
