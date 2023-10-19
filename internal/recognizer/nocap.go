package recognizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Implex-ltd/hcsolver/internal/utils"
)

var (
	ImMut, SmMut, AsMut sync.RWMutex
	Client              *http.Client
)

var (
	nocap  = true
	apikey = "rorm-5cd874d2-0d3b-e49c-f760-f1e74b316467"
	proapi = "https://pro.nocaptchaai.com/solve"
)

type Base64JSON struct {
	Images  map[string]string `json:"images"`
	Target  string            `json:"target"`
	Method  string            `json:"method"`
	Sitekey string            `json:"sitekey"`
	Site    string            `json:"site"`
	Ln      string            `json:"ln"`
	Type    string            `json:"type"`
}

type NoCapAnswer struct {
	Answer         []any  `json:"answer"`
	ID             string `json:"id"`
	Message        string `json:"message"`
	ProcessingTime string `json:"processing_time"`
	Solution       []int  `json:"solution"`
	Status         string `json:"status"`
	Target         string `json:"target"`
	URL            string `json:"url"`
}

type NoCapAnswerSelect struct {
	Answer         []int   `json:"answer"`
	Answers        [][]int `json:"answers"`
	ID             string  `json:"id"`
	ProcessingTime string  `json:"processing_time"`
	Solution       []any   `json:"solution"`
	Status         string  `json:"status"`
	Url            string  `json:"url"`
}

func SolvePic(toSolve map[string]map[string]string, prompt, target string) (map[string]string, error) {
	imgs := map[string]string{}
	out := map[string]string{}

	j := 0
	for _, v := range toSolve {
		imgs[fmt.Sprintf("%d", j)] = v["body"]
		j++
	}

	jsonBody, err := json.Marshal(Base64JSON{
		Images:  imgs,
		Target:  prompt,
		Method:  "hcaptcha_base64",
		Sitekey: "4c672d35-0701-42b2-88c3-78380b0db560",
		Site:    "discord.com",
		Ln:      "en",
		Type:    "grid",
	})
	if err != nil {
		return imgs, err
	}

	req, err := http.NewRequest("POST", proapi, bytes.NewBuffer(jsonBody))
	if err != nil {
		return imgs, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("apikey", apikey)

	resp, err := Client.Do(req)
	if err != nil {
		return imgs, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("nocap", err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return imgs, err
	}

	var answer NoCapAnswer
	if err := json.Unmarshal(bodyBytes, &answer); err != nil {
		return imgs, err
	}

	i := 0
	for _, v := range toSolve {
		found := false

		for _, a := range answer.Solution {
			if a == i {
				found = true
				break
			}
		}

		go func(v map[string]string) {
			if found {
				SmMut.Lock()
				defer SmMut.Unlock()

				currentValue, _ := Hashlist.Load(target)

				var updatedValue []string
				if currentValue != nil {
					updatedValue = append(currentValue.([]string), v["hash"])
				} else {
					updatedValue = []string{v["hash"]}
				}

				Hashlist.Store(target, updatedValue)
				utils.AppendLine(fmt.Sprintf("%s,%s", v["hash"], target), "hash.csv")
			} else {
				SmMut.Lock()
				defer SmMut.Unlock()
				currentValue, _ := Hashlist.Load(target)

				var updatedValue []string
				if currentValue != nil {
					updatedValue = append(currentValue.([]string), "noy_"+v["hash"])
				} else {
					updatedValue = []string{"noy_" + v["hash"]}
				}

				Hashlist.Store(target, updatedValue)
				utils.AppendLine(fmt.Sprintf("%s,not_%s", v["hash"], target), "hash.csv")
			}
		}(v)

		ImMut.Lock()
		out[v["key"]] = fmt.Sprintf(`%v`, found)
		ImMut.Unlock()
		i++
	}

	return out, nil
}

func SolvePicSelect(toSolve map[string]map[string]string, prompt, target string) (map[string]HashData, error) {
	imgs := map[string]string{}
	out := map[string]HashData{}

	j := 0
	for _, v := range toSolve {
		imgs[fmt.Sprintf("%d", j)] = v["body"]
		j++
	}

	jsonBody, err := json.Marshal(Base64JSON{
		Images:  imgs,
		Target:  prompt,
		Method:  "hcaptcha_base64",
		Sitekey: "4c672d35-0701-42b2-88c3-78380b0db560",
		Site:    "discord.com",
		Ln:      "en",
		Type:    "bbox",
	})
	if err != nil {
		return out, err
	}

	req, err := http.NewRequest("POST", proapi, bytes.NewBuffer(jsonBody))
	if err != nil {
		return out, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("apikey", apikey)

	resp, err := Client.Do(req)
	if err != nil {
		return out, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("nocap2", err)
		}
	}(resp.Body)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, err
	}

	var answer NoCapAnswerSelect
	if err := json.Unmarshal(bodyBytes, &answer); err != nil {
		return out, err
	}

	if answer.Status == "new" {
		time.Sleep(time.Second)

		req, err := http.NewRequest("GET", answer.Url, nil)
		if err != nil {
			return out, err
		}

		req.Header.Set("Content-type", "application/json")
		req.Header.Set("apikey", apikey)

		resp, err := Client.Do(req)
		if err != nil {
			return out, err
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("nocap2", err)
			}
		}(resp.Body)

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return out, err
		}

		if err := json.Unmarshal(bodyBytes, &answer); err != nil {
			return out, err
		}
	}

	if answer.Status == "skip" {
		return out, fmt.Errorf("cant solve images (skip)")
	}

	if len(answer.Answers) == 0 {
		return out, fmt.Errorf("can solve images, (0 found)")
	}

	i := 0
	for _, v := range toSolve {
		go func(i int, v map[string]string) {
			AsMut.Lock()
			defer AsMut.Unlock()

			currentValue, _ := Selectlist.Load(target)
			newHashData := HashData{
				Hash: v["hash"],
				X:    answer.Answers[i][0],
				Y:    answer.Answers[i][1],
			}
			var updatedValue []HashData
			if currentValue != nil {
				updatedValue = append(currentValue.([]HashData), newHashData)
			} else {
				updatedValue = []HashData{newHashData}
			}
			Selectlist.Store(target, updatedValue)

			utils.AppendLine(fmt.Sprintf("%s,%s,%d,%d", v["hash"], target, answer.Answers[i][0], answer.Answers[i][1]), "area_hash.csv")
		}(i, v)

		if i > len(answer.Answers) {
			fmt.Println("doesn't got all answers")
			break
		}

		ImMut.Lock()
		out[v["key"]] = HashData{
			Hash: v["hash"],
			X:    answer.Answers[i][0],
			Y:    answer.Answers[i][1],
		}
		ImMut.Unlock()
		i++

		if i > len(answer.Answers) {
			fmt.Println("doesn't got all answers")
			break
		}
	}

	return out, nil
}
