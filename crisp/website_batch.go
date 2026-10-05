// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
)

// WebsiteBatchConversationsOperation mapping
type WebsiteBatchConversationsOperation struct {
	InboxID  *string  `json:"inbox_id,omitempty"`
	Sessions []string `json:"sessions,omitempty"`
}

// WebsiteBatchPeopleOperation mapping
type WebsiteBatchPeopleOperation struct {
	People *WebsiteBatchPeopleOperationInner `json:"people,omitempty"`
}

// WebsiteBatchPeopleOperationInner mapping
type WebsiteBatchPeopleOperationInner struct {
	Profiles *[]string                               `json:"profiles,omitempty"`
	Search   *WebsiteBatchPeopleOperationInnerSearch `json:"search,omitempty"`
}

// WebsiteBatchPeopleOperationInnerSearch mapping
type WebsiteBatchPeopleOperationInnerSearch struct {
	Filter   []WebsiteFilter `json:"filter,omitempty"`
	Operator string          `json:"operator,omitempty"`
}

// WebsiteBatchReportOperation mapping
type WebsiteBatchReportOperation struct {
	Sessions []string `json:"sessions"`
	Flag     string   `json:"flag"`
}

// WebsiteBatchBlockOperation mapping
type WebsiteBatchBlockOperation struct {
	Sessions []string `json:"sessions"`
	Blocked  bool     `json:"blocked"`
}

// WebsiteBatchRoutingOperationAssigned mapping
type WebsiteBatchRoutingOperationAssigned struct {
	UserID string `json:"user_id"`
}

// WebsiteBatchRoutingOperation mapping
type WebsiteBatchRoutingOperation struct {
	Sessions []string                              `json:"sessions"`
	Assigned *WebsiteBatchRoutingOperationAssigned `json:"assigned"`
}

// WebsiteBatchInboxOperation mapping
type WebsiteBatchInboxOperation struct {
	Sessions []string `json:"sessions"`
	InboxID  *string  `json:"inbox_id"`
}

// WebsiteBatchDataPayload mapping
type WebsiteBatchDataPayload map[string]interface{}

// WebsiteBatchConversationsDataOperation mapping
type WebsiteBatchConversationsDataOperation struct {
	Sessions []string                `json:"sessions"`
	Data     WebsiteBatchDataPayload `json:"data"`
}

// WebsiteBatchPeopleDataOperation mapping
type WebsiteBatchPeopleDataOperation struct {
	People *WebsiteBatchPeopleOperationInner `json:"people,omitempty"`
	Data   WebsiteBatchDataPayload           `json:"data"`
}

// WebsiteBatchConversationsSegmentsOperation mapping
type WebsiteBatchConversationsSegmentsOperation struct {
	Sessions []string `json:"sessions"`
	Segments []string `json:"segments"`
}

// WebsiteBatchPeopleSegmentsOperation mapping
type WebsiteBatchPeopleSegmentsOperation struct {
	People   *WebsiteBatchPeopleOperationInner `json:"people,omitempty"`
	Segments []string                          `json:"segments"`
}

// BatchResolveConversations resolves given (or all) items in website (conversation variant).
func (service *WebsiteService) BatchResolveConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/resolve", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchUnresolveConversations unresolves given items in website (conversation variant).
func (service *WebsiteService) BatchUnresolveConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/unresolve", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchReadConversations marks given (or all) items as read in website (conversation variant).
func (service *WebsiteService) BatchReadConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/read", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchUnreadConversations marks given items as unread in website (conversation variant).
func (service *WebsiteService) BatchUnreadConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/unread", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchRemoveConversations removes given items in website (conversation variant).
func (service *WebsiteService) BatchRemoveConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/remove", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchRemovePeople removes given items in website (people variant).
func (service *WebsiteService) BatchRemovePeople(websiteID string, people WebsiteBatchPeopleOperationInner) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/remove", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchPeopleOperation{People: &people})

	return service.client.Do(req, nil)
}

// BatchReportConversations reports given items in website (conversation variant).
func (service *WebsiteService) BatchReportConversations(websiteID string, operation WebsiteBatchReportOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/report", websiteID)
	req, _ := service.client.NewRequest("POST", url, operation)

	return service.client.Do(req, nil)
}

// BatchBlockConversations blocks or unblocks given items in website (conversation variant).
func (service *WebsiteService) BatchBlockConversations(websiteID string, operation WebsiteBatchBlockOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/block", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchRoutingConversations assigns or unassigns operator for given items in website (conversation variant).
func (service *WebsiteService) BatchRoutingConversations(websiteID string, operation WebsiteBatchRoutingOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/routing", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchInboxConversations moves given items to a different inbox in website (conversation variant).
func (service *WebsiteService) BatchInboxConversations(websiteID string, operation WebsiteBatchInboxOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/inbox", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchUpdateConversationsData updates custom data on given items in website (conversation variant).
func (service *WebsiteService) BatchUpdateConversationsData(websiteID string, operation WebsiteBatchConversationsDataOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/data", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchUpdatePeopleData updates custom data on given items in website (people variant).
func (service *WebsiteService) BatchUpdatePeopleData(websiteID string, people WebsiteBatchPeopleOperationInner, data WebsiteBatchDataPayload) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/data", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchPeopleDataOperation{People: &people, Data: data})

	return service.client.Do(req, nil)
}

// BatchUpdateConversationsSegments updates segments on given items in website (conversation variant).
func (service *WebsiteService) BatchUpdateConversationsSegments(websiteID string, operation WebsiteBatchConversationsSegmentsOperation) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/segments", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, operation)

	return service.client.Do(req, nil)
}

// BatchUpdatePeopleSegments updates segments on given items in website (people variant).
func (service *WebsiteService) BatchUpdatePeopleSegments(websiteID string, people WebsiteBatchPeopleOperationInner, segments []string) (*Response, error) {
	url := fmt.Sprintf("website/%s/batch/segments", websiteID)
	req, _ := service.client.NewRequest("PATCH", url, WebsiteBatchPeopleSegmentsOperation{People: &people, Segments: segments})

	return service.client.Do(req, nil)
}
