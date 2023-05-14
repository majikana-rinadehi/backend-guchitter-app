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

// 可変長引数を、DialOptの方法で記載する
// cf. https://wp.jmuk.org/2018/01/06/go%E3%81%AE%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%8A%E3%83%AB%E5%BC%95%E6%95%B0/
type ginOptions struct {
	withParam bool
	params    []gin.Param

	withBody bool
	model    any

	withQuery bool
	queryList []gin.Param
}

type GinOption func(o *ginOptions)

// gin用のRecorder, Contextを作成
func SetupGinContext(options ...GinOption) (*httptest.ResponseRecorder, *gin.Context) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	opt := ginOptions{}
	// setup options
	for _, o := range options {
		o(&opt)
	}

	if opt.withParam {
		// リクエストパラメータを追加する
		for _, p := range opt.params {
			c.Params = append(
				c.Params, p,
			)
		}
	}

	if opt.withBody {
		// リクエストボディを追加する
		jsonValue, _ := json.Marshal(opt.model)
		reqBody := bytes.NewBuffer(jsonValue)
		req, _ := http.NewRequest(
			"POST",
			"",
			reqBody,
		)
		c.Request = req
	}

	if opt.withQuery {
		// リクエストクエリを設定する
		req, _ := http.NewRequest(
			http.MethodPost,
			"",
			nil,
		)
		q := req.URL.Query()
		for _, query := range opt.queryList {
			q.Add(query.Key, query.Value)
		}
		req.URL.RawQuery = q.Encode()
		c.Request = req
	}
	return w, c
}

func WithParam(params []gin.Param) GinOption {
	return func(opt *ginOptions) {
		opt.withParam = true
		opt.params = params
	}
}

func WithBody(model any) GinOption {
	return func(opt *ginOptions) {
		opt.withBody = true
		opt.model = model
	}
}

func WithQuery(queryList []gin.Param) GinOption {
	return func(opt *ginOptions) {
		opt.withQuery = true
		opt.queryList = queryList
	}
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
