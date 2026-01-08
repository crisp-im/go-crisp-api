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
  Data  *[]WebsiteInbox  `json:"data,omitempty"`
}

// WebsiteInboxData mapping
type WebsiteInboxData struct {
  Data  *WebsiteInbox  `json:"data,omitempty"`
}

// WebsiteInboxNewData mapping
type WebsiteInboxNewData struct {
  Data  *WebsiteInboxNew  `json:"data,omitempty"`
}

// WebsiteInbox mapping
type WebsiteInbox struct {
  WebsiteInboxItem
  InboxID    *string  `json:"inbox_id,omitempty"`
  CreatedAt  *uint64  `json:"created_at,omitempty"`
  UpdatedAt  *uint64  `json:"updated_at,omitempty"`
}

// WebsiteInboxNew mapping
type WebsiteInboxNew struct {
  InboxID  *string  `json:"inbox_id,omitempty"`
}

// WebsiteInboxItem mapping
type WebsiteInboxItem struct {
  Name        *string                       `json:"name,omitempty"`
  Emoji       *string                       `json:"emoji,omitempty"`
  Order       *uint16                       `json:"order,omitempty"`
  Operators   *[]string                     `json:"operators,omitempty"`
  Conditions  *[]WebsiteInboxItemCondition  `json:"conditions,omitempty"`
}

// WebsiteInboxItemCondition mapping
type WebsiteInboxItemCondition struct {
  Model      *string    `json:"model,omitempty"`
  Criterion  *string    `json:"criterion,omitempty"`
  Operator   *string    `json:"operator,omitempty"`
  Query      *[]string  `json:"query,omitempty"`
}

// WebsiteInboxOrdersItem mapping
type WebsiteInboxOrdersItem struct {
  Orders  WebsiteInboxOrdersItem  `json:"orders"`
}

// WebsiteInboxOrderItem mapping
type WebsiteInboxOrderItem struct {
  InboxID  string  `json:"inbox_id"`
  Order    uint16  `json:"order"`
}


// String returns the string representation of WebsiteInbox
func (instance WebsiteInbox) String() string {
  return Stringify(instance)
}


// String returns the string representation of WebsiteInboxNew
func (instance WebsiteInboxNew) String() string {
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


// BatchOrderInboxes changes order for multiple matching inboxes on website.
func (service *WebsiteService) BatchOrderInboxes(websiteID string, orders WebsiteInboxOrdersItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/inboxes/batch/order", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, orders)

  return service.client.Do(req, nil)
}


// CreateNewInbox creates a new inbox for website.
func (service *WebsiteService) CreateNewInbox(websiteID string, inbox WebsiteInboxItem) (*WebsiteInboxNew, *Response, error) {
  url := fmt.Sprintf("website/%s/inbox", websiteID)
  req, _ := service.client.NewRequest("POST", url, inbox)

  inboxNew := new(WebsiteInboxNewData)
  resp, err := service.client.Do(req, inboxNew)
  if err != nil {
    return nil, resp, err
  }

  return inboxNew.Data, resp, err
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


// SaveInbox saves inbox for website.
func (service *WebsiteService) SaveInbox(websiteID string, inboxID string, inbox WebsiteInboxItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/inbox/%s", websiteID, inboxID)
  req, _ := service.client.NewRequest("PUT", url, inbox)

  return service.client.Do(req, nil)
}


// DeleteInbox deletes inbox for website.
func (service *WebsiteService) DeleteInbox(websiteID string, inboxID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/inbox/%s", websiteID, inboxID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
