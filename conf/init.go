package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	JWTSecret string `json:"jwt_secret"`
}

var Config Conf

func Init() {
	file, err := os.Open("./conf/config.json")
	if err != nil {
		panic(err)
	}

	decoder := json.NewDecoder(file)
	Config = Conf{}
	err = decoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
}
