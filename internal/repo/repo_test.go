package repo

import (
	"event-diary/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestRepo(t *testing.T) {
	log := zap.NewExample().Sugar()
	c, err := NewClient(log)
	assert.NoError(t, err)
	testRecord := &models.Record{
		EventTime:  "eventTime",
		RecordTime: "recordTime",
		Reporter:   "reporter",
	}
	err = c.Create(testRecord)
	assert.NoError(t, err)

	records, err := c.AllRecords()
	t.Log(records[0])
	assert.NoError(t, err)
	assert.Equal(t, testRecord, records[0])

	err = c.Delete(records[0])
	assert.NoError(t, err)

	_, err = c.AllRecords()
	assert.NoError(t, err)

}
