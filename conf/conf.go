package conf

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfFile struct {
	Mysql struct {
		Host      string   `yaml:"host"`
		Port      int      `yaml:"port"`
		User      string   `yaml:"user"`
		Password  string   `yaml:"password"`
		Charset   string   `yaml:"charset"`
		ParseTime bool     `yaml:"parseTime"`
		Location  string   `yaml:"loc"`
		Db        []string `yaml:"db"`
	}
	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Auth     string `yaml:"auth"`
		Protocol string `yaml:"protocol"`
		Db       int    `yaml:"db"`
	}
	Http struct {
		Host           string `yaml:"host"`
		Port           int    `yaml:"port"`
		ReadTimeOut    int    `yaml:"read_timeout"`
		WriteTimeOut   int    `yaml:"write_timeout"`
		MaxHeaderBytes uint64 `yaml:"max_header_bytes"`
	}
}

type AppConf struct {
	MysqlDsn     string
	RedisOptions struct {
		Addr     string
		Password string
		Db       int
	}
	HttpOptions struct {
		Addr           string
		ReadTimeout    int
		WriteTimeout   int
		MaxHeaderBytes uint64
	}
}

func (c *AppConf) parseConfFile() *ConfFile {
	var conf ConfFile
	confFile, err := ioutil.ReadFile("app.yml")

	if err != nil {
		log.Printf("无法读取配置文件 #%v ", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(confFile, &conf)
	if err != nil {
		log.Printf("解析配置文件出错 #%v ", err)
		os.Exit(1)
	}

	return &conf
}

func (c *AppConf) loadMysqlDsn(conf *ConfFile) string {
	return fmt.Sprintf("%s@(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s", conf.Mysql.User, conf.Mysql.Host, conf.Mysql.Port, conf.Mysql.Db, conf.Mysql.Charset, conf.Mysql.ParseTime, conf.Mysql.Location)
}

func (c *AppConf) loadRedisOptions(conf *ConfFile) {
	c.RedisOptions.Addr = fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port)
	c.RedisOptions.Password = conf.Redis.Auth
	c.RedisOptions.Db = conf.Redis.Db
}

func (c *AppConf) loadHttpOptions(conf *ConfFile) {
	c.HttpOptions.Addr = fmt.Sprintf("%s:%d", conf.Http.Host, conf.Http.Port)
	c.HttpOptions.ReadTimeout = conf.Http.ReadTimeOut
	c.HttpOptions.WriteTimeout = conf.Http.WriteTimeOut
	c.HttpOptions.MaxHeaderBytes = conf.Http.MaxHeaderBytes
}

func (c *AppConf) Load() {
	conf := c.parseConfFile()
	c.loadMysqlDsn(conf)
	c.loadRedisOptions(conf)
	c.loadHttpOptions(conf)
}
