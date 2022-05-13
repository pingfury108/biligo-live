package main

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/iyear/biligo-live"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	app := createCli()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createCli() *cli.App {
	app := &cli.App{
		Name:   "biligo-live",
		Usage:  "biligo-live",
		Action: run,
		Flags: []cli.Flag{
			&cli.Int64Flag{
				Name:  "room",
				Value: 0,
				Usage: "room ID",
			},
			&cli.Int64Flag{
				Name:  "uid",
				Value: 123456,
				Usage: "user id",
			},
			&cli.StringFlag{
				Name:  "user-key",
				Value: "",
				Usage: "user mark",
			},
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "debug mode",
			},
		},
	}
	return app
}

func run(c *cli.Context) error {
	room := c.Int64("room")
	user_key := c.String("user-key")
	uid := c.Int64("uid")

	// 获取一个Live实例
	// debug: debug模式，输出一些额外的信息
	// heartbeat: 心跳包发送间隔。不发送心跳包，70 秒之后会断开连接，通常每 30 秒发送 1 次
	// cache: Rev channel 的缓存
	// recover: panic recover后的操作函数
	l := live.NewLive(c.Bool("debug"), 30*time.Second, 0, func(err error) {
		log.Println("panic:", err)
		// do something...
	})

	// 连接ws服务器
	// dialer: ws dialer
	// host: bilibili live ws host
	if err := l.Conn(websocket.DefaultDialer, live.WsDefaultHost); err != nil {
		log.Fatal(err)
		return err
	}

	ctx, stop := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	ifError := make(chan error)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// 进入房间
		// room: room id(真实ID，短号需自行转换)
		// key: 用户标识，可留空
		// uid: 用户UID，可随机生成
		if err := l.Enter(ctx, room, user_key, uid); err != nil {
			log.Println("Error Encountered: ", err)
			log.Println("Room Disconnected")
			ifError <- err
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		rev(ctx, l)
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	select {
	case <-sc:
		fmt.Println("I want to stop")
		// 关闭ws连接与相关协程
		stop()
		break
	case err := <-ifError:
		fmt.Println("I don't want to stop, but I encountered an error: ", err)
		break
	}

	wg.Wait()
	return nil
}

func rev(ctx context.Context, l *live.Live) {
	for {
		select {
		case tp := <-l.Rev:
			if tp.Error != nil {
				// do something...
				log.Println(tp.Error)
				continue
			}
			handle(tp.Msg)
		case <-ctx.Done():
			log.Println("rev func stopped")
			return
		}
	}
}
func handle(msg live.Msg) {
	// 使用 msg.(type) 进行事件跳转和处理，常见事件基本都完成了解析(Parse)功能，不常见的功能有一些实在太难抓取
	// 更多注释和说明等待添加
	switch msg.(type) {
	// 心跳回应直播间人气值
	case *live.MsgHeartbeatReply:
		fmt.Printf("HOT: %d\n", msg.(*live.MsgHeartbeatReply).GetHot())
	// 弹幕消息
	case *live.MsgDanmaku:
		dm, err := msg.(*live.MsgDanmaku).Parse()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("弹幕[%v]: (%s[%s]) %s\n", time.UnixMilli(dm.Time).Format("2006-01-02T15:04:05"), dm.Uname, dm.MedalName, dm.Content)
	// 礼物消息
	case *live.MsgSendGift:
		g, err := msg.(*live.MsgSendGift).Parse()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%s[%v]: %s %d个%s\n", g.Action, time.Unix(g.Timestamp, 0).Format("2006-01-02T15:04:05"), g.Uname, g.Num, g.GiftName)
	// 直播间粉丝数变化消息
	case *live.MsgFansUpdate:
		f, err := msg.(*live.MsgFansUpdate).Parse()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("Room: %d,fans: %d,fansClub: %d\n", f.RoomID, f.Fans, f.FansClub)
	// case:......

	// General 表示live未实现的CMD命令，请自行处理raw数据。也可以提issue更新这个CMD
	case *live.MsgGeneral:
		//fmt.Println("unknown msg type|raw:", string(msg.Raw()))
		//fmt.Println()
	}
}
