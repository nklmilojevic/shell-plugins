package llm

import (
	"testing"

	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/plugintest"
)

func TestCredentialsProvisioner(t *testing.T) {
	plugintest.TestProvisioner(t, Credentials().DefaultProvisioner, map[string]plugintest.ProvisionCase{
		"with OpenAI": {
			ItemFields: map[sdk.FieldName]string{
				OpenAIAPIKey: "sk-1234567890abcdefghijklmnopqrstuvwxyzEXAMPLEKEY",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"OPENAI_API_KEY": "sk-1234567890abcdefghijklmnopqrstuvwxyzEXAMPLEKEY",
				},
			},
		},
		"with OpenRouter": {
			ItemFields: map[sdk.FieldName]string{
				OpenRouterAPIKey: "sk-or-1234567890abcdefghijklmnopqrstuvwxEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"OPENROUTER_API_KEY": "sk-or-1234567890abcdefghijklmnopqrstuvwxEXAMPLE",
				},
			},
		},
		"with both": {
			ItemFields: map[sdk.FieldName]string{
				OpenAIAPIKey: "sk-1234567890abcdefghijklmnopqrstuvwxyzEXAMPLEKEY",
				OpenRouterAPIKey: "sk-or-1234567890abcdefghijklmnopqrstuvwxEXAMPLE",
			},
			ExpectedOutput: sdk.ProvisionOutput{
				Environment: map[string]string{
					"OPENAI_API_KEY":     "sk-1234567890abcdefghijklmnopqrstuvwxyzEXAMPLEKEY",
					"OPENROUTER_API_KEY": "sk-or-1234567890abcdefghijklmnopqrstuvwxEXAMPLE",
				},
			},
		},
	})
}

func TestCredentialsImporter(t *testing.T) {
	plugintest.TestImporter(t, Credentials().Importer, map[string]plugintest.ImportCase{
		"environment": {
			Environment: map[string]string{
				"OPENAI_API_KEY":     "sk-1234567890abcdefghijklmnopqrstuvwxyzEXAMPLEKEY",
				"OPENROUTER_API_KEY": "sk-or-1234567890abcdefghijklmnopqrstuvwxEXAMPLE",
			},
			ExpectedCandidates: []sdk.ImportCandidate{
				{
					Fields: map[sdk.FieldName]string{
						OpenAIAPIKey: "sk-1234567890abcdefghijklmnopqrstuvwxyzEXAMPLEKEY",
						OpenRouterAPIKey: "sk-or-1234567890abcdefghijklmnopqrstuvwxEXAMPLE",
					},
				},
			},
		},
	})
}
