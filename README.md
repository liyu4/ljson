#profile
ljson is json parse library use golang,it is easy to use and extension.

###how to use 
```
package main

import (
    "fmt"
    "ljson"
)

func main() {
    //imitate data, it can be download from other website.
    data := `{"name":"light","weigth":"maybe65kg","result":["light","fish","dylan"]}`
    json := ljson.NewJson(data)

    // mq is data structure map[string]interface{} so if you want to get value assert as need.
    mp := json.GetMapData()

    fmt.Println(mp)
}

```
