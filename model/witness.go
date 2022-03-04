package model

type (
	/*
		所有的HTTP请求都需要定义Https请求的报文体和响应的报文体
		写在结构体属性后面的叫做json tag，这是因为Go语言中的public变量首字母大写
		而一般的Json的结构体的属性名称都是小写的。通过设置json tag在序列化的时候
		可以自动转换成需要的字段名称。
	*/
	WitnessTransactionsRequest struct {
		Id string `json:"id"`
	}

	WitnessTransactionResponse struct {
		TransactionList []string `json:"transctionList"`
	}
)
