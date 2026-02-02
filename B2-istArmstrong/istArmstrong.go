package main

import (
	"strconv"
)

// istArmstrong bestimmt, ob zahl eine Armstrong-Zahl ist.
// zahl muss eine natürliche Zahl sein.

// Eine Armstrong‑Zahl ist eine Zahl, bei der die Summe ihrer einzelnen Ziffern, jeweils potenziert mit
// der Ziffernanzahl der Zahl, gleich der Zahl selbst ist.
func istArmstrong(zahl int) bool {
	var zahlLength int= len(strconv.Itoa(zahl) ) 
  var zahlToSlice int = zahl
   sumOfPotenz:=0

  for i:=0; i<=zahlLength; i++{
     potenzOfDigit:=potenziere(zahlToSlice%10, zahlLength)
   sumOfPotenz+=potenzOfDigit
    zahlToSlice= zahlToSlice/10
}
    

	return sumOfPotenz==zahl
}
