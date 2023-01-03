func validIPAddress(queryIP string) string {
    ip := strings.Split(queryIP,".")
    if len(ip) != 4{
        ip = strings.Split(queryIP,":")
    }
    if len(ip) == 4{
        for _,v := range(ip){
            n,ok := strconv.Atoi(v)
            if ok != nil{
                return "Neither"
            }else if n <= 255 && n >= 0 && (v[0] != '0' || (n == 0 && len(v) == 1)){
                continue
            }else{
                return "Neither"
            }
        }
        return "IPv4"
    }else if len(ip) == 8{
        for _,v := range(ip){
            if len(v) > 4 || len(v) == 0{
                return "Neither"
            }
            for _,x := range(v){
                if (x >= '0' && x <= '9') || (x >= 'a' && x <= 'f') || (x >= 'A' && x <= 'F'){
                    continue
                }else{
                    return "Neither"
                }
            }
        }
        return "IPv6"
    }
    return "Neither"
}