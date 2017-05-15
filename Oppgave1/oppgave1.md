Vi har lest begge filene vha. fileutils og fant ut at filene er ganske like med noen forskjell i koden.
I koden har vi lagt til en printf metode som formaterer teksten til heksadesimale tegn og string.
Slik kan vi finne ut at text1.txt filen inneholder flere heksadesimal tegn enn text2.txt filen. Ved hjelp av fmt.Prtintf(%x) funksjonen, ser vi at det ligger fler heksadesimale tegn i  text1.txt enn text2.txt filen.

Formateringen ser slik ut:
Tekst 1 er øverst.
