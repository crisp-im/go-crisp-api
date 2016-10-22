# go-crisp-api

[![Build Status](https://travis-ci.org/crisp-im/go-crisp-api.svg?branch=master)](https://travis-ci.org/crisp-im/go-crisp-api)

The Crisp API Golang wrapper. Authenticate, send messages, fetch conversations, access your agent accounts from your Go code.

Copyright 2016 Crisp IM, Inc. See LICENSE for copying information.

* **📝 Implements**: [Crisp Platform - API ~ v1](https://docs.crisp.im/api/v1/) at reference revision: 10/22/2016
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

To authenticate to the API, generate your session identifier and session key [using this guide](https://docs.crisp.im/api/v1/#header-authentication).

Then, add authentication parameters to your `client` instance right after you create it:

```go
client := crisp.New()

// Authenticate to API (identifier, key)
client.Authenticate("7c3ef21c-1e04-41ce-8c06-5605c346f73e", "cc29e1a5086e428fcc6a697d5837a66d82808e65c5cce006fbf2191ceea80a0a")

// Now, you can use authenticated API sections.
```

## Resource Methods

All the available Crisp API resources are fully implemented. **Programmatic methods names are named after their label name in the [API Reference](https://docs.crisp.im/api/v1/)**.

Thus, it is straightforward to look for them in the library while reading the [API Reference](https://docs.crisp.im/api/v1/).

In the following method prototypes, `crisp` is to be replaced with your Crisp API instance. For example, instanciate `client := crisp.New()` and then call eg: `client.User.CheckSessionValidity()`.

When calling a method that writes data to the API (eg: `crisp.User.CreateUserAccount()`), you need to build an account instance this way:

```go
userAccount = crisp.UserAccountCreate{
  Email: "john@acme-inc.com"
  Passowrd: "SecurePassword"
  FirstName: "John"
  LastName: "Doe"
}
```

Refer directly to [the library source code](https://github.com/crisp-im/go-crisp-api/tree/master/crisp) to know more about those structures.

### Email

* **Email Subscription**
  * **Get Subscription Status**: `crisp.Email.GetSubscriptionStatus(emailHash string, key string) (*SubscriptionStatus, *Response, error)`
  * **Update Subscription Status**: `crisp.Email.UpdateSubscriptionStatus(emailHash string, key string, subscribed bool) (*Response, error)`

### Bucket

* **Bucket URL**
  * **Generate Bucket URL**: `crisp.Bucket.GenerateBucketURL(bucketData BucketURLRequest) (*Response, error)`

### User

* **User Availability**
  * **Get User Availability**: `crisp.User.GetUserAvailability() (*UserAvailability, *Response, error)`
  * **Update User Availability**: `crisp.User.UpdateUserAvailability(availabilityType string, timeFor uint) (*Response, error)`
  * **Get User Availability Status**: `crisp.User.GetUserAvailabilityStatus() (*UserAvailabilityStatus, *Response, error)`

* **User Account Base**
  * **Get User Account**: `crisp.User.GetUserAccount() (*UserAccount, *Response, error)`
  * **Create User Account**: `crisp.User.CreateUserAccount(user UserAccountCreate) (*Response, error)`

* **User Account Billing**
  * **List All Billing Methods**: `crisp.User.ListAllBillingMethods() (*[]BillingMethodAll, *Response, error)`
  * **Add New Billing Method**: `crisp.User.AddNewBillingMethod(billing BillingMethodCreate) (*Response, error)`
  * **Get A Billing Method**: `crisp.User.GetBillingMethod(cardID string) (*BillingMethod, *Response, error)`
  * **Remove A Billing Method**: `crisp.User.RemoveBillingMethod(cardID string) (*Response, error)`
  * **List Invoices For Billing Method**: `crisp.User.ListInvoicesForBillingMethod(cardID string, pageNumber int) (*BillingMethodInvoiceAll, *Response, error)`
  * **Get An Invoice For Billing Method**: `crisp.User.GetInvoiceForBillingMethod(cardID string, invoiceID string) (*BillingMethodInvoice, *Response, error)`
  * **Link To An External Billing Method**: `crisp.User.LinkToExternalBillingMethod(billingService string) (*Response, error)`
  * **Finish Linking External Billing Method**: `crisp.User.FinishLinkingExternalBillingMethod(cardID string) (*Response, error)`
  * **Finish Unlinking External Billing Method**: `crisp.User.FinishUnlinkingExternalBillingMethod(cardID string) (*Response, error)`

* **User Account Notification**
  * **Get Notification Settings**: `crisp.User.GetNotificationSettings() (*UserNotificationSettings, *Response, error)`
  * **Update Notification Settings**: `crisp.User.UpdateNotificationSettings(notifications UserNotificationSettingsUpdate) (*Response, error)`
  * **Add A Notification Provider**: `crisp.User.AddNotificationProvider(notificationID string) (*Response, error)`

* **User Account Websites**
  * **List Websites**: `crisp.User.ListWebsites() (*[]UserWebsite, *Response, error)`

* **User Account Profile**
  * **Get Profile**: `crisp.User.GetProfile() (*UserProfile, *Response, error)`
  * **Update Profile**: `crisp.User.UpdateProfile(profile UserProfileUpdate) (*Response, error)`

* **User Account Recover**
  * **Get Recovery Details**: `crisp.User.GetRecoveryDetails(userID string, recoverIdentifier string, recoverKey string) (*Response, error)`
  * **Send Recovery Password**: `crisp.User.SendRecoveryPassword(userID string, recoverIdentifier string, recoverKey string, password string) (*Response, error)`
  * **Delete Recovery Keypair**: `crisp.User.DeleteRecoveryKeypair(userID string, recoverIdentifier string, recoverKey string) (*Response, error)`

* **User Account Schedule**
  * **Get Schedule Settings**: `crisp.User.GetScheduleSettings() (*UserScheduleSettings, *Response, error)`
  * **Update Schedule Settings**: `crisp.User.UpdateScheduleSettings(schedule UserScheduleSettingsUpdate) (*Response, error)`

* **User Session**
  * **Check Session Validity**: `crisp.User.CheckSessionValidity() (*Response, error)`
  * **Create A New Session**: `crisp.User.CreateNewSession(email string, password string) (*UserSessionParameters, *Response, error)`
  * **Destroy A Session**: `crisp.User.DestroySession() (*Response, error)`
  * **Recover A Session**: `crisp.User.RecoverSession(email string)`

* **User Statistics**
  * **Count Total Unread Messages**: `crisp.User.CountTotalUnreadMessages() (*UserStatistics, *Response, error)`

### Website

* **Website Base**
  * **Create Website**: `crisp.Website.CreateWebsite(websiteData WebsiteCreate) (*Website, *Response, error)`
  * **Delete A Website**: `crisp.Website.DeleteWebsite(websiteID string) (*Response, error)`

* **Website Batch**
  * **Resolve All Conversations**: `crisp.Website.ResolveAllConversations(websiteID string) (*Response, error)`
  * **Read All Conversations**: `crisp.Website.ReadAllConversations(websiteID string) (*Response, error)`

* **Website Billing**
  * **Get Website Billing**: `crisp.Website.GetWebsiteBilling(websiteID string) (*WebsiteBilling, *Response, error)`
  * **Update Website Billing**: `crisp.Website.UpdateWebsiteBilling(websiteID string, cardID string) (*Response, error)`
  * **Unlink Website Billing**: `crisp.Website.UnlinkWebsiteBilling(websiteID string) (*Response, error)`

* **Website Availability**
  * **Get Website Availability Status**: `crisp.Website.GetWebsiteAvailabilityStatus(websiteID string) (*WebsiteAvailabilityStatus, *Response, error)`

* **Website Operator**
  * **List Website Operators**: `crisp.Website.ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error)`
  * **Get A Website Operator**: `crisp.Website.GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error)`
  * **Invite A Website Operator**: `crisp.Website.InviteWebsiteOperator(websiteID string, email string, role string) (*Response, error)`
  * **Change Operator Role**: `crisp.Website.ChangeOperatorRole(websiteID string, userID string, role string) (*Response, error)`
  * **Unlink Operator From Website**: `crisp.Website.UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error)`

* **Website Invite**
  * **Get Invite Details**: `crisp.Website.GetInviteDetails(websiteID string, inviteIdentifier string, inviteKey string) (*WebsiteInvite, *Response, error)`
  * **Redeem Invite**: `crisp.Website.RedeemInvite(websiteID string, inviteIdentifier string, inviteKey string) (*Response, error)`
  * **Delete Invite Keypair**: `crisp.Website.DeleteInviteKeypair(websiteID string, inviteIdentifier string, inviteKey string) (*Response, error)`

* **Website Statistics**
  * **Get All Statistics**: `crisp.Website.GetAllStatistics(websiteID string) (*WebsiteStatisticsAll, *Response, error)`
  * **Count Total Number Of Conversations**: `crisp.Website.CountTotalNumberOfConversations(websiteID string) (*WebsiteStatisticsTotal, *Response, error)`
  * **Count Number Of Pending Conversations**: `crisp.Website.CountNumberOfPendingConversations(websiteID string) (*WebsiteStatisticsPending, *Response, error)`
  * **Count Number Of Unresolved Conversations**: `crisp.Website.CountNumberOfUnresolvedConversations(websiteID string) (*WebsiteStatisticsUnresolved, *Response, error)`
  * **Count Number Of Resolved Conversations**: `crisp.Website.CountNumberOfResolvedConversations(websiteID string) (*WebsiteStatisticsResolved, *Response, error)`
  * **Count Number Of Unread Messages**: `crisp.Website.CountNumberOfUnreadMessages(websiteID string) (*WebsiteStatisticsUnread, *Response, error)`

* **Website Settings**
  * **Get Website Settings**: `crisp.Website.GetWebsiteSettings(websiteID string) (*WebsiteSettings, *Response, error)`
  * **Update Website Settings**: `crisp.Website.UpdateWebsiteSettings(websiteID string, settings WebsiteSettingsUpdate) (*Response, error)`

* **Website Visitors**
  * **Count Visitors**: `crisp.Website.CountVisitors(websiteID string) (*WebsiteVisitorCount, *Response, error)`
  * **List Visitors**: `crisp.Website.ListVisitors(websiteID string, pageNumber uint) (*[]WebsiteVisitor, *Response, error)`
  * **Request Visitor Details**: `crisp.Website.RequestVisitorDetails(websiteID string) (*Response, error)`

* **Website States**
  * **Request Website States**: `crisp.Website.RequestWebsiteStates(websiteID string) (*Response, error)`

* **Website Conversations**
  * **List Conversations**: `crisp.Website.ListConversations(websiteID string, pageNumber uint) (*[]Conversation, *Response, error)`
  * **List Conversations (Search Variant)**: `crisp.Website.SearchConversations(websiteID string, pageNumber uint, searchQuery string, searchType string) (*[]Conversation, *Response, error)`
  * **Export Conversation Emails**: `crisp.Website.ExportConversationEmails(websiteID string) (*[]ConversationExportEmail, *Response, error)`

* **Website Conversation**
  * **Create A New Conversation**: `crisp.Website.CreateNewConversation(websiteID string) (*ConversationNew, *Response, error)`
  * **Check If Conversation Exists**: `crisp.Website.CheckConversationExists(websiteID string, sessionID string) (*Response, error)`
  * **Get A Conversation**: `crisp.Website.GetConversation(websiteID string, sessionID string) (*Conversation, *Response, error)`
  * **Remove A Conversation**: `crisp.Website.RemoveConversation(websiteID string, sessionID string) (*Response, error)`
  * **Initiate A Conversation With Existing Session**: `crisp.Website.InitiateConversationWithExistingSession(websiteID string, sessionID string) (*Response, error)`
  * **Get Messages In Conversation**: `crisp.Website.GetMessagesInConversation(websiteID string, sessionID string) (*[]ConversationMessage, *Response, error)`
  * **Send A Message In Conversation (Text Variant)**: `crisp.Website.SendTextMessageInConversation(websiteID string, sessionID string, message ConversationTextMessageNew) (*Response, error)`
  * **Send A Message In Conversation (File Variant)**: `crisp.Website.SendFileMessageInConversation(websiteID string, sessionID string, message ConversationFileMessageNew) (*Response, error)`
  * **Suggest Message Completion**: `crisp.Website.SuggestMessageCompletion(websiteID string, sessionID string, message string) (*ConversationCompletion, *Response, error)`
  * **Compose A Message In Conversation**: `crisp.Website.ComposeMessageInConversation(websiteID string, sessionID string, compose ConversationComposeMessageNew) (*Response, error)`
  * **Mark Messages As Read In Conversation**: `crisp.Website.MarkMessagesReadInConversation(websiteID string, sessionID string, read ConversationReadMessageMark) (*Response, error)`
  * **Update Conversation Open State**: `crisp.Website.UpdateConversationOpenState(websiteID string, sessionID string, opened bool) (*Response, error)`
  * **Get Conversation Metas**: `crisp.Website.GetConversationMetas(websiteID string, sessionID string) (*ConversationMeta, *Response, error)`
  * **Update Conversation Metas**: `crisp.Website.UpdateConversationMetas(websiteID string, sessionID string, metas ConversationMetaUpdate) (*Response, error)`
  * **Get Conversation State**: `crisp.Website.GetConversationState(websiteID string, sessionID string) (*ConversationState, *Response, error)`
  * **Change Conversation State**: `crisp.Website.ChangeConversationState(websiteID string, sessionID string, state string) (*Response, error)`
  * **Get Block Status For Conversation**: `crisp.Website.GetBlockStatusForConversation(websiteID string, sessionID string) (*ConversationBlock, *Response, error)`
  * **Block Incoming Messages For Conversation**: `crisp.Website.BlockIncomingMessagesForConversation(websiteID string, sessionID string, blocked bool) (*Response, error)`

### Plugin

* **One Plugin**
  * **Get Plugin Information**: `crisp.Plugin.GetPluginInformation(pluginID string) (*PluginInformation, *Response, error)`
  * **Get Plugin Stars**: `crisp.Plugin.GetPluginStars(pluginID string)`
  * **Get Personal Plugin Rank**: `crisp.Plugin.GetPersonalPluginRank(pluginID string) (*PluginPersonalPluginRank, *Response, error)`
  * **Rank A Plugin**: `crisp.Plugin.RankPlugin(pluginID string, rank uint) (*Response, error)`
  * **Delete Plugin Rank**: `crisp.Plugin.DeletePluginRank(pluginID string) (*Response, error)`

* **Plugin List**
  * **List All Plugins**: `crisp.Plugin.ListAllPlugins(pageNumber uint) (*[]PluginInformation, *Response, error)`
  * **List Featured Plugins**: `crisp.Plugin.ListFeaturedPlugins(pageNumber uint) (*[]PluginInformation, *Response, error)`
  * **Search Plugins**: `crisp.Plugin.SearchPlugins(query string, pageNumber uint) (*[]PluginInformation, *Response, error)`

* **Plugin Subscription**
  * **List All Active Subscriptions**: `crisp.Plugin.ListAllActiveSubscriptions() (*[]PluginSubscription, *Response, error)`
  * **List Subscriptions For A Website**: `crisp.Plugin.ListSubscriptionsForWebsite(websiteID string) (*[]PluginSubscription, *Response, error)`
  * **Get Subscription Details**: `crisp.Plugin.GetSubscriptionDetails(websiteID string, pluginID string) (*PluginSubscription, *Response, error)`
  * **Subscribe Website To Plugin**: `crisp.Plugin.SubscribeWebsiteToPlugin(websiteID string, pluginID string) (*Response, error)`
  * **Unsubscribe Plugin From Website**: `crisp.Plugin.UnsubscribePluginFromWebsite(websiteID string, pluginID string) (*Response, error)`
  * **Get Subscription Settings**: `crisp.Plugin.GetSubscriptionSettings(websiteID string, pluginID string) (*PluginSubscriptionSettings, *Response, error)`
  * **Save Subscription Settings**: `crisp.Plugin.SaveSubscriptionSettings(websiteID string, pluginID string, settings interface{}) (*Response, error)`

* **Plugin Connect**
  * **Get Connect Account**: `crisp.Plugin.GetConnectAccount() (*PluginConnectAccount, *Response, error)`
  * **Check Connect Session Validity**: `crisp.Plugin.CheckConnectSessionValidity() (*Response, error)`
  * **List All Connect Websites**: `crisp.Plugin.ListAllConnectWebsites(pageNumber uint) (*[]PluginConnectAllWebsites, *Response, error)`
  * **List Connect Websites Since**: `crisp.Plugin.ListConnectWebsitesSince(dateSince time.Time) (*[]PluginConnectWebsitesSince, *Response, error)`

### Plan

* **One Plan**
  * **Get Plan Information**: `crisp.Plan.GetPlanInformation(planID string) (*PlanInformation, *Response, error)`

* **Plan List**
  * **List Plans**: `crisp.Plan.ListPlans() (*[]PlanInformation, *Response, error)`

* **Plan Subscription**
  * **List All Active Subscriptions**: `crisp.Plan.ListAllActiveSubscriptions() (*[]PlanSubscription, *Response, error)`
  * **Get Subscription For A Website**: `crisp.Plan.GetSubscriptionForWebsite(websiteID string) (*PlanSubscription, *Response, error)`
  * **Subscribe Website To Plan**: `crisp.Plan.SubscribeWebsiteToPlan(websiteID string, planID string) (*Response, error)`
  * **Unsubscribe Plan From Website**: `crisp.Plan.UnsubscribePlanFromWebsite(websiteID string) (*Response, error)`

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
})
```

Available events are listed below:

* **Session Events**
  * **Session Update Availability**: `session:update_availability`
  * **Session Request Initiated**: `session:request:initiated`
  * **Session Set Email**: `session:set_email`
  * **Session Set Avatar**: `session:set_avatar`
  * **Session Set Cover**: `session:set_cover`
  * **Session Set Nickname**: `session:set_nickname`
  * **Session Set Data**: `session:set_data`
  * **Session Sync Pages**: `session:sync:pages`
  * **Session Sync Geolocation**: `session:sync:geolocation`
  * **Session Sync System**: `session:sync:system`
  * **Session Sync Network**: `session:sync:network`
  * **Session Sync Extended Information**: `session:sync:extended_informations`
  * **Session Set State**: `session:set_state`
  * **Session Set Block**: `session:set_block`
  * **Session Set Tags**: `session:set_tags`
  * **Session Set Opened**: `session:set_opened`
  * **Session Set Closed**: `session:set_closed`
  * **Session Removed**: `session:removed`

* **Message Events**
  * **Message Send (Text Variant)**: `message:send/text`
  * **Message Send (File Variant)**: `message:send/file`
  * **Message Received (Text Variant)**: `message:received/text`
  * **Message Received (File Variant)**: `message:received/file`
  * **Message Compose Send**: `message:compose:send`
  * **Message Compose Receive**: `message:compose:receive`
  * **Message Acknowledge Read Send**: `message:acknowledge:read:send`
  * **Message Acknowledge Read Received**: `message:acknowledge:read:received`
  * **Message Acknowledge Delivered**: `message:acknowledge:delivered`

* **Website Events**
  * **Website Update Visitors Count**: `website:update_visitors_count`
  * **Website Update Visitors List**: `website:update_visitors_list`

* **Bucket Events**
  * **Bucket URL Upload Generated**: `bucket:url:upload:generated`
  * **Bucket URL Avatar Generated**: `bucket:url:avatar:generated`

* **Billing Events**
  * **Billing Link Redirect**: `billing:link:redirect`
