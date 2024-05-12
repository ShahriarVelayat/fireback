import Promises
import Combine
import SwiftUI
 
func {{ .r.GetFuncNameUpper}}Fetcher() -> AnyPublisher<ArrayResponse<{{ .r.ResponseEntityComputed }}>, Error> {
    {{ template "rpcActionCommon" .r }}
    request.addValue(AuthService.shared.TokenSnapShot, forHTTPHeaderField: "Authorization")
    request.addValue("root", forHTTPHeaderField: "workspace-id")
        
    return URLSession.shared
        .dataTaskPublisher(for: request)
        .map(\.data)
        .decode(type: ArrayResponse<{{ .r.ResponseEntityComputed }}>.self, decoder: JSONDecoder())
        .receive(on: DispatchQueue.main)
        .eraseToAnyPublisher()
}

/*

I have commented this for now, because this is not returning correctly, as well as the function above is good enough

func {{ .r.GetFuncNameUpper}}() -> Promise<[{{ .r.ResponseEntityComputed}}]?> {
    
    return Promise<[{{ .r.ResponseEntityComputed}}]?>(on: .main) { fulfill, reject in
    
        {{ template "rpcActionCommon" .r }}
        request.httpMethod = "GET"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")

        URLSession.shared.dataTask(with: request) { data, response, error in
            if let data = data {
                let decoder = JSONDecoder()
                
                if let str = String(bytes: data, encoding: .utf8) {
                    print(str)
                }

                do {
                    let result = try decoder.decode(SingleResponse<{{ .r.ResponseEntityComputed}}>.self, from: data)
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