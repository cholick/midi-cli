package cli

import (
	"errors"
	"fmt"

	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/internal/util"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

func NewPanicCommand(opener midi.Opener, con ui.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "panic",
		Short: "Send all notes off on all channels to all visible ports",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := util.ParentPreRun(cmd, args)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			out, err := opener.NewDefaultOut()
			if err != nil {
				return fmt.Errorf("error creting midi out: %w", err)
			}

			ports, err := out.ListPorts()
			if err != nil {
				return fmt.Errorf("error listing ports: %w", err)
			}
			out.Close()

			errs := make([]error, 0)
			for _, port := range ports {
				out, err = opener.NewOutForPort(port)
				if err != nil {
					errs = append(errs, err)
				} else {
					con.Debugf("Sending panic on channels 1-16 to %v", port)
					err = out.PanicAll()
					errs = append(errs, err)
					out.Close()
				}
			}

			return errors.Join(errs...)
		},
	}

	return cmd
}
