package guffers
import "fmt"


type Props struct {
  W, H,
  X, Y,
  CompX, CompY,
  Scroll,

  BgCl, FontStyle, Cl,

  TopPadding, RightPadding,
  BottomPadding, LeftPadding int

  // abs - absolute (can move free inside the parent)
  // rel - relative (cannot move)
  // out - outside (can move free inside the screen)
  Pos,

  Padding string
}

type Buffer struct {
  // This only accepts *Buffer and *string
  Children []interface {}
  buff []string

  Styles Props
  styles Props

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

    self.Children = append(self.Children, item)
    return
  }

  _ = fmt.Errorf("'%v' is not accepted", item)
}

func (self *Buffer) Refresh() {
  self.scr.refresh(self)
}
