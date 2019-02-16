package mem

 import "errors"

type Memory []uint16

func Fetch(address uint8, mem Memory) (data uint16, err error) {
  if(uint8(len(mem)) < address) {
    err = errors.New("-1")
    data = 0
  } else {
    data = mem[address]
  }
  return
}
