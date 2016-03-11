package main

import (
	"encoding/json"

	"github.com/ericdaugherty/gotsport-scraper"
	"github.com/jasonmoo/lambda_proc"
)

func main() {

	lambda_proc.Run(func(context *lambda_proc.Context, eventJSON json.RawMessage) (interface{}, error) {

		var params map[string]string
		if err := json.Unmarshal(eventJSON, &params); err != nil {
			return nil, err
		}

		schedule, err := scraper.GetSchedule(params)

		if err != nil {
			return nil, err
		}

		return schedule, nil
	})

}
