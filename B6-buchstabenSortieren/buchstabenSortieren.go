package main

import (
	"slices"
	"strings"
)

// buchstabenSortieren gibt einen String zurück, in dem alle Buchstaben
// von wort alphabetisch sortiert sind. In wort mehrfach vorkommende
// Buchstaben kommen im resultierendem String genauso oft vor. Großbuchstaben
// in wort kommen im resultierendem String als Kleinbuchstaben vor.
func buchstabenSortieren(wort string) string {
lowerCaseWort:=strings.ToLower(wort)
arrayOfLetters:=strings.Split(lowerCaseWort, "")
slices.Sort(arrayOfLetters)
return strings.Join(arrayOfLetters, "")
}
