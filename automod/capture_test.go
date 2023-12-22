package automod

import (
	"testing"

	"github.com/bluesky-social/indigo/automod/countstore"
	"github.com/stretchr/testify/assert"
)

func TestNoOpCaptureReplyRule(t *testing.T) {
	assert := assert.New(t)

	engine := EngineTestFixture()
	capture := MustLoadCapture("testdata/capture_atprotocom.json")
	assert.NoError(ProcessCaptureRules(&engine, capture))
	c, err := engine.GetCount("automod-quota", "report", countstore.PeriodDay)
	assert.NoError(err)
	assert.Equal(0, c)
	c, err = engine.GetCount("automod-quota", "takedown", countstore.PeriodDay)
	assert.NoError(err)
	assert.Equal(0, c)
}