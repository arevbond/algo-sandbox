func destCity(paths [][]string) string {
	mp := make(map[string]struct{})    

	for _, path := range paths {
		mp[path[0]] = struct{}{}
	}

	for _, path := range paths {
		if _, ok := mp[path[1]]; !ok {
			return path[1]
		}
	}
	return ""
}
