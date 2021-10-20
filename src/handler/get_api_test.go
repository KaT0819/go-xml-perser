package handler

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/google/go-cmp/cmp"

	"go-xml-parser/src/gen/restapi/operations"
	"go-xml-parser/src/testUtil"
)

const (
	getTestUrl = "/test"
)

// 正常系（検索条件なし）
func TestGetApiNoParam(t *testing.T) {
	// spanner接続用
	dbclient := testUtil.DbClient()
	defer dbclient.Close()

	// リクエストパラメータ設定
	params := operations.NewGetTestParams()
	request := httptest.NewRequest("GET", getTestUrl, nil)
	params.HTTPRequest = request
	t.Log(params)

	// 呼び出し
	r := GetTest(params)
	w := httptest.NewRecorder()
	r.WriteResponse(w, runtime.JSONProducer())

	// レスポンスをStructに変換
	var ret string
	err := json.Unmarshal(w.Body.Bytes(), &ret)
	if err != nil {
		t.Error(err)
	}

	// 期待値の定義
	extract := "テストです"
	// t.Log(reflect.DeepEqual(ret[0], extract))

	// 差分確認
	if diff := cmp.Diff(ret, extract, nil); diff != "" {
		t.Errorf("Value is mismatch (-結果 +期待値):\n%s", diff)
	}
}
