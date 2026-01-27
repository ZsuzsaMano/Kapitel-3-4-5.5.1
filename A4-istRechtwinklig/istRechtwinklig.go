package main

// istRechtwinklig prüft, ob das durch die Seitenlängen x, y, z gebildete
// Dreieck rechtwinklig ist.
func istRechtwinklig(x, y, z int) bool {
	// bitte korrigieren Sie den Rumpf dieser Funktion.

if(y>z && y>x){
return y*y==x*x+z*z}
if(x>y && x>z){
return x*x==y*y+z*z
}else
{return z*z==x*x+y*y}
}
