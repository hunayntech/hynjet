package metadata

// Table metadata struct
type Table struct {
	Name      string `sql:"primary_key"`
	Columns   []Column
	Relations []Relation
}

func (t Table) SetRelations(relations []Relation) {
	t.Relations = relations
}

// MutableColumns returns list of mutable columns for table
func (t Table) MutableColumns() []Column {
	var ret []Column

	for _, column := range t.Columns {
		if column.IsPrimaryKey || column.IsGenerated {
			continue
		}

		ret = append(ret, column)
	}

	return ret
}
