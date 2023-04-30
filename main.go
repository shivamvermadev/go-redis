package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gomodule/redigo/redis"
	"github.com/pelletier/go-toml"
)

// var (
// 	COUNTRY = flag.String("ctr", "IND", "Country code, IND/RUS")
// 	ENV     = flag.String("env", "LOCAL", "LOCAL,TEST,PROD")
// )

func main() {

	// var name string
	// var age int
	// flag.StringVar(&name, "name", "dev", "Name of person")
	// flag.IntVar(&age, "age", 24, "Age of person")

	// flag.Parse()

	// fmt.Println(*COUNTRY)
	// fmt.Println(*ENV)
	// fmt.Println(name)
	// fmt.Println(age)

	// fmt.Println(os.Args[0])                             // this will print the current running file ie ./main
	// fmt.Println(filepath.Abs(os.Args[0]))               // will print the abs path of running file
	// fmt.Println(filepath.Abs(filepath.Dir(os.Args[0]))) // will print the abs path to the directory of running file
	absPath, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		panic(err)
	}

	configTree, err := toml.LoadFile(absPath + "/main.toml")
	if err != nil {
		panic(err)
	}

	redisHost := configTree.Get(fmt.Sprint("REDIS", ".", "HOST")).(string)
	redisPort := configTree.Get(fmt.Sprint("REDIS", ".", "PORT")).(string)

	redisPool := newPool(fmt.Sprint(redisHost, ":", redisPort))
	c := redisPool.Get()
	defer c.Close()

	c.Do("SET", "name", "shivammm")
}

func newPool(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:   3,
		MaxActive: 240,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return con, err
		},
	}
}
