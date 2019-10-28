package auth

type Config struct {
	IsAccount bool
}

type Auth struct {
	config *Config
}

func New(conf *Config) *Auth {
	if conf == nil {
		conf = &Config{IsAccount: false}
	}
	auth := &Auth{
		conf,
	}
	return auth
}
