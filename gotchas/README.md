based on "50 Shades of Go: ..." 

 - lambda fn. in go ? `func(unused string) {}` 
 - `unused_vars.go` ciekwe wywołanie lambda? fn, syntax
 - *unused inmports* nie znajduje w tej chwili zastosowania // A jednak bardzo pomocne dla golings do template rozwiązań dla algoexpert
 - `redeclaring_vars.go` unikaj takiej redeklaracji w multi-line
 - aha czyli moge zrobić redeklaracje w multi-line ale nie moge jeśli
 chce przypisać do Field Value // działa short decl. dla nie field struct
 - use `go tool vet -shadow ./` to shadow catching; NO LONGER included as of Go .12+ USE staticcheck.dev
 - You can't use the `cap()` on maps. Implementacja hash mapy powoduje że nie ma fn. `cap()` nie ma sensu tak jak dla arrays i slices.
 - `011_arr_fn_args` przykład kóry ma w sobie pointery, pass by reference, pass by copy oraz lambde.
 - korzystając z `range` nie lusisz odbierać dwóch zmiennych, dla manie to taki enumerate