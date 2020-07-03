package xdriveri

type DbDriverType int

const (
	MySqlDriver DbDriverType = iota
	MsSqlDriver
	OracleDriver
	PostgresDriver
	Sqlite3Driver
)

var DbDriverNames = map[DbDriverType]string{
	MySqlDriver:    "mysql",
	MsSqlDriver:    "mssql",
	OracleDriver:   "oracle",
	PostgresDriver: "postgres",
	Sqlite3Driver:  "sqlite3",
}

func (dr DbDriverType) String() string {
	if s, ok := DbDriverNames[dr]; ok {
		return s
	}
	return "unknown driver"
}

func DbDriverTypeByName(name string) DbDriverType {
	switch name {
	case "mysql":
		return MySqlDriver
	case "mssql":
		return MsSqlDriver
	case "oracle":
		return OracleDriver
	case "postgres":
		return PostgresDriver
	case "sqlite3":
		return Sqlite3Driver
	}
	return MySqlDriver
}
