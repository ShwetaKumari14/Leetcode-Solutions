package main

import (
	"fmt"
	"strings"
)

func main() {

	emails := []string{"fg.r.u.uzj+o.pw@kziczvh.com", "r.cyo.g+d.h+b.ja@tgsg.z.com", "fg.r.u.uzj+o.f.d@kziczvh.com", "r.cyo.g+ng.r.iq@tgsg.z.com", "fg.r.u.uzj+lp.k@kziczvh.com", "r.cyo.g+n.h.e+n.g@tgsg.z.com", "fg.r.u.uzj+k+p.j@kziczvh.com", "fg.r.u.uzj+w.y+b@kziczvh.com", "r.cyo.g+x+d.c+f.t@tgsg.z.com", "r.cyo.g+x+t.y.l.i@tgsg.z.com", "r.cyo.g+brxxi@tgsg.z.com", "r.cyo.g+z+dr.k.u@tgsg.z.com", "r.cyo.g+d+l.c.n+g@tgsg.z.com", "fg.r.u.uzj+vq.o@kziczvh.com", "fg.r.u.uzj+uzq@kziczvh.com", "fg.r.u.uzj+mvz@kziczvh.com", "fg.r.u.uzj+taj@kziczvh.com", "fg.r.u.uzj+fek@kziczvh.com"}
	result := numUniqueEmails(emails)
	fmt.Println(result)
}

func numUniqueEmails(emails []string) int {

	ansMap := make(map[string]int)

	for _, email := range emails {
		count, localName := checkValidEmail(email)
		if count > 0 {
			ansMap[localName]++
		}
	}

	return len(ansMap)

}

func checkValidEmail(str string) (int, string) {

	email, domainName, localName := strings.Split(str, "@"), "", ""
	if len(email) > 1 {
		localName = email[0]
		domainName = email[1]
	} else {
		return 0, ""
	}

	if strings.Contains(localName, ".") {
		localName = strings.ReplaceAll(localName, ".", "")
	}
	if strings.Contains(localName, "+") {
		if idx := strings.Index(localName, "+"); idx != -1 {
			localName = localName[:idx]
		}
	}

	return 1, localName + "@" + domainName
}
