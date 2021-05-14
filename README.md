# go-crisp-api

[![Test and Build](https://github.com/crisp-im/go-crisp-api/workflows/Test%20and%20Build/badge.svg?branch=master)](https://github.com/crisp-im/go-crisp-api/actions?query=workflow%3A%22Test+and+Build%22)

The Crisp API Golang wrapper. Authenticate, send messages, fetch conversations, access your agent accounts from your Go code.

Copyright 2021 Crisp IM SARL. See LICENSE for copying information.

* **📝 Implements**: [Crisp Platform - API ~ v1](https://docs.crisp.chat/api/v1/) at reference revision: 14/05/2021
* **😘 Maintainer**: [@valeriansaliou](https://github.com/valeriansaliou)

## Usage

```go
import "github.com/crisp-im/go-crisp-api/crisp"
```

Construct a new Crisp client, then use the various services on the client to
access different parts of the Crisp API. For example:

```go
client := crisp.New()

// List plugin subscriptions for website
subscriptions, _, err := client.Plugin.ListSubscriptionsForWebsite("5d02a3ef-ea86-47b8-a6eb-809f787abab5")
```

## Authentication

To authenticate against the API, generate your session identifier and session key **once** using the [Crisp token generation utility](https://go.crisp.chat/account/token/). You'll get a token keypair made of 2 values.

**Keep your token keypair values private, and store them safely for long-term use.**

Then, add authentication parameters to your `client` instance right after you create it:

```go
client := crisp.New()

// Authenticate to API (identifier, key)
// eg. client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")
client.Authenticate(identifier, key)

// Now, you can use authenticated API sections.
```

**🔴 Important: Make sure to generate your token once, and use the same token keys in all your subsequent requests to the API. Do not generate too many tokens, as we may invalidate your older tokens to make room for newer tokens.**

## Resource Methods

All the available Crisp API resources are fully implemented. **Programmatic methods names are named after their label name in the [API Reference](https://docs.crisp.chat/api/v1/)**.

Thus, it is straightforward to look for them in the library while reading the [API Reference](https://docs.crisp.chat/api/v1/).

**⚠️ Note that, depending on your authentication token tier, which is either `user` or `plugin`, you may not be allowed to use all methods from the library. When in doubt, refer to the library method descriptions below.**

In the following method prototypes, `crisp` is to be replaced with your Crisp API instance. For example, instanciate `client := crisp.New()` and then call eg: `client.Website.ListWebsiteOperators(websiteID)`.

When calling a method that writes data to the API (eg: `client.Website.CreateWebsite()`), you need to build an account instance and submit it this way:

```go
websiteData := crisp.WebsiteCreate{
  Name: "Acme Inc.",
  Domain: "acme-inc.com",
}

client.Website.CreateWebsite(websiteData)
```

Refer directly to [the library source code](https://github.com/crisp-im/go-crisp-api/tree/master/crisp) to know more about those structures.

### Bucket

* **Bucket URL**
  * **Generate Bucket URL** [`user`, `plugin`]: `client.Bucket.GenerateBucketURL(bucketData BucketURLRequest) (*Response, error)`

### Media

* **Media Animation**
  * **List Animation Medias** [`user`]: `client.Media.ListAnimationMedias(pageNumber uint, listID int, searchQuery string) (*Response, error)`

### Website

* **Website Base**
  * **Check If Website Exists** [`user`, `plugin`]: `client.Website.CheckWebsiteExists(domain string) (*Response, error)`
  * **Create Website** [`user`]: `client.Website.CreateWebsite(websiteData WebsiteCreate) (*Website, *Response, error)`
  * **Get A Website** [`user`, `plugin`]: `client.Website.GetWebsite(websiteID string) (*Website, *Response, error)`
  * **Delete A Website** [`user`]: `client.Website.DeleteWebsite(websiteID string, verify string) (*Response, error)`

* **Website Batch**
  * **Batch Resolve Conversations** [`user`]: `client.Website.BatchResolveConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Read Conversations** [`user`]: `client.Website.BatchReadConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Remove Conversations** [`user`]: `client.Website.BatchRemoveConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Remove People** [`user`]: `client.Website.BatchRemovePeople(websiteID string, people WebsiteBatchPeopleOperationInner) (*Response, error)`

* **Website Availability**
  * **Get Website Availability Status** [`user`, `plugin`]: `client.Website.GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error)`
  * **List Website Operator Availabilities** [`user`, `plugin`]: `client.Website.ListWebsiteOperatorAvailabilities(websiteID string) (*[]WebsiteAvailabilityOperator, *Response, error)`

* **Website Operator**
  * **List Website Operators** [`user`, `plugin`]: `client.Website.ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error)`
  * **List Last Active Website Operators** [`user`, `plugin`]: `client.Website.ListLastActiveWebsiteOperators(websiteID string) (*[]WebsiteOperatorsLastActiveListOne, *Response, error)`
  * **Flush Last Active Website Operators** [`user`]: `client.Website.FlushLastActiveWebsiteOperators(websiteID string) (*Response, error)`
  * **Send Email To Website Operators** [`user`, `plugin`]: `client.Website.SendEmailToWebsiteOperators(websiteID string, email WebsiteOperatorEmail) (*Response, error)`
  * **Get A Website Operator** [`user`, `plugin`]: `client.Website.GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error)`
  * **Invite A Website Operator** [`user`]: `client.Website.InviteWebsiteOperator(websiteID string, email string, role string, verify string) (*Response, error)`
  * **Change Operator Membership** [`user`]: `client.Website.ChangeOperatorMembership(websiteID string, userID string, role string, title *string) (*Response, error)`
  * **Unlink Operator From Website** [`user`]: `client.Website.UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error)`

* **Website Verify**
  * **Get Verify Settings** [`user`, `plugin`]: `client.Website.GetVerifySettings(websiteID string) (*WebsiteVerifySettings, *Response, error)`
  * **Update Verify Settings** [`user`, `plugin`]: `client.Website.UpdateVerifySettings(websiteID string, settings WebsiteVerifySettingsUpdate) (*Response, error)`
  * **Get Verify Key** [`user`, `plugin`]: `client.Website.GetVerifyKey(websiteID string) (*WebsiteVerifyKey, *Response, error)`
  * **Roll Verify Key** [`user`, `plugin`]: `client.Website.RollVerifyKey(websiteID string) (*Response, error)`

* **Website Settings**
  * **Get Website Settings** [`user`, `plugin`]: `client.Website.GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error)`
  * **Update Website Settings** [`user`, `plugin`]: `client.Website.UpdateWebsiteSettings(websiteID string, settings WebsiteSettingsUpdate) (*Response, error)`

* **Website Visitors**
  * **Count Visitors** [`user`, `plugin`]: `client.Website.CountVisitors(websiteID string) (*WebsiteVisitorCount, *Response, error)`
  * **List Visitors** [`user`, `plugin`]: `client.Website.ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error)`
  * **Pinpoint Visitors On A Map (Wide Variant)** [`user`, `plugin`]: `client.Website.PinpointVisitorsOnMapWide(websiteID string) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Pinpoint Visitors On A Map (Area Variant)** [`user`, `plugin`]: `client.Website.PinpointVisitorsOnMapArea(websiteID string, centerLongitude float32, centerLatitude float32, centerRadius uint) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Get Session Identifier From Token** [`user`, `plugin`]: `client.Website.GetSessionIdentifierFromToken(websiteID string, tokenID string) (*WebsiteVisitorsToken, *Response, error)`
  * **Count Blocked Visitors** [`user`]: `client.Website.CountBlockedVisitors(websiteID string) (*[]WebsiteVisitorsBlocked, *Response, error)`
  * **Count Blocked Visitors In Rule** [`user`]: `client.Website.CountBlockedVisitorsInRule(websiteID string, rule string) (*WebsiteVisitorsBlocked, *Response, error)`
  * **Clear Blocked Visitors In Rule** [`user`]: `client.Website.ClearBlockedVisitorsInRule(websiteID string, rule string) (*Response, error)`

* **Website Campaigns**
  * **List Campaigns** [`user`]: `client.Website.ListCampaigns(websiteID string, pageNumber uint) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaigns (Filter Variant)** [`user`]: `client.Website.FilterCampaigns(websiteID string, pageNumber uint, searchName string, filterTag string, filterTypeOneShot bool, filterTypeAutomated bool, filterStatusNotConfigured bool, filterStatusReady bool, filterStatusPaused bool, filterStatusSending bool, filterStatusDone bool) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaign Tags** [`user`]: `client.Website.ListCampaignTags(websiteID string) (*[]string, *Response, error)`
  * **List Campaign Templates** [`user`]: `client.Website.ListCampaignTemplates(websiteID string, pageNumber uint) (*[]WebsiteCampaignTemplateExcerpt, *Response, error)`
  * **List Campaign Templates (Search Variant)** [`user`]: `client.Website.SearchCampaignTemplates(websiteID string, pageNumber uint, searchName string) (*[]WebsiteCampaignTemplateExcerpt, *Response, error)`
  * **Create A New Campaign Template** [`user`]: `client.Website.CreateNewCampaignTemplate(websiteID string, templateFormat string, templateName string) (*WebsiteCampaignTemplateNew, *Response, error)`
  * **Check If Campaign Template Exists** [`user`]: `client.Website.CheckCampaignTemplateExists(websiteID string, templateID string) (*Response, error)`
  * **Get A Campaign Template** [`user`]: `client.Website.GetCampaignTemplate(websiteID string, templateID string) (*WebsiteCampaignTemplateItem, *Response, error)`
  * **Save A Campaign Template** [`user`]: `client.Website.SaveCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error)`
  * **Update A Campaign Template** [`user`]: `client.Website.UpdateCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error)`
  * **Remove A Campaign Template** [`user`]: `client.Website.RemoveCampaignTemplate(websiteID string, templateID string) (*Response, error)`

* **Website Campaign**
  * **Create A New Campaign** [`user`]: `client.Website.CreateNewCampaign(websiteID string, campaignType string, campaignName string) (*WebsiteCampaignNew, *Response, error)`
  * **Check If Campaign Exists** [`user`]: `client.Website.CheckCampaignExists(websiteID string, campaignID string) (*Response, error)`
  * **Get A Campaign** [`user`]: `client.Website.GetCampaign(websiteID string, campaignID string) (*WebsiteCampaignItem, *Response, error)`
  * **Save A Campaign** [`user`]: `client.Website.SaveCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Update A Campaign** [`user`]: `client.Website.UpdateCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Remove A Campaign** [`user`]: `client.Website.RemoveCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Dispatch A Campaign** [`user`]: `client.Website.DispatchCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Resume A Campaign** [`user`]: `client.Website.ResumeCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Pause A Campaign** [`user`]: `client.Website.PauseCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Test A Campaign** [`user`]: `client.Website.TestCampaign(websiteID string, campaignID string) (*Response, error)`
  * **List Campaign Recipients** [`user`]: `client.Website.ListCampaignRecipients(websiteID string, campaignID string, pageNumber uint) (*[]WebsiteCampaignRecipient, *Response, error)`
  * **List Campaign Statistics** [`user`]: `client.Website.ListCampaignStatistics(websiteID string, campaignID string, action string, pageNumber uint) (*[]WebsiteCampaignStatistic, *Response, error)`

* **Website Conversations**
  * **List Conversations** [`user`, `plugin`]: `client.Website.ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error)`
  * **List Conversations (Search Variant)** [`user`, `plugin`]: `client.Website.SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error)`
  * **List Suggested Conversation Segments** [`user`, `plugin`]: `client.Website.ListSuggestedConversationSegments(websiteID string, pageNumber uint) (*[]ConversationSuggestedSegment, *Response, error)`
  * **List Suggested Conversation Data Keys** [`user`, `plugin`]: `client.Website.ListSuggestedConversationDataKeys(websiteID string, pageNumber uint) (*[]ConversationSuggestedData, *Response, error)`

* **Website Conversation**
  * **Create A New Conversation** [`user`, `plugin`]: `client.Website.CreateNewConversation(websiteID string) (*ConversationNew, *Response, error)`
  * **Check If Conversation Exists** [`user`, `plugin`]: `client.Website.CheckConversationExists(websiteID string, sessionID string) (*Response, error)`
  * **Get A Conversation** [`user`, `plugin`]: `client.Website.GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error)`
  * **Remove A Conversation** [`user`, `plugin`]: `client.Website.RemoveConversation(websiteID string, sessionID string) (*Response, error)`
  * **Initiate A Conversation With Existing Session** [`user`, `plugin`]: `client.Website.InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error)`
  * **Get Messages In Conversation (Last Variant)** [`user`, `plugin`]: `client.Website.GetMessagesInConversationLast(websiteID string, sessionID string) (*[]ConversationMessage, *Response, error)`
  * **Get Messages In Conversation (Before Variant)** [`user`, `plugin`]: `client.Website.GetMessagesInConversationBefore(websiteID string, sessionID string, timestampBefore uint32) (*[]ConversationMessage, *Response, error)`
  * **Send A Message In Conversation (Text Variant)** [`user`, `plugin`]: `client.Website.SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (File Variant)** [`user`, `plugin`]: `client.Website.SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Animation Variant)** [`user`, `plugin`]: `client.Website.SendAnimationMessageInConversation(websiteID string, sessionID string, message ConversationAnimationMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Audio Variant)** [`user`, `plugin`]: `client.Website.SendAudioMessageInConversation(websiteID string, sessionID string, message ConversationAudioMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Picker Variant)** [`user`, `plugin`]: `client.Website.SendPickerMessageInConversation(websiteID string, sessionID string, message ConversationPickerMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Field Variant)** [`user`, `plugin`]: `client.Website.SendFieldMessageInConversation(websiteID string, sessionID string, message ConversationFieldMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Note Variant)** [`user`, `plugin`]: `client.Website.SendNoteMessageInConversation(websiteID string, sessionID string, message ConversationNoteMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Event Variant)** [`user`, `plugin`]: `client.Website.SendEventMessageInConversation(websiteID string, sessionID string, message ConversationEventMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Get A Message In Conversation** [`user`, `plugin`]: `client.Website.GetMessageInConversation(websiteID string, sessionID string, fingerprint int) (*ConversationMessage, *Response, error)`
  * **Update A Message In Conversation (Text Variant)** [`user`, `plugin`]: `client.Website.UpdateTextMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (File Variant)** [`user`, `plugin`]: `client.Website.UpdateFileMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFileMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Animation Variant)** [`user`, `plugin`]: `client.Website.UpdateAnimationMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAnimationMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Audio Variant)** [`user`, `plugin`]: `client.Website.UpdateAudioMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAudioMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Picker Variant)** [`user`, `plugin`]: `client.Website.UpdatePickerMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationPickerMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Field Variant)** [`user`, `plugin`]: `client.Website.UpdateFieldMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFieldMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Note Variant)** [`user`, `plugin`]: `client.Website.UpdateNoteMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (Event Variant)** [`user`, `plugin`]: `client.Website.UpdateEventMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationEventMessageNewContent) (*Response, error)`
  * **Compose A Message In Conversation** [`user`, `plugin`]: `client.Website.ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error)`
  * **Mark Messages As Read In Conversation** [`user`, `plugin`]: `client.Website.MarkMessagesReadInConversation(websiteID string, sessionID string, read ConversationReadMessageMark) (*Response, error)`
  * **Mark Messages As Delivered In Conversation** [`user`, `plugin`]: `client.Website.MarkMessagesDeliveredInConversation(websiteID string, sessionID string, delivered ConversationDeliveredMessageMark) (*Response, error)`
  * **Update Conversation Open State** [`user`, `plugin`]: `client.Website.UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error)`
  * **Get Conversation Routing Assign** [`user`, `plugin`]: `client.Website.GetConversationRoutingAssign(websiteID string, sessionID string) (*ConversationRoutingAssign, *Response, error)`
  * **Assign Conversation Routing** [`user`, `plugin`]: `client.Website.AssignConversationRouting(websiteID string, sessionID string, assign ConversationRoutingAssignUpdate) (*Response, error)`
  * **Get Conversation Metas** [`user`, `plugin`]: `client.Website.GetConversationMetas(websiteID string, sessionID string) (*ConversationMeta, *Response, error)`
  * **Update Conversation Metas** [`user`, `plugin`]: `client.Website.UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error)`
  * **Get An Original Message In Conversation** [`user`, `plugin`]: `client.Website.GetOriginalMessageInConversation(websiteID string, sessionID string, originalID string) (*ConversationOriginal, *Response, error)`
  * **List Conversation Pages** [`user`, `plugin`]: `client.Website.ListConversationPages(websiteID string, sessionID string, pageNumber uint) (*[]ConversationPage, *Response, error)`
  * **List Conversation Events** [`user`, `plugin`]: `client.Website.ListConversationEvents(websiteID string, sessionID string, pageNumber uint) (*[]ConversationEvent, *Response, error)`
  * **Get Conversation State** [`user`, `plugin`]: `client.Website.GetConversationState(websiteID string, sessionID string) (*ConversationState, *Response, error)`
  * **Change Conversation State** [`user`, `plugin`]: `client.Website.ChangeConversationState(websiteID string, sessionID string, state string) (*Response, error)`
  * **Get Conversation Participants** [`user`, `plugin`]: `client.Website.GetConversationParticipants(websiteID string, sessionID string) (*ConversationParticipants, *Response, error)`
  * **Save Conversation Participants** [`user`, `plugin`]: `client.Website.SaveConversationParticipants(websiteID string, sessionID string, participants ConversationParticipantsSave) (*Response, error)`
  * **Get Block Status For Conversation** [`user`, `plugin`]: `client.Website.GetBlockStatusForConversation(websiteID string, sessionID string) (*ConversationBlock, *Response, error)`
  * **Block Incoming Messages For Conversation** [`user`, `plugin`]: `client.Website.BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool) (*Response, error)`
  * **Request Email Transcript For Conversation** [`user`, `plugin`]: `client.Website.RequestEmailTranscriptForConversation(websiteID string, sessionID string, to string, email string) (*Response, error)`
  * **Request Chatbox Binding Purge For Conversation** [`user`, `plugin`]: `client.Website.RequestChatboxBindingPurgeForConversation(websiteID string, sessionID string) (*Response, error)`
  * **List Browsing Sessions For Conversation** [`user`, `plugin`]: `client.Website.ListBrowsingSessionsForConversation(websiteID string, sessionID string) (*[]ConversationBrowsing, *Response, error)`
  * **Initiate Browsing Session For Conversation** [`user`, `plugin`]: `client.Website.InitiateBrowsingSessionForConversation(websiteID string, sessionID string) (*Response, error)`
  * **Send Action To An Existing Browsing Session** [`user`, `plugin`]: `client.Website.SendActionToExistingBrowsingSession(websiteID string, sessionID string, browsingID string, action string) (*Response, error)`
  * **Debug Existing Browsing Session** [`user`, `plugin`]: `client.Website.DebugExistingBrowsingSession(websiteID string, sessionID string, browsingID string, debug ConversationBrowsingDebug) (*Response, error)`
  * **Assist Existing Browsing Session** [`user`, `plugin`]: `client.Website.AssistExistingBrowsingSession(websiteID string, sessionID string, browsingID string, assist ConversationBrowsingAssist) (*Response, error)`
  * **Initiate New Call Session For Conversation** [`user`, `plugin`]: `client.Website.InitiateNewCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Get Ongoing Call Session For Conversation** [`user`, `plugin`]: `client.Website.GetOngoingCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Abort Ongoing Call Session For Conversation** [`user`, `plugin`]: `client.Website.AbortOngoingCallSessionForConversation(websiteID string, sessionID string, callID string) (*Response, error)`
  * **Transmit Signaling On Ongoing Call Session** [`user`, `plugin`]: `client.Website.TransmitSignalingOnOngoingCallSession(websiteID string, sessionID string, callID string, payload ConversationCallSignalingPayload) (*Response, error)`
  * **Deliver Widget Button Action For Conversation** [`user`]: `client.Website.DeliverWidgetButtonActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}) (*ConversationWidgetAction, *Response, error)`
  * **Deliver Widget Data Fetch Action For Conversation** [`user`]: `client.Website.DeliverWidgetDataFetchActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string) (*ConversationWidgetAction, *Response, error)`
  * **Deliver Widget Data Edit Action For Conversation** [`user`]: `client.Website.DeliverWidgetDataEditActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, value string) (*ConversationWidgetAction, *Response, error)`
  * **Schedule A Reminder For Conversation** [`user`, `plugin`]: `client.Website.ScheduleReminderForConversation(websiteID string, sessionID string, date string, note string) (*Response, error)`

* **Website Analytics**
  * **Acquire Analytics Points** [`user`]: `client.Website.AcquireAnalyticsPoints(websiteID string, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time, dateSplit string, classifier string, filterPrimary string, filterSecondary string, filterTertiary string) (*WebsiteAnalyticsPoints, *Response, error)`
  * **List Analytics Filters** [`user`]: `client.Website.ListAnalyticsFilters(websiteID string, pageNumber uint, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time) (*[]WebsiteAnalyticsFilter, *Response, error)`
  * **List Analytics Classifiers** [`user`]: `client.Website.ListAnalyticsClassifiers(websiteID string, pageNumber uint, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time) (*[]WebsiteAnalyticsClassifier, *Response, error)`

* **Website People**
  * **Get People Statistics** [`user`, `plugin`]: `client.Website.GetPeopleStatistics(websiteID string) (*PeopleStatistics, *Response, error)`
  * **List Suggested People Segments** [`user`, `plugin`]: `client.Website.ListSuggestedPeopleSegments(websiteID string, pageNumber uint) (*[]PeopleSuggestedSegment, *Response, error)`
  * **List Suggested People Data Keys** [`user`, `plugin`]: `client.Website.ListSuggestedPeopleDataKeys(websiteID string, pageNumber uint) (*[]PeopleSuggestedData, *Response, error)`
  * **List People Profiles** [`user`, `plugin`]: `client.Website.ListPeopleProfiles(websiteID string, pageNumber uint, searchField string, searchOrder string, searchOperator string, searchFilter []PeopleFilter, searchText string) (*[]PeopleProfile, *Response, error)`
  * **Add New People Profile** [`user`, `plugin`]: `client.Website.AddNewPeopleProfile(websiteID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **Check If People Profile Exists** [`user`, `plugin`]: `client.Website.CheckPeopleProfileExists(websiteID string, peopleID string) (*Response, error)`
  * **Get People Profile** [`user`, `plugin`]: `client.Website.GetPeopleProfile(websiteID string, peopleID string) (*PeopleProfile, *Response, error)`
  * **Save People Profile** [`user`, `plugin`]: `client.Website.SavePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **Update People Profile** [`user`, `plugin`]: `client.Website.UpdatePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **Remove People Profile** [`user`, `plugin`]: `client.Website.RemovePeopleProfile(websiteID string, peopleID string) (*Response, error)`
  * **List People Conversations** [`user`, `plugin`]: `client.Website.ListPeopleConversations(websiteID string, peopleID string, pageNumber uint) ([]string, *Response, error)`
  * **List People Campaigns** [`user`]: `client.Website.ListPeopleCampaigns(websiteID string, peopleID string, pageNumber uint) (*[]PeopleCampaign, *Response, error)`
  + **Add A People Event** [`user`, `plugin`]: `client.Website.AddPeopleEvent(websiteID string, peopleID string, peopleEvent PeopleEventAdd) (*Response, error)`
  + **List People Events** [`user`, `plugin`]: `client.Website.ListPeopleEvents(websiteID string, peopleID string, pageNumber uint) (*[]PeopleEvent, *Response, error)`
  + **Get People Data** [`user`, `plugin`]: `client.Website.GetPeopleData(websiteID string, peopleID string) (*PeopleData, *Response, error)`
  + **Save People Data** [`user`, `plugin`]: `client.Website.SavePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error)`
  + **Get People Subscription Status** [`user`, `plugin`]: `client.Website.GetPeopleSubscriptionStatus(websiteID string, peopleID string) (*PeopleSubscription, *Response, error)`
  + **Update People Subscription Status** [`user`, `plugin`]: `client.Website.UpdatePeopleSubscriptionStatus(websiteID string, peopleID string, peopleSubscription PeopleSubscriptionUpdate)`
  * **Export People Profiles** [`user`]: `client.Website.ExportPeopleProfiles(websiteID string) (*Response, error)`
  * **Import People Profiles** [`user`]: `client.Website.ImportPeopleProfiles(websiteID string, profileImportSetup PeopleProfileImportSetup) (*PeopleProfileImport, *Response, error)`

### Plugin

* **Plugin Subscription**
  * **List All Active Subscriptions** [`user`]: `client.Plugin.ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error)`
  * **List Subscriptions For A Website** [`user`]: `client.Plugin.ListSubscriptionsForWebsite(websiteID string) (*[]PluginSubscription, *Response, error)`
  * **Get Subscription Details** [`user`]: `client.Plugin.GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error)`
  * **Subscribe Website To Plugin** [`user`]: `client.Plugin.SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error)`
  * **Unsubscribe Plugin From Website** [`user`]: `client.Plugin.UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error)`
  * **Get Subscription Settings** [`user`, `plugin`]: `client.Plugin.GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error)`
  * **Save Subscription Settings** [`user`, `plugin`]: `client.Plugin.SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Update Subscription Settings** [`user`, `plugin`]: `client.Plugin.UpdateSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Forward Plugin Payload To Channel** [`user`, `plugin`]: `client.Plugin.ForwardPluginPayloadToChannel(websiteID string, pluginID string, payload PluginSubscriptionChannelForward) (*Response, error)`
  * **Dispatch Plugin Event** [`user`, `plugin`]: `client.Plugin.DispatchPluginEvent(websiteID string, pluginID string, payload PluginSubscriptionEventDispatch) (*Response, error)`

* **Plugin Connect**
  * **Get Connect Account** [`user`, `plugin`]: `client.Plugin.GetConnectAccount() (*PluginConnectAccount, *Response, error)`
  * **Check Connect Session Validity** [`user`, `plugin`]: `client.Plugin.CheckConnectSessionValidity() (*Response, error)`
  * **List All Connect Websites** [`user`, `plugin`]: `client.Plugin.ListAllConnectWebsites(pageNumber uint, filterConfigured bool) (*[]PluginConnectAllWebsites, *Response, error)`
  * **List Connect Websites Since** [`user`, `plugin`]: `client.Plugin.ListConnectWebsitesSince(dateSince time.Time, filterConfigured bool) (*[]PluginConnectWebsitesSince, *Response, error)`

## Realtime Events

You can bind to realtime events from Crisp, in order to get notified of incoming messages and updates in websites.

You won't receive any event if you don't explicitly subscribe to realtime events, as the library doesn't connect to the realtime backend automatically.

To start listening for events and bind an handler, proceed as follow:

```go
client := crisp.New()

// Set authentication parameters
client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")

// Connect to realtime events backend and listen (only to 'message:send' namespace)
client.Events.Listen(
  []string{
    "message:send",
  },

  func(reg *crisp.EventsRegister) {
    // Socket is connected: now listening for events

    // Notice: if the realtime socket breaks at any point, this function will be called again upon reconnect (to re-bind events)
    // Thus, ensure you only use this to register handlers

    // Register handler on 'message:send/text' namespace
    reg.On("message:send/text", func(evt crisp.EventsReceiveTextMessage) {
      // Handle text message from visitor
    })

    // Register handler on 'message:send/file' namespace
    reg.On("message:send/file", func(evt crisp.EventsReceiveFileMessage) {
      // Handle file message from visitor
    })

    // Register handler on 'message:send/animation' namespace
    reg.On("message:send/animation", func(evt crisp.EventsReceiveAnimationMessage) {
      // Handle animation message from visitor
    })

    // Register handler on 'message:send/audio' namespace
    reg.On("message:send/audio", func(evt crisp.EventsReceiveAudioMessage) {
      // Handle audio message from visitor
    })

    // Register handler on 'message:send/picker' namespace
    reg.On("message:send/picker", func(evt crisp.EventsReceivePickerMessage) {
      // Handle picker message from visitor
    })

    // Register handler on 'message:send/field' namespace
    reg.On("message:send/field", func(evt crisp.EventsReceiveFieldMessage) {
      // Handle field message from visitor
    })

    // Register handler on 'message:send/event' namespace
    reg.On("message:send/event", func(evt crisp.EventsReceiveEventMessage) {
      // Handle event message from visitor
    })
  },

  func() {
    // Socket is disconnected: will try to reconnect
  },

  func() {
    // Socket error: may be broken
  },
)
```

Available events are listed below:

* **Session Events**
  * **Session Update Availability** [`user`, `plugin`]: `session:update_availability`
  * **Session Update Verify** [`user`, `plugin`]: `session:update_verify`
  * **Session Request Initiated** [`user`, `plugin`]: `session:request:initiated`
  * **Session Set Email** [`user`, `plugin`]: `session:set_email`
  * **Session Set Phone** [`user`, `plugin`]: `session:set_phone`
  * **Session Set Address** [`user`, `plugin`]: `session:set_address`
  * **Session Set Avatar** [`user`, `plugin`]: `session:set_avatar`
  * **Session Set Nickname** [`user`, `plugin`]: `session:set_nickname`
  * **Session Set Data** [`user`, `plugin`]: `session:set_data`
  * **Session Sync Pages** [`user`, `plugin`]: `session:sync:pages`
  * **Session Sync Events** [`user`, `plugin`]: `session:sync:events`
  * **Session Sync Capabilities** [`user`, `plugin`]: `session:sync:capabilities`
  * **Session Sync Geolocation** [`user`, `plugin`]: `session:sync:geolocation`
  * **Session Sync System** [`user`, `plugin`]: `session:sync:system`
  * **Session Sync Network** [`user`, `plugin`]: `session:sync:network`
  * **Session Sync Timezone** [`user`, `plugin`]: `session:sync:timezone`
  * **Session Sync Locales** [`user`, `plugin`]: `session:sync:locales`
  * **Session Sync Rating** [`user`, `plugin`]: `session:sync:rating`
  * **Session Set State** [`user`, `plugin`]: `session:set_state`
  * **Session Set Block** [`user`, `plugin`]: `session:set_block`
  * **Session Set Segments** [`user`, `plugin`]: `session:set_segments`
  * **Session Set Opened** [`user`, `plugin`]: `session:set_opened`
  * **Session Set Closed** [`user`, `plugin`]: `session:set_closed`
  * **Session Set Participants** [`user`, `plugin`]: `session:set_participants`
  * **Session Set Mentions** [`user`, `plugin`]: `session:set_mentions`
  * **Session Set Routing** [`user`, `plugin`]: `session:set_routing`
  * **Session Removed** [`user`, `plugin`]: `session:removed`

* **Message Events**
  * **Message Updated** [`user`, `plugin`]: `message:updated`
  * **Message Send (Text Variant)** [`user`, `plugin`]: `message:send/text`
  * **Message Send (File Variant)** [`user`, `plugin`]: `message:send/file`
  * **Message Send (Animation Variant)** [`user`, `plugin`]: `message:send/animation`
  * **Message Send (Audio Variant)** [`user`, `plugin`]: `message:send/audio`
  * **Message Send (Picker Variant)** [`user`, `plugin`]: `message:send/picker`
  * **Message Send (Field Variant)** [`user`, `plugin`]: `message:send/field`
  * **Message Send (Note Variant)** [`user`, `plugin`]: `message:send/note`
  * **Message Send (Event Variant)** [`user`, `plugin`]: `message:send/event`
  * **Message Received (Text Variant)** [`user`, `plugin`]: `message:received/text`
  * **Message Received (File Variant)** [`user`, `plugin`]: `message:received/file`
  * **Message Received (Animation Variant)** [`user`, `plugin`]: `message:received/animation`
  * **Message Received (Audio Variant)** [`user`, `plugin`]: `message:received/audio`
  * **Message Received (Picker Variant)** [`user`, `plugin`]: `message:received/picker`
  * **Message Received (Field Variant)** [`user`, `plugin`]: `message:received/field`
  * **Message Received (Note Variant)** [`user`, `plugin`]: `message:received/note`
  * **Message Received (Event Variant)** [`user`, `plugin`]: `message:received/event`
  * **Message Compose Send** [`user`, `plugin`]: `message:compose:send`
  * **Message Compose Receive** [`user`, `plugin`]: `message:compose:receive`
  * **Message Acknowledge Read Send** [`user`, `plugin`]: `message:acknowledge:read:send`
  * **Message Acknowledge Read Received** [`user`, `plugin`]: `message:acknowledge:read:received`
  * **Message Acknowledge Delivered** [`user`, `plugin`]: `message:acknowledge:delivered`
  * **Message Notify Unread Send** [`user`, `plugin`]: `message:notify:unread:send`
  * **Message Notify Unread Received** [`user`, `plugin`]: `message:notify:unread:received`

* **People Events**
  * **People Profile Created** [`user`, `plugin`]: `people:profile:created`
  * **People Profile Updated** [`user`, `plugin`]: `people:profile:updated`
  * **People Profile Removed** [`user`, `plugin`]: `people:profile:removed`
  * **People Bind Session** [`user`, `plugin`]: `people:bind:session`
  * **People Sync Profile** [`user`, `plugin`]: `people:sync:profile`
  * **People Import Progress** [`user`]: `people:import:progress`
  * **People Import Done** [`user`]: `people:import:done`

* **Campaign Events**
  * **Campaign Progress** [`user`]: `campaign:progress`
  * **Campaign Dispatched** [`user`]: `campaign:dispatched`
  * **Campaign Running** [`user`]: `campaign:running`

* **Browsing Events**
  * **Browsing Request Initiated** [`user`, `plugin`]: `browsing:request:initiated`
  * **Browsing Request Rejected** [`user`, `plugin`]: `browsing:request:rejected`

* **Call Events**
  * **Call Request Initiated** [`user`, `plugin`]: `call:request:initiated`
  * **Call Request Rejected** [`user`, `plugin`]: `call:request:rejected`

* **Widget Events**
  * **Widget Action Processed** [`user`]: `widget:action:processed`

* **Status Events**
  * **Status Health Changed** [`user`]: `status:health:changed`

* **Website Events**
  * **Website Update Visitors Count** [`user`, `plugin`]: `website:update_visitors_count`
  * **Website Update Operators Availability** [`user`, `plugin`]: `website:update_operators_availability`
  * **Website Users Available** [`user`, `plugin`]: `website:users:available`

* **Bucket Events**
  * **Bucket URL Upload Generated** [`user`, `plugin`]: `bucket:url:upload:generated`
  * **Bucket URL Avatar Generated** [`user`, `plugin`]: `bucket:url:avatar:generated`
  * **Bucket URL Website Generated** [`user`, `plugin`]: `bucket:url:website:generated`
  * **Bucket URL Campaign Generated** [`user`, `plugin`]: `bucket:url:campaign:generated`
  * **Bucket URL Helpdesk Generated** [`user`, `plugin`]: `bucket:url:helpdesk:generated`
  * **Bucket URL Status Generated** [`user`, `plugin`]: `bucket:url:status:generated`
  * **Bucket URL Processing Generated** [`user`, `plugin`]: `bucket:url:processing:generated`

* **Media Events**
  * **Media Animation Listed** [`user`]: `media:animation:listed`

* **Email Events**
  * **Email Subscribe** [`user`, `plugin`]: `email:subscribe`
  * **Email Track View** [`user`, `plugin`]: `email:track:view`

* **Plugin Events**
  * **Plugin Channel** [`user`, `plugin`]: `plugin:channel`
  * **Plugin Event** [`user`, `plugin`]: `plugin:event`
  * **Plugin Settings Saved** [`user`, `plugin`]: `plugin:settings:saved`
