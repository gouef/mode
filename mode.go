package mode

import (
	"errors"
	"os"
)

const EnvMode = "ENV_MODE"

const (
	// DebugMode mode is debug.
	DebugMode = "debug"
	// ReleaseMode mode is release.
	ReleaseMode = "release"
	// TestMode mode is test.
	TestMode = "test"
)

type Mode struct {
	mode  string
	modes []string
}

// NewBasicMode create basic Mode with debug, release, test
// Example:
//
//	mode := mode.NewBaseMode()
func NewBasicMode() *Mode {
	return NewMode(nil)
}

// NewMode create Mode
// Example:
//
//	 modes := []string{"staging"}
//		mode := mode.NewMode(modes)
func NewMode(additionalModes []string) *Mode {
	mode := os.Getenv(EnvMode)
	modes := []string{
		DebugMode,
		TestMode,
		ReleaseMode,
	}
	modes = append(modes, additionalModes...)

	if !modeExists(modes, mode) {
		panic("mode from env doest not exists in available modes")
	}

	return &Mode{
		mode:  mode,
		modes: modes,
	}
}

// AddMode add next available mode
func (m *Mode) AddMode(mode string) *Mode {
	m.modes = append(m.modes, mode)

	return m
}

// IsMode check if mode is your specific
func (m *Mode) IsMode(mode string) (bool, error) {
	if !m.modeExists(mode) {
		return false, errors.New("mode doest not exists")
	}
	return m.mode == mode, nil
}

// EnableMode enable specific mode
func (m *Mode) EnableMode(mode string) (bool, error) {
	if !m.modeExists(mode) {
		return false, errors.New("mode doest not exists")
	}

	m.mode = mode
	return true, nil
}

// IsDebug check if mode is debug
func (m *Mode) IsDebug() (bool, error) {
	return m.IsMode(DebugMode)
}

// EnableDebug enable debug mode
func (m *Mode) EnableDebug() (bool, error) {
	return m.EnableMode(DebugMode)
}

// IsTest check if mode is test
func (m *Mode) IsTest() (bool, error) {
	return m.IsMode(TestMode)
}

// EnableTest enable test mode
func (m *Mode) EnableTest() (bool, error) {
	return m.EnableMode(TestMode)
}

// IsRelease check if mode is release
func (m *Mode) IsRelease() (bool, error) {
	return m.IsMode(ReleaseMode)
}

// EnableRelease enable release mode
func (m *Mode) EnableRelease() (bool, error) {
	return m.EnableMode(ReleaseMode)
}

// GetMode return current set mode
func (m *Mode) GetMode() string {
	return m.mode
}

// modeExists internal function for check if mode is inside available modes
func (m *Mode) modeExists(mode string) bool {
	return modeExists(m.modes, mode)
}

// modeExists internal function for check if mode is inside available modes
func modeExists(modes []string, mode string) bool {
	for _, availableMode := range modes {
		if availableMode == mode {
			return true
		}
	}

	return false
}