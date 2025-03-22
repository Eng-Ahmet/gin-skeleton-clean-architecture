// config/config.go
package config

func LoadConfig() {

	LoadEnv()
	ConnectToDatabase()
}
