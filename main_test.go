package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	// テストサーバーをセットアップ
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))
	defer ts.Close()

	// テストサーバーにリクエストを送信
	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	// レスポンスボディを読み込む
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// レスポンスボディを検証
	expected := "Hello, World!"
	if string(body) != expected {
		t.Errorf("Expected response body to be %s, got %s", expected, body)
	}
}
