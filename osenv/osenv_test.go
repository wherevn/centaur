package osenv

import (
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_OK(t *testing.T) {
	var (
		key      = "TEST_BOOL_VALUE"
		expected = true
	)
	_ = os.Setenv(key, "true")

	result, err := Bool(key)
	{
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}

}

func TestInt_OK(t *testing.T) {
	var (
		key      = "TEST_INT_VALUE"
		expected = 20
	)
	_ = os.Setenv(key, strconv.Itoa(expected))

	result, err := Int(key)
	{
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestInt32_OK(t *testing.T) {
	var (
		key      = "TEST_INT32_VALUE"
		expected = int32(69)
	)
	_ = os.Setenv(key, strconv.Itoa(int(expected)))

	result, err := Int32(key)
	{
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestInt64_OK(t *testing.T) {
	var (
		key      = "TEST_INT64_VALUE"
		expected = int64(1323)
	)
	_ = os.Setenv(key, "1323")

	result, err := Int64(key)
	{
		assert.NoError(t, err)
		assert.Equal(t, expected, result)
	}
}

func TestString_OK(t *testing.T) {
	var (
		key      = "TEST_INT64_VALUE"
		expected = "test string"
	)
	_ = os.Setenv(key, expected)

	result := String(key)
	{
		assert.Equal(t, expected, result)
	}
}
