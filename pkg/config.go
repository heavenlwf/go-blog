package Conf

import (
	"io/ioutil"
	"path"
	"os"
	"log"
	"gopkg.in/yaml.v2"
)

var (
	c Conf
)

type Mysql struct {
	USER 			string						`yaml:"USER"`
	PASSWORD		string						`yaml:"PASSWORD"`
	HOST			string						`yaml:"HOST"`
	PORT 			int							`yaml:"PORT"`
	NAME			string						`yaml:"NAME"`
	TablePrefix		string						`yaml:"TABLE_PREFIX"`
}

type Conf struct {
	RunMode   		bool   						`yaml:"RUN_MODE"`
	PageSize  		int    						`yaml:"PAGE_SIZE"`
	JwtSecret 		string 						`yaml:"JWT_SECRET"`
	HttpPort  		int	 						`yaml:"HTTP_PORT"`
	ReadTimeout		int							`yaml:"READ_TIMEOUT"`
	WriteTimeout	int							`yaml:"WRITE_TIMEOUT"`
	MySQL 			map[string]*Mysql			`yaml:"MySQL"`
}

func (c *Conf) getConf() Conf {
	abspath, err := os.Getwd()
	if err != nil {
		log.Fatal("获取配置文件路径失败: ", err)
		os.Exit(1)
	}
	configPath := path.Join(abspath, "config/deploy.yml")
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal("读取配置文件失败: ", err)
		os.Exit(1)
	}

	yaml.UnmarshalStrict(content, c)
	return *c
}

func init() {
	c = c.getConf()

	//data, err := json.Marshal(conf)
	//if err != nil {
	//	log.Fatal("json conf失败: ", err)
	//}
}
