package main

 import "cpu"

func main() {
  var start uint16 = 0x07
  cpu.BeginExecution(start)
}
