// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// WebsiteService service
type WebsiteService service


// WebsiteFilter mapping
type WebsiteFilter struct {
  Model      *string         `json:"model,omitempty"`
  Criterion  *string         `json:"criterion,omitempty"`
  Operator   *string         `json:"operator,omitempty"`
  Query      *[]interface{}  `json:"query,omitempty"`
}
