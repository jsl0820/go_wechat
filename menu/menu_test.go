package menu

import (
	"encoding/json"
	"testing"

	. "github.com/jsl0820/wechat"
)

func init() {
	config := Config{
		WxAppId:     "wx582ef3694f7a7546",
		WxAppSecret: "148ee9063222674ef03e4c21776e02cd",
	}

	WxConfig(config)
}

func TestCurrent(t *testing.T) {
	if info, err := Current(); err != nil {
		t.Log(info)
	} else {
		t.Log(err)
	}
}

func TestCreate(t *testing.T) {

	mune := Item{
		"button": []Item{
			Item{
				"type": "click",
				"name": "今日歌曲",
				"key":  "V1001_TODAY_MUSIC",
			},
			Item{
				"name": "测试菜单",
				"sub_button": []Item{
					Item{
						"type": "view",
						"name": "搜索",
						"url":  "http://www.soso.com/",
					},
					Item{
						"type": "click",
						"name": "赞一下我们",
						"key":  "V1001_GOOD",
					},
				},
			},
		},
	}

	b, err := json.Marshal(mune)
	t.Log(string(b))
	t.Log(err)

	isCreated := Create(mune)
	t.Log(isCreated)
}

func TestDel(t *testing.T) {
	isDeled := Del()
	t.Log("是否删除", isDeled)
}

func TestCreateCondMenu(t *testing.T) {

	mune := Item{
		"button": []Item{
			Item{
				"type": "click",
				"name": "今日歌曲",
				"key":  "V1001_TODAY_MUSIC",
			},
			Item{
				"name": "测试菜单",
				"sub_button": []Item{
					Item{
						"type": "view",
						"name": "搜索",
						"url":  "http://www.soso.com/",
					},
					Item{
						"type": "click",
						"name": "赞一下我们",
						"key":  "V1001_GOOD",
					},
				},
			},
		},
		"matchrule": Item{
			"country": "中国",
		},
	}

	menuid, created := CreateCondMenu(mune)
	t.Log(menuid)
	t.Log(created)
}
