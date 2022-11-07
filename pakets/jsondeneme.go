package pakets

import (
	"encoding/json"
	"fmt"
)

func JsonMap() {
	futbolcuJson := `{
"formano" :10,
"adi":"Hagi",
"takım":"Galatasaray",
"ulke":"Türkiye",
"deger":1500000
}`
	//fmt.Println(futbolcuJson)

	var futbolcu map[string]interface{}

	json.Unmarshal([]byte(futbolcuJson), &futbolcu)
	fmt.Println("Forma No: ", futbolcu["formano"])
	fmt.Println("adi : ", futbolcu["adi"])
	fmt.Println("takım : ", futbolcu["takım"])
	fmt.Printf("deger : %.2f", futbolcu["deger"])

}
