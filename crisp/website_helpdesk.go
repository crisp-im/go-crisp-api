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
  LocaleID    *string  `json:"locale_id,omitempty"`
  Locale      *string  `json:"locale,omitempty"`
  URL         *string  `json:"url,omitempty"`
  Articles    *uint16  `json:"articles,omitempty"`
  Categories  *uint16  `json:"categories,omitempty"`
}

// HelpdeskLocaleArticleListData mapping
type HelpdeskLocaleArticleListData struct {
  Data  *[]HelpdeskLocaleArticle  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleData mapping
type HelpdeskLocaleArticleData struct {
  Data  *HelpdeskLocaleArticle  `json:"data,omitempty"`
}

// HelpdeskLocaleArticle mapping
type HelpdeskLocaleArticle struct {
  ArticleID    *string  `json:"article_id,omitempty"`
  Title        *string  `json:"title,omitempty"`
  Description  *string  `json:"description,omitempty"`
  Content      *string  `json:"content,omitempty"`
  Status       *string  `json:"status,omitempty"`
  Visibility   *string  `json:"visibility,omitempty"`
  Visits       *uint32  `json:"visits,omitempty"`
  Order        *uint16  `json:"order,omitempty"`
  URL          *string  `json:"url,omitempty"`
  CreatedAt    *uint64  `json:"created_at,omitempty"`
  UpdatedAt    *uint64  `json:"updated_at,omitempty"`
  PublishedAt  *uint64  `json:"published_at,omitempty"`
}

// HelpdeskLocaleArticleNewData mapping
type HelpdeskLocaleArticleNewData struct {
  Data  *HelpdeskLocaleArticleNew  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleNew mapping
type HelpdeskLocaleArticleNew struct {
  ArticleID  *string  `json:"article_id,omitempty"`
}

// HelpdeskLocaleArticleCategoryListData mapping
type HelpdeskLocaleArticleCategoryListData struct {
  Data  *[]HelpdeskLocaleArticleCategory  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleCategoryData mapping
type HelpdeskLocaleArticleCategoryData struct {
  Data  *HelpdeskLocaleArticleCategory  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleCategory mapping
type HelpdeskLocaleArticleCategory struct {
  CategoryID  *string  `json:"category_id,omitempty"`
  SectionID   *string  `json:"section_id,omitempty"`
}

// HelpdeskLocaleArticleCategoryNewData mapping
type HelpdeskLocaleArticleCategoryNewData struct {
  Data  *HelpdeskLocaleArticleCategoryNew  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleCategoryNew mapping
type HelpdeskLocaleArticleCategoryNew struct {
  CategoryID  *string  `json:"category_id,omitempty"`
}

// HelpdeskLocaleArticleAlternateListData mapping
type HelpdeskLocaleArticleAlternateListData struct {
  Data  *[]HelpdeskLocaleArticleAlternate  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleAlternateData mapping
type HelpdeskLocaleArticleAlternateData struct {
  Data  *HelpdeskLocaleArticleAlternate  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleAlternate mapping
type HelpdeskLocaleArticleAlternate struct {
  Locale     *string  `json:"locale,omitempty"`
  ArticleID  *string  `json:"article_id,omitempty"`
}

// HelpdeskLocaleArticlePublishData mapping
type HelpdeskLocaleArticlePublishData struct {
  Data  *HelpdeskLocaleArticlePublish  `json:"data,omitempty"`
}

// HelpdeskLocaleArticlePublish mapping
type HelpdeskLocaleArticlePublish struct {
  URL  *string  `json:"url,omitempty"`
}

// HelpdeskLocaleSectionListData mapping
type HelpdeskLocaleSectionListData struct {
  Data  *[]HelpdeskLocaleSection  `json:"data,omitempty"`
}

// HelpdeskLocaleSectionData mapping
type HelpdeskLocaleSectionData struct {
  Data  *HelpdeskLocaleSection  `json:"data,omitempty"`
}

// HelpdeskLocaleSection mapping
type HelpdeskLocaleSection struct {
  SectionID  *string  `json:"section_id,omitempty"`
  Name       *string  `json:"name,omitempty"`
  Order      *uint16  `json:"order,omitempty"`
  CreatedAt  *uint64  `json:"created_at,omitempty"`
  UpdatedAt  *uint64  `json:"updated_at,omitempty"`
}

// HelpdeskLocaleSectionNewData mapping
type HelpdeskLocaleSectionNewData struct {
  Data  *HelpdeskLocaleSectionNew  `json:"data,omitempty"`
}

// HelpdeskLocaleSectionNew mapping
type HelpdeskLocaleSectionNew struct {
  SectionID  *string  `json:"section_id,omitempty"`
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
  Data  *HelpdeskLocaleFeedbackItem  `json:"data,omitempty"`
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
  Name        *string                      `json:"name,omitempty"`
  Appearance  *HelpdeskSettingsAppearance  `json:"appearance,omitempty"`
  Behavior    *HelpdeskSettingsBehavior    `json:"behavior,omitempty"`
  Include     *HelpdeskSettingsInclude     `json:"include,omitempty"`
  Access      *HelpdeskSettingsAccess      `json:"access,omitempty"`
}

// HelpdeskSettingsAppearance mapping
type HelpdeskSettingsAppearance struct {
  Logos   *HelpdeskSettingsAppearanceLogos  `json:"logos,omitempty"`
  Banner  *string                           `json:"banner,omitempty"`
}

// HelpdeskSettingsAppearanceLogos mapping
type HelpdeskSettingsAppearanceLogos struct {
  Header  *string  `json:"header,omitempty"`
  Footer  *string  `json:"footer,omitempty"`
}

// HelpdeskSettingsBehavior mapping
type HelpdeskSettingsBehavior struct {
  FrequentlyRead      *bool  `json:"frequently_read,omitempty"`
  ShowCategoryImages  *bool  `json:"show_category_images,omitempty"`
  ShowChatbox         *bool  `json:"show_chatbox,omitempty"`
  AskFeedback         *bool  `json:"ask_feedback,omitempty"`
  LocalePicker        *bool  `json:"locale_picker,omitempty"`
  ReferLink           *bool  `json:"refer_link,omitempty"`
  ForbidIndexing      *bool  `json:"forbid_indexing,omitempty"`
  StatusHealthDead    *bool  `json:"status_health_dead,omitempty"`
}

// HelpdeskSettingsInclude mapping
type HelpdeskSettingsInclude struct {
  HTML  *string  `json:"html,omitempty"`
}

// HelpdeskSettingsAccess mapping
type HelpdeskSettingsAccess struct {
  Password  *string  `json:"password,omitempty"`
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
  Verify  *string  `json:"verify,omitempty"`
}

// HelpdeskLocaleAdd mapping
type HelpdeskLocaleAdd struct {
  Locale  string  `json:"locale"`
}

// HelpdeskLocaleArticleAdd mapping
type HelpdeskLocaleArticleAdd struct {
  Title  string  `json:"title"`
}

// HelpdeskLocaleArticleCategoryUpdate mapping
type HelpdeskLocaleArticleCategoryUpdate struct {
  CategoryID  string   `json:"category_id"`
  SectionID   *string  `json:"section_id,omitempty"`
}

// HelpdeskLocaleArticleAlternateSave mapping
type HelpdeskLocaleArticleAlternateSave struct {
  ArticleID  string  `json:"article_id"`
}

// HelpdeskLocaleCategoryAdd mapping
type HelpdeskLocaleCategoryAdd struct {
  Title  string  `json:"title"`
}

// HelpdeskLocaleSectionAdd mapping
type HelpdeskLocaleSectionAdd struct {
  Name  string  `json:"name"`
}

// HelpdeskLocaleExternalImport mapping
type HelpdeskLocaleExternalImport struct {
  HelpdeskURL  string  `json:"helpdesk_url"`
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


// String returns the string representation of HelpdeskLocaleArticle
func (instance HelpdeskLocaleArticle) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleArticleNew
func (instance HelpdeskLocaleArticleNew) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleArticleCategory
func (instance HelpdeskLocaleArticleCategory) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleArticleCategoryNew
func (instance HelpdeskLocaleArticleCategoryNew) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleArticleAlternate
func (instance HelpdeskLocaleArticleAlternate) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleArticlePublish
func (instance HelpdeskLocaleArticlePublish) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleSection
func (instance HelpdeskLocaleSection) String() string {
  return Stringify(instance)
}


// String returns the string representation of HelpdeskLocaleSectionNew
func (instance HelpdeskLocaleSectionNew) String() string {
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
func (service *WebsiteService) DeleteHelpdesk(websiteID string, verify string) (*Response, error) {
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

  locale := new(HelpdeskLocaleData)
  resp, err := service.client.Do(req, locale)
  if err != nil {
    return nil, resp, err
  }

  return locale.Data, resp, err
}


// DeleteHelpdeskLocale deletes a locale for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskLocaleArticles lists articles for a helpdesk locale in website.
func (service *WebsiteService) ListHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticle, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/articles/%d", websiteID, locale, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  articles := new(HelpdeskLocaleArticleListData)
  resp, err := service.client.Do(req, articles)
  if err != nil {
    return nil, resp, err
  }

  return articles.Data, resp, err
}


// AddNewHelpdeskLocaleArticle adds a new locale article for helpdesk in website.
func (service *WebsiteService) AddNewHelpdeskLocaleArticle(websiteID string, locale string, title string) (*HelpdeskLocaleArticleNew, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article", websiteID, locale)
  req, _ := service.client.NewRequest("POST", url, HelpdeskLocaleArticleAdd{Title: title})

  articleNew := new(HelpdeskLocaleArticleNewData)
  resp, err := service.client.Do(req, articleNew)
  if err != nil {
    return nil, resp, err
  }

  return articleNew.Data, resp, err
}


// CheckHelpdeskLocaleArticleExists checks if a locale article exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleArticleExists(websiteID string, locale string, articleId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleArticle resolves a locale article for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticle, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("GET", url, nil)

  article := new(HelpdeskLocaleArticleData)
  resp, err := service.client.Do(req, article)
  if err != nil {
    return nil, resp, err
  }

  return article.Data, resp, err
}


// SaveHelpdeskLocaleArticle saves a locale article for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("PUT", url, article)

  return service.client.Do(req, nil)
}


// UpdateHelpdeskLocaleArticle updates a locale article for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("PATCH", url, article)

  return service.client.Do(req, nil)
}


// DeleteHelpdeskLocaleArticle deletes a locale article for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleArticleCategory resolves a locale article category for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/category", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("GET", url, nil)

  category := new(HelpdeskLocaleArticleCategoryData)
  resp, err := service.client.Do(req, category)
  if err != nil {
    return nil, resp, err
  }

  return category.Data, resp, err
}


// UpdateHelpdeskLocaleArticleCategory updates a locale article category for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string, categoryId string, sectionId *string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/category", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("PATCH", url, HelpdeskLocaleArticleCategoryUpdate{CategoryID: categoryId, SectionID: sectionID})

  return service.client.Do(req, nil)
}


// ListHelpdeskLocaleArticleAlternates lists alternate locales on a locale article for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleArticleAlternates(websiteID string, locale string, articleId string) (*[]HelpdeskLocaleArticleAlternate, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/alternates", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("GET", url, nil)

  alternates := new(HelpdeskLocaleArticleAlternateListData)
  resp, err := service.client.Do(req, alternates)
  if err != nil {
    return nil, resp, err
  }

  return alternates.Data, resp, err
}


// CheckHelpdeskLocaleArticleAlternateExists checks if alternate locale exists on a locale article for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleArticleAlternateExists(websiteID string, locale string, articleId string, localeLinked string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/alternate/%s", websiteID, locale, articleId, localeLinked)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleArticleAlternate resolves alternate locale on a locale article for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*HelpdeskLocaleArticleAlternate, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/alternate/%s", websiteID, locale, articleId, localeLinked)
  req, _ := service.client.NewRequest("GET", url, nil)

  alternate := new(HelpdeskLocaleArticleAlternateData)
  resp, err := service.client.Do(req, alternate)
  if err != nil {
    return nil, resp, err
  }

  return alternate.Data, resp, err
}


// SaveHelpdeskLocaleArticleAlternate saves alternate locale on a locale article for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string, articleIdLinked string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/alternate/%s", websiteID, locale, articleId, localeLinked)
  req, _ := service.client.NewRequest("PUT", url, HelpdeskLocaleArticleAlternateSave{ArticleID: articleIdLinked})

  return service.client.Do(req, nil)
}


// DeleteHelpdeskLocaleArticleAlternate deletes alternate locale on a locale article for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/alternate/%s", websiteID, locale, articleId, localeLinked)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// PublishHelpdeskLocaleArticle publishes a locale article for helpdesk in website.
func (service *WebsiteService) PublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticlePublish, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/publish", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("POST", url, nil)

  articlePublish := new(HelpdeskLocaleArticlePublishData)
  resp, err := service.client.Do(req, articlePublish)
  if err != nil {
    return nil, resp, err
  }

  return articlePublish.Data, resp, err
}


// UnpublishHelpdeskLocaleArticle unpublishes a locale article for helpdesk in website.
func (service *WebsiteService) UnpublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/unpublish", websiteID, locale, articleId)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskLocaleCategories lists locale categories for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/categories/%d", websiteID, locale, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  categories := new(HelpdeskLocaleArticleCategoryListData)
  resp, err := service.client.Do(req, categories)
  if err != nil {
    return nil, resp, err
  }

  return categories.Data, resp, err
}


// AddHelpdeskLocaleCategory adds a locale category for helpdesk in website.
func (service *WebsiteService) AddHelpdeskLocaleCategory(websiteID string, locale string, name string) (*HelpdeskLocaleArticleCategoryNew, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category", websiteID, locale)
  req, _ := service.client.NewRequest("POST", url, HelpdeskLocaleCategoryAdd{Name: name})

  categoryNew := new(HelpdeskLocaleArticleCategoryNewData)
  resp, err := service.client.Do(req, categoryNew)
  if err != nil {
    return nil, resp, err
  }

  return categoryNew.Data, resp, err
}


// CheckHelpdeskLocaleCategoryExists checks if a locale category exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleCategoryExists(websiteID string, locale string, categoryId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryId)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleCategory resolves a locale category for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*HelpdeskLocaleArticleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryId)
  req, _ := service.client.NewRequest("GET", url, nil)

  category := new(HelpdeskLocaleArticleCategoryData)
  resp, err := service.client.Do(req, category)
  if err != nil {
    return nil, resp, err
  }

  return category.Data, resp, err
}


// SaveHelpdeskLocaleCategory saves a locale category for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryId)
  req, _ := service.client.NewRequest("PUT", url, category)

  return service.client.Do(req, nil)
}


// UpdateHelpdeskLocaleCategory updates a locale category for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryId)
  req, _ := service.client.NewRequest("PATCH", url, category)

  return service.client.Do(req, nil)
}


// DeleteHelpdeskLocaleCategory deletes a locale category for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryId)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskLocaleSections lists locale sections for helpdesk in website and category.
func (service *WebsiteService) ListHelpdeskLocaleSections(websiteID string, locale string, categoryId string, pageNumber uint) (*[]HelpdeskLocaleSection, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/sections/%d", websiteID, locale, categoryId, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  sections := new(HelpdeskLocaleSectionListData)
  resp, err := service.client.Do(req, sections)
  if err != nil {
    return nil, resp, err
  }

  return sections.Data, resp, err
}


// AddHelpdeskLocaleSection adds a locale section for helpdesk in website and category.
func (service *WebsiteService) AddHelpdeskLocaleSection(websiteID string, locale string, categoryId string, name string) (*HelpdeskLocaleSectionNew, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/section", websiteID, locale, categoryId)
  req, _ := service.client.NewRequest("POST", url, HelpdeskLocaleSectionAdd{Name: name})

  sectionNew := new(HelpdeskLocaleSectionNewData)
  resp, err := service.client.Do(req, sectionNew)
  if err != nil {
    return nil, resp, err
  }

  return sectionNew.Data, resp, err
}


// CheckHelpdeskLocaleSectionExists checks if a locale section exists for helpdesk in website and category.
func (service *WebsiteService) CheckHelpdeskLocaleSectionExists(websiteID string, locale string, categoryId string, sectionId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/section/%s", websiteID, locale, categoryId, sectionId)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleSection resolves a locale section for helpdesk in website and category.
func (service *WebsiteService) ResolveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*HelpdeskLocaleSection, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/section/%s", websiteID, locale, categoryId, sectionId)
  req, _ := service.client.NewRequest("GET", url, nil)

  section := new(HelpdeskLocaleSectionData)
  resp, err := service.client.Do(req, section)
  if err != nil {
    return nil, resp, err
  }

  return section.Data, resp, err
}


// SaveHelpdeskLocaleSection saves a locale section for helpdesk in website and category.
func (service *WebsiteService) SaveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/section/%s", websiteID, locale, categoryId, sectionId)
  req, _ := service.client.NewRequest("PUT", url, section)

  return service.client.Do(req, nil)
}


// UpdateHelpdeskLocaleSection updates a locale section for helpdesk in website and category.
func (service *WebsiteService) UpdateHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/section/%s", websiteID, locale, categoryId, sectionId)
  req, _ := service.client.NewRequest("PATCH", url, section)

  return service.client.Do(req, nil)
}


// DeleteHelpdeskLocaleSection deletes a locale section for helpdesk in website and category.
func (service *WebsiteService) DeleteHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s/section/%s", websiteID, locale, categoryId, sectionId)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// MapHelpdeskLocaleFeedbackRatings map locale feedback ratings for helpdesk in website.
func (service *WebsiteService) MapHelpdeskLocaleFeedbackRatings(websiteID string, locale string, filterDateStart string, filterDateEnd string) (*HelpdeskLocaleFeedbackRatings, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/feedback/ratings?filter_date_start=%s&filter_date_end=%s", websiteID, locale, url.QueryEscape(filterDateStart), url.QueryEscape(filterDateEnd))
  req, _ := service.client.NewRequest("GET", url, nil)

  ratings := new(HelpdeskLocaleFeedbackRatingsData)
  resp, err := service.client.Do(req, ratings)
  if err != nil {
    return nil, resp, err
  }

  return ratings.Data, resp, err
}


// ListHelpdeskLocaleFeedbacks lists locale feedbacks for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleFeedbacks(websiteID string, locale string, pageNumber uint, filterDateStart string, filterDateEnd string) (*[]HelpdeskLocaleFeedbackItem, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/feedback/list/%d?filter_date_start=%s&filter_date_end=%s", websiteID, locale, pageNumber, url.QueryEscape(filterDateStart), url.QueryEscape(filterDateEnd))
  req, _ := service.client.NewRequest("GET", url, nil)

  feedbacks := new(HelpdeskLocaleFeedbackListData)
  resp, err := service.client.Do(req, feedbacks)
  if err != nil {
    return nil, resp, err
  }

  return feedbacks.Data, resp, err
}


// ImportExternalHelpdeskToLocale imports a whole external helpdesk to Crisp, as a Crisp Helpdesk.
func (service *WebsiteService) ImportExternalHelpdeskToLocale(websiteID string, locale string, helpdeskUrl string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/import", websiteID, locale)
  req, _ := service.client.NewRequest("POST", url, HelpdeskLocaleExternalImport{HelpdeskURL: helpdeskUrl})

  return service.client.Do(req, nil)
}


// ExportHelpdeskLocaleArticles exports helpdesk articles for locale.
func (service *WebsiteService) ExportHelpdeskLocaleArticles(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/export", websiteID, locale)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskRedirections lists redirections for helpdesk in website.
func (service *WebsiteService) ListHelpdeskRedirections(websiteID string, pageNumber uint) (*[]HelpdeskRedirection, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/redirections/%d", websiteID, locale, pageNumber)
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

