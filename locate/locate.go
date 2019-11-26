package locate

import (
    "os"
    "strconv"
    "upload/util/rabbitmq"
)

// 定位对象,验证文件是否存在本地
func Locate(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}

// 监听定位消息
func StartLocate() {
    q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
    defer q.Close()

    q.Bind("dataServers")
    c := q.Consume()
    for msg := range c {

        object, err := strconv.Unquote(string(msg.Body))
        if err != nil {
            panic(err)
        }
        if Locate(os.Getenv("STORAGE_ROOT") + "/objects/" + object) {
            q.Send(msg.ReplyTo,os.Getenv("LISTEN_ADDRESS"))
        }
    }
}
