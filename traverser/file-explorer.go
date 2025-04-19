package traverser

import (
	"os"
	"start/hasher"
	"start/history"
)

func MapFolder(path string, record history.Record) {
	files, err := os.ReadDir(path)
	if err != nil {
		return
	}

	record.IncFolder()
	for _, file := range files {
		if file.IsDir() {
			MapFolder(path+"/"+file.Name(), record)
			continue
		}

		MapFileContents(path+"/"+file.Name(), record)
	}
}

func MapFileContents(filePath string, record history.Record)  {
	file, err := os.Open(filePath) // For read access.
	if err != nil {
		return
	}

	data := make([]byte, 300)
	_, errRead := file.Read(data)
	if errRead != nil {
		return
	}

	hash := string(hasher.HashIt(string(data)))
	record.HandleHash(filePath, hash)
	record.IncFile()
}