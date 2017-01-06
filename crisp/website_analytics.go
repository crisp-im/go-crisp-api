// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "time"
  "net/url"
)


// WebsiteAnalyticsData mapping
type WebsiteAnalyticsData struct {
  Data  *WebsiteAnalytics  `json:"data,omitempty"`
}

// WebsiteAnalytics mapping
type WebsiteAnalytics struct {
  Metrics  *[]WebsiteAnalyticsMetric  `json:"metrics,omitempty"`
}

// WebsiteAnalyticsMetric mapping
type WebsiteAnalyticsMetric struct {
  Type    *string                         `json:"type,omitempty"`
  Points  *[]WebsiteAnalyticsMetricPoint  `json:"points,omitempty"`
}

// WebsiteAnalyticsMetricPoint mapping
type WebsiteAnalyticsMetricPoint struct {
  Value      *int     `json:"value,omitempty"`
  Unit       *string  `json:"unit,omitempty"`
  Timestamp  *uint    `json:"timestamp,omitempty"`
}


// String returns the string representation of WebsiteAnalytics
func (instance WebsiteAnalytics) String() string {
  return Stringify(instance)
}


// AcquireChatsAnalytics acquires analytics on chats in website.
func (service *WebsiteService) AcquireChatsAnalytics(websiteID string, filterMetric string, filterOperator string, filterDateStart time.Time, filterDateEnd time.Time) (*WebsiteAnalytics, *Response, error) {
  filterDateStartFormat, err := filterDateStart.UTC().MarshalText()
  if err != nil {
    return nil, nil, err
  }

  filterDateEndFormat, err := filterDateEnd.UTC().MarshalText()
  if err != nil {
    return nil, nil, err
  }

  url := fmt.Sprintf("website/%s/analytics/chats/?filter_metric=%s&filter_operator=%s&filter_date_start=%s&filter_date_end=%s", websiteID, url.QueryEscape(filterMetric), url.QueryEscape(filterOperator), url.QueryEscape(string(filterDateStartFormat[:])), url.QueryEscape(string(filterDateEndFormat[:])))

  req, _ := service.client.NewRequest("GET", url, nil)

  analytics := new(WebsiteAnalyticsData)
  resp, err := service.client.Do(req, analytics)
  if err != nil {
    return nil, resp, err
  }

  return analytics.Data, resp, err
}
