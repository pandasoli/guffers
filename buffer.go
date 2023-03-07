package guffers
import "fmt"


type Props struct {
  W, H,
  X, Y,
  Margin, Padding,
  Scroll int

  BgCl, Cl int

  // abs - absolute (can move free inside the parent)
  // rel - relative (cannot move)
  // out - outside (can move free inside the screen)
  Pos string
}

type Buffer struct {
  // This only accepts *Buffer and *string
  children []interface {}
  buff []string

  Props
  final_props Props

  parent *Buffer
  scr *Screen
}


func (self *Buffer) Add(item interface {}) {
  _, isBuff := item.(*Buffer)
  _, isStr := item.(*string)

  if isBuff || isStr {
    defer self.scr.refresh(self)

    if item, ok := item.(*Buffer); ok {
      item.parent = self
    }

    self.children = append(self.children, item)
    return
  }

  _ = fmt.Errorf("'%v' is not accepted", item)
}

func (self *Buffer) Refresh() {
  self.scr.refresh(self)
}
