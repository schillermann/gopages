package gopages

import "fmt"

type Metadata struct {
	Key   string
	Value string
}

type Page struct {
	MetadataList []Metadata
}

func (p Page) Metadata() Page {
	return p
}

func (p Page) Print() {
	fmt.Println(p.MetadataList[0].Key)
}
