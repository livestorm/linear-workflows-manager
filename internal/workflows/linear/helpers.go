package linear

import "strings"

func (t *Ticket) GetLabel(name string) string {
	for _, label := range t.Data.Labels {
		if strings.Contains(label.Name, name+": ") {
			kvList := strings.Split(label.Name, ": ")
			return kvList[1]
		}
	}
	return ""
}
