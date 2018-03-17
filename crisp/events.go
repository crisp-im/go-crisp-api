// Copyright 2017 Crisp IM SARL All rights reserved.
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

// EventsSendAuthentication sends authentication
type EventsSendAuthentication struct {
  Tier      string    `json:"tier"`
  Username  string    `json:"username"`
  Password  string    `json:"password"`
  Events    []string  `json:"events"`
}

// EventsSendBind sends bind
type EventsSendBind struct {
  Events  []string  `json:"events"`
}

// EventsReceiveGenericMessageType receive a generic message type
type EventsReceiveGenericMessageType struct {
  Type  *string  `json:"type"`
}

// EventsGeneric maps a generic event
type EventsGeneric struct {
  RoutingIDs  *[]string  `json:"routing_ids,omitempty"`
}

// EventsUserGeneric maps a generic user
type EventsUserGeneric struct {
  UserID  *string  `json:"user_id"`
}

// EventsWebsiteGeneric maps a generic website
type EventsWebsiteGeneric struct {
  WebsiteID  *string  `json:"website_id"`
}

// EventsImportGeneric maps a generic import
type EventsImportGeneric struct {
  ImportID  *string  `json:"import_id"`
}

// EventsSessionGeneric maps a generic session (unbound from website)
type EventsSessionGenericUnbound struct {
  SessionID  *string  `json:"session_id"`
}

// EventsSessionGeneric maps a generic session
type EventsSessionGeneric struct {
  EventsWebsiteGeneric
  EventsSessionGenericUnbound
}

// EventsBrowsingGeneric maps a generic browsing
type EventsBrowsingGeneric struct {
  EventsSessionGeneric
  BrowsingID  *string  `json:"browsing_id"`
}

// EventsCallGeneric maps a generic call
type EventsCallGeneric struct {
  EventsSessionGeneric
  CallID  *string  `json:"call_id"`
}

// EventsPeopleGeneric maps a generic people
type EventsPeopleGeneric struct {
  EventsWebsiteGeneric
  PeopleID  *string  `json:"people_id"`
}

// EventsCampaignGeneric maps a generic campaign
type EventsCampaignGeneric struct {
  EventsWebsiteGeneric
  CampaignID  *string  `json:"campaign_id"`
}

// EventsReceiveGenericMessage maps a generic message
type EventsReceiveGenericMessage struct {
  EventsReceiveGenericMessageType
  EventsSessionGeneric
  From         *string                          `json:"from"`
  Origin       *string                          `json:"origin"`
  Mentions     *[]string                        `json:"mentions"`
  Stamped      *bool                            `json:"stamped"`
  Timestamp    *uint                            `json:"timestamp"`
  Fingerprint  *int                             `json:"fingerprint"`
  User         *EventsReceiveCommonMessageUser  `json:"user"`
}

// EventsReceiveCommonMessageUser maps a message user
type EventsReceiveCommonMessageUser struct {
  Type      *string  `json:"type"`
  UserID    *string  `json:"user_id"`
  Nickname  *string  `json:"nickname"`
  Avatar    *string  `json:"avatar"`
}

// EventsReceiveAuthenticationUnauthorized maps unauthorized
type EventsReceiveAuthenticationUnauthorized struct {
  Message  *string  `json:"message"`
}

// EventsReceiveSessionUpdateAvailability maps session:update_availability
type EventsReceiveSessionUpdateAvailability struct {
  EventsGeneric
  EventsSessionGeneric
  Availability  *string  `json:"availability"`
}

// EventsReceiveSessionUpdateVerify maps session:update_verify
type EventsReceiveSessionUpdateVerify struct {
  EventsGeneric
  EventsSessionGeneric
  IsVerified  *bool  `json:"is_verified"`
}

// EventsReceiveSessionRequestInitiated maps session:request:initiated
type EventsReceiveSessionRequestInitiated struct {
  EventsGeneric
  EventsSessionGeneric
}

// EventsReceiveSessionSetEmail maps session:set_email
type EventsReceiveSessionSetEmail struct {
  EventsGeneric
  EventsSessionGeneric
  Email  *string  `json:"email"`
}

// EventsReceiveSessionSetPhone maps session:set_phone
type EventsReceiveSessionSetPhone struct {
  EventsGeneric
  EventsSessionGeneric
  Phone  *string  `json:"phone"`
}

// EventsReceiveSessionSetAddress maps session:set_address
type EventsReceiveSessionSetAddress struct {
  EventsGeneric
  EventsSessionGeneric
  Address  *string  `json:"address"`
}

// EventsReceiveSessionSetAvatar maps session:set_avatar
type EventsReceiveSessionSetAvatar struct {
  EventsGeneric
  EventsSessionGeneric
  Avatar  *string  `json:"avatar"`
}

// EventsReceiveSessionSetNickname maps session:set_nickname
type EventsReceiveSessionSetNickname struct {
  EventsGeneric
  EventsSessionGeneric
  Nickname  *string  `json:"nickname"`
}

// EventsReceiveSessionSetData maps session:set_data
type EventsReceiveSessionSetData struct {
  EventsGeneric
  EventsSessionGeneric
  Data  *interface{}  `json:"data"`
}

// EventsReceiveSessionSyncPages maps session:sync:pages
type EventsReceiveSessionSyncPages struct {
  EventsGeneric
  EventsSessionGeneric
  Pages  *[]EventsReceiveSessionSyncPagesOne  `json:"pages"`
}

// EventsReceiveSessionSyncPagesOne maps session:sync:pages/pages
type EventsReceiveSessionSyncPagesOne struct {
  PageTitle     *string  `json:"page_title"`
  PageURL       *string  `json:"page_url"`
  PageReferrer  *string  `json:"page_referrer"`
  Timestamp     *uint    `json:"timestamp"`
}

// EventsReceiveSessionSyncEvents maps session:sync:events
type EventsReceiveSessionSyncEvents struct {
  EventsGeneric
  EventsSessionGeneric
  Events  *EventsReceiveSessionSyncEventsOne  `json:"events"`
}

// EventsReceiveSessionSyncEventsOne maps session:sync:events/events
type EventsReceiveSessionSyncEventsOne struct {
  Text       *string       `json:"text"`
  Data       *interface{}  `json:"data"`
  Color      *string       `json:"color"`
  Timestamp  *uint         `json:"timestamp"`
}

// EventsReceiveSessionSyncCapabilities maps session:sync:capabilities
type EventsReceiveSessionSyncCapabilities struct {
  EventsGeneric
  EventsSessionGeneric
  Capabilities  *EventsReceiveSessionSyncCapabilitiesData  `json:"capabilities"`
}

// EventsReceiveSessionSyncCapabilitiesData maps session:sync:capabilities/capabilities
type EventsReceiveSessionSyncCapabilitiesData struct {
  Capabilities  *[]string  `json:"capabilities,omitempty"`
}

// EventsReceiveSessionSyncGeolocation maps session:sync:geolocation
type EventsReceiveSessionSyncGeolocation struct {
  EventsGeneric
  EventsSessionGeneric
  Geolocation  *EventsReceiveSessionSyncGeolocationData  `json:"geolocation"`
}

// EventsReceiveSessionSyncGeolocationData maps session:sync:geolocation/geolocation
type EventsReceiveSessionSyncGeolocationData struct {
  Coordinates  *EventsReceiveSessionSyncGeolocationDataCoordinates  `json:"coordinates,omitempty"`
  City         *string                                              `json:"city,omitempty"`
  Region       *string                                              `json:"region,omitempty"`
  Country      *string                                              `json:"country,omitempty"`
}

// EventsReceiveSessionSyncGeolocationDataCoordinates maps session:sync:geolocation/geolocation/coordinates
type EventsReceiveSessionSyncGeolocationDataCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}

// EventsReceiveSessionSyncSystem maps session:sync:system
type EventsReceiveSessionSyncSystem struct {
  EventsGeneric
  EventsSessionGeneric
  System  *EventsReceiveSessionSyncSystemData  `json:"system"`
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
  EventsGeneric
  EventsSessionGeneric
  Network  *EventsReceiveSessionSyncNetworkData  `json:"network"`
}

// EventsReceiveSessionSyncNetworkData maps session:sync:network/network
type EventsReceiveSessionSyncNetworkData struct {
  IP  *string  `json:"ip,omitempty"`
}

// EventsReceiveSessionSyncTimezone maps session:sync:timezone
type EventsReceiveSessionSyncTimezone struct {
  EventsGeneric
  EventsSessionGeneric
  Timezone  *EventsReceiveSessionSyncTimezoneData  `json:"timezone"`
}

// EventsReceiveSessionSyncTimezoneData maps session:sync:timezone/timezone
type EventsReceiveSessionSyncTimezoneData struct {
  Offset  *int16  `json:"offset,omitempty"`
}

// EventsReceiveSessionSyncLocales maps session:sync:locales
type EventsReceiveSessionSyncLocales struct {
  EventsGeneric
  EventsSessionGeneric
  Locales  *EventsReceiveSessionSyncLocalesData  `json:"locales"`
}

// EventsReceiveSessionSyncLocalesData maps session:sync:locales/locales
type EventsReceiveSessionSyncLocalesData struct {
  Locales  *[]string  `json:"locales,omitempty"`
}

// EventsReceiveSessionSetState maps session:set_state
type EventsReceiveSessionSetState struct {
  EventsGeneric
  EventsSessionGeneric
  State  *string  `json:"state"`
}

// EventsReceiveSessionSetBlock maps session:set_block
type EventsReceiveSessionSetBlock struct {
  EventsGeneric
  EventsSessionGeneric
  IsBlocked  *bool  `json:"is_blocked"`
}

// EventsReceiveSessionSetSegments maps session:set_segments
type EventsReceiveSessionSetSegments struct {
  EventsGeneric
  EventsSessionGeneric
  Segments  *[]string  `json:"segments"`
}

// EventsReceiveSessionSetOpened maps session:set_opened
type EventsReceiveSessionSetOpened struct {
  EventsGeneric
  EventsSessionGeneric
  Operator  *EventsReceiveSessionSetOpenStateOperator  `json:"operator"`
}

// EventsReceiveSessionSetClosed maps session:set_closed
type EventsReceiveSessionSetClosed struct {
  EventsGeneric
  EventsSessionGeneric
  Operator  *EventsReceiveSessionSetOpenStateOperator  `json:"operator"`
}

// EventsReceiveSessionSetOpenStateOperator maps session:set_{opened,closed}/operator
type EventsReceiveSessionSetOpenStateOperator struct {
  EventsUserGeneric
}

// EventsReceiveSessionSetMentions maps session:set_mentions
type EventsReceiveSessionSetMentions struct {
  EventsGeneric
  EventsSessionGeneric
  Mentions  *[]string  `json:"mentions"`
}

// EventsReceiveSessionSetRouting maps session:set_routing
type EventsReceiveSessionSetRouting struct {
  EventsGeneric
  EventsSessionGeneric
  RoutingID  *string  `json:"routing_id"`
}

// EventsReceiveSessionRemoved maps session:removed
type EventsReceiveSessionRemoved struct {
  EventsGeneric
  EventsSessionGeneric
}

// EventsReceiveMessageUpdated maps message:updated
type EventsReceiveMessageUpdated struct {
  EventsGeneric
  EventsSessionGeneric
  Fingerprint  *int          `json:"fingerprint"`
  Content      *interface{}  `json:"content"`
}

// EventsReceiveTextMessage maps message:{send,received} (text type)
type EventsReceiveTextMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *string  `json:"content"`
}

// EventsReceiveFileMessage maps message:{send,received} (file type)
type EventsReceiveFileMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceiveFileMessageContent  `json:"content"`
}

// EventsReceiveFileMessageContent maps message:{send,received}/content (file type)
type EventsReceiveFileMessageContent struct {
  Name  string  `json:"name"`
  URL   string  `json:"url"`
  Type  string  `json:"type"`
}

// EventsReceiveAnimationMessage maps message:{send,received} (animation type)
type EventsReceiveAnimationMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceiveAnimationMessageContent  `json:"content"`
}

// EventsReceiveAnimationMessageContent maps message:{send,received}/content (animation type)
type EventsReceiveAnimationMessageContent struct {
  URL   string  `json:"url"`
  Type  string  `json:"type"`
}

// EventsReceiveAudioMessage maps message:{send,received} (audio type)
type EventsReceiveAudioMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceiveAudioMessageContent  `json:"content"`
}

// EventsReceiveAudioMessageContent maps message:{send,received}/content (audio type)
type EventsReceiveAudioMessageContent struct {
  URL       string  `json:"url"`
  Type      string  `json:"type"`
  Duration  uint16  `json:"duration"`
}

// EventsReceivePickerMessage maps message:{send,received} (picker type)
type EventsReceivePickerMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceivePickerMessageContent  `json:"content"`
}

// EventsReceivePickerMessageContent maps message:{send,received}/content (picker type)
type EventsReceivePickerMessageContent struct {
  ID       string                                      `json:"id"`
  Text     string                                      `json:"text"`
  Choices  *[]EventsReceivePickerMessageContentChoice  `json:"choices"`
}

// EventsReceivePickerMessageContentChoice maps message:{send,received}/content/choices (picker type)
type EventsReceivePickerMessageContentChoice struct {
  Value     string  `json:"value"`
  Label     string  `json:"label"`
  Selected  bool    `json:"selected"`
}

// EventsReceiveFieldMessage maps message:{send,received} (field type)
type EventsReceiveFieldMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceiveFieldMessageContent  `json:"content"`
}

// EventsReceiveFieldMessageContent maps message:{send,received}/content (field type)
type EventsReceiveFieldMessageContent struct {
  ID       string  `json:"id"`
  Text     string  `json:"text"`
  Explain  string  `json:"explain"`
  Value    string  `json:"value"`
}

// EventsReceiveNoteMessage maps message:{send,received} (note type)
type EventsReceiveNoteMessage EventsReceiveTextMessage

// EventsReceiveMessageComposeSend maps message:compose:send
type EventsReceiveMessageComposeSend struct {
  EventsGeneric
  EventsSessionGeneric
  Type       *string  `json:"type"`
  Excerpt    *string  `json:"excerpt"`
  Timestamp  *uint    `json:"timestamp"`
}

// EventsReceiveMessageComposeReceive maps message:compose:receive
type EventsReceiveMessageComposeReceive struct {
  EventsGeneric
  EventsSessionGeneric
  Type       *string                                  `json:"type"`
  Excerpt    *string                                  `json:"excerpt"`
  Timestamp  *uint                                    `json:"timestamp"`
  User       *EventsReceiveMessageComposeReceiveUser  `json:"user"`
}

// EventsReceiveMessageComposeReceiveUser maps message:compose:receive/user
type EventsReceiveMessageComposeReceiveUser struct {
  UserID    *string  `json:"user_id"`
  Nickname  *string  `json:"nickname"`
  Avatar    *string  `json:"avatar"`
}

// EventsReceiveMessageAcknowledge maps message:acknowledge:*
type EventsReceiveMessageAcknowledge struct {
  EventsGeneric
  EventsSessionGeneric
  Origin        *string  `json:"origin"`
  Fingerprints  *[]int   `json:"fingerprints"`
}

// EventsReceiveMessageNotify maps message:notify:*
type EventsReceiveMessageNotify struct {
  EventsGeneric
  EventsSessionGeneric
}

// EventsReceivePeopleProfileCreated maps people:profile:created
type EventsReceivePeopleProfileCreated struct {
  EventsGeneric
  EventsPeopleGeneric
  Email  *string  `json:"email"`
}

// EventsReceivePeopleProfileRemoved maps people:profile:removed
type EventsReceivePeopleProfileRemoved EventsReceivePeopleProfileCreated

// EventsPeopleBindSession maps people:bind:session
type EventsPeopleBindSession struct {
  EventsGeneric
  EventsPeopleGeneric
  EventsSessionGenericUnbound
}

// EventsPeopleSyncProfile maps people:sync:profile
type EventsPeopleSyncProfile struct {
  EventsGeneric
  EventsPeopleGeneric
  EventsSessionGenericUnbound
  Identity  *PeopleProfileCard  `json:"identity"`
}

// EventsPeopleImportProgress maps people:import:progress
type EventsPeopleImportProgress struct {
  EventsGeneric
  EventsImportGeneric
  Progress  *uint8                            `json:"progress"`
  Count     *EventsPeopleImportProgressCount  `json:"count"`
}

// EventsPeopleImportProgressCount maps people:import:progress/count
type EventsPeopleImportProgressCount struct {
  Total      *uint32  `json:"total"`
  Remaining  *uint32  `json:"remaining"`
}

// EventsPeopleImportDone maps people:import:done
type EventsPeopleImportDone struct {
  EventsGeneric
  EventsImportGeneric
  Error  *bool  `json:"error"`
}

// EventsCampaignProgress maps campaign:progress
type EventsCampaignProgress struct {
  EventsGeneric
  EventsCampaignGeneric
  Progress  *uint8  `json:"progress"`
}

// EventsCampaignDispatched maps campaign:dispatched
type EventsCampaignDispatched struct {
  EventsGeneric
  EventsCampaignGeneric
}

// EventsCampaignRunning maps campaign:running
type EventsCampaignRunning struct {
  EventsGeneric
  EventsCampaignGeneric
  Running  *bool  `json:"running"`
}

// EventsBrowsingRequestInitiated maps browsing:request:initiated
type EventsBrowsingRequestInitiated struct {
  EventsGeneric
  EventsBrowsingGeneric
}

// EventsBrowsingRequestRejected maps browsing:request:rejected
type EventsBrowsingRequestRejected struct {
  EventsGeneric
  EventsSessionGeneric
}

// EventsCallRequestInitiated maps call:request:initiated
type EventsCallRequestInitiated struct {
  EventsGeneric
  EventsCallGeneric
}

// EventsCallRequestRejected maps call:request:rejected
type EventsCallRequestRejected struct {
  EventsGeneric
  EventsCallGeneric
}

// EventsServiceTranslateProcessed maps service:translate:processed
type EventsServiceTranslateProcessed struct {
  EventsGeneric
  EventsWebsiteGeneric
  ID    *string  `json:"id"`
  Text  *string  `json:"text"`
}

// EventsReceiveWebsiteUpdateVisitorsCount maps website:update_visitors_count
type EventsReceiveWebsiteUpdateVisitorsCount struct {
  EventsGeneric
  EventsWebsiteGeneric
  VisitorsCount  *uint    `json:"visitors_count"`
}

// EventsReceiveWebsiteUpdateOperatorsAvailability maps website:update_operators_availability
type EventsReceiveWebsiteUpdateOperatorsAvailability struct {
  EventsGeneric
  EventsWebsiteGeneric
  EventsUserGeneric
  Availability  *EventsReceiveWebsiteUpdateOperatorsAvailabilityItself  `json:"availability"`
}

// EventsReceiveWebsiteUpdateOperatorsAvailabilityItself maps website:update_operators_availability/availability
type EventsReceiveWebsiteUpdateOperatorsAvailabilityItself struct {
  Type  *string  `json:"type"`
}

// EventsReceiveWebsiteUsersAvailable maps website:users:available
type EventsReceiveWebsiteUsersAvailable struct {
  EventsGeneric
  EventsWebsiteGeneric
  Available  *bool  `json:"available"`
}

// EventsReceiveWebsiteValidateDomainValid maps website:validate:domain:valid
type EventsReceiveWebsiteValidateDomainValid struct {
  EventsGeneric
  EventsWebsiteGeneric
  Domain   *string                                           `json:"domain"`
  Records  *[]EventsReceiveWebsiteValidateDomainValidRecord  `json:"records,omitempty"`
}

// EventsReceiveWebsiteValidateDomainValidRecord maps website:validate:domain:valid/records
type EventsReceiveWebsiteValidateDomainValidRecord struct {
  Type   *string  `json:"type,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// EventsReceiveWebsiteValidateDomainInvalid maps website:validate:domain:invalid
type EventsReceiveWebsiteValidateDomainInvalid EventsReceiveWebsiteValidateDomainValid

// EventsReceiveBucketURLUploadGenerated maps bucket:url:upload:generated
type EventsReceiveBucketURLUploadGenerated struct {
  EventsGeneric
  From        *string                                         `json:"from"`
  ID          *string                                         `json:"id"`
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
type EventsReceiveBucketURLAvatarGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLCampaignGenerated maps bucket:url:campaign:generated
type EventsReceiveBucketURLCampaignGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLHelpdeskGenerated maps bucket:url:helpdesk:generated
type EventsReceiveBucketURLHelpdeskGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLProcessingGenerated maps bucket:url:processing:generated
type EventsReceiveBucketURLProcessingGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveMediaAnimationListed maps media:animation:listed
type EventsReceiveMediaAnimationListed struct {
  EventsGeneric
  ID       *int                                        `json:"id"`
  Results  *[]EventsReceiveMediaAnimationListedResult  `json:"results"`
}

// EventsReceiveMediaAnimationListedResult maps media:animation:listed/results
type EventsReceiveMediaAnimationListedResult struct {
  Type  *string  `json:"type"`
  URL   *string  `json:"url"`
}

// EventsReceiveEmailSubscribe maps email:subscribe
type EventsReceiveEmailSubscribe struct {
  EventsGeneric
  EventsWebsiteGeneric
  Email       *string  `json:"email"`
  Subscribed  *bool    `json:"subscribed"`
}

// EventsReceiveEmailTrackView maps email:track:view
type EventsReceiveEmailTrackView struct {
  EventsGeneric
  EventsWebsiteGeneric
  Type        *string  `json:"type"`
  Identifier  *string  `json:"identifier"`
  Mode        *string  `json:"mode"`
}

// EventsReceiveBillingLinkRedirect maps billing:link:redirect
type EventsReceiveBillingLinkRedirect struct {
  EventsGeneric
  Service  *string  `json:"service"`
  URL      *string  `json:"url"`
}


// String returns the string representation of EventsReceiveSessionUpdateAvailability
func (evt EventsReceiveSessionUpdateAvailability) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionUpdateVerify
func (evt EventsReceiveSessionUpdateVerify) String() string {
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


// String returns the string representation of EventsReceiveSessionSyncEvents
func (evt EventsReceiveSessionSyncEvents) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSyncCapabilities
func (evt EventsReceiveSessionSyncCapabilities) String() string {
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


// String returns the string representation of EventsReceiveSessionSetMentions
func (evt EventsReceiveSessionSetMentions) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionSetRouting
func (evt EventsReceiveSessionSetRouting) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveSessionRemoved
func (evt EventsReceiveSessionRemoved) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageUpdated
func (evt EventsReceiveMessageUpdated) String() string {
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


// String returns the string representation of EventsReceiveAudioMessage
func (evt EventsReceiveAudioMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceivePickerMessage
func (evt EventsReceivePickerMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveFieldMessage
func (evt EventsReceiveFieldMessage) String() string {
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


// String returns the string representation of EventsReceiveMessageNotify
func (evt EventsReceiveMessageNotify) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceivePeopleProfileCreated
func (evt EventsReceivePeopleProfileCreated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceivePeopleProfileRemoved
func (evt EventsReceivePeopleProfileRemoved) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsPeopleBindSession
func (evt EventsPeopleBindSession) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsPeopleSyncProfile
func (evt EventsPeopleSyncProfile) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsPeopleImportProgress
func (evt EventsPeopleImportProgress) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsPeopleImportDone
func (evt EventsPeopleImportDone) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsCampaignProgress
func (evt EventsCampaignProgress) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsCampaignDispatched
func (evt EventsCampaignDispatched) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsCampaignRunning
func (evt EventsCampaignRunning) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsBrowsingRequestInitiated
func (evt EventsBrowsingRequestInitiated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsBrowsingRequestRejected
func (evt EventsBrowsingRequestRejected) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsCallRequestInitiated
func (evt EventsCallRequestInitiated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsCallRequestRejected
func (evt EventsCallRequestRejected) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveWebsiteUpdateVisitorsCount
func (evt EventsReceiveWebsiteUpdateVisitorsCount) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveWebsiteUpdateOperatorsAvailability
func (evt EventsReceiveWebsiteUpdateOperatorsAvailability) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveWebsiteUsersAvailable
func (evt EventsReceiveWebsiteUsersAvailable) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveWebsiteValidateDomainValid
func (evt EventsReceiveWebsiteValidateDomainValid) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveWebsiteValidateDomainInvalid
func (evt EventsReceiveWebsiteValidateDomainInvalid) String() string {
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


// String returns the string representation of EventsReceiveBucketURLCampaignGenerated
func (evt EventsReceiveBucketURLCampaignGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBucketURLHelpdeskGenerated
func (evt EventsReceiveBucketURLHelpdeskGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBucketURLProcessingGenerated
func (evt EventsReceiveBucketURLProcessingGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMediaAnimationListed
func (evt EventsReceiveMediaAnimationListed) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveEmailSubscribe
func (evt EventsReceiveEmailSubscribe) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveEmailTrackView
func (evt EventsReceiveEmailTrackView) String() string {
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

  so.On("session:update_verify", func(chnl *gosocketio.Channel, evt EventsReceiveSessionUpdateVerify) {
    if hdl, ok := register.Handlers["session:update_verify"]; ok {
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

  so.On("session:sync:events", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncEvents) {
    if hdl, ok := register.Handlers["session:sync:events"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:sync:capabilities", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSyncCapabilities) {
    if hdl, ok := register.Handlers["session:sync:capabilities"]; ok {
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

  so.On("session:set_mentions", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetMentions) {
    if hdl, ok := register.Handlers["session:set_mentions"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:set_routing", func(chnl *gosocketio.Channel, evt EventsReceiveSessionSetRouting) {
    if hdl, ok := register.Handlers["session:set_routing"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("session:removed", func(chnl *gosocketio.Channel, evt EventsReceiveSessionRemoved) {
    if hdl, ok := register.Handlers["session:removed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:updated", func(chnl *gosocketio.Channel, evt EventsReceiveMessageUpdated) {
    if hdl, ok := register.Handlers["message:updated"]; ok {
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

    case "audio":
      if hdl, ok := register.Handlers["message:send/audio"]; ok {
        var messageSendAudio EventsReceiveAudioMessage
        json.Unmarshal(*evt, &messageSendAudio)

        go hdl.callFunc(&messageSendAudio)
      }

    case "picker":
      if hdl, ok := register.Handlers["message:send/picker"]; ok {
        var messageSendPicker EventsReceivePickerMessage
        json.Unmarshal(*evt, &messageSendPicker)

        go hdl.callFunc(&messageSendPicker)
      }

    case "field":
      if hdl, ok := register.Handlers["message:send/field"]; ok {
        var messageSendField EventsReceiveFieldMessage
        json.Unmarshal(*evt, &messageSendField)

        go hdl.callFunc(&messageSendField)
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

    case "audio":
      if hdl, ok := register.Handlers["message:received/audio"]; ok {
        var messageReceivedAudio EventsReceiveAudioMessage
        json.Unmarshal(*evt, &messageReceivedAudio)

        go hdl.callFunc(&messageReceivedAudio)
      }

    case "picker":
      if hdl, ok := register.Handlers["message:received/picker"]; ok {
        var messageReceivedPicker EventsReceivePickerMessage
        json.Unmarshal(*evt, &messageReceivedPicker)

        go hdl.callFunc(&messageReceivedPicker)
      }

    case "field":
      if hdl, ok := register.Handlers["message:received/field"]; ok {
        var messageReceivedField EventsReceiveFieldMessage
        json.Unmarshal(*evt, &messageReceivedField)

        go hdl.callFunc(&messageReceivedField)
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

  so.On("message:notify:unread:send", func(chnl *gosocketio.Channel, evt EventsReceiveMessageNotify) {
    if hdl, ok := register.Handlers["message:notify:unread:send"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("message:notify:unread:received", func(chnl *gosocketio.Channel, evt EventsReceiveMessageNotify) {
    if hdl, ok := register.Handlers["message:notify:unread:received"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("people:profile:created", func(chnl *gosocketio.Channel, evt EventsReceivePeopleProfileCreated) {
    if hdl, ok := register.Handlers["people:profile:created"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("people:profile:removed", func(chnl *gosocketio.Channel, evt EventsReceivePeopleProfileRemoved) {
    if hdl, ok := register.Handlers["people:profile:removed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("people:bind:session", func(chnl *gosocketio.Channel, evt EventsPeopleBindSession) {
    if hdl, ok := register.Handlers["people:bind:session"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("people:sync:profile", func(chnl *gosocketio.Channel, evt EventsPeopleSyncProfile) {
    if hdl, ok := register.Handlers["people:sync:profile"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("people:import:progress", func(chnl *gosocketio.Channel, evt EventsPeopleImportProgress) {
    if hdl, ok := register.Handlers["people:import:progress"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("people:import:done", func(chnl *gosocketio.Channel, evt EventsPeopleImportDone) {
    if hdl, ok := register.Handlers["people:import:done"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("campaign:progress", func(chnl *gosocketio.Channel, evt EventsCampaignProgress) {
    if hdl, ok := register.Handlers["campaign:progress"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("campaign:dispatched", func(chnl *gosocketio.Channel, evt EventsCampaignDispatched) {
    if hdl, ok := register.Handlers["campaign:dispatched"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("campaign:running", func(chnl *gosocketio.Channel, evt EventsCampaignRunning) {
    if hdl, ok := register.Handlers["campaign:running"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("browsing:request:initiated", func(chnl *gosocketio.Channel, evt EventsBrowsingRequestInitiated) {
    if hdl, ok := register.Handlers["browsing:request:initiated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("browsing:request:rejected", func(chnl *gosocketio.Channel, evt EventsBrowsingRequestRejected) {
    if hdl, ok := register.Handlers["browsing:request:rejected"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("call:request:initiated", func(chnl *gosocketio.Channel, evt EventsCallRequestInitiated) {
    if hdl, ok := register.Handlers["call:request:initiated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("call:request:rejected", func(chnl *gosocketio.Channel, evt EventsCallRequestRejected) {
    if hdl, ok := register.Handlers["call:request:rejected"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("service:translate:processed", func(chnl *gosocketio.Channel, evt EventsServiceTranslateProcessed) {
    if hdl, ok := register.Handlers["service:translate:processed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:update_visitors_count", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteUpdateVisitorsCount) {
    if hdl, ok := register.Handlers["website:update_visitors_count"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:update_operators_availability", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteUpdateOperatorsAvailability) {
    if hdl, ok := register.Handlers["website:update_operators_availability"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:users:available", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteUsersAvailable) {
    if hdl, ok := register.Handlers["website:users:available"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:validate:domain:valid", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteValidateDomainValid) {
    if hdl, ok := register.Handlers["website:validate:domain:valid"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("website:validate:domain:invalid", func(chnl *gosocketio.Channel, evt EventsReceiveWebsiteValidateDomainInvalid) {
    if hdl, ok := register.Handlers["website:validate:domain:invalid"]; ok {
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

  so.On("bucket:url:campaign:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLCampaignGenerated) {
    if hdl, ok := register.Handlers["bucket:url:campaign:generated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("bucket:url:helpdesk:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLHelpdeskGenerated) {
    if hdl, ok := register.Handlers["bucket:url:helpdesk:generated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("bucket:url:processing:generated", func(chnl *gosocketio.Channel, evt EventsReceiveBucketURLProcessingGenerated) {
    if hdl, ok := register.Handlers["bucket:url:processing:generated"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("media:animation:listed", func(chnl *gosocketio.Channel, evt EventsReceiveMediaAnimationListed) {
    if hdl, ok := register.Handlers["media:animation:listed"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("email:subscribe", func(chnl *gosocketio.Channel, evt EventsReceiveEmailSubscribe) {
    if hdl, ok := register.Handlers["email:subscribe"]; ok {
      go hdl.callFunc(&evt)
    }
  })

  so.On("email:track:view", func(chnl *gosocketio.Channel, evt EventsReceiveEmailTrackView) {
    if hdl, ok := register.Handlers["email:track:view"]; ok {
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
        so.Channel.Emit("authentication", EventsSendAuthentication{Tier: service.client.auth.Tier, Username: service.client.auth.Username, Password: service.client.auth.Password, Events: events})
      }
    })
  } else {
    service.reconnect(events, handleDone, connectorChild, connectedSocket, child)
  }
}


// Rebind emits an empty socket bind event which associates the socket to its channels (without modifying allowed events)
func (service *EventsService) Rebind() {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind", "")
  }
}


// Bind emits a socket bind event which associates the socket to its channels (with allowed events)
func (service *EventsService) Bind(events []string) {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind", EventsSendBind{Events: events})
  }
}


// BindPush emits a socket bind push event which adds allowed events to socket
func (service *EventsService) BindPush(events []string) {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind:push", EventsSendBind{Events: events})
  }
}


// BindPop emits a socket bind pop event which removes allowed events from socket
func (service *EventsService) BindPop(events []string) {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind:pop", EventsSendBind{Events: events})
  }
}


// Listen starts listening for incoming realtime events.
func (service *EventsService) Listen(events []string, handleDone func(*EventsRegister)) {
  connectorChild := 0
  connectedSocket := false

  service.connect(events, handleDone, &connectorChild, &connectedSocket)
}
