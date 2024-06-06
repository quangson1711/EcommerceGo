package user

import (
	"Ecommerce-Go/types"
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockUserStore struct {
}

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run("should fail if the user payload invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			FirstName: "dsa",
			LastName:  "dsa",
			Email:     "",
			Password:  "dsa",
		}

		marshaledPayload, _ := json.Marshal(payload)
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(marshaledPayload))

		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v", http.StatusBadRequest, status)
		}
	})
}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
