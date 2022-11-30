package ChainOfResponsibility

import "testing"

func TestBoarding(t *testing.T) {
	chain := BuildBoardingProcessorChain()
	passenger := &Passenger{
		name:       "李",
		hasLuggage: true,
	}
	chain.ProcessorFunc(passenger)
}
