package note

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/cli/base"
	"github.com/cholick/midi-cli/internal/util"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

func NewNoteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "note",
		Short: "Send note messages",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := util.ParentPreRun(cmd, args)
			if err != nil {
				return err
			}

			fv, err := getNoteFlagValues(cmd)
			if err != nil {
				return err
			}

			_, err = midi.NoteNumberFromName(fv.Note)
			if err != nil {
				return fmt.Errorf("note is invalid: %w", err)
			}

			if fv.Velocity < 0 || fv.Velocity > 127 {
				return fmt.Errorf("velocity must be 0-127 (inclusive)")
			}

			//todo: move common validation to make re-usable
			if fv.Channel < 1 || fv.Channel > 16 {
				return fmt.Errorf("channel must be be 1-16 (inclusive)")
			}

			return nil
		},
	}

	//todo: move common flags (port, channel) to make re-usable
	cmd.PersistentFlags().StringP("port", "p", "", "Port to send message")
	cmd.PersistentFlags().IntP("channel", "c", 1, "MIDI channel")
	cmd.PersistentFlags().StringP("note", "n", "c4", "Note name (eg c4, C4, C#4, d♭4, or Db4)")
	cmd.PersistentFlags().IntP("velocity", "o", 127, "Note velocity")

	_ = cmd.MarkPersistentFlagRequired("port")

	return cmd
}

type flagNoteValues struct {
	*base.FlagValues
	Note     string
	Velocity int
}

func getNoteFlagValues(cmd *cobra.Command) (*flagNoteValues, error) {
	bfv, err := base.GetFlagValues(cmd)
	if err != nil {
		return nil, err
	}

	note, err := cmd.Flags().GetString("note")
	if err != nil {
		return nil, fmt.Errorf("error getting note flag: %w", err)
	}

	velocity, err := cmd.Flags().GetInt("velocity")
	if err != nil {
		return nil, fmt.Errorf("error getting velocity flag: %w", err)
	}

	return &flagNoteValues{
		FlagValues: bfv,
		Note:       note,
		Velocity:   velocity,
	}, nil
}
