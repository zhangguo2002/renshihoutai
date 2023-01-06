# 使用方法

###使用条件

配置好go环境

下载了bee工具和beego框架（即那两个go get命令和一个go install 命令,不会的私信我，我的微信号为wxid_e61d7opkki2822)

下载了MySQL

### 具体操作

1、先创建一个数据库名为renshihoutai

2、再利用navicat或者其他MySQL图形化软件运行那三个SQL文件（运行完之后可以删掉)

3、打开该项目文件的终端并输入

`bee run`

4、打开浏览器并输入localhost:8080/register进入到注册页面，注册账号密码，注册完了后就可以登录进去

5、这里要注意管理员一定要注册admin为账号，普通用户注册1、2、3、4、5等为账号（因为这是根yuan_gong_zi_liao那张表中的id字段所关联的，这里如果有兴趣的朋友可以优化一下，也可以通知作者该怎么改才好,互相学习吗^ - ^)

以下是users表的内容，这里的id为主键后面的name和pwd分别是账号密码

![1672992692178](D:\computer_language_study\码云\GO\renshihoutai\Readme.assets\1672992692178.png)

以下是bu_men_xin_xi表的内容

![1672992784163](D:\computer_language_study\码云\GO\renshihoutai\Readme.assets\1672992784163.png)

以下是yuan_gong_zi_liao表的内容

![1672992840346](D:\computer_language_study\码云\GO\renshihoutai\Readme.assets\1672992840346.png)