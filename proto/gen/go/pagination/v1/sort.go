package pagination_v1

func (s SortOrder) Asc() bool {
	switch s {
	case SortOrder_SORT_ORDER_DESC:
		return false
	case SortOrder_SORT_ORDER_ASC:
		return true
	}
	return true
}
