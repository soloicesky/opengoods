package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type DBConfig struct {
	Database     string `json:"database"`  //database name
	User         string `json:"user"`      //login user
	Password     string `json:"passoword"` //login password
	Host         string `json:"host"`      //host
	Port         int    `json:"port"`      //port
	Charset      string `json:"charset"`   //charset
	URL          string `json:"url"`       //url
	MaxIdleConns int    `json:"maxconn"`   //max idle connections
	MaxOpenConns int    `json:"maxopne"`   //max open connections
}

func (cfg *DBConfig) ToString() string {
	data, err := json.Marshal(cfg)

	if nil != err {
		return ""
	}

	return string(data)
}

type ServerConfig struct {
	Version     string `json:"version"`     //version
	ReleaseDate string `json:"releasedate"` //release date
	Url         string `json:"url"`         //url
	Port        int    `json:port`          //port
}

func (cfg *ServerConfig) ToString() string {
	data, err := json.Marshal(cfg)

	if nil != err {
		return ""
	}

	return string(data)
}

var dbConfig DBConfig
var serverConfig ServerConfig

func DBCfg() *DBConfig {
	return &dbConfig
}

func ServerCfg() *ServerConfig {
	return &serverConfig
}

func init_db() error {
	bytes, err := ioutil.ReadFile("./db.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &dbConfig); err != nil {
		fmt.Println("invalid config: ", err.Error())
		return err
	}

	return nil
}

func init_server() error {
	bytes, err := ioutil.ReadFile("./server.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &serverConfig); err != nil {
		fmt.Println("invalid config: ", err.Error())
		return err
	}

	return nil
}

func init() {
	err := init_db()

	if nil != err {
		fmt.Println("init dabase config failed", err.Error())
		os.Exit(-1)
	}

	err = init_server()

	if nil != err {
		fmt.Println("init server config failed", err.Error())
		os.Exit(-1)
	}
}
