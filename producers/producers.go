package producers

import "github.com/simsor/twitchplays/controllers"

// KeysProducer defines a interface for a key Producer
type KeysProducer interface {
	Init()
	ReadKeys(keys chan controllers.KeyInput)
}
