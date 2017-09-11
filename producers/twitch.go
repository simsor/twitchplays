package producers

import (
	"fmt"
	"strings"

	"github.com/simsor/twitchplays/controllers"
	irc "github.com/thoj/go-ircevent"
)

// TwitchKeysProducer connects to a Twitch salon
type TwitchKeysProducer struct {
	Username string
	Password string
	Channel  string

	irc         *irc.Connection
	realChannel string
}

// NewTwitch instantiates a new TwitchKeysProducer
func NewTwitch(username, password, channel string) (t *TwitchKeysProducer) {
	t = &TwitchKeysProducer{}
	t.Username = strings.ToLower(username)
	t.Password = password
	t.Channel = channel

	return t
}

// Init inits an IRC connection to the Twitch chat
func (tkp *TwitchKeysProducer) Init() {
	tkp.irc = irc.IRC(tkp.Username, tkp.Username)
	tkp.realChannel = fmt.Sprintf("#%s", tkp.Channel)
	tkp.irc.Password = tkp.Password
}

// ReadKeys starts reading keys from the chat
func (tkp *TwitchKeysProducer) ReadKeys(keys chan controllers.KeyInput) {
	tkp.irc.AddCallback("001", func(event *irc.Event) {
		tkp.irc.Join(tkp.realChannel)
	})

	tkp.irc.AddCallback("PRIVMSG", func(event *irc.Event) {
		channel := event.Arguments[0]
		sender := event.Nick
		message := event.Message()

		if controllers.IsValidButton(message) && channel == tkp.realChannel {
			keys <- controllers.KeyInput{
				Sender: sender,
				Key:    message,
			}
		}
	})

	tkp.irc.Connect("irc.chat.twitch.tv:6667")
}
