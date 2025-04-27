package System

import (
	"VGBackendFramework/Database"
	"encoding/json"
	"fmt"
	"github.com/TylerBrock/colorjson"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
)

type TestUnit struct {
}

const TestUnitAuthorizationBearerToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJkYXRhIjoiMSIsImV4cCI6MTU4ODg4MjAxMTczNzIxOTAwMSwiaXNzIjoidnVlZ2hvc3QifQ.DuWl1Yd7sF1kCbFBgO6RyWU4YS7TfPJmYY4BXRg2jlw"

func (t *TestUnit) HttpAuthorizationBearer(httpRequest *http.Request, withSession ...bool) {
	if len(withSession) > 0 && withSession[0] {
		httpRequest.Header.Add("Authorization", TestUnitAuthorizationBearerToken)
	}
}

func (t *TestUnit) HttpPostApiController(apiController interface{}, method string, values url.Values, withSession ...bool) ApiOutputStore {
	var finalOutput ApiOutputStore
	httpTest := httptest.NewServer(http.HandlerFunc(func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		t.HttpAuthorizationBearer(httpRequest, withSession...)
		Run(apiController, method, httpResponseWriter, httpRequest)
	}))
	defer httpTest.Close()

	res, err := http.PostForm(httpTest.URL, values)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var obj map[string]interface{}
	_ = json.Unmarshal([]byte(greeting), &obj)
	_ = json.Unmarshal([]byte(greeting), &finalOutput)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
	return finalOutput
}

func (t *TestUnit) HttpGetApiController(apiController interface{}, method string, values interface{}) ApiOutputStore {
	var finalOutput ApiOutputStore
	httpTest := httptest.NewServer(http.HandlerFunc(func(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
		Run(apiController, method, httpResponseWriter, httpRequest)
	}))
	defer httpTest.Close()

	res, err := http.Get(httpTest.URL)
	if err != nil {
		log.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var obj map[string]interface{}
	_ = json.Unmarshal([]byte(greeting), &obj)
	_ = json.Unmarshal([]byte(greeting), &finalOutput)

	// Make a custom formatter with indent set
	f := colorjson.NewFormatter()
	f.Indent = 4

	// Marshall the Colorized JSON
	s, _ := f.Marshal(obj)
	fmt.Println(string(s))
	return finalOutput
}

func (t *TestUnit) DBExecute(query string) {
	t.DB(func(db Database.DB) {
		_, _ = db.Execute(query)
	})
}
func (t *TestUnit) IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	return reflect.DeepEqual(value, reflect.Zero(reflect.TypeOf(value)).Interface())
}

func (t *TestUnit) Contains(value string, substr string) (isContain bool) {
	return strings.Contains(value, substr)
}

func (t *TestUnit) SomeError() error {
	err := TestSomeError{}
	return err
}

func (t *TestUnit) DB(run func(db Database.DB)) {
	db := Database.NewDB(Database.NewPostgresqlDrive())
	db.Connect()
	defer db.Close()
	run(db)
}

type TestSomeError struct {
	error
}

func (te TestSomeError) Error() string {
	return "This just demo dummy error handler"
}
