package cli

import ( 
	"os" 
	"fmt"
	"start/traverser"
	"start/history" 
)

func Entrypoint()  {
	Reset := "\033[0m"
	Green := "\033[32m"
	BrightMagenta := "\033[35;1m"
	Pink := "\033[38;5;199m"
	Lime := "\033[38;5;76m"
	LightPink := "\033[38;5;225m"

	path := os.Args[1]
	r := history.New()
	traverser.MapFolder(path, r)

	fmt.Println(LightPink, "*********************************")
	fmt.Println(Green, "* Files checked:   ", Lime, *r.FilesCount)
	fmt.Println(Green, "* Folders checked: ", Lime, *r.FoldersCount)

	for _, v := range r.Occurencies {
		if len(v) > 1 {
			fmt.Println("")
			fmt.Println(BrightMagenta, "*** Duplication found: ")

			for _, fileDupPath := range v {
				fmt.Println(BrightMagenta, "  >>> ", Pink, fileDupPath)
			}

		}
	}
	fmt.Println("")
	fmt.Println(LightPink, "*********************************", Reset)
}