package engine

type ParseFunc func(body []byte, url string) ParseResult

type Parser interface {
	Parse(content []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}
type Proxy struct {
	Ip   string
	Port string
}

type Request struct {
	Url   string
	Parse Parser
}

type ParseResult struct {
	Requests    Request
	Proxys      []Proxy
	NextRequest Request
}

// 默认的解析方法
func NilParseFunc(body []byte, url string) ParseResult {
	return ParseResult{}
}

type NilParse struct {
}

func (NilParse) Serialize() (name string, args interface{}) {
	return "NilParse", nil
}

func (NilParse) Parse(content []byte, url string) ParseResult {
	return ParseResult{}
}

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (p *FuncParser) Parse(content []byte, url string) ParseResult {
	return p.parser(content, url)
}

func (p *FuncParser) Serialize() (name string, args interface{}) {
	return p.name, nil
}

func NewFuncParse(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}
