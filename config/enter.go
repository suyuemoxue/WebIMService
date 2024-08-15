package config

type Config struct {
	MySQL  MySQL  `yaml:"mysql"`
	System System `yaml:"system"`
	Logger Logger `yaml:"logger"`
}
