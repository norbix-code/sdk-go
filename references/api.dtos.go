//go:build ignore

/* Options:
Date: 2026-11-19 15:30:30
Version: 10.20
Tip: To override a DTO option, remove "//" prefix before updating
BaseUrl: http://localhost:5002

//GlobalNamespace:
//MakePropertiesOptional: False
//AddServiceStackTypes: True
//AddResponseStatus: False
//AddImplicitVersion:
//AddDescriptionAsComments: True
//IncludeTypes:
//ExcludeTypes:
//DefaultImports:
*/

package dtos

import (
	"time"
)

type JsonSchemaFieldDto struct {
	// @DataMember
	FieldName string `json:"fieldName"`
}

type StringFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Format *string `json:"format,omitempty"`
	// @DataMember
	Pattern *string `json:"pattern,omitempty"`
	// @DataMember
	MinLength *int `json:"minLength,omitempty"`
	// @DataMember
	MaxLength *int `json:"maxLength,omitempty"`
	// @DataMember
	TranslateOptions *IReadOnlyDictionary[string, string] `json:"translateOptions,omitempty"`
	// @DataMember
	Default *string `json:"default,omitempty"`
	// @DataMember
	Unique *bool `json:"unique,omitempty"`
}

type DecimalFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *float64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *float64 `json:"maximum,omitempty"`
	// @DataMember
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	// @DataMember
	Default *float64 `json:"default,omitempty"`
	// @DataMember
	Unique *bool `json:"unique,omitempty"`
}

type CurrencyDefaultDto struct {
	// @DataMember
	Value float64 `json:"value,omitempty"`
	// @DataMember
	Currency string `json:"currency"`
}

type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	AllowedCurrencies *IReadOnlyList[string] `json:"allowedCurrencies,omitempty"`
	// @DataMember
	MultipleOf *float64 `json:"multipleOf,omitempty"`
	// @DataMember
	Minimum *float64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *float64 `json:"maximum,omitempty"`
	// @DataMember
	Default *CurrencyDefaultDto `json:"default,omitempty"`
}

type BooleanFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Default *bool `json:"default,omitempty"`
}

type DateFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *int64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *int64 `json:"maximum,omitempty"`
	// @DataMember
	Default *int64 `json:"default,omitempty"`
}

type IntegerFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *int64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *int64 `json:"maximum,omitempty"`
	// @DataMember
	Default *int64 `json:"default,omitempty"`
	// @DataMember
	Unique *bool `json:"unique,omitempty"`
}

type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	AllowedTypes *IReadOnlyList[string] `json:"allowedTypes,omitempty"`
}

type TagsFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	MinItems *int `json:"minItems,omitempty"`
	// @DataMember
	MaxItems *int `json:"maxItems,omitempty"`
	// @DataMember
	Default *IReadOnlyList[string] `json:"default,omitempty"`
}

type FileFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Storages *IReadOnlyList[string] `json:"storages,omitempty"`
	// @DataMember
	MinItems *int `json:"minItems,omitempty"`
	// @DataMember
	MaxItems *int `json:"maxItems,omitempty"`
	// @DataMember
	AllowedFileType *string `json:"allowedFileType,omitempty"`
	// @DataMember
	MaxSizeMb *float64 `json:"maxSizeMb,omitempty"`
}

type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
	// @DataMember
	DisplayField *string `json:"displayField,omitempty"`
}

type CollectionSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	CollectionId *string `json:"collectionId,omitempty"`
	// @DataMember
	DisplayField *string `json:"displayField,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type UserSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
	// @DataMember
	DisplayField *string `json:"displayField,omitempty"`
}

type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
	// @DataMember
	DisplayField *string `json:"displayField,omitempty"`
}

type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Values *IReadOnlyList[string] `json:"values,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
	// @DataMember
	Default *IReadOnlyList[string] `json:"default,omitempty"`
}

type ObjectFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Properties IReadOnlyList[interface{}] `json:"properties"`
	// @DataMember
	Required *IReadOnlyList[string] `json:"required,omitempty"`
}

type ArrayFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Items interface{} `json:"items"`
	// @DataMember
	MinItems *int `json:"minItems,omitempty"`
	// @DataMember
	MaxItems *int `json:"maxItems,omitempty"`
	// @DataMember
	UniqueItems *bool `json:"uniqueItems,omitempty"`
}

type JsonFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	MaxBytes *int `json:"maxBytes,omitempty"`
}

// @DataContract(Namespace="http://codemash.io/types/")
type RequestBase struct {
	/** @description TimeZone */
	// @DataMember
	// @ApiMember(DataType="string", Description="TimeZone", Name="TimeZoneId", ParameterType="header")
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	/** @description The CodeMash API version used to fetch data from the API. If not specified, the last version will be used.  E.g.: v3 */
	// @DataMember
	// @ApiMember(DataType="string", Description="The CodeMash API version used to fetch data from the API. If not specified, the last version will be used.  E.g.: v3", IsRequired=true, Name="version", ParameterType="path")
	Version string `json:"version"`
	/** @description CorrelationId for each request */
	// @DataMember
	// @ApiMember(DataType="string", Description="CorrelationId for each request", Name="CorrelationId", ParameterType="header")
	CorrelationId *string `json:"correlationId,omitempty"`
}

// @DataContract(Namespace="http://codemash.io/types/")
type CodeMashRequestBase struct {
	RequestBase
	/** @description ID of your project. Can be passed in a header as norbix-project-id. */
	// @DataMember
	// @ApiMember(DataType="string", Description="ID of your project. Can be passed in a header as norbix-project-id.", IsRequired=true, Name="norbix-project-id", ParameterType="header")
	ProjectId string `json:"projectId"`
	/** @description Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env.", Name="norbix-env", ParameterType="header")
	Env *string `json:"env,omitempty"`
}

type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale        = "Female"
	GenderOther         = "Other"
)

type MarketingBlockReason string

const (
	MarketingBlockReasonUnspecified  MarketingBlockReason = "Unspecified"
	MarketingBlockReasonUnsubscribed                      = "Unsubscribed"
	MarketingBlockReasonComplaint                         = "Complaint"
	MarketingBlockReasonHardBounce                        = "HardBounce"
	MarketingBlockReasonInvalidEmail                      = "InvalidEmail"
	MarketingBlockReasonAdminBlock                        = "AdminBlock"
)

type UserGeneralInfoDto struct {
	Phone                     *string                `json:"phone,omitempty"`
	PrimaryEmail              *string                `json:"primaryEmail,omitempty"`
	DisplayName               *string                `json:"displayName,omitempty"`
	FirstName                 *string                `json:"firstName,omitempty"`
	LastName                  *string                `json:"lastName,omitempty"`
	FullName                  *string                `json:"fullName,omitempty"`
	AddressLine1              *string                `json:"addressLine1,omitempty"`
	AddressLine2              *string                `json:"addressLine2,omitempty"`
	Country                   *string                `json:"country,omitempty"`
	City                      *string                `json:"city,omitempty"`
	State                     *string                `json:"state,omitempty"`
	PostalCode                *string                `json:"postalCode,omitempty"`
	Company                   *string                `json:"company,omitempty"`
	Gender                    *Gender                `json:"gender,omitempty"`
	BirthDate                 *int64                 `json:"birthDate,omitempty"`
	TimeZone                  *string                `json:"timeZone,omitempty"`
	Language                  *string                `json:"language,omitempty"`
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
	ExtraMetadata             *string                `json:"extraMetadata,omitempty"`
	Notes                     *string                `json:"notes,omitempty"`
}

// @DataContract
type SaveUser struct {
	CodeMashRequestBase
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description User Info */
	// @DataMember
	// @ApiMember(DataType="object", Description="User Info", Name="UserGeneralInfo", ParameterType="body")
	UserGeneralInfo *UserGeneralInfoDto `json:"userGeneralInfo,omitempty"`
	/** @description Attach this login to an existing user id. Optional. */
	// @DataMember
	// @ApiMember(Description="Attach this login to an existing user id. Optional.")
	UserId *string `json:"userId,omitempty"`
	/** @description Ignore UserRegistersAsRole from Membership Settings */
	// @DataMember
	// @ApiMember(DataType="boolean", Description="Ignore UserRegistersAsRole from Membership Settings", Name="IgnoreUserRegistersAsRole", ParameterType="body")
	IgnoreUserRegistersAsRole bool `json:"ignoreUserRegistersAsRole,omitempty"`
}

// @DataContract
type SaveUserWithRolesBase struct {
	SaveUser
	// @DataMember
	Roles []string `json:"roles"`
	/** @description Tags to put on the new user's contact (free labels campaigns and segments select by, normalised like any tag). Only here, on the routes that need membership:create — a public self sign-up cannot tag itself. Ignored for a service user (no contact). */
	// @DataMember
	// @ApiMember(Description="Tags to put on the new user's contact (free labels campaigns and segments select by, normalised like any tag). Only here, on the routes that need membership:create — a public self sign-up cannot tag itself. Ignored for a service user (no contact).")
	Tags []string `json:"tags,omitempty"`
}

type CodeMashListPaginationRequestBase struct {
	RequestBase
	/** @description ID of your project. Can be passed in a header as norbix-project-id. */
	// @DataMember
	// @ApiMember(DataType="string", Description="ID of your project. Can be passed in a header as norbix-project-id.", IsRequired=true, Name="norbix-project-id", ParameterType="header")
	ProjectId string `json:"projectId"`
	/** @description Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env.", Name="norbix-env", ParameterType="header")
	Env *string `json:"env,omitempty"`
	/** @description Cursor token — fetch the page AFTER this item. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Cursor token — fetch the page AFTER this item.", Name="startingAfter", ParameterType="query")
	StartingAfter *string `json:"startingAfter,omitempty"`
	/** @description Cursor token — fetch the page BEFORE this item. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Cursor token — fetch the page BEFORE this item.", Name="endingBefore", ParameterType="query")
	EndingBefore *string `json:"endingBefore,omitempty"`
	/** @description Amount of records to return. */
	// @DataMember
	// @ApiMember(DataType="integer", Description="Amount of records to return.", Format="int32", Name="pageSize", ParameterType="query")
	PageSize *int `json:"pageSize,omitempty"`
}

type CursorArgs struct {
	Field string `json:"field"`
	Order int    `json:"order,omitempty"`
}

type PagingArgs struct {
	CursorArgs    *CursorArgs `json:"cursorArgs,omitempty"`
	PageSize      *int        `json:"pageSize,omitempty"`
	StartingAfter *string     `json:"startingAfter,omitempty"`
	EndingBefore  *string     `json:"endingBefore,omitempty"`
}

// @DataContract
type RegisterFileItem struct {
	// @DataMember
	Path *string `json:"path,omitempty"`
	// @DataMember
	OwnerUserId *string `json:"ownerUserId,omitempty"`
	// @DataMember
	ContentType *string `json:"contentType,omitempty"`
	// @DataMember
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember
	FileName *string `json:"fileName,omitempty"`
}

type WriteLogEntry struct {
	/** @description information (default), warning or error. */
	// @ApiMember(Description="information (default), warning or error.")
	Level *string `json:"level,omitempty"`
	/** @description A short line, 1–200 characters — what the Logs list shows. */
	// @ApiMember(Description="A short line, 1–200 characters — what the Logs list shows.", IsRequired=true)
	Title string `json:"title"`
	/** @description Details, at most 8 000 characters. */
	// @ApiMember(Description="Details, at most 8 000 characters.")
	Message *string `json:"message,omitempty"`
	/** @description Your own code for this kind of line, e.g. order:paid — lower case letters, digits and _ . : -, at most 64. Stored as app:<code>; app:log when empty. */
	// @ApiMember(Description="Your own code for this kind of line, e.g. order:paid — lower case letters, digits and _ . : -, at most 64. Stored as app:<code>; app:log when empty.")
	EventCode *string `json:"eventCode,omitempty"`
	/** @description Groups your lines into one flow in the Logs details view: at most 64 letters, digits, _ or -. Empty = the id of this request. */
	// @ApiMember(Description="Groups your lines into one flow in the Logs details view: at most 64 letters, digits, _ or -. Empty = the id of this request.")
	CorrelationId *string `json:"correlationId,omitempty"`
	/** @description When it happened (UTC). Empty = when Norbix received it. At most 7 days old and at most 5 minutes ahead. */
	// @ApiMember(Description="When it happened (UTC). Empty = when Norbix received it. At most 7 days old and at most 5 minutes ahead.")
	Timestamp *time.Time `json:"timestamp,omitempty"`
	/** @description Extra fields, at most 20; key at most 64, value at most 2 000 characters. Values under secret-looking keys (password, token, apiKey …) are masked. */
	// @ApiMember(Description="Extra fields, at most 20; key at most 64, value at most 2 000 characters. Values under secret-looking keys (password, token, apiKey …) are masked.")
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type PaymentOrderLineInput struct {
	// @DataMember
	Sku *string `json:"sku,omitempty"`
	// @DataMember
	Title *string `json:"title,omitempty"`
	// @DataMember
	Quantity int `json:"quantity,omitempty"`
	// @DataMember
	UnitAmountMinor int64 `json:"unitAmountMinor,omitempty"`
}

// @DataContract
type FileChecksumDto struct {
	// @DataMember(Order=1)
	Algorithm string `json:"algorithm"`
	// @DataMember(Order=2)
	Hash string `json:"hash"`
}

// @DataContract
type FileResourceDto struct {
	// @DataMember(Order=1)
	Id *string `json:"id,omitempty"`
	// @DataMember(Order=2)
	OriginalFileName string `json:"originalFileName"`
	// @DataMember(Order=3)
	Extension string `json:"extension"`
	// @DataMember(Order=4)
	StoredFileName string `json:"storedFileName"`
	// @DataMember(Order=5)
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember(Order=6)
	Checksum *FileChecksumDto `json:"checksum,omitempty"`
}

type FileProvider string

const (
	FileProviderLocal              FileProvider = "Local"
	FileProviderAwsS3                           = "AwsS3"
	FileProviderAzureBlobStorage                = "AzureBlobStorage"
	FileProviderGoogleCloudStorage              = "GoogleCloudStorage"
	FileProviderFtp                             = "Ftp"
	FileProviderAppleICloud                     = "AppleICloud"
	FileProviderDropBox                         = "DropBox"
	FileProviderGoogleDrive                     = "GoogleDrive"
)

// @DataContract
type FileResourceRefDto struct {
	// @DataMember(Order=1)
	Resource FileResourceDto `json:"resource"`
	// @DataMember(Order=2)
	IntegrationId string `json:"integrationId"`
	// @DataMember(Order=3)
	Provider FileProvider `json:"provider,omitempty"`
	// @DataMember(Order=4)
	Path string `json:"path"`
	// @DataMember(Order=5)
	PublicUrl *string `json:"publicUrl,omitempty"`
	// @DataMember(Order=6)
	IsPublic bool `json:"isPublic,omitempty"`
}

type ImportColumnMappingDto struct {
	// @DataMember
	CsvColumnIndex int `json:"csvColumnIndex,omitempty"`
	// @DataMember
	CsvHeader *string `json:"csvHeader,omitempty"`
	// @DataMember
	PropertyName *string `json:"propertyName,omitempty"`
	// @DataMember
	DontImportOnError bool `json:"dontImportOnError,omitempty"`
}

type EmailCampaignRecipientsSourceTypes string

const (
	EmailCampaignRecipientsSourceTypesAllUsers       EmailCampaignRecipientsSourceTypes = "AllUsers"
	EmailCampaignRecipientsSourceTypesSpecifiedUsers                                    = "SpecifiedUsers"
	EmailCampaignRecipientsSourceTypesAccountUsers                                      = "AccountUsers"
	EmailCampaignRecipientsSourceTypesEmail                                             = "Email"
	EmailCampaignRecipientsSourceTypesCollection                                        = "Collection"
	EmailCampaignRecipientsSourceTypesTriggerSubject                                    = "TriggerSubject"
)

type TokenMappingResolverType string

const (
	TokenMappingResolverTypeNotSet            TokenMappingResolverType = "NotSet"
	TokenMappingResolverTypeCustom                                     = "Custom"
	TokenMappingResolverTypeProject                                    = "Project"
	TokenMappingResolverTypeProjectSocials                             = "ProjectSocials"
	TokenMappingResolverTypeInitiator                                  = "Initiator"
	TokenMappingResolverTypeRecipient                                  = "Recipient"
	TokenMappingResolverTypeSchemaRecord                               = "SchemaRecord"
	TokenMappingResolverTypeTargetUser                                 = "TargetUser"
	TokenMappingResolverTypeTagDefinitions                             = "TagDefinitions"
	TokenMappingResolverTypeEmailSignatures                            = "EmailSignatures"
	TokenMappingResolverTypeCampaign                                   = "Campaign"
	TokenMappingResolverTypeTemplate                                   = "Template"
	TokenMappingResolverTypeEmailFooters                               = "EmailFooters"
	TokenMappingResolverTypeOld                                        = "Old"
	TokenMappingResolverTypeNew                                        = "New"
	TokenMappingResolverTypeRequestInput                               = "RequestInput"
	TokenMappingResolverTypeIntegrationConfig                          = "IntegrationConfig"
	TokenMappingResolverTypeIntegrationSecret                          = "IntegrationSecret"
	TokenMappingResolverTypeStepOutput                                 = "StepOutput"
)

type SecretValueFormat string

const (
	SecretValueFormatRaw      SecretValueFormat = "Raw"
	SecretValueFormatBearer                     = "Bearer"
	SecretValueFormatBasic                      = "Basic"
	SecretValueFormatPrefixed                   = "Prefixed"
)

// @DataContract
type TokenMappingDto struct {
	// @DataMember
	Key string `json:"key"`
	// @DataMember
	Value string `json:"value"`
	// @DataMember
	Resolver TokenMappingResolverType `json:"resolver,omitempty"`
	// @DataMember
	SecretKeys []string `json:"secretKeys,omitempty"`
	// @DataMember
	SecretFormat *SecretValueFormat `json:"secretFormat,omitempty"`
}

type EmailCampaignRequest struct {
	Source                  EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	TemplateId              string                             `json:"templateId"`
	IntegrationId           *string                            `json:"integrationId,omitempty"`
	ValidationIntegrationId *string                            `json:"validationIntegrationId,omitempty"`
	Language                *string                            `json:"language,omitempty"`
	InitiatorId             *string                            `json:"initiatorId,omitempty"`
	Notes                   *string                            `json:"notes,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
	// @DataMember
	CampaignTime *int64 `json:"campaignTime,omitempty"`
}

type SmsCampaignRecipientsSourceTypes string

const (
	SmsCampaignRecipientsSourceTypesAllUsers       SmsCampaignRecipientsSourceTypes = "AllUsers"
	SmsCampaignRecipientsSourceTypesSpecifiedUsers                                  = "SpecifiedUsers"
	SmsCampaignRecipientsSourceTypesAccountUsers                                    = "AccountUsers"
	SmsCampaignRecipientsSourceTypesPhoneNumbers                                    = "PhoneNumbers"
	SmsCampaignRecipientsSourceTypesCollection                                      = "Collection"
	SmsCampaignRecipientsSourceTypesTriggerSubject                                  = "TriggerSubject"
)

// @Flags()
type RespectTimeZoneSettings int

const (
	RespectTimeZoneSettingsRespectToLastLoginZone           RespectTimeZoneSettings = 1
	RespectTimeZoneSettingsRespectToRegistrationZone        RespectTimeZoneSettings = 2
	RespectTimeZoneSettingsRespectToRegistrationProjectZone RespectTimeZoneSettings = 4
)

// @DataContract
type SmsCampaignDeliverySettingsDto struct {
	// @DataMember
	RecipientsSourceType SmsCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
	// @DataMember
	CampaignTime *int64 `json:"campaignTime,omitempty"`
	// @DataMember
	RespectTimeZoneSettings *RespectTimeZoneSettings `json:"respectTimeZoneSettings,omitempty"`
}

// @DataContract
type SmsToAllUsersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	// @DataMember
	RolesNames []string `json:"rolesNames,omitempty"`
	// @DataMember
	UserTags []string `json:"userTags,omitempty"`
}

// @DataContract
type SmsToUsersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	// @DataMember
	Recipients []string `json:"recipients"`
}

// @DataContract
type SmsToAccountUsersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	// @DataMember
	Recipients []string `json:"recipients"`
}

type CollectionEmailCampaignRecipientField string

const (
	CollectionEmailCampaignRecipientFieldUser  CollectionEmailCampaignRecipientField = "User"
	CollectionEmailCampaignRecipientFieldEmail                                       = "Email"
)

// @DataContract
type SmsToCollectionRecordsDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	// @DataMember
	Fields []string `json:"fields"`
	// @DataMember
	FieldType CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	// @DataMember
	SchemaName string `json:"schemaName"`
	// @DataMember
	RoleNames []string `json:"roleNames,omitempty"`
	// @DataMember
	Languages []string `json:"languages,omitempty"`
	// @DataMember
	RecordIds []string `json:"recordIds,omitempty"`
}

// @DataContract
type SmsToPhoneNumbersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	// @DataMember
	PhoneNumbers []string `json:"phoneNumbers"`
}

type PushCampaignRecipientsSourceTypes string

const (
	PushCampaignRecipientsSourceTypesAllUsers       PushCampaignRecipientsSourceTypes = "AllUsers"
	PushCampaignRecipientsSourceTypesSpecifiedUsers                                   = "SpecifiedUsers"
	PushCampaignRecipientsSourceTypesCollection                                       = "Collection"
	PushCampaignRecipientsSourceTypesDevices                                          = "Devices"
	PushCampaignRecipientsSourceTypesAccountUsers                                     = "AccountUsers"
	PushCampaignRecipientsSourceTypesTriggerSubject                                   = "TriggerSubject"
)

type PushCampaignRequest struct {
	Source        PushCampaignRecipientsSourceTypes `json:"source,omitempty"`
	TemplateId    string                            `json:"templateId"`
	IntegrationId *string                           `json:"integrationId,omitempty"`
	Language      *string                           `json:"language,omitempty"`
	InitiatorId   *string                           `json:"initiatorId,omitempty"`
	Notes         *string                           `json:"notes,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
	// @DataMember
	CampaignTime *int64 `json:"campaignTime,omitempty"`
}

// @DataContract
type CodeMashRelease string

const (
	CodeMashReleaseNotSet         CodeMashRelease = "NotSet"
	CodeMashReleaseCommunity                      = "Community"
	CodeMashReleaseManagedService                 = "ManagedService"
	CodeMashReleaseEnterprise                     = "Enterprise"
)

type CodeMashRuntime string

const (
	CodeMashRuntimeDevelopment CodeMashRuntime = "Development"
	CodeMashRuntimeCI                          = "CI"
	CodeMashRuntimeStaging                     = "Staging"
	CodeMashRuntimeProduction                  = "Production"
)

type EchoRegionDto struct {
	Code        string `json:"code"`
	DisplayName string `json:"displayName"`
	ApiUrl      string `json:"apiUrl"`
	HubUrl      string `json:"hubUrl"`
}

type EchoAgentSnippetDto struct {
	Client      string `json:"client"`
	Kind        string `json:"kind"`
	Recommended bool   `json:"recommended,omitempty"`
	Auth        string `json:"auth"`
	Config      string `json:"config"`
}

type EchoAgentDto struct {
	McpUrl                 string                `json:"mcpUrl"`
	OAuthMetadataUrl       *string               `json:"oAuthMetadataUrl,omitempty"`
	InstallationType       string                `json:"installationType"`
	OnboardingDocsUrl      string                `json:"onboardingDocsUrl"`
	ToolsUrl               string                `json:"toolsUrl"`
	OpenApiUrl             *string               `json:"openApiUrl,omitempty"`
	CliPackage             string                `json:"cliPackage"`
	CliInstallCommand      string                `json:"cliInstallCommand"`
	DeviceAuthorizationUrl string                `json:"deviceAuthorizationUrl"`
	DeviceTokenUrl         string                `json:"deviceTokenUrl"`
	Snippets               []EchoAgentSnippetDto `json:"snippets"`
}

type EchoHostingPlatformDto struct {
	Provider string `json:"provider"`
	IsAws    bool   `json:"isAws,omitempty"`
}

// @DataContract
type QuickConnectProviderDto struct {
	// @DataMember
	Module string `json:"module"`
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	Available bool `json:"available,omitempty"`
	// @DataMember
	CallbackKind string `json:"callbackKind"`
	// @DataMember
	RequiredInputs []string `json:"requiredInputs"`
	// @DataMember
	Modes []string `json:"modes"`
	// @DataMember
	RedirectUrl string `json:"redirectUrl"`
	// @DataMember
	PlatformWebhookUrls []string `json:"platformWebhookUrls"`
}

type EchoConnectDto struct {
	Providers []QuickConnectProviderDto `json:"providers"`
}

type EchoKeylessDto struct {
	Issuer            string  `json:"issuer"`
	Available         bool    `json:"available,omitempty"`
	UnavailableReason *string `json:"unavailableReason,omitempty"`
}

type PublicBrandDto struct {
	DisplayName string  `json:"displayName"`
	MainColor   *string `json:"mainColor,omitempty"`
	AccentColor *string `json:"accentColor,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty"`
	IconUrl     *string `json:"iconUrl,omitempty"`
}

type PublicSocialSignInDto struct {
	Provider    string `json:"provider"`
	DisplayName string `json:"displayName"`
	StartPath   string `json:"startPath"`
}

type PublicPasswordPolicyDto struct {
	MinLength      int     `json:"minLength,omitempty"`
	MaxLength      *int    `json:"maxLength,omitempty"`
	MinNumbers     *int    `json:"minNumbers,omitempty"`
	MinUpper       *int    `json:"minUpper,omitempty"`
	MinLower       *int    `json:"minLower,omitempty"`
	MinSpecial     *int    `json:"minSpecial,omitempty"`
	AllowedSpecial *string `json:"allowedSpecial,omitempty"`
}

type PublicAuthDto struct {
	SocialProviders []string                 `json:"socialProviders"`
	SocialSignIn    []PublicSocialSignInDto  `json:"socialSignIn"`
	Passkey         bool                     `json:"passkey,omitempty"`
	Methods         []string                 `json:"methods,omitempty"`
	PasswordPolicy  *PublicPasswordPolicyDto `json:"passwordPolicy,omitempty"`
}

type PublicAiAssistantDto struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Welcome *string `json:"welcome,omitempty"`
}

type PublicAiChatDto struct {
	Enabled    bool                   `json:"enabled,omitempty"`
	Assistants []PublicAiAssistantDto `json:"assistants"`
}

type ErrorDto struct {
	Message    string            `json:"message"`
	ErrorCode  *string           `json:"errorCode,omitempty"`
	Context    map[string]string `json:"context,omitempty"`
	StackTrace []ErrorDto        `json:"stackTrace,omitempty"`
}

type CodeMashResponseStatus struct {
	IsSuccess bool       `json:"isSuccess,omitempty"`
	Errors    []ErrorDto `json:"errors,omitempty"`
}

// @DataContract
type ResponseBase struct {
	// @DataMember
	ResponseStatus CodeMashResponseStatus `json:"responseStatus"`
}

type EndUserChatAttachment struct {
	Id           string    `json:"id"`
	SessionId    string    `json:"sessionId"`
	FileName     string    `json:"fileName"`
	ContentType  string    `json:"contentType"`
	Kind         string    `json:"kind"`
	Size         int64     `json:"size,omitempty"`
	Summary      *string   `json:"summary,omitempty"`
	CreatedAtUtc time.Time `json:"createdAtUtc,omitempty"`
}

type EndUserChatMemoryNote struct {
	Id           string    `json:"id"`
	SessionId    string    `json:"sessionId"`
	Kind         string    `json:"kind"`
	Text         string    `json:"text"`
	CreatedAtUtc time.Time `json:"createdAtUtc,omitempty"`
}

type EndUserChatAssistant struct {
	Id             string  `json:"id"`
	Name           string  `json:"name"`
	WelcomeMessage *string `json:"welcomeMessage,omitempty"`
	IsDefault      bool    `json:"isDefault,omitempty"`
	MemoryEnabled  bool    `json:"memoryEnabled,omitempty"`
}

type EndUserChatPlan struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	QuotaUnit    string `json:"quotaUnit"`
	MonthlyQuota int64  `json:"monthlyQuota,omitempty"`
	Used         int64  `json:"used,omitempty"`
	Remaining    int64  `json:"remaining,omitempty"`
	Attachments  bool   `json:"attachments,omitempty"`
	Rag          bool   `json:"rag,omitempty"`
	Memory       bool   `json:"memory,omitempty"`
}

type EndUserChatSession struct {
	Id           string    `json:"id"`
	AssistantId  *string   `json:"assistantId,omitempty"`
	Title        *string   `json:"title,omitempty"`
	IsPinned     bool      `json:"isPinned,omitempty"`
	IsArchived   bool      `json:"isArchived,omitempty"`
	LastSeq      int64     `json:"lastSeq,omitempty"`
	CreatedAtUtc time.Time `json:"createdAtUtc,omitempty"`
	UpdatedAtUtc time.Time `json:"updatedAtUtc,omitempty"`
}

type AiChatEntryWireDto struct {
	Kind                 string     `json:"kind"`
	Id                   string     `json:"id"`
	Seq                  int64      `json:"seq,omitempty"`
	AtUtc                time.Time  `json:"atUtc,omitempty"`
	RefEntryId           *string    `json:"refEntryId,omitempty"`
	WorkItemId           *string    `json:"workItemId,omitempty"`
	Feedback             *string    `json:"feedback,omitempty"`
	FeedbackAtUtc        *time.Time `json:"feedbackAtUtc,omitempty"`
	FeedbackByUserAuthId *string    `json:"feedbackByUserAuthId,omitempty"`
}

type EndUserAiToolParameter struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Required    bool    `json:"required,omitempty"`
	Description *string `json:"description,omitempty"`
}

type EndUserAiTool struct {
	Name                 string                   `json:"name"`
	Description          string                   `json:"description"`
	Toolsets             []string                 `json:"toolsets"`
	RequiresConfirmation bool                     `json:"requiresConfirmation,omitempty"`
	Parameters           []EndUserAiToolParameter `json:"parameters"`
}

type AuthType string

const (
	AuthTypeService  AuthType = "Service"
	AuthTypeEmail             = "Email"
	AuthTypeUserName          = "UserName"
	AuthTypePhone             = "Phone"
	AuthTypeGuest             = "Guest"
	AuthTypeSocial            = "Social"
)

type AccessInformationDto struct {
	Ip       *string    `json:"ip,omitempty"`
	Date     *time.Time `json:"date,omitempty"`
	TimeZone *string    `json:"timeZone,omitempty"`
}

type RegistrationDto struct {
	RegistrationInformation AccessInformationDto `json:"registrationInformation"`
}

type LoginDto struct {
	NeedChangePasswordOnNextLogin bool                  `json:"needChangePasswordOnNextLogin,omitempty"`
	LastAccessInformation         *AccessInformationDto `json:"lastAccessInformation,omitempty"`
}

type AuthStatus string

const (
	AuthStatusRegistered        AuthStatus = "Registered"
	AuthStatusPendingValidation            = "PendingValidation"
	AuthStatusInvited                      = "Invited"
	AuthStatusActive                       = "Active"
	AuthStatusUnregistered                 = "Unregistered"
	AuthStatusSuspended                    = "Suspended"
	AuthStatusInActive                     = "InActive"
	AuthStatusBlocked                      = "Blocked"
)

type AuthDto struct {
	Id                string              `json:"id"`
	Type              AuthType            `json:"type,omitempty"`
	Email             *string             `json:"email,omitempty"`
	UserName          *string             `json:"userName,omitempty"`
	Registration      *RegistrationDto    `json:"registration,omitempty"`
	Login             *LoginDto           `json:"login,omitempty"`
	GeneralInfo       *UserGeneralInfoDto `json:"generalInfo,omitempty"`
	Roles             []string            `json:"roles,omitempty"`
	PushDevices       []string            `json:"pushDevices,omitempty"`
	Tags              []string            `json:"tags,omitempty"`
	Status            AuthStatus          `json:"status,omitempty"`
	CreatedOn         time.Time           `json:"createdOn,omitempty"`
	ModifiedOn        time.Time           `json:"modifiedOn,omitempty"`
	CredentialVersion *string             `json:"credentialVersion,omitempty"`
}

type PaginatedResponse[TViewModelProjection any] struct {
	Items         IList[TViewModelProjection] `json:"items"`
	HasMore       bool                        `json:"hasMore,omitempty"`
	HasPrevious   bool                        `json:"hasPrevious,omitempty"`
	StartingAfter *string                     `json:"startingAfter,omitempty"`
	EndingBefore  *string                     `json:"endingBefore,omitempty"`
}

type UserMarketingPreferencesDto struct {
	BlockAllMarketingMessages bool                   `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[string][]string    `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason `json:"blockReasons,omitempty"`
}

type PasskeyListItemDto struct {
	CredentialId    string    `json:"credentialId"`
	FriendlyName    string    `json:"friendlyName"`
	RegisteredOnUtc time.Time `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   time.Time `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool      `json:"isRevoked,omitempty"`
}

type TermMultiParentDto struct {
	// @DataMember
	TaxonomyId string `json:"taxonomyId"`
	// @DataMember
	ParentId string `json:"parentId"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Names map[string]string `json:"names,omitempty"`
}

type TermTreeDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	TaxonomyName *string `json:"taxonomyName,omitempty"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Order *int `json:"order,omitempty"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Names map[string]string `json:"names,omitempty"`
	// @DataMember
	Slug *string `json:"slug,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Descriptions map[string]string `json:"descriptions,omitempty"`
	// @DataMember
	MultiParents []TermMultiParentDto `json:"multiParents,omitempty"`
	// @DataMember
	Meta *interface{} `json:"meta,omitempty"`
	// @DataMember
	Children []TermTreeDto `json:"children,omitempty"`
}

type TaxonomyTreeDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	TaxonomySlug string `json:"taxonomySlug"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Children []TaxonomyTreeDto `json:"children,omitempty"`
	// @DataMember
	Terms []TermTreeDto `json:"terms,omitempty"`
}

type TermDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	TaxonomyName *string `json:"taxonomyName,omitempty"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Order *int `json:"order,omitempty"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Names map[string]string `json:"names,omitempty"`
	// @DataMember
	Slug *string `json:"slug,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Descriptions map[string]string `json:"descriptions,omitempty"`
	// @DataMember
	MultiParents []TermMultiParentDto `json:"multiParents,omitempty"`
	// @DataMember
	Meta *interface{} `json:"meta,omitempty"`
	// @DataMember
	Origin *string `json:"origin,omitempty"`
}

type DataSchemaDto struct {
	// @DataMember
	Json string `json:"json"`
	// @DataMember
	Fields []interface{} `json:"fields"`
}

type VisualSchemaDto struct {
	// @DataMember
	Json string `json:"json"`
}

type SchemaUniqueIndexDto struct {
	// @DataMember
	Fields []string `json:"fields"`
}

type SchemaSettingsDto struct {
	// @DataMember
	SoftDelete bool `json:"softDelete,omitempty"`
	// @DataMember
	HasRecordOwner bool `json:"hasRecordOwner,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	UniqueIndexes []SchemaUniqueIndexDto `json:"uniqueIndexes,omitempty"`
	// @DataMember
	EndUserOwnRecords *bool `json:"endUserOwnRecords,omitempty"`
}

type SchemaEmbedSettingsDto struct {
	// @DataMember
	Enabled bool `json:"enabled,omitempty"`
	// @DataMember
	Fields []string `json:"fields"`
	// @DataMember
	EmbeddingIntegrationId *string `json:"embeddingIntegrationId,omitempty"`
	// @DataMember
	PerUser bool `json:"perUser,omitempty"`
}

type TriggerType string

const (
	TriggerTypeMembership TriggerType = "Membership"
	TriggerTypeSchema                 = "Schema"
	TriggerTypeFiles                  = "Files"
	TriggerTypePayments               = "Payments"
	TriggerTypeAi                     = "Ai"
)

type TriggerActionType string

const (
	TriggerActionTypeCode        TriggerActionType = "Code"
	TriggerActionTypePush                          = "Push"
	TriggerActionTypeSms                           = "Sms"
	TriggerActionTypeEmail                         = "Email"
	TriggerActionTypeWebhookCall                   = "WebhookCall"
	TriggerActionTypeSseCall                       = "SseCall"
	TriggerActionTypeMarketplace                   = "Marketplace"
	TriggerActionTypeWorkflow                      = "Workflow"
)

// @DataContract
type TriggerActionDto struct {
	// @DataMember
	Type TriggerActionType `json:"type,omitempty"`
	// @DataMember
	IntegrationId *string `json:"integrationId,omitempty"`
}

// @DataContract
type ExecutionIdentitySummaryDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	IsDefault bool `json:"isDefault,omitempty"`
	// @DataMember
	Status string `json:"status"`
}

// @DataContract
type TriggerDto struct {
	// @DataMember
	Type TriggerType `json:"type,omitempty"`
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	ThenAction TriggerActionDto `json:"thenAction"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	ActivationCode *string `json:"activationCode,omitempty"`
	// @DataMember
	SavedByAuthId *string `json:"savedByAuthId,omitempty"`
	// @DataMember
	ExecutionIdentityId *string `json:"executionIdentityId,omitempty"`
	// @DataMember
	ExecutionIdentity *ExecutionIdentitySummaryDto `json:"executionIdentity,omitempty"`
}

type SchemaDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	SchemaName string `json:"schemaName"`
	// @DataMember
	SchemaSlug *string `json:"schemaSlug,omitempty"`
	// @DataMember
	Version int `json:"version,omitempty"`
	// @DataMember
	MetaSchemaVersion int `json:"metaSchemaVersion,omitempty"`
	// @DataMember
	DataSchema DataSchemaDto `json:"dataSchema"`
	// @DataMember
	VisualSchema VisualSchemaDto `json:"visualSchema"`
	// @DataMember
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	// @DataMember
	Settings *SchemaSettingsDto `json:"settings,omitempty"`
	// @DataMember
	Embed *SchemaEmbedSettingsDto `json:"embed,omitempty"`
	// @DataMember
	Triggers []TriggerDto `json:"triggers,omitempty"`
}

type SchemaListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	SchemaName string `json:"schemaName"`
	// @DataMember
	SchemaTitle string `json:"schemaTitle"`
	// @DataMember
	LatestVersion *int `json:"latestVersion,omitempty"`
	// @DataMember
	HasDraft bool `json:"hasDraft,omitempty"`
	// @DataMember
	MetaSchemaVersion int `json:"metaSchemaVersion,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Env *string `json:"env,omitempty"`
}

// @DataContract
type PublicFolderDto struct {
	// @DataMember(Order=1)
	Path string `json:"path"`
	// @DataMember(Order=2)
	PublicId string `json:"publicId"`
	// @DataMember(Order=3)
	PublicUrl *string `json:"publicUrl,omitempty"`
	// @DataMember(Order=4)
	Inherited bool `json:"inherited,omitempty"`
}

type RegisterFileResult struct {
	Index       int     `json:"index,omitempty"`
	Path        *string `json:"path,omitempty"`
	OwnerUserId *string `json:"ownerUserId,omitempty"`
	Status      *string `json:"status,omitempty"`
	FileId      *string `json:"fileId,omitempty"`
	Verified    *bool   `json:"verified,omitempty"`
	Warning     *string `json:"warning,omitempty"`
	ErrorCode   *string `json:"errorCode,omitempty"`
	ErrorStatus *int    `json:"errorStatus,omitempty"`
	Error       *string `json:"error,omitempty"`
}

// @DataContract
type OwnPushDeviceDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	Token string `json:"token"`
	// @DataMember
	Platform *string `json:"platform,omitempty"`
	// @DataMember
	OsName *string `json:"osName,omitempty"`
	// @DataMember
	OsVersion *string `json:"osVersion,omitempty"`
	// @DataMember
	AppVersion *string `json:"appVersion,omitempty"`
	// @DataMember
	DeviceName *string `json:"deviceName,omitempty"`
	// @DataMember
	ModelName *string `json:"modelName,omitempty"`
	// @DataMember
	Brand *string `json:"brand,omitempty"`
	// @DataMember
	InstallationId *string `json:"installationId,omitempty"`
	// @DataMember
	Locale *string `json:"locale,omitempty"`
	// @DataMember
	TimeZone *string `json:"timeZone,omitempty"`
	// @DataMember
	UpdatedAtUtc *time.Time `json:"updatedAtUtc,omitempty"`
}

// @DataContract
type OwnPushMessageDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	CampaignId string `json:"campaignId"`
	// @DataMember
	Title *string `json:"title,omitempty"`
	// @DataMember
	Body *string `json:"body,omitempty"`
	// @DataMember
	SentAtUtc *time.Time `json:"sentAtUtc,omitempty"`
	// @DataMember
	IsRead bool `json:"isRead,omitempty"`
	// @DataMember
	ReadAtUtc *time.Time `json:"readAtUtc,omitempty"`
}

// @DataContract
type PaymentOrderLineDto struct {
	// @DataMember
	Sku string `json:"sku"`
	// @DataMember
	Title string `json:"title"`
	// @DataMember
	Quantity int `json:"quantity,omitempty"`
	// @DataMember
	UnitAmountMinor int64 `json:"unitAmountMinor,omitempty"`
	// @DataMember
	AmountMinor int64 `json:"amountMinor,omitempty"`
}

// @DataContract
type PaymentOrderCheckoutDto struct {
	// @DataMember
	Attempt int `json:"attempt,omitempty"`
	// @DataMember
	IntegrationId string `json:"integrationId"`
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	ProviderOrderId string `json:"providerOrderId"`
	// @DataMember
	ProviderReference string `json:"providerReference"`
	// @DataMember
	PaymentUrl string `json:"paymentUrl"`
	// @DataMember
	Method string `json:"method"`
	// @DataMember
	StartedOnUtc time.Time `json:"startedOnUtc,omitempty"`
	// @DataMember
	ExpiresAtUtc time.Time `json:"expiresAtUtc,omitempty"`
}

// @DataContract
type PaymentOrderDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	Env string `json:"env"`
	// @DataMember
	BuyerUserId string `json:"buyerUserId"`
	// @DataMember
	MerchantReference string `json:"merchantReference"`
	// @DataMember
	Currency string `json:"currency"`
	// @DataMember
	PriceVersion string `json:"priceVersion"`
	// @DataMember
	Lines []PaymentOrderLineDto `json:"lines"`
	// @DataMember
	TotalMinor int64 `json:"totalMinor,omitempty"`
	// @DataMember
	Metadata map[string]string `json:"metadata"`
	// @DataMember
	Status string `json:"status"`
	// @DataMember
	CreatedOnUtc time.Time `json:"createdOnUtc,omitempty"`
	// @DataMember
	UpdatedOnUtc time.Time `json:"updatedOnUtc,omitempty"`
	// @DataMember
	Checkouts []PaymentOrderCheckoutDto `json:"checkouts"`
	// @DataMember
	SettledMinor *int64 `json:"settledMinor,omitempty"`
	// @DataMember
	RefundedMinor int64 `json:"refundedMinor,omitempty"`
	// @DataMember
	PaidOnUtc *time.Time `json:"paidOnUtc,omitempty"`
	// @DataMember
	ProviderPaymentId *string `json:"providerPaymentId,omitempty"`
	// @DataMember
	Mismatches []string `json:"mismatches"`
}

// @DataContract
type WorkflowRunErrorDto struct {
	// @DataMember
	Code *string `json:"code,omitempty"`
	// @DataMember
	StepId *string `json:"stepId,omitempty"`
	// @DataMember
	Reason *string `json:"reason,omitempty"`
	// @DataMember
	Message *string `json:"message,omitempty"`
}

// @DataContract
type WorkflowRunActorDto struct {
	// @DataMember
	UserId *string `json:"userId,omitempty"`
	// @DataMember
	Kind string `json:"kind"`
}

// @DataContract
type WorkflowWaitSlotDto struct {
	/** @description As the author wrote it: a usr_ id, role:<name>, tag:<name>, role:<name>+tag:<name> or teamrole:<name>. Show this. */
	// @DataMember
	// @ApiMember(Description="As the author wrote it: a usr_ id, role:<name>, tag:<name>, role:<name>+tag:<name> or teamrole:<name>. Show this.")
	Label string `json:"label"`
	// @DataMember
	UserId *string `json:"userId,omitempty"`
	// @DataMember
	Role *string `json:"role,omitempty"`
	// @DataMember
	Tag *string `json:"tag,omitempty"`
	// @DataMember
	TeamRole *bool `json:"teamRole,omitempty"`
	/** @description Set when this slot is an AI agent's: the agent's id in the step. `userId` is then the agent's service user. */
	// @DataMember
	// @ApiMember(Description="Set when this slot is an AI agent's: the agent's id in the step. `userId` is then the agent's service user.")
	Agent *string `json:"agent,omitempty"`
	/** @description The reason given with the answer that filled this slot. An AI agent always gives one. */
	// @DataMember
	// @ApiMember(Description="The reason given with the answer that filled this slot. An AI agent always gives one.")
	Reason *string `json:"reason,omitempty"`
	// @DataMember
	AnsweredBy *string `json:"answeredBy,omitempty"`
	// @DataMember
	Decision *string `json:"decision,omitempty"`
	// @DataMember
	At *time.Time `json:"at,omitempty"`
	/** @description True = somebody who was not asked answered for this slot, using the override right. */
	// @DataMember
	// @ApiMember(Description="True = somebody who was not asked answered for this slot, using the override right.")
	Override *bool `json:"override,omitempty"`
}

// @DataContract
type WorkflowWaitAnswerDto struct {
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	Kind *string `json:"kind,omitempty"`
	// @DataMember
	Reason *string `json:"reason,omitempty"`
	// @DataMember
	Decision string `json:"decision"`
	/** @description Which slots this answer filled (places in `slots`, from 0). One person who matches two rules fills both with one answer. */
	// @DataMember
	// @ApiMember(Description="Which slots this answer filled (places in `slots`, from 0). One person who matches two rules fills both with one answer.")
	Slots []int `json:"slots,omitempty"`
	/** @description True = this person was not asked and answered on the override right. */
	// @DataMember
	// @ApiMember(Description="True = this person was not asked and answered on the override right.")
	Override *bool `json:"override,omitempty"`
	// @DataMember
	At time.Time `json:"at,omitempty"`
}

// @DataContract
type WorkflowWaitingOnDto struct {
	// @DataMember
	StepId string `json:"stepId"`
	// @DataMember
	Kind string `json:"kind"`
	// @DataMember
	Until *time.Time `json:"until,omitempty"`
	// @DataMember
	Key *string `json:"key,omitempty"`
	/** @description The people told when the wait parked. An approval addressed to a role tells everybody who held that role then, so this list can be longer or shorter than `slots` — count `slots`, never this. */
	// @DataMember
	// @ApiMember(Description="The people told when the wait parked. An approval addressed to a role tells everybody who held that role then, so this list can be longer or shorter than `slots` — count `slots`, never this.")
	Assignees []string `json:"assignees,omitempty"`
	/** @description A group approval: how many of `slots` must approve. Absent = the first answer decides. */
	// @DataMember
	// @ApiMember(Description="A group approval: how many of `slots` must approve. Absent = the first answer decides.")
	Require *int `json:"require,omitempty"`
	/** @description One entry per thing the approval asks for: a person or a rule, and whether it is filled. The unit `require`, `approvals` and `rejections` count. */
	// @DataMember
	// @ApiMember(Description="One entry per thing the approval asks for: a person or a rule, and whether it is filled. The unit `require`, `approvals` and `rejections` count.")
	Slots []WorkflowWaitSlotDto `json:"slots,omitempty"`
	// @DataMember
	Answers []WorkflowWaitAnswerDto `json:"answers,omitempty"`
	/** @description Inbox rows only: true = this approval is yours; false = none of it is yours, so answering it is an override, if you hold the right. Absent = nobody is asked for a slot here, or it was not worked out. */
	// @DataMember
	// @ApiMember(Description="Inbox rows only: true = this approval is yours; false = none of it is yours, so answering it is an override, if you hold the right. Absent = nobody is asked for a slot here, or it was not worked out.")
	Mine *bool `json:"mine,omitempty"`
	// @DataMember
	ChildRunId *string `json:"childRunId,omitempty"`
}

// @DataContract
type WorkflowRunStepDto struct {
	// @DataMember
	StepId string `json:"stepId"`
	// @DataMember
	Kind string `json:"kind"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	ForeachStepId *string `json:"foreachStepId,omitempty"`
	// @DataMember
	ItemIndex *int `json:"itemIndex,omitempty"`
	// @DataMember
	Status string `json:"status"`
	// @DataMember
	Attempt int `json:"attempt,omitempty"`
	// @DataMember
	Pass int `json:"pass,omitempty"`
	// @DataMember
	StartedAt *time.Time `json:"startedAt,omitempty"`
	// @DataMember
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// @DataMember
	DurationMs *int64 `json:"durationMs,omitempty"`
	// @DataMember
	Input *interface{} `json:"input,omitempty"`
	// @DataMember
	InputTruncated bool `json:"inputTruncated,omitempty"`
	// @DataMember
	Output *interface{} `json:"output,omitempty"`
	// @DataMember
	OutputTruncated bool `json:"outputTruncated,omitempty"`
	// @DataMember
	Error *WorkflowRunErrorDto `json:"error,omitempty"`
	// @DataMember
	ResumedBy *WorkflowRunActorDto `json:"resumedBy,omitempty"`
	// @DataMember
	Decision *string `json:"decision,omitempty"`
}

// @DataContract
type WorkflowRunDto struct {
	// @DataMember
	RunId string `json:"runId"`
	// @DataMember
	WorkflowId string `json:"workflowId"`
	// @DataMember
	Version *int `json:"version,omitempty"`
	// @DataMember
	IsDraft bool `json:"isDraft,omitempty"`
	// @DataMember
	Status string `json:"status"`
	// @DataMember
	Source string `json:"source"`
	// @DataMember
	StartedBy *WorkflowRunActorDto `json:"startedBy,omitempty"`
	// @DataMember
	TriggerId *string `json:"triggerId,omitempty"`
	// @DataMember
	TaskId *string `json:"taskId,omitempty"`
	// @DataMember
	ParentRunId *string `json:"parentRunId,omitempty"`
	// @DataMember
	Depth int `json:"depth,omitempty"`
	// @DataMember
	Env string `json:"env"`
	// @DataMember
	CorrelationId *string `json:"correlationId,omitempty"`
	// @DataMember
	ExecutionIdentityId *string `json:"executionIdentityId,omitempty"`
	// @DataMember
	ExecutionIdentityName *string `json:"executionIdentityName,omitempty"`
	// @DataMember
	ExecutionIdentityIsDefault bool `json:"executionIdentityIsDefault,omitempty"`
	// @DataMember
	StartedAt time.Time `json:"startedAt,omitempty"`
	// @DataMember
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// @DataMember
	DurationMs *int64 `json:"durationMs,omitempty"`
	// @DataMember
	WaitingOn []WorkflowWaitingOnDto `json:"waitingOn"`
	// @DataMember
	Error *WorkflowRunErrorDto `json:"error,omitempty"`
	// @DataMember
	Output *interface{} `json:"output,omitempty"`
	// @DataMember
	StepRuns int `json:"stepRuns,omitempty"`
	// @DataMember
	Steps []WorkflowRunStepDto `json:"steps"`
	// @DataMember
	Request map[string]interface{} `json:"request,omitempty"`
	// @DataMember
	Trigger map[string]interface{} `json:"trigger,omitempty"`
}

// @DataContract
type WorkflowRunListProjection struct {
	// @DataMember
	RunId string `json:"runId"`
	// @DataMember
	WorkflowId string `json:"workflowId"`
	// @DataMember
	Version *int `json:"version,omitempty"`
	// @DataMember
	IsDraft bool `json:"isDraft,omitempty"`
	// @DataMember
	Status string `json:"status"`
	// @DataMember
	Source string `json:"source"`
	// @DataMember
	StartedAt time.Time `json:"startedAt,omitempty"`
	// @DataMember
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// @DataMember
	DurationMs *int64 `json:"durationMs,omitempty"`
	// @DataMember
	StartedBy *WorkflowRunActorDto `json:"startedBy,omitempty"`
	// @DataMember
	ExecutionIdentityId *string `json:"executionIdentityId,omitempty"`
	// @DataMember
	ExecutionIdentityName *string `json:"executionIdentityName,omitempty"`
	// @DataMember
	ExecutionIdentityIsDefault bool `json:"executionIdentityIsDefault,omitempty"`
	// @DataMember
	CurrentStepIds []string `json:"currentStepIds"`
	// @DataMember
	WaitingOn []WorkflowWaitingOnDto `json:"waitingOn"`
	// @DataMember
	Error *WorkflowRunErrorDto `json:"error,omitempty"`
}

// @DataContract
type TenantLogEntryDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	Timestamp time.Time `json:"timestamp,omitempty"`
	// @DataMember
	Module string `json:"module"`
	// @DataMember
	Level string `json:"level"`
	// @DataMember
	EventCode string `json:"eventCode"`
	// @DataMember
	Title string `json:"title"`
	// @DataMember
	Message string `json:"message"`
	// @DataMember
	CorrelationId *string `json:"correlationId,omitempty"`
	// @DataMember
	TraceId *string `json:"traceId,omitempty"`
	// @DataMember
	SpanId *string `json:"spanId,omitempty"`
	// @DataMember
	Env string `json:"env"`
	// @DataMember
	Origin string `json:"origin"`
	// @DataMember
	Meta *IReadOnlyDictionary[string, string] `json:"meta,omitempty"`
}

type InAppNotificationDto struct {
	Id           string     `json:"id"`
	ProjectId    string     `json:"projectId"`
	UserId       string     `json:"userId"`
	Channel      string     `json:"channel"`
	EventName    string     `json:"eventName"`
	PayloadJson  *string    `json:"payloadJson,omitempty"`
	IsRead       bool       `json:"isRead,omitempty"`
	SeenAtUtc    *time.Time `json:"seenAtUtc,omitempty"`
	ReadAtUtc    *time.Time `json:"readAtUtc,omitempty"`
	CreatedAtUtc time.Time  `json:"createdAtUtc,omitempty"`
}

// @DataContract
type ChannelSubscriptionStateDto struct {
	// @DataMember
	Unsubscribed bool `json:"unsubscribed,omitempty"`
	// @DataMember
	BlockedTags map[string][]string `json:"blockedTags"`
	// @DataMember
	BounceState string `json:"bounceState"`
	// @DataMember
	LastBounceAt *time.Time `json:"lastBounceAt,omitempty"`
	// @DataMember
	SpamComplaint bool `json:"spamComplaint,omitempty"`
}

// @DataContract
type UserDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	ProjectId string `json:"projectId"`
	// @DataMember
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	// @DataMember
	PrimaryPhone *string `json:"primaryPhone,omitempty"`
	// @DataMember
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember
	FirstName *string `json:"firstName,omitempty"`
	// @DataMember
	LastName *string `json:"lastName,omitempty"`
	// @DataMember
	FullName *string `json:"fullName,omitempty"`
	// @DataMember
	Company *string `json:"company,omitempty"`
	// @DataMember
	Locale *string `json:"locale,omitempty"`
	// @DataMember
	TimeZone *string `json:"timeZone,omitempty"`
	// @DataMember
	Gender *string `json:"gender,omitempty"`
	// @DataMember
	BirthDate *int64 `json:"birthDate,omitempty"`
	// @DataMember
	AddressLine1 *string `json:"addressLine1,omitempty"`
	// @DataMember
	AddressLine2 *string `json:"addressLine2,omitempty"`
	// @DataMember
	Country *string `json:"country,omitempty"`
	// @DataMember
	City *string `json:"city,omitempty"`
	// @DataMember
	State *string `json:"state,omitempty"`
	// @DataMember
	PostalCode *string `json:"postalCode,omitempty"`
	// @DataMember
	Tags []string `json:"tags,omitempty"`
	// @DataMember
	Roles []string `json:"roles,omitempty"`
	// @DataMember
	Lifecycle string `json:"lifecycle"`
	// @DataMember
	SourceOfCreation string `json:"sourceOfCreation"`
	// @DataMember
	MergedIntoContactId *string `json:"mergedIntoContactId,omitempty"`
	// @DataMember
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// @DataMember
	ModifiedOn time.Time `json:"modifiedOn,omitempty"`
	// @DataMember
	Auths []AuthDto `json:"auths,omitempty"`
	// @DataMember
	MarketingPreferences map[string]ChannelSubscriptionStateDto `json:"marketingPreferences,omitempty"`
}

// @DataContract
type ContactRolesRepairDto struct {
	// @DataMember(Name="dryRun")
	DryRun bool `json:"dryRun,omitempty"`
	// @DataMember(Name="contactsWithoutRoles")
	ContactsWithoutRoles int `json:"contactsWithoutRoles,omitempty"`
	// @DataMember(Name="contactsToRepair")
	ContactsToRepair int `json:"contactsToRepair,omitempty"`
	// @DataMember(Name="contactsRepaired")
	ContactsRepaired int `json:"contactsRepaired,omitempty"`
	// @DataMember(Name="contactsWithoutLoginRoleIds")
	ContactsWithoutLoginRoleIds int `json:"contactsWithoutLoginRoleIds,omitempty"`
	// @DataMember(Name="legacyRoleNamesIgnored")
	LegacyRoleNamesIgnored int `json:"legacyRoleNamesIgnored,omitempty"`
}

// @DataContract
type SchemaRefDto struct {
	// @DataMember(Order=1)
	SchemaId string `json:"schemaId"`
	// @DataMember(Order=2)
	SchemaName string `json:"schemaName"`
	// @DataMember(Order=3)
	DatabaseIntegrationId string `json:"databaseIntegrationId"`
}

type CollectionImportDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	Schema SchemaRefDto `json:"schema"`
	// @DataMember
	File FileResourceRefDto `json:"file"`
	// @DataMember
	ErrorFile *FileResourceRefDto `json:"errorFile,omitempty"`
	// @DataMember
	Delimiter string `json:"delimiter"`
	// @DataMember
	HasHeader bool `json:"hasHeader,omitempty"`
	// @DataMember
	Status string `json:"status"`
	// @DataMember
	TotalRows int64 `json:"totalRows,omitempty"`
	// @DataMember
	TotalImported int64 `json:"totalImported,omitempty"`
	// @DataMember
	TotalErrors int64 `json:"totalErrors,omitempty"`
	// @DataMember
	FailureReason *string `json:"failureReason,omitempty"`
	// @DataMember
	Mapping []ImportColumnMappingDto `json:"mapping,omitempty"`
	// @DataMember
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// @DataMember
	StartedOn *time.Time `json:"startedOn,omitempty"`
	// @DataMember
	CompletedOn *time.Time `json:"completedOn,omitempty"`
}

type ImportUploadTargetDto struct {
	// @DataMember
	Url string `json:"url"`
	// @DataMember
	ContentType string `json:"contentType"`
	// @DataMember
	File FileResourceRefDto `json:"file"`
}

type ImportFileColumnDto struct {
	// @DataMember
	Index int `json:"index,omitempty"`
	// @DataMember
	Header string `json:"header"`
	// @DataMember
	Samples []string `json:"samples"`
	// @DataMember
	DetectedType string `json:"detectedType"`
}

type ImportFileAnalysisDto struct {
	// @DataMember
	File FileResourceRefDto `json:"file"`
	// @DataMember
	Columns []ImportFileColumnDto `json:"columns"`
	// @DataMember
	SampleRowCount int `json:"sampleRowCount,omitempty"`
}

type CampaignStatus string

const (
	CampaignStatusPending    CampaignStatus = "Pending"
	CampaignStatusRegistered                = "Registered"
	CampaignStatusScheduled                 = "Scheduled"
	CampaignStatusStarted                   = "Started"
	CampaignStatusStopped                   = "Stopped"
	CampaignStatusProcessing                = "Processing"
	CampaignStatusCompleted                 = "Completed"
	CampaignStatusFailed                    = "Failed"
)

type CampaignStatusChangeEntryDto struct {
	Time   time.Time      `json:"time,omitempty"`
	Status CampaignStatus `json:"status,omitempty"`
	Errors []ErrorDto     `json:"errors,omitempty"`
}

// @DataContract
type CampaignOutcomeDto struct {
	// @DataMember
	Sent int `json:"sent,omitempty"`
	// @DataMember
	Failed int `json:"failed,omitempty"`
	// @DataMember
	Blocked int `json:"blocked,omitempty"`
	// @DataMember
	Cancelled int `json:"cancelled,omitempty"`
	// @DataMember
	RemovedRecipients int `json:"removedRecipients,omitempty"`
	// @DataMember
	BlockedReason *string `json:"blockedReason,omitempty"`
}

// @DataContract
type CampaignDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	ForceCampaignLanguage bool `json:"forceCampaignLanguage,omitempty"`
	// @DataMember
	CampaignProcessingIntegrationId *string `json:"campaignProcessingIntegrationId,omitempty"`
	// @DataMember
	StatusHistory []CampaignStatusChangeEntryDto `json:"statusHistory"`
	// @DataMember
	Status *CampaignStatusChangeEntryDto `json:"status,omitempty"`
	// @DataMember
	TokenMappingValues []TokenMappingDto `json:"tokenMappingValues,omitempty"`
	// @DataMember
	Notes *string `json:"notes,omitempty"`
	// @DataMember
	CreatedById *string `json:"createdById,omitempty"`
	// @DataMember
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	TriggerId *string `json:"triggerId,omitempty"`
	// @DataMember
	TriggerType *TriggerType `json:"triggerType,omitempty"`
	// @DataMember
	Outcome *CampaignOutcomeDto `json:"outcome,omitempty"`
	// @DataMember
	Id *string `json:"id,omitempty"`
}

type EmailCampaignDeliverySettingsDto struct {
	// @DataMember
	RecipientsSourceType EmailCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
	// @DataMember
	CampaignTime *int64 `json:"campaignTime,omitempty"`
	// @DataMember
	RespectTimeZoneSettings *RespectTimeZoneSettings `json:"respectTimeZoneSettings,omitempty"`
}

type CommunicationChannel string

const (
	CommunicationChannelTransactional CommunicationChannel = "Transactional"
	CommunicationChannelMarketing                          = "Marketing"
	CommunicationChannelSystem                             = "System"
)

type NotificationMedium string

const (
	NotificationMediumEmail NotificationMedium = "Email"
	NotificationMediumSms                      = "Sms"
	NotificationMediumPush                     = "Push"
)

// @DataContract
type TemplateDto struct {
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	TemplateName string `json:"templateName"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	// @DataMember
	Medium NotificationMedium `json:"medium,omitempty"`
	// @DataMember
	IsActive bool `json:"isActive,omitempty"`
	// @DataMember
	Tags []string `json:"tags,omitempty"`
	// @DataMember
	Env *string `json:"env,omitempty"`
}

type EmailTemplateEngine string

const (
	EmailTemplateEngineNotSet     EmailTemplateEngine = "NotSet"
	EmailTemplateEngineHandlebars                     = "Handlebars"
	EmailTemplateEngineMjml                           = "Mjml"
	EmailTemplateEngineLiquid                         = "Liquid"
	EmailTemplateEngineRazor                          = "Razor"
	EmailTemplateEngineMustache                       = "Mustache"
)

// @DataContract
type EmailBodyDto struct {
	// @DataMember
	Structure *string `json:"structure,omitempty"`
	// @DataMember
	Code string `json:"code"`
	// @DataMember
	TemplateEngine EmailTemplateEngine `json:"templateEngine,omitempty"`
}

// @DataContract
type EmailMessageContentDto struct {
	// @DataMember
	Subject string `json:"subject"`
	// @DataMember
	Body EmailBodyDto `json:"body"`
	// @DataMember
	StaticAttachments []FileResourceRefDto `json:"staticAttachments,omitempty"`
}

// @DataContract
type EmailMessageTranslationDto struct {
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	Content EmailMessageContentDto `json:"content"`
	// @DataMember
	StaticAttachments []FileResourceRefDto `json:"staticAttachments,omitempty"`
}

// @DataContract
type EmailTemplateDto struct {
	TemplateDto
	// @DataMember
	Translations []EmailMessageTranslationDto `json:"translations"`
	// @DataMember
	StaticAttachments []FileResourceRefDto `json:"staticAttachments,omitempty"`
}

// @DataContract
type EmailCampaignDto struct {
	CampaignDto
	// @DataMember
	DeliverySettings EmailCampaignDeliverySettingsDto `json:"deliverySettings"`
	// @DataMember
	Template EmailTemplateDto `json:"template"`
	// @DataMember
	ValidationIntegrationId *string `json:"validationIntegrationId,omitempty"`
	// @DataMember
	TemplateIsSystem bool `json:"templateIsSystem,omitempty"`
}

// @DataContract
type EmailCampaignListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	TemplateName string `json:"templateName"`
	// @DataMember
	TemplateId string `json:"templateId"`
	// @DataMember
	IntegrationId *string `json:"integrationId,omitempty"`
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	Strategy string `json:"strategy"`
	// @DataMember
	LatestStatus *CampaignStatus `json:"latestStatus,omitempty"`
	// @DataMember
	CreatedOn time.Time `json:"createdOn,omitempty"`
	// @DataMember
	Outcome *CampaignOutcomeDto `json:"outcome,omitempty"`
}

type CampaignBatchStatus string

const (
	CampaignBatchStatusRegistered CampaignBatchStatus = "Registered"
	CampaignBatchStatusProcessing                     = "Processing"
	CampaignBatchStatusCompleted                      = "Completed"
	CampaignBatchStatusFailed                         = "Failed"
)

type BatchStatusChangeEntryDto struct {
	Time   time.Time           `json:"time,omitempty"`
	Status CampaignBatchStatus `json:"status,omitempty"`
	Errors []ErrorDto          `json:"errors,omitempty"`
}

// @DataContract
type CampaignBatchDto struct {
	// @DataMember
	CampaignId string `json:"campaignId"`
	// @DataMember
	BatchId string `json:"batchId"`
	// @DataMember
	StartAfter *string `json:"startAfter,omitempty"`
	// @DataMember
	StatusHistory []BatchStatusChangeEntryDto `json:"statusHistory"`
	// @DataMember
	Id string `json:"id"`
}

// @DataContract
type EmailRecipientDto struct {
	// @DataMember
	EmailAddress string `json:"emailAddress"`
	// @DataMember
	Language *string `json:"language,omitempty"`
	// @DataMember
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	// @DataMember
	UserTokenMappings []TokenMappingDto `json:"userTokenMappings,omitempty"`
}

// @DataContract
type EmailRecipientsDto struct {
	// @DataMember
	To []EmailRecipientDto `json:"to,omitempty"`
	// @DataMember
	Cc []EmailRecipientDto `json:"cc,omitempty"`
	// @DataMember
	Bcc []EmailRecipientDto `json:"bcc,omitempty"`
	// @DataMember
	StartingAfter *string `json:"startingAfter,omitempty"`
	// @DataMember
	HasMore bool `json:"hasMore,omitempty"`
}

// @DataContract
type EmailCampaignBatchDto struct {
	CampaignBatchDto
	// @DataMember
	Recipients *EmailRecipientsDto `json:"recipients,omitempty"`
}

type CampaignNotificationStatus string

const (
	CampaignNotificationStatusCompleted                         CampaignNotificationStatus = "Completed"
	CampaignNotificationStatusBlockedByUserPreferenceBlockAll                              = "BlockedByUserPreferenceBlockAll"
	CampaignNotificationStatusBlockedByUserPreferenceBlockByTag                            = "BlockedByUserPreferenceBlockByTag"
	CampaignNotificationStatusFailed                                                       = "Failed"
	CampaignNotificationStatusViewed                                                       = "Viewed"
	CampaignNotificationStatusClicked                                                      = "Clicked"
	CampaignNotificationStatusBlockedByValidation                                          = "BlockedByValidation"
)

type NotificationStatusChangeEntryDto struct {
	Time       time.Time                  `json:"time,omitempty"`
	Status     CampaignNotificationStatus `json:"status,omitempty"`
	SourceId   *string                    `json:"sourceId,omitempty"`
	Errors     []ErrorDto                 `json:"errors,omitempty"`
	Tags       []string                   `json:"tags,omitempty"`
	Flag       *string                    `json:"flag,omitempty"`
	FlagReason *string                    `json:"flagReason,omitempty"`
}

type NotificationDeliveryState string

const (
	NotificationDeliveryStatePending   NotificationDeliveryState = "Pending"
	NotificationDeliveryStateSending                             = "Sending"
	NotificationDeliveryStateRetrying                            = "Retrying"
	NotificationDeliveryStateAccepted                            = "Accepted"
	NotificationDeliveryStateFailed                              = "Failed"
	NotificationDeliveryStateBlocked                             = "Blocked"
	NotificationDeliveryStateCancelled                           = "Cancelled"
)

type NotificationAttemptOutcome string

const (
	NotificationAttemptOutcomeAccepted NotificationAttemptOutcome = "Accepted"
	NotificationAttemptOutcomeFailed                              = "Failed"
	NotificationAttemptOutcomeBlocked                             = "Blocked"
)

type SendErrorCategory string

const (
	SendErrorCategoryUnknown   SendErrorCategory = "Unknown"
	SendErrorCategoryTransient                   = "Transient"
	SendErrorCategoryPermanent                   = "Permanent"
	SendErrorCategoryAmbiguous                   = "Ambiguous"
)

type NotificationAttemptTrigger string

const (
	NotificationAttemptTriggerInitial     NotificationAttemptTrigger = "Initial"
	NotificationAttemptTriggerAutoRetry                              = "AutoRetry"
	NotificationAttemptTriggerManualRetry                            = "ManualRetry"
)

// @DataContract
type NotificationAttemptDto struct {
	// @DataMember
	AttemptId string `json:"attemptId"`
	// @DataMember
	Number int `json:"number,omitempty"`
	// @DataMember
	StartedAt time.Time `json:"startedAt,omitempty"`
	// @DataMember
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	// @DataMember
	Outcome NotificationAttemptOutcome `json:"outcome,omitempty"`
	// @DataMember
	Category *SendErrorCategory `json:"category,omitempty"`
	// @DataMember
	StatusCode *int `json:"statusCode,omitempty"`
	// @DataMember
	ErrorCode *string `json:"errorCode,omitempty"`
	// @DataMember
	Message *string `json:"message,omitempty"`
	// @DataMember
	ProviderRef *string `json:"providerRef,omitempty"`
	// @DataMember
	RetryAfterSeconds *int `json:"retryAfterSeconds,omitempty"`
	// @DataMember
	Exhausted *bool `json:"exhausted,omitempty"`
	// @DataMember
	NextRetryAt *time.Time `json:"nextRetryAt,omitempty"`
	// @DataMember
	CorrelationId *string `json:"correlationId,omitempty"`
	// @DataMember
	Trigger NotificationAttemptTrigger `json:"trigger,omitempty"`
	// @DataMember
	RetryRequestId *string `json:"retryRequestId,omitempty"`
}

type DeliveryEventOutcome string

const (
	DeliveryEventOutcomeDelivered  DeliveryEventOutcome = "Delivered"
	DeliveryEventOutcomeBounced                         = "Bounced"
	DeliveryEventOutcomeComplained                      = "Complained"
	DeliveryEventOutcomeRejected                        = "Rejected"
	DeliveryEventOutcomeDeferred                        = "Deferred"
)

// @DataContract
type NotificationDeliveryEventDto struct {
	// @DataMember
	EventId *string `json:"eventId,omitempty"`
	// @DataMember
	Outcome DeliveryEventOutcome `json:"outcome,omitempty"`
	// @DataMember
	At time.Time `json:"at,omitempty"`
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	ProviderRef *string `json:"providerRef,omitempty"`
	// @DataMember
	Reason *string `json:"reason,omitempty"`
}

// @DataContract
type CampaignBatchNotificationDto struct {
	// @DataMember
	CampaignId string `json:"campaignId"`
	// @DataMember
	BatchId string `json:"batchId"`
	// @DataMember
	NotificationId string `json:"notificationId"`
	// @DataMember
	RefNotificationId *string `json:"refNotificationId,omitempty"`
	// @DataMember
	Body *string `json:"body,omitempty"`
	// @DataMember
	Model map[string]string `json:"model,omitempty"`
	// @DataMember
	StatusHistory []NotificationStatusChangeEntryDto `json:"statusHistory"`
	// @DataMember
	State *NotificationDeliveryState `json:"state,omitempty"`
	// @DataMember
	StateChangedAt *time.Time `json:"stateChangedAt,omitempty"`
	// @DataMember
	Attempts []NotificationAttemptDto `json:"attempts,omitempty"`
	// @DataMember
	DeliveryEvents []NotificationDeliveryEventDto `json:"deliveryEvents,omitempty"`
	// @DataMember
	DeliveryOutcome *DeliveryEventOutcome `json:"deliveryOutcome,omitempty"`
	// @DataMember
	DeliveryOutcomeRank *int `json:"deliveryOutcomeRank,omitempty"`
	// @DataMember
	Id string `json:"id"`
}

type EmailRecipientField string

const (
	EmailRecipientFieldNotSet EmailRecipientField = "NotSet"
	EmailRecipientFieldTo                         = "To"
	EmailRecipientFieldCc                         = "Cc"
	EmailRecipientFieldBcc                        = "Bcc"
)

type RecipientRemovalReason string

const (
	RecipientRemovalReasonNotSet        RecipientRemovalReason = "NotSet"
	RecipientRemovalReasonOptedOut                             = "OptedOut"
	RecipientRemovalReasonCompliance                           = "Compliance"
	RecipientRemovalReasonUndeliverable                        = "Undeliverable"
)

// @DataContract
type RemovedRecipientDto struct {
	// @DataMember
	Address string `json:"address"`
	// @DataMember
	Field EmailRecipientField `json:"field,omitempty"`
	// @DataMember
	Reason RecipientRemovalReason `json:"reason,omitempty"`
	// @DataMember
	Tag *string `json:"tag,omitempty"`
}

// @DataContract
type EmailCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	// @DataMember
	Subject *string `json:"subject,omitempty"`
	// @DataMember
	Recipients EmailRecipientsDto `json:"recipients"`
	// @DataMember
	RemovedRecipients []RemovedRecipientDto `json:"removedRecipients,omitempty"`
	// @DataMember
	Content *EmailMessageContentDto `json:"content,omitempty"`
	// @DataMember
	TransactionalKind *string `json:"transactionalKind,omitempty"`
}

// @DataContract
type CampaignStatsDto struct {
	// @DataMember
	Batches int `json:"batches,omitempty"`
	// @DataMember
	Sent int `json:"sent,omitempty"`
	// @DataMember
	Failed int `json:"failed,omitempty"`
	// @DataMember
	FailedRetryable int `json:"failedRetryable,omitempty"`
	// @DataMember
	Cancelled int `json:"cancelled,omitempty"`
	// @DataMember
	Blocked int `json:"blocked,omitempty"`
	// @DataMember
	InProgress int `json:"inProgress,omitempty"`
	// @DataMember
	RemovedRecipients int `json:"removedRecipients,omitempty"`
	// @DataMember
	Delivered int `json:"delivered,omitempty"`
	// @DataMember
	Bounced int `json:"bounced,omitempty"`
	// @DataMember
	Complained int `json:"complained,omitempty"`
	// @DataMember
	Rejected int `json:"rejected,omitempty"`
	// @DataMember
	SuccessRate float64 `json:"successRate,omitempty"`
}

// @DataContract
type SmsMessageContentDto struct {
	// @DataMember
	Body string `json:"body"`
}

// @DataContract
type SmsMessageTranslationDto struct {
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	Content SmsMessageContentDto `json:"content"`
}

// @DataContract
type SmsTemplateDto struct {
	TemplateDto
	// @DataMember
	Translations []SmsMessageTranslationDto `json:"translations"`
}

// @DataContract
type SmsCampaignDto struct {
	CampaignDto
	// @DataMember
	Recipients SmsCampaignDeliverySettingsDto `json:"recipients"`
	// @DataMember
	Template SmsTemplateDto `json:"template"`
}

// @DataContract
type SmsRecipientDto struct {
	// @DataMember
	PhoneNumber string `json:"phoneNumber"`
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	Language *string `json:"language,omitempty"`
	// @DataMember
	UserTokenMappings []TokenMappingDto `json:"userTokenMappings,omitempty"`
	// @DataMember
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	// @DataMember
	Record *string `json:"record,omitempty"`
}

// @DataContract
type SmsRecipientsDto struct {
	// @DataMember
	To []SmsRecipientDto `json:"to,omitempty"`
	// @DataMember
	StartingAfter *string `json:"startingAfter,omitempty"`
	// @DataMember
	HasMore bool `json:"hasMore,omitempty"`
}

// @DataContract
type SmsCampaignBatchDto struct {
	CampaignBatchDto
	// @DataMember
	Recipients SmsRecipientsDto `json:"recipients"`
}

// @DataContract
type SmsCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	// @DataMember
	Recipients SmsRecipientsDto `json:"recipients"`
	// @DataMember
	Content *SmsMessageContentDto `json:"content,omitempty"`
}

// @DataContract
type PushAudienceCountDto struct {
	// @DataMember
	Devices int `json:"devices,omitempty"`
	// @DataMember
	Recipients int `json:"recipients,omitempty"`
	// @DataMember
	SkippedUserIds int `json:"skippedUserIds,omitempty"`
	// @DataMember
	IsCapped bool `json:"isCapped,omitempty"`
}

// @DataContract
type PushCampaignDeliverySettingsDto struct {
	// @DataMember
	RecipientsSourceType PushCampaignRecipientsSourceTypes `json:"recipientsSourceType,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
	// @DataMember
	CampaignTime *int64 `json:"campaignTime,omitempty"`
	// @DataMember
	RespectTimeZoneSettings *RespectTimeZoneSettings `json:"respectTimeZoneSettings,omitempty"`
}

// @DataContract
type PushMessageContentDto struct {
	// @DataMember
	Title string `json:"title"`
	// @DataMember
	Body string `json:"body"`
}

// @DataContract
type PushMessageTranslationDto struct {
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	Content PushMessageContentDto `json:"content"`
}

// @DataContract
type PushTemplateDto struct {
	TemplateDto
	// @DataMember
	Translations []PushMessageTranslationDto `json:"translations"`
}

// @DataContract
type PushCampaignDto struct {
	CampaignDto
	// @DataMember
	Recipients PushCampaignDeliverySettingsDto `json:"recipients"`
	// @DataMember
	Template PushTemplateDto `json:"template"`
}

type PushDeviceDeliveryFamily string

const (
	PushDeviceDeliveryFamilyIos     PushDeviceDeliveryFamily = "Ios"
	PushDeviceDeliveryFamilyAndroid                          = "Android"
	PushDeviceDeliveryFamilyChrome                           = "Chrome"
	PushDeviceDeliveryFamilySafari                           = "Safari"
	PushDeviceDeliveryFamilyExpo                             = "Expo"
)

// @DataContract
type PushDeviceDeliveryTokenDto struct {
	// @DataMember
	PushDeviceToken string `json:"pushDeviceToken"`
	// @DataMember
	DeliveryFamily PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
}

// @DataContract
type PushRecipientDto struct {
	// @DataMember
	DeviceTokens []PushDeviceDeliveryTokenDto `json:"deviceTokens"`
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	Language *string `json:"language,omitempty"`
	// @DataMember
	UserTokenMappings []TokenMappingDto `json:"userTokenMappings,omitempty"`
	// @DataMember
	TimeZoneId *string `json:"timeZoneId,omitempty"`
	// @DataMember
	Record *string `json:"record,omitempty"`
}

// @DataContract
type PushRecipientsDto struct {
	// @DataMember
	To []PushRecipientDto `json:"to,omitempty"`
	// @DataMember
	StartingAfter *string `json:"startingAfter,omitempty"`
	// @DataMember
	HasMore bool `json:"hasMore,omitempty"`
}

// @DataContract
type PushCampaignBatchDto struct {
	CampaignBatchDto
	// @DataMember
	Recipients PushRecipientsDto `json:"recipients"`
}

// @DataContract
type NotificationReadMarkDto struct {
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	ReadAtUtc time.Time `json:"readAtUtc,omitempty"`
}

// @DataContract
type PushCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	// @DataMember
	Subject *string `json:"subject,omitempty"`
	// @DataMember
	Recipients PushRecipientsDto `json:"recipients"`
	// @DataMember
	Content *PushMessageContentDto `json:"content,omitempty"`
	// @DataMember
	ReadBy []NotificationReadMarkDto `json:"readBy,omitempty"`
}

type EchoResponse struct {
	Release                     CodeMashRelease         `json:"release,omitempty"`
	Runtime                     CodeMashRuntime         `json:"runtime,omitempty"`
	ManagedServiceHubUrl        string                  `json:"managedServiceHubUrl"`
	ManagedServiceApiUrl        string                  `json:"managedServiceApiUrl"`
	HubUrl                      string                  `json:"hubUrl"`
	ApiUrl                      string                  `json:"apiUrl"`
	ApiVersion                  string                  `json:"apiVersion"`
	HubVersion                  string                  `json:"hubVersion"`
	ApiUrlTemplate              *string                 `json:"apiUrlTemplate,omitempty"`
	EmailServiceConfigured      bool                    `json:"emailServiceConfigured,omitempty"`
	RootBootstrapPasswordSource *string                 `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                     []EchoRegionDto         `json:"regions,omitempty"`
	Agent                       *EchoAgentDto           `json:"agent,omitempty"`
	HostingPlatform             *EchoHostingPlatformDto `json:"hostingPlatform,omitempty"`
	Connect                     *EchoConnectDto         `json:"connect,omitempty"`
	Keyless                     *EchoKeylessDto         `json:"keyless,omitempty"`
}

type PublicProjectConfigDto struct {
	DisplayName string          `json:"displayName"`
	Branding    *PublicBrandDto `json:"branding,omitempty"`
	Auth        PublicAuthDto   `json:"auth"`
	AiChat      PublicAiChatDto `json:"aiChat"`
}

type PublicLegalDocumentDto struct {
	Kind      string  `json:"kind"`
	Title     *string `json:"title,omitempty"`
	Body      string  `json:"body"`
	Available bool    `json:"available,omitempty"`
}

// @DataContract
type IdResponse struct {
	ResponseBase
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	Status *string `json:"status,omitempty"`
	// @DataMember
	Warnings []string `json:"warnings,omitempty"`
}

type ListEndUserChatAttachmentsResponse struct {
	ResponseBase
	Attachments []EndUserChatAttachment `json:"attachments"`
}

type EmptyResponse struct {
	ResponseBase
}

type ListEndUserChatMemoryResponse struct {
	ResponseBase
	Notes []EndUserChatMemoryNote `json:"notes"`
}

type GetEndUserChatAvailabilityResponse struct {
	ResponseBase
	Enabled            bool                   `json:"enabled,omitempty"`
	Available          bool                   `json:"available,omitempty"`
	Reason             *string                `json:"reason,omitempty"`
	DefaultAssistantId *string                `json:"defaultAssistantId,omitempty"`
	Assistants         []EndUserChatAssistant `json:"assistants"`
	Plan               *EndUserChatPlan       `json:"plan,omitempty"`
}

type ListEndUserChatSessionsResponse struct {
	ResponseBase
	Sessions []EndUserChatSession `json:"sessions"`
}

type GetEndUserChatSessionResponse struct {
	ResponseBase
	Session *EndUserChatSession `json:"session,omitempty"`
}

type GetEndUserChatEntriesResponse struct {
	ResponseBase
	SessionId *string              `json:"sessionId,omitempty"`
	Entries   []AiChatEntryWireDto `json:"entries"`
	LastSeq   int64                `json:"lastSeq,omitempty"`
	HasMore   bool                 `json:"hasMore,omitempty"`
}

type StartEndUserChatTurnResponse struct {
	ResponseBase
	TurnId    *string `json:"turnId,omitempty"`
	SessionId *string `json:"sessionId,omitempty"`
	Channel   *string `json:"channel,omitempty"`
}

type GetEndUserAiToolsResponse struct {
	ResponseBase
	Tools []EndUserAiTool `json:"tools,omitempty"`
}

type InvokeEndUserAiToolResponse struct {
	ResponseBase
	Result *string `json:"result,omitempty"`
}

// @DataContract
type SaveGuestUserResponse struct {
	IdResponse
	// @DataMember
	UserName *string `json:"userName,omitempty"`
	// @DataMember
	Password *string `json:"password,omitempty"`
}

type GetUserResponse struct {
	ResponseBase
	User *AuthDto `json:"user,omitempty"`
}

type GetUsersResponse struct {
	ResponseBase
	List *PaginatedResponse[AuthDto] `json:"list,omitempty"`
}

type GetUserPreferencesResponse struct {
	ResponseBase
	Preferences *UserMarketingPreferencesDto `json:"preferences,omitempty"`
}

// @DataContract
type StartLinkIdentityResponse struct {
	ResponseBase
	// @DataMember
	AuthorizeUrl *string `json:"authorizeUrl,omitempty"`
}

// @DataContract
type SocialSignInStartResponse struct {
	ResponseBase
}

// @DataContract
type SocialSignInTokenResponse struct {
	ResponseBase
	// @DataMember
	BearerToken *string `json:"bearerToken,omitempty"`
	// @DataMember
	UserId *string `json:"userId,omitempty"`
	// @DataMember
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	// @DataMember
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type CredentialChangedResponse struct {
	ResponseBase
	// @DataMember
	AuthId *string `json:"authId,omitempty"`
	// @DataMember
	CredentialVersion *string `json:"credentialVersion,omitempty"`
}

type PasskeyOkResponse struct {
	ResponseBase
}

// @DataContract
type SendTestPasswordResetResponse struct {
	ResponseBase
	// @DataMember
	Queued bool `json:"queued,omitempty"`
	// @DataMember
	Reason *string `json:"reason,omitempty"`
	// @DataMember
	Env string `json:"env"`
	// @DataMember
	CorrelationId *string `json:"correlationId,omitempty"`
	// @DataMember
	ExpiresAtUtc *time.Time `json:"expiresAtUtc,omitempty"`
	// @DataMember
	IntegrationId *string `json:"integrationId,omitempty"`
	// @DataMember
	IntegrationName *string `json:"integrationName,omitempty"`
	// @DataMember
	Provider *string `json:"provider,omitempty"`
	// @DataMember
	IntegrationProblem *string `json:"integrationProblem,omitempty"`
	// @DataMember
	TemplateId *string `json:"templateId,omitempty"`
	// @DataMember
	TemplateName *string `json:"templateName,omitempty"`
	// @DataMember
	TemplateSource *string `json:"templateSource,omitempty"`
	// @DataMember
	Language *string `json:"language,omitempty"`
	// @DataMember
	LanguageRule *string `json:"languageRule,omitempty"`
	// @DataMember
	TemplateLanguages []string `json:"templateLanguages,omitempty"`
	// @DataMember
	Fallback *string `json:"fallback,omitempty"`
	// @DataMember
	MjmlRenderer *string `json:"mjmlRenderer,omitempty"`
}

type PasskeyCeremonyOptionsResponse struct {
	ResponseBase
	CeremonyId  string `json:"ceremonyId"`
	OptionsJson string `json:"optionsJson"`
}

type PasskeyAuthTokensResponse struct {
	ResponseBase
	AccessToken      string   `json:"accessToken"`
	RefreshToken     string   `json:"refreshToken"`
	ExpiresInSeconds int      `json:"expiresInSeconds,omitempty"`
	RecoveryCodes    []string `json:"recoveryCodes,omitempty"`
}

type PasskeyListResponse struct {
	ResponseBase
	Passkeys []PasskeyListItemDto `json:"passkeys"`
}

type PasskeyRecoveryResponse struct {
	ResponseBase
	AccessToken      string `json:"accessToken"`
	RefreshToken     string `json:"refreshToken"`
	ExpiresInSeconds int    `json:"expiresInSeconds,omitempty"`
	RemainingCodes   int    `json:"remainingCodes,omitempty"`
}

type PasskeyVerificationTokenResponse struct {
	ResponseBase
	VerificationToken string `json:"verificationToken"`
}

type FindMergedTermTreeResponse struct {
	ResponseBase
	Tree []TermTreeDto `json:"tree,omitempty"`
}

type FindTaxonomyTreeResponse struct {
	ResponseBase
	Tree []TaxonomyTreeDto `json:"tree,omitempty"`
}

type FindTermsResponse struct {
	ResponseBase
	List *PaginatedResponse[TermDto] `json:"list,omitempty"`
}

type FindTermsChildrenResponse struct {
	ResponseBase
	List *PaginatedResponse[TermDto] `json:"list,omitempty"`
}

type FindTermTreeResponse struct {
	ResponseBase
	Tree []TermTreeDto `json:"tree,omitempty"`
}

type GetDatabaseSchemaResponse struct {
	ResponseBase
	Item *SchemaDto `json:"item,omitempty"`
}

type GetDatabaseSchemasResponse struct {
	ResponseBase
	List *PaginatedResponse[SchemaListProjection] `json:"list,omitempty"`
}

type AggregateResponse struct {
	ResponseBase
	Result []interface{} `json:"result,omitempty"`
}

type CountResponse struct {
	ResponseBase
	Count int64 `json:"count,omitempty"`
}

type DistinctResponse struct {
	ResponseBase
	Values []interface{} `json:"values,omitempty"`
}

type ExecuteAggregateResponse struct {
	ResponseBase
	Result []interface{} `json:"result,omitempty"`
}

type FindResponse struct {
	ResponseBase
	List *PaginatedResponse[interface{}] `json:"list,omitempty"`
}

type FindOneResponse struct {
	ResponseBase
	Result *interface{} `json:"result,omitempty"`
}

type GetFileByIdResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl *string             `json:"publicUrl,omitempty"`
}

type GetFileInfoResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl *string             `json:"publicUrl,omitempty"`
}

type GetSignedUrlResponse struct {
	ResponseBase
	Url *string `json:"url,omitempty"`
}

type ListFilesResponse struct {
	ResponseBase
	List          *PaginatedResponse[FileResourceRefDto] `json:"list,omitempty"`
	Folders       *IList[string]                         `json:"folders,omitempty"`
	PublicFolders *IList[PublicFolderDto]                `json:"publicFolders,omitempty"`
}

type RegisterFilesResponse struct {
	ResponseBase
	DryRun     bool                 `json:"dryRun,omitempty"`
	Registered int                  `json:"registered,omitempty"`
	Unchanged  int                  `json:"unchanged,omitempty"`
	Failed     int                  `json:"failed,omitempty"`
	Items      []RegisterFileResult `json:"items,omitempty"`
}

type RequestUploadUrlResponse struct {
	ResponseBase
	Url *string `json:"url,omitempty"`
}

type GetOwnPushDevicesResponse struct {
	ResponseBase
	Items []OwnPushDeviceDto `json:"items"`
}

type GetOwnPushMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse[OwnPushMessageDto] `json:"list,omitempty"`
}

type WriteLogsResponse struct {
	ResponseBase
	Accepted int `json:"accepted,omitempty"`
}

type GetPaymentOrderResponse struct {
	ResponseBase
	Item *PaymentOrderDto `json:"item,omitempty"`
}

type GetPaymentOrdersResponse struct {
	ResponseBase
	List *PaginatedResponse[PaymentOrderDto] `json:"list,omitempty"`
}

type StartPaymentCheckoutResponse struct {
	ResponseBase
	Checkout *PaymentOrderCheckoutDto `json:"checkout,omitempty"`
	Reused   bool                     `json:"reused,omitempty"`
	Order    *PaymentOrderDto         `json:"order,omitempty"`
}

type InvokeMarketplaceFunctionResponse struct {
	ResponseBase
	IsSuccess       bool         `json:"isSuccess,omitempty"`
	Output          *interface{} `json:"output,omitempty"`
	VendorRequestId *string      `json:"vendorRequestId,omitempty"`
}

// @DataContract
type InvokeCodeWorkflowResponse struct {
	ResponseBase
	// @DataMember
	RunId *string `json:"runId,omitempty"`
	// @DataMember
	Status *string `json:"status,omitempty"`
	// @DataMember
	Output *interface{} `json:"output,omitempty"`
	// @DataMember
	DurationMs *int64 `json:"durationMs,omitempty"`
	// @DataMember
	Error *WorkflowRunErrorDto `json:"error,omitempty"`
	// @DataMember
	Run *WorkflowRunDto `json:"run,omitempty"`
}

// @DataContract
type WorkflowRunsResponse struct {
	ResponseBase
	// @DataMember
	List *PaginatedResponse[WorkflowRunListProjection] `json:"list,omitempty"`
}

// @DataContract
type WorkflowRunResponse struct {
	ResponseBase
	// @DataMember
	Run *WorkflowRunDto `json:"run,omitempty"`
}

// @DataContract
type CodeWorkflowResumeLinkResponse struct {
	ResponseBase
	// @DataMember
	RunId *string `json:"runId,omitempty"`
	// @DataMember
	WorkflowId *string `json:"workflowId,omitempty"`
	// @DataMember
	WorkflowName *string `json:"workflowName,omitempty"`
	// @DataMember
	StepId *string `json:"stepId,omitempty"`
	// @DataMember
	StepName *string `json:"stepName,omitempty"`
	// @DataMember
	Mode *string `json:"mode,omitempty"`
	// @DataMember
	Decision *string `json:"decision,omitempty"`
	// @DataMember
	Key *string `json:"key,omitempty"`
	// @DataMember
	AssigneeUserId *string `json:"assigneeUserId,omitempty"`
	// @DataMember
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
}

// @DataContract
type WorkflowRunLogsResponse struct {
	ResponseBase
	// @DataMember
	List *PaginatedResponse[TenantLogEntryDto] `json:"list,omitempty"`
}

type GetNotificationInboxResponse struct {
	ResponseBase
	List *PaginatedResponse[InAppNotificationDto] `json:"list,omitempty"`
}

type GetUnreadCountResponse struct {
	ResponseBase
	Count int64 `json:"count,omitempty"`
}

type MarkNotificationReadResponse struct {
	ResponseBase
}

type MarkAllNotificationsReadResponse struct {
	ResponseBase
}

type GetContactResponse struct {
	ResponseBase
	Item *UserDto `json:"item,omitempty"`
}

type GetAllContactsResponse struct {
	ResponseBase
	Items      *IReadOnlyList[UserDto] `json:"items,omitempty"`
	NextCursor *string                 `json:"nextCursor,omitempty"`
}

type RepairContactRolesResponse struct {
	ResponseBase
	Result *ContactRolesRepairDto `json:"result,omitempty"`
}

type GetFileResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl *string             `json:"publicUrl,omitempty"`
}

type GetFolderFilesResponse struct {
	ResponseBase
	List          *PaginatedResponse[FileResourceRefDto] `json:"list,omitempty"`
	Folders       *IList[string]                         `json:"folders,omitempty"`
	PublicFolders *IList[PublicFolderDto]                `json:"publicFolders,omitempty"`
}

type GetCollectionImportResponse struct {
	ResponseBase
	Result *CollectionImportDto `json:"result,omitempty"`
}

type GetCollectionImportsResponse struct {
	ResponseBase
	Result *PaginatedResponse[CollectionImportDto] `json:"result,omitempty"`
}

type RequestImportUploadUrlResponse struct {
	ResponseBase
	Result *ImportUploadTargetDto `json:"result,omitempty"`
}

type AnalyzeImportFileResponse struct {
	ResponseBase
	Result *ImportFileAnalysisDto `json:"result,omitempty"`
}

type GetEmailCampaignResponse struct {
	ResponseBase
	Item *EmailCampaignDto `json:"item,omitempty"`
}

type GetEmailCampaignsResponse struct {
	ResponseBase
	List *PaginatedResponse[EmailCampaignListProjection] `json:"list,omitempty"`
}

type GetEmailCampaignBatchesResponse struct {
	ResponseBase
	List *PaginatedResponse[EmailCampaignBatchDto] `json:"list,omitempty"`
}

type GetEmailCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *EmailCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

type GetEmailCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []BatchStatusChangeEntryDto                           `json:"batchStatusHistory,omitempty"`
	List               *PaginatedResponse[EmailCampaignBatchNotificationDto] `json:"list,omitempty"`
}

type GetEmailCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

type PreviewEmailNotificationResponse struct {
	ResponseBase
	Subject *string `json:"subject,omitempty"`
	Body    *string `json:"body,omitempty"`
}

// @DataContract
type RetryFailedEmailNotificationsResponse struct {
	ResponseBase
	// @DataMember
	RetryRequestId *string `json:"retryRequestId,omitempty"`
	// @DataMember
	Selected int `json:"selected,omitempty"`
	// @DataMember
	Skipped map[string]int `json:"skipped,omitempty"`
	// @DataMember
	Replayed bool `json:"replayed,omitempty"`
	// @DataMember
	Warning *string `json:"warning,omitempty"`
}

type GetEmailCampaignMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse[EmailCampaignBatchNotificationDto] `json:"list,omitempty"`
}

type GetSmsCampaignResponse struct {
	ResponseBase
	SmsCampaign *SmsCampaignDto `json:"smsCampaign,omitempty"`
}

type GetSmsCampaignsResponse struct {
	ResponseBase
	List *PaginatedResponse[SmsCampaignDto] `json:"list,omitempty"`
}

type GetSmsCampaignBatchesResponse struct {
	ResponseBase
	List *PaginatedResponse[SmsCampaignBatchDto] `json:"list,omitempty"`
}

type GetSmsCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *SmsCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

type GetSmsCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []BatchStatusChangeEntryDto                         `json:"batchStatusHistory,omitempty"`
	List               *PaginatedResponse[SmsCampaignBatchNotificationDto] `json:"list,omitempty"`
}

type GetSmsCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

type PreviewSmsNotificationResponse struct {
	ResponseBase
	Body *string `json:"body,omitempty"`
}

// @DataContract
type RetryFailedSmsNotificationsResponse struct {
	ResponseBase
	// @DataMember
	RetryRequestId *string `json:"retryRequestId,omitempty"`
	// @DataMember
	Selected int `json:"selected,omitempty"`
	// @DataMember
	Skipped map[string]int `json:"skipped,omitempty"`
	// @DataMember
	Replayed bool `json:"replayed,omitempty"`
	// @DataMember
	Warning *string `json:"warning,omitempty"`
}

type GetSmsCampaignMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse[SmsCampaignBatchNotificationDto] `json:"list,omitempty"`
}

type GetPushCampaignAudienceCountResponse struct {
	ResponseBase
	Result *PushAudienceCountDto `json:"result,omitempty"`
}

type GetPushCampaignResponse struct {
	ResponseBase
	Item *PushCampaignDto `json:"item,omitempty"`
}

type GetPushCampaignsResponse struct {
	ResponseBase
	List *PaginatedResponse[PushCampaignDto] `json:"list,omitempty"`
}

type GetPushCampaignBatchesResponse struct {
	ResponseBase
	List *PaginatedResponse[PushCampaignBatchDto] `json:"list,omitempty"`
}

type GetPushCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *PushCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

type GetPushCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []BatchStatusChangeEntryDto                          `json:"batchStatusHistory,omitempty"`
	List               *PaginatedResponse[PushCampaignBatchNotificationDto] `json:"list,omitempty"`
}

type GetPushCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

type PreviewPushNotificationResponse struct {
	ResponseBase
	Title    *string `json:"title,omitempty"`
	Body     *string `json:"body,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
}

// @DataContract
type RetryFailedPushNotificationsResponse struct {
	ResponseBase
	// @DataMember
	RetryRequestId *string `json:"retryRequestId,omitempty"`
	// @DataMember
	Selected int `json:"selected,omitempty"`
	// @DataMember
	Skipped map[string]int `json:"skipped,omitempty"`
	// @DataMember
	Replayed bool `json:"replayed,omitempty"`
	// @DataMember
	Warning *string `json:"warning,omitempty"`
}

type GetPushCampaignMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse[PushCampaignBatchNotificationDto] `json:"list,omitempty"`
}

type SharedTypeGen struct {
	TypegenStringFieldDto              *StringFieldDto              `json:"typegen_StringFieldDto,omitempty"`
	TypegenDecimalFieldDto             *DecimalFieldDto             `json:"typegen_DecimalFieldDto,omitempty"`
	TypegenCurrencyFieldDto            *CurrencyFieldDto            `json:"typegen_CurrencyFieldDto,omitempty"`
	TypegenBooleanFieldDto             *BooleanFieldDto             `json:"typegen_BooleanFieldDto,omitempty"`
	TypegenDateFieldDto                *DateFieldDto                `json:"typegen_DateFieldDto,omitempty"`
	TypegenIntegerFieldDto             *IntegerFieldDto             `json:"typegen_IntegerFieldDto,omitempty"`
	TypegenGeolocationFieldDto         *GeolocationFieldDto         `json:"typegen_GeolocationFieldDto,omitempty"`
	TypegenTagsFieldDto                *TagsFieldDto                `json:"typegen_TagsFieldDto,omitempty"`
	TypegenFileFieldDto                *FileFieldDto                `json:"typegen_FileFieldDto,omitempty"`
	TypegenTaxonomySelectionFieldDto   *TaxonomySelectionFieldDto   `json:"typegen_TaxonomySelectionFieldDto,omitempty"`
	TypegenCollectionSelectionFieldDto *CollectionSelectionFieldDto `json:"typegen_CollectionSelectionFieldDto,omitempty"`
	TypegenUserSelectionFieldDto       *UserSelectionFieldDto       `json:"typegen_UserSelectionFieldDto,omitempty"`
	TypegenRoleSelectionFieldDto       *RoleSelectionFieldDto       `json:"typegen_RoleSelectionFieldDto,omitempty"`
	TypegenEnumSelectionFieldDto       *EnumSelectionFieldDto       `json:"typegen_EnumSelectionFieldDto,omitempty"`
	TypegenObjectFieldDto              *ObjectFieldDto              `json:"typegen_ObjectFieldDto,omitempty"`
	TypegenArrayFieldDto               *ArrayFieldDto               `json:"typegen_ArrayFieldDto,omitempty"`
	TypegenJsonFieldDto                *JsonFieldDto                `json:"typegen_JsonFieldDto,omitempty"`
}

func (SharedTypeGen) CreateResponseVoid() {}
func (SharedTypeGen) HttpMethod() string  { return "POST" }

// @Route("/{version}/echo", "GET")
type Echo struct {
	RequestBase
}

func (Echo) CreateResponse() (r EchoResponse) { return }
func (Echo) HttpMethod() string               { return "GET" }

// @Route("/{version}/public/projects/{ProjectId}/brand/{Kind}", "GET")
type GetPublicProjectBrandAsset struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
	Kind      *string `json:"kind,omitempty"`
	V         *string `json:"v,omitempty"`
}

func (GetPublicProjectBrandAsset) CreateResponse() (r []byte) { return }
func (GetPublicProjectBrandAsset) HttpMethod() string         { return "GET" }

// @Route("/{version}/public/projects/{ProjectId}/config", "GET")
type GetPublicProjectConfig struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
	Env       *string `json:"env,omitempty"`
}

func (GetPublicProjectConfig) CreateResponse() (r PublicProjectConfigDto) { return }
func (GetPublicProjectConfig) HttpMethod() string                         { return "GET" }

// @Route("/{version}/public/projects/{ProjectId}/legal/{Kind}", "GET")
type GetPublicProjectLegal struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
	Kind      *string `json:"kind,omitempty"`
}

func (GetPublicProjectLegal) CreateResponse() (r PublicLegalDocumentDto) { return }
func (GetPublicProjectLegal) HttpMethod() string                         { return "GET" }

/** @description Adds a file to one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}/attachments", "POST")
// @Api(Description="Adds a file to one of the caller's own AI chats.")
type UploadEndUserChatAttachmentRequest struct {
	CodeMashRequestBase
	SessionId     string `json:"sessionId"`
	FileName      string `json:"fileName"`
	ContentType   string `json:"contentType"`
	Base64Content string `json:"base64Content"`
}

func (UploadEndUserChatAttachmentRequest) CreateResponse() (r IdResponse) { return }
func (UploadEndUserChatAttachmentRequest) HttpMethod() string             { return "POST" }

/** @description Lists the files in one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}/attachments", "GET")
// @Api(Description="Lists the files in one of the caller's own AI chats.")
type ListEndUserChatAttachmentsRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId"`
}

func (ListEndUserChatAttachmentsRequest) CreateResponse() (r ListEndUserChatAttachmentsResponse) {
	return
}
func (ListEndUserChatAttachmentsRequest) HttpMethod() string { return "GET" }

/** @description Removes a file from one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/attachments/{AttachmentId}", "DELETE")
// @Api(Description="Removes a file from one of the caller's own AI chats.")
type DeleteEndUserChatAttachmentRequest struct {
	CodeMashRequestBase
	AttachmentId string `json:"attachmentId"`
}

func (DeleteEndUserChatAttachmentRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteEndUserChatAttachmentRequest) HttpMethod() string                { return "DELETE" }

/** @description Likes, dislikes or clears one message of the caller's own AI chat. */
// @Route("/{version}/ai/chat/sessions/{SessionId}/entries/{EntryId}/feedback", "PUT")
// @Api(Description="Likes, dislikes or clears one message of the caller's own AI chat.")
type SetEndUserChatEntryFeedbackRequest struct {
	CodeMashRequestBase
	SessionId string  `json:"sessionId"`
	EntryId   string  `json:"entryId"`
	Feedback  *string `json:"feedback,omitempty"`
}

func (SetEndUserChatEntryFeedbackRequest) CreateResponse() (r EmptyResponse) { return }
func (SetEndUserChatEntryFeedbackRequest) HttpMethod() string                { return "PUT" }

/** @description Lists what the AI chat remembers about the caller. */
// @Route("/{version}/ai/chat/memory", "GET")
// @Api(Description="Lists what the AI chat remembers about the caller.")
type ListEndUserChatMemoryRequest struct {
	CodeMashRequestBase
	Take *int `json:"take,omitempty"`
}

func (ListEndUserChatMemoryRequest) CreateResponse() (r ListEndUserChatMemoryResponse) { return }
func (ListEndUserChatMemoryRequest) HttpMethod() string                                { return "GET" }

/** @description Forgets one thing the AI chat remembers about the caller. */
// @Route("/{version}/ai/chat/memory/{NoteId}", "DELETE")
// @Api(Description="Forgets one thing the AI chat remembers about the caller.")
type ForgetEndUserChatMemoryRequest struct {
	CodeMashRequestBase
	NoteId string `json:"noteId"`
}

func (ForgetEndUserChatMemoryRequest) CreateResponse() (r EmptyResponse) { return }
func (ForgetEndUserChatMemoryRequest) HttpMethod() string                { return "DELETE" }

/** @description Whether the AI chat can run for the caller, and which assistants it offers. */
// @Route("/{version}/ai/chat/availability", "GET")
// @Api(Description="Whether the AI chat can run for the caller, and which assistants it offers.")
type GetEndUserChatAvailabilityRequest struct {
	CodeMashRequestBase
}

func (GetEndUserChatAvailabilityRequest) CreateResponse() (r GetEndUserChatAvailabilityResponse) {
	return
}
func (GetEndUserChatAvailabilityRequest) HttpMethod() string { return "GET" }

/** @description Lists the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions", "GET")
// @Api(Description="Lists the caller's own AI chats.")
type ListEndUserChatSessionsRequest struct {
	CodeMashRequestBase
	Take            *int `json:"take,omitempty"`
	IncludeArchived bool `json:"includeArchived,omitempty"`
}

func (ListEndUserChatSessionsRequest) CreateResponse() (r ListEndUserChatSessionsResponse) { return }
func (ListEndUserChatSessionsRequest) HttpMethod() string                                  { return "GET" }

/** @description Opens a new AI chat for the caller. */
// @Route("/{version}/ai/chat/sessions", "POST")
// @Api(Description="Opens a new AI chat for the caller.")
type CreateEndUserChatSessionRequest struct {
	CodeMashRequestBase
	AssistantId *string `json:"assistantId,omitempty"`
	Title       *string `json:"title,omitempty"`
}

func (CreateEndUserChatSessionRequest) CreateResponse() (r IdResponse) { return }
func (CreateEndUserChatSessionRequest) HttpMethod() string             { return "POST" }

/** @description Returns one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}", "GET")
// @Api(Description="Returns one of the caller's own AI chats.")
type GetEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId"`
}

func (GetEndUserChatSessionRequest) CreateResponse() (r GetEndUserChatSessionResponse) { return }
func (GetEndUserChatSessionRequest) HttpMethod() string                                { return "GET" }

/** @description Renames one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}", "PATCH")
// @Api(Description="Renames one of the caller's own AI chats.")
type RenameEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string  `json:"sessionId"`
	Title     *string `json:"title,omitempty"`
}

func (RenameEndUserChatSessionRequest) CreateResponse() (r EmptyResponse) { return }
func (RenameEndUserChatSessionRequest) HttpMethod() string                { return "PATCH" }

/** @description Pins or unpins one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}/pin", "PUT")
// @Api(Description="Pins or unpins one of the caller's own AI chats.")
type PinEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId"`
	Pinned    bool   `json:"pinned,omitempty"`
}

func (PinEndUserChatSessionRequest) CreateResponse() (r EmptyResponse) { return }
func (PinEndUserChatSessionRequest) HttpMethod() string                { return "PUT" }

/** @description Archives or unarchives one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}/archive", "PUT")
// @Api(Description="Archives or unarchives one of the caller's own AI chats.")
type ArchiveEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId"`
	Archived  bool   `json:"archived,omitempty"`
}

func (ArchiveEndUserChatSessionRequest) CreateResponse() (r EmptyResponse) { return }
func (ArchiveEndUserChatSessionRequest) HttpMethod() string                { return "PUT" }

/** @description Deletes one of the caller's own AI chats. */
// @Route("/{version}/ai/chat/sessions/{SessionId}", "DELETE")
// @Api(Description="Deletes one of the caller's own AI chats.")
type DeleteEndUserChatSessionRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId"`
}

func (DeleteEndUserChatSessionRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteEndUserChatSessionRequest) HttpMethod() string                { return "DELETE" }

/** @description Returns a page of one of the caller's own AI chat transcripts. */
// @Route("/{version}/ai/chat/sessions/{SessionId}/entries", "GET")
// @Api(Description="Returns a page of one of the caller's own AI chat transcripts.")
type GetEndUserChatEntriesRequest struct {
	CodeMashRequestBase
	SessionId string `json:"sessionId"`
	AfterSeq  *int64 `json:"afterSeq,omitempty"`
	Take      *int   `json:"take,omitempty"`
}

func (GetEndUserChatEntriesRequest) CreateResponse() (r GetEndUserChatEntriesResponse) { return }
func (GetEndUserChatEntriesRequest) HttpMethod() string                                { return "GET" }

/** @description Sends a message to the AI chat; the answer streams on the caller's channel. */
// @Route("/{version}/ai/chat/turn", "POST")
// @Api(Description="Sends a message to the AI chat; the answer streams on the caller's channel.")
type StartEndUserChatTurnRequest struct {
	CodeMashRequestBase
	SessionId   *string `json:"sessionId,omitempty"`
	AssistantId *string `json:"assistantId,omitempty"`
	Message     string  `json:"message"`
}

func (StartEndUserChatTurnRequest) CreateResponse() (r StartEndUserChatTurnResponse) { return }
func (StartEndUserChatTurnRequest) HttpMethod() string                               { return "POST" }

/** @description Lists the AI tools a project user may use: only their own data (own:* toolsets). */
// @Route("/{version}/ai/tools", "GET")
// @Api(Description="Lists the AI tools a project user may use: only their own data (own:* toolsets).")
type GetEndUserAiToolsRequest struct {
	RequestBase
}

func (GetEndUserAiToolsRequest) CreateResponse() (r GetEndUserAiToolsResponse) { return }
func (GetEndUserAiToolsRequest) HttpMethod() string                            { return "GET" }

/** @description Invokes one own-scope AI tool as the calling project user. */
// @Route("/{version}/ai/tools/{ToolName}", "POST")
// @Api(Description="Invokes one own-scope AI tool as the calling project user.")
type InvokeEndUserAiToolRequest struct {
	RequestBase
	ToolName      string  `json:"toolName"`
	ArgumentsJson *string `json:"argumentsJson,omitempty"`
}

func (InvokeEndUserAiToolRequest) CreateResponse() (r InvokeEndUserAiToolResponse) { return }
func (InvokeEndUserAiToolRequest) HttpMethod() string                              { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/block", "PATCH")
// @Api(Description="Membership")
// @DataContract
type BlockUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to block, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to block, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (BlockUserRequest) CreateResponse() (r EmptyResponse) { return }
func (BlockUserRequest) HttpMethod() string                { return "PATCH" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/service", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveSystemUserWithPermissions struct {
	SaveUserWithRolesBase
}

func (SaveSystemUserWithPermissions) CreateResponse() (r IdResponse) { return }
func (SaveSystemUserWithPermissions) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/guest", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveGuestUser struct {
	SaveUser
}

func (SaveGuestUser) CreateResponse() (r SaveGuestUserResponse) { return }
func (SaveGuestUser) HttpMethod() string                        { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/phone", "POST")
// @Api(Description="Membership")
// @DataContract
type SavePhoneUser struct {
	SaveUser
	/** @description Phone number for the new user, in E.164 format. */
	// @DataMember
	// @ApiMember(Description="Phone number for the new user, in E.164 format.", IsRequired=true)
	Phone string `json:"phone"`
}

func (SavePhoneUser) CreateResponse() (r IdResponse) { return }
func (SavePhoneUser) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/user-name", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveUserNameUser struct {
	SaveUser
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	UserName string `json:"userName"`
	/** @description Optional roles from the project's "Permitted user registration roles". Empty = the default registration role. */
	// @DataMember
	// @ApiMember(Description="Optional roles from the project's \"Permitted user registration roles\". Empty = the default registration role.")
	Roles []string `json:"roles,omitempty"`
}

func (SaveUserNameUser) CreateResponse() (r IdResponse) { return }
func (SaveUserNameUser) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/email", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveEmailUser struct {
	SaveUser
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	Email string `json:"email"`
	/** @description Optional roles from the project's "Permitted user registration roles". Empty = the default registration role. */
	// @DataMember
	// @ApiMember(Description="Optional roles from the project's \"Permitted user registration roles\". Empty = the default registration role.")
	Roles []string `json:"roles,omitempty"`
}

func (SaveEmailUser) CreateResponse() (r IdResponse) { return }
func (SaveEmailUser) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/phone-with-permissions", "POST")
// @Api(Description="Membership")
// @DataContract
type SavePhoneUserNameWithPermissions struct {
	SaveUserWithRolesBase
	// @DataMember
	Phone string `json:"phone"`
}

func (SavePhoneUserNameWithPermissions) CreateResponse() (r IdResponse) { return }
func (SavePhoneUserNameWithPermissions) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/email-with-permissions", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveEmailUserNameWithPermissions struct {
	SaveUserWithRolesBase
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	Email string `json:"email"`
}

func (SaveEmailUserNameWithPermissions) CreateResponse() (r IdResponse) { return }
func (SaveEmailUserNameWithPermissions) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/register/user-name-with-permissions", "POST")
// @Api(Description="Membership")
// @DataContract
type SaveUserNameWithPermissions struct {
	SaveUserWithRolesBase
	// @DataMember
	Password string `json:"password"`
	// @DataMember
	UserName string `json:"userName"`
}

func (SaveUserNameWithPermissions) CreateResponse() (r IdResponse) { return }
func (SaveUserNameWithPermissions) HttpMethod() string             { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth", "DELETE")
// @Api(Description="Membership")
// @DataContract
type DeleteUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to delete, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to delete, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (DeleteUserRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteUserRequest) HttpMethod() string                { return "DELETE" }

/** @description Membership */
// @Route("/{version}/membership/auth/{id}", "GET")
// @Api(Description="Membership")
// @DataContract
type GetUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to fetch, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to fetch, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetUserRequest) CreateResponse() (r GetUserResponse) { return }
func (GetUserRequest) HttpMethod() string                  { return "GET" }

/** @description Membership */
// @Route("/{version}/membership/auth", "GET")
// @Api(Description="Membership")
// @DataContract
type GetUsersRequest struct {
	CodeMashListPaginationRequestBase
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Include each user's effective permissions in the result. */
	// @DataMember
	// @ApiMember(Description="Include each user's effective permissions in the result.")
	IncludePermissions bool `json:"includePermissions,omitempty"`
	/** @description Only return users that have a registered push device. */
	// @DataMember
	// @ApiMember(Description="Only return users that have a registered push device.")
	UserShouldHavePushDevice bool `json:"userShouldHavePushDevice,omitempty"`
	/** @description Only return users that have an email address. */
	// @DataMember
	// @ApiMember(Description="Only return users that have an email address.")
	UserShouldHaveEmail bool `json:"userShouldHaveEmail,omitempty"`
	/** @description Include each user's metadata in the result. */
	// @DataMember
	// @ApiMember(Description="Include each user's metadata in the result.")
	IncludeMeta bool `json:"includeMeta,omitempty"`
	/** @description Filter to users that have any of these role names. */
	// @DataMember
	// @ApiMember(Description="Filter to users that have any of these role names.")
	RoleNames []string `json:"roleNames,omitempty"`
	/** @description Filter to these specific user ids. */
	// @DataMember
	// @ApiMember(Description="Filter to these specific user ids.")
	UserIds []string `json:"userIds,omitempty"`
	/** @description Free text, case-insensitive: matches email, user name, display / first / last name or phone. */
	// @DataMember
	// @ApiMember(Description="Free text, case-insensitive: matches email, user name, display / first / last name or phone.")
	Search *string `json:"search,omitempty"`
	/** @description Case-insensitive part of the email. */
	// @DataMember
	// @ApiMember(Description="Case-insensitive part of the email.")
	Email *string `json:"email,omitempty"`
	/** @description Case-insensitive part of the user name. */
	// @DataMember
	// @ApiMember(Description="Case-insensitive part of the user name.")
	UserName *string `json:"userName,omitempty"`
	/** @description Status names, any of: Registered, PendingValidation, Active, Unregistered, Suspended, InActive, Blocked. */
	// @DataMember
	// @ApiMember(Description="Status names, any of: Registered, PendingValidation, Active, Unregistered, Suspended, InActive, Blocked.")
	Statuses []string `json:"statuses,omitempty"`
	/** @description Login type names, any of: Service, Email, UserName, Phone, Guest, Social. */
	// @DataMember
	// @ApiMember(Description="Login type names, any of: Service, Email, UserName, Phone, Guest, Social.")
	Types []string `json:"types,omitempty"`
}

func (GetUsersRequest) CreateResponse() (r GetUsersResponse) { return }
func (GetUsersRequest) HttpMethod() string                   { return "GET" }

/** @description Membership */
// @Route("/{version}/membership/auth/{id}/preferences", "GET")
// @Api(Description="Membership")
// @DataContract
type GetUserPreferencesRequest struct {
	CodeMashRequestBase
	/** @description Id of the user whose preferences to fetch, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user whose preferences to fetch, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the project's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the project's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetUserPreferencesRequest) CreateResponse() (r GetUserPreferencesResponse) { return }
func (GetUserPreferencesRequest) HttpMethod() string                             { return "GET" }

// @Route("/{version}/membership/users/{contactId}/marketing-state/{channel}/consent", "POST")
type GrantContactConsentRequest struct {
	CodeMashRequestBase
	/** @description Id of the user (contact) to grant consent for. */
	// @ApiMember(Description="Id of the user (contact) to grant consent for.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Delivery channel to grant consent on: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel to grant consent on: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
	/** @description Lawful basis for the consent, e.g. Consent. Defaults to Consent. */
	// @ApiMember(Description="Lawful basis for the consent, e.g. Consent. Defaults to Consent.")
	LawfulBasis string `json:"lawfulBasis"`
	/** @description Source of the consent, e.g. UserOptIn. Defaults to UserOptIn. */
	// @ApiMember(Description="Source of the consent, e.g. UserOptIn. Defaults to UserOptIn.")
	Source string `json:"source"`
	/** @description Optional free-text reference to evidence of consent (e.g. a form submission id). */
	// @ApiMember(Description="Optional free-text reference to evidence of consent (e.g. a form submission id).")
	EvidenceRef *string `json:"evidenceRef,omitempty"`
}

func (GrantContactConsentRequest) CreateResponse() (r EmptyResponse) { return }
func (GrantContactConsentRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/invite", "POST")
// @Api(Description="Membership")
// @DataContract
type InviteUserRequest struct {
	CodeMashRequestBase
	/** @description Email address the invitation is sent to. It cannot be changed later. */
	// @DataMember
	// @ApiMember(Description="Email address the invitation is sent to. It cannot be changed later.", IsRequired=true)
	Email string `json:"email"`
	/** @description The invited person's full name, e.g. "Rūta Petraitė". Stored as the full name, first name (first word) and last name (the rest). The user may change it when finishing. */
	// @DataMember
	// @ApiMember(Description="The invited person's full name, e.g. \"Rūta Petraitė\". Stored as the full name, first name (first word) and last name (the rest). The user may change it when finishing.")
	FullName *string `json:"fullName,omitempty"`
	/** @description Optional. How e-mails greet the person ("Hi {salutation},"), e.g. "Rūta". Stored as the display name. When empty the full name is used. The user may change it when finishing. */
	// @DataMember
	// @ApiMember(Description="Optional. How e-mails greet the person (\"Hi {salutation},\"), e.g. \"Rūta\". Stored as the display name. When empty the full name is used. The user may change it when finishing.")
	Salutation *string `json:"salutation,omitempty"`
	/** @description Optional. Roles the invited user gets, by name ("Editor") or role id. Empty = the project's default registration role. The same rules as creating a user: an administrator may give any role, anyone else only the allowed registration roles. */
	// @DataMember
	// @ApiMember(Description="Optional. Roles the invited user gets, by name (\"Editor\") or role id. Empty = the project's default registration role. The same rules as creating a user: an administrator may give any role, anyone else only the allowed registration roles.")
	Roles []string `json:"roles,omitempty"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (InviteUserRequest) CreateResponse() (r EmptyResponse) { return }
func (InviteUserRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/invitation/finish", "POST")
// @Api(Description="Membership")
// @DataContract
type FinishInvitationRequest struct {
	CodeMashRequestBase
	/** @description One-time invitation token from the e-mail link (the #token= fragment). */
	// @DataMember
	// @ApiMember(Description="One-time invitation token from the e-mail link (the #token= fragment).", IsRequired=true)
	Token string `json:"token"`
	/** @description The user's password. Validated against the project's complexity policy. */
	// @DataMember
	// @ApiMember(Description="The user's password. Validated against the project's complexity policy.", IsRequired=true)
	Password string `json:"password"`
	/** @description Optional. The user's full name; replaces the one the admin entered (full name, first name = first word, last name = the rest). Empty = keep the admin's. */
	// @DataMember
	// @ApiMember(Description="Optional. The user's full name; replaces the one the admin entered (full name, first name = first word, last name = the rest). Empty = keep the admin's.")
	FullName *string `json:"fullName,omitempty"`
	/** @description Optional. How e-mails greet the user (stored as the display name). Empty = keep the admin's, or the full name when the admin gave none. */
	// @DataMember
	// @ApiMember(Description="Optional. How e-mails greet the user (stored as the display name). Empty = keep the admin's, or the full name when the admin gave none.")
	Salutation *string `json:"salutation,omitempty"`
	/** @description Optional. Older form of Salutation (the display name); Salutation wins when both are sent. */
	// @DataMember
	// @ApiMember(Description="Optional. Older form of Salutation (the display name); Salutation wins when both are sent.")
	DisplayName *string `json:"displayName,omitempty"`
	/** @description Optional first name; wins over the one taken from FullName. */
	// @DataMember
	// @ApiMember(Description="Optional first name; wins over the one taken from FullName.")
	FirstName *string `json:"firstName,omitempty"`
	/** @description Optional last name; wins over the one taken from FullName. */
	// @DataMember
	// @ApiMember(Description="Optional last name; wins over the one taken from FullName.")
	LastName *string `json:"lastName,omitempty"`
}

func (FinishInvitationRequest) CreateResponse() (r EmptyResponse) { return }
func (FinishInvitationRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/{UserId}/link-identity/start", "POST")
// @Api(Description="Membership")
// @DataContract
type StartLinkIdentityRequest struct {
	CodeMashRequestBase
	// @DataMember
	UserId string `json:"userId"`
	/** @description google, apple, microsoft, github, facebook, x or okta. */
	// @DataMember
	// @ApiMember(Description="google, apple, microsoft, github, facebook, x or okta.", IsRequired=true)
	Provider string `json:"provider"`
	// @DataMember
	Mode *string `json:"mode,omitempty"`
	// @DataMember
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// @DataMember
	CodeChallenge *string `json:"codeChallenge,omitempty"`
}

func (StartLinkIdentityRequest) CreateResponse() (r StartLinkIdentityResponse) { return }
func (StartLinkIdentityRequest) HttpMethod() string                            { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/{UserId}/link-identity", "POST")
// @Api(Description="Membership")
// @DataContract
type LinkIdentityRequest struct {
	CodeMashRequestBase
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	Code string `json:"code"`
	// @DataMember
	CodeVerifier *string `json:"codeVerifier,omitempty"`
}

func (LinkIdentityRequest) CreateResponse() (r EmptyResponse) { return }
func (LinkIdentityRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/users/{userId}/map-auth", "POST")
// @Api(Description="Membership")
type MapAuthToUserRequest struct {
	CodeMashRequestBase
	UserId                string  `json:"userId"`
	AuthId                string  `json:"authId"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (MapAuthToUserRequest) CreateResponse() (r EmptyResponse) { return }
func (MapAuthToUserRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/assign-roles", "PUT")
// @Api(Description="Membership")
// @DataContract
type AssignRolePermissionsRequest struct {
	CodeMashRequestBase
	/** @description Id of the user login to assign roles to, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user login to assign roles to, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The complete new list of role names (full replacement), from get_roles. */
	// @DataMember
	// @ApiMember(Description="The complete new list of role names (full replacement), from get_roles.")
	Roles []string `json:"roles,omitempty"`
}

func (AssignRolePermissionsRequest) CreateResponse() (r EmptyResponse) { return }
func (AssignRolePermissionsRequest) HttpMethod() string                { return "PUT" }

/** @description Membership */
// @Route("/{version}/membership/users/{userId}/roles", "PUT")
// @Api(Description="Membership")
// @DataContract
type SetContactRolesRequest struct {
	CodeMashRequestBase
	/** @description Id of the human user to assign roles to. */
	// @DataMember
	// @ApiMember(Description="Id of the human user to assign roles to.", IsRequired=true)
	UserId string `json:"userId"`
	/** @description The complete new list of role ids (full replacement), from get_roles. Empty/omitted clears all roles. */
	// @DataMember
	// @ApiMember(Description="The complete new list of role ids (full replacement), from get_roles. Empty/omitted clears all roles.")
	Roles []string `json:"roles,omitempty"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (SetContactRolesRequest) CreateResponse() (r EmptyResponse) { return }
func (SetContactRolesRequest) HttpMethod() string                { return "PUT" }

// @Route("/{version}/membership/users/{contactId}/marketing-state/{commChannel}/{channel}/tags/{tag}", "PUT")
type SetContactTagSubscriptionRequest struct {
	CodeMashRequestBase
	/** @description Id of the user (contact) to update. */
	// @ApiMember(Description="Id of the user (contact) to update.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Communication channel type: Marketing or Transactional. */
	// @ApiMember(Description="Communication channel type: Marketing or Transactional.", IsRequired=true)
	CommChannel string `json:"commChannel"`
	/** @description Delivery channel: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
	/** @description The tag name; must already exist for the communication channel. */
	// @ApiMember(Description="The tag name; must already exist for the communication channel.", IsRequired=true)
	Tag string `json:"tag"`
	/** @description True to subscribe (unblock) the tag, false to block it. */
	// @ApiMember(Description="True to subscribe (unblock) the tag, false to block it.")
	Subscribed bool `json:"subscribed,omitempty"`
}

func (SetContactTagSubscriptionRequest) CreateResponse() (r EmptyResponse) { return }
func (SetContactTagSubscriptionRequest) HttpMethod() string                { return "PUT" }

/** @description Membership */
// @Route("/{version}/membership/auth/social/{Provider}", "GET")
// @Api(Description="Membership")
// @DataContract
type StartSocialSignInRequest struct {
	CodeMashRequestBase
	/** @description google, apple, microsoft, github, facebook, x or okta. */
	// @DataMember
	// @ApiMember(Description="google, apple, microsoft, github, facebook, x or okta.", IsRequired=true)
	Provider string `json:"provider"`
	/** @description Name of the integration's OAuth mode (which app page to come back to). Default: the only mode, else the first by name. */
	// @DataMember
	// @ApiMember(Description="Name of the integration's OAuth mode (which app page to come back to). Default: the only mode, else the first by name.")
	Mode *string `json:"mode,omitempty"`
	/** @description Optional app page to come back to (success and failure). Absolute https (http only on localhost), origin in the project's allowed origins or the project's own hosted Admin Portal. Wins over the mode's pages. Refused: the browser goes to the mode's failure page with error=return_url_not_allowed (JSON 052 when the mode has none). */
	// @DataMember
	// @ApiMember(Description="Optional app page to come back to (success and failure). Absolute https (http only on localhost), origin in the project's allowed origins or the project's own hosted Admin Portal. Wins over the mode's pages. Refused: the browser goes to the mode's failure page with error=return_url_not_allowed (JSON 052 when the mode has none).")
	ReturnUrl *string `json:"returnUrl,omitempty"`
	/** @description Optional PKCE S256 challenge of the app: base64url (no padding) SHA-256 of a random code verifier the app keeps (43 characters). When set, the exchange needs the matching codeVerifier, so only the browser / app that started can use the one-time code. Strongly recommended for web apps. */
	// @DataMember
	// @ApiMember(Description="Optional PKCE S256 challenge of the app: base64url (no padding) SHA-256 of a random code verifier the app keeps (43 characters). When set, the exchange needs the matching codeVerifier, so only the browser / app that started can use the one-time code. Strongly recommended for web apps.")
	CodeChallenge *string `json:"codeChallenge,omitempty"`
}

func (StartSocialSignInRequest) CreateResponse() (r SocialSignInStartResponse) { return }
func (StartSocialSignInRequest) HttpMethod() string                            { return "GET" }

/** @description Membership */
// @Route("/{version}/membership/auth/social/{Provider}/callback/{ProjectId}", "GET POST")
// @Api(Description="Membership")
// @DataContract
type SocialSignInCallbackRequest struct {
	CodeMashRequestBase
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	Code *string `json:"code,omitempty"`
	// @DataMember
	State *string `json:"state,omitempty"`
	// @DataMember
	Error *string `json:"error,omitempty"`
}

func (SocialSignInCallbackRequest) CreateResponse() (r SocialSignInStartResponse) { return }
func (SocialSignInCallbackRequest) HttpMethod() string                            { return "GET" }

/** @description Membership */
// @Route("/{version}/membership/auth/social/{Provider}/exchange", "POST")
// @Api(Description="Membership")
// @DataContract
type ExchangeSocialSignInCodeRequest struct {
	CodeMashRequestBase
	/** @description google, apple, microsoft, github, facebook, x or okta. */
	// @DataMember
	// @ApiMember(Description="google, apple, microsoft, github, facebook, x or okta.", IsRequired=true)
	Provider string `json:"provider"`
	/** @description The one-time code from the callback URL (valid 2 minutes, once). */
	// @DataMember
	// @ApiMember(Description="The one-time code from the callback URL (valid 2 minutes, once).", IsRequired=true)
	Code string `json:"code"`
	/** @description The PKCE code verifier whose S256 was sent as codeChallenge at the start. Required when the start had a codeChallenge; must be left out when it had none (CM-ERRORS-MEMBERSHIP-058, the code is used up). */
	// @DataMember
	// @ApiMember(Description="The PKCE code verifier whose S256 was sent as codeChallenge at the start. Required when the start had a codeChallenge; must be left out when it had none (CM-ERRORS-MEMBERSHIP-058, the code is used up).")
	CodeVerifier *string `json:"codeVerifier,omitempty"`
}

func (ExchangeSocialSignInCodeRequest) CreateResponse() (r SocialSignInTokenResponse) { return }
func (ExchangeSocialSignInCodeRequest) HttpMethod() string                            { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/unblock", "PATCH")
// @Api(Description="Membership")
// @DataContract
type UnblockUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to unblock, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to unblock, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (UnblockUserRequest) CreateResponse() (r EmptyResponse) { return }
func (UnblockUserRequest) HttpMethod() string                { return "PATCH" }

// @Route("/{version}/membership/users/{contactId}/marketing-state/{channel}/unsubscribe", "POST")
type UnsubscribeContactRequest struct {
	CodeMashRequestBase
	/** @description Id of the user (contact) to unsubscribe. */
	// @ApiMember(Description="Id of the user (contact) to unsubscribe.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Delivery channel to unsubscribe from: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel to unsubscribe from: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
	/** @description Optional suppression reason name explaining why consent was revoked. */
	// @ApiMember(Description="Optional suppression reason name explaining why consent was revoked.")
	Reason *string `json:"reason,omitempty"`
}

func (UnsubscribeContactRequest) CreateResponse() (r EmptyResponse) { return }
func (UnsubscribeContactRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth", "PUT")
// @Api(Description="Membership")
// @DataContract
type UpdateUserRequest struct {
	SaveUser
	/** @description Id of the user to update, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to update, from get_users.", IsRequired=true)
	Id string `json:"id"`
}

func (UpdateUserRequest) CreateResponse() (r IdResponse) { return }
func (UpdateUserRequest) HttpMethod() string             { return "PUT" }

/** @description Membership */
// @Route("/{version}/membership/auth/{id}/preferences", "PUT")
// @Api(Description="Membership")
// @DataContract
type UpdateUserPreferencesRequest struct {
	CodeMashRequestBase
	/** @description Id of the user to update, from get_users. */
	// @DataMember
	// @ApiMember(Description="Id of the user to update, from get_users.", IsRequired=true)
	Id string `json:"id"`
	/** @description When true, blocks all marketing messages to this user. */
	// @DataMember
	// @ApiMember(Description="When true, blocks all marketing messages to this user.")
	BlockAllMarketingMessages bool `json:"blockAllMarketingMessages,omitempty"`
	/** @description Per communication channel, the set of tags blocked for this user. Full replacement. */
	// @DataMember
	// @ApiMember(Description="Per communication channel, the set of tags blocked for this user. Full replacement.")
	BlockedTags map[string][]string `json:"blockedTags,omitempty"`
	/** @description Database integration id. Optional — defaults to the project's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the project's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (UpdateUserPreferencesRequest) CreateResponse() (r EmptyResponse) { return }
func (UpdateUserPreferencesRequest) HttpMethod() string                { return "PUT" }

/** @description Membership · Password */
// @Route("/{version}/membership/userauth/password/change", "POST")
// @Api(Description="Membership · Password")
// @DataContract
type ChangePasswordRequest struct {
	CodeMashRequestBase
	/** @description The member's current password. */
	// @DataMember
	// @ApiMember(Description="The member's current password.", IsRequired=true)
	CurrentPassword string `json:"currentPassword"`
	/** @description The new password. Validated against the project's complexity policy. */
	// @DataMember
	// @ApiMember(Description="The new password. Validated against the project's complexity policy.", IsRequired=true)
	NewPassword string `json:"newPassword"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (ChangePasswordRequest) CreateResponse() (r CredentialChangedResponse) { return }
func (ChangePasswordRequest) HttpMethod() string                            { return "POST" }

/** @description Membership · Password */
// @Route("/{version}/membership/userauth/password/reset/request", "POST")
// @Api(Description="Membership · Password")
// @DataContract
type RequestPasswordResetRequest struct {
	CodeMashRequestBase
	/** @description Email address to send the reset link to. */
	// @DataMember
	// @ApiMember(Description="Email address to send the reset link to.", IsRequired=true)
	Email string `json:"email"`
}

func (RequestPasswordResetRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (RequestPasswordResetRequest) HttpMethod() string                    { return "POST" }

/** @description Membership · Password */
// @Route("/{version}/membership/userauth/password/reset/confirm", "POST")
// @Api(Description="Membership · Password")
// @DataContract
type ConfirmPasswordResetRequest struct {
	CodeMashRequestBase
	/** @description One-time reset token from the email link. */
	// @DataMember
	// @ApiMember(Description="One-time reset token from the email link.", IsRequired=true)
	Token string `json:"token"`
	/** @description The new password. Validated against the project's complexity policy. */
	// @DataMember
	// @ApiMember(Description="The new password. Validated against the project's complexity policy.", IsRequired=true)
	NewPassword string `json:"newPassword"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (ConfirmPasswordResetRequest) CreateResponse() (r CredentialChangedResponse) { return }
func (ConfirmPasswordResetRequest) HttpMethod() string                            { return "POST" }

/** @description Membership · Password */
// @Route("/{version}/membership/userauth/password/reset/test", "POST")
// @Api(Description="Membership · Password")
// @DataContract
type SendTestPasswordResetRequest struct {
	CodeMashRequestBase
	/** @description Email address of the project user (a login in the request environment). */
	// @DataMember
	// @ApiMember(Description="Email address of the project user (a login in the request environment).", IsRequired=true)
	Email string `json:"email"`
}

func (SendTestPasswordResetRequest) CreateResponse() (r SendTestPasswordResetResponse) { return }
func (SendTestPasswordResetRequest) HttpMethod() string                                { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/authentication-options", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type PasskeyAuthenticationOptionsRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

func (PasskeyAuthenticationOptionsRequest) CreateResponse() (r PasskeyCeremonyOptionsResponse) {
	return
}
func (PasskeyAuthenticationOptionsRequest) HttpMethod() string { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/verify-authentication", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type VerifyPasskeyAuthenticationRequest struct {
	CodeMashRequestBase
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AssertionResponse string `json:"assertionResponse"`
}

func (VerifyPasskeyAuthenticationRequest) CreateResponse() (r PasskeyAuthTokensResponse) { return }
func (VerifyPasskeyAuthenticationRequest) HttpMethod() string                            { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkeys", "GET")
// @Api(Description="Membership · Passkey")
// @DataContract
type ListPasskeysRequest struct {
	CodeMashRequestBase
}

func (ListPasskeysRequest) CreateResponse() (r PasskeyListResponse) { return }
func (ListPasskeysRequest) HttpMethod() string                      { return "GET" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkeys/{CredentialId}/rename", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RenamePasskeyRequest struct {
	CodeMashRequestBase
	/** @description Base64 credential id of the passkey to rename, from list_passkeys. */
	// @DataMember
	// @ApiMember(Description="Base64 credential id of the passkey to rename, from list_passkeys.", IsRequired=true)
	CredentialId string `json:"credentialId"`
	/** @description The new friendly name for the passkey. */
	// @DataMember
	// @ApiMember(Description="The new friendly name for the passkey.", IsRequired=true)
	FriendlyName string `json:"friendlyName"`
}

func (RenamePasskeyRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (RenamePasskeyRequest) HttpMethod() string                    { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkeys/{CredentialId}/revoke", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RevokePasskeyRequest struct {
	CodeMashRequestBase
	/** @description Base64 credential id of the passkey to revoke, from list_passkeys. */
	// @DataMember
	// @ApiMember(Description="Base64 credential id of the passkey to revoke, from list_passkeys.", IsRequired=true)
	CredentialId string `json:"credentialId"`
}

func (RevokePasskeyRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (RevokePasskeyRequest) HttpMethod() string                    { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/recovery/use-code", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type UseRecoveryCodeRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
	// @DataMember
	RecoveryCode string `json:"recoveryCode"`
}

func (UseRecoveryCodeRequest) CreateResponse() (r PasskeyRecoveryResponse) { return }
func (UseRecoveryCodeRequest) HttpMethod() string                          { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/recovery/magic-link/request", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RequestMagicLinkRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

func (RequestMagicLinkRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (RequestMagicLinkRequest) HttpMethod() string                    { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/recovery/magic-link/consume", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type ConsumeMagicLinkRequest struct {
	CodeMashRequestBase
	// @DataMember
	Token string `json:"token"`
}

func (ConsumeMagicLinkRequest) CreateResponse() (r PasskeyRecoveryResponse) { return }
func (ConsumeMagicLinkRequest) HttpMethod() string                          { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/has-passkey", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type HasPasskeyRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

func (HasPasskeyRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (HasPasskeyRequest) HttpMethod() string                    { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/email/start-verification", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type StartEmailVerificationRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
}

func (StartEmailVerificationRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (StartEmailVerificationRequest) HttpMethod() string                    { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/email/confirm-verification", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type ConfirmEmailVerificationRequest struct {
	CodeMashRequestBase
	// @DataMember
	Email string `json:"email"`
	// @DataMember
	Code string `json:"code"`
}

func (ConfirmEmailVerificationRequest) CreateResponse() (r PasskeyVerificationTokenResponse) { return }
func (ConfirmEmailVerificationRequest) HttpMethod() string                                   { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/registration-options", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type PasskeyRegistrationOptionsRequest struct {
	CodeMashRequestBase
	// @DataMember
	VerificationToken string `json:"verificationToken"`
}

func (PasskeyRegistrationOptionsRequest) CreateResponse() (r PasskeyCeremonyOptionsResponse) { return }
func (PasskeyRegistrationOptionsRequest) HttpMethod() string                                 { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/passkey/verify-registration", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type VerifyPasskeyRegistrationRequest struct {
	CodeMashRequestBase
	// @DataMember
	VerificationToken string `json:"verificationToken"`
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AttestationResponse string `json:"attestationResponse"`
	// @DataMember
	FriendlyName *string `json:"friendlyName,omitempty"`
}

func (VerifyPasskeyRegistrationRequest) CreateResponse() (r PasskeyAuthTokensResponse) { return }
func (VerifyPasskeyRegistrationRequest) HttpMethod() string                            { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/token/refresh", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type RefreshPasskeyTokenRequest struct {
	CodeMashRequestBase
	// @DataMember
	RefreshToken *string `json:"refreshToken,omitempty"`
}

func (RefreshPasskeyTokenRequest) CreateResponse() (r PasskeyAuthTokensResponse) { return }
func (RefreshPasskeyTokenRequest) HttpMethod() string                            { return "POST" }

/** @description Membership · Passkey */
// @Route("/{version}/membership/userauth/logout", "POST")
// @Api(Description="Membership · Passkey")
// @DataContract
type PasskeyLogoutRequest struct {
	CodeMashRequestBase
	// @DataMember
	RefreshToken *string `json:"refreshToken,omitempty"`
}

func (PasskeyLogoutRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (PasskeyLogoutRequest) HttpMethod() string                    { return "POST" }

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/merged-tree", "GET")
// @Api(Description="Database")
// @DataContract
type FindMergedTermTreeRequest struct {
	CodeMashRequestBase
	/** @description Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions. */
	// @DataMember
	// @ApiMember(Description="Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions.")
	AllLanguages bool `json:"allLanguages,omitempty"`
	/** @description The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…). */
	// @DataMember
	// @ApiMember(Description="The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…).", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (FindMergedTermTreeRequest) CreateResponse() (r FindMergedTermTreeResponse) { return }
func (FindMergedTermTreeRequest) HttpMethod() string                             { return "GET" }

/** @description Database */
// @Route("/{version}/database/taxonomies/tree", "GET")
// @Api(Description="Database")
// @DataContract
type FindTaxonomyTreeRequest struct {
	CodeMashRequestBase
	/** @description Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions. */
	// @DataMember
	// @ApiMember(Description="Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions.")
	AllLanguages bool `json:"allLanguages,omitempty"`
	// @DataMember
	IncludeTerms bool `json:"includeTerms,omitempty"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (FindTaxonomyTreeRequest) CreateResponse() (r FindTaxonomyTreeResponse) { return }
func (FindTaxonomyTreeRequest) HttpMethod() string                           { return "GET" }

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/terms", "GET")
// @Api(Description="Database")
// @DataContract
type FindTermsRequest struct {
	CodeMashListPaginationRequestBase
	/** @description Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions. */
	// @DataMember
	// @ApiMember(Description="Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions.")
	AllLanguages bool `json:"allLanguages,omitempty"`
	/** @description The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…). */
	// @DataMember
	// @ApiMember(Description="The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…).", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SortDescending bool `json:"sortDescending,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

func (FindTermsRequest) CreateResponse() (r FindTermsResponse) { return }
func (FindTermsRequest) HttpMethod() string                    { return "GET" }

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/terms/{parentId}/children", "GET")
// @Api(Description="Database")
// @DataContract
type FindTermsChildrenRequest struct {
	CodeMashListPaginationRequestBase
	/** @description Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions. */
	// @DataMember
	// @ApiMember(Description="Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions.")
	AllLanguages bool `json:"allLanguages,omitempty"`
	/** @description The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…). */
	// @DataMember
	// @ApiMember(Description="The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…).", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	ParentId string `json:"parentId"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

func (FindTermsChildrenRequest) CreateResponse() (r FindTermsChildrenResponse) { return }
func (FindTermsChildrenRequest) HttpMethod() string                            { return "GET" }

/** @description Database */
// @Route("/{version}/database/taxonomies/{taxonomyName}/terms/tree", "GET")
// @Api(Description="Database")
// @DataContract
type FindTermTreeRequest struct {
	CodeMashRequestBase
	/** @description Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions. */
	// @DataMember
	// @ApiMember(Description="Set true to keep name / description as stored even when Accept-Language is sent (editors). The full translations are always in names / descriptions.")
	AllLanguages bool `json:"allLanguages,omitempty"`
	/** @description The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…). */
	// @DataMember
	// @ApiMember(Description="The taxonomy: its name (slug or title; URL-encode spaces) or its id (txn_…).", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	RootTermId *string `json:"rootTermId,omitempty"`
	// @DataMember
	Depth *int `json:"depth,omitempty"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (FindTermTreeRequest) CreateResponse() (r FindTermTreeResponse) { return }
func (FindTermTreeRequest) HttpMethod() string                       { return "GET" }

/** @description Database */
// @Route("/{version}/database/schemas/{id}", "GET")
// @Api(Description="Database")
// @DataContract
type GetDatabaseSchemaRequest struct {
	CodeMashRequestBase
	// @DataMember
	Id string `json:"id"`
}

func (GetDatabaseSchemaRequest) CreateResponse() (r GetDatabaseSchemaResponse) { return }
func (GetDatabaseSchemaRequest) HttpMethod() string                            { return "GET" }

/** @description Database */
// @Route("/{version}/database/schemas", "GET")
// @Api(Description="Database")
// @DataContract
type GetDatabaseSchemasRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

func (GetDatabaseSchemasRequest) CreateResponse() (r GetDatabaseSchemasResponse) { return }
func (GetDatabaseSchemasRequest) HttpMethod() string                             { return "GET" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/aggregate", "POST")
// @Api(Description="Database")
// @DataContract
type AggregateRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Pipeline string `json:"pipeline"`
}

func (AggregateRequest) CreateResponse() (r AggregateResponse) { return }
func (AggregateRequest) HttpMethod() string                    { return "POST" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}/responsibility", "PUT")
// @Api(Description="Database")
// @DataContract
type ChangeResponsibilityRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	NewResponsibleUserId string `json:"newResponsibleUserId"`
}

func (ChangeResponsibilityRequest) CreateResponse() (r EmptyResponse) { return }
func (ChangeResponsibilityRequest) HttpMethod() string                { return "PUT" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/count", "GET")
// @Api(Description="Database")
// @DataContract
type CountRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
}

func (CountRequest) CreateResponse() (r CountResponse) { return }
func (CountRequest) HttpMethod() string                { return "GET" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "DELETE")
// @Api(Description="Database")
// @DataContract
type DeleteManyRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter string `json:"filter"`
	// @DataMember
	AllRecords *bool `json:"allRecords,omitempty"`
}

func (DeleteManyRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteManyRequest) HttpMethod() string                { return "DELETE" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "DELETE")
// @Api(Description="Database")
// @DataContract
type DeleteOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (DeleteOneRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteOneRequest) HttpMethod() string                { return "DELETE" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/distinct", "GET")
// @Api(Description="Database")
// @DataContract
type DistinctRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Field string `json:"field"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
}

func (DistinctRequest) CreateResponse() (r DistinctResponse) { return }
func (DistinctRequest) HttpMethod() string                   { return "GET" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/aggregates/{aggregateId}/execute", "POST")
// @Api(Description="Database")
// @DataContract
type ExecuteAggregateRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	AggregateId string `json:"aggregateId"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Tokens map[string]string `json:"tokens,omitempty"`
}

func (ExecuteAggregateRequest) CreateResponse() (r ExecuteAggregateResponse) { return }
func (ExecuteAggregateRequest) HttpMethod() string                           { return "POST" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}", "GET")
// @Api(Description="Database")
// @DataContract
type FindRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
	// @DataMember
	SortBy *string `json:"sortBy,omitempty"`
	// @DataMember
	SortOrder *int `json:"sortOrder,omitempty"`
	// @DataMember
	ExpandReferences bool `json:"expandReferences,omitempty"`
	// @DataMember
	AllLanguages bool `json:"allLanguages,omitempty"`
}

func (FindRequest) CreateResponse() (r FindResponse) { return }
func (FindRequest) HttpMethod() string               { return "GET" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "GET")
// @Api(Description="Database")
// @DataContract
type FindOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	ExpandReferences bool `json:"expandReferences,omitempty"`
	// @DataMember
	AllLanguages bool `json:"allLanguages,omitempty"`
}

func (FindOneRequest) CreateResponse() (r FindOneResponse) { return }
func (FindOneRequest) HttpMethod() string                  { return "GET" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/own", "GET")
// @Api(Description="Database")
// @DataContract
type FindOwnRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter *string `json:"filter,omitempty"`
	// @DataMember
	SchemaVersion *int `json:"schemaVersion,omitempty"`
	// @DataMember
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
	/** @description Set true to get every reference value as { id, display } (display = the target's displayField per the schema; null when the target is gone). Needs read permission on every source the schema links to (users, roles, taxonomy, collection, files) — otherwise the read is refused with CM-ERRORS-DATABASE-056 naming the source. Default false returns the stored ids. */
	// @DataMember
	// @ApiMember(Description="Set true to get every reference value as { id, display } (display = the target's displayField per the schema; null when the target is gone). Needs read permission on every source the schema links to (users, roles, taxonomy, collection, files) — otherwise the read is refused with CM-ERRORS-DATABASE-056 naming the source. Default false returns the stored ids.")
	ExpandReferences bool `json:"expandReferences,omitempty"`
	// @DataMember
	AllLanguages bool `json:"allLanguages,omitempty"`
}

func (FindOwnRequest) CreateResponse() (r FindResponse) { return }
func (FindOwnRequest) HttpMethod() string               { return "GET" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "POST")
// @Api(Description="Database")
// @DataContract
type InsertManyRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Documents string `json:"documents"`
}

func (InsertManyRequest) CreateResponse() (r EmptyResponse) { return }
func (InsertManyRequest) HttpMethod() string                { return "POST" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}", "POST")
// @Api(Description="Database")
// @DataContract
type InsertOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Document string `json:"document"`
}

func (InsertOneRequest) CreateResponse() (r IdResponse) { return }
func (InsertOneRequest) HttpMethod() string             { return "POST" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}/replace", "PUT")
// @Api(Description="Database")
// @DataContract
type ReplaceOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Replacement string `json:"replacement"`
}

func (ReplaceOneRequest) CreateResponse() (r EmptyResponse) { return }
func (ReplaceOneRequest) HttpMethod() string                { return "PUT" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "PUT")
// @Api(Description="Database")
// @DataContract
type UpdateManyRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Filter string `json:"filter"`
	// @DataMember
	AllRecords *bool `json:"allRecords,omitempty"`
	// @DataMember
	Update string `json:"update"`
	// @DataMember
	ArrayFilters *string `json:"arrayFilters,omitempty"`
}

func (UpdateManyRequest) CreateResponse() (r EmptyResponse) { return }
func (UpdateManyRequest) HttpMethod() string                { return "PUT" }

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "PUT")
// @Api(Description="Database")
// @DataContract
type UpdateOneRequest struct {
	CodeMashRequestBase
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Update string `json:"update"`
	// @DataMember
	ArrayFilters *string `json:"arrayFilters,omitempty"`
}

func (UpdateOneRequest) CreateResponse() (r EmptyResponse) { return }
func (UpdateOneRequest) HttpMethod() string                { return "PUT" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/commit", "POST")
// @Api(Description="Files")
// @DataContract
type CommitUploadRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	ContentType *string `json:"contentType,omitempty"`
	// @DataMember
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember
	FileName *string `json:"fileName,omitempty"`
	// @DataMember
	OwnerUserId *string `json:"ownerUserId,omitempty"`
	// @DataMember
	ReassignOwner *bool `json:"reassignOwner,omitempty"`
}

func (CommitUploadRequest) CreateResponse() (r EmptyResponse) { return }
func (CommitUploadRequest) HttpMethod() string                { return "POST" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/content", "GET")
// @Api(Description="Files")
// @DataContract
type GetFileContentRequest struct {
	RequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	Token *string `json:"token,omitempty"`
}

func (GetFileContentRequest) CreateResponse() (r []byte) { return }
func (GetFileContentRequest) HttpMethod() string         { return "GET" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/content", "PUT")
// @Api(Description="Files")
// @DataContract
type PutFileContentRequest struct {
	RequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	Token *string `json:"token,omitempty"`
}

func (PutFileContentRequest) CreateResponse() (r EmptyResponse) { return }
func (PutFileContentRequest) HttpMethod() string                { return "PUT" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}", "DELETE")
// @Api(Description="Files")
// @DataContract
type DeleteFileApiRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
}

func (DeleteFileApiRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteFileApiRequest) HttpMethod() string                { return "DELETE" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/bulk", "DELETE")
// @Api(Description="Files")
// @DataContract
type DeleteManyFilesApiRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember(Name="paths[]")
	Paths []string `json:"paths__"`
}

func (DeleteManyFilesApiRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteManyFilesApiRequest) HttpMethod() string                { return "DELETE" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/download", "GET")
// @Api(Description="Files")
// @DataContract
type DownloadFileApiRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
}

func (DownloadFileApiRequest) CreateResponse() (r []byte) { return }
func (DownloadFileApiRequest) HttpMethod() string         { return "GET" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/by-id/{id}", "GET")
// @Api(Description="Files")
// @DataContract
type GetFileByIdRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Id string `json:"id"`
}

func (GetFileByIdRequest) CreateResponse() (r GetFileByIdResponse) { return }
func (GetFileByIdRequest) HttpMethod() string                      { return "GET" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/info", "GET")
// @Api(Description="Files")
// @DataContract
type GetFileInfoRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
}

func (GetFileInfoRequest) CreateResponse() (r GetFileInfoResponse) { return }
func (GetFileInfoRequest) HttpMethod() string                      { return "GET" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/sign", "GET")
// @Api(Description="Files")
// @DataContract
type GetSignedUrlRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	ExpirationSeconds *int `json:"expirationSeconds,omitempty"`
}

func (GetSignedUrlRequest) CreateResponse() (r GetSignedUrlResponse) { return }
func (GetSignedUrlRequest) HttpMethod() string                       { return "GET" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}", "GET")
// @Api(Description="Files")
// @DataContract
type ListFilesRequest struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path *string `json:"path,omitempty"`
}

func (ListFilesRequest) CreateResponse() (r ListFilesResponse) { return }
func (ListFilesRequest) HttpMethod() string                    { return "GET" }

/** @description Files */
// @Route("/{version}/files/public/{PublicId}/{Name*}", "GET")
// @Api(Description="Files")
// @DataContract
type GetPublicFileRequest struct {
	RequestBase
	// @DataMember
	PublicId *string `json:"publicId,omitempty"`
	// @DataMember
	Name *string `json:"name,omitempty"`
}

func (GetPublicFileRequest) CreateResponse() (r []byte) { return }
func (GetPublicFileRequest) HttpMethod() string         { return "GET" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/register", "POST")
// @Api(Description="Files")
// @DataContract
type RegisterFilesRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Items []RegisterFileItem `json:"items,omitempty"`
	// @DataMember
	ReassignOwner *bool `json:"reassignOwner,omitempty"`
	// @DataMember
	DryRun *bool `json:"dryRun,omitempty"`
}

func (RegisterFilesRequest) CreateResponse() (r RegisterFilesResponse) { return }
func (RegisterFilesRequest) HttpMethod() string                        { return "POST" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/upload-url", "POST")
// @Api(Description="Files")
// @DataContract
type RequestUploadUrlRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	ContentType string `json:"contentType"`
	// @DataMember
	ExpirationSeconds *int `json:"expirationSeconds,omitempty"`
	// @DataMember
	OwnerUserId *string `json:"ownerUserId,omitempty"`
	// @DataMember
	ReassignOwner *bool `json:"reassignOwner,omitempty"`
}

func (RequestUploadUrlRequest) CreateResponse() (r RequestUploadUrlResponse) { return }
func (RequestUploadUrlRequest) HttpMethod() string                           { return "POST" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/use", "POST")
// @Api(Description="Files")
// @DataContract
type UseFileRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
	// @DataMember
	Path string `json:"path"`
	// @DataMember
	OwnerUserId *string `json:"ownerUserId,omitempty"`
	// @DataMember
	ReassignOwner *bool `json:"reassignOwner,omitempty"`
}

func (UseFileRequest) CreateResponse() (r IdResponse) { return }
func (UseFileRequest) HttpMethod() string             { return "POST" }

/** @description Registers or refreshes a push device for the signed-in user. */
// @Route("/{version}/notifications/push/devices", "POST")
// @Api(Description="Registers or refreshes a push device for the signed-in user.")
type RegisterOwnPushDeviceRequest struct {
	CodeMashRequestBase
	/** @description Who issued the token: fcm (Android / Firebase) or apns (iOS). The NATIVE device token — in an Expo app use getDevicePushTokenAsync(). Expo push tokens (provider 'expo') are refused. */
	// @ApiMember(Description="Who issued the token: fcm (Android / Firebase) or apns (iOS). The NATIVE device token — in an Expo app use getDevicePushTokenAsync(). Expo push tokens (provider 'expo') are refused.", IsRequired=true)
	Provider string `json:"provider"`
	/** @description The push token, exactly as the provider gave it. */
	// @ApiMember(Description="The push token, exactly as the provider gave it.", IsRequired=true)
	Token string `json:"token"`
	/** @description The app's platform: ios, android or web. */
	// @ApiMember(Description="The app's platform: ios, android or web.")
	Platform *string `json:"platform,omitempty"`
	/** @description Operating system name, e.g. iOS. */
	// @ApiMember(Description="Operating system name, e.g. iOS.")
	OsName *string `json:"osName,omitempty"`
	/** @description Operating system version, e.g. 18.1. */
	// @ApiMember(Description="Operating system version, e.g. 18.1.")
	OsVersion *string `json:"osVersion,omitempty"`
	/** @description The app version, e.g. 2.4.1. */
	// @ApiMember(Description="The app version, e.g. 2.4.1.")
	AppVersion *string `json:"appVersion,omitempty"`
	/** @description The device name the user sees, e.g. Anna's iPhone. */
	// @ApiMember(Description="The device name the user sees, e.g. Anna's iPhone.")
	DeviceName *string `json:"deviceName,omitempty"`
	/** @description The device model, e.g. iPhone 15. */
	// @ApiMember(Description="The device model, e.g. iPhone 15.")
	ModelName    *string `json:"modelName,omitempty"`
	Brand        *string `json:"brand,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	/** @description A stable id the app keeps for this installation (not the token). */
	// @ApiMember(Description="A stable id the app keeps for this installation (not the token).")
	InstallationId *string `json:"installationId,omitempty"`
	/** @description The device locale, e.g. lt-LT. */
	// @ApiMember(Description="The device locale, e.g. lt-LT.")
	Locale *string `json:"locale,omitempty"`
	/** @description The device time zone (IANA), e.g. Europe/Vilnius. */
	// @ApiMember(Description="The device time zone (IANA), e.g. Europe/Vilnius.")
	TimeZone *string `json:"timeZone,omitempty"`
}

func (RegisterOwnPushDeviceRequest) CreateResponse() (r IdResponse) { return }
func (RegisterOwnPushDeviceRequest) HttpMethod() string             { return "POST" }

/** @description Lists the signed-in user's own push devices. */
// @Route("/{version}/notifications/push/devices", "GET")
// @Api(Description="Lists the signed-in user's own push devices.")
type GetOwnPushDevicesRequest struct {
	CodeMashRequestBase
}

func (GetOwnPushDevicesRequest) CreateResponse() (r GetOwnPushDevicesResponse) { return }
func (GetOwnPushDevicesRequest) HttpMethod() string                            { return "GET" }

/** @description Removes one of the signed-in user's own push devices by id. */
// @Route("/{version}/notifications/push/devices/{Id}", "DELETE")
// @Api(Description="Removes one of the signed-in user's own push devices by id.")
type DeleteOwnPushDeviceRequest struct {
	CodeMashRequestBase
	/** @description The device id (pnd_…) from the devices list or the register answer. */
	// @ApiMember(Description="The device id (pnd_…) from the devices list or the register answer.", IsRequired=true)
	Id string `json:"id"`
}

func (DeleteOwnPushDeviceRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteOwnPushDeviceRequest) HttpMethod() string                { return "DELETE" }

/** @description Removes the signed-in user's push device registered with a token (sign-out). */
// @Route("/{version}/notifications/push/devices", "DELETE")
// @Api(Description="Removes the signed-in user's push device registered with a token (sign-out).")
type DeleteOwnPushDeviceByTokenRequest struct {
	CodeMashRequestBase
	/** @description The push token the device was registered with. */
	// @ApiMember(Description="The push token the device was registered with.", IsRequired=true)
	Token string `json:"token"`
}

func (DeleteOwnPushDeviceByTokenRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteOwnPushDeviceByTokenRequest) HttpMethod() string                { return "DELETE" }

/** @description Lists the push messages sent to the signed-in user. */
// @Route("/{version}/notifications/push/messages", "GET")
// @Api(Description="Lists the push messages sent to the signed-in user.")
type GetOwnPushMessagesRequest struct {
	CodeMashListPaginationRequestBase
	/** @description Only messages the caller has not marked read. */
	// @ApiMember(Description="Only messages the caller has not marked read.")
	UnreadOnly bool `json:"unreadOnly,omitempty"`
}

func (GetOwnPushMessagesRequest) CreateResponse() (r GetOwnPushMessagesResponse) { return }
func (GetOwnPushMessagesRequest) HttpMethod() string                             { return "GET" }

/** @description Marks one push message sent to the signed-in user read. */
// @Route("/{version}/notifications/push/messages/{Id}/read", "POST")
// @Api(Description="Marks one push message sent to the signed-in user read.")
type MarkOwnPushMessageReadRequest struct {
	CodeMashRequestBase
	/** @description The message id (ntf_…) from the messages list. */
	// @ApiMember(Description="The message id (ntf_…) from the messages list.", IsRequired=true)
	Id string `json:"id"`
}

func (MarkOwnPushMessageReadRequest) CreateResponse() (r EmptyResponse) { return }
func (MarkOwnPushMessageReadRequest) HttpMethod() string                { return "POST" }

/** @description Marks every push message sent to the signed-in user read. */
// @Route("/{version}/notifications/push/messages/read", "POST")
// @Api(Description="Marks every push message sent to the signed-in user read.")
type MarkAllOwnPushMessagesReadRequest struct {
	CodeMashRequestBase
}

func (MarkAllOwnPushMessagesReadRequest) CreateResponse() (r EmptyResponse) { return }
func (MarkAllOwnPushMessagesReadRequest) HttpMethod() string                { return "POST" }

/** @description Send your app's own log entries to the project's logs. */
// @Route("/{version}/logs", "POST")
// @Api(Description="Send your app's own log entries to the project's logs.")
type WriteLogs struct {
	CodeMashRequestBase
	/** @description 1–100 entries. */
	// @ApiMember(Description="1–100 entries.", IsRequired=true)
	Entries []WriteLogEntry `json:"entries"`
}

func (WriteLogs) CreateResponse() (r WriteLogsResponse) { return }
func (WriteLogs) HttpMethod() string                    { return "POST" }

/** @description Stores a payment order with a fixed price for one of the project's users. Server only (service user API key). */
// @Route("/{version}/payments/orders", "POST")
// @Api(Description="Stores a payment order with a fixed price for one of the project's users. Server only (service user API key).")
type CreatePaymentOrderRequest struct {
	CodeMashRequestBase
	/** @description The user the order is for (usr_…). */
	// @ApiMember(Description="The user the order is for (usr_…).", IsRequired=true)
	BuyerUserId string `json:"buyerUserId"`
	/** @description The project's own order reference, unique per project and environment (1–64 of A–Z a–z 0–9 . _ : -). */
	// @ApiMember(Description="The project's own order reference, unique per project and environment (1–64 of A–Z a–z 0–9 . _ : -).", IsRequired=true)
	MerchantReference string `json:"merchantReference"`
	/** @description 8–128 characters. Same key + same body = the stored order; same key + other body = 409. May be sent as the Idempotency-Key header instead. */
	// @ApiMember(Description="8–128 characters. Same key + same body = the stored order; same key + other body = 409. May be sent as the Idempotency-Key header instead.")
	IdempotencyKey *string `json:"idempotencyKey,omitempty"`
	/** @description EUR. */
	// @ApiMember(Description="EUR.", IsRequired=true)
	Currency string `json:"currency"`
	/** @description The version of the price list the server priced the lines with. */
	// @ApiMember(Description="The version of the price list the server priced the lines with.", IsRequired=true)
	PriceVersion string `json:"priceVersion"`
	/** @description 1–50 lines: sku, title, quantity, unitAmountMinor (cents). Norbix computes the amounts and the total. */
	// @ApiMember(Description="1–50 lines: sku, title, quantity, unitAmountMinor (cents). Norbix computes the amounts and the total.", IsRequired=true)
	Lines []PaymentOrderLineInput `json:"lines"`
	/** @description Up to 20 project values (e.g. adId). No personal or billing data. */
	// @ApiMember(Description="Up to 20 project values (e.g. adId). No personal or billing data.")
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (CreatePaymentOrderRequest) CreateResponse() (r IdResponse) { return }
func (CreatePaymentOrderRequest) HttpMethod() string             { return "POST" }

/** @description Reads one payment order. A buyer reads only their own orders. */
// @Route("/{version}/payments/orders/{Id}", "GET")
// @Api(Description="Reads one payment order. A buyer reads only their own orders.")
type GetPaymentOrderRequest struct {
	CodeMashRequestBase
	/** @description The order id (ord_…). */
	// @ApiMember(Description="The order id (ord_…).", IsRequired=true)
	Id string `json:"id"`
}

func (GetPaymentOrderRequest) CreateResponse() (r GetPaymentOrderResponse) { return }
func (GetPaymentOrderRequest) HttpMethod() string                          { return "GET" }

/** @description Lists payment orders, newest first. A buyer gets only their own orders. */
// @Route("/{version}/payments/orders", "GET")
// @Api(Description="Lists payment orders, newest first. A buyer gets only their own orders.")
type GetPaymentOrdersRequest struct {
	CodeMashListPaginationRequestBase
	/** @description Server only: one buyer's orders (usr_…). A buyer always gets their own. */
	// @ApiMember(Description="Server only: one buyer's orders (usr_…). A buyer always gets their own.")
	BuyerUserId *string `json:"buyerUserId,omitempty"`
	/** @description Only the order with this merchant reference. */
	// @ApiMember(Description="Only the order with this merchant reference.")
	MerchantReference *string `json:"merchantReference,omitempty"`
	/** @description Only orders in this status (Created, CheckoutStarted, Paid, …). */
	// @ApiMember(Description="Only orders in this status (Created, CheckoutStarted, Paid, …).")
	Status *string `json:"status,omitempty"`
}

func (GetPaymentOrdersRequest) CreateResponse() (r GetPaymentOrdersResponse) { return }
func (GetPaymentOrdersRequest) HttpMethod() string                           { return "GET" }

/** @description Opens a provider checkout (Montonio or Neopay) for a stored order and answers where the buyer pays. The buyer (own order) or the server. */
// @Route("/{version}/payments/orders/{Id}/checkout", "POST")
// @Api(Description="Opens a provider checkout (Montonio or Neopay) for a stored order and answers where the buyer pays. The buyer (own order) or the server.")
type StartPaymentCheckoutRequest struct {
	CodeMashRequestBase
	/** @description The order id (ord_…). */
	// @ApiMember(Description="The order id (ord_…).", IsRequired=true)
	Id string `json:"id"`
	/** @description The checkout integration (int_…, Montonio or Neopay) of the project in the request's environment. */
	// @ApiMember(Description="The checkout integration (int_…, Montonio or Neopay) of the project in the request's environment.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
	/** @description bankLink (default), card, applePay or googlePay (the wallets: Montonio only). */
	// @ApiMember(Description="bankLink (default), card, applePay or googlePay (the wallets: Montonio only).")
	Method *string `json:"method,omitempty"`
	/** @description Where the buyer comes back after paying; its origin must be one of the project's allowed origins. */
	// @ApiMember(Description="Where the buyer comes back after paying; its origin must be one of the project's allowed origins.", IsRequired=true)
	ReturnUrl string `json:"returnUrl"`
	/** @description Bank links only: the country whose banks are shown first (default LT). */
	// @ApiMember(Description="Bank links only: the country whose banks are shown first (default LT).")
	PreferredCountry *string `json:"preferredCountry,omitempty"`
	/** @description Bank links only: one bank to open directly (the provider's bank code). */
	// @ApiMember(Description="Bank links only: one bank to open directly (the provider's bank code).")
	PreferredProvider *string `json:"preferredProvider,omitempty"`
}

func (StartPaymentCheckoutRequest) CreateResponse() (r StartPaymentCheckoutResponse) { return }
func (StartPaymentCheckoutRequest) HttpMethod() string                               { return "POST" }

// @Route("/{version}/code/marketplace/functions/{FunctionViewId}/invoke", "POST")
// @DataContract
type InvokeMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Function view id (func_…), from get_marketplace_functions. */
	// @DataMember
	// @ApiMember(Description="Function view id (func_…), from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
	// @DataMember
	Payload map[string]interface{} `json:"payload"`
}

func (InvokeMarketplaceFunction) CreateResponse() (r InvokeMarketplaceFunctionResponse) { return }
func (InvokeMarketplaceFunction) HttpMethod() string                                    { return "POST" }

// @Route("/{version}/code/workflows/{Id}/invoke", "POST")
// @DataContract
type InvokeCodeWorkflow struct {
	CodeMashRequestBase
	/** @description wf_… */
	// @DataMember
	// @ApiMember(Description="wf_…", IsRequired=true)
	Id string `json:"id"`
	/** @description The @Model.Request object; checked against the definition's inputs when it has one. */
	// @DataMember
	// @ApiMember(Description="The @Model.Request object; checked against the definition's inputs when it has one.")
	Payload map[string]interface{} `json:"payload,omitempty"`
	/** @description 0–30 seconds to wait for the run; default 10; 0 answers 202 at once. */
	// @DataMember
	// @ApiMember(DataType="integer", Description="0–30 seconds to wait for the run; default 10; 0 answers 202 at once.")
	WaitSeconds *int `json:"waitSeconds,omitempty"`
	/** @description "draft" to test the draft, a version number to pin one; absent = the latest published. */
	// @DataMember
	// @ApiMember(Description="\"draft\" to test the draft, a version number to pin one; absent = the latest published.")
	RunVersion *string `json:"runVersion,omitempty"`
}

func (InvokeCodeWorkflow) CreateResponse() (r InvokeCodeWorkflowResponse) { return }
func (InvokeCodeWorkflow) HttpMethod() string                             { return "POST" }

// @Route("/{version}/code/workflows/runs", "GET")
// @DataContract
type GetCodeWorkflowWaitingRuns struct {
	CodeMashListPaginationRequestBase
	/** @description running | waiting (default) | succeeded | failed | cancelled */
	// @DataMember
	// @ApiMember(Description="running | waiting (default) | succeeded | failed | cancelled")
	Status *string `json:"status,omitempty"`
	/** @description approval | event | delay | child */
	// @DataMember
	// @ApiMember(Description="approval | event | delay | child")
	WaitingKind *string `json:"waitingKind,omitempty"`
	/** @description "me" or a usr_… id: only approvals assigned to that user. Without code:read only "me" (the default then). */
	// @DataMember
	// @ApiMember(Description="\"me\" or a usr_… id: only approvals assigned to that user. Without code:read only \"me\" (the default then).")
	Assignee *string `json:"assignee,omitempty"`
	/** @description Only runs of this workflow (wf_…). */
	// @DataMember
	// @ApiMember(Description="Only runs of this workflow (wf_…).")
	WorkflowId *string `json:"workflowId,omitempty"`
	/** @description The wait's rendered key, exact match. */
	// @DataMember
	// @ApiMember(Description="The wait's rendered key, exact match.")
	Key *string `json:"key,omitempty"`
}

func (GetCodeWorkflowWaitingRuns) CreateResponse() (r WorkflowRunsResponse) { return }
func (GetCodeWorkflowWaitingRuns) HttpMethod() string                       { return "GET" }

// @Route("/{version}/code/workflows/runs/{RunId}/resume", "POST")
// @DataContract
type ResumeCodeWorkflowRun struct {
	CodeMashRequestBase
	/** @description wfr_… */
	// @DataMember
	// @ApiMember(Description="wfr_…", IsRequired=true)
	RunId string `json:"runId"`
	/** @description Required when several steps wait. */
	// @DataMember
	// @ApiMember(Description="Required when several steps wait.")
	StepId *string `json:"stepId,omitempty"`
	/** @description approved | rejected — required for an approval wait. */
	// @DataMember
	// @ApiMember(Description="approved | rejected — required for an approval wait.")
	Decision *string `json:"decision,omitempty"`
	// @DataMember
	Payload map[string]interface{} `json:"payload,omitempty"`
	/** @description Optional: why — one or two sentences. Kept with the answer and shown on the run (at most 2 000 characters). */
	// @DataMember
	// @ApiMember(Description="Optional: why — one or two sentences. Kept with the answer and shown on the run (at most 2 000 characters).")
	Reason *string `json:"reason,omitempty"`
}

func (ResumeCodeWorkflowRun) CreateResponse() (r WorkflowRunResponse) { return }
func (ResumeCodeWorkflowRun) HttpMethod() string                      { return "POST" }

/** @description Beta: what a workflow's one-click link will do — no sign-in, nothing changes. POST the same URL to answer. */
// @Route("/{version}/code/workflows/resume/{Token}", "GET")
// @Api(Description="Beta: what a workflow's one-click link will do — no sign-in, nothing changes. POST the same URL to answer.")
// @DataContract
type OpenCodeWorkflowResumeLink struct {
	RequestBase
	/** @description The signed link token from the wait's notification (@Model.Wait.ApproveUrl / RejectUrl / ResumeUrl). */
	// @DataMember
	// @ApiMember(Description="The signed link token from the wait's notification (@Model.Wait.ApproveUrl / RejectUrl / ResumeUrl).", IsRequired=true)
	Token string `json:"token"`
}

func (OpenCodeWorkflowResumeLink) CreateResponse() (r CodeWorkflowResumeLinkResponse) { return }
func (OpenCodeWorkflowResumeLink) HttpMethod() string                                 { return "GET" }

/** @description Beta: answers a workflow's waiting step with a one-click link — no sign-in; the run goes on as the workflow's execution identity. */
// @Route("/{version}/code/workflows/resume/{Token}", "POST")
// @Api(Description="Beta: answers a workflow's waiting step with a one-click link — no sign-in; the run goes on as the workflow's execution identity.")
// @DataContract
type UseCodeWorkflowResumeLink struct {
	RequestBase
	/** @description The signed link token. */
	// @DataMember
	// @ApiMember(Description="The signed link token.", IsRequired=true)
	Token string `json:"token"`
}

func (UseCodeWorkflowResumeLink) CreateResponse() (r WorkflowRunResponse) { return }
func (UseCodeWorkflowResumeLink) HttpMethod() string                      { return "POST" }

// @Route("/{version}/code/workflows/{Id}/runs", "GET")
// @DataContract
type GetCodeWorkflowRuns struct {
	CodeMashListPaginationRequestBase
	/** @description wf_… */
	// @DataMember
	// @ApiMember(Description="wf_…", IsRequired=true)
	Id string `json:"id"`
	/** @description running | waiting | succeeded | failed | cancelled */
	// @DataMember
	// @ApiMember(Description="running | waiting | succeeded | failed | cancelled")
	Status *string `json:"status,omitempty"`
	// @DataMember
	From *time.Time `json:"from,omitempty"`
	// @DataMember
	To *time.Time `json:"to,omitempty"`
}

func (GetCodeWorkflowRuns) CreateResponse() (r WorkflowRunsResponse) { return }
func (GetCodeWorkflowRuns) HttpMethod() string                       { return "GET" }

// @Route("/{version}/code/workflows/runs/{RunId}", "GET")
// @DataContract
type GetCodeWorkflowRun struct {
	CodeMashRequestBase
	/** @description wfr_… */
	// @DataMember
	// @ApiMember(Description="wfr_…", IsRequired=true)
	RunId string `json:"runId"`
}

func (GetCodeWorkflowRun) CreateResponse() (r WorkflowRunResponse) { return }
func (GetCodeWorkflowRun) HttpMethod() string                      { return "GET" }

// @Route("/{version}/code/workflows/runs/{RunId}/logs", "GET")
// @DataContract
type GetCodeWorkflowRunLogs struct {
	CodeMashListPaginationRequestBase
	/** @description wfr_… */
	// @DataMember
	// @ApiMember(Description="wfr_…", IsRequired=true)
	RunId string `json:"runId"`
	/** @description Only this step's lines. */
	// @DataMember
	// @ApiMember(Description="Only this step's lines.")
	StepId *string `json:"stepId,omitempty"`
	/** @description Information | Warning | Error */
	// @DataMember
	// @ApiMember(Description="Information | Warning | Error")
	Level *string `json:"level,omitempty"`
}

func (GetCodeWorkflowRunLogs) CreateResponse() (r WorkflowRunLogsResponse) { return }
func (GetCodeWorkflowRunLogs) HttpMethod() string                          { return "GET" }

// @Route("/{version}/code/workflows/{Id}/resume", "POST")
// @DataContract
type ResumeCodeWorkflowRunByKey struct {
	CodeMashRequestBase
	/** @description wf_… */
	// @DataMember
	// @ApiMember(Description="wf_…", IsRequired=true)
	Id string `json:"id"`
	/** @description The wait's rendered key. */
	// @DataMember
	// @ApiMember(Description="The wait's rendered key.", IsRequired=true)
	Key string `json:"key"`
	// @DataMember
	StepId *string `json:"stepId,omitempty"`
	/** @description approved | rejected — required for an approval wait. */
	// @DataMember
	// @ApiMember(Description="approved | rejected — required for an approval wait.")
	Decision *string `json:"decision,omitempty"`
	// @DataMember
	Payload map[string]interface{} `json:"payload,omitempty"`
	/** @description Optional: why — one or two sentences. Kept with the answer and shown on the run (at most 2 000 characters). */
	// @DataMember
	// @ApiMember(Description="Optional: why — one or two sentences. Kept with the answer and shown on the run (at most 2 000 characters).")
	Reason *string `json:"reason,omitempty"`
}

func (ResumeCodeWorkflowRunByKey) CreateResponse() (r WorkflowRunResponse) { return }
func (ResumeCodeWorkflowRunByKey) HttpMethod() string                      { return "POST" }

// @Route("/{version}/code/workflows/runs/{RunId}/retry", "POST")
// @DataContract
type RetryCodeWorkflowRunStep struct {
	CodeMashRequestBase
	/** @description wfr_… */
	// @DataMember
	// @ApiMember(Description="wfr_…", IsRequired=true)
	RunId string `json:"runId"`
	/** @description The failed step; default = the step named in the run's error. */
	// @DataMember
	// @ApiMember(Description="The failed step; default = the step named in the run's error.")
	StepId *string `json:"stepId,omitempty"`
}

func (RetryCodeWorkflowRunStep) CreateResponse() (r WorkflowRunResponse) { return }
func (RetryCodeWorkflowRunStep) HttpMethod() string                      { return "POST" }

// @Route("/{version}/code/workflows/runs/{RunId}/cancel", "POST")
// @DataContract
type CancelCodeWorkflowRun struct {
	CodeMashRequestBase
	/** @description wfr_… */
	// @DataMember
	// @ApiMember(Description="wfr_…", IsRequired=true)
	RunId string `json:"runId"`
	// @DataMember
	Reason *string `json:"reason,omitempty"`
}

func (CancelCodeWorkflowRun) CreateResponse() (r WorkflowRunResponse) { return }
func (CancelCodeWorkflowRun) HttpMethod() string                      { return "POST" }

// @Route("/{version}/notifications/inbox", "GET")
type GetNotificationInboxRequest struct {
	CodeMashListPaginationRequestBase
	UnreadOnly bool    `json:"unreadOnly,omitempty"`
	Channel    *string `json:"channel,omitempty"`
}

func (GetNotificationInboxRequest) CreateResponse() (r GetNotificationInboxResponse) { return }
func (GetNotificationInboxRequest) HttpMethod() string                               { return "GET" }

// @Route("/{version}/notifications/inbox/unread-count", "GET")
type GetUnreadCountRequest struct {
	CodeMashRequestBase
}

func (GetUnreadCountRequest) CreateResponse() (r GetUnreadCountResponse) { return }
func (GetUnreadCountRequest) HttpMethod() string                         { return "GET" }

// @Route("/{version}/notifications/inbox/{notificationId}/read", "POST")
type MarkNotificationReadRequest struct {
	CodeMashRequestBase
	NotificationId string `json:"notificationId"`
}

func (MarkNotificationReadRequest) CreateResponse() (r MarkNotificationReadResponse) { return }
func (MarkNotificationReadRequest) HttpMethod() string                               { return "POST" }

// @Route("/{version}/notifications/inbox/read-all", "POST")
type MarkAllNotificationsReadRequest struct {
	CodeMashRequestBase
}

func (MarkAllNotificationsReadRequest) CreateResponse() (r MarkAllNotificationsReadResponse) { return }
func (MarkAllNotificationsReadRequest) HttpMethod() string                                   { return "POST" }

/** @description Sets or removes one user's own end-user AI plan */
// @Route("/{version}/membership/users/{userId}/ai-plan", "PUT")
// @Api(Description="Sets or removes one user's own end-user AI plan")
type SetUserAiPlan struct {
	CodeMashRequestBase
	/** @description The human user id (ct_…). */
	// @ApiMember(Description="The human user id (ct_…).", IsRequired=true)
	UserId string `json:"userId"`
	/** @description The plan id (aip_…). Empty removes the user's own plan. */
	// @ApiMember(Description="The plan id (aip_…). Empty removes the user's own plan.")
	PlanId *string `json:"planId,omitempty"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (SetUserAiPlan) CreateResponse() (r EmptyResponse) { return }
func (SetUserAiPlan) HttpMethod() string                { return "PUT" }

/** @description Clear a contact's hard bounce */
// @Route("/{version}/membership/users/{contactId}/marketing-state/{channel}/clear-bounce", "POST")
// @Api(Description="Clear a contact's hard bounce")
type ClearContactBounceRequest struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) whose hard bounce is cleared. */
	// @ApiMember(Description="The contact id (ct_…) whose hard bounce is cleared.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Delivery channel: Email, Sms, or Push. */
	// @ApiMember(Description="Delivery channel: Email, Sms, or Push.", IsRequired=true)
	Channel string `json:"channel"`
}

func (ClearContactBounceRequest) CreateResponse() (r EmptyResponse) { return }
func (ClearContactBounceRequest) HttpMethod() string                { return "POST" }

/** @description Create a contact */
// @Route("/{version}/membership/users", "POST")
// @Api(Description="Create a contact")
type CreateContactRequest struct {
	CodeMashRequestBase
	/** @description Primary email address of the contact (optional if a phone is given). */
	// @ApiMember(Description="Primary email address of the contact (optional if a phone is given).")
	PrimaryEmail *string `json:"primaryEmail,omitempty"`
	/** @description Primary phone number in international format, e.g. +14155550123 (optional if an email is given). */
	// @ApiMember(Description="Primary phone number in international format, e.g. +14155550123 (optional if an email is given).")
	PrimaryPhone *string `json:"primaryPhone,omitempty"`
	/** @description Display name shown in the dashboard (optional). */
	// @ApiMember(Description="Display name shown in the dashboard (optional).")
	DisplayName *string `json:"displayName,omitempty"`
	/** @description Contact's first name (optional). */
	// @ApiMember(Description="Contact's first name (optional).")
	FirstName *string `json:"firstName,omitempty"`
	/** @description Contact's last name (optional). */
	// @ApiMember(Description="Contact's last name (optional).")
	LastName *string `json:"lastName,omitempty"`
	/** @description Tags to put on the new contact — free labels campaigns and segments select by, normalised like any tag. */
	// @ApiMember(Description="Tags to put on the new contact — free labels campaigns and segments select by, normalised like any tag.")
	Tags []string `json:"tags,omitempty"`
}

func (CreateContactRequest) CreateResponse() (r EmptyResponse) { return }
func (CreateContactRequest) HttpMethod() string                { return "POST" }

/** @description Archive a contact */
// @Route("/{version}/membership/users/{contactId}", "DELETE")
// @Api(Description="Archive a contact")
type DeleteContact struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to archive. Get it from get_all_contacts. */
	// @ApiMember(Description="The contact id (ct_…) to archive. Get it from get_all_contacts.", IsRequired=true)
	ContactId string `json:"contactId"`
}

func (DeleteContact) CreateResponse() (r EmptyResponse) { return }
func (DeleteContact) HttpMethod() string                { return "DELETE" }

/** @description Get a contact */
// @Route("/{version}/membership/users/{contactId}", "GET")
// @Api(Description="Get a contact")
type GetContact struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to fetch, or a login id (usr_…) — then the contact that owns that login is returned. Get it from get_all_contacts. */
	// @ApiMember(Description="The contact id (ct_…) to fetch, or a login id (usr_…) — then the contact that owns that login is returned. Get it from get_all_contacts.", IsRequired=true)
	ContactId string `json:"contactId"`
}

func (GetContact) CreateResponse() (r GetContactResponse) { return }
func (GetContact) HttpMethod() string                     { return "GET" }

/** @description List contacts */
// @Route("/{version}/membership/users", "GET")
// @Api(Description="List contacts")
type GetAllContacts struct {
	CodeMashRequestBase
	/** @description Cursor for the next page: pass the nextCursor from the previous call. Omit for the first page. */
	// @ApiMember(Description="Cursor for the next page: pass the nextCursor from the previous call. Omit for the first page.")
	StartingAfter *string `json:"startingAfter,omitempty"`
	/** @description How many contacts to return per page (default 50). */
	// @ApiMember(Description="How many contacts to return per page (default 50).")
	PageSize *int `json:"pageSize,omitempty"`
	/** @description Free text, case-insensitive: matches the email, phone, display / first / last / full name, or any login's user name or email; an exact contact id (ct_…) or login id (usr_…) finds that contact. */
	// @ApiMember(Description="Free text, case-insensitive: matches the email, phone, display / first / last / full name, or any login's user name or email; an exact contact id (ct_…) or login id (usr_…) finds that contact.")
	Search *string `json:"search,omitempty"`
	/** @description Case-insensitive part of the email (primary email or any login's email). */
	// @ApiMember(Description="Case-insensitive part of the email (primary email or any login's email).")
	Email *string `json:"email,omitempty"`
	/** @description Case-insensitive part of any login's user name. */
	// @ApiMember(Description="Case-insensitive part of any login's user name.")
	UserName *string `json:"userName,omitempty"`
	/** @description Login status names, any of: Registered, PendingValidation, Active, Unregistered, Suspended, InActive, Blocked. */
	// @ApiMember(Description="Login status names, any of: Registered, PendingValidation, Active, Unregistered, Suspended, InActive, Blocked.")
	Statuses []string `json:"statuses,omitempty"`
	/** @description Login type names, any of: Service, Email, UserName, Phone, Guest, Social. */
	// @ApiMember(Description="Login type names, any of: Service, Email, UserName, Phone, Guest, Social.")
	Types []string `json:"types,omitempty"`
	/** @description Composite project-role ids (pr_{project}_nr_{role}); the contact holds any of them. */
	// @ApiMember(Description="Composite project-role ids (pr_{project}_nr_{role}); the contact holds any of them.")
	RoleNames []string `json:"roleNames,omitempty"`
	/** @description Contact tags; the contact carries any of them. */
	// @ApiMember(Description="Contact tags; the contact carries any of them.")
	Tags []string `json:"tags,omitempty"`
}

func (GetAllContacts) CreateResponse() (r GetAllContactsResponse) { return }
func (GetAllContacts) HttpMethod() string                         { return "GET" }

/** @description Merge contacts */
// @Route("/{version}/membership/users/merge", "POST")
// @Api(Description="Merge contacts")
type MergeContactsRequest struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) that will remain after the merge (the survivor). */
	// @ApiMember(Description="The contact id (ct_…) that will remain after the merge (the survivor).", IsRequired=true)
	SurvivorId string `json:"survivorId"`
	/** @description The contact ids (ct_…) to merge into the survivor and archive. At least one. */
	// @ApiMember(Description="The contact ids (ct_…) to merge into the survivor and archive. At least one.", IsRequired=true)
	MergedIds []string `json:"mergedIds"`
}

func (MergeContactsRequest) CreateResponse() (r EmptyResponse) { return }
func (MergeContactsRequest) HttpMethod() string                { return "POST" }

/** @description Repair contact roles from their logins */
// @Route("/{version}/membership/users/repair-roles", "POST")
// @Api(Description="Repair contact roles from their logins")
type RepairContactRolesRequest struct {
	CodeMashRequestBase
	/** @description True = only count what would change. False = write. Empty = true (only count). */
	// @ApiMember(Description="True = only count what would change. False = write. Empty = true (only count).")
	DryRun *bool `json:"dryRun,omitempty"`
}

func (RepairContactRolesRequest) CreateResponse() (r RepairContactRolesResponse) { return }
func (RepairContactRolesRequest) HttpMethod() string                             { return "POST" }

/** @description Update a contact */
// @Route("/{version}/membership/users/{contactId}", "PATCH")
// @Api(Description="Update a contact")
type UpdateContactRequest struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to update. Get it from get_all_contacts. */
	// @ApiMember(Description="The contact id (ct_…) to update. Get it from get_all_contacts.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description Display name shown in the dashboard. */
	// @ApiMember(Description="Display name shown in the dashboard.")
	DisplayName *string `json:"displayName,omitempty"`
	/** @description First name. */
	// @ApiMember(Description="First name.")
	FirstName *string `json:"firstName,omitempty"`
	/** @description Last name. */
	// @ApiMember(Description="Last name.")
	LastName *string `json:"lastName,omitempty"`
	/** @description Full name (overrides first/last when set). */
	// @ApiMember(Description="Full name (overrides first/last when set).")
	FullName *string `json:"fullName,omitempty"`
	/** @description Company or organisation name. */
	// @ApiMember(Description="Company or organisation name.")
	Company *string `json:"company,omitempty"`
	/** @description Free-text internal notes about the contact. */
	// @ApiMember(Description="Free-text internal notes about the contact.")
	Notes *string `json:"notes,omitempty"`
	/** @description Gender: Male, Female or Other. */
	// @ApiMember(Description="Gender: Male, Female or Other.")
	Gender *string `json:"gender,omitempty"`
	/** @description Birth date as a unix timestamp in MILLISECONDS (UTC). */
	// @ApiMember(Description="Birth date as a unix timestamp in MILLISECONDS (UTC).")
	BirthDate *int64 `json:"birthDate,omitempty"`
	/** @description IANA time zone id, e.g. Europe/Vilnius. */
	// @ApiMember(Description="IANA time zone id, e.g. Europe/Vilnius.")
	TimeZone *string `json:"timeZone,omitempty"`
	/** @description Preferred language/locale code, e.g. en or en-US. */
	// @ApiMember(Description="Preferred language/locale code, e.g. en or en-US.")
	Language *string `json:"language,omitempty"`
	/** @description Address line 1 (street). */
	// @ApiMember(Description="Address line 1 (street).")
	AddressLine1 *string `json:"addressLine1,omitempty"`
	/** @description Address line 2 (apartment, suite, etc.). */
	// @ApiMember(Description="Address line 2 (apartment, suite, etc.).")
	AddressLine2 *string `json:"addressLine2,omitempty"`
	/** @description Country name or code. */
	// @ApiMember(Description="Country name or code.")
	Country *string `json:"country,omitempty"`
	/** @description City. */
	// @ApiMember(Description="City.")
	City *string `json:"city,omitempty"`
	/** @description State, region or province. */
	// @ApiMember(Description="State, region or province.")
	State *string `json:"state,omitempty"`
	/** @description Postal or ZIP code. */
	// @ApiMember(Description="Postal or ZIP code.")
	PostalCode *string `json:"postalCode,omitempty"`
}

func (UpdateContactRequest) CreateResponse() (r EmptyResponse) { return }
func (UpdateContactRequest) HttpMethod() string                { return "PATCH" }

/** @description Link a login to a contact */
// @Route("/{version}/membership/users/{contactId}/identities", "POST")
// @Api(Description="Link a login to a contact")
type AddContactIdentityRequest struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to link the login to. */
	// @ApiMember(Description="The contact id (ct_…) to link the login to.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description The login (identity) id (usr_…) to link to the contact. */
	// @ApiMember(Description="The login (identity) id (usr_…) to link to the contact.", IsRequired=true)
	AuthId string `json:"authId"`
}

func (AddContactIdentityRequest) CreateResponse() (r EmptyResponse) { return }
func (AddContactIdentityRequest) HttpMethod() string                { return "POST" }

/** @description Make a login the contact's primary */
// @Route("/{version}/membership/users/{contactId}/identities/{authId}/promote", "POST")
// @Api(Description="Make a login the contact's primary")
type PromoteContactIdentityRequest struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) whose login is being promoted. */
	// @ApiMember(Description="The contact id (ct_…) whose login is being promoted.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description The linked login (identity) id (usr_…) to make primary. */
	// @ApiMember(Description="The linked login (identity) id (usr_…) to make primary.", IsRequired=true)
	AuthId string `json:"authId"`
}

func (PromoteContactIdentityRequest) CreateResponse() (r EmptyResponse) { return }
func (PromoteContactIdentityRequest) HttpMethod() string                { return "POST" }

/** @description Unlink a login from a contact */
// @Route("/{version}/membership/users/{contactId}/identities/{authId}", "DELETE")
// @Api(Description="Unlink a login from a contact")
type RemoveContactIdentityRequest struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to unlink the login from. */
	// @ApiMember(Description="The contact id (ct_…) to unlink the login from.", IsRequired=true)
	ContactId string `json:"contactId"`
	/** @description The login (identity) id (usr_…) to unlink from the contact. */
	// @ApiMember(Description="The login (identity) id (usr_…) to unlink from the contact.", IsRequired=true)
	AuthId string `json:"authId"`
}

func (RemoveContactIdentityRequest) CreateResponse() (r EmptyResponse) { return }
func (RemoveContactIdentityRequest) HttpMethod() string                { return "DELETE" }

// @Route("/{version}/files/item/private", "POST")
type MakeFilePrivateRequest struct {
	CodeMashRequestBase
	/** @description The files integration the file lives on. */
	// @ApiMember(Description="The files integration the file lives on.", IsRequired=true)
	FilesIntegrationId string `json:"filesIntegrationId"`
	/** @description Path of the file, relative to the integration. */
	// @ApiMember(Description="Path of the file, relative to the integration.", IsRequired=true)
	Path string `json:"path"`
}

func (MakeFilePrivateRequest) CreateResponse() (r EmptyResponse) { return }
func (MakeFilePrivateRequest) HttpMethod() string                { return "POST" }

// @Route("/{version}/files/item/public", "POST")
type MakeFilePublicRequest struct {
	CodeMashRequestBase
	/** @description The files integration the file lives on. */
	// @ApiMember(Description="The files integration the file lives on.", IsRequired=true)
	FilesIntegrationId string `json:"filesIntegrationId"`
	/** @description Path of the file to publish, relative to the integration. */
	// @ApiMember(Description="Path of the file to publish, relative to the integration.", IsRequired=true)
	Path string `json:"path"`
}

func (MakeFilePublicRequest) CreateResponse() (r IdResponse) { return }
func (MakeFilePublicRequest) HttpMethod() string             { return "POST" }

// @Route("/{version}/files/folder/private", "POST")
type MakeFolderPrivateRequest struct {
	CodeMashRequestBase
	/** @description The files integration the folder lives on. */
	// @ApiMember(Description="The files integration the folder lives on.", IsRequired=true)
	FilesIntegrationId string `json:"filesIntegrationId"`
	/** @description Folder prefix, relative to the integration. */
	// @ApiMember(Description="Folder prefix, relative to the integration.", IsRequired=true)
	Path string `json:"path"`
}

func (MakeFolderPrivateRequest) CreateResponse() (r EmptyResponse) { return }
func (MakeFolderPrivateRequest) HttpMethod() string                { return "POST" }

// @Route("/{version}/files/folder/public", "POST")
type MakeFolderPublicRequest struct {
	CodeMashRequestBase
	/** @description The files integration the folder lives on. */
	// @ApiMember(Description="The files integration the folder lives on.", IsRequired=true)
	FilesIntegrationId string `json:"filesIntegrationId"`
	/** @description Folder prefix to publish, relative to the integration. */
	// @ApiMember(Description="Folder prefix to publish, relative to the integration.", IsRequired=true)
	Path string `json:"path"`
}

func (MakeFolderPublicRequest) CreateResponse() (r IdResponse) { return }
func (MakeFolderPublicRequest) HttpMethod() string             { return "POST" }

// @Route("/{version}/files/item", "GET")
type GetFile struct {
	CodeMashRequestBase
	/** @description The files integration id to read from, from get_files_integrations. */
	// @ApiMember(Description="The files integration id to read from, from get_files_integrations.", IsRequired=true)
	FilesIntegrationId string `json:"filesIntegrationId"`
	/** @description The path of the file to fetch metadata for. */
	// @ApiMember(Description="The path of the file to fetch metadata for.", IsRequired=true)
	Path string `json:"path"`
}

func (GetFile) CreateResponse() (r GetFileResponse) { return }
func (GetFile) HttpMethod() string                  { return "GET" }

// @Route("/{version}/files/folder", "GET")
type GetFolderFiles struct {
	CodeMashListPaginationRequestBase
	/** @description The files integration id to list from, from get_files_integrations. */
	// @ApiMember(Description="The files integration id to list from, from get_files_integrations.", IsRequired=true)
	FilesIntegrationId string `json:"filesIntegrationId"`
	/** @description Path prefix to list. Empty / null lists the root. */
	// @ApiMember(Description="Path prefix to list. Empty / null lists the root.")
	Path *string `json:"path,omitempty"`
}

func (GetFolderFiles) CreateResponse() (r GetFolderFilesResponse) { return }
func (GetFolderFiles) HttpMethod() string                         { return "GET" }

/** @description Database */
// @Route("/{version}/database/imports", "POST")
// @Api(Description="Database")
type CreateCollectionImport struct {
	CodeMashRequestBase
	/** @description The uploaded CSV's file ref, from the upload call. */
	// @ApiMember(Description="The uploaded CSV's file ref, from the upload call.", IsRequired=true)
	File FileResourceRefDto `json:"file"`
	/** @description The target schema id. */
	// @ApiMember(Description="The target schema id.", IsRequired=true)
	SchemaId string `json:"schemaId"`
	/** @description The target collection (schema) name. */
	// @ApiMember(Description="The target collection (schema) name.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The CSV delimiter used at upload time. */
	// @ApiMember(Description="The CSV delimiter used at upload time.", IsRequired=true)
	Delimiter string `json:"delimiter"`
	HasHeader bool   `json:"hasHeader,omitempty"`
	/** @description Column → property mapping, frozen for this import. */
	// @ApiMember(Description="Column → property mapping, frozen for this import.", IsRequired=true)
	Mapping []ImportColumnMappingDto `json:"mapping"`
}

func (CreateCollectionImport) CreateResponse() (r IdResponse) { return }
func (CreateCollectionImport) HttpMethod() string             { return "POST" }

/** @description Database */
// @Route("/{version}/database/imports/{Id}", "DELETE")
// @Api(Description="Database")
type DeleteCollectionImportRequest struct {
	CodeMashRequestBase
	/** @description The import id (imp_…). */
	// @ApiMember(Description="The import id (imp_…).", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (DeleteCollectionImportRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteCollectionImportRequest) HttpMethod() string                { return "DELETE" }

/** @description Database */
// @Route("/{version}/database/imports/{Id}", "GET")
// @Api(Description="Database")
type GetCollectionImport struct {
	CodeMashRequestBase
	/** @description The import id (imp_…). */
	// @ApiMember(Description="The import id (imp_…).", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetCollectionImport) CreateResponse() (r GetCollectionImportResponse) { return }
func (GetCollectionImport) HttpMethod() string                              { return "GET" }

/** @description Database */
// @Route("/{version}/database/imports", "GET")
// @Api(Description="Database")
type GetCollectionImports struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetCollectionImports) CreateResponse() (r GetCollectionImportsResponse) { return }
func (GetCollectionImports) HttpMethod() string                               { return "GET" }

/** @description Database */
// @Route("/{version}/database/imports/upload-url", "POST")
// @Api(Description="Database")
type RequestImportUploadUrlRequest struct {
	CodeMashRequestBase
	FileAccountId *string `json:"fileAccountId,omitempty"`
	/** @description The original CSV file name, e.g. people.csv. */
	// @ApiMember(Description="The original CSV file name, e.g. people.csv.", IsRequired=true)
	FileName string `json:"fileName"`
}

func (RequestImportUploadUrlRequest) CreateResponse() (r RequestImportUploadUrlResponse) { return }
func (RequestImportUploadUrlRequest) HttpMethod() string                                 { return "POST" }

/** @description Database */
// @Route("/{version}/database/imports/analyze", "POST")
// @Api(Description="Database")
type AnalyzeImportFileRequest struct {
	CodeMashRequestBase
	/** @description The uploaded CSV's file ref, from the upload-url call. */
	// @ApiMember(Description="The uploaded CSV's file ref, from the upload-url call.", IsRequired=true)
	File FileResourceRefDto `json:"file"`
	/** @description The CSV delimiter, e.g. "," or ";". */
	// @ApiMember(Description="The CSV delimiter, e.g. \",\" or \";\".", IsRequired=true)
	Delimiter string `json:"delimiter"`
	/** @description Whether the first row is a header row. */
	// @ApiMember(Description="Whether the first row is a header row.")
	HasHeader bool `json:"hasHeader,omitempty"`
}

func (AnalyzeImportFileRequest) CreateResponse() (r AnalyzeImportFileResponse) { return }
func (AnalyzeImportFileRequest) HttpMethod() string                            { return "POST" }

/** @description Create email campaign */
// @Route("/{version}/notifications/email/campaigns", "POST")
// @Api(Description="Create email campaign")
// @DataContract
type CreateEmailCampaignRequest struct {
	CodeMashRequestBase
	// @DataMember
	Campaign EmailCampaignRequest `json:"campaign"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional. true sends the campaign now (or at a campaignTime sooner than the review window). Without it, a campaign created by an AI tool must be scheduled at least 24 hours ahead for role-based / all-users / collection audiences and at least 1 hour ahead for direct recipient lists, or it is rejected with CM-ERRORS-CAMPAIGNS-017. Set it only when the user explicitly asked to send now. */
	// @DataMember
	// @ApiMember(Description="Optional. true sends the campaign now (or at a campaignTime sooner than the review window). Without it, a campaign created by an AI tool must be scheduled at least 24 hours ahead for role-based / all-users / collection audiences and at least 1 hour ahead for direct recipient lists, or it is rejected with CM-ERRORS-CAMPAIGNS-017. Set it only when the user explicitly asked to send now.")
	SendNow *bool `json:"sendNow,omitempty"`
}

func (CreateEmailCampaignRequest) CreateResponse() (r IdResponse) { return }
func (CreateEmailCampaignRequest) HttpMethod() string             { return "POST" }

/** @description Deletes emails campaign from queue */
// @Route("/{version}/notifications/email/campaigns/{Id}", "DELETE")
// @Api(Description="Deletes emails campaign from queue")
// @DataContract
type DeleteEmailCampaignRequest struct {
	CodeMashRequestBase
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (DeleteEmailCampaignRequest) CreateResponse() (r EmptyResponse) { return }
func (DeleteEmailCampaignRequest) HttpMethod() string                { return "DELETE" }

/** @description Gets email campaign by id */
// @Route("/{version}/notifications/email/campaigns/{id}", "GET")
// @Api(Description="Gets email campaign by id")
type GetEmailCampaign struct {
	CodeMashRequestBase
	/** @description The campaign id. */
	// @ApiMember(Description="The campaign id.")
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetEmailCampaign) CreateResponse() (r GetEmailCampaignResponse) { return }
func (GetEmailCampaign) HttpMethod() string                           { return "GET" }

/** @description Gets email campaigns */
// @Route("/{version}/notifications/email/campaigns", "GET")
// @Api(Description="Gets email campaigns")
type GetEmailCampaigns struct {
	CodeMashListPaginationRequestBase
	/** @description Optional. When omitted, the project's default database integration is used (resolved server-side from the project state). */
	// @ApiMember(Description="Optional. When omitted, the project's default database integration is used (resolved server-side from the project state).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional: return only the campaign with this id. */
	// @ApiMember(Description="Optional: return only the campaign with this id.")
	CampaignId *string `json:"campaignId,omitempty"`
	/** @description Optional: only campaigns that targeted this email address. */
	// @ApiMember(Description="Optional: only campaigns that targeted this email address.")
	EmailAddress *string `json:"emailAddress,omitempty"`
	/** @description Optional: only campaigns built on this email template id. */
	// @ApiMember(Description="Optional: only campaigns built on this email template id.")
	TemplateId *string `json:"templateId,omitempty"`
	/** @description Optional lower bound for the campaign time, unix timestamp in seconds (UTC). */
	// @ApiMember(Description="Optional lower bound for the campaign time, unix timestamp in seconds (UTC).")
	From *int64 `json:"from,omitempty"`
	/** @description Optional upper bound for the campaign time, unix timestamp in seconds (UTC). */
	// @ApiMember(Description="Optional upper bound for the campaign time, unix timestamp in seconds (UTC).")
	To *int64 `json:"to,omitempty"`
}

func (GetEmailCampaigns) CreateResponse() (r GetEmailCampaignsResponse) { return }
func (GetEmailCampaigns) HttpMethod() string                            { return "GET" }

/** @description Get email campaign batches */
// @Route("/{version}/notifications/email/campaigns/{id}/batches", "GET")
// @Api(Description="Get email campaign batches")
type GetEmailCampaignBatches struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id to list batches for. Get it from get_email_campaigns. */
	// @ApiMember(Description="The email campaign id to list batches for. Get it from get_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional batch id to filter to a single batch. Get it from a prior call to this tool. */
	// @ApiMember(Description="Optional batch id to filter to a single batch. Get it from a prior call to this tool.")
	BatchId *string `json:"batchId,omitempty"`
	/** @description Optional recipient email address to filter batches by. */
	// @ApiMember(Description="Optional recipient email address to filter batches by.")
	EmailAddress *string `json:"emailAddress,omitempty"`
}

func (GetEmailCampaignBatches) CreateResponse() (r GetEmailCampaignBatchesResponse) { return }
func (GetEmailCampaignBatches) HttpMethod() string                                  { return "GET" }

/** @description Get an email campaign batch notification */
// @Route("/{version}/notifications/email/campaigns/{id}/batches/{batchId}/{notificationId}", "GET")
// @Api(Description="Get an email campaign batch notification")
type GetEmailCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id. Get it from get_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description The campaign batch id. Get it from get_email_campaign_batches. */
	// @ApiMember(Description="The campaign batch id. Get it from get_email_campaign_batches.", IsRequired=true)
	BatchId string `json:"batchId"`
	/** @description The notification id within the batch. Get it from get_email_campaign_batch_notifications. */
	// @ApiMember(Description="The notification id within the batch. Get it from get_email_campaign_batch_notifications.", IsRequired=true)
	NotificationId string `json:"notificationId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetEmailCampaignBatchNotification) CreateResponse() (r GetEmailCampaignBatchNotificationResponse) {
	return
}
func (GetEmailCampaignBatchNotification) HttpMethod() string { return "GET" }

/** @description Get email campaign batch notifications */
// @Route("/{version}/notifications/email/campaigns/{id}/batches/{batchId}", "GET")
// @Api(Description="Get email campaign batch notifications")
type GetEmailCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id. Get it from get_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description The campaign batch id to list notifications for. Get it from get_email_campaign_batches. */
	// @ApiMember(Description="The campaign batch id to list notifications for. Get it from get_email_campaign_batches.", IsRequired=true)
	BatchId string `json:"batchId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetEmailCampaignBatchNotifications) CreateResponse() (r GetEmailCampaignBatchNotificationsResponse) {
	return
}
func (GetEmailCampaignBatchNotifications) HttpMethod() string { return "GET" }

/** @description Get email campaign statistics */
// @Route("/{version}/notifications/email/campaigns/{id}/stats", "GET")
// @Api(Description="Get email campaign statistics")
type GetEmailCampaignStatistics struct {
	CodeMashRequestBase
	/** @description The email campaign id to get statistics for. Get it from get_email_campaigns. */
	// @ApiMember(Description="The email campaign id to get statistics for. Get it from get_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetEmailCampaignStatistics) CreateResponse() (r GetEmailCampaignStatisticsResponse) { return }
func (GetEmailCampaignStatistics) HttpMethod() string                                     { return "GET" }

/** @description Preview an email notification */
// @Route("/{version}/notifications/email/preview", "GET")
// @Api(Description="Preview an email notification")
type PreviewEmailNotification struct {
	RequestBase
	/** @description The signed preview link hash. With it the preview opens without sign-in: exactly one notification, until the link expires (30 days after sending). */
	// @ApiMember(Description="The signed preview link hash. With it the preview opens without sign-in: exactly one notification, until the link expires (30 days after sending).")
	Hash string `json:"hash"`
	/** @description Signed-in preview by id: the project id. With a hash it must match the hash. */
	// @ApiMember(Description="Signed-in preview by id: the project id. With a hash it must match the hash.")
	ProjectId *string `json:"projectId,omitempty"`
	/** @description Signed-in preview by id: the notification id (needs email:read). With a hash it must match the hash. */
	// @ApiMember(Description="Signed-in preview by id: the notification id (needs email:read). With a hash it must match the hash.")
	NotificationId *string `json:"notificationId,omitempty"`
}

func (PreviewEmailNotification) CreateResponse() (r PreviewEmailNotificationResponse) { return }
func (PreviewEmailNotification) HttpMethod() string                                   { return "GET" }

/** @description Retries the failed recipients of an email campaign */
// @Route("/{version}/notifications/email/campaigns/{Id}/retry-failed", "POST")
// @Api(Description="Retries the failed recipients of an email campaign")
// @DataContract
type RetryFailedEmailNotificationsRequest struct {
	CodeMashRequestBase
	/** @description The campaign id. Get it from get_email_campaigns. */
	// @DataMember
	// @ApiMember(Description="The campaign id. Get it from get_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @DataMember
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Your key for this request, at most 64 characters. The same key again returns the first answer and sends nothing; the same key with another body is refused (409). */
	// @DataMember
	// @ApiMember(Description="Your key for this request, at most 64 characters. The same key again returns the first answer and sends nothing; the same key with another body is refused (409).", IsRequired=true)
	IdempotencyKey string `json:"idempotencyKey"`
	/** @description Optional. Only these notification ids. Default: every failed recipient of the campaign. */
	// @DataMember
	// @ApiMember(Description="Optional. Only these notification ids. Default: every failed recipient of the campaign.")
	NotificationIds []string `json:"notificationIds,omitempty"`
	/** @description Optional, default false. Also retry recipients whose failure is Ambiguous (a timeout after the request was sent) — this MAY SEND A DUPLICATE. */
	// @DataMember
	// @ApiMember(Description="Optional, default false. Also retry recipients whose failure is Ambiguous (a timeout after the request was sent) — this MAY SEND A DUPLICATE.")
	IncludeAmbiguous *bool `json:"includeAmbiguous,omitempty"`
	/** @description Optional. Only these failure classes: Transient, Permanent, Unknown. Default: all three (Ambiguous only through includeAmbiguous). */
	// @DataMember
	// @ApiMember(Description="Optional. Only these failure classes: Transient, Permanent, Unknown. Default: all three (Ambiguous only through includeAmbiguous).")
	Categories []string `json:"categories,omitempty"`
}

func (RetryFailedEmailNotificationsRequest) CreateResponse() (r RetryFailedEmailNotificationsResponse) {
	return
}
func (RetryFailedEmailNotificationsRequest) HttpMethod() string { return "POST" }

/** @description Stops a running email campaign */
// @Route("/{version}/notifications/email/campaigns/{Id}/stop", "POST")
// @Api(Description="Stops a running email campaign")
// @DataContract
type StopEmailCampaignRequest struct {
	CodeMashRequestBase
	/** @description The campaign id to stop. */
	// @DataMember
	// @ApiMember(Description="The campaign id to stop.")
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @DataMember
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (StopEmailCampaignRequest) CreateResponse() (r EmptyResponse) { return }
func (StopEmailCampaignRequest) HttpMethod() string                { return "POST" }

/** @description Get email campaign messages */
// @Route("/{version}/notifications/emails/campaigns/{campaignId}/messages", "GET")
// @Api(Description="Get email campaign messages")
type GetEmailCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id. Get it from get_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_email_campaigns.", IsRequired=true)
	CampaignId string `json:"campaignId"`
	/** @description The campaign batch id to list messages for. Get it from get_email_campaign_batches. */
	// @ApiMember(Description="The campaign batch id to list messages for. Get it from get_email_campaign_batches.", IsRequired=true)
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetEmailCampaignMessagesRequest) CreateResponse() (r GetEmailCampaignMessagesResponse) { return }
func (GetEmailCampaignMessagesRequest) HttpMethod() string                                   { return "GET" }

/** @description Create SMS campaign */
// @Route("/{version}/notifications/sms/campaigns", "POST")
// @Api(Description="Create SMS campaign")
// @DataContract
type CreateSmsCampaignRequest struct {
	CodeMashRequestBase
	/** @description SMS template id to send — pick one with get_sms_templates. Never invent it. */
	// @DataMember
	// @ApiMember(Description="SMS template id to send — pick one with get_sms_templates. Never invent it.")
	TemplateId string `json:"templateId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @DataMember
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description SMS provider integration id the campaign sends through — pick one with get_sms_integrations (the project default unless the user named another). Required; never invent it. */
	// @DataMember
	// @ApiMember(Description="SMS provider integration id the campaign sends through — pick one with get_sms_integrations (the project default unless the user named another). Required; never invent it.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
	/** @description Optional language code forcing one template translation for every recipient. */
	// @DataMember
	// @ApiMember(Description="Optional language code forcing one template translation for every recipient.")
	Language *string `json:"language,omitempty"`
	// @DataMember
	InitiatorId *string `json:"initiatorId,omitempty"`
	/** @description Audience type: 'AllUsers' (every project member subscribed to the SMS channel — role-based delivery can address MILLIONS of contacts), 'SpecifiedUsers' (exact member ids), or 'PhoneNumbers' (raw phone numbers). Fill EXACTLY the settings object matching this value. 'Collection' delivery is not available from chat. */
	// @DataMember
	// @ApiMember(Description="Audience type: 'AllUsers' (every project member subscribed to the SMS channel — role-based delivery can address MILLIONS of contacts), 'SpecifiedUsers' (exact member ids), or 'PhoneNumbers' (raw phone numbers). Fill EXACTLY the settings object matching this value. 'Collection' delivery is not available from chat.")
	DeliveryType SmsCampaignRecipientsSourceTypes `json:"deliveryType,omitempty"`
	/** @description For deliveryType 'AllUsers'. JSON object: {"recipientsSourceType":"AllUsers","rolesNames":["authenticated"],"userTags":[],"campaignTime":<unix seconds UTC>}. rolesNames/userTags are optional narrowing filters — verify exact role names with get_roles. */
	// @DataMember
	// @ApiMember(Description="For deliveryType 'AllUsers'. JSON object: {\"recipientsSourceType\":\"AllUsers\",\"rolesNames\":[\"authenticated\"],\"userTags\":[],\"campaignTime\":<unix seconds UTC>}. rolesNames/userTags are optional narrowing filters — verify exact role names with get_roles.")
	AllUsers *SmsToAllUsersDeliverySettingsDto `json:"allUsers,omitempty"`
	/** @description For deliveryType 'SpecifiedUsers'. JSON object: {"recipientsSourceType":"SpecifiedUsers","recipients":[<member ids>],"campaignTime":<unix seconds UTC>}. */
	// @DataMember
	// @ApiMember(Description="For deliveryType 'SpecifiedUsers'. JSON object: {\"recipientsSourceType\":\"SpecifiedUsers\",\"recipients\":[<member ids>],\"campaignTime\":<unix seconds UTC>}.")
	SpecifiedUsers *SmsToUsersDeliverySettingsDto `json:"specifiedUsers,omitempty"`
	/** @description For deliveryType 'AccountUsers'. JSON object: {"recipientsSourceType":"AccountUsers","recipients":[<account owner / team member ids>],"campaignTime":<unix seconds UTC>}. Members without a phone number are skipped. */
	// @DataMember
	// @ApiMember(Description="For deliveryType 'AccountUsers'. JSON object: {\"recipientsSourceType\":\"AccountUsers\",\"recipients\":[<account owner / team member ids>],\"campaignTime\":<unix seconds UTC>}. Members without a phone number are skipped.")
	AccountUsers *SmsToAccountUsersDeliverySettingsDto `json:"accountUsers,omitempty"`
	// @DataMember
	Collection *SmsToCollectionRecordsDeliverySettingsDto `json:"collection,omitempty"`
	/** @description For deliveryType 'PhoneNumbers'. JSON object: {"recipientsSourceType":"PhoneNumbers","phoneNumbers":["+37060000000"],"campaignTime":<unix seconds UTC>}. Numbers in international format. */
	// @DataMember
	// @ApiMember(Description="For deliveryType 'PhoneNumbers'. JSON object: {\"recipientsSourceType\":\"PhoneNumbers\",\"phoneNumbers\":[\"+37060000000\"],\"campaignTime\":<unix seconds UTC>}. Numbers in international format.")
	PhoneNumbers *SmsToPhoneNumbersDeliverySettingsDto `json:"phoneNumbers,omitempty"`
	/** @description Optional. true sends the campaign now (or at a campaignTime sooner than the review window). Without it, a campaign created by an AI tool must be scheduled at least 24 hours ahead for role-based / all-users / collection audiences and at least 1 hour ahead for direct recipient lists, or it is rejected with CM-ERRORS-CAMPAIGNS-017. Set it only when the user explicitly asked to send now. */
	// @DataMember
	// @ApiMember(Description="Optional. true sends the campaign now (or at a campaignTime sooner than the review window). Without it, a campaign created by an AI tool must be scheduled at least 24 hours ahead for role-based / all-users / collection audiences and at least 1 hour ahead for direct recipient lists, or it is rejected with CM-ERRORS-CAMPAIGNS-017. Set it only when the user explicitly asked to send now.")
	SendNow *bool `json:"sendNow,omitempty"`
}

func (CreateSmsCampaignRequest) CreateResponse() (r IdResponse) { return }
func (CreateSmsCampaignRequest) HttpMethod() string             { return "POST" }

/** @description Deletes sms campaign from queue */
// @Route("/{version}/notifications/sms/campaigns/{id}", "DELETE")
// @Api(Description="Deletes sms campaign from queue")
type DeleteSmsCampaign struct {
	CodeMashRequestBase
	/** @description The campaign id to delete. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id to delete. Get it from get_sms_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (DeleteSmsCampaign) CreateResponse() (r EmptyResponse) { return }
func (DeleteSmsCampaign) HttpMethod() string                { return "DELETE" }

/** @description Get sms campaign by id */
// @Route("/{version}/notifications/sms/campaigns/{id}", "GET")
// @Api(Description="Get sms campaign by id")
type GetSmsCampaign struct {
	CodeMashRequestBase
	/** @description The campaign id. */
	// @ApiMember(Description="The campaign id.")
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetSmsCampaign) CreateResponse() (r GetSmsCampaignResponse) { return }
func (GetSmsCampaign) HttpMethod() string                         { return "GET" }

/** @description Gets sms campaigns */
// @Route("/{version}/notifications/sms/campaigns", "GET")
// @Api(Description="Gets sms campaigns")
type GetSmsCampaigns struct {
	CodeMashListPaginationRequestBase
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional: return only the campaign with this id. */
	// @ApiMember(Description="Optional: return only the campaign with this id.")
	CampaignId *string `json:"campaignId,omitempty"`
	/** @description Optional: only campaigns built on this SMS template id. */
	// @ApiMember(Description="Optional: only campaigns built on this SMS template id.")
	TemplateId *string `json:"templateId,omitempty"`
	/** @description Optional lower bound for the campaign time, unix timestamp in seconds (UTC). */
	// @ApiMember(Description="Optional lower bound for the campaign time, unix timestamp in seconds (UTC).")
	From *int64 `json:"from,omitempty"`
	/** @description Optional upper bound for the campaign time, unix timestamp in seconds (UTC). */
	// @ApiMember(Description="Optional upper bound for the campaign time, unix timestamp in seconds (UTC).")
	To *int64 `json:"to,omitempty"`
}

func (GetSmsCampaigns) CreateResponse() (r GetSmsCampaignsResponse) { return }
func (GetSmsCampaigns) HttpMethod() string                          { return "GET" }

/** @description Gets sms campaign batches */
// @Route("/{version}/notifications/sms/campaigns/{id}/batches", "GET")
// @Api(Description="Gets sms campaign batches")
type GetSmsCampaignBatches struct {
	CodeMashListPaginationRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.")
	Id *string `json:"id,omitempty"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetSmsCampaignBatches) CreateResponse() (r GetSmsCampaignBatchesResponse) { return }
func (GetSmsCampaignBatches) HttpMethod() string                                { return "GET" }

/** @description Gets sms campaign batch notification */
// @Route("/{version}/notifications/sms/campaigns/{id}/batches/{batchId}/{notificationId}", "GET")
// @Api(Description="Gets sms campaign batch notification")
type GetSmsCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description The campaign batch id. Get it from get_sms_campaign_batches. */
	// @ApiMember(Description="The campaign batch id. Get it from get_sms_campaign_batches.", IsRequired=true)
	BatchId string `json:"batchId"`
	/** @description The notification id. Get it from get_sms_campaign_batch_notifications. */
	// @ApiMember(Description="The notification id. Get it from get_sms_campaign_batch_notifications.", IsRequired=true)
	NotificationId string `json:"notificationId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetSmsCampaignBatchNotification) CreateResponse() (r GetSmsCampaignBatchNotificationResponse) {
	return
}
func (GetSmsCampaignBatchNotification) HttpMethod() string { return "GET" }

/** @description Gets sms campaign batch notifications */
// @Route("/{version}/notifications/sms/campaigns/{id}/batches/{batchId}", "GET")
// @Api(Description="Gets sms campaign batch notifications")
type GetSmsCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description The campaign batch id. Get it from get_sms_campaign_batches. */
	// @ApiMember(Description="The campaign batch id. Get it from get_sms_campaign_batches.", IsRequired=true)
	BatchId string `json:"batchId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetSmsCampaignBatchNotifications) CreateResponse() (r GetSmsCampaignBatchNotificationsResponse) {
	return
}
func (GetSmsCampaignBatchNotifications) HttpMethod() string { return "GET" }

/** @description Get sms campaign statistics */
// @Route("/{version}/notifications/sms/campaigns/{id}/stats", "GET")
// @Api(Description="Get sms campaign statistics")
type GetSmsCampaignStatistics struct {
	CodeMashRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetSmsCampaignStatistics) CreateResponse() (r GetSmsCampaignStatisticsResponse) { return }
func (GetSmsCampaignStatistics) HttpMethod() string                                   { return "GET" }

/** @description Returns SMS preview notification body */
// @Route("/{version}/notifications/sms/preview", "GET")
// @Api(Description="Returns SMS preview notification body")
type PreviewSmsNotification struct {
	RequestBase
	/** @description The signed preview link hash. With it the preview opens without sign-in: exactly one notification, until the link expires (30 days after sending). */
	// @ApiMember(Description="The signed preview link hash. With it the preview opens without sign-in: exactly one notification, until the link expires (30 days after sending).")
	Hash string `json:"hash"`
	/** @description Signed-in preview by id: the project id. With a hash it must match the hash. */
	// @ApiMember(Description="Signed-in preview by id: the project id. With a hash it must match the hash.")
	ProjectId *string `json:"projectId,omitempty"`
	/** @description Signed-in preview by id: the notification id (needs sms:read). With a hash it must match the hash. */
	// @ApiMember(Description="Signed-in preview by id: the notification id (needs sms:read). With a hash it must match the hash.")
	NotificationId *string `json:"notificationId,omitempty"`
}

func (PreviewSmsNotification) CreateResponse() (r PreviewSmsNotificationResponse) { return }
func (PreviewSmsNotification) HttpMethod() string                                 { return "GET" }

/** @description Retries the failed recipients of an SMS campaign */
// @Route("/{version}/notifications/sms/campaigns/{Id}/retry-failed", "POST")
// @Api(Description="Retries the failed recipients of an SMS campaign")
// @DataContract
type RetryFailedSmsNotificationsRequest struct {
	CodeMashRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @DataMember
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @DataMember
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Your key for this request, at most 64 characters. The same key again returns the first answer and sends nothing; the same key with another body is refused (409). */
	// @DataMember
	// @ApiMember(Description="Your key for this request, at most 64 characters. The same key again returns the first answer and sends nothing; the same key with another body is refused (409).", IsRequired=true)
	IdempotencyKey string `json:"idempotencyKey"`
	/** @description Optional. Only these notification ids. Default: every failed recipient of the campaign. */
	// @DataMember
	// @ApiMember(Description="Optional. Only these notification ids. Default: every failed recipient of the campaign.")
	NotificationIds []string `json:"notificationIds,omitempty"`
	/** @description Optional, default false. Also retry recipients whose failure is Ambiguous (a timeout after the request was sent) — this MAY SEND A DUPLICATE. */
	// @DataMember
	// @ApiMember(Description="Optional, default false. Also retry recipients whose failure is Ambiguous (a timeout after the request was sent) — this MAY SEND A DUPLICATE.")
	IncludeAmbiguous *bool `json:"includeAmbiguous,omitempty"`
	/** @description Optional. Only these failure classes: Transient, Permanent, Unknown. Default: all three (Ambiguous only through includeAmbiguous). */
	// @DataMember
	// @ApiMember(Description="Optional. Only these failure classes: Transient, Permanent, Unknown. Default: all three (Ambiguous only through includeAmbiguous).")
	Categories []string `json:"categories,omitempty"`
}

func (RetryFailedSmsNotificationsRequest) CreateResponse() (r RetryFailedSmsNotificationsResponse) {
	return
}
func (RetryFailedSmsNotificationsRequest) HttpMethod() string { return "POST" }

/** @description Stops a running SMS campaign */
// @Route("/{version}/notifications/sms/campaigns/{Id}/stop", "POST")
// @Api(Description="Stops a running SMS campaign")
type StopSmsCampaignRequest struct {
	CodeMashRequestBase
	/** @description The campaign id to stop. */
	// @ApiMember(Description="The campaign id to stop.")
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (StopSmsCampaignRequest) CreateResponse() (r EmptyResponse) { return }
func (StopSmsCampaignRequest) HttpMethod() string                { return "POST" }

/** @description Gets the sms notifications */
// @Route("/{version}/notifications/sms/campaigns/{campaignId}/messages", "GET")
// @Api(Description="Gets the sms notifications")
type GetSmsCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.", IsRequired=true)
	CampaignId string `json:"campaignId"`
	/** @description The campaign batch id. Get it from get_sms_campaign_batches. */
	// @ApiMember(Description="The campaign batch id. Get it from get_sms_campaign_batches.", IsRequired=true)
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetSmsCampaignMessagesRequest) CreateResponse() (r GetSmsCampaignMessagesResponse) { return }
func (GetSmsCampaignMessagesRequest) HttpMethod() string                                 { return "GET" }

/** @description Count the devices a push campaign audience would reach */
// @Route("/{version}/notifications/push/campaigns/audience-count", "POST")
// @Api(Description="Count the devices a push campaign audience would reach")
// @DataContract
type GetPushCampaignAudienceCountRequest struct {
	CodeMashRequestBase
	/** @description The audience, in the same shape as CreatePushCampaignRequest.campaign. Template and send options are ignored. */
	// @DataMember
	// @ApiMember(Description="The audience, in the same shape as CreatePushCampaignRequest.campaign. Template and send options are ignored.")
	Campaign PushCampaignRequest `json:"campaign"`
	/** @description Optional. When omitted, the default database integration for the request's environment is used. */
	// @DataMember
	// @ApiMember(Description="Optional. When omitted, the default database integration for the request's environment is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetPushCampaignAudienceCountRequest) CreateResponse() (r GetPushCampaignAudienceCountResponse) {
	return
}
func (GetPushCampaignAudienceCountRequest) HttpMethod() string { return "POST" }

/** @description Create push campaign */
// @Route("/{version}/notifications/push/campaigns", "POST")
// @Api(Description="Create push campaign")
// @DataContract
type CreatePushCampaignRequest struct {
	CodeMashRequestBase
	// @DataMember
	Campaign PushCampaignRequest `json:"campaign"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional. true sends the campaign now (or at a campaignTime sooner than the review window). Without it, a campaign created by an AI tool must be scheduled at least 24 hours ahead for role-based / all-users / collection audiences and at least 1 hour ahead for direct recipient lists, or it is rejected with CM-ERRORS-CAMPAIGNS-017. Set it only when the user explicitly asked to send now. */
	// @DataMember
	// @ApiMember(Description="Optional. true sends the campaign now (or at a campaignTime sooner than the review window). Without it, a campaign created by an AI tool must be scheduled at least 24 hours ahead for role-based / all-users / collection audiences and at least 1 hour ahead for direct recipient lists, or it is rejected with CM-ERRORS-CAMPAIGNS-017. Set it only when the user explicitly asked to send now.")
	SendNow *bool `json:"sendNow,omitempty"`
}

func (CreatePushCampaignRequest) CreateResponse() (r IdResponse) { return }
func (CreatePushCampaignRequest) HttpMethod() string             { return "POST" }

/** @description Deletes push campaign from queue */
// @Route("/{version}/notifications/push/campaigns/{Id}", "DELETE")
// @Api(Description="Deletes push campaign from queue")
// @DataContract
type DeletePushCampaignRequest struct {
	CodeMashRequestBase
	/** @description The push campaign id to delete. Get it from get_push_campaigns. */
	// @DataMember
	// @ApiMember(Description="The push campaign id to delete. Get it from get_push_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (DeletePushCampaignRequest) CreateResponse() (r EmptyResponse) { return }
func (DeletePushCampaignRequest) HttpMethod() string                { return "DELETE" }

/** @description Gets push campaign by id */
// @Route("/{version}/notifications/push/campaigns/{id}", "GET")
// @Api(Description="Gets push campaign by id")
type GetPushCampaign struct {
	CodeMashRequestBase
	/** @description The campaign id. */
	// @ApiMember(Description="The campaign id.")
	Id string `json:"id"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetPushCampaign) CreateResponse() (r GetPushCampaignResponse) { return }
func (GetPushCampaign) HttpMethod() string                          { return "GET" }

/** @description Gets push campaigns */
// @Route("/{version}/notifications/push/campaigns", "GET")
// @Api(Description="Gets push campaigns")
type GetPushCampaigns struct {
	CodeMashListPaginationRequestBase
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional: only campaigns built on this push template id. */
	// @ApiMember(Description="Optional: only campaigns built on this push template id.")
	TemplateId *string `json:"templateId,omitempty"`
	/** @description Optional lower bound for the campaign time, unix timestamp in seconds (UTC). */
	// @ApiMember(Description="Optional lower bound for the campaign time, unix timestamp in seconds (UTC).")
	From *int64 `json:"from,omitempty"`
	/** @description Optional upper bound for the campaign time, unix timestamp in seconds (UTC). */
	// @ApiMember(Description="Optional upper bound for the campaign time, unix timestamp in seconds (UTC).")
	To *int64 `json:"to,omitempty"`
}

func (GetPushCampaigns) CreateResponse() (r GetPushCampaignsResponse) { return }
func (GetPushCampaigns) HttpMethod() string                           { return "GET" }

/** @description Gets push campaign batches */
// @Route("/{version}/notifications/push/campaigns/{id}/batches", "GET")
// @Api(Description="Gets push campaign batches")
type GetPushCampaignBatches struct {
	CodeMashListPaginationRequestBase
	/** @description The push campaign id to list batches for. Get it from get_push_campaigns. */
	// @ApiMember(Description="The push campaign id to list batches for. Get it from get_push_campaigns.")
	Id string `json:"id"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional: only return the batch with this id. */
	// @ApiMember(Description="Optional: only return the batch with this id.")
	BatchId *string `json:"batchId,omitempty"`
}

func (GetPushCampaignBatches) CreateResponse() (r GetPushCampaignBatchesResponse) { return }
func (GetPushCampaignBatches) HttpMethod() string                                 { return "GET" }

/** @description Gets a push campaign batch notification */
// @Route("/{version}/notifications/push/campaigns/{id}/batches/{batchId}/{notificationId}", "GET")
// @Api(Description="Gets a push campaign batch notification")
type GetPushCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	/** @description The push campaign id. Get it from get_push_campaigns. */
	// @ApiMember(Description="The push campaign id. Get it from get_push_campaigns.")
	Id string `json:"id"`
	/** @description The batch id. Get it from get_push_campaign_batches. */
	// @ApiMember(Description="The batch id. Get it from get_push_campaign_batches.")
	BatchId string `json:"batchId"`
	/** @description The notification id within the batch. */
	// @ApiMember(Description="The notification id within the batch.")
	NotificationId string `json:"notificationId"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetPushCampaignBatchNotification) CreateResponse() (r GetPushCampaignBatchNotificationResponse) {
	return
}
func (GetPushCampaignBatchNotification) HttpMethod() string { return "GET" }

/** @description Gets push campaign batch notifications */
// @Route("/{version}/notifications/push/campaigns/{id}/batches/{batchId}", "GET")
// @Api(Description="Gets push campaign batch notifications")
type GetPushCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	/** @description The push campaign id. Get it from get_push_campaigns. */
	// @ApiMember(Description="The push campaign id. Get it from get_push_campaigns.")
	Id string `json:"id"`
	/** @description The batch id to list notifications for. Get it from get_push_campaign_batches. */
	// @ApiMember(Description="The batch id to list notifications for. Get it from get_push_campaign_batches.")
	BatchId string `json:"batchId"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetPushCampaignBatchNotifications) CreateResponse() (r GetPushCampaignBatchNotificationsResponse) {
	return
}
func (GetPushCampaignBatchNotifications) HttpMethod() string { return "GET" }

/** @description Get push campaign statistics */
// @Route("/{version}/notifications/push/campaigns/{id}/stats", "GET")
// @Api(Description="Get push campaign statistics")
type GetPushCampaignStatistics struct {
	CodeMashRequestBase
	/** @description The push campaign id to get statistics for. Get it from get_push_campaigns. */
	// @ApiMember(Description="The push campaign id to get statistics for. Get it from get_push_campaigns.")
	Id string `json:"id"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetPushCampaignStatistics) CreateResponse() (r GetPushCampaignStatisticsResponse) { return }
func (GetPushCampaignStatistics) HttpMethod() string                                    { return "GET" }

/** @description Returns push preview notification */
// @Route("/{version}/notifications/push/preview", "GET")
// @Api(Description="Returns push preview notification")
type PreviewPushNotification struct {
	RequestBase
	/** @description The signed preview link hash. With it the preview opens without sign-in: exactly one notification, until the link expires (30 days after sending). */
	// @ApiMember(Description="The signed preview link hash. With it the preview opens without sign-in: exactly one notification, until the link expires (30 days after sending).")
	Hash string `json:"hash"`
	/** @description Signed-in preview by id: the project id. With a hash it must match the hash. */
	// @ApiMember(Description="Signed-in preview by id: the project id. With a hash it must match the hash.")
	ProjectId *string `json:"projectId,omitempty"`
	/** @description Signed-in preview by id: the notification id (needs push:read). With a hash it must match the hash. */
	// @ApiMember(Description="Signed-in preview by id: the notification id (needs push:read). With a hash it must match the hash.")
	NotificationId *string `json:"notificationId,omitempty"`
}

func (PreviewPushNotification) CreateResponse() (r PreviewPushNotificationResponse) { return }
func (PreviewPushNotification) HttpMethod() string                                  { return "GET" }

/** @description Retries the failed recipients of a push campaign */
// @Route("/{version}/notifications/push/campaigns/{Id}/retry-failed", "POST")
// @Api(Description="Retries the failed recipients of a push campaign")
// @DataContract
type RetryFailedPushNotificationsRequest struct {
	CodeMashRequestBase
	/** @description The campaign id. Get it from get_push_campaigns. */
	// @DataMember
	// @ApiMember(Description="The campaign id. Get it from get_push_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @DataMember
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Your key for this request, at most 64 characters. The same key again returns the first answer and sends nothing; the same key with another body is refused (409). */
	// @DataMember
	// @ApiMember(Description="Your key for this request, at most 64 characters. The same key again returns the first answer and sends nothing; the same key with another body is refused (409).", IsRequired=true)
	IdempotencyKey string `json:"idempotencyKey"`
	/** @description Optional. Only these notification ids. Default: every failed recipient of the campaign. */
	// @DataMember
	// @ApiMember(Description="Optional. Only these notification ids. Default: every failed recipient of the campaign.")
	NotificationIds []string `json:"notificationIds,omitempty"`
	/** @description Optional, default false. Also retry recipients whose failure is Ambiguous (a timeout after the request was sent) — this MAY SEND A DUPLICATE. */
	// @DataMember
	// @ApiMember(Description="Optional, default false. Also retry recipients whose failure is Ambiguous (a timeout after the request was sent) — this MAY SEND A DUPLICATE.")
	IncludeAmbiguous *bool `json:"includeAmbiguous,omitempty"`
	/** @description Optional. Only these failure classes: Transient, Permanent, Unknown. Default: all three (Ambiguous only through includeAmbiguous). */
	// @DataMember
	// @ApiMember(Description="Optional. Only these failure classes: Transient, Permanent, Unknown. Default: all three (Ambiguous only through includeAmbiguous).")
	Categories []string `json:"categories,omitempty"`
}

func (RetryFailedPushNotificationsRequest) CreateResponse() (r RetryFailedPushNotificationsResponse) {
	return
}
func (RetryFailedPushNotificationsRequest) HttpMethod() string { return "POST" }

/** @description Stops a running push campaign */
// @Route("/{version}/notifications/push/campaigns/{Id}/stop", "POST")
// @Api(Description="Stops a running push campaign")
// @DataContract
type StopPushCampaignRequest struct {
	CodeMashRequestBase
	/** @description The campaign id to stop. */
	// @DataMember
	// @ApiMember(Description="The campaign id to stop.")
	Id string `json:"id"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (StopPushCampaignRequest) CreateResponse() (r EmptyResponse) { return }
func (StopPushCampaignRequest) HttpMethod() string                { return "POST" }

/** @description Gets push campaign messages */
// @Route("/{version}/notifications/push/campaigns/{campaignId}/messages", "GET")
// @Api(Description="Gets push campaign messages")
type GetPushCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	/** @description The push campaign id. Get it from get_push_campaigns. */
	// @ApiMember(Description="The push campaign id. Get it from get_push_campaigns.")
	CampaignId string `json:"campaignId"`
	/** @description Optional: restrict results to this batch id. Get it from get_push_campaign_batches. */
	// @ApiMember(Description="Optional: restrict results to this batch id. Get it from get_push_campaign_batches.")
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (GetPushCampaignMessagesRequest) CreateResponse() (r GetPushCampaignMessagesResponse) { return }
func (GetPushCampaignMessagesRequest) HttpMethod() string                                  { return "GET" }
