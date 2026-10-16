// Copyright 2026 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
)

// WebsiteInboxListData mapping
type WebsiteInboxListData struct {
	Data *[]WebsiteInbox `json:"data,omitempty"`
}

// WebsiteInboxData mapping
type WebsiteInboxData struct {
	Data *WebsiteInbox `json:"data,omitempty"`
}

// WebsiteInbox mapping
type WebsiteInbox struct {
	WebsiteInboxItem
	InboxID   *string `json:"inbox_id,omitempty"`
	CreatedAt *uint64 `json:"created_at,omitempty"`
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// WebsiteInboxItem mapping
type WebsiteInboxItem struct {
	Name          *string                      `json:"name,omitempty"`
	Emoji         *string                      `json:"emoji,omitempty"`
	Order         *uint16                      `json:"order,omitempty"`
	Operators     *[]string                    `json:"operators,omitempty"`
	Teams         *[]string                    `json:"teams,omitempty"`
	TeamOperators *[]string                    `json:"team_operators,omitempty"`
	Operator      *string                      `json:"operator,omitempty"`
	Conditions    *[]WebsiteInboxItemCondition `json:"conditions,omitempty"`
}

// WebsiteInboxItemCondition mapping
type WebsiteInboxItemCondition struct {
	Model     *string   `json:"model,omitempty"`
	Criterion *string   `json:"criterion,omitempty"`
	Operator  *string   `json:"operator,omitempty"`
	Query     *[]string `json:"query,omitempty"`
}

// String returns the string representation of WebsiteInbox
func (instance WebsiteInbox) String() string {
	return Stringify(instance)
}

// ListInboxes lists all inboxes for website.
func (service *WebsiteService) ListInboxes(websiteID string, pageNumber uint) (*[]WebsiteInbox, *Response, error) {
	url := fmt.Sprintf("website/%s/inboxes/list/%d", websiteID, pageNumber)
	req, _ := service.client.NewRequest("GET", url, nil)

	inboxes := new(WebsiteInboxListData)
	resp, err := service.client.Do(req, inboxes)
	if err != nil {
		return nil, resp, err
	}

	return inboxes.Data, resp, err
}

// CheckInboxExists checks if inbox exists for website.
func (service *WebsiteService) CheckInboxExists(websiteID string, inboxID string) (*Response, error) {
	url := fmt.Sprintf("website/%s/inbox/%s", websiteID, inboxID)
	req, _ := service.client.NewRequest("HEAD", url, nil)

	return service.client.Do(req, nil)
}

// GetInbox resolves inbox for website.
func (service *WebsiteService) GetInbox(websiteID string, inboxID string) (*WebsiteInbox, *Response, error) {
	url := fmt.Sprintf("website/%s/inbox/%s", websiteID, inboxID)
	req, _ := service.client.NewRequest("GET", url, nil)

	inbox := new(WebsiteInboxData)
	resp, err := service.client.Do(req, inbox)
	if err != nil {
		return nil, resp, err
	}

	return inbox.Data, resp, err
}
