package service

const HelloServiceName = "HelloService"

type HelloService interface {
	Hello(request *Request, reply *Response) error
}
