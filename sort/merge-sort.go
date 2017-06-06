package main
import "fmt"

func merge(array []int) ([]int) {
  if len(array) <= 1 {
    return array
  }
  middle := len(array) / 2

  left := merge(array[:middle])
  right := merge(array[middle:])

  return merge_segs(left, right)
}

func merge_segs(left []int, right []int) ([]int) {
  var result []int

  for len(left) > 0 && len(right) > 0 {
    if left[0] <= right[0] {
      result = append(result, left[0])
      left = left[1:]
    } else {
      result = append(result, right[0])
      right = right[1:]
    }
  }

  for len(right) > 0 {
    result = append(result, right[0])
    right = right[1:]
  }

  for len(left) > 0 {
    result = append(result, left[0])
    left = left[1:]
  }
  
  return result
}

/*
  Testing function
*/
func main() {
  array := []int{11, 6, 8, 1, 2, 2, 4, 9}

  fmt.Println("Unsorted array")
  fmt.Println(array)

  fmt.Println("\n Post Merge basic")
  array = merge(array)

  fmt.Println(array)
}