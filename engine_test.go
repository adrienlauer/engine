package engine

import (
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestComponentDownload(t *testing.T) {
	logger := log.New(os.Stdout, "TEST: ", log.Ldate|log.Ltime)
	lagoon, err := Create(logger, "./testdata/work", "./testdata/git/main", "1.2.3")
	assert.Nil(t, err)
	assert.NotNil(t, lagoon)
}
