//go:build ignore

/* Options:
Date: 2026-10-07 10:13:50
Version: 10.20
Tip: To override a DTO option, remove "//" prefix before updating
BaseUrl: http://localhost:49877

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

// @DataContract(Namespace="http://codemash.io/types/")
type RequestBase struct {
	/** @description Specify culture code when your response from the API should be localised. E.g.: en */
	// @DataMember
	// @ApiMember(DataType="string", Description="Specify culture code when your response from the API should be localised. E.g.: en", Name="CultureCode", ParameterType="header")
	CultureCode *string `json:"cultureCode,omitempty"`
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

// @DataContract
type EchoLicenseDto struct {
	// @DataMember(Name="domain")
	Domain *string `json:"domain,omitempty"`
	// @DataMember(Name="accountId")
	AccountId *string `json:"accountId,omitempty"`
	// @DataMember(Name="email")
	Email *string `json:"email,omitempty"`
	// @DataMember(Name="release")
	Release *string `json:"release,omitempty"`
	// @DataMember(Name="expire")
	Expire int64 `json:"expire,omitempty"`
	// @DataMember(Name="isTrial")
	IsTrial bool `json:"isTrial,omitempty"`
	// @DataMember(Name="cap")
	Cap int `json:"cap,omitempty"`
}

type EchoRegionDto struct {
	Code        string `json:"code"`
	DisplayName string `json:"displayName"`
	ApiUrl      string `json:"apiUrl"`
	HubUrl      string `json:"hubUrl"`
}

type EchoAgentDto struct {
	McpUrl            string  `json:"mcpUrl"`
	OAuthMetadataUrl  *string `json:"oAuthMetadataUrl,omitempty"`
	InstallationType  string  `json:"installationType"`
	OnboardingDocsUrl string  `json:"onboardingDocsUrl"`
	ToolsUrl          string  `json:"toolsUrl"`
}

type PublicBrandDto struct {
	DisplayName string  `json:"displayName"`
	MainColor   *string `json:"mainColor,omitempty"`
	AccentColor *string `json:"accentColor,omitempty"`
	LogoUrl     *string `json:"logoUrl,omitempty"`
	IconUrl     *string `json:"iconUrl,omitempty"`
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
	AuthStatusActive                       = "Active"
	AuthStatusUnregistered                 = "Unregistered"
	AuthStatusSuspended                    = "Suspended"
	AuthStatusInActive                     = "InActive"
	AuthStatusBlocked                      = "Blocked"
)

type AuthDto struct {
	Id           string              `json:"id"`
	Type         AuthType            `json:"type,omitempty"`
	Email        *string             `json:"email,omitempty"`
	UserName     *string             `json:"userName,omitempty"`
	Registration *RegistrationDto    `json:"registration,omitempty"`
	Login        *LoginDto           `json:"login,omitempty"`
	GeneralInfo  *UserGeneralInfoDto `json:"generalInfo,omitempty"`
	Roles        []string            `json:"roles,omitempty"`
	PushDevices  []string            `json:"pushDevices,omitempty"`
	Tags         []string            `json:"tags,omitempty"`
	Status       AuthStatus          `json:"status,omitempty"`
	CreatedOn    time.Time           `json:"createdOn,omitempty"`
	ModifiedOn   time.Time           `json:"modifiedOn,omitempty"`
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
}

type JsonSchemaFieldDto struct {
	// @DataMember
	FieldName string `json:"fieldName"`
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

type SchemaSettingsDto struct {
	// @DataMember
	SoftDelete bool `json:"softDelete,omitempty"`
	// @DataMember
	HasRecordOwner bool `json:"hasRecordOwner,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
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
)

// @DataContract
type TriggerActionDto struct {
	// @DataMember
	Type TriggerActionType `json:"type,omitempty"`
	// @DataMember
	IntegrationId *string `json:"integrationId,omitempty"`
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
type FileChecksumDto struct {
	// @DataMember(Order=1)
	Algorithm string `json:"algorithm"`
	// @DataMember(Order=2)
	Hash string `json:"hash"`
}

// @DataContract
type FileResourceDto struct {
	// @DataMember(Order=1)
	Id string `json:"id"`
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

// @DataContract
type IntegrationTestResultItemDto struct {
	// @DataMember
	Operation string `json:"operation"`
	// @DataMember
	Result string `json:"result"`
	// @DataMember
	Errors *IReadOnlyList[string] `json:"errors,omitempty"`
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

type EchoResponse struct {
	ContainerName                *string         `json:"containerName,omitempty"`
	Ip                           string          `json:"ip"`
	Release                      CodeMashRelease `json:"release,omitempty"`
	Runtime                      CodeMashRuntime `json:"runtime,omitempty"`
	ManagedServiceHubUrl         string          `json:"managedServiceHubUrl"`
	ManagedServiceApiUrl         string          `json:"managedServiceApiUrl"`
	HubUrl                       string          `json:"hubUrl"`
	ApiUrl                       string          `json:"apiUrl"`
	ApiVersion                   string          `json:"apiVersion"`
	HubVersion                   string          `json:"hubVersion"`
	MjmlUrl                      string          `json:"mjmlUrl"`
	AdminUrlTemplate             *string         `json:"adminUrlTemplate,omitempty"`
	License                      *EchoLicenseDto `json:"license,omitempty"`
	AskForEnterpriseLicenseEmail *string         `json:"askForEnterpriseLicenseEmail,omitempty"`
	EmailServiceConfigured       bool            `json:"emailServiceConfigured,omitempty"`
	RootBootstrapPasswordSource  *string         `json:"rootBootstrapPasswordSource,omitempty"`
	Regions                      []EchoRegionDto `json:"regions,omitempty"`
	IsProductionInstallation     bool            `json:"isProductionInstallation,omitempty"`
	LicensingMode                string          `json:"licensingMode"`
	GraceDaysLeft                *int            `json:"graceDaysLeft,omitempty"`
	InstallationDomain           *string         `json:"installationDomain,omitempty"`
	LicensingDocsUrl             *string         `json:"licensingDocsUrl,omitempty"`
	Agent                        *EchoAgentDto   `json:"agent,omitempty"`
}

type PublicProjectConfigDto struct {
	DisplayName        string          `json:"displayName"`
	AdminPortalEnabled bool            `json:"adminPortalEnabled,omitempty"`
	Branding           *PublicBrandDto `json:"branding,omitempty"`
	Auth               PublicAuthDto   `json:"auth"`
	AiChat             PublicAiChatDto `json:"aiChat"`
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

type PasskeyOkResponse struct {
	ResponseBase
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

type RequestUploadUrlResponse struct {
	ResponseBase
	Url *string `json:"url,omitempty"`
}

// @DataContract
type TestFilesIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList[IntegrationTestResultItemDto] `json:"items,omitempty"`
}

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

func (SaveGuestUser) CreateResponse() (r IdResponse) { return }
func (SaveGuestUser) HttpMethod() string             { return "POST" }

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
}

func (SaveEmailUser) CreateResponse() (r IdResponse) { return }
func (SaveEmailUser) HttpMethod() string             { return "POST" }

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
	/** @description Email address the invitation is sent to. */
	// @DataMember
	// @ApiMember(Description="Email address the invitation is sent to.", IsRequired=true)
	Email string `json:"email"`
	/** @description Database integration id. Optional — defaults to the request environment's default integration. */
	// @DataMember
	// @ApiMember(Description="Database integration id. Optional — defaults to the request environment's default integration.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

func (InviteUserRequest) CreateResponse() (r EmptyResponse) { return }
func (InviteUserRequest) HttpMethod() string                { return "POST" }

/** @description Membership */
// @Route("/{version}/membership/auth/{userId}/link-identity", "POST")
// @Api(Description="Membership")
// @DataContract
type LinkIdentityRequest struct {
	CodeMashRequestBase
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	Provider string `json:"provider"`
	// @DataMember
	ProviderToken string `json:"providerToken"`
	// @DataMember
	EmailToVerify *string `json:"emailToVerify,omitempty"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
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

func (ChangePasswordRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (ChangePasswordRequest) HttpMethod() string                    { return "POST" }

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

func (ConfirmPasswordResetRequest) CreateResponse() (r PasskeyOkResponse) { return }
func (ConfirmPasswordResetRequest) HttpMethod() string                    { return "POST" }

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
	// @DataMember
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
	// @DataMember
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
	// @DataMember
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
	// @DataMember
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
}

func (RequestUploadUrlRequest) CreateResponse() (r RequestUploadUrlResponse) { return }
func (RequestUploadUrlRequest) HttpMethod() string                           { return "POST" }

/** @description Files */
// @Route("/{version}/files/{filesIntegrationId}/test", "POST")
// @Api(Description="Files")
// @DataContract
type TestFilesIntegrationRequest struct {
	CodeMashRequestBase
	// @DataMember
	FilesIntegrationId string `json:"filesIntegrationId"`
}

func (TestFilesIntegrationRequest) CreateResponse() (r TestFilesIntegrationResponse) { return }
func (TestFilesIntegrationRequest) HttpMethod() string                               { return "POST" }
