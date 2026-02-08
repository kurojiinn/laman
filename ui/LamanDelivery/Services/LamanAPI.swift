import Foundation

final class LamanAPI {
    private let baseURL: URL
    private let session: URLSession

    init(baseURL: URL = URL(string: "http://localhost:8080")!, session: URLSession = .shared) {
        self.baseURL = baseURL
        self.session = session
    }

    func getCategories() async throws -> [Category] {
        let url = baseURL.appendingPathComponent("api/v1/catalog/categories")
        return try await fetch(url: url, responseType: [Category].self)
    }

    func getProducts() async throws -> [Product] {
        var components = URLComponents(url: baseURL.appendingPathComponent("api/v1/catalog/products"), resolvingAgainstBaseURL: false)
        components?.queryItems = [URLQueryItem(name: "available_only", value: "true")]
        guard let url = components?.url else { throw LamanAPIError.invalidURL }
        return try await fetch(url: url, responseType: [Product].self)
    }

    func createOrder(request: CreateOrderRequest) async throws -> Order {
        let url = baseURL.appendingPathComponent("api/v1/orders")
        var req = URLRequest(url: url)
        req.httpMethod = "POST"
        req.setValue("application/json", forHTTPHeaderField: "Content-Type")
        req.httpBody = try JSONEncoder.laman.encode(request)

        let (data, response) = try await session.data(for: req)
        try validate(response: response, data: data)
        return try JSONDecoder.laman.decode(Order.self, from: data)
    }

    private func fetch<T: Decodable>(url: URL, responseType: T.Type) async throws -> T {
        let (data, response) = try await session.data(from: url)
        try validate(response: response, data: data)
        return try JSONDecoder.laman.decode(T.self, from: data)
    }

    private func validate(response: URLResponse, data: Data) throws {
        guard let http = response as? HTTPURLResponse else {
            throw LamanAPIError.invalidResponse
        }
        guard (200...299).contains(http.statusCode) else {
            let message = String(data: data, encoding: .utf8) ?? "HTTP \(http.statusCode)"
            throw LamanAPIError.serverError(message)
        }
    }
}

enum LamanAPIError: LocalizedError {
    case invalidURL
    case invalidResponse
    case serverError(String)

    var errorDescription: String? {
        switch self {
        case .invalidURL:
            return "Неверный URL"
        case .invalidResponse:
            return "Некорректный ответ сервера"
        case .serverError(let message):
            return message
        }
    }
}

extension JSONDecoder {
    static var laman: JSONDecoder {
        let decoder = JSONDecoder()
        decoder.keyDecodingStrategy = .convertFromSnakeCase
        decoder.dateDecodingStrategy = .iso8601
        return decoder
    }
}

extension JSONEncoder {
    static var laman: JSONEncoder {
        let encoder = JSONEncoder()
        encoder.keyEncodingStrategy = .convertToSnakeCase
        return encoder
    }
}
