# Homework for JiKeShiJian Golang Training Camp
## Week 2
sql.ErrNoRows 属于 Sentinel error，当业务层关系这个error发生的原因时，
应该要包装起来往上层抛出，由业务代码可以根据包装后包含的上下文信息来决定怎么处理这个error。
可以使用 pkg/errors 包的wrap 或是 fmt.Errorf 来包装error。
