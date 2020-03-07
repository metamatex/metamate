package types

type Config struct {
	Log              bool
	DriverName       string
	DataSource       string
	MaxOpenConns     int
	TypeNames        []string
}
