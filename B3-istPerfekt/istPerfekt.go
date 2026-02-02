package main

// istPerfekt bestimmt, ob zahl eine perfekte Zahl ist.
// zahl muss eine natürliche Zahl sein.

// Eine natürliche Zahl 𝑛 heißt perfekt, wenn 𝑛 die Summe aller natürlichen Zahlen 𝑖 mit 1 ≤ 𝑖 < 𝑛 ist,
// durch die 𝑛 ohne Rest teilbar ist. Das heißt, eine Zahl ist perfekt, wenn sie gleicher der Summe all ihrer
// Teiler (außer sich selbst) ist
func istPerfekt(zahl int) bool {
	sumOfTeiler:=0
	for i:=1; i<zahl; i++{
		if(zahl%i ==0){
		sumOfTeiler+=i}
	}
	return sumOfTeiler==zahl
}