func numUniqueEmails(emails []string) int {
    emailsMap := make(map[string]int)
    for i:=0; i < len(emails); i++ {
        var address string
        var domain string
        emailAddress := emails[i]
        for j:=0; j < len(emailAddress); j++ {
            if emailAddress[j] == '@' {
                address = emailAddress[0:j]
                domain = emailAddress[j+1:]
                break
            }
        }
        
        parsedAddress := strings.Replace(address, ".","",-1)
        for j:=0; j < len(parsedAddress); j++ {
            if parsedAddress[j] == '+' {
                parsedAddress = parsedAddress[0:j]
            }
        }
        parsedAddress = parsedAddress + "@"+domain
        if _, ok := emailsMap[parsedAddress]; !ok {
            emailsMap[parsedAddress] = 0
        } 
    }
    return len(emailsMap)
}