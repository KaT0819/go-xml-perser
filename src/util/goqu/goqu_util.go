package goqu

import (
	"reflect"
	"strings"
	"time"

	"github.com/doug-martin/goqu/v9"

	"github.com/go-openapi/strfmt"
)

// goquを使用するときのgoqu.Recordを生成します。
// p : swaggerモデル構造体
// targetColumns : 更新対象キー
// excludeColumns : 更新対象外キー
func CreateGokuUpdateRecoed(p interface{}, targetColumns []string, excludeColumns []string) goqu.Record {

	updateColumns := goqu.Record{}
	val := reflect.ValueOf(p).Elem()

	for _, s := range targetColumns {

		for i := 0; i < val.NumField(); i++ {

			typeField := val.Type().Field(i)

			if strings.EqualFold(s, typeField.Name) && !contains(excludeColumns, s) {

				valueField := val.Field(i)  // 値
				t := valueField.Interface() // 型

				// fmt.Print(" key: ", typeField.Name)

				switch value := t.(type) {
				case int64:
					// println(" Value: ", value)
					updateColumns[typeField.Name] = value
				case string:
					// println(" Value: " + valueField.String())
					updateColumns[typeField.Name] = value
				case float64:
					// println(" Value: " + strconv.FormatFloat(value, 'f', -1, 64))
					// updateColumns[typeField.Name] = strconv.FormatFloat(value, 'f', -1, 64)
					updateColumns[typeField.Name] = value
				case bool:
					// println(" Value: ", value)
					updateColumns[typeField.Name] = value
				case strfmt.DateTime:
					// println(" Value: ", value.String())
					updateColumns[typeField.Name] = value
				}

				break
			}
		}
	}

	// 更新日時を追加
	updateColumns["UpdateDate"] = time.Now()

	return updateColumns
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
