package module

import (
	"fmt"
	"github.com/liangyaopei/checker"
	"reflect"
	"strings"
)

type fieldRule struct {
	fieldExpr string
	rule      main.Rule

	ruleName string
}

func (r fieldRule) Check(param interface{}) (bool, string) {
	exprValue, kind := fetchField(param, r.fieldExpr)
	if kind == reflect.Invalid {
		return false,
			fmt.Sprintf("[%s]:'%s' cannot be found", r.ruleName, r.fieldExpr)
	}
	if exprValue == nil {
		return false,
			fmt.Sprintf("[%s]:'%s' is nil", r.ruleName, r.fieldExpr)
	}
	return r.rule.Check(exprValue)
}

// Field applies rule to fieldExpr
func Field(fieldExpr string, rule main.Rule) main.Rule {
	return fieldRule{
		fieldExpr: fieldExpr,
		rule:      rule,
		ruleName:  "fieldRule",
	}
}

type andRule struct {
	rules []main.Rule
}

func (r andRule) Check(param interface{}) (bool, string) {
	for i := 0; i < len(r.rules); i++ {
		isValid, msg := r.rules[i].Check(param)
		if !isValid {
			return isValid, msg
		}
	}
	return true, ""
}

// And accepts slice of rules
// is passed when all rules passed
func And(rules ...main.Rule) main.Rule {
	return andRule{
		rules: rules,
	}
}

type orRule struct {
	rules []main.Rule
}

func (r orRule) Check(param interface{}) (bool, string) {
	messages := make([]string, 0, len(r.rules))
	for i := 0; i < len(r.rules); i++ {
		isValid, msg := r.rules[i].Check(param)
		if isValid {
			return true, ""
		}
		messages = append(messages, msg)
	}
	return false,
		fmt.Sprintf("%s, at least one ot them should be true",
			strings.Join(messages, " or "))
}

// Or accepts slice of rules
// is failed when all rules failed
func Or(rules ...main.Rule) main.Rule {
	return orRule{
		rules: rules,
	}
}

type notRule struct {
	innerRule main.Rule
}

func (r notRule) Check(param interface{}) (bool, string) {
	isInnerValid, errMsg := r.innerRule.Check(param)
	isValid := !isInnerValid
	if !isValid {
		return false,
			fmt.Sprintf("[notRule]:{%s}", errMsg)
	}
	return true, ""
}

// Not returns the opposite if innerRule
func Not(innerRule main.Rule) main.Rule {
	return notRule{
		innerRule: innerRule,
	}
}
