package db

type DatabaseList struct {
	Master struct {
		Mysql Database
	}
	Slave struct {
		Mysql Database
	}
}

type Database struct {
	User                 string `yaml:"username"`
	Password             string `yaml:"password"`
	Net                  string `yaml:"net"`
	Host                 string `yaml:"host"`
	Port                 string `yaml:"port"`
	DBName               string `yaml:"dbname"`
	AllowNativePasswords bool   `yaml:"allowNativePassword"`
	Adapter              string `yaml:"adapter"`
	Params               struct {
		ParseTime string
	}
}
