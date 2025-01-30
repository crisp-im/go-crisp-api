// Copyright 2018 Crisp IM SAS All rights reserved.
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
  SessionID       *string                      `json:"session_id,omitempty"`
  WebsiteID       *string                      `json:"website_id,omitempty"`
  InboxID         *string                      `json:"inbox_id,omitempty"`
  PeopleID        *string                      `json:"people_id,omitempty"`
  State           *string                      `json:"state,omitempty"`
  Status          *uint8                       `json:"status,omitempty"`
  IsVerified      *bool                        `json:"is_verified,omitempty"`
  IsBlocked       *bool                        `json:"is_blocked,omitempty"`
  Availability    *string                      `json:"availability,omitempty"`
  Active          *ConversationActive          `json:"active,omitempty"`
  LastMessage     *string                      `json:"last_message,omitempty"`
  PreviewMessage  *ConversationPreviewMessage  `json:"preview_message,omitempty"`
  Topic           *string                      `json:"topic,omitempty"`
  Participants    *[]ConversationParticipant   `json:"participants,omitempty"`
  Mentions        *[]string                    `json:"mentions,omitempty"`
  Compose         *ConversationCompose         `json:"compose,omitempty"`
  Unread          *ConversationUnread          `json:"unread,omitempty"`
  Assigned        *ConversationAssigned        `json:"assigned,omitempty"`
  Meta            *ConversationMeta            `json:"meta,omitempty"`
  CreatedAt       *uint64                      `json:"created_at,omitempty"`
  UpdatedAt       *uint64                      `json:"updated_at,omitempty"`
  WaitingSince    *uint64                      `json:"waiting_since,omitempty"`
}

// ConversationActive mapping
type ConversationActive struct {
  Now   *bool    `json:"now,omitempty"`
  Last  *uint64  `json:"last,omitempty"`
}

// ConversationPreviewMessage mapping
type ConversationPreviewMessage struct {
  Type         *string  `json:"type,omitempty"`
  From         *string  `json:"from,omitempty"`
  Excerpt      *string  `json:"excerpt,omitempty"`
  Fingerprint  *int     `json:"fingerprint,omitempty"`
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
  Timestamp    *uint64                       `json:"timestamp,omitempty"`
  User         *ConversationComposeAtomUser  `json:"user,omitempty"`
  Automated    *bool                         `json:"automated,omitempty"`
}

// ConversationComposeAtomUser mapping
type ConversationComposeAtomUser struct {
  UserID    *string  `json:"user_id,omitempty"`
  Nickname  *string  `json:"nickname,omitempty"`
  Avatar    *string  `json:"avatar,omitempty"`
}

// ConversationUnread mapping
type ConversationUnread struct {
  Operator  *uint16  `json:"operator,omitempty"`
  Visitor   *uint16  `json:"visitor,omitempty"`
}

// ConversationAssigned mapping
type ConversationAssigned struct {
  UserID  *string  `json:"user_id,omitempty"`
}

// ConversationMeta mapping
type ConversationMeta struct {
  Nickname    *string                      `json:"nickname,omitempty"`
  Email       *string                      `json:"email,omitempty"`
  Phone       *string                      `json:"phone,omitempty"`
  Address     *string                      `json:"address,omitempty"`
  Subject     *string                      `json:"subject,omitempty"`
  IP          *string                      `json:"ip,omitempty"`
  Connection  *ConversationMetaConnection  `json:"connection,omitempty"`
  Data        *interface{}                 `json:"data,omitempty"`
  Avatar      *string                      `json:"avatar,omitempty"`
  Device      *ConversationMetaDevice      `json:"device,omitempty"`
  Segments    *[]string                    `json:"segments,omitempty"`
}

// ConversationOriginal mapping
type ConversationOriginal struct {
  WebsiteID   *string       `json:"website_id,omitempty"`
  SessionID   *string       `json:"session_id,omitempty"`
  OriginalID  *string       `json:"original_id,omitempty"`
  Type        *string       `json:"type,omitempty"`
  Headers     *interface{}  `json:"headers,omitempty"`
  Content     *string       `json:"content,omitempty"`
  Timestamp   *uint64       `json:"timestamp,omitempty"`
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
  Timestamp     *uint64  `json:"timestamp,omitempty"`
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
  Timestamp  *uint64       `json:"timestamp,omitempty"`
}

// ConversationFilesData mapping
type ConversationFilesData struct {
  Data  *[]ConversationFile  `json:"data,omitempty"`
}

// ConversationFile mapping
type ConversationFile struct {
  Name         *string  `json:"name,omitempty"`
  Type         *string  `json:"type,omitempty"`
  URL          *string  `json:"url,omitempty"`
  Fingerprint  *int     `json:"fingerprint,omitempty"`
  Timestamp    *uint64  `json:"timestamp,omitempty"`
}

// ConversationMessage mapping
type ConversationMessage struct {
  SessionID    *string                                 `json:"session_id,omitempty"`
  WebsiteID    *string                                 `json:"website_id,omitempty"`
  Type         *string                                 `json:"type,omitempty"`
  From         *string                                 `json:"from,omitempty"`
  Origin       *string                                 `json:"origin,omitempty"`
  Content      *interface{}                            `json:"content,omitempty"`
  Preview      *[]ConversationMessagePreview           `json:"preview,omitempty"`
  Mentions     *[]string                               `json:"mentions,omitempty"`
  Read         *string                                 `json:"read,omitempty"`
  Delivered    *string                                 `json:"delivered,omitempty"`
  Ignored      *map[string]ConversationMessageIgnored  `json:"ignored,omitempty"`
  Edited       *bool                                   `json:"edited,omitempty"`
  Translated   *bool                                   `json:"translated,omitempty"`
  Automated    *bool                                   `json:"automated,omitempty"`
  Fingerprint  *int                                    `json:"fingerprint,omitempty"`
  Timestamp    *uint64                                 `json:"timestamp,omitempty"`
  User         *ConversationMessageUser                `json:"user,omitempty"`
  References   *[]ConversationMessageReference         `json:"references,omitempty"`
  Original     *ConversationMessageOriginal            `json:"original,omitempty"`
}

// ConversationMessageDispatchedData mapping
type ConversationMessageDispatchedData struct {
  Data  *ConversationMessageDispatched  `json:"data,omitempty"`
}

// ConversationMessageDispatched mapping
type ConversationMessageDispatched struct {
  Fingerprint  *int  `json:"fingerprint,omitempty"`
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
  ID        *string                                    `json:"id"`
  Text      *string                                    `json:"text"`
  Choices   *[]ConversationMessagePickerContentChoice  `json:"choices"`
  Required  *bool                                      `json:"required,omitempty"`
}

// ConversationMessagePickerContentChoice mapping
type ConversationMessagePickerContentChoice struct {
  Value     *string                                        `json:"value"`
  Icon      *string                                        `json:"icon,omitempty"`
  Label     *string                                        `json:"label"`
  Selected  *bool                                          `json:"selected"`
  Action    *ConversationMessagePickerContentChoiceAction  `json:"action,omitempty"`
}

// ConversationMessagePickerContentChoiceAction mapping
type ConversationMessagePickerContentChoiceAction struct {
  Type    *string  `json:"type"`
  Target  *string  `json:"target"`
}

// ConversationMessageFieldContent mapping
type ConversationMessageFieldContent struct {
  ID        *string  `json:"id"`
  Text      *string  `json:"text"`
  Explain   *string  `json:"explain"`
  Value     *string  `json:"value"`
  Required  *bool    `json:"required,omitempty"`
}

// ConversationMessageCarouselContent mapping
type ConversationMessageCarouselContent struct {
  Text     *string                                      `json:"text"`
  Targets  *[]ConversationMessageCarouselContentTarget  `json:"targets"`
}

// ConversationMessageCarouselContentTarget mapping
type ConversationMessageCarouselContentTarget struct {
  Title        *string                                            `json:"title"`
  Description  *string                                            `json:"description"`
  Image        *string                                            `json:"image,omitempty"`
  Actions      *[]ConversationMessageCarouselContentTargetAction  `json:"actions"`
}

// ConversationMessageCarouselContentTargetAction mapping
type ConversationMessageCarouselContentTargetAction struct {
  Label  *string  `json:"label"`
  URL    *string  `json:"url"`
}

// ConversationMessageNoteContent mapping
type ConversationMessageNoteContent string

// ConversationMessageEventContent mapping
type ConversationMessageEventContent struct {
  Namespace  *string  `json:"namespace"`
  Text       *string  `json:"text,omitempty"`
}

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

// ConversationMessageIgnored mapping
type ConversationMessageIgnored struct {
  Type    *string  `json:"type,omitempty"`
  Reason  *string  `json:"reason,omitempty"`
}

// ConversationMessageUser mapping
type ConversationMessageUser struct {
  Type      *string  `json:"type,omitempty"`
  UserID    *string  `json:"user_id,omitempty"`
  Nickname  *string  `json:"nickname,omitempty"`
  Avatar    *string  `json:"avatar,omitempty"`
}

// ConversationMessageReference mapping
type ConversationMessageReference struct {
  Type    *string  `json:"type,omitempty"`
  Name    *string  `json:"name,omitempty"`
  Target  *string  `json:"target,omitempty"`
}

// ConversationMessageOriginal mapping
type ConversationMessageOriginal struct {
  OriginalID  *string  `json:"original_id,omitempty"`
}

// ConversationMetaConnection mapping
type ConversationMetaConnection struct {
  ISP  *string  `json:"isp,omitempty"`
  ASN  *string  `json:"asn,omitempty"`
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

// ConversationSuggestedSegmentData mapping
type ConversationSuggestedSegmentData struct {
  Data  *[]ConversationSuggestedSegment  `json:"data,omitempty"`
}

// ConversationSuggestedSegment mapping
type ConversationSuggestedSegment struct {
  Segment  *string  `json:"segment,omitempty"`
  Count    *int32   `json:"count,omitempty"`
}

// ConversationSuggestedDataData mapping
type ConversationSuggestedDataData struct {
  Data  *[]ConversationSuggestedData  `json:"data,omitempty"`
}

// ConversationSuggestedData mapping
type ConversationSuggestedData struct {
  Key    *string  `json:"key,omitempty"`
  Count  *int32   `json:"count,omitempty"`
}

// ConversationSpamData mapping
type ConversationSpamData struct {
  Data  *[]ConversationSpam  `json:"data,omitempty"`
}

// ConversationSpam mapping
type ConversationSpam struct {
  SpamID     *string       `json:"spam_id,omitempty"`
  Type       *string       `json:"type,omitempty"`
  Reason     *string       `json:"reason,omitempty"`
  Metadata   *interface{}  `json:"metadata,omitempty"`
  Headers    *interface{}  `json:"headers,omitempty"`
  Timestamp  *uint64       `json:"timestamp,omitempty"`
}

// ConversationSpamContentData mapping
type ConversationSpamContentData struct {
  Data  *ConversationSpamContent  `json:"data,omitempty"`
}

// ConversationSpamContent mapping
type ConversationSpamContent struct {
  ConversationSpam
  Content  *string  `json:"content,omitempty"`
}

// ConversationSpamDecision mapping
type ConversationSpamDecision struct {
  Action  *string  `json:"action,omitempty"`
}

// ConversationRoutingAssignData mapping
type ConversationRoutingAssignData struct {
  Data  *ConversationRoutingAssign  `json:"data,omitempty"`
}

// ConversationRoutingAssign mapping
type ConversationRoutingAssign struct {
  Assigned  *ConversationRoutingAssignAssigned  `json:"assigned,omitempty"`
}

// ConversationRoutingAssignAssigned mapping
type ConversationRoutingAssignAssigned struct {
  UserID  *string  `json:"user_id,omitempty"`
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

// ConversationMessageData mapping
type ConversationMessageData struct {
  Data  *ConversationMessage  `json:"data,omitempty"`
}

// ConversationSuggestedSegmentDelete mapping
type ConversationSuggestedSegmentDelete struct {
  Segment  string  `json:"segment,omitempty"`
}

// ConversationSuggestedDataDelete mapping
type ConversationSuggestedDataDelete struct {
  Key  string  `json:"key,omitempty"`
}

// ConversationFileMessageNewContent mapping
type ConversationFileMessageNewContent struct {
  Name  string  `json:"name,omitempty"`
  URL   string  `json:"url,omitempty"`
  Type  string  `json:"type,omitempty"`
}

// ConversationAnimationMessageNewContent mapping
type ConversationAnimationMessageNewContent struct {
  URL   string  `json:"url,omitempty"`
  Type  string  `json:"type,omitempty"`
}

// ConversationAudioMessageNewContent mapping
type ConversationAudioMessageNewContent struct {
  URL       string  `json:"url,omitempty"`
  Type      string  `json:"type,omitempty"`
  Duration  uint16  `json:"duration,omitempty"`
}

// ConversationPickerMessageNewContent mapping
type ConversationPickerMessageNewContent struct {
  ID        string                                       `json:"id,omitempty"`
  Text      string                                       `json:"text,omitempty"`
  Choices   []ConversationPickerMessageNewContentChoice  `json:"choices,omitempty"`
  Required  *bool                                        `json:"required,omitempty"`
}

// ConversationPickerMessageNewContentChoice mapping
type ConversationPickerMessageNewContentChoice struct {
  Value     string                                            `json:"value,omitempty"`
  Icon      *string                                           `json:"icon,omitempty"`
  Label     string                                            `json:"label,omitempty"`
  Selected  bool                                              `json:"selected"`
  Action    *ConversationPickerMessageNewContentChoiceAction  `json:"action,omitempty"`
}

// ConversationPickerMessageNewContentChoiceAction mapping
type ConversationPickerMessageNewContentChoiceAction struct {
  Type    string  `json:"type"`
  Target  string  `json:"target"`
}

// ConversationFieldMessageNewContent mapping
type ConversationFieldMessageNewContent struct {
  ID        string  `json:"id,omitempty"`
  Text      string  `json:"text,omitempty"`
  Explain   string  `json:"explain,omitempty"`
  Value     string  `json:"value,omitempty"`
  Required  *bool   `json:"required,omitempty"`
}

// ConversationCarouselMessageNewContent mapping
type ConversationCarouselMessageNewContent struct {
  Text     string                                         `json:"text,omitempty"`
  Targets  []ConversationCarouselMessageNewContentTarget  `json:"targets,omitempty"`
}

// ConversationCarouselMessageNewContentTarget mapping
type ConversationCarouselMessageNewContentTarget struct {
  Title        string                                               `json:"title,omitempty"`
  Description  string                                               `json:"description,omitempty"`
  Image        *string                                              `json:"image,omitempty"`
  Actions      []ConversationCarouselMessageNewContentTargetAction  `json:"actions,omitempty"`
}

// ConversationCarouselMessageNewContentTargetAction mapping
type ConversationCarouselMessageNewContentTargetAction struct {
  Label  string  `json:"label,omitempty"`
  URL    string  `json:"url,omitempty"`
}

// ConversationEventMessageNewContent mapping
type ConversationEventMessageNewContent struct {
  Namespace  string   `json:"namespace,omitempty"`
  Text       *string  `json:"text,omitempty"`
}

// ConversationTextMessageNew mapping
type ConversationTextMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      string                                 `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
}

// ConversationFileMessageNew mapping
type ConversationFileMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      ConversationFileMessageNewContent      `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
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
  References   *[]ConversationAllMessageNewReference   `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal      `json:"original,omitempty"`
  Timestamp    *uint64                                 `json:"timestamp,omitempty"`
  Stealth      *bool                                   `json:"stealth,omitempty"`
  Translated   *bool                                   `json:"translated,omitempty"`
  Automated    *bool                                   `json:"automated,omitempty"`
}

// ConversationAudioMessageNew mapping
type ConversationAudioMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      ConversationAudioMessageNewContent     `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
}

// ConversationPickerMessageNew mapping
type ConversationPickerMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      ConversationPickerMessageNewContent    `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
}

// ConversationFieldMessageNew mapping
type ConversationFieldMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      ConversationFieldMessageNewContent     `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
}

// ConversationCarouselMessageNew mapping
type ConversationCarouselMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      ConversationCarouselMessageNewContent  `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
}

// ConversationNoteMessageNew mapping
type ConversationNoteMessageNew ConversationTextMessageNew

// ConversationEventMessageNew mapping
type ConversationEventMessageNew struct {
  Type         string                                 `json:"type,omitempty"`
  From         string                                 `json:"from,omitempty"`
  Origin       string                                 `json:"origin,omitempty"`
  Content      ConversationEventMessageNewContent     `json:"content,omitempty"`
  Mentions     []string                               `json:"mentions,omitempty"`
  Fingerprint  int                                    `json:"fingerprint,omitempty"`
  User         ConversationAllMessageNewUser          `json:"user,omitempty"`
  References   *[]ConversationAllMessageNewReference  `json:"references,omitempty"`
  Original     *ConversationAllMessageNewOriginal     `json:"original,omitempty"`
  Timestamp    *uint64                                `json:"timestamp,omitempty"`
  Stealth      *bool                                  `json:"stealth,omitempty"`
  Translated   *bool                                  `json:"translated,omitempty"`
  Automated    *bool                                  `json:"automated,omitempty"`
}

// ConversationTextMessageUpdate mapping
type ConversationTextMessageUpdate struct {
  Content  string  `json:"content,omitempty"`
}

// ConversationFileMessageUpdate mapping
type ConversationFileMessageUpdate struct {
  Content  ConversationFileMessageNewContent  `json:"content,omitempty"`
}

// ConversationAnimationMessageUpdate mapping
type ConversationAnimationMessageUpdate struct {
  Content  ConversationAnimationMessageNewContent  `json:"content,omitempty"`
}

// ConversationAudioMessageUpdate mapping
type ConversationAudioMessageUpdate struct {
  Content  ConversationAudioMessageNewContent  `json:"content,omitempty"`
}

// ConversationPickerMessageUpdate mapping
type ConversationPickerMessageUpdate struct {
  Content  ConversationPickerMessageNewContent  `json:"content,omitempty"`
}

// ConversationFieldMessageUpdate mapping
type ConversationFieldMessageUpdate struct {
  Content  ConversationFieldMessageNewContent  `json:"content,omitempty"`
}

// ConversationCarouselMessageUpdate mapping
type ConversationCarouselMessageUpdate struct {
  Content  ConversationCarouselMessageNewContent  `json:"content,omitempty"`
}

// ConversationNoteMessageUpdate mapping
type ConversationNoteMessageUpdate ConversationTextMessageUpdate

// ConversationEventMessageUpdate mapping
type ConversationEventMessageUpdate struct {
  Content  ConversationEventMessageNewContent  `json:"content,omitempty"`
}

// ConversationAllMessageNewUser mapping
type ConversationAllMessageNewUser struct {
  Type      string  `json:"type,omitempty"`
  Nickname  string  `json:"nickname,omitempty"`
  Avatar    string  `json:"avatar,omitempty"`
}

// ConversationAllMessageNewReference mapping
type ConversationAllMessageNewReference struct {
  Type    string  `json:"type,omitempty"`
  Name    string  `json:"name,omitempty"`
  Target  string  `json:"target,omitempty"`
}

// ConversationAllMessageNewOriginal mapping
type ConversationAllMessageNewOriginal struct {
  Type     string  `json:"type,omitempty"`
  Content  string  `json:"content,omitempty"`
}

// ConversationComposeMessageNew mapping
type ConversationComposeMessageNew struct {
  Type       string  `json:"type,omitempty"`
  From       string  `json:"from,omitempty"`
  Excerpt    string  `json:"excerpt,omitempty"`
  Stealth    *bool   `json:"stealth,omitempty"`
  Automated  *bool   `json:"automated,omitempty"`
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

// ConversationRoutingAssignUpdate mapping
type ConversationRoutingAssignUpdate struct {
  Assigned  *ConversationRoutingAssignUpdateAssigned  `json:"assigned"`
  Silent    *bool                                     `json:"silent,omitempty"`
}

// ConversationRoutingAssignUpdateAssigned mapping
type ConversationRoutingAssignUpdateAssigned struct {
  UserID  string  `json:"user_id,omitempty"`
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
  Subject   string                        `json:"subject,omitempty"`
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

// ConversationOriginalData mapping
type ConversationOriginalData struct {
  Data  *ConversationOriginal  `json:"data,omitempty"`
}

// ConversationInboxUpdate mapping
type ConversationInboxUpdate struct {
  InboxID  *string  `json:"inbox_id,omitempty"`
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

// ConversationParticipantsData mapping
type ConversationParticipantsData struct {
  Data  *ConversationParticipants  `json:"data,omitempty"`
}

// ConversationParticipants mapping
type ConversationParticipants struct {
  Participants  *[]ConversationParticipant  `json:"participants,omitempty"`
}

// ConversationParticipant mapping
type ConversationParticipant struct {
  Type    *string  `json:"type,omitempty"`
  Target  *string  `json:"target,omitempty"`
}

// ConversationParticipantsSave mapping
type ConversationParticipantsSave struct {
  Participants  *[]ConversationParticipant  `json:"participants,omitempty"`
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

// ConversationVerifyData mapping
type ConversationVerifyData struct {
  Data  *ConversationVerify  `json:"data,omitempty"`
}

// ConversationVerify mapping
type ConversationVerify struct {
  Verified  *bool  `json:"verified,omitempty"`
}

// ConversationVerifyUpdate mapping
type ConversationVerifyUpdate struct {
  Verified  *bool  `json:"verified,omitempty"`
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
  BrowsingID     *string  `json:"browsing_id,omitempty"`
  BrowsingToken  *string  `json:"browsing_token,omitempty"`
  Useragent      *string  `json:"useragent,omitempty"`
}

// ConversationBrowsingAction mapping
type ConversationBrowsingAction struct {
  Action  *string  `json:"action,omitempty"`
}

// ConversationBrowsingAssist mapping
type ConversationBrowsingAssist struct {
  Action  *string                            `json:"action,omitempty"`
  Mouse   *ConversationBrowsingAssistMouse   `json:"mouse,omitempty"`
  Click   *ConversationBrowsingAssistClick   `json:"click,omitempty"`
  Scroll  *ConversationBrowsingAssistScroll  `json:"scroll,omitempty"`
}

// ConversationBrowsingAssistMouse mapping
type ConversationBrowsingAssistMouse struct {
  X  *uint16  `json:"x,omitempty"`
  Y  *uint16  `json:"y,omitempty"`
}

// ConversationBrowsingAssistClick mapping
type ConversationBrowsingAssistClick struct {
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

// ConversationCallPayload mapping
type ConversationCallPayload struct {
  Mode  string  `json:"mode,omitempty"`
}

// ConversationCallSignalingPayload mapping
type ConversationCallSignalingPayload struct {
  Type     string       `json:"type,omitempty"`
  Payload  interface{}  `json:"payload,omitempty"`
}

// ConversationWidgetActionData mapping
type ConversationWidgetActionData struct {
  Data  *ConversationWidgetAction  `json:"data,omitempty"`
}

// ConversationWidgetAction mapping
type ConversationWidgetAction struct {
  CorrelationID  *string                            `json:"correlation_id,omitempty"`
  Strategy       *ConversationWidgetActionStrategy  `json:"strategy,omitempty"`
}

// ConversationWidgetActionStrategy mapping
type ConversationWidgetActionStrategy struct {
  Attempts    *uint8   `json:"attempts,omitempty"`
  Exhaustion  *uint16  `json:"exhaustion,omitempty"`
}

// ConversationWidgetButtonPayload mapping
type ConversationWidgetButtonPayload struct {
  SectionID  string        `json:"section_id,omitempty"`
  ItemID     string        `json:"item_id,omitempty"`
  Data       interface{}   `json:"data,omitempty"`
  Value      *interface{}  `json:"value,omitempty"`
}

// ConversationWidgetDataFetchPayload mapping
type ConversationWidgetDataFetchPayload struct {
  SectionID  string       `json:"section_id,omitempty"`
  ItemID     string       `json:"item_id,omitempty"`
  Action     string       `json:"action,omitempty"`
  Data       interface{}  `json:"data,omitempty"`
}

// ConversationWidgetDataEditPayload mapping
type ConversationWidgetDataEditPayload struct {
  SectionID  string  `json:"section_id,omitempty"`
  ItemID     string  `json:"item_id,omitempty"`
  Action     string  `json:"action,omitempty"`
  Value      string  `json:"value,omitempty"`
}

// ConversationReminderPayload mapping
type ConversationReminderPayload struct {
  Date  string  `json:"date,omitempty"`
  Note  string  `json:"note,omitempty"`
}

// ConversationReportPayload mapping
type ConversationReportPayload struct {
  Flag  string  `json:"flag,omitempty"`
}


// String returns the string representation of Conversation
func (instance Conversation) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationSuggestedSegment
func (instance ConversationSuggestedSegment) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationSuggestedData
func (instance ConversationSuggestedData) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationSpam
func (instance ConversationSpam) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationSpamContent
func (instance ConversationSpamContent) String() string {
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


// String returns the string representation of ConversationOriginal
func (instance ConversationOriginal) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationPage
func (instance ConversationPage) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationEvent
func (instance ConversationEvent) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationFile
func (instance ConversationFile) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationState
func (instance ConversationState) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationParticipants
func (instance ConversationParticipants) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationRoutingAssign
func (instance ConversationRoutingAssign) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationBlock
func (instance ConversationBlock) String() string {
  return Stringify(instance)
}


// String returns the string representation of ConversationVerify
func (instance ConversationVerify) String() string {
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


// String returns the string representation of ConversationWidgetAction
func (instance ConversationWidgetAction) String() string {
  return Stringify(instance)
}


// ListConversations lists conversations for website.
func (service *WebsiteService) ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  conversations := new(ConversationListData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// SearchConversations searches conversations for website.
func (service *WebsiteService) SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/%d?search_query=%s&search_type=%s", websiteID, pageNumber, url.QueryEscape(searchQuery), url.QueryEscape(searchType))
  req, _ := service.client.NewRequest("GET", url, nil)

  conversations := new(ConversationListData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// FilterConversations filters conversations for website.
func (service *WebsiteService) FilterConversations(websiteID string, pageNumber uint, filterInboxID string, filterUnread bool, filterResolved bool, filterNotResolved bool, filterMention bool, filterAssigned bool, filterUnassigned bool) (*[]Conversation, *Response, error) {
  var (
    filterInboxIDValue string
    filterUnreadValue string
    filterResolvedValue string
    filterNotResolvedValue string
    filterMentionValue string
    filterAssignedValue string
    filterUnassignedValue string
  )

  filterInboxIDValue = filterInboxID

  if filterUnread == true {
    filterUnreadValue = "1"
  } else {
    filterUnreadValue = "0"
  }

  if filterResolved == true {
    filterResolvedValue = "1"
  } else {
    filterResolvedValue = "0"
  }

  if filterNotResolved == true {
    filterNotResolvedValue = "1"
  } else {
    filterNotResolvedValue = "0"
  }

  if filterMention == true {
    filterMentionValue = "1"
  } else {
    filterMentionValue = "0"
  }

  if filterAssigned == true {
    filterAssignedValue = "1"
  } else {
    filterAssignedValue = "0"
  }

  if filterUnassigned == true {
    filterUnassignedValue = "1"
  } else {
    filterUnassignedValue = "0"
  }

  url := fmt.Sprintf("website/%s/conversations/%d?filter_inbox_id=%s&filter_unread=%s&filter_resolved=%s&filter_not_resolved=%s&filter_mention=%s&filter_assigned=%s&filter_unassigned=%s", websiteID, pageNumber, url.QueryEscape(filterInboxIDValue), url.QueryEscape(filterUnreadValue), url.QueryEscape(filterResolvedValue), url.QueryEscape(filterNotResolvedValue), url.QueryEscape(filterMentionValue), url.QueryEscape(filterAssignedValue), url.QueryEscape(filterUnassignedValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  conversations := new(ConversationListData)
  resp, err := service.client.Do(req, conversations)
  if err != nil {
    return nil, resp, err
  }

  return conversations.Data, resp, err
}


// ListSuggestedConversationSegments lists suggested conversation segments for website.
func (service *WebsiteService) ListSuggestedConversationSegments(websiteID string, pageNumber uint) (*[]ConversationSuggestedSegment, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/suggest/segments/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  segments := new(ConversationSuggestedSegmentData)
  resp, err := service.client.Do(req, segments)
  if err != nil {
    return nil, resp, err
  }

  return segments.Data, resp, err
}


// DeleteSuggestedConversationSegment deletes a suggested conversation segment for website.
func (service *WebsiteService) DeleteSuggestedConversationSegment(websiteID string, segment string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversations/suggest/segment", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, ConversationSuggestedSegmentDelete{Segment: segment})

  return service.client.Do(req, nil)
}


// ListSuggestedConversationDataKeys lists suggested conversation data keys for website.
func (service *WebsiteService) ListSuggestedConversationDataKeys(websiteID string, pageNumber uint) (*[]ConversationSuggestedData, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/suggest/data/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  data := new(ConversationSuggestedDataData)
  resp, err := service.client.Do(req, data)
  if err != nil {
    return nil, resp, err
  }

  return data.Data, resp, err
}


// DeleteSuggestedConversationDataKey deletes a suggested conversation data key for website.
func (service *WebsiteService) DeleteSuggestedConversationDataKey(websiteID string, key string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversations/suggest/data", websiteID)
  req, _ := service.client.NewRequest("DELETE", url, ConversationSuggestedDataDelete{Key: key})

  return service.client.Do(req, nil)
}


// ListSpamConversations lists spam conversations in website.
func (service *WebsiteService) ListSpamConversations(websiteID string, pageNumber uint) (*[]ConversationSpam, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/spams/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  spams := new(ConversationSpamData)
  resp, err := service.client.Do(req, spams)
  if err != nil {
    return nil, resp, err
  }

  return spams.Data, resp, err
}


// ResolveSpamConversationContent resolves content for spam conversation.
func (service *WebsiteService) ResolveSpamConversationContent(websiteID string, spamID string) (*ConversationSpamContent, *Response, error) {
  url := fmt.Sprintf("website/%s/conversations/spam/%s/content", websiteID, spamID)
  req, _ := service.client.NewRequest("GET", url, nil)

  content := new(ConversationSpamContentData)
  resp, err := service.client.Do(req, content)
  if err != nil {
    return nil, resp, err
  }

  return content.Data, resp, err
}


// SubmitSpamConversationDecision submits decision on spam conversation.
func (service *WebsiteService) SubmitSpamConversationDecision(websiteID string, spamID string, action string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversations/spam/%s/decision", websiteID, spamID)
  req, _ := service.client.NewRequest("POST", url, ConversationSpamDecision{&action})

  return service.client.Do(req, nil)
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
func (service *WebsiteService) GetMessagesInConversationBefore(websiteID string, sessionID string, timestampBefore uint) (*[]ConversationMessage, *Response, error) {
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
func (service *WebsiteService) SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendFileMessageInConversation sends a message in an existing conversation (file variant).
func (service *WebsiteService) SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendAnimationMessageInConversation sends a message in an existing conversation (animation variant).
func (service *WebsiteService) SendAnimationMessageInConversation(websiteID string, sessionID string, message ConversationAnimationMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendAudioMessageInConversation sends a message in an existing conversation (audio variant).
func (service *WebsiteService) SendAudioMessageInConversation(websiteID string, sessionID string, message ConversationAudioMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendPickerMessageInConversation sends a message in an existing conversation (picker variant).
func (service *WebsiteService) SendPickerMessageInConversation(websiteID string, sessionID string, message ConversationPickerMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendFieldMessageInConversation sends a message in an existing conversation (field variant).
func (service *WebsiteService) SendFieldMessageInConversation(websiteID string, sessionID string, message ConversationFieldMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendCarouselMessageInConversation sends a message in an existing conversation (carousel variant).
func (service *WebsiteService) SendCarouselMessageInConversation(websiteID string, sessionID string, message ConversationCarouselMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendNoteMessageInConversation sends a message in an existing conversation (note variant).
func (service *WebsiteService) SendNoteMessageInConversation(websiteID string, sessionID string, message ConversationNoteMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// SendEventMessageInConversation sends a message in an existing conversation (event variant).
func (service *WebsiteService) SendEventMessageInConversation(websiteID string, sessionID string, message ConversationEventMessageNew) (*ConversationMessageDispatched, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, message)

  dispatched := new(ConversationMessageDispatchedData)
  resp, err := service.client.Do(req, dispatched)
  if err != nil {
    return nil, resp, err
  }

  return dispatched.Data, resp, err
}


// GetMessageInConversation resolves an existing message in an existing conversation.
func (service *WebsiteService) GetMessageInConversation(websiteID string, sessionID string, fingerprint int) (*ConversationMessage, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("GET", url, nil)

  message := new(ConversationMessageData)
  resp, err := service.client.Do(req, message)
  if err != nil {
    return nil, resp, err
  }

  return message.Data, resp, err
}


// UpdateTextMessageInConversation edits an existing message in an existing conversation (text variant).
func (service *WebsiteService) UpdateTextMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationTextMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateFileMessageInConversation edits an existing message in an existing conversation (file variant).
func (service *WebsiteService) UpdateFileMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFileMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationFileMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateAnimationMessageInConversation edits an existing message in an existing conversation (animation variant).
func (service *WebsiteService) UpdateAnimationMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAnimationMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationAnimationMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateAudioMessageInConversation edits an existing message in an existing conversation (audio variant).
func (service *WebsiteService) UpdateAudioMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAudioMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationAudioMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdatePickerMessageInConversation edits an existing message in an existing conversation (picker variant).
func (service *WebsiteService) UpdatePickerMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationPickerMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationPickerMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateFieldMessageInConversation edits an existing message in an existing conversation (field variant).
func (service *WebsiteService) UpdateFieldMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFieldMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationFieldMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateCarouselMessageInConversation edits an existing message in an existing conversation (carousel variant).
func (service *WebsiteService) UpdateCarouselMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationCarouselMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationCarouselMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateNoteMessageInConversation edits an existing message in an existing conversation (note variant).
func (service *WebsiteService) UpdateNoteMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationNoteMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// UpdateEventMessageInConversation edits an existing message in an existing conversation (event variant).
func (service *WebsiteService) UpdateEventMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationEventMessageNewContent) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("PATCH", url, ConversationEventMessageUpdate{Content: content})

  return service.client.Do(req, nil)
}


// RemoveMessageInConversation removes an existing message in an existing conversation.
func (service *WebsiteService) RemoveMessageInConversation(websiteID string, sessionID string, fingerprint int) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/message/%d", websiteID, sessionID, fingerprint)
  req, _ := service.client.NewRequest("DELETE", url, nil)

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


// UpdateConversationOpenState updates conversation open state for authenticated operator user.
func (service *WebsiteService) UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/open", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationOpenUpdate{&opened})

  return service.client.Do(req, nil)
}


// GetConversationRoutingAssign resolves assigned operator for conversation routing.
func (service *WebsiteService) GetConversationRoutingAssign(websiteID string, sessionID string) (*ConversationRoutingAssign, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/routing", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  assign := new(ConversationRoutingAssignData)
  resp, err := service.client.Do(req, assign)
  if err != nil {
    return nil, resp, err
  }

  return assign.Data, resp, err
}


// AssignConversationRouting assigns conversation routing to an operator, or unassign.
func (service *WebsiteService) AssignConversationRouting(websiteID string, sessionID string, assign ConversationRoutingAssignUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/routing", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, assign)

  return service.client.Do(req, nil)
}


// UpdateConversationInbox updates inbox used for conversation.
func (service *WebsiteService) UpdateConversationInbox(websiteID string, sessionID string, inboxID *string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/inbox", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationInboxUpdate{inboxID})

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


// GetOriginalMessageInConversation resolves conversation original message.
func (service *WebsiteService) GetOriginalMessageInConversation(websiteID string, sessionID string, originalID string) (*ConversationOriginal, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/original/%s", websiteID, sessionID, originalID)
  req, _ := service.client.NewRequest("GET", url, nil)

  original := new(ConversationOriginalData)
  resp, err := service.client.Do(req, original)
  if err != nil {
    return nil, resp, err
  }

  return original.Data, resp, err
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


// ListConversationFiles lists files in conversation (extracted from messages).
func (service *WebsiteService) ListConversationFiles(websiteID string, sessionID string, pageNumber uint) (*[]ConversationFile, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/files/%d", websiteID, sessionID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  files := new(ConversationFilesData)
  resp, err := service.client.Do(req, files)
  if err != nil {
    return nil, resp, err
  }

  return files.Data, resp, err
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


// GetConversationParticipants resolves conversation participants.
func (service *WebsiteService) GetConversationParticipants(websiteID string, sessionID string) (*ConversationParticipants, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/participants", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  participants := new(ConversationParticipantsData)
  resp, err := service.client.Do(req, participants)
  if err != nil {
    return nil, resp, err
  }

  return participants.Data, resp, err
}


// SaveConversationParticipants saves conversation participants.
func (service *WebsiteService) SaveConversationParticipants(websiteID string, sessionID string, participants ConversationParticipantsSave) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/participants", websiteID, sessionID)
  req, _ := service.client.NewRequest("PUT", url, participants)

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


// GetVerifyStatusForConversation resolves conversation verify status.
func (service *WebsiteService) GetVerifyStatusForConversation(websiteID string, sessionID string) (*ConversationVerify, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/verify", websiteID, sessionID)
  req, _ := service.client.NewRequest("GET", url, nil)

  verify := new(ConversationVerifyData)
  resp, err := service.client.Do(req, verify)
  if err != nil {
    return nil, resp, err
  }

  return verify.Data, resp, err
}


// UpdateVerifyStatusForConversation updates conversation verify status.
func (service *WebsiteService) UpdateVerifyStatusForConversation(websiteID string, sessionID string, verified bool) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/verify", websiteID, sessionID)
  req, _ := service.client.NewRequest("PATCH", url, ConversationVerifyUpdate{&verified})

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


// RequestChatboxBindingPurgeForConversation requests a chatbox binding purge for conversation.
func (service *WebsiteService) RequestChatboxBindingPurgeForConversation(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/purge", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// RequestUserFeedbackForConversation requests feedback from user for conversation.
func (service *WebsiteService) RequestUserFeedbackForConversation(websiteID string, sessionID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/feedback", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, nil)

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


// AssistExistingBrowsingSession assists an existing browsing session.
func (service *WebsiteService) AssistExistingBrowsingSession(websiteID string, sessionID string, browsingID string, assist ConversationBrowsingAssist) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/browsing/%s/assist", websiteID, sessionID, browsingID)
  req, _ := service.client.NewRequest("PATCH", url, assist)

  return service.client.Do(req, nil)
}


// InitiateNewCallSessionForConversation initiates a new audio/video call session for conversation.
func (service *WebsiteService) InitiateNewCallSessionForConversation(websiteID string, sessionID string, mode string) (*ConversationCall, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/call", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, ConversationCallPayload{mode})

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


// DeliverWidgetButtonActionForConversation delivers a button action on plugin widget for conversation.
func (service *WebsiteService) DeliverWidgetButtonActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}, value *interface{}) (*ConversationWidgetAction, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/widget/%s/button", websiteID, sessionID, pluginID)
  req, _ := service.client.NewRequest("POST", url, ConversationWidgetButtonPayload{SectionID: sectionID, ItemID: itemID, Data: data, Value: value})

  action := new(ConversationWidgetActionData)
  resp, err := service.client.Do(req, action)
  if err != nil {
    return nil, resp, err
  }

  return action.Data, resp, err
}


// DeliverWidgetDataFetchActionForConversation delivers a data action on plugin widget for conversation (fetch action).
func (service *WebsiteService) DeliverWidgetDataFetchActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}) (*ConversationWidgetAction, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/widget/%s/data", websiteID, sessionID, pluginID)
  req, _ := service.client.NewRequest("POST", url, ConversationWidgetDataFetchPayload{SectionID: sectionID, ItemID: itemID, Data: data})

  action := new(ConversationWidgetActionData)
  resp, err := service.client.Do(req, action)
  if err != nil {
    return nil, resp, err
  }

  return action.Data, resp, err
}


// DeliverWidgetDataEditActionForConversation delivers a data action on plugin widget for conversation (edit action).
func (service *WebsiteService) DeliverWidgetDataEditActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, value string) (*ConversationWidgetAction, *Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/widget/%s/data", websiteID, sessionID, pluginID)
  req, _ := service.client.NewRequest("POST", url, ConversationWidgetDataEditPayload{SectionID: sectionID, ItemID: itemID, Value: value})

  action := new(ConversationWidgetActionData)
  resp, err := service.client.Do(req, action)
  if err != nil {
    return nil, resp, err
  }

  return action.Data, resp, err
}


// ScheduleReminderForConversation schedules a reminder in the future for conversation.
func (service *WebsiteService) ScheduleReminderForConversation(websiteID string, sessionID string, date string, note string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/reminder", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, ConversationReminderPayload{Date: date, Note: note})

  return service.client.Do(req, nil)
}


// ReportConversation reports a conversation to Crisp and remove it from the inbox.
func (service *WebsiteService) ReportConversation(websiteID string, sessionID string, flag string) (*Response, error) {
  url := fmt.Sprintf("website/%s/conversation/%s/report", websiteID, sessionID)
  req, _ := service.client.NewRequest("POST", url, ConversationReportPayload{Flag: flag})

  return service.client.Do(req, nil)
}
