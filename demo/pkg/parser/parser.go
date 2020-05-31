package parser

//新的爬虫需要
type parser interface {
	Parse(parseInfo interface{}) interface{}
	//NewParser() error
}

func DoParse(p parser, parseInfo interface{}) interface{} {
	return p.Parse(parseInfo)
}
//func Parse(parseInfo interface{}) error{
//	return parse()
//}
//
//// Derived class have to implement parse
//func parse() error {
//	return nil
//}

func NewParser() parser {
	var p parser
	return p
}

