package cli

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/cli/common"
	"github.com/cholick/midi-cli/internal/ui"
	"github.com/cholick/midi-cli/internal/util"
	"github.com/cholick/midi-cli/pkg/midi"

	"github.com/spf13/cobra"
)

func NewBSCommand(opener midi.Opener, con ui.Console) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bs",
		Short: "Send bank select (cc 0) message",
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

			val, err := cmd.Flags().GetInt("value")
			if err != nil {
				return fmt.Errorf("error getting value flag: %w", err)
			}
			if val < 0 || val > 127 {
				return fmt.Errorf("bank value must be be 0-127 (inclusive)")
			}

			con.Debugf("Bankd select %v on %v", val, fv.Channel)

			out, err := opener.NewOutForPort(fv.Port)
			if err != nil {
				return fmt.Errorf("error opening MIDI out: %w", err)
			}
			defer out.Close()

			err = out.ControlChange(0, val, fv.Channel)
			if err != nil {
				return fmt.Errorf("error sending control change: %w", err)
			}

			return nil
		},
	}

	common.AddFlags(cmd)
	cmd.PersistentFlags().IntP("value", "l", 0, "Bank value")

	_ = cmd.MarkPersistentFlagRequired("value")

	return cmd
}
