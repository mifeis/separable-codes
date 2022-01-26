# Separable-Codes
Repository for the final grade thesis in partial fulfilment of the requirements for the degree in Telecomunnications Services and Technology Engineering.
Here you can find the coded algorithms to obtain such data as the number of combinations, dependence and unfavourable cases to apply in the Lovász Local Lemma, which proves the existence in a space of probability where the dependent events can be avoided.

Constants:
WORDS Defines the words of the code
GROUP References the subset size of rows 
REPS Maximum num of repetitions in a pair for a fixed GROUP elements

The different results are exported in some files on the /out folder.

**Install**

go get github.com/mifeis/separable-codes

or, to download:

go clone https://github.com/mifeis/separable-codes

**Use**

go run main.go

or

go build

./separable-codes.exe

**Other libraries needed**

- github.com/gonum/gonum/blob/master/stat/combin

- github.com/qax-os/excelizegithub.com/360EntSecGroup-Skylar/excelizev1.4
