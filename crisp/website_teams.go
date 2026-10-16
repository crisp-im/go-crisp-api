// Copyright 2026 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp

import (
	"fmt"
)

// WebsiteTeamListData mapping
type WebsiteTeamListData struct {
	Data *[]WebsiteTeam `json:"data,omitempty"`
}

// WebsiteTeamData mapping
type WebsiteTeamData struct {
	Data *WebsiteTeam `json:"data,omitempty"`
}

// WebsiteTeam mapping
type WebsiteTeam struct {
	WebsiteTeamItem
	TeamID    *string `json:"team_id,omitempty"`
	CreatedAt *uint64 `json:"created_at,omitempty"`
	UpdatedAt *uint64 `json:"updated_at,omitempty"`
}

// WebsiteTeamItem mapping
type WebsiteTeamItem struct {
	Name      *string   `json:"name,omitempty"`
	Emoji     *string   `json:"emoji,omitempty"`
	Operators *[]string `json:"operators,omitempty"`
}

// String returns the string representation of WebsiteTeam
func (instance WebsiteTeam) String() string {
	return Stringify(instance)
}

// ListTeams lists all teams for website.
func (service *WebsiteService) ListTeams(websiteID string, pageNumber uint) (*[]WebsiteTeam, *Response, error) {
	url := fmt.Sprintf("website/%s/teams/list/%d", websiteID, pageNumber)
	req, _ := service.client.NewRequest("GET", url, nil)

	teams := new(WebsiteTeamListData)
	resp, err := service.client.Do(req, teams)
	if err != nil {
		return nil, resp, err
	}

	return teams.Data, resp, err
}

// CheckTeamExists checks if team exists for website.
func (service *WebsiteService) CheckTeamExists(websiteID string, teamID string) (*Response, error) {
	url := fmt.Sprintf("website/%s/team/%s", websiteID, teamID)
	req, _ := service.client.NewRequest("HEAD", url, nil)

	return service.client.Do(req, nil)
}

// GetTeam resolves team for website.
func (service *WebsiteService) GetTeam(websiteID string, teamID string) (*WebsiteTeam, *Response, error) {
	url := fmt.Sprintf("website/%s/team/%s", websiteID, teamID)
	req, _ := service.client.NewRequest("GET", url, nil)

	team := new(WebsiteTeamData)
	resp, err := service.client.Do(req, team)
	if err != nil {
		return nil, resp, err
	}

	return team.Data, resp, err
}
