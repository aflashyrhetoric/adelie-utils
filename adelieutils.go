package adelieutils

func ArrayContainsInt(a int, list []int) bool {
  for _, b := range list {
    if b == a {
      return true
    }
  }
  return false
}
