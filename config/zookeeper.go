package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	log "project1/logger"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
	"gopkg.in/yaml.v2"
)

type ZookeeperLoader struct{}

func NewZookeeperLoader() Loader {
	return new(ZookeeperLoader)
}

var (
	zkCon *zk.Conn
)

func init() {

	if os.Getenv(`ZK_CONFIG`) == `true` {
		c, _, err := zk.Connect(strings.Split(os.Getenv(`ZK_HOSTS`), `,`), time.Second) //*10)
		if err != nil {
			log.Fatal(err)
		}

		log.Info(`Zookeeper connected`)
		zkCon = c
	}
}

func (l *ZookeeperLoader) Load(path string, i interface{}) error {

	if os.Getenv(`ZK_CONFIG`) == `true` {
		return l.fromZookeeper(os.Getenv(`ZK_CONFIG_PATH`)+path, i)
	}

	return l.fromFile(path, i)
}

func (l *ZookeeperLoader) fromFile(path string, i interface{}) error {

	byt, err := ioutil.ReadFile(path + `.yaml`)
	if err != nil {
		log.Fatal(`cannot read file `+path, err)
	}

	if err := yaml.UnmarshalStrict(byt, i); err != nil {
		log.Error(err)
	}

	return err
}

func (l *ZookeeperLoader) fromZookeeper(path string, i interface{}) error {
	byt, _, err := zkCon.Get("/" + path)
	if err != nil {
		log.Fatal(`zookeeper cannot read path `+path, err)
	}

	if err := json.Unmarshal(byt, i); err != nil {
		log.Error(err)
	}

	return err
}
