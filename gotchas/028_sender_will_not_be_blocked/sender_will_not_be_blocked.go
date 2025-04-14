package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
			time.Sleep(time.Second * 1)
		}
	}()

	ch <- "one"
	ch <- "two" // Sender will not be block !!!
	// jeśli receiver nie będzie gotowy do odbioru to payload idze w kosmos
	// rozwiązaniem jest buffer channel
	// ^^^^^KOMPLETNIE ZŁE ROZUMOWANIE `ch <- 'two'` jest blocking !!! tylko app zsązy wyjść
	// zobacz z time.Sleep
	// ^^^^ JESZCZE INACZEJ, `ch <- "two"` jest blocking ALE po jeśli tylko receiver jest ready to
	// "pewnie przekazuje" leci od razu dalej

	// musisz odpalić kilka razy żeby zobaczyć że czasem udaje się print "two" przed exitem

	// time.Sleep(time.Second * 1)
}
