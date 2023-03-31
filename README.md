# URL-Test
检测URL列表能否正常打开

# 检测速度
### 测试设备配置  
**CPU** 13th Gen Intel(R) Core(TM) i5-13600K 3.50 GHz  
**内存** 32GB 3200MHz  
**系统** Windows 10 22H2

平均速度120万条URL/小时

# 工作原理
1. 一个Reader逐行读取需要验证URL的CSV表，将URL添加到一个具有1024缓冲容量的channel
2. 启动1024个Worker从channel中取出URL，对URL发起GET请求，如果响应超时（10s）或错误，则认定为无法打开，否则视为可以打开
3. 每个worker持有一个独立的http.Client，避免每次请求都创建一个Client增加内存开销
4. Worker获得验证结果后异步调用writer，将结果追加至结果CSV文件中
5. writer在初始化时创建一个文件对象用于保存结果，每次调用write都会追加写入到该文件中
6. 按 Ctrl + C 停止程序后会通知Reader不要再添加URL到channel中，并等待所有worker消费完channel中剩下的URL后将有效和无效的数量保存至kv.txt文件

# 现有问题
内存占用会随着时间增加逐步上升，实测100min跑约200万条URL后内存增长至14GB+