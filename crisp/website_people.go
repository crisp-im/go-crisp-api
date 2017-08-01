// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "encoding/json"
  "net/url"
)


// PeopleStatisticsData mapping
type PeopleStatisticsData struct {
  Data  *PeopleStatistics  `json:"data,omitempty"`
}

// PeopleStatistics mapping
type PeopleStatistics struct {
  Total  *uint  `json:"total,omitempty"`
}

// PeopleSegmentData mapping
type PeopleSegmentData struct {
  Data  *[]PeopleSegment  `json:"data,omitempty"`
}

// PeopleSegment mapping
type PeopleSegment struct {
  Segment  *string  `json:"segment,omitempty"`
  Count    *int32   `json:"count,omitempty"`
}

// PeopleProfileData mapping
type PeopleProfileData struct {
  Data  *PeopleProfile  `json:"data,omitempty"`
}

// PeopleProfileListData mapping
type PeopleProfileListData struct {
  Data  *[]PeopleProfile  `json:"data,omitempty"`
}

// PeopleEventsData mapping
type PeopleEventsData struct {
  Data  *[]PeopleEvent  `json:"data,omitempty"`
}

// PeopleDataData mapping
type PeopleDataData struct {
  Data  *PeopleData  `json:"data,omitempty"`
}

// PeopleProfile mapping
type PeopleProfile struct {
  PeopleProfileCard
  PeopleID  *string  `json:"people_id,omitempty"`
}

// PeopleProfileCard mapping
type PeopleProfileCard struct {
  Email     *string                    `json:"email,omitempty"`
  Person    *PeopleProfileCardPerson   `json:"person,omitempty"`
  Company   *PeopleProfileCardCompany  `json:"company,omitempty"`
  Segments  *[]string                  `json:"segments,omitempty"`
  Active    *PeopleProfileCardActive   `json:"active,omitempty"`
  Score     *uint8                     `json:"score,omitempty"`
}

// PeopleProfileCardPerson mapping
type PeopleProfileCardPerson struct {
  Nickname     *string                             `json:"nickname,omitempty"`
  Avatar       *string                             `json:"avatar,omitempty"`
  Gender       *string                             `json:"gender,omitempty"`
  Phone        *string                             `json:"phone,omitempty"`
  Address      *string                             `json:"address,omitempty"`
  Description  *string                             `json:"description,omitempty"`
  Website      *string                             `json:"website,omitempty"`
  Timezone     *int16                              `json:"timezone,omitempty"`
  Profiles     *[]PeopleProfileCardPersonProfile   `json:"profiles,omitempty"`
  Employment   *PeopleProfileCardPersonEmployment  `json:"employment,omitempty"`
  Geolocation  *PeopleProfileCardGeolocation       `json:"geolocation,omitempty"`
  Locales      *[]string                           `json:"locales,omitempty"`
}

// PeopleProfileCardPersonProfile mapping
type PeopleProfileCardPersonProfile struct {
  Type    *string  `json:"type,omitempty"`
  Handle  *string  `json:"handle,omitempty"`
}

// PeopleProfileCardPersonEmployment mapping
type PeopleProfileCardPersonEmployment struct {
  Name       *string  `json:"name,omitempty"`
  Domain     *string  `json:"domain,omitempty"`
  Title      *string  `json:"title,omitempty"`
  Role       *string  `json:"role,omitempty"`
  Seniority  *string  `json:"seniority,omitempty"`
}

// PeopleProfileCardCompany mapping
type PeopleProfileCardCompany struct {
  Name         *string                           `json:"name,omitempty"`
  LegalName    *string                           `json:"legal_name,omitempty"`
  Domain       *string                           `json:"domain,omitempty"`
  URL          *string                           `json:"url,omitempty"`
  Timezone     *int16                            `json:"timezone,omitempty"`
  Phones       *[]string                         `json:"phones,omitempty"`
  Emails       *[]string                         `json:"emails,omitempty"`
  Geolocation  *PeopleProfileCardGeolocation     `json:"geolocation,omitempty"`
  Metrics      *PeopleProfileCardCompanyMetrics  `json:"metrics,omitempty"`
  Tags         *[]string                         `json:"tags,omitempty"`
}

// PeopleProfileCardCompanyMetrics mapping
type PeopleProfileCardCompanyMetrics struct {
  Employees  *int16  `json:"employees,omitempty"`
  MarketCap  *int16  `json:"market_cap,omitempty"`
  Raised     *int16  `json:"raised,omitempty"`
  ARR        *int16  `json:"arr,omitempty"`
}

// PeopleProfileCardGeolocation mapping
type PeopleProfileCardGeolocation struct {
  Country      *string                                   `json:"country,omitempty"`
  Region       *string                                   `json:"region,omitempty"`
  City         *string                                   `json:"city,omitempty"`
  Coordinates  *PeopleProfileCardGeolocationCoordinates  `json:"coordinates,omitempty"`
}

// PeopleProfileCardGeolocationCoordinates mapping
type PeopleProfileCardGeolocationCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}

// PeopleProfileCardActive mapping
type PeopleProfileCardActive struct {
  Now   *bool  `json:"now,omitempty"`
  Last  *uint  `json:"last,omitempty"`
}

// PeopleConversationsData mapping
type PeopleConversationsData struct {
  Data  []string  `json:"data,omitempty"`
}

// PeopleEvent mapping
type PeopleEvent struct {
  Text       *string       `json:"text,omitempty"`
  Data       *interface{}  `json:"data,omitempty"`
  Color      *string       `json:"color,omitempty"`
  Timestamp  *uint         `json:"timestamp,omitempty"`
}

// PeopleData mapping
type PeopleData struct {
  Data  *interface{}  `json:"data,omitempty"`
}

// PeopleProfileImportData mapping
type PeopleProfileImportData struct {
  Data  *PeopleProfileImport  `json:"data,omitempty"`
}

// PeopleProfileImport mapping
type PeopleProfileImport struct {
  ImportID  *string  `json:"import_id,omitempty"`
}

// PeopleProfileUpdateCard mapping
type PeopleProfileUpdateCard struct {
  Email     string                     `json:"email,omitempty"`
  Person    *PeopleProfileCardPerson   `json:"person,omitempty"`
  Company   *PeopleProfileCardCompany  `json:"company,omitempty"`
  Segments  []string                   `json:"segments,omitempty"`
}

// PeopleEventAdd mapping
type PeopleEventAdd struct {
  Text   string       `json:"text,omitempty"`
  Data   interface{}  `json:"data,omitempty"`
  Color  string       `json:"color,omitempty"`
}

// PeopleProfileImportSetup mapping
type PeopleProfileImportSetup struct {
  URL      string                              `json:"url,omitempty"`
  Mapping  *[]PeopleProfileImportSetupMapping  `json:"mapping,omitempty"`
  Options  *PeopleProfileImportSetupOptions    `json:"options,omitempty"`
}

// PeopleProfileImportSetupMapping mapping
type PeopleProfileImportSetupMapping struct {
  Column  uint8  `json:"column,omitempty"`
  Field   bool   `json:"field,omitempty"`
}

// PeopleProfileImportSetupOptions mapping
type PeopleProfileImportSetupOptions struct {
  ColumnSeparator  string  `json:"column_separator,omitempty"`
  SkipHeader       bool    `json:"skip_header,omitempty"`
}

// PeopleFilter mapping
type PeopleFilter struct {
  Model      string    `json:"model,omitempty"`
  Criterion  string    `json:"criterion,omitempty"`
  Operator   string    `json:"operator,omitempty"`
  Query      []string  `json:"query,omitempty"`
}


// String returns the string representation of PeopleStatistics
func (instance PeopleStatistics) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleSegment
func (instance PeopleSegment) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleProfile
func (instance PeopleProfile) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleEvent
func (instance PeopleEvent) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleData
func (instance PeopleData) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleProfileImport
func (instance PeopleProfileImport) String() string {
  return Stringify(instance)
}


// GetPeopleStatistics resolves statistics on people in website.
func (service *WebsiteService) GetPeopleStatistics(websiteID string) (*PeopleStatistics, *Response, error) {
  url := fmt.Sprintf("website/%s/people/stats", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  stats := new(PeopleStatisticsData)
  resp, err := service.client.Do(req, stats)
  if err != nil {
    return nil, resp, err
  }

  return stats.Data, resp, err
}


// ListPeopleSegments lists segments in use for people.
func (service *WebsiteService) ListPeopleSegments(websiteID string, pageNumber uint) (*[]PeopleSegment, *Response, error) {
  url := fmt.Sprintf("website/%s/people/segments/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  segments := new(PeopleSegmentData)
  resp, err := service.client.Do(req, segments)
  if err != nil {
    return nil, resp, err
  }

  return segments.Data, resp, err
}


// ListPeopleProfiles lists people profiles for website.
func (service *WebsiteService) ListPeopleProfiles(websiteID string, pageNumber uint, searchField string, searchOrder string, searchOperator string, searchFilter []PeopleFilter) (*[]PeopleProfile, *Response, error) {
  searchFilterString := ""

  if len(searchFilter) > 0 {
    searchFilterBytes, err := json.Marshal(searchFilter)

    if err == nil {
      searchFilterString = string(searchFilterBytes)
    }
  }

  url := fmt.Sprintf("website/%s/people/profiles/%d?sort_field=%s&sort_order=%s&search_operator=%s&search_filter=%s", websiteID, pageNumber, url.QueryEscape(searchField), url.QueryEscape(searchOrder), url.QueryEscape(searchOperator), url.QueryEscape(searchFilterString))
  req, _ := service.client.NewRequest("GET", url, nil)

  people := new(PeopleProfileListData)
  resp, err := service.client.Do(req, people)
  if err != nil {
    return nil, resp, err
  }

  return people.Data, resp, err
}


// AddNewPeopleProfile adds a new people profile.
func (service *WebsiteService) AddNewPeopleProfile(websiteID string, peopleProfile PeopleProfileUpdateCard) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/profile", websiteID)
  req, _ := service.client.NewRequest("POST", url, peopleProfile)

  return service.client.Do(req, nil)
}


// CheckPeopleProfileExists checks if given people profile exists.
func (service *WebsiteService) CheckPeopleProfileExists(websiteID string, peopleID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/profile/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// GetPeopleProfile resolves people profile.
func (service *WebsiteService) GetPeopleProfile(websiteID string, peopleID string) (*PeopleProfile, *Response, error) {
  url := fmt.Sprintf("website/%s/people/profile/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("GET", url, nil)

  profile := new(PeopleProfileData)
  resp, err := service.client.Do(req, profile)
  if err != nil {
    return nil, resp, err
  }

  return profile.Data, resp, err
}


// SavePeopleProfile saves people profile, and overwrite all previous information.
func (service *WebsiteService) SavePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/profile/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("PUT", url, peopleProfile)

  return service.client.Do(req, nil)
}


// UpdatePeopleProfile updates people profile, and save only changed fields on the previous profile revision.
func (service *WebsiteService) UpdatePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/profile/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("PATCH", url, peopleProfile)

  return service.client.Do(req, nil)
}


// RemovePeopleProfile removes people profile in website.
func (service *WebsiteService) RemovePeopleProfile(websiteID string, peopleID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/profile/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListPeopleConversations lists conversations linked to people.
func (service *WebsiteService) ListPeopleConversations(websiteID string, peopleID string, pageNumber uint) ([]string, *Response, error) {
  url := fmt.Sprintf("website/%s/people/conversations/%s/list/%d", websiteID, peopleID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  conversations := new(PeopleConversationsData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// AddPeopleEvent stacks an event for people.
func (service *WebsiteService) AddPeopleEvent(websiteID string, peopleID string, peopleEvent PeopleEventAdd) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/events/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("POST", url, peopleEvent)

  return service.client.Do(req, nil)
}


// ListPeopleEvents lists stacked events for people.
func (service *WebsiteService) ListPeopleEvents(websiteID string, peopleID string, pageNumber uint) (*[]PeopleEvent, *Response, error) {
  url := fmt.Sprintf("website/%s/people/events/%s/list/%d", websiteID, peopleID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  events := new(PeopleEventsData)
  resp, err := service.client.Do(req, events)
  if err != nil {
    return nil, resp, err
  }

  return events.Data, resp, err
}


// GetPeopleData gets stored data for people.
func (service *WebsiteService) GetPeopleData(websiteID string, peopleID string) (*PeopleData, *Response, error) {
  url := fmt.Sprintf("website/%s/people/data/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(PeopleDataData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data.Data, resp, err
}


// SavePeopleData saves stored data for people.
func (service *WebsiteService) SavePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/data/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("PUT", url, peopleData)

  return service.client.Do(req, nil)
}


// ExportPeopleProfiles exports people profiles.
func (service *WebsiteService) ExportPeopleProfiles(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/export/profiles", websiteID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ImportPeopleProfiles imports people profiles.
func (service *WebsiteService) ImportPeopleProfiles(websiteID string, profileImportSetup PeopleProfileImportSetup) (*PeopleProfileImport, *Response, error) {
  url := fmt.Sprintf("website/%s/people/import/profiles", websiteID)
  req, _ := service.client.NewRequest("POST", url, profileImportSetup)

  profileImport := new(PeopleProfileImportData)
  resp, err := service.client.Do(req, profileImport)
  if err != nil {
    return nil, resp, err
  }

  return profileImport.Data, resp, err
}
