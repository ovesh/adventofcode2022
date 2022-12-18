package day15

import (
	"fmt"
	"sort"
	"strings"
)

type intervalsArray [][2]int

func merge(intervals [][2]int) [][2]int {
	intA := intervalsArray(intervals)

	intervalsSorted := [][2]int(intA)

	var output [][2]int
	currentIntervalStart := intervalsSorted[0][0]
	currentIntervalEnd := intervalsSorted[0][1]
	for j := 1; j < len(intervalsSorted); j++ {
		if currentIntervalEnd >= intervalsSorted[j][0] {
			if intervalsSorted[j][1] > currentIntervalEnd {
				currentIntervalEnd = intervalsSorted[j][1]
			}
		} else {
			output = append(output, [2]int{currentIntervalStart, currentIntervalEnd})
			currentIntervalStart = intervalsSorted[j][0]
			currentIntervalEnd = intervalsSorted[j][1]
		}
	}
	output = append(output, [2]int{currentIntervalStart, currentIntervalEnd})
	if len(output) == 1 {
		return output
	}

	// edge case where diff between segments is 1 (e.g {[1,4], [5,9])
	realOutput := [][2]int{}
	for i := 1; i < len(output); i++ {
		if output[i-1][1]+1 != output[i][0] {
			realOutput = append(realOutput, output[i-1])
			continue
		}
		nextPair := [2]int{output[i-1][0], output[i][1]}
		i++
		for ; i < len(output); i++ {
			if output[i-1][1]+1 != output[i][0] {
				if i == len(output)-1 {
					realOutput = append(realOutput, output[i-1])
				}
				break
			}
			nextPair[1] = output[i][1]
		}
		realOutput = append(realOutput, nextPair)
	}
	if realOutput[len(realOutput)-1][1] != output[len(output)-1][1] { // add last one if needed
		realOutput = append(realOutput, output[len(output)-1])
	}
	return realOutput

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Part2() int {
	const maxXY = 4000000
	//const maxXY = 20
	lines := strings.Split(input0, "\n")
	allBeacons := make([]beacon, 0, len(lines))
	for _, line := range lines {
		rawSensor := line[len("Sensor at "):strings.Index(line, ":")]
		sensorX := atoi(rawSensor[2:strings.Index(rawSensor, ",")])
		sensorY := atoi(rawSensor[strings.Index(rawSensor, ", ")+4:])
		//fmt.Println("Sensor at x=", sensorX, "y=", sensorY)
		rawBeacon := line[strings.Index(line, "closest beacon is at ")+len("closest beacon is at "):]
		beaconX := atoi(rawBeacon[2:strings.Index(rawBeacon, ",")])
		beaconY := atoi(rawBeacon[strings.Index(rawBeacon, ", ")+4:])

		allBeacons = append(allBeacons, beacon{beaconX, beaconY, sensorX, sensorY})
	}

	for targetLineIdx := 0; targetLineIdx < maxXY; targetLineIdx++ {
		intervals := [][2]int{}
		for _, beacon := range allBeacons {
			distance := abs(beacon.x-beacon.sensorX) + abs(beacon.y-beacon.sensorY)
			oneSide := abs(distance - abs(targetLineIdx-beacon.sensorY))
			//fmt.Println("Beacon at x=", beacon.x, "y=", beacon.y, "sensorX=", beacon.sensorX, "sensorY=", beacon.sensorY, "distance=", distance, "oneSide=", oneSide)
			if abs(targetLineIdx-beacon.sensorY) >= distance {
				//fmt.Println("too far")
				continue
			}
			intervals = append(intervals, [2]int{max(beacon.sensorX-oneSide, 0), min(beacon.sensorX+oneSide, maxXY)})
		}
		sort.Slice(intervals, func(i, j int) bool {
			if intervals[i][0] == intervals[j][0] {
				return intervals[i][1] < intervals[j][1]
			}
			return intervals[i][0] < intervals[j][0]
		})
		intervals = merge(intervals)
		if len(intervals) < 1 && len(intervals) > 2 {
			panic("invalid intervals")
		}
		if len(intervals) == 2 {
			fmt.Println("y=", targetLineIdx, "intervals=", intervals)
			return ((intervals[0][1] + 1) * 4000000) + targetLineIdx
		}
	}

	return 0
}
