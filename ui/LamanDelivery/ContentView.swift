import SwiftUI

final class AppState: ObservableObject {
    @Published var categories: [Category] = []
    @Published var products: [Product] = []
    @Published var cart: [UUID: Int] = [:]
    @Published var orders: [Order] = []
    @Published var isLoading: Bool = false
    @Published var errorMessage: String? = nil

    private let api = LamanAPI()

    @MainActor
    func loadCatalog() async {
        isLoading = true
        errorMessage = nil
        do {
            async let categoriesTask = api.getCategories()
            async let productsTask = api.getProducts()
            categories = try await categoriesTask
            products = try await productsTask
        } catch {
            errorMessage = error.localizedDescription
        }
        isLoading = false
    }

    func quantity(for productID: UUID) -> Int {
        cart[productID] ?? 0
    }

    func setQuantity(_ quantity: Int, for productID: UUID) {
        if quantity <= 0 {
            cart.removeValue(forKey: productID)
        } else {
            cart[productID] = quantity
        }
    }

    var cartItems: [CartItem] {
        products.compactMap { product in
            let qty = cart[product.id] ?? 0
            guard qty > 0 else { return nil }
            return CartItem(product: product, quantity: qty)
        }
    }

    var subtotal: Double {
        cartItems.reduce(0) { $0 + (Double($1.quantity) * $1.product.price) }
    }

    var deliveryFee: Double {
        200
    }

    var serviceFee: Double {
        max(0, subtotal * 0.05)
    }

    var total: Double {
        subtotal + deliveryFee + serviceFee
    }

    var totalItems: Int {
        cartItems.reduce(0) { $0 + $1.quantity }
    }

    @MainActor
    func submitOrder(request: CreateOrderRequest) async throws -> Order {
        let order = try await api.createOrder(request: request)
        orders.insert(order, at: 0)
        cart.removeAll()
        return order
    }
}

struct ContentView: View {
    @EnvironmentObject private var appState: AppState

    var body: some View {
        TabView {
            NavigationStack {
                CatalogView()
            }
            .tabItem {
                Label("Каталог", systemImage: "building.2")
            }

            NavigationStack {
                CartView()
            }
            .tabItem {
                Label("Корзина", systemImage: "cart")
            }

            NavigationStack {
                OrdersView()
            }
            .tabItem {
                Label("Заказы", systemImage: "list.bullet.rectangle")
            }
        }
        .task {
            await appState.loadCatalog()
        }
    }
}

struct CartItem: Identifiable {
    let id = UUID()
    let product: Product
    let quantity: Int
}

#Preview {
    ContentView()
        .environmentObject(AppState())
}
