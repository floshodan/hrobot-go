package schema

type Error struct {
	Error struct {
		Status  int    `json:"status"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

type InvalidInputError struct {
	Error struct {
		Status  int      `json:"status"`
		Code    string   `json:"code"`
		Message string   `json:"message"`
		Missing []string `json:"missing"`
		Invalid []string `json:"invalid"`
	} `json:"error"`
}

type RequestLimitError struct {
	Error struct {
		Status      int    `json:"status"`
		Code        string `json:"code"`
		Max_Request int    `json:"max_request"`
		Interval    int    `json:"interval"`
		Message     string `json:"message"`
	}
}
