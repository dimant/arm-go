package util

func Find(arr []interface{}, predicate func(interface{}) bool) interface{} {
	for _, value := range arr {
		if predicate(value) {
			return value
		}
	}

	return nil
}