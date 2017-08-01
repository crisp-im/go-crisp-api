// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// ListAnimationMedias lists animation medias.
func (service *WebsiteService) ListAnimationMedias(pageNumber uint, listID string, searchQuery string) (*Response, error) {
  url := fmt.Sprintf("media/animation/list/%d?list_id=%s&search_query=%s", pageNumber, url.QueryEscape(listID), url.QueryEscape(searchQuery))
  req, _ := service.client.NewRequest("GET", url, nil)

  return service.client.Do(req, nil)
}
