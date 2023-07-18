package query_test

import (
	"demo/gen/query"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
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

func TestQuerier1(t *testing.T) {
	db := getDb()
	var ids = []int64{1529710641459326980, 1529710641459326979}
	u := query.Use(db).OdsDict
	result, err := u.Select().Where(u.DictID.In(ids...)).Find()
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

func Test(t *testing.T) {
	var body = "{\"type\":\"任务类型，配置任务类型code\",\"typeName\":1,\"typeLogo\":\"任务类型展示logo\",\"eventType\":\"类型下具体的事件类型code\",\"eventName\":\"类型下具体的事件类型名称\",\"eventLogo\":\"类型下具体的事件类型logo\",\"showTemplate\":{\"titleTemplate\":\"任务展示标题模板 例如关注Twitter任务：可以定义展示内容：Follow @{UserName} on twitter。其中userName为接口返回参数，具体定义方式可和前端确认\",\"description\":\"任务描述,可配置任务描述信息，指引用户完成任务\",\"buttons\":[{\"type\":\"按钮类型 跳转链接按钮、验证按钮等\",\"name\":\"按钮名称\"}]},\"params\":[{\"title\":\"参数展示标题\",\"description\":\"参数描述，指引项目方正确填写参数\",\"type\":\"输入框类型，用户输入、下拉、开关、时间选择等\",\"enum\":\"枚举类型，下拉选择时使用\",\"backDesc\":\"输入框背景描述\",\"paramName\":\"接口传参字段名称\",\"paramType\":\"参数值类型\",\"verifyRules\":[{\"type\":\"校验规则 如：maxLength\",\"value\":500},{\"type\":\"maxValue\",\"value\":10}]}]}"
	var data map[string]any
	_ = json.Unmarshal([]byte(body), &data)
	fmt.Println(getFiled(data, "showTemplate.buttons"))
}

func getFiled(data map[string]any, filed string) string {
	before, after, found := strings.Cut(filed, ".")
	if found {
		a := data[before]
		v, ok := a.(map[string]any)
		if ok {
			return getFiled(v, after)
		} else {
			return ""
		}
	} else {
		marshal, err := json.Marshal(data[filed])
		if err != nil {
			return ""
		} else {
			return string(marshal)
		}
	}
}
