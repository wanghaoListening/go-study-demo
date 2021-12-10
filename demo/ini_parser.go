package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     string `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     string `ini:"port"`
	Database string `ini:"database"`
	Password string `ini:"password"`
}

func loadIni(fileName string, data interface{}) error {

	//
	t := reflect.TypeOf(data)
	fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		return fmt.Errorf("the data must be ptr type") //格式化输出之后返回一个error类型
	}
	//获取值的类型
	if t.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("the data must be struct type")
	}

	//读文件得到的字节类型数据
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		return fmt.Errorf("read file failed %v", err)
	}

	fileText := string(file)

	lines := strings.Split(fileText, "\n")
	fmt.Println(lines)

	return nil
}
