package port

import (
	"github.com/spf13/cobra"
)

func NewPortCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "port",
		Short: "Manage MIDI ports",
	}
}
