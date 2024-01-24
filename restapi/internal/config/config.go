package config

type Config struct {
	Env string 'yaml:"env" env-default: "local"'
	StoragePath string 'yaml:"storage_path" env-required:"true"'
}

type HTTPServer struct {
	Address string 'yaml:"address" env-default:"localhost:8080"'
	
}