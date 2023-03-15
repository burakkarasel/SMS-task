package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	db "github.com/burakkarasel/SMS-task/internal/db/mock"
	"github.com/burakkarasel/SMS-task/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestCreateStudent tests createStudent function
func TestCreateStudent(t *testing.T) {
	student := util.RandomStudent()
	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"fullName":   student.FullName,
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusCreated {
					t.Fatalf("Expected %d, got %d", http.StatusCreated, w.Code)
				}
			},
		},
		{
			name: "Invalid Input",
			body: gin.H{
				"fullName":   "",
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"fullName":   "YOLO",
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
	}

	for _, testCase := range testCases {
		store := db.MockStore{}

		server := newTestServer(t, store)
		w := httptest.NewRecorder()

		data, err := json.Marshal(testCase.body)
		if err != nil {
			t.Fatal(err)
		}

		url := "/api/v1/students"
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

// TestGetStudent tests getStudent function
func TestGetStudent(t *testing.T) {
	id := util.RandomInt(20, 1000)
	testCases := []struct {
		name          string
		id            int
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   id,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusOK {
					t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
				}
			},
		},
		{
			name: "Invalid Input",
			id:   -3,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name: "Internal Server Error",
			id:   17,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
		{
			name: "Not Found",
			id:   18,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusNotFound {
					t.Fatalf("Expected %d, got %d", http.StatusNotFound, w.Code)
				}
			},
		},
	}

	for _, testCase := range testCases {
		store := db.MockStore{}

		server := newTestServer(t, store)
		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/v1/students/%d", testCase.id)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

// TestListStudents tests listStudents
func TestListStudents(t *testing.T) {
	pageId := 1
	pageLimit := 5
	testCases := []struct {
		name          string
		query         string
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name:  "OK",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d", pageId, pageLimit),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusOK {
					t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
				}
			},
		},
		{
			name:  "Invalid Input",
			query: fmt.Sprintf("?pageId=-3&pageLimit=%d", pageLimit),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name:  "Internal Server Error",
			query: fmt.Sprintf("?pageId=%d&pageLimit=17", pageId),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
	}

	for _, testCase := range testCases {
		store := db.MockStore{}

		server := newTestServer(t, store)
		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/v1/students%s", testCase.query)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

// TestUpdateStudent tests updateStudent function
func TestUpdateStudent(t *testing.T) {
	student := util.RandomStudent()
	testCases := []struct {
		name          string
		id            int
		body          gin.H
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"fullName":   student.FullName,
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			id: student.Id,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusOK {
					t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
				}
			},
		},
		{
			name: "Invalid Body Input",
			body: gin.H{
				"fullName":   "",
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			id: student.Id,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name: "Invalid URI Input",
			body: gin.H{
				"fullName":   student.FullName,
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			id: -1,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name: "Internal Server Error",
			body: gin.H{
				"fullName":   "YOLO",
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			id: student.Id,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},

		{
			name: "Not Found",
			body: gin.H{
				"fullName":   "OH NO",
				"department": student.Department,
				"year":       student.Year,
				"email":      student.Email,
			},
			id: student.Id,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusNotFound {
					t.Fatalf("Expected %d, got %d", http.StatusNotFound, w.Code)
				}
			},
		},
	}

	for _, testCase := range testCases {
		store := db.MockStore{}

		server := newTestServer(t, store)
		w := httptest.NewRecorder()

		data, err := json.Marshal(testCase.body)
		if err != nil {
			t.Fatal(err)
		}

		url := fmt.Sprintf("/api/v1/students/%d", testCase.id)
		req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

// TestDeleteStudent tests deleteStudent function
func TestDeleteStudent(t *testing.T) {
	id := util.RandomInt(20, 1000)
	testCases := []struct {
		name          string
		id            int
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			id:   id,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusNoContent {
					t.Fatalf("Expected %d, got %d", http.StatusNoContent, w.Code)
				}
			},
		},
		{
			name: "Invalid Input",
			id:   -3,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name: "Internal Server Error",
			id:   17,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
		{
			name: "Not Found",
			id:   18,
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusNotFound {
					t.Fatalf("Expected %d, got %d", http.StatusNotFound, w.Code)
				}
			},
		},
	}

	for _, testCase := range testCases {
		store := db.MockStore{}

		server := newTestServer(t, store)
		w := httptest.NewRecorder()

		url := fmt.Sprintf("/api/v1/studentClasses/%d", testCase.id)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}
