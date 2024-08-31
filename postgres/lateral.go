package postgres

import "github.com/hunayntech/hynjet/v2/internal/jet"

// LATERAL derived tables constructor from select statement
func LATERAL(selectStmt SelectStatement) lateralImpl {
	return lateralImpl{
		selectStmt: selectStmt,
	}
}

type lateralImpl struct {
	selectStmt SelectStatement
}

func (l lateralImpl) AS(alias string) SelectTable {
	subQuery := &selectTableImpl{
		SelectTable: jet.NewLateral(l.selectStmt, alias),
	}

	subQuery.readableTableInterfaceImpl.parent = subQuery

	return subQuery
}
