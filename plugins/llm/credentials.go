package llm

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/importer"
	"github.com/1Password/shell-plugins/sdk/provision"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

// Custom field names for llm-specific fields
const (
	OpenAIAPIKey     sdk.FieldName = "OpenAI API Key"
	OpenRouterAPIKey sdk.FieldName = "OpenRouter API Key"
)

func Credentials() schema.CredentialType {
	return schema.CredentialType{
		Name:          credname.Credentials,
		DocsURL:       sdk.URL("https://llm.datasette.io/en/stable/setup.html"),
		ManagementURL: nil,
		Fields: []schema.CredentialField{
			{
				Name:                OpenAIAPIKey,
				MarkdownDescription: "The OpenAI API key used to authenticate to OpenAI models.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 51,
					Prefix: "sk-",
					Charset: schema.Charset{
						Uppercase: true,
						Lowercase: true,
						Digits:    true,
					},
				},
			},
			{
				Name:                OpenRouterAPIKey,
				MarkdownDescription: "The OpenRouter API key used to authenticate to OpenRouter models.",
				Secret:              true,
				Optional:            true,
				Composition: &schema.ValueComposition{
					Length: 51,
					Prefix: "sk-or-",
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
	"OPENAI_API_KEY":     OpenAIAPIKey,
	"OPENROUTER_API_KEY": OpenRouterAPIKey,
}
