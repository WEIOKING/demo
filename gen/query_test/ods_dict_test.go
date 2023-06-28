package query_test

import (
	"demo/gen/query"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"testing"
)

var instanceMysql *gorm.DB
var onceMysql sync.Once

func getDb() *gorm.DB {
	onceMysql.Do(func() {
		db, err := gorm.Open(mysql.Open("root:28W0f7e@@tcp(159.138.122.113:3306)/space_planckx_new?charset=utf8"))
		if err != nil {
			panic(fmt.Errorf("cannot establish db connection: %w", err))
		}
		instanceMysql = db
	})
	return instanceMysql
}

func TestQuery(t *testing.T) {
	db := getDb()
	dictTypes, err := query.Use(db).OdsDict.Find()
	if err != nil {
		t.Error("query error")
	}
	fmt.Println(len(dictTypes))
	for i := range dictTypes {
		marshal, err := json.Marshal(dictTypes[i])
		if err != nil {
			t.Error("json format error")
		}
		fmt.Println(string(marshal))
	}
}

func TestQuerier(t *testing.T) {
	db := getDb()
	result, err := query.Use(db).OdsDict.CountGroupByType()
	if err != nil {
		t.Error("query error")
	}
	for i := range result {
		marshal, err := json.Marshal(result[i])
		if err != nil {
			t.Error("json format error")
		}
		fmt.Println(string(marshal))
	}
}
