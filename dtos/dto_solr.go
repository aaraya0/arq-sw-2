package dtos

type ResponseDto struct {
	NumFound int      `json:"numFound"`
	Docs     ItemsDTO `json:"docs"`
}

type SolrResponseDto struct {
	Response ResponseDto `json:"response"`
}

type SolrResponsesDto []SolrResponseDto
