package mysql

import (
	"fmt"
	"strings"

	"github.com/HarrisonEagle/HawkORM/utils"
)

type Condition struct {
	parser   *utils.Parser
	andConds []string
	orConds  []string
	notConds []string
}

func newCondition(parser *utils.Parser) *Condition {
	return &Condition{parser: parser}
}

func (c *Condition) SetAND(conditions interface{}) {
	fields := c.parser.ExtractAllColumnsFromStructOrSlice(conditions, true)
	values := c.parser.ExtractAllValuesFromStruct(conditions, true)
	for i := 0; i < len(fields); i++ {
		condPair := fmt.Sprintf("%s = \"%s\"", fields[i], values[i])
		c.andConds = append(c.andConds, condPair)
	}
}

func (c *Condition) SetOR(conditions interface{}) {
	fields := c.parser.ExtractAllColumnsFromStructOrSlice(conditions, true)
	values := c.parser.ExtractAllValuesFromStruct(conditions, true)
	for i := 0; i < len(fields); i++ {
		var condPair string
		condPair = fmt.Sprintf("%s = \"%s\"", fields[i], values[i])
		c.orConds = append(c.orConds, condPair)
	}
}

func (c *Condition) SetNOT(conditions interface{}) {
	fields := c.parser.ExtractAllColumnsFromStructOrSlice(conditions, true)
	values := c.parser.ExtractAllValuesFromStruct(conditions, true)
	for i := 0; i < len(fields); i++ {
		var condPair string
		condPair = fmt.Sprintf("%s <> \"%s\"", fields[i], values[i])
		c.notConds = append(c.notConds, condPair)
	}
}

func (c *Condition) getConditionQuery() string {
	orCondsStr := ""
	allConds := c.andConds
	if len(c.orConds) > 0 {
		orCondsStr = "(" + strings.Join(c.orConds, " OR ") + ")"

		allConds = append(allConds, orCondsStr)
	}
	allConds = append(allConds, c.notConds...)
	if len(allConds) == 0 {
		return ""
	}
	return "WHERE " + strings.Join(allConds, " AND ")
}
