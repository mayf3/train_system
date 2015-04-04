### SYSU Train System

## 一些不确定的地方
暂定的设计方法是
#model：
包括负责和数据库交互的类、经常用到的类
#controller： 
分为三个文件，假定该系统名为table，则有
table\_data ： 负责路由并将数据传给template
table\_handle : table\_data内调用table\_handle的函数，暂定是包含复杂逻辑的函数
table\_utils ： 暂定是table\_handle和table\_data中重复使用的逻辑

问题在于utils和handle之间的区分，经常会遇到将handle中的复杂逻辑拆分之后，就完全变成多个utils了
所以暂定想法是，handle内重复出现三次以上的逻辑就抽出来放入utils

## TODO
timeline
