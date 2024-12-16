package common

import (
	"fmt"

	"github.com/spf13/cobra"
)

type FlagValues struct {
	Port    string
	Channel int
}

func GetFlagValues(cmd *cobra.Command) (*FlagValues, error) {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return nil, fmt.Errorf("error getting port flag: %w", err)
	}

	channel, err := cmd.Flags().GetInt("channel")
	if err != nil {
		return nil, fmt.Errorf("error getting channel flag: %w", err)
	}

	if channel < 1 || channel > 16 {
		return nil, fmt.Errorf("channel must be be 1-16 (inclusive)")
	}

	return &FlagValues{
		Port:    port,
		Channel: channel,
	}, nil
}

func AddFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("port", "p", "", "Port to send message")
	cmd.PersistentFlags().IntP("channel", "c", 1, "MIDI channel")

	// suppressing this error to not pollute signature, tests would catch
	_ = cmd.MarkPersistentFlagRequired("port")
}
