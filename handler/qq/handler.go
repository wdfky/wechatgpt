package qq

import (
	"strings"
	"sync"

	"github.com/Logiase/MiraiGo-Template/bot"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	log "github.com/sirupsen/logrus"
	"github.com/wechatgpt/wechatbot/config"
	"github.com/wechatgpt/wechatbot/openai"
)

func init() {
	instance = &qqchatgpt{}
	bot.RegisterModule(instance)
}

type qqchatgpt struct {
}

var instance *qqchatgpt

func (q *qqchatgpt) MiraiGoModule() bot.ModuleInfo {
	return bot.ModuleInfo{
		ID:       "internal.qq",
		Instance: instance,
	}
}

func (q *qqchatgpt) Init() {
	// 初始化过程
	// 在此处可以进行 Module 的初始化配置
	// 如配置读取
}

func (q *qqchatgpt) PostInit() {
	// 第二次初始化
	// 再次过程中可以进行跨Module的动作
	// 如通用数据库等等
}

func (q *qqchatgpt) Serve(b *bot.Bot) {
	// 注册服务函数部分
	registerchatgpt(b)
}

func (q *qqchatgpt) Start(b *bot.Bot) {
	// 此函数会新开携程进行调用
	// ```go
	// 		go exampleModule.Start()
	// ```

	// 可以利用此部分进行后台操作
	// 如http服务器等等
}

func (q *qqchatgpt) Stop(b *bot.Bot, wg *sync.WaitGroup) {
	// 别忘了解锁
	// defer wg.Done()
	// 结束部分
	// 一般调用此函数时，程序接收到 os.Interrupt 信号
	// 即将退出
	// 在此处应该释放相应的资源或者对状态进行保存
}

func qqChatGptGroupMessageEvent(qqClient *client.QQClient, msg *message.GroupMessage) {
	sender := msg.Sender.Uin
	group := msg.GroupCode

	log.Infof("收到群 %v(%v) 内 %v(%v) 的消息: %v", group, qqClient.FindGroup(group).Name, sender, qqClient.FindGroup(group).FindMember(sender).DisplayName(), msg.ToString())

	qqKey := config.GetQQKeyword()
	requestText := msg.ToString()
	if qqKey != nil {
		if ok := strings.Contains(requestText, *qqKey); !ok {
			return
		}

		splitItems := strings.Split(requestText, *qqKey)
		if len(splitItems) < 2 {
			return
		}
		requestText = strings.TrimSpace(splitItems[1])
	}

	log.Info("问题：", requestText)
	reply, err := openai.Completions(requestText)
	if err != nil {
		log.Println(err)
		if reply != nil {
			result := *reply
			// 如果文字超过4000个字会回错，截取前4000个文字进行回复
			if len(result) > 4000 {
				qqClient.SendGroupMessage(msg.GroupCode, &message.SendingMessage{
					Elements: []message.IMessageElement{
						message.NewAt(msg.Sender.Uin),
						message.NewText(result[:4000]),
					},
				})
			}
		} else {
			qqClient.SendGroupMessage(msg.GroupCode, &message.SendingMessage{
				Elements: []message.IMessageElement{
					message.NewAt(msg.Sender.Uin),
					message.NewText("bot error: " + err.Error()),
				},
			})
		}
		log.Warning("chatgpt error: ", err)
		return
	}
	if reply == nil {
		return
	}
	// 如果在提问的时候没有包含？,AI会自动在开头补充个？看起来很奇怪
	result := *reply

	if strings.HasPrefix(result, "?") {
		result = strings.Replace(result, "?", "", -1)
	}
	if strings.HasPrefix(result, "？") {
		result = strings.Replace(result, "？", "", -1)
	}
	// 微信不支持markdown格式，所以把反引号直接去掉
	result = strings.Replace(result, "`", "", -1)

	qqClient.SendGroupMessage(msg.GroupCode, &message.SendingMessage{
		Elements: []message.IMessageElement{
			message.NewAt(msg.Sender.Uin),
			message.NewText(result),
		},
	})
}
func registerchatgpt(b *bot.Bot) {
	// 群消息撤回
	// b.GroupMessageRecalledEvent.Subscribe(func(qqClient *client.QQClient, event *client.GroupMessageRecalledEvent) {
	// 	logGroupMessageRecallEvent(event)
	// })

	b.GroupMessageEvent.Subscribe(qqChatGptGroupMessageEvent)

	// b.GroupMuteEvent.Subscribe(func(qqClient *client.QQClient, event *client.GroupMuteEvent) {
	// 	logGroupMuteEvent(event)
	// })
	// //b.OnGroupMuted(func(qqClient *client.QQClient, event *client.GroupMuteEvent) {
	// //	logGroupMuteEvent(event)
	// //})

	// b.PrivateMessageEvent.Subscribe(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
	// 	logPrivateMessage(privateMessage)
	// })
	// //b.OnPrivateMessage(func(qqClient *client.QQClient, privateMessage *message.PrivateMessage) {
	// //	logPrivateMessage(privateMessage)
	// //})

	// b.FriendMessageRecalledEvent.Subscribe(func(qqClient *client.QQClient, event *client.FriendMessageRecalledEvent) {
	// 	logFriendMessageRecallEvent(event)
	// })
	// //b.OnFriendMessageRecalled(func(qqClient *client.QQClient, event *client.FriendMessageRecalledEvent) {
	// //	logFriendMessageRecallEvent(event)
	// //})

	// b.DisconnectedEvent.Subscribe(func(qqClient *client.QQClient, event *client.ClientDisconnectedEvent) {
	// 	logDisconnect(event)
	// })
	// //b.OnDisconnected(func(qqClient *client.QQClient, event *client.ClientDisconnectedEvent) {
	// //	logDisconnect(event)
	// //})
}
