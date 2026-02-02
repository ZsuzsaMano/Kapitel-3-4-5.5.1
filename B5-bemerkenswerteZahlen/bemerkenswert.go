package main

// bemerkenswert gibt die nächste natürliche Zahl größer start zurück, deren
// Quadrat die Summe der natürlichen Zahlen von 1 bis n ist (mit beliebigem n).
func bemerkenswert(start int) int {
   zahl:=start+1

   zahlQuadrat:= zahl*zahl
   var dreieckszahl int
for i:=0; i<=zahlQuadrat; i++{
  dreieckszahl=i*(i+1)/2
 if (dreieckszahl==zahlQuadrat){
	return zahl
}} 
 return bemerkenswert(zahl+1)

}
