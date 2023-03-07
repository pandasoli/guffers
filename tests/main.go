package main
import "guffers"


func main() {
  scr := guffers.Screen {}

  buff1 := guffers.Buffer {
    Props: guffers.Props {
      W: 60,
      H: 10,
      X: 4,
      Y: 2,
      BgCl: 7,
      Cl: 0,
    },
  }

  lorem := "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."

  buff1.Add(&lorem)
  scr.Add(&buff1)
}
