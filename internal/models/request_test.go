package models

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		startDate string
		endDate   string
		err       error
	}{
		{
			name:      "good dates",
			startDate: "2006-01-02",
			endDate:   "2006-01-05",
			err:       nil,
		},
		{
			name:      "end date before start date",
			startDate: "2006-01-06",
			endDate:   "2006-01-02",
			err:       ErrEndDateToEearly,
		},
		{
			name:      "start date in future",
			startDate: "2055-01-06",
			endDate:   "2006-01-02",
			err:       ErrDateInFuture,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request, _ = http.NewRequest("GET", "/", nil)
			q := c.Request.URL.Query()
			q.Add("start_date", tt.startDate)
			q.Add("end_date", tt.endDate)
			c.Request.URL.RawQuery = q.Encode()
			var req Request
			err := c.ShouldBind(&req)
			assert.NoError(t, err, "Should bind should not return error")
			err = req.Validate()
			assert.Equal(t, tt.err, err, "Expected error: %s, actual error: %s", tt.err, err)
		})
	}
}
