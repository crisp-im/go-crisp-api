// Copyright 2016 Crisp IM, Inc. All rights reserved.
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
  "encoding/json"
  "strconv"
)


const (
  reconnectWait = 4
)

var activeSocket *gosocketio.Client


// EventsService service
type EventsService service

// EventsRegister stores event handlers
type EventsRegister struct {
  Handlers  map[string]*caller
}

type eventsSendAuthentication struct {
  Tier      string    `json:"tier"`
  Username  string    `json:"username"`
  Password  string    `json:"password"`
  Events    []string  `json:"events"`
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

// EventsReceiveSessionSetPhone maps session:set_phone
type EventsReceiveSessionSetPhone struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  Phone      *string  `json:"phone"`
}

// EventsReceiveSessionSetAddress maps session:set_address
type EventsReceiveSessionSetAddress struct {
  WebsiteID  *string  `json:"website_id"`
  SessionID  *string  `json:"session_id"`
  Address    *string  `json:"address"`
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

// EventsReceiveSessionSetData maps session:set_data
type EventsReceiveSessionSetData struct {
  WebsiteID  *string       `json:"website_id"`
  SessionID  *string       `json:"session_id"`
  Data       *interface{}  `json:"data"`
}

// EventsReceiveSessionSyncPages maps session:sync:pages
type EventsReceiveSessionSyncPages struct {
  WebsiteID  *string                              `json:"website_id"`
  SessionID  *string                              `json:"session_id"`
  Pages      *[]EventsReceiveSessionSyncPagesOne  `json:"pages"`
}

// EventsReceiveSessionSyncPagesOne maps session:sync:pages/pages
type EventsReceiveSessionSyncPagesOne struct {
  PageTitle     *string  `json:"page_title"`
  PageURL       *string  `json:"page_url"`
  PageReferrer  *string  `json:"page_referrer"`
  Timestamp     *uint    `json:"timestamp"`
}

// EventsReceiveSessionSyncGeolocation maps session:sync:geolocation
type EventsReceiveSessionSyncGeolocation struct {
  WebsiteID    *string                                   `json:"website_id"`
  SessionID    *string                                   `json:"session_id"`
  Geolocation  *EventsReceiveSessionSyncGeolocationData  `json:"geolocation"`
}

// EventsReceiveSessionSyncGeolocationData maps session:sync:geolocation/geolocation
type EventsReceiveSessionSyncGeolocationData struct {
  Coordinates  *EventsReceiveSessionSyncGeolocationDataCoordinates  `json:"coordinates,omitempty"`
  City         *string                                              `json:"city,omitempty"`
  Region       *string                                              `json:"region,omitempty"`
  Country      *string                                              `json:"country,omitempty"`
}

// EventsReceiveSessionSyncGeolocationDataCoordinates mapping
type EventsReceiveSessionSyncGeolocationDataCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}

// EventsReceiveSessionSyncSystem maps session:sync:system
type EventsReceiveSessionSyncSystem struct {
  WebsiteID    *string                              `json:"website_id"`
  SessionID    *string                              `json:"session_id"`
  System       *EventsReceiveSessionSyncSystemData  `json:"system"`
}

// EventsReceiveSessionSyncSystemData maps session:sync:system/system
type EventsReceiveSessionSyncSystemData struct {
  OS         *EventsReceiveSessionSyncSystemDataOS       `json:"os,omitempty"`
  Engine     *EventsReceiveSessionSyncSystemDataEngine   `json:"engine,omitempty"`
  Browser    *EventsReceiveSessionSyncSystemDataBrowser  `json:"browser,omitempty"`
  Useragent  *string                                     `json:"useragent,omitempty"`
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

// EventsReceiveSessionSyncNetwork maps session:sync:network
type EventsReceiveSessionSyncNetwork struct {
  WebsiteID    *string                               `json:"website_id"`
  SessionID    *string                               `json:"session_id"`
  Network      *EventsReceiveSessionSyncNetworkData  `json:"network"`
}

// EventsReceiveSessionSyncNetworkData maps session:sync:network/network
type EventsReceiveSessionSyncNetworkData struct {
  IP  *string  `json:"ip,omitempty"`
}

// EventsReceiveSessionSyncTimezone maps session:sync:timezone
type EventsReceiveSessionSyncTimezone struct {
  WebsiteID    *string                                `json:"website_id"`
  SessionID    *string                                `json:"session_id"`
  Timezone     *EventsReceiveSessionSyncTimezoneData  `json:"timezone"`
}

// EventsReceiveSessionSyncTimezoneData maps session:sync:timezone/timezone
type EventsReceiveSessionSyncTimezoneData struct {
  Offset  *int16  `json:"offset,omitempty"`
}

// EventsReceiveSessionSyncLocales maps session:sync:locales
type EventsReceiveSessionSyncLocales struct {
  WebsiteID    *string                               `json:"website_id"`
  SessionID    *string                               `json:"session_id"`
  Locales      *EventsReceiveSessionSyncLocalesData  `json:"locales"`
}

// EventsReceiveSessionSyncLocalesData maps session:sync:locales/locales
type EventsReceiveSessionSyncLocalesData struct {
  Locales  *[]string  `json:"locales,omitempty"`
}

// EventsReceiveSessionSyncIdentity maps session:sync:identity
type EventsReceiveSessionSyncIdentity struct {
  WebsiteID  *string       `json:"website_id"`
  SessionID  *string       `json:"session_id"`
  Identity   *interface{}  `json:"identity"`
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

// EventsReceiveSessionSetSegments maps session:set_segments
type EventsReceiveSessionSetSegments struct {
  WebsiteID  *string    `json:"website_id"`
  SessionID  *string    `json:"session_id"`
  Segments   *[]string  `json:"segments"`
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

// EventsReceiveGenericMessageType maps message:{send,received} (generic type, type key only)
type EventsReceiveGenericMessageType struct {
  Type  *string  `json:"type"`
}

// EventsReceiveTextMessage maps message:{send,received} (text type)
type EventsReceiveTextMessage struct {
  WebsiteID    *string                          `json:"website_id"`
  SessionID    *string                          `json:"session_id"`
  From         *string                          `json:"from"`
  Type         *string                          `json:"type"`
  Origin       *string                          `json:"origin"`
  Content      *string                          `json:"content"`
  Stamped      *bool                            `json:"stamped"`
  Timestamp    *uint                            `json:"timestamp"`
  Fingerprint  *int                             `json:"fingerprint"`
  User         *EventsReceiveCommonMessageUser  `json:"user"`
}

// EventsReceiveFileMessage maps message:{send,received} (file type)
type EventsReceiveFileMessage struct {
  WebsiteID    *string                           `json:"website_id"`
  SessionID    *string                           `json:"session_id"`
  From         *string                           `json:"from"`
  Type         *string                           `json:"type"`
  Origin       *string                           `json:"origin"`
  Content      *EventsReceiveFileMessageContent  `json:"content"`
  Stamped      *bool                             `json:"stamped"`
  Timestamp    *uint                             `json:"timestamp"`
  Fingerprint  *int                              `json:"fingerprint"`
  User         *EventsReceiveCommonMessageUser   `json:"user"`
}

// EventsReceiveFileMessageContent maps message:{send,received}/content (file type)
type EventsReceiveFileMessageContent struct {
  Name  string  `json:"name"`
  URL   string  `json:"url"`
  Type  string  `json:"type"`
}

// EventsReceiveAnimationMessage maps message:{send,received} (animation type)
type EventsReceiveAnimationMessage struct {
  WebsiteID    *string                                `json:"website_id"`
  SessionID    *string                                `json:"session_id"`
  From         *string                                `json:"from"`
  Type         *string                                `json:"type"`
  Origin       *string                                `json:"origin"`
  Content      *EventsReceiveAnimationMessageContent  `json:"content"`
  Stamped      *bool                                  `json:"stamped"`
  Timestamp    *uint                                  `json:"timestamp"`
  Fingerprint  *int                                   `json:"fingerprint"`
  User         *EventsReceiveCommonMessageUser        `json:"user"`
}

// EventsReceiveAnimationMessageContent maps message:{send,received}/content (animation type)
type EventsReceiveAnimationMessageContent struct {
  URL   string  `json:"url"`
  Type  string  `json:"type"`
}

// EventsReceiveNoteMessage maps message:{send,received} (note type)
type EventsReceiveNoteMessage EventsReceiveTextMessage

// EventsReceiveCommonMessageUser maps message:{send,received}/user
type EventsReceiveCommonMessageUser struct {
  UserID    *string  `json:"user_id"`
  Nickname  *string  `json:"nickname"`
  Avatar    *string  `json:"avatar"`
}

// EventsReceiveMessageComposeSend maps message:compose:send
type EventsReceiveMessageComposeSend struct {
  WebsiteID    *string  `json:"website_id"`
  SessionID    *string  `json:"session_id"`
  Type         *string  `json:"type"`
  Excerpt      *string  `json:"excerpt"`
  Timestamp    *uint    `json:"timestamp"`
}

// EventsReceiveMessageComposeReceive maps message:compose:receive
type EventsReceiveMessageComposeReceive struct {
  WebsiteID    *string  `json:"website_id"`
  SessionID    *string  `json:"session_id"`
  UserID       *string  `json:"user_id"`
  Type         *string  `json:"type"`
  Excerpt      *string  `json:"excerpt"`
  Timestamp    *uint    `json:"timestamp"`
}

// EventsReceiveMessageAcknowledge maps message:acknowledge:*
type EventsReceiveMessageAcknowledge struct {
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
  Time   *uint                                         `json:"time"`
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
  PageTitle     *string  `json:"page_title"`
  PageURL       *string  `json:"page_url"`
  PageReferrer  *string  `json:"page_referrer"`
  Timestamp     *uint    `json:"timestamp"`
}

// EventsReceiveUpdateOperatorsAvailability maps website:update_operators_availability
type EventsReceiveUpdateOperatorsAvailability struct {
  WebsiteID     *string                                          `json:"website_id"`
  UserID        *string                                          `json:"user_id"`
  Availability  *EventsReceiveUpdateOperatorsAvailabilityItself  `json:"availability"`
}

// EventsReceiveUpdateOperatorsAvailabilityItself maps website:update_operators_availability/availability
type EventsReceiveUpdateOperatorsAvailabilityItself struct {
  Type  *string  `json:"type"`
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

// EventsReceiveMediaAnimationListed maps media:animation:listed
type EventsReceiveMediaAnimationListed struct {
  ID       *int                                        `json:"id"`
  Results  *[]EventsReceiveMediaAnimationListedResult  `json:"results"`
}

// EventsReceiveMediaAnimationListedResult maps media:animation:listed/results
type EventsReceiveMediaAnimationListedResult struct {
  Type  *string  `json:"type"`
  URL   *string  `json:"url"`
}

// EventsReceiveBillingLinkRedirect maps billing:link:redirect
type EventsReceiveBillingLinkRedirect struct {
  Service  *string  `json:"service"`
  URL      *string  `json:"url"`
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


// String returns the string representation of EventsReceiveSessionSetData
func (evt EventsReceiveSessionSetData) String() string {
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


// String returns the string representation of EventsReceiveSessionSyncNetwork
func (evt EventsReceiveSessionSyncNetwork) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncTimezone
func (evt EventsReceiveSessionSyncTimezone) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncLocales
func (evt EventsReceiveSessionSyncLocales) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncIdentity
func (evt EventsReceiveSessionSyncIdentity) String() string {
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


// String returns the string representation of EventsReceiveSessionSetSegments
func (evt EventsReceiveSessionSetSegments) String() string {
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


// String returns the string representation of EventsReceiveGenericMessageType
func (evt EventsReceiveGenericMessageType) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveTextMessage
func (evt EventsReceiveTextMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveFileMessage
func (evt EventsReceiveFileMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveAnimationMessage
func (evt EventsReceiveAnimationMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveNoteMessage
func (evt EventsReceiveNoteMessage) String() string {
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


// String returns the string representation of EventsReceiveMessageAcknowledge
func (evt EventsReceiveMessageAcknowledge) String() string {
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


// String returns the string representation of EventsReceiveMediaAnimationListed
func (evt EventsReceiveMediaAnimationListed) String() string {
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
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:request:initiated", func(chnl *gosocketio.Channel, evt EventsReceiveSessionRequestInitiated) {
    if hdl, ok := register.Handlers["session:request:initiated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_email", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetEmail) {
    if hdl, ok := register.Handlers["session:set_email"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_phone", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetPhone) {
    if hdl, ok := register.Handlers["session:set_phone"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_address", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetAddress) {
    if hdl, ok := register.Handlers["session:set_address"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_avatar", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetAvatar) {
    if hdl, ok := register.Handlers["session:set_avatar"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_cover", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetCover) {
    if hdl, ok := register.Handlers["session:set_cover"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_nickname", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetNickname) {
    if hdl, ok := register.Handlers["session:set_nickname"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_data", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetData) {
    if hdl, ok := register.Handlers["session:set_data"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:pages", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncPages) {
    if hdl, ok := register.Handlers["session:sync:pages"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:geolocation", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncGeolocation) {
    if hdl, ok := register.Handlers["session:sync:geolocation"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:system", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncSystem) {
    if hdl, ok := register.Handlers["session:sync:system"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:network", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncNetwork) {
    if hdl, ok := register.Handlers["session:sync:network"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:timezone", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncTimezone) {
    if hdl, ok := register.Handlers["session:sync:timezone"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:locales", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncLocales) {
    if hdl, ok := register.Handlers["session:sync:locales"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:identity", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncIdentity) {
    if hdl, ok := register.Handlers["session:sync:identity"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_state", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetState) {
    if hdl, ok := register.Handlers["session:set_state"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_block", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetBlock) {
    if hdl, ok := register.Handlers["session:set_block"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_segments", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetSegments) {
    if hdl, ok := register.Handlers["session:set_segments"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_opened", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetOpened) {
    if hdl, ok := register.Handlers["session:set_opened"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_closed", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetClosed) {
    if hdl, ok := register.Handlers["session:set_closed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:removed", func(chnl *gosocketio.Channel, evt EventsReceiveSessionRemoved) {
    if hdl, ok := register.Handlers["session:removed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:send", func(chnl *gosocketio.Channel, evt *json.RawMessage) {
    var messageGenericType EventsReceiveGenericMessageType
    json.Unmarshal(*evt, &messageGenericType)

    switch *messageGenericType.Type {
    case "text":
      if hdl, ok := register.Handlers["message:send/text"]; ok {
        var messageSendText EventsReceiveTextMessage
        json.Unmarshal(*evt, &messageSendText)

        go hdl.callFunc(&messageSendText)
      }

    case "file":
      if hdl, ok := register.Handlers["message:send/file"]; ok {
        var messageSendFile EventsReceiveFileMessage
        json.Unmarshal(*evt, &messageSendFile)

        go hdl.callFunc(&messageSendFile)
      }

    case "animation":
      if hdl, ok := register.Handlers["message:send/animation"]; ok {
        var messageSendAnimation EventsReceiveAnimationMessage
        json.Unmarshal(*evt, &messageSendAnimation)

        go hdl.callFunc(&messageSendAnimation)
      }

    case "note":
      if hdl, ok := register.Handlers["message:send/note"]; ok {
        var messageSendNote EventsReceiveNoteMessage
        json.Unmarshal(*evt, &messageSendNote)

        go hdl.callFunc(&messageSendNote)
      }
    }
  })

  so.On("message:received", func(chnl *gosocketio.Channel, evt *json.RawMessage) {
    var messageGenericType EventsReceiveGenericMessageType
    json.Unmarshal(*evt, &messageGenericType)

    switch *messageGenericType.Type {
    case "text":
      if hdl, ok := register.Handlers["message:received/text"]; ok {
        var messageReceivedText EventsReceiveTextMessage
        json.Unmarshal(*evt, &messageReceivedText)

        go hdl.callFunc(&messageReceivedText)
      }

    case "file":
      if hdl, ok := register.Handlers["message:received/file"]; ok {
        var messageReceivedFile EventsReceiveFileMessage
        json.Unmarshal(*evt, &messageReceivedFile)

        go hdl.callFunc(&messageReceivedFile)
      }

    case "animation":
      if hdl, ok := register.Handlers["message:received/animation"]; ok {
        var messageReceivedAnimation EventsReceiveAnimationMessage
        json.Unmarshal(*evt, &messageReceivedAnimation)

        go hdl.callFunc(&messageReceivedAnimation)
      }

    case "note":
      if hdl, ok := register.Handlers["message:received/note"]; ok {
        var messageReceivedNote EventsReceiveNoteMessage
        json.Unmarshal(*evt, &messageReceivedNote)

        go hdl.callFunc(&messageReceivedNote)
      }
    }
  })

  so.On("message:compose:send", func(chnl *gosocketio.Channel, evt EventsReceiveMessageComposeSend) {
    if hdl, ok := register.Handlers["message:compose:send"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:compose:receive", func(chnl *gosocketio.Channel, evt EventsReceiveMessageComposeReceive) {
    if hdl, ok := register.Handlers["message:compose:receive"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:acknowledge:read:send", func(chnl *gosocketio.Channel, evt EventsReceiveMessageAcknowledge) {
    if hdl, ok := register.Handlers["message:acknowledge:read:send"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:acknowledge:read:received", func(chnl *gosocketio.Channel, evt EventsReceiveMessageAcknowledge) {
    if hdl, ok := register.Handlers["message:acknowledge:read:received"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:acknowledge:delivered", func(chnl *gosocketio.Channel, evt EventsReceiveMessageAcknowledge) {
    if hdl, ok := register.Handlers["message:acknowledge:delivered"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:update_visitors_count", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteVisitorsCount) {
    if hdl, ok := register.Handlers["website:update_visitors_count"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:update_visitors_list", func(chnl *gosocketio.Channel, evt EventsReceiveUpdateVisitorsList) {
    if hdl, ok := register.Handlers["website:update_visitors_list"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:update_operators_availability", func(chnl *gosocketio.Channel, evt EventsReceiveUpdateOperatorsAvailability) {
    if hdl, ok := register.Handlers["website:update_operators_availability"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("bucket:url:upload:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLUploadGenerated) {
    if hdl, ok := register.Handlers["bucket:url:upload:generated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("bucket:url:avatar:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLAvatarGenerated) {
    if hdl, ok := register.Handlers["bucket:url:avatar:generated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("media:animation:listed", func(chnl *gosocketio.Channel, evt EventsReceiveMediaAnimationListed) {
    if hdl, ok := register.Handlers["media:animation:listed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("billing:link:redirect", func(chnl *gosocketio.Channel, evt EventsReceiveBillingLinkRedirect) {
    if hdl, ok := register.Handlers["billing:link:redirect"]; ok {
      go hdl.callFunc(&evt)
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


func (service *EventsService) reconnect(events []string, handleDone func(*EventsRegister), connectorChild *int, connectedSocket *bool, child int) {
  // Attempt to reconnect
  for *connectedSocket == false && child == *connectorChild {
    // Hold on.
    time.Sleep(reconnectWait * time.Second)

    *connectorChild++

    service.connect(events, handleDone, connectorChild, connectedSocket)
  }
}


func (service *EventsService) connect(events []string, handleDone func(*EventsRegister), connectorChild *int, connectedSocket *bool) {
  child := *connectorChild

  endpointURL := service.getEndpointURL()
  transport := transport.GetDefaultWebsocketTransport()

  so, err := gosocketio.Dial(endpointURL, transport)

  if err == nil {
    so.On("authenticated", func(chnl *gosocketio.Channel, authenticated bool) {
      if authenticated == true {
        reg := EventsRegister{Handlers: make(map[string]*caller)}

        activeSocket = so

        reg.BindEvents(so)

        handleDone(&reg)
      }
    })

    so.On(gosocketio.OnDisconnection, func(chnl *gosocketio.Channel) {
      *connectedSocket = false

      service.reconnect(events, handleDone, connectorChild, connectedSocket, child)
    })

    so.On(gosocketio.OnConnection, func(chnl *gosocketio.Channel) {
      *connectedSocket = true

      // Authenticate to socket
      if service.client.auth.Available == true {
        so.Channel.Emit("authentication", eventsSendAuthentication{Tier: service.client.auth.Tier, Username: service.client.auth.Username, Password: service.client.auth.Password, Events: events})
      }
    })
  } else {
    service.reconnect(events, handleDone, connectorChild, connectedSocket, child)
  }
}


// Bind emits a socket bind event which associates the socket to its channels
func (service *EventsService) Bind() {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind", "")
  }
}


// Listen starts listening for incoming realtime events.
func (service *EventsService) Listen(events []string, handleDone func(*EventsRegister)) {
  connectorChild := 0
  connectedSocket := false

  service.connect(events, handleDone, &connectorChild, &connectedSocket)
}
