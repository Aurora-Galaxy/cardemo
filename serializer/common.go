package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  error       `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}
