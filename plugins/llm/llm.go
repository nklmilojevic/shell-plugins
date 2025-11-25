package llm

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func LLMCLI() schema.Executable {
	return schema.Executable{
		Name:      "LLM CLI",
		Runs:      []string{"llm"},
		DocsURL:   sdk.URL("https://llm.datasette.io/en/stable/"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.Credentials,
			},
		},
	}
}
