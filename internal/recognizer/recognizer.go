package recognizer

import (
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/cespare/xxhash"
	"github.com/zenthangplus/goccm"
)

var (
	re      = regexp.MustCompile(`Please click (on all|each) image(s)? containing (a\s)?(.+)`)
	badCode = map[string]string{
		"а": "a",
		"е": "e",
		"e": "e",
		"i": "i",
		"і": "i",
		"ο": "o",
		"с": "c",
		"ԁ": "d",
		"ѕ": "s",
		"һ": "h",
		"у": "y",
		"р": "p",
		"ϳ": "j",
		"х": "x",
		"ー": "一",
		"土": "士",
	}
)

func NewRecognizer(Proxy, Type, Question string, Requester map[string]map[string]string, Task []TaskList) (*Recognizer, error) {
	proxy, err := url.Parse(Proxy)
	if err != nil {
		return nil, err
	}

	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 200
	t.MaxConnsPerHost = 500
	t.MaxIdleConnsPerHost = 200

	client := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	return &Recognizer{
		TaskType:  Type,
		Question:  Question,
		TaskList:  Task,
		Requester: Requester,
		OMut:      sync.RWMutex{},
		TMut:      sync.RWMutex{},
		Http:      client,
	}, nil
}

func (R *Recognizer) ExtractTarget(question string) string {
	matches := re.FindStringSubmatch(question)
	out := R.Question

	if len(matches) >= 4 {
		out = strings.ReplaceAll(matches[4], " ", "_")
	}

	return out
}

func (R *Recognizer) LabelCleaning(rawLabel string) string {
	runes := []rune(rawLabel)

	for i, r := range runes {
		if replacement, ok := badCode[string(r)]; ok {
			runes[i] = []rune(replacement)[0]
		}
	}

	cleanLabel := string(runes)
	return cleanLabel
}

func (R *Recognizer) Recognize() (*SolveResponse, error) {
	var solved *SolveResponse
	var err error

	R.Question = R.LabelCleaning(R.Question)

	switch R.TaskType {
	case "image_label_binary":
		R.Target = R.ExtractTarget(R.Question)

		solved, err = R.LabelBinary()
	case "image_label_area_select":
		var entity string

		for _, innerMap := range R.Requester {
			if value, ok := innerMap["en"]; ok {
				entity = value
			}
		}

		R.Target = strings.ReplaceAll(strings.Split(R.Question, "Please click on the ")[1], " ", "_")
		R.EntityType = entity

		solved, err = R.LabelAreaSelect()
	case "text_free_entry":
		solved, err = R.TextFreeEntry()
	default:
		err = fmt.Errorf("invalid task-type: %v", R.TaskType)
	}

	return solved, err
}

func (R *Recognizer) HashExistBinary(hashString string) (bool, error) {
	hashesInterface, _ := Hashlist.Load(R.Target)

	if hashesInterface != nil {
		hashes := hashesInterface.([]string)
		for _, h := range hashes {
			if h == hashString {
				return true, nil
			}
		}
	}

	Hashlist.Range(func(key, value interface{}) bool {
		prompt := key.(string)
		if prompt != R.Target {
			hashes := value.([]string)
			for _, h := range hashes {
				if h == hashString {
					return false
				}
			}
		}
		return true
	})

	return false, fmt.Errorf("hash not solved yet")
}

func (R *Recognizer) DownloadAndCheckBinary(Url string) (bool, string, error) {
	resp, err := R.Http.Get(Url)

	if err != nil {
		return false, "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("dl_and_check", err)
		}
	}(resp.Body)

	buff := make([]byte, 650)
	_, err = io.ReadFull(resp.Body, buff)
	if err != nil {
		return false, "", err
	}

	hash := xxhash.Sum64(buff)
	hashString := fmt.Sprintf("%x", hash)

	exist, err := R.HashExistBinary(hashString)
	if err != nil {
		return false, hashString, err
	}

	return exist, hashString, nil
}

func (R *Recognizer) LabelBinary() (*SolveResponse, error) {
	out := map[string]any{}
	toResolve := map[string]map[string]string{}

	c := goccm.New(len(R.TaskList))

	for _, task := range R.TaskList {
		c.Wait()

		go func(t TaskList) {
			defer c.Done()

			var is bool
			var err error
			var hash string

			for {
				is, hash, err = R.DownloadAndCheckBinary(t.DatapointURI)
				if err != nil && err.Error() != "hash not solved yet" {
					fmt.Println("blank hash", err)
					time.Sleep(time.Second)
					continue
				}

				if err != nil {
					if err.Error() == "hash not solved yet" && nocap {
						for {
							resp, err := R.Http.Get(t.DatapointURI)

							if err != nil {
								continue
							}
							defer func(Body io.ReadCloser) {
								err := Body.Close()
								if err != nil {
									fmt.Println("xd1", err)
								}
							}(resp.Body)

							body, err := io.ReadAll(resp.Body)
							if err != nil {
								continue
							}

							R.TMut.Lock()
							toResolve[t.TaskKey] = map[string]string{
								"body": base64.StdEncoding.EncodeToString(body),
								"hash": hash,
								"key":  t.TaskKey,
							}
							R.TMut.Unlock()
							break
						}
					}
				}
				break
			}

			if is {
				R.OMut.Lock()
				out[t.TaskKey] = fmt.Sprintf(`%v`, is)
				R.OMut.Unlock()
			}
		}(task)
	}

	c.WaitAllDone()

	if nocap && len(toResolve) > 1 {
		response, err := SolvePic(toResolve, R.Question, R.Target)
		if err != nil {
			return nil, err
		}

		R.OMut.Lock()
		for k, v := range response {
			out[k] = v
		}
		R.OMut.Unlock()
	}

	return &SolveResponse{
		Success: true,
		Data:    out,
	}, nil
}

func (R *Recognizer) HashExistSelect(hashString string) (*HashData, error) {
	// Load the Selectlist value for the target prompt.
	value, _ := Selectlist.Load(R.Target)

	if value != nil {
		hashDataList := value.([]HashData)
		for _, k := range hashDataList {
			if hashString == k.Hash {
				return &k, nil
			}
		}
	}

	// Iterate over the sync.Map to check other prompts.
	Selectlist.Range(func(key, value interface{}) bool {
		prompt := key.(string)
		if prompt != R.Target {
			hashDataList := value.([]HashData)
			for _, k := range hashDataList {
				if hashString == k.Hash {
					return false
				}
			}
		}
		return true
	})

	return nil, fmt.Errorf("hash not solved yet")
}

func (R *Recognizer) DownloadAndCheckSelect(Url string) (*HashData, string, error) {
	resp, err := R.Http.Get(Url)

	if err != nil {
		return nil, "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("dl_and_check0", err)
		}
	}(resp.Body)

	buff := make([]byte, 650)
	_, err = io.ReadFull(resp.Body, buff)
	if err != nil {
		fmt.Println(resp.StatusCode)
		return nil, "", err
	}

	hash := xxhash.Sum64(buff)
	hashString := fmt.Sprintf("%x", hash)

	data, err := R.HashExistSelect(hashString)
	if err != nil {
		return nil, hashString, err
	}

	return data, hashString, nil
}

func (R *Recognizer) LabelAreaSelect() (*SolveResponse, error) {
	out := make(map[string][]map[string]interface{})
	toResolve := map[string]map[string]string{}

	c := goccm.New(len(R.TaskList))

	for _, task := range R.TaskList {
		c.Wait()

		go func(t TaskList) {
			defer c.Done()

			var data *HashData
			var err error
			var hash string

			for {
				data, hash, err = R.DownloadAndCheckSelect(t.DatapointURI)

				if err != nil && err.Error() != "hash not solved yet" {
					fmt.Println("blank hash", err)
					time.Sleep(time.Second)
					continue
				}

				if err != nil {
					if err.Error() == "hash not solved yet" && nocap {
						fmt.Println("nocap")
						for {
							resp, err := R.Http.Get(t.DatapointURI)
							if err != nil {
								continue
							}
							defer func(Body io.ReadCloser) {
								err := Body.Close()
								if err != nil {
									fmt.Println("nocap2", err)
								}
							}(resp.Body)

							body, err := io.ReadAll(resp.Body)
							if err != nil {
								continue
							}

							R.TMut.Lock()
							toResolve[t.TaskKey] = map[string]string{
								"body": base64.StdEncoding.EncodeToString(body),
								"hash": hash,
								"key":  t.TaskKey,
							}
							R.TMut.Unlock()
							break
						}
					}
				}

				break
			}

			if data != nil {
				R.OMut.Lock()
				out[t.TaskKey] = []map[string]interface{}{
					{
						"entity_name": 0,
						"entity_type": R.EntityType,
						"entity_coords": []int{
							data.X,
							data.Y,
						},
					},
				}
				R.OMut.Unlock()
			}
		}(task)
	}

	c.WaitAllDone()

	if nocap && len(toResolve) > 1 {
		response, err := SolvePicSelect(toResolve, R.Question, R.Target)
		if err != nil {
			return nil, err
		}

		if len(response) < 1 {
			return &SolveResponse{
				Success: false,
			}, fmt.Errorf("cant recognize")
		}

		R.OMut.Lock()
		for k, v := range response {
			out[k] = []map[string]interface{}{
				{
					"entity_name": 0,
					"entity_type": R.EntityType,
					"entity_coords": []int{
						v.X,
						v.Y,
					},
				},
			}
		}
		R.OMut.Unlock()
	}

	return &SolveResponse{
		Success: true,
		Data:    out,
	}, nil
}

func (R *Recognizer) TextFreeEntry() (*SolveResponse, error) {
	answers := map[string]AnswerStruct{}
	resp := []string{"oui", "non"}

	for _, questions := range R.TaskList {
		res := resp[rand.Int()%len(resp)]
		question := strings.ReplaceAll(questions.DatapointText["fr"], " ", "_")

		val, ok := Answerlist.Load(question)

		if ok {
			res = val.(string)
		} else {
			file, err := os.OpenFile("../../assets/nop.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return nil, err
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("%s|%s", question, res+"\n"))
		}
		
		answers[questions.TaskKey] = AnswerStruct{
			Text: res,
		}
	}

	return &SolveResponse{
		Success: true,
		Data:    answers,
	}, nil
}
