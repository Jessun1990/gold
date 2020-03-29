package db

import (
	"fmt"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dsnFuncMap["mysql"] = getMysqlDsn
}

// getMysqlDsn 输出格式:"用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
// "[username]:[password]@[protocol[(address:port)]]/[dbname][?param1=value1&...&paramN=valueN]"
func getMysqlDsn(stdConfig StdConfig) string {
	if stdConfig.Port != "" {
		stdConfig.Address = fmt.Sprintf("%s:%s", stdConfig.Address, stdConfig.Port)
	}
	return fmt.Sprintf("%s:%s@%s(%s)/%s%s",
		stdConfig.Username, stdConfig.Password, stdConfig.Protocol,
		stdConfig.Address, stdConfig.Dbname, stdConfig.Params)
}
