package guffers
import "strconv"


func strArr2IntArr(list []string) (res []int, done bool) {
  res = make([]int, len(list))

  for i, str := range list {
    intValue, err := strconv.Atoi(str)

    if err != nil {
      return nil, false
    }
  
    res[i] = intValue
  }
  return res, true
}

func strInArr(item string, list []string) bool {
  for _, i := range list {
    if i == item {
      return true
    }
  }

  return false
}
