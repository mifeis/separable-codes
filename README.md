# Separable-Codes

Repositori pel treball de fi de grau d'Enginyeria de Telecomunicacions.

Es tracta d'uns algoritmes que calculen l'existencia de codis separables a partir del Lovász Local Lemma.

El codi proveeix totes les combinacions possibles d'un codi de WORDS paraules per GROUP elements (WORDS sobre GROUP en combinatoria), retorna els casos favorables i desfavorables i, la dependencia entre events.

Els resultats dels tests s'exporten a diferents fitxers excel dins la carpeta /out.

**Install**

go get github.com/mifeis/separable-codes

o, per descarregar el repositori i les seves carpetes:

go clone https://github.com/mifeis/separable-codes

**Use**

Des de la carpeta arrel del repositori:

go run main.go

o

go build

./separable-codes.exe
