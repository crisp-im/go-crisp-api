// Copyright 2023 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// HelpdeskData mapping
type HelpdeskData struct {
  Data  *Helpdesk  `json:"data,omitempty"`
}

// Helpdesk mapping
type Helpdesk struct {
  Name  *string  `json:"name,omitempty"`
  URL   *string  `json:"url,omitempty"`
}

// HelpdeskLocaleListData mapping
type HelpdeskLocaleListData struct {
  Data  *[]HelpdeskLocale  `json:"data,omitempty"`
}

// HelpdeskLocaleData mapping
type HelpdeskLocaleData struct {
  Data  *HelpdeskLocale  `json:"data,omitempty"`
}

// HelpdeskLocale mapping
type HelpdeskLocale struct {
  LocaleID    *string                     `json:"locale_id,omitempty"`
  Locale      *string                     `json:"locale,omitempty"`
  URL         *string                     `json:"url,omitempty"`
  Statistics  *HelpdeskLocaleStatistics   `json:"statistics,omitempty"`
}

// HelpdeskLocaleStatistics mapping
type HelpdeskLocaleStatistics struct {
  Articles    *HelpdeskLocaleStatisticsContent  `json:"articles,omitempty"`
  Guides      *HelpdeskLocaleStatisticsContent  `json:"guides,omitempty"`
  References  *HelpdeskLocaleStatisticsContent  `json:"references,omitempty"`
  News        *HelpdeskLocaleStatisticsContent  `json:"news,omitempty"`
}

// HelpdeskLocaleStatisticsContent mapping
type HelpdeskLocaleStatisticsContent struct {
  Entries  *uint32  `json:"entries,omitempty"`
  Groups   *uint32  `json:"groups,omitempty"`
}

// HelpdeskLocaleFeedbackRatingsData mapping
type HelpdeskLocaleFeedbackRatingsData struct {
  Data  *HelpdeskLocaleFeedbackRatings  `json:"data,omitempty"`
}

// HelpdeskLocaleFeedbackRatings mapping
type HelpdeskLocaleFeedbackRatings struct {
  Ratings  *HelpdeskLocaleFeedbackRatingsRatings  `json:"ratings,omitempty"`
}

// HelpdeskLocaleFeedbackRatingsRatings mapping
type HelpdeskLocaleFeedbackRatingsRatings struct {
  Helpful    *uint32  `json:"helpful,omitempty"`
  Unhelpful  *uint32  `json:"unhelpful,omitempty"`
}

// HelpdeskLocaleFeedbackListData mapping
type HelpdeskLocaleFeedbackListData struct {
  Data  *[]HelpdeskLocaleFeedbackItem  `json:"data,omitempty"`
}

// HelpdeskLocaleFeedbackItem mapping
type HelpdeskLocaleFeedbackItem struct {
  Rating     *string                             `json:"rating,omitempty"`
  Comment    *string                             `json:"comment,omitempty"`
  Article    *HelpdeskLocaleFeedbackItemArticle  `json:"article,omitempty"`
  Session    *HelpdeskLocaleFeedbackItemSession  `json:"session,omitempty"`
  CreatedAt  *uint64                             `json:"created_at,omitempty"`
}

// HelpdeskLocaleFeedbackItemArticle mapping
type HelpdeskLocaleFeedbackItemArticle struct {
  ArticleID  *string  `json:"article_id,omitempty"`
  Title      *string  `json:"title,omitempty"`
  URL        *string  `json:"url,omitempty"`
}

// HelpdeskLocaleFeedbackItemSession mapping
type HelpdeskLocaleFeedbackItemSession struct {
  SessionID    *string                                        `json:"session_id,omitempty"`
  Nickname     *string                                        `json:"nickname,omitempty"`
  Email        *string                                        `json:"email,omitempty"`
  Avatar       *string                                        `json:"avatar,omitempty"`
  Geolocation  *HelpdeskLocaleFeedbackItemSessionGeolocation  `json:"geolocation,omitempty"`
  Assigned     *HelpdeskLocaleFeedbackItemSessionAssigned     `json:"assigned,omitempty"`
}

// HelpdeskLocaleFeedbackItemSessionGeolocation mapping
type HelpdeskLocaleFeedbackItemSessionGeolocation struct {
  Country      *string                                                   `json:"country,omitempty"`
  Region       *string                                                   `json:"region,omitempty"`
  City         *string                                                   `json:"city,omitempty"`
  Coordinates  *HelpdeskLocaleFeedbackItemSessionGeolocationCoordinates  `json:"coordinates,omitempty"`
}

// HelpdeskLocaleFeedbackItemSessionGeolocationCoordinates mapping
type HelpdeskLocaleFeedbackItemSessionGeolocationCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}

// HelpdeskLocaleFeedbackItemSessionAssigned mapping
type HelpdeskLocaleFeedbackItemSessionAssigned struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// HelpdeskPageListData mapping
type HelpdeskPageListData struct {
  Data  *[]HelpdeskPage  `json:"data,omitempty"`
}

// HelpdeskPage mapping
type HelpdeskPage struct {
  EntityID  *string  `json:"entity_id,omitempty"`
  Title     *string  `json:"title,omitempty"`
  URL       *string  `json:"url,omitempty"`
}

// HelpdeskTreeListData mapping
type HelpdeskTreeListData struct {
  Data  *[]HelpdeskTreeEntry  `json:"data,omitempty"`
}

// HelpdeskTreeData mapping
type HelpdeskTreeData struct {
  Data  *HelpdeskTreeEntry  `json:"data,omitempty"`
}

// HelpdeskTreeEntry mapping
type HelpdeskTreeEntry struct {
  Type      *string                 `json:"type,omitempty"`
  Slug      *string                 `json:"slug,omitempty"`
  Title     *string                 `json:"title,omitempty"`
  State     *HelpdeskTreeState      `json:"state,omitempty"`
  Children  *[]HelpdeskTreeEntry    `json:"children,omitempty"`
}

// HelpdeskTreeState mapping
type HelpdeskTreeState struct {
  Published  *bool  `json:"published,omitempty"`
  Hidden     *bool  `json:"hidden,omitempty"`
  Featured   *bool  `json:"featured,omitempty"`
}

// HelpdeskTreeContentData mapping
type HelpdeskTreeContentData struct {
  Data  *HelpdeskTreeContent  `json:"data,omitempty"`
}

// HelpdeskTreeContent mapping
type HelpdeskTreeContent struct {
  Content  *string  `json:"content,omitempty"`
}

// HelpdeskTreeMetadataData mapping
type HelpdeskTreeMetadataData struct {
  Data  *HelpdeskTreeMetadata  `json:"data,omitempty"`
}

// HelpdeskTreeMetadata mapping
type HelpdeskTreeMetadata struct {
  Format       *string                       `json:"format,omitempty"`
  Title        *string                       `json:"title,omitempty"`
  Description  *string                       `json:"description,omitempty"`
  State        *HelpdeskTreeState            `json:"state,omitempty"`
  Author       *HelpdeskTreeMetadataAuthor   `json:"author,omitempty"`
  Color        *string                       `json:"color,omitempty"`
  Image        *string                       `json:"image,omitempty"`
}

// HelpdeskTreeMetadataAuthor mapping
type HelpdeskTreeMetadataAuthor struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// HelpdeskTreePageData mapping
type HelpdeskTreePageData struct {
  Data  *HelpdeskTreePage  `json:"data,omitempty"`
}

// HelpdeskTreePage mapping
type HelpdeskTreePage struct {
  Title  *string  `json:"title,omitempty"`
  URL    *string  `json:"url,omitempty"`
}

// HelpdeskHistoryChangeListData mapping
type HelpdeskHistoryChangeListData struct {
  Data  *[]HelpdeskHistoryChange  `json:"data,omitempty"`
}

// HelpdeskHistoryChangeData mapping
type HelpdeskHistoryChangeData struct {
  Data  *HelpdeskHistoryChange  `json:"data,omitempty"`
}

// HelpdeskHistoryChange mapping
type HelpdeskHistoryChange struct {
  ChangeID   *string                       `json:"change_id,omitempty"`
  Message    *string                       `json:"message,omitempty"`
  Author     *HelpdeskHistoryChangeAuthor  `json:"author,omitempty"`
  Edits      *HelpdeskHistoryChangeEdits   `json:"edits,omitempty"`
  CreatedAt  *uint64                       `json:"created_at,omitempty"`
  Files      *[]HelpdeskHistoryChangeFile  `json:"files,omitempty"`
}

// HelpdeskHistoryChangeAuthor mapping
type HelpdeskHistoryChangeAuthor struct {
  Email  *string  `json:"email,omitempty"`
  Name   *string  `json:"name,omitempty"`
}

// HelpdeskHistoryChangeEdits mapping
type HelpdeskHistoryChangeEdits struct {
  Insertions  *uint32  `json:"insertions,omitempty"`
  Deletions   *uint32  `json:"deletions,omitempty"`
}

// HelpdeskHistoryChangeFile mapping
type HelpdeskHistoryChangeFile struct {
  Change    *string                 `json:"change,omitempty"`
  Path      *string                 `json:"path,omitempty"`
  Metadata  *HelpdeskTreeMetadata   `json:"metadata,omitempty"`
  Content   *string                 `json:"content,omitempty"`
}

// HelpdeskRedirectionListData mapping
type HelpdeskRedirectionListData struct {
  Data  *[]HelpdeskRedirection  `json:"data,omitempty"`
}

// HelpdeskRedirectionData mapping
type HelpdeskRedirectionData struct {
  Data  *HelpdeskRedirection  `json:"data,omitempty"`
}

// HelpdeskRedirection mapping
type HelpdeskRedirection struct {
  RedirectionID  *string  `json:"redirection_id,omitempty"`
  Path           *string  `json:"path,omitempty"`
  Target         *string  `json:"target,omitempty"`
  CreatedAt      *uint64  `json:"created_at,omitempty"`
  UpdatedAt      *uint64  `json:"updated_at,omitempty"`
}

// HelpdeskRedirectionNewData mapping
type HelpdeskRedirectionNewData struct {
  Data  *HelpdeskRedirectionNew  `json:"data,omitempty"`
}

// HelpdeskRedirectionNew mapping
type HelpdeskRedirectionNew struct {
  RedirectionID  *string  `json:"redirection_id,omitempty"`
}

// HelpdeskSettingsData mapping
type HelpdeskSettingsData struct {
  Data  *HelpdeskSettings  `json:"data,omitempty"`
}

// HelpdeskSettings mapping
type HelpdeskSettings struct {
  Name          *string                        `json:"name,omitempty"`
  Appearance    *HelpdeskSettingsAppearance    `json:"appearance,omitempty"`
  Localization  *HelpdeskSettingsLocalization  `json:"localization,omitempty"`
  Section       *HelpdeskSettingsSection       `json:"section,omitempty"`
  Behavior      *HelpdeskSettingsBehavior      `json:"behavior,omitempty"`
  Service       *HelpdeskSettingsService       `json:"service,omitempty"`
  Include       *HelpdeskSettingsInclude       `json:"include,omitempty"`
  Access        *HelpdeskSettingsAccess        `json:"access,omitempty"`
}

// HelpdeskSettingsAppearance mapping
type HelpdeskSettingsAppearance struct {
  Color   *HelpdeskSettingsAppearanceColor  `json:"color,omitempty"`
  Logos   *HelpdeskSettingsAppearanceLogos  `json:"logos,omitempty"`
  Banner  *string                           `json:"banner,omitempty"`
}

// HelpdeskSettingsAppearanceColor mapping
type HelpdeskSettingsAppearanceColor struct {
  Mode            *string  `json:"mode,omitempty"`
  ModeChangeable  *bool    `json:"mode_changeable,omitempty"`
}

// HelpdeskSettingsAppearanceLogos mapping
type HelpdeskSettingsAppearanceLogos struct {
  Favicon  *string  `json:"favicon,omitempty"`
  Header   *string  `json:"header,omitempty"`
  Footer   *string  `json:"footer,omitempty"`
}

// HelpdeskSettingsLocalization mapping
type HelpdeskSettingsLocalization struct {
  WritingLocale               *string  `json:"writing_locale,omitempty"`
  TranslatedAutomatic         *bool    `json:"translated_automatic,omitempty"`
  TranslatedLocalesReadonly   *bool    `json:"translated_locales_readonly,omitempty"`
}

// HelpdeskSettingsSection mapping
type HelpdeskSettingsSection struct {
  Articles    *string  `json:"articles,omitempty"`
  Guides      *string  `json:"guides,omitempty"`
  References  *string  `json:"references,omitempty"`
  News        *string  `json:"news,omitempty"`
}

// HelpdeskSettingsBehavior mapping
type HelpdeskSettingsBehavior struct {
  FrequentlyRead      *bool  `json:"frequently_read,omitempty"`
  ShowCategoryImages  *bool  `json:"show_category_images,omitempty"`
  ShowChatbox         *bool  `json:"show_chatbox,omitempty"`
  AskFeedback         *bool  `json:"ask_feedback,omitempty"`
  ReportIncorrect     *bool  `json:"report_incorrect,omitempty"`
  ServeMarkdown       *bool  `json:"serve_markdown,omitempty"`
  AgentChatBar        *bool  `json:"agent_chat_bar,omitempty"`
  AgentCopyButton     *bool  `json:"agent_copy_button,omitempty"`
  TableOfContents     *bool  `json:"table_of_contents,omitempty"`
  LocalePicker        *bool  `json:"locale_picker,omitempty"`
  ReferLink           *bool  `json:"refer_link,omitempty"`
  ForbidIndexing      *bool  `json:"forbid_indexing,omitempty"`
  StatusHealthDead    *bool  `json:"status_health_dead,omitempty"`
}

// HelpdeskSettingsService mapping
type HelpdeskSettingsService struct {
  MCPServer  *bool  `json:"mcp_server,omitempty"`
}

// HelpdeskSettingsInclude mapping
type HelpdeskSettingsInclude struct {
  HTML  *string  `json:"html,omitempty"`
}

// HelpdeskSettingsAccess mapping
type HelpdeskSettingsAccess struct {
  RestrictMode  *string  `json:"restrict_mode,omitempty"`
  Password      *string  `json:"password,omitempty"`
  JWTSecret     *string  `json:"jwt_secret,omitempty"`
}

// HelpdeskDomainData mapping
type HelpdeskDomainData struct {
  Data  *HelpdeskDomain  `json:"data,omitempty"`
}

// HelpdeskDomain mapping
type HelpdeskDomain struct {
  Root      *string  `json:"root,omitempty"`
  Basic     *string  `json:"basic,omitempty"`
  Custom    *string  `json:"custom,omitempty"`
  Verified  *bool    `json:"verified,omitempty"`
}

// HelpdeskDomainSetupFlowData mapping
type HelpdeskDomainSetupFlowData struct {
  Data  *HelpdeskDomainSetupFlow  `json:"data,omitempty"`
}

// HelpdeskDomainSetupFlow mapping
type HelpdeskDomainSetupFlow struct {
  Custom  *string                        `json:"custom,omitempty"`
  Setup   *HelpdeskDomainSetupFlowSetup  `json:"setup,omitempty"`
}

// HelpdeskDomainSetupFlowSetup mapping
type HelpdeskDomainSetupFlowSetup struct {
  Records  *[]HelpdeskDomainSetupFlowSetupRecord  `json:"records,omitempty"`
}

// HelpdeskDomainSetupFlowSetupRecord mapping
type HelpdeskDomainSetupFlowSetupRecord struct {
  Type   *string  `json:"type,omitempty"`
  Query  *string  `json:"query,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// HelpdeskCreate mapping
type HelpdeskCreate struct {
  Name         string  `json:"name"`
  DomainBasic  string  `json:"domain_basic"`
}

// HelpdeskRemove mapping
type HelpdeskRemove struct {
  Verify  *HelpdeskRemoveVerify  `json:"verify,omitempty"`
}

// HelpdeskRemoveVerify mapping
type HelpdeskRemoveVerify struct {
  Method  string  `json:"method"`
  Secret  string  `json:"secret"`
}

// HelpdeskLocaleAdd mapping
type HelpdeskLocaleAdd struct {
  Locale  string  `json:"locale"`
}

// HelpdeskLocaleExternalImport mapping
type HelpdeskLocaleExternalImport struct {
  HelpdeskURL    string     `json:"helpdesk_url"`
  DetectLocales  *bool      `json:"detect_locales,omitempty"`
  OtherLocales   *[]string  `json:"other_locales,omitempty"`
}

// HelpdeskTreePathUpdate mapping
type HelpdeskTreePathUpdate struct {
  Action    string                        `json:"action"`
  Path      *HelpdeskTreePathUpdatePath   `json:"path,omitempty"`
  Position  *int16                        `json:"position,omitempty"`
}

// HelpdeskTreePathUpdatePath mapping
type HelpdeskTreePathUpdatePath struct {
  To  string  `json:"to"`
}

// HelpdeskTreeContentSave mapping
type HelpdeskTreeContentSave struct {
  Content  string  `json:"content"`
}

// HelpdeskHistoryChangeCancel mapping
type HelpdeskHistoryChangeCancel struct {
  Action  string  `json:"action"`
}

// HelpdeskRedirectionAdd mapping
type HelpdeskRedirectionAdd struct {
  Path    string  `json:"path"`
  Target  string  `json:"target"`
}

// HelpdeskDomainChangeRequest mapping
type HelpdeskDomainChangeRequest struct {
  Basic   *string  `json:"basic,omitempty"`
  Custom  *string  `json:"custom,omitempty"`
}


// String returns the string representation of Helpdesk
func (instance Helpdesk) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocale
func (instance HelpdeskLocale) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleFeedbackRatings
func (instance HelpdeskLocaleFeedbackRatings) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleFeedbackItem
func (instance HelpdeskLocaleFeedbackItem) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskPage
func (instance HelpdeskPage) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskTreeEntry
func (instance HelpdeskTreeEntry) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskTreeContent
func (instance HelpdeskTreeContent) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskTreeMetadata
func (instance HelpdeskTreeMetadata) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskTreePage
func (instance HelpdeskTreePage) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskHistoryChange
func (instance HelpdeskHistoryChange) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskRedirection
func (instance HelpdeskRedirection) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskRedirectionNew
func (instance HelpdeskRedirectionNew) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskSettings
func (instance HelpdeskSettings) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskDomain
func (instance HelpdeskDomain) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskDomainSetupFlow
func (instance HelpdeskDomainSetupFlow) String() string {
  return Stringify(instance)
}


// CheckHelpdeskExists checks if helpdesk exists for website.
func (service *WebsiteService) CheckHelpdeskExists(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk", websiteID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdesk resolves helpdesk information for website.
func (service *WebsiteService) ResolveHelpdesk(websiteID string) (*Helpdesk, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  helpdesk := new(HelpdeskData)
  resp, err := service.client.Do(req, helpdesk)
  if err != nil {
    return nil, resp, err
  }

  return helpdesk.Data, resp, err
}


// InitializeHelpdesk initializes a new helpdesk for website.
func (service *WebsiteService) InitializeHelpdesk(websiteID string, name string, domainBasic string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk", websiteID)
  req, _ := service.client.NewRequest("POST", url, HelpdeskCreate{Name: name, DomainBasic: domainBasic})

  return service.client.Do(req, nil)
}


// DeleteHelpdesk deletes helpdesk for website.
func (service *WebsiteService) DeleteHelpdesk(websiteID string, verify HelpdeskRemoveVerify) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, HelpdeskRemove{Verify: &verify})

  return service.client.Do(req, nil)
}


// ListHelpdeskLocales lists locales for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocales(websiteID string, pageNumber uint) (*[]HelpdeskLocale, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locales/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  locales := new(HelpdeskLocaleListData)
  resp, err := service.client.Do(req, locales)
  if err != nil {
    return nil, resp, err
  }

  return locales.Data, resp, err
}


// AddHelpdeskLocale adds a locale for helpdesk in website.
func (service *WebsiteService) AddHelpdeskLocale(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale", websiteID)
  req, _ := service.client.NewRequest("POST", url, HelpdeskLocaleAdd{Locale: locale})

  return service.client.Do(req, nil)
}


// CheckHelpdeskLocaleExists checks if a helpdesk locale exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleExists(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocale resolves a locale for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocale(websiteID string, locale string) (*HelpdeskLocale, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("GET", url, nil)

  helpdeskLocale := new(HelpdeskLocaleData)
  resp, err := service.client.Do(req, helpdeskLocale)
  if err != nil {
    return nil, resp, err
  }

  return helpdeskLocale.Data, resp, err
}


// DeleteHelpdeskLocale deletes a locale for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskPages lists publicly visible helpdesk pages for a locale and content type.
func (service *WebsiteService) ListHelpdeskPages(websiteID string, locale string, contentType string, pageNumber uint) (*[]HelpdeskPage, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/page/list/%s/%s/%d", websiteID, locale, contentType, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  pages := new(HelpdeskPageListData)
  resp, err := service.client.Do(req, pages)
  if err != nil {
    return nil, resp, err
  }

  return pages.Data, resp, err
}


// ListHelpdeskTree lists tree entries for a helpdesk locale and content type.
func (service *WebsiteService) ListHelpdeskTree(websiteID string, locale string, contentType string, pageNumber uint, subPath string, searchTitle string, filterDateStart string, filterDateEnd string) (*[]HelpdeskTreeEntry, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/list/%s/%s/%d?sub_path=%s&search_title=%s&filter_date_start=%s&filter_date_end=%s", websiteID, locale, contentType, pageNumber, url.QueryEscape(subPath), url.QueryEscape(searchTitle), url.QueryEscape(filterDateStart), url.QueryEscape(filterDateEnd))
  req, _ := service.client.NewRequest("GET", url, nil)

  tree := new(HelpdeskTreeListData)
  resp, err := service.client.Do(req, tree)
  if err != nil {
    return nil, resp, err
  }

  return tree.Data, resp, err
}


// CreateHelpdeskTreePath creates an empty helpdesk tree path.
func (service *WebsiteService) CreateHelpdeskTreePath(websiteID string, locale string, contentType string, path string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/path/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskTreePath resolves a helpdesk tree path.
func (service *WebsiteService) ResolveHelpdeskTreePath(websiteID string, locale string, contentType string, path string) (*HelpdeskTreeEntry, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/path/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("GET", url, nil)

  tree := new(HelpdeskTreeData)
  resp, err := service.client.Do(req, tree)
  if err != nil {
    return nil, resp, err
  }

  return tree.Data, resp, err
}


// UpdateHelpdeskTreePath moves a helpdesk tree path or changes its display order.
func (service *WebsiteService) UpdateHelpdeskTreePath(websiteID string, locale string, contentType string, path string, update HelpdeskTreePathUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/path/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("PATCH", url, update)

  return service.client.Do(req, nil)
}


// DeleteHelpdeskTreePath deletes a helpdesk tree path recursively.
func (service *WebsiteService) DeleteHelpdeskTreePath(websiteID string, locale string, contentType string, path string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/path/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskTreeContent resolves the body content of a helpdesk tree file.
func (service *WebsiteService) ResolveHelpdeskTreeContent(websiteID string, locale string, contentType string, path string) (*HelpdeskTreeContent, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/content/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("GET", url, nil)

  content := new(HelpdeskTreeContentData)
  resp, err := service.client.Do(req, content)
  if err != nil {
    return nil, resp, err
  }

  return content.Data, resp, err
}


// SaveHelpdeskTreeContent replaces the body content of a helpdesk tree file.
func (service *WebsiteService) SaveHelpdeskTreeContent(websiteID string, locale string, contentType string, path string, content string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/content/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("PUT", url, HelpdeskTreeContentSave{Content: content})

  return service.client.Do(req, nil)
}


// ResolveHelpdeskTreeMetadata resolves parsed metadata for a helpdesk tree path.
func (service *WebsiteService) ResolveHelpdeskTreeMetadata(websiteID string, locale string, contentType string, path string) (*HelpdeskTreeMetadata, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/metadata/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("GET", url, nil)

  metadata := new(HelpdeskTreeMetadataData)
  resp, err := service.client.Do(req, metadata)
  if err != nil {
    return nil, resp, err
  }

  return metadata.Data, resp, err
}


// UpdateHelpdeskTreeMetadata updates metadata for a helpdesk tree path.
func (service *WebsiteService) UpdateHelpdeskTreeMetadata(websiteID string, locale string, contentType string, path string, metadata HelpdeskTreeMetadata) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/metadata/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("PATCH", url, metadata)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskTreePage resolves page information for a helpdesk tree file.
func (service *WebsiteService) ResolveHelpdeskTreePage(websiteID string, locale string, contentType string, path string) (*HelpdeskTreePage, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/tree/page/%s/%s/%s", websiteID, locale, contentType, path)
  req, _ := service.client.NewRequest("GET", url, nil)

  page := new(HelpdeskTreePageData)
  resp, err := service.client.Do(req, page)
  if err != nil {
    return nil, resp, err
  }

  return page.Data, resp, err
}


// ListHelpdeskHistoryChanges lists helpdesk content history changes.
func (service *WebsiteService) ListHelpdeskHistoryChanges(websiteID string, pageNumber uint, filterLocale string, filterType string, filterTreePath string) (*[]HelpdeskHistoryChange, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/history/changes/%d?filter_locale=%s&filter_type=%s&filter_tree_path=%s", websiteID, pageNumber, url.QueryEscape(filterLocale), url.QueryEscape(filterType), url.QueryEscape(filterTreePath))
  req, _ := service.client.NewRequest("GET", url, nil)

  changes := new(HelpdeskHistoryChangeListData)
  resp, err := service.client.Do(req, changes)
  if err != nil {
    return nil, resp, err
  }

  return changes.Data, resp, err
}


// ResolveHelpdeskHistoryChange resolves a helpdesk content history change.
func (service *WebsiteService) ResolveHelpdeskHistoryChange(websiteID string, changeId string) (*HelpdeskHistoryChange, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/history/change/%s", websiteID, changeId)
  req, _ := service.client.NewRequest("GET", url, nil)

  change := new(HelpdeskHistoryChangeData)
  resp, err := service.client.Do(req, change)
  if err != nil {
    return nil, resp, err
  }

  return change.Data, resp, err
}


// CancelHelpdeskHistoryChange cancels a helpdesk content history change.
func (service *WebsiteService) CancelHelpdeskHistoryChange(websiteID string, changeId string, action string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/history/change/%s", websiteID, changeId)
  req, _ := service.client.NewRequest("DELETE", url, HelpdeskHistoryChangeCancel{Action: action})

  return service.client.Do(req, nil)
}


// RequestHelpdeskContentRefresh queues an asynchronous full refresh of helpdesk content.
func (service *WebsiteService) RequestHelpdeskContentRefresh(websiteID string, locale string, contentType string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/refresh/%s/%s", websiteID, locale, contentType)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// MapHelpdeskLocaleFeedbackRatings map locale feedback ratings for helpdesk in website.
func (service *WebsiteService) MapHelpdeskLocaleFeedbackRatings(websiteID string, locale string, contentType string, filterDateStart string, filterDateEnd string) (*HelpdeskLocaleFeedbackRatings, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/feedback/ratings/%s/%s?filter_date_start=%s&filter_date_end=%s", websiteID, locale, contentType, url.QueryEscape(filterDateStart), url.QueryEscape(filterDateEnd))
  req, _ := service.client.NewRequest("GET", url, nil)

  ratings := new(HelpdeskLocaleFeedbackRatingsData)
  resp, err := service.client.Do(req, ratings)
  if err != nil {
    return nil, resp, err
  }

  return ratings.Data, resp, err
}


// ListHelpdeskLocaleFeedbacks lists locale feedbacks for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleFeedbacks(websiteID string, locale string, contentType string, pageNumber uint, filterDateStart string, filterDateEnd string) (*[]HelpdeskLocaleFeedbackItem, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/feedback/list/%s/%s/%d?filter_date_start=%s&filter_date_end=%s", websiteID, locale, contentType, pageNumber, url.QueryEscape(filterDateStart), url.QueryEscape(filterDateEnd))
  req, _ := service.client.NewRequest("GET", url, nil)

  feedbacks := new(HelpdeskLocaleFeedbackListData)
  resp, err := service.client.Do(req, feedbacks)
  if err != nil {
    return nil, resp, err
  }

  return feedbacks.Data, resp, err
}


// ImportExternalHelpdeskToLocale imports a whole external helpdesk to Crisp, as a Crisp Helpdesk.
func (service *WebsiteService) ImportExternalHelpdeskToLocale(websiteID string, locale string, contentType string, helpdeskUrl string, detectLocales *bool, otherLocales *[]string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/import/%s/%s", websiteID, locale, contentType)
  req, _ := service.client.NewRequest("POST", url, HelpdeskLocaleExternalImport{HelpdeskURL: helpdeskUrl, DetectLocales: detectLocales, OtherLocales: otherLocales})

  return service.client.Do(req, nil)
}


// ExportHelpdeskLocaleArticles exports helpdesk articles for locale.
func (service *WebsiteService) ExportHelpdeskLocaleArticles(websiteID string, locale string, contentType string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/export/%s/%s", websiteID, locale, contentType)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskRedirections lists redirections for helpdesk in website.
func (service *WebsiteService) ListHelpdeskRedirections(websiteID string, pageNumber uint) (*[]HelpdeskRedirection, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/redirections/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  redirections := new(HelpdeskRedirectionListData)
  resp, err := service.client.Do(req, redirections)
  if err != nil {
    return nil, resp, err
  }

  return redirections.Data, resp, err
}


// AddHelpdeskRedirection adds a redirection for helpdesk in website.
func (service *WebsiteService) AddHelpdeskRedirection(websiteID string, redirectionPath string, redirectionTarget string) (*HelpdeskRedirectionNew, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/redirection", websiteID)
  req, _ := service.client.NewRequest("POST", url, HelpdeskRedirectionAdd{Path: redirectionPath, Target: redirectionTarget})

  redirectionNew := new(HelpdeskRedirectionNewData)
  resp, err := service.client.Do(req, redirectionNew)
  if err != nil {
    return nil, resp, err
  }

  return redirectionNew.Data, resp, err
}


// CheckHelpdeskRedirectionExists checks if a helpdesk redirection exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskRedirectionExists(websiteID string, redirectionId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/redirection/%s", websiteID, redirectionId)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskRedirection resolves a redirection for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskRedirection(websiteID string, redirectionId string) (*HelpdeskRedirection, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/redirection/%s", websiteID, redirectionId)
  req, _ := service.client.NewRequest("GET", url, nil)

  redirection := new(HelpdeskRedirectionData)
  resp, err := service.client.Do(req, redirection)
  if err != nil {
    return nil, resp, err
  }

  return redirection.Data, resp, err
}


// DeleteHelpdeskRedirection deletes a redirection for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskRedirection(websiteID string, redirectionId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/redirection/%s", websiteID, redirectionId)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskSettings resolves settings for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskSettings(websiteID string) (*HelpdeskSettings, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(HelpdeskSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// SaveHelpdeskSettings saves settings for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskSettings(websiteID string, settings HelpdeskSettings) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskDomain resolves domain for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskDomain(websiteID string) (*HelpdeskDomain, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/domain", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  domain := new(HelpdeskDomainData)
  resp, err := service.client.Do(req, domain)
  if err != nil {
    return nil, resp, err
  }

  return domain.Data, resp, err
}


// RequestHelpdeskDomainChange requests a change in the domain used for helpdesk.
func (service *WebsiteService) RequestHelpdeskDomainChange(websiteID string, basic *string, custom *string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/domain", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, HelpdeskDomainChangeRequest{Basic: basic, Custom: custom})

  return service.client.Do(req, nil)
}


// GenerateHelpdeskDomainSetupFlow retrieves the domain setup flow for helpdesk.
func (service *WebsiteService) GenerateHelpdeskDomainSetupFlow(websiteID string, custom string) (*HelpdeskDomainSetupFlow, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/domain/setup?custom=%s", websiteID, url.QueryEscape(custom))
  req, _ := service.client.NewRequest("GET", url, nil)

  flow := new(HelpdeskDomainSetupFlowData)
  resp, err := service.client.Do(req, flow)
  if err != nil {
    return nil, resp, err
  }

  return flow.Data, resp, err
}

