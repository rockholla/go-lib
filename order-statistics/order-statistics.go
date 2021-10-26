package main

import (
  "fmt"
  "math"
)

// FindNthSmallest will find the nth smallest number in an un-ordered slice of integers
// Order Statistics, with Hoare-based partitioning and quickselect to minimize time complexity
func FindNthSmallest(collection []int, n int) (int, error) {
  var pivotIndex, pivot int
  // We'll rebuild our isolated collection into an internal-only variable in this function
  var col []int
  if n < 1 {
    return -1, fmt.Errorf("n should be a positive number")
  }
  // If we only have a single item slice/collection, just return that item
  if len(collection) == 1 {
    return collection[0], nil
  }
  // Remove duplicates, tracking existing entries in a map/hash table
  existing := make(map[int]bool)
  for _, item := range collection {
    if _, exists := existing[item]; exists == false {
      col = append(col, item)
      existing[item] = true
    }
  }
  if (n > len(col)) {
    return -1, fmt.Errorf("supplied n of %v is greater than the length of the unique (duplicates removed) collection: %v", n, len(col))
  }

  // We'll immediately reduce our nth to a zero-based index for ease of implementation
  n--
  // Set up our tracking low and high bounds for partioning at beginning and end of slice
  low := 0
  high := len(col) - 1
  var resetPivot = func() {
    // the pivot is a middle value where we'll place items less than it to the left, greater
    // than it to the right in a given partitioning iteration
    pivotIndex = int(math.Floor(float64((low + high) / 2)))
    pivot = col[pivotIndex]
  }
  resetPivot()
  for {
    // We're going to start at the beginning of the collection and move our way up
    // until we find a value that is greater than the pivot value
    for col[low] < pivot {
      low++
    }
    // We're going to start at the end of the collection and move our way down
    // until we find a value that is less than the pivot value
    for col[high] > pivot {
      high--
    }
    if high <= low {
      // this is where the "partitioning" is finished with one round, high = pivot index
      if n == high {
        return col[n], nil
      }
      if n < high {
        // if our n is less than the high/pivot index, we want to shift our high
        // down by one, reset to our baseline low, and begin again with partioning
        low = 0
        high--
      } else {
        // our n is greater than the high pivot index, we can move our low bound
        // up by one and reset our high to original, and begin again with partitioning
        low = high + 1
        high = len(col) - 1
      }
      resetPivot()
      continue
    }
    // we've identified a value in both the upper and lower sections to switch
    // moving each to the correct side of the pivot
    col[high], col[low] = col[low], col[high]
  }
}
