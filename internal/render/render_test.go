package render

import (
	"bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	// get the request with session context
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	// get context from request
	ctx := r.Context()
	// load session header into context
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	// put context back into request
	r = r.WithContext(ctx)

	return r, nil
}
