package ChainOfResponsibility

import "fmt"

type BoardingProcessor interface {
	SetNextProcessor(processor BoardingProcessor)
	ProcessorFunc(passenger *Passenger)
}

// Passenger 旅客
type Passenger struct {
	name                  string // 姓名
	hasBoardingPass       bool   // 是否办理登机牌
	hasLuggage            bool   // 是否有行李需要托运
	isPassIdentityCheck   bool   // 是否通过身份校验
	isPassSecurityCheck   bool   // 是否通过安检
	isCompleteForBoarding bool   // 是否完成登机
}

// 登机处理器基类
type baseBoardingProcessor struct {
	next BoardingProcessor
}

func (b *baseBoardingProcessor) SetNextProcessor(next BoardingProcessor) {
	b.next = next
}

func (b *baseBoardingProcessor) ProcessorFunc(passenger *Passenger) {
	if b.next != nil {
		b.next.ProcessorFunc(passenger)
	}
}

//登机牌处理器
type boardingPassProcessor struct {
	baseBoardingProcessor
}

func (b *boardingPassProcessor) ProcessorFunc(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("为客户 %s 办理登机牌手续\n", passenger.name)
		passenger.hasBoardingPass = true
	}
	b.baseBoardingProcessor.ProcessorFunc(passenger)
}

// luggageCheckInProcessor 托运行李处理器
type luggageCheckInProcessor struct {
	baseBoardingProcessor
}

func (l *luggageCheckInProcessor) ProcessorFunc(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("客户 %s 未办理登机牌手续，不能托运行李\n", passenger.name)
		return
	}
	if passenger.hasLuggage {
		fmt.Printf("为客户 %s 办理托运行李\n", passenger.name)
	}
	l.next.ProcessorFunc(passenger)
}

// identityCheckProcessor 校验身份处理器
type identityCheckProcessor struct {
	baseBoardingProcessor
}

func (i identityCheckProcessor) ProcessorFunc(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌，不能办理身份校验;\n", passenger.name)
		return
	}
	if !passenger.isPassIdentityCheck {
		fmt.Printf("旅客%s 校验身份\n", passenger.name)
		passenger.isPassIdentityCheck = true
	}
	i.next.ProcessorFunc(passenger)
}

// securityCheckProcessor 安检处理器
type securityCheckProcessor struct {
	baseBoardingProcessor
}

func (s *securityCheckProcessor) ProcessorFunc(passenger *Passenger) {
	if !passenger.hasBoardingPass {
		fmt.Printf("旅客%s未办理登机牌，不能进行安检;\n", passenger.name)
		return
	}
	if !passenger.isPassSecurityCheck {
		fmt.Printf("为旅客%s进行安检;\n", passenger.name)
		passenger.isPassSecurityCheck = true
	}
	s.baseBoardingProcessor.ProcessorFunc(passenger)
}

// completeBoardingProcessor 完成登机处理器
type completeBoardingProcessor struct {
	baseBoardingProcessor
}

func (c *completeBoardingProcessor) ProcessorFunc(passenger *Passenger) {
	if !passenger.hasBoardingPass ||
		!passenger.isPassIdentityCheck ||
		!passenger.isPassSecurityCheck {
		fmt.Printf("旅客%s登机检查过程未完成，不能登机;\n", passenger.name)
		return
	}
	passenger.isCompleteForBoarding = true
	fmt.Printf("旅客%s成功登机;\n", passenger.name)
}

// BuildBoardingProcessorChain 构建登机流程处理链
func BuildBoardingProcessorChain() BoardingProcessor {
	completeBoardingNode := &completeBoardingProcessor{}

	securityCheckNode := &securityCheckProcessor{}
	securityCheckNode.SetNextProcessor(completeBoardingNode)

	identityCheckNode := &identityCheckProcessor{}
	identityCheckNode.SetNextProcessor(securityCheckNode)

	luggageCheckInNode := &luggageCheckInProcessor{}
	luggageCheckInNode.SetNextProcessor(identityCheckNode)

	boardingPassNode := &boardingPassProcessor{}
	boardingPassNode.SetNextProcessor(luggageCheckInNode)
	return boardingPassNode
}
