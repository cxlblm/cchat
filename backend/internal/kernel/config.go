package kernel

type Config struct {
	Database database `toml:"database"`
	Cache    cache    `toml:"cache"`
	Server   server   `toml:"server"`
	Nsq      nsq      `toml:"nsq"`
}

type database struct {
	Host      string `toml:"host"`
	DB        string `toml:"db"`
	User      string `toml:"user"`
	Password  string `toml:"password"`
	Port      int    `toml:"port"`
	Charset   string `toml:"charset"`
	ParseTime bool   `toml:"parseTime"`
	Loc       string `toml:"loc"`
}

type cache struct {
	Host     string `toml:"host"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Port     int    `toml:"port"`
	DB       int    `toml:"db"`
}

type server struct {
	Port int    `toml:"port"`
	Host string `toml:"host"`
}

type nsq struct {
	Use     bool   `toml:"use"`
	Topic   string `toml:"topic"`
	Lookup  string `toml:"lookup"`
	Channel string `toml:"channel"`
}
