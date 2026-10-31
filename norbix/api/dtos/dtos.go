// Code generated from the Norbix API contract. DO NOT EDIT.
//
// Regenerate with the type generator described in README.md ("Regenerate the
// types"). Editing this file by hand is lost on the next run.

// Package dtos holds the data types of the Norbix API surface. A field
// that holds another generated type is a pointer. A nullable true/false or
// number field is a pointer too (*bool, *float64): nil is left out of the
// request, false and 0 are sent. A required true/false field is always sent.
// Every other field carries omitempty, because the gateway leaves absent
// values out.
package dtos

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

// EmailCampaignRecipientsSourceTypes enum.
type EmailCampaignRecipientsSourceTypes string

const (
	EmailCampaignRecipientsSourceTypesAllUsers       EmailCampaignRecipientsSourceTypes = "AllUsers"
	EmailCampaignRecipientsSourceTypesSpecifiedUsers EmailCampaignRecipientsSourceTypes = "SpecifiedUsers"
	EmailCampaignRecipientsSourceTypesAccountUsers   EmailCampaignRecipientsSourceTypes = "AccountUsers"
	EmailCampaignRecipientsSourceTypesEmail          EmailCampaignRecipientsSourceTypes = "Email"
	EmailCampaignRecipientsSourceTypesCollection     EmailCampaignRecipientsSourceTypes = "Collection"
	EmailCampaignRecipientsSourceTypesTriggerSubject EmailCampaignRecipientsSourceTypes = "TriggerSubject"
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

// SmsCampaignRecipientsSourceTypes enum.
type SmsCampaignRecipientsSourceTypes string

const (
	SmsCampaignRecipientsSourceTypesAllUsers       SmsCampaignRecipientsSourceTypes = "AllUsers"
	SmsCampaignRecipientsSourceTypesSpecifiedUsers SmsCampaignRecipientsSourceTypes = "SpecifiedUsers"
	SmsCampaignRecipientsSourceTypesAccountUsers   SmsCampaignRecipientsSourceTypes = "AccountUsers"
	SmsCampaignRecipientsSourceTypesPhoneNumbers   SmsCampaignRecipientsSourceTypes = "PhoneNumbers"
	SmsCampaignRecipientsSourceTypesCollection     SmsCampaignRecipientsSourceTypes = "Collection"
	SmsCampaignRecipientsSourceTypesTriggerSubject SmsCampaignRecipientsSourceTypes = "TriggerSubject"
)

// RespectTimeZoneSettings enum.
type RespectTimeZoneSettings string

const (
	RespectTimeZoneSettingsRespectToLastLoginZone           RespectTimeZoneSettings = "RespectToLastLoginZone"
	RespectTimeZoneSettingsRespectToRegistrationZone        RespectTimeZoneSettings = "RespectToRegistrationZone"
	RespectTimeZoneSettingsRespectToRegistrationProjectZone RespectTimeZoneSettings = "RespectToRegistrationProjectZone"
)

// CollectionEmailCampaignRecipientField enum.
type CollectionEmailCampaignRecipientField string

const (
	CollectionEmailCampaignRecipientFieldUser  CollectionEmailCampaignRecipientField = "User"
	CollectionEmailCampaignRecipientFieldEmail CollectionEmailCampaignRecipientField = "Email"
)

// PushCampaignRecipientsSourceTypes enum.
type PushCampaignRecipientsSourceTypes string

const (
	PushCampaignRecipientsSourceTypesAllUsers       PushCampaignRecipientsSourceTypes = "AllUsers"
	PushCampaignRecipientsSourceTypesSpecifiedUsers PushCampaignRecipientsSourceTypes = "SpecifiedUsers"
	PushCampaignRecipientsSourceTypesCollection     PushCampaignRecipientsSourceTypes = "Collection"
	PushCampaignRecipientsSourceTypesDevices        PushCampaignRecipientsSourceTypes = "Devices"
	PushCampaignRecipientsSourceTypesAccountUsers   PushCampaignRecipientsSourceTypes = "AccountUsers"
	PushCampaignRecipientsSourceTypesTriggerSubject PushCampaignRecipientsSourceTypes = "TriggerSubject"
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

// EmailRecipientField enum.
type EmailRecipientField string

const (
	EmailRecipientFieldTo  EmailRecipientField = "To"
	EmailRecipientFieldCc  EmailRecipientField = "Cc"
	EmailRecipientFieldBcc EmailRecipientField = "Bcc"
)

// RecipientRemovalReason enum.
type RecipientRemovalReason string

const (
	RecipientRemovalReasonOptedOut      RecipientRemovalReason = "OptedOut"
	RecipientRemovalReasonCompliance    RecipientRemovalReason = "Compliance"
	RecipientRemovalReasonUndeliverable RecipientRemovalReason = "Undeliverable"
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

// RequestBase DTO.
type RequestBase struct {
	TimeZoneId    string `json:"timeZoneId,omitempty"`
	Version       string `json:"version,omitempty"`
	CorrelationId string `json:"correlationId,omitempty"`
}

// IVersionBasedRequest DTO.
type IVersionBasedRequest struct {
}

// IHasCorrelationIdRequest DTO.
type IHasCorrelationIdRequest struct {
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

// SaveUser DTO.
type SaveUser struct {
	CodeMashRequestBase
	DatabaseIntegrationId     string              `json:"databaseIntegrationId,omitempty"`
	UserGeneralInfo           *UserGeneralInfoDto `json:"userGeneralInfo,omitempty"`
	UserId                    string              `json:"userId,omitempty"`
	IgnoreUserRegistersAsRole bool                `json:"ignoreUserRegistersAsRole"`
}

// SaveUserWithRolesBase DTO.
type SaveUserWithRolesBase struct {
	SaveUser
	Roles []string `json:"roles,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

// IPublicRegistrationRequest DTO.
type IPublicRegistrationRequest struct {
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

// IPasskeyCeremonyRequest DTO.
type IPasskeyCeremonyRequest struct {
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

// RegisterFileItem DTO.
type RegisterFileItem struct {
	Path        string   `json:"path,omitempty"`
	OwnerUserId string   `json:"ownerUserId,omitempty"`
	ContentType string   `json:"contentType,omitempty"`
	SizeBytes   *float64 `json:"sizeBytes,omitempty"`
	FileName    string   `json:"fileName,omitempty"`
}

// WriteLogEntry DTO.
type WriteLogEntry struct {
	Level         string            `json:"level,omitempty"`
	Title         string            `json:"title,omitempty"`
	Message       string            `json:"message,omitempty"`
	EventCode     string            `json:"eventCode,omitempty"`
	CorrelationId string            `json:"correlationId,omitempty"`
	Timestamp     string            `json:"timestamp,omitempty"`
	Meta          map[string]string `json:"meta,omitempty"`
}

// PaymentOrderLineInput DTO.
type PaymentOrderLineInput struct {
	Sku             string  `json:"sku,omitempty"`
	Title           string  `json:"title,omitempty"`
	Quantity        float64 `json:"quantity,omitempty"`
	UnitAmountMinor float64 `json:"unitAmountMinor,omitempty"`
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

// ImportColumnMappingDto DTO.
type ImportColumnMappingDto struct {
	CsvColumnIndex    float64 `json:"csvColumnIndex,omitempty"`
	CsvHeader         string  `json:"csvHeader,omitempty"`
	PropertyName      string  `json:"propertyName,omitempty"`
	DontImportOnError bool    `json:"dontImportOnError"`
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

// SmsCampaignDeliverySettingsDto DTO.
type SmsCampaignDeliverySettingsDto struct {
	RecipientsSourceType    SmsCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto               `json:"mappedTokens,omitempty"`
	CampaignTime            *float64                         `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings          `json:"respectTimeZoneSettings,omitempty"`
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

// EndUserChatAttachment DTO.
type EndUserChatAttachment struct {
	Id           string  `json:"id,omitempty"`
	SessionId    string  `json:"sessionId,omitempty"`
	FileName     string  `json:"fileName,omitempty"`
	ContentType  string  `json:"contentType,omitempty"`
	Kind         string  `json:"kind,omitempty"`
	Size         float64 `json:"size,omitempty"`
	Summary      string  `json:"summary,omitempty"`
	CreatedAtUtc string  `json:"createdAtUtc,omitempty"`
}

// EndUserChatMemoryNote DTO.
type EndUserChatMemoryNote struct {
	Id           string `json:"id,omitempty"`
	SessionId    string `json:"sessionId,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Text         string `json:"text,omitempty"`
	CreatedAtUtc string `json:"createdAtUtc,omitempty"`
}

// EndUserChatAssistant DTO.
type EndUserChatAssistant struct {
	Id             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	WelcomeMessage string `json:"welcomeMessage,omitempty"`
	IsDefault      bool   `json:"isDefault"`
	MemoryEnabled  bool   `json:"memoryEnabled"`
}

// EndUserChatPlan DTO.
type EndUserChatPlan struct {
	Id           string  `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	QuotaUnit    string  `json:"quotaUnit,omitempty"`
	MonthlyQuota float64 `json:"monthlyQuota,omitempty"`
	Used         float64 `json:"used,omitempty"`
	Remaining    float64 `json:"remaining,omitempty"`
	Attachments  bool    `json:"attachments"`
	Rag          bool    `json:"rag"`
	Memory       bool    `json:"memory"`
}

// EndUserChatSession DTO.
type EndUserChatSession struct {
	Id           string  `json:"id,omitempty"`
	AssistantId  string  `json:"assistantId,omitempty"`
	Title        string  `json:"title,omitempty"`
	IsPinned     bool    `json:"isPinned"`
	IsArchived   bool    `json:"isArchived"`
	LastSeq      float64 `json:"lastSeq,omitempty"`
	CreatedAtUtc string  `json:"createdAtUtc,omitempty"`
	UpdatedAtUtc string  `json:"updatedAtUtc,omitempty"`
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

// EndUserAiToolParameter DTO.
type EndUserAiToolParameter struct {
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Required    bool   `json:"required"`
	Description string `json:"description,omitempty"`
}

// EndUserAiTool DTO.
type EndUserAiTool struct {
	Name                 string                    `json:"name,omitempty"`
	Description          string                    `json:"description,omitempty"`
	Toolsets             []string                  `json:"toolsets,omitempty"`
	RequiresConfirmation bool                      `json:"requiresConfirmation"`
	Parameters           []*EndUserAiToolParameter `json:"parameters,omitempty"`
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

// PaginatedResponse DTO.
type PaginatedResponse[TViewModelProjection any] struct {
	Items         []TViewModelProjection `json:"items,omitempty"`
	HasMore       bool                   `json:"hasMore"`
	HasPrevious   bool                   `json:"hasPrevious"`
	StartingAfter string                 `json:"startingAfter,omitempty"`
	EndingBefore  string                 `json:"endingBefore,omitempty"`
}

// UserMarketingPreferencesDto DTO.
type UserMarketingPreferencesDto struct {
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
}

// PasskeyListItemDto DTO.
type PasskeyListItemDto struct {
	CredentialId    string `json:"credentialId,omitempty"`
	FriendlyName    string `json:"friendlyName,omitempty"`
	RegisteredOnUtc string `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   string `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool   `json:"isRevoked"`
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

// DataSchemaDto DTO.
type DataSchemaDto struct {
	Json   string                `json:"json,omitempty"`
	Fields []*JsonSchemaFieldDto `json:"fields,omitempty"`
}

// VisualSchemaDto DTO.
type VisualSchemaDto struct {
	Json string `json:"json,omitempty"`
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

// TriggerActionDto DTO.
type TriggerActionDto struct {
	Type          TriggerActionType `json:"type,omitempty"`
	IntegrationId string            `json:"integrationId,omitempty"`
}

// ExecutionIdentitySummaryDto DTO.
type ExecutionIdentitySummaryDto struct {
	Id        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	IsDefault bool   `json:"isDefault"`
	Status    string `json:"status,omitempty"`
}

// TriggerDto DTO.
type TriggerDto struct {
	Type                TriggerType                  `json:"type,omitempty"`
	ViewId              string                       `json:"viewId,omitempty"`
	Name                string                       `json:"name,omitempty"`
	ThenAction          *TriggerActionDto            `json:"thenAction,omitempty"`
	Description         string                       `json:"description,omitempty"`
	IsEnabled           bool                         `json:"isEnabled"`
	ActivationCode      string                       `json:"activationCode,omitempty"`
	Order               *float64                     `json:"order,omitempty"`
	BreakOnError        bool                         `json:"breakOnError"`
	SavedByAuthId       string                       `json:"savedByAuthId,omitempty"`
	ExecutionIdentityId string                       `json:"executionIdentityId,omitempty"`
	ExecutionIdentity   *ExecutionIdentitySummaryDto `json:"executionIdentity,omitempty"`
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

// PublicFolderDto DTO.
type PublicFolderDto struct {
	Path      string `json:"path,omitempty"`
	PublicId  string `json:"publicId,omitempty"`
	PublicUrl string `json:"publicUrl,omitempty"`
	Inherited bool   `json:"inherited"`
}

// RegisterFileResult DTO.
type RegisterFileResult struct {
	Index       float64  `json:"index,omitempty"`
	Path        string   `json:"path,omitempty"`
	OwnerUserId string   `json:"ownerUserId,omitempty"`
	Status      string   `json:"status,omitempty"`
	FileId      string   `json:"fileId,omitempty"`
	Verified    *bool    `json:"verified,omitempty"`
	Warning     string   `json:"warning,omitempty"`
	ErrorCode   string   `json:"errorCode,omitempty"`
	ErrorStatus *float64 `json:"errorStatus,omitempty"`
	Error       string   `json:"error,omitempty"`
}

// OwnPushDeviceDto DTO.
type OwnPushDeviceDto struct {
	Id             string `json:"id,omitempty"`
	Provider       string `json:"provider,omitempty"`
	Token          string `json:"token,omitempty"`
	Platform       string `json:"platform,omitempty"`
	OsName         string `json:"osName,omitempty"`
	OsVersion      string `json:"osVersion,omitempty"`
	AppVersion     string `json:"appVersion,omitempty"`
	DeviceName     string `json:"deviceName,omitempty"`
	ModelName      string `json:"modelName,omitempty"`
	Brand          string `json:"brand,omitempty"`
	InstallationId string `json:"installationId,omitempty"`
	Locale         string `json:"locale,omitempty"`
	TimeZone       string `json:"timeZone,omitempty"`
	UpdatedAtUtc   string `json:"updatedAtUtc,omitempty"`
}

// OwnPushMessageDto DTO.
type OwnPushMessageDto struct {
	Id         string `json:"id,omitempty"`
	CampaignId string `json:"campaignId,omitempty"`
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	SentAtUtc  string `json:"sentAtUtc,omitempty"`
	IsRead     bool   `json:"isRead"`
	ReadAtUtc  string `json:"readAtUtc,omitempty"`
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

// WorkflowRunErrorDto DTO.
type WorkflowRunErrorDto struct {
	Code    string `json:"code,omitempty"`
	StepId  string `json:"stepId,omitempty"`
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

// WorkflowRunListProjection DTO.
type WorkflowRunListProjection struct {
	RunId                      string                  `json:"runId,omitempty"`
	WorkflowId                 string                  `json:"workflowId,omitempty"`
	Version                    *float64                `json:"version,omitempty"`
	IsDraft                    bool                    `json:"isDraft"`
	Status                     string                  `json:"status,omitempty"`
	Source                     string                  `json:"source,omitempty"`
	StartedAt                  string                  `json:"startedAt,omitempty"`
	FinishedAt                 string                  `json:"finishedAt,omitempty"`
	DurationMs                 *float64                `json:"durationMs,omitempty"`
	StartedBy                  *WorkflowRunActorDto    `json:"startedBy,omitempty"`
	ExecutionIdentityId        string                  `json:"executionIdentityId,omitempty"`
	ExecutionIdentityName      string                  `json:"executionIdentityName,omitempty"`
	ExecutionIdentityIsDefault bool                    `json:"executionIdentityIsDefault"`
	CurrentStepIds             []string                `json:"currentStepIds,omitempty"`
	WaitingOn                  []*WorkflowWaitingOnDto `json:"waitingOn,omitempty"`
	Error                      *WorkflowRunErrorDto    `json:"error,omitempty"`
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
	RunId                      string                  `json:"runId,omitempty"`
	WorkflowId                 string                  `json:"workflowId,omitempty"`
	Version                    *float64                `json:"version,omitempty"`
	IsDraft                    bool                    `json:"isDraft"`
	Status                     string                  `json:"status,omitempty"`
	Source                     string                  `json:"source,omitempty"`
	StartedBy                  *WorkflowRunActorDto    `json:"startedBy,omitempty"`
	TriggerId                  string                  `json:"triggerId,omitempty"`
	TaskId                     string                  `json:"taskId,omitempty"`
	ParentRunId                string                  `json:"parentRunId,omitempty"`
	Depth                      float64                 `json:"depth,omitempty"`
	Env                        string                  `json:"env,omitempty"`
	CorrelationId              string                  `json:"correlationId,omitempty"`
	ExecutionIdentityId        string                  `json:"executionIdentityId,omitempty"`
	ExecutionIdentityName      string                  `json:"executionIdentityName,omitempty"`
	ExecutionIdentityIsDefault bool                    `json:"executionIdentityIsDefault"`
	StartedAt                  string                  `json:"startedAt,omitempty"`
	FinishedAt                 string                  `json:"finishedAt,omitempty"`
	DurationMs                 *float64                `json:"durationMs,omitempty"`
	WaitingOn                  []*WorkflowWaitingOnDto `json:"waitingOn,omitempty"`
	Error                      *WorkflowRunErrorDto    `json:"error,omitempty"`
	Output                     any                     `json:"output,omitempty"`
	StepRuns                   float64                 `json:"stepRuns,omitempty"`
	Steps                      []*WorkflowRunStepDto   `json:"steps,omitempty"`
	Request                    map[string]any          `json:"request,omitempty"`
	Trigger                    map[string]any          `json:"trigger,omitempty"`
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

// CampaignStatusChangeEntryDto DTO.
type CampaignStatusChangeEntryDto struct {
	Time   string         `json:"time,omitempty"`
	Status CampaignStatus `json:"status,omitempty"`
	Errors []*ErrorDto    `json:"errors,omitempty"`
}

// CampaignOutcomeDto DTO.
type CampaignOutcomeDto struct {
	Sent              float64 `json:"sent,omitempty"`
	Failed            float64 `json:"failed,omitempty"`
	Blocked           float64 `json:"blocked,omitempty"`
	Cancelled         float64 `json:"cancelled,omitempty"`
	RemovedRecipients float64 `json:"removedRecipients,omitempty"`
	BlockedReason     string  `json:"blockedReason,omitempty"`
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
	TriggerId                       string                          `json:"triggerId,omitempty"`
	TriggerType                     TriggerType                     `json:"triggerType,omitempty"`
	Outcome                         *CampaignOutcomeDto             `json:"outcome,omitempty"`
	Id                              string                          `json:"id,omitempty"`
}

// EmailCampaignDeliverySettingsDto DTO.
type EmailCampaignDeliverySettingsDto struct {
	RecipientsSourceType    EmailCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto                 `json:"mappedTokens,omitempty"`
	CampaignTime            *float64                           `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings            `json:"respectTimeZoneSettings,omitempty"`
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
	ViewId        string              `json:"viewId,omitempty"`
	TemplateName  string              `json:"templateName,omitempty"`
	TemplateId    string              `json:"templateId,omitempty"`
	IntegrationId string              `json:"integrationId,omitempty"`
	Language      string              `json:"language,omitempty"`
	Strategy      string              `json:"strategy,omitempty"`
	LatestStatus  CampaignStatus      `json:"latestStatus,omitempty"`
	CreatedOn     string              `json:"createdOn,omitempty"`
	Outcome       *CampaignOutcomeDto `json:"outcome,omitempty"`
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

// RemovedRecipientDto DTO.
type RemovedRecipientDto struct {
	Address string                 `json:"address,omitempty"`
	Field   EmailRecipientField    `json:"field,omitempty"`
	Reason  RecipientRemovalReason `json:"reason,omitempty"`
	Tag     string                 `json:"tag,omitempty"`
}

// EmailCampaignBatchNotificationDto DTO.
type EmailCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	Subject           string                  `json:"subject,omitempty"`
	Recipients        *EmailRecipientsDto     `json:"recipients,omitempty"`
	RemovedRecipients []*RemovedRecipientDto  `json:"removedRecipients,omitempty"`
	Content           *EmailMessageContentDto `json:"content,omitempty"`
	TransactionalKind string                  `json:"transactionalKind,omitempty"`
}

// CampaignStatsDto DTO.
type CampaignStatsDto struct {
	Batches           float64 `json:"batches,omitempty"`
	Sent              float64 `json:"sent,omitempty"`
	Failed            float64 `json:"failed,omitempty"`
	FailedRetryable   float64 `json:"failedRetryable,omitempty"`
	Cancelled         float64 `json:"cancelled,omitempty"`
	Blocked           float64 `json:"blocked,omitempty"`
	InProgress        float64 `json:"inProgress,omitempty"`
	RemovedRecipients float64 `json:"removedRecipients,omitempty"`
	SuccessRate       float64 `json:"successRate,omitempty"`
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

// PushAudienceCountDto DTO.
type PushAudienceCountDto struct {
	Devices        float64 `json:"devices,omitempty"`
	Recipients     float64 `json:"recipients,omitempty"`
	SkippedUserIds float64 `json:"skippedUserIds,omitempty"`
	IsCapped       bool    `json:"isCapped"`
}

// PushCampaignDeliverySettingsDto DTO.
type PushCampaignDeliverySettingsDto struct {
	RecipientsSourceType    PushCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	MappedTokens            []*TokenMappingDto                `json:"mappedTokens,omitempty"`
	CampaignTime            *float64                          `json:"campaignTime,omitempty"`
	RespectTimeZoneSettings RespectTimeZoneSettings           `json:"respectTimeZoneSettings,omitempty"`
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

// PushCampaignDto DTO.
type PushCampaignDto struct {
	CampaignDto
	Recipients *PushCampaignDeliverySettingsDto `json:"recipients,omitempty"`
	Template   *PushTemplateDto                 `json:"template,omitempty"`
}

// PushDeviceDeliveryTokenDto DTO.
type PushDeviceDeliveryTokenDto struct {
	PushDeviceToken string                   `json:"pushDeviceToken,omitempty"`
	DeliveryFamily  PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
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

// IBindableContract DTO.
type IBindableContract struct {
}

// IHasViewId DTO.
type IHasViewId struct {
}

// IHasResponsibleUserId DTO.
type IHasResponsibleUserId struct {
}

// IHasDatabaseId DTO.
type IHasDatabaseId struct {
}

// ICursorArgs DTO.
type ICursorArgs struct {
}

// IQueuedTrigger DTO.
type IQueuedTrigger struct {
}

// IHasRazorTemplateCode DTO.
type IHasRazorTemplateCode struct {
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

// IdResponse DTO.
type IdResponse struct {
	ResponseBase
	Id       string   `json:"id,omitempty"`
	Status   string   `json:"status,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

// ListEndUserChatAttachmentsResponse DTO.
type ListEndUserChatAttachmentsResponse struct {
	ResponseBase
	Attachments []*EndUserChatAttachment `json:"attachments,omitempty"`
}

// EmptyResponse DTO.
type EmptyResponse struct {
	ResponseBase
}

// ListEndUserChatMemoryResponse DTO.
type ListEndUserChatMemoryResponse struct {
	ResponseBase
	Notes []*EndUserChatMemoryNote `json:"notes,omitempty"`
}

// GetEndUserChatAvailabilityResponse DTO.
type GetEndUserChatAvailabilityResponse struct {
	ResponseBase
	Enabled            bool                    `json:"enabled"`
	Available          bool                    `json:"available"`
	Reason             string                  `json:"reason,omitempty"`
	DefaultAssistantId string                  `json:"defaultAssistantId,omitempty"`
	Assistants         []*EndUserChatAssistant `json:"assistants,omitempty"`
	Plan               *EndUserChatPlan        `json:"plan,omitempty"`
}

// ListEndUserChatSessionsResponse DTO.
type ListEndUserChatSessionsResponse struct {
	ResponseBase
	Sessions []*EndUserChatSession `json:"sessions,omitempty"`
}

// GetEndUserChatSessionResponse DTO.
type GetEndUserChatSessionResponse struct {
	ResponseBase
	Session *EndUserChatSession `json:"session,omitempty"`
}

// GetEndUserChatEntriesResponse DTO.
type GetEndUserChatEntriesResponse struct {
	ResponseBase
	SessionId string                `json:"sessionId,omitempty"`
	Entries   []*AiChatEntryWireDto `json:"entries,omitempty"`
	LastSeq   float64               `json:"lastSeq,omitempty"`
	HasMore   bool                  `json:"hasMore"`
}

// StartEndUserChatTurnResponse DTO.
type StartEndUserChatTurnResponse struct {
	ResponseBase
	TurnId    string `json:"turnId,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
	Channel   string `json:"channel,omitempty"`
}

// GetEndUserAiToolsResponse DTO.
type GetEndUserAiToolsResponse struct {
	ResponseBase
	Tools []*EndUserAiTool `json:"tools,omitempty"`
}

// InvokeEndUserAiToolResponse DTO.
type InvokeEndUserAiToolResponse struct {
	ResponseBase
	Result string `json:"result,omitempty"`
}

// SaveGuestUserResponse DTO.
type SaveGuestUserResponse struct {
	IdResponse
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// GetUserResponse DTO.
type GetUserResponse struct {
	ResponseBase
	User *AuthDto `json:"user,omitempty"`
}

// GetUsersResponse DTO.
type GetUsersResponse struct {
	ResponseBase
	List PaginatedResponse[*AuthDto] `json:"list,omitempty"`
}

// GetUserPreferencesResponse DTO.
type GetUserPreferencesResponse struct {
	ResponseBase
	Preferences *UserMarketingPreferencesDto `json:"preferences,omitempty"`
}

// CredentialChangedResponse DTO.
type CredentialChangedResponse struct {
	ResponseBase
	AuthId            string `json:"authId,omitempty"`
	CredentialVersion string `json:"credentialVersion,omitempty"`
}

// PasskeyOkResponse DTO.
type PasskeyOkResponse struct {
	ResponseBase
}

// SendTestPasswordResetResponse DTO.
type SendTestPasswordResetResponse struct {
	ResponseBase
	Queued             bool     `json:"queued"`
	Reason             string   `json:"reason,omitempty"`
	Env                string   `json:"env,omitempty"`
	CorrelationId      string   `json:"correlationId,omitempty"`
	ExpiresAtUtc       string   `json:"expiresAtUtc,omitempty"`
	IntegrationId      string   `json:"integrationId,omitempty"`
	IntegrationName    string   `json:"integrationName,omitempty"`
	Provider           string   `json:"provider,omitempty"`
	IntegrationProblem string   `json:"integrationProblem,omitempty"`
	TemplateId         string   `json:"templateId,omitempty"`
	TemplateName       string   `json:"templateName,omitempty"`
	TemplateSource     string   `json:"templateSource,omitempty"`
	Language           string   `json:"language,omitempty"`
	LanguageRule       string   `json:"languageRule,omitempty"`
	TemplateLanguages  []string `json:"templateLanguages,omitempty"`
	Fallback           string   `json:"fallback,omitempty"`
	MjmlRenderer       string   `json:"mjmlRenderer,omitempty"`
}

// PasskeyCeremonyOptionsResponse DTO.
type PasskeyCeremonyOptionsResponse struct {
	ResponseBase
	CeremonyId  string `json:"ceremonyId,omitempty"`
	OptionsJson string `json:"optionsJson,omitempty"`
}

// PasskeyAuthTokensResponse DTO.
type PasskeyAuthTokensResponse struct {
	ResponseBase
	AccessToken      string   `json:"accessToken,omitempty"`
	RefreshToken     string   `json:"refreshToken,omitempty"`
	ExpiresInSeconds float64  `json:"expiresInSeconds,omitempty"`
	RecoveryCodes    []string `json:"recoveryCodes,omitempty"`
}

// PasskeyListResponse DTO.
type PasskeyListResponse struct {
	ResponseBase
	Passkeys []*PasskeyListItemDto `json:"passkeys,omitempty"`
}

// PasskeyRecoveryResponse DTO.
type PasskeyRecoveryResponse struct {
	ResponseBase
	AccessToken      string  `json:"accessToken,omitempty"`
	RefreshToken     string  `json:"refreshToken,omitempty"`
	ExpiresInSeconds float64 `json:"expiresInSeconds,omitempty"`
	RemainingCodes   float64 `json:"remainingCodes,omitempty"`
}

// PasskeyVerificationTokenResponse DTO.
type PasskeyVerificationTokenResponse struct {
	ResponseBase
	VerificationToken string `json:"verificationToken,omitempty"`
}

// FindMergedTermTreeResponse DTO.
type FindMergedTermTreeResponse struct {
	ResponseBase
	Tree []*TermTreeDto `json:"tree,omitempty"`
}

// FindTaxonomyTreeResponse DTO.
type FindTaxonomyTreeResponse struct {
	ResponseBase
	Tree []*TaxonomyTreeDto `json:"tree,omitempty"`
}

// FindTermsResponse DTO.
type FindTermsResponse struct {
	ResponseBase
	List PaginatedResponse[*TermDto] `json:"list,omitempty"`
}

// FindTermsChildrenResponse DTO.
type FindTermsChildrenResponse struct {
	ResponseBase
	List PaginatedResponse[*TermDto] `json:"list,omitempty"`
}

// FindTermTreeResponse DTO.
type FindTermTreeResponse struct {
	ResponseBase
	Tree []*TermTreeDto `json:"tree,omitempty"`
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

// AggregateResponse DTO.
type AggregateResponse struct {
	ResponseBase
	Result []any `json:"result,omitempty"`
}

// CountResponse DTO.
type CountResponse struct {
	ResponseBase
	Count float64 `json:"count,omitempty"`
}

// DistinctResponse DTO.
type DistinctResponse struct {
	ResponseBase
	Values []any `json:"values,omitempty"`
}

// ExecuteAggregateResponse DTO.
type ExecuteAggregateResponse struct {
	ResponseBase
	Result []any `json:"result,omitempty"`
}

// FindResponse DTO.
type FindResponse struct {
	ResponseBase
	List PaginatedResponse[any] `json:"list,omitempty"`
}

// FindOneResponse DTO.
type FindOneResponse struct {
	ResponseBase
	Result any `json:"result,omitempty"`
}

// GetFileByIdResponse DTO.
type GetFileByIdResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl string              `json:"publicUrl,omitempty"`
}

// GetFileInfoResponse DTO.
type GetFileInfoResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl string              `json:"publicUrl,omitempty"`
}

// GetSignedUrlResponse DTO.
type GetSignedUrlResponse struct {
	ResponseBase
	Url string `json:"url,omitempty"`
}

// ListFilesResponse DTO.
type ListFilesResponse struct {
	ResponseBase
	List          PaginatedResponse[*FileResourceRefDto] `json:"list,omitempty"`
	Folders       []string                               `json:"folders,omitempty"`
	PublicFolders []*PublicFolderDto                     `json:"publicFolders,omitempty"`
}

// RegisterFilesResponse DTO.
type RegisterFilesResponse struct {
	ResponseBase
	DryRun     bool                  `json:"dryRun"`
	Registered float64               `json:"registered,omitempty"`
	Unchanged  float64               `json:"unchanged,omitempty"`
	Failed     float64               `json:"failed,omitempty"`
	Items      []*RegisterFileResult `json:"items,omitempty"`
}

// RequestUploadUrlResponse DTO.
type RequestUploadUrlResponse struct {
	ResponseBase
	Url string `json:"url,omitempty"`
}

// GetOwnPushDevicesResponse DTO.
type GetOwnPushDevicesResponse struct {
	ResponseBase
	Items []*OwnPushDeviceDto `json:"items,omitempty"`
}

// GetOwnPushMessagesResponse DTO.
type GetOwnPushMessagesResponse struct {
	ResponseBase
	List PaginatedResponse[*OwnPushMessageDto] `json:"list,omitempty"`
}

// WriteLogsResponse DTO.
type WriteLogsResponse struct {
	ResponseBase
	Accepted float64 `json:"accepted,omitempty"`
}

// GetPaymentOrderResponse DTO.
type GetPaymentOrderResponse struct {
	ResponseBase
	Item *PaymentOrderDto `json:"item,omitempty"`
}

// GetPaymentOrdersResponse DTO.
type GetPaymentOrdersResponse struct {
	ResponseBase
	List PaginatedResponse[*PaymentOrderDto] `json:"list,omitempty"`
}

// StartPaymentCheckoutResponse DTO.
type StartPaymentCheckoutResponse struct {
	ResponseBase
	Checkout *PaymentOrderCheckoutDto `json:"checkout,omitempty"`
	Reused   bool                     `json:"reused"`
	Order    *PaymentOrderDto         `json:"order,omitempty"`
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

// WorkflowRunLogsResponse DTO.
type WorkflowRunLogsResponse struct {
	ResponseBase
	List PaginatedResponse[*TenantLogEntryDto] `json:"list,omitempty"`
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

// GetFileResponse DTO.
type GetFileResponse struct {
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

// UploadEndUserChatAttachmentRequest DTO.
type UploadEndUserChatAttachmentRequest struct {
	CodeMashRequestBase
	SessionId     string `json:"sessionId,omitempty"`
	FileName      string `json:"fileName,omitempty"`
	ContentType   string `json:"contentType,omitempty"`
	Base64Content string `json:"base64Content,omitempty"`
}

// ListEndUserChatAttachmentsRequest DTO.
type ListEndUserChatAttachmentsRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// DeleteEndUserChatAttachmentRequest DTO.
type DeleteEndUserChatAttachmentRequest struct {
	CodeMashRequestBase
	AttachmentId string `json:"attachmentId,omitempty"`
}

// SetEndUserChatEntryFeedbackRequest DTO.
type SetEndUserChatEntryFeedbackRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	EntryId   string `json:"entryId,omitempty"`
	Feedback  string `json:"feedback,omitempty"`
}

// ListEndUserChatMemoryRequest DTO.
type ListEndUserChatMemoryRequest struct {
	CodeMashRequestBase
	Take *float64 `json:"take,omitempty"`
}

// ForgetEndUserChatMemoryRequest DTO.
type ForgetEndUserChatMemoryRequest struct {
	CodeMashRequestBase
	NoteId string `json:"noteId,omitempty"`
}

// GetEndUserChatAvailabilityRequest DTO.
type GetEndUserChatAvailabilityRequest struct {
	CodeMashRequestBase
}

// ListEndUserChatSessionsRequest DTO.
type ListEndUserChatSessionsRequest struct {
	CodeMashRequestBase
	Take            *float64 `json:"take,omitempty"`
	IncludeArchived bool     `json:"includeArchived"`
}

// CreateEndUserChatSessionRequest DTO.
type CreateEndUserChatSessionRequest struct {
	CodeMashRequestBase
	AssistantId string `json:"assistantId,omitempty"`
	Title       string `json:"title,omitempty"`
}

// GetEndUserChatSessionRequest DTO.
type GetEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// RenameEndUserChatSessionRequest DTO.
type RenameEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	Title     string `json:"title,omitempty"`
}

// PinEndUserChatSessionRequest DTO.
type PinEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	Pinned    bool   `json:"pinned"`
}

// ArchiveEndUserChatSessionRequest DTO.
type ArchiveEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
	Archived  bool   `json:"archived"`
}

// DeleteEndUserChatSessionRequest DTO.
type DeleteEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId,omitempty"`
}

// GetEndUserChatEntriesRequest DTO.
type GetEndUserChatEntriesRequest struct {
	CodeMashRequestBase
	SessionId string   `json:"sessionId,omitempty"`
	AfterSeq  *float64 `json:"afterSeq,omitempty"`
	Take      *float64 `json:"take,omitempty"`
}

// StartEndUserChatTurnRequest DTO.
type StartEndUserChatTurnRequest struct {
	CodeMashRequestBase
	SessionId   string `json:"sessionId,omitempty"`
	AssistantId string `json:"assistantId,omitempty"`
	Message     string `json:"message,omitempty"`
}

// GetEndUserAiToolsRequest DTO.
type GetEndUserAiToolsRequest struct {
	RequestBase
}

// InvokeEndUserAiToolRequest DTO.
type InvokeEndUserAiToolRequest struct {
	RequestBase
	ToolName      string `json:"toolName,omitempty"`
	ArgumentsJson string `json:"argumentsJson,omitempty"`
}

// BlockUserRequest DTO.
type BlockUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// SaveSystemUserWithPermissions DTO.
type SaveSystemUserWithPermissions struct {
	SaveUserWithRolesBase
}

// SaveGuestUser DTO.
type SaveGuestUser struct {
	SaveUser
}

// SavePhoneUser DTO.
type SavePhoneUser struct {
	SaveUser
	Phone string `json:"phone,omitempty"`
}

// SaveUserNameUser DTO.
type SaveUserNameUser struct {
	SaveUser
	Password string   `json:"password,omitempty"`
	UserName string   `json:"userName,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}

// SaveEmailUser DTO.
type SaveEmailUser struct {
	SaveUser
	Password string   `json:"password,omitempty"`
	Email    string   `json:"email,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}

// SavePhoneUserNameWithPermissions DTO.
type SavePhoneUserNameWithPermissions struct {
	SaveUserWithRolesBase
	Phone string `json:"phone,omitempty"`
}

// SaveEmailUserNameWithPermissions DTO.
type SaveEmailUserNameWithPermissions struct {
	SaveUserWithRolesBase
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

// SaveUserNameWithPermissions DTO.
type SaveUserNameWithPermissions struct {
	SaveUserWithRolesBase
	Password string `json:"password,omitempty"`
	UserName string `json:"userName,omitempty"`
}

// DeleteUserRequest DTO.
type DeleteUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetUserRequest DTO.
type GetUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GetUsersRequest DTO.
type GetUsersRequest struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId    string   `json:"databaseIntegrationId,omitempty"`
	IncludePermissions       bool     `json:"includePermissions"`
	UserShouldHavePushDevice bool     `json:"userShouldHavePushDevice"`
	UserShouldHaveEmail      bool     `json:"userShouldHaveEmail"`
	IncludeMeta              bool     `json:"includeMeta"`
	RoleNames                []string `json:"roleNames,omitempty"`
	UserIds                  []string `json:"userIds,omitempty"`
	Search                   string   `json:"search,omitempty"`
	Email                    string   `json:"email,omitempty"`
	UserName                 string   `json:"userName,omitempty"`
	Statuses                 []string `json:"statuses,omitempty"`
	Types                    []string `json:"types,omitempty"`
}

// GetUserPreferencesRequest DTO.
type GetUserPreferencesRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// GrantContactConsentRequest DTO.
type GrantContactConsentRequest struct {
	CodeMashRequestBase
	ContactId   string `json:"contactId,omitempty"`
	Channel     string `json:"channel,omitempty"`
	LawfulBasis string `json:"lawfulBasis,omitempty"`
	Source      string `json:"source,omitempty"`
	EvidenceRef string `json:"evidenceRef,omitempty"`
}

// InviteUserRequest DTO.
type InviteUserRequest struct {
	CodeMashRequestBase
	Email                 string `json:"email,omitempty"`
	FullName              string `json:"fullName,omitempty"`
	Salutation            string `json:"salutation,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FinishInvitationRequest DTO.
type FinishInvitationRequest struct {
	CodeMashRequestBase
	Token       string `json:"token,omitempty"`
	Password    string `json:"password,omitempty"`
	FullName    string `json:"fullName,omitempty"`
	Salutation  string `json:"salutation,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
}

// LinkIdentityRequest DTO.
type LinkIdentityRequest struct {
	CodeMashRequestBase
	UserId                string `json:"userId,omitempty"`
	Provider              string `json:"provider,omitempty"`
	ProviderToken         string `json:"providerToken,omitempty"`
	EmailToVerify         string `json:"emailToVerify,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// MapAuthToUserRequest DTO.
type MapAuthToUserRequest struct {
	CodeMashRequestBase
	UserId                string `json:"userId,omitempty"`
	AuthId                string `json:"authId,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// AssignRolePermissionsRequest DTO.
type AssignRolePermissionsRequest struct {
	CodeMashRequestBase
	Id                    string   `json:"id,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Roles                 []string `json:"roles,omitempty"`
}

// SetContactRolesRequest DTO.
type SetContactRolesRequest struct {
	CodeMashRequestBase
	UserId                string   `json:"userId,omitempty"`
	Roles                 []string `json:"roles,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
}

// SetContactTagSubscriptionRequest DTO.
type SetContactTagSubscriptionRequest struct {
	CodeMashRequestBase
	ContactId   string `json:"contactId,omitempty"`
	CommChannel string `json:"commChannel,omitempty"`
	Channel     string `json:"channel,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Subscribed  bool   `json:"subscribed"`
}

// UnblockUserRequest DTO.
type UnblockUserRequest struct {
	CodeMashRequestBase
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// UnsubscribeContactRequest DTO.
type UnsubscribeContactRequest struct {
	CodeMashRequestBase
	ContactId string `json:"contactId,omitempty"`
	Channel   string `json:"channel,omitempty"`
	Reason    string `json:"reason,omitempty"`
}

// UpdateUserRequest DTO.
type UpdateUserRequest struct {
	SaveUser
	Id string `json:"id,omitempty"`
}

// UpdateUserPreferencesRequest DTO.
type UpdateUserPreferencesRequest struct {
	CodeMashRequestBase
	Id                        string              `json:"id,omitempty"`
	BlockAllMarketingMessages bool                `json:"blockAllMarketingMessages"`
	BlockedTags               map[string][]string `json:"blockedTags,omitempty"`
	DatabaseIntegrationId     string              `json:"databaseIntegrationId,omitempty"`
}

// ChangePasswordRequest DTO.
type ChangePasswordRequest struct {
	CodeMashRequestBase
	CurrentPassword       string `json:"currentPassword,omitempty"`
	NewPassword           string `json:"newPassword,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// RequestPasswordResetRequest DTO.
type RequestPasswordResetRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// ConfirmPasswordResetRequest DTO.
type ConfirmPasswordResetRequest struct {
	CodeMashRequestBase
	Token                 string `json:"token,omitempty"`
	NewPassword           string `json:"newPassword,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// SendTestPasswordResetRequest DTO.
type SendTestPasswordResetRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// PasskeyAuthenticationOptionsRequest DTO.
type PasskeyAuthenticationOptionsRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// VerifyPasskeyAuthenticationRequest DTO.
type VerifyPasskeyAuthenticationRequest struct {
	CodeMashRequestBase
	CeremonyId        string `json:"ceremonyId,omitempty"`
	AssertionResponse string `json:"assertionResponse,omitempty"`
}

// ListPasskeysRequest DTO.
type ListPasskeysRequest struct {
	CodeMashRequestBase
}

// RenamePasskeyRequest DTO.
type RenamePasskeyRequest struct {
	CodeMashRequestBase
	CredentialId string `json:"credentialId,omitempty"`
	FriendlyName string `json:"friendlyName,omitempty"`
}

// RevokePasskeyRequest DTO.
type RevokePasskeyRequest struct {
	CodeMashRequestBase
	CredentialId string `json:"credentialId,omitempty"`
}

// UseRecoveryCodeRequest DTO.
type UseRecoveryCodeRequest struct {
	CodeMashRequestBase
	Email        string `json:"email,omitempty"`
	RecoveryCode string `json:"recoveryCode,omitempty"`
}

// RequestMagicLinkRequest DTO.
type RequestMagicLinkRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// ConsumeMagicLinkRequest DTO.
type ConsumeMagicLinkRequest struct {
	CodeMashRequestBase
	Token string `json:"token,omitempty"`
}

// HasPasskeyRequest DTO.
type HasPasskeyRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// StartEmailVerificationRequest DTO.
type StartEmailVerificationRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
}

// ConfirmEmailVerificationRequest DTO.
type ConfirmEmailVerificationRequest struct {
	CodeMashRequestBase
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

// PasskeyRegistrationOptionsRequest DTO.
type PasskeyRegistrationOptionsRequest struct {
	CodeMashRequestBase
	VerificationToken string `json:"verificationToken,omitempty"`
}

// VerifyPasskeyRegistrationRequest DTO.
type VerifyPasskeyRegistrationRequest struct {
	CodeMashRequestBase
	VerificationToken   string `json:"verificationToken,omitempty"`
	CeremonyId          string `json:"ceremonyId,omitempty"`
	AttestationResponse string `json:"attestationResponse,omitempty"`
	FriendlyName        string `json:"friendlyName,omitempty"`
}

// RefreshPasskeyTokenRequest DTO.
type RefreshPasskeyTokenRequest struct {
	CodeMashRequestBase
	RefreshToken string `json:"refreshToken,omitempty"`
}

// PasskeyLogoutRequest DTO.
type PasskeyLogoutRequest struct {
	CodeMashRequestBase
	RefreshToken string `json:"refreshToken,omitempty"`
}

// FindMergedTermTreeRequest DTO.
type FindMergedTermTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool   `json:"allLanguages"`
	TaxonomyName          string `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindTaxonomyTreeRequest DTO.
type FindTaxonomyTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool   `json:"allLanguages"`
	IncludeTerms          bool   `json:"includeTerms"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// FindTermsRequest DTO.
type FindTermsRequest struct {
	CodeMashListPaginationRequestBase
	AllLanguages          bool        `json:"allLanguages"`
	TaxonomyName          string      `json:"taxonomyName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SortDescending        bool        `json:"sortDescending"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// FindTermsChildrenRequest DTO.
type FindTermsChildrenRequest struct {
	CodeMashListPaginationRequestBase
	AllLanguages          bool        `json:"allLanguages"`
	TaxonomyName          string      `json:"taxonomyName,omitempty"`
	ParentId              string      `json:"parentId,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
}

// FindTermTreeRequest DTO.
type FindTermTreeRequest struct {
	CodeMashRequestBase
	AllLanguages          bool     `json:"allLanguages"`
	TaxonomyName          string   `json:"taxonomyName,omitempty"`
	RootTermId            string   `json:"rootTermId,omitempty"`
	Depth                 *float64 `json:"depth,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
}

// GetDatabaseSchemaRequest DTO.
type GetDatabaseSchemaRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetDatabaseSchemasRequest DTO.
type GetDatabaseSchemasRequest struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// AggregateRequest DTO.
type AggregateRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Pipeline              string `json:"pipeline,omitempty"`
}

// ChangeResponsibilityRequest DTO.
type ChangeResponsibilityRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	NewResponsibleUserId  string `json:"newResponsibleUserId,omitempty"`
}

// CountRequest DTO.
type CountRequest struct {
	CodeMashRequestBase
	CollectionName        string   `json:"collectionName,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	SchemaVersion         *float64 `json:"schemaVersion,omitempty"`
}

// DeleteManyRequest DTO.
type DeleteManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
	AllRecords            *bool  `json:"allRecords,omitempty"`
}

// DeleteOneRequest DTO.
type DeleteOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
}

// DistinctRequest DTO.
type DistinctRequest struct {
	CodeMashRequestBase
	CollectionName        string   `json:"collectionName,omitempty"`
	DatabaseIntegrationId string   `json:"databaseIntegrationId,omitempty"`
	Field                 string   `json:"field,omitempty"`
	Filter                string   `json:"filter,omitempty"`
	SchemaVersion         *float64 `json:"schemaVersion,omitempty"`
}

// ExecuteAggregateRequest DTO.
type ExecuteAggregateRequest struct {
	CodeMashRequestBase
	CollectionName        string            `json:"collectionName,omitempty"`
	AggregateId           string            `json:"aggregateId,omitempty"`
	DatabaseIntegrationId string            `json:"databaseIntegrationId,omitempty"`
	Tokens                map[string]string `json:"tokens,omitempty"`
}

// FindRequest DTO.
type FindRequest struct {
	CodeMashListPaginationRequestBase
	CollectionName        string      `json:"collectionName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SchemaVersion         *float64    `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	SortBy                string      `json:"sortBy,omitempty"`
	SortOrder             *float64    `json:"sortOrder,omitempty"`
	ExpandReferences      bool        `json:"expandReferences"`
	AllLanguages          bool        `json:"allLanguages"`
}

// FindOneRequest DTO.
type FindOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	ExpandReferences      bool   `json:"expandReferences"`
	AllLanguages          bool   `json:"allLanguages"`
}

// FindOwnRequest DTO.
type FindOwnRequest struct {
	CodeMashListPaginationRequestBase
	CollectionName        string      `json:"collectionName,omitempty"`
	DatabaseIntegrationId string      `json:"databaseIntegrationId,omitempty"`
	Filter                string      `json:"filter,omitempty"`
	SchemaVersion         *float64    `json:"schemaVersion,omitempty"`
	PagingArgs            *PagingArgs `json:"pagingArgs,omitempty"`
	ExpandReferences      bool        `json:"expandReferences"`
	AllLanguages          bool        `json:"allLanguages"`
}

// InsertManyRequest DTO.
type InsertManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Documents             string `json:"documents,omitempty"`
}

// InsertOneRequest DTO.
type InsertOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Document              string `json:"document,omitempty"`
}

// ReplaceOneRequest DTO.
type ReplaceOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Replacement           string `json:"replacement,omitempty"`
}

// UpdateManyRequest DTO.
type UpdateManyRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Filter                string `json:"filter,omitempty"`
	AllRecords            *bool  `json:"allRecords,omitempty"`
	Update                string `json:"update,omitempty"`
	ArrayFilters          string `json:"arrayFilters,omitempty"`
}

// UpdateOneRequest DTO.
type UpdateOneRequest struct {
	CodeMashRequestBase
	CollectionName        string `json:"collectionName,omitempty"`
	Id                    string `json:"id,omitempty"`
	DatabaseIntegrationId string `json:"databaseIntegrationId,omitempty"`
	Update                string `json:"update,omitempty"`
	ArrayFilters          string `json:"arrayFilters,omitempty"`
}

// CommitUploadRequest DTO.
type CommitUploadRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Path               string   `json:"path,omitempty"`
	ContentType        string   `json:"contentType,omitempty"`
	SizeBytes          *float64 `json:"sizeBytes,omitempty"`
	FileName           string   `json:"fileName,omitempty"`
	OwnerUserId        string   `json:"ownerUserId,omitempty"`
	ReassignOwner      *bool    `json:"reassignOwner,omitempty"`
}

// GetFileContentRequest DTO.
type GetFileContentRequest struct {
	RequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
	Token              string `json:"token,omitempty"`
}

// PutFileContentRequest DTO.
type PutFileContentRequest struct {
	RequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
	Token              string `json:"token,omitempty"`
}

// DeleteFileApiRequest DTO.
type DeleteFileApiRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// DeleteManyFilesApiRequest DTO.
type DeleteManyFilesApiRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Paths__            []string `json:"paths__,omitempty"`
}

// DownloadFileApiRequest DTO.
type DownloadFileApiRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetFileByIdRequest DTO.
type GetFileByIdRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Id                 string `json:"id,omitempty"`
}

// GetFileInfoRequest DTO.
type GetFileInfoRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetSignedUrlRequest DTO.
type GetSignedUrlRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Path               string   `json:"path,omitempty"`
	ExpirationSeconds  *float64 `json:"expirationSeconds,omitempty"`
}

// ListFilesRequest DTO.
type ListFilesRequest struct {
	CodeMashListPaginationRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
}

// GetPublicFileRequest DTO.
type GetPublicFileRequest struct {
	RequestBase
	PublicId string `json:"publicId,omitempty"`
	Name     string `json:"name,omitempty"`
}

// RegisterFilesRequest DTO.
type RegisterFilesRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string              `json:"filesIntegrationId,omitempty"`
	Items              []*RegisterFileItem `json:"items,omitempty"`
	ReassignOwner      *bool               `json:"reassignOwner,omitempty"`
	DryRun             *bool               `json:"dryRun,omitempty"`
}

// RequestUploadUrlRequest DTO.
type RequestUploadUrlRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string   `json:"filesIntegrationId,omitempty"`
	Path               string   `json:"path,omitempty"`
	ContentType        string   `json:"contentType,omitempty"`
	ExpirationSeconds  *float64 `json:"expirationSeconds,omitempty"`
	OwnerUserId        string   `json:"ownerUserId,omitempty"`
	ReassignOwner      *bool    `json:"reassignOwner,omitempty"`
}

// UseFileRequest DTO.
type UseFileRequest struct {
	CodeMashRequestBase
	FilesIntegrationId string `json:"filesIntegrationId,omitempty"`
	Path               string `json:"path,omitempty"`
	OwnerUserId        string `json:"ownerUserId,omitempty"`
	ReassignOwner      *bool  `json:"reassignOwner,omitempty"`
}

// RegisterOwnPushDeviceRequest DTO.
type RegisterOwnPushDeviceRequest struct {
	CodeMashRequestBase
	Provider       string `json:"provider,omitempty"`
	Token          string `json:"token,omitempty"`
	Platform       string `json:"platform,omitempty"`
	OsName         string `json:"osName,omitempty"`
	OsVersion      string `json:"osVersion,omitempty"`
	AppVersion     string `json:"appVersion,omitempty"`
	DeviceName     string `json:"deviceName,omitempty"`
	ModelName      string `json:"modelName,omitempty"`
	Brand          string `json:"brand,omitempty"`
	Manufacturer   string `json:"manufacturer,omitempty"`
	InstallationId string `json:"installationId,omitempty"`
	Locale         string `json:"locale,omitempty"`
	TimeZone       string `json:"timeZone,omitempty"`
}

// GetOwnPushDevicesRequest DTO.
type GetOwnPushDevicesRequest struct {
	CodeMashRequestBase
}

// DeleteOwnPushDeviceRequest DTO.
type DeleteOwnPushDeviceRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// DeleteOwnPushDeviceByTokenRequest DTO.
type DeleteOwnPushDeviceByTokenRequest struct {
	CodeMashRequestBase
	Token string `json:"token,omitempty"`
}

// GetOwnPushMessagesRequest DTO.
type GetOwnPushMessagesRequest struct {
	CodeMashListPaginationRequestBase
	UnreadOnly bool `json:"unreadOnly"`
}

// MarkOwnPushMessageReadRequest DTO.
type MarkOwnPushMessageReadRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// MarkAllOwnPushMessagesReadRequest DTO.
type MarkAllOwnPushMessagesReadRequest struct {
	CodeMashRequestBase
}

// WriteLogs DTO.
type WriteLogs struct {
	CodeMashRequestBase
	Entries []*WriteLogEntry `json:"entries,omitempty"`
}

// CreatePaymentOrderRequest DTO.
type CreatePaymentOrderRequest struct {
	CodeMashRequestBase
	BuyerUserId       string                   `json:"buyerUserId,omitempty"`
	MerchantReference string                   `json:"merchantReference,omitempty"`
	IdempotencyKey    string                   `json:"idempotencyKey,omitempty"`
	Currency          string                   `json:"currency,omitempty"`
	PriceVersion      string                   `json:"priceVersion,omitempty"`
	Lines             []*PaymentOrderLineInput `json:"lines,omitempty"`
	Metadata          map[string]string        `json:"metadata,omitempty"`
}

// GetPaymentOrderRequest DTO.
type GetPaymentOrderRequest struct {
	CodeMashRequestBase
	Id string `json:"id,omitempty"`
}

// GetPaymentOrdersRequest DTO.
type GetPaymentOrdersRequest struct {
	CodeMashListPaginationRequestBase
	BuyerUserId       string `json:"buyerUserId,omitempty"`
	MerchantReference string `json:"merchantReference,omitempty"`
	Status            string `json:"status,omitempty"`
}

// StartPaymentCheckoutRequest DTO.
type StartPaymentCheckoutRequest struct {
	CodeMashRequestBase
	Id                string `json:"id,omitempty"`
	IntegrationId     string `json:"integrationId,omitempty"`
	Method            string `json:"method,omitempty"`
	ReturnUrl         string `json:"returnUrl,omitempty"`
	PreferredCountry  string `json:"preferredCountry,omitempty"`
	PreferredProvider string `json:"preferredProvider,omitempty"`
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
