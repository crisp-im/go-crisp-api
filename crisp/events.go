// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "github.com/graarh/golang-socketio"
  "github.com/graarh/golang-socketio/transport"
  "time"
  "net"
  "net/url"
  "strconv"
)


const (
  reconnectWait = 4
)

// EventsService service
type EventsService service

// EventsRegister stores event handlers
type EventsRegister struct {
  Handlers  map[string]*caller
}

type eventsSendAuthentication struct {
  Tier      string  `json:"tier"`
  Username  string  `json:"username"`
  Password  string  `json:"password"`
}

// EventsReceiveAuthenticationUnauthorized maps unauthorized
type EventsReceiveAuthenticationUnauthorized struct {
  Message  *string  `json:"message"`
}

// EventsReceiveSessionUpdateAvailability maps session:update_availability
type EventsReceiveSessionUpdateAvailability struct {
  WebsiteID     *string  `json:"website_id"`
  SessionID     *string  `json:"session_id"`
  Availability  *string  `json:"availability"`
}

// EventsReceiveSessionRequestInitiated maps session:request:initiated
type EventsReceiveSessionRequestInitiated struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
}

// EventsReceiveSessionSetEmail maps session:set_email
type EventsReceiveSessionSetEmail struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  Email      *string  `json:"email"`
}

// EventsReceiveSessionSetAvatar maps session:set_avatar
type EventsReceiveSessionSetAvatar struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  Avatar     *string  `json:"avatar"`
}

// EventsReceiveSessionSetCover maps session:set_cover
type EventsReceiveSessionSetCover struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  Cover      *string  `json:"cover"`
}

// EventsReceiveSessionSetNickname maps session:set_nickname
type EventsReceiveSessionSetNickname struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  Nickname   *string  `json:"nickname"`
}

// EventsReceiveSessionSyncPages maps session:sync:pages
type EventsReceiveSessionSyncPages struct {
  WebsiteID  *string                              `json:"website_id"`
  SessionID  *string                              `json:"session_id"`
  Pages      *[]EventsReceiveSessionSyncPagesOne  `json:"pages"`
}

// EventsReceiveSessionSyncPagesOne maps session:sync:pages/pages
type EventsReceiveSessionSyncPagesOne struct {
  PageTitle  *string  `json:"page_title"`
  PageURL    *string  `json:"page_url"`
  Timestamp  *uint    `json:"timestamp"`
}

// EventsReceiveSessionSyncGeolocation maps session:sync:geolocation
type EventsReceiveSessionSyncGeolocation struct {
  WebsiteID    *string                                   `json:"website_id"`
  SessionID    *string                                   `json:"session_id"`
  Geolocation  *EventsReceiveSessionSyncGeolocationData  `json:"geolocation"`
}

// EventsReceiveSessionSyncGeolocationData maps session:sync:geolocation/geolocation
type EventsReceiveSessionSyncGeolocationData struct {
  Metro    *int        `json:"metro,omitempty"`
  LL       *[]float32  `json:"ll,omitempty"`
  City     *string     `json:"city,omitempty"`
  Region   *string     `json:"region,omitempty"`
  Country  *string     `json:"country,omitempty"`
  Range    *[]uint     `json:"range,omitempty"`
}

// EventsReceiveSessionSyncSystem maps session:sync:system
type EventsReceiveSessionSyncSystem struct {
  WebsiteID    *string                              `json:"website_id"`
  SessionID    *string                              `json:"session_id"`
  System       *EventsReceiveSessionSyncSystemData  `json:"system"`
}

// EventsReceiveSessionSyncSystemData maps session:sync:system/system
type EventsReceiveSessionSyncSystemData struct {
  OS       *EventsReceiveSessionSyncSystemDataOS       `json:"os,omitempty"`
  Engine   *EventsReceiveSessionSyncSystemDataEngine   `json:"engine,omitempty"`
  Browser  *EventsReceiveSessionSyncSystemDataBrowser  `json:"browser,omitempty"`
  UA       *string                                             `json:"ua,omitempty"`
}

// EventsReceiveSessionSyncSystemDataOS maps session:sync:system/system/os
type EventsReceiveSessionSyncSystemDataOS struct {
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// EventsReceiveSessionSyncSystemDataEngine maps session:sync:system/system/engine
type EventsReceiveSessionSyncSystemDataEngine struct {
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// EventsReceiveSessionSyncSystemDataBrowser maps session:sync:system/system/browser
type EventsReceiveSessionSyncSystemDataBrowser struct {
  Major    *string  `json:"major,omitempty"`
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// EventsReceiveSessionSyncExtendedInformation maps session:sync:extended_informations
type EventsReceiveSessionSyncExtendedInformation struct {
  WebsiteID            *string       `json:"website_id"`
  SessionID            *string       `json:"session_id"`
  ExtendedInformation  *interface{}  `json:"extended_informations"`
}

// EventsReceiveSessionSetState maps session:set_state
type EventsReceiveSessionSetState struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  State      *string  `json:"state"`
}

// EventsReceiveSessionSetBlock maps session:set_block
type EventsReceiveSessionSetBlock struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  IsBlocked  *bool    `json:"is_blocked"`
}

// EventsReceiveSessionSetTags maps session:set_tags
type EventsReceiveSessionSetTags struct {
  WebsiteID  *string    `json:"website_id"`
  SessionID  *string    `json:"session_id"`
  Tags       *[]string  `json:"tags"`
}

// EventsReceiveSessionSetOpened maps session:set_opened
type EventsReceiveSessionSetOpened struct {
  WebsiteID  *string                                    `json:"website_id"`
  SessionID  *string                                    `json:"session_id"`
  Operator   *EventsReceiveSessionSetOpenStateOperator  `json:"operator"`
}

// EventsReceiveSessionSetClosed maps session:set_closed
type EventsReceiveSessionSetClosed struct {
  WebsiteID  *string                                    `json:"website_id"`
  SessionID  *string                                    `json:"session_id"`
  Operator   *EventsReceiveSessionSetOpenStateOperator  `json:"operator"`
}

// EventsReceiveSessionSetOpenStateOperator maps session:set_{opened,closed}/operator
type EventsReceiveSessionSetOpenStateOperator struct {
  UserID  *string  `json:"user_id"`
}

// EventsReceiveSessionRemoved maps session:removed
type EventsReceiveSessionRemoved struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
}

// EventsReceiveMessageSend maps message:send
type EventsReceiveMessageSend struct {
  WebsiteID    *string                          `json:"website_id"`
  SessionID    *string                          `json:"session_id"`
  From         *string                          `json:"from"`
  Type         *string                          `json:"type"`
  Origin       *string                          `json:"origin"`
  Content      *string                          `json:"content"`
  Stamped      *bool                            `json:"stamped"`
  Timestamp    *uint                            `json:"timestamp"`
  Fingerprint  *int                             `json:"fingerprint"`
  User         *EventsReceiveMessageCommonUser  `json:"user"`
}

// EventsReceiveMessageReceived maps message:received
type EventsReceiveMessageReceived struct {
  WebsiteID    *string                          `json:"website_id"`
  SessionID    *string                          `json:"session_id"`
  From         *string                          `json:"from"`
  Type         *string                          `json:"type"`
  Origin       *string                          `json:"origin"`
  Content      *string                          `json:"content"`
  Stamped      *bool                            `json:"stamped"`
  Timestamp    *uint                            `json:"timestamp"`
  Fingerprint  *int                             `json:"fingerprint"`
  User         *EventsReceiveMessageCommonUser  `json:"user"`
}

// EventsReceiveMessageComposeSend maps message:compose:send
type EventsReceiveMessageComposeSend struct {
  WebsiteID    *string  `json:"website_id"`
  SessionID    *string  `json:"session_id"`
  Type         *string  `json:"type"`
  Excerpt      *string  `json:"excerpt"`
}

// EventsReceiveMessageComposeReceive maps message:compose:receive
type EventsReceiveMessageComposeReceive struct {
  WebsiteID    *string                                     `json:"website_id"`
  SessionID    *string                                     `json:"session_id"`
  Compose      *EventsReceiveMessageComposeReceiveCompose  `json:"compose"`
}

// EventsReceiveMessageComposeReceiveCompose maps message:compose:receive/compose
type EventsReceiveMessageComposeReceiveCompose struct {
  UserID       *string  `json:"user_id"`
  Type         *string  `json:"type"`
  Timestamp    *uint    `json:"timestamp"`
}

// EventsReceiveMessageAcknowledgeReadSend maps message:acknowledge:read:send
type EventsReceiveMessageAcknowledgeReadSend struct {
  WebsiteID     *string  `json:"website_id"`
  SessionID     *string  `json:"session_id"`
  Origin        *string  `json:"origin"`
  Fingerprints  *[]int   `json:"fingerprints"`
}

// EventsReceiveMessageAcknowledgeReadReceived maps message:acknowledge:read:received
type EventsReceiveMessageAcknowledgeReadReceived struct {
  WebsiteID     *string  `json:"website_id"`
  SessionID     *string  `json:"session_id"`
  Origin        *string  `json:"origin"`
  Fingerprints  *[]int   `json:"fingerprints"`
}

// EventsReceiveMessageAcknowledgeDelivered maps message:acknowledge:read:delivered
type EventsReceiveMessageAcknowledgeDelivered struct {
  WebsiteID     *string  `json:"website_id"`
  SessionID     *string  `json:"session_id"`
  Origin        *string  `json:"origin"`
  Fingerprints  *[]int   `json:"fingerprints"`
}

// EventsReceiveWebsiteVisitorsCount maps website:update_visitors_count
type EventsReceiveWebsiteVisitorsCount struct {
  WebsiteID      *string  `json:"website_id"`
  VisitorsCount  *uint    `json:"visitors_count"`
}

// EventsReceiveUpdateVisitorsList maps website:update_visitors_list
type EventsReceiveUpdateVisitorsList struct {
  WebsiteID     *string                                 `json:"website_id"`
  VisitorsList  *EventsReceiveUpdateVisitorsListItself  `json:"visitors_list"`
}

// EventsReceiveUpdateVisitorsListItself maps website:update_visitors_list/visitors_list
type EventsReceiveUpdateVisitorsListItself struct {
  Items  *[]EventsReceiveUpdateVisitorsListItselfItem  `json:"items"`
  Tule   *uint                                         `json:"time"`
}

// EventsReceiveUpdateVisitorsListItselfItem maps website:update_visitors_list/items
type EventsReceiveUpdateVisitorsListItselfItem struct {
  SessionID  *string                                          `json:"session_id"`
  Nickname   *string                                          `json:"nickname"`
  City       *string                                          `json:"city"`
  Country    *string                                          `json:"country"`
  Initiated  *bool                                            `json:"initiated"`
  Pages      *EventsReceiveUpdateVisitorsListItselfItemPages  `json:"pages"`
}

// EventsReceiveUpdateVisitorsListItselfItemPages maps website:update_visitors_list/items/pages
type EventsReceiveUpdateVisitorsListItselfItemPages struct {
  Count    *uint                                             `json:"count"`
  Current  *EventsReceiveUpdateVisitorsListItselfItemPagesCurrent   `json:"current"`
}

// EventsReceiveUpdateVisitorsListItselfItemPagesCurrent maps website:update_visitors_list/items/pages/current
type EventsReceiveUpdateVisitorsListItselfItemPagesCurrent struct {
  PageTitle  *string  `json:"page_title"`
  PageURL    *string  `json:"page_url"`
  Timestamp  *uint    `json:"timestamp"`
}

// EventsReceiveBucketURLUploadGenerated maps bucket:url:upload:generated
type EventsReceiveBucketURLUploadGenerated struct {
  From        *string                                         `json:"from"`
  ID          *int                                            `json:"id"`
  Identifier  *string                                         `json:"identifier"`
  Policy      *EventsReceiveBucketURLUploadGeneratedPolicy    `json:"policy"`
  Type        *string                                         `json:"type"`
  URL         *EventsReceiveBucketURLUploadGeneratedURL       `json:"url"`
}

// EventsReceiveBucketURLUploadGeneratedPolicy maps bucket:url:upload:generated/policy
type EventsReceiveBucketURLUploadGeneratedPolicy struct {
  SizeLimit  *uint  `json:"size_limit"`
}

// EventsReceiveBucketURLUploadGeneratedURL maps bucket:url:upload:generated/url
type EventsReceiveBucketURLUploadGeneratedURL struct {
  Resource  *string  `json:"resource"`
  Signed    *string  `json:"signed"`
}

// EventsReceiveBucketURLAvatarGenerated maps bucket:url:avatar:generated
type EventsReceiveBucketURLAvatarGenerated struct {
  From        *string                                       `json:"from"`
  ID          *int                                          `json:"id"`
  Identifier  *string                                       `json:"identifier"`
  Policy      *EventsReceiveBucketURLAvatarGeneratedPolicy  `json:"policy"`
  Type        *string                                       `json:"type"`
  URL         *EventsReceiveBucketURLAvatarGeneratedURL     `json:"url"`
}

// EventsReceiveBucketURLAvatarGeneratedPolicy maps bucket:url:avatar:generated/policy
type EventsReceiveBucketURLAvatarGeneratedPolicy struct {
  SizeLimit  *uint  `json:"size_limit"`
}

// EventsReceiveBucketURLAvatarGeneratedURL maps bucket:url:avatar:generated/url
type EventsReceiveBucketURLAvatarGeneratedURL struct {
  Resource  *string  `json:"resource"`
  Signed    *string  `json:"signed"`
}

// EventsReceiveBillingLinkRedirect maps billing:link:redirect
type EventsReceiveBillingLinkRedirect struct {
  Service  *string  `json:"service"`
  URL      *string  `json:"url"`
}

// EventsReceiveMessageCommonUser maps message:*/user
type EventsReceiveMessageCommonUser struct {
  UserID    *string  `json:"user_id"`
  Nickname  *string  `json:"nickname"`
}


// String returns the string representation of EventsReceiveSessionUpdateAvailability
func (evt EventsReceiveSessionUpdateAvailability) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionRequestInitiated
func (evt EventsReceiveSessionRequestInitiated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetEmail
func (evt EventsReceiveSessionSetEmail) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetAvatar
func (evt EventsReceiveSessionSetAvatar) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetCover
func (evt EventsReceiveSessionSetCover) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetNickname
func (evt EventsReceiveSessionSetNickname) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncPages
func (evt EventsReceiveSessionSyncPages) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncGeolocation
func (evt EventsReceiveSessionSyncGeolocation) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncSystem
func (evt EventsReceiveSessionSyncSystem) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncExtendedInformation
func (evt EventsReceiveSessionSyncExtendedInformation) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetState
func (evt EventsReceiveSessionSetState) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetBlock
func (evt EventsReceiveSessionSetBlock) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetTags
func (evt EventsReceiveSessionSetTags) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetOpened
func (evt EventsReceiveSessionSetOpened) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetClosed
func (evt EventsReceiveSessionSetClosed) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionRemoved
func (evt EventsReceiveSessionRemoved) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageSend
func (evt EventsReceiveMessageSend) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageReceived
func (evt EventsReceiveMessageReceived) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageComposeSend
func (evt EventsReceiveMessageComposeSend) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageComposeReceive
func (evt EventsReceiveMessageComposeReceive) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageAcknowledgeReadSend
func (evt EventsReceiveMessageAcknowledgeReadSend) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageAcknowledgeReadReceived
func (evt EventsReceiveMessageAcknowledgeReadReceived) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageAcknowledgeDelivered
func (evt EventsReceiveMessageAcknowledgeDelivered) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveWebsiteVisitorsCount
func (evt EventsReceiveWebsiteVisitorsCount) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveUpdateVisitorsList
func (evt EventsReceiveUpdateVisitorsList) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBucketURLUploadGenerated
func (evt EventsReceiveBucketURLUploadGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBucketURLAvatarGenerated
func (evt EventsReceiveBucketURLAvatarGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBillingLinkRedirect
func (evt EventsReceiveBillingLinkRedirect) String() string {
  return Stringify(evt)
}


// On registers an event handler on event name
func (register *EventsRegister) On(eventName string, handler interface{}) error {
  c, err := newCaller(handler)
  if err != nil {
    return err
  }

  register.Handlers[eventName] = c

  return nil
}

// BindEvents listens for recognized incoming events
func (register *EventsRegister) BindEvents(so *gosocketio.Client) {
  so.On("session:update_availability", func(chnl *gosocketio.Channel, evt EventsReceiveSessionUpdateAvailability) {
    if hdl, ok := register.Handlers["session:update_availability"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:request:initiated", func(chnl *gosocketio.Channel, evt EventsReceiveSessionRequestInitiated) {
    if hdl, ok := register.Handlers["session:request:initiated"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_email", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetEmail) {
    if hdl, ok := register.Handlers["session:set_email"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_avatar", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetAvatar) {
    if hdl, ok := register.Handlers["session:set_avatar"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_cover", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetCover) {
    if hdl, ok := register.Handlers["session:set_cover"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_nickname", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetNickname) {
    if hdl, ok := register.Handlers["session:set_nickname"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:pages", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncPages) {
    if hdl, ok := register.Handlers["session:sync:pages"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:geolocation", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncGeolocation) {
    if hdl, ok := register.Handlers["session:sync:geolocation"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:system", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncSystem) {
    if hdl, ok := register.Handlers["session:sync:system"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:extended_informations", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncExtendedInformation) {
    if hdl, ok := register.Handlers["session:sync:extended_informations"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_state", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetState) {
    if hdl, ok := register.Handlers["session:set_state"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_block", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetBlock) {
    if hdl, ok := register.Handlers["session:set_block"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_tags", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetTags) {
    if hdl, ok := register.Handlers["session:set_tags"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_opened", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetOpened) {
    if hdl, ok := register.Handlers["session:set_opened"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:set_closed", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetClosed) {
    if hdl, ok := register.Handlers["session:set_closed"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("session:removed", func(chnl *gosocketio.Channel, evt EventsReceiveSessionRemoved) {
    if hdl, ok := register.Handlers["session:removed"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:send", func(chnl *gosocketio.Channel, evt EventsReceiveMessageSend) {
    if hdl, ok := register.Handlers["message:send"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:received", func(chnl *gosocketio.Channel, evt EventsReceiveMessageReceived) {
    if hdl, ok := register.Handlers["message:received"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:compose:send", func(chnl *gosocketio.Channel, evt EventsReceiveMessageComposeSend) {
    if hdl, ok := register.Handlers["message:compose:send"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:compose:receive", func(chnl *gosocketio.Channel, evt EventsReceiveMessageComposeReceive) {
    if hdl, ok := register.Handlers["message:compose:receive"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:acknowledge:read:send", func(chnl *gosocketio.Channel, evt EventsReceiveMessageAcknowledgeReadSend) {
    if hdl, ok := register.Handlers["message:acknowledge:read:send"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:acknowledge:read:received", func(chnl *gosocketio.Channel, evt EventsReceiveMessageAcknowledgeReadReceived) {
    if hdl, ok := register.Handlers["message:acknowledge:read:received"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("message:acknowledge:delivered", func(chnl *gosocketio.Channel, evt EventsReceiveMessageAcknowledgeDelivered) {
    if hdl, ok := register.Handlers["message:acknowledge:delivered"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("website:update_visitors_count", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteVisitorsCount) {
    if hdl, ok := register.Handlers["website:update_visitors_count"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("website:update_visitors_list", func(chnl *gosocketio.Channel, evt EventsReceiveUpdateVisitorsList) {
    if hdl, ok := register.Handlers["website:update_visitors_list"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("bucket:url:upload:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLUploadGenerated) {
    if hdl, ok := register.Handlers["bucket:url:upload:generated"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("bucket:url:avatar:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLAvatarGenerated) {
    if hdl, ok := register.Handlers["bucket:url:avatar:generated"]; ok {
      hdl.callFunc(&evt)
    }
  })

  so.On("billing:link:redirect", func(chnl *gosocketio.Channel, evt EventsReceiveBillingLinkRedirect) {
    if hdl, ok := register.Handlers["billing:link:redirect"]; ok {
      hdl.callFunc(&evt)
    }
  })
}


func (service *EventsService) getEndpointURL() (string) {
  var secure bool

  u, _ := url.Parse(service.client.config.RealtimeEndpointURL)
  host, portString, _ := net.SplitHostPort(u.Host)

  port, _ := strconv.ParseInt(portString, 10, 64)

  if u.Scheme == "https" {
    secure = true
  } else {
    secure = false
  }

  return gosocketio.GetUrl(host, int(port), secure)
}


func (service *EventsService) reconnect(handleDone func(*EventsRegister), connectorChild *int, connectedSocket *bool, child int) {
  // Attempt to reconnect
  for *connectedSocket == false && child == *connectorChild {
    // Hold on.
    time.Sleep(reconnectWait * time.Second)

    *connectorChild++

    service.connect(handleDone, connectorChild, connectedSocket)
  }
}


func (service *EventsService) connect(handleDone func(*EventsRegister), connectorChild *int, connectedSocket *bool) {
  child := *connectorChild

  endpointURL := service.getEndpointURL()
  transport := transport.GetDefaultWebsocketTransport()

  so, err := gosocketio.Dial(endpointURL, transport)

  if err == nil {
    so.On("authenticated", func(chnl *gosocketio.Channel, authenticated bool) {
      if authenticated == true {
        reg := EventsRegister{Handlers: make(map[string]*caller)}

        reg.BindEvents(so)

        handleDone(&reg)
      }
    })

    so.On(gosocketio.OnDisconnection, func(chnl *gosocketio.Channel) {
      *connectedSocket = false

      service.reconnect(handleDone, connectorChild, connectedSocket, child)
    })

    so.On(gosocketio.OnConnection, func(chnl *gosocketio.Channel) {
      *connectedSocket = true

      // Authenticate to socket
      if service.client.auth.Available == true {
        so.Channel.Emit("authentication", eventsSendAuthentication{service.client.auth.Tier, service.client.auth.Username, service.client.auth.Password})
      }
    })
  } else {
    service.reconnect(handleDone, connectorChild, connectedSocket, child)
  }
}


// Listen starts listening for incoming realtime events.
func (service *EventsService) Listen(handleDone func(*EventsRegister)) {
  connectorChild := 0
  connectedSocket := false

  service.connect(handleDone, &connectorChild, &connectedSocket)
}
