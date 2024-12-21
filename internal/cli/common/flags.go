package common

import (
	"errors"
	"fmt"
	"os"

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

	if port == "" {
		port = os.Getenv("MIDI_CLI_PORT")
	}

	if port == "" {
		return nil, errors.New("'port' flag or MIDI_CLI_PORT environment variable is required")
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
	cmd.PersistentFlags().StringP("port", "p", "", "Port to send message (also specifiable via MIDI_CLI_PORT)")
	cmd.PersistentFlags().IntP("channel", "c", 1, "MIDI channel")
}
