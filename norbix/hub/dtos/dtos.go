// Code generated from the Norbix Hub contract. DO NOT EDIT.
//
// Regenerate with the type generator described in README.md ("Regenerate the
// types"). Editing this file by hand is lost on the next run.

// Package dtos holds the data types of the Norbix Hub surface. A field
// that holds another generated type is a pointer. A nullable true/false or
// number field is a pointer too (*bool, *float64): nil is left out of the
// request, false and 0 are sent. A required true/false field is always sent.
// Every other field carries omitempty, because the gateway leaves absent
// values out.
package dtos

// EmailProvider enum.
type EmailProvider string

const (
	EmailProviderSmtp               EmailProvider = "Smtp"
	EmailProviderSendGrid           EmailProvider = "SendGrid"
	EmailProviderMailGun            EmailProvider = "MailGun"
	EmailProviderAwsSes             EmailProvider = "AwsSes"
	EmailProviderFake               EmailProvider = "Fake"
	EmailProviderAzureCommunication EmailProvider = "AzureCommunication"
)

// SmtpPorts enum.
type SmtpPorts string

const (
	SmtpPortsDefault  SmtpPorts = "Default"
	SmtpPortsSsl      SmtpPorts = "Ssl"
	SmtpPortsTls      SmtpPorts = "Tls"
	SmtpPortsFallback SmtpPorts = "Fallback"
)

// CloudAuthMode enum.
type CloudAuthMode string

const (
	CloudAuthModeAccessKey CloudAuthMode = "AccessKey"
	CloudAuthModeKeyless   CloudAuthMode = "Keyless"
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
	TokenMappingResolverTypeNotSet            TokenMappingResolverType = "NotSet"
	TokenMappingResolverTypeCustom            TokenMappingResolverType = "Custom"
	TokenMappingResolverTypeProject           TokenMappingResolverType = "Project"
	TokenMappingResolverTypeProjectSocials    TokenMappingResolverType = "ProjectSocials"
	TokenMappingResolverTypeInitiator         TokenMappingResolverType = "Initiator"
	TokenMappingResolverTypeRecipient         TokenMappingResolverType = "Recipient"
	TokenMappingResolverTypeSchemaRecord      TokenMappingResolverType = "SchemaRecord"
	TokenMappingResolverTypeTargetUser        TokenMappingResolverType = "TargetUser"
	TokenMappingResolverTypeTagDefinitions    TokenMappingResolverType = "TagDefinitions"
	TokenMappingResolverTypeEmailSignatures   TokenMappingResolverType = "EmailSignatures"
	TokenMappingResolverTypeCampaign          TokenMappingResolverType = "Campaign"
	TokenMappingResolverTypeTemplate          TokenMappingResolverType = "Template"
	TokenMappingResolverTypeEmailFooters      TokenMappingResolverType = "EmailFooters"
	TokenMappingResolverTypeOld               TokenMappingResolverType = "Old"
	TokenMappingResolverTypeNew               TokenMappingResolverType = "New"
	TokenMappingResolverTypeRequestInput      TokenMappingResolverType = "RequestInput"
	TokenMappingResolverTypeIntegrationConfig TokenMappingResolverType = "IntegrationConfig"
	TokenMappingResolverTypeIntegrationSecret TokenMappingResolverType = "IntegrationSecret"
	TokenMappingResolverTypeStepOutput        TokenMappingResolverType = "StepOutput"
)

// SecretValueFormat enum.
type SecretValueFormat string

const (
	SecretValueFormatRaw      SecretValueFormat = "Raw"
	SecretValueFormatBearer   SecretValueFormat = "Bearer"
	SecretValueFormatBasic    SecretValueFormat = "Basic"
	SecretValueFormatPrefixed SecretValueFormat = "Prefixed"
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
	TriggerTypeAi         TriggerType = "Ai"
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
	TriggerActionTypeWorkflow    TriggerActionType = "Workflow"
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
	SchemaTriggerTypeOnInserted     SchemaTriggerType = "OnInserted"
	SchemaTriggerTypeOnDeleted      SchemaTriggerType = "OnDeleted"
	SchemaTriggerTypeOnUpdated      SchemaTriggerType = "OnUpdated"
	SchemaTriggerTypeOnInsertedMany SchemaTriggerType = "OnInsertedMany"
	SchemaTriggerTypeOnUpdatedMany  SchemaTriggerType = "OnUpdatedMany"
	SchemaTriggerTypeOnDeletedMany  SchemaTriggerType = "OnDeletedMany"
)

// FilesTriggerType enum.
type FilesTriggerType string

const (
	FilesTriggerTypeOnFileUploaded FilesTriggerType = "OnFileUploaded"
	FilesTriggerTypeOnFileDeleted  FilesTriggerType = "OnFileDeleted"
)

// PaymentTriggerType enum.
type PaymentTriggerType string

const (
	PaymentTriggerTypeOnOrderCreated        PaymentTriggerType = "OnOrderCreated"
	PaymentTriggerTypeOnOrderPaid           PaymentTriggerType = "OnOrderPaid"
	PaymentTriggerTypeOnWebhookCallReceived PaymentTriggerType = "OnWebhookCallReceived"
	PaymentTriggerTypeOnOrderPaymentChanged PaymentTriggerType = "OnOrderPaymentChanged"
)

// DatabaseProvider enum.
type DatabaseProvider string

const (
	DatabaseProviderMongoDbConnectionString         DatabaseProvider = "MongoDbConnectionString"
	DatabaseProviderCodeMashMongoDbAtlasFlexManaged DatabaseProvider = "CodeMashMongoDbAtlasFlexManaged"
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
	PaymentGatewayPlatformMontonio     PaymentGatewayPlatform = "Montonio"
	PaymentGatewayPlatformNeopay       PaymentGatewayPlatform = "Neopay"
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
	CodeProviderFake                 CodeProvider = "Fake"
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

// CodeInvocationType enum.
type CodeInvocationType string

const (
	CodeInvocationTypeSync  CodeInvocationType = "Sync"
	CodeInvocationTypeAsync CodeInvocationType = "Async"
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

// PushDeviceDeliveryFamily enum.
type PushDeviceDeliveryFamily string

const (
	PushDeviceDeliveryFamilyIos     PushDeviceDeliveryFamily = "Ios"
	PushDeviceDeliveryFamilyAndroid PushDeviceDeliveryFamily = "Android"
	PushDeviceDeliveryFamilyChrome  PushDeviceDeliveryFamily = "Chrome"
	PushDeviceDeliveryFamilySafari  PushDeviceDeliveryFamily = "Safari"
	PushDeviceDeliveryFamilyExpo    PushDeviceDeliveryFamily = "Expo"
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

// SchedulerTaskType enum.
type SchedulerTaskType string

const (
	SchedulerTaskTypeEmailCampaign           SchedulerTaskType = "EmailCampaign"
	SchedulerTaskTypePushCampaign            SchedulerTaskType = "PushCampaign"
	SchedulerTaskTypeSmsCampaign             SchedulerTaskType = "SmsCampaign"
	SchedulerTaskTypeCodeFunctionalCall      SchedulerTaskType = "CodeFunctionalCall"
	SchedulerTaskTypeWebhookCall             SchedulerTaskType = "WebhookCall"
	SchedulerTaskTypeMarketplaceFunctionCall SchedulerTaskType = "MarketplaceFunctionCall"
	SchedulerTaskTypeWorkflowCall            SchedulerTaskType = "WorkflowCall"
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

// MarketplaceAvailability enum.
type MarketplaceAvailability string

const (
	MarketplaceAvailabilityDevelopment MarketplaceAvailability = "Development"
	MarketplaceAvailabilityStaging     MarketplaceAvailability = "Staging"
	MarketplaceAvailabilityProduction  MarketplaceAvailability = "Production"
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

// MarketplaceFormKind enum.
type MarketplaceFormKind string

const (
	MarketplaceFormKindGeneric        MarketplaceFormKind = "Generic"
	MarketplaceFormKindDatabaseQuery  MarketplaceFormKind = "DatabaseQuery"
	MarketplaceFormKindDatabaseWrite  MarketplaceFormKind = "DatabaseWrite"
	MarketplaceFormKindMessage        MarketplaceFormKind = "Message"
	MarketplaceFormKindPrompt         MarketplaceFormKind = "Prompt"
	MarketplaceFormKindWeb            MarketplaceFormKind = "Web"
	MarketplaceFormKindMembershipUser MarketplaceFormKind = "MembershipUser"
)

// CodeFunctionSource enum.
type CodeFunctionSource string

const (
	CodeFunctionSourceDatabase   CodeFunctionSource = "Database"
	CodeFunctionSourceMembership CodeFunctionSource = "Membership"
	CodeFunctionSourceFiles      CodeFunctionSource = "Files"
	CodeFunctionSourcePayments   CodeFunctionSource = "Payments"
	CodeFunctionSourceAi         CodeFunctionSource = "Ai"
	CodeFunctionSourceScheduler  CodeFunctionSource = "Scheduler"
	CodeFunctionSourceManual     CodeFunctionSource = "Manual"
	CodeFunctionSourceWorkflow   CodeFunctionSource = "Workflow"
)

// CodeFunctionCallerKind enum.
type CodeFunctionCallerKind string

const (
	CodeFunctionCallerKindUser   CodeFunctionCallerKind = "User"
	CodeFunctionCallerKindApiKey CodeFunctionCallerKind = "ApiKey"
	CodeFunctionCallerKindSystem CodeFunctionCallerKind = "System"
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

// AiPlanQuotaUnit enum.
type AiPlanQuotaUnit string

const (
	AiPlanQuotaUnitNone    AiPlanQuotaUnit = "None"
	AiPlanQuotaUnitCredits AiPlanQuotaUnit = "Credits"
	AiPlanQuotaUnitTokens  AiPlanQuotaUnit = "Tokens"
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

// EmailValidationProvider enum.
type EmailValidationProvider string

const (
	EmailValidationProviderZeroBounce      EmailValidationProvider = "ZeroBounce"
	EmailValidationProviderNeverBounce     EmailValidationProvider = "NeverBounce"
	EmailValidationProviderBouncer         EmailValidationProvider = "Bouncer"
	EmailValidationProviderMailgunValidate EmailValidationProvider = "MailgunValidate"
)

// MarketplaceAccessLevel enum.
type MarketplaceAccessLevel string

const (
	MarketplaceAccessLevelNone          MarketplaceAccessLevel = "None"
	MarketplaceAccessLevelViewer        MarketplaceAccessLevel = "Viewer"
	MarketplaceAccessLevelFunctionsOnly MarketplaceAccessLevel = "FunctionsOnly"
	MarketplaceAccessLevelReadOnly      MarketplaceAccessLevel = "ReadOnly"
	MarketplaceAccessLevelFullAccess    MarketplaceAccessLevel = "FullAccess"
	MarketplaceAccessLevelManage        MarketplaceAccessLevel = "Manage"
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

// EmbeddingProvider enum.
type EmbeddingProvider string

const (
	EmbeddingProviderVoyage EmbeddingProvider = "Voyage"
	EmbeddingProviderOpenAI EmbeddingProvider = "OpenAI"
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
	ApplicationModuleDiagnostics  ApplicationModule = "Diagnostics"
)

// LanguageGapKind enum.
type LanguageGapKind string

const (
	LanguageGapKindTemplate  LanguageGapKind = "Template"
	LanguageGapKindFooter    LanguageGapKind = "Footer"
	LanguageGapKindSignature LanguageGapKind = "Signature"
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

// AuthStatus enum.
type AuthStatus string

const (
	AuthStatusRegistered        AuthStatus = "Registered"
	AuthStatusPendingValidation AuthStatus = "PendingValidation"
	AuthStatusInvited           AuthStatus = "Invited"
	AuthStatusActive            AuthStatus = "Active"
	AuthStatusUnregistered      AuthStatus = "Unregistered"
	AuthStatusSuspended         AuthStatus = "Suspended"
	AuthStatusInActive          AuthStatus = "InActive"
	AuthStatusBlocked           AuthStatus = "Blocked"
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

// NotificationDeliveryState enum.
type NotificationDeliveryState string

const (
	NotificationDeliveryStatePending   NotificationDeliveryState = "Pending"
	NotificationDeliveryStateSending   NotificationDeliveryState = "Sending"
	NotificationDeliveryStateRetrying  NotificationDeliveryState = "Retrying"
	NotificationDeliveryStateAccepted  NotificationDeliveryState = "Accepted"
	NotificationDeliveryStateFailed    NotificationDeliveryState = "Failed"
	NotificationDeliveryStateBlocked   NotificationDeliveryState = "Blocked"
	NotificationDeliveryStateCancelled NotificationDeliveryState = "Cancelled"
)

// NotificationAttemptOutcome enum.
type NotificationAttemptOutcome string

const (
	NotificationAttemptOutcomeAccepted NotificationAttemptOutcome = "Accepted"
	NotificationAttemptOutcomeFailed   NotificationAttemptOutcome = "Failed"
)

// SendErrorCategory enum.
type SendErrorCategory string

const (
	SendErrorCategoryUnknown   SendErrorCategory = "Unknown"
	SendErrorCategoryTransient SendErrorCategory = "Transient"
	SendErrorCategoryPermanent SendErrorCategory = "Permanent"
	SendErrorCategoryAmbiguous SendErrorCategory = "Ambiguous"
)

// NotificationAttemptTrigger enum.
type NotificationAttemptTrigger string

const (
	NotificationAttemptTriggerInitial     NotificationAttemptTrigger = "Initial"
	NotificationAttemptTriggerAutoRetry   NotificationAttemptTrigger = "AutoRetry"
	NotificationAttemptTriggerManualRetry NotificationAttemptTrigger = "ManualRetry"
)

// DeliveryEventOutcome enum.
type DeliveryEventOutcome string

const (
	DeliveryEventOutcomeDelivered  DeliveryEventOutcome = "Delivered"
	DeliveryEventOutcomeBounced    DeliveryEventOutcome = "Bounced"
	DeliveryEventOutcomeComplained DeliveryEventOutcome = "Complained"
	DeliveryEventOutcomeRejected   DeliveryEventOutcome = "Rejected"
	DeliveryEventOutcomeDeferred   DeliveryEventOutcome = "Deferred"
)

// CodeFunctionQualifierKindDto enum.
type CodeFunctionQualifierKindDto string

const (
	CodeFunctionQualifierKindDtoAlias   CodeFunctionQualifierKindDto = "Alias"
	CodeFunctionQualifierKindDtoVersion CodeFunctionQualifierKindDto = "Version"
)

// AiTriggerType enum.
type AiTriggerType string

const (
	AiTriggerTypeOnCreditsWarning   AiTriggerType = "OnCreditsWarning"
	AiTriggerTypeOnCreditsExhausted AiTriggerType = "OnCreditsExhausted"
	AiTriggerTypeOnQuotaWarning     AiTriggerType = "OnQuotaWarning"
	AiTriggerTypeOnQuotaExhausted   AiTriggerType = "OnQuotaExhausted"
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

// SchedulerTaskRunOutcome enum.
type SchedulerTaskRunOutcome string

const (
	SchedulerTaskRunOutcomeFired   SchedulerTaskRunOutcome = "Fired"
	SchedulerTaskRunOutcomeFailed  SchedulerTaskRunOutcome = "Failed"
	SchedulerTaskRunOutcomeSkipped SchedulerTaskRunOutcome = "Skipped"
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

// SupportCustomerStatus enum.
type SupportCustomerStatus string

const (
	SupportCustomerStatusPending SupportCustomerStatus = "Pending"
	SupportCustomerStatusOpen    SupportCustomerStatus = "Open"
	SupportCustomerStatusSolved  SupportCustomerStatus = "Solved"
	SupportCustomerStatusClosed  SupportCustomerStatus = "Closed"
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
	IsEnabled       bool          `json:"isEnabled"`
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
	Provider              EmailProvider `json:"provider,omitempty"`
	AuthMode              CloudAuthMode `json:"authMode,omitempty"`
	AwsRegion             string        `json:"awsRegion,omitempty"`
	EmailIdentityArn      string        `json:"emailIdentityArn,omitempty"`
	ConfigurationSet      string        `json:"configurationSet,omitempty"`
	RoleArn               string        `json:"roleArn,omitempty"`
	AccessKey             string        `json:"accessKey,omitempty"`
	SecretKey             string        `json:"secretKey,omitempty"`
	AcceptKeyRisk         *bool         `json:"acceptKeyRisk,omitempty"`
	DeliveryEventsEnabled *bool         `json:"deliveryEventsEnabled,omitempty"`
	SnsTopicArn           string        `json:"snsTopicArn,omitempty"`
}

// SendGridEmailIntegrationRequest DTO.
type SendGridEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider              EmailProvider `json:"provider,omitempty"`
	ApiKey                string        `json:"apiKey,omitempty"`
	DeliveryEventsEnabled *bool         `json:"deliveryEventsEnabled,omitempty"`
	EventWebhookPublicKey string        `json:"eventWebhookPublicKey,omitempty"`
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

// AzureCommunicationEmailIntegrationRequest DTO.
type AzureCommunicationEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider              EmailProvider `json:"provider,omitempty"`
	AuthMode              CloudAuthMode `json:"authMode,omitempty"`
	Endpoint              string        `json:"endpoint,omitempty"`
	TenantId              string        `json:"tenantId,omitempty"`
	ClientId              string        `json:"clientId,omitempty"`
	AcceptKeyRisk         *bool         `json:"acceptKeyRisk,omitempty"`
	AccessKey             string        `json:"accessKey,omitempty"`
	DeliveryEventsEnabled *bool         `json:"deliveryEventsEnabled,omitempty"`
	EventGridKey          string        `json:"eventGridKey,omitempty"`
}

// FakeEmailIntegrationRequest DTO.
type FakeEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider EmailProvider `json:"provider,omitempty"`
}

// TokenMappingDto DTO.
type TokenMappingDto struct {
	Key          string                   `json:"key,omitempty"`
	Value        string                   `json:"value,omitempty"`
	Resolver     TokenMappingResolverType `json:"resolver,omitempty"`
	SecretKeys   []string                 `json:"secretKeys,omitempty"`
	SecretFormat SecretValueFormat        `json:"secretFormat,omitempty"`
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
	CampaignTime            *float64                           `json:"campaignTime,omitempty"`
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
	SingleEmailStrategy bool                               `json:"singleEmailStrategy"`
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
	SingleEmailStrategy bool                               `json:"singleEmailStrategy"`
}

// EmailToUsersDeliverySettingsRequest DTO.
type EmailToUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	UserRecipients      []string                           `json:"userRecipients,omitempty"`
	UserCc              []string                           `json:"userCc,omitempty"`
	UserBcc             []string                           `json:"userBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy"`
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
	IsEnabled      bool              `json:"isEnabled"`
	PreExecuteCode string            `json:"preExecuteCode,omitempty"`
	Order          *float64          `json:"order,omitempty"`
	BreakOnError   bool              `json:"breakOnError"`
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
	IsEnabled       bool             `json:"isEnabled"`
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
	IsEnabled       bool         `json:"isEnabled"`
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
	UseSsl   bool         `json:"useSsl"`
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
	Provider      FileProvider  `json:"provider,omitempty"`
	AuthMode      CloudAuthMode `json:"authMode,omitempty"`
	BucketName    string        `json:"bucketName,omitempty"`
	Region        string        `json:"region,omitempty"`
	RoleArn       string        `json:"roleArn,omitempty"`
	AccessKey     string        `json:"accessKey,omitempty"`
	SecretKey     string        `json:"secretKey,omitempty"`
	AcceptKeyRisk *bool         `json:"acceptKeyRisk,omitempty"`
}

// GoogleCloudFilesIntegrationRequest DTO.
type GoogleCloudFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider              FileProvider  `json:"provider,omitempty"`
	AuthMode              CloudAuthMode `json:"authMode,omitempty"`
	BucketName            string        `json:"bucketName,omitempty"`
	ProjectNumber         string        `json:"projectNumber,omitempty"`
	PoolId                string        `json:"poolId,omitempty"`
	ProviderId            string        `json:"providerId,omitempty"`
	ServiceAccountEmail   string        `json:"serviceAccountEmail,omitempty"`
	ServiceAccountJsonKey string        `json:"serviceAccountJsonKey,omitempty"`
	AcceptKeyRisk         *bool         `json:"acceptKeyRisk,omitempty"`
}

// AzureBlobFilesIntegrationRequest DTO.
type AzureBlobFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider         FileProvider  `json:"provider,omitempty"`
	AuthMode         CloudAuthMode `json:"authMode,omitempty"`
	BlobName         string        `json:"blobName,omitempty"`
	StorageAccount   string        `json:"storageAccount,omitempty"`
	TenantId         string        `json:"tenantId,omitempty"`
	ClientId         string        `json:"clientId,omitempty"`
	ConnectionString string        `json:"connectionString,omitempty"`
	AcceptKeyRisk    *bool         `json:"acceptKeyRisk,omitempty"`
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
	IsEnabled       bool            `json:"isEnabled"`
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
	IsEnabled       bool               `json:"isEnabled"`
}

// DisplayName DTO.
type DisplayName struct {
	Value string `json:"value,omitempty"`
}

// RoleName DTO.
type RoleName struct {
	Name                string `json:"name,omitempty"`
	DisplayName         string `json:"displayName,omitempty"`
	IsAdministrator     bool   `json:"isAdministrator"`
	IsAuthenticated     bool   `json:"isAuthenticated"`
	IsGuest             bool   `json:"isGuest"`
	IsRootRole          bool   `json:"isRootRole"`
	IsCollaboratorRole  bool   `json:"isCollaboratorRole"`
	IsProjectSystemRole bool   `json:"isProjectSystemRole"`
	IsAccountSystemRole bool   `json:"isAccountSystemRole"`
	IsSystemRole        bool   `json:"isSystemRole"`
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
	IsProduction bool               `json:"isProduction"`
	OAuthModes   []*OAuthModeConfig `json:"oAuthModes,omitempty"`
}

// PaymentIntegrationRequest DTO.
type PaymentIntegrationRequest struct {
	IntegrationId   string                 `json:"integrationId,omitempty"`
	Provider        PaymentGatewayPlatform `json:"provider,omitempty"`
	IntegrationName string                 `json:"integrationName,omitempty"`
	IsEnabled       bool                   `json:"isEnabled"`
}

// LemonSqueezyPaymentIntegrationRequest DTO.
type LemonSqueezyPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	StoreId              string                 `json:"storeId,omitempty"`
	ApiKey               string                 `json:"apiKey,omitempty"`
	WebhookSigningSecret string                 `json:"webhookSigningSecret,omitempty"`
	IsTestMode           bool                   `json:"isTestMode"`
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
	IsTestMode           bool                   `json:"isTestMode"`
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
	WebhookBundleId                       string                 `json:"webhookBundleId,omitempty"`
	MerchantIdentifier                    string                 `json:"merchantIdentifier,omitempty"`
	MerchantDomain                        string                 `json:"merchantDomain,omitempty"`
	DisplayName                           string                 `json:"displayName,omitempty"`
	MerchantIdentityCertificateP12Base64  string                 `json:"merchantIdentityCertificateP12Base64,omitempty"`
	MerchantIdentityCertificatePassword   string                 `json:"merchantIdentityCertificatePassword,omitempty"`
	PaymentProcessingCertificateP12Base64 string                 `json:"paymentProcessingCertificateP12Base64,omitempty"`
	PaymentProcessingCertificatePassword  string                 `json:"paymentProcessingCertificatePassword,omitempty"`
}

// GoogleInAppPaymentIntegrationRequest DTO.
type GoogleInAppPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider           PaymentGatewayPlatform `json:"provider,omitempty"`
	WebhookPackageName string                 `json:"webhookPackageName,omitempty"`
	MerchantId         string                 `json:"merchantId,omitempty"`
	MerchantName       string                 `json:"merchantName,omitempty"`
	Gateway            string                 `json:"gateway,omitempty"`
	PrivateKeyOrToken  string                 `json:"privateKeyOrToken,omitempty"`
	GatewayMerchantId  string                 `json:"gatewayMerchantId,omitempty"`
}

// MontonioPaymentIntegrationRequest DTO.
type MontonioPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider  PaymentGatewayPlatform `json:"provider,omitempty"`
	AccessKey string                 `json:"accessKey,omitempty"`
	SecretKey string                 `json:"secretKey,omitempty"`
	IsSandbox *bool                  `json:"isSandbox,omitempty"`
	Locale    string                 `json:"locale,omitempty"`
}

// NeopayPaymentIntegrationRequest DTO.
type NeopayPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider   PaymentGatewayPlatform `json:"provider,omitempty"`
	ProjectId  string                 `json:"projectId,omitempty"`
	ProjectKey string                 `json:"projectKey,omitempty"`
	IsSandbox  *bool                  `json:"isSandbox,omitempty"`
	Country    string                 `json:"country,omitempty"`
	Locale     string                 `json:"locale,omitempty"`
}

// ShopifyPaymentIntegrationRequest DTO.
type ShopifyPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider      PaymentGatewayPlatform `json:"provider,omitempty"`
	ShopDomain    string                 `json:"shopDomain,omitempty"`
	WebhookSecret string                 `json:"webhookSecret,omitempty"`
}

// WooCommercePaymentIntegrationRequest DTO.
type WooCommercePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider       PaymentGatewayPlatform `json:"provider,omitempty"`
	StoreUrl       string                 `json:"storeUrl,omitempty"`
	WebhookSecret  string                 `json:"webhookSecret,omitempty"`
	ConsumerKey    string                 `json:"consumerKey,omitempty"`
	ConsumerSecret string                 `json:"consumerSecret,omitempty"`
}

// MagentoPaymentIntegrationRequest DTO.
type MagentoPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider      PaymentGatewayPlatform `json:"provider,omitempty"`
	StoreUrl      string                 `json:"storeUrl,omitempty"`
	WebhookSecret string                 `json:"webhookSecret,omitempty"`
}

// BraintreePaymentIntegrationRequest DTO.
type BraintreePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider      PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantId    string                 `json:"merchantId,omitempty"`
	Environment   string                 `json:"environment,omitempty"`
	PrivateKey    string                 `json:"privateKey,omitempty"`
	WebhookSecret string                 `json:"webhookSecret,omitempty"`
}

// AuthorizeNetPaymentIntegrationRequest DTO.
type AuthorizeNetPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider            PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantLoginId     string                 `json:"merchantLoginId,omitempty"`
	Environment         string                 `json:"environment,omitempty"`
	TransactionKey      string                 `json:"transactionKey,omitempty"`
	WebhookSignatureKey string                 `json:"webhookSignatureKey,omitempty"`
}

// CheckOutComPaymentIntegrationRequest DTO.
type CheckOutComPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider        PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantAccount string                 `json:"merchantAccount,omitempty"`
	Environment     string                 `json:"environment,omitempty"`
	SecretKey       string                 `json:"secretKey,omitempty"`
	WebhookSecret   string                 `json:"webhookSecret,omitempty"`
}

// WorldpayPaymentIntegrationRequest DTO.
type WorldpayPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider      PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantCode  string                 `json:"merchantCode,omitempty"`
	Environment   string                 `json:"environment,omitempty"`
	ApiKey        string                 `json:"apiKey,omitempty"`
	WebhookSecret string                 `json:"webhookSecret,omitempty"`
}

// PushIntegrationRequest DTO.
type PushIntegrationRequest struct {
	IntegrationId   string       `json:"integrationId,omitempty"`
	Provider        PushProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName,omitempty"`
	IsEnabled       bool         `json:"isEnabled"`
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
	IsProduction bool         `json:"isProduction"`
}

// FakePushIntegrationRequest DTO.
type FakePushIntegrationRequest struct {
	PushIntegrationRequest
	Provider PushProvider `json:"provider,omitempty"`
}

// SmsIntegrationRequest DTO.
type SmsIntegrationRequest struct {
	IntegrationId   string      `json:"integrationId,omitempty"`
	Provider        SmsProvider `json:"provider,omitempty"`
	IntegrationName string      `json:"integrationName,omitempty"`
	IsEnabled       bool        `json:"isEnabled"`
}

// BirdSmsIntegrationRequest DTO.
type BirdSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider   SmsProvider `json:"provider,omitempty"`
	Originator string      `json:"originator,omitempty"`
	Region     string      `json:"region,omitempty"`
	ApiKey     string      `json:"apiKey,omitempty"`
}

// PlivoSmsIntegrationRequest DTO.
type PlivoSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider        SmsProvider `json:"provider,omitempty"`
	AuthId          string      `json:"authId,omitempty"`
	FromPhoneNumber string      `json:"fromPhoneNumber,omitempty"`
	AuthToken       string      `json:"authToken,omitempty"`
}

// SinchSmsIntegrationRequest DTO.
type SinchSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider        SmsProvider `json:"provider,omitempty"`
	ServicePlanId   string      `json:"servicePlanId,omitempty"`
	FromPhoneNumber string      `json:"fromPhoneNumber,omitempty"`
	ApiKey          string      `json:"apiKey,omitempty"`
	ApiSecret       string      `json:"apiSecret,omitempty"`
}

// TelesignSmsIntegrationRequest DTO.
type TelesignSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider   SmsProvider `json:"provider,omitempty"`
	CustomerId string      `json:"customerId,omitempty"`
	FromSender string      `json:"fromSender,omitempty"`
	ApiKey     string      `json:"apiKey,omitempty"`
}

// TelnyxSmsIntegrationRequest DTO.
type TelnyxSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider           SmsProvider `json:"provider,omitempty"`
	MessagingProfileId string      `json:"messagingProfileId,omitempty"`
	FromPhoneNumber    string      `json:"fromPhoneNumber,omitempty"`
	ApiKey             string      `json:"apiKey,omitempty"`
}

// TwilioSmsIntegrationRequest DTO.
type TwilioSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider              SmsProvider `json:"provider,omitempty"`
	AccountSid            string      `json:"accountSid,omitempty"`
	FromPhoneNumber       string      `json:"fromPhoneNumber,omitempty"`
	AuthToken             string      `json:"authToken,omitempty"`
	DeliveryEventsEnabled *bool       `json:"deliveryEventsEnabled,omitempty"`
	StatusCallbackUrl     string      `json:"statusCallbackUrl,omitempty"`
}

// VonageSmsIntegrationRequest DTO.
type VonageSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider   SmsProvider `json:"provider,omitempty"`
	ApiKey     string      `json:"apiKey,omitempty"`
	FromSender string      `json:"fromSender,omitempty"`
	ApiSecret  string      `json:"apiSecret,omitempty"`
}

// FakeSmsIntegrationRequest DTO.
type FakeSmsIntegrationRequest struct {
	SmsIntegrationRequest
	Provider SmsProvider `json:"provider,omitempty"`
}

// CodeIntegrationRequest DTO.
type CodeIntegrationRequest struct {
	IntegrationId   string       `json:"integrationId,omitempty"`
	Provider        CodeProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName,omitempty"`
	IsEnabled       bool         `json:"isEnabled"`
}

// AwsLambdaCodeIntegrationRequest DTO.
type AwsLambdaCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider      CodeProvider  `json:"provider,omitempty"`
	AuthMode      CloudAuthMode `json:"authMode,omitempty"`
	Region        string        `json:"region,omitempty"`
	RoleArn       string        `json:"roleArn,omitempty"`
	AccessKey     string        `json:"accessKey,omitempty"`
	SecretKey     string        `json:"secretKey,omitempty"`
	AcceptKeyRisk *bool         `json:"acceptKeyRisk,omitempty"`
}

// AzureFunctionsCodeIntegrationRequest DTO.
type AzureFunctionsCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider              CodeProvider  `json:"provider,omitempty"`
	AuthMode              CloudAuthMode `json:"authMode,omitempty"`
	FunctionAppName       string        `json:"functionAppName,omitempty"`
	ResourceGroup         string        `json:"resourceGroup,omitempty"`
	SubscriptionId        string        `json:"subscriptionId,omitempty"`
	TenantId              string        `json:"tenantId,omitempty"`
	ClientId              string        `json:"clientId,omitempty"`
	ConnectionStringOrKey string        `json:"connectionStringOrKey,omitempty"`
	AcceptKeyRisk         *bool         `json:"acceptKeyRisk,omitempty"`
}

// GoogleCloudFunctionsCodeIntegrationRequest DTO.
type GoogleCloudFunctionsCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider              CodeProvider  `json:"provider,omitempty"`
	AuthMode              CloudAuthMode `json:"authMode,omitempty"`
	ProjectId             string        `json:"projectId,omitempty"`
	Region                string        `json:"region,omitempty"`
	ProjectNumber         string        `json:"projectNumber,omitempty"`
	PoolId                string        `json:"poolId,omitempty"`
	ProviderId            string        `json:"providerId,omitempty"`
	ServiceAccountEmail   string        `json:"serviceAccountEmail,omitempty"`
	ServiceAccountJsonKey string        `json:"serviceAccountJsonKey,omitempty"`
	AcceptKeyRisk         *bool         `json:"acceptKeyRisk,omitempty"`
}

// LlmModelOptionRequest DTO.
type LlmModelOptionRequest struct {
	Id          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// LlmIntegrationRequest DTO.
type LlmIntegrationRequest struct {
	IntegrationId   string                   `json:"integrationId,omitempty"`
	Provider        LlmProvider              `json:"provider,omitempty"`
	IntegrationName string                   `json:"integrationName,omitempty"`
	IsEnabled       bool                     `json:"isEnabled"`
	Endpoint        string                   `json:"endpoint,omitempty"`
	DefaultModel    string                   `json:"defaultModel,omitempty"`
	IsDefault       bool                     `json:"isDefault"`
	Models          []*LlmModelOptionRequest `json:"models,omitempty"`
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
	IsEnabled       bool         `json:"isEnabled"`
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
	IsActive             bool                 `json:"isActive"`
	Tags                 []string             `json:"tags,omitempty"`
	Env                  string               `json:"env,omitempty"`
}

// EmailBodyDto DTO.
type EmailBodyDto struct {
	Structure      string              `json:"structure,omitempty"`
	Code           string              `json:"code,omitempty"`
	TemplateEngine EmailTemplateEngine `json:"templateEngine,omitempty"`
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
	SizeBytes        *float64         `json:"sizeBytes,omitempty"`
	Checksum         *FileChecksumDto `json:"checksum,omitempty"`
}

// FileResourceRefDto DTO.
type FileResourceRefDto struct {
	Resource      *FileResourceDto `json:"resource,omitempty"`
	IntegrationId string           `json:"integrationId,omitempty"`
	Provider      FileProvider     `json:"provider,omitempty"`
	Path          string           `json:"path,omitempty"`
	PublicUrl     string           `json:"publicUrl,omitempty"`
	IsPublic      bool             `json:"isPublic"`
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
	Body string `json:"body,omitempty"`
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
	HiddenSystemEmailTemplate bool                     `json:"hiddenSystemEmailTemplate"`
}

// EmailCampaignDeliverySettingsDto DTO.
type EmailCampaignDeliverySettingsDto struct {
	RecipientsSourceType    EmailCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto                 `json:"mappedTokens,omitempty"`
	CampaignTime            *float64                           `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings            `json:"respectTimeZoneSettings,omitempty"`
}

// TriggerActionEmailDto DTO.
type TriggerActionEmailDto struct {
	TriggerActionDto
	TemplateId       string                            `json:"templateId,omitempty"`
	DeliverySettings *EmailCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
	Language         string                            `json:"language,omitempty"`
	InitiatorId      string                            `json:"initiatorId,omitempty"`
}

// PushCampaignDeliverySettingsDto DTO.
type PushCampaignDeliverySettingsDto struct {
	RecipientsSourceType    PushCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto                `json:"mappedTokens,omitempty"`
	CampaignTime            *float64                          `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings           `json:"respectTimeZoneSettings,omitempty"`
}

// TriggerActionPushDto DTO.
type TriggerActionPushDto struct {
	TriggerActionDto
	TemplateId       string                           `json:"templateId,omitempty"`
	DeliverySettings *PushCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
	Language         string                           `json:"language,omitempty"`
	InitiatorId      string                           `json:"initiatorId,omitempty"`
}

// FunctionRefDto DTO.
type FunctionRefDto struct {
	IntegrationId string       `json:"integrationId,omitempty"`
	Provider      CodeProvider `json:"provider,omitempty"`
	FunctionId    string       `json:"functionId,omitempty"`
	Qualifier     string       `json:"qualifier,omitempty"`
	DisplayName   string       `json:"displayName,omitempty"`
}

// CodeDeliverySettingsDto DTO.
type CodeDeliverySettingsDto struct {
	MappedTokens []*TokenMappingDto `json:"mappedTokens,omitempty"`
}

// TriggerActionCodeDto DTO.
type TriggerActionCodeDto struct {
	TriggerActionDto
	FunctionRef      *FunctionRefDto          `json:"functionRef,omitempty"`
	RunAsUserId      string                   `json:"runAsUserId,omitempty"`
	InvocationType   CodeInvocationType       `json:"invocationType,omitempty"`
	Config           string                   `json:"config,omitempty"`
	DeliverySettings *CodeDeliverySettingsDto `json:"deliverySettings,omitempty"`
	FunctionId       string                   `json:"functionId,omitempty"`
	ActivationCode   string                   `json:"activationCode,omitempty"`
	Meta             map[string]string        `json:"meta,omitempty"`
}

// WebhookDeliverySettingsDto DTO.
type WebhookDeliverySettingsDto struct {
	DestinationIds    []string           `json:"destinationIds,omitempty"`
	EventName         string             `json:"eventName,omitempty"`
	ContentType       string             `json:"contentType,omitempty"`
	IncludeRawPayload *bool              `json:"includeRawPayload,omitempty"`
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
	CampaignTime            *float64                         `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings          `json:"respectTimeZoneSettings,omitempty"`
}

// TriggerActionSmsDto DTO.
type TriggerActionSmsDto struct {
	TriggerActionDto
	TemplateId       string                          `json:"templateId,omitempty"`
	DeliverySettings *SmsCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
	Language         string                          `json:"language,omitempty"`
	InitiatorId      string                          `json:"initiatorId,omitempty"`
}

// SseDeliverySettingsDto DTO.
type SseDeliverySettingsDto struct {
	Audience        string             `json:"audience,omitempty"`
	UserAuthIds     []string           `json:"userAuthIds,omitempty"`
	EventName       string             `json:"eventName,omitempty"`
	PayloadType     string             `json:"payloadType,omitempty"`
	PayloadTemplate string             `json:"payloadTemplate,omitempty"`
	Persist         bool               `json:"persist"`
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
	FunctionId  string            `json:"functionId,omitempty"`
	Payload     map[string]string `json:"payload,omitempty"`
	RunAsUserId string            `json:"runAsUserId,omitempty"`
}

// TriggerActionWorkflowDto DTO.
type TriggerActionWorkflowDto struct {
	TriggerActionDto
	WorkflowRef string            `json:"workflowRef,omitempty"`
	Payload     map[string]string `json:"payload,omitempty"`
}

// RequestBase DTO.
type RequestBase struct {
	TimeZoneId    string `json:"timeZoneId,omitempty"`
	Version       string `json:"version,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}

// CodeMashListPaginationRequestBase DTO.
type CodeMashListPaginationRequestBase struct {
	RequestBase
	ProjectId     string   `json:"projectId,omitempty"`
	Env           string   `json:"env,omitempty"`
	StartingAfter string   `json:"startingAfter,omitempty"`
	EndingBefore  string   `json:"endingBefore,omitempty"`
	PageSize      *float64 `json:"pageSize,omitempty"`
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
	IsSuccess bool        `json:"isSuccess"`
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
	SingleEmailStrategy bool     `json:"singleEmailStrategy"`
}

// EmailToUsersDeliverySettingsDto DTO.
type EmailToUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	UserRecipients      []string `json:"userRecipients,omitempty"`
	UserCc              []string `json:"userCc,omitempty"`
	UserBcc             []string `json:"userBcc,omitempty"`
	SingleEmailStrategy bool     `json:"singleEmailStrategy"`
}

// EmailToEmailAddressesDeliverySettingsDto DTO.
type EmailToEmailAddressesDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	Recipients          []string `json:"recipients,omitempty"`
	RecipientsCc        []string `json:"recipientsCc,omitempty"`
	RecipientsBcc       []string `json:"recipientsBcc,omitempty"`
	SingleEmailStrategy bool     `json:"singleEmailStrategy"`
}

// EmailToCollectionRecordsDeliverySettingsDto DTO.
type EmailToCollectionRecordsDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	Fields     []string                              `json:"fields,omitempty"`
	SchemaName string                                `json:"schemaName,omitempty"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
	RecordIds  []string                              `json:"recordIds,omitempty"`
}

// PushToAllUsersDeliverySettingsDto DTO.
type PushToAllUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	RolesNames []string                   `json:"rolesNames,omitempty"`
	UserTags   []string                   `json:"userTags,omitempty"`
	Platforms  []PushDeviceDeliveryFamily `json:"platforms,omitempty"`
}

// PushToUsersDeliverySettingsDto DTO.
type PushToUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Recipients []string `json:"recipients,omitempty"`
}

// PushToAccountUsersDeliverySettingsDto DTO.
type PushToAccountUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Recipients []string                   `json:"recipients,omitempty"`
	Platforms  []PushDeviceDeliveryFamily `json:"platforms,omitempty"`
}

// PushToCollectionRecordsDeliverySettingsDto DTO.
type PushToCollectionRecordsDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	Fields     []string                              `json:"fields,omitempty"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	SchemaName string                                `json:"schemaName,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
	RecordIds  []string                              `json:"recordIds,omitempty"`
}

// PushDeviceDeliveryTokenDto DTO.
type PushDeviceDeliveryTokenDto struct {
	PushDeviceToken string                   `json:"pushDeviceToken,omitempty"`
	DeliveryFamily  PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
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

// SmsToAccountUsersDeliverySettingsDto DTO.
type SmsToAccountUsersDeliverySettingsDto struct {
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
	RecordIds  []string                              `json:"recordIds,omitempty"`
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
	IsEnabled                         bool     `json:"isEnabled"`
	Env                               string   `json:"env,omitempty"`
	LastIntegrationTestAtUtc          string   `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      *bool    `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         []string `json:"lastIntegrationTestErrors,omitempty"`
	HumanDeliveryConfirmedAtUtc       string   `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool     `json:"requiresHumanDeliveryConfirmation"`
}

// LlmModelOptionDto DTO.
type LlmModelOptionDto struct {
	Id               string  `json:"id,omitempty"`
	DisplayName      string  `json:"displayName,omitempty"`
	InputCreditRate  float64 `json:"inputCreditRate,omitempty"`
	OutputCreditRate float64 `json:"outputCreditRate,omitempty"`
}

// LlmIntegrationDto DTO.
type LlmIntegrationDto struct {
	IntegrationDto
	Provider      LlmProvider          `json:"provider,omitempty"`
	BaseUrl       string               `json:"baseUrl,omitempty"`
	DefaultModel  string               `json:"defaultModel,omitempty"`
	IsConfigured  bool                 `json:"isConfigured"`
	IsSystemOwned bool                 `json:"isSystemOwned"`
	IsDefault     bool                 `json:"isDefault"`
	Models        []*LlmModelOptionDto `json:"models,omitempty"`
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
	IsConfigured  bool         `json:"isConfigured"`
	IsSystemOwned bool         `json:"isSystemOwned"`
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

// AwsLambdaKeylessCodeIntegrationDto DTO.
type AwsLambdaKeylessCodeIntegrationDto struct {
	CodeIntegrationDto
	Region   string        `json:"region,omitempty"`
	RoleArn  string        `json:"roleArn,omitempty"`
	AuthMode CloudAuthMode `json:"authMode,omitempty"`
}

// AwsLambdaIamCodeIntegrationDto DTO.
type AwsLambdaIamCodeIntegrationDto struct {
	CodeIntegrationDto
	Region   string        `json:"region,omitempty"`
	AuthMode CloudAuthMode `json:"authMode,omitempty"`
}

// AzureFunctionsCodeIntegrationDto DTO.
type AzureFunctionsCodeIntegrationDto struct {
	CodeIntegrationDto
	FunctionAppName string        `json:"functionAppName,omitempty"`
	ResourceGroup   string        `json:"resourceGroup,omitempty"`
	AuthMode        CloudAuthMode `json:"authMode,omitempty"`
	SubscriptionId  string        `json:"subscriptionId,omitempty"`
	TenantId        string        `json:"tenantId,omitempty"`
	ClientId        string        `json:"clientId,omitempty"`
}

// GoogleCloudFunctionsCodeIntegrationDto DTO.
type GoogleCloudFunctionsCodeIntegrationDto struct {
	CodeIntegrationDto
	ProjectId           string        `json:"projectId,omitempty"`
	Region              string        `json:"region,omitempty"`
	AuthMode            CloudAuthMode `json:"authMode,omitempty"`
	ProjectNumber       string        `json:"projectNumber,omitempty"`
	PoolId              string        `json:"poolId,omitempty"`
	ProviderId          string        `json:"providerId,omitempty"`
	ServiceAccountEmail string        `json:"serviceAccountEmail,omitempty"`
}

// PaymentsIntegrationConnectionDto DTO.
type PaymentsIntegrationConnectionDto struct {
	Provider          string `json:"provider,omitempty"`
	AccountId         string `json:"accountId,omitempty"`
	AccountName       string `json:"accountName,omitempty"`
	ConnectedBy       string `json:"connectedBy,omitempty"`
	ConnectedAtUtc    string `json:"connectedAtUtc,omitempty"`
	IsConnected       bool   `json:"isConnected"`
	DisconnectedAtUtc string `json:"disconnectedAtUtc,omitempty"`
	DisconnectReason  string `json:"disconnectReason,omitempty"`
}

// PaymentsIntegrationDto DTO.
type PaymentsIntegrationDto struct {
	IntegrationDto
	GatewayPlatform PaymentGatewayPlatform            `json:"gatewayPlatform,omitempty"`
	Connection      *PaymentsIntegrationConnectionDto `json:"connection,omitempty"`
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
	WebhookBundleId    string `json:"webhookBundleId,omitempty"`
	MerchantIdentifier string `json:"merchantIdentifier,omitempty"`
	MerchantDomain     string `json:"merchantDomain,omitempty"`
	DisplayName        string `json:"displayName,omitempty"`
}

// GoogleInAppPaymentIntegrationDto DTO.
type GoogleInAppPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	WebhookPackageName string `json:"webhookPackageName,omitempty"`
	MerchantId         string `json:"merchantId,omitempty"`
	MerchantName       string `json:"merchantName,omitempty"`
	Gateway            string `json:"gateway,omitempty"`
	GatewayMerchantId  string `json:"gatewayMerchantId,omitempty"`
}

// LemonSqueezyPaymentIntegrationDto DTO.
type LemonSqueezyPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreId    string `json:"storeId,omitempty"`
	IsTestMode bool   `json:"isTestMode"`
}

// MolliePaymentIntegrationDto DTO.
type MolliePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ProfileId  string `json:"profileId,omitempty"`
	IsTestMode bool   `json:"isTestMode"`
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
	ShopDomain string `json:"shopDomain,omitempty"`
}

// WooCommercePaymentIntegrationDto DTO.
type WooCommercePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreUrl string `json:"storeUrl,omitempty"`
}

// MagentoPaymentIntegrationDto DTO.
type MagentoPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreUrl string `json:"storeUrl,omitempty"`
}

// BraintreePaymentIntegrationDto DTO.
type BraintreePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantId  string `json:"merchantId,omitempty"`
	Environment string `json:"environment,omitempty"`
}

// AuthorizeNetPaymentIntegrationDto DTO.
type AuthorizeNetPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantLoginId string `json:"merchantLoginId,omitempty"`
	Environment     string `json:"environment,omitempty"`
}

// CheckOutComPaymentIntegrationDto DTO.
type CheckOutComPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantAccount string `json:"merchantAccount,omitempty"`
	Environment     string `json:"environment,omitempty"`
}

// WorldpayPaymentIntegrationDto DTO.
type WorldpayPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantCode string `json:"merchantCode,omitempty"`
	Environment  string `json:"environment,omitempty"`
}

// MontonioPaymentIntegrationDto DTO.
type MontonioPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	AccessKey string `json:"accessKey,omitempty"`
	IsSandbox bool   `json:"isSandbox"`
	Locale    string `json:"locale,omitempty"`
}

// NeopayPaymentIntegrationDto DTO.
type NeopayPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ProjectId string `json:"projectId,omitempty"`
	IsSandbox bool   `json:"isSandbox"`
	Country   string `json:"country,omitempty"`
	Locale    string `json:"locale,omitempty"`
}

// PaymentOrderLineDto DTO.
type PaymentOrderLineDto struct {
	Sku             string  `json:"sku,omitempty"`
	Title           string  `json:"title,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	UnitAmountMinor float64 `json:"unitAmountMinor,omitempty"`
	AmountMinor     float64 `json:"amountMinor,omitempty"`
}

// PaymentOrderCheckoutDto DTO.
type PaymentOrderCheckoutDto struct {
	Attempt           float64 `json:"attempt,omitempty"`
	IntegrationId     string  `json:"integrationId,omitempty"`
	Provider          string  `json:"provider,omitempty"`
	ProviderOrderId   string  `json:"providerOrderId,omitempty"`
	ProviderReference string  `json:"providerReference,omitempty"`
	PaymentUrl        string  `json:"paymentUrl,omitempty"`
	Method            string  `json:"method,omitempty"`
	StartedOnUtc      string  `json:"startedOnUtc,omitempty"`
	ExpiresAtUtc      string  `json:"expiresAtUtc,omitempty"`
}

// PaymentOrderDto DTO.
type PaymentOrderDto struct {
	Id                string                     `json:"id,omitempty"`
	Env               string                     `json:"env,omitempty"`
	BuyerUserId       string                     `json:"buyerUserId,omitempty"`
	MerchantReference string                     `json:"merchantReference,omitempty"`
	Currency          string                     `json:"currency,omitempty"`
	PriceVersion      string                     `json:"priceVersion,omitempty"`
	Lines             []*PaymentOrderLineDto     `json:"lines,omitempty"`
	TotalMinor        float64                    `json:"totalMinor,omitempty"`
	Metadata          map[string]string          `json:"metadata,omitempty"`
	Status            string                     `json:"status,omitempty"`
	CreatedOnUtc      string                     `json:"createdOnUtc,omitempty"`
	UpdatedOnUtc      string                     `json:"updatedOnUtc,omitempty"`
	Checkouts         []*PaymentOrderCheckoutDto `json:"checkouts,omitempty"`
	SettledMinor      *float64                   `json:"settledMinor,omitempty"`
	RefundedMinor     float64                    `json:"refundedMinor,omitempty"`
	PaidOnUtc         string                     `json:"paidOnUtc,omitempty"`
	ProviderPaymentId string                     `json:"providerPaymentId,omitempty"`
	Mismatches        []string                   `json:"mismatches,omitempty"`
}

// PaymentEventDto DTO.
type PaymentEventDto struct {
	IntegrationId   string            `json:"integrationId,omitempty"`
	Provider        string            `json:"provider,omitempty"`
	ProviderEventId string            `json:"providerEventId,omitempty"`
	OrderId         string            `json:"orderId,omitempty"`
	BuyerUserId     string            `json:"buyerUserId,omitempty"`
	EventName       string            `json:"eventName,omitempty"`
	Status          string            `json:"status,omitempty"`
	AmountMinor     *float64          `json:"amountMinor,omitempty"`
	Currency        string            `json:"currency,omitempty"`
	ExternalIds     map[string]string `json:"externalIds,omitempty"`
	OccurredOn      string            `json:"occurredOn,omitempty"`
	ReceivedOn      string            `json:"receivedOn,omitempty"`
}

// PaymentOrderEventDto DTO.
type PaymentOrderEventDto struct {
	OrderEventId string           `json:"orderEventId,omitempty"`
	Event        string           `json:"event,omitempty"`
	From         string           `json:"from,omitempty"`
	To           string           `json:"to,omitempty"`
	Reason       string           `json:"reason,omitempty"`
	Order        *PaymentOrderDto `json:"order,omitempty"`
	Payment      *PaymentEventDto `json:"payment,omitempty"`
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

// AwsS3KeylessFilesIntegrationDto DTO.
type AwsS3KeylessFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string        `json:"bucketName,omitempty"`
	Region     string        `json:"region,omitempty"`
	RoleArn    string        `json:"roleArn,omitempty"`
	AuthMode   CloudAuthMode `json:"authMode,omitempty"`
}

// AwsS3IamFilesIntegrationDto DTO.
type AwsS3IamFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string        `json:"bucketName,omitempty"`
	Region     string        `json:"region,omitempty"`
	AuthMode   CloudAuthMode `json:"authMode,omitempty"`
}

// AzureBlobFilesIntegrationDto DTO.
type AzureBlobFilesIntegrationDto struct {
	FilesIntegrationDto
	BlobName       string        `json:"blobName,omitempty"`
	AuthMode       CloudAuthMode `json:"authMode,omitempty"`
	StorageAccount string        `json:"storageAccount,omitempty"`
	TenantId       string        `json:"tenantId,omitempty"`
	ClientId       string        `json:"clientId,omitempty"`
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
	UseSsl   bool    `json:"useSsl"`
}

// GoogleCloudFilesIntegrationDto DTO.
type GoogleCloudFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName          string        `json:"bucketName,omitempty"`
	AuthMode            CloudAuthMode `json:"authMode,omitempty"`
	ProjectNumber       string        `json:"projectNumber,omitempty"`
	PoolId              string        `json:"poolId,omitempty"`
	ProviderId          string        `json:"providerId,omitempty"`
	ServiceAccountEmail string        `json:"serviceAccountEmail,omitempty"`
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
	Provider         DatabaseProvider `json:"provider,omitempty"`
	IsSystemOwned    bool             `json:"isSystemOwned"`
	CanManageIndexes *bool            `json:"canManageIndexes,omitempty"`
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
	AccountSid            string `json:"accountSid,omitempty"`
	FromPhoneNumber       string `json:"fromPhoneNumber,omitempty"`
	DeliveryEventsEnabled *bool  `json:"deliveryEventsEnabled,omitempty"`
	StatusCallbackUrl     string `json:"statusCallbackUrl,omitempty"`
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
	TeamId       string `json:"teamId,omitempty"`
	AppBundleId  string `json:"appBundleId,omitempty"`
	KeyId        string `json:"keyId,omitempty"`
	IsProduction bool   `json:"isProduction"`
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
	Region                string `json:"region,omitempty"`
	IdentityArn           string `json:"identityArn,omitempty"`
	ConfigurationSetName  string `json:"configurationSetName,omitempty"`
	DeliveryEventsEnabled *bool  `json:"deliveryEventsEnabled,omitempty"`
	SnsTopicArn           string `json:"snsTopicArn,omitempty"`
}

// AwsKeylessEmailIntegrationDto DTO.
type AwsKeylessEmailIntegrationDto struct {
	AwsSesEmailIntegrationDto
	RoleArn  string        `json:"roleArn,omitempty"`
	AuthMode CloudAuthMode `json:"authMode,omitempty"`
}

// AwsIamEmailIntegrationDto DTO.
type AwsIamEmailIntegrationDto struct {
	AwsSesEmailIntegrationDto
	AuthMode CloudAuthMode `json:"authMode,omitempty"`
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
	DeliveryEventsEnabled *bool  `json:"deliveryEventsEnabled,omitempty"`
	EventWebhookPublicKey string `json:"eventWebhookPublicKey,omitempty"`
}

// SmtpEmailIntegrationDto DTO.
type SmtpEmailIntegrationDto struct {
	EmailIntegrationDto
	HostName string  `json:"hostName,omitempty"`
	Port     float64 `json:"port,omitempty"`
}

// AzureCommunicationEmailIntegrationDto DTO.
type AzureCommunicationEmailIntegrationDto struct {
	EmailIntegrationDto
	Endpoint              string        `json:"endpoint,omitempty"`
	DeliveryEventsEnabled *bool         `json:"deliveryEventsEnabled,omitempty"`
	AuthMode              CloudAuthMode `json:"authMode,omitempty"`
	TenantId              string        `json:"tenantId,omitempty"`
	ClientId              string        `json:"clientId,omitempty"`
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
	IsEnabled       bool              `json:"isEnabled"`
}

// WebhookIntegrationDto DTO.
type WebhookIntegrationDto struct {
	IntegrationDto
	IsConfigured bool                     `json:"isConfigured"`
	Destinations []*WebhookDestinationDto `json:"destinations,omitempty"`
	ExtraHeaders map[string]string        `json:"extraHeaders,omitempty"`
}

// SchedulerTaskDto DTO.
type SchedulerTaskDto struct {
	ProjectId         string            `json:"projectId,omitempty"`
	TaskId            string            `json:"taskId,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	Cron              string            `json:"cron,omitempty"`
	Type              SchedulerTaskType `json:"type,omitempty"`
	PayloadJson       string            `json:"payloadJson,omitempty"`
	InitiatorId       string            `json:"initiatorId,omitempty"`
	IsEnabled         bool              `json:"isEnabled"`
	StopOnError       bool              `json:"stopOnError"`
	Env               string            `json:"env,omitempty"`
	CreatedAtUnix     *float64          `json:"createdAtUnix,omitempty"`
	UpdatedAtUnix     *float64          `json:"updatedAtUnix,omitempty"`
	EndsAtUnix        *float64          `json:"endsAtUnix,omitempty"`
	HasEnded          bool              `json:"hasEnded"`
	IsFailing         bool              `json:"isFailing"`
	LastFailureReason string            `json:"lastFailureReason,omitempty"`
	LastFailedAtUnix  *float64          `json:"lastFailedAtUnix,omitempty"`
}

// SchedulerTaskRequest DTO.
type SchedulerTaskRequest struct {
	Type SchedulerTaskType `json:"type,omitempty"`
}

// EmailCampaignSchedulerTaskRequest DTO.
type EmailCampaignSchedulerTaskRequest struct {
	SchedulerTaskRequest
	Type                  SchedulerTaskType     `json:"type,omitempty"`
	Campaign              *EmailCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string                `json:"databaseIntegrationId,omitempty"`
}

// IdResponse DTO.
type IdResponse struct {
	ResponseBase
	Id       string   `json:"id,omitempty"`
	Status   string   `json:"status,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// CodeMashRequestBase DTO.
type CodeMashRequestBase struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Env       string `json:"env,omitempty"`
}

// CreateSmsCampaignRequest DTO.
type CreateSmsCampaignRequest struct {
	CodeMashRequestBase
	TemplateId            string                                     `json:"templateId,omitempty"`
	DatabaseIntegrationId string                                     `json:"databaseIntegrationId,omitempty"`
	IntegrationId         string                                     `json:"integrationId,omitempty"`
	Language              string                                     `json:"language,omitempty"`
	InitiatorId           string                                     `json:"initiatorId,omitempty"`
	DeliveryType          SmsCampaignRecipientsSourceTypes           `json:"deliveryType,omitempty"`
	AllUsers              *SmsToAllUsersDeliverySettingsDto          `json:"allUsers,omitempty"`
	SpecifiedUsers        *SmsToUsersDeliverySettingsDto             `json:"specifiedUsers,omitempty"`
	AccountUsers          *SmsToAccountUsersDeliverySettingsDto      `json:"accountUsers,omitempty"`
	Collection            *SmsToCollectionRecordsDeliverySettingsDto `json:"collection,omitempty"`
	PhoneNumbers          *SmsToPhoneNumbersDeliverySettingsDto      `json:"phoneNumbers,omitempty"`
	SendNow               *bool                                      `json:"sendNow,omitempty"`
}

// SmsCampaignSchedulerTaskRequest DTO.
type SmsCampaignSchedulerTaskRequest struct {
	SchedulerTaskRequest
	Type                  SchedulerTaskType         `json:"type,omitempty"`
	Campaign              *CreateSmsCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string                    `json:"databaseIntegrationId,omitempty"`
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
	CampaignTime  *float64                          `json:"campaignTime,omitempty"`
}

// PushCampaignSchedulerTaskRequest DTO.
type PushCampaignSchedulerTaskRequest struct {
	SchedulerTaskRequest
	Type                  SchedulerTaskType    `json:"type,omitempty"`
	Campaign              *PushCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string               `json:"databaseIntegrationId,omitempty"`
}

// MongoDbAggregateDto DTO.
type MongoDbAggregateDto struct {
	ViewId            string   `json:"viewId,omitempty"`
	DisplayName       string   `json:"displayName,omitempty"`
	Description       string   `json:"description,omitempty"`
	SchemaViewId      string   `json:"schemaViewId,omitempty"`
	Pipeline          string   `json:"pipeline,omitempty"`
	JoinedCollections []string `json:"joinedCollections,omitempty"`
}

// MarketplaceValidationStepResultDto DTO.
type MarketplaceValidationStepResultDto struct {
	Key                  string `json:"key,omitempty"`
	Label                string `json:"label,omitempty"`
	RequiredForFunctions bool   `json:"requiredForFunctions"`
	Succeeded            bool   `json:"succeeded"`
	Skipped              bool   `json:"skipped"`
	Message              string `json:"message,omitempty"`
	ErrorCode            string `json:"errorCode,omitempty"`
}

// MarketplaceIntegrationDto DTO.
type MarketplaceIntegrationDto struct {
	IntegrationDto
	ListingViewId       string                                `json:"listingViewId,omitempty"`
	Transport           MarketplaceTransport                  `json:"transport,omitempty"`
	Vendor              string                                `json:"vendor,omitempty"`
	Category            MarketplaceCategory                   `json:"category,omitempty"`
	Description         string                                `json:"description,omitempty"`
	Config              map[string]string                     `json:"config,omitempty"`
	LastValidationSteps []*MarketplaceValidationStepResultDto `json:"lastValidationSteps,omitempty"`
	SignedIn            bool                                  `json:"signedIn"`
	SignedInBy          string                                `json:"signedInBy,omitempty"`
	SignedInAt          string                                `json:"signedInAt,omitempty"`
	TokenExpiresAt      string                                `json:"tokenExpiresAt,omitempty"`
}

// MarketplaceFunctionParameterDto DTO.
type MarketplaceFunctionParameterDto struct {
	Name         string `json:"name,omitempty"`
	Type         string `json:"type,omitempty"`
	Description  string `json:"description,omitempty"`
	IsRequired   bool   `json:"isRequired"`
	DefaultValue string `json:"defaultValue,omitempty"`
}

// MarketplaceFunctionDto DTO.
type MarketplaceFunctionDto struct {
	ViewId            string                             `json:"viewId,omitempty"`
	IntegrationViewId string                             `json:"integrationViewId,omitempty"`
	FunctionKey       string                             `json:"functionKey,omitempty"`
	DisplayName       string                             `json:"displayName,omitempty"`
	Description       string                             `json:"description,omitempty"`
	IsEnabled         bool                               `json:"isEnabled"`
	RequestTemplate   string                             `json:"requestTemplate,omitempty"`
	Arguments         map[string]string                  `json:"arguments,omitempty"`
	Inputs            []*MarketplaceFunctionParameterDto `json:"inputs,omitempty"`
	MappedTokens      []*TokenMappingDto                 `json:"mappedTokens,omitempty"`
}

// MarketplaceFieldDefinitionDto DTO.
type MarketplaceFieldDefinitionDto struct {
	Key               string               `json:"key,omitempty"`
	Label             string               `json:"label,omitempty"`
	Description       string               `json:"description,omitempty"`
	DocumentationUrl  string               `json:"documentationUrl,omitempty"`
	Type              MarketplaceFieldType `json:"type,omitempty"`
	IsRequired        bool                 `json:"isRequired"`
	DefaultValue      string               `json:"defaultValue,omitempty"`
	Placeholder       string               `json:"placeholder,omitempty"`
	ValidationPattern string               `json:"validationPattern,omitempty"`
	AllowedValues     []string             `json:"allowedValues,omitempty"`
}

// MarketplaceInstallFormDto DTO.
type MarketplaceInstallFormDto struct {
	DataSchema string   `json:"dataSchema,omitempty"`
	UiSchema   string   `json:"uiSchema,omitempty"`
	SecretKeys []string `json:"secretKeys,omitempty"`
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
	IsRequired       bool                         `json:"isRequired"`
}

// MarketplaceHttpRequestSpecDto DTO.
type MarketplaceHttpRequestSpecDto struct {
	Method       string                         `json:"method,omitempty"`
	PathTemplate string                         `json:"pathTemplate,omitempty"`
	Parameters   []*MarketplaceParameterSpecDto `json:"parameters,omitempty"`
	ContentType  string                         `json:"contentType,omitempty"`
}

// MarketplaceValidationStepDto DTO.
type MarketplaceValidationStepDto struct {
	Key                  string                         `json:"key,omitempty"`
	Label                string                         `json:"label,omitempty"`
	Description          string                         `json:"description,omitempty"`
	RequiredForFunctions bool                           `json:"requiredForFunctions"`
	Probe                *MarketplaceHttpRequestSpecDto `json:"probe,omitempty"`
}

// MarketplaceFunctionDefinitionDto DTO.
type MarketplaceFunctionDefinitionDto struct {
	DefinitionId         string                             `json:"definitionId,omitempty"`
	FunctionKey          string                             `json:"functionKey,omitempty"`
	DisplayName          string                             `json:"displayName,omitempty"`
	Description          string                             `json:"description,omitempty"`
	Group                string                             `json:"group,omitempty"`
	FormKind             MarketplaceFormKind                `json:"formKind,omitempty"`
	Inputs               []*MarketplaceFunctionParameterDto `json:"inputs,omitempty"`
	Parameters           []*MarketplaceFunctionParameterDto `json:"parameters,omitempty"`
	RequestSchema        string                             `json:"requestSchema,omitempty"`
	RequestTemplate      string                             `json:"requestTemplate,omitempty"`
	Request              *MarketplaceHttpRequestSpecDto     `json:"request,omitempty"`
	DefaultTokenMappings []*TokenMappingDto                 `json:"defaultTokenMappings,omitempty"`
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
	VideoUrl         string                              `json:"videoUrl,omitempty"`
	Availability     MarketplaceAvailability             `json:"availability,omitempty"`
	ConnectSteps     []string                            `json:"connectSteps,omitempty"`
	OAuthRedirectUrl string                              `json:"oAuthRedirectUrl,omitempty"`
	IsOfficial       bool                                `json:"isOfficial"`
	Tags             []string                            `json:"tags,omitempty"`
	SpecVersion      float64                             `json:"specVersion,omitempty"`
	ConfigFields     []*MarketplaceFieldDefinitionDto    `json:"configFields,omitempty"`
	SecretFields     []*MarketplaceFieldDefinitionDto    `json:"secretFields,omitempty"`
	InstallForm      *MarketplaceInstallFormDto          `json:"installForm,omitempty"`
	ValidationSteps  []*MarketplaceValidationStepDto     `json:"validationSteps,omitempty"`
	Functions        []*MarketplaceFunctionDefinitionDto `json:"functions,omitempty"`
}

// AdminPortalModuleDto DTO.
type AdminPortalModuleDto struct {
	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Enabled     bool   `json:"enabled"`
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
	Kind          string   `json:"kind,omitempty"`
	RequirementId string   `json:"requirementId,omitempty"`
	SessionId     string   `json:"sessionId,omitempty"`
	EntryId       string   `json:"entryId,omitempty"`
	EntrySeq      *float64 `json:"entrySeq,omitempty"`
	ArtifactId    string   `json:"artifactId,omitempty"`
	Label         string   `json:"label,omitempty"`
	Step          *float64 `json:"step,omitempty"`
	Number        *float64 `json:"number,omitempty"`
	SourceKind    string   `json:"sourceKind,omitempty"`
	SourceId      string   `json:"sourceId,omitempty"`
	Score         *float64 `json:"score,omitempty"`
	Cited         *bool    `json:"cited,omitempty"`
}

// AssistantTextEntryWireDto DTO.
type AssistantTextEntryWireDto struct {
	AiChatEntryWireDto
	Kind        string                      `json:"kind,omitempty"`
	Text        string                      `json:"text,omitempty"`
	IsStreaming bool                        `json:"isStreaming"`
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
	AllowFreeText bool                           `json:"allowFreeText"`
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
	MaxIterations *float64 `json:"maxIterations,omitempty"`
	MaxToolCalls  float64  `json:"maxToolCalls,omitempty"`
}

// AiChatPlanStepWireDto DTO.
type AiChatPlanStepWireDto struct {
	N          float64                           `json:"n,omitempty"`
	Tool       string                            `json:"tool,omitempty"`
	Title      string                            `json:"title,omitempty"`
	Goal       string                            `json:"goal,omitempty"`
	Inputs     *AiChatPlanStepInputsWireDto      `json:"inputs,omitempty"`
	DependsOn  []float64                         `json:"dependsOn,omitempty"`
	Replaces   *float64                          `json:"replaces,omitempty"`
	Done       []*AiChatPlanStepDoneCheckWireDto `json:"done,omitempty"`
	Loop       *AiChatPlanStepLoopWireDto        `json:"loop,omitempty"`
	Difficulty *float64                          `json:"difficulty,omitempty"`
}

// PlanEntryWireDto DTO.
type PlanEntryWireDto struct {
	AiChatEntryWireDto
	Kind             string                   `json:"kind,omitempty"`
	Goal             string                   `json:"goal,omitempty"`
	Steps            []*AiChatPlanStepWireDto `json:"steps,omitempty"`
	Status           string                   `json:"status,omitempty"`
	Gate             *AiChatGateResultWireDto `json:"gate,omitempty"`
	Difficulty       *float64                 `json:"difficulty,omitempty"`
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

// CodeFunctionCallerDto DTO.
type CodeFunctionCallerDto struct {
	UserId string                 `json:"userId,omitempty"`
	Email  string                 `json:"email,omitempty"`
	Kind   CodeFunctionCallerKind `json:"kind,omitempty"`
}

// CodeFunctionRunAsDto DTO.
type CodeFunctionRunAsDto struct {
	UserId      string `json:"userId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	ExpiresAt   string `json:"expiresAt,omitempty"`
}

// CodeFunctionContextDto DTO.
type CodeFunctionContextDto struct {
	AccountId     string                 `json:"accountId,omitempty"`
	ProjectId     string                 `json:"projectId,omitempty"`
	Environment   string                 `json:"environment,omitempty"`
	Region        string                 `json:"region,omitempty"`
	HubUrl        string                 `json:"hubUrl,omitempty"`
	ApiUrl        string                 `json:"apiUrl,omitempty"`
	EchoUrl       string                 `json:"echoUrl,omitempty"`
	TriggerId     string                 `json:"triggerId,omitempty"`
	TaskId        string                 `json:"taskId,omitempty"`
	CorrelationId string                 `json:"correlationId,omitempty"`
	Caller        *CodeFunctionCallerDto `json:"caller,omitempty"`
	RunAs         *CodeFunctionRunAsDto  `json:"runAs,omitempty"`
	WorkflowId    string                 `json:"workflowId,omitempty"`
	RunId         string                 `json:"runId,omitempty"`
	StepId        string                 `json:"stepId,omitempty"`
}

// ResourceRefDto DTO.
type ResourceRefDto struct {
	ProjectId     string          `json:"projectId,omitempty"`
	IntegrationId string          `json:"integrationId,omitempty"`
	Kind          ResourceKindDto `json:"kind,omitempty"`
}

// DocumentRefDto DTO.
type DocumentRefDto struct {
	ResourceRefDto
	Kind                  ResourceKindDto `json:"kind,omitempty"`
	DatabaseIntegrationId string          `json:"databaseIntegrationId,omitempty"`
	SchemaId              string          `json:"schemaId,omitempty"`
	ExternalId            string          `json:"externalId,omitempty"`
}

// DatabaseRecordsBatchDto DTO.
type DatabaseRecordsBatchDto struct {
	Ids       []string         `json:"ids,omitempty"`
	Count     float64          `json:"count,omitempty"`
	Truncated bool             `json:"truncated"`
	Documents []map[string]any `json:"documents,omitempty"`
}

// DatabaseRecordChangeDto DTO.
type DatabaseRecordChangeDto struct {
	Collection    string                   `json:"collection,omitempty"`
	SchemaId      string                   `json:"schemaId,omitempty"`
	SchemaVersion *float64                 `json:"schemaVersion,omitempty"`
	RecordId      string                   `json:"recordId,omitempty"`
	Old           map[string]any           `json:"old,omitempty"`
	New           map[string]any           `json:"new,omitempty"`
	Truncated     bool                     `json:"truncated"`
	OldRef        *DocumentRefDto          `json:"oldRef,omitempty"`
	NewRef        *DocumentRefDto          `json:"newRef,omitempty"`
	Records       *DatabaseRecordsBatchDto `json:"records,omitempty"`
}

// UserChangeDto DTO.
type UserChangeDto struct {
	UserId string         `json:"userId,omitempty"`
	Old    map[string]any `json:"old,omitempty"`
	New    map[string]any `json:"new,omitempty"`
}

// FileChangeDto DTO.
type FileChangeDto struct {
	Old  map[string]any `json:"old,omitempty"`
	New  map[string]any `json:"new,omitempty"`
	Path string         `json:"path,omitempty"`
}

// PaymentFactDto DTO.
type PaymentFactDto struct {
	IntegrationId   string            `json:"integrationId,omitempty"`
	Provider        string            `json:"provider,omitempty"`
	ProviderEventId string            `json:"providerEventId,omitempty"`
	OrderId         string            `json:"orderId,omitempty"`
	EventName       string            `json:"eventName,omitempty"`
	Status          string            `json:"status,omitempty"`
	AmountMinor     *float64          `json:"amountMinor,omitempty"`
	Currency        string            `json:"currency,omitempty"`
	ExternalIds     map[string]string `json:"externalIds,omitempty"`
	OccurredOn      string            `json:"occurredOn,omitempty"`
	ReceivedOn      string            `json:"receivedOn,omitempty"`
}

// AiUsageEventDto DTO.
type AiUsageEventDto struct {
	Payload      map[string]any `json:"payload,omitempty"`
	TargetUserId string         `json:"targetUserId,omitempty"`
}

// ScheduledTaskInfoDto DTO.
type ScheduledTaskInfoDto struct {
	TaskId  string `json:"taskId,omitempty"`
	Name    string `json:"name,omitempty"`
	FiredAt string `json:"firedAt,omitempty"`
}

// MarketplaceFunctionCallDto DTO.
type MarketplaceFunctionCallDto struct {
	FunctionId    string         `json:"functionId,omitempty"`
	FunctionKey   string         `json:"functionKey,omitempty"`
	IntegrationId string         `json:"integrationId,omitempty"`
	Payload       map[string]any `json:"payload,omitempty"`
}

// WorkflowStepInfoDto DTO.
type WorkflowStepInfoDto struct {
	WorkflowId string         `json:"workflowId,omitempty"`
	RunId      string         `json:"runId,omitempty"`
	StepId     string         `json:"stepId,omitempty"`
	Version    *float64       `json:"version,omitempty"`
	Request    map[string]any `json:"request,omitempty"`
}

// CodeFunctionDataDto DTO.
type CodeFunctionDataDto struct {
	Database    *DatabaseRecordChangeDto    `json:"database,omitempty"`
	User        *UserChangeDto              `json:"user,omitempty"`
	File        *FileChangeDto              `json:"file,omitempty"`
	Payment     *PaymentFactDto             `json:"payment,omitempty"`
	Order       *PaymentOrderEventDto       `json:"order,omitempty"`
	Ai          *AiUsageEventDto            `json:"ai,omitempty"`
	Task        *ScheduledTaskInfoDto       `json:"task,omitempty"`
	Payload     map[string]any              `json:"payload,omitempty"`
	Marketplace *MarketplaceFunctionCallDto `json:"marketplace,omitempty"`
	Workflow    *WorkflowStepInfoDto        `json:"workflow,omitempty"`
}

// CodeFunctionEventDto DTO.
type CodeFunctionEventDto struct {
	SpecVersion string                  `json:"specVersion,omitempty"`
	Id          string                  `json:"id,omitempty"`
	EventId     string                  `json:"eventId,omitempty"`
	Event       string                  `json:"event,omitempty"`
	Source      CodeFunctionSource      `json:"source,omitempty"`
	CreatedOn   string                  `json:"createdOn,omitempty"`
	Context     *CodeFunctionContextDto `json:"context,omitempty"`
	Data        *CodeFunctionDataDto    `json:"data,omitempty"`
	Config      map[string]any          `json:"config,omitempty"`
}

// FakeCodeIntegrationDto DTO.
type FakeCodeIntegrationDto struct {
	CodeIntegrationDto
}

// FakeCodeIntegrationRequest DTO.
type FakeCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider CodeProvider `json:"provider,omitempty"`
}

// CodeFunctionSchedulerTaskRequest DTO.
type CodeFunctionSchedulerTaskRequest struct {
	SchedulerTaskRequest
	Type           SchedulerTaskType  `json:"type,omitempty"`
	FunctionRef    *FunctionRefDto    `json:"functionRef,omitempty"`
	InvocationType CodeInvocationType `json:"invocationType,omitempty"`
	Config         string             `json:"config,omitempty"`
	MappedTokens   []*TokenMappingDto `json:"mappedTokens,omitempty"`
}

// MarketplaceFunctionSchedulerTaskRequest DTO.
type MarketplaceFunctionSchedulerTaskRequest struct {
	SchedulerTaskRequest
	Type       SchedulerTaskType `json:"type,omitempty"`
	FunctionId string            `json:"functionId,omitempty"`
	Payload    map[string]string `json:"payload,omitempty"`
}

// WorkflowSchedulerTaskRequest DTO.
type WorkflowSchedulerTaskRequest struct {
	SchedulerTaskRequest
	Type        SchedulerTaskType `json:"type,omitempty"`
	WorkflowRef string            `json:"workflowRef,omitempty"`
	Payload     map[string]string `json:"payload,omitempty"`
}

// AwsCloudWatchLoggingIntegrationRequest DTO.
type AwsCloudWatchLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider  LoggingProvider `json:"provider,omitempty"`
	Region    string          `json:"region,omitempty"`
	LogGroup  string          `json:"logGroup,omitempty"`
	LogStream string          `json:"logStream,omitempty"`
	AccessKey string          `json:"accessKey,omitempty"`
	SecretKey string          `json:"secretKey,omitempty"`
}

// AwsCloudWatchLoggingIntegrationDto DTO.
type AwsCloudWatchLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Region    string `json:"region,omitempty"`
	LogGroup  string `json:"logGroup,omitempty"`
	LogStream string `json:"logStream,omitempty"`
}

// QuickConnectProviderDto DTO.
type QuickConnectProviderDto struct {
	Module              string   `json:"module,omitempty"`
	Provider            string   `json:"provider,omitempty"`
	Available           bool     `json:"available"`
	CallbackKind        string   `json:"callbackKind,omitempty"`
	RequiredInputs      []string `json:"requiredInputs,omitempty"`
	Modes               []string `json:"modes,omitempty"`
	RedirectUrl         string   `json:"redirectUrl,omitempty"`
	PlatformWebhookUrls []string `json:"platformWebhookUrls,omitempty"`
}

// QuickConnectTargetDto DTO.
type QuickConnectTargetDto struct {
	Id     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// JsonSchemaFieldDto DTO.
type JsonSchemaFieldDto struct {
	FieldName string `json:"fieldName,omitempty"`
}

// StringFieldDto DTO.
type StringFieldDto struct {
	JsonSchemaFieldDto
	Format           string            `json:"format,omitempty"`
	Pattern          string            `json:"pattern,omitempty"`
	MinLength        *float64          `json:"minLength,omitempty"`
	MaxLength        *float64          `json:"maxLength,omitempty"`
	TranslateOptions map[string]string `json:"translateOptions,omitempty"`
	Default          string            `json:"default,omitempty"`
	Unique           *bool             `json:"unique,omitempty"`
}

// DecimalFieldDto DTO.
type DecimalFieldDto struct {
	JsonSchemaFieldDto
	Minimum    *float64 `json:"minimum,omitempty"`
	Maximum    *float64 `json:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	Default    *float64 `json:"default,omitempty"`
	Unique     *bool    `json:"unique,omitempty"`
}

// CurrencyDefaultDto DTO.
type CurrencyDefaultDto struct {
	Value    float64 `json:"value,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

// CurrencyFieldDto DTO.
type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	AllowedCurrencies []string            `json:"allowedCurrencies,omitempty"`
	MultipleOf        *float64            `json:"multipleOf,omitempty"`
	Minimum           *float64            `json:"minimum,omitempty"`
	Maximum           *float64            `json:"maximum,omitempty"`
	Default           *CurrencyDefaultDto `json:"default,omitempty"`
}

// BooleanFieldDto DTO.
type BooleanFieldDto struct {
	JsonSchemaFieldDto
	Default *bool `json:"default,omitempty"`
}

// DateFieldDto DTO.
type DateFieldDto struct {
	JsonSchemaFieldDto
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	Default *float64 `json:"default,omitempty"`
}

// IntegerFieldDto DTO.
type IntegerFieldDto struct {
	JsonSchemaFieldDto
	Minimum *float64 `json:"minimum,omitempty"`
	Maximum *float64 `json:"maximum,omitempty"`
	Default *float64 `json:"default,omitempty"`
	Unique  *bool    `json:"unique,omitempty"`
}

// GeolocationFieldDto DTO.
type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	AllowedTypes []string `json:"allowedTypes,omitempty"`
}

// TagsFieldDto DTO.
type TagsFieldDto struct {
	JsonSchemaFieldDto
	MinItems *float64 `json:"minItems,omitempty"`
	MaxItems *float64 `json:"maxItems,omitempty"`
	Default  []string `json:"default,omitempty"`
}

// FileFieldDto DTO.
type FileFieldDto struct {
	JsonSchemaFieldDto
	Storages        []string `json:"storages,omitempty"`
	MinItems        *float64 `json:"minItems,omitempty"`
	MaxItems        *float64 `json:"maxItems,omitempty"`
	AllowedFileType string   `json:"allowedFileType,omitempty"`
	MaxSizeMb       *float64 `json:"maxSizeMb,omitempty"`
}

// TaxonomySelectionFieldDto DTO.
type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	TaxonomyId   string `json:"taxonomyId,omitempty"`
	Multiple     bool   `json:"multiple"`
	DisplayField string `json:"displayField,omitempty"`
}

// CollectionSelectionFieldDto DTO.
type CollectionSelectionFieldDto struct {
	JsonSchemaFieldDto
	CollectionId string `json:"collectionId,omitempty"`
	DisplayField string `json:"displayField,omitempty"`
	Multiple     bool   `json:"multiple"`
}

// UserSelectionFieldDto DTO.
type UserSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple     bool   `json:"multiple"`
	DisplayField string `json:"displayField,omitempty"`
}

// RoleSelectionFieldDto DTO.
type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	Multiple     bool   `json:"multiple"`
	DisplayField string `json:"displayField,omitempty"`
}

// EnumSelectionFieldDto DTO.
type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	Values   []string `json:"values,omitempty"`
	Multiple bool     `json:"multiple"`
	Default  []string `json:"default,omitempty"`
}

// ObjectFieldDto DTO.
type ObjectFieldDto struct {
	JsonSchemaFieldDto
	Properties []*JsonSchemaFieldDto `json:"properties,omitempty"`
	Required   []string              `json:"required,omitempty"`
}

// ArrayFieldDto DTO.
type ArrayFieldDto struct {
	JsonSchemaFieldDto
	Items       *JsonSchemaFieldDto `json:"items,omitempty"`
	MinItems    *float64            `json:"minItems,omitempty"`
	MaxItems    *float64            `json:"maxItems,omitempty"`
	UniqueItems *bool               `json:"uniqueItems,omitempty"`
}

// JsonFieldDto DTO.
type JsonFieldDto struct {
	JsonSchemaFieldDto
	MaxBytes *float64 `json:"maxBytes,omitempty"`
}

// IVersionBasedRequest DTO.
type IVersionBasedRequest struct {
}

// IHasCorrelationIdRequest DTO.
type IHasCorrelationIdRequest struct {
}

// IHasProjectId DTO.
type IHasProjectId struct {
}

// IHasEnv DTO.
type IHasEnv struct {
}

// IHasAccountId DTO.
type IHasAccountId struct {
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

// PromotionIntegrationChoiceDto DTO.
type PromotionIntegrationChoiceDto struct {
	Id      string `json:"id,omitempty"`
	Action  string `json:"action,omitempty"`
	MapToId string `json:"mapToId,omitempty"`
}

// AiPlanModelDto DTO.
type AiPlanModelDto struct {
	LlmIntegrationId string `json:"llmIntegrationId,omitempty"`
	Model            string `json:"model,omitempty"`
}

// AiPlanFeaturesDto DTO.
type AiPlanFeaturesDto struct {
	Attachments bool `json:"attachments"`
	Rag         bool `json:"rag"`
	Memory      bool `json:"memory"`
}

// AiPlanDto DTO.
type AiPlanDto struct {
	Id                  string             `json:"id,omitempty"`
	Name                string             `json:"name,omitempty"`
	AllowedAssistantIds []string           `json:"allowedAssistantIds,omitempty"`
	AllowedModels       []*AiPlanModelDto  `json:"allowedModels,omitempty"`
	QuotaUnit           AiPlanQuotaUnit    `json:"quotaUnit,omitempty"`
	MonthlyQuota        float64            `json:"monthlyQuota,omitempty"`
	Features            *AiPlanFeaturesDto `json:"features,omitempty"`
	QuotaReachedMessage string             `json:"quotaReachedMessage,omitempty"`
}

// AiPlanRoleAssignmentDto DTO.
type AiPlanRoleAssignmentDto struct {
	RoleId string `json:"roleId,omitempty"`
	PlanId string `json:"planId,omitempty"`
}

// ProjectAiAssistantRequestBase DTO.
type ProjectAiAssistantRequestBase struct {
	CodeMashRequestBase
	Name               string   `json:"name,omitempty"`
	WelcomeMessage     string   `json:"welcomeMessage,omitempty"`
	SystemPrompt       string   `json:"systemPrompt,omitempty"`
	Toolsets           []string `json:"toolsets,omitempty"`
	LlmIntegrationId   string   `json:"llmIntegrationId,omitempty"`
	Model              string   `json:"model,omitempty"`
	MemoryEnabled      bool     `json:"memoryEnabled"`
	RagSourceIds       []string `json:"ragSourceIds,omitempty"`
	WeakMatchThreshold *float64 `json:"weakMatchThreshold,omitempty"`
	PlanId             string   `json:"planId,omitempty"`
	IsDefault          bool     `json:"isDefault"`
}

// TeamMemberRoleInput DTO.
type TeamMemberRoleInput struct {
	RoleId       string   `json:"roleId,omitempty"`
	Environments []string `json:"environments,omitempty"`
}

// AiProjectRoleIdsDto DTO.
type AiProjectRoleIdsDto struct {
	ProjectId string   `json:"projectId,omitempty"`
	RoleIds   []string `json:"roleIds,omitempty"`
}

// RoleEnvironmentsDto DTO.
type RoleEnvironmentsDto struct {
	RoleId       string   `json:"roleId,omitempty"`
	Environments []string `json:"environments,omitempty"`
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
	BirthDate                 *float64               `json:"birthDate,omitempty"`
	TimeZone                  string                 `json:"timeZone,omitempty"`
	Language                  string                 `json:"language,omitempty"`
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
	ExtraMetadata             string                 `json:"extraMetadata,omitempty"`
	Notes                     string                 `json:"notes,omitempty"`
}

// ImportUserRowDto DTO.
type ImportUserRowDto struct {
	Email           string              `json:"email,omitempty"`
	UserName        string              `json:"userName,omitempty"`
	UserGeneralInfo *UserGeneralInfoDto `json:"userGeneralInfo,omitempty"`
	Roles           []string            `json:"roles,omitempty"`
	CreatedOn       *float64            `json:"createdOn,omitempty"`
	HashFormat      string              `json:"hashFormat,omitempty"`
	PasswordHash    string              `json:"passwordHash,omitempty"`
	PasswordSalt    string              `json:"passwordSalt,omitempty"`
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

// CursorArgs DTO.
type CursorArgs struct {
	Field string  `json:"field,omitempty"`
	Order float64 `json:"order,omitempty"`
}

// PagingArgs DTO.
type PagingArgs struct {
	CursorArgs    *CursorArgs `json:"cursorArgs,omitempty"`
	PageSize      *float64    `json:"pageSize,omitempty"`
	StartingAfter string      `json:"startingAfter,omitempty"`
	EndingBefore  string      `json:"endingBefore,omitempty"`
}

// SchemaUniqueIndexDto DTO.
type SchemaUniqueIndexDto struct {
	Fields []string `json:"fields,omitempty"`
}

// SchemaSettingsDto DTO.
type SchemaSettingsDto struct {
	SoftDelete        bool                    `json:"softDelete"`
	HasRecordOwner    bool                    `json:"hasRecordOwner"`
	Description       string                  `json:"description,omitempty"`
	UniqueIndexes     []*SchemaUniqueIndexDto `json:"uniqueIndexes,omitempty"`
	EndUserOwnRecords *bool                   `json:"endUserOwnRecords,omitempty"`
}

// SchemaEmbedSettingsDto DTO.
type SchemaEmbedSettingsDto struct {
	Enabled                bool     `json:"enabled"`
	Fields                 []string `json:"fields,omitempty"`
	EmbeddingIntegrationId string   `json:"embeddingIntegrationId,omitempty"`
	PerUser                bool     `json:"perUser"`
}

// SchemaListColumnDto DTO.
type SchemaListColumnDto struct {
	Field         string `json:"field,omitempty"`
	Primary       bool   `json:"primary"`
	LinkReference bool   `json:"linkReference"`
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
	Filters     []string               `json:"filters,omitempty"`
}

// ImportColumnMappingDto DTO.
type ImportColumnMappingDto struct {
	CsvColumnIndex    float64 `json:"csvColumnIndex,omitempty"`
	CsvHeader         string  `json:"csvHeader,omitempty"`
	PropertyName      string  `json:"propertyName,omitempty"`
	DontImportOnError bool    `json:"dontImportOnError"`
}

// EmailValidationIntegrationRequest DTO.
type EmailValidationIntegrationRequest struct {
	IntegrationId   string                  `json:"integrationId,omitempty"`
	Provider        EmailValidationProvider `json:"provider,omitempty"`
	IntegrationName string                  `json:"integrationName,omitempty"`
	IsEnabled       *bool                   `json:"isEnabled,omitempty"`
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

// SaveSmsTemplate DTO.
type SaveSmsTemplate struct {
	CodeMashRequestBase
	TemplateName         string                      `json:"templateName,omitempty"`
	Description          string                      `json:"description,omitempty"`
	CommunicationChannel CommunicationChannel        `json:"communicationChannel,omitempty"`
	Tags                 []string                    `json:"tags,omitempty"`
	Translations         []*SmsMessageTranslationDto `json:"translations,omitempty"`
}

// WorkflowStepRetryDto DTO.
type WorkflowStepRetryDto struct {
	Count          float64 `json:"count,omitempty"`
	BackoffSeconds float64 `json:"backoffSeconds,omitempty"`
}

// WorkflowOnErrorDto DTO.
type WorkflowOnErrorDto struct {
	Action string `json:"action,omitempty"`
	Goto   string `json:"goto,omitempty"`
}

// WorkflowStepDto DTO.
type WorkflowStepDto struct {
	Id             string                `json:"id,omitempty"`
	Kind           string                `json:"kind,omitempty"`
	Name           string                `json:"name,omitempty"`
	Next           string                `json:"next,omitempty"`
	Then           string                `json:"then,omitempty"`
	Else           string                `json:"else,omitempty"`
	Rejected       string                `json:"rejected,omitempty"`
	OnTimeout      string                `json:"onTimeout,omitempty"`
	Retry          *WorkflowStepRetryDto `json:"retry,omitempty"`
	TimeoutSeconds *float64              `json:"timeoutSeconds,omitempty"`
	OnError        *WorkflowOnErrorDto   `json:"onError,omitempty"`
	Action         map[string]any        `json:"action,omitempty"`
	Config         map[string]any        `json:"config,omitempty"`
}

// WorkflowLimitsDto DTO.
type WorkflowLimitsDto struct {
	MaxStepRuns         *float64 `json:"maxStepRuns,omitempty"`
	MaxDurationSeconds  *float64 `json:"maxDurationSeconds,omitempty"`
	MaxWaitSeconds      *float64 `json:"maxWaitSeconds,omitempty"`
	MaxParallelBranches *float64 `json:"maxParallelBranches,omitempty"`
	MaxDepth            *float64 `json:"maxDepth,omitempty"`
}

// WorkflowDefinitionDto DTO.
type WorkflowDefinitionDto struct {
	SchemaVersion       string              `json:"schemaVersion,omitempty"`
	WorkflowId          string              `json:"workflowId,omitempty"`
	Name                string              `json:"name,omitempty"`
	Description         string              `json:"description,omitempty"`
	Version             *float64            `json:"version,omitempty"`
	IsEnabled           bool                `json:"isEnabled"`
	ExecutionIdentityId string              `json:"executionIdentityId,omitempty"`
	Inputs              map[string]any      `json:"inputs,omitempty"`
	Start               string              `json:"start,omitempty"`
	Steps               []*WorkflowStepDto  `json:"steps,omitempty"`
	OnError             *WorkflowOnErrorDto `json:"onError,omitempty"`
	Limits              *WorkflowLimitsDto  `json:"limits,omitempty"`
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
	PlatformApiLevel *float64   `json:"platformApiLevel,omitempty"`
}

// ChatScreenContextDto DTO.
type ChatScreenContextDto struct {
	Kind   string `json:"kind,omitempty"`
	ViewId string `json:"viewId,omitempty"`
}

// INoSessionCookies DTO.
type INoSessionCookies struct {
}

// EmbeddingIntegrationRequest DTO.
type EmbeddingIntegrationRequest struct {
	IntegrationId   string            `json:"integrationId,omitempty"`
	Provider        EmbeddingProvider `json:"provider,omitempty"`
	IntegrationName string            `json:"integrationName,omitempty"`
	IsEnabled       *bool             `json:"isEnabled,omitempty"`
	Endpoint        string            `json:"endpoint,omitempty"`
	Model           string            `json:"model,omitempty"`
	Dimension       float64           `json:"dimension,omitempty"`
	ApiKey          string            `json:"apiKey,omitempty"`
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

// EchoRegionDto DTO.
type EchoRegionDto struct {
	Code        string `json:"code,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	ApiUrl      string `json:"apiUrl,omitempty"`
	HubUrl      string `json:"hubUrl,omitempty"`
}

// EchoAgentSnippetDto DTO.
type EchoAgentSnippetDto struct {
	Client      string `json:"client,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Recommended bool   `json:"recommended"`
	Auth        string `json:"auth,omitempty"`
	Config      string `json:"config,omitempty"`
}

// EchoAgentDto DTO.
type EchoAgentDto struct {
	McpUrl                 string                 `json:"mcpUrl,omitempty"`
	OAuthMetadataUrl       string                 `json:"oAuthMetadataUrl,omitempty"`
	InstallationType       string                 `json:"installationType,omitempty"`
	OnboardingDocsUrl      string                 `json:"onboardingDocsUrl,omitempty"`
	ToolsUrl               string                 `json:"toolsUrl,omitempty"`
	CliPackage             string                 `json:"cliPackage,omitempty"`
	CliInstallCommand      string                 `json:"cliInstallCommand,omitempty"`
	DeviceAuthorizationUrl string                 `json:"deviceAuthorizationUrl,omitempty"`
	DeviceTokenUrl         string                 `json:"deviceTokenUrl,omitempty"`
	Snippets               []*EchoAgentSnippetDto `json:"snippets,omitempty"`
}

// EchoHostingPlatformDto DTO.
type EchoHostingPlatformDto struct {
	Provider string `json:"provider,omitempty"`
	IsAws    bool   `json:"isAws"`
}

// EchoConnectDto DTO.
type EchoConnectDto struct {
	Providers []*QuickConnectProviderDto `json:"providers,omitempty"`
}

// EchoKeylessDto DTO.
type EchoKeylessDto struct {
	Issuer            string `json:"issuer,omitempty"`
	Available         bool   `json:"available"`
	UnavailableReason string `json:"unavailableReason,omitempty"`
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
	MinLength      float64  `json:"minLength,omitempty"`
	MaxLength      *float64 `json:"maxLength,omitempty"`
	MinNumbers     *float64 `json:"minNumbers,omitempty"`
	MinUpper       *float64 `json:"minUpper,omitempty"`
	MinLower       *float64 `json:"minLower,omitempty"`
	MinSpecial     *float64 `json:"minSpecial,omitempty"`
	AllowedSpecial string   `json:"allowedSpecial,omitempty"`
}

// PublicAuthDto DTO.
type PublicAuthDto struct {
	SocialProviders []string                 `json:"socialProviders,omitempty"`
	Passkey         bool                     `json:"passkey"`
	Methods         []string                 `json:"methods,omitempty"`
	PasswordPolicy  *PublicPasswordPolicyDto `json:"passwordPolicy,omitempty"`
}

// PublicAiAssistantDto DTO.
type PublicAiAssistantDto struct {
	Id      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Welcome string `json:"welcome,omitempty"`
}

// PublicAiChatDto DTO.
type PublicAiChatDto struct {
	Enabled    bool                    `json:"enabled"`
	Assistants []*PublicAiAssistantDto `json:"assistants,omitempty"`
}

// TriggerAttentionDto DTO.
type TriggerAttentionDto struct {
	TriggerId   string      `json:"triggerId,omitempty"`
	TriggerType TriggerType `json:"triggerType,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	AtUtc       string      `json:"atUtc,omitempty"`
}

// TriggerActivationTestProblem DTO.
type TriggerActivationTestProblem struct {
	Code    string  `json:"code,omitempty"`
	Message string  `json:"message,omitempty"`
	Line    float64 `json:"line,omitempty"`
	Column  float64 `json:"column,omitempty"`
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
	AccountId          string        `json:"accountId,omitempty"`
	AccountIdAsGuid    string        `json:"accountIdAsGuid,omitempty"`
	UserId             string        `json:"userId,omitempty"`
	LoggedInUserId     string        `json:"loggedInUserId,omitempty"`
	LoggedInUserEmail  string        `json:"loggedInUserEmail,omitempty"`
	Status             AccountStatus `json:"status,omitempty"`
	ProjectCap         float64       `json:"projectCap,omitempty"`
	Permissions        []string      `json:"permissions,omitempty"`
	Roles              []string      `json:"roles,omitempty"`
	AllowedProjects    []string      `json:"allowedProjects,omitempty"`
	TrialWasIssued     bool          `json:"trialWasIssued"`
	EmailVerified      bool          `json:"emailVerified"`
	JoinedByInvitation bool          `json:"joinedByInvitation"`
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

// PromotionSummaryDto DTO.
type PromotionSummaryDto struct {
	Added               float64 `json:"added,omitempty"`
	Changed             float64 `json:"changed,omitempty"`
	Unchanged           float64 `json:"unchanged,omitempty"`
	Deleted             float64 `json:"deleted,omitempty"`
	Kept                float64 `json:"kept,omitempty"`
	Skipped             float64 `json:"skipped,omitempty"`
	RecordsAtRisk       float64 `json:"recordsAtRisk,omitempty"`
	IntegrationsSeeded  float64 `json:"integrationsSeeded,omitempty"`
	IntegrationsSkipped float64 `json:"integrationsSkipped,omitempty"`
	IntegrationsMapped  float64 `json:"integrationsMapped,omitempty"`
	IntegrationsRemoved float64 `json:"integrationsRemoved,omitempty"`
	IsNoOp              bool    `json:"isNoOp"`
}

// PromotionItemChangeDto DTO.
type PromotionItemChangeDto struct {
	Type             string   `json:"type,omitempty"`
	Id               string   `json:"id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Status           string   `json:"status,omitempty"`
	MatchedBy        string   `json:"matchedBy,omitempty"`
	PreviousTargetId string   `json:"previousTargetId,omitempty"`
	BuiltIn          bool     `json:"builtIn"`
	RecordCount      *float64 `json:"recordCount,omitempty"`
	Note             string   `json:"note,omitempty"`
}

// PromotionIntegrationTargetDto DTO.
type PromotionIntegrationTargetDto struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Provider  string `json:"provider,omitempty"`
	IsDefault bool   `json:"isDefault"`
}

// PromotionIntegrationChangeDto DTO.
type PromotionIntegrationChangeDto struct {
	Module        string                           `json:"module,omitempty"`
	Id            string                           `json:"id,omitempty"`
	Name          string                           `json:"name,omitempty"`
	Provider      string                           `json:"provider,omitempty"`
	Status        string                           `json:"status,omitempty"`
	DefaultAction string                           `json:"defaultAction,omitempty"`
	Action        string                           `json:"action,omitempty"`
	TargetId      string                           `json:"targetId,omitempty"`
	MapToId       string                           `json:"mapToId,omitempty"`
	HoldsSecret   bool                             `json:"holdsSecret"`
	IsFake        bool                             `json:"isFake"`
	RealSender    bool                             `json:"realSender"`
	Targets       []*PromotionIntegrationTargetDto `json:"targets,omitempty"`
	Warning       string                           `json:"warning,omitempty"`
	Note          string                           `json:"note,omitempty"`
}

// PromotionNotPromotedDto DTO.
type PromotionNotPromotedDto struct {
	Area   string   `json:"area,omitempty"`
	Scope  string   `json:"scope,omitempty"`
	Reason string   `json:"reason,omitempty"`
	Count  *float64 `json:"count,omitempty"`
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

// PromotionTermChangeDto DTO.
type PromotionTermChangeDto struct {
	TaxonomyId string `json:"taxonomyId,omitempty"`
	TermId     string `json:"termId,omitempty"`
	Slug       string `json:"slug,omitempty"`
	Change     string `json:"change,omitempty"`
}

// PromotionResultDto DTO.
type PromotionResultDto struct {
	Summary                 *PromotionSummaryDto             `json:"summary,omitempty"`
	Items                   []*PromotionItemChangeDto        `json:"items,omitempty"`
	Integrations            []*PromotionIntegrationChangeDto `json:"integrations,omitempty"`
	NotPromoted             []*PromotionNotPromotedDto       `json:"notPromoted,omitempty"`
	RequiresConfirmation    bool                             `json:"requiresConfirmation"`
	ContentMirrored         []*PromotionItemDto              `json:"contentMirrored,omitempty"`
	ContentDeleted          []*PromotionItemDto              `json:"contentDeleted,omitempty"`
	IntegrationsSeeded      []*PromotionItemDto              `json:"integrationsSeeded,omitempty"`
	IntegrationsToProvision []*PromotionItemDto              `json:"integrationsToProvision,omitempty"`
	IntegrationsSkipped     []*PromotionItemDto              `json:"integrationsSkipped,omitempty"`
	Blockers                []*PromotionBlockerDto           `json:"blockers,omitempty"`
	TermChanges             []*PromotionTermChangeDto        `json:"termChanges,omitempty"`
	Warnings                []*PromotionBlockerDto           `json:"warnings,omitempty"`
	FromVersion             *float64                         `json:"fromVersion,omitempty"`
	WasDryRun               bool                             `json:"wasDryRun"`
	Promoted                *bool                            `json:"promoted,omitempty"`
	Message                 string                           `json:"message,omitempty"`
}

// BuiltInRepairDto DTO.
type BuiltInRepairDto struct {
	Env    string                    `json:"env,omitempty"`
	Items  []*PromotionItemChangeDto `json:"items,omitempty"`
	IsNoOp bool                      `json:"isNoOp"`
}

// EnvironmentAdminUrlDto DTO.
type EnvironmentAdminUrlDto struct {
	Env               string `json:"env,omitempty"`
	AdminUrl          string `json:"adminUrl,omitempty"`
	EffectiveAdminUrl string `json:"effectiveAdminUrl,omitempty"`
	Source            string `json:"source,omitempty"`
	Note              string `json:"note,omitempty"`
}

// ProjectEnvironmentsDto DTO.
type ProjectEnvironmentsDto struct {
	Environments   []string                  `json:"environments,omitempty"`
	PromotionOrder []string                  `json:"promotionOrder,omitempty"`
	AdminUrls      []*EnvironmentAdminUrlDto `json:"adminUrls,omitempty"`
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

// ProjectEnvironmentAdminUrlDto DTO.
type ProjectEnvironmentAdminUrlDto struct {
	Env string `json:"env,omitempty"`
	Url string `json:"url,omitempty"`
}

// AuthenticationFlowPasswordPolicyDto DTO.
type AuthenticationFlowPasswordPolicyDto struct {
	MinLength      float64  `json:"minLength,omitempty"`
	MaxLength      *float64 `json:"maxLength,omitempty"`
	MinNumbers     *float64 `json:"minNumbers,omitempty"`
	MinUpper       *float64 `json:"minUpper,omitempty"`
	MinLower       *float64 `json:"minLower,omitempty"`
	MinSpecial     *float64 `json:"minSpecial,omitempty"`
	AllowedSpecial string   `json:"allowedSpecial,omitempty"`
}

// AuthenticationFlowSummaryDto DTO.
type AuthenticationFlowSummaryDto struct {
	Type               string                               `json:"type,omitempty"`
	Provider           string                               `json:"provider,omitempty"`
	PasswordComplexity *AuthenticationFlowPasswordPolicyDto `json:"passwordComplexity,omitempty"`
}

// AiAssistantDto DTO.
type AiAssistantDto struct {
	Id                 string   `json:"id,omitempty"`
	Name               string   `json:"name,omitempty"`
	WelcomeMessage     string   `json:"welcomeMessage,omitempty"`
	SystemPrompt       string   `json:"systemPrompt,omitempty"`
	Toolsets           []string `json:"toolsets,omitempty"`
	LlmIntegrationId   string   `json:"llmIntegrationId,omitempty"`
	Model              string   `json:"model,omitempty"`
	MemoryEnabled      bool     `json:"memoryEnabled"`
	RagSourceIds       []string `json:"ragSourceIds,omitempty"`
	WeakMatchThreshold *float64 `json:"weakMatchThreshold,omitempty"`
	PlanId             string   `json:"planId,omitempty"`
	IsDefault          bool     `json:"isDefault"`
}

// ProjectAiSettingsDto DTO.
type ProjectAiSettingsDto struct {
	Enabled                 bool                       `json:"enabled"`
	DefaultLlmIntegrationId string                     `json:"defaultLlmIntegrationId,omitempty"`
	DefaultModel            string                     `json:"defaultModel,omitempty"`
	Assistants              []*AiAssistantDto          `json:"assistants,omitempty"`
	EmbedFiles              bool                       `json:"embedFiles"`
	Plans                   []*AiPlanDto               `json:"plans,omitempty"`
	PlanRoleAssignments     []*AiPlanRoleAssignmentDto `json:"planRoleAssignments,omitempty"`
	DefaultPlanId           string                     `json:"defaultPlanId,omitempty"`
}

// TriggerDto DTO.
type TriggerDto struct {
	Type           TriggerType       `json:"type,omitempty"`
	ViewId         string            `json:"viewId,omitempty"`
	Name           string            `json:"name,omitempty"`
	ThenAction     *TriggerActionDto `json:"thenAction,omitempty"`
	Description    string            `json:"description,omitempty"`
	IsEnabled      bool              `json:"isEnabled"`
	ActivationCode string            `json:"activationCode,omitempty"`
	Order          *float64          `json:"order,omitempty"`
	BreakOnError   bool              `json:"breakOnError"`
	SavedByAuthId  string            `json:"savedByAuthId,omitempty"`
}

// SchemaTriggerDto DTO.
type SchemaTriggerDto struct {
	TriggerDto
	SchemaId          string            `json:"schemaId,omitempty"`
	When              SchemaTriggerType `json:"when,omitempty"`
	ConfigurationCode string            `json:"configurationCode,omitempty"`
	Env               string            `json:"env,omitempty"`
}

// DatabaseDto DTO.
type DatabaseDto struct {
	IsEnabled                 bool                `json:"isEnabled"`
	Triggers                  []*SchemaTriggerDto `json:"triggers,omitempty"`
	DefaultIntegrationViewIds map[string]string   `json:"defaultIntegrationViewIds,omitempty"`
}

// EmailDto DTO.
type EmailDto struct {
	IsEnabled                 bool              `json:"isEnabled"`
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds,omitempty"`
}

// AiDto DTO.
type AiDto struct {
	IsEnabled                 bool              `json:"isEnabled"`
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds,omitempty"`
	DefaultIntegrationViewId  string            `json:"defaultIntegrationViewId,omitempty"`
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
	IsSystem         bool     `json:"isSystem"`
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
	IsSystem    bool             `json:"isSystem"`
	Version     string           `json:"version,omitempty"`
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
	IsEnabled              bool                    `json:"isEnabled"`
	Triggers               []*MembershipTriggerDto `json:"triggers,omitempty"`
	CustomRoles            []*RoleItemDto          `json:"customRoles,omitempty"`
	CustomPolicies         []*PolicyItemDto        `json:"customPolicies,omitempty"`
	SystemRoleAttachments  []*RoleItemDto          `json:"systemRoleAttachments,omitempty"`
	Authorization          *AuthorizationDto       `json:"authorization,omitempty"`
	RequireEmailValidation bool                    `json:"requireEmailValidation"`
}

// LoggingDto DTO.
type LoggingDto struct {
	IsEnabled     bool `json:"isEnabled"`
	IsEstablished bool `json:"isEstablished"`
}

// ServerEventsDto DTO.
type ServerEventsDto struct {
	IsEnabled bool `json:"isEnabled"`
}

// PushDto DTO.
type PushDto struct {
	IsEnabled                 bool                `json:"isEnabled"`
	DefaultIntegrationViewIds map[string]string   `json:"defaultIntegrationViewIds,omitempty"`
	MarketingTags             []*TagDefinitionDto `json:"marketingTags,omitempty"`
	TransactionalTags         []*TagDefinitionDto `json:"transactionalTags,omitempty"`
}

// SchedulerDto DTO.
type SchedulerDto struct {
	IsEnabled bool `json:"isEnabled"`
}

// CodeDto DTO.
type CodeDto struct {
	IsEnabled        bool `json:"isEnabled"`
	WorkflowsEnabled bool `json:"workflowsEnabled"`
}

// FilesTriggerDto DTO.
type FilesTriggerDto struct {
	TriggerDto
	When   FilesTriggerType `json:"when,omitempty"`
	Folder string           `json:"folder,omitempty"`
}

// FilesDto DTO.
type FilesDto struct {
	IsEnabled                 bool               `json:"isEnabled"`
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
	IsEnabled bool                 `json:"isEnabled"`
	Triggers  []*PaymentTriggerDto `json:"triggers,omitempty"`
}

// SmsDto DTO.
type SmsDto struct {
	IsEnabled                 bool              `json:"isEnabled"`
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds,omitempty"`
}

// ProjectDto DTO.
type ProjectDto struct {
	AccountViewId                 string                           `json:"accountViewId,omitempty"`
	ProjectStatus                 ProjectStatus                    `json:"projectStatus,omitempty"`
	IsActive                      bool                             `json:"isActive"`
	ViewId                        string                           `json:"viewId,omitempty"`
	Name                          string                           `json:"name,omitempty"`
	UniqueName                    string                           `json:"uniqueName,omitempty"`
	HostLabel                     string                           `json:"hostLabel,omitempty"`
	ApiHost                       string                           `json:"apiHost,omitempty"`
	Description                   string                           `json:"description,omitempty"`
	MarketingUrl                  string                           `json:"marketingUrl,omitempty"`
	CanonicalAdminUrl             string                           `json:"canonicalAdminUrl,omitempty"`
	AdminUrl                      string                           `json:"adminUrl,omitempty"`
	EffectiveAdminUrl             string                           `json:"effectiveAdminUrl,omitempty"`
	DefaultLanguage               string                           `json:"defaultLanguage,omitempty"`
	Languages                     []string                         `json:"languages,omitempty"`
	PrimaryRegion                 *ProjectRegionDto                `json:"primaryRegion,omitempty"`
	AdditionalRegions             []*ProjectRegionDto              `json:"additionalRegions,omitempty"`
	IsMultiRegionEligible         bool                             `json:"isMultiRegionEligible"`
	Brand                         *ProjectBrandDto                 `json:"brand,omitempty"`
	NotificationSettings          *NotificationSettingsDto         `json:"notificationSettings,omitempty"`
	AllowedOrigins                []string                         `json:"allowedOrigins,omitempty"`
	EnvironmentAdminUrls          []*ProjectEnvironmentAdminUrlDto `json:"environmentAdminUrls,omitempty"`
	ExposeBrandToAdminPortal      bool                             `json:"exposeBrandToAdminPortal"`
	ExposeAuthToAdminPortal       bool                             `json:"exposeAuthToAdminPortal"`
	AdminPortalEnabled            bool                             `json:"adminPortalEnabled"`
	AdminPortalServiceUserId      string                           `json:"adminPortalServiceUserId,omitempty"`
	MembershipAuthenticationFlows []*AuthenticationFlowSummaryDto  `json:"membershipAuthenticationFlows,omitempty"`
	ExposeLegalToAdminPortal      bool                             `json:"exposeLegalToAdminPortal"`
	AiChat                        *ProjectAiSettingsDto            `json:"aiChat,omitempty"`
	LegalTermsMarkdown            string                           `json:"legalTermsMarkdown,omitempty"`
	LegalPrivacyMarkdown          string                           `json:"legalPrivacyMarkdown,omitempty"`
	Environments                  []string                         `json:"environments,omitempty"`
	EnvironmentRanks              map[string]float64               `json:"environmentRanks,omitempty"`
	Database                      *DatabaseDto                     `json:"database,omitempty"`
	Email                         *EmailDto                        `json:"email,omitempty"`
	Ai                            *AiDto                           `json:"ai,omitempty"`
	Membership                    *MembershipDto                   `json:"membership,omitempty"`
	Logging                       *LoggingDto                      `json:"logging,omitempty"`
	ServerEvents                  *ServerEventsDto                 `json:"serverEvents,omitempty"`
	Push                          *PushDto                         `json:"push,omitempty"`
	Scheduler                     *SchedulerDto                    `json:"scheduler,omitempty"`
	Code                          *CodeDto                         `json:"code,omitempty"`
	Files                         *FilesDto                        `json:"files,omitempty"`
	Payments                      *PaymentsDto                     `json:"payments,omitempty"`
	Sms                           *SmsDto                          `json:"sms,omitempty"`
	DatabaseEnabled               bool                             `json:"databaseEnabled"`
	EmailEnabled                  bool                             `json:"emailEnabled"`
	MembershipEnabled             bool                             `json:"membershipEnabled"`
	LoggingEnabled                bool                             `json:"loggingEnabled"`
	ServerEventsEnabled           bool                             `json:"serverEventsEnabled"`
	PushEnabled                   bool                             `json:"pushEnabled"`
	SchedulerEnabled              bool                             `json:"schedulerEnabled"`
	CodeEnabled                   bool                             `json:"codeEnabled"`
	FilesEnabled                  bool                             `json:"filesEnabled"`
	PaymentsEnabled               bool                             `json:"paymentsEnabled"`
	SmsEnabled                    bool                             `json:"smsEnabled"`
	DefaultLlmIntegrationViewId   string                           `json:"defaultLlmIntegrationViewId,omitempty"`
	Connections                   float64                          `json:"connections,omitempty"`
}

// ProjectListItemDto DTO.
type ProjectListItemDto struct {
	ViewId            string              `json:"viewId,omitempty"`
	IsActive          bool                `json:"isActive"`
	ProjectStatus     ProjectStatus       `json:"projectStatus,omitempty"`
	Name              string              `json:"name,omitempty"`
	UniqueName        string              `json:"uniqueName,omitempty"`
	PrimaryRegion     *ProjectRegionDto   `json:"primaryRegion,omitempty"`
	AdditionalRegions []*ProjectRegionDto `json:"additionalRegions,omitempty"`
	ModulesCount      float64             `json:"modulesCount,omitempty"`
	UsersCount        float64             `json:"usersCount,omitempty"`
}

// AiCreditPackCheckoutDto DTO.
type AiCreditPackCheckoutDto struct {
	Url            string  `json:"url,omitempty"`
	SessionId      string  `json:"sessionId,omitempty"`
	Pack           string  `json:"pack,omitempty"`
	Credits        float64 `json:"credits,omitempty"`
	PriceEuroCents float64 `json:"priceEuroCents,omitempty"`
	PurchaseId     string  `json:"purchaseId,omitempty"`
}

// ProjectAiPlansDto DTO.
type ProjectAiPlansDto struct {
	Plans         []*AiPlanDto               `json:"plans,omitempty"`
	Roles         []*AiPlanRoleAssignmentDto `json:"roles,omitempty"`
	DefaultPlanId string                     `json:"defaultPlanId,omitempty"`
}

// AiUserPlanAssignmentDto DTO.
type AiUserPlanAssignmentDto struct {
	UserId string `json:"userId,omitempty"`
	PlanId string `json:"planId,omitempty"`
}

// AiUserPlansDto DTO.
type AiUserPlansDto struct {
	Users []*AiUserPlanAssignmentDto `json:"users,omitempty"`
}

// AiUsageGroupDto DTO.
type AiUsageGroupDto struct {
	Id                    string  `json:"id,omitempty"`
	LlmInputTokens        float64 `json:"llmInputTokens,omitempty"`
	LlmOutputTokens       float64 `json:"llmOutputTokens,omitempty"`
	EmbeddingTokens       float64 `json:"embeddingTokens,omitempty"`
	RerankCalls           float64 `json:"rerankCalls,omitempty"`
	TotalTokens           float64 `json:"totalTokens,omitempty"`
	ChargeableTokens      float64 `json:"chargeableTokens,omitempty"`
	ChargeableRerankCalls float64 `json:"chargeableRerankCalls,omitempty"`
	Credits               float64 `json:"credits,omitempty"`
}

// AiWalletLineDto DTO.
type AiWalletLineDto struct {
	Kind             string   `json:"kind,omitempty"`
	Credits          float64  `json:"credits,omitempty"`
	Reference        string   `json:"reference,omitempty"`
	PurchasedBalance float64  `json:"purchasedBalance,omitempty"`
	IncludedUsed     *float64 `json:"includedUsed,omitempty"`
	PurchasedUsed    *float64 `json:"purchasedUsed,omitempty"`
	AtUtc            string   `json:"atUtc,omitempty"`
}

// ProjectAiWalletDto DTO.
type ProjectAiWalletDto struct {
	Period           string             `json:"period,omitempty"`
	IncludedCredits  float64            `json:"includedCredits,omitempty"`
	PurchasedCredits float64            `json:"purchasedCredits,omitempty"`
	ConsumedCredits  float64            `json:"consumedCredits,omitempty"`
	RemainingCredits float64            `json:"remainingCredits,omitempty"`
	Status           string             `json:"status,omitempty"`
	Lines            []*AiWalletLineDto `json:"lines,omitempty"`
}

// ProjectAiUsageDto DTO.
type ProjectAiUsageDto struct {
	Period     string              `json:"period,omitempty"`
	Totals     *AiUsageGroupDto    `json:"totals,omitempty"`
	Assistants []*AiUsageGroupDto  `json:"assistants,omitempty"`
	TopUsers   []*AiUsageGroupDto  `json:"topUsers,omitempty"`
	Models     []*AiUsageGroupDto  `json:"models,omitempty"`
	Wallet     *ProjectAiWalletDto `json:"wallet,omitempty"`
}

// TemplateLanguageGapDto DTO.
type TemplateLanguageGapDto struct {
	Module           ApplicationModule `json:"module,omitempty"`
	Kind             LanguageGapKind   `json:"kind,omitempty"`
	TemplateId       string            `json:"templateId,omitempty"`
	TemplateName     string            `json:"templateName,omitempty"`
	MissingLanguages []string          `json:"missingLanguages,omitempty"`
}

// PaginatedResponse DTO.
type PaginatedResponse[TViewModelProjection any] struct {
	Items         []TViewModelProjection `json:"items,omitempty"`
	HasMore       bool                   `json:"hasMore"`
	HasPrevious   bool                   `json:"hasPrevious"`
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
	NeedChangePasswordOnNextLogin bool                  `json:"needChangePasswordOnNextLogin"`
	LastAccessInformation         *AccessInformationDto `json:"lastAccessInformation,omitempty"`
}

// AuthDto DTO.
type AuthDto struct {
	Id                string              `json:"id,omitempty"`
	Type              AuthType            `json:"type,omitempty"`
	Email             string              `json:"email,omitempty"`
	UserName          string              `json:"userName,omitempty"`
	Registration      *RegistrationDto    `json:"registration,omitempty"`
	Login             *LoginDto           `json:"login,omitempty"`
	GeneralInfo       *UserGeneralInfoDto `json:"generalInfo,omitempty"`
	Roles             []string            `json:"roles,omitempty"`
	PushDevices       []string            `json:"pushDevices,omitempty"`
	Tags              []string            `json:"tags,omitempty"`
	Status            AuthStatus          `json:"status,omitempty"`
	CreatedOn         string              `json:"createdOn,omitempty"`
	ModifiedOn        string              `json:"modifiedOn,omitempty"`
	CredentialVersion string              `json:"credentialVersion,omitempty"`
}

// AccountPasswordPolicyDto DTO.
type AccountPasswordPolicyDto struct {
	MinLength      float64  `json:"minLength,omitempty"`
	MaxLength      *float64 `json:"maxLength,omitempty"`
	MinNumbers     *float64 `json:"minNumbers,omitempty"`
	MaxNumbers     *float64 `json:"maxNumbers,omitempty"`
	MinUpper       *float64 `json:"minUpper,omitempty"`
	MaxUpper       *float64 `json:"maxUpper,omitempty"`
	MinLower       *float64 `json:"minLower,omitempty"`
	MaxLower       *float64 `json:"maxLower,omitempty"`
	MinSpecial     *float64 `json:"minSpecial,omitempty"`
	MaxSpecial     *float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial string   `json:"allowedSpecial,omitempty"`
}

// AccountTeamRoleDto DTO.
type AccountTeamRoleDto struct {
	Id          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	IsSystem    bool     `json:"isSystem"`
	Policies    []string `json:"policies,omitempty"`
}

// TeamInvitationDto DTO.
type TeamInvitationDto struct {
	Id        string   `json:"id,omitempty"`
	Email     string   `json:"email,omitempty"`
	Roles     []string `json:"roles,omitempty"`
	SentOn    string   `json:"sentOn,omitempty"`
	ExpiresOn string   `json:"expiresOn,omitempty"`
}

// AccountPasskeyListItemDto DTO.
type AccountPasskeyListItemDto struct {
	CredentialId    string `json:"credentialId,omitempty"`
	FriendlyName    string `json:"friendlyName,omitempty"`
	RegisteredOnUtc string `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   string `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool   `json:"isRevoked"`
}

// LicenseDomainDnsRecordDto DTO.
type LicenseDomainDnsRecordDto struct {
	Host       string `json:"host,omitempty"`
	RecordType string `json:"recordType,omitempty"`
	Resolved   bool   `json:"resolved"`
	Required   bool   `json:"required"`
}

// LicenseDomainDnsStatusDto DTO.
type LicenseDomainDnsStatusDto struct {
	Domain              string                       `json:"domain,omitempty"`
	AllRequiredResolved bool                         `json:"allRequiredResolved"`
	Records             []*LicenseDomainDnsRecordDto `json:"records,omitempty"`
}

// LicenseDomainVerificationChallengeDto DTO.
type LicenseDomainVerificationChallengeDto struct {
	Domain        string `json:"domain,omitempty"`
	TxtHost       string `json:"txtHost,omitempty"`
	TxtValue      string `json:"txtValue,omitempty"`
	ExpiresAtUtc  string `json:"expiresAtUtc,omitempty"`
	Verified      bool   `json:"verified"`
	VerifiedAtUtc string `json:"verifiedAtUtc,omitempty"`
	Skipped       bool   `json:"skipped"`
}

// LicenseDomainVerificationStatusDto DTO.
type LicenseDomainVerificationStatusDto struct {
	Domain           string `json:"domain,omitempty"`
	Verified         bool   `json:"verified"`
	Skipped          bool   `json:"skipped"`
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
	IsTrial           bool   `json:"isTrial"`
	SubscriptionRefId string `json:"subscriptionRefId,omitempty"`
}

// LicenseDto DTO.
type LicenseDto struct {
	CodeMashSubscriptionDto
	IsEnterprise bool    `json:"isEnterprise"`
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

// InstallationLicenseDetailsDto DTO.
type InstallationLicenseDetailsDto struct {
	Domain    string  `json:"domain,omitempty"`
	AccountId string  `json:"accountId,omitempty"`
	Email     string  `json:"email,omitempty"`
	Release   string  `json:"release,omitempty"`
	Expire    float64 `json:"expire,omitempty"`
	IsTrial   bool    `json:"isTrial"`
	Cap       float64 `json:"cap,omitempty"`
}

// InstallationLicenseStatusDto DTO.
type InstallationLicenseStatusDto struct {
	StoredMode                   string                         `json:"storedMode,omitempty"`
	EffectiveMode                string                         `json:"effectiveMode,omitempty"`
	IsProduction                 bool                           `json:"isProduction"`
	GraceDaysLeft                *float64                       `json:"graceDaysLeft,omitempty"`
	GraceUntilUtc                string                         `json:"graceUntilUtc,omitempty"`
	LastProvenAtUtc              string                         `json:"lastProvenAtUtc,omitempty"`
	LastHeartbeatAtUtc           string                         `json:"lastHeartbeatAtUtc,omitempty"`
	InstallationDomain           string                         `json:"installationDomain,omitempty"`
	LicensedDomain               string                         `json:"licensedDomain,omitempty"`
	HostKind                     string                         `json:"hostKind,omitempty"`
	IsTrialLicense               bool                           `json:"isTrialLicense"`
	LicenseExpireUtc             string                         `json:"licenseExpireUtc,omitempty"`
	LastHeartbeatOutcome         string                         `json:"lastHeartbeatOutcome,omitempty"`
	LastHeartbeatError           string                         `json:"lastHeartbeatError,omitempty"`
	Message                      string                         `json:"message,omitempty"`
	License                      *InstallationLicenseDetailsDto `json:"license,omitempty"`
	LicensingDocsUrl             string                         `json:"licensingDocsUrl,omitempty"`
	AskForEnterpriseLicenseEmail string                         `json:"askForEnterpriseLicenseEmail,omitempty"`
}

// AiRoleDto DTO.
type AiRoleDto struct {
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	IsSystem     bool     `json:"isSystem"`
	Environments []string `json:"environments,omitempty"`
}

// AiProjectRolesDto DTO.
type AiProjectRolesDto struct {
	Id    string       `json:"id,omitempty"`
	Name  string       `json:"name,omitempty"`
	Roles []*AiRoleDto `json:"roles,omitempty"`
}

// AiServiceUserKeyDto DTO.
type AiServiceUserKeyDto struct {
	Id       string `json:"id,omitempty"`
	Hint     string `json:"hint,omitempty"`
	IssuedAt string `json:"issuedAt,omitempty"`
}

// AiServiceUserDto DTO.
type AiServiceUserDto struct {
	Id           string                 `json:"id,omitempty"`
	Name         string                 `json:"name,omitempty"`
	AccountRoles []*AiRoleDto           `json:"accountRoles,omitempty"`
	Projects     []*AiProjectRolesDto   `json:"projects,omitempty"`
	CreatedAt    string                 `json:"createdAt,omitempty"`
	Keys         []*AiServiceUserKeyDto `json:"keys,omitempty"`
}

// ImportUserErrorDto DTO.
type ImportUserErrorDto struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ImportUserOutcomeDto DTO.
type ImportUserOutcomeDto struct {
	Row     float64               `json:"row,omitempty"`
	Login   string                `json:"login,omitempty"`
	Outcome string                `json:"outcome,omitempty"`
	UserId  string                `json:"userId,omitempty"`
	Errors  []*ImportUserErrorDto `json:"errors,omitempty"`
}

// UsersImportDto DTO.
type UsersImportDto struct {
	DryRun   bool                    `json:"dryRun"`
	Received float64                 `json:"received,omitempty"`
	Imported float64                 `json:"imported,omitempty"`
	Valid    float64                 `json:"valid,omitempty"`
	Existing float64                 `json:"existing,omitempty"`
	Invalid  float64                 `json:"invalid,omitempty"`
	Failed   float64                 `json:"failed,omitempty"`
	Rows     []*ImportUserOutcomeDto `json:"rows,omitempty"`
}

// LoginProjectRepairDto DTO.
type LoginProjectRepairDto struct {
	DryRun                    bool    `json:"dryRun"`
	LoginsWithoutProject      float64 `json:"loginsWithoutProject,omitempty"`
	LoginsToRepair            float64 `json:"loginsToRepair,omitempty"`
	LoginsRepaired            float64 `json:"loginsRepaired,omitempty"`
	AccountLoginsSkipped      float64 `json:"accountLoginsSkipped,omitempty"`
	OtherProjectLoginsSkipped float64 `json:"otherProjectLoginsSkipped,omitempty"`
	LoginsWithoutEvidence     float64 `json:"loginsWithoutEvidence,omitempty"`
}

// ServiceUserApiKeyDto DTO.
type ServiceUserApiKeyDto struct {
	Id            float64  `json:"id,omitempty"`
	Name          string   `json:"name,omitempty"`
	VisibleKey    string   `json:"visibleKey,omitempty"`
	Scopes        []string `json:"scopes,omitempty"`
	Environment   string   `json:"environment,omitempty"`
	CreatedDate   string   `json:"createdDate,omitempty"`
	ExpiryDate    string   `json:"expiryDate,omitempty"`
	CancelledDate string   `json:"cancelledDate,omitempty"`
	Active        bool     `json:"active"`
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
	HasPreExecuteCode bool              `json:"hasPreExecuteCode"`
	IsEnabled         bool              `json:"isEnabled"`
	Order             *float64          `json:"order,omitempty"`
	BreakOnError      bool              `json:"breakOnError"`
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
	IsSystem    bool    `json:"isSystem"`
	PolicyCount float64 `json:"policyCount,omitempty"`
}

// PasskeySettingsDto DTO.
type PasskeySettingsDto struct {
	Enabled                       bool    `json:"enabled"`
	CodeTtlMinutes                float64 `json:"codeTtlMinutes,omitempty"`
	MaxCredentialsPerUser         float64 `json:"maxCredentialsPerUser,omitempty"`
	RecoveryCodeCount             float64 `json:"recoveryCodeCount,omitempty"`
	GenerateRecoveryCodesAtSignup bool    `json:"generateRecoveryCodesAtSignup"`
	AuthenticatorAttachment       string  `json:"authenticatorAttachment,omitempty"`
	AllowMagicLinkRecovery        bool    `json:"allowMagicLinkRecovery"`
	RefreshTokenTtlDays           float64 `json:"refreshTokenTtlDays,omitempty"`
	RpId                          string  `json:"rpId,omitempty"`
}

// IntegrationListProjection DTO.
type IntegrationListProjection struct {
	ViewId                            string   `json:"viewId,omitempty"`
	IntegrationName                   string   `json:"integrationName,omitempty"`
	IsEnabled                         bool     `json:"isEnabled"`
	LastIntegrationTestAtUtc          string   `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      *bool    `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         []string `json:"lastIntegrationTestErrors,omitempty"`
	HumanDeliveryConfirmedAtUtc       string   `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool     `json:"requiresHumanDeliveryConfirmation"`
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
	SendEmail bool                          `json:"sendEmail"`
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
	MinLength      float64  `json:"minLength,omitempty"`
	MaxLength      *float64 `json:"maxLength,omitempty"`
	MinNumbers     *float64 `json:"minNumbers,omitempty"`
	MaxNumbers     *float64 `json:"maxNumbers,omitempty"`
	MinUpper       *float64 `json:"minUpper,omitempty"`
	MaxUpper       *float64 `json:"maxUpper,omitempty"`
	MinLower       *float64 `json:"minLower,omitempty"`
	MaxLower       *float64 `json:"maxLower,omitempty"`
	MinSpecial     *float64 `json:"minSpecial,omitempty"`
	MaxSpecial     *float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial string   `json:"allowedSpecial,omitempty"`
}

// MembershipAuthorizationViewDto DTO.
type MembershipAuthorizationViewDto struct {
	EmailPreferences             *MembershipEmailPreferencesDto `json:"emailPreferences,omitempty"`
	UserRegistersAsRole          string                         `json:"userRegistersAsRole,omitempty"`
	GuestRegistersAsRole         string                         `json:"guestRegistersAsRole,omitempty"`
	AllowedRegisterRoles         []string                       `json:"allowedRegisterRoles,omitempty"`
	AllowedProviderRegisterRoles []string                       `json:"allowedProviderRegisterRoles,omitempty"`
	ResetPasswordTokenExpiration *float64                       `json:"resetPasswordTokenExpiration,omitempty"`
	InvitationExpiration         *float64                       `json:"invitationExpiration,omitempty"`
	EmailVerificationExpiration  *float64                       `json:"emailVerificationExpiration,omitempty"`
	DeactivationExpiration       *float64                       `json:"deactivationExpiration,omitempty"`
	DefaultSubscribeToNews       bool                           `json:"defaultSubscribeToNews"`
	PasswordComplexity           *PasswordComplexityDto         `json:"passwordComplexity,omitempty"`
	DefaultRoles                 []string                       `json:"defaultRoles,omitempty"`
	AllowedRegistrationRoles     []string                       `json:"allowedRegistrationRoles,omitempty"`
	AllowGuestUsers              bool                           `json:"allowGuestUsers"`
	GuestCleanupPeriodDays       *float64                       `json:"guestCleanupPeriodDays,omitempty"`
	AllowEmailRegistration       bool                           `json:"allowEmailRegistration"`
	AllowUsernameRegistration    bool                           `json:"allowUsernameRegistration"`
	InviteOnly                   bool                           `json:"inviteOnly"`
	AllowAnyoneToSignUp          bool                           `json:"allowAnyoneToSignUp"`
}

// MembershipCredentialsSettingsModeDto DTO.
type MembershipCredentialsSettingsModeDto struct {
	Name      string `json:"name,omitempty"`
	LogoutUrl string `json:"logoutUrl,omitempty"`
}

// MembershipCredentialsSettingsDto DTO.
type MembershipCredentialsSettingsDto struct {
	LogoutUrl      string                                  `json:"logoutUrl,omitempty"`
	AllowUsernames bool                                    `json:"allowUsernames"`
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
	Env  string            `json:"env,omitempty"`
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
	Env                   string           `json:"env,omitempty"`
	TermsMetaDataSchema   *DataSchemaDto   `json:"termsMetaDataSchema,omitempty"`
	TermsMetaVisualSchema *VisualSchemaDto `json:"termsMetaVisualSchema,omitempty"`
	Dependencies          []string         `json:"dependencies,omitempty"`
}

// TaxonomyRef DTO.
type TaxonomyRef struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// TaxonomyListProjection DTO.
type TaxonomyListProjection struct {
	ViewId         string         `json:"viewId,omitempty"`
	TaxonomyName   string         `json:"taxonomyName,omitempty"`
	TaxonomySlug   string         `json:"taxonomySlug,omitempty"`
	ParentId       string         `json:"parentId,omitempty"`
	Description    string         `json:"description,omitempty"`
	Dependencies   []string       `json:"dependencies,omitempty"`
	ParentName     string         `json:"parentName,omitempty"`
	DependencyRefs []*TaxonomyRef `json:"dependencyRefs,omitempty"`
	Env            string         `json:"env,omitempty"`
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
	Order        *float64              `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Slug         string                `json:"slug,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         any                   `json:"meta,omitempty"`
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
	Order        *float64              `json:"order,omitempty"`
	Name         string                `json:"name,omitempty"`
	Names        map[string]string     `json:"names,omitempty"`
	Slug         string                `json:"slug,omitempty"`
	Description  string                `json:"description,omitempty"`
	Descriptions map[string]string     `json:"descriptions,omitempty"`
	MultiParents []*TermMultiParentDto `json:"multiParents,omitempty"`
	Meta         any                   `json:"meta,omitempty"`
	Origin       string                `json:"origin,omitempty"`
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
	Published    bool     `json:"published"`
	LinkedFields []string `json:"linkedFields,omitempty"`
}

// EndUserOwnRecordsRowDto DTO.
type EndUserOwnRecordsRowDto struct {
	SchemaId          string `json:"schemaId,omitempty"`
	SchemaName        string `json:"schemaName,omitempty"`
	HasRecordOwner    bool   `json:"hasRecordOwner"`
	EndUserOwnRecords bool   `json:"endUserOwnRecords"`
	Action            string `json:"action,omitempty"`
}

// EndUserOwnRecordsReportDto DTO.
type EndUserOwnRecordsReportDto struct {
	DryRun  bool                       `json:"dryRun"`
	Env     string                     `json:"env,omitempty"`
	Schemas []*EndUserOwnRecordsRowDto `json:"schemas,omitempty"`
}

// SchemaDto DTO.
type SchemaDto struct {
	ViewId            string                  `json:"viewId,omitempty"`
	SchemaName        string                  `json:"schemaName,omitempty"`
	SchemaSlug        string                  `json:"schemaSlug,omitempty"`
	Version           float64                 `json:"version,omitempty"`
	MetaSchemaVersion float64                 `json:"metaSchemaVersion,omitempty"`
	DataSchema        *DataSchemaDto          `json:"dataSchema,omitempty"`
	VisualSchema      *VisualSchemaDto        `json:"visualSchema,omitempty"`
	PublishedAt       string                  `json:"publishedAt,omitempty"`
	Settings          *SchemaSettingsDto      `json:"settings,omitempty"`
	Embed             *SchemaEmbedSettingsDto `json:"embed,omitempty"`
	Triggers          []*TriggerDto           `json:"triggers,omitempty"`
}

// SchemaListProjection DTO.
type SchemaListProjection struct {
	ViewId            string   `json:"viewId,omitempty"`
	SchemaName        string   `json:"schemaName,omitempty"`
	SchemaTitle       string   `json:"schemaTitle,omitempty"`
	LatestVersion     *float64 `json:"latestVersion,omitempty"`
	HasDraft          bool     `json:"hasDraft"`
	MetaSchemaVersion float64  `json:"metaSchemaVersion,omitempty"`
	Description       string   `json:"description,omitempty"`
	Env               string   `json:"env,omitempty"`
}

// SchemaDraftDto DTO.
type SchemaDraftDto struct {
	DataSchema   *DataSchemaDto   `json:"dataSchema,omitempty"`
	VisualSchema *VisualSchemaDto `json:"visualSchema,omitempty"`
	UpdatedAt    string           `json:"updatedAt,omitempty"`
}

// SchemaIndexDto DTO.
type SchemaIndexDto struct {
	Name   string   `json:"name,omitempty"`
	Keys   []string `json:"keys,omitempty"`
	Unique bool     `json:"unique"`
	Reason string   `json:"reason,omitempty"`
	Field  string   `json:"field,omitempty"`
}

// SchemaIndexIntegrationStatusDto DTO.
type SchemaIndexIntegrationStatusDto struct {
	IntegrationId string            `json:"integrationId,omitempty"`
	State         string            `json:"state,omitempty"`
	Wanted        []*SchemaIndexDto `json:"wanted,omitempty"`
	Created       []string          `json:"created,omitempty"`
	Dropped       []string          `json:"dropped,omitempty"`
	OverCap       []*SchemaIndexDto `json:"overCap,omitempty"`
	RefusedReason string            `json:"refusedReason,omitempty"`
	LastRunUtc    string            `json:"lastRunUtc,omitempty"`
}

// SchemaIndexStatusDto DTO.
type SchemaIndexStatusDto struct {
	SchemaId     string                             `json:"schemaId,omitempty"`
	Env          string                             `json:"env,omitempty"`
	Collection   string                             `json:"collection,omitempty"`
	State        string                             `json:"state,omitempty"`
	Integrations []*SchemaIndexIntegrationStatusDto `json:"integrations,omitempty"`
	LastRunUtc   string                             `json:"lastRunUtc,omitempty"`
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
	Name            string                   `json:"name,omitempty"`
	Keys            []*CollectionIndexKeyDto `json:"keys,omitempty"`
	CreatedByNorbix bool                     `json:"createdByNorbix"`
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
	Env      string           `json:"env,omitempty"`
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
	HasHeader     bool                      `json:"hasHeader"`
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

// KeylessIntegrationSetupTemplate DTO.
type KeylessIntegrationSetupTemplate struct {
	Kind     string `json:"kind,omitempty"`
	FileName string `json:"fileName,omitempty"`
	Content  string `json:"content,omitempty"`
}

// PublicFolderDto DTO.
type PublicFolderDto struct {
	Path      string `json:"path,omitempty"`
	PublicId  string `json:"publicId,omitempty"`
	PublicUrl string `json:"publicUrl,omitempty"`
	Inherited bool   `json:"inherited"`
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
	Address string   `json:"address,omitempty"`
	Verdict string   `json:"verdict,omitempty"`
	Reason  string   `json:"reason,omitempty"`
	Score   *float64 `json:"score,omitempty"`
}

// TemplateListProjection DTO.
type TemplateListProjection struct {
	Id           string               `json:"id,omitempty"`
	ViewId       string               `json:"viewId,omitempty"`
	TemplateName string               `json:"templateName,omitempty"`
	IsActive     bool                 `json:"isActive"`
	Type         CommunicationChannel `json:"type,omitempty"`
	Tags         []string             `json:"tags,omitempty"`
	Env          string               `json:"env,omitempty"`
}

// EmailTemplateListProjection DTO.
type EmailTemplateListProjection struct {
	TemplateListProjection
	HasAttachments bool     `json:"hasAttachments"`
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
	ImagePreview string                   `json:"imagePreview,omitempty"`
	Description  string                   `json:"description,omitempty"`
	Theme        SystemEmailTemplateTheme `json:"theme,omitempty"`
	SystemTags   []string                 `json:"systemTags,omitempty"`
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

// EmailLinkPreferencesDto DTO.
type EmailLinkPreferencesDto struct {
	EmailAddress              string   `json:"emailAddress,omitempty"`
	UnsubscribedFromMarketing bool     `json:"unsubscribedFromMarketing"`
	BlockReasons              []string `json:"blockReasons,omitempty"`
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
	ForceCampaignLanguage           bool                            `json:"forceCampaignLanguage"`
	CampaignProcessingIntegrationId string                          `json:"campaignProcessingIntegrationId,omitempty"`
	StatusHistory                   []*CampaignStatusChangeEntryDto `json:"statusHistory,omitempty"`
	Status                          *CampaignStatusChangeEntryDto   `json:"status,omitempty"`
	TokenMappingValues              []*TokenMappingDto              `json:"tokenMappingValues,omitempty"`
	Notes                           string                          `json:"notes,omitempty"`
	CreatedById                     string                          `json:"createdById,omitempty"`
	TimeZoneId                      string                          `json:"timeZoneId,omitempty"`
	UserId                          string                          `json:"userId,omitempty"`
	Id                              string                          `json:"id,omitempty"`
}

// EmailCampaignDto DTO.
type EmailCampaignDto struct {
	CampaignDto
	DeliverySettings        *EmailCampaignDeliverySettingsDto `json:"deliverySettings,omitempty"`
	Template                *EmailTemplateDto                 `json:"template,omitempty"`
	ValidationIntegrationId string                            `json:"validationIntegrationId,omitempty"`
	TemplateIsSystem        bool                              `json:"templateIsSystem"`
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
	HasMore       bool                 `json:"hasMore"`
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

// NotificationAttemptDto DTO.
type NotificationAttemptDto struct {
	AttemptId         string                     `json:"attemptId,omitempty"`
	Number            float64                    `json:"number,omitempty"`
	StartedAt         string                     `json:"startedAt,omitempty"`
	FinishedAt        string                     `json:"finishedAt,omitempty"`
	Outcome           NotificationAttemptOutcome `json:"outcome,omitempty"`
	Category          SendErrorCategory          `json:"category,omitempty"`
	StatusCode        *float64                   `json:"statusCode,omitempty"`
	ErrorCode         string                     `json:"errorCode,omitempty"`
	Message           string                     `json:"message,omitempty"`
	ProviderRef       string                     `json:"providerRef,omitempty"`
	RetryAfterSeconds *float64                   `json:"retryAfterSeconds,omitempty"`
	Exhausted         *bool                      `json:"exhausted,omitempty"`
	NextRetryAt       string                     `json:"nextRetryAt,omitempty"`
	CorrelationId     string                     `json:"correlationId,omitempty"`
	Trigger           NotificationAttemptTrigger `json:"trigger,omitempty"`
	RetryRequestId    string                     `json:"retryRequestId,omitempty"`
}

// NotificationDeliveryEventDto DTO.
type NotificationDeliveryEventDto struct {
	EventId     string               `json:"eventId,omitempty"`
	Outcome     DeliveryEventOutcome `json:"outcome,omitempty"`
	At          string               `json:"at,omitempty"`
	Provider    string               `json:"provider,omitempty"`
	ProviderRef string               `json:"providerRef,omitempty"`
	Reason      string               `json:"reason,omitempty"`
}

// CampaignBatchNotificationDto DTO.
type CampaignBatchNotificationDto struct {
	CampaignId          string                              `json:"campaignId,omitempty"`
	BatchId             string                              `json:"batchId,omitempty"`
	NotificationId      string                              `json:"notificationId,omitempty"`
	RefNotificationId   string                              `json:"refNotificationId,omitempty"`
	Body                string                              `json:"body,omitempty"`
	Model               map[string]string                   `json:"model,omitempty"`
	StatusHistory       []*NotificationStatusChangeEntryDto `json:"statusHistory,omitempty"`
	State               NotificationDeliveryState           `json:"state,omitempty"`
	StateChangedAt      string                              `json:"stateChangedAt,omitempty"`
	Attempts            []*NotificationAttemptDto           `json:"attempts,omitempty"`
	DeliveryEvents      []*NotificationDeliveryEventDto     `json:"deliveryEvents,omitempty"`
	DeliveryOutcome     DeliveryEventOutcome                `json:"deliveryOutcome,omitempty"`
	DeliveryOutcomeRank *float64                            `json:"deliveryOutcomeRank,omitempty"`
	Id                  string                              `json:"id,omitempty"`
}

// EmailCampaignBatchNotificationDto DTO.
type EmailCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	Subject           string                  `json:"subject,omitempty"`
	Recipients        *EmailRecipientsDto     `json:"recipients,omitempty"`
	Content           *EmailMessageContentDto `json:"content,omitempty"`
	TransactionalKind string                  `json:"transactionalKind,omitempty"`
}

// CampaignStatsDto DTO.
type CampaignStatsDto struct {
	Batches         float64 `json:"batches,omitempty"`
	Sent            float64 `json:"sent,omitempty"`
	Failed          float64 `json:"failed,omitempty"`
	FailedRetryable float64 `json:"failedRetryable,omitempty"`
	Cancelled       float64 `json:"cancelled,omitempty"`
	SuccessRate     float64 `json:"successRate,omitempty"`
}

// SmsTemplateListProjection DTO.
type SmsTemplateListProjection struct {
	TemplateListProjection
	Languages []string `json:"languages,omitempty"`
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
	HasMore       bool               `json:"hasMore"`
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

// WorkflowVersionSummaryDto DTO.
type WorkflowVersionSummaryDto struct {
	Version           float64 `json:"version,omitempty"`
	Name              string  `json:"name,omitempty"`
	StepCount         float64 `json:"stepCount,omitempty"`
	PublishedAt       string  `json:"publishedAt,omitempty"`
	PublishedByAuthId string  `json:"publishedByAuthId,omitempty"`
}

// WorkflowValidationIssueDto DTO.
type WorkflowValidationIssueDto struct {
	Code    string   `json:"code,omitempty"`
	Message string   `json:"message,omitempty"`
	Path    string   `json:"path,omitempty"`
	StepId  string   `json:"stepId,omitempty"`
	Line    *float64 `json:"line,omitempty"`
	Column  *float64 `json:"column,omitempty"`
}

// WorkflowDto DTO.
type WorkflowDto struct {
	WorkflowId     string                        `json:"workflowId,omitempty"`
	Env            string                        `json:"env,omitempty"`
	Definition     *WorkflowDefinitionDto        `json:"definition,omitempty"`
	Version        float64                       `json:"version,omitempty"`
	IsPublished    bool                          `json:"isPublished"`
	HasDraft       bool                          `json:"hasDraft"`
	IsEnabled      bool                          `json:"isEnabled"`
	DisabledReason string                        `json:"disabledReason,omitempty"`
	DraftUpdatedAt string                        `json:"draftUpdatedAt,omitempty"`
	PublishedAt    string                        `json:"publishedAt,omitempty"`
	CreatedAt      string                        `json:"createdAt,omitempty"`
	UpdatedAt      string                        `json:"updatedAt,omitempty"`
	SavedByAuthId  string                        `json:"savedByAuthId,omitempty"`
	Versions       []*WorkflowVersionSummaryDto  `json:"versions,omitempty"`
	Warnings       []*WorkflowValidationIssueDto `json:"warnings,omitempty"`
}

// WorkflowStepDiffDto DTO.
type WorkflowStepDiffDto struct {
	StepId string   `json:"stepId,omitempty"`
	Paths  []string `json:"paths,omitempty"`
}

// WorkflowVersionDiffDto DTO.
type WorkflowVersionDiffDto struct {
	FromVersion   float64                `json:"fromVersion,omitempty"`
	ToVersion     float64                `json:"toVersion,omitempty"`
	AddedSteps    []string               `json:"addedSteps,omitempty"`
	RemovedSteps  []string               `json:"removedSteps,omitempty"`
	ChangedSteps  []*WorkflowStepDiffDto `json:"changedSteps,omitempty"`
	ChangedFields []string               `json:"changedFields,omitempty"`
}

// WorkflowLastRunDto DTO.
type WorkflowLastRunDto struct {
	RunId      string `json:"runId,omitempty"`
	Status     string `json:"status,omitempty"`
	FinishedAt string `json:"finishedAt,omitempty"`
}

// WorkflowListProjection DTO.
type WorkflowListProjection struct {
	ViewId              string              `json:"viewId,omitempty"`
	WorkflowId          string              `json:"workflowId,omitempty"`
	Name                string              `json:"name,omitempty"`
	Env                 string              `json:"env,omitempty"`
	Version             float64             `json:"version,omitempty"`
	IsPublished         bool                `json:"isPublished"`
	HasDraft            bool                `json:"hasDraft"`
	IsEnabled           bool                `json:"isEnabled"`
	DisabledReason      string              `json:"disabledReason,omitempty"`
	ExecutionIdentityId string              `json:"executionIdentityId,omitempty"`
	UpdatedAt           string              `json:"updatedAt,omitempty"`
	LastRun             *WorkflowLastRunDto `json:"lastRun,omitempty"`
	WaitingRuns         float64             `json:"waitingRuns,omitempty"`
}

// WorkflowStepKindSchemaDto DTO.
type WorkflowStepKindSchemaDto struct {
	Kind         string         `json:"kind,omitempty"`
	Description  string         `json:"description,omitempty"`
	ConfigSchema map[string]any `json:"configSchema,omitempty"`
	OutputSchema map[string]any `json:"outputSchema,omitempty"`
	Requires     []string       `json:"requires,omitempty"`
}

// WorkflowBindingFamilyDto DTO.
type WorkflowBindingFamilyDto struct {
	Family      string `json:"family,omitempty"`
	Description string `json:"description,omitempty"`
}

// WorkflowSchemaDto DTO.
type WorkflowSchemaDto struct {
	SchemaVersion string                       `json:"schemaVersion,omitempty"`
	Definition    map[string]any               `json:"definition,omitempty"`
	StepKinds     []*WorkflowStepKindSchemaDto `json:"stepKinds,omitempty"`
	Bindings      []*WorkflowBindingFamilyDto  `json:"bindings,omitempty"`
}

// WorkflowRunErrorDto DTO.
type WorkflowRunErrorDto struct {
	Code    string `json:"code,omitempty"`
	StepId  string `json:"stepId,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

// WorkflowRunActorDto DTO.
type WorkflowRunActorDto struct {
	UserId string `json:"userId,omitempty"`
	Kind   string `json:"kind,omitempty"`
}

// WorkflowWaitingOnDto DTO.
type WorkflowWaitingOnDto struct {
	StepId    string   `json:"stepId,omitempty"`
	Kind      string   `json:"kind,omitempty"`
	Until     string   `json:"until,omitempty"`
	Key       string   `json:"key,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

// WorkflowRunStepDto DTO.
type WorkflowRunStepDto struct {
	StepId          string               `json:"stepId,omitempty"`
	Kind            string               `json:"kind,omitempty"`
	Name            string               `json:"name,omitempty"`
	Status          string               `json:"status,omitempty"`
	Attempt         float64              `json:"attempt,omitempty"`
	Pass            float64              `json:"pass,omitempty"`
	StartedAt       string               `json:"startedAt,omitempty"`
	FinishedAt      string               `json:"finishedAt,omitempty"`
	DurationMs      *float64             `json:"durationMs,omitempty"`
	Input           any                  `json:"input,omitempty"`
	InputTruncated  bool                 `json:"inputTruncated"`
	Output          any                  `json:"output,omitempty"`
	OutputTruncated bool                 `json:"outputTruncated"`
	Error           *WorkflowRunErrorDto `json:"error,omitempty"`
	ResumedBy       *WorkflowRunActorDto `json:"resumedBy,omitempty"`
	Decision        string               `json:"decision,omitempty"`
}

// WorkflowRunDto DTO.
type WorkflowRunDto struct {
	RunId               string                  `json:"runId,omitempty"`
	WorkflowId          string                  `json:"workflowId,omitempty"`
	Version             *float64                `json:"version,omitempty"`
	IsDraft             bool                    `json:"isDraft"`
	Status              string                  `json:"status,omitempty"`
	Source              string                  `json:"source,omitempty"`
	StartedBy           *WorkflowRunActorDto    `json:"startedBy,omitempty"`
	TriggerId           string                  `json:"triggerId,omitempty"`
	TaskId              string                  `json:"taskId,omitempty"`
	ParentRunId         string                  `json:"parentRunId,omitempty"`
	Depth               float64                 `json:"depth,omitempty"`
	Env                 string                  `json:"env,omitempty"`
	CorrelationId       string                  `json:"correlationId,omitempty"`
	ExecutionIdentityId string                  `json:"executionIdentityId,omitempty"`
	StartedAt           string                  `json:"startedAt,omitempty"`
	FinishedAt          string                  `json:"finishedAt,omitempty"`
	DurationMs          *float64                `json:"durationMs,omitempty"`
	WaitingOn           []*WorkflowWaitingOnDto `json:"waitingOn,omitempty"`
	Error               *WorkflowRunErrorDto    `json:"error,omitempty"`
	Output              any                     `json:"output,omitempty"`
	StepRuns            float64                 `json:"stepRuns,omitempty"`
	Steps               []*WorkflowRunStepDto   `json:"steps,omitempty"`
	Request             map[string]any          `json:"request,omitempty"`
	Trigger             map[string]any          `json:"trigger,omitempty"`
}

// WorkflowRunListProjection DTO.
type WorkflowRunListProjection struct {
	RunId          string                  `json:"runId,omitempty"`
	WorkflowId     string                  `json:"workflowId,omitempty"`
	Version        *float64                `json:"version,omitempty"`
	IsDraft        bool                    `json:"isDraft"`
	Status         string                  `json:"status,omitempty"`
	Source         string                  `json:"source,omitempty"`
	StartedAt      string                  `json:"startedAt,omitempty"`
	FinishedAt     string                  `json:"finishedAt,omitempty"`
	DurationMs     *float64                `json:"durationMs,omitempty"`
	CurrentStepIds []string                `json:"currentStepIds,omitempty"`
	WaitingOn      []*WorkflowWaitingOnDto `json:"waitingOn,omitempty"`
	Error          *WorkflowRunErrorDto    `json:"error,omitempty"`
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
	Env           string            `json:"env,omitempty"`
	Origin        string            `json:"origin,omitempty"`
	Meta          map[string]string `json:"meta,omitempty"`
}

// MarketplaceListingProjection DTO.
type MarketplaceListingProjection struct {
	ViewId        string                  `json:"viewId,omitempty"`
	Slug          string                  `json:"slug,omitempty"`
	DisplayName   string                  `json:"displayName,omitempty"`
	Vendor        string                  `json:"vendor,omitempty"`
	Category      MarketplaceCategory     `json:"category,omitempty"`
	Transport     MarketplaceTransport    `json:"transport,omitempty"`
	IconUrl       string                  `json:"iconUrl,omitempty"`
	IsOfficial    bool                    `json:"isOfficial"`
	Tags          []string                `json:"tags,omitempty"`
	FunctionCount float64                 `json:"functionCount,omitempty"`
	Availability  MarketplaceAvailability `json:"availability,omitempty"`
	VideoUrl      string                  `json:"videoUrl,omitempty"`
}

// MarketplaceAccessCustomPolicyDto DTO.
type MarketplaceAccessCustomPolicyDto struct {
	PolicyId    string  `json:"policyId,omitempty"`
	Name        string  `json:"name,omitempty"`
	UsedByRoles float64 `json:"usedByRoles,omitempty"`
}

// MarketplaceAccessRoleDto DTO.
type MarketplaceAccessRoleDto struct {
	RoleId            string                              `json:"roleId,omitempty"`
	Name              string                              `json:"name,omitempty"`
	Scope             string                              `json:"scope,omitempty"`
	IsSystem          bool                                `json:"isSystem"`
	Fixed             bool                                `json:"fixed"`
	Package           MarketplaceAccessLevel              `json:"package,omitempty"`
	IsCustom          bool                                `json:"isCustom"`
	Manage            bool                                `json:"manage"`
	Custom            []*MarketplaceAccessCustomPolicyDto `json:"custom,omitempty"`
	DefaultNotApplied string                              `json:"defaultNotApplied,omitempty"`
	CanChange         bool                                `json:"canChange"`
}

// MarketplaceAccessPackageDto DTO.
type MarketplaceAccessPackageDto struct {
	Kind     MarketplaceAccessLevel `json:"kind,omitempty"`
	PolicyId string                 `json:"policyId,omitempty"`
	Name     string                 `json:"name,omitempty"`
	Exists   bool                   `json:"exists"`
}

// MarketplaceAccessToolDto DTO.
type MarketplaceAccessToolDto struct {
	Name                   string `json:"name,omitempty"`
	ReadOnly               bool   `json:"readOnly"`
	Grantable              bool   `json:"grantable"`
	Curated                bool   `json:"curated"`
	ExcludedFromFullAccess bool   `json:"excludedFromFullAccess"`
}

// MarketplaceAccessFunctionDto DTO.
type MarketplaceAccessFunctionDto struct {
	FunctionId  string `json:"functionId,omitempty"`
	Name        string `json:"name,omitempty"`
	FunctionKey string `json:"functionKey,omitempty"`
}

// MarketplaceIntegrationListProjection DTO.
type MarketplaceIntegrationListProjection struct {
	ViewId                            string               `json:"viewId,omitempty"`
	IntegrationName                   string               `json:"integrationName,omitempty"`
	IsEnabled                         bool                 `json:"isEnabled"`
	ListingViewId                     string               `json:"listingViewId,omitempty"`
	Vendor                            string               `json:"vendor,omitempty"`
	Category                          MarketplaceCategory  `json:"category,omitempty"`
	Transport                         MarketplaceTransport `json:"transport,omitempty"`
	LastIntegrationTestAtUtc          string               `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      *bool                `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         []string             `json:"lastIntegrationTestErrors,omitempty"`
	HumanDeliveryConfirmedAtUtc       string               `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool                 `json:"requiresHumanDeliveryConfirmation"`
	SignedIn                          bool                 `json:"signedIn"`
	SignedInAt                        string               `json:"signedInAt,omitempty"`
}

// MarketplaceMcpToolDto DTO.
type MarketplaceMcpToolDto struct {
	Name            string `json:"name,omitempty"`
	Title           string `json:"title,omitempty"`
	Description     string `json:"description,omitempty"`
	InputSchema     string `json:"inputSchema,omitempty"`
	ReadOnlyHint    *bool  `json:"readOnlyHint,omitempty"`
	DestructiveHint *bool  `json:"destructiveHint,omitempty"`
}

// MarketplaceAccessChangesDto DTO.
type MarketplaceAccessChangesDto struct {
	ReadOnlyAdded          []string `json:"readOnlyAdded,omitempty"`
	ReadOnlyRemoved        []string `json:"readOnlyRemoved,omitempty"`
	NotGrantable           []string `json:"notGrantable,omitempty"`
	ReadOnlyHintOverridden []string `json:"readOnlyHintOverridden,omitempty"`
	ExcludedFromFullAccess []string `json:"excludedFromFullAccess,omitempty"`
}

// MarketplaceFunctionProjection DTO.
type MarketplaceFunctionProjection struct {
	ViewId            string  `json:"viewId,omitempty"`
	IntegrationViewId string  `json:"integrationViewId,omitempty"`
	FunctionKey       string  `json:"functionKey,omitempty"`
	DisplayName       string  `json:"displayName,omitempty"`
	IsEnabled         bool    `json:"isEnabled"`
	MappingCount      float64 `json:"mappingCount,omitempty"`
}

// MarketplaceFunctionRunDto DTO.
type MarketplaceFunctionRunDto struct {
	FiredAt         string   `json:"firedAt,omitempty"`
	Event           string   `json:"event,omitempty"`
	Source          string   `json:"source,omitempty"`
	FunctionId      string   `json:"functionId,omitempty"`
	RunAsUserId     string   `json:"runAsUserId,omitempty"`
	Succeeded       bool     `json:"succeeded"`
	Status          string   `json:"status,omitempty"`
	DurationMs      *float64 `json:"durationMs,omitempty"`
	Error           string   `json:"error,omitempty"`
	VendorRequestId string   `json:"vendorRequestId,omitempty"`
	CorrelationId   string   `json:"correlationId,omitempty"`
	HasDetails      bool     `json:"hasDetails"`
	TriggerId       string   `json:"triggerId,omitempty"`
	TaskId          string   `json:"taskId,omitempty"`
	Input           string   `json:"input,omitempty"`
	IntegrationId   string   `json:"integrationId,omitempty"`
	ToolName        string   `json:"toolName,omitempty"`
}

// CodeIntegrationListProjection DTO.
type CodeIntegrationListProjection struct {
	IntegrationListProjection
	Provider CodeProvider `json:"provider,omitempty"`
}

// CodeFunctionDescriptorDto DTO.
type CodeFunctionDescriptorDto struct {
	FunctionId   string            `json:"functionId,omitempty"`
	DisplayName  string            `json:"displayName,omitempty"`
	Runtime      string            `json:"runtime,omitempty"`
	Region       string            `json:"region,omitempty"`
	Description  string            `json:"description,omitempty"`
	LastModified string            `json:"lastModified,omitempty"`
	Tags         map[string]string `json:"tags,omitempty"`
	Arn          string            `json:"arn,omitempty"`
}

// CodeFunctionQualifierDto DTO.
type CodeFunctionQualifierDto struct {
	Name          string                       `json:"name,omitempty"`
	Kind          CodeFunctionQualifierKindDto `json:"kind,omitempty"`
	Description   string                       `json:"description,omitempty"`
	TargetVersion string                       `json:"targetVersion,omitempty"`
}

// CodeFunctionRunDto DTO.
type CodeFunctionRunDto struct {
	FiredAt    string   `json:"firedAt,omitempty"`
	Event      string   `json:"event,omitempty"`
	Source     string   `json:"source,omitempty"`
	Succeeded  bool     `json:"succeeded"`
	StatusCode *float64 `json:"statusCode,omitempty"`
	DurationMs *float64 `json:"durationMs,omitempty"`
	Error      string   `json:"error,omitempty"`
	RequestId  string   `json:"requestId,omitempty"`
	Input      string   `json:"input,omitempty"`
}

// PushTemplateListProjection DTO.
type PushTemplateListProjection struct {
	TemplateListProjection
	Languages []string `json:"languages,omitempty"`
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

// PushDeviceListProjection DTO.
type PushDeviceListProjection struct {
	Id               string   `json:"id,omitempty"`
	UserId           string   `json:"userId,omitempty"`
	DeviceOs         string   `json:"deviceOs,omitempty"`
	Token            string   `json:"token,omitempty"`
	Brand            string   `json:"brand,omitempty"`
	Manufacturer     string   `json:"manufacturer,omitempty"`
	ModelName        string   `json:"modelName,omitempty"`
	DeviceName       string   `json:"deviceName,omitempty"`
	DeviceType       string   `json:"deviceType,omitempty"`
	OsName           string   `json:"osName,omitempty"`
	OsVersion        string   `json:"osVersion,omitempty"`
	PlatformApiLevel *float64 `json:"platformApiLevel,omitempty"`
	Platform         string   `json:"platform,omitempty"`
	AppVersion       string   `json:"appVersion,omitempty"`
	InstallationId   string   `json:"installationId,omitempty"`
	Locale           string   `json:"locale,omitempty"`
	TimeZone         string   `json:"timeZone,omitempty"`
	UpdatedAtUtc     string   `json:"updatedAtUtc,omitempty"`
}

// TeamPushDeviceListProjection DTO.
type TeamPushDeviceListProjection struct {
	PushDeviceListProjection
	OwnerKind  string `json:"ownerKind,omitempty"`
	OwnerName  string `json:"ownerName,omitempty"`
	OwnerEmail string `json:"ownerEmail,omitempty"`
}

// PushAudienceCountDto DTO.
type PushAudienceCountDto struct {
	Devices        float64 `json:"devices,omitempty"`
	Recipients     float64 `json:"recipients,omitempty"`
	SkippedUserIds float64 `json:"skippedUserIds,omitempty"`
	IsCapped       bool    `json:"isCapped"`
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
	HasMore       bool                `json:"hasMore"`
}

// PushCampaignBatchDto DTO.
type PushCampaignBatchDto struct {
	CampaignBatchDto
	Recipients *PushRecipientsDto `json:"recipients,omitempty"`
}

// NotificationReadMarkDto DTO.
type NotificationReadMarkDto struct {
	UserId    string `json:"userId,omitempty"`
	ReadAtUtc string `json:"readAtUtc,omitempty"`
}

// PushCampaignBatchNotificationDto DTO.
type PushCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	Subject    string                     `json:"subject,omitempty"`
	Recipients *PushRecipientsDto         `json:"recipients,omitempty"`
	Content    *PushMessageContentDto     `json:"content,omitempty"`
	ReadBy     []*NotificationReadMarkDto `json:"readBy,omitempty"`
}

// InAppNotificationDto DTO.
type InAppNotificationDto struct {
	Id           string `json:"id,omitempty"`
	ProjectId    string `json:"projectId,omitempty"`
	UserId       string `json:"userId,omitempty"`
	Channel      string `json:"channel,omitempty"`
	EventName    string `json:"eventName,omitempty"`
	PayloadJson  string `json:"payloadJson,omitempty"`
	IsRead       bool   `json:"isRead"`
	SeenAtUtc    string `json:"seenAtUtc,omitempty"`
	ReadAtUtc    string `json:"readAtUtc,omitempty"`
	CreatedAtUtc string `json:"createdAtUtc,omitempty"`
}

// PaymentsWebhookLogEntry DTO.
type PaymentsWebhookLogEntry struct {
	RunId           string  `json:"runId,omitempty"`
	IntegrationId   string  `json:"integrationId,omitempty"`
	Source          string  `json:"source,omitempty"`
	Env             string  `json:"env,omitempty"`
	EventName       string  `json:"eventName,omitempty"`
	ProviderEventId string  `json:"providerEventId,omitempty"`
	StatusCode      float64 `json:"statusCode,omitempty"`
	Description     string  `json:"description,omitempty"`
	ReceivedOn      string  `json:"receivedOn,omitempty"`
}

// PaymentsWebhookRunOrderDto DTO.
type PaymentsWebhookRunOrderDto struct {
	OrderId string `json:"orderId,omitempty"`
	Kind    string `json:"kind,omitempty"`
	From    string `json:"from,omitempty"`
	To      string `json:"to,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

// PaymentsWebhookRunActionDto DTO.
type PaymentsWebhookRunActionDto struct {
	TriggerId   string   `json:"triggerId,omitempty"`
	TriggerName string   `json:"triggerName,omitempty"`
	Action      string   `json:"action,omitempty"`
	Status      string   `json:"status,omitempty"`
	Reason      string   `json:"reason,omitempty"`
	Queue       string   `json:"queue,omitempty"`
	MessageId   string   `json:"messageId,omitempty"`
	Deliveries  *float64 `json:"deliveries,omitempty"`
	NorbixEvent string   `json:"norbixEvent,omitempty"`
	At          string   `json:"at,omitempty"`
}

// PaymentsWebhookRunDto DTO.
type PaymentsWebhookRunDto struct {
	Id              string                         `json:"id,omitempty"`
	CallId          string                         `json:"callId,omitempty"`
	Env             string                         `json:"env,omitempty"`
	IntegrationId   string                         `json:"integrationId,omitempty"`
	Source          string                         `json:"source,omitempty"`
	ReceivedOn      string                         `json:"receivedOn,omitempty"`
	StatusCode      float64                        `json:"statusCode,omitempty"`
	Outcome         string                         `json:"outcome,omitempty"`
	Reason          string                         `json:"reason,omitempty"`
	Description     string                         `json:"description,omitempty"`
	EventName       string                         `json:"eventName,omitempty"`
	ProviderEventId string                         `json:"providerEventId,omitempty"`
	CorrelationId   string                         `json:"correlationId,omitempty"`
	Fact            *PaymentEventDto               `json:"fact,omitempty"`
	Order           *PaymentsWebhookRunOrderDto    `json:"order,omitempty"`
	Actions         []*PaymentsWebhookRunActionDto `json:"actions,omitempty"`
	Part            float64                        `json:"part,omitempty"`
	Parts           float64                        `json:"parts,omitempty"`
}

// PaymentsWebhookRawCallDto DTO.
type PaymentsWebhookRawCallDto struct {
	CallId       string            `json:"callId,omitempty"`
	ReceivedOn   string            `json:"receivedOn,omitempty"`
	ContentType  string            `json:"contentType,omitempty"`
	Body         string            `json:"body,omitempty"`
	Truncated    bool              `json:"truncated"`
	MaskedFields float64           `json:"maskedFields,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
	AnswerStatus float64           `json:"answerStatus,omitempty"`
	AnswerBody   string            `json:"answerBody,omitempty"`
	Parts        float64           `json:"parts,omitempty"`
}

// PaymentsWebhookRunDetailsDto DTO.
type PaymentsWebhookRunDetailsDto struct {
	Run *PaymentsWebhookRunDto     `json:"run,omitempty"`
	Raw *PaymentsWebhookRawCallDto `json:"raw,omitempty"`
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

// AgentOnboardingSnippet DTO.
type AgentOnboardingSnippet struct {
	Client      string `json:"client,omitempty"`
	Config      string `json:"config,omitempty"`
	Auth        string `json:"auth,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Recommended bool   `json:"recommended"`
}

// AiProjectContextCallerWireDto DTO.
type AiProjectContextCallerWireDto struct {
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	AccountId string `json:"accountId,omitempty"`
}

// AiProjectContextIntegrationWireDto DTO.
type AiProjectContextIntegrationWireDto struct {
	Id        string   `json:"id,omitempty"`
	Name      string   `json:"name,omitempty"`
	Kind      string   `json:"kind,omitempty"`
	Enabled   *bool    `json:"enabled,omitempty"`
	DefaultIn []string `json:"defaultIn,omitempty"`
	Env       string   `json:"env,omitempty"`
}

// AiProjectContextModuleWireDto DTO.
type AiProjectContextModuleWireDto struct {
	Name                 string                                `json:"name,omitempty"`
	IntegrationsReadable *bool                                 `json:"integrationsReadable,omitempty"`
	Integrations         []*AiProjectContextIntegrationWireDto `json:"integrations,omitempty"`
}

// AiProjectContextProjectWireDto DTO.
type AiProjectContextProjectWireDto struct {
	Id                 string                           `json:"id,omitempty"`
	Name               string                           `json:"name,omitempty"`
	UniqueName         string                           `json:"uniqueName,omitempty"`
	Access             string                           `json:"access,omitempty"`
	Description        string                           `json:"description,omitempty"`
	DefaultLanguage    string                           `json:"defaultLanguage,omitempty"`
	Languages          []string                         `json:"languages,omitempty"`
	Environments       []string                         `json:"environments,omitempty"`
	PrimaryRegion      string                           `json:"primaryRegion,omitempty"`
	AdditionalRegions  []string                         `json:"additionalRegions,omitempty"`
	CorsOrigins        []string                         `json:"corsOrigins,omitempty"`
	ApiHost            string                           `json:"apiHost,omitempty"`
	AdminUrl           string                           `json:"adminUrl,omitempty"`
	MarketingUrl       string                           `json:"marketingUrl,omitempty"`
	AdminPortalEnabled *bool                            `json:"adminPortalEnabled,omitempty"`
	Modules            []*AiProjectContextModuleWireDto `json:"modules,omitempty"`
}

// AiProjectContextWireDto DTO.
type AiProjectContextWireDto struct {
	Version  string                            `json:"version,omitempty"`
	Caller   *AiProjectContextCallerWireDto    `json:"caller,omitempty"`
	Projects []*AiProjectContextProjectWireDto `json:"projects,omitempty"`
	Hints    []string                          `json:"hints,omitempty"`
}

// AiToolManifestParameter DTO.
type AiToolManifestParameter struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required"`
	Description string `json:"description,omitempty"`
}

// AiToolManifestItem DTO.
type AiToolManifestItem struct {
	Name                 string                     `json:"name,omitempty"`
	Description          string                     `json:"description,omitempty"`
	Toolsets             []string                   `json:"toolsets,omitempty"`
	RequiresConfirmation bool                       `json:"requiresConfirmation"`
	Parameters           []*AiToolManifestParameter `json:"parameters,omitempty"`
}

// ChatModelOption DTO.
type ChatModelOption struct {
	LlmIntegrationId string  `json:"llmIntegrationId,omitempty"`
	Kind             string  `json:"kind,omitempty"`
	Provider         string  `json:"provider,omitempty"`
	Model            string  `json:"model,omitempty"`
	Label            string  `json:"label,omitempty"`
	IsDefault        bool    `json:"isDefault"`
	IsAuto           bool    `json:"isAuto"`
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
	IsArchived   bool   `json:"isArchived"`
	IsPinned     bool   `json:"isPinned"`
}

// ScaffoldStep DTO.
type ScaffoldStep struct {
	Order       float64 `json:"order,omitempty"`
	Kind        string  `json:"kind,omitempty"`
	Title       string  `json:"title,omitempty"`
	Tool        string  `json:"tool,omitempty"`
	Arguments   any     `json:"arguments,omitempty"`
	CheckTool   string  `json:"checkTool,omitempty"`
	TemplateRef string  `json:"templateRef,omitempty"`
}

// ScaffoldIssue DTO.
type ScaffoldIssue struct {
	Where   string `json:"where,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ScaffoldStepReportWithLink DTO.
type ScaffoldStepReportWithLink struct {
	Order        float64  `json:"order,omitempty"`
	Title        string   `json:"title,omitempty"`
	Tool         string   `json:"tool,omitempty"`
	Status       string   `json:"status,omitempty"`
	Id           string   `json:"id,omitempty"`
	DashboardUrl string   `json:"dashboardUrl,omitempty"`
	Note         string   `json:"note,omitempty"`
	Errors       []string `json:"errors,omitempty"`
}

// ScaffoldApplyReport DTO.
type ScaffoldApplyReport struct {
	Completed  bool                          `json:"completed"`
	ProjectId  string                        `json:"projectId,omitempty"`
	ProjectUrl string                        `json:"projectUrl,omitempty"`
	Summary    string                        `json:"summary,omitempty"`
	Steps      []*ScaffoldStepReportWithLink `json:"steps,omitempty"`
}

// TemplatePreviewPartResult DTO.
type TemplatePreviewPartResult struct {
	Name     string   `json:"name,omitempty"`
	Rendered string   `json:"rendered,omitempty"`
	Errors   []string `json:"errors,omitempty"`
}

// TemplatePreview DTO.
type TemplatePreview struct {
	Channel string                       `json:"channel,omitempty"`
	Ok      bool                         `json:"ok"`
	Parts   []*TemplatePreviewPartResult `json:"parts,omitempty"`
	Errors  []string                     `json:"errors,omitempty"`
}

// ProjectBriefSourceWireDto DTO.
type ProjectBriefSourceWireDto struct {
	Kind       string   `json:"kind,omitempty"`
	SessionId  string   `json:"sessionId,omitempty"`
	EntryId    string   `json:"entryId,omitempty"`
	EntrySeq   *float64 `json:"entrySeq,omitempty"`
	EventId    string   `json:"eventId,omitempty"`
	UserAuthId string   `json:"userAuthId,omitempty"`
	AtUtc      string   `json:"atUtc,omitempty"`
	Surface    string   `json:"surface,omitempty"`
	Quote      string   `json:"quote,omitempty"`
	WorkItemId string   `json:"workItemId,omitempty"`
}

// ProjectBriefSatisfiedByWireDto DTO.
type ProjectBriefSatisfiedByWireDto struct {
	ArtifactId string `json:"artifactId,omitempty"`
	Tool       string `json:"tool,omitempty"`
	Orphaned   bool   `json:"orphaned"`
}

// ProjectBriefRequirementWireDto DTO.
type ProjectBriefRequirementWireDto struct {
	Id           string                            `json:"id,omitempty"`
	Text         string                            `json:"text,omitempty"`
	Status       string                            `json:"status,omitempty"`
	Confidence   float64                           `json:"confidence,omitempty"`
	IsAssumption bool                              `json:"isAssumption"`
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
	Done             bool                     `json:"done"`
	DoneAtUtc        string                   `json:"doneAtUtc,omitempty"`
	DoneByUserAuthId string                   `json:"doneByUserAuthId,omitempty"`
}

// WorkItemOpenQuestionWireDto DTO.
type WorkItemOpenQuestionWireDto struct {
	EntryRef *WorkItemEntryRefWireDto `json:"entryRef,omitempty"`
	Blocking bool                     `json:"blocking"`
	Text     string                   `json:"text,omitempty"`
	AtUtc    string                   `json:"atUtc,omitempty"`
}

// WorkItemDoneConditionWireDto DTO.
type WorkItemDoneConditionWireDto struct {
	Condition float64 `json:"condition,omitempty"`
	Holds     bool    `json:"holds"`
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
	Difficulty          *float64                        `json:"difficulty,omitempty"`
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

// AiTriggerProjectionList DTO.
type AiTriggerProjectionList struct {
	TriggerProjectionList
	Type AiTriggerType `json:"type,omitempty"`
}

// AiTriggerDto DTO.
type AiTriggerDto struct {
	TriggerDto
	When AiTriggerType `json:"when,omitempty"`
}

// EmbeddingIntegrationDto DTO.
type EmbeddingIntegrationDto struct {
	IntegrationDto
	Provider     EmbeddingProvider `json:"provider,omitempty"`
	Model        string            `json:"model,omitempty"`
	Dimension    float64           `json:"dimension,omitempty"`
	BaseUrl      string            `json:"baseUrl,omitempty"`
	IsConfigured bool              `json:"isConfigured"`
}

// EmbeddingIntegrationListProjection DTO.
type EmbeddingIntegrationListProjection struct {
	IntegrationListProjection
	EmbeddingProvider EmbeddingProvider `json:"embeddingProvider,omitempty"`
	Model             string            `json:"model,omitempty"`
	Dimension         float64           `json:"dimension,omitempty"`
	IsConfigured      bool              `json:"isConfigured"`
}

// LlmIntegrationListProjection DTO.
type LlmIntegrationListProjection struct {
	IntegrationListProjection
	LlmProvider  LlmProvider          `json:"llmProvider,omitempty"`
	BaseUrl      string               `json:"baseUrl,omitempty"`
	DefaultModel string               `json:"defaultModel,omitempty"`
	IsDefault    bool                 `json:"isDefault"`
	Models       []*LlmModelOptionDto `json:"models,omitempty"`
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
	TaskId            string            `json:"taskId,omitempty"`
	Name              string            `json:"name,omitempty"`
	Cron              string            `json:"cron,omitempty"`
	Type              SchedulerTaskType `json:"type,omitempty"`
	IsEnabled         bool              `json:"isEnabled"`
	IsFailing         bool              `json:"isFailing"`
	LastFailureReason string            `json:"lastFailureReason,omitempty"`
	LastFailedAtUnix  *float64          `json:"lastFailedAtUnix,omitempty"`
	EndsAtUnix        *float64          `json:"endsAtUnix,omitempty"`
	HasEnded          bool              `json:"hasEnded"`
	ViewId            string            `json:"viewId,omitempty"`
}

// SchedulerTaskRunDto DTO.
type SchedulerTaskRunDto struct {
	LogId         string                  `json:"logId,omitempty"`
	AtUtc         string                  `json:"atUtc,omitempty"`
	AtUnix        float64                 `json:"atUnix,omitempty"`
	Outcome       SchedulerTaskRunOutcome `json:"outcome,omitempty"`
	Reason        string                  `json:"reason,omitempty"`
	Detail        string                  `json:"detail,omitempty"`
	CorrelationId string                  `json:"correlationId,omitempty"`
}

// ResolvedResourceEntry DTO.
type ResolvedResourceEntry struct {
	Ref        *ResourceRefDto   `json:"ref,omitempty"`
	Status     ResolvedRefStatus `json:"status,omitempty"`
	Resolved   any               `json:"resolved,omitempty"`
	Diagnostic string            `json:"diagnostic,omitempty"`
}

// ChannelSubscriptionStateDto DTO.
type ChannelSubscriptionStateDto struct {
	Unsubscribed bool                `json:"unsubscribed"`
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
	BirthDate            *float64                                `json:"birthDate,omitempty"`
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

// ContactRolesRepairDto DTO.
type ContactRolesRepairDto struct {
	DryRun                      bool    `json:"dryRun"`
	ContactsWithoutRoles        float64 `json:"contactsWithoutRoles,omitempty"`
	ContactsToRepair            float64 `json:"contactsToRepair,omitempty"`
	ContactsRepaired            float64 `json:"contactsRepaired,omitempty"`
	ContactsWithoutLoginRoleIds float64 `json:"contactsWithoutLoginRoleIds,omitempty"`
	LegacyRoleNamesIgnored      float64 `json:"legacyRoleNamesIgnored,omitempty"`
}

// ConsentPurposeDto DTO.
type ConsentPurposeDto struct {
	Key             string   `json:"key,omitempty"`
	Name            string   `json:"name,omitempty"`
	Channel         string   `json:"channel,omitempty"`
	MappedTags      []string `json:"mappedTags,omitempty"`
	RegulatoryBasis []string `json:"regulatoryBasis,omitempty"`
	Description     string   `json:"description,omitempty"`
	IsDeprecated    bool     `json:"isDeprecated"`
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
	AutoForwardAdvisories bool    `json:"autoForwardAdvisories"`
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
	FirstResponseOn *float64              `json:"firstResponseOn,omitempty"`
	ResolvedOn      *float64              `json:"resolvedOn,omitempty"`
	ClosedOn        *float64              `json:"closedOn,omitempty"`
	Resolution      string                `json:"resolution,omitempty"`
	MessageCount    float64               `json:"messageCount,omitempty"`
	LastMessageOn   *float64              `json:"lastMessageOn,omitempty"`
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
	LastMessageOn  *float64              `json:"lastMessageOn,omitempty"`
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
	StepId       string `json:"stepId,omitempty"`
	Kind         string `json:"kind,omitempty"`
	IsSuccess    bool   `json:"isSuccess"`
	Result       any    `json:"result,omitempty"`
	ErrorMessage string `json:"errorMessage,omitempty"`
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
	IsManagedService bool                       `json:"isManagedService"`
	ApiVersion       string                     `json:"apiVersion,omitempty"`
	HubVersion       string                     `json:"hubVersion,omitempty"`
	Release          string                     `json:"release,omitempty"`
	Runtime          string                     `json:"runtime,omitempty"`
	HubUrl           string                     `json:"hubUrl,omitempty"`
	ApiUrl           string                     `json:"apiUrl,omitempty"`
	LicensePresent   bool                       `json:"licensePresent"`
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
	HasMore  bool                      `json:"hasMore"`
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
	IsList        bool                          `json:"isList"`
	Item          map[string]any                `json:"item,omitempty"`
	ListItems     []*DiagnosticRedisListItemDto `json:"listItems,omitempty"`
	HasMore       bool                          `json:"hasMore"`
	StartingAfter string                        `json:"startingAfter,omitempty"`
}

// DiagnosticHealthCheckDto DTO.
type DiagnosticHealthCheckDto struct {
	CheckId    string  `json:"checkId,omitempty"`
	IsHealthy  bool    `json:"isHealthy"`
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

// IQueuedTrigger DTO.
type IQueuedTrigger struct {
}

// IHasResponsibleUserId DTO.
type IHasResponsibleUserId struct {
}

// ICursorArgs DTO.
type ICursorArgs struct {
}

// EchoResponse DTO.
type EchoResponse struct {
	Release                     CodeMashRelease         `json:"release,omitempty"`
	Runtime                     CodeMashRuntime         `json:"runtime,omitempty"`
	ManagedServiceHubUrl        string                  `json:"managedServiceHubUrl,omitempty"`
	ManagedServiceApiUrl        string                  `json:"managedServiceApiUrl,omitempty"`
	HubUrl                      string                  `json:"hubUrl,omitempty"`
	ApiUrl                      string                  `json:"apiUrl,omitempty"`
	ApiVersion                  string                  `json:"apiVersion,omitempty"`
	HubVersion                  string                  `json:"hubVersion,omitempty"`
	AdminUrlTemplate            string                  `json:"adminUrlTemplate,omitempty"`
	EmailServiceConfigured      bool                    `json:"emailServiceConfigured"`
	RootBootstrapPasswordSource string                  `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                     []*EchoRegionDto        `json:"regions,omitempty"`
	Agent                       *EchoAgentDto           `json:"agent,omitempty"`
	HostingPlatform             *EchoHostingPlatformDto `json:"hostingPlatform,omitempty"`
	Connect                     *EchoConnectDto         `json:"connect,omitempty"`
	Keyless                     *EchoKeylessDto         `json:"keyless,omitempty"`
}

// PublicProjectConfigDto DTO.
type PublicProjectConfigDto struct {
	DisplayName        string           `json:"displayName,omitempty"`
	AdminPortalEnabled bool             `json:"adminPortalEnabled"`
	Branding           *PublicBrandDto  `json:"branding,omitempty"`
	Auth               *PublicAuthDto   `json:"auth,omitempty"`
	AiChat             *PublicAiChatDto `json:"aiChat,omitempty"`
}

// PublicLegalDocumentDto DTO.
type PublicLegalDocumentDto struct {
	Kind      string `json:"kind,omitempty"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
	Available bool   `json:"available"`
}

// GetTriggerDefinitionResponse DTO.
type GetTriggerDefinitionResponse struct {
	ResponseBase
	Trigger *SaveTriggerRequest `json:"trigger,omitempty"`
}

// GetTriggersNeedingAttentionResponse DTO.
type GetTriggersNeedingAttentionResponse struct {
	ResponseBase
	Items []*TriggerAttentionDto `json:"items,omitempty"`
}

// TestTriggerPreExecuteCodeResponse DTO.
type TestTriggerPreExecuteCodeResponse struct {
	ResponseBase
	Outcome     string                          `json:"outcome,omitempty"`
	TriggerRuns bool                            `json:"triggerRuns"`
	Problems    []*TriggerActivationTestProblem `json:"problems,omitempty"`
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

// CreateStripeCheckoutSessionResponse DTO.
type CreateStripeCheckoutSessionResponse struct {
	IdResponse
	TrialIncluded bool   `json:"trialIncluded"`
	NoTrialReason string `json:"noTrialReason,omitempty"`
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

// RepairEnvironmentBuiltInsResponse DTO.
type RepairEnvironmentBuiltInsResponse struct {
	ResponseBase
	Items []*BuiltInRepairDto `json:"items,omitempty"`
}

// GetProjectEnvironmentsResponse DTO.
type GetProjectEnvironmentsResponse struct {
	ResponseBase
	Item *ProjectEnvironmentsDto `json:"item,omitempty"`
}

// GetEnvironmentAdminUrlResponse DTO.
type GetEnvironmentAdminUrlResponse struct {
	ResponseBase
	Item *EnvironmentAdminUrlDto `json:"item,omitempty"`
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
	IsActive      bool    `json:"isActive"`
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
	AdminPortalEnabled bool                    `json:"adminPortalEnabled"`
	DisplayName        string                  `json:"displayName,omitempty"`
	Modules            []*AdminPortalModuleDto `json:"modules,omitempty"`
}

// CreateAiCreditPackCheckoutResponse DTO.
type CreateAiCreditPackCheckoutResponse struct {
	ResponseBase
	Result *AiCreditPackCheckoutDto `json:"result,omitempty"`
}

// GetProjectAiPlansResponse DTO.
type GetProjectAiPlansResponse struct {
	ResponseBase
	Result *ProjectAiPlansDto `json:"result,omitempty"`
}

// UpdateProjectAiPlansResponse DTO.
type UpdateProjectAiPlansResponse struct {
	ResponseBase
	PlanIds []string `json:"planIds,omitempty"`
}

// GetProjectAiUserPlansResponse DTO.
type GetProjectAiUserPlansResponse struct {
	ResponseBase
	Result *AiUserPlansDto `json:"result,omitempty"`
}

// GetProjectAiSettingsResponse DTO.
type GetProjectAiSettingsResponse struct {
	ResponseBase
	Result *ProjectAiSettingsDto `json:"result,omitempty"`
}

// GetProjectAiUsageResponse DTO.
type GetProjectAiUsageResponse struct {
	ResponseBase
	Result *ProjectAiUsageDto `json:"result,omitempty"`
}

// CheckProjectLanguagesResponse DTO.
type CheckProjectLanguagesResponse struct {
	ResponseBase
	Templates []*TemplateLanguageGapDto `json:"templates,omitempty"`
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

// GetAccountTeamInvitationsResponse DTO.
type GetAccountTeamInvitationsResponse struct {
	ResponseBase
	Items []*TeamInvitationDto `json:"items,omitempty"`
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

// GetMyAccountUserProfileResponse DTO.
type GetMyAccountUserProfileResponse struct {
	ResponseBase
	Item *AuthDto `json:"item,omitempty"`
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

// GetLicenseKeyResponse DTO.
type GetLicenseKeyResponse struct {
	ResponseBase
	Key string `json:"key,omitempty"`
}

// CreateAiServiceUserResponse DTO.
type CreateAiServiceUserResponse struct {
	ResponseBase
	Id    string `json:"id,omitempty"`
	KeyId string `json:"keyId,omitempty"`
	Key   string `json:"key,omitempty"`
}

// ListAiServiceUsersResponse DTO.
type ListAiServiceUsersResponse struct {
	ResponseBase
	Items []*AiServiceUserDto `json:"items,omitempty"`
}

// RotateAiServiceUserKeyResponse DTO.
type RotateAiServiceUserKeyResponse struct {
	ResponseBase
	KeyId string `json:"keyId,omitempty"`
	Key   string `json:"key,omitempty"`
}

// ImportUsersResponse DTO.
type ImportUsersResponse struct {
	ResponseBase
	Result *UsersImportDto `json:"result,omitempty"`
}

// RepairLoginProjectIdsResponse DTO.
type RepairLoginProjectIdsResponse struct {
	ResponseBase
	Result *LoginProjectRepairDto `json:"result,omitempty"`
}

// IssueServiceUserApiKeyResponse DTO.
type IssueServiceUserApiKeyResponse struct {
	Id             float64                 `json:"id,omitempty"`
	Name           string                  `json:"name,omitempty"`
	Key            string                  `json:"key,omitempty"`
	Environment    string                  `json:"environment,omitempty"`
	ResponseStatus *CodeMashResponseStatus `json:"responseStatus,omitempty"`
}

// ListServiceUserApiKeysResponse DTO.
type ListServiceUserApiKeysResponse struct {
	Keys           []*ServiceUserApiKeyDto `json:"keys,omitempty"`
	ResponseStatus *CodeMashResponseStatus `json:"responseStatus,omitempty"`
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
	Result bool `json:"result"`
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

// EndUserOwnRecordsReportResponse DTO.
type EndUserOwnRecordsReportResponse struct {
	ResponseBase
	Result *EndUserOwnRecordsReportDto `json:"result,omitempty"`
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

// GetDatabaseSchemaIndexStatusResponse DTO.
type GetDatabaseSchemaIndexStatusResponse struct {
	ResponseBase
	Status *SchemaIndexStatusDto `json:"status,omitempty"`
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
	Result []any `json:"result,omitempty"`
}

// CountRecordsResponse DTO.
type CountRecordsResponse struct {
	ResponseBase
	Count float64 `json:"count,omitempty"`
}

// DistinctRecordValuesResponse DTO.
type DistinctRecordValuesResponse struct {
	ResponseBase
	Values []any `json:"values,omitempty"`
}

// ExecuteRecordsAggregateResponse DTO.
type ExecuteRecordsAggregateResponse struct {
	ResponseBase
	Result []any `json:"result,omitempty"`
}

// FindRecordsResponse DTO.
type FindRecordsResponse struct {
	ResponseBase
	List PaginatedResponse[any] `json:"list,omitempty"`
}

// FindOneRecordResponse DTO.
type FindOneRecordResponse struct {
	ResponseBase
	Result any `json:"result,omitempty"`
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
	Result []any `json:"result,omitempty"`
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

// GetKeylessIntegrationSetupResponse DTO.
type GetKeylessIntegrationSetupResponse struct {
	ResponseBase
	IntegrationId     string                             `json:"integrationId,omitempty"`
	Env               string                             `json:"env,omitempty"`
	Module            string                             `json:"module,omitempty"`
	Cloud             string                             `json:"cloud,omitempty"`
	Issuer            string                             `json:"issuer,omitempty"`
	Available         bool                               `json:"available"`
	UnavailableReason string                             `json:"unavailableReason,omitempty"`
	Subject           string                             `json:"subject,omitempty"`
	Audience          string                             `json:"audience,omitempty"`
	Script            string                             `json:"script,omitempty"`
	Templates         []*KeylessIntegrationSetupTemplate `json:"templates,omitempty"`
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
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl string              `json:"publicUrl,omitempty"`
}

// GetFileByIdResponse DTO.
type GetFileByIdResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
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

// GetEmailPreferencesByLinkResponse DTO.
type GetEmailPreferencesByLinkResponse struct {
	ResponseBase
	Item *EmailLinkPreferencesDto `json:"item,omitempty"`
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

// RetryFailedEmailNotificationsResponse DTO.
type RetryFailedEmailNotificationsResponse struct {
	ResponseBase
	RetryRequestId string             `json:"retryRequestId,omitempty"`
	Selected       float64            `json:"selected,omitempty"`
	Skipped        map[string]float64 `json:"skipped,omitempty"`
	Replayed       bool               `json:"replayed"`
	Warning        string             `json:"warning,omitempty"`
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

// RetryFailedSmsNotificationsResponse DTO.
type RetryFailedSmsNotificationsResponse struct {
	ResponseBase
	RetryRequestId string             `json:"retryRequestId,omitempty"`
	Selected       float64            `json:"selected,omitempty"`
	Skipped        map[string]float64 `json:"skipped,omitempty"`
	Replayed       bool               `json:"replayed"`
	Warning        string             `json:"warning,omitempty"`
}

// GetSmsCampaignMessagesResponse DTO.
type GetSmsCampaignMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*SmsCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// WorkflowResponse DTO.
type WorkflowResponse struct {
	ResponseBase
	Workflow *WorkflowDto `json:"workflow,omitempty"`
}

// WorkflowVersionsResponse DTO.
type WorkflowVersionsResponse struct {
	ResponseBase
	Items []*WorkflowVersionSummaryDto `json:"items,omitempty"`
}

// WorkflowVersionDiffResponse DTO.
type WorkflowVersionDiffResponse struct {
	ResponseBase
	Diff *WorkflowVersionDiffDto `json:"diff,omitempty"`
}

// ValidateCodeWorkflowResponse DTO.
type ValidateCodeWorkflowResponse struct {
	ResponseBase
	IsValid  bool                          `json:"isValid"`
	Errors   []*WorkflowValidationIssueDto `json:"errors,omitempty"`
	Warnings []*WorkflowValidationIssueDto `json:"warnings,omitempty"`
}

// WorkflowsResponse DTO.
type WorkflowsResponse struct {
	ResponseBase
	List PaginatedResponse[*WorkflowListProjection] `json:"list,omitempty"`
}

// WorkflowSchemaResponse DTO.
type WorkflowSchemaResponse struct {
	ResponseBase
	Schema *WorkflowSchemaDto `json:"schema,omitempty"`
}

// InvokeCodeWorkflowResponse DTO.
type InvokeCodeWorkflowResponse struct {
	ResponseBase
	RunId      string               `json:"runId,omitempty"`
	Status     string               `json:"status,omitempty"`
	Output     any                  `json:"output,omitempty"`
	DurationMs *float64             `json:"durationMs,omitempty"`
	Error      *WorkflowRunErrorDto `json:"error,omitempty"`
	Run        *WorkflowRunDto      `json:"run,omitempty"`
}

// WorkflowRunsResponse DTO.
type WorkflowRunsResponse struct {
	ResponseBase
	List PaginatedResponse[*WorkflowRunListProjection] `json:"list,omitempty"`
}

// WorkflowRunResponse DTO.
type WorkflowRunResponse struct {
	ResponseBase
	Run *WorkflowRunDto `json:"run,omitempty"`
}

// WorkflowRunLogsResponse DTO.
type WorkflowRunLogsResponse struct {
	ResponseBase
	List PaginatedResponse[*TenantLogEntryDto] `json:"list,omitempty"`
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

// GetMarketplaceIntegrationAccessResponse DTO.
type GetMarketplaceIntegrationAccessResponse struct {
	ResponseBase
	Roles     []*MarketplaceAccessRoleDto     `json:"roles,omitempty"`
	Packages  []*MarketplaceAccessPackageDto  `json:"packages,omitempty"`
	Tools     []*MarketplaceAccessToolDto     `json:"tools,omitempty"`
	Functions []*MarketplaceAccessFunctionDto `json:"functions,omitempty"`
	CanManage bool                            `json:"canManage"`
}

// SaveMarketplaceIntegrationCustomAccessResponse DTO.
type SaveMarketplaceIntegrationCustomAccessResponse struct {
	ResponseBase
	PolicyId    string  `json:"policyId,omitempty"`
	Created     bool    `json:"created"`
	UsedByRoles float64 `json:"usedByRoles,omitempty"`
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

// GetMarketplaceIntegrationToolsResponse DTO.
type GetMarketplaceIntegrationToolsResponse struct {
	ResponseBase
	Tools       []*MarketplaceMcpToolDto `json:"tools,omitempty"`
	ListedAt    string                   `json:"listedAt,omitempty"`
	CatalogHash string                   `json:"catalogHash,omitempty"`
	Truncated   bool                     `json:"truncated"`
}

// StartMarketplaceIntegrationSignInResponse DTO.
type StartMarketplaceIntegrationSignInResponse struct {
	ResponseBase
	AuthorizeUrl string `json:"authorizeUrl,omitempty"`
	RedirectUrl  string `json:"redirectUrl,omitempty"`
	ExpiresAt    string `json:"expiresAt,omitempty"`
}

// SignOutMarketplaceIntegrationResponse DTO.
type SignOutMarketplaceIntegrationResponse struct {
	ResponseBase
	WasSignedIn bool   `json:"wasSignedIn"`
	Revoke      string `json:"revoke,omitempty"`
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
	Items         []*IntegrationTestResultItemDto       `json:"items,omitempty"`
	Steps         []*MarketplaceValidationStepResultDto `json:"steps,omitempty"`
	AccessChanges *MarketplaceAccessChangesDto          `json:"accessChanges,omitempty"`
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

// GetMarketplaceFunctionRunsResponse DTO.
type GetMarketplaceFunctionRunsResponse struct {
	ResponseBase
	Runs []*MarketplaceFunctionRunDto `json:"runs,omitempty"`
}

// InvokeMarketplaceFunctionResponse DTO.
type InvokeMarketplaceFunctionResponse struct {
	ResponseBase
	IsSuccess       bool   `json:"isSuccess"`
	Output          any    `json:"output,omitempty"`
	VendorRequestId string `json:"vendorRequestId,omitempty"`
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

// GetCodeFunctionsResponse DTO.
type GetCodeFunctionsResponse struct {
	ResponseBase
	List PaginatedResponse[*CodeFunctionDescriptorDto] `json:"list,omitempty"`
}

// GetCodeFunctionEventSampleResponse DTO.
type GetCodeFunctionEventSampleResponse struct {
	ResponseBase
	Sample *CodeFunctionEventDto `json:"sample,omitempty"`
}

// GetCodeFunctionQualifiersResponse DTO.
type GetCodeFunctionQualifiersResponse struct {
	ResponseBase
	Qualifiers []*CodeFunctionQualifierDto `json:"qualifiers,omitempty"`
}

// GetCodeFunctionRunsResponse DTO.
type GetCodeFunctionRunsResponse struct {
	ResponseBase
	Runs []*CodeFunctionRunDto `json:"runs,omitempty"`
}

// InvokeCodeFunctionResponse DTO.
type InvokeCodeFunctionResponse struct {
	ResponseBase
	StatusCode float64 `json:"statusCode,omitempty"`
	Output     any     `json:"output,omitempty"`
	DurationMs float64 `json:"durationMs,omitempty"`
	RequestId  string  `json:"requestId,omitempty"`
	Logs       string  `json:"logs,omitempty"`
}

// CodeWorkflowResumeLinkResponse DTO.
type CodeWorkflowResumeLinkResponse struct {
	ResponseBase
	RunId          string `json:"runId,omitempty"`
	WorkflowId     string `json:"workflowId,omitempty"`
	WorkflowName   string `json:"workflowName,omitempty"`
	StepId         string `json:"stepId,omitempty"`
	StepName       string `json:"stepName,omitempty"`
	Mode           string `json:"mode,omitempty"`
	Decision       string `json:"decision,omitempty"`
	Key            string `json:"key,omitempty"`
	AssigneeUserId string `json:"assigneeUserId,omitempty"`
	ExpiresAt      string `json:"expiresAt,omitempty"`
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

// GetPushDeviceResponse DTO.
type GetPushDeviceResponse struct {
	ResponseBase
	Item   *PushDeviceDto `json:"item,omitempty"`
	UserId string         `json:"userId,omitempty"`
}

// GetPushDevicesResponse DTO.
type GetPushDevicesResponse struct {
	ResponseBase
	List        PaginatedResponse[*PushDeviceListProjection] `json:"list,omitempty"`
	TeamDevices []*TeamPushDeviceListProjection              `json:"teamDevices,omitempty"`
}

// GetPushCampaignAudienceCountResponse DTO.
type GetPushCampaignAudienceCountResponse struct {
	ResponseBase
	Result *PushAudienceCountDto `json:"result,omitempty"`
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

// RetryFailedPushNotificationsResponse DTO.
type RetryFailedPushNotificationsResponse struct {
	ResponseBase
	RetryRequestId string             `json:"retryRequestId,omitempty"`
	Selected       float64            `json:"selected,omitempty"`
	Skipped        map[string]float64 `json:"skipped,omitempty"`
	Replayed       bool               `json:"replayed"`
	Warning        string             `json:"warning,omitempty"`
}

// GetPushCampaignMessagesResponse DTO.
type GetPushCampaignMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*PushCampaignBatchNotificationDto] `json:"list,omitempty"`
}

// GetNotificationInboxResponse DTO.
type GetNotificationInboxResponse struct {
	ResponseBase
	List PaginatedResponse[*InAppNotificationDto] `json:"list,omitempty"`
}

// GetUnreadCountResponse DTO.
type GetUnreadCountResponse struct {
	ResponseBase
	Count float64 `json:"count,omitempty"`
}

// MarkNotificationReadResponse DTO.
type MarkNotificationReadResponse struct {
	ResponseBase
}

// MarkAllNotificationsReadResponse DTO.
type MarkAllNotificationsReadResponse struct {
	ResponseBase
}

// GetPaymentsWebhookLogResponse DTO.
type GetPaymentsWebhookLogResponse struct {
	ResponseBase
	List []*PaymentsWebhookLogEntry `json:"list,omitempty"`
}

// GetPaymentsWebhookRunsResponse DTO.
type GetPaymentsWebhookRunsResponse struct {
	ResponseBase
	List PaginatedResponse[*PaymentsWebhookRunDto] `json:"list,omitempty"`
}

// GetPaymentsWebhookRunResponse DTO.
type GetPaymentsWebhookRunResponse struct {
	ResponseBase
	Item *PaymentsWebhookRunDetailsDto `json:"item,omitempty"`
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

// GetQuickConnectProvidersResponse DTO.
type GetQuickConnectProvidersResponse struct {
	ResponseBase
	Providers []*QuickConnectProviderDto `json:"providers,omitempty"`
}

// StartIntegrationConnectResponse DTO.
type StartIntegrationConnectResponse struct {
	ResponseBase
	AuthorizeUrl string `json:"authorizeUrl,omitempty"`
	ConnectId    string `json:"connectId,omitempty"`
	RedirectUrl  string `json:"redirectUrl,omitempty"`
	ExpiresAt    string `json:"expiresAt,omitempty"`
}

// GetIntegrationConnectTargetsResponse DTO.
type GetIntegrationConnectTargetsResponse struct {
	ResponseBase
	ConnectId       string                   `json:"connectId,omitempty"`
	Module          string                   `json:"module,omitempty"`
	Provider        string                   `json:"provider,omitempty"`
	IntegrationName string                   `json:"integrationName,omitempty"`
	AccountName     string                   `json:"accountName,omitempty"`
	Targets         []*QuickConnectTargetDto `json:"targets,omitempty"`
}

// DisconnectPaymentsIntegrationConnectResponse DTO.
type DisconnectPaymentsIntegrationConnectResponse struct {
	ResponseBase
	WasConnected     bool `json:"wasConnected"`
	ProviderAnswered bool `json:"providerAnswered"`
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

// GetLogResponse DTO.
type GetLogResponse struct {
	ResponseBase
	Item *TenantLogEntryDto `json:"item,omitempty"`
}

// GetLogsResponse DTO.
type GetLogsResponse struct {
	ResponseBase
	List PaginatedResponse[*TenantLogEntryDto] `json:"list,omitempty"`
}

// GetLogSettingsResponse DTO.
type GetLogSettingsResponse struct {
	ResponseBase
	SkipCloudDashboardLogs bool   `json:"skipCloudDashboardLogs"`
	SkipHttpBodyMeta       bool   `json:"skipHttpBodyMeta"`
	AiChatLoggingEnabled   bool   `json:"aiChatLoggingEnabled"`
	HasNorbixLogging       bool   `json:"hasNorbixLogging"`
	Env                    string `json:"env,omitempty"`
	LastEntryAtUtc         string `json:"lastEntryAtUtc,omitempty"`
	LastStoredAtUtc        string `json:"lastStoredAtUtc,omitempty"`
	StorageWarning         string `json:"storageWarning,omitempty"`
}

// SaveLogSettingsResponse DTO.
type SaveLogSettingsResponse struct {
	ResponseBase
}

// GetAgentOnboardingResponse DTO.
type GetAgentOnboardingResponse struct {
	ResponseBase
	Prompt           string                    `json:"prompt,omitempty"`
	HubUrl           string                    `json:"hubUrl,omitempty"`
	ApiUrl           string                    `json:"apiUrl,omitempty"`
	McpUrl           string                    `json:"mcpUrl,omitempty"`
	DocsUrl          string                    `json:"docsUrl,omitempty"`
	InstallationType string                    `json:"installationType,omitempty"`
	Snippets         []*AgentOnboardingSnippet `json:"snippets,omitempty"`
}

// GetAiContextResponse DTO.
type GetAiContextResponse struct {
	ResponseBase
	Version   string                   `json:"version,omitempty"`
	Unchanged bool                     `json:"unchanged"`
	Context   *AiProjectContextWireDto `json:"context,omitempty"`
	Markdown  string                   `json:"markdown,omitempty"`
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

// UploadChatAttachmentResponse DTO.
type UploadChatAttachmentResponse struct {
	ResponseBase
	Id        string `json:"id,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
}

// ChatAvailabilityResponse DTO.
type ChatAvailabilityResponse struct {
	ResponseBase
	Available bool               `json:"available"`
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

// ScaffoldProjectResponse DTO.
type ScaffoldProjectResponse struct {
	ResponseBase
	Mode     string               `json:"mode,omitempty"`
	Env      string               `json:"env,omitempty"`
	Valid    bool                 `json:"valid"`
	Plan     []*ScaffoldStep      `json:"plan,omitempty"`
	Defaults []string             `json:"defaults,omitempty"`
	Issues   []*ScaffoldIssue     `json:"issues,omitempty"`
	Report   *ScaffoldApplyReport `json:"report,omitempty"`
}

// ValidateSchemaResponse DTO.
type ValidateSchemaResponse struct {
	ResponseBase
	Valid       bool             `json:"valid"`
	Issues      []*ScaffoldIssue `json:"issues,omitempty"`
	Collections []string         `json:"collections,omitempty"`
	Warnings    []*ScaffoldIssue `json:"warnings,omitempty"`
}

// RenderTemplatePreviewResponse DTO.
type RenderTemplatePreviewResponse struct {
	ResponseBase
	Preview *TemplatePreview `json:"preview,omitempty"`
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

// GetAiTriggersResponse DTO.
type GetAiTriggersResponse struct {
	GetTriggersResponse
	List PaginatedResponse[*AiTriggerProjectionList] `json:"list,omitempty"`
}

// GetAiTriggerResponse DTO.
type GetAiTriggerResponse struct {
	GetTriggerResponse
	Trigger *AiTriggerDto `json:"trigger,omitempty"`
}

// AuthDeviceStartResponse DTO.
type AuthDeviceStartResponse struct {
	ResponseBase
	DeviceCode              string  `json:"deviceCode,omitempty"`
	UserCode                string  `json:"userCode,omitempty"`
	VerificationUri         string  `json:"verificationUri,omitempty"`
	VerificationUriComplete string  `json:"verificationUriComplete,omitempty"`
	ExpiresIn               float64 `json:"expiresIn,omitempty"`
	Interval                float64 `json:"interval,omitempty"`
}

// AuthDeviceTokenResponse DTO.
type AuthDeviceTokenResponse struct {
	ResponseBase
	Error            string   `json:"error,omitempty"`
	ErrorDescription string   `json:"errorDescription,omitempty"`
	ErrorCode        string   `json:"errorCode,omitempty"`
	Reason           string   `json:"reason,omitempty"`
	BearerToken      string   `json:"bearerToken,omitempty"`
	RefreshToken     string   `json:"refreshToken,omitempty"`
	ExpiresIn        *float64 `json:"expiresIn,omitempty"`
	ClientId         string   `json:"clientId,omitempty"`
	UserId           string   `json:"userId,omitempty"`
	UserName         string   `json:"userName,omitempty"`
	DisplayName      string   `json:"displayName,omitempty"`
	AccountId        string   `json:"accountId,omitempty"`
	ProjectId        string   `json:"projectId,omitempty"`
}

// AuthDeviceConsentResponse DTO.
type AuthDeviceConsentResponse struct {
	ResponseBase
	ClientName   string               `json:"clientName,omitempty"`
	DeviceName   string               `json:"deviceName,omitempty"`
	UserCode     string               `json:"userCode,omitempty"`
	ProjectId    string               `json:"projectId,omitempty"`
	AccountRoles []*AiRoleDto         `json:"accountRoles,omitempty"`
	Projects     []*AiProjectRolesDto `json:"projects,omitempty"`
}

// AuthDeviceConsentDecisionResponse DTO.
type AuthDeviceConsentDecisionResponse struct {
	ResponseBase
}

// OAuthConsentResponse DTO.
type OAuthConsentResponse struct {
	ResponseBase
	ClientName         string               `json:"clientName,omitempty"`
	RedirectHost       string               `json:"redirectHost,omitempty"`
	ClientHost         string               `json:"clientHost,omitempty"`
	RedirectIsLoopback bool                 `json:"redirectIsLoopback"`
	ConsentTicket      string               `json:"consentTicket,omitempty"`
	AccountRoles       []*AiRoleDto         `json:"accountRoles,omitempty"`
	Projects           []*AiProjectRolesDto `json:"projects,omitempty"`
}

// OAuthConsentDecisionResponse DTO.
type OAuthConsentDecisionResponse struct {
	ResponseBase
	RedirectUrl string `json:"redirectUrl,omitempty"`
}

// GetEmbeddingIntegrationResponse DTO.
type GetEmbeddingIntegrationResponse struct {
	ResponseBase
	Item *EmbeddingIntegrationDto `json:"item,omitempty"`
}

// GetEmbeddingIntegrationsResponse DTO.
type GetEmbeddingIntegrationsResponse struct {
	ResponseBase
	List PaginatedResponse[*EmbeddingIntegrationListProjection] `json:"list,omitempty"`
}

// TestEmbeddingIntegrationResponse DTO.
type TestEmbeddingIntegrationResponse struct {
	ResponseBase
	Dimension   *float64 `json:"dimension,omitempty"`
	LatencyMs   *float64 `json:"latencyMs,omitempty"`
	TotalTokens *float64 `json:"totalTokens,omitempty"`
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
	AllowsPartialResponse bool                `json:"allowsPartialResponse"`
	Options               map[string]string   `json:"options,omitempty"`
	Status                float64             `json:"status,omitempty"`
	StatusCode            any                 `json:"statusCode,omitempty"`
	StatusDescription     string              `json:"statusDescription,omitempty"`
	Response              any                 `json:"response,omitempty"`
	ResponseFilter        *IContentTypeWriter `json:"responseFilter,omitempty"`
	RequestContext        *IRequest           `json:"requestContext,omitempty"`
	View                  string              `json:"view,omitempty"`
	Template              string              `json:"template,omitempty"`
	PaddingLength         float64             `json:"paddingLength,omitempty"`
	IsPartialRequest      bool                `json:"isPartialRequest"`
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

// GetSchedulerTaskRunsResponse DTO.
type GetSchedulerTaskRunsResponse struct {
	ResponseBase
	LogsEnabled bool                   `json:"logsEnabled"`
	Runs        []*SchedulerTaskRunDto `json:"runs,omitempty"`
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

// RepairContactRolesResponse DTO.
type RepairContactRolesResponse struct {
	ResponseBase
	Result *ContactRolesRepairDto `json:"result,omitempty"`
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
	OfficialOnly *bool                  `json:"officialOnly,omitempty"`
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
	FunctionViewId string         `json:"functionViewId,omitempty"`
	Payload        map[string]any `json:"payload,omitempty"`
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

// GetQuickConnectProviders DTO.
type GetQuickConnectProviders struct {
	RequestBase
	Module string `json:"module,omitempty"`
}

// StartPaymentsIntegrationConnect DTO.
type StartPaymentsIntegrationConnect struct {
	CodeMashRequestBase
	Provider        string            `json:"provider,omitempty"`
	IntegrationName string            `json:"integrationName,omitempty"`
	Mode            string            `json:"mode,omitempty"`
	Input           map[string]string `json:"input,omitempty"`
	IntegrationId   string            `json:"integrationId,omitempty"`
}

// GetIntegrationConnectTargets DTO.
type GetIntegrationConnectTargets struct {
	CodeMashRequestBase
	ConnectId string `json:"connectId,omitempty"`
}

// CompletePaymentsIntegrationConnect DTO.
type CompletePaymentsIntegrationConnect struct {
	CodeMashRequestBase
	ConnectId string `json:"connectId,omitempty"`
	TargetId  string `json:"targetId,omitempty"`
	IsEnabled bool   `json:"isEnabled"`
}

// DisconnectPaymentsIntegrationConnect DTO.
type DisconnectPaymentsIntegrationConnect struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// InternalsTypeGen DTO.
type InternalsTypeGen struct {
	Typegen_0_SmtpEmailIntegrationRequest                        *SmtpEmailIntegrationRequest                       `json:"typegen_0_SmtpEmailIntegrationRequest,omitempty"`
	Typegen_1_AwsSesEmailIntegrationRequest                      *AwsSesEmailIntegrationRequest                     `json:"typegen_1_AwsSesEmailIntegrationRequest,omitempty"`
	Typegen_2_SendGridEmailIntegrationRequest                    *SendGridEmailIntegrationRequest                   `json:"typegen_2_SendGridEmailIntegrationRequest,omitempty"`
	Typegen_3_MailGunEmailIntegrationRequest                     *MailGunEmailIntegrationRequest                    `json:"typegen_3_MailGunEmailIntegrationRequest,omitempty"`
	Typegen_3a_AzureCommunicationEmailIntegrationRequest         *AzureCommunicationEmailIntegrationRequest         `json:"typegen_3a_AzureCommunicationEmailIntegrationRequest,omitempty"`
	Typegen_297_FakeEmailIntegrationRequest                      *FakeEmailIntegrationRequest                       `json:"typegen_297_FakeEmailIntegrationRequest,omitempty"`
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
	Typegen_285_MontonioPaymentIntegrationRequest                *MontonioPaymentIntegrationRequest                 `json:"typegen_285_MontonioPaymentIntegrationRequest,omitempty"`
	Typegen_295_NeopayPaymentIntegrationRequest                  *NeopayPaymentIntegrationRequest                   `json:"typegen_295_NeopayPaymentIntegrationRequest,omitempty"`
	Typegen_287_ShopifyPaymentIntegrationRequest                 *ShopifyPaymentIntegrationRequest                  `json:"typegen_287_ShopifyPaymentIntegrationRequest,omitempty"`
	Typegen_288_WooCommercePaymentIntegrationRequest             *WooCommercePaymentIntegrationRequest              `json:"typegen_288_WooCommercePaymentIntegrationRequest,omitempty"`
	Typegen_289_MagentoPaymentIntegrationRequest                 *MagentoPaymentIntegrationRequest                  `json:"typegen_289_MagentoPaymentIntegrationRequest,omitempty"`
	Typegen_290_BraintreePaymentIntegrationRequest               *BraintreePaymentIntegrationRequest                `json:"typegen_290_BraintreePaymentIntegrationRequest,omitempty"`
	Typegen_291_AuthorizeNetPaymentIntegrationRequest            *AuthorizeNetPaymentIntegrationRequest             `json:"typegen_291_AuthorizeNetPaymentIntegrationRequest,omitempty"`
	Typegen_292_CheckOutComPaymentIntegrationRequest             *CheckOutComPaymentIntegrationRequest              `json:"typegen_292_CheckOutComPaymentIntegrationRequest,omitempty"`
	Typegen_293_WorldpayPaymentIntegrationRequest                *WorldpayPaymentIntegrationRequest                 `json:"typegen_293_WorldpayPaymentIntegrationRequest,omitempty"`
	Typegen_56_EdgeWebPushIntegrationRequest                     *EdgeWebPushIntegrationRequest                     `json:"typegen_56_EdgeWebPushIntegrationRequest,omitempty"`
	Typegen_57_ChromePluginPushIntegrationRequest                *ChromePluginPushIntegrationRequest                `json:"typegen_57_ChromePluginPushIntegrationRequest,omitempty"`
	Typegen_58_SafariPushIntegrationRequest                      *SafariPushIntegrationRequest                      `json:"typegen_58_SafariPushIntegrationRequest,omitempty"`
	Typegen_59_ChromeWebPushIntegrationRequest                   *ChromeWebPushIntegrationRequest                   `json:"typegen_59_ChromeWebPushIntegrationRequest,omitempty"`
	Typegen_60_FirefoxWebPushIntegrationRequest                  *FirefoxWebPushIntegrationRequest                  `json:"typegen_60_FirefoxWebPushIntegrationRequest,omitempty"`
	Typegen_61_AndroidFirebasePushIntegrationRequest             *AndroidFirebasePushIntegrationRequest             `json:"typegen_61_AndroidFirebasePushIntegrationRequest,omitempty"`
	Typegen_62_AppleApnsPushIntegrationRequest                   *AppleApnsPushIntegrationRequest                   `json:"typegen_62_AppleApnsPushIntegrationRequest,omitempty"`
	Typegen_298_FakePushIntegrationRequest                       *FakePushIntegrationRequest                        `json:"typegen_298_FakePushIntegrationRequest,omitempty"`
	Typegen_299_BirdSmsIntegrationRequest                        *BirdSmsIntegrationRequest                         `json:"typegen_299_BirdSmsIntegrationRequest,omitempty"`
	Typegen_300_PlivoSmsIntegrationRequest                       *PlivoSmsIntegrationRequest                        `json:"typegen_300_PlivoSmsIntegrationRequest,omitempty"`
	Typegen_301_SinchSmsIntegrationRequest                       *SinchSmsIntegrationRequest                        `json:"typegen_301_SinchSmsIntegrationRequest,omitempty"`
	Typegen_302_TelesignSmsIntegrationRequest                    *TelesignSmsIntegrationRequest                     `json:"typegen_302_TelesignSmsIntegrationRequest,omitempty"`
	Typegen_303_TelnyxSmsIntegrationRequest                      *TelnyxSmsIntegrationRequest                       `json:"typegen_303_TelnyxSmsIntegrationRequest,omitempty"`
	Typegen_304_TwilioSmsIntegrationRequest                      *TwilioSmsIntegrationRequest                       `json:"typegen_304_TwilioSmsIntegrationRequest,omitempty"`
	Typegen_305_VonageSmsIntegrationRequest                      *VonageSmsIntegrationRequest                       `json:"typegen_305_VonageSmsIntegrationRequest,omitempty"`
	Typegen_306_FakeSmsIntegrationRequest                        *FakeSmsIntegrationRequest                         `json:"typegen_306_FakeSmsIntegrationRequest,omitempty"`
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
	Typegen_308_TriggerActionWorkflowDto                         *TriggerActionWorkflowDto                          `json:"typegen_308_TriggerActionWorkflowDto,omitempty"`
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
	Typegen_249_SmsToAccountUsersDeliverySettingsDto             *SmsToAccountUsersDeliverySettingsDto              `json:"typegen_249_SmsToAccountUsersDeliverySettingsDto,omitempty"`
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
	Typegen_114_AwsLambdaKeylessCodeIntegrationDto               *AwsLambdaKeylessCodeIntegrationDto                `json:"typegen_114_AwsLambdaKeylessCodeIntegrationDto,omitempty"`
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
	Typegen_286_MontonioPaymentIntegrationDto                    *MontonioPaymentIntegrationDto                     `json:"typegen_286_MontonioPaymentIntegrationDto,omitempty"`
	Typegen_296_NeopayPaymentIntegrationDto                      *NeopayPaymentIntegrationDto                       `json:"typegen_296_NeopayPaymentIntegrationDto,omitempty"`
	Typegen_294_PaymentOrderEventDto                             *PaymentOrderEventDto                              `json:"typegen_294_PaymentOrderEventDto,omitempty"`
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
	Typegen_154_AwsS3KeylessFilesIntegrationDto                  *AwsS3KeylessFilesIntegrationDto                   `json:"typegen_154_AwsS3KeylessFilesIntegrationDto,omitempty"`
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
	Typegen_179_AwsKeylessEmailIntegrationDto                    *AwsKeylessEmailIntegrationDto                     `json:"typegen_179_AwsKeylessEmailIntegrationDto,omitempty"`
	Typegen_180_AwsIamEmailIntegrationDto                        *AwsIamEmailIntegrationDto                         `json:"typegen_180_AwsIamEmailIntegrationDto,omitempty"`
	Typegen_181_MailGunEmailIntegrationDto                       *MailGunEmailIntegrationDto                        `json:"typegen_181_MailGunEmailIntegrationDto,omitempty"`
	Typegen_182_SendGridEmailIntegrationDto                      *SendGridEmailIntegrationDto                       `json:"typegen_182_SendGridEmailIntegrationDto,omitempty"`
	Typegen_183_SmtpEmailIntegrationDto                          *SmtpEmailIntegrationDto                           `json:"typegen_183_SmtpEmailIntegrationDto,omitempty"`
	Typegen_307_AzureCommunicationEmailIntegrationDto            *AzureCommunicationEmailIntegrationDto             `json:"typegen_307_AzureCommunicationEmailIntegrationDto,omitempty"`
	Typegen_248_FakeEmailIntegrationDto                          *FakeEmailIntegrationDto                           `json:"typegen_248_FakeEmailIntegrationDto,omitempty"`
	Typegen_192_WebhookIntegrationDto                            *WebhookIntegrationDto                             `json:"typegen_192_WebhookIntegrationDto,omitempty"`
	Typegen_193_WebhookDestinationDto                            *WebhookDestinationDto                             `json:"typegen_193_WebhookDestinationDto,omitempty"`
	Typegen_194_SchedulerTaskDto                                 *SchedulerTaskDto                                  `json:"typegen_194_SchedulerTaskDto,omitempty"`
	Typegen_249_EmailCampaignSchedulerTaskRequest                *EmailCampaignSchedulerTaskRequest                 `json:"typegen_249_EmailCampaignSchedulerTaskRequest,omitempty"`
	Typegen_250_SmsCampaignSchedulerTaskRequest                  *SmsCampaignSchedulerTaskRequest                   `json:"typegen_250_SmsCampaignSchedulerTaskRequest,omitempty"`
	Typegen_251_PushCampaignSchedulerTaskRequest                 *PushCampaignSchedulerTaskRequest                  `json:"typegen_251_PushCampaignSchedulerTaskRequest,omitempty"`
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
	Typegen_269_FunctionRefDto                                   *FunctionRefDto                                    `json:"typegen_269_FunctionRefDto,omitempty"`
	Typegen_270_CodeFunctionEventDto                             *CodeFunctionEventDto                              `json:"typegen_270_CodeFunctionEventDto,omitempty"`
	Typegen_271_CodeFunctionContextDto                           *CodeFunctionContextDto                            `json:"typegen_271_CodeFunctionContextDto,omitempty"`
	Typegen_272_CodeFunctionCallerDto                            *CodeFunctionCallerDto                             `json:"typegen_272_CodeFunctionCallerDto,omitempty"`
	Typegen_273_CodeFunctionRunAsDto                             *CodeFunctionRunAsDto                              `json:"typegen_273_CodeFunctionRunAsDto,omitempty"`
	Typegen_274_CodeFunctionDataDto                              *CodeFunctionDataDto                               `json:"typegen_274_CodeFunctionDataDto,omitempty"`
	Typegen_275_DatabaseRecordChangeDto                          *DatabaseRecordChangeDto                           `json:"typegen_275_DatabaseRecordChangeDto,omitempty"`
	Typegen_276_UserChangeDto                                    *UserChangeDto                                     `json:"typegen_276_UserChangeDto,omitempty"`
	Typegen_277_FileChangeDto                                    *FileChangeDto                                     `json:"typegen_277_FileChangeDto,omitempty"`
	Typegen_278_PaymentFactDto                                   *PaymentFactDto                                    `json:"typegen_278_PaymentFactDto,omitempty"`
	Typegen_279_AiUsageEventDto                                  *AiUsageEventDto                                   `json:"typegen_279_AiUsageEventDto,omitempty"`
	Typegen_280_ScheduledTaskInfoDto                             *ScheduledTaskInfoDto                              `json:"typegen_280_ScheduledTaskInfoDto,omitempty"`
	Typegen_281_FakeCodeIntegrationDto                           *FakeCodeIntegrationDto                            `json:"typegen_281_FakeCodeIntegrationDto,omitempty"`
	Typegen_282_FakeCodeIntegrationRequest                       *FakeCodeIntegrationRequest                        `json:"typegen_282_FakeCodeIntegrationRequest,omitempty"`
	Typegen_283_CodeFunctionSchedulerTaskRequest                 *CodeFunctionSchedulerTaskRequest                  `json:"typegen_283_CodeFunctionSchedulerTaskRequest,omitempty"`
	Typegen_284_MarketplaceFunctionSchedulerTaskRequest          *MarketplaceFunctionSchedulerTaskRequest           `json:"typegen_284_MarketplaceFunctionSchedulerTaskRequest,omitempty"`
	Typegen_309_WorkflowSchedulerTaskRequest                     *WorkflowSchedulerTaskRequest                      `json:"typegen_309_WorkflowSchedulerTaskRequest,omitempty"`
	Typegen_285_AwsCloudWatchLoggingIntegrationRequest           *AwsCloudWatchLoggingIntegrationRequest            `json:"typegen_285_AwsCloudWatchLoggingIntegrationRequest,omitempty"`
	Typegen_286_AwsCloudWatchLoggingIntegrationDto               *AwsCloudWatchLoggingIntegrationDto                `json:"typegen_286_AwsCloudWatchLoggingIntegrationDto,omitempty"`
	Typegen_310_GetQuickConnectProviders                         *GetQuickConnectProviders                          `json:"typegen_310_GetQuickConnectProviders,omitempty"`
	Typegen_311_StartPaymentsIntegrationConnect                  *StartPaymentsIntegrationConnect                   `json:"typegen_311_StartPaymentsIntegrationConnect,omitempty"`
	Typegen_312_GetIntegrationConnectTargets                     *GetIntegrationConnectTargets                      `json:"typegen_312_GetIntegrationConnectTargets,omitempty"`
	Typegen_313_CompletePaymentsIntegrationConnect               *CompletePaymentsIntegrationConnect                `json:"typegen_313_CompletePaymentsIntegrationConnect,omitempty"`
	Typegen_314_QuickConnectProviderDto                          *QuickConnectProviderDto                           `json:"typegen_314_QuickConnectProviderDto,omitempty"`
	Typegen_315_QuickConnectTargetDto                            *QuickConnectTargetDto                             `json:"typegen_315_QuickConnectTargetDto,omitempty"`
	Typegen_316_DisconnectPaymentsIntegrationConnect             *DisconnectPaymentsIntegrationConnect              `json:"typegen_316_DisconnectPaymentsIntegrationConnect,omitempty"`
	Typegen_317_PaymentsIntegrationConnectionDto                 *PaymentsIntegrationConnectionDto                  `json:"typegen_317_PaymentsIntegrationConnectionDto,omitempty"`
}

// SharedTypeGen DTO.
type SharedTypeGen struct {
	Typegen_StringFieldDto              *StringFieldDto              `json:"typegen_StringFieldDto,omitempty"`
	Typegen_DecimalFieldDto             *DecimalFieldDto             `json:"typegen_DecimalFieldDto,omitempty"`
	Typegen_CurrencyFieldDto            *CurrencyFieldDto            `json:"typegen_CurrencyFieldDto,omitempty"`
	Typegen_BooleanFieldDto             *BooleanFieldDto             `json:"typegen_BooleanFieldDto,omitempty"`
	Typegen_DateFieldDto                *DateFieldDto                `json:"typegen_DateFieldDto,omitempty"`
	Typegen_IntegerFieldDto             *IntegerFieldDto             `json:"typegen_IntegerFieldDto,omitempty"`
	Typegen_GeolocationFieldDto         *GeolocationFieldDto         `json:"typegen_GeolocationFieldDto,omitempty"`
	Typegen_TagsFieldDto                *TagsFieldDto                `json:"typegen_TagsFieldDto,omitempty"`
	Typegen_FileFieldDto                *FileFieldDto                `json:"typegen_FileFieldDto,omitempty"`
	Typegen_TaxonomySelectionFieldDto   *TaxonomySelectionFieldDto   `json:"typegen_TaxonomySelectionFieldDto,omitempty"`
	Typegen_CollectionSelectionFieldDto *CollectionSelectionFieldDto `json:"typegen_CollectionSelectionFieldDto,omitempty"`
	Typegen_UserSelectionFieldDto       *UserSelectionFieldDto       `json:"typegen_UserSelectionFieldDto,omitempty"`
	Typegen_RoleSelectionFieldDto       *RoleSelectionFieldDto       `json:"typegen_RoleSelectionFieldDto,omitempty"`
	Typegen_EnumSelectionFieldDto       *EnumSelectionFieldDto       `json:"typegen_EnumSelectionFieldDto,omitempty"`
	Typegen_ObjectFieldDto              *ObjectFieldDto              `json:"typegen_ObjectFieldDto,omitempty"`
	Typegen_ArrayFieldDto               *ArrayFieldDto               `json:"typegen_ArrayFieldDto,omitempty"`
	Typegen_JsonFieldDto                *JsonFieldDto                `json:"typegen_JsonFieldDto,omitempty"`
}

// Echo DTO.
type Echo struct {
	RequestBase
}

// GetPublicProjectBrandAsset DTO.
type GetPublicProjectBrandAsset struct {
	RequestBase
	ProjectId string `json:"projectId,omitempty"`
	Kind      string `json:"kind,omitempty"`
	V         string `json:"v,omitempty"`
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

// GetTriggerDefinition DTO.
type GetTriggerDefinition struct {
	CodeMashRequestBase
	Id          string      `json:"id,omitempty"`
	TriggerType TriggerType `json:"triggerType,omitempty"`
	SchemaId    string      `json:"schemaId,omitempty"`
}

// GetTriggersNeedingAttention DTO.
type GetTriggersNeedingAttention struct {
	CodeMashRequestBase
	TriggerType TriggerType `json:"triggerType,omitempty"`
}

// TestTriggerPreExecuteCode DTO.
type TestTriggerPreExecuteCode struct {
	CodeMashRequestBase
	Code        string      `json:"code,omitempty"`
	TriggerType TriggerType `json:"triggerType,omitempty"`
	NewRecord   string      `json:"newRecord,omitempty"`
	OldRecord   string      `json:"oldRecord,omitempty"`
	EventKind   string      `json:"eventKind,omitempty"`
	EventName   string      `json:"eventName,omitempty"`
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
	SourceEnv          string                           `json:"sourceEnv,omitempty"`
	TargetEnv          string                           `json:"targetEnv,omitempty"`
	DryRun             bool                             `json:"dryRun"`
	ConfirmRecordLoss  bool                             `json:"confirmRecordLoss"`
	IntegrationChoices []*PromotionIntegrationChoiceDto `json:"integrationChoices,omitempty"`
}

// RollbackPromotionRequest DTO.
type RollbackPromotionRequest struct {
	CodeMashRequestBase
	TargetEnv         string  `json:"targetEnv,omitempty"`
	FromVersion       float64 `json:"fromVersion,omitempty"`
	DryRun            bool    `json:"dryRun"`
	ConfirmRecordLoss bool    `json:"confirmRecordLoss"`
}

// RepairEnvironmentBuiltInsRequest DTO.
type RepairEnvironmentBuiltInsRequest struct {
	CodeMashRequestBase
	EnvironmentName string `json:"environmentName,omitempty"`
	DryRun          bool   `json:"dryRun"`
}

// GetProjectEnvironments DTO.
type GetProjectEnvironments struct {
	CodeMashRequestBase
}

// UpdateEnvironmentAdminUrl DTO.
type UpdateEnvironmentAdminUrl struct {
	CodeMashRequestBase
	EnvironmentName string `json:"environmentName,omitempty"`
	Url             string `json:"url,omitempty"`
}

// GetEnvironmentAdminUrl DTO.
type GetEnvironmentAdminUrl struct {
	CodeMashRequestBase
	EnvironmentName string `json:"environmentName,omitempty"`
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
	TimeoutSeconds *float64 `json:"timeoutSeconds,omitempty"`
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

// SetAdminPortalEnabledRequest DTO.
type SetAdminPortalEnabledRequest struct {
	CodeMashRequestBase
	Enabled bool `json:"enabled"`
}

// UpdateProjectExposeBrand DTO.
type UpdateProjectExposeBrand struct {
	CodeMashRequestBase
	Exposed bool `json:"exposed"`
}

// UpdateProjectExposeAuth DTO.
type UpdateProjectExposeAuth struct {
	CodeMashRequestBase
	Exposed bool `json:"exposed"`
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

// CreateAiCreditPackCheckoutRequest DTO.
type CreateAiCreditPackCheckoutRequest struct {
	CodeMashRequestBase
	Pack      string `json:"pack,omitempty"`
	ReturnUrl string `json:"returnUrl,omitempty"`
}

// GetProjectAiPlans DTO.
type GetProjectAiPlans struct {
	CodeMashRequestBase
}

// UpdateProjectAiPlans DTO.
type UpdateProjectAiPlans struct {
	CodeMashRequestBase
	Plans []*AiPlanDto `json:"plans,omitempty"`
}

// UpdateProjectAiPlanAssignments DTO.
type UpdateProjectAiPlanAssignments struct {
	CodeMashRequestBase
	Roles         []*AiPlanRoleAssignmentDto `json:"roles,omitempty"`
	DefaultPlanId string                     `json:"defaultPlanId,omitempty"`
}

// GetProjectAiUserPlans DTO.
type GetProjectAiUserPlans struct {
	CodeMashRequestBase
}

// GetProjectAiSettings DTO.
type GetProjectAiSettings struct {
	CodeMashRequestBase
}

// UpdateProjectAiSettings DTO.
type UpdateProjectAiSettings struct {
	CodeMashRequestBase
	Enabled                 bool   `json:"enabled"`
	DefaultLlmIntegrationId string `json:"defaultLlmIntegrationId,omitempty"`
	DefaultModel            string `json:"defaultModel,omitempty"`
}

// UpdateProjectAiKnowledge DTO.
type UpdateProjectAiKnowledge struct {
	CodeMashRequestBase
	EmbedFiles bool `json:"embedFiles"`
}

// CreateProjectAiAssistant DTO.
type CreateProjectAiAssistant struct {
	ProjectAiAssistantRequestBase
}

// UpdateProjectAiAssistant DTO.
type UpdateProjectAiAssistant struct {
	ProjectAiAssistantRequestBase
	AssistantId string `json:"assistantId,omitempty"`
}

// DeleteProjectAiAssistant DTO.
type DeleteProjectAiAssistant struct {
	CodeMashRequestBase
	AssistantId string `json:"assistantId,omitempty"`
}

// GetProjectAiUsage DTO.
type GetProjectAiUsage struct {
	CodeMashRequestBase
	Top *float64 `json:"top,omitempty"`
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
	Origins                 []string `json:"origins,omitempty"`
	RemoveAdminPortalOrigin *bool    `json:"removeAdminPortalOrigin,omitempty"`
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

// CheckProjectLanguages DTO.
type CheckProjectLanguages struct {
	CodeMashRequestBase
	DefaultLanguage string   `json:"defaultLanguage,omitempty"`
	Languages       []string `json:"languages,omitempty"`
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
	Exposed bool `json:"exposed"`
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
	SendInvitation bool     `json:"sendInvitation"`
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
	IncludeAccountOwner      bool     `json:"includeAccountOwner"`
	UserShouldHavePushDevice bool     `json:"userShouldHavePushDevice"`
	ProjectId                string   `json:"projectId,omitempty"`
	UserIds                  []string `json:"userIds,omitempty"`
	RoleNames                []string `json:"roleNames,omitempty"`
	StartingAfter            string   `json:"startingAfter,omitempty"`
	EndingBefore             string   `json:"endingBefore,omitempty"`
	PageSize                 *float64 `json:"pageSize,omitempty"`
}

// GetAccountPasswordPolicy DTO.
type GetAccountPasswordPolicy struct {
	RequestBase
}

// GetAccountTeamPolicies DTO.
type GetAccountTeamPolicies struct {
	RequestBase
}

// RemoveTeamMember DTO.
type RemoveTeamMember struct {
	RequestBase
	MemberId string `json:"memberId,omitempty"`
}

// SetTeamMemberRoleEnvironmentsRequest DTO.
type SetTeamMemberRoleEnvironmentsRequest struct {
	RequestBase
	MemberId     string   `json:"memberId,omitempty"`
	RoleId       string   `json:"roleId,omitempty"`
	Environments []string `json:"environments,omitempty"`
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

// GetAccountTeamInvitations DTO.
type GetAccountTeamInvitations struct {
	RequestBase
}

// CancelTeamInvitation DTO.
type CancelTeamInvitation struct {
	RequestBase
	InvitationId string `json:"invitationId,omitempty"`
}

// UpdateTeamMember DTO.
type UpdateTeamMember struct {
	RequestBase
	MemberId    string                 `json:"memberId,omitempty"`
	DisplayName string                 `json:"displayName,omitempty"`
	Roles       []*TeamMemberRoleInput `json:"roles,omitempty"`
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

// GetMyAccountUserProfile DTO.
type GetMyAccountUserProfile struct {
	RequestBase
}

// UpdateMyAccountUserPhone DTO.
type UpdateMyAccountUserPhone struct {
	RequestBase
	Phone string `json:"phone,omitempty"`
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

// GetLicenseKey DTO.
type GetLicenseKey struct {
	RequestBase
	Id string `json:"id,omitempty"`
}

// CreateAiServiceUserRequest DTO.
type CreateAiServiceUserRequest struct {
	RequestBase
	Name             string                 `json:"name,omitempty"`
	AccountRoleIds   []string               `json:"accountRoleIds,omitempty"`
	ProjectRoles     []*AiProjectRoleIdsDto `json:"projectRoles,omitempty"`
	RoleEnvironments []*RoleEnvironmentsDto `json:"roleEnvironments,omitempty"`
}

// ListAiServiceUsersRequest DTO.
type ListAiServiceUsersRequest struct {
	RequestBase
}

// RotateAiServiceUserKeyRequest DTO.
type RotateAiServiceUserKeyRequest struct {
	RequestBase
	Id          string `json:"id,omitempty"`
	RevokeKeyId string `json:"revokeKeyId,omitempty"`
}

// RevokeAiServiceUserKeyRequest DTO.
type RevokeAiServiceUserKeyRequest struct {
	RequestBase
	Id    string `json:"id,omitempty"`
	KeyId string `json:"keyId,omitempty"`
}

// UpdateAiServiceUserRequest DTO.
type UpdateAiServiceUserRequest struct {
	RequestBase
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// DeleteAiServiceUserRequest DTO.
type DeleteAiServiceUserRequest struct {
	RequestBase
	Id string `json:"id,omitempty"`
}

// DisableMembership DTO.
type DisableMembership struct {
	CodeMashRequestBase
}

// EnableMembership DTO.
type EnableMembership struct {
	CodeMashRequestBase
}

// ImportUsersRequest DTO.
type ImportUsersRequest struct {
	CodeMashRequestBase
	DryRun             *bool               `json:"dryRun,omitempty"`
	IgnoreDefaultRoles bool                `json:"ignoreDefaultRoles"`
	Users              []*ImportUserRowDto `json:"users,omitempty"`
}

// RepairLoginProjectIdsRequest DTO.
type RepairLoginProjectIdsRequest struct {
	CodeMashRequestBase
	DryRun *bool `json:"dryRun,omitempty"`
}

// IssueServiceUserApiKeyRequest DTO.
type IssueServiceUserApiKeyRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Scopes                []string `json:"scopes,omitempty"`
	ExpiresInDays         *float64 `json:"expiresInDays,omitempty"`
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
	Enabled                       bool    `json:"enabled"`
	CodeTtlMinutes                float64 `json:"codeTtlMinutes,omitempty"`
	MaxCredentialsPerUser         float64 `json:"maxCredentialsPerUser,omitempty"`
	RecoveryCodeCount             float64 `json:"recoveryCodeCount,omitempty"`
	GenerateRecoveryCodesAtSignup *bool   `json:"generateRecoveryCodesAtSignup,omitempty"`
	AuthenticatorAttachment       string  `json:"authenticatorAttachment,omitempty"`
	AllowMagicLinkRecovery        bool    `json:"allowMagicLinkRecovery"`
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
	AllowGuestUsers             *bool    `json:"allowGuestUsers,omitempty"`
	GuestCleanupPeriodDays      *float64 `json:"guestCleanupPeriodDays,omitempty"`
	UserRegistersAsRole         string   `json:"userRegistersAsRole,omitempty"`
	GuestRegistersAsRole        string   `json:"guestRegistersAsRole,omitempty"`
	AllowedRegisterRoles        []string `json:"allowedRegisterRoles,omitempty"`
	NeedVerification            bool     `json:"needVerification"`
	VerificationEmailTemplate   string   `json:"verificationEmailTemplate,omitempty"`
	DeactivationEmailTemplate   string   `json:"deactivationEmailTemplate,omitempty"`
	AllowInviteUsers            bool     `json:"allowInviteUsers"`
	AllowDeactivateUsers        bool     `json:"allowDeactivateUsers"`
	InviteUserEmailTemplate     string   `json:"inviteUserEmailTemplate,omitempty"`
	InvitationExpiration        *float64 `json:"invitationExpiration,omitempty"`
	EmailVerificationExpiration *float64 `json:"emailVerificationExpiration,omitempty"`
	DeactivationExpiration      *float64 `json:"deactivationExpiration,omitempty"`
	DefaultSubscribeToNews      bool     `json:"defaultSubscribeToNews"`
	AllowEmailRegistration      *bool    `json:"allowEmailRegistration,omitempty"`
	InviteOnly                  *bool    `json:"inviteOnly,omitempty"`
	AllowAnyoneToSignUp         *bool    `json:"allowAnyoneToSignUp,omitempty"`
	MinLength                   float64  `json:"minLength,omitempty"`
	MaxLength                   *float64 `json:"maxLength,omitempty"`
	MinNumbers                  *float64 `json:"minNumbers,omitempty"`
	MaxNumbers                  *float64 `json:"maxNumbers,omitempty"`
	MinUpper                    *float64 `json:"minUpper,omitempty"`
	MaxUpper                    *float64 `json:"maxUpper,omitempty"`
	MinLower                    *float64 `json:"minLower,omitempty"`
	MaxLower                    *float64 `json:"maxLower,omitempty"`
	MinSpecial                  *float64 `json:"minSpecial,omitempty"`
	MaxSpecial                  *float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial              string   `json:"allowedSpecial,omitempty"`
}

// UpdatePasswordComplexity DTO.
type UpdatePasswordComplexity struct {
	CodeMashRequestBase
	MinLength      float64  `json:"minLength,omitempty"`
	MaxLength      *float64 `json:"maxLength,omitempty"`
	MinNumbers     *float64 `json:"minNumbers,omitempty"`
	MaxNumbers     *float64 `json:"maxNumbers,omitempty"`
	MinUpper       *float64 `json:"minUpper,omitempty"`
	MaxUpper       *float64 `json:"maxUpper,omitempty"`
	MinLower       *float64 `json:"minLower,omitempty"`
	MaxLower       *float64 `json:"maxLower,omitempty"`
	MinSpecial     *float64 `json:"minSpecial,omitempty"`
	MaxSpecial     *float64 `json:"maxSpecial,omitempty"`
	AllowedSpecial string   `json:"allowedSpecial,omitempty"`
}

// GetAuthenticationSettings DTO.
type GetAuthenticationSettings struct {
	CodeMashRequestBase
}

// UpdateAuthenticationSettings DTO.
type UpdateAuthenticationSettings struct {
	CodeMashRequestBase
	LogoutUrl      string                        `json:"logoutUrl,omitempty"`
	AllowUsernames bool                          `json:"allowUsernames"`
	Modes          []*CredentialsSettingsModeDto `json:"modes,omitempty"`
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
	AllLanguages          bool   `json:"allLanguages"`
	IncludeTerms          bool   `json:"includeTerms"`
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
	AllLanguages          bool   `json:"allLanguages"`
	TaxonomyId            string `json:"taxonomyId,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetDatabaseMergedTermTreeRequest DTO.
type GetDatabaseMergedTermTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool   `json:"allLanguages"`
	TaxonomyName          string `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetDatabaseTaxonomyTermTreeRequest DTO.
type GetDatabaseTaxonomyTermTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool     `json:"allLanguages"`
	TaxonomyName          string   `json:"taxonomyName,omitempty"`
	RootTermId            string   `json:"rootTermId,omitempty"`
	Depth                 *float64 `json:"depth,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
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
	Translatable bool   `json:"translatable"`
	Publish      *bool  `json:"publish,omitempty"`
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

// EndUserOwnRecordsReportRequest DTO.
type EndUserOwnRecordsReportRequest struct {
	CodeMashRequestBase
	DryRun *bool    `json:"dryRun,omitempty"`
	Enable []string `json:"enable,omitempty"`
}

// GetDatabaseSchema DTO.
type GetDatabaseSchema struct {
	CodeMashRequestBase
	Id      string   `json:"id,omitempty"`
	Version *float64 `json:"version,omitempty"`
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

// GetDatabaseSchemaIndexStatus DTO.
type GetDatabaseSchemaIndexStatus struct {
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
	Confirmed bool   `json:"confirmed"`
}

// RenameDatabaseSchemaRequest DTO.
type RenameDatabaseSchemaRequest struct {
	CodeMashRequestBase
	Id    string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
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

// UpdateDatabaseSchemaEmbedRequest DTO.
type UpdateDatabaseSchemaEmbedRequest struct {
	CodeMashRequestBase
	Id    string                  `json:"id,omitempty"`
	Embed *SchemaEmbedSettingsDto `json:"embed,omitempty"`
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
	CollectionName        string   `json:"collectionName,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	SchemaVersion         *float64 `json:"schemaVersion,omitempty"`
}

// DeleteManyRecords DTO.
type DeleteManyRecords struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
	AllRecords            *bool  `json:"allRecords,omitempty"`
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
	CollectionName        string   `json:"collectionName,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Field                 string   `json:"field,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	SchemaVersion         *float64 `json:"schemaVersion,omitempty"`
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
	SchemaVersion         *float64    `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	SortBy                string      `json:"sortBy,omitempty"`
	SortOrder             *float64    `json:"sortOrder,omitempty"`
	ExpandReferences      bool        `json:"expandReferences"`
	AllLanguages          bool        `json:"allLanguages"`
}

// FindOneRecord DTO.
type FindOneRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	ExpandReferences      bool   `json:"expandReferences"`
	AllLanguages          bool   `json:"allLanguages"`
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
	AllRecords            *bool  `json:"allRecords,omitempty"`
	Update                string `json:"update,omitempty"`
	ArrayFilters          string `json:"arrayFilters,omitempty"`
}

// UpdateOneRecord DTO.
type UpdateOneRecord struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Update                string `json:"update,omitempty"`
	ArrayFilters          string `json:"arrayFilters,omitempty"`
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
	HasHeader             bool                      `json:"hasHeader"`
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
	HasHeader bool                `json:"hasHeader"`
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

// GetKeylessIntegrationSetup DTO.
type GetKeylessIntegrationSetup struct {
	CodeMashRequestBase
	Module        string `json:"module,omitempty"`
	Cloud         string `json:"cloud,omitempty"`
	IntegrationId string `json:"integrationId,omitempty"`
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

// GetFileById DTO.
type GetFileById struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Id                 string `json:"id,omitempty"`
}

// GetFolderFiles DTO.
type GetFolderFiles struct {
	CodeMashListPaginationRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// UseFileRequest DTO.
type UseFileRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
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
	ShowArchived *bool  `json:"showArchived,omitempty"`
	TemplateId   string `json:"templateId,omitempty"`
}

// GetMjml DTO.
type GetMjml struct {
	CodeMashRequestBase
	Code         string             `json:"code,omitempty"`
	Tokens       []*TokenMappingDto `json:"tokens,omitempty"`
	IsForPreview bool               `json:"isForPreview"`
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

// GetEmailPreferencesByLinkRequest DTO.
type GetEmailPreferencesByLinkRequest struct {
	RequestBase
	Token string `json:"token,omitempty"`
}

// CreateEmailCampaignRequest DTO.
type CreateEmailCampaignRequest struct {
	CodeMashRequestBase
	Campaign              *EmailCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string                `json:"databaseIntegrationId,omitempty"`
	SendNow               *bool                 `json:"sendNow,omitempty"`
}

// DeleteEmailCampaignRequest DTO.
type DeleteEmailCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
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
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	CampaignId            string   `json:"campaignId,omitempty"`
	EmailAddress          string   `json:"emailAddress,omitempty"`
	TemplateId            string   `json:"templateId,omitempty"`
	From                  *float64 `json:"from,omitempty"`
	To                    *float64 `json:"to,omitempty"`
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
	Hash           string `json:"hash,omitempty"`
	ProjectId      string `json:"projectId,omitempty"`
	NotificationId string `json:"notificationId,omitempty"`
}

// RetryFailedEmailNotificationsRequest DTO.
type RetryFailedEmailNotificationsRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	IdempotencyKey        string   `json:"idempotencyKey,omitempty"`
	NotificationIds       []string `json:"notificationIds,omitempty"`
	IncludeAmbiguous      *bool    `json:"includeAmbiguous,omitempty"`
	Categories            []string `json:"categories,omitempty"`
}

// StopEmailCampaignRequest DTO.
type StopEmailCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetEmailCampaignMessagesRequest DTO.
type GetEmailCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
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
	ShowArchived *bool  `json:"showArchived,omitempty"`
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
	IsForPreview bool               `json:"isForPreview"`
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
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	CampaignId            string   `json:"campaignId,omitempty"`
	TemplateId            string   `json:"templateId,omitempty"`
	From                  *float64 `json:"from,omitempty"`
	To                    *float64 `json:"to,omitempty"`
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
	Hash           string `json:"hash,omitempty"`
	ProjectId      string `json:"projectId,omitempty"`
	NotificationId string `json:"notificationId,omitempty"`
}

// RetryFailedSmsNotificationsRequest DTO.
type RetryFailedSmsNotificationsRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	IdempotencyKey        string   `json:"idempotencyKey,omitempty"`
	NotificationIds       []string `json:"notificationIds,omitempty"`
	IncludeAmbiguous      *bool    `json:"includeAmbiguous,omitempty"`
	Categories            []string `json:"categories,omitempty"`
}

// StopSmsCampaignRequest DTO.
type StopSmsCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetSmsCampaignMessagesRequest DTO.
type GetSmsCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// SaveCodeWorkflow DTO.
type SaveCodeWorkflow struct {
	CodeMashRequestBase
	Definition *WorkflowDefinitionDto `json:"definition,omitempty"`
}

// PublishCodeWorkflow DTO.
type PublishCodeWorkflow struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetCodeWorkflowDraft DTO.
type GetCodeWorkflowDraft struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DiscardCodeWorkflowDraft DTO.
type DiscardCodeWorkflowDraft struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetCodeWorkflowVersions DTO.
type GetCodeWorkflowVersions struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetCodeWorkflowVersionDiff DTO.
type GetCodeWorkflowVersionDiff struct {
	CodeMashRequestBase
	Id          string  `json:"id,omitempty"`
	FromVersion float64 `json:"fromVersion,omitempty"`
	ToVersion   float64 `json:"toVersion,omitempty"`
}

// ValidateCodeWorkflow DTO.
type ValidateCodeWorkflow struct {
	CodeMashRequestBase
	Definition *WorkflowDefinitionDto `json:"definition,omitempty"`
}

// GetCodeWorkflows DTO.
type GetCodeWorkflows struct {
	CodeMashListPaginationRequestBase
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// GetCodeWorkflow DTO.
type GetCodeWorkflow struct {
	CodeMashRequestBase
	Id            string   `json:"id,omitempty"`
	PinnedVersion *float64 `json:"pinnedVersion,omitempty"`
}

// DeleteCodeWorkflow DTO.
type DeleteCodeWorkflow struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// EnableCodeWorkflow DTO.
type EnableCodeWorkflow struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DisableCodeWorkflow DTO.
type DisableCodeWorkflow struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetCodeWorkflowSchema DTO.
type GetCodeWorkflowSchema struct {
	CodeMashRequestBase
}

// EnableCodeWorkflows DTO.
type EnableCodeWorkflows struct {
	CodeMashRequestBase
}

// DisableCodeWorkflows DTO.
type DisableCodeWorkflows struct {
	CodeMashRequestBase
}

// InvokeCodeWorkflow DTO.
type InvokeCodeWorkflow struct {
	CodeMashRequestBase
	Id          string         `json:"id,omitempty"`
	Payload     map[string]any `json:"payload,omitempty"`
	WaitSeconds *float64       `json:"waitSeconds,omitempty"`
	RunVersion  string         `json:"runVersion,omitempty"`
}

// GetCodeWorkflowRuns DTO.
type GetCodeWorkflowRuns struct {
	CodeMashListPaginationRequestBase
	Id     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
}

// GetCodeWorkflowRun DTO.
type GetCodeWorkflowRun struct {
	CodeMashRequestBase
	RunId string `json:"runId,omitempty"`
}

// GetCodeWorkflowRunLogs DTO.
type GetCodeWorkflowRunLogs struct {
	CodeMashListPaginationRequestBase
	RunId  string `json:"runId,omitempty"`
	StepId string `json:"stepId,omitempty"`
	Level  string `json:"level,omitempty"`
}

// ResumeCodeWorkflowRunByKey DTO.
type ResumeCodeWorkflowRunByKey struct {
	CodeMashRequestBase
	Id       string         `json:"id,omitempty"`
	Key      string         `json:"key,omitempty"`
	StepId   string         `json:"stepId,omitempty"`
	Decision string         `json:"decision,omitempty"`
	Payload  map[string]any `json:"payload,omitempty"`
}

// RetryCodeWorkflowRunStep DTO.
type RetryCodeWorkflowRunStep struct {
	CodeMashRequestBase
	RunId  string `json:"runId,omitempty"`
	StepId string `json:"stepId,omitempty"`
}

// CancelCodeWorkflowRun DTO.
type CancelCodeWorkflowRun struct {
	CodeMashRequestBase
	RunId  string `json:"runId,omitempty"`
	Reason string `json:"reason,omitempty"`
}

// GetMarketplaceIntegrationAccess DTO.
type GetMarketplaceIntegrationAccess struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// SetMarketplaceIntegrationRoleAccess DTO.
type SetMarketplaceIntegrationRoleAccess struct {
	CodeMashRequestBase
	IntegrationViewId string                 `json:"integrationViewId,omitempty"`
	RoleId            string                 `json:"roleId,omitempty"`
	Package           MarketplaceAccessLevel `json:"package,omitempty"`
	Manage            bool                   `json:"manage"`
}

// SaveMarketplaceIntegrationCustomAccess DTO.
type SaveMarketplaceIntegrationCustomAccess struct {
	CodeMashRequestBase
	IntegrationViewId        string                 `json:"integrationViewId,omitempty"`
	RoleId                   string                 `json:"roleId,omitempty"`
	BasePackage              MarketplaceAccessLevel `json:"basePackage,omitempty"`
	Tools                    []string               `json:"tools,omitempty"`
	Functions                []string               `json:"functions,omitempty"`
	Manage                   bool                   `json:"manage"`
	NewPolicyForThisRoleOnly bool                   `json:"newPolicyForThisRoleOnly"`
}

// GetMarketplaceIntegrationTools DTO.
type GetMarketplaceIntegrationTools struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// StartMarketplaceIntegrationSignIn DTO.
type StartMarketplaceIntegrationSignIn struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// SignOutMarketplaceIntegration DTO.
type SignOutMarketplaceIntegration struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
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
	IntegrationViewId string             `json:"integrationViewId,omitempty"`
	TokenMappings     []*TokenMappingDto `json:"tokenMappings,omitempty"`
}

// GetMarketplaceFunctionCatalog DTO.
type GetMarketplaceFunctionCatalog struct {
	CodeMashRequestBase
	IntegrationViewId string `json:"integrationViewId,omitempty"`
}

// GetMarketplaceFunctionRuns DTO.
type GetMarketplaceFunctionRuns struct {
	CodeMashRequestBase
	FunctionId    string   `json:"functionId,omitempty"`
	TriggerId     string   `json:"triggerId,omitempty"`
	TaskId        string   `json:"taskId,omitempty"`
	IntegrationId string   `json:"integrationId,omitempty"`
	Limit         *float64 `json:"limit,omitempty"`
}

// GetCodeFunctions DTO.
type GetCodeFunctions struct {
	CodeMashListPaginationRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
	Prefix        string `json:"prefix,omitempty"`
	Tag           string `json:"tag,omitempty"`
}

// GetCodeFunctionEventSample DTO.
type GetCodeFunctionEventSample struct {
	CodeMashRequestBase
	Source   CodeFunctionSource `json:"source,omitempty"`
	Event    string             `json:"event,omitempty"`
	SchemaId string             `json:"schemaId,omitempty"`
}

// GetCodeFunctionQualifiers DTO.
type GetCodeFunctionQualifiers struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
	FunctionId    string `json:"functionId,omitempty"`
}

// GetCodeFunctionRuns DTO.
type GetCodeFunctionRuns struct {
	CodeMashRequestBase
	TriggerId string   `json:"triggerId,omitempty"`
	TaskId    string   `json:"taskId,omitempty"`
	Limit     *float64 `json:"limit,omitempty"`
}

// InvokeCodeFunction DTO.
type InvokeCodeFunction struct {
	CodeMashRequestBase
	FunctionRef    *FunctionRefDto    `json:"functionRef,omitempty"`
	Payload        map[string]any     `json:"payload,omitempty"`
	InvocationType CodeInvocationType `json:"invocationType,omitempty"`
	RunAsUserId    string             `json:"runAsUserId,omitempty"`
	Config         string             `json:"config,omitempty"`
}

// GetCodeWorkflowWaitingRuns DTO.
type GetCodeWorkflowWaitingRuns struct {
	CodeMashListPaginationRequestBase
	Status      string `json:"status,omitempty"`
	WaitingKind string `json:"waitingKind,omitempty"`
	Assignee    string `json:"assignee,omitempty"`
	WorkflowId  string `json:"workflowId,omitempty"`
	Key         string `json:"key,omitempty"`
}

// ResumeCodeWorkflowRun DTO.
type ResumeCodeWorkflowRun struct {
	CodeMashRequestBase
	RunId    string         `json:"runId,omitempty"`
	StepId   string         `json:"stepId,omitempty"`
	Decision string         `json:"decision,omitempty"`
	Payload  map[string]any `json:"payload,omitempty"`
}

// OpenCodeWorkflowResumeLink DTO.
type OpenCodeWorkflowResumeLink struct {
	RequestBase
	Token string `json:"token,omitempty"`
}

// UseCodeWorkflowResumeLink DTO.
type UseCodeWorkflowResumeLink struct {
	RequestBase
	Token string `json:"token,omitempty"`
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
	ShowArchived *bool  `json:"showArchived,omitempty"`
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
	IsForPreview bool               `json:"isForPreview"`
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

// RegisterDevice DTO.
type RegisterDevice struct {
	RequestBase
	PushDeviceDto         *PushDeviceDto `json:"pushDeviceDto,omitempty"`
	UserId                string         `json:"userId,omitempty"`
	ProjectId             string         `json:"projectId,omitempty"`
	AccountId             string         `json:"accountId,omitempty"`
	DatabaseIntegrationId string         `json:"databaseIntegrationId,omitempty"`
}

// DeletePushDevice DTO.
type DeletePushDevice struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushDevice DTO.
type GetPushDevice struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushDevices DTO.
type GetPushDevices struct {
	CodeMashListPaginationRequestBase
	UserId                string `json:"userId,omitempty"`
	DeviceKey             string `json:"deviceKey,omitempty"`
	Platform              string `json:"platform,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushCampaignAudienceCountRequest DTO.
type GetPushCampaignAudienceCountRequest struct {
	CodeMashRequestBase
	Campaign              *PushCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string               `json:"databaseIntegrationId,omitempty"`
}

// CreatePushCampaignRequest DTO.
type CreatePushCampaignRequest struct {
	CodeMashRequestBase
	Campaign              *PushCampaignRequest `json:"campaign,omitempty"`
	DatabaseIntegrationId string               `json:"databaseIntegrationId,omitempty"`
	SendNow               *bool                `json:"sendNow,omitempty"`
}

// DeletePushCampaignRequest DTO.
type DeletePushCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
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
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	TemplateId            string   `json:"templateId,omitempty"`
	From                  *float64 `json:"from,omitempty"`
	To                    *float64 `json:"to,omitempty"`
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
	Hash           string `json:"hash,omitempty"`
	ProjectId      string `json:"projectId,omitempty"`
	NotificationId string `json:"notificationId,omitempty"`
}

// RetryFailedPushNotificationsRequest DTO.
type RetryFailedPushNotificationsRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	IdempotencyKey        string   `json:"idempotencyKey,omitempty"`
	NotificationIds       []string `json:"notificationIds,omitempty"`
	IncludeAmbiguous      *bool    `json:"includeAmbiguous,omitempty"`
	Categories            []string `json:"categories,omitempty"`
}

// StopPushCampaignRequest DTO.
type StopPushCampaignRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetPushCampaignMessagesRequest DTO.
type GetPushCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	CampaignId            string `json:"campaignId,omitempty"`
	CampaignBatchId       string `json:"campaignBatchId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetNotificationInboxRequest DTO.
type GetNotificationInboxRequest struct {
	CodeMashListPaginationRequestBase
	UnreadOnly bool   `json:"unreadOnly"`
	Channel    string `json:"channel,omitempty"`
}

// GetUnreadCountRequest DTO.
type GetUnreadCountRequest struct {
	CodeMashRequestBase
}

// MarkNotificationReadRequest DTO.
type MarkNotificationReadRequest struct {
	CodeMashRequestBase
	NotificationId string `json:"notificationId,omitempty"`
}

// MarkAllNotificationsReadRequest DTO.
type MarkAllNotificationsReadRequest struct {
	CodeMashRequestBase
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
	IntegrationId string   `json:"integrationId,omitempty"`
	Limit         *float64 `json:"limit,omitempty"`
}

// GetPaymentsWebhookRuns DTO.
type GetPaymentsWebhookRuns struct {
	CodeMashListPaginationRequestBase
	IntegrationId string `json:"integrationId,omitempty"`
	EventName     string `json:"eventName,omitempty"`
	TriggerId     string `json:"triggerId,omitempty"`
	Outcome       string `json:"outcome,omitempty"`
	FromUtc       string `json:"fromUtc,omitempty"`
	ToUtc         string `json:"toUtc,omitempty"`
}

// GetPaymentsWebhookRun DTO.
type GetPaymentsWebhookRun struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
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

// DisableLogging DTO.
type DisableLogging struct {
	CodeMashRequestBase
}

// EnableLogging DTO.
type EnableLogging struct {
	CodeMashRequestBase
	CreateNorbixLogging bool `json:"createNorbixLogging"`
}

// DeleteLoggingIntegrationRequest DTO.
type DeleteLoggingIntegrationRequest struct {
	CodeMashRequestBase
	Id       string `json:"id,omitempty"`
	WipeLogs bool   `json:"wipeLogs"`
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

// GetLog DTO.
type GetLog struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetLogs DTO.
type GetLogs struct {
	CodeMashListPaginationRequestBase
	Level            string `json:"level,omitempty"`
	Module           string `json:"module,omitempty"`
	LogCorrelationId string `json:"logCorrelationId,omitempty"`
	EventCode        string `json:"eventCode,omitempty"`
	Search           string `json:"search,omitempty"`
	MetaKey          string `json:"metaKey,omitempty"`
	MetaValue        string `json:"metaValue,omitempty"`
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
	SkipCloudDashboardLogs *bool `json:"skipCloudDashboardLogs,omitempty"`
	SkipHttpBodyMeta       *bool `json:"skipHttpBodyMeta,omitempty"`
	AiChatLoggingEnabled   *bool `json:"aiChatLoggingEnabled,omitempty"`
}

// GetAgentOnboardingRequest DTO.
type GetAgentOnboardingRequest struct {
	RequestBase
}

// GetAiContextRequest DTO.
type GetAiContextRequest struct {
	RequestBase
	ProjectId    string `json:"projectId,omitempty"`
	Format       string `json:"format,omitempty"`
	KnownVersion string `json:"knownVersion,omitempty"`
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
	Archived  bool   `json:"archived"`
}

// SetChatSessionPinnedRequest DTO.
type SetChatSessionPinnedRequest struct {
	RequestBase
	SessionId string `json:"sessionId,omitempty"`
	Pinned    bool   `json:"pinned"`
}

// SetChatSessionSharingRequest DTO.
type SetChatSessionSharingRequest struct {
	RequestBase
	SessionId  string `json:"sessionId,omitempty"`
	DoNotShare bool   `json:"doNotShare"`
}

// GetChatSessionsRequest DTO.
type GetChatSessionsRequest struct {
	RequestBase
	Take            *float64 `json:"take,omitempty"`
	IncludeArchived *bool    `json:"includeArchived,omitempty"`
}

// GetChatSessionEntriesRequest DTO.
type GetChatSessionEntriesRequest struct {
	RequestBase
	SessionId string   `json:"sessionId,omitempty"`
	SinceSeq  *float64 `json:"sinceSeq,omitempty"`
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

// ScaffoldProjectRequest DTO.
type ScaffoldProjectRequest struct {
	RequestBase
	Spec string `json:"spec,omitempty"`
	Mode string `json:"mode,omitempty"`
	Env  string `json:"env,omitempty"`
}

// ValidateSchemaRequest DTO.
type ValidateSchemaRequest struct {
	RequestBase
	SchemaJson string `json:"schemaJson,omitempty"`
}

// RenderTemplatePreviewRequest DTO.
type RenderTemplatePreviewRequest struct {
	RequestBase
	ProjectId  string `json:"projectId,omitempty"`
	Env        string `json:"env,omitempty"`
	Channel    string `json:"channel,omitempty"`
	TemplateId string `json:"templateId,omitempty"`
	Body       string `json:"body,omitempty"`
	Subject    string `json:"subject,omitempty"`
	SampleData string `json:"sampleData,omitempty"`
}

// McpRequest DTO.
type McpRequest struct {
	Version       string `json:"version,omitempty"`
	Toolsets      string `json:"toolsets,omitempty"`
	RequestStream string `json:"requestStream,omitempty"`
}

// GetProjectBriefRequest DTO.
type GetProjectBriefRequest struct {
	CodeMashRequestBase
	SinceSeq *float64 `json:"sinceSeq,omitempty"`
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
	Done       *bool   `json:"done,omitempty"`
}

// SaveAiProjectTrigger DTO.
type SaveAiProjectTrigger struct {
	SaveTrigger
}

// GetAiProjectTriggers DTO.
type GetAiProjectTriggers struct {
	GetTriggers
}

// GetAiProjectTrigger DTO.
type GetAiProjectTrigger struct {
	GetTrigger
}

// EnableAiProjectTrigger DTO.
type EnableAiProjectTrigger struct {
	EnableTrigger
}

// DisableAiProjectTrigger DTO.
type DisableAiProjectTrigger struct {
	DisableTrigger
}

// DeleteAiProjectTrigger DTO.
type DeleteAiProjectTrigger struct {
	DeleteTrigger
}

// AuthDeviceStartRequest DTO.
type AuthDeviceStartRequest struct {
	RequestBase
	ClientName string `json:"clientName,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	ProjectId  string `json:"projectId,omitempty"`
}

// AuthDeviceTokenRequest DTO.
type AuthDeviceTokenRequest struct {
	RequestBase
	DeviceCode string `json:"deviceCode,omitempty"`
}

// AuthDeviceConsentRequest DTO.
type AuthDeviceConsentRequest struct {
	RequestBase
	UserCode string `json:"userCode,omitempty"`
}

// AuthDeviceConsentDecisionRequest DTO.
type AuthDeviceConsentDecisionRequest struct {
	RequestBase
	UserCode       string                 `json:"userCode,omitempty"`
	Decision       string                 `json:"decision,omitempty"`
	AccountRoleIds []string               `json:"accountRoleIds,omitempty"`
	ProjectRoles   []*AiProjectRoleIdsDto `json:"projectRoles,omitempty"`
}

// OAuthProtectedResourceMetadataRequest DTO.
type OAuthProtectedResourceMetadataRequest struct {
	Path string `json:"path,omitempty"`
}

// OAuthAuthorizationServerMetadataRequest DTO.
type OAuthAuthorizationServerMetadataRequest struct {
	Path string `json:"path,omitempty"`
}

// OAuthRegisterRequest DTO.
type OAuthRegisterRequest struct {
	Version       string `json:"version,omitempty"`
	RequestStream string `json:"requestStream,omitempty"`
}

// OAuthAuthorizeRequest DTO.
type OAuthAuthorizeRequest struct {
	Version string `json:"version,omitempty"`
}

// OAuthConsentRequest DTO.
type OAuthConsentRequest struct {
	RequestBase
	ResponseType        string `json:"responseType,omitempty"`
	ClientId            string `json:"clientId,omitempty"`
	RedirectUri         string `json:"redirectUri,omitempty"`
	CodeChallenge       string `json:"codeChallenge,omitempty"`
	CodeChallengeMethod string `json:"codeChallengeMethod,omitempty"`
	State               string `json:"state,omitempty"`
	Resource            string `json:"resource,omitempty"`
	Scope               string `json:"scope,omitempty"`
}

// OAuthConsentDecisionRequest DTO.
type OAuthConsentDecisionRequest struct {
	RequestBase
	ResponseType        string                 `json:"responseType,omitempty"`
	ClientId            string                 `json:"clientId,omitempty"`
	RedirectUri         string                 `json:"redirectUri,omitempty"`
	CodeChallenge       string                 `json:"codeChallenge,omitempty"`
	CodeChallengeMethod string                 `json:"codeChallengeMethod,omitempty"`
	State               string                 `json:"state,omitempty"`
	Resource            string                 `json:"resource,omitempty"`
	Scope               string                 `json:"scope,omitempty"`
	ConsentTicket       string                 `json:"consentTicket,omitempty"`
	Decision            string                 `json:"decision,omitempty"`
	AccountRoleIds      []string               `json:"accountRoleIds,omitempty"`
	ProjectRoles        []*AiProjectRoleIdsDto `json:"projectRoles,omitempty"`
}

// OAuthTokenRequest DTO.
type OAuthTokenRequest struct {
	Version string `json:"version,omitempty"`
}

// OAuthRevokeRequest DTO.
type OAuthRevokeRequest struct {
	Version string `json:"version,omitempty"`
}

// DeleteEmbeddingIntegrationRequest DTO.
type DeleteEmbeddingIntegrationRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmbeddingIntegration DTO.
type GetEmbeddingIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetEmbeddingIntegrations DTO.
type GetEmbeddingIntegrations struct {
	CodeMashListPaginationRequestBase
}

// SaveEmbeddingIntegration DTO.
type SaveEmbeddingIntegration struct {
	CodeMashRequestBase
	Integration *EmbeddingIntegrationRequest `json:"integration,omitempty"`
}

// TestEmbeddingIntegration DTO.
type TestEmbeddingIntegration struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
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

// SetLlmIntegrationAsDefaultRequest DTO.
type SetLlmIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
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
	Version               string `json:"version,omitempty"`
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
	IsEnabled       *bool             `json:"isEnabled,omitempty"`
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
	Enabled *bool             `json:"enabled,omitempty"`
}

// GetSchedulerTaskRuns DTO.
type GetSchedulerTaskRuns struct {
	CodeMashRequestBase
	Id   string   `json:"id,omitempty"`
	Take *float64 `json:"take,omitempty"`
}

// SaveSchedulerTaskRequest DTO.
type SaveSchedulerTaskRequest struct {
	CodeMashRequestBase
	TaskId          string                `json:"taskId,omitempty"`
	Name            string                `json:"name,omitempty"`
	Description     string                `json:"description,omitempty"`
	Cron            string                `json:"cron,omitempty"`
	InitiatorUserId string                `json:"initiatorUserId,omitempty"`
	IsEnabled       bool                  `json:"isEnabled"`
	StopOnError     bool                  `json:"stopOnError"`
	EndsAtUnix      *float64              `json:"endsAtUnix,omitempty"`
	Task            *SchedulerTaskRequest `json:"task,omitempty"`
}

// ResolveResources DTO.
type ResolveResources struct {
	CodeMashRequestBase
	Refs []*ResourceRefDto `json:"refs,omitempty"`
}

// SetUserAiPlan DTO.
type SetUserAiPlan struct {
	CodeMashRequestBase
	UserId                string `json:"userId,omitempty"`
	PlanId                string `json:"planId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// CreateContactRequest DTO.
type CreateContactRequest struct {
	CodeMashRequestBase
	PrimaryEmail string   `json:"primaryEmail,omitempty"`
	PrimaryPhone string   `json:"primaryPhone,omitempty"`
	DisplayName  string   `json:"displayName,omitempty"`
	FirstName    string   `json:"firstName,omitempty"`
	LastName     string   `json:"lastName,omitempty"`
	Tags         []string `json:"tags,omitempty"`
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
	StartingAfter string   `json:"startingAfter,omitempty"`
	PageSize      *float64 `json:"pageSize,omitempty"`
	Search        string   `json:"search,omitempty"`
	Email         string   `json:"email,omitempty"`
	UserName      string   `json:"userName,omitempty"`
	Statuses      []string `json:"statuses,omitempty"`
	Types         []string `json:"types,omitempty"`
	RoleNames     []string `json:"roleNames,omitempty"`
}

// MergeContactsRequest DTO.
type MergeContactsRequest struct {
	CodeMashRequestBase
	SurvivorId string   `json:"survivorId,omitempty"`
	MergedIds  []string `json:"mergedIds,omitempty"`
}

// RepairContactRolesRequest DTO.
type RepairContactRolesRequest struct {
	CodeMashRequestBase
	DryRun *bool `json:"dryRun,omitempty"`
}

// UpdateContactRequest DTO.
type UpdateContactRequest struct {
	CodeMashRequestBase
	ContactId    string   `json:"contactId,omitempty"`
	DisplayName  string   `json:"displayName,omitempty"`
	FirstName    string   `json:"firstName,omitempty"`
	LastName     string   `json:"lastName,omitempty"`
	FullName     string   `json:"fullName,omitempty"`
	Company      string   `json:"company,omitempty"`
	Notes        string   `json:"notes,omitempty"`
	Gender       string   `json:"gender,omitempty"`
	BirthDate    *float64 `json:"birthDate,omitempty"`
	TimeZone     string   `json:"timeZone,omitempty"`
	Language     string   `json:"language,omitempty"`
	AddressLine1 string   `json:"addressLine1,omitempty"`
	AddressLine2 string   `json:"addressLine2,omitempty"`
	Country      string   `json:"country,omitempty"`
	City         string   `json:"city,omitempty"`
	State        string   `json:"state,omitempty"`
	PostalCode   string   `json:"postalCode,omitempty"`
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
	AutoForwardAdvisories bool   `json:"autoForwardAdvisories"`
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

// GetDiagnosticPacks DTO.
type GetDiagnosticPacks struct {
	CodeMashRequestBase
}

// RunDiagnosticPackRequest DTO.
type RunDiagnosticPackRequest struct {
	CodeMashRequestBase
	PackName    string   `json:"packName,omitempty"`
	PackVersion *float64 `json:"packVersion,omitempty"`
	CaseId      string   `json:"caseId,omitempty"`
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
	RememberMe        *bool             `json:"rememberMe,omitempty"`
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
