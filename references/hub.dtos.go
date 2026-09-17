//go:build ignore

/* Options:
Date: 2026-09-04 14:56:08
Version: 10.08
Tip: To override a DTO option, remove "//" prefix before updating
BaseUrl: http://localhost:5001

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

type IReturn struct {
}

type IReturnVoid struct {
}

type IHasSessionId struct {
	SessionId *string `json:"sessionId,omitempty"`
}

type IHasBearerToken struct {
	BearerToken *string `json:"bearerToken,omitempty"`
}

type IPost struct {
}

// @DataContract
type EmailProvider string

const (
	EmailProviderSmtp     EmailProvider = "Smtp"
	EmailProviderSendGrid               = "SendGrid"
	EmailProviderMailGun                = "MailGun"
	EmailProviderAwsSes                 = "AwsSes"
	EmailProviderFake                   = "Fake"
)

type EmailIntegrationRequest struct {
	IntegrationId   *string       `json:"integrationId,omitempty"`
	Provider        EmailProvider `json:"provider,omitempty"`
	IntegrationName string        `json:"integrationName"`
	IsEnabled       bool          `json:"isEnabled,omitempty"`
	EmailAddress    string        `json:"emailAddress"`
	EmailSenderName *string       `json:"emailSenderName,omitempty"`
}

type SmtpPorts string

const (
	SmtpPortsDefault  SmtpPorts = "Default"
	SmtpPortsSsl                = "Ssl"
	SmtpPortsTls                = "Tls"
	SmtpPortsFallback           = "Fallback"
)

type SmtpEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider EmailProvider `json:"provider,omitempty"`
	Domain   string        `json:"domain"`
	Port     SmtpPorts     `json:"port,omitempty"`
	UserName string        `json:"userName"`
	Password string        `json:"password"`
}

type AwsIntegrationType string

const (
	AwsIntegrationTypeIam              AwsIntegrationType = "Iam"
	AwsIntegrationTypeCrossAccountRole                    = "CrossAccountRole"
)

type AwsSesEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider         EmailProvider      `json:"provider,omitempty"`
	IntegrationType  AwsIntegrationType `json:"integrationType,omitempty"`
	AwsRegion        string             `json:"awsRegion"`
	EmailIdentityArn string             `json:"emailIdentityArn"`
	ConfigurationSet *string            `json:"configurationSet,omitempty"`
	RoleArn          *string            `json:"roleArn,omitempty"`
	ExternalId       *string            `json:"externalId,omitempty"`
	AccessKey        *string            `json:"accessKey,omitempty"`
	SecretKey        *string            `json:"secretKey,omitempty"`
}

type SendGridEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider EmailProvider `json:"provider,omitempty"`
	ApiKey   string        `json:"apiKey"`
}

type MailGunRegion string

const (
	MailGunRegionUs MailGunRegion = "Us"
	MailGunRegionEu               = "Eu"
)

type MailGunEmailIntegrationRequest struct {
	EmailIntegrationRequest
	Provider          EmailProvider `json:"provider,omitempty"`
	Domain            string        `json:"domain"`
	ApiKey            string        `json:"apiKey"`
	WebhookSigningKey string        `json:"webhookSigningKey"`
	Region            MailGunRegion `json:"region,omitempty"`
}

type EmailCampaignRecipientsSourceTypes string

const (
	EmailCampaignRecipientsSourceTypesAllUsers       EmailCampaignRecipientsSourceTypes = "AllUsers"
	EmailCampaignRecipientsSourceTypesSpecifiedUsers                                    = "SpecifiedUsers"
	EmailCampaignRecipientsSourceTypesAccountUsers                                      = "AccountUsers"
	EmailCampaignRecipientsSourceTypesEmail                                             = "Email"
	EmailCampaignRecipientsSourceTypesCollection                                        = "Collection"
)

type TokenMappingResolverType string

const (
	TokenMappingResolverTypeNotSet          TokenMappingResolverType = "NotSet"
	TokenMappingResolverTypeCustom                                   = "Custom"
	TokenMappingResolverTypeProject                                  = "Project"
	TokenMappingResolverTypeProjectSocials                           = "ProjectSocials"
	TokenMappingResolverTypeInitiator                                = "Initiator"
	TokenMappingResolverTypeRecipient                                = "Recipient"
	TokenMappingResolverTypeSchemaRecord                             = "SchemaRecord"
	TokenMappingResolverTypeTargetUser                               = "TargetUser"
	TokenMappingResolverTypeTagDefinitions                           = "TagDefinitions"
	TokenMappingResolverTypeEmailSignatures                          = "EmailSignatures"
	TokenMappingResolverTypeCampaign                                 = "Campaign"
	TokenMappingResolverTypeTemplate                                 = "Template"
	TokenMappingResolverTypeEmailFooters                             = "EmailFooters"
	TokenMappingResolverTypeOld                                      = "Old"
	TokenMappingResolverTypeNew                                      = "New"
)

// @DataContract
type TokenMappingDto struct {
	// @DataMember
	Key string `json:"key"`
	// @DataMember
	Value string `json:"value"`
	// @DataMember
	Resolver TokenMappingResolverType `json:"resolver,omitempty"`
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

type EmailToAllUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source     EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	RolesNames []string                           `json:"rolesNames,omitempty"`
	UserTags   []string                           `json:"userTags,omitempty"`
}

type EmailToAccountUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	UserRecipients      []string                           `json:"userRecipients"`
	UserCc              []string                           `json:"userCc,omitempty"`
	UserBcc             []string                           `json:"userBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy,omitempty"`
}

type CollectionEmailCampaignRecipientField string

const (
	CollectionEmailCampaignRecipientFieldUser  CollectionEmailCampaignRecipientField = "User"
	CollectionEmailCampaignRecipientFieldEmail                                       = "Email"
)

type EmailToCollectionRecordsDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source     EmailCampaignRecipientsSourceTypes    `json:"source,omitempty"`
	Fields     []string                              `json:"fields"`
	SchemaName string                                `json:"schemaName"`
	FieldType  CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	RoleNames  []string                              `json:"roleNames,omitempty"`
	Languages  []string                              `json:"languages,omitempty"`
}

type EmailToEmailsDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	Recipients          []string                           `json:"recipients"`
	RecipientsCc        []string                           `json:"recipientsCc,omitempty"`
	RecipientsBcc       []string                           `json:"recipientsBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy,omitempty"`
}

type EmailToUsersDeliverySettingsRequest struct {
	EmailCampaignRequest
	Source              EmailCampaignRecipientsSourceTypes `json:"source,omitempty"`
	UserRecipients      []string                           `json:"userRecipients"`
	UserCc              []string                           `json:"userCc,omitempty"`
	UserBcc             []string                           `json:"userBcc,omitempty"`
	SingleEmailStrategy bool                               `json:"singleEmailStrategy,omitempty"`
}

type TriggerType string

const (
	TriggerTypeMembership TriggerType = "Membership"
	TriggerTypeSchema                 = "Schema"
	TriggerTypeFiles                  = "Files"
	TriggerTypePayments               = "Payments"
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

type SaveTriggerRequest struct {
	Type           TriggerType      `json:"type,omitempty"`
	TriggerId      *string          `json:"triggerId,omitempty"`
	Name           string           `json:"name"`
	Description    *string          `json:"description,omitempty"`
	IsEnabled      bool             `json:"isEnabled,omitempty"`
	PreExecuteCode *string          `json:"preExecuteCode,omitempty"`
	Action         TriggerActionDto `json:"action"`
}

type MembershipTriggerType string

const (
	MembershipTriggerTypeOnRegistered  MembershipTriggerType = "OnRegistered"
	MembershipTriggerTypeOnInvited                           = "OnInvited"
	MembershipTriggerTypeOnVerified                          = "OnVerified"
	MembershipTriggerTypeOnUpdated                           = "OnUpdated"
	MembershipTriggerTypeOnDeleted                           = "OnDeleted"
	MembershipTriggerTypeOnBlocked                           = "OnBlocked"
	MembershipTriggerTypeOnReactivated                       = "OnReactivated"
	MembershipTriggerTypeOnUserCreated                       = "OnUserCreated"
)

type MembershipTriggerRequest struct {
	SaveTriggerRequest
	Type TriggerType           `json:"type,omitempty"`
	When MembershipTriggerType `json:"when,omitempty"`
}

type SchemaTriggerType string

const (
	SchemaTriggerTypeOnInserted SchemaTriggerType = "OnInserted"
	SchemaTriggerTypeOnDeleted                    = "OnDeleted"
	SchemaTriggerTypeOnUpdated                    = "OnUpdated"
)

type SchemaTriggerRequest struct {
	SaveTriggerRequest
	Type              TriggerType       `json:"type,omitempty"`
	SchemaId          string            `json:"schemaId"`
	When              SchemaTriggerType `json:"when,omitempty"`
	ConfigurationCode *string           `json:"configurationCode,omitempty"`
}

type FilesTriggerType string

const (
	FilesTriggerTypeOnFileUploaded FilesTriggerType = "OnFileUploaded"
	FilesTriggerTypeOnFileDeleted                   = "OnFileDeleted"
)

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
}

type FilesTriggerRequest struct {
	SaveTriggerRequest
	Type    TriggerType        `json:"type,omitempty"`
	When    FilesTriggerType   `json:"when,omitempty"`
	FileRef FileResourceRefDto `json:"fileRef"`
}

type PaymentTriggerType string

const (
	PaymentTriggerTypeOnOrderCreated        PaymentTriggerType = "OnOrderCreated"
	PaymentTriggerTypeOnOrderPaid                              = "OnOrderPaid"
	PaymentTriggerTypeOnWebhookCallReceived                    = "OnWebhookCallReceived"
)

type PaymentTriggerRequest struct {
	SaveTriggerRequest
	Type         TriggerType        `json:"type,omitempty"`
	When         PaymentTriggerType `json:"when,omitempty"`
	Integrations []string           `json:"integrations,omitempty"`
	Events       []string           `json:"events,omitempty"`
}

type DatabaseProvider string

const (
	DatabaseProviderMongoDbConnectionString         DatabaseProvider = "MongoDbConnectionString"
	DatabaseProviderCodeMashMongoDbAtlasFlexManaged                  = "CodeMashMongoDbAtlasFlexManaged"
)

type DatabaseIntegrationRequest struct {
	IntegrationId   *string          `json:"integrationId,omitempty"`
	Provider        DatabaseProvider `json:"provider,omitempty"`
	IntegrationName string           `json:"integrationName"`
	IsEnabled       bool             `json:"isEnabled,omitempty"`
}

type MongoDbConnectionStringDatabaseIntegrationRequest struct {
	DatabaseIntegrationRequest
	Provider         DatabaseProvider `json:"provider,omitempty"`
	DatabaseName     *string          `json:"databaseName,omitempty"`
	ConnectionString string           `json:"connectionString"`
}

type MongoDbAtlasFlexManagedDatabaseIntegrationRequest struct {
	DatabaseIntegrationRequest
	Provider         DatabaseProvider `json:"provider,omitempty"`
	NorbixRegionCode string           `json:"norbixRegionCode"`
}

type FilesIntegrationRequest struct {
	IntegrationId   *string      `json:"integrationId,omitempty"`
	Provider        FileProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
}

type GoogleDriveFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider              FileProvider `json:"provider,omitempty"`
	RootFolderId          *string      `json:"rootFolderId,omitempty"`
	ServiceAccountJsonKey string       `json:"serviceAccountJsonKey"`
}

type FtpFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider FileProvider `json:"provider,omitempty"`
	Host     string       `json:"host"`
	Port     int          `json:"port,omitempty"`
	RootPath *string      `json:"rootPath,omitempty"`
	UseSsl   bool         `json:"useSsl,omitempty"`
	Username string       `json:"username"`
	Password string       `json:"password"`
}

type DropBoxFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider    FileProvider `json:"provider,omitempty"`
	RootPath    *string      `json:"rootPath,omitempty"`
	AccessToken string       `json:"accessToken"`
}

type AppleICloudFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider            FileProvider `json:"provider,omitempty"`
	ContainerIdentifier string       `json:"containerIdentifier"`
	RelativePath        *string      `json:"relativePath,omitempty"`
	KeyId               string       `json:"keyId"`
	TeamId              string       `json:"teamId"`
	BundleId            string       `json:"bundleId"`
	P8PrivateKey        string       `json:"p8PrivateKey"`
}

type AwsS3IntegrationType string

const (
	AwsS3IntegrationTypeIam              AwsS3IntegrationType = "Iam"
	AwsS3IntegrationTypeCrossAccountRole                      = "CrossAccountRole"
)

type AwsS3FilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider        FileProvider         `json:"provider,omitempty"`
	IntegrationType AwsS3IntegrationType `json:"integrationType,omitempty"`
	BucketName      string               `json:"bucketName"`
	Region          string               `json:"region"`
	RoleArn         *string              `json:"roleArn,omitempty"`
	ExternalId      *string              `json:"externalId,omitempty"`
	AccessKey       *string              `json:"accessKey,omitempty"`
	SecretKey       *string              `json:"secretKey,omitempty"`
}

type GoogleCloudFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider              FileProvider `json:"provider,omitempty"`
	BucketName            string       `json:"bucketName"`
	ServiceAccountJsonKey string       `json:"serviceAccountJsonKey"`
}

type AzureBlobFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider         FileProvider `json:"provider,omitempty"`
	BlobName         string       `json:"blobName"`
	ConnectionString string       `json:"connectionString"`
}

type LocalFilesIntegrationRequest struct {
	FilesIntegrationRequest
	Provider FileProvider `json:"provider,omitempty"`
	RootPath *string      `json:"rootPath,omitempty"`
}

type LoggingProvider string

const (
	LoggingProviderConsole          LoggingProvider = "Console"
	LoggingProviderNorbixLogging                    = "NorbixLogging"
	LoggingProviderDataDog                          = "DataDog"
	LoggingProviderNewRelic                         = "NewRelic"
	LoggingProviderSentry                           = "Sentry"
	LoggingProviderGrafanaLoki                      = "GrafanaLoki"
	LoggingProviderAxiom                            = "Axiom"
	LoggingProviderElasticCloud                     = "ElasticCloud"
	LoggingProviderAWSCloudWatch                    = "AWSCloudWatch"
	LoggingProviderGCPCloudLogging                  = "GCPCloudLogging"
	LoggingProviderAzureMonitorLogs                 = "AzureMonitorLogs"
	LoggingProviderGenericHttp                      = "GenericHttp"
	LoggingProviderKafka                            = "Kafka"
	LoggingProviderAMQP                             = "AMQP"
	LoggingProviderPrometheus                       = "Prometheus"
	LoggingProviderAzureOTel                        = "AzureOTel"
	LoggingProviderSplunk                           = "Splunk"
	LoggingProviderElasticSearch                    = "ElasticSearch"
	LoggingProviderKibana                           = "Kibana"
	LoggingProviderLocalFile                        = "LocalFile"
	LoggingProviderAWSS3                            = "AWSS3"
	LoggingProviderAWSKinesis                       = "AWSKinesis"
	LoggingProviderMongoDB                          = "MongoDB"
	LoggingProviderInternalKafka                    = "InternalKafka"
)

type LoggingIntegrationRequest struct {
	IntegrationId   *string         `json:"integrationId,omitempty"`
	Provider        LoggingProvider `json:"provider,omitempty"`
	IntegrationName string          `json:"integrationName"`
	IsEnabled       bool            `json:"isEnabled,omitempty"`
}

type AmqpLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	Host        string          `json:"host"`
	Port        int             `json:"port,omitempty"`
	VirtualHost string          `json:"virtualHost"`
	Exchange    string          `json:"exchange"`
	RoutingKey  string          `json:"routingKey"`
	Username    string          `json:"username"`
	Password    string          `json:"password"`
}

type AwsKinesisLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider   LoggingProvider `json:"provider,omitempty"`
	StreamName string          `json:"streamName"`
	Region     string          `json:"region"`
	AccessKey  string          `json:"accessKey"`
	SecretKey  string          `json:"secretKey"`
}

type AwsS3LoggingIntegrationType string

const (
	AwsS3LoggingIntegrationTypeIam              AwsS3LoggingIntegrationType = "Iam"
	AwsS3LoggingIntegrationTypeCrossAccountRole                             = "CrossAccountRole"
)

type AwsS3LoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider        LoggingProvider             `json:"provider,omitempty"`
	IntegrationType AwsS3LoggingIntegrationType `json:"integrationType,omitempty"`
	BucketName      string                      `json:"bucketName"`
	Region          string                      `json:"region"`
	RoleArn         *string                     `json:"roleArn,omitempty"`
	ExternalId      *string                     `json:"externalId,omitempty"`
	AccessKey       *string                     `json:"accessKey,omitempty"`
	SecretKey       *string                     `json:"secretKey,omitempty"`
}

type NewRelicLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	Region      string          `json:"region"`
	ServiceName string          `json:"serviceName"`
	ApiKey      string          `json:"apiKey"`
}

type MongoDbLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	DatabaseName     *string         `json:"databaseName,omitempty"`
	ConnectionString string          `json:"connectionString"`
}

type KafkaLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	BootstrapServers string          `json:"bootstrapServers"`
	Topic            string          `json:"topic"`
	SecurityProtocol *string         `json:"securityProtocol,omitempty"`
	SaslUsername     *string         `json:"saslUsername,omitempty"`
	SaslPassword     *string         `json:"saslPassword,omitempty"`
}

type PrometheusLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	EndpointUrl string          `json:"endpointUrl"`
	JobName     *string         `json:"jobName,omitempty"`
	BearerToken *string         `json:"bearerToken,omitempty"`
}

type DataDogLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider    LoggingProvider `json:"provider,omitempty"`
	Site        string          `json:"site"`
	ServiceName string          `json:"serviceName"`
	Environment string          `json:"environment"`
	ApiKey      string          `json:"apiKey"`
}

type InternalKafkaLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	BootstrapServers string          `json:"bootstrapServers"`
	Topic            string          `json:"topic"`
	SecurityProtocol *string         `json:"securityProtocol,omitempty"`
	SaslUsername     *string         `json:"saslUsername,omitempty"`
	SaslPassword     *string         `json:"saslPassword,omitempty"`
}

type ElasticSearchLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider LoggingProvider `json:"provider,omitempty"`
	Uri      string          `json:"uri"`
	Index    string          `json:"index"`
	Username *string         `json:"username,omitempty"`
	Password *string         `json:"password,omitempty"`
}

type SplunkLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider       LoggingProvider `json:"provider,omitempty"`
	HecEndpointUrl string          `json:"hecEndpointUrl"`
	Index          string          `json:"index"`
	HecToken       string          `json:"hecToken"`
}

type AzureOtelLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider         LoggingProvider `json:"provider,omitempty"`
	EndpointUrl      string          `json:"endpointUrl"`
	ResourceName     string          `json:"resourceName"`
	ConnectionString string          `json:"connectionString"`
}

type KibanaLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider LoggingProvider `json:"provider,omitempty"`
	Uri      string          `json:"uri"`
	SpaceId  *string         `json:"spaceId,omitempty"`
	ApiKey   string          `json:"apiKey"`
}

type LocalFileLoggingIntegrationRequest struct {
	LoggingIntegrationRequest
	Provider LoggingProvider `json:"provider,omitempty"`
	RootPath *string         `json:"rootPath,omitempty"`
}

type MembershipProvider string

const (
	MembershipProviderAppleSignIn  MembershipProvider = "AppleSignIn"
	MembershipProviderGoogleSignIn                    = "GoogleSignIn"
	MembershipProviderGoogle                          = "Google"
	MembershipProviderFacebook                        = "Facebook"
	MembershipProviderX                               = "X"
	MembershipProviderGitHub                          = "GitHub"
	MembershipProviderLinkedIn                        = "LinkedIn"
	MembershipProviderOkta                            = "Okta"
	MembershipProviderMicrosoft                       = "Microsoft"
)

type MembershipIntegrationRequest struct {
	IntegrationId   *string            `json:"integrationId,omitempty"`
	Provider        MembershipProvider `json:"provider,omitempty"`
	IntegrationName string             `json:"integrationName"`
	IsEnabled       bool               `json:"isEnabled,omitempty"`
}

type DisplayName struct {
	Value string `json:"value"`
}

type RoleName struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	// @Ignore()
	IsAdministrator bool `json:"isAdministrator,omitempty"`
	// @Ignore()
	IsAuthenticated bool `json:"isAuthenticated,omitempty"`
	// @Ignore()
	IsGuest bool `json:"isGuest,omitempty"`
	// @Ignore()
	IsRootRole bool `json:"isRootRole,omitempty"`
	// @Ignore()
	IsCollaboratorRole bool `json:"isCollaboratorRole,omitempty"`
	// @Ignore()
	IsProjectSystemRole bool `json:"isProjectSystemRole,omitempty"`
	// @Ignore()
	IsAccountSystemRole bool `json:"isAccountSystemRole,omitempty"`
	// @Ignore()
	IsSystemRole bool `json:"isSystemRole,omitempty"`
}

type OAuthModeConfig struct {
	Name               DisplayName `json:"name"`
	CallbackUrl        string      `json:"callbackUrl"`
	LogoutUrl          *string     `json:"logoutUrl,omitempty"`
	FailureRedirectUrl *string     `json:"failureRedirectUrl,omitempty"`
	RoleName           *RoleName   `json:"roleName,omitempty"`
}

type OktaMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	Domain       string             `json:"domain"`
	ClientId     string             `json:"clientId"`
	ClientSecret string             `json:"clientSecret"`
	OAuthModes   []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type XMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	ApiKey       string             `json:"apiKey"`
	ApiSecretKey string             `json:"apiSecretKey"`
	OAuthModes   []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type GoogleMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	ClientId     string             `json:"clientId"`
	ClientSecret string             `json:"clientSecret"`
	OAuthModes   []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type MicrosoftMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	TenantId     string             `json:"tenantId"`
	ClientId     string             `json:"clientId"`
	ClientSecret string             `json:"clientSecret"`
	OAuthModes   []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type GitHubMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	ClientId     string             `json:"clientId"`
	ClientSecret string             `json:"clientSecret"`
	OAuthModes   []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type MetaMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider   MembershipProvider `json:"provider,omitempty"`
	AppId      string             `json:"appId"`
	AppSecret  string             `json:"appSecret"`
	OAuthModes []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type AppleMembershipIntegrationRequest struct {
	MembershipIntegrationRequest
	Provider     MembershipProvider `json:"provider,omitempty"`
	TeamId       string             `json:"teamId"`
	AppBundleId  string             `json:"appBundleId"`
	ServiceId    string             `json:"serviceId"`
	KeyId        string             `json:"keyId"`
	PrivateKey   string             `json:"privateKey"`
	IsProduction bool               `json:"isProduction,omitempty"`
	OAuthModes   []OAuthModeConfig  `json:"oAuthModes,omitempty"`
}

type PaymentGatewayPlatform string

const (
	PaymentGatewayPlatformStripe       PaymentGatewayPlatform = "Stripe"
	PaymentGatewayPlatformAdyen                               = "Adyen"
	PaymentGatewayPlatformPaddle                              = "Paddle"
	PaymentGatewayPlatformLemonSqueezy                        = "LemonSqueezy"
	PaymentGatewayPlatformAppleInApp                          = "AppleInApp"
	PaymentGatewayPlatformGoogleInApp                         = "GoogleInApp"
	PaymentGatewayPlatformShopify                             = "Shopify"
	PaymentGatewayPlatformWooCommerce                         = "WooCommerce"
	PaymentGatewayPlatformMagento                             = "Magento"
	PaymentGatewayPlatformPayPal                              = "PayPal"
	PaymentGatewayPlatformBraintree                           = "Braintree"
	PaymentGatewayPlatformAuthorizeNet                        = "AuthorizeNet"
	PaymentGatewayPlatformCheckOutCom                         = "CheckOutCom"
	PaymentGatewayPlatformMollie                              = "Mollie"
	PaymentGatewayPlatformWorldpay                            = "Worldpay"
)

type PaymentIntegrationRequest struct {
	IntegrationId   *string                `json:"integrationId,omitempty"`
	Provider        PaymentGatewayPlatform `json:"provider,omitempty"`
	IntegrationName string                 `json:"integrationName"`
	IsEnabled       bool                   `json:"isEnabled,omitempty"`
}

type LemonSqueezyPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	StoreId              string                 `json:"storeId"`
	ApiKey               string                 `json:"apiKey"`
	WebhookSigningSecret string                 `json:"webhookSigningSecret"`
	IsTestMode           bool                   `json:"isTestMode,omitempty"`
}

type AdyenPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider        PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantAccount string                 `json:"merchantAccount"`
	ApiKey          string                 `json:"apiKey"`
	Environment     string                 `json:"environment"`
	WebhookId       *string                `json:"webhookId,omitempty"`
	WebhookHmacKey  *string                `json:"webhookHmacKey,omitempty"`
}

type MolliePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	ProfileId            string                 `json:"profileId"`
	ApiKey               string                 `json:"apiKey"`
	IsTestMode           bool                   `json:"isTestMode,omitempty"`
	WebhookSigningSecret *string                `json:"webhookSigningSecret,omitempty"`
}

type PaddlePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider                 PaymentGatewayPlatform `json:"provider,omitempty"`
	ApiKey                   string                 `json:"apiKey"`
	WebhookEndpointSecretKey string                 `json:"webhookEndpointSecretKey"`
	Environment              string                 `json:"environment"`
	ClientSideToken          *string                `json:"clientSideToken,omitempty"`
}

type PayPalPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider     PaymentGatewayPlatform `json:"provider,omitempty"`
	ClientId     string                 `json:"clientId"`
	ClientSecret string                 `json:"clientSecret"`
	Environment  string                 `json:"environment"`
	BrandName    *string                `json:"brandName,omitempty"`
	WebhookId    *string                `json:"webhookId,omitempty"`
}

type StripePaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider             PaymentGatewayPlatform `json:"provider,omitempty"`
	PublishableKey       string                 `json:"publishableKey"`
	SecretKey            string                 `json:"secretKey"`
	WebhookSigningSecret string                 `json:"webhookSigningSecret"`
	WebhookEndpointId    *string                `json:"webhookEndpointId,omitempty"`
	DefaultCurrency      *string                `json:"defaultCurrency,omitempty"`
}

type AppleInAppPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider                              PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantIdentifier                    string                 `json:"merchantIdentifier"`
	MerchantDomain                        string                 `json:"merchantDomain"`
	DisplayName                           string                 `json:"displayName"`
	MerchantIdentityCertificateP12Base64  string                 `json:"merchantIdentityCertificateP12Base64"`
	MerchantIdentityCertificatePassword   string                 `json:"merchantIdentityCertificatePassword"`
	PaymentProcessingCertificateP12Base64 string                 `json:"paymentProcessingCertificateP12Base64"`
	PaymentProcessingCertificatePassword  string                 `json:"paymentProcessingCertificatePassword"`
	WebhookBundleId                       *string                `json:"webhookBundleId,omitempty"`
}

type GoogleInAppPaymentIntegrationRequest struct {
	PaymentIntegrationRequest
	Provider           PaymentGatewayPlatform `json:"provider,omitempty"`
	MerchantId         string                 `json:"merchantId"`
	MerchantName       string                 `json:"merchantName"`
	Gateway            string                 `json:"gateway"`
	PrivateKeyOrToken  string                 `json:"privateKeyOrToken"`
	GatewayMerchantId  *string                `json:"gatewayMerchantId,omitempty"`
	WebhookPackageName *string                `json:"webhookPackageName,omitempty"`
}

// @DataContract
type PushProvider string

const (
	PushProviderAppleApns            PushProvider = "AppleApns"
	PushProviderSafariWeb                         = "SafariWeb"
	PushProviderSafariPush                        = "SafariPush"
	PushProviderAndroidFirebase                   = "AndroidFirebase"
	PushProviderChromeWeb                         = "ChromeWeb"
	PushProviderFirefoxWeb                        = "FirefoxWeb"
	PushProviderEdgeWeb                           = "EdgeWeb"
	PushProviderChromePush                        = "ChromePush"
	PushProviderCodeMashIosApp                    = "CodeMashIosApp"
	PushProviderCodeMashAndroidApp                = "CodeMashAndroidApp"
	PushProviderCodeMashSafariPlugin              = "CodeMashSafariPlugin"
	PushProviderCodeMashSafariWeb                 = "CodeMashSafariWeb"
	PushProviderCodeMashChromePlugin              = "CodeMashChromePlugin"
	PushProviderCodeMashChromeWeb                 = "CodeMashChromeWeb"
	PushProviderExpo                              = "Expo"
	PushProviderFake                              = "Fake"
)

type PushIntegrationRequest struct {
	IntegrationId   *string      `json:"integrationId,omitempty"`
	Provider        PushProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
}

type EdgeWebPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey"`
	VapidPrivateKey string       `json:"vapidPrivateKey"`
	Subject         *string      `json:"subject,omitempty"`
}

type ChromePluginPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	ExtensionId     string       `json:"extensionId"`
	VapidPublicKey  string       `json:"vapidPublicKey"`
	VapidPrivateKey string       `json:"vapidPrivateKey"`
	Subject         *string      `json:"subject,omitempty"`
}

type SafariPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider             PushProvider `json:"provider,omitempty"`
	WebsitePushId        string       `json:"websitePushId"`
	CertificateP12Base64 string       `json:"certificateP12Base64"`
	CertificatePassword  string       `json:"certificatePassword"`
}

type ChromeWebPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey"`
	VapidPrivateKey string       `json:"vapidPrivateKey"`
	Subject         *string      `json:"subject,omitempty"`
}

type FirefoxWebPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider        PushProvider `json:"provider,omitempty"`
	VapidPublicKey  string       `json:"vapidPublicKey"`
	VapidPrivateKey string       `json:"vapidPrivateKey"`
	Subject         *string      `json:"subject,omitempty"`
}

type AndroidFirebasePushIntegrationRequest struct {
	PushIntegrationRequest
	Provider           PushProvider `json:"provider,omitempty"`
	ProjectId          string       `json:"projectId"`
	ClientEmail        string       `json:"clientEmail"`
	ServiceAccountJson string       `json:"serviceAccountJson"`
}

type AppleApnsPushIntegrationRequest struct {
	PushIntegrationRequest
	Provider     PushProvider `json:"provider,omitempty"`
	TeamId       string       `json:"teamId"`
	AppBundleId  string       `json:"appBundleId"`
	KeyId        string       `json:"keyId"`
	PrivateKey   string       `json:"privateKey"`
	IsProduction bool         `json:"isProduction,omitempty"`
}

type CodeProvider string

const (
	CodeProviderAwsLambda            CodeProvider = "AwsLambda"
	CodeProviderAzureFunctions                    = "AzureFunctions"
	CodeProviderGoogleCloudFunctions              = "GoogleCloudFunctions"
	CodeProviderPipedream                         = "Pipedream"
	CodeProviderZapier                            = "Zapier"
	CodeProviderCloudflareWorkers                 = "CloudflareWorkers"
	CodeProviderVercel                            = "Vercel"
	CodeProviderNetlify                           = "Netlify"
	CodeProviderSupabaseEdge                      = "SupabaseEdge"
	CodeProviderModal                             = "Modal"
)

type CodeIntegrationRequest struct {
	IntegrationId   *string      `json:"integrationId,omitempty"`
	Provider        CodeProvider `json:"provider,omitempty"`
	IntegrationName string       `json:"integrationName"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
}

type AwsLambdaIntegrationType string

const (
	AwsLambdaIntegrationTypeIam              AwsLambdaIntegrationType = "Iam"
	AwsLambdaIntegrationTypeCrossAccountRole                          = "CrossAccountRole"
)

type AwsLambdaCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider        CodeProvider             `json:"provider,omitempty"`
	IntegrationType AwsLambdaIntegrationType `json:"integrationType,omitempty"`
	Region          string                   `json:"region"`
	RoleArn         *string                  `json:"roleArn,omitempty"`
	ExternalId      *string                  `json:"externalId,omitempty"`
	AccessKey       *string                  `json:"accessKey,omitempty"`
	SecretKey       *string                  `json:"secretKey,omitempty"`
}

type AzureFunctionsCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider              CodeProvider `json:"provider,omitempty"`
	FunctionAppName       string       `json:"functionAppName"`
	ResourceGroup         *string      `json:"resourceGroup,omitempty"`
	ConnectionStringOrKey string       `json:"connectionStringOrKey"`
}

type GoogleCloudFunctionsCodeIntegrationRequest struct {
	CodeIntegrationRequest
	Provider              CodeProvider `json:"provider,omitempty"`
	ProjectId             string       `json:"projectId"`
	Region                *string      `json:"region,omitempty"`
	ServiceAccountJsonKey string       `json:"serviceAccountJsonKey"`
}

type LlmProvider string

const (
	LlmProviderOpenAI       LlmProvider = "OpenAI"
	LlmProviderAnthropic                = "Anthropic"
	LlmProviderOllama                   = "Ollama"
	LlmProviderGroq                     = "Groq"
	LlmProviderGoogle                   = "Google"
	LlmProviderMistral                  = "Mistral"
	LlmProviderOpenRouter               = "OpenRouter"
	LlmProviderGrok                     = "Grok"
	LlmProviderNorbixHosted             = "NorbixHosted"
)

type LlmIntegrationRequest struct {
	IntegrationId   *string     `json:"integrationId,omitempty"`
	Provider        LlmProvider `json:"provider,omitempty"`
	IntegrationName string      `json:"integrationName"`
	IsEnabled       bool        `json:"isEnabled,omitempty"`
	Endpoint        *string     `json:"endpoint,omitempty"`
	DefaultModel    *string     `json:"defaultModel,omitempty"`
}

type OllamaLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
}

type OpenRouterLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type MistralLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type GrokLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type GroqLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type GoogleLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type AnthropicLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type OpenAiLlmIntegrationRequest struct {
	LlmIntegrationRequest
	Provider LlmProvider `json:"provider,omitempty"`
	ApiKey   string      `json:"apiKey"`
}

type McpProvider string

const (
	McpProviderDocker         McpProvider = "Docker"
	McpProviderObsidian                   = "Obsidian"
	McpProviderGoogleCalendar             = "GoogleCalendar"
	McpProviderStripe                     = "Stripe"
	McpProviderGitHub                     = "GitHub"
	McpProviderMongoDb                    = "MongoDb"
	McpProviderPlaywright                 = "Playwright"
	McpProviderBraveSearch                = "BraveSearch"
)

type McpTransport string

const (
	McpTransportSse        McpTransport = "Sse"
	McpTransportHttpStream              = "HttpStream"
	McpTransportStdio                   = "Stdio"
)

type McpIntegrationRequest struct {
	IntegrationId   *string      `json:"integrationId,omitempty"`
	Provider        McpProvider  `json:"provider,omitempty"`
	Transport       McpTransport `json:"transport,omitempty"`
	IntegrationName string       `json:"integrationName"`
	IsEnabled       bool         `json:"isEnabled,omitempty"`
	Name            string       `json:"name"`
	Category        string       `json:"category"`
	Description     string       `json:"description"`
	Icon            string       `json:"icon"`
}

type PlaywrightMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	Command   *string      `json:"command,omitempty"`
	Args      []string     `json:"args,omitempty"`
	Headless  string       `json:"headless"`
}

type MongoDbMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider         McpProvider  `json:"provider,omitempty"`
	Transport        McpTransport `json:"transport,omitempty"`
	Command          *string      `json:"command,omitempty"`
	Args             []string     `json:"args,omitempty"`
	ConnectionString string       `json:"connectionString"`
}

type GitHubMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider    McpProvider  `json:"provider,omitempty"`
	Transport   McpTransport `json:"transport,omitempty"`
	ServerUrl   string       `json:"serverUrl"`
	AccessToken string       `json:"accessToken"`
}

type StripeMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	ServerUrl string       `json:"serverUrl"`
	ApiKey    string       `json:"apiKey"`
}

type BraveSearchMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	ServerUrl string       `json:"serverUrl"`
	ApiKey    string       `json:"apiKey"`
}

type ObsidianMcpIntegrationRequest struct {
	McpIntegrationRequest
	Provider             McpProvider       `json:"provider,omitempty"`
	Transport            McpTransport      `json:"transport,omitempty"`
	Command              *string           `json:"command,omitempty"`
	Args                 []string          `json:"args,omitempty"`
	EnvironmentVariables map[string]string `json:"environmentVariables,omitempty"`
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
type SmsMessageContentDto struct {
	// @DataMember
	Subject string `json:"subject"`
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

type SystemEmailTemplateTheme string

const (
	SystemEmailTemplateThemeText     SystemEmailTemplateTheme = "Text"
	SystemEmailTemplateThemeBranded                           = "Branded"
	SystemEmailTemplateThemeCreative                          = "Creative"
)

// @DataContract
type SystemEmailTemplateDto struct {
	EmailTemplateDto
	// @DataMember
	ImagePreview string `json:"imagePreview"`
	// @DataMember
	Theme SystemEmailTemplateTheme `json:"theme,omitempty"`
	// @DataMember
	SystemGroup string `json:"systemGroup"`
	// @DataMember
	SystemTags []string `json:"systemTags,omitempty"`
	// @DataMember
	ForTrigger *TriggerType `json:"forTrigger,omitempty"`
	// @DataMember
	HiddenSystemEmailTemplate bool `json:"hiddenSystemEmailTemplate,omitempty"`
}

// @Flags()
type RespectTimeZoneSettings int

const (
	RespectTimeZoneSettingsRespectToLastLoginZone           RespectTimeZoneSettings = 1
	RespectTimeZoneSettingsRespectToRegistrationZone        RespectTimeZoneSettings = 2
	RespectTimeZoneSettingsRespectToRegistrationProjectZone RespectTimeZoneSettings = 4
)

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

// @DataContract
type TriggerActionEmailDto struct {
	TriggerActionDto
	// @DataMember
	TemplateId string `json:"templateId"`
	// @DataMember
	DeliverySettings EmailCampaignDeliverySettingsDto `json:"deliverySettings"`
}

type PushCampaignRecipientsSourceTypes string

const (
	PushCampaignRecipientsSourceTypesAllUsers       PushCampaignRecipientsSourceTypes = "AllUsers"
	PushCampaignRecipientsSourceTypesSpecifiedUsers                                   = "SpecifiedUsers"
	PushCampaignRecipientsSourceTypesCollection                                       = "Collection"
	PushCampaignRecipientsSourceTypesDevices                                          = "Devices"
	PushCampaignRecipientsSourceTypesAccountUsers                                     = "AccountUsers"
)

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
type TriggerActionPushDto struct {
	TriggerActionDto
	// @DataMember
	TemplateId string `json:"templateId"`
	// @DataMember
	DeliverySettings PushCampaignDeliverySettingsDto `json:"deliverySettings"`
}

// @DataContract
type CodeDeliverySettingsDto struct {
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
}

// @DataContract
type TriggerActionCodeDto struct {
	TriggerActionDto
	// @DataMember
	FunctionId string `json:"functionId"`
	// @DataMember
	DeliverySettings CodeDeliverySettingsDto `json:"deliverySettings"`
}

// @DataContract
type WebhookDeliverySettingsDto struct {
	// @DataMember
	DestinationIds []string `json:"destinationIds,omitempty"`
	// @DataMember
	EventName *string `json:"eventName,omitempty"`
	// @DataMember
	ContentType *string `json:"contentType,omitempty"`
	// @DataMember
	IncludeRawPayload bool `json:"includeRawPayload,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
}

// @DataContract
type TriggerActionWebhookDto struct {
	TriggerActionDto
	// @DataMember
	DeliverySettings *WebhookDeliverySettingsDto `json:"deliverySettings,omitempty"`
}

type SmsCampaignRecipientsSourceTypes string

const (
	SmsCampaignRecipientsSourceTypesAllUsers       SmsCampaignRecipientsSourceTypes = "AllUsers"
	SmsCampaignRecipientsSourceTypesSpecifiedUsers                                  = "SpecifiedUsers"
	SmsCampaignRecipientsSourceTypesAccountUsers                                    = "AccountUsers"
	SmsCampaignRecipientsSourceTypesPhoneNumbers                                    = "PhoneNumbers"
	SmsCampaignRecipientsSourceTypesCollection                                      = "Collection"
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
type TriggerActionSmsDto struct {
	TriggerActionDto
	// @DataMember
	TemplateId string `json:"templateId"`
	// @DataMember
	DeliverySettings SmsCampaignDeliverySettingsDto `json:"deliverySettings"`
}

// @DataContract
type SseDeliverySettingsDto struct {
	// @DataMember
	Audience string `json:"audience"`
	// @DataMember
	UserAuthIds []string `json:"userAuthIds,omitempty"`
	// @DataMember
	EventName *string `json:"eventName,omitempty"`
	// @DataMember
	PayloadType *string `json:"payloadType,omitempty"`
	// @DataMember
	PayloadTemplate *string `json:"payloadTemplate,omitempty"`
	// @DataMember
	Persist bool `json:"persist,omitempty"`
	// @DataMember
	MappedTokens []TokenMappingDto `json:"mappedTokens,omitempty"`
}

// @DataContract
type TriggerActionSseDto struct {
	TriggerActionDto
	// @DataMember
	DeliverySettings SseDeliverySettingsDto `json:"deliverySettings"`
}

// @DataContract
type TriggerActionMarketplaceDto struct {
	TriggerActionDto
	// @DataMember
	FunctionId string `json:"functionId"`
	// @DataMember
	Payload map[string]string `json:"payload,omitempty"`
}

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

type Env struct {
	Value  string `json:"value"`
	IsProd bool   `json:"isProd,omitempty"`
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

type CodeMashListPaginationRequestBase struct {
	RequestBase
	/** @description ID of your project. Can be passed in a header as norbix-project-id. */
	// @DataMember
	// @ApiMember(DataType="string", Description="ID of your project. Can be passed in a header as norbix-project-id.", IsRequired=true, Name="norbix-project-id", ParameterType="header")
	ProjectId string `json:"projectId"`
	/** @description Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env. */
	// @DataMember
	// @ApiMember(DataType="string", Description="Target environment for this request (e.g. TEST, STAGING). Optional — when omitted the request runs against PROD. Can be passed in a header as norbix-env.", Name="norbix-env", ParameterType="header")
	Env         *string `json:"env,omitempty"`
	ResolvedEnv Env     `json:"resolvedEnv"`
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
	/** @description Paging */
	// @ApiMember(DataType="object", Description="Paging", Name="paging", ParameterType="body")
	Paging *PagingArgs `json:"paging,omitempty"`
}

// @DataContract
type GetTriggers struct {
	CodeMashListPaginationRequestBase
	// @DataMember
	SchemaId *string `json:"schemaId,omitempty"`
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

// @DataContract
type GetTriggersResponse struct {
	ResponseBase
}

// @DataContract
type EmailToAllUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	// @DataMember
	RolesNames []string `json:"rolesNames,omitempty"`
	// @DataMember
	UserTags []string `json:"userTags,omitempty"`
}

// @DataContract
type EmailToAccountUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	// @DataMember
	UserRecipients []string `json:"userRecipients"`
	// @DataMember
	UserCc []string `json:"userCc,omitempty"`
	// @DataMember
	UserBcc []string `json:"userBcc,omitempty"`
	// @DataMember
	SingleEmailStrategy bool `json:"singleEmailStrategy,omitempty"`
}

// @DataContract
type EmailToUsersDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	// @DataMember
	UserRecipients []string `json:"userRecipients"`
	// @DataMember
	UserCc []string `json:"userCc,omitempty"`
	// @DataMember
	UserBcc []string `json:"userBcc,omitempty"`
	// @DataMember
	SingleEmailStrategy bool `json:"singleEmailStrategy,omitempty"`
}

// @DataContract
type EmailToEmailAddressesDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	// @DataMember
	Recipients []string `json:"recipients"`
	// @DataMember
	RecipientsCc []string `json:"recipientsCc,omitempty"`
	// @DataMember
	RecipientsBcc []string `json:"recipientsBcc,omitempty"`
	// @DataMember
	SingleEmailStrategy bool `json:"singleEmailStrategy,omitempty"`
}

// @DataContract
type EmailToCollectionRecordsDeliverySettingsDto struct {
	EmailCampaignDeliverySettingsDto
	// @DataMember
	Fields []string `json:"fields"`
	// @DataMember
	SchemaName string `json:"schemaName"`
	// @DataMember
	FieldType CollectionEmailCampaignRecipientField `json:"fieldType,omitempty"`
	// @DataMember
	RoleNames []string `json:"roleNames,omitempty"`
	// @DataMember
	Languages []string `json:"languages,omitempty"`
}

// @DataContract
type PushToAllUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	// @DataMember
	RolesNames []string `json:"rolesNames,omitempty"`
	// @DataMember
	UserTags []string `json:"userTags,omitempty"`
}

// @DataContract
type PushToUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	// @DataMember
	Recipients []string `json:"recipients"`
}

// @DataContract
type PushToAccountUsersDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	// @DataMember
	Recipients []string `json:"recipients"`
}

// @DataContract
type PushToCollectionRecordsDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
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
}

// @DataContract
type PushDeviceDeliveryTokenDto struct {
}

// @DataContract
type PushToDevicesDeliverySettingsDto struct {
	PushCampaignDeliverySettingsDto
	// @DataMember
	Devices []PushDeviceDeliveryTokenDto `json:"devices"`
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
}

// @DataContract
type SmsToPhoneNumbersDeliverySettingsDto struct {
	SmsCampaignDeliverySettingsDto
	// @DataMember
	PhoneNumbers []string `json:"phoneNumbers"`
}

type IntegrationDto struct {
	ViewId                            string        `json:"viewId"`
	IntegrationName                   string        `json:"integrationName"`
	IsEnabled                         bool          `json:"isEnabled,omitempty"`
	Env                               *string       `json:"env,omitempty"`
	LastIntegrationTestAtUtc          *time.Time    `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      *bool         `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         IReadOnlyList `json:"lastIntegrationTestErrors"`
	HumanDeliveryConfirmedAtUtc       *time.Time    `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool          `json:"requiresHumanDeliveryConfirmation,omitempty"`
}

type LlmIntegrationDto struct {
	IntegrationDto
	Provider      LlmProvider `json:"provider,omitempty"`
	BaseUrl       *string     `json:"baseUrl,omitempty"`
	DefaultModel  *string     `json:"defaultModel,omitempty"`
	IsConfigured  bool        `json:"isConfigured,omitempty"`
	IsSystemOwned bool        `json:"isSystemOwned,omitempty"`
}

type OpenAiLlmIntegrationDto struct {
	LlmIntegrationDto
}

type AnthropicLlmIntegrationDto struct {
	LlmIntegrationDto
}

type OllamaLlmIntegrationDto struct {
	LlmIntegrationDto
}

type GroqLlmIntegrationDto struct {
	LlmIntegrationDto
}

type GoogleLlmIntegrationDto struct {
	LlmIntegrationDto
}

type MistralLlmIntegrationDto struct {
	LlmIntegrationDto
}

type OpenRouterLlmIntegrationDto struct {
	LlmIntegrationDto
}

type GrokLlmIntegrationDto struct {
	LlmIntegrationDto
}

type McpMetadata struct {
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type McpAuth string

const (
	McpAuthOAuth2 McpAuth = "OAuth2"
	McpAuthApiKey         = "ApiKey"
	McpAuthNone           = "None"
)

type McpIntegrationDto struct {
	IntegrationDto
	Provider      McpProvider  `json:"provider,omitempty"`
	Transport     McpTransport `json:"transport,omitempty"`
	Metadata      McpMetadata  `json:"metadata"`
	IsConfigured  bool         `json:"isConfigured,omitempty"`
	IsSystemOwned bool         `json:"isSystemOwned,omitempty"`
	Command       *string      `json:"command,omitempty"`
	Args          []string     `json:"args,omitempty"`
	ServerUrl     *string      `json:"serverUrl,omitempty"`
	Auth          *McpAuth     `json:"auth,omitempty"`
}

type DockerMcpIntegrationDto struct {
	McpIntegrationDto
}

type GoogleCalendarMcpIntegrationDto struct {
	McpIntegrationDto
}

type ObsidianMcpIntegrationDto struct {
	McpIntegrationDto
}

type CodeIntegrationDto struct {
	IntegrationDto
	Provider CodeProvider `json:"provider,omitempty"`
}

type AwsLambdaCrossAccountRoleCodeIntegrationDto struct {
	CodeIntegrationDto
	Region     string `json:"region"`
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type AwsLambdaIamCodeIntegrationDto struct {
	CodeIntegrationDto
	Region string `json:"region"`
}

type AzureFunctionsCodeIntegrationDto struct {
	CodeIntegrationDto
	FunctionAppName string  `json:"functionAppName"`
	ResourceGroup   *string `json:"resourceGroup,omitempty"`
}

type GoogleCloudFunctionsCodeIntegrationDto struct {
	CodeIntegrationDto
	ProjectId string  `json:"projectId"`
	Region    *string `json:"region,omitempty"`
}

type PaymentsIntegrationDto struct {
	IntegrationDto
	GatewayPlatform PaymentGatewayPlatform `json:"gatewayPlatform,omitempty"`
}

type AdyenPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantAccount string  `json:"merchantAccount"`
	Environment     string  `json:"environment"`
	WebhookId       *string `json:"webhookId,omitempty"`
}

type AppleInAppPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantIdentifier string `json:"merchantIdentifier"`
	MerchantDomain     string `json:"merchantDomain"`
	DisplayName        string `json:"displayName"`
}

type GoogleInAppPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantId        string  `json:"merchantId"`
	MerchantName      string  `json:"merchantName"`
	Gateway           string  `json:"gateway"`
	GatewayMerchantId *string `json:"gatewayMerchantId,omitempty"`
}

type LemonSqueezyPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreId    string `json:"storeId"`
	IsTestMode bool   `json:"isTestMode,omitempty"`
}

type MolliePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ProfileId  string `json:"profileId"`
	IsTestMode bool   `json:"isTestMode,omitempty"`
}

type PaddlePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	Environment     string  `json:"environment"`
	ClientSideToken *string `json:"clientSideToken,omitempty"`
}

type PayPalPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ClientId    string  `json:"clientId"`
	Environment string  `json:"environment"`
	BrandName   *string `json:"brandName,omitempty"`
}

type StripePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	PublishableKey    string  `json:"publishableKey"`
	WebhookEndpointId *string `json:"webhookEndpointId,omitempty"`
	DefaultCurrency   *string `json:"defaultCurrency,omitempty"`
}

type ShopifyPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	ShopDomain    string  `json:"shopDomain"`
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

type WooCommercePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreUrl      string  `json:"storeUrl"`
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

type MagentoPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	StoreUrl      string  `json:"storeUrl"`
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

type BraintreePaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantId    string  `json:"merchantId"`
	Environment   string  `json:"environment"`
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

type AuthorizeNetPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantLoginId     string  `json:"merchantLoginId"`
	Environment         string  `json:"environment"`
	WebhookSignatureKey *string `json:"webhookSignatureKey,omitempty"`
}

type CheckOutComPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantAccount string  `json:"merchantAccount"`
	Environment     string  `json:"environment"`
	WebhookSecret   *string `json:"webhookSecret,omitempty"`
}

type WorldpayPaymentIntegrationDto struct {
	PaymentsIntegrationDto
	MerchantCode  string  `json:"merchantCode"`
	Environment   string  `json:"environment"`
	WebhookSecret *string `json:"webhookSecret,omitempty"`
}

type MembershipIntegrationDto struct {
	IntegrationDto
	Provider MembershipProvider `json:"provider,omitempty"`
}

type AppleSignInMembershipIntegrationDto struct {
	MembershipIntegrationDto
	TeamId      string `json:"teamId"`
	AppBundleId string `json:"appBundleId"`
	ServiceId   string `json:"serviceId"`
}

type GitHubMembershipIntegrationDto struct {
	MembershipIntegrationDto
	ClientId string `json:"clientId"`
}

type GoogleMembershipIntegrationDto struct {
	MembershipIntegrationDto
	ClientId string `json:"clientId"`
}

type MetaMembershipIntegrationDto struct {
	MembershipIntegrationDto
	AppId string `json:"appId"`
}

type MicrosoftMembershipIntegrationDto struct {
	MembershipIntegrationDto
	TenantId string `json:"tenantId"`
	ClientId string `json:"clientId"`
}

type OktaMembershipIntegrationDto struct {
	MembershipIntegrationDto
	Domain   string `json:"domain"`
	ClientId string `json:"clientId"`
}

type XMembershipIntegrationDto struct {
	MembershipIntegrationDto
	ApiKey string `json:"apiKey"`
}

type LoggingIntegrationDto struct {
	IntegrationDto
	Provider LoggingProvider `json:"provider,omitempty"`
}

type AmqpLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Host        string `json:"host"`
	Port        int    `json:"port,omitempty"`
	VirtualHost string `json:"virtualHost"`
	Exchange    string `json:"exchange"`
	RoutingKey  string `json:"routingKey"`
}

type AwsKinesisLoggingIntegrationDto struct {
	LoggingIntegrationDto
	StreamName string `json:"streamName"`
	Region     string `json:"region"`
}

type AwsS3CrossAccountRoleLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BucketName string `json:"bucketName"`
	Region     string `json:"region"`
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type AwsS3IamLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BucketName string `json:"bucketName"`
	Region     string `json:"region"`
}

type AzureOtelLoggingIntegrationDto struct {
	LoggingIntegrationDto
	EndpointUrl  string `json:"endpointUrl"`
	ResourceName string `json:"resourceName"`
}

type DataDogLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Site        string `json:"site"`
	ServiceName string `json:"serviceName"`
	Environment string `json:"environment"`
}

type ElasticSearchLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Uri   string `json:"uri"`
	Index string `json:"index"`
}

type InternalKafkaLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BootstrapServers string  `json:"bootstrapServers"`
	Topic            string  `json:"topic"`
	SecurityProtocol *string `json:"securityProtocol,omitempty"`
}

type KafkaLoggingIntegrationDto struct {
	LoggingIntegrationDto
	BootstrapServers string  `json:"bootstrapServers"`
	Topic            string  `json:"topic"`
	SecurityProtocol *string `json:"securityProtocol,omitempty"`
}

type KibanaLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Uri     string  `json:"uri"`
	SpaceId *string `json:"spaceId,omitempty"`
}

type LocalFileLoggingIntegrationDto struct {
	LoggingIntegrationDto
	RootPath *string `json:"rootPath,omitempty"`
}

type MongoDbLoggingIntegrationDto struct {
	LoggingIntegrationDto
	DatabaseName *string `json:"databaseName,omitempty"`
}

type NewRelicLoggingIntegrationDto struct {
	LoggingIntegrationDto
	Region      string `json:"region"`
	ServiceName string `json:"serviceName"`
}

type PrometheusLoggingIntegrationDto struct {
	LoggingIntegrationDto
	EndpointUrl string  `json:"endpointUrl"`
	JobName     *string `json:"jobName,omitempty"`
}

type SplunkLoggingIntegrationDto struct {
	LoggingIntegrationDto
	HecEndpointUrl string `json:"hecEndpointUrl"`
	Index          string `json:"index"`
}

type FilesIntegrationDto struct {
	IntegrationDto
	Provider FileProvider `json:"provider,omitempty"`
}

type AppleICloudFilesIntegrationDto struct {
	FilesIntegrationDto
	ContainerIdentifier string  `json:"containerIdentifier"`
	RelativePath        *string `json:"relativePath,omitempty"`
}

type AwsS3CrossAccountRoleFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string `json:"bucketName"`
	Region     string `json:"region"`
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type AwsS3IamFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string `json:"bucketName"`
	Region     string `json:"region"`
}

type AzureBlobFilesIntegrationDto struct {
	FilesIntegrationDto
	BlobName string `json:"blobName"`
}

type DropBoxFilesIntegrationDto struct {
	FilesIntegrationDto
	RootPath *string `json:"rootPath,omitempty"`
}

type FtpFilesIntegrationDto struct {
	FilesIntegrationDto
	Host     string  `json:"host"`
	Port     int     `json:"port,omitempty"`
	RootPath *string `json:"rootPath,omitempty"`
	UseSsl   bool    `json:"useSsl,omitempty"`
}

type GoogleCloudFilesIntegrationDto struct {
	FilesIntegrationDto
	BucketName string `json:"bucketName"`
}

type GoogleDriveFilesIntegrationDto struct {
	FilesIntegrationDto
	RootFolderId *string `json:"rootFolderId,omitempty"`
}

type LocalFilesIntegrationDto struct {
	FilesIntegrationDto
	RootPath *string `json:"rootPath,omitempty"`
}

type DatabaseIntegrationDto struct {
	IntegrationDto
	Provider DatabaseProvider `json:"provider,omitempty"`
}

type MongoDbConnectionStringIntegrationDto struct {
	DatabaseIntegrationDto
	DatabaseName *string `json:"databaseName,omitempty"`
}

type IntegrationStatus string

const (
	IntegrationStatusUnknown        IntegrationStatus = "Unknown"
	IntegrationStatusPending                          = "Pending"
	IntegrationStatusProvisioning                     = "Provisioning"
	IntegrationStatusActive                           = "Active"
	IntegrationStatusFailed                           = "Failed"
	IntegrationStatusDeprovisioning                   = "Deprovisioning"
)

type MongoDbAtlasFlexManagedIntegrationDto struct {
	DatabaseIntegrationDto
	DatabaseName     *string           `json:"databaseName,omitempty"`
	NorbixRegionCode string            `json:"norbixRegionCode"`
	FlexTierCode     string            `json:"flexTierCode"`
	Status           IntegrationStatus `json:"status,omitempty"`
	AtlasProjectId   *string           `json:"atlasProjectId,omitempty"`
	AtlasClusterName *string           `json:"atlasClusterName,omitempty"`
	FailureReason    *string           `json:"failureReason,omitempty"`
}

type SmsProvider string

const (
	SmsProviderTwilio   SmsProvider = "Twilio"
	SmsProviderVonage               = "Vonage"
	SmsProviderPlivo                = "Plivo"
	SmsProviderTelnyx               = "Telnyx"
	SmsProviderBird                 = "Bird"
	SmsProviderTelesign             = "Telesign"
	SmsProviderSinch                = "Sinch"
	SmsProviderFake                 = "Fake"
)

type SmsIntegrationDto struct {
	IntegrationDto
	Provider SmsProvider `json:"provider,omitempty"`
}

type BirdSmsIntegrationDto struct {
	SmsIntegrationDto
	Originator string `json:"originator"`
	Region     string `json:"region"`
}

type PlivoSmsIntegrationDto struct {
	SmsIntegrationDto
	AuthId          string `json:"authId"`
	FromPhoneNumber string `json:"fromPhoneNumber"`
}

type SinchSmsIntegrationDto struct {
	SmsIntegrationDto
	ServicePlanId   string `json:"servicePlanId"`
	FromPhoneNumber string `json:"fromPhoneNumber"`
}

type TelesignSmsIntegrationDto struct {
	SmsIntegrationDto
	CustomerId string `json:"customerId"`
	FromSender string `json:"fromSender"`
}

type TelnyxSmsIntegrationDto struct {
	SmsIntegrationDto
	MessagingProfileId string `json:"messagingProfileId"`
	FromPhoneNumber    string `json:"fromPhoneNumber"`
}

type TwilioSmsIntegrationDto struct {
	SmsIntegrationDto
	AccountSid      string `json:"accountSid"`
	FromPhoneNumber string `json:"fromPhoneNumber"`
}

type VonageSmsIntegrationDto struct {
	SmsIntegrationDto
	ApiKey     string `json:"apiKey"`
	FromSender string `json:"fromSender"`
}

type PushIntegrationDto struct {
	IntegrationDto
	Provider PushProvider `json:"provider,omitempty"`
}

type AndroidFirebasePushIntegrationDto struct {
	PushIntegrationDto
	ProjectId   string `json:"projectId"`
	ClientEmail string `json:"clientEmail"`
}

type AppleApnsPushIntegrationDto struct {
	PushIntegrationDto
	TeamId      string `json:"teamId"`
	AppBundleId string `json:"appBundleId"`
}

type ChromePluginPushIntegrationDto struct {
	PushIntegrationDto
	ExtensionId    string  `json:"extensionId"`
	VapidPublicKey string  `json:"vapidPublicKey"`
	Subject        *string `json:"subject,omitempty"`
}

type ChromeWebPushIntegrationDto struct {
	PushIntegrationDto
	VapidPublicKey string  `json:"vapidPublicKey"`
	Subject        *string `json:"subject,omitempty"`
}

type EdgeWebPushIntegrationDto struct {
	PushIntegrationDto
	VapidPublicKey string  `json:"vapidPublicKey"`
	Subject        *string `json:"subject,omitempty"`
}

type FirefoxWebPushIntegrationDto struct {
	PushIntegrationDto
	VapidPublicKey string  `json:"vapidPublicKey"`
	Subject        *string `json:"subject,omitempty"`
}

type SafariPushIntegrationDto struct {
	PushIntegrationDto
	WebsitePushId string `json:"websitePushId"`
}

type EmailIntegrationDto struct {
	IntegrationDto
	Provider        EmailProvider `json:"provider,omitempty"`
	EmailAddress    string        `json:"emailAddress"`
	EmailSenderName *string       `json:"emailSenderName,omitempty"`
}

type AwsSesEmailIntegrationDto struct {
	EmailIntegrationDto
	Region               string  `json:"region"`
	IdentityArn          string  `json:"identityArn"`
	ConfigurationSetName *string `json:"configurationSetName,omitempty"`
}

type AwsCrossAccountRoleEmailIntegrationDto struct {
	AwsSesEmailIntegrationDto
	RoleArn    string `json:"roleArn"`
	ExternalId string `json:"externalId"`
}

type AwsIamEmailIntegrationDto struct {
	AwsSesEmailIntegrationDto
}

type MailGunEmailIntegrationDto struct {
	EmailIntegrationDto
	Domain string        `json:"domain"`
	Region MailGunRegion `json:"region,omitempty"`
}

type SendGridEmailIntegrationDto struct {
	EmailIntegrationDto
}

type SmtpEmailIntegrationDto struct {
	EmailIntegrationDto
	HostName string `json:"hostName"`
	Port     int    `json:"port,omitempty"`
}

// @DataContract
type WebhookDestinationDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	DestinationName string `json:"destinationName"`
	// @DataMember
	EndpointUrl string `json:"endpointUrl"`
	// @DataMember
	SelectedEvents IReadOnlyList `json:"selectedEvents"`
	// @DataMember
	ExtraHeaders *IReadOnlyDictionary `json:"extraHeaders,omitempty"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
}

type WebhookIntegrationDto struct {
	IntegrationDto
	IsConfigured bool                 `json:"isConfigured,omitempty"`
	Destinations IReadOnlyList        `json:"destinations"`
	ExtraHeaders *IReadOnlyDictionary `json:"extraHeaders,omitempty"`
}

type SchedulerTaskType string

const (
	SchedulerTaskTypeEmailCampaign      SchedulerTaskType = "EmailCampaign"
	SchedulerTaskTypePushCampaign                         = "PushCampaign"
	SchedulerTaskTypeSmsCampaign                          = "SmsCampaign"
	SchedulerTaskTypeCodeFunctionalCall                   = "CodeFunctionalCall"
	SchedulerTaskTypeWebhookCall                          = "WebhookCall"
)

// @DataContract
type SchedulerTaskDto struct {
	// @DataMember
	ProjectId string `json:"projectId"`
	// @DataMember
	TaskId string `json:"taskId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Cron string `json:"cron"`
	// @DataMember
	Type SchedulerTaskType `json:"type,omitempty"`
	// @DataMember
	PayloadJson string `json:"payloadJson"`
	// @DataMember
	InitiatorId string `json:"initiatorId"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	StopOnError bool `json:"stopOnError,omitempty"`
	// @DataMember
	CreatedAtUnix *int64 `json:"createdAtUnix,omitempty"`
	// @DataMember
	UpdatedAtUnix *int64 `json:"updatedAtUnix,omitempty"`
}

type MongoDbAggregateDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	SchemaViewId string `json:"schemaViewId"`
	// @DataMember
	Pipeline string `json:"pipeline"`
}

type MarketplaceTransport string

const (
	MarketplaceTransportMcp      MarketplaceTransport = "Mcp"
	MarketplaceTransportRest                          = "Rest"
	MarketplaceTransportCode                          = "Code"
	MarketplaceTransportInternal                      = "Internal"
	MarketplaceTransportSdk                           = "Sdk"
)

type MarketplaceCategory string

const (
	MarketplaceCategoryOther         MarketplaceCategory = "Other"
	MarketplaceCategoryCrm                               = "Crm"
	MarketplaceCategoryErp                               = "Erp"
	MarketplaceCategoryMarketing                         = "Marketing"
	MarketplaceCategoryCommunication                     = "Communication"
	MarketplaceCategoryProductivity                      = "Productivity"
	MarketplaceCategoryStorage                           = "Storage"
	MarketplaceCategoryAnalytics                         = "Analytics"
	MarketplaceCategoryIdentity                          = "Identity"
	MarketplaceCategoryPayments                          = "Payments"
	MarketplaceCategoryDevTools                          = "DevTools"
	MarketplaceCategoryAi                                = "Ai"
	MarketplaceCategoryFiles                             = "Files"
	MarketplaceCategoryDatabase                          = "Database"
	MarketplaceCategoryCalendar                          = "Calendar"
)

// @DataContract
type MarketplaceIntegrationDto struct {
	IntegrationDto
	// @DataMember
	ListingViewId string `json:"listingViewId"`
	// @DataMember
	Transport MarketplaceTransport `json:"transport,omitempty"`
	// @DataMember
	Vendor string `json:"vendor"`
	// @DataMember
	Category MarketplaceCategory `json:"category,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Config IReadOnlyDictionary `json:"config"`
}

type MarketplaceTokenResolverKind string

const (
	MarketplaceTokenResolverKindStatic            MarketplaceTokenResolverKind = "Static"
	MarketplaceTokenResolverKindRequest                                        = "Request"
	MarketplaceTokenResolverKindProject                                        = "Project"
	MarketplaceTokenResolverKindInitiator                                      = "Initiator"
	MarketplaceTokenResolverKindCustom                                         = "Custom"
	MarketplaceTokenResolverKindIntegrationConfig                              = "IntegrationConfig"
	MarketplaceTokenResolverKindIntegrationSecret                              = "IntegrationSecret"
)

type MarketplaceSecretValueFormat string

const (
	MarketplaceSecretValueFormatRaw      MarketplaceSecretValueFormat = "Raw"
	MarketplaceSecretValueFormatBearer                                = "Bearer"
	MarketplaceSecretValueFormatBasic                                 = "Basic"
	MarketplaceSecretValueFormatPrefixed                              = "Prefixed"
)

// @DataContract
type MarketplaceTokenMappingDto struct {
	// @DataMember
	Token string `json:"token"`
	// @DataMember
	Resolver MarketplaceTokenResolverKind `json:"resolver,omitempty"`
	// @DataMember
	Value *string `json:"value,omitempty"`
	// @DataMember
	SecretKeys []string `json:"secretKeys,omitempty"`
	// @DataMember
	Format MarketplaceSecretValueFormat `json:"format,omitempty"`
}

// @DataContract
type MarketplaceFunctionDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	IntegrationViewId string `json:"integrationViewId"`
	// @DataMember
	FunctionKey string `json:"functionKey"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	RequestTemplate string `json:"requestTemplate"`
	// @DataMember
	MappedTokens []MarketplaceTokenMappingDto `json:"mappedTokens"`
}

type MarketplaceFieldType string

const (
	MarketplaceFieldTypeString        MarketplaceFieldType = "String"
	MarketplaceFieldTypeNumber                             = "Number"
	MarketplaceFieldTypeBoolean                            = "Boolean"
	MarketplaceFieldTypeUrl                                = "Url"
	MarketplaceFieldTypeEmail                              = "Email"
	MarketplaceFieldTypeJson                               = "Json"
	MarketplaceFieldTypeMultilineText                      = "MultilineText"
)

// @DataContract
type MarketplaceFieldDefinitionDto struct {
	// @DataMember
	Key string `json:"key"`
	// @DataMember
	Label string `json:"label"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
	// @DataMember
	Type MarketplaceFieldType `json:"type,omitempty"`
	// @DataMember
	IsRequired bool `json:"isRequired,omitempty"`
	// @DataMember
	DefaultValue *string `json:"defaultValue,omitempty"`
	// @DataMember
	Placeholder *string `json:"placeholder,omitempty"`
	// @DataMember
	ValidationPattern *string `json:"validationPattern,omitempty"`
	// @DataMember
	AllowedValues *IReadOnlyList `json:"allowedValues,omitempty"`
}

// @DataContract
type MarketplaceFunctionParameterDto struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Type string `json:"type"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	IsRequired bool `json:"isRequired,omitempty"`
	// @DataMember
	DefaultValue *string `json:"defaultValue,omitempty"`
}

type MarketplaceParameterLocation string

const (
	MarketplaceParameterLocationBody   MarketplaceParameterLocation = "Body"
	MarketplaceParameterLocationHeader                              = "Header"
	MarketplaceParameterLocationQuery                               = "Query"
	MarketplaceParameterLocationPath                                = "Path"
)

// @DataContract
type MarketplaceParameterSpecDto struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Location MarketplaceParameterLocation `json:"location,omitempty"`
	// @DataMember
	ValueTemplate *string `json:"valueTemplate,omitempty"`
	// @DataMember
	Label *string `json:"label,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
	// @DataMember
	Type *string `json:"type,omitempty"`
	// @DataMember
	IsRequired bool `json:"isRequired,omitempty"`
}

// @DataContract
type MarketplaceHttpRequestSpecDto struct {
	// @DataMember
	Method string `json:"method"`
	// @DataMember
	PathTemplate string `json:"pathTemplate"`
	// @DataMember
	Parameters IReadOnlyList `json:"parameters"`
	// @DataMember
	ContentType *string `json:"contentType,omitempty"`
}

// @DataContract
type MarketplaceFunctionDefinitionDto struct {
	// @DataMember
	DefinitionId *string `json:"definitionId,omitempty"`
	// @DataMember
	FunctionKey string `json:"functionKey"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Group *string `json:"group,omitempty"`
	// @DataMember
	Parameters IReadOnlyList `json:"parameters"`
	// @DataMember
	RequestSchema *string `json:"requestSchema,omitempty"`
	// @DataMember
	RequestTemplate *string `json:"requestTemplate,omitempty"`
	// @DataMember
	Request *MarketplaceHttpRequestSpecDto `json:"request,omitempty"`
	// @DataMember
	DefaultTokenMappings []MarketplaceTokenMappingDto `json:"defaultTokenMappings"`
}

// @DataContract
type MarketplaceListingDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Slug string `json:"slug"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	Vendor string `json:"vendor"`
	// @DataMember
	Category MarketplaceCategory `json:"category,omitempty"`
	// @DataMember
	Transport MarketplaceTransport `json:"transport,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	IconUrl *string `json:"iconUrl,omitempty"`
	// @DataMember
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
	// @DataMember
	IsOfficial bool `json:"isOfficial,omitempty"`
	// @DataMember
	Tags IReadOnlyList `json:"tags"`
	// @DataMember
	SpecVersion int `json:"specVersion,omitempty"`
	// @DataMember
	ConfigFields IReadOnlyList `json:"configFields"`
	// @DataMember
	SecretFields IReadOnlyList `json:"secretFields"`
	// @DataMember
	Functions IReadOnlyList `json:"functions"`
}

type AdminPortalModuleDto struct {
	Key         string `json:"key"`
	DisplayName string `json:"displayName"`
	Enabled     bool   `json:"enabled,omitempty"`
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

type AiChatEntryAttachmentWireDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type UserMessageEntryWireDto struct {
	AiChatEntryWireDto
	Kind        string                         `json:"kind"`
	Text        string                         `json:"text"`
	Attachments []AiChatEntryAttachmentWireDto `json:"attachments"`
}

type AiChatEntrySourceWireDto struct {
	Kind          string  `json:"kind"`
	RequirementId *string `json:"requirementId,omitempty"`
	SessionId     *string `json:"sessionId,omitempty"`
	EntryId       *string `json:"entryId,omitempty"`
	EntrySeq      *int64  `json:"entrySeq,omitempty"`
	ArtifactId    *string `json:"artifactId,omitempty"`
	Label         *string `json:"label,omitempty"`
	Step          *int    `json:"step,omitempty"`
}

type AssistantTextEntryWireDto struct {
	AiChatEntryWireDto
	Kind        string                     `json:"kind"`
	Text        string                     `json:"text"`
	IsStreaming bool                       `json:"isStreaming,omitempty"`
	Sources     []AiChatEntrySourceWireDto `json:"sources"`
}

type AiChatQuestionOptionWireDto struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type AiChatQuestionWireDto struct {
	Id            string                        `json:"id"`
	Text          string                        `json:"text"`
	Options       []AiChatQuestionOptionWireDto `json:"options"`
	Default       *string                       `json:"default,omitempty"`
	AllowFreeText bool                          `json:"allowFreeText,omitempty"`
}

type AiChatGateResultWireDto struct {
	Class                  string   `json:"class"`
	Reason                 string   `json:"reason"`
	AffectedRequirementIds []string `json:"affectedRequirementIds"`
}

type AssistantQuestionEntryWireDto struct {
	AiChatEntryWireDto
	Kind      string                   `json:"kind"`
	Questions []AiChatQuestionWireDto  `json:"questions"`
	Status    string                   `json:"status"`
	Scope     string                   `json:"scope"`
	Gate      *AiChatGateResultWireDto `json:"gate,omitempty"`
}

type UserAnswerEntryWireDto struct {
	AiChatEntryWireDto
	Kind    string            `json:"kind"`
	Answers map[string]string `json:"answers"`
}

type AiChatPlanStepInputsWireDto struct {
	Artifacts    []string `json:"artifacts"`
	Requirements []string `json:"requirements"`
}

type AiChatPlanStepDoneCheckWireDto struct {
	Check    string `json:"check"`
	ArgsJson string `json:"argsJson"`
}

type AiChatPlanStepLoopWireDto struct {
	MaxIterations *int `json:"maxIterations,omitempty"`
	MaxToolCalls  int  `json:"maxToolCalls,omitempty"`
}

type AiChatPlanStepWireDto struct {
	N          int                              `json:"n,omitempty"`
	Tool       string                           `json:"tool"`
	Title      string                           `json:"title"`
	Goal       *string                          `json:"goal,omitempty"`
	Inputs     AiChatPlanStepInputsWireDto      `json:"inputs"`
	DependsOn  []int                            `json:"dependsOn"`
	Replaces   *int                             `json:"replaces,omitempty"`
	Done       []AiChatPlanStepDoneCheckWireDto `json:"done"`
	Loop       AiChatPlanStepLoopWireDto        `json:"loop"`
	Difficulty *int                             `json:"difficulty,omitempty"`
}

type PlanEntryWireDto struct {
	AiChatEntryWireDto
	Kind             string                   `json:"kind"`
	Goal             string                   `json:"goal"`
	Steps            []AiChatPlanStepWireDto  `json:"steps"`
	Status           string                   `json:"status"`
	Gate             *AiChatGateResultWireDto `json:"gate,omitempty"`
	Difficulty       *int                     `json:"difficulty,omitempty"`
	DifficultyReason *string                  `json:"difficultyReason,omitempty"`
	DeltaOf          *string                  `json:"deltaOf,omitempty"`
}

type UserDecisionEntryWireDto struct {
	AiChatEntryWireDto
	Kind     string  `json:"kind"`
	Decision string  `json:"decision"`
	Comment  *string `json:"comment,omitempty"`
}

type AiChatStepLogLineWireDto struct {
	Seq    int     `json:"seq,omitempty"`
	Tool   string  `json:"tool"`
	Agent  *string `json:"agent,omitempty"`
	Status string  `json:"status"`
	Detail *string `json:"detail,omitempty"`
}

type RunStepEntryWireDto struct {
	AiChatEntryWireDto
	Kind          string                     `json:"kind"`
	N             int                        `json:"n,omitempty"`
	Tool          string                     `json:"tool"`
	Title         string                     `json:"title"`
	Status        string                     `json:"status"`
	ResultSummary *string                    `json:"resultSummary,omitempty"`
	Error         *string                    `json:"error,omitempty"`
	Log           []AiChatStepLogLineWireDto `json:"log"`
}

type ActionPendingEntryWireDto struct {
	AiChatEntryWireDto
	Kind          string `json:"kind"`
	Tool          string `json:"tool"`
	ArgumentsJson string `json:"argumentsJson"`
	Status        string `json:"status"`
}

type NoticeEntryWireDto struct {
	AiChatEntryWireDto
	Kind  string `json:"kind"`
	Text  string `json:"text"`
	Level string `json:"level"`
}

type ConversationSnapshotEntryWireDto struct {
	AiChatEntryWireDto
	Kind          string `json:"kind"`
	SnapshotId    string `json:"snapshotId"`
	CoversUpToSeq int64  `json:"coversUpToSeq,omitempty"`
}

type ICultureBasedRequest struct {
	CultureCode *string `json:"cultureCode,omitempty"`
}

type IVersionBasedRequest struct {
	Version string `json:"version"`
}

type IHasCorrelationIdRequest struct {
	CorrelationId *string `json:"correlationId,omitempty"`
}

type SubscriptionType string

const (
	SubscriptionTypeManagedService SubscriptionType = "ManagedService"
	SubscriptionTypeLicense                         = "License"
)

type IHasAccountId struct {
	AccountId string `json:"accountId"`
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

type IHasProjectId struct {
	ProjectId string `json:"projectId"`
}

type IHasEnv struct {
	Env *string `json:"env,omitempty"`
}

// @DataContract
type TagDescriptionDto struct {
	// @DataMember
	Title string `json:"title"`
	// @DataMember
	Description *string `json:"description,omitempty"`
}

// @DataContract
type TagTranslationDto struct {
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	Content TagDescriptionDto `json:"content"`
}

// @DataContract
type TagDefinitionBaseDto struct {
	// @DataMember
	Tag string `json:"tag"`
	// @DataMember
	Translations []TagTranslationDto `json:"translations"`
}

// @DataContract
type GroupDefinitionDto struct {
	TagDefinitionBaseDto
}

type DeliveryChannel string

const (
	DeliveryChannelEmail        DeliveryChannel = "Email"
	DeliveryChannelPush                         = "Push"
	DeliveryChannelSms                          = "Sms"
	DeliveryChannelWebPush                      = "WebPush"
	DeliveryChannelInApp                        = "InApp"
	DeliveryChannelChatBot                      = "ChatBot"
	DeliveryChannelChatPlatform                 = "ChatPlatform"
)

// @DataContract
type TagDefinitionDto struct {
	TagDefinitionBaseDto
	// @DataMember
	DefaultDelivery map[DeliveryChannel]bool `json:"defaultDelivery"`
}

type EmailAddress struct {
	Address string `json:"address"`
}

type AggregateId struct {
	Value string `json:"value,omitempty"`
}

type AccountId struct {
	AggregateId
}

type UtcDateTime struct {
}

type TimeUnit string

const (
	TimeUnitTicks        TimeUnit = "Ticks"
	TimeUnitMilliseconds          = "Milliseconds"
	TimeUnitSeconds               = "Seconds"
	TimeUnitMinutes               = "Minutes"
	TimeUnitHours                 = "Hours"
)

type ExpirationToken struct {
	Items int64    `json:"items,omitempty"`
	Unit  TimeUnit `json:"unit,omitempty"`
	Value int64    `json:"value,omitempty"`
}

type CodeMashSubscriptionId struct {
	AggregateId
}

type ProjectId struct {
	AggregateId
}

type IntegrationId struct {
	AggregateId
}

type ResourceRefKind string

const (
	ResourceRefKindContact         ResourceRefKind = "Contact"
	ResourceRefKindDocument                        = "Document"
	ResourceRefKindFile                            = "File"
	ResourceRefKindPaymentCustomer                 = "PaymentCustomer"
	ResourceRefKindOrder                           = "Order"
	ResourceRefKindPayment                         = "Payment"
	ResourceRefKindProduct                         = "Product"
	ResourceRefKindIntegration                     = "Integration"
)

type ResourceRef struct {
	ProjectId     ProjectId       `json:"projectId"`
	IntegrationId *IntegrationId  `json:"integrationId,omitempty"`
	Kind          ResourceRefKind `json:"kind,omitempty"`
}

type ResourceSource string

const (
	ResourceSourceNorbix       ResourceSource = "Norbix"
	ResourceSourceStripe                      = "Stripe"
	ResourceSourceShopify                     = "Shopify"
	ResourceSourcePayPal                      = "PayPal"
	ResourceSourceAdyen                       = "Adyen"
	ResourceSourceMollie                      = "Mollie"
	ResourceSourcePaddle                      = "Paddle"
	ResourceSourceLemonSqueezy                = "LemonSqueezy"
	ResourceSourceAppleInApp                  = "AppleInApp"
	ResourceSourceGoogleInApp                 = "GoogleInApp"
	ResourceSourceAuthorizeNet                = "AuthorizeNet"
	ResourceSourceBraintree                   = "Braintree"
	ResourceSourceCheckOutCom                 = "CheckOutCom"
	ResourceSourceWooCommerce                 = "WooCommerce"
	ResourceSourceMagento                     = "Magento"
	ResourceSourceWorldpay                    = "Worldpay"
)

type PaymentCustomerRef struct {
	ResourceRef
	Kind       ResourceRefKind `json:"kind,omitempty"`
	Source     ResourceSource  `json:"source,omitempty"`
	ExternalId string          `json:"externalId"`
}

type Quantity struct {
	Value int `json:"value,omitempty"`
}

type CodeMashManagedServiceSubscription struct {
	SubscriptionId     CodeMashSubscriptionId `json:"subscriptionId"`
	PaymentCustomerRef PaymentCustomerRef     `json:"paymentCustomerRef"`
	RefSubscriptionId  string                 `json:"refSubscriptionId"`
	IssuedOn           UtcDateTime            `json:"issuedOn"`
	WillExpireOn       UtcDateTime            `json:"willExpireOn"`
	ProjectCap         Quantity               `json:"projectCap"`
	IsTrial            bool                   `json:"isTrial,omitempty"`
}

type DomainUrl struct {
	Value string `json:"value"`
}

type CodeMashLicense struct {
	CodeMashManagedServiceSubscription
	Domain       DomainUrl `json:"domain"`
	AccountId    AccountId `json:"accountId"`
	IsEnterprise bool      `json:"isEnterprise,omitempty"`
}

// @DataContract
type ProjectName struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	UniqueName string `json:"uniqueName"`
}

type NorbixRegion struct {
	Code string `json:"code"`
}

type Continent string

const (
	ContinentAfrica       Continent = "Africa"
	ContinentAntarctica             = "Antarctica"
	ContinentAsia                   = "Asia"
	ContinentEurope                 = "Europe"
	ContinentNorthAmerica           = "NorthAmerica"
	ContinentOceania                = "Oceania"
	ContinentSouthAmerica           = "SouthAmerica"
)

// @DataContract
type ProjectRegion struct {
	// @DataMember
	Region NorbixRegion `json:"region"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Continent *Continent `json:"continent,omitempty"`
}

type ProjectLegalDocuments struct {
	TermsMarkdown   *string `json:"termsMarkdown,omitempty"`
	PrivacyMarkdown *string `json:"privacyMarkdown,omitempty"`
}

type AuthId struct {
	Value string `json:"value,omitempty"`
}

type Language struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type FileResourceId struct {
	Value string `json:"value,omitempty"`
}

type FileChecksum struct {
	Algorithm string `json:"algorithm"`
	Hash      string `json:"hash"`
}

// @DataContract
type FileResource struct {
	// @DataMember
	Id FileResourceId `json:"id"`
	// @DataMember
	OriginalFileName string `json:"originalFileName"`
	// @DataMember
	Extension string `json:"extension"`
	// @DataMember
	SizeBytes *int64 `json:"sizeBytes,omitempty"`
	// @DataMember
	Checksum *FileChecksum `json:"checksum,omitempty"`
	// @DataMember
	StoredFileName string `json:"storedFileName"`
}

// @DataContract
type FileResourceRef struct {
	// @DataMember(Order=1)
	Resource FileResource `json:"resource"`
	// @DataMember(Order=2)
	IntegrationId IntegrationId `json:"integrationId"`
	// @DataMember(Order=3)
	Provider FileProvider `json:"provider,omitempty"`
	// @DataMember(Order=4)
	Path string `json:"path"`
}

type ProjectLogo struct {
	FileResource FileResourceRef `json:"fileResource"`
	PublicUrl    string          `json:"publicUrl"`
}

type ProjectIcon struct {
	FileResource FileResourceRef `json:"fileResource"`
	PublicUrl    string          `json:"publicUrl"`
}

// @DataContract
type BrandColor struct {
	// @DataMember
	Value string `json:"value"`
}

type Tag struct {
}

// @DataContract
type GroupTags struct {
	// @DataMember
	Group Tag `json:"group"`
	// @DataMember
	Tags []Tag `json:"tags"`
}

// @DataContract
type ProjectCommunicationChannel struct {
	// @DataMember
	Channel CommunicationChannel `json:"channel,omitempty"`
	// @DataMember
	Groups []GroupTags `json:"groups"`
}

type TagDescription struct {
	DisplayName DisplayName `json:"displayName"`
	Description *string     `json:"description,omitempty"`
}

// @DataContract
type MessageTranslation struct {
}

type TagTranslation struct {
	MessageTranslation
}

type BaseTagDefinition struct {
	Tag          Tag              `json:"tag"`
	Translations []TagTranslation `json:"translations"`
}

type GroupDefinition struct {
	BaseTagDefinition
}

type TagDefinition struct {
	BaseTagDefinition
	DefaultDelivery map[DeliveryChannel]bool `json:"defaultDelivery"`
}

// @DataContract
type ProjectCommunication struct {
	// @DataMember
	Channels []ProjectCommunicationChannel `json:"channels"`
	// @DataMember
	Groups []GroupDefinition `json:"groups"`
	// @DataMember
	Tags []TagDefinition `json:"tags"`
}

// @DataContract
type TimeZone struct {
	// @DataMember
	ZoneId string `json:"zoneId"`
}

type PolicyId struct {
	Template           string `json:"template,omitempty"`
	TenancyScopeViewId string `json:"tenancyScopeViewId"`
	ViewId             string `json:"viewId"`
	IsSystem           bool   `json:"isSystem,omitempty"`
}

type PermissionEffect string

const (
	PermissionEffectAllow PermissionEffect = "Allow"
	PermissionEffectDeny                   = "Deny"
)

// @Flags()
type ApplicationModule int

const (
	ApplicationModuleAccount      ApplicationModule = 0
	ApplicationModuleMembership   ApplicationModule = 1
	ApplicationModuleDatabase     ApplicationModule = 2
	ApplicationModuleFiles        ApplicationModule = 4
	ApplicationModuleCode         ApplicationModule = 8
	ApplicationModuleEmail        ApplicationModule = 16
	ApplicationModulePush         ApplicationModule = 32
	ApplicationModulePayment      ApplicationModule = 64
	ApplicationModuleScheduler    ApplicationModule = 128
	ApplicationModuleLogging      ApplicationModule = 256
	ApplicationModuleServerEvents ApplicationModule = 512
	ApplicationModuleAi           ApplicationModule = 1024
	ApplicationModuleSms          ApplicationModule = 2048
	ApplicationModuleProject      ApplicationModule = 4096
	ApplicationModuleCompliance   ApplicationModule = 8192
	ApplicationModuleContacts     ApplicationModule = 16384
	ApplicationModuleMarketplace  ApplicationModule = 32768
)

type PermissionAction struct {
	Module              *ApplicationModule `json:"module,omitempty"`
	Operation           *string            `json:"operation,omitempty"`
	IsModuleWildcard    bool               `json:"isModuleWildcard,omitempty"`
	IsOperationWildcard bool               `json:"isOperationWildcard,omitempty"`
	IsConcrete          bool               `json:"isConcrete,omitempty"`
	Specificity         int                `json:"specificity,omitempty"`
}

type ResourceKind struct {
	Name string `json:"name"`
}

type ResourceIdentifier struct {
	Value string `json:"value"`
}

type ResourcePattern struct {
	Account           *AccountId          `json:"account,omitempty"`
	Project           *ProjectId          `json:"project,omitempty"`
	Module            *ApplicationModule  `json:"module,omitempty"`
	Kind              *ResourceKind       `json:"kind,omitempty"`
	Id                *ResourceIdentifier `json:"id,omitempty"`
	IsAccountWildcard bool                `json:"isAccountWildcard,omitempty"`
	IsProjectWildcard bool                `json:"isProjectWildcard,omitempty"`
	IsModuleWildcard  bool                `json:"isModuleWildcard,omitempty"`
	IsKindWildcard    bool                `json:"isKindWildcard,omitempty"`
	IsIdWildcard      bool                `json:"isIdWildcard,omitempty"`
	IsConcrete        bool                `json:"isConcrete,omitempty"`
	IsFullWildcard    bool                `json:"isFullWildcard,omitempty"`
	Specificity       int                 `json:"specificity,omitempty"`
}

type Permission struct {
	Sid       *string            `json:"sid,omitempty"`
	Effect    PermissionEffect   `json:"effect,omitempty"`
	Actions   []PermissionAction `json:"actions"`
	Resources []ResourcePattern  `json:"resources"`
}

type MembershipPolicy struct {
	Id          PolicyId     `json:"id"`
	Name        DisplayName  `json:"name"`
	Description *string      `json:"description,omitempty"`
	Permissions []Permission `json:"permissions"`
	Disabled    bool         `json:"disabled,omitempty"`
	IsSystem    bool         `json:"isSystem,omitempty"`
}

type RoleId struct {
	Template           string `json:"template,omitempty"`
	TenancyScopeViewId string `json:"tenancyScopeViewId"`
	ViewId             string `json:"viewId"`
	IsSystem           bool   `json:"isSystem,omitempty"`
}

type MembershipRole struct {
	Id               RoleId      `json:"id"`
	Name             DisplayName `json:"name"`
	Description      *string     `json:"description,omitempty"`
	AttachedPolicies []PolicyId  `json:"attachedPolicies"`
	Disabled         bool        `json:"disabled,omitempty"`
	IsSystem         bool        `json:"isSystem,omitempty"`
}

type BillingPeriod struct {
	Year            int       `json:"year,omitempty"`
	Month           int       `json:"month,omitempty"`
	StartUtc        time.Time `json:"startUtc,omitempty"`
	EndExclusiveUtc time.Time `json:"endExclusiveUtc,omitempty"`
	LastInstantUtc  time.Time `json:"lastInstantUtc,omitempty"`
}

type AtlasClusterChargeRecord struct {
	AtlasProjectId   string `json:"atlasProjectId"`
	AtlasClusterName string `json:"atlasClusterName"`
	Cents            int64  `json:"cents,omitempty"`
}

type AtlasUsageRecord struct {
	Period        BillingPeriod `json:"period"`
	TotalCents    int64         `json:"totalCents,omitempty"`
	PerCluster    IReadOnlyList `json:"perCluster"`
	RecordedAtUtc UtcDateTime   `json:"recordedAtUtc"`
}

type UsageIngestionFailureReason string

const (
	UsageIngestionFailureReasonUnknownCustomer  UsageIngestionFailureReason = "UnknownCustomer"
	UsageIngestionFailureReasonMeterNotFound                                = "MeterNotFound"
	UsageIngestionFailureReasonValidationFailed                             = "ValidationFailed"
	UsageIngestionFailureReasonImportSetFailed                              = "ImportSetFailed"
)

type UsageIngestionFailure struct {
	Reason        UsageIngestionFailureReason `json:"reason,omitempty"`
	Period        *BillingPeriod              `json:"period,omitempty"`
	StripeEventId string                      `json:"stripeEventId"`
	Message       string                      `json:"message"`
	ReportedAtUtc UtcDateTime                 `json:"reportedAtUtc"`
}

// @DataContract
type DeleteTrigger struct {
	CodeMashRequestBase
	// @DataMember
	TriggerId string `json:"triggerId"`
	// @DataMember
	TriggerType TriggerType `json:"triggerType,omitempty"`
	// @DataMember
	SchemaId *string `json:"schemaId,omitempty"`
}

// @DataContract
type DisableTrigger struct {
	CodeMashRequestBase
	// @DataMember
	TriggerId string `json:"triggerId"`
	// @DataMember
	TriggerType TriggerType `json:"triggerType,omitempty"`
	// @DataMember
	SchemaId *string `json:"schemaId,omitempty"`
}

// @DataContract
type EnableTrigger struct {
	CodeMashRequestBase
	// @DataMember
	TriggerId string `json:"triggerId"`
	// @DataMember
	TriggerType TriggerType `json:"triggerType,omitempty"`
	// @DataMember
	SchemaId *string `json:"schemaId,omitempty"`
}

type GetTrigger struct {
	CodeMashRequestBase
	Id       string  `json:"id"`
	SchemaId *string `json:"schemaId,omitempty"`
}

// @DataContract
type SaveTrigger struct {
	CodeMashRequestBase
	// @DataMember
	Trigger SaveTriggerRequest `json:"trigger"`
}

type CredentialsSettingsModeDto struct {
	Name      *string `json:"name,omitempty"`
	LogoutUrl *string `json:"logoutUrl,omitempty"`
}

type Integration struct {
	IntegrationId                    IntegrationId `json:"integrationId"`
	Env                              Env           `json:"env"`
	Capability                       string        `json:"capability,omitempty"`
	IsSystemOwned                    bool          `json:"isSystemOwned,omitempty"`
	IntegrationName                  DisplayName   `json:"integrationName"`
	IsEnabled                        bool          `json:"isEnabled,omitempty"`
	IsConfigured                     bool          `json:"isConfigured,omitempty"`
	LastIntegrationTestAtUtc         *time.Time    `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded     *bool         `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrorMessages IReadOnlyList `json:"lastIntegrationTestErrorMessages"`
	HumanDeliveryConfirmedAtUtc      *time.Time    `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	IsApprovedThatItWorks            bool          `json:"isApprovedThatItWorks,omitempty"`
}

type MembershipIntegration struct {
	Integration
	Provider MembershipProvider `json:"provider,omitempty"`
}

type TriggerId struct {
	AggregateId
}

type TriggerAction struct {
	Type          TriggerActionType `json:"type,omitempty"`
	IntegrationId *IntegrationId    `json:"integrationId,omitempty"`
}

// @DataContract
type TemplateCode struct {
}

type Trigger struct {
	TriggerId      TriggerId      `json:"triggerId"`
	Name           DisplayName    `json:"name"`
	TriggerAction  TriggerAction  `json:"triggerAction"`
	ActivationCode *TemplateCode  `json:"activationCode,omitempty"`
	Description    *string        `json:"description,omitempty"`
	IsEnabled      bool           `json:"isEnabled,omitempty"`
	Env            Env            `json:"env"`
	IntegrationId  *IntegrationId `json:"integrationId,omitempty"`
}

type MembershipTrigger struct {
	Trigger
	When MembershipTriggerType `json:"when,omitempty"`
}

type TriggerByIdEventBase struct {
	TriggerId TriggerId `json:"triggerId"`
}

type SchemaSettingsDto struct {
	// @DataMember
	SoftDelete bool `json:"softDelete,omitempty"`
	// @DataMember
	HasRecordOwner bool `json:"hasRecordOwner,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
}

type SchemaListColumnDto struct {
	// @DataMember
	Field string `json:"field"`
}

type SchemaListSortDto struct {
	// @DataMember
	Field string `json:"field"`
	// @DataMember
	Order int `json:"order,omitempty"`
}

type SchemaListSettingsDto struct {
	// @DataMember
	Columns []SchemaListColumnDto `json:"columns"`
	// @DataMember
	DefaultSort *SchemaListSortDto `json:"defaultSort,omitempty"`
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

type MongoDbAggregateId struct {
	AggregateId
}

type MongoDbAggregateQuery struct {
	Value string `json:"value"`
}

type SchemaId struct {
	AggregateId
}

type MongoDbAggregate struct {
	Id          MongoDbAggregateId    `json:"id"`
	DisplayName DisplayName           `json:"displayName"`
	Description *string               `json:"description,omitempty"`
	Query       MongoDbAggregateQuery `json:"query"`
	SchemaId    SchemaId              `json:"schemaId"`
}

type DatabaseIntegration struct {
	Integration
	Provider         DatabaseProvider  `json:"provider,omitempty"`
	Status           IntegrationStatus `json:"status,omitempty"`
	AtlasProjectId   *string           `json:"atlasProjectId,omitempty"`
	AtlasClusterName *string           `json:"atlasClusterName,omitempty"`
	FailureReason    *string           `json:"failureReason,omitempty"`
}

type ProjectStatus string

const (
	ProjectStatusActive             ProjectStatus = "Active"
	ProjectStatusProvisioning                     = "Provisioning"
	ProjectStatusProvisioningFailed               = "ProvisioningFailed"
	ProjectStatusNoDatabase                       = "NoDatabase"
	ProjectStatusDisabled                         = "Disabled"
	ProjectStatusSuspended                        = "Suspended"
	ProjectStatusRemoved                          = "Removed"
)

type SchemaName struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

type JsonSchemaFieldName struct {
	FieldName string `json:"fieldName"`
}

type JsonSchemaField struct {
	FieldName JsonSchemaFieldName `json:"fieldName"`
}

type DataSchema struct {
	RawJson string            `json:"rawJson"`
	Fields  []JsonSchemaField `json:"fields"`
}

type VisualSchema struct {
	RawJson string `json:"rawJson"`
}

type SchemaDraft struct {
	DataSchema   DataSchema   `json:"dataSchema"`
	VisualSchema VisualSchema `json:"visualSchema"`
	UpdatedAt    time.Time    `json:"updatedAt,omitempty"`
}

type SchemaVersion struct {
	Value int `json:"value,omitempty"`
}

type MetaSchemaVersion struct {
	Value int `json:"value,omitempty"`
}

type PublishedSchemaVersion struct {
	Version           SchemaVersion     `json:"version"`
	DataSchema        DataSchema        `json:"dataSchema"`
	VisualSchema      VisualSchema      `json:"visualSchema"`
	MetaSchemaVersion MetaSchemaVersion `json:"metaSchemaVersion"`
	PublishedAt       time.Time         `json:"publishedAt,omitempty"`
}

type SchemaSettings struct {
	SoftDelete     bool    `json:"softDelete,omitempty"`
	HasRecordOwner bool    `json:"hasRecordOwner,omitempty"`
	Description    *string `json:"description,omitempty"`
}

type Schema struct {
	SchemaName        SchemaName      `json:"schemaName"`
	Id                SchemaId        `json:"id"`
	Env               Env             `json:"env"`
	Draft             *SchemaDraft    `json:"draft,omitempty"`
	PublishedVersions IReadOnlyList   `json:"publishedVersions"`
	Triggers          []Trigger       `json:"triggers,omitempty"`
	Settings          *SchemaSettings `json:"settings,omitempty"`
}

type SchemaDiff struct {
	AddedFields              IReadOnlyList `json:"addedFields"`
	RemovedFields            IReadOnlyList `json:"removedFields"`
	TypeChangedFields        IReadOnlyList `json:"typeChangedFields"`
	ValidatorTightenedFields IReadOnlyList `json:"validatorTightenedFields"`
	IsEmpty                  bool          `json:"isEmpty,omitempty"`
}

type TaxonomyId struct {
	AggregateId
}

type TaxonomyName struct {
	Value string `json:"value"`
	Title string `json:"title"`
}

type RecordId struct {
	Id string `json:"id"`
}

type Taxonomy struct {
	ParentId              *TaxonomyId   `json:"parentId,omitempty"`
	Id                    TaxonomyId    `json:"id"`
	Name                  TaxonomyName  `json:"name"`
	Description           *string       `json:"description,omitempty"`
	TermsMetaVisualSchema *VisualSchema `json:"termsMetaVisualSchema,omitempty"`
	TermsMetaDataSchema   *DataSchema   `json:"termsMetaDataSchema,omitempty"`
	Dependencies          []TaxonomyId  `json:"dependencies,omitempty"`
	RecordId              *RecordId     `json:"recordId,omitempty"`
}

type SchemaTrigger struct {
	Trigger
	SchemaId      SchemaId          `json:"schemaId"`
	When          SchemaTriggerType `json:"when,omitempty"`
	Configuration *TemplateCode     `json:"configuration,omitempty"`
}

type IPasskeyMessage struct {
}

type AuthUserName struct {
	Value string `json:"value"`
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

type IpAddress struct {
	Ip string `json:"ip"`
}

type AccessInformation struct {
	Ip   *IpAddress   `json:"ip,omitempty"`
	Date *UtcDateTime `json:"date,omitempty"`
	Zone *TimeZone    `json:"zone,omitempty"`
}

type Registration struct {
	RegistrationInformation AccessInformation `json:"registrationInformation"`
}

type Login struct {
	NeedChangePasswordOnNextLogin bool               `json:"needChangePasswordOnNextLogin,omitempty"`
	LastAccessInformation         *AccessInformation `json:"lastAccessInformation,omitempty"`
}

type Phone struct {
	Value string `json:"value"`
}

// @DataContract
type FirstName struct {
	// @DataMember
	Value string `json:"value"`
}

// @DataContract
type LastName struct {
	// @DataMember
	Value string `json:"value"`
}

type MidName struct {
	Value string `json:"value"`
}

type FullName struct {
	FirstName *FirstName `json:"firstName,omitempty"`
	MidName   *MidName   `json:"midName,omitempty"`
	LastName  *LastName  `json:"lastName,omitempty"`
	Title     *string    `json:"title,omitempty"`
}

type City struct {
	Value string `json:"value"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type AddressLine struct {
	Value string `json:"value"`
}

type PostalCode struct {
	Value string `json:"value"`
}

type CountryState struct {
	Value string `json:"value"`
}

type Address struct {
	City         *City         `json:"city,omitempty"`
	Country      *Country      `json:"country,omitempty"`
	AddressLine1 *AddressLine  `json:"addressLine1,omitempty"`
	AddressLine2 *AddressLine  `json:"addressLine2,omitempty"`
	PostalCode   *PostalCode   `json:"postalCode,omitempty"`
	State        *CountryState `json:"state,omitempty"`
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

type UserMarketingPreferences struct {
	BlockAllMarketingMessages bool                      `json:"blockAllMarketingMessages,omitempty"`
	BlockedTags               map[DeliveryChannel][]Tag `json:"blockedTags,omitempty"`
	BlockReasons              []MarketingBlockReason    `json:"blockReasons,omitempty"`
}

type UserGeneralInfo struct {
	Phone                *Phone                    `json:"phone,omitempty"`
	PrimaryEmail         *EmailAddress             `json:"primaryEmail,omitempty"`
	DisplayName          *DisplayName              `json:"displayName,omitempty"`
	FirstName            *FirstName                `json:"firstName,omitempty"`
	LastName             *LastName                 `json:"lastName,omitempty"`
	FullName             *FullName                 `json:"fullName,omitempty"`
	Address              *Address                  `json:"address,omitempty"`
	Company              *string                   `json:"company,omitempty"`
	Gender               *Gender                   `json:"gender,omitempty"`
	BirthDate            *UtcDateTime              `json:"birthDate,omitempty"`
	TimeZone             *TimeZone                 `json:"timeZone,omitempty"`
	Language             *Language                 `json:"language,omitempty"`
	MarketingPreferences *UserMarketingPreferences `json:"marketingPreferences,omitempty"`
	Notes                *string                   `json:"notes,omitempty"`
	ExtraMetadata        *string                   `json:"extraMetadata,omitempty"`
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

type DeviceId struct {
	Id string `json:"id,omitempty"`
}

type DeviceType string

const (
	DeviceTypeUnknown DeviceType = "Unknown"
	DeviceTypePhone              = "Phone"
	DeviceTypeTablet             = "Tablet"
	DeviceTypeDesktop            = "Desktop"
	DeviceTypeTv                 = "Tv"
)

type PushDeviceToken struct {
	Token string `json:"token"`
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
type PushDeviceDeliveryToken struct {
	// @DataMember
	PushDeviceToken PushDeviceToken `json:"pushDeviceToken"`
	// @DataMember
	DeliveryFamily PushDeviceDeliveryFamily `json:"deliveryFamily,omitempty"`
}

// @DataContract
type PushDevice struct {
	// @DataMember
	Id DeviceId `json:"id"`
	// @DataMember
	Brand *string `json:"brand,omitempty"`
	// @DataMember
	Manufacturer *string `json:"manufacturer,omitempty"`
	// @DataMember
	ModelName *string `json:"modelName,omitempty"`
	// @DataMember
	DeviceName *string `json:"deviceName,omitempty"`
	// @DataMember
	DeviceType *DeviceType `json:"deviceType,omitempty"`
	// @DataMember
	OsName *string `json:"osName,omitempty"`
	// @DataMember
	OsVersion *string `json:"osVersion,omitempty"`
	// @DataMember
	PlatformApiLevel *int `json:"platformApiLevel,omitempty"`
	// @DataMember
	Token PushDeviceDeliveryToken `json:"token"`
}

type PushDevices []PushDevice

type UserId struct {
	Value string `json:"value,omitempty"`
}

type UserRef struct {
	ResourceRef
	Kind   ResourceRefKind `json:"kind,omitempty"`
	UserId UserId          `json:"userId"`
}

type Auth struct {
	Id           AuthId           `json:"id"`
	Roles        []RoleName       `json:"roles,omitempty"`
	Email        *EmailAddress    `json:"email,omitempty"`
	UserName     *AuthUserName    `json:"userName,omitempty"`
	Type         AuthType         `json:"type,omitempty"`
	Registration Registration     `json:"registration"`
	Login        *Login           `json:"login,omitempty"`
	GeneralInfo  *UserGeneralInfo `json:"generalInfo,omitempty"`
	Status       AuthStatus       `json:"status,omitempty"`
	CreatedOn    UtcDateTime      `json:"createdOn"`
	ModifiedOn   UtcDateTime      `json:"modifiedOn"`
	PushDevices  *PushDevices     `json:"pushDevices,omitempty"`
	Tags         []Tag            `json:"tags,omitempty"`
	UserRef      *UserRef         `json:"userRef,omitempty"`
}

type FileIntegration struct {
	Integration
	Provider FileProvider `json:"provider,omitempty"`
}

type FileTrigger struct {
	Trigger
	When            FilesTriggerType `json:"when,omitempty"`
	FileResourceRef *FileResourceRef `json:"fileResourceRef,omitempty"`
}

// @DataContract
type EmailValidationProvider string

const (
	EmailValidationProviderZeroBounce      EmailValidationProvider = "ZeroBounce"
	EmailValidationProviderNeverBounce                             = "NeverBounce"
	EmailValidationProviderBouncer                                 = "Bouncer"
	EmailValidationProviderMailgunValidate                         = "MailgunValidate"
)

type EmailValidationIntegrationRequest struct {
	IntegrationId   *string                 `json:"integrationId,omitempty"`
	Provider        EmailValidationProvider `json:"provider,omitempty"`
	IntegrationName string                  `json:"integrationName"`
	IsEnabled       bool                    `json:"isEnabled,omitempty"`
}

type SaveEmailTemplate struct {
	CodeMashRequestBase
	/** @description The display name of the email template. */
	// @ApiMember(Description="The display name of the email template.", IsRequired=true)
	TemplateName string `json:"templateName"`
	/** @description Optional free-text description of what the template is used for. */
	// @ApiMember(Description="Optional free-text description of what the template is used for.")
	Description *string `json:"description,omitempty"`
	/** @description The communication channel the template is intended for (e.g. Transactional, Marketing). */
	// @ApiMember(Description="The communication channel the template is intended for (e.g. Transactional, Marketing).")
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	/** @description Optional tags to organize/filter the template by. */
	// @ApiMember(Description="Optional tags to organize/filter the template by.")
	Tags []string `json:"tags,omitempty"`
	/** @description Optional static file attachments to send with every email using this template. */
	// @DataMember
	// @ApiMember(Description="Optional static file attachments to send with every email using this template.")
	StaticAttachments []FileResourceRefDto `json:"staticAttachments,omitempty"`
	/** @description The per-language content translations (subject/body) for this template. */
	// @ApiMember(Description="The per-language content translations (subject/body) for this template.", IsRequired=true)
	Translations []EmailMessageTranslationDto `json:"translations"`
}

// @DataContract
type TranslationDto struct {
	// @DataMember
	Language string `json:"language"`
	// @DataMember
	Content string `json:"content"`
}

type EmailFooterId struct {
	Value string `json:"value,omitempty"`
}

type EmailFooter struct {
	Id           EmailFooterId        `json:"id"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Env          Env                  `json:"env"`
}

// @DataContract
type EmailSenderName struct {
}

type EmailIntegration struct {
	Integration
	Provider        EmailProvider    `json:"provider,omitempty"`
	EmailAddress    EmailAddress     `json:"emailAddress"`
	EmailSenderName *EmailSenderName `json:"emailSenderName,omitempty"`
}

type EmailSignatureId struct {
	Value string `json:"value,omitempty"`
}

type EmailSignature struct {
	Id           EmailSignatureId     `json:"id"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Env          Env                  `json:"env"`
}

type TemplateId struct {
	Value string `json:"value,omitempty"`
}

// @DataContract
type Template struct {
	// @DataMember
	TemplateId TemplateId `json:"templateId"`
	// @DataMember
	TemplateName DisplayName `json:"templateName"`
	// @DataMember
	Translations []MessageTranslation `json:"translations"`
	// @DataMember
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	// @DataMember
	IsActive bool `json:"isActive,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Tags []Tag `json:"tags,omitempty"`
	// @DataMember
	FileIntegrationId *IntegrationId `json:"fileIntegrationId,omitempty"`
	// @DataMember
	Env Env `json:"env"`
}

// @DataContract
type EmailSubject struct {
}

// @DataContract
type EmailBody struct {
	// @DataMember
	Code TemplateCode `json:"code"`
	// @DataMember
	Structure *string `json:"structure,omitempty"`
	// @DataMember
	EmailTemplateEngine EmailTemplateEngine `json:"emailTemplateEngine,omitempty"`
}

// @DataContract
type EmailMessageContent struct {
	// @DataMember(Order=1)
	Subject EmailSubject `json:"subject"`
	// @DataMember(Order=2)
	Body EmailBody `json:"body"`
	// @DataMember(Order=3)
	StaticAttachments []FileResourceRef `json:"staticAttachments,omitempty"`
}

// @DataContract
type EmailTemplate struct {
	Template
	// @DataMember
	StaticAttachments []FileResourceRef `json:"staticAttachments,omitempty"`
}

type EmailValidationIntegration struct {
	Integration
	Provider EmailValidationProvider `json:"provider,omitempty"`
}

type CampaignId struct {
	Id string `json:"id,omitempty"`
}

type CampaignBatchId struct {
	Id string `json:"id,omitempty"`
}

type NotificationId struct {
	AggregateId
}

type CampaignStopReason string

const (
	CampaignStopReasonUserRequested  CampaignStopReason = "UserRequested"
	CampaignStopReasonModuleDisabled                    = "ModuleDisabled"
)

type EmailDeliveryEventType string

const (
	EmailDeliveryEventTypeUnknown      EmailDeliveryEventType = "Unknown"
	EmailDeliveryEventTypeDelivered                           = "Delivered"
	EmailDeliveryEventTypeOpen                                = "Open"
	EmailDeliveryEventTypeClick                               = "Click"
	EmailDeliveryEventTypeSoftBounce                          = "SoftBounce"
	EmailDeliveryEventTypeHardBounce                          = "HardBounce"
	EmailDeliveryEventTypeComplaint                           = "Complaint"
	EmailDeliveryEventTypeUnsubscribed                        = "Unsubscribed"
)

type SaveSmsTemplate struct {
	CodeMashRequestBase
	/** @description Display name for the SMS template. */
	// @ApiMember(Description="Display name for the SMS template.", IsRequired=true)
	TemplateName string `json:"templateName"`
	/** @description Optional free-text description of the template's purpose. */
	// @ApiMember(Description="Optional free-text description of the template's purpose.")
	Description *string `json:"description,omitempty"`
	/** @description Whether this template is Transactional or Marketing SMS. */
	// @ApiMember(Description="Whether this template is Transactional or Marketing SMS.", IsRequired=true)
	CommunicationChannel CommunicationChannel `json:"communicationChannel"`
	/** @description Optional tags to organize/filter templates. */
	// @ApiMember(Description="Optional tags to organize/filter templates.")
	Tags []string `json:"tags,omitempty"`
	/** @description The template's per-language translations (each with its own SMS body). */
	// @ApiMember(Description="The template's per-language translations (each with its own SMS body).", IsRequired=true)
	Translations []SmsMessageTranslationDto `json:"translations"`
}

type SmsIntegrationRequest struct {
	IntegrationId   *string     `json:"integrationId,omitempty"`
	Provider        SmsProvider `json:"provider,omitempty"`
	IntegrationName string      `json:"integrationName"`
	IsEnabled       bool        `json:"isEnabled,omitempty"`
}

type SmsIntegration struct {
	Integration
	Provider SmsProvider `json:"provider,omitempty"`
}

type SmsTitle struct {
	Value TemplateCode `json:"value"`
}

type SmsBody struct {
	Value TemplateCode `json:"value"`
}

// @DataContract
type SmsMessageContent struct {
	// @DataMember(Order=1)
	Title SmsTitle `json:"title"`
	// @DataMember(Order=2)
	Body SmsBody `json:"body"`
}

// @DataContract
type SmsTemplate struct {
	Template
}

type CodeIntegration struct {
	Integration
	Provider CodeProvider `json:"provider,omitempty"`
}

type MarketplaceIntegrationTransport string

const (
	MarketplaceIntegrationTransportMcp      MarketplaceIntegrationTransport = "Mcp"
	MarketplaceIntegrationTransportRest                                     = "Rest"
	MarketplaceIntegrationTransportCode                                     = "Code"
	MarketplaceIntegrationTransportInternal                                 = "Internal"
	MarketplaceIntegrationTransportSdk                                      = "Sdk"
)

type MarketplaceIntegrationCategory string

const (
	MarketplaceIntegrationCategoryOther         MarketplaceIntegrationCategory = "Other"
	MarketplaceIntegrationCategoryCrm                                          = "Crm"
	MarketplaceIntegrationCategoryErp                                          = "Erp"
	MarketplaceIntegrationCategoryMarketing                                    = "Marketing"
	MarketplaceIntegrationCategoryCommunication                                = "Communication"
	MarketplaceIntegrationCategoryProductivity                                 = "Productivity"
	MarketplaceIntegrationCategoryStorage                                      = "Storage"
	MarketplaceIntegrationCategoryAnalytics                                    = "Analytics"
	MarketplaceIntegrationCategoryIdentity                                     = "Identity"
	MarketplaceIntegrationCategoryPayments                                     = "Payments"
	MarketplaceIntegrationCategoryDevTools                                     = "DevTools"
	MarketplaceIntegrationCategoryAi                                           = "Ai"
	MarketplaceIntegrationCategoryFiles                                        = "Files"
	MarketplaceIntegrationCategoryDatabase                                     = "Database"
	MarketplaceIntegrationCategoryCalendar                                     = "Calendar"
)

type MarketplaceTokenResolver string

const (
	MarketplaceTokenResolverStatic            MarketplaceTokenResolver = "Static"
	MarketplaceTokenResolverRequest                                    = "Request"
	MarketplaceTokenResolverProject                                    = "Project"
	MarketplaceTokenResolverInitiator                                  = "Initiator"
	MarketplaceTokenResolverCustom                                     = "Custom"
	MarketplaceTokenResolverIntegrationConfig                          = "IntegrationConfig"
	MarketplaceTokenResolverIntegrationSecret                          = "IntegrationSecret"
)

type SecretValueFormat string

const (
	SecretValueFormatRaw      SecretValueFormat = "Raw"
	SecretValueFormatBearer                     = "Bearer"
	SecretValueFormatBasic                      = "Basic"
	SecretValueFormatPrefixed                   = "Prefixed"
)

type MarketplaceTokenMapping struct {
	Token      string                   `json:"token"`
	Resolver   MarketplaceTokenResolver `json:"resolver,omitempty"`
	Value      *string                  `json:"value,omitempty"`
	SecretKeys *IReadOnlyList           `json:"secretKeys,omitempty"`
	Format     SecretValueFormat        `json:"format,omitempty"`
}

type MarketplaceIntegration struct {
	Integration
	Capability    string                          `json:"capability,omitempty"`
	ListingViewId string                          `json:"listingViewId"`
	Transport     MarketplaceIntegrationTransport `json:"transport,omitempty"`
	Vendor        string                          `json:"vendor"`
	Category      MarketplaceIntegrationCategory  `json:"category,omitempty"`
	Description   *string                         `json:"description,omitempty"`
	Config        IReadOnlyDictionary             `json:"config"`
	TokenMappings IReadOnlyList                   `json:"tokenMappings"`
}

type MarketplaceFunctionId struct {
	Value string `json:"value,omitempty"`
}

type MarketplaceFunction struct {
	FunctionId      MarketplaceFunctionId `json:"functionId"`
	IntegrationId   IntegrationId         `json:"integrationId"`
	Env             Env                   `json:"env"`
	FunctionKey     string                `json:"functionKey"`
	DisplayName     DisplayName           `json:"displayName"`
	Description     *string               `json:"description,omitempty"`
	IsEnabled       bool                  `json:"isEnabled,omitempty"`
	RequestTemplate string                `json:"requestTemplate"`
	MappedTokens    IReadOnlyList         `json:"mappedTokens"`
	ViewId          string                `json:"viewId"`
}

type SavePushTemplate struct {
	CodeMashRequestBase
	/** @description The template's display name. */
	// @ApiMember(Description="The template's display name.", IsRequired=true)
	TemplateName string `json:"templateName"`
	/** @description Optional free-text description of the template's purpose. */
	// @ApiMember(Description="Optional free-text description of the template's purpose.")
	Description *string `json:"description,omitempty"`
	/** @description Whether the template is Transactional or Marketing. */
	// @ApiMember(Description="Whether the template is Transactional or Marketing.")
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
	/** @description Optional tags for organizing/filtering templates. */
	// @ApiMember(Description="Optional tags for organizing/filtering templates.")
	Tags []string `json:"tags,omitempty"`
	/** @description The per-locale translations (title/body/subtitle) that make up the template content. */
	// @ApiMember(Description="The per-locale translations (title/body/subtitle) that make up the template content.", IsRequired=true)
	Translations []PushMessageTranslationDto `json:"translations"`
}

type IHasAccountId struct {
	AccountId string `json:"accountId"`
}

// @DataContract
type PushDeviceDto struct {
	// @DataMember
	DeviceId *string `json:"deviceId,omitempty"`
	// @DataMember
	DeviceOs string `json:"deviceOs"`
	// @DataMember
	Token string `json:"token"`
	// @DataMember
	Brand *string `json:"brand,omitempty"`
	// @DataMember
	Manufacturer *string `json:"manufacturer,omitempty"`
	// @DataMember
	ModelName *string `json:"modelName,omitempty"`
	// @DataMember
	DeviceName *string `json:"deviceName,omitempty"`
	// @DataMember
	DeviceType *DeviceType `json:"deviceType,omitempty"`
	// @DataMember
	OsName *string `json:"osName,omitempty"`
	// @DataMember
	OsVersion *string `json:"osVersion,omitempty"`
	// @DataMember
	PlatformApiLevel *int `json:"platformApiLevel,omitempty"`
}

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

type PushIntegration struct {
	Integration
	Provider PushProvider `json:"provider,omitempty"`
}

// @DataContract
type PushTitle struct {
	// @DataMember
	Value TemplateCode `json:"value"`
}

type PushBody struct {
	Value TemplateCode `json:"value"`
}

// @DataContract
type PushMessageContent struct {
	// @DataMember(Order=1)
	Title PushTitle `json:"title"`
	// @DataMember(Order=1)
	SubTitle *PushTitle `json:"subTitle,omitempty"`
	// @DataMember(Order=2)
	Body PushBody `json:"body"`
}

// @DataContract
type PushTemplate struct {
	Template
}

type PaymentIntegration struct {
	Integration
	Provider PaymentGatewayPlatform `json:"provider,omitempty"`
}

type PaymentTrigger struct {
	Trigger
	When         PaymentTriggerType `json:"when,omitempty"`
	Integrations []IntegrationId    `json:"integrations,omitempty"`
	Events       []string           `json:"events,omitempty"`
}

type LoggingIntegration struct {
	Integration
	Provider LoggingProvider `json:"provider,omitempty"`
}

type ChatScreenContextDto struct {
	Kind   string  `json:"kind"`
	ViewId *string `json:"viewId,omitempty"`
}

type LlmIntegration struct {
	Integration
	Provider     LlmProvider `json:"provider,omitempty"`
	DefaultModel string      `json:"defaultModel"`
}

type McpIntegration struct {
	Integration
	Provider  McpProvider  `json:"provider,omitempty"`
	Transport McpTransport `json:"transport,omitempty"`
	Metadata  McpMetadata  `json:"metadata"`
}

type WebhookDestinationId struct {
	AggregateId
}

type TriggerEventName struct {
	Value string `json:"value"`
}

type WebhookDestination struct {
	DestinationId   WebhookDestinationId `json:"destinationId"`
	DestinationName DisplayName          `json:"destinationName"`
	EndpointUrl     DomainUrl            `json:"endpointUrl"`
	SelectedEvents  []TriggerEventName   `json:"selectedEvents"`
	ExtraHeaders    *IReadOnlyDictionary `json:"extraHeaders,omitempty"`
	IsEnabled       bool                 `json:"isEnabled,omitempty"`
}

type WebhookIntegration struct {
	Integration
	Capability   string               `json:"capability,omitempty"`
	Destinations []WebhookDestination `json:"destinations"`
	ExtraHeaders *IReadOnlyDictionary `json:"extraHeaders,omitempty"`
}

type SchedulerTaskRequest struct {
	Type SchedulerTaskType `json:"type,omitempty"`
}

type TaskId struct {
	AggregateId
}

type CronExpression struct {
}

type SchedulerTask struct {
	Id          TaskId            `json:"id"`
	Type        SchedulerTaskType `json:"type,omitempty"`
	Name        DisplayName       `json:"name"`
	Description *string           `json:"description,omitempty"`
	Cron        CronExpression    `json:"cron"`
	PayloadJson string            `json:"payloadJson"`
	InitiatorId AuthId            `json:"initiatorId"`
	IsEnabled   bool              `json:"isEnabled,omitempty"`
	StopOnError bool              `json:"stopOnError,omitempty"`
}

type ResourceKindDto string

const (
	ResourceKindDtoContact         ResourceKindDto = "contact"
	ResourceKindDtoDocument                        = "document"
	ResourceKindDtoFile                            = "file"
	ResourceKindDtoPaymentCustomer                 = "paymentCustomer"
	ResourceKindDtoOrder                           = "order"
	ResourceKindDtoPayment                         = "payment"
	ResourceKindDtoProduct                         = "product"
	ResourceKindDtoIntegration                     = "integration"
)

type ResourceRefDto struct {
	ProjectId     string          `json:"projectId"`
	IntegrationId *string         `json:"integrationId,omitempty"`
	Kind          ResourceKindDto `json:"kind,omitempty"`
}

type CaseResolutionFixKind string

const (
	CaseResolutionFixKindCodeFix             CaseResolutionFixKind = "CodeFix"
	CaseResolutionFixKindConfigChange                              = "ConfigChange"
	CaseResolutionFixKindCustomerInstruction                       = "CustomerInstruction"
	CaseResolutionFixKindKnownLimitation                           = "KnownLimitation"
	CaseResolutionFixKindDuplicate                                 = "Duplicate"
)

// @DataContract
type CaseResolutionDto struct {
	// @DataMember
	Problem string `json:"problem"`
	// @DataMember
	Symptoms []string `json:"symptoms"`
	// @DataMember
	RootCause string `json:"rootCause"`
	// @DataMember
	Fix CaseResolutionFixKind `json:"fix,omitempty"`
	// @DataMember
	FixDetail *string `json:"fixDetail,omitempty"`
	// @DataMember
	AffectedVersions []string `json:"affectedVersions,omitempty"`
}

type SupportCaseId struct {
	AggregateId
	ViewId string `json:"viewId"`
}

type SupportCaseKind string

const (
	SupportCaseKindQuestion       SupportCaseKind = "Question"
	SupportCaseKindBug                            = "Bug"
	SupportCaseKindIncident                       = "Incident"
	SupportCaseKindBilling                        = "Billing"
	SupportCaseKindSecurity                       = "Security"
	SupportCaseKindFeatureRequest                 = "FeatureRequest"
)

type SupportCaseSeverity string

const (
	SupportCaseSeverityS1 SupportCaseSeverity = "S1"
	SupportCaseSeverityS2                     = "S2"
	SupportCaseSeverityS3                     = "S3"
	SupportCaseSeverityS4                     = "S4"
)

type DeploymentMode string

const (
	DeploymentModeManaged    DeploymentMode = "Managed"
	DeploymentModeSelfHosted                = "SelfHosted"
	DeploymentModeEnterprise                = "Enterprise"
)

type SupportMessageAuthorKind string

const (
	SupportMessageAuthorKindCustomer SupportMessageAuthorKind = "Customer"
	SupportMessageAuthorKindStaff                             = "Staff"
	SupportMessageAuthorKindAi                                = "Ai"
	SupportMessageAuthorKindSystem                            = "System"
)

type SupportMessageRef struct {
	MessageId  string                   `json:"messageId"`
	AuthorKind SupportMessageAuthorKind `json:"authorKind,omitempty"`
	AuthorId   *string                  `json:"authorId,omitempty"`
	SentOn     UtcDateTime              `json:"sentOn"`
}

type SupportCaseStatus string

const (
	SupportCaseStatusOpen              SupportCaseStatus = "Open"
	SupportCaseStatusTriaged                             = "Triaged"
	SupportCaseStatusInProgress                          = "InProgress"
	SupportCaseStatusWaitingOnCustomer                   = "WaitingOnCustomer"
	SupportCaseStatusResolved                            = "Resolved"
	SupportCaseStatusClosed                              = "Closed"
)

type CaseResolution struct {
	Problem          string                `json:"problem"`
	Symptoms         IReadOnlyList         `json:"symptoms"`
	RootCause        string                `json:"rootCause"`
	Fix              CaseResolutionFixKind `json:"fix,omitempty"`
	FixDetail        *string               `json:"fixDetail,omitempty"`
	Module           *string               `json:"module,omitempty"`
	Kind             SupportCaseKind       `json:"kind,omitempty"`
	Severity         SupportCaseSeverity   `json:"severity,omitempty"`
	AffectedVersions IReadOnlyList         `json:"affectedVersions"`
	ResolvedBy       *string               `json:"resolvedBy,omitempty"`
}

type SupportCaseCloseReason string

const (
	SupportCaseCloseReasonManual                 SupportCaseCloseReason = "Manual"
	SupportCaseCloseReasonAutoClosedAfterResolve                        = "AutoClosedAfterResolve"
)

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

// @DataContract
type AccountOwnerDto struct {
	// @DataMember
	Email string `json:"email"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	BillingEmail *string `json:"billingEmail,omitempty"`
	// @DataMember
	OperationsEmail *string `json:"operationsEmail,omitempty"`
	// @DataMember
	SecurityEmail *string `json:"securityEmail,omitempty"`
}

// @Flags()
type AccountStatus int

const (
	AccountStatusRegistered        AccountStatus = 1
	AccountStatusPendingValidation AccountStatus = 2
	AccountStatusActive            AccountStatus = 8
	AccountStatusInActive          AccountStatus = 16
	AccountStatusBlocked           AccountStatus = 32
	AccountStatusUnregistered      AccountStatus = 64
)

// @DataContract
type AccountStatusDto struct {
	// @DataMember
	AccountId string `json:"accountId"`
	// @DataMember
	AccountIdAsGuid string `json:"accountIdAsGuid,omitempty"`
	// @DataMember
	UserId string `json:"userId"`
	// @DataMember
	LoggedInUserId string `json:"loggedInUserId"`
	// @DataMember
	LoggedInUserEmail *string `json:"loggedInUserEmail,omitempty"`
	// @DataMember
	Status AccountStatus `json:"status,omitempty"`
	// @DataMember
	ProjectCap int `json:"projectCap,omitempty"`
	// @DataMember
	Permissions []string `json:"permissions"`
	// @DataMember
	Roles []string `json:"roles"`
	// @DataMember
	AllowedProjects []string `json:"allowedProjects,omitempty"`
	// @DataMember
	TrialWasIssued bool `json:"trialWasIssued,omitempty"`
}

// @DataContract
type UsageBillingClusterChargeDto struct {
	// @DataMember
	AtlasProjectId string `json:"atlasProjectId"`
	// @DataMember
	AtlasClusterName string `json:"atlasClusterName"`
	// @DataMember
	Cents int64 `json:"cents,omitempty"`
}

// @DataContract
type UsageBillingPeriodDto struct {
	// @DataMember
	Period string `json:"period"`
	// @DataMember
	TotalCents int64 `json:"totalCents,omitempty"`
	// @DataMember
	PerCluster []UsageBillingClusterChargeDto `json:"perCluster"`
	// @DataMember
	RecordedAtUtc time.Time `json:"recordedAtUtc,omitempty"`
}

// @DataContract
type UsageBillingIngestionFailureDto struct {
	// @DataMember
	Reason string `json:"reason"`
	// @DataMember
	Period *string `json:"period,omitempty"`
	// @DataMember
	StripeEventId string `json:"stripeEventId"`
	// @DataMember
	Message string `json:"message"`
	// @DataMember
	ReportedAtUtc time.Time `json:"reportedAtUtc,omitempty"`
}

// @DataContract
type UsageBillingDto struct {
	// @DataMember
	AccountId string `json:"accountId"`
	// @DataMember
	Atlas map[string]UsageBillingPeriodDto `json:"atlas"`
	// @DataMember
	IngestionFailures []UsageBillingIngestionFailureDto `json:"ingestionFailures"`
}

// @DataContract
type PromotionItemDto struct {
	// @DataMember
	Type string `json:"type"`
	// @DataMember
	Id string `json:"id"`
}

// @DataContract
type PromotionBlockerDto struct {
	// @DataMember
	ContentType string `json:"contentType"`
	// @DataMember
	ContentId string `json:"contentId"`
	// @DataMember
	RefKind string `json:"refKind"`
	// @DataMember
	UnresolvedRef string `json:"unresolvedRef"`
}

// @DataContract
type PromotionResultDto struct {
	// @DataMember
	ContentMirrored []PromotionItemDto `json:"contentMirrored"`
	// @DataMember
	ContentDeleted []PromotionItemDto `json:"contentDeleted"`
	// @DataMember
	IntegrationsSeeded []PromotionItemDto `json:"integrationsSeeded"`
	// @DataMember
	IntegrationsSkipped []PromotionItemDto `json:"integrationsSkipped"`
	// @DataMember
	Blockers []PromotionBlockerDto `json:"blockers"`
	// @DataMember
	FromVersion *int64 `json:"fromVersion,omitempty"`
	// @DataMember
	WasDryRun bool `json:"wasDryRun,omitempty"`
}

// @DataContract
type ProjectEnvironmentsDto struct {
	// @DataMember
	Environments []string `json:"environments"`
}

type ProjectRegionDto struct {
	Id        string     `json:"id"`
	Continent *Continent `json:"continent,omitempty"`
	Name      *string    `json:"name,omitempty"`
}

// @DataContract
type ProjectBrandDto struct {
	// @DataMember
	MainColor *string `json:"mainColor,omitempty"`
	// @DataMember
	AccentColor *string `json:"accentColor,omitempty"`
	// @DataMember
	Logo *FileResourceRefDto `json:"logo,omitempty"`
	// @DataMember
	Icon *FileResourceRefDto `json:"icon,omitempty"`
}

type NotificationsSettingsGroupDto struct {
	Tag  string   `json:"tag"`
	Tags []string `json:"tags"`
}

type NotificationSettingsChannelDto struct {
	Channel CommunicationChannel            `json:"channel,omitempty"`
	Groups  []NotificationsSettingsGroupDto `json:"groups"`
}

type NotificationSettingsDto struct {
	Channels  []NotificationSettingsChannelDto `json:"channels"`
	AllGroups []GroupDefinitionDto             `json:"allGroups"`
	AllTags   []TagDefinitionDto               `json:"allTags"`
}

// @DataContract
type AuthenticationFlowPasswordPolicyDto struct {
	// @DataMember
	MinLength int `json:"minLength,omitempty"`
	// @DataMember
	MaxLength *int `json:"maxLength,omitempty"`
	// @DataMember
	MinNumbers *int `json:"minNumbers,omitempty"`
	// @DataMember
	MinUpper *int `json:"minUpper,omitempty"`
	// @DataMember
	MinLower *int `json:"minLower,omitempty"`
	// @DataMember
	MinSpecial *int `json:"minSpecial,omitempty"`
	// @DataMember
	AllowedSpecial *string `json:"allowedSpecial,omitempty"`
}

// @DataContract
type AuthenticationFlowSummaryDto struct {
	// @DataMember
	Type string `json:"type"`
	// @DataMember
	Provider *string `json:"provider,omitempty"`
	// @DataMember
	PasswordComplexity *AuthenticationFlowPasswordPolicyDto `json:"passwordComplexity,omitempty"`
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
}

// @DataContract
type SchemaTriggerDto struct {
	TriggerDto
	// @DataMember
	SchemaId string `json:"schemaId"`
	// @DataMember
	When SchemaTriggerType `json:"when,omitempty"`
	// @DataMember
	ConfigurationCode *string `json:"configurationCode,omitempty"`
}

// @DataContract
type DatabaseDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	Triggers []SchemaTriggerDto `json:"triggers,omitempty"`
	// @DataMember
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds"`
}

// @DataContract
type EmailDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds"`
}

// @DataContract
type AiDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	DefaultIntegrationViewId *string `json:"defaultIntegrationViewId,omitempty"`
}

// @DataContract
type MembershipTriggerDto struct {
	TriggerDto
	// @DataMember
	When MembershipTriggerType `json:"when,omitempty"`
}

type RoleItemDto struct {
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	DisplayName      *string  `json:"displayName,omitempty"`
	Description      *string  `json:"description,omitempty"`
	IsSystem         bool     `json:"isSystem,omitempty"`
	AttachedPolicies []string `json:"attachedPolicies,omitempty"`
}

type PermissionDto struct {
	Sid       *string          `json:"sid,omitempty"`
	Effect    PermissionEffect `json:"effect,omitempty"`
	Actions   []string         `json:"actions"`
	Resources []string         `json:"resources"`
}

type PolicyItemDto struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	IsSystem    bool            `json:"isSystem,omitempty"`
	Permissions []PermissionDto `json:"permissions,omitempty"`
}

// @DataContract
type AuthorizationDto struct {
	// @DataMember
	UserRegistersAsRole *string `json:"userRegistersAsRole,omitempty"`
	// @DataMember
	AllowedRegisterRoles []string `json:"allowedRegisterRoles,omitempty"`
	// @DataMember
	AllowedProviderRegisterRoles []string `json:"allowedProviderRegisterRoles,omitempty"`
}

// @DataContract
type MembershipDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	Triggers []MembershipTriggerDto `json:"triggers,omitempty"`
	// @DataMember
	CustomRoles []RoleItemDto `json:"customRoles,omitempty"`
	// @DataMember
	CustomPolicies []PolicyItemDto `json:"customPolicies,omitempty"`
	// @DataMember
	Authorization *AuthorizationDto `json:"authorization,omitempty"`
	// @DataMember
	RequireEmailValidation bool `json:"requireEmailValidation,omitempty"`
}

// @DataContract
type LoggingDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	IsEstablished bool `json:"isEstablished,omitempty"`
}

// @DataContract
type ServerEventsDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// @DataContract
type PushDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds"`
	// @DataMember
	MarketingTags []TagDefinitionDto `json:"marketingTags,omitempty"`
	// @DataMember
	TransactionalTags []TagDefinitionDto `json:"transactionalTags,omitempty"`
}

// @DataContract
type SchedulerDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// @DataContract
type CodeDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// @DataContract
type FilesTriggerDto struct {
	TriggerDto
	// @DataMember
	When FilesTriggerType `json:"when,omitempty"`
}

// @DataContract
type FilesDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	Triggers []FilesTriggerDto `json:"triggers,omitempty"`
	// @DataMember
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds"`
}

// @DataContract
type PaymentTriggerDto struct {
	TriggerDto
	// @DataMember
	When PaymentTriggerType `json:"when,omitempty"`
	// @DataMember
	Integrations []string `json:"integrations,omitempty"`
	// @DataMember
	Events []string `json:"events,omitempty"`
}

// @DataContract
type PaymentsDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	Triggers []PaymentTriggerDto `json:"triggers,omitempty"`
}

// @DataContract
type SmsDto struct {
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	DefaultIntegrationViewIds map[string]string `json:"defaultIntegrationViewIds"`
}

// @DataContract
type ProjectDto struct {
	// @DataMember
	AccountViewId string `json:"accountViewId"`
	// @DataMember
	ProjectStatus ProjectStatus `json:"projectStatus,omitempty"`
	// @DataMember
	IsActive bool `json:"isActive,omitempty"`
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	UniqueName string `json:"uniqueName"`
	// @DataMember
	HostLabel *string `json:"hostLabel,omitempty"`
	// @DataMember
	ApiHost *string `json:"apiHost,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	MarketingUrl *string `json:"marketingUrl,omitempty"`
	// @DataMember
	CanonicalAdminUrl *string `json:"canonicalAdminUrl,omitempty"`
	// @DataMember
	AdminUrl *string `json:"adminUrl,omitempty"`
	// @DataMember
	EffectiveAdminUrl *string `json:"effectiveAdminUrl,omitempty"`
	// @DataMember
	DefaultLanguage string `json:"defaultLanguage"`
	// @DataMember
	Languages []string `json:"languages"`
	// @DataMember
	PrimaryRegion *ProjectRegionDto `json:"primaryRegion,omitempty"`
	// @DataMember
	AdditionalRegions []ProjectRegionDto `json:"additionalRegions,omitempty"`
	// @DataMember
	IsMultiRegionEligible bool `json:"isMultiRegionEligible,omitempty"`
	// @DataMember
	Brand *ProjectBrandDto `json:"brand,omitempty"`
	// @DataMember
	NotificationSettings *NotificationSettingsDto `json:"notificationSettings,omitempty"`
	// @DataMember
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`
	// @DataMember
	ExposeBrandToAdminPortal bool `json:"exposeBrandToAdminPortal,omitempty"`
	// @DataMember
	ExposeAuthToAdminPortal bool `json:"exposeAuthToAdminPortal,omitempty"`
	// @DataMember
	AdminPortalEnabled bool `json:"adminPortalEnabled,omitempty"`
	// @DataMember
	AdminPortalServiceUserId *string `json:"adminPortalServiceUserId,omitempty"`
	// @DataMember
	MembershipAuthenticationFlows []AuthenticationFlowSummaryDto `json:"membershipAuthenticationFlows,omitempty"`
	// @DataMember
	ExposeLegalToAdminPortal bool `json:"exposeLegalToAdminPortal,omitempty"`
	// @DataMember
	LegalTermsMarkdown *string `json:"legalTermsMarkdown,omitempty"`
	// @DataMember
	LegalPrivacyMarkdown *string `json:"legalPrivacyMarkdown,omitempty"`
	// @DataMember
	Environments []string `json:"environments"`
	// @DataMember
	EnvironmentRanks map[string]int `json:"environmentRanks"`
	// @DataMember
	Database *DatabaseDto `json:"database,omitempty"`
	// @DataMember
	Email *EmailDto `json:"email,omitempty"`
	// @DataMember
	Ai *AiDto `json:"ai,omitempty"`
	// @DataMember
	Membership *MembershipDto `json:"membership,omitempty"`
	// @DataMember
	Logging *LoggingDto `json:"logging,omitempty"`
	// @DataMember
	ServerEvents *ServerEventsDto `json:"serverEvents,omitempty"`
	// @DataMember
	Push *PushDto `json:"push,omitempty"`
	// @DataMember
	Scheduler *SchedulerDto `json:"scheduler,omitempty"`
	// @DataMember
	Code *CodeDto `json:"code,omitempty"`
	// @DataMember
	Files *FilesDto `json:"files,omitempty"`
	// @DataMember
	Payments *PaymentsDto `json:"payments,omitempty"`
	// @DataMember
	Sms *SmsDto `json:"sms,omitempty"`
	// @DataMember
	DatabaseEnabled bool `json:"databaseEnabled,omitempty"`
	// @DataMember
	EmailEnabled bool `json:"emailEnabled,omitempty"`
	// @DataMember
	MembershipEnabled bool `json:"membershipEnabled,omitempty"`
	// @DataMember
	LoggingEnabled bool `json:"loggingEnabled,omitempty"`
	// @DataMember
	ServerEventsEnabled bool `json:"serverEventsEnabled,omitempty"`
	// @DataMember
	PushEnabled bool `json:"pushEnabled,omitempty"`
	// @DataMember
	SchedulerEnabled bool `json:"schedulerEnabled,omitempty"`
	// @DataMember
	CodeEnabled bool `json:"codeEnabled,omitempty"`
	// @DataMember
	FilesEnabled bool `json:"filesEnabled,omitempty"`
	// @DataMember
	PaymentsEnabled bool `json:"paymentsEnabled,omitempty"`
	// @DataMember
	SmsEnabled bool `json:"smsEnabled,omitempty"`
	// @DataMember
	DefaultLlmIntegrationViewId *string `json:"defaultLlmIntegrationViewId,omitempty"`
	// @DataMember
	Connections int `json:"connections,omitempty"`
}

// @DataContract
type ProjectListItemDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	IsActive bool `json:"isActive,omitempty"`
	// @DataMember
	ProjectStatus ProjectStatus `json:"projectStatus,omitempty"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	UniqueName string `json:"uniqueName"`
	// @DataMember
	PrimaryRegion *ProjectRegionDto `json:"primaryRegion,omitempty"`
	// @DataMember
	AdditionalRegions []ProjectRegionDto `json:"additionalRegions,omitempty"`
}

type PaginatedResponse struct {
	Items         IList   `json:"items"`
	HasMore       bool    `json:"hasMore,omitempty"`
	HasPrevious   bool    `json:"hasPrevious,omitempty"`
	StartingAfter *string `json:"startingAfter,omitempty"`
	EndingBefore  *string `json:"endingBefore,omitempty"`
}

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

type AccountPasswordPolicyDto struct {
	MinLength      int     `json:"minLength,omitempty"`
	MaxLength      *int    `json:"maxLength,omitempty"`
	MinNumbers     *int    `json:"minNumbers,omitempty"`
	MaxNumbers     *int    `json:"maxNumbers,omitempty"`
	MinUpper       *int    `json:"minUpper,omitempty"`
	MaxUpper       *int    `json:"maxUpper,omitempty"`
	MinLower       *int    `json:"minLower,omitempty"`
	MaxLower       *int    `json:"maxLower,omitempty"`
	MinSpecial     *int    `json:"minSpecial,omitempty"`
	MaxSpecial     *int    `json:"maxSpecial,omitempty"`
	AllowedSpecial *string `json:"allowedSpecial,omitempty"`
}

type AccountTeamRoleDto struct {
	Id          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	IsSystem    bool     `json:"isSystem,omitempty"`
	Policies    []string `json:"policies,omitempty"`
}

type AccountPasskeyListItemDto struct {
	CredentialId    string    `json:"credentialId"`
	FriendlyName    string    `json:"friendlyName"`
	RegisteredOnUtc time.Time `json:"registeredOnUtc,omitempty"`
	LastUsedOnUtc   time.Time `json:"lastUsedOnUtc,omitempty"`
	IsRevoked       bool      `json:"isRevoked,omitempty"`
}

// @DataContract
type LicenseDomainDnsRecordDto struct {
	// @DataMember
	Host string `json:"host"`
	// @DataMember
	RecordType string `json:"recordType"`
	// @DataMember
	Resolved bool `json:"resolved,omitempty"`
	// @DataMember
	Required bool `json:"required,omitempty"`
}

// @DataContract
type LicenseDomainDnsStatusDto struct {
	// @DataMember
	Domain string `json:"domain"`
	// @DataMember
	AllRequiredResolved bool `json:"allRequiredResolved,omitempty"`
	// @DataMember
	Records []LicenseDomainDnsRecordDto `json:"records"`
}

// @DataContract
type LicenseDomainVerificationChallengeDto struct {
	// @DataMember
	Domain string `json:"domain"`
	// @DataMember
	TxtHost string `json:"txtHost"`
	// @DataMember
	TxtValue string `json:"txtValue"`
	// @DataMember
	ExpiresAtUtc time.Time `json:"expiresAtUtc,omitempty"`
	// @DataMember
	Verified bool `json:"verified,omitempty"`
	// @DataMember
	VerifiedAtUtc *time.Time `json:"verifiedAtUtc,omitempty"`
	// @DataMember
	Skipped bool `json:"skipped,omitempty"`
}

// @DataContract
type LicenseDomainVerificationStatusDto struct {
	// @DataMember
	Domain string `json:"domain"`
	// @DataMember
	Verified bool `json:"verified,omitempty"`
	// @DataMember
	Skipped bool `json:"skipped,omitempty"`
	// @DataMember
	TxtHost *string `json:"txtHost,omitempty"`
	// @DataMember
	ExpectedTxtValue *string `json:"expectedTxtValue,omitempty"`
	// @DataMember
	ObservedTxtValue *string `json:"observedTxtValue,omitempty"`
	// @DataMember
	ExpiresAtUtc *time.Time `json:"expiresAtUtc,omitempty"`
	// @DataMember
	VerifiedAtUtc *time.Time `json:"verifiedAtUtc,omitempty"`
	// @DataMember
	Message *string `json:"message,omitempty"`
}

type CodeMashSubscriptionDto struct {
	ViewId            string    `json:"viewId"`
	Domain            string    `json:"domain"`
	WillExpireOn      time.Time `json:"willExpireOn,omitempty"`
	IssuedOn          time.Time `json:"issuedOn,omitempty"`
	IsTrial           bool      `json:"isTrial,omitempty"`
	SubscriptionRefId string    `json:"subscriptionRefId"`
}

type LicenseDto struct {
	CodeMashSubscriptionDto
	IsEnterprise bool `json:"isEnterprise,omitempty"`
	ProjectCap   int  `json:"projectCap,omitempty"`
}

// @DataContract
type LicenseHeartbeatVerdictDto struct {
	// @DataMember(Name="status")
	Status string `json:"status"`
	// @DataMember(Name="proofToken")
	ProofToken *string `json:"proofToken,omitempty"`
	// @DataMember(Name="serverTimeUtc")
	ServerTimeUtc time.Time `json:"serverTimeUtc,omitempty"`
	// @DataMember(Name="graceUntilUtc")
	GraceUntilUtc *time.Time `json:"graceUntilUtc,omitempty"`
	// @DataMember(Name="installationId")
	InstallationId string `json:"installationId,omitempty"`
	// @DataMember(Name="licenseAccountId")
	LicenseAccountId *string `json:"licenseAccountId,omitempty"`
	// @DataMember(Name="domain")
	Domain *string `json:"domain,omitempty"`
	// @DataMember(Name="signature")
	Signature *string `json:"signature,omitempty"`
	// @DataMember(Name="message")
	Message *string `json:"message,omitempty"`
}

// @DataContract
type InstallationLicenseStatusDto struct {
	// @DataMember(Name="storedMode")
	StoredMode string `json:"storedMode"`
	// @DataMember(Name="effectiveMode")
	EffectiveMode string `json:"effectiveMode"`
	// @DataMember(Name="isProduction")
	IsProduction bool `json:"isProduction,omitempty"`
	// @DataMember(Name="graceDaysLeft")
	GraceDaysLeft *int `json:"graceDaysLeft,omitempty"`
	// @DataMember(Name="graceUntilUtc")
	GraceUntilUtc *time.Time `json:"graceUntilUtc,omitempty"`
	// @DataMember(Name="lastProvenAtUtc")
	LastProvenAtUtc *time.Time `json:"lastProvenAtUtc,omitempty"`
	// @DataMember(Name="lastHeartbeatAtUtc")
	LastHeartbeatAtUtc *time.Time `json:"lastHeartbeatAtUtc,omitempty"`
	// @DataMember(Name="installationDomain")
	InstallationDomain *string `json:"installationDomain,omitempty"`
	// @DataMember(Name="licensedDomain")
	LicensedDomain *string `json:"licensedDomain,omitempty"`
	// @DataMember(Name="hostKind")
	HostKind *string `json:"hostKind,omitempty"`
	// @DataMember(Name="isTrialLicense")
	IsTrialLicense bool `json:"isTrialLicense,omitempty"`
	// @DataMember(Name="licenseExpireUtc")
	LicenseExpireUtc *time.Time `json:"licenseExpireUtc,omitempty"`
	// @DataMember(Name="message")
	Message *string `json:"message,omitempty"`
}

type ServiceUserApiKeyDto struct {
	Id            int        `json:"id,omitempty"`
	Name          string     `json:"name"`
	VisibleKey    *string    `json:"visibleKey,omitempty"`
	Scopes        []string   `json:"scopes"`
	CreatedDate   time.Time  `json:"createdDate,omitempty"`
	ExpiryDate    *time.Time `json:"expiryDate,omitempty"`
	CancelledDate *time.Time `json:"cancelledDate,omitempty"`
	Active        bool       `json:"active,omitempty"`
}

type GetTriggerResponse struct {
	ResponseBase
}

// @DataContract
type TriggerProjectionList struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	ActionType TriggerActionType `json:"actionType,omitempty"`
	// @DataMember
	HasPreExecuteCode bool `json:"hasPreExecuteCode,omitempty"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// @DataContract
type MembershipTriggerProjectionList struct {
	TriggerProjectionList
	// @DataMember
	Type MembershipTriggerType `json:"type,omitempty"`
	// @DataMember
	DestinationIds []string `json:"destinationIds,omitempty"`
}

type RoleListProjectionDto struct {
	ViewId      string  `json:"viewId"`
	Name        string  `json:"name"`
	DisplayName *string `json:"displayName,omitempty"`
	IsSystem    bool    `json:"isSystem,omitempty"`
	PolicyCount int     `json:"policyCount,omitempty"`
}

type PasskeySettingsDto struct {
	Enabled                       bool    `json:"enabled,omitempty"`
	CodeTtlMinutes                int     `json:"codeTtlMinutes,omitempty"`
	MaxCredentialsPerUser         int     `json:"maxCredentialsPerUser,omitempty"`
	RecoveryCodeCount             int     `json:"recoveryCodeCount,omitempty"`
	GenerateRecoveryCodesAtSignup bool    `json:"generateRecoveryCodesAtSignup,omitempty"`
	AuthenticatorAttachment       string  `json:"authenticatorAttachment"`
	AllowMagicLinkRecovery        bool    `json:"allowMagicLinkRecovery,omitempty"`
	RefreshTokenTtlDays           int     `json:"refreshTokenTtlDays,omitempty"`
	RpId                          *string `json:"rpId,omitempty"`
}

type IntegrationListProjection struct {
	ViewId                            string        `json:"viewId"`
	IntegrationName                   string        `json:"integrationName"`
	IsEnabled                         bool          `json:"isEnabled,omitempty"`
	LastIntegrationTestAtUtc          *time.Time    `json:"lastIntegrationTestAtUtc,omitempty"`
	LastIntegrationTestSucceeded      *bool         `json:"lastIntegrationTestSucceeded,omitempty"`
	LastIntegrationTestErrors         IReadOnlyList `json:"lastIntegrationTestErrors"`
	HumanDeliveryConfirmedAtUtc       *time.Time    `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	RequiresHumanDeliveryConfirmation bool          `json:"requiresHumanDeliveryConfirmation,omitempty"`
}

type MembershipIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider MembershipProvider `json:"provider,omitempty"`
}

// @DataContract
type MembershipMessageTemplateDto struct {
	// @DataMember
	Id *string `json:"id,omitempty"`
}

// @DataContract
type MembershipEmailActionSettingsDto struct {
	// @DataMember
	SendEmail bool `json:"sendEmail,omitempty"`
	// @DataMember
	Template *MembershipMessageTemplateDto `json:"template,omitempty"`
	// @DataMember
	Callback *string `json:"callback,omitempty"`
}

// @DataContract
type MembershipEmailPreferencesDto struct {
	// @DataMember
	RegistrationViaEmail *MembershipEmailActionSettingsDto `json:"registrationViaEmail,omitempty"`
	// @DataMember
	VerificationViaEmail *MembershipEmailActionSettingsDto `json:"verificationViaEmail,omitempty"`
	// @DataMember
	PasswordResetViaEmail *MembershipEmailActionSettingsDto `json:"passwordResetViaEmail,omitempty"`
	// @DataMember
	InvitationViaEmail *MembershipEmailActionSettingsDto `json:"invitationViaEmail,omitempty"`
	// @DataMember
	DeactivationViaEmail *MembershipEmailActionSettingsDto `json:"deactivationViaEmail,omitempty"`
}

// @DataContract
type PasswordComplexityDto struct {
	// @DataMember
	MinLength int `json:"minLength,omitempty"`
	// @DataMember
	MaxLength *int `json:"maxLength,omitempty"`
	// @DataMember
	MinNumbers *int `json:"minNumbers,omitempty"`
	// @DataMember
	MaxNumbers *int `json:"maxNumbers,omitempty"`
	// @DataMember
	MinUpper *int `json:"minUpper,omitempty"`
	// @DataMember
	MaxUpper *int `json:"maxUpper,omitempty"`
	// @DataMember
	MinLower *int `json:"minLower,omitempty"`
	// @DataMember
	MaxLower *int `json:"maxLower,omitempty"`
	// @DataMember
	MinSpecial *int `json:"minSpecial,omitempty"`
	// @DataMember
	MaxSpecial *int `json:"maxSpecial,omitempty"`
	// @DataMember
	AllowedSpecial *string `json:"allowedSpecial,omitempty"`
}

// @DataContract
type MembershipAuthorizationViewDto struct {
	// @DataMember
	EmailPreferences *MembershipEmailPreferencesDto `json:"emailPreferences,omitempty"`
	// @DataMember
	UserRegistersAsRole *string `json:"userRegistersAsRole,omitempty"`
	// @DataMember
	GuestRegistersAsRole *string `json:"guestRegistersAsRole,omitempty"`
	// @DataMember
	AllowedRegisterRoles []string `json:"allowedRegisterRoles,omitempty"`
	// @DataMember
	AllowedProviderRegisterRoles []string `json:"allowedProviderRegisterRoles,omitempty"`
	// @DataMember
	ResetPasswordTokenExpiration *int `json:"resetPasswordTokenExpiration,omitempty"`
	// @DataMember
	InvitationExpiration *int `json:"invitationExpiration,omitempty"`
	// @DataMember
	EmailVerificationExpiration *int `json:"emailVerificationExpiration,omitempty"`
	// @DataMember
	DeactivationExpiration *int `json:"deactivationExpiration,omitempty"`
	// @DataMember
	DefaultSubscribeToNews bool `json:"defaultSubscribeToNews,omitempty"`
	// @DataMember
	PasswordComplexity *PasswordComplexityDto `json:"passwordComplexity,omitempty"`
}

// @DataContract
type MembershipCredentialsSettingsModeDto struct {
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	LogoutUrl *string `json:"logoutUrl,omitempty"`
}

// @DataContract
type MembershipCredentialsSettingsDto struct {
	// @DataMember
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// @DataMember
	AllowUsernames bool `json:"allowUsernames,omitempty"`
	// @DataMember
	Modes []MembershipCredentialsSettingsModeDto `json:"modes,omitempty"`
}

// @DataContract
type MembershipAuthenticationViewDto struct {
	// @DataMember
	CredentialsSettings *MembershipCredentialsSettingsDto `json:"credentialsSettings,omitempty"`
	// @DataMember
	Flows []string `json:"flows"`
}

// @DataContract
type SchemaTriggerProjectionList struct {
	TriggerProjectionList
	// @DataMember
	Type SchemaTriggerType `json:"type,omitempty"`
}

type JsonSchemaFieldDto struct {
	// @DataMember
	FieldName string `json:"fieldName"`
}

type DataSchemaDto struct {
	// @DataMember
	Json string `json:"json"`
	// @DataMember
	Fields []JsonSchemaFieldDto `json:"fields"`
}

type VisualSchemaDto struct {
	// @DataMember
	Json string `json:"json"`
}

type TaxonomyDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Slug *string `json:"slug,omitempty"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	TermsMetaDataSchema *DataSchemaDto `json:"termsMetaDataSchema,omitempty"`
	// @DataMember
	TermsMetaVisualSchema *VisualSchemaDto `json:"termsMetaVisualSchema,omitempty"`
	// @DataMember
	Dependencies []string `json:"dependencies,omitempty"`
}

type TaxonomyListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	TaxonomyName string `json:"taxonomyName"`
	// @DataMember
	TaxonomySlug string `json:"taxonomySlug"`
	// @DataMember
	ParentId *string `json:"parentId,omitempty"`
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
	Description *string `json:"description,omitempty"`
	// @DataMember
	Descriptions map[string]string `json:"descriptions,omitempty"`
	// @DataMember
	MultiParents []TermMultiParentDto `json:"multiParents,omitempty"`
	// @DataMember
	Meta *Object `json:"meta,omitempty"`
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
	Description *string `json:"description,omitempty"`
	// @DataMember
	Descriptions map[string]string `json:"descriptions,omitempty"`
	// @DataMember
	MultiParents []TermMultiParentDto `json:"multiParents,omitempty"`
	// @DataMember
	Meta *Object `json:"meta,omitempty"`
}

// @DataContract
type AppliedTaxonomyDto struct {
	// @DataMember
	CatalogId string `json:"catalogId"`
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	Slug string `json:"slug"`
	// @DataMember
	Title string `json:"title"`
	// @DataMember
	Action string `json:"action"`
	// @DataMember
	TermsCreated int `json:"termsCreated,omitempty"`
}

// @DataContract
type AppliedCollectionDto struct {
	// @DataMember
	Entity string `json:"entity"`
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Title string `json:"title"`
	// @DataMember
	Action string `json:"action"`
	// @DataMember
	Published bool `json:"published,omitempty"`
	// @DataMember
	LinkedFields []string `json:"linkedFields"`
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
}

type SchemaDraftDto struct {
	// @DataMember
	DataSchema DataSchemaDto `json:"dataSchema"`
	// @DataMember
	VisualSchema VisualSchemaDto `json:"visualSchema"`
	// @DataMember
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

type SchemaDiffDto struct {
	// @DataMember
	FromVersion int `json:"fromVersion,omitempty"`
	// @DataMember
	ToVersion int `json:"toVersion,omitempty"`
	// @DataMember
	Added []string `json:"added"`
	// @DataMember
	Removed []string `json:"removed"`
	// @DataMember
	TypeChanged []string `json:"typeChanged"`
	// @DataMember
	ValidatorTightened []string `json:"validatorTightened"`
}

type SchemaVersionSummaryDto struct {
	// @DataMember
	Version int `json:"version,omitempty"`
	// @DataMember
	MetaSchemaVersion int `json:"metaSchemaVersion,omitempty"`
	// @DataMember
	PublishedAt time.Time `json:"publishedAt,omitempty"`
}

type CollectionIndexKeyDto struct {
	// @DataMember
	Field string `json:"field"`
	// @DataMember
	Order int `json:"order,omitempty"`
}

type CollectionIndexDto struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Keys []CollectionIndexKeyDto `json:"keys"`
}

// @DataContract
type SeedCollectionReportItemDto struct {
	// @DataMember
	CollectionName string `json:"collectionName"`
	// @DataMember
	Requested int `json:"requested,omitempty"`
	// @DataMember
	Inserted int `json:"inserted,omitempty"`
	// @DataMember
	Ids []string `json:"ids"`
	// @DataMember
	Errors []string `json:"errors"`
}

// @DataContract
type SeedCollectionRecordsResultDto struct {
	// @DataMember
	InsertOrder []string `json:"insertOrder"`
	// @DataMember
	Report []SeedCollectionReportItemDto `json:"report"`
}

type DatabaseIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider DatabaseProvider `json:"provider,omitempty"`
}

// @DataContract
type FlexTierDto struct {
	// @DataMember
	Code string `json:"code"`
	// @DataMember
	Step int `json:"step,omitempty"`
	// @DataMember
	DisplayName string `json:"displayName"`
}

// @DataContract
type IntegrationTestResultItemDto struct {
	// @DataMember
	Operation string `json:"operation"`
	// @DataMember
	Result string `json:"result"`
	// @DataMember
	Errors *IReadOnlyList `json:"errors,omitempty"`
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

type MongoDbAggregateListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	SchemaViewId string `json:"schemaViewId"`
}

// @DataContract
type FilesTriggerProjectionList struct {
	TriggerProjectionList
	// @DataMember
	Type FilesTriggerType `json:"type,omitempty"`
}

type FilesIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider FileProvider `json:"provider,omitempty"`
}

// @DataContract
type NotificationModuleDependencyItemDto struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	ViewId *string `json:"viewId,omitempty"`
	// @DataMember
	Category *string `json:"category,omitempty"`
}

// @DataContract
type NotificationModuleDisableDependenciesDto struct {
	// @DataMember
	Triggers []NotificationModuleDependencyItemDto `json:"triggers"`
	// @DataMember
	SchedulerTasks []NotificationModuleDependencyItemDto `json:"schedulerTasks"`
	// @DataMember
	InFlightCampaigns []NotificationModuleDependencyItemDto `json:"inFlightCampaigns"`
	// @DataMember
	MembershipSettings []NotificationModuleDependencyItemDto `json:"membershipSettings"`
}

// @DataContract
type TestEmailValidationItemDto struct {
	// @DataMember
	Address string `json:"address"`
	// @DataMember
	Verdict string `json:"verdict"`
	// @DataMember
	Reason *string `json:"reason,omitempty"`
	// @DataMember
	Score *float64 `json:"score,omitempty"`
}

// @DataContract
type TemplateListProjection struct {
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	TemplateName string `json:"templateName"`
	// @DataMember
	IsActive bool `json:"isActive,omitempty"`
	// @DataMember
	Type CommunicationChannel `json:"type,omitempty"`
	// @DataMember
	Tags []string `json:"tags,omitempty"`
}

type EmailTemplateListProjection struct {
	TemplateListProjection
	HasAttachments bool          `json:"hasAttachments,omitempty"`
	Languages      IReadOnlyList `json:"languages"`
}

// @DataContract
type MjmlParseError struct {
	// @DataMember(Name="line")
	Line int `json:"line,omitempty"`
	// @DataMember(Name="message")
	Message *string `json:"message,omitempty"`
	// @DataMember(Name="tagName")
	TagName *string `json:"tagName,omitempty"`
	// @DataMember(Name="formattedMessage")
	FormattedMessage *string `json:"formattedMessage,omitempty"`
}

// @DataContract
type HtmlFromMjmlResponse struct {
	// @DataMember(Name="html")
	Html string `json:"html"`
	// @DataMember(Name="errors")
	Errors []MjmlParseError `json:"errors"`
}

type SystemEmailTemplateListProjection struct {
	EmailTemplateListProjection
}

// @DataContract
type EmailSignatureDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember
	Translations []TranslationDto `json:"translations"`
}

// @DataContract
type ListItemProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	DisplayName *string `json:"displayName,omitempty"`
}

// @DataContract
type ListItemWithTranslationsProjection struct {
	ListItemProjection
	// @DataMember
	Translations []string `json:"translations"`
}

// @DataContract
type EmailFooterDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember
	Translations []TranslationDto `json:"translations"`
}

type EmailSettings struct {
	Signatures *IList `json:"signatures,omitempty"`
	Footers    *IList `json:"footers,omitempty"`
}

// @DataContract
type DomainHealthRecordItemDto struct {
	// @DataMember
	Record string `json:"record"`
	// @DataMember
	Value string `json:"value"`
}

type EmailIntegrationListProjection struct {
	IntegrationListProjection
	EmailProvider      EmailProvider `json:"emailProvider,omitempty"`
	SenderEmailAddress string        `json:"senderEmailAddress"`
	SenderDisplayName  *string       `json:"senderDisplayName,omitempty"`
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
	UserId string `json:"userId"`
	// @DataMember
	Id *string `json:"id,omitempty"`
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
	Time     time.Time                  `json:"time,omitempty"`
	Status   CampaignNotificationStatus `json:"status,omitempty"`
	SourceId *string                    `json:"sourceId,omitempty"`
	Errors   []ErrorDto                 `json:"errors,omitempty"`
	Tags     []string                   `json:"tags,omitempty"`
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
	Subject *string `json:"subject,omitempty"`
	// @DataMember
	Body *string `json:"body,omitempty"`
	// @DataMember
	Model map[string]string `json:"model,omitempty"`
	// @DataMember
	StatusHistory []NotificationStatusChangeEntryDto `json:"statusHistory"`
	// @DataMember
	Id string `json:"id"`
}

// @DataContract
type EmailCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	// @DataMember
	Recipients EmailRecipientsDto `json:"recipients"`
	// @DataMember
	Content *EmailMessageContentDto `json:"content,omitempty"`
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
	SuccessRate float64 `json:"successRate,omitempty"`
}

// @DataContract
type SmsTemplateListProjection struct {
	TemplateListProjection
}

type SmsSettings struct {
}

type SmsIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider SmsProvider `json:"provider,omitempty"`
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
type MarketplaceListingProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Slug string `json:"slug"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	Vendor string `json:"vendor"`
	// @DataMember
	Category MarketplaceCategory `json:"category,omitempty"`
	// @DataMember
	Transport MarketplaceTransport `json:"transport,omitempty"`
	// @DataMember
	IconUrl *string `json:"iconUrl,omitempty"`
	// @DataMember
	IsOfficial bool `json:"isOfficial,omitempty"`
	// @DataMember
	Tags IReadOnlyList `json:"tags"`
	// @DataMember
	FunctionCount int `json:"functionCount,omitempty"`
}

// @DataContract
type MarketplaceIntegrationListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	IntegrationName string `json:"integrationName"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	ListingViewId string `json:"listingViewId"`
	// @DataMember
	Vendor string `json:"vendor"`
	// @DataMember
	Category MarketplaceCategory `json:"category,omitempty"`
	// @DataMember
	Transport MarketplaceTransport `json:"transport,omitempty"`
	// @DataMember
	LastIntegrationTestAtUtc *time.Time `json:"lastIntegrationTestAtUtc,omitempty"`
	// @DataMember
	LastIntegrationTestSucceeded *bool `json:"lastIntegrationTestSucceeded,omitempty"`
	// @DataMember
	LastIntegrationTestErrors IReadOnlyList `json:"lastIntegrationTestErrors"`
	// @DataMember
	HumanDeliveryConfirmedAtUtc *time.Time `json:"humanDeliveryConfirmedAtUtc,omitempty"`
	// @DataMember
	RequiresHumanDeliveryConfirmation bool `json:"requiresHumanDeliveryConfirmation,omitempty"`
}

// @DataContract
type MarketplaceFunctionProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	IntegrationViewId string `json:"integrationViewId"`
	// @DataMember
	FunctionKey string `json:"functionKey"`
	// @DataMember
	DisplayName string `json:"displayName"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	MappingCount int `json:"mappingCount,omitempty"`
}

type CodeIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider CodeProvider `json:"provider,omitempty"`
}

// @DataContract
type PushTemplateListProjection struct {
	TemplateListProjection
}

type PushSettings struct {
	MarketingTags     []TagDefinitionDto `json:"marketingTags,omitempty"`
	TransactionalTags []TagDefinitionDto `json:"transactionalTags,omitempty"`
}

type PushIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider PushProvider `json:"provider,omitempty"`
}

// @DataContract
type PushCampaignDto struct {
	CampaignDto
	// @DataMember
	Recipients PushCampaignDeliverySettingsDto `json:"recipients"`
	// @DataMember
	Template PushTemplateDto `json:"template"`
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
type PushCampaignBatchNotificationDto struct {
	CampaignBatchNotificationDto
	// @DataMember
	Recipients PushRecipientsDto `json:"recipients"`
	// @DataMember
	Content *PushMessageContentDto `json:"content,omitempty"`
}

type PaymentsWebhookLogEntry struct {
	IntegrationId   string    `json:"integrationId"`
	Source          string    `json:"source"`
	EventName       *string   `json:"eventName,omitempty"`
	ProviderEventId *string   `json:"providerEventId,omitempty"`
	StatusCode      int       `json:"statusCode,omitempty"`
	Description     string    `json:"description"`
	ReceivedOn      time.Time `json:"receivedOn,omitempty"`
}

// @DataContract
type PaymentTriggerProjectionList struct {
	TriggerProjectionList
	// @DataMember
	Type PaymentTriggerType `json:"type,omitempty"`
	// @DataMember
	Integrations []string `json:"integrations,omitempty"`
	// @DataMember
	Events []string `json:"events,omitempty"`
}

type PaymentsIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	GatewayPlatform PaymentGatewayPlatform `json:"gatewayPlatform,omitempty"`
}

type LoggingIntegrationListProjection struct {
	IntegrationListProjection
	// @DataMember
	Provider LoggingProvider `json:"provider,omitempty"`
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
	Meta *IReadOnlyDictionary `json:"meta,omitempty"`
}

type AiToolManifestParameter struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Required    bool    `json:"required,omitempty"`
	Description *string `json:"description,omitempty"`
}

type AiToolManifestItem struct {
	Name                 string                    `json:"name"`
	Description          string                    `json:"description"`
	Toolsets             []string                  `json:"toolsets"`
	RequiresConfirmation bool                      `json:"requiresConfirmation,omitempty"`
	Parameters           []AiToolManifestParameter `json:"parameters"`
}

type ChatModelOption struct {
	LlmIntegrationId *string `json:"llmIntegrationId,omitempty"`
	Kind             string  `json:"kind"`
	Provider         string  `json:"provider"`
	Model            string  `json:"model"`
	Label            string  `json:"label"`
	IsDefault        bool    `json:"isDefault,omitempty"`
	IsAuto           bool    `json:"isAuto,omitempty"`
	ContextWindow    int     `json:"contextWindow,omitempty"`
}

type ChatMemoryNote struct {
	Id           string    `json:"id"`
	Kind         string    `json:"kind"`
	Text         string    `json:"text"`
	ProjectId    *string   `json:"projectId,omitempty"`
	CreatedAtUtc time.Time `json:"createdAtUtc,omitempty"`
}

type ChatSessionListItem struct {
	SessionId    string    `json:"sessionId"`
	Profile      string    `json:"profile"`
	ProjectId    *string   `json:"projectId,omitempty"`
	Env          *string   `json:"env,omitempty"`
	Title        *string   `json:"title,omitempty"`
	UpdatedAtUtc time.Time `json:"updatedAtUtc,omitempty"`
	IsArchived   bool      `json:"isArchived,omitempty"`
	IsPinned     bool      `json:"isPinned,omitempty"`
}

type ProjectBriefSourceWireDto struct {
	Kind       string    `json:"kind"`
	SessionId  *string   `json:"sessionId,omitempty"`
	EntryId    *string   `json:"entryId,omitempty"`
	EntrySeq   *int64    `json:"entrySeq,omitempty"`
	EventId    *string   `json:"eventId,omitempty"`
	UserAuthId *string   `json:"userAuthId,omitempty"`
	AtUtc      time.Time `json:"atUtc,omitempty"`
	Surface    *string   `json:"surface,omitempty"`
	Quote      *string   `json:"quote,omitempty"`
	WorkItemId *string   `json:"workItemId,omitempty"`
}

type ProjectBriefSatisfiedByWireDto struct {
	ArtifactId string `json:"artifactId"`
	Tool       string `json:"tool"`
	Orphaned   bool   `json:"orphaned,omitempty"`
}

type ProjectBriefRequirementWireDto struct {
	Id           string                           `json:"id"`
	Text         string                           `json:"text"`
	Status       string                           `json:"status"`
	Confidence   float64                          `json:"confidence,omitempty"`
	IsAssumption bool                             `json:"isAssumption,omitempty"`
	Sources      []ProjectBriefSourceWireDto      `json:"sources"`
	SatisfiedBy  []ProjectBriefSatisfiedByWireDto `json:"satisfiedBy"`
	SinceEventId string                           `json:"sinceEventId"`
}

type ProjectBriefDecisionWireDto struct {
	EventId    string                      `json:"eventId"`
	Text       string                      `json:"text"`
	Confidence float64                     `json:"confidence,omitempty"`
	AtUtc      time.Time                   `json:"atUtc,omitempty"`
	Sources    []ProjectBriefSourceWireDto `json:"sources"`
}

type ProjectBriefAssumptionWireDto struct {
	RequirementId string    `json:"requirementId"`
	EventId       string    `json:"eventId"`
	Text          string    `json:"text"`
	Confidence    float64   `json:"confidence,omitempty"`
	AtUtc         time.Time `json:"atUtc,omitempty"`
}

type ProjectBriefSnapshotWireDto struct {
	ProjectId       string                           `json:"projectId"`
	UpToSeq         int64                            `json:"upToSeq,omitempty"`
	AtUtc           time.Time                        `json:"atUtc,omitempty"`
	Requirements    []ProjectBriefRequirementWireDto `json:"requirements"`
	Decisions       []ProjectBriefDecisionWireDto    `json:"decisions"`
	OpenAssumptions []ProjectBriefAssumptionWireDto  `json:"openAssumptions"`
	Summary         *string                          `json:"summary,omitempty"`
}

type ProjectBriefEventWireDto struct {
	Id                string                           `json:"id"`
	ProjectId         string                           `json:"projectId"`
	Seq               int64                            `json:"seq,omitempty"`
	AtUtc             time.Time                        `json:"atUtc,omitempty"`
	Kind              string                           `json:"kind"`
	RequirementId     *string                          `json:"requirementId,omitempty"`
	SupersedesEventId *string                          `json:"supersedesEventId,omitempty"`
	Text              string                           `json:"text"`
	Confidence        float64                          `json:"confidence,omitempty"`
	Sources           []ProjectBriefSourceWireDto      `json:"sources"`
	Origin            string                           `json:"origin"`
	SatisfiedBy       []ProjectBriefSatisfiedByWireDto `json:"satisfiedBy"`
}

type WorkItemEntryRefWireDto struct {
	SessionId string `json:"sessionId"`
	EntryId   string `json:"entryId"`
}

type WorkItemRunRefWireDto struct {
	SessionId   string  `json:"sessionId"`
	PlanEntryId string  `json:"planEntryId"`
	RunId       *string `json:"runId,omitempty"`
}

type WorkItemArtifactWireDto struct {
	ArtifactId  string  `json:"artifactId"`
	What        string  `json:"what"`
	Step        int     `json:"step,omitempty"`
	PlanEntryId *string `json:"planEntryId,omitempty"`
	Kind        *string `json:"kind,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type WorkItemMovedOutWireDto struct {
	Text     string                   `json:"text"`
	Reason   string                   `json:"reason"`
	MovedTo  *string                  `json:"movedTo,omitempty"`
	Source   string                   `json:"source"`
	EntryRef *WorkItemEntryRefWireDto `json:"entryRef,omitempty"`
	AtUtc    time.Time                `json:"atUtc,omitempty"`
}

type WorkItemNeedsYouWireDto struct {
	Text             string                   `json:"text"`
	Kind             string                   `json:"kind"`
	EntryRef         *WorkItemEntryRefWireDto `json:"entryRef,omitempty"`
	Done             bool                     `json:"done,omitempty"`
	DoneAtUtc        *time.Time               `json:"doneAtUtc,omitempty"`
	DoneByUserAuthId *string                  `json:"doneByUserAuthId,omitempty"`
}

type WorkItemOpenQuestionWireDto struct {
	EntryRef WorkItemEntryRefWireDto `json:"entryRef"`
	Blocking bool                    `json:"blocking,omitempty"`
	Text     string                  `json:"text"`
	AtUtc    time.Time               `json:"atUtc,omitempty"`
}

type WorkItemDoneConditionWireDto struct {
	Condition int    `json:"condition,omitempty"`
	Holds     bool   `json:"holds,omitempty"`
	Reason    string `json:"reason"`
}

type WorkItemWireDto struct {
	Id                  string                         `json:"id"`
	ProjectId           *string                        `json:"projectId,omitempty"`
	Status              string                         `json:"status"`
	Goal                string                         `json:"goal"`
	NotInScope          []string                       `json:"notInScope"`
	ScopeRequirementIds []string                       `json:"scopeRequirementIds"`
	Difficulty          *int                           `json:"difficulty,omitempty"`
	DifficultyReason    *string                        `json:"difficultyReason,omitempty"`
	PlanEntryRefs       []WorkItemEntryRefWireDto      `json:"planEntryRefs"`
	RunRefs             []WorkItemRunRefWireDto        `json:"runRefs"`
	Artifacts           []WorkItemArtifactWireDto      `json:"artifacts"`
	MovedOut            []WorkItemMovedOutWireDto      `json:"movedOut"`
	NeedsYou            []WorkItemNeedsYouWireDto      `json:"needsYou"`
	OpenQuestions       []WorkItemOpenQuestionWireDto  `json:"openQuestions"`
	SessionIds          []string                       `json:"sessionIds"`
	ParentId            *string                        `json:"parentId,omitempty"`
	Children            []string                       `json:"children"`
	SummaryEntryRef     *WorkItemEntryRefWireDto       `json:"summaryEntryRef,omitempty"`
	CreatedBy           *string                        `json:"createdBy,omitempty"`
	CreatedAtUtc        time.Time                      `json:"createdAtUtc,omitempty"`
	UpdatedAtUtc        time.Time                      `json:"updatedAtUtc,omitempty"`
	DoneVerdict         *string                        `json:"doneVerdict,omitempty"`
	DoneConditions      []WorkItemDoneConditionWireDto `json:"doneConditions"`
}

type LlmIntegrationListProjection struct {
	IntegrationListProjection
	LlmProvider  LlmProvider `json:"llmProvider,omitempty"`
	BaseUrl      *string     `json:"baseUrl,omitempty"`
	DefaultModel *string     `json:"defaultModel,omitempty"`
}

type McpIntegrationListProjection struct {
	IntegrationListProjection
	McpProvider McpProvider  `json:"mcpProvider,omitempty"`
	Transport   McpTransport `json:"transport,omitempty"`
	Category    *string      `json:"category,omitempty"`
	Description *string      `json:"description,omitempty"`
	Icon        *string      `json:"icon,omitempty"`
}

type IVirtualDirectory struct {
}

type IVirtualPathProvider struct {
	RootDirectory        *IVirtualDirectory `json:"rootDirectory,omitempty"`
	VirtualPathSeparator *string            `json:"virtualPathSeparator,omitempty"`
	RealPathSeparator    *string            `json:"realPathSeparator,omitempty"`
}

type IVirtualFile struct {
	VirtualPathProvider *IVirtualPathProvider `json:"virtualPathProvider,omitempty"`
	Extension           *string               `json:"extension,omitempty"`
	Length              int64                 `json:"length,omitempty"`
}

// @Flags()
type CacheControl int

const (
	CacheControlNone            CacheControl = 0
	CacheControlPublic          CacheControl = 1
	CacheControlPrivate         CacheControl = 2
	CacheControlMustRevalidate  CacheControl = 4
	CacheControlNoCache         CacheControl = 8
	CacheControlNoStore         CacheControl = 16
	CacheControlNoTransform     CacheControl = 32
	CacheControlProxyRevalidate CacheControl = 64
)

type IContentTypeWriter struct {
}

// @Flags()
type RequestAttributes int

const (
	RequestAttributesNone                  RequestAttributes = 0
	RequestAttributesLocalhost             RequestAttributes = 1
	RequestAttributesLocalSubnet           RequestAttributes = 2
	RequestAttributesExternal              RequestAttributes = 4
	RequestAttributesSecure                RequestAttributes = 8
	RequestAttributesInSecure              RequestAttributes = 16
	RequestAttributesAnySecurityMode       RequestAttributes = 24
	RequestAttributesHttpHead              RequestAttributes = 32
	RequestAttributesHttpGet               RequestAttributes = 64
	RequestAttributesHttpPost              RequestAttributes = 128
	RequestAttributesHttpPut               RequestAttributes = 256
	RequestAttributesHttpDelete            RequestAttributes = 512
	RequestAttributesHttpPatch             RequestAttributes = 1024
	RequestAttributesHttpOptions           RequestAttributes = 2048
	RequestAttributesHttpOther             RequestAttributes = 4096
	RequestAttributesAnyHttpMethod         RequestAttributes = 8160
	RequestAttributesOneWay                RequestAttributes = 8192
	RequestAttributesReply                 RequestAttributes = 16384
	RequestAttributesAnyCallStyle          RequestAttributes = 24576
	RequestAttributesSoap11                RequestAttributes = 32768
	RequestAttributesSoap12                RequestAttributes = 65536
	RequestAttributesXml                   RequestAttributes = 131072
	RequestAttributesJson                  RequestAttributes = 262144
	RequestAttributesJsv                   RequestAttributes = 524288
	RequestAttributesProtoBuf              RequestAttributes = 1048576
	RequestAttributesCsv                   RequestAttributes = 2097152
	RequestAttributesHtml                  RequestAttributes = 4194304
	RequestAttributesJsonl                 RequestAttributes = 8388608
	RequestAttributesMsgPack               RequestAttributes = 16777216
	RequestAttributesFormatOther           RequestAttributes = 33554432
	RequestAttributesAnyFormat             RequestAttributes = 67076096
	RequestAttributesHttp                  RequestAttributes = 67108864
	RequestAttributesMessageQueue          RequestAttributes = 134217728
	RequestAttributesTcp                   RequestAttributes = 268435456
	RequestAttributesGrpc                  RequestAttributes = 536870912
	RequestAttributesEndpointOther         RequestAttributes = 1073741824
	RequestAttributesAnyEndpoint           RequestAttributes = 2080374784
	RequestAttributesInProcess             RequestAttributes = -2147483648
	RequestAttributesInternalNetworkAccess RequestAttributes = -2147483645
	RequestAttributesAnyNetworkAccessType  RequestAttributes = -2147483641
	RequestAttributesAny                   RequestAttributes = -1
)

type IRequestPreferences struct {
	AcceptsBrotli  bool `json:"acceptsBrotli,omitempty"`
	AcceptsDeflate bool `json:"acceptsDeflate,omitempty"`
	AcceptsGzip    bool `json:"acceptsGzip,omitempty"`
}

type IHttpFile struct {
	Name          *string `json:"name,omitempty"`
	FileName      *string `json:"fileName,omitempty"`
	ContentLength int64   `json:"contentLength,omitempty"`
	ContentType   *string `json:"contentType,omitempty"`
	InputStream   []byte  `json:"inputStream,omitempty"`
}

type IRequest struct {
	OriginalRequest                *Object              `json:"originalRequest,omitempty"`
	Response                       *IResponse           `json:"response,omitempty"`
	OperationName                  *string              `json:"operationName,omitempty"`
	Verb                           *string              `json:"verb,omitempty"`
	RequestAttributes              RequestAttributes    `json:"requestAttributes,omitempty"`
	RequestPreferences             *IRequestPreferences `json:"requestPreferences,omitempty"`
	Dto                            *Object              `json:"dto,omitempty"`
	ContentType                    *string              `json:"contentType,omitempty"`
	IsLocal                        bool                 `json:"isLocal,omitempty"`
	UserAgent                      *string              `json:"userAgent,omitempty"`
	Cookies                        map[string]Cookie    `json:"cookies,omitempty"`
	ResponseContentType            *string              `json:"responseContentType,omitempty"`
	HasExplicitResponseContentType bool                 `json:"hasExplicitResponseContentType,omitempty"`
	Items                          map[string]Object    `json:"items,omitempty"`
	Headers                        *NameValueCollection `json:"headers,omitempty"`
	QueryString                    *NameValueCollection `json:"queryString,omitempty"`
	FormData                       *NameValueCollection `json:"formData,omitempty"`
	UseBufferedStream              bool                 `json:"useBufferedStream,omitempty"`
	RawUrl                         *string              `json:"rawUrl,omitempty"`
	AbsoluteUri                    *string              `json:"absoluteUri,omitempty"`
	UserHostAddress                *string              `json:"userHostAddress,omitempty"`
	RemoteIp                       *string              `json:"remoteIp,omitempty"`
	Authorization                  *string              `json:"authorization,omitempty"`
	IsSecureConnection             bool                 `json:"isSecureConnection,omitempty"`
	AcceptTypes                    []string             `json:"acceptTypes,omitempty"`
	PathInfo                       *string              `json:"pathInfo,omitempty"`
	OriginalPathInfo               *string              `json:"originalPathInfo,omitempty"`
	InputStream                    []byte               `json:"inputStream,omitempty"`
	ContentLength                  int64                `json:"contentLength,omitempty"`
	Files                          []IHttpFile          `json:"files,omitempty"`
	UrlReferrer                    *string              `json:"urlReferrer,omitempty"`
	RequestAborted                 CancellationToken    `json:"requestAborted,omitempty"`
}

type IResponse struct {
	OriginalResponse  *Object           `json:"originalResponse,omitempty"`
	Request           *IRequest         `json:"request,omitempty"`
	StatusCode        int               `json:"statusCode,omitempty"`
	StatusDescription *string           `json:"statusDescription,omitempty"`
	ContentType       *string           `json:"contentType,omitempty"`
	OutputStream      []byte            `json:"outputStream,omitempty"`
	Dto               *Object           `json:"dto,omitempty"`
	UseBufferedStream bool              `json:"useBufferedStream,omitempty"`
	IsClosed          bool              `json:"isClosed,omitempty"`
	KeepAlive         bool              `json:"keepAlive,omitempty"`
	HasStarted        bool              `json:"hasStarted,omitempty"`
	Items             map[string]Object `json:"items,omitempty"`
}

// @DataContract
type SchedulerTaskListProjection struct {
	// @DataMember
	TaskId string `json:"taskId"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Cron string `json:"cron"`
	// @DataMember
	Type SchedulerTaskType `json:"type,omitempty"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	ViewId string `json:"viewId"`
}

type ResolvedRefStatus string

const (
	ResolvedRefStatusOk           ResolvedRefStatus = "ok"
	ResolvedRefStatusNotFound                       = "notFound"
	ResolvedRefStatusUnauthorized                   = "unauthorized"
	ResolvedRefStatusSourceError                    = "sourceError"
	ResolvedRefStatusErased                         = "erased"
)

type ResolvedResourceEntry struct {
	Ref        ResourceRefDto    `json:"ref"`
	Status     ResolvedRefStatus `json:"status,omitempty"`
	Resolved   *Object           `json:"resolved,omitempty"`
	Diagnostic *string           `json:"diagnostic,omitempty"`
}

// @DataContract
type ChannelSubscriptionStateDto struct {
	// @DataMember
	Unsubscribed bool `json:"unsubscribed,omitempty"`
	// @DataMember
	BlockedTags map[string][]string `json:"blockedTags"`
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

type ConsentPurposeDto struct {
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Channel         string   `json:"channel"`
	MappedTags      []string `json:"mappedTags"`
	RegulatoryBasis []string `json:"regulatoryBasis"`
	Description     *string  `json:"description,omitempty"`
	IsDeprecated    bool     `json:"isDeprecated,omitempty"`
}

type RetentionWindowDto struct {
	DataKind string `json:"dataKind"`
	Days     int    `json:"days,omitempty"`
	Action   string `json:"action"`
}

type ProjectComplianceDto struct {
	Regimes          []string             `json:"regimes"`
	ConsentPurposes  []ConsentPurposeDto  `json:"consentPurposes"`
	RetentionWindows []RetentionWindowDto `json:"retentionWindows"`
}

type LegalHoldDto struct {
	Id          string     `json:"id"`
	SubjectKind string     `json:"subjectKind"`
	SubjectId   string     `json:"subjectId"`
	Reason      string     `json:"reason"`
	PlacedAt    time.Time  `json:"placedAt,omitempty"`
	PlacedBy    *string    `json:"placedBy,omitempty"`
	ReleasedAt  *time.Time `json:"releasedAt,omitempty"`
	ReleasedBy  *string    `json:"releasedBy,omitempty"`
}

type DsarRequestDto struct {
	Id              string     `json:"id"`
	SubjectKind     string     `json:"subjectKind"`
	SubjectId       string     `json:"subjectId"`
	Status          string     `json:"status"`
	ReceivedAt      time.Time  `json:"receivedAt,omitempty"`
	SlaDeadline     time.Time  `json:"slaDeadline,omitempty"`
	AutoApproveAt   *time.Time `json:"autoApproveAt,omitempty"`
	DecidedAt       *time.Time `json:"decidedAt,omitempty"`
	DecidedBy       *string    `json:"decidedBy,omitempty"`
	RejectionReason *string    `json:"rejectionReason,omitempty"`
}

type ComplianceAuditEntryDto struct {
	Id          string            `json:"id"`
	Timestamp   time.Time         `json:"timestamp,omitempty"`
	Action      string            `json:"action"`
	SubjectKind *string           `json:"subjectKind,omitempty"`
	SubjectId   *string           `json:"subjectId,omitempty"`
	Reason      *string           `json:"reason,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

type AccountComplianceDto struct {
	DsarMode              string  `json:"dsarMode"`
	DsarDelayDays         int     `json:"dsarDelayDays,omitempty"`
	AutoForwardAdvisories bool    `json:"autoForwardAdvisories,omitempty"`
	SecurityContact       *string `json:"securityContact,omitempty"`
}

type SupportCustomerStatus string

const (
	SupportCustomerStatusPending SupportCustomerStatus = "Pending"
	SupportCustomerStatusOpen                          = "Open"
	SupportCustomerStatusSolved                        = "Solved"
	SupportCustomerStatusClosed                        = "Closed"
)

// @DataContract
type SupportCaseDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	AccountId string `json:"accountId"`
	// @DataMember
	ProjectId *string `json:"projectId,omitempty"`
	// @DataMember
	ReporterId *string `json:"reporterId,omitempty"`
	// @DataMember
	Kind SupportCaseKind `json:"kind,omitempty"`
	// @DataMember
	Severity SupportCaseSeverity `json:"severity,omitempty"`
	// @DataMember
	Status SupportCaseStatus `json:"status,omitempty"`
	// @DataMember
	CustomerStatus SupportCustomerStatus `json:"customerStatus,omitempty"`
	// @DataMember
	Subject string `json:"subject"`
	// @DataMember
	AffectedModule *string `json:"affectedModule,omitempty"`
	// @DataMember
	DeploymentMode DeploymentMode `json:"deploymentMode,omitempty"`
	// @DataMember
	GatewayVersion *string `json:"gatewayVersion,omitempty"`
	// @DataMember
	Region *string `json:"region,omitempty"`
	// @DataMember
	PlanTier *string `json:"planTier,omitempty"`
	// @DataMember
	OpenedOn int64 `json:"openedOn,omitempty"`
	// @DataMember
	FirstResponseOn *int64 `json:"firstResponseOn,omitempty"`
	// @DataMember
	ResolvedOn *int64 `json:"resolvedOn,omitempty"`
	// @DataMember
	ClosedOn *int64 `json:"closedOn,omitempty"`
	// @DataMember
	Resolution *string `json:"resolution,omitempty"`
	// @DataMember
	MessageCount int `json:"messageCount,omitempty"`
	// @DataMember
	LastMessageOn *int64 `json:"lastMessageOn,omitempty"`
}

// @DataContract
type SupportCaseMessageDto struct {
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	CaseId string `json:"caseId"`
	// @DataMember
	AuthorKind SupportMessageAuthorKind `json:"authorKind,omitempty"`
	// @DataMember
	AuthorId *string `json:"authorId,omitempty"`
	// @DataMember
	AuthorDisplayName *string `json:"authorDisplayName,omitempty"`
	// @DataMember
	Body string `json:"body"`
	// @DataMember
	SentOn int64 `json:"sentOn,omitempty"`
}

// @DataContract
type SupportCaseDetailDto struct {
	// @DataMember
	Case SupportCaseDto `json:"case"`
	// @DataMember
	Messages []SupportCaseMessageDto `json:"messages"`
}

// @DataContract
type SupportCaseListProjection struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	ProjectId *string `json:"projectId,omitempty"`
	// @DataMember
	Kind SupportCaseKind `json:"kind,omitempty"`
	// @DataMember
	Severity SupportCaseSeverity `json:"severity,omitempty"`
	// @DataMember
	Status SupportCaseStatus `json:"status,omitempty"`
	// @DataMember
	CustomerStatus SupportCustomerStatus `json:"customerStatus,omitempty"`
	// @DataMember
	Subject string `json:"subject"`
	// @DataMember
	OpenedOn int64 `json:"openedOn,omitempty"`
	// @DataMember
	MessageCount int `json:"messageCount,omitempty"`
	// @DataMember
	LastMessageOn *int64 `json:"lastMessageOn,omitempty"`
}

// @DataContract
type DiagnosticPackStepDescriptorDto struct {
	// @DataMember
	StepId string `json:"stepId"`
	// @DataMember
	Kind string `json:"kind"`
	// @DataMember
	Description string `json:"description"`
	// @DataMember
	Parameters map[string]string `json:"parameters"`
}

// @DataContract
type DiagnosticPackDescriptorDto struct {
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Version int `json:"version,omitempty"`
	// @DataMember
	Summary string `json:"summary"`
	// @DataMember
	Steps []DiagnosticPackStepDescriptorDto `json:"steps"`
}

// @DataContract
type DiagnosticPackStepResultDto struct {
	// @DataMember
	StepId string `json:"stepId"`
	// @DataMember
	Kind string `json:"kind"`
	// @DataMember
	IsSuccess bool `json:"isSuccess,omitempty"`
	// @DataMember
	Result *Object `json:"result,omitempty"`
	// @DataMember
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// @DataContract
type DiagnosticPackRunResultDto struct {
	// @DataMember
	PackName string `json:"packName"`
	// @DataMember
	PackVersion int `json:"packVersion,omitempty"`
	// @DataMember
	CaseId *string `json:"caseId,omitempty"`
	// @DataMember
	Steps []DiagnosticPackStepResultDto `json:"steps"`
}

// @DataContract
type DiagnosticEchoRegionDto struct {
}

// @DataContract
type DiagnosticEchoDto struct {
	// @DataMember
	ContainerName *string `json:"containerName,omitempty"`
	// @DataMember
	IsManagedService bool `json:"isManagedService,omitempty"`
	// @DataMember
	ApiVersion string `json:"apiVersion"`
	// @DataMember
	HubVersion string `json:"hubVersion"`
	// @DataMember
	Release string `json:"release"`
	// @DataMember
	Runtime string `json:"runtime"`
	// @DataMember
	HubUrl string `json:"hubUrl"`
	// @DataMember
	ApiUrl string `json:"apiUrl"`
	// @DataMember
	LicensePresent bool `json:"licensePresent,omitempty"`
	// @DataMember
	Regions []DiagnosticEchoRegionDto `json:"regions,omitempty"`
}

// @DataContract
type DiagnosticEventItemDto struct {
	// @DataMember
	Position int64 `json:"position,omitempty"`
	// @DataMember
	EventType string `json:"eventType"`
	// @DataMember
	Payload map[string]string `json:"payload,omitempty"`
}

// @DataContract
type DiagnosticEventsPageDto struct {
	// @DataMember
	Stream string `json:"stream"`
	// @DataMember
	From int64 `json:"from,omitempty"`
	// @DataMember
	Count int `json:"count,omitempty"`
	// @DataMember
	HasMore bool `json:"hasMore,omitempty"`
	// @DataMember
	NextFrom int64 `json:"nextFrom,omitempty"`
	// @DataMember
	Items []DiagnosticEventItemDto `json:"items"`
}

// @DataContract
type DiagnosticLogsResponse struct {
	// @DataMember
	List *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type DiagnosticRedisListItemDto struct {
	// @DataMember
	ViewId string `json:"viewId"`
	// @DataMember
	Name *string `json:"name,omitempty"`
	// @DataMember
	Status *string `json:"status,omitempty"`
}

// @DataContract
type DiagnosticRedisInspectDto struct {
	// @DataMember
	KeyPattern string `json:"keyPattern"`
	// @DataMember
	CacheKey string `json:"cacheKey"`
	// @DataMember
	IsList bool `json:"isList,omitempty"`
	// @DataMember
	Item map[string]Object `json:"item,omitempty"`
	// @DataMember
	ListItems []DiagnosticRedisListItemDto `json:"listItems,omitempty"`
	// @DataMember
	HasMore bool `json:"hasMore,omitempty"`
	// @DataMember
	StartingAfter *string `json:"startingAfter,omitempty"`
}

// @DataContract
type DiagnosticHealthCheckDto struct {
	// @DataMember
	CheckId string `json:"checkId"`
	// @DataMember
	IsHealthy bool `json:"isHealthy,omitempty"`
	// @DataMember
	StatusCode int `json:"statusCode,omitempty"`
	// @DataMember
	Detail *string `json:"detail,omitempty"`
}

// @DataContract
type ResponseError struct {
	// @DataMember(Order=1)
	ErrorCode string `json:"errorCode"`
	// @DataMember(Order=2)
	FieldName string `json:"fieldName"`
	// @DataMember(Order=3)
	Message string `json:"message"`
	// @DataMember(Order=4)
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type ResponseStatus struct {
	// @DataMember(Order=1)
	ErrorCode string `json:"errorCode"`
	// @DataMember(Order=2)
	Message *string `json:"message,omitempty"`
	// @DataMember(Order=3)
	StackTrace *string `json:"stackTrace,omitempty"`
	// @DataMember(Order=4)
	Errors []ResponseError `json:"errors,omitempty"`
	// @DataMember(Order=5)
	Meta map[string]string `json:"meta,omitempty"`
}

type ILlmApiKeyRequest struct {
	ApiKey string `json:"apiKey"`
}

type IHasViewId struct {
	ViewId string `json:"viewId"`
}

type IHasDatabaseId struct {
	Id *string `json:"id,omitempty"`
}

type IBindableContract struct {
}

type IHasRazorTemplateCode struct {
}

type IHasDomainEntityId struct {
	ViewId string `json:"viewId"`
}

type IIntegrationIdentification struct {
	IntegrationId IntegrationId `json:"integrationId"`
	Capability    string        `json:"capability,omitempty"`
	IsSystemOwned bool          `json:"isSystemOwned,omitempty"`
}

type CronExpression struct {
	Value  string         `json:"value"`
	Parsed CronExpression `json:"parsed"`
}

type IHasResponsibleUserId struct {
	UserId string `json:"userId"`
}

type ICursorArgs struct {
	Field string `json:"field"`
	Order int    `json:"order,omitempty"`
}

type StringField struct {
	JsonSchemaField
	Format           *string              `json:"format,omitempty"`
	Pattern          *string              `json:"pattern,omitempty"`
	MinLength        *int                 `json:"minLength,omitempty"`
	MaxLength        *int                 `json:"maxLength,omitempty"`
	TranslateOptions *IReadOnlyDictionary `json:"translateOptions,omitempty"`
}

type DecimalField struct {
	JsonSchemaField
	Minimum    *float64 `json:"minimum,omitempty"`
	Maximum    *float64 `json:"maximum,omitempty"`
	MultipleOf *float64 `json:"multipleOf,omitempty"`
}

type CurrencyField struct {
	JsonSchemaField
	AllowedCurrencies *IReadOnlyList `json:"allowedCurrencies,omitempty"`
}

type BooleanField struct {
	JsonSchemaField
}

type DateField struct {
	JsonSchemaField
	Minimum *int64 `json:"minimum,omitempty"`
	Maximum *int64 `json:"maximum,omitempty"`
}

type IntegerField struct {
	JsonSchemaField
	Minimum *int64 `json:"minimum,omitempty"`
	Maximum *int64 `json:"maximum,omitempty"`
}

type GeolocationField struct {
	JsonSchemaField
	AllowedTypes *IReadOnlyList `json:"allowedTypes,omitempty"`
}

type TagsField struct {
	JsonSchemaField
}

type FileField struct {
	JsonSchemaField
	Storages *IReadOnlyList `json:"storages,omitempty"`
}

type TaxonomySelectionField struct {
	JsonSchemaField
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	Multiple   bool    `json:"multiple,omitempty"`
}

type CollectionSelectionField struct {
	JsonSchemaField
	CollectionId *string `json:"collectionId,omitempty"`
	DisplayField *string `json:"displayField,omitempty"`
	Multiple     bool    `json:"multiple,omitempty"`
}

type UserSelectionField struct {
	JsonSchemaField
	Multiple bool `json:"multiple,omitempty"`
}

type RoleSelectionField struct {
	JsonSchemaField
	Multiple bool `json:"multiple,omitempty"`
}

type EnumSelectionField struct {
	JsonSchemaField
	Values   *IReadOnlyList `json:"values,omitempty"`
	Multiple bool           `json:"multiple,omitempty"`
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
	TranslateOptions *IReadOnlyDictionary `json:"translateOptions,omitempty"`
}

type DecimalFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *float64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *float64 `json:"maximum,omitempty"`
	// @DataMember
	MultipleOf *float64 `json:"multipleOf,omitempty"`
}

type CurrencyFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	AllowedCurrencies *IReadOnlyList `json:"allowedCurrencies,omitempty"`
}

type BooleanFieldDto struct {
	JsonSchemaFieldDto
}

type DateFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *int64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *int64 `json:"maximum,omitempty"`
}

type IntegerFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Minimum *int64 `json:"minimum,omitempty"`
	// @DataMember
	Maximum *int64 `json:"maximum,omitempty"`
}

type GeolocationFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	AllowedTypes *IReadOnlyList `json:"allowedTypes,omitempty"`
}

type TagsFieldDto struct {
	JsonSchemaFieldDto
}

type FileFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Storages *IReadOnlyList `json:"storages,omitempty"`
}

type TaxonomySelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	TaxonomyId *string `json:"taxonomyId,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
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
}

type RoleSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
}

type EnumSelectionFieldDto struct {
	JsonSchemaFieldDto
	// @DataMember
	Values *IReadOnlyList `json:"values,omitempty"`
	// @DataMember
	Multiple bool `json:"multiple,omitempty"`
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
}

type PublicProjectConfigDto struct {
	DisplayName        string          `json:"displayName"`
	AdminPortalEnabled bool            `json:"adminPortalEnabled,omitempty"`
	Branding           *PublicBrandDto `json:"branding,omitempty"`
	Auth               PublicAuthDto   `json:"auth"`
}

type PublicLegalDocumentDto struct {
	Kind      string  `json:"kind"`
	Title     *string `json:"title,omitempty"`
	Body      string  `json:"body"`
	Available bool    `json:"available,omitempty"`
}

type GetAccountProfileResponse struct {
	ResponseBase
	Item *AccountOwnerDto `json:"item,omitempty"`
}

type EmptyResponse struct {
	ResponseBase
}

type GetAccountStatusResponse struct {
	ResponseBase
	Item *AccountStatusDto `json:"item,omitempty"`
}

// @DataContract
type IdResponse struct {
	ResponseBase
	// @DataMember
	Id *string `json:"id,omitempty"`
	// @DataMember
	Status *string `json:"status,omitempty"`
}

// @DataContract
type CreateStripeCheckoutSessionResponse struct {
	IdResponse
}

// @DataContract
type GetStripeBillingPortalUrlResponse struct {
	IdResponse
}

// @DataContract
type CreateTeamMemberFromInvitationResponse struct {
	IdResponse
	// @DataMember
	Token *string `json:"token,omitempty"`
}

type GetAccountUsageBillingResponse struct {
	ResponseBase
	Item *UsageBillingDto `json:"item,omitempty"`
}

type PromoteEnvironmentResponse struct {
	ResponseBase
	Item *PromotionResultDto `json:"item,omitempty"`
}

type GetProjectEnvironmentsResponse struct {
	ResponseBase
	Item *ProjectEnvironmentsDto `json:"item,omitempty"`
}

type GetProjectResponse struct {
	ResponseBase
	Item *ProjectDto `json:"item,omitempty"`
}

type GetProjectsResponse struct {
	ResponseBase
	List []ProjectListItemDto `json:"list,omitempty"`
}

type GetAccountRegionsResponse struct {
	ResponseBase
	Items []ProjectRegionDto `json:"items,omitempty"`
}

type WaitForProjectActiveResponse struct {
	ResponseBase
	Status        *string `json:"status,omitempty"`
	IsActive      bool    `json:"isActive,omitempty"`
	WaitedSeconds int     `json:"waitedSeconds,omitempty"`
	Message       *string `json:"message,omitempty"`
}

type GetProjectTokensResponse struct {
	ResponseBase
	Tokens []TokenMappingDto `json:"tokens,omitempty"`
}

type AdminPortalStructureDto struct {
	ProjectId          string                 `json:"projectId"`
	AdminPortalEnabled bool                   `json:"adminPortalEnabled,omitempty"`
	DisplayName        string                 `json:"displayName"`
	Modules            []AdminPortalModuleDto `json:"modules"`
}

// @DataContract
type CreateAccountResponse struct {
	IdResponse
	// @DataMember
	Token *string `json:"token,omitempty"`
}

type GetAccountCollaboratorsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetAccountPasswordPolicyResponse struct {
	ResponseBase
	Policy *AccountPasswordPolicyDto `json:"policy,omitempty"`
}

type GetAccountTeamPoliciesResponse struct {
	ResponseBase
	Policies []PolicyItemDto `json:"policies,omitempty"`
}

type GetAccountTeamRolesResponse struct {
	ResponseBase
	Roles []AccountTeamRoleDto `json:"roles,omitempty"`
}

type AccountPasskeyOkResponse struct {
	ResponseBase
}

type AccountPasskeyVerificationTokenResponse struct {
	ResponseBase
	VerificationToken string `json:"verificationToken"`
}

type AccountPasskeyCeremonyOptionsResponse struct {
	ResponseBase
	CeremonyId  string `json:"ceremonyId"`
	OptionsJson string `json:"optionsJson"`
}

type AccountPasskeyAuthTokensResponse struct {
	ResponseBase
	AccessToken      string   `json:"accessToken"`
	RefreshToken     string   `json:"refreshToken"`
	ExpiresInSeconds int      `json:"expiresInSeconds,omitempty"`
	RecoveryCodes    []string `json:"recoveryCodes,omitempty"`
}

type AccountPasskeyListResponse struct {
	ResponseBase
	Passkeys []AccountPasskeyListItemDto `json:"passkeys"`
}

type AccountPasskeyEnrollmentResponse struct {
	ResponseBase
	RecoveryCodes []string `json:"recoveryCodes,omitempty"`
}

type GetLicenseDomainDnsStatusResponse struct {
	ResponseBase
	Status *LicenseDomainDnsStatusDto `json:"status,omitempty"`
}

type StartLicenseDomainVerificationResponse struct {
	ResponseBase
	Challenge *LicenseDomainVerificationChallengeDto `json:"challenge,omitempty"`
}

type GetLicenseDomainVerificationStatusResponse struct {
	ResponseBase
	Status *LicenseDomainVerificationStatusDto `json:"status,omitempty"`
}

type GetLicensesResponse struct {
	ResponseBase
	List []LicenseDto `json:"list,omitempty"`
}

type PostLicenseHeartbeatResponse struct {
	ResponseBase
	Verdict *LicenseHeartbeatVerdictDto `json:"verdict,omitempty"`
}

type GetInstallationLicenseStatusResponse struct {
	ResponseBase
	Status *InstallationLicenseStatusDto `json:"status,omitempty"`
}

type IssueServiceUserApiKeyResponse struct {
	Id   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type ListServiceUserApiKeysResponse struct {
	Keys []ServiceUserApiKeyDto `json:"keys"`
}

type GetMembershipTriggerResponse struct {
	GetTriggerResponse
	Trigger *MembershipTriggerDto `json:"trigger,omitempty"`
}

// @DataContract
type GetMembershipTriggersResponse struct {
	GetTriggersResponse
	// @DataMember
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetRoleResponse struct {
	ResponseBase
	Role *RoleItemDto `json:"role,omitempty"`
}

type GetRolesResponse struct {
	ResponseBase
	Roles []RoleListProjectionDto `json:"roles"`
}

type GetPolicyResponse struct {
	ResponseBase
	Policy *PolicyItemDto `json:"policy,omitempty"`
}

type GetPoliciesResponse struct {
	ResponseBase
	Policies []PolicyItemDto `json:"policies"`
}

type GetPasskeySettingsResponse struct {
	ResponseBase
	Result *PasskeySettingsDto `json:"result,omitempty"`
}

type GetMembershipIntegrationResponse struct {
	ResponseBase
	Item *MembershipIntegrationDto `json:"item,omitempty"`
}

type GetMembershipIntegrationsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetAuthorizationSettingsResponse struct {
	ResponseBase
	Result *MembershipAuthorizationViewDto `json:"result,omitempty"`
}

type UpdatePasswordComplexityResponse struct {
	ResponseBase
	Result bool `json:"result,omitempty"`
}

type GetAuthenticationSettingsResponse struct {
	ResponseBase
	Result *MembershipAuthenticationViewDto `json:"result,omitempty"`
}

type GetSchemaTriggerResponse struct {
	GetTriggerResponse
	Trigger *SchemaTriggerDto `json:"trigger,omitempty"`
}

type GetSchemaTriggersResponse struct {
	GetTriggersResponse
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetDatabaseTaxonomyResponse struct {
	ResponseBase
	Item *TaxonomyDto `json:"item,omitempty"`
}

type GetDatabaseTaxonomiesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetDatabaseTaxonomyTreeResponse struct {
	ResponseBase
	Tree []TaxonomyTreeDto `json:"tree,omitempty"`
}

type GetDatabaseTaxonomyTermResponse struct {
	ResponseBase
	Item *TermDto `json:"item,omitempty"`
}

type GetDatabaseMergedTermTreeResponse struct {
	ResponseBase
	Tree []TermTreeDto `json:"tree,omitempty"`
}

type GetDatabaseTaxonomyTermTreeResponse struct {
	ResponseBase
	Tree []TermTreeDto `json:"tree,omitempty"`
}

// @DataContract
type ApplyDatabaseSchemaBundleResponse struct {
	ResponseBase
	// @DataMember
	Tier *string `json:"tier,omitempty"`
	// @DataMember
	Taxonomies []AppliedTaxonomyDto `json:"taxonomies"`
	// @DataMember
	Collections []AppliedCollectionDto `json:"collections"`
	// @DataMember
	Decisions []string `json:"decisions"`
	// @DataMember
	Errors []string `json:"errors"`
}

type GetDatabaseSchemaResponse struct {
	ResponseBase
	Item *SchemaDto `json:"item,omitempty"`
}

type GetDatabaseSchemasResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetDatabaseSchemaDraftResponse struct {
	ResponseBase
	Item *SchemaDraftDto `json:"item,omitempty"`
}

type GetDatabaseSchemaListSettingsResponse struct {
	ResponseBase
	Settings *SchemaListSettingsDto `json:"settings,omitempty"`
}

type GetDatabaseSchemaVersionDiffResponse struct {
	ResponseBase
	Item *SchemaDiffDto `json:"item,omitempty"`
}

type GetDatabaseSchemaVersionsResponse struct {
	ResponseBase
	Items []SchemaVersionSummaryDto `json:"items,omitempty"`
}

type AggregateRecordsResponse struct {
	ResponseBase
	Result []Object `json:"result,omitempty"`
}

type CountRecordsResponse struct {
	ResponseBase
	Count int64 `json:"count,omitempty"`
}

type DistinctRecordValuesResponse struct {
	ResponseBase
	Values []Object `json:"values,omitempty"`
}

type ExecuteRecordsAggregateResponse struct {
	ResponseBase
	Result []Object `json:"result,omitempty"`
}

type FindRecordsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type FindOneRecordResponse struct {
	ResponseBase
	Result *Object `json:"result,omitempty"`
}

type GetCollectionIndexesResponse struct {
	ResponseBase
	Indexes []CollectionIndexDto `json:"indexes,omitempty"`
}

// @DataContract
type SeedCollectionRecordsResponse struct {
	ResponseBase
	// @DataMember
	Result *SeedCollectionRecordsResultDto `json:"result,omitempty"`
}

type GetDatabaseIntegrationResponse struct {
	ResponseBase
	Item *DatabaseIntegrationDto `json:"item,omitempty"`
}

type GetDatabaseIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

type GetAllowedFlexTiersResponse struct {
	ResponseBase
	Tiers []FlexTierDto `json:"tiers,omitempty"`
}

type RevealManagedFlexConnectionStringResponse struct {
	ResponseBase
	ConnectionString *string `json:"connectionString,omitempty"`
}

// @DataContract
type TestDatabaseIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetCollectionImportResponse struct {
	ResponseBase
	Result *CollectionImportDto `json:"result,omitempty"`
}

type GetCollectionImportsResponse struct {
	ResponseBase
	Result *PaginatedResponse `json:"result,omitempty"`
}

type RequestImportUploadUrlResponse struct {
	ResponseBase
	Result *ImportUploadTargetDto `json:"result,omitempty"`
}

type AnalyzeImportFileResponse struct {
	ResponseBase
	Result *ImportFileAnalysisDto `json:"result,omitempty"`
}

type GetDatabaseAggregateResponse struct {
	ResponseBase
	Item *MongoDbAggregateDto `json:"item,omitempty"`
}

type GetDatabaseAggregatesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type TestDatabaseAggregateResponse struct {
	ResponseBase
	Result []Object `json:"result,omitempty"`
}

type GetFilesTriggerResponse struct {
	GetTriggerResponse
	Trigger *FilesTriggerDto `json:"trigger,omitempty"`
}

type GetFilesTriggersResponse struct {
	GetTriggersResponse
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetFilesIntegrationResponse struct {
	ResponseBase
	Item *FilesIntegrationDto `json:"item,omitempty"`
}

type GetFilesIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestFilesIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetFileResponse struct {
	ResponseBase
	File      *FileResourceRefDto `json:"file,omitempty"`
	IsPublic  *bool               `json:"isPublic,omitempty"`
	PublicUrl *string             `json:"publicUrl,omitempty"`
}

type GetFolderFilesResponse struct {
	ResponseBase
	List    *PaginatedResponse `json:"list,omitempty"`
	Folders *IList             `json:"folders,omitempty"`
}

type GetNotificationModuleDisableDependenciesResponse struct {
	ResponseBase
	Dependencies *NotificationModuleDisableDependenciesDto `json:"dependencies,omitempty"`
}

// @DataContract
type TestEmailValidationIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items []TestEmailValidationItemDto `json:"items"`
}

type GetEmailTemplateResponse struct {
	ResponseBase
	Item *EmailTemplateDto `json:"item,omitempty"`
}

type GetEmailTemplatesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetHtmlFromMjmlResponse struct {
	ResponseBase
	Variables            []string              `json:"variables,omitempty"`
	HtmlFromMjmlResponse *HtmlFromMjmlResponse `json:"htmlFromMjmlResponse,omitempty"`
}

type GetSystemEmailTemplateResponse struct {
	ResponseBase
	Item *SystemEmailTemplateDto `json:"item,omitempty"`
}

type GetSystemEmailTemplatesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetEmailTemplateAvailableTokensResponse struct {
	ResponseBase
	Tokens map[string][]string `json:"tokens,omitempty"`
}

type GetEmailSignatureResponse struct {
	ResponseBase
	Item *EmailSignatureDto `json:"item,omitempty"`
}

type GetEmailSignaturesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetEmailSettingsResponse struct {
	ResponseBase
	Settings   *EmailSettings       `json:"settings,omitempty"`
	SystemTags []GroupDefinitionDto `json:"systemTags,omitempty"`
}

// @DataContract
type CheckEmailIntegrationDomainHealthResponse struct {
	ResponseBase
	// @DataMember
	Domain *string `json:"domain,omitempty"`
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetEmailIntegrationResponse struct {
	ResponseBase
	Item *EmailIntegrationDto `json:"item,omitempty"`
}

type GetEmailIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestEmailIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetEmailFooterResponse struct {
	ResponseBase
	Item *EmailFooterDto `json:"item,omitempty"`
}

type GetEmailFootersResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetEmailCampaignResponse struct {
	ResponseBase
	Item *EmailCampaignDto `json:"item,omitempty"`
}

type GetEmailCampaignsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetEmailCampaignBatchesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetEmailCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *EmailCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

type GetEmailCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []BatchStatusChangeEntryDto `json:"batchStatusHistory,omitempty"`
	List               *PaginatedResponse          `json:"list,omitempty"`
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

type GetEmailCampaignMessageResponse struct {
	ResponseBase
	EmailMessageEntity *EmailCampaignBatchNotificationDto `json:"emailMessageEntity,omitempty"`
}

type GetEmailCampaignMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetSmsTemplateResponse struct {
	ResponseBase
	Item *SmsTemplateDto `json:"item,omitempty"`
}

type GetSmsTemplatesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetSmsMessageContentTokensResponse struct {
	ResponseBase
	Tokens map[string][]string `json:"tokens,omitempty"`
}

type RenderSmsTextResponse struct {
	ResponseBase
	Variables []string `json:"variables,omitempty"`
	Text      *string  `json:"text,omitempty"`
}

type GetSmsSettingsResponse struct {
	ResponseBase
	Settings *SmsSettings `json:"settings,omitempty"`
}

type GetSmsIntegrationResponse struct {
	ResponseBase
	Item *SmsIntegrationDto `json:"item,omitempty"`
}

type GetSmsIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestSmsIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetSmsCampaignResponse struct {
	ResponseBase
	SmsCampaign *SmsCampaignDto `json:"smsCampaign,omitempty"`
}

type GetSmsCampaignsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetSmsCampaignBatchesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetSmsCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *SmsCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

type GetSmsCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []BatchStatusChangeEntryDto `json:"batchStatusHistory,omitempty"`
	List               *PaginatedResponse          `json:"list,omitempty"`
}

type GetSmsCampaignStatisticsResponse struct {
	ResponseBase
	Stats *CampaignStatsDto `json:"stats,omitempty"`
}

type PreviewSmsNotificationResponse struct {
	ResponseBase
	Body *string `json:"body,omitempty"`
}

type GetSmsCampaignMessageResponse struct {
	ResponseBase
	SmsMessageEntity *SmsCampaignBatchNotificationDto `json:"smsMessageEntity,omitempty"`
}

type GetSmsCampaignMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetMarketplaceListingResponse struct {
	ResponseBase
	Listing *MarketplaceListingDto `json:"listing,omitempty"`
}

type GetMarketplaceTokensResponse struct {
	ResponseBase
	Tokens []string `json:"tokens,omitempty"`
}

type GetMarketplaceListingsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetMarketplaceIntegrationResponse struct {
	ResponseBase
	Integration *MarketplaceIntegrationDto `json:"integration,omitempty"`
}

type GetMarketplaceIntegrationsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type EmptyMarketplaceSecretsResponse struct {
	ResponseBase
}

type RevealMarketplaceIntegrationSecretsResponse struct {
	ResponseBase
	Secrets *IReadOnlyDictionary `json:"secrets,omitempty"`
}

// @DataContract
type TestMarketplaceIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type SetMarketplaceIntegrationTokenMappingsResponse struct {
	ResponseBase
}

type GetMarketplaceFunctionResponse struct {
	ResponseBase
	Function *MarketplaceFunctionDto `json:"function,omitempty"`
}

type GetMarketplaceFunctionsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetMarketplaceFunctionCatalogResponse struct {
	ResponseBase
	Functions IReadOnlyList `json:"functions"`
}

type InvokeMarketplaceFunctionResponse struct {
	ResponseBase
	IsSuccess       bool    `json:"isSuccess,omitempty"`
	Output          *Object `json:"output,omitempty"`
	VendorRequestId *string `json:"vendorRequestId,omitempty"`
}

type GetCodeIntegrationResponse struct {
	ResponseBase
	Item *CodeIntegrationDto `json:"item,omitempty"`
}

type GetCodeIntegrationsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestCodeIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetPushTemplateResponse struct {
	ResponseBase
	Item *PushTemplateDto `json:"item,omitempty"`
}

type GetPushTemplatesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetPushMessageContentTokensResponse struct {
	ResponseBase
	Tokens map[string][]string `json:"tokens,omitempty"`
}

type RenderPushResponse struct {
	ResponseBase
	Variables []string `json:"variables,omitempty"`
	Title     *string  `json:"title,omitempty"`
	Body      *string  `json:"body,omitempty"`
	Subtitle  *string  `json:"subtitle,omitempty"`
}

type GetPushSettingsResponse struct {
	ResponseBase
	Settings *PushSettings `json:"settings,omitempty"`
}

type GetPushIntegrationResponse struct {
	ResponseBase
	Item *PushIntegrationDto `json:"item,omitempty"`
}

type GetPushIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

type GetPushCampaignResponse struct {
	ResponseBase
	Item *PushCampaignDto `json:"item,omitempty"`
}

type GetPushCampaignsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetPushCampaignBatchesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetPushCampaignBatchNotificationResponse struct {
	ResponseBase
	CampaignNotification *PushCampaignBatchNotificationDto `json:"campaignNotification,omitempty"`
}

type GetPushCampaignBatchNotificationsResponse struct {
	ResponseBase
	BatchStatusHistory []BatchStatusChangeEntryDto `json:"batchStatusHistory,omitempty"`
	List               *PaginatedResponse          `json:"list,omitempty"`
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

type GetPushCampaignMessageResponse struct {
	ResponseBase
	PushMessageEntity *PushCampaignBatchNotificationDto `json:"pushMessageEntity,omitempty"`
}

type GetPushCampaignMessagesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetPaymentsWebhookLogResponse struct {
	ResponseBase
	List *IReadOnlyList `json:"list,omitempty"`
}

type GetPaymentsTriggerResponse struct {
	GetTriggerResponse
	Trigger *PaymentTriggerDto `json:"trigger,omitempty"`
}

type GetPaymentsTriggersResponse struct {
	GetTriggersResponse
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetPaymentsIntegrationResponse struct {
	ResponseBase
	Item *PaymentsIntegrationDto `json:"item,omitempty"`
}

type GetPaymentsIntegrationsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestPaymentsIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetLoggingIntegrationResponse struct {
	ResponseBase
	Item *LoggingIntegrationDto `json:"item,omitempty"`
}

type GetLoggingIntegrationsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestLoggingIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type CleanLogsResponse struct {
	ResponseBase
}

type GetLogsByCorrelationIdResponse struct {
	ResponseBase
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetLogsResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetLogSettingsResponse struct {
	ResponseBase
	SkipCloudDashboardLogs bool `json:"skipCloudDashboardLogs,omitempty"`
	SkipHttpBodyMeta       bool `json:"skipHttpBodyMeta,omitempty"`
	AiChatLoggingEnabled   bool `json:"aiChatLoggingEnabled,omitempty"`
	HasNorbixLogging       bool `json:"hasNorbixLogging,omitempty"`
}

type SaveLogSettingsResponse struct {
	ResponseBase
}

type GetAiToolsResponse struct {
	ResponseBase
	Tools []AiToolManifestItem `json:"tools,omitempty"`
}

type InvokeAiToolResponse struct {
	ResponseBase
	Result *string `json:"result,omitempty"`
}

type AskChatResponse struct {
	ResponseBase
	Result *string `json:"result,omitempty"`
}

type UploadChatAttachmentResponse struct {
	ResponseBase
	Id        *string `json:"id,omitempty"`
	SessionId *string `json:"sessionId,omitempty"`
}

type ChatAvailabilityResponse struct {
	ResponseBase
	Available bool              `json:"available,omitempty"`
	Reason    *string           `json:"reason,omitempty"`
	Profiles  []string          `json:"profiles,omitempty"`
	Models    []ChatModelOption `json:"models,omitempty"`
}

type GetChatMemoryResponse struct {
	ResponseBase
	Notes []ChatMemoryNote `json:"notes,omitempty"`
}

type GetChatSessionsResponse struct {
	ResponseBase
	Sessions []ChatSessionListItem `json:"sessions,omitempty"`
}

type GetChatSessionEntriesResponse struct {
	ResponseBase
	SessionId         *string              `json:"sessionId,omitempty"`
	Profile           *string              `json:"profile,omitempty"`
	ProjectId         *string              `json:"projectId,omitempty"`
	Env               *string              `json:"env,omitempty"`
	Entries           []AiChatEntryWireDto `json:"entries,omitempty"`
	LastSeq           int64                `json:"lastSeq,omitempty"`
	ActiveWorkItemIds []string             `json:"activeWorkItemIds,omitempty"`
}

type ChatTurnResponse struct {
	ResponseBase
	SessionId   *string               `json:"sessionId,omitempty"`
	Reply       *string               `json:"reply,omitempty"`
	ScreenPatch *ChatScreenContextDto `json:"screenPatch,omitempty"`
	ToolTrace   []string              `json:"toolTrace,omitempty"`
}

type GetProjectBriefResponse struct {
	ResponseBase
	ProjectId *string                      `json:"projectId,omitempty"`
	Snapshot  *ProjectBriefSnapshotWireDto `json:"snapshot,omitempty"`
	Events    []ProjectBriefEventWireDto   `json:"events,omitempty"`
	LastSeq   int64                        `json:"lastSeq,omitempty"`
}

type GetWorkItemsResponse struct {
	ResponseBase
	ProjectId *string           `json:"projectId,omitempty"`
	WorkItems []WorkItemWireDto `json:"workItems,omitempty"`
}

type GetWorkItemResponse struct {
	ResponseBase
	WorkItem *WorkItemWireDto     `json:"workItem,omitempty"`
	Plans    []AiChatEntryWireDto `json:"plans,omitempty"`
	Steps    []AiChatEntryWireDto `json:"steps,omitempty"`
}

type ExportWorkItemResponse struct {
	ResponseBase
	WorkItemId *string `json:"workItemId,omitempty"`
	Markdown   *string `json:"markdown,omitempty"`
}

type GetLlmIntegrationResponse struct {
	ResponseBase
	Item *LlmIntegrationDto `json:"item,omitempty"`
}

type GetLlmIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

// @DataContract
type TestLlmIntegrationResponse struct {
	ResponseBase
	// @DataMember
	Items *IReadOnlyList `json:"items,omitempty"`
}

type GetMcpIntegrationResponse struct {
	ResponseBase
	Item *McpIntegrationDto `json:"item,omitempty"`
}

type GetMcpIntegrationsResponse struct {
	ResponseBase
	DefaultIntegrationId *string            `json:"defaultIntegrationId,omitempty"`
	List                 *PaginatedResponse `json:"list,omitempty"`
}

type GetWebhookIntegrationResponse struct {
	ResponseBase
	Item *WebhookIntegrationDto `json:"item,omitempty"`
}

type RevealWebhookIntegrationSecretResponse struct {
	ResponseBase
	SigningSecret *string `json:"signingSecret,omitempty"`
}

type RotateWebhookIntegrationSecretResponse struct {
	ResponseBase
	SigningSecret *string `json:"signingSecret,omitempty"`
}

type HttpResult struct {
	ResponseText          *string             `json:"responseText,omitempty"`
	ResponseStream        []byte              `json:"responseStream,omitempty"`
	FileInfo              *FileInfo           `json:"fileInfo,omitempty"`
	VirtualFile           *IVirtualFile       `json:"virtualFile,omitempty"`
	ContentType           *string             `json:"contentType,omitempty"`
	Headers               map[string]string   `json:"headers,omitempty"`
	Cookies               []Cookie            `json:"cookies,omitempty"`
	ETag                  *string             `json:"eTag,omitempty"`
	Age                   *time.Duration      `json:"age,omitempty"`
	MaxAge                *time.Duration      `json:"maxAge,omitempty"`
	Expires               *time.Time          `json:"expires,omitempty"`
	LastModified          *time.Time          `json:"lastModified,omitempty"`
	CacheControl          CacheControl        `json:"cacheControl,omitempty"`
	ResultScope           *Func               `json:"resultScope,omitempty"`
	AllowsPartialResponse bool                `json:"allowsPartialResponse,omitempty"`
	Options               map[string]string   `json:"options,omitempty"`
	Status                int                 `json:"status,omitempty"`
	StatusCode            HttpStatusCode      `json:"statusCode,omitempty"`
	StatusDescription     *string             `json:"statusDescription,omitempty"`
	Response              *Object             `json:"response,omitempty"`
	ResponseFilter        *IContentTypeWriter `json:"responseFilter,omitempty"`
	RequestContext        *IRequest           `json:"requestContext,omitempty"`
	View                  *string             `json:"view,omitempty"`
	Template              *string             `json:"template,omitempty"`
	PaddingLength         int                 `json:"paddingLength,omitempty"`
	IsPartialRequest      bool                `json:"isPartialRequest,omitempty"`
}

type SaveWebhookDestinationResponse struct {
	ResponseBase
	DestinationId *string `json:"destinationId,omitempty"`
}

type GetSchedulerTaskResponse struct {
	ResponseBase
	Item *SchedulerTaskDto `json:"item,omitempty"`
}

type GetSchedulerTasksResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type ResolveResourcesResponse struct {
	ResponseBase
	Resolved IReadOnlyList `json:"resolved"`
}

type GetContactResponse struct {
	ResponseBase
	Item *UserDto `json:"item,omitempty"`
}

type GetAllContactsResponse struct {
	ResponseBase
	Items      *IReadOnlyList `json:"items,omitempty"`
	NextCursor *string        `json:"nextCursor,omitempty"`
}

type GetComplianceSettingsResponse struct {
	ResponseBase
	Settings *ProjectComplianceDto `json:"settings,omitempty"`
}

type GetLegalHoldsResponse struct {
	ResponseBase
	Holds []LegalHoldDto `json:"holds"`
}

type GetDsarRequestsResponse struct {
	ResponseBase
	Requests []DsarRequestDto `json:"requests"`
}

type GetComplianceAuditLogResponse struct {
	ResponseBase
	Entries []ComplianceAuditEntryDto `json:"entries"`
}

type GetAccountComplianceResponse struct {
	ResponseBase
	Settings *AccountComplianceDto `json:"settings,omitempty"`
}

type GetSupportCaseResponse struct {
	ResponseBase
	Result *SupportCaseDetailDto `json:"result,omitempty"`
}

type GetSupportCasesResponse struct {
	ResponseBase
	List *PaginatedResponse `json:"list,omitempty"`
}

type GetDiagnosticPacksResponse struct {
	ResponseBase
	Packs []DiagnosticPackDescriptorDto `json:"packs,omitempty"`
}

type RunDiagnosticPackResponse struct {
	ResponseBase
	Result *DiagnosticPackRunResultDto `json:"result,omitempty"`
}

type GetDiagnosticEchoResponse struct {
	ResponseBase
	Result *DiagnosticEchoDto `json:"result,omitempty"`
}

type ReadDiagnosticEventsResponse struct {
	ResponseBase
	Result *DiagnosticEventsPageDto `json:"result,omitempty"`
}

type QueryDiagnosticLogsResponse struct {
	ResponseBase
	Result *DiagnosticLogsResponse `json:"result,omitempty"`
}

type InspectDiagnosticRedisResponse struct {
	ResponseBase
	Result *DiagnosticRedisInspectDto `json:"result,omitempty"`
}

type RunDiagnosticHealthCheckResponse struct {
	ResponseBase
	Result *DiagnosticHealthCheckDto `json:"result,omitempty"`
}

// @DataContract
type AuthenticateResponse struct {
	// @DataMember(Order=1)
	UserId *string `json:"userId,omitempty"`
	// @DataMember(Order=2)
	SessionId *string `json:"sessionId,omitempty"`
	// @DataMember(Order=3)
	UserName *string `json:"userName,omitempty"`
	// @DataMember(Order=4)
	DisplayName *string `json:"displayName,omitempty"`
	// @DataMember(Order=5)
	ReferrerUrl *string `json:"referrerUrl,omitempty"`
	// @DataMember(Order=6)
	BearerToken *string `json:"bearerToken,omitempty"`
	// @DataMember(Order=7)
	RefreshToken *string `json:"refreshToken,omitempty"`
	// @DataMember(Order=8)
	RefreshTokenExpiry *time.Time `json:"refreshTokenExpiry,omitempty"`
	// @DataMember(Order=9)
	ProfileUrl *string `json:"profileUrl,omitempty"`
	// @DataMember(Order=10)
	Roles []string `json:"roles,omitempty"`
	// @DataMember(Order=11)
	Permissions []string `json:"permissions,omitempty"`
	// @DataMember(Order=12)
	AuthProvider *string `json:"authProvider,omitempty"`
	// @DataMember(Order=13)
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
	// @DataMember(Order=14)
	Meta map[string]string `json:"meta,omitempty"`
}

// @DataContract
type GetAccessTokenResponse struct {
	// @DataMember(Order=1)
	AccessToken *string `json:"accessToken,omitempty"`
	// @DataMember(Order=2)
	Meta map[string]string `json:"meta,omitempty"`
	// @DataMember(Order=3)
	ResponseStatus *ResponseStatus `json:"responseStatus,omitempty"`
}

// @Route("/{version}/code/enable", "GET")
type EnableCode struct {
	CodeMashRequestBase
}

// @Route("/{version}/code/disable", "GET")
type DisableCode struct {
	CodeMashRequestBase
}

// @Route("/{version}/code/integrations", "GET")
type GetCodeIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/code/integrations/{id}", "GET")
type GetCodeIntegration struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/code/integrations", "POST")
// @DataContract
type SaveCodeIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration CodeIntegrationRequest `json:"integration"`
}

// @Route("/{version}/code/integrations/test", "POST")
type TestCodeIntegration struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

// @Route("/{version}/code/integrations/confirm-human-delivery", "POST")
// @DataContract
type ConfirmCodeIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @DataMember
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

// @Route("/{version}/code/integrations/{Id}/default", "PUT")
type SetCodeIntegrationAsDefault struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/code/integrations/{Id}", "DELETE")
type DeleteCodeIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/code/integrations/{Id}/enable", "PUT")
type EnableCodeIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/code/integrations/{Id}/disable", "PUT")
type DisableCodeIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_code_integrations. */
	// @ApiMember(Description="Integration id, from get_code_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/code/marketplace/listings", "GET")
type GetMarketplaceListings struct {
	CodeMashListPaginationRequestBase
	/** @description Filter by one or more categories (Crm, Erp, Communication, etc.). */
	// @ApiMember(Description="Filter by one or more categories (Crm, Erp, Communication, etc.).")
	Categories []MarketplaceCategory `json:"categories,omitempty"`
	/** @description Filter by transport (Mcp, Rest, Code). */
	// @ApiMember(Description="Filter by transport (Mcp, Rest, Code).")
	Transports []MarketplaceTransport `json:"transports,omitempty"`
	/** @description Free-text search over the listing's display name, vendor, and description. */
	// @ApiMember(Description="Free-text search over the listing's display name, vendor, and description.")
	Search *string `json:"search,omitempty"`
	/** @description If true, return only listings curated and verified by Norbix. */
	// @ApiMember(Description="If true, return only listings curated and verified by Norbix.")
	OfficialOnly *bool `json:"officialOnly,omitempty"`
	/** @description Filter by curated tag slugs (e.g. ai-llm, messaging, crm). Matches listings carrying any of the given tags. */
	// @ApiMember(Description="Filter by curated tag slugs (e.g. ai-llm, messaging, crm). Matches listings carrying any of the given tags.")
	Tags []string `json:"tags,omitempty"`
}

// @Route("/{version}/code/marketplace/listings/{ListingViewId}/functions/{FunctionKey}/tokens", "GET")
type GetMarketplaceListingFunctionTokens struct {
	CodeMashRequestBase
	/** @description Marketplace listing view id, from get_marketplace_listings. */
	// @ApiMember(Description="Marketplace listing view id, from get_marketplace_listings.", IsRequired=true)
	ListingViewId string `json:"listingViewId"`
	/** @description Function key on the listing, from get_marketplace_listings. */
	// @ApiMember(Description="Function key on the listing, from get_marketplace_listings.", IsRequired=true)
	FunctionKey string `json:"functionKey"`
}

// @Route("/{version}/code/marketplace/integrations", "GET")
type GetMarketplaceIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}", "GET")
type GetMarketplaceIntegration struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/{version}/code/marketplace/integrations", "POST")
// @DataContract
type SaveMarketplaceIntegration struct {
	CodeMashRequestBase
	/** @description The marketplace integration to install, from a get_marketplace_listings entry. */
	// @DataMember
	// @ApiMember(Description="The marketplace integration to install, from a get_marketplace_listings entry.", IsRequired=true)
	Integration MarketplaceIntegrationDto `json:"integration"`
	// @DataMember
	Secrets map[string]string `json:"secrets"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}", "DELETE")
type DeleteMarketplaceIntegration struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/enable", "POST")
type EnableMarketplaceIntegration struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/disable", "POST")
type DisableMarketplaceIntegration struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions", "GET")
type GetMarketplaceFunctions struct {
	CodeMashListPaginationRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions/{FunctionViewId}", "GET")
type GetMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	/** @description Function view id, from get_marketplace_functions. */
	// @ApiMember(Description="Function view id, from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions", "POST")
// @DataContract
type SaveMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @DataMember
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	/** @description The function to create: functionKey, displayName, description, and one mapping per vendor function parameter. */
	// @DataMember
	// @ApiMember(Description="The function to create: functionKey, displayName, description, and one mapping per vendor function parameter.", IsRequired=true)
	Function MarketplaceFunctionDto `json:"function"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions/{FunctionViewId}", "DELETE")
type DeleteMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	/** @description Function view id, from get_marketplace_functions. */
	// @ApiMember(Description="Function view id, from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions/{FunctionViewId}/enable", "POST")
type EnableMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	/** @description Function view id, from get_marketplace_functions. */
	// @ApiMember(Description="Function view id, from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions/{FunctionViewId}/disable", "POST")
type DisableMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	/** @description Function view id, from get_marketplace_functions. */
	// @ApiMember(Description="Function view id, from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/functions/{FunctionViewId}/tokens", "GET")
type GetMarketplaceFunctionTokens struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	/** @description Function view id, from get_marketplace_functions. */
	// @ApiMember(Description="Function view id, from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
}

// @Route("/{version}/code/marketplace/functions/{FunctionViewId}/invoke", "POST")
// @DataContract
type InvokeMarketplaceFunction struct {
	CodeMashRequestBase
	/** @description Function view id (func_…), from get_marketplace_functions. */
	// @DataMember
	// @ApiMember(Description="Function view id (func_…), from get_marketplace_functions.", IsRequired=true)
	FunctionViewId string `json:"functionViewId"`
	// @DataMember
	Payload map[string]Object `json:"payload"`
}

// @Route("/{version}/code/marketplace/listings/{ListingViewId}", "GET")
type GetMarketplaceListing struct {
	CodeMashRequestBase
	/** @description Listing view id (ml_…), from get_marketplace_listings. */
	// @ApiMember(Description="Listing view id (ml_…), from get_marketplace_listings.", IsRequired=true)
	ListingViewId string `json:"listingViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/test", "POST")
type TestMarketplaceIntegration struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/internal/_typegen", "GET")
type InternalsTypeGen struct {
	Typegen0SmtpEmailIntegrationRequest                        *SmtpEmailIntegrationRequest                       `json:"typegen_0_SmtpEmailIntegrationRequest,omitempty"`
	Typegen1AwsSesEmailIntegrationRequest                      *AwsSesEmailIntegrationRequest                     `json:"typegen_1_AwsSesEmailIntegrationRequest,omitempty"`
	Typegen2SendGridEmailIntegrationRequest                    *SendGridEmailIntegrationRequest                   `json:"typegen_2_SendGridEmailIntegrationRequest,omitempty"`
	Typegen3MailGunEmailIntegrationRequest                     *MailGunEmailIntegrationRequest                    `json:"typegen_3_MailGunEmailIntegrationRequest,omitempty"`
	Typegen4EmailToAllUsersDeliverySettingsRequest             *EmailToAllUsersDeliverySettingsRequest            `json:"typegen_4_EmailToAllUsersDeliverySettingsRequest,omitempty"`
	Typegen5EmailToAccountUsersDeliverySettingsRequest         *EmailToAccountUsersDeliverySettingsRequest        `json:"typegen_5_EmailToAccountUsersDeliverySettingsRequest,omitempty"`
	Typegen6EmailToCollectionRecordsDeliverySettingsRequest    *EmailToCollectionRecordsDeliverySettingsRequest   `json:"typegen_6_EmailToCollectionRecordsDeliverySettingsRequest,omitempty"`
	Typegen7EmailToEmailsDeliverySettingsRequest               *EmailToEmailsDeliverySettingsRequest              `json:"typegen_7_EmailToEmailsDeliverySettingsRequest,omitempty"`
	Typegen8EmailToUsersDeliverySettingsRequest                *EmailToUsersDeliverySettingsRequest               `json:"typegen_8_EmailToUsersDeliverySettingsRequest,omitempty"`
	Typegen9MembershipTriggerRequest                           *MembershipTriggerRequest                          `json:"typegen_9_MembershipTriggerRequest,omitempty"`
	Typegen10SchemaTriggerRequest                              *SchemaTriggerRequest                              `json:"typegen_10_SchemaTriggerRequest,omitempty"`
	Typegen11FilesTriggerRequest                               *FilesTriggerRequest                               `json:"typegen_11_FilesTriggerRequest,omitempty"`
	Typegen12PaymentTriggerRequest                             *PaymentTriggerRequest                             `json:"typegen_12_PaymentTriggerRequest,omitempty"`
	Typegen15MongoDbConnectionStringDatabaseIntegrationRequest *MongoDbConnectionStringDatabaseIntegrationRequest `json:"typegen_15_MongoDbConnectionStringDatabaseIntegrationRequest,omitempty"`
	Typegen16MongoDbAtlasFlexManagedDatabaseIntegrationRequest *MongoDbAtlasFlexManagedDatabaseIntegrationRequest `json:"typegen_16_MongoDbAtlasFlexManagedDatabaseIntegrationRequest,omitempty"`
	Typegen16GoogleDriveFilesIntegrationRequest                *GoogleDriveFilesIntegrationRequest                `json:"typegen_16_GoogleDriveFilesIntegrationRequest,omitempty"`
	Typegen17FtpFilesIntegrationRequest                        *FtpFilesIntegrationRequest                        `json:"typegen_17_FtpFilesIntegrationRequest,omitempty"`
	Typegen18DropBoxFilesIntegrationRequest                    *DropBoxFilesIntegrationRequest                    `json:"typegen_18_DropBoxFilesIntegrationRequest,omitempty"`
	Typegen19AppleICloudFilesIntegrationRequest                *AppleICloudFilesIntegrationRequest                `json:"typegen_19_AppleICloudFilesIntegrationRequest,omitempty"`
	Typegen20AwsS3FilesIntegrationRequest                      *AwsS3FilesIntegrationRequest                      `json:"typegen_20_AwsS3FilesIntegrationRequest,omitempty"`
	Typegen21GoogleCloudFilesIntegrationRequest                *GoogleCloudFilesIntegrationRequest                `json:"typegen_21_GoogleCloudFilesIntegrationRequest,omitempty"`
	Typegen22AzureBlobFilesIntegrationRequest                  *AzureBlobFilesIntegrationRequest                  `json:"typegen_22_AzureBlobFilesIntegrationRequest,omitempty"`
	Typegen23LocalFilesIntegrationRequest                      *LocalFilesIntegrationRequest                      `json:"typegen_23_LocalFilesIntegrationRequest,omitempty"`
	Typegen24AmqpLoggingIntegrationRequest                     *AmqpLoggingIntegrationRequest                     `json:"typegen_24_AmqpLoggingIntegrationRequest,omitempty"`
	Typegen25AwsKinesisLoggingIntegrationRequest               *AwsKinesisLoggingIntegrationRequest               `json:"typegen_25_AwsKinesisLoggingIntegrationRequest,omitempty"`
	Typegen26AwsS3LoggingIntegrationRequest                    *AwsS3LoggingIntegrationRequest                    `json:"typegen_26_AwsS3LoggingIntegrationRequest,omitempty"`
	Typegen28NewRelicLoggingIntegrationRequest                 *NewRelicLoggingIntegrationRequest                 `json:"typegen_28_NewRelicLoggingIntegrationRequest,omitempty"`
	Typegen30MongoDbLoggingIntegrationRequest                  *MongoDbLoggingIntegrationRequest                  `json:"typegen_30_MongoDbLoggingIntegrationRequest,omitempty"`
	Typegen31KafkaLoggingIntegrationRequest                    *KafkaLoggingIntegrationRequest                    `json:"typegen_31_KafkaLoggingIntegrationRequest,omitempty"`
	Typegen32PrometheusLoggingIntegrationRequest               *PrometheusLoggingIntegrationRequest               `json:"typegen_32_PrometheusLoggingIntegrationRequest,omitempty"`
	Typegen33DataDogLoggingIntegrationRequest                  *DataDogLoggingIntegrationRequest                  `json:"typegen_33_DataDogLoggingIntegrationRequest,omitempty"`
	Typegen34InternalKafkaLoggingIntegrationRequest            *InternalKafkaLoggingIntegrationRequest            `json:"typegen_34_InternalKafkaLoggingIntegrationRequest,omitempty"`
	Typegen35ElasticSearchLoggingIntegrationRequest            *ElasticSearchLoggingIntegrationRequest            `json:"typegen_35_ElasticSearchLoggingIntegrationRequest,omitempty"`
	Typegen37SplunkLoggingIntegrationRequest                   *SplunkLoggingIntegrationRequest                   `json:"typegen_37_SplunkLoggingIntegrationRequest,omitempty"`
	Typegen38AzureOtelLoggingIntegrationRequest                *AzureOtelLoggingIntegrationRequest                `json:"typegen_38_AzureOtelLoggingIntegrationRequest,omitempty"`
	Typegen39KibanaLoggingIntegrationRequest                   *KibanaLoggingIntegrationRequest                   `json:"typegen_39_KibanaLoggingIntegrationRequest,omitempty"`
	Typegen40LocalFileLoggingIntegrationRequest                *LocalFileLoggingIntegrationRequest                `json:"typegen_40_LocalFileLoggingIntegrationRequest,omitempty"`
	Typegen41OktaMembershipIntegrationRequest                  *OktaMembershipIntegrationRequest                  `json:"typegen_41_OktaMembershipIntegrationRequest,omitempty"`
	Typegen42XMembershipIntegrationRequest                     *XMembershipIntegrationRequest                     `json:"typegen_42_XMembershipIntegrationRequest,omitempty"`
	Typegen43GoogleMembershipIntegrationRequest                *GoogleMembershipIntegrationRequest                `json:"typegen_43_GoogleMembershipIntegrationRequest,omitempty"`
	Typegen44MicrosoftMembershipIntegrationRequest             *MicrosoftMembershipIntegrationRequest             `json:"typegen_44_MicrosoftMembershipIntegrationRequest,omitempty"`
	Typegen45GitHubMembershipIntegrationRequest                *GitHubMembershipIntegrationRequest                `json:"typegen_45_GitHubMembershipIntegrationRequest,omitempty"`
	Typegen46MetaMembershipIntegrationRequest                  *MetaMembershipIntegrationRequest                  `json:"typegen_46_MetaMembershipIntegrationRequest,omitempty"`
	Typegen47AppleMembershipIntegrationRequest                 *AppleMembershipIntegrationRequest                 `json:"typegen_47_AppleMembershipIntegrationRequest,omitempty"`
	Typegen48LemonSqueezyPaymentIntegrationRequest             *LemonSqueezyPaymentIntegrationRequest             `json:"typegen_48_LemonSqueezyPaymentIntegrationRequest,omitempty"`
	Typegen49AdyenPaymentIntegrationRequest                    *AdyenPaymentIntegrationRequest                    `json:"typegen_49_AdyenPaymentIntegrationRequest,omitempty"`
	Typegen50MolliePaymentIntegrationRequest                   *MolliePaymentIntegrationRequest                   `json:"typegen_50_MolliePaymentIntegrationRequest,omitempty"`
	Typegen51PaddlePaymentIntegrationRequest                   *PaddlePaymentIntegrationRequest                   `json:"typegen_51_PaddlePaymentIntegrationRequest,omitempty"`
	Typegen52PayPalPaymentIntegrationRequest                   *PayPalPaymentIntegrationRequest                   `json:"typegen_52_PayPalPaymentIntegrationRequest,omitempty"`
	Typegen53StripePaymentIntegrationRequest                   *StripePaymentIntegrationRequest                   `json:"typegen_53_StripePaymentIntegrationRequest,omitempty"`
	Typegen54AppleInAppPaymentIntegrationRequest               *AppleInAppPaymentIntegrationRequest               `json:"typegen_54_AppleInAppPaymentIntegrationRequest,omitempty"`
	Typegen55GoogleInAppPaymentIntegrationRequest              *GoogleInAppPaymentIntegrationRequest              `json:"typegen_55_GoogleInAppPaymentIntegrationRequest,omitempty"`
	Typegen56EdgeWebPushIntegrationRequest                     *EdgeWebPushIntegrationRequest                     `json:"typegen_56_EdgeWebPushIntegrationRequest,omitempty"`
	Typegen57ChromePluginPushIntegrationRequest                *ChromePluginPushIntegrationRequest                `json:"typegen_57_ChromePluginPushIntegrationRequest,omitempty"`
	Typegen58SafariPushIntegrationRequest                      *SafariPushIntegrationRequest                      `json:"typegen_58_SafariPushIntegrationRequest,omitempty"`
	Typegen59ChromeWebPushIntegrationRequest                   *ChromeWebPushIntegrationRequest                   `json:"typegen_59_ChromeWebPushIntegrationRequest,omitempty"`
	Typegen60FirefoxWebPushIntegrationRequest                  *FirefoxWebPushIntegrationRequest                  `json:"typegen_60_FirefoxWebPushIntegrationRequest,omitempty"`
	Typegen61AndroidFirebasePushIntegrationRequest             *AndroidFirebasePushIntegrationRequest             `json:"typegen_61_AndroidFirebasePushIntegrationRequest,omitempty"`
	Typegen62AppleApnsPushIntegrationRequest                   *AppleApnsPushIntegrationRequest                   `json:"typegen_62_AppleApnsPushIntegrationRequest,omitempty"`
	Typegen65AwsLambdaCodeIntegrationRequest                   *AwsLambdaCodeIntegrationRequest                   `json:"typegen_65_AwsLambdaCodeIntegrationRequest,omitempty"`
	Typegen66AzureFunctionsCodeIntegrationRequest              *AzureFunctionsCodeIntegrationRequest              `json:"typegen_66_AzureFunctionsCodeIntegrationRequest,omitempty"`
	Typegen67GoogleCloudFunctionsCodeIntegrationRequest        *GoogleCloudFunctionsCodeIntegrationRequest        `json:"typegen_67_GoogleCloudFunctionsCodeIntegrationRequest,omitempty"`
	Typegen68OllamaLlmIntegrationRequest                       *OllamaLlmIntegrationRequest                       `json:"typegen_68_OllamaLlmIntegrationRequest,omitempty"`
	Typegen69OpenRouterLlmIntegrationRequest                   *OpenRouterLlmIntegrationRequest                   `json:"typegen_69_OpenRouterLlmIntegrationRequest,omitempty"`
	Typegen70MistralLlmIntegrationRequest                      *MistralLlmIntegrationRequest                      `json:"typegen_70_MistralLlmIntegrationRequest,omitempty"`
	Typegen71GrokLlmIntegrationRequest                         *GrokLlmIntegrationRequest                         `json:"typegen_71_GrokLlmIntegrationRequest,omitempty"`
	Typegen72GroqLlmIntegrationRequest                         *GroqLlmIntegrationRequest                         `json:"typegen_72_GroqLlmIntegrationRequest,omitempty"`
	Typegen73GoogleLlmIntegrationRequest                       *GoogleLlmIntegrationRequest                       `json:"typegen_73_GoogleLlmIntegrationRequest,omitempty"`
	Typegen74AnthropicLlmIntegrationRequest                    *AnthropicLlmIntegrationRequest                    `json:"typegen_74_AnthropicLlmIntegrationRequest,omitempty"`
	Typegen75OpenAiLlmIntegrationRequest                       *OpenAiLlmIntegrationRequest                       `json:"typegen_75_OpenAiLlmIntegrationRequest,omitempty"`
	Typegen76PlaywrightMcpIntegrationRequest                   *PlaywrightMcpIntegrationRequest                   `json:"typegen_76_PlaywrightMcpIntegrationRequest,omitempty"`
	Typegen77MongoDbMcpIntegrationRequest                      *MongoDbMcpIntegrationRequest                      `json:"typegen_77_MongoDbMcpIntegrationRequest,omitempty"`
	Typegen78GitHubMcpIntegrationRequest                       *GitHubMcpIntegrationRequest                       `json:"typegen_78_GitHubMcpIntegrationRequest,omitempty"`
	Typegen79StripeMcpIntegrationRequest                       *StripeMcpIntegrationRequest                       `json:"typegen_79_StripeMcpIntegrationRequest,omitempty"`
	Typegen80BraveSearchMcpIntegrationRequest                  *BraveSearchMcpIntegrationRequest                  `json:"typegen_80_BraveSearchMcpIntegrationRequest,omitempty"`
	Typegen81ObsidianMcpIntegrationRequest                     *ObsidianMcpIntegrationRequest                     `json:"typegen_81_ObsidianMcpIntegrationRequest,omitempty"`
	Typegen82EmailTemplateDto                                  *EmailTemplateDto                                  `json:"typegen_82_EmailTemplateDto,omitempty"`
	Typegen83PushTemplateDto                                   *PushTemplateDto                                   `json:"typegen_83_PushTemplateDto,omitempty"`
	Typegen84SmsTemplateDto                                    *SmsTemplateDto                                    `json:"typegen_84_SmsTemplateDto,omitempty"`
	Typegen85SystemEmailTemplateDto                            *SystemEmailTemplateDto                            `json:"typegen_85_SystemEmailTemplateDto,omitempty"`
	Typegen86TriggerActionEmailDto                             *TriggerActionEmailDto                             `json:"typegen_86_TriggerActionEmailDto,omitempty"`
	Typegen87TriggerActionPushDto                              *TriggerActionPushDto                              `json:"typegen_87_TriggerActionPushDto,omitempty"`
	Typegen88TriggerActionCodeDto                              *TriggerActionCodeDto                              `json:"typegen_88_TriggerActionCodeDto,omitempty"`
	Typegen89TriggerActionWebhookDto                           *TriggerActionWebhookDto                           `json:"typegen_89_TriggerActionWebhookDto,omitempty"`
	Typegen236TriggerActionSmsDto                              *TriggerActionSmsDto                               `json:"typegen_236_TriggerActionSmsDto,omitempty"`
	Typegen237TriggerActionSseDto                              *TriggerActionSseDto                               `json:"typegen_237_TriggerActionSseDto,omitempty"`
	Typegen238TriggerActionMarketplaceDto                      *TriggerActionMarketplaceDto                       `json:"typegen_238_TriggerActionMarketplaceDto,omitempty"`
	Typegen239SseDeliverySettingsDto                           *SseDeliverySettingsDto                            `json:"typegen_239_SseDeliverySettingsDto,omitempty"`
	Typegen240GetTriggers                                      *GetTriggers                                       `json:"typegen_240_GetTriggers,omitempty"`
	Typegen241GetTriggersResponse                              *GetTriggersResponse                               `json:"typegen_241_GetTriggersResponse,omitempty"`
	Typegen90EmailToAllUsersDeliverySettingsDto                *EmailToAllUsersDeliverySettingsDto                `json:"typegen_90_EmailToAllUsersDeliverySettingsDto,omitempty"`
	Typegen91EmailToAccountUsersDeliverySettingsDto            *EmailToAccountUsersDeliverySettingsDto            `json:"typegen_91_EmailToAccountUsersDeliverySettingsDto,omitempty"`
	Typegen92EmailToUsersDeliverySettingsDto                   *EmailToUsersDeliverySettingsDto                   `json:"typegen_92_EmailToUsersDeliverySettingsDto,omitempty"`
	Typegen93EmailToEmailAddressesDeliverySettingsDto          *EmailToEmailAddressesDeliverySettingsDto          `json:"typegen_93_EmailToEmailAddressesDeliverySettingsDto,omitempty"`
	Typegen94EmailToCollectionRecordsDeliverySettingsDto       *EmailToCollectionRecordsDeliverySettingsDto       `json:"typegen_94_EmailToCollectionRecordsDeliverySettingsDto,omitempty"`
	Typegen95PushToAllUsersDeliverySettingsDto                 *PushToAllUsersDeliverySettingsDto                 `json:"typegen_95_PushToAllUsersDeliverySettingsDto,omitempty"`
	Typegen96PushToUsersDeliverySettingsDto                    *PushToUsersDeliverySettingsDto                    `json:"typegen_96_PushToUsersDeliverySettingsDto,omitempty"`
	Typegen229PushToAccountUsersDeliverySettingsDto            *PushToAccountUsersDeliverySettingsDto             `json:"typegen_229_PushToAccountUsersDeliverySettingsDto,omitempty"`
	Typegen97PushToCollectionRecordsDeliverySettingsDto        *PushToCollectionRecordsDeliverySettingsDto        `json:"typegen_97_PushToCollectionRecordsDeliverySettingsDto,omitempty"`
	Typegen98PushToDevicesDeliverySettingsDto                  *PushToDevicesDeliverySettingsDto                  `json:"typegen_98_PushToDevicesDeliverySettingsDto,omitempty"`
	Typegen99SmsToAllUsersDeliverySettingsDto                  *SmsToAllUsersDeliverySettingsDto                  `json:"typegen_99_SmsToAllUsersDeliverySettingsDto,omitempty"`
	Typegen100SmsToUsersDeliverySettingsDto                    *SmsToUsersDeliverySettingsDto                     `json:"typegen_100_SmsToUsersDeliverySettingsDto,omitempty"`
	Typegen101SmsToCollectionRecordsDeliverySettingsDto        *SmsToCollectionRecordsDeliverySettingsDto         `json:"typegen_101_SmsToCollectionRecordsDeliverySettingsDto,omitempty"`
	Typegen102SmsToPhoneNumbersDeliverySettingsDto             *SmsToPhoneNumbersDeliverySettingsDto              `json:"typegen_102_SmsToPhoneNumbersDeliverySettingsDto,omitempty"`
	Typegen103OpenAiLlmIntegrationDto                          *OpenAiLlmIntegrationDto                           `json:"typegen_103_OpenAiLlmIntegrationDto,omitempty"`
	Typegen104AnthropicLlmIntegrationDto                       *AnthropicLlmIntegrationDto                        `json:"typegen_104_AnthropicLlmIntegrationDto,omitempty"`
	Typegen105OllamaLlmIntegrationDto                          *OllamaLlmIntegrationDto                           `json:"typegen_105_OllamaLlmIntegrationDto,omitempty"`
	Typegen106GroqLlmIntegrationDto                            *GroqLlmIntegrationDto                             `json:"typegen_106_GroqLlmIntegrationDto,omitempty"`
	Typegen107GoogleLlmIntegrationDto                          *GoogleLlmIntegrationDto                           `json:"typegen_107_GoogleLlmIntegrationDto,omitempty"`
	Typegen108MistralLlmIntegrationDto                         *MistralLlmIntegrationDto                          `json:"typegen_108_MistralLlmIntegrationDto,omitempty"`
	Typegen109OpenRouterLlmIntegrationDto                      *OpenRouterLlmIntegrationDto                       `json:"typegen_109_OpenRouterLlmIntegrationDto,omitempty"`
	Typegen110GrokLlmIntegrationDto                            *GrokLlmIntegrationDto                             `json:"typegen_110_GrokLlmIntegrationDto,omitempty"`
	Typegen111DockerMcpIntegrationDto                          *DockerMcpIntegrationDto                           `json:"typegen_111_DockerMcpIntegrationDto,omitempty"`
	Typegen112GoogleCalendarMcpIntegrationDto                  *GoogleCalendarMcpIntegrationDto                   `json:"typegen_112_GoogleCalendarMcpIntegrationDto,omitempty"`
	Typegen113ObsidianMcpIntegrationDto                        *ObsidianMcpIntegrationDto                         `json:"typegen_113_ObsidianMcpIntegrationDto,omitempty"`
	Typegen114AwsLambdaCrossAccountRoleCodeIntegrationDto      *AwsLambdaCrossAccountRoleCodeIntegrationDto       `json:"typegen_114_AwsLambdaCrossAccountRoleCodeIntegrationDto,omitempty"`
	Typegen115AwsLambdaIamCodeIntegrationDto                   *AwsLambdaIamCodeIntegrationDto                    `json:"typegen_115_AwsLambdaIamCodeIntegrationDto,omitempty"`
	Typegen116AzureFunctionsCodeIntegrationDto                 *AzureFunctionsCodeIntegrationDto                  `json:"typegen_116_AzureFunctionsCodeIntegrationDto,omitempty"`
	Typegen118GoogleCloudFunctionsCodeIntegrationDto           *GoogleCloudFunctionsCodeIntegrationDto            `json:"typegen_118_GoogleCloudFunctionsCodeIntegrationDto,omitempty"`
	Typegen120AdyenPaymentIntegrationDto                       *AdyenPaymentIntegrationDto                        `json:"typegen_120_AdyenPaymentIntegrationDto,omitempty"`
	Typegen121AppleInAppPaymentIntegrationDto                  *AppleInAppPaymentIntegrationDto                   `json:"typegen_121_AppleInAppPaymentIntegrationDto,omitempty"`
	Typegen122GoogleInAppPaymentIntegrationDto                 *GoogleInAppPaymentIntegrationDto                  `json:"typegen_122_GoogleInAppPaymentIntegrationDto,omitempty"`
	Typegen123LemonSqueezyPaymentIntegrationDto                *LemonSqueezyPaymentIntegrationDto                 `json:"typegen_123_LemonSqueezyPaymentIntegrationDto,omitempty"`
	Typegen124MolliePaymentIntegrationDto                      *MolliePaymentIntegrationDto                       `json:"typegen_124_MolliePaymentIntegrationDto,omitempty"`
	Typegen125PaddlePaymentIntegrationDto                      *PaddlePaymentIntegrationDto                       `json:"typegen_125_PaddlePaymentIntegrationDto,omitempty"`
	Typegen126PayPalPaymentIntegrationDto                      *PayPalPaymentIntegrationDto                       `json:"typegen_126_PayPalPaymentIntegrationDto,omitempty"`
	Typegen127StripePaymentIntegrationDto                      *StripePaymentIntegrationDto                       `json:"typegen_127_StripePaymentIntegrationDto,omitempty"`
	Typegen184ShopifyPaymentIntegrationDto                     *ShopifyPaymentIntegrationDto                      `json:"typegen_184_ShopifyPaymentIntegrationDto,omitempty"`
	Typegen185WooCommercePaymentIntegrationDto                 *WooCommercePaymentIntegrationDto                  `json:"typegen_185_WooCommercePaymentIntegrationDto,omitempty"`
	Typegen186MagentoPaymentIntegrationDto                     *MagentoPaymentIntegrationDto                      `json:"typegen_186_MagentoPaymentIntegrationDto,omitempty"`
	Typegen187BraintreePaymentIntegrationDto                   *BraintreePaymentIntegrationDto                    `json:"typegen_187_BraintreePaymentIntegrationDto,omitempty"`
	Typegen188AuthorizeNetPaymentIntegrationDto                *AuthorizeNetPaymentIntegrationDto                 `json:"typegen_188_AuthorizeNetPaymentIntegrationDto,omitempty"`
	Typegen189CheckOutComPaymentIntegrationDto                 *CheckOutComPaymentIntegrationDto                  `json:"typegen_189_CheckOutComPaymentIntegrationDto,omitempty"`
	Typegen190WorldpayPaymentIntegrationDto                    *WorldpayPaymentIntegrationDto                     `json:"typegen_190_WorldpayPaymentIntegrationDto,omitempty"`
	Typegen128AppleSignInMembershipIntegrationDto              *AppleSignInMembershipIntegrationDto               `json:"typegen_128_AppleSignInMembershipIntegrationDto,omitempty"`
	Typegen129GitHubMembershipIntegrationDto                   *GitHubMembershipIntegrationDto                    `json:"typegen_129_GitHubMembershipIntegrationDto,omitempty"`
	Typegen130GoogleMembershipIntegrationDto                   *GoogleMembershipIntegrationDto                    `json:"typegen_130_GoogleMembershipIntegrationDto,omitempty"`
	Typegen131MetaMembershipIntegrationDto                     *MetaMembershipIntegrationDto                      `json:"typegen_131_MetaMembershipIntegrationDto,omitempty"`
	Typegen132MicrosoftMembershipIntegrationDto                *MicrosoftMembershipIntegrationDto                 `json:"typegen_132_MicrosoftMembershipIntegrationDto,omitempty"`
	Typegen133OktaMembershipIntegrationDto                     *OktaMembershipIntegrationDto                      `json:"typegen_133_OktaMembershipIntegrationDto,omitempty"`
	Typegen134XMembershipIntegrationDto                        *XMembershipIntegrationDto                         `json:"typegen_134_XMembershipIntegrationDto,omitempty"`
	Typegen135AmqpLoggingIntegrationDto                        *AmqpLoggingIntegrationDto                         `json:"typegen_135_AmqpLoggingIntegrationDto,omitempty"`
	Typegen136AwsKinesisLoggingIntegrationDto                  *AwsKinesisLoggingIntegrationDto                   `json:"typegen_136_AwsKinesisLoggingIntegrationDto,omitempty"`
	Typegen137AwsS3CrossAccountRoleLoggingIntegrationDto       *AwsS3CrossAccountRoleLoggingIntegrationDto        `json:"typegen_137_AwsS3CrossAccountRoleLoggingIntegrationDto,omitempty"`
	Typegen138AwsS3IamLoggingIntegrationDto                    *AwsS3IamLoggingIntegrationDto                     `json:"typegen_138_AwsS3IamLoggingIntegrationDto,omitempty"`
	Typegen139AzureOtelLoggingIntegrationDto                   *AzureOtelLoggingIntegrationDto                    `json:"typegen_139_AzureOtelLoggingIntegrationDto,omitempty"`
	Typegen140DataDogLoggingIntegrationDto                     *DataDogLoggingIntegrationDto                      `json:"typegen_140_DataDogLoggingIntegrationDto,omitempty"`
	Typegen141ElasticSearchLoggingIntegrationDto               *ElasticSearchLoggingIntegrationDto                `json:"typegen_141_ElasticSearchLoggingIntegrationDto,omitempty"`
	Typegen142InternalKafkaLoggingIntegrationDto               *InternalKafkaLoggingIntegrationDto                `json:"typegen_142_InternalKafkaLoggingIntegrationDto,omitempty"`
	Typegen143KafkaLoggingIntegrationDto                       *KafkaLoggingIntegrationDto                        `json:"typegen_143_KafkaLoggingIntegrationDto,omitempty"`
	Typegen144KibanaLoggingIntegrationDto                      *KibanaLoggingIntegrationDto                       `json:"typegen_144_KibanaLoggingIntegrationDto,omitempty"`
	Typegen145LocalFileLoggingIntegrationDto                   *LocalFileLoggingIntegrationDto                    `json:"typegen_145_LocalFileLoggingIntegrationDto,omitempty"`
	Typegen147MongoDbLoggingIntegrationDto                     *MongoDbLoggingIntegrationDto                      `json:"typegen_147_MongoDbLoggingIntegrationDto,omitempty"`
	Typegen148NewRelicLoggingIntegrationDto                    *NewRelicLoggingIntegrationDto                     `json:"typegen_148_NewRelicLoggingIntegrationDto,omitempty"`
	Typegen149PrometheusLoggingIntegrationDto                  *PrometheusLoggingIntegrationDto                   `json:"typegen_149_PrometheusLoggingIntegrationDto,omitempty"`
	Typegen150SplunkLoggingIntegrationDto                      *SplunkLoggingIntegrationDto                       `json:"typegen_150_SplunkLoggingIntegrationDto,omitempty"`
	Typegen153AppleICloudFilesIntegrationDto                   *AppleICloudFilesIntegrationDto                    `json:"typegen_153_AppleICloudFilesIntegrationDto,omitempty"`
	Typegen154AwsS3CrossAccountRoleFilesIntegrationDto         *AwsS3CrossAccountRoleFilesIntegrationDto          `json:"typegen_154_AwsS3CrossAccountRoleFilesIntegrationDto,omitempty"`
	Typegen155AwsS3IamFilesIntegrationDto                      *AwsS3IamFilesIntegrationDto                       `json:"typegen_155_AwsS3IamFilesIntegrationDto,omitempty"`
	Typegen156AzureBlobFilesIntegrationDto                     *AzureBlobFilesIntegrationDto                      `json:"typegen_156_AzureBlobFilesIntegrationDto,omitempty"`
	Typegen157DropBoxFilesIntegrationDto                       *DropBoxFilesIntegrationDto                        `json:"typegen_157_DropBoxFilesIntegrationDto,omitempty"`
	Typegen158FtpFilesIntegrationDto                           *FtpFilesIntegrationDto                            `json:"typegen_158_FtpFilesIntegrationDto,omitempty"`
	Typegen159GoogleCloudFilesIntegrationDto                   *GoogleCloudFilesIntegrationDto                    `json:"typegen_159_GoogleCloudFilesIntegrationDto,omitempty"`
	Typegen160GoogleDriveFilesIntegrationDto                   *GoogleDriveFilesIntegrationDto                    `json:"typegen_160_GoogleDriveFilesIntegrationDto,omitempty"`
	Typegen161LocalFilesIntegrationDto                         *LocalFilesIntegrationDto                          `json:"typegen_161_LocalFilesIntegrationDto,omitempty"`
	Typegen164MongoDbConnectionStringIntegrationDto            *MongoDbConnectionStringIntegrationDto             `json:"typegen_164_MongoDbConnectionStringIntegrationDto,omitempty"`
	Typegen165MongoDbAtlasFlexManagedIntegrationDto            *MongoDbAtlasFlexManagedIntegrationDto             `json:"typegen_165_MongoDbAtlasFlexManagedIntegrationDto,omitempty"`
	Typegen165BirdSmsIntegrationDto                            *BirdSmsIntegrationDto                             `json:"typegen_165_BirdSmsIntegrationDto,omitempty"`
	Typegen166PlivoSmsIntegrationDto                           *PlivoSmsIntegrationDto                            `json:"typegen_166_PlivoSmsIntegrationDto,omitempty"`
	Typegen167SinchSmsIntegrationDto                           *SinchSmsIntegrationDto                            `json:"typegen_167_SinchSmsIntegrationDto,omitempty"`
	Typegen168TelesignSmsIntegrationDto                        *TelesignSmsIntegrationDto                         `json:"typegen_168_TelesignSmsIntegrationDto,omitempty"`
	Typegen169TelnyxSmsIntegrationDto                          *TelnyxSmsIntegrationDto                           `json:"typegen_169_TelnyxSmsIntegrationDto,omitempty"`
	Typegen170TwilioSmsIntegrationDto                          *TwilioSmsIntegrationDto                           `json:"typegen_170_TwilioSmsIntegrationDto,omitempty"`
	Typegen171VonageSmsIntegrationDto                          *VonageSmsIntegrationDto                           `json:"typegen_171_VonageSmsIntegrationDto,omitempty"`
	Typegen172AndroidFirebasePushIntegrationDto                *AndroidFirebasePushIntegrationDto                 `json:"typegen_172_AndroidFirebasePushIntegrationDto,omitempty"`
	Typegen173AppleApnsPushIntegrationDto                      *AppleApnsPushIntegrationDto                       `json:"typegen_173_AppleApnsPushIntegrationDto,omitempty"`
	Typegen174ChromePluginPushIntegrationDto                   *ChromePluginPushIntegrationDto                    `json:"typegen_174_ChromePluginPushIntegrationDto,omitempty"`
	Typegen175ChromeWebPushIntegrationDto                      *ChromeWebPushIntegrationDto                       `json:"typegen_175_ChromeWebPushIntegrationDto,omitempty"`
	Typegen176EdgeWebPushIntegrationDto                        *EdgeWebPushIntegrationDto                         `json:"typegen_176_EdgeWebPushIntegrationDto,omitempty"`
	Typegen177FirefoxWebPushIntegrationDto                     *FirefoxWebPushIntegrationDto                      `json:"typegen_177_FirefoxWebPushIntegrationDto,omitempty"`
	Typegen178SafariPushIntegrationDto                         *SafariPushIntegrationDto                          `json:"typegen_178_SafariPushIntegrationDto,omitempty"`
	Typegen179AwsCrossAccountRoleEmailIntegrationDto           *AwsCrossAccountRoleEmailIntegrationDto            `json:"typegen_179_AwsCrossAccountRoleEmailIntegrationDto,omitempty"`
	Typegen180AwsIamEmailIntegrationDto                        *AwsIamEmailIntegrationDto                         `json:"typegen_180_AwsIamEmailIntegrationDto,omitempty"`
	Typegen181MailGunEmailIntegrationDto                       *MailGunEmailIntegrationDto                        `json:"typegen_181_MailGunEmailIntegrationDto,omitempty"`
	Typegen182SendGridEmailIntegrationDto                      *SendGridEmailIntegrationDto                       `json:"typegen_182_SendGridEmailIntegrationDto,omitempty"`
	Typegen183SmtpEmailIntegrationDto                          *SmtpEmailIntegrationDto                           `json:"typegen_183_SmtpEmailIntegrationDto,omitempty"`
	Typegen192WebhookIntegrationDto                            *WebhookIntegrationDto                             `json:"typegen_192_WebhookIntegrationDto,omitempty"`
	Typegen193WebhookDestinationDto                            *WebhookDestinationDto                             `json:"typegen_193_WebhookDestinationDto,omitempty"`
	Typegen194SchedulerTaskDto                                 *SchedulerTaskDto                                  `json:"typegen_194_SchedulerTaskDto,omitempty"`
	Typegen195MongoDbAggregateDto                              *MongoDbAggregateDto                               `json:"typegen_195_MongoDbAggregateDto,omitempty"`
	Typegen196MarketplaceIntegrationDto                        *MarketplaceIntegrationDto                         `json:"typegen_196_MarketplaceIntegrationDto,omitempty"`
	Typegen197MarketplaceFunctionDto                           *MarketplaceFunctionDto                            `json:"typegen_197_MarketplaceFunctionDto,omitempty"`
	Typegen198MarketplaceListingDto                            *MarketplaceListingDto                             `json:"typegen_198_MarketplaceListingDto,omitempty"`
	Typegen199MarketplaceFunctionDefinitionDto                 *MarketplaceFunctionDefinitionDto                  `json:"typegen_199_MarketplaceFunctionDefinitionDto,omitempty"`
	Typegen200MarketplaceFunctionParameterDto                  *MarketplaceFunctionParameterDto                   `json:"typegen_200_MarketplaceFunctionParameterDto,omitempty"`
	Typegen201EnableCode                                       *EnableCode                                        `json:"typegen_201_EnableCode,omitempty"`
	Typegen202DisableCode                                      *DisableCode                                       `json:"typegen_202_DisableCode,omitempty"`
	Typegen203GetCodeIntegrations                              *GetCodeIntegrations                               `json:"typegen_203_GetCodeIntegrations,omitempty"`
	Typegen204GetCodeIntegration                               *GetCodeIntegration                                `json:"typegen_204_GetCodeIntegration,omitempty"`
	Typegen205SaveCodeIntegration                              *SaveCodeIntegration                               `json:"typegen_205_SaveCodeIntegration,omitempty"`
	Typegen206TestCodeIntegration                              *TestCodeIntegration                               `json:"typegen_206_TestCodeIntegration,omitempty"`
	Typegen207ConfirmCodeIntegrationHumanDeliveryRequest       *ConfirmCodeIntegrationHumanDeliveryRequest        `json:"typegen_207_ConfirmCodeIntegrationHumanDeliveryRequest,omitempty"`
	Typegen208SetCodeIntegrationAsDefault                      *SetCodeIntegrationAsDefault                       `json:"typegen_208_SetCodeIntegrationAsDefault,omitempty"`
	Typegen209DeleteCodeIntegrationRequest                     *DeleteCodeIntegrationRequest                      `json:"typegen_209_DeleteCodeIntegrationRequest,omitempty"`
	Typegen210EnableCodeIntegrationRequest                     *EnableCodeIntegrationRequest                      `json:"typegen_210_EnableCodeIntegrationRequest,omitempty"`
	Typegen211DisableCodeIntegrationRequest                    *DisableCodeIntegrationRequest                     `json:"typegen_211_DisableCodeIntegrationRequest,omitempty"`
	Typegen212GetMarketplaceListings                           *GetMarketplaceListings                            `json:"typegen_212_GetMarketplaceListings,omitempty"`
	Typegen213GetMarketplaceListingFunctionTokens              *GetMarketplaceListingFunctionTokens               `json:"typegen_213_GetMarketplaceListingFunctionTokens,omitempty"`
	Typegen214GetMarketplaceIntegrations                       *GetMarketplaceIntegrations                        `json:"typegen_214_GetMarketplaceIntegrations,omitempty"`
	Typegen215GetMarketplaceIntegration                        *GetMarketplaceIntegration                         `json:"typegen_215_GetMarketplaceIntegration,omitempty"`
	Typegen216SaveMarketplaceIntegration                       *SaveMarketplaceIntegration                        `json:"typegen_216_SaveMarketplaceIntegration,omitempty"`
	Typegen217DeleteMarketplaceIntegration                     *DeleteMarketplaceIntegration                      `json:"typegen_217_DeleteMarketplaceIntegration,omitempty"`
	Typegen218EnableMarketplaceIntegration                     *EnableMarketplaceIntegration                      `json:"typegen_218_EnableMarketplaceIntegration,omitempty"`
	Typegen219DisableMarketplaceIntegration                    *DisableMarketplaceIntegration                     `json:"typegen_219_DisableMarketplaceIntegration,omitempty"`
	Typegen221GetMarketplaceFunctions                          *GetMarketplaceFunctions                           `json:"typegen_221_GetMarketplaceFunctions,omitempty"`
	Typegen222GetMarketplaceFunction                           *GetMarketplaceFunction                            `json:"typegen_222_GetMarketplaceFunction,omitempty"`
	Typegen223SaveMarketplaceFunction                          *SaveMarketplaceFunction                           `json:"typegen_223_SaveMarketplaceFunction,omitempty"`
	Typegen224DeleteMarketplaceFunction                        *DeleteMarketplaceFunction                         `json:"typegen_224_DeleteMarketplaceFunction,omitempty"`
	Typegen225EnableMarketplaceFunction                        *EnableMarketplaceFunction                         `json:"typegen_225_EnableMarketplaceFunction,omitempty"`
	Typegen226DisableMarketplaceFunction                       *DisableMarketplaceFunction                        `json:"typegen_226_DisableMarketplaceFunction,omitempty"`
	Typegen227GetMarketplaceFunctionTokens                     *GetMarketplaceFunctionTokens                      `json:"typegen_227_GetMarketplaceFunctionTokens,omitempty"`
	Typegen228InvokeMarketplaceFunction                        *InvokeMarketplaceFunction                         `json:"typegen_228_InvokeMarketplaceFunction,omitempty"`
	Typegen232GetMarketplaceListing                            *GetMarketplaceListing                             `json:"typegen_232_GetMarketplaceListing,omitempty"`
	Typegen233TestMarketplaceIntegration                       *TestMarketplaceIntegration                        `json:"typegen_233_TestMarketplaceIntegration,omitempty"`
	Typegen234TestMarketplaceIntegrationResponse               *TestMarketplaceIntegrationResponse                `json:"typegen_234_TestMarketplaceIntegrationResponse,omitempty"`
	Typegen235GetMarketplaceListingResponse                    *GetMarketplaceListingResponse                     `json:"typegen_235_GetMarketplaceListingResponse,omitempty"`
	Typegen230AdminPortalStructureDto                          *AdminPortalStructureDto                           `json:"typegen_230_AdminPortalStructureDto,omitempty"`
	Typegen231AdminPortalModuleDto                             *AdminPortalModuleDto                              `json:"typegen_231_AdminPortalModuleDto,omitempty"`
	Typegen236UserMessageEntryWireDto                          *UserMessageEntryWireDto                           `json:"typegen_236_UserMessageEntryWireDto,omitempty"`
	Typegen237AssistantTextEntryWireDto                        *AssistantTextEntryWireDto                         `json:"typegen_237_AssistantTextEntryWireDto,omitempty"`
	Typegen238AssistantQuestionEntryWireDto                    *AssistantQuestionEntryWireDto                     `json:"typegen_238_AssistantQuestionEntryWireDto,omitempty"`
	Typegen239UserAnswerEntryWireDto                           *UserAnswerEntryWireDto                            `json:"typegen_239_UserAnswerEntryWireDto,omitempty"`
	Typegen240PlanEntryWireDto                                 *PlanEntryWireDto                                  `json:"typegen_240_PlanEntryWireDto,omitempty"`
	Typegen241UserDecisionEntryWireDto                         *UserDecisionEntryWireDto                          `json:"typegen_241_UserDecisionEntryWireDto,omitempty"`
	Typegen242RunStepEntryWireDto                              *RunStepEntryWireDto                               `json:"typegen_242_RunStepEntryWireDto,omitempty"`
	Typegen243ActionPendingEntryWireDto                        *ActionPendingEntryWireDto                         `json:"typegen_243_ActionPendingEntryWireDto,omitempty"`
	Typegen244NoticeEntryWireDto                               *NoticeEntryWireDto                                `json:"typegen_244_NoticeEntryWireDto,omitempty"`
	Typegen245ConversationSnapshotEntryWireDto                 *ConversationSnapshotEntryWireDto                  `json:"typegen_245_ConversationSnapshotEntryWireDto,omitempty"`
}

// @Route("/{version}/echo", "GET")
type Echo struct {
	RequestBase
}

// @Route("/{version}/public/projects/{ProjectId}/config", "GET")
type GetPublicProjectConfig struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
}

// @Route("/{version}/public/projects/{ProjectId}/legal/{Kind}", "GET")
type GetPublicProjectLegal struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
	Kind      *string `json:"kind,omitempty"`
}

// @Route("/{version}/account/profile", "GET")
type GetAccountProfile struct {
	RequestBase
}

// @Route("/{version}/account/profile", "PUT")
// @DataContract
type UpdateAccountProfile struct {
	RequestBase
	/** @description Account owner's display name. */
	// @DataMember
	// @ApiMember(Description="Account owner's display name.", IsRequired=true)
	DisplayName string `json:"displayName"`
	/** @description Email address used for billing communications. */
	// @DataMember
	// @ApiMember(Description="Email address used for billing communications.")
	BillingEmail *string `json:"billingEmail,omitempty"`
	/** @description Email address used for operations communications. */
	// @DataMember
	// @ApiMember(Description="Email address used for operations communications.")
	OperationsEmail *string `json:"operationsEmail,omitempty"`
	/** @description Email address used for security-related communications. */
	// @DataMember
	// @ApiMember(Description="Email address used for security-related communications.")
	SecurityEmail *string `json:"securityEmail,omitempty"`
}

// @Route("/{version}/account/verify/resend", "GET")
type ResendAccountVerificationToken struct {
	RequestBase
}

/** @description Get Account Status. */
// @Route("/{version}/account/status", "GET")
// @Api(Description="Get Account Status.")
type GetAccountStatus struct {
	RequestBase
}

// @Route("/{version}/account/stripe/create-checkout-session", "POST")
// @DataContract
type CreateStripeCheckoutSession struct {
	RequestBase
	// @DataMember
	SubscriptionType SubscriptionType `json:"subscriptionType,omitempty"`
	// @DataMember
	Domain *string `json:"domain,omitempty"`
	// @DataMember
	ProjectCap int `json:"projectCap,omitempty"`
	// @DataMember
	NewProjectSessionId *string `json:"newProjectSessionId,omitempty"`
	// @DataMember
	ReturnUrl *string `json:"returnUrl,omitempty"`
}

// @Route("/{version}/account/stripe/get-portal-url", "POST")
// @DataContract
type GetStripeBillingPortalUrl struct {
	RequestBase
	/** @description Which subscription (e.g. main account plan) to open the billing portal for. */
	// @DataMember
	// @ApiMember(Description="Which subscription (e.g. main account plan) to open the billing portal for.")
	SubscriptionType SubscriptionType `json:"subscriptionType,omitempty"`
	/** @description URL to return to after the customer leaves the billing portal. */
	// @DataMember
	// @ApiMember(Description="URL to return to after the customer leaves the billing portal.")
	ReturnUrl *string `json:"returnUrl,omitempty"`
}

// @Route("/{version}/account/team/member", "POST")
type CreateTeamMemberFromInvitation struct {
	RequestBase
	/** @description Display name of the account holder */
	// @ApiMember(DataType="string", Description="Display name of the account holder", IsRequired=true, Name="DisplayName", ParameterType="form")
	DisplayName string `json:"displayName"`
	/** @description Token from invitation email */
	// @ApiMember(DataType="string", Description="Token from invitation email", IsRequired=true, Name="Token", ParameterType="form")
	Token string `json:"token"`
	/** @description Set password for a new account */
	// @ApiMember(DataType="string", Description="Set password for a new account", Format="password", IsRequired=true, Name="Password", ParameterType="form")
	Password string `json:"password"`
}

/** @description Get Account Usage Billing. */
// @Route("/{version}/account/usage-billing", "GET")
// @Api(Description="Get Account Usage Billing.")
type GetAccountUsageBilling struct {
	RequestBase
}

// @Route("/{version}/account/verify", "GET")
type VerifyAccount struct {
	RequestBase
	Token     string `json:"token"`
	AccountId string `json:"accountId"`
}

// @Route("/{version}/account/projects/{projectId}/notifications/settings/group", "DELETE")
type DeleteNotificationsGroup struct {
	CodeMashRequestBase
	/** @description Tag identifying the notification group to remove. */
	// @ApiMember(Description="Tag identifying the notification group to remove.", IsRequired=true)
	GroupTag string `json:"groupTag"`
}

// @Route("/{version}/account/projects/{projectId}/notifications/settings/tag", "DELETE")
type DeleteNotificationsTag struct {
	CodeMashRequestBase
	/** @description Tag identifying the notification tag to delete. */
	// @ApiMember(Description="Tag identifying the notification tag to delete.", IsRequired=true)
	Tag string `json:"tag"`
}

// @Route("/{version}/account/projects/{projectId}/notifications/settings/group/tag", "DELETE")
type RemoveTagFromNotificationsGroup struct {
	CodeMashRequestBase
	/** @description Tag identifying the notification group. */
	// @ApiMember(Description="Tag identifying the notification group.", IsRequired=true)
	GroupTag string `json:"groupTag"`
	/** @description Tag identifying the notification tag to remove from the group. */
	// @ApiMember(Description="Tag identifying the notification tag to remove from the group.", IsRequired=true)
	Tag string `json:"tag"`
}

// @Route("/{version}/account/projects/{projectId}/notifications/settings/group", "POST")
type SaveNotificationsGroup struct {
	CodeMashRequestBase
	/** @description The group's tag and translations to save. The tag identifies the group; translations provide its display name per locale. */
	// @ApiMember(Description="The group's tag and translations to save. The tag identifies the group; translations provide its display name per locale.", IsRequired=true)
	GroupDefinition GroupDefinitionDto `json:"groupDefinition"`
	/** @description Communication channel (e.g. Email, Push) this group belongs to. */
	// @ApiMember(Description="Communication channel (e.g. Email, Push) this group belongs to.")
	Channel CommunicationChannel `json:"channel,omitempty"`
	/** @description If moving the group to a different channel, the channel it currently belongs to. */
	// @ApiMember(Description="If moving the group to a different channel, the channel it currently belongs to.")
	OriginChannel *CommunicationChannel `json:"originChannel,omitempty"`
}

// @Route("/{version}/account/projects/{projectId}/notifications/settings/tag", "POST")
type SaveNotificationsTag struct {
	CodeMashRequestBase
	/** @description The tag's identifier, translations, and default per-delivery-channel enabled/disabled settings. */
	// @ApiMember(Description="The tag's identifier, translations, and default per-delivery-channel enabled/disabled settings.", IsRequired=true)
	TagDefinition TagDefinitionDto `json:"tagDefinition"`
	/** @description Communication channel (e.g. Email, Push) this tag belongs to. */
	// @ApiMember(Description="Communication channel (e.g. Email, Push) this tag belongs to.")
	Channel *CommunicationChannel `json:"channel,omitempty"`
	/** @description Tag of the group this notification tag should be placed under, if any. */
	// @ApiMember(Description="Tag of the group this notification tag should be placed under, if any.")
	GroupTag *string `json:"groupTag,omitempty"`
}

/** @description Create a new backend project. */
// @Route("/{version}/account/projects", "POST")
// @Api(Description="Create a new backend project.")
// @DataContract
type CreateProjectRequest struct {
	RequestBase
	// @DataMember
	Integration DatabaseIntegrationRequest `json:"integration"`
	/** @description Project name, unique per account. */
	// @DataMember
	// @ApiMember(Description="Project name, unique per account.")
	ProjectName string `json:"projectName"`
	/** @description Region code for the primary region, e.g. 'nb-eu-germany'. Use a code from get_account_regions. */
	// @DataMember
	// @ApiMember(Description="Region code for the primary region, e.g. 'nb-eu-germany'. Use a code from get_account_regions.")
	PrimaryRegion *string `json:"primaryRegion,omitempty"`
	// @DataMember
	AdditionalRegions []string `json:"additionalRegions,omitempty"`
	// @DataMember
	Description *string `json:"description,omitempty"`
}

/** @description Deletes project */
// @Route("/{version}/account/projects/{projectId}", "DELETE")
// @Api(Description="Deletes project")
type DeleteProject struct {
	CodeMashRequestBase
}

// @Route("/{version}/account/projects/environments", "POST")
// @DataContract
type CreateProjectEnvironmentRequest struct {
	CodeMashRequestBase
	/** @description Name for the new environment (e.g. 'TEST', 'STAGING'). A-Z/0-9/space, up to 15 chars, cannot be PROD. */
	// @DataMember
	// @ApiMember(Description="Name for the new environment (e.g. 'TEST', 'STAGING'). A-Z/0-9/space, up to 15 chars, cannot be PROD.", IsRequired=true)
	EnvironmentName string `json:"environmentName"`
	// @DataMember(Name="integration")
	Integration DatabaseIntegrationRequest `json:"integration"`
}

// @Route("/{version}/account/projects/environments/{environmentName}", "DELETE")
// @DataContract
type DeleteProjectEnvironmentRequest struct {
	CodeMashRequestBase
	/** @description Name of the environment to delete (from get_project_environments), e.g. 'TEST'. PROD is rejected. */
	// @DataMember
	// @ApiMember(Description="Name of the environment to delete (from get_project_environments), e.g. 'TEST'. PROD is rejected.", IsRequired=true)
	EnvironmentName string `json:"environmentName"`
}

// @Route("/{version}/account/projects/environments/{environmentName}/rank", "PATCH")
// @DataContract
type SetEnvironmentRankRequest struct {
	CodeMashRequestBase
	/** @description Name of the environment to re-rank (from get_project_environments). */
	// @DataMember
	// @ApiMember(Description="Name of the environment to re-rank (from get_project_environments).", IsRequired=true)
	EnvironmentName string `json:"environmentName"`
	/** @description New promotion-ladder rank. Out-of-range values are clamped and ranks re-normalized. */
	// @DataMember
	// @ApiMember(Description="New promotion-ladder rank. Out-of-range values are clamped and ranks re-normalized.")
	Rank int `json:"rank,omitempty"`
}

// @Route("/{version}/account/projects/environments/promote", "POST")
// @DataContract
type PromoteEnvironmentRequest struct {
	CodeMashRequestBase
	/** @description Environment to promote FROM (source of truth for this promotion). */
	// @DataMember
	// @ApiMember(Description="Environment to promote FROM (source of truth for this promotion).", IsRequired=true)
	SourceEnv string `json:"sourceEnv"`
	/** @description Environment to promote INTO. Must be higher on the promotion ladder than SourceEnv. */
	// @DataMember
	// @ApiMember(Description="Environment to promote INTO. Must be higher on the promotion ladder than SourceEnv.", IsRequired=true)
	TargetEnv string `json:"targetEnv"`
	/** @description When true, only returns the promotion plan (including deletions) without applying it or copying secrets. Use this to preview before a real run. */
	// @DataMember
	// @ApiMember(Description="When true, only returns the promotion plan (including deletions) without applying it or copying secrets. Use this to preview before a real run.")
	DryRun bool `json:"dryRun,omitempty"`
}

// @Route("/{version}/account/projects/environments/promote/rollback", "POST")
// @DataContract
type RollbackPromotionRequest struct {
	CodeMashRequestBase
	/** @description Environment whose content should be rolled back. */
	// @DataMember
	// @ApiMember(Description="Environment whose content should be rolled back.", IsRequired=true)
	TargetEnv string `json:"targetEnv"`
	/** @description The fromVersion anchor returned by the promote_environment call being rolled back. */
	// @DataMember
	// @ApiMember(Description="The fromVersion anchor returned by the promote_environment call being rolled back.", IsRequired=true)
	FromVersion int64 `json:"fromVersion"`
}

// @Route("/{version}/account/projects/environments", "GET")
// @DataContract
type GetProjectEnvironments struct {
	CodeMashRequestBase
}

/** @description Gets project info. */
// @Route("/{version}/account/projects/{projectId}", "GET")
// @Api(Description="Gets project info.")
type GetProject struct {
	CodeMashRequestBase
}

/** @description Retrieve projects list. */
// @Route("/{version}/account/projects", "GET")
// @Api(Description="Retrieve projects list.")
type GetProjects struct {
	RequestBase
}

/** @description Get available project regions. */
// @Route("/{version}/account/regions", "GET")
// @Api(Description="Get available project regions.")
type GetAccountRegions struct {
	RequestBase
}

/** @description Waits (bounded, server-side) for a project to finish provisioning and become active. */
// @Route("/{version}/account/projects/{projectId}/wait-active", "GET")
// @Api(Description="Waits (bounded, server-side) for a project to finish provisioning and become active.")
type WaitForProjectActiveRequest struct {
	CodeMashRequestBase
	/** @description Max seconds to wait before returning 'not active yet' (default 30, capped at 90). */
	// @ApiMember(Description="Max seconds to wait before returning 'not active yet' (default 30, capped at 90).")
	TimeoutSeconds *int `json:"timeoutSeconds,omitempty"`
}

/** @description Gets project tokens. */
// @Route("/{version}/account/projects/{projectId}/tokens", "GET")
// @Api(Description="Gets project tokens.")
type GetProjectTokens struct {
	CodeMashRequestBase
	InitiatorId                *string `json:"initiatorId,omitempty"`
	RecipientId                *string `json:"recipientId,omitempty"`
	TargetUserId               *string `json:"targetUserId,omitempty"`
	MembershipTriggerOldUserId *string `json:"membershipTriggerOldUserId,omitempty"`
	MembershipTriggerNewUserId *string `json:"membershipTriggerNewUserId,omitempty"`
}

/** @description Assigns the project's Admin Portal service user */
// @Route("/{version}/account/projects/{projectId}/settings/admin-portal/service-user", "PUT")
// @Api(Description="Assigns the project's Admin Portal service user")
type AssignAdminPortalServiceUserRequest struct {
	CodeMashRequestBase
	/** @description Id of the existing AuthType.Service user to assign as the project's Admin Portal service user. */
	// @ApiMember(Description="Id of the existing AuthType.Service user to assign as the project's Admin Portal service user.", IsRequired=true)
	ServiceUserId string `json:"serviceUserId"`
}

/** @description Reads the Admin Portal layout/structure (service user only) */
// @Route("/{version}/account/projects/{projectId}/admin-portal/structure", "GET")
// @Api(Description="Reads the Admin Portal layout/structure (service user only)")
type GetAdminPortalStructure struct {
	CodeMashRequestBase
}

/** @description Updates the project's admin-portal URL override */
// @Route("/{version}/account/projects/{projectId}/settings/admin-url", "PATCH")
// @Api(Description="Updates the project's admin-portal URL override")
type UpdateProjectAdminUrl struct {
	CodeMashRequestBase
	/** @description Custom admin-portal URL to use instead of the canonical address. Null/empty restores the canonical pr_{id}.admin.{host} address. */
	// @ApiMember(Description="Custom admin-portal URL to use instead of the canonical address. Null/empty restores the canonical pr_{id}.admin.{host} address.")
	Url *string `json:"url,omitempty"`
}

/** @description Updates project accent color */
// @Route("/{version}/account/projects/{projectId}/settings/accent-color", "PATCH")
// @Api(Description="Updates project accent color")
type UpdateProjectAccentColor struct {
	CodeMashRequestBase
	/** @description Hex color code, e.g. '#FF6D00'. */
	// @ApiMember(Description="Hex color code, e.g. '#FF6D00'.", IsRequired=true)
	Color string `json:"color"`
}

/** @description Updates project icon */
// @Route("/{version}/account/projects/{projectId}/settings/icon", "PATCH")
// @Api(Description="Updates project icon")
type UpdateProjectIcon struct {
	CodeMashRequestBase
	FileResource *FileResourceRefDto `json:"fileResource,omitempty"`
}

/** @description Updates project logo */
// @Route("/{version}/account/projects/{projectId}/settings/logo", "PATCH")
// @Api(Description="Updates project logo")
type UpdateProjectLogo struct {
	CodeMashRequestBase
	FileResource *FileResourceRefDto `json:"fileResource,omitempty"`
}

/** @description Updates project main color */
// @Route("/{version}/account/projects/{projectId}/settings/main-color", "PATCH")
// @Api(Description="Updates project main color")
type UpdateProjectMainColor struct {
	CodeMashRequestBase
	/** @description Hex color code, e.g. '#1A73E8'. */
	// @ApiMember(Description="Hex color code, e.g. '#1A73E8'.", IsRequired=true)
	Color string `json:"color"`
}

/** @description Updates project CORS settings */
// @Route("/{version}/account/projects/{projectId}/settings/origins", "PATCH")
// @Api(Description="Updates project CORS settings")
type UpdateProjectAllowedOrigins struct {
	CodeMashRequestBase
	/** @description The complete new list of allowed origin URLs, e.g. ["https://app.example.com", "https://example.com"]. An entry with no scheme (e.g. "example.com") defaults to https. Whatever is not in this list stops being allowed. */
	// @ApiMember(Description="The complete new list of allowed origin URLs, e.g. [\"https://app.example.com\", \"https://example.com\"]. An entry with no scheme (e.g. \"example.com\") defaults to https. Whatever is not in this list stops being allowed.")
	Origins []string `json:"origins,omitempty"`
}

/** @description Update project default language */
// @Route("/{version}/account/projects/{projectId}/settings/default-language", "PATCH")
// @Api(Description="Update project default language")
type UpdateProjectDefaultLanguage struct {
	CodeMashRequestBase
	/** @description Language code, e.g. 'en' or 'de'. */
	// @ApiMember(Description="Language code, e.g. 'en' or 'de'.", IsRequired=true)
	DefaultLanguage string `json:"defaultLanguage"`
}

/** @description Updates project description */
// @Route("/{version}/account/projects/{projectId}/settings/description", "PATCH")
// @Api(Description="Updates project description")
type UpdateProjectDescription struct {
	CodeMashRequestBase
	/** @description The new description text. Omit (null) to clear the description. */
	// @ApiMember(Description="The new description text. Omit (null) to clear the description.")
	Description *string `json:"description,omitempty"`
}

/** @description Disables project */
// @Route("/{version}/account/projects/{projectId}/disable", "PATCH")
// @Api(Description="Disables project")
type DisableProject struct {
	CodeMashRequestBase
}

/** @description Enables project */
// @Route("/{version}/account/projects/{projectId}/enable", "PATCH")
// @Api(Description="Enables project")
type EnableProject struct {
	CodeMashRequestBase
}

/** @description Updates project languages */
// @Route("/{version}/account/projects/{projectId}/settings/languages", "PATCH")
// @Api(Description="Updates project languages")
type UpdateProjectLanguages struct {
	CodeMashRequestBase
	/** @description The complete new list of language codes, e.g. ["en", "de", "lt"]. */
	// @ApiMember(Description="The complete new list of language codes, e.g. [\"en\", \"de\", \"lt\"].", IsRequired=true)
	Languages []string `json:"languages"`
}

/** @description Updates the project's public legal documents (Terms & Conditions, Privacy Policy) */
// @Route("/{version}/account/projects/{projectId}/settings/legal", "PATCH")
// @Api(Description="Updates the project's public legal documents (Terms & Conditions, Privacy Policy)")
type UpdateProjectLegalDocuments struct {
	CodeMashRequestBase
	/** @description Terms & Conditions document, Markdown. Null/empty clears it. */
	// @ApiMember(Description="Terms & Conditions document, Markdown. Null/empty clears it.")
	TermsMarkdown *string `json:"termsMarkdown,omitempty"`
	/** @description Privacy Policy document, Markdown. Null/empty clears it. */
	// @ApiMember(Description="Privacy Policy document, Markdown. Null/empty clears it.")
	PrivacyMarkdown *string `json:"privacyMarkdown,omitempty"`
}

/** @description Sets whether the project's legal documents are publicly readable via the Admin Portal */
// @Route("/{version}/account/projects/{projectId}/settings/legal/expose", "PATCH")
// @Api(Description="Sets whether the project's legal documents are publicly readable via the Admin Portal")
type UpdateProjectExposeLegal struct {
	CodeMashRequestBase
	/** @description True to make the legal documents publicly readable via the Admin Portal, false to hide them. */
	// @ApiMember(Description="True to make the legal documents publicly readable via the Admin Portal, false to hide them.")
	Exposed bool `json:"exposed,omitempty"`
}

/** @description Updates project marketing url */
// @Route("/{version}/account/projects/{projectId}/settings/url", "PATCH")
// @Api(Description="Updates project marketing url")
type UpdateProjectUrl struct {
	CodeMashRequestBase
	/** @description The marketing site URL, e.g. 'https://example.com'. Omit (null) to clear. */
	// @ApiMember(Description="The marketing site URL, e.g. 'https://example.com'. Omit (null) to clear.")
	Url *string `json:"url,omitempty"`
}

/** @description Updates project name */
// @Route("/{version}/account/projects/{projectId}/settings/name", "PATCH")
// @Api(Description="Updates project name")
type UpdateProjectName struct {
	CodeMashRequestBase
	/** @description The new project name, unique per account. */
	// @ApiMember(Description="The new project name, unique per account.", IsRequired=true)
	Name string `json:"name"`
}

/** @description Updates project regions */
// @Route("/{version}/account/projects/{projectId}/settings/regions", "PATCH")
// @Api(Description="Updates project regions")
type UpdateProjectRegions struct {
	CodeMashRequestBase
	/** @description Primary region code, e.g. 'nb-eu-germany'. Immutable once set — omit to keep the current one; only set it on a project that has none. */
	// @ApiMember(Description="Primary region code, e.g. 'nb-eu-germany'. Immutable once set — omit to keep the current one; only set it on a project that has none.")
	PrimaryRegion *string `json:"primaryRegion,omitempty"`
	/** @description The complete new list of additional region codes (full replacement). A region that still hosts a provisioned database cluster cannot be removed. */
	// @ApiMember(Description="The complete new list of additional region codes (full replacement). A region that still hosts a provisioned database cluster cannot be removed.")
	AdditionalRegions []string `json:"additionalRegions,omitempty"`
}

/** @description This API endpoint allows users to create a new CodeMash account. */
// @Route("/{version}/account", "POST")
// @Api(Description="This API endpoint allows users to create a new CodeMash account.")
type CreateAccount struct {
	RequestBase
	/** @description Display name of the account holder */
	// @ApiMember(DataType="string", Description="Display name of the account holder", IsRequired=true, Name="DisplayName", ParameterType="form")
	DisplayName string `json:"displayName"`
	/** @description Real email of account holder */
	// @ApiMember(DataType="string", Description="Real email of account holder", IsRequired=true, Name="Email", ParameterType="form")
	Email string `json:"email"`
	/** @description Set password for a new account */
	// @ApiMember(DataType="string", Description="Set password for a new account", Format="password", IsRequired=true, Name="Password", ParameterType="form")
	Password string `json:"password"`
}

// @Route("/{version}/account/team/member/password", "POST")
type ChangeTeamMemberPassword struct {
	RequestBase
	Email           string `json:"email"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// @Route("/{version}/account/team/member/create", "POST")
type CreateTeamMember struct {
	RequestBase
	Email          string   `json:"email"`
	DisplayName    *string  `json:"displayName,omitempty"`
	Password       *string  `json:"password,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	SendInvitation bool     `json:"sendInvitation,omitempty"`
}

// @Route("/{version}/account/team/policies", "POST")
type CreateAccountPolicy struct {
	RequestBase
	/** @description Name for the new custom account policy. */
	// @ApiMember(Description="Name for the new custom account policy.", IsRequired=true)
	PolicyName string `json:"policyName"`
	/** @description Optional human-readable description of the policy's purpose. */
	// @ApiMember(Description="Optional human-readable description of the policy's purpose.")
	Description *string `json:"description,omitempty"`
	/** @description Raw JSON policy document, AWS-IAM style (Statement array of Effect/Action/Resource entries), matching PolicyDocument.schema.json. This defines which permissions the policy grants. */
	// @ApiMember(Description="Raw JSON policy document, AWS-IAM style (Statement array of Effect/Action/Resource entries), matching PolicyDocument.schema.json. This defines which permissions the policy grants.")
	PolicyDocumentJson *string `json:"policyDocumentJson,omitempty"`
}

// @Route("/{version}/account/team/roles", "POST")
type CreateAccountRole struct {
	RequestBase
	/** @description Name for the new custom account team role. */
	// @ApiMember(Description="Name for the new custom account team role.", IsRequired=true)
	RoleName string `json:"roleName"`
	/** @description Optional human-readable description of the role's purpose. */
	// @ApiMember(Description="Optional human-readable description of the role's purpose.")
	Description *string `json:"description,omitempty"`
	/** @description Public policy ids (from get_account_team_policies) to attach to this role. */
	// @ApiMember(Description="Public policy ids (from get_account_team_policies) to attach to this role.")
	Policies []string `json:"policies,omitempty"`
}

// @Route("/{version}/account/team/policies/{Id}", "DELETE")
type DeleteAccountPolicy struct {
	RequestBase
	/** @description Public policy id (from get_account_team_policies) to delete. */
	// @ApiMember(Description="Public policy id (from get_account_team_policies) to delete.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/account/team/roles/{Id}", "DELETE")
type DeleteAccountRole struct {
	RequestBase
	/** @description Role template id (from get_account_team_roles) to delete. */
	// @ApiMember(Description="Role template id (from get_account_team_roles) to delete.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets account team members (collaborators) */
// @Route("/{version}/account/collaborators", "GET")
// @Api(Description="Gets account team members (collaborators)")
type GetAccountCollaborators struct {
	RequestBase
	/** @description Set true to also include the account owner in the list. */
	// @ApiMember(Description="Set true to also include the account owner in the list.")
	IncludeAccountOwner bool `json:"includeAccountOwner,omitempty"`
	/** @description Set true to only return members that have a registered push device. */
	// @ApiMember(Description="Set true to only return members that have a registered push device.")
	UserShouldHavePushDevice bool `json:"userShouldHavePushDevice,omitempty"`
	/** @description Optional project id — only members with access to that project. */
	// @ApiMember(Description="Optional project id — only members with access to that project.")
	ProjectId *string `json:"projectId,omitempty"`
	/** @description Optional filter: only these user ids. */
	// @ApiMember(Description="Optional filter: only these user ids.")
	UserIds []string `json:"userIds,omitempty"`
	/** @description Optional filter: only members having one of these role names. */
	// @ApiMember(Description="Optional filter: only members having one of these role names.")
	RoleNames  []string    `json:"roleNames,omitempty"`
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// @Route("/{version}/account/team/password-policy", "GET")
type GetAccountPasswordPolicy struct {
	RequestBase
}

// @Route("/{version}/account/team/policies", "GET")
type GetAccountTeamPolicies struct {
	RequestBase
}

// @Route("/{version}/account/team/roles", "GET")
type GetAccountTeamRoles struct {
	RequestBase
}

/** @description Send invite to team member */
// @Route("/{version}/account/team/member/invite", "POST")
// @Api(Description="Send invite to team member")
type SendInviteToTeamMember struct {
	RequestBase
	/** @description Email address the invitation is sent to. */
	// @ApiMember(Description="Email address the invitation is sent to.", IsRequired=true)
	Email string `json:"email"`
	/** @description Account role ids (GUIDs from get_account_team_roles) the member gets on accepting. Omit for the default member role. */
	// @ApiMember(Description="Account role ids (GUIDs from get_account_team_roles) the member gets on accepting. Omit for the default member role.")
	Roles []string `json:"roles,omitempty"`
}

// @Route("/{version}/account/team/policies", "PUT")
type UpdateAccountPolicy struct {
	RequestBase
	/** @description Public policy id (from get_account_team_policies) to update. */
	// @ApiMember(Description="Public policy id (from get_account_team_policies) to update.", IsRequired=true)
	Id string `json:"id"`
	/** @description New name for the policy. */
	// @ApiMember(Description="New name for the policy.", IsRequired=true)
	PolicyName string `json:"policyName"`
	/** @description Optional human-readable description of the policy's purpose. */
	// @ApiMember(Description="Optional human-readable description of the policy's purpose.")
	Description *string `json:"description,omitempty"`
	/** @description Raw JSON policy document, AWS-IAM style (Statement array of Effect/Action/Resource entries) — replaces the policy's current statement set. */
	// @ApiMember(Description="Raw JSON policy document, AWS-IAM style (Statement array of Effect/Action/Resource entries) — replaces the policy's current statement set.")
	PolicyDocumentJson *string `json:"policyDocumentJson,omitempty"`
}

// @Route("/{version}/account/team/roles", "PUT")
type UpdateAccountRole struct {
	RequestBase
	/** @description Role template id (from get_account_team_roles) to update. */
	// @ApiMember(Description="Role template id (from get_account_team_roles) to update.", IsRequired=true)
	Id string `json:"id"`
	/** @description New name for the role. */
	// @ApiMember(Description="New name for the role.", IsRequired=true)
	RoleName string `json:"roleName"`
	/** @description Optional human-readable description of the role's purpose. */
	// @ApiMember(Description="Optional human-readable description of the role's purpose.")
	Description *string `json:"description,omitempty"`
	/** @description Public policy ids (from get_account_team_policies) to attach — replaces the current set. */
	// @ApiMember(Description="Public policy ids (from get_account_team_policies) to attach — replaces the current set.")
	Policies []string `json:"policies,omitempty"`
}

// @Route("/{version}/account/userauth/has-passkey", "POST")
// @DataContract
type AccountHasPasskeyRequest struct {
	RequestBase
	// @DataMember
	Email string `json:"email"`
}

// @Route("/{version}/account/userauth/email/start-verification", "POST")
// @DataContract
type AccountStartEmailVerificationRequest struct {
	RequestBase
	// @DataMember
	Email string `json:"email"`
}

// @Route("/{version}/account/userauth/email/confirm-verification", "POST")
// @DataContract
type AccountConfirmEmailVerificationRequest struct {
	RequestBase
	// @DataMember
	Email string `json:"email"`
	// @DataMember
	Code string `json:"code"`
}

// @Route("/{version}/account/userauth/passkey/registration-options", "POST")
// @DataContract
type AccountPasskeyRegistrationOptionsRequest struct {
	RequestBase
	// @DataMember
	VerificationToken string `json:"verificationToken"`
}

// @Route("/{version}/account/userauth/passkey/verify-registration", "POST")
// @DataContract
type AccountVerifyPasskeyRegistrationRequest struct {
	RequestBase
	// @DataMember
	VerificationToken string `json:"verificationToken"`
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AttestationResponse string `json:"attestationResponse"`
	// @DataMember
	FriendlyName *string `json:"friendlyName,omitempty"`
}

// @Route("/{version}/account/userauth/passkey/authentication-options", "POST")
// @DataContract
type AccountPasskeyAuthenticationOptionsRequest struct {
	RequestBase
	// @DataMember
	Email string `json:"email"`
}

// @Route("/{version}/account/userauth/passkey/verify-authentication", "POST")
// @DataContract
type AccountVerifyPasskeyAuthenticationRequest struct {
	RequestBase
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AssertionResponse string `json:"assertionResponse"`
}

// @Route("/{version}/account/userauth/passkeys", "GET")
// @DataContract
type ListAccountPasskeysRequest struct {
	RequestBase
}

// @Route("/{version}/account/userauth/passkeys/{CredentialId}/rename", "POST")
// @DataContract
type RenameAccountPasskeyRequest struct {
	RequestBase
	// @DataMember
	CredentialId string `json:"credentialId"`
	// @DataMember
	FriendlyName string `json:"friendlyName"`
}

// @Route("/{version}/account/userauth/passkeys/{CredentialId}/revoke", "POST")
// @DataContract
type RevokeAccountPasskeyRequest struct {
	RequestBase
	// @DataMember
	CredentialId string `json:"credentialId"`
}

// @Route("/{version}/account/userauth/passkey/enrollment-options", "POST")
// @DataContract
type AccountPasskeyEnrollmentOptionsRequest struct {
	RequestBase
}

// @Route("/{version}/account/userauth/passkey/verify-enrollment", "POST")
// @DataContract
type AccountVerifyPasskeyEnrollmentRequest struct {
	RequestBase
	// @DataMember
	CeremonyId string `json:"ceremonyId"`
	// @DataMember
	AttestationResponse string `json:"attestationResponse"`
	// @DataMember
	FriendlyName *string `json:"friendlyName,omitempty"`
}

// @Route("/{version}/account/licensing/dns-status", "GET")
type GetLicenseDomainDnsStatus struct {
	RequestBase
	Domain *string `json:"domain,omitempty"`
}

// @Route("/{version}/licensing/domain-verification/start", "POST")
type StartLicenseDomainVerificationRequest struct {
	RequestBase
	Domain *string `json:"domain,omitempty"`
}

// @Route("/{version}/licensing/domain-verification/status", "GET")
type GetLicenseDomainVerificationStatus struct {
	RequestBase
	Domain *string `json:"domain,omitempty"`
}

// @Route("/{version}/account/licenses", "GET")
type GetLicenses struct {
	RequestBase
}

// @Route("/{version}/licensing/heartbeat", "POST")
type PostLicenseHeartbeat struct {
	RequestBase
	License          *string `json:"license,omitempty"`
	LicenseAccountId *string `json:"licenseAccountId,omitempty"`
	InstallationId   *string `json:"installationId,omitempty"`
	Domain           *string `json:"domain,omitempty"`
	HostKind         *string `json:"hostKind,omitempty"`
	Release          *string `json:"release,omitempty"`
	InstanceVersion  *string `json:"instanceVersion,omitempty"`
}

// @Route("/{version}/account/licensing/status", "GET")
type GetInstallationLicenseStatus struct {
	RequestBase
}

type AccountCreated struct {
	Email       EmailAddress `json:"email"`
	DisplayName DisplayName  `json:"displayName"`
	AccountId   AccountId    `json:"accountId"`
	CreatedOn   UtcDateTime  `json:"createdOn"`
}

type AccountProfileUpdated struct {
	DisplayName     DisplayName   `json:"displayName"`
	BillingEmail    *EmailAddress `json:"billingEmail,omitempty"`
	OperationsEmail *EmailAddress `json:"operationsEmail,omitempty"`
	SecurityEmail   *EmailAddress `json:"securityEmail,omitempty"`
}

type AccountSetAsActive struct {
}

type AccountValidationTokenIssued struct {
	Expiration ExpirationToken `json:"expiration"`
}

type AccountVerified struct {
}

type AccountBlocked struct {
}

type AccountSetAsInactive struct {
}

type AccountUnregistered struct {
}

type LicenseCreated struct {
	License CodeMashLicense `json:"license"`
}

type ProjectCreated struct {
	Id                    ProjectId       `json:"id"`
	Name                  ProjectName     `json:"name"`
	DatabaseIntegrationId IntegrationId   `json:"databaseIntegrationId"`
	PrimaryRegion         *ProjectRegion  `json:"primaryRegion,omitempty"`
	AdditionalRegions     []ProjectRegion `json:"additionalRegions,omitempty"`
	Description           *string         `json:"description,omitempty"`
	IsProvisioning        bool            `json:"isProvisioning,omitempty"`
}

type ProjectActivated struct {
}

type ProjectSuspendedByLicense struct {
}

type ProjectResumedFromLicenseSuspension struct {
}

type ProjectDisabled struct {
}

type ProjectDeleted struct {
}

type ProjectNameChanged struct {
	ProjectName ProjectName `json:"projectName"`
}

type ProjectDescriptionChanged struct {
	Description *string `json:"description,omitempty"`
}

type ProjectMarketingUrlChanged struct {
	Url *DomainUrl `json:"url,omitempty"`
}

type ProjectAdminUrlChanged struct {
	Url *DomainUrl `json:"url,omitempty"`
}

type ProjectLegalDocumentsChanged struct {
	Documents ProjectLegalDocuments `json:"documents"`
}

type ProjectExposeLegalToAdminPortalChanged struct {
	Exposed bool `json:"exposed,omitempty"`
}

type ProjectAdminPortalServiceUserAssigned struct {
	ServiceUserId AuthId `json:"serviceUserId"`
}

type ProjectAllowedOriginsChanged struct {
	Origins []DomainUrl `json:"origins,omitempty"`
}

type ProjectEnvironmentCreated struct {
	Env   Env            `json:"env"`
	Ranks map[string]int `json:"ranks"`
}

type ProjectEnvironmentDeleted struct {
	Env   Env            `json:"env"`
	Ranks map[string]int `json:"ranks"`
}

type ProjectEnvironmentRanksChanged struct {
	Ranks map[string]int `json:"ranks"`
}

type ProjectDefaultLanguageChanged struct {
	Language Language `json:"language"`
}

type ProjectLanguagesChanged struct {
	Languages []Language `json:"languages"`
}

type ProjectLogoChanged struct {
	Logo *ProjectLogo `json:"logo,omitempty"`
}

type ProjectIconChanged struct {
	Icon *ProjectIcon `json:"icon,omitempty"`
}

type ProjectMainColorChanged struct {
	Color BrandColor `json:"color"`
}

type ProjectAccentColorChanged struct {
	Color BrandColor `json:"color"`
}

type ProjectRegionsChanged struct {
	PrimaryRegion     *ProjectRegion  `json:"primaryRegion,omitempty"`
	AdditionalRegions []ProjectRegion `json:"additionalRegions,omitempty"`
}

type ProjectCommunicationSet struct {
	ProjectCommunication ProjectCommunication `json:"projectCommunication"`
}

type ProjectTimeZoneChanged struct {
	TimeZone *TimeZone `json:"timeZone,omitempty"`
}

type ProjectPaymentZonesChanged struct {
	PaymentZones []TimeZone `json:"paymentZones,omitempty"`
}

type ProjectCommunicationGroupSaved struct {
	Group         GroupDefinition       `json:"group"`
	Channel       CommunicationChannel  `json:"channel,omitempty"`
	OriginChannel *CommunicationChannel `json:"originChannel,omitempty"`
}

type ProjectCommunicationTagFromGroupDeleted struct {
	GroupTag   Tag `json:"groupTag"`
	RemovedTag Tag `json:"removedTag"`
}

type ProjectCommunicationGroupDeleted struct {
	GroupTag Tag `json:"groupTag"`
}

type ProjectCommunicationTagSaved struct {
	Tag      TagDefinition         `json:"tag"`
	GroupTag *Tag                  `json:"groupTag,omitempty"`
	Channel  *CommunicationChannel `json:"channel,omitempty"`
}

type ProjectCommunicationTagDeleted struct {
	Tag Tag `json:"tag"`
}

type CustomerCreated struct {
	PaymentCustomerRef PaymentCustomerRef `json:"paymentCustomerRef"`
}

type SubscriptionChanged struct {
	Subscription CodeMashManagedServiceSubscription `json:"subscription"`
}

type SubscriptionCanceled struct {
	PaymentCustomerRef PaymentCustomerRef `json:"paymentCustomerRef"`
	SubscriptionId     string             `json:"subscriptionId"`
}

type AccountTeamPolicyCreated struct {
	Policy MembershipPolicy `json:"policy"`
}

type AccountTeamPolicyUpdated struct {
	Policy MembershipPolicy `json:"policy"`
}

type AccountTeamPolicyDeleted struct {
	PolicyId PolicyId `json:"policyId"`
}

type AccountTeamRoleCreated struct {
	Role MembershipRole `json:"role"`
}

type AccountTeamRoleUpdated struct {
	Role MembershipRole `json:"role"`
}

type AccountTeamRoleDeleted struct {
	RoleId RoleId `json:"roleId"`
}

type AtlasUsageRecorded struct {
	Record AtlasUsageRecord `json:"record"`
}

type UsageBillingIngestionFailed struct {
	Failure UsageIngestionFailure `json:"failure"`
}

// @Route("/{version}/membership/disable", "GET")
type DisableMembership struct {
	CodeMashRequestBase
}

// @Route("/{version}/membership/enable", "GET")
type EnableMembership struct {
	CodeMashRequestBase
}

/** @description Membership */
// @Route("/{version}/membership/users/{Id}/api-keys", "POST")
// @Api(Description="Membership")
// @DataContract
type IssueServiceUserApiKeyRequest struct {
	CodeMashRequestBase
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Scopes []string `json:"scopes,omitempty"`
	// @DataMember
	ExpiresInDays *int `json:"expiresInDays,omitempty"`
	// @DataMember
	Notes *string `json:"notes,omitempty"`
}

/** @description Membership */
// @Route("/{version}/membership/users/{Id}/api-keys", "GET")
// @Api(Description="Membership")
// @DataContract
type ListServiceUserApiKeysRequest struct {
	CodeMashRequestBase
	/** @description The service user's auth id. */
	// @DataMember
	// @ApiMember(Description="The service user's auth id.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Membership */
// @Route("/{version}/membership/users/{Id}/api-keys/{KeyId}", "DELETE")
// @Api(Description="Membership")
// @DataContract
type DeleteServiceUserApiKeyRequest struct {
	CodeMashRequestBase
	/** @description The service user's auth id. */
	// @DataMember
	// @ApiMember(Description="The service user's auth id.", IsRequired=true)
	Id string `json:"id"`
	/** @description The key id to delete, from list_service_user_api_keys. */
	// @DataMember
	// @ApiMember(Description="The key id to delete, from list_service_user_api_keys.", IsRequired=true)
	KeyId int `json:"keyId"`
}

// @Route("/{version}/membership/triggers/{triggerId}", "DELETE")
// @DataContract
type DeleteMembershipTrigger struct {
	DeleteTrigger
}

// @Route("/{version}/membership/triggers/{triggerId}/disable", "PATCH")
// @DataContract
type DisableMembershipTrigger struct {
	DisableTrigger
}

// @Route("/{version}/membership/triggers/{triggerId}/enable", "PATCH")
// @DataContract
type EnableMembershipTrigger struct {
	EnableTrigger
}

/** @description Gets membership trigger by specified Id */
// @Route("/{version}/membership/triggers/{id}", "GET")
// @Api(Description="Gets membership trigger by specified Id")
type GetMembershipTrigger struct {
	GetTrigger
}

/** @description Gets membership triggers */
// @Route("/{version}/membership/triggers", "GET")
// @Api(Description="Gets membership triggers")
type GetMembershipTriggers struct {
	GetTriggers
}

// @Route("/{version}/membership/triggers", "POST")
// @DataContract
type SaveMembershipTrigger struct {
	SaveTrigger
}

/** @description Create a new custom role for project. */
// @Route("/{version}/membership/roles", "POST")
// @Api(Description="Create a new custom role for project.")
type CreateRole struct {
	CodeMashRequestBase
	/** @description Display name of the new role, unique within the project. */
	// @ApiMember(Description="Display name of the new role, unique within the project.", IsRequired=true)
	RoleName    string  `json:"roleName"`
	Description *string `json:"description,omitempty"`
	/** @description Policy ids to attach. These are OPAQUE ids from get_policies (e.g. 'pol_3kJ9xJ2mQ0aBcDeFgHiJk') — NEVER invent them or guess from a policy name. Omit this to create a role with no policies and attach them later. */
	// @ApiMember(Description="Policy ids to attach. These are OPAQUE ids from get_policies (e.g. 'pol_3kJ9xJ2mQ0aBcDeFgHiJk') — NEVER invent them or guess from a policy name. Omit this to create a role with no policies and attach them later.")
	Policies []string `json:"policies,omitempty"`
}

/** @description Deletes custom role from project. */
// @Route("/{version}/membership/roles", "DELETE")
// @Api(Description="Deletes custom role from project.")
type DeleteRole struct {
	CodeMashRequestBase
	/** @description Id of the role to delete, from get_roles. */
	// @ApiMember(Description="Id of the role to delete, from get_roles.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets project role details. */
// @Route("/{version}/membership/roles/{Id}", "GET")
// @Api(Description="Gets project role details.")
type GetRole struct {
	CodeMashRequestBase
	/** @description Role id from get_roles. */
	// @ApiMember(Description="Role id from get_roles.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets project roles. */
// @Route("/{version}/membership/roles", "GET")
// @Api(Description="Gets project roles.")
type GetRoles struct {
	CodeMashRequestBase
}

/** @description Updates role policies */
// @Route("/{version}/membership/roles", "PATCH")
// @Api(Description="Updates role policies")
type UpdateRolePolicies struct {
	CodeMashRequestBase
	/** @description Id of the role to update, from get_roles. */
	// @ApiMember(Description="Id of the role to update, from get_roles.", IsRequired=true)
	Id string `json:"id"`
	/** @description The role's name — required; resend the current name to keep it. */
	// @ApiMember(Description="The role's name — required; resend the current name to keep it.", IsRequired=true)
	RoleName    string  `json:"roleName"`
	Description *string `json:"description,omitempty"`
	/** @description The complete new list of attached policy ids (full replacement), opaque ids from get_policies (e.g. 'pol_3kJ9xJ2mQ0aBcDeFgHiJk') — never invent or guess them. */
	// @ApiMember(Description="The complete new list of attached policy ids (full replacement), opaque ids from get_policies (e.g. 'pol_3kJ9xJ2mQ0aBcDeFgHiJk') — never invent or guess them.")
	Policies []string `json:"policies,omitempty"`
}

/** @description Create a new custom policy for project. */
// @Route("/{version}/membership/policies", "POST")
// @Api(Description="Create a new custom policy for project.")
type CreatePolicy struct {
	CodeMashRequestBase
	/** @description Display name of the new policy, unique within the project. */
	// @ApiMember(Description="Display name of the new policy, unique within the project.", IsRequired=true)
	PolicyName string `json:"policyName"`
	/** @description Optional human-readable description of what the policy grants. */
	// @ApiMember(Description="Optional human-readable description of what the policy grants.")
	Description *string `json:"description,omitempty"`
	/** @description AWS-IAM-style policy document as a raw JSON string (an object with a permission statement list) matching PolicyDocument.schema.json. Malformed or invalid documents are rejected before any change is made. */
	// @ApiMember(Description="AWS-IAM-style policy document as a raw JSON string (an object with a permission statement list) matching PolicyDocument.schema.json. Malformed or invalid documents are rejected before any change is made.")
	PolicyDocumentJson *string `json:"policyDocumentJson,omitempty"`
}

/** @description Deletes custom policy from project. */
// @Route("/{version}/membership/policies", "DELETE")
// @Api(Description="Deletes custom policy from project.")
type DeletePolicy struct {
	CodeMashRequestBase
	/** @description Public policy id, e.g. 'pol_database-read', from get_policies. */
	// @ApiMember(Description="Public policy id, e.g. 'pol_database-read', from get_policies.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets project policy details. */
// @Route("/{version}/membership/policies/{Id}", "GET")
// @Api(Description="Gets project policy details.")
type GetPolicy struct {
	CodeMashRequestBase
	/** @description Public policy id, e.g. 'pol_database-read', from get_policies. */
	// @ApiMember(Description="Public policy id, e.g. 'pol_database-read', from get_policies.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets project policies. */
// @Route("/{version}/membership/policies", "GET")
// @Api(Description="Gets project policies.")
type GetPolicies struct {
	CodeMashRequestBase
}

/** @description Updates a custom policy for project. */
// @Route("/{version}/membership/policies", "PUT")
// @Api(Description="Updates a custom policy for project.")
type UpdatePolicy struct {
	CodeMashRequestBase
	/** @description Public policy id, e.g. 'pol_database-read', from get_policies. */
	// @ApiMember(Description="Public policy id, e.g. 'pol_database-read', from get_policies.", IsRequired=true)
	Id string `json:"id"`
	/** @description The policy's name — required; resend the current name to keep it. */
	// @ApiMember(Description="The policy's name — required; resend the current name to keep it.", IsRequired=true)
	PolicyName string `json:"policyName"`
	/** @description Optional human-readable description of what the policy grants. */
	// @ApiMember(Description="Optional human-readable description of what the policy grants.")
	Description *string `json:"description,omitempty"`
	/** @description AWS-IAM-style policy document as a raw JSON string (an object with a permission statement list) matching PolicyDocument.schema.json — this is a FULL replacement of the policy's current permissions. */
	// @ApiMember(Description="AWS-IAM-style policy document as a raw JSON string (an object with a permission statement list) matching PolicyDocument.schema.json — this is a FULL replacement of the policy's current permissions.")
	PolicyDocumentJson *string `json:"policyDocumentJson,omitempty"`
}

/** @description Gets the project's passkey authentication settings. */
// @Route("/{version}/membership/passkey/settings", "GET")
// @Api(Description="Gets the project's passkey authentication settings.")
type GetPasskeySettings struct {
	CodeMashRequestBase
}

/** @description Saves the project's passkey authentication settings. */
// @Route("/{version}/membership/passkey/settings", "POST")
// @Api(Description="Saves the project's passkey authentication settings.")
type SavePasskeySettings struct {
	CodeMashRequestBase
	/** @description Whether the email + passkey sign-in flow is enabled for the project. */
	// @ApiMember(Description="Whether the email + passkey sign-in flow is enabled for the project.")
	Enabled bool `json:"enabled,omitempty"`
	/** @description Email verification code lifetime, in minutes. Allowed range: 3-15. */
	// @ApiMember(Description="Email verification code lifetime, in minutes. Allowed range: 3-15.")
	CodeTtlMinutes int `json:"codeTtlMinutes,omitempty"`
	/** @description Maximum passkeys a single user may register. Allowed range: 1-20. */
	// @ApiMember(Description="Maximum passkeys a single user may register. Allowed range: 1-20.")
	MaxCredentialsPerUser int `json:"maxCredentialsPerUser,omitempty"`
	/** @description Number of recovery codes generated at signup. Allowed range: 5-20. */
	// @ApiMember(Description="Number of recovery codes generated at signup. Allowed range: 5-20.")
	RecoveryCodeCount int `json:"recoveryCodeCount,omitempty"`
	/** @description Whether recovery codes are generated automatically at signup. */
	// @ApiMember(Description="Whether recovery codes are generated automatically at signup.")
	GenerateRecoveryCodesAtSignup bool `json:"generateRecoveryCodesAtSignup,omitempty"`
	/** @description Accepted authenticator types: 'Any', 'Platform', or 'CrossPlatform'. */
	// @ApiMember(Description="Accepted authenticator types: 'Any', 'Platform', or 'CrossPlatform'.")
	AuthenticatorAttachment string `json:"authenticatorAttachment"`
	/** @description Per-project opt-in for magic-link account recovery (off by default). */
	// @ApiMember(Description="Per-project opt-in for magic-link account recovery (off by default).")
	AllowMagicLinkRecovery bool `json:"allowMagicLinkRecovery,omitempty"`
	/** @description Absolute refresh-token lifetime, in days. Allowed range: 7-90. */
	// @ApiMember(Description="Absolute refresh-token lifetime, in days. Allowed range: 7-90.")
	RefreshTokenTtlDays int `json:"refreshTokenTtlDays,omitempty"`
	/** @description Optional explicit WebAuthn RP-ID — a bare DNS host (e.g. 'app.example.com', no scheme/port/path). Leave null/empty to derive it from the project's CORS origins. WARNING: changing this value invalidates every existing passkey on the project — set it once before going live and avoid changing it afterward. */
	// @ApiMember(Description="Optional explicit WebAuthn RP-ID — a bare DNS host (e.g. 'app.example.com', no scheme/port/path). Leave null/empty to derive it from the project's CORS origins. WARNING: changing this value invalidates every existing passkey on the project — set it once before going live and avoid changing it afterward.")
	RpId *string `json:"rpId,omitempty"`
}

// @Route("/{version}/membership/integrations/{Id}", "DELETE")
type DeleteMembershipIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_membership_integrations. */
	// @ApiMember(Description="Integration id, from get_membership_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/membership/integrations/{Id}/disable", "PUT")
type DisableMembershipIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_membership_integrations. */
	// @ApiMember(Description="Integration id, from get_membership_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/membership/integrations/{Id}/enable", "PUT")
type EnableMembershipIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_membership_integrations. */
	// @ApiMember(Description="Integration id, from get_membership_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/membership/integrations/{id}", "GET")
type GetMembershipIntegration struct {
	CodeMashRequestBase
	/** @description Integration id, from get_membership_integrations. */
	// @ApiMember(Description="Integration id, from get_membership_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/membership/integrations", "GET")
type GetMembershipIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/membership/integrations", "POST")
// @DataContract
type SaveMembershipIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration MembershipIntegrationRequest `json:"integration"`
}

// @Route("/{version}/membership/integrations/{Id}/default", "PUT")
type SetMembershipIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	/** @description Integration id, from get_membership_integrations. */
	// @ApiMember(Description="Integration id, from get_membership_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets the project's membership authorization (role-assignment) settings. */
// @Route("/{version}/membership/authorization", "GET")
// @Api(Description="Gets the project's membership authorization (role-assignment) settings.")
type GetAuthorizationSettings struct {
	CodeMashRequestBase
}

/** @description Updates the project's membership authorization settings. */
// @Route("/{version}/membership/authorization", "PUT")
// @Api(Description="Updates the project's membership authorization settings.")
type UpdateAuthorizationSettings struct {
	CodeMashRequestBase
	Setting                     *string  `json:"setting,omitempty"`
	DefaultRoles                []string `json:"defaultRoles,omitempty"`
	AllowedRegistrationRoles    []string `json:"allowedRegistrationRoles,omitempty"`
	AllowGuestUsers             bool     `json:"allowGuestUsers,omitempty"`
	GuestCleanupPeriodDays      *int     `json:"guestCleanupPeriodDays,omitempty"`
	UserRegistersAsRole         *string  `json:"userRegistersAsRole,omitempty"`
	GuestRegistersAsRole        *string  `json:"guestRegistersAsRole,omitempty"`
	AllowedRegisterRoles        []string `json:"allowedRegisterRoles,omitempty"`
	NeedVerification            bool     `json:"needVerification,omitempty"`
	VerificationEmailTemplate   *string  `json:"verificationEmailTemplate,omitempty"`
	DeactivationEmailTemplate   *string  `json:"deactivationEmailTemplate,omitempty"`
	AllowInviteUsers            bool     `json:"allowInviteUsers,omitempty"`
	AllowDeactivateUsers        bool     `json:"allowDeactivateUsers,omitempty"`
	InviteUserEmailTemplate     *string  `json:"inviteUserEmailTemplate,omitempty"`
	InvitationExpiration        *int     `json:"invitationExpiration,omitempty"`
	EmailVerificationExpiration *int     `json:"emailVerificationExpiration,omitempty"`
	DeactivationExpiration      *int     `json:"deactivationExpiration,omitempty"`
	DefaultSubscribeToNews      bool     `json:"defaultSubscribeToNews,omitempty"`
	MinLength                   int      `json:"minLength,omitempty"`
	MaxLength                   *int     `json:"maxLength,omitempty"`
	MinNumbers                  *int     `json:"minNumbers,omitempty"`
	MaxNumbers                  *int     `json:"maxNumbers,omitempty"`
	MinUpper                    *int     `json:"minUpper,omitempty"`
	MaxUpper                    *int     `json:"maxUpper,omitempty"`
	MinLower                    *int     `json:"minLower,omitempty"`
	MaxLower                    *int     `json:"maxLower,omitempty"`
	MinSpecial                  *int     `json:"minSpecial,omitempty"`
	MaxSpecial                  *int     `json:"maxSpecial,omitempty"`
	AllowedSpecial              *string  `json:"allowedSpecial,omitempty"`
}

/** @description Updates the project's membership password complexity policy. */
// @Route("/{version}/membership/authorization/password-complexity", "PUT")
// @Api(Description="Updates the project's membership password complexity policy.")
type UpdatePasswordComplexity struct {
	CodeMashRequestBase
	/** @description Minimum password length. */
	// @ApiMember(Description="Minimum password length.")
	MinLength int `json:"minLength,omitempty"`
	/** @description Maximum password length, if capped. */
	// @ApiMember(Description="Maximum password length, if capped.")
	MaxLength *int `json:"maxLength,omitempty"`
	/** @description Minimum number of numeric characters required. */
	// @ApiMember(Description="Minimum number of numeric characters required.")
	MinNumbers *int `json:"minNumbers,omitempty"`
	/** @description Maximum number of numeric characters allowed. */
	// @ApiMember(Description="Maximum number of numeric characters allowed.")
	MaxNumbers *int `json:"maxNumbers,omitempty"`
	/** @description Minimum number of uppercase characters required. */
	// @ApiMember(Description="Minimum number of uppercase characters required.")
	MinUpper *int `json:"minUpper,omitempty"`
	/** @description Maximum number of uppercase characters allowed. */
	// @ApiMember(Description="Maximum number of uppercase characters allowed.")
	MaxUpper *int `json:"maxUpper,omitempty"`
	/** @description Minimum number of lowercase characters required. */
	// @ApiMember(Description="Minimum number of lowercase characters required.")
	MinLower *int `json:"minLower,omitempty"`
	/** @description Maximum number of lowercase characters allowed. */
	// @ApiMember(Description="Maximum number of lowercase characters allowed.")
	MaxLower *int `json:"maxLower,omitempty"`
	/** @description Minimum number of special characters required. */
	// @ApiMember(Description="Minimum number of special characters required.")
	MinSpecial *int `json:"minSpecial,omitempty"`
	/** @description Maximum number of special characters allowed. */
	// @ApiMember(Description="Maximum number of special characters allowed.")
	MaxSpecial *int `json:"maxSpecial,omitempty"`
	/** @description The set of characters counted as 'special', if restricted. */
	// @ApiMember(Description="The set of characters counted as 'special', if restricted.")
	AllowedSpecial *string `json:"allowedSpecial,omitempty"`
}

/** @description Gets the project's configured membership authentication sign-in flows. */
// @Route("/{version}/membership/authentication", "GET")
// @Api(Description="Gets the project's configured membership authentication sign-in flows.")
type GetAuthenticationSettings struct {
	CodeMashRequestBase
}

/** @description Updates the project's membership authentication preferences. */
// @Route("/{version}/membership/authentication", "PUT")
// @Api(Description="Updates the project's membership authentication preferences.")
type UpdateAuthenticationSettings struct {
	CodeMashRequestBase
	/** @description Default URL to redirect end users to after logout. */
	// @ApiMember(Description="Default URL to redirect end users to after logout.")
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	/** @description Whether end users may sign in with a username in addition to email. */
	// @ApiMember(Description="Whether end users may sign in with a username in addition to email.")
	AllowUsernames bool `json:"allowUsernames,omitempty"`
	/** @description Per-authentication-mode logout URL overrides. */
	// @ApiMember(Description="Per-authentication-mode logout URL overrides.")
	Modes []CredentialsSettingsModeDto `json:"modes,omitempty"`
}

type MembershipIntegrationSaved struct {
	Integration MembershipIntegration `json:"integration"`
}

type MembershipIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type MembershipIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type MembershipIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MembershipIntegrationSetAsDefault struct {
	Id IntegrationId `json:"id"`
}

type MembershipIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MembershipIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MembershipEstablished struct {
}

type MembershipEnabled struct {
}

type MembershipDisabled struct {
}

type SetUserRegistersAsRole struct {
	ProjectId *ProjectId `json:"projectId,omitempty"`
	Role      RoleName   `json:"role"`
}

type PolicyCreated struct {
	Policy MembershipPolicy `json:"policy"`
}

type PolicyUpdated struct {
	Policy MembershipPolicy `json:"policy"`
}

type PolicyDeleted struct {
	PolicyId PolicyId `json:"policyId"`
}

type RoleCreated struct {
	Role MembershipRole `json:"role"`
}

type RoleUpdated struct {
	Role MembershipRole `json:"role"`
}

type RoleDeleted struct {
	RoleId RoleId `json:"roleId"`
}

type MembershipTriggerSaved struct {
	Trigger MembershipTrigger `json:"trigger"`
}

type MembershipTriggerMirrored struct {
	Trigger Trigger `json:"trigger"`
}

type MembershipTriggerEnabled struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type MembershipTriggerDisabled struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type MembershipTriggerDeleted struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

/** @description Disable database service */
// @Route("/{version}/database/disable", "GET")
// @Api(Description="Disable database service")
type DisableDatabase struct {
	CodeMashRequestBase
}

// @Route("/{version}/database/enable", "GET")
type EnableDatabase struct {
	CodeMashRequestBase
}

/** @description Delete database trigger */
// @Route("/{version}/database/schemas/triggers/{triggerId}", "DELETE")
// @Api(Description="Delete database trigger")
// @DataContract
type DeleteSchemaTrigger struct {
	DeleteTrigger
}

/** @description Disable database trigger */
// @Route("/{version}/database/schemas/triggers/{triggerId}/disable", "PATCH")
// @Api(Description="Disable database trigger")
// @DataContract
type DisableSchemaTrigger struct {
	DisableTrigger
}

/** @description Enable database trigger */
// @Route("/{version}/database/schemas/triggers/{triggerId}/enable", "PATCH")
// @Api(Description="Enable database trigger")
// @DataContract
type EnableSchemaTrigger struct {
	EnableTrigger
}

/** @description Gets database trigger by specified Id */
// @Route("/{version}/database/schemas/triggers/{id}", "GET")
// @Api(Description="Gets database trigger by specified Id")
type GetSchemaTrigger struct {
	GetTrigger
}

/** @description Gets database triggers */
// @Route("/{version}/database/schemas/triggers", "GET")
// @Api(Description="Gets database triggers")
type GetSchemaTriggers struct {
	GetTriggers
}

// @Route("/{version}/database/schemas/triggers", "POST")
// @DataContract
type SaveSchemaTrigger struct {
	SaveTrigger
}

/** @description Delete database taxonomy */
// @Route("/{version}/database/taxonomies/{Id}", "DELETE")
// @Api(Description="Delete database taxonomy")
type DeleteDatabaseTaxonomyRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy id to delete, from get_database_taxonomies. */
	// @ApiMember(Description="Taxonomy id to delete, from get_database_taxonomies.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets database taxonomy by id */
// @Route("/{version}/database/taxonomies/{id}", "GET")
// @Api(Description="Gets database taxonomy by id")
type GetDatabaseTaxonomy struct {
	CodeMashRequestBase
	/** @description Taxonomy id from get_database_taxonomies. */
	// @ApiMember(Description="Taxonomy id from get_database_taxonomies.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets database taxonomies */
// @Route("/{version}/database/taxonomies", "GET")
// @Api(Description="Gets database taxonomies")
type GetDatabaseTaxonomies struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Returns the single-parent taxonomy structure tree */
// @Route("/{version}/database/taxonomies/tree", "GET")
// @Api(Description="Returns the single-parent taxonomy structure tree")
// @DataContract
type GetDatabaseTaxonomyTreeRequest struct {
	CodeMashRequestBase
	/** @description When true, each taxonomy node also carries its own term tree (heavier response). */
	// @DataMember
	// @ApiMember(Description="When true, each taxonomy node also carries its own term tree (heavier response).")
	IncludeTerms bool `json:"includeTerms,omitempty"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Creates or updates a database taxonomy */
// @Route("/{version}/database/taxonomies", "POST")
// @Api(Description="Creates or updates a database taxonomy")
// @DataContract
type SaveDatabaseTaxonomyRequest struct {
	CodeMashRequestBase
	/** @description Empty to create a new taxonomy; set to an existing taxonomy id (from get_database_taxonomies) to update it. */
	// @DataMember
	// @ApiMember(Description="Empty to create a new taxonomy; set to an existing taxonomy id (from get_database_taxonomies) to update it.")
	ViewId *string `json:"viewId,omitempty"`
	/** @description Human-entered taxonomy title (e.g. "Countries"); a slug is derived server-side. */
	// @DataMember
	// @ApiMember(Description="Human-entered taxonomy title (e.g. \"Countries\"); a slug is derived server-side.", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	/** @description Optional free-text description of the taxonomy. */
	// @DataMember
	// @ApiMember(Description="Optional free-text description of the taxonomy.")
	Description *string `json:"description,omitempty"`
	/** @description Optional raw JSON string (Norbix data meta-schema) describing custom meta fields for terms in this taxonomy. Omit to leave the taxonomy structural-only. */
	// @DataMember
	// @ApiMember(Description="Optional raw JSON string (Norbix data meta-schema) describing custom meta fields for terms in this taxonomy. Omit to leave the taxonomy structural-only.")
	TermsMetaDataSchema *string `json:"termsMetaDataSchema,omitempty"`
	/** @description Optional raw JSON string (Norbix UI/visual meta-schema) describing the term meta edit form. */
	// @DataMember
	// @ApiMember(Description="Optional raw JSON string (Norbix UI/visual meta-schema) describing the term meta edit form.")
	TermsMetaVisualSchema *string `json:"termsMetaVisualSchema,omitempty"`
	/** @description Optional parent taxonomy id. The child taxonomy points to its parent — e.g. set the Countries taxonomy's parentId to the Regions taxonomy id so each country term can be parented by a region term. Omit for a root taxonomy. */
	// @DataMember
	// @ApiMember(Description="Optional parent taxonomy id. The child taxonomy points to its parent — e.g. set the Countries taxonomy's parentId to the Regions taxonomy id so each country term can be parented by a region term. Omit for a root taxonomy.")
	ParentId *string `json:"parentId,omitempty"`
	/** @description Optional list of other taxonomy ids this taxonomy depends on for multi-parent terms. Omit for a self-contained taxonomy. */
	// @DataMember
	// @ApiMember(Description="Optional list of other taxonomy ids this taxonomy depends on for multi-parent terms. Omit for a self-contained taxonomy.")
	Dependencies []string `json:"dependencies,omitempty"`
}

/** @description Delete a single term from a taxonomy by id */
// @Route("/{version}/database/taxonomies/{TaxonomyId}/terms/{Id}", "DELETE")
// @Api(Description="Delete a single term from a taxonomy by id")
// @DataContract
type DeleteDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy id that owns the term, from get_database_taxonomies. */
	// @DataMember
	// @ApiMember(Description="Taxonomy id that owns the term, from get_database_taxonomies.", IsRequired=true)
	TaxonomyId string `json:"taxonomyId"`
	/** @description Term id to delete, from get_database_taxonomy_term_tree. */
	// @DataMember
	// @ApiMember(Description="Term id to delete, from get_database_taxonomy_term_tree.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Delete many terms in a taxonomy matching the given filter */
// @Route("/{version}/database/taxonomies/{TaxonomyId}/terms/many", "DELETE")
// @Api(Description="Delete many terms in a taxonomy matching the given filter")
// @DataContract
type DeleteManyDatabaseTaxonomyTermsRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy id whose terms to delete, from get_database_taxonomies. */
	// @DataMember
	// @ApiMember(Description="Taxonomy id whose terms to delete, from get_database_taxonomies.", IsRequired=true)
	TaxonomyId string `json:"taxonomyId"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description MongoDB extended-JSON match filter (a raw JSON object, e.g. {"active":false}) selecting which terms to delete. Automatically ANDed server-side with the taxonomyId, so it cannot affect other taxonomies. */
	// @DataMember
	// @ApiMember(Description="MongoDB extended-JSON match filter (a raw JSON object, e.g. {\"active\":false}) selecting which terms to delete. Automatically ANDed server-side with the taxonomyId, so it cannot affect other taxonomies.", IsRequired=true)
	Filter string `json:"filter"`
}

/** @description Get a single term from a taxonomy by id */
// @Route("/{version}/database/taxonomies/{TaxonomyId}/terms/{Id}", "GET")
// @Api(Description="Get a single term from a taxonomy by id")
// @DataContract
type GetDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy id from get_database_taxonomies. */
	// @DataMember
	// @ApiMember(Description="Taxonomy id from get_database_taxonomies.", IsRequired=true)
	TaxonomyId string `json:"taxonomyId"`
	/** @description Term id from get_database_taxonomy_term_tree. */
	// @DataMember
	// @ApiMember(Description="Term id from get_database_taxonomy_term_tree.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Returns a merged term tree across a taxonomy and its child taxonomies */
// @Route("/{version}/database/taxonomies/{TaxonomyName}/merged-tree", "GET")
// @Api(Description="Returns a merged term tree across a taxonomy and its child taxonomies")
// @DataContract
type GetDatabaseMergedTermTreeRequest struct {
	CodeMashRequestBase
	/** @description Root taxonomy slug/name (from get_database_taxonomies). Its terms are the roots; child-taxonomy terms nest under them. */
	// @DataMember
	// @ApiMember(Description="Root taxonomy slug/name (from get_database_taxonomies). Its terms are the roots; child-taxonomy terms nest under them.", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Returns the whole term tree of a taxonomy (or a sub-tree) in one call */
// @Route("/{version}/database/taxonomies/{TaxonomyName}/terms/tree", "GET")
// @Api(Description="Returns the whole term tree of a taxonomy (or a sub-tree) in one call")
// @DataContract
type GetDatabaseTaxonomyTermTreeRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy slug/name to fetch the term tree for, from get_database_taxonomies. */
	// @DataMember
	// @ApiMember(Description="Taxonomy slug/name to fetch the term tree for, from get_database_taxonomies.", IsRequired=true)
	TaxonomyName string `json:"taxonomyName"`
	/** @description Optional term id to root the returned tree at a sub-tree instead of the whole taxonomy. */
	// @DataMember
	// @ApiMember(Description="Optional term id to root the returned tree at a sub-tree instead of the whole taxonomy.")
	RootTermId *string `json:"rootTermId,omitempty"`
	/** @description Optional maximum depth to return below the root. */
	// @DataMember
	// @ApiMember(DataType="integer", Description="Optional maximum depth to return below the root.")
	Depth *int `json:"depth,omitempty"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Insert a single term into a taxonomy */
// @Route("/{version}/database/taxonomies/{TaxonomyId}/terms", "POST")
// @Api(Description="Insert a single term into a taxonomy")
// @DataContract
type SaveDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy id to insert the term into, from get_database_taxonomies. */
	// @DataMember
	// @ApiMember(Description="Taxonomy id to insert the term into, from get_database_taxonomies.", IsRequired=true)
	TaxonomyId string `json:"taxonomyId"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The term to insert, as a MongoDB extended-JSON document string (a raw JSON object). Supported term fields: name (string, or a {lang:value} map — required); description; order (integer sort position, lower shows first — omit for unordered); parentId (id of the single parent term); multiParents ("additional categories": array of {taxonomyId, parentId}). The server stamps taxonomyId/taxonomyName automatically — do not include them. Example: {"name":"France","order":1}. */
	// @DataMember
	// @ApiMember(Description="The term to insert, as a MongoDB extended-JSON document string (a raw JSON object). Supported term fields: name (string, or a {lang:value} map — required); description; order (integer sort position, lower shows first — omit for unordered); parentId (id of the single parent term); multiParents (\"additional categories\": array of {taxonomyId, parentId}). The server stamps taxonomyId/taxonomyName automatically — do not include them. Example: {\"name\":\"France\",\"order\":1}.", IsRequired=true)
	Document string `json:"document"`
}

/** @description Update a single term in a taxonomy by id */
// @Route("/{version}/database/taxonomies/{TaxonomyId}/terms/{Id}", "PUT")
// @Api(Description="Update a single term in a taxonomy by id")
// @DataContract
type UpdateDatabaseTaxonomyTermRequest struct {
	CodeMashRequestBase
	/** @description Taxonomy id that owns the term, from get_database_taxonomies. */
	// @DataMember
	// @ApiMember(Description="Taxonomy id that owns the term, from get_database_taxonomies.", IsRequired=true)
	TaxonomyId string `json:"taxonomyId"`
	/** @description Term id to update, from get_database_taxonomy_term_tree. */
	// @DataMember
	// @ApiMember(Description="Term id to update, from get_database_taxonomy_term_tree.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Partial update document as MongoDB extended-JSON (a raw JSON object of fields to change), applied with $set — only the given fields change. Updatable term fields: name (string or {lang:value} map); description; order (integer sort position, lower shows first — use this to numerate/rank terms; set null to clear); parentId (single parent term id — a term from THIS taxonomy's parent taxonomy; e.g. link a country to its region by setting the country term's parentId to the region term id); multiParents ("additional categories": array of {taxonomyId, parentId}). Example to rank a term: {"order":1}. */
	// @DataMember
	// @ApiMember(Description="Partial update document as MongoDB extended-JSON (a raw JSON object of fields to change), applied with $set — only the given fields change. Updatable term fields: name (string or {lang:value} map); description; order (integer sort position, lower shows first — use this to numerate/rank terms; set null to clear); parentId (single parent term id — a term from THIS taxonomy's parent taxonomy; e.g. link a country to its region by setting the country term's parentId to the region term id); multiParents (\"additional categories\": array of {taxonomyId, parentId}). Example to rank a term: {\"order\":1}.", IsRequired=true)
	Update string `json:"update"`
}

/** @description Creates every collection and taxonomy of a compiled IF bundle, linked and published */
// @Route("/{version}/database/schemas/apply-bundle", "POST")
// @Api(Description="Creates every collection and taxonomy of a compiled IF bundle, linked and published")
// @DataContract
type ApplyDatabaseSchemaBundleRequest struct {
	CodeMashRequestBase
	/** @description Comma-separated catalog entity ids to create from the reviewed catalog (e.g. "blog_posts,comments"). The usual input. */
	// @DataMember
	// @ApiMember(Description="Comma-separated catalog entity ids to create from the reviewed catalog (e.g. \"blog_posts,comments\"). The usual input.")
	Entities *string `json:"entities,omitempty"`
	/** @description Only for entities the catalog lacks: one IF entity object or an array of IF objects (JSON string). May also hold catalog refs with add_fields / remove_fields. */
	// @DataMember
	// @ApiMember(Description="Only for entities the catalog lacks: one IF entity object or an array of IF objects (JSON string). May also hold catalog refs with add_fields / remove_fields.")
	BundleJson *string `json:"bundleJson,omitempty"`
	/** @description Field tier to compile: minimal | standard (default) | extended. */
	// @DataMember
	// @ApiMember(Description="Field tier to compile: minimal | standard (default) | extended.")
	Tier *string `json:"tier,omitempty"`
	/** @description true = mark free-text fields (title, body, excerpt …) translatable for multilingual content. Default false. */
	// @DataMember
	// @ApiMember(Description="true = mark free-text fields (title, body, excerpt …) translatable for multilingual content. Default false.")
	Translatable bool `json:"translatable,omitempty"`
	/** @description false = leave every collection as a draft instead of publishing v1. Default true. */
	// @DataMember
	// @ApiMember(Description="false = leave every collection as a draft instead of publishing v1. Default true.")
	Publish bool `json:"publish,omitempty"`
}

/** @description Delete database schema (collection) */
// @Route("/{version}/database/schemas/{Id}", "DELETE")
// @Api(Description="Delete database schema (collection)")
type DeleteDatabaseSchemaRequest struct {
	CodeMashRequestBase
	/** @description Schema id to delete, from get_database_schemas. */
	// @ApiMember(Description="Schema id to delete, from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Discards the working-copy draft of a database schema without publishing */
// @Route("/{version}/database/schemas/{Id}/draft", "DELETE")
// @Api(Description="Discards the working-copy draft of a database schema without publishing")
// @DataContract
type DiscardDatabaseSchemaDraftRequest struct {
	CodeMashRequestBase
	/** @description Schema id whose draft to discard, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id whose draft to discard, from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets database schema by id */
// @Route("/{version}/database/schemas/{id}", "GET")
// @Api(Description="Gets database schema by id")
// @DataContract
type GetDatabaseSchema struct {
	CodeMashRequestBase
	/** @description Schema id from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional published version number to pin; omit for the latest published version. */
	// @DataMember(Name="version")
	// @ApiMember(DataType="integer", Description="Optional published version number to pin; omit for the latest published version.", Name="version", ParameterType="query")
	Version *int `json:"version,omitempty"`
}

/** @description Gets database schemas (collections) */
// @Route("/{version}/database/schemas", "GET")
// @Api(Description="Gets database schemas (collections)")
type GetDatabaseSchemas struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Gets the current draft of a database schema */
// @Route("/{version}/database/schemas/{Id}/draft", "GET")
// @Api(Description="Gets the current draft of a database schema")
// @DataContract
type GetDatabaseSchemaDraft struct {
	CodeMashRequestBase
	/** @description Schema id from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets database schema records-list display settings */
// @Route("/{version}/database/schemas/{Id}/list-settings", "GET")
// @Api(Description="Gets database schema records-list display settings")
// @DataContract
type GetDatabaseSchemaListSettings struct {
	CodeMashRequestBase
	/** @description Schema id from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Structural diff between two published versions of a database schema */
// @Route("/{version}/database/schemas/{Id}/versions/diff", "GET")
// @Api(Description="Structural diff between two published versions of a database schema")
// @DataContract
type GetDatabaseSchemaVersionDiff struct {
	CodeMashRequestBase
	/** @description Schema id from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
	/** @description Earlier published version number to diff from. Get valid values from get_database_schema_versions. */
	// @DataMember
	// @ApiMember(DataType="integer", Description="Earlier published version number to diff from. Get valid values from get_database_schema_versions.", IsRequired=true)
	FromVersion int `json:"fromVersion"`
	/** @description Later published version number to diff to. Get valid values from get_database_schema_versions. */
	// @DataMember
	// @ApiMember(DataType="integer", Description="Later published version number to diff to. Get valid values from get_database_schema_versions.", IsRequired=true)
	ToVersion int `json:"toVersion"`
}

/** @description Lists published version summaries for a database schema */
// @Route("/{version}/database/schemas/{Id}/versions", "GET")
// @Api(Description="Lists published version summaries for a database schema")
// @DataContract
type GetDatabaseSchemaVersions struct {
	CodeMashRequestBase
	/** @description Schema id from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/database/schemas/{Id}/publish", "POST")
// @DataContract
type PublishDatabaseSchemaRequest struct {
	CodeMashRequestBase
	// @DataMember
	Id string `json:"id"`
	// @DataMember
	Confirmed bool `json:"confirmed,omitempty"`
}

/** @description Renames a database schema (collection) */
// @Route("/{version}/database/schemas/{Id}/rename", "PUT")
// @Api(Description="Renames a database schema (collection)")
// @DataContract
type RenameDatabaseSchemaRequest struct {
	CodeMashRequestBase
	/** @description Schema id to rename, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id to rename, from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
	/** @description New human-entered title (e.g. "Company Employees"); the slug is derived server-side. */
	// @DataMember
	// @ApiMember(Description="New human-entered title (e.g. \"Company Employees\"); the slug is derived server-side.", IsRequired=true)
	Title string `json:"title"`
	/** @description When true (default), rejects the rename if another schema already owns the derived slug. Leave true unless explicitly asked to bypass the uniqueness check. */
	// @DataMember
	// @ApiMember(Description="When true (default), rejects the rename if another schema already owns the derived slug. Leave true unless explicitly asked to bypass the uniqueness check.")
	RenameUniqueName bool `json:"renameUniqueName,omitempty"`
}

// @Route("/{version}/database/schemas", "POST")
// @DataContract
type SaveDatabaseSchemaRequest struct {
	CodeMashRequestBase
	/** @description Empty to create a new schema; set to an existing schema id (from get_database_schemas) to update its draft. */
	// @DataMember
	// @ApiMember(Description="Empty to create a new schema; set to an existing schema id (from get_database_schemas) to update its draft.")
	ViewId *string `json:"viewId,omitempty"`
	/** @description Human-entered schema title (e.g. "Company Employees"); a slug is derived server-side. */
	// @DataMember
	// @ApiMember(Description="Human-entered schema title (e.g. \"Company Employees\"); a slug is derived server-side.", IsRequired=true)
	SchemaName string `json:"schemaName"`
	/** @description Raw JSON string matching the Norbix data meta-schema (https://norbix.ai/schemas/meta/v1.json). When unsure of the shape, read an existing schema with get_database_schema first. */
	// @DataMember
	// @ApiMember(Description="Raw JSON string matching the Norbix data meta-schema (https://norbix.ai/schemas/meta/v1.json). When unsure of the shape, read an existing schema with get_database_schema first.")
	DataSchema *string `json:"dataSchema,omitempty"`
	/** @description OPTIONAL raw JSON string matching the Norbix UI/visual meta-schema (https://norbix.ai/schemas/ui/v1.json), describing the record form layout. If omitted or invalid, the backend auto-generates a flat-list form from the data schema; provide it to control the layout. */
	// @DataMember
	// @ApiMember(Description="OPTIONAL raw JSON string matching the Norbix UI/visual meta-schema (https://norbix.ai/schemas/ui/v1.json), describing the record form layout. If omitted or invalid, the backend auto-generates a flat-list form from the data schema; provide it to control the layout.")
	VisualSchema *string `json:"visualSchema,omitempty"`
	/** @description Optional schema-level settings (e.g. record validation behavior). */
	// @DataMember
	// @ApiMember(Description="Optional schema-level settings (e.g. record validation behavior).")
	Settings *SchemaSettingsDto `json:"settings,omitempty"`
}

/** @description Saves the working-copy draft of a database schema */
// @Route("/{version}/database/schemas/{Id}/draft", "PUT")
// @Api(Description="Saves the working-copy draft of a database schema")
// @DataContract
type UpdateDatabaseSchemaDraftRequest struct {
	CodeMashRequestBase
	/** @description Schema id whose draft to replace, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id whose draft to replace, from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
	/** @description Raw JSON string matching the Norbix data meta-schema (https://norbix.ai/schemas/meta/v1.json) for the draft's data schema. */
	// @DataMember
	// @ApiMember(Description="Raw JSON string matching the Norbix data meta-schema (https://norbix.ai/schemas/meta/v1.json) for the draft's data schema.")
	DataSchema *string `json:"dataSchema,omitempty"`
	/** @description Raw JSON string matching the Norbix UI/visual meta-schema (https://norbix.ai/schemas/ui/v1.json) for the draft's record form. */
	// @DataMember
	// @ApiMember(Description="Raw JSON string matching the Norbix UI/visual meta-schema (https://norbix.ai/schemas/ui/v1.json) for the draft's record form.")
	VisualSchema *string `json:"visualSchema,omitempty"`
}

/** @description Updates database schema records-list display settings */
// @Route("/{version}/database/schemas/{Id}/list-settings", "PUT")
// @Api(Description="Updates database schema records-list display settings")
// @DataContract
type UpdateDatabaseSchemaListSettingsRequest struct {
	CodeMashRequestBase
	/** @description Schema id whose list settings to update, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id whose list settings to update, from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
	/** @description The complete new list settings object (full replace). */
	// @DataMember
	// @ApiMember(Description="The complete new list settings object (full replace).", IsRequired=true)
	Settings SchemaListSettingsDto `json:"settings"`
}

/** @description Updates database schema settings */
// @Route("/{version}/database/schemas/{Id}/settings", "PUT")
// @Api(Description="Updates database schema settings")
// @DataContract
type UpdateDatabaseSchemaSettingsRequest struct {
	CodeMashRequestBase
	/** @description Schema id whose settings to update, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id whose settings to update, from get_database_schemas.", IsRequired=true)
	Id string `json:"id"`
	/** @description The new schema settings object. */
	// @DataMember
	// @ApiMember(Description="The new schema settings object.", IsRequired=true)
	Settings SchemaSettingsDto `json:"settings"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/aggregate", "POST")
// @Api(Description="Database")
type AggregateRecords struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to run the aggregation against. */
	// @ApiMember(Description="The collection (schema) name to run the aggregation against.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The aggregation pipeline as a MongoDB extended-JSON array of stages. */
	// @ApiMember(Description="The aggregation pipeline as a MongoDB extended-JSON array of stages.", IsRequired=true)
	Pipeline string `json:"pipeline"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}/responsibility", "PUT")
// @Api(Description="Database")
type ChangeRecordResponsibility struct {
	CodeMashRequestBase
	/** @description The collection (schema) name the record lives in. */
	// @ApiMember(Description="The collection (schema) name the record lives in.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description The id of the record whose responsibility changes. */
	// @ApiMember(Description="The id of the record whose responsibility changes.", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The new responsible user (owner) id. */
	// @ApiMember(Description="The new responsible user (owner) id.", IsRequired=true)
	NewResponsibleUserId string `json:"newResponsibleUserId"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/count", "GET")
// @Api(Description="Database")
type CountRecords struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to count records in. */
	// @ApiMember(Description="The collection (schema) name to count records in.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional MongoDB extended-JSON filter. Empty means count all records. */
	// @ApiMember(Description="Optional MongoDB extended-JSON filter. Empty means count all records.")
	Filter        *string `json:"filter,omitempty"`
	SchemaVersion *int    `json:"schemaVersion,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "DELETE")
// @Api(Description="Database")
type DeleteManyRecords struct {
	CodeMashRequestBase
	/** @description The collection (schema) name the records live in. */
	// @ApiMember(Description="The collection (schema) name the records live in.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The match filter as a MongoDB extended-JSON document. Required. */
	// @ApiMember(Description="The match filter as a MongoDB extended-JSON document. Required.", IsRequired=true)
	Filter string `json:"filter"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "DELETE")
// @Api(Description="Database")
type DeleteRecord struct {
	CodeMashRequestBase
	/** @description The collection (schema) name the record lives in. */
	// @ApiMember(Description="The collection (schema) name the record lives in.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description The id of the record to delete. */
	// @ApiMember(Description="The id of the record to delete.", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/distinct", "GET")
// @Api(Description="Database")
type DistinctRecordValues struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to read from. */
	// @ApiMember(Description="The collection (schema) name to read from.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The document field (dotted path allowed) to get distinct values for, e.g. 'status'. */
	// @ApiMember(Description="The document field (dotted path allowed) to get distinct values for, e.g. 'status'.", IsRequired=true)
	Field string `json:"field"`
	/** @description Optional MongoDB extended-JSON filter. Empty means consider all records. */
	// @ApiMember(Description="Optional MongoDB extended-JSON filter. Empty means consider all records.")
	Filter        *string `json:"filter,omitempty"`
	SchemaVersion *int    `json:"schemaVersion,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/aggregates/{aggregateId}/execute", "POST")
// @Api(Description="Database")
type ExecuteRecordsAggregate struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to run the saved aggregation against. */
	// @ApiMember(Description="The collection (schema) name to run the saved aggregation against.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description The saved aggregate id (maggr_…) to execute. */
	// @ApiMember(Description="The saved aggregate id (maggr_…) to execute.", IsRequired=true)
	AggregateId           string  `json:"aggregateId"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional key/value substitutions for {TokenKey} placeholders in the saved pipeline. */
	// @ApiMember(Description="Optional key/value substitutions for {TokenKey} placeholders in the saved pipeline.")
	Tokens map[string]string `json:"tokens,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}", "GET")
// @Api(Description="Database")
type FindRecords struct {
	CodeMashListPaginationRequestBase
	/** @description The collection (schema) name to read from. */
	// @ApiMember(Description="The collection (schema) name to read from.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Optional MongoDB extended-JSON filter. Empty means match all records. */
	// @ApiMember(Description="Optional MongoDB extended-JSON filter. Empty means match all records.")
	Filter *string `json:"filter,omitempty"`
	/** @description Optional contact id (ct_…) — only that contact's records are returned. */
	// @ApiMember(Description="Optional contact id (ct_…) — only that contact's records are returned.")
	ContactId     *string     `json:"contactId,omitempty"`
	SchemaVersion *int        `json:"schemaVersion,omitempty"`
	PagingArgs    *PagingArgs `json:"pagingArgs,omitempty"`
	SortBy        *string     `json:"sortBy,omitempty"`
	SortOrder     *int        `json:"sortOrder,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "GET")
// @Api(Description="Database")
type FindOneRecord struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to read from. */
	// @ApiMember(Description="The collection (schema) name to read from.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description The id of the record to fetch. */
	// @ApiMember(Description="The id of the record to fetch.", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/indexes", "GET")
// @Api(Description="Database")
// @DataContract
type GetCollectionIndexes struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to inspect. */
	// @DataMember
	// @ApiMember(Description="The collection (schema) name to inspect.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	// @DataMember
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "POST")
// @Api(Description="Database")
type InsertManyRecords struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to insert into. */
	// @ApiMember(Description="The collection (schema) name to insert into.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The records to insert as a MongoDB extended-JSON array of documents. */
	// @ApiMember(Description="The records to insert as a MongoDB extended-JSON array of documents.", IsRequired=true)
	Documents string `json:"documents"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}", "POST")
// @Api(Description="Database")
type InsertRecord struct {
	CodeMashRequestBase
	/** @description The collection (schema) name to insert into. */
	// @ApiMember(Description="The collection (schema) name to insert into.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The record to insert, as a MongoDB extended-JSON document string. */
	// @ApiMember(Description="The record to insert, as a MongoDB extended-JSON document string.", IsRequired=true)
	Document string `json:"document"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}/replace", "PUT")
// @Api(Description="Database")
type ReplaceRecord struct {
	CodeMashRequestBase
	/** @description The collection (schema) name the record lives in. */
	// @ApiMember(Description="The collection (schema) name the record lives in.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description The id of the record to replace. */
	// @ApiMember(Description="The id of the record to replace.", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The replacement document as MongoDB extended-JSON. */
	// @ApiMember(Description="The replacement document as MongoDB extended-JSON.", IsRequired=true)
	Replacement string `json:"replacement"`
}

/** @description Database */
// @Route("/{version}/database/collections/seed", "POST")
// @Api(Description="Database")
type SeedCollectionRecords struct {
	CodeMashRequestBase
	/** @description Seeding mode: 'dummy' (server-generated sample data, default) or 'realistic' (caller-supplied documents). */
	// @ApiMember(Description="Seeding mode: 'dummy' (server-generated sample data, default) or 'realistic' (caller-supplied documents).")
	Mode                  *string `json:"mode,omitempty"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description JSON array of {collectionName, count?, documents?}. count applies to dummy mode (max 100 per collection); documents (extended-JSON objects, may contain $seedRef placeholders) apply to realistic mode. */
	// @ApiMember(Description="JSON array of {collectionName, count?, documents?}. count applies to dummy mode (max 100 per collection); documents (extended-JSON objects, may contain $seedRef placeholders) apply to realistic mode.", IsRequired=true)
	Collections string `json:"collections"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/many", "PUT")
// @Api(Description="Database")
type UpdateManyRecords struct {
	CodeMashRequestBase
	/** @description The collection (schema) name the records live in. */
	// @ApiMember(Description="The collection (schema) name the records live in.", IsRequired=true)
	CollectionName        string  `json:"collectionName"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The match filter as a MongoDB extended-JSON document. Empty object means match all. */
	// @ApiMember(Description="The match filter as a MongoDB extended-JSON document. Empty object means match all.", IsRequired=true)
	Filter string `json:"filter"`
	/** @description The partial update document (applied with $set), as MongoDB extended-JSON. */
	// @ApiMember(Description="The partial update document (applied with $set), as MongoDB extended-JSON.", IsRequired=true)
	Update string `json:"update"`
}

/** @description Database */
// @Route("/{version}/database/collections/{collectionName}/{id}", "PUT")
// @Api(Description="Database")
type UpdateOneRecord struct {
	CodeMashRequestBase
	/** @description The collection (schema) name the record lives in. */
	// @ApiMember(Description="The collection (schema) name the record lives in.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description The id of the record to update. */
	// @ApiMember(Description="The id of the record to update.", IsRequired=true)
	Id                    string  `json:"id"`
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description The partial update document (applied with $set), as MongoDB extended-JSON. */
	// @ApiMember(Description="The partial update document (applied with $set), as MongoDB extended-JSON.", IsRequired=true)
	Update string `json:"update"`
}

/** @description Delete integration for particular project */
// @Route("/{version}/database/integrations/{Id}", "DELETE")
// @Api(Description="Delete integration for particular project")
type DeleteDatabaseIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Database integration id to delete, from get_database_integrations. */
	// @ApiMember(Description="Database integration id to delete, from get_database_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Disable integration for particular project */
// @Route("/{version}/database/integrations/{Id}/disable", "PUT")
// @Api(Description="Disable integration for particular project")
type DisableDatabaseIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Database integration id to disable, from get_database_integrations. */
	// @ApiMember(Description="Database integration id to disable, from get_database_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Enable integration for particular project */
// @Route("/{version}/database/integrations/{Id}/enable", "PUT")
// @Api(Description="Enable integration for particular project")
type EnableDatabaseIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Database integration id to enable, from get_database_integrations. */
	// @ApiMember(Description="Database integration id to enable, from get_database_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integration by specified Id */
// @Route("/{version}/database/integrations/{id}", "GET")
// @Api(Description="Gets integration by specified Id")
type GetDatabaseIntegration struct {
	CodeMashRequestBase
	/** @description Database integration id from get_database_integrations. */
	// @ApiMember(Description="Database integration id from get_database_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets database integrations */
// @Route("/{version}/database/integrations", "GET")
// @Api(Description="Gets database integrations")
type GetDatabaseIntegrations struct {
	CodeMashListPaginationRequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Returns the Flex tiers this account is entitled to pick */
// @Route("/{version}/database/integrations/flex-tiers", "GET")
// @Api(Description="Returns the Flex tiers this account is entitled to pick")
type GetAllowedFlexTiers struct {
	CodeMashRequestBase
}

// @Route("/{version}/database/integrations/{Id}/connection-string", "GET")
type RevealManagedFlexConnectionString struct {
	CodeMashRequestBase
	Id string `json:"id"`
}

/** @description Saves database integration */
// @Route("/{version}/database/integrations", "POST")
// @Api(Description="Saves database integration")
// @DataContract
type SaveDatabaseIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration DatabaseIntegrationRequest `json:"integration"`
}

/** @description Sets integration as default */
// @Route("/{version}/database/integrations/{Id}/default", "PUT")
// @Api(Description="Sets integration as default")
type SetDatabaseIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	/** @description Database integration id to set as default, from get_database_integrations. */
	// @ApiMember(Description="Database integration id to set as default, from get_database_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Test database integration */
// @Route("/{version}/database/integrations/test", "POST")
// @Api(Description="Test database integration")
type TestDatabaseIntegration struct {
	CodeMashRequestBase
	/** @description Database integration id to test, from get_database_integrations. */
	// @ApiMember(Description="Database integration id to test, from get_database_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

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

/** @description Database */
// @Route("/{version}/database/imports", "GET")
// @Api(Description="Database")
type GetCollectionImports struct {
	CodeMashListPaginationRequestBase
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

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

/** @description Delete saved Mongo aggregation */
// @Route("/{version}/database/aggregates/{Id}", "DELETE")
// @Api(Description="Delete saved Mongo aggregation")
// @DataContract
type DeleteDatabaseAggregateRequest struct {
	CodeMashRequestBase
	/** @description Aggregate id to delete, from get_database_aggregates. */
	// @DataMember
	// @ApiMember(Description="Aggregate id to delete, from get_database_aggregates.", IsRequired=true)
	Id string `json:"id"`
	/** @description Schema id that owns this aggregate, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id that owns this aggregate, from get_database_schemas.", IsRequired=true)
	SchemaId string `json:"schemaId"`
}

/** @description Get saved Mongo aggregation by id */
// @Route("/{version}/database/aggregates/{Id}", "GET")
// @Api(Description="Get saved Mongo aggregation by id")
type GetDatabaseAggregate struct {
	CodeMashRequestBase
	/** @description Aggregate id from get_database_aggregates. */
	// @ApiMember(Description="Aggregate id from get_database_aggregates.", IsRequired=true)
	Id string `json:"id"`
	/** @description Schema id that owns this aggregate, from get_database_schemas. */
	// @ApiMember(Description="Schema id that owns this aggregate, from get_database_schemas.", IsRequired=true)
	SchemaId string `json:"schemaId"`
}

/** @description Lists saved Mongo aggregations for a schema */
// @Route("/{version}/database/aggregates", "GET")
// @Api(Description="Lists saved Mongo aggregations for a schema")
type GetDatabaseAggregates struct {
	CodeMashListPaginationRequestBase
	/** @description Schema id whose saved aggregates to list, from get_database_schemas. */
	// @ApiMember(Description="Schema id whose saved aggregates to list, from get_database_schemas.", IsRequired=true)
	SchemaId   string      `json:"schemaId"`
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

/** @description Creates or updates a saved Mongo aggregation */
// @Route("/{version}/database/aggregates", "POST")
// @Api(Description="Creates or updates a saved Mongo aggregation")
// @DataContract
type SaveDatabaseAggregateRequest struct {
	CodeMashRequestBase
	/** @description Empty to create a new saved aggregate; set to an existing aggregate id (from get_database_aggregates) to update it. */
	// @DataMember
	// @ApiMember(Description="Empty to create a new saved aggregate; set to an existing aggregate id (from get_database_aggregates) to update it.")
	ViewId *string `json:"viewId,omitempty"`
	/** @description Schema id that owns this aggregate, from get_database_schemas. */
	// @DataMember
	// @ApiMember(Description="Schema id that owns this aggregate, from get_database_schemas.", IsRequired=true)
	SchemaId string `json:"schemaId"`
	/** @description Human-readable display name for the saved aggregate. */
	// @DataMember
	// @ApiMember(Description="Human-readable display name for the saved aggregate.", IsRequired=true)
	DisplayName string `json:"displayName"`
	/** @description Optional free-text description of what the aggregate does. */
	// @DataMember
	// @ApiMember(Description="Optional free-text description of what the aggregate does.")
	Description *string `json:"description,omitempty"`
	/** @description MongoDB aggregation pipeline JSON, optionally containing {TokenKey} placeholders substituted at execute time. */
	// @DataMember
	// @ApiMember(Description="MongoDB aggregation pipeline JSON, optionally containing {TokenKey} placeholders substituted at execute time.", IsRequired=true)
	Pipeline string `json:"pipeline"`
}

/** @description Test-run an aggregation pipeline with caller-supplied tokens */
// @Route("/{version}/database/aggregates/test", "POST")
// @Api(Description="Test-run an aggregation pipeline with caller-supplied tokens")
// @DataContract
type TestDatabaseAggregateRequest struct {
	CodeMashRequestBase
	/** @description Optional database integration id. When omitted, the project's default database integration is used. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id. When omitted, the project's default database integration is used.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
	/** @description Name of the collection (schema) to run the aggregation against. */
	// @DataMember
	// @ApiMember(Description="Name of the collection (schema) to run the aggregation against.", IsRequired=true)
	CollectionName string `json:"collectionName"`
	/** @description MongoDB aggregation pipeline JSON, optionally containing {TokenKey} placeholders to be substituted from tokens. */
	// @DataMember
	// @ApiMember(Description="MongoDB aggregation pipeline JSON, optionally containing {TokenKey} placeholders to be substituted from tokens.", IsRequired=true)
	Pipeline string `json:"pipeline"`
	/** @description Optional key/value substitutions for {TokenKey} placeholders in the pipeline. */
	// @DataMember
	// @ApiMember(Description="Optional key/value substitutions for {TokenKey} placeholders in the pipeline.")
	Tokens map[string]string `json:"tokens,omitempty"`
}

type MongoDbAggregateCreated struct {
	Aggregate MongoDbAggregate `json:"aggregate"`
}

type MongoDbAggregateUpdated struct {
	Aggregate MongoDbAggregate `json:"aggregate"`
}

type MongoDbAggregateDeleted struct {
	SchemaId SchemaId           `json:"schemaId"`
	Id       MongoDbAggregateId `json:"id"`
}

type DatabaseEstablished struct {
}

type DatabaseEnabled struct {
}

type DatabaseDisabled struct {
}

type DatabaseIntegrationSaved struct {
	Integration DatabaseIntegration `json:"integration"`
}

type DatabaseIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type DatabaseIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type DatabaseIntegrationSetAsDefault struct {
	Env Env           `json:"env"`
	Id  IntegrationId `json:"id"`
}

type DatabaseIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type DatabaseIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type DatabaseIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type DatabaseIntegrationProvisioningStarted struct {
	IntegrationId    IntegrationId `json:"integrationId"`
	AtlasProjectId   string        `json:"atlasProjectId"`
	AtlasClusterName string        `json:"atlasClusterName"`
}

type DatabaseIntegrationProvisioningCompleted struct {
	IntegrationId            IntegrationId `json:"integrationId"`
	ConnectionStringTemplate string        `json:"connectionStringTemplate"`
}

type DatabaseIntegrationProvisioningFailed struct {
	IntegrationId IntegrationId `json:"integrationId"`
	Reason        string        `json:"reason"`
	Retryable     bool          `json:"retryable,omitempty"`
}

type DatabaseIntegrationDeprovisioned struct {
	IntegrationId    IntegrationId `json:"integrationId"`
	AtlasProjectId   string        `json:"atlasProjectId"`
	AtlasClusterName string        `json:"atlasClusterName"`
}

type ProjectStatusChanged struct {
	Status ProjectStatus `json:"status,omitempty"`
}

type SchemaCreated struct {
	Schema Schema `json:"schema"`
}

type SchemaMirrored struct {
	Schema Schema `json:"schema"`
}

type SchemaDraftUpdated struct {
	Id    SchemaId    `json:"id"`
	Draft SchemaDraft `json:"draft"`
	Env   Env         `json:"env"`
}

type SchemaDraftDiscarded struct {
	Id  SchemaId `json:"id"`
	Env Env      `json:"env"`
}

type SchemaVersionPublished struct {
	Id      SchemaId               `json:"id"`
	Version PublishedSchemaVersion `json:"version"`
	Diff    SchemaDiff             `json:"diff"`
	Env     Env                    `json:"env"`
}

type SchemaSettingsUpdated struct {
	Id       SchemaId       `json:"id"`
	Settings SchemaSettings `json:"settings"`
	Env      Env            `json:"env"`
}

type SchemaDeleted struct {
	Id  SchemaId `json:"id"`
	Env Env      `json:"env"`
}

type SchemaRenamed struct {
	SchemaId         SchemaId   `json:"schemaId"`
	NewName          SchemaName `json:"newName"`
	RenameUniqueName bool       `json:"renameUniqueName,omitempty"`
	Env              Env        `json:"env"`
}

type SchemaDataCleared struct {
	Id           SchemaId        `json:"id"`
	Integrations []IntegrationId `json:"integrations"`
	Env          Env             `json:"env"`
}

type TaxonomyCreated struct {
	Taxonomy Taxonomy `json:"taxonomy"`
}

type TaxonomyUpdated struct {
	Taxonomy Taxonomy `json:"taxonomy"`
}

type TaxonomyDeleted struct {
	TaxonomyId TaxonomyId `json:"taxonomyId"`
}

type TaxonomyDataCleared struct {
	TaxonomyId   TaxonomyId      `json:"taxonomyId"`
	Integrations []IntegrationId `json:"integrations"`
}

type SchemaTriggerSaved struct {
	Trigger SchemaTrigger `json:"trigger"`
}

type DatabaseTriggerMirrored struct {
	Trigger Trigger `json:"trigger"`
}

type SchemaTriggerEnabled struct {
	TriggerByIdEventBase
	SchemaId SchemaId `json:"schemaId"`
	Env      Env      `json:"env"`
}

type SchemaTriggerDisabled struct {
	TriggerByIdEventBase
	SchemaId SchemaId `json:"schemaId"`
	Env      Env      `json:"env"`
}

type SchemaTriggerDeleted struct {
	TriggerByIdEventBase
	SchemaId SchemaId `json:"schemaId"`
	Env      Env      `json:"env"`
}

type ProcessCollectionImport struct {
	ImportId              string  `json:"importId"`
	ProjectId             string  `json:"projectId"`
	AccountId             string  `json:"accountId"`
	DatabaseIntegrationId string  `json:"databaseIntegrationId"`
	Env                   *string `json:"env,omitempty"`
}

type RecordInserted struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	Id                    string        `json:"id"`
	Document              Object        `json:"document"`
}

type RecordUpdated struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	Id                    string        `json:"id"`
	From                  Object        `json:"from"`
	To                    Object        `json:"to"`
}

type RecordDeleted struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	Id                    string        `json:"id"`
	Document              Object        `json:"document"`
}

type RecordReplaced struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	Id                    string        `json:"id"`
	From                  Object        `json:"from"`
	To                    Object        `json:"to"`
}

type RecordResponsibilityChanged struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	Id                    string        `json:"id"`
	FromOwner             AuthId        `json:"fromOwner"`
	ToOwner               AuthId        `json:"toOwner"`
}

type RecordsInserted struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	Ids                   IReadOnlyList `json:"ids"`
	Documents             IReadOnlyList `json:"documents"`
}

type RecordsUpdated struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	MatchedCount          int64         `json:"matchedCount,omitempty"`
	ModifiedCount         int64         `json:"modifiedCount,omitempty"`
	Update                Object        `json:"update"`
}

type RecordsDeleted struct {
	ProjectId             ProjectId     `json:"projectId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
	SchemaName            SchemaName    `json:"schemaName"`
	DeletedCount          int64         `json:"deletedCount,omitempty"`
	Filter                Object        `json:"filter"`
}

type EmailVerificationCodeRequested struct {
	Email        string    `json:"email"`
	ProjectId    string    `json:"projectId,omitempty"`
	Code         string    `json:"code"`
	ExpiresAtUtc time.Time `json:"expiresAtUtc,omitempty"`
}

type MagicLinkRequested struct {
	Email        string    `json:"email"`
	ProjectId    string    `json:"projectId,omitempty"`
	Token        string    `json:"token"`
	ExpiresAtUtc time.Time `json:"expiresAtUtc,omitempty"`
}

type PasswordResetRequested struct {
	Email        string    `json:"email"`
	ProjectId    string    `json:"projectId,omitempty"`
	Token        string    `json:"token"`
	ExpiresAtUtc time.Time `json:"expiresAtUtc,omitempty"`
}

type PasswordChanged struct {
	Email     string `json:"email"`
	ProjectId string `json:"projectId,omitempty"`
}

type SseCallTriggered struct {
	ProjectId        ProjectId            `json:"projectId"`
	AccountId        AccountId            `json:"accountId"`
	TriggerId        TriggerId            `json:"triggerId"`
	TriggerType      TriggerType          `json:"triggerType,omitempty"`
	SourceEvent      string               `json:"sourceEvent"`
	TargetUserAuthId *string              `json:"targetUserAuthId,omitempty"`
	SchemaId         *string              `json:"schemaId,omitempty"`
	TokenMappings    *IReadOnlyDictionary `json:"tokenMappings,omitempty"`
	CorrelationId    *string              `json:"correlationId,omitempty"`
}

type UserRegistered struct {
	Auth       Auth    `json:"auth"`
	LinkToUser *UserId `json:"linkToUser,omitempty"`
}

type UserCreated struct {
	UserId    UserId    `json:"userId"`
	ProjectId ProjectId `json:"projectId"`
	AuthId    *AuthId   `json:"authId,omitempty"`
}

type UserUpdated struct {
	AuthId AuthId          `json:"authId"`
	From   UserGeneralInfo `json:"from"`
	To     UserGeneralInfo `json:"to"`
}

type UserBlocked struct {
	User   *UserGeneralInfo `json:"user,omitempty"`
	AuthId AuthId           `json:"authId"`
}

type UserUnblocked struct {
	User   *UserGeneralInfo `json:"user,omitempty"`
	AuthId AuthId           `json:"authId"`
}

type UserInvited struct {
	EmailAddress EmailAddress `json:"emailAddress"`
}

type UserVerified struct {
	AuthId AuthId           `json:"authId"`
	User   *UserGeneralInfo `json:"user,omitempty"`
}

type UserDeleted struct {
	User   *UserGeneralInfo `json:"user,omitempty"`
	AuthId AuthId           `json:"authId"`
}

// @Route("/{version}/files/disable", "GET")
type DisableFiles struct {
	CodeMashRequestBase
}

// @Route("/{version}/files/enable", "GET")
type EnableFiles struct {
	CodeMashRequestBase
}

// @Route("/{version}/files/triggers/{triggerId}", "DELETE")
// @DataContract
type DeleteFilesTrigger struct {
	DeleteTrigger
}

// @Route("/{version}/files/triggers/{triggerId}/disable", "PATCH")
// @DataContract
type DisableFilesTrigger struct {
	DisableTrigger
}

// @Route("/{version}/files/triggers/{triggerId}/enable", "PATCH")
// @DataContract
type EnableFilesTrigger struct {
	EnableTrigger
}

/** @description Gets files trigger by specified Id */
// @Route("/{version}/files/triggers/{id}", "GET")
// @Api(Description="Gets files trigger by specified Id")
type GetFilesTrigger struct {
	GetTrigger
}

/** @description Gets files triggers */
// @Route("/{version}/files/triggers", "GET")
// @Api(Description="Gets files triggers")
type GetFilesTriggers struct {
	GetTriggers
}

// @Route("/{version}/files/triggers", "POST")
// @DataContract
type SaveFilesTrigger struct {
	SaveTrigger
}

// @Route("/{version}/files/integrations/{Id}", "DELETE")
type DeleteFilesIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Files integration id to delete, from get_files_integrations. */
	// @ApiMember(Description="Files integration id to delete, from get_files_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/files/integrations/{Id}/disable", "PUT")
type DisableFilesIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Files integration id to disable, from get_files_integrations. */
	// @ApiMember(Description="Files integration id to disable, from get_files_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/files/integrations/{Id}/enable", "PUT")
type EnableFilesIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Files integration id to enable, from get_files_integrations. */
	// @ApiMember(Description="Files integration id to enable, from get_files_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integration by specified Id */
// @Route("/{version}/files/integrations/{id}", "GET")
// @Api(Description="Gets integration by specified Id")
type GetFilesIntegration struct {
	CodeMashRequestBase
	/** @description Files integration id to fetch, from get_files_integrations. */
	// @ApiMember(Description="Files integration id to fetch, from get_files_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integrations */
// @Route("/{version}/files/integrations", "GET")
// @Api(Description="Gets integrations")
type GetFilesIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/files/integrations", "POST")
// @DataContract
type SaveFilesIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration FilesIntegrationRequest `json:"integration"`
}

// @Route("/{version}/files/integrations/{Id}/default", "PUT")
type SetFilesIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	/** @description Files integration id to set as default, from get_files_integrations. */
	// @ApiMember(Description="Files integration id to set as default, from get_files_integrations.", IsRequired=true)
	Id string `json:"id"`
}

// @Route("/{version}/files/integrations/test", "POST")
type TestFilesIntegration struct {
	CodeMashRequestBase
	/** @description Integration id, from get_files_integrations. */
	// @ApiMember(Description="Integration id, from get_files_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

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

type FilesEstablished struct {
}

type FilesEnabled struct {
}

type FilesDisabled struct {
}

type FilesIntegrationSaved struct {
	Integration FileIntegration `json:"integration"`
}

type FilesIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type FilesIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type FilesIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type FilesIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type FilesIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type FilesIntegrationSetAsDefault struct {
	Env Env           `json:"env"`
	Id  IntegrationId `json:"id"`
}

type FilesTriggerSaved struct {
	Trigger FileTrigger `json:"trigger"`
}

type FilesTriggerMirrored struct {
	Trigger Trigger `json:"trigger"`
}

type FilesTriggerEnabled struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type FilesTriggerDisabled struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type FilesTriggerDeleted struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type FileUploaded struct {
	ProjectId     ProjectId       `json:"projectId"`
	IntegrationId IntegrationId   `json:"integrationId"`
	FileRef       FileResourceRef `json:"fileRef"`
}

type FileDeleted struct {
	ProjectId     ProjectId     `json:"projectId"`
	IntegrationId IntegrationId `json:"integrationId"`
	Path          string        `json:"path"`
}

/** @description Disable email service */
// @Route("/{version}/notifications/email/disable", "GET")
// @Api(Description="Disable email service")
type DisableEmail struct {
	CodeMashRequestBase
}

/** @description Get email disable dependencies */
// @Route("/{version}/notifications/email/disable-dependencies", "GET")
// @Api(Description="Get email disable dependencies")
type GetEmailDisableDependencies struct {
	CodeMashRequestBase
}

/** @description Enable email service */
// @Route("/{version}/notifications/email/enable", "GET")
// @Api(Description="Enable email service")
type EnableEmail struct {
	CodeMashRequestBase
}

// @Route("/{version}/notifications/email/validation/integrations", "POST")
// @DataContract
type SaveEmailValidationIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration EmailValidationIntegrationRequest `json:"integration"`
}

// @Route("/{version}/notifications/email/validation/integrations/test", "POST")
type TestEmailValidationIntegration struct {
	CodeMashRequestBase
	IntegrationId string `json:"integrationId"`
}

/** @description Attach a file to an email template */
// @Route("/{version}/notifications/email/templates/attachments", "POST")
// @Api(Description="Attach a file to an email template")
type AttachFileToTemplateRequest struct {
	CodeMashRequestBase
	/** @description Optional language code to scope the attachment to a single translation. Omit to attach at the template level. */
	// @ApiMember(Description="Optional language code to scope the attachment to a single translation. Omit to attach at the template level.")
	Language *string `json:"language,omitempty"`
	/** @description The email template id to attach the file to. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to attach the file to. Get it from get_email_templates.", IsRequired=true)
	TemplateId string `json:"templateId"`
	/** @description The file resource reference to attach (from a prior file upload). */
	// @ApiMember(Description="The file resource reference to attach (from a prior file upload).", IsRequired=true)
	FileRef FileResourceRefDto `json:"fileRef"`
}

/** @description Create an email template */
// @Route("/{version}/notifications/email/templates", "POST")
// @Api(Description="Create an email template")
type CreateEmailTemplateRequest struct {
	SaveEmailTemplate
}

/** @description Delete an email template */
// @Route("/{version}/notifications/email/templates/{Id}", "DELETE")
// @Api(Description="Delete an email template")
type DeleteEmailTemplateRequest struct {
	CodeMashRequestBase
	/** @description The email template id to delete. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to delete. Get it from get_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get an email template */
// @Route("/{version}/notifications/email/templates/{id}", "GET")
// @Api(Description="Get an email template")
type GetEmailTemplate struct {
	CodeMashRequestBase
	/** @description The email template id to fetch. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to fetch. Get it from get_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets email templates */
// @Route("/{version}/notifications/email/templates", "GET")
// @Api(Description="Gets email templates")
type GetEmailTemplates struct {
	CodeMashListPaginationRequestBase
	/** @description Set true to include archived templates. */
	// @ApiMember(Description="Set true to include archived templates.")
	ShowArchived *bool `json:"showArchived,omitempty"`
	/** @description Optional: return only the template with this id. */
	// @ApiMember(Description="Optional: return only the template with this id.")
	TemplateId *string `json:"templateId,omitempty"`
}

/** @description Render MJML email template */
// @Route("/{version}/notifications/email/templates/mjml", "POST")
// @Api(Description="Render MJML email template")
type GetMjml struct {
	CodeMashRequestBase
	/** @description The MJML/Razor template source code to render. */
	// @ApiMember(Description="The MJML/Razor template source code to render.", IsRequired=true)
	Code string `json:"code"`
	/** @description Optional token values to bind into the template while rendering. */
	// @ApiMember(Description="Optional token values to bind into the template while rendering.")
	Tokens []TokenMappingDto `json:"tokens,omitempty"`
	/** @description Set true when rendering for a preview (vs. a final save), to affect how missing tokens are handled. */
	// @ApiMember(Description="Set true when rendering for a preview (vs. a final save), to affect how missing tokens are handled.")
	IsForPreview bool `json:"isForPreview,omitempty"`
}

/** @description Get a system email template */
// @Route("/{version}/notifications/email/system-templates/{id}", "GET")
// @Api(Description="Get a system email template")
type GetSystemEmailTemplate struct {
	CodeMashRequestBase
	/** @description The system email template id to fetch. Get it from get_system_email_templates. */
	// @ApiMember(Description="The system email template id to fetch. Get it from get_system_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get system email templates */
// @Route("/{version}/notifications/email/system-templates", "GET")
// @Api(Description="Get system email templates")
type GetSystemEmailTemplates struct {
	CodeMashListPaginationRequestBase
	/** @description Optional group tags to filter templates by (e.g. newsletter, onboarding). */
	// @ApiMember(Description="Optional group tags to filter templates by (e.g. newsletter, onboarding).")
	GroupTags []string `json:"groupTags,omitempty"`
	/** @description Optional visual themes to filter templates by. */
	// @ApiMember(Description="Optional visual themes to filter templates by.")
	Themes []string `json:"themes,omitempty"`
	/** @description Optional communication channel to filter templates by (e.g. Transactional, Marketing). */
	// @ApiMember(Description="Optional communication channel to filter templates by (e.g. Transactional, Marketing).")
	CommunicationChannel *CommunicationChannel `json:"communicationChannel,omitempty"`
	/** @description Optional trigger type to filter templates that are designed for a specific automated trigger. */
	// @ApiMember(Description="Optional trigger type to filter templates that are designed for a specific automated trigger.")
	ForTrigger *TriggerType `json:"forTrigger,omitempty"`
}

/** @description Gets the tokens used by an email template */
// @Route("/{version}/notifications/email/templates/{id}/tokens", "GET")
// @Api(Description="Gets the tokens used by an email template")
type GetEmailTemplateAvailableTokens struct {
	CodeMashRequestBase
	/** @description Template id from get_email_templates. */
	// @ApiMember(Description="Template id from get_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Update an email template */
// @Route("/{version}/notifications/email/templates", "PUT")
// @Api(Description="Update an email template")
type UpdateEmailTemplateRequest struct {
	SaveEmailTemplate
	/** @description The email template id to update. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to update. Get it from get_email_templates.", IsRequired=true)
	ViewId string `json:"viewId"`
}

/** @description Delete an email signature */
// @Route("/{version}/notifications/email/signatures/{id}", "DELETE")
// @Api(Description="Delete an email signature")
type DeleteEmailSignature struct {
	CodeMashRequestBase
	/** @description The email signature id to delete. Get it from get_email_signatures. */
	// @ApiMember(Description="The email signature id to delete. Get it from get_email_signatures.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get an email signature */
// @Route("/{version}/notifications/email/signatures/{id}", "GET")
// @Api(Description="Get an email signature")
type GetEmailSignature struct {
	CodeMashRequestBase
	/** @description The email signature id to fetch. Get it from get_email_signatures. */
	// @ApiMember(Description="The email signature id to fetch. Get it from get_email_signatures.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get email signatures */
// @Route("/{version}/notifications/email/signatures", "GET")
// @Api(Description="Get email signatures")
type GetEmailSignatures struct {
	CodeMashListPaginationRequestBase
}

/** @description Save an email signature */
// @Route("/{version}/notifications/email/signatures", "POST")
// @Api(Description="Save an email signature")
type SaveEmailSignatureRequest struct {
	CodeMashRequestBase
	/** @description The signature id to update. Omit to create a new signature. Get it from get_email_signatures. */
	// @ApiMember(Description="The signature id to update. Omit to create a new signature. Get it from get_email_signatures.")
	ViewId *string `json:"viewId,omitempty"`
	/** @description The display name of the signature. */
	// @ApiMember(Description="The display name of the signature.", IsRequired=true)
	DisplayName string `json:"displayName"`
	/** @description The per-language content translations for this signature. */
	// @ApiMember(Description="The per-language content translations for this signature.", IsRequired=true)
	Translations []TranslationDto `json:"translations"`
}

/** @description Get email settings */
// @Route("/{version}/notifications/email/settings", "GET")
// @Api(Description="Get email settings")
type GetEmailSettings struct {
	CodeMashRequestBase
	/** @description Unused legacy field; leave empty. */
	// @ApiMember(Description="Unused legacy field; leave empty.")
	Id string `json:"id"`
}

/** @description Confirm human delivery of a test email */
// @Route("/{version}/notifications/email/integrations/confirm-human-delivery", "POST")
// @Api(Description="Confirm human delivery of a test email")
// @DataContract
type ConfirmEmailIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	/** @description The email integration id the test delivery was confirmed for. Get it from get_email_integrations. */
	// @DataMember
	// @ApiMember(Description="The email integration id the test delivery was confirmed for. Get it from get_email_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

/** @description Delete an email integration */
// @Route("/{version}/notifications/email/integrations/{Id}", "DELETE")
// @Api(Description="Delete an email integration")
type DeleteEmailIntegration struct {
	CodeMashRequestBase
	/** @description The email integration id to delete. Get it from get_email_integrations. */
	// @ApiMember(Description="The email integration id to delete. Get it from get_email_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Disable an email integration */
// @Route("/{version}/notifications/email/integrations/{Id}/disable", "PUT")
// @Api(Description="Disable an email integration")
type DisableEmailIntegration struct {
	CodeMashRequestBase
	/** @description The email integration id to disable. Get it from get_email_integrations. */
	// @ApiMember(Description="The email integration id to disable. Get it from get_email_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Check email integration domain health */
// @Route("/{version}/notifications/email/integrations/domain-health", "POST")
// @Api(Description="Check email integration domain health")
// @DataContract
type CheckEmailIntegrationDomainHealthRequest struct {
	CodeMashRequestBase
	/** @description The email integration id to check DNS health for. Get it from get_email_integrations. */
	// @DataMember
	// @ApiMember(Description="The email integration id to check DNS health for. Get it from get_email_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

/** @description Enable an email integration */
// @Route("/{version}/notifications/email/integrations/{Id}/enable", "PUT")
// @Api(Description="Enable an email integration")
type EnableEmailIntegration struct {
	CodeMashRequestBase
	/** @description The email integration id to enable. Get it from get_email_integrations. */
	// @ApiMember(Description="The email integration id to enable. Get it from get_email_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get an email integration */
// @Route("/{version}/notifications/email/integrations/{id}", "GET")
// @Api(Description="Get an email integration")
type GetEmailIntegration struct {
	CodeMashRequestBase
	/** @description The email integration id to fetch. Get it from get_email_integrations. */
	// @ApiMember(Description="The email integration id to fetch. Get it from get_email_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets email integrations */
// @Route("/{version}/notifications/email/integrations", "GET")
// @Api(Description="Gets email integrations")
type GetEmailIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/notifications/email/integrations", "POST")
// @DataContract
type SaveEmailIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration EmailIntegrationRequest `json:"integration"`
}

/** @description Set an email integration as default */
// @Route("/{version}/notifications/email/integrations/{Id}/default", "PUT")
// @Api(Description="Set an email integration as default")
type SetEmailsIntegrationAsDefault struct {
	CodeMashRequestBase
	/** @description The email integration id to set as default. Get it from get_email_integrations. */
	// @ApiMember(Description="The email integration id to set as default. Get it from get_email_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Test an email integration */
// @Route("/{version}/notifications/email/integrations/test", "POST")
// @Api(Description="Test an email integration")
type TestEmailIntegration struct {
	CodeMashRequestBase
	/** @description The email integration id to test. Get it from get_email_integrations. */
	// @ApiMember(Description="The email integration id to test. Get it from get_email_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
	/** @description The recipient email address to send the test email to. */
	// @ApiMember(Description="The recipient email address to send the test email to.", IsRequired=true)
	To string `json:"to"`
}

/** @description Archive an email template */
// @Route("/{version}/notifications/email/templates/{Id}/archive", "PUT")
// @Api(Description="Archive an email template")
type ArchiveEmailTemplateRequest struct {
	CodeMashRequestBase
	/** @description The email template id to archive. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to archive. Get it from get_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Clone an email template */
// @Route("/{version}/notifications/email/templates/{Id}/clone", "POST")
// @Api(Description="Clone an email template")
type CloneEmailTemplateRequest struct {
	CodeMashRequestBase
	/** @description The email template id to clone. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to clone. Get it from get_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Un-archive an email template */
// @Route("/{version}/notifications/email/templates/{Id}/unarchive", "PUT")
// @Api(Description="Un-archive an email template")
type UnArchiveEmailTemplateRequest struct {
	CodeMashRequestBase
	/** @description The email template id to unarchive. Get it from get_email_templates. */
	// @ApiMember(Description="The email template id to unarchive. Get it from get_email_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Delete an email footer */
// @Route("/{version}/notifications/email/footers/{id}", "DELETE")
// @Api(Description="Delete an email footer")
type DeleteEmailFooter struct {
	CodeMashRequestBase
	/** @description The email footer id to delete. Get it from get_email_footers. */
	// @ApiMember(Description="The email footer id to delete. Get it from get_email_footers.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get an email footer */
// @Route("/{version}/notifications/email/footers/{id}", "GET")
// @Api(Description="Get an email footer")
type GetEmailFooter struct {
	CodeMashRequestBase
	/** @description The email footer id to fetch. Get it from get_email_footers. */
	// @ApiMember(Description="The email footer id to fetch. Get it from get_email_footers.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Get email footers */
// @Route("/{version}/notifications/email/footers", "GET")
// @Api(Description="Get email footers")
type GetEmailFooters struct {
	CodeMashListPaginationRequestBase
}

/** @description Save an email footer */
// @Route("/{version}/notifications/email/footers", "POST")
// @Api(Description="Save an email footer")
type SaveEmailFooterRequest struct {
	CodeMashRequestBase
	/** @description The footer id to update. Omit to create a new footer. Get it from get_email_footers. */
	// @ApiMember(Description="The footer id to update. Omit to create a new footer. Get it from get_email_footers.")
	ViewId *string `json:"viewId,omitempty"`
	/** @description The display name of the footer. */
	// @ApiMember(Description="The display name of the footer.", IsRequired=true)
	DisplayName string `json:"displayName"`
	/** @description The per-language content translations for this footer. */
	// @ApiMember(Description="The per-language content translations for this footer.", IsRequired=true)
	Translations []TranslationDto `json:"translations"`
}

/** @description This endpoint implements the RFC 8058 one-click unsubscribe flow used by mailbox providers. */
// @Route("/{version}/email/one-click-unsubscribe", "POST")
// @Api(Description="This endpoint implements the RFC 8058 one-click unsubscribe flow used by mailbox providers.")
// @DataContract
type OneClickUnsubscribeRequest struct {
	RequestBase
	/** @description Encrypted unsubscribe token. The campaign batcher embedded this value in the List-Unsubscribe header. */
	// @DataMember
	// @ApiMember(Description="Encrypted unsubscribe token. The campaign batcher embedded this value in the List-Unsubscribe header.", IsRequired=true, Name="token", ParameterType="query")
	Token string `json:"token"`
}

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
}

/** @description Deletes emails campaign from queue */
// @Route("/{version}/notifications/email/campaigns/{Id}", "DELETE")
// @Api(Description="Deletes emails campaign from queue")
// @DataContract
type DeleteEmailCampaignRequest struct {
	CodeMashRequestBase
}

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

/** @description Get email campaign batches */
// @Route("/{version}/notifications/email/campaigns/{id}/batches", "GET")
// @Api(Description="Get email campaign batches")
type GetEmailCampaignBatches struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id to list batches for. Get it from get_all_email_campaigns. */
	// @ApiMember(Description="The email campaign id to list batches for. Get it from get_all_email_campaigns.", IsRequired=true)
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

/** @description Get an email campaign batch notification */
// @Route("/{version}/notifications/email/campaigns/{id}/batches/{batchId}/{notificationId}", "GET")
// @Api(Description="Get an email campaign batch notification")
type GetEmailCampaignBatchNotification struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id. Get it from get_all_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_all_email_campaigns.", IsRequired=true)
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

/** @description Get email campaign batch notifications */
// @Route("/{version}/notifications/email/campaigns/{id}/batches/{batchId}", "GET")
// @Api(Description="Get email campaign batch notifications")
type GetEmailCampaignBatchNotifications struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id. Get it from get_all_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_all_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description The campaign batch id to list notifications for. Get it from get_email_campaign_batches. */
	// @ApiMember(Description="The campaign batch id to list notifications for. Get it from get_email_campaign_batches.", IsRequired=true)
	BatchId string `json:"batchId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Get email campaign statistics */
// @Route("/{version}/notifications/email/campaigns/{id}/stats", "GET")
// @Api(Description="Get email campaign statistics")
type GetEmailCampaignStatistics struct {
	CodeMashRequestBase
	/** @description The email campaign id to get statistics for. Get it from get_all_email_campaigns. */
	// @ApiMember(Description="The email campaign id to get statistics for. Get it from get_all_email_campaigns.", IsRequired=true)
	Id string `json:"id"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Preview an email notification */
// @Route("/{version}/notifications/email/preview", "GET")
// @Api(Description="Preview an email notification")
type PreviewEmailNotification struct {
	RequestBase
	/** @description The opaque, pre-signed preview hash identifying the project and notification to preview. */
	// @ApiMember(Description="The opaque, pre-signed preview hash identifying the project and notification to preview.", IsRequired=true)
	Hash string `json:"hash"`
}

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

/** @description Get an email campaign message */
// @Route("/{version}/notifications/emails/campaigns/{campaignId}/messages/{id}", "GET")
// @Api(Description="Get an email campaign message")
type GetEmailCampaignMessage struct {
	CodeMashRequestBase
	/** @description The email campaign id. Get it from get_all_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_all_email_campaigns.", IsRequired=true)
	CampaignId string `json:"campaignId"`
	/** @description The campaign batch id. Get it from get_email_campaign_batches. */
	// @ApiMember(Description="The campaign batch id. Get it from get_email_campaign_batches.", IsRequired=true)
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description The notification (message) id to fetch. Get it from get_email_campaign_messages. */
	// @ApiMember(Description="The notification (message) id to fetch. Get it from get_email_campaign_messages.", IsRequired=true)
	NotificationId string `json:"notificationId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

/** @description Get email campaign messages */
// @Route("/{version}/notifications/emails/campaigns/{campaignId}/messages", "GET")
// @Api(Description="Get email campaign messages")
type GetEmailCampaignMessagesRequest struct {
	CodeMashListPaginationRequestBase
	/** @description The email campaign id. Get it from get_all_email_campaigns. */
	// @ApiMember(Description="The email campaign id. Get it from get_all_email_campaigns.", IsRequired=true)
	CampaignId string `json:"campaignId"`
	/** @description The campaign batch id to list messages for. Get it from get_email_campaign_batches. */
	// @ApiMember(Description="The campaign batch id to list messages for. Get it from get_email_campaign_batches.", IsRequired=true)
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

type EmailServiceEstablished struct {
}

type ProjectDatabaseConnected struct {
	Env Env `json:"env"`
}

type EmailServiceEnabled struct {
}

type EmailServiceDisabled struct {
}

type EmailFooterSaved struct {
	Id           EmailFooterId        `json:"id"`
	Name         DisplayName          `json:"name"`
	Translations []MessageTranslation `json:"translations"`
	Env          *Env                 `json:"env,omitempty"`
}

type EmailFooterMirrored struct {
	Footer EmailFooter `json:"footer"`
}

type EmailFooterDeleted struct {
	Id  EmailFooterId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailIntegrationSaved struct {
	Integration EmailIntegration `json:"integration"`
}

type EmailIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type EmailIntegrationHumanDeliveryConfirmed struct {
	Id             IntegrationId `json:"id"`
	ConfirmedAtUtc time.Time     `json:"confirmedAtUtc,omitempty"`
}

type EmailIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type EmailIntegrationSetAsDefault struct {
	Env Env           `json:"env"`
	Id  IntegrationId `json:"id"`
}

type EmailIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailSignatureSaved struct {
	Id           EmailSignatureId     `json:"id"`
	Name         DisplayName          `json:"name"`
	Translations []MessageTranslation `json:"translations"`
	Env          *Env                 `json:"env,omitempty"`
}

type EmailSignatureMirrored struct {
	Signature EmailSignature `json:"signature"`
}

type EmailSignatureDeleted struct {
	Id  EmailSignatureId `json:"id"`
	Env *Env             `json:"env,omitempty"`
}

type EmailTemplateCreated struct {
	TemplateId                  TemplateId           `json:"templateId"`
	DisplayName                 DisplayName          `json:"displayName"`
	Translations                []MessageTranslation `json:"translations"`
	Channel                     CommunicationChannel `json:"channel,omitempty"`
	Description                 *string              `json:"description,omitempty"`
	Tags                        []Tag                `json:"tags,omitempty"`
	LanguageAgnosticAttachments []FileResourceRef    `json:"languageAgnosticAttachments,omitempty"`
	Env                         *Env                 `json:"env,omitempty"`
}

type EmailTemplateUpdated struct {
	TemplateId                  TemplateId           `json:"templateId"`
	DisplayName                 DisplayName          `json:"displayName"`
	Translations                []MessageTranslation `json:"translations"`
	Channel                     CommunicationChannel `json:"channel,omitempty"`
	Description                 *string              `json:"description,omitempty"`
	Tags                        []Tag                `json:"tags,omitempty"`
	LanguageAgnosticAttachments []FileResourceRef    `json:"languageAgnosticAttachments,omitempty"`
	AttachmentsToBeDeleted      []FileResourceRef    `json:"attachmentsToBeDeleted,omitempty"`
	Env                         *Env                 `json:"env,omitempty"`
}

type EmailTemplateMirrored struct {
	Template EmailTemplate `json:"template"`
}

type EmailTemplateBackfilled struct {
	Template EmailTemplate `json:"template"`
}

type EmailTemplateDeleted struct {
	TemplateId        TemplateId        `json:"templateId"`
	FilesToBeDeleted  []FileResourceRef `json:"filesToBeDeleted,omitempty"`
	FileIntegrationId *IntegrationId    `json:"fileIntegrationId,omitempty"`
	Env               *Env              `json:"env,omitempty"`
}

type EmailTemplateArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type EmailTemplateUnArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type EmailValidationIntegrationSaved struct {
	Integration EmailValidationIntegration `json:"integration"`
}

type EmailValidationIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailValidationIntegrationSecretsConfigured struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailValidationIntegrationSecretsConfigurationFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type EmailValidationIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type EmailBatchRegistered struct {
	ProjectId       ProjectId       `json:"projectId"`
	CampaignId      CampaignId      `json:"campaignId"`
	CampaignBatchId CampaignBatchId `json:"campaignBatchId"`
	StartingAfter   *string         `json:"startingAfter,omitempty"`
}

type EmailNotificationRead struct {
	ProjectId       ProjectId       `json:"projectId"`
	CampaignId      CampaignId      `json:"campaignId"`
	CampaignBatchId CampaignBatchId `json:"campaignBatchId"`
	NotificationId  NotificationId  `json:"notificationId"`
}

type EmailNotificationClicked struct {
	ProjectId       ProjectId       `json:"projectId"`
	CampaignId      CampaignId      `json:"campaignId"`
	CampaignBatchId CampaignBatchId `json:"campaignBatchId"`
	NotificationId  NotificationId  `json:"notificationId"`
	SourceId        *string         `json:"sourceId,omitempty"`
}

type EmailCampaignStarted struct {
	ProjectId  ProjectId  `json:"projectId"`
	CampaignId CampaignId `json:"campaignId"`
}

type EmailCampaignStopped struct {
	ProjectId  ProjectId           `json:"projectId"`
	CampaignId CampaignId          `json:"campaignId"`
	Reason     *CampaignStopReason `json:"reason,omitempty"`
}

type EmailCampaignCompleted struct {
	ProjectId  ProjectId  `json:"projectId"`
	CampaignId CampaignId `json:"campaignId"`
	Errors     []ErrorDto `json:"errors,omitempty"`
}

type EmailCampaignFailed struct {
	ProjectId  ProjectId  `json:"projectId"`
	CampaignId CampaignId `json:"campaignId"`
	Errors     []ErrorDto `json:"errors"`
}

type EmailCampaignTriggered struct {
	ProjectId     ProjectId            `json:"projectId"`
	TriggerId     TriggerId            `json:"triggerId"`
	TriggerType   TriggerType          `json:"triggerType,omitempty"`
	SourceEvent   string               `json:"sourceEvent"`
	SchemaId      *string              `json:"schemaId,omitempty"`
	TokenMappings *IReadOnlyDictionary `json:"tokenMappings,omitempty"`
}

type EmailDeliveryEventReceived struct {
	ProjectId         ProjectId              `json:"projectId"`
	IntegrationId     IntegrationId          `json:"integrationId"`
	Recipient         EmailAddress           `json:"recipient"`
	Type              EmailDeliveryEventType `json:"type,omitempty"`
	OccurredAt        time.Time              `json:"occurredAt,omitempty"`
	ProviderMessageId *string                `json:"providerMessageId,omitempty"`
	Reason            *string                `json:"reason,omitempty"`
}

/** @description Disable SMS service */
// @Route("/{version}/notifications/sms/disable", "GET")
// @Api(Description="Disable SMS service")
type DisableSms struct {
	CodeMashRequestBase
}

/** @description Lists SMS-module dependencies shown before disable */
// @Route("/{version}/notifications/sms/disable-dependencies", "GET")
// @Api(Description="Lists SMS-module dependencies shown before disable")
type GetSmsDisableDependencies struct {
	CodeMashRequestBase
}

/** @description Enable SMS service */
// @Route("/{version}/notifications/sms/enable", "GET")
// @Api(Description="Enable SMS service")
type EnableSms struct {
	CodeMashRequestBase
}

/** @description Archives sms template */
// @Route("/{version}/notifications/sms/templates/{Id}/archive", "PUT")
// @Api(Description="Archives sms template")
type ArchiveSmsTemplateRequest struct {
	CodeMashRequestBase
	/** @description The SMS template id to archive. Get it from get_sms_templates. */
	// @ApiMember(Description="The SMS template id to archive. Get it from get_sms_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Clones sms template */
// @Route("/{version}/notifications/sms/templates/{Id}/clone", "POST")
// @Api(Description="Clones sms template")
type CloneSmsTemplateRequest struct {
	CodeMashRequestBase
	/** @description The SMS template id to clone. Get it from get_sms_templates. */
	// @ApiMember(Description="The SMS template id to clone. Get it from get_sms_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Create SMS template */
// @Route("/{version}/notifications/sms/templates", "POST")
// @Api(Description="Create SMS template")
type CreateSmsTemplateRequest struct {
	SaveSmsTemplate
}

/** @description Delete Sms Template for particular project */
// @Route("/{version}/notifications/sms/templates/{Id}", "DELETE")
// @Api(Description="Delete Sms Template for particular project")
type DeleteSmsTemplateRequest struct {
	CodeMashRequestBase
	/** @description The SMS template id to delete. Get it from get_sms_templates. */
	// @ApiMember(Description="The SMS template id to delete. Get it from get_sms_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets sms template by id */
// @Route("/{version}/notifications/sms/templates/{id}", "GET")
// @Api(Description="Gets sms template by id")
type GetSmsTemplate struct {
	CodeMashRequestBase
	/** @description The SMS template id to fetch. Get it from get_sms_templates. */
	// @ApiMember(Description="The SMS template id to fetch. Get it from get_sms_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets sms templates */
// @Route("/{version}/notifications/sms/templates", "GET")
// @Api(Description="Gets sms templates")
type GetSmsTemplates struct {
	CodeMashListPaginationRequestBase
	/** @description Set true to include archived templates. */
	// @ApiMember(Description="Set true to include archived templates.")
	ShowArchived *bool `json:"showArchived,omitempty"`
	/** @description Optional: return only the template with this id. */
	// @ApiMember(Description="Optional: return only the template with this id.")
	TemplateId *string `json:"templateId,omitempty"`
}

/** @description Goes through the Sms template and returns all the tokens that are used in the template translations */
// @Route("/{version}/notifications/sms/templates/{id}/tokens", "GET")
// @Api(Description="Goes through the Sms template and returns all the tokens that are used in the template translations")
type GetSmsMessageContentTokens struct {
	CodeMashRequestBase
	/** @description The SMS template id. Get it from get_sms_templates. */
	// @ApiMember(Description="The SMS template id. Get it from get_sms_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Runs the SMS Razor template, returns the bound text or the list of unresolved tokens. */
// @Route("/{version}/notifications/sms/templates/render", "POST")
// @Api(Description="Runs the SMS Razor template, returns the bound text or the list of unresolved tokens.")
type RenderSms struct {
	CodeMashRequestBase
	/** @description The Razor SMS template code to render. */
	// @ApiMember(Description="The Razor SMS template code to render.")
	Code string `json:"code"`
	/** @description Token name/value pairs to bind into the template. */
	// @ApiMember(Description="Token name/value pairs to bind into the template.")
	Tokens []TokenMappingDto `json:"tokens,omitempty"`
	/** @description Set true when rendering for a preview (relaxes some validation). */
	// @ApiMember(Description="Set true when rendering for a preview (relaxes some validation).")
	IsForPreview bool `json:"isForPreview,omitempty"`
}

/** @description Un-archives sms template */
// @Route("/{version}/notifications/sms/templates/{Id}/unarchive", "PUT")
// @Api(Description="Un-archives sms template")
type UnArchiveSmsTemplateRequest struct {
	CodeMashRequestBase
	/** @description The SMS template id to unarchive. */
	// @ApiMember(Description="The SMS template id to unarchive.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Edit sms template */
// @Route("/{version}/notifications/sms/templates", "PUT")
// @Api(Description="Edit sms template")
type UpdateSmsTemplateRequest struct {
	SaveSmsTemplate
	/** @description The SMS template id to update. Get it from get_sms_templates. */
	// @ApiMember(Description="The SMS template id to update. Get it from get_sms_templates.", IsRequired=true)
	ViewId string `json:"viewId"`
}

/** @description Gets SMS settings */
// @Route("/{version}/notifications/sms/settings", "GET")
// @Api(Description="Gets SMS settings")
type GetSmsSettings struct {
	CodeMashRequestBase
	/** @description Unused legacy parameter; leave empty. */
	// @ApiMember(Description="Unused legacy parameter; leave empty.")
	Id string `json:"id"`
}

/** @description Confirm that you received the test SMS delivery. */
// @Route("/{version}/notifications/sms/integrations/confirm-human-delivery", "POST")
// @Api(Description="Confirm that you received the test SMS delivery.")
// @DataContract
type ConfirmSmsIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	/** @description The SMS integration id being verified. Get it from get_sms_integrations. */
	// @DataMember
	// @ApiMember(Description="The SMS integration id being verified. Get it from get_sms_integrations.")
	IntegrationId string `json:"integrationId"`
}

/** @description Delete integration for particular project */
// @Route("/{version}/notifications/sms/integrations/{Id}", "DELETE")
// @Api(Description="Delete integration for particular project")
type DeleteSmsIntegrationRequest struct {
	CodeMashRequestBase
	/** @description The SMS integration id to delete. Get it from get_sms_integrations. */
	// @ApiMember(Description="The SMS integration id to delete. Get it from get_sms_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Disable integration for particular project */
// @Route("/{version}/notifications/sms/integrations/{Id}/disable", "PUT")
// @Api(Description="Disable integration for particular project")
type DisableSmsIntegrationRequest struct {
	CodeMashRequestBase
	/** @description The SMS integration id to disable. Get it from get_sms_integrations. */
	// @ApiMember(Description="The SMS integration id to disable. Get it from get_sms_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Enable integration for particular project */
// @Route("/{version}/notifications/sms/integrations/{Id}/enable", "PUT")
// @Api(Description="Enable integration for particular project")
type EnableSmsIntegrationRequest struct {
	CodeMashRequestBase
	/** @description The SMS integration id to enable. Get it from get_sms_integrations. */
	// @ApiMember(Description="The SMS integration id to enable. Get it from get_sms_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integration by specified Id */
// @Route("/{version}/notifications/sms/integrations/{id}", "GET")
// @Api(Description="Gets integration by specified Id")
type GetSmsIntegration struct {
	CodeMashRequestBase
	/** @description The SMS integration id to fetch. Get it from get_sms_integrations. */
	// @ApiMember(Description="The SMS integration id to fetch. Get it from get_sms_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets sms integrations */
// @Route("/{version}/notifications/sms/integrations", "GET")
// @Api(Description="Gets sms integrations")
type GetSmsIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/notifications/sms/integrations", "POST")
// @DataContract
type SaveSmsIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration SmsIntegrationRequest `json:"integration"`
}

/** @description Sets integration as default */
// @Route("/{version}/notifications/sms/integrations/{Id}/default", "PUT")
// @Api(Description="Sets integration as default")
type SetSmsIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	/** @description The SMS integration id to set as default. Get it from get_sms_integrations. */
	// @ApiMember(Description="The SMS integration id to set as default. Get it from get_sms_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Test SMS integration */
// @Route("/{version}/notifications/sms/integrations/test", "POST")
// @Api(Description="Test SMS integration")
type TestSmsIntegration struct {
	CodeMashRequestBase
	/** @description The SMS integration id to test. Get it from get_sms_integrations. */
	// @ApiMember(Description="The SMS integration id to test. Get it from get_sms_integrations.")
	IntegrationId string `json:"integrationId"`
	/** @description Optional phone number (international format) to send the test SMS to. */
	// @ApiMember(Description="Optional phone number (international format) to send the test SMS to.")
	To *string `json:"to,omitempty"`
}

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
	// @DataMember
	Collection *SmsToCollectionRecordsDeliverySettingsDto `json:"collection,omitempty"`
	/** @description For deliveryType 'PhoneNumbers'. JSON object: {"recipientsSourceType":"PhoneNumbers","phoneNumbers":["+37060000000"],"campaignTime":<unix seconds UTC>}. Numbers in international format. */
	// @DataMember
	// @ApiMember(Description="For deliveryType 'PhoneNumbers'. JSON object: {\"recipientsSourceType\":\"PhoneNumbers\",\"phoneNumbers\":[\"+37060000000\"],\"campaignTime\":<unix seconds UTC>}. Numbers in international format.")
	PhoneNumbers *SmsToPhoneNumbersDeliverySettingsDto `json:"phoneNumbers,omitempty"`
}

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

/** @description Gets sms campaigns */
// @Route("/{version}/notifications/sms/campaigns", "GET")
// @Api(Description="Gets sms campaigns")
type GetSmsCampaigns struct {
	CodeMashListPaginationRequestBase
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
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

/** @description Returns SMS preview notification body */
// @Route("/{version}/notifications/sms/preview", "GET")
// @Api(Description="Returns SMS preview notification body")
type PreviewSmsNotification struct {
	RequestBase
	/** @description Signed preview hash identifying the notification to render. */
	// @ApiMember(Description="Signed preview hash identifying the notification to render.")
	Hash string `json:"hash"`
}

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

/** @description Gets campaign sms message details */
// @Route("/{version}/notifications/sms/campaigns/{campaignId}/messages/{id}", "GET")
// @Api(Description="Gets campaign sms message details")
type GetSmsCampaignMessage struct {
	CodeMashRequestBase
	/** @description The campaign id. Get it from get_sms_campaigns. */
	// @ApiMember(Description="The campaign id. Get it from get_sms_campaigns.", IsRequired=true)
	CampaignId string `json:"campaignId"`
	/** @description The campaign batch id. Get it from get_sms_campaign_batches. */
	// @ApiMember(Description="The campaign batch id. Get it from get_sms_campaign_batches.", IsRequired=true)
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description The notification (message) id. Get it from get_sms_campaign_messages. */
	// @ApiMember(Description="The notification (message) id. Get it from get_sms_campaign_messages.", IsRequired=true)
	NotificationId string `json:"notificationId"`
	/** @description Optional. Omit to use the project default database integration (resolved per environment). */
	// @ApiMember(Description="Optional. Omit to use the project default database integration (resolved per environment).")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

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

type SmsIntegrationSaved struct {
	Integration SmsIntegration `json:"integration"`
}

type SmsIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type SmsIntegrationHumanDeliveryConfirmed struct {
	Id             IntegrationId `json:"id"`
	ConfirmedAtUtc time.Time     `json:"confirmedAtUtc,omitempty"`
}

type SmsIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type SmsIntegrationSetAsDefault struct {
	Env Env           `json:"env"`
	Id  IntegrationId `json:"id"`
}

type SmsIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type SmsIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type SmsIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type SmsServiceEstablished struct {
	DefaultTemplates []SmsTemplate `json:"defaultTemplates,omitempty"`
}

type SmsServiceEnabled struct {
}

type SmsServiceDisabled struct {
}

type SmsTemplateCreated struct {
	TemplateId   TemplateId           `json:"templateId"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Channel      CommunicationChannel `json:"channel,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Tags         []Tag                `json:"tags,omitempty"`
	Env          *Env                 `json:"env,omitempty"`
}

type SmsTemplateUpdated struct {
	TemplateId   TemplateId           `json:"templateId"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Channel      CommunicationChannel `json:"channel,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Tags         []Tag                `json:"tags,omitempty"`
	Env          *Env                 `json:"env,omitempty"`
}

type SmsTemplateMirrored struct {
	Template SmsTemplate `json:"template"`
}

type SmsTemplateDeleted struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type SmsTemplateArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type SmsTemplateUnArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type SmsBatchRegistered struct {
	CampaignId      CampaignId      `json:"campaignId"`
	CampaignBatchId CampaignBatchId `json:"campaignBatchId"`
	StartingAfter   *string         `json:"startingAfter,omitempty"`
}

type SmsNotificationRead struct {
	CampaignId      CampaignId      `json:"campaignId"`
	CampaignBatchId CampaignBatchId `json:"campaignBatchId"`
	NotificationId  NotificationId  `json:"notificationId"`
}

type SmsNotificationClicked struct {
	CampaignId      CampaignId      `json:"campaignId"`
	CampaignBatchId CampaignBatchId `json:"campaignBatchId"`
	NotificationId  NotificationId  `json:"notificationId"`
	SourceId        *string         `json:"sourceId,omitempty"`
}

type SmsCampaignStarted struct {
	CampaignId CampaignId `json:"campaignId"`
}

type SmsCampaignStopped struct {
	CampaignId CampaignId          `json:"campaignId"`
	Reason     *CampaignStopReason `json:"reason,omitempty"`
}

type SmsCampaignCompleted struct {
	CampaignId CampaignId `json:"campaignId"`
	Errors     []ErrorDto `json:"errors,omitempty"`
}

type SmsCampaignFailed struct {
	CampaignId CampaignId `json:"campaignId"`
	Errors     []ErrorDto `json:"errors"`
}

type SmsCampaignTriggered struct {
	ProjectId     ProjectId            `json:"projectId"`
	TriggerId     TriggerId            `json:"triggerId"`
	TriggerType   TriggerType          `json:"triggerType,omitempty"`
	SourceEvent   string               `json:"sourceEvent"`
	SchemaId      *string              `json:"schemaId,omitempty"`
	TokenMappings *IReadOnlyDictionary `json:"tokenMappings,omitempty"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/secrets", "PUT")
// @DataContract
type ReplaceMarketplaceIntegrationSecretsRequest struct {
	CodeMashRequestBase
	/** @description Integration view id (int_…). */
	// @DataMember
	// @ApiMember(Description="Integration view id (int_…).", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	// @DataMember
	Secrets map[string]string `json:"secrets"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/secrets/reveal", "POST")
// @DataContract
type RevealMarketplaceIntegrationSecretsRequest struct {
	CodeMashRequestBase
	/** @description Integration view id (int_…). */
	// @DataMember
	// @ApiMember(Description="Integration view id (int_…).", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/token-mappings", "PUT")
// @DataContract
type SetMarketplaceIntegrationTokenMappingsRequest struct {
	CodeMashRequestBase
	/** @description Integration view id (int_…). */
	// @DataMember
	// @ApiMember(Description="Integration view id (int_…).", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
	// @DataMember
	TokenMappings []MarketplaceTokenMappingDto `json:"tokenMappings"`
}

// @Route("/{version}/code/marketplace/integrations/{IntegrationViewId}/catalog", "GET")
type GetMarketplaceFunctionCatalog struct {
	CodeMashRequestBase
	/** @description Integration view id, from get_marketplace_integrations. */
	// @ApiMember(Description="Integration view id, from get_marketplace_integrations.", IsRequired=true)
	IntegrationViewId string `json:"integrationViewId"`
}

type CodeIntegrationSaved struct {
	Integration CodeIntegration `json:"integration"`
}

type CodeIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type CodeIntegrationHumanDeliveryConfirmed struct {
	Id             IntegrationId `json:"id"`
	ConfirmedAtUtc time.Time     `json:"confirmedAtUtc,omitempty"`
}

type CodeIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type CodeIntegrationSetAsDefault struct {
	Id IntegrationId `json:"id"`
}

type CodeIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type CodeIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type CodeIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MarketplaceIntegrationSaved struct {
	Integration MarketplaceIntegration `json:"integration"`
}

type MarketplaceIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MarketplaceIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MarketplaceIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MarketplaceIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type MarketplaceIntegrationSecretsConfigured struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MarketplaceIntegrationSecretsConfigurationFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type MarketplaceFunctionSaved struct {
	Function MarketplaceFunction `json:"function"`
}

type MarketplaceFunctionDeleted struct {
	IntegrationId IntegrationId         `json:"integrationId"`
	FunctionId    MarketplaceFunctionId `json:"functionId"`
}

type MarketplaceFunctionEnabled struct {
	IntegrationId IntegrationId         `json:"integrationId"`
	FunctionId    MarketplaceFunctionId `json:"functionId"`
}

type MarketplaceFunctionDisabled struct {
	IntegrationId IntegrationId         `json:"integrationId"`
	FunctionId    MarketplaceFunctionId `json:"functionId"`
}

type ServerlessEnabled struct {
}

type ServerlessDisabled struct {
}

/** @description Disable push service */
// @Route("/{version}/notifications/push/disable", "GET")
// @Api(Description="Disable push service")
type DisablePush struct {
	CodeMashRequestBase
}

/** @description Lists push disable dependencies */
// @Route("/{version}/notifications/push/disable-dependencies", "GET")
// @Api(Description="Lists push disable dependencies")
type GetPushDisableDependencies struct {
	CodeMashRequestBase
}

/** @description Enable push service */
// @Route("/{version}/notifications/push/enable", "GET")
// @Api(Description="Enable push service")
type EnablePush struct {
	CodeMashRequestBase
}

/** @description Archives push template */
// @Route("/{version}/notifications/push/templates/{Id}/archive", "PUT")
// @Api(Description="Archives push template")
type ArchivePushTemplateRequest struct {
	CodeMashRequestBase
	/** @description The push template id to archive. Get it from get_push_templates. */
	// @ApiMember(Description="The push template id to archive. Get it from get_push_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Clones push template */
// @Route("/{version}/notifications/push/templates/{Id}/clone", "POST")
// @Api(Description="Clones push template")
type ClonePushTemplateRequest struct {
	CodeMashRequestBase
	/** @description The push template id to clone. Get it from get_push_templates. */
	// @ApiMember(Description="The push template id to clone. Get it from get_push_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Create push template */
// @Route("/{version}/notifications/push/templates", "POST")
// @Api(Description="Create push template")
type CreatePushTemplateRequest struct {
	SavePushTemplate
}

/** @description Delete push template */
// @Route("/{version}/notifications/push/templates/{Id}", "DELETE")
// @Api(Description="Delete push template")
type DeletePushTemplateRequest struct {
	CodeMashRequestBase
	/** @description The push template id to delete. Get it from get_push_templates. */
	// @ApiMember(Description="The push template id to delete. Get it from get_push_templates.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets a push template */
// @Route("/{version}/notifications/push/templates/{id}", "GET")
// @Api(Description="Gets a push template")
type GetPushTemplate struct {
	CodeMashRequestBase
	/** @description The push template id to fetch. Get it from get_push_templates. */
	// @ApiMember(Description="The push template id to fetch. Get it from get_push_templates.")
	Id string `json:"id"`
}

/** @description Gets push templates */
// @Route("/{version}/notifications/push/templates", "GET")
// @Api(Description="Gets push templates")
type GetPushTemplates struct {
	CodeMashListPaginationRequestBase
	/** @description Set true to include archived templates. */
	// @ApiMember(Description="Set true to include archived templates.")
	ShowArchived *bool `json:"showArchived,omitempty"`
	/** @description Optional: return only the template with this id. */
	// @ApiMember(Description="Optional: return only the template with this id.")
	TemplateId *string `json:"templateId,omitempty"`
}

/** @description Gets push template content tokens */
// @Route("/{version}/notifications/push/templates/{id}/tokens", "GET")
// @Api(Description="Gets push template content tokens")
type GetPushMessageContentTokens struct {
	CodeMashRequestBase
	/** @description The push template id to scan for tokens. Get it from get_push_templates. */
	// @ApiMember(Description="The push template id to scan for tokens. Get it from get_push_templates.")
	Id string `json:"id"`
}

/** @description Renders a push template field */
// @Route("/{version}/notifications/push/templates/render", "POST")
// @Api(Description="Renders a push template field")
type RenderPush struct {
	CodeMashRequestBase
	/** @description The Razor template source for the field being rendered (Title, Body, or Subtitle). */
	// @ApiMember(Description="The Razor template source for the field being rendered (Title, Body, or Subtitle).", IsRequired=true)
	Code string `json:"code"`
	/** @description Optional token values already bound for this render pass. */
	// @ApiMember(Description="Optional token values already bound for this render pass.")
	Tokens []TokenMappingDto `json:"tokens,omitempty"`
	/** @description Set true when rendering for a preview (relaxes strict validation). */
	// @ApiMember(Description="Set true when rendering for a preview (relaxes strict validation).")
	IsForPreview bool `json:"isForPreview,omitempty"`
}

/** @description Un-archives push template */
// @Route("/{version}/notifications/push/templates/{Id}/unarchive", "PUT")
// @Api(Description="Un-archives push template")
type UnArchivePushTemplateRequest struct {
	CodeMashRequestBase
	/** @description The push template id to unarchive. */
	// @ApiMember(Description="The push template id to unarchive.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Edit push template */
// @Route("/{version}/notifications/push/templates", "PUT")
// @Api(Description="Edit push template")
type UpdatePushTemplateRequest struct {
	SavePushTemplate
	/** @description The push template id to update. Get it from get_push_templates. */
	// @ApiMember(Description="The push template id to update. Get it from get_push_templates.", IsRequired=true)
	ViewId string `json:"viewId"`
}

/** @description Gets push settings */
// @Route("/{version}/notifications/push/settings", "GET")
// @Api(Description="Gets push settings")
type GetPushSettings struct {
	CodeMashRequestBase
	/** @description The push settings id to fetch. */
	// @ApiMember(Description="The push settings id to fetch.")
	Id string `json:"id"`
}

/** @description Confirm human delivery of a test push */
// @Route("/{version}/notifications/push/integrations/confirm-human-delivery", "POST")
// @Api(Description="Confirm human delivery of a test push")
// @DataContract
type ConfirmPushIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	/** @description The push integration id being verified. Get it from get_push_integrations. */
	// @DataMember
	// @ApiMember(Description="The push integration id being verified. Get it from get_push_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

/** @description Delete push integration */
// @Route("/{version}/notifications/push/integrations/{Id}", "DELETE")
// @Api(Description="Delete push integration")
type DeletePushIntegrationRequest struct {
	CodeMashRequestBase
	/** @description The push integration id to delete. Get it from get_push_integrations. */
	// @ApiMember(Description="The push integration id to delete. Get it from get_push_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Disable push integration */
// @Route("/{version}/notifications/push/integrations/{Id}/disable", "PUT")
// @Api(Description="Disable push integration")
type DisablePushIntegrationRequest struct {
	CodeMashRequestBase
	/** @description The push integration id to disable. Get it from get_push_integrations. */
	// @ApiMember(Description="The push integration id to disable. Get it from get_push_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Enable push integration */
// @Route("/{version}/notifications/push/integrations/{Id}/enable", "PUT")
// @Api(Description="Enable push integration")
type EnablePushIntegrationRequest struct {
	CodeMashRequestBase
	/** @description The push integration id to enable. Get it from get_push_integrations. */
	// @ApiMember(Description="The push integration id to enable. Get it from get_push_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets a push integration */
// @Route("/{version}/notifications/push/integrations/{id}", "GET")
// @Api(Description="Gets a push integration")
type GetPushIntegration struct {
	CodeMashRequestBase
	/** @description The push integration id to fetch. Get it from get_push_integrations. */
	// @ApiMember(Description="The push integration id to fetch. Get it from get_push_integrations.")
	Id string `json:"id"`
}

/** @description Gets push integrations */
// @Route("/{version}/notifications/push/integrations", "GET")
// @Api(Description="Gets push integrations")
type GetPushIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/notifications/push/integrations", "POST")
// @DataContract
type SavePushIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration PushIntegrationRequest `json:"integration"`
}

/** @description Sets push integration as default */
// @Route("/{version}/notifications/push/integrations/{Id}/default", "PUT")
// @Api(Description="Sets push integration as default")
type SetPushIntegrationAsDefaultRequest struct {
	CodeMashRequestBase
	/** @description The push integration id to set as default. Get it from get_push_integrations. */
	// @ApiMember(Description="The push integration id to set as default. Get it from get_push_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Test push integration */
// @Route("/{version}/notifications/push/integrations/test", "POST")
// @Api(Description="Test push integration")
type TestPushIntegration struct {
	CodeMashRequestBase
	/** @description The push integration id to test. Get it from get_push_integrations. */
	// @DataMember
	// @ApiMember(Description="The push integration id to test. Get it from get_push_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
	/** @description Optional device token to send the test notification to. Requires DeliveryFamily when set. */
	// @DataMember
	// @ApiMember(Description="Optional device token to send the test notification to. Requires DeliveryFamily when set.")
	TestToken *string `json:"testToken,omitempty"`
	/** @description Optional delivery family for the test token (e.g. Ios, Android, Chrome, Safari, Expo). Requires TestToken when set. */
	// @DataMember
	// @ApiMember(Description="Optional delivery family for the test token (e.g. Ios, Android, Chrome, Safari, Expo). Requires TestToken when set.")
	DeliveryFamily *string `json:"deliveryFamily,omitempty"`
}

// @Route("/{version}/notifications/push/integrations/app/request", "POST")
type RegisterCodeMashAppPushIntegration struct {
	CodeMashRequestBase
	AccountId string    `json:"accountId"`
	UserId    string    `json:"userId"`
	RequestId string    `json:"requestId"`
	Pin       int       `json:"pin,omitempty"`
	ValidTill time.Time `json:"validTill,omitempty"`
	PublicKey string    `json:"publicKey"`
}

/** @description Registers a device for push notifications */
// @Route("/{version}/notifications/push/devices", "POST")
// @Api(Description="Registers a device for push notifications")
// @DataContract
type RegisterDevice struct {
	RequestBase
	/** @description The device details: OS, token, model, and delivery family. */
	// @DataMember
	// @ApiMember(Description="The device details: OS, token, model, and delivery family.", IsRequired=true)
	PushDeviceDto PushDeviceDto `json:"pushDeviceDto"`
	/** @description The id of the user this device belongs to. */
	// @DataMember
	// @ApiMember(Description="The id of the user this device belongs to.", IsRequired=true)
	UserId string `json:"userId"`
	// @DataMember
	ProjectId string `json:"projectId"`
	/** @description Optional account id to associate with the device. */
	// @DataMember
	// @ApiMember(Description="Optional account id to associate with the device.")
	AccountId *string `json:"accountId,omitempty"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @DataMember
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

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
}

/** @description Deletes push campaign from queue */
// @Route("/{version}/notifications/push/campaigns/{Id}", "DELETE")
// @Api(Description="Deletes push campaign from queue")
// @DataContract
type DeletePushCampaignRequest struct {
	CodeMashRequestBase
}

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

/** @description Returns push preview notification */
// @Route("/{version}/notifications/push/preview", "GET")
// @Api(Description="Returns push preview notification")
type PreviewPushNotification struct {
	RequestBase
	/** @description The encrypted preview hash identifying the project and notification. */
	// @ApiMember(Description="The encrypted preview hash identifying the project and notification.")
	Hash string `json:"hash"`
}

/** @description Stops a running push campaign */
// @Route("/{version}/notifications/push/campaigns/{Id}/stop", "POST")
// @Api(Description="Stops a running push campaign")
// @DataContract
type StopPushCampaignRequest struct {
	CodeMashRequestBase
}

/** @description Gets campaign push notification details */
// @Route("/{version}/notifications/push/campaigns/{campaignId}/messages/{id}", "GET")
// @Api(Description="Gets campaign push notification details")
type GetPushCampaignMessage struct {
	CodeMashRequestBase
	/** @description The push campaign id. Get it from get_push_campaigns. */
	// @ApiMember(Description="The push campaign id. Get it from get_push_campaigns.")
	CampaignId string `json:"campaignId"`
	/** @description The batch id. Get it from get_push_campaign_batches. */
	// @ApiMember(Description="The batch id. Get it from get_push_campaign_batches.")
	CampaignBatchId string `json:"campaignBatchId"`
	/** @description The notification id within the batch. */
	// @ApiMember(Description="The notification id within the batch.")
	NotificationId string `json:"notificationId"`
	/** @description Optional database integration id; omit to use the project's default. */
	// @ApiMember(Description="Optional database integration id; omit to use the project's default.")
	DatabaseIntegrationId *string `json:"databaseIntegrationId,omitempty"`
}

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

type PushIntegrationSaved struct {
	Integration PushIntegration `json:"integration"`
}

type PushIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type PushIntegrationHumanDeliveryConfirmed struct {
	Id             IntegrationId `json:"id"`
	ConfirmedAtUtc time.Time     `json:"confirmedAtUtc,omitempty"`
}

type PushIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type PushIntegrationSetAsDefault struct {
	Env Env           `json:"env"`
	Id  IntegrationId `json:"id"`
}

type PushIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PushIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PushIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PushServiceEstablished struct {
	DefaultTemplates []PushTemplate `json:"defaultTemplates,omitempty"`
}

type PushServiceEnabled struct {
}

type PushServiceDisabled struct {
}

type PushModuleTagSaved struct {
	Tag                  TagDefinition        `json:"tag"`
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
}

type PushModuleTagDeleted struct {
	Tag                  Tag                  `json:"tag"`
	CommunicationChannel CommunicationChannel `json:"communicationChannel,omitempty"`
}

type PushTemplateCreated struct {
	TemplateId   TemplateId           `json:"templateId"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Channel      CommunicationChannel `json:"channel,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Tags         []Tag                `json:"tags,omitempty"`
	Env          *Env                 `json:"env,omitempty"`
}

type PushTemplateUpdated struct {
	TemplateId   TemplateId           `json:"templateId"`
	DisplayName  DisplayName          `json:"displayName"`
	Translations []MessageTranslation `json:"translations"`
	Channel      CommunicationChannel `json:"channel,omitempty"`
	Description  *string              `json:"description,omitempty"`
	Tags         []Tag                `json:"tags,omitempty"`
	Env          *Env                 `json:"env,omitempty"`
}

type PushTemplateMirrored struct {
	Template PushTemplate `json:"template"`
}

type PushTemplateDeleted struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type PushTemplateArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

type PushTemplateUnArchived struct {
	TemplateId TemplateId `json:"templateId"`
	Env        *Env       `json:"env,omitempty"`
}

/** @description Disable payments service */
// @Route("/{version}/payments/disable", "GET")
// @Api(Description="Disable payments service")
type DisablePayments struct {
	CodeMashRequestBase
}

/** @description Enable payments service */
// @Route("/{version}/payments/enable", "GET")
// @Api(Description="Enable payments service")
type EnablePayments struct {
	CodeMashRequestBase
}

/** @description Gets the received payment webhooks log */
// @Route("/{version}/payments/webhooks/log", "GET")
// @Api(Description="Gets the received payment webhooks log")
type GetPaymentsWebhookLog struct {
	CodeMashRequestBase
	/** @description Only rows for this payments integration (view id). Omit for the whole project. */
	// @ApiMember(DataType="string", Description="Only rows for this payments integration (view id). Omit for the whole project.")
	IntegrationId *string `json:"integrationId,omitempty"`
	/** @description Max rows to return, newest first. Default 50, ceiling 200. */
	// @ApiMember(DataType="int", Description="Max rows to return, newest first. Default 50, ceiling 200.")
	Limit *int `json:"limit,omitempty"`
}

/** @description Delete payments trigger */
// @Route("/{version}/payments/triggers/{triggerId}", "DELETE")
// @Api(Description="Delete payments trigger")
// @DataContract
type DeletePaymentsTrigger struct {
	DeleteTrigger
}

/** @description Disable payments trigger */
// @Route("/{version}/payments/triggers/{triggerId}/disable", "PATCH")
// @Api(Description="Disable payments trigger")
// @DataContract
type DisablePaymentsTrigger struct {
	DisableTrigger
}

/** @description Enable payments trigger */
// @Route("/{version}/payments/triggers/{triggerId}/enable", "PATCH")
// @Api(Description="Enable payments trigger")
// @DataContract
type EnablePaymentsTrigger struct {
	EnableTrigger
}

/** @description Gets payments trigger by specified Id */
// @Route("/{version}/payments/triggers/{id}", "GET")
// @Api(Description="Gets payments trigger by specified Id")
type GetPaymentsTrigger struct {
	GetTrigger
}

/** @description Gets payments triggers */
// @Route("/{version}/payments/triggers", "GET")
// @Api(Description="Gets payments triggers")
type GetPaymentsTriggers struct {
	GetTriggers
}

// @Route("/{version}/payments/triggers", "POST")
// @DataContract
type SavePaymentsTrigger struct {
	SaveTrigger
}

/** @description Confirm that you received or verified the test payment integration outcome. */
// @Route("/{version}/payments/integrations/confirm-human-delivery", "POST")
// @Api(Description="Confirm that you received or verified the test payment integration outcome.")
// @DataContract
type ConfirmPaymentsIntegrationHumanDeliveryRequest struct {
	CodeMashRequestBase
	/** @description The id of the payments integration whose test outcome is being confirmed. */
	// @DataMember
	// @ApiMember(Description="The id of the payments integration whose test outcome is being confirmed.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

/** @description Delete integration for particular project */
// @Route("/{version}/payments/integrations/{Id}", "DELETE")
// @Api(Description="Delete integration for particular project")
type DeletePaymentsIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Payments integration id to delete, from get_payments_integrations. */
	// @ApiMember(Description="Payments integration id to delete, from get_payments_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Disable integration for particular project */
// @Route("/{version}/payments/integrations/{Id}/disable", "PUT")
// @Api(Description="Disable integration for particular project")
type DisablePaymentsIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Payments integration id to disable, from get_payments_integrations. */
	// @ApiMember(Description="Payments integration id to disable, from get_payments_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Enable integration for particular project */
// @Route("/{version}/payments/integrations/{Id}/enable", "PUT")
// @Api(Description="Enable integration for particular project")
type EnablePaymentsIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Payments integration id to enable, from get_payments_integrations. */
	// @ApiMember(Description="Payments integration id to enable, from get_payments_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integration by specified Id */
// @Route("/{version}/payments/integrations/{id}", "GET")
// @Api(Description="Gets integration by specified Id")
type GetPaymentsIntegration struct {
	CodeMashRequestBase
	/** @description Payments integration id to fetch, from get_payments_integrations. */
	// @ApiMember(Description="Payments integration id to fetch, from get_payments_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integrations */
// @Route("/{version}/payments/integrations", "GET")
// @Api(Description="Gets integrations")
type GetPaymentsIntegrations struct {
	CodeMashListPaginationRequestBase
}

/** @description Saves payments integration */
// @Route("/{version}/payments/integrations", "POST")
// @Api(Description="Saves payments integration")
// @DataContract
type SavePaymentsIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration PaymentIntegrationRequest `json:"integration"`
}

/** @description Test payments integration */
// @Route("/{version}/payments/integrations/test", "POST")
// @Api(Description="Test payments integration")
type TestPaymentsIntegration struct {
	CodeMashRequestBase
	/** @description The id of the payments integration to test. */
	// @ApiMember(Description="The id of the payments integration to test.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

type PaymentsIntegrationSaved struct {
	Integration PaymentIntegration `json:"integration"`
}

type PaymentsIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type PaymentsIntegrationHumanDeliveryConfirmed struct {
	Id             IntegrationId `json:"id"`
	ConfirmedAtUtc time.Time     `json:"confirmedAtUtc,omitempty"`
}

type PaymentsIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type PaymentsIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PaymentsIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PaymentsIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type PaymentsEstablished struct {
}

type PaymentsEnabled struct {
}

type PaymentsDisabled struct {
}

type PaymentsTriggerSaved struct {
	Trigger PaymentTrigger `json:"trigger"`
}

type PaymentTriggerMirrored struct {
	Trigger Trigger `json:"trigger"`
}

type PaymentsTriggerEnabled struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type PaymentsTriggerDisabled struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

type PaymentsTriggerDeleted struct {
	TriggerByIdEventBase
	Env Env `json:"env"`
}

/** @description Disable logging service */
// @Route("/{version}/logs/disable", "GET")
// @Api(Description="Disable logging service")
type DisableLogging struct {
	CodeMashRequestBase
}

// @Route("/{version}/logs/enable", "GET")
type EnableLogging struct {
	CodeMashRequestBase
	/** @description When true, also create a Norbix Logging integration backed by the project's default database. */
	// @ApiMember(DataType="boolean", Description="When true, also create a Norbix Logging integration backed by the project's default database.", Name="createNorbixLogging", ParameterType="query")
	CreateNorbixLogging bool `json:"createNorbixLogging,omitempty"`
}

/** @description Delete integration for particular project */
// @Route("/{version}/logs/integrations/{Id}", "DELETE")
// @Api(Description="Delete integration for particular project")
type DeleteLoggingIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Logging integration id to delete, from get_logging_integrations. */
	// @ApiMember(Description="Logging integration id to delete, from get_logging_integrations.", IsRequired=true)
	Id string `json:"id"`
	/** @description When true and this is a Norbix Logging integration, also permanently wipes the stored log entries in its backing database. Ignored for other providers. */
	// @ApiMember(Description="When true and this is a Norbix Logging integration, also permanently wipes the stored log entries in its backing database. Ignored for other providers.")
	WipeLogs bool `json:"wipeLogs,omitempty"`
}

/** @description Disable integration for particular project */
// @Route("/{version}/logs/integrations/{Id}/disable", "PUT")
// @Api(Description="Disable integration for particular project")
type DisableLoggingIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Logging integration id to disable, from get_logging_integrations. */
	// @ApiMember(Description="Logging integration id to disable, from get_logging_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Enable integration for particular project */
// @Route("/{version}/logs/integrations/{Id}/enable", "PUT")
// @Api(Description="Enable integration for particular project")
type EnableLoggingIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Logging integration id to enable, from get_logging_integrations. */
	// @ApiMember(Description="Logging integration id to enable, from get_logging_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integration by specified Id */
// @Route("/{version}/logs/integrations/{id}", "GET")
// @Api(Description="Gets integration by specified Id")
type GetLoggingIntegration struct {
	CodeMashRequestBase
	/** @description Logging integration id to fetch, from get_logging_integrations. */
	// @ApiMember(Description="Logging integration id to fetch, from get_logging_integrations.", IsRequired=true)
	Id string `json:"id"`
}

/** @description Gets integrations */
// @Route("/{version}/logs/integrations", "GET")
// @Api(Description="Gets integrations")
type GetLoggingIntegrations struct {
	CodeMashListPaginationRequestBase
}

/** @description Saves logging integration */
// @Route("/{version}/logs/integrations", "POST")
// @Api(Description="Saves logging integration")
// @DataContract
type SaveLoggingIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration LoggingIntegrationRequest `json:"integration"`
}

/** @description Test logging integration */
// @Route("/{version}/logs/integrations/test", "POST")
// @Api(Description="Test logging integration")
type TestLoggingIntegration struct {
	CodeMashRequestBase
	/** @description Logging integration id to test, from get_logging_integrations. */
	// @ApiMember(Description="Logging integration id to test, from get_logging_integrations.", IsRequired=true)
	IntegrationId string `json:"integrationId"`
}

/** @description Delete every log entry stored in the project's Norbix Logging integration */
// @Route("/{version}/logs/clean", "POST")
// @Api(Description="Delete every log entry stored in the project's Norbix Logging integration")
type CleanLogs struct {
	CodeMashRequestBase
}

/** @description Fetch the audit trail for a correlation id */
// @Route("/{version}/logs/audit", "GET")
// @Api(Description="Fetch the audit trail for a correlation id")
type GetLogsByCorrelationId struct {
	CodeMashRequestBase
	/** @description The correlation id whose full request trail you want. */
	// @ApiMember(DataType="string", Description="The correlation id whose full request trail you want.", IsRequired=true, Name="correlationId", ParameterType="query")
	TargetCorrelationId string `json:"targetCorrelationId"`
}

/** @description Fetch a filtered, cursor-paged list of tenant log entries */
// @Route("/{version}/logs", "GET")
// @Api(Description="Fetch a filtered, cursor-paged list of tenant log entries")
type GetLogs struct {
	CodeMashListPaginationRequestBase
	/** @description Severity filter: Information, Warning or Error. */
	// @ApiMember(DataType="string", Description="Severity filter: Information, Warning or Error.", Name="level", ParameterType="query")
	Level *string `json:"level,omitempty"`
	/** @description Module filter: Database, Email, Membership, etc. */
	// @ApiMember(DataType="string", Description="Module filter: Database, Email, Membership, etc.", Name="module", ParameterType="query")
	Module *string `json:"module,omitempty"`
	/** @description Correlation id filter. Empty = no filter (show all). */
	// @ApiMember(DataType="string", Description="Correlation id filter. Empty = no filter (show all).", Name="logCorrelationId", ParameterType="query")
	LogCorrelationId *string `json:"logCorrelationId,omitempty"`
	/** @description Exact event code filter (e.g. db:record:insert). */
	// @ApiMember(DataType="string", Description="Exact event code filter (e.g. db:record:insert).", Name="eventCode", ParameterType="query")
	EventCode *string `json:"eventCode,omitempty"`
	/** @description Free-text search over title and message. */
	// @ApiMember(DataType="string", Description="Free-text search over title and message.", Name="search", ParameterType="query")
	Search *string `json:"search,omitempty"`
	/** @description Start of the timestamp range (inclusive, UTC). Optional. */
	// @ApiMember(Description="Start of the timestamp range (inclusive, UTC). Optional.")
	FromUtc *time.Time `json:"fromUtc,omitempty"`
	/** @description End of the timestamp range (inclusive, UTC). Optional. */
	// @ApiMember(Description="End of the timestamp range (inclusive, UTC). Optional.")
	ToUtc *time.Time `json:"toUtc,omitempty"`
}

/** @description Fetch the per-project log settings (flags) */
// @Route("/{version}/logs/settings", "GET")
// @Api(Description="Fetch the per-project log settings (flags)")
type GetLogSettings struct {
	CodeMashRequestBase
}

/** @description Update the per-project log settings (flags) */
// @Route("/{version}/logs/settings", "POST")
// @Api(Description="Update the per-project log settings (flags)")
type SaveLogSettings struct {
	CodeMashRequestBase
	/** @description Drop log entries from requests originating from the Norbix studio (cloud dashboard). */
	// @ApiMember(Description="Drop log entries from requests originating from the Norbix studio (cloud dashboard).")
	SkipCloudDashboardLogs bool `json:"skipCloudDashboardLogs,omitempty"`
	/** @description Strip request/response body meta off http:request / http:response log entries. */
	// @ApiMember(Description="Strip request/response body meta off http:request / http:response log entries.")
	SkipHttpBodyMeta bool `json:"skipHttpBodyMeta,omitempty"`
	/** @description Turn on tenant-visible log entries for AI chat turns. Default false. */
	// @ApiMember(Description="Turn on tenant-visible log entries for AI chat turns. Default false.")
	AiChatLoggingEnabled bool `json:"aiChatLoggingEnabled,omitempty"`
}

type LoggingIntegrationSaved struct {
	Integration LoggingIntegration `json:"integration"`
}

type LoggingIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type LoggingIntegrationRenamed struct {
	Id   IntegrationId `json:"id"`
	Name DisplayName   `json:"name"`
	Env  *Env          `json:"env,omitempty"`
}

type LoggingIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationSecretsConfigured struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationSecretsConfigurationFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationSecretsCleared struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationSecretsClearingFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LoggingIntegrationSetAsDefault struct {
	Id IntegrationId `json:"id"`
}

type NorbixLoggingLogsWipeRequested struct {
	DeletedIntegrationId  IntegrationId `json:"deletedIntegrationId"`
	DatabaseIntegrationId IntegrationId `json:"databaseIntegrationId"`
}

type LoggingEstablished struct {
}

type LoggingEnabled struct {
}

type LoggingDisabled struct {
}

/** @description Lists the AI tools this host exposes (external-agent bridge). */
// @Route("/{version}/account/ai/tools", "GET")
// @Api(Description="Lists the AI tools this host exposes (external-agent bridge).")
type GetAiToolsRequest struct {
	RequestBase
	Toolset *string `json:"toolset,omitempty"`
}

/** @description Invokes one AI tool directly (external-agent bridge). */
// @Route("/{version}/account/ai/tools/{ToolName}", "POST")
// @Api(Description="Invokes one AI tool directly (external-agent bridge).")
type InvokeAiToolRequest struct {
	RequestBase
	ToolName      string  `json:"toolName"`
	ArgumentsJson *string `json:"argumentsJson,omitempty"`
}

/** @description Gets account info. */
// @Route("/{version}/account/chat/complete", "POST")
// @Api(Description="Gets account info.")
type AskChatRequest struct {
	RequestBase
	Prompt  string  `json:"prompt"`
	Profile *string `json:"profile,omitempty"`
}

/** @description Uploads a file into an AI chat session. */
// @Route("/{version}/account/chat/attachments", "POST")
// @Api(Description="Uploads a file into an AI chat session.")
type UploadChatAttachmentRequest struct {
	RequestBase
	SessionId     *string `json:"sessionId,omitempty"`
	FileName      string  `json:"fileName"`
	ContentType   string  `json:"contentType"`
	Base64Content string  `json:"base64Content"`
	Profile       *string `json:"profile,omitempty"`
	Topic         *string `json:"topic,omitempty"`
	ProjectId     *string `json:"projectId,omitempty"`
	Env           *string `json:"env,omitempty"`
}

/** @description Reports AI chat availability and the model-picker menu. */
// @Route("/{version}/account/chat/availability", "GET")
// @Api(Description="Reports AI chat availability and the model-picker menu.")
type ChatAvailabilityRequest struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
	Env       *string `json:"env,omitempty"`
}

/** @description Lists what the AI assistant remembers about this account. */
// @Route("/{version}/account/chat/memory", "GET")
// @Api(Description="Lists what the AI assistant remembers about this account.")
type GetChatMemoryRequest struct {
	RequestBase
	ProjectId *string `json:"projectId,omitempty"`
}

/** @description Deletes one AI memory note ('forget this'). */
// @Route("/{version}/account/chat/memory/{NoteId}", "DELETE")
// @Api(Description="Deletes one AI memory note ('forget this').")
type ForgetChatMemoryRequest struct {
	RequestBase
	NoteId string `json:"noteId"`
}

/** @description Deletes an AI chat session (soft delete). */
// @Route("/{version}/account/chat/sessions/{SessionId}", "DELETE")
// @Api(Description="Deletes an AI chat session (soft delete).")
type DeleteChatSessionRequest struct {
	RequestBase
	SessionId string `json:"sessionId"`
}

/** @description Archives or unarchives an AI chat session. */
// @Route("/{version}/account/chat/sessions/{SessionId}/archive", "PATCH")
// @Api(Description="Archives or unarchives an AI chat session.")
type SetChatSessionArchivedRequest struct {
	RequestBase
	SessionId string `json:"sessionId"`
	Archived  bool   `json:"archived,omitempty"`
}

/** @description Pins or unpins an AI chat session. */
// @Route("/{version}/account/chat/sessions/{SessionId}/pin", "PATCH")
// @Api(Description="Pins or unpins an AI chat session.")
type SetChatSessionPinnedRequest struct {
	RequestBase
	SessionId string `json:"sessionId"`
	Pinned    bool   `json:"pinned,omitempty"`
}

/** @description Marks or unmarks an AI chat session as "do not share". */
// @Route("/{version}/account/chat/sessions/{SessionId}/sharing", "PATCH")
// @Api(Description="Marks or unmarks an AI chat session as \"do not share\".")
type SetChatSessionSharingRequest struct {
	RequestBase
	SessionId  string `json:"sessionId"`
	DoNotShare bool   `json:"doNotShare,omitempty"`
}

/** @description Lists the account's recent AI chat sessions. */
// @Route("/{version}/account/chat/sessions", "GET")
// @Api(Description="Lists the account's recent AI chat sessions.")
type GetChatSessionsRequest struct {
	RequestBase
	Take            *int  `json:"take,omitempty"`
	IncludeArchived *bool `json:"includeArchived,omitempty"`
}

/** @description Returns one AI chat session's conversation entries — the transcript. */
// @Route("/{version}/account/chat/sessions/{SessionId}/entries", "GET")
// @Api(Description="Returns one AI chat session's conversation entries — the transcript.")
type GetChatSessionEntriesRequest struct {
	RequestBase
	SessionId string `json:"sessionId"`
	SinceSeq  *int64 `json:"sinceSeq,omitempty"`
}

/** @description Records like / dislike feedback on one AI chat entry, or clears it. */
// @Route("/{version}/account/chat/sessions/{SessionId}/entries/{EntryId}/feedback", "POST")
// @Api(Description="Records like / dislike feedback on one AI chat entry, or clears it.")
type SetChatEntryFeedbackRequest struct {
	RequestBase
	SessionId string  `json:"sessionId"`
	EntryId   string  `json:"entryId"`
	Feedback  *string `json:"feedback,omitempty"`
}

/** @description Answers one open AI chat question — or a prepared change's Apply / Skip — and continues the conversation. */
// @Route("/{version}/account/chat/sessions/{SessionId}/questions/{EntryId}/answer", "POST")
// @Api(Description="Answers one open AI chat question — or a prepared change's Apply / Skip — and continues the conversation.")
type AnswerChatQuestionRequest struct {
	RequestBase
	SessionId string            `json:"sessionId"`
	EntryId   string            `json:"entryId"`
	Answers   map[string]string `json:"answers,omitempty"`
}

/** @description Approves or rejects a proposed AI chat plan. */
// @Route("/{version}/account/chat/sessions/{SessionId}/plans/{EntryId}/decision", "POST")
// @Api(Description="Approves or rejects a proposed AI chat plan.")
type DecideChatPlanRequest struct {
	RequestBase
	SessionId string  `json:"sessionId"`
	EntryId   string  `json:"entryId"`
	Decision  string  `json:"decision"`
	Comment   *string `json:"comment,omitempty"`
}

/** @description Stops one running step of an AI chat plan run. */
// @Route("/{version}/account/chat/sessions/{SessionId}/steps/{EntryId}/stop", "POST")
// @Api(Description="Stops one running step of an AI chat plan run.")
type StopChatRunStepRequest struct {
	RequestBase
	SessionId string `json:"sessionId"`
	EntryId   string `json:"entryId"`
}

/** @description Runs one AI chat conversation turn. */
// @Route("/{version}/account/chat/turn", "POST")
// @Api(Description="Runs one AI chat conversation turn.")
type ChatTurnRequest struct {
	RequestBase
	SessionId        *string               `json:"sessionId,omitempty"`
	Message          string                `json:"message"`
	Profile          *string               `json:"profile,omitempty"`
	Topic            *string               `json:"topic,omitempty"`
	LlmIntegrationId *string               `json:"llmIntegrationId,omitempty"`
	Model            *string               `json:"model,omitempty"`
	ProjectId        *string               `json:"projectId,omitempty"`
	Env              *string               `json:"env,omitempty"`
	ScreenContext    *ChatScreenContextDto `json:"screenContext,omitempty"`
}

/** @description MCP server endpoint — JSON-RPC 2.0 over HTTP POST exposing the AI tool catalog. */
// @Route("/{version}/account/mcp", "POST")
// @Api(Description="MCP server endpoint — JSON-RPC 2.0 over HTTP POST exposing the AI tool catalog.")
type McpRequest struct {
	Version       *string `json:"version,omitempty"`
	RequestStream []byte  `json:"requestStream"`
}

/** @description Reads a project's AI Brief: the requirements, decisions and assumptions the assistant recorded from conversations, each with the chat turn, user and time it came from. */
// @Route("/{version}/projects/{projectId}/ai/brief", "GET")
// @Api(Description="Reads a project's AI Brief: the requirements, decisions and assumptions the assistant recorded from conversations, each with the chat turn, user and time it came from.")
type GetProjectBriefRequest struct {
	CodeMashRequestBase
	/** @description Return the Brief events after this sequence number as well (0 = all). Omit for the snapshot only. */
	// @ApiMember(Description="Return the Brief events after this sequence number as well (0 = all). Omit for the snapshot only.")
	SinceSeq *int64 `json:"sinceSeq,omitempty"`
}

/** @description Lists a project's AI work items: one serious ask each, with its goal, plans, changes, moved-out items, needs-you list and open questions. */
// @Route("/{version}/projects/{projectId}/ai/work-items", "GET")
// @Api(Description="Lists a project's AI work items: one serious ask each, with its goal, plans, changes, moved-out items, needs-you list and open questions.")
type GetWorkItemsRequest struct {
	CodeMashRequestBase
	/** @description Filter by status: proposed, active, waiting, done, partly-done or dropped. Omit for all. */
	// @ApiMember(Description="Filter by status: proposed, active, waiting, done, partly-done or dropped. Omit for all.")
	Status *string `json:"status,omitempty"`
}

/** @description Reads one AI work item: the six long-task sections and the definition-of-done verdict. */
// @Route("/{version}/projects/{projectId}/ai/work-items/{WorkItemId}", "GET")
// @Api(Description="Reads one AI work item: the six long-task sections and the definition-of-done verdict.")
type GetWorkItemRequest struct {
	CodeMashRequestBase
	WorkItemId string `json:"workItemId"`
}

/** @description Exports one AI work item as markdown in the long-task shape: Goal, Plan, Changes, Rejected / moved out, Needs you, Open questions. */
// @Route("/{version}/projects/{projectId}/ai/work-items/{WorkItemId}/export.md", "GET")
// @Api(Description="Exports one AI work item as markdown in the long-task shape: Goal, Plan, Changes, Rejected / moved out, Needs you, Open questions.")
type ExportWorkItemRequest struct {
	CodeMashRequestBase
	WorkItemId string `json:"workItemId"`
}

/** @description Ticks one manual line of a work item's "Needs you" checklist — the one write a human makes to a work item directly. */
// @Route("/{version}/projects/{projectId}/ai/work-items/{WorkItemId}/needs-you/{Index}/done", "POST")
// @Api(Description="Ticks one manual line of a work item's \"Needs you\" checklist — the one write a human makes to a work item directly.")
type MarkNeedsYouDoneRequest struct {
	CodeMashRequestBase
	WorkItemId string `json:"workItemId"`
	Index      int    `json:"index,omitempty"`
	/** @description Set false to un-tick the line. Default true. */
	// @ApiMember(Description="Set false to un-tick the line. Default true.")
	Done *bool `json:"done,omitempty"`
}

// @Route("/{version}/ai/integrations/llms/{Id}", "DELETE")
type DeleteLlmIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Id of the LLM integration to delete. */
	// @ApiMember(Description="Id of the LLM integration to delete.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/llms/{Id}/disable", "PUT")
type DisableLlmIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Id of the LLM integration to disable. */
	// @ApiMember(Description="Id of the LLM integration to disable.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/llms/{Id}/enable", "PUT")
type EnableLlmIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Id of the LLM integration to enable. */
	// @ApiMember(Description="Id of the LLM integration to enable.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/llms/{id}", "GET")
type GetLlmIntegration struct {
	CodeMashRequestBase
	/** @description Id of the LLM integration to fetch. */
	// @ApiMember(Description="Id of the LLM integration to fetch.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/llms/integrations", "GET")
type GetLlmIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/ai/integrations/llms/", "POST")
// @DataContract
type SaveLlmIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration LlmIntegrationRequest `json:"integration"`
}

// @Route("/{version}/ai/integrations/llms/test", "POST")
type TestLlmIntegration struct {
	CodeMashRequestBase
	/** @description Id of the LLM integration to test. */
	// @ApiMember(Description="Id of the LLM integration to test.")
	IntegrationId string `json:"integrationId"`
}

// @Route("/{version}/ai/integrations/mcp/{Id}", "DELETE")
type DeleteMcpIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Id of the MCP integration to delete. */
	// @ApiMember(Description="Id of the MCP integration to delete.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/mcp/{Id}/disable", "PUT")
type DisableMcpIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Id of the MCP integration to disable. */
	// @ApiMember(Description="Id of the MCP integration to disable.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/mcp/{Id}/enable", "PUT")
type EnableMcpIntegrationRequest struct {
	CodeMashRequestBase
	/** @description Id of the MCP integration to enable. */
	// @ApiMember(Description="Id of the MCP integration to enable.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/mcp/{id}", "GET")
type GetMcpIntegration struct {
	CodeMashRequestBase
	/** @description Id of the MCP integration to fetch. */
	// @ApiMember(Description="Id of the MCP integration to fetch.")
	Id string `json:"id"`
}

// @Route("/{version}/ai/integrations/mcp/integrations", "GET")
type GetMcpIntegrations struct {
	CodeMashListPaginationRequestBase
}

// @Route("/{version}/ai/integrations/mcp/", "POST")
// @DataContract
type SaveMcpIntegration struct {
	CodeMashRequestBase
	// @DataMember(Name="integration")
	Integration McpIntegrationRequest `json:"integration"`
}

// @Route("/{version}/ai/integrations/mcp/test", "POST")
type TestMcpIntegration struct {
	CodeMashRequestBase
	/** @description Id of the MCP integration to test. */
	// @ApiMember(Description="Id of the MCP integration to test.")
	IntegrationId string `json:"integrationId"`
}

type LlmIntegrationSaved struct {
	LlmIntegration LlmIntegration `json:"llmIntegration"`
}

type LlmIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LlmIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LlmIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LlmIntegrationSecretsConfigured struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LlmIntegrationSecretsConfigurationFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type LlmIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type McpIntegrationSaved struct {
	McpIntegration McpIntegration `json:"mcpIntegration"`
}

type McpIntegrationDeleted struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type McpIntegrationEnabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type McpIntegrationDisabled struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type McpIntegrationSecretsConfigured struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type McpIntegrationSecretsConfigurationFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type McpIntegrationTested struct {
	Id            IntegrationId `json:"id"`
	Succeeded     bool          `json:"succeeded,omitempty"`
	ErrorMessages IReadOnlyList `json:"errorMessages"`
	TestedAtUtc   time.Time     `json:"testedAtUtc,omitempty"`
	Env           *Env          `json:"env,omitempty"`
}

type WebhookIntegrationSaved struct {
	Integration WebhookIntegration `json:"integration"`
}

type WebhookIntegrationExtraHeadersChanged struct {
	Id           IntegrationId        `json:"id"`
	ExtraHeaders *IReadOnlyDictionary `json:"extraHeaders,omitempty"`
}

type WebhookIntegrationSecretsConfigured struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type WebhookIntegrationSecretsConfigurationFailed struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type WebhookIntegrationSecretsCleared struct {
	Id  IntegrationId `json:"id"`
	Env *Env          `json:"env,omitempty"`
}

type WebhookDestinationSaved struct {
	IntegrationId IntegrationId      `json:"integrationId"`
	Destination   WebhookDestination `json:"destination"`
}

type WebhookDestinationRemoved struct {
	IntegrationId IntegrationId        `json:"integrationId"`
	DestinationId WebhookDestinationId `json:"destinationId"`
}

type WebhookDestinationEnabled struct {
	IntegrationId IntegrationId        `json:"integrationId"`
	DestinationId WebhookDestinationId `json:"destinationId"`
}

type WebhookDestinationDisabled struct {
	IntegrationId IntegrationId        `json:"integrationId"`
	DestinationId WebhookDestinationId `json:"destinationId"`
}

/** @description Gets the project's webhook integration */
// @Route("/{version}/webhooks/integration", "GET")
// @Api(Description="Gets the project's webhook integration")
type GetWebhookIntegration struct {
	CodeMashRequestBase
}

// @Route("/{version}/webhooks/integration/secret", "GET")
type RevealWebhookIntegrationSecretRequest struct {
	CodeMashRequestBase
}

// @Route("/{version}/webhooks/integration/secret/rotate", "POST")
type RotateWebhookIntegrationSecretRequest struct {
	CodeMashRequestBase
}

// @Route("/{version}/webhooks/integration/extra-headers", "PUT")
type UpdateWebhookIntegrationExtraHeadersRequest struct {
	CodeMashRequestBase
	/** @description The integration-wide static headers to send with every delivery. Pass an empty dictionary to clear all extra headers. */
	// @ApiMember(Description="The integration-wide static headers to send with every delivery. Pass an empty dictionary to clear all extra headers.")
	ExtraHeaders map[string]string `json:"extraHeaders,omitempty"`
}

// @Route("/{version}/webhooks/{source}/{integrationInstanceId}", "POST")
type ReceiveWebhook struct {
	Source                string `json:"source"`
	IntegrationInstanceId string `json:"integrationInstanceId"`
	RequestStream         []byte `json:"requestStream"`
}

// @Route("/{version}/webhooks/destinations/{DestinationId}/disable", "PUT")
type DisableWebhookDestinationRequest struct {
	CodeMashRequestBase
	/** @description The webhook destination id to disable, from get_webhook_integration. */
	// @ApiMember(Description="The webhook destination id to disable, from get_webhook_integration.", IsRequired=true)
	DestinationId string `json:"destinationId"`
}

// @Route("/{version}/webhooks/destinations/{DestinationId}/enable", "PUT")
type EnableWebhookDestinationRequest struct {
	CodeMashRequestBase
	/** @description The webhook destination id to enable, from get_webhook_integration. */
	// @ApiMember(Description="The webhook destination id to enable, from get_webhook_integration.", IsRequired=true)
	DestinationId string `json:"destinationId"`
}

// @Route("/{version}/webhooks/destinations/{DestinationId}", "DELETE")
type RemoveWebhookDestinationRequest struct {
	CodeMashRequestBase
	/** @description The webhook destination id to remove, from get_webhook_integration. */
	// @ApiMember(Description="The webhook destination id to remove, from get_webhook_integration.", IsRequired=true)
	DestinationId string `json:"destinationId"`
}

// @Route("/{version}/webhooks/destinations", "POST")
type SaveWebhookDestinationRequest struct {
	CodeMashRequestBase
	/** @description Existing destination id to overwrite, from get_webhook_integration. Omit to create a new destination. */
	// @ApiMember(Description="Existing destination id to overwrite, from get_webhook_integration. Omit to create a new destination.")
	DestinationId *string `json:"destinationId,omitempty"`
	/** @description Display name for the destination. */
	// @ApiMember(Description="Display name for the destination.", IsRequired=true)
	DestinationName string `json:"destinationName"`
	/** @description The HTTPS endpoint URL that will receive the webhook deliveries. */
	// @ApiMember(Description="The HTTPS endpoint URL that will receive the webhook deliveries.", IsRequired=true)
	EndpointUrl string `json:"endpointUrl"`
	/** @description The event names this destination subscribes to. Empty subscribes to none. */
	// @ApiMember(Description="The event names this destination subscribes to. Empty subscribes to none.")
	SelectedEvents []string `json:"selectedEvents"`
	/** @description Destination-specific static headers sent with every delivery to this destination. These win over the integration-wide extra headers on duplicate keys. */
	// @ApiMember(Description="Destination-specific static headers sent with every delivery to this destination. These win over the integration-wide extra headers on duplicate keys.")
	ExtraHeaders map[string]string `json:"extraHeaders,omitempty"`
	/** @description Whether this destination is enabled for delivery. Defaults to true. */
	// @ApiMember(Description="Whether this destination is enabled for delivery. Defaults to true.")
	IsEnabled bool `json:"isEnabled,omitempty"`
}

// @Route("/{version}/scheduler/disable", "GET")
type DisableScheduler struct {
	CodeMashRequestBase
}

// @Route("/{version}/scheduler/enable", "GET")
type EnableScheduler struct {
	CodeMashRequestBase
}

// @Route("/{version}/scheduler/tasks/{Id}", "DELETE")
type DeleteSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id"`
}

// @Route("/{version}/scheduler/tasks/{Id}/disable", "PUT")
type DisableSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id"`
}

// @Route("/{version}/scheduler/tasks/{Id}/enable", "PUT")
type EnableSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id"`
}

/** @description Gets a scheduled task by id */
// @Route("/{version}/scheduler/tasks/{id}", "GET")
// @Api(Description="Gets a scheduled task by id")
type GetSchedulerTask struct {
	CodeMashRequestBase
	Id string `json:"id"`
}

/** @description Gets scheduled tasks */
// @Route("/{version}/scheduler/tasks", "GET")
// @Api(Description="Gets scheduled tasks")
type GetSchedulerTasks struct {
	CodeMashListPaginationRequestBase
	/** @description Optional filter — only return tasks of this type. */
	// @ApiMember(Description="Optional filter — only return tasks of this type.")
	Type *SchedulerTaskType `json:"type,omitempty"`
	/** @description Optional filter — only return tasks whose enabled state matches this value. */
	// @ApiMember(Description="Optional filter — only return tasks whose enabled state matches this value.")
	Enabled *bool `json:"enabled,omitempty"`
}

/** @description Save scheduled task */
// @Route("/{version}/scheduler/tasks", "POST")
// @Api(Description="Save scheduled task")
// @DataContract
type SaveSchedulerTaskRequest struct {
	CodeMashRequestBase
	// @DataMember
	TaskId *string `json:"taskId,omitempty"`
	// @DataMember
	Name string `json:"name"`
	// @DataMember
	Description *string `json:"description,omitempty"`
	// @DataMember
	Cron string `json:"cron"`
	// @DataMember
	InitiatorUserId string `json:"initiatorUserId"`
	// @DataMember
	IsEnabled bool `json:"isEnabled,omitempty"`
	// @DataMember
	StopOnError bool `json:"stopOnError,omitempty"`
	// @DataMember
	Task SchedulerTaskRequest `json:"task"`
}

type SchedulerEnabled struct {
}

type SchedulerDisabled struct {
}

type SchedulerTaskSaved struct {
	Task SchedulerTask `json:"task"`
}

type SchedulerTaskEnabled struct {
	TaskId TaskId `json:"taskId"`
}

type SchedulerTaskDisabled struct {
	TaskId TaskId `json:"taskId"`
}

type SchedulerTaskDeleted struct {
	TaskId TaskId `json:"taskId"`
}

// @Route("/{version}/resources/resolve", "POST")
type ResolveResources struct {
	CodeMashRequestBase
	Refs IReadOnlyList `json:"refs"`
}

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
}

/** @description Archive a contact */
// @Route("/{version}/membership/users/{contactId}", "DELETE")
// @Api(Description="Archive a contact")
type DeleteContact struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to archive. Get it from get_all_contacts. */
	// @ApiMember(Description="The contact id (ct_…) to archive. Get it from get_all_contacts.", IsRequired=true)
	ContactId string `json:"contactId"`
}

/** @description Get a contact */
// @Route("/{version}/membership/users/{contactId}", "GET")
// @Api(Description="Get a contact")
type GetContact struct {
	CodeMashRequestBase
	/** @description The contact id (ct_…) to fetch. Get it from get_all_contacts. */
	// @ApiMember(Description="The contact id (ct_…) to fetch. Get it from get_all_contacts.", IsRequired=true)
	ContactId string `json:"contactId"`
}

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
}

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

// @Route("/{version}/compliance/settings", "GET")
type GetComplianceSettings struct {
	CodeMashRequestBase
}

// @Route("/{version}/compliance/retention", "DELETE")
type RemoveRetentionWindowRequest struct {
	CodeMashRequestBase
	DataKind string `json:"dataKind"`
}

// @Route("/{version}/compliance/retention", "POST")
type SaveRetentionWindowRequest struct {
	CodeMashRequestBase
	DataKind string `json:"dataKind"`
	Days     int    `json:"days,omitempty"`
	Action   string `json:"action"`
}

// @Route("/{version}/compliance/regimes", "POST")
type AssignRegimeRequest struct {
	CodeMashRequestBase
	Regime string `json:"regime"`
}

// @Route("/{version}/compliance/regimes", "DELETE")
type ClearRegimeRequest struct {
	CodeMashRequestBase
	Regime string `json:"regime"`
}

// @Route("/{version}/compliance/purposes", "POST")
type DefineConsentPurposeRequest struct {
	CodeMashRequestBase
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Channel         string   `json:"channel"`
	MappedTags      []string `json:"mappedTags"`
	RegulatoryBasis []string `json:"regulatoryBasis"`
	Description     *string  `json:"description,omitempty"`
}

// @Route("/{version}/compliance/purposes/deprecate", "POST")
type DeprecateConsentPurposeRequest struct {
	CodeMashRequestBase
	Key string `json:"key"`
}

// @Route("/{version}/compliance/holds", "GET")
type GetLegalHolds struct {
	CodeMashRequestBase
}

// @Route("/{version}/compliance/holds", "POST")
type PlaceLegalHoldRequest struct {
	CodeMashRequestBase
	SubjectKind string `json:"subjectKind"`
	SubjectId   string `json:"subjectId"`
	Reason      string `json:"reason"`
}

// @Route("/{version}/compliance/holds/release", "POST")
type ReleaseLegalHoldRequest struct {
	CodeMashRequestBase
	HoldId string `json:"holdId"`
}

// @Route("/{version}/compliance/dsar/approve", "POST")
type ApproveDsarRequestRequest struct {
	CodeMashRequestBase
	RequestId string `json:"requestId"`
}

// @Route("/{version}/compliance/dsar/reject", "POST")
type RejectDsarRequestRequest struct {
	CodeMashRequestBase
	RequestId string `json:"requestId"`
	Reason    string `json:"reason"`
}

// @Route("/{version}/compliance/dsar", "GET")
type GetDsarRequests struct {
	CodeMashRequestBase
}

// @Route("/{version}/compliance/dsar", "POST")
type OpenDsarRequestRequest struct {
	CodeMashRequestBase
	SubjectKind string `json:"subjectKind"`
	SubjectId   string `json:"subjectId"`
}

// @Route("/{version}/compliance/audit", "GET")
type GetComplianceAuditLog struct {
	CodeMashRequestBase
	From        *time.Time `json:"from,omitempty"`
	To          *time.Time `json:"to,omitempty"`
	SubjectKind *string    `json:"subjectKind,omitempty"`
	SubjectId   *string    `json:"subjectId,omitempty"`
	Limit       int        `json:"limit,omitempty"`
}

// @Route("/{version}/compliance/account", "GET")
type GetAccountCompliance struct {
	RequestBase
}

// @Route("/{version}/compliance/account/dsar-policy", "POST")
type SaveDsarPolicyRequest struct {
	RequestBase
	Mode      string `json:"mode"`
	DelayDays int    `json:"delayDays,omitempty"`
}

// @Route("/{version}/compliance/account/incident-routing", "POST")
type SaveIncidentRoutingRequest struct {
	RequestBase
	AutoForwardAdvisories bool    `json:"autoForwardAdvisories,omitempty"`
	SecurityContact       *string `json:"securityContact,omitempty"`
}

// @Route("/{version}/support/cases/{CaseId}/close", "POST")
type CloseSupportCaseRequest struct {
	RequestBase
	CaseId string `json:"caseId"`
}

// @Route("/{version}/support/cases/{CaseId}/reopen", "POST")
type ReopenSupportCaseRequest struct {
	RequestBase
	CaseId string `json:"caseId"`
	Reason string `json:"reason"`
}

// @Route("/{version}/support/cases/{CaseId}/resolve", "POST")
type ResolveSupportCaseRequest struct {
	RequestBase
	CaseId     string             `json:"caseId"`
	Resolution *CaseResolutionDto `json:"resolution,omitempty"`
}

// @Route("/{version}/support/cases/{CaseId}/messages", "POST")
type AppendSupportCaseMessageRequest struct {
	RequestBase
	CaseId  string `json:"caseId"`
	Message string `json:"message"`
}

// @Route("/{version}/support/cases/{CaseId}", "GET")
type GetSupportCase struct {
	RequestBase
	CaseId string `json:"caseId"`
}

// @Route("/{version}/support/cases", "GET")
type GetSupportCases struct {
	RequestBase
	PagingArgs *PagingArgs `json:"pagingArgs,omitempty"`
}

// @Route("/{version}/support/cases", "POST")
type OpenSupportCaseRequest struct {
	RequestBase
	Kind      string  `json:"kind"`
	Severity  string  `json:"severity"`
	Subject   string  `json:"subject"`
	Message   string  `json:"message"`
	ProjectId *string `json:"projectId,omitempty"`
}

type SupportCaseOpened struct {
	CaseId         SupportCaseId       `json:"caseId"`
	AccountId      AccountId           `json:"accountId"`
	ProjectId      *ProjectId          `json:"projectId,omitempty"`
	ReporterId     *string             `json:"reporterId,omitempty"`
	Kind           SupportCaseKind     `json:"kind,omitempty"`
	Severity       SupportCaseSeverity `json:"severity,omitempty"`
	Subject        string              `json:"subject"`
	DeploymentMode DeploymentMode      `json:"deploymentMode,omitempty"`
	GatewayVersion *string             `json:"gatewayVersion,omitempty"`
	Region         *string             `json:"region,omitempty"`
	PlanTier       *string             `json:"planTier,omitempty"`
	OpenedOn       UtcDateTime         `json:"openedOn"`
}

type SupportCaseTriaged struct {
	CaseId         SupportCaseId       `json:"caseId"`
	Kind           SupportCaseKind     `json:"kind,omitempty"`
	Severity       SupportCaseSeverity `json:"severity,omitempty"`
	AffectedModule *string             `json:"affectedModule,omitempty"`
	TriagedBy      *string             `json:"triagedBy,omitempty"`
	TriagedOn      UtcDateTime         `json:"triagedOn"`
}

type SupportCaseMessageAppended struct {
	CaseId  SupportCaseId     `json:"caseId"`
	Message SupportMessageRef `json:"message"`
}

type SupportCaseStatusChanged struct {
	CaseId    SupportCaseId     `json:"caseId"`
	From      SupportCaseStatus `json:"from,omitempty"`
	To        SupportCaseStatus `json:"to,omitempty"`
	ChangedOn UtcDateTime       `json:"changedOn"`
}

type SupportCaseResolved struct {
	CaseId     SupportCaseId  `json:"caseId"`
	Resolution CaseResolution `json:"resolution"`
	ResolvedOn UtcDateTime    `json:"resolvedOn"`
}

type SupportCaseClosed struct {
	CaseId   SupportCaseId          `json:"caseId"`
	ClosedBy *string                `json:"closedBy,omitempty"`
	ClosedOn UtcDateTime            `json:"closedOn"`
	Reason   SupportCaseCloseReason `json:"reason,omitempty"`
}

type SupportCaseReopened struct {
	CaseId     SupportCaseId `json:"caseId"`
	Reason     string        `json:"reason"`
	ReopenedOn UtcDateTime   `json:"reopenedOn"`
}

type SupportCaseWaitingReminderSent struct {
	CaseId   SupportCaseId `json:"caseId"`
	TierDays int           `json:"tierDays,omitempty"`
	SentOn   UtcDateTime   `json:"sentOn"`
}

type SupportCaseAttachmentLinked struct {
	CaseId        SupportCaseId `json:"caseId"`
	AttachmentRef string        `json:"attachmentRef"`
	FileName      *string       `json:"fileName,omitempty"`
	LinkedOn      UtcDateTime   `json:"linkedOn"`
}

// @Route("/{version}/diagnostics/packs", "GET")
type GetDiagnosticPacks struct {
	CodeMashRequestBase
}

// @Route("/{version}/diagnostics/packs/{PackName}/run", "POST")
type RunDiagnosticPackRequest struct {
	CodeMashRequestBase
	PackName    string  `json:"packName"`
	PackVersion *int    `json:"packVersion,omitempty"`
	CaseId      *string `json:"caseId,omitempty"`
}

// @Route("/{version}/diagnostics/echo", "GET")
type GetDiagnosticEcho struct {
	CodeMashRequestBase
	CaseId *string `json:"caseId,omitempty"`
}

// @Route("/{version}/diagnostics/events", "GET")
type ReadDiagnosticEventsRequest struct {
	CodeMashRequestBase
	Stream string  `json:"stream"`
	From   int64   `json:"from,omitempty"`
	Count  int     `json:"count,omitempty"`
	CaseId *string `json:"caseId,omitempty"`
}

// @Route("/{version}/diagnostics/logs", "GET")
type QueryDiagnosticLogsRequest struct {
	CodeMashListPaginationRequestBase
	Level            *string    `json:"level,omitempty"`
	Module           *string    `json:"module,omitempty"`
	LogCorrelationId *string    `json:"logCorrelationId,omitempty"`
	EventCode        *string    `json:"eventCode,omitempty"`
	Search           *string    `json:"search,omitempty"`
	FromUtc          *time.Time `json:"fromUtc,omitempty"`
	ToUtc            *time.Time `json:"toUtc,omitempty"`
	CaseId           *string    `json:"caseId,omitempty"`
}

// @Route("/{version}/diagnostics/redis", "GET")
type InspectDiagnosticRedisRequest struct {
	CodeMashRequestBase
	KeyPattern string  `json:"keyPattern"`
	CaseId     *string `json:"caseId,omitempty"`
}

// @Route("/{version}/diagnostics/health/{CheckId}", "POST")
type RunDiagnosticHealthCheckRequest struct {
	CodeMashRequestBase
	CheckId string  `json:"checkId"`
	CaseId  *string `json:"caseId,omitempty"`
}

/** @description Sign In */
// @Route("/auth", "GET,POST")
// @Route("/auth/{provider}", "GET,POST")
// @Route("/v3/auth", "POST,GET,OPTIONS")
// @Route("/v3/auth/{provider}", "POST,GET,OPTIONS")
// @Route("/v3/staff/auth", "POST,GET,OPTIONS")
// @Route("/v3/staff/auth/{provider}", "POST,GET,OPTIONS")
// @Api(Description="Sign In")
// @DataContract
type Authenticate struct {
	/** @description AuthProvider, e.g. credentials */
	// @DataMember(Order=1)
	Provider *string `json:"provider,omitempty"`
	// @DataMember(Order=2)
	UserName *string `json:"userName,omitempty"`
	// @DataMember(Order=3)
	Password *string `json:"password,omitempty"`
	// @DataMember(Order=4)
	RememberMe *bool `json:"rememberMe,omitempty"`
	// @DataMember(Order=5)
	AccessToken *string `json:"accessToken,omitempty"`
	// @DataMember(Order=6)
	AccessTokenSecret *string `json:"accessTokenSecret,omitempty"`
	// @DataMember(Order=7)
	ReturnUrl *string `json:"returnUrl,omitempty"`
	// @DataMember(Order=8)
	ErrorView *string `json:"errorView,omitempty"`
	// @DataMember(Order=9)
	Meta map[string]string `json:"meta,omitempty"`
}

// @Route("/access-token")
// @DataContract
type GetAccessToken struct {
	// @DataMember(Order=1)
	RefreshToken *string `json:"refreshToken,omitempty"`
	// @DataMember(Order=2)
	Meta map[string]string `json:"meta,omitempty"`
}
