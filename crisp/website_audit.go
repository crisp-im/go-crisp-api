// Copyright 2026 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
	"net/url"
)

// WebsiteAuditLogListData mapping
type WebsiteAuditLogListData struct {
	Data *[]WebsiteAuditLog `json:"data,omitempty"`
}

// WebsiteAuditLogCreate mapping
type WebsiteAuditLogCreate struct {
	Action    string                 `json:"action"`
	Namespace string                 `json:"namespace"`
	UserID    string                 `json:"user_id"`
	IP        *string                `json:"ip,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

// WebsiteAuditLog mapping
type WebsiteAuditLog struct {
	UserID    *string                `json:"user_id,omitempty"`
	Action    *string                `json:"action,omitempty"`
	Namespace *string                `json:"namespace,omitempty"`
	Origin    *string                `json:"origin,omitempty"`
	IP        *string                `json:"ip,omitempty"`
	Country   *string                `json:"country,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
	Timestamp *uint64                `json:"timestamp,omitempty"`
}

// String returns the string representation of WebsiteAuditLog
func (instance WebsiteAuditLog) String() string {
	return Stringify(instance)
}

// CreateAuditLog creates an audit-log entry for a website.
func (service *WebsiteService) CreateAuditLog(websiteID string, auditLog WebsiteAuditLogCreate) (*Response, error) {
	url := fmt.Sprintf("website/%s/audit/log", websiteID)
	req, _ := service.client.NewRequest("POST", url, auditLog)

	return service.client.Do(req, nil)
}

// ListAuditLogs lists audit-log entries for a website.
func (service *WebsiteService) ListAuditLogs(websiteID string, pageNumber uint, filterUserID string, filterNamespace string, filterDateStart string, filterDateEnd string) (*[]WebsiteAuditLog, *Response, error) {
	url := fmt.Sprintf("website/%s/audit/log/list/%d?filter_user_id=%s&filter_namespace=%s&filter_date_start=%s&filter_date_end=%s", websiteID, pageNumber, url.QueryEscape(filterUserID), url.QueryEscape(filterNamespace), url.QueryEscape(filterDateStart), url.QueryEscape(filterDateEnd))
	req, _ := service.client.NewRequest("GET", url, nil)

	auditLogs := new(WebsiteAuditLogListData)
	resp, err := service.client.Do(req, auditLogs)
	if err != nil {
		return nil, resp, err
	}

	return auditLogs.Data, resp, err
}
