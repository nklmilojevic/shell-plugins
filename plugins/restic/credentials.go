package restic

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

// Custom field names for restic-specific fields
const (
	AzureAccountName sdk.FieldName = "Azure Account Name"
	AzureAccountKey  sdk.FieldName = "Azure Account Key"
	AzureAccountSAS  sdk.FieldName = "Azure Account SAS"
	B2AccountID      sdk.FieldName = "B2 Account ID"
	B2AccountKey     sdk.FieldName = "B2 Account Key"
	RESTUsername     sdk.FieldName = "REST Username"
	RESTPassword     sdk.FieldName = "REST Password"
	SessionToken     sdk.FieldName = "Session Token"
)

func Credentials() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.Credentials,
		DocsURL:       sdk.URL("https://restic.readthedocs.io/en/latest/040_backup.html"),
		ManagementURL: nil,
		Fields: []schema.CredentialField{
			// Core - Required
			{
				Name:                fieldname.Password,
				MarkdownDescription: "The password used to encrypt/decrypt the restic repository.",
				Secret:              true,
				Composition: &schema.ValueComposition{
					Length: 32,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			// AWS S3 / Cloudflare R2 / MinIO - Optional
			{
				Name:                fieldname.AccessKeyID,
				MarkdownDescription: "The AWS Access Key ID for S3-compatible backends (AWS S3, Cloudflare R2, MinIO).",
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 20,
					Prefix: "AKIA",
					Charset: schema.Charset{
						Uppercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.SecretAccessKey,
				MarkdownDescription: "The AWS Secret Access Key for S3-compatible backends.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 40,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                SessionToken,
				MarkdownDescription: "The AWS Session Token for temporary credentials.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 64,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			// Azure Blob Storage - Optional
			{
				Name:                AzureAccountName,
				MarkdownDescription: "The Azure Storage account name.",
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 24,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                AzureAccountKey,
				MarkdownDescription: "The Azure Storage account key.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 88,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                AzureAccountSAS,
				MarkdownDescription: "The Azure Storage SAS token.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 64,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			// Backblaze B2 - Optional
			{
				Name:                B2AccountID,
				MarkdownDescription: "The Backblaze B2 Account ID.",
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 12,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                B2AccountKey,
				MarkdownDescription: "The Backblaze B2 Application Key.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 31,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			// Google Cloud Storage - Optional
			{
				Name:                fieldname.ProjectID,
				MarkdownDescription: "The Google Cloud project ID.",
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 20,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                fieldname.Credentials,
				MarkdownDescription: "The path to Google Cloud application credentials JSON file.",
				Optional:            true,
			},
			// REST Server - Optional
			{
				Name:                RESTUsername,
				MarkdownDescription: "The username for REST server authentication.",
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 12,
					Charset: schema.Charset{
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                RESTPassword,
				MarkdownDescription: "The password for REST server authentication.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 24,
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
		},
		DefaultProvisioner: provision.EnvVars(defaultEnvVarMapping),
		Importer:           importer.TryEnvVarPair(defaultEnvVarMapping),
	}
}

var defaultEnvVarMapping = map[string]sdk.FieldName{
	// Core
	"RESTIC_PASSWORD": fieldname.Password,
	// AWS S3 / Cloudflare R2 / MinIO
	"AWS_ACCESS_KEY_ID":     fieldname.AccessKeyID,
	"AWS_SECRET_ACCESS_KEY": fieldname.SecretAccessKey,
	"AWS_SESSION_TOKEN":     SessionToken,
	// Azure
	"AZURE_ACCOUNT_NAME": AzureAccountName,
	"AZURE_ACCOUNT_KEY":  AzureAccountKey,
	"AZURE_ACCOUNT_SAS":  AzureAccountSAS,
	// Backblaze B2
	"B2_ACCOUNT_ID":  B2AccountID,
	"B2_ACCOUNT_KEY": B2AccountKey,
	// Google Cloud
	"GOOGLE_PROJECT_ID":              fieldname.ProjectID,
	"GOOGLE_APPLICATION_CREDENTIALS": fieldname.Credentials,
	// REST Server
	"RESTIC_REST_USERNAME": RESTUsername,
	"RESTIC_REST_PASSWORD": RESTPassword,
}
