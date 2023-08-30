package hcaptcha

import (
	"fmt"
	"github.com/Implex-ltd/hcsolver/internal/recognizer"
)

var (
	RETRY = 0
)

func (c *Hcap) SolveImages(captcha *Captcha) (any, error) {
	retry := 0

	for retry <= RETRY {
		if len(captcha.Tasklist) <= 0 {
			return nil, fmt.Errorf("no images found")
		}

		r := recognizer.NewRecognizer(captcha.RequestType, captcha.RequesterQuestion.En, captcha.RequesterRestrictedAnswerSet, captcha.Tasklist)
		response, err := r.Recognize()
		if err != nil {
			return map[string]any{}, err
		}

		if !response.Success {
			retry++
			continue
		}

		return response.Data, nil
	}

	return nil, fmt.Errorf("max retry exceded")
}
