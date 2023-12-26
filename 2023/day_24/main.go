package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	items := parseInput("inputs.txt")

	res1 := solvePart1(items, 200000000000000, 400000000000000)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2(items)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Item struct {
	x, y, z    float64
	dx, dy, dz float64
}

func solvePart1(items []Item, min float64, max float64) int {
	numIntersect := 0
	n := len(items)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			intersection := findFuturePathCrossing2D(items[i], items[j])
			if intersection == nil {
				continue
			}

			if intersection[0] < min || intersection[0] > max || intersection[1] < min || intersection[1] > max {
				continue
			}

			numIntersect++
		}
	}

	return numIntersect
}

func solvePart2(items []Item) int64 {
	//group points with similar velocities together
	xVelocities := map[float64][]float64{}
	yVelocities := map[float64][]float64{}
	zVelocities := map[float64][]float64{}
	for _, item := range items {
		if curr, ok := xVelocities[item.dx]; ok {
			curr = append(curr, item.x)
			xVelocities[item.dx] = curr
		} else {
			xVelocities[item.dx] = []float64{item.x}
		}

		if curr, ok := yVelocities[item.dy]; ok {
			curr = append(curr, item.y)
			yVelocities[item.dy] = curr
		} else {
			yVelocities[item.dy] = []float64{item.y}
		}

		if curr, ok := zVelocities[item.dz]; ok {
			curr = append(curr, item.z)
			zVelocities[item.dz] = curr
		} else {
			zVelocities[item.dz] = []float64{item.z}
		}
	}

	//calculate possible rock velocities for each axis
	rockDxs := calculatePossibleVelocities(xVelocities)
	rockDys := calculatePossibleVelocities(yVelocities)
	rockDzs := calculatePossibleVelocities(zVelocities)

	//test for valid speed combinations on each axis
	n := len(items)
	for _, rdx := range rockDxs {
		for _, rdy := range rockDys {
			for _, rdz := range rockDzs {

				for i := 0; i < n-1; i++ {
					for j := i + 1; j < n; j++ {
						a, b := items[i], items[j]

						//line equations from rock to target points
						ma, ba := lineEquation2D(a.x, a.y, a.dx-rdx, a.dy-rdy)
						mb, bb := lineEquation2D(b.x, b.y, b.dx-rdx, b.dy-rdy)

						rock := Item{
							dx: rdx,
							dy: rdy,
							dz: rdz,
						}

						//rock.y = ma * rock.x + ba = mb * rock.x + bb
						rock.x = (bb - ba) / (ma - mb)
						rock.y = ma*rock.x + ba

						//rock.x = a.x + time * (a.dx - rdx)
						time := (rock.x - a.x) / (a.dx - rdx)
						rock.z = a.z + time*(a.dz-rdz)

						//test whether rock does collide with all points
						valid := true
						for _, test := range items {
							if _, ok := hasCollission(rock, test); !ok {
								valid = false
								break
							}
						}

						if valid {
							//there can be 2 valid rocks
							//each having negative and positive totals
							//valid answer is the positive one
							total := int64(rock.x + rock.y + rock.z)
							if total > 0 {
								return total
							}
						}
					}
				}
			}
		}
	}

	return -1
}

func parseInput(path string) []Item {
	strToNumbers := func(str string) []float64 {
		re := regexp.MustCompile(`-?\d+`)

		numbers := []float64{}
		for _, val := range re.FindAllString(str, -1) {
			num, _ := strconv.ParseFloat(val, 64)
			numbers = append(numbers, num)
		}

		return numbers
	}

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := []Item{}
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " @ ")
		position := strToNumbers(data[0])
		velocity := strToNumbers(data[1])

		item := Item{
			x:  position[0],
			y:  position[1],
			z:  position[2],
			dx: velocity[0],
			dy: velocity[1],
			dz: velocity[2],
		}
		result = append(result, item)
	}

	return result
}

/*
Line Equations :
y2 - y1 = m * (x2-x1) || m = dy / dx
y = mx + b || b = y - mx
*/
func lineEquation2D(x, y, dx, dy float64) (m float64, b float64) {
	if dx == 0 {
		return math.MaxInt32, math.MaxInt32
	}

	m = dy / dx
	b = y - (m * x)
	return m, b
}

func findFuturePathCrossing2D(a, b Item) []float64 {
	m1, b1 := lineEquation2D(a.x, a.y, a.dx, a.dy)
	m2, b2 := lineEquation2D(b.x, b.y, b.dx, b.dy)

	if m1 == m2 { //parallel lines
		return nil
	}

	posX := (b2 - b1) / (m1 - m2)
	posY := (m1 * posX) + b1

	var time1, time2 float64
	if m1 == 0 {
		time1 = (posX - a.x) / a.dx
	} else {
		time1 = (posY - a.y) / a.dy
	}

	if m2 == 0 {
		time2 = (posX - b.x) / b.dx
	} else {
		time2 = (posY - b.y) / b.dy
	}

	if time1 < 0 || time2 < 0 {
		return nil
	}

	return []float64{posX, posY}
}

func hasCollission(a, b Item) (float64, bool) {
	//on collission : x2 = a.x + time * a.dx = b.x + time * b.dx
	const tolerance = 1e-9
	var time float64
	if math.Abs(a.dx-b.dx) > tolerance {
		time = (b.x - a.x) / (a.dx - b.dx)
	} else if math.Abs(a.dy-b.dy) > tolerance {
		time = (b.y - a.y) / (a.dy - b.dy)
	} else {
		time = (b.z - a.z) / (a.dz - b.dz)
	}

	//verify collission on all axis
	ax2 := a.x + (time * a.dx)
	bx2 := b.x + (time * b.dx)
	ay2 := a.y + (time * a.dy)
	by2 := b.y + (time * b.dy)
	az2 := a.z + (time * a.dz)
	bz2 := b.z + (time * b.dz)

	if math.Abs(ax2-bx2) > tolerance || math.Abs(ay2-by2) > tolerance || math.Abs(az2-bz2) > tolerance {
		return 0, false
	}

	return time, true
}

/*
Premise :
Given 2 points that have same velocities on an axis,
we can guarantee that the distance between those 2 points
is always constant on said axis.

Subsequently, we can infer that a valid rock speed which can
collide with both points on that axis will satisfy the equation :
distance % (pointVelocity - rockVelocity) == 0 .

This assumes non float time and non float velocities
to travel between these 2 points.
*/
func calculatePossibleVelocities(velocities map[float64][]float64) []float64 {
	rockDvs := make([]float64, 2001)
	for rdv := float64(-1000); rdv <= float64(1000); rdv++ {
		rockDvs[1000+int(rdv)] = rdv
	}

	for vel, pos := range velocities {
		n := len(pos)
		if n < 2 {
			continue
		}

		for i := 0; i < n-1; i++ {
			for j := i + 1; j < n; j++ {
				a, b := pos[i], pos[j]
				newDvs := []float64{}
				for _, dv := range rockDvs {
					if dv == vel || int64(a-b)%int64(vel-dv) != 0 {
						continue
					}

					newDvs = append(newDvs, dv)
				}

				rockDvs = newDvs
			}
		}
	}

	//fmt.Println("rockDvs", rockDvs)
	return rockDvs
}
