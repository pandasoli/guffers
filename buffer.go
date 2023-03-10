package guffers
import "fmt"


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

  panic(fmt.Errorf("'%v' is not accepted as a child", item))
}

func (self *Buffer) Refresh() {
  self.scr.refresh(self)
}
