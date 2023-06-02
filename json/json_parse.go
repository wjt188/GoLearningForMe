package main

//GO解析Json文件的示例
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// 定义配置文件解析后的结构
type MongoConfig struct {
	MongoAddr       string
	MongoPoolLimit  int
	MongoDb         string
	MongoCollection string
}
type Config struct {
	Port  string
	Mongo MongoConfig
}
type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}
func (js *JsonStruct) Load(filename string, v interface{}) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}

func main() {
	JsonParse := NewJsonStruct()
	//v := Config{}
	var v = new(Config)
	JsonParse.Load("./json/json_parse.json", &v)
	fmt.Println(v.Port)
	fmt.Println(v.Mongo.MongoDb)
	fmt.Println(v.Mongo.MongoAddr)
	fmt.Println(v.Mongo.MongoPoolLimit)
	fmt.Println(v.Mongo.MongoCollection)
	fmt.Printf("%v", *v)

}
