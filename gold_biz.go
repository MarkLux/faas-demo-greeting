package gold

import (
	"log"
)

func (s *GoldService) OnInit() {
}

func (s *GoldService) OnHandle(req *goldrpc.GoldRequest, rsp *goldrpc.GoldResponse) error {
	userName := req.Data["name"].(string)
	log.Println("userName: " + userName)
	greeting := "hello, " + userName
	rsp.Data = make(map[string]interface{})
	rsp.Data["greeting"] = greeting
	return nil
}

func (s *GoldService) OnError(err error) bool {
	log.Println(err)
	return false
}