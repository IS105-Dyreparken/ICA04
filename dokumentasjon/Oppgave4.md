# Oppgave 4

a)

|Fakultet  |Studenter 2014|Sannsynlighet i %|
|----------------------------|-------|------|
|Økonomi og samfunnsvitenskap|3093   |29,35%|
|Teknologi og realfag        |2166   |20,55%|
|Helse og idrettsvitenskap   |1829   |17,35%|
|Humanioria og pedagogikk    |1525   |14,47%|
|Lærerutdanning              |1506   |14,29%|
|Kunstfag                    |420    |3,999%|
|Sum                         |10539  |100%  |

b)

N = Totale valg?
M = Studenter på en gitt linje?

Helse og idrett
Infomengde = log2(1/(1829/10539)) = log2(10 539/1829) = 2,526611002194811 = 3 bits

Humaniora og pedagogikk
Infomengde = log2(1/(1525/10539)) = log2(10539/1525) = 2,788856834605916 = 3 bits

Kunstfag
Infomengde = log2 (1/(420/10539)) = log2(10539/420) = 4,649204844277404 = 5 bits

Teknologi og realfag
Infomengde = log2 (1/(2166/10539)) = log2(10539/2166) = 2,2826328343352 = 3 bits

Lærerutdanning
Infomengde = log2 (1/(1506/10539)) = log2(10539/1506) = 2,806944307271599 = 3 bits

Økonomi og samfunnsvitenskap = log2(1/(3093/10539)) = log2(10539/3093) = 1,768659243854263 = 2 bits

Det vi lærer er at fakultetet for Økonomi og samfunnsvitenskap returnerer oss minst informasjon. Dette gir mening i forhold til Huffman-kode, da det som har høyest sannsynlighet for å gå igjen blir lagret med minst mulig bits, slik at vi får maks ut av komprimeringen.

c)  
![](ICA04/dokumentasjon/images/Huffman Tree riktig.png)

d)  
Hvis vi tar utgangspunkt i tallene over, så får vi 16 bits totalt. Hvis vi deler disse 16 bit på 6, så får vi en gjennomsnittslengde på 2.67 bit for fakultetskodene

Ved 100 mulige utfall vil vi få et gjennomsnitt på 2.67*100 utfall = 267 bit
Altså vil den gjennomsnittlige størrelsen på en lengde med 100 fakultetskoder være 267 bit.
