// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "time"
  "net/url"
)


// WebsiteAnalyticsPointsData mapping
type WebsiteAnalyticsPointsData struct {
  Data  *WebsiteAnalyticsPoints  `json:"data,omitempty"`
}

// WebsiteAnalyticsPoints mapping
type WebsiteAnalyticsPoints struct {
  Pipeline  *WebsiteAnalyticsPointsPipeline  `json:"pipeline,omitempty"`
  Points    *[]WebsiteAnalyticsPointsPoint   `json:"points,omitempty"`
}

// WebsiteAnalyticsPointsPipeline mapping
type WebsiteAnalyticsPointsPipeline struct {
  Aggregator  *string  `json:"aggregator,omitempty"`
}

// WebsiteAnalyticsPointsPoint mapping
type WebsiteAnalyticsPointsPoint struct {
  Value  *uint32                           `json:"value,omitempty"`
  Hits   *uint64                           `json:"hits,omitempty"`
  Date   *WebsiteAnalyticsPointsPointDate  `json:"date,omitempty"`
}

// WebsiteAnalyticsPointsPointDate mapping
type WebsiteAnalyticsPointsPointDate struct {
  From  *uint64  `json:"from,omitempty"`
  To    *uint64  `json:"to,omitempty"`
}

// WebsiteAnalyticsFilterData mapping
type WebsiteAnalyticsFilterData struct {
  Data  *[]WebsiteAnalyticsFilter  `json:"data,omitempty"`
}

// WebsiteAnalyticsFilter mapping
type WebsiteAnalyticsFilter struct {
  Primary    *string  `json:"primary,omitempty"`
  Secondary  *string  `json:"secondary,omitempty"`
  Tertiary   *string  `json:"tertiary,omitempty"`
}


// String returns the string representation of WebsiteAnalyticsPoints
func (instance WebsiteAnalyticsPoints) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteAnalyticsFilter
func (instance WebsiteAnalyticsFilter) String() string {
  return Stringify(instance)
}


// AcquireAnalyticsPoints acquires analytics points for given type and metric in website.
func (service *WebsiteService) AcquireAnalyticsPoints(websiteID string, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time, dateSplit string, classifier string, filterPrimary string, filterSecondary string, filterTertiary string) (*WebsiteAnalyticsPoints, *Response, error) {
  dateFromFormat, err := dateFrom.UTC().MarshalText()
  if err != nil {
    return nil, nil, err
  }

  dateToFormat, err := dateTo.UTC().MarshalText()
  if err != nil {
    return nil, nil, err
  }

  url := fmt.Sprintf("website/%s/analytics/%s/%s/points?date_from=%s&date_to=%s&date_split=%s&classifier=%s&filter_primary=%s&filter_secondary=%s&filter_tertiary=%s", websiteID, pointType, pointMetric, url.QueryEscape(string(dateFromFormat[:])), url.QueryEscape(string(dateToFormat[:])), url.QueryEscape(dateSplit), url.QueryEscape(classifier), url.QueryEscape(filterPrimary), url.QueryEscape(filterSecondary), url.QueryEscape(filterTertiary))

  req, _ := service.client.NewRequest("GET", url, nil)

  points := new(WebsiteAnalyticsPointsData)
  resp, err := service.client.Do(req, points)
  if err != nil {
    return nil, resp, err
  }

  return points.Data, resp, err
}


// ListAnalyticsFilters acquires available analytics filters in website.
func (service *WebsiteService) ListAnalyticsFilters(websiteID string, pointType string, pointMetric string, pageNumber uint) (*[]WebsiteAnalyticsFilter, *Response, error) {
  url := fmt.Sprintf("website/%s/analytics/%s/%s/filters/%d", websiteID, pointType, pointMetric, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  filters := new(WebsiteAnalyticsFilterData)
  resp, err := service.client.Do(req, filters)
  if err != nil {
    return nil, resp, err
  }

  return filters.Data, resp, err
}
