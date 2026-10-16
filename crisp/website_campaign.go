// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
	"net/url"
)

// WebsiteCampaignTemplateNewData mapping
type WebsiteCampaignTemplateNewData struct {
	Data *WebsiteCampaignTemplateNew `json:"data,omitempty"`
}

// WebsiteCampaignTemplateNew mapping
type WebsiteCampaignTemplateNew struct {
	TemplateID *string `json:"template_id,omitempty"`
}

// WebsiteCampaignTemplateNewItem mapping
type WebsiteCampaignTemplateNewItem struct {
	Format string `json:"format"`
	Name   string `json:"name"`
}

// WebsiteCampaignTemplateExcerptsData mapping
type WebsiteCampaignTemplateExcerptsData struct {
	Data *[]WebsiteCampaignTemplateExcerpt `json:"data,omitempty"`
}

// WebsiteCampaignTemplateExcerpt mapping
type WebsiteCampaignTemplateExcerpt struct {
	TemplateID *string `json:"template_id,omitempty"`
	Type       *string `json:"type,omitempty"`
	Name       *string `json:"name,omitempty"`
	Format     *string `json:"format,omitempty"`
	CreatedAt  *uint64 `json:"created_at,omitempty"`
	UpdatedAt  *uint64 `json:"updated_at,omitempty"`
}

// WebsiteCampaignTemplateItemData mapping
type WebsiteCampaignTemplateItemData struct {
	Data *WebsiteCampaignTemplateItem `json:"data,omitempty"`
}

// WebsiteCampaignTemplateItem mapping
type WebsiteCampaignTemplateItem struct {
	WebsiteCampaignTemplateExcerpt
	Content *string `json:"content,omitempty"`
}

// String returns the string representation of WebsiteCampaignTemplateNew
func (instance WebsiteCampaignTemplateNew) String() string {
	return Stringify(instance)
}

// String returns the string representation of WebsiteCampaignTemplateExcerpt
func (instance WebsiteCampaignTemplateExcerpt) String() string {
	return Stringify(instance)
}

// String returns the string representation of WebsiteCampaignTemplateItem
func (instance WebsiteCampaignTemplateItem) String() string {
	return Stringify(instance)
}

// ListCampaignTemplates lists campaign templates for website.
func (service *WebsiteService) ListCampaignTemplates(websiteID string, pageNumber uint) (*[]WebsiteCampaignTemplateExcerpt, *Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/templates/%d", websiteID, pageNumber)
	req, _ := service.client.NewRequest("GET", url, nil)

	templates := new(WebsiteCampaignTemplateExcerptsData)
	resp, err := service.client.Do(req, templates)
	if err != nil {
		return nil, resp, err
	}

	return templates.Data, resp, err
}

// FilterCampaignTemplates lists campaign templates for website (filter variant).
func (service *WebsiteService) FilterCampaignTemplates(websiteID string, pageNumber uint, searchName string, filterTypeStatic bool, filterTypeCustom bool) (*[]WebsiteCampaignTemplateExcerpt, *Response, error) {
	var (
		filterTypeStaticValue string
		filterTypeCustomValue string
	)

	if filterTypeStatic == true {
		filterTypeStaticValue = "1"
	} else {
		filterTypeStaticValue = "0"
	}

	if filterTypeCustom == true {
		filterTypeCustomValue = "1"
	} else {
		filterTypeCustomValue = "0"
	}

	url := fmt.Sprintf("website/%s/campaigns/templates/%d?search_name=%s&filter_type_static=%s&filter_type_custom=%s", websiteID, pageNumber, url.QueryEscape(searchName), url.QueryEscape(filterTypeStaticValue), url.QueryEscape(filterTypeCustomValue))
	req, _ := service.client.NewRequest("GET", url, nil)

	templates := new(WebsiteCampaignTemplateExcerptsData)
	resp, err := service.client.Do(req, templates)
	if err != nil {
		return nil, resp, err
	}

	return templates.Data, resp, err
}

// CreateNewCampaignTemplate creates a new campaign template.
func (service *WebsiteService) CreateNewCampaignTemplate(websiteID string, templateFormat string, templateName string) (*WebsiteCampaignTemplateNew, *Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/template", websiteID)
	req, _ := service.client.NewRequest("POST", url, WebsiteCampaignTemplateNewItem{Format: templateFormat, Name: templateName})

	template := new(WebsiteCampaignTemplateNewData)
	resp, err := service.client.Do(req, template)
	if err != nil {
		return nil, resp, err
	}

	return template.Data, resp, err
}

// CheckCampaignTemplateExists checks if given campaign template exists.
func (service *WebsiteService) CheckCampaignTemplateExists(websiteID string, templateID string) (*Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/template/%s", websiteID, templateID)
	req, _ := service.client.NewRequest("HEAD", url, nil)

	return service.client.Do(req, nil)
}

// GetCampaignTemplate resolves campaign template information.
func (service *WebsiteService) GetCampaignTemplate(websiteID string, templateID string) (*WebsiteCampaignTemplateItem, *Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/template/%s", websiteID, templateID)
	req, _ := service.client.NewRequest("GET", url, nil)

	template := new(WebsiteCampaignTemplateItemData)
	resp, err := service.client.Do(req, template)
	if err != nil {
		return nil, resp, err
	}

	return template.Data, resp, err
}

// SaveCampaignTemplate saves a campaign template in website, and overwrite previous template information.
func (service *WebsiteService) SaveCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/template/%s", websiteID, templateID)
	req, _ := service.client.NewRequest("PUT", url, websiteCampaignTemplateItem)

	return service.client.Do(req, nil)
}

// UpdateCampaignTemplate updates a campaign template in website, and save only changed fields.
func (service *WebsiteService) UpdateCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/template/%s", websiteID, templateID)
	req, _ := service.client.NewRequest("PATCH", url, websiteCampaignTemplateItem)

	return service.client.Do(req, nil)
}

// RemoveCampaignTemplate removes a campaign template in website.
func (service *WebsiteService) RemoveCampaignTemplate(websiteID string, templateID string) (*Response, error) {
	url := fmt.Sprintf("website/%s/campaigns/template/%s", websiteID, templateID)
	req, _ := service.client.NewRequest("DELETE", url, nil)

	return service.client.Do(req, nil)
}
