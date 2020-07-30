package iniconfig

import (
	"io/ioutil"
	"testing"
)

type Config struct {
	ServerConf ServerConfig `ini:"server"`
	MysqlConf  MysqlConfig  `ini:"mysql"`
}

type ServerConfig struct {
	Ip   string
	Port int
}

type MysqlConfig struct {
	UserName string `ini:"username"`
	Passwd   string `ini:"passwd"`
	Database string `ini:"database"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func TestIniConfig(t *testing.T) {
	data, err := ioutil.ReadFile("./config.ini")
	if err != nil {
		t.Error("read file failed")
	}

	var conf Config
	err = UnMarshal(data, &conf)
	if err != nil {
		t.Errorf("unmarshal failed,err:%v", err)
		return
	}

	t.Logf("unmarshal success, conf:%#v", conf)
}
