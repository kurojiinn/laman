import Foundation

struct Order: Codable, Identifiable {
    let id: UUID
    let guestName: String?
    let guestPhone: String?
    let guestAddress: String?
    let comment: String?
    let status: String?
    let itemsTotal: Double?
    let serviceFee: Double?
    let deliveryFee: Double?
    let finalTotal: Double?
    let createdAt: Date?
    let items: [OrderItem]?
}

struct CreateOrderRequest: Codable {
    let guestName: String
    let guestPhone: String
    let guestAddress: String
    let deliveryAddress: String
    let comment: String?
    let paymentMethod: PaymentMethod
    let items: [CreateOrderItem]
}

struct CreateOrderItem: Codable {
    let productId: UUID
    let quantity: Int
}

enum PaymentMethod: String, Codable {
    case cash = "CASH"
    case transfer = "TRANSFER"
}
