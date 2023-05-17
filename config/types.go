package config

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type snowflake struct {
	WorkerID      int64 `yaml:"worker-id"`
	DatancenterID int64 `yaml:"datancenter-id"`
}

type service struct {
	Name string
	Addr string
	LB   bool `yaml:"load-balance"`
}

type mysql struct {
	Addr     string
	Database string
	Username string
	Password string
	Charset  string
}

type etcd struct {
	Addr string
}

type config struct {
	Server    server
	Snowflake snowflake
	Mysql     mysql
	Etcd      etcd
}
