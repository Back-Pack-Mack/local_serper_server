package main

type SearchOption func(map[string]interface{})

func WithNum(num int) SearchOption {
	return func(params map[string]interface{}) {
		params["num"] = num
	}
}

func WithPage(page int) SearchOption {
	return func(params map[string]interface{}) {
		params["page"] = page
	}
}

func WithCountry(gl string) SearchOption {
	return func(params map[string]interface{}) {
		params["gl"] = gl
	}
}

func WithLanguage(hl string) SearchOption {
	return func(params map[string]interface{}) {
		params["hl"] = hl
	}
}

func WithLocation(location string) SearchOption {
	return func(params map[string]interface{}) {
		params["location"] = location
	}
}

func WithType(searchType string) SearchOption {
	return func(params map[string]interface{}) {
		params["type"] = searchType
	}
}

func WithAutocorrect(autocorrect bool) SearchOption {
	return func(params map[string]interface{}) {
		params["autocorrect"] = autocorrect
	}
}
