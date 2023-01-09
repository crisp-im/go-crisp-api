// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


// BucketURLRequest mapping
type BucketURLRequest struct {
  Namespace   string                     `json:"namespace,omitempty"`
  ID          string                     `json:"id,omitempty"`
  File        BucketURLRequestFile       `json:"file,omitempty"`
  Resource    *BucketURLRequestResource  `json:"resource,omitempty"`
}

// BucketURLRequestFile mapping
type BucketURLRequestFile struct {
  Name  string  `json:"name,omitempty"`
  Type  string  `json:"type,omitempty"`
}

// BucketURLRequestResource mapping
type BucketURLRequestResource struct {
  Type        string  `json:"type,omitempty"`
  Identifier  string  `json:"identifier,omitempty"`
}


// GenerateBucketURL generates a bucket URL.
func (service *BucketService) GenerateBucketURL(bucketData BucketURLRequest) (*Response, error) {
  url := "bucket/url/generate"
  req, _ := service.client.NewRequest("POST", url, bucketData)

  return service.client.Do(req, nil)
}
