package recognizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

var (
	ImMut, SmMut, AsMut sync.RWMutex
)

var (
	nocap  = true
	apikey = "rorm-8473d243-790d-9184-3fa2-76e4ff8424df"
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

	client := &http.Client{}

	Ccm.Wait()
	resp, err := client.Do(req)
	Ccm.Done()
	if err != nil {
		return imgs, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
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

				Hashlist[target] = append(Hashlist[target], v["hash"])

				file, err := os.OpenFile("../../assets/hash.csv", os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return
				}
				defer file.Close()

				file.WriteString(fmt.Sprintf("%s,%s", v["hash"], target) + "\n")
			} else {
				SmMut.Lock()
				defer SmMut.Unlock()
				Hashlist["not_"+target] = append(Hashlist["not_"+target], v["hash"])

				file, err := os.OpenFile("../../assets/hash.csv", os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return
				}
				defer file.Close()

				file.WriteString(fmt.Sprintf("%s,not_%s", v["hash"], target) + "\n")
			}
		}(v)

		ImMut.Lock()
		out[v["key"]] = fmt.Sprintf(`"%v"`, found)
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

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return out, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
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

	if len(answer.Answers) == 0 {
		return out, fmt.Errorf("cant solve")
	}

	i := 0
	for _, v := range toSolve {
		go func(i int, v map[string]string) {
			AsMut.Lock()
			defer AsMut.Unlock()

			Selectlist[target] = append(Selectlist[target], HashData{
				Hash: v["hash"],
				X:    answer.Answers[i][0],
				Y:    answer.Answers[i][1],
			})

			file, err := os.OpenFile("../../assets/area_hash.csv", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("%s,%s,%d,%d", v["hash"], target, answer.Answers[i][0], answer.Answers[i][1]) + "\n")
		}(i, v)

		ImMut.Lock()
		out[v["key"]] = HashData{
			Hash: v["hash"],
			X:    answer.Answers[i][0],
			Y:    answer.Answers[i][1],
		}
		ImMut.Unlock()
		i++
	}

	return out, nil
}