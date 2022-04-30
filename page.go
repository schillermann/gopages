package gopages

import "fmt"

type PageMetadata struct {
	Key   string
	Value string
}

type Page struct {
	MetadataList []PageMetadata
}

type PageInterface interface {
	Metadata() Page
	Print()
}

func (p Page) Metadata() Page {
	return p
}

func (p Page) Print() {
	fmt.Println(p.MetadataList[0].Key)
}
