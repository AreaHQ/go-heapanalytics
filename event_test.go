package heapanalytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEventSetsValuesCorrectly(t *testing.T) {
	expectedAppID := "testAppID"
	expectedIdentity := "testIdentity"
	expectecdEvent := "testEvent"
	expectedProperties := map[string]interface{}{"TestString": "This value", "TestNumber": 10}

	e := NewEvent(expectedAppID, expectedIdentity, expectecdEvent, expectedProperties)

	assert.Equal(t, expectedAppID, e.AppID)
	assert.Equal(t, expectedIdentity, e.Identity)
	assert.Equal(t, expectecdEvent, e.Event)
	assert.Equal(t, expectedProperties, e.Properties)
}
