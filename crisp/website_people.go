// Copyright 2018 Crisp IM SAS All rights reserved.
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
  Total  *uint64  `json:"total,omitempty"`
}

// PeopleSuggestedSegmentData mapping
type PeopleSuggestedSegmentData struct {
  Data  *[]PeopleSuggestedSegment  `json:"data,omitempty"`
}

// PeopleSuggestedSegment mapping
type PeopleSuggestedSegment struct {
  Segment  *string  `json:"segment,omitempty"`
  Count    *int32   `json:"count,omitempty"`
}

// PeopleSuggestedDataData mapping
type PeopleSuggestedDataData struct {
  Data  *[]PeopleSuggestedData  `json:"data,omitempty"`
}

// PeopleSuggestedData mapping
type PeopleSuggestedData struct {
  Key    *string  `json:"key,omitempty"`
  Count  *int32   `json:"count,omitempty"`
}

// PeopleSuggestedEventData mapping
type PeopleSuggestedEventData struct {
  Data  *[]PeopleSuggestedEvent  `json:"data,omitempty"`
}

// PeopleSuggestedEvent mapping
type PeopleSuggestedEvent struct {
  Text   *string  `json:"text,omitempty"`
  Count  *int32   `json:"count,omitempty"`
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

// PeopleSubscriptionData mapping
type PeopleSubscriptionData struct {
  Data  *PeopleSubscription  `json:"data,omitempty"`
}

// PeopleProfile mapping
type PeopleProfile struct {
  PeopleProfileCard
  PeopleID  *string  `json:"people_id,omitempty"`
}

// PeopleProfileNewData mapping
type PeopleProfileNewData struct {
  Data  *PeopleProfileNew  `json:"data,omitempty"`
}

// PeopleProfileNew mapping
type PeopleProfileNew struct {
  PeopleID  *string  `json:"people_id,omitempty"`
}

// PeopleProfileCard mapping
type PeopleProfileCard struct {
  Email      *string                    `json:"email,omitempty"`
  Person     *PeopleProfileCardPerson   `json:"person,omitempty"`
  Company    *PeopleProfileCardCompany  `json:"company,omitempty"`
  Segments   *[]string                  `json:"segments,omitempty"`
  Notepad    *string                    `json:"notepad,omitempty"`
  Active     *PeopleProfileCardActive   `json:"active,omitempty"`
  Score      *uint8                     `json:"score,omitempty"`
  CreatedAt  *uint64                    `json:"created_at,omitempty"`
  UpdatedAt  *uint64                    `json:"updated_at,omitempty"`
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
  Description  *string                           `json:"description,omitempty"`
  Timezone     *int16                            `json:"timezone,omitempty"`
  Phones       *[]string                         `json:"phones,omitempty"`
  Emails       *[]string                         `json:"emails,omitempty"`
  Geolocation  *PeopleProfileCardGeolocation     `json:"geolocation,omitempty"`
  Metrics      *PeopleProfileCardCompanyMetrics  `json:"metrics,omitempty"`
  Tags         *[]string                         `json:"tags,omitempty"`
}

// PeopleProfileCardCompanyMetrics mapping
type PeopleProfileCardCompanyMetrics struct {
  Employees  *uint32  `json:"employees,omitempty"`
  MarketCap  *uint32  `json:"market_cap,omitempty"`
  Raised     *uint32  `json:"raised,omitempty"`
  ARR        *uint32  `json:"arr,omitempty"`
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
  Now   *bool    `json:"now,omitempty"`
  Last  *uint64  `json:"last,omitempty"`
}

// PeopleConversationsData mapping
type PeopleConversationsData struct {
  Data  []string  `json:"data,omitempty"`
}

// PeopleCampaignsData mapping
type PeopleCampaignsData struct {
  Data  *[]PeopleCampaign  `json:"data,omitempty"`
}

// PeopleCampaign mapping
type PeopleCampaign struct {
  CampaignID    *string    `json:"campaign_id,omitempty"`
  Type          *string    `json:"type,omitempty"`
  Name          *string    `json:"name,omitempty"`
  CreatedAt     *uint64    `json:"created_at,omitempty"`
  UpdatedAt     *uint64    `json:"updated_at,omitempty"`
  DispatchedAt  *uint64    `json:"dispatched_at,omitempty"`
  OccurredAt    *uint64    `json:"occurred_at,omitempty"`
  Statistics    *[]string  `json:"statistics,omitempty"`
}

// PeopleEvent mapping
type PeopleEvent struct {
  Text       *string       `json:"text,omitempty"`
  Data       *interface{}  `json:"data,omitempty"`
  Color      *string       `json:"color,omitempty"`
  Timestamp  *uint64       `json:"timestamp,omitempty"`
}

// PeopleData mapping
type PeopleData struct {
  Data  *interface{}  `json:"data,omitempty"`
}

// PeopleSubscription mapping
type PeopleSubscription struct {
  Email  *bool  `json:"email,omitempty"`
}

// PeopleProfileImportData mapping
type PeopleProfileImportData struct {
  Data  *PeopleProfileImport  `json:"data,omitempty"`
}

// PeopleProfileImport mapping
type PeopleProfileImport struct {
  ImportID  *string  `json:"import_id,omitempty"`
}

// PeopleSuggestedSegmentDelete mapping
type PeopleSuggestedSegmentDelete struct {
  Segment  string  `json:"segment,omitempty"`
}

// PeopleSuggestedDataDelete mapping
type PeopleSuggestedDataDelete struct {
  Key  string  `json:"key,omitempty"`
}

// PeopleSuggestedEventDelete mapping
type PeopleSuggestedEventDelete struct {
  Text  string  `json:"text,omitempty"`
}

// PeopleProfileUpdateCard mapping
type PeopleProfileUpdateCard struct {
  Email     string                     `json:"email,omitempty"`
  Person    *PeopleProfileCardPerson   `json:"person,omitempty"`
  Company   *PeopleProfileCardCompany  `json:"company,omitempty"`
  Segments  []string                   `json:"segments,omitempty"`
  Notepad   string                     `json:"notepad,omitempty"`
  Active    uint64                     `json:"active,omitempty"`
}

// PeopleEventAdd mapping
type PeopleEventAdd struct {
  Text   string       `json:"text,omitempty"`
  Data   interface{}  `json:"data,omitempty"`
  Color  string       `json:"color,omitempty"`
}

// PeopleSubscriptionUpdate mapping
type PeopleSubscriptionUpdate struct {
  Email  bool  `json:"email,omitempty"`
}

// PeopleProfileImportSetup mapping
type PeopleProfileImportSetup struct {
  URL      string                              `json:"url,omitempty"`
  Mapping  *[]PeopleProfileImportSetupMapping  `json:"mapping,omitempty"`
  Options  *PeopleProfileImportSetupOptions    `json:"options,omitempty"`
}

// PeopleProfileImportSetupMapping mapping
type PeopleProfileImportSetupMapping struct {
  Column  uint8   `json:"column,omitempty"`
  Field   string  `json:"field,omitempty"`
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


// String returns the string representation of PeopleSuggestedSegment
func (instance PeopleSuggestedSegment) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleSuggestedData
func (instance PeopleSuggestedData) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleSuggestedEvent
func (instance PeopleSuggestedEvent) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleProfile
func (instance PeopleProfile) String() string {
  return Stringify(instance)
}


// String returns the string representation of PeopleCampaign
func (instance PeopleCampaign) String() string {
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


// String returns the string representation of PeopleSubscription
func (instance PeopleSubscription) String() string {
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


// ListSuggestedPeopleSegments lists suggested segments for people.
func (service *WebsiteService) ListSuggestedPeopleSegments(websiteID string, pageNumber uint) (*[]PeopleSuggestedSegment, *Response, error) {
  url := fmt.Sprintf("website/%s/people/suggest/segments/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  segments := new(PeopleSuggestedSegmentData)
  resp, err := service.client.Do(req, segments)
  if err != nil {
    return nil, resp, err
  }

  return segments.Data, resp, err
}


// DeleteSuggestedPeopleSegment deletes a suggested segment for people.
func (service *WebsiteService) DeleteSuggestedPeopleSegment(websiteID string, segment string) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/suggest/segment", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, PeopleSuggestedSegmentDelete{Segment: segment})

  return service.client.Do(req, nil)
}


// ListSuggestedPeopleDataKeys lists suggested data keys for people.
func (service *WebsiteService) ListSuggestedPeopleDataKeys(websiteID string, pageNumber uint) (*[]PeopleSuggestedData, *Response, error) {
  url := fmt.Sprintf("website/%s/people/suggest/data/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(PeopleSuggestedDataData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data.Data, resp, err
}


// DeleteSuggestedPeopleDataKey deletes a suggested data key for people.
func (service *WebsiteService) DeleteSuggestedPeopleDataKey(websiteID string, key string) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/suggest/data", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, PeopleSuggestedDataDelete{Key: key})

  return service.client.Do(req, nil)
}


// ListSuggestedPeopleEvents lists suggested events for people.
func (service *WebsiteService) ListSuggestedPeopleEvents(websiteID string, pageNumber uint) (*[]PeopleSuggestedEvent, *Response, error) {
  url := fmt.Sprintf("website/%s/people/suggest/events/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  events := new(PeopleSuggestedEventData)
  resp, err := service.client.Do(req, events)
  if err != nil {
    return nil, resp, err
  }

  return events.Data, resp, err
}


// DeleteSuggestedPeopleEvent deletes a suggested event for people.
func (service *WebsiteService) DeleteSuggestedPeopleEvent(websiteID string, text string) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/suggest/event", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, PeopleSuggestedEventDelete{Text: text})

  return service.client.Do(req, nil)
}


// ListPeopleProfiles lists people profiles for website.
func (service *WebsiteService) ListPeopleProfiles(websiteID string, pageNumber uint, searchField string, searchOrder string, searchOperator string, searchFilter []PeopleFilter, searchText string) (*[]PeopleProfile, *Response, error) {
  searchFilterString := ""

  if len(searchFilter) > 0 {
    searchFilterBytes, err := json.Marshal(searchFilter)

    if err == nil {
      searchFilterString = string(searchFilterBytes)
    }
  }

  url := fmt.Sprintf("website/%s/people/profiles/%d?sort_field=%s&sort_order=%s&search_operator=%s&search_filter=%s&search_text=%s", websiteID, pageNumber, url.QueryEscape(searchField), url.QueryEscape(searchOrder), url.QueryEscape(searchOperator), url.QueryEscape(searchFilterString), searchText)
  req, _ := service.client.NewRequest("GET", url, nil)

  people := new(PeopleProfileListData)
  resp, err := service.client.Do(req, people)
  if err != nil {
    return nil, resp, err
  }

  return people.Data, resp, err
}


// AddNewPeopleProfile adds a new people profile.
func (service *WebsiteService) AddNewPeopleProfile(websiteID string, peopleProfile PeopleProfileUpdateCard) (*PeopleProfileNew, *Response, error) {
  url := fmt.Sprintf("website/%s/people/profile", websiteID)
  req, _ := service.client.NewRequest("POST", url, peopleProfile)

  profile := new(PeopleProfileNewData)
  resp, err := service.client.Do(req, profile)
  if err != nil {
    return nil, resp, err
  }

  return profile.Data, resp, err
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


// FilterPeopleConversations filters conversations linked to people.
func (service *WebsiteService) FilterPeopleConversations(websiteID string, peopleID string, pageNumber uint, filterUnread bool, filterResolved bool, filterNotResolved bool) ([]string, *Response, error) {
  var (
    filterUnreadValue string
    filterResolvedValue string
    filterNotResolvedValue string
  )

  if filterUnread == true {
    filterUnreadValue = "1"
  } else {
    filterUnreadValue = "0"
  }

  if filterResolved == true {
    filterResolvedValue = "1"
  } else {
    filterResolvedValue = "0"
  }

  if filterNotResolved == true {
    filterNotResolvedValue = "1"
  } else {
    filterNotResolvedValue = "0"
  }

  url := fmt.Sprintf("website/%s/people/conversations/%s/list/%d?filter_unread=%s&filter_resolved=%s&filter_not_resolved=%s", websiteID, peopleID, pageNumber, url.QueryEscape(filterUnreadValue), url.QueryEscape(filterResolvedValue), url.QueryEscape(filterNotResolvedValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  conversations := new(PeopleConversationsData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// ListPeopleCampaigns lists campaigns linked to people.
func (service *WebsiteService) ListPeopleCampaigns(websiteID string, peopleID string, pageNumber uint) (*[]PeopleCampaign, *Response, error) {
  url := fmt.Sprintf("website/%s/people/campaigns/%s/list/%d", websiteID, peopleID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  campaigns := new(PeopleCampaignsData)
  resp, err := service.client.Do(req, campaigns)
  if err != nil {
    return nil, resp, err
  }

  return campaigns.Data, resp, err
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


// UpdatePeopleData updates stored data for people.
func (service *WebsiteService) UpdatePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/data/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("PATCH", url, peopleData)

  return service.client.Do(req, nil)
}


// GetPeopleSubscriptionStatus resolves subscription status for people.
func (service *WebsiteService) GetPeopleSubscriptionStatus(websiteID string, peopleID string) (*PeopleSubscription, *Response, error) {
  url := fmt.Sprintf("website/%s/people/subscription/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("GET", url, nil)

  subscription := new(PeopleSubscriptionData)
  resp, err := service.client.Do(req, subscription)
  if err != nil {
    return nil, resp, err
  }

  return subscription.Data, resp, err
}


// UpdatePeopleSubscriptionStatus updates current subscription status for people.
func (service *WebsiteService) UpdatePeopleSubscriptionStatus(websiteID string, peopleID string, peopleSubscription PeopleSubscriptionUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/people/subscription/%s", websiteID, peopleID)
  req, _ := service.client.NewRequest("PATCH", url, peopleSubscription)

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
