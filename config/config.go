package config

import (
	"encoding/json"
	"fmt"
	"go-delegator/log"
	"io"
	"os"
)

type Configuration struct {
	MainnetUrls                []string `json:"mainnet_urls"`
	TestnetUrls                []string `json:"testnet_urls"`
	AccountPath                string   `json:"account_path"`
	AccountPassword            string   `json:"account_password"`
	ProviderAddresses          []string `json:"provider_addrs"`
	IdentityRegistryAddress    string   `json:"identity_registry_addr"`
	PublicKeyResolverAddresses []string `json:"publickey_resolver_addrs"`
	ServieKeyResolverAddresses []string `json:"servicekey_resolver_addrs"`
	UseReCaptcha               bool     `json:"use_recaptcha"`
	ReCaptchaSecret            string   `json:"recaptcha_secret"`
	ReCaptchaURL               string   `json:"recaptcha_url"`
}

const configFile = "/home/trident/go/src/go-delegator/config.json"

var Config Configuration

func ReadConfig() {
	confFile, _ := os.Open(configFile)
	defer confFile.Close()
	log.Info("Reading config file: ", configFile)

	conf, err := io.ReadAll((confFile))
	Config = Configuration{}
	c := Configuration{}
	err = json.Unmarshal(conf, &Config)
	err = json.Unmarshal(conf, &c)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
