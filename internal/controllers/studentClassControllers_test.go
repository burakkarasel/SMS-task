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

// TestCreateStudentClass tests createStudentClass function
func TestCreateStudentClass(t *testing.T) {
	sc := util.RandomStudentClassResponse()
	testCases := []struct {
		name          string
		body          gin.H
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"studentId": sc.Student.Id,
				"classId":   sc.Class.Id,
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
				"studentId": -1,
				"classId":   sc.Class.Id,
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
				"studentId": 17,
				"classId":   sc.Class.Id,
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
		{
			name: "Conflict",
			body: gin.H{
				"studentId": 18,
				"classId":   sc.Class.Id,
			},
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusConflict {
					t.Fatalf("Expected %d, got %d", http.StatusConflict, w.Code)
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

		url := "/api/v1/studentClasses"
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

// TestGetStudentClass tests getStudentClass function
func TestGetStudentClass(t *testing.T) {
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

		url := fmt.Sprintf("/api/v1/studentClasses/%d", testCase.id)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

// TestListStudentClasses tests listStudentClasses function
func TestListStudentClasses(t *testing.T) {
	pageId := 1
	pageLimit := 5
	studentId := 5
	classId := 5
	testCases := []struct {
		name          string
		query         string
		checkResponse func(t *testing.T, w *httptest.ResponseRecorder)
	}{
		{
			name:  "Valid StudentId",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&studentId=%d&classId=0", pageId, pageLimit, studentId),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusOK {
					t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
				}
			},
		},
		{
			name:  "Valid ClassId",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&studentId=0&classId=%d", pageId, pageLimit, classId),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusOK {
					t.Fatalf("Expected %d, got %d", http.StatusOK, w.Code)
				}
			},
		},
		{
			name:  "Invalid Input",
			query: fmt.Sprintf("?pageId=-3&pageLimit=%d&studentId=%d&classId=0", pageLimit, studentId),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name:  "Both 0 Input",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&studentId=0&classId=0", pageId, pageLimit),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name:  "Both Positive Input",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&studentId=%d&classId=%d", pageId, pageLimit, studentId, classId),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusBadRequest {
					t.Fatalf("Expected %d, got %d", http.StatusBadRequest, w.Code)
				}
			},
		},
		{
			name:  "Student Internal Server Error",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&studentId=%d", pageId, pageLimit, 17),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
		{
			name:  "Student Not Found",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&studentId=%d", pageId, pageLimit, 18),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusNotFound {
					t.Fatalf("Expected %d, got %d", http.StatusNotFound, w.Code)
				}
			},
		},
		{
			name:  "Class Internal Server Error",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&classId=%d", pageId, pageLimit, 17),
			checkResponse: func(t *testing.T, w *httptest.ResponseRecorder) {
				if w.Code != http.StatusInternalServerError {
					t.Fatalf("Expected %d, got %d", http.StatusInternalServerError, w.Code)
				}
			},
		},
		{
			name:  "Class Not Found",
			query: fmt.Sprintf("?pageId=%d&pageLimit=%d&classId=%d", pageId, pageLimit, 18),
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

		url := fmt.Sprintf("/api/v1/studentClasses%s", testCase.query)
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}

func TestDeleteStudentClass(t *testing.T) {
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

		url := fmt.Sprintf("/api/v1/students/%d", testCase.id)
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			t.Fatal(err)
		}

		server.router.ServeHTTP(w, req)

		testCase.checkResponse(t, w)
	}
}
