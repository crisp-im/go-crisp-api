// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteServiceTranslateItem mapping
type WebsiteServiceTranslateItem struct {
  ID      *string                             `json:"id,omitempty"`
  Locale  *WebsiteServiceTranslateItemLocale  `json:"locale,omitempty"`
  Text    *string                             `json:"text,omitempty"`
}

// WebsiteServiceTranslateItemLocale mapping
type WebsiteServiceTranslateItemLocale struct {
  From  *string  `json:"from,omitempty"`
  To    *string  `json:"to,omitempty"`
}


// RequestTranslationService requests a translation service.
func (service *WebsiteService) RequestTranslationService(websiteID string, translate WebsiteServiceTranslateItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/service/translate", websiteID)
  req, _ := service.client.NewRequest("POST", url, translate)

  return service.client.Do(req, nil)
}
