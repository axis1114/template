package chat

import (
	"github.com/DanPlayer/randomname"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
	"template/global"
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
