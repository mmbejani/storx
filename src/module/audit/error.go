package audit

import "fmt"

type AuditResult[T any] struct {
	Value T
	Error error
}

type AuditErrorKind int

const (
	ErrConfiguration AuditErrorKind = iota
	ErrConfigNotLoaded
	ErrTarget
	ErrNotInitialized
	ErrAlreadyInitialized
	ErrStorageNotAvailable
	ErrSaveConfig
	ErrLoadConfig
	ErrSerialization
	ErrIO
	ErrJoin
)

type AuditError struct {
	Kind    AuditErrorKind
	Message string
	Err     error
}

func (e *AuditError) Error() string {
	switch e.Kind {
	case ErrConfiguration:
		return fmt.Sprintf("configuration error: %s", e.Message)
	case ErrConfigNotLoaded:
		return "config not loaded"
	case ErrTarget:
		return fmt.Sprintf("target error: %v", e.Err)
	case ErrNotInitialized:
		return fmt.Sprintf("system not initialized: %s", e.Message)
	case ErrAlreadyInitialized:
		return "system already initialized"
	case ErrStorageNotAvailable:
		return fmt.Sprintf("storage not available: %s", e.Message)
	case ErrSaveConfig:
		return fmt.Sprintf("failed to save configuration: %v", e.Err)
	case ErrLoadConfig:
		return fmt.Sprintf("failed to load configuration: %v", e.Err)
	case ErrSerialization:
		return fmt.Sprintf("serialization error: %v", e.Err)
	case ErrIO:
		return fmt.Sprintf("I/O error: %v", e.Err)
	case ErrJoin:
		return fmt.Sprintf("join error: %v", e.Err)
	default:
		return "unknown audit error"
	}
}
