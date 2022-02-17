package httpx

type BasePageModel struct {
	Data 			interface{}		`json:"data"`
	Count 			int64			`json:"count"`
}

