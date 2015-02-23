package unet
package unet


import (
import (
	"bytes"
	"bytes"
	"fmt"
	"fmt"
type XsdtString struct{ string }


	"net/http"
	"net/http"
	"strings"
	"strings"


	"github.com/go-utils/ufs"
	"github.com/go-utils/ufs"
type XsdtString struct{ string }	"github.com/go-utils/ugo"
)
)


//	Returns a human-readable URL representation of the specified TCP address.
//	Returns a human-readable URL representation of the specified TCP address.
//
//
//	Examples:
//	Examples:
//
//
//	`unet.Addr("http", ":8080")` = `http://localhost:8080`
//	`unet.Addr("http", ":8080")` = `http://localhost:8080`
//
//
//	`unet.Addr("https", "testserver:9090")` = `https://testserver:9090`
//	`unet.Addr("https", "testserver:9090")` = `https://testserver:9090`
//
//
//	`unet.Addr("http", ":http")` = `http://localhost`
//	`unet.Addr("http", ":http")` = `http://localhost`
//
//
//	`unet.Addr("https", "demomachine:https")` = `https://demomachine`
//	`unet.Addr("https", "demomachine:https")` = `https://demomachine`
func Addr(protocol, tcpAddr string) (fullAddr string) {
func Addr(protocol, tcpAddr string) (fullAddr string) {
	localhost := ugo.HostName()
	localhost := ugo.HostName()
	both := strings.Split(tcpAddr, ":")
	both := strings.Split(tcpAddr, ":")
	if len(both) < 1 {
	if len(both) < 1 {
		both = []string{localhost}
		both = []string{localhost}
	} else if len(both[0]) == 0 {
	} else if len(both[0]) == 0 {
		both[0] = localhost
		both[0] = localhost
	}
	}
	if len(both) > 1 {
	if len(both) > 1 {
		if both[1] == protocol {
		if both[1] == protocol {
			both[1] = ""
			both[1] = ""
		}
		}
		if len(both[1]) == 0 {
		if len(both[1]) == 0 {
			both = both[:1]
			both = both[:1]
		}
		}
	}
	}
	if fullAddr = strings.Join(both, ":"); len(protocol) > 0 {
	if fullAddr = strings.Join(both, ":"); len(protocol) > 0 {
		fullAddr = fmt.Sprintf("%s://%s", protocol, fullAddr)
		fullAddr = fmt.Sprintf("%s://%s", protocol, fullAddr)
	}
	}
	return
	return
}
}


//	Downloads a remote file at the specified (`net/http`-compatible) `srcFileUrl` to the specified `dstFilePath`.
//	Downloads a remote file at the specified (`net/http`-compatible) `srcFileUrl` to the specified `dstFilePath`.
func DownloadFile(srcFileUrl, dstFilePath string) (err error) {
func DownloadFile(srcFileUrl, dstFilePath string) (err error) {
	var rc io.ReadCloser
	var rc io.ReadCloser
	if rc, err = OpenRemoteFile(srcFileUrl); err == nil {
	if rc, err = OpenRemoteFile(srcFileUrl); err == nil {
		defer rc.Close()
		defer rc.Close()
		ufs.SaveToFile(rc, dstFilePath)
		ufs.SaveToFile(rc, dstFilePath)
	}
	}
	return
	return
}
}


//	Opens a remote file at the specified (`net/http`-compatible) `srcFileUrl` and returns its `io.ReadCloser`.
//	Opens a remote file at the specified (`net/http`-compatible) `srcFileUrl` and returns its `io.ReadCloser`.
func OpenRemoteFile(srcFileUrl string) (src io.ReadCloser, err error) {
func OpenRemoteFile(srcFileUrl string) (src io.ReadCloser, err error) {
	var resp *http.Response
	var resp *http.Response
	if resp, err = new(http.Client).Get(srcFileUrl); (err == nil) && (resp != nil) {
	if resp, err = new(http.Client).Get(srcFileUrl); (err == nil) && (resp != nil) {
		src = resp.Body
		src = resp.Body
	}
	}
	return
	return
}
}


//	Implements `http.ResponseWriter` with a `bytes.Buffer`.
//	Implements `http.ResponseWriter` with a `bytes.Buffer`.
type ResponseBuffer struct {
type ResponseBuffer struct {
	//	Used to implement the `http.ResponseWriter.Write` method.
	//	Used to implement the `http.ResponseWriter.Write` method.
	bytes.Buffer
	bytes.Buffer


	//	Used to implement the `http.ResponseWriter.Header` method.
	//	Used to implement the `http.ResponseWriter.Header` method.
	Resp http.Response
	Resp http.Response
}
}


//	Returns `me.Resp.Header`.
//	Returns `me.Resp.Header`.
func (me *ResponseBuffer) Header() http.Header {
func (me *ResponseBuffer) Header() http.Header {
	return me.Resp.Header
	return me.Resp.Header
}
}


//	No-op -- currently, headers aren't written to the underlying `bytes.Buffer`.
//	No-op -- currently, headers aren't written to the underlying `bytes.Buffer`.
func (_ *ResponseBuffer) WriteHeader(_ int) {
func (_ *ResponseBuffer) WriteHeader(_ int) {
}
}
