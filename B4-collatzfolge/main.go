package main

import "fmt"

// Eine sog. Collatz‑Folge kann nach einem einfachen Bildungsgesetz konstruiert werden:
// • Beginne mit einer beliebigen natürlichen Zahl 𝑛 > 0.
// • Ist 𝑛 gerade, so nimm als nächstes 𝑛/2
// • Ist 𝑛 ungerade, so nimm als nächstes 3𝑛 + 1.
// • Wiederhole die Vorgehensweise mit der erhaltenen Zahl.
func collatzFolge (zahl int){
fmt.Printf("Die Collatz-Folge mit dem Startwert %v lautet:\n", zahl)
result:=zahl
fmt.Print(zahl, ", ")
for result!=1 {
if(result%2==0){
result/=2
}else{
result= result*3+1}
fmt.Printf("%v, ", result)
}
fmt.Printf("...\n\n")
}



func main(){
collatzFolge(19)
collatzFolge(23)
collatzFolge(42)
collatzFolge(122)
}