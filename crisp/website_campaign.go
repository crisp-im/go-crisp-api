// Copyright 2016 Crisp IM, Inc. All rights reserved.
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
  Name  string  `json:"name"`
}

// WebsiteCampaignDispatch mapping
type WebsiteCampaignDispatch struct {
  Type    string    `json:"type,omitempty"`
  Emails  []string  `json:"emails,omitempty"`
}

// WebsiteCampaignExcerptsData mapping
type WebsiteCampaignExcerptsData struct {
  Data  *[]WebsiteCampaignExcerpt  `json:"data,omitempty"`
}

// WebsiteCampaignExcerpt mapping
type WebsiteCampaignExcerpt struct {
  CampaignID    *string  `json:"campaign_id,omitempty"`
  Name          *string  `json:"name,omitempty"`
  Ready         *bool    `json:"ready,omitempty"`
  Dispatched    *bool    `json:"dispatched,omitempty"`
  Running       *bool    `json:"running,omitempty"`
  Progress      *uint8   `json:"progress,omitempty"`
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
  Sender      *string  `json:"sender,omitempty"`
  Recipients  *string  `json:"recipients,omitempty"`
  Message     *string  `json:"message,omitempty"`
  Options     *string  `json:"options,omitempty"`
}

// WebsiteCampaignItemSender mapping
type WebsiteCampaignItemSender struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// WebsiteCampaignItemRecipients mapping
type WebsiteCampaignItemRecipients struct {
  Type      *string    `json:"type,omitempty"`
  Segments  *[]string  `json:"segments,omitempty"`
  People    *[]string  `json:"people,omitempty"`
}

// WebsiteCampaignItemOptions mapping
type WebsiteCampaignItemOptions struct {
  DeliverToChatbox  *bool  `json:"deliver_to_chatbox,omitempty"`
  DeliverToEmail    *bool  `json:"deliver_to_email,omitempty"`
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
func (service *WebsiteService) FilterCampaigns(websiteID string, pageNumber uint, filterName string, filterReady bool, filterDispatched bool, filterRunning bool) (*[]WebsiteCampaignExcerpt, *Response, error) {
  var (
    filterReadyValue string
    filterDispatchedValue string
    filterRunningValue string
  )

  if filterReady == true {
    filterReadyValue = "1"
  } else {
    filterReadyValue = "0"
  }

  if filterDispatched == true {
    filterDispatchedValue = "1"
  } else {
    filterDispatchedValue = "0"
  }

  if filterRunning == true {
    filterRunningValue = "1"
  } else {
    filterRunningValue = "0"
  }

  url := fmt.Sprintf("website/%s/campaigns/%d?filter_name=%s&filter_ready=%s&filter_dispatched=%s&filter_running=%s", websiteID, pageNumber, url.QueryEscape(filterName), url.QueryEscape(filterReadyValue), url.QueryEscape(filterDispatchedValue), url.QueryEscape(filterRunningValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  campaigns := new(WebsiteCampaignExcerptsData)
  resp, err := service.client.Do(req, campaigns)
  if err != nil {
    return nil, resp, err
  }

  return campaigns.Data, resp, err
}


// CreateNewCampaign creates a new campaign.
func (service *WebsiteService) CreateNewCampaign(websiteID string, name string) (*WebsiteCampaignNew, *Response, error) {
  url := fmt.Sprintf("website/%s/campaign", websiteID)
  req, _ := service.client.NewRequest("POST", url, WebsiteCampaignNewItem{Name: name})

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


// DispatchCampaignLive dispatches a ready campaign (live variant).
func (service *WebsiteService) DispatchCampaignLive(websiteID string, campaignID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/dispatch", websiteID, campaignID)
  req, _ := service.client.NewRequest("POST", url, WebsiteCampaignDispatch{Type: "live"})

  return service.client.Do(req, nil)
}


// DispatchCampaignTest dispatches a ready campaign (test variant).
func (service *WebsiteService) DispatchCampaignTest(websiteID string, campaignID string, emails []string) (*Response, error) {
  url := fmt.Sprintf("website/%s/campaign/%s/dispatch", websiteID, campaignID)
  req, _ := service.client.NewRequest("POST", url, WebsiteCampaignDispatch{Type: "test", Emails: emails})

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
