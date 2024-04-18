func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    var base = strs[0]

    for i := range base {
        for _, word := range strs[1:] {
            if i == len(word) || word[i] != base[i] {
                return base[0:i]
            }
        }
    }
    return base
}