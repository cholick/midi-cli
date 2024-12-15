package note

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

func NewOffCommand(opener midi.Opener, con ui.Console) *cobra.Command {
	return &cobra.Command{
		Use:   "off",
		Short: "Send note off",
		RunE: func(cmd *cobra.Command, args []string) error {
			fv, err := getFlagValues(cmd)
			if err != nil {
				return err
			}
			con.Debugf("Note off %v on %v", fv.Note, fv.Channel)

			out, err := opener.NewOutForPort(fv.Port)
			if err != nil {
				return fmt.Errorf("error opening MIDI out: %w", err)
			}
			defer out.Close()

			err = out.NoteOff(fv.Note, fv.Velocity, fv.Channel)
			if err != nil {
				return fmt.Errorf("error playing note: %w", err)
			}

			return nil
		},
	}
}
