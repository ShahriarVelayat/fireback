import Promises
import Combine
import SwiftUI
func GetModbusConnectionTypesExportFetcher() -> AnyPublisher<ArrayResponse<ModbusConnectionTypeEntity>, Error> {
    var computedUrl = "http://localhost:61901/modbus-connection-types/export"
    var request = URLRequest(url: URL(string: computedUrl)!)
    print("Token:", AuthService.shared.TokenSnapShot)
    request.addValue(AuthService.shared.TokenSnapShot, forHTTPHeaderField: "Authorization")
    request.addValue("root", forHTTPHeaderField: "workspace-id")
    return URLSession.shared
        .dataTaskPublisher(for: request)
        .map(\.data)
        .decode(type: ArrayResponse<ModbusConnectionTypeEntity>.self, decoder: JSONDecoder())
        .receive(on: DispatchQueue.main)
        .eraseToAnyPublisher()
}
/*
I have commented this for now, because this is not returning correctly, as well as the function above is good enough
func GetModbusConnectionTypesExport() -> Promise<[ModbusConnectionTypeEntity]?> {
    return Promise<[ModbusConnectionTypeEntity]?>(on: .main) { fulfill, reject in
    var computedUrl = "http://localhost:61901/modbus-connection-types/export"
        var request = URLRequest(url: URL(string: computedUrl)!)
        request.httpMethod = "GET"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                let decoder = JSONDecoder()
                if let str = String(bytes: data, encoding: .utf8) {
                    print(str)
                }
                do {
                    let result = try decoder.decode(SingleResponse<ModbusConnectionTypeEntity>.self, from: data)
                    if result.error != nil {
                        reject(result.error!)
                    } else {
                        fulfill(result.data)
                    }
                } catch {
                    print(error)
                }
            }
        }.resume()
    }
}
*/