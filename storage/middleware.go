package storage

type InterceptorContext struct {
	Value string
	Key   string
	Err   error
}

type Interceptor interface {
	Intercept(method string, i *InterceptorContext)
}
