// Copyright 2017 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// BucketURLRequest mapping
type BucketURLRequest struct {
  Namespace   string                `json:"namespace,omitempty"`
  From        string                `json:"from,omitempty"`
  Identifier  string                `json:"identifier,omitempty"`
  ID          string                `json:"id,omitempty"`
  File        BucketURLRequestFile  `json:"file,omitempty"`
}

// BucketURLRequestFile mapping
type BucketURLRequestFile struct {
  Name  string  `json:"name,omitempty"`
  Type  string  `json:"type,omitempty"`
}


// GenerateBucketURL generates a bucket URL.
func (service *BucketService) GenerateBucketURL(bucketData BucketURLRequest) (*Response, error) {
  url := "bucket/url/generate"
  req, _ := service.client.NewRequest("POST", url, bucketData)

  return service.client.Do(req, nil)
}
