# Twitch Plays Anything

A little bot written in Go to make your own TwitchPlaysPokemon!

## Dependencies

* `vJoy` if you want to use the vJoy ouput (recommended, windows-only)
* Windows if you plan on using the "keyboard" output

## Compiling

```
$ go get github.com/simsor/twitchplays
```

Then, rename `twitchplays.toml.example` to `twitchplays.toml` and edit it to fit your needs.

Finally, run `twitchplays`

## Setting up your game

If you chose the vJoy output, you can enter "setup mode" to simulate joystick presses to configure your game.

Run "twitchplays -setup" to enter this mode.

For now, only the buttons present on the original GameBoy are supported, but you can relatively easily extend it by tinkering with the code.