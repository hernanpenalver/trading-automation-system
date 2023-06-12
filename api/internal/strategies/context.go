package strategies

import "fmt"

const (
	parameterTypeInteger = "integer"
	parameterTypeString  = "string"
)

type Context struct {
	Name       string
	Parameters []*Parameter
}

func NewContext(name string, parameters []*Parameter) *Context {
	return &Context{Name: name, Parameters: parameters}
}

type Parameter struct {
	Type  string
	Name  string
	Value interface{}
	//Data interface{}
}

//type ParameterInteger struct {
//	Min int
//	Max int
//}
//
//type ParameterString struct {
//	Values []string
//}

func (p *Parameter) GetIntValue() int {
	return p.GetValue().(int)
}

func (p *Parameter) GetStringValue() string {
	return p.GetValue().(string)
}

func (p *Parameter) GetValue() interface{} {
	//switch p.Type {
	//case parameterTypeInteger:
	//	result := p.Data.(ParameterInteger)
	//	return result.Min
	//
	//case parameterTypeString:
	//	result := p.Data.(ParameterString)
	//	return result.Values[0]
	//
	//default:
	//	panic("Parameter type invalid")
	//}
	return p.Value
}

func (s *Context) GetParameter(name string) *Parameter {
	for _, parameter := range s.Parameters {
		if parameter.Name == name {
			return parameter
		}
	}
	panic("Parameter not found")
}

func (s *Context) SetParameter(name string, value interface{}) {
	for _, parameter := range s.Parameters {
		if parameter.Name == name {
			parameter.Value = value
			return
		}
	}
	panic("Parameter not found")
}

func (s *Context) ToString() string {
	var result string
	for _, parameter := range s.Parameters {
		result += fmt.Sprintf("%s_%v_", parameter.Name, parameter.Value)
	}

	return result
}

//func (s *Context) StringifyParams() string {
//	var stringifyParams string
//	for _, parameter := range s.Parameters {
//		stringifyParams += fmt.Sprintf("%s_%d", parameter.Name, parameter.Value)
//	}
//	return stringifyParams
//}
