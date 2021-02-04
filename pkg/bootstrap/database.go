package bootstrap

import (
	"fmt"
	"github.com/ecnuvj/vhoj_db/pkg/dao/datasource"
)

func InitDatabase() {
	err := datasource.ConnectDB("../../../config/mysql.yaml")
	if err != nil {
		panic(fmt.Sprintf("connect db error: %v", err))
	}
}
