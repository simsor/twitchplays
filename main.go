package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/simsor/twitchplays/controllers"
	"github.com/simsor/twitchplays/producers"
)

var (
	producer   producers.KeysProducer
	controller controllers.Controller
	configFile string
)

func init() {
	var setupMode = flag.Bool("setup", false, "Enter vJoy setup mode")
	flag.StringVar(&configFile, "conf", "twitchplays.toml", "Path to the configuration file")
	flag.Parse()
	if *setupMode {
		c, err := controllers.NewVJoyController()
		if err != nil {
			fmt.Println("Error setting up vJoy:", err)
			os.Exit(1)
		}
		c.InteractiveMode()
		os.Exit(0)
	}
}

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	config, err := ReadConfig(configFile)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return
	}

	if config.Reader == "twitch" {
		producer = producers.NewTwitch(config.Twitch.Username, config.Twitch.Password, config.Twitch.Channel)
	} else if config.Reader == "irc" {
		producer = &producers.IRCKeysProducer{
			Server:  config.IRC.Server,
			Port:    config.IRC.Port,
			Channel: config.IRC.Channel,
		}
	} else if config.Reader == "random" {
		producer = producers.RandomKeysProducer{}
	} else {
		fmt.Println("Error: reader", config.Reader, "unknown")
		return
	}

	if config.Controller == "keyboard" {
		controller, _ = controllers.NewKeyboardController()
	} else if config.Controller == "vjoy" {
		controller, _ = controllers.NewVJoyController()
	} else {
		fmt.Println("Error: controller", config.Controller, "unknown")
		return
	}

	keys := make(chan controllers.KeyInput)

	producer.Init()

	go processKeys(keys)
	go producer.ReadKeys(keys)

	fmt.Println("TwitchPlays is running. Press Ctrl+C to stop")
	<-quit

	close(keys)

	time.Sleep(500 * time.Millisecond)
}

func processKeys(keyChan chan controllers.KeyInput) {
	for k := range keyChan {
		fmt.Printf("<%s> %s\n", k.Sender, k.Key)
		go controller.Press(k.Key)
	}

	fmt.Println("No more keys...")
}
