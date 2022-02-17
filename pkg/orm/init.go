package orm

import "github.com/whileW/core-go/pkg/flag"

var (
	// 最多空闲连接
	defaultDbMaxIdleConns		int
	// 最多打开的连接
	defaultDbMaxOpenConns		int
	// 数据库类型：mysql,mssql,sqlserver
	defaultDbAdapter			string
)

func init()  {
	flag.Register(
		&flag.IntFlag{
			Name:	"dbmopenc",
			Usage:	"--dbmopenc=100，数据库默认最大打开连接数",
			EnvVar: "DB_MAX_OPEN_CONNS",
			Default: 100,
			Variable:&MaxOpenConns,
		},
		&flag.IntFlag{
			Name:	"dbmidlec",
			Usage:	"--dbmidlec=10，数据库默认最大空闲连接数",
			EnvVar: "DB_MAX_IDLE_CONNS",
			Default: 10,
			Variable:&MaxIdleConns,
		},
		&flag.StringFlag{
			Name:   "dbdsn",
			Usage:  "--dbdsn=root:secret@tcp(127.0.0.1:3306)/mysql?charset=utf8&parseTime=True&loc=Local，数据库连接字符串",
			EnvVar: "DB_DSN",
			Variable: &DSN,
		},
	)
}