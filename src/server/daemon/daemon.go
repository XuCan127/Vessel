package daemon

import (
	"Vessel/src/common/term"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

type Daemon struct {
	TcpConfig term.TcpConfig
}

func (daemon *Daemon) ActivateListeners() error {
	// 启动 HTTP 监听器并启动 HTTP 服务
	httpListener, err := daemon.newHTTPListener()
	if err != nil {
		return err
	}
	if httpListener != nil {
		go func() {
			if err := daemon.serveHTTP(httpListener); err != nil {
				logrus.Errorf(term.ErrorFmt, "startDemon", err)
				daemon.Stop(httpListener)
			}
		}()
	}
	return nil
}

func (daemon *Daemon) newHTTPListener() (net.Listener, error) {
	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: daemon.TcpConfig.IP, Port: daemon.TcpConfig.Port})
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func (daemon *Daemon) serveHTTP(listener net.Listener) error {
	// 创建 HTTP 服务
	httpServer := &http.Server{
		Handler: daemon.newHTTPRouter(),
	}

	// 启动 HTTP 服务
	if err := httpServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (daemon *Daemon) newHTTPRouter() *mux.Router {
	// 创建 HTTP 路由器
	router := mux.NewRouter()

	// 添加中间件函数进行验证
	router.Use(daemon.authenticateMiddleware)

	// 注册 HTTP 路由
	router.HandleFunc("/version", daemon.VersionHandler).Methods("GET")
	router.HandleFunc("/imageBase/add/{name}/{path}", daemon.ImageBaseAddHandler).Methods("POST")
	router.HandleFunc("/imageBase/remove/{imageId}", daemon.ImageBaseRemoveHandler).Methods("POST")
	router.HandleFunc("/imageBase/list", daemon.ImageBaseListHandler).Methods("POST")

	router.HandleFunc("/containers/launch/{path}", daemon.ContainerLaunchHandler).Methods("POST")
	router.HandleFunc("/containers/ps", daemon.ContainerPSHandler).Methods("POST")

	return router
}

func (daemon *Daemon) Stop(listener net.Listener) {
	err := listener.Close()
	if err != nil {
		return
	}
}
