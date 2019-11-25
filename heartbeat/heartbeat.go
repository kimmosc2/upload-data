package heartbeat

import (
    "os"
    "time"
    "upload/util/rabbitmq"
)

// StartHeartBeat()每隔5秒向apiServers exchange发送心跳服务
func StartHeartBeat() {
    q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
    defer q.Close()

    for {
        q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"), )
        time.Sleep(5 * time.Second)
    }
}
