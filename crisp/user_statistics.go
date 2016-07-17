// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// UserStatisticsData mapping
type UserStatisticsData struct {
  Data  *UserStatistics  `json:"data"`
}

// UserStatistics mapping
type UserStatistics struct {
  UserID  *string  `json:"user_id"`
  Unread  *int     `json:"unread"`
}


// CountTotalUnreadMessages counts the total number of unread messages, cross-website.
// Reference: https://docs.crisp.im/api/v1/#user-user-statistics-get
func (service *UserService) CountTotalUnreadMessages() (*UserStatistics, *Response, error) {
  url := "user/stats/unread"
  req, _ := service.client.NewRequest("GET", url, nil)

  stats := new(UserStatisticsData)
  resp, err := service.client.Do(req, stats)
  if err != nil {
    return nil, resp, err
  }

  return stats.Data, resp, err
}
