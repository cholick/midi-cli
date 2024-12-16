import os
import subprocess
import unittest

import mido
import pretty_midi
import rtmidi


class TestE2E(unittest.TestCase):
    def setUp(self):
        self.messages = []
        self.port_name = "midi-integration-test"

        self.midi_in = rtmidi.MidiIn()
        self.midi_in.set_callback(self.midi_callback)
        self.midi_in.open_virtual_port(self.port_name)

    def tearDown(self):
        self.midi_in.close_port()

    def midi_callback(self, event, data=None):
        message, delta_time = event
        midi_message = mido.Message.from_bytes(message)
        self.messages.append(midi_message)

    def test_port_list(self):
        dir_path = os.path.dirname(os.path.realpath(__file__))
        go_path = os.path.join(dir_path, "..", "..")
        result = subprocess.run(
            "go run cmd/midi-cli/main.go port list",
            shell=True, capture_output=True, cwd=go_path,
        )

        if result.returncode != 0:
            print("Command stdout")
            print(decode(result.stdout))
            print("Command stderr")
            print(decode(result.stderr))

        self.assertEqual(0, result.returncode)
        self.assertIn(self.port_name, decode(result.stdout))

    def test_send_note_defaults(self):
        dir_path = os.path.dirname(os.path.realpath(__file__))
        go_path = os.path.join(dir_path, "..", "..")
        result = subprocess.run(
            f"go run cmd/midi-cli/main.go note on -n c4 --port {self.port_name}",
            shell=True, capture_output=True, cwd=go_path,
        )

        if result.returncode != 0:
            print("Command stdout")
            print(decode(result.stdout))
            print("Command stderr")
            print(decode(result.stderr))

        self.assertEqual(1, len(self.messages))
        self.assertEqual('note_on', self.messages[0].type)
        self.assertEqual(127, self.messages[0].velocity)
        self.assertEqual(0, self.messages[0].channel)
        self.assertEqual("C4", pretty_midi.note_number_to_name(self.messages[0].note))

    def test_send_note(self):
        dir_path = os.path.dirname(os.path.realpath(__file__))
        go_path = os.path.join(dir_path, "..", "..")
        result = subprocess.run(
            f"go run cmd/midi-cli/main.go note on -n d4 -o 121 -c 4 --port {self.port_name}",
            shell=True, capture_output=True, cwd=go_path,
        )

        if result.returncode != 0:
            print("Command stdout")
            print(decode(result.stdout))
            print("Command stderr")
            print(decode(result.stderr))

        self.assertEqual(1, len(self.messages))
        self.assertEqual('note_on', self.messages[0].type)
        self.assertEqual(121, self.messages[0].velocity)
        self.assertEqual(3, self.messages[0].channel)
        self.assertEqual("D4", pretty_midi.note_number_to_name(self.messages[0].note))

    def test_program_change(self):
        dir_path = os.path.dirname(os.path.realpath(__file__))
        go_path = os.path.join(dir_path, "..", "..")
        result = subprocess.run(
            f"go run cmd/midi-cli/main.go pc -n 4 --port {self.port_name}",
            shell=True, capture_output=True, cwd=go_path,
        )

        if result.returncode != 0:
            print("Command stdout")
            print(decode(result.stdout))
            print("Command stderr")
            print(decode(result.stderr))

        self.assertEqual(1, len(self.messages))
        self.assertEqual('program_change', self.messages[0].type)
        self.assertEqual(4, self.messages[0].program)

    def test_control_change(self):
        dir_path = os.path.dirname(os.path.realpath(__file__))
        go_path = os.path.join(dir_path, "..", "..")
        result = subprocess.run(
            f"go run cmd/midi-cli/main.go cc -n 3 -l 33 --port {self.port_name}",
            shell=True, capture_output=True, cwd=go_path,
        )

        if result.returncode != 0:
            print("Command stdout")
            print(decode(result.stdout))
            print("Command stderr")
            print(decode(result.stderr))

        self.assertEqual(1, len(self.messages))
        self.assertEqual('control_change', self.messages[0].type)
        self.assertEqual(3, self.messages[0].control)
        self.assertEqual(33, self.messages[0].value)


def decode(data: bytes):
    return data.decode("utf-8")
