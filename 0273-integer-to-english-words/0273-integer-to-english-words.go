func numberToWords(num int) string {
    r := ""
    if num == 0{
        return "Zero"
    }
    for cnt := 0;num != 0;cnt++{
        cur := num%1000
        num = num/1000      
        if cnt%3 == 1 {
            if cur != 0{
                r = "Thousand "+r
            }          
        }else if cnt%3 == 2 {
            if cur != 0{
                r = "Million "+r
            }           
        }else if cnt > 0 &&  cur != 0{
            r = "Billion "+r
        }
        r = decode_3digit(cur)+r
    }
    return r[:len(r)-1]
}

func decode_3digit(n int) string{
    res := ""
    hun := n/100
    if hun != 0{
        if hun == 1{
            res += "One Hundred "
        }else  if hun == 2{
            res += "Two Hundred "
        }else  if hun == 3{
            res += "Three Hundred "
        }else  if hun == 4{
            res += "Four Hundred "
        }else  if hun == 5{
            res += "Five Hundred "
        }else  if hun == 6{
            res += "Six Hundred "
        }else  if hun == 7{
            res += "Seven Hundred "
        }else  if hun == 8{
            res += "Eight Hundred "
        }else  if hun == 9{
            res += "Nine Hundred "
        }
    }
    tens := n%100
    if tens/10 == 1{
        if tens == 10{
        res += "Ten "
      }else  if tens == 11{
        res += "Eleven "
      }else if tens == 12{
        res += "Twelve "
      }else if tens == 13{
        res += "Thirteen "
      }else if tens == 14{
        res += "Fourteen "
      }else if tens == 15{
        res += "Fifteen "
      }else if tens == 16{
        res += "Sixteen "
      }else if tens == 17{
        res += "Seventeen "
      }else if tens == 18{
        res += "Eighteen "
      }else if tens == 19{
        res += "Nineteen "
      }        
      return res  
    }
    if tens < 10{
        goto kk
    }
    if tens/10 == 2{
        res += "Twenty "
        goto kk
    }
    if tens/10 == 3{
        res += "Thirty "
        goto kk
    }
    if tens/10 == 4{
        res += "Forty "
        goto kk
    }
    if tens/10 == 5{
        res += "Fifty "
        goto kk
    }
    if tens/10 == 6{
        res += "Sixty "
        goto kk
    }
    if tens/10 == 7{
        res += "Seventy "
        goto kk
    }
    if tens/10 == 8{
        res += "Eighty "
        goto kk
    }
    if tens/10 == 9{
        res += "Ninety "
        goto kk
    }
    kk:sin := tens %10    
    //fmt.Println(sin)
    if sin == 0{
        return res
        }
    if sin == 1{
            res += "One "
        return res
        }
    if sin == 2{
            res += "Two "
        return res
        }
    if sin == 3{
            res += "Three "
        return res
        }
    if sin == 4{
            res += "Four "
        return res
        }
    if sin == 5{
            res += "Five "
        return res
        }
    if sin == 6{
            res += "Six "
        return res
        }
    if sin == 7{
            res += "Seven "
        return res
        }
    if sin == 8{
            res += "Eight "
        return res
        }
    if sin == 9{
            res += "Nine "
        return res
        } 
    return res
    
}