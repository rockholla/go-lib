package main

import (
  "testing"
)

func TestFindNthSmallestBasic(t *testing.T) {

  var col []int
  var err error
  var n, result, expectedResult int

  // 5, 10, 15, 20, 25, 30
  col = []int{30, 20, 10, 5, 15, 25}
  n = 5
  expectedResult = 25
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestBasic, col = %v, n = %v: %s", col, n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestBasic, col = %v, n = %v, expected %v, got %v", col, n, expectedResult, result)
  }

  // 0, 1, 5, 6, 8
  col = []int{1, 5, 0, 6, 8}
  n = 2
  expectedResult = 1
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestBasic, col = %v, n = %v: %s", col, n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestBasic, col = %v, n = %v, expected %v, got %v", col, n, expectedResult, result)
  }

  // 1, 2, 5, 6, 8
  col = []int{1, 5, 2, 6, 8}
  n = 4
  expectedResult = 6
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestBasic, col = %v, n = %v: %s", col, n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestBasic, col = %v, n = %v, expected %v, got %v", col, n, expectedResult, result)
  }

}

func TestFindNthSmallestFirst(t *testing.T) {
  // 1, 2, 5, 6, 8
  collection := []int{2, 5, 1, 6, 8}
  var n int
  var err error
  var result, expectedResult int
  n = 1
  expectedResult = 1
  result, err = FindNthSmallest(collection, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestBasic, n = %v: %s", n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestBasic, n = %v, expected %v, got %v", n, expectedResult, result)
  }
}

func TestFindNthSmallestLast(t *testing.T) {
  // 1, 2, 5, 6, 8
  collection := []int{2, 5, 1, 6, 8}
  var n int
  var err error
  var result, expectedResult int

  n = 5
  expectedResult = 8
  result, err = FindNthSmallest(collection, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestBasic, n = %v: %s", n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestBasic, n = %v, expected %v, got %v", n, expectedResult, result)
  }
}

func TestFindNthSmallestComplex(t *testing.T) {
  // 0, 1, 2, 3, 5, 12, 14, 16, 20, 21, 22, 29, 31, 34, 56
  collection := []int{20, 34, 16, 3, 29, 14, 56, 12, 1, 5, 21, 2, 0, 22, 31}
  var n int
  var err error
  var result, expectedResult int

  n = 4
  expectedResult = 3
  result, err = FindNthSmallest(collection, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestUnsortedDuplicates:, n = %v, %s", n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestUnsortedDuplicates, n = %v, expected %v, got %v", n, expectedResult, result)
  }

  n = 10
  expectedResult = 21
  result, err = FindNthSmallest(collection, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestUnsortedDuplicates:, n = %v, %s", n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestUnsortedDuplicates, n = %v, expected %v, got %v", n, expectedResult, result)
  }

  n = 1
  expectedResult = 0
  result, err = FindNthSmallest(collection, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestUnsortedDuplicates:, n = %v, %s", n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestUnsortedDuplicates, n = %v, expected %v, got %v", n, expectedResult, result)
  }

  n = 15
  expectedResult = 56
  result, err = FindNthSmallest(collection, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestUnsortedDuplicates:, n = %v, %s", n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestUnsortedDuplicates, n = %v, expected %v, got %v", n, expectedResult, result)
  }
}

func TestFindNthSmallestDuplicates(t *testing.T) {
  // Duplicates in order statistics mean that duplicates are ignored in order, so
  // the duplicated 14 in our col below becomes only the 7th smallest instead of both the 7th and 8th
  // 0, 1, 2, 3, 5, 12, 14, 16, 20, 21, 22, 29, 31, 34, 56
  col := []int{20, 34, 16, 3, 29, 14, 56, 12, 14, 1, 5, 21, 2, 0, 22, 31}
  var n int
  var err error
  var result, expectedResult int

  n = 4
  expectedResult = 3
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestDuplicates: col = %v, n = %v, %s", col, n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestDuplicates: col = %v, n = %v, expected %v, got %v", col, n, expectedResult, result)
  }

  // 1, 3, 4, 8, 12
  col = []int{3, 4, 12, 1, 8, 1}
  n = 2
  expectedResult = 3
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestDuplicates: col = %v, n = %v, %s", col, n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestDuplicates: col = %v, n = %v, expected %v, got %v", col, n, expectedResult, result)
  }

  n = 4
  expectedResult = 8
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestDuplicates: col = %v, n = %v, %s", col, n, err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestDuplicates: col = %v, n = %v, expected %v, got %v", col, n, expectedResult, result)
  }

}

func TestFindNthSmallestSingleItemSlice(t *testing.T) {
  col := []int{15}
  var n int
  var err error
  var result, expectedResult int
  n = 1
  expectedResult = 15
  result, err = FindNthSmallest(col, n)
  if err != nil {
    t.Errorf("Got unexpected error from TestFindNthSmallestSingleItemSlice: %s", err)
  }
  if result != expectedResult {
    t.Errorf("Got unexpected result from TestFindNthSmallestSingleItemSlice, n = %v, expected %v, got %v", n, expectedResult, result)
  }
}

func TestFindNthSmallestNOutofBounds(t *testing.T) {
  col := []int{1, 2, 5}
  n := 4
  _, err := FindNthSmallest(col, n)
  if err == nil {
    t.Errorf("Didn't get expected error in TestFindNthSmallestNOutofBounds")
  }
}

func TestFindNthSmallestDuplicatesNOutofBounds(t *testing.T) {
  col := []int{1, 1, 2, 5}
  n := 4
  _, err := FindNthSmallest(col, n)
  if err == nil {
    t.Errorf("Didn't get expected error in TestFindNthSmallestDuplicatesNOutofBounds")
  }
}

func TestFindNthSmallestEmptySlice(t *testing.T) {
  col := []int{}
  n := 1
  _, err := FindNthSmallest(col, n)
  if err == nil {
    t.Errorf("Didn't get expected error in TestFindNthSmallestEmptySlice")
  }
}

func TestFindNthSmallestZeroN(t *testing.T) {
  collection := []int{}
  n := 0
  _, err := FindNthSmallest(collection, n)
  if err == nil {
    t.Errorf("Didn't get expected error in TestFindNthSmallestZeroN")
  }
}

func TestFindNthSmallestNegativeN(t *testing.T) {
  collection := []int{}
  n := -1
  _, err := FindNthSmallest(collection, n)
  if err == nil {
    t.Errorf("Didn't get expected error in TestFindNthSmallestNegativeN")
  }
}
