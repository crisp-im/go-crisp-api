// Copyright 2024 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// AnalyticsGenerateQuery mapping
type AnalyticsGenerateQuery struct {
  Metric      string                       `json:"metric"`
  Type        string                       `json:"type"`
  Aggregator  *string                      `json:"aggregator,omitempty"`
  SplitBy     *string                      `json:"split_by,omitempty"`
  Filter      *interface{}                 `json:"filter,omitempty"`
  Date        AnalyticsGenerateQueryDate   `json:"date"`
  If          *AnalyticsGenerateQueryIf    `json:"if,omitempty"`
  Then        *AnalyticsGenerateQueryThen  `json:"then,omitempty"`
  RawExport   *bool                        `json:"raw_export,omitempty"`
}

// AnalyticsGenerateQueryDate mapping
type AnalyticsGenerateQueryDate struct {
  From      string                           `json:"from"`
  To        string                           `json:"to"`
  Split     string                           `json:"split"`
  Timezone  string                           `json:"timezone"`
  Days      *AnalyticsGenerateQueryDateDays  `json:"days,omitempty"`
}

// AnalyticsGenerateQueryDateDays mapping
type AnalyticsGenerateQueryDateDays struct {
  Monday     *AnalyticsGenerateQueryDateDay  `json:"monday,omitempty"`
  Tuesday    *AnalyticsGenerateQueryDateDay  `json:"tuesday,omitempty"`
  Wednesday  *AnalyticsGenerateQueryDateDay  `json:"wednesday,omitempty"`
  Thursday   *AnalyticsGenerateQueryDateDay  `json:"thursday,omitempty"`
  Friday     *AnalyticsGenerateQueryDateDay  `json:"friday,omitempty"`
  Saturday   *AnalyticsGenerateQueryDateDay  `json:"saturday,omitempty"`
  Sunday     *AnalyticsGenerateQueryDateDay  `json:"sunday,omitempty"`
}

// AnalyticsGenerateQueryDateDay mapping
type AnalyticsGenerateQueryDateDay struct {
  HourFrom  string  `json:"hour_from"`
  HourTo    string  `json:"hour_to"`
}

// AnalyticsGenerateQueryIf mapping
type AnalyticsGenerateQueryIf struct {
  Metric   string  `json:"metric"`
  SplitBy  string  `json:"split_by"`
}

// AnalyticsGenerateQueryThen mapping
type AnalyticsGenerateQueryThen struct {
  Type    *string       `json:"type,omitempty"`
  Filter  *interface{}  `json:"filter,omitempty"`
}

// WebsiteAnalyticsGenerateData mapping
type WebsiteAnalyticsGenerateData struct {
  Data  *[]WebsiteAnalyticsGeneratePoint  `json:"data,omitempty"`
}

// WebsiteAnalyticsGeneratePoint mapping
type WebsiteAnalyticsGeneratePoint struct {
  Period      *string  `json:"period,omitempty"`
  Value       *uint32  `json:"value,omitempty"`
  SplitBy     *string  `json:"split_by,omitempty"`
  Hits        *uint32  `json:"hits,omitempty"`
  UniqueHits  *uint32  `json:"unique_hits,omitempty"`
  SumValue    *uint32  `json:"sum_value,omitempty"`
  SumHits     *uint32  `json:"sum_hits,omitempty"`
}


// String returns the string representation of WebsiteAnalyticsGeneratePoint
func (instance WebsiteAnalyticsGeneratePoint) String() string {
  return Stringify(instance)
}


// GenerateAnalytics generates analytics for given type and metric in website.
func (service *WebsiteService) GenerateAnalytics(websiteID string, query AnalyticsGenerateQuery) (*[]WebsiteAnalyticsGeneratePoint, *Response, error) {
  url := fmt.Sprintf("website/%s/analytics/generate", websiteID)
  req, _ := service.client.NewRequest("POST", url, query)

  points := new(WebsiteAnalyticsGenerateData)
  resp, err := service.client.Do(req, points)
  if err != nil {
    return nil, resp, err
  }

  return points.Data, resp, err
}
