package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicMetrics(t *testing.T) {
	docs := map[FileName]string {
		"A":"The car is driven on the road",
		"B": "The truck is driven on the highway",
	}

	res := Parse(docs)
	t.Run("Noise ignored", func(t *testing.T) {
		exp := map[FileName]float64{
			"A": 0.0,
			"B": 0.0,
		}
		assert.Equal(t, exp, res["the"])
		assert.Equal(t, exp, res["is"])
		assert.Equal(t, exp, res["on"])
		assert.Equal(t, exp, res["driven"])
	})

	t.Run("unique words in doc A", func(t *testing.T) {
		assert.Equal(t, 0.0, res["truck"]["A"])
		assert.Equal(t, 0.0, res["highway"]["A"])
		assert.GreaterOrEqual(t, 0.2, res["car"]["A"])
		assert.GreaterOrEqual(t, 0.2, res["road"]["A"])
	})

	t.Run("unique words in doc B", func(t *testing.T) {
		assert.GreaterOrEqual(t, 0.2, res["truck"]["B"])
		assert.GreaterOrEqual(t, 0.2, res["highway"]["B"])
		assert.Equal(t, 0.0, res["car"]["B"])
		assert.Equal(t, 0.0, res["road"]["B"])
	})
}

func TestPunctuation(t *testing.T) {
	docs := map[FileName]string {
		"A": "foobar, FooBar    foobar. foobar foo-bar -,foo-bar.,",
	}

	res := Parse(docs)
	exp := TfIdf{
		"foobar": {
			"A": 0.0,
		},
	}
	assert.Len(t, res, 1)
	assert.Equal(t, exp, res)
}