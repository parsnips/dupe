package main

import (
	"fmt"

	"github.com/parsnips/dupe/util"
)

type xyzVector struct {
	x float64
	y float64
	z float64
}

const accelg float64 = 9.80665 // m/s^2

func calc3dimVector(timeDelta float64, objectVector, forceVec xyzVector) xyzVector {

	returnVector := xyzVector{x: 0.0, y: 0.0, z: 0.0}

	returnVector.x = objectVector.x + (forceVec.x * timeDelta)
	returnVector.y = objectVector.y + (forceVec.y * timeDelta)
	returnVector.z = objectVector.z + (forceVec.z * timeDelta)

	return returnVector
}

func main() {
	momentVector := xyzVector{
		x: 150.0,
		y: -90000.0,
		z: -110.0,
	} // 1000lb capsule mass(lb)*vel(m/s)

	forceInstant := xyzVector{
		x: -24.0,
		y: 5000.0,
		z: 17.6,
	} // thrust(lb)

	newMoment := xyzVector{
		x: 0.0,
		y: 0.0,
		z: 0.0,
	}

	var deltaTime float64 = 6.25 // seconds

	var checker util.Dupe
	checker = &util.RealDupeChecker{}

	newMoment = calc3dimVector(deltaTime, momentVector, forceInstant)
	fmt.Println(checker.IsDupe("a.txt", "b.txt"))
	fmt.Println("New x = ", newMoment.x)
	fmt.Println("New y = ", newMoment.y)
	fmt.Println("New z = ", newMoment.z)
}
