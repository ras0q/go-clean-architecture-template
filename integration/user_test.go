package integration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Ras96/go-clean-architecture-template/internal/interfaces/controller"
	"github.com/google/go-cmp/cmp"
)

func Test_userControllerImpl_GetUser(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		userID      string
		reqJSONBody string
		statusCode  int
		wantResBody *controller.GetUserResponse
	}{
		"200": {
			userID:      "1",
			reqJSONBody: "",
			statusCode:  http.StatusOK,
			wantResBody: &controller.GetUserResponse{
				ID:    "1",
				Name:  "Ras",
				Email: "ras@example.com",
			},
		},
		"400: invalid json": {
			userID:      "1",
			reqJSONBody: "{",
			statusCode:  http.StatusBadRequest,
			wantResBody: nil,
		},
		"404: user not found": {
			userID:      "notExistUserID",
			reqJSONBody: "",
			statusCode:  http.StatusNotFound,
			wantResBody: nil,
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res, err := doRequest(t, http.MethodGet, fmt.Sprintf("/api/users/%s", tt.userID), tt.reqJSONBody)
			if err != nil {
				t.Errorf("doRequest: %v", err)
			}
			if res.Code != tt.statusCode {
				t.Errorf("expected status code %d, but got %d", tt.statusCode, res.Code)
			}
			if tt.wantResBody != nil {
				var got controller.GetUserResponse
				if err := json.NewDecoder(res.Body).Decode(&got); err != nil {
					t.Errorf("json.NewDecoder.Decode: %v", err)
				}
				if diff := cmp.Diff(tt.wantResBody, &got); diff != "" {
					t.Errorf("unexpected response body: %s", diff)
				}
			}
		})
	}
}
