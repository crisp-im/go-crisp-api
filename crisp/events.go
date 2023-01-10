// Copyright 2018 Crisp IM SAS All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "github.com/graarh/golang-socketio"
  "github.com/graarh/golang-socketio/transport"
  "time"
  "strconv"
  "net/url"
  "encoding/json"
  "encoding/hex"
  "math/rand"
  "crypto/hmac"
  "crypto/sha256"
  "sync"
  "errors"
  "fmt"
)


type eventsLoopback struct {
  register  *EventsRegister
  events    map[string]bool
}

const (
  reconnectWaitBaseline = 4
  reconnectWaitRandomization = 8
  socketioURLQuery = "?EIO=3&transport=websocket"
)

var activeSocket    *gosocketio.Client
var activeLoopback  *eventsLoopback


// EventsMode mode
type EventsMode uint8

const (
  EventsModeWebSockets  EventsMode = 0
  EventsModeWebHooks    EventsMode = 1
)


// EventsService service
type EventsService service

// EventsRegister stores event handlers
type EventsRegister struct {
  handlers      map[string]*caller
  handlersLock  sync.RWMutex
  fnRaiseError  *func(err error)
}

// EventsPayloadHook model mapping
type EventsPayloadHook struct {
  Event  string            `json:"event"`
  Data   *json.RawMessage  `json:"data"`
}

// EventsPullEndpointsData model mapping
type EventsPullEndpointsData struct {
  Data  *EventsPullEndpoints  `json:"data,omitempty"`
}

// EventsPullEndpoints model mapping
type EventsPullEndpoints struct {
  Socket  *EventsPullEndpointsSocket  `json:"socket,omitempty"`
}

// EventsPullEndpointsSocket model mapping
type EventsPullEndpointsSocket struct {
  App  *string  `json:"app,omitempty"`
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

// EventsPluginGeneric maps a generic plugin
type EventsPluginGeneric struct {
  PluginID  *string  `json:"plugin_id"`
}

// EventsImportGeneric maps a generic import
type EventsImportGeneric struct {
  ImportID  *string  `json:"import_id"`
}

// EventsSessionGenericUnbound maps a generic session (unbound from website)
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
  From         *string                              `json:"from"`
  Origin       *string                              `json:"origin"`
  Mentions     *[]string                            `json:"mentions"`
  Stamped      *bool                                `json:"stamped"`
  Timestamp    *uint64                              `json:"timestamp"`
  Fingerprint  *int                                 `json:"fingerprint"`
  User         *EventsReceiveCommonMessageUser      `json:"user"`
  Original     *EventsReceiveCommonMessageOriginal  `json:"original"`
  Edited       *bool                                `json:"edited,omitempty"`
  Translated   *bool                                `json:"translated,omitempty"`
}

// EventsReceiveCommonMessageUser maps a message user
type EventsReceiveCommonMessageUser struct {
  Type      *string  `json:"type"`
  UserID    *string  `json:"user_id"`
  Nickname  *string  `json:"nickname"`
  Avatar    *string  `json:"avatar"`
}

// EventsReceiveCommonMessageOriginal maps a message original
type EventsReceiveCommonMessageOriginal struct {
  OriginalID  *string  `json:"original_id"`
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

// EventsReceiveSessionSetSubject maps session:set_subject
type EventsReceiveSessionSetSubject struct {
  EventsGeneric
  EventsSessionGeneric
  Subject  *string  `json:"subject"`
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
  Timestamp     *uint64  `json:"timestamp"`
}

// EventsReceiveSessionSyncEvents maps session:sync:events
type EventsReceiveSessionSyncEvents struct {
  EventsGeneric
  EventsSessionGeneric
  Events  *[]EventsReceiveSessionSyncEventsOne  `json:"events"`
}

// EventsReceiveSessionSyncEventsOne maps session:sync:events/events
type EventsReceiveSessionSyncEventsOne struct {
  Text       *string       `json:"text"`
  Data       *interface{}  `json:"data,omitempty"`
  Color      *string       `json:"color,omitempty"`
  Timestamp  *uint64       `json:"timestamp"`
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

// EventsReceiveSessionSyncRating maps session:sync:rating
type EventsReceiveSessionSyncRating struct {
  EventsGeneric
  EventsSessionGeneric
  Rating  *EventsReceiveSessionSyncRatingData  `json:"rating"`
}

// EventsReceiveSessionSyncRatingData maps session:sync:rating/rating
type EventsReceiveSessionSyncRatingData struct {
  Stars    uint8    `json:"stars"`
  Comment  *string  `json:"comment"`
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

// EventsReceiveSessionSetParticipants maps session:set_participants
type EventsReceiveSessionSetParticipants struct {
  EventsGeneric
  EventsSessionGeneric
  Participants  *[]EventsReceiveSessionSetParticipantsOne  `json:"participants"`
}

// EventsReceiveSessionSetParticipantsOne maps session:set_participants/participants
type EventsReceiveSessionSetParticipantsOne struct {
  Type    *string  `json:"type"`
  Target  *string  `json:"target"`
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
  Value     string                                          `json:"value"`
  Icon      *string                                         `json:"icon,omitempty"`
  Label     string                                          `json:"label"`
  Selected  bool                                            `json:"selected"`
  Action    *EventsReceivePickerMessageContentChoiceAction  `json:"action,omitempty"`
}

// EventsReceivePickerMessageContentChoiceAction maps message:{send,received}/content/choices/action (picker type)
type EventsReceivePickerMessageContentChoiceAction struct {
  Type    string  `json:"type"`
  Target  string  `json:"target"`
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

// EventsReceiveCarouselMessage maps message:{send,received} (carousel type)
type EventsReceiveCarouselMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceiveCarouselMessageContent  `json:"content"`
}

// EventsReceiveCarouselMessageContent maps message:{send,received}/content (carousel type)
type EventsReceiveCarouselMessageContent struct {
  Text     string                                        `json:"text"`
  Targets  *[]EventsReceiveCarouselMessageContentTarget  `json:"targets"`
}

// EventsReceiveCarouselMessageContentTarget maps message:{send,received}/content/targets (carousel type)
type EventsReceiveCarouselMessageContentTarget struct {
  Title        string                                              `json:"title"`
  Description  string                                              `json:"description"`
  Image        *string                                             `json:"image,omitempty"`
  Actions      *[]EventsReceiveCarouselMessageContentTargetAction  `json:"actions"`
}

// EventsReceiveCarouselMessageContentTargetAction maps message:{send,received}/content/targets/actions (carousel type)
type EventsReceiveCarouselMessageContentTargetAction struct {
  Label  string  `json:"label"`
  URL    string  `json:"url"`
}

// EventsReceiveNoteMessage maps message:{send,received} (note type)
type EventsReceiveNoteMessage EventsReceiveTextMessage

// EventsReceiveEventMessage maps message:{send,received} (event type)
type EventsReceiveEventMessage struct {
  EventsGeneric
  EventsReceiveGenericMessage
  Content  *EventsReceiveEventMessageContent  `json:"content"`
}

// EventsReceiveMessageRemoved maps message:removed
type EventsReceiveMessageRemoved struct {
  EventsGeneric
  EventsSessionGeneric
  Fingerprint  *int  `json:"fingerprint"`
}

// EventsReceiveEventMessageContent maps message:{send,received}/content (event type)
type EventsReceiveEventMessageContent struct {
  Namespace  string   `json:"namespace"`
  Text       *string  `json:"text,omitempty"`
}

// EventsReceiveMessageComposeSend maps message:compose:send
type EventsReceiveMessageComposeSend struct {
  EventsGeneric
  EventsSessionGeneric
  Type       *string  `json:"type"`
  Excerpt    *string  `json:"excerpt"`
  Timestamp  *uint64  `json:"timestamp"`
}

// EventsReceiveMessageComposeReceive maps message:compose:receive
type EventsReceiveMessageComposeReceive struct {
  EventsGeneric
  EventsSessionGeneric
  Type       *string                                  `json:"type"`
  Excerpt    *string                                  `json:"excerpt"`
  Timestamp  *uint64                                  `json:"timestamp"`
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

// EventsReceivePeopleProfileUpdated maps people:profile:updated
type EventsReceivePeopleProfileUpdated struct {
  EventsGeneric
  EventsPeopleGeneric
  Email   *string             `json:"email"`
  Update  *PeopleProfileCard  `json:"update"`
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
  Progress  *uint8                             `json:"progress"`
  Count     *EventsPeopleImportProgressCount   `json:"count"`
  Items     *[]EventsPeopleImportProgressItem  `json:"items"`
}

// EventsPeopleImportProgressCount maps people:import:progress/count
type EventsPeopleImportProgressCount struct {
  Total      *uint32  `json:"total"`
  Skipped    *uint32  `json:"skipped"`
  Remaining  *uint32  `json:"remaining"`
}

// EventsPeopleImportProgressItem maps people:import:progress/items
type EventsPeopleImportProgressItem struct {
  Identifier  *string  `json:"identifier,omitempty"`
  Reason      *string  `json:"reason"`
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
  BrowsingToken  *string  `json:"browsing_token"`
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

// EventsWidgetActionProcessed maps widget:action:processed
type EventsWidgetActionProcessed struct {
  EventsGeneric
  EventsSessionGeneric
  EventsPluginGeneric
  CorrelationID  *string       `json:"correlation_id"`
  Outcome        *string       `json:"outcome"`
  Result         *interface{}  `json:"result,omitempty"`
}

// EventsStatusHealthChanged maps status:health:changed
type EventsStatusHealthChanged struct {
  EventsGeneric
  EventsWebsiteGeneric
  Health  *string                           `json:"health"`
  Nodes   *[]EventsStatusHealthChangedNode  `json:"nodes"`
}

// EventsStatusHealthChangedNode maps status:health:changed/nodes
type EventsStatusHealthChangedNode struct {
  Label    *string  `json:"label,omitempty"`
  Replica  *string  `json:"replica,omitempty"`
}

// EventsReceiveWebsiteUpdateVisitorsCount maps website:update_visitors_count
type EventsReceiveWebsiteUpdateVisitorsCount struct {
  EventsGeneric
  EventsWebsiteGeneric
  VisitorsCount  *uint32  `json:"visitors_count"`
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

// EventsReceiveBucketURLUploadGenerated maps bucket:url:upload:generated
type EventsReceiveBucketURLUploadGenerated struct {
  EventsGeneric
  Type        *string                                         `json:"type"`
  From        *string                                         `json:"from"`
  Identifier  *string                                         `json:"identifier"`
  ID          *string                                         `json:"id"`
  Resource    *EventsReceiveBucketURLUploadGeneratedResource  `json:"resource"`
  URL         *EventsReceiveBucketURLUploadGeneratedURL       `json:"url"`
  Policy      *EventsReceiveBucketURLUploadGeneratedPolicy    `json:"policy"`
}

// EventsReceiveBucketURLUploadGeneratedResource maps bucket:url:upload:generated/resource
type EventsReceiveBucketURLUploadGeneratedResource struct {
  Type  *string  `json:"type"`
  ID    *string  `json:"id"`
}

// EventsReceiveBucketURLUploadGeneratedURL maps bucket:url:upload:generated/url
type EventsReceiveBucketURLUploadGeneratedURL struct {
  Resource  *string  `json:"resource"`
  Signed    *string  `json:"signed"`
}

// EventsReceiveBucketURLUploadGeneratedPolicy maps bucket:url:upload:generated/policy
type EventsReceiveBucketURLUploadGeneratedPolicy struct {
  SizeLimit  *uint64  `json:"size_limit"`
}

// EventsReceiveBucketURLAvatarGenerated maps bucket:url:avatar:generated
type EventsReceiveBucketURLAvatarGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLWebsiteGenerated maps bucket:url:website:generated
type EventsReceiveBucketURLWebsiteGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLCampaignGenerated maps bucket:url:campaign:generated
type EventsReceiveBucketURLCampaignGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLHelpdeskGenerated maps bucket:url:helpdesk:generated
type EventsReceiveBucketURLHelpdeskGenerated EventsReceiveBucketURLUploadGenerated

// EventsReceiveBucketURLStatusGenerated maps bucket:url:status:generated
type EventsReceiveBucketURLStatusGenerated EventsReceiveBucketURLUploadGenerated

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

// EventsReceivePluginChannel maps plugin:channel
type EventsReceivePluginChannel struct {
  EventsGeneric
  EventsWebsiteGeneric
  EventsPluginGeneric
  Namespace   *string       `json:"namespace"`
  Identifier  *string       `json:"identifier,omitempty"`
  Payload     *interface{}  `json:"payload"`
}

// EventsReceivePluginEvent maps plugin:event
type EventsReceivePluginEvent struct {
  EventsGeneric
  EventsWebsiteGeneric
  EventsPluginGeneric
  URN   *string       `json:"urn"`
  Name  *string       `json:"name"`
  Data  *interface{}  `json:"data"`
}

// EventsReceivePluginSettingsSaved maps plugin:settings:saved
type EventsReceivePluginSettingsSaved struct {
  EventsGeneric
  EventsWebsiteGeneric
  EventsPluginGeneric
  Settings  *interface{}  `json:"settings"`
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


// String returns the string representation of EventsReceiveSessionSyncRating
func (evt EventsReceiveSessionSyncRating) String() string {
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


// String returns the string representation of EventsReceiveSessionSetParticipants
func (evt EventsReceiveSessionSetParticipants) String() string {
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


// String returns the string representation of EventsReceiveCarouselMessage
func (evt EventsReceiveCarouselMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveNoteMessage
func (evt EventsReceiveNoteMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveEventMessage
func (evt EventsReceiveEventMessage) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveMessageRemoved
func (evt EventsReceiveMessageRemoved) String() string {
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


// String returns the string representation of EventsReceivePeopleProfileUpdated
func (evt EventsReceivePeopleProfileUpdated) String() string {
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


// String returns the string representation of EventsWidgetActionProcessed
func (evt EventsWidgetActionProcessed) String() string {
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


// String returns the string representation of EventsReceiveBucketURLUploadGenerated
func (evt EventsReceiveBucketURLUploadGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBucketURLAvatarGenerated
func (evt EventsReceiveBucketURLAvatarGenerated) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceiveBucketURLWebsiteGenerated
func (evt EventsReceiveBucketURLWebsiteGenerated) String() string {
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


// String returns the string representation of EventsReceiveBucketURLStatusGenerated
func (evt EventsReceiveBucketURLStatusGenerated) String() string {
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


// String returns the string representation of EventsReceivePluginChannel
func (evt EventsReceivePluginChannel) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceivePluginEvent
func (evt EventsReceivePluginEvent) String() string {
  return Stringify(evt)
}


// String returns the string representation of EventsReceivePluginSettingsSaved
func (evt EventsReceivePluginSettingsSaved) String() string {
  return Stringify(evt)
}


func (register *EventsRegister) segmentEventName(event string, rawData *json.RawMessage) (string, error) {
  switch event {
  case "message:send", "message:received":
    // Messages event names are segmented (send + received)
    var messageGenericType EventsReceiveGenericMessageType

    if decodeError := json.Unmarshal(*rawData, &messageGenericType); decodeError != nil {
      return "", decodeError
    }

    if messageGenericType.Type == nil || *messageGenericType.Type == "" {
      return "", errors.New("empty message type for segment")
    }

    return event + "/" + *messageGenericType.Type, nil
  default:
    // Non-segmented event, return identity function
    return event, nil
  }
}


func (register *EventsRegister) emit(event string, rawData *json.RawMessage) (bool, error) {
  // Acquire segmented event name (eg. certain events as 'namespace' may become 'namespace/segment')
  eventSegmented, segmentErr := register.segmentEventName(event, rawData)

  if segmentErr != nil {
    return false, segmentErr
  }

  // Acquire event handler
  register.handlersLock.RLock()
  fnHandle, ok := register.handlers[eventSegmented]
  register.handlersLock.RUnlock()

  // Route to event handler?
  // Important: go asynchronous when emitting so that event processing does not block parents
  if ok == true {
    // Decode raw data
    data := fnHandle.getArgs()

    if decodeErr := json.Unmarshal(*rawData, &data); decodeErr != nil {
      return false, decodeErr
    }

    // Route to handler
    go fnHandle.callFunc(data)

    return true, nil
  }

  return false, nil
}


func (service *EventsService) pullRemoteEndpointURL() (string, error) {
  var (
    protocol string
    port int
    endpointURL string
  )

  // Fetch endpoint URLs from remote (for tier)
  endpointsUrl := fmt.Sprintf("%s/connect/endpoints", service.client.auth.Tier)
  endpointsReq, _ := service.client.NewRequest("GET", endpointsUrl, nil)

  endpoints := new(EventsPullEndpointsData)
  _, err := service.client.Do(endpointsReq, endpoints)

  if err != nil {
    return "", err
  }

  if endpoints.Data != nil && endpoints.Data.Socket != nil && endpoints.Data.Socket.App != nil {
    endpointURL = *endpoints.Data.Socket.App
  }

  if endpointURL == "" {
    return "", errors.New("empty endpoint url")
  }

  // Generate WebSocket URL
  u, _ := url.Parse(endpointURL)
  host := u.Host
  path := u.Path

  if path == "" {
    path = "/"
  }

  if u.Scheme == "https" || u.Scheme == "wss" {
    protocol = "wss"
    port = 443
  } else {
    protocol = "ws"
    port = 80
  }

  wsEndpointUrl := protocol + "://" + host + ":" + strconv.Itoa(port) + path + socketioURLQuery;

  return wsEndpointUrl, nil;
}


func (service *EventsService) reconnect(events []string, handleConnected func(*EventsRegister), handleDisconnected func(), handleError func(err error), connectorChild *int, connectedSocket *bool, endpointURL *string, child int) {
  // Attempt to reconnect
  for *connectedSocket == false && child == *connectorChild {
    // Compute a random wait interval
    reconnectWait := reconnectWaitBaseline + rand.Intn(reconnectWaitRandomization)

    // Hold on
    time.Sleep(time.Duration(reconnectWait) * time.Second)

    *connectorChild++

    service.connect(events, handleConnected, handleDisconnected, handleError, connectorChild, connectedSocket, endpointURL)
  }
}


func (service *EventsService) connect(events []string, handleConnected func(*EventsRegister), handleDisconnected func(), handleError func(err error), connectorChild *int, connectedSocket *bool, endpointURL *string) {
  child := *connectorChild

  // Pull endpoint URL? (for the first time if not set)
  if *endpointURL == "" {
    pulledEndpointURL, pullErr := service.pullRemoteEndpointURL()

    if pullErr == nil && pulledEndpointURL != "" {
      *endpointURL = pulledEndpointURL
    }
  }

  // Acquire Socket.IO client? (endpoint URL is available)
  var (
    so *gosocketio.Client
    dialErr error
  )

  if *endpointURL != "" {
    so, dialErr = gosocketio.Dial(*endpointURL, transport.GetDefaultWebsocketTransport())
  }

  // Listen for events? (socket has been dialed)
  if dialErr == nil && so != nil {
    so.On("authenticated", func(_ *gosocketio.Channel, authenticated bool) {
      if authenticated == true {
        reg := EventsRegister{handlers: make(map[string]*caller), fnRaiseError: &handleError}

        activeSocket = so

        // Bind all listened-for events
        for _, event := range events {
          // Important: copy event name in a local (as we are in a loop binding an asynchronous handler)
          boundEvent := event

          so.On(event, func(_ *gosocketio.Channel, evt *json.RawMessage) {
            // Dispatch event to event bus
            // Important: emit method is already asynchronous
            _, routeErr := reg.emit(boundEvent, evt)

            // Raise error in asynchronous error handler?
            if routeErr != nil && reg.fnRaiseError != nil {
              (*reg.fnRaiseError)(fmt.Errorf("[socket->%s] %w on payload ~> %s", boundEvent, routeErr, *evt))
            }
          })
        }

        handleConnected(&reg)
      }
    })

    so.On(gosocketio.OnError, func(_ *gosocketio.Channel) {
      handleError(errors.New("socket error"))
    })

    so.On(gosocketio.OnDisconnection, func(_ *gosocketio.Channel) {
      *connectedSocket = false

      handleDisconnected()

      service.reconnect(events, handleConnected, handleDisconnected, handleError, connectorChild, connectedSocket, endpointURL, child)
    })

    so.On(gosocketio.OnConnection, func(_ *gosocketio.Channel) {
      *connectedSocket = true

      // Authenticate to socket
      if service.client.auth.Available == true {
        so.Channel.Emit("authentication", EventsSendAuthentication{Tier: service.client.auth.Tier, Username: service.client.auth.Username, Password: service.client.auth.Password, Events: events})
      }
    })
  } else {
    // Reconnect later (might re-pull endpoint URL and/or re-dial socket)
    service.reconnect(events, handleConnected, handleDisconnected, handleError, connectorChild, connectedSocket, endpointURL, child)
  }
}


func (service *EventsService) loopback(events []string, handleBound func(*EventsRegister), handleError func(err error)) {
  // Prepare register
  reg := EventsRegister{handlers: make(map[string]*caller), fnRaiseError: &handleError}

  // Prepare events map (used for optimized lookup)
  var eventsMap = make(map[string]bool)

  for _, event := range events {
    eventsMap[event] = true
  }

  // Bind loopback
  lo := eventsLoopback{
    register: &reg,
    events: eventsMap,
  }

  activeLoopback = &lo

  // Mark as bound
  handleBound(&reg)
}


// On registers an event handler on event name
func (register *EventsRegister) On(eventName string, handler interface{}) error {
  c, err := newCaller(handler)
  if err != nil {
    return err
  }

  register.handlersLock.Lock()
  register.handlers[eventName] = c
  register.handlersLock.Unlock()

  return nil
}


// Receive receives a raw event and dispatches it to the listener (used for Web Hooks)
func (service *EventsService) ReceiveHook(payload *[]byte) (bool, error) {
  if activeLoopback != nil {
    // Ensure payload is readable
    if payload == nil || len(*payload) == 0 {
      return false, errors.New("empty hook payload")
    }

    // Parse event name and data contained in payload
    var parsedPayload EventsPayloadHook

    parseErr := json.Unmarshal(*payload, &parsedPayload)

    if parseErr != nil || parsedPayload.Event == "" || parsedPayload.Data == nil {
      return false, errors.New("malformatted hook payload")
    }

    // Check if event is subscribed to? (in routing table)
    // Notice: if not in routing table, then silently discard the event w/o any error, as we do not want an HTTP failure status to be sent in response by the implementor.
    if _, ok := activeLoopback.events[parsedPayload.Event]; ok == false {
      return false, nil
    }

    // Dispatch event to event bus
    // Important: emit method is already asynchronous, as the hook might be received synchronously over HTTP
    if activeLoopback.register == nil {
      return false, errors.New("could not emit to register")
    }

    hasRouted, routeErr := activeLoopback.register.emit(parsedPayload.Event, parsedPayload.Data)

    // Raise error in asynchronous error handler?
    if routeErr != nil && activeLoopback.register.fnRaiseError != nil {
      (*activeLoopback.register.fnRaiseError)(fmt.Errorf("[hook->%s] %w on payload ~> %s", parsedPayload.Event, routeErr, *parsedPayload.Data))
    }

    return hasRouted, routeErr
  }

  return false, errors.New("hook loopback not bound")
}


// Verify verifies an event string and checks that signatures match (used for Web Hooks)
func (service *EventsService) VerifyHook(localSecret string, remotePayload *[]byte, remoteRawTimestamp string, remoteSignature string) bool {
  if activeLoopback != nil {
    // Ensure remote payload is readable
    if remotePayload == nil || len(*remotePayload) == 0 {
      return false
    }

    // Decode remote signature
    remoteSignatureBytes, decodeErr := hex.DecodeString(remoteSignature)

    if decodeErr != nil {
      return false
    }

    // Compute local trace
    localTrace := fmt.Sprintf("[%s;%s]", remoteRawTimestamp, string(*remotePayload))

    // Create local HMAC
    localMac := hmac.New(sha256.New, []byte(localSecret))

    localMac.Write([]byte(localTrace))

    // Compute local signature, and compare
    localSignatureBytes := localMac.Sum(nil)

    return hmac.Equal(localSignatureBytes, remoteSignatureBytes)
  }

  // Default: not verified (signatures do not match or loopback not /yet?/ bound)
  return false
}


// Rebind emits an empty socket bind event which associates the socket to its channels, without modifying allowed events (used for WebSockets)
func (service *EventsService) RebindSocket() {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind", "")
  }
}


// Bind emits a socket bind event which associates the socket to its channels, with allowed events (used for WebSockets)
func (service *EventsService) BindSocket(events []string) {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind", EventsSendBind{Events: events})
  }
}


// BindPush emits a socket bind push event which adds allowed events to socket (used for WebSockets)
func (service *EventsService) BindPushSocket(events []string) {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind:push", EventsSendBind{Events: events})
  }
}


// BindPop emits a socket bind pop event which removes allowed events from socket (used for WebSockets)
func (service *EventsService) BindPopSocket(events []string) {
  if activeSocket != nil {
    activeSocket.Channel.Emit("socket:bind:pop", EventsSendBind{Events: events})
  }
}


// Listen starts listening for incoming realtime events (used for both WebSockets and Web Hooks)
func (service *EventsService) Listen(mode EventsMode, events []string, handleBound func(*EventsRegister), handleUnbound func(), handleError func(err error)) (error) {
  if len(events) == 0 {
    return errors.New("no events provided")
  }

  switch mode {
  case EventsModeWebSockets:
    // Storage variables
    connectorChild := 0
    connectedSocket := false
    endpointURL := ""

    // Connect to remote WebSocket server and bind listeners once connected
    service.connect(events, handleBound, handleUnbound, handleError, &connectorChild, &connectedSocket, &endpointURL)

    return nil
  case EventsModeWebHooks:
    // Bind to local loopback, acting as an event bus listening for events dispatched using 'Receive()'
    service.loopback(events, handleBound, handleError)

    return nil
  default:
    return errors.New("invalid events mode")
  }
}
