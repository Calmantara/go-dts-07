package jsonplaceholder

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Calmantara/go-common/pkg/logger"
)

type JsonPlaceholderClient interface {
	GetTodo(ctx context.Context, todoNum int) (res JsonPlaceholderResp, err error)
}

type JsonPlaceholderClientImpl struct {
	baseUrl string
}

func NewJsonPlaceholderClient() JsonPlaceholderClient {
	return &JsonPlaceholderClientImpl{
		baseUrl: "https://jsonplaceholder.typicode.com/todos",
	}
}

func (ph *JsonPlaceholderClientImpl) GetTodo(ctx context.Context, todoNum int) (resMdl JsonPlaceholderResp, err error) {
	logger.Info(ctx, "GetTodo invoked")

	req, _ := http.NewRequest("GET", fmt.Sprintf(ph.baseUrl+"/%v", todoNum), nil)
	req.Header.Add("Accept", "*/*")
	// sending http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		err = errors.New("request not success")
		return
	}
	err = json.Unmarshal(body, &resMdl)
	return
}
