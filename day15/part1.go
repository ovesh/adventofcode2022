package day15

import (
	"strconv"
	"strings"
)

func atoi(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

type beacon struct {
	x, y, sensorX, sensorY int
}

func Part1() int {
	const targetLineIdx = 2000000
	lines := strings.Split(input0, "\n")
	minX, maxX := 0, 0
	allBeacons := make([]beacon, 0, len(lines))
	for _, line := range lines {
		rawSensor := line[len("Sensor at "):strings.Index(line, ":")]
		sensorX := atoi(rawSensor[2:strings.Index(rawSensor, ",")])
		sensorY := atoi(rawSensor[strings.Index(rawSensor, ", ")+4:])
		//fmt.Println("Sensor at x=", sensorX, "y=", sensorY)
		rawBeacon := line[strings.Index(line, "closest beacon is at ")+len("closest beacon is at "):]
		beaconX := atoi(rawBeacon[2:strings.Index(rawBeacon, ",")])
		beaconY := atoi(rawBeacon[strings.Index(rawBeacon, ", ")+4:])
		//fmt.Println("Beacon at x=", beaconX, "y=", beaconY)

		distance := abs(beaconX-sensorX) + abs(beaconY-sensorY)
		remainingDistance := distance - abs(targetLineIdx-sensorY)
		if remainingDistance > 0 {
			curMin := sensorX - remainingDistance
			if curMin < minX {
				minX = curMin
			}
			curMax := sensorX + remainingDistance
			if curMax > maxX {
				maxX = curMax
			}
		}

		allBeacons = append(allBeacons, beacon{beaconX, beaconY, sensorX, sensorY})
	}
	//fmt.Println("minX=", minX, "maxX=", maxX)

	targetLine := make([]bool, maxX-minX+1)

	for _, beacon := range allBeacons {
		distance := abs(beacon.x-beacon.sensorX) + abs(beacon.y-beacon.sensorY)
		oneSide := abs(distance - abs(targetLineIdx-beacon.sensorY))
		//fmt.Println("Beacon at x=", beacon.x, "y=", beacon.y, "sensorX=", beacon.sensorX, "sensorY=", beacon.sensorY, "distance=", distance, "oneSide=", oneSide)
		if abs(targetLineIdx-beacon.sensorY) >= distance {
			//fmt.Println("too far")
			continue
		}
		for i := beacon.sensorX - oneSide; i <= beacon.sensorX+oneSide; i++ {
			relativeX := i - minX
			if relativeX >= 0 && relativeX < len(targetLine) {
				targetLine[relativeX] = true
			}
		}
	}
	for _, beacon := range allBeacons {
		if beacon.y == targetLineIdx {
			targetLine[beacon.x-minX] = false
		}
	}

	res := 0
	for _, b := range targetLine {
		if b {
			res++
		}
	}

	//fmt.Println("  0    5    0    5    0    5")
	//for _, b := range targetLine {
	//	if b {
	//		fmt.Print("#")
	//	} else {
	//		fmt.Print(".")
	//	}
	//}
	//fmt.Println()

	return res
}
