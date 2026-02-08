import SwiftUI

struct OrdersView: View {
    @EnvironmentObject private var appState: AppState

    private let priceColor = Color(red: 0.06, green: 0.73, blue: 0.51)

    var body: some View {
        List {
            if appState.orders.isEmpty {
                ContentUnavailableView("Заказов пока нет", systemImage: "list.bullet.rectangle")
            } else {
                ForEach(appState.orders) { order in
                    VStack(alignment: .leading, spacing: 6) {
                        Text("Заказ #\(order.id.uuidString.prefix(8))")
                            .font(.headline)
                        if let name = order.guestName {
                            Text(name)
                                .font(.subheadline)
                                .foregroundStyle(.secondary)
                        }
                        HStack {
                            Text("Итого")
                            Spacer()
                            Text(priceText(order.finalTotal ?? 0))
                                .foregroundStyle(priceColor)
                        }
                    }
                    .padding(.vertical, 6)
                }
            }
        }
        .navigationTitle("Заказы")
        .listStyle(.insetGrouped)
    }

    private func priceText(_ price: Double) -> String {
        if price == Double(Int(price)) {
            return "\(Int(price))₽"
        }
        return String(format: "%.2f₽", price)
    }
}

#Preview {
    NavigationStack { OrdersView() }
        .environmentObject(AppState())
}
