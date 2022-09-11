package baseNetWork

import (
	"net/http"
)

// NetWork 继承基类
type NetWork struct {
	BaseNetWork
	req *http.Client
}

func (n NetWork) name() {

}
