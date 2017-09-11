package producers

import (
	"math/rand"
	"time"

	"github.com/simsor/twitchplays/controllers"
)

// RandomKeysProducer is a test KeysReader sending random inputs
type RandomKeysProducer struct {
}

// Init does nothing
func (kp RandomKeysProducer) Init() {

}

// ReadKeys always sends the A key
func (kp RandomKeysProducer) ReadKeys(keys chan controllers.KeyInput) {
	var validKeys = []string{"a", "b", "start", "select", "up", "down", "left", "right"}
	for {
		keys <- controllers.KeyInput{
			Sender: "Ghost",
			Key:    getRandomChoice(validKeys),
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func getRandomChoice(a []string) string {
	var l = len(a)
	var i = rand.Intn(l)
	return a[i]
}
