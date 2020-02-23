package minimumsubstringwindow

func tightenFromBegining(s string, start, end int, tSet map[byte]int, seen map[byte]int) (int, int) {

	for start < end {
		c, in := seen[s[start]]
		if !in {
			start++
		} else {
			cnt := tSet[s[start]]

			if c > cnt {
				c--
				seen[s[start]] = c
				start++
			} else {
				break
			}
		}
	}

	return start, end
}

func windowLen(mp map[byte]int) (ret int) {

	for _, val := range mp {
		ret += val
	}

	return
}

func passes(seen map[byte]int, tSet map[byte]int) bool {
	if len(seen) < len(tSet) {
		return false
	}

	for key := range seen {
		c := seen[key]
		c1 := tSet[key]

		if c < c1 {
			return false
		}
	}

	return true
}

// MinWindow returns the min substring that contains all chars in t
func MinWindow(s string, t string) string {

	start := 0
	end := start

	tSet := make(map[byte]int, len(t))

	for i := 0; i < len(t); i++ {
		count := tSet[t[i]]
		tSet[t[i]] = count + 1
	}

	seen := make(map[byte]int, len(t))

	for ; end < len(s); end++ {
		if _, in := tSet[s[end]]; in {
			c := seen[s[end]]
			c++
			seen[s[end]] = c

			if passes(seen, tSet) {
				break
			}
		}
	}

	if !passes(seen, tSet) {
		return ""
	}

	start, end = tightenFromBegining(s, start, end, tSet, seen)

	stTmp := start
	endTmp := end

	for endTmp < len(s) {
		if c, in := seen[s[stTmp]]; in {
			c--
			seen[s[stTmp]] = c
		}

		stTmp++
		endTmp++

		if endTmp < len(s) {
			if c, in := seen[s[endTmp]]; in {
				c++
				seen[s[endTmp]] = c
			}
		}

		if passes(seen, tSet) {
			stTmp, endTmp = tightenFromBegining(s, stTmp, endTmp, tSet, seen)

			if endTmp-stTmp < end-start {
				start = stTmp
				end = endTmp
			}
		}
	}

	return s[start : end+1]
}
