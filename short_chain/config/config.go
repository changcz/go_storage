package config

type SysConfig struct {
	MysqlConfig MysqlConfig `yaml:"mysql_config"`
}

type MysqlConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DbName   string `yaml:"db_name"`
}
