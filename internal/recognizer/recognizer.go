package recognizer

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cespare/xxhash"
	"github.com/zenthangplus/goccm"
)

func NewRecognizer(Type, Question string, Requester map[string]map[string]string, Task []TaskList) *Recognizer {
	return &Recognizer{
		TaskType:  Type,
		Question:  Question,
		TaskList:  Task,
		Requester: Requester,
		OMut:      sync.RWMutex{},
		TMut:      sync.RWMutex{},
	}
}

func (R *Recognizer) Recognize() (*SolveRepsonse, error) {
	var solved *SolveRepsonse
	var err error

	switch R.TaskType {
	case "image_label_binary":
		if strings.Contains("Please click each image containing a ", R.Question) {
			R.Target = strings.ReplaceAll(strings.Split(R.Question, "Please click each image containing a ")[1], " ", "_")
		} else {
			R.Target = strings.ReplaceAll(strings.Split(R.Question, "Please click each image containing ")[1], " ", "_")
		}

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
	default:
		err = fmt.Errorf("invalid task-type: %v", R.TaskType)
	}

	return solved, err
}

func (R *Recognizer) HashExistBinary(hashString string) (bool, error) {
	HMut.Lock()
	defer HMut.Unlock()

	hash, exist := Hashlist[R.Target]
	if exist {
		for _, h := range hash {
			if h == hashString {
				return true, nil
			}
		}
	}

	for prompt, hash := range Hashlist {
		if prompt != R.Target {
			for _, h := range hash {
				if h == hashString {
					return false, nil
				}
			}
		}
	}

	return false, fmt.Errorf("hash not solved yet")
}

func (R *Recognizer) DownloadAndCheckBinary(Url string) (bool, string, error) {
	Ccm.Wait()
	resp, err := http.Get(Url)
	Ccm.Done()
	if err != nil {
		return false, "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
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

func (R *Recognizer) LabelBinary() (*SolveRepsonse, error) {
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

				for {
					Ccm.Wait()
					resp, err := http.Get(t.DatapointURI)
					Ccm.Done()

					if err != nil {
						continue
					}
					defer func(Body io.ReadCloser) {
						err := Body.Close()
						if err != nil {
							fmt.Println(err)
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

	response, err := SolvePic(toResolve, R.Question, R.Target)
	if err != nil {
		return nil, err
	}

	R.OMut.Lock()
	for k, v := range response {
		out[k] = v
	}
	R.OMut.Unlock()

	return &SolveRepsonse{
		Success: true,
		Data:    out,
	}, nil
}

func (R *Recognizer) HashExistSelect(hashString string) (*HashData, error) {
	HMut.Lock()
	defer HMut.Unlock()

	hash, exist := Selectlist[R.Target]
	if exist {
		for _, k := range hash {
			if hashString == k.Hash {
				return &k, nil
			}
		}
	}

	for prompt, hash := range Selectlist {
		if prompt != R.Target {
			for _, k := range hash {
				if hashString == k.Hash {
					return &k, nil
				}
			}
		}
	}

	return nil, fmt.Errorf("hash not solved yet")
}

func (R *Recognizer) DownloadAndCheckSelect(Url string) (*HashData, string, error) {
	Ccm.Wait()
	resp, err := http.Get(Url)
	Ccm.Done()

	if err != nil {
		return nil, "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	buff := make([]byte, 650)
	_, err = io.ReadFull(resp.Body, buff)
	if err != nil {
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

func (R *Recognizer) LabelAreaSelect() (*SolveRepsonse, error) {
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

			data, hash, err = R.DownloadAndCheckSelect(t.DatapointURI)
			for {
				if err != nil && err.Error() != "hash not solved yet" {
					fmt.Println("blank hash", err)
					time.Sleep(time.Second)
					continue
				}

				for {
					Ccm.Wait()
					resp, err := http.Get(t.DatapointURI)
					Ccm.Done()
					if err != nil {
						continue
					}
					defer func(Body io.ReadCloser) {
						err := Body.Close()
						if err != nil {
							fmt.Println(err)
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

	response, err := SolvePicSelect(toResolve, R.Question, R.Target)
	if err != nil {
		return nil, err
	}

	if len(response) < 1 {
		return &SolveRepsonse{
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

	return &SolveRepsonse{
		Success: true,
		Data:    out,
	}, nil
}
