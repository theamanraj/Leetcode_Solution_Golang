import (
    "strings"
    "strconv"
)

func subdomainVisits(cpdomains []string) []string {
    domainsMap := make(map[string]int, len(cpdomains) * 2)
    
    for _, cpdomain := range cpdomains {
        raw        := strings.Split(cpdomain, " ")
        visits,  _ := strconv.Atoi(raw[0])
        domains    := strings.Split(raw[1], ".")
        
        for i, _:= range domains {
            domainsMap[strings.Join(domains[i:], ".")] += visits
        }
    }
    
    answer := make([]string, 0, len(domainsMap))
    
    for domain, visits := range domainsMap {
        answer = append(answer, strconv.Itoa(visits) + " " + domain)
    }
    
    return answer
}