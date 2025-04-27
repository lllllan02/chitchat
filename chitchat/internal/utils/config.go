package utils

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config 全局配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Upload   UploadConfig   `mapstructure:"upload"`
	Redis    RedisConfig    `mapstructure:"redis"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver      string `mapstructure:"driver"`
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname"`
	Charset     string `mapstructure:"charset"`
	MaxIdleConn int    `mapstructure:"max_idle_conns"`
	MaxOpenConn int    `mapstructure:"max_open_conns"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire string `mapstructure:"expire"`
}

// UploadConfig 文件上传配置
type UploadConfig struct {
	MaxSize      int      `mapstructure:"max_size"`
	AllowedTypes []string `mapstructure:"allowed_types"`
	StoragePath  string   `mapstructure:"storage_path"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Enabled  bool   `mapstructure:"enabled"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

var AppConfig Config

// LoadConfig 加载配置文件
func LoadConfig(path string) error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("未找到配置文件，将使用默认配置和环境变量")
			return createDefaultConfig()
		}
		return fmt.Errorf("读取配置文件错误: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		return fmt.Errorf("解析配置文件错误: %v", err)
	}

	return nil
}

// 创建默认配置
func createDefaultConfig() error {
	AppConfig = Config{
		Server: ServerConfig{
			Port: 8080,
			Mode: "development",
		},
		Database: DatabaseConfig{
			Driver:      "mysql",
			Host:        "localhost",
			Port:        3306,
			Username:    "root",
			Password:    "password",
			DBName:      "chitchat",
			Charset:     "utf8mb4",
			MaxIdleConn: 10,
			MaxOpenConn: 100,
		},
		JWT: JWTConfig{
			Secret: "default_secret_key",
			Expire: "24h",
		},
		Upload: UploadConfig{
			MaxSize:      5,
			AllowedTypes: []string{"image/jpeg", "image/png", "image/gif"},
			StoragePath:  "./uploads",
		},
		Redis: RedisConfig{
			Enabled:  false,
			Host:     "localhost",
			Port:     6379,
			Password: "",
			DB:       0,
		},
	}
	return nil
}

// GetDSN 获取数据库连接字符串
func GetDSN() string {
	db := AppConfig.Database
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		db.Username, db.Password, db.Host, db.Port, db.DBName, db.Charset)
}
