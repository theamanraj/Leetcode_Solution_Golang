func shipWithinDays(weights []int, days int) ( res int)  {
    var sum , max int 
    for _, v := range weights {
        if v>max { max=v} 
        sum+=v
    }
    
    return sort.Search( sum, func(cap int)bool{ 
        if cap < max { return false} 
        loaded, d := 0, 1  
        for _, pack := range weights {
			if cap-loaded >=pack {
				loaded+=pack
            }else{
                d++
                loaded=pack
            } 
		} 
        return d<=days 
    })
}