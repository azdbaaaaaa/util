package pagination

func (s SortOrder) Asc() bool {
	switch s {
	case SortOrder_SortDesc:
		return false
	case SortOrder_SortAsc:
		return true
	}
	return true
}
