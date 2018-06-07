package engine

import (
	"log"
	"os"
	"testing"

	"github.com/lagoon-platform/model"
	"github.com/stretchr/testify/assert"
)

/*
func TestEngineRemoteNoTag(t *testing.T) {
	os.RemoveAll("testdata/work")
	engine1, e := Create(log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime), "testdata/work", "https://github.com/lagoon-platform/demomode/l", "")
	assertOnlyWarnings(t, e)
	assert.NotNil(t, engine1)

	engine2, e := Create(log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime), "testdata/work", "https://github.com/lagoon-platform/demomodel/", "")
	assertOnlyWarnings(t, e)
	assert.NotNil(t, engine2)
}
*/
func TestEngineLocalWithTag(t *testing.T) {
	os.RemoveAll("testdata/work")
	engine, e := Create(log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime), "testdata/work", "testdata/sample", "v1.0.0")
	assertOnlyWarnings(t, e)
	engine.ComponentManager().Ensure()
}

func TestEngineLocalNoTag(t *testing.T) {
	os.RemoveAll("testdata/work")
	engine, e := Create(log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime), "testdata/work", "testdata/sample", "")
	assertOnlyWarnings(t, e)
	engine.ComponentManager().Ensure()
}

func assertOnlyWarnings(t *testing.T, e error) {
	if e != nil {
		switch e.(type) {
		case model.ValidationErrors:
			assert.False(t, e.(model.ValidationErrors).HasErrors())
		default:
			assert.Nil(t, e)
		}
	}
}
