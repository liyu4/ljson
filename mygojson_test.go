package mygojson

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	json := `{"name":"light","weigth":"maybe55kg","result":["light","fish","dylan"]}`
	jsonArray := `["1","2",3]`
	j1 := Json(json)
	fmt.Println(j1)   //stdout &{map[name:light weigth:maybe55kg result:[light fish dylan]]}
	Json(json).Type() // stdou map[string]initerface{}

	//-----------------------------------------------
	j2 := Json(json).Get("name")
	fmt.Println(j2) //stdout &{light}
	//-----------------------------------------------
	j3 := Json(json).Get("result").GetIndex(2)
	fmt.Println(j3) //stdout &{fish}
	//-----------------------------------------------
	j4 := Json(jsonArray).ArrayIndex(3)
	Json(jsonArray).Type()
	fmt.Println(j4) //stdout 3
	//-----------------------------------------------
	json1 := `{"from":"en","to":"zh","trans_result":[{"src": 1.1433,"dst":"\u4eca\u5929"},{"src":"tomorrow","dst":"\u660e\u5929"}]}`
	j5 := Json(json1).Get("trans_result").GetKey("dst", 2)
	fmt.Println(j5)
	//-----------------------------------------------
	// 递归map
	json2 := `{"trans_result":{"src": 3.141596,"dst":"\u4eca\u5929"}}`
	j6 := Json(json2).GetPath("trans_result", "src").ToString()
	fmt.Println(j6)
	//-----------------------------------------------
	j7, j8 := Json(json2).Get("trans_result").ToArray()
	fmt.Println(j7, j8)
	//-----------------------------------------------

	j9 := Json(jsonArray).StringtoArray()
	fmt.Println(j9)
	Json(jsonArray).Type()

	//-----------------------------------------------
	jsonArrayMap := `[{"name":"light"},{"sex":"male"}]`
	fmt.Println(Json(jsonArrayMap))
	Json(jsonArrayMap).Type()
	j10, j11 := Json(jsonArrayMap).ToArray()
	fmt.Println(j10, j11)
}
