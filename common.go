package ck_orm

const (
	DriverMysql  = "mysql"
	DriverSQLite = "sqlite"
)
//数据库配置
type DBConfig struct {
	DBDriver   string `json:"db_driver" yaml:"db_driver"`
	DBServer   string `json:"db_server" yaml:"db_server"`
	DBPort     string `json:"db_port" yaml:"db_port"`
	DBName     string `json:"db_name" yaml:"db_name"`
	DBUser     string `json:"db_user" yaml:"db_user"`
	DBPassword string `json:"db_password" yaml:"db_password"`
	DBPoolSize int    `json:"db_pool_size" yaml:"db_pool_size"`
	DBIdleSize int    `json:"db_Idle_size" yaml:"db_Idle_size"`
	MaxLife    int    `json:"max_life" yaml:"max_life"`
	DBDebug    bool   `json:"db_debug" yaml:"db_debug"`
}


