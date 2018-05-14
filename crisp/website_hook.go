// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteHookData mapping
type WebsiteHookData struct {
  Data  *WebsiteHook  `json:"data,omitempty"`
}

// WebsiteHookListData mapping
type WebsiteHookListData struct {
  Data  *[]WebsiteHook  `json:"data,omitempty"`
}

// WebsiteHookTagsData mapping
type WebsiteHookTagsData struct {
  Data  *[]string  `json:"data,omitempty"`
}

// WebsiteHook mapping
type WebsiteHook struct {
  WebsiteHookItem
  HookID  *string  `json:"hook_id,omitempty"`
  Status  *string  `json:"status,omitempty"`
}

// WebsiteHookItem mapping
type WebsiteHookItem struct {
  Label   *string    `json:"label,omitempty"`
  Target  *string    `json:"target,omitempty"`
  Events  *[]string  `json:"events,omitempty"`
}


// String returns the string representation of WebsiteHook
func (instance WebsiteHook) String() string {
  return Stringify(instance)
}


// ListHooks lists hooks for website.
func (service *WebsiteService) ListHooks(websiteID string, pageNumber uint) (*[]WebsiteHook, *Response, error) {
  url := fmt.Sprintf("website/%s/hooks/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  hooks := new(WebsiteHookListData)
  resp, err := service.client.Do(req, hooks)
  if err != nil {
    return nil, resp, err
  }

  return hooks.Data, resp, err
}


// CreateNewHook creates a new hook.
func (service *WebsiteService) CreateNewHook(websiteID string, websiteHookItem WebsiteHookItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/hook", websiteID)
  req, _ := service.client.NewRequest("POST", url, websiteHookItem)

  return service.client.Do(req, nil)
}


// CheckHookExists checks if given hook exists.
func (service *WebsiteService) CheckHookExists(websiteID string, hookID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/hook/%s", websiteID, hookID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// GetHook resolves hook information.
func (service *WebsiteService) GetHook(websiteID string, hookID string) (*WebsiteHook, *Response, error) {
  url := fmt.Sprintf("website/%s/hook/%s", websiteID, hookID)
  req, _ := service.client.NewRequest("GET", url, nil)

  hook := new(WebsiteHookData)
  resp, err := service.client.Do(req, hook)
  if err != nil {
    return nil, resp, err
  }

  return hook.Data, resp, err
}


// SaveHook saves a hook in website.
func (service *WebsiteService) SaveHook(websiteID string, hookID string, websiteHookItem WebsiteHookItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/hook/%s", websiteID, hookID)
  req, _ := service.client.NewRequest("PUT", url, websiteHookItem)

  return service.client.Do(req, nil)
}


// UpdateHook updates a hook in website.
func (service *WebsiteService) UpdateHook(websiteID string, hookID string, websiteHookItem WebsiteHookItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/hook/%s", websiteID, hookID)
  req, _ := service.client.NewRequest("PATCH", url, websiteHookItem)

  return service.client.Do(req, nil)
}


// RemoveHook removes a hook in website.
func (service *WebsiteService) RemoveHook(websiteID string, hookID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/hook/%s", websiteID, hookID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
