package testutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path"
	"runtime"
	"testing"

	config "github.com/backend-guchitter-app/config"
	"github.com/gin-gonic/gin"
	"github.com/go-testfixtures/testfixtures/v3"
)

const fixturesDirRelativePathFormat = "%s/../../infrastructure/persistence/fixtures"

// gin用のRecorder, Contextを作成
func SetupGinContext() (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return w, c
}

// gin用のRecorder, Contextを作成(リクエストパラメータ)
func SetupGinContextWithParam(params []gin.Param) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// リクエストパラメータを追加する
	for _, p := range params {
		c.Params = append(
			c.Params, p,
		)
	}
	return w, c
}

// gin用のRecorder, Contextを作成(リクエストボディ)
func SetupGinContextWithBody(model any) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// リクエストボディを追加する
	jsonValue, _ := json.Marshal(model)
	reqBody := bytes.NewBuffer(jsonValue)
	req, _ := http.NewRequest(
		"POST",
		"",
		reqBody,
	)
	c.Request = req
	return w, c
}

// gin用のRecorder, Contextを作成(リクエストクエリ)
func SetupGinContextWithQuery(queryList []gin.Param) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	// リクエストクエリを設定する
	req, _ := http.NewRequest(
		http.MethodPost,
		"",
		nil,
	)
	q := req.URL.Query()
	for _, query := range queryList {
		q.Add(query.Key, query.Value)
	}
	req.URL.RawQuery = q.Encode()
	c.Request = req
	return w, c
}

// SetTestEnv sets env values of db conn for repository unit test
func SetTestEnv(t *testing.T) {
	t.Setenv("guchitter_USER", "root")
	t.Setenv("guchitter_PASS", "root")
	t.Setenv("guchitter_DBNAME", "guchitter_test")
	t.Setenv("guchitter_HOST", "localhost")
	t.Setenv("guchitter_PORT", "3307")
}

// SetupFixtures setup conn to test DB and insert test data into it
func SetupFixtures() {

	// *gorm.DB.DB() で *sql.DB がしゅとくできたのか...
	db, _ := config.ConnectTest().DB()

	_, pwd, _, _ := runtime.Caller(0)
	dir := fmt.Sprintf(fixturesDirRelativePathFormat, path.Dir(pwd))

	fmt.Println("dir:", dir)

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),     // You database connection
		testfixtures.Dialect("mysql"), // Available: "postgresql", "timescaledb", "mysql", "mariadb", "sqlite" and "sqlserver"
		testfixtures.Directory(dir),   // The directory containing the YAML files
	)
	if err != nil {
		panic(err)
	}

	fixtures.Load()
}

// ToString converts structure to string
func ToString(v interface{}) string {
	jsonByte, _ := json.Marshal(v)
	jsonStr := bytes.NewBuffer(jsonByte).String()
	return jsonStr
}
