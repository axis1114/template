package chat

import (
	"encoding/json"
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"template/global"
	"template/utils"
	"time"
)

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

type GroupResponse struct {
	NickName    string    `json:"nick_name"`
	Avatar      string    `json:"avatar"`
	Content     string    `json:"content"`
	OnlineCount int       `json:"online_count"`
	Date        time.Time `json:"created_at"`
}

var ConnGroupMap = map[string]ChatUser{}

func (Chat) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Log.Error("websocket upgrade fail", zap.Error(err))
		return
	}
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	nickName := randomname.GenerateName()
	avatar := "upload/chat_avatar/default_avatar.jpg"

	chatUser := ChatUser{Conn: conn, NickName: nickName, Avatar: avatar}
	ConnGroupMap[addr] = chatUser
	global.Log.Infof("%s连接成功", chatUser.NickName)
}

func SendMsg(_addr string, response GroupResponse) {
	byteData, err := json.Marshal(response)
	if err != nil {
		global.Log.Error("json marshal fail", zap.Error(err))
		return
	}
	chatUser := ConnGroupMap[_addr]
	err = chatUser.Conn.WriteMessage(websocket.TextMessage, byteData)
	if err != nil {
		global.Log.Error("write message fail", zap.Error(err))
		return
	}
	ip, addr := getIPAndAddr(_addr)
	global.DB.Create(&models.ChatModel{
		NickName: response.NickName,
		Avatar:   response.Avatar,
		Content:  response.Content,
		IP:       ip,
		Addr:     addr,
		IsGroup:  false,
		MsgType:  response.MsgType,
	})
}

func getIPAndAddr(_addr string) (ip string, addr string) {
	addrList := strings.Split(_addr, ":")
	ip = addrList[0]
	addr = utils.GetAddr(ip)
	return ip, addr
}
