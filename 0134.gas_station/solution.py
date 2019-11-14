class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        station = 0
        total_gas = current_gas = 0
        for i in range(len(gas)):
            total_gas += gas[i] - cost[i]
            current_gas += gas[i] - cost[i]
            if current_gas < 0:
                station = i + 1
                current_gas = 0
        return station if total_gas >= 0 else -1
