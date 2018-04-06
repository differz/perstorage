package usecases

import "../contracts/usecases"

type PlaceOrderUseCase struct {
	projectId   int64
	subject     string
	description string
}

func NewPlaceOrderUseCase() PlaceOrderUseCase {
	return PlaceOrderUseCase{
		projectId: 1,
	}
}

func (u *PlaceOrderUseCase) placeOrder(request contracts.PlaceOrderRequest, output contracts.PlaceOrderOutput) {
	/*      int projectId = request.getProjectId();
	        String subject = request.getSubject();
	        String description = request.getDescription();

	        orderRepository.saveIssue(issue);

			output.onResponse(orderID);
	*/
}
