package conf

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/cloudwego/kitex/pkg/klog"
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v3"
)

const (
	EnvDev  = "dev"
	EnvTest = "test"
	EnvProd = "prod"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	Env string

	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type Mysql struct {
	Master         MysqlConnConf `yaml:"master"`
	Slave          MysqlConnConf `yaml:"slave"`
	ConnectTimeOut int           `yaml:"connectionTimeOut"`
	MaxIdle        int           `yaml:"maxIdle"`
}

type MysqlConnConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbName"`
}

type Redis struct {
	Host        string `yaml:"host"`
	Password    string `yaml:"password"`
	DB          int    `yaml:"db"`
	MaxIdle     int    `yaml:"maxIdle"`
	MaxActivate int    `yaml:"maxActivate"`
	IdleTimeOut int    `yaml:"idleTimeOut"`
}

func GetInstance() *Config {
	once.Do(initConf)
	return config
}

func initConf() {
	confDir := filepath.Join(getProjectRoot(), "conf")
	path := filepath.Join(confDir, filepath.Join(Env(), "base.yaml"))
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	config = new(Config)
	err = yaml.Unmarshal(content, config)
	if err != nil {
		klog.Error("parse yaml error - %v", err)
		panic(err)
	}
	if err := validator.Validate(config); err != nil {
		klog.Error("validate config %s error -%v", err)
		panic(err)
	}

	config.Env = Env()
}

// 获取当前项目根目录
func getProjectRoot() string {
	dirExist := func(path string) (bool, error) {
		stat, err := os.Stat(path)
		if err == nil && stat.IsDir() {
			return true, nil
		}
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	filenae, err := os.Getwd()
	if err == nil {
		dir := path.Dir(filenae)
		for i := 0; i < 3; i++ {
			p := path.Join(dir, "conf")
			exist, err := dirExist(p)
			if exist {
				return dir
			}
			if err != nil {
				klog.Errorf("GetRootPath error: %v", err)
			}
			if len(dir) <= 1 {
				break
			}
			dir = path.Dir(dir)
		}
	}

	return "."
}

// 获取执行环境
func Env() string {
	var defaultEnv = EnvProd
	if e := os.Getenv("env"); e == EnvDev || e == EnvTest {
		defaultEnv = EnvDev
	}
	return defaultEnv
}
