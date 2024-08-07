package constant

type Dialect string

func (d Dialect) String() string {
	return string(d)
}

const (
	PostgresDialect Dialect = "postgres"
)

type TableName string

func (t TableName) String() string {
	return string(t)
}

const (
	PlatformTableName        TableName = "platform"
	CategoryTableName        TableName = "category"
	GameTableName            TableName = "game"
	RoleTableName            TableName = "role"
	TransaksiTableName       TableName = "transaksi"
	TransaksiDetailTableName TableName = "transaksi_detail"
	UserDetailTableName      TableName = "users"
)

// type DateTimeFormat string

// func (d DateTimeFormat) String() string {
// 	return string(d)
// }

// type RegexFormat string

// func (d RegexFormat) String() string {
// 	return string(d)
// }
