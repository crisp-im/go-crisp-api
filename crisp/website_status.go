// Copyright 2018 Crisp IM SARL All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteStatusData mapping
type WebsiteStatusData struct {
  Data  *WebsiteStatus  `json:"data,omitempty"`
}

// WebsiteStatus mapping
type WebsiteStatus struct {
  Name  *string  `json:"name,omitempty"`
  URL   *string  `json:"url,omitempty"`
}

// WebsiteStatusInitialize mapping
type WebsiteStatusInitialize struct {
  Name         string  `json:"name,omitempty"`
  DomainBasic  string  `json:"domain_basic,omitempty"`
}

// WebsiteStatusServiceData mapping
type WebsiteStatusServiceData struct {
  Data  *WebsiteStatusService  `json:"data,omitempty"`
}

// WebsiteStatusServicesData mapping
type WebsiteStatusServicesData struct {
  Data  *[]WebsiteStatusService  `json:"data,omitempty"`
}

// WebsiteStatusService mapping
type WebsiteStatusService struct {
  ServiceID  *string  `json:"service_id,omitempty"`
  Name       *string  `json:"name,omitempty"`
  Order      *uint16  `json:"order,omitempty"`
  CreatedAt  *uint64  `json:"created_at,omitempty"`
  UpdatedAt  *uint64  `json:"updated_at,omitempty"`
}

// WebsiteStatusServiceNewData mapping
type WebsiteStatusServiceNewData struct {
  Data  *WebsiteStatusServiceNew  `json:"data,omitempty"`
}

// WebsiteStatusServiceNew mapping
type WebsiteStatusServiceNew struct {
  ServiceID  *string  `json:"service_id,omitempty"`
}

// WebsiteStatusServiceNodeData mapping
type WebsiteStatusServiceNodeData struct {
  Data  *WebsiteStatusServiceNode  `json:"data,omitempty"`
}

// WebsiteStatusServiceNodesData mapping
type WebsiteStatusServiceNodesData struct {
  Data  *[]WebsiteStatusServiceNode  `json:"data,omitempty"`
}

// WebsiteStatusServiceNode mapping
type WebsiteStatusServiceNode struct {
  NodeID     *string                        `json:"node_id,omitempty"`
  Label      *string                        `json:"label,omitempty"`
  Mode       *string                        `json:"mode,omitempty"`
  Order      *uint16                        `json:"order,omitempty"`
  Replicas   *[]string                      `json:"replicas,omitempty"`
  HTTP       *WebsiteStatusServiceNodeHTTP  `json:"http,omitempty"`
  CreatedAt  *uint64                        `json:"created_at,omitempty"`
  UpdatedAt  *uint64                        `json:"updated_at,omitempty"`
}

// WebsiteStatusServiceNodeHTTP mapping
type WebsiteStatusServiceNodeHTTP struct {
  Status  *WebsiteStatusServiceNodeHTTPStatus  `json:"status,omitempty"`
  Body    *WebsiteStatusServiceNodeHTTPBody    `json:"body,omitempty"`
}

// WebsiteStatusServiceNodeHTTPStatus mapping
type WebsiteStatusServiceNodeHTTPStatus struct {
  HealthyAbove  *uint16  `json:"healthy_above,omitempty"`
  HealthyBelow  *uint16  `json:"healthy_below,omitempty"`
}

// WebsiteStatusServiceNodeHTTPBody mapping
type WebsiteStatusServiceNodeHTTPBody struct {
  HealthyMatch  *string  `json:"healthy_match,omitempty"`
}

// WebsiteStatusServiceNodeNewData mapping
type WebsiteStatusServiceNodeNewData struct {
  Data  *WebsiteStatusServiceNodeNew  `json:"data,omitempty"`
}

// WebsiteStatusServiceNodeNew mapping
type WebsiteStatusServiceNodeNew struct {
  NodeID  *string  `json:"node_id,omitempty"`
}

// WebsiteStatusServiceNodeReplicasData mapping
type WebsiteStatusServiceNodeReplicasData struct {
  Data  *[]WebsiteStatusServiceNodeReplica  `json:"data,omitempty"`
}

// WebsiteStatusServiceNodeReplica mapping
type WebsiteStatusServiceNodeReplica string;

// WebsiteStatusAnnouncementData mapping
type WebsiteStatusAnnouncementData struct {
  Data  *WebsiteStatusAnnouncement  `json:"data,omitempty"`
}

// WebsiteStatusAnnouncement mapping
type WebsiteStatusAnnouncement struct {
  Title      *string  `json:"title,omitempty"`
  Message    *string  `json:"message,omitempty"`
  CreatedAt  *uint64  `json:"created_at,omitempty"`
  UpdatedAt  *uint64  `json:"updated_at,omitempty"`
  ExpireAt   *uint64  `json:"expire_at,omitempty"`
}

// WebsiteStatusSettingsData mapping
type WebsiteStatusSettingsData struct {
  Data  *WebsiteStatusSettings  `json:"data,omitempty"`
}

// WebsiteStatusSettings mapping
type WebsiteStatusSettings struct {
  Name        *string                           `json:"name,omitempty"`
  Appearance  *WebsiteStatusSettingsAppearance  `json:"appearance,omitempty"`
  Behavior    *WebsiteStatusSettingsBehavior    `json:"behavior,omitempty"`
  Include     *WebsiteStatusSettingsInclude     `json:"include,omitempty"`
  Metrics     *WebsiteStatusSettingsMetrics     `json:"metrics,omitempty"`
  Notify      *WebsiteStatusSettingsNotify      `json:"notify,omitempty"`
}

// WebsiteStatusSettingsAppearance mapping
type WebsiteStatusSettingsAppearance struct {
  Logos   *WebsiteStatusSettingsAppearanceLogos  `json:"logos,omitempty"`
  Banner  *string                                `json:"banner,omitempty"`
}

// WebsiteStatusSettingsAppearanceLogos mapping
type WebsiteStatusSettingsAppearanceLogos struct {
  Header  *string  `json:"header,omitempty"`
  Footer  *string  `json:"footer,omitempty"`
}

// WebsiteStatusSettingsBehavior mapping
type WebsiteStatusSettingsBehavior struct {
  ShowChatbox     *bool  `json:"show_chatbox,omitempty"`
  LocalePicker    *bool  `json:"locale_picker,omitempty"`
  ReferLink       *bool  `json:"refer_link,omitempty"`
  ForbidIndexing  *bool  `json:"forbid_indexing,omitempty"`
}

// WebsiteStatusSettingsInclude mapping
type WebsiteStatusSettingsInclude struct {
  HTML  *string  `json:"html,omitempty"`
}

// WebsiteStatusSettingsMetrics mapping
type WebsiteStatusSettingsMetrics struct {
  Poll  *WebsiteStatusSettingsMetricsPoll  `json:"poll,omitempty"`
  Push  *WebsiteStatusSettingsMetricsPush  `json:"push,omitempty"`
}

// WebsiteStatusSettingsMetricsPoll mapping
type WebsiteStatusSettingsMetricsPoll struct {
  Interval   *uint16  `json:"interval,omitempty"`
  Retry      *uint8   `json:"retry,omitempty"`
  DelayDead  *uint16  `json:"delay_dead,omitempty"`
  DelaySick  *uint16  `json:"delay_sick,omitempty"`
}

// WebsiteStatusSettingsMetricsPush mapping
type WebsiteStatusSettingsMetricsPush struct {
  DelayDead           *uint16   `json:"delay_dead,omitempty"`
  SystemCPUSickAbove  *float32  `json:"system_cpu_sick_above,omitempty"`
  SystemRAMSickAbove  *float32  `json:"system_ram_sick_above,omitempty"`
}

// WebsiteStatusSettingsNotify mapping
type WebsiteStatusSettingsNotify struct {
  Slack  *WebsiteStatusSettingsNotifySlack  `json:"slack,omitempty"`
  Email  *WebsiteStatusSettingsNotifyEmail  `json:"email,omitempty"`
}

// WebsiteStatusSettingsNotifySlack mapping
type WebsiteStatusSettingsNotifySlack struct {
  HookURL  *string  `json:"hook_url,omitempty"`
}

// WebsiteStatusSettingsNotifyEmail mapping
type WebsiteStatusSettingsNotifyEmail struct {
  To  *string  `json:"to,omitempty"`
}

// WebsiteStatusDomainData mapping
type WebsiteStatusDomainData struct {
  Data  *WebsiteStatusDomain  `json:"data,omitempty"`
}

// WebsiteStatusDomain mapping
type WebsiteStatusDomain struct {
  Root    *string  `json:"root,omitempty"`
  Basic   *string  `json:"basic,omitempty"`
  Custom  *string  `json:"custom,omitempty"`
}

// WebsiteStatusDomainSetupFlowData mapping
type WebsiteStatusDomainSetupFlowData struct {
  Data  *WebsiteStatusDomainSetupFlow  `json:"data,omitempty"`
}

// WebsiteStatusDomainSetupFlow mapping
type WebsiteStatusDomainSetupFlow struct {
  Custom  *string                             `json:"custom,omitempty"`
  Setup   *WebsiteStatusDomainSetupFlowSetup  `json:"setup,omitempty"`
}

// WebsiteStatusDomainSetupFlowSetup mapping
type WebsiteStatusDomainSetupFlowSetup struct {
  Records  *[]WebsiteStatusDomainSetupFlowSetupRecord  `json:"records,omitempty"`
}

// WebsiteStatusDomainSetupFlowSetupRecord mapping
type WebsiteStatusDomainSetupFlowSetupRecord struct {
  Type   *string  `json:"type,omitempty"`
  Query  *string  `json:"query,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// WebsiteStatusReporterTokenData mapping
type WebsiteStatusReporterTokenData struct {
  Data  *WebsiteStatusReporterToken  `json:"data,omitempty"`
}

// WebsiteStatusReporterToken mapping
type WebsiteStatusReporterToken struct {
  Token  *string  `json:"token,omitempty"`
}

// WebsiteStatusServiceItem mapping
type WebsiteStatusServiceItem struct {
  Name  *string  `json:"name,omitempty"`
}

// WebsiteStatusServiceUpdate mapping
type WebsiteStatusServiceUpdate struct {
  Name   string  `json:"name,omitempty"`
  Order  uint16  `json:"order,omitempty"`
}

// WebsiteStatusServiceNodeItem mapping
type WebsiteStatusServiceNodeItem struct {
  Label  *string  `json:"label,omitempty"`
  Mode   *string  `json:"mode,omitempty"`
}

// WebsiteStatusServiceNodeUpdate mapping
type WebsiteStatusServiceNodeUpdate struct {
  Label     string                               `json:"label,omitempty"`
  Order     uint16                               `json:"order,omitempty"`
  Replicas  []string                             `json:"replicas,omitempty"`
  HTTP      *WebsiteStatusServiceNodeUpdateHTTP  `json:"http,omitempty"`
}

// WebsiteStatusServiceNodeUpdateHTTP mapping
type WebsiteStatusServiceNodeUpdateHTTP struct {
  Status  WebsiteStatusServiceNodeUpdateHTTPStatus  `json:"status,omitempty"`
  Body    WebsiteStatusServiceNodeUpdateHTTPBody    `json:"body,omitempty"`
}

// WebsiteStatusServiceNodeUpdateHTTPStatus mapping
type WebsiteStatusServiceNodeUpdateHTTPStatus struct {
  HealthyAbove  uint16  `json:"healthy_above,omitempty"`
  HealthyBelow  uint16  `json:"healthy_below,omitempty"`
}

// WebsiteStatusServiceNodeUpdateHTTPBody mapping
type WebsiteStatusServiceNodeUpdateHTTPBody struct {
  HealthyMatch  string  `json:"healthy_match,omitempty"`
}

// WebsiteStatusAnnouncementItem mapping
type WebsiteStatusAnnouncementItem struct {
  Title     string   `json:"title,omitempty"`
  Message   string   `json:"message,omitempty"`
  ExpireAt  *uint64  `json:"expire_at,omitempty"`
}

// WebsiteStatusSettingsUpdate mapping
type WebsiteStatusSettingsUpdate struct {
  Name        string                                 `json:"name,omitempty"`
  Appearance  WebsiteStatusSettingsUpdateAppearance  `json:"appearance,omitempty"`
  Behavior    WebsiteStatusSettingsUpdateBehavior    `json:"behavior,omitempty"`
  Include     WebsiteStatusSettingsUpdateInclude     `json:"include,omitempty"`
  Metrics     WebsiteStatusSettingsUpdateMetrics     `json:"metrics,omitempty"`
  Notify      WebsiteStatusSettingsUpdateNotify      `json:"notify,omitempty"`
}

// WebsiteStatusSettingsUpdateAppearance mapping
type WebsiteStatusSettingsUpdateAppearance struct {
  Logos   WebsiteStatusSettingsUpdateAppearanceLogos  `json:"logos,omitempty"`
  Banner  string                                      `json:"banner,omitempty"`
}

// WebsiteStatusSettingsUpdateAppearanceLogos mapping
type WebsiteStatusSettingsUpdateAppearanceLogos struct {
  Header  string  `json:"header,omitempty"`
  Footer  string  `json:"footer,omitempty"`
}

// WebsiteStatusSettingsUpdateBehavior mapping
type WebsiteStatusSettingsUpdateBehavior struct {
  ShowChatbox     bool  `json:"show_chatbox,omitempty"`
  LocalePicker    bool  `json:"locale_picker,omitempty"`
  ReferLink       bool  `json:"refer_link,omitempty"`
  ForbidIndexing  bool  `json:"forbid_indexing,omitempty"`
}

// WebsiteStatusSettingsUpdateInclude mapping
type WebsiteStatusSettingsUpdateInclude struct {
  HTML  string  `json:"html,omitempty"`
}

// WebsiteStatusSettingsUpdateMetrics mapping
type WebsiteStatusSettingsUpdateMetrics struct {
  Poll  WebsiteStatusSettingsUpdateMetricsPoll  `json:"poll,omitempty"`
  Push  WebsiteStatusSettingsUpdateMetricsPush  `json:"push,omitempty"`
}

// WebsiteStatusSettingsUpdateMetricsPoll mapping
type WebsiteStatusSettingsUpdateMetricsPoll struct {
  Interval   uint16  `json:"interval,omitempty"`
  Retry      uint8   `json:"retry,omitempty"`
  DelayDead  uint16  `json:"delay_dead,omitempty"`
  DelaySick  uint16  `json:"delay_sick,omitempty"`
}

// WebsiteStatusSettingsUpdateMetricsPush mapping
type WebsiteStatusSettingsUpdateMetricsPush struct {
  DelayDead           uint16   `json:"delay_dead,omitempty"`
  SystemCPUSickAbove  float32  `json:"system_cpu_sick_above,omitempty"`
  SystemRAMSickAbove  float32  `json:"system_ram_sick_above,omitempty"`
}

// WebsiteStatusSettingsUpdateNotify mapping
type WebsiteStatusSettingsUpdateNotify struct {
  Slack  *WebsiteStatusSettingsUpdateNotifySlack  `json:"slack,omitempty"`
  Email  *WebsiteStatusSettingsUpdateNotifyEmail  `json:"email,omitempty"`
}

// WebsiteStatusSettingsUpdateNotifySlack mapping
type WebsiteStatusSettingsUpdateNotifySlack struct {
  HookURL  string  `json:"hook_url,omitempty"`
}

// WebsiteStatusSettingsUpdateNotifyEmail mapping
type WebsiteStatusSettingsUpdateNotifyEmail struct {
  To  string  `json:"to,omitempty"`
}

// WebsiteStatusDomainChange mapping
type WebsiteStatusDomainChange struct {
  Basic   string  `json:"basic,omitempty"`
  Custom  string  `json:"custom,omitempty"`
}


// String returns the string representation of WebsiteStatus
func (instance WebsiteStatus) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusService
func (instance WebsiteStatusService) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusServiceNode
func (instance WebsiteStatusServiceNode) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusServiceNodeNew
func (instance WebsiteStatusServiceNodeNew) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusServiceNodeReplica
func (instance WebsiteStatusServiceNodeReplica) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusServiceNodeHTTP
func (instance WebsiteStatusServiceNodeHTTP) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusServiceNodeHTTPStatus
func (instance WebsiteStatusServiceNodeHTTPStatus) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusServiceNodeHTTPBody
func (instance WebsiteStatusServiceNodeHTTPBody) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusAnnouncement
func (instance WebsiteStatusAnnouncement) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettings
func (instance WebsiteStatusSettings) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsAppearance
func (instance WebsiteStatusSettingsAppearance) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsAppearanceLogos
func (instance WebsiteStatusSettingsAppearanceLogos) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsBehavior
func (instance WebsiteStatusSettingsBehavior) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsInclude
func (instance WebsiteStatusSettingsInclude) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsMetrics
func (instance WebsiteStatusSettingsMetrics) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsMetricsPoll
func (instance WebsiteStatusSettingsMetricsPoll) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsMetricsPush
func (instance WebsiteStatusSettingsMetricsPush) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsNotify
func (instance WebsiteStatusSettingsNotify) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsNotifySlack
func (instance WebsiteStatusSettingsNotifySlack) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusSettingsNotifyEmail
func (instance WebsiteStatusSettingsNotifyEmail) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusDomain
func (instance WebsiteStatusDomain) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusDomainSetupFlow
func (instance WebsiteStatusDomainSetupFlow) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusDomainSetupFlowSetup
func (instance WebsiteStatusDomainSetupFlowSetup) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusDomainSetupFlowSetupRecord
func (instance WebsiteStatusDomainSetupFlowSetupRecord) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteStatusReporterToken
func (instance WebsiteStatusReporterToken) String() string {
  return Stringify(instance)
}


// CheckIfStatusPageExists checks if status page exists for website
func (service *WebsiteService) CheckIfStatusPageExists(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status", websiteID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveStatusPage resolves status page information for website.
func (service *WebsiteService) ResolveStatusPage(websiteID string) (*WebsiteStatus, *Response, error) {
  url := fmt.Sprintf("website/%s/status", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  status := new(WebsiteStatusData)
  resp, err := service.client.Do(req, status)
  if err != nil {
    return nil, resp, err
  }

  return status.Data, resp, err
}


// InitializeStatusPage initializes a new status page for website.
func (service *WebsiteService) InitializeStatusPage(websiteID string, statusPage WebsiteStatusInitialize) (*Response, error) {
  url := fmt.Sprintf("website/%s/status", websiteID)
  req, _ := service.client.NewRequest("POST", url, statusPage)

  return service.client.Do(req, nil)
}


// DeleteStatusPage deletes status page for website.
func (service *WebsiteService) DeleteStatusPage(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListStatusPageServices lists services for status page in website.
func (service *WebsiteService) ListStatusPageServices(websiteID string, pageNumber uint) (*[]WebsiteStatusService, *Response, error) {
  url := fmt.Sprintf("website/%s/status/services/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  services := new(WebsiteStatusServicesData)
  resp, err := service.client.Do(req, services)
  if err != nil {
    return nil, resp, err
  }

  return services.Data, resp, err
}


// AddNewStatusPageService adds a new service for status page in website.
func (service *WebsiteService) AddNewStatusPageService(websiteID string, name string) (*WebsiteStatusServiceNew, *Response, error) {
  url := fmt.Sprintf("website/%s/status/service", websiteID)
  req, _ := service.client.NewRequest("POST", url, WebsiteStatusServiceItem{Name: &name})

  serviceNew := new(WebsiteStatusServiceNewData)
  resp, err := service.client.Do(req, serviceNew)
  if err != nil {
    return nil, resp, err
  }

  return serviceNew.Data, resp, err
}


// CheckStatusPageServiceExists checks if a service exists for status page in website.
func (service *WebsiteService) CheckStatusPageServiceExists(websiteID string, serviceID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s", websiteID, serviceID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveStatusPageService resolves a service for status page in website.
func (service *WebsiteService) ResolveStatusPageService(websiteID string, serviceID string) (*WebsiteStatusService, *Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s", websiteID, serviceID)
  req, _ := service.client.NewRequest("GET", url, nil)

  serviceData := new(WebsiteStatusServiceData)
  resp, err := service.client.Do(req, serviceData)
  if err != nil {
    return nil, resp, err
  }

  return serviceData.Data, resp, err
}


// SaveStatusPageService saves a service for status page in website.
func (service *WebsiteService) SaveStatusPageService(websiteID string, serviceID string, serviceData WebsiteStatusServiceUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s", websiteID, serviceID)
  req, _ := service.client.NewRequest("PUT", url, serviceData)

  return service.client.Do(req, nil)
}


// UpdateStatusPageService updates a service for status page in website.
func (service *WebsiteService) UpdateStatusPageService(websiteID string, serviceID string, serviceData WebsiteStatusServiceUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s", websiteID, serviceID)
  req, _ := service.client.NewRequest("PATCH", url, serviceData)

  return service.client.Do(req, nil)
}


// DeleteStatusPageService deletes a service for status page in website.
func (service *WebsiteService) DeleteStatusPageService(websiteID string, serviceID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s", websiteID, serviceID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListStatusPageServiceNodes lists nodes in service for status page in website.
func (service *WebsiteService) ListStatusPageServiceNodes(websiteID string, serviceID string, pageNumber uint) (*[]WebsiteStatusServiceNode, *Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/nodes/%d", websiteID, serviceID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  nodes := new(WebsiteStatusServiceNodesData)
  resp, err := service.client.Do(req, nodes)
  if err != nil {
    return nil, resp, err
  }

  return nodes.Data, resp, err
}


// AddNewStatusPageServiceNode adds a new node in service for status page in website.
func (service *WebsiteService) AddNewStatusPageServiceNode(websiteID string, serviceID string, label string, mode string) (*WebsiteStatusServiceNodeNew, *Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node", websiteID, serviceID)
  req, _ := service.client.NewRequest("POST", url, WebsiteStatusServiceNodeItem{Label: &label, Mode: &mode})

  nodeNew := new(WebsiteStatusServiceNodeNewData)
  resp, err := service.client.Do(req, nodeNew)
  if err != nil {
    return nil, resp, err
  }

  return nodeNew.Data, resp, err
}


// CheckStatusPageServiceNodeExists checks if a node in service exists for status page in website.
func (service *WebsiteService) CheckStatusPageServiceNodeExists(websiteID string, serviceID string, nodeID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveStatusPageServiceNode resolves a node in service for status page in website.
func (service *WebsiteService) ResolveStatusPageServiceNode(websiteID string, serviceID string, nodeID string) (*WebsiteStatusServiceNode, *Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("GET", url, nil)

  nodeData := new(WebsiteStatusServiceNodeData)
  resp, err := service.client.Do(req, nodeData)
  if err != nil {
    return nil, resp, err
  }

  return nodeData.Data, resp, err
}


// SaveStatusPageServiceNode saves a node in service for status page in website.
func (service *WebsiteService) SaveStatusPageServiceNode(websiteID string, serviceID string, nodeID string, node WebsiteStatusServiceNodeUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("PUT", url, node)

  return service.client.Do(req, nil)
}


// UpdateStatusPageServiceNode updates a node in service for status page in website.
func (service *WebsiteService) UpdateStatusPageServiceNode(websiteID string, serviceID string, nodeID string, node WebsiteStatusServiceNodeUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("PATCH", url, node)

  return service.client.Do(req, nil)
}


// DeleteStatusPageServiceNode deletes a node in service for status page in website.
func (service *WebsiteService) DeleteStatusPageServiceNode(websiteID string, serviceID string, nodeID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListStatusPageServiceNodeReplicas lists replicas health for node in service for status page in website.
func (service *WebsiteService) ListStatusPageServiceNodeReplicas(websiteID string, serviceID string, nodeID string) (*[]WebsiteStatusServiceNodeReplica, *Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s/replicas", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("GET", url, nil)

  replicas := new(WebsiteStatusServiceNodeReplicasData)
  resp, err := service.client.Do(req, replicas)
  if err != nil {
    return nil, resp, err
  }

  return replicas.Data, resp, err
}


// FlushStatusPageServiceNodeReplicas flushes the list of replicas health for node in service for status page in website.
func (service *WebsiteService) FlushStatusPageServiceNodeReplicas(websiteID string, serviceID string, nodeID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/service/%s/node/%s/replicas", websiteID, serviceID, nodeID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// CheckIfStatusPageAnnouncementExists checks if an announcement exists for status page in website.
func (service *WebsiteService) CheckIfStatusPageAnnouncementExists(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/announcement", websiteID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveStatusPageAnnouncement resolves an announcement for status page in website.
func (service *WebsiteService) ResolveStatusPageAnnouncement(websiteID string) (*WebsiteStatusAnnouncement, *Response, error) {
  url := fmt.Sprintf("website/%s/status/announcement", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  announcement := new(WebsiteStatusAnnouncementData)
  resp, err := service.client.Do(req, announcement)
  if err != nil {
    return nil, resp, err
  }

  return announcement.Data, resp, err
}


// SaveStatusPageAnnouncement saves an announcement for status page in website.
func (service *WebsiteService) SaveStatusPageAnnouncement(websiteID string, announcement WebsiteStatusAnnouncementItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/announcement", websiteID)
  req, _ := service.client.NewRequest("PUT", url, announcement)

  return service.client.Do(req, nil)
}


// DeleteStatusPageAnnouncement deletes an announcement for status page in website.
func (service *WebsiteService) DeleteStatusPageAnnouncement(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/announcement", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ResolveStatusPageSettings resolves settings for status page in website.
func (service *WebsiteService) ResolveStatusPageSettings(websiteID string) (*WebsiteStatusSettings, *Response, error) {
  url := fmt.Sprintf("website/%s/status/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(WebsiteStatusSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// SaveStatusPageSettings saves settings for status page in website.
func (service *WebsiteService) SaveStatusPageSettings(websiteID string, settings WebsiteStatusSettingsUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}


// ResolveStatusPageDomain resolves domain for status page in website.
func (service *WebsiteService) ResolveStatusPageDomain(websiteID string) (*WebsiteStatusDomain, *Response, error) {
  url := fmt.Sprintf("website/%s/status/domain", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  domain := new(WebsiteStatusDomainData)
  resp, err := service.client.Do(req, domain)
  if err != nil {
    return nil, resp, err
  }

  return domain.Data, resp, err
}


// RequestStatusPageDomainChange requests a change in the domain used for status page.
func (service *WebsiteService) RequestStatusPageDomainChange(websiteID string, domain WebsiteStatusDomainChange) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/domain", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, domain)

  return service.client.Do(req, nil)
}


// GenerateStatusPageDomainSetupFlow retrieves the domain setup flow for status page.
func (service *WebsiteService) GenerateStatusPageDomainSetupFlow(websiteID string, custom string) (*WebsiteStatusDomainSetupFlow, *Response, error) {
  url := fmt.Sprintf("website/%s/status/domain/setup?custom=%s", websiteID, url.QueryEscape(custom))
  req, _ := service.client.NewRequest("GET", url, nil)

  domainSetupFlow := new(WebsiteStatusDomainSetupFlowData)
  resp, err := service.client.Do(req, domainSetupFlow)
  if err != nil {
    return nil, resp, err
  }

  return domainSetupFlow.Data, resp, err
}


// GetStatusPageReporterToken resolves status page reporter token, that can be used in Crisp Status Report libraries.
func (service *WebsiteService) GetStatusPageReporterToken(websiteID string) (*WebsiteStatusReporterToken, *Response, error) {
  url := fmt.Sprintf("website/%s/status/reporter/token", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  token := new(WebsiteStatusReporterTokenData)
  resp, err := service.client.Do(req, token)
  if err != nil {
    return nil, resp, err
  }

  return token.Data, resp, err
}


// RollStatusPageReporterToken rolls status page reporter token.
func (service *WebsiteService) RollStatusPageReporterToken(websiteID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/status/reporter/token", websiteID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}
