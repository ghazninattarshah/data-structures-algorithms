package problems

// you can also use imports, for example:
import "strconv"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func BinaryGap(N int) int {
    // write your code in Go 1.4
    
    str := strconv.FormatInt(int64(N), 2)
    max := 0
    distance := 0
    inprogress := false
    for i := 0; i < len(str); i++ {
        
        if str[i] == '1' && !inprogress {
            inprogress = true
        } else if str[i] == '0' {
            distance++
        } else if str[i] == '1' && inprogress {
            inprogress = false
            if distance > max {
				max = distance
				//fmt.Printf("Distance=%d, Max=%d\n", distance, max)
            }
			distance = 0
			inprogress=true
        }
    }
    
    return max
}