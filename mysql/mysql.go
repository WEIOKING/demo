package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
}

func mysql1() {
	db, err := sql.Open("mysql", "root:28W0f7e@@tcp(159.138.122.113:3306)/space_planckx_new?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM ods_dict_type LIMIT 10")
	if err != nil {
		log.Fatal()
	}
	defer rows.Close()
	for rows.Next() {
		var odsDictType OdsDictType
		err := rows.Scan(&odsDictType.DictTypeId, &odsDictType.TypeCode, &odsDictType.TypeName, &odsDictType.TypeDes, &odsDictType.IsValid)
		if err != nil {
			log.Fatal(err)
		}
		marshal, err := json.Marshal(&odsDictType)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(marshal)
		fmt.Println(string(marshal))
	}
}

type OdsDictType struct {
	DictTypeId uint64 `json:"dictTypeId"`
	TypeCode   string `json:"typeCode"`
	TypeName   string `json:"typeName"`
	TypeDes    string `json:"typeDes"`
	IsValid    int    `json:"isValid"`
}
