package cli

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/cli/common"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/internal/util"
	"github.com/cholick/midi-cli/pkg/midi"
	"github.com/spf13/cobra"
)

func NewCCCommand(opener midi.Opener, con ui.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cc",
		Short: "Send control change messages",
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
				return fmt.Errorf("controller number must be be 0-127 (inclusive)")
			}

			val, err := cmd.Flags().GetInt("value")
			if err != nil {
				return fmt.Errorf("error getting value flag: %w", err)
			}
			if val < 0 || val > 127 {
				return fmt.Errorf("controller value must be be 0-127 (inclusive)")
			}

			con.Debugf("Control change %v of %v on %v", num, val, fv.Channel)

			out, err := opener.NewOutForPort(fv.Port)
			if err != nil {
				return fmt.Errorf("error opening MIDI out: %w", err)
			}
			defer out.Close()

			err = out.ControlChange(num, val, fv.Channel)
			if err != nil {
				return fmt.Errorf("error sending program change: %w", err)
			}

			return nil
		},
	}

	common.AddFlags(cmd)
	cmd.PersistentFlags().IntP("number", "n", 0, "Controller number")
	cmd.PersistentFlags().IntP("value", "l", 0, "Controller value")

	_ = cmd.MarkPersistentFlagRequired("number")
	_ = cmd.MarkPersistentFlagRequired("value")

	return cmd
}
