// Copyright 2023 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
	"errors"
	// TODO
  // "fmt"
  // "net/http"
  // "net/url"
)


// HelpdeskData mapping
type HelpdeskData struct {
  Data  *Helpdesk  `json:"data,omitempty"`
}

// Helpdesk mapping
type Helpdesk struct {
	/* TODO */
}

// HelpdeskLocaleData mapping
type HelpdeskLocaleData struct {
  Data  *HelpdeskLocale  `json:"data,omitempty"`
}

// HelpdeskLocale mapping
type HelpdeskLocale struct {
	/* TODO */
}

// HelpdeskLocaleArticleData mapping
type HelpdeskLocaleArticleData struct {
  Data  *HelpdeskLocaleArticle  `json:"data,omitempty"`
}

// HelpdeskLocaleArticle mapping
type HelpdeskLocaleArticle struct {
	/* TODO */
}

// HelpdeskLocaleArticleNewData mapping
type HelpdeskLocaleArticleNewData struct {
  Data  *HelpdeskLocaleArticleNew  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleNew mapping
type HelpdeskLocaleArticleNew struct {
	/* TODO */
}

// HelpdeskLocaleArticleCategoryData mapping
type HelpdeskLocaleArticleCategoryData struct {
  Data  *HelpdeskLocaleArticleCategory  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleCategory mapping
type HelpdeskLocaleArticleCategory struct {
	/* TODO */
}

// HelpdeskLocaleArticleCategoryNewData mapping
type HelpdeskLocaleArticleCategoryNewData struct {
  Data  *HelpdeskLocaleArticleCategoryNew  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleCategoryNew mapping
type HelpdeskLocaleArticleCategoryNew struct {
	/* TODO */
}

// HelpdeskLocaleArticleAlternateData mapping
type HelpdeskLocaleArticleAlternateData struct {
  Data  *HelpdeskLocaleArticleAlternate  `json:"data,omitempty"`
}

// HelpdeskLocaleArticleAlternate mapping
type HelpdeskLocaleArticleAlternate struct {
	/* TODO */
}

// HelpdeskLocaleArticlePublishData mapping
type HelpdeskLocaleArticlePublishData struct {
  Data  *HelpdeskLocaleArticlePublish  `json:"data,omitempty"`
}

// HelpdeskLocaleArticlePublish mapping
type HelpdeskLocaleArticlePublish struct {
	/* TODO */
}

// HelpdeskLocaleSectionData mapping
type HelpdeskLocaleSectionData struct {
  Data  *HelpdeskLocaleSection  `json:"data,omitempty"`
}

// HelpdeskLocaleSection mapping
type HelpdeskLocaleSection struct {
	/* TODO */
}

// HelpdeskLocaleSectionNewData mapping
type HelpdeskLocaleSectionNewData struct {
  Data  *HelpdeskLocaleSectionNew  `json:"data,omitempty"`
}

// HelpdeskLocaleSectionNew mapping
type HelpdeskLocaleSectionNew struct {
	/* TODO */
}

// HelpdeskLocaleFeedbackRatingsData mapping
type HelpdeskLocaleFeedbackRatingsData struct {
  Data  *HelpdeskLocaleFeedbackRatings  `json:"data,omitempty"`
}

// HelpdeskLocaleFeedbackRatings mapping
type HelpdeskLocaleFeedbackRatings struct {
	/* TODO */
}

// HelpdeskLocaleFeedbackItemData mapping
type HelpdeskLocaleFeedbackItemData struct {
  Data  *HelpdeskLocaleFeedbackItem  `json:"data,omitempty"`
}

// HelpdeskLocaleFeedbackItem mapping
type HelpdeskLocaleFeedbackItem struct {
	/* TODO */
}

// HelpdeskRedirectionData mapping
type HelpdeskRedirectionData struct {
  Data  *HelpdeskRedirection  `json:"data,omitempty"`
}

// HelpdeskRedirection mapping
type HelpdeskRedirection struct {
	/* TODO */
}

// HelpdeskRedirectionNewData mapping
type HelpdeskRedirectionNewData struct {
  Data  *HelpdeskRedirectionNew  `json:"data,omitempty"`
}

// HelpdeskRedirectionNew mapping
type HelpdeskRedirectionNew struct {
	/* TODO */
}

// HelpdeskSettingsData mapping
type HelpdeskSettingsData struct {
  Data  *HelpdeskSettings  `json:"data,omitempty"`
}

// HelpdeskSettings mapping
type HelpdeskSettings struct {
	/* TODO */
}

// HelpdeskDomainData mapping
type HelpdeskDomainData struct {
  Data  *HelpdeskDomain  `json:"data,omitempty"`
}

// HelpdeskDomain mapping
type HelpdeskDomain struct {
	/* TODO */
}

// HelpdeskDomainSetupFlowData mapping
type HelpdeskDomainSetupFlowData struct {
  Data  *HelpdeskDomainSetupFlow  `json:"data,omitempty"`
}

// HelpdeskDomainSetupFlow mapping
type HelpdeskDomainSetupFlow struct {
	/* TODO */
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
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdesk resolves helpdesk information for website.
func (service *WebsiteService) ResolveHelpdesk(websiteID string) (*Helpdesk, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// InitializeHelpdesk initializes a new helpdesk for website.
func (service *WebsiteService) InitializeHelpdesk(websiteID string, name string, domainBasic string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdesk deletes helpdesk for website.
func (service *WebsiteService) DeleteHelpdesk(websiteID string, verify string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskLocales lists locales for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocales(websiteID string, pageNumber uint) (*[]HelpdeskLocale, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// AddHelpdeskLocale adds a locale for helpdesk in website.
func (service *WebsiteService) AddHelpdeskLocale(websiteID string, locale string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// CheckHelpdeskLocaleExists checks if a helpdesk locale exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleExists(websiteID string, locale string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskLocale resolves a locale for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocale(websiteID string, locale string) (*HelpdeskLocale, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdeskLocale deletes a locale for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskLocaleArticles lists articles for a helpdesk locale in website.
func (service *WebsiteService) ListHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticle, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// AddNewHelpdeskLocaleArticle adds a new locale article for helpdesk in website.
func (service *WebsiteService) AddNewHelpdeskLocaleArticle(websiteID string, locale string, title string) (*HelpdeskLocaleArticleNew, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// CheckHelpdeskLocaleArticleExists checks if a locale article exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleArticleExists(websiteID string, locale string, articleId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskLocaleArticle resolves a locale article for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticle, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// SaveHelpdeskLocaleArticle saves a locale article for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// UpdateHelpdeskLocaleArticle updates a locale article for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdeskLocaleArticle deletes a locale article for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskLocaleArticleCategory resolves a locale article category for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticleCategory, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// UpdateHelpdeskLocaleArticleCategory updates a locale article category for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string, categoryId string, sectionId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskLocaleArticleAlternates lists alternate locales on a locale article for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleArticleAlternates(websiteID string, locale string, articleId string) (*[]HelpdeskLocaleArticleAlternate, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// CheckHelpdeskLocaleArticleAlternateExists checks if alternate locale exists on a locale article for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleArticleAlternateExists(websiteID string, locale string, articleId string, localeLinked string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskLocaleArticleAlternate resolves alternate locale on a locale article for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*HelpdeskLocaleArticleAlternate, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// SaveHelpdeskLocaleArticleAlternate saves alternate locale on a locale article for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string, articleIdLinked string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdeskLocaleArticleAlternate deletes alternate locale on a locale article for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// PublishHelpdeskLocaleArticle publishes a locale article for helpdesk in website.
func (service *WebsiteService) PublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticlePublish, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// UnpublishHelpdeskLocaleArticle unpublishes a locale article for helpdesk in website.
func (service *WebsiteService) UnpublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskLocaleCategories lists locale categories for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticleCategory, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// AddHelpdeskLocaleCategory adds a locale category for helpdesk in website.
func (service *WebsiteService) AddHelpdeskLocaleCategory(websiteID string, locale string) (*HelpdeskLocaleArticleCategoryNew, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// CheckHelpdeskLocaleCategoryExists checks if a locale category exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskLocaleCategoryExists(websiteID string, locale string, categoryId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskLocaleCategory resolves a locale category for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*HelpdeskLocaleArticleCategory, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// SaveHelpdeskLocaleCategory saves a locale category for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// UpdateHelpdeskLocaleCategory updates a locale category for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdeskLocaleCategory deletes a locale category for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskLocaleSections lists locale sections for helpdesk in website and category.
func (service *WebsiteService) ListHelpdeskLocaleSections(websiteID string, locale string, categoryId string, pageNumber uint) (*[]HelpdeskLocaleSection, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// AddHelpdeskLocaleSection adds a locale section for helpdesk in website and category.
func (service *WebsiteService) AddHelpdeskLocaleSection(websiteID string, locale string, categoryId string, name string) (*HelpdeskLocaleSectionNew, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// CheckHelpdeskLocaleSectionExists checks if a locale section exists for helpdesk in website and category.
func (service *WebsiteService) CheckHelpdeskLocaleSectionExists(websiteID string, locale string, categoryId string, sectionId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskLocaleSection resolves a locale section for helpdesk in website and category.
func (service *WebsiteService) ResolveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*HelpdeskLocaleSection, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// SaveHelpdeskLocaleSection saves a locale section for helpdesk in website and category.
func (service *WebsiteService) SaveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// UpdateHelpdeskLocaleSection updates a locale section for helpdesk in website and category.
func (service *WebsiteService) UpdateHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdeskLocaleSection deletes a locale section for helpdesk in website and category.
func (service *WebsiteService) DeleteHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// MapHelpdeskLocaleFeedbackRatings map locale feedback ratings for helpdesk in website.
func (service *WebsiteService) MapHelpdeskLocaleFeedbackRatings(websiteID string, locale string, filterDateStart string, filterDateEnd string) (*HelpdeskLocaleFeedbackRatings, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskLocaleFeedbacks lists locale feedbacks for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleFeedbacks(websiteID string, locale string, pageNumber uint, filterDateStart string, filterDateEnd string) (*[]HelpdeskLocaleFeedbackItem, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ImportExternalHelpdeskToLocale imports a whole external helpdesk to Crisp, as a Crisp Helpdesk.
func (service *WebsiteService) ImportExternalHelpdeskToLocale(websiteID string, locale string, helpdeskUrl string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ExportHelpdeskLocaleArticles exports helpdesk articles for locale.
func (service *WebsiteService) ExportHelpdeskLocaleArticles(websiteID string, locale string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ListHelpdeskRedirections lists redirections for helpdesk in website.
func (service *WebsiteService) ListHelpdeskRedirections(websiteID string, pageNumber uint) (*[]HelpdeskRedirection, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// AddHelpdeskRedirection adds a redirection for helpdesk in website.
func (service *WebsiteService) AddHelpdeskRedirection(websiteID string, redirectionPath string, redirectionTarget string) (*HelpdeskRedirectionNew, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// CheckHelpdeskRedirectionExists checks if a helpdesk redirection exists for helpdesk in website.
func (service *WebsiteService) CheckHelpdeskRedirectionExists(websiteID string, redirectionId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskRedirection resolves a redirection for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskRedirection(websiteID string, redirectionId string) (*HelpdeskRedirection, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// DeleteHelpdeskRedirection deletes a redirection for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskRedirection(websiteID string, redirectionId string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskSettings resolves settings for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskSettings(websiteID string) (*HelpdeskSettings, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// SaveHelpdeskSettings saves settings for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskSettings(websiteID string, settings HelpdeskSettings) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// ResolveHelpdeskDomain resolves domain for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskDomain(websiteID string) (*HelpdeskDomain, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// RequestHelpdeskDomainChange requests a change in the domain used for helpdesk.
func (service *WebsiteService) RequestHelpdeskDomainChange(websiteID string, basic string, custom string) (*Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}


// GenerateHelpdeskDomainSetupFlow retrieves the domain setup flow for helpdesk.
func (service *WebsiteService) GenerateHelpdeskDomainSetupFlow(websiteID string, custom string) (*HelpdeskDomainSetupFlow, *Response, error) {
	// TODO
	return errors.New("TODO: not implemented")
}

