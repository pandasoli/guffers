package guffers
import (
  "fmt"
  "strings"
)


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

func (self *Buffer) Show() {
  for i := 0; i < self.styles.CompH; i++ {
    line := ""

    if len(self.buff) > i {
      line = self.buff[i]
    }

    // Fill line with whitespaces if its width is less than the buffer's
    if len(line) < self.styles.CompW {
      line += strings.Repeat(" ", self.styles.CompW - len(line))
    }

    line =
      fmt.Sprintf("\033[%d;%dH", self.styles.CompY + i, self.styles.CompX) + // Set position
      fmt.Sprintf("\033[%d;4%d;3%dm", self.styles.FontStyle, self.styles.BgCl, self.styles.Cl) + // Set colors
      line + // Print buffer's line
      "\033[0m" // Reset color

    fmt.Print(line)
  }
}

func (self *Buffer) Refresh() {
  self.scr.refresh(self)
}
