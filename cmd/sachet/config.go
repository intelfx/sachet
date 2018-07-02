package main

import (
	"io/ioutil"

	"github.com/intelfx/sachet/provider/cm"
	"github.com/intelfx/sachet/provider/exotel"
	"github.com/intelfx/sachet/provider/freemobile"
	"github.com/intelfx/sachet/provider/infobip"
	"github.com/intelfx/sachet/provider/mediaburst"
	"github.com/intelfx/sachet/provider/messagebird"
	"github.com/intelfx/sachet/provider/nexmo"
	"github.com/intelfx/sachet/provider/otc"
	"github.com/intelfx/sachet/provider/telegram"
	"github.com/intelfx/sachet/provider/turbosms"
	"github.com/intelfx/sachet/provider/twilio"

	"gopkg.in/yaml.v2"
)

type ReceiverConf struct {
	Name     string
	Provider string
	To       []string
	From     string
}

var config struct {
	Providers struct {
		MessageBird messagebird.MessageBirdConfig
		Nexmo       nexmo.NexmoConfig
		Twilio      twilio.TwilioConfig
		Infobip     infobip.InfobipConfig
		Exotel      exotel.ExotelConfig
		CM          cm.CMConfig
		Telegram    telegram.TelegramConfig
		Turbosms    turbosms.TurbosmsConfig
		OTC         otc.OTCConfig
		MediaBurst  mediaburst.MediaBurstConfig
		FreeMobile  freemobile.Config
	}

	Receivers []ReceiverConf
}

// LoadConfig loads the specified YAML configuration file.
func LoadConfig(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, &config)
	if err != nil {
		return err
	}

	return nil
}
