package tools

import (
	"gopkg.in/olivere/elastic.v3"
	"log"
	"strconv"
)

var (
	eshost = beego.AppConfig.String("eshost")
	esport = beego.AppConfig.String("esport")
	client *elastic.Client
)

type ESHelper struct{}

func (e *ESHelper)GetClient() *elastic.Client{
	if client == nil {
		client, err = elastic.NewClient(elastic.SetURL("http://" + eshost + ":" + esport))
		if err != nil{
			log.Println(err.Error())
			return nil
		}
	}
	return client
}

func (e *ESHelper)CreateByBulk() {
	var arr = [500]stirng
	for i int = 0; i < 500; i++ {
		arr[i] = "person"+ strconv.Itoa(i)
	}
}
