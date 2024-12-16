package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gouef/mode" // Nahraď správnou cestou k balíčku "mode"
)

func TestNewBasicMode_DefaultMode(t *testing.T) {
	os.Unsetenv(mode.EnvMode)

	m, err := mode.NewBasicMode()

	require.Error(t, err, "mode from env doest not exists in available modes, used \"\", \"debug\" will be used.\n")
	assert.Equal(t, mode.DebugMode, m.GetMode())
	assert.ElementsMatch(t, []string{mode.DebugMode, mode.TestMode, mode.ReleaseMode}, m.Modes())
}

func TestNewBasicMode_InvalidEnvMode(t *testing.T) {
	os.Setenv(mode.EnvMode, "invalid")

	m, err := mode.NewBasicMode()

	assert.Error(t, err)
	assert.Equal(t, mode.DebugMode, m.GetMode())
	assert.Contains(t, err.Error(), "mode from env doest not exists in available modes")
}

func TestNewMode_ValidEnvMode(t *testing.T) {
	os.Setenv(mode.EnvMode, mode.TestMode)

	m, err := mode.NewMode([]string{"staging"})

	require.NoError(t, err)
	assert.Equal(t, mode.TestMode, m.GetMode())
	assert.Contains(t, m.Modes(), "staging")
}

func TestMode_AddMode(t *testing.T) {
	os.Unsetenv(mode.EnvMode)
	m, _ := mode.NewBasicMode()

	m.AddMode("staging")

	assert.Contains(t, m.Modes(), "staging")
}

func TestMode_EnableMode(t *testing.T) {
	os.Unsetenv(mode.EnvMode)
	m, _ := mode.NewBasicMode()

	success, err := m.EnableMode(mode.TestMode)

	require.NoError(t, err)
	assert.True(t, success)
	assert.Equal(t, mode.TestMode, m.GetMode())
}

func TestMode_EnableDebug(t *testing.T) {
	os.Unsetenv(mode.EnvMode)
	m, _ := mode.NewBasicMode()

	success, err := m.EnableDebug()

	require.NoError(t, err)
	assert.True(t, success)
	assert.Equal(t, mode.DebugMode, m.GetMode())
}

func TestMode_EnableTest(t *testing.T) {
	os.Unsetenv(mode.EnvMode)
	m, _ := mode.NewBasicMode()

	success, err := m.EnableTest()

	require.NoError(t, err)
	assert.True(t, success)
	assert.Equal(t, mode.TestMode, m.GetMode())
}

func TestMode_EnableRelease(t *testing.T) {
	os.Unsetenv(mode.EnvMode)
	m, _ := mode.NewBasicMode()

	success, err := m.EnableRelease()

	require.NoError(t, err)
	assert.True(t, success)
	assert.Equal(t, mode.ReleaseMode, m.GetMode())
}

func TestMode_EnableMode_Invalid(t *testing.T) {
	os.Unsetenv(mode.EnvMode)
	m, _ := mode.NewBasicMode()

	success, err := m.EnableMode("invalid")

	assert.Error(t, err)
	assert.False(t, success)
	assert.Contains(t, err.Error(), "mode doest not exists")
}

func TestMode_IsMode(t *testing.T) {
	os.Setenv(mode.EnvMode, mode.DebugMode)
	m, _ := mode.NewBasicMode()

	isMode, err := m.IsMode(mode.DebugMode)

	require.NoError(t, err)
	assert.True(t, isMode)

	isMode, err = m.IsMode("invalid")
	assert.Error(t, err)
	assert.False(t, isMode)
}

func TestMode_IsDebug(t *testing.T) {
	os.Setenv(mode.EnvMode, mode.DebugMode)
	m, _ := mode.NewBasicMode()

	isDebug, err := m.IsDebug()

	require.NoError(t, err)
	assert.True(t, isDebug)
}

func TestMode_IsTest(t *testing.T) {
	os.Setenv(mode.EnvMode, mode.TestMode)
	m, _ := mode.NewBasicMode()

	isTest, err := m.IsTest()

	require.NoError(t, err)
	assert.True(t, isTest)
}

func TestMode_IsRelease(t *testing.T) {
	os.Setenv(mode.EnvMode, mode.ReleaseMode)
	m, _ := mode.NewBasicMode()

	isRelease, err := m.IsRelease()

	require.NoError(t, err)
	assert.True(t, isRelease)
}
