import SwiftUI

struct CartView: View {
    @EnvironmentObject private var appState: AppState

    @State private var showOrderForm = false

    private let priceColor = Color(red: 0.06, green: 0.73, blue: 0.51)
    private let accentBlue = Color(red: 0.23, green: 0.51, blue: 0.96)

    var body: some View {
        ZStack {
            LinearGradient(
                colors: [accentBlue.opacity(0.15), Color.orange.opacity(0.12)],
                startPoint: .topLeading,
                endPoint: .bottomTrailing
            )
            .ignoresSafeArea()

            VStack(spacing: 16) {
                if appState.cartItems.isEmpty {
                    ContentUnavailableView("Корзина пуста", systemImage: "cart")
                } else {
                    List {
                        Section("Товары") {
                            ForEach(appState.cartItems) { item in
                                HStack {
                                    VStack(alignment: .leading, spacing: 4) {
                                        Text(item.product.name)
                                            .font(.headline)
                                        Text("\(item.quantity) × \(priceText(item.product.price))")
                                            .font(.subheadline)
                                            .foregroundStyle(.secondary)
                                    }
                                    Spacer()
                                    Text(priceText(Double(item.quantity) * item.product.price))
                                        .foregroundStyle(priceColor)
                                }
                            }
                        }

                        Section {
                            HStack {
                                Text("Подытог")
                                Spacer()
                                Text(priceText(appState.subtotal))
                            }
                            HStack {
                                Text("Сервисный сбор")
                                Spacer()
                                Text(priceText(appState.serviceFee))
                            }
                            HStack {
                                Text("Доставка")
                                Spacer()
                                Text(priceText(appState.deliveryFee))
                            }
                            HStack {
                                Text("Итого")
                                    .font(.headline)
                                Spacer()
                                Text(priceText(appState.total))
                                    .font(.headline)
                                    .foregroundStyle(priceColor)
                            }
                        }
                    }
                    .listStyle(.insetGrouped)
                }

                Button {
                    showOrderForm = true
                } label: {
                    Text("Оформить заказ (\(appState.totalItems) товара)")
                        .frame(maxWidth: .infinity)
                        .padding()
                        .background(appState.totalItems == 0 ? Color.gray.opacity(0.3) : accentBlue)
                        .foregroundStyle(.white)
                        .clipShape(RoundedRectangle(cornerRadius: 12))
                }
                .disabled(appState.totalItems == 0)
                .padding(.horizontal)
            }
        }
        .navigationTitle("Корзина")
        .sheet(isPresented: $showOrderForm) {
            NavigationStack {
                OrderView()
            }
        }
    }

    private func priceText(_ price: Double) -> String {
        if price == Double(Int(price)) {
            return "\(Int(price))₽"
        }
        return String(format: "%.2f₽", price)
    }
}

#Preview {
    NavigationStack { CartView() }
        .environmentObject(AppState())
}
