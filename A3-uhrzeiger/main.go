package main

// Aufgabe A3: Uhrzeiger
// Schreiben Sie ein Programm, das eine Uhrzeit (Stunde 0 bis 23, Minute 0 bis 59) einliest und den Winkel
// des Stunden- und Minutenzeigers zu dieser Uhrzeit ausgibt. Die Winkel sind 0°, wenn die Zeiger oben
// stehen, also um 0 und 12 Uhr.
// Die Ausgabe Ihres Programmes könnte so aussehen:
// Geben Sie bitte eine Uhrzeit an, indem Sie zunächst
// die Stunde (von 0 bis 23) und dann die Minute (von 0 bis 59) eingeben:
// 8
// 5
// Zeigerstellung um 08:05 Uhr:
// Winkel des Stundenzeigers: 242.5°
// Winkel des Minutenzeigers: 30°

import "fmt"

const maxWinkel float32 = 360

func uhrZeiger (uhr int) float32{
return maxWinkel/12*float32(uhr%12)
}

func minutenZeiger ( minute int) float32{
return maxWinkel/60*float32(minute)
}
func main(){
var uhr int
var minute int

fmt.Println("Geben Sie bitte eine Uhrzeit an, indem Sie erst die Stunde (von 0 bis 23) eingeben:")
fmt.Scan(&uhr)
var uhrWinkel float32 = uhrZeiger(uhr)
if(uhr< 0 || uhr>23){
fmt.Println("Eingegebene Nummer ist nicht gültig, Geben Sie bitte eine nummer zwischen 0 bis 23 ein:")
fmt.Scan(&uhr)
}

fmt.Println("Geben Sie bitte zunächst die die Minute (von 0 bis 59) eingeben:")
fmt.Scan(&minute)
if(minute< 0 || minute>59){
fmt.Println("Eingegebene Nummer ist nicht gültig, Geben Sie bitte eine nummer zwischen 0 bis 59 ein:")
fmt.Scan(&minute)
}
var minutenWinkel float32 = minutenZeiger(minute)

fmt.Printf("Zeigerstellung um %v:%v Uhr \n", uhr, minute)
fmt.Printf("Winkel des Stundenzeigers: %g Grad \n", uhrWinkel)
fmt.Printf("Winkel des Minutenzeigers: %g Grad \n" , minutenWinkel)

}




