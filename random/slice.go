package random

// PickString in random postion of slice
func PickString(items []string) string {
	if len(items) == 0 {
		return ""
	}

	return items[rangeInt(0, len(items)-1)]
}
