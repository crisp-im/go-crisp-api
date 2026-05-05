# go-crisp-api

[![Test and Build](https://github.com/crisp-im/go-crisp-api/actions/workflows/test.yml/badge.svg)](https://github.com/crisp-im/go-crisp-api/actions/workflows/test.yml)

The Crisp API Golang wrapper. Authenticate, send messages, fetch conversations, access your agent accounts from your Go code.

Copyright 2026 Crisp IM SAS. See LICENSE for copying information.

* **📝 Implements**: [REST API Reference (V1)](https://docs.crisp.chat/references/rest-api/v1/) at revision: 05/03/2026
* **😘 Maintainer**: [@valeriansaliou](https://github.com/valeriansaliou)

## Usage

You may follow the [REST API Quickstart](https://docs.crisp.chat/guides/rest-api/quickstart/) guide, which will get you running with the REST API in minutes.

```go
import "github.com/crisp-im/go-crisp-api/crisp/v3"
```

Construct a new Crisp client, then use the various services on the client to
access different parts of the Crisp API. For example:

```go
client := crisp.New()

// List plugin subscriptions for website
subscriptions, _, err := client.Plugin.ListSubscriptionsForWebsite("5d02a3ef-ea86-47b8-a6eb-809f787abab5")
```

## Authentication

To authenticate against the API, obtain your authentication token keypair by following the [REST API Authentication](https://docs.crisp.chat/guides/rest-api/authentication/) guide. You'll get a token keypair made of 2 values.

**Keep your token keypair values private, and store them safely for long-term use.**

Then, add authentication parameters to your `client` instance right after you create it:

```go
client := crisp.New()

// Authenticate to API with your token (identifier, key)
// eg. client.AuthenticateTier("plugin", "43f34724-9eeb-4474-9cec-560250754dec", "d12e60c5d2aa264b90997a641b6474ffd6602b66d8e8abc49634c404f06fa7d0")
client.AuthenticateTier(tier, identifier, key)

// Now, you can use authenticated API sections.
```

## Resource Methods

All the available Crisp API resources are fully implemented. **Programmatic methods names are named after their label name in the [REST API Reference](https://docs.crisp.chat/references/rest-api/v1/)**.

All methods that you will most likely need when building a Crisp integration are prefixed with a star symbol (⭐).

**⚠️ Note that, depending on your authentication token tier, which is either `user`, `website` or `plugin`, you may not be allowed to use all methods from the library. When in doubt, refer to the library method descriptions below. Most likely, you are using a `plugin` token.**

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

* #### **Bucket URL**
  * **Generate Bucket URL**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#bucket-url)
    * `client.Bucket.GenerateBucketURL(bucketData BucketURLRequest) (*Response, error)`

### Media

* #### **Media Animation**
  * **List Animation Medias**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-animation-medias)
    * `client.Media.ListAnimationMedias(pageNumber uint, listID int, searchQuery string) (*Response, error)`

### Website

* #### **Website Base**
  * **Check If Website Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-website-exists)
    * `client.Website.CheckWebsiteExists(domain string) (*Response, error)`
  * **Create Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-website)
    * `client.Website.CreateWebsite(websiteData WebsiteCreate) (*Website, *Response, error)`
  * **Get A Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-website)
    * `client.Website.GetWebsite(websiteID string) (*Website, *Response, error)`
  * **Delete A Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-a-website)
    * `client.Website.DeleteWebsite(websiteID string, verify WebsiteRemoveVerify) (*Response, error)`
  * **Abort Website Deletion**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#abort-website-deletion)
    * `client.Website.AbortWebsiteDeletion(websiteID string) (*Response, error)`

* #### **Website Batch**
  * **Batch Resolve Conversations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-resolve-items)
    * `client.Website.BatchResolveConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error)`
  * **Batch Read Conversations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-read-items)
    * `client.Website.BatchReadConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error)`
  * **Batch Remove Conversations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-remove-items)
    * `client.Website.BatchRemoveConversations(websiteID string, operation WebsiteBatchConversationsOperation) (*Response, error)`
  * **Batch Remove People**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-remove-items)
    * `client.Website.BatchRemovePeople(websiteID string, people WebsiteBatchPeopleOperationInner) (*Response, error)`

* #### **Website Availability**
  * **Get Website Availability Status**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-website-availability-status)
    * `client.Website.GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error)`
  * **List Website Operator Availabilities**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-website-operator-availabilities)
    * `client.Website.ListWebsiteOperatorAvailabilities(websiteID string) (*[]WebsiteAvailabilityOperator, *Response, error)`

* #### **Website Operator**
  * **List Website Operators**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-website-operators)
    * `client.Website.ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error)`
  * **List Last Active Website Operators**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-last-active-website-operators)
    * `client.Website.ListLastActiveWebsiteOperators(websiteID string) (*[]WebsiteOperatorsLastActiveListOne, *Response, error)`
  * **Flush Last Active Website Operators**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#flush-last-active-website-operators)
    * `client.Website.FlushLastActiveWebsiteOperators(websiteID string) (*Response, error)`
  * **Send Email To Website Operators**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-email-to-website-operators)
    * `client.Website.SendEmailToWebsiteOperators(websiteID string, email WebsiteOperatorEmail) (*Response, error)`
  * **Get A Website Operator**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-website-operator)
    * `client.Website.GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error)`
  * **Invite A Website Operator**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#invite-a-website-operator)
    * `client.Website.InviteWebsiteOperator(websiteID string, email string, role string, verify WebsiteOperatorInviteVerify) (*Response, error)`
  * **Change Operator Membership**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#change-operator-membership)
    * `client.Website.ChangeOperatorMembership(websiteID string, userID string, role string, title *string) (*Response, error)`
  * **Unlink Operator From Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unlink-operator-from-website)
    * `client.Website.UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error)`

* #### **Website Verify**
  * **Get Verify Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-verify-settings)
    * `client.Website.GetVerifySettings(websiteID string) (*WebsiteVerifySettings, *Response, error)`
  * **Update Verify Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-verify-settings)
    * `client.Website.UpdateVerifySettings(websiteID string, settings WebsiteVerifySettingsUpdate) (*Response, error)`
  * **Get Verify Key**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-verify-key)
    * `client.Website.GetVerifyKey(websiteID string) (*WebsiteVerifyKey, *Response, error)`
  * **Roll Verify Key**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#roll-verify-key)
    * `client.Website.RollVerifyKey(websiteID string) (*Response, error)`

* #### **Website Inbox**
  * **List Inboxes**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-website-inboxes)
    * `client.Website.ListInboxes(websiteID string, pageNumber uint) (*[]WebsiteInbox, *Response, error)`
  * **Batch Order Inboxes**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-order-website-inboxes)
    * `client.Website.BatchOrderInboxes(websiteID string, orders WebsiteInboxOrdersItem) (*Response, error)`
  * **Create A New Inbox**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-website-inbox)
    * `client.Website.CreateNewInbox(websiteID string, inbox WebsiteInboxItem) (*WebsiteInboxNew, *Response, error)`
  * **Check If Inbox Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-website-inbox-exists)
    * `client.Website.CheckInboxExists(websiteID string, inboxID string) (*Response, error)`
  * **Get Inbox**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-website-inbox)
    * `client.Website.GetInbox(websiteID string, inboxID string) (*WebsiteInbox, *Response, error)`
  * **Save Inbox**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-website-inbox)
    * `client.Website.SaveInbox(websiteID string, inboxID string, inbox WebsiteInboxItem) (*Response, error)`
  * **Delete Inbox**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-website-inbox)
    * `client.Website.DeleteInbox(websiteID string, inboxID string) (*Response, error)`

* #### **Website Settings**
  * **Get Website Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-website-settings)
    * `client.Website.GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error)`
  * **Update Website Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-website-settings)
    * `client.Website.UpdateWebsiteSettings(websiteID string, settings WebsiteSettingsUpdate) (*Response, error)`

* #### **Website Visitors**
  * **Count Visitors**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#count-visitors)
    * `client.Website.CountVisitors(websiteID string) (*WebsiteVisitorCount, *Response, error)`
  * **List Visitors**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-visitors)
    * `client.Website.ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error)`
  * **Pinpoint Visitors On A Map (Wide Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#pinpoint-visitors-on-a-map)
    * `client.Website.PinpointVisitorsOnMapWide(websiteID string) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Pinpoint Visitors On A Map (Area Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#pinpoint-visitors-on-a-map)
    * `client.Website.PinpointVisitorsOnMapArea(websiteID string, centerLongitude float32, centerLatitude float32, centerRadius uint) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Get Session Identifier From Token**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-session-identifier-from-token)
    * `client.Website.GetSessionIdentifierFromToken(websiteID string, tokenID string) (*WebsiteVisitorsToken, *Response, error)`
  * **Count Blocked Visitors**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#count-blocked-visitors/)
    * `client.Website.CountBlockedVisitors(websiteID string) (*[]WebsiteVisitorsBlocked, *Response, error)`
  * **Count Blocked Visitors In Rule**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#count-blocked-visitors-in-rule)
    * `client.Website.CountBlockedVisitorsInRule(websiteID string, rule string) (*WebsiteVisitorsBlocked, *Response, error)`
  * **Clear Blocked Visitors In Rule**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#clear-blocked-visitors-in-rule)
    * `client.Website.ClearBlockedVisitorsInRule(websiteID string, rule string) (*Response, error)`

* #### **Website Campaigns**
  * **List Campaigns**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaigns)
    * `client.Website.ListCampaigns(websiteID string, pageNumber uint) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaigns (Filter Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaigns)
    * `client.Website.FilterCampaigns(websiteID string, pageNumber uint, searchName string, filterTag string, filterTypeOneShot bool, filterTypeAutomated bool, filterStatusNotConfigured bool, filterStatusReady bool, filterStatusPaused bool, filterStatusSending bool, filterStatusDone bool) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaign Tags**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-tags)
    * `client.Website.ListCampaignTags(websiteID string) (*[]string, *Response, error)`
  * **List Campaign Templates**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-templates)
    * `client.Website.ListCampaignTemplates(websiteID string, pageNumber uint) (*[]WebsiteCampaignTemplateExcerpt, *Response, error)`
  * **List Campaign Templates (Filter Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-templates)
    * `client.Website.FilterCampaignTemplates(websiteID string, pageNumber uint, searchName string, filterTypeStatic bool, filterTypeCustom bool) (*[]WebsiteCampaignTemplateExcerpt, *Response, error)`
  * **Create A New Campaign Template**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-campaign-template)
    * `client.Website.CreateNewCampaignTemplate(websiteID string, templateFormat string, templateName string) (*WebsiteCampaignTemplateNew, *Response, error)`
  * **Check If Campaign Template Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-campaign-template-exists)
    * `client.Website.CheckCampaignTemplateExists(websiteID string, templateID string) (*Response, error)`
  * **Get A Campaign Template**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-campaign-template)
    * `client.Website.GetCampaignTemplate(websiteID string, templateID string) (*WebsiteCampaignTemplateItem, *Response, error)`
  * **Save A Campaign Template**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-a-campaign-template)
    * `client.Website.SaveCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error)`
  * **Update A Campaign Template**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-campaign-template)
    * `client.Website.UpdateCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error)`
  * **Remove A Campaign Template**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-campaign-template)
    * `client.Website.RemoveCampaignTemplate(websiteID string, templateID string) (*Response, error)`

* #### **Website Campaign**
  * **Create A New Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-campaign)
    * `client.Website.CreateNewCampaign(websiteID string, campaignType string, campaignName string) (*WebsiteCampaignNew, *Response, error)`
  * **Check If Campaign Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-campaign-exists)
    * `client.Website.CheckCampaignExists(websiteID string, campaignID string) (*Response, error)`
  * **Get A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-campaign)
    * `client.Website.GetCampaign(websiteID string, campaignID string) (*WebsiteCampaignItem, *Response, error)`
  * **Save A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-a-campaign)
    * `client.Website.SaveCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Update A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-campaign)
    * `client.Website.UpdateCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Remove A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-campaign)
    * `client.Website.RemoveCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Dispatch A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#dispatch-a-campaign)
    * `client.Website.DispatchCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Resume A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resume-a-campaign)
    * `client.Website.ResumeCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Pause A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#pause-a-campaign)
    * `client.Website.PauseCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Test A Campaign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#test-a-campaign)
    * `client.Website.TestCampaign(websiteID string, campaignID string) (*Response, error)`
  * **List Campaign Recipients**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-recipients)
    * `client.Website.ListCampaignRecipients(websiteID string, campaignID string, pageNumber uint) (*[]WebsiteCampaignRecipient, *Response, error)`
  * **List Campaign Statistics**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-statistics)
    * `client.Website.ListCampaignStatistics(websiteID string, campaignID string, action string, pageNumber uint) (*[]WebsiteCampaignStatistic, *Response, error)`

* #### **Website Conversations**
  * **⭐ List Conversations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversations)
    * `client.Website.ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error)`
  * **List Conversations (Search Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversations)
    * `client.Website.SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error)`
  * **List Conversations (Filter Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversations)
    * `client.Website.FilterConversations(websiteID string, pageNumber uint, filterInboxID string, filterUnread bool, filterResolved bool, filterNotResolved bool, filterMention bool, filterAssigned bool, filterUnassigned bool) (*[]Conversation, *Response, error)`
  * **List Suggested Conversation Segments**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-conversation-segments)
    * `client.Website.ListSuggestedConversationSegments(websiteID string, pageNumber uint) (*[]ConversationSuggestedSegment, *Response, error)`
  * **Delete Suggested Conversation Segment**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-conversation-segment)
    * `client.Website.DeleteSuggestedConversationSegment(websiteID string, segment string) (*Response, error)`
  * **List Suggested Conversation Data Keys**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-conversation-data-keys)
    * `client.Website.ListSuggestedConversationDataKeys(websiteID string, pageNumber uint) (*[]ConversationSuggestedData, *Response, error)`
  * **Delete Suggested Conversation Data Key**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-conversation-data-key)
    * `client.Website.DeleteSuggestedConversationDataKey(websiteID string, key string) (*Response, error)`
  * **List Spam Conversations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-spam-conversations)
    * `client.Website.ListSpamConversations(websiteID string, pageNumber uint) (*[]ConversationSpam, *Response, error)`
  * **Resolve Spam Conversation Content**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-spam-conversation-content)
    * `client.Website.ResolveSpamConversationContent(websiteID string, spamID string) (*ConversationSpamContent, *Response, error)`
  * **Submit Spam Conversation Decision**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#submit-spam-conversation-decision)
    * `client.Website.SubmitSpamConversationDecision(websiteID string, spamID string, action string) (*Response, error)`

* #### **Website Conversation**
  * **⭐ Create A New Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-conversation)
    * `client.Website.CreateNewConversation(websiteID string) (*ConversationNew, *Response, error)`
  * **Check If Conversation Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-conversation-exists)
    * `client.Website.CheckConversationExists(websiteID string, sessionID string) (*Response, error)`
  * **⭐ Get A Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-conversation)
    * `client.Website.GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error)`
  * **Remove A Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-conversation)
    * `client.Website.RemoveConversation(websiteID string, sessionID string) (*Response, error)`
  * **Initiate A Conversation With Existing Session**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initiate-a-conversation-with-existing-session)
    * `client.Website.InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error)`
  * **⭐ Get Messages In Conversation (Last Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-messages-in-conversation)
    * `client.Website.GetMessagesInConversationLast(websiteID string, sessionID string) (*[]ConversationMessage, *Response, error)`
  * **Get Messages In Conversation (Before Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-messages-in-conversation)
    * `client.Website.GetMessagesInConversationBefore(websiteID string, sessionID string, timestampBefore uint) (*[]ConversationMessage, *Response, error)`
  * **⭐ Send A Message In Conversation (Text Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **⭐ Send A Message In Conversation (File Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Animation Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendAnimationMessageInConversation(websiteID string, sessionID string, message ConversationAnimationMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Audio Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendAudioMessageInConversation(websiteID string, sessionID string, message ConversationAudioMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Picker Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendPickerMessageInConversation(websiteID string, sessionID string, message ConversationPickerMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Field Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendFieldMessageInConversation(websiteID string, sessionID string, message ConversationFieldMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Carousel Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendCarouselMessageInConversation(websiteID string, sessionID string, message ConversationCarouselMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **⭐ Send A Message In Conversation (Note Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendNoteMessageInConversation(websiteID string, sessionID string, message ConversationNoteMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Event Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendEventMessageInConversation(websiteID string, sessionID string, message ConversationEventMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Get A Message In Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-message-in-conversation)
    * `client.Website.GetMessageInConversation(websiteID string, sessionID string, fingerprint int) (*ConversationMessage, *Response, error)`
  * **Update A Message In Conversation (Text Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateTextMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (File Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateFileMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFileMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Animation Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateAnimationMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAnimationMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Audio Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateAudioMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAudioMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Picker Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdatePickerMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationPickerMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Field Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateFieldMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFieldMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Carousel Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateCarouselMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationCarouselMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Note Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateNoteMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (Event Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateEventMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationEventMessageNewContent) (*Response, error)`
  * **Remove A Message In Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-message-in-conversation)
    * `client.Website.RemoveMessageInConversation(websiteID string, sessionID string, fingerprint int) (*Response, error)`
  * **Compose A Message In Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#compose-a-message-in-conversation)
    * `client.Website.ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error)`
  * **⭐ Mark Messages As Read In Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#mark-messages-as-read-in-conversation)
    * `client.Website.MarkMessagesReadInConversation(websiteID string, sessionID string, read ConversationReadMessageMark) (*Response, error)`
  * **Mark Conversation As Unread**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#mark-conversation-as-unread)
    * `client.Website.MarkConversationAsUnread(websiteID string, sessionID string, unread ConversationUnreadMessageMark) (*Response, error)`
  * **⭐ Mark Messages As Delivered In Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#mark-messages-as-delivered-in-conversation)
    * `client.Website.MarkMessagesDeliveredInConversation(websiteID string, sessionID string, delivered ConversationDeliveredMessageMark) (*Response, error)`
  * **Update Conversation Open State**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-conversation-open-state)
    * `client.Website.UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error)`
  * **⭐ Get Conversation Routing Assign**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-routing-assign)
    * `client.Website.GetConversationRoutingAssign(websiteID string, sessionID string) (*ConversationRoutingAssign, *Response, error)`
  * **⭐ Assign Conversation Routing**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#assign-conversation-routing)
    * `client.Website.AssignConversationRouting(websiteID string, sessionID string, assign ConversationRoutingAssignUpdate) (*Response, error)`
  * **Update Conversation Inbox**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-conversation-inbox)
    * `client.Website.UpdateConversationInbox(websiteID string, sessionID string, inboxID *string) (*Response, error)`
  * **⭐ Get Conversation Metas**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-metas)
    * `client.Website.GetConversationMetas(websiteID string, sessionID string) (*ConversationMeta, *Response, error)`
  * **⭐ Update Conversation Metas**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-conversation-metas)
    * `client.Website.UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error)`
  * **Get An Original Message In Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-an-original-message-in-conversation)
    * `client.Website.GetOriginalMessageInConversation(websiteID string, sessionID string, originalID string) (*ConversationOriginal, *Response, error)`
  * **List Conversation Pages**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversation-pages)
    * `client.Website.ListConversationPages(websiteID string, sessionID string, pageNumber uint) (*[]ConversationPage, *Response, error)`
  * **List Conversation Events**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversation-events)
    * `client.Website.ListConversationEvents(websiteID string, sessionID string, pageNumber uint) (*[]ConversationEvent, *Response, error)`
  * **List Conversation Files**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversation-files)
    * `client.Website.ListConversationFiles(websiteID string, sessionID string, pageNumber uint) (*[]ConversationFile, *Response, error)`
  * **Get Conversation State**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-state)
    * `client.Website.GetConversationState(websiteID string, sessionID string) (*ConversationState, *Response, error)`
  * **⭐ Change Conversation State**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#change-conversation-state)
    * `client.Website.ChangeConversationState(websiteID string, sessionID string, state string, user *ConversationMessageUser, origin *string) (*Response, error)`
  * **Get Conversation Relations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-relations)
    * `client.Website.GetConversationRelations(websiteID string, sessionID string) (*ConversationRelations, *Response, error)`
  * **Get Conversation Participants**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-participants)
    * `client.Website.GetConversationParticipants(websiteID string, sessionID string) (*ConversationParticipants, *Response, error)`
  * **Save Conversation Participants**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-conversation-participants)
    * `client.Website.SaveConversationParticipants(websiteID string, sessionID string, participants ConversationParticipantsSave) (*Response, error)`
  * **Get Block Status For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-block-status-for-conversation)
    * `client.Website.GetBlockStatusForConversation(websiteID string, sessionID string) (*ConversationBlock, *Response, error)`
  * **Block Incoming Messages For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#block-incoming-messages-for-conversation)
    * `client.Website.BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool, user *ConversationMessageUser, origin *string) (*Response, error)`
  * **Get Verify Status For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-verify-status-for-conversation)
    * `client.Website.GetVerifyStatusForConversation(websiteID string, sessionID string) (*ConversationVerify, *Response, error)`
  * **Update Verify Status For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-verify-status-for-conversation)
    * `client.Website.UpdateVerifyStatusForConversation(websiteID string, sessionID string, verified bool) (*Response, error)`
  * **Request Identity Verification For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-identity-verification-for-conversation)
    * `client.Website.RequestIdentityVerificationForConversation(websiteID string, sessionID string, request ConversationVerifyIdentityRequest) (*Response, error)`
  * **Redeem Identity Verification Link For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#redeem-identity-verification-link-for-conversation)
    * `client.Website.RedeemIdentityVerificationLinkForConversation(websiteID string, sessionID string, redeem ConversationVerifyIdentityRedeem) (*Response, error)`
  * **Request Email Transcript For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-email-transcript-for-conversation)
    * `client.Website.RequestEmailTranscriptForConversation(websiteID string, sessionID string, to string, email string) (*Response, error)`
  * **Request Chatbox Binding Purge For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-chatbox-binding-purge-for-conversation)
    * `client.Website.RequestChatboxBindingPurgeForConversation(websiteID string, sessionID string) (*Response, error)`
  * **Request User Feedback For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-user-feedback-for-conversation)
    * `client.Website.RequestUserFeedbackForConversation(websiteID string, sessionID string) (*Response, error)`
  * **List Browsing Sessions For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-browsing-sessions-for-conversation)
    * `client.Website.ListBrowsingSessionsForConversation(websiteID string, sessionID string) (*[]ConversationBrowsing, *Response, error)`
  * **Initiate Browsing Session For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initiate-browsing-session-for-conversation)
    * `client.Website.InitiateBrowsingSessionForConversation(websiteID string, sessionID string) (*Response, error)`
  * **Send Action To An Existing Browsing Session**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-action-to-an-existing-browsing-session)
    * `client.Website.SendActionToExistingBrowsingSession(websiteID string, sessionID string, browsingID string, action string) (*Response, error)`
  * **Assist Existing Browsing Session**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#assist-an-existing-browsing-session)
    * `client.Website.AssistExistingBrowsingSession(websiteID string, sessionID string, browsingID string, assist ConversationBrowsingAssist) (*Response, error)`
  * **Initiate New Call Session For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initiate-new-call-session-for-conversation)
    * `client.Website.InitiateNewCallSessionForConversation(websiteID string, sessionID string, mode string, user *ConversationMessageUser, origin *string) (*ConversationCall, *Response, error)`
  * **Get Ongoing Call Session For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-ongoing-call-session-for-conversation)
    * `client.Website.GetOngoingCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Abort Ongoing Call Session For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#abort-ongoing-call-session-for-conversation)
    * `client.Website.AbortOngoingCallSessionForConversation(websiteID string, sessionID string, callID string, user *ConversationMessageUser, origin *string) (*Response, error)`
  * **Transmit Signaling On Ongoing Call Session**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#transmit-signaling-on-ongoing-call-session)
    * `client.Website.TransmitSignalingOnOngoingCallSession(websiteID string, sessionID string, callID string, payload ConversationCallSignalingPayload) (*Response, error)`
  * **Deliver Widget Button Action For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#deliver-widget-button-action-for-conversation)
    * `client.Website.DeliverWidgetButtonActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}, value *interface{}) (*ConversationWidgetAction, *Response, error)`
  * **Deliver Widget Data Fetch Action For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#deliver-widget-data-action-for-conversation)
    * `client.Website.DeliverWidgetDataFetchActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}) (*ConversationWidgetAction, *Response, error)`
  * **Deliver Widget Data Edit Action For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#deliver-widget-data-action-for-conversation)
    * `client.Website.DeliverWidgetDataEditActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, value string) (*ConversationWidgetAction, *Response, error)`
  * **Schedule A Reminder For Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#schedule-a-reminder-for-conversation)
    * `client.Website.ScheduleReminderForConversation(websiteID string, sessionID string, date string, note string, user *ConversationMessageUser, origin *string) (*Response, error)`
  * **Report Conversation**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#report-conversation)
    * `client.Website.ReportConversation(websiteID string, sessionID string, flag string) (*Response, error)`

* #### **Website Analytics**
  * **Generate Analytics**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#generate-analytics)
    * `client.Website.GenerateAnalytics(websiteID string, query AnalyticsGenerateQuery) (*[]WebsiteAnalyticsGeneratePoint, *Response, error)`

* #### **Website People**
  * **Get People Statistics**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-statistics)
    * `client.Website.GetPeopleStatistics(websiteID string) (*PeopleStatistics, *Response, error)`
  * **List Suggested People Segments**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-people-segments)
    * `client.Website.ListSuggestedPeopleSegments(websiteID string, pageNumber uint) (*[]PeopleSuggestedSegment, *Response, error)`
  * **Delete Suggested People Segment**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-people-segment)
    * `client.Website.DeleteSuggestedPeopleSegment(websiteID string, segment string) (*Response, error)`
  * **List Suggested People Data Keys**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-people-data-keys/)
    * `client.Website.ListSuggestedPeopleDataKeys(websiteID string, pageNumber uint) (*[]PeopleSuggestedData, *Response, error)`
  * **Delete Suggested People Data Key**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-people-data-key)
    * `client.Website.DeleteSuggestedPeopleDataKey(websiteID string, key string) (*Response, error)`
  * **List Suggested People Events**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-people-events)
    * `client.Website.ListSuggestedPeopleEvents(websiteID string, pageNumber uint) (*[]PeopleSuggestedEvent, *Response, error)`
  * **Delete Suggested People Event**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-people-event)
    * `client.Website.DeleteSuggestedPeopleEvent(websiteID string, text string) (*Response, error)`
  * **⭐ List People Profiles**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-profiles)
    * `client.Website.ListPeopleProfiles(websiteID string, pageNumber uint, searchField string, searchOrder string, searchOperator string, searchFilter []PeopleFilter, searchText string) (*[]PeopleProfile, *Response, error)`
  * **⭐ Add New People Profile**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-new-people-profile)
    * `client.Website.AddNewPeopleProfile(websiteID string, peopleProfile PeopleProfileUpdateCard) (*PeopleProfileNew, *Response, error)`
  * **⭐ Check If People Profile Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-people-profile-exists)
    * `client.Website.CheckPeopleProfileExists(websiteID string, peopleID string) (*Response, error)`
  * **⭐ Get People Profile**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-profile)
    * `client.Website.GetPeopleProfile(websiteID string, peopleID string) (*PeopleProfile, *Response, error)`
  * **⭐ Save People Profile**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-people-profile)
    * `client.Website.SavePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **⭐ Update People Profile**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-people-profile)
    * `client.Website.UpdatePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **⭐ Remove People Profile**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-people-profile)
    * `client.Website.RemovePeopleProfile(websiteID string, peopleID string) (*Response, error)`
  * **List People Conversations**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-conversations)
    * `client.Website.ListPeopleConversations(websiteID string, peopleID string, pageNumber uint) ([]string, *Response, error)`
  * **List People Conversations (Filter Variant)**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-conversations)
    * `client.Website.FilterPeopleConversations(websiteID string, peopleID string, pageNumber uint, filterUnread bool, filterResolved bool, filterNotResolved bool) ([]string, *Response, error)`
  * **List People Campaigns**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-campaigns)
    * `client.Website.ListPeopleCampaigns(websiteID string, peopleID string, pageNumber uint) (*[]PeopleCampaign, *Response, error)`
  * **Add A People Event**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-a-people-event)
    * `client.Website.AddPeopleEvent(websiteID string, peopleID string, peopleEvent PeopleEventAdd) (*Response, error)`
  * **List People Events**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-events)
    * `client.Website.ListPeopleEvents(websiteID string, peopleID string, pageNumber uint) (*[]PeopleEvent, *Response, error)`
  * **Get People Data**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-data)
    * `client.Website.GetPeopleData(websiteID string, peopleID string) (*PeopleData, *Response, error)`
  * **Save People Data**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-people-data)
    * `client.Website.SavePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error)`
  * **Update People Data**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-people-data)
    * `client.Website.UpdatePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error)`
  * **Get People Subscription Status**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-subscription-status)
    * `client.Website.GetPeopleSubscriptionStatus(websiteID string, peopleID string) (*PeopleSubscription, *Response, error)`
  * **Update People Subscription Status**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-people-subscription-status)
    * `client.Website.UpdatePeopleSubscriptionStatus(websiteID string, peopleID string, peopleSubscription PeopleSubscriptionUpdate)`
  * **Export People Profiles**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#export-people-profiles)
    * `client.Website.ExportPeopleProfiles(websiteID string) (*Response, error)`
  * **Import People Profiles**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#import-people-profiles)
    * `client.Website.ImportPeopleProfiles(websiteID string, profileImportSetup PeopleProfileImportSetup) (*PeopleProfileImport, *Response, error)`

* #### **Website Helpdesk**
  * **Check If Helpdesk Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-exists)
    * `client.Website.CheckHelpdeskExists(websiteID string) (*Response, error)`
  * **Resolve Helpdesk**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk)
    * `client.Website.ResolveHelpdesk(websiteID string) (*Helpdesk, *Response, error)`
  * **Initialize Helpdesk**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initialize-helpdesk)
    * `client.Website.InitializeHelpdesk(websiteID string, name string, domainBasic string) (*Response, error)`
  * **Delete Helpdesk**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk)
    * `client.Website.DeleteHelpdesk(websiteID string, verify HelpdeskRemoveVerify) (*Response, error)`
  * **List Helpdesk Locales**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locales)
    * `client.Website.ListHelpdeskLocales(websiteID string, pageNumber uint) (*[]HelpdeskLocale, *Response, error)`
  * **Add Helpdesk Locale**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-locale)
    * `client.Website.AddHelpdeskLocale(websiteID string, locale string) (*Response, error)`
  * **Check If Helpdesk Locale Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-exists)
    * `client.Website.CheckHelpdeskLocaleExists(websiteID string, locale string) (*Response, error)`
  * **Resolve Helpdesk Locale**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale)
    * `client.Website.ResolveHelpdeskLocale(websiteID string, locale string) (*HelpdeskLocale, *Response, error)`
  * **Delete Helpdesk Locale**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale)
    * `client.Website.DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error)`
  * **List Helpdesk Locale Articles**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-articles)
    * `client.Website.ListHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticle, *Response, error)`
  * **Add A New Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-a-new-helpdesk-locale-article)
    * `client.Website.AddNewHelpdeskLocaleArticle(websiteID string, locale string, title string) (*HelpdeskLocaleArticleNew, *Response, error)`
  * **Check If Helpdesk Locale Article Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-article-exists)
    * `client.Website.CheckHelpdeskLocaleArticleExists(websiteID string, locale string, articleId string) (*Response, error)`
  * **Resolve Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article)
    * `client.Website.ResolveHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticle, *Response, error)`
  * **Save Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-article)
    * `client.Website.SaveHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error)`
  * **Update Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-article)
    * `client.Website.UpdateHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error)`
  * **Delete Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-article)
    * `client.Website.DeleteHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error)`
  * **Resolve Helpdesk Locale Article Page**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article-page)
    * `client.Website.ResolveHelpdeskLocaleArticlePage(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticlePage, *Response, error)`
  * **Resolve Helpdesk Locale Article Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article-category)
    * `client.Website.ResolveHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticleCategory, *Response, error)`
  * **Update Helpdesk Locale Article Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-article-category)
    * `client.Website.UpdateHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string, categoryId string, sectionId *string) (*Response, error)`
  * **List Helpdesk Locale Article Alternates**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-article-alternates)
    * `client.Website.ListHelpdeskLocaleArticleAlternates(websiteID string, locale string, articleId string) (*[]HelpdeskLocaleArticleAlternate, *Response, error)`
  * **Check If Helpdesk Locale Article Alternate Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-article-alternate-exists)
    * `client.Website.CheckHelpdeskLocaleArticleAlternateExists(websiteID string, locale string, articleId string, localeLinked string) (*Response, error)`
  * **Resolve Helpdesk Locale Article Alternate**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article-alternate)
    * `client.Website.ResolveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*HelpdeskLocaleArticleAlternate, *Response, error)`
  * **Save Helpdesk Locale Article Alternate**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-article-alternate)
    * `client.Website.SaveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string, articleIdLinked string) (*Response, error)`
  * **Delete Helpdesk Locale Article Alternate**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-article-alternate)
    * `client.Website.DeleteHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*Response, error)`
  * **Publish Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#publish-helpdesk-locale-article)
    * `client.Website.PublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticlePublish, *Response, error)`
  * **Unpublish Helpdesk Locale Article**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unpublish-helpdesk-locale-article)
    * `client.Website.UnpublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error)`
  * **List Helpdesk Locale Categories**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-categories)
    * `client.Website.ListHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticleCategory, *Response, error)`
  * **Add Helpdesk Locale Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-locale-category)
    * `client.Website.AddHelpdeskLocaleCategory(websiteID string, locale string, name string) (*HelpdeskLocaleArticleCategoryNew, *Response, error)`
  * **Check If Helpdesk Locale Category Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-category-exists)
    * `client.Website.CheckHelpdeskLocaleCategoryExists(websiteID string, locale string, categoryId string) (*Response, error)`
  * **Resolve Helpdesk Locale Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-category)
    * `client.Website.ResolveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*HelpdeskLocaleArticleCategory, *Response, error)`
  * **Save Helpdesk Locale Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-category)
    * `client.Website.SaveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error)`
  * **Update Helpdesk Locale Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-category)
    * `client.Website.UpdateHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error)`
  * **Delete Helpdesk Locale Category**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-category)
    * `client.Website.DeleteHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*Response, error)`
  * **List Helpdesk Locale Sections**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-sections)
    * `client.Website.ListHelpdeskLocaleSections(websiteID string, locale string, categoryId string, pageNumber uint) (*[]HelpdeskLocaleSection, *Response, error)`
  * **Add Helpdesk Locale Section**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-locale-section)
    * `client.Website.AddHelpdeskLocaleSection(websiteID string, locale string, categoryId string, name string) (*HelpdeskLocaleSectionNew, *Response, error)`
  * **Check If Helpdesk Locale Section Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-section-exists)
    * `client.Website.CheckHelpdeskLocaleSectionExists(websiteID string, locale string, categoryId string, sectionId string) (*Response, error)`
  * **Resolve Helpdesk Locale Section**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-section)
    * `client.Website.ResolveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*HelpdeskLocaleSection, *Response, error)`
  * **Save Helpdesk Locale Section**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-section)
    * `client.Website.SaveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error)`
  * **Update Helpdesk Locale Section**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-section)
    * `client.Website.UpdateHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error)`
  * **Delete Helpdesk Locale Section**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-section)
    * `client.Website.DeleteHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*Response, error)`
  * **Map Helpdesk Locale Feedback Ratings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#map-helpdesk-locale-feedback-ratings)
    * `client.Website.MapHelpdeskLocaleFeedbackRatings(websiteID string, locale string, filterDateStart string, filterDateEnd string) (*HelpdeskLocaleFeedbackRatings, *Response, error)`
  * **List Helpdesk Locale Feedbacks**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-feedbacks)
    * `client.Website.ListHelpdeskLocaleFeedbacks(websiteID string, locale string, pageNumber uint, filterDateStart string, filterDateEnd string) (*[]HelpdeskLocaleFeedbackItem, *Response, error)`
  * **Import External Helpdesk To Locale**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#import-external-helpdesk-to-locale)
    * `client.Website.ImportExternalHelpdeskToLocale(websiteID string, locale string, helpdeskUrl string) (*Response, error)`
  * **Export Helpdesk Locale Articles**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#export-helpdesk-locale-articles)
    * `client.Website.ExportHelpdeskLocaleArticles(websiteID string, locale string) (*Response, error)`
  * **List Helpdesk Redirections**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-redirections)
    * `client.Website.ListHelpdeskRedirections(websiteID string, pageNumber uint) (*[]HelpdeskRedirection, *Response, error)`
  * **Add Helpdesk Redirection**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-redirection)
    * `client.Website.AddHelpdeskRedirection(websiteID string, redirectionPath string, redirectionTarget string) (*HelpdeskRedirectionNew, *Response, error)`
  * **Check If Helpdesk Redirection Exists**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-redirection-exists)
    * `client.Website.CheckHelpdeskRedirectionExists(websiteID string, redirectionId string) (*Response, error)`
  * **Resolve Helpdesk Redirection**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-redirection)
    * `client.Website.ResolveHelpdeskRedirection(websiteID string, redirectionId string) (*HelpdeskRedirection, *Response, error)`
  * **Delete Helpdesk Redirection**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-redirection)
    * `client.Website.DeleteHelpdeskRedirection(websiteID string, redirectionId string) (*Response, error)`
  * **Resolve Helpdesk Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-settings)
    * `client.Website.ResolveHelpdeskSettings(websiteID string) (*HelpdeskSettings, *Response, error)`
  * **Save Helpdesk Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-settings)
    * `client.Website.SaveHelpdeskSettings(websiteID string, settings HelpdeskSettings) (*Response, error)`
  * **Resolve Helpdesk Domain**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-domain)
    * `client.Website.ResolveHelpdeskDomain(websiteID string) (*HelpdeskDomain, *Response, error)`
  * **Request Helpdesk Domain Change**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-helpdesk-domain-change)
    * `client.Website.RequestHelpdeskDomainChange(websiteID string, basic *string, custom *string) (*Response, error)`
  * **Generate Helpdesk Domain Setup Flow**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#generate-helpdesk-domain-setup-flow)
    * `client.Website.GenerateHelpdeskDomainSetupFlow(websiteID string, custom string) (*HelpdeskDomainSetupFlow, *Response, error)`

* #### **Website Connect**
  * **⭐ Get Connect Endpoints**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-website-connect-endpoints)
    * `client.Website.GetConnectEndpoints(websiteID string) (*WebsiteConnectEndpoints, *Response, error)`

### Plugin

* #### **Plugin Subscription**
  * **List All Active Subscriptions**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-active-subscriptions)
    * `client.Plugin.ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error)`
  * **List Subscriptions For A Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-subscriptions-for-a-website)
    * `client.Plugin.ListSubscriptionsForWebsite(websiteID string) (*[]PluginSubscription, *Response, error)`
  * **Get Subscription Details**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-subscription-details)
    * `client.Plugin.GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error)`
  * **Subscribe Website To Plugin**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#subscribe-website-to-plugin)
    * `client.Plugin.SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error)`
  * **Unsubscribe Plugin From Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unsubscribe-plugin-from-website)
    * `client.Plugin.UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error)`
  * **Get Subscription Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-subscription-settings)
    * `client.Plugin.GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error)`
  * **Save Subscription Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-subscription-settings)
    * `client.Plugin.SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Update Subscription Settings**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-subscription-settings)
    * `client.Plugin.UpdateSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Get Plugin Usage Bills**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-plugin-usage-bills)
    * `client.Plugin.GetPluginUsageBills(websiteID string, pluginID string) (*[]PluginBillUsage, *Response, error)`
  * **Report Plugin Usage To Bill**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#report-plugin-usage-to-bill)
    * `client.Plugin.ReportPluginUsageToBill(websiteID string, pluginID string, usage PluginSubscriptionBillUsageReport) (*Response, error)`
  * **Get Plugin Attest Provenance**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-plugin-attest-provenance)
    * `client.Plugin.GetPluginAttestProvenance(websiteID string, pluginID string) (*PluginAttestProvenance, *Response, error)`
  * **Forward Plugin Payload To Channel**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#forward-plugin-payload-to-channel)
    * `client.Plugin.ForwardPluginPayloadToChannel(websiteID string, pluginID string, payload PluginSubscriptionChannelForward) (*Response, error)`
  * **Dispatch Plugin Event**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#dispatch-plugin-event)
    * `client.Plugin.DispatchPluginEvent(websiteID string, pluginID string, payload PluginSubscriptionEventDispatch) (*Response, error)`

* #### **Plugin Connect**
  * **⭐ Get Connect Account**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-connect-account)
    * `client.Plugin.GetConnectAccount() (*PluginConnectAccount, *Response, error)`
  * **⭐ Check Connect Session Validity**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-connect-session-validity)
    * `client.Plugin.CheckConnectSessionValidity() (*Response, error)`
  * **⭐ List All Connect Websites**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-connect-websites)
    * `client.Plugin.ListAllConnectWebsites(pageNumber uint, filterConfigured bool, includePlan bool) (*[]PluginConnectAllWebsites, *Response, error)`
  * **⭐ List Connect Websites Since**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-connect-websites)
    * `client.Plugin.ListConnectWebsitesSince(dateSince time.Time, filterConfigured bool, includePlan bool) (*[]PluginConnectWebsitesSince, *Response, error)`
  * **⭐ Get Connect Endpoints**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-connect-endpoints)
    * `client.Plugin.GetConnectEndpoints() (*PluginConnectEndpoints, *Response, error)`

### Plan

* #### **Plan Subscription**
  * **List All Active Plan Subscriptions**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-active-plan-subscriptions)
    * `client.Plan.ListAllActiveSubscriptions() (*[]PlanSubscription, *Response, error)`
  * **Get Plan Subscription For A Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-plan-subscription-for-a-website)
    * `client.Plan.GetPlanSubscriptionForWebsite(websiteID string) (*PlanSubscription, *Response, error)`
  * **Subscribe Website To Plan**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#subscribe-website-to-plan)
    * `client.Plan.SubscribeWebsiteToPlan(websiteID string, planID string) (*Response, error)`
  * **Unsubscribe Plan From Website**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unsubscribe-plan-from-website)
    * `client.Plan.UnsubscribePlanFromWebsite(websiteID string) (*Response, error)`
  * **Change Bill Period For Website Plan Subscription**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#change-bill-period-for-website-plan-subscription)
    * `client.Plan.ChangeBillPeriodForWebsitePlanSubscription(websiteID string, period string) (*Response, error)`
  * **Check Coupon Availability For Website Plan Subscription**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-coupon-availability-for-website-plan-subscription)
    * `client.Plan.CheckCouponAvailabilityForWebsitePlanSubscription(websiteID string, code string) (*PlanSubscriptionCoupon, *Response, error)`
  * **Redeem Coupon For Website Plan Subscription**: [Reference](https://docs.crisp.chat/references/rest-api/v1/#redeem-coupon-for-website-plan-subscription)
    * `client.Plan.RedeemCouponForWebsitePlanSubscription(websiteID string, code string) (*Response, error)`

## Realtime Events

You can bind to realtime events from Crisp, in order to get notified of incoming messages and updates in websites.

You won't receive any event if you don't explicitly subscribe to realtime events, as the library doesn't connect to the realtime backend automatically.

**There are two ways to receive realtime events:**

1. Using Web Hooks (**⭐ recommended**)
2. Using WebSockets with the RTM API

### Receive realtime events

#### Receive events over Web Hooks

To start listening for events and bind a handler, check out the [events over Web Hooks example](https://github.com/crisp-im/go-crisp-api/blob/master/examples/events_webhooks/main.go).

Plugin Web Hooks will need to be configured first for this to work. Check out our [Web Hooks Quickstart guide](https://docs.crisp.chat/guides/web-hooks/quickstart/) and our [Web Hooks Reference](https://docs.crisp.chat/references/web-hooks/v1/) to get started.

#### Receive events over WebSockets (RTM API)

To start listening for events and bind a handler, check out the [events over WebSockets example](https://github.com/crisp-im/go-crisp-api/blob/master/examples/events_websockets/main.go).

### Available realtime events

Available events are listed below:

* #### **Session Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#session-events)
  * **Session Update Availability**:
    * `session:update_availability`
  * **Session Update Verify**:
    * `session:update_verify`
  * **Session Request Initiated**:
    * `session:request:initiated`
  * **Session Set Email**:
    * `session:set_email`
  * **Session Set Phone**:
    * `session:set_phone`
  * **Session Set Address**:
    * `session:set_address`
  * **Session Set Subject**:
    * `session:set_subject`
  * **Session Set Avatar**:
    * `session:set_avatar`
  * **Session Set Nickname**:
    * `session:set_nickname`
  * **Session Set Origin**:
    * `session:set_origin`
  * **Session Set Data**:
    * `session:set_data`
  * **Session Sync Pages**:
    * `session:sync:pages`
  * **Session Sync Events**:
    * `session:sync:events`
  * **Session Sync Capabilities**:
    * `session:sync:capabilities`
  * **Session Sync Geolocation**:
    * `session:sync:geolocation`
  * **Session Sync System**:
    * `session:sync:system`
  * **Session Sync Network**:
    * `session:sync:network`
  * **Session Sync Timezone**:
    * `session:sync:timezone`
  * **Session Sync Locales**:
    * `session:sync:locales`
  * **Session Sync Rating**:
    * `session:sync:rating`
  * **Session Sync Topic**:
    * `session:sync:topic`
  * **Session Set State**:
    * `session:set_state`
  * **Session Set Block**:
    * `session:set_block`
  * **Session Set Segments**:
    * `session:set_segments`
  * **Session Set Opened**:
    * `session:set_opened`
  * **Session Set Closed**:
    * `session:set_closed`
  * **Session Set Participants**:
    * `session:set_participants`
  * **Session Set Mentions**:
    * `session:set_mentions`
  * **Session Set Routing**:
    * `session:set_routing`
  * **Session Set Inbox**:
    * `session:set_inbox`
  * **Session Removed**:
    * `session:removed`
  * **Session Error**:
    * `session:error`

* #### **Message Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#message-events)
  * **Message Updated**:
    * `message:updated`
  * **Message Send (Text Variant)**:
    * `message:send/text`
  * **Message Send (File Variant)**:
    * `message:send/file`
  * **Message Send (Animation Variant)**:
    * `message:send/animation`
  * **Message Send (Audio Variant)**:
    * `message:send/audio`
  * **Message Send (Picker Variant)**:
    * `message:send/picker`
  * **Message Send (Field Variant)**:
    * `message:send/field`
  * **Message Send (Carousel Variant)**:
    * `message:send/carousel`
  * **Message Send (Note Variant)**:
    * `message:send/note`
  * **Message Send (Event Variant)**:
    * `message:send/event`
  * **Message Received (Text Variant)**:
    * `message:received/text`
  * **Message Received (File Variant)**:
    * `message:received/file`
  * **Message Received (Animation Variant)**:
    * `message:received/animation`
  * **Message Received (Audio Variant)**:
    * `message:received/audio`
  * **Message Received (Picker Variant)**:
    * `message:received/picker`
  * **Message Received (Field Variant)**:
    * `message:received/field`
  * **Message Received (Carousel Variant)**:
    * `message:received/carousel`
  * **Message Received (Note Variant)**:
    * `message:received/note`
  * **Message Received (Event Variant)**:
    * `message:received/event`
  * **Message Removed**:
    * `message:removed`
  * **Message Compose Send**:
    * `message:compose:send`
  * **Message Compose Receive**:
    * `message:compose:receive`
  * **Message Acknowledge Read Send**:
    * `message:acknowledge:read:send`
  * **Message Acknowledge Read Received**:
    * `message:acknowledge:read:received`
  * **Message Acknowledge Delivered**:
    * `message:acknowledge:delivered`
  * **Message Acknowledge Ignored**:
    * `message:acknowledge:ignored`
  * **Message Notify Unread Send**:
    * `message:notify:unread:send`
  * **Message Notify Unread Received**:
    * `message:notify:unread:received`

* #### **Spam Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#spam-events)
  * **Spam Message**:
    * `spam:message`
  * **Spam Decision**:
    * `spam:decision`

* #### **People Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#people-events)
  * **People Profile Created**:
    * `people:profile:created`
  * **People Profile Updated**:
    * `people:profile:updated`
  * **People Profile Removed**:
    * `people:profile:removed`
  * **People Bind Session**:
    * `people:bind:session`
  * **People Sync Profile**:
    * `people:sync:profile`
  * **People Import Progress**:
    * `people:import:progress`
  * **People Import Done**:
    * `people:import:done`

* #### **Campaign Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#campaign-events)
  * **Campaign Progress**:
    * `campaign:progress`
  * **Campaign Dispatched**:
    * `campaign:dispatched`
  * **Campaign Running**:
    * `campaign:running`

* #### **Browsing Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#browsing-events)
  * **Browsing Request Initiated**:
    * `browsing:request:initiated`
  * **Browsing Request Rejected**:
    * `browsing:request:rejected`

* #### **Call Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#call-events)
  * **Call Request Initiated**:
    * `call:request:initiated`
  * **Call Request Rejected**:
    * `call:request:rejected`

* #### **Identity Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#identity-events)
  * **Identity Verify Request**:
    * `identity:verify:request`

* #### **Widget Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#widget-events)
  * **Widget Action Processed**:
    * `widget:action:processed`

* #### **Status Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#status-events)
  * **Status Health Changed**:
    * `status:health:changed`

* #### **Website Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#website-events)
  * **Website Update Visitors Count**:
    * `website:update_visitors_count`
  * **Website Update Operators Availability**:
    * `website:update_operators_availability`
  * **Website Users Available**:
    * `website:users:available`

* #### **Bucket Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#bucket-events)
  * **Bucket URL Upload Generated**:
    * `bucket:url:upload:generated`
  * **Bucket URL Avatar Generated**:
    * `bucket:url:avatar:generated`
  * **Bucket URL Website Generated**:
    * `bucket:url:website:generated`
  * **Bucket URL Campaign Generated**:
    * `bucket:url:campaign:generated`
  * **Bucket URL Helpdesk Generated**:
    * `bucket:url:helpdesk:generated`
  * **Bucket URL Status Generated**:
    * `bucket:url:status:generated`
  * **Bucket URL Processing Generated**:
    * `bucket:url:processing:generated`

* #### **Media Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#media-events)
  * **Media Animation Listed**:
    * `media:animation:listed`

* #### **Email Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#email-events)
  * **Email Subscribe**:
    * `email:subscribe`
  * **Email Track View**:
    * `email:track:view`

* #### **Plugin Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#plugin-events)
  * **Plugin Channel**:
    * `plugin:channel`
  * **Plugin Event**:
    * `plugin:event`
  * **Plugin Subscription Updated**:
    * `plugin:subscription:updated`
  * **Plugin Settings Saved**:
    * `plugin:settings:saved`

* #### **Plan Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#plan-events)
  * **Plan Subscription Updated**:
    * `plan:subscription:updated`
