package response

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/opengoats/goat/exception"
)

// GetDataFromBody 获取body中的数据
func GetMessageFromBody(body io.ReadCloser, v interface{}) error {
	defer body.Close()

	bytesB, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	data := NewMessage(v)

	if err := json.Unmarshal(bytesB, data); err != nil {
		return err
	}

	if data.Code == nil {
		return errors.New("reponse code is nil")
	}

	if *data.Code != 0 {
		return exception.NewAPIException(data.Namespace, *data.Code, data.Reason, data.Describe)
	}

	return nil
}
