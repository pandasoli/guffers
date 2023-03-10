package guffers

import (
  "fmt"
  "strings"
)


type Props struct {
  W, H,
  CompW, CompH,
  X, Y,
  CompX, CompY,
  Scroll,

  BgCl, FontStyle, Cl,

  TopPadding, RightPadding,
  BottomPadding, LeftPadding,

  TopMargin, RightMargin,
  BottomMargin, LeftMargin int

  // abs - absolute (can move free inside the parent)
  // rel - relative (cannot move)
  // out - outside (can move free inside the screen)
  Pos,

  Padding, Margin string
}

func (self *Props) process_compounds() {
  paddings := strings.Split(self.Padding, " ")
  if paddings, ok := strArr2IntArr(paddings); ok {
    switch len(paddings) {
    case 1:
      self.TopPadding = paddings[0]
      self.RightPadding = paddings[0]
      self.BottomPadding = paddings[0]
      self.LeftPadding = paddings[0]
    case 2:
      self.TopPadding = paddings[0]
      self.RightPadding = paddings[1]
      self.BottomPadding = paddings[0]
      self.LeftPadding = paddings[1]
    case 3:  
      self.TopPadding = paddings[0]
      self.RightPadding = paddings[1]
      self.BottomPadding = paddings[2]
      self.LeftPadding = paddings[1]
    case 4:  
      self.TopPadding = paddings[0]
      self.RightPadding = paddings[1]
      self.BottomPadding = paddings[2]
      self.LeftPadding = paddings[3]
    default:
      panic(fmt.Errorf("The padding property cannot accept more than 4 items\n"))
    }    
  } else { 
    // panic(fmt.Errorf("Could not convert paddings to int\n"))
    self.TopPadding = 0
    self.RightPadding = 0
    self.BottomPadding = 0
    self.LeftPadding = 0
  }

  margins := strings.Split(self.Margin, " ")
  if margins, ok := strArr2IntArr(margins); ok {
    switch len(margins) {
    case 1:
      self.TopMargin = margins[0]
      self.RightMargin = margins[0]
      self.BottomMargin = margins[0]
      self.LeftMargin = margins[0]
    case 2:       
      self.TopMargin = margins[0]
      self.RightMargin = margins[1]
      self.BottomMargin = margins[0]
      self.LeftMargin = margins[1]
    case 3:            
      self.TopMargin = margins[0] 
      self.RightMargin = margins[1]
      self.BottomMargin = margins[2]
      self.LeftMargin = margins[1] 
    case 4:            
      self.TopMargin = margins[0]
      self.RightMargin = margins[1]
      self.BottomMargin = margins[2]
      self.LeftMargin = margins[3]
    default:    
      panic(fmt.Errorf("The padding property cannot accept more than 4 items\n"))
    }        
  } else {     
    // panic(fmt.Errorf("Could not convert paddings to int\n"))
    self.TopMargin = 0
    self.RightMargin = 0
    self.BottomMargin = 0
    self.LeftMargin = 0
  } 
}
