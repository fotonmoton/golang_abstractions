package search

type SearchValidator interface {
	Validate(string) error
}

type Searcher interface {
	Find(string) []string
}

type SearchSerializer interface {
	Serialize([]string) string
}

type Search struct {
	// validator  SearchValidator
	searcher   Searcher
	serializer SearchSerializer
}

// func NewSearch(validator SearchValidator, searcher Searcher, serializer SearchSerializer) *search {
// 	return &search{validator, searcher, serializer}
// }

func NewSearch(searcher Searcher, serializer SearchSerializer) *Search {
	return &Search{searcher, serializer}
}

// func (s *search) GetProducts(searchTerm string) string {
// 	if err := s.validator.Validate(searchTerm); err != nil {
// 		return ""
// 	}

// 	products := s.searcher.Find(searchTerm)

// 	return s.serializer.Serialize(products)
// }

func (s *Search) GetProducts(searchTerm string) string {

	products := s.searcher.Find(searchTerm)

	return s.serializer.Serialize(products)
}
