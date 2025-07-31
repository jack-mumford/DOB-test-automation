package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
 // XOR all drone IDs - pairs will cancel out, leaving only the unique one
 result := 0
 for _, id := range droneIds {
  result ^= id
 }
 return result
}