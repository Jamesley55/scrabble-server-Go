package models

type Dictionary struct {
	Title       string
	Description string
	Words       []string
}

type DictionaryErrors struct {
	InvalidFileType            bool
	InvalidJson                bool
	InvalidTitle               bool
	InvalidDictionaryStructure bool
}

type DictHeaders struct {
	Title       string
	Description string
}
