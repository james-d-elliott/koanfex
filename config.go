package main

type Config struct {
	Sub SubConfig `koanf:"sub"`
}

type SubConfig struct {
	Example    string `koanf:"example"`
	Additional bool   `koanf:"additional"`
}
