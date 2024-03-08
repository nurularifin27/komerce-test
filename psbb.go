package main

import (
	"fmt"
  "strings"
  "strconv"
  "sort"
)

func main() {
  n := 8 // jumlah keluarga
  members := "2 3 4 4 2 1 3 1" // jumlah member per keluarga
  arrMembers := strings.Split(members," ")
  if len(arrMembers) != n {
    fmt.Println("Input must be equal with count of family")
    return
  }
  families := make([]int , n)
  for i , m := range arrMembers {
    sizes, err := strconv.Atoi(m)
    if err != nil {
      fmt.Println("Input not valid")
    }
    families[i] = sizes
  }

  bus := minBus(families)
  fmt.Printf("Minimum bus required is %d", bus)
}

func minBus(families []int) int {
  sort.Ints(families)
  bus := 0

  i , j := 0 , len(families) - 1

  for i <= j {
    if families[i] + families[j] <= 4 && i != j {
      i++
    }
    j--
    bus++
  }

  return bus
}