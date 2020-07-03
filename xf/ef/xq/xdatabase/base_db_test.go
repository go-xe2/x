package xdatabase_test

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xe2/x/encoding/xjson"
	"github.com/go-xe2/x/xf/ef/xq/xdatabase"
	"github.com/go-xe2/x/xf/ef/xqi"
	"testing"
)

var config = xdatabase.NewMysqlConfigFromMap(map[string]interface{}{
	"host":        "127.0.1",
	"port":        3306,
	"user":        "root",
	"password":    "123456",
	"db":          "xqdb",
	"maxOpenCons": 0,
	"maxIdleCons": 0,
	"driver":      "mysql",
	"charset":     "utf8",
	"parseTime":   true,
})

func TestTBaseDb(t *testing.T) {
	conn, err := xdatabase.NewDbConn(config)
	if err != nil {
		t.Fatal(err)
	}
	szSql := "select user_id, name, sex, age, base_salary, " +
		"b_bool, n_int,n_bigint,n_tinyint, n_float,n_double,t_datetime,t_date,t_time,c_char,c_text,c_tinytext,c_longtext," +
		"c_blob,c_tinyblob,c_longblob,c_binary,c_varbinary " +
		" from tableA where user_id in (?,?,?,?,?)"
	db := xdatabase.NewDbWithConn(conn)
	ret, err := db.Query(xqi.MapBinder, szSql, 1, 2, 3, 4, 5)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := json.Marshal(ret)
	fmt.Println("result map:", string(bytes))

	ret1, err1 := db.Query(xqi.JsonBinder, szSql, 1, 2, 3, 4, 5)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("result json:", ret1)
	if _, ok := ret1.(xjson.JsonStr); ok {
		t.Log("ret1 is JsonStr=====>>>")
	}
	var TestData = map[string]interface{}{
		"total": 33,
		"rows":  ret1,
	}

	bytes, err = json.Marshal(TestData)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("json marshal:", string(bytes))

	ret2, err2 := db.Query(xqi.XmlBinder, szSql, 1, 2, 3, 4, 5)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("result xml:", ret2)
}
