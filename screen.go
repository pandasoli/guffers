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
  buff.styles.CompX = buff.styles.X
  buff.styles.CompY = buff.styles.Y

  // Clear Buffer's buffer
  buff.buff = []string {}

  // Load properties
  paddings := strings.Split(buff.styles.Padding, " ")
  if paddings, ok := strArr2IntArr(paddings); ok {
    switch len(paddings) {
    case 1:
      buff.styles.TopPadding = paddings[0]
      buff.styles.RightPadding = paddings[0]
      buff.styles.BottomPadding = paddings[0]
      buff.styles.LeftPadding = paddings[0]
    case 2:
      buff.styles.TopPadding = paddings[0]
      buff.styles.RightPadding = paddings[1]
      buff.styles.BottomPadding = paddings[0]
      buff.styles.LeftPadding = paddings[1]
    case 3:
      buff.styles.TopPadding = paddings[0]
      buff.styles.RightPadding = paddings[1]
      buff.styles.BottomPadding = paddings[2]
      buff.styles.LeftPadding = paddings[1]
    case 4:
      buff.styles.TopPadding = paddings[0]
      buff.styles.RightPadding = paddings[1]
      buff.styles.BottomPadding = paddings[2]
      buff.styles.LeftPadding = paddings[3]
    default:
      fmt.Errorf("The padding property cannot accept more than 4 items\n")
      return
    }
  } else {
    fmt.Errorf("Could not convert paddings to int\n")
    return
  }

  // Show on screen
  defer func() {
    for i := 0; i < buff.styles.H; i++ {
      line := buff.buff[i]

      // Fill line with whitespaces if its width is less than the buffer's
      if len(line) < buff.styles.W {
        line += strings.Repeat(" ", buff.styles.W - len(line))
      }

      fmt.Printf("\033[%d;%dH", buff.styles.CompY + i, buff.styles.CompX) // Go to position
      fmt.Printf("\033[%d;4%d;3%dm", buff.styles.FontStyle, buff.styles.BgCl, buff.styles.Cl) // Set colors
      fmt.Printf("%s", line) // Print buffer's line
      fmt.Printf("\033[0m") // Reset color
    }
  }()

  // Loading children
  usable_w := buff.styles.W - (buff.styles.LeftPadding + buff.styles.RightPadding)

  for _, child := range buff.Children {
    switch child.(type) {
      case *string:
        if child, ok := child.(*string); ok {
          for i := 0; i < len(*child); i += usable_w {
            line := (*child)[i:]

            if len(line) > usable_w {
              line = line[:usable_w]
            }

            buff.buff = append(buff.buff, line)
          }
        }

      default:
        fmt.Errorf("Type '%v' is not acceptable\n", child)
        return
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
    if buff.styles.CompY + buff.styles.H > yest {
      yest = buff.styles.CompY + buff.styles.H
    }
  }

  fmt.Printf("\033[%d;0H", yest + 2)
}
