// Code generated from the Norbix Hub contract. DO NOT EDIT.
//
// Regenerate with the type generator described in README.md ("Regenerate the
// types"). Editing this file by hand is lost on the next run.

// Package dtos holds the data types of the Norbix Hub surface. A field
// that holds another generated type is a pointer; every field carries a JSON
// tag with omitempty, because the gateway leaves absent values out.
package dtos

// EmailProvider enum.
type EmailProvider string

const (
	EmailProviderSmtp     EmailProvider = "Smtp"
	EmailProviderSendGrid EmailProvider = "SendGrid"
	EmailProviderMailGun  EmailProvider = "MailGun"
	EmailProviderAwsSes   EmailProvider = "AwsSes"
	EmailProviderFake     EmailProvider = "Fake"
)

// SmtpPorts enum.
type SmtpPorts string

const (
	SmtpPortsDefault  SmtpPorts = "Default"
	SmtpPortsSsl      SmtpPorts = "Ssl"
	SmtpPortsTls      SmtpPorts = "Tls"
	SmtpPortsFallback SmtpPorts = "Fallback"
)

// AwsIntegrationType enum.
type AwsIntegrationType string

const (
	AwsIntegrationTypeIam              AwsIntegrationType = "Iam"
	AwsIntegrationTypeCrossAccountRole AwsIntegrationType = "CrossAccountRole"
)

// MailGunRegion enum.
type MailGunRegion string

const (
	MailGunRegionUs MailGunRegion = "Us"
	MailGunRegionEu MailGunRegion = "Eu"
)

// EmailCampaignRecipientsSourceTypes enum.
type EmailCampaignRecipientsSourceTypes string

const (
	EmailCampaignRecipientsSourceTypesAllUsers       EmailCampaignRecipientsSourceTypes = "AllUsers"
	EmailCampaignRecipientsSourceTypesSpecifiedUsers EmailCampaignRecipientsSourceTypes = "SpecifiedUsers"
	EmailCampaignRecipientsSourceTypesAccountUsers   EmailCampaignRecipientsSourceTypes = "AccountUsers"
	EmailCampaignRecipientsSourceTypesEmail          EmailCampaignRecipientsSourceTypes = "Email"
	EmailCampaignRecipientsSourceTypesCollection     EmailCampaignRecipientsSourceTypes = "Collection"
)

// TokenMappingResolverType enum.
type TokenMappingResolverType string

const (
	TokenMappingResolverTypeNotSet          TokenMappingResolverType = "NotSet"
	TokenMappingResolverTypeCustom          TokenMappingResolverType = "Custom"
	TokenMappingResolverTypeProject         TokenMappingResolverType = "Project"
	TokenMappingResolverTypeProjectSocials  TokenMappingResolverType = "ProjectSocials"
	TokenMappingResolverTypeInitiator       TokenMappingResolverType = "Initiator"
	TokenMappingResolverTypeRecipient       TokenMappingResolverType = "Recipient"
	TokenMappingResolverTypeSchemaRecord    TokenMappingResolverType = "SchemaRecord"
	TokenMappingResolverTypeTargetUser      TokenMappingResolverType = "TargetUser"
	TokenMappingResolverTypeTagDefinitions  TokenMappingResolverType = "TagDefinitions"
	TokenMappingResolverTypeEmailSignatures TokenMappingResolverType = "EmailSignatures"
	TokenMappingResolverTypeCampaign        TokenMappingResolverType = "Campaign"
	TokenMappingResolverTypeTemplate        TokenMappingResolverType = "Template"
	TokenMappingResolverTypeEmailFooters    TokenMappingResolverType = "EmailFooters"
	TokenMappingResolverTypeOld             TokenMappingResolverType = "Old"
	TokenMappingResolverTypeNew             TokenMappingResolverType = "New"
)

// CollectionEmailCampaignRecipientField enum.
type CollectionEmailCampaignRecipientField string

const (
	CollectionEmailCampaignRecipientFieldUser  CollectionEmailCampaignRecipientField = "User"
	CollectionEmailCampaignRecipientFieldEmail CollectionEmailCampaignRecipientField = "Email"
)

// TriggerType enum.
type TriggerType string

const (
	TriggerTypeMembership TriggerType = "Membership"
	TriggerTypeSchema     TriggerType = "Schema"
	TriggerTypeFiles      TriggerType = "Files"
	TriggerTypePayments   TriggerType = "Payments"
)

// TriggerActionType enum.
type TriggerActionType string

const (
	TriggerActionTypeCode        TriggerActionType = "Code"
	TriggerActionTypePush        TriggerActionType = "Push"
	TriggerActionTypeSms         TriggerActionType = "Sms"
	TriggerActionTypeEmail       TriggerActionType = "Email"
	TriggerActionTypeWebhookCall TriggerActionType = "WebhookCall"
	TriggerActionTypeSseCall     TriggerActionType = "SseCall"
	TriggerActionTypeMarketplace TriggerActionType = "Marketplace"
)

// MembershipTriggerType enum.
type MembershipTriggerType string

const (
	MembershipTriggerTypeOnRegistered  MembershipTriggerType = "OnRegistered"
	MembershipTriggerTypeOnInvited     MembershipTriggerType = "OnInvited"
	MembershipTriggerTypeOnVerified    MembershipTriggerType = "OnVerified"
	MembershipTriggerTypeOnUpdated     MembershipTriggerType = "OnUpdated"
	MembershipTriggerTypeOnDeleted     MembershipTriggerType = "OnDeleted"
	MembershipTriggerTypeOnBlocked     MembershipTriggerType = "OnBlocked"
	MembershipTriggerTypeOnReactivated MembershipTriggerType = "OnReactivated"
	MembershipTriggerTypeOnUserCreated MembershipTriggerType = "OnUserCreated"
)

// SchemaTriggerType enum.
type SchemaTriggerType string

const (
	SchemaTriggerTypeOnInserted SchemaTriggerType = "OnInserted"
	SchemaTriggerTypeOnDeleted  SchemaTriggerType = "OnDeleted"
	SchemaTriggerTypeOnUpdated  SchemaTriggerType = "OnUpdated"
)

// FilesTriggerType enum.
type FilesTriggerType string

const (
	FilesTriggerTypeOnFileUploaded FilesTriggerType = "OnFileUploaded"
	FilesTriggerTypeOnFileDeleted  FilesTriggerType = "OnFileDeleted"
)

// FileProvider enum.
type FileProvider string

const (
	FileProviderLocal              FileProvider = "Local"
	FileProviderAwsS3              FileProvider = "AwsS3"
	FileProviderAzureBlobStorage   FileProvider = "AzureBlobStorage"
	FileProviderGoogleCloudStorage FileProvider = "GoogleCloudStorage"
	FileProviderFtp                FileProvider = "Ftp"
	FileProviderAppleICloud        FileProvider = "AppleICloud"
	FileProviderDropBox            FileProvider = "DropBox"
	FileProviderGoogleDrive        FileProvider = "GoogleDrive"
)

// PaymentTriggerType enum.
type PaymentTriggerType string

const (
	PaymentTriggerTypeOnOrderCreated        PaymentTriggerType = "OnOrderCreated"
	PaymentTriggerTypeOnOrderPaid           PaymentTriggerType = "OnOrderPaid"
	PaymentTriggerTypeOnWebhookCallReceived PaymentTriggerType = "OnWebhookCallReceived"
)

// DatabaseProvider enum.
type DatabaseProvider string

const (
	DatabaseProviderMongoDbConnectionString         DatabaseProvider = "MongoDbConnectionString"
	DatabaseProviderCodeMashMongoDbAtlasFlexManaged DatabaseProvider = "CodeMashMongoDbAtlasFlexManaged"
)

// AwsS3IntegrationType enum.
type AwsS3IntegrationType string

const (
	AwsS3IntegrationTypeIam              AwsS3IntegrationType = "Iam"
	AwsS3IntegrationTypeCrossAccountRole AwsS3IntegrationType = "CrossAccountRole"
)

// LoggingProvider enum.
type LoggingProvider string

const (
	LoggingProviderConsole          LoggingProvider = "Console"
	LoggingProviderNorbixLogging    LoggingProvider = "NorbixLogging"
	LoggingProviderDataDog          LoggingProvider = "DataDog"
	LoggingProviderNewRelic         LoggingProvider = "NewRelic"
	LoggingProviderSentry           LoggingProvider = "Sentry"
	LoggingProviderGrafanaLoki      LoggingProvider = "GrafanaLoki"
	LoggingProviderAxiom            LoggingProvider = "Axiom"
	LoggingProviderElasticCloud     LoggingProvider = "ElasticCloud"
	LoggingProviderAWSCloudWatch    LoggingProvider = "AWSCloudWatch"
	LoggingProviderGCPCloudLogging  LoggingProvider = "GCPCloudLogging"
	LoggingProviderAzureMonitorLogs LoggingProvider = "AzureMonitorLogs"
	LoggingProviderGenericHttp      LoggingProvider = "GenericHttp"
	LoggingProviderKafka            LoggingProvider = "Kafka"
	LoggingProviderAMQP             LoggingProvider = "AMQP"
	LoggingProviderPrometheus       LoggingProvider = "Prometheus"
	LoggingProviderAzureOTel        LoggingProvider = "AzureOTel"
	LoggingProviderSplunk           LoggingProvider = "Splunk"
	LoggingProviderElasticSearch    LoggingProvider = "ElasticSearch"
	LoggingProviderKibana           LoggingProvider = "Kibana"
	LoggingProviderLocalFile        LoggingProvider = "LocalFile"
	LoggingProviderAWSS3            LoggingProvider = "AWSS3"
	LoggingProviderAWSKinesis       LoggingProvider = "AWSKinesis"
	LoggingProviderMongoDB          LoggingProvider = "MongoDB"
	LoggingProviderInternalKafka    LoggingProvider = "InternalKafka"
)

// AwsS3LoggingIntegrationType enum.
type AwsS3LoggingIntegrationType string

const (
	AwsS3LoggingIntegrationTypeIam              AwsS3LoggingIntegrationType = "Iam"
	AwsS3LoggingIntegrationTypeCrossAccountRole AwsS3LoggingIntegrationType = "CrossAccountRole"
)

// MembershipProvider enum.
type MembershipProvider string

const (
	MembershipProviderAppleSignIn  MembershipProvider = "AppleSignIn"
	MembershipProviderGoogleSignIn MembershipProvider = "GoogleSignIn"
	MembershipProviderGoogle       MembershipProvider = "Google"
	MembershipProviderFacebook     MembershipProvider = "Facebook"
	MembershipProviderX            MembershipProvider = "X"
	MembershipProviderGitHub       MembershipProvider = "GitHub"
	MembershipProviderLinkedIn     MembershipProvider = "LinkedIn"
	MembershipProviderOkta         MembershipProvider = "Okta"
	MembershipProviderMicrosoft    MembershipProvider = "Microsoft"
)

// PaymentGatewayPlatform enum.
type PaymentGatewayPlatform string

const (
	PaymentGatewayPlatformStripe       PaymentGatewayPlatform = "Stripe"
	PaymentGatewayPlatformAdyen        PaymentGatewayPlatform = "Adyen"
	PaymentGatewayPlatformPaddle       PaymentGatewayPlatform = "Paddle"
	PaymentGatewayPlatformLemonSqueezy PaymentGatewayPlatform = "LemonSqueezy"
	PaymentGatewayPlatformAppleInApp   PaymentGatewayPlatform = "AppleInApp"
	PaymentGatewayPlatformGoogleInApp  PaymentGatewayPlatform = "GoogleInApp"
	PaymentGatewayPlatformShopify      PaymentGatewayPlatform = "Shopify"
	PaymentGatewayPlatformWooCommerce  PaymentGatewayPlatform = "WooCommerce"
	PaymentGatewayPlatformMagento      PaymentGatewayPlatform = "Magento"
	PaymentGatewayPlatformPayPal       PaymentGatewayPlatform = "PayPal"
	PaymentGatewayPlatformBraintree    PaymentGatewayPlatform = "Braintree"
	PaymentGatewayPlatformAuthorizeNet PaymentGatewayPlatform = "AuthorizeNet"
	PaymentGatewayPlatformCheckOutCom  PaymentGatewayPlatform = "CheckOutCom"
	PaymentGatewayPlatformMollie       PaymentGatewayPlatform = "Mollie"
	PaymentGatewayPlatformWorldpay     PaymentGatewayPlatform = "Worldpay"
)

// PushProvider enum.
type PushProvider string

const (
	PushProviderAppleApns            PushProvider = "AppleApns"
	PushProviderSafariWeb            PushProvider = "SafariWeb"
	PushProviderSafariPush           PushProvider = "SafariPush"
	PushProviderAndroidFirebase      PushProvider = "AndroidFirebase"
	PushProviderChromeWeb            PushProvider = "ChromeWeb"
	PushProviderFirefoxWeb           PushProvider = "FirefoxWeb"
	PushProviderEdgeWeb              PushProvider = "EdgeWeb"
	PushProviderChromePush           PushProvider = "ChromePush"
	PushProviderCodeMashIosApp       PushProvider = "CodeMashIosApp"
	PushProviderCodeMashAndroidApp   PushProvider = "CodeMashAndroidApp"
	PushProviderCodeMashSafariPlugin PushProvider = "CodeMashSafariPlugin"
	PushProviderCodeMashSafariWeb    PushProvider = "CodeMashSafariWeb"
	PushProviderCodeMashChromePlugin PushProvider = "CodeMashChromePlugin"
	PushProviderCodeMashChromeWeb    PushProvider = "CodeMashChromeWeb"
	PushProviderExpo                 PushProvider = "Expo"
	PushProviderFake                 PushProvider = "Fake"
)

// CodeProvider enum.
type CodeProvider string

const (
	CodeProviderAwsLambda            CodeProvider = "AwsLambda"
	CodeProviderAzureFunctions       CodeProvider = "AzureFunctions"
	CodeProviderGoogleCloudFunctions CodeProvider = "GoogleCloudFunctions"
	CodeProviderPipedream            CodeProvider = "Pipedream"
	CodeProviderZapier               CodeProvider = "Zapier"
	CodeProviderCloudflareWorkers    CodeProvider = "CloudflareWorkers"
	CodeProviderVercel               CodeProvider = "Vercel"
	CodeProviderNetlify              CodeProvider = "Netlify"
	CodeProviderSupabaseEdge         CodeProvider = "SupabaseEdge"
	CodeProviderModal                CodeProvider = "Modal"
)

// AwsLambdaIntegrationType enum.
type AwsLambdaIntegrationType string

const (
	AwsLambdaIntegrationTypeIam              AwsLambdaIntegrationType = "Iam"
	AwsLambdaIntegrationTypeCrossAccountRole AwsLambdaIntegrationType = "CrossAccountRole"
)

// LlmProvider enum.
type LlmProvider string

const (
	LlmProviderOpenAI       LlmProvider = "OpenAI"
	LlmProviderAnthropic    LlmProvider = "Anthropic"
	LlmProviderOllama       LlmProvider = "Ollama"
	LlmProviderGroq         LlmProvider = "Groq"
	LlmProviderGoogle       LlmProvider = "Google"
	LlmProviderMistral      LlmProvider = "Mistral"
	LlmProviderOpenRouter   LlmProvider = "OpenRouter"
	LlmProviderGrok         LlmProvider = "Grok"
	LlmProviderNorbixHosted LlmProvider = "NorbixHosted"
)

// McpProvider enum.
type McpProvider string

const (
	McpProviderDocker         McpProvider = "Docker"
	McpProviderObsidian       McpProvider = "Obsidian"
	McpProviderGoogleCalendar McpProvider = "GoogleCalendar"
	McpProviderStripe         McpProvider = "Stripe"
	McpProviderGitHub         McpProvider = "GitHub"
	McpProviderMongoDb        McpProvider = "MongoDb"
	McpProviderPlaywright     McpProvider = "Playwright"
	McpProviderBraveSearch    McpProvider = "BraveSearch"
)

// McpTransport enum.
type McpTransport string

const (
	McpTransportSse        McpTransport = "Sse"
	McpTransportHttpStream McpTransport = "HttpStream"
	McpTransportStdio      McpTransport = "Stdio"
)

// CommunicationChannel enum.
type CommunicationChannel string

const (
	CommunicationChannelTransactional CommunicationChannel = "Transactional"
	CommunicationChannelMarketing     CommunicationChannel = "Marketing"
	CommunicationChannelSystem        CommunicationChannel = "System"
)

// NotificationMedium enum.
type NotificationMedium string

const (
	NotificationMediumEmail NotificationMedium = "Email"
	NotificationMediumSms   NotificationMedium = "Sms"
	NotificationMediumPush  NotificationMedium = "Push"
)

// EmailTemplateEngine enum.
type EmailTemplateEngine string

const (
	EmailTemplateEngineNotSet     EmailTemplateEngine = "NotSet"
	EmailTemplateEngineHandlebars EmailTemplateEngine = "Handlebars"
	EmailTemplateEngineMjml       EmailTemplateEngine = "Mjml"
	EmailTemplateEngineLiquid     EmailTemplateEngine = "Liquid"
	EmailTemplateEngineRazor      EmailTemplateEngine = "Razor"
	EmailTemplateEngineMustache   EmailTemplateEngine = "Mustache"
)

// SystemEmailTemplateTheme enum.
type SystemEmailTemplateTheme string

const (
	SystemEmailTemplateThemeText     SystemEmailTemplateTheme = "Text"
	SystemEmailTemplateThemeBranded  SystemEmailTemplateTheme = "Branded"
	SystemEmailTemplateThemeCreative SystemEmailTemplateTheme = "Creative"
)

// RespectTimeZoneSettings enum.
type RespectTimeZoneSettings string

const (
	RespectTimeZoneSettingsRespectToLastLoginZone           RespectTimeZoneSettings = "RespectToLastLoginZone"
	RespectTimeZoneSettingsRespectToRegistrationZone        RespectTimeZoneSettings = "RespectToRegistrationZone"
	RespectTimeZoneSettingsRespectToRegistrationProjectZone RespectTimeZoneSettings = "RespectToRegistrationProjectZone"
)

// PushCampaignRecipientsSourceTypes enum.
type PushCampaignRecipientsSourceTypes string

const (
	PushCampaignRecipientsSourceTypesAllUsers       PushCampaignRecipientsSourceTypes = "AllUsers"
	PushCampaignRecipientsSourceTypesSpecifiedUsers PushCampaignRecipientsSourceTypes = "SpecifiedUsers"
	PushCampaignRecipientsSourceTypesCollection     PushCampaignRecipientsSourceTypes = "Collection"
	PushCampaignRecipientsSourceTypesDevices        PushCampaignRecipientsSourceTypes = "Devices"
	PushCampaignRecipientsSourceTypesAccountUsers   PushCampaignRecipientsSourceTypes = "AccountUsers"
)

// SmsCampaignRecipientsSourceTypes enum.
type SmsCampaignRecipientsSourceTypes string

const (
	SmsCampaignRecipientsSourceTypesAllUsers       SmsCampaignRecipientsSourceTypes = "AllUsers"
	SmsCampaignRecipientsSourceTypesSpecifiedUsers SmsCampaignRecipientsSourceTypes = "SpecifiedUsers"
	SmsCampaignRecipientsSourceTypesAccountUsers   SmsCampaignRecipientsSourceTypes = "AccountUsers"
	SmsCampaignRecipientsSourceTypesPhoneNumbers   SmsCampaignRecipientsSourceTypes = "PhoneNumbers"
	SmsCampaignRecipientsSourceTypesCollection     SmsCampaignRecipientsSourceTypes = "Collection"
)

// McpAuth enum.
type McpAuth string

const (
	McpAuthOAuth2 McpAuth = "OAuth2"
	McpAuthApiKey McpAuth = "ApiKey"
	McpAuthNone   McpAuth = "None"
)

// IntegrationStatus enum.
type IntegrationStatus string

const (
	IntegrationStatusUnknown        IntegrationStatus = "Unknown"
	IntegrationStatusPending        IntegrationStatus = "Pending"
	IntegrationStatusProvisioning   IntegrationStatus = "Provisioning"
	IntegrationStatusActive         IntegrationStatus = "Active"
	IntegrationStatusFailed         IntegrationStatus = "Failed"
	IntegrationStatusDeprovisioning IntegrationStatus = "Deprovisioning"
)

// SmsProvider enum.
type SmsProvider string

const (
	SmsProviderTwilio   SmsProvider = "Twilio"
	SmsProviderVonage   SmsProvider = "Vonage"
	SmsProviderPlivo    SmsProvider = "Plivo"
	SmsProviderTelnyx   SmsProvider = "Telnyx"
	SmsProviderBird     SmsProvider = "Bird"
	SmsProviderTelesign SmsProvider = "Telesign"
	SmsProviderSinch    SmsProvider = "Sinch"
	SmsProviderFake     SmsProvider = "Fake"
)

// SchedulerTaskType enum.
type SchedulerTaskType string

const (
	SchedulerTaskTypeEmailCampaign      SchedulerTaskType = "EmailCampaign"
	SchedulerTaskTypePushCampaign       SchedulerTaskType = "PushCampaign"
	SchedulerTaskTypeSmsCampaign        SchedulerTaskType = "SmsCampaign"
	SchedulerTaskTypeCodeFunctionalCall SchedulerTaskType = "CodeFunctionalCall"
	SchedulerTaskTypeWebhookCall        SchedulerTaskType = "WebhookCall"
)

// MarketplaceTransport enum.
type MarketplaceTransport string

const (
	MarketplaceTransportMcp      MarketplaceTransport = "Mcp"
	MarketplaceTransportRest     MarketplaceTransport = "Rest"
	MarketplaceTransportCode     MarketplaceTransport = "Code"
	MarketplaceTransportInternal MarketplaceTransport = "Internal"
	MarketplaceTransportSdk      MarketplaceTransport = "Sdk"
)

// MarketplaceCategory enum.
type MarketplaceCategory string

const (
	MarketplaceCategoryOther         MarketplaceCategory = "Other"
	MarketplaceCategoryCrm           MarketplaceCategory = "Crm"
	MarketplaceCategoryErp           MarketplaceCategory = "Erp"
	MarketplaceCategoryMarketing     MarketplaceCategory = "Marketing"
	MarketplaceCategoryCommunication MarketplaceCategory = "Communication"
	MarketplaceCategoryProductivity  MarketplaceCategory = "Productivity"
	MarketplaceCategoryStorage       MarketplaceCategory = "Storage"
	MarketplaceCategoryAnalytics     MarketplaceCategory = "Analytics"
	MarketplaceCategoryIdentity      MarketplaceCategory = "Identity"
	MarketplaceCategoryPayments      MarketplaceCategory = "Payments"
	MarketplaceCategoryDevTools      MarketplaceCategory = "DevTools"
	MarketplaceCategoryAi            MarketplaceCategory = "Ai"
	MarketplaceCategoryFiles         MarketplaceCategory = "Files"
	MarketplaceCategoryDatabase      MarketplaceCategory = "Database"
	MarketplaceCategoryCalendar      MarketplaceCategory = "Calendar"
)

// MarketplaceTokenResolverKind enum.
type MarketplaceTokenResolverKind string

const (
	MarketplaceTokenResolverKindStatic            MarketplaceTokenResolverKind = "Static"
	MarketplaceTokenResolverKindRequest           MarketplaceTokenResolverKind = "Request"
	MarketplaceTokenResolverKindProject           MarketplaceTokenResolverKind = "Project"
	MarketplaceTokenResolverKindInitiator         MarketplaceTokenResolverKind = "Initiator"
	MarketplaceTokenResolverKindCustom            MarketplaceTokenResolverKind = "Custom"
	MarketplaceTokenResolverKindIntegrationConfig MarketplaceTokenResolverKind = "IntegrationConfig"
	MarketplaceTokenResolverKindIntegrationSecret MarketplaceTokenResolverKind = "IntegrationSecret"
)

// MarketplaceSecretValueFormat enum.
type MarketplaceSecretValueFormat string

const (
	MarketplaceSecretValueFormatRaw      MarketplaceSecretValueFormat = "Raw"
	MarketplaceSecretValueFormatBearer   MarketplaceSecretValueFormat = "Bearer"
	MarketplaceSecretValueFormatBasic    MarketplaceSecretValueFormat = "Basic"
	MarketplaceSecretValueFormatPrefixed MarketplaceSecretValueFormat = "Prefixed"
)

// MarketplaceFieldType enum.
type MarketplaceFieldType string

const (
	MarketplaceFieldTypeString        MarketplaceFieldType = "String"
	MarketplaceFieldTypeNumber        MarketplaceFieldType = "Number"
	MarketplaceFieldTypeBoolean       MarketplaceFieldType = "Boolean"
	MarketplaceFieldTypeUrl           MarketplaceFieldType = "Url"
	MarketplaceFieldTypeEmail         MarketplaceFieldType = "Email"
	MarketplaceFieldTypeJson          MarketplaceFieldType = "Json"
	MarketplaceFieldTypeMultilineText MarketplaceFieldType = "MultilineText"
)

// MarketplaceParameterLocation enum.
type MarketplaceParameterLocation string

const (
	MarketplaceParameterLocationBody   MarketplaceParameterLocation = "Body"
	MarketplaceParameterLocationHeader MarketplaceParameterLocation = "Header"
	MarketplaceParameterLocationQuery  MarketplaceParameterLocation = "Query"
	MarketplaceParameterLocationPath   MarketplaceParameterLocation = "Path"
)

// SubscriptionType enum.
type SubscriptionType string

const (
	SubscriptionTypeManagedService SubscriptionType = "ManagedService"
	SubscriptionTypeLicense        SubscriptionType = "License"
)

// DeliveryChannel enum.
type DeliveryChannel string

const (
	DeliveryChannelEmail        DeliveryChannel = "Email"
	DeliveryChannelPush         DeliveryChannel = "Push"
	DeliveryChannelSms          DeliveryChannel = "Sms"
	DeliveryChannelWebPush      DeliveryChannel = "WebPush"
	DeliveryChannelInApp        DeliveryChannel = "InApp"
	DeliveryChannelChatBot      DeliveryChannel = "ChatBot"
	DeliveryChannelChatPlatform DeliveryChannel = "ChatPlatform"
)

// TimeUnit enum.
type TimeUnit string

const (
	TimeUnitTicks        TimeUnit = "Ticks"
	TimeUnitMilliseconds TimeUnit = "Milliseconds"
	TimeUnitSeconds      TimeUnit = "Seconds"
	TimeUnitMinutes      TimeUnit = "Minutes"
	TimeUnitHours        TimeUnit = "Hours"
)

// ResourceRefKind enum.
type ResourceRefKind string

const (
	ResourceRefKindContact         ResourceRefKind = "Contact"
	ResourceRefKindDocument        ResourceRefKind = "Document"
	ResourceRefKindFile            ResourceRefKind = "File"
	ResourceRefKindPaymentCustomer ResourceRefKind = "PaymentCustomer"
	ResourceRefKindOrder           ResourceRefKind = "Order"
	ResourceRefKindPayment         ResourceRefKind = "Payment"
	ResourceRefKindProduct         ResourceRefKind = "Product"
	ResourceRefKindIntegration     ResourceRefKind = "Integration"
)

// ResourceSource enum.
type ResourceSource string

const (
	ResourceSourceNorbix       ResourceSource = "Norbix"
	ResourceSourceStripe       ResourceSource = "Stripe"
	ResourceSourceShopify      ResourceSource = "Shopify"
	ResourceSourcePayPal       ResourceSource = "PayPal"
	ResourceSourceAdyen        ResourceSource = "Adyen"
	ResourceSourceMollie       ResourceSource = "Mollie"
	ResourceSourcePaddle       ResourceSource = "Paddle"
	ResourceSourceLemonSqueezy ResourceSource = "LemonSqueezy"
	ResourceSourceAppleInApp   ResourceSource = "AppleInApp"
	ResourceSourceGoogleInApp  ResourceSource = "GoogleInApp"
	ResourceSourceAuthorizeNet ResourceSource = "AuthorizeNet"
	ResourceSourceBraintree    ResourceSource = "Braintree"
	ResourceSourceCheckOutCom  ResourceSource = "CheckOutCom"
	ResourceSourceWooCommerce  ResourceSource = "WooCommerce"
	ResourceSourceMagento      ResourceSource = "Magento"
	ResourceSourceWorldpay     ResourceSource = "Worldpay"
)

// Continent enum.
type Continent string

const (
	ContinentAfrica       Continent = "Africa"
	ContinentAntarctica   Continent = "Antarctica"
	ContinentAsia         Continent = "Asia"
	ContinentEurope       Continent = "Europe"
	ContinentNorthAmerica Continent = "NorthAmerica"
	ContinentOceania      Continent = "Oceania"
	ContinentSouthAmerica Continent = "SouthAmerica"
)

// PermissionEffect enum.
type PermissionEffect string

const (
	PermissionEffectAllow PermissionEffect = "Allow"
	PermissionEffectDeny  PermissionEffect = "Deny"
)

// ApplicationModule enum.
type ApplicationModule string

const (
	ApplicationModuleAccount      ApplicationModule = "Account"
	ApplicationModuleMembership   ApplicationModule = "Membership"
	ApplicationModuleDatabase     ApplicationModule = "Database"
	ApplicationModuleFiles        ApplicationModule = "Files"
	ApplicationModuleCode         ApplicationModule = "Code"
	ApplicationModuleEmail        ApplicationModule = "Email"
	ApplicationModulePush         ApplicationModule = "Push"
	ApplicationModulePayment      ApplicationModule = "Payment"
	ApplicationModuleScheduler    ApplicationModule = "Scheduler"
	ApplicationModuleLogging      ApplicationModule = "Logging"
	ApplicationModuleServerEvents ApplicationModule = "ServerEvents"
	ApplicationModuleAi           ApplicationModule = "Ai"
	ApplicationModuleSms          ApplicationModule = "Sms"
	ApplicationModuleProject      ApplicationModule = "Project"
	ApplicationModuleCompliance   ApplicationModule = "Compliance"
	ApplicationModuleContacts     ApplicationModule = "Contacts"
	ApplicationModuleMarketplace  ApplicationModule = "Marketplace"
)

// UsageIngestionFailureReason enum.
type UsageIngestionFailureReason string

const (
	UsageIngestionFailureReasonUnknownCustomer  UsageIngestionFailureReason = "UnknownCustomer"
	UsageIngestionFailureReasonMeterNotFound    UsageIngestionFailureReason = "MeterNotFound"
	UsageIngestionFailureReasonValidationFailed UsageIngestionFailureReason = "ValidationFailed"
	UsageIngestionFailureReasonImportSetFailed  UsageIngestionFailureReason = "ImportSetFailed"
)

// ProjectStatus enum.
type ProjectStatus string

const (
	ProjectStatusActive             ProjectStatus = "Active"
	ProjectStatusProvisioning       ProjectStatus = "Provisioning"
	ProjectStatusProvisioningFailed ProjectStatus = "ProvisioningFailed"
	ProjectStatusNoDatabase         ProjectStatus = "NoDatabase"
	ProjectStatusDisabled           ProjectStatus = "Disabled"
	ProjectStatusSuspended          ProjectStatus = "Suspended"
	ProjectStatusRemoved            ProjectStatus = "Removed"
)

// AuthType enum.
type AuthType string

const (
	AuthTypeService  AuthType = "Service"
	AuthTypeEmail    AuthType = "Email"
	AuthTypeUserName AuthType = "UserName"
	AuthTypePhone    AuthType = "Phone"
	AuthTypeGuest    AuthType = "Guest"
	AuthTypeSocial   AuthType = "Social"
)

// Gender enum.
type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
	GenderOther  Gender = "Other"
)

// MarketingBlockReason enum.
type MarketingBlockReason string

const (
	MarketingBlockReasonUnspecified  MarketingBlockReason = "Unspecified"
	MarketingBlockReasonUnsubscribed MarketingBlockReason = "Unsubscribed"
	MarketingBlockReasonComplaint    MarketingBlockReason = "Complaint"
	MarketingBlockReasonHardBounce   MarketingBlockReason = "HardBounce"
	MarketingBlockReasonInvalidEmail MarketingBlockReason = "InvalidEmail"
	MarketingBlockReasonAdminBlock   MarketingBlockReason = "AdminBlock"
)

// AuthStatus enum.
type AuthStatus string

const (
	AuthStatusRegistered        AuthStatus = "Registered"
	AuthStatusPendingValidation AuthStatus = "PendingValidation"
	AuthStatusActive            AuthStatus = "Active"
	AuthStatusUnregistered      AuthStatus = "Unregistered"
	AuthStatusSuspended         AuthStatus = "Suspended"
	AuthStatusInActive          AuthStatus = "InActive"
	AuthStatusBlocked           AuthStatus = "Blocked"
)

// DeviceType enum.
type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "Unknown"
	DeviceTypePhone   DeviceType = "Phone"
	DeviceTypeTablet  DeviceType = "Tablet"
	DeviceTypeDesktop DeviceType = "Desktop"
	DeviceTypeTv      DeviceType = "Tv"
)

// PushDeviceDeliveryFamily enum.
type PushDeviceDeliveryFamily string

const (
	PushDeviceDeliveryFamilyIos     PushDeviceDeliveryFamily = "Ios"
	PushDeviceDeliveryFamilyAndroid PushDeviceDeliveryFamily = "Android"
	PushDeviceDeliveryFamilyChrome  PushDeviceDeliveryFamily = "Chrome"
	PushDeviceDeliveryFamilySafari  PushDeviceDeliveryFamily = "Safari"
	PushDeviceDeliveryFamilyExpo    PushDeviceDeliveryFamily = "Expo"
)

// EmailValidationProvider enum.
type EmailValidationProvider string

const (
	EmailValidationProviderZeroBounce      EmailValidationProvider = "ZeroBounce"
	EmailValidationProviderNeverBounce     EmailValidationProvider = "NeverBounce"
	EmailValidationProviderBouncer         EmailValidationProvider = "Bouncer"
	EmailValidationProviderMailgunValidate EmailValidationProvider = "MailgunValidate"
)

// CampaignStopReason enum.
type CampaignStopReason string

const (
	CampaignStopReasonUserRequested  CampaignStopReason = "UserRequested"
	CampaignStopReasonModuleDisabled CampaignStopReason = "ModuleDisabled"
)

// EmailDeliveryEventType enum.
type EmailDeliveryEventType string

const (
	EmailDeliveryEventTypeUnknown      EmailDeliveryEventType = "Unknown"
	EmailDeliveryEventTypeDelivered    EmailDeliveryEventType = "Delivered"
	EmailDeliveryEventTypeOpen         EmailDeliveryEventType = "Open"
	EmailDeliveryEventTypeClick        EmailDeliveryEventType = "Click"
	EmailDeliveryEventTypeSoftBounce   EmailDeliveryEventType = "SoftBounce"
	EmailDeliveryEventTypeHardBounce   EmailDeliveryEventType = "HardBounce"
	EmailDeliveryEventTypeComplaint    EmailDeliveryEventType = "Complaint"
	EmailDeliveryEventTypeUnsubscribed EmailDeliveryEventType = "Unsubscribed"
)

// MarketplaceIntegrationTransport enum.
type MarketplaceIntegrationTransport string

const (
	MarketplaceIntegrationTransportMcp      MarketplaceIntegrationTransport = "Mcp"
	MarketplaceIntegrationTransportRest     MarketplaceIntegrationTransport = "Rest"
	MarketplaceIntegrationTransportCode     MarketplaceIntegrationTransport = "Code"
	MarketplaceIntegrationTransportInternal MarketplaceIntegrationTransport = "Internal"
	MarketplaceIntegrationTransportSdk      MarketplaceIntegrationTransport = "Sdk"
)

// MarketplaceIntegrationCategory enum.
type MarketplaceIntegrationCategory string

const (
	MarketplaceIntegrationCategoryOther         MarketplaceIntegrationCategory = "Other"
	MarketplaceIntegrationCategoryCrm           MarketplaceIntegrationCategory = "Crm"
	MarketplaceIntegrationCategoryErp           MarketplaceIntegrationCategory = "Erp"
	MarketplaceIntegrationCategoryMarketing     MarketplaceIntegrationCategory = "Marketing"
	MarketplaceIntegrationCategoryCommunication MarketplaceIntegrationCategory = "Communication"
	MarketplaceIntegrationCategoryProductivity  MarketplaceIntegrationCategory = "Productivity"
	MarketplaceIntegrationCategoryStorage       MarketplaceIntegrationCategory = "Storage"
	MarketplaceIntegrationCategoryAnalytics     MarketplaceIntegrationCategory = "Analytics"
	MarketplaceIntegrationCategoryIdentity      MarketplaceIntegrationCategory = "Identity"
	MarketplaceIntegrationCategoryPayments      MarketplaceIntegrationCategory = "Payments"
	MarketplaceIntegrationCategoryDevTools      MarketplaceIntegrationCategory = "DevTools"
	MarketplaceIntegrationCategoryAi            MarketplaceIntegrationCategory = "Ai"
	MarketplaceIntegrationCategoryFiles         MarketplaceIntegrationCategory = "Files"
	MarketplaceIntegrationCategoryDatabase      MarketplaceIntegrationCategory = "Database"
	MarketplaceIntegrationCategoryCalendar      MarketplaceIntegrationCategory = "Calendar"
)

// MarketplaceTokenResolver enum.
type MarketplaceTokenResolver string

const (
	MarketplaceTokenResolverStatic            MarketplaceTokenResolver = "Static"
	MarketplaceTokenResolverRequest           MarketplaceTokenResolver = "Request"
	MarketplaceTokenResolverProject           MarketplaceTokenResolver = "Project"
	MarketplaceTokenResolverInitiator         MarketplaceTokenResolver = "Initiator"
	MarketplaceTokenResolverCustom            MarketplaceTokenResolver = "Custom"
	MarketplaceTokenResolverIntegrationConfig MarketplaceTokenResolver = "IntegrationConfig"
	MarketplaceTokenResolverIntegrationSecret MarketplaceTokenResolver = "IntegrationSecret"
)

// SecretValueFormat enum.
type SecretValueFormat string

const (
	SecretValueFormatRaw      SecretValueFormat = "Raw"
	SecretValueFormatBearer   SecretValueFormat = "Bearer"
	SecretValueFormatBasic    SecretValueFormat = "Basic"
	SecretValueFormatPrefixed SecretValueFormat = "Prefixed"
)

// ResourceKindDto enum.
type ResourceKindDto string

const (
	ResourceKindDtoContact         ResourceKindDto = "contact"
	ResourceKindDtoDocument        ResourceKindDto = "document"
	ResourceKindDtoFile            ResourceKindDto = "file"
	ResourceKindDtoPaymentCustomer ResourceKindDto = "paymentCustomer"
	ResourceKindDtoOrder           ResourceKindDto = "order"
	ResourceKindDtoPayment         ResourceKindDto = "payment"
	ResourceKindDtoProduct         ResourceKindDto = "product"
	ResourceKindDtoIntegration     ResourceKindDto = "integration"
)

// CaseResolutionFixKind enum.
type CaseResolutionFixKind string

const (
	CaseResolutionFixKindCodeFix             CaseResolutionFixKind = "CodeFix"
	CaseResolutionFixKindConfigChange        CaseResolutionFixKind = "ConfigChange"
	CaseResolutionFixKindCustomerInstruction CaseResolutionFixKind = "CustomerInstruction"
	CaseResolutionFixKindKnownLimitation     CaseResolutionFixKind = "KnownLimitation"
	CaseResolutionFixKindDuplicate           CaseResolutionFixKind = "Duplicate"
)

// SupportCaseKind enum.
type SupportCaseKind string

const (
	SupportCaseKindQuestion       SupportCaseKind = "Question"
	SupportCaseKindBug            SupportCaseKind = "Bug"
	SupportCaseKindIncident       SupportCaseKind = "Incident"
	SupportCaseKindBilling        SupportCaseKind = "Billing"
	SupportCaseKindSecurity       SupportCaseKind = "Security"
	SupportCaseKindFeatureRequest SupportCaseKind = "FeatureRequest"
)

// SupportCaseSeverity enum.
type SupportCaseSeverity string

const (
	SupportCaseSeverityS1 SupportCaseSeverity = "S1"
	SupportCaseSeverityS2 SupportCaseSeverity = "S2"
	SupportCaseSeverityS3 SupportCaseSeverity = "S3"
	SupportCaseSeverityS4 SupportCaseSeverity = "S4"
)

// DeploymentMode enum.
type DeploymentMode string

const (
	DeploymentModeManaged    DeploymentMode = "Managed"
	DeploymentModeSelfHosted DeploymentMode = "SelfHosted"
	DeploymentModeEnterprise DeploymentMode = "Enterprise"
)

// SupportMessageAuthorKind enum.
type SupportMessageAuthorKind string

const (
	SupportMessageAuthorKindCustomer SupportMessageAuthorKind = "Customer"
	SupportMessageAuthorKindStaff    SupportMessageAuthorKind = "Staff"
	SupportMessageAuthorKindAi       SupportMessageAuthorKind = "Ai"
	SupportMessageAuthorKindSystem   SupportMessageAuthorKind = "System"
)

// SupportCaseStatus enum.
type SupportCaseStatus string

const (
	SupportCaseStatusOpen              SupportCaseStatus = "Open"
	SupportCaseStatusTriaged           SupportCaseStatus = "Triaged"
	SupportCaseStatusInProgress        SupportCaseStatus = "InProgress"
	SupportCaseStatusWaitingOnCustomer SupportCaseStatus = "WaitingOnCustomer"
	SupportCaseStatusResolved          SupportCaseStatus = "Resolved"
	SupportCaseStatusClosed            SupportCaseStatus = "Closed"
)

// SupportCaseCloseReason enum.
type SupportCaseCloseReason string

const (
	SupportCaseCloseReasonManual                 SupportCaseCloseReason = "Manual"
	SupportCaseCloseReasonAutoClosedAfterResolve SupportCaseCloseReason = "AutoClosedAfterResolve"
)

// CodeMashRelease enum.
type CodeMashRelease string

const (
	CodeMashReleaseNotSet         CodeMashRelease = "NotSet"
	CodeMashReleaseCommunity      CodeMashRelease = "Community"
	CodeMashReleaseManagedService CodeMashRelease = "ManagedService"
	CodeMashReleaseEnterprise     CodeMashRelease = "Enterprise"
)

// CodeMashRuntime enum.
type CodeMashRuntime string

const (
	CodeMashRuntimeDevelopment CodeMashRuntime = "Development"
	CodeMashRuntimeCI          CodeMashRuntime = "CI"
	CodeMashRuntimeStaging     CodeMashRuntime = "Staging"
	CodeMashRuntimeProduction  CodeMashRuntime = "Production"
)

// AccountStatus enum.
type AccountStatus string

const (
	AccountStatusRegistered        AccountStatus = "Registered"
	AccountStatusPendingValidation AccountStatus = "PendingValidation"
	AccountStatusActive            AccountStatus = "Active"
	AccountStatusInActive          AccountStatus = "InActive"
	AccountStatusBlocked           AccountStatus = "Blocked"
	AccountStatusUnregistered      AccountStatus = "Unregistered"
)

// CampaignStatus enum.
type CampaignStatus string

const (
	CampaignStatusPending    CampaignStatus = "Pending"
	CampaignStatusRegistered CampaignStatus = "Registered"
	CampaignStatusScheduled  CampaignStatus = "Scheduled"
	CampaignStatusStarted    CampaignStatus = "Started"
	CampaignStatusStopped    CampaignStatus = "Stopped"
	CampaignStatusProcessing CampaignStatus = "Processing"
	CampaignStatusCompleted  CampaignStatus = "Completed"
	CampaignStatusFailed     CampaignStatus = "Failed"
)

// CampaignBatchStatus enum.
type CampaignBatchStatus string

const (
	CampaignBatchStatusRegistered CampaignBatchStatus = "Registered"
	CampaignBatchStatusProcessing CampaignBatchStatus = "Processing"
	CampaignBatchStatusCompleted  CampaignBatchStatus = "Completed"
	CampaignBatchStatusFailed     CampaignBatchStatus = "Failed"
)

// CampaignNotificationStatus enum.
type CampaignNotificationStatus string

const (
	CampaignNotificationStatusCompleted                         CampaignNotificationStatus = "Completed"
	CampaignNotificationStatusBlockedByUserPreferenceBlockAll   CampaignNotificationStatus = "BlockedByUserPreferenceBlockAll"
	CampaignNotificationStatusBlockedByUserPreferenceBlockByTag CampaignNotificationStatus = "BlockedByUserPreferenceBlockByTag"
	CampaignNotificationStatusFailed                            CampaignNotificationStatus = "Failed"
	CampaignNotificationStatusViewed                            CampaignNotificationStatus = "Viewed"
	CampaignNotificationStatusClicked                           CampaignNotificationStatus = "Clicked"
	CampaignNotificationStatusBlockedByValidation               CampaignNotificationStatus = "BlockedByValidation"
)

// CacheControl enum.
type CacheControl string

const (
	CacheControlNone            CacheControl = "None"
	CacheControlPublic          CacheControl = "Public"
	CacheControlPrivate         CacheControl = "Private"
	CacheControlMustRevalidate  CacheControl = "MustRevalidate"
	CacheControlNoCache         CacheControl = "NoCache"
	CacheControlNoStore         CacheControl = "NoStore"
	CacheControlNoTransform     CacheControl = "NoTransform"
	CacheControlProxyRevalidate CacheControl = "ProxyRevalidate"
)

// RequestAttributes enum.
type RequestAttributes string

const (
	RequestAttributesNone                  RequestAttributes = "None"
	RequestAttributesLocalhost             RequestAttributes = "Localhost"
	RequestAttributesLocalSubnet           RequestAttributes = "LocalSubnet"
	RequestAttributesExternal              RequestAttributes = "External"
	RequestAttributesSecure                RequestAttributes = "Secure"
	RequestAttributesInSecure              RequestAttributes = "InSecure"
	RequestAttributesAnySecurityMode       RequestAttributes = "AnySecurityMode"
	RequestAttributesHttpHead              RequestAttributes = "HttpHead"
	RequestAttributesHttpGet               RequestAttributes = "HttpGet"
	RequestAttributesHttpPost              RequestAttributes = "HttpPost"
	RequestAttributesHttpPut               RequestAttributes = "HttpPut"
	RequestAttributesHttpDelete            RequestAttributes = "HttpDelete"
	RequestAttributesHttpPatch             RequestAttributes = "HttpPatch"
	RequestAttributesHttpOptions           RequestAttributes = "HttpOptions"
	RequestAttributesHttpOther             RequestAttributes = "HttpOther"
	RequestAttributesAnyHttpMethod         RequestAttributes = "AnyHttpMethod"
	RequestAttributesOneWay                RequestAttributes = "OneWay"
	RequestAttributesReply                 RequestAttributes = "Reply"
	RequestAttributesAnyCallStyle          RequestAttributes = "AnyCallStyle"
	RequestAttributesSoap11                RequestAttributes = "Soap11"
	RequestAttributesSoap12                RequestAttributes = "Soap12"
	RequestAttributesXml                   RequestAttributes = "Xml"
	RequestAttributesJson                  RequestAttributes = "Json"
	RequestAttributesJsv                   RequestAttributes = "Jsv"
	RequestAttributesProtoBuf              RequestAttributes = "ProtoBuf"
	RequestAttributesCsv                   RequestAttributes = "Csv"
	RequestAttributesHtml                  RequestAttributes = "Html"
	RequestAttributesJsonl                 RequestAttributes = "Jsonl"
	RequestAttributesMsgPack               RequestAttributes = "MsgPack"
	RequestAttributesFormatOther           RequestAttributes = "FormatOther"
	RequestAttributesAnyFormat             RequestAttributes = "AnyFormat"
	RequestAttributesHttp                  RequestAttributes = "Http"
	RequestAttributesMessageQueue          RequestAttributes = "MessageQueue"
	RequestAttributesTcp                   RequestAttributes = "Tcp"
	RequestAttributesGrpc                  RequestAttributes = "Grpc"
	RequestAttributesEndpointOther         RequestAttributes = "EndpointOther"
	RequestAttributesAnyEndpoint           RequestAttributes = "AnyEndpoint"
	RequestAttributesInProcess             RequestAttributes = "InProcess"
	RequestAttributesInternalNetworkAccess RequestAttributes = "InternalNetworkAccess"
	RequestAttributesAnyNetworkAccessType  RequestAttributes = "AnyNetworkAccessType"
	RequestAttributesAny                   RequestAttributes = "Any"
)

// ResolvedRefStatus enum.
type ResolvedRefStatus string

const (
	ResolvedRefStatusOk           ResolvedRefStatus = "ok"
	ResolvedRefStatusNotFound     ResolvedRefStatus = "notFound"
	ResolvedRefStatusUnauthorized ResolvedRefStatus = "unauthorized"
	ResolvedRefStatusSourceError  ResolvedRefStatus = "sourceError"
	ResolvedRefStatusErased       ResolvedRefStatus = "erased"
)

// SupportCustomerStatus enum.
type SupportCustomerStatus string

const (
	SupportCustomerStatusPending SupportCustomerStatus = "Pending"
	SupportCustomerStatusOpen    SupportCustomerStatus = "Open"
	SupportCustomerStatusSolved  SupportCustomerStatus = "Solved"
	SupportCustomerStatusClosed  SupportCustomerStatus = "Closed"
)

// IReturn DTO.
type IReturn[T any] struct {
}

// IReturnVoid DTO.
type IReturnVoid struct {
}

// IHasSessionId DTO.
type IHasSessionId struct {
}

// IHasBearerToken DTO.
type IHasBearerToken struct {
}

// IPost DTO.
type IPost struct {
}

// EmailIntegrationRequest DTO.
type EmailIntegrationRequest struct {
	IntegrationId   string        `json:"integrationId,omitempty"`
	Provider        EmailProvider `json:"provider,omitempty"`
	IntegrationName string        `json:"integrationName,omitempty"`
	IsEnabled       bool          `json:"isEnabled,omitempty"`
	EmailAddress    string        `json:"emailAddress,omitempty"`
	EmailSenderName string        `json:"emailSenderName,omitempty"`
}

// SmtpEmailIntegrationRequest DTO.
type SmtpEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider EmailProvider `json:"provider,omitempty"`
	Domain   string        `json:"domain,omitempty"`
	Port     SmtpPorts     `json:"port,omitempty"`
	UserName string        `json:"userName,omitempty"`
	Password string        `json:"password,omitempty"`
}

// AwsSesEmailIntegrationRequest DTO.
type AwsSesEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider         EmailProvider      `json:"provider,omitempty"`
	IntegrationType  AwsIntegrationType `json:"integrationType,omitempty"`
	AwsRegion        string             `json:"awsRegion,omitempty"`
	EmailIdentityArn string             `json:"emailIdentityArn,omitempty"`
	ConfigurationSet string             `json:"configurationSet,omitempty"`
	RoleArn          string             `json:"roleArn,omitempty"`
	ExternalId       string             `json:"externalId,omitempty"`
	AccessKey        string             `json:"accessKey,omitempty"`
	SecretKey        string             `json:"secretKey,omitempty"`
}

// SendGridEmailIntegrationRequest DTO.
type SendGridEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider EmailProvider `json:"provider,omitempty"`
	ApiKey   string        `json:"apiKey,omitempty"`
}

// MailGunEmailIntegrationRequest DTO.
type MailGunEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider          EmailProvider `json:"provider,omitempty"`
	Domain            string        `json:"domain,omitempty"`
	ApiKey            string        `json:"apiKey,omitempty"`
	WebhookSigningKey string        `json:"webhookSigningKey,omitempty"`
	Region            MailGunRegion `json:"region,omitempty"`
}

// TokenMappingDto DTO.
type TokenMappingDto struct {
	Key      string                   `json:"key,omitempty"`
	Value    string                   `json:"value,omitempty"`
	Resolver TokenMappingResolverType `json:"resolver,omitempty"`
}

// EmailCampaignRequest DTO.
type EmailCampaignRequest struct {
	Source                  EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	TemplateId              string                             `json:"templateId,omitempty"`
	IntegrationId           string                             `json:"integrationId,omitempty"`
	ValidationIntegrationId string                             `json:"validationIntegrationId,omitempty"`
	Language                string                             `json:"language,omitempty"`
	InitiatorId             string                             `json:"initiatorId,omitempty"`
	Notes                   string                             `json:"notes,omitempty"`
	MappedTokens            []*TokenMappingDto                 `json:"mappedTokens,omitempty"`
	CampaignTime            float64                            `json:"campaignTime,omitempty"`
}

// EmailToAllUsersDeliverySettingsRequest DTO.
type EmailToAllUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source     EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	RolesNames []string                           `json:"rolesNames,omitempty"`
	UserTags   []string                           `json:"userTags,omitempty"`
}

// EmailToAccountUsersDeliverySettingsRequest DTO.
type EmailToAccountUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	UserRecipients      []string                           `json:"userRecipients,omitempty"`
	UserCc              []string                           `json:"userCc,omitempty"`
	UserBcc             []string                           `json:"userBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy,omitempty"`
}

// EmailToCollectionRecordsDeliverySettingsRequest DTO.
type EmailToCollectionRecordsDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source     EmailCampaignRecipientsSourceTypes    `json:"source,omitempty"`
	Fields     []string                              `json:"fields,omitempty"`
	SchemaName string                                `json:"schemaName,omitempty"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
}

// EmailToEmailsDeliverySettingsRequest DTO.
type EmailToEmailsDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	Recipients          []string                           `json:"recipients,omitempty"`
	RecipientsCc        []string                           `json:"recipientsCc,omitempty"`
	RecipientsBcc       []string                           `json:"recipientsBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy,omitempty"`
}

// EmailToUsersDeliverySettingsRequest DTO.
type EmailToUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	UserRecipients      []string                           `json:"userRecipients,omitempty"`
	UserCc              []string                           `json:"userCc,omitempty"`
	UserBcc             []string                           `json:"userBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy,omitempty"`
}

// TriggerActionDto DTO.
type TriggerActionDto struct {
	Type          TriggerActionType `json:"type,omitempty"`
	IntegrationId string            `json:"integrationId,omitempty"`
}

// SaveTriggerRequest DTO.
type SaveTriggerRequest struct {
	Type           TriggerType       `json:"type,omitempty"`
	TriggerId      string            `json:"triggerId,omitempty"`
	Name           string            `json:"name,omitempty"`
	Description    string            `json:"description,omitempty"`
	IsEnabled      bool              `json:"isEnabled,omitempty"`
	PreExecuteCode string            `json:"preExecuteCode,omitempty"`
	Action         *TriggerActionDto `json:"action,omitempty"`
}

// MembershipTriggerRequest DTO.
type MembershipTriggerRequest struct {
	SaveTriggerRequest
	Type TriggerType           `json:"type,omitempty"`
	When MembershipTriggerType `json:"when,omitempty"`
}

// SchemaTriggerRequest DTO.
type SchemaTriggerRequest struct {
	SaveTriggerRequest
	Type              TriggerType       `json:"type,omitempty"`
	SchemaId          string            `json:"schemaId,omitempty"`
	When              SchemaTriggerType `json:"when,omitempty"`
	ConfigurationCode string            `json:"configurationCode,omitempty"`
}

// FileChecksumDto DTO.
type FileChecksumDto struct {
	Algorithm string `json:"algorithm,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

// FileResourceDto DTO.
type FileResourceDto struct {
	Id               string           `json:"id,omitempty"`
	OriginalFileName string           `json:"originalFileName,omitempty"`
	Extension        string           `json:"extension,omitempty"`
	StoredFileName   string           `json:"storedFileName,omitempty"`
	SizeBytes        float64          `json:"sizeBytes,omitempty"`
	Checksum         *FileChecksumDto `json:"checksum,omitempty"`
}

// FileResourceRefDto DTO.
type FileResourceRefDto struct {
	Resource      *FileResourceDto `json:"resource,omitempty"`
	IntegrationId string           `json:"integrationId,omitempty"`
	Provider      FileProvider     `json:"provider,omitempty"`
	Path          string           `json:"path,omitempty"`
	PublicUrl     string           `json:"publicUrl,omitempty"`
	IsPublic      bool             `json:"isPublic,omitempty"`
}

// FilesTriggerRequest DTO.
type FilesTriggerRequest struct {
	SaveTriggerRequest
	Type   TriggerType      `json:"type,omitempty"`
	When   FilesTriggerType `json:"when,omitempty"`
	Folder string           `json:"folder,omitempty"`
}

// PaymentTriggerRequest DTO.
type PaymentTriggerRequest struct {
	SaveTriggerRequest
	Type         TriggerType        `json:"type,omitempty"`
	When         PaymentTriggerType `json:"when,omitempty"`
	Integrations []string           `json:"integrations,omitempty"`
	Events       []string           `json:"events,omitempty"`
}

// DatabaseIntegrationRequest DTO.
type DatabaseIntegrationRequest struct {
	IntegrationId   string           `json:"integrationId,omitempty"`
	Provider        DatabaseProvider `json:"provider,omitempty"`
	IntegrationName string           `json:"integrationName,omitempty"`
	IsEnabled       bool             `json:"isEnabled,omitempty"`
}

// MongoDbConnectionStringDatabaseIntegrationRequest DTO.
type MongoDbConnectionStringDatabaseIntegrationRequest struct {
	DatabaseIntegrationRequest
	Provider         DatabaseProvider `json:"provider,omitempty"`
	DatabaseName     string           `json:"databaseName,omitempty"`
	ConnectionString string           `json:"connectionString,omitempty"`
}

// MongoDbAtlasFlexManagedDatabaseIntegrationRequest DTO.
type MongoDbAtlasFlexManagedDatabaseIntegrationRequest struct {
	DatabaseIntegrationRequest
	Provider         DatabaseProvider `json:"provider,omitempty"`
	NorbixRegionCode string           `json:"norbixRegionCode,omitempty"`
}

// FilesIntegrationRequest DTO.
type FilesIntegrationRequest struct {
	IntegrationId   string       `json:"integrationId,omitempty"`
	Provider        FileProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName,omitempty"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
}

// GoogleDriveFilesIntegrationRequest DTO.
type GoogleDriveFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider              FileProvider `json:"provider,omitempty"`
	RootFolderId          string       `json:"rootFolderId,omitempty"`
	ServiceAccountJsonKey string       `json:"serviceAccountJsonKey,omitempty"`
}

// FtpFilesIntegrationRequest DTO.
type FtpFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider FileProvider `json:"provider,omitempty"`
	Host     string       `json:"host,omitempty"`
	Port     float64      `json:"port,omitempty"`
	RootPath string       `json:"rootPath,omitempty"`
	UseSsl   bool         `json:"useSsl,omitempty"`
	Username string       `json:"username,omitempty"`
	Password string       `json:"password,omitempty"`
}

// DropBoxFilesIntegrationRequest DTO.
type DropBoxFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider    FileProvider `json:"provider,omitempty"`
	RootPath    string       `json:"rootPath,omitempty"`
	AccessToken string       `json:"accessToken,omitempty"`
}

// AppleICloudFilesIntegrationRequest DTO.
type AppleICloudFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider            FileProvider `json:"provider,omitempty"`
	ContainerIdentifier string       `json:"containerIdentifier,omitempty"`
	RelativePath        string       `json:"relativePath,omitempty"`
	KeyId               string       `json:"keyId,omitempty"`
	TeamId              string       `json:"teamId,omitempty"`
	BundleId            string       `json:"bundleId,omitempty"`
	P8PrivateKey        string       `json:"p8PrivateKey,omitempty"`
}

// AwsS3FilesIntegrationRequest DTO.
type AwsS3FilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider        FileProvider         `json:"provider,omitempty"`
	IntegrationType AwsS3IntegrationType `json:"integrationType,omitempty"`
	BucketName      string               `json:"bucketName,omitempty"`
	Region          string               `json:"region,omitempty"`
	RoleArn         string               `json:"roleArn,omitempty"`
	ExternalId      string               `json:"externalId,omitempty"`
	AccessKey       string               `json:"accessKey,omitempty"`
	SecretKey       string               `json:"secretKey,omitempty"`
}

// GoogleCloudFilesIntegrationRequest DTO.
type GoogleCloudFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider              FileProvider `json:"provider,omitempty"`
	BucketName            string       `json:"bucketName,omitempty"`
	ServiceAccountJsonKey string       `json:"serviceAccountJsonKey,omitempty"`
}

// AzureBlobFilesIntegrationRequest DTO.
type AzureBlobFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider         FileProvider `json:"provider,omitempty"`
	BlobName         string       `json:"blobName,omitempty"`
	ConnectionString string       `json:"connectionString,omitempty"`
}

// LocalFilesIntegrationRequest DTO.
type LocalFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider FileProvider `json:"provider,omitempty"`
	RootPath string       `json:"rootPath,omitempty"`
}

// LoggingIntegrationRequest DTO.
type LoggingIntegrationRequest struct {
	IntegrationId   string          `json:"integrationId,omitempty"`
	Provider        LoggingProvider `json:"provider,omitempty"`
	IntegrationName string          `json:"integrationName,omitempty"`
	IsEnabled       bool            `json:"isEnabled,omitempty"`
}

// AmqpLoggingIntegrationRequest DTO.
type AmqpLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	Host        string          `json:"host,omitempty"`
	Port        float64         `json:"port,omitempty"`
	VirtualHost string          `json:"virtualHost,omitempty"`
	Exchange    string          `json:"exchange,omitempty"`
	RoutingKey  string          `json:"routingKey,omitempty"`
	Username    string          `json:"username,omitempty"`
	Password    string          `json:"password,omitempty"`
}

// AwsKinesisLoggingIntegrationRequest DTO.
type AwsKinesisLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider   LoggingProvider `json:"provider,omitempty"`
	StreamName string          `json:"streamName,omitempty"`
	Region     string          `json:"region,omitempty"`
	AccessKey  string          `json:"accessKey,omitempty"`
	SecretKey  string          `json:"secretKey,omitempty"`
}

// AwsS3LoggingIntegrationRequest DTO.
type AwsS3LoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider        LoggingProvider             `json:"provider,omitempty"`
	IntegrationType AwsS3LoggingIntegrationType `json:"integrationType,omitempty"`
	BucketName      string                      `json:"bucketName,omitempty"`
	Region          string                      `json:"region,omitempty"`
	RoleArn         string                      `json:"roleArn,omitempty"`
	ExternalId      string                      `json:"externalId,omitempty"`
	AccessKey       string                      `json:"accessKey,omitempty"`
	SecretKey       string                      `json:"secretKey,omitempty"`
}

// NewRelicLoggingIntegrationRequest DTO.
type NewRelicLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	Region      string          `json:"region,omitempty"`
	ServiceName string          `json:"serviceName,omitempty"`
	ApiKey      string          `json:"apiKey,omitempty"`
}

// MongoDbLoggingIntegrationRequest DTO.
type MongoDbLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	DatabaseName     string          `json:"databaseName,omitempty"`
	ConnectionString string          `json:"connectionString,omitempty"`
}

// KafkaLoggingIntegrationRequest DTO.
type KafkaLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	BootstrapServers string          `json:"bootstrapServers,omitempty"`
	Topic            string          `json:"topic,omitempty"`
	SecurityProtocol string          `json:"securityProtocol,omitempty"`
	SaslUsername     string          `json:"saslUsername,omitempty"`
	SaslPassword     string          `json:"saslPassword,omitempty"`
}

// PrometheusLoggingIntegrationRequest DTO.
type PrometheusLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	EndpointUrl string          `json:"endpointUrl,omitempty"`
	JobName     string          `json:"jobName,omitempty"`
	BearerToken string          `json:"bearerToken,omitempty"`
}

// DataDogLoggingIntegrationRequest DTO.
type DataDogLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	Site        string          `json:"site,omitempty"`
	ServiceName string          `json:"serviceName,omitempty"`
	Environment string          `json:"environment,omitempty"`
	ApiKey      string          `json:"apiKey,omitempty"`
}

// InternalKafkaLoggingIntegrationRequest DTO.
type InternalKafkaLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	BootstrapServers string          `json:"bootstrapServers,omitempty"`
	Topic            string          `json:"topic,omitempty"`
	SecurityProtocol string          `json:"securityProtocol,omitempty"`
	SaslUsername     string          `json:"saslUsername,omitempty"`
	SaslPassword     string          `json:"saslPassword,omitempty"`
}

// ElasticSearchLoggingIntegrationRequest DTO.
type ElasticSearchLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider LoggingProvider `json:"provider,omitempty"`
	Uri      string          `json:"uri,omitempty"`
	Index    string          `json:"index,omitempty"`
	Username string          `json:"username,omitempty"`
	Password string          `json:"password,omitempty"`
}

// SplunkLoggingIntegrationRequest DTO.
type SplunkLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider       LoggingProvider `json:"provider,omitempty"`
	HecEndpointUrl string          `json:"hecEndpointUrl,omitempty"`
	Index          string          `json:"index,omitempty"`
	HecToken       string          `json:"hecToken,omitempty"`
}

// AzureOtelLoggingIntegrationRequest DTO.
type AzureOtelLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	EndpointUrl      string          `json:"endpointUrl,omitempty"`
	ResourceName     string          `json:"resourceName,omitempty"`
	ConnectionString string          `json:"connectionString,omitempty"`
}

// KibanaLoggingIntegrationRequest DTO.
type KibanaLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider LoggingProvider `json:"provider,omitempty"`
	Uri      string          `json:"uri,omitempty"`
	SpaceId  string          `json:"spaceId,omitempty"`
	ApiKey   string          `json:"apiKey,omitempty"`
}

// LocalFileLoggingIntegrationRequest DTO.
type LocalFileLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider LoggingProvider `json:"provider,omitempty"`
	RootPath string          `json:"rootPath,omitempty"`
}

// MembershipIntegrationRequest DTO.
type MembershipIntegrationRequest struct {
	IntegrationId   string             `json:"integrationId,omitempty"`
	Provider        MembershipProvider `json:"provider,omitempty"`
	IntegrationName string             `json:"integrationName,omitempty"`
	IsEnabled       bool               `json:"isEnabled,omitempty"`
}

// DisplayName DTO.
type DisplayName struct {
	Value string `json:"value,omitempty"`
}

// RoleName DTO.
type RoleName struct {
	Name                string `json:"name,omitempty"`
	DisplayName         string `json:"displayName,omitempty"`
	IsAdministrator     bool   `json:"isAdministrator,omitempty"`
	IsAuthenticated     bool   `json:"isAuthenticated,omitempty"`
	IsGuest             bool   `json:"isGuest,omitempty"`
	IsRootRole          bool   `json:"isRootRole,omitempty"`
	IsCollaboratorRole  bool   `json:"isCollaboratorRole,omitempty"`
	IsProjectSystemRole bool   `json:"isProjectSystemRole,omitempty"`
	IsAccountSystemRole bool   `json:"isAccountSystemRole,omitempty"`
	IsSystemRole        bool   `json:"isSystemRole,omitempty"`
}

// OAuthModeConfig DTO.
type OAuthModeConfig struct {
	Name               *DisplayName `json:"name,omitempty"`
	CallbackUrl        string       `json:"callbackUrl,omitempty"`
	LogoutUrl          string       `json:"logoutUrl,omitempty"`
	FailureRedirectUrl string       `json:"failureRedirectUrl,omitempty"`
	RoleName           *RoleName    `json:"roleName,omitempty"`
}

// OktaMembershipIntegrationRequest DTO.
type OktaMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	Domain       string             `json:"domain,omitempty"`
	ClientId     string             `json:"clientId,omitempty"`
	ClientSecret string             `json:"clientSecret,omitempty"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// XMembershipIntegrationRequest DTO.
type XMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	ApiKey       string             `json:"apiKey,omitempty"`
	ApiSecretKey string             `json:"apiSecretKey,omitempty"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// GoogleMembershipIntegrationRequest DTO.
type GoogleMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	ClientId     string             `json:"clientId,omitempty"`
	ClientSecret string             `json:"clientSecret,omitempty"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// MicrosoftMembershipIntegrationRequest DTO.
type MicrosoftMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	TenantId     string             `json:"tenantId,omitempty"`
	ClientId     string             `json:"clientId,omitempty"`
	ClientSecret string             `json:"clientSecret,omitempty"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// GitHubMembershipIntegrationRequest DTO.
type GitHubMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	ClientId     string             `json:"clientId,omitempty"`
	ClientSecret string             `json:"clientSecret,omitempty"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// MetaMembershipIntegrationRequest DTO.
type MetaMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider   MembershipProvider `json:"provider,omitempty"`
	AppId      string             `json:"appId,omitempty"`
	AppSecret  string             `json:"appSecret,omitempty"`
	OAuthModes []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// AppleMembershipIntegrationRequest DTO.
type AppleMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	TeamId       string             `json:"teamId,omitempty"`
	AppBundleId  string             `json:"appBundleId,omitempty"`
	ServiceId    string             `json:"serviceId,omitempty"`
	KeyId        string             `json:"keyId,omitempty"`
	PrivateKey   string             `json:"privateKey,omitempty"`
	IsProduction bool               `json:"isProduction,omitempty"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// PaymentIntegrationRequest DTO.
type PaymentIntegrationRequest struct {
	IntegrationId   string                 `json:"integrationId,omitempty"`
	Provider        PaymentGatewayPlatform `json:"provider,omitempty"`
	IntegrationName string                 `json:"integrationName,omitempty"`
	IsEnabled       bool                   `json:"isEnabled,omitempty"`
}

// LemonSqueezyPaymentIntegrationRequest DTO.
type LemonSqueezyPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	StoreId              string                 `json:"storeId,omitempty"`
	ApiKey               string                 `json:"apiKey,omitempty"`
	WebhookSigningSecret string                 `json:"webhookSigningSecret,omitempty"`
	IsTestMode           bool                   `json:"isTestMode,omitempty"`
}

// AdyenPaymentIntegrationRequest DTO.
type AdyenPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider        PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantAccount string                 `json:"merchantAccount,omitempty"`
	ApiKey          string                 `json:"apiKey,omitempty"`
	Environment     string                 `json:"environment,omitempty"`
	WebhookId       string                 `json:"webhookId,omitempty"`
	WebhookHmacKey  string                 `json:"webhookHmacKey,omitempty"`
}

// MolliePaymentIntegrationRequest DTO.
type MolliePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	ProfileId            string                 `json:"profileId,omitempty"`
	ApiKey               string                 `json:"apiKey,omitempty"`
	IsTestMode           bool                   `json:"isTestMode,omitempty"`
	WebhookSigningSecret string                 `json:"webhookSigningSecret,omitempty"`
}

// PaddlePaymentIntegrationRequest DTO.
type PaddlePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider                 PaymentGatewayPlatform `json:"provider,omitempty"`
	ApiKey                   string                 `json:"apiKey,omitempty"`
	WebhookEndpointSecretKey string                 `json:"webhookEndpointSecretKey,omitempty"`
	Environment              string                 `json:"environment,omitempty"`
	ClientSideToken          string                 `json:"clientSideToken,omitempty"`
}

// PayPalPaymentIntegrationRequest DTO.
type PayPalPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider     PaymentGatewayPlatform `json:"provider,omitempty"`
	ClientId     string                 `json:"clientId,omitempty"`
	ClientSecret string                 `json:"clientSecret,omitempty"`
	Environment  string                 `json:"environment,omitempty"`
	BrandName    string                 `json:"brandName,omitempty"`
	WebhookId    string                 `json:"webhookId,omitempty"`
}

// StripePaymentIntegrationRequest DTO.
type StripePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	PublishableKey       string                 `json:"publishableKey,omitempty"`
	SecretKey            string                 `json:"secretKey,omitempty"`
	WebhookSigningSecret string                 `json:"webhookSigningSecret,omitempty"`
	WebhookEndpointId    string                 `json:"webhookEndpointId,omitempty"`
	DefaultCurrency      string                 `json:"defaultCurrency,omitempty"`
}

// AppleInAppPaymentIntegrationRequest DTO.
type AppleInAppPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider                              PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantIdentifier                    string                 `json:"merchantIdentifier,omitempty"`
	MerchantDomain                        string                 `json:"merchantDomain,omitempty"`
	DisplayName                           string                 `json:"displayName,omitempty"`
	MerchantIdentityCertificateP12Base64  string                 `json:"merchantIdentityCertificateP12Base64,omitempty"`
	MerchantIdentityCertificatePassword   string                 `json:"merchantIdentityCertificatePassword,omitempty"`
	PaymentProcessingCertificateP12Base64 string                 `json:"paymentProcessingCertificateP12Base64,omitempty"`
	PaymentProcessingCertificatePassword  string                 `json:"paymentProcessingCertificatePassword,omitempty"`
	WebhookBundleId                       string                 `json:"webhookBundleId,omitempty"`
}

// GoogleInAppPaymentIntegrationRequest DTO.
type GoogleInAppPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider           PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantId         string                 `json:"merchantId,omitempty"`
	MerchantName       string                 `json:"merchantName,omitempty"`
	Gateway            string                 `json:"gateway,omitempty"`
	PrivateKeyOrToken  string                 `json:"privateKeyOrToken,omitempty"`
	GatewayMerchantId  string                 `json:"gatewayMerchantId,omitempty"`
	WebhookPackageName string                 `json:"webhookPackageName,omitempty"`
}

// PushIntegrationRequest DTO.
type PushIntegrationRequest struct {
	IntegrationId   string       `json:"integrationId,omitempty"`
	Provider        PushProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName,omitempty"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
}

// EdgeWebPushIntegrationRequest DTO.
type EdgeWebPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey,omitempty"`
	VapidPrivateKey string       `json:"vapidPrivateKey,omitempty"`
	Subject         string       `json:"subject,omitempty"`
}

// ChromePluginPushIntegrationRequest DTO.
type ChromePluginPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	ExtensionId     string       `json:"extensionId,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey,omitempty"`
	VapidPrivateKey string       `json:"vapidPrivateKey,omitempty"`
	Subject         string       `json:"subject,omitempty"`
}

// SafariPushIntegrationRequest DTO.
type SafariPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider             PushProvider `json:"provider,omitempty"`
	WebsitePushId        string       `json:"websitePushId,omitempty"`
	CertificateP12Base64 string       `json:"certificateP12Base64,omitempty"`
	CertificatePassword  string       `json:"certificatePassword,omitempty"`
}

// ChromeWebPushIntegrationRequest DTO.
type ChromeWebPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey,omitempty"`
	VapidPrivateKey string       `json:"vapidPrivateKey,omitempty"`
	Subject         string       `json:"subject,omitempty"`
}

// FirefoxWebPushIntegrationRequest DTO.
type FirefoxWebPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey,omitempty"`
	VapidPrivateKey string       `json:"vapidPrivateKey,omitempty"`
	Subject         string       `json:"subject,omitempty"`
}

// AndroidFirebasePushIntegrationRequest DTO.
type AndroidFirebasePushIntegrationRequest struct {
	PushIntegrationRequest
	Provider           PushProvider `json:"provider,omitempty"`
	ProjectId          string       `json:"projectId,omitempty"`
	ClientEmail        string       `json:"clientEmail,omitempty"`
	ServiceAccountJson string       `json:"serviceAccountJson,omitempty"`
}

// AppleApnsPushIntegrationRequest DTO.
type AppleApnsPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider     PushProvider `json:"provider,omitempty"`
	TeamId       string       `json:"teamId,omitempty"`
	AppBundleId  string       `json:"appBundleId,omitempty"`
	KeyId        string       `json:"keyId,omitempty"`
	PrivateKey   string       `json:"privateKey,omitempty"`
	IsProduction bool         `json:"isProduction,omitempty"`
}

// CodeIntegrationRequest DTO.
type CodeIntegrationRequest struct {
	IntegrationId   string       `json:"integrationId,omitempty"`
	Provider        CodeProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName,omitempty"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
}

// AwsLambdaCodeIntegrationRequest DTO.
type AwsLambdaCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider        CodeProvider             `json:"provider,omitempty"`
	IntegrationType AwsLambdaIntegrationType `json:"integrationType,omitempty"`
	Region          string                   `json:"region,omitempty"`
	RoleArn         string                   `json:"roleArn,omitempty"`
	ExternalId      string                   `json:"externalId,omitempty"`
	AccessKey       string                   `json:"accessKey,omitempty"`
	SecretKey       string                   `json:"secretKey,omitempty"`
}

// AzureFunctionsCodeIntegrationRequest DTO.
type AzureFunctionsCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider              CodeProvider `json:"provider,omitempty"`
	FunctionAppName       string       `json:"functionAppName,omitempty"`
	ResourceGroup         string       `json:"resourceGroup,omitempty"`
	ConnectionStringOrKey string       `json:"connectionStringOrKey,omitempty"`
}

// GoogleCloudFunctionsCodeIntegrationRequest DTO.
type GoogleCloudFunctionsCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider              CodeProvider `json:"provider,omitempty"`
	ProjectId             string       `json:"projectId,omitempty"`
	Region                string       `json:"region,omitempty"`
	ServiceAccountJsonKey string       `json:"serviceAccountJsonKey,omitempty"`
}

// LlmIntegrationRequest DTO.
type LlmIntegrationRequest struct {
	IntegrationId   string      `json:"integrationId,omitempty"`
	Provider        LlmProvider `json:"provider,omitempty"`
	IntegrationName string      `json:"integrationName,omitempty"`
	IsEnabled       bool        `json:"isEnabled,omitempty"`
	Endpoint        string      `json:"endpoint,omitempty"`
	DefaultModel    string      `json:"defaultModel,omitempty"`
}

// OllamaLlmIntegrationRequest DTO.
type OllamaLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
}

// OpenRouterLlmIntegrationRequest DTO.
type OpenRouterLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// MistralLlmIntegrationRequest DTO.
type MistralLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// GrokLlmIntegrationRequest DTO.
type GrokLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// GroqLlmIntegrationRequest DTO.
type GroqLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// GoogleLlmIntegrationRequest DTO.
type GoogleLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// AnthropicLlmIntegrationRequest DTO.
type AnthropicLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// OpenAiLlmIntegrationRequest DTO.
type OpenAiLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey,omitempty"`
}

// McpIntegrationRequest DTO.
type McpIntegrationRequest struct {
	IntegrationId   string       `json:"integrationId,omitempty"`
	Provider        McpProvider  `json:"provider,omitempty"`
	Transport       McpTransport `json:"transport,omitempty"`
	IntegrationName string       `json:"integrationName,omitempty"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
	Name            string       `json:"name,omitempty"`
	Category        string       `json:"category,omitempty"`
	Description     string       `json:"description,omitempty"`
	Icon            string       `json:"icon,omitempty"`
}

// PlaywrightMcpIntegrationRequest DTO.
type PlaywrightMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	Command   string       `json:"command,omitempty"`
	Args      []string     `json:"args,omitempty"`
	Headless  string       `json:"headless,omitempty"`
}

// MongoDbMcpIntegrationRequest DTO.
type MongoDbMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider         McpProvider  `json:"provider,omitempty"`
	Transport        McpTransport `json:"transport,omitempty"`
	Command          string       `json:"command,omitempty"`
	Args             []string     `json:"args,omitempty"`
	ConnectionString string       `json:"connectionString,omitempty"`
}

// GitHubMcpIntegrationRequest DTO.
type GitHubMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider    McpProvider  `json:"provider,omitempty"`
	Transport   McpTransport `json:"transport,omitempty"`
	ServerUrl   string       `json:"serverUrl,omitempty"`
	AccessToken string       `json:"accessToken,omitempty"`
}

// StripeMcpIntegrationRequest DTO.
type StripeMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	ServerUrl string       `json:"serverUrl,omitempty"`
	ApiKey    string       `json:"apiKey,omitempty"`
}

// BraveSearchMcpIntegrationRequest DTO.
type BraveSearchMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	ServerUrl string       `json:"serverUrl,omitempty"`
	ApiKey    string       `json:"apiKey,omitempty"`
}

// ObsidianMcpIntegrationRequest DTO.
type ObsidianMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider             McpProvider       `json:"provider,omitempty"`
	Transport            McpTransport      `json:"transport,omitempty"`
	Command              string            `json:"command,omitempty"`
	Args                 []string          `json:"args,omitempty"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
}

// TemplateDto DTO.
type TemplateDto struct {
	Id                   string               `json:"id,omitempty"`
	ViewId               string               `json:"viewId,omitempty"`
	TemplateName         string               `json:"templateName,omitempty"`
	Description          string               `json:"description,omitempty"`
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	Medium               NotificationMedium   `json:"medium,omitempty"`
	IsActive             bool                 `json:"isActive,omitempty"`
	Tags                 []string             `json:"tags,omitempty"`
}

// EmailBodyDto DTO.
type EmailBodyDto struct {
	Structure      string              `json:"structure,omitempty"`
	Code           string              `json:"code,omitempty"`
	TemplateEngine EmailTemplateEngine `json:"templateEngine,omitempty"`
}

// EmailMessageContentDto DTO.
type EmailMessageContentDto struct {
	Subject           string                `json:"subject,omitempty"`
	Body              *EmailBodyDto         `json:"body,omitempty"`
	StaticAttachments []*FileResourceRefDto `json:"staticAttachments,omitempty"`
}

// EmailMessageTranslationDto DTO.
type EmailMessageTranslationDto struct {
	Language          string                  `json:"language,omitempty"`
	Content           *EmailMessageContentDto `json:"content,omitempty"`
	StaticAttachments []*FileResourceRefDto   `json:"staticAttachments,omitempty"`
}

// EmailTemplateDto DTO.
type EmailTemplateDto struct {
	TemplateDto
	Translations      []*EmailMessageTranslationDto `json:"translations,omitempty"`
	StaticAttachments []*FileResourceRefDto         `json:"staticAttachments,omitempty"`
}

// PushMessageContentDto DTO.
type PushMessageContentDto struct {
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

// PushMessageTranslationDto DTO.
type PushMessageTranslationDto struct {
	Language string                 `json:"language,omitempty"`
	Content  *PushMessageContentDto `json:"content,omitempty"`
}

// PushTemplateDto DTO.
type PushTemplateDto struct {
	TemplateDto
	Translations []*PushMessageTranslationDto `json:"translations,omitempty"`
}

// SmsMessageContentDto DTO.
type SmsMessageContentDto struct {
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

// SmsMessageTranslationDto DTO.
type SmsMessageTranslationDto struct {
	Language string                `json:"language,omitempty"`
	Content  *SmsMessageContentDto `json:"content,omitempty"`
}

// SmsTemplateDto DTO.
type SmsTemplateDto struct {
	TemplateDto
	Translations []*SmsMessageTranslationDto `json:"translations,omitempty"`
}

// SystemEmailTemplateDto DTO.
type SystemEmailTemplateDto struct {
	EmailTemplateDto
	ImagePreview              string                   `json:"imagePreview,omitempty"`
	Theme                     SystemEmailTemplateTheme `json:"theme,omitempty"`
	SystemGroup               string                   `json:"systemGroup,omitempty"`
	SystemTags                []string                 `json:"systemTags,omitempty"`
	ForTrigger                TriggerType              `json:"forTrigger,omitempty"`
	HiddenSystemEmailTemplate bool                     `json:"hiddenSystemEmailTemplate,omitempty"`
}

// EmailCampaignDeliverySettingsDto DTO.
type EmailCampaignDeliverySettingsDto struct {
	RecipientsSourceType    EmailCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto                 `json:"mappedTokens,omitempty"`
	CampaignTime            float64                            `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings            `json:"respectTimeZoneSettings,omitempty"`
}

// TriggerActionEmailDto DTO.
type TriggerActionEmailDto struct {
	TriggerActionDto
	TemplateId       string                            `json:"templateId,omitempty"`
	DeliverySettings *EmailCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

// PushCampaignDeliverySettingsDto DTO.
type PushCampaignDeliverySettingsDto struct {
	RecipientsSourceType    PushCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto                `json:"mappedTokens,omitempty"`
	CampaignTime            float64                           `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings           `json:"respectTimeZoneSettings,omitempty"`
}

// TriggerActionPushDto DTO.
type TriggerActionPushDto struct {
	TriggerActionDto
	TemplateId       string                           `json:"templateId,omitempty"`
	DeliverySettings *PushCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

// CodeDeliverySettingsDto DTO.
type CodeDeliverySettingsDto struct {
	MappedTokens []*TokenMappingDto `json:"mappedTokens,omitempty"`
}

// TriggerActionCodeDto DTO.
type TriggerActionCodeDto struct {
	TriggerActionDto
	FunctionId       string                   `json:"functionId,omitempty"`
	DeliverySettings *CodeDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

// WebhookDeliverySettingsDto DTO.
type WebhookDeliverySettingsDto struct {
	DestinationIds    []string           `json:"destinationIds,omitempty"`
	EventName         string             `json:"eventName,omitempty"`
	ContentType       string             `json:"contentType,omitempty"`
	IncludeRawPayload bool               `json:"includeRawPayload,omitempty"`
	MappedTokens      []*TokenMappingDto `json:"mappedTokens,omitempty"`
}

// TriggerActionWebhookDto DTO.
type TriggerActionWebhookDto struct {
	TriggerActionDto
	DeliverySettings *WebhookDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

// SmsCampaignDeliverySettingsDto DTO.
type SmsCampaignDeliverySettingsDto struct {
	RecipientsSourceType    SmsCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto               `json:"mappedTokens,omitempty"`
	CampaignTime            float64                          `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings          `json:"respectTimeZoneSettings,omitempty"`
}

// TriggerActionSmsDto DTO.
type TriggerActionSmsDto struct {
	TriggerActionDto
	TemplateId       string                          `json:"templateId,omitempty"`
	DeliverySettings *SmsCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

// SseDeliverySettingsDto DTO.
type SseDeliverySettingsDto struct {
	Audience        string             `json:"audience,omitempty"`
	UserAuthIds     []string           `json:"userAuthIds,omitempty"`
	EventName       string             `json:"eventName,omitempty"`
	PayloadType     string             `json:"payloadType,omitempty"`
	PayloadTemplate string             `json:"payloadTemplate,omitempty"`
	Persist         bool               `json:"persist,omitempty"`
	MappedTokens    []*TokenMappingDto `json:"mappedTokens,omitempty"`
}

// TriggerActionSseDto DTO.
type TriggerActionSseDto struct {
	TriggerActionDto
	DeliverySettings *SseDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

// TriggerActionMarketplaceDto DTO.
type TriggerActionMarketplaceDto struct {
	TriggerActionDto
	FunctionId string            `json:"functionId,omitempty"`
	Payload    map[string]string `json:"payload,omitempty"`
}

// RequestBase DTO.
type RequestBase struct {
	CultureCode   string `json:"cultureCode,omitempty"`
	TimeZoneId    string `json:"timeZoneId,omitempty"`
	Version       string `json:"version,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}

// Env DTO.
type Env struct {
	Value  string `json:"value,omitempty"`
	IsProd bool   `json:"isProd,omitempty"`
}

// CursorArgs DTO.
type CursorArgs struct {
	Field string  `json:"field,omitempty"`
	Order float64 `json:"order,omitempty"`
}

// PagingArgs DTO.
type PagingArgs struct {
	CursorArgs    *CursorArgs `json:"cursorArgs,omitempty"`
	PageSize      float64     `json:"pageSize,omitempty"`
	StartingAfter string      `json:"startingAfter,omitempty"`
	EndingBefore  string      `json:"endingBefore,omitempty"`
}

// CodeMashListPaginationRequestBase DTO.
type CodeMashListPaginationRequestBase struct {
	RequestBase
	ProjectId     string      `json:"projectId,omitempty"`
	Env           string      `json:"env,omitempty"`
	ResolvedEnv   *Env        `json:"resolvedEnv,omitempty"`
	StartingAfter string      `json:"startingAfter,omitempty"`
	EndingBefore  string      `json:"endingBefore,omitempty"`
	PageSize      float64     `json:"pageSize,omitempty"`
	Paging        *PagingArgs `json:"paging,omitempty"`
}

// GetTriggers DTO.
type GetTriggers struct {
	CodeMashListPaginationRequestBase
	SchemaId string `json:"schemaId,omitempty"`
}

// ErrorDto DTO.
type ErrorDto struct {
	Message    string            `json:"message,omitempty"`
	ErrorCode  string            `json:"errorCode,omitempty"`
	Context    map[string]string `json:"context,omitempty"`
	StackTrace []*ErrorDto       `json:"stackTrace,omitempty"`
}

// CodeMashResponseStatus DTO.
type CodeMashResponseStatus struct {
	IsSuccess bool        `json:"isSuccess,omitempty"`
	Errors    []*ErrorDto `json:"errors,omitempty"`
}

// ResponseBase DTO.
type ResponseBase struct {
	ResponseStatus *CodeMashResponseStatus `json:"responseStatus,omitempty"`
}

// GetTriggersResponse DTO.
type GetTriggersResponse struct {
	ResponseBase
}

// EmailToAllUsersDeliverySettingsDto DTO.
type EmailToAllUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	RolesNames []string `json:"rolesNames,omitempty"`
	UserTags   []string `json:"userTags,omitempty"`
}

// EmailToAccountUsersDeliverySettingsDto DTO.
type EmailToAccountUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	UserRecipients      []string `json:"userRecipients,omitempty"`
	UserCc              []string `json:"userCc,omitempty"`
	UserBcc             []string `json:"userBcc,omitempty"`
	SingleEmailStrategy bool     `json:"singleEmailStrategy,omitempty"`
}

// EmailToUsersDeliverySettingsDto DTO.
type EmailToUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	UserRecipients      []string `json:"userRecipients,omitempty"`
	UserCc              []string `json:"userCc,omitempty"`
	UserBcc             []string `json:"userBcc,omitempty"`
	SingleEmailStrategy bool     `json:"singleEmailStrategy,omitempty"`
}

// EmailToEmailAddressesDeliverySettingsDto DTO.
type EmailToEmailAddressesDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	Recipients          []string `json:"recipients,omitempty"`
	RecipientsCc        []string `json:"recipientsCc,omitempty"`
	RecipientsBcc       []string `json:"recipientsBcc,omitempty"`
	SingleEmailStrategy bool     `json:"singleEmailStrategy,omitempty"`
}

// EmailToCollectionRecordsDeliverySettingsDto DTO.
type EmailToCollectionRecordsDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	Fields     []string                              `json:"fields,omitempty"`
	SchemaName string                                `json:"schemaName,omitempty"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
}

// PushToAllUsersDeliverySettingsDto DTO.
type PushToAllUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	RolesNames []string `json:"rolesNames,omitempty"`
	UserTags   []string `json:"userTags,omitempty"`
}

// PushToUsersDeliverySettingsDto DTO.
type PushToUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Recipients []string `json:"recipients,omitempty"`
}

// PushToAccountUsersDeliverySettingsDto DTO.
type PushToAccountUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Recipients []string `json:"recipients,omitempty"`
}

// PushToCollectionRecordsDeliverySettingsDto DTO.
type PushToCollectionRecordsDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Fields     []string                              `json:"fields,omitempty"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	SchemaName string                                `json:"schemaName,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
}

// PushDeviceDeliveryTokenDto DTO.
type PushDeviceDeliveryTokenDto struct {
}

// PushToDevicesDeliverySettingsDto DTO.
type PushToDevicesDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Devices []*PushDeviceDeliveryTokenDto `json:"devices,omitempty"`
}

// SmsToAllUsersDeliverySettingsDto DTO.
type SmsToAllUsersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	RolesNames []string `json:"rolesNames,omitempty"`
	UserTags   []string `json:"userTags,omitempty"`
}

// SmsToUsersDeliverySettingsDto DTO.
type SmsToUsersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	Recipients []string `json:"recipients,omitempty"`
}

// SmsToCollectionRecordsDeliverySettingsDto DTO.
type SmsToCollectionRecordsDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	Fields     []string                              `json:"fields,omitempty"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	SchemaName string                                `json:"schemaName,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
}

// SmsToPhoneNumbersDeliverySettingsDto DTO.
type SmsToPhoneNumbersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	PhoneNumbers []string `json:"phoneNumbers,omitempty"`
}

// IntegrationDto DTO.
type IntegrationDto struct {
	ViewId                            string   `json:"viewId,omitempty"`
	IntegrationName                   string   `json:"integrationName,omitempty"`
	IsEnabled                         bool     `json:"isEnabled,omitempty"`
	Env                               string   `json:"env,omitempty"`
	LastIntegrationTestAtUtc          string   `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      bool     `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         []string `json:"lastIntegrationTestErrors,omitempty"`
	HumanDeliveryConfirmedAtUtc       string   `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool     `json:"requiresHumanDeliveryConfirmation,omitempty"`
}

// LlmIntegrationDto DTO.
type LlmIntegrationDto struct {
	IntegrationDto
	Provider      LlmProvider `json:"provider,omitempty"`
	BaseUrl       string      `json:"baseUrl,omitempty"`
	DefaultModel  string      `json:"defaultModel,omitempty"`
	IsConfigured  bool        `json:"isConfigured,omitempty"`
	IsSystemOwned bool        `json:"isSystemOwned,omitempty"`
}

// OpenAiLlmIntegrationDto DTO.
type OpenAiLlmIntegrationDto struct {
	LlmIntegrationDto
}

// AnthropicLlmIntegrationDto DTO.
type AnthropicLlmIntegrationDto struct {
	LlmIntegrationDto
}

// OllamaLlmIntegrationDto DTO.
type OllamaLlmIntegrationDto struct {
	LlmIntegrationDto
}

// GroqLlmIntegrationDto DTO.
type GroqLlmIntegrationDto struct {
	LlmIntegrationDto
}

// GoogleLlmIntegrationDto DTO.
type GoogleLlmIntegrationDto struct {
	LlmIntegrationDto
}

// MistralLlmIntegrationDto DTO.
type MistralLlmIntegrationDto struct {
	LlmIntegrationDto
}

// OpenRouterLlmIntegrationDto DTO.
type OpenRouterLlmIntegrationDto struct {
	LlmIntegrationDto
}

// GrokLlmIntegrationDto DTO.
type GrokLlmIntegrationDto struct {
	LlmIntegrationDto
}

// McpMetadata DTO.
type McpMetadata struct {
	Name        string `json:"name,omitempty"`
	Category    string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
}

// McpIntegrationDto DTO.
type McpIntegrationDto struct {
	IntegrationDto
	Provider      McpProvider  `json:"provider,omitempty"`
	Transport     McpTransport `json:"transport,omitempty"`
	Metadata      *McpMetadata `json:"metadata,omitempty"`
	IsConfigured  bool         `json:"isConfigured,omitempty"`
	IsSystemOwned bool         `json:"isSystemOwned,omitempty"`
	Command       string       `json:"command,omitempty"`
	Args          []string     `json:"args,omitempty"`
	ServerUrl     string       `json:"serverUrl,omitempty"`
	Auth          McpAuth      `json:"auth,omitempty"`
}

// DockerMcpIntegrationDto DTO.
type DockerMcpIntegrationDto struct {
	McpIntegrationDto
}

// GoogleCalendarMcpIntegrationDto DTO.
type GoogleCalendarMcpIntegrationDto struct {
	McpIntegrationDto
}

// ObsidianMcpIntegrationDto DTO.
type ObsidianMcpIntegrationDto struct {
	McpIntegrationDto
}

// CodeIntegrationDto DTO.
type CodeIntegrationDto struct {
	IntegrationDto
	Provider CodeProvider `json:"provider,omitempty"`
}

// AwsLambdaCrossAccountRoleCodeIntegrationDto DTO.
type AwsLambdaCrossAccountRoleCodeIntegrationDto struct {
	CodeIntegrationDto
	Region     string `json:"region,omitempty"`
	RoleArn    string `json:"roleArn,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
}

// AwsLambdaIamCodeIntegrationDto DTO.
type AwsLambdaIamCodeIntegrationDto struct {
	CodeIntegrationDto
	Region string `json:"region,omitempty"`
}

// AzureFunctionsCodeIntegrationDto DTO.
type AzureFunctionsCodeIntegrationDto struct {
	CodeIntegrationDto
	FunctionAppName string `json:"functionAppName,omitempty"`
	ResourceGroup   string `json:"resourceGroup,omitempty"`
}

// GoogleCloudFunctionsCodeIntegrationDto DTO.
type GoogleCloudFunctionsCodeIntegrationDto struct {
	CodeIntegrationDto
	ProjectId string `json:"projectId,omitempty"`
	Region    string `json:"region,omitempty"`
}

// PaymentsIntegrationDto DTO.
type PaymentsIntegrationDto struct {
	IntegrationDto
	GatewayPlatform PaymentGatewayPlatform `json:"gatewayPlatform,omitempty"`
}

// AdyenPaymentIntegrationDto DTO.
type AdyenPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantAccount string `json:"merchantAccount,omitempty"`
	Environment     string `json:"environment,omitempty"`
	WebhookId       string `json:"webhookId,omitempty"`
}

// AppleInAppPaymentIntegrationDto DTO.
type AppleInAppPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantIdentifier string `json:"merchantIdentifier,omitempty"`
	MerchantDomain     string `json:"merchantDomain,omitempty"`
	DisplayName        string `json:"displayName,omitempty"`
}

// GoogleInAppPaymentIntegrationDto DTO.
type GoogleInAppPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantId        string `json:"merchantId,omitempty"`
	MerchantName      string `json:"merchantName,omitempty"`
	Gateway           string `json:"gateway,omitempty"`
	GatewayMerchantId string `json:"gatewayMerchantId,omitempty"`
}

// LemonSqueezyPaymentIntegrationDto DTO.
type LemonSqueezyPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreId    string `json:"storeId,omitempty"`
	IsTestMode bool   `json:"isTestMode,omitempty"`
}

// MolliePaymentIntegrationDto DTO.
type MolliePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ProfileId  string `json:"profileId,omitempty"`
	IsTestMode bool   `json:"isTestMode,omitempty"`
}

// PaddlePaymentIntegrationDto DTO.
type PaddlePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	Environment     string `json:"environment,omitempty"`
	ClientSideToken string `json:"clientSideToken,omitempty"`
}

// PayPalPaymentIntegrationDto DTO.
type PayPalPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ClientId    string `json:"clientId,omitempty"`
	Environment string `json:"environment,omitempty"`
	BrandName   string `json:"brandName,omitempty"`
}

// StripePaymentIntegrationDto DTO.
type StripePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	PublishableKey    string `json:"publishableKey,omitempty"`
	WebhookEndpointId string `json:"webhookEndpointId,omitempty"`
	DefaultCurrency   string `json:"defaultCurrency,omitempty"`
}

// ShopifyPaymentIntegrationDto DTO.
type ShopifyPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ShopDomain    string `json:"shopDomain,omitempty"`
	WebhookSecret string `json:"webhookSecret,omitempty"`
}

// WooCommercePaymentIntegrationDto DTO.
type WooCommercePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreUrl      string `json:"storeUrl,omitempty"`
	WebhookSecret string `json:"webhookSecret,omitempty"`
}

// MagentoPaymentIntegrationDto DTO.
type MagentoPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreUrl      string `json:"storeUrl,omitempty"`
	WebhookSecret string `json:"webhookSecret,omitempty"`
}

// BraintreePaymentIntegrationDto DTO.
type BraintreePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantId    string `json:"merchantId,omitempty"`
	Environment   string `json:"environment,omitempty"`
	WebhookSecret string `json:"webhookSecret,omitempty"`
}

// AuthorizeNetPaymentIntegrationDto DTO.
type AuthorizeNetPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantLoginId     string `json:"merchantLoginId,omitempty"`
	Environment         string `json:"environment,omitempty"`
	WebhookSignatureKey string `json:"webhookSignatureKey,omitempty"`
}

// CheckOutComPaymentIntegrationDto DTO.
type CheckOutComPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantAccount string `json:"merchantAccount,omitempty"`
	Environment     string `json:"environment,omitempty"`
	WebhookSecret   string `json:"webhookSecret,omitempty"`
}

// WorldpayPaymentIntegrationDto DTO.
type WorldpayPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantCode  string `json:"merchantCode,omitempty"`
	Environment   string `json:"environment,omitempty"`
	WebhookSecret string `json:"webhookSecret,omitempty"`
}

// MembershipIntegrationDto DTO.
type MembershipIntegrationDto struct {
	IntegrationDto
	Provider MembershipProvider `json:"provider,omitempty"`
}

// AppleSignInMembershipIntegrationDto DTO.
type AppleSignInMembershipIntegrationDto struct {
	MembershipIntegrationDto
	TeamId      string `json:"teamId,omitempty"`
	AppBundleId string `json:"appBundleId,omitempty"`
	ServiceId   string `json:"serviceId,omitempty"`
}

// GitHubMembershipIntegrationDto DTO.
type GitHubMembershipIntegrationDto struct {
	MembershipIntegrationDto
	ClientId string `json:"clientId,omitempty"`
}

// GoogleMembershipIntegrationDto DTO.
type GoogleMembershipIntegrationDto struct {
	MembershipIntegrationDto
	ClientId string `json:"clientId,omitempty"`
}

// MetaMembershipIntegrationDto DTO.
type MetaMembershipIntegrationDto struct {
	MembershipIntegrationDto
	AppId string `json:"appId,omitempty"`
}

// MicrosoftMembershipIntegrationDto DTO.
type MicrosoftMembershipIntegrationDto struct {
	MembershipIntegrationDto
	TenantId string `json:"tenantId,omitempty"`
	ClientId string `json:"clientId,omitempty"`
}

// OktaMembershipIntegrationDto DTO.
type OktaMembershipIntegrationDto struct {
	MembershipIntegrationDto
	Domain   string `json:"domain,omitempty"`
	ClientId string `json:"clientId,omitempty"`
}

// XMembershipIntegrationDto DTO.
type XMembershipIntegrationDto struct {
	MembershipIntegrationDto
	ApiKey string `json:"apiKey,omitempty"`
}

// LoggingIntegrationDto DTO.
type LoggingIntegrationDto struct {
	IntegrationDto
	Provider LoggingProvider `json:"provider,omitempty"`
}

// AmqpLoggingIntegrationDto DTO.
type AmqpLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Host        string  `json:"host,omitempty"`
	Port        float64 `json:"port,omitempty"`
	VirtualHost string  `json:"virtualHost,omitempty"`
	Exchange    string  `json:"exchange,omitempty"`
	RoutingKey  string  `json:"routingKey,omitempty"`
}

// AwsKinesisLoggingIntegrationDto DTO.
type AwsKinesisLoggingIntegrationDto struct {
	LoggingIntegrationDto
	StreamName string `json:"streamName,omitempty"`
	Region     string `json:"region,omitempty"`
}

// AwsS3CrossAccountRoleLoggingIntegrationDto DTO.
type AwsS3CrossAccountRoleLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BucketName string `json:"bucketName,omitempty"`
	Region     string `json:"region,omitempty"`
	RoleArn    string `json:"roleArn,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
}

// AwsS3IamLoggingIntegrationDto DTO.
type AwsS3IamLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BucketName string `json:"bucketName,omitempty"`
	Region     string `json:"region,omitempty"`
}

// AzureOtelLoggingIntegrationDto DTO.
type AzureOtelLoggingIntegrationDto struct {
	LoggingIntegrationDto
	EndpointUrl  string `json:"endpointUrl,omitempty"`
	ResourceName string `json:"resourceName,omitempty"`
}

// DataDogLoggingIntegrationDto DTO.
type DataDogLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Site        string `json:"site,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
	Environment string `json:"environment,omitempty"`
}

// ElasticSearchLoggingIntegrationDto DTO.
type ElasticSearchLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Uri   string `json:"uri,omitempty"`
	Index string `json:"index,omitempty"`
}

// InternalKafkaLoggingIntegrationDto DTO.
type InternalKafkaLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BootstrapServers string `json:"bootstrapServers,omitempty"`
	Topic            string `json:"topic,omitempty"`
	SecurityProtocol string `json:"securityProtocol,omitempty"`
}

// KafkaLoggingIntegrationDto DTO.
type KafkaLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BootstrapServers string `json:"bootstrapServers,omitempty"`
	Topic            string `json:"topic,omitempty"`
	SecurityProtocol string `json:"securityProtocol,omitempty"`
}

// KibanaLoggingIntegrationDto DTO.
type KibanaLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Uri     string `json:"uri,omitempty"`
	SpaceId string `json:"spaceId,omitempty"`
}

// LocalFileLoggingIntegrationDto DTO.
type LocalFileLoggingIntegrationDto struct {
	LoggingIntegrationDto
	RootPath string `json:"rootPath,omitempty"`
}

// MongoDbLoggingIntegrationDto DTO.
type MongoDbLoggingIntegrationDto struct {
	LoggingIntegrationDto
	DatabaseName string `json:"databaseName,omitempty"`
}

// NewRelicLoggingIntegrationDto DTO.
type NewRelicLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Region      string `json:"region,omitempty"`
	ServiceName string `json:"serviceName,omitempty"`
}

// PrometheusLoggingIntegrationDto DTO.
type PrometheusLoggingIntegrationDto struct {
	LoggingIntegrationDto
	EndpointUrl string `json:"endpointUrl,omitempty"`
	JobName     string `json:"jobName,omitempty"`
}

// SplunkLoggingIntegrationDto DTO.
type SplunkLoggingIntegrationDto struct {
	LoggingIntegrationDto
	HecEndpointUrl string `json:"hecEndpointUrl,omitempty"`
	Index          string `json:"index,omitempty"`
}

// FilesIntegrationDto DTO.
type FilesIntegrationDto struct {
	IntegrationDto
	Provider FileProvider `json:"provider,omitempty"`
}

// AppleICloudFilesIntegrationDto DTO.
type AppleICloudFilesIntegrationDto struct {
	FilesIntegrationDto
	ContainerIdentifier string `json:"containerIdentifier,omitempty"`
	RelativePath        string `json:"relativePath,omitempty"`
}

// AwsS3CrossAccountRoleFilesIntegrationDto DTO.
type AwsS3CrossAccountRoleFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string `json:"bucketName,omitempty"`
	Region     string `json:"region,omitempty"`
	RoleArn    string `json:"roleArn,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
}

// AwsS3IamFilesIntegrationDto DTO.
type AwsS3IamFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string `json:"bucketName,omitempty"`
	Region     string `json:"region,omitempty"`
}

// AzureBlobFilesIntegrationDto DTO.
type AzureBlobFilesIntegrationDto struct {
	FilesIntegrationDto
	BlobName string `json:"blobName,omitempty"`
}

// DropBoxFilesIntegrationDto DTO.
type DropBoxFilesIntegrationDto struct {
	FilesIntegrationDto
	RootPath string `json:"rootPath,omitempty"`
}

// FtpFilesIntegrationDto DTO.
type FtpFilesIntegrationDto struct {
	FilesIntegrationDto
	Host     string  `json:"host,omitempty"`
	Port     float64 `json:"port,omitempty"`
	RootPath string  `json:"rootPath,omitempty"`
	UseSsl   bool    `json:"useSsl,omitempty"`
}

// GoogleCloudFilesIntegrationDto DTO.
type GoogleCloudFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string `json:"bucketName,omitempty"`
}

// GoogleDriveFilesIntegrationDto DTO.
type GoogleDriveFilesIntegrationDto struct {
	FilesIntegrationDto
	RootFolderId string `json:"rootFolderId,omitempty"`
}

// LocalFilesIntegrationDto DTO.
type LocalFilesIntegrationDto struct {
	FilesIntegrationDto
	RootPath string `json:"rootPath,omitempty"`
}

// DatabaseIntegrationDto DTO.
type DatabaseIntegrationDto struct {
	IntegrationDto
	Provider DatabaseProvider `json:"provider,omitempty"`
}

// MongoDbConnectionStringIntegrationDto DTO.
type MongoDbConnectionStringIntegrationDto struct {
	DatabaseIntegrationDto
	DatabaseName string `json:"databaseName,omitempty"`
}

// MongoDbAtlasFlexManagedIntegrationDto DTO.
type MongoDbAtlasFlexManagedIntegrationDto struct {
	DatabaseIntegrationDto
	DatabaseName     string            `json:"databaseName,omitempty"`
	NorbixRegionCode string            `json:"norbixRegionCode,omitempty"`
	FlexTierCode     string            `json:"flexTierCode,omitempty"`
	Status           IntegrationStatus `json:"status,omitempty"`
	AtlasProjectId   string            `json:"atlasProjectId,omitempty"`
	AtlasClusterName string            `json:"atlasClusterName,omitempty"`
	FailureReason    string            `json:"failureReason,omitempty"`
}

// SmsIntegrationDto DTO.
type SmsIntegrationDto struct {
	IntegrationDto
	Provider SmsProvider `json:"provider,omitempty"`
}

// BirdSmsIntegrationDto DTO.
type BirdSmsIntegrationDto struct {
	SmsIntegrationDto
	Originator string `json:"originator,omitempty"`
	Region     string `json:"region,omitempty"`
}

// PlivoSmsIntegrationDto DTO.
type PlivoSmsIntegrationDto struct {
	SmsIntegrationDto
	AuthId          string `json:"authId,omitempty"`
	FromPhoneNumber string `json:"fromPhoneNumber,omitempty"`
}

// SinchSmsIntegrationDto DTO.
type SinchSmsIntegrationDto struct {
	SmsIntegrationDto
	ServicePlanId   string `json:"servicePlanId,omitempty"`
	FromPhoneNumber string `json:"fromPhoneNumber,omitempty"`
}

// TelesignSmsIntegrationDto DTO.
type TelesignSmsIntegrationDto struct {
	SmsIntegrationDto
	CustomerId string `json:"customerId,omitempty"`
	FromSender string `json:"fromSender,omitempty"`
}

// TelnyxSmsIntegrationDto DTO.
type TelnyxSmsIntegrationDto struct {
	SmsIntegrationDto
	MessagingProfileId string `json:"messagingProfileId,omitempty"`
	FromPhoneNumber    string `json:"fromPhoneNumber,omitempty"`
}

// TwilioSmsIntegrationDto DTO.
type TwilioSmsIntegrationDto struct {
	SmsIntegrationDto
	AccountSid      string `json:"accountSid,omitempty"`
	FromPhoneNumber string `json:"fromPhoneNumber,omitempty"`
}

// VonageSmsIntegrationDto DTO.
type VonageSmsIntegrationDto struct {
	SmsIntegrationDto
	ApiKey     string `json:"apiKey,omitempty"`
	FromSender string `json:"fromSender,omitempty"`
}

// FakeSmsIntegrationDto DTO.
type FakeSmsIntegrationDto struct {
	SmsIntegrationDto
}

// PushIntegrationDto DTO.
type PushIntegrationDto struct {
	IntegrationDto
	Provider PushProvider `json:"provider,omitempty"`
}

// AndroidFirebasePushIntegrationDto DTO.
type AndroidFirebasePushIntegrationDto struct {
	PushIntegrationDto
	ProjectId   string `json:"projectId,omitempty"`
	ClientEmail string `json:"clientEmail,omitempty"`
}

// AppleApnsPushIntegrationDto DTO.
type AppleApnsPushIntegrationDto struct {
	PushIntegrationDto
	TeamId      string `json:"teamId,omitempty"`
	AppBundleId string `json:"appBundleId,omitempty"`
}

// ChromePluginPushIntegrationDto DTO.
type ChromePluginPushIntegrationDto struct {
	PushIntegrationDto
	ExtensionId    string `json:"extensionId,omitempty"`
	VapidPublicKey string `json:"vapidPublicKey,omitempty"`
	Subject        string `json:"subject,omitempty"`
}

// ChromeWebPushIntegrationDto DTO.
type ChromeWebPushIntegrationDto struct {
	PushIntegrationDto
	VapidPublicKey string `json:"vapidPublicKey,omitempty"`
	Subject        string `json:"subject,omitempty"`
}

// EdgeWebPushIntegrationDto DTO.
type EdgeWebPushIntegrationDto struct {
	PushIntegrationDto
	VapidPublicKey string `json:"vapidPublicKey,omitempty"`
	Subject        string `json:"subject,omitempty"`
}

// FirefoxWebPushIntegrationDto DTO.
type FirefoxWebPushIntegrationDto struct {
	PushIntegrationDto
	VapidPublicKey string `json:"vapidPublicKey,omitempty"`
	Subject        string `json:"subject,omitempty"`
}

// SafariPushIntegrationDto DTO.
type SafariPushIntegrationDto struct {
	PushIntegrationDto
	WebsitePushId string `json:"websitePushId,omitempty"`
}

// FakePushIntegrationDto DTO.
type FakePushIntegrationDto struct {
	PushIntegrationDto
}

// EmailIntegrationDto DTO.
type EmailIntegrationDto struct {
	IntegrationDto
	Provider        EmailProvider `json:"provider,omitempty"`
	EmailAddress    string        `json:"emailAddress,omitempty"`
	EmailSenderName string        `json:"emailSenderName,omitempty"`
}

// AwsSesEmailIntegrationDto DTO.
type AwsSesEmailIntegrationDto struct {
	EmailIntegrationDto
	Region               string `json:"region,omitempty"`
	IdentityArn          string `json:"identityArn,omitempty"`
	ConfigurationSetName string `json:"configurationSetName,omitempty"`
}

// AwsCrossAccountRoleEmailIntegrationDto DTO.
type AwsCrossAccountRoleEmailIntegrationDto struct {
	AwsSesEmailIntegrationDto
	RoleArn    string `json:"roleArn,omitempty"`
	ExternalId string `json:"externalId,omitempty"`
}

// AwsIamEmailIntegrationDto DTO.
type AwsIamEmailIntegrationDto struct {
	AwsSesEmailIntegrationDto
}

// MailGunEmailIntegrationDto DTO.
type MailGunEmailIntegrationDto struct {
	EmailIntegrationDto
	Domain string        `json:"domain,omitempty"`
	Region MailGunRegion `json:"region,omitempty"`
}

// SendGridEmailIntegrationDto DTO.
type SendGridEmailIntegrationDto struct {
	EmailIntegrationDto
}

// SmtpEmailIntegrationDto DTO.
type SmtpEmailIntegrationDto struct {
	EmailIntegrationDto
	HostName string  `json:"hostName,omitempty"`
	Port     float64 `json:"port,omitempty"`
}

// FakeEmailIntegrationDto DTO.
type FakeEmailIntegrationDto struct {
	EmailIntegrationDto
}

// WebhookDestinationDto DTO.
type WebhookDestinationDto struct {
	ViewId          string            `json:"viewId,omitempty"`
	DestinationName string            `json:"destinationName,omitempty"`
	EndpointUrl     string            `json:"endpointUrl,omitempty"`
	SelectedEvents  []string          `json:"selectedEvents,omitempty"`
	ExtraHeaders    map[string]string `json:"extraHeaders,omitempty"`
	IsEnabled       bool              `json:"isEnabled,omitempty"`
}

// WebhookIntegrationDto DTO.
type WebhookIntegrationDto struct {
	IntegrationDto
	IsConfigured bool                     `json:"isConfigured,omitempty"`
	Destinations []*WebhookDestinationDto `json:"destinations,omitempty"`
	ExtraHeaders map[string]string        `json:"extraHeaders,omitempty"`
}

// SchedulerTaskDto DTO.
type SchedulerTaskDto struct {
	ProjectId     string            `json:"projectId,omitempty"`
	TaskId        string            `json:"taskId,omitempty"`
	Name          string            `json:"name,omitempty"`
	Description   string            `json:"description,omitempty"`
	Cron          string            `json:"cron,omitempty"`
	Type          SchedulerTaskType `json:"type,omitempty"`
	PayloadJson   string            `json:"payloadJson,omitempty"`
	InitiatorId   string            `json:"initiatorId,omitempty"`
	IsEnabled     bool              `json:"isEnabled,omitempty"`
	StopOnError   bool              `json:"stopOnError,omitempty"`
	CreatedAtUnix float64           `json:"createdAtUnix,omitempty"`
	UpdatedAtUnix float64           `json:"updatedAtUnix,omitempty"`
}

// MongoDbAggregateDto DTO.
type MongoDbAggregateDto struct {
	ViewId       string `json:"viewId,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Description  string `json:"description,omitempty"`
	SchemaViewId string `json:"schemaViewId,omitempty"`
	Pipeline     string `json:"pipeline,omitempty"`
}

// MarketplaceIntegrationDto DTO.
type MarketplaceIntegrationDto struct {
	IntegrationDto
	ListingViewId string               `json:"listingViewId,omitempty"`
	Transport     MarketplaceTransport `json:"transport,omitempty"`
	Vendor        string               `json:"vendor,omitempty"`
	Category      MarketplaceCategory  `json:"category,omitempty"`
	Description   string               `json:"description,omitempty"`
	Config        map[string]string    `json:"config,omitempty"`
}

// MarketplaceTokenMappingDto DTO.
type MarketplaceTokenMappingDto struct {
	Token      string                       `json:"token,omitempty"`
	Resolver   MarketplaceTokenResolverKind `json:"resolver,omitempty"`
	Value      string                       `json:"value,omitempty"`
	SecretKeys []string                     `json:"secretKeys,omitempty"`
	Format     MarketplaceSecretValueFormat `json:"format,omitempty"`
}

// MarketplaceFunctionDto DTO.
type MarketplaceFunctionDto struct {
	ViewId            string                        `json:"viewId,omitempty"`
	IntegrationViewId string                        `json:"integrationViewId,omitempty"`
	FunctionKey       string                        `json:"functionKey,omitempty"`
	DisplayName       string                        `json:"displayName,omitempty"`
	Description       string                        `json:"description,omitempty"`
	IsEnabled         bool                          `json:"isEnabled,omitempty"`
	RequestTemplate   string                        `json:"requestTemplate,omitempty"`
	MappedTokens      []*MarketplaceTokenMappingDto `json:"mappedTokens,omitempty"`
}

// MarketplaceFieldDefinitionDto DTO.
type MarketplaceFieldDefinitionDto struct {
	Key               string               `json:"key,omitempty"`
	Label             string               `json:"label,omitempty"`
	Description       string               `json:"description,omitempty"`
	DocumentationUrl  string               `json:"documentationUrl,omitempty"`
	Type              MarketplaceFieldType `json:"type,omitempty"`
	IsRequired        bool                 `json:"isRequired,omitempty"`
	DefaultValue      string               `json:"defaultValue,omitempty"`
	Placeholder       string               `json:"placeholder,omitempty"`
	ValidationPattern string               `json:"validationPattern,omitempty"`
	AllowedValues     []string             `json:"allowedValues,omitempty"`
}

// MarketplaceFunctionParameterDto DTO.
type MarketplaceFunctionParameterDto struct {
	Name         string `json:"name,omitempty"`
	Type         string `json:"type,omitempty"`
	Description  string `json:"description,omitempty"`
	IsRequired   bool   `json:"isRequired,omitempty"`
	DefaultValue string `json:"defaultValue,omitempty"`
}

// MarketplaceParameterSpecDto DTO.
type MarketplaceParameterSpecDto struct {
	Name             string                       `json:"name,omitempty"`
	Location         MarketplaceParameterLocation `json:"location,omitempty"`
	ValueTemplate    string                       `json:"valueTemplate,omitempty"`
	Label            string                       `json:"label,omitempty"`
	Description      string                       `json:"description,omitempty"`
	DocumentationUrl string                       `json:"documentationUrl,omitempty"`
	Type             string                       `json:"type,omitempty"`
	IsRequired       bool                         `json:"isRequired,omitempty"`
}

// MarketplaceHttpRequestSpecDto DTO.
type MarketplaceHttpRequestSpecDto struct {
	Method       string                         `json:"method,omitempty"`
	PathTemplate string                         `json:"pathTemplate,omitempty"`
	Parameters   []*MarketplaceParameterSpecDto `json:"parameters,omitempty"`
	ContentType  string                         `json:"contentType,omitempty"`
}

// MarketplaceFunctionDefinitionDto DTO.
type MarketplaceFunctionDefinitionDto struct {
	DefinitionId         string                             `json:"definitionId,omitempty"`
	FunctionKey          string                             `json:"functionKey,omitempty"`
	DisplayName          string                             `json:"displayName,omitempty"`
	Description          string                             `json:"description,omitempty"`
	Group                string                             `json:"group,omitempty"`
	Parameters           []*MarketplaceFunctionParameterDto `json:"parameters,omitempty"`
	RequestSchema        string                             `json:"requestSchema,omitempty"`
	RequestTemplate      string                             `json:"requestTemplate,omitempty"`
	Request              *MarketplaceHttpRequestSpecDto     `json:"request,omitempty"`
	DefaultTokenMappings []*MarketplaceTokenMappingDto      `json:"defaultTokenMappings,omitempty"`
}

// MarketplaceListingDto DTO.
type MarketplaceListingDto struct {
	ViewId           string                              `json:"viewId,omitempty"`
	Slug             string                              `json:"slug,omitempty"`
	DisplayName      string                              `json:"displayName,omitempty"`
	Vendor           string                              `json:"vendor,omitempty"`
	Category         MarketplaceCategory                 `json:"category,omitempty"`
	Transport        MarketplaceTransport                `json:"transport,omitempty"`
	Description      string                              `json:"description,omitempty"`
	IconUrl          string                              `json:"iconUrl,omitempty"`
	DocumentationUrl string                              `json:"documentationUrl,omitempty"`
	IsOfficial       bool                                `json:"isOfficial,omitempty"`
	Tags             []string                            `json:"tags,omitempty"`
	SpecVersion      float64                             `json:"specVersion,omitempty"`
	ConfigFields     []*MarketplaceFieldDefinitionDto    `json:"configFields,omitempty"`
	SecretFields     []*MarketplaceFieldDefinitionDto    `json:"secretFields,omitempty"`
	Functions        []*MarketplaceFunctionDefinitionDto `json:"functions,omitempty"`
}

// AdminPortalModuleDto DTO.
type AdminPortalModuleDto struct {
	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Enabled     bool   `json:"enabled,omitempty"`
}

// AiChatEntryWireDto DTO.
type AiChatEntryWireDto struct {
	Kind                 string  `json:"kind,omitempty"`
	Id                   string  `json:"id,omitempty"`
	Seq                  float64 `json:"seq,omitempty"`
	AtUtc                string  `json:"atUtc,omitempty"`
	RefEntryId           string  `json:"refEntryId,omitempty"`
	WorkItemId           string  `json:"workItemId,omitempty"`
	Feedback             string  `json:"feedback,omitempty"`
	FeedbackAtUtc        string  `json:"feedbackAtUtc,omitempty"`
	FeedbackByUserAuthId string  `json:"feedbackByUserAuthId,omitempty"`
}

// AiChatEntryAttachmentWireDto DTO.
type AiChatEntryAttachmentWireDto struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// UserMessageEntryWireDto DTO.
type UserMessageEntryWireDto struct {
	AiChatEntryWireDto
	Kind        string                          `json:"kind,omitempty"`
	Text        string                          `json:"text,omitempty"`
	Attachments []*AiChatEntryAttachmentWireDto `json:"attachments,omitempty"`
}

// AiChatEntrySourceWireDto DTO.
type AiChatEntrySourceWireDto struct {
	Kind          string  `json:"kind,omitempty"`
	RequirementId string  `json:"requirementId,omitempty"`
	SessionId     string  `json:"sessionId,omitempty"`
	EntryId       string  `json:"entryId,omitempty"`
	EntrySeq      float64 `json:"entrySeq,omitempty"`
	ArtifactId    string  `json:"artifactId,omitempty"`
	Label         string  `json:"label,omitempty"`
	Step          float64 `json:"step,omitempty"`
}

// AssistantTextEntryWireDto DTO.
type AssistantTextEntryWireDto struct {
	AiChatEntryWireDto
	Kind        string                      `json:"kind,omitempty"`
	Text        string                      `json:"text,omitempty"`
	IsStreaming bool                        `json:"isStreaming,omitempty"`
	Sources     []*AiChatEntrySourceWireDto `json:"sources,omitempty"`
}

// AiChatQuestionOptionWireDto DTO.
type AiChatQuestionOptionWireDto struct {
	Value string `json:"value,omitempty"`
	Label string `json:"label,omitempty"`
}

// AiChatQuestionWireDto DTO.
type AiChatQuestionWireDto struct {
	Id            string                         `json:"id,omitempty"`
	Text          string                         `json:"text,omitempty"`
	Options       []*AiChatQuestionOptionWireDto `json:"options,omitempty"`
	Default       string                         `json:"default,omitempty"`
	AllowFreeText bool                           `json:"allowFreeText,omitempty"`
}

// AiChatGateResultWireDto DTO.
type AiChatGateResultWireDto struct {
	Class                  string   `json:"class,omitempty"`
	Reason                 string   `json:"reason,omitempty"`
	AffectedRequirementIds []string `json:"affectedRequirementIds,omitempty"`
}

// AssistantQuestionEntryWireDto DTO.
type AssistantQuestionEntryWireDto struct {
	AiChatEntryWireDto
	Kind      string                   `json:"kind,omitempty"`
	Questions []*AiChatQuestionWireDto `json:"questions,omitempty"`
	Status    string                   `json:"status,omitempty"`
	Scope     string                   `json:"scope,omitempty"`
	Gate      *AiChatGateResultWireDto `json:"gate,omitempty"`
}

// UserAnswerEntryWireDto DTO.
type UserAnswerEntryWireDto struct {
	AiChatEntryWireDto
	Kind    string            `json:"kind,omitempty"`
	Answers map[string]string `json:"answers,omitempty"`
}

// AiChatPlanStepInputsWireDto DTO.
type AiChatPlanStepInputsWireDto struct {
	Artifacts    []string `json:"artifacts,omitempty"`
	Requirements []string `json:"requirements,omitempty"`
}

// AiChatPlanStepDoneCheckWireDto DTO.
type AiChatPlanStepDoneCheckWireDto struct {
	Check    string `json:"check,omitempty"`
	ArgsJson string `json:"argsJson,omitempty"`
}

// AiChatPlanStepLoopWireDto DTO.
type AiChatPlanStepLoopWireDto struct {
	MaxIterations float64 `json:"maxIterations,omitempty"`
	MaxToolCalls  float64 `json:"maxToolCalls,omitempty"`
}

// AiChatPlanStepWireDto DTO.
type AiChatPlanStepWireDto struct {
	N          float64                           `json:"n,omitempty"`
	Tool       string                            `json:"tool,omitempty"`
	Title      string                            `json:"title,omitempty"`
	Goal       string                            `json:"goal,omitempty"`
	Inputs     *AiChatPlanStepInputsWireDto      `json:"inputs,omitempty"`
	DependsOn  []float64                         `json:"dependsOn,omitempty"`
	Replaces   float64                           `json:"replaces,omitempty"`
	Done       []*AiChatPlanStepDoneCheckWireDto `json:"done,omitempty"`
	Loop       *AiChatPlanStepLoopWireDto        `json:"loop,omitempty"`
	Difficulty float64                           `json:"difficulty,omitempty"`
}

// PlanEntryWireDto DTO.
type PlanEntryWireDto struct {
	AiChatEntryWireDto
	Kind             string                   `json:"kind,omitempty"`
	Goal             string                   `json:"goal,omitempty"`
	Steps            []*AiChatPlanStepWireDto `json:"steps,omitempty"`
	Status           string                   `json:"status,omitempty"`
	Gate             *AiChatGateResultWireDto `json:"gate,omitempty"`
	Difficulty       float64                  `json:"difficulty,omitempty"`
	DifficultyReason string                   `json:"difficultyReason,omitempty"`
	DeltaOf          string                   `json:"deltaOf,omitempty"`
}

// UserDecisionEntryWireDto DTO.
type UserDecisionEntryWireDto struct {
	AiChatEntryWireDto
	Kind     string `json:"kind,omitempty"`
	Decision string `json:"decision,omitempty"`
	Comment  string `json:"comment,omitempty"`
}

// AiChatStepLogLineWireDto DTO.
type AiChatStepLogLineWireDto struct {
	Seq    float64 `json:"seq,omitempty"`
	Tool   string  `json:"tool,omitempty"`
	Agent  string  `json:"agent,omitempty"`
	Status string  `json:"status,omitempty"`
	Detail string  `json:"detail,omitempty"`
}

// RunStepEntryWireDto DTO.
type RunStepEntryWireDto struct {
	AiChatEntryWireDto
	Kind          string                      `json:"kind,omitempty"`
	N             float64                     `json:"n,omitempty"`
	Tool          string                      `json:"tool,omitempty"`
	Title         string                      `json:"title,omitempty"`
	Status        string                      `json:"status,omitempty"`
	ResultSummary string                      `json:"resultSummary,omitempty"`
	Error         string                      `json:"error,omitempty"`
	Log           []*AiChatStepLogLineWireDto `json:"log,omitempty"`
}

// ActionPendingEntryWireDto DTO.
type ActionPendingEntryWireDto struct {
	AiChatEntryWireDto
	Kind          string `json:"kind,omitempty"`
	Tool          string `json:"tool,omitempty"`
	ArgumentsJson string `json:"argumentsJson,omitempty"`
	Status        string `json:"status,omitempty"`
}

// NoticeEntryWireDto DTO.
type NoticeEntryWireDto struct {
	AiChatEntryWireDto
	Kind  string `json:"kind,omitempty"`
	Text  string `json:"text,omitempty"`
	Level string `json:"level,omitempty"`
}

// ConversationSnapshotEntryWireDto DTO.
type ConversationSnapshotEntryWireDto struct {
	AiChatEntryWireDto
	Kind          string  `json:"kind,omitempty"`
	SnapshotId    string  `json:"snapshotId,omitempty"`
	CoversUpToSeq float64 `json:"coversUpToSeq,omitempty"`
}

// ICultureBasedRequest DTO.
type ICultureBasedRequest struct {
}

// IVersionBasedRequest DTO.
type IVersionBasedRequest struct {
}

// IHasCorrelationIdRequest DTO.
type IHasCorrelationIdRequest struct {
}

// IHasAccountId DTO.
type IHasAccountId struct {
}

// CodeMashRequestBase DTO.
type CodeMashRequestBase struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Env       string `json:"env,omitempty"`
}

// IHasProjectId DTO.
type IHasProjectId struct {
}

// IHasEnv DTO.
type IHasEnv struct {
}

// TagDescriptionDto DTO.
type TagDescriptionDto struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// TagTranslationDto DTO.
type TagTranslationDto struct {
	Language string             `json:"language,omitempty"`
	Content  *TagDescriptionDto `json:"content,omitempty"`
}

// TagDefinitionBaseDto DTO.
type TagDefinitionBaseDto struct {
	Tag          string               `json:"tag,omitempty"`
	Translations []*TagTranslationDto `json:"translations,omitempty"`
}

// GroupDefinitionDto DTO.
type GroupDefinitionDto struct {
	TagDefinitionBaseDto
}

// TagDefinitionDto DTO.
type TagDefinitionDto struct {
	TagDefinitionBaseDto
	DefaultDelivery map[string]bool `json:"defaultDelivery,omitempty"`
}

// EmailAddress DTO.
type EmailAddress struct {
	Address string `json:"address,omitempty"`
}

// AggregateId DTO.
type AggregateId struct {
	Value string `json:"value,omitempty"`
}

// AccountId DTO.
type AccountId struct {
	AggregateId
}

// UtcDateTime DTO.
type UtcDateTime struct {
}

// ExpirationToken DTO.
type ExpirationToken struct {
	Items float64  `json:"items,omitempty"`
	Unit  TimeUnit `json:"unit,omitempty"`
	Value float64  `json:"value,omitempty"`
}

// CodeMashSubscriptionId DTO.
type CodeMashSubscriptionId struct {
	AggregateId
}

// ProjectId DTO.
type ProjectId struct {
	AggregateId
}

// IntegrationId DTO.
type IntegrationId struct {
	AggregateId
}

// ResourceRef DTO.
type ResourceRef struct {
	ProjectId     *ProjectId      `json:"projectId,omitempty"`
	IntegrationId *IntegrationId  `json:"integrationId,omitempty"`
	Kind          ResourceRefKind `json:"kind,omitempty"`
}

// PaymentCustomerRef DTO.
type PaymentCustomerRef struct {
	ResourceRef
	Kind       ResourceRefKind `json:"kind,omitempty"`
	Source     ResourceSource  `json:"source,omitempty"`
	ExternalId string          `json:"externalId,omitempty"`
}

// Quantity DTO.
type Quantity struct {
	Value float64 `json:"value,omitempty"`
}

// CodeMashManagedServiceSubscription DTO.
type CodeMashManagedServiceSubscription struct {
	SubscriptionId     *CodeMashSubscriptionId `json:"subscriptionId,omitempty"`
	PaymentCustomerRef *PaymentCustomerRef     `json:"paymentCustomerRef,omitempty"`
	RefSubscriptionId  string                  `json:"refSubscriptionId,omitempty"`
	IssuedOn           *UtcDateTime            `json:"issuedOn,omitempty"`
	WillExpireOn       *UtcDateTime            `json:"willExpireOn,omitempty"`
	ProjectCap         *Quantity               `json:"projectCap,omitempty"`
	IsTrial            bool                    `json:"isTrial,omitempty"`
}

// DomainUrl DTO.
type DomainUrl struct {
	Value string `json:"value,omitempty"`
}

// CodeMashLicense DTO.
type CodeMashLicense struct {
	CodeMashManagedServiceSubscription
	Domain       *DomainUrl `json:"domain,omitempty"`
	AccountId    *AccountId `json:"accountId,omitempty"`
	IsEnterprise bool       `json:"isEnterprise,omitempty"`
}

// ProjectName DTO.
type ProjectName struct {
	Name       string `json:"name,omitempty"`
	UniqueName string `json:"uniqueName,omitempty"`
}

// NorbixRegion DTO.
type NorbixRegion struct {
	Code string `json:"code,omitempty"`
}

// ProjectRegion DTO.
type ProjectRegion struct {
	Region    *NorbixRegion `json:"region,omitempty"`
	Name      string        `json:"name,omitempty"`
	Continent Continent     `json:"continent,omitempty"`
}

// ProjectLegalDocuments DTO.
type ProjectLegalDocuments struct {
	TermsMarkdown   string `json:"termsMarkdown,omitempty"`
	PrivacyMarkdown string `json:"privacyMarkdown,omitempty"`
}

// AuthId DTO.
type AuthId struct {
	Value string `json:"value,omitempty"`
}

// Language DTO.
type Language struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// FileResourceId DTO.
type FileResourceId struct {
	Value string `json:"value,omitempty"`
}

// FileChecksum DTO.
type FileChecksum struct {
	Algorithm string `json:"algorithm,omitempty"`
	Hash      string `json:"hash,omitempty"`
}

// FileResource DTO.
type FileResource struct {
	Id               *FileResourceId `json:"id,omitempty"`
	OriginalFileName string          `json:"originalFileName,omitempty"`
	Extension        string          `json:"extension,omitempty"`
	SizeBytes        float64         `json:"sizeBytes,omitempty"`
	Checksum         *FileChecksum   `json:"checksum,omitempty"`
	StoredFileName   string          `json:"storedFileName,omitempty"`
}

// FileResourceRef DTO.
type FileResourceRef struct {
	Resource      *FileResource  `json:"resource,omitempty"`
	IntegrationId *IntegrationId `json:"integrationId,omitempty"`
	Provider      FileProvider   `json:"provider,omitempty"`
	Path          string         `json:"path,omitempty"`
}

// ProjectLogo DTO.
type ProjectLogo struct {
	FileResource *FileResourceRef `json:"fileResource,omitempty"`
	PublicUrl    string           `json:"publicUrl,omitempty"`
}

// ProjectIcon DTO.
type ProjectIcon struct {
	FileResource *FileResourceRef `json:"fileResource,omitempty"`
	PublicUrl    string           `json:"publicUrl,omitempty"`
}

// BrandColor DTO.
type BrandColor struct {
	Value string `json:"value,omitempty"`
}

// Tag DTO.
type Tag struct {
}

// GroupTags DTO.
type GroupTags struct {
	Group *Tag   `json:"group,omitempty"`
	Tags  []*Tag `json:"tags,omitempty"`
}

// ProjectCommunicationChannel DTO.
type ProjectCommunicationChannel struct {
	Channel CommunicationChannel `json:"channel,omitempty"`
	Groups  []*GroupTags         `json:"groups,omitempty"`
}

// TagDescription DTO.
type TagDescription struct {
	DisplayName *DisplayName `json:"displayName,omitempty"`
	Description string       `json:"description,omitempty"`
}

// MessageTranslation DTO.
type MessageTranslation[TContent any] struct {
}

// TagTranslation DTO.
type TagTranslation struct {
	MessageTranslation[*TagDescription]
}

// BaseTagDefinition DTO.
type BaseTagDefinition struct {
	Tag          *Tag              `json:"tag,omitempty"`
	Translations []*TagTranslation `json:"translations,omitempty"`
}

// GroupDefinition DTO.
type GroupDefinition struct {
	BaseTagDefinition
}

// TagDefinition DTO.
type TagDefinition struct {
	BaseTagDefinition
	DefaultDelivery map[string]bool `json:"defaultDelivery,omitempty"`
}

// ProjectCommunication DTO.
type ProjectCommunication struct {
	Channels []*ProjectCommunicationChannel `json:"channels,omitempty"`
	Groups   []*GroupDefinition             `json:"groups,omitempty"`
	Tags     []*TagDefinition               `json:"tags,omitempty"`
}

// TimeZone DTO.
type TimeZone struct {
	ZoneId string `json:"zoneId,omitempty"`
}

// PolicyId DTO.
type PolicyId struct {
	Template           string `json:"template,omitempty"`
	TenancyScopeViewId string `json:"tenancyScopeViewId,omitempty"`
	ViewId             string `json:"viewId,omitempty"`
	IsSystem           bool   `json:"isSystem,omitempty"`
}

// PermissionAction DTO.
type PermissionAction struct {
	Module              ApplicationModule `json:"module,omitempty"`
	Operation           string            `json:"operation,omitempty"`
	IsModuleWildcard    bool              `json:"isModuleWildcard,omitempty"`
	IsOperationWildcard bool              `json:"isOperationWildcard,omitempty"`
	IsConcrete          bool              `json:"isConcrete,omitempty"`
	Specificity         float64           `json:"specificity,omitempty"`
}

// ResourceKind DTO.
type ResourceKind struct {
	Name string `json:"name,omitempty"`
}

// ResourceIdentifier DTO.
type ResourceIdentifier struct {
	Value string `json:"value,omitempty"`
}

// ResourcePattern DTO.
type ResourcePattern struct {
	Account           *AccountId          `json:"account,omitempty"`
	Project           *ProjectId          `json:"project,omitempty"`
	Module            ApplicationModule   `json:"module,omitempty"`
	Kind              *ResourceKind       `json:"kind,omitempty"`
	Id                *ResourceIdentifier `json:"id,omitempty"`
	IsAccountWildcard bool                `json:"isAccountWildcard,omitempty"`
	IsProjectWildcard bool                `json:"isProjectWildcard,omitempty"`
	IsModuleWildcard  bool                `json:"isModuleWildcard,omitempty"`
	IsKindWildcard    bool                `json:"isKindWildcard,omitempty"`
	IsIdWildcard      bool                `json:"isIdWildcard,omitempty"`
	IsConcrete        bool                `json:"isConcrete,omitempty"`
	IsFullWildcard    bool                `json:"isFullWildcard,omitempty"`
	Specificity       float64             `json:"specificity,omitempty"`
}

// Permission DTO.
type Permission struct {
	Sid       string              `json:"sid,omitempty"`
	Effect    PermissionEffect    `json:"effect,omitempty"`
	Actions   []*PermissionAction `json:"actions,omitempty"`
	Resources []*ResourcePattern  `json:"resources,omitempty"`
}

// MembershipPolicy DTO.
type MembershipPolicy struct {
	Id          *PolicyId     `json:"id,omitempty"`
	Name        *DisplayName  `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Permissions []*Permission `json:"permissions,omitempty"`
	Disabled    bool          `json:"disabled,omitempty"`
	IsSystem    bool          `json:"isSystem,omitempty"`
}

// RoleId DTO.
type RoleId struct {
	Template           string `json:"template,omitempty"`
	TenancyScopeViewId string `json:"tenancyScopeViewId,omitempty"`
	ViewId             string `json:"viewId,omitempty"`
	IsSystem           bool   `json:"isSystem,omitempty"`
}

// MembershipRole DTO.
type MembershipRole struct {
	Id               *RoleId      `json:"id,omitempty"`
	Name             *DisplayName `json:"name,omitempty"`
	Description      string       `json:"description,omitempty"`
	AttachedPolicies []*PolicyId  `json:"attachedPolicies,omitempty"`
	Disabled         bool         `json:"disabled,omitempty"`
	IsSystem         bool         `json:"isSystem,omitempty"`
}

// BillingPeriod DTO.
type BillingPeriod struct {
	Year            float64 `json:"year,omitempty"`
	Month           float64 `json:"month,omitempty"`
	StartUtc        string  `json:"startUtc,omitempty"`
	EndExclusiveUtc string  `json:"endExclusiveUtc,omitempty"`
	LastInstantUtc  string  `json:"lastInstantUtc,omitempty"`
}

// AtlasClusterChargeRecord DTO.
type AtlasClusterChargeRecord struct {
	AtlasProjectId   string  `json:"atlasProjectId,omitempty"`
	AtlasClusterName string  `json:"atlasClusterName,omitempty"`
	Cents            float64 `json:"cents,omitempty"`
}

// AtlasUsageRecord DTO.
type AtlasUsageRecord struct {
	Period        *BillingPeriod              `json:"period,omitempty"`
	TotalCents    float64                     `json:"totalCents,omitempty"`
	PerCluster    []*AtlasClusterChargeRecord `json:"perCluster,omitempty"`
	RecordedAtUtc *UtcDateTime                `json:"recordedAtUtc,omitempty"`
}

// UsageIngestionFailure DTO.
type UsageIngestionFailure struct {
	Reason        UsageIngestionFailureReason `json:"reason,omitempty"`
	Period        *BillingPeriod              `json:"period,omitempty"`
	StripeEventId string                      `json:"stripeEventId,omitempty"`
	Message       string                      `json:"message,omitempty"`
	ReportedAtUtc *UtcDateTime                `json:"reportedAtUtc,omitempty"`
}

// DeleteTrigger DTO.
type DeleteTrigger struct {
	CodeMashRequestBase
	TriggerId   string      `json:"triggerId,omitempty"`
	TriggerType TriggerType `json:"triggerType,omitempty"`
	SchemaId    string      `json:"schemaId,omitempty"`
}

// DisableTrigger DTO.
type DisableTrigger struct {
	CodeMashRequestBase
	TriggerId   string      `json:"triggerId,omitempty"`
	TriggerType TriggerType `json:"triggerType,omitempty"`
	SchemaId    string      `json:"schemaId,omitempty"`
}

// EnableTrigger DTO.
type EnableTrigger struct {
	CodeMashRequestBase
	TriggerId   string      `json:"triggerId,omitempty"`
	TriggerType TriggerType `json:"triggerType,omitempty"`
	SchemaId    string      `json:"schemaId,omitempty"`
}

// GetTrigger DTO.
type GetTrigger struct {
	CodeMashRequestBase
	Id       string `json:"id,omitempty"`
	SchemaId string `json:"schemaId,omitempty"`
}

// SaveTrigger DTO.
type SaveTrigger struct {
	CodeMashRequestBase
	Trigger *SaveTriggerRequest `json:"trigger,omitempty"`
}

// CredentialsSettingsModeDto DTO.
type CredentialsSettingsModeDto struct {
	Name      string `json:"name,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
}

// Integration DTO.
type Integration struct {
	IntegrationId                    *IntegrationId `json:"integrationId,omitempty"`
	Env                              *Env           `json:"env,omitempty"`
	Capability                       string         `json:"capability,omitempty"`
	IsSystemOwned                    bool           `json:"isSystemOwned,omitempty"`
	IntegrationName                  *DisplayName   `json:"integrationName,omitempty"`
	IsEnabled                        bool           `json:"isEnabled,omitempty"`
	IsConfigured                     bool           `json:"isConfigured,omitempty"`
	LastIntegrationTestAtUtc         string         `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded     bool           `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrorMessages []string       `json:"lastIntegrationTestErrorMessages,omitempty"`
	HumanDeliveryConfirmedAtUtc      string         `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	IsApprovedThatItWorks            bool           `json:"isApprovedThatItWorks,omitempty"`
}

// MembershipIntegration DTO.
type MembershipIntegration struct {
	Integration
	Provider MembershipProvider `json:"provider,omitempty"`
}

// TriggerId DTO.
type TriggerId struct {
	AggregateId
}

// TriggerAction DTO.
type TriggerAction struct {
	Type          TriggerActionType `json:"type,omitempty"`
	IntegrationId *IntegrationId    `json:"integrationId,omitempty"`
}

// TemplateCode DTO.
type TemplateCode struct {
}

// Trigger DTO.
type Trigger struct {
	TriggerId      *TriggerId     `json:"triggerId,omitempty"`
	Name           *DisplayName   `json:"name,omitempty"`
	TriggerAction  *TriggerAction `json:"triggerAction,omitempty"`
	ActivationCode *TemplateCode  `json:"activationCode,omitempty"`
	Description    string         `json:"description,omitempty"`
	IsEnabled      bool           `json:"isEnabled,omitempty"`
	Env            *Env           `json:"env,omitempty"`
	IntegrationId  *IntegrationId `json:"integrationId,omitempty"`
}

// MembershipTrigger DTO.
type MembershipTrigger struct {
	Trigger
	When MembershipTriggerType `json:"when,omitempty"`
}

// TriggerByIdEventBase DTO.
type TriggerByIdEventBase struct {
	TriggerId *TriggerId `json:"triggerId,omitempty"`
}

// SchemaSettingsDto DTO.
type SchemaSettingsDto struct {
	SoftDelete     bool   `json:"softDelete,omitempty"`
	HasRecordOwner bool   `json:"hasRecordOwner,omitempty"`
	Description    string `json:"description,omitempty"`
}

// SchemaListColumnDto DTO.
type SchemaListColumnDto struct {
	Field string `json:"field,omitempty"`
}

// SchemaListSortDto DTO.
type SchemaListSortDto struct {
	Field string  `json:"field,omitempty"`
	Order float64 `json:"order,omitempty"`
}

// SchemaListSettingsDto DTO.
type SchemaListSettingsDto struct {
	Columns     []*SchemaListColumnDto `json:"columns,omitempty"`
	DefaultSort *SchemaListSortDto     `json:"defaultSort,omitempty"`
}

// ImportColumnMappingDto DTO.
type ImportColumnMappingDto struct {
	CsvColumnIndex    float64 `json:"csvColumnIndex,omitempty"`
	CsvHeader         string  `json:"csvHeader,omitempty"`
	PropertyName      string  `json:"propertyName,omitempty"`
	DontImportOnError bool    `json:"dontImportOnError,omitempty"`
}

// MongoDbAggregateId DTO.
type MongoDbAggregateId struct {
	AggregateId
}

// MongoDbAggregateQuery DTO.
type MongoDbAggregateQuery struct {
	Value string `json:"value,omitempty"`
}

// SchemaId DTO.
type SchemaId struct {
	AggregateId
}

// MongoDbAggregate DTO.
type MongoDbAggregate struct {
	Id          *MongoDbAggregateId    `json:"id,omitempty"`
	DisplayName *DisplayName           `json:"displayName,omitempty"`
	Description string                 `json:"description,omitempty"`
	Query       *MongoDbAggregateQuery `json:"query,omitempty"`
	SchemaId    *SchemaId              `json:"schemaId,omitempty"`
}

// DatabaseIntegration DTO.
type DatabaseIntegration struct {
	Integration
	Provider         DatabaseProvider  `json:"provider,omitempty"`
	Status           IntegrationStatus `json:"status,omitempty"`
	AtlasProjectId   string            `json:"atlasProjectId,omitempty"`
	AtlasClusterName string            `json:"atlasClusterName,omitempty"`
	FailureReason    string            `json:"failureReason,omitempty"`
}

// SchemaName DTO.
type SchemaName struct {
	Value string `json:"value,omitempty"`
	Title string `json:"title,omitempty"`
}

// JsonSchemaFieldName DTO.
type JsonSchemaFieldName struct {
	FieldName string `json:"fieldName,omitempty"`
}

// JsonSchemaField DTO.
type JsonSchemaField struct {
	FieldName *JsonSchemaFieldName `json:"fieldName,omitempty"`
}

// DataSchema DTO.
type DataSchema struct {
	RawJson string             `json:"rawJson,omitempty"`
	Fields  []*JsonSchemaField `json:"fields,omitempty"`
}

// VisualSchema DTO.
type VisualSchema struct {
	RawJson string `json:"rawJson,omitempty"`
}

// SchemaDraft DTO.
type SchemaDraft struct {
	DataSchema   *DataSchema   `json:"dataSchema,omitempty"`
	VisualSchema *VisualSchema `json:"visualSchema,omitempty"`
	UpdatedAt    string        `json:"updatedAt,omitempty"`
}

// SchemaVersion DTO.
type SchemaVersion struct {
	Value float64 `json:"value,omitempty"`
}

// MetaSchemaVersion DTO.
type MetaSchemaVersion struct {
	Value float64 `json:"value,omitempty"`
}

// PublishedSchemaVersion DTO.
type PublishedSchemaVersion struct {
	Version           *SchemaVersion     `json:"version,omitempty"`
	DataSchema        *DataSchema        `json:"dataSchema,omitempty"`
	VisualSchema      *VisualSchema      `json:"visualSchema,omitempty"`
	MetaSchemaVersion *MetaSchemaVersion `json:"metaSchemaVersion,omitempty"`
	PublishedAt       string             `json:"publishedAt,omitempty"`
}

// SchemaSettings DTO.
type SchemaSettings struct {
	SoftDelete     bool   `json:"softDelete,omitempty"`
	HasRecordOwner bool   `json:"hasRecordOwner,omitempty"`
	Description    string `json:"description,omitempty"`
}

// Schema DTO.
type Schema struct {
	SchemaName        *SchemaName               `json:"schemaName,omitempty"`
	Id                *SchemaId                 `json:"id,omitempty"`
	Env               *Env                      `json:"env,omitempty"`
	Draft             *SchemaDraft              `json:"draft,omitempty"`
	PublishedVersions []*PublishedSchemaVersion `json:"publishedVersions,omitempty"`
	Triggers          []*Trigger                `json:"triggers,omitempty"`
	Settings          *SchemaSettings           `json:"settings,omitempty"`
}

// SchemaDiff DTO.
type SchemaDiff struct {
	AddedFields              []string `json:"addedFields,omitempty"`
	RemovedFields            []string `json:"removedFields,omitempty"`
	TypeChangedFields        []string `json:"typeChangedFields,omitempty"`
	ValidatorTightenedFields []string `json:"validatorTightenedFields,omitempty"`
	IsEmpty                  bool     `json:"isEmpty,omitempty"`
}

// TaxonomyId DTO.
type TaxonomyId struct {
	AggregateId
}

// TaxonomyName DTO.
type TaxonomyName struct {
	Value string `json:"value,omitempty"`
	Title string `json:"title,omitempty"`
}

// RecordId DTO.
type RecordId struct {
	Id string `json:"id,omitempty"`
}

// Taxonomy DTO.
type Taxonomy struct {
	ParentId              *TaxonomyId   `json:"parentId,omitempty"`
	Id                    *TaxonomyId   `json:"id,omitempty"`
	Name                  *TaxonomyName `json:"name,omitempty"`
	Description           string        `json:"description,omitempty"`
	TermsMetaVisualSchema *VisualSchema `json:"termsMetaVisualSchema,omitempty"`
	TermsMetaDataSchema   *DataSchema   `json:"termsMetaDataSchema,omitempty"`
	Dependencies          []*TaxonomyId `json:"dependencies,omitempty"`
	RecordId              *RecordId     `json:"recordId,omitempty"`
}

// SchemaTrigger DTO.
type SchemaTrigger struct {
	Trigger
	SchemaId      *SchemaId         `json:"schemaId,omitempty"`
	When          SchemaTriggerType `json:"when,omitempty"`
	Configuration *TemplateCode     `json:"configuration,omitempty"`
}

// IPasskeyMessage DTO.
type IPasskeyMessage struct {
}

// AuthUserName DTO.
type AuthUserName struct {
	Value string `json:"value,omitempty"`
}

// IpAddress DTO.
type IpAddress struct {
	Ip string `json:"ip,omitempty"`
}

// AccessInformation DTO.
type AccessInformation struct {
	Ip   *IpAddress   `json:"ip,omitempty"`
	Date *UtcDateTime `json:"date,omitempty"`
	Zone *TimeZone    `json:"zone,omitempty"`
}

// Registration DTO.
type Registration struct {
	RegistrationInformation *AccessInformation `json:"registrationInformation,omitempty"`
}

// Login DTO.
type Login struct {
	NeedChangePasswordOnNextLogin bool               `json:"needChangePasswordOnNextLogin,omitempty"`
	LastAccessInformation         *AccessInformation `json:"lastAccessInformation,omitempty"`
}

// Phone DTO.
type Phone struct {
	Value string `json:"value,omitempty"`
}

// FirstName DTO.
type FirstName struct {
	Value string `json:"value,omitempty"`
}

// LastName DTO.
type LastName struct {
	Value string `json:"value,omitempty"`
}

// MidName DTO.
type MidName struct {
	Value string `json:"value,omitempty"`
}

// FullName DTO.
type FullName struct {
	FirstName *FirstName `json:"firstName,omitempty"`
	MidName   *MidName   `json:"midName,omitempty"`
	LastName  *LastName  `json:"lastName,omitempty"`
	Title     string     `json:"title,omitempty"`
}

// City DTO.
type City struct {
	Value string `json:"value,omitempty"`
}

// Country DTO.
type Country struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// AddressLine DTO.
type AddressLine struct {
	Value string `json:"value,omitempty"`
}

// PostalCode DTO.
type PostalCode struct {
	Value string `json:"value,omitempty"`
}

// CountryState DTO.
type CountryState struct {
	Value string `json:"value,omitempty"`
}

// Address DTO.
type Address struct {
	City         *City         `json:"city,omitempty"`
	Country      *Country      `json:"country,omitempty"`
	AddressLine1 *AddressLine  `json:"addressLine1,omitempty"`
	AddressLine2 *AddressLine  `json:"addressLine2,omitempty"`
	PostalCode   *PostalCode   `json:"postalCode,omitempty"`
	State        *CountryState `json:"state,omitempty"`
}

// UserMarketingPreferences DTO.
type UserMarketingPreferences struct {
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]*Tag      `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
}

// UserGeneralInfo DTO.
type UserGeneralInfo struct {
	Phone                *Phone                    `json:"phone,omitempty"`
	PrimaryEmail         *EmailAddress             `json:"primaryEmail,omitempty"`
	DisplayName          *DisplayName              `json:"displayName,omitempty"`
	FirstName            *FirstName                `json:"firstName,omitempty"`
	LastName             *LastName                 `json:"lastName,omitempty"`
	FullName             *FullName                 `json:"fullName,omitempty"`
	Address              *Address                  `json:"address,omitempty"`
	Company              string                    `json:"company,omitempty"`
	Gender               Gender                    `json:"gender,omitempty"`
	BirthDate            *UtcDateTime              `json:"birthDate,omitempty"`
	TimeZone             *TimeZone                 `json:"timeZone,omitempty"`
	Language             *Language                 `json:"language,omitempty"`
	MarketingPreferences *UserMarketingPreferences `json:"marketingPreferences,omitempty"`
	Notes                string                    `json:"notes,omitempty"`
	ExtraMetadata        string                    `json:"extraMetadata,omitempty"`
}

// DeviceId DTO.
type DeviceId struct {
	Id string `json:"id,omitempty"`
}

// PushDeviceToken DTO.
type PushDeviceToken struct {
	Token string `json:"token,omitempty"`
}

// PushDeviceDeliveryToken DTO.
type PushDeviceDeliveryToken struct {
	PushDeviceToken *PushDeviceToken         `json:"pushDeviceToken,omitempty"`
	DeliveryFamily  PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
}

// PushDevice DTO.
type PushDevice struct {
	Id               *DeviceId                `json:"id,omitempty"`
	Brand            string                   `json:"brand,omitempty"`
	Manufacturer     string                   `json:"manufacturer,omitempty"`
	ModelName        string                   `json:"modelName,omitempty"`
	DeviceName       string                   `json:"deviceName,omitempty"`
	DeviceType       DeviceType               `json:"deviceType,omitempty"`
	OsName           string                   `json:"osName,omitempty"`
	OsVersion        string                   `json:"osVersion,omitempty"`
	PlatformApiLevel float64                  `json:"platformApiLevel,omitempty"`
	Token            *PushDeviceDeliveryToken `json:"token,omitempty"`
}

// PushDevices DTO.
type PushDevices []*PushDevice

// UserId DTO.
type UserId struct {
	Value string `json:"value,omitempty"`
}

// UserRef DTO.
type UserRef struct {
	ResourceRef
	Kind   ResourceRefKind `json:"kind,omitempty"`
	UserId *UserId         `json:"userId,omitempty"`
}

// Auth DTO.
type Auth struct {
	Id           *AuthId          `json:"id,omitempty"`
	Roles        []*RoleName      `json:"roles,omitempty"`
	Email        *EmailAddress    `json:"email,omitempty"`
	UserName     *AuthUserName    `json:"userName,omitempty"`
	Type         AuthType         `json:"type,omitempty"`
	Registration *Registration    `json:"registration,omitempty"`
	Login        *Login           `json:"login,omitempty"`
	GeneralInfo  *UserGeneralInfo `json:"generalInfo,omitempty"`
	Status       AuthStatus       `json:"status,omitempty"`
	CreatedOn    *UtcDateTime     `json:"createdOn,omitempty"`
	ModifiedOn   *UtcDateTime     `json:"modifiedOn,omitempty"`
	PushDevices  *PushDevices     `json:"pushDevices,omitempty"`
	Tags         []*Tag           `json:"tags,omitempty"`
	UserRef      *UserRef         `json:"userRef,omitempty"`
}

// FileIntegration DTO.
type FileIntegration struct {
	Integration
	Provider FileProvider `json:"provider,omitempty"`
}

// FileTrigger DTO.
type FileTrigger struct {
	Trigger
	When            FilesTriggerType `json:"when,omitempty"`
	FileResourceRef *FileResourceRef `json:"fileResourceRef,omitempty"`
}

// EmailValidationIntegrationRequest DTO.
type EmailValidationIntegrationRequest struct {
	IntegrationId   string                  `json:"integrationId,omitempty"`
	Provider        EmailValidationProvider `json:"provider,omitempty"`
	IntegrationName string                  `json:"integrationName,omitempty"`
	IsEnabled       bool                    `json:"isEnabled,omitempty"`
}

// SaveEmailTemplate DTO.
type SaveEmailTemplate struct {
	CodeMashRequestBase
	TemplateName         string                        `json:"templateName,omitempty"`
	Description          string                        `json:"description,omitempty"`
	CommunicationChannel CommunicationChannel          `json:"communicationChannel,omitempty"`
	Tags                 []string                      `json:"tags,omitempty"`
	StaticAttachments    []*FileResourceRefDto         `json:"staticAttachments,omitempty"`
	Translations         []*EmailMessageTranslationDto `json:"translations,omitempty"`
}

// TranslationDto DTO.
type TranslationDto struct {
	Language string `json:"language,omitempty"`
	Content  string `json:"content,omitempty"`
}

// EmailFooterId DTO.
type EmailFooterId struct {
	Value string `json:"value,omitempty"`
}

// EmailFooter DTO.
type EmailFooter struct {
	Id           *EmailFooterId                      `json:"id,omitempty"`
	DisplayName  *DisplayName                        `json:"displayName,omitempty"`
	Translations []MessageTranslation[*TemplateCode] `json:"translations,omitempty"`
	Env          *Env                                `json:"env,omitempty"`
}

// EmailSenderName DTO.
type EmailSenderName struct {
}

// EmailIntegration DTO.
type EmailIntegration struct {
	Integration
	Provider        EmailProvider    `json:"provider,omitempty"`
	EmailAddress    *EmailAddress    `json:"emailAddress,omitempty"`
	EmailSenderName *EmailSenderName `json:"emailSenderName,omitempty"`
}

// EmailSignatureId DTO.
type EmailSignatureId struct {
	Value string `json:"value,omitempty"`
}

// EmailSignature DTO.
type EmailSignature struct {
	Id           *EmailSignatureId                   `json:"id,omitempty"`
	DisplayName  *DisplayName                        `json:"displayName,omitempty"`
	Translations []MessageTranslation[*TemplateCode] `json:"translations,omitempty"`
	Env          *Env                                `json:"env,omitempty"`
}

// TemplateId DTO.
type TemplateId struct {
	Value string `json:"value,omitempty"`
}

// Template DTO.
type Template[TMessageContent any] struct {
	TemplateId           *TemplateId                           `json:"templateId,omitempty"`
	TemplateName         *DisplayName                          `json:"templateName,omitempty"`
	Translations         []MessageTranslation[TMessageContent] `json:"translations,omitempty"`
	CommunicationChannel CommunicationChannel                  `json:"communicationChannel,omitempty"`
	IsActive             bool                                  `json:"isActive,omitempty"`
	Description          string                                `json:"description,omitempty"`
	Tags                 []*Tag                                `json:"tags,omitempty"`
	FileIntegrationId    *IntegrationId                        `json:"fileIntegrationId,omitempty"`
	Env                  *Env                                  `json:"env,omitempty"`
}

// EmailSubject DTO.
type EmailSubject struct {
}

// EmailBody DTO.
type EmailBody struct {
	Code                *TemplateCode       `json:"code,omitempty"`
	Structure           string              `json:"structure,omitempty"`
	EmailTemplateEngine EmailTemplateEngine `json:"emailTemplateEngine,omitempty"`
}

// EmailMessageContent DTO.
type EmailMessageContent struct {
	Subject           *EmailSubject      `json:"subject,omitempty"`
	Body              *EmailBody         `json:"body,omitempty"`
	StaticAttachments []*FileResourceRef `json:"staticAttachments,omitempty"`
}

// EmailTemplate DTO.
type EmailTemplate struct {
	Template[*EmailMessageContent]
	StaticAttachments []*FileResourceRef `json:"staticAttachments,omitempty"`
}

// EmailValidationIntegration DTO.
type EmailValidationIntegration struct {
	Integration
	Provider EmailValidationProvider `json:"provider,omitempty"`
}

// CampaignId DTO.
type CampaignId struct {
	Id string `json:"id,omitempty"`
}

// CampaignBatchId DTO.
type CampaignBatchId struct {
	Id string `json:"id,omitempty"`
}

// NotificationId DTO.
type NotificationId struct {
	AggregateId
}

// SaveSmsTemplate DTO.
type SaveSmsTemplate struct {
	CodeMashRequestBase
	TemplateName         string                      `json:"templateName,omitempty"`
	Description          string                      `json:"description,omitempty"`
	CommunicationChannel CommunicationChannel        `json:"communicationChannel,omitempty"`
	Tags                 []string                    `json:"tags,omitempty"`
	Translations         []*SmsMessageTranslationDto `json:"translations,omitempty"`
}

// SmsIntegrationRequest DTO.
type SmsIntegrationRequest struct {
	IntegrationId   string      `json:"integrationId,omitempty"`
	Provider        SmsProvider `json:"provider,omitempty"`
	IntegrationName string      `json:"integrationName,omitempty"`
	IsEnabled       bool        `json:"isEnabled,omitempty"`
}

// SmsIntegration DTO.
type SmsIntegration struct {
	Integration
	Provider SmsProvider `json:"provider,omitempty"`
}

// SmsTitle DTO.
type SmsTitle struct {
	Value *TemplateCode `json:"value,omitempty"`
}

// SmsBody DTO.
type SmsBody struct {
	Value *TemplateCode `json:"value,omitempty"`
}

// SmsMessageContent DTO.
type SmsMessageContent struct {
	Title *SmsTitle `json:"title,omitempty"`
	Body  *SmsBody  `json:"body,omitempty"`
}

// SmsTemplate DTO.
type SmsTemplate struct {
	Template[*SmsMessageContent]
}

// CodeIntegration DTO.
type CodeIntegration struct {
	Integration
	Provider CodeProvider `json:"provider,omitempty"`
}

// MarketplaceTokenMapping DTO.
type MarketplaceTokenMapping struct {
	Token      string                   `json:"token,omitempty"`
	Resolver   MarketplaceTokenResolver `json:"resolver,omitempty"`
	Value      string                   `json:"value,omitempty"`
	SecretKeys []string                 `json:"secretKeys,omitempty"`
	Format     SecretValueFormat        `json:"format,omitempty"`
}

// MarketplaceIntegration DTO.
type MarketplaceIntegration struct {
	Integration
	Capability    string                          `json:"capability,omitempty"`
	ListingViewId string                          `json:"listingViewId,omitempty"`
	Transport     MarketplaceIntegrationTransport `json:"transport,omitempty"`
	Vendor        string                          `json:"vendor,omitempty"`
	Category      MarketplaceIntegrationCategory  `json:"category,omitempty"`
	Description   string                          `json:"description,omitempty"`
	Config        map[string]string               `json:"config,omitempty"`
	TokenMappings []*MarketplaceTokenMapping      `json:"tokenMappings,omitempty"`
}

// MarketplaceFunctionId DTO.
type MarketplaceFunctionId struct {
	Value string `json:"value,omitempty"`
}

// MarketplaceFunction DTO.
type MarketplaceFunction struct {
	FunctionId      *MarketplaceFunctionId     `json:"functionId,omitempty"`
	IntegrationId   *IntegrationId             `json:"integrationId,omitempty"`
	Env             *Env                       `json:"env,omitempty"`
	FunctionKey     string                     `json:"functionKey,omitempty"`
	DisplayName     *DisplayName               `json:"displayName,omitempty"`
	Description     string                     `json:"description,omitempty"`
	IsEnabled       bool                       `json:"isEnabled,omitempty"`
	RequestTemplate string                     `json:"requestTemplate,omitempty"`
	MappedTokens    []*MarketplaceTokenMapping `json:"mappedTokens,omitempty"`
	ViewId          string                     `json:"viewId,omitempty"`
}

// SavePushTemplate DTO.
type SavePushTemplate struct {
	CodeMashRequestBase
	TemplateName         string                       `json:"templateName,omitempty"`
	Description          string                       `json:"description,omitempty"`
	CommunicationChannel CommunicationChannel         `json:"communicationChannel,omitempty"`
	Tags                 []string                     `json:"tags,omitempty"`
	Translations         []*PushMessageTranslationDto `json:"translations,omitempty"`
}

// PushDeviceDto DTO.
type PushDeviceDto struct {
	DeviceId         string     `json:"deviceId,omitempty"`
	DeviceOs         string     `json:"deviceOs,omitempty"`
	Token            string     `json:"token,omitempty"`
	Brand            string     `json:"brand,omitempty"`
	Manufacturer     string     `json:"manufacturer,omitempty"`
	ModelName        string     `json:"modelName,omitempty"`
	DeviceName       string     `json:"deviceName,omitempty"`
	DeviceType       DeviceType `json:"deviceType,omitempty"`
	OsName           string     `json:"osName,omitempty"`
	OsVersion        string     `json:"osVersion,omitempty"`
	PlatformApiLevel float64    `json:"platformApiLevel,omitempty"`
}

// PushCampaignRequest DTO.
type PushCampaignRequest struct {
	Source        PushCampaignRecipientsSourceTypes `json:"source,omitempty"`
	TemplateId    string                            `json:"templateId,omitempty"`
	IntegrationId string                            `json:"integrationId,omitempty"`
	Language      string                            `json:"language,omitempty"`
	InitiatorId   string                            `json:"initiatorId,omitempty"`
	Notes         string                            `json:"notes,omitempty"`
	MappedTokens  []*TokenMappingDto                `json:"mappedTokens,omitempty"`
	CampaignTime  float64                           `json:"campaignTime,omitempty"`
}

// PushIntegration DTO.
type PushIntegration struct {
	Integration
	Provider PushProvider `json:"provider,omitempty"`
}

// PushTitle DTO.
type PushTitle struct {
	Value *TemplateCode `json:"value,omitempty"`
}

// PushBody DTO.
type PushBody struct {
	Value *TemplateCode `json:"value,omitempty"`
}

// PushMessageContent DTO.
type PushMessageContent struct {
	Title    *PushTitle `json:"title,omitempty"`
	SubTitle *PushTitle `json:"subTitle,omitempty"`
	Body     *PushBody  `json:"body,omitempty"`
}

// PushTemplate DTO.
type PushTemplate struct {
	Template[*PushMessageContent]
}

// PaymentIntegration DTO.
type PaymentIntegration struct {
	Integration
	Provider PaymentGatewayPlatform `json:"provider,omitempty"`
}

// PaymentTrigger DTO.
type PaymentTrigger struct {
	Trigger
	When         PaymentTriggerType `json:"when,omitempty"`
	Integrations []*IntegrationId   `json:"integrations,omitempty"`
	Events       []string           `json:"events,omitempty"`
}

// LoggingIntegration DTO.
type LoggingIntegration struct {
	Integration
	Provider LoggingProvider `json:"provider,omitempty"`
}

// ChatScreenContextDto DTO.
type ChatScreenContextDto struct {
	Kind   string `json:"kind,omitempty"`
	ViewId string `json:"viewId,omitempty"`
}

// LlmIntegration DTO.
type LlmIntegration struct {
	Integration
	Provider     LlmProvider `json:"provider,omitempty"`
	DefaultModel string      `json:"defaultModel,omitempty"`
}

// McpIntegration DTO.
type McpIntegration struct {
	Integration
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	Metadata  *McpMetadata `json:"metadata,omitempty"`
}

// WebhookDestinationId DTO.
type WebhookDestinationId struct {
	AggregateId
}

// TriggerEventName DTO.
type TriggerEventName struct {
	Value string `json:"value,omitempty"`
}

// WebhookDestination DTO.
type WebhookDestination struct {
	DestinationId   *WebhookDestinationId `json:"destinationId,omitempty"`
	DestinationName *DisplayName          `json:"destinationName,omitempty"`
	EndpointUrl     *DomainUrl            `json:"endpointUrl,omitempty"`
	SelectedEvents  []*TriggerEventName   `json:"selectedEvents,omitempty"`
	ExtraHeaders    map[string]string     `json:"extraHeaders,omitempty"`
	IsEnabled       bool                  `json:"isEnabled,omitempty"`
}

// WebhookIntegration DTO.
type WebhookIntegration struct {
	Integration
	Capability   string                `json:"capability,omitempty"`
	Destinations []*WebhookDestination `json:"destinations,omitempty"`
	ExtraHeaders map[string]string     `json:"extraHeaders,omitempty"`
}

// SchedulerTaskRequest DTO.
type SchedulerTaskRequest struct {
	Type SchedulerTaskType `json:"type,omitempty"`
}

// TaskId DTO.
type TaskId struct {
	AggregateId
}

// CronExpression DTO.
type CronExpression struct {
	Value string `json:"value,omitempty"`
}

// SchedulerTask DTO.
type SchedulerTask struct {
	Id          *TaskId           `json:"id,omitempty"`
	Type        SchedulerTaskType `json:"type,omitempty"`
	Name        *DisplayName      `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Cron        *CronExpression   `json:"cron,omitempty"`
	PayloadJson string            `json:"payloadJson,omitempty"`
	InitiatorId *AuthId           `json:"initiatorId,omitempty"`
	IsEnabled   bool              `json:"isEnabled,omitempty"`
	StopOnError bool              `json:"stopOnError,omitempty"`
}

// ResourceRefDto DTO.
type ResourceRefDto struct {
	ProjectId     string          `json:"projectId,omitempty"`
	IntegrationId string          `json:"integrationId,omitempty"`
	Kind          ResourceKindDto `json:"kind,omitempty"`
}

// CaseResolutionDto DTO.
type CaseResolutionDto struct {
	Problem          string                `json:"problem,omitempty"`
	Symptoms         []string              `json:"symptoms,omitempty"`
	RootCause        string                `json:"rootCause,omitempty"`
	Fix              CaseResolutionFixKind `json:"fix,omitempty"`
	FixDetail        string                `json:"fixDetail,omitempty"`
	AffectedVersions []string              `json:"affectedVersions,omitempty"`
}

// SupportCaseId DTO.
type SupportCaseId struct {
	AggregateId
	ViewId string `json:"viewId,omitempty"`
}

// SupportMessageRef DTO.
type SupportMessageRef struct {
	MessageId  string                   `json:"messageId,omitempty"`
	AuthorKind SupportMessageAuthorKind `json:"authorKind,omitempty"`
	AuthorId   string                   `json:"authorId,omitempty"`
	SentOn     *UtcDateTime             `json:"sentOn,omitempty"`
}

// CaseResolution DTO.
type CaseResolution struct {
	Problem          string                `json:"problem,omitempty"`
	Symptoms         []string              `json:"symptoms,omitempty"`
	RootCause        string                `json:"rootCause,omitempty"`
	Fix              CaseResolutionFixKind `json:"fix,omitempty"`
	FixDetail        string                `json:"fixDetail,omitempty"`
	Module           string                `json:"module,omitempty"`
	Kind             SupportCaseKind       `json:"kind,omitempty"`
	Severity         SupportCaseSeverity   `json:"severity,omitempty"`
	AffectedVersions []string              `json:"affectedVersions,omitempty"`
	ResolvedBy       string                `json:"resolvedBy,omitempty"`
}

// EchoLicenseDto DTO.
type EchoLicenseDto struct {
	Domain    string  `json:"domain,omitempty"`
	AccountId string  `json:"accountId,omitempty"`
	Email     string  `json:"email,omitempty"`
	Release   string  `json:"release,omitempty"`
	Expire    float64 `json:"expire,omitempty"`
	IsTrial   bool    `json:"isTrial,omitempty"`
	Cap       float64 `json:"cap,omitempty"`
}

// EchoRegionDto DTO.
type EchoRegionDto struct {
	Code        string `json:"code,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	HubUrl      string `json:"hubUrl,omitempty"`
}

// PublicBrandDto DTO.
type PublicBrandDto struct {
	DisplayName string `json:"displayName,omitempty"`
	MainColor   string `json:"mainColor,omitempty"`
	AccentColor string `json:"accentColor,omitempty"`
	LogoUrl     string `json:"logoUrl,omitempty"`
	IconUrl     string `json:"iconUrl,omitempty"`
}

// PublicPasswordPolicyDto DTO.
type PublicPasswordPolicyDto struct {
	MinLength      float64 `json:"minLength,omitempty"`
	MaxLength      float64 `json:"maxLength,omitempty"`
	MinNumbers     float64 `json:"minNumbers,omitempty"`
	MinUpper       float64 `json:"minUpper,omitempty"`
	MinLower       float64 `json:"minLower,omitempty"`
	MinSpecial     float64 `json:"minSpecial,omitempty"`
	AllowedSpecial string  `json:"allowedSpecial,omitempty"`
}

// PublicAuthDto DTO.
type PublicAuthDto struct {
	SocialProviders []string                 `json:"socialProviders,omitempty"`
	Passkey         bool                     `json:"passkey,omitempty"`
	Methods         []string                 `json:"methods,omitempty"`
	PasswordPolicy  *PublicPasswordPolicyDto `json:"passwordPolicy,omitempty"`
}

// AccountOwnerDto DTO.
type AccountOwnerDto struct {
	Email           string `json:"email,omitempty"`
	DisplayName     string `json:"displayName,omitempty"`
	BillingEmail    string `json:"billingEmail,omitempty"`
	OperationsEmail string `json:"operationsEmail,omitempty"`
	SecurityEmail   string `json:"securityEmail,omitempty"`
}

// AccountStatusDto DTO.
type AccountStatusDto struct {
	AccountId         string        `json:"accountId,omitempty"`
	AccountIdAsGuid   string        `json:"accountIdAsGuid,omitempty"`
	UserId            string        `json:"userId,omitempty"`
	LoggedInUserId    string        `json:"loggedInUserId,omitempty"`
	LoggedInUserEmail string        `json:"loggedInUserEmail,omitempty"`
	Status            AccountStatus `json:"status,omitempty"`
	ProjectCap        float64       `json:"projectCap,omitempty"`
	Permissions       []string      `json:"permissions,omitempty"`
	Roles             []string      `json:"roles,omitempty"`
	AllowedProjects   []string      `json:"allowedProjects,omitempty"`
	TrialWasIssued    bool          `json:"trialWasIssued,omitempty"`
}

// UsageBillingClusterChargeDto DTO.
type UsageBillingClusterChargeDto struct {
	AtlasProjectId   string  `json:"atlasProjectId,omitempty"`
	AtlasClusterName string  `json:"atlasClusterName,omitempty"`
	Cents            float64 `json:"cents,omitempty"`
}

// UsageBillingPeriodDto DTO.
type UsageBillingPeriodDto struct {
	Period        string                          `json:"period,omitempty"`
	TotalCents    float64                         `json:"totalCents,omitempty"`
	PerCluster    []*UsageBillingClusterChargeDto `json:"perCluster,omitempty"`
	RecordedAtUtc string                          `json:"recordedAtUtc,omitempty"`
}

// UsageBillingIngestionFailureDto DTO.
type UsageBillingIngestionFailureDto struct {
	Reason        string `json:"reason,omitempty"`
	Period        string `json:"period,omitempty"`
	StripeEventId string `json:"stripeEventId,omitempty"`
	Message       string `json:"message,omitempty"`
	ReportedAtUtc string `json:"reportedAtUtc,omitempty"`
}

// UsageBillingDto DTO.
type UsageBillingDto struct {
	AccountId         string                             `json:"accountId,omitempty"`
	Atlas             map[string]*UsageBillingPeriodDto  `json:"atlas,omitempty"`
	IngestionFailures []*UsageBillingIngestionFailureDto `json:"ingestionFailures,omitempty"`
}

// PromotionItemDto DTO.
type PromotionItemDto struct {
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`
}

// PromotionBlockerDto DTO.
type PromotionBlockerDto struct {
	ContentType   string `json:"contentType,omitempty"`
	ContentId     string `json:"contentId,omitempty"`
	RefKind       string `json:"refKind,omitempty"`
	UnresolvedRef string `json:"unresolvedRef,omitempty"`
}

// PromotionResultDto DTO.
type PromotionResultDto struct {
	ContentMirrored     []*PromotionItemDto    `json:"contentMirrored,omitempty"`
	ContentDeleted      []*PromotionItemDto    `json:"contentDeleted,omitempty"`
	IntegrationsSeeded  []*PromotionItemDto    `json:"integrationsSeeded,omitempty"`
	IntegrationsSkipped []*PromotionItemDto    `json:"integrationsSkipped,omitempty"`
	Blockers            []*PromotionBlockerDto `json:"blockers,omitempty"`
	FromVersion         float64                `json:"fromVersion,omitempty"`
	WasDryRun           bool                   `json:"wasDryRun,omitempty"`
}

// ProjectEnvironmentsDto DTO.
type ProjectEnvironmentsDto struct {
	Environments []string `json:"environments,omitempty"`
}

// ProjectRegionDto DTO.
type ProjectRegionDto struct {
	Id        string    `json:"id,omitempty"`
	Continent Continent `json:"continent,omitempty"`
	Name      string    `json:"name,omitempty"`
}

// ProjectBrandDto DTO.
type ProjectBrandDto struct {
	MainColor   string              `json:"mainColor,omitempty"`
	AccentColor string              `json:"accentColor,omitempty"`
	Logo        *FileResourceRefDto `json:"logo,omitempty"`
	Icon        *FileResourceRefDto `json:"icon,omitempty"`
}

// NotificationsSettingsGroupDto DTO.
type NotificationsSettingsGroupDto struct {
	Tag  string   `json:"tag,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NotificationSettingsChannelDto DTO.
type NotificationSettingsChannelDto struct {
	Channel CommunicationChannel             `json:"channel,omitempty"`
	Groups  []*NotificationsSettingsGroupDto `json:"groups,omitempty"`
}

// NotificationSettingsDto DTO.
type NotificationSettingsDto struct {
	Channels  []*NotificationSettingsChannelDto `json:"channels,omitempty"`
	AllGroups []*GroupDefinitionDto             `json:"allGroups,omitempty"`
	AllTags   []*TagDefinitionDto               `json:"allTags,omitempty"`
}

// AuthenticationFlowPasswordPolicyDto DTO.
type AuthenticationFlowPasswordPolicyDto struct {
	MinLength      float64 `json:"minLength,omitempty"`
	MaxLength      float64 `json:"maxLength,omitempty"`
	MinNumbers     float64 `json:"minNumbers,omitempty"`
	MinUpper       float64 `json:"minUpper,omitempty"`
	MinLower       float64 `json:"minLower,omitempty"`
	MinSpecial     float64 `json:"minSpecial,omitempty"`
	AllowedSpecial string  `json:"allowedSpecial,omitempty"`
}

// AuthenticationFlowSummaryDto DTO.
type AuthenticationFlowSummaryDto struct {
	Type               string                               `json:"type,omitempty"`
	Provider           string                               `json:"provider,omitempty"`
	PasswordComplexity *AuthenticationFlowPasswordPolicyDto `json:"passwordComplexity,omitempty"`
}

// TriggerDto DTO.
type TriggerDto struct {
	Type           TriggerType       `json:"type,omitempty"`
	ViewId         string            `json:"viewId,omitempty"`
	Name           string            `json:"name,omitempty"`
	ThenAction     *TriggerActionDto `json:"thenAction,omitempty"`
	Description    string            `json:"description,omitempty"`
	IsEnabled      bool              `json:"isEnabled,omitempty"`
	ActivationCode string            `json:"activationCode,omitempty"`
	SavedByAuthId  string            `json:"savedByAuthId,omitempty"`
}

// SchemaTriggerDto DTO.
type SchemaTriggerDto struct {
	TriggerDto
	SchemaId          string            `json:"schemaId,omitempty"`
	When              SchemaTriggerType `json:"when,omitempty"`
	ConfigurationCode string            `json:"configurationCode,omitempty"`
}

// DatabaseDto DTO.
type DatabaseDto struct {
	IsEnabled                 bool                `json:"isEnabled,omitempty"`
	Triggers                  []*SchemaTriggerDto `json:"triggers,omitempty"`
	DefaultIntegrationViewIds map[string]string   `json:"defaultIntegrationViewIds,omitempty"`
}

// EmailDto DTO.
type EmailDto struct {
	IsEnabled                 bool              `json:"isEnabled,omitempty"`
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds,omitempty"`
}

// AiDto DTO.
type AiDto struct {
	IsEnabled                bool   `json:"isEnabled,omitempty"`
	DefaultIntegrationViewId string `json:"defaultIntegrationViewId,omitempty"`
}

// MembershipTriggerDto DTO.
type MembershipTriggerDto struct {
	TriggerDto
	When MembershipTriggerType `json:"when,omitempty"`
}

// RoleItemDto DTO.
type RoleItemDto struct {
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	DisplayName      string   `json:"displayName,omitempty"`
	Description      string   `json:"description,omitempty"`
	IsSystem         bool     `json:"isSystem,omitempty"`
	AttachedPolicies []string `json:"attachedPolicies,omitempty"`
}

// PermissionDto DTO.
type PermissionDto struct {
	Sid       string           `json:"sid,omitempty"`
	Effect    PermissionEffect `json:"effect,omitempty"`
	Actions   []string         `json:"actions,omitempty"`
	Resources []string         `json:"resources,omitempty"`
}

// PolicyItemDto DTO.
type PolicyItemDto struct {
	Id          string           `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	IsSystem    bool             `json:"isSystem,omitempty"`
	Permissions []*PermissionDto `json:"permissions,omitempty"`
}

// AuthorizationDto DTO.
type AuthorizationDto struct {
	UserRegistersAsRole          string   `json:"userRegistersAsRole,omitempty"`
	AllowedRegisterRoles         []string `json:"allowedRegisterRoles,omitempty"`
	AllowedProviderRegisterRoles []string `json:"allowedProviderRegisterRoles,omitempty"`
}

// MembershipDto DTO.
type MembershipDto struct {
	IsEnabled              bool                    `json:"isEnabled,omitempty"`
	Triggers               []*MembershipTriggerDto `json:"triggers,omitempty"`
	CustomRoles            []*RoleItemDto          `json:"customRoles,omitempty"`
	CustomPolicies         []*PolicyItemDto        `json:"customPolicies,omitempty"`
	Authorization          *AuthorizationDto       `json:"authorization,omitempty"`
	RequireEmailValidation bool                    `json:"requireEmailValidation,omitempty"`
}

// LoggingDto DTO.
type LoggingDto struct {
	IsEnabled     bool `json:"isEnabled,omitempty"`
	IsEstablished bool `json:"isEstablished,omitempty"`
}

// ServerEventsDto DTO.
type ServerEventsDto struct {
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// PushDto DTO.
type PushDto struct {
	IsEnabled                 bool                `json:"isEnabled,omitempty"`
	DefaultIntegrationViewIds map[string]string   `json:"defaultIntegrationViewIds,omitempty"`
	MarketingTags             []*TagDefinitionDto `json:"marketingTags,omitempty"`
	TransactionalTags         []*TagDefinitionDto `json:"transactionalTags,omitempty"`
}

// SchedulerDto DTO.
type SchedulerDto struct {
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// CodeDto DTO.
type CodeDto struct {
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// FilesTriggerDto DTO.
type FilesTriggerDto struct {
	TriggerDto
	When   FilesTriggerType `json:"when,omitempty"`
	Folder string           `json:"folder,omitempty"`
}

// FilesDto DTO.
type FilesDto struct {
	IsEnabled                 bool               `json:"isEnabled,omitempty"`
	Triggers                  []*FilesTriggerDto `json:"triggers,omitempty"`
	DefaultIntegrationViewIds map[string]string  `json:"defaultIntegrationViewIds,omitempty"`
}

// PaymentTriggerDto DTO.
type PaymentTriggerDto struct {
	TriggerDto
	When         PaymentTriggerType `json:"when,omitempty"`
	Integrations []string           `json:"integrations,omitempty"`
	Events       []string           `json:"events,omitempty"`
}

// PaymentsDto DTO.
type PaymentsDto struct {
	IsEnabled bool                 `json:"isEnabled,omitempty"`
	Triggers  []*PaymentTriggerDto `json:"triggers,omitempty"`
}

// SmsDto DTO.
type SmsDto struct {
	IsEnabled                 bool              `json:"isEnabled,omitempty"`
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds,omitempty"`
}

// ProjectDto DTO.
type ProjectDto struct {
	AccountViewId                 string                          `json:"accountViewId,omitempty"`
	ProjectStatus                 ProjectStatus                   `json:"projectStatus,omitempty"`
	IsActive                      bool                            `json:"isActive,omitempty"`
	ViewId                        string                          `json:"viewId,omitempty"`
	Name                          string                          `json:"name,omitempty"`
	UniqueName                    string                          `json:"uniqueName,omitempty"`
	HostLabel                     string                          `json:"hostLabel,omitempty"`
	ApiHost                       string                          `json:"apiHost,omitempty"`
	Description                   string                          `json:"description,omitempty"`
	MarketingUrl                  string                          `json:"marketingUrl,omitempty"`
	CanonicalAdminUrl             string                          `json:"canonicalAdminUrl,omitempty"`
	AdminUrl                      string                          `json:"adminUrl,omitempty"`
	EffectiveAdminUrl             string                          `json:"effectiveAdminUrl,omitempty"`
	DefaultLanguage               string                          `json:"defaultLanguage,omitempty"`
	Languages                     []string                        `json:"languages,omitempty"`
	PrimaryRegion                 *ProjectRegionDto               `json:"primaryRegion,omitempty"`
	AdditionalRegions             []*ProjectRegionDto             `json:"additionalRegions,omitempty"`
	IsMultiRegionEligible         bool                            `json:"isMultiRegionEligible,omitempty"`
	Brand                         *ProjectBrandDto                `json:"brand,omitempty"`
	NotificationSettings          *NotificationSettingsDto        `json:"notificationSettings,omitempty"`
	AllowedOrigins                []string                        `json:"allowedOrigins,omitempty"`
	ExposeBrandToAdminPortal      bool                            `json:"exposeBrandToAdminPortal,omitempty"`
	ExposeAuthToAdminPortal       bool                            `json:"exposeAuthToAdminPortal,omitempty"`
	AdminPortalEnabled            bool                            `json:"adminPortalEnabled,omitempty"`
	AdminPortalServiceUserId      string                          `json:"adminPortalServiceUserId,omitempty"`
	MembershipAuthenticationFlows []*AuthenticationFlowSummaryDto `json:"membershipAuthenticationFlows,omitempty"`
	ExposeLegalToAdminPortal      bool                            `json:"exposeLegalToAdminPortal,omitempty"`
	LegalTermsMarkdown            string                          `json:"legalTermsMarkdown,omitempty"`
	LegalPrivacyMarkdown          string                          `json:"legalPrivacyMarkdown,omitempty"`
	Environments                  []string                        `json:"environments,omitempty"`
	EnvironmentRanks              map[string]float64              `json:"environmentRanks,omitempty"`
	Database                      *DatabaseDto                    `json:"database,omitempty"`
	Email                         *EmailDto                       `json:"email,omitempty"`
	Ai                            *AiDto                          `json:"ai,omitempty"`
	Membership                    *MembershipDto                  `json:"membership,omitempty"`
	Logging                       *LoggingDto                     `json:"logging,omitempty"`
	ServerEvents                  *ServerEventsDto                `json:"serverEvents,omitempty"`
	Push                          *PushDto                        `json:"push,omitempty"`
	Scheduler                     *SchedulerDto                   `json:"scheduler,omitempty"`
	Code                          *CodeDto                        `json:"code,omitempty"`
	Files                         *FilesDto                       `json:"files,omitempty"`
	Payments                      *PaymentsDto                    `json:"payments,omitempty"`
	Sms                           *SmsDto                         `json:"sms,omitempty"`
	DatabaseEnabled               bool                            `json:"databaseEnabled,omitempty"`
	EmailEnabled                  bool                            `json:"emailEnabled,omitempty"`
	MembershipEnabled             bool                            `json:"membershipEnabled,omitempty"`
	LoggingEnabled                bool                            `json:"loggingEnabled,omitempty"`
	ServerEventsEnabled           bool                            `json:"serverEventsEnabled,omitempty"`
	PushEnabled                   bool                            `json:"pushEnabled,omitempty"`
	SchedulerEnabled              bool                            `json:"schedulerEnabled,omitempty"`
	CodeEnabled                   bool                            `json:"codeEnabled,omitempty"`
	FilesEnabled                  bool                            `json:"filesEnabled,omitempty"`
	PaymentsEnabled               bool                            `json:"paymentsEnabled,omitempty"`
	SmsEnabled                    bool                            `json:"smsEnabled,omitempty"`
	DefaultLlmIntegrationViewId   string                          `json:"defaultLlmIntegrationViewId,omitempty"`
	Connections                   float64                         `json:"connections,omitempty"`
}

// ProjectListItemDto DTO.
type ProjectListItemDto struct {
	ViewId            string              `json:"viewId,omitempty"`
	IsActive          bool                `json:"isActive,omitempty"`
	ProjectStatus     ProjectStatus       `json:"projectStatus,omitempty"`
	Name              string              `json:"name,omitempty"`
	UniqueName        string              `json:"uniqueName,omitempty"`
	PrimaryRegion     *ProjectRegionDto   `json:"primaryRegion,omitempty"`
	AdditionalRegions []*ProjectRegionDto `json:"additionalRegions,omitempty"`
}

// PaginatedResponse DTO.
type PaginatedResponse[TViewModelProjection any] struct {
	Items         []TViewModelProjection `json:"items,omitempty"`
	HasMore       bool                   `json:"hasMore,omitempty"`
	HasPrevious   bool                   `json:"hasPrevious,omitempty"`
	StartingAfter string                 `json:"startingAfter,omitempty"`
	EndingBefore  string                 `json:"endingBefore,omitempty"`
}

// AccessInformationDto DTO.
type AccessInformationDto struct {
	Ip       string `json:"ip,omitempty"`
	Date     string `json:"date,omitempty"`
	TimeZone string `json:"timeZone,omitempty"`
}

// RegistrationDto DTO.
type RegistrationDto struct {
	RegistrationInformation *AccessInformationDto `json:"registrationInformation,omitempty"`
}

// LoginDto DTO.
type LoginDto struct {
	NeedChangePasswordOnNextLogin bool                  `json:"needChangePasswordOnNextLogin,omitempty"`
	LastAccessInformation         *AccessInformationDto `json:"lastAccessInformation,omitempty"`
}

// UserGeneralInfoDto DTO.
type UserGeneralInfoDto struct {
	Phone                     string                 `json:"phone,omitempty"`
	PrimaryEmail              string                 `json:"primaryEmail,omitempty"`
	DisplayName               string                 `json:"displayName,omitempty"`
	FirstName                 string                 `json:"firstName,omitempty"`
	LastName                  string                 `json:"lastName,omitempty"`
	FullName                  string                 `json:"fullName,omitempty"`
	AddressLine1              string                 `json:"addressLine1,omitempty"`
	AddressLine2              string                 `json:"addressLine2,omitempty"`
	Country                   string                 `json:"country,omitempty"`
	City                      string                 `json:"city,omitempty"`
	State                     string                 `json:"state,omitempty"`
	PostalCode                string                 `json:"postalCode,omitempty"`
	Company                   string                 `json:"company,omitempty"`
	Gender                    Gender                 `json:"gender,omitempty"`
	BirthDate                 float64                `json:"birthDate,omitempty"`
	TimeZone                  string                 `json:"timeZone,omitempty"`
	Language                  string                 `json:"language,omitempty"`
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
	ExtraMetadata             string                 `json:"extraMetadata,omitempty"`
	Notes                     string                 `json:"notes,omitempty"`
}

// AuthDto DTO.
type AuthDto struct {
	Id           string              `json:"id,omitempty"`
	Type         AuthType            `json:"type,omitempty"`
	Email        string              `json:"email,omitempty"`
	UserName     string              `json:"userName,omitempty"`
	Registration *RegistrationDto    `json:"registration,omitempty"`
	Login        *LoginDto           `json:"login,omitempty"`
	GeneralInfo  *UserGeneralInfoDto `json:"generalInfo,omitempty"`
	Roles        []string            `json:"roles,omitempty"`
	PushDevices  []string            `json:"pushDevices,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
	Status       AuthStatus          `json:"status,omitempty"`
	CreatedOn    string              `json:"createdOn,omitempty"`
	ModifiedOn   string              `json:"modifiedOn,omitempty"`
}

// AccountPasswordPolicyDto DTO.
type AccountPasswordPolicyDto struct {
	MinLength      float64 `json:"minLength,omitempty"`
	MaxLength      float64 `json:"maxLength,omitempty"`
	MinNumbers     float64 `json:"minNumbers,omitempty"`
	MaxNumbers     float64 `json:"maxNumbers,omitempty"`
	MinUpper       float64 `json:"minUpper,omitempty"`
	MaxUpper       float64 `json:"maxUpper,omitempty"`
	MinLower       float64 `json:"minLower,omitempty"`
	MaxLower       float64 `json:"maxLower,omitempty"`
	MinSpecial     float64 `json:"minSpecial,omitempty"`
	MaxSpecial     float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial string  `json:"allowedSpecial,omitempty"`
}

// AccountTeamRoleDto DTO.
type AccountTeamRoleDto struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	IsSystem    bool     `json:"isSystem,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

// AccountPasskeyListItemDto DTO.
type AccountPasskeyListItemDto struct {
	CredentialId    string `json:"credentialId,omitempty"`
	FriendlyName    string `json:"friendlyName,omitempty"`
	RegisteredOnUtc string `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   string `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool   `json:"isRevoked,omitempty"`
}

// LicenseDomainDnsRecordDto DTO.
type LicenseDomainDnsRecordDto struct {
	Host       string `json:"host,omitempty"`
	RecordType string `json:"recordType,omitempty"`
	Resolved   bool   `json:"resolved,omitempty"`
	Required   bool   `json:"required,omitempty"`
}

// LicenseDomainDnsStatusDto DTO.
type LicenseDomainDnsStatusDto struct {
	Domain              string                       `json:"domain,omitempty"`
	AllRequiredResolved bool                         `json:"allRequiredResolved,omitempty"`
	Records             []*LicenseDomainDnsRecordDto `json:"records,omitempty"`
}

// LicenseDomainVerificationChallengeDto DTO.
type LicenseDomainVerificationChallengeDto struct {
	Domain        string `json:"domain,omitempty"`
	TxtHost       string `json:"txtHost,omitempty"`
	TxtValue      string `json:"txtValue,omitempty"`
	ExpiresAtUtc  string `json:"expiresAtUtc,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	VerifiedAtUtc string `json:"verifiedAtUtc,omitempty"`
	Skipped       bool   `json:"skipped,omitempty"`
}

// LicenseDomainVerificationStatusDto DTO.
type LicenseDomainVerificationStatusDto struct {
	Domain           string `json:"domain,omitempty"`
	Verified         bool   `json:"verified,omitempty"`
	Skipped          bool   `json:"skipped,omitempty"`
	TxtHost          string `json:"txtHost,omitempty"`
	ExpectedTxtValue string `json:"expectedTxtValue,omitempty"`
	ObservedTxtValue string `json:"observedTxtValue,omitempty"`
	ExpiresAtUtc     string `json:"expiresAtUtc,omitempty"`
	VerifiedAtUtc    string `json:"verifiedAtUtc,omitempty"`
	Message          string `json:"message,omitempty"`
}

// CodeMashSubscriptionDto DTO.
type CodeMashSubscriptionDto struct {
	ViewId            string `json:"viewId,omitempty"`
	Domain            string `json:"domain,omitempty"`
	WillExpireOn      string `json:"willExpireOn,omitempty"`
	IssuedOn          string `json:"issuedOn,omitempty"`
	IsTrial           bool   `json:"isTrial,omitempty"`
	SubscriptionRefId string `json:"subscriptionRefId,omitempty"`
}

// LicenseDto DTO.
type LicenseDto struct {
	CodeMashSubscriptionDto
	IsEnterprise bool    `json:"isEnterprise,omitempty"`
	ProjectCap   float64 `json:"projectCap,omitempty"`
}

// LicenseHeartbeatVerdictDto DTO.
type LicenseHeartbeatVerdictDto struct {
	Status           string `json:"status,omitempty"`
	ProofToken       string `json:"proofToken,omitempty"`
	ServerTimeUtc    string `json:"serverTimeUtc,omitempty"`
	GraceUntilUtc    string `json:"graceUntilUtc,omitempty"`
	InstallationId   string `json:"installationId,omitempty"`
	LicenseAccountId string `json:"licenseAccountId,omitempty"`
	Domain           string `json:"domain,omitempty"`
	Signature        string `json:"signature,omitempty"`
	Message          string `json:"message,omitempty"`
}

// InstallationLicenseStatusDto DTO.
type InstallationLicenseStatusDto struct {
	StoredMode         string  `json:"storedMode,omitempty"`
	EffectiveMode      string  `json:"effectiveMode,omitempty"`
	IsProduction       bool    `json:"isProduction,omitempty"`
	GraceDaysLeft      float64 `json:"graceDaysLeft,omitempty"`
	GraceUntilUtc      string  `json:"graceUntilUtc,omitempty"`
	LastProvenAtUtc    string  `json:"lastProvenAtUtc,omitempty"`
	LastHeartbeatAtUtc string  `json:"lastHeartbeatAtUtc,omitempty"`
	InstallationDomain string  `json:"installationDomain,omitempty"`
	LicensedDomain     string  `json:"licensedDomain,omitempty"`
	HostKind           string  `json:"hostKind,omitempty"`
	IsTrialLicense     bool    `json:"isTrialLicense,omitempty"`
	LicenseExpireUtc   string  `json:"licenseExpireUtc,omitempty"`
	Message            string  `json:"message,omitempty"`
}

// ServiceUserApiKeyDto DTO.
type ServiceUserApiKeyDto struct {
	Id            float64  `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	VisibleKey    string   `json:"visibleKey,omitempty"`
	Scopes        []string `json:"scopes,omitempty"`
	CreatedDate   string   `json:"createdDate,omitempty"`
	ExpiryDate    string   `json:"expiryDate,omitempty"`
	CancelledDate string   `json:"cancelledDate,omitempty"`
	Active        bool     `json:"active,omitempty"`
}

// GetTriggerResponse DTO.
type GetTriggerResponse struct {
	ResponseBase
}

// TriggerProjectionList DTO.
type TriggerProjectionList struct {
	ViewId            string            `json:"viewId,omitempty"`
	Name              string            `json:"name,omitempty"`
	ActionType        TriggerActionType `json:"actionType,omitempty"`
	HasPreExecuteCode bool              `json:"hasPreExecuteCode,omitempty"`
	IsEnabled         bool              `json:"isEnabled,omitempty"`
}

// MembershipTriggerProjectionList DTO.
type MembershipTriggerProjectionList struct {
	TriggerProjectionList
	Type           MembershipTriggerType `json:"type,omitempty"`
	DestinationIds []string              `json:"destinationIds,omitempty"`
}

// RoleListProjectionDto DTO.
type RoleListProjectionDto struct {
	ViewId      string  `json:"viewId,omitempty"`
	Name        string  `json:"name,omitempty"`
	DisplayName string  `json:"displayName,omitempty"`
	IsSystem    bool    `json:"isSystem,omitempty"`
	PolicyCount float64 `json:"policyCount,omitempty"`
}

// PasskeySettingsDto DTO.
type PasskeySettingsDto struct {
	Enabled                       bool    `json:"enabled,omitempty"`
	CodeTtlMinutes                float64 `json:"codeTtlMinutes,omitempty"`
	MaxCredentialsPerUser         float64 `json:"maxCredentialsPerUser,omitempty"`
	RecoveryCodeCount             float64 `json:"recoveryCodeCount,omitempty"`
	GenerateRecoveryCodesAtSignup bool    `json:"generateRecoveryCodesAtSignup,omitempty"`
	AuthenticatorAttachment       string  `json:"authenticatorAttachment,omitempty"`
	AllowMagicLinkRecovery        bool    `json:"allowMagicLinkRecovery,omitempty"`
	RefreshTokenTtlDays           float64 `json:"refreshTokenTtlDays,omitempty"`
	RpId                          string  `json:"rpId,omitempty"`
}

// IntegrationListProjection DTO.
type IntegrationListProjection struct {
	ViewId                            string   `json:"viewId,omitempty"`
	IntegrationName                   string   `json:"integrationName,omitempty"`
	IsEnabled                         bool     `json:"isEnabled,omitempty"`
	LastIntegrationTestAtUtc          string   `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      bool     `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         []string `json:"lastIntegrationTestErrors,omitempty"`
	HumanDeliveryConfirmedAtUtc       string   `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool     `json:"requiresHumanDeliveryConfirmation,omitempty"`
}

// MembershipIntegrationListProjection DTO.
type MembershipIntegrationListProjection struct {
	IntegrationListProjection
	Provider MembershipProvider `json:"provider,omitempty"`
}

// MembershipMessageTemplateDto DTO.
type MembershipMessageTemplateDto struct {
	Id string `json:"id,omitempty"`
}

// MembershipEmailActionSettingsDto DTO.
type MembershipEmailActionSettingsDto struct {
	SendEmail bool                          `json:"sendEmail,omitempty"`
	Template  *MembershipMessageTemplateDto `json:"template,omitempty"`
	Callback  string                        `json:"callback,omitempty"`
}

// MembershipEmailPreferencesDto DTO.
type MembershipEmailPreferencesDto struct {
	RegistrationViaEmail  *MembershipEmailActionSettingsDto `json:"registrationViaEmail,omitempty"`
	VerificationViaEmail  *MembershipEmailActionSettingsDto `json:"verificationViaEmail,omitempty"`
	PasswordResetViaEmail *MembershipEmailActionSettingsDto `json:"passwordResetViaEmail,omitempty"`
	InvitationViaEmail    *MembershipEmailActionSettingsDto `json:"invitationViaEmail,omitempty"`
	DeactivationViaEmail  *MembershipEmailActionSettingsDto `json:"deactivationViaEmail,omitempty"`
}

// PasswordComplexityDto DTO.
type PasswordComplexityDto struct {
	MinLength      float64 `json:"minLength,omitempty"`
	MaxLength      float64 `json:"maxLength,omitempty"`
	MinNumbers     float64 `json:"minNumbers,omitempty"`
	MaxNumbers     float64 `json:"maxNumbers,omitempty"`
	MinUpper       float64 `json:"minUpper,omitempty"`
	MaxUpper       float64 `json:"maxUpper,omitempty"`
	MinLower       float64 `json:"minLower,omitempty"`
	MaxLower       float64 `json:"maxLower,omitempty"`
	MinSpecial     float64 `json:"minSpecial,omitempty"`
	MaxSpecial     float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial string  `json:"allowedSpecial,omitempty"`
}

// MembershipAuthorizationViewDto DTO.
type MembershipAuthorizationViewDto struct {
	EmailPreferences             *MembershipEmailPreferencesDto `json:"emailPreferences,omitempty"`
	UserRegistersAsRole          string                         `json:"userRegistersAsRole,omitempty"`
	GuestRegistersAsRole         string                         `json:"guestRegistersAsRole,omitempty"`
	AllowedRegisterRoles         []string                       `json:"allowedRegisterRoles,omitempty"`
	AllowedProviderRegisterRoles []string                       `json:"allowedProviderRegisterRoles,omitempty"`
	ResetPasswordTokenExpiration float64                        `json:"resetPasswordTokenExpiration,omitempty"`
	InvitationExpiration         float64                        `json:"invitationExpiration,omitempty"`
	EmailVerificationExpiration  float64                        `json:"emailVerificationExpiration,omitempty"`
	DeactivationExpiration       float64                        `json:"deactivationExpiration,omitempty"`
	DefaultSubscribeToNews       bool                           `json:"defaultSubscribeToNews,omitempty"`
	PasswordComplexity           *PasswordComplexityDto         `json:"passwordComplexity,omitempty"`
}

// MembershipCredentialsSettingsModeDto DTO.
type MembershipCredentialsSettingsModeDto struct {
	Name      string `json:"name,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
}

// MembershipCredentialsSettingsDto DTO.
type MembershipCredentialsSettingsDto struct {
	LogoutUrl      string                                  `json:"logoutUrl,omitempty"`
	AllowUsernames bool                                    `json:"allowUsernames,omitempty"`
	Modes          []*MembershipCredentialsSettingsModeDto `json:"modes,omitempty"`
}

// MembershipAuthenticationViewDto DTO.
type MembershipAuthenticationViewDto struct {
	CredentialsSettings *MembershipCredentialsSettingsDto `json:"credentialsSettings,omitempty"`
	Flows               []string                          `json:"flows,omitempty"`
}

// SchemaTriggerProjectionList DTO.
type SchemaTriggerProjectionList struct {
	TriggerProjectionList
	Type SchemaTriggerType `json:"type,omitempty"`
}

// JsonSchemaFieldDto DTO.
type JsonSchemaFieldDto struct {
	FieldName string `json:"fieldName,omitempty"`
}

// DataSchemaDto DTO.
type DataSchemaDto struct {
	Json   string                `json:"json,omitempty"`
	Fields []*JsonSchemaFieldDto `json:"fields,omitempty"`
}

// VisualSchemaDto DTO.
type VisualSchemaDto struct {
	Json string `json:"json,omitempty"`
}

// TaxonomyDto DTO.
type TaxonomyDto struct {
	ViewId                string           `json:"viewId,omitempty"`
	Name                  string           `json:"name,omitempty"`
	Slug                  string           `json:"slug,omitempty"`
	ParentId              string           `json:"parentId,omitempty"`
	Description           string           `json:"description,omitempty"`
	TermsMetaDataSchema   *DataSchemaDto   `json:"termsMetaDataSchema,omitempty"`
	TermsMetaVisualSchema *VisualSchemaDto `json:"termsMetaVisualSchema,omitempty"`
	Dependencies          []string         `json:"dependencies,omitempty"`
}

// TaxonomyListProjection DTO.
type TaxonomyListProjection struct {
	ViewId       string `json:"viewId,omitempty"`
	TaxonomyName string `json:"taxonomyName,omitempty"`
	TaxonomySlug string `json:"taxonomySlug,omitempty"`
	ParentId     string `json:"parentId,omitempty"`
}

// TermMultiParentDto DTO.
type TermMultiParentDto struct {
	TaxonomyId string            `json:"taxonomyId,omitempty"`
	ParentId   string            `json:"parentId,omitempty"`
	Name       string            `json:"name,omitempty"`
	Names      map[string]string `json:"names,omitempty"`
}

// TermTreeDto DTO.
type TermTreeDto struct {
	Id           string                `json:"id,omitempty"`
	TaxonomyId   string                `json:"taxonomyId,omitempty"`
	TaxonomyName string                `json:"taxonomyName,omitempty"`
	ParentId     string                `json:"parentId,omitempty"`
	Order        float64               `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         map[string]any        `json:"meta,omitempty"`
	Children     []*TermTreeDto        `json:"children,omitempty"`
}

// TaxonomyTreeDto DTO.
type TaxonomyTreeDto struct {
	ViewId       string             `json:"viewId,omitempty"`
	TaxonomyName string             `json:"taxonomyName,omitempty"`
	TaxonomySlug string             `json:"taxonomySlug,omitempty"`
	ParentId     string             `json:"parentId,omitempty"`
	Children     []*TaxonomyTreeDto `json:"children,omitempty"`
	Terms        []*TermTreeDto     `json:"terms,omitempty"`
}

// TermDto DTO.
type TermDto struct {
	Id           string                `json:"id,omitempty"`
	TaxonomyId   string                `json:"taxonomyId,omitempty"`
	TaxonomyName string                `json:"taxonomyName,omitempty"`
	ParentId     string                `json:"parentId,omitempty"`
	Order        float64               `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         map[string]any        `json:"meta,omitempty"`
}

// AppliedTaxonomyDto DTO.
type AppliedTaxonomyDto struct {
	CatalogId    string  `json:"catalogId,omitempty"`
	Id           string  `json:"id,omitempty"`
	Slug         string  `json:"slug,omitempty"`
	Title        string  `json:"title,omitempty"`
	Action       string  `json:"action,omitempty"`
	TermsCreated float64 `json:"termsCreated,omitempty"`
}

// AppliedCollectionDto DTO.
type AppliedCollectionDto struct {
	Entity       string   `json:"entity,omitempty"`
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Title        string   `json:"title,omitempty"`
	Action       string   `json:"action,omitempty"`
	Published    bool     `json:"published,omitempty"`
	LinkedFields []string `json:"linkedFields,omitempty"`
}

// SchemaDto DTO.
type SchemaDto struct {
	ViewId            string             `json:"viewId,omitempty"`
	SchemaName        string             `json:"schemaName,omitempty"`
	SchemaSlug        string             `json:"schemaSlug,omitempty"`
	Version           float64            `json:"version,omitempty"`
	MetaSchemaVersion float64            `json:"metaSchemaVersion,omitempty"`
	DataSchema        *DataSchemaDto     `json:"dataSchema,omitempty"`
	VisualSchema      *VisualSchemaDto   `json:"visualSchema,omitempty"`
	PublishedAt       string             `json:"publishedAt,omitempty"`
	Settings          *SchemaSettingsDto `json:"settings,omitempty"`
	Triggers          []*TriggerDto      `json:"triggers,omitempty"`
}

// SchemaListProjection DTO.
type SchemaListProjection struct {
	ViewId            string  `json:"viewId,omitempty"`
	SchemaName        string  `json:"schemaName,omitempty"`
	SchemaTitle       string  `json:"schemaTitle,omitempty"`
	LatestVersion     float64 `json:"latestVersion,omitempty"`
	HasDraft          bool    `json:"hasDraft,omitempty"`
	MetaSchemaVersion float64 `json:"metaSchemaVersion,omitempty"`
	Description       string  `json:"description,omitempty"`
}

// SchemaDraftDto DTO.
type SchemaDraftDto struct {
	DataSchema   *DataSchemaDto   `json:"dataSchema,omitempty"`
	VisualSchema *VisualSchemaDto `json:"visualSchema,omitempty"`
	UpdatedAt    string           `json:"updatedAt,omitempty"`
}

// SchemaDiffDto DTO.
type SchemaDiffDto struct {
	FromVersion        float64  `json:"fromVersion,omitempty"`
	ToVersion          float64  `json:"toVersion,omitempty"`
	Added              []string `json:"added,omitempty"`
	Removed            []string `json:"removed,omitempty"`
	TypeChanged        []string `json:"typeChanged,omitempty"`
	ValidatorTightened []string `json:"validatorTightened,omitempty"`
}

// SchemaVersionSummaryDto DTO.
type SchemaVersionSummaryDto struct {
	Version           float64 `json:"version,omitempty"`
	MetaSchemaVersion float64 `json:"metaSchemaVersion,omitempty"`
	PublishedAt       string  `json:"publishedAt,omitempty"`
}

// CollectionIndexKeyDto DTO.
type CollectionIndexKeyDto struct {
	Field string  `json:"field,omitempty"`
	Order float64 `json:"order,omitempty"`
}

// CollectionIndexDto DTO.
type CollectionIndexDto struct {
	Name string                   `json:"name,omitempty"`
	Keys []*CollectionIndexKeyDto `json:"keys,omitempty"`
}

// SeedCollectionReportItemDto DTO.
type SeedCollectionReportItemDto struct {
	CollectionName string   `json:"collectionName,omitempty"`
	Requested      float64  `json:"requested,omitempty"`
	Inserted       float64  `json:"inserted,omitempty"`
	Ids            []string `json:"ids,omitempty"`
	Errors         []string `json:"errors,omitempty"`
}

// SeedCollectionRecordsResultDto DTO.
type SeedCollectionRecordsResultDto struct {
	InsertOrder []string                       `json:"insertOrder,omitempty"`
	Report      []*SeedCollectionReportItemDto `json:"report,omitempty"`
}

// DatabaseIntegrationListProjection DTO.
type DatabaseIntegrationListProjection struct {
	IntegrationListProjection
	Provider DatabaseProvider `json:"provider,omitempty"`
}

// FlexTierDto DTO.
type FlexTierDto struct {
	Code        string  `json:"code,omitempty"`
	Step        float64 `json:"step,omitempty"`
	DisplayName string  `json:"displayName,omitempty"`
}

// IntegrationTestResultItemDto DTO.
type IntegrationTestResultItemDto struct {
	Operation string   `json:"operation,omitempty"`
	Result    string   `json:"result,omitempty"`
	Errors    []string `json:"errors,omitempty"`
}

// SchemaRefDto DTO.
type SchemaRefDto struct {
	SchemaId              string `json:"schemaId,omitempty"`
	SchemaName            string `json:"schemaName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// CollectionImportDto DTO.
type CollectionImportDto struct {
	Id            string                    `json:"id,omitempty"`
	Schema        *SchemaRefDto             `json:"schema,omitempty"`
	File          *FileResourceRefDto       `json:"file,omitempty"`
	ErrorFile     *FileResourceRefDto       `json:"errorFile,omitempty"`
	Delimiter     string                    `json:"delimiter,omitempty"`
	HasHeader     bool                      `json:"hasHeader,omitempty"`
	Status        string                    `json:"status,omitempty"`
	TotalRows     float64                   `json:"totalRows,omitempty"`
	TotalImported float64                   `json:"totalImported,omitempty"`
	TotalErrors   float64                   `json:"totalErrors,omitempty"`
	FailureReason string                    `json:"failureReason,omitempty"`
	Mapping       []*ImportColumnMappingDto `json:"mapping,omitempty"`
	CreatedOn     string                    `json:"createdOn,omitempty"`
	StartedOn     string                    `json:"startedOn,omitempty"`
	CompletedOn   string                    `json:"completedOn,omitempty"`
}

// ImportUploadTargetDto DTO.
type ImportUploadTargetDto struct {
	Url         string              `json:"url,omitempty"`
	ContentType string              `json:"contentType,omitempty"`
	File        *FileResourceRefDto `json:"file,omitempty"`
}

// ImportFileColumnDto DTO.
type ImportFileColumnDto struct {
	Index        float64  `json:"index,omitempty"`
	Header       string   `json:"header,omitempty"`
	Samples      []string `json:"samples,omitempty"`
	DetectedType string   `json:"detectedType,omitempty"`
}

// ImportFileAnalysisDto DTO.
type ImportFileAnalysisDto struct {
	File           *FileResourceRefDto    `json:"file,omitempty"`
	Columns        []*ImportFileColumnDto `json:"columns,omitempty"`
	SampleRowCount float64                `json:"sampleRowCount,omitempty"`
}

// MongoDbAggregateListProjection DTO.
type MongoDbAggregateListProjection struct {
	ViewId       string `json:"viewId,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	SchemaViewId string `json:"schemaViewId,omitempty"`
}

// FilesTriggerProjectionList DTO.
type FilesTriggerProjectionList struct {
	TriggerProjectionList
	Type   FilesTriggerType `json:"type,omitempty"`
	Folder string           `json:"folder,omitempty"`
}

// FilesIntegrationListProjection DTO.
type FilesIntegrationListProjection struct {
	IntegrationListProjection
	Provider FileProvider `json:"provider,omitempty"`
}

// PublicFolderDto DTO.
type PublicFolderDto struct {
	Path      string `json:"path,omitempty"`
	PublicId  string `json:"publicId,omitempty"`
	PublicUrl string `json:"publicUrl,omitempty"`
	Inherited bool   `json:"inherited,omitempty"`
}

// NotificationModuleDependencyItemDto DTO.
type NotificationModuleDependencyItemDto struct {
	Name     string `json:"name,omitempty"`
	ViewId   string `json:"viewId,omitempty"`
	Category string `json:"category,omitempty"`
}

// NotificationModuleDisableDependenciesDto DTO.
type NotificationModuleDisableDependenciesDto struct {
	Triggers           []*NotificationModuleDependencyItemDto `json:"triggers,omitempty"`
	SchedulerTasks     []*NotificationModuleDependencyItemDto `json:"schedulerTasks,omitempty"`
	InFlightCampaigns  []*NotificationModuleDependencyItemDto `json:"inFlightCampaigns,omitempty"`
	MembershipSettings []*NotificationModuleDependencyItemDto `json:"membershipSettings,omitempty"`
}

// TestEmailValidationItemDto DTO.
type TestEmailValidationItemDto struct {
	Address string  `json:"address,omitempty"`
	Verdict string  `json:"verdict,omitempty"`
	Reason  string  `json:"reason,omitempty"`
	Score   float64 `json:"score,omitempty"`
}

// TemplateListProjection DTO.
type TemplateListProjection struct {
	Id           string               `json:"id,omitempty"`
	ViewId       string               `json:"viewId,omitempty"`
	TemplateName string               `json:"templateName,omitempty"`
	IsActive     bool                 `json:"isActive,omitempty"`
	Type         CommunicationChannel `json:"type,omitempty"`
	Tags         []string             `json:"tags,omitempty"`
}

// EmailTemplateListProjection DTO.
type EmailTemplateListProjection struct {
	TemplateListProjection
	HasAttachments bool     `json:"hasAttachments,omitempty"`
	Languages      []string `json:"languages,omitempty"`
}

// MjmlParseError DTO.
type MjmlParseError struct {
	Line             float64 `json:"line,omitempty"`
	Message          string  `json:"message,omitempty"`
	TagName          string  `json:"tagName,omitempty"`
	FormattedMessage string  `json:"formattedMessage,omitempty"`
}

// HtmlFromMjmlResponse DTO.
type HtmlFromMjmlResponse struct {
	Html   string            `json:"html,omitempty"`
	Errors []*MjmlParseError `json:"errors,omitempty"`
}

// SystemEmailTemplateListProjection DTO.
type SystemEmailTemplateListProjection struct {
	EmailTemplateListProjection
}

// EmailSignatureDto DTO.
type EmailSignatureDto struct {
	ViewId       string            `json:"viewId,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Translations []*TranslationDto `json:"translations,omitempty"`
}

// ListItemProjection DTO.
type ListItemProjection struct {
	ViewId      string `json:"viewId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// ListItemWithTranslationsProjection DTO.
type ListItemWithTranslationsProjection struct {
	ListItemProjection
	Translations []string `json:"translations,omitempty"`
}

// EmailFooterDto DTO.
type EmailFooterDto struct {
	ViewId       string            `json:"viewId,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Translations []*TranslationDto `json:"translations,omitempty"`
}

// EmailSettings DTO.
type EmailSettings struct {
	Signatures []*EmailSignatureDto `json:"signatures,omitempty"`
	Footers    []*EmailFooterDto    `json:"footers,omitempty"`
}

// DomainHealthRecordItemDto DTO.
type DomainHealthRecordItemDto struct {
	Record string `json:"record,omitempty"`
	Value  string `json:"value,omitempty"`
}

// EmailIntegrationListProjection DTO.
type EmailIntegrationListProjection struct {
	IntegrationListProjection
	EmailProvider      EmailProvider `json:"emailProvider,omitempty"`
	SenderEmailAddress string        `json:"senderEmailAddress,omitempty"`
	SenderDisplayName  string        `json:"senderDisplayName,omitempty"`
}

// CampaignStatusChangeEntryDto DTO.
type CampaignStatusChangeEntryDto struct {
	Time   string         `json:"time,omitempty"`
	Status CampaignStatus `json:"status,omitempty"`
	Errors []*ErrorDto    `json:"errors,omitempty"`
}

// CampaignDto DTO.
type CampaignDto struct {
	ViewId                          string                          `json:"viewId,omitempty"`
	CreatedOn                       string                          `json:"createdOn,omitempty"`
	Language                        string                          `json:"language,omitempty"`
	ForceCampaignLanguage           bool                            `json:"forceCampaignLanguage,omitempty"`
	CampaignProcessingIntegrationId string                          `json:"campaignProcessingIntegrationId,omitempty"`
	StatusHistory                   []*CampaignStatusChangeEntryDto `json:"statusHistory,omitempty"`
	Status                          *CampaignStatusChangeEntryDto   `json:"status,omitempty"`
	TokenMappingValues              []*TokenMappingDto              `json:"tokenMappingValues,omitempty"`
	Notes                           string                          `json:"notes,omitempty"`
	UserId                          string                          `json:"userId,omitempty"`
	Id                              string                          `json:"id,omitempty"`
}

// EmailCampaignDto DTO.
type EmailCampaignDto struct {
	CampaignDto
	DeliverySettings        *EmailCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
	Template                *EmailTemplateDto                 `json:"template,omitempty"`
	ValidationIntegrationId string                            `json:"validationIntegrationId,omitempty"`
	TemplateIsSystem        bool                              `json:"templateIsSystem,omitempty"`
}

// EmailCampaignListProjection DTO.
type EmailCampaignListProjection struct {
	ViewId        string         `json:"viewId,omitempty"`
	TemplateName  string         `json:"templateName,omitempty"`
	TemplateId    string         `json:"templateId,omitempty"`
	IntegrationId string         `json:"integrationId,omitempty"`
	Language      string         `json:"language,omitempty"`
	Strategy      string         `json:"strategy,omitempty"`
	LatestStatus  CampaignStatus `json:"latestStatus,omitempty"`
	CreatedOn     string         `json:"createdOn,omitempty"`
}

// BatchStatusChangeEntryDto DTO.
type BatchStatusChangeEntryDto struct {
	Time   string              `json:"time,omitempty"`
	Status CampaignBatchStatus `json:"status,omitempty"`
	Errors []*ErrorDto         `json:"errors,omitempty"`
}

// CampaignBatchDto DTO.
type CampaignBatchDto struct {
	CampaignId    string                       `json:"campaignId,omitempty"`
	BatchId       string                       `json:"batchId,omitempty"`
	StartAfter    string                       `json:"startAfter,omitempty"`
	StatusHistory []*BatchStatusChangeEntryDto `json:"statusHistory,omitempty"`
	Id            string                       `json:"id,omitempty"`
}

// EmailRecipientDto DTO.
type EmailRecipientDto struct {
	EmailAddress      string             `json:"emailAddress,omitempty"`
	Language          string             `json:"language,omitempty"`
	TimeZoneId        string             `json:"timeZoneId,omitempty"`
	UserTokenMappings []*TokenMappingDto `json:"userTokenMappings,omitempty"`
}

// EmailRecipientsDto DTO.
type EmailRecipientsDto struct {
	To            []*EmailRecipientDto `json:"to,omitempty"`
	Cc            []*EmailRecipientDto `json:"cc,omitempty"`
	Bcc           []*EmailRecipientDto `json:"bcc,omitempty"`
	StartingAfter string               `json:"startingAfter,omitempty"`
	HasMore       bool                 `json:"hasMore,omitempty"`
}

// EmailCampaignBatchDto DTO.
type EmailCampaignBatchDto struct {
	CampaignBatchDto
	Recipients *EmailRecipientsDto `json:"recipients,omitempty"`
}

// NotificationStatusChangeEntryDto DTO.
type NotificationStatusChangeEntryDto struct {
	Time     string                     `json:"time,omitempty"`
	Status   CampaignNotificationStatus `json:"status,omitempty"`
	SourceId string                     `json:"sourceId,omitempty"`
	Errors   []*ErrorDto                `json:"errors,omitempty"`
	Tags     []string                   `json:"tags,omitempty"`
}

// CampaignBatchNotificationDto DTO.
type CampaignBatchNotificationDto struct {
	CampaignId        string                              `json:"campaignId,omitempty"`
	BatchId           string                              `json:"batchId,omitempty"`
	NotificationId    string                              `json:"notificationId,omitempty"`
	RefNotificationId string                              `json:"refNotificationId,omitempty"`
	Subject           string                              `json:"subject,omitempty"`
	Body              string                              `json:"body,omitempty"`
	Model             map[string]string                   `json:"model,omitempty"`
	StatusHistory     []*NotificationStatusChangeEntryDto `json:"statusHistory,omitempty"`
	Id                string                              `json:"id,omitempty"`
}

// EmailCampaignBatchNotificationDto DTO.
type EmailCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	Recipients *EmailRecipientsDto     `json:"recipients,omitempty"`
	Content    *EmailMessageContentDto `json:"content,omitempty"`
}

// CampaignStatsDto DTO.
type CampaignStatsDto struct {
	Batches     float64 `json:"batches,omitempty"`
	Sent        float64 `json:"sent,omitempty"`
	Failed      float64 `json:"failed,omitempty"`
	SuccessRate float64 `json:"successRate,omitempty"`
}

// SmsTemplateListProjection DTO.
type SmsTemplateListProjection struct {
	TemplateListProjection
}

// SmsSettings DTO.
type SmsSettings struct {
}

// SmsIntegrationListProjection DTO.
type SmsIntegrationListProjection struct {
	IntegrationListProjection
	Provider SmsProvider `json:"provider,omitempty"`
}

// SmsCampaignDto DTO.
type SmsCampaignDto struct {
	CampaignDto
	Recipients *SmsCampaignDeliverySettingsDto `json:"recipients,omitempty"`
	Template   *SmsTemplateDto                 `json:"template,omitempty"`
}

// SmsRecipientDto DTO.
type SmsRecipientDto struct {
	PhoneNumber       string             `json:"phoneNumber,omitempty"`
	UserId            string             `json:"userId,omitempty"`
	Language          string             `json:"language,omitempty"`
	UserTokenMappings []*TokenMappingDto `json:"userTokenMappings,omitempty"`
	TimeZoneId        string             `json:"timeZoneId,omitempty"`
	Record            string             `json:"record,omitempty"`
}

// SmsRecipientsDto DTO.
type SmsRecipientsDto struct {
	To            []*SmsRecipientDto `json:"to,omitempty"`
	StartingAfter string             `json:"startingAfter,omitempty"`
	HasMore       bool               `json:"hasMore,omitempty"`
}

// SmsCampaignBatchDto DTO.
type SmsCampaignBatchDto struct {
	CampaignBatchDto
	Recipients *SmsRecipientsDto `json:"recipients,omitempty"`
}

// SmsCampaignBatchNotificationDto DTO.
type SmsCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	Recipients *SmsRecipientsDto     `json:"recipients,omitempty"`
	Content    *SmsMessageContentDto `json:"content,omitempty"`
}

// MarketplaceListingProjection DTO.
type MarketplaceListingProjection struct {
	ViewId        string               `json:"viewId,omitempty"`
	Slug          string               `json:"slug,omitempty"`
	DisplayName   string               `json:"displayName,omitempty"`
	Vendor        string               `json:"vendor,omitempty"`
	Category      MarketplaceCategory  `json:"category,omitempty"`
	Transport     MarketplaceTransport `json:"transport,omitempty"`
	IconUrl       string               `json:"iconUrl,omitempty"`
	IsOfficial    bool                 `json:"isOfficial,omitempty"`
	Tags          []string             `json:"tags,omitempty"`
	FunctionCount float64              `json:"functionCount,omitempty"`
}

// MarketplaceIntegrationListProjection DTO.
type MarketplaceIntegrationListProjection struct {
	ViewId                            string               `json:"viewId,omitempty"`
	IntegrationName                   string               `json:"integrationName,omitempty"`
	IsEnabled                         bool                 `json:"isEnabled,omitempty"`
	ListingViewId                     string               `json:"listingViewId,omitempty"`
	Vendor                            string               `json:"vendor,omitempty"`
	Category                          MarketplaceCategory  `json:"category,omitempty"`
	Transport                         MarketplaceTransport `json:"transport,omitempty"`
	LastIntegrationTestAtUtc          string               `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      bool                 `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         []string             `json:"lastIntegrationTestErrors,omitempty"`
	HumanDeliveryConfirmedAtUtc       string               `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool                 `json:"requiresHumanDeliveryConfirmation,omitempty"`
}

// MarketplaceFunctionProjection DTO.
type MarketplaceFunctionProjection struct {
	ViewId            string  `json:"viewId,omitempty"`
	IntegrationViewId string  `json:"integrationViewId,omitempty"`
	FunctionKey       string  `json:"functionKey,omitempty"`
	DisplayName       string  `json:"displayName,omitempty"`
	IsEnabled         bool    `json:"isEnabled,omitempty"`
	MappingCount      float64 `json:"mappingCount,omitempty"`
}

// CodeIntegrationListProjection DTO.
type CodeIntegrationListProjection struct {
	IntegrationListProjection
	Provider CodeProvider `json:"provider,omitempty"`
}

// PushTemplateListProjection DTO.
type PushTemplateListProjection struct {
	TemplateListProjection
}

// PushSettings DTO.
type PushSettings struct {
	MarketingTags     []*TagDefinitionDto `json:"marketingTags,omitempty"`
	TransactionalTags []*TagDefinitionDto `json:"transactionalTags,omitempty"`
}

// PushIntegrationListProjection DTO.
type PushIntegrationListProjection struct {
	IntegrationListProjection
	Provider PushProvider `json:"provider,omitempty"`
}

// PushCampaignDto DTO.
type PushCampaignDto struct {
	CampaignDto
	Recipients *PushCampaignDeliverySettingsDto `json:"recipients,omitempty"`
	Template   *PushTemplateDto                 `json:"template,omitempty"`
}

// PushRecipientDto DTO.
type PushRecipientDto struct {
	DeviceTokens      []*PushDeviceDeliveryTokenDto `json:"deviceTokens,omitempty"`
	UserId            string                        `json:"userId,omitempty"`
	Language          string                        `json:"language,omitempty"`
	UserTokenMappings []*TokenMappingDto            `json:"userTokenMappings,omitempty"`
	TimeZoneId        string                        `json:"timeZoneId,omitempty"`
	Record            string                        `json:"record,omitempty"`
}

// PushRecipientsDto DTO.
type PushRecipientsDto struct {
	To            []*PushRecipientDto `json:"to,omitempty"`
	StartingAfter string              `json:"startingAfter,omitempty"`
	HasMore       bool                `json:"hasMore,omitempty"`
}

// PushCampaignBatchDto DTO.
type PushCampaignBatchDto struct {
	CampaignBatchDto
	Recipients *PushRecipientsDto `json:"recipients,omitempty"`
}

// PushCampaignBatchNotificationDto DTO.
type PushCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	Recipients *PushRecipientsDto     `json:"recipients,omitempty"`
	Content    *PushMessageContentDto `json:"content,omitempty"`
}

// PaymentsWebhookLogEntry DTO.
type PaymentsWebhookLogEntry struct {
	IntegrationId   string  `json:"integrationId,omitempty"`
	Source          string  `json:"source,omitempty"`
	EventName       string  `json:"eventName,omitempty"`
	ProviderEventId string  `json:"providerEventId,omitempty"`
	StatusCode      float64 `json:"statusCode,omitempty"`
	Description     string  `json:"description,omitempty"`
	ReceivedOn      string  `json:"receivedOn,omitempty"`
}

// PaymentTriggerProjectionList DTO.
type PaymentTriggerProjectionList struct {
	TriggerProjectionList
	Type         PaymentTriggerType `json:"type,omitempty"`
	Integrations []string           `json:"integrations,omitempty"`
	Events       []string           `json:"events,omitempty"`
}

// PaymentsIntegrationListProjection DTO.
type PaymentsIntegrationListProjection struct {
	IntegrationListProjection
	GatewayPlatform PaymentGatewayPlatform `json:"gatewayPlatform,omitempty"`
}

// LoggingIntegrationListProjection DTO.
type LoggingIntegrationListProjection struct {
	IntegrationListProjection
	Provider LoggingProvider `json:"provider,omitempty"`
}

// TenantLogEntryDto DTO.
type TenantLogEntryDto struct {
	Id            string            `json:"id,omitempty"`
	Timestamp     string            `json:"timestamp,omitempty"`
	Module        string            `json:"module,omitempty"`
	Level         string            `json:"level,omitempty"`
	EventCode     string            `json:"eventCode,omitempty"`
	Title         string            `json:"title,omitempty"`
	Message       string            `json:"message,omitempty"`
	CorrelationId string            `json:"correlationId,omitempty"`
	TraceId       string            `json:"traceId,omitempty"`
	SpanId        string            `json:"spanId,omitempty"`
	Meta          map[string]string `json:"meta,omitempty"`
}

// AiToolManifestParameter DTO.
type AiToolManifestParameter struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Description string `json:"description,omitempty"`
}

// AiToolManifestItem DTO.
type AiToolManifestItem struct {
	Name                 string                     `json:"name,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Toolsets             []string                   `json:"toolsets,omitempty"`
	RequiresConfirmation bool                       `json:"requiresConfirmation,omitempty"`
	Parameters           []*AiToolManifestParameter `json:"parameters,omitempty"`
}

// ChatModelOption DTO.
type ChatModelOption struct {
	LlmIntegrationId string  `json:"llmIntegrationId,omitempty"`
	Kind             string  `json:"kind,omitempty"`
	Provider         string  `json:"provider,omitempty"`
	Model            string  `json:"model,omitempty"`
	Label            string  `json:"label,omitempty"`
	IsDefault        bool    `json:"isDefault,omitempty"`
	IsAuto           bool    `json:"isAuto,omitempty"`
	ContextWindow    float64 `json:"contextWindow,omitempty"`
}

// ChatMemoryNote DTO.
type ChatMemoryNote struct {
	Id           string `json:"id,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Text         string `json:"text,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	CreatedAtUtc string `json:"createdAtUtc,omitempty"`
}

// ChatSessionListItem DTO.
type ChatSessionListItem struct {
	SessionId    string `json:"sessionId,omitempty"`
	Profile      string `json:"profile,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	Env          string `json:"env,omitempty"`
	Title        string `json:"title,omitempty"`
	UpdatedAtUtc string `json:"updatedAtUtc,omitempty"`
	IsArchived   bool   `json:"isArchived,omitempty"`
	IsPinned     bool   `json:"isPinned,omitempty"`
}

// ProjectBriefSourceWireDto DTO.
type ProjectBriefSourceWireDto struct {
	Kind       string  `json:"kind,omitempty"`
	SessionId  string  `json:"sessionId,omitempty"`
	EntryId    string  `json:"entryId,omitempty"`
	EntrySeq   float64 `json:"entrySeq,omitempty"`
	EventId    string  `json:"eventId,omitempty"`
	UserAuthId string  `json:"userAuthId,omitempty"`
	AtUtc      string  `json:"atUtc,omitempty"`
	Surface    string  `json:"surface,omitempty"`
	Quote      string  `json:"quote,omitempty"`
	WorkItemId string  `json:"workItemId,omitempty"`
}

// ProjectBriefSatisfiedByWireDto DTO.
type ProjectBriefSatisfiedByWireDto struct {
	ArtifactId string `json:"artifactId,omitempty"`
	Tool       string `json:"tool,omitempty"`
	Orphaned   bool   `json:"orphaned,omitempty"`
}

// ProjectBriefRequirementWireDto DTO.
type ProjectBriefRequirementWireDto struct {
	Id           string                            `json:"id,omitempty"`
	Text         string                            `json:"text,omitempty"`
	Status       string                            `json:"status,omitempty"`
	Confidence   float64                           `json:"confidence,omitempty"`
	IsAssumption bool                              `json:"isAssumption,omitempty"`
	Sources      []*ProjectBriefSourceWireDto      `json:"sources,omitempty"`
	SatisfiedBy  []*ProjectBriefSatisfiedByWireDto `json:"satisfiedBy,omitempty"`
	SinceEventId string                            `json:"sinceEventId,omitempty"`
}

// ProjectBriefDecisionWireDto DTO.
type ProjectBriefDecisionWireDto struct {
	EventId    string                       `json:"eventId,omitempty"`
	Text       string                       `json:"text,omitempty"`
	Confidence float64                      `json:"confidence,omitempty"`
	AtUtc      string                       `json:"atUtc,omitempty"`
	Sources    []*ProjectBriefSourceWireDto `json:"sources,omitempty"`
}

// ProjectBriefAssumptionWireDto DTO.
type ProjectBriefAssumptionWireDto struct {
	RequirementId string  `json:"requirementId,omitempty"`
	EventId       string  `json:"eventId,omitempty"`
	Text          string  `json:"text,omitempty"`
	Confidence    float64 `json:"confidence,omitempty"`
	AtUtc         string  `json:"atUtc,omitempty"`
}

// ProjectBriefSnapshotWireDto DTO.
type ProjectBriefSnapshotWireDto struct {
	ProjectId       string                            `json:"projectId,omitempty"`
	UpToSeq         float64                           `json:"upToSeq,omitempty"`
	AtUtc           string                            `json:"atUtc,omitempty"`
	Requirements    []*ProjectBriefRequirementWireDto `json:"requirements,omitempty"`
	Decisions       []*ProjectBriefDecisionWireDto    `json:"decisions,omitempty"`
	OpenAssumptions []*ProjectBriefAssumptionWireDto  `json:"openAssumptions,omitempty"`
	Summary         string                            `json:"summary,omitempty"`
}

// ProjectBriefEventWireDto DTO.
type ProjectBriefEventWireDto struct {
	Id                string                            `json:"id,omitempty"`
	ProjectId         string                            `json:"projectId,omitempty"`
	Seq               float64                           `json:"seq,omitempty"`
	AtUtc             string                            `json:"atUtc,omitempty"`
	Kind              string                            `json:"kind,omitempty"`
	RequirementId     string                            `json:"requirementId,omitempty"`
	SupersedesEventId string                            `json:"supersedesEventId,omitempty"`
	Text              string                            `json:"text,omitempty"`
	Confidence        float64                           `json:"confidence,omitempty"`
	Sources           []*ProjectBriefSourceWireDto      `json:"sources,omitempty"`
	Origin            string                            `json:"origin,omitempty"`
	SatisfiedBy       []*ProjectBriefSatisfiedByWireDto `json:"satisfiedBy,omitempty"`
}

// WorkItemEntryRefWireDto DTO.
type WorkItemEntryRefWireDto struct {
	SessionId string `json:"sessionId,omitempty"`
	EntryId   string `json:"entryId,omitempty"`
}

// WorkItemRunRefWireDto DTO.
type WorkItemRunRefWireDto struct {
	SessionId   string `json:"sessionId,omitempty"`
	PlanEntryId string `json:"planEntryId,omitempty"`
	RunId       string `json:"runId,omitempty"`
}

// WorkItemArtifactWireDto DTO.
type WorkItemArtifactWireDto struct {
	ArtifactId  string  `json:"artifactId,omitempty"`
	What        string  `json:"what,omitempty"`
	Step        float64 `json:"step,omitempty"`
	PlanEntryId string  `json:"planEntryId,omitempty"`
	Kind        string  `json:"kind,omitempty"`
	Name        string  `json:"name,omitempty"`
}

// WorkItemMovedOutWireDto DTO.
type WorkItemMovedOutWireDto struct {
	Text     string                   `json:"text,omitempty"`
	Reason   string                   `json:"reason,omitempty"`
	MovedTo  string                   `json:"movedTo,omitempty"`
	Source   string                   `json:"source,omitempty"`
	EntryRef *WorkItemEntryRefWireDto `json:"entryRef,omitempty"`
	AtUtc    string                   `json:"atUtc,omitempty"`
}

// WorkItemNeedsYouWireDto DTO.
type WorkItemNeedsYouWireDto struct {
	Text             string                   `json:"text,omitempty"`
	Kind             string                   `json:"kind,omitempty"`
	EntryRef         *WorkItemEntryRefWireDto `json:"entryRef,omitempty"`
	Done             bool                     `json:"done,omitempty"`
	DoneAtUtc        string                   `json:"doneAtUtc,omitempty"`
	DoneByUserAuthId string                   `json:"doneByUserAuthId,omitempty"`
}

// WorkItemOpenQuestionWireDto DTO.
type WorkItemOpenQuestionWireDto struct {
	EntryRef *WorkItemEntryRefWireDto `json:"entryRef,omitempty"`
	Blocking bool                     `json:"blocking,omitempty"`
	Text     string                   `json:"text,omitempty"`
	AtUtc    string                   `json:"atUtc,omitempty"`
}

// WorkItemDoneConditionWireDto DTO.
type WorkItemDoneConditionWireDto struct {
	Condition float64 `json:"condition,omitempty"`
	Holds     bool    `json:"holds,omitempty"`
	Reason    string  `json:"reason,omitempty"`
}

// WorkItemWireDto DTO.
type WorkItemWireDto struct {
	Id                  string                          `json:"id,omitempty"`
	ProjectId           string                          `json:"projectId,omitempty"`
	Status              string                          `json:"status,omitempty"`
	Goal                string                          `json:"goal,omitempty"`
	NotInScope          []string                        `json:"notInScope,omitempty"`
	ScopeRequirementIds []string                        `json:"scopeRequirementIds,omitempty"`
	Difficulty          float64                         `json:"difficulty,omitempty"`
	DifficultyReason    string                          `json:"difficultyReason,omitempty"`
	PlanEntryRefs       []*WorkItemEntryRefWireDto      `json:"planEntryRefs,omitempty"`
	RunRefs             []*WorkItemRunRefWireDto        `json:"runRefs,omitempty"`
	Artifacts           []*WorkItemArtifactWireDto      `json:"artifacts,omitempty"`
	MovedOut            []*WorkItemMovedOutWireDto      `json:"movedOut,omitempty"`
	NeedsYou            []*WorkItemNeedsYouWireDto      `json:"needsYou,omitempty"`
	OpenQuestions       []*WorkItemOpenQuestionWireDto  `json:"openQuestions,omitempty"`
	SessionIds          []string                        `json:"sessionIds,omitempty"`
	ParentId            string                          `json:"parentId,omitempty"`
	Children            []string                        `json:"children,omitempty"`
	SummaryEntryRef     *WorkItemEntryRefWireDto        `json:"summaryEntryRef,omitempty"`
	CreatedBy           string                          `json:"createdBy,omitempty"`
	CreatedAtUtc        string                          `json:"createdAtUtc,omitempty"`
	UpdatedAtUtc        string                          `json:"updatedAtUtc,omitempty"`
	DoneVerdict         string                          `json:"doneVerdict,omitempty"`
	DoneConditions      []*WorkItemDoneConditionWireDto `json:"doneConditions,omitempty"`
}

// LlmIntegrationListProjection DTO.
type LlmIntegrationListProjection struct {
	IntegrationListProjection
	LlmProvider  LlmProvider `json:"llmProvider,omitempty"`
	BaseUrl      string      `json:"baseUrl,omitempty"`
	DefaultModel string      `json:"defaultModel,omitempty"`
}

// McpIntegrationListProjection DTO.
type McpIntegrationListProjection struct {
	IntegrationListProjection
	McpProvider McpProvider  `json:"mcpProvider,omitempty"`
	Transport   McpTransport `json:"transport,omitempty"`
	Category    string       `json:"category,omitempty"`
	Description string       `json:"description,omitempty"`
	Icon        string       `json:"icon,omitempty"`
}

// IVirtualDirectory DTO.
type IVirtualDirectory struct {
}

// IVirtualPathProvider DTO.
type IVirtualPathProvider struct {
}

// IVirtualFile DTO.
type IVirtualFile struct {
}

// IContentTypeWriter DTO.
type IContentTypeWriter struct {
}

// IRequestPreferences DTO.
type IRequestPreferences struct {
}

// IHttpFile DTO.
type IHttpFile struct {
}

// IRequest DTO.
type IRequest struct {
}

// IResponse DTO.
type IResponse struct {
}

// SchedulerTaskListProjection DTO.
type SchedulerTaskListProjection struct {
	TaskId    string            `json:"taskId,omitempty"`
	Name      string            `json:"name,omitempty"`
	Cron      string            `json:"cron,omitempty"`
	Type      SchedulerTaskType `json:"type,omitempty"`
	IsEnabled bool              `json:"isEnabled,omitempty"`
	ViewId    string            `json:"viewId,omitempty"`
}

// ResolvedResourceEntry DTO.
type ResolvedResourceEntry struct {
	Ref        *ResourceRefDto   `json:"ref,omitempty"`
	Status     ResolvedRefStatus `json:"status,omitempty"`
	Resolved   map[string]any    `json:"resolved,omitempty"`
	Diagnostic string            `json:"diagnostic,omitempty"`
}

// ChannelSubscriptionStateDto DTO.
type ChannelSubscriptionStateDto struct {
	Unsubscribed bool                `json:"unsubscribed,omitempty"`
	BlockedTags  map[string][]string `json:"blockedTags,omitempty"`
}

// UserDto DTO.
type UserDto struct {
	Id                   string                                  `json:"id,omitempty"`
	ProjectId            string                                  `json:"projectId,omitempty"`
	PrimaryEmail         string                                  `json:"primaryEmail,omitempty"`
	PrimaryPhone         string                                  `json:"primaryPhone,omitempty"`
	DisplayName          string                                  `json:"displayName,omitempty"`
	FirstName            string                                  `json:"firstName,omitempty"`
	LastName             string                                  `json:"lastName,omitempty"`
	FullName             string                                  `json:"fullName,omitempty"`
	Company              string                                  `json:"company,omitempty"`
	Locale               string                                  `json:"locale,omitempty"`
	TimeZone             string                                  `json:"timeZone,omitempty"`
	Gender               string                                  `json:"gender,omitempty"`
	BirthDate            float64                                 `json:"birthDate,omitempty"`
	AddressLine1         string                                  `json:"addressLine1,omitempty"`
	AddressLine2         string                                  `json:"addressLine2,omitempty"`
	Country              string                                  `json:"country,omitempty"`
	City                 string                                  `json:"city,omitempty"`
	State                string                                  `json:"state,omitempty"`
	PostalCode           string                                  `json:"postalCode,omitempty"`
	Tags                 []string                                `json:"tags,omitempty"`
	Roles                []string                                `json:"roles,omitempty"`
	Lifecycle            string                                  `json:"lifecycle,omitempty"`
	SourceOfCreation     string                                  `json:"sourceOfCreation,omitempty"`
	MergedIntoContactId  string                                  `json:"mergedIntoContactId,omitempty"`
	CreatedOn            string                                  `json:"createdOn,omitempty"`
	ModifiedOn           string                                  `json:"modifiedOn,omitempty"`
	Auths                []*AuthDto                              `json:"auths,omitempty"`
	MarketingPreferences map[string]*ChannelSubscriptionStateDto `json:"marketingPreferences,omitempty"`
}

// ConsentPurposeDto DTO.
type ConsentPurposeDto struct {
	Key             string   `json:"key,omitempty"`
	Name            string   `json:"name,omitempty"`
	Channel         string   `json:"channel,omitempty"`
	MappedTags      []string `json:"mappedTags,omitempty"`
	RegulatoryBasis []string `json:"regulatoryBasis,omitempty"`
	Description     string   `json:"description,omitempty"`
	IsDeprecated    bool     `json:"isDeprecated,omitempty"`
}

// RetentionWindowDto DTO.
type RetentionWindowDto struct {
	DataKind string  `json:"dataKind,omitempty"`
	Days     float64 `json:"days,omitempty"`
	Action   string  `json:"action,omitempty"`
}

// ProjectComplianceDto DTO.
type ProjectComplianceDto struct {
	Regimes          []string              `json:"regimes,omitempty"`
	ConsentPurposes  []*ConsentPurposeDto  `json:"consentPurposes,omitempty"`
	RetentionWindows []*RetentionWindowDto `json:"retentionWindows,omitempty"`
}

// LegalHoldDto DTO.
type LegalHoldDto struct {
	Id          string `json:"id,omitempty"`
	SubjectKind string `json:"subjectKind,omitempty"`
	SubjectId   string `json:"subjectId,omitempty"`
	Reason      string `json:"reason,omitempty"`
	PlacedAt    string `json:"placedAt,omitempty"`
	PlacedBy    string `json:"placedBy,omitempty"`
	ReleasedAt  string `json:"releasedAt,omitempty"`
	ReleasedBy  string `json:"releasedBy,omitempty"`
}

// DsarRequestDto DTO.
type DsarRequestDto struct {
	Id              string `json:"id,omitempty"`
	SubjectKind     string `json:"subjectKind,omitempty"`
	SubjectId       string `json:"subjectId,omitempty"`
	Status          string `json:"status,omitempty"`
	ReceivedAt      string `json:"receivedAt,omitempty"`
	SlaDeadline     string `json:"slaDeadline,omitempty"`
	AutoApproveAt   string `json:"autoApproveAt,omitempty"`
	DecidedAt       string `json:"decidedAt,omitempty"`
	DecidedBy       string `json:"decidedBy,omitempty"`
	RejectionReason string `json:"rejectionReason,omitempty"`
}

// ComplianceAuditEntryDto DTO.
type ComplianceAuditEntryDto struct {
	Id          string            `json:"id,omitempty"`
	Timestamp   string            `json:"timestamp,omitempty"`
	Action      string            `json:"action,omitempty"`
	SubjectKind string            `json:"subjectKind,omitempty"`
	SubjectId   string            `json:"subjectId,omitempty"`
	Reason      string            `json:"reason,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// AccountComplianceDto DTO.
type AccountComplianceDto struct {
	DsarMode              string  `json:"dsarMode,omitempty"`
	DsarDelayDays         float64 `json:"dsarDelayDays,omitempty"`
	AutoForwardAdvisories bool    `json:"autoForwardAdvisories,omitempty"`
	SecurityContact       string  `json:"securityContact,omitempty"`
}

// SupportCaseDto DTO.
type SupportCaseDto struct {
	ViewId          string                `json:"viewId,omitempty"`
	AccountId       string                `json:"accountId,omitempty"`
	ProjectId       string                `json:"projectId,omitempty"`
	ReporterId      string                `json:"reporterId,omitempty"`
	Kind            SupportCaseKind       `json:"kind,omitempty"`
	Severity        SupportCaseSeverity   `json:"severity,omitempty"`
	Status          SupportCaseStatus     `json:"status,omitempty"`
	CustomerStatus  SupportCustomerStatus `json:"customerStatus,omitempty"`
	Subject         string                `json:"subject,omitempty"`
	AffectedModule  string                `json:"affectedModule,omitempty"`
	DeploymentMode  DeploymentMode        `json:"deploymentMode,omitempty"`
	GatewayVersion  string                `json:"gatewayVersion,omitempty"`
	Region          string                `json:"region,omitempty"`
	PlanTier        string                `json:"planTier,omitempty"`
	OpenedOn        float64               `json:"openedOn,omitempty"`
	FirstResponseOn float64               `json:"firstResponseOn,omitempty"`
	ResolvedOn      float64               `json:"resolvedOn,omitempty"`
	ClosedOn        float64               `json:"closedOn,omitempty"`
	Resolution      string                `json:"resolution,omitempty"`
	MessageCount    float64               `json:"messageCount,omitempty"`
	LastMessageOn   float64               `json:"lastMessageOn,omitempty"`
}

// SupportCaseMessageDto DTO.
type SupportCaseMessageDto struct {
	Id                string                   `json:"id,omitempty"`
	CaseId            string                   `json:"caseId,omitempty"`
	AuthorKind        SupportMessageAuthorKind `json:"authorKind,omitempty"`
	AuthorId          string                   `json:"authorId,omitempty"`
	AuthorDisplayName string                   `json:"authorDisplayName,omitempty"`
	Body              string                   `json:"body,omitempty"`
	SentOn            float64                  `json:"sentOn,omitempty"`
}

// SupportCaseDetailDto DTO.
type SupportCaseDetailDto struct {
	Case     *SupportCaseDto          `json:"case,omitempty"`
	Messages []*SupportCaseMessageDto `json:"messages,omitempty"`
}

// SupportCaseListProjection DTO.
type SupportCaseListProjection struct {
	ViewId         string                `json:"viewId,omitempty"`
	ProjectId      string                `json:"projectId,omitempty"`
	Kind           SupportCaseKind       `json:"kind,omitempty"`
	Severity       SupportCaseSeverity   `json:"severity,omitempty"`
	Status         SupportCaseStatus     `json:"status,omitempty"`
	CustomerStatus SupportCustomerStatus `json:"customerStatus,omitempty"`
	Subject        string                `json:"subject,omitempty"`
	OpenedOn       float64               `json:"openedOn,omitempty"`
	MessageCount   float64               `json:"messageCount,omitempty"`
	LastMessageOn  float64               `json:"lastMessageOn,omitempty"`
}

// DiagnosticPackStepDescriptorDto DTO.
type DiagnosticPackStepDescriptorDto struct {
	StepId      string            `json:"stepId,omitempty"`
	Kind        string            `json:"kind,omitempty"`
	Description string            `json:"description,omitempty"`
	Parameters  map[string]string `json:"parameters,omitempty"`
}

// DiagnosticPackDescriptorDto DTO.
type DiagnosticPackDescriptorDto struct {
	Name    string                             `json:"name,omitempty"`
	Version float64                            `json:"version,omitempty"`
	Summary string                             `json:"summary,omitempty"`
	Steps   []*DiagnosticPackStepDescriptorDto `json:"steps,omitempty"`
}

// DiagnosticPackStepResultDto DTO.
type DiagnosticPackStepResultDto struct {
	StepId       string         `json:"stepId,omitempty"`
	Kind         string         `json:"kind,omitempty"`
	IsSuccess    bool           `json:"isSuccess,omitempty"`
	Result       map[string]any `json:"result,omitempty"`
	ErrorMessage string         `json:"errorMessage,omitempty"`
}

// DiagnosticPackRunResultDto DTO.
type DiagnosticPackRunResultDto struct {
	PackName    string                         `json:"packName,omitempty"`
	PackVersion float64                        `json:"packVersion,omitempty"`
	CaseId      string                         `json:"caseId,omitempty"`
	Steps       []*DiagnosticPackStepResultDto `json:"steps,omitempty"`
}

// DiagnosticEchoRegionDto DTO.
type DiagnosticEchoRegionDto struct {
}

// DiagnosticEchoDto DTO.
type DiagnosticEchoDto struct {
	ContainerName    string                     `json:"containerName,omitempty"`
	IsManagedService bool                       `json:"isManagedService,omitempty"`
	ApiVersion       string                     `json:"apiVersion,omitempty"`
	HubVersion       string                     `json:"hubVersion,omitempty"`
	Release          string                     `json:"release,omitempty"`
	Runtime          string                     `json:"runtime,omitempty"`
	HubUrl           string                     `json:"hubUrl,omitempty"`
	ApiUrl           string                     `json:"apiUrl,omitempty"`
	LicensePresent   bool                       `json:"licensePresent,omitempty"`
	Regions          []*DiagnosticEchoRegionDto `json:"regions,omitempty"`
}

// DiagnosticEventItemDto DTO.
type DiagnosticEventItemDto struct {
	Position  float64           `json:"position,omitempty"`
	EventType string            `json:"eventType,omitempty"`
	Payload   map[string]string `json:"payload,omitempty"`
}

// DiagnosticEventsPageDto DTO.
type DiagnosticEventsPageDto struct {
	Stream   string                    `json:"stream,omitempty"`
	From     float64                   `json:"from,omitempty"`
	Count    float64                   `json:"count,omitempty"`
	HasMore  bool                      `json:"hasMore,omitempty"`
	NextFrom float64                   `json:"nextFrom,omitempty"`
	Items    []*DiagnosticEventItemDto `json:"items,omitempty"`
}

// DiagnosticLogsResponse DTO.
type DiagnosticLogsResponse struct {
	List PaginatedResponse[*TenantLogEntryDto] `json:"list,omitempty"`
}

// DiagnosticRedisListItemDto DTO.
type DiagnosticRedisListItemDto struct {
	ViewId string `json:"viewId,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}

// DiagnosticRedisInspectDto DTO.
type DiagnosticRedisInspectDto struct {
	KeyPattern    string                        `json:"keyPattern,omitempty"`
	CacheKey      string                        `json:"cacheKey,omitempty"`
	IsList        bool                          `json:"isList,omitempty"`
	Item          map[string]map[string]any     `json:"item,omitempty"`
	ListItems     []*DiagnosticRedisListItemDto `json:"listItems,omitempty"`
	HasMore       bool                          `json:"hasMore,omitempty"`
	StartingAfter string                        `json:"startingAfter,omitempty"`
}

// DiagnosticHealthCheckDto DTO.
type DiagnosticHealthCheckDto struct {
	CheckId    string  `json:"checkId,omitempty"`
	IsHealthy  bool    `json:"isHealthy,omitempty"`
	StatusCode float64 `json:"statusCode,omitempty"`
	Detail     string  `json:"detail,omitempty"`
}

// ResponseError DTO.
type ResponseError struct {
	ErrorCode string            `json:"errorCode,omitempty"`
	FieldName string            `json:"fieldName,omitempty"`
	Message   string            `json:"message,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
}

// ResponseStatus DTO.
type ResponseStatus struct {
	ErrorCode  string            `json:"errorCode,omitempty"`
	Message    string            `json:"message,omitempty"`
	StackTrace string            `json:"stackTrace,omitempty"`
	Errors     []*ResponseError  `json:"errors,omitempty"`
	Meta       map[string]string `json:"meta,omitempty"`
}

// ILlmApiKeyRequest DTO.
type ILlmApiKeyRequest struct {
}

// IHasViewId DTO.
type IHasViewId struct {
}

// IHasDatabaseId DTO.
type IHasDatabaseId struct {
}

// IBindableContract DTO.
type IBindableContract struct {
}

// IHasRazorTemplateCode DTO.
type IHasRazorTemplateCode struct {
}

// IHasDomainEntityId DTO.
type IHasDomainEntityId struct {
}

// IIntegrationIdentification DTO.
type IIntegrationIdentification struct {
}

// IHasResponsibleUserId DTO.
type IHasResponsibleUserId struct {
}

// ICursorArgs DTO.
type ICursorArgs struct {
}

// StringField DTO.
type StringField struct {
	JsonSchemaField
	Format           string            `json:"format,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MinLength        float64           `json:"minLength,omitempty"`
	MaxLength        float64           `json:"maxLength,omitempty"`
	TranslateOptions map[string]string `json:"translateOptions,omitempty"`
}

// DecimalField DTO.
type DecimalField struct {
	JsonSchemaField
	Minimum    float64 `json:"minimum,omitempty"`
	Maximum    float64 `json:"maximum,omitempty"`
	MultipleOf float64 `json:"multipleOf,omitempty"`
}

// CurrencyField DTO.
type CurrencyField struct {
	JsonSchemaField
	AllowedCurrencies []string `json:"allowedCurrencies,omitempty"`
}

// BooleanField DTO.
type BooleanField struct {
	JsonSchemaField
}

// DateField DTO.
type DateField struct {
	JsonSchemaField
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

// IntegerField DTO.
type IntegerField struct {
	JsonSchemaField
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

// GeolocationField DTO.
type GeolocationField struct {
	JsonSchemaField
	AllowedTypes []string `json:"allowedTypes,omitempty"`
}

// TagsField DTO.
type TagsField struct {
	JsonSchemaField
}

// FileField DTO.
type FileField struct {
	JsonSchemaField
	Storages []string `json:"storages,omitempty"`
}

// TaxonomySelectionField DTO.
type TaxonomySelectionField struct {
	JsonSchemaField
	TaxonomyId string `json:"taxonomyId,omitempty"`
	Multiple   bool   `json:"multiple,omitempty"`
}

// CollectionSelectionField DTO.
type CollectionSelectionField struct {
	JsonSchemaField
	CollectionId string `json:"collectionId,omitempty"`
	DisplayField string `json:"displayField,omitempty"`
	Multiple     bool   `json:"multiple,omitempty"`
}

// UserSelectionField DTO.
type UserSelectionField struct {
	JsonSchemaField
	Multiple bool `json:"multiple,omitempty"`
}

// RoleSelectionField DTO.
type RoleSelectionField struct {
	JsonSchemaField
	Multiple bool `json:"multiple,omitempty"`
}

// EnumSelectionField DTO.
type EnumSelectionField struct {
	JsonSchemaField
	Values   []string `json:"values,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
}

// StringFieldDto DTO.
type StringFieldDto struct {
	JsonSchemaFieldDto
	Format           string            `json:"format,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MinLength        float64           `json:"minLength,omitempty"`
	MaxLength        float64           `json:"maxLength,omitempty"`
	TranslateOptions map[string]string `json:"translateOptions,omitempty"`
}

// DecimalFieldDto DTO.
type DecimalFieldDto struct {
	JsonSchemaFieldDto
	Minimum    float64 `json:"minimum,omitempty"`
	Maximum    float64 `json:"maximum,omitempty"`
	MultipleOf float64 `json:"multipleOf,omitempty"`
}

// CurrencyFieldDto DTO.
type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	AllowedCurrencies []string `json:"allowedCurrencies,omitempty"`
}

// BooleanFieldDto DTO.
type BooleanFieldDto struct {
	JsonSchemaFieldDto
}

// DateFieldDto DTO.
type DateFieldDto struct {
	JsonSchemaFieldDto
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

// IntegerFieldDto DTO.
type IntegerFieldDto struct {
	JsonSchemaFieldDto
	Minimum float64 `json:"minimum,omitempty"`
	Maximum float64 `json:"maximum,omitempty"`
}

// GeolocationFieldDto DTO.
type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	AllowedTypes []string `json:"allowedTypes,omitempty"`
}

// TagsFieldDto DTO.
type TagsFieldDto struct {
	JsonSchemaFieldDto
}

// FileFieldDto DTO.
type FileFieldDto struct {
	JsonSchemaFieldDto
	Storages []string `json:"storages,omitempty"`
}

// TaxonomySelectionFieldDto DTO.
type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	TaxonomyId string `json:"taxonomyId,omitempty"`
	Multiple   bool   `json:"multiple,omitempty"`
}

// CollectionSelectionFieldDto DTO.
type CollectionSelectionFieldDto struct {
	JsonSchemaFieldDto
	CollectionId string `json:"collectionId,omitempty"`
	DisplayField string `json:"displayField,omitempty"`
	Multiple     bool   `json:"multiple,omitempty"`
}

// UserSelectionFieldDto DTO.
type UserSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple bool `json:"multiple,omitempty"`
}

// RoleSelectionFieldDto DTO.
type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple bool `json:"multiple,omitempty"`
}

// EnumSelectionFieldDto DTO.
type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	Values   []string `json:"values,omitempty"`
	Multiple bool     `json:"multiple,omitempty"`
}

// EchoResponse DTO.
type EchoResponse struct {
	ContainerName                string           `json:"containerName,omitempty"`
	Ip                           string           `json:"ip,omitempty"`
	Release                      CodeMashRelease  `json:"release,omitempty"`
	Runtime                      CodeMashRuntime  `json:"runtime,omitempty"`
	ManagedServiceHubUrl         string           `json:"managedServiceHubUrl,omitempty"`
	ManagedServiceApiUrl         string           `json:"managedServiceApiUrl,omitempty"`
	HubUrl                       string           `json:"hubUrl,omitempty"`
	ApiUrl                       string           `json:"apiUrl,omitempty"`
	ApiVersion                   string           `json:"apiVersion,omitempty"`
	HubVersion                   string           `json:"hubVersion,omitempty"`
	MjmlUrl                      string           `json:"mjmlUrl,omitempty"`
	AdminUrlTemplate             string           `json:"adminUrlTemplate,omitempty"`
	License                      *EchoLicenseDto  `json:"license,omitempty"`
	AskForEnterpriseLicenseEmail string           `json:"askForEnterpriseLicenseEmail,omitempty"`
	EmailServiceConfigured       bool             `json:"emailServiceConfigured,omitempty"`
	RootBootstrapPasswordSource  string           `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                      []*EchoRegionDto `json:"regions,omitempty"`
	IsProductionInstallation     bool             `json:"isProductionInstallation,omitempty"`
	LicensingMode                string           `json:"licensingMode,omitempty"`
	GraceDaysLeft                float64          `json:"graceDaysLeft,omitempty"`
	InstallationDomain           string           `json:"installationDomain,omitempty"`
	LicensingDocsUrl             string           `json:"licensingDocsUrl,omitempty"`
}

// PublicProjectConfigDto DTO.
type PublicProjectConfigDto struct {
	DisplayName        string          `json:"displayName,omitempty"`
	AdminPortalEnabled bool            `json:"adminPortalEnabled,omitempty"`
	Branding           *PublicBrandDto `json:"branding,omitempty"`
	Auth               *PublicAuthDto  `json:"auth,omitempty"`
}

// PublicLegalDocumentDto DTO.
type PublicLegalDocumentDto struct {
	Kind      string `json:"kind,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	Available bool   `json:"available,omitempty"`
}

// GetAccountProfileResponse DTO.
type GetAccountProfileResponse struct {
	ResponseBase
	Item *AccountOwnerDto `json:"item,omitempty"`
}

// EmptyResponse DTO.
type EmptyResponse struct {
	ResponseBase
}

// GetAccountStatusResponse DTO.
type GetAccountStatusResponse struct {
	ResponseBase
	Item *AccountStatusDto `json:"item,omitempty"`
}

// IdResponse DTO.
type IdResponse struct {
	ResponseBase
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

// CreateStripeCheckoutSessionResponse DTO.
type CreateStripeCheckoutSessionResponse struct {
	IdResponse
}

// GetStripeBillingPortalUrlResponse DTO.
type GetStripeBillingPortalUrlResponse struct {
	IdResponse
}

// CreateTeamMemberFromInvitationResponse DTO.
type CreateTeamMemberFromInvitationResponse struct {
	IdResponse
	Token string `json:"token,omitempty"`
}

// GetAccountUsageBillingResponse DTO.
type GetAccountUsageBillingResponse struct {
	ResponseBase
	Item *UsageBillingDto `json:"item,omitempty"`
}

// PromoteEnvironmentResponse DTO.
type PromoteEnvironmentResponse struct {
	ResponseBase
	Item *PromotionResultDto `json:"item,omitempty"`
}

// GetProjectEnvironmentsResponse DTO.
type GetProjectEnvironmentsResponse struct {
	ResponseBase
	Item *ProjectEnvironmentsDto `json:"item,omitempty"`
}

// GetProjectResponse DTO.
type GetProjectResponse struct {
	ResponseBase
	Item *ProjectDto `json:"item,omitempty"`
}

// GetProjectsResponse DTO.
type GetProjectsResponse struct {
	ResponseBase
	List []*ProjectListItemDto `json:"list,omitempty"`
}

// GetAccountRegionsResponse DTO.
type GetAccountRegionsResponse struct {
	ResponseBase
	Items []*ProjectRegionDto `json:"items,omitempty"`
}

// WaitForProjectActiveResponse DTO.
type WaitForProjectActiveResponse struct {
	ResponseBase
	Status        string  `json:"status,omitempty"`
	IsActive      bool    `json:"isActive,omitempty"`
	WaitedSeconds float64 `json:"waitedSeconds,omitempty"`
	Message       string  `json:"message,omitempty"`
}

// GetProjectTokensResponse DTO.
type GetProjectTokensResponse struct {
	ResponseBase
	Tokens []*TokenMappingDto `json:"tokens,omitempty"`
}

// AdminPortalStructureDto DTO.
type AdminPortalStructureDto struct {
	ProjectId          string                  `json:"projectId,omitempty"`
	AdminPortalEnabled bool                    `json:"adminPortalEnabled,omitempty"`
	DisplayName        string                  `json:"displayName,omitempty"`
	Modules            []*AdminPortalModuleDto `json:"modules,omitempty"`
}

// CreateAccountResponse DTO.
type CreateAccountResponse struct {
	IdResponse
	Token string `json:"token,omitempty"`
}

// GetAccountCollaboratorsResponse DTO.
type GetAccountCollaboratorsResponse struct {
	ResponseBase
	List PaginatedResponse[*AuthDto] `json:"list,omitempty"`
}

// GetAccountPasswordPolicyResponse DTO.
type GetAccountPasswordPolicyResponse struct {
	ResponseBase
	Policy *AccountPasswordPolicyDto `json:"policy,omitempty"`
}

// GetAccountTeamPoliciesResponse DTO.
type GetAccountTeamPoliciesResponse struct {
	ResponseBase
	Policies []*PolicyItemDto `json:"policies,omitempty"`
}

// GetAccountTeamRolesResponse DTO.
type GetAccountTeamRolesResponse struct {
	ResponseBase
	Roles []*AccountTeamRoleDto `json:"roles,omitempty"`
}

// AccountPasskeyOkResponse DTO.
type AccountPasskeyOkResponse struct {
	ResponseBase
}

// AccountPasskeyVerificationTokenResponse DTO.
type AccountPasskeyVerificationTokenResponse struct {
	ResponseBase
	VerificationToken string `json:"verificationToken,omitempty"`
}

// AccountPasskeyCeremonyOptionsResponse DTO.
type AccountPasskeyCeremonyOptionsResponse struct {
	ResponseBase
	CeremonyId  string `json:"ceremonyId,omitempty"`
	OptionsJson string `json:"optionsJson,omitempty"`
}

// AccountPasskeyAuthTokensResponse DTO.
type AccountPasskeyAuthTokensResponse struct {
	ResponseBase
	AccessToken      string   `json:"accessToken,omitempty"`
	RefreshToken     string   `json:"refreshToken,omitempty"`
	ExpiresInSeconds float64  `json:"expiresInSeconds,omitempty"`
	RecoveryCodes    []string `json:"recoveryCodes,omitempty"`
}

// AccountPasskeyListResponse DTO.
type AccountPasskeyListResponse struct {
	ResponseBase
	Passkeys []*AccountPasskeyListItemDto `json:"passkeys,omitempty"`
}

// AccountPasskeyEnrollmentResponse DTO.
type AccountPasskeyEnrollmentResponse struct {
	ResponseBase
	RecoveryCodes []string `json:"recoveryCodes,omitempty"`
}

// GetLicenseDomainDnsStatusResponse DTO.
type GetLicenseDomainDnsStatusResponse struct {
	ResponseBase
	Status *LicenseDomainDnsStatusDto `json:"status,omitempty"`
}

// StartLicenseDomainVerificationResponse DTO.
type StartLicenseDomainVerificationResponse struct {
	ResponseBase
	Challenge *LicenseDomainVerificationChallengeDto `json:"challenge,omitempty"`
}

// GetLicenseDomainVerificationStatusResponse DTO.
type GetLicenseDomainVerificationStatusResponse struct {
	ResponseBase
	Status *LicenseDomainVerificationStatusDto `json:"status,omitempty"`
}

// GetLicensesResponse DTO.
type GetLicensesResponse struct {
	ResponseBase
	List []*LicenseDto `json:"list,omitempty"`
}

// PostLicenseHeartbeatResponse DTO.
type PostLicenseHeartbeatResponse struct {
	ResponseBase
	Verdict *LicenseHeartbeatVerdictDto `json:"verdict,omitempty"`
}

// GetInstallationLicenseStatusResponse DTO.
type GetInstallationLicenseStatusResponse struct {
	ResponseBase
	Status *InstallationLicenseStatusDto `json:"status,omitempty"`
}

// IssueServiceUserApiKeyResponse DTO.
type IssueServiceUserApiKeyResponse struct {
	Id   float64 `json:"id,omitempty"`
	Name string  `json:"name,omitempty"`
	Key  string  `json:"key,omitempty"`
}

// ListServiceUserApiKeysResponse DTO.
type ListServiceUserApiKeysResponse struct {
	Keys []*ServiceUserApiKeyDto `json:"keys,omitempty"`
}

// GetMembershipTriggerResponse DTO.
type GetMembershipTriggerResponse struct {
	GetTriggerResponse
	Trigger *MembershipTriggerDto `json:"trigger,omitempty"`
}

// GetMembershipTriggersResponse DTO.
type GetMembershipTriggersResponse struct {
	GetTriggersResponse
	List PaginatedResponse[*MembershipTriggerProjectionList] `json:"list,omitempty"`
}

// GetRoleResponse DTO.
type GetRoleResponse struct {
	ResponseBase
	Role *RoleItemDto `json:"role,omitempty"`
}

// GetRolesResponse DTO.
type GetRolesResponse struct {
	ResponseBase
	Roles []*RoleListProjectionDto `json:"roles,omitempty"`
}

// GetPolicyResponse DTO.
type GetPolicyResponse struct {
	ResponseBase
	Policy *PolicyItemDto `json:"policy,omitempty"`
}

// GetPoliciesResponse DTO.
type GetPoliciesResponse struct {
	ResponseBase
	Policies []*PolicyItemDto `json:"policies,omitempty"`
}

// GetPasskeySettingsResponse DTO.
type GetPasskeySettingsResponse struct {
	ResponseBase
	Result *PasskeySettingsDto `json:"result,omitempty"`
}

// GetMembershipIntegrationResponse DTO.
type GetMembershipIntegrationResponse struct {
	ResponseBase
	Item *MembershipIntegrationDto `json:"item,omitempty"`
}

// GetMembershipIntegrationsResponse DTO.
type GetMembershipIntegrationsResponse struct {
	ResponseBase
	List PaginatedResponse[*MembershipIntegrationListProjection] `json:"list,omitempty"`
}

// GetAuthorizationSettingsResponse DTO.
type GetAuthorizationSettingsResponse struct {
	ResponseBase
	Result *MembershipAuthorizationViewDto `json:"result,omitempty"`
}

// UpdatePasswordComplexityResponse DTO.
type UpdatePasswordComplexityResponse struct {
	ResponseBase
	Result bool `json:"result,omitempty"`
}

// GetAuthenticationSettingsResponse DTO.
type GetAuthenticationSettingsResponse struct {
	ResponseBase
	Result *MembershipAuthenticationViewDto `json:"result,omitempty"`
}

// GetSchemaTriggerResponse DTO.
type GetSchemaTriggerResponse struct {
	GetTriggerResponse
	Trigger *SchemaTriggerDto `json:"trigger,omitempty"`
}

// GetSchemaTriggersResponse DTO.
type GetSchemaTriggersResponse struct {
	GetTriggersResponse
	List PaginatedResponse[*SchemaTriggerProjectionList] `json:"list,omitempty"`
}

// GetDatabaseTaxonomyResponse DTO.
type GetDatabaseTaxonomyResponse struct {
	ResponseBase
	Item *TaxonomyDto `json:"item,omitempty"`
}

// GetDatabaseTaxonomiesResponse DTO.
type GetDatabaseTaxonomiesResponse struct {
	ResponseBase
	List PaginatedResponse[*TaxonomyListProjection] `json:"list,omitempty"`
}

// GetDatabaseTaxonomyTreeResponse DTO.
type GetDatabaseTaxonomyTreeResponse struct {
	ResponseBase
	Tree []*TaxonomyTreeDto `json:"tree,omitempty"`
}

// GetDatabaseTaxonomyTermResponse DTO.
type GetDatabaseTaxonomyTermResponse struct {
	ResponseBase
	Item *TermDto `json:"item,omitempty"`
}

// GetDatabaseMergedTermTreeResponse DTO.
type GetDatabaseMergedTermTreeResponse struct {
	ResponseBase
	Tree []*TermTreeDto `json:"tree,omitempty"`
}

// GetDatabaseTaxonomyTermTreeResponse DTO.
type GetDatabaseTaxonomyTermTreeResponse struct {
	ResponseBase
	Tree []*TermTreeDto `json:"tree,omitempty"`
}

// ApplyDatabaseSchemaBundleResponse DTO.
type ApplyDatabaseSchemaBundleResponse struct {
	ResponseBase
	Tier        string                  `json:"tier,omitempty"`
	Taxonomies  []*AppliedTaxonomyDto   `json:"taxonomies,omitempty"`
	Collections []*AppliedCollectionDto `json:"collections,omitempty"`
	Decisions   []string                `json:"decisions,omitempty"`
	Errors      []string                `json:"errors,omitempty"`
}

// GetDatabaseSchemaResponse DTO.
type GetDatabaseSchemaResponse struct {
	ResponseBase
	Item *SchemaDto `json:"item,omitempty"`
}

// GetDatabaseSchemasResponse DTO.
type GetDatabaseSchemasResponse struct {
	ResponseBase
	List PaginatedResponse[*SchemaListProjection] `json:"list,omitempty"`
}

// GetDatabaseSchemaDraftResponse DTO.
type GetDatabaseSchemaDraftResponse struct {
	ResponseBase
	Item *SchemaDraftDto `json:"item,omitempty"`
}

// GetDatabaseSchemaListSettingsResponse DTO.
type GetDatabaseSchemaListSettingsResponse struct {
	ResponseBase
	Settings *SchemaListSettingsDto `json:"settings,omitempty"`
}

// GetDatabaseSchemaVersionDiffResponse DTO.
type GetDatabaseSchemaVersionDiffResponse struct {
	ResponseBase
	Item *SchemaDiffDto `json:"item,omitempty"`
}

// GetDatabaseSchemaVersionsResponse DTO.
type GetDatabaseSchemaVersionsResponse struct {
	ResponseBase
	Items []*SchemaVersionSummaryDto `json:"items,omitempty"`
}

// AggregateRecordsResponse DTO.
type AggregateRecordsResponse struct {
	ResponseBase
	Result []map[string]any `json:"result,omitempty"`
}

// CountRecordsResponse DTO.
type CountRecordsResponse struct {
	ResponseBase
	Count float64 `json:"count,omitempty"`
}

// DistinctRecordValuesResponse DTO.
type DistinctRecordValuesResponse struct {
	ResponseBase
	Values []map[string]any `json:"values,omitempty"`
}

// ExecuteRecordsAggregateResponse DTO.
type ExecuteRecordsAggregateResponse struct {
	ResponseBase
	Result []map[string]any `json:"result,omitempty"`
}

// FindRecordsResponse DTO.
type FindRecordsResponse struct {
	ResponseBase
	List PaginatedResponse[map[string]any] `json:"list,omitempty"`
}

// FindOneRecordResponse DTO.
type FindOneRecordResponse struct {
	ResponseBase
	Result map[string]any `json:"result,omitempty"`
}

// GetCollectionIndexesResponse DTO.
type GetCollectionIndexesResponse struct {
	ResponseBase
	Indexes []*CollectionIndexDto `json:"indexes,omitempty"`
}

// SeedCollectionRecordsResponse DTO.
type SeedCollectionRecordsResponse struct {
	ResponseBase
	Result *SeedCollectionRecordsResultDto `json:"result,omitempty"`
}

// GetDatabaseIntegrationResponse DTO.
type GetDatabaseIntegrationResponse struct {
	ResponseBase
	Item *DatabaseIntegrationDto `json:"item,omitempty"`
}

// GetDatabaseIntegrationsResponse DTO.
type GetDatabaseIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                                `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*DatabaseIntegrationListProjection] `json:"list,omitempty"`
}

// GetAllowedFlexTiersResponse DTO.
type GetAllowedFlexTiersResponse struct {
	ResponseBase
	Tiers []*FlexTierDto `json:"tiers,omitempty"`
}

// RevealManagedFlexConnectionStringResponse DTO.
type RevealManagedFlexConnectionStringResponse struct {
	ResponseBase
	ConnectionString string `json:"connectionString,omitempty"`
}

// TestDatabaseIntegrationResponse DTO.
type TestDatabaseIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetCollectionImportResponse DTO.
type GetCollectionImportResponse struct {
	ResponseBase
	Result *CollectionImportDto `json:"result,omitempty"`
}

// GetCollectionImportsResponse DTO.
type GetCollectionImportsResponse struct {
	ResponseBase
	Result PaginatedResponse[*CollectionImportDto] `json:"result,omitempty"`
}

// RequestImportUploadUrlResponse DTO.
type RequestImportUploadUrlResponse struct {
	ResponseBase
	Result *ImportUploadTargetDto `json:"result,omitempty"`
}

// AnalyzeImportFileResponse DTO.
type AnalyzeImportFileResponse struct {
	ResponseBase
	Result *ImportFileAnalysisDto `json:"result,omitempty"`
}

// GetDatabaseAggregateResponse DTO.
type GetDatabaseAggregateResponse struct {
	ResponseBase
	Item *MongoDbAggregateDto `json:"item,omitempty"`
}

// GetDatabaseAggregatesResponse DTO.
type GetDatabaseAggregatesResponse struct {
	ResponseBase
	List PaginatedResponse[*MongoDbAggregateListProjection] `json:"list,omitempty"`
}

// TestDatabaseAggregateResponse DTO.
type TestDatabaseAggregateResponse struct {
	ResponseBase
	Result []map[string]any `json:"result,omitempty"`
}

// GetFilesTriggerResponse DTO.
type GetFilesTriggerResponse struct {
	GetTriggerResponse
	Trigger *FilesTriggerDto `json:"trigger,omitempty"`
}

// GetFilesTriggersResponse DTO.
type GetFilesTriggersResponse struct {
	GetTriggersResponse
	List PaginatedResponse[*FilesTriggerProjectionList] `json:"list,omitempty"`
}

// GetFilesIntegrationResponse DTO.
type GetFilesIntegrationResponse struct {
	ResponseBase
	Item *FilesIntegrationDto `json:"item,omitempty"`
}

// GetFilesIntegrationsResponse DTO.
type GetFilesIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                             `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*FilesIntegrationListProjection] `json:"list,omitempty"`
	AvailableProviders   []FileProvider                                     `json:"availableProviders,omitempty"`
}

// TestFilesIntegrationResponse DTO.
type TestFilesIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetFileResponse DTO.
type GetFileResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  bool                `json:"isPublic,omitempty"`
	PublicUrl string              `json:"publicUrl,omitempty"`
}

// GetFolderFilesResponse DTO.
type GetFolderFilesResponse struct {
	ResponseBase
	List          PaginatedResponse[*FileResourceRefDto] `json:"list,omitempty"`
	Folders       []string                               `json:"folders,omitempty"`
	PublicFolders []*PublicFolderDto                     `json:"publicFolders,omitempty"`
}

// GetNotificationModuleDisableDependenciesResponse DTO.
type GetNotificationModuleDisableDependenciesResponse struct {
	ResponseBase
	Dependencies *NotificationModuleDisableDependenciesDto `json:"dependencies,omitempty"`
}

// TestEmailValidationIntegrationResponse DTO.
type TestEmailValidationIntegrationResponse struct {
	ResponseBase
	Items []*TestEmailValidationItemDto `json:"items,omitempty"`
}

// GetEmailTemplateResponse DTO.
type GetEmailTemplateResponse struct {
	ResponseBase
	Item *EmailTemplateDto `json:"item,omitempty"`
}

// GetEmailTemplatesResponse DTO.
type GetEmailTemplatesResponse struct {
	ResponseBase
	List PaginatedResponse[*EmailTemplateListProjection] `json:"list,omitempty"`
}

// GetHtmlFromMjmlResponse DTO.
type GetHtmlFromMjmlResponse struct {
	ResponseBase
	Variables            []string              `json:"variables,omitempty"`
	HtmlFromMjmlResponse *HtmlFromMjmlResponse `json:"htmlFromMjmlResponse,omitempty"`
}

// GetSystemEmailTemplateResponse DTO.
type GetSystemEmailTemplateResponse struct {
	ResponseBase
	Item *SystemEmailTemplateDto `json:"item,omitempty"`
}

// GetSystemEmailTemplatesResponse DTO.
type GetSystemEmailTemplatesResponse struct {
	ResponseBase
	List PaginatedResponse[*SystemEmailTemplateListProjection] `json:"list,omitempty"`
}

// GetEmailTemplateAvailableTokensResponse DTO.
type GetEmailTemplateAvailableTokensResponse struct {
	ResponseBase
	Tokens map[string][]string `json:"tokens,omitempty"`
}

// GetEmailSignatureResponse DTO.
type GetEmailSignatureResponse struct {
	ResponseBase
	Item *EmailSignatureDto `json:"item,omitempty"`
}

// GetEmailSignaturesResponse DTO.
type GetEmailSignaturesResponse struct {
	ResponseBase
	List PaginatedResponse[*ListItemWithTranslationsProjection] `json:"list,omitempty"`
}

// GetEmailSettingsResponse DTO.
type GetEmailSettingsResponse struct {
	ResponseBase
	Settings   *EmailSettings        `json:"settings,omitempty"`
	SystemTags []*GroupDefinitionDto `json:"systemTags,omitempty"`
}

// CheckEmailIntegrationDomainHealthResponse DTO.
type CheckEmailIntegrationDomainHealthResponse struct {
	ResponseBase
	Domain string                       `json:"domain,omitempty"`
	Items  []*DomainHealthRecordItemDto `json:"items,omitempty"`
}

// GetEmailIntegrationResponse DTO.
type GetEmailIntegrationResponse struct {
	ResponseBase
	Item *EmailIntegrationDto `json:"item,omitempty"`
}

// GetEmailIntegrationsResponse DTO.
type GetEmailIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                             `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*EmailIntegrationListProjection] `json:"list,omitempty"`
}

// TestEmailIntegrationResponse DTO.
type TestEmailIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetEmailFooterResponse DTO.
type GetEmailFooterResponse struct {
	ResponseBase
	Item *EmailFooterDto `json:"item,omitempty"`
}

// GetEmailFootersResponse DTO.
type GetEmailFootersResponse struct {
	ResponseBase
	List PaginatedResponse[*ListItemWithTranslationsProjection] `json:"list,omitempty"`
}

// GetEmailCampaignResponse DTO.
type GetEmailCampaignResponse struct {
	ResponseBase
	Item *EmailCampaignDto `json:"item,omitempty"`
}

// GetEmailCampaignsResponse DTO.
type GetEmailCampaignsResponse struct {
	ResponseBase
	List PaginatedResponse[*EmailCampaignListProjection] `json:"list,omitempty"`
}

// GetEmailCampaignBatchesResponse DTO.
type GetEmailCampaignBatchesResponse struct {
	ResponseBase
	List PaginatedResponse[*EmailCampaignBatchDto] `json:"list,omitempty"`
}

// GetEmailCampaignBatchNotificationResponse DTO.
type GetEmailCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *EmailCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

// GetEmailCampaignBatchNotificationsResponse DTO.
type GetEmailCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []*BatchStatusChangeEntryDto                          `json:"batchStatusHistory,omitempty"`
	List               PaginatedResponse[*EmailCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetEmailCampaignStatisticsResponse DTO.
type GetEmailCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

// PreviewEmailNotificationResponse DTO.
type PreviewEmailNotificationResponse struct {
	ResponseBase
	Subject string `json:"subject,omitempty"`
	Body    string `json:"body,omitempty"`
}

// GetEmailCampaignMessageResponse DTO.
type GetEmailCampaignMessageResponse struct {
	ResponseBase
	EmailMessageEntity *EmailCampaignBatchNotificationDto `json:"emailMessageEntity,omitempty"`
}

// GetEmailCampaignMessagesResponse DTO.
type GetEmailCampaignMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*EmailCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetSmsTemplateResponse DTO.
type GetSmsTemplateResponse struct {
	ResponseBase
	Item *SmsTemplateDto `json:"item,omitempty"`
}

// GetSmsTemplatesResponse DTO.
type GetSmsTemplatesResponse struct {
	ResponseBase
	List PaginatedResponse[*SmsTemplateListProjection] `json:"list,omitempty"`
}

// GetSmsMessageContentTokensResponse DTO.
type GetSmsMessageContentTokensResponse struct {
	ResponseBase
	Tokens map[string][]string `json:"tokens,omitempty"`
}

// RenderSmsTextResponse DTO.
type RenderSmsTextResponse struct {
	ResponseBase
	Variables []string `json:"variables,omitempty"`
	Text      string   `json:"text,omitempty"`
}

// GetSmsSettingsResponse DTO.
type GetSmsSettingsResponse struct {
	ResponseBase
	Settings *SmsSettings `json:"settings,omitempty"`
}

// GetSmsIntegrationResponse DTO.
type GetSmsIntegrationResponse struct {
	ResponseBase
	Item *SmsIntegrationDto `json:"item,omitempty"`
}

// GetSmsIntegrationsResponse DTO.
type GetSmsIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                           `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*SmsIntegrationListProjection] `json:"list,omitempty"`
}

// TestSmsIntegrationResponse DTO.
type TestSmsIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetSmsCampaignResponse DTO.
type GetSmsCampaignResponse struct {
	ResponseBase
	SmsCampaign *SmsCampaignDto `json:"smsCampaign,omitempty"`
}

// GetSmsCampaignsResponse DTO.
type GetSmsCampaignsResponse struct {
	ResponseBase
	List PaginatedResponse[*SmsCampaignDto] `json:"list,omitempty"`
}

// GetSmsCampaignBatchesResponse DTO.
type GetSmsCampaignBatchesResponse struct {
	ResponseBase
	List PaginatedResponse[*SmsCampaignBatchDto] `json:"list,omitempty"`
}

// GetSmsCampaignBatchNotificationResponse DTO.
type GetSmsCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *SmsCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

// GetSmsCampaignBatchNotificationsResponse DTO.
type GetSmsCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []*BatchStatusChangeEntryDto                        `json:"batchStatusHistory,omitempty"`
	List               PaginatedResponse[*SmsCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetSmsCampaignStatisticsResponse DTO.
type GetSmsCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

// PreviewSmsNotificationResponse DTO.
type PreviewSmsNotificationResponse struct {
	ResponseBase
	Body string `json:"body,omitempty"`
}

// GetSmsCampaignMessageResponse DTO.
type GetSmsCampaignMessageResponse struct {
	ResponseBase
	SmsMessageEntity *SmsCampaignBatchNotificationDto `json:"smsMessageEntity,omitempty"`
}

// GetSmsCampaignMessagesResponse DTO.
type GetSmsCampaignMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*SmsCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetMarketplaceListingResponse DTO.
type GetMarketplaceListingResponse struct {
	ResponseBase
	Listing *MarketplaceListingDto `json:"listing,omitempty"`
}

// GetMarketplaceTokensResponse DTO.
type GetMarketplaceTokensResponse struct {
	ResponseBase
	Tokens []string `json:"tokens,omitempty"`
}

// GetMarketplaceListingsResponse DTO.
type GetMarketplaceListingsResponse struct {
	ResponseBase
	List PaginatedResponse[*MarketplaceListingProjection] `json:"list,omitempty"`
}

// GetMarketplaceIntegrationResponse DTO.
type GetMarketplaceIntegrationResponse struct {
	ResponseBase
	Integration *MarketplaceIntegrationDto `json:"integration,omitempty"`
}

// GetMarketplaceIntegrationsResponse DTO.
type GetMarketplaceIntegrationsResponse struct {
	ResponseBase
	List PaginatedResponse[*MarketplaceIntegrationListProjection] `json:"list,omitempty"`
}

// EmptyMarketplaceSecretsResponse DTO.
type EmptyMarketplaceSecretsResponse struct {
	ResponseBase
}

// RevealMarketplaceIntegrationSecretsResponse DTO.
type RevealMarketplaceIntegrationSecretsResponse struct {
	ResponseBase
	Secrets map[string]string `json:"secrets,omitempty"`
}

// TestMarketplaceIntegrationResponse DTO.
type TestMarketplaceIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// SetMarketplaceIntegrationTokenMappingsResponse DTO.
type SetMarketplaceIntegrationTokenMappingsResponse struct {
	ResponseBase
}

// GetMarketplaceFunctionResponse DTO.
type GetMarketplaceFunctionResponse struct {
	ResponseBase
	Function *MarketplaceFunctionDto `json:"function,omitempty"`
}

// GetMarketplaceFunctionsResponse DTO.
type GetMarketplaceFunctionsResponse struct {
	ResponseBase
	List PaginatedResponse[*MarketplaceFunctionProjection] `json:"list,omitempty"`
}

// GetMarketplaceFunctionCatalogResponse DTO.
type GetMarketplaceFunctionCatalogResponse struct {
	ResponseBase
	Functions []*MarketplaceFunctionDefinitionDto `json:"functions,omitempty"`
}

// InvokeMarketplaceFunctionResponse DTO.
type InvokeMarketplaceFunctionResponse struct {
	ResponseBase
	IsSuccess       bool           `json:"isSuccess,omitempty"`
	Output          map[string]any `json:"output,omitempty"`
	VendorRequestId string         `json:"vendorRequestId,omitempty"`
}

// GetCodeIntegrationResponse DTO.
type GetCodeIntegrationResponse struct {
	ResponseBase
	Item *CodeIntegrationDto `json:"item,omitempty"`
}

// GetCodeIntegrationsResponse DTO.
type GetCodeIntegrationsResponse struct {
	ResponseBase
	List PaginatedResponse[*CodeIntegrationListProjection] `json:"list,omitempty"`
}

// TestCodeIntegrationResponse DTO.
type TestCodeIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetPushTemplateResponse DTO.
type GetPushTemplateResponse struct {
	ResponseBase
	Item *PushTemplateDto `json:"item,omitempty"`
}

// GetPushTemplatesResponse DTO.
type GetPushTemplatesResponse struct {
	ResponseBase
	List PaginatedResponse[*PushTemplateListProjection] `json:"list,omitempty"`
}

// GetPushMessageContentTokensResponse DTO.
type GetPushMessageContentTokensResponse struct {
	ResponseBase
	Tokens map[string][]string `json:"tokens,omitempty"`
}

// RenderPushResponse DTO.
type RenderPushResponse struct {
	ResponseBase
	Variables []string `json:"variables,omitempty"`
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Subtitle  string   `json:"subtitle,omitempty"`
}

// GetPushSettingsResponse DTO.
type GetPushSettingsResponse struct {
	ResponseBase
	Settings *PushSettings `json:"settings,omitempty"`
}

// GetPushIntegrationResponse DTO.
type GetPushIntegrationResponse struct {
	ResponseBase
	Item *PushIntegrationDto `json:"item,omitempty"`
}

// GetPushIntegrationsResponse DTO.
type GetPushIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                            `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*PushIntegrationListProjection] `json:"list,omitempty"`
}

// GetPushCampaignResponse DTO.
type GetPushCampaignResponse struct {
	ResponseBase
	Item *PushCampaignDto `json:"item,omitempty"`
}

// GetPushCampaignsResponse DTO.
type GetPushCampaignsResponse struct {
	ResponseBase
	List PaginatedResponse[*PushCampaignDto] `json:"list,omitempty"`
}

// GetPushCampaignBatchesResponse DTO.
type GetPushCampaignBatchesResponse struct {
	ResponseBase
	List PaginatedResponse[*PushCampaignBatchDto] `json:"list,omitempty"`
}

// GetPushCampaignBatchNotificationResponse DTO.
type GetPushCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *PushCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

// GetPushCampaignBatchNotificationsResponse DTO.
type GetPushCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []*BatchStatusChangeEntryDto                         `json:"batchStatusHistory,omitempty"`
	List               PaginatedResponse[*PushCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetPushCampaignStatisticsResponse DTO.
type GetPushCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

// PreviewPushNotificationResponse DTO.
type PreviewPushNotificationResponse struct {
	ResponseBase
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	Subtitle string `json:"subtitle,omitempty"`
}

// GetPushCampaignMessageResponse DTO.
type GetPushCampaignMessageResponse struct {
	ResponseBase
	PushMessageEntity *PushCampaignBatchNotificationDto `json:"pushMessageEntity,omitempty"`
}

// GetPushCampaignMessagesResponse DTO.
type GetPushCampaignMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*PushCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetPaymentsWebhookLogResponse DTO.
type GetPaymentsWebhookLogResponse struct {
	ResponseBase
	List []*PaymentsWebhookLogEntry `json:"list,omitempty"`
}

// GetPaymentsTriggerResponse DTO.
type GetPaymentsTriggerResponse struct {
	GetTriggerResponse
	Trigger *PaymentTriggerDto `json:"trigger,omitempty"`
}

// GetPaymentsTriggersResponse DTO.
type GetPaymentsTriggersResponse struct {
	GetTriggersResponse
	List PaginatedResponse[*PaymentTriggerProjectionList] `json:"list,omitempty"`
}

// GetPaymentsIntegrationResponse DTO.
type GetPaymentsIntegrationResponse struct {
	ResponseBase
	Item *PaymentsIntegrationDto `json:"item,omitempty"`
}

// GetPaymentsIntegrationsResponse DTO.
type GetPaymentsIntegrationsResponse struct {
	ResponseBase
	List PaginatedResponse[*PaymentsIntegrationListProjection] `json:"list,omitempty"`
}

// TestPaymentsIntegrationResponse DTO.
type TestPaymentsIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetLoggingIntegrationResponse DTO.
type GetLoggingIntegrationResponse struct {
	ResponseBase
	Item *LoggingIntegrationDto `json:"item,omitempty"`
}

// GetLoggingIntegrationsResponse DTO.
type GetLoggingIntegrationsResponse struct {
	ResponseBase
	List PaginatedResponse[*LoggingIntegrationListProjection] `json:"list,omitempty"`
}

// TestLoggingIntegrationResponse DTO.
type TestLoggingIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// CleanLogsResponse DTO.
type CleanLogsResponse struct {
	ResponseBase
}

// GetLogsByCorrelationIdResponse DTO.
type GetLogsByCorrelationIdResponse struct {
	ResponseBase
	Items []*TenantLogEntryDto `json:"items,omitempty"`
}

// GetLogsResponse DTO.
type GetLogsResponse struct {
	ResponseBase
	List PaginatedResponse[*TenantLogEntryDto] `json:"list,omitempty"`
}

// GetLogSettingsResponse DTO.
type GetLogSettingsResponse struct {
	ResponseBase
	SkipCloudDashboardLogs bool `json:"skipCloudDashboardLogs,omitempty"`
	SkipHttpBodyMeta       bool `json:"skipHttpBodyMeta,omitempty"`
	AiChatLoggingEnabled   bool `json:"aiChatLoggingEnabled,omitempty"`
	HasNorbixLogging       bool `json:"hasNorbixLogging,omitempty"`
}

// SaveLogSettingsResponse DTO.
type SaveLogSettingsResponse struct {
	ResponseBase
}

// GetAiToolsResponse DTO.
type GetAiToolsResponse struct {
	ResponseBase
	Tools []*AiToolManifestItem `json:"tools,omitempty"`
}

// InvokeAiToolResponse DTO.
type InvokeAiToolResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// AskChatResponse DTO.
type AskChatResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// UploadChatAttachmentResponse DTO.
type UploadChatAttachmentResponse struct {
	ResponseBase
	Id        string `json:"id,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
}

// ChatAvailabilityResponse DTO.
type ChatAvailabilityResponse struct {
	ResponseBase
	Available bool               `json:"available,omitempty"`
	Reason    string             `json:"reason,omitempty"`
	Profiles  []string           `json:"profiles,omitempty"`
	Models    []*ChatModelOption `json:"models,omitempty"`
}

// GetChatMemoryResponse DTO.
type GetChatMemoryResponse struct {
	ResponseBase
	Notes []*ChatMemoryNote `json:"notes,omitempty"`
}

// GetChatSessionsResponse DTO.
type GetChatSessionsResponse struct {
	ResponseBase
	Sessions []*ChatSessionListItem `json:"sessions,omitempty"`
}

// GetChatSessionEntriesResponse DTO.
type GetChatSessionEntriesResponse struct {
	ResponseBase
	SessionId         string                `json:"sessionId,omitempty"`
	Profile           string                `json:"profile,omitempty"`
	ProjectId         string                `json:"projectId,omitempty"`
	Env               string                `json:"env,omitempty"`
	Entries           []*AiChatEntryWireDto `json:"entries,omitempty"`
	LastSeq           float64               `json:"lastSeq,omitempty"`
	ActiveWorkItemIds []string              `json:"activeWorkItemIds,omitempty"`
}

// ChatTurnResponse DTO.
type ChatTurnResponse struct {
	ResponseBase
	SessionId   string                `json:"sessionId,omitempty"`
	Reply       string                `json:"reply,omitempty"`
	ScreenPatch *ChatScreenContextDto `json:"screenPatch,omitempty"`
	ToolTrace   []string              `json:"toolTrace,omitempty"`
}

// GetProjectBriefResponse DTO.
type GetProjectBriefResponse struct {
	ResponseBase
	ProjectId string                       `json:"projectId,omitempty"`
	Snapshot  *ProjectBriefSnapshotWireDto `json:"snapshot,omitempty"`
	Events    []*ProjectBriefEventWireDto  `json:"events,omitempty"`
	LastSeq   float64                      `json:"lastSeq,omitempty"`
}

// GetWorkItemsResponse DTO.
type GetWorkItemsResponse struct {
	ResponseBase
	ProjectId string             `json:"projectId,omitempty"`
	WorkItems []*WorkItemWireDto `json:"workItems,omitempty"`
}

// GetWorkItemResponse DTO.
type GetWorkItemResponse struct {
	ResponseBase
	WorkItem *WorkItemWireDto      `json:"workItem,omitempty"`
	Plans    []*AiChatEntryWireDto `json:"plans,omitempty"`
	Steps    []*AiChatEntryWireDto `json:"steps,omitempty"`
}

// ExportWorkItemResponse DTO.
type ExportWorkItemResponse struct {
	ResponseBase
	WorkItemId string `json:"workItemId,omitempty"`
	Markdown   string `json:"markdown,omitempty"`
}

// GetLlmIntegrationResponse DTO.
type GetLlmIntegrationResponse struct {
	ResponseBase
	Item *LlmIntegrationDto `json:"item,omitempty"`
}

// GetLlmIntegrationsResponse DTO.
type GetLlmIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                           `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*LlmIntegrationListProjection] `json:"list,omitempty"`
}

// TestLlmIntegrationResponse DTO.
type TestLlmIntegrationResponse struct {
	ResponseBase
	Items []*IntegrationTestResultItemDto `json:"items,omitempty"`
}

// GetMcpIntegrationResponse DTO.
type GetMcpIntegrationResponse struct {
	ResponseBase
	Item *McpIntegrationDto `json:"item,omitempty"`
}

// GetMcpIntegrationsResponse DTO.
type GetMcpIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId string                                           `json:"defaultIntegrationId,omitempty"`
	List                 PaginatedResponse[*McpIntegrationListProjection] `json:"list,omitempty"`
}

// GetWebhookIntegrationResponse DTO.
type GetWebhookIntegrationResponse struct {
	ResponseBase
	Item *WebhookIntegrationDto `json:"item,omitempty"`
}

// RevealWebhookIntegrationSecretResponse DTO.
type RevealWebhookIntegrationSecretResponse struct {
	ResponseBase
	SigningSecret string `json:"signingSecret,omitempty"`
}

// RotateWebhookIntegrationSecretResponse DTO.
type RotateWebhookIntegrationSecretResponse struct {
	ResponseBase
	SigningSecret string `json:"signingSecret,omitempty"`
}

// HttpResult DTO.
type HttpResult struct {
	ResponseText          string              `json:"responseText,omitempty"`
	ResponseStream        string              `json:"responseStream,omitempty"`
	FileInfo              any                 `json:"fileInfo,omitempty"`
	VirtualFile           *IVirtualFile       `json:"virtualFile,omitempty"`
	ContentType           string              `json:"contentType,omitempty"`
	Headers               map[string]string   `json:"headers,omitempty"`
	Cookies               []any               `json:"cookies,omitempty"`
	ETag                  string              `json:"eTag,omitempty"`
	Age                   string              `json:"age,omitempty"`
	MaxAge                string              `json:"maxAge,omitempty"`
	Expires               string              `json:"expires,omitempty"`
	LastModified          string              `json:"lastModified,omitempty"`
	CacheControl          CacheControl        `json:"cacheControl,omitempty"`
	ResultScope           any                 `json:"resultScope,omitempty"`
	AllowsPartialResponse bool                `json:"allowsPartialResponse,omitempty"`
	Options               map[string]string   `json:"options,omitempty"`
	Status                float64             `json:"status,omitempty"`
	StatusCode            any                 `json:"statusCode,omitempty"`
	StatusDescription     string              `json:"statusDescription,omitempty"`
	Response              map[string]any      `json:"response,omitempty"`
	ResponseFilter        *IContentTypeWriter `json:"responseFilter,omitempty"`
	RequestContext        *IRequest           `json:"requestContext,omitempty"`
	View                  string              `json:"view,omitempty"`
	Template              string              `json:"template,omitempty"`
	PaddingLength         float64             `json:"paddingLength,omitempty"`
	IsPartialRequest      bool                `json:"isPartialRequest,omitempty"`
}

// SaveWebhookDestinationResponse DTO.
type SaveWebhookDestinationResponse struct {
	ResponseBase
	DestinationId string `json:"destinationId,omitempty"`
}

// GetSchedulerTaskResponse DTO.
type GetSchedulerTaskResponse struct {
	ResponseBase
	Item *SchedulerTaskDto `json:"item,omitempty"`
}

// GetSchedulerTasksResponse DTO.
type GetSchedulerTasksResponse struct {
	ResponseBase
	List PaginatedResponse[*SchedulerTaskListProjection] `json:"list,omitempty"`
}

// ResolveResourcesResponse DTO.
type ResolveResourcesResponse struct {
	ResponseBase
	Resolved []*ResolvedResourceEntry `json:"resolved,omitempty"`
}

// GetContactResponse DTO.
type GetContactResponse struct {
	ResponseBase
	Item *UserDto `json:"item,omitempty"`
}

// GetAllContactsResponse DTO.
type GetAllContactsResponse struct {
	ResponseBase
	Items      []*UserDto `json:"items,omitempty"`
	NextCursor string     `json:"nextCursor,omitempty"`
}

// GetComplianceSettingsResponse DTO.
type GetComplianceSettingsResponse struct {
	ResponseBase
	Settings *ProjectComplianceDto `json:"settings,omitempty"`
}

// GetLegalHoldsResponse DTO.
type GetLegalHoldsResponse struct {
	ResponseBase
	Holds []*LegalHoldDto `json:"holds,omitempty"`
}

// GetDsarRequestsResponse DTO.
type GetDsarRequestsResponse struct {
	ResponseBase
	Requests []*DsarRequestDto `json:"requests,omitempty"`
}

// GetComplianceAuditLogResponse DTO.
type GetComplianceAuditLogResponse struct {
	ResponseBase
	Entries []*ComplianceAuditEntryDto `json:"entries,omitempty"`
}

// GetAccountComplianceResponse DTO.
type GetAccountComplianceResponse struct {
	ResponseBase
	Settings *AccountComplianceDto `json:"settings,omitempty"`
}

// GetSupportCaseResponse DTO.
type GetSupportCaseResponse struct {
	ResponseBase
	Result *SupportCaseDetailDto `json:"result,omitempty"`
}

// GetSupportCasesResponse DTO.
type GetSupportCasesResponse struct {
	ResponseBase
	List PaginatedResponse[*SupportCaseListProjection] `json:"list,omitempty"`
}

// GetDiagnosticPacksResponse DTO.
type GetDiagnosticPacksResponse struct {
	ResponseBase
	Packs []*DiagnosticPackDescriptorDto `json:"packs,omitempty"`
}

// RunDiagnosticPackResponse DTO.
type RunDiagnosticPackResponse struct {
	ResponseBase
	Result *DiagnosticPackRunResultDto `json:"result,omitempty"`
}

// GetDiagnosticEchoResponse DTO.
type GetDiagnosticEchoResponse struct {
	ResponseBase
	Result *DiagnosticEchoDto `json:"result,omitempty"`
}

// ReadDiagnosticEventsResponse DTO.
type ReadDiagnosticEventsResponse struct {
	ResponseBase
	Result *DiagnosticEventsPageDto `json:"result,omitempty"`
}

// QueryDiagnosticLogsResponse DTO.
type QueryDiagnosticLogsResponse struct {
	ResponseBase
	Result *DiagnosticLogsResponse `json:"result,omitempty"`
}

// InspectDiagnosticRedisResponse DTO.
type InspectDiagnosticRedisResponse struct {
	ResponseBase
	Result *DiagnosticRedisInspectDto `json:"result,omitempty"`
}

// RunDiagnosticHealthCheckResponse DTO.
type RunDiagnosticHealthCheckResponse struct {
	ResponseBase
	Result *DiagnosticHealthCheckDto `json:"result,omitempty"`
}

// AuthenticateResponse DTO.
type AuthenticateResponse struct {
	UserId             string            `json:"userId,omitempty"`
	SessionId          string            `json:"sessionId,omitempty"`
	UserName           string            `json:"userName,omitempty"`
	DisplayName        string            `json:"displayName,omitempty"`
	ReferrerUrl        string            `json:"referrerUrl,omitempty"`
	BearerToken        string            `json:"bearerToken,omitempty"`
	RefreshToken       string            `json:"refreshToken,omitempty"`
	RefreshTokenExpiry string            `json:"refreshTokenExpiry,omitempty"`
	ProfileUrl         string            `json:"profileUrl,omitempty"`
	Roles              []string          `json:"roles,omitempty"`
	Permissions        []string          `json:"permissions,omitempty"`
	AuthProvider       string            `json:"authProvider,omitempty"`
	ResponseStatus     *ResponseStatus   `json:"responseStatus,omitempty"`
	Meta               map[string]string `json:"meta,omitempty"`
}

// GetAccessTokenResponse DTO.
type GetAccessTokenResponse struct {
	AccessToken    string            `json:"accessToken,omitempty"`
	Meta           map[string]string `json:"meta,omitempty"`
	ResponseStatus *ResponseStatus   `json:"responseStatus,omitempty"`
}

// EnableCode DTO.
type EnableCode struct {
	CodeMashRequestBase
}

// DisableCode DTO.
type DisableCode struct {
	CodeMashRequestBase
}

// GetCodeIntegrations DTO.
type GetCodeIntegrations struct {
	CodeMashListPaginationRequestBase
}

// GetCodeIntegration DTO.
type GetCodeIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// SaveCodeIntegration DTO.
type SaveCodeIntegration struct {
	CodeMashRequestBase
	Integration *CodeIntegrationRequest `json:"integration,omitempty"`
}

// TestCodeIntegration DTO.
type TestCodeIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// ConfirmCodeIntegrationHumanDeliveryRequest DTO.
type ConfirmCodeIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// SetCodeIntegrationAsDefault DTO.
type SetCodeIntegrationAsDefault struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DeleteCodeIntegrationRequest DTO.
type DeleteCodeIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableCodeIntegrationRequest DTO.
type EnableCodeIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableCodeIntegrationRequest DTO.
type DisableCodeIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetMarketplaceListings DTO.
type GetMarketplaceListings struct {
	CodeMashListPaginationRequestBase
	Categories   []MarketplaceCategory  `json:"categories,omitempty"`
	Transports   []MarketplaceTransport `json:"transports,omitempty"`
	Search       string                 `json:"search,omitempty"`
	OfficialOnly bool                   `json:"officialOnly,omitempty"`
	Tags         []string               `json:"tags,omitempty"`
}

// GetMarketplaceListingFunctionTokens DTO.
type GetMarketplaceListingFunctionTokens struct {
	CodeMashRequestBase
	ListingViewId string `json:"listingViewId,omitempty"`
	FunctionKey   string `json:"functionKey,omitempty"`
}

// GetMarketplaceIntegrations DTO.
type GetMarketplaceIntegrations struct {
	CodeMashListPaginationRequestBase
}

// GetMarketplaceIntegration DTO.
type GetMarketplaceIntegration struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// SaveMarketplaceIntegration DTO.
type SaveMarketplaceIntegration struct {
	CodeMashRequestBase
	Integration *MarketplaceIntegrationDto `json:"integration,omitempty"`
	Secrets     map[string]string          `json:"secrets,omitempty"`
}

// DeleteMarketplaceIntegration DTO.
type DeleteMarketplaceIntegration struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// EnableMarketplaceIntegration DTO.
type EnableMarketplaceIntegration struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// DisableMarketplaceIntegration DTO.
type DisableMarketplaceIntegration struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// GetMarketplaceFunctions DTO.
type GetMarketplaceFunctions struct {
	CodeMashListPaginationRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// GetMarketplaceFunction DTO.
type GetMarketplaceFunction struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
	FunctionViewId    string `json:"functionViewId,omitempty"`
}

// SaveMarketplaceFunction DTO.
type SaveMarketplaceFunction struct {
	CodeMashRequestBase
	IntegrationViewId string                  `json:"integrationViewId,omitempty"`
	Function          *MarketplaceFunctionDto `json:"function,omitempty"`
}

// DeleteMarketplaceFunction DTO.
type DeleteMarketplaceFunction struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
	FunctionViewId    string `json:"functionViewId,omitempty"`
}

// EnableMarketplaceFunction DTO.
type EnableMarketplaceFunction struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
	FunctionViewId    string `json:"functionViewId,omitempty"`
}

// DisableMarketplaceFunction DTO.
type DisableMarketplaceFunction struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
	FunctionViewId    string `json:"functionViewId,omitempty"`
}

// GetMarketplaceFunctionTokens DTO.
type GetMarketplaceFunctionTokens struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
	FunctionViewId    string `json:"functionViewId,omitempty"`
}

// InvokeMarketplaceFunction DTO.
type InvokeMarketplaceFunction struct {
	CodeMashRequestBase
	FunctionViewId string                    `json:"functionViewId,omitempty"`
	Payload        map[string]map[string]any `json:"payload,omitempty"`
}

// GetMarketplaceListing DTO.
type GetMarketplaceListing struct {
	CodeMashRequestBase
	ListingViewId string `json:"listingViewId,omitempty"`
}

// TestMarketplaceIntegration DTO.
type TestMarketplaceIntegration struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// InternalsTypeGen DTO.
type InternalsTypeGen struct {
	Typegen_0_SmtpEmailIntegrationRequest                        *SmtpEmailIntegrationRequest                       `json:"typegen_0_SmtpEmailIntegrationRequest,omitempty"`
	Typegen_1_AwsSesEmailIntegrationRequest                      *AwsSesEmailIntegrationRequest                     `json:"typegen_1_AwsSesEmailIntegrationRequest,omitempty"`
	Typegen_2_SendGridEmailIntegrationRequest                    *SendGridEmailIntegrationRequest                   `json:"typegen_2_SendGridEmailIntegrationRequest,omitempty"`
	Typegen_3_MailGunEmailIntegrationRequest                     *MailGunEmailIntegrationRequest                    `json:"typegen_3_MailGunEmailIntegrationRequest,omitempty"`
	Typegen_4_EmailToAllUsersDeliverySettingsRequest             *EmailToAllUsersDeliverySettingsRequest            `json:"typegen_4_EmailToAllUsersDeliverySettingsRequest,omitempty"`
	Typegen_5_EmailToAccountUsersDeliverySettingsRequest         *EmailToAccountUsersDeliverySettingsRequest        `json:"typegen_5_EmailToAccountUsersDeliverySettingsRequest,omitempty"`
	Typegen_6_EmailToCollectionRecordsDeliverySettingsRequest    *EmailToCollectionRecordsDeliverySettingsRequest   `json:"typegen_6_EmailToCollectionRecordsDeliverySettingsRequest,omitempty"`
	Typegen_7_EmailToEmailsDeliverySettingsRequest               *EmailToEmailsDeliverySettingsRequest              `json:"typegen_7_EmailToEmailsDeliverySettingsRequest,omitempty"`
	Typegen_8_EmailToUsersDeliverySettingsRequest                *EmailToUsersDeliverySettingsRequest               `json:"typegen_8_EmailToUsersDeliverySettingsRequest,omitempty"`
	Typegen_9_MembershipTriggerRequest                           *MembershipTriggerRequest                          `json:"typegen_9_MembershipTriggerRequest,omitempty"`
	Typegen_10_SchemaTriggerRequest                              *SchemaTriggerRequest                              `json:"typegen_10_SchemaTriggerRequest,omitempty"`
	Typegen_11_FilesTriggerRequest                               *FilesTriggerRequest                               `json:"typegen_11_FilesTriggerRequest,omitempty"`
	Typegen_12_PaymentTriggerRequest                             *PaymentTriggerRequest                             `json:"typegen_12_PaymentTriggerRequest,omitempty"`
	Typegen_15_MongoDbConnectionStringDatabaseIntegrationRequest *MongoDbConnectionStringDatabaseIntegrationRequest `json:"typegen_15_MongoDbConnectionStringDatabaseIntegrationRequest,omitempty"`
	Typegen_16_MongoDbAtlasFlexManagedDatabaseIntegrationRequest *MongoDbAtlasFlexManagedDatabaseIntegrationRequest `json:"typegen_16_MongoDbAtlasFlexManagedDatabaseIntegrationRequest,omitempty"`
	Typegen_16_GoogleDriveFilesIntegrationRequest                *GoogleDriveFilesIntegrationRequest                `json:"typegen_16_GoogleDriveFilesIntegrationRequest,omitempty"`
	Typegen_17_FtpFilesIntegrationRequest                        *FtpFilesIntegrationRequest                        `json:"typegen_17_FtpFilesIntegrationRequest,omitempty"`
	Typegen_18_DropBoxFilesIntegrationRequest                    *DropBoxFilesIntegrationRequest                    `json:"typegen_18_DropBoxFilesIntegrationRequest,omitempty"`
	Typegen_19_AppleICloudFilesIntegrationRequest                *AppleICloudFilesIntegrationRequest                `json:"typegen_19_AppleICloudFilesIntegrationRequest,omitempty"`
	Typegen_20_AwsS3FilesIntegrationRequest                      *AwsS3FilesIntegrationRequest                      `json:"typegen_20_AwsS3FilesIntegrationRequest,omitempty"`
	Typegen_21_GoogleCloudFilesIntegrationRequest                *GoogleCloudFilesIntegrationRequest                `json:"typegen_21_GoogleCloudFilesIntegrationRequest,omitempty"`
	Typegen_22_AzureBlobFilesIntegrationRequest                  *AzureBlobFilesIntegrationRequest                  `json:"typegen_22_AzureBlobFilesIntegrationRequest,omitempty"`
	Typegen_23_LocalFilesIntegrationRequest                      *LocalFilesIntegrationRequest                      `json:"typegen_23_LocalFilesIntegrationRequest,omitempty"`
	Typegen_24_AmqpLoggingIntegrationRequest                     *AmqpLoggingIntegrationRequest                     `json:"typegen_24_AmqpLoggingIntegrationRequest,omitempty"`
	Typegen_25_AwsKinesisLoggingIntegrationRequest               *AwsKinesisLoggingIntegrationRequest               `json:"typegen_25_AwsKinesisLoggingIntegrationRequest,omitempty"`
	Typegen_26_AwsS3LoggingIntegrationRequest                    *AwsS3LoggingIntegrationRequest                    `json:"typegen_26_AwsS3LoggingIntegrationRequest,omitempty"`
	Typegen_28_NewRelicLoggingIntegrationRequest                 *NewRelicLoggingIntegrationRequest                 `json:"typegen_28_NewRelicLoggingIntegrationRequest,omitempty"`
	Typegen_30_MongoDbLoggingIntegrationRequest                  *MongoDbLoggingIntegrationRequest                  `json:"typegen_30_MongoDbLoggingIntegrationRequest,omitempty"`
	Typegen_31_KafkaLoggingIntegrationRequest                    *KafkaLoggingIntegrationRequest                    `json:"typegen_31_KafkaLoggingIntegrationRequest,omitempty"`
	Typegen_32_PrometheusLoggingIntegrationRequest               *PrometheusLoggingIntegrationRequest               `json:"typegen_32_PrometheusLoggingIntegrationRequest,omitempty"`
	Typegen_33_DataDogLoggingIntegrationRequest                  *DataDogLoggingIntegrationRequest                  `json:"typegen_33_DataDogLoggingIntegrationRequest,omitempty"`
	Typegen_34_InternalKafkaLoggingIntegrationRequest            *InternalKafkaLoggingIntegrationRequest            `json:"typegen_34_InternalKafkaLoggingIntegrationRequest,omitempty"`
	Typegen_35_ElasticSearchLoggingIntegrationRequest            *ElasticSearchLoggingIntegrationRequest            `json:"typegen_35_ElasticSearchLoggingIntegrationRequest,omitempty"`
	Typegen_37_SplunkLoggingIntegrationRequest                   *SplunkLoggingIntegrationRequest                   `json:"typegen_37_SplunkLoggingIntegrationRequest,omitempty"`
	Typegen_38_AzureOtelLoggingIntegrationRequest                *AzureOtelLoggingIntegrationRequest                `json:"typegen_38_AzureOtelLoggingIntegrationRequest,omitempty"`
	Typegen_39_KibanaLoggingIntegrationRequest                   *KibanaLoggingIntegrationRequest                   `json:"typegen_39_KibanaLoggingIntegrationRequest,omitempty"`
	Typegen_40_LocalFileLoggingIntegrationRequest                *LocalFileLoggingIntegrationRequest                `json:"typegen_40_LocalFileLoggingIntegrationRequest,omitempty"`
	Typegen_41_OktaMembershipIntegrationRequest                  *OktaMembershipIntegrationRequest                  `json:"typegen_41_OktaMembershipIntegrationRequest,omitempty"`
	Typegen_42_XMembershipIntegrationRequest                     *XMembershipIntegrationRequest                     `json:"typegen_42_XMembershipIntegrationRequest,omitempty"`
	Typegen_43_GoogleMembershipIntegrationRequest                *GoogleMembershipIntegrationRequest                `json:"typegen_43_GoogleMembershipIntegrationRequest,omitempty"`
	Typegen_44_MicrosoftMembershipIntegrationRequest             *MicrosoftMembershipIntegrationRequest             `json:"typegen_44_MicrosoftMembershipIntegrationRequest,omitempty"`
	Typegen_45_GitHubMembershipIntegrationRequest                *GitHubMembershipIntegrationRequest                `json:"typegen_45_GitHubMembershipIntegrationRequest,omitempty"`
	Typegen_46_MetaMembershipIntegrationRequest                  *MetaMembershipIntegrationRequest                  `json:"typegen_46_MetaMembershipIntegrationRequest,omitempty"`
	Typegen_47_AppleMembershipIntegrationRequest                 *AppleMembershipIntegrationRequest                 `json:"typegen_47_AppleMembershipIntegrationRequest,omitempty"`
	Typegen_48_LemonSqueezyPaymentIntegrationRequest             *LemonSqueezyPaymentIntegrationRequest             `json:"typegen_48_LemonSqueezyPaymentIntegrationRequest,omitempty"`
	Typegen_49_AdyenPaymentIntegrationRequest                    *AdyenPaymentIntegrationRequest                    `json:"typegen_49_AdyenPaymentIntegrationRequest,omitempty"`
	Typegen_50_MolliePaymentIntegrationRequest                   *MolliePaymentIntegrationRequest                   `json:"typegen_50_MolliePaymentIntegrationRequest,omitempty"`
	Typegen_51_PaddlePaymentIntegrationRequest                   *PaddlePaymentIntegrationRequest                   `json:"typegen_51_PaddlePaymentIntegrationRequest,omitempty"`
	Typegen_52_PayPalPaymentIntegrationRequest                   *PayPalPaymentIntegrationRequest                   `json:"typegen_52_PayPalPaymentIntegrationRequest,omitempty"`
	Typegen_53_StripePaymentIntegrationRequest                   *StripePaymentIntegrationRequest                   `json:"typegen_53_StripePaymentIntegrationRequest,omitempty"`
	Typegen_54_AppleInAppPaymentIntegrationRequest               *AppleInAppPaymentIntegrationRequest               `json:"typegen_54_AppleInAppPaymentIntegrationRequest,omitempty"`
	Typegen_55_GoogleInAppPaymentIntegrationRequest              *GoogleInAppPaymentIntegrationRequest              `json:"typegen_55_GoogleInAppPaymentIntegrationRequest,omitempty"`
	Typegen_56_EdgeWebPushIntegrationRequest                     *EdgeWebPushIntegrationRequest                     `json:"typegen_56_EdgeWebPushIntegrationRequest,omitempty"`
	Typegen_57_ChromePluginPushIntegrationRequest                *ChromePluginPushIntegrationRequest                `json:"typegen_57_ChromePluginPushIntegrationRequest,omitempty"`
	Typegen_58_SafariPushIntegrationRequest                      *SafariPushIntegrationRequest                      `json:"typegen_58_SafariPushIntegrationRequest,omitempty"`
	Typegen_59_ChromeWebPushIntegrationRequest                   *ChromeWebPushIntegrationRequest                   `json:"typegen_59_ChromeWebPushIntegrationRequest,omitempty"`
	Typegen_60_FirefoxWebPushIntegrationRequest                  *FirefoxWebPushIntegrationRequest                  `json:"typegen_60_FirefoxWebPushIntegrationRequest,omitempty"`
	Typegen_61_AndroidFirebasePushIntegrationRequest             *AndroidFirebasePushIntegrationRequest             `json:"typegen_61_AndroidFirebasePushIntegrationRequest,omitempty"`
	Typegen_62_AppleApnsPushIntegrationRequest                   *AppleApnsPushIntegrationRequest                   `json:"typegen_62_AppleApnsPushIntegrationRequest,omitempty"`
	Typegen_65_AwsLambdaCodeIntegrationRequest                   *AwsLambdaCodeIntegrationRequest                   `json:"typegen_65_AwsLambdaCodeIntegrationRequest,omitempty"`
	Typegen_66_AzureFunctionsCodeIntegrationRequest              *AzureFunctionsCodeIntegrationRequest              `json:"typegen_66_AzureFunctionsCodeIntegrationRequest,omitempty"`
	Typegen_67_GoogleCloudFunctionsCodeIntegrationRequest        *GoogleCloudFunctionsCodeIntegrationRequest        `json:"typegen_67_GoogleCloudFunctionsCodeIntegrationRequest,omitempty"`
	Typegen_68_OllamaLlmIntegrationRequest                       *OllamaLlmIntegrationRequest                       `json:"typegen_68_OllamaLlmIntegrationRequest,omitempty"`
	Typegen_69_OpenRouterLlmIntegrationRequest                   *OpenRouterLlmIntegrationRequest                   `json:"typegen_69_OpenRouterLlmIntegrationRequest,omitempty"`
	Typegen_70_MistralLlmIntegrationRequest                      *MistralLlmIntegrationRequest                      `json:"typegen_70_MistralLlmIntegrationRequest,omitempty"`
	Typegen_71_GrokLlmIntegrationRequest                         *GrokLlmIntegrationRequest                         `json:"typegen_71_GrokLlmIntegrationRequest,omitempty"`
	Typegen_72_GroqLlmIntegrationRequest                         *GroqLlmIntegrationRequest                         `json:"typegen_72_GroqLlmIntegrationRequest,omitempty"`
	Typegen_73_GoogleLlmIntegrationRequest                       *GoogleLlmIntegrationRequest                       `json:"typegen_73_GoogleLlmIntegrationRequest,omitempty"`
	Typegen_74_AnthropicLlmIntegrationRequest                    *AnthropicLlmIntegrationRequest                    `json:"typegen_74_AnthropicLlmIntegrationRequest,omitempty"`
	Typegen_75_OpenAiLlmIntegrationRequest                       *OpenAiLlmIntegrationRequest                       `json:"typegen_75_OpenAiLlmIntegrationRequest,omitempty"`
	Typegen_76_PlaywrightMcpIntegrationRequest                   *PlaywrightMcpIntegrationRequest                   `json:"typegen_76_PlaywrightMcpIntegrationRequest,omitempty"`
	Typegen_77_MongoDbMcpIntegrationRequest                      *MongoDbMcpIntegrationRequest                      `json:"typegen_77_MongoDbMcpIntegrationRequest,omitempty"`
	Typegen_78_GitHubMcpIntegrationRequest                       *GitHubMcpIntegrationRequest                       `json:"typegen_78_GitHubMcpIntegrationRequest,omitempty"`
	Typegen_79_StripeMcpIntegrationRequest                       *StripeMcpIntegrationRequest                       `json:"typegen_79_StripeMcpIntegrationRequest,omitempty"`
	Typegen_80_BraveSearchMcpIntegrationRequest                  *BraveSearchMcpIntegrationRequest                  `json:"typegen_80_BraveSearchMcpIntegrationRequest,omitempty"`
	Typegen_81_ObsidianMcpIntegrationRequest                     *ObsidianMcpIntegrationRequest                     `json:"typegen_81_ObsidianMcpIntegrationRequest,omitempty"`
	Typegen_82_EmailTemplateDto                                  *EmailTemplateDto                                  `json:"typegen_82_EmailTemplateDto,omitempty"`
	Typegen_83_PushTemplateDto                                   *PushTemplateDto                                   `json:"typegen_83_PushTemplateDto,omitempty"`
	Typegen_84_SmsTemplateDto                                    *SmsTemplateDto                                    `json:"typegen_84_SmsTemplateDto,omitempty"`
	Typegen_85_SystemEmailTemplateDto                            *SystemEmailTemplateDto                            `json:"typegen_85_SystemEmailTemplateDto,omitempty"`
	Typegen_86_TriggerActionEmailDto                             *TriggerActionEmailDto                             `json:"typegen_86_TriggerActionEmailDto,omitempty"`
	Typegen_87_TriggerActionPushDto                              *TriggerActionPushDto                              `json:"typegen_87_TriggerActionPushDto,omitempty"`
	Typegen_88_TriggerActionCodeDto                              *TriggerActionCodeDto                              `json:"typegen_88_TriggerActionCodeDto,omitempty"`
	Typegen_89_TriggerActionWebhookDto                           *TriggerActionWebhookDto                           `json:"typegen_89_TriggerActionWebhookDto,omitempty"`
	Typegen_236_TriggerActionSmsDto                              *TriggerActionSmsDto                               `json:"typegen_236_TriggerActionSmsDto,omitempty"`
	Typegen_237_TriggerActionSseDto                              *TriggerActionSseDto                               `json:"typegen_237_TriggerActionSseDto,omitempty"`
	Typegen_238_TriggerActionMarketplaceDto                      *TriggerActionMarketplaceDto                       `json:"typegen_238_TriggerActionMarketplaceDto,omitempty"`
	Typegen_239_SseDeliverySettingsDto                           *SseDeliverySettingsDto                            `json:"typegen_239_SseDeliverySettingsDto,omitempty"`
	Typegen_240_GetTriggers                                      *GetTriggers                                       `json:"typegen_240_GetTriggers,omitempty"`
	Typegen_241_GetTriggersResponse                              *GetTriggersResponse                               `json:"typegen_241_GetTriggersResponse,omitempty"`
	Typegen_90_EmailToAllUsersDeliverySettingsDto                *EmailToAllUsersDeliverySettingsDto                `json:"typegen_90_EmailToAllUsersDeliverySettingsDto,omitempty"`
	Typegen_91_EmailToAccountUsersDeliverySettingsDto            *EmailToAccountUsersDeliverySettingsDto            `json:"typegen_91_EmailToAccountUsersDeliverySettingsDto,omitempty"`
	Typegen_92_EmailToUsersDeliverySettingsDto                   *EmailToUsersDeliverySettingsDto                   `json:"typegen_92_EmailToUsersDeliverySettingsDto,omitempty"`
	Typegen_93_EmailToEmailAddressesDeliverySettingsDto          *EmailToEmailAddressesDeliverySettingsDto          `json:"typegen_93_EmailToEmailAddressesDeliverySettingsDto,omitempty"`
	Typegen_94_EmailToCollectionRecordsDeliverySettingsDto       *EmailToCollectionRecordsDeliverySettingsDto       `json:"typegen_94_EmailToCollectionRecordsDeliverySettingsDto,omitempty"`
	Typegen_95_PushToAllUsersDeliverySettingsDto                 *PushToAllUsersDeliverySettingsDto                 `json:"typegen_95_PushToAllUsersDeliverySettingsDto,omitempty"`
	Typegen_96_PushToUsersDeliverySettingsDto                    *PushToUsersDeliverySettingsDto                    `json:"typegen_96_PushToUsersDeliverySettingsDto,omitempty"`
	Typegen_229_PushToAccountUsersDeliverySettingsDto            *PushToAccountUsersDeliverySettingsDto             `json:"typegen_229_PushToAccountUsersDeliverySettingsDto,omitempty"`
	Typegen_97_PushToCollectionRecordsDeliverySettingsDto        *PushToCollectionRecordsDeliverySettingsDto        `json:"typegen_97_PushToCollectionRecordsDeliverySettingsDto,omitempty"`
	Typegen_98_PushToDevicesDeliverySettingsDto                  *PushToDevicesDeliverySettingsDto                  `json:"typegen_98_PushToDevicesDeliverySettingsDto,omitempty"`
	Typegen_99_SmsToAllUsersDeliverySettingsDto                  *SmsToAllUsersDeliverySettingsDto                  `json:"typegen_99_SmsToAllUsersDeliverySettingsDto,omitempty"`
	Typegen_100_SmsToUsersDeliverySettingsDto                    *SmsToUsersDeliverySettingsDto                     `json:"typegen_100_SmsToUsersDeliverySettingsDto,omitempty"`
	Typegen_101_SmsToCollectionRecordsDeliverySettingsDto        *SmsToCollectionRecordsDeliverySettingsDto         `json:"typegen_101_SmsToCollectionRecordsDeliverySettingsDto,omitempty"`
	Typegen_102_SmsToPhoneNumbersDeliverySettingsDto             *SmsToPhoneNumbersDeliverySettingsDto              `json:"typegen_102_SmsToPhoneNumbersDeliverySettingsDto,omitempty"`
	Typegen_103_OpenAiLlmIntegrationDto                          *OpenAiLlmIntegrationDto                           `json:"typegen_103_OpenAiLlmIntegrationDto,omitempty"`
	Typegen_104_AnthropicLlmIntegrationDto                       *AnthropicLlmIntegrationDto                        `json:"typegen_104_AnthropicLlmIntegrationDto,omitempty"`
	Typegen_105_OllamaLlmIntegrationDto                          *OllamaLlmIntegrationDto                           `json:"typegen_105_OllamaLlmIntegrationDto,omitempty"`
	Typegen_106_GroqLlmIntegrationDto                            *GroqLlmIntegrationDto                             `json:"typegen_106_GroqLlmIntegrationDto,omitempty"`
	Typegen_107_GoogleLlmIntegrationDto                          *GoogleLlmIntegrationDto                           `json:"typegen_107_GoogleLlmIntegrationDto,omitempty"`
	Typegen_108_MistralLlmIntegrationDto                         *MistralLlmIntegrationDto                          `json:"typegen_108_MistralLlmIntegrationDto,omitempty"`
	Typegen_109_OpenRouterLlmIntegrationDto                      *OpenRouterLlmIntegrationDto                       `json:"typegen_109_OpenRouterLlmIntegrationDto,omitempty"`
	Typegen_110_GrokLlmIntegrationDto                            *GrokLlmIntegrationDto                             `json:"typegen_110_GrokLlmIntegrationDto,omitempty"`
	Typegen_111_DockerMcpIntegrationDto                          *DockerMcpIntegrationDto                           `json:"typegen_111_DockerMcpIntegrationDto,omitempty"`
	Typegen_112_GoogleCalendarMcpIntegrationDto                  *GoogleCalendarMcpIntegrationDto                   `json:"typegen_112_GoogleCalendarMcpIntegrationDto,omitempty"`
	Typegen_113_ObsidianMcpIntegrationDto                        *ObsidianMcpIntegrationDto                         `json:"typegen_113_ObsidianMcpIntegrationDto,omitempty"`
	Typegen_114_AwsLambdaCrossAccountRoleCodeIntegrationDto      *AwsLambdaCrossAccountRoleCodeIntegrationDto       `json:"typegen_114_AwsLambdaCrossAccountRoleCodeIntegrationDto,omitempty"`
	Typegen_115_AwsLambdaIamCodeIntegrationDto                   *AwsLambdaIamCodeIntegrationDto                    `json:"typegen_115_AwsLambdaIamCodeIntegrationDto,omitempty"`
	Typegen_116_AzureFunctionsCodeIntegrationDto                 *AzureFunctionsCodeIntegrationDto                  `json:"typegen_116_AzureFunctionsCodeIntegrationDto,omitempty"`
	Typegen_118_GoogleCloudFunctionsCodeIntegrationDto           *GoogleCloudFunctionsCodeIntegrationDto            `json:"typegen_118_GoogleCloudFunctionsCodeIntegrationDto,omitempty"`
	Typegen_120_AdyenPaymentIntegrationDto                       *AdyenPaymentIntegrationDto                        `json:"typegen_120_AdyenPaymentIntegrationDto,omitempty"`
	Typegen_121_AppleInAppPaymentIntegrationDto                  *AppleInAppPaymentIntegrationDto                   `json:"typegen_121_AppleInAppPaymentIntegrationDto,omitempty"`
	Typegen_122_GoogleInAppPaymentIntegrationDto                 *GoogleInAppPaymentIntegrationDto                  `json:"typegen_122_GoogleInAppPaymentIntegrationDto,omitempty"`
	Typegen_123_LemonSqueezyPaymentIntegrationDto                *LemonSqueezyPaymentIntegrationDto                 `json:"typegen_123_LemonSqueezyPaymentIntegrationDto,omitempty"`
	Typegen_124_MolliePaymentIntegrationDto                      *MolliePaymentIntegrationDto                       `json:"typegen_124_MolliePaymentIntegrationDto,omitempty"`
	Typegen_125_PaddlePaymentIntegrationDto                      *PaddlePaymentIntegrationDto                       `json:"typegen_125_PaddlePaymentIntegrationDto,omitempty"`
	Typegen_126_PayPalPaymentIntegrationDto                      *PayPalPaymentIntegrationDto                       `json:"typegen_126_PayPalPaymentIntegrationDto,omitempty"`
	Typegen_127_StripePaymentIntegrationDto                      *StripePaymentIntegrationDto                       `json:"typegen_127_StripePaymentIntegrationDto,omitempty"`
	Typegen_184_ShopifyPaymentIntegrationDto                     *ShopifyPaymentIntegrationDto                      `json:"typegen_184_ShopifyPaymentIntegrationDto,omitempty"`
	Typegen_185_WooCommercePaymentIntegrationDto                 *WooCommercePaymentIntegrationDto                  `json:"typegen_185_WooCommercePaymentIntegrationDto,omitempty"`
	Typegen_186_MagentoPaymentIntegrationDto                     *MagentoPaymentIntegrationDto                      `json:"typegen_186_MagentoPaymentIntegrationDto,omitempty"`
	Typegen_187_BraintreePaymentIntegrationDto                   *BraintreePaymentIntegrationDto                    `json:"typegen_187_BraintreePaymentIntegrationDto,omitempty"`
	Typegen_188_AuthorizeNetPaymentIntegrationDto                *AuthorizeNetPaymentIntegrationDto                 `json:"typegen_188_AuthorizeNetPaymentIntegrationDto,omitempty"`
	Typegen_189_CheckOutComPaymentIntegrationDto                 *CheckOutComPaymentIntegrationDto                  `json:"typegen_189_CheckOutComPaymentIntegrationDto,omitempty"`
	Typegen_190_WorldpayPaymentIntegrationDto                    *WorldpayPaymentIntegrationDto                     `json:"typegen_190_WorldpayPaymentIntegrationDto,omitempty"`
	Typegen_128_AppleSignInMembershipIntegrationDto              *AppleSignInMembershipIntegrationDto               `json:"typegen_128_AppleSignInMembershipIntegrationDto,omitempty"`
	Typegen_129_GitHubMembershipIntegrationDto                   *GitHubMembershipIntegrationDto                    `json:"typegen_129_GitHubMembershipIntegrationDto,omitempty"`
	Typegen_130_GoogleMembershipIntegrationDto                   *GoogleMembershipIntegrationDto                    `json:"typegen_130_GoogleMembershipIntegrationDto,omitempty"`
	Typegen_131_MetaMembershipIntegrationDto                     *MetaMembershipIntegrationDto                      `json:"typegen_131_MetaMembershipIntegrationDto,omitempty"`
	Typegen_132_MicrosoftMembershipIntegrationDto                *MicrosoftMembershipIntegrationDto                 `json:"typegen_132_MicrosoftMembershipIntegrationDto,omitempty"`
	Typegen_133_OktaMembershipIntegrationDto                     *OktaMembershipIntegrationDto                      `json:"typegen_133_OktaMembershipIntegrationDto,omitempty"`
	Typegen_134_XMembershipIntegrationDto                        *XMembershipIntegrationDto                         `json:"typegen_134_XMembershipIntegrationDto,omitempty"`
	Typegen_135_AmqpLoggingIntegrationDto                        *AmqpLoggingIntegrationDto                         `json:"typegen_135_AmqpLoggingIntegrationDto,omitempty"`
	Typegen_136_AwsKinesisLoggingIntegrationDto                  *AwsKinesisLoggingIntegrationDto                   `json:"typegen_136_AwsKinesisLoggingIntegrationDto,omitempty"`
	Typegen_137_AwsS3CrossAccountRoleLoggingIntegrationDto       *AwsS3CrossAccountRoleLoggingIntegrationDto        `json:"typegen_137_AwsS3CrossAccountRoleLoggingIntegrationDto,omitempty"`
	Typegen_138_AwsS3IamLoggingIntegrationDto                    *AwsS3IamLoggingIntegrationDto                     `json:"typegen_138_AwsS3IamLoggingIntegrationDto,omitempty"`
	Typegen_139_AzureOtelLoggingIntegrationDto                   *AzureOtelLoggingIntegrationDto                    `json:"typegen_139_AzureOtelLoggingIntegrationDto,omitempty"`
	Typegen_140_DataDogLoggingIntegrationDto                     *DataDogLoggingIntegrationDto                      `json:"typegen_140_DataDogLoggingIntegrationDto,omitempty"`
	Typegen_141_ElasticSearchLoggingIntegrationDto               *ElasticSearchLoggingIntegrationDto                `json:"typegen_141_ElasticSearchLoggingIntegrationDto,omitempty"`
	Typegen_142_InternalKafkaLoggingIntegrationDto               *InternalKafkaLoggingIntegrationDto                `json:"typegen_142_InternalKafkaLoggingIntegrationDto,omitempty"`
	Typegen_143_KafkaLoggingIntegrationDto                       *KafkaLoggingIntegrationDto                        `json:"typegen_143_KafkaLoggingIntegrationDto,omitempty"`
	Typegen_144_KibanaLoggingIntegrationDto                      *KibanaLoggingIntegrationDto                       `json:"typegen_144_KibanaLoggingIntegrationDto,omitempty"`
	Typegen_145_LocalFileLoggingIntegrationDto                   *LocalFileLoggingIntegrationDto                    `json:"typegen_145_LocalFileLoggingIntegrationDto,omitempty"`
	Typegen_147_MongoDbLoggingIntegrationDto                     *MongoDbLoggingIntegrationDto                      `json:"typegen_147_MongoDbLoggingIntegrationDto,omitempty"`
	Typegen_148_NewRelicLoggingIntegrationDto                    *NewRelicLoggingIntegrationDto                     `json:"typegen_148_NewRelicLoggingIntegrationDto,omitempty"`
	Typegen_149_PrometheusLoggingIntegrationDto                  *PrometheusLoggingIntegrationDto                   `json:"typegen_149_PrometheusLoggingIntegrationDto,omitempty"`
	Typegen_150_SplunkLoggingIntegrationDto                      *SplunkLoggingIntegrationDto                       `json:"typegen_150_SplunkLoggingIntegrationDto,omitempty"`
	Typegen_153_AppleICloudFilesIntegrationDto                   *AppleICloudFilesIntegrationDto                    `json:"typegen_153_AppleICloudFilesIntegrationDto,omitempty"`
	Typegen_154_AwsS3CrossAccountRoleFilesIntegrationDto         *AwsS3CrossAccountRoleFilesIntegrationDto          `json:"typegen_154_AwsS3CrossAccountRoleFilesIntegrationDto,omitempty"`
	Typegen_155_AwsS3IamFilesIntegrationDto                      *AwsS3IamFilesIntegrationDto                       `json:"typegen_155_AwsS3IamFilesIntegrationDto,omitempty"`
	Typegen_156_AzureBlobFilesIntegrationDto                     *AzureBlobFilesIntegrationDto                      `json:"typegen_156_AzureBlobFilesIntegrationDto,omitempty"`
	Typegen_157_DropBoxFilesIntegrationDto                       *DropBoxFilesIntegrationDto                        `json:"typegen_157_DropBoxFilesIntegrationDto,omitempty"`
	Typegen_158_FtpFilesIntegrationDto                           *FtpFilesIntegrationDto                            `json:"typegen_158_FtpFilesIntegrationDto,omitempty"`
	Typegen_159_GoogleCloudFilesIntegrationDto                   *GoogleCloudFilesIntegrationDto                    `json:"typegen_159_GoogleCloudFilesIntegrationDto,omitempty"`
	Typegen_160_GoogleDriveFilesIntegrationDto                   *GoogleDriveFilesIntegrationDto                    `json:"typegen_160_GoogleDriveFilesIntegrationDto,omitempty"`
	Typegen_161_LocalFilesIntegrationDto                         *LocalFilesIntegrationDto                          `json:"typegen_161_LocalFilesIntegrationDto,omitempty"`
	Typegen_164_MongoDbConnectionStringIntegrationDto            *MongoDbConnectionStringIntegrationDto             `json:"typegen_164_MongoDbConnectionStringIntegrationDto,omitempty"`
	Typegen_165_MongoDbAtlasFlexManagedIntegrationDto            *MongoDbAtlasFlexManagedIntegrationDto             `json:"typegen_165_MongoDbAtlasFlexManagedIntegrationDto,omitempty"`
	Typegen_165_BirdSmsIntegrationDto                            *BirdSmsIntegrationDto                             `json:"typegen_165_BirdSmsIntegrationDto,omitempty"`
	Typegen_166_PlivoSmsIntegrationDto                           *PlivoSmsIntegrationDto                            `json:"typegen_166_PlivoSmsIntegrationDto,omitempty"`
	Typegen_167_SinchSmsIntegrationDto                           *SinchSmsIntegrationDto                            `json:"typegen_167_SinchSmsIntegrationDto,omitempty"`
	Typegen_168_TelesignSmsIntegrationDto                        *TelesignSmsIntegrationDto                         `json:"typegen_168_TelesignSmsIntegrationDto,omitempty"`
	Typegen_169_TelnyxSmsIntegrationDto                          *TelnyxSmsIntegrationDto                           `json:"typegen_169_TelnyxSmsIntegrationDto,omitempty"`
	Typegen_170_TwilioSmsIntegrationDto                          *TwilioSmsIntegrationDto                           `json:"typegen_170_TwilioSmsIntegrationDto,omitempty"`
	Typegen_171_VonageSmsIntegrationDto                          *VonageSmsIntegrationDto                           `json:"typegen_171_VonageSmsIntegrationDto,omitempty"`
	Typegen_246_FakeSmsIntegrationDto                            *FakeSmsIntegrationDto                             `json:"typegen_246_FakeSmsIntegrationDto,omitempty"`
	Typegen_172_AndroidFirebasePushIntegrationDto                *AndroidFirebasePushIntegrationDto                 `json:"typegen_172_AndroidFirebasePushIntegrationDto,omitempty"`
	Typegen_173_AppleApnsPushIntegrationDto                      *AppleApnsPushIntegrationDto                       `json:"typegen_173_AppleApnsPushIntegrationDto,omitempty"`
	Typegen_174_ChromePluginPushIntegrationDto                   *ChromePluginPushIntegrationDto                    `json:"typegen_174_ChromePluginPushIntegrationDto,omitempty"`
	Typegen_175_ChromeWebPushIntegrationDto                      *ChromeWebPushIntegrationDto                       `json:"typegen_175_ChromeWebPushIntegrationDto,omitempty"`
	Typegen_176_EdgeWebPushIntegrationDto                        *EdgeWebPushIntegrationDto                         `json:"typegen_176_EdgeWebPushIntegrationDto,omitempty"`
	Typegen_177_FirefoxWebPushIntegrationDto                     *FirefoxWebPushIntegrationDto                      `json:"typegen_177_FirefoxWebPushIntegrationDto,omitempty"`
	Typegen_178_SafariPushIntegrationDto                         *SafariPushIntegrationDto                          `json:"typegen_178_SafariPushIntegrationDto,omitempty"`
	Typegen_247_FakePushIntegrationDto                           *FakePushIntegrationDto                            `json:"typegen_247_FakePushIntegrationDto,omitempty"`
	Typegen_179_AwsCrossAccountRoleEmailIntegrationDto           *AwsCrossAccountRoleEmailIntegrationDto            `json:"typegen_179_AwsCrossAccountRoleEmailIntegrationDto,omitempty"`
	Typegen_180_AwsIamEmailIntegrationDto                        *AwsIamEmailIntegrationDto                         `json:"typegen_180_AwsIamEmailIntegrationDto,omitempty"`
	Typegen_181_MailGunEmailIntegrationDto                       *MailGunEmailIntegrationDto                        `json:"typegen_181_MailGunEmailIntegrationDto,omitempty"`
	Typegen_182_SendGridEmailIntegrationDto                      *SendGridEmailIntegrationDto                       `json:"typegen_182_SendGridEmailIntegrationDto,omitempty"`
	Typegen_183_SmtpEmailIntegrationDto                          *SmtpEmailIntegrationDto                           `json:"typegen_183_SmtpEmailIntegrationDto,omitempty"`
	Typegen_248_FakeEmailIntegrationDto                          *FakeEmailIntegrationDto                           `json:"typegen_248_FakeEmailIntegrationDto,omitempty"`
	Typegen_192_WebhookIntegrationDto                            *WebhookIntegrationDto                             `json:"typegen_192_WebhookIntegrationDto,omitempty"`
	Typegen_193_WebhookDestinationDto                            *WebhookDestinationDto                             `json:"typegen_193_WebhookDestinationDto,omitempty"`
	Typegen_194_SchedulerTaskDto                                 *SchedulerTaskDto                                  `json:"typegen_194_SchedulerTaskDto,omitempty"`
	Typegen_195_MongoDbAggregateDto                              *MongoDbAggregateDto                               `json:"typegen_195_MongoDbAggregateDto,omitempty"`
	Typegen_196_MarketplaceIntegrationDto                        *MarketplaceIntegrationDto                         `json:"typegen_196_MarketplaceIntegrationDto,omitempty"`
	Typegen_197_MarketplaceFunctionDto                           *MarketplaceFunctionDto                            `json:"typegen_197_MarketplaceFunctionDto,omitempty"`
	Typegen_198_MarketplaceListingDto                            *MarketplaceListingDto                             `json:"typegen_198_MarketplaceListingDto,omitempty"`
	Typegen_199_MarketplaceFunctionDefinitionDto                 *MarketplaceFunctionDefinitionDto                  `json:"typegen_199_MarketplaceFunctionDefinitionDto,omitempty"`
	Typegen_200_MarketplaceFunctionParameterDto                  *MarketplaceFunctionParameterDto                   `json:"typegen_200_MarketplaceFunctionParameterDto,omitempty"`
	Typegen_201_EnableCode                                       *EnableCode                                        `json:"typegen_201_EnableCode,omitempty"`
	Typegen_202_DisableCode                                      *DisableCode                                       `json:"typegen_202_DisableCode,omitempty"`
	Typegen_203_GetCodeIntegrations                              *GetCodeIntegrations                               `json:"typegen_203_GetCodeIntegrations,omitempty"`
	Typegen_204_GetCodeIntegration                               *GetCodeIntegration                                `json:"typegen_204_GetCodeIntegration,omitempty"`
	Typegen_205_SaveCodeIntegration                              *SaveCodeIntegration                               `json:"typegen_205_SaveCodeIntegration,omitempty"`
	Typegen_206_TestCodeIntegration                              *TestCodeIntegration                               `json:"typegen_206_TestCodeIntegration,omitempty"`
	Typegen_207_ConfirmCodeIntegrationHumanDeliveryRequest       *ConfirmCodeIntegrationHumanDeliveryRequest        `json:"typegen_207_ConfirmCodeIntegrationHumanDeliveryRequest,omitempty"`
	Typegen_208_SetCodeIntegrationAsDefault                      *SetCodeIntegrationAsDefault                       `json:"typegen_208_SetCodeIntegrationAsDefault,omitempty"`
	Typegen_209_DeleteCodeIntegrationRequest                     *DeleteCodeIntegrationRequest                      `json:"typegen_209_DeleteCodeIntegrationRequest,omitempty"`
	Typegen_210_EnableCodeIntegrationRequest                     *EnableCodeIntegrationRequest                      `json:"typegen_210_EnableCodeIntegrationRequest,omitempty"`
	Typegen_211_DisableCodeIntegrationRequest                    *DisableCodeIntegrationRequest                     `json:"typegen_211_DisableCodeIntegrationRequest,omitempty"`
	Typegen_212_GetMarketplaceListings                           *GetMarketplaceListings                            `json:"typegen_212_GetMarketplaceListings,omitempty"`
	Typegen_213_GetMarketplaceListingFunctionTokens              *GetMarketplaceListingFunctionTokens               `json:"typegen_213_GetMarketplaceListingFunctionTokens,omitempty"`
	Typegen_214_GetMarketplaceIntegrations                       *GetMarketplaceIntegrations                        `json:"typegen_214_GetMarketplaceIntegrations,omitempty"`
	Typegen_215_GetMarketplaceIntegration                        *GetMarketplaceIntegration                         `json:"typegen_215_GetMarketplaceIntegration,omitempty"`
	Typegen_216_SaveMarketplaceIntegration                       *SaveMarketplaceIntegration                        `json:"typegen_216_SaveMarketplaceIntegration,omitempty"`
	Typegen_217_DeleteMarketplaceIntegration                     *DeleteMarketplaceIntegration                      `json:"typegen_217_DeleteMarketplaceIntegration,omitempty"`
	Typegen_218_EnableMarketplaceIntegration                     *EnableMarketplaceIntegration                      `json:"typegen_218_EnableMarketplaceIntegration,omitempty"`
	Typegen_219_DisableMarketplaceIntegration                    *DisableMarketplaceIntegration                     `json:"typegen_219_DisableMarketplaceIntegration,omitempty"`
	Typegen_221_GetMarketplaceFunctions                          *GetMarketplaceFunctions                           `json:"typegen_221_GetMarketplaceFunctions,omitempty"`
	Typegen_222_GetMarketplaceFunction                           *GetMarketplaceFunction                            `json:"typegen_222_GetMarketplaceFunction,omitempty"`
	Typegen_223_SaveMarketplaceFunction                          *SaveMarketplaceFunction                           `json:"typegen_223_SaveMarketplaceFunction,omitempty"`
	Typegen_224_DeleteMarketplaceFunction                        *DeleteMarketplaceFunction                         `json:"typegen_224_DeleteMarketplaceFunction,omitempty"`
	Typegen_225_EnableMarketplaceFunction                        *EnableMarketplaceFunction                         `json:"typegen_225_EnableMarketplaceFunction,omitempty"`
	Typegen_226_DisableMarketplaceFunction                       *DisableMarketplaceFunction                        `json:"typegen_226_DisableMarketplaceFunction,omitempty"`
	Typegen_227_GetMarketplaceFunctionTokens                     *GetMarketplaceFunctionTokens                      `json:"typegen_227_GetMarketplaceFunctionTokens,omitempty"`
	Typegen_228_InvokeMarketplaceFunction                        *InvokeMarketplaceFunction                         `json:"typegen_228_InvokeMarketplaceFunction,omitempty"`
	Typegen_232_GetMarketplaceListing                            *GetMarketplaceListing                             `json:"typegen_232_GetMarketplaceListing,omitempty"`
	Typegen_233_TestMarketplaceIntegration                       *TestMarketplaceIntegration                        `json:"typegen_233_TestMarketplaceIntegration,omitempty"`
	Typegen_234_TestMarketplaceIntegrationResponse               *TestMarketplaceIntegrationResponse                `json:"typegen_234_TestMarketplaceIntegrationResponse,omitempty"`
	Typegen_235_GetMarketplaceListingResponse                    *GetMarketplaceListingResponse                     `json:"typegen_235_GetMarketplaceListingResponse,omitempty"`
	Typegen_230_AdminPortalStructureDto                          *AdminPortalStructureDto                           `json:"typegen_230_AdminPortalStructureDto,omitempty"`
	Typegen_231_AdminPortalModuleDto                             *AdminPortalModuleDto                              `json:"typegen_231_AdminPortalModuleDto,omitempty"`
	Typegen_236_UserMessageEntryWireDto                          *UserMessageEntryWireDto                           `json:"typegen_236_UserMessageEntryWireDto,omitempty"`
	Typegen_237_AssistantTextEntryWireDto                        *AssistantTextEntryWireDto                         `json:"typegen_237_AssistantTextEntryWireDto,omitempty"`
	Typegen_238_AssistantQuestionEntryWireDto                    *AssistantQuestionEntryWireDto                     `json:"typegen_238_AssistantQuestionEntryWireDto,omitempty"`
	Typegen_239_UserAnswerEntryWireDto                           *UserAnswerEntryWireDto                            `json:"typegen_239_UserAnswerEntryWireDto,omitempty"`
	Typegen_240_PlanEntryWireDto                                 *PlanEntryWireDto                                  `json:"typegen_240_PlanEntryWireDto,omitempty"`
	Typegen_241_UserDecisionEntryWireDto                         *UserDecisionEntryWireDto                          `json:"typegen_241_UserDecisionEntryWireDto,omitempty"`
	Typegen_242_RunStepEntryWireDto                              *RunStepEntryWireDto                               `json:"typegen_242_RunStepEntryWireDto,omitempty"`
	Typegen_243_ActionPendingEntryWireDto                        *ActionPendingEntryWireDto                         `json:"typegen_243_ActionPendingEntryWireDto,omitempty"`
	Typegen_244_NoticeEntryWireDto                               *NoticeEntryWireDto                                `json:"typegen_244_NoticeEntryWireDto,omitempty"`
	Typegen_245_ConversationSnapshotEntryWireDto                 *ConversationSnapshotEntryWireDto                  `json:"typegen_245_ConversationSnapshotEntryWireDto,omitempty"`
}

// Echo DTO.
type Echo struct {
	RequestBase
}

// GetPublicProjectConfig DTO.
type GetPublicProjectConfig struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
}

// GetPublicProjectLegal DTO.
type GetPublicProjectLegal struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Kind      string `json:"kind,omitempty"`
}

// GetAccountProfile DTO.
type GetAccountProfile struct {
	RequestBase
}

// UpdateAccountProfile DTO.
type UpdateAccountProfile struct {
	RequestBase
	DisplayName     string `json:"displayName,omitempty"`
	BillingEmail    string `json:"billingEmail,omitempty"`
	OperationsEmail string `json:"operationsEmail,omitempty"`
	SecurityEmail   string `json:"securityEmail,omitempty"`
}

// ResendAccountVerificationToken DTO.
type ResendAccountVerificationToken struct {
	RequestBase
}

// GetAccountStatus DTO.
type GetAccountStatus struct {
	RequestBase
}

// CreateStripeCheckoutSession DTO.
type CreateStripeCheckoutSession struct {
	RequestBase
	SubscriptionType    SubscriptionType `json:"subscriptionType,omitempty"`
	Domain              string           `json:"domain,omitempty"`
	ProjectCap          float64          `json:"projectCap,omitempty"`
	NewProjectSessionId string           `json:"newProjectSessionId,omitempty"`
	ReturnUrl           string           `json:"returnUrl,omitempty"`
}

// GetStripeBillingPortalUrl DTO.
type GetStripeBillingPortalUrl struct {
	RequestBase
	SubscriptionType SubscriptionType `json:"subscriptionType,omitempty"`
	ReturnUrl        string           `json:"returnUrl,omitempty"`
}

// CreateTeamMemberFromInvitation DTO.
type CreateTeamMemberFromInvitation struct {
	RequestBase
	DisplayName string `json:"displayName,omitempty"`
	Token       string `json:"token,omitempty"`
	Password    string `json:"password,omitempty"`
}

// GetAccountUsageBilling DTO.
type GetAccountUsageBilling struct {
	RequestBase
}

// VerifyAccount DTO.
type VerifyAccount struct {
	RequestBase
	Token     string `json:"token,omitempty"`
	AccountId string `json:"accountId,omitempty"`
}

// DeleteNotificationsGroup DTO.
type DeleteNotificationsGroup struct {
	CodeMashRequestBase
	GroupTag string `json:"groupTag,omitempty"`
}

// DeleteNotificationsTag DTO.
type DeleteNotificationsTag struct {
	CodeMashRequestBase
	Tag string `json:"tag,omitempty"`
}

// RemoveTagFromNotificationsGroup DTO.
type RemoveTagFromNotificationsGroup struct {
	CodeMashRequestBase
	GroupTag string `json:"groupTag,omitempty"`
	Tag      string `json:"tag,omitempty"`
}

// SaveNotificationsGroup DTO.
type SaveNotificationsGroup struct {
	CodeMashRequestBase
	GroupDefinition *GroupDefinitionDto  `json:"groupDefinition,omitempty"`
	Channel         CommunicationChannel `json:"channel,omitempty"`
	OriginChannel   CommunicationChannel `json:"originChannel,omitempty"`
}

// SaveNotificationsTag DTO.
type SaveNotificationsTag struct {
	CodeMashRequestBase
	TagDefinition *TagDefinitionDto    `json:"tagDefinition,omitempty"`
	Channel       CommunicationChannel `json:"channel,omitempty"`
	GroupTag      string               `json:"groupTag,omitempty"`
}

// CreateProjectRequest DTO.
type CreateProjectRequest struct {
	RequestBase
	Integration       *DatabaseIntegrationRequest `json:"integration,omitempty"`
	ProjectName       string                      `json:"projectName,omitempty"`
	PrimaryRegion     string                      `json:"primaryRegion,omitempty"`
	AdditionalRegions []string                    `json:"additionalRegions,omitempty"`
	Description       string                      `json:"description,omitempty"`
}

// DeleteProject DTO.
type DeleteProject struct {
	CodeMashRequestBase
}

// CreateProjectEnvironmentRequest DTO.
type CreateProjectEnvironmentRequest struct {
	CodeMashRequestBase
	EnvironmentName string                      `json:"environmentName,omitempty"`
	Integration     *DatabaseIntegrationRequest `json:"integration,omitempty"`
}

// DeleteProjectEnvironmentRequest DTO.
type DeleteProjectEnvironmentRequest struct {
	CodeMashRequestBase
	EnvironmentName string `json:"environmentName,omitempty"`
}

// SetEnvironmentRankRequest DTO.
type SetEnvironmentRankRequest struct {
	CodeMashRequestBase
	EnvironmentName string  `json:"environmentName,omitempty"`
	Rank            float64 `json:"rank,omitempty"`
}

// PromoteEnvironmentRequest DTO.
type PromoteEnvironmentRequest struct {
	CodeMashRequestBase
	SourceEnv string `json:"sourceEnv,omitempty"`
	TargetEnv string `json:"targetEnv,omitempty"`
	DryRun    bool   `json:"dryRun,omitempty"`
}

// RollbackPromotionRequest DTO.
type RollbackPromotionRequest struct {
	CodeMashRequestBase
	TargetEnv   string  `json:"targetEnv,omitempty"`
	FromVersion float64 `json:"fromVersion,omitempty"`
}

// GetProjectEnvironments DTO.
type GetProjectEnvironments struct {
	CodeMashRequestBase
}

// GetProject DTO.
type GetProject struct {
	CodeMashRequestBase
}

// GetProjects DTO.
type GetProjects struct {
	RequestBase
}

// GetAccountRegions DTO.
type GetAccountRegions struct {
	RequestBase
}

// WaitForProjectActiveRequest DTO.
type WaitForProjectActiveRequest struct {
	CodeMashRequestBase
	TimeoutSeconds float64 `json:"timeoutSeconds,omitempty"`
}

// GetProjectTokens DTO.
type GetProjectTokens struct {
	CodeMashRequestBase
	InitiatorId                string `json:"initiatorId,omitempty"`
	RecipientId                string `json:"recipientId,omitempty"`
	TargetUserId               string `json:"targetUserId,omitempty"`
	MembershipTriggerOldUserId string `json:"membershipTriggerOldUserId,omitempty"`
	MembershipTriggerNewUserId string `json:"membershipTriggerNewUserId,omitempty"`
}

// AssignAdminPortalServiceUserRequest DTO.
type AssignAdminPortalServiceUserRequest struct {
	CodeMashRequestBase
	ServiceUserId string `json:"serviceUserId,omitempty"`
}

// GetAdminPortalStructure DTO.
type GetAdminPortalStructure struct {
	CodeMashRequestBase
}

// UpdateProjectAdminUrl DTO.
type UpdateProjectAdminUrl struct {
	CodeMashRequestBase
	Url string `json:"url,omitempty"`
}

// UpdateProjectAccentColor DTO.
type UpdateProjectAccentColor struct {
	CodeMashRequestBase
	Color string `json:"color,omitempty"`
}

// UpdateProjectIcon DTO.
type UpdateProjectIcon struct {
	CodeMashRequestBase
	FileResource *FileResourceRefDto `json:"fileResource,omitempty"`
}

// UpdateProjectLogo DTO.
type UpdateProjectLogo struct {
	CodeMashRequestBase
	FileResource *FileResourceRefDto `json:"fileResource,omitempty"`
}

// UpdateProjectMainColor DTO.
type UpdateProjectMainColor struct {
	CodeMashRequestBase
	Color string `json:"color,omitempty"`
}

// UpdateProjectAllowedOrigins DTO.
type UpdateProjectAllowedOrigins struct {
	CodeMashRequestBase
	Origins []string `json:"origins,omitempty"`
}

// UpdateProjectDefaultLanguage DTO.
type UpdateProjectDefaultLanguage struct {
	CodeMashRequestBase
	DefaultLanguage string `json:"defaultLanguage,omitempty"`
}

// UpdateProjectDescription DTO.
type UpdateProjectDescription struct {
	CodeMashRequestBase
	Description string `json:"description,omitempty"`
}

// DisableProject DTO.
type DisableProject struct {
	CodeMashRequestBase
}

// EnableProject DTO.
type EnableProject struct {
	CodeMashRequestBase
}

// UpdateProjectLanguages DTO.
type UpdateProjectLanguages struct {
	CodeMashRequestBase
	Languages []string `json:"languages,omitempty"`
}

// UpdateProjectLegalDocuments DTO.
type UpdateProjectLegalDocuments struct {
	CodeMashRequestBase
	TermsMarkdown   string `json:"termsMarkdown,omitempty"`
	PrivacyMarkdown string `json:"privacyMarkdown,omitempty"`
}

// UpdateProjectExposeLegal DTO.
type UpdateProjectExposeLegal struct {
	CodeMashRequestBase
	Exposed bool `json:"exposed,omitempty"`
}

// UpdateProjectUrl DTO.
type UpdateProjectUrl struct {
	CodeMashRequestBase
	Url string `json:"url,omitempty"`
}

// UpdateProjectName DTO.
type UpdateProjectName struct {
	CodeMashRequestBase
	Name string `json:"name,omitempty"`
}

// UpdateProjectRegions DTO.
type UpdateProjectRegions struct {
	CodeMashRequestBase
	PrimaryRegion     string   `json:"primaryRegion,omitempty"`
	AdditionalRegions []string `json:"additionalRegions,omitempty"`
}

// CreateAccount DTO.
type CreateAccount struct {
	RequestBase
	DisplayName string `json:"displayName,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
}

// ChangeTeamMemberPassword DTO.
type ChangeTeamMemberPassword struct {
	RequestBase
	Email           string `json:"email,omitempty"`
	CurrentPassword string `json:"currentPassword,omitempty"`
	NewPassword     string `json:"newPassword,omitempty"`
}

// CreateTeamMember DTO.
type CreateTeamMember struct {
	RequestBase
	Email          string   `json:"email,omitempty"`
	DisplayName    string   `json:"displayName,omitempty"`
	Password       string   `json:"password,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	SendInvitation bool     `json:"sendInvitation,omitempty"`
}

// CreateAccountPolicy DTO.
type CreateAccountPolicy struct {
	RequestBase
	PolicyName         string `json:"policyName,omitempty"`
	Description        string `json:"description,omitempty"`
	PolicyDocumentJson string `json:"policyDocumentJson,omitempty"`
}

// CreateAccountRole DTO.
type CreateAccountRole struct {
	RequestBase
	RoleName    string   `json:"roleName,omitempty"`
	Description string   `json:"description,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

// DeleteAccountPolicy DTO.
type DeleteAccountPolicy struct {
	RequestBase
	Id string `json:"id,omitempty"`
}

// DeleteAccountRole DTO.
type DeleteAccountRole struct {
	RequestBase
	Id string `json:"id,omitempty"`
}

// GetAccountCollaborators DTO.
type GetAccountCollaborators struct {
	RequestBase
	IncludeAccountOwner      bool        `json:"includeAccountOwner,omitempty"`
	UserShouldHavePushDevice bool        `json:"userShouldHavePushDevice,omitempty"`
	ProjectId                string      `json:"projectId,omitempty"`
	UserIds                  []string    `json:"userIds,omitempty"`
	RoleNames                []string    `json:"roleNames,omitempty"`
	PagingArgs               *PagingArgs `json:"pagingArgs,omitempty"`
}

// GetAccountPasswordPolicy DTO.
type GetAccountPasswordPolicy struct {
	RequestBase
}

// GetAccountTeamPolicies DTO.
type GetAccountTeamPolicies struct {
	RequestBase
}

// GetAccountTeamRoles DTO.
type GetAccountTeamRoles struct {
	RequestBase
}

// SendInviteToTeamMember DTO.
type SendInviteToTeamMember struct {
	RequestBase
	Email string   `json:"email,omitempty"`
	Roles []string `json:"roles,omitempty"`
}

// UpdateAccountPolicy DTO.
type UpdateAccountPolicy struct {
	RequestBase
	Id                 string `json:"id,omitempty"`
	PolicyName         string `json:"policyName,omitempty"`
	Description        string `json:"description,omitempty"`
	PolicyDocumentJson string `json:"policyDocumentJson,omitempty"`
}

// UpdateAccountRole DTO.
type UpdateAccountRole struct {
	RequestBase
	Id          string   `json:"id,omitempty"`
	RoleName    string   `json:"roleName,omitempty"`
	Description string   `json:"description,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

// AccountHasPasskeyRequest DTO.
type AccountHasPasskeyRequest struct {
	RequestBase
	Email string `json:"email,omitempty"`
}

// AccountStartEmailVerificationRequest DTO.
type AccountStartEmailVerificationRequest struct {
	RequestBase
	Email string `json:"email,omitempty"`
}

// AccountConfirmEmailVerificationRequest DTO.
type AccountConfirmEmailVerificationRequest struct {
	RequestBase
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

// AccountPasskeyRegistrationOptionsRequest DTO.
type AccountPasskeyRegistrationOptionsRequest struct {
	RequestBase
	VerificationToken string `json:"verificationToken,omitempty"`
}

// AccountVerifyPasskeyRegistrationRequest DTO.
type AccountVerifyPasskeyRegistrationRequest struct {
	RequestBase
	VerificationToken   string `json:"verificationToken,omitempty"`
	CeremonyId          string `json:"ceremonyId,omitempty"`
	AttestationResponse string `json:"attestationResponse,omitempty"`
	FriendlyName        string `json:"friendlyName,omitempty"`
}

// AccountPasskeyAuthenticationOptionsRequest DTO.
type AccountPasskeyAuthenticationOptionsRequest struct {
	RequestBase
	Email string `json:"email,omitempty"`
}

// AccountVerifyPasskeyAuthenticationRequest DTO.
type AccountVerifyPasskeyAuthenticationRequest struct {
	RequestBase
	CeremonyId        string `json:"ceremonyId,omitempty"`
	AssertionResponse string `json:"assertionResponse,omitempty"`
}

// ListAccountPasskeysRequest DTO.
type ListAccountPasskeysRequest struct {
	RequestBase
}

// RenameAccountPasskeyRequest DTO.
type RenameAccountPasskeyRequest struct {
	RequestBase
	CredentialId string `json:"credentialId,omitempty"`
	FriendlyName string `json:"friendlyName,omitempty"`
}

// RevokeAccountPasskeyRequest DTO.
type RevokeAccountPasskeyRequest struct {
	RequestBase
	CredentialId string `json:"credentialId,omitempty"`
}

// AccountPasskeyEnrollmentOptionsRequest DTO.
type AccountPasskeyEnrollmentOptionsRequest struct {
	RequestBase
}

// AccountVerifyPasskeyEnrollmentRequest DTO.
type AccountVerifyPasskeyEnrollmentRequest struct {
	RequestBase
	CeremonyId          string `json:"ceremonyId,omitempty"`
	AttestationResponse string `json:"attestationResponse,omitempty"`
	FriendlyName        string `json:"friendlyName,omitempty"`
}

// GetLicenseDomainDnsStatus DTO.
type GetLicenseDomainDnsStatus struct {
	RequestBase
	Domain string `json:"domain,omitempty"`
}

// StartLicenseDomainVerificationRequest DTO.
type StartLicenseDomainVerificationRequest struct {
	RequestBase
	Domain string `json:"domain,omitempty"`
}

// GetLicenseDomainVerificationStatus DTO.
type GetLicenseDomainVerificationStatus struct {
	RequestBase
	Domain string `json:"domain,omitempty"`
}

// GetLicenses DTO.
type GetLicenses struct {
	RequestBase
}

// PostLicenseHeartbeat DTO.
type PostLicenseHeartbeat struct {
	RequestBase
	License          string `json:"license,omitempty"`
	LicenseAccountId string `json:"licenseAccountId,omitempty"`
	InstallationId   string `json:"installationId,omitempty"`
	Domain           string `json:"domain,omitempty"`
	HostKind         string `json:"hostKind,omitempty"`
	Release          string `json:"release,omitempty"`
	InstanceVersion  string `json:"instanceVersion,omitempty"`
}

// GetInstallationLicenseStatus DTO.
type GetInstallationLicenseStatus struct {
	RequestBase
}

// AccountCreated DTO.
type AccountCreated struct {
	Email       *EmailAddress `json:"email,omitempty"`
	DisplayName *DisplayName  `json:"displayName,omitempty"`
	AccountId   *AccountId    `json:"accountId,omitempty"`
	CreatedOn   *UtcDateTime  `json:"createdOn,omitempty"`
}

// AccountProfileUpdated DTO.
type AccountProfileUpdated struct {
	DisplayName     *DisplayName  `json:"displayName,omitempty"`
	BillingEmail    *EmailAddress `json:"billingEmail,omitempty"`
	OperationsEmail *EmailAddress `json:"operationsEmail,omitempty"`
	SecurityEmail   *EmailAddress `json:"securityEmail,omitempty"`
}

// AccountSetAsActive DTO.
type AccountSetAsActive struct {
}

// AccountValidationTokenIssued DTO.
type AccountValidationTokenIssued struct {
	Expiration *ExpirationToken `json:"expiration,omitempty"`
}

// AccountVerified DTO.
type AccountVerified struct {
}

// AccountBlocked DTO.
type AccountBlocked struct {
}

// AccountSetAsInactive DTO.
type AccountSetAsInactive struct {
}

// AccountUnregistered DTO.
type AccountUnregistered struct {
}

// LicenseCreated DTO.
type LicenseCreated struct {
	License *CodeMashLicense `json:"license,omitempty"`
}

// ProjectCreated DTO.
type ProjectCreated struct {
	Id                    *ProjectId       `json:"id,omitempty"`
	Name                  *ProjectName     `json:"name,omitempty"`
	DatabaseIntegrationId *IntegrationId   `json:"databaseIntegrationId,omitempty"`
	PrimaryRegion         *ProjectRegion   `json:"primaryRegion,omitempty"`
	AdditionalRegions     []*ProjectRegion `json:"additionalRegions,omitempty"`
	Description           string           `json:"description,omitempty"`
	IsProvisioning        bool             `json:"isProvisioning,omitempty"`
}

// ProjectActivated DTO.
type ProjectActivated struct {
}

// ProjectSuspendedByLicense DTO.
type ProjectSuspendedByLicense struct {
}

// ProjectResumedFromLicenseSuspension DTO.
type ProjectResumedFromLicenseSuspension struct {
}

// ProjectDisabled DTO.
type ProjectDisabled struct {
}

// ProjectDeleted DTO.
type ProjectDeleted struct {
}

// ProjectNameChanged DTO.
type ProjectNameChanged struct {
	ProjectName *ProjectName `json:"projectName,omitempty"`
}

// ProjectDescriptionChanged DTO.
type ProjectDescriptionChanged struct {
	Description string `json:"description,omitempty"`
}

// ProjectMarketingUrlChanged DTO.
type ProjectMarketingUrlChanged struct {
	Url *DomainUrl `json:"url,omitempty"`
}

// ProjectAdminUrlChanged DTO.
type ProjectAdminUrlChanged struct {
	Url *DomainUrl `json:"url,omitempty"`
}

// ProjectLegalDocumentsChanged DTO.
type ProjectLegalDocumentsChanged struct {
	Documents *ProjectLegalDocuments `json:"documents,omitempty"`
}

// ProjectExposeLegalToAdminPortalChanged DTO.
type ProjectExposeLegalToAdminPortalChanged struct {
	Exposed bool `json:"exposed,omitempty"`
}

// ProjectAdminPortalServiceUserAssigned DTO.
type ProjectAdminPortalServiceUserAssigned struct {
	ServiceUserId *AuthId `json:"serviceUserId,omitempty"`
}

// ProjectAllowedOriginsChanged DTO.
type ProjectAllowedOriginsChanged struct {
	Origins []*DomainUrl `json:"origins,omitempty"`
}

// ProjectEnvironmentCreated DTO.
type ProjectEnvironmentCreated struct {
	Env   *Env               `json:"env,omitempty"`
	Ranks map[string]float64 `json:"ranks,omitempty"`
}

// ProjectEnvironmentDeleted DTO.
type ProjectEnvironmentDeleted struct {
	Env   *Env               `json:"env,omitempty"`
	Ranks map[string]float64 `json:"ranks,omitempty"`
}

// ProjectEnvironmentRanksChanged DTO.
type ProjectEnvironmentRanksChanged struct {
	Ranks map[string]float64 `json:"ranks,omitempty"`
}

// ProjectDefaultLanguageChanged DTO.
type ProjectDefaultLanguageChanged struct {
	Language *Language `json:"language,omitempty"`
}

// ProjectLanguagesChanged DTO.
type ProjectLanguagesChanged struct {
	Languages []*Language `json:"languages,omitempty"`
}

// ProjectLogoChanged DTO.
type ProjectLogoChanged struct {
	Logo *ProjectLogo `json:"logo,omitempty"`
}

// ProjectIconChanged DTO.
type ProjectIconChanged struct {
	Icon *ProjectIcon `json:"icon,omitempty"`
}

// ProjectMainColorChanged DTO.
type ProjectMainColorChanged struct {
	Color *BrandColor `json:"color,omitempty"`
}

// ProjectAccentColorChanged DTO.
type ProjectAccentColorChanged struct {
	Color *BrandColor `json:"color,omitempty"`
}

// ProjectRegionsChanged DTO.
type ProjectRegionsChanged struct {
	PrimaryRegion     *ProjectRegion   `json:"primaryRegion,omitempty"`
	AdditionalRegions []*ProjectRegion `json:"additionalRegions,omitempty"`
}

// ProjectCommunicationSet DTO.
type ProjectCommunicationSet struct {
	ProjectCommunication *ProjectCommunication `json:"projectCommunication,omitempty"`
}

// ProjectTimeZoneChanged DTO.
type ProjectTimeZoneChanged struct {
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

// ProjectPaymentZonesChanged DTO.
type ProjectPaymentZonesChanged struct {
	PaymentZones []*TimeZone `json:"paymentZones,omitempty"`
}

// ProjectCommunicationGroupSaved DTO.
type ProjectCommunicationGroupSaved struct {
	Group         *GroupDefinition     `json:"group,omitempty"`
	Channel       CommunicationChannel `json:"channel,omitempty"`
	OriginChannel CommunicationChannel `json:"originChannel,omitempty"`
}

// ProjectCommunicationTagFromGroupDeleted DTO.
type ProjectCommunicationTagFromGroupDeleted struct {
	GroupTag   *Tag `json:"groupTag,omitempty"`
	RemovedTag *Tag `json:"removedTag,omitempty"`
}

// ProjectCommunicationGroupDeleted DTO.
type ProjectCommunicationGroupDeleted struct {
	GroupTag *Tag `json:"groupTag,omitempty"`
}

// ProjectCommunicationTagSaved DTO.
type ProjectCommunicationTagSaved struct {
	Tag      *TagDefinition       `json:"tag,omitempty"`
	GroupTag *Tag                 `json:"groupTag,omitempty"`
	Channel  CommunicationChannel `json:"channel,omitempty"`
}

// ProjectCommunicationTagDeleted DTO.
type ProjectCommunicationTagDeleted struct {
	Tag *Tag `json:"tag,omitempty"`
}

// CustomerCreated DTO.
type CustomerCreated struct {
	PaymentCustomerRef *PaymentCustomerRef `json:"paymentCustomerRef,omitempty"`
}

// SubscriptionChanged DTO.
type SubscriptionChanged struct {
	Subscription *CodeMashManagedServiceSubscription `json:"subscription,omitempty"`
}

// SubscriptionCanceled DTO.
type SubscriptionCanceled struct {
	PaymentCustomerRef *PaymentCustomerRef `json:"paymentCustomerRef,omitempty"`
	SubscriptionId     string              `json:"subscriptionId,omitempty"`
}

// AccountTeamPolicyCreated DTO.
type AccountTeamPolicyCreated struct {
	Policy *MembershipPolicy `json:"policy,omitempty"`
}

// AccountTeamPolicyUpdated DTO.
type AccountTeamPolicyUpdated struct {
	Policy *MembershipPolicy `json:"policy,omitempty"`
}

// AccountTeamPolicyDeleted DTO.
type AccountTeamPolicyDeleted struct {
	PolicyId *PolicyId `json:"policyId,omitempty"`
}

// AccountTeamRoleCreated DTO.
type AccountTeamRoleCreated struct {
	Role *MembershipRole `json:"role,omitempty"`
}

// AccountTeamRoleUpdated DTO.
type AccountTeamRoleUpdated struct {
	Role *MembershipRole `json:"role,omitempty"`
}

// AccountTeamRoleDeleted DTO.
type AccountTeamRoleDeleted struct {
	RoleId *RoleId `json:"roleId,omitempty"`
}

// AtlasUsageRecorded DTO.
type AtlasUsageRecorded struct {
	Record *AtlasUsageRecord `json:"record,omitempty"`
}

// UsageBillingIngestionFailed DTO.
type UsageBillingIngestionFailed struct {
	Failure *UsageIngestionFailure `json:"failure,omitempty"`
}

// DisableMembership DTO.
type DisableMembership struct {
	CodeMashRequestBase
}

// EnableMembership DTO.
type EnableMembership struct {
	CodeMashRequestBase
}

// IssueServiceUserApiKeyRequest DTO.
type IssueServiceUserApiKeyRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Scopes                []string `json:"scopes,omitempty"`
	ExpiresInDays         float64  `json:"expiresInDays,omitempty"`
	Notes                 string   `json:"notes,omitempty"`
}

// ListServiceUserApiKeysRequest DTO.
type ListServiceUserApiKeysRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DeleteServiceUserApiKeyRequest DTO.
type DeleteServiceUserApiKeyRequest struct {
	CodeMashRequestBase
	Id    string  `json:"id,omitempty"`
	KeyId float64 `json:"keyId,omitempty"`
}

// DeleteMembershipTrigger DTO.
type DeleteMembershipTrigger struct {
	DeleteTrigger
}

// DisableMembershipTrigger DTO.
type DisableMembershipTrigger struct {
	DisableTrigger
}

// EnableMembershipTrigger DTO.
type EnableMembershipTrigger struct {
	EnableTrigger
}

// GetMembershipTrigger DTO.
type GetMembershipTrigger struct {
	GetTrigger
}

// GetMembershipTriggers DTO.
type GetMembershipTriggers struct {
	GetTriggers
}

// SaveMembershipTrigger DTO.
type SaveMembershipTrigger struct {
	SaveTrigger
}

// CreateRole DTO.
type CreateRole struct {
	CodeMashRequestBase
	RoleName    string   `json:"roleName,omitempty"`
	Description string   `json:"description,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

// DeleteRole DTO.
type DeleteRole struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetRole DTO.
type GetRole struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetRoles DTO.
type GetRoles struct {
	CodeMashRequestBase
}

// UpdateRolePolicies DTO.
type UpdateRolePolicies struct {
	CodeMashRequestBase
	Id          string   `json:"id,omitempty"`
	RoleName    string   `json:"roleName,omitempty"`
	Description string   `json:"description,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

// CreatePolicy DTO.
type CreatePolicy struct {
	CodeMashRequestBase
	PolicyName         string `json:"policyName,omitempty"`
	Description        string `json:"description,omitempty"`
	PolicyDocumentJson string `json:"policyDocumentJson,omitempty"`
}

// DeletePolicy DTO.
type DeletePolicy struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPolicy DTO.
type GetPolicy struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPolicies DTO.
type GetPolicies struct {
	CodeMashRequestBase
}

// UpdatePolicy DTO.
type UpdatePolicy struct {
	CodeMashRequestBase
	Id                 string `json:"id,omitempty"`
	PolicyName         string `json:"policyName,omitempty"`
	Description        string `json:"description,omitempty"`
	PolicyDocumentJson string `json:"policyDocumentJson,omitempty"`
}

// GetPasskeySettings DTO.
type GetPasskeySettings struct {
	CodeMashRequestBase
}

// SavePasskeySettings DTO.
type SavePasskeySettings struct {
	CodeMashRequestBase
	Enabled                       bool    `json:"enabled,omitempty"`
	CodeTtlMinutes                float64 `json:"codeTtlMinutes,omitempty"`
	MaxCredentialsPerUser         float64 `json:"maxCredentialsPerUser,omitempty"`
	RecoveryCodeCount             float64 `json:"recoveryCodeCount,omitempty"`
	GenerateRecoveryCodesAtSignup bool    `json:"generateRecoveryCodesAtSignup,omitempty"`
	AuthenticatorAttachment       string  `json:"authenticatorAttachment,omitempty"`
	AllowMagicLinkRecovery        bool    `json:"allowMagicLinkRecovery,omitempty"`
	RefreshTokenTtlDays           float64 `json:"refreshTokenTtlDays,omitempty"`
	RpId                          string  `json:"rpId,omitempty"`
}

// DeleteMembershipIntegrationRequest DTO.
type DeleteMembershipIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableMembershipIntegrationRequest DTO.
type DisableMembershipIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableMembershipIntegrationRequest DTO.
type EnableMembershipIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetMembershipIntegration DTO.
type GetMembershipIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetMembershipIntegrations DTO.
type GetMembershipIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveMembershipIntegration DTO.
type SaveMembershipIntegration struct {
	CodeMashRequestBase
	Integration *MembershipIntegrationRequest `json:"integration,omitempty"`
}

// SetMembershipIntegrationAsDefaultRequest DTO.
type SetMembershipIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetAuthorizationSettings DTO.
type GetAuthorizationSettings struct {
	CodeMashRequestBase
}

// UpdateAuthorizationSettings DTO.
type UpdateAuthorizationSettings struct {
	CodeMashRequestBase
	Setting                     string   `json:"setting,omitempty"`
	DefaultRoles                []string `json:"defaultRoles,omitempty"`
	AllowedRegistrationRoles    []string `json:"allowedRegistrationRoles,omitempty"`
	AllowGuestUsers             bool     `json:"allowGuestUsers,omitempty"`
	GuestCleanupPeriodDays      float64  `json:"guestCleanupPeriodDays,omitempty"`
	UserRegistersAsRole         string   `json:"userRegistersAsRole,omitempty"`
	GuestRegistersAsRole        string   `json:"guestRegistersAsRole,omitempty"`
	AllowedRegisterRoles        []string `json:"allowedRegisterRoles,omitempty"`
	NeedVerification            bool     `json:"needVerification,omitempty"`
	VerificationEmailTemplate   string   `json:"verificationEmailTemplate,omitempty"`
	DeactivationEmailTemplate   string   `json:"deactivationEmailTemplate,omitempty"`
	AllowInviteUsers            bool     `json:"allowInviteUsers,omitempty"`
	AllowDeactivateUsers        bool     `json:"allowDeactivateUsers,omitempty"`
	InviteUserEmailTemplate     string   `json:"inviteUserEmailTemplate,omitempty"`
	InvitationExpiration        float64  `json:"invitationExpiration,omitempty"`
	EmailVerificationExpiration float64  `json:"emailVerificationExpiration,omitempty"`
	DeactivationExpiration      float64  `json:"deactivationExpiration,omitempty"`
	DefaultSubscribeToNews      bool     `json:"defaultSubscribeToNews,omitempty"`
	MinLength                   float64  `json:"minLength,omitempty"`
	MaxLength                   float64  `json:"maxLength,omitempty"`
	MinNumbers                  float64  `json:"minNumbers,omitempty"`
	MaxNumbers                  float64  `json:"maxNumbers,omitempty"`
	MinUpper                    float64  `json:"minUpper,omitempty"`
	MaxUpper                    float64  `json:"maxUpper,omitempty"`
	MinLower                    float64  `json:"minLower,omitempty"`
	MaxLower                    float64  `json:"maxLower,omitempty"`
	MinSpecial                  float64  `json:"minSpecial,omitempty"`
	MaxSpecial                  float64  `json:"maxSpecial,omitempty"`
	AllowedSpecial              string   `json:"allowedSpecial,omitempty"`
}

// UpdatePasswordComplexity DTO.
type UpdatePasswordComplexity struct {
	CodeMashRequestBase
	MinLength      float64 `json:"minLength,omitempty"`
	MaxLength      float64 `json:"maxLength,omitempty"`
	MinNumbers     float64 `json:"minNumbers,omitempty"`
	MaxNumbers     float64 `json:"maxNumbers,omitempty"`
	MinUpper       float64 `json:"minUpper,omitempty"`
	MaxUpper       float64 `json:"maxUpper,omitempty"`
	MinLower       float64 `json:"minLower,omitempty"`
	MaxLower       float64 `json:"maxLower,omitempty"`
	MinSpecial     float64 `json:"minSpecial,omitempty"`
	MaxSpecial     float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial string  `json:"allowedSpecial,omitempty"`
}

// GetAuthenticationSettings DTO.
type GetAuthenticationSettings struct {
	CodeMashRequestBase
}

// UpdateAuthenticationSettings DTO.
type UpdateAuthenticationSettings struct {
	CodeMashRequestBase
	LogoutUrl      string                        `json:"logoutUrl,omitempty"`
	AllowUsernames bool                          `json:"allowUsernames,omitempty"`
	Modes          []*CredentialsSettingsModeDto `json:"modes,omitempty"`
}

// MembershipIntegrationSaved DTO.
type MembershipIntegrationSaved struct {
	Integration *MembershipIntegration `json:"integration,omitempty"`
}

// MembershipIntegrationTested DTO.
type MembershipIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// MembershipIntegrationRenamed DTO.
type MembershipIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// MembershipIntegrationDeleted DTO.
type MembershipIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MembershipIntegrationSetAsDefault DTO.
type MembershipIntegrationSetAsDefault struct {
	Id *IntegrationId `json:"id,omitempty"`
}

// MembershipIntegrationEnabled DTO.
type MembershipIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MembershipIntegrationDisabled DTO.
type MembershipIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MembershipEstablished DTO.
type MembershipEstablished struct {
}

// MembershipEnabled DTO.
type MembershipEnabled struct {
}

// MembershipDisabled DTO.
type MembershipDisabled struct {
}

// SetUserRegistersAsRole DTO.
type SetUserRegistersAsRole struct {
	ProjectId *ProjectId `json:"projectId,omitempty"`
	Role      *RoleName  `json:"role,omitempty"`
}

// PolicyCreated DTO.
type PolicyCreated struct {
	Policy *MembershipPolicy `json:"policy,omitempty"`
}

// PolicyUpdated DTO.
type PolicyUpdated struct {
	Policy *MembershipPolicy `json:"policy,omitempty"`
}

// PolicyDeleted DTO.
type PolicyDeleted struct {
	PolicyId *PolicyId `json:"policyId,omitempty"`
}

// RoleCreated DTO.
type RoleCreated struct {
	Role *MembershipRole `json:"role,omitempty"`
}

// RoleUpdated DTO.
type RoleUpdated struct {
	Role *MembershipRole `json:"role,omitempty"`
}

// RoleDeleted DTO.
type RoleDeleted struct {
	RoleId *RoleId `json:"roleId,omitempty"`
}

// MembershipTriggerSaved DTO.
type MembershipTriggerSaved struct {
	Trigger *MembershipTrigger `json:"trigger,omitempty"`
}

// MembershipTriggerMirrored DTO.
type MembershipTriggerMirrored struct {
	Trigger *Trigger `json:"trigger,omitempty"`
}

// MembershipTriggerEnabled DTO.
type MembershipTriggerEnabled struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// MembershipTriggerDisabled DTO.
type MembershipTriggerDisabled struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// MembershipTriggerDeleted DTO.
type MembershipTriggerDeleted struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// DisableDatabase DTO.
type DisableDatabase struct {
	CodeMashRequestBase
}

// EnableDatabase DTO.
type EnableDatabase struct {
	CodeMashRequestBase
}

// DeleteSchemaTrigger DTO.
type DeleteSchemaTrigger struct {
	DeleteTrigger
}

// DisableSchemaTrigger DTO.
type DisableSchemaTrigger struct {
	DisableTrigger
}

// EnableSchemaTrigger DTO.
type EnableSchemaTrigger struct {
	EnableTrigger
}

// GetSchemaTrigger DTO.
type GetSchemaTrigger struct {
	GetTrigger
}

// GetSchemaTriggers DTO.
type GetSchemaTriggers struct {
	GetTriggers
}

// SaveSchemaTrigger DTO.
type SaveSchemaTrigger struct {
	SaveTrigger
}

// DeleteDatabaseTaxonomyRequest DTO.
type DeleteDatabaseTaxonomyRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseTaxonomy DTO.
type GetDatabaseTaxonomy struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseTaxonomies DTO.
type GetDatabaseTaxonomies struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// GetDatabaseTaxonomyTreeRequest DTO.
type GetDatabaseTaxonomyTreeRequest struct {
	CodeMashRequestBase
	IncludeTerms          bool   `json:"includeTerms,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// SaveDatabaseTaxonomyRequest DTO.
type SaveDatabaseTaxonomyRequest struct {
	CodeMashRequestBase
	ViewId                string   `json:"viewId,omitempty"`
	TaxonomyName          string   `json:"taxonomyName,omitempty"`
	Description           string   `json:"description,omitempty"`
	TermsMetaDataSchema   string   `json:"termsMetaDataSchema,omitempty"`
	TermsMetaVisualSchema string   `json:"termsMetaVisualSchema,omitempty"`
	ParentId              string   `json:"parentId,omitempty"`
	Dependencies          []string `json:"dependencies,omitempty"`
}

// DeleteDatabaseTaxonomyTermRequest DTO.
type DeleteDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	TaxonomyId            string `json:"taxonomyId,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// DeleteManyDatabaseTaxonomyTermsRequest DTO.
type DeleteManyDatabaseTaxonomyTermsRequest struct {
	CodeMashRequestBase
	TaxonomyId            string `json:"taxonomyId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
}

// GetDatabaseTaxonomyTermRequest DTO.
type GetDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	TaxonomyId            string `json:"taxonomyId,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetDatabaseMergedTermTreeRequest DTO.
type GetDatabaseMergedTermTreeRequest struct {
	CodeMashRequestBase
	TaxonomyName          string `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetDatabaseTaxonomyTermTreeRequest DTO.
type GetDatabaseTaxonomyTermTreeRequest struct {
	CodeMashRequestBase
	TaxonomyName          string  `json:"taxonomyName,omitempty"`
	RootTermId            string  `json:"rootTermId,omitempty"`
	Depth                 float64 `json:"depth,omitempty"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
}

// SaveDatabaseTaxonomyTermRequest DTO.
type SaveDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	TaxonomyId            string `json:"taxonomyId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Document              string `json:"document,omitempty"`
}

// UpdateDatabaseTaxonomyTermRequest DTO.
type UpdateDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	TaxonomyId            string `json:"taxonomyId,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Update                string `json:"update,omitempty"`
}

// ApplyDatabaseSchemaBundleRequest DTO.
type ApplyDatabaseSchemaBundleRequest struct {
	CodeMashRequestBase
	Entities     string `json:"entities,omitempty"`
	BundleJson   string `json:"bundleJson,omitempty"`
	Tier         string `json:"tier,omitempty"`
	Translatable bool   `json:"translatable,omitempty"`
	Publish      bool   `json:"publish,omitempty"`
}

// DeleteDatabaseSchemaRequest DTO.
type DeleteDatabaseSchemaRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DiscardDatabaseSchemaDraftRequest DTO.
type DiscardDatabaseSchemaDraftRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseSchema DTO.
type GetDatabaseSchema struct {
	CodeMashRequestBase
	Id      string  `json:"id,omitempty"`
	Version float64 `json:"version,omitempty"`
}

// GetDatabaseSchemas DTO.
type GetDatabaseSchemas struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// GetDatabaseSchemaDraft DTO.
type GetDatabaseSchemaDraft struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseSchemaListSettings DTO.
type GetDatabaseSchemaListSettings struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseSchemaVersionDiff DTO.
type GetDatabaseSchemaVersionDiff struct {
	CodeMashRequestBase
	Id          string  `json:"id,omitempty"`
	FromVersion float64 `json:"fromVersion,omitempty"`
	ToVersion   float64 `json:"toVersion,omitempty"`
}

// GetDatabaseSchemaVersions DTO.
type GetDatabaseSchemaVersions struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// PublishDatabaseSchemaRequest DTO.
type PublishDatabaseSchemaRequest struct {
	CodeMashRequestBase
	Id        string `json:"id,omitempty"`
	Confirmed bool   `json:"confirmed,omitempty"`
}

// RenameDatabaseSchemaRequest DTO.
type RenameDatabaseSchemaRequest struct {
	CodeMashRequestBase
	Id               string `json:"id,omitempty"`
	Title            string `json:"title,omitempty"`
	RenameUniqueName bool   `json:"renameUniqueName,omitempty"`
}

// SaveDatabaseSchemaRequest DTO.
type SaveDatabaseSchemaRequest struct {
	CodeMashRequestBase
	ViewId       string             `json:"viewId,omitempty"`
	SchemaName   string             `json:"schemaName,omitempty"`
	DataSchema   string             `json:"dataSchema,omitempty"`
	VisualSchema string             `json:"visualSchema,omitempty"`
	Settings     *SchemaSettingsDto `json:"settings,omitempty"`
}

// UpdateDatabaseSchemaDraftRequest DTO.
type UpdateDatabaseSchemaDraftRequest struct {
	CodeMashRequestBase
	Id           string `json:"id,omitempty"`
	DataSchema   string `json:"dataSchema,omitempty"`
	VisualSchema string `json:"visualSchema,omitempty"`
}

// UpdateDatabaseSchemaListSettingsRequest DTO.
type UpdateDatabaseSchemaListSettingsRequest struct {
	CodeMashRequestBase
	Id       string                 `json:"id,omitempty"`
	Settings *SchemaListSettingsDto `json:"settings,omitempty"`
}

// UpdateDatabaseSchemaSettingsRequest DTO.
type UpdateDatabaseSchemaSettingsRequest struct {
	CodeMashRequestBase
	Id       string             `json:"id,omitempty"`
	Settings *SchemaSettingsDto `json:"settings,omitempty"`
}

// AggregateRecords DTO.
type AggregateRecords struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Pipeline              string `json:"pipeline,omitempty"`
}

// ChangeRecordResponsibility DTO.
type ChangeRecordResponsibility struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	NewResponsibleUserId  string `json:"newResponsibleUserId,omitempty"`
}

// CountRecords DTO.
type CountRecords struct {
	CodeMashRequestBase
	CollectionName        string  `json:"collectionName,omitempty"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	Filter                string  `json:"filter,omitempty"`
	SchemaVersion         float64 `json:"schemaVersion,omitempty"`
}

// DeleteManyRecords DTO.
type DeleteManyRecords struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
}

// DeleteRecord DTO.
type DeleteRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// DistinctRecordValues DTO.
type DistinctRecordValues struct {
	CodeMashRequestBase
	CollectionName        string  `json:"collectionName,omitempty"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	Field                 string  `json:"field,omitempty"`
	Filter                string  `json:"filter,omitempty"`
	SchemaVersion         float64 `json:"schemaVersion,omitempty"`
}

// ExecuteRecordsAggregate DTO.
type ExecuteRecordsAggregate struct {
	CodeMashRequestBase
	CollectionName        string            `json:"collectionName,omitempty"`
	AggregateId           string            `json:"aggregateId,omitempty"`
	DatabaseIntegrationId string            `json:"databaseIntegrationId,omitempty"`
	Tokens                map[string]string `json:"tokens,omitempty"`
}

// FindRecords DTO.
type FindRecords struct {
	CodeMashListPaginationRequestBase
	CollectionName        string      `json:"collectionName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	ContactId             string      `json:"contactId,omitempty"`
	SchemaVersion         float64     `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	SortBy                string      `json:"sortBy,omitempty"`
	SortOrder             float64     `json:"sortOrder,omitempty"`
}

// FindOneRecord DTO.
type FindOneRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetCollectionIndexes DTO.
type GetCollectionIndexes struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// InsertManyRecords DTO.
type InsertManyRecords struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Documents             string `json:"documents,omitempty"`
}

// InsertRecord DTO.
type InsertRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Document              string `json:"document,omitempty"`
}

// ReplaceRecord DTO.
type ReplaceRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Replacement           string `json:"replacement,omitempty"`
}

// SeedCollectionRecords DTO.
type SeedCollectionRecords struct {
	CodeMashRequestBase
	Mode                  string `json:"mode,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Collections           string `json:"collections,omitempty"`
}

// UpdateManyRecords DTO.
type UpdateManyRecords struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
	Update                string `json:"update,omitempty"`
}

// UpdateOneRecord DTO.
type UpdateOneRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Update                string `json:"update,omitempty"`
}

// DeleteDatabaseIntegrationRequest DTO.
type DeleteDatabaseIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableDatabaseIntegrationRequest DTO.
type DisableDatabaseIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableDatabaseIntegrationRequest DTO.
type EnableDatabaseIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseIntegration DTO.
type GetDatabaseIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseIntegrations DTO.
type GetDatabaseIntegrations struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// GetAllowedFlexTiers DTO.
type GetAllowedFlexTiers struct {
	CodeMashRequestBase
}

// RevealManagedFlexConnectionString DTO.
type RevealManagedFlexConnectionString struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// SaveDatabaseIntegration DTO.
type SaveDatabaseIntegration struct {
	CodeMashRequestBase
	Integration *DatabaseIntegrationRequest `json:"integration,omitempty"`
}

// SetDatabaseIntegrationAsDefaultRequest DTO.
type SetDatabaseIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// TestDatabaseIntegration DTO.
type TestDatabaseIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// CreateCollectionImport DTO.
type CreateCollectionImport struct {
	CodeMashRequestBase
	File                  *FileResourceRefDto       `json:"file,omitempty"`
	SchemaId              string                    `json:"schemaId,omitempty"`
	CollectionName        string                    `json:"collectionName,omitempty"`
	DatabaseIntegrationId string                    `json:"databaseIntegrationId,omitempty"`
	Delimiter             string                    `json:"delimiter,omitempty"`
	HasHeader             bool                      `json:"hasHeader,omitempty"`
	Mapping               []*ImportColumnMappingDto `json:"mapping,omitempty"`
}

// DeleteCollectionImportRequest DTO.
type DeleteCollectionImportRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetCollectionImport DTO.
type GetCollectionImport struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetCollectionImports DTO.
type GetCollectionImports struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// RequestImportUploadUrlRequest DTO.
type RequestImportUploadUrlRequest struct {
	CodeMashRequestBase
	FileAccountId string `json:"fileAccountId,omitempty"`
	FileName      string `json:"fileName,omitempty"`
}

// AnalyzeImportFileRequest DTO.
type AnalyzeImportFileRequest struct {
	CodeMashRequestBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	Delimiter string              `json:"delimiter,omitempty"`
	HasHeader bool                `json:"hasHeader,omitempty"`
}

// DeleteDatabaseAggregateRequest DTO.
type DeleteDatabaseAggregateRequest struct {
	CodeMashRequestBase
	Id       string `json:"id,omitempty"`
	SchemaId string `json:"schemaId,omitempty"`
}

// GetDatabaseAggregate DTO.
type GetDatabaseAggregate struct {
	CodeMashRequestBase
	Id       string `json:"id,omitempty"`
	SchemaId string `json:"schemaId,omitempty"`
}

// GetDatabaseAggregates DTO.
type GetDatabaseAggregates struct {
	CodeMashListPaginationRequestBase
	SchemaId   string      `json:"schemaId,omitempty"`
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// SaveDatabaseAggregateRequest DTO.
type SaveDatabaseAggregateRequest struct {
	CodeMashRequestBase
	ViewId      string `json:"viewId,omitempty"`
	SchemaId    string `json:"schemaId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Description string `json:"description,omitempty"`
	Pipeline    string `json:"pipeline,omitempty"`
}

// TestDatabaseAggregateRequest DTO.
type TestDatabaseAggregateRequest struct {
	CodeMashRequestBase
	DatabaseIntegrationId string            `json:"databaseIntegrationId,omitempty"`
	CollectionName        string            `json:"collectionName,omitempty"`
	Pipeline              string            `json:"pipeline,omitempty"`
	Tokens                map[string]string `json:"tokens,omitempty"`
}

// MongoDbAggregateCreated DTO.
type MongoDbAggregateCreated struct {
	Aggregate *MongoDbAggregate `json:"aggregate,omitempty"`
}

// MongoDbAggregateUpdated DTO.
type MongoDbAggregateUpdated struct {
	Aggregate *MongoDbAggregate `json:"aggregate,omitempty"`
}

// MongoDbAggregateDeleted DTO.
type MongoDbAggregateDeleted struct {
	SchemaId *SchemaId           `json:"schemaId,omitempty"`
	Id       *MongoDbAggregateId `json:"id,omitempty"`
}

// DatabaseEstablished DTO.
type DatabaseEstablished struct {
}

// DatabaseEnabled DTO.
type DatabaseEnabled struct {
}

// DatabaseDisabled DTO.
type DatabaseDisabled struct {
}

// DatabaseIntegrationSaved DTO.
type DatabaseIntegrationSaved struct {
	Integration *DatabaseIntegration `json:"integration,omitempty"`
}

// DatabaseIntegrationTested DTO.
type DatabaseIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// DatabaseIntegrationRenamed DTO.
type DatabaseIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// DatabaseIntegrationSetAsDefault DTO.
type DatabaseIntegrationSetAsDefault struct {
	Env *Env           `json:"env,omitempty"`
	Id  *IntegrationId `json:"id,omitempty"`
}

// DatabaseIntegrationDeleted DTO.
type DatabaseIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// DatabaseIntegrationEnabled DTO.
type DatabaseIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// DatabaseIntegrationDisabled DTO.
type DatabaseIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// DatabaseIntegrationProvisioningStarted DTO.
type DatabaseIntegrationProvisioningStarted struct {
	IntegrationId    *IntegrationId `json:"integrationId,omitempty"`
	AtlasProjectId   string         `json:"atlasProjectId,omitempty"`
	AtlasClusterName string         `json:"atlasClusterName,omitempty"`
}

// DatabaseIntegrationProvisioningCompleted DTO.
type DatabaseIntegrationProvisioningCompleted struct {
	IntegrationId            *IntegrationId `json:"integrationId,omitempty"`
	ConnectionStringTemplate string         `json:"connectionStringTemplate,omitempty"`
}

// DatabaseIntegrationProvisioningFailed DTO.
type DatabaseIntegrationProvisioningFailed struct {
	IntegrationId *IntegrationId `json:"integrationId,omitempty"`
	Reason        string         `json:"reason,omitempty"`
	Retryable     bool           `json:"retryable,omitempty"`
}

// DatabaseIntegrationDeprovisioned DTO.
type DatabaseIntegrationDeprovisioned struct {
	IntegrationId    *IntegrationId `json:"integrationId,omitempty"`
	AtlasProjectId   string         `json:"atlasProjectId,omitempty"`
	AtlasClusterName string         `json:"atlasClusterName,omitempty"`
}

// ProjectStatusChanged DTO.
type ProjectStatusChanged struct {
	Status ProjectStatus `json:"status,omitempty"`
}

// SchemaCreated DTO.
type SchemaCreated struct {
	Schema *Schema `json:"schema,omitempty"`
}

// SchemaMirrored DTO.
type SchemaMirrored struct {
	Schema *Schema `json:"schema,omitempty"`
}

// SchemaDraftUpdated DTO.
type SchemaDraftUpdated struct {
	Id    *SchemaId    `json:"id,omitempty"`
	Draft *SchemaDraft `json:"draft,omitempty"`
	Env   *Env         `json:"env,omitempty"`
}

// SchemaDraftDiscarded DTO.
type SchemaDraftDiscarded struct {
	Id  *SchemaId `json:"id,omitempty"`
	Env *Env      `json:"env,omitempty"`
}

// SchemaVersionPublished DTO.
type SchemaVersionPublished struct {
	Id      *SchemaId               `json:"id,omitempty"`
	Version *PublishedSchemaVersion `json:"version,omitempty"`
	Diff    *SchemaDiff             `json:"diff,omitempty"`
	Env     *Env                    `json:"env,omitempty"`
}

// SchemaSettingsUpdated DTO.
type SchemaSettingsUpdated struct {
	Id       *SchemaId       `json:"id,omitempty"`
	Settings *SchemaSettings `json:"settings,omitempty"`
	Env      *Env            `json:"env,omitempty"`
}

// SchemaDeleted DTO.
type SchemaDeleted struct {
	Id  *SchemaId `json:"id,omitempty"`
	Env *Env      `json:"env,omitempty"`
}

// SchemaRenamed DTO.
type SchemaRenamed struct {
	SchemaId         *SchemaId   `json:"schemaId,omitempty"`
	NewName          *SchemaName `json:"newName,omitempty"`
	RenameUniqueName bool        `json:"renameUniqueName,omitempty"`
	Env              *Env        `json:"env,omitempty"`
}

// SchemaDataCleared DTO.
type SchemaDataCleared struct {
	Id           *SchemaId        `json:"id,omitempty"`
	Integrations []*IntegrationId `json:"integrations,omitempty"`
	Env          *Env             `json:"env,omitempty"`
}

// TaxonomyCreated DTO.
type TaxonomyCreated struct {
	Taxonomy *Taxonomy `json:"taxonomy,omitempty"`
}

// TaxonomyUpdated DTO.
type TaxonomyUpdated struct {
	Taxonomy *Taxonomy `json:"taxonomy,omitempty"`
}

// TaxonomyDeleted DTO.
type TaxonomyDeleted struct {
	TaxonomyId *TaxonomyId `json:"taxonomyId,omitempty"`
}

// TaxonomyDataCleared DTO.
type TaxonomyDataCleared struct {
	TaxonomyId   *TaxonomyId      `json:"taxonomyId,omitempty"`
	Integrations []*IntegrationId `json:"integrations,omitempty"`
}

// SchemaTriggerSaved DTO.
type SchemaTriggerSaved struct {
	Trigger *SchemaTrigger `json:"trigger,omitempty"`
}

// DatabaseTriggerMirrored DTO.
type DatabaseTriggerMirrored struct {
	Trigger *Trigger `json:"trigger,omitempty"`
}

// SchemaTriggerEnabled DTO.
type SchemaTriggerEnabled struct {
	TriggerByIdEventBase
	SchemaId *SchemaId `json:"schemaId,omitempty"`
	Env      *Env      `json:"env,omitempty"`
}

// SchemaTriggerDisabled DTO.
type SchemaTriggerDisabled struct {
	TriggerByIdEventBase
	SchemaId *SchemaId `json:"schemaId,omitempty"`
	Env      *Env      `json:"env,omitempty"`
}

// SchemaTriggerDeleted DTO.
type SchemaTriggerDeleted struct {
	TriggerByIdEventBase
	SchemaId *SchemaId `json:"schemaId,omitempty"`
	Env      *Env      `json:"env,omitempty"`
}

// ProcessCollectionImport DTO.
type ProcessCollectionImport struct {
	ImportId              string `json:"importId,omitempty"`
	ProjectId             string `json:"projectId,omitempty"`
	AccountId             string `json:"accountId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Env                   string `json:"env,omitempty"`
}

// RecordInserted DTO.
type RecordInserted struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	Id                    string         `json:"id,omitempty"`
	Document              map[string]any `json:"document,omitempty"`
}

// RecordUpdated DTO.
type RecordUpdated struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	Id                    string         `json:"id,omitempty"`
	From                  map[string]any `json:"from,omitempty"`
	To                    map[string]any `json:"to,omitempty"`
}

// RecordDeleted DTO.
type RecordDeleted struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	Id                    string         `json:"id,omitempty"`
	Document              map[string]any `json:"document,omitempty"`
}

// RecordReplaced DTO.
type RecordReplaced struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	Id                    string         `json:"id,omitempty"`
	From                  map[string]any `json:"from,omitempty"`
	To                    map[string]any `json:"to,omitempty"`
}

// RecordResponsibilityChanged DTO.
type RecordResponsibilityChanged struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	Id                    string         `json:"id,omitempty"`
	FromOwner             *AuthId        `json:"fromOwner,omitempty"`
	ToOwner               *AuthId        `json:"toOwner,omitempty"`
}

// RecordsInserted DTO.
type RecordsInserted struct {
	ProjectId             *ProjectId       `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId   `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName      `json:"schemaName,omitempty"`
	Ids                   []string         `json:"ids,omitempty"`
	Documents             []map[string]any `json:"documents,omitempty"`
}

// RecordsUpdated DTO.
type RecordsUpdated struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	MatchedCount          float64        `json:"matchedCount,omitempty"`
	ModifiedCount         float64        `json:"modifiedCount,omitempty"`
	Update                map[string]any `json:"update,omitempty"`
}

// RecordsDeleted DTO.
type RecordsDeleted struct {
	ProjectId             *ProjectId     `json:"projectId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
	SchemaName            *SchemaName    `json:"schemaName,omitempty"`
	DeletedCount          float64        `json:"deletedCount,omitempty"`
	Filter                map[string]any `json:"filter,omitempty"`
}

// EmailVerificationCodeRequested DTO.
type EmailVerificationCodeRequested struct {
	Email        string `json:"email,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	Code         string `json:"code,omitempty"`
	ExpiresAtUtc string `json:"expiresAtUtc,omitempty"`
}

// MagicLinkRequested DTO.
type MagicLinkRequested struct {
	Email        string `json:"email,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	Token        string `json:"token,omitempty"`
	ExpiresAtUtc string `json:"expiresAtUtc,omitempty"`
}

// PasswordResetRequested DTO.
type PasswordResetRequested struct {
	Email        string `json:"email,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	Token        string `json:"token,omitempty"`
	ExpiresAtUtc string `json:"expiresAtUtc,omitempty"`
}

// PasswordChanged DTO.
type PasswordChanged struct {
	Email     string `json:"email,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
}

// SseCallTriggered DTO.
type SseCallTriggered struct {
	ProjectId        *ProjectId        `json:"projectId,omitempty"`
	AccountId        *AccountId        `json:"accountId,omitempty"`
	TriggerId        *TriggerId        `json:"triggerId,omitempty"`
	TriggerType      TriggerType       `json:"triggerType,omitempty"`
	SourceEvent      string            `json:"sourceEvent,omitempty"`
	TargetUserAuthId string            `json:"targetUserAuthId,omitempty"`
	SchemaId         string            `json:"schemaId,omitempty"`
	TokenMappings    map[string]string `json:"tokenMappings,omitempty"`
	CorrelationId    string            `json:"correlationId,omitempty"`
}

// UserRegistered DTO.
type UserRegistered struct {
	Auth       *Auth   `json:"auth,omitempty"`
	LinkToUser *UserId `json:"linkToUser,omitempty"`
}

// UserCreated DTO.
type UserCreated struct {
	UserId    *UserId    `json:"userId,omitempty"`
	ProjectId *ProjectId `json:"projectId,omitempty"`
	AuthId    *AuthId    `json:"authId,omitempty"`
}

// UserUpdated DTO.
type UserUpdated struct {
	AuthId *AuthId          `json:"authId,omitempty"`
	From   *UserGeneralInfo `json:"from,omitempty"`
	To     *UserGeneralInfo `json:"to,omitempty"`
}

// UserBlocked DTO.
type UserBlocked struct {
	User   *UserGeneralInfo `json:"user,omitempty"`
	AuthId *AuthId          `json:"authId,omitempty"`
}

// UserUnblocked DTO.
type UserUnblocked struct {
	User   *UserGeneralInfo `json:"user,omitempty"`
	AuthId *AuthId          `json:"authId,omitempty"`
}

// UserInvited DTO.
type UserInvited struct {
	EmailAddress *EmailAddress `json:"emailAddress,omitempty"`
}

// UserVerified DTO.
type UserVerified struct {
	AuthId *AuthId          `json:"authId,omitempty"`
	User   *UserGeneralInfo `json:"user,omitempty"`
}

// UserDeleted DTO.
type UserDeleted struct {
	User   *UserGeneralInfo `json:"user,omitempty"`
	AuthId *AuthId          `json:"authId,omitempty"`
}

// DisableFiles DTO.
type DisableFiles struct {
	CodeMashRequestBase
}

// EnableFiles DTO.
type EnableFiles struct {
	CodeMashRequestBase
}

// DeleteFilesTrigger DTO.
type DeleteFilesTrigger struct {
	DeleteTrigger
}

// DisableFilesTrigger DTO.
type DisableFilesTrigger struct {
	DisableTrigger
}

// EnableFilesTrigger DTO.
type EnableFilesTrigger struct {
	EnableTrigger
}

// GetFilesTrigger DTO.
type GetFilesTrigger struct {
	GetTrigger
}

// GetFilesTriggers DTO.
type GetFilesTriggers struct {
	GetTriggers
}

// SaveFilesTrigger DTO.
type SaveFilesTrigger struct {
	SaveTrigger
}

// MakeFilePrivateRequest DTO.
type MakeFilePrivateRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// MakeFilePublicRequest DTO.
type MakeFilePublicRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// MakeFolderPrivateRequest DTO.
type MakeFolderPrivateRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// MakeFolderPublicRequest DTO.
type MakeFolderPublicRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// DeleteFilesIntegrationRequest DTO.
type DeleteFilesIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableFilesIntegrationRequest DTO.
type DisableFilesIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableFilesIntegrationRequest DTO.
type EnableFilesIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetFilesIntegration DTO.
type GetFilesIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetFilesIntegrations DTO.
type GetFilesIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveFilesIntegration DTO.
type SaveFilesIntegration struct {
	CodeMashRequestBase
	Integration *FilesIntegrationRequest `json:"integration,omitempty"`
}

// SetFilesIntegrationAsDefaultRequest DTO.
type SetFilesIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// TestFilesIntegration DTO.
type TestFilesIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// GetFile DTO.
type GetFile struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetFolderFiles DTO.
type GetFolderFiles struct {
	CodeMashListPaginationRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// FilesEstablished DTO.
type FilesEstablished struct {
}

// FilesEnabled DTO.
type FilesEnabled struct {
}

// FilesDisabled DTO.
type FilesDisabled struct {
}

// FilesIntegrationSaved DTO.
type FilesIntegrationSaved struct {
	Integration *FileIntegration `json:"integration,omitempty"`
}

// FilesIntegrationTested DTO.
type FilesIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// FilesIntegrationRenamed DTO.
type FilesIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// FilesIntegrationDeleted DTO.
type FilesIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// FilesIntegrationEnabled DTO.
type FilesIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// FilesIntegrationDisabled DTO.
type FilesIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// FilesIntegrationSetAsDefault DTO.
type FilesIntegrationSetAsDefault struct {
	Env *Env           `json:"env,omitempty"`
	Id  *IntegrationId `json:"id,omitempty"`
}

// FilesTriggerSaved DTO.
type FilesTriggerSaved struct {
	Trigger *FileTrigger `json:"trigger,omitempty"`
}

// FilesTriggerMirrored DTO.
type FilesTriggerMirrored struct {
	Trigger *Trigger `json:"trigger,omitempty"`
}

// FilesTriggerEnabled DTO.
type FilesTriggerEnabled struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// FilesTriggerDisabled DTO.
type FilesTriggerDisabled struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// FilesTriggerDeleted DTO.
type FilesTriggerDeleted struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// FileUploaded DTO.
type FileUploaded struct {
	ProjectId     *ProjectId       `json:"projectId,omitempty"`
	IntegrationId *IntegrationId   `json:"integrationId,omitempty"`
	FileRef       *FileResourceRef `json:"fileRef,omitempty"`
	Verified      bool             `json:"verified,omitempty"`
}

// FileDeleted DTO.
type FileDeleted struct {
	ProjectId     *ProjectId     `json:"projectId,omitempty"`
	IntegrationId *IntegrationId `json:"integrationId,omitempty"`
	Path          string         `json:"path,omitempty"`
}

// DisableEmail DTO.
type DisableEmail struct {
	CodeMashRequestBase
}

// GetEmailDisableDependencies DTO.
type GetEmailDisableDependencies struct {
	CodeMashRequestBase
}

// EnableEmail DTO.
type EnableEmail struct {
	CodeMashRequestBase
}

// SaveEmailValidationIntegration DTO.
type SaveEmailValidationIntegration struct {
	CodeMashRequestBase
	Integration *EmailValidationIntegrationRequest `json:"integration,omitempty"`
}

// TestEmailValidationIntegration DTO.
type TestEmailValidationIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// AttachFileToTemplateRequest DTO.
type AttachFileToTemplateRequest struct {
	CodeMashRequestBase
	Language   string              `json:"language,omitempty"`
	TemplateId string              `json:"templateId,omitempty"`
	FileRef    *FileResourceRefDto `json:"fileRef,omitempty"`
}

// CreateEmailTemplateRequest DTO.
type CreateEmailTemplateRequest struct {
	SaveEmailTemplate
}

// DeleteEmailTemplateRequest DTO.
type DeleteEmailTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailTemplate DTO.
type GetEmailTemplate struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailTemplates DTO.
type GetEmailTemplates struct {
	CodeMashListPaginationRequestBase
	ShowArchived bool   `json:"showArchived,omitempty"`
	TemplateId   string `json:"templateId,omitempty"`
}

// GetMjml DTO.
type GetMjml struct {
	CodeMashRequestBase
	Code         string             `json:"code,omitempty"`
	Tokens       []*TokenMappingDto `json:"tokens,omitempty"`
	IsForPreview bool               `json:"isForPreview,omitempty"`
}

// GetSystemEmailTemplate DTO.
type GetSystemEmailTemplate struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSystemEmailTemplates DTO.
type GetSystemEmailTemplates struct {
	CodeMashListPaginationRequestBase
	GroupTags            []string             `json:"groupTags,omitempty"`
	Themes               []string             `json:"themes,omitempty"`
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	ForTrigger           TriggerType          `json:"forTrigger,omitempty"`
}

// GetEmailTemplateAvailableTokens DTO.
type GetEmailTemplateAvailableTokens struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// UpdateEmailTemplateRequest DTO.
type UpdateEmailTemplateRequest struct {
	SaveEmailTemplate
	ViewId string `json:"viewId,omitempty"`
}

// DeleteEmailSignature DTO.
type DeleteEmailSignature struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailSignature DTO.
type GetEmailSignature struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailSignatures DTO.
type GetEmailSignatures struct {
	CodeMashListPaginationRequestBase
}

// SaveEmailSignatureRequest DTO.
type SaveEmailSignatureRequest struct {
	CodeMashRequestBase
	ViewId       string            `json:"viewId,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Translations []*TranslationDto `json:"translations,omitempty"`
}

// GetEmailSettings DTO.
type GetEmailSettings struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// ConfirmEmailIntegrationHumanDeliveryRequest DTO.
type ConfirmEmailIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// DeleteEmailIntegration DTO.
type DeleteEmailIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableEmailIntegration DTO.
type DisableEmailIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// CheckEmailIntegrationDomainHealthRequest DTO.
type CheckEmailIntegrationDomainHealthRequest struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// EnableEmailIntegration DTO.
type EnableEmailIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailIntegration DTO.
type GetEmailIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailIntegrations DTO.
type GetEmailIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveEmailIntegration DTO.
type SaveEmailIntegration struct {
	CodeMashRequestBase
	Integration *EmailIntegrationRequest `json:"integration,omitempty"`
}

// SetEmailsIntegrationAsDefault DTO.
type SetEmailsIntegrationAsDefault struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// TestEmailIntegration DTO.
type TestEmailIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
	To            string `json:"to,omitempty"`
}

// ArchiveEmailTemplateRequest DTO.
type ArchiveEmailTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// CloneEmailTemplateRequest DTO.
type CloneEmailTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// UnArchiveEmailTemplateRequest DTO.
type UnArchiveEmailTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DeleteEmailFooter DTO.
type DeleteEmailFooter struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailFooter DTO.
type GetEmailFooter struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmailFooters DTO.
type GetEmailFooters struct {
	CodeMashListPaginationRequestBase
}

// SaveEmailFooterRequest DTO.
type SaveEmailFooterRequest struct {
	CodeMashRequestBase
	ViewId       string            `json:"viewId,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Translations []*TranslationDto `json:"translations,omitempty"`
}

// OneClickUnsubscribeRequest DTO.
type OneClickUnsubscribeRequest struct {
	RequestBase
	Token string `json:"token,omitempty"`
}

// CreateEmailCampaignRequest DTO.
type CreateEmailCampaignRequest struct {
	CodeMashRequestBase
	Campaign              *EmailCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string                `json:"databaseIntegrationId,omitempty"`
}

// DeleteEmailCampaignRequest DTO.
type DeleteEmailCampaignRequest struct {
	CodeMashRequestBase
}

// GetEmailCampaign DTO.
type GetEmailCampaign struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetEmailCampaigns DTO.
type GetEmailCampaigns struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	CampaignId            string  `json:"campaignId,omitempty"`
	EmailAddress          string  `json:"emailAddress,omitempty"`
	TemplateId            string  `json:"templateId,omitempty"`
	From                  float64 `json:"from,omitempty"`
	To                    float64 `json:"to,omitempty"`
}

// GetEmailCampaignBatches DTO.
type GetEmailCampaignBatches struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	EmailAddress          string `json:"emailAddress,omitempty"`
}

// GetEmailCampaignBatchNotification DTO.
type GetEmailCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	NotificationId        string `json:"notificationId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetEmailCampaignBatchNotifications DTO.
type GetEmailCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetEmailCampaignStatistics DTO.
type GetEmailCampaignStatistics struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// PreviewEmailNotification DTO.
type PreviewEmailNotification struct {
	RequestBase
	Hash string `json:"hash,omitempty"`
}

// StopEmailCampaignRequest DTO.
type StopEmailCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetEmailCampaignMessage DTO.
type GetEmailCampaignMessage struct {
	CodeMashRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	NotificationId        string `json:"notificationId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetEmailCampaignMessagesRequest DTO.
type GetEmailCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// EmailServiceEstablished DTO.
type EmailServiceEstablished struct {
}

// ProjectDatabaseConnected DTO.
type ProjectDatabaseConnected struct {
	Env *Env `json:"env,omitempty"`
}

// EmailServiceEnabled DTO.
type EmailServiceEnabled struct {
}

// EmailServiceDisabled DTO.
type EmailServiceDisabled struct {
}

// EmailFooterSaved DTO.
type EmailFooterSaved struct {
	Id           *EmailFooterId                      `json:"id,omitempty"`
	Name         *DisplayName                        `json:"name,omitempty"`
	Translations []MessageTranslation[*TemplateCode] `json:"translations,omitempty"`
	Env          *Env                                `json:"env,omitempty"`
}

// EmailFooterMirrored DTO.
type EmailFooterMirrored struct {
	Footer *EmailFooter `json:"footer,omitempty"`
}

// EmailFooterDeleted DTO.
type EmailFooterDeleted struct {
	Id  *EmailFooterId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailIntegrationSaved DTO.
type EmailIntegrationSaved struct {
	Integration *EmailIntegration `json:"integration,omitempty"`
}

// EmailIntegrationTested DTO.
type EmailIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// EmailIntegrationHumanDeliveryConfirmed DTO.
type EmailIntegrationHumanDeliveryConfirmed struct {
	Id             *IntegrationId `json:"id,omitempty"`
	ConfirmedAtUtc string         `json:"confirmedAtUtc,omitempty"`
}

// EmailIntegrationRenamed DTO.
type EmailIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// EmailIntegrationSetAsDefault DTO.
type EmailIntegrationSetAsDefault struct {
	Env *Env           `json:"env,omitempty"`
	Id  *IntegrationId `json:"id,omitempty"`
}

// EmailIntegrationDeleted DTO.
type EmailIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailIntegrationEnabled DTO.
type EmailIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailIntegrationDisabled DTO.
type EmailIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailSignatureSaved DTO.
type EmailSignatureSaved struct {
	Id           *EmailSignatureId                   `json:"id,omitempty"`
	Name         *DisplayName                        `json:"name,omitempty"`
	Translations []MessageTranslation[*TemplateCode] `json:"translations,omitempty"`
	Env          *Env                                `json:"env,omitempty"`
}

// EmailSignatureMirrored DTO.
type EmailSignatureMirrored struct {
	Signature *EmailSignature `json:"signature,omitempty"`
}

// EmailSignatureDeleted DTO.
type EmailSignatureDeleted struct {
	Id  *EmailSignatureId `json:"id,omitempty"`
	Env *Env              `json:"env,omitempty"`
}

// EmailTemplateCreated DTO.
type EmailTemplateCreated struct {
	TemplateId                  *TemplateId                                `json:"templateId,omitempty"`
	DisplayName                 *DisplayName                               `json:"displayName,omitempty"`
	Translations                []MessageTranslation[*EmailMessageContent] `json:"translations,omitempty"`
	Channel                     CommunicationChannel                       `json:"channel,omitempty"`
	Description                 string                                     `json:"description,omitempty"`
	Tags                        []*Tag                                     `json:"tags,omitempty"`
	LanguageAgnosticAttachments []*FileResourceRef                         `json:"languageAgnosticAttachments,omitempty"`
	Env                         *Env                                       `json:"env,omitempty"`
}

// EmailTemplateUpdated DTO.
type EmailTemplateUpdated struct {
	TemplateId                  *TemplateId                                `json:"templateId,omitempty"`
	DisplayName                 *DisplayName                               `json:"displayName,omitempty"`
	Translations                []MessageTranslation[*EmailMessageContent] `json:"translations,omitempty"`
	Channel                     CommunicationChannel                       `json:"channel,omitempty"`
	Description                 string                                     `json:"description,omitempty"`
	Tags                        []*Tag                                     `json:"tags,omitempty"`
	LanguageAgnosticAttachments []*FileResourceRef                         `json:"languageAgnosticAttachments,omitempty"`
	AttachmentsToBeDeleted      []*FileResourceRef                         `json:"attachmentsToBeDeleted,omitempty"`
	Env                         *Env                                       `json:"env,omitempty"`
}

// EmailTemplateMirrored DTO.
type EmailTemplateMirrored struct {
	Template *EmailTemplate `json:"template,omitempty"`
}

// EmailTemplateBackfilled DTO.
type EmailTemplateBackfilled struct {
	Template *EmailTemplate `json:"template,omitempty"`
}

// EmailTemplateDeleted DTO.
type EmailTemplateDeleted struct {
	TemplateId        *TemplateId        `json:"templateId,omitempty"`
	FilesToBeDeleted  []*FileResourceRef `json:"filesToBeDeleted,omitempty"`
	FileIntegrationId *IntegrationId     `json:"fileIntegrationId,omitempty"`
	Env               *Env               `json:"env,omitempty"`
}

// EmailTemplateArchived DTO.
type EmailTemplateArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// EmailTemplateUnArchived DTO.
type EmailTemplateUnArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// EmailValidationIntegrationSaved DTO.
type EmailValidationIntegrationSaved struct {
	Integration *EmailValidationIntegration `json:"integration,omitempty"`
}

// EmailValidationIntegrationDeleted DTO.
type EmailValidationIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailValidationIntegrationSecretsConfigured DTO.
type EmailValidationIntegrationSecretsConfigured struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailValidationIntegrationSecretsConfigurationFailed DTO.
type EmailValidationIntegrationSecretsConfigurationFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// EmailValidationIntegrationTested DTO.
type EmailValidationIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// EmailBatchRegistered DTO.
type EmailBatchRegistered struct {
	ProjectId       *ProjectId       `json:"projectId,omitempty"`
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	StartingAfter   string           `json:"startingAfter,omitempty"`
}

// EmailNotificationRead DTO.
type EmailNotificationRead struct {
	ProjectId       *ProjectId       `json:"projectId,omitempty"`
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	NotificationId  *NotificationId  `json:"notificationId,omitempty"`
}

// EmailNotificationClicked DTO.
type EmailNotificationClicked struct {
	ProjectId       *ProjectId       `json:"projectId,omitempty"`
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	NotificationId  *NotificationId  `json:"notificationId,omitempty"`
	SourceId        string           `json:"sourceId,omitempty"`
}

// EmailCampaignStarted DTO.
type EmailCampaignStarted struct {
	ProjectId  *ProjectId  `json:"projectId,omitempty"`
	CampaignId *CampaignId `json:"campaignId,omitempty"`
}

// EmailCampaignStopped DTO.
type EmailCampaignStopped struct {
	ProjectId  *ProjectId         `json:"projectId,omitempty"`
	CampaignId *CampaignId        `json:"campaignId,omitempty"`
	Reason     CampaignStopReason `json:"reason,omitempty"`
}

// EmailCampaignCompleted DTO.
type EmailCampaignCompleted struct {
	ProjectId  *ProjectId  `json:"projectId,omitempty"`
	CampaignId *CampaignId `json:"campaignId,omitempty"`
	Errors     []*ErrorDto `json:"errors,omitempty"`
}

// EmailCampaignFailed DTO.
type EmailCampaignFailed struct {
	ProjectId  *ProjectId  `json:"projectId,omitempty"`
	CampaignId *CampaignId `json:"campaignId,omitempty"`
	Errors     []*ErrorDto `json:"errors,omitempty"`
}

// EmailCampaignTriggered DTO.
type EmailCampaignTriggered struct {
	ProjectId     *ProjectId        `json:"projectId,omitempty"`
	AccountId     *AccountId        `json:"accountId,omitempty"`
	TriggerId     *TriggerId        `json:"triggerId,omitempty"`
	TriggerType   TriggerType       `json:"triggerType,omitempty"`
	SourceEvent   string            `json:"sourceEvent,omitempty"`
	SchemaId      string            `json:"schemaId,omitempty"`
	TokenMappings map[string]string `json:"tokenMappings,omitempty"`
}

// EmailDeliveryEventReceived DTO.
type EmailDeliveryEventReceived struct {
	ProjectId         *ProjectId             `json:"projectId,omitempty"`
	IntegrationId     *IntegrationId         `json:"integrationId,omitempty"`
	Recipient         *EmailAddress          `json:"recipient,omitempty"`
	Type              EmailDeliveryEventType `json:"type,omitempty"`
	OccurredAt        string                 `json:"occurredAt,omitempty"`
	ProviderMessageId string                 `json:"providerMessageId,omitempty"`
	Reason            string                 `json:"reason,omitempty"`
}

// DisableSms DTO.
type DisableSms struct {
	CodeMashRequestBase
}

// GetSmsDisableDependencies DTO.
type GetSmsDisableDependencies struct {
	CodeMashRequestBase
}

// EnableSms DTO.
type EnableSms struct {
	CodeMashRequestBase
}

// ArchiveSmsTemplateRequest DTO.
type ArchiveSmsTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// CloneSmsTemplateRequest DTO.
type CloneSmsTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// CreateSmsTemplateRequest DTO.
type CreateSmsTemplateRequest struct {
	SaveSmsTemplate
}

// DeleteSmsTemplateRequest DTO.
type DeleteSmsTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSmsTemplate DTO.
type GetSmsTemplate struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSmsTemplates DTO.
type GetSmsTemplates struct {
	CodeMashListPaginationRequestBase
	ShowArchived bool   `json:"showArchived,omitempty"`
	TemplateId   string `json:"templateId,omitempty"`
}

// GetSmsMessageContentTokens DTO.
type GetSmsMessageContentTokens struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// RenderSms DTO.
type RenderSms struct {
	CodeMashRequestBase
	Code         string             `json:"code,omitempty"`
	Tokens       []*TokenMappingDto `json:"tokens,omitempty"`
	IsForPreview bool               `json:"isForPreview,omitempty"`
}

// UnArchiveSmsTemplateRequest DTO.
type UnArchiveSmsTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// UpdateSmsTemplateRequest DTO.
type UpdateSmsTemplateRequest struct {
	SaveSmsTemplate
	ViewId string `json:"viewId,omitempty"`
}

// GetSmsSettings DTO.
type GetSmsSettings struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// ConfirmSmsIntegrationHumanDeliveryRequest DTO.
type ConfirmSmsIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// DeleteSmsIntegrationRequest DTO.
type DeleteSmsIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableSmsIntegrationRequest DTO.
type DisableSmsIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableSmsIntegrationRequest DTO.
type EnableSmsIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSmsIntegration DTO.
type GetSmsIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSmsIntegrations DTO.
type GetSmsIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveSmsIntegration DTO.
type SaveSmsIntegration struct {
	CodeMashRequestBase
	Integration *SmsIntegrationRequest `json:"integration,omitempty"`
}

// SetSmsIntegrationAsDefaultRequest DTO.
type SetSmsIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// TestSmsIntegration DTO.
type TestSmsIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
	To            string `json:"to,omitempty"`
}

// CreateSmsCampaignRequest DTO.
type CreateSmsCampaignRequest struct {
	CodeMashRequestBase
	TemplateId            string                                     `json:"templateId,omitempty"`
	DatabaseIntegrationId string                                     `json:"databaseIntegrationId,omitempty"`
	Language              string                                     `json:"language,omitempty"`
	InitiatorId           string                                     `json:"initiatorId,omitempty"`
	DeliveryType          SmsCampaignRecipientsSourceTypes           `json:"deliveryType,omitempty"`
	AllUsers              *SmsToAllUsersDeliverySettingsDto          `json:"allUsers,omitempty"`
	SpecifiedUsers        *SmsToUsersDeliverySettingsDto             `json:"specifiedUsers,omitempty"`
	Collection            *SmsToCollectionRecordsDeliverySettingsDto `json:"collection,omitempty"`
	PhoneNumbers          *SmsToPhoneNumbersDeliverySettingsDto      `json:"phoneNumbers,omitempty"`
}

// DeleteSmsCampaign DTO.
type DeleteSmsCampaign struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaign DTO.
type GetSmsCampaign struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaigns DTO.
type GetSmsCampaigns struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	TemplateId            string  `json:"templateId,omitempty"`
	From                  float64 `json:"from,omitempty"`
	To                    float64 `json:"to,omitempty"`
}

// GetSmsCampaignBatches DTO.
type GetSmsCampaignBatches struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaignBatchNotification DTO.
type GetSmsCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	NotificationId        string `json:"notificationId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaignBatchNotifications DTO.
type GetSmsCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaignStatistics DTO.
type GetSmsCampaignStatistics struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// PreviewSmsNotification DTO.
type PreviewSmsNotification struct {
	RequestBase
	Hash string `json:"hash,omitempty"`
}

// StopSmsCampaignRequest DTO.
type StopSmsCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaignMessage DTO.
type GetSmsCampaignMessage struct {
	CodeMashRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	NotificationId        string `json:"notificationId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaignMessagesRequest DTO.
type GetSmsCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// SmsIntegrationSaved DTO.
type SmsIntegrationSaved struct {
	Integration *SmsIntegration `json:"integration,omitempty"`
}

// SmsIntegrationTested DTO.
type SmsIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// SmsIntegrationHumanDeliveryConfirmed DTO.
type SmsIntegrationHumanDeliveryConfirmed struct {
	Id             *IntegrationId `json:"id,omitempty"`
	ConfirmedAtUtc string         `json:"confirmedAtUtc,omitempty"`
}

// SmsIntegrationRenamed DTO.
type SmsIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// SmsIntegrationSetAsDefault DTO.
type SmsIntegrationSetAsDefault struct {
	Env *Env           `json:"env,omitempty"`
	Id  *IntegrationId `json:"id,omitempty"`
}

// SmsIntegrationDeleted DTO.
type SmsIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// SmsIntegrationEnabled DTO.
type SmsIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// SmsIntegrationDisabled DTO.
type SmsIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// SmsServiceEstablished DTO.
type SmsServiceEstablished struct {
	DefaultTemplates []*SmsTemplate `json:"defaultTemplates,omitempty"`
}

// SmsServiceEnabled DTO.
type SmsServiceEnabled struct {
}

// SmsServiceDisabled DTO.
type SmsServiceDisabled struct {
}

// SmsTemplateCreated DTO.
type SmsTemplateCreated struct {
	TemplateId   *TemplateId                              `json:"templateId,omitempty"`
	DisplayName  *DisplayName                             `json:"displayName,omitempty"`
	Translations []MessageTranslation[*SmsMessageContent] `json:"translations,omitempty"`
	Channel      CommunicationChannel                     `json:"channel,omitempty"`
	Description  string                                   `json:"description,omitempty"`
	Tags         []*Tag                                   `json:"tags,omitempty"`
	Env          *Env                                     `json:"env,omitempty"`
}

// SmsTemplateUpdated DTO.
type SmsTemplateUpdated struct {
	TemplateId   *TemplateId                              `json:"templateId,omitempty"`
	DisplayName  *DisplayName                             `json:"displayName,omitempty"`
	Translations []MessageTranslation[*SmsMessageContent] `json:"translations,omitempty"`
	Channel      CommunicationChannel                     `json:"channel,omitempty"`
	Description  string                                   `json:"description,omitempty"`
	Tags         []*Tag                                   `json:"tags,omitempty"`
	Env          *Env                                     `json:"env,omitempty"`
}

// SmsTemplateMirrored DTO.
type SmsTemplateMirrored struct {
	Template *SmsTemplate `json:"template,omitempty"`
}

// SmsTemplateDeleted DTO.
type SmsTemplateDeleted struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// SmsTemplateArchived DTO.
type SmsTemplateArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// SmsTemplateUnArchived DTO.
type SmsTemplateUnArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// SmsBatchRegistered DTO.
type SmsBatchRegistered struct {
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	StartingAfter   string           `json:"startingAfter,omitempty"`
}

// SmsNotificationRead DTO.
type SmsNotificationRead struct {
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	NotificationId  *NotificationId  `json:"notificationId,omitempty"`
}

// SmsNotificationClicked DTO.
type SmsNotificationClicked struct {
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	NotificationId  *NotificationId  `json:"notificationId,omitempty"`
	SourceId        string           `json:"sourceId,omitempty"`
}

// SmsCampaignStarted DTO.
type SmsCampaignStarted struct {
	CampaignId *CampaignId `json:"campaignId,omitempty"`
}

// SmsCampaignStopped DTO.
type SmsCampaignStopped struct {
	CampaignId *CampaignId        `json:"campaignId,omitempty"`
	Reason     CampaignStopReason `json:"reason,omitempty"`
}

// SmsCampaignCompleted DTO.
type SmsCampaignCompleted struct {
	CampaignId *CampaignId `json:"campaignId,omitempty"`
	Errors     []*ErrorDto `json:"errors,omitempty"`
}

// SmsCampaignFailed DTO.
type SmsCampaignFailed struct {
	CampaignId *CampaignId `json:"campaignId,omitempty"`
	Errors     []*ErrorDto `json:"errors,omitempty"`
}

// SmsCampaignTriggered DTO.
type SmsCampaignTriggered struct {
	ProjectId     *ProjectId        `json:"projectId,omitempty"`
	AccountId     *AccountId        `json:"accountId,omitempty"`
	TriggerId     *TriggerId        `json:"triggerId,omitempty"`
	TriggerType   TriggerType       `json:"triggerType,omitempty"`
	SourceEvent   string            `json:"sourceEvent,omitempty"`
	SchemaId      string            `json:"schemaId,omitempty"`
	TokenMappings map[string]string `json:"tokenMappings,omitempty"`
}

// ReplaceMarketplaceIntegrationSecretsRequest DTO.
type ReplaceMarketplaceIntegrationSecretsRequest struct {
	CodeMashRequestBase
	IntegrationViewId string            `json:"integrationViewId,omitempty"`
	Secrets           map[string]string `json:"secrets,omitempty"`
}

// RevealMarketplaceIntegrationSecretsRequest DTO.
type RevealMarketplaceIntegrationSecretsRequest struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// SetMarketplaceIntegrationTokenMappingsRequest DTO.
type SetMarketplaceIntegrationTokenMappingsRequest struct {
	CodeMashRequestBase
	IntegrationViewId string                        `json:"integrationViewId,omitempty"`
	TokenMappings     []*MarketplaceTokenMappingDto `json:"tokenMappings,omitempty"`
}

// GetMarketplaceFunctionCatalog DTO.
type GetMarketplaceFunctionCatalog struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// CodeIntegrationSaved DTO.
type CodeIntegrationSaved struct {
	Integration *CodeIntegration `json:"integration,omitempty"`
}

// CodeIntegrationTested DTO.
type CodeIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// CodeIntegrationHumanDeliveryConfirmed DTO.
type CodeIntegrationHumanDeliveryConfirmed struct {
	Id             *IntegrationId `json:"id,omitempty"`
	ConfirmedAtUtc string         `json:"confirmedAtUtc,omitempty"`
}

// CodeIntegrationRenamed DTO.
type CodeIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// CodeIntegrationSetAsDefault DTO.
type CodeIntegrationSetAsDefault struct {
	Id *IntegrationId `json:"id,omitempty"`
}

// CodeIntegrationDeleted DTO.
type CodeIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// CodeIntegrationEnabled DTO.
type CodeIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// CodeIntegrationDisabled DTO.
type CodeIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MarketplaceIntegrationSaved DTO.
type MarketplaceIntegrationSaved struct {
	Integration *MarketplaceIntegration `json:"integration,omitempty"`
}

// MarketplaceIntegrationDeleted DTO.
type MarketplaceIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MarketplaceIntegrationEnabled DTO.
type MarketplaceIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MarketplaceIntegrationDisabled DTO.
type MarketplaceIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MarketplaceIntegrationTested DTO.
type MarketplaceIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// MarketplaceIntegrationSecretsConfigured DTO.
type MarketplaceIntegrationSecretsConfigured struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MarketplaceIntegrationSecretsConfigurationFailed DTO.
type MarketplaceIntegrationSecretsConfigurationFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// MarketplaceFunctionSaved DTO.
type MarketplaceFunctionSaved struct {
	Function *MarketplaceFunction `json:"function,omitempty"`
}

// MarketplaceFunctionDeleted DTO.
type MarketplaceFunctionDeleted struct {
	IntegrationId *IntegrationId         `json:"integrationId,omitempty"`
	FunctionId    *MarketplaceFunctionId `json:"functionId,omitempty"`
}

// MarketplaceFunctionEnabled DTO.
type MarketplaceFunctionEnabled struct {
	IntegrationId *IntegrationId         `json:"integrationId,omitempty"`
	FunctionId    *MarketplaceFunctionId `json:"functionId,omitempty"`
}

// MarketplaceFunctionDisabled DTO.
type MarketplaceFunctionDisabled struct {
	IntegrationId *IntegrationId         `json:"integrationId,omitempty"`
	FunctionId    *MarketplaceFunctionId `json:"functionId,omitempty"`
}

// ServerlessEnabled DTO.
type ServerlessEnabled struct {
}

// ServerlessDisabled DTO.
type ServerlessDisabled struct {
}

// MarketplaceFunctionTriggered DTO.
type MarketplaceFunctionTriggered struct {
	ProjectId        *ProjectId  `json:"projectId,omitempty"`
	AccountId        *AccountId  `json:"accountId,omitempty"`
	TriggerId        *TriggerId  `json:"triggerId,omitempty"`
	TriggerType      TriggerType `json:"triggerType,omitempty"`
	SourceEvent      string      `json:"sourceEvent,omitempty"`
	SchemaId         string      `json:"schemaId,omitempty"`
	TargetUserAuthId string      `json:"targetUserAuthId,omitempty"`
	OldDocumentJson  string      `json:"oldDocumentJson,omitempty"`
	NewDocumentJson  string      `json:"newDocumentJson,omitempty"`
	Collection       string      `json:"collection,omitempty"`
	CorrelationId    string      `json:"correlationId,omitempty"`
}

// DisablePush DTO.
type DisablePush struct {
	CodeMashRequestBase
}

// GetPushDisableDependencies DTO.
type GetPushDisableDependencies struct {
	CodeMashRequestBase
}

// EnablePush DTO.
type EnablePush struct {
	CodeMashRequestBase
}

// ArchivePushTemplateRequest DTO.
type ArchivePushTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// ClonePushTemplateRequest DTO.
type ClonePushTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// CreatePushTemplateRequest DTO.
type CreatePushTemplateRequest struct {
	SavePushTemplate
}

// DeletePushTemplateRequest DTO.
type DeletePushTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPushTemplate DTO.
type GetPushTemplate struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPushTemplates DTO.
type GetPushTemplates struct {
	CodeMashListPaginationRequestBase
	ShowArchived bool   `json:"showArchived,omitempty"`
	TemplateId   string `json:"templateId,omitempty"`
}

// GetPushMessageContentTokens DTO.
type GetPushMessageContentTokens struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// RenderPush DTO.
type RenderPush struct {
	CodeMashRequestBase
	Code         string             `json:"code,omitempty"`
	Tokens       []*TokenMappingDto `json:"tokens,omitempty"`
	IsForPreview bool               `json:"isForPreview,omitempty"`
}

// UnArchivePushTemplateRequest DTO.
type UnArchivePushTemplateRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// UpdatePushTemplateRequest DTO.
type UpdatePushTemplateRequest struct {
	SavePushTemplate
	ViewId string `json:"viewId,omitempty"`
}

// GetPushSettings DTO.
type GetPushSettings struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// ConfirmPushIntegrationHumanDeliveryRequest DTO.
type ConfirmPushIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// DeletePushIntegrationRequest DTO.
type DeletePushIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisablePushIntegrationRequest DTO.
type DisablePushIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnablePushIntegrationRequest DTO.
type EnablePushIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPushIntegration DTO.
type GetPushIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPushIntegrations DTO.
type GetPushIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SavePushIntegration DTO.
type SavePushIntegration struct {
	CodeMashRequestBase
	Integration *PushIntegrationRequest `json:"integration,omitempty"`
}

// SetPushIntegrationAsDefaultRequest DTO.
type SetPushIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// TestPushIntegration DTO.
type TestPushIntegration struct {
	CodeMashRequestBase
	IntegrationId  string `json:"integrationId,omitempty"`
	TestToken      string `json:"testToken,omitempty"`
	DeliveryFamily string `json:"deliveryFamily,omitempty"`
}

// RegisterCodeMashAppPushIntegration DTO.
type RegisterCodeMashAppPushIntegration struct {
	CodeMashRequestBase
	AccountId string  `json:"accountId,omitempty"`
	UserId    string  `json:"userId,omitempty"`
	RequestId string  `json:"requestId,omitempty"`
	Pin       float64 `json:"pin,omitempty"`
	ValidTill string  `json:"validTill,omitempty"`
	PublicKey string  `json:"publicKey,omitempty"`
}

// RegisterDevice DTO.
type RegisterDevice struct {
	RequestBase
	PushDeviceDto         *PushDeviceDto `json:"pushDeviceDto,omitempty"`
	UserId                string         `json:"userId,omitempty"`
	ProjectId             string         `json:"projectId,omitempty"`
	AccountId             string         `json:"accountId,omitempty"`
	DatabaseIntegrationId string         `json:"databaseIntegrationId,omitempty"`
}

// CreatePushCampaignRequest DTO.
type CreatePushCampaignRequest struct {
	CodeMashRequestBase
	Campaign              *PushCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string               `json:"databaseIntegrationId,omitempty"`
}

// DeletePushCampaignRequest DTO.
type DeletePushCampaignRequest struct {
	CodeMashRequestBase
}

// GetPushCampaign DTO.
type GetPushCampaign struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushCampaigns DTO.
type GetPushCampaigns struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId string  `json:"databaseIntegrationId,omitempty"`
	TemplateId            string  `json:"templateId,omitempty"`
	From                  float64 `json:"from,omitempty"`
	To                    float64 `json:"to,omitempty"`
}

// GetPushCampaignBatches DTO.
type GetPushCampaignBatches struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
}

// GetPushCampaignBatchNotification DTO.
type GetPushCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	NotificationId        string `json:"notificationId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushCampaignBatchNotifications DTO.
type GetPushCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	Id                    string `json:"id,omitempty"`
	BatchId               string `json:"batchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushCampaignStatistics DTO.
type GetPushCampaignStatistics struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// PreviewPushNotification DTO.
type PreviewPushNotification struct {
	RequestBase
	Hash string `json:"hash,omitempty"`
}

// StopPushCampaignRequest DTO.
type StopPushCampaignRequest struct {
	CodeMashRequestBase
}

// GetPushCampaignMessage DTO.
type GetPushCampaignMessage struct {
	CodeMashRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	NotificationId        string `json:"notificationId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushCampaignMessagesRequest DTO.
type GetPushCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// PushIntegrationSaved DTO.
type PushIntegrationSaved struct {
	Integration *PushIntegration `json:"integration,omitempty"`
}

// PushIntegrationTested DTO.
type PushIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// PushIntegrationHumanDeliveryConfirmed DTO.
type PushIntegrationHumanDeliveryConfirmed struct {
	Id             *IntegrationId `json:"id,omitempty"`
	ConfirmedAtUtc string         `json:"confirmedAtUtc,omitempty"`
}

// PushIntegrationRenamed DTO.
type PushIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// PushIntegrationSetAsDefault DTO.
type PushIntegrationSetAsDefault struct {
	Env *Env           `json:"env,omitempty"`
	Id  *IntegrationId `json:"id,omitempty"`
}

// PushIntegrationDeleted DTO.
type PushIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PushIntegrationEnabled DTO.
type PushIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PushIntegrationDisabled DTO.
type PushIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PushServiceEstablished DTO.
type PushServiceEstablished struct {
	DefaultTemplates []*PushTemplate `json:"defaultTemplates,omitempty"`
}

// PushServiceEnabled DTO.
type PushServiceEnabled struct {
}

// PushServiceDisabled DTO.
type PushServiceDisabled struct {
}

// PushModuleTagSaved DTO.
type PushModuleTagSaved struct {
	Tag                  *TagDefinition       `json:"tag,omitempty"`
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
}

// PushModuleTagDeleted DTO.
type PushModuleTagDeleted struct {
	Tag                  *Tag                 `json:"tag,omitempty"`
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
}

// PushTemplateCreated DTO.
type PushTemplateCreated struct {
	TemplateId   *TemplateId                               `json:"templateId,omitempty"`
	DisplayName  *DisplayName                              `json:"displayName,omitempty"`
	Translations []MessageTranslation[*PushMessageContent] `json:"translations,omitempty"`
	Channel      CommunicationChannel                      `json:"channel,omitempty"`
	Description  string                                    `json:"description,omitempty"`
	Tags         []*Tag                                    `json:"tags,omitempty"`
	Env          *Env                                      `json:"env,omitempty"`
}

// PushTemplateUpdated DTO.
type PushTemplateUpdated struct {
	TemplateId   *TemplateId                               `json:"templateId,omitempty"`
	DisplayName  *DisplayName                              `json:"displayName,omitempty"`
	Translations []MessageTranslation[*PushMessageContent] `json:"translations,omitempty"`
	Channel      CommunicationChannel                      `json:"channel,omitempty"`
	Description  string                                    `json:"description,omitempty"`
	Tags         []*Tag                                    `json:"tags,omitempty"`
	Env          *Env                                      `json:"env,omitempty"`
}

// PushTemplateMirrored DTO.
type PushTemplateMirrored struct {
	Template *PushTemplate `json:"template,omitempty"`
}

// PushTemplateDeleted DTO.
type PushTemplateDeleted struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// PushTemplateArchived DTO.
type PushTemplateArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// PushTemplateUnArchived DTO.
type PushTemplateUnArchived struct {
	TemplateId *TemplateId `json:"templateId,omitempty"`
	Env        *Env        `json:"env,omitempty"`
}

// PushBatchRegistered DTO.
type PushBatchRegistered struct {
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	StartingAfter   string           `json:"startingAfter,omitempty"`
}

// PushNotificationRead DTO.
type PushNotificationRead struct {
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	NotificationId  *NotificationId  `json:"notificationId,omitempty"`
}

// PushNotificationClicked DTO.
type PushNotificationClicked struct {
	CampaignId      *CampaignId      `json:"campaignId,omitempty"`
	CampaignBatchId *CampaignBatchId `json:"campaignBatchId,omitempty"`
	NotificationId  *NotificationId  `json:"notificationId,omitempty"`
	SourceId        string           `json:"sourceId,omitempty"`
}

// PushCampaignStarted DTO.
type PushCampaignStarted struct {
	CampaignId *CampaignId `json:"campaignId,omitempty"`
}

// PushCampaignStopped DTO.
type PushCampaignStopped struct {
	CampaignId *CampaignId        `json:"campaignId,omitempty"`
	Reason     CampaignStopReason `json:"reason,omitempty"`
}

// PushCampaignCompleted DTO.
type PushCampaignCompleted struct {
	CampaignId *CampaignId `json:"campaignId,omitempty"`
	Errors     []*ErrorDto `json:"errors,omitempty"`
}

// PushCampaignFailed DTO.
type PushCampaignFailed struct {
	CampaignId *CampaignId `json:"campaignId,omitempty"`
	Errors     []*ErrorDto `json:"errors,omitempty"`
}

// PushCampaignTriggered DTO.
type PushCampaignTriggered struct {
	ProjectId     *ProjectId        `json:"projectId,omitempty"`
	AccountId     *AccountId        `json:"accountId,omitempty"`
	TriggerId     *TriggerId        `json:"triggerId,omitempty"`
	TriggerType   TriggerType       `json:"triggerType,omitempty"`
	SourceEvent   string            `json:"sourceEvent,omitempty"`
	SchemaId      string            `json:"schemaId,omitempty"`
	TokenMappings map[string]string `json:"tokenMappings,omitempty"`
}

// DisablePayments DTO.
type DisablePayments struct {
	CodeMashRequestBase
}

// EnablePayments DTO.
type EnablePayments struct {
	CodeMashRequestBase
}

// GetPaymentsWebhookLog DTO.
type GetPaymentsWebhookLog struct {
	CodeMashRequestBase
	IntegrationId string  `json:"integrationId,omitempty"`
	Limit         float64 `json:"limit,omitempty"`
}

// DeletePaymentsTrigger DTO.
type DeletePaymentsTrigger struct {
	DeleteTrigger
}

// DisablePaymentsTrigger DTO.
type DisablePaymentsTrigger struct {
	DisableTrigger
}

// EnablePaymentsTrigger DTO.
type EnablePaymentsTrigger struct {
	EnableTrigger
}

// GetPaymentsTrigger DTO.
type GetPaymentsTrigger struct {
	GetTrigger
}

// GetPaymentsTriggers DTO.
type GetPaymentsTriggers struct {
	GetTriggers
}

// SavePaymentsTrigger DTO.
type SavePaymentsTrigger struct {
	SaveTrigger
}

// ConfirmPaymentsIntegrationHumanDeliveryRequest DTO.
type ConfirmPaymentsIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// DeletePaymentsIntegrationRequest DTO.
type DeletePaymentsIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisablePaymentsIntegrationRequest DTO.
type DisablePaymentsIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnablePaymentsIntegrationRequest DTO.
type EnablePaymentsIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPaymentsIntegration DTO.
type GetPaymentsIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPaymentsIntegrations DTO.
type GetPaymentsIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SavePaymentsIntegration DTO.
type SavePaymentsIntegration struct {
	CodeMashRequestBase
	Integration *PaymentIntegrationRequest `json:"integration,omitempty"`
}

// TestPaymentsIntegration DTO.
type TestPaymentsIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// PaymentsIntegrationSaved DTO.
type PaymentsIntegrationSaved struct {
	Integration *PaymentIntegration `json:"integration,omitempty"`
}

// PaymentsIntegrationTested DTO.
type PaymentsIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// PaymentsIntegrationHumanDeliveryConfirmed DTO.
type PaymentsIntegrationHumanDeliveryConfirmed struct {
	Id             *IntegrationId `json:"id,omitempty"`
	ConfirmedAtUtc string         `json:"confirmedAtUtc,omitempty"`
}

// PaymentsIntegrationRenamed DTO.
type PaymentsIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// PaymentsIntegrationDeleted DTO.
type PaymentsIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PaymentsIntegrationEnabled DTO.
type PaymentsIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PaymentsIntegrationDisabled DTO.
type PaymentsIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// PaymentsEstablished DTO.
type PaymentsEstablished struct {
}

// PaymentsEnabled DTO.
type PaymentsEnabled struct {
}

// PaymentsDisabled DTO.
type PaymentsDisabled struct {
}

// PaymentsTriggerSaved DTO.
type PaymentsTriggerSaved struct {
	Trigger *PaymentTrigger `json:"trigger,omitempty"`
}

// PaymentTriggerMirrored DTO.
type PaymentTriggerMirrored struct {
	Trigger *Trigger `json:"trigger,omitempty"`
}

// PaymentsTriggerEnabled DTO.
type PaymentsTriggerEnabled struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// PaymentsTriggerDisabled DTO.
type PaymentsTriggerDisabled struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// PaymentsTriggerDeleted DTO.
type PaymentsTriggerDeleted struct {
	TriggerByIdEventBase
	Env *Env `json:"env,omitempty"`
}

// DisableLogging DTO.
type DisableLogging struct {
	CodeMashRequestBase
}

// EnableLogging DTO.
type EnableLogging struct {
	CodeMashRequestBase
	CreateNorbixLogging bool `json:"createNorbixLogging,omitempty"`
}

// DeleteLoggingIntegrationRequest DTO.
type DeleteLoggingIntegrationRequest struct {
	CodeMashRequestBase
	Id       string `json:"id,omitempty"`
	WipeLogs bool   `json:"wipeLogs,omitempty"`
}

// DisableLoggingIntegrationRequest DTO.
type DisableLoggingIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableLoggingIntegrationRequest DTO.
type EnableLoggingIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetLoggingIntegration DTO.
type GetLoggingIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetLoggingIntegrations DTO.
type GetLoggingIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveLoggingIntegration DTO.
type SaveLoggingIntegration struct {
	CodeMashRequestBase
	Integration *LoggingIntegrationRequest `json:"integration,omitempty"`
}

// TestLoggingIntegration DTO.
type TestLoggingIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// CleanLogs DTO.
type CleanLogs struct {
	CodeMashRequestBase
}

// GetLogsByCorrelationId DTO.
type GetLogsByCorrelationId struct {
	CodeMashRequestBase
	TargetCorrelationId string `json:"targetCorrelationId,omitempty"`
}

// GetLogs DTO.
type GetLogs struct {
	CodeMashListPaginationRequestBase
	Level            string `json:"level,omitempty"`
	Module           string `json:"module,omitempty"`
	LogCorrelationId string `json:"logCorrelationId,omitempty"`
	EventCode        string `json:"eventCode,omitempty"`
	Search           string `json:"search,omitempty"`
	FromUtc          string `json:"fromUtc,omitempty"`
	ToUtc            string `json:"toUtc,omitempty"`
}

// GetLogSettings DTO.
type GetLogSettings struct {
	CodeMashRequestBase
}

// SaveLogSettings DTO.
type SaveLogSettings struct {
	CodeMashRequestBase
	SkipCloudDashboardLogs bool `json:"skipCloudDashboardLogs,omitempty"`
	SkipHttpBodyMeta       bool `json:"skipHttpBodyMeta,omitempty"`
	AiChatLoggingEnabled   bool `json:"aiChatLoggingEnabled,omitempty"`
}

// LoggingIntegrationSaved DTO.
type LoggingIntegrationSaved struct {
	Integration *LoggingIntegration `json:"integration,omitempty"`
}

// LoggingIntegrationTested DTO.
type LoggingIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// LoggingIntegrationRenamed DTO.
type LoggingIntegrationRenamed struct {
	Id   *IntegrationId `json:"id,omitempty"`
	Name *DisplayName   `json:"name,omitempty"`
	Env  *Env           `json:"env,omitempty"`
}

// LoggingIntegrationDeleted DTO.
type LoggingIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationEnabled DTO.
type LoggingIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationDisabled DTO.
type LoggingIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationSecretsConfigured DTO.
type LoggingIntegrationSecretsConfigured struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationSecretsConfigurationFailed DTO.
type LoggingIntegrationSecretsConfigurationFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationSecretsCleared DTO.
type LoggingIntegrationSecretsCleared struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationSecretsClearingFailed DTO.
type LoggingIntegrationSecretsClearingFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LoggingIntegrationSetAsDefault DTO.
type LoggingIntegrationSetAsDefault struct {
	Id *IntegrationId `json:"id,omitempty"`
}

// NorbixLoggingLogsWipeRequested DTO.
type NorbixLoggingLogsWipeRequested struct {
	DeletedIntegrationId  *IntegrationId `json:"deletedIntegrationId,omitempty"`
	DatabaseIntegrationId *IntegrationId `json:"databaseIntegrationId,omitempty"`
}

// LoggingEstablished DTO.
type LoggingEstablished struct {
}

// LoggingEnabled DTO.
type LoggingEnabled struct {
}

// LoggingDisabled DTO.
type LoggingDisabled struct {
}

// GetAiToolsRequest DTO.
type GetAiToolsRequest struct {
	RequestBase
	Toolset string `json:"toolset,omitempty"`
}

// InvokeAiToolRequest DTO.
type InvokeAiToolRequest struct {
	RequestBase
	ToolName      string `json:"toolName,omitempty"`
	ArgumentsJson string `json:"argumentsJson,omitempty"`
}

// AskChatRequest DTO.
type AskChatRequest struct {
	RequestBase
	Prompt  string `json:"prompt,omitempty"`
	Profile string `json:"profile,omitempty"`
}

// UploadChatAttachmentRequest DTO.
type UploadChatAttachmentRequest struct {
	RequestBase
	SessionId     string `json:"sessionId,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	Base64Content string `json:"base64Content,omitempty"`
	Profile       string `json:"profile,omitempty"`
	Topic         string `json:"topic,omitempty"`
	ProjectId     string `json:"projectId,omitempty"`
	Env           string `json:"env,omitempty"`
}

// ChatAvailabilityRequest DTO.
type ChatAvailabilityRequest struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Env       string `json:"env,omitempty"`
}

// GetChatMemoryRequest DTO.
type GetChatMemoryRequest struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
}

// ForgetChatMemoryRequest DTO.
type ForgetChatMemoryRequest struct {
	RequestBase
	NoteId string `json:"noteId,omitempty"`
}

// DeleteChatSessionRequest DTO.
type DeleteChatSessionRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// SetChatSessionArchivedRequest DTO.
type SetChatSessionArchivedRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
	Archived  bool   `json:"archived,omitempty"`
}

// SetChatSessionPinnedRequest DTO.
type SetChatSessionPinnedRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
	Pinned    bool   `json:"pinned,omitempty"`
}

// SetChatSessionSharingRequest DTO.
type SetChatSessionSharingRequest struct {
	RequestBase
	SessionId  string `json:"sessionId,omitempty"`
	DoNotShare bool   `json:"doNotShare,omitempty"`
}

// GetChatSessionsRequest DTO.
type GetChatSessionsRequest struct {
	RequestBase
	Take            float64 `json:"take,omitempty"`
	IncludeArchived bool    `json:"includeArchived,omitempty"`
}

// GetChatSessionEntriesRequest DTO.
type GetChatSessionEntriesRequest struct {
	RequestBase
	SessionId string  `json:"sessionId,omitempty"`
	SinceSeq  float64 `json:"sinceSeq,omitempty"`
}

// SetChatEntryFeedbackRequest DTO.
type SetChatEntryFeedbackRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
	EntryId   string `json:"entryId,omitempty"`
	Feedback  string `json:"feedback,omitempty"`
}

// AnswerChatQuestionRequest DTO.
type AnswerChatQuestionRequest struct {
	RequestBase
	SessionId string            `json:"sessionId,omitempty"`
	EntryId   string            `json:"entryId,omitempty"`
	Answers   map[string]string `json:"answers,omitempty"`
}

// DecideChatPlanRequest DTO.
type DecideChatPlanRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
	EntryId   string `json:"entryId,omitempty"`
	Decision  string `json:"decision,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

// StopChatRunStepRequest DTO.
type StopChatRunStepRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
	EntryId   string `json:"entryId,omitempty"`
}

// ChatTurnRequest DTO.
type ChatTurnRequest struct {
	RequestBase
	SessionId        string                `json:"sessionId,omitempty"`
	Message          string                `json:"message,omitempty"`
	Profile          string                `json:"profile,omitempty"`
	Topic            string                `json:"topic,omitempty"`
	LlmIntegrationId string                `json:"llmIntegrationId,omitempty"`
	Model            string                `json:"model,omitempty"`
	ProjectId        string                `json:"projectId,omitempty"`
	Env              string                `json:"env,omitempty"`
	ScreenContext    *ChatScreenContextDto `json:"screenContext,omitempty"`
}

// McpRequest DTO.
type McpRequest struct {
	Version       string `json:"version,omitempty"`
	RequestStream string `json:"requestStream,omitempty"`
}

// GetProjectBriefRequest DTO.
type GetProjectBriefRequest struct {
	CodeMashRequestBase
	SinceSeq float64 `json:"sinceSeq,omitempty"`
}

// GetWorkItemsRequest DTO.
type GetWorkItemsRequest struct {
	CodeMashRequestBase
	Status string `json:"status,omitempty"`
}

// GetWorkItemRequest DTO.
type GetWorkItemRequest struct {
	CodeMashRequestBase
	WorkItemId string `json:"workItemId,omitempty"`
}

// ExportWorkItemRequest DTO.
type ExportWorkItemRequest struct {
	CodeMashRequestBase
	WorkItemId string `json:"workItemId,omitempty"`
}

// MarkNeedsYouDoneRequest DTO.
type MarkNeedsYouDoneRequest struct {
	CodeMashRequestBase
	WorkItemId string  `json:"workItemId,omitempty"`
	Index      float64 `json:"index,omitempty"`
	Done       bool    `json:"done,omitempty"`
}

// DeleteLlmIntegrationRequest DTO.
type DeleteLlmIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableLlmIntegrationRequest DTO.
type DisableLlmIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableLlmIntegrationRequest DTO.
type EnableLlmIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetLlmIntegration DTO.
type GetLlmIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetLlmIntegrations DTO.
type GetLlmIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveLlmIntegration DTO.
type SaveLlmIntegration struct {
	CodeMashRequestBase
	Integration *LlmIntegrationRequest `json:"integration,omitempty"`
}

// TestLlmIntegration DTO.
type TestLlmIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// DeleteMcpIntegrationRequest DTO.
type DeleteMcpIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableMcpIntegrationRequest DTO.
type DisableMcpIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableMcpIntegrationRequest DTO.
type EnableMcpIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetMcpIntegration DTO.
type GetMcpIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetMcpIntegrations DTO.
type GetMcpIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveMcpIntegration DTO.
type SaveMcpIntegration struct {
	CodeMashRequestBase
	Integration *McpIntegrationRequest `json:"integration,omitempty"`
}

// TestMcpIntegration DTO.
type TestMcpIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
}

// LlmIntegrationSaved DTO.
type LlmIntegrationSaved struct {
	LlmIntegration *LlmIntegration `json:"llmIntegration,omitempty"`
}

// LlmIntegrationDeleted DTO.
type LlmIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LlmIntegrationEnabled DTO.
type LlmIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LlmIntegrationDisabled DTO.
type LlmIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LlmIntegrationSecretsConfigured DTO.
type LlmIntegrationSecretsConfigured struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LlmIntegrationSecretsConfigurationFailed DTO.
type LlmIntegrationSecretsConfigurationFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// LlmIntegrationTested DTO.
type LlmIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// McpIntegrationSaved DTO.
type McpIntegrationSaved struct {
	McpIntegration *McpIntegration `json:"mcpIntegration,omitempty"`
}

// McpIntegrationDeleted DTO.
type McpIntegrationDeleted struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// McpIntegrationEnabled DTO.
type McpIntegrationEnabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// McpIntegrationDisabled DTO.
type McpIntegrationDisabled struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// McpIntegrationSecretsConfigured DTO.
type McpIntegrationSecretsConfigured struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// McpIntegrationSecretsConfigurationFailed DTO.
type McpIntegrationSecretsConfigurationFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// McpIntegrationTested DTO.
type McpIntegrationTested struct {
	Id            *IntegrationId `json:"id,omitempty"`
	Succeeded     bool           `json:"succeeded,omitempty"`
	ErrorMessages []string       `json:"errorMessages,omitempty"`
	TestedAtUtc   string         `json:"testedAtUtc,omitempty"`
	Env           *Env           `json:"env,omitempty"`
}

// WebhookIntegrationSaved DTO.
type WebhookIntegrationSaved struct {
	Integration *WebhookIntegration `json:"integration,omitempty"`
}

// WebhookIntegrationExtraHeadersChanged DTO.
type WebhookIntegrationExtraHeadersChanged struct {
	Id           *IntegrationId    `json:"id,omitempty"`
	ExtraHeaders map[string]string `json:"extraHeaders,omitempty"`
}

// WebhookIntegrationSecretsConfigured DTO.
type WebhookIntegrationSecretsConfigured struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// WebhookIntegrationSecretsConfigurationFailed DTO.
type WebhookIntegrationSecretsConfigurationFailed struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// WebhookIntegrationSecretsCleared DTO.
type WebhookIntegrationSecretsCleared struct {
	Id  *IntegrationId `json:"id,omitempty"`
	Env *Env           `json:"env,omitempty"`
}

// WebhookDestinationSaved DTO.
type WebhookDestinationSaved struct {
	IntegrationId *IntegrationId      `json:"integrationId,omitempty"`
	Destination   *WebhookDestination `json:"destination,omitempty"`
}

// WebhookDestinationRemoved DTO.
type WebhookDestinationRemoved struct {
	IntegrationId *IntegrationId        `json:"integrationId,omitempty"`
	DestinationId *WebhookDestinationId `json:"destinationId,omitempty"`
}

// WebhookDestinationEnabled DTO.
type WebhookDestinationEnabled struct {
	IntegrationId *IntegrationId        `json:"integrationId,omitempty"`
	DestinationId *WebhookDestinationId `json:"destinationId,omitempty"`
}

// WebhookDestinationDisabled DTO.
type WebhookDestinationDisabled struct {
	IntegrationId *IntegrationId        `json:"integrationId,omitempty"`
	DestinationId *WebhookDestinationId `json:"destinationId,omitempty"`
}

// GetWebhookIntegration DTO.
type GetWebhookIntegration struct {
	CodeMashRequestBase
}

// RevealWebhookIntegrationSecretRequest DTO.
type RevealWebhookIntegrationSecretRequest struct {
	CodeMashRequestBase
}

// RotateWebhookIntegrationSecretRequest DTO.
type RotateWebhookIntegrationSecretRequest struct {
	CodeMashRequestBase
}

// UpdateWebhookIntegrationExtraHeadersRequest DTO.
type UpdateWebhookIntegrationExtraHeadersRequest struct {
	CodeMashRequestBase
	ExtraHeaders map[string]string `json:"extraHeaders,omitempty"`
}

// ReceiveWebhook DTO.
type ReceiveWebhook struct {
	Source                string `json:"source,omitempty"`
	IntegrationInstanceId string `json:"integrationInstanceId,omitempty"`
	RequestStream         string `json:"requestStream,omitempty"`
}

// DisableWebhookDestinationRequest DTO.
type DisableWebhookDestinationRequest struct {
	CodeMashRequestBase
	DestinationId string `json:"destinationId,omitempty"`
}

// EnableWebhookDestinationRequest DTO.
type EnableWebhookDestinationRequest struct {
	CodeMashRequestBase
	DestinationId string `json:"destinationId,omitempty"`
}

// RemoveWebhookDestinationRequest DTO.
type RemoveWebhookDestinationRequest struct {
	CodeMashRequestBase
	DestinationId string `json:"destinationId,omitempty"`
}

// SaveWebhookDestinationRequest DTO.
type SaveWebhookDestinationRequest struct {
	CodeMashRequestBase
	DestinationId   string            `json:"destinationId,omitempty"`
	DestinationName string            `json:"destinationName,omitempty"`
	EndpointUrl     string            `json:"endpointUrl,omitempty"`
	SelectedEvents  []string          `json:"selectedEvents,omitempty"`
	ExtraHeaders    map[string]string `json:"extraHeaders,omitempty"`
	IsEnabled       bool              `json:"isEnabled,omitempty"`
}

// DisableScheduler DTO.
type DisableScheduler struct {
	CodeMashRequestBase
}

// EnableScheduler DTO.
type EnableScheduler struct {
	CodeMashRequestBase
}

// DeleteSchedulerTask DTO.
type DeleteSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableSchedulerTask DTO.
type DisableSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableSchedulerTask DTO.
type EnableSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSchedulerTask DTO.
type GetSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetSchedulerTasks DTO.
type GetSchedulerTasks struct {
	CodeMashListPaginationRequestBase
	Type    SchedulerTaskType `json:"type,omitempty"`
	Enabled bool              `json:"enabled,omitempty"`
}

// SaveSchedulerTaskRequest DTO.
type SaveSchedulerTaskRequest struct {
	CodeMashRequestBase
	TaskId          string                `json:"taskId,omitempty"`
	Name            string                `json:"name,omitempty"`
	Description     string                `json:"description,omitempty"`
	Cron            string                `json:"cron,omitempty"`
	InitiatorUserId string                `json:"initiatorUserId,omitempty"`
	IsEnabled       bool                  `json:"isEnabled,omitempty"`
	StopOnError     bool                  `json:"stopOnError,omitempty"`
	Task            *SchedulerTaskRequest `json:"task,omitempty"`
}

// SchedulerEnabled DTO.
type SchedulerEnabled struct {
}

// SchedulerDisabled DTO.
type SchedulerDisabled struct {
}

// SchedulerTaskSaved DTO.
type SchedulerTaskSaved struct {
	Task *SchedulerTask `json:"task,omitempty"`
}

// SchedulerTaskEnabled DTO.
type SchedulerTaskEnabled struct {
	TaskId *TaskId `json:"taskId,omitempty"`
}

// SchedulerTaskDisabled DTO.
type SchedulerTaskDisabled struct {
	TaskId *TaskId `json:"taskId,omitempty"`
}

// SchedulerTaskDeleted DTO.
type SchedulerTaskDeleted struct {
	TaskId *TaskId `json:"taskId,omitempty"`
}

// ResolveResources DTO.
type ResolveResources struct {
	CodeMashRequestBase
	Refs []*ResourceRefDto `json:"refs,omitempty"`
}

// CreateContactRequest DTO.
type CreateContactRequest struct {
	CodeMashRequestBase
	PrimaryEmail string `json:"primaryEmail,omitempty"`
	PrimaryPhone string `json:"primaryPhone,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	LastName     string `json:"lastName,omitempty"`
}

// DeleteContact DTO.
type DeleteContact struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
}

// GetContact DTO.
type GetContact struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
}

// GetAllContacts DTO.
type GetAllContacts struct {
	CodeMashRequestBase
	StartingAfter string  `json:"startingAfter,omitempty"`
	PageSize      float64 `json:"pageSize,omitempty"`
}

// MergeContactsRequest DTO.
type MergeContactsRequest struct {
	CodeMashRequestBase
	SurvivorId string   `json:"survivorId,omitempty"`
	MergedIds  []string `json:"mergedIds,omitempty"`
}

// UpdateContactRequest DTO.
type UpdateContactRequest struct {
	CodeMashRequestBase
	ContactId    string  `json:"contactId,omitempty"`
	DisplayName  string  `json:"displayName,omitempty"`
	FirstName    string  `json:"firstName,omitempty"`
	LastName     string  `json:"lastName,omitempty"`
	FullName     string  `json:"fullName,omitempty"`
	Company      string  `json:"company,omitempty"`
	Notes        string  `json:"notes,omitempty"`
	Gender       string  `json:"gender,omitempty"`
	BirthDate    float64 `json:"birthDate,omitempty"`
	TimeZone     string  `json:"timeZone,omitempty"`
	Language     string  `json:"language,omitempty"`
	AddressLine1 string  `json:"addressLine1,omitempty"`
	AddressLine2 string  `json:"addressLine2,omitempty"`
	Country      string  `json:"country,omitempty"`
	City         string  `json:"city,omitempty"`
	State        string  `json:"state,omitempty"`
	PostalCode   string  `json:"postalCode,omitempty"`
}

// AddContactIdentityRequest DTO.
type AddContactIdentityRequest struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
	AuthId    string `json:"authId,omitempty"`
}

// PromoteContactIdentityRequest DTO.
type PromoteContactIdentityRequest struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
	AuthId    string `json:"authId,omitempty"`
}

// RemoveContactIdentityRequest DTO.
type RemoveContactIdentityRequest struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
	AuthId    string `json:"authId,omitempty"`
}

// GetComplianceSettings DTO.
type GetComplianceSettings struct {
	CodeMashRequestBase
}

// RemoveRetentionWindowRequest DTO.
type RemoveRetentionWindowRequest struct {
	CodeMashRequestBase
	DataKind string `json:"dataKind,omitempty"`
}

// SaveRetentionWindowRequest DTO.
type SaveRetentionWindowRequest struct {
	CodeMashRequestBase
	DataKind string  `json:"dataKind,omitempty"`
	Days     float64 `json:"days,omitempty"`
	Action   string  `json:"action,omitempty"`
}

// AssignRegimeRequest DTO.
type AssignRegimeRequest struct {
	CodeMashRequestBase
	Regime string `json:"regime,omitempty"`
}

// ClearRegimeRequest DTO.
type ClearRegimeRequest struct {
	CodeMashRequestBase
	Regime string `json:"regime,omitempty"`
}

// DefineConsentPurposeRequest DTO.
type DefineConsentPurposeRequest struct {
	CodeMashRequestBase
	Key             string   `json:"key,omitempty"`
	Name            string   `json:"name,omitempty"`
	Channel         string   `json:"channel,omitempty"`
	MappedTags      []string `json:"mappedTags,omitempty"`
	RegulatoryBasis []string `json:"regulatoryBasis,omitempty"`
	Description     string   `json:"description,omitempty"`
}

// DeprecateConsentPurposeRequest DTO.
type DeprecateConsentPurposeRequest struct {
	CodeMashRequestBase
	Key string `json:"key,omitempty"`
}

// GetLegalHolds DTO.
type GetLegalHolds struct {
	CodeMashRequestBase
}

// PlaceLegalHoldRequest DTO.
type PlaceLegalHoldRequest struct {
	CodeMashRequestBase
	SubjectKind string `json:"subjectKind,omitempty"`
	SubjectId   string `json:"subjectId,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

// ReleaseLegalHoldRequest DTO.
type ReleaseLegalHoldRequest struct {
	CodeMashRequestBase
	HoldId string `json:"holdId,omitempty"`
}

// ApproveDsarRequestRequest DTO.
type ApproveDsarRequestRequest struct {
	CodeMashRequestBase
	RequestId string `json:"requestId,omitempty"`
}

// RejectDsarRequestRequest DTO.
type RejectDsarRequestRequest struct {
	CodeMashRequestBase
	RequestId string `json:"requestId,omitempty"`
	Reason    string `json:"reason,omitempty"`
}

// GetDsarRequests DTO.
type GetDsarRequests struct {
	CodeMashRequestBase
}

// OpenDsarRequestRequest DTO.
type OpenDsarRequestRequest struct {
	CodeMashRequestBase
	SubjectKind string `json:"subjectKind,omitempty"`
	SubjectId   string `json:"subjectId,omitempty"`
}

// GetComplianceAuditLog DTO.
type GetComplianceAuditLog struct {
	CodeMashRequestBase
	From        string  `json:"from,omitempty"`
	To          string  `json:"to,omitempty"`
	SubjectKind string  `json:"subjectKind,omitempty"`
	SubjectId   string  `json:"subjectId,omitempty"`
	Limit       float64 `json:"limit,omitempty"`
}

// GetAccountCompliance DTO.
type GetAccountCompliance struct {
	RequestBase
}

// SaveDsarPolicyRequest DTO.
type SaveDsarPolicyRequest struct {
	RequestBase
	Mode      string  `json:"mode,omitempty"`
	DelayDays float64 `json:"delayDays,omitempty"`
}

// SaveIncidentRoutingRequest DTO.
type SaveIncidentRoutingRequest struct {
	RequestBase
	AutoForwardAdvisories bool   `json:"autoForwardAdvisories,omitempty"`
	SecurityContact       string `json:"securityContact,omitempty"`
}

// CloseSupportCaseRequest DTO.
type CloseSupportCaseRequest struct {
	RequestBase
	CaseId string `json:"caseId,omitempty"`
}

// ReopenSupportCaseRequest DTO.
type ReopenSupportCaseRequest struct {
	RequestBase
	CaseId string `json:"caseId,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// ResolveSupportCaseRequest DTO.
type ResolveSupportCaseRequest struct {
	RequestBase
	CaseId     string             `json:"caseId,omitempty"`
	Resolution *CaseResolutionDto `json:"resolution,omitempty"`
}

// AppendSupportCaseMessageRequest DTO.
type AppendSupportCaseMessageRequest struct {
	RequestBase
	CaseId  string `json:"caseId,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetSupportCase DTO.
type GetSupportCase struct {
	RequestBase
	CaseId string `json:"caseId,omitempty"`
}

// GetSupportCases DTO.
type GetSupportCases struct {
	RequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// OpenSupportCaseRequest DTO.
type OpenSupportCaseRequest struct {
	RequestBase
	Kind      string `json:"kind,omitempty"`
	Severity  string `json:"severity,omitempty"`
	Subject   string `json:"subject,omitempty"`
	Message   string `json:"message,omitempty"`
	ProjectId string `json:"projectId,omitempty"`
}

// SupportCaseOpened DTO.
type SupportCaseOpened struct {
	CaseId         *SupportCaseId      `json:"caseId,omitempty"`
	AccountId      *AccountId          `json:"accountId,omitempty"`
	ProjectId      *ProjectId          `json:"projectId,omitempty"`
	ReporterId     string              `json:"reporterId,omitempty"`
	Kind           SupportCaseKind     `json:"kind,omitempty"`
	Severity       SupportCaseSeverity `json:"severity,omitempty"`
	Subject        string              `json:"subject,omitempty"`
	DeploymentMode DeploymentMode      `json:"deploymentMode,omitempty"`
	GatewayVersion string              `json:"gatewayVersion,omitempty"`
	Region         string              `json:"region,omitempty"`
	PlanTier       string              `json:"planTier,omitempty"`
	OpenedOn       *UtcDateTime        `json:"openedOn,omitempty"`
}

// SupportCaseTriaged DTO.
type SupportCaseTriaged struct {
	CaseId         *SupportCaseId      `json:"caseId,omitempty"`
	Kind           SupportCaseKind     `json:"kind,omitempty"`
	Severity       SupportCaseSeverity `json:"severity,omitempty"`
	AffectedModule string              `json:"affectedModule,omitempty"`
	TriagedBy      string              `json:"triagedBy,omitempty"`
	TriagedOn      *UtcDateTime        `json:"triagedOn,omitempty"`
}

// SupportCaseMessageAppended DTO.
type SupportCaseMessageAppended struct {
	CaseId  *SupportCaseId     `json:"caseId,omitempty"`
	Message *SupportMessageRef `json:"message,omitempty"`
}

// SupportCaseStatusChanged DTO.
type SupportCaseStatusChanged struct {
	CaseId    *SupportCaseId    `json:"caseId,omitempty"`
	From      SupportCaseStatus `json:"from,omitempty"`
	To        SupportCaseStatus `json:"to,omitempty"`
	ChangedOn *UtcDateTime      `json:"changedOn,omitempty"`
}

// SupportCaseResolved DTO.
type SupportCaseResolved struct {
	CaseId     *SupportCaseId  `json:"caseId,omitempty"`
	Resolution *CaseResolution `json:"resolution,omitempty"`
	ResolvedOn *UtcDateTime    `json:"resolvedOn,omitempty"`
}

// SupportCaseClosed DTO.
type SupportCaseClosed struct {
	CaseId   *SupportCaseId         `json:"caseId,omitempty"`
	ClosedBy string                 `json:"closedBy,omitempty"`
	ClosedOn *UtcDateTime           `json:"closedOn,omitempty"`
	Reason   SupportCaseCloseReason `json:"reason,omitempty"`
}

// SupportCaseReopened DTO.
type SupportCaseReopened struct {
	CaseId     *SupportCaseId `json:"caseId,omitempty"`
	Reason     string         `json:"reason,omitempty"`
	ReopenedOn *UtcDateTime   `json:"reopenedOn,omitempty"`
}

// SupportCaseWaitingReminderSent DTO.
type SupportCaseWaitingReminderSent struct {
	CaseId   *SupportCaseId `json:"caseId,omitempty"`
	TierDays float64        `json:"tierDays,omitempty"`
	SentOn   *UtcDateTime   `json:"sentOn,omitempty"`
}

// SupportCaseAttachmentLinked DTO.
type SupportCaseAttachmentLinked struct {
	CaseId        *SupportCaseId `json:"caseId,omitempty"`
	AttachmentRef string         `json:"attachmentRef,omitempty"`
	FileName      string         `json:"fileName,omitempty"`
	LinkedOn      *UtcDateTime   `json:"linkedOn,omitempty"`
}

// GetDiagnosticPacks DTO.
type GetDiagnosticPacks struct {
	CodeMashRequestBase
}

// RunDiagnosticPackRequest DTO.
type RunDiagnosticPackRequest struct {
	CodeMashRequestBase
	PackName    string  `json:"packName,omitempty"`
	PackVersion float64 `json:"packVersion,omitempty"`
	CaseId      string  `json:"caseId,omitempty"`
}

// GetDiagnosticEcho DTO.
type GetDiagnosticEcho struct {
	CodeMashRequestBase
	CaseId string `json:"caseId,omitempty"`
}

// ReadDiagnosticEventsRequest DTO.
type ReadDiagnosticEventsRequest struct {
	CodeMashRequestBase
	Stream string  `json:"stream,omitempty"`
	From   float64 `json:"from,omitempty"`
	Count  float64 `json:"count,omitempty"`
	CaseId string  `json:"caseId,omitempty"`
}

// QueryDiagnosticLogsRequest DTO.
type QueryDiagnosticLogsRequest struct {
	CodeMashListPaginationRequestBase
	Level            string `json:"level,omitempty"`
	Module           string `json:"module,omitempty"`
	LogCorrelationId string `json:"logCorrelationId,omitempty"`
	EventCode        string `json:"eventCode,omitempty"`
	Search           string `json:"search,omitempty"`
	FromUtc          string `json:"fromUtc,omitempty"`
	ToUtc            string `json:"toUtc,omitempty"`
	CaseId           string `json:"caseId,omitempty"`
}

// InspectDiagnosticRedisRequest DTO.
type InspectDiagnosticRedisRequest struct {
	CodeMashRequestBase
	KeyPattern string `json:"keyPattern,omitempty"`
	CaseId     string `json:"caseId,omitempty"`
}

// RunDiagnosticHealthCheckRequest DTO.
type RunDiagnosticHealthCheckRequest struct {
	CodeMashRequestBase
	CheckId string `json:"checkId,omitempty"`
	CaseId  string `json:"caseId,omitempty"`
}

// Authenticate DTO.
type Authenticate struct {
	Provider          string            `json:"provider,omitempty"`
	UserName          string            `json:"userName,omitempty"`
	Password          string            `json:"password,omitempty"`
	RememberMe        bool              `json:"rememberMe,omitempty"`
	AccessToken       string            `json:"accessToken,omitempty"`
	AccessTokenSecret string            `json:"accessTokenSecret,omitempty"`
	ReturnUrl         string            `json:"returnUrl,omitempty"`
	ErrorView         string            `json:"errorView,omitempty"`
	Meta              map[string]string `json:"meta,omitempty"`
}

// GetAccessToken DTO.
type GetAccessToken struct {
	RefreshToken string            `json:"refreshToken,omitempty"`
	Meta         map[string]string `json:"meta,omitempty"`
}
