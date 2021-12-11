package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
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

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
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

	var structName string
	for idx, line := range lines {

		if len(line) == 0 || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			//注释空行忽略
			continue
		}
		if strings.HasPrefix(line, "[") {

			selectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(selectionName) == 0 {
				return fmt.Errorf("line:%d syntax error", idx+1)
			}
			//根据selectionName去data当中根据反射找到结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if selectionName == field.Tag.Get("ini") {
					//说明找到了对应的嵌套结构体
					structName = field.Name
				}
			}
		} else {
			//以等号分割某一行
			v := reflect.ValueOf(data)
			structObj := v.Elem().FieldByName(structName)
			if structObj.Kind() != reflect.Struct {

				return fmt.Errorf("%s 不是一个结构体类型", structName)
			}
			segs := strings.Split(line, "=")
			//遍历结构体当中的每一个字段
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < structObj.NumField(); i++ {
				field := structObj.Type().Field(i)
				fieldType = field
				if field.Tag.Get("ini") == segs[0] {
					fieldName = field.Name
					break
				}
			}
			fieldObj := structObj.FieldByName(fieldName)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(segs[1])
			case reflect.Int:
				valInt, err := strconv.ParseInt(segs[1], 10, 64)
				if err != nil {
					return fmt.Errorf("line:%d value type error", idx+1)
				}
				fieldObj.SetInt(valInt)
			case reflect.Bool:
				valBool, err := strconv.ParseBool(segs[1])
				if err != nil {
					return fmt.Errorf("line:%d value type error", idx+1)
				}
				fieldObj.SetBool(valBool)
			}

		}

	}

	return nil
}

/*func main() {

	var cfg Config
	err := loadIni("/Users/bytedance/dev/code/demo_code/demo/conf.ini", &cfg)
	if err != nil {

		fmt.Printf("load ini failed err: %v\n", err)
	}

	fmt.Printf("%#v", cfg)
}*/
