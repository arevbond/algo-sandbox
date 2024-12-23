package findtheprefixcommonarrayoftwoarrays

func findThePrefixCommonArray(A []int, B []int) []int {
    counter := make(map[int]int)

    result := make([]int, len(A))

    for i := 0; i < len(A); i++ {
        counter[A[i]]++
        counter[B[i]]++

        var sum int
        if i > 0 {
            sum += result[i-1]
        }

        if cnt, ok := counter[A[i]]; ok && cnt == 2 {
            sum++
        } else if A[i] != B[i] {
            if cnt, ok := counter[B[i]]; ok && cnt == 2 {
                sum++
            }
        }


        result[i] = sum
    }
    return result
}
