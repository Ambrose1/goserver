package baseNetWork

import "net/http"

type AMHTTPTYPE uint32

//enum
const (
	AMHTTPTYPE_GET AMHTTPTYPE = 1 << (32 - 1 - iota)
	AMHTTPTYPE_POST
	AMHTTPTYPE_PUT
	AMHTTPTYPE_DELETE
	AMHTTPTYPE_OPTION
	AMHTTPTYPE_HEAD
	AMHTTPTYPE_TRACE
	AMHTTPTYPE_CONNECT
)

// BaseNetWork base class
type BaseNetWork struct {
}

type FuncType *func(argc ...any) any

type NetWorkInterface interface {
	AMHTTPRequest(params any, url string, reqType AMHTTPTYPE) (resp *http.Response, err error)
	AMHTTPRequestWithBlock(params any, url string, reqType AMHTTPTYPE, block *FuncType) (resp *http.Response, err error)
}
