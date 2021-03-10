/*
 * @lc app=leetcode.cn id=224 lang=golang
 *
 * [224] 基本计算器
 */

// @lc code=start
func calculate(s string) int {
	res, _ := calculateFunction(s)
	return res
}
func calculateFunction(s string) (int, int) {
	i := 0
	sum := 0
	// fmt.Printf("=== %s\n", s)
	tokenType, s2, i := getNextToken(s, i)
	// fmt.Printf("type=%d content=%s i=%d\n", tokenType, s2, i)
	if tokenType != NUMBER && tokenType != BRACKETS {
		return 0, 0
	}

	if s2 == "(" {
		a, iAdd := calculateFunction(s[i:])
		sum += a
		i += iAdd
	} else if s2 == ")" {
		return sum, i
	}

	a, _ := strconv.Atoi(s2)
	sum += a
	for tokenType != END {
		tokenType, s2, i = getNextToken(s, i)
		// fmt.Printf("type=%d content=%s i=%d\n",tokenType,s2,i)
		switch tokenType {
		case OPERATION:
			operation := s2
			// 继续获取下一个操作数
			tokenType, s2, i = getNextToken(s, i)
			if tokenType != NUMBER && tokenType != BRACKETS {
				print("ERROR")
				return 0, 0
			}
			if tokenType == BRACKETS {
				// fmt.Printf("type=%d content=%s i=%d\n",tokenType,s2,i)
				a, iAdd := calculateFunction(s[i:])
				sum += a
				i += iAdd
			}

			switch operation {
			case "+":
				a, _ := strconv.Atoi(s2)
				sum += a
			case "-":
				a, _ := strconv.Atoi(s2)
				sum -= a
			}
		case BRACKETS:
			if s2 == "(" {
				a, iAdd := calculateFunction(s[i:])
				sum += a
				i += iAdd
			} else if s2 == ")" {
				return sum, i
			}
		case END:
			return sum, i
		default:
			// fmt.Printf("type=%d content=%s i=%d  ERROR!!!\n",tokenType,s2,i)
			return 0, 0
		}
		// fmt.Printf("sum=%d\n",sum)
	}
	return sum, i
}

type TokenType int

const (
	NUMBER TokenType = iota
	OPERATION
	BRACKETS
	END
)

func getNextToken(s string, i int) (TokenType, string, int) {
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			start := i
			i++
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
			end := i
			numberStr := s[start:end]
			return NUMBER, numberStr, i
		} else if s[i] == '(' || s[i] == ')' {
			str := fmt.Sprintf("%c", s[i])
			i++
			return BRACKETS, str, i
		} else if s[i] == '+' || s[i] == '-' {
			str := fmt.Sprintf("%c", s[i])
			i++
			return OPERATION, str, i
		}
	}
	return END, "END", i

}

// @lc code=end

