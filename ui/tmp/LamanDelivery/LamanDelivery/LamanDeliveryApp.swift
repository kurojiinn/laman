import SwiftUI

@main
struct LamanDeliveryApp: App {
    @StateObject private var appState: AppState
    @StateObject private var catalogVM: CatalogViewModel

    init() {
        let appState = AppState()
        _appState = StateObject(wrappedValue: appState)
        _catalogVM = StateObject(wrappedValue: CatalogViewModel(appState: appState))
    }

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environmentObject(appState)
                .environmentObject(catalogVM)
        }
    }
}
