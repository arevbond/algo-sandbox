package stringcompression

import "strconv"

/*
Input: chars = ["a","a","b","b","c","c","c"]
Output: Return 6, and the first 6 characters of the input array should be: ["a","2","b","2","c","3"]
Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".
*/

//   ["a","2","b","b","c","c","c"]
//l:           ^
//r:                  ^
func compress(chars []byte) int {
    l := 0 // index to write in chars 

    for r := 0; r < len(chars); {
        i := r + 1
        for ; i < len(chars) && chars[i] == chars[i-1]; i++ {}

        if i - r > 1 {
            for _, n := range strconv.Itoa(i-r) {
                l++
                chars[l] = byte(n)
            }
        }

        r = i
        l = i
    }

    return l
}
