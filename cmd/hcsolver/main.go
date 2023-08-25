package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Implex-ltd/hcsolver/internal/hcaptcha"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func Solve(config *hcaptcha.Config) (*TaskResponse, error) {
	st := time.Now()
	task := &TaskResponse{}

	hc, err := hcaptcha.NewHcaptcha(config)
	if err != nil {
		return nil, err
	}

	var response *hcaptcha.ResponseCheckCaptcha

	site, err := hc.CheckSiteConfig()
	if err != nil {
		logger.Error("checksiteconfig",
			zap.String("error", err.Error()),
		)
		task.Data.Err.Errors = append(task.Data.Err.Errors, err.Error())
		return task, err
	}

	captcha, err := hc.GetChallenge(site)
	if err != nil {
		logger.Error("getchallenge",
			zap.String("error", err.Error()),
		)
		task.Data.Err.Errors = append(task.Data.Err.Errors, err.Error())
		return task, err
	}

	task.Data.Task.TaskType = captcha.RequestType
	task.Data.Task.TaskPrompt = captcha.RequesterQuestion.En

	if captcha.RequestType != "image_label_binary" {
		logger.Error("getchallenge",
			zap.String("error", fmt.Sprintf("invalid request-type: %s", captcha.RequestType)),
		)
		task.Data.Err.Errors = append(task.Data.Err.Errors, fmt.Sprintf("invalid request-type: %s", captcha.RequestType))
		return task, fmt.Errorf("invalid request-type: %s", captcha.RequestType)
	}

	response, err = hc.CheckCaptcha(captcha)
	if err != nil {
		logger.Error("checkcaptcha",
			zap.String("error", err.Error()),
		)
		task.Data.Err.Errors = append(task.Data.Err.Errors, err.Error())
		return task, err
	}

	if response.Pass {
		task.Data.Token.CAPTCHAKey = response.GeneratedPassUUID
		task.Data.Token.Expiration = int64(response.Expiration)
		log.Println("ai:", hc.AnswerProcessing.Milliseconds(), "hsw:", hc.HswProcessing.Milliseconds())
	}

	return &TaskResponse{
		Success: true,
		Data: Data{
			Task: Task{
				TaskType:   task.Data.Task.TaskType,
				TaskPrompt: task.Data.Task.TaskPrompt,
			},
			Metrics: Metrics{
				TaskProcess: time.Since(st).Milliseconds(),
				HswProcess:  hc.HswProcessing.Milliseconds(),
				ImgProcess:  hc.AnswerProcessing.Milliseconds(),
				StartTime:   st.Unix(),
				TTLProcess:  time.Since(st).Milliseconds(),
			},
			Token: Token{
				CAPTCHAKey: task.Data.Token.CAPTCHAKey,
				Expiration: task.Data.Token.Expiration,
			},
			Err: Err{
				Retry:  task.Data.Err.Retry,
				Errors: task.Data.Err.Errors,
			},
		},
	}, nil
}

func HandlerSolve(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestBody BodyNewSolveTask

	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	logger.Info("new task",
		//	zap.String("useragent", requestBody.UserAgent),
		zap.String("sitekey", requestBody.SiteKey),
		zap.String("domain", requestBody.Domain),
	//	zap.String("proxy", requestBody.Proxy),
	)

	resp, err := Solve(&hcaptcha.Config{
		UserAgent: requestBody.UserAgent,
		SiteKey:   requestBody.SiteKey,
		Domain:    requestBody.Domain,
		Proxy:     requestBody.Proxy,
	})

	if err != nil || !resp.Success {
		logger.Error("solve",
			zap.String("error", fmt.Sprintf("%s", err.Error())),
			zap.Bool("success", resp.Success),
		)

		resp.Success = false

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(resp)
		return
	}

	logger.Info("solved",
		zap.String("key", resp.Data.Token.CAPTCHAKey[:15]),
		zap.String("prompt", resp.Data.Task.TaskPrompt),
		//zap.String("type", resp.Data.Task.TaskType),
		zap.Int64("hsw_process", resp.Data.Metrics.HswProcess),
		zap.Int64("img_process", resp.Data.Metrics.ImgProcess),
		//zap.Int64("ttl_process", resp.Data.Metrics.TTLProcess),
		zap.Int64("task_process", resp.Data.Metrics.TaskProcess),

		//zap.Int64("retry", resp.Data.Err.Retry),

		//zap.String("useragent", requestBody.UserAgent),
		//zap.String("sitekey", requestBody.SiteKey),
		//zap.String("domain", requestBody.Domain),
		//zap.String("proxy", requestBody.Proxy),
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func main() {
	LoadSettings()

	defer func() {
		logger.Sync()
		if err := logger.Core().Sync(); err != nil {
			log.Fatalf("Erreur lors de la synchronisation du fichier de log : %v", err)
		}
	}()

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Post("/solve", HandlerSolve)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", 1337), r); err != nil {
		panic(err)
	}
}
