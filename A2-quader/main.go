// Aufgabe A2: Quader
// Schreiben Sie ein Programm, daß die Längen der drei Seiten eines Quaders einliest und folgende
// Größen dieses Quaders ausgibt:
// • Volumen
// • Kantensumme
// • Oberfläche
// • Umkugelradius
// • Raumdiagonale
// Tipps: Wenn Ihnen die Formeln zur Bestimmung dieser Größen nicht bekannt sind, finden Sie sie in einer
// mathematischen Formelsammlung oder auch der Wikipedia. Verwenden Sie geeignete Funktionen
// aus dem Package math.
// Die Ausgabe des Programmes könnte so aussehen:
// Bitte geben Sie die drei Seitenlängen des Quaders ein:
// 3
// 4.5
// 8
// Ein 3 × 4.5 × 8 Quader hat die geometrischen Größen:
// Volumen: 108
// Kantensumme: 62
// Oberfläche: 147
// Umkugelradius: 4.83
// Raumdiagonale: 9.66

package main

import "fmt"

func quader(){
var a float64
var b float64
var c float64
fmt.Printf("Bitte geben Sie die drei Seitenlängen des Quaders ein: \n")
fmt.Scan(&a, &b, &c)
fmt.Printf(`Ein %g × %g × %g Quader hat die geometrischen Größen:`, a,b,c)
}

func main (){
quader()}
