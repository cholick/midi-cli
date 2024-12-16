## midi-cli

Probably use https://github.com/gbevin/SendMIDI instead of this; it has considerably more features. I only
created this because I wanted to send program changes and a few other MIDI commands using a Stream Deck, and used
that as an excuse to do some MIDI work in Go (and get myself a bit more familiar with the messages for
some hardware work).

```shell
$ midi-cli --help
Send MIDI messages

Usage:
  midi-cli [command]

Available Commands:
  cc          Send control change messages
  help        Help about any command
  note        Send note messages
  panic       Send all notes off on all channels to all visible ports
  pc          Send program change messages
  port        Manage MIDI ports

Flags:
  -h, --help      help for midi-cli
  -v, --verbose   Verbose output

Use "midi-cli [command] --help" for more information about a command.
```

### Setup

See [detailed-setup.md](docs/detailed-setup.md) for some detailed setup instructions for a Mac. It amounts to
making sure the build-in IAC driver is on in MIDI Studio, sending signals to that driver, and making sure the
stand-alone instrument / DAW is listening.
