package restic

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func ResticCLI() schema.Executable {
	return schema.Executable{
		Name:      "Restic CLI",
		Runs:      []string{"restic"},
		DocsURL:   sdk.URL("https://restic.readthedocs.io"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.Credentials,
			},
		},
	}
}
