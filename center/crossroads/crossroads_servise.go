package crossroads

import (
	"fmt"
	"time"
)

func StartCrossRoads(crossRoads CrossRoads) error {

	for {
		fmt.Println("start crossroads")

		err := crossRoads.CollectInformationFromTLs()
		if err != nil {
			return err
		}

		times, commands := crossRoads.FindOptimal()

		timer := time.NewTimer(time.Duration(times) * 100 * time.Millisecond)

		err = crossRoads.Execute(times, commands)
		if err != nil {
			return err
		}

		<-timer.C
	}

}
