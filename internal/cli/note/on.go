package note

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

func NewOnCommand(opener midi.Opener, con ui.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "on",
		Short: "Send note on",
		RunE: func(cmd *cobra.Command, args []string) error {
			fv, err := getFlagValues(cmd)
			if err != nil {
				return err
			}
			con.Debugf("Note on %v at %v on %v", fv.Note, fv.Velocity, fv.Channel)

			out, err := opener.NewOutForPort(fv.Port)
			if err != nil {
				return fmt.Errorf("error opening MIDI out: %w", err)
			}
			defer out.Close()

			err = out.NoteOn(fv.Note, fv.Velocity, fv.Channel)
			if err != nil {
				return fmt.Errorf("error playing note: %w", err)
			}

			return nil
		},
	}

	return cmd
}
