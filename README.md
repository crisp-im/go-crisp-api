# go-crisp-api

[![Test and Build](https://github.com/crisp-im/go-crisp-api/workflows/Test%20and%20Build/badge.svg?branch=master)](https://github.com/crisp-im/go-crisp-api/actions?query=workflow%3A%22Test+and+Build%22)

The Crisp API Golang wrapper. Authenticate, send messages, fetch conversations, access your agent accounts from your Go code.

Copyright 2023 Crisp IM SAS. See LICENSE for copying information.

* **📝 Implements**: [REST API Reference (V1)](https://docs.crisp.chat/references/rest-api/v1/) at revision: 17/07/2023
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

// Authenticate to API with your plugin token (identifier, key)
// eg. client.AuthenticateTier("plugin", "43f34724-9eeb-4474-9cec-560250754dec", "d12e60c5d2aa264b90997a641b6474ffd6602b66d8e8abc49634c404f06fa7d0")
client.AuthenticateTier("plugin", identifier, key)

// Now, you can use authenticated API sections.
```

## Resource Methods

All the available Crisp API resources are fully implemented. **Programmatic methods names are named after their label name in the [REST API Reference](https://docs.crisp.chat/references/rest-api/v1/)**.

All methods that you will most likely need when building a Crisp integration are prefixed with a star symbol (⭐).

**⚠️ Note that, depending on your authentication token tier, which is either `user` or `plugin`, you may not be allowed to use all methods from the library. When in doubt, refer to the library method descriptions below. Most likely, you are using a `plugin` token.**

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
  * **Generate Bucket URL** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#bucket-url)
    * `client.Bucket.GenerateBucketURL(bucketData BucketURLRequest) (*Response, error)`

### Media

* #### **Media Animation**
  * **List Animation Medias** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-animation-medias)
    * `client.Media.ListAnimationMedias(pageNumber uint, listID int, searchQuery string) (*Response, error)`

### Website

* #### **Website Base**
  * **Check If Website Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-website-exists)
    * `client.Website.CheckWebsiteExists(domain string) (*Response, error)`
  * **Create Website** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-website)
    * `client.Website.CreateWebsite(websiteData WebsiteCreate) (*Website, *Response, error)`
  * **Get A Website** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-website)
    * `client.Website.GetWebsite(websiteID string) (*Website, *Response, error)`
  * **Delete A Website** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-a-website)
    * `client.Website.DeleteWebsite(websiteID string, verify string) (*Response, error)`

* #### **Website Batch**
  * **Batch Resolve Conversations** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-resolve-items)
    * `client.Website.BatchResolveConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Read Conversations** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-read-items)
    * `client.Website.BatchReadConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Remove Conversations** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-remove-items)
    * `client.Website.BatchRemoveConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Remove People** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#batch-remove-items)
    * `client.Website.BatchRemovePeople(websiteID string, people WebsiteBatchPeopleOperationInner) (*Response, error)`

* #### **Website Availability**
  * **Get Website Availability Status** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-website-availability-status)
    * `client.Website.GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error)`
  * **List Website Operator Availabilities** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-website-operator-availabilities)
    * `client.Website.ListWebsiteOperatorAvailabilities(websiteID string) (*[]WebsiteAvailabilityOperator, *Response, error)`

* #### **Website Operator**
  * **List Website Operators** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-website-operators)
    * `client.Website.ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error)`
  * **List Last Active Website Operators** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-last-active-website-operators)
    * `client.Website.ListLastActiveWebsiteOperators(websiteID string) (*[]WebsiteOperatorsLastActiveListOne, *Response, error)`
  * **Flush Last Active Website Operators** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#flush-last-active-website-operators)
    * `client.Website.FlushLastActiveWebsiteOperators(websiteID string) (*Response, error)`
  * **Send Email To Website Operators** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-email-to-website-operators)
    * `client.Website.SendEmailToWebsiteOperators(websiteID string, email WebsiteOperatorEmail) (*Response, error)`
  * **Get A Website Operator** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-website-operator)
    * `client.Website.GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error)`
  * **Invite A Website Operator** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#invite-a-website-operator)
    * `client.Website.InviteWebsiteOperator(websiteID string, email string, role string, verify string) (*Response, error)`
  * **Change Operator Membership** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#change-operator-membership)
    * `client.Website.ChangeOperatorMembership(websiteID string, userID string, role string, title *string) (*Response, error)`
  * **Unlink Operator From Website** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unlink-operator-from-website)
    * `client.Website.UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error)`

* #### **Website Verify**
  * **Get Verify Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-verify-settings)
    * `client.Website.GetVerifySettings(websiteID string) (*WebsiteVerifySettings, *Response, error)`
  * **Update Verify Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-verify-settings)
    * `client.Website.UpdateVerifySettings(websiteID string, settings WebsiteVerifySettingsUpdate) (*Response, error)`
  * **Get Verify Key** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-verify-key)
    * `client.Website.GetVerifyKey(websiteID string) (*WebsiteVerifyKey, *Response, error)`
  * **Roll Verify Key** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#roll-verify-key)
    * `client.Website.RollVerifyKey(websiteID string) (*Response, error)`

* #### **Website Settings**
  * **Get Website Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-website-settings)
    * `client.Website.GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error)`
  * **Update Website Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-website-settings)
    * `client.Website.UpdateWebsiteSettings(websiteID string, settings WebsiteSettingsUpdate) (*Response, error)`

* #### **Website Visitors**
  * **Count Visitors** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#count-visitors)
    * `client.Website.CountVisitors(websiteID string) (*WebsiteVisitorCount, *Response, error)`
  * **List Visitors** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-visitors)
    * `client.Website.ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error)`
  * **Pinpoint Visitors On A Map (Wide Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#pinpoint-visitors-on-a-map)
    * `client.Website.PinpointVisitorsOnMapWide(websiteID string) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Pinpoint Visitors On A Map (Area Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#pinpoint-visitors-on-a-map)
    * `client.Website.PinpointVisitorsOnMapArea(websiteID string, centerLongitude float32, centerLatitude float32, centerRadius uint) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Get Session Identifier From Token** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-session-identifier-from-token)
    * `client.Website.GetSessionIdentifierFromToken(websiteID string, tokenID string) (*WebsiteVisitorsToken, *Response, error)`
  * **Count Blocked Visitors** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#count-blocked-visitors/)
    * `client.Website.CountBlockedVisitors(websiteID string) (*[]WebsiteVisitorsBlocked, *Response, error)`
  * **Count Blocked Visitors In Rule** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#count-blocked-visitors-in-rule)
    * `client.Website.CountBlockedVisitorsInRule(websiteID string, rule string) (*WebsiteVisitorsBlocked, *Response, error)`
  * **Clear Blocked Visitors In Rule** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#clear-blocked-visitors-in-rule)
    * `client.Website.ClearBlockedVisitorsInRule(websiteID string, rule string) (*Response, error)`

* #### **Website Campaigns**
  * **List Campaigns** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaigns)
    * `client.Website.ListCampaigns(websiteID string, pageNumber uint) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaigns (Filter Variant)** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaigns)
    * `client.Website.FilterCampaigns(websiteID string, pageNumber uint, searchName string, filterTag string, filterTypeOneShot bool, filterTypeAutomated bool, filterStatusNotConfigured bool, filterStatusReady bool, filterStatusPaused bool, filterStatusSending bool, filterStatusDone bool) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaign Tags** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-tags)
    * `client.Website.ListCampaignTags(websiteID string) (*[]string, *Response, error)`
  * **List Campaign Templates** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-templates)
    * `client.Website.ListCampaignTemplates(websiteID string, pageNumber uint) (*[]WebsiteCampaignTemplateExcerpt, *Response, error)`
  * **List Campaign Templates (Filter Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-templates)
    * `client.Website.FilterCampaignTemplates(websiteID string, pageNumber uint, searchName string, filterTypeStatic bool, filterTypeCustom bool) (*[]WebsiteCampaignTemplateExcerpt, *Response, error)`
  * **Create A New Campaign Template** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-campaign-template)
    * `client.Website.CreateNewCampaignTemplate(websiteID string, templateFormat string, templateName string) (*WebsiteCampaignTemplateNew, *Response, error)`
  * **Check If Campaign Template Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-campaign-template-exists)
    * `client.Website.CheckCampaignTemplateExists(websiteID string, templateID string) (*Response, error)`
  * **Get A Campaign Template** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-campaign-template)
    * `client.Website.GetCampaignTemplate(websiteID string, templateID string) (*WebsiteCampaignTemplateItem, *Response, error)`
  * **Save A Campaign Template** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-a-campaign-template)
    * `client.Website.SaveCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error)`
  * **Update A Campaign Template** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-campaign-template)
    * `client.Website.UpdateCampaignTemplate(websiteID string, templateID string, websiteCampaignTemplateItem WebsiteCampaignTemplateItem) (*Response, error)`
  * **Remove A Campaign Template** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-campaign-template)
    * `client.Website.RemoveCampaignTemplate(websiteID string, templateID string) (*Response, error)`

* #### **Website Campaign**
  * **Create A New Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-campaign)
    * `client.Website.CreateNewCampaign(websiteID string, campaignType string, campaignName string) (*WebsiteCampaignNew, *Response, error)`
  * **Check If Campaign Exists** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-campaign-exists)
    * `client.Website.CheckCampaignExists(websiteID string, campaignID string) (*Response, error)`
  * **Get A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-campaign)
    * `client.Website.GetCampaign(websiteID string, campaignID string) (*WebsiteCampaignItem, *Response, error)`
  * **Save A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-a-campaign)
    * `client.Website.SaveCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Update A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-campaign)
    * `client.Website.UpdateCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Remove A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-campaign)
    * `client.Website.RemoveCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Dispatch A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#dispatch-a-campaign)
    * `client.Website.DispatchCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Resume A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resume-a-campaign)
    * `client.Website.ResumeCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Pause A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#pause-a-campaign)
    * `client.Website.PauseCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Test A Campaign** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#test-a-campaign)
    * `client.Website.TestCampaign(websiteID string, campaignID string) (*Response, error)`
  * **List Campaign Recipients** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-recipients)
    * `client.Website.ListCampaignRecipients(websiteID string, campaignID string, pageNumber uint) (*[]WebsiteCampaignRecipient, *Response, error)`
  * **List Campaign Statistics** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-campaign-statistics)
    * `client.Website.ListCampaignStatistics(websiteID string, campaignID string, action string, pageNumber uint) (*[]WebsiteCampaignStatistic, *Response, error)`

* #### **Website Conversations**
  * **⭐ List Conversations** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversations)
    * `client.Website.ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error)`
  * **List Conversations (Search Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversations)
    * `client.Website.SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error)`
  * **List Conversations (Filter Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversations)
    * `client.Website.FilterConversations(websiteID string, pageNumber uint, filterUnread bool, filterResolved bool, filterNotResolved bool, filterMention bool, filterAssigned bool, filterUnassigned bool) (*[]Conversation, *Response, error)`
  * **List Suggested Conversation Segments** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-conversation-segments)
    * `client.Website.ListSuggestedConversationSegments(websiteID string, pageNumber uint) (*[]ConversationSuggestedSegment, *Response, error)`
  * **Delete Suggested Conversation Segment** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-conversation-segment)
    * `client.Website.DeleteSuggestedConversationSegment(websiteID string, segment string) (*Response, error)`
  * **List Suggested Conversation Data Keys** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-conversation-data-keys)
    * `client.Website.ListSuggestedConversationDataKeys(websiteID string, pageNumber uint) (*[]ConversationSuggestedData, *Response, error)`
  * **Delete Suggested Conversation Data Key** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-conversation-data-key)
    * `client.Website.DeleteSuggestedConversationDataKey(websiteID string, key string) (*Response, error)`

* #### **Website Conversation**
  * **⭐ Create A New Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#create-a-new-conversation)
    * `client.Website.CreateNewConversation(websiteID string) (*ConversationNew, *Response, error)`
  * **Check If Conversation Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-conversation-exists)
    * `client.Website.CheckConversationExists(websiteID string, sessionID string) (*Response, error)`
  * **⭐ Get A Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-conversation)
    * `client.Website.GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error)`
  * **Remove A Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-conversation)
    * `client.Website.RemoveConversation(websiteID string, sessionID string) (*Response, error)`
  * **Initiate A Conversation With Existing Session** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initiate-a-conversation-with-existing-session)
    * `client.Website.InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error)`
  * **⭐ Get Messages In Conversation (Last Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-messages-in-conversation)
    * `client.Website.GetMessagesInConversationLast(websiteID string, sessionID string) (*[]ConversationMessage, *Response, error)`
  * **Get Messages In Conversation (Before Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-messages-in-conversation)
    * `client.Website.GetMessagesInConversationBefore(websiteID string, sessionID string, timestampBefore uint) (*[]ConversationMessage, *Response, error)`
  * **⭐ Send A Message In Conversation (Text Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **⭐ Send A Message In Conversation (File Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Animation Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendAnimationMessageInConversation(websiteID string, sessionID string, message ConversationAnimationMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Audio Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendAudioMessageInConversation(websiteID string, sessionID string, message ConversationAudioMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Picker Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendPickerMessageInConversation(websiteID string, sessionID string, message ConversationPickerMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Field Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendFieldMessageInConversation(websiteID string, sessionID string, message ConversationFieldMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Carousel Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendCarouselMessageInConversation(websiteID string, sessionID string, message ConversationCarouselMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **⭐ Send A Message In Conversation (Note Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendNoteMessageInConversation(websiteID string, sessionID string, message ConversationNoteMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Event Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-a-message-in-conversation)
    * `client.Website.SendEventMessageInConversation(websiteID string, sessionID string, message ConversationEventMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Get A Message In Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-a-message-in-conversation)
    * `client.Website.GetMessageInConversation(websiteID string, sessionID string, fingerprint int) (*ConversationMessage, *Response, error)`
  * **Update A Message In Conversation (Text Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateTextMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (File Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateFileMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFileMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Animation Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateAnimationMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAnimationMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Audio Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateAudioMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAudioMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Picker Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdatePickerMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationPickerMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Field Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateFieldMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFieldMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Carousel Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateCarouselMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationCarouselMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Note Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateNoteMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (Event Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-a-message-in-conversation)
    * `client.Website.UpdateEventMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationEventMessageNewContent) (*Response, error)`
  * **Remove A Message In Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-a-message-in-conversation)
    * `client.Website.RemoveMessageInConversation(websiteID string, sessionID string, fingerprint int) (*Response, error)`
  * **Compose A Message In Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#compose-a-message-in-conversation)
    * `client.Website.ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error)`
  * **⭐ Mark Messages As Read In Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#mark-messages-as-read-in-conversation)
    * `client.Website.MarkMessagesReadInConversation(websiteID string, sessionID string, read ConversationReadMessageMark) (*Response, error)`
  * **⭐ Mark Messages As Delivered In Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#mark-messages-as-delivered-in-conversation)
    * `client.Website.MarkMessagesDeliveredInConversation(websiteID string, sessionID string, delivered ConversationDeliveredMessageMark) (*Response, error)`
  * **Update Conversation Open State** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-conversation-open-state)
    * `client.Website.UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error)`
  * **⭐ Get Conversation Routing Assign** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-routing-assign)
    * `client.Website.GetConversationRoutingAssign(websiteID string, sessionID string) (*ConversationRoutingAssign, *Response, error)`
  * **⭐ Assign Conversation Routing** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#assign-conversation-routing)
    * `client.Website.AssignConversationRouting(websiteID string, sessionID string, assign ConversationRoutingAssignUpdate) (*Response, error)`
  * **⭐ Get Conversation Metas** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-metas)
    * `client.Website.GetConversationMetas(websiteID string, sessionID string) (*ConversationMeta, *Response, error)`
  * **⭐ Update Conversation Metas** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-conversation-metas)
    * `client.Website.UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error)`
  * **Get An Original Message In Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-an-original-message-in-conversation)
    * `client.Website.GetOriginalMessageInConversation(websiteID string, sessionID string, originalID string) (*ConversationOriginal, *Response, error)`
  * **List Conversation Pages** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversation-pages)
    * `client.Website.ListConversationPages(websiteID string, sessionID string, pageNumber uint) (*[]ConversationPage, *Response, error)`
  * **List Conversation Events** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversation-events)
    * `client.Website.ListConversationEvents(websiteID string, sessionID string, pageNumber uint) (*[]ConversationEvent, *Response, error)`
  * **List Conversation Files** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-conversation-files)
    * `client.Website.ListConversationFiles(websiteID string, sessionID string, pageNumber uint) (*[]ConversationFile, *Response, error)`
  * **Get Conversation State** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-state)
    * `client.Website.GetConversationState(websiteID string, sessionID string) (*ConversationState, *Response, error)`
  * **⭐ Change Conversation State** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#change-conversation-state)
    * `client.Website.ChangeConversationState(websiteID string, sessionID string, state string) (*Response, error)`
  * **Get Conversation Participants** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-conversation-participants)
    * `client.Website.GetConversationParticipants(websiteID string, sessionID string) (*ConversationParticipants, *Response, error)`
  * **Save Conversation Participants** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-conversation-participants)
    * `client.Website.SaveConversationParticipants(websiteID string, sessionID string, participants ConversationParticipantsSave) (*Response, error)`
  * **Get Block Status For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-block-status-for-conversation)
    * `client.Website.GetBlockStatusForConversation(websiteID string, sessionID string) (*ConversationBlock, *Response, error)`
  * **Block Incoming Messages For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#block-incoming-messages-for-conversation)
    * `client.Website.BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool) (*Response, error)`
  * **Get Verify Status For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-verify-status-for-conversation)
    * `client.Website.GetVerifyStatusForConversation(websiteID string, sessionID string) (*ConversationVerify, *Response, error)`
  * **Update Verify Status For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-verify-status-for-conversation)
    * `client.Website.UpdateVerifyStatusForConversation(websiteID string, sessionID string, verified bool) (*Response, error)`
  * **Request Email Transcript For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-email-transcript-for-conversation)
    * `client.Website.RequestEmailTranscriptForConversation(websiteID string, sessionID string, to string, email string) (*Response, error)`
  * **Request Chatbox Binding Purge For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-chatbox-binding-purge-for-conversation)
    * `client.Website.RequestChatboxBindingPurgeForConversation(websiteID string, sessionID string) (*Response, error)`
  * **Request User Feedback For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-user-feedback-for-conversation)
    * `client.Website.RequestUserFeedbackForConversation(websiteID string, sessionID string) (*Response, error)`
  * **List Browsing Sessions For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-browsing-sessions-for-conversation)
    * `client.Website.ListBrowsingSessionsForConversation(websiteID string, sessionID string) (*[]ConversationBrowsing, *Response, error)`
  * **Initiate Browsing Session For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initiate-browsing-session-for-conversation)
    * `client.Website.InitiateBrowsingSessionForConversation(websiteID string, sessionID string) (*Response, error)`
  * **Send Action To An Existing Browsing Session** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#send-action-to-an-existing-browsing-session)
    * `client.Website.SendActionToExistingBrowsingSession(websiteID string, sessionID string, browsingID string, action string) (*Response, error)`
  * **Assist Existing Browsing Session** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#assist-an-existing-browsing-session)
    * `client.Website.AssistExistingBrowsingSession(websiteID string, sessionID string, browsingID string, assist ConversationBrowsingAssist) (*Response, error)`
  * **Initiate New Call Session For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initiate-new-call-session-for-conversation)
    * `client.Website.InitiateNewCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Get Ongoing Call Session For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-ongoing-call-session-for-conversation)
    * `client.Website.GetOngoingCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Abort Ongoing Call Session For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#abort-ongoing-call-session-for-conversation)
    * `client.Website.AbortOngoingCallSessionForConversation(websiteID string, sessionID string, callID string) (*Response, error)`
  * **Transmit Signaling On Ongoing Call Session** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#transmit-signaling-on-ongoing-call-session)
    * `client.Website.TransmitSignalingOnOngoingCallSession(websiteID string, sessionID string, callID string, payload ConversationCallSignalingPayload) (*Response, error)`
  * **Deliver Widget Button Action For Conversation** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#deliver-widget-button-action-for-conversation)
    * `client.Website.DeliverWidgetButtonActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}, value *interface{}) (*ConversationWidgetAction, *Response, error)`
  * **Deliver Widget Data Fetch Action For Conversation** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#deliver-widget-data-action-for-conversation)
    * `client.Website.DeliverWidgetDataFetchActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, data interface{}) (*ConversationWidgetAction, *Response, error)`
  * **Deliver Widget Data Edit Action For Conversation** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#deliver-widget-data-action-for-conversation)
    * `client.Website.DeliverWidgetDataEditActionForConversation(websiteID string, sessionID string, pluginID string, sectionID string, itemID string, value string) (*ConversationWidgetAction, *Response, error)`
  * **Schedule A Reminder For Conversation** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#schedule-a-reminder-for-conversation)
    * `client.Website.ScheduleReminderForConversation(websiteID string, sessionID string, date string, note string) (*Response, error)`

* #### **Website Analytics**
  * **Acquire Analytics Points** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#acquire-analytics-points)
    * `client.Website.AcquireAnalyticsPoints(websiteID string, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time, dateSplit string, classifier string, filterPrimary string, filterSecondary string, filterTertiary string) (*WebsiteAnalyticsPoints, *Response, error)`
  * **List Analytics Filters** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-analytics-filters)
    * `client.Website.ListAnalyticsFilters(websiteID string, pageNumber uint, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time) (*[]WebsiteAnalyticsFilter, *Response, error)`
  * **List Analytics Classifiers** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-analytics-classifiers)
    * `client.Website.ListAnalyticsClassifiers(websiteID string, pageNumber uint, pointType string, pointMetric string, dateFrom time.Time, dateTo time.Time) (*[]WebsiteAnalyticsClassifier, *Response, error)`

* #### **Website People**
  * **Get People Statistics** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-statistics)
    * `client.Website.GetPeopleStatistics(websiteID string) (*PeopleStatistics, *Response, error)`
  * **List Suggested People Segments** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-people-segments)
    * `client.Website.ListSuggestedPeopleSegments(websiteID string, pageNumber uint) (*[]PeopleSuggestedSegment, *Response, error)`
  * **Delete Suggested People Segment** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-people-segment)
    * `client.Website.DeleteSuggestedPeopleSegment(websiteID string, segment string) (*Response, error)`
  * **List Suggested People Data Keys** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-people-data-keys/)
    * `client.Website.ListSuggestedPeopleDataKeys(websiteID string, pageNumber uint) (*[]PeopleSuggestedData, *Response, error)`
  * **Delete Suggested People Data Key** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-people-data-key)
    * `client.Website.DeleteSuggestedPeopleDataKey(websiteID string, key string) (*Response, error)`
  * **List Suggested People Events** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-suggested-people-events)
    * `client.Website.ListSuggestedPeopleEvents(websiteID string, pageNumber uint) (*[]PeopleSuggestedEvent, *Response, error)`
  * **Delete Suggested People Event** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-suggested-people-event)
    * `client.Website.DeleteSuggestedPeopleEvent(websiteID string, text string) (*Response, error)`
  * **⭐ List People Profiles** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-profiles)
    * `client.Website.ListPeopleProfiles(websiteID string, pageNumber uint, searchField string, searchOrder string, searchOperator string, searchFilter []PeopleFilter, searchText string) (*[]PeopleProfile, *Response, error)`
  * **⭐ Add New People Profile** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-new-people-profile)
    * `client.Website.AddNewPeopleProfile(websiteID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **⭐ Check If People Profile Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-people-profile-exists)
    * `client.Website.CheckPeopleProfileExists(websiteID string, peopleID string) (*Response, error)`
  * **⭐ Get People Profile** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-profile)
    * `client.Website.GetPeopleProfile(websiteID string, peopleID string) (*PeopleProfile, *Response, error)`
  * **⭐ Save People Profile** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-people-profile)
    * `client.Website.SavePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **⭐ Update People Profile** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-people-profile)
    * `client.Website.UpdatePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **⭐ Remove People Profile** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#remove-people-profile)
    * `client.Website.RemovePeopleProfile(websiteID string, peopleID string) (*Response, error)`
  * **List People Conversations** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-conversations)
    * `client.Website.ListPeopleConversations(websiteID string, peopleID string, pageNumber uint) ([]string, *Response, error)`
  * **List People Conversations (Filter Variant)** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-conversations)
    * `client.Website.FilterPeopleConversations(websiteID string, peopleID string, pageNumber uint, filterUnread bool, filterResolved bool, filterNotResolved bool) ([]string, *Response, error)`
  * **List People Campaigns** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-campaigns)
    * `client.Website.ListPeopleCampaigns(websiteID string, peopleID string, pageNumber uint) (*[]PeopleCampaign, *Response, error)`
  * **Add A People Event** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-a-people-event)
    * `client.Website.AddPeopleEvent(websiteID string, peopleID string, peopleEvent PeopleEventAdd) (*Response, error)`
  * **List People Events** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-people-events)
    * `client.Website.ListPeopleEvents(websiteID string, peopleID string, pageNumber uint) (*[]PeopleEvent, *Response, error)`
  * **Get People Data** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-data)
    * `client.Website.GetPeopleData(websiteID string, peopleID string) (*PeopleData, *Response, error)`
  * **Save People Data** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-people-data)
    * `client.Website.SavePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error)`
  * **Update People Data** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-people-data)
    * `client.Website.UpdatePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error)`
  * **Get People Subscription Status** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-people-subscription-status)
    * `client.Website.GetPeopleSubscriptionStatus(websiteID string, peopleID string) (*PeopleSubscription, *Response, error)`
  * **Update People Subscription Status** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-people-subscription-status)
    * `client.Website.UpdatePeopleSubscriptionStatus(websiteID string, peopleID string, peopleSubscription PeopleSubscriptionUpdate)`
  * **Export People Profiles** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#export-people-profiles)
    * `client.Website.ExportPeopleProfiles(websiteID string) (*Response, error)`
  * **Import People Profiles** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#import-people-profiles)
    * `client.Website.ImportPeopleProfiles(websiteID string, profileImportSetup PeopleProfileImportSetup) (*PeopleProfileImport, *Response, error)`

* #### **Website Helpdesk**
  * **Check If Helpdesk Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-exists)
    * `client.Website.CheckHelpdeskExists(websiteID string) (*Response, error)`
  * **Resolve Helpdesk** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk)
    * `client.Website.ResolveHelpdesk(websiteID string) (*Helpdesk, *Response, error)`
  * **Initialize Helpdesk** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#initialize-helpdesk)
    * `client.Website.InitializeHelpdesk(websiteID string, name string, domainBasic string) (*Response, error)`
  * **Delete Helpdesk** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk)
    * `client.Website.DeleteHelpdesk(websiteID string, verify string) (*Response, error)`
  * **List Helpdesk Locales** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locales)
    * `client.Website.ListHelpdeskLocales(websiteID string, pageNumber uint) (*[]HelpdeskLocale, *Response, error)`
  * **Add Helpdesk Locale** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-locale)
    * `client.Website.AddHelpdeskLocale(websiteID string, locale string) (*Response, error)`
  * **Check If Helpdesk Locale Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-exists)
    * `client.Website.CheckHelpdeskLocaleExists(websiteID string, locale string) (*Response, error)`
  * **Resolve Helpdesk Locale** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale)
    * `client.Website.ResolveHelpdeskLocale(websiteID string, locale string) (*HelpdeskLocale, *Response, error)`
  * **Delete Helpdesk Locale** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale)
    * `client.Website.DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error)`
  * **List Helpdesk Locale Articles** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-articles)
    * `client.Website.ListHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticle, *Response, error)`
  * **Add A New Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-a-new-helpdesk-locale-article)
    * `client.Website.AddNewHelpdeskLocaleArticle(websiteID string, locale string, title string) (*HelpdeskLocaleArticleNew, *Response, error)`
  * **Check If Helpdesk Locale Article Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-article-exists)
    * `client.Website.CheckHelpdeskLocaleArticleExists(websiteID string, locale string, articleId string) (*Response, error)`
  * **Resolve Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article)
    * `client.Website.ResolveHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticle, *Response, error)`
  * **Save Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-article)
    * `client.Website.SaveHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error)`
  * **Update Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-article)
    * `client.Website.UpdateHelpdeskLocaleArticle(websiteID string, locale string, articleId string, article HelpdeskLocaleArticle) (*Response, error)`
  * **Delete Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-article)
    * `client.Website.DeleteHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error)`
  * **Resolve Helpdesk Locale Article Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article-category)
    * `client.Website.ResolveHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticleCategory, *Response, error)`
  * **Update Helpdesk Locale Article Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-article-category)
    * `client.Website.UpdateHelpdeskLocaleArticleCategory(websiteID string, locale string, articleId string, categoryId string, sectionId *string) (*Response, error)`
  * **List Helpdesk Locale Article Alternates** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-article-alternates)
    * `client.Website.ListHelpdeskLocaleArticleAlternates(websiteID string, locale string, articleId string) (*[]HelpdeskLocaleArticleAlternate, *Response, error)`
  * **Check If Helpdesk Locale Article Alternate Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-article-alternate-exists)
    * `client.Website.CheckHelpdeskLocaleArticleAlternateExists(websiteID string, locale string, articleId string, localeLinked string) (*Response, error)`
  * **Resolve Helpdesk Locale Article Alternate** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-article-alternate)
    * `client.Website.ResolveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*HelpdeskLocaleArticleAlternate, *Response, error)`
  * **Save Helpdesk Locale Article Alternate** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-article-alternate)
    * `client.Website.SaveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string, articleIdLinked string) (*Response, error)`
  * **Delete Helpdesk Locale Article Alternate** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-article-alternate)
    * `client.Website.DeleteHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleId string, localeLinked string) (*Response, error)`
  * **Publish Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#publish-helpdesk-locale-article)
    * `client.Website.PublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*HelpdeskLocaleArticlePublish, *Response, error)`
  * **Unpublish Helpdesk Locale Article** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unpublish-helpdesk-locale-article)
    * `client.Website.UnpublishHelpdeskLocaleArticle(websiteID string, locale string, articleId string) (*Response, error)`
  * **List Helpdesk Locale Categories** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-categories)
    * `client.Website.ListHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint) (*[]HelpdeskLocaleArticleCategory, *Response, error)`
  * **Add Helpdesk Locale Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-locale-category)
    * `client.Website.AddHelpdeskLocaleCategory(websiteID string, locale string, name string) (*HelpdeskLocaleArticleCategoryNew, *Response, error)`
  * **Check If Helpdesk Locale Category Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-category-exists)
    * `client.Website.CheckHelpdeskLocaleCategoryExists(websiteID string, locale string, categoryId string) (*Response, error)`
  * **Resolve Helpdesk Locale Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-category)
    * `client.Website.ResolveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*HelpdeskLocaleArticleCategory, *Response, error)`
  * **Save Helpdesk Locale Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-category)
    * `client.Website.SaveHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error)`
  * **Update Helpdesk Locale Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-category)
    * `client.Website.UpdateHelpdeskLocaleCategory(websiteID string, locale string, categoryId string, category HelpdeskLocaleArticleCategory) (*Response, error)`
  * **Delete Helpdesk Locale Category** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-category)
    * `client.Website.DeleteHelpdeskLocaleCategory(websiteID string, locale string, categoryId string) (*Response, error)`
  * **List Helpdesk Locale Sections** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-sections)
    * `client.Website.ListHelpdeskLocaleSections(websiteID string, locale string, categoryId string, pageNumber uint) (*[]HelpdeskLocaleSection, *Response, error)`
  * **Add Helpdesk Locale Section** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-locale-section)
    * `client.Website.AddHelpdeskLocaleSection(websiteID string, locale string, categoryId string, name string) (*HelpdeskLocaleSectionNew, *Response, error)`
  * **Check If Helpdesk Locale Section Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-locale-section-exists)
    * `client.Website.CheckHelpdeskLocaleSectionExists(websiteID string, locale string, categoryId string, sectionId string) (*Response, error)`
  * **Resolve Helpdesk Locale Section** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-locale-section)
    * `client.Website.ResolveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*HelpdeskLocaleSection, *Response, error)`
  * **Save Helpdesk Locale Section** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-locale-section)
    * `client.Website.SaveHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error)`
  * **Update Helpdesk Locale Section** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-helpdesk-locale-section)
    * `client.Website.UpdateHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string, section HelpdeskLocaleSection) (*Response, error)`
  * **Delete Helpdesk Locale Section** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-locale-section)
    * `client.Website.DeleteHelpdeskLocaleSection(websiteID string, locale string, categoryId string, sectionId string) (*Response, error)`
  * **Map Helpdesk Locale Feedback Ratings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#map-helpdesk-locale-feedback-ratings)
    * `client.Website.MapHelpdeskLocaleFeedbackRatings(websiteID string, locale string, filterDateStart string, filterDateEnd string) (*HelpdeskLocaleFeedbackRatings, *Response, error)`
  * **List Helpdesk Locale Feedbacks** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-locale-feedbacks)
    * `client.Website.ListHelpdeskLocaleFeedbacks(websiteID string, locale string, pageNumber uint, filterDateStart string, filterDateEnd string) (*[]HelpdeskLocaleFeedbackItem, *Response, error)`
  * **Import External Helpdesk To Locale** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#import-external-helpdesk-to-locale)
    * `client.Website.ImportExternalHelpdeskToLocale(websiteID string, locale string, helpdeskUrl string) (*Response, error)`
  * **Export Helpdesk Locale Articles** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#export-helpdesk-locale-articles)
    * `client.Website.ExportHelpdeskLocaleArticles(websiteID string, locale string) (*Response, error)`
  * **List Helpdesk Redirections** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-helpdesk-redirections)
    * `client.Website.ListHelpdeskRedirections(websiteID string, pageNumber uint) (*[]HelpdeskRedirection, *Response, error)`
  * **Add Helpdesk Redirection** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#add-helpdesk-redirection)
    * `client.Website.AddHelpdeskRedirection(websiteID string, redirectionPath string, redirectionTarget string) (*HelpdeskRedirectionNew, *Response, error)`
  * **Check If Helpdesk Redirection Exists** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-if-helpdesk-redirection-exists)
    * `client.Website.CheckHelpdeskRedirectionExists(websiteID string, redirectionId string) (*Response, error)`
  * **Resolve Helpdesk Redirection** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-redirection)
    * `client.Website.ResolveHelpdeskRedirection(websiteID string, redirectionId string) (*HelpdeskRedirection, *Response, error)`
  * **Delete Helpdesk Redirection** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#delete-helpdesk-redirection)
    * `client.Website.DeleteHelpdeskRedirection(websiteID string, redirectionId string) (*Response, error)`
  * **Resolve Helpdesk Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-settings)
    * `client.Website.ResolveHelpdeskSettings(websiteID string) (*HelpdeskSettings, *Response, error)`
  * **Save Helpdesk Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-helpdesk-settings)
    * `client.Website.SaveHelpdeskSettings(websiteID string, settings HelpdeskSettings) (*Response, error)`
  * **Resolve Helpdesk Domain** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#resolve-helpdesk-domain)
    * `client.Website.ResolveHelpdeskDomain(websiteID string) (*HelpdeskDomain, *Response, error)`
  * **Request Helpdesk Domain Change** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#request-helpdesk-domain-change)
    * `client.Website.RequestHelpdeskDomainChange(websiteID string, basic *string, custom *string) (*Response, error)`
  * **Generate Helpdesk Domain Setup Flow** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#generate-helpdesk-domain-setup-flow)
    * `client.Website.GenerateHelpdeskDomainSetupFlow(websiteID string, custom string) (*HelpdeskDomainSetupFlow, *Response, error)`

### Plugin

* #### **Plugin Subscription**
  * **List All Active Subscriptions** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-active-subscriptions)
    * `client.Plugin.ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error)`
  * **List Subscriptions For A Website** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-subscriptions-for-a-website)
    * `client.Plugin.ListSubscriptionsForWebsite(websiteID string) (*[]PluginSubscription, *Response, error)`
  * **Get Subscription Details** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-subscription-details)
    * `client.Plugin.GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error)`
  * **Subscribe Website To Plugin** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#subscribe-website-to-plugin)
    * `client.Plugin.SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error)`
  * **Unsubscribe Plugin From Website** [`user`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#unsubscribe-plugin-from-website)
    * `client.Plugin.UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error)`
  * **Get Subscription Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-subscription-settings)
    * `client.Plugin.GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error)`
  * **Save Subscription Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#save-subscription-settings)
    * `client.Plugin.SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Update Subscription Settings** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#update-subscription-settings)
    * `client.Plugin.UpdateSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Forward Plugin Payload To Channel** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#forward-plugin-payload-to-channel)
    * `client.Plugin.ForwardPluginPayloadToChannel(websiteID string, pluginID string, payload PluginSubscriptionChannelForward) (*Response, error)`
  * **Dispatch Plugin Event** [`user`, `plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#dispatch-plugin-event)
    * `client.Plugin.DispatchPluginEvent(websiteID string, pluginID string, payload PluginSubscriptionEventDispatch) (*Response, error)`

* #### **Plugin Connect**
  * **⭐ Get Connect Account** [`plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-connect-account)
    * `client.Plugin.GetConnectAccount() (*PluginConnectAccount, *Response, error)`
  * **⭐ Check Connect Session Validity** [`plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#check-connect-session-validity)
    * `client.Plugin.CheckConnectSessionValidity() (*Response, error)`
  * **⭐ List All Connect Websites** [`plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-connect-websites)
    * `client.Plugin.ListAllConnectWebsites(pageNumber uint, filterConfigured bool) (*[]PluginConnectAllWebsites, *Response, error)`
  * **⭐ List Connect Websites Since** [`plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#list-all-connect-websites)
    * `client.Plugin.ListConnectWebsitesSince(dateSince time.Time, filterConfigured bool) (*[]PluginConnectWebsitesSince, *Response, error)`
  * **⭐ Get Connect Endpoints** [`plugin`]: [Reference](https://docs.crisp.chat/references/rest-api/v1/#get-connect-endpoints)
    * `client.Plugin.GetConnectEndpoints() (*PluginConnectEndpoints, *Response, error)`

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
  * **Session Update Availability** [`user`, `plugin`]:
    * `session:update_availability`
  * **Session Update Verify** [`user`, `plugin`]:
    * `session:update_verify`
  * **Session Request Initiated** [`user`, `plugin`]:
    * `session:request:initiated`
  * **Session Set Email** [`user`, `plugin`]:
    * `session:set_email`
  * **Session Set Phone** [`user`, `plugin`]:
    * `session:set_phone`
  * **Session Set Address** [`user`, `plugin`]:
    * `session:set_address`
  * **Session Set Subject** [`user`, `plugin`]:
    * `session:set_subject`
  * **Session Set Avatar** [`user`, `plugin`]:
    * `session:set_avatar`
  * **Session Set Nickname** [`user`, `plugin`]:
    * `session:set_nickname`
  * **Session Set Data** [`user`, `plugin`]:
    * `session:set_data`
  * **Session Sync Pages** [`user`, `plugin`]:
    * `session:sync:pages`
  * **Session Sync Events** [`user`, `plugin`]:
    * `session:sync:events`
  * **Session Sync Capabilities** [`user`, `plugin`]:
    * `session:sync:capabilities`
  * **Session Sync Geolocation** [`user`, `plugin`]:
    * `session:sync:geolocation`
  * **Session Sync System** [`user`, `plugin`]:
    * `session:sync:system`
  * **Session Sync Network** [`user`, `plugin`]:
    * `session:sync:network`
  * **Session Sync Timezone** [`user`, `plugin`]:
    * `session:sync:timezone`
  * **Session Sync Locales** [`user`, `plugin`]:
    * `session:sync:locales`
  * **Session Sync Rating** [`user`, `plugin`]:
    * `session:sync:rating`
  * **Session Sync Topic** [`user`, `plugin`]:
    * `session:sync:topic`
  * **Session Set State** [`user`, `plugin`]:
    * `session:set_state`
  * **Session Set Block** [`user`, `plugin`]:
    * `session:set_block`
  * **Session Set Segments** [`user`, `plugin`]:
    * `session:set_segments`
  * **Session Set Opened** [`user`, `plugin`]:
    * `session:set_opened`
  * **Session Set Closed** [`user`, `plugin`]:
    * `session:set_closed`
  * **Session Set Participants** [`user`, `plugin`]:
    * `session:set_participants`
  * **Session Set Mentions** [`user`, `plugin`]:
    * `session:set_mentions`
  * **Session Set Routing** [`user`, `plugin`]:
    * `session:set_routing`
  * **Session Removed** [`user`, `plugin`]:
    * `session:removed`

* #### **Message Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#message-events)
  * **Message Updated** [`user`, `plugin`]:
    * `message:updated`
  * **Message Send (Text Variant)** [`user`, `plugin`]:
    * `message:send/text`
  * **Message Send (File Variant)** [`user`, `plugin`]:
    * `message:send/file`
  * **Message Send (Animation Variant)** [`user`, `plugin`]:
    * `message:send/animation`
  * **Message Send (Audio Variant)** [`user`, `plugin`]:
    * `message:send/audio`
  * **Message Send (Picker Variant)** [`user`, `plugin`]:
    * `message:send/picker`
  * **Message Send (Field Variant)** [`user`, `plugin`]:
    * `message:send/field`
  * **Message Send (Carousel Variant)** [`user`, `plugin`]:
    * `message:send/carousel`
  * **Message Send (Note Variant)** [`user`, `plugin`]:
    * `message:send/note`
  * **Message Send (Event Variant)** [`user`, `plugin`]:
    * `message:send/event`
  * **Message Received (Text Variant)** [`user`, `plugin`]:
    * `message:received/text`
  * **Message Received (File Variant)** [`user`, `plugin`]:
    * `message:received/file`
  * **Message Received (Animation Variant)** [`user`, `plugin`]:
    * `message:received/animation`
  * **Message Received (Audio Variant)** [`user`, `plugin`]:
    * `message:received/audio`
  * **Message Received (Picker Variant)** [`user`, `plugin`]:
    * `message:received/picker`
  * **Message Received (Field Variant)** [`user`, `plugin`]:
    * `message:received/field`
  * **Message Received (Carousel Variant)** [`user`, `plugin`]:
    * `message:received/carousel`
  * **Message Received (Note Variant)** [`user`, `plugin`]:
    * `message:received/note`
  * **Message Received (Event Variant)** [`user`, `plugin`]:
    * `message:received/event`
  * **Message Removed** [`user`, `plugin`]:
    * `message:removed`
  * **Message Compose Send** [`user`, `plugin`]:
    * `message:compose:send`
  * **Message Compose Receive** [`user`, `plugin`]:
    * `message:compose:receive`
  * **Message Acknowledge Read Send** [`user`, `plugin`]:
    * `message:acknowledge:read:send`
  * **Message Acknowledge Read Received** [`user`, `plugin`]:
    * `message:acknowledge:read:received`
  * **Message Acknowledge Delivered** [`user`, `plugin`]:
    * `message:acknowledge:delivered`
  * **Message Notify Unread Send** [`user`, `plugin`]:
    * `message:notify:unread:send`
  * **Message Notify Unread Received** [`user`, `plugin`]:
    * `message:notify:unread:received`

* #### **People Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#people-events)
  * **People Profile Created** [`user`, `plugin`]:
    * `people:profile:created`
  * **People Profile Updated** [`user`, `plugin`]:
    * `people:profile:updated`
  * **People Profile Removed** [`user`, `plugin`]:
    * `people:profile:removed`
  * **People Bind Session** [`user`, `plugin`]:
    * `people:bind:session`
  * **People Sync Profile** [`user`, `plugin`]:
    * `people:sync:profile`
  * **People Import Progress** [`user`]:
    * `people:import:progress`
  * **People Import Done** [`user`]:
    * `people:import:done`

* #### **Campaign Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#campaign-events)
  * **Campaign Progress** [`user`]:
    * `campaign:progress`
  * **Campaign Dispatched** [`user`]:
    * `campaign:dispatched`
  * **Campaign Running** [`user`]:
    * `campaign:running`

* #### **Browsing Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#browsing-events)
  * **Browsing Request Initiated** [`user`, `plugin`]:
    * `browsing:request:initiated`
  * **Browsing Request Rejected** [`user`, `plugin`]:
    * `browsing:request:rejected`

* #### **Call Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#call-events)
  * **Call Request Initiated** [`user`, `plugin`]:
    * `call:request:initiated`
  * **Call Request Rejected** [`user`, `plugin`]:
    * `call:request:rejected`

* #### **Widget Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#widget-events)
  * **Widget Action Processed** [`user`]:
    * `widget:action:processed`

* #### **Status Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#status-events)
  * **Status Health Changed** [`user`]:
    * `status:health:changed`

* #### **Website Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#website-events)
  * **Website Update Visitors Count** [`user`, `plugin`]:
    * `website:update_visitors_count`
  * **Website Update Operators Availability** [`user`, `plugin`]:
    * `website:update_operators_availability`
  * **Website Users Available** [`user`, `plugin`]:
    * `website:users:available`

* #### **Bucket Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#bucket-events)
  * **Bucket URL Upload Generated** [`user`, `plugin`]:
    * `bucket:url:upload:generated`
  * **Bucket URL Avatar Generated** [`user`, `plugin`]:
    * `bucket:url:avatar:generated`
  * **Bucket URL Website Generated** [`user`, `plugin`]:
    * `bucket:url:website:generated`
  * **Bucket URL Campaign Generated** [`user`, `plugin`]:
    * `bucket:url:campaign:generated`
  * **Bucket URL Helpdesk Generated** [`user`, `plugin`]:
    * `bucket:url:helpdesk:generated`
  * **Bucket URL Status Generated** [`user`, `plugin`]:
    * `bucket:url:status:generated`
  * **Bucket URL Processing Generated** [`user`, `plugin`]:
    * `bucket:url:processing:generated`

* #### **Media Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#media-events)
  * **Media Animation Listed** [`user`]:
    * `media:animation:listed`

* #### **Email Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#email-events)
  * **Email Subscribe** [`user`, `plugin`]:
    * `email:subscribe`
  * **Email Track View** [`user`, `plugin`]:
    * `email:track:view`

* #### **Plugin Events**: [Reference](https://docs.crisp.chat/references/rtm-api/v1/#plugin-events)
  * **Plugin Channel** [`user`, `plugin`]:
    * `plugin:channel`
  * **Plugin Event** [`user`, `plugin`]:
    * `plugin:event`
  * **Plugin Settings Saved** [`user`, `plugin`]:
    * `plugin:settings:saved`
