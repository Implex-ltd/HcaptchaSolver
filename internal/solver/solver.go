package solver

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/cespare/xxhash"
	"go.uber.org/zap"
)

const (
	apikey = "rorm-8473d243-790d-9184-3fa2-76e4ff8424df"
	proapi = "https://pro.nocaptchaai.com/solve"
)

var (
	nocap = true
)

func HashExists(prompt string, contentHash uint64) bool {
	hashStr := fmt.Sprintf("%x", contentHash)

	hashlistMutex.RLock()
	defer hashlistMutex.RUnlock()

	hashes, exists := Hashlist[prompt]
	if exists {
		for _, h := range hashes {
			if h == hashStr {
				return true
			}
		}
	}
	return false
}

func ParallelHashExists(prompt string, contentHash uint64, wg *sync.WaitGroup, resultChan chan<- bool) {
	defer wg.Done()

	result := HashExists(prompt, contentHash)
	resultChan <- result
}

func ParallelHashExistsAllThreads(prompt string, contentHash uint64) bool {
	hashStr := fmt.Sprintf("%x", contentHash)

	hashlistMutex.RLock()
	defer hashlistMutex.RUnlock()

	for otherPrompt, hashes := range Hashlist {
		if otherPrompt != prompt && !strings.HasPrefix(prompt, "not_") {
			for _, h := range hashes {
				if h == hashStr {
					return true
				}
			}
		}
	}
	return false
}

func SolvePic(url, prompt string) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	base64_json := Base64JSON{
		Images: map[string]string{
			"0": base64.StdEncoding.EncodeToString(body),
		},
		Target:  "Please click each image containing a " + prompt,
		Method:  "hcaptcha_base64",
		Sitekey: "4c672d35-0701-42b2-88c3-78380b0db560",
		Site:    "discord.com",
		Ln:      "en",
	}
	jsonBody, err := json.Marshal(base64_json)
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", proapi, bytes.NewBuffer(jsonBody))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("apikey", apikey)

	client := &http.Client{}

	resp, err = client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var answer NoCapAnswer
	if err := json.Unmarshal(bodyBytes, &answer); err != nil {
		return false, err
	}

	if len(answer.Solution) > 0 {
		return true, nil
	}

	return false, nil
}

func DownloadAndClassify(url, key, prompt, fullprompt string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	st := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		results <- Result{Hash: "", Match: false, Err: err, Url: url, St: time.Since(st), Key: key}
		return
	}
	defer resp.Body.Close()

	buff := make([]byte, 650)
	_, err = io.ReadFull(resp.Body, buff)
	if err != nil {
		results <- Result{Hash: "", Match: false, Err: err, Url: url, St: time.Since(st), Key: key}
		return
	}

	contentHash := xxhash.Sum64(buff)
	hashStr := fmt.Sprintf("%x", contentHash)

	if HashExists(prompt, contentHash) {
		results <- Result{Hash: hashStr, Match: true, Err: nil, Url: url, St: time.Since(st), Key: key}
		return
	}

	if HashExists(fmt.Sprintf("not_%s", prompt), contentHash) {
		results <- Result{Hash: hashStr, Match: false, Err: nil, Url: url, St: time.Since(st), Key: key}
		return
	}

	if ParallelHashExistsAllThreads(prompt, contentHash) {
		results <- Result{Hash: hashStr, Match: false, Err: nil, Url: url, St: time.Since(st), Key: key}
		return
	}

	// if not solved
	if nocap {
		answer, err := SolvePic(url, fullprompt)
		if err != nil {
			results <- Result{Hash: fmt.Sprintf("%x", contentHash), Match: false, Err: nil, Url: url, St: time.Since(st), Key: key}
			return
		}

		if answer {
			go func() {
				mu.Lock()
				defer mu.Unlock()

				Hashlist[prompt] = append(Hashlist[prompt], fmt.Sprintf("%x", contentHash))

				file, err := os.OpenFile("../../assets/hash.csv", os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					return
				}
				defer file.Close()

				file.WriteString(fmt.Sprintf("%s,%s", fmt.Sprintf("%x", contentHash), prompt) + "\n")
			}()
		} else {
			mu.Lock()
			defer mu.Unlock()
			Hashlist["not_"+prompt] = append(Hashlist["not_"+prompt], fmt.Sprintf("%x", contentHash))

			file, err := os.OpenFile("../../assets/hash.csv", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				return
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("%s,not_%s", fmt.Sprintf("%x", contentHash), prompt) + "\n")
		}

		results <- Result{Hash: fmt.Sprintf("%x", contentHash), Match: answer, Err: nil, Url: url, St: time.Since(st), Key: key}
		return
	}

	results <- Result{Hash: fmt.Sprintf("%x", contentHash), Match: false, Err: nil, Url: url, St: time.Since(st), Key: key}
}

func Task(task *BodyNewSolveTask, logger *zap.Logger) *SolveRepsonse {
	results := make(chan Result, len(task.TaskList))
	t := time.Now()

	var prompt string
	if strings.Contains(task.Question, "Please click each image containing a ") {
		prompt = strings.ReplaceAll(strings.Split(task.Question, "Please click each image containing a ")[1], " ", "_")
	}

	var wg sync.WaitGroup
	for _, t := range task.TaskList {
		wg.Add(1)
		go DownloadAndClassify(t.DatapointURI, t.TaskKey, prompt, task.Question, results, &wg)
	}

	wg.Wait()
	close(results)

	resp := map[string]string{}

	for result := range results {
		resp[result.Key] = fmt.Sprintf("%v", result.Match)

		if result.Err != nil {
			fmt.Println("Image download failed:", result.Err)
			return nil
		}

		logger.Info("Image classified",
			zap.String("hash", result.Hash),
			zap.String("prompt", prompt),
			zap.Bool("match", result.Match),
			zap.Int64("st", result.St.Milliseconds()),
			zap.String("url", result.Url),
		)
	}

	logger.Info("Task classified",
		zap.Int64("st", time.Since(t).Milliseconds()),
		zap.String("prompt", prompt),
	)

	return &SolveRepsonse{
		Success: true,
		Data:    resp,
	}
}

func LoadHash(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		if len(parts) != 2 {
			continue
		}

		hash, prompt := parts[0], parts[1]

		if _, exists := Hashlist[prompt]; !exists {
			Hashlist[prompt] = []string{}
		}

		Hashlist[prompt] = append(Hashlist[prompt], hash)
		count++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
