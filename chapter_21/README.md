## 20-1 context包的使用
- WithCancel   创建一个带有Clear()关闭函数的ctx
- WithDeadline 创建一个带有超时时间点的 ctx
- WithTimeout  创建一个带有超时时间的 ctx
- WithValue    创建一个携带了参数的 ctx
- TODO         在不确定使用上下文时使用
- Background   在确定使用上下文时使用
- Context      基础结构 实现方法
    - Deadline()
    - Done()
    - Err()
    - Value()
    
