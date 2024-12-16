package note

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/cli/common"
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

			return nil
		},
	}

	common.AddFlags(cmd)
	cmd.PersistentFlags().StringP("note", "n", "c4", "Note name (eg c4, C4, C#4, d♭4, or Db4)")
	cmd.PersistentFlags().IntP("velocity", "o", 127, "Note velocity")

	return cmd
}

type flagNoteValues struct {
	*common.FlagValues
	Note     string
	Velocity int
}

func getNoteFlagValues(cmd *cobra.Command) (*flagNoteValues, error) {
	bfv, err := common.GetFlagValues(cmd)
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
