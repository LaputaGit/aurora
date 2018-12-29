package main

import (
	"log"
	"sync"
)

/**
	路由
 */
type Route struct {
	mutex sync.Mutex
	clients map[int64]*Client
}

//创建一个route
func NewRoute() *Route{
	route := new(Route)
	route.clients = make(map[int64]*Client)
	return route
}

//把client加入route
func (route *Route) AddClient(client *Client){
	route.mutex.Lock()
	defer route.mutex.Unlock()
	if _,ok := route.clients[client.uid]; ok{
		log.Println("client has exists")
	}
	route.clients[client.uid] = client;
}

//删除client
func (route *Route) RemoveClient(client *Client){
	route.mutex.Lock()
	defer route.mutex.Unlock()
	if _,ok := route.clients[client.uid]; ok{
		delete(route.clients,client.uid)
	}else{
		log.Println("client no exists")
	}
}

//个根据uid查找client
func (route *Route) FindClient(uid int64) *Client{
	route.mutex.Lock()
	defer route.mutex.Unlock()
	c,ok := route.clients[uid];
	if ok{
		return c
	}else{
		return nil
	}
}

