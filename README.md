# go-crisp-api

[![Build Status](https://travis-ci.org/crisp-im/go-crisp-api.svg?branch=master)](https://travis-ci.org/crisp-im/go-crisp-api)

The Crisp API Golang wrapper. Authenticate, send messages, fetch conversations, access your agent accounts from your Go code.

Copyright 2018 Crisp IM SARL. See LICENSE for copying information.

* **📝 Implements**: [Crisp Platform - API ~ v1](https://docs.crisp.chat/api/v1/) at reference revision: 05/14/2018
* **😘 Maintainer**: [@valeriansaliou](https://github.com/valeriansaliou)

## Usage

```go
import "github.com/crisp-im/go-crisp-api/crisp"
```

Construct a new Crisp client, then use the various services on the client to
access different parts of the Crisp API. For example:

```go
client := crisp.New()

// Get plugin information
plugin, _, err := client.Plugin.GetPluginInformation("185fe7ee-7cc6-4b8b-884d-fda9df632c13")
```

## Authentication

To authenticate against the API, generate your session identifier and session key **once** using the following cURL request in your terminal (replace `YOUR_ACCOUNT_EMAIL` and `YOUR_ACCOUNT_PASSWORD`):

```bash
curl -H "Content-Type: application/json" -X POST -d '{"email":"YOUR_ACCOUNT_EMAIL","password":"YOUR_ACCOUNT_PASSWORD"}' https://api.crisp.chat/v1/user/session/login
```

If authentication succeeds, you will get a JSON response containing your authentication keys: `identifier` and `key`. **Keep those 2 values private, and store them safely for long-term use**.

Then, add authentication parameters to your `client` instance right after you create it:

```go
client := crisp.New()

// Authenticate to API (identifier, key)
// eg. client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")
client.Authenticate(identifier, key)

// Now, you can use authenticated API sections.
```

**🔴 Important: Be sure to login once, and re-use the same authentication keys (same `identifier` + `key`) in all your subsequent requests to the API. Do not generate new tokens from your code for every new request to the API (you will be heavily rate-limited; that will induce HTTP failures for some of your API calls).**

## Resource Methods

All the available Crisp API resources are fully implemented. **Programmatic methods names are named after their label name in the [API Reference](https://docs.crisp.chat/api/v1/)**.

Thus, it is straightforward to look for them in the library while reading the [API Reference](https://docs.crisp.chat/api/v1/).

In the following method prototypes, `crisp` is to be replaced with your Crisp API instance. For example, instanciate `client := crisp.New()` and then call eg: `client.User.CheckSessionValidity()`.

When calling a method that writes data to the API (eg: `client.User.CreateUserAccount()`), you need to build an account instance and submit it this way:

```go
userAccount := crisp.UserAccountCreate{
  Email: "john@acme-inc.com",
  Password: "SecurePassword",
  FirstName: "John",
  LastName: "Doe",
}

client.User.CreateUserAccount(userAccount)
```

Refer directly to [the library source code](https://github.com/crisp-im/go-crisp-api/tree/master/crisp) to know more about those structures.

### Email

* **Email Subscription**
  * **Get Subscription Status**: `client.Email.GetSubscriptionStatus(emailHash string, key string, websiteID string) (*SubscriptionStatus, *Response, error)`
  * **Update Subscription Status**: `client.Email.UpdateSubscriptionStatus(emailHash string, key string, websiteID string, subscribed bool) (*Response, error)`

### Bucket

* **Bucket URL**
  * **Generate Bucket URL**: `client.Bucket.GenerateBucketURL(bucketData BucketURLRequest) (*Response, error)`

### Media

* **Media Animation**
  * **List Animation Medias**: `client.Media.ListAnimationMedias(pageNumber uint, listID int, searchQuery string) (*Response, error)`

### User

* **User Availability**
  * **Get User Availability**: `client.User.GetUserAvailability() (*UserAvailability, *Response, error)`
  * **Update User Availability**: `client.User.UpdateUserAvailability(availabilityType string, timeFor uint) (*Response, error)`
  * **Get User Availability Status**: `client.User.GetUserAvailabilityStatus() (*UserAvailabilityStatus, *Response, error)`

* **User Account Base**
  * **Get User Account**: `client.User.GetUserAccount() (*UserAccount, *Response, error)`
  * **Create User Account**: `client.User.CreateUserAccount(user UserAccountCreate) (*Response, error)`
  * **Delete User Account**: `client.User.DeleteUserAccount() (*Response, error)`

* **User Account Billing**
  * **List All Billing Methods**: `client.User.ListAllBillingMethods() (*[]BillingMethodAll, *Response, error)`
  * **Add New Billing Method**: `client.User.AddNewBillingMethod(billing BillingMethodCreate) (*Response, error)`
  * **Get A Billing Method**: `client.User.GetBillingMethod(cardID string) (*BillingMethod, *Response, error)`
  * **Update A Billing Method**: `client.User.UpdateBillingMethod(cardID string, billing BillingMethodUpdate) (*Response, error)`
  * **Remove A Billing Method**: `client.User.RemoveBillingMethod(cardID string) (*Response, error)`
  * **List Invoices For Billing Method**: `client.User.ListInvoicesForBillingMethod(cardID string, pageNumber int) (*BillingMethodInvoiceAll, *Response, error)`
  * **Get An Invoice For Billing Method**: `client.User.GetInvoiceForBillingMethod(cardID string, invoiceID string) (*BillingMethodInvoice, *Response, error)`
  * **Link To An External Billing Method**: `client.User.LinkToExternalBillingMethod(billingService string) (*Response, error)`
  * **Finish Linking External Billing Method**: `client.User.FinishLinkingExternalBillingMethod(cardID string) (*Response, error)`
  * **Finish Unlinking External Billing Method**: `client.User.FinishUnlinkingExternalBillingMethod(cardID string) (*Response, error)`

* **User Account Notification**
  * **Get Notification Settings**: `client.User.GetNotificationSettings() (*UserNotificationSettings, *Response, error)`
  * **Update Notification Settings**: `client.User.UpdateNotificationSettings(notifications UserNotificationSettingsUpdate) (*Response, error)`
  * **Add A Notification Provider**: `client.User.AddNotificationProvider(notificationID string) (*Response, error)`
  * **Delete A Notification Provider**: `client.User.DeleteNotificationProvider(notificationID string) (*Response, error)`

* **User Account Websites**
  * **List Websites**: `client.User.ListWebsites() (*[]UserWebsite, *Response, error)`

* **User Account Profile**
  * **Get Profile**: `client.User.GetProfile() (*UserProfile, *Response, error)`
  * **Update Profile**: `client.User.UpdateProfile(profile UserProfileUpdate) (*Response, error)`

* **User Account Token**
  * **Check Account Token Configured**: `client.User.CheckAccountTokenConfigured() (*Response, error)`
  * **Configure Account Token**: `client.User.ConfigureAccountToken(secret string) (*Response, error)`
  * **Unconfigure Account Token**: `client.User.UnconfigureAccountToken() (*Response, error)`
  * **Generate Account Token**: `client.User.GenerateAccountToken() (*UserTokenGenerate, *Response, error)`
  * **Verify Account Token**: `client.User.VerifyAccountToken(secret string, token string) (*Response, error)`

* **User Account Recover**
  * **Get Recovery Details**: `client.User.GetRecoveryDetails(userID string, recoverIdentifier string, recoverKey string) (*Response, error)`
  * **Send Recovery Password**: `client.User.SendRecoveryPassword(userID string, recoverIdentifier string, recoverKey string, password string) (*Response, error)`
  * **Delete Recovery Keypair**: `client.User.DeleteRecoveryKeypair(userID string, recoverIdentifier string, recoverKey string) (*Response, error)`

* **User Account Schedule**
  * **Get Schedule Settings**: `client.User.GetScheduleSettings() (*UserScheduleSettings, *Response, error)`
  * **Update Schedule Settings**: `client.User.UpdateScheduleSettings(schedule UserScheduleSettingsUpdate) (*Response, error)`

* **User Session**
  * **Check Session Validity**: `client.User.CheckSessionValidity() (*Response, error)`
  * **Create A New Session**: `client.User.CreateNewSession(email string, password string) (*UserSessionParameters, *Response, error)`
  * **Create A New Session (Token Variant)**: `client.User.CreateNewSessionWithToken(email string, password string, token string) (*UserSessionParameters, *Response, error)`
  * **Destroy A Session**: `client.User.DestroySession() (*Response, error)`
  * **Recover A Session**: `client.User.RecoverSession(email string)`

* **User Statistics**
  * **Count Total Unread Messages**: `client.User.CountTotalUnreadMessages() (*UserStatistics, *Response, error)`

### Website

* **Website Base**
  * **Check If Website Exists**: `client.Website.CheckWebsiteExists(domain string) (*Response, error)`
  * **Create Website**: `client.Website.CreateWebsite(websiteData WebsiteCreate) (*Website, *Response, error)`
  * **Get A Website**: `client.Website.GetWebsite(websiteID string) (*Website, *Response, error)`
  * **Delete A Website**: `client.Website.DeleteWebsite(websiteID string) (*Response, error)`

* **Website Batch**
  * **Batch Resolve Conversations**: `client.Website.BatchResolveConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Read Conversations**: `client.Website.BatchReadConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Remove Conversations**: `client.Website.BatchRemoveConversations(websiteID string, sessions []string) (*Response, error)`
  * **Batch Remove People**: `client.Website.BatchRemovePeople(websiteID string, people WebsiteBatchPeopleOperationInner) (*Response, error)`

* **Website Billing**
  * **Get Website Billing**: `client.Website.GetWebsiteBilling(websiteID string) (*WebsiteBilling, *Response, error)`
  * **Update Website Billing**: `client.Website.UpdateWebsiteBilling(websiteID string, cardID string) (*Response, error)`
  * **Unlink Website Billing**: `client.Website.UnlinkWebsiteBilling(websiteID string) (*Response, error)`

* **Website Availability**
  * **Get Website Availability Status**: `client.Website.GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error)`

* **Website Channel**
  * **Get Website Email Channel**: `client.Website.GetWebsiteEmailChannel(websiteID string) (*WebsiteChannelEmail, *Response, error)`
  * **Get Website Email Channel Domain**: `client.Website.GetWebsiteEmailChannelDomain(websiteID string) (*WebsiteChannelEmailDomain, *Response, error)`
  * **Request Website Email Channel Domain Change**: `client.Website.RequestWebsiteEmailChannelDomainChange(websiteID string, domain WebsiteChannelEmailDomainChange) (*Response, error)`
  * **Generate Website Email Channel Setup Flow**: `client.Website.GenerateWebsiteEmailChannelSetupFlow(websiteID string, domain string) (*WebsiteChannelEmailSetupFlow, *Response, error)`
  * **Get Website Email Channel Relay**: `client.Website.GetWebsiteEmailChannelRelay(websiteID string) (*WebsiteChannelEmailRelay, *Response, error)`
  * **Request Website Email Channel Relay Change**: `client.Website.RequestWebsiteEmailChannelRelayChange(websiteID string, relay WebsiteChannelEmailRelayRequest) (*Response, error)`

* **Website Contract**
  * **List Website Contracts**: `client.Website.ListWebsiteContracts(websiteID string) (*[]WebsiteContract, *Response, error)`
  * **Check If Agreed Website Contract Exists**: `client.Website.CheckIfAgreedWebsiteContractExists(websiteID string, contractID string) (*Response, error)`
  * **Resolve Agreed Website Contract**: `client.Website.ResolveAgreedWebsiteContract(websiteID string, contractID string) (*WebsiteContract, *Response, error)`
  * **Agree To Website Contract**: `client.Website.AgreeToWebsiteContract(websiteID string, contractID string, agreementURL string) (*Response, error)`
  * **Delete Agreed Website Contract**: `client.Website.DeleteAgreedWebsiteContract(websiteID string, contractID string) (*Response, error)`

* **Website Operator**
  * **List Website Operators**: `client.Website.ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error)`
  * **List Last Active Website Operators**: `client.Website.ListLastActiveWebsiteOperators(websiteID string) (*[]WebsiteOperatorsLastActiveListOne, *Response, error)`
  * **Flush Last Active Website Operators**: `client.Website.FlushLastActiveWebsiteOperators(websiteID string) (*Response, error)`
  * **Send Email To Website Operators**: `client.Website.SendEmailToWebsiteOperators(websiteID string, email WebsiteOperatorEmail) (*Response, error)`
  * **Get A Website Operator**: `client.Website.GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error)`
  * **Invite A Website Operator**: `client.Website.InviteWebsiteOperator(websiteID string, email string, role string) (*Response, error)`
  * **Change Operator Role**: `client.Website.ChangeOperatorRole(websiteID string, userID string, role string) (*Response, error)`
  * **Unlink Operator From Website**: `client.Website.UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error)`

* **Website Invite**
  * **Get Invite Details**: `client.Website.GetInviteDetails(websiteID string, inviteIdentifier string, inviteKey string) (*WebsiteInvite, *Response, error)`
  * **Redeem Invite**: `client.Website.RedeemInvite(websiteID string, inviteIdentifier string, inviteKey string) (*Response, error)`
  * **Delete Invite Keypair**: `client.Website.DeleteInviteKeypair(websiteID string, inviteIdentifier string, inviteKey string) (*Response, error)`

* **Website Rating**
  * **Resolve Session Rating**: `client.Website.ResolveSessionRating(websiteID string, sessionID string) (*WebsiteRatingSession, *Response, error)`
  * **Submit Session Rating**: `client.Website.SubmitSessionRating(websiteID string, sessionID string, websiteRatingSession WebsiteRatingCreateSession) (*Response, error)`
  * **Delete Session Rating**: `client.Website.DeleteSessionRating(websiteID string, sessionID string) (*Response, error)`

* **Website Routing**
  * **GetWebsite Routing Settings**: `client.Website.GetWebsiteRoutingSettings(websiteID string) (*WebsiteRoutingSettings, *Response, error)`
  * **Update Website Routing Settings**: `client.Website.UpdateWebsiteRoutingSettings(websiteID string, settings WebsiteRoutingSettingsUpdate) (*Response, error)`
  * **List Website Routing Rules**: `client.Website.ListWebsiteRoutingRules(websiteID string) (*WebsiteRoutingRules, *Response, error)`
  * **Save Website Routing Rules**: `client.Website.SaveWebsiteRoutingRules(websiteID string, rules WebsiteRoutingRulesUpdate) (*Response, error)`

* **Website Settings**
  * **Get Website Settings**: `client.Website.GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error)`
  * **Update Website Settings**: `client.Website.UpdateWebsiteSettings(websiteID string, settings WebsiteSettingsUpdate) (*Response, error)`

* **Website Visitors**
  * **Count Visitors**: `client.Website.CountVisitors(websiteID string) (*WebsiteVisitorCount, *Response, error)`
  * **List Visitors**: `client.Website.ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error)`
  * **Pinpoint Visitors On A Map (Wide Variant)**: `client.Website.PinpointVisitorsOnMapWide(websiteID string) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Pinpoint Visitors On A Map (Area Variant)**: `client.Website.PinpointVisitorsOnMapArea(websiteID string, centerLongitude float32, centerLatitude float32, centerRadius uint) (*[]WebsiteVisitorsMapPoint, *Response, error)`
  * **Count Blocked Visitors**: `client.Website.CountBlockedVisitors(websiteID string) (*[]WebsiteVisitorsBlocked, *Response, error)`
  * **Count Blocked Visitors In Rule**: `client.Website.CountBlockedVisitorsInRule(websiteID string, rule string) (*WebsiteVisitorsBlocked, *Response, error)`
  * **Clear Blocked Visitors In Rule**: `client.Website.ClearBlockedVisitorsInRule(websiteID string, rule string) (*Response, error)`

* **Website States**
  * **Request Website States**: `client.Website.RequestWebsiteStates(websiteID string) (*Response, error)`

* **Website Hooks**
  * **List Hooks**: `client.Website.ListHooks(websiteID string, pageNumber uint) (*[]WebsiteHook, *Response, error)`

* **Website Hook**
  * **Create A New Hook**: `client.Website.CreateNewHook(websiteID string, websiteHookItem WebsiteHookItem) (*Response, error)`
  * **Check If Hook Exists**: `client.Website.CheckHookExists(websiteID string, hookID string) (*Response, error)`
  * **Get A Hook**: `client.Website.GetHook(websiteID string, hookID string) (*WebsiteHook, *Response, error)`
  * **Save A Hook**: `client.Website.SaveHook(websiteID string, hookID string, websiteHookItem WebsiteHookItem) (*Response, error)`
  * **Update A Hook**: `client.Website.UpdateHook(websiteID string, hookID string, websiteHookItem WebsiteHookItem) (*Response, error)`
  * **Remove A Hook**: `client.Website.RemoveHook(websiteID string, hookID string) (*Response, error)`

* **Website Verify**
  * **Get Verify Settings**: `client.Website.GetVerifySettings(websiteID string) (*WebsiteVerifySettings, *Response, error)`
  * **Update Verify Settings**: `client.Website.UpdateVerifySettings(websiteID string, verify WebsiteVerifySettingsUpdate) (*Response, error)`
  * **Get Verify Key**: `client.Website.GetVerifyKey(websiteID string) (*WebsiteVerifyKey, *Response, error)`
  * **Roll Verify Key**: `client.Website.RollVerifyKey(websiteID string) (*Response, error)`

* **Website Shortcuts**
  * **List Shortcuts**: `client.Website.ListShortcuts(websiteID string, pageNumber uint) (*[]WebsiteShortcut, *Response, error)`
  * **List Shortcuts (Search Variant)**: `client.Website.SearchShortcuts(websiteID string, pageNumber uint, searchQuery string, searchTag string) (*[]WebsiteShortcut, *Response, error)`
  * **List Shortcut Tags**: `client.Website.ListShortcutTags(websiteID string) (*[]string, *Response, error)`

* **Website Shortcut**
  * **Create A New Shortcut**: `client.Website.CreateNewShortcut(websiteID string, websiteShortcutItem WebsiteShortcutItem) (*Response, error)`
  * **Check If Shortcut Exists**: `client.Website.CheckShortcutExists(websiteID string, shortcutID string) (*Response, error)`
  * **Get A Shortcut**: `client.Website.GetShortcut(websiteID string, shortcutID string) (*WebsiteShortcut, *Response, error)`
  * **Save A Shortcut**: `client.Website.SaveShortcut(websiteID string, shortcutID string, websiteShortcutItem WebsiteShortcutItem) (*Response, error)`
  * **Update A Shortcut**: `client.Website.UpdateShortcut(websiteID string, shortcutID string, websiteShortcutItem WebsiteShortcutItem) (*Response, error)`
  * **Remove A Shortcut**: `client.Website.RemoveShortcut(websiteID string, shortcutID string) (*Response, error)`

* **Website Campaigns**
  * **List Campaigns**: `client.Website.ListCampaigns(websiteID string, pageNumber uint) (*[]WebsiteCampaignExcerpt, *Response, error)`
  * **List Campaigns (Filter Variant)**: `client.Website.FilterCampaigns(websiteID string, pageNumber uint, searchName string, filterTypeOneShot bool, filterTypeAutomated bool, filterStatusNotConfigured bool, filterStatusReady bool, filterStatusPaused bool, filterStatusSending bool, filterStatusDone bool) (*[]WebsiteCampaignExcerpt, *Response, error)`

* **Website Campaign**
  * **Create A New Campaign**: `client.Website.CreateNewCampaign(websiteID string, campaignType string, campaignName string) (*WebsiteCampaignNew, *Response, error)`
  * **Check If Campaign Exists**: `client.Website.CheckCampaignExists(websiteID string, campaignID string) (*Response, error)`
  * **Get A Campaign**: `client.Website.GetCampaign(websiteID string, campaignID string) (*WebsiteCampaignItem, *Response, error)`
  * **Save A Campaign**: `client.Website.SaveCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Update A Campaign**: `client.Website.UpdateCampaign(websiteID string, campaignID string, websiteCampaignItem WebsiteCampaignItem) (*Response, error)`
  * **Remove A Campaign**: `client.Website.RemoveCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Dispatch A Campaign**: `client.Website.DispatchCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Resume A Campaign**: `client.Website.ResumeCampaign(websiteID string, campaignID string) (*Response, error)`
  * **Pause A Campaign**: `client.Website.PauseCampaign(websiteID string, campaignID string) (*Response, error)`
  * **List Campaign Statistics**: `client.Website.ListCampaignStatistics(websiteID string, campaignID string, action string, pageNumber uint) (*[]WebsiteCampaignStatistic, *Response, error)`

* **Website Conversations**
  * **List Conversations**: `client.Website.ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error)`
  * **List Conversations (Search Variant)**: `client.Website.SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error)`
  * **List Conversation Segments In Meta**: `client.Website.ListConversationSegmentsInMeta(websiteID string, pageNumber uint) (*[]ConversationMetaSegment, *Response, error)`

* **Website Conversation**
  * **Create A New Conversation**: `client.Website.CreateNewConversation(websiteID string) (*ConversationNew, *Response, error)`
  * **Check If Conversation Exists**: `client.Website.CheckConversationExists(websiteID string, sessionID string) (*Response, error)`
  * **Get A Conversation**: `client.Website.GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error)`
  * **Remove A Conversation**: `client.Website.RemoveConversation(websiteID string, sessionID string) (*Response, error)`
  * **Initiate A Conversation With Existing Session**: `client.Website.InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error)`
  * **Get Messages In Conversation (Last Variant)**: `client.Website.GetMessagesInConversationLast(websiteID string, sessionID string) (*[]ConversationMessage, *Response, error)`
  * **Get Messages In Conversation (Before Variant)**: `client.Website.GetMessagesInConversationBefore(websiteID string, sessionID string, timestampBefore uint32) (*[]ConversationMessage, *Response, error)`
  * **Send A Message In Conversation (Text Variant)**: `client.Website.SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (File Variant)**: `client.Website.SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Animation Variant)**: `client.Website.SendAnimationMessageInConversation(websiteID string, sessionID string, message ConversationAnimationMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Audio Variant)**: `client.Website.SendAudioMessageInConversation(websiteID string, sessionID string, message ConversationAudioMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Picker Variant)**: `client.Website.SendPickerMessageInConversation(websiteID string, sessionID string, message ConversationPickerMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Field Variant)**: `client.Website.SendFieldMessageInConversation(websiteID string, sessionID string, message ConversationFieldMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Send A Message In Conversation (Note Variant)**: `client.Website.SendNoteMessageInConversation(websiteID string, sessionID string, message ConversationNoteMessageNew) (*ConversationMessageDispatched, *Response, error)`
  * **Get A Message In Conversation**: `client.Website.GetMessageInConversation(websiteID string, sessionID string, fingerprint int) (*ConversationMessage, *Response, error)`
  * **Update A Message In Conversation (Text Variant)**: `client.Website.UpdateTextMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Update A Message In Conversation (File Variant)**: `client.Website.UpdateFileMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFileMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Animation Variant)**: `client.Website.UpdateAnimationMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAnimationMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Audio Variant)**: `client.Website.UpdateAudioMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationAudioMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Picker Variant)**: `client.Website.UpdatePickerMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationPickerMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Field Variant)**: `client.Website.UpdateFieldMessageInConversation(websiteID string, sessionID string, fingerprint int, content ConversationFieldMessageNewContent) (*Response, error)`
  * **Update A Message In Conversation (Note Variant)**: `client.Website.UpdateNoteMessageInConversation(websiteID string, sessionID string, fingerprint int, content string) (*Response, error)`
  * **Compose A Message In Conversation**: `client.Website.ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error)`
  * **Mark Messages As Read In Conversation**: `client.Website.MarkMessagesReadInConversation(websiteID string, sessionID string, read ConversationReadMessageMark) (*Response, error)`
  * **Mark Messages As Delivered In Conversation**: `client.Website.MarkMessagesDeliveredInConversation(websiteID string, sessionID string, delivered ConversationDeliveredMessageMark) (*Response, error)`
  * **Update Conversation Open State**: `client.Website.UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error)`
  * **Get Conversation Routing Assign**: `client.Website.GetConversationRoutingAssign(websiteID string, sessionID string) (*ConversationRoutingAssign, *Response, error)`
  * **Assign Conversation Routing**: `client.Website.AssignConversationRouting(websiteID string, sessionID string, assign ConversationRoutingAssignUpdate) (*Response, error)`
  * **Request Translation Service**: `client.Website.RequestTranslationService(websiteID string, translate WebsiteServiceTranslateItem) (*Response, error)`
  * **Get Conversation Metas**: `client.Website.GetConversationMetas(websiteID string, sessionID string) (*ConversationMeta, *Response, error)`
  * **Update Conversation Metas**: `client.Website.UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error)`
  * **List Conversation Pages**: `client.Website.ListConversationPages(websiteID string, sessionID string, pageNumber uint) (*[]ConversationPage, *Response, error)`
  * **List Conversation Events**: `client.Website.ListConversationEvents(websiteID string, sessionID string, pageNumber uint) (*[]ConversationEvent, *Response, error)`
  * **Get Conversation State**: `client.Website.GetConversationState(websiteID string, sessionID string) (*ConversationState, *Response, error)`
  * **Change Conversation State**: `client.Website.ChangeConversationState(websiteID string, sessionID string, state string) (*Response, error)`
  * **Get Block Status For Conversation**: `client.Website.GetBlockStatusForConversation(websiteID string, sessionID string) (*ConversationBlock, *Response, error)`
  * **Block Incoming Messages For Conversation**: `client.Website.BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool) (*Response, error)`
  * **Request Email Transcript For Conversation**: `client.Website.RequestEmailTranscriptForConversation(websiteID string, sessionID string, to string, email string) (*Response, error)`
  * **Request Chatbox Binding Purge For Conversation**: `client.Website.RequestChatboxBindingPurgeForConversation(websiteID string, sessionID string) (*Response, error)`
  * **List Browsing Sessions For Conversation**: `client.Website.ListBrowsingSessionsForConversation(websiteID string, sessionID string) (*[]ConversationBrowsing, *Response, error)`
  * **Initiate Browsing Session For Conversation**: `client.Website.InitiateBrowsingSessionForConversation(websiteID string, sessionID string) (*Response, error)`
  * **Send Action To An Existing Browsing Session**: `client.Website.SendActionToExistingBrowsingSession(websiteID string, sessionID string, browsingID string, action string) (*Response, error)`
  * **Debug Existing Browsing Session**: `client.Website.DebugExistingBrowsingSession(websiteID string, sessionID string, browsingID string, debug ConversationBrowsingDebug) (*Response, error)`
  * **Assist Existing Browsing Session**: `client.Website.AssistExistingBrowsingSession(websiteID string, sessionID string, browsingID string, assist ConversationBrowsingAssist) (*Response, error)`
  * **Initiate New Call Session For Conversation**: `client.Website.InitiateNewCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Get Ongoing Call Session For Conversation**: `client.Website.GetOngoingCallSessionForConversation(websiteID string, sessionID string) (*ConversationCall, *Response, error)`
  * **Abort Ongoing Call Session For Conversation**: `client.Website.AbortOngoingCallSessionForConversation(websiteID string, sessionID string, callID string) (*Response, error)`
  * **Transmit Signaling On Ongoing Call Session**: `client.Website.TransmitSignalingOnOngoingCallSession(websiteID string, sessionID string, callID string, payload ConversationCallSignalingPayload) (*Response, error)`

* **Website Analytics**
  * **Acquire Chats Analytics**: `client.Website.AcquireChatsAnalytics(websiteID string, filterMetric string, filterOperator string, filterDateStart time.Time, filterDateEnd time.Time, filterDateSplit uint8) (*WebsiteAnalytics, *Response, error)`

* **Website People**
  * **Get People Statistics**: `client.Website.GetPeopleStatistics(websiteID string) (*PeopleStatistics, *Response, error)`
  * **List People Segments**: `client.Website.ListPeopleSegments(websiteID string, pageNumber uint) (*[]PeopleSegment, *Response, error)`
  * **List People Profiles**: `client.Website.ListPeopleProfiles(websiteID string, pageNumber uint, searchField string, searchOrder string, searchOperator string, searchFilter []PeopleFilter) (*[]PeopleProfile, *Response, error)`
  * **Add New People Profile**: `client.Website.AddNewPeopleProfile(websiteID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **Check If People Profile Exists**: `client.Website.CheckPeopleProfileExists(websiteID string, peopleID string) (*Response, error)`
  * **Get People Profile**: `client.Website.GetPeopleProfile(websiteID string, peopleID string) (*PeopleProfile, *Response, error)`
  * **Save People Profile**: `client.Website.SavePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **Update People Profile**: `client.Website.UpdatePeopleProfile(websiteID string, peopleID string, peopleProfile PeopleProfileUpdateCard) (*Response, error)`
  * **Remove People Profile**: `client.Website.RemovePeopleProfile(websiteID string, peopleID string) (*Response, error)`
  * **List People Conversations**: `client.Website.ListPeopleConversations(websiteID string, peopleID string, pageNumber uint) ([]string, *Response, error)`
  + **Add A People Event**: `client.Website.AddPeopleEvent(websiteID string, peopleID string, peopleEvent PeopleEventAdd) (*Response, error)`
  + **List People Events**: `client.Website.ListPeopleEvents(websiteID string, peopleID string, pageNumber uint) (*[]PeopleEvent, *Response, error)`
  + **Get People Data**: `client.Website.GetPeopleData(websiteID string, peopleID string) (*PeopleData, *Response, error)`
  + **Save People Data**: `client.Website.SavePeopleData(websiteID string, peopleID string, peopleData interface{}) (*Response, error)`
  + **Get People Subscription Status**: `client.Website.GetPeopleSubscriptionStatus(websiteID string, peopleID string) (*PeopleSubscription, *Response, error)`
  + **Update People Subscription Status**: `client.Website.UpdatePeopleSubscriptionStatus(websiteID string, peopleID string, peopleSubscription PeopleSubscriptionUpdate)`
  * **Export People Profiles**: `client.Website.ExportPeopleProfiles(websiteID string) (*Response, error)`
  * **Import People Profiles**: `client.Website.ImportPeopleProfiles(websiteID string, profileImportSetup PeopleProfileImportSetup) (*PeopleProfileImport, *Response, error)`

* **Website Helpdesk**
  * **Check If Helpdesk Exists**: `client.Website.CheckIfHelpdeskExists(websiteID string) (*Response, error)`
  * **Resolve Helpdesk**: `client.Website.ResolveHelpdesk(websiteID string) (*WebsiteHelpdesk, *Response, error)`
  * **Initialize Helpdesk**: `client.Website.InitializeHelpdesk(websiteID string, helpdesk WebsiteHelpdeskInitialize) (*Response, error)`
  * **Delete Helpdesk**: `client.Website.DeleteHelpdesk(websiteID string) (*Response, error)`
  * **List Helpdesk Locales**: `client.Website.ListHelpdeskLocales(websiteID string, pageNumber uint) (*[]WebsiteHelpdeskLocale, *Response, error)`
  * **Add Helpdesk Locale**: `client.Website.AddHelpdeskLocale(websiteID string, locale string) (*Response, error)`
  * **Check If Helpdesk Locale Exists**: `client.Website.CheckIfHelpdeskLocaleExists(websiteID string, locale string) (*Response, error)`
  * **Resolve Helpdesk Locale**: `client.Website.ResolveHelpdeskLocale(websiteID string, locale string) (*WebsiteHelpdeskLocale, *Response, error)`
  * **Delete Helpdesk Locale**: `client.Website.DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error)`
  * **List Helpdesk Locale Articles**: `client.Website.ListHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint) (*[]WebsiteHelpdeskLocaleArticleExcerpt, *Response, error)`
  * **List Helpdesk Locale Articles (Filter Variant)**: `client.Website.FilterHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint, searchTitle string, filterCategoryID string, filterStatusPublished bool, filterStatusDraft bool, filterVisibilityVisible bool, filterVisibilityHidden bool) (*[]WebsiteHelpdeskLocaleArticleExcerpt, *Response, error)`
  * **Add A New Helpdesk Locale Article**: `client.Website.AddNewHelpdeskLocaleArticle(websiteID string, locale string, articleTitle string) (*WebsiteHelpdeskLocaleArticleNew, *Response, error)`
  * **Check If Helpdesk Locale Article Exists**: `client.Website.CheckIfHelpdeskLocaleArticleExists(websiteID string, locale string, articleID string) (*Response, error)`
  * **Resolve Helpdesk Locale Article**: `client.Website.ResolveHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*WebsiteHelpdeskLocaleArticle, *Response, error)`
  * **Save Helpdesk Locale Article**: `client.Website.SaveHelpdeskLocaleArticle(websiteID string, locale string, articleID string, article WebsiteHelpdeskLocaleArticleUpdate) (*Response, error)`
  * **Update Helpdesk Locale Article**: `client.Website.UpdateHelpdeskLocaleArticle(websiteID string, locale string, articleID string, article WebsiteHelpdeskLocaleArticleUpdate) (*Response, error)`
  * **Delete Helpdesk Locale Article**: `client.Website.DeleteHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*Response, error)`
  * **Resolve Helpdesk Locale Article Category**: `client.Website.ResolveHelpdeskLocaleArticleCategory(websiteID string, locale string, articleID string) (*WebsiteHelpdeskLocaleArticleCategory, *Response, error)`
  * **Update Helpdesk Locale Article Category**: `client.Website.UpdateHelpdeskLocaleArticleCategory(websiteID string, locale string, articleID string, categoryID string) (*Response, error)`
  * **List Helpdesk Locale Article Alternates**: `client.Website.ListHelpdeskLocaleArticleAlternates(websiteID string, locale string, articleID string) (*[]WebsiteHelpdeskLocaleArticleAlternate, *Response, error)`
  * **Check If Helpdesk Locale Article Alternate Exists**: `client.Website.CheckIfHelpdeskLocaleArticleAlternateExists(websiteID string, locale string, articleID string, localeLinked string) (*Response, error)`
  * **Resolve Helpdesk Locale Article Alternate**: `client.Website.ResolveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleID string, localeLinked string) (*WebsiteHelpdeskLocaleArticleAlternate, *Response, error)`
  * **Save Helpdesk Locale Article Alternate**: `client.Website.SaveHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleID string, localeLinked string, articleIDAlternate string) (*Response, error)`
  * **Delete Helpdesk Locale Article Alternate**: `client.Website.DeleteHelpdeskLocaleArticleAlternate(websiteID string, locale string, articleID string, localeLinked string) (*Response, error)`
  * **Publish Helpdesk Locale Article**: `client.Website.PublishHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*WebsiteHelpdeskLocaleArticlePublish, *Response, error)`
  * **Unpublish Helpdesk Locale Article**: `client.Website.UnpublishHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*Response, error)`
  * **List Helpdesk Locale Categories**: `client.Website.ListHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint) (*[]WebsiteHelpdeskLocaleCategory, *Response, error)`
  * **List Helpdesk Locale Categories (Filter Variant)**: `client.Website.FilterHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint, searchName string) (*[]WebsiteHelpdeskLocaleCategory, *Response, error)`
  * **Add Helpdesk Locale Category**: `client.Website.AddHelpdeskLocaleCategory(websiteID string, locale string, categoryName string) (*WebsiteHelpdeskLocaleCategoryNew, *Response, error)`
  * **Check If Helpdesk Locale Category Exists**: `client.Website.CheckIfHelpdeskLocaleCategoryExists(websiteID string, locale string, categoryID string) (*Response, error)`
  * **Resolve Helpdesk Locale Category**: `client.Website.ResolveHelpdeskLocaleCategory(websiteID string, locale string, categoryID string) (*WebsiteHelpdeskLocaleCategory, *Response, error)`
  * **Save Helpdesk Locale Category**: `client.Website.SaveHelpdeskLocaleCategory(websiteID string, locale string, categoryID string, category WebsiteHelpdeskLocaleCategoryUpdate) (*Response, error)`
  * **Update Helpdesk Locale Category**: `client.Website.UpdateHelpdeskLocaleCategory(websiteID string, locale string, categoryID string, category WebsiteHelpdeskLocaleCategoryUpdate) (*Response, error)`
  * **Delete Helpdesk Locale Category**: `client.Website.DeleteHelpdeskLocaleCategory(websiteID string, locale string, categoryID string) (*Response, error)`
  * **Resolve Helpdesk Settings**: `client.Website.ResolveHelpdeskSettings(websiteID string) (*WebsiteHelpdeskSettings, *Response, error)`
  * **Save Helpdesk Settings**: `client.Website.SaveHelpdeskSettings(websiteID string, settings WebsiteHelpdeskSettingsUpdate) (*Response, error)`
  * **Resolve Helpdesk Domain**: `client.Website.ResolveHelpdeskDomain(websiteID string) (*WebsiteHelpdeskDomain, *Response, error)`
  * **Request Helpdesk Domain Change**: `client.Website.RequestHelpdeskDomainChange(websiteID string, domain WebsiteHelpdeskDomainChange) (*Response, error)`
  * **Generate Helpdesk Domain Setup Flow**: `client.Website.GenerateHelpdeskDomainSetupFlow(websiteID string, custom string) (*WebsiteHelpdeskDomainSetupFlow, *Response, error)`

### Plugin

* **One Plugin**
  * **Get Plugin Information**: `client.Plugin.GetPluginInformation(pluginID string) (*PluginInformation, *Response, error)`

* **Plugin List**
  * **List All Plugins**: `client.Plugin.ListAllPlugins(pageNumber uint) (*[]PluginInformation, *Response, error)`
  * **Search Plugins**: `client.Plugin.SearchPlugins(query string, pageNumber uint) (*[]PluginInformation, *Response, error)`

* **Plugin Subscription**
  * **List All Active Subscriptions**: `client.Plugin.ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error)`
  * **List Subscriptions For A Website**: `client.Plugin.ListSubscriptionsForWebsite(websiteID string) (*[]PluginSubscription, *Response, error)`
  * **Get Subscription Details**: `client.Plugin.GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error)`
  * **Subscribe Website To Plugin**: `client.Plugin.SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error)`
  * **Unsubscribe Plugin From Website**: `client.Plugin.UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error)`
  * **Get Subscription Settings**: `client.Plugin.GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error)`
  * **Save Subscription Settings**: `client.Plugin.SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Update Subscription Settings**: `client.Plugin.UpdateSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`
  * **Forward Plugin Payload To Channel**: `client.Plugin.ForwardPluginPayloadToChannel(websiteID string, pluginID string, payload PluginSubscriptionChannelForward) (*Response, error)`

* **Plugin Connect**
  * **Get Connect Account**: `client.Plugin.GetConnectAccount() (*PluginConnectAccount, *Response, error)`
  * **Check Connect Session Validity**: `client.Plugin.CheckConnectSessionValidity() (*Response, error)`
  * **List All Connect Websites**: `client.Plugin.ListAllConnectWebsites(pageNumber uint, filterConfigured bool) (*[]PluginConnectAllWebsites, *Response, error)`
  * **List Connect Websites Since**: `client.Plugin.ListConnectWebsitesSince(dateSince time.Time, filterConfigured bool) (*[]PluginConnectWebsitesSince, *Response, error)`

### Plan

* **One Plan**
  * **Get Plan Information**: `client.Plan.GetPlanInformation(planID string) (*PlanInformation, *Response, error)`

* **Plan List**
  * **List Plans**: `client.Plan.ListPlans() (*[]PlanInformation, *Response, error)`

* **Plan Subscription**
  * **List All Active Subscriptions**: `client.Plan.ListAllActiveSubscriptions() (*[]PlanSubscription, *Response, error)`
  * **Get Subscription For A Website**: `client.Plan.GetSubscriptionForWebsite(websiteID string) (*PlanSubscription, *Response, error)`
  * **Subscribe Website To Plan**: `client.Plan.SubscribeWebsiteToPlan(websiteID string, planID string) (*Response, error)`
  * **Unsubscribe Plan From Website**: `client.Plan.UnsubscribePlanFromWebsite(websiteID string) (*Response, error)`
  * **Change Bill Period For Website Subscription**: `client.Plan.ChangeBillPeriodForWebsiteSubscription(websiteID string, period string) (*Response, error)`
  * **Check Coupon Availability For Website Subscription**: `client.Plan.CheckCouponAvailabilityForWebsiteSubscription(websiteID string, code string) (*PlanSubscriptionCoupon, *Response, error)`
  * **Redeem Coupon For Website Subscription**: `client.Plan.RedeemCouponForWebsiteSubscription(websiteID string, code string) (*Response, error)`

## Realtime Events

You can bind to realtime events from Crisp, in order to get notified of incoming messages and updates in websites.

You won't receive any event if you don't explicitly subscribe to realtime events, as the library doesn't connect to the realtime backend automatically.

To start listening for events and bind an handler, proceed as follow:

```go
client := crisp.New()

// Set authentication parameters
client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")

// Connect to realtime events backend and listen (only to 'message:send' namespace)
client.Events.Listen([]string{"message:send"}, func(reg *crisp.EventsRegister) {
  // Now listening for events

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
})
```

Available events are listed below:

* **Session Events**
  * **Session Update Availability**: `session:update_availability`
  * **Session Update Verify**: `session:update_verify`
  * **Session Request Initiated**: `session:request:initiated`
  * **Session Set Email**: `session:set_email`
  * **Session Set Phone**: `session:set_phone`
  * **Session Set Address**: `session:set_address`
  * **Session Set Avatar**: `session:set_avatar`
  * **Session Set Nickname**: `session:set_nickname`
  * **Session Set Data**: `session:set_data`
  * **Session Sync Pages**: `session:sync:pages`
  * **Session Sync Events**: `session:sync:events`
  * **Session Sync Capabilities**: `session:sync:capabilities`
  * **Session Sync Geolocation**: `session:sync:geolocation`
  * **Session Sync System**: `session:sync:system`
  * **Session Sync Network**: `session:sync:network`
  * **Session Sync Timezone**: `session:sync:timezone`
  * **Session Sync Locales**: `session:sync:locales`
  * **Session Set State**: `session:set_state`
  * **Session Set Block**: `session:set_block`
  * **Session Set Segments**: `session:set_segments`
  * **Session Set Opened**: `session:set_opened`
  * **Session Set Closed**: `session:set_closed`
  * **Session Set Mention**: `session:set_mentions`
  * **Session Set Routing**: `session:set_routing`
  * **Session Removed**: `session:removed`

* **Message Events**
  * **Message Updated**: `message:updated`
  * **Message Send (Text Variant)**: `message:send/text`
  * **Message Send (File Variant)**: `message:send/file`
  * **Message Send (Animation Variant)**: `message:send/animation`
  * **Message Send (Audio Variant)**: `message:send/audio`
  * **Message Send (Picker Variant)**: `message:send/picker`
  * **Message Send (Field Variant)**: `message:send/field`
  * **Message Send (Note Variant)**: `message:send/note`
  * **Message Received (Text Variant)**: `message:received/text`
  * **Message Received (File Variant)**: `message:received/file`
  * **Message Received (Animation Variant)**: `message:received/animation`
  * **Message Received (Audio Variant)**: `message:received/audio`
  * **Message Received (Picker Variant)**: `message:received/picker`
  * **Message Received (Field Variant)**: `message:received/field`
  * **Message Received (Note Variant)**: `message:received/note`
  * **Message Compose Send**: `message:compose:send`
  * **Message Compose Receive**: `message:compose:receive`
  * **Message Acknowledge Read Send**: `message:acknowledge:read:send`
  * **Message Acknowledge Read Received**: `message:acknowledge:read:received`
  * **Message Acknowledge Delivered**: `message:acknowledge:delivered`
  * **Message Notify Unread Send**: `message:notify:unread:send`
  * **Message Notify Unread Received**: `message:notify:unread:received`

* **People Events**
  * **People Profile Created**: `people:profile:created`
  * **People Profile Removed**: `people:profile:removed`
  * **People Bind Session**: `people:bind:session`
  * **People Sync Profile**: `people:sync:profile`
  * **People Import Progress**: `people:import:progress`
  * **People Import Done**: `people:import:done`

* **Campaign Events**
  * **Campaign Progress**: `campaign:progress`
  * **Campaign Dispatched**: `campaign:dispatched`
  * **Campaign Running**: `campaign:running`

* **Browsing Events**
  * **Browsing Request Initiated**: `browsing:request:initiated`
  * **Browsing Request Rejected**: `browsing:request:rejected`

* **Call Events**
  * **Call Request Initiated**: `call:request:initiated`
  * **Call Request Rejected**: `call:request:rejected`

* **Service Events**
  * **Service Translate Processed**: `service:translate:processed`

* **Website Events**
  * **Website Update Visitors Count**: `website:update_visitors_count`
  * **Website Update Operators Availability**: `website:update_operators_availability`
  * **Website Users Available**: `website:users:available`
  * **Website Validate Domain Valid**: `website:validate:domain:valid`
  * **Website Validate Domain Invalid**: `website:validate:domain:invalid`

* **Bucket Events**
  * **Bucket URL Upload Generated**: `bucket:url:upload:generated`
  * **Bucket URL Avatar Generated**: `bucket:url:avatar:generated`
  * **Bucket URL Website Generated**: `bucket:url:website:generated`
  * **Bucket URL Campaign Generated**: `bucket:url:campaign:generated`
  * **Bucket URL Helpdesk Generated**: `bucket:url:helpdesk:generated`
  * **Bucket URL Processing Generated**: `bucket:url:processing:generated`

* **Media Events**
  * **Media Animation Listed**: `media:animation:listed`

* **Email Events**
  * **Email Subscribe**: `email:subscribe`
  * **Email Track View**: `email:track:view`

* **Billing Events**
  * **Billing Link Redirect**: `billing:link:redirect`

* **Plugin Events**
  * **Plugin Channel**: `plugin:channel`
  * **Plugin Settings Saved**: `plugin:settings:saved`
