// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// RequestWebsiteStates requests website states to be pushed to clients connected to realtime socket.
func (service *WebsiteService) RequestWebsiteStates(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/states", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  resp, err := service.client.Do(req, nil)
  return resp, err
}
