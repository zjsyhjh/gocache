package view

/*
 * 对byte数组以及string的包装
 * 如果bytes != nil, 则使用bytes，否则使用str
 */
type ByteView struct {
	bytes []byte
	str   string
}
