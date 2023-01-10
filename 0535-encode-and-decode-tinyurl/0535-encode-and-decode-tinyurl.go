type Codec struct {
    longToShort map[string]string
    shortToLong map[string]string
    baseUrl string
    usableStr string
}

func Constructor() Codec {
    var c Codec
    c.baseUrl = "http://tinyurl.com/"
    c.usableStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
    c.longToShort = make(map[string]string)
    c.shortToLong = make(map[string]string)
    return c
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    _, isIn := this.longToShort[longUrl]
    
    if isIn{
        return this.baseUrl + this.longToShort[longUrl]
    }
    
    var encoded string
    usableStrRune := []rune(this.usableStr)
    
    for {
        encoded = "";
        
        for i:= 0; i < 6; i++{
            encoded += string(usableStrRune[rand.Intn(1000) % 62])
        }
        
        _, isIn := this.shortToLong[encoded]

        if !isIn {
            break;
        }
    }
    
    this.longToShort[longUrl] = encoded
    this.shortToLong[encoded] = longUrl
    
    return this.baseUrl + encoded
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    shortUrlRune := []rune(shortUrl)
    return this.shortToLong[string(shortUrlRune[len(shortUrlRune) - 6:])]
}

