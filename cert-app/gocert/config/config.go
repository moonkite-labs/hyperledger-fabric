package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	MspID        string `envconfig:"MSP_ID"`
	CertPath     string `envconfig:"CERT_PATH"`
	KeystorePath string `envconfig:"KEYSTORE_PATH"`
	CCPPath      string `envconfig:"CCP_PATH"`
	TlsCertPath  string `envconfig:"TLSCERT_PATH"`
	Label        string `envconfig:"LABEL"`
	GatewayPeer  string `envconfig:"GATEWAY_PEER"`
	PeerEndpoint string `envconfig:"PEER_ENDPOINT"`
	DB_NAME      string `envconfig:"DB_NAME"`
	DB_HOST      string `envconfig:"DB_HOST"`
	DB_PORT      string `envconfig:"DB_PORT"`
	DB_USER      string `envconfig:"DB_USER"`
	DB_PASS      string `envconfig:"DB_PASS"`
}

// Parse environment variables into a Config struct
func (c *Config) ParseEnv() error {
	return envconfig.Process("", c)
}
