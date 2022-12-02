package constant

type ResponseCode int

const (
	SUCCESS      ResponseCode = 20000 //成功
	Error        ResponseCode = 1     //参数校验失败 暂时20000
	Handle_Error ResponseCode = 101   //处理失败
	TokenFail    ResponseCode = 102   // token校验失败
	NotFind      ResponseCode = 404   // 页面未找到
	ServerFail   ResponseCode = 500   // 服务器错误

)
