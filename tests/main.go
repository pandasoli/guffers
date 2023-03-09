package main
import "guffers"


func main() {
  scr := guffers.Screen {}

  lorem1 := "Lorem ipsum dolor sit amet, qui minim labore adipisicing minim sint cillum sint consectetur cupidatat."
  lorem2 := "Lorem ipsum dolor sit amet, officia excepteur ex fugiat reprehenderit enim labore culpa sint ad nisi Lorem pariatur mollit ex esse exercitation amet. Nisi anim cupidatat excepteur officia. Reprehenderit nostrud nostrud ipsum Lorem est aliquip amet voluptate voluptate dolor minim nulla est proident. Nostrud officia pariatur ut officia. Sit irure elit esse ea nulla sunt ex occaecat reprehenderit commodo officia dolor Lorem duis laboris cupidatat officia voluptate. Culpa proident adipisicing id nulla nisi laboris ex in Lorem sunt duis officia eiusmod. Aliqua reprehenderit commodo ex non excepteur duis sunt velit enim. Voluptate laboris sint cupidatat ullamco ut ea consectetur et est culpa et culpa duis."

  buff1 := guffers.Buffer {
    Styles: guffers.Props {
      BgCl: 2,
      Cl: 1,
    },
    Children: []interface {} {
      &lorem2,
    },
  }

  mainbuff1 := guffers.Buffer {
    Styles: guffers.Props {
      W: 60,
      H: 6,
      X: 4,
      Y: 2,
      BgCl: 7,
      Cl: 0,
      Padding: "1 2",
    },
    Children: []interface {} {
      &lorem1,
      &buff1,
    },
  }

  scr.Add(&mainbuff1)
  scr.Refresh()
  scr.End()
}
