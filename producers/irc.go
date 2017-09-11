package producers

import (
	"fmt"

	"github.com/simsor/twitchplays/controllers"
	irc "github.com/thoj/go-ircevent"
)

// IRCKeysProducer connects to an IRC channel and reads keys from there
type IRCKeysProducer struct {
	Server  string
	Port    int
	Channel string

	irc *irc.Connection
}

// Init creates a IRC object
func (ikp *IRCKeysProducer) Init() {
	ikp.irc = irc.IRC("TwitchBot", "BotBOt")
}

// ReadKeys connects to the IRC server and starts reading keys
func (ikp *IRCKeysProducer) ReadKeys(keys chan controllers.KeyInput) {
	ikp.irc.AddCallback("001", func(event *irc.Event) {
		ikp.irc.Join(ikp.Channel)
	})

	ikp.irc.AddCallback("PRIVMSG", func(event *irc.Event) {
		channel := event.Arguments[0]
		sender := event.Nick
		message := event.Message()

		if controllers.IsValidButton(message) && channel == ikp.Channel {
			keys <- controllers.KeyInput{
				Sender: sender,
				Key:    message,
			}
		}
	})

	ikp.irc.Connect(fmt.Sprintf("%s:%d", ikp.Server, ikp.Port))
}
