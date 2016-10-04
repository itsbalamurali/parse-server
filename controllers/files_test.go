package controllers

import (
	"testing"
	"net/http/httptest"
	"fmt"
	"net/http"
	"bytes"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"github.com/goamz/goamz/s3/s3test"
)

func startServer() (*s3test.Server, *s3.Bucket) {
	testServer, _ := s3test.NewServer(new(s3test.Config))
	auth := aws.Auth{"abc", "123", ""}
	server := s3.New(auth, aws.Region{Name: "faux-region-1", S3Endpoint: testServer.URL(), S3LocationConstraint: true}) // Needs true for s3test
	b := server.Bucket("foobar.com")
	b.PutBucket(s3.Private)
	return testServer, b
}

func TestPut(t *testing.T) {
	s, b := startServer()
	defer s.Quit()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/files/foo.txt", bytes.NewBufferString("foobar"))
	r.Header.Set("Content-Type", "text/plain")
	upload.Service(b).ServeHTTP(w, r)

	expCode := 200
	if w.Code != expCode {
		t.Fatalf("Response status expected: %#v, got: %#v", expCode, w.Code)
	}

	expLoc := b.URL("/foo.txt")
	if w.Header().Get("Location") != expLoc {
		t.Fatalf("Response Location header expected: %#v, got: %#v", expLoc, w.Header().Get("Location"))
	}

	expBody := fmt.Sprintf(`{"url":"%s"}`, expLoc)
	if w.Body.String() != expBody {
		t.Fatalf("Response body expected: %#v, got: %#v", expBody, w.Body.String())
	}

	data, _ := b.Get("/foo.txt")
	if string(data) != "foobar" {
		t.Fatalf("Stored file content expected: %#v, got: %#v", "foobar", string(data))
	}
}

func TestPutBodyNil(t *testing.T) {
	s, b := startServer()
	defer s.Quit()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("PUT", "/files/foo.txt", nil)
	upload.Service(b).ServeHTTP(w, r)

	expCode := 400
	if w.Code != expCode {
		t.Fatalf("Response status expected: %#v, got: %#v", expCode, w.Code)
	}

	expBody := "Body must be set\n"
	if w.Body.String() != expBody {
		t.Fatalf("Response body expected: %#v, got: %#v", expBody, w.Body.String())
	}
}

func TestGet(t *testing.T) {
	s, b := startServer()
	defer s.Quit()

	b.Put("/foo.txt", []byte("foobar"), "text/plain", s3.PublicRead)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/files/foo.txt", nil)
	upload.Service(b).ServeHTTP(w, r)

	expCode := 200
	if w.Code != expCode {
		t.Fatalf("Response status expected: %#v, got: %#v", expCode, w.Code)
	}

	expBody := fmt.Sprintf(`{"url":"%s"}`, b.URL("/foo.txt"))
	if w.Body.String() != expBody {
		t.Fatalf("Response body expected: %#v, got: %#v", expBody, w.Body.String())
	}
}

func TestDelete(t *testing.T) {
	s, b := startServer()
	defer s.Quit()

	b.Put("/foo.txt", []byte("foobar"), "text/plain", s3.PublicRead)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("DELETE", "/files/foo.txt", nil)
	upload.Service(b).ServeHTTP(w, r)

	expCode := 204
	if w.Code != expCode {
		t.Fatalf("Response status expected: %#v, got: %#v", expCode, w.Code)
	}

	res, _ := b.Head("/foo.txt")
	if res != nil {
		t.Fatalf("File /foo.txt has not been deleted from S3.")
	}
}