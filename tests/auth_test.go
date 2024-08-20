package test

import (
	"reflect"
	"testing"
	"net/http"
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestSplit(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://dummy.restapiexample.com/api/v1/employees", nil)
	req.Header.Add("Authorization", "ApiKey xxxxx")
    got, _ := auth.GetAPIKey(req.Header)
    want := "xxxxx"
    if !reflect.DeepEqual(want, got) {
         t.Fatalf("expected: %v, got: %v", want, got)
    }
}