package elastic

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Request struct {
	UserName string `json:"userName"`
	Message  string `json:"message"`
	Hour     string `json:"hour"`
	DateFrom string `json:"dateFrom"`
	DateEnd  string `json:"dateEnd"`
}

func Query(ctx *fiber.Ctx, subject string) map[string]interface{} {
	var req Request
	if err := ctx.BodyParser(&req); err != nil {
		println(err)
	}

	countWords := func(str string) bool {
		parsedString := strings.Split(str, " ")
		if len(parsedString) > 1 {
			return true
		} else {
			return false
		}
	}

body := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"match": map[string]interface{}{
							"subject": subject,
						},
					},
				},
			},
		},
		"sort": []interface{}{
			map[string]interface{}{
				"date": "desc",
			},
		},
	}
	if countWords(req.UserName) {
		query := map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"userName": map[string]interface{}{
					"query": req.UserName,
					"boost": 4.0,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	if userName := req.UserName; userName != "" && ctx.FormValue("userName") == "match" {
		query := map[string]interface{}{
			"match": map[string]interface{}{
				"userName": map[string]interface{}{
					"query": userName,
					"boost": 4.0,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	if userName := req.UserName; userName != "" && ctx.FormValue("userName") == "fuzz" {
		query := map[string]interface{}{
			"fuzzy": map[string]interface{}{
				"userName": map[string]interface{}{
					"value":     userName,
					"fuzziness": 1,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	if userName := req.UserName; userName != "" && ctx.FormValue("userName") == "prefix" {
		query := map[string]interface{}{
			"prefix": map[string]interface{}{
				"userName": map[string]interface{}{
					"value": userName,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	if message := req.Message; message != "" && ctx.FormValue("message") == "matchPhrase" {
		query := map[string]interface{}{
			"match_phrase": map[string]interface{}{
				"text": map[string]interface{}{
					"query": message,
					"boost": 4.0,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	if message := req.Message; message != "" && ctx.FormValue("message") == "matchPhrasePrefix" {
		query := map[string]interface{}{
			"match_phrase_prefix": map[string]interface{}{
				"text": map[string]interface{}{
					"query": message,
					"boost": 4.0,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	dateFrom := req.DateFrom
	dateEnd := req.DateEnd

	if dateFrom != "" && dateEnd != "" {
		query := map[string]interface{}{
			"range": map[string]interface{}{
				"date": map[string]interface{}{
					"gte": dateFrom,
					"lte": dateEnd,
				},
			},
		}
		body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(body["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}), query)
	}

	return body
}
