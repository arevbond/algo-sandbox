package tasks

func canVisitAllRooms2(rooms [][]int) bool {
	// room: keys
	paths := make(map[int][]int)

	for roomNumber, keys := range rooms {
		paths[roomNumber] = keys
	}

	seen := make(map[int]struct{})

	var dfs func(roomNumber int)
	
	dfs = func(roomNumber int) {
		seen[roomNumber] = struct{}{}

		nextRooms := paths[roomNumber]

		for _, nextRoomNumber := range nextRooms {
			if _, ok := seen[nextRoomNumber]; !ok {
				dfs(nextRoomNumber)
			}
		}

		return
	}

	dfs(0)

	for roomNumber := 0; roomNumber < len(rooms); roomNumber++ {
		if _, ok := seen[roomNumber]; !ok {
			return false
		}
	}

	return true

}
