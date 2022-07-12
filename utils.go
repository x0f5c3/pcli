package pcli

import (
	"github.com/spf13/cobra"
	"github.com/x0f5c3/pterm"
)

// generateTitleString generates a pretty looking title string.
func generateTitleString(rootCmd *cobra.Command) string {
	return pterm.Sprintf("\n# %s | %s\n", pterm.Cyan(rootCmd.Name()), pterm.Green(rootCmd.Version))
}
