package main

import (
	"GoLearning/redis/search/common"
	"GoLearning/redis/search/utils"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"log"
	"reflect"
	"sort"
	"strings"
	"time"
)

// 创建一个redis类型的客户端结构体
type Client struct {
	Conn *redis.Client
}

// 初始化一个客户端
func NewClient(conn *redis.Client) *Client {
	return &Client{Conn: conn}
}

// 将内容中的非常用词过滤掉，其他词放到集合中
func Tokenize(content string) []string {
	words := utils.Set{}
	for _, match := range common.WORDSRE.FindAllString(strings.ToLower(content), -1) {
		if len(match) < 2 {
			continue
		}
		words.Add(match)
	}
	return words.Diff(&common.STOPWORDS)
}

// 对内容进行一个标记话的处理，将文档添加到正确的反向索引集合里面，返回的结果是为集合添加了多少个单词
func (c *Client) IndexDocument(docid, content string) int64 {
	words := Tokenize(content)

	pipeline := c.Conn.TxPipeline()
	for _, word := range words {
		pipeline.SAdd(context.Background(), "idx:"+word, docid)
	}
	res, err := pipeline.Exec(context.Background())
	if err != nil {
		log.Println("pipeline err in IndexDocument: ", err)
		return 0
	}
	return int64(len(res))
}
func (c *Client) setCommon(method string, names *[]string, ttl int) string {
	id := uuid.NewV4().String()
	pipeline := c.Conn.TxPipeline()

	namelist := make([]reflect.Value, 0, len(*names)+1)
	namelist = append(namelist, reflect.ValueOf("idx:"+id))
	for _, name := range *names {
		namelist = append(namelist, reflect.ValueOf("idx:"+name))
	}
	fmt.Println(namelist)
	//methodValue := reflect.ValueOf(pipeline).MethodByName(method)
	//methodValue.Call(namelist)
	pipeline.Expire(context.Background(), "idx:"+id, time.Duration(ttl)*time.Second)
	if _, err := pipeline.Exec(context.Background()); err != nil {
		log.Println("pipeline err in setCommon: ", err)
		return ""
	}
	return id
}

func (c *Client) Intersect(items []string, ttl int) string {
	return c.setCommon("SInterStore", &items, ttl)
}

func (c *Client) Union(items []string, ttl int) string {
	return c.setCommon("SUnionStore", &items, ttl)
}

func (c *Client) Difference(items []string, ttl int) string {
	return c.setCommon("SDiffStore", &items, ttl)
}

func Parse(query string) (all [][]string, unwantedlist []string) {
	unwanted, current := utils.Set{}, utils.Set{}
	for _, word := range common.QUERYRE.FindAllString(strings.ToLower(query), -1) {
		prefix := word[0]
		if prefix == '+' || prefix == '-' {
			word = word[1:]
		} else {
			prefix = 0
		}

		switch {
		case len(word) < 2 || common.STOPWORDS[sort.SearchStrings(common.STOPWORDS, word)] == word:
			continue
		case prefix == '-':
			unwanted.Add(word)
			continue
		case len(current) != 0 && prefix == 0:
			all = append(all, current.Getkeys())
			current = utils.Set{}
		}

		current.Add(word)
	}

	if len(current) != 0 {
		all = append(all, current.Getkeys())
	}

	unwantedlist = append(unwantedlist, unwanted.Getkeys()...)
	return
}

// 返回一个ID作为搜索结果，该ID表示对应的集合里包含了与用户给定的搜索参数的相匹配的文档
func (c *Client) ParseAndSearch(query string, ttl int) string {
	all, unwanted := Parse(query)

	if len(all) == 0 {
		return ""
	}

	toIntersect := []string{}
	for _, syn := range all {
		if len(syn) > 1 {
			toIntersect = append(toIntersect, c.Union(syn, ttl))
		} else {
			toIntersect = append(toIntersect, syn[0])
		}
	}

	var intersectResult string
	if len(toIntersect) > 1 {
		intersectResult = c.Intersect(toIntersect, ttl)
	} else {
		intersectResult = toIntersect[0]
	}

	if len(unwanted) != 0 {
		unwanted = append([]string{intersectResult}, unwanted...)
		return c.Difference(unwanted, ttl)
	}
	//fmt.Println(reflect.TypeOf(intersectResult))
	return intersectResult
}

func (c *Client) SearchAndSort(query string, id string, ttl int, sort string, start, num int64) ([]string, string) {
	var order string
	if strings.HasPrefix(sort, "-") {
		order = "DESC"
	} else {
		order = "ASC"
	}

	sort = strings.TrimPrefix(sort, "-")
	by := sort

	alpha := !strings.Contains(sort, "updated") && !strings.Contains(sort, "id") &&
		!strings.Contains(sort, "created")

	if id != "" && !c.Conn.Expire(context.Background(), id, time.Duration(ttl)*time.Second).Val() {
		id = ""
	}

	if id == "" {
		id = c.ParseAndSearch(query, ttl)
		fmt.Println(id)
	}

	var res *redis.StringSliceCmd
	pipeline := c.Conn.TxPipeline()
	pipeline.SCard(context.Background(), "idx:"+id)
	res = pipeline.Sort(context.Background(), "idx:"+id, &redis.Sort{By: by, Alpha: alpha, Order: order, Offset: start, Count: num})
	if _, err := pipeline.Exec(context.Background()); err != nil {
		log.Println("pipeline err in SearchAndSort: ", err)
		return nil, ""
	}
	fmt.Println("result:", res.Val())
	return res.Val(), id
}

func main() {
	//fmt.Println(Tokenize(common.CONTENT))
	ctx := NewClient(redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	}))
	//var content1 = "The Curry is the greatest basketball player"
	//var content2 = "you should write more and more program to improve you program design"
	//var content3 = "today id so hot"
	//ctx.IndexDocument("文档A", content1)
	//ctx.IndexDocument("文档B", content2)
	//ctx.IndexDocument("文档C", content3)

	fmt.Println(ctx.SearchAndSort("basketball", "", 300, "-updated", 0, 10))

}
