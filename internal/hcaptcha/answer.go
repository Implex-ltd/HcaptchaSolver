package hcaptcha

import (
	"fmt"
	"github.com/Implex-ltd/hcsolver/internal/solver"
)

var (
	RETRY   = 0
)

func (c *Hcap) SolveImages(captcha *Captcha) (map[string]string, error) {
	retry := 0

	for retry <= RETRY {
		if len(captcha.Tasklist) <= 0 {
			return nil, fmt.Errorf("no images found")
		}

		response := solver.Task(&solver.BodyNewSolveTask{
			Question: captcha.RequesterQuestion.En,
			TaskType: captcha.RequestType,
			TaskList: captcha.Tasklist,
		}, c.Logger)

		if !response.Success {
			retry++
			continue
		}

		return response.Data, nil
	}

	return nil, fmt.Errorf("max retry exceded")
}
