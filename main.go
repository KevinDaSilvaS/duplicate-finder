package main

import (
	"start/cli" 
)

func main() {
	/* r := history.New()
	traverser.MapFolder(".", r)
	fmt.Println(*r.FilesCount, *r.FoldersCount)
	for _, v := range r.Occurencies {
		if len(v) > 1 {
			fmt.Println(v)
		}
	} */
	
	cli.Entrypoint()
}