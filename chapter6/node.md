##### 1.GO语言的并发同步模型来源于通信顺序进程 (CSP),CSP是一种消息传递模型，通过在goroutine之间传递数据来传递消息，不加锁
##### 2.进程可以看作是包含应用程序在运行中需要用到和维护的各种资源的容器
##### 3.
![图片](https://raw.githubusercontent.com/joeeeeee/goinaction/master/chapter6/images/%E5%BA%94%E7%94%A8%E7%A8%8B%E5%BA%8F%E8%BF%9B%E7%A8%8B%E5%92%8C%E7%BA%BF%E7%A8%8B%E7%AE%80%E8%A6%81%E6%8F%8F%E7%BB%98.png)