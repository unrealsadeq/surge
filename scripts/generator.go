package scripts

import "fmt"

// portRange generates port range rules with specified protocol.
// portRange("TCP", [][2]int{{10000, 10019}})
// result:
// AND,((DEST-PORT,10000), (PROTOCOL,TCP))
// AND,((DEST-PORT,10001), (PROTOCOL,TCP))
// ...
// // AND,((DEST-PORT,10019), (PROTOCOL,TCP))
func portRange(protocol string, ranges [][2]int) []string {
	var rules []string

	for _, ports := range ranges {
		start, stop := ports[0], ports[1]
		for i := start; i < stop+1; i++ {
			rules = append(rules, fmt.Sprintf("AND,((DEST-PORT,%d), (PROTOCOL,%s))", i, protocol))
		}
	}

	return rules
}
