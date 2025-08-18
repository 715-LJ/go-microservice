package config

type MysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetPort() int
	GetUser() string
	GetPassword() string
	GetDbname() string
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
}

// defaultMysqlConfig mysql 配置
type defaultMysqlConfig struct {
	Hosts             string
	Port              int
	User              string
	Password          string
	Dbname            string
	Enable            bool
	MaxidleConnection int
	MaxopenConnection int
}

// URL pgsql 连接
func (m defaultMysqlConfig) GetURL() string {
	return m.Hosts
}
func (m defaultMysqlConfig) GetPort() int {
	return m.Port
}
func (m defaultMysqlConfig) GetUser() string {
	return m.User
}

func (m defaultMysqlConfig) GetPassword() string {
	return m.Password
}

func (m defaultMysqlConfig) GetDbname() string {
	return m.Dbname
}

// Enabled 激活
func (m defaultMysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m defaultMysqlConfig) GetMaxIdleConnection() int {
	return m.MaxidleConnection
}

// 打开连接数
func (m defaultMysqlConfig) GetMaxOpenConnection() int {
	return m.MaxopenConnection
}
