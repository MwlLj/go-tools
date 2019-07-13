package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MwlLj/gotools/timetool"
	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
	"net"
	"os"
	"strings"
	"time"
)

var (
	FormatJson string = "json"
)

type CSub struct {
}

func (this *CSub) Run(param *CRunParam) {
	for {
		for {
			conn, err := net.Dial("tcp", *this.joinHost(param))
			if err != nil {
				fmt.Fprintln(os.Stderr, "dial: ", err)
				break
			}
			cc := mqtt.NewClientConn(conn)
			cc.Dump = param.Dump
			cc.ClientId = param.Id

			tq := make([]proto.TopicQos, 1)
			tq[0].Topic = param.Topic
			tq[0].Qos = proto.QosAtMostOnce

			if err := cc.Connect(param.UserName, param.UserPwd); err != nil {
				fmt.Fprintf(os.Stderr, "connect: %v\n", err)
				break
			}
			cc.Subscribe(tq)

			for m := range cc.Incoming {
				fmt.Printf("%s, [%s]\n", m.TopicName, timetool.GetNowSecondFormat())
				payload := m.Payload.(proto.BytesPayload)
				if len(payload) == 0 {
					fmt.Println("null")
				} else {
					this.format(&(param.Format), payload)
				}
				fmt.Println()
			}
			fmt.Println("read break")
			break
		}
		time.Sleep(time.Second * time.Duration(1))
	}
}

func (this *CSub) format(f *string, payload []byte) {
	if *f == FormatJson {
		var out bytes.Buffer
		json.Indent(&out, payload, "", "\t")
		fmt.Println(out.String())
	} else {
		fmt.Println(string(payload))
	}
}

func (this *CSub) joinHost(param *CRunParam) *string {
	host := strings.Join([]string{param.Host, param.Port}, ":")
	return &host
}

func NewSub() *CSub {
	return &CSub{}
}
