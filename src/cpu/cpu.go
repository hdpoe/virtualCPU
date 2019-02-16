package cpu

import "mem"
import "fmt"

func BeginExecution(strtAdd uint16) (string) {
  memoryTable := mem.Memory{1,2,3,4,5,6,7,100,9,10}
//  len := len(memoryTable)
//  fmt.Println(len)
//  fmt.Println(uint8(strtAdd))
  data, err := mem.Fetch(uint8(strtAdd), memoryTable)
  if err == nil {
    fmt.Println("start address was", strtAdd)
    fmt.Println("data retrieved was", data)
  } else {
    fmt.Println("Error")
  }
  return "Yep"
}
