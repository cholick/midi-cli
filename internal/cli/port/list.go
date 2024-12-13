package port

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

func NewListCommand(con ui.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List ports",
		RunE: func(cmd *cobra.Command, args []string) error {
			out, err := midi.NewOut()
			if err != nil {
				return fmt.Errorf("error creting midi out: %w", err)
			}

			ports, err := out.ListPorts()
			if err != nil {
				return fmt.Errorf("error listing midi ports: %w", err)
			}

			for _, port := range ports {
				con.Info(port)
			}

			return nil
		},
	}
}
