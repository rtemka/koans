// #134
// https://leetcode.com/problems/gas-station/description/
package leetcode

func canCompleteCircuit(gas []int, cost []int) int {
	// If the total gas is less than the total cost, completing the circuit is impossible.
	if sum(gas) < sum(cost) {
		return -1
	}
	start, tank := 0, 0
	for i := range gas {
		tank += gas[i] - cost[i]
		// If our tank has negative gas, we cannot continue through the
		// circuit from the current start point, nor from any station
		// before or including the current station 'i'.
		if tank < 0 {
			// Set the next station as the new start point and reset the tank.
			start = i + 1
			tank = 0
		}
	}
	return start
}

func sum(s []int) int {
	sum := 0
	for i := range s {
		sum += s[i]
	}
	return sum
}
