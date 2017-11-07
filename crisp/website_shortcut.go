// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteShortcutData mapping
type WebsiteShortcutData struct {
  Data  *WebsiteShortcut  `json:"data,omitempty"`
}

// WebsiteShortcutListData mapping
type WebsiteShortcutListData struct {
  Data  *[]WebsiteShortcut  `json:"data,omitempty"`
}

// WebsiteShortcutTagsData mapping
type WebsiteShortcutTagsData struct {
  Data  *[]string  `json:"data,omitempty"`
}

// WebsiteShortcut mapping
type WebsiteShortcut struct {
  WebsiteShortcutItem
  ShortcutID  *string  `json:"shortcut_id,omitempty"`
}

// WebsiteShortcutItem mapping
type WebsiteShortcutItem struct {
  Bang      *string  `json:"bang,omitempty"`
  Text      *string  `json:"text,omitempty"`
  Tag       *string  `json:"tag,omitempty"`
  Disabled  *bool    `json:"disabled,omitempty"`
}


// String returns the string representation of WebsiteShortcut
func (instance WebsiteShortcut) String() string {
  return Stringify(instance)
}


// SearchShortcuts lists shortcuts for website (search variant).
func (service *WebsiteService) SearchShortcuts(websiteID string, pageNumber uint, searchQuery string, searchTag string) (*[]WebsiteShortcut, *Response, error) {
  var resourceURL string

  if searchQuery != "" && searchTag != "" {
    resourceURL = fmt.Sprintf("website/%s/shortcuts/list/%d?search_query=%s&search_tag=%s", websiteID, pageNumber, url.QueryEscape(searchQuery), url.QueryEscape(searchTag))
  } else {
    resourceURL = fmt.Sprintf("website/%s/shortcuts/list/%d", websiteID, pageNumber)
  }

  req, _ := service.client.NewRequest("GET", resourceURL, nil)

  shortcuts := new(WebsiteShortcutListData)
  resp, err := service.client.Do(req, shortcuts)
  if err != nil {
    return nil, resp, err
  }

  return shortcuts.Data, resp, err
}


// ListShortcuts lists shortcuts for website.
func (service *WebsiteService) ListShortcuts(websiteID string, pageNumber uint) (*[]WebsiteShortcut, *Response, error) {
  return service.SearchShortcuts(websiteID, pageNumber, "", "")
}


// ListShortcutTags lists shortcut tags for website.
func (service *WebsiteService) ListShortcutTags(websiteID string) (*[]string, *Response, error) {
  url := fmt.Sprintf("website/%s/shortcuts/tags", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  tags := new(WebsiteShortcutTagsData)
  resp, err := service.client.Do(req, tags)
  if err != nil {
    return nil, resp, err
  }

  return tags.Data, resp, err
}


// CreateNewShortcut creates a new shortcut.
func (service *WebsiteService) CreateNewShortcut(websiteID string, websiteShortcutItem WebsiteShortcutItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/shortcut", websiteID)
  req, _ := service.client.NewRequest("POST", url, websiteShortcutItem)

  return service.client.Do(req, nil)
}


// CheckShortcutExists checks if given shortcut exists.
func (service *WebsiteService) CheckShortcutExists(websiteID string, shortcutID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/shortcut/%s", websiteID, shortcutID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// GetShortcut resolves shortcut information.
func (service *WebsiteService) GetShortcut(websiteID string, shortcutID string) (*WebsiteShortcut, *Response, error) {
  url := fmt.Sprintf("website/%s/shortcut/%s", websiteID, shortcutID)
  req, _ := service.client.NewRequest("GET", url, nil)

  shortcut := new(WebsiteShortcutData)
  resp, err := service.client.Do(req, shortcut)
  if err != nil {
    return nil, resp, err
  }

  return shortcut.Data, resp, err
}


// SaveShortcut saves a shortcut in website.
func (service *WebsiteService) SaveShortcut(websiteID string, shortcutID string, websiteShortcutItem WebsiteShortcutItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/shortcut/%s", websiteID, shortcutID)
  req, _ := service.client.NewRequest("PUT", url, websiteShortcutItem)

  return service.client.Do(req, nil)
}


// UpdateShortcut updates a shortcut in website.
func (service *WebsiteService) UpdateShortcut(websiteID string, shortcutID string, websiteShortcutItem WebsiteShortcutItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/shortcut/%s", websiteID, shortcutID)
  req, _ := service.client.NewRequest("PATCH", url, websiteShortcutItem)

  return service.client.Do(req, nil)
}


// RemoveShortcut removes a shortcut in website.
func (service *WebsiteService) RemoveShortcut(websiteID string, shortcutID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/shortcut/%s", websiteID, shortcutID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
