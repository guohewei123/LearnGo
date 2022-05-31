## 20-1 context包的使用
划分任务边界
- WithCancel   创建一个带有cancel()关闭函数的ctx，由外部控制什么时候结束(只有调用cancel()就会结束)
- WithDeadline 创建一个带有超时时间点的 ctx, 表示在什么时间点结束
- WithTimeout  创建一个带有超时时间的 ctx， 表示在多久后结束，同时具有 WithCancel 的特性
- WithValue    创建一个携带了参数的 ctx， 表示要携带的值
- TODO         在不确定使用上下文时使用
- Background   在确定使用上下文时使用
- Context      基础结构 实现方法
    - Deadline()
    - Done()
    - Err()
    - Value()
    
