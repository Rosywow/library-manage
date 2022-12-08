package cmn

import "github.com/spf13/viper"

type DBConfig struct {
	IpAddress string `json:"ipAddress"`
	Port      string `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	Dbname    string `json:"dbname"`
}

type DBType struct {
	ContainerDB DBConfig `json:"containerDB"`
	LocalContainerDB DBConfig `json:"local_container_db"`
	LocalDB DBConfig `json:"local_db"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type Config struct {
	DB DBType `json:"db"`
	Server ServerConfig `json:"server"`
}

var vi *viper.Viper

func GetViperValue() (Config,error){
	vi = viper.New()
	var config Config

	vi.SetConfigName("config")
	vi.SetConfigType("json")
	//docker
	//vi.AddConfigPath("/app/config")
	//local
	vi.AddConfigPath("./cmn")
	err := vi.ReadInConfig()
	if err!=nil {
		return Config{},err
	}

	err = vi.Unmarshal(&config)
	if err!=nil {
		return Config{},err
	}

	return config,nil
}
