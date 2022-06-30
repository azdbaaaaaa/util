package common

func (s SortOrder) Asc() bool {
	switch s {
	case SORT_ORDER_DESC:
		return false
	case SORT_ORDER_ASC:
		return true
	}
	return true
}
