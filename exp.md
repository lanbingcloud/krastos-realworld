API:
1. protobuf3的规范缺失了很多编码时常用的内容,如必填检查,默认值等,不确定对象等,一定程度上可以import其他补充包来实现这些功能.
2. Kratos通过protobuf生成出来的数据类型除了变量本身外还会有Get方法, 该方法目测主要是用于定义接口用的,即大家都有的某个变量可以定义成同一个接口一起处理.所以定义api的时候变量名注意有相同意义的,最好定义成同一个名字
3. 根据不知道哪里的规范,Delete接口不能带body,只能通过url传参,所以太过复杂的值和参数要考虑api拆分

Data:
1. 根据依赖注入需要,这里New函数需要返回biz层定义的接口,如果需要在单元测试的时候使用该对象进行测试环境的初始化时,需要另外定义一个初始化方法(你返回的接口只能使用biz层定义的函数,不能任意调用data层的资源)

BIZ:
1. 需要进行单元测试的方法如果过程中有多个路径会返回失败,请定义一个返回错误码的值,就算业务中用不上,你单元测试的时候也好测试错误路径

TEST:
1. 该模板下的测试需要单独建一个package,不然会出现循环依赖. 设计数据结构和方法时需要注意开放必要的函数和变量出来,否则测试用例没法写.
2. Vscode里面可以对着方法右键在该package下生成一个test文件,可以参考里面的结构,需要注意的是,如果你的测试用例不能并行,请去掉t.run的代码块
3. 定义一个初始化方法,并且在里面定义t.Cleanup的方法,能让你的测试用例更加的简洁
```go
func initTestCluster(t *testing.T, vc *vault.Client, uc *vp.VaultUsercase, bc *conf.Bootstrap) {
	vc.Sys().DisableAuth("test_cluster")

	uc.EnableAuth(context.Background(), &test_cluster_01)
	t.Cleanup(func() {
		vc.Sys().DisableAuth("test_cluster")
		os.Remove(bc.Server.Authorization.Resource.Acl)
	})
}
```

OTHER:
1. 有些函数err是否为空只是标识该操作是否成功,但是返回的值是否有效需要另外判断.(Vault和casbin的有不少都是这样)