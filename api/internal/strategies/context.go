package strategies


const (
	parameterTypeInteger = "integer"
	parameterTypeString = "string"
)

type Context struct {
	Name       string
	Parameters []Parameter
}

func NewContext(name string, parameters []Parameter) *Context {
	return &Context{Name: name, Parameters: parameters}
}

type Parameter struct {
	Type string
	Name string
	Data interface{}
}

type ParameterInteger struct {
	Min int
	Max int
}

type ParameterString struct {
	Values []string
}

func (p *Parameter) GetIntValue() int {
	return p.GetValue().(int)
}

func (p *Parameter) GetStringValue() string {
	return p.GetValue().(string)
}

func (p *Parameter) GetValue() interface{} {
	switch p.Type {
	case parameterTypeInteger:
		result := p.Data.(ParameterInteger)
		return result.Min

	case parameterTypeString:
		result := p.Data.(ParameterString)
		return result.Values[0]

	default:
		panic("Parameter type invalid")
	}
}


func (s *Context) GetParameter(name string) *Parameter {
	for _, parameter := range s.Parameters {
		if parameter.Name == name {
			return &parameter
		}
	}
	return nil
}

//func (s *Context) StringifyParams() string {
//	var stringifyParams string
//	for _, parameter := range s.Parameters {
//		stringifyParams += fmt.Sprintf("%s_%d", parameter.Name, parameter.Value)
//	}
//	return stringifyParams
//}