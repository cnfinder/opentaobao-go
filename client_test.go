package opentaobao

import (
	"fmt"
	"testing"
)

func init() {
	AppKey = "29807915"
	AppSecret = ""
	Router = "http://gw.api.taobao.com/router/rest"
}
func TestExecute(t *testing.T) {
	// http://open.taobao.com/docs/api.htm?apiId=24515
	/*res, err := Execute("taobao.tbk.item.get", Parameter{
		"fields": "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick",
		"q":      "女装",
		"cat":    "16,18",
	})*/

	res, err := Execute("taobao.tbk.shop.get",Parameter{
		"fields":"user_id,shop_title,shop_type,seller_nick,pict_url,shop_url",
		"q":"女装",
		"is_tmall":true,
		"start_credit":1,

	})

	if err != nil {
		t.Fatal(err)
	}



/*	fmt.Println("商品数量:", res.Get("tbk_item_get_response").Get("total_results").MustInt())
	var imtes []interface{}
	imtes, _ = res.Get("tbk_item_get_response").Get("results").Get("n_tbk_item").Array()
	for _, v := range imtes {
		fmt.Println("======")
		item := v.(map[string]interface{})
		fmt.Println("商品名称:", item["title"])
		fmt.Println("商品价格:", item["reserve_price"])
		fmt.Println("商品链接:", item["item_url"])
	}*/

	fmt.Printf("%+v",string(res))
}


