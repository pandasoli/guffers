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
  buff.Pos = "out"

  self.buffs = append(self.buffs, buff)
}

func (self *Screen) refresh(buff *Buffer) {
  w := buff.final_props.W
  h := buff.final_props.H

  // Call parent
  if buff.W != w || buff.H != h {
    buff.final_props.W = buff.W
    buff.final_props.H = buff.H
    buff.final_props.Y = buff.Y
    buff.final_props.X = buff.X

    if buff.parent != nil {
      buff.parent.Refresh()
      return
    }
  }

  // Set new style props
  buff.final_props.BgCl = buff.BgCl
  buff.final_props.Cl = buff.Cl
  buff.final_props.Margin = buff.Margin
  buff.final_props.Padding = buff.Padding

  // Create or clear buffer
  if len(buff.buff) < h {
    for i := 0; i < h - len(buff.buff); i++ {
      buff.buff = append(buff.buff, "")
    }
  } else if len(buff.buff) > h {
    buff.buff = buff.buff[0:h]
  } else {
    for i := 0; i < len(buff.buff); i++ {
      buff.buff[i] = strings.Repeat("", w)
    }
  }

  buff_line := 0
  for childi := 0; childi < len(buff.children); childi++ {
    child_buff, isBuff := buff.children[childi].(*Buffer)
    str, isStr := buff.children[childi].(*string)

    if isBuff {
      child_buff.Refresh()

      for buffi := 0; buffi < len(child_buff.buff); buffi++ {
        buff.buff[buff_line] = child_buff.buff[buffi]
        buff_line++
      }
    } else if str := *str; isStr {
      for chi := 0; chi < len(str); chi += w {
        strlen := len(str[chi:])

        if strlen < w {
          buff.buff[buff_line] = str[chi:chi + strlen] + strings.Repeat("", w - strlen)
        } else {
          buff.buff[buff_line] = str[chi:chi + w]
        }

        buff_line++
      }
    }
  }

  // Show in screen
  for y := 0; y < len(buff.buff); y++ {
    fmt.Printf(
      "\033[4%d;3%dm\033[%d;%dH%s\033[0m",
      buff.final_props.BgCl,
      buff.final_props.Cl,
      buff.final_props.Y + y,
      buff.final_props.X,
      buff.buff[y],
    )
  }
}

func (self *Screen) Refresh() {
  for _, buff := range self.buffs {
    buff.Refresh()
  }
}
