func compareVersion(version1 string, version2 string) int {
  v1 := strings.Split(version1, ".")
  v2 := strings.Split(version2, ".")

  for (len(v1) > 0) && (len(v2) > 0) {
    v1Int, _ := strconv.Atoi(v1[0])
    v2Int, _ := strconv.Atoi(v2[0])

    if v1Int == v2Int {
      v1 = v1[1:]
      v2 = v2[1:]
      continue
    }

    if v1Int > v2Int{
      return 1
    }

    if v1Int < v2Int {
      return -1
    }

    v1 = v1[1:]
    v2 = v2[1:]
  }

  v1String := strings.Join(v1, "")
  v2String := strings.Join(v2, "")

  v1Int, _ := strconv.Atoi(v1String)
  v2Int, _ := strconv.Atoi(v2String)

  if v1Int == v2Int {
    return 0
  } else if v1Int > v2Int {
    return 1
  } else {
    return -1
  }
}