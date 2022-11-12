package dtos

type ResponseDto struct {
	NumFound int      `json:"numFound"`
	Docs     ItemsDTO `json:"docs"`
}

type SolrResponseDto struct {
	Response ResponseDto `json:"response"`
}

type DocDto struct {
	Doc ItemDTO `json:"doc"`
}

type AddDto struct {
	Add DocDto `json:"add"`
}

type SolrResponsesDto []SolrResponseDto
