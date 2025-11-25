package restic

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
	"github.com/1Password/shell-plugins/sdk/schema/fieldname"
)

func TestCredentialsProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, Credentials().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"default with AWS S3": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Password:        "mysecretpassword",
				fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
				fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"RESTIC_PASSWORD":       "mysecretpassword",
					"AWS_ACCESS_KEY_ID":     "AKIAIOSFODNN7EXAMPLE",
					"AWS_SECRET_ACCESS_KEY": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
				},
			},
		},
		"with Azure": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Password: "mysecretpassword",
				AzureAccountName:   "myaccount",
				AzureAccountKey:    "myaccountkey",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"RESTIC_PASSWORD":    "mysecretpassword",
					"AZURE_ACCOUNT_NAME": "myaccount",
					"AZURE_ACCOUNT_KEY":  "myaccountkey",
				},
			},
		},
		"with B2": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Password: "mysecretpassword",
				B2AccountID:        "myaccountid",
				B2AccountKey:       "myaccountkey",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"RESTIC_PASSWORD": "mysecretpassword",
					"B2_ACCOUNT_ID":   "myaccountid",
					"B2_ACCOUNT_KEY":  "myaccountkey",
				},
			},
		},
		"with REST server": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Password: "mysecretpassword",
				RESTUsername:       "myuser",
				RESTPassword:       "myrestpassword",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"RESTIC_PASSWORD":      "mysecretpassword",
					"RESTIC_REST_USERNAME": "myuser",
					"RESTIC_REST_PASSWORD": "myrestpassword",
				},
			},
		},
		"password only": {
			ItemFields: map[sdk.FieldName]string{
				fieldname.Password: "mysecretpassword",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"RESTIC_PASSWORD": "mysecretpassword",
				},
			},
		},
	})
}

func TestCredentialsImporter(t *testing.T) {
	plugintest.TestImporter(t, Credentials().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"RESTIC_PASSWORD":       "mysecretpassword",
				"AWS_ACCESS_KEY_ID":     "AKIAIOSFODNN7EXAMPLE",
				"AWS_SECRET_ACCESS_KEY": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						fieldname.Password:        "mysecretpassword",
						fieldname.AccessKeyID:     "AKIAIOSFODNN7EXAMPLE",
						fieldname.SecretAccessKey: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
					},
				},
			},
		},
	})
}
