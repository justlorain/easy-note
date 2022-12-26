server:
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/easy_note idl/user.thrift
	kitex --thrift-plugin validator -module github.com/cloudwego/biz-demo/easy_note -service demouser -use github.com/cloudwego/biz-demo/easy_note/kitex_gen ../../idl/user.thrift