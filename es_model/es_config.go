package es_model

type esConfig struct {
	Properties esConfigProperties `json:"properties"`
}

type esConfigProperties struct {
	Content  esConfigContent  `json:"content"`
	FileName esConfigFileName `json:"file_name"`
}

type esConfigContent struct {
	Type           string `json:"type"`
	Analyzer       string `json:"analyzer"`
	SearchAnalyzer string `json:"search_analyzer"`
}

type esConfigFileName struct {
	Type           string `json:"type"`
	Analyzer       string `json:"analyzer"`
	SearchAnalyzer string `json:"search_analyzer"`
}
