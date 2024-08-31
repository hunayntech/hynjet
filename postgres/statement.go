package postgres

import "github.com/hunayntech/hynjet/v2/internal/jet"

// RawStatement creates new sql statements from raw query and optional map of named arguments
func RawStatement(rawQuery string, namedArguments ...RawArgs) Statement {
	return jet.RawStatement(Dialect, rawQuery, namedArguments...)
}
