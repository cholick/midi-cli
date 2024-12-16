package cc

import (
	"fmt"

	"github.com/cholick/midi-cli/internal/cli/base"
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

			fv, err := base.GetFlagValues(cmd)
			if err != nil {
				return err
			}

			// todo: move duplicated validation up to base
			if fv.Channel < 1 || fv.Channel > 16 {
				return fmt.Errorf("channel must be be 1-16 (inclusive)")
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fv, err := base.GetFlagValues(cmd)
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

	cmd.PersistentFlags().StringP("port", "p", "", "Port to send message")
	cmd.PersistentFlags().IntP("channel", "c", 1, "MIDI channel")
	cmd.PersistentFlags().IntP("number", "n", 0, "Controller number")
	cmd.PersistentFlags().IntP("value", "l", 0, "Controller value")

	_ = cmd.MarkPersistentFlagRequired("port")
	_ = cmd.MarkPersistentFlagRequired("number")
	_ = cmd.MarkPersistentFlagRequired("number")
	_ = cmd.MarkPersistentFlagRequired("value")

	return cmd
}
