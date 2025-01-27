package handlers


import "dopc/internal/services"

type HandlerTest struct {
	Service services.CalculateService
}

func DeliveryHandlerTest(service services.CalculateService) *HandlerTest {
	return &HandlerTest{Service: service}
}

