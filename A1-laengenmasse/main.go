// Aufgabe A1: Längenmaße
// Schreiben Sie ein Programm, das eine Länge in Metern einliest und diese umrechnet in:
// • mm
// • km
// • Zoll (1 Zoll = 2.54 cm)
// • Seemeilen (1 sm = 1852 m)
// • Lichtjahre (1 Lj = 9 460 730 472 580 800 m)
// Die Ausgabe Ihres Programmes könnte so aussehen:
// Geben Sie eine Länge (in Metern) ein:
// 3755
// 3755 Meter entsprechen:
// 3.755e+06 mm
// 3.755 km
// 147834.646 Zoll
// 2.03 sm
// 3.97e-13 Lj

package main

import "fmt"


func umrechnenVonMeter() {
var eingabe float64
fmt.Printf("Geben Sie eine Länge (in Metern) ein: ")
fmt.Scan(&eingabe)
fmt.Printf(`%v Meter entsprechen:
%g mm,
%g km,
%3f Zoll,
%3g sm,
%3g Lj `, 
eingabe, eingabe*1000, eingabe/1000 , eingabe*39.37008, eingabe*0.00053995680345572 , eingabe*1.0570008340247E-16 )
}

func main(){
     umrechnenVonMeter()}