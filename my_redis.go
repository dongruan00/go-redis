package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//连接redis
	conn, err := redis.Dial("tcp", "localhost:6379")

	if err!=nil{
		fmt.Println("conn redis failed", err)
		return
	}

	fmt.Println("conn redis success")

	defer conn.Close()
	//redis set操作
	_,err=conn.Do("Set","abc", 100)

	if err!=nil{
		fmt.Println(err)
		return
	}
	//redis get操作
	r, err:=redis.Int(conn.Do("Get", "abc"))

	if err!=nil{
		fmt.Println("get abc failed,",err)
		return
	}
	fmt.Println(r)  //100


	//redis 批量set
	_, err=conn.Do("MSet", "abc", 100, "efg", 300)

	if err!=nil{
		fmt.Println(err)
		return
	}
	//redis 批量get
	rs, err:=redis.Ints(conn.Do("MGet", "abc", "efg"))
	if err!=nil{
		fmt.Println("get abc failed", err)
		return
	}
	//获取批量的值
	//100
	//300
	for _,v :=range rs{
		fmt.Println(v)
	}


	//redis设置过期时间
	_,err=conn.Do("expire", "abc", 10)

	if err!=nil{
		fmt.Println(err)
		return
	}

	//redis list操作lpush
	_,err=conn.Do("lpush", "book_list", "a", "b", 100)

	if err!=nil{
		fmt.Println(err)
		return
	}
	//redis list操作lpop
	list, err:=redis.String(conn.Do("lpop", "book_list"))

	if err!=nil{
		fmt.Println("get abc failed", err)
		return
	}

	fmt.Println(list)

	//redis操作hash设置
	_, err=conn.Do("HSet", "books", "abc", 100)

	if err!=nil{
		fmt.Println(err)
		return
	}
	//redis获取hash的值
	books, err:=redis.Int(conn.Do("HGet", "books", "abc"))

	if err!=nil{
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(books)

	
}
