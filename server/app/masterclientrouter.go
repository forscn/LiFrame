
package app

import (
	"encoding/json"
	"github.com/llr104/LiFrame/core/liFace"
	"github.com/llr104/LiFrame/core/liNet"
	"github.com/llr104/LiFrame/proto"
	"github.com/llr104/LiFrame/utils"
	"os"
	"time"
)

var MClientRouter MasterClientRouter

type MasterClientRouter struct {
	liNet.BaseRouter
	isShutDown bool
}

func (s *MasterClientRouter) NameSpace() string {
	return "System"
}

func (s *MasterClientRouter) ServerListAck(rsp liFace.IRespond) {

	ackInfo := proto.ServerListAck{}
	err := json.Unmarshal(rsp.GetData(), &ackInfo)
	utils.Log.Info("ServerListAck: %v", ackInfo)
	if err != nil{
		utils.Log.Info("ServerListAck error:%s",err.Error())
	}else{
		ServerMgr.Update(ackInfo.ServerMap)
	}

}

func (s *MasterClientRouter) ShutDown(req liFace.IRequest, rsp liFace.IMessage) {
	utils.Log.Info("ShutDown:%s", req.GetMessage().GetMsgName())

	if s.isShutDown == false {
		//是否需要做一些退出操作
		s.isShutDown = true
		f := GetShutDownFunc()
		if f != nil{
			f()
		}
		time.Sleep(5*time.Second)
		os.Exit(0)
	}

}
