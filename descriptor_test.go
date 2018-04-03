package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLabelCreate(t *testing.T) {
	assert.Equal(t, []string{"label1", "label2"}, createLabels("label1", "label2").AsStrings())
}

func TestLabelContains(t *testing.T) {
	f := createLabels("label1", "label2")
	assert.Equal(t, true, f.Contains("label1"))
	assert.Equal(t, true, f.Contains("label2"))
	assert.Equal(t, false, f.Contains("label3"))
}
