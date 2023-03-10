package guffers

import (
	"fmt"
	"strings"
)


type Screen struct {
  buffs []*Buffer
}


func (self *Screen) Add(buff *Buffer) {
  buff.scr = self
  buff.Styles.Pos = "out"

  self.buffs = append(self.buffs, buff)
}

func (self *Screen) refresh(buff *Buffer) {
  // Set old/final styles equals current styles
  buff.styles = buff.Styles

  // Clear Buffer's buffer
  buff.buff = []string {}

  // Load properties
  buff.styles.process_compounds()

  // Show on screen
  defer func() {
    if buff.parent != nil {
      return
    }

    for i := 0; i < buff.styles.CompH; i++ {
      line := ""

      if len(buff.buff) > i {
        line = buff.buff[i]
      }

      // Fill line with whitespaces if its width is less than the buffer's
      if len(line) < buff.styles.CompW {
        line += strings.Repeat(" ", buff.styles.CompW - len(line))
      }

      line =
        fmt.Sprintf("\033[%d;%dH", buff.styles.CompY + i, buff.styles.CompX) + // Set position
        fmt.Sprintf("\033[%d;4%d;3%dm", buff.styles.FontStyle, buff.styles.BgCl, buff.styles.Cl) + // Set colors
        line + // Print buffer's line
        "\033[0m" // Reset color

      fmt.Print(line)
    }
  }()

  // Loading children
  usable_w := buff.styles.CompW - (buff.styles.LeftPadding + buff.styles.RightPadding)

  for _, child := range buff.Children {
    switch child := child.(type) {
    case *string:
      for i := 0; i < len(*child); i += usable_w {
        line := (*child)[i:]

        if len(line) > usable_w {
          line = line[:usable_w]
        }

        buff.buff = append(buff.buff, line)
      }

    case *Buffer:
      child.Styles.process_compounds()
      child.Styles.CompW = usable_w - (child.Styles.LeftMargin + child.Styles.RightMargin)
      child.Refresh()

      for range make([]int, child.Styles.TopMargin) {
        buff.buff = append(buff.buff, "")
      }

      for _, line := range child.buff {
        line =
          strings.Repeat(" ", child.Styles.LeftMargin) + // Create child's left margin
          fmt.Sprintf("\033[%d;4%d;3%dm", child.styles.FontStyle, child.styles.BgCl, child.styles.Cl) + // Set child's color
          line +
          fmt.Sprintf("\033[%d;4%d;3%dm", buff.styles.FontStyle, buff.styles.BgCl, buff.styles.Cl) + // Set parent's color
          strings.Repeat(" ", child.Styles.RightMargin) // Create child's right margin

        buff.buff = append(buff.buff, line)
      }

      for range make([]int, child.Styles.BottomMargin) {
        buff.buff = append(buff.buff, "")
      }

    default:
      panic(fmt.Errorf("Type '%v' is not accepted as a child\n", child))
    }
  }

  // Making paddings
  for linei := range buff.buff {
    buff.buff[linei] =
      strings.Repeat(" ", buff.styles.LeftPadding) +
      buff.buff[linei] +
      strings.Repeat(" ", buff.styles.RightPadding)
  }

  for range make([]int, buff.styles.TopPadding) {
    buff.buff = append([]string {""}, buff.buff...)
  }

  for range make([]int, buff.styles.BottomPadding) {
    buff.buff = append(buff.buff, "")
  }
}

func (self *Screen) Refresh() {
  for _, buff := range self.buffs {
    buff.Refresh()
  }
}

func (self *Screen) End() {
  yest := 0

  for _, buff := range self.buffs {
    if buff.styles.CompY + buff.styles.CompH > yest {
      yest = buff.styles.CompY + buff.styles.CompH
    }
  }

  fmt.Printf("\033[%d;0H", yest + 2)
}
