package main

import (
	"practice_gogeneric/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	t.Run("employee", func(t *testing.T) {
		want := &model.Employee{}
		got := TableName(want)
		assert.IsType(t, want, got)
	})
	t.Run("department", func(t *testing.T) {
		want := &model.Department{}
		got := TableName(want)
		assert.IsType(t, want, got)
	})
}
