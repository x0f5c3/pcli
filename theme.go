package pcli

import "github.com/x0f5c3/pterm"

// HelpSectionPrinter is used to print the different sections in the CLIs help output.
// You can overwrite it with any Sprint function.
var HelpSectionPrinter = pterm.DefaultSection.WithLevel(2).Sprint
