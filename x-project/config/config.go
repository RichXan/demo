package config

type Config struct {
	System   SystemConfig   `yaml:"system"`
	Log      LogConfig      `yaml:"log"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
	Smtp     SmtpConfig     `yaml:"smtp"`
}

type SystemConfig struct {
	Name        string     `yaml:"name"`
	Version     string     `yaml:"version"`
	Description string     `yaml:"description"`
	Env         string     `yaml:"env"`
	Port        int        `yaml:"port"`
	Http        HttpConfig `yaml:"http"`
}

type HttpConfig struct {
	ReadTimeout  int `yaml:"read_timeout"`
	WriteTimeout int `yaml:"write_timeout"`
	IdleTimeout  int `yaml:"idle_timeout"`
}

type LogConfig struct {
	Level     string `yaml:"level"`
	IsConsole bool   `yaml:"is_console"`
	Directory string `yaml:"directory"`
	MaxSize   int    `yaml:"max_size"`
	MaxBackup int    `yaml:"max_backups"`
}

type DatabaseConfig struct {
	Path        string `yaml:"path"` // host + port
	Database    string `yaml:"database"`
	Config      string `yaml:"config"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	MaxIdleConn int    `yaml:"max_idle_conns"`
	MaxOpenConn int    `yaml:"max_open_conns"`
	MaxLifeSec  int    `yaml:"max_life_seconds"`
	IsConsole   bool   `yaml:"is_console"`
}

type RedisConfig struct {
	Sentinels  string `yaml:"sentinels"` //非哨兵模式时，配置多个地址表示集群模式
	Password   string `yaml:"password"`
	MasterName string `yaml:"master"` //有配置master表示哨兵模式
	Db         int    `yaml:"db"`
}

type SmtpConfig struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	From     string   `yaml:"from"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	ToEmails []string `yaml:"to_emails"`
}
