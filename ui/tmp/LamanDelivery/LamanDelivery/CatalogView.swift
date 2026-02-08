import SwiftUI

struct CatalogView: View {
    @EnvironmentObject private var appState: AppState

    private let priceColor = Color(red: 0.06, green: 0.73, blue: 0.51)
    private let accentBlue = Color(red: 0.23, green: 0.51, blue: 0.96)
    @State private var selectedCategoryId: UUID? = nil

    var body: some View {
        ZStack {
            LinearGradient(
                colors: [accentBlue.opacity(0.15), Color.orange.opacity(0.12)],
                startPoint: .topLeading,
                endPoint: .bottomTrailing
            )
            .ignoresSafeArea()

            List {
                if !appState.categories.isEmpty {
                    Section("Категории") {
                        ScrollView(.horizontal, showsIndicators: false) {
                            HStack(spacing: 8) {
                                categoryChip(title: "Все", isSelected: selectedCategoryId == nil) {
                                    selectedCategoryId = nil
                                    Task { await appState.loadProducts(categoryId: nil) }
                                }

                                ForEach(appState.categories) { category in
                                    categoryChip(title: category.name, isSelected: selectedCategoryId == category.id) {
                                        selectedCategoryId = category.id
                                        Task { await appState.loadProducts(categoryId: category.id) }
                                    }
                                }
                            }
                            .padding(.vertical, 4)
                        }
                    }
                }

                Section("Товары") {
                    ForEach(appState.products) { product in
                        HStack(alignment: .center, spacing: 12) {
                            Image(systemName: productIcon(for: product))
                                .font(.title2)
                                .foregroundStyle(accentBlue)
                                .frame(width: 28)

                            VStack(alignment: .leading, spacing: 4) {
                                Text(product.name)
                                    .font(.headline)
                                Text(productSubtitle(for: product))
                                    .font(.subheadline)
                                    .foregroundStyle(.secondary)
                            }

                            Spacer()

                            VStack(alignment: .trailing, spacing: 6) {
                                Text(priceText(product.price))
                                    .font(.headline)
                                    .foregroundStyle(priceColor)

                                Stepper(value: Binding(
                                    get: { appState.quantity(for: product.id) },
                                    set: { appState.setQuantity($0, for: product.id) }
                                ), in: 0...99) {
                                    Text("\(appState.quantity(for: product.id))")
                                        .font(.subheadline)
                                }
                                .labelsHidden()
                            }
                        }
                        .padding(.vertical, 4)
                    }
                }
            }
            .listStyle(.insetGrouped)
        }
        .navigationTitle("Каталог")
        .toolbar {
            ToolbarItem(placement: .topBarTrailing) {
                NavigationLink {
                    CartView()
                } label: {
                    Text("Оформить (\(appState.totalItems))")
                }
                .disabled(appState.totalItems == 0)
            }
        }
        .overlay {
            if appState.isLoading {
                ProgressView("Загрузка каталога...")
                    .padding()
                    .background(.ultraThinMaterial)
                    .clipShape(RoundedRectangle(cornerRadius: 12))
            }
        }
        .alert("Ошибка сети", isPresented: Binding(
            get: { appState.errorMessage != nil },
            set: { _ in appState.errorMessage = nil }
        )) {
            Button("Ок", role: .cancel) {}
        } message: {
            Text(appState.errorMessage ?? "Неизвестная ошибка")
        }
    }

    private func productSubtitle(for product: Product) -> String {
        let weight = product.weight != nil ? "• \(String(format: "%.1f", product.weight!)) кг" : ""
        return "\(product.description ?? "") \(weight)".trimmingCharacters(in: .whitespaces)
    }

    private func priceText(_ price: Double) -> String {
        if price == Double(Int(price)) {
            return "\(Int(price))₽"
        }
        return String(format: "%.2f₽", price)
    }

    private func productIcon(for product: Product) -> String {
        let name = product.name.lowercased()
        if name.contains("цемент") {
            return "bag.fill"
        }
        if name.contains("хлеб") {
            return "leaf"
        }
        if name.contains("молоко") {
            return "drop.fill"
        }
        if name.contains("чипс") {
            return "flame.fill"
        }
        return "shippingbox.fill"
    }

    private func categoryChip(title: String, isSelected: Bool, action: @escaping () -> Void) -> some View {
        Button(action: action) {
            Text(title)
                .font(.subheadline)
                .padding(.horizontal, 12)
                .padding(.vertical, 6)
                .background(isSelected ? accentBlue : Color.gray.opacity(0.2))
                .foregroundStyle(isSelected ? .white : .primary)
                .clipShape(Capsule())
        }
        .buttonStyle(.plain)
    }
}

#Preview {
    NavigationStack { CatalogView() }
        .environmentObject(AppState())
}
