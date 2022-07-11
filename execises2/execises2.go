package execises2

func CanFastAttack(IsAwake bool) bool {
	if !IsAwake {
		return true
	} else {
		return false
	}
}

func CanSpy(KnightIsAwake, prisonerIsAwake, archerIsAwake bool) bool {
	if KnightIsAwake || prisonerIsAwake || archerIsAwake {
		return true
	} else {
		return false
	}
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if prisonerIsAwake && !archerIsAwake {
		return true
	} else {
		return false
	}

}

func CanFreePresoner(KnightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {

	if petDogIsPresent && !archerIsAwake {

		return true
	} else if !petDogIsPresent && prisonerIsAwake && !KnightIsAwake && !archerIsAwake {
		return true
	} else if !prisonerIsAwake {
		return false
	} else {
		return false
	}
}
