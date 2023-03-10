package response

import "fmt"

// NewMessage new实例
func NewMessage(data interface{}) *Message {
	code := -1
	return &Message{
		Code:    &code,
		Message: data,
	}
}

// Data to be used by controllers.
type Message struct {
	RequestId string      `json:"request_id,omitempty"` // 请求Id
	Code      *int        `json:"code"`                 // 自定义返回码  0:表示正常
	Type      string      `json:"type,omitempty"`       // 数据类型, 可以缺省
	Namespace string      `json:"namespace,omitempty"`  // 异常的范围
	Reason    string      `json:"reason,omitempty"`     // 异常原因
	Recommend string      `json:"recommend,omitempty"`  // 推荐链接
	Describe  string      `json:"describe,omitempty"`   // 关于这次响应的说明信息
	Message   interface{} `json:"message,omitempty"`    // 返回的具体数据
	Meta      interface{} `json:"meta,omitempty"`       // 数据meta
}

func (d *Message) Error() error {
	if d.Code == nil {
		return nil
	}
	if *d.Code != 0 {
		return fmt.Errorf("code: %d message: %s, meta: %s",
			d.Code, d.Message, d.Meta)
	}
	return nil
}

// Option configures how we set up the data.
type Option interface {
	Apply(*Message)
}

func newFuncOption(f func(*Message)) Option {
	return &funcOption{
		f: f,
	}
}

type funcOption struct {
	f func(*Message)
}

func (fdo *funcOption) Apply(do *Message) {
	fdo.f(do)
}

func WithRequestId(rid string) Option {
	return newFuncOption(func(o *Message) {
		o.RequestId = rid
	})
}

func WithRecommend(msg string) Option {
	return newFuncOption(func(o *Message) {
		o.Recommend = msg
	})
}

func WithMeta(meta interface{}) Option {
	return newFuncOption(func(o *Message) {
		o.Meta = meta
	})
}
