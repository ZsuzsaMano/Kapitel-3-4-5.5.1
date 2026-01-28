package main

// potenziere gibt den Potenzwert basis hoch exponent zurück.
// exponent muss >= 0 sein.

// welche die 𝑛‑te Potenz einer Zahl berechnet.
// Beachten Sie dabei:
// • Der Exponent 𝑛 ist eine nicht‑negative ganze Zahl. Dies muss Ihre Funktion jedoch nicht über‑
// prüfen.
// • Wir definieren 00 als 1.
// • Implementieren Sie Ihre Funktion mit einer Schleife, d.h. benutzen Sie nicht die Bibliotheks‑
// funktion math.pow().
// • Sie dürfen davon ausgehen, dass Ihrer Funktion nur Argumente übergeben werden, die einen
// Potenzwert im zulässigen Integer‑Wertebereich ergeben.

func potenziere(basis, exponent int) int {
	if(exponent==0){
return 1}
var result int = basis
for i:=2; i<=exponent; i++ {
result= result*basis
}
	return result
}
