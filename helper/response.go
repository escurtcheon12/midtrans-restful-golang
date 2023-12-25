package helper

// func ToAuthResponse(user domain.User) web.AuthResponse {
// 	return web.AuthResponse{
// 		Id:         user.Id,
// 		Username:   user.Username,
// 		Email:      user.Email,
// 		Phone:      user.Phone,
// 		Password:   user.Password,
// 		Created_at: user.Created_at,
// 		Updated_at: user.Updated_at,
// 	}
// }

// func ToAuthResponses(user []domain.User) []web.AuthResponse {
// 	var authResponses []web.AuthResponse
// 	for _, user := range user {
// 		authResponses = append(authResponses, ToAuthResponse(user))
// 	}
// 	return authResponses
// }

// type ToResponseP struct {
// 	UserP          domain.User
// 	OrdersP        domain.Orders
// 	NotificationsP domain.Notifications
// }

// type ToResponseR struct {
// 	UserR          web.UserResponse
// 	OrdersR        web.OrdersResponse
// 	NotificationsR web.NotificationsResponse
// 	AuthR          web.AuthResponse
// }

// func ToResponse(p ToResponseP) *ToResponseR {
// 	return &ToResponseR{
// 		UserR:          p.UserP,
// 		OrdersR:        p.OrdersP,
// 		NotificationsR: p.NotificationsP,
// 	}
// }

// func ToResponses(user []domain.User) []web.AuthResponse {
// 	var authResponses []web.AuthResponse
// 	for _, user := range user {
// 		authResponses = append(authResponses, ToAuthResponse(user))
// 	}
// 	return authResponses
// }
