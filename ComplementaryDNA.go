package kata

import (
      "strings"
)

func DNAStrand(dna string) string {
  // your code here 
  var text = ""
  x := strings.Split(dna, "")
  for i:= 0; i < len(dna); i++ {
    if x[i] == "A" {  
      text = text + "T"
    }
    if x[i] == "T" {  
      text = text + "A"
    }
    if x[i] == "G" {  
      text = text + "C"
    }
    if x[i] == "C" {  
      text = text + "G"
    }
  }
  return text
  }