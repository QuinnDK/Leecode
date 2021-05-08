package main

import "fmt"

type ParkingSystem struct {
	//left [4]int
	big int
	medium int
	small int
}

func Constructor(big int, medium int, small int)ParkingSystem {
	//return ParkingSystem{[4]int{0,big,medium,small}}
	return ParkingSystem{big:big,small:small,medium:medium}
}

func (this *ParkingSystem)AddCar(carType int)bool {
	//if this.left[carType]>0{
	//	this.left[carType]--
	//	return true
	//}
	//return false
	switch carType {
	case 1:
		if this.big>0{
			this.big--
			return true
		}else {
			return false
		}
	case 2:
		if this.small>0 {
			this.small--
			return true
		}else {
			return false
		}
	case 3:
		if this.medium>0{
			this.medium--
			return true
		}else {
			return false
		}
	default:
		return false
	}
}

func main() {
	parking:=Constructor(2,1,0)
	bo:=parking.AddCar(3)
	fmt.Printf("%v",bo)

}