package utils

type Config struct {
	Port       int    `yaml:"port"`
	Nic        string `yaml:"nic"`
	Url        string `yaml:"url"`
	MacAddress string `yaml:"macAddress"`
}
