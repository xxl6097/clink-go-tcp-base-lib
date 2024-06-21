package server

import (
	"github.com/aceld/zinx/zconf"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

func ServeTCP(funcConn func(begin bool, conn ziface.IConnection), funcConf func(conf *zconf.Config), funcDecoder func(func(ziface.IDecoder)), funcRouter func(func(uint32, ziface.IRouter))) {
	conf := &zconf.Config{}
	funcConf(conf)
	s := znet.NewUserConfServer(conf)
	//注册链接hook回调函数
	s.SetOnConnStart(func(connection ziface.IConnection) {
		funcConn(true, connection)

	})
	s.SetOnConnStop(func(connection ziface.IConnection) {
		funcConn(false, connection)
	})

	funcDecoder(func(decoder ziface.IDecoder) {
		s.SetDecoder(decoder)
	})

	funcRouter(func(msgID uint32, router ziface.IRouter) {
		s.AddRouter(msgID, router)
	})
	//开启服务
	s.Serve()
}

func test() {
	ServeTCP(func(begin bool, conn ziface.IConnection) {
		// 1. 设置链接属性
	}, func(conf *zconf.Config) {
		// 2. 设置服务器配置
	}, func(f func(ziface.IDecoder)) {
		// 3. 设置解码器
	}, func(f func(uint32, ziface.IRouter)) {
		// 4. 设置路由
	})
}
