// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteContractsData mapping
type WebsiteContractsData struct {
  Data  *[]WebsiteContract  `json:"data,omitempty"`
}

// WebsiteContractData mapping
type WebsiteContractData struct {
  Data  *WebsiteContract  `json:"data,omitempty"`
}

// WebsiteContract mapping
type WebsiteContract struct {
  ID            *string  `json:"id,omitempty"`
  Title         *string  `json:"title,omitempty"`
  ContractURL   *string  `json:"contract_url,omitempty"`
  AgreementURL  *string  `json:"agreement_url,omitempty"`
  Description   *string  `json:"description,omitempty"`
}

// WebsiteContractSave mapping
type WebsiteContractSave struct {
  AgreementURL  *string  `json:"agreement_url,omitempty"`
}


// String returns the string representation of WebsiteContract
func (instance WebsiteContract) String() string {
  return Stringify(instance)
}


// ListWebsiteContracts lists all contracts agreed or not for website.
func (service *WebsiteService) ListWebsiteContracts(websiteID string) (*[]WebsiteContract, *Response, error) {
  url := fmt.Sprintf("website/%s/contracts", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  contracts := new(WebsiteContractsData)
  resp, err := service.client.Do(req, contracts)
  if err != nil {
    return nil, resp, err
  }

  return contracts.Data, resp, err
}


// CheckIfAgreedWebsiteContractExists checks if an agreed contract exists for website.
func (service *WebsiteService) CheckIfAgreedWebsiteContractExists(websiteID string, contractID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/contract/%s", websiteID, contractID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveAgreedWebsiteContract resolves an agreed contract for website.
func (service *WebsiteService) ResolveAgreedWebsiteContract(websiteID string, contractID string) (*WebsiteContract, *Response, error) {
  url := fmt.Sprintf("website/%s/contract/%s", websiteID, contractID)
  req, _ := service.client.NewRequest("GET", url, nil)

  contract := new(WebsiteContractData)
  resp, err := service.client.Do(req, contract)
  if err != nil {
    return nil, resp, err
  }

  return contract.Data, resp, err
}


// AgreeToWebsiteContract agrees to a contract for website.
func (service *WebsiteService) AgreeToWebsiteContract(websiteID string, contractID string, agreementURL string) (*Response, error) {
  url := fmt.Sprintf("website/%s/contract/%s", websiteID, contractID)
  req, _ := service.client.NewRequest("PUT", url, WebsiteContractSave{AgreementURL: &agreementURL})

  return service.client.Do(req, nil)
}


// DeleteAgreedWebsiteContract deletes an agreed contract for website.
func (service *WebsiteService) DeleteAgreedWebsiteContract(websiteID string, contractID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/contract/%s", websiteID, contractID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}
