package config

type server struct {
	Secret  []byte
	Version string
	Name    string
}

type snowflake struct {
	WorkerID      int64
	DatancenterID int64
}

type service struct {
	Name string
	Addr string
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
