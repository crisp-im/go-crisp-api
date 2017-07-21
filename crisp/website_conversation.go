// Copyright 2016 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/http"
  "net/url"
)


// ConversationListData mapping
type ConversationListData struct {
  Data  *[]Conversation  `json:"data,omitempty"`
}

// ConversationData mapping
type ConversationData struct {
  Data  *Conversation  `json:"data,omitempty"`
}

// Conversation mapping
type Conversation struct {
  SessionID     *string               `json:"session_id,omitempty"`
  WebsiteID     *string               `json:"website_id,omitempty"`
  PeopleID      *string               `json:"people_id,omitempty"`
  State         *string               `json:"state,omitempty"`
  Status        *uint8                `json:"status,omitempty"`
  IsBlocked     *bool                 `json:"is_blocked,omitempty"`
  Availability  *string               `json:"availability,omitempty"`
  Active        *ConversationActive   `json:"active,omitempty"`
  LastMessage   *string               `json:"last_message,omitempty"`
  Mentions      *[]string             `json:"mentions,omitempty"`
  Compose       *ConversationCompose  `json:"compose,omitempty"`
  Unread        *ConversationUnread   `json:"unread,omitempty"`
  Meta          *ConversationMeta     `json:"meta,omitempty"`
  CreatedAt     *uint                 `json:"created_at,omitempty"`
  UpdatedAt     *uint                 `json:"updated_at,omitempty"`
}

// ConversationActive mapping
type ConversationActive struct {
  Now   *bool  `json:"now,omitempty"`
  Last  *uint  `json:"last,omitempty"`
}

// ConversationCompose mapping
type ConversationCompose struct {
  Operator  *ConversationComposeAtom  `json:"operator,omitempty"`
  Visitor   *ConversationComposeAtom  `json:"visitor,omitempty"`
}

// ConversationComposeAtom mapping
type ConversationComposeAtom struct {
  Type         *string                       `json:"type,omitempty"`
  Excerpt      *string                       `json:"excerpt,omitempty"`
  Timestamp    *uint                         `json:"timestamp,omitempty"`
  User         *ConversationComposeAtomUser  `json:"user,omitempty"`
}

// ConversationComposeAtomUser mapping
type ConversationComposeAtomUser struct {
  UserID    *string  `json:"user_id,omitempty"`
  Nickname  *string  `json:"nickname,omitempty"`
  Avatar    *string  `json:"avatar,omitempty"`
}

// ConversationUnread mapping
type ConversationUnread struct {
  Operator  *uint  `json:"operator,omitempty"`
  Visitor   *uint  `json:"visitor,omitempty"`
}

// ConversationMeta mapping
type ConversationMeta struct {
  Nickname  *string                  `json:"nickname,omitempty"`
  Email     *string                  `json:"email,omitempty"`
  Phone     *string                  `json:"phone,omitempty"`
  Address   *string                  `json:"address,omitempty"`
  IP        *string                  `json:"ip,omitempty"`
  Data      *interface{}             `json:"data,omitempty"`
  Avatar    *string                  `json:"avatar,omitempty"`
  Device    *ConversationMetaDevice  `json:"device,omitempty"`
  Segments  *[]string                `json:"segments,omitempty"`
}

// ConversationPagesData mapping
type ConversationPagesData struct {
  Data  *[]ConversationPage  `json:"data,omitempty"`
}

// ConversationPage mapping
type ConversationPage struct {
  PageTitle     *string  `json:"page_title,omitempty"`
  PageURL       *string  `json:"page_url,omitempty"`
  PageReferrer  *string  `json:"page_referrer,omitempty"`
  Timestamp     *uint    `json:"timestamp,omitempty"`
}

// ConversationEventsData mapping
type ConversationEventsData struct {
  Data  *[]ConversationEvent  `json:"data,omitempty"`
}

// ConversationEvent mapping
type ConversationEvent struct {
  Text       *string       `json:"text,omitempty"`
  Data       *interface{}  `json:"data,omitempty"`
  Color      *string       `json:"color,omitempty"`
  Timestamp  *uint         `json:"timestamp,omitempty"`
}

// ConversationMessage mapping
type ConversationMessage struct {
  SessionID    *string                        `json:"session_id,omitempty"`
  WebsiteID    *string                        `json:"website_id,omitempty"`
  Type         *string                        `json:"type,omitempty"`
  From         *string                        `json:"from,omitempty"`
  Origin       *string                        `json:"origin,omitempty"`
  Content      *interface{}                   `json:"content,omitempty"`
  Preview      *[]ConversationMessagePreview  `json:"preview,omitempty"`
  Mentions     *[]string                      `json:"mentions,omitempty"`
  Read         *string                        `json:"read,omitempty"`
  Delivered    *string                        `json:"delivered,omitempty"`
  Fingerprint  *int                           `json:"fingerprint,omitempty"`
  Timestamp    *uint                          `json:"timestamp,omitempty"`
  User         *ConversationMessageUser       `json:"user,omitempty"`
}

// ConversationMessageTextContent mapping
type ConversationMessageTextContent string

// ConversationMessageFileContent mapping
type ConversationMessageFileContent struct {
  Name  *string  `json:"name"`
  URL   *string  `json:"url"`
  Type  *string  `json:"type"`
}

// ConversationMessageAnimationContent mapping
type ConversationMessageAnimationContent struct {
  URL   *string  `json:"url"`
  Type  *string  `json:"type"`
}

// ConversationMessageAudioContent mapping
type ConversationMessageAudioContent struct {
  URL       *string  `json:"url"`
  Type      *string  `json:"type"`
  Duration  *uint16  `json:"duration"`
}

// ConversationMessagePickerContent mapping
type ConversationMessagePickerContent struct {
  ID       *string                                    `json:"id"`
  Text     *string                                    `json:"text"`
  Choices  *[]ConversationMessagePickerContentChoice  `json:"choices"`
}

// ConversationMessagePickerContentChoice mapping
type ConversationMessagePickerContentChoice struct {
  Value     *string  `json:"value"`
  Label     *string  `json:"label"`
  Selected  *bool    `json:"selected"`
}

// ConversationMessageFieldContent mapping
type ConversationMessageFieldContent struct {
  ID       *string  `json:"id"`
  Text     *string  `json:"text"`
  Explain  *string  `json:"explain"`
  Value    *string  `json:"value"`
}

// ConversationMessageNoteContent mapping
type ConversationMessageNoteContent string

// ConversationMessagePreview mapping
type ConversationMessagePreview struct {
  URL      *string                                 `json:"url,omitempty"`
  Website  *string                                 `json:"website,omitempty"`
  Title    *string                                 `json:"title,omitempty"`
  Preview  *ConversationMessagePreviewInformation  `json:"preview,omitempty"`
  Stamped  *bool                                   `json:"stamped,omitempty"`
}

// ConversationMessagePreviewInformation mapping
type ConversationMessagePreviewInformation struct {
  Excerpt  *string  `json:"excerpt,omitempty"`
  Image    *string  `json:"image,omitempty"`
  Embed    *string  `json:"embed,omitempty"`
}

// ConversationMessageUser mapping
type ConversationMessageUser struct {
  Type      *string  `json:"type,omitempty"`
  UserID    *string  `json:"user_id,omitempty"`
  Nickname  *string  `json:"nickname,omitempty"`
  Avatar    *string  `json:"avatar,omitempty"`
}

// ConversationMetaDevice mapping
type ConversationMetaDevice struct {
  Capabilities  *[]string                           `json:"capabilities,omitempty"`
  Geolocation   *ConversationMetaDeviceGeolocation  `json:"geolocation,omitempty"`
  System        *ConversationMetaDeviceSystem       `json:"system,omitempty"`
  Timezone      *int16                              `json:"timezone,omitempty"`
  Locales       *[]string                           `json:"locales,omitempty"`
}

// ConversationMetaDeviceGeolocation mapping
type ConversationMetaDeviceGeolocation struct {
  Coordinates  *ConversationMetaDeviceGeolocationCoordinates  `json:"coordinates,omitempty"`
  City         *string                                        `json:"city,omitempty"`
  Region       *string                                        `json:"region,omitempty"`
  Country      *string                                        `json:"country,omitempty"`
}

// ConversationMetaDeviceGeolocationCoordinates mapping
type ConversationMetaDeviceGeolocationCoordinates struct {
  Latitude   *float32  `json:"latitude,omitempty"`
  Longitude  *float32  `json:"longitude,omitempty"`
}

// ConversationMetaDeviceSystem mapping
type ConversationMetaDeviceSystem struct {
  OS         *ConversationMetaDeviceSystemOS       `json:"os,omitempty"`
  Engine     *ConversationMetaDeviceSystemEngine   `json:"engine,omitempty"`
  Browser    *ConversationMetaDeviceSystemBrowser  `json:"browser,omitempty"`
  Useragent  *string                               `json:"useragent,omitempty"`
}

// ConversationMetaDeviceSystemOS mapping
type ConversationMetaDeviceSystemOS struct {
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// ConversationMetaDeviceSystemEngine mapping
type ConversationMetaDeviceSystemEngine struct {
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// ConversationMetaDeviceSystemBrowser mapping
type ConversationMetaDeviceSystemBrowser struct {
  Major    *string  `json:"major,omitempty"`
  Version  *string  `json:"version,omitempty"`
  Name     *string  `json:"name,omitempty"`
}

// ConversationMetaSegmentData mapping
type ConversationMetaSegmentData struct {
  Data  *[]ConversationMetaSegment  `json:"data,omitempty"`
}

// ConversationMetaSegment mapping
type ConversationMetaSegment struct {
  Segment  *string  `json:"segment,omitempty"`
  Count    *int32   `json:"count,omitempty"`
}

// ConversationNewData mapping
type ConversationNewData struct {
  Data  *ConversationNew  `json:"data,omitempty"`
}

// ConversationNew mapping
type ConversationNew struct {
  SessionID  *string  `json:"session_id,omitempty"`
}

// ConversationMessagesData mapping
type ConversationMessagesData struct {
  Data  *[]ConversationMessage  `json:"data,omitempty"`
}

// ConversationTextMessageNew mapping
type ConversationTextMessageNew struct {
  Type         string                         `json:"type,omitempty"`
  From         string                         `json:"from,omitempty"`
  Origin       string                         `json:"origin,omitempty"`
  Content      string                         `json:"content,omitempty"`
  Mentions     []string                       `json:"mentions,omitempty"`
  Fingerprint  int                            `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser  `json:"user,omitempty"`
}

// ConversationFileMessageNew mapping
type ConversationFileMessageNew struct {
  Type         string                             `json:"type,omitempty"`
  From         string                             `json:"from,omitempty"`
  Origin       string                             `json:"origin,omitempty"`
  Content      ConversationFileMessageNewContent  `json:"content,omitempty"`
  Mentions     []string                           `json:"mentions,omitempty"`
  Fingerprint  int                                `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser      `json:"user,omitempty"`
}

// ConversationFileMessageNewContent mapping
type ConversationFileMessageNewContent struct {
  Name  string  `json:"name,omitempty"`
  URL   string  `json:"url,omitempty"`
  Type  string  `json:"type,omitempty"`
}

// ConversationAnimationMessageNew mapping
type ConversationAnimationMessageNew struct {
  Type         string                                  `json:"type,omitempty"`
  From         string                                  `json:"from,omitempty"`
  Origin       string                                  `json:"origin,omitempty"`
  Content      ConversationAnimationMessageNewContent  `json:"content,omitempty"`
  Mentions     []string                                `json:"mentions,omitempty"`
  Fingerprint  int                                     `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser           `json:"user,omitempty"`
}

// ConversationAnimationMessageNewContent mapping
type ConversationAnimationMessageNewContent struct {
  URL   string  `json:"url,omitempty"`
  Type  string  `json:"type,omitempty"`
}

// ConversationAudioMessageNew mapping
type ConversationAudioMessageNew struct {
  Type         string                              `json:"type,omitempty"`
  From         string                              `json:"from,omitempty"`
  Origin       string                              `json:"origin,omitempty"`
  Content      ConversationAudioMessageNewContent  `json:"content,omitempty"`
  Mentions     []string                            `json:"mentions,omitempty"`
  Fingerprint  int                                 `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser       `json:"user,omitempty"`
}

// ConversationAudioMessageNewContent mapping
type ConversationAudioMessageNewContent struct {
  URL       string  `json:"url,omitempty"`
  Type      string  `json:"type,omitempty"`
  Duration  uint16  `json:"duration,omitempty"`
}

// ConversationPickerMessageNew mapping
type ConversationPickerMessageNew struct {
  Type         string                               `json:"type,omitempty"`
  From         string                               `json:"from,omitempty"`
  Origin       string                               `json:"origin,omitempty"`
  Content      ConversationPickerMessageNewContent  `json:"content,omitempty"`
  Mentions     []string                             `json:"mentions,omitempty"`
  Fingerprint  int                                  `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser        `json:"user,omitempty"`
}

// ConversationPickerMessageNewContent mapping
type ConversationPickerMessageNewContent struct {
  ID       string                                       `json:"id,omitempty"`
  Text     string                                       `json:"text,omitempty"`
  Choices  []ConversationPickerMessageNewContentChoice  `json:"choices,omitempty"`
}

// ConversationPickerMessageNewContentChoice mapping
type ConversationPickerMessageNewContentChoice struct {
  Value     string  `json:"value,omitempty"`
  Label     string  `json:"label,omitempty"`
  Selected  bool    `json:"selected"`
}

// ConversationFieldMessageNew mapping
type ConversationFieldMessageNew struct {
  Type         string                              `json:"type,omitempty"`
  From         string                              `json:"from,omitempty"`
  Origin       string                              `json:"origin,omitempty"`
  Content      ConversationFieldMessageNewContent  `json:"content,omitempty"`
  Mentions     []string                            `json:"mentions,omitempty"`
  Fingerprint  int                                 `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser       `json:"user,omitempty"`
}

// ConversationFieldMessageNewContent mapping
type ConversationFieldMessageNewContent struct {
  ID       string  `json:"id,omitempty"`
  Text     string  `json:"text,omitempty"`
  Explain  string  `json:"explain,omitempty"`
  Value    string  `json:"value,omitempty"`
}

// ConversationNoteMessageNew mapping
type ConversationNoteMessageNew ConversationTextMessageNew

// ConversationAllMessageNewUser mapping
type ConversationAllMessageNewUser struct {
  Type      string  `json:"type,omitempty"`
  Nickname  string  `json:"nickname,omitempty"`
  Avatar    string  `json:"avatar,omitempty"`
}

// ConversationComposeMessageNew mapping
type ConversationComposeMessageNew struct {
  Type     string  `json:"type,omitempty"`
  From     string  `json:"from,omitempty"`
  Excerpt  string  `json:"excerpt,omitempty"`
}

// ConversationReadMessageMark mapping
type ConversationReadMessageMark struct {
  From          string  `json:"from,omitempty"`
  Origin        string  `json:"origin,omitempty"`
  Fingerprints  []int   `json:"fingerprints,omitempty"`
}

// ConversationDeliveredMessageMark mapping
type ConversationDeliveredMessageMark ConversationReadMessageMark

// ConversationOpenUpdate mapping
type ConversationOpenUpdate struct {
  Opened  *bool  `json:"blocked,omitempty"`
}

// ConversationMetaData mapping
type ConversationMetaData struct {
  Data  *ConversationMeta  `json:"data,omitempty"`
}

// ConversationMetaUpdate mapping
type ConversationMetaUpdate struct {
  Nickname  string                        `json:"nickname,omitempty"`
  Email     string                        `json:"email,omitempty"`
  Avatar    string                        `json:"avatar,omitempty"`
  IP        string                        `json:"ip,omitempty"`
  Phone     string                        `json:"phone,omitempty"`
  Address   string                        `json:"address,omitempty"`
  Segments  []string                      `json:"segments,omitempty"`
  Data      interface{}                   `json:"data,omitempty"`
  Device    *ConversationMetaUpdateDevice `json:"device,omitempty"`
}

// ConversationMetaUpdateDevice mapping
type ConversationMetaUpdateDevice struct {
  Capabilities  []string                                  `json:"capabilities,omitempty"`
  Geolocation   *ConversationMetaUpdateDeviceGeolocation  `json:"geolocation,omitempty"`
  System        *ConversationMetaUpdateDeviceSystem       `json:"system,omitempty"`
  Timezone      int16                                     `json:"timezone,omitempty"`
  Locales       []string                                  `json:"locales,omitempty"`
}

// ConversationMetaUpdateDeviceGeolocation mapping
type ConversationMetaUpdateDeviceGeolocation struct {
  Coordinates  ConversationMetaUpdateDeviceGeolocationCoordinates  `json:"coordinates,omitempty"`
  City         string                                              `json:"city,omitempty"`
  Region       string                                              `json:"region,omitempty"`
  Country      string                                              `json:"country,omitempty"`
}

// ConversationMetaUpdateDeviceGeolocationCoordinates mapping
type ConversationMetaUpdateDeviceGeolocationCoordinates struct {
  Latitude   float32  `json:"latitude,omitempty"`
  Longitude  float32  `json:"longitude,omitempty"`
}

// ConversationMetaUpdateDeviceSystem mapping
type ConversationMetaUpdateDeviceSystem struct {
  OS         ConversationMetaUpdateDeviceSystemOS       `json:"os,omitempty"`
  Engine     ConversationMetaUpdateDeviceSystemEngine   `json:"engine,omitempty"`
  Browser    ConversationMetaUpdateDeviceSystemBrowser  `json:"browser,omitempty"`
  Useragent  string                                     `json:"useragent,omitempty"`
}

// ConversationMetaUpdateDeviceSystemOS mapping
type ConversationMetaUpdateDeviceSystemOS struct {
  Version  string  `json:"version,omitempty"`
  Name     string  `json:"name,omitempty"`
}

// ConversationMetaUpdateDeviceSystemEngine mapping
type ConversationMetaUpdateDeviceSystemEngine struct {
  Version  string  `json:"version,omitempty"`
  Name     string  `json:"name,omitempty"`
}

// ConversationMetaUpdateDeviceSystemBrowser mapping
type ConversationMetaUpdateDeviceSystemBrowser struct {
  Major    string  `json:"major,omitempty"`
  Version  string  `json:"version,omitempty"`
  Name     string  `json:"name,omitempty"`
}

// ConversationStateData mapping
type ConversationStateData struct {
  Data  *ConversationState  `json:"data,omitempty"`
}

// ConversationState mapping
type ConversationState struct {
  State  *string  `json:"state,omitempty"`
}

// ConversationStateUpdate mapping
type ConversationStateUpdate struct {
  State  *string  `json:"state,omitempty"`
}

// ConversationBlockData mapping
type ConversationBlockData struct {
  Data  *ConversationBlock  `json:"data,omitempty"`
}

// ConversationBlock mapping
type ConversationBlock struct {
  Blocked  *bool  `json:"blocked,omitempty"`
}

// ConversationBlockUpdate mapping
type ConversationBlockUpdate struct {
  Blocked  *bool  `json:"blocked,omitempty"`
}

// ConversationTranscriptRequest mapping
type ConversationTranscriptRequest struct {
  To     *string  `json:"to,omitempty"`
  Email  *string  `json:"email,omitempty"`
}

// ConversationBrowsingData mapping
type ConversationBrowsingData struct {
  Data  *[]ConversationBrowsing  `json:"data,omitempty"`
}

// ConversationBrowsing mapping
type ConversationBrowsing struct {
  BrowsingID  *string  `json:"browsing_id,omitempty"`
  Useragent   *string  `json:"useragent,omitempty"`
}

// ConversationBrowsingAction mapping
type ConversationBrowsingAction struct {
  Action  *string  `json:"action,omitempty"`
}

// ConversationBrowsingDebug mapping
type ConversationBrowsingDebug struct {
  Action   *string  `json:"action,omitempty"`
  Execute  *string  `json:"execute,omitempty"`
}

// ConversationBrowsingAssist mapping
type ConversationBrowsingAssist struct {
  Action  *string                            `json:"action,omitempty"`
  Mouse   *ConversationBrowsingAssistMouse   `json:"mouse,omitempty"`
  Scroll  *ConversationBrowsingAssistScroll  `json:"scroll,omitempty"`
}

// ConversationBrowsingAssistMouse mapping
type ConversationBrowsingAssistMouse struct {
  X  *uint16  `json:"x,omitempty"`
  Y  *uint16  `json:"y,omitempty"`
}

// ConversationBrowsingAssistScroll mapping
type ConversationBrowsingAssistScroll struct {
  X  *uint16  `json:"x,omitempty"`
  Y  *uint16  `json:"y,omitempty"`
}

// ConversationCallData mapping
type ConversationCallData struct {
  Data  *ConversationCall  `json:"data,omitempty"`
}

// ConversationCall mapping
type ConversationCall struct {
  CallID  *string  `json:"call_id,omitempty"`
}

// ConversationCallSignalingPayload mapping
type ConversationCallSignalingPayload struct {
  Type     string       `json:"type,omitempty"`
  Payload  interface{}  `json:"payload,omitempty"`
}


// String returns the string representation of Conversation
func (instance Conversation) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationMetaSegment
func (instance ConversationMetaSegment) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationNew
func (instance ConversationNew) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationMessage
func (instance ConversationMessage) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationMeta
func (instance ConversationMeta) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationState
func (instance ConversationState) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationBlock
func (instance ConversationBlock) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationBrowsing
func (instance ConversationBrowsing) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationCall
func (instance ConversationCall) String() string {
  return Stringify(instance)
}


// SearchConversations searches conversations for website.
func (service *WebsiteService) SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error) {
  var resourceURL string

  if searchQuery != "" && searchType != "" {
    resourceURL = fmt.Sprintf("website/%s/conversations/%d?search_query=%s&search_type=%s", websiteID, pageNumber, url.QueryEscape(searchQuery), url.QueryEscape(searchType))
  } else {
    resourceURL = fmt.Sprintf("website/%s/conversations/%d", websiteID, pageNumber)
  }

  req, _ := service.client.NewRequest("GET", resourceURL, nil)

  conversations := new(ConversationListData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// ListConversations lists conversations for website.
func (service *WebsiteService) ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error) {
  return service.SearchConversations(websiteID, pageNumber, "", "")
}


// ListConversationSegmentsInMeta lists conversation segments in meta for website.
func (service *WebsiteService) ListConversationSegmentsInMeta(websiteID string, pageNumber uint) (*[]ConversationMetaSegment, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/meta/segments/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  segments := new(ConversationMetaSegmentData)
  resp, err := service.client.Do(req, segments)
  if err != nil {
    return nil, resp, err
  }

  return segments.Data, resp, err
}


// CreateNewConversation creates a new conversation.
func (service *WebsiteService) CreateNewConversation(websiteID string) (*ConversationNew, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation", websiteID)
  req, _ := service.client.NewRequest("POST", url, nil)

  conversation := new(ConversationNewData)
  resp, err := service.client.Do(req, conversation)
  if err != nil {
    return nil, resp, err
  }

  return conversation.Data, resp, err
}


// CheckConversationExists checks if given conversation session identifier exists.
func (service *WebsiteService) CheckConversationExists(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// GetConversation resolves conversation information and messages.
func (service *WebsiteService) GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  conversation := new(ConversationData)
  resp, err := service.client.Do(req, conversation)
  if err != nil {
    return nil, resp, err
  }

  return conversation.Data, resp, err
}


// RemoveConversation removes a conversation in website.
func (service *WebsiteService) RemoveConversation(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s", websiteID, sessionID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// InitiateConversationWithExistingSession initiates a conversation from an existing session.
func (service *WebsiteService) InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/initiate", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// GetMessagesInConversationLast resolves messages in an existing conversation (last variant).
func (service *WebsiteService) GetMessagesInConversationLast(websiteID string, sessionID string) (*[]ConversationMessage, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/messages", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  messages := new(ConversationMessagesData)
  resp, err := service.client.Do(req, messages)
  if err != nil {
    return nil, resp, err
  }

  return messages.Data, resp, err
}


// GetMessagesInConversationBefore resolves messages in an existing conversation (before variant).
func (service *WebsiteService) GetMessagesInConversationBefore(websiteID string, sessionID string, timestampBefore uint32) (*[]ConversationMessage, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/messages?timestamp_before=%d", websiteID, sessionID, timestampBefore)
  req, _ := service.client.NewRequest("GET", url, nil)

  messages := new(ConversationMessagesData)
  resp, err := service.client.Do(req, messages)
  if err != nil {
    return nil, resp, err
  }

  return messages.Data, resp, err
}


// SendTextMessageInConversation sends a message in an existing conversation (text variant).
func (service *WebsiteService) SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendFileMessageInConversation sends a message in an existing conversation (file variant).
func (service *WebsiteService) SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendAnimationMessageInConversation sends a message in an existing conversation (animation variant).
func (service *WebsiteService) SendAnimationMessageInConversation(websiteID string, sessionID string, message ConversationAnimationMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendAudioMessageInConversation sends a message in an existing conversation (audio variant).
func (service *WebsiteService) SendAudioMessageInConversation(websiteID string, sessionID string, message ConversationAudioMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendPickerMessageInConversation sends a message in an existing conversation (picker variant).
func (service *WebsiteService) SendPickerMessageInConversation(websiteID string, sessionID string, message ConversationPickerMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendFieldMessageInConversation sends a message in an existing conversation (field variant).
func (service *WebsiteService) SendFieldMessageInConversation(websiteID string, sessionID string, message ConversationFieldMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// SendNoteMessageInConversation sends a message in an existing conversation (note variant).
func (service *WebsiteService) SendNoteMessageInConversation(websiteID string, sessionID string, message ConversationNoteMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  return service.client.Do(req, nil)
}


// ComposeMessageInConversation starts or stop composing a message in an existing conversation.
func (service *WebsiteService) ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/compose", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, compose)

  return service.client.Do(req, nil)
}


// MarkMessagesReadInConversation marks messages as read in conversation. Either using given message fingerprints, or all messages.
func (service *WebsiteService) MarkMessagesReadInConversation(websiteID string, sessionID string, read ConversationReadMessageMark) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/read", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, read)

  return service.client.Do(req, nil)
}


// MarkMessagesDeliveredInConversation marks messages as delivered in conversation. Either using given message fingerprints, or all messages.
func (service *WebsiteService) MarkMessagesDeliveredInConversation(websiteID string, sessionID string, delivered ConversationDeliveredMessageMark) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/delivered", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, delivered)

  return service.client.Do(req, nil)
}


// UpdateConversationOpenState updates conversation open state for authenticated operator user
func (service *WebsiteService) UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/open", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationOpenUpdate{&opened})

  return service.client.Do(req, nil)
}


// GetConversationMetas resolves conversation meta information.
func (service *WebsiteService) GetConversationMetas(websiteID string, sessionID string) (*ConversationMeta, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/meta", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  meta := new(ConversationMetaData)
  resp, err := service.client.Do(req, meta)
  if err != nil {
    return nil, resp, err
  }

  return meta.Data, resp, err
}


// UpdateConversationMetas updates conversation meta information.
func (service *WebsiteService) UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/meta", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, metas)

  return service.client.Do(req, nil)
}


// ListConversationPages lists browsed pages in conversation.
func (service *WebsiteService) ListConversationPages(websiteID string, sessionID string, pageNumber uint) (*[]ConversationPage, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/pages/%d", websiteID, sessionID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  pages := new(ConversationPagesData)
  resp, err := service.client.Do(req, pages)
  if err != nil {
    return nil, resp, err
  }

  return pages.Data, resp, err
}


// ListConversationEvents lists stacked events in conversation.
func (service *WebsiteService) ListConversationEvents(websiteID string, sessionID string, pageNumber uint) (*[]ConversationEvent, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/events/%d", websiteID, sessionID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  events := new(ConversationEventsData)
  resp, err := service.client.Do(req, events)
  if err != nil {
    return nil, resp, err
  }

  return events.Data, resp, err
}


// GetConversationState resolves conversation state.
func (service *WebsiteService) GetConversationState(websiteID string, sessionID string) (*ConversationState, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/state", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  state := new(ConversationStateData)
  resp, err := service.client.Do(req, state)
  if err != nil {
    return nil, resp, err
  }

  return state.Data, resp, err
}


// ChangeConversationState updates conversation state.
func (service *WebsiteService) ChangeConversationState(websiteID string, sessionID string, state string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/state", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationStateUpdate{&state})

  return service.client.Do(req, nil)
}


// GetBlockStatusForConversation resolves conversation block status.
func (service *WebsiteService) GetBlockStatusForConversation(websiteID string, sessionID string) (*ConversationBlock, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/block", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  block := new(ConversationBlockData)
  resp, err := service.client.Do(req, block)
  if err != nil {
    return nil, resp, err
  }

  return block.Data, resp, err
}


// BlockIncomingMessagesForConversation blocks further incoming messages from a conversation.
func (service *WebsiteService) BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/block", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationBlockUpdate{&blocked})

  return service.client.Do(req, nil)
}


// RequestEmailTranscriptForConversation requests an email transcript for a conversation.
func (service *WebsiteService) RequestEmailTranscriptForConversation(websiteID string, sessionID string, to string, email string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/transcript", websiteID, sessionID)

  var req *http.Request

  transcriptRequest := ConversationTranscriptRequest{To: &to}

  if email != "" {
    transcriptRequest.Email = &email
  }

  req, _ = service.client.NewRequest("POST", url, transcriptRequest)

  return service.client.Do(req, nil)
}


// ListBrowsingSessionsForConversation lists available browsing sessions for conversation.
func (service *WebsiteService) ListBrowsingSessionsForConversation(websiteID string, sessionID string) (*[]ConversationBrowsing, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/browsing", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  browsing := new(ConversationBrowsingData)
  resp, err := service.client.Do(req, browsing)
  if err != nil {
    return nil, resp, err
  }

  return browsing.Data, resp, err
}


// InitiateBrowsingSessionForConversation initiates browsing session for conversation.
func (service *WebsiteService) InitiateBrowsingSessionForConversation(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/browsing", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// SendActionToExistingBrowsingSession sends an action to an existing browsing session.
func (service *WebsiteService) SendActionToExistingBrowsingSession(websiteID string, sessionID string, browsingID string, action string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/browsing/%s", websiteID, sessionID, browsingID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationBrowsingAction{&action})

  return service.client.Do(req, nil)
}


// DebugExistingBrowsingSession debugs an existing browsing session.
func (service *WebsiteService) DebugExistingBrowsingSession(websiteID string, sessionID string, browsingID string, debug ConversationBrowsingDebug) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/browsing/%s/debug", websiteID, sessionID, browsingID)
  req, _ := service.client.NewRequest("PATCH", url, debug)

  return service.client.Do(req, nil)
}


// AssistExistingBrowsingSession assists an existing browsing session.
func (service *WebsiteService) AssistExistingBrowsingSession(websiteID string, sessionID string, browsingID string, assist ConversationBrowsingAssist) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/browsing/%s/assist", websiteID, sessionID, browsingID)
  req, _ := service.client.NewRequest("PATCH", url, assist)

  return service.client.Do(req, nil)
}


// InitiateNewCallSessionForConversation initiates a new audio/video call session for conversation.
func (service *WebsiteService) InitiateNewCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/call", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, nil)

  call := new(ConversationCallData)
  resp, err := service.client.Do(req, call)
  if err != nil {
    return nil, resp, err
  }

  return call.Data, resp, err
}


// GetOngoingCallSessionForConversation gets the ongoing audio/video call session for conversation.
func (service *WebsiteService) GetOngoingCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/call", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  call := new(ConversationCallData)
  resp, err := service.client.Do(req, call)
  if err != nil {
    return nil, resp, err
  }

  return call.Data, resp, err
}


// AbortOngoingCallSessionForConversation aborts the ongoing audio/video call session for conversation.
func (service *WebsiteService) AbortOngoingCallSessionForConversation(websiteID string, sessionID string, callID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/call/%s", websiteID, sessionID, callID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// TransmitSignalingOnOngoingCallSession transmits a signaling payload for the ongoing audio/video call session for conversation.
func (service *WebsiteService) TransmitSignalingOnOngoingCallSession(websiteID string, sessionID string, callID string, payload ConversationCallSignalingPayload) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/call/%s", websiteID, sessionID, callID)
  req, _ := service.client.NewRequest("PATCH", url, payload)

  return service.client.Do(req, nil)
}
