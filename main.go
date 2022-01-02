package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"wol/utils"
)

func main() {

	var config utils.Config
	configFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(configFile, &config); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.GET(config.Url, func(c *gin.Context) {
		magicPacket, err := utils.GetMagicPacket(config.MacAddress)
		if err != nil {
			utils.Result(0, err.Error(), nil, c)
			return
		}
		ok, err := utils.SendMagicPacket(magicPacket, config.Nic)
		if err != nil {
			utils.Result(-1, err.Error(), nil, c)
			return
		}
		if ok {
			utils.Result(0, "MagicPacket 发送成功", nil, c)
		}
	})

	if err := router.Run(fmt.Sprintf(":%v", config.Port)); err != nil {
		panic(err)
	}

}
