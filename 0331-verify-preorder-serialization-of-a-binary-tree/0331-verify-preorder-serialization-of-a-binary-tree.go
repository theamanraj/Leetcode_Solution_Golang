var flag = true

func isValid(nodes *[]string) {
  if len(*nodes) == 0 {
    flag = false
    return
  }
  val := (*nodes)[0]
  *nodes = (*nodes)[1:]
  
  if val == "#" {
    return
  }
  
  
  isValid(nodes)
  isValid(nodes)
  
}

func isValidSerialization(preorder string) bool {  
  flag = true
  nodes := strings.Split(preorder,",")
  isValid(&nodes)
  return flag && len(nodes) < 1
}