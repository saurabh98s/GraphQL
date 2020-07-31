package config


type config struct {
	mongoConfig mongo
}

var configuration = &config{}

// Load loads the config into the configuration object
func Load() {
	configuration.mongoConfig.load()
}

// Mongo returns the mongo config
func Mongo() mongo {
	return configuration.mongoConfig
}