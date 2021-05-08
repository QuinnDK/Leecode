package main

type ParkingSystem1 struct {
	big int
	medium int
	small int
}
func Constructor1(big int, medium int, small int) ParkingSystem1 {
	return ParkingSystem1{big:big,small:small,medium:medium}
}

func (this *ParkingSystem1) Addcar(carType int )bool  {
	switch carType {
	case 1:
		if this.big>0{
			this.big--
			return true
		}else {
			return false
		}
	case 2:
		if this.medium>0{
			this.medium--
			return true
		}else {
			return false
		}
	case 3:
		if this.small>0{
			this.small--
			return true
		}else {
			return false
		}
	default:
		return false

	}

}