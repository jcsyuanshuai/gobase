package data

type MysqlOpts struct {
	Driver   string
	Host     string
	Database string
	Port     string
	Username string
	Password string
}

type RedisOpts struct {
	Database int
	Host     string
	Password string
}

type MongoOpts struct {
	AppUrl string
}

type MysqlEngine struct {
}

func (engine *MysqlEngine) New(b Buildable) (*Database, error) {
	return b.BuildMysqlEngine(&MysqlOpts{})
}
