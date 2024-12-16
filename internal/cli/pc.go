package cli

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/cli/common"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/internal/util"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

// todo: This just sends 0-127 data byte now, but there's special meaning for the
// todo: values for some hardware (number corresponding to general MIDI instruments).
// todo: Potentially give shorthand for that if come up with a reason for it being useful
func NewPCCommand(opener midi.Opener, con ui.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pc",
		Short: "Send program change messages",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			err := util.ParentPreRun(cmd, args)
			if err != nil {
				return err
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fv, err := common.GetFlagValues(cmd)
			if err != nil {
				return err
			}

			num, err := cmd.Flags().GetInt("number")
			if err != nil {
				return fmt.Errorf("error getting number flag: %w", err)
			}
			if num < 0 || num > 127 {
				return fmt.Errorf("program number must be be 0-127 (inclusive)")
			}

			con.Debugf("Program change %v on %v", num, fv.Channel)

			out, err := opener.NewOutForPort(fv.Port)
			if err != nil {
				return fmt.Errorf("error opening MIDI out: %w", err)
			}
			defer out.Close()

			err = out.ProgramChange(num, fv.Channel)
			if err != nil {
				return fmt.Errorf("error sending program change: %w", err)
			}

			return nil
		},
	}

	common.AddFlags(cmd)
	cmd.PersistentFlags().IntP("number", "n", 0, "Program/preset number")

	_ = cmd.MarkPersistentFlagRequired("number")

	return cmd
}
