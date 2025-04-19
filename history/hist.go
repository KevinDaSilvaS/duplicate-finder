package history

type Record struct {
	Occurencies map[string][]string
	FoldersCount *uint32
	FilesCount *uint32
}

func (r *Record) IncFile() {
	*r.FilesCount = *r.FilesCount+1
}

func (r *Record) IncFolder() {
	*r.FoldersCount = *r.FoldersCount+1
}

func (r *Record) HandleHash(path string, hash string) {
	key := hash
    if val, ok := r.Occurencies[key]; ok {
		r.Occurencies[key] = append(val, path)
		return
	}

	r.Occurencies[key] = []string{path}
}

func New() Record {
	s := uint32(0)
	v := uint32(0)
	return Record{ Occurencies: make(map[string][]string), FilesCount: &s, FoldersCount: &v }
}