# JSON
a json lib
## Warning
in order to google proto,  int64, uint64 and other type will be transfer 
to string in protoMarshal,  so if you want to give a int64 to javascript
make sure that the number less than 2^52 -1
 
for more information:  
https://developers.google.com/protocol-buffers/docs/proto3#json  
  
https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Number/MAX_SAFE_INTEGER

#### desc
sometimes we should use the same proto.go file to return http response,
and the json data will avoid omitempty key, but if we want to both run
grpc or other rpc service and http service to return answer through 
proto struct,  we can use this lib api  **MarshalAvoidOmit(v interface{})**

#### example
```go
package main

import (
	"fmt"
	myjson "github.com/chi-chu/json"
)

type person struct {
	Name string `json:"name,omitempty"`
	Age  int `json:"age,omitempty"`
}

func main() {
	a := myjson.ConfigCompatibleWithStandardLibrary
	b, err := a.Marshal(map[string]string{"abc":"cs"})
	fmt.Println(string(b), err)
	b, err = a.MarshalAvoidOmit(&person{
		Name: "",
		Age:  0,
	})
	fmt.Println(string(b), err)
	b, err = a.Marshal(&person{
		Name: "",
		Age:  0,
	})
	fmt.Println(string(b), err)
}
```

and this will output
```
{"abc":"cs"} <nil>
{"name":"","age":0} <nil>
{} <nil>
```


#### In Chinese

开启http和rpc服务，通过同一套 proto 结构定义返回， 对于http的服务， json编码会被proto.go
文件的omitempty忽略， 通过 **MarshalAvoidOmit(v interface{})** 这个方法，将不会忽略
输出完整的json