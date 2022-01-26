package problems

func DutchFlagSort(balls []string) []string {
	i, j := 0, 0
	k := len(balls) - 1
	for j <= k {
		switch balls[j] {
		case "R":
			balls[i], balls[j] = balls[j], balls[i]
			i++
			j++
		case "B":
			balls[j], balls[k] = balls[k], balls[j]
			k--
		default:
			j++
		}
	}
	return balls
}
