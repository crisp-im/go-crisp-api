// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// BillingMethodAllData mapping
type BillingMethodAllData struct {
  Data  *[]BillingMethodAll  `json:"data"`
}

// BillingMethodAll mapping
type BillingMethodAll struct {
  UserID  *string         `json:"user_id"`
  Cards   *BillingMethod  `json:"cards"`
}

// BillingMethodData mapping
type BillingMethodData struct {
  Data  *BillingMethod  `json:"data"`
}

// BillingMethod mapping
type BillingMethod struct {
  CardID             *string                  `json:"card_id"`
  NameOnCard         *string                  `json:"name_on_card"`
  Address            *string                  `json:"address"`
  CardNumberPreview  *string                  `json:"card_number_preview"`
  CardProvider       *string                  `json:"card_provider"`
  ExpirationDate     *string                  `json:"expiration_date"`
  AddedDate          *string                  `json:"added_date"`
  IsExpired          *bool                    `json:"is_expired"`
  LinkedWebsites     *[]BillingMethodWebsite  `json:"linked_websites"`
}

// BillingMethodWebsite mapping
type BillingMethodWebsite struct {
  WebsiteID  *string  `json:"website_id"`
  Name       *string  `json:"name"`
  Domain     *string  `json:"domain"`
}

// BillingMethodCreate mapping
type BillingMethodCreate struct {
  NameOnCard      *string                             `json:"name_on_card"`
  Address         *string                             `json:"address"`
  CardNumber      *string                             `json:"card_number"`
  SecurityCode    *string                             `json:"security_code"`
  ExpirationDate  *BillingMethodCreateExpirationDate  `json:"expiration_date"`
}

// BillingMethodCreateExpirationDate mapping
type BillingMethodCreateExpirationDate struct {
  Month  *string  `json:"month"`
  Year   *string  `json:"year"`
}

// BillingMethodInvoiceAllData mapping
type BillingMethodInvoiceAllData struct {
  Data  *BillingMethodInvoiceAll  `json:"data"`
}

// BillingMethodInvoiceAll mapping
type BillingMethodInvoiceAll struct {
  Results  *[]BillingMethodInvoice      `json:"results"`
  Paging   *BillingMethodInvoicePaging  `json:"paging"`
}

// BillingMethodInvoiceData mapping
type BillingMethodInvoiceData struct {
  Data  *BillingMethodInvoice  `json:"data"`
}

// BillingMethodInvoice mapping
type BillingMethodInvoice struct {
  InvoiceID   *string                       `json:"invoice_id"`
  BillID      *string                       `json:"bill_id"`
  Date        *string                       `json:"date"`
  Payment     *BillingMethodInvoicePayment  `json:"payment"`
  Invoice     *BillingMethodInvoiceContent  `json:"invoice"`
}

// BillingMethodInvoicePayment mapping
type BillingMethodInvoicePayment struct {
  IsPaid         *bool  `json:"is_paid"`
  NumberRetries  *uint  `json:"number_retries"`
}

// BillingMethodInvoiceContent mapping
type BillingMethodInvoiceContent struct {
  ID        *string                          `json:"id"`
  Currency  *string                          `json:"currency"`
  Due       *BillingMethodInvoiceContentDue  `json:"due"`
}

// BillingMethodInvoiceContentDue mapping
type BillingMethodInvoiceContentDue struct {
  Total   *uint                                  `json:"total"`
  Parts   *[]BillingMethodInvoiceContentDuePart  `json:"parts"`
}

// BillingMethodInvoiceContentDuePart mapping
type BillingMethodInvoiceContentDuePart struct {
  ID        *string                                      `json:"id"`
  Name      *string                                      `json:"name"`
  Domain    *string                                      `json:"domain"`
  Plugins   *[]BillingMethodInvoiceContentDuePartPlugin  `json:"plugins"`
}

// BillingMethodInvoiceContentDuePartPlugin mapping
type BillingMethodInvoiceContentDuePartPlugin struct {
  ID     *string  `json:"id"`
  Name   *string  `json:"name"`
  Price  *uint    `json:"price"`
}

// BillingMethodInvoicePaging mapping
type BillingMethodInvoicePaging struct {
  Range    *uint  `json:"range"`
  PerPage  *uint  `json:"per_page"`
  Total    *uint  `json:"total"`
}

// BillingMethodLink mapping
type BillingMethodLink struct {
  Service  *string  `json:"service"`
}


// ListAllBillingMethods resolves the current user account information.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-base-get
func (service *UserService) ListAllBillingMethods() (*[]BillingMethodAll, *Response, error) {
  url := "user/account/account"
  req, _ := service.client.NewRequest("GET", url, nil)

  billing := new(BillingMethodAllData)
  resp, err := service.client.Do(req, billing)
  if err != nil {
    return nil, resp, err
  }

  return billing.Data, resp, err
}


// AddNewBillingMethod adds a payment method (credit card) to the user account.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-billing-post
func (service *UserService) AddNewBillingMethod(nameOnCard string, address string, cardNumber string, securityCode string, expirationDateMonth string, expirationDateYear string) (*Response, error) {
  url := "user/account/billing"
  req, _ := service.client.NewRequest("POST", url, BillingMethodCreate{NameOnCard: &nameOnCard, Address: &address, CardNumber: &cardNumber, SecurityCode: &securityCode, ExpirationDate: &BillingMethodCreateExpirationDate{Month: &expirationDateMonth, Year: &expirationDateYear}})

  return service.client.Do(req, nil)
}


// GetBillingMethod acquires information about a saved billing method (eg: credit card).
// Reference: https://docs.crisp.im/api/v1/#user-user-account-billing-get-1
func (service *UserService) GetBillingMethod(cardID string) (*BillingMethod, *Response, error) {
  url := fmt.Sprintf("user/account/billing/%s", cardID)
  req, _ := service.client.NewRequest("GET", url, nil)

  billing := new(BillingMethodData)
  resp, err := service.client.Do(req, billing)
  if err != nil {
    return nil, resp, err
  }

  return billing.Data, resp, err
}


// RemoveBillingMethod deletes a saved billing method.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-billing-delete
func (service *UserService) RemoveBillingMethod(cardID string) (*Response, error) {
  url := fmt.Sprintf("user/account/billing/%s", cardID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListInvoicesForBillingMethod lists saved invoices for billing method.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-billing-get-2
func (service *UserService) ListInvoicesForBillingMethod(cardID string, pageNumber int) (*BillingMethodInvoiceAll, *Response, error) {
  url := fmt.Sprintf("user/account/billing/%s/invoices/%d", cardID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  invoicing := new(BillingMethodInvoiceAllData)
  resp, err := service.client.Do(req, invoicing)
  if err != nil {
    return nil, resp, err
  }

  return invoicing.Data, resp, err
}


// GetInvoiceForBillingMethod get given saved invoice for billing method.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-billing-get-3
func (service *UserService) GetInvoiceForBillingMethod(cardID string, invoiceID string) (*BillingMethodInvoice, *Response, error) {
  url := fmt.Sprintf("user/account/billing/%s/invoice/%s", cardID, invoiceID)
  req, _ := service.client.NewRequest("GET", url, nil)

  invoicing := new(BillingMethodInvoiceData)
  resp, err := service.client.Do(req, invoicing)
  if err != nil {
    return nil, resp, err
  }

  return invoicing.Data, resp, err
}


// LinkToExternalBillingMethod links to an external billing method. Used to for services which need an external approval (eg: PayPal), and that cannot be added directly via a simple form submit.
// Reference: https://docs.crisp.im/api/v1/#user-user-account-billing-post-1
func (service *UserService) LinkToExternalBillingMethod(billingService string) (*Response, error) {
  url := "user/account/billing/link"
  req, _ := service.client.NewRequest("POST", url, BillingMethodLink{Service: &billingService})

  return service.client.Do(req, nil)
}
