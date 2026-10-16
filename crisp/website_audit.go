// Copyright 2026 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
)

// WebsiteAuditLogCreate mapping
type WebsiteAuditLogCreate struct {
	Action    string                 `json:"action"`
	Namespace string                 `json:"namespace"`
	UserID    string                 `json:"user_id"`
	IP        *string                `json:"ip,omitempty"`
	Context   map[string]interface{} `json:"context,omitempty"`
}

// CreateAuditLog creates an audit-log entry for a website.
func (service *WebsiteService) CreateAuditLog(websiteID string, auditLog WebsiteAuditLogCreate) (*Response, error) {
	url := fmt.Sprintf("website/%s/audit/log", websiteID)
	req, _ := service.client.NewRequest("POST", url, auditLog)

	return service.client.Do(req, nil)
}
