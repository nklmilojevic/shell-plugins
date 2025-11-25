package restic

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "restic",
		Platform: schema.PlatformInfo{
			Name:     "Restic",
			Homepage: sdk.URL("https://restic.net"),
		},
		Credentials: []schema.CredentialType{
			Credentials(),
		},
		Executables: []schema.Executable{
			ResticCLI(),
		},
	}
}
