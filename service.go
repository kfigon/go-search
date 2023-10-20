package main

type Service struct {
	t *Trie
	index *TfIdf
}

func NewService(docs map[FileName]string) *Service {
	index := Parse(docs)
	t := NewTrie()
	for word := range index {
		t.Add(string(word))
	}

	return &Service{
		t: t,
		index: &index,
	}
}

func (s *Service) Suggest(prefix string) []string {
	return s.t.FindPrefixed(prefix)
}

func (s *Service) Results(word string) []FileName {
	names := (*s.index)[Word(word)]
	out := make([]FileName, len(names))

	for fileName,v := range names {
		if v > 0 {
			out = append(out, fileName)
		}
	}
	return out
}