package ljson

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	json := `{"name":"light","weigth":"maybe55kg","result":["light","fish","dylan"]}`
	jsonArray := `["1","2",3]`
	j1 := NewJson(json)
	fmt.Println(j1)      //stdout &{map[name:light weigth:maybe55kg result:[light fish dylan]]}
	NewJson(json).Type() // stdou map[string]initerface{}

	//-----------------------------------------------
	j2 := NewJson(json).Get("name")
	fmt.Println(j2) //stdout &{light}
	//-----------------------------------------------
	j3 := NewJson(json).Get("result").GetIndex(2)
	fmt.Println(j3) //stdout &{fish}
	//-----------------------------------------------
	j4, _ := NewJson(jsonArray).ArrayIndex(3)
	NewJson(jsonArray).Type()
	fmt.Println(j4) //stdout 3
	//-----------------------------------------------
	json1 := `{"from":"en","to":"zh","trans_result":[{"src": 1.1433,"dst":"\u4eca\u5929"},{"src":"tomorrow","dst":"\u660e\u5929"}]}`
	j5 := NewJson(json1).Get("trans_result").GetKey("dst", 2)
	fmt.Println(j5)
	//-----------------------------------------------
	// 递归map
	json2 := `{"trans_result":{"src": 3.141596,"dst":"\u4eca\u5929"}}`
	j6 := NewJson(json2).GetPath("trans_result", "src").String()
	fmt.Println(j6)
	//-----------------------------------------------
	j7, j8 := NewJson(json2).Get("trans_result").ToArray()
	fmt.Println(j7, j8)
	//-----------------------------------------------

	j9, _ := NewJson(jsonArray).Array()
	fmt.Println(j9)
	NewJson(jsonArray).Type()

	//-----------------------------------------------
	jsonArrayMap := `[{"name":"light"},{"sex":"male"}]`
	fmt.Println(NewJson(jsonArrayMap))
	NewJson(jsonArrayMap).Type()
	j10, j11 := NewJson(jsonArrayMap).ToArray()
	fmt.Println(j10, j11)
}
