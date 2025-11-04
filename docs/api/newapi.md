# GPT-4o-image 图像生成

>  - 异步处理模式，返回任务ID用于后续查询
- 支持文本转图片、图生图、图像编辑等多种生成模式。
- 生成的图像链接，有效期为24小时，请尽快保存 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1/images/generations \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
      "model": "gpt-4o-image",
      "prompt": "星空下的古老城堡",
      "size": "1:1",
      "n": 1,
      "image_urls": [
        "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"
      ]
    }'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/images/generations"

  payload = {
      "model": "gpt-4o-image",
      "prompt": "星空下的古老城堡",
      "size": "1:1",
      "n": 1,
      "image_urls": [
          "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"    ]}

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/images/generations";

  const payload = {
    model: "gpt-4o-image",
    prompt: "星空下的古老城堡",
    size: "1:1",
    n: 1,
    image_urls: [
      "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"  ]};

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1/images/generations"

      payload := map[string]interface{}{
          "model":  "gpt-4o-image",
          "prompt": "星空下的古老城堡",
          "size":   "1:1",
          "n":      1,
          "image_urls": []string{
              "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"
    }    }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/images/generations";

          String payload = """
          {
            "model": "gpt-4o-image",
            "prompt": "星空下的古老城堡",
            "size": "1:1",
            "n": 1,
            "image_urls": [
              "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"
            ]        }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/images/generations";

  $payload = [
      "model" => "gpt-4o-image",
      "prompt" => "星空下的古老城堡",
      "size" => "1:1",
      "n" => 1,
      "image_urls" => [
          "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"    ]];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/images/generations")

  payload = {
    model: "gpt-4o-image",
    prompt: "星空下的古老城堡",
    size: "1:1",
    n: 1,
    image_urls: [
      "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"  ]}

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/images/generations")!

  let payload: [String: Any] = [
      "model": "gpt-4o-image",
      "prompt": "星空下的古老城堡",
      "size": "1:1",
      "n": 1,
      "image_urls": [
          "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"    ]]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/images/generations";

          var payload = @"{
              ""model"": ""gpt-4o-image"",
              ""prompt"": ""星空下的古老城堡"",
              ""size"": ""1:1"",
              ""n"": 1,
              ""image_urls"": [
                  ""https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"",
                  ""
              ]        }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/images/generations";
          const char *payload = "{"
              "\"model\":\"gpt-4o-image\","
              "\"prompt\":\"星空下的古老城堡\","
              "\"size\":\"1:1\","
              "\"n\":1,"
              "\"image_urls\":[\"https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png\"]"
          "}";

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, "Authorization: Bearer <token>");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/images/generations"];
          
          NSDictionary *payload = @{
              @"model": @"gpt-4o-image",
              @"prompt": @"星空下的古老城堡",
              @"size": @"1:1",
              @"n": @1,
              @"image_urls": @[
                  @"https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png",
                  @
              ]        };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/images/generations"

  let payload = {|{
    "model": "gpt-4o-image",
    "prompt": "星空下的古老城堡",
    "size": "1:1",
    "n": 1,
    "image_urls": [
      "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"  ]}|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/images/generations');
    
    final payload = {
      'model': 'gpt-4o-image',
      'prompt': '星空下的古老城堡',
      'size': '1:1',
      'n': 1,
      'image_urls': [
        'https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png'    ]  };
    
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
        'Content-Type': 'application/json'
    },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/images/generations"

  payload <- list(
    model = "gpt-4o-image",
    prompt = "星空下的古老城堡",
    size = "1:1",
    n = 1,
    image_urls = list(
      "https://cdn.apimart.ai/doc/9998238784524590-549804af-14bc-4fbc-bae5-9d8469d35de3-image_task_01K889V97T2YF8RHQW21XMVMS1_0.png"  ))

  response <- POST(
    url,
    add_headers(
      Authorization = "Bearer <token>",
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": [
      {
        "status": "submitted",
        "task_id": "task_01K8SGYNNNVBQTXNR4MM964S7K"
      }
    ]
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" default="gpt-4o-image" required>
  图像生成模型名称

  Example: `"gpt-4o-image"`
</ParamField>

<ParamField body="prompt" type="string" required>
  图像生成的文本描述

  最长 1000 个字符
</ParamField>

<ParamField body="size" type="string">
  图像生成的尺寸

  支持的格式：

  * 比例：`1:1`、`2:3`、`3:2`
</ParamField>

<ParamField body="n" type="integer">
  生成图像的数量

  支持 1、2、4，会根据生成数量进行预扣费
</ParamField>

<ParamField body="image_urls" type="array">
  参考图像的 URL 列表，用于图生图或图像编辑

  * 最多 5 张图片
  * 每张图片不超过 10MB
  * 支持的格式：.jpeg、.jpg、.png、.webp
</ParamField>

<ParamField body="mask_url" type="string">
  蒙版图像的 URL

  * 必须是 PNG 格式
  * 大小需要与参考图像一致
  * 不超过 4MB
</ParamField>

## Response

<ResponseField name="code" type="integer">
  响应状态码
</ResponseField>

<ResponseField name="data" type="array">
  返回数据数组

  <Expandable title="属性">
    <ResponseField name="status" type="string">
      任务状态

      * `submitted` - 已提交
    </ResponseField>

    <ResponseField name="task_id" type="string">
      任务唯一标识符
    </ResponseField>
  </Expandable>
</ResponseField>


# Gemini-2.5-Flash-Image（Nano banana）图像生成

>  - 异步处理模式，返回任务ID用于后续查询
- 快速生成速度，针对快速图像创建进行优化
- 生成的图像链接，有效期为24小时，请尽快保存 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1/images/generations \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
      "model": "gemini-2.5-flash-image-preview",
      "prompt": "月光下的竹林小径",
      "size": "1:1",
      "n": 1
    }'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/images/generations"

  payload = {
      "model": "gemini-2.5-flash-image-preview",
      "prompt": "月光下的竹林小径",
      "size": "1:1",
      "n": 1
    }

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/images/generations";

  const payload = {
    model: "gemini-2.5-flash-image-preview",
    prompt: "月光下的竹林小径",
    size: "1:1",
    n: 1
    };

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1/images/generations"

      payload := map[string]interface{}{
          "model":  "gemini-2.5-flash-image-preview",
          "prompt": "月光下的竹林小径",
          "size":   "1:1",
          "n":      1string{
              "https://example.com/image1.png",
              "https://example.com/image2.png"
    }
    }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/images/generations";

          String payload = """
          {
            "model": "gemini-2.5-flash-image-preview",
            "prompt": "月光下的竹林小径",
            "size": "1:1",
            "n": 1
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/images/generations";

  $payload = [
      "model" => "gemini-2.5-flash-image-preview",
      "prompt" => "月光下的竹林小径",
      "size" => "1:1",
      "n" => 1
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/images/generations")

  payload = {
    model: "gemini-2.5-flash-image-preview",
    prompt: "月光下的竹林小径",
    size: "1:1",
    n: 1
    }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/images/generations")!

  let payload: [String: Any] = [
      "model": "gemini-2.5-flash-image-preview",
      "prompt": "月光下的竹林小径",
      "size": "1:1",
      "n": 1
      ]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/images/generations";

          var payload = @"{
              ""model"": ""gemini-2.5-flash-image-preview"",
              ""prompt"": ""月光下的竹林小径"",
              ""size"": ""1:1"",
              ""n"": 1
          }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/images/generations";
          const char *payload = "{"
              "\"model\":\"gemini-2.5-flash-image-preview\","
              "\"prompt\":\"月光下的竹林小径\","
              "\"size\":\"1:1\","
              "\"n\":1,"
              ""
          "}";

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, "Authorization: Bearer <token>");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/images/generations"];
          
          NSDictionary *payload = @{
              @"model": @"gemini-2.5-flash-image-preview",
              @"prompt": @"月光下的竹林小径",
              @"size": @"1:1",
              @"n": @1
          };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/images/generations"

  let payload = {|{
    "model": "gemini-2.5-flash-image-preview",
    "prompt": "月光下的竹林小径",
    "size": "1:1",
    "n": 1
  }|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/images/generations');
    
    final payload = {
      'model': 'gemini-2.5-flash-image-preview',
      'prompt': '月光下的竹林小径',
      'size': '1:1',
      'n': 1
    };
    
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
        'Content-Type': 'application/json'
    },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/images/generations"

  payload <- list(
    model = "gemini-2.5-flash-image-preview",
    prompt = "月光下的竹林小径",
    size = "1:1",
    n = 1
  )

  response <- POST(
    url,
    add_headers(
      Authorization = "Bearer <token>",
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": [
      {
        "status": "submitted",
        "task_id": "task_01K8SGYNNNVBQTXNR4MM964S7K"
      }
    ]
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" default="gemini-2.5-flash-image-preview" required>
  图像生成模型名称

  Example: `"gemini-2.5-flash-image-preview"`
</ParamField>

<ParamField body="prompt" type="string" required>
  图像生成的文本描述

  最长 1000 个字符
</ParamField>

<ParamField body="size" type="string">
  图像生成的尺寸

  支持的格式：

  * 比例：`1:1`、`2:3`、`3:2`、`3:4`、`4:3`、`4:5`、`5:4`、`9:16`、`16:9`、`21:9`
</ParamField>

<ParamField body="n" type="integer" default="1">
  生成图像的数量

  固定为 1
</ParamField>

<ParamField body="image_urls" type="array">
  参考图像的 URL 列表，用于图生图或图像编辑

  * 最多 5 张图片
  * 每张图片不超过 10MB
  * 支持的格式：.jpeg、.jpg、.png、.webp
</ParamField>

<ParamField body="mask_url" type="string">
  蒙版图像的 URL

  * 必须是 PNG 格式
  * 大小需要与参考图像一致
  * 不超过 4MB
</ParamField>

## Response

<ResponseField name="code" type="integer">
  响应状态码
</ResponseField>

<ResponseField name="data" type="array">
  返回数据数组

  <Expandable title="属性">
    <ResponseField name="status" type="string">
      任务状态

      * `submitted` - 已提交
    </ResponseField>

    <ResponseField name="task_id" type="string">
      任务唯一标识符
    </ResponseField>
  </Expandable>
</ResponseField>


# Sora2 视频生成

>  - 异步处理模式，返回任务ID用于后续查询
- 支持文本转视频、图生视频等多种生成模式
- 生成的视频链接，有效期为24小时，请尽快保存 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1/videos/generations \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
      "model": "sora-2",
      "prompt": "瀑布倾泻而下形成彩虹",
      "duration": 10,
      "aspect_ratio": "16:9",
      "image_urls": ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
    }'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/videos/generations"

  payload = {
      "model": "sora-2",
      "prompt": "瀑布倾泻而下形成彩虹",
      "duration": 10,
      "aspect_ratio": "16:9",
      "image_urls": ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
  }

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/videos/generations";

  const payload = {
    model: "sora-2",
    prompt: "瀑布倾泻而下形成彩虹",
    duration: 15,
    aspect_ratio: "16:9",
    image_urls: ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
  };

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1/videos/generations"

      payload := map[string]interface{}{
          "model":      "sora-2",
          "prompt":     "瀑布倾泻而下形成彩虹",
          "duration":   15,
          "aspect_ratio": "16:9",
          "image_urls":  ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"],
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/videos/generations";

          String payload = """
          {
            "model": "sora-2",
            "prompt": "瀑布倾泻而下形成彩虹",
            "duration": 10,
            "aspect_ratio": "16:9",
            "image_urls": ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/videos/generations";

  $payload = [
      "model" => "sora-2",
      "prompt" => "瀑布倾泻而下形成彩虹",
      "duration" => 15,
      "aspect_ratio" => "16:9",
      "image_urls" => ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/videos/generations")

  payload = {
    model: "sora-2",
    prompt: "瀑布倾泻而下形成彩虹",
    duration: 15,
    aspect_ratio: "16:9",
    image_urls: ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
  }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/videos/generations")!

  let payload: [String: Any] = [
      "model": "sora-2",
      "prompt": "瀑布倾泻而下形成彩虹",
      "duration": 10,
      "aspect_ratio": "16:9",
      "image_urls": ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
  ]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/videos/generations";

          var payload = @"{
              ""model"": ""sora-2"",
              ""prompt"": ""瀑布倾泻而下形成彩虹"",
              ""duration"": 15,
              ""aspect_ratio"": ""16:9"",
              ""image_urls"": [""https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png""]
          }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/videos/generations";
          const char *payload = "{"
              "\"model\":\"sora-2\","
              "\"prompt\":\"瀑布倾泻而下形成彩虹\","
              "\"duration\":15,"
              "\"aspect_ratio\":\"16:9\","
              "\"image_urls\":[\"https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png\"]"
          "}";

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, "Authorization: Bearer <token>");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/videos/generations"];
          
          NSDictionary *payload = @{
              @"model": @"sora-2",
              @"prompt": @"瀑布倾泻而下形成彩虹",
              @"duration": @15,
              @"aspect_ratio": @"16:9",
              @"image_urls": @[@"https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
          };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/videos/generations"

  let payload = {|{
    "model": "sora-2",
    "prompt": "瀑布倾泻而下形成彩虹",
    "duration": 10,
          "aspect_ratio": "16:9",
            "image_urls": ["https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png"]
  }|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/videos/generations');
    
    final payload = {
      'model': 'sora-2',
      'prompt': '瀑布倾泻而下形成彩虹',
      'duration': 15,
      'aspect_ratio': '16:9',
      'image_urls': ['https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png']
    };
    
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/videos/generations"

  payload <- list(
    model = "sora-2",
    prompt = "瀑布倾泻而下形成彩虹",
    duration = 15,
    aspect_ratio = "16:9",
    image_urls = c("https://cdn.apimart.ai/doc/9998238782946594-f62f70ce-348c-4b13-bb5f-15f17bee676b-image_task_01K88BEGZHVJWJ3ZV6HY99SWQR_0.png")
  )

  response <- POST(
    url,
    add_headers(
      Authorization = "Bearer <token>",
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": [
      {
        "status": "submitted",
        "task_id": "task_01K8SGYNNNVBQTXNR4MM964S7K"
      }
    ]
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" default="sora-2" required>
  视频生成模型名称

  Example: `"sora-2"`
</ParamField>

<ParamField body="prompt" type="string" required>
  视频生成的文本描述

  最长 1000 个字符
</ParamField>

<ParamField body="duration" type="integer">
  视频时长（秒）

  支持的值：`10` 或 `15`

  Example: `10`
</ParamField>

<ParamField body="aspect_ratio" type="string">
  视频分辨率

  支持的格式：

  * `16:9` (横屏)
  * `9:16` (竖屏)
</ParamField>

<ParamField body="image_urls" type="array">
  用于图生视频的参考图像 URL 数组

  * 每个元素必须是有效的图像 URL
  * 支持的格式：.jpeg、.jpg、.png、.webp
  * 最大文件大小：10MB
</ParamField>

<ParamField body="watermark" type="boolean" default="false">
  是否在生成的视频中添加水印

  * `false`：不添加水印
  * `true`：在视频中添加Sora官方水印

  默认值：`false`
</ParamField>

## Response

<ResponseField name="code" type="integer">
  响应状态码，成功时为 200
</ResponseField>

<ResponseField name="data" type="array">
  返回数据数组

  <Expandable title="数组元素">
    <ResponseField name="status" type="string">
      任务状态，初始提交时为 `submitted`
    </ResponseField>

    <ResponseField name="task_id" type="string">
      任务唯一标识符，用于查询任务状态和结果
    </ResponseField>
  </Expandable>
</ResponseField>


# VEO3 视频生成

>  - 异步处理模式，返回任务ID用于后续查询
- 支持文本转视频、图生视频等多种生成模式
- 生成的视频链接，有效期为24小时，请尽快保存 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1/videos/generations \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
      "model": "veo3.1-fast",
      "prompt": "海豚在碧蓝海洋中跳跃",
      "duration": 8,
      "aspect_ratio": "16:9",
      "image_urls": ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
    }'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/videos/generations"

  payload = {
      "model": "veo3.1-fast",
      "prompt": "海豚在碧蓝海洋中跳跃",
      "duration": 8,
      "aspect_ratio": "16:9",
      "image_urls": ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
  }

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/videos/generations";

  const payload = {
    model: "veo3.1",
    prompt: "海豚在碧蓝海洋中跳跃",
    duration: 8,
    aspect_ratio: "16:9",
    image_urls: ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
  };

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1/videos/generations"

      payload := map[string]interface{}{
          "model":      "veo3.1",
          "prompt":     "海豚在碧蓝海洋中跳跃",
          "duration":   8,
          "aspect_ratio": "16:9",
          "image_urls":  ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"],
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/videos/generations";

          String payload = """
          {
            "model": "veo3.1-fast",
            "prompt": "海豚在碧蓝海洋中跳跃",
            "duration": 8,
            "aspect_ratio": "16:9",
            "image_urls": ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/videos/generations";

  $payload = [
      "model" => "veo3.1",
      "prompt" => "海豚在碧蓝海洋中跳跃",
      "duration" => 8,
      "aspect_ratio" => "16:9",
      "image_urls" => ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/videos/generations")

  payload = {
    model: "veo3.1",
    prompt: "海豚在碧蓝海洋中跳跃",
    duration: 8,
    aspect_ratio: "16:9",
    image_urls: ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
  }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/videos/generations")!

  let payload: [String: Any] = [
      "model": "veo3.1-fast",
      "prompt": "海豚在碧蓝海洋中跳跃",
      "duration": 8,
      "aspect_ratio": "16:9",
      "image_urls": ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
  ]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/videos/generations";

          var payload = @"{
              ""model"": ""veo3.1"",
              ""prompt"": ""海豚在碧蓝海洋中跳跃"",
              ""duration"": 8,
              ""aspect_ratio"": ""16:9"",
              ""image_url"": ""https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png""
          }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/videos/generations";
          const char *payload = "{"
              "\"model\":\"veo3.1-fast\","
              "\"prompt\":\"海豚在碧蓝海洋中跳跃\","
              "\"duration\":8,"
              "\"aspect_ratio\":\"16:9\","
              "\"image_urls\":[\"https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png\"]"
          "}";

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, "Authorization: Bearer <token>");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/videos/generations"];
          
          NSDictionary *payload = @{
              @"model": @"veo3.1",
              @"prompt": @"海豚在碧蓝海洋中跳跃",
              @"duration": @8,
              @"aspect_ratio": @"16:9",
              @"image_urls": @[@"https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
          };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/videos/generations"

  let payload = {|{
    "model": "veo3.1",
    "prompt": "海豚在碧蓝海洋中跳跃",
    "duration": 8,
    "aspect_ratio": "16:9",
              "image_urls": ["https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png"]
  }|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/videos/generations');
    
    final payload = {
      'model': 'veo3.1',
      'prompt': '海豚在碧蓝海洋中跳跃',
      'duration': 8,
      'aspect_ratio': '16:9',
      'image_urls': ['https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png']
    };
    
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/videos/generations"

  payload <- list(
    model = "veo3.1",
    prompt = "海豚在碧蓝海洋中跳跃",
    duration = 8,
    aspect_ratio = "16:9",
    image_urls = c("https://cdn.apimart.ai/doc/9998238783208208-9972597b-255d-4e7e-9649-e6ee38a837aa-image_task_01K88B53MTK41PP5KGDTG2PA5P_0.png")
  )

  response <- POST(
    url,
    add_headers(
      Authorization = "Bearer <token>",
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": [
      {
        "status": "submitted",
        "task_id": "task_01K8SGYNNNVBQTXNR4MM964S7K"
      }
    ]
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" default="veo3.1-fast" required>
  视频生成模型名称

  可用模型：

  * `veo3.1-fast` - 快速生成模型，适用于快速预览和迭代
  * `veo3.1-quality` - 高质量生成模型，适用于最终制作

  Example: `"veo3.1-fast"`
</ParamField>

<ParamField body="prompt" type="string" required>
  视频生成的文本描述

  最长 1000 个字符
</ParamField>

<ParamField body="duration" type="integer">
  视频时长（秒）

  固定值：`8`（VEO3 仅支持 8 秒时长）
</ParamField>

<ParamField body="aspect_ratio" type="string">
  视频分辨率

  支持的格式：

  * `16:9` (横屏)
  * `9:16` (竖屏)
</ParamField>

<ParamField body="image_urls" type="array">
  用于图生视频的参考图像 URL 数组

  * 每个元素必须是有效的图像 URL
  * 支持的格式：.jpeg、.jpg、.png、.webp
  * 最大文件大小：10MB
</ParamField>

## Response

<ResponseField name="code" type="integer">
  响应状态码
</ResponseField>

<ResponseField name="data" type="array">
  返回数据数组

  <Expandable title="属性">
    <ResponseField name="status" type="string">
      任务状态

      * `submitted` - 已提交
    </ResponseField>

    <ResponseField name="task_id" type="string">
      任务唯一标识符
    </ResponseField>
  </Expandable>
</ResponseField>


# Whisper-1 音频转录

>  - 支持 99 种语言的语音识别
- 多种输出格式：json、text、srt、vtt 等
- 最大文件大小 25 MB 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1/audio/transcriptions \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: multipart/form-data' \
    --form 'file=@/path/to/audio.mp3' \
    --form 'model=whisper-1' \
    --form 'language=zh' \
    --form 'response_format=json'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/audio/transcriptions"

  files = {
      "file": open("/path/to/audio.mp3", "rb")
  }

  data = {
      "model": "whisper-1",
      "language": "zh",
      "response_format": "json"
  }

  headers = {
      "Authorization": "Bearer <token>"
  }

  response = requests.post(url, files=files, data=data, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/audio/transcriptions";

  const formData = new FormData();
  formData.append("file", audioFile);
  formData.append("model", "whisper-1");
  formData.append("language", "zh");
  formData.append("response_format", "json");

  const headers = {
    "Authorization": "Bearer <token>"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: formData
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "fmt"
      "io"
      "mime/multipart"
      "net/http"
      "os"
  )

  func main() {
      url := "https://api.apimart.ai/v1/audio/transcriptions"

      file, _ := os.Open("/path/to/audio.mp3")
      defer file.Close()

      body := &bytes.Buffer{}
      writer := multipart.NewWriter(body)
      
      part, _ := writer.CreateFormFile("file", "audio.mp3")
      io.Copy(part, file)
      
      writer.WriteField("model", "whisper-1")
      writer.WriteField("language", "zh")
      writer.WriteField("response_format", "json")
      writer.Close()

      req, _ := http.NewRequest("POST", url, body)
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", writer.FormDataContentType())

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      responseBody, _ := io.ReadAll(resp.Body)
      fmt.Println(string(responseBody))
  }
  ```

  ```java Java theme={null}
  import java.io.File;
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/audio/transcriptions";
          
          File audioFile = new File("/path/to/audio.mp3");
          
          // 使用 Apache HttpClient 或 OkHttp 库来发送 multipart/form-data 请求
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/audio/transcriptions";

  $file = new CURLFile('/path/to/audio.mp3', 'audio/mpeg', 'audio.mp3');

  $data = [
      "file" => $file,
      "model" => "whisper-1",
      "language" => "zh",
      "response_format" => "json"
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, $data);
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/audio/transcriptions")

  File.open('/path/to/audio.mp3', 'rb') do |file|
    request = Net::HTTP::Post.new(url)
    request["Authorization"] = "Bearer <token>"
    
    form_data = [
      ['file', file, { filename: 'audio.mp3', content_type: 'audio/mpeg' }],
      ['model', 'whisper-1'],
      ['language', 'zh'],
      ['response_format', 'json']
    ]
    
    request.set_form form_data, 'multipart/form-data'
    
    http = Net::HTTP.new(url.host, url.port)
    http.use_ssl = true
    
    response = http.request(request)
    puts response.body
  end
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/audio/transcriptions")!

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")

  let boundary = "Boundary-\(UUID().uuidString)"
  request.setValue("multipart/form-data; boundary=\(boundary)", forHTTPHeaderField: "Content-Type")

  var body = Data()

  // Add file
  let fileURL = URL(fileURLWithPath: "/path/to/audio.mp3")
  if let fileData = try? Data(contentsOf: fileURL) {
      body.append("--\(boundary)\r\n".data(using: .utf8)!)
      body.append("Content-Disposition: form-data; name=\"file\"; filename=\"audio.mp3\"\r\n".data(using: .utf8)!)
      body.append("Content-Type: audio/mpeg\r\n\r\n".data(using: .utf8)!)
      body.append(fileData)
      body.append("\r\n".data(using: .utf8)!)
  }

  // Add other fields
  let fields = ["model": "whisper-1", "language": "zh", "response_format": "json"]
  for (key, value) in fields {
      body.append("--\(boundary)\r\n".data(using: .utf8)!)
      body.append("Content-Disposition: form-data; name=\"\(key)\"\r\n\r\n".data(using: .utf8)!)
      body.append("\(value)\r\n".data(using: .utf8)!)
  }

  body.append("--\(boundary)--\r\n".data(using: .utf8)!)

  request.httpBody = body

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.IO;
  using System.Net.Http;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/audio/transcriptions";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          using var form = new MultipartFormDataContent();
          
          var fileStream = File.OpenRead("/path/to/audio.mp3");
          form.Add(new StreamContent(fileStream), "file", "audio.mp3");
          form.Add(new StringContent("whisper-1"), "model");
          form.Add(new StringContent("zh"), "language");
          form.Add(new StringContent("json"), "response_format");

          var response = await client.PostAsync(url, form);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;
      struct curl_httppost *formpost = NULL;
      struct curl_httppost *lastptr = NULL;
      struct curl_slist *headers = NULL;

      curl_global_init(CURL_GLOBAL_ALL);

      curl_formadd(&formpost, &lastptr,
                   CURLFORM_COPYNAME, "file",
                   CURLFORM_FILE, "/path/to/audio.mp3",
                   CURLFORM_END);

      curl_formadd(&formpost, &lastptr,
                   CURLFORM_COPYNAME, "model",
                   CURLFORM_COPYCONTENTS, "whisper-1",
                   CURLFORM_END);

      curl_formadd(&formpost, &lastptr,
                   CURLFORM_COPYNAME, "language",
                   CURLFORM_COPYCONTENTS, "zh",
                   CURLFORM_END);

      curl_formadd(&formpost, &lastptr,
                   CURLFORM_COPYNAME, "response_format",
                   CURLFORM_COPYCONTENTS, "json",
                   CURLFORM_END);

      curl = curl_easy_init();
      headers = curl_slist_append(headers, "Authorization: Bearer <token>");

      if(curl) {
          curl_easy_setopt(curl, CURLOPT_URL, "https://api.apimart.ai/v1/audio/transcriptions");
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
          curl_easy_setopt(curl, CURLOPT_HTTPPOST, formpost);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_easy_cleanup(curl);
          curl_formfree(formpost);
          curl_slist_free_all(headers);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/audio/transcriptions"];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          
          NSString *boundary = @"Boundary-12345";
          NSString *contentType = [NSString stringWithFormat:@"multipart/form-data; boundary=%@", boundary];
          [request setValue:contentType forHTTPHeaderField:@"Content-Type"];
          
          NSMutableData *body = [NSMutableData data];
          
          // Add file
          NSData *fileData = [NSData dataWithContentsOfFile:@"/path/to/audio.mp3"];
          [body appendData:[[NSString stringWithFormat:@"--%@\r\n", boundary] dataUsingEncoding:NSUTF8StringEncoding]];
          [body appendData:[@"Content-Disposition: form-data; name=\"file\"; filename=\"audio.mp3\"\r\n" dataUsingEncoding:NSUTF8StringEncoding]];
          [body appendData:[@"Content-Type: audio/mpeg\r\n\r\n" dataUsingEncoding:NSUTF8StringEncoding]];
          [body appendData:fileData];
          [body appendData:[@"\r\n" dataUsingEncoding:NSUTF8StringEncoding]];
          
          // Add other fields
          NSDictionary *fields = @{@"model": @"whisper-1", @"language": @"zh", @"response_format": @"json"};
          for (NSString *key in fields) {
              [body appendData:[[NSString stringWithFormat:@"--%@\r\n", boundary] dataUsingEncoding:NSUTF8StringEncoding]];
              [body appendData:[[NSString stringWithFormat:@"Content-Disposition: form-data; name=\"%@\"\r\n\r\n", key] dataUsingEncoding:NSUTF8StringEncoding]];
              [body appendData:[[NSString stringWithFormat:@"%@\r\n", fields[key]] dataUsingEncoding:NSUTF8StringEncoding]];
          }
          
          [body appendData:[[NSString stringWithFormat:@"--%@--\r\n", boundary] dataUsingEncoding:NSUTF8StringEncoding]];
          [request setHTTPBody:body];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/audio/transcriptions"

  (* Note: Multipart form data handling in OCaml requires additional libraries *)
  let () =
    print_endline "使用 multipart_form 库来处理文件上传"
  ```

  ```dart Dart theme={null}
  import 'dart:io';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/audio/transcriptions');
    
    var request = http.MultipartRequest('POST', url);
    request.headers['Authorization'] = 'Bearer <token>';
    
    request.files.add(await http.MultipartFile.fromPath('file', '/path/to/audio.mp3'));
    request.fields['model'] = 'whisper-1';
    request.fields['language'] = 'zh';
    request.fields['response_format'] = 'json';
    
    var response = await request.send();
    var responseData = await response.stream.bytesToString();
    
    print(responseData);
  }
  ```

  ```r R theme={null}
  library(httr)

  url <- "https://api.apimart.ai/v1/audio/transcriptions"

  response <- POST(
    url,
    add_headers(Authorization = "Bearer <token>"),
    body = list(
      file = upload_file("/path/to/audio.mp3"),
      model = "whisper-1",
      language = "zh",
      response_format = "json"
    ),
    encode = "multipart"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "text": "这是一段测试音频的转录文本内容。"
  }
  ```

  ```json 200 (详细格式) theme={null}
  {
    "task": "transcribe",
    "language": "zh",
    "duration": 8.5,
    "text": "这是一段测试音频的转录文本内容。",
    "segments": [
      {
        "id": 0,
        "seek": 0,
        "start": 0.0,
        "end": 3.5,
        "text": "这是一段测试音频",
        "tokens": [50364, 1234, 5678],
        "temperature": 0.0,
        "avg_logprob": -0.3,
        "compression_ratio": 1.2,
        "no_speech_prob": 0.01
      }
    ]
  }
  ```

  ```srt 200 (SRT 字幕格式) theme={null}
  1
  00:00:00,000 --> 00:00:03,500
  这是一段测试音频

  2
  00:00:03,500 --> 00:00:08,500
  的转录文本内容。
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 413 theme={null}
  {
    "error": {
      "code": 413,
      "message": "文件大小超过限制（最大 25MB）",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="file" type="file" required>
  要转录的音频文件

  支持的格式：mp3, mp4, mpeg, mpga, m4a, wav, webm

  最大文件大小：25 MB
</ParamField>

<ParamField body="model" type="string" default="whisper-1" required>
  语音识别模型名称

  Example: `"whisper-1"`
</ParamField>

<ParamField body="language" type="string">
  音频的语言代码（ISO-639-1 格式）

  指定语言可以提高准确率和速度

  支持的语言包括：zh（中文）、en（英文）、ja（日文）、ko（韩文）等 99 种语言

  Example: `"zh"`
</ParamField>

<ParamField body="prompt" type="string">
  可选的文本提示，用于指导模型的转录风格或延续前一段音频

  最长 224 个 tokens
</ParamField>

<ParamField body="response_format" type="string" default="json">
  输出格式

  支持的格式：

  * `json` - JSON 格式（仅包含文本）
  * `text` - 纯文本
  * `srt` - SRT 字幕格式
  * `verbose_json` - 详细的 JSON 格式（包含时间戳和其他元数据）
  * `vtt` - WebVTT 字幕格式
</ParamField>

<ParamField body="temperature" type="number" default="0">
  采样温度，范围 0 到 1

  较高的值（如 0.8）会使输出更随机，较低的值（如 0.2）会使其更加确定和一致
</ParamField>

## Response

<ResponseField name="text" type="string">
  转录后的文本内容
</ResponseField>

<ResponseField name="task" type="string">
  任务类型，固定为 `transcribe`

  仅在 verbose\_json 格式时返回
</ResponseField>

<ResponseField name="language" type="string">
  检测到的或指定的语言代码

  仅在 verbose\_json 格式时返回
</ResponseField>

<ResponseField name="duration" type="number">
  音频时长（秒）

  仅在 verbose\_json 格式时返回
</ResponseField>

<ResponseField name="segments" type="array">
  文本片段数组

  仅在 verbose\_json 格式时返回

  <Expandable title="属性">
    <ResponseField name="id" type="integer">
      片段ID
    </ResponseField>

    <ResponseField name="start" type="number">
      片段开始时间（秒）
    </ResponseField>

    <ResponseField name="end" type="number">
      片段结束时间（秒）
    </ResponseField>

    <ResponseField name="text" type="string">
      片段文本内容
    </ResponseField>

    <ResponseField name="temperature" type="number">
      使用的采样温度
    </ResponseField>

    <ResponseField name="avg_logprob" type="number">
      平均对数概率
    </ResponseField>

    <ResponseField name="compression_ratio" type="number">
      压缩比
    </ResponseField>

    <ResponseField name="no_speech_prob" type="number">
      无语音概率
    </ResponseField>
  </Expandable>
</ResponseField>


# TTS 文字转语音

>  - 支持多种语音模型和音色选择
- 输出高质量音频格式：mp3、opus、aac、flac、wav、pcm
- 最大输入文本 4096 个字符 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1/audio/speech \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
      "model": "gpt-4o-mini-tts",
      "input": "今天天气真好,适合出去散步。",
      "voice": "alloy",
      "response_format": "opus",
      "speed": 1.0
    }' \
    --output speech.opus
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/audio/speech"

  payload = {
      "model": "gpt-4o-mini-tts",
      "input": "今天天气真好,适合出去散步。",
      "voice": "alloy",
      "response_format": "opus",
      "speed": 1.0
  }

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  with open("speech.opus", "wb") as f:
      f.write(response.content)
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/audio/speech";

  const payload = {
    model: "tts-1",
    input: "今天天气真好,适合出去散步。",
    voice: "alloy",
    response_format: "opus",
    speed: 1.0
  };

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.blob())
    .then(blob => {
      const url = window.URL.createObjectURL(blob);
      const a = document.createElement('a');
      a.href = url;
      a.download = 'speech.opus';
      a.click();
    })
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io"
      "net/http"
      "os"
  )

  func main() {
      url := "https://api.apimart.ai/v1/audio/speech"

      payload := map[string]interface{}{
          "model":           "tts-1",
          "input":           "今天天气真好,适合出去散步。",
          "voice":           "alloy",
          "response_format": "opus",
          "speed":           1.0,
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      out, _ := os.Create("speech.opus")
      defer out.Close()

      io.Copy(out, resp.Body)
      fmt.Println("Audio saved to speech.opus")
  }
  ```

  ```java Java theme={null}
  import java.io.FileOutputStream;
  import java.io.InputStream;
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/audio/speech";

          String json = """
          {
              "model": "gpt-4o-mini-tts",
              "input": "今天天气真好,适合出去散步。",
              "voice": "alloy",
              "response_format": "opus",
              "speed": 1.0
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(json))
              .build();

          HttpResponse<InputStream> response = client.send(request,
              HttpResponse.BodyHandlers.ofInputStream());

          try (FileOutputStream fos = new FileOutputStream("speech.opus")) {
              response.body().transferTo(fos);
          }
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/audio/speech";

  $data = [
      "model" => "tts-1",
      "input" => "今天天气真好,适合出去散步。",
      "voice" => "alloy",
      "response_format" => "opus",
      "speed" => 1.0
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  file_put_contents("speech.opus", $response);
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'uri'
  require 'json'

  url = URI("https://api.apimart.ai/v1/audio/speech")

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"

  request.body = {
    model: "tts-1",
    input: "今天天气真好,适合出去散步。",
    voice: "alloy",
    response_format: "opus",
    speed: 1.0
  }.to_json

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  response = http.request(request)

  File.open("speech.opus", "wb") do |file|
    file.write(response.body)
  end
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/audio/speech")!

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")

  let payload: [String: Any] = [
      "model": "gpt-4o-mini-tts",
      "input": "今天天气真好,适合出去散步。",
      "voice": "alloy",
      "response_format": "opus",
      "speed": 1.0
  ]

  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }

      if let data = data {
          let fileURL = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask)[0]
              .appendingPathComponent("speech.opus")
          try? data.write(to: fileURL)
          print("Audio saved to \(fileURL)")
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.IO;
  using System.Net.Http;
  using System.Text;
  using System.Text.Json;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/audio/speech";

          var payload = new
          {
              model = "tts-1",
              input = "今天天气真好,适合出去散步。",
              voice = "alloy",
              response_format = "opus",
              speed = 1.0
          };

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var json = JsonSerializer.Serialize(payload);
          var content = new StringContent(json, Encoding.UTF8, "application/json");

          var response = await client.PostAsync(url, content);
          var audioBytes = await response.Content.ReadAsByteArrayAsync();

          await File.WriteAllBytesAsync("speech.opus", audioBytes);
          Console.WriteLine("Audio saved to speech.opus");
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  size_t write_data(void *ptr, size_t size, size_t nmemb, FILE *stream) {
      return fwrite(ptr, size, nmemb, stream);
  }

  int main(void) {
      CURL *curl;
      CURLcode res;
      struct curl_slist *headers = NULL;

      curl_global_init(CURL_GLOBAL_ALL);
      curl = curl_easy_init();

      if(curl) {
          FILE *fp = fopen("speech.opus", "wb");

          headers = curl_slist_append(headers, "Authorization: Bearer <token>");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          const char *json_data = "{\"model\":\"gpt-4o-mini-tts\",\"input\":\"今天天气真好,适合出去散步。\",\"voice\":\"alloy\",\"response_format\":\"opus\",\"speed\":1.0}";

          curl_easy_setopt(curl, CURLOPT_URL, "https://api.apimart.ai/v1/audio/speech");
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, json_data);
          curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_data);
          curl_easy_setopt(curl, CURLOPT_WRITEDATA, fp);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          fclose(fp);
          curl_easy_cleanup(curl);
          curl_slist_free_all(headers);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/audio/speech"];

          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];

          NSDictionary *payload = @{
              @"model": @"tts-1",
              @"input": @"今天天气真好,适合出去散步。",
              @"voice": @"alloy",
              @"response_format": @"opus",
              @"speed": @1.0
          };

          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload options:0 error:nil];
          [request setHTTPBody:jsonData];

          NSURLSessionDataTask *task = [[NSURLSession sharedSession]
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }

                  NSString *filePath = [NSHomeDirectory() stringByAppendingPathComponent:@"speech.opus"];
                  [data writeToFile:filePath atomically:YES];
                  NSLog(@"Audio saved to %@", filePath);
              }];

          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/audio/speech"

  let json_body = `Assoc [
    ("model", `String "tts-1");
    ("input", `String "今天天气真好,适合出去散步。");
    ("voice", `String "alloy");
    ("response_format", `String "mp3");
    ("speed", `Float 1.0)
  ]

  let () =
    let body = Cohttp_lwt.Body.of_string (Yojson.Safe.to_string json_body) in
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in

    Lwt_main.run (
      Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      let oc = open_out_bin "speech.opus" in
      output_string oc body_str;
      close_out oc;
      print_endline "Audio saved to speech.opus"
    )
  ```

  ```dart Dart theme={null}
  import 'dart:io';
  import 'package:http/http.dart' as http;
  import 'dart:convert';

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/audio/speech');

    final payload = {
      'model': 'tts-1',
      'input': '今天天气真好,适合出去散步。',
      'voice': 'alloy',
      'response_format': 'opus',
      'speed': 1.0
    };

    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
        'Content-Type': 'application/json'
      },
      body: jsonEncode(payload)
    );

    await File('speech.opus').writeAsBytes(response.bodyBytes);
    print('Audio saved to speech.opus');
  }
  ```

  ```r R theme={null}
  library(httr)

  url <- "https://api.apimart.ai/v1/audio/speech"

  payload <- list(
    model = "tts-1",
    input = "今天天气真好,适合出去散步。",
    voice = "alloy",
    response_format = "opus",
    speed = 1.0
  )

  response <- POST(
    url,
    add_headers(
      Authorization = "Bearer <token>",
      `Content-Type` = "application/json"
    ),
    body = payload,
    encode = "json"
  )

  writeBin(content(response, "raw"), "speech.opus")
  cat("Audio saved to speech.opus\n")
  ```
</RequestExample>

<ResponseExample>
  ```binary 200 theme={null}
  二进制音频数据流
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败,请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足,请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 413 theme={null}
  {
    "error": {
      "code": 413,
      "message": "输入文本超过限制(最大 4096 字符)",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁,请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误,请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误,服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key:

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加:

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" required>
  TTS 模型名称

  可选值:

  * `gpt-4o-mini-tts` - GPT-4o Mini TTS 模型(暂不支持 mp3 格式)

  Example: `"gpt-4o-mini-tts"`
</ParamField>

<ParamField body="input" type="string" required>
  要转换为语音的文本内容

  最大长度: 4096 个字符

  Example: `"今天天气真好,适合出去散步。"`
</ParamField>

<ParamField body="voice" type="string" required>
  语音音色选择

  可选音色:

  * `alloy` - 中性、平衡的音色
  * `echo` - 男性、沉稳的音色
  * `fable` - 英式、叙述性的音色
  * `onyx` - 男性、深沉的音色
  * `nova` - 女性、活力的音色
  * `shimmer` - 女性、温柔的音色

  Example: `"alloy"`
</ParamField>

<ParamField body="response_format" type="string" default="mp3">
  音频输出格式

  支持的格式:

  * `mp3` - MP3 格式(默认)
  * `opus` - Opus 格式,用于互联网流媒体
  * `aac` - AAC 格式
  * `flac` - FLAC 格式,无损压缩
  * `wav` - WAV 格式,未压缩
  * `pcm` - PCM 格式,原始音频数据

  Example: `"mp3"`
</ParamField>

<ParamField body="speed" type="number" default="1.0">
  语音播放速度

  范围: 0.25 到 4.0

  * `0.25` - 最慢速度(1/4倍速)
  * `1.0` - 正常速度(默认)
  * `4.0` - 最快速度(4倍速)

  Example: `1.0`
</ParamField>

## Response

成功时返回二进制音频数据流,可直接保存为音频文件或播放。

错误时返回 JSON 格式的错误信息,包含错误代码、消息和类型。


# 获取任务状态

>  - 查询异步任务的执行状态和结果
- 实时状态更新和进度跟踪
- 任务完成时获取生成结果 

<RequestExample>
  ```bash cURL theme={null}
  curl --request GET \
    --url https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt \
    --header 'Authorization: Bearer <token>'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt"

  headers = {
      "Authorization": "Bearer <token>"
  }

  response = requests.get(url, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt";

  const headers = {
    "Authorization": "Bearer <token>"
  };

  fetch(url, {
    method: "GET",
    headers: headers
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt"

      req, _ := http.NewRequest("GET", url, nil)
      req.Header.Set("Authorization", "Bearer <token>")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt";

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .GET()
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt";

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt")

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Get.new(url)
  request["Authorization"] = "Bearer <token>"

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt")!

  var request = URLRequest(url: url)
  request.httpMethod = "GET"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var response = await client.GetAsync(url);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt";

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, "Authorization: Bearer <token>");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt"];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"GET"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt"

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
    in
    
    let response = Client.get ~headers (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt');
    
    final response = await http.get(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
      },
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)

  url <- "https://api.apimart.ai/v1/tasks/task-unified-1757156493-imcg5zqt"

  response <- GET(
    url,
    add_headers(
      Authorization = "Bearer <token>"
    )
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": {
      "id": "task-unified-1757156493-k9m2xpvw",
      "status": "completed",
      "progress": 100,
      "result": {
        "images": [
          {
            "url": "https://example.com/generated-image.png",
            "expires_at": 1757242893
          }
        ]
      },
      "created": 1757156493,
      "completed": 1757156593,
      "estimated_time": 100,
      "actual_time": 100
    }
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "无效的任务ID",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Path parameters

<ParamField path="task_id" type="string" required>
  生成API返回的任务ID
</ParamField>

## Response

<ResponseField name="id" type="string">
  任务唯一标识符
</ResponseField>

<ResponseField name="status" type="string">
  任务状态值：

  * `pending` - 排队等待处理
  * `processing` - 处理中
  * `completed` - 成功完成
  * `failed` - 失败
  * `cancelled` - 用户取消
</ResponseField>

<ResponseField name="progress" type="integer">
  任务进度百分比（0–100）
</ResponseField>

<ResponseField name="result" type="object">
  任务结果，仅在状态为 `completed` 时返回

  <Expandable title="属性">
    <ResponseField name="images" type="array">
      生成的图像对象数组（图像生成任务）
    </ResponseField>

    <ResponseField name="videos" type="array">
      生成的视频对象数组（视频生成任务）
    </ResponseField>
  </Expandable>
</ResponseField>

<ResponseField name="created" type="integer">
  任务创建时间戳
</ResponseField>

<ResponseField name="completed" type="integer">
  任务完成时间戳（仅在完成时存在）
</ResponseField>

<ResponseField name="estimated_time" type="integer">
  预计完成时间（秒）
</ResponseField>

<ResponseField name="actual_time" type="integer">
  实际完成时间（秒）（仅在完成时存在）
</ResponseField>

<ResponseField name="error" type="object">
  错误详情（仅在状态为 `failed` 时存在）

  <Expandable title="属性">
    <ResponseField name="code" type="integer">
      错误代码
    </ResponseField>

    <ResponseField name="message" type="string">
      错误消息
    </ResponseField>

    <ResponseField name="type" type="string">
      错误类型
    </ResponseField>
  </Expandable>
</ResponseField>

# OpenAI 多模态响应接口

>  - 完全兼容 OpenAI Responses API 格式
- 支持文本和图像的多模态输入
- 支持工具扩展：网络搜索、文件搜索、函数调用、远程MCP 

<RequestExample>
  ```bash cURL theme={null}
  curl https://api.apimart.ai/v1/responses \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $OPENAI_API_KEY" \
    -d '{
      "model": "gpt-5",
      "input": [
        {
          "role": "user",
          "content": [
            {
              "type": "input_text",
              "text": "这张图片里有什么？"
            },
            {
              "type": "input_image",
              "image_url": "https://openai-documentation.vercel.app/images/cat_and_otter.png"
            }
          ]
        }
      ]
    }'
  ```

  ```python Python theme={null}
  import requests
  import os

  url = "https://api.apimart.ai/v1/responses"

  payload = {
      "model": "gpt-5",
      "input": [
          {
              "role": "user",
              "content": [
                  {
                      "type": "input_text",
                      "text": "这张图片里有什么？"
                  },
                  {
                      "type": "input_image",
                      "image_url": "https://openai-documentation.vercel.app/images/cat_and_otter.png"
                  }
              ]
          }
      ]
  }

  headers = {
      "Authorization": f"Bearer {os.environ.get('OPENAI_API_KEY')}",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/responses";

  const payload = {
    model: "gpt-5",
    input: [
      {
        role: "user",
        content: [
          {
            type: "input_text",
            text: "这张图片里有什么？"
          },
          {
            type: "input_image",
            image_url: "https://openai-documentation.vercel.app/images/cat_and_otter.png"
          }
        ]
      }
    ]
  };

  const headers = {
    "Authorization": `Bearer ${process.env.OPENAI_API_KEY}`,
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
      "os"
  )

  func main() {
      url := "https://api.apimart.ai/v1/responses"

      payload := map[string]interface{}{
          "model": "gpt-5",
          "input": []map[string]interface{}{
              {
                  "role": "user",
                  "content": []map[string]string{
                      {
                          "type": "input_text",
                          "text": "这张图片里有什么？",
                      },
                      {
                          "type":      "input_image",
                          "image_url": "https://openai-documentation.vercel.app/images/cat_and_otter.png",
                      },
                  },
              },
          },
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/responses";
          String apiKey = System.getenv("OPENAI_API_KEY");

          String payload = """
          {
            "model": "gpt-5",
            "input": [
              {
                "role": "user",
                "content": [
                  {
                    "type": "input_text",
                    "text": "这张图片里有什么？"
                  },
                  {
                    "type": "input_image",
                    "image_url": "https://openai-documentation.vercel.app/images/cat_and_otter.png"
                  }
                ]
              }
            ]
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer " + apiKey)
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/responses";
  $apiKey = getenv('OPENAI_API_KEY');

  $payload = [
      "model" => "gpt-5",
      "input" => [
          [
              "role" => "user",
              "content" => [
                  [
                      "type" => "input_text",
                      "text" => "这张图片里有什么？"
                  ],
                  [
                      "type" => "input_image",
                      "image_url" => "https://openai-documentation.vercel.app/images/cat_and_otter.png"
                  ]
              ]
          ]
      ]
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer " . $apiKey,
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/responses")
  api_key = ENV['OPENAI_API_KEY']

  payload = {
    model: "gpt-5",
    input: [
      {
        role: "user",
        content: [
          {
            type: "input_text",
            text: "这张图片里有什么？"
          },
          {
            type: "input_image",
            image_url: "https://openai-documentation.vercel.app/images/cat_and_otter.png"
          }
        ]
      }
    ]
  }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer #{api_key}"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/responses")!
  let apiKey = ProcessInfo.processInfo.environment["OPENAI_API_KEY"] ?? ""

  let payload: [String: Any] = [
      "model": "gpt-5",
      "input": [
          [
              "role": "user",
              "content": [
                  [
                      "type": "input_text",
                      "text": "这张图片里有什么？"
                  ],
                  [
                      "type": "input_image",
                      "image_url": "https://openai-documentation.vercel.app/images/cat_and_otter.png"
                  ]
              ]
          ]
      ]
  ]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer \(apiKey)", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/responses";
          var apiKey = Environment.GetEnvironmentVariable("OPENAI_API_KEY");

          var payload = @"{
              ""model"": ""gpt-5"",
              ""input"": [
                  {
                      ""role"": ""user"",
                      ""content"": [
                          {
                              ""type"": ""input_text"",
                              ""text"": ""这张图片里有什么？""
                          },
                          {
                              ""type"": ""input_image"",
                              ""image_url"": ""https://openai-documentation.vercel.app/images/cat_and_otter.png""
                          }
                      ]
                  }
              ]
          }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", $"Bearer {apiKey}");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>
  #include <stdlib.h>

  int main(void) {
      CURL *curl;
      CURLcode res;
      const char *api_key = getenv("OPENAI_API_KEY");

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/responses";
          const char *payload = "{"
              "\"model\":\"gpt-5\","
              "\"input\":[{\"role\":\"user\",\"content\":[{\"type\":\"input_text\",\"text\":\"这张图片里有什么？\"},{\"type\":\"input_image\",\"image_url\":\"https://openai-documentation.vercel.app/images/cat_and_otter.png\"}]}]"
          "}";

          char auth_header[256];
          snprintf(auth_header, sizeof(auth_header), "Authorization: Bearer %s", api_key);

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, auth_header);
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/responses"];
          NSString *apiKey = [NSProcessInfo processInfo].environment[@"OPENAI_API_KEY"];
          
          NSDictionary *payload = @{
              @"model": @"gpt-5",
              @"input": @[
                  @{
                      @"role": @"user",
                      @"content": @[
                          @{
                              @"type": @"input_text",
                              @"text": @"这张图片里有什么？"
                          },
                          @{
                              @"type": @"input_image",
                              @"image_url": @"https://openai-documentation.vercel.app/images/cat_and_otter.png"
                          }
                      ]
                  }
              ]
          };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:[NSString stringWithFormat:@"Bearer %@", apiKey] 
              forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/responses"
  let api_key = Sys.getenv "OPENAI_API_KEY"

  let payload = {|{
    "model": "gpt-5",
    "input": [
      {
        "role": "user",
        "content": [
          {
            "type": "input_text",
            "text": "这张图片里有什么？"
          },
          {
            "type": "input_image",
            "image_url": "https://openai-documentation.vercel.app/images/cat_and_otter.png"
          }
        ]
      }
    ]
  }|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" ("Bearer " ^ api_key)
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'dart:io';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/responses');
    final apiKey = Platform.environment['OPENAI_API_KEY'];
    
    final payload = {
      'model': 'gpt-5',
      'input': [
        {
          'role': 'user',
          'content': [
            {
              'type': 'input_text',
              'text': '这张图片里有什么？'
            },
            {
              'type': 'input_image',
              'image_url': 'https://openai-documentation.vercel.app/images/cat_and_otter.png'
            }
          ]
        }
      ]
    };
    
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer $apiKey',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/responses"
  api_key <- Sys.getenv("OPENAI_API_KEY")

  payload <- list(
    model = "gpt-5",
    input = list(
      list(
        role = "user",
        content = list(
          list(
            type = "input_text",
            text = "这张图片里有什么？"
          ),
          list(
            type = "input_image",
            image_url = "https://openai-documentation.vercel.app/images/cat_and_otter.png"
          )
        )
      )
    )
  )

  response <- POST(
    url,
    add_headers(
      Authorization = paste("Bearer", api_key),
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": {
      "id": "resp-9876543210",
      "object": "response",
      "created": 1677652288,
      "model": "gpt-5",
      "choices": [
        {
          "index": 0,
          "message": {
            "role": "assistant",
            "content": "这张图片中有一只猫和一只水獭。它们看起来正在互动，场景非常可爱和温馨。猫咪和水獭似乎相处得很融洽。"
          },
          "finish_reason": "stop"
        }
      ],
      "usage": {
        "prompt_tokens": 156,
        "completion_tokens": 45,
        "total_tokens": 201
      }
    }
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" required>
  模型名称

  支持的模型包括：

  * `gpt-5` - OpenAI 最新多模态模型
  * `GPT-4o-image` - GPT-4 优化版多模态模型
  * `gpt-4-vision` - GPT-4 视觉理解模型
  * 更多模型持续更新中...
</ParamField>

<ParamField body="input" type="array" required>
  输入内容列表

  每个输入项包含：

  * `role`: 角色类型（`user`、`assistant`、`system`）
  * `content`: 内容数组，支持多种类型：
    * `input_text`: 文本输入
    * `input_image`: 图像输入
</ParamField>

<ParamField body="temperature" type="number">
  控制输出随机性，范围 0-2

  * 较低的值（如 0.2）使输出更确定
  * 较高的值（如 1.8）使输出更随机

  默认值：1.0
</ParamField>

<ParamField body="max_tokens" type="integer">
  生成的最大token数量

  不同模型有不同的最大值限制，请参考具体模型文档
</ParamField>

<ParamField body="stream" type="boolean">
  是否使用流式输出

  * `true`: 流式返回（SSE格式）
  * `false`: 一次性返回完整响应

  默认值：false
</ParamField>

<ParamField body="top_p" type="number">
  核采样参数，范围 0-1

  控制生成文本的多样性，建议与 temperature 二选一使用

  默认值：1.0
</ParamField>

<ParamField body="tools" type="array">
  工具列表，用于扩展模型能力

  支持的工具类型：

  * **网络搜索** (`web_search`): 实时搜索互联网信息
  * **文件搜索** (`file_search`): 搜索已上传的文件内容
  * **函数调用** (`function`): 调用自定义函数
  * **远程MCP** (`remote_mcp`): 连接远程模型上下文协议服务

  示例：`[{"type": "web_search"}]`
</ParamField>

## Response

<ResponseField name="id" type="string">
  响应的唯一标识符
</ResponseField>

<ResponseField name="object" type="string">
  对象类型，固定为 `response`
</ResponseField>

<ResponseField name="created" type="integer">
  创建时间戳
</ResponseField>

<ResponseField name="model" type="string">
  实际使用的模型名称
</ResponseField>

<ResponseField name="choices" type="array">
  生成的回复列表

  <Expandable title="属性">
    <ResponseField name="index" type="integer">
      选项索引
    </ResponseField>

    <ResponseField name="message" type="object">
      消息内容

      <Expandable title="属性">
        <ResponseField name="role" type="string">
          角色类型（assistant）
        </ResponseField>

        <ResponseField name="content" type="string">
          生成的文本内容
        </ResponseField>
      </Expandable>
    </ResponseField>

    <ResponseField name="finish_reason" type="string">
      结束原因

      可能的值：

      * `stop` - 自然结束
      * `length` - 达到最大长度
      * `content_filter` - 内容过滤
    </ResponseField>
  </Expandable>
</ResponseField>

<ResponseField name="usage" type="object">
  token使用统计

  <Expandable title="属性">
    <ResponseField name="prompt_tokens" type="integer">
      输入内容的token数
    </ResponseField>

    <ResponseField name="completion_tokens" type="integer">
      生成内容的token数
    </ResponseField>

    <ResponseField name="total_tokens" type="integer">
      总token数
    </ResponseField>
  </Expandable>
</ResponseField>

## 使用示例

### 纯文本输入

```json  theme={null}
{
  "model": "gpt-5",
  "input": [
    {
      "role": "user",
      "content": [
        {
          "type": "input_text",
          "text": "你好，介绍一下人工智能"
        }
      ]
    }
  ]
}
```

### 使用网络搜索工具

```json  theme={null}
{
  "model": "gpt-5",
  "tools": [{"type": "web_search"}],
  "input": "今天有什么正面的新闻？"
}
```

```bash cURL示例 theme={null}
curl "https://api.apimart.ai/v1/responses" \
    -H "Content-Type: application/json" \
    -H "Authorization: Bearer $OPENAI_API_KEY" \
    -d '{
        "model": "gpt-5",
        "tools": [{"type": "web_search"}],
        "input": "今天有什么正面的新闻？"
    }'
```

### 图像理解

```json  theme={null}
{
  "model": "gpt-5",
  "input": [
    {
      "role": "user",
      "content": [
        {
          "type": "input_text",
          "text": "描述这张图片"
        },
        {
          "type": "input_image",
          "image_url": "https://example.com/image.jpg"
        }
      ]
    }
  ]
}
```

### 多图像分析

```json  theme={null}
{
  "model": "gpt-5",
  "input": [
    {
      "role": "user",
      "content": [
        {
          "type": "input_text",
          "text": "比较这两张图片的异同"
        },
        {
          "type": "input_image",
          "image_url": "https://example.com/image1.jpg"
        },
        {
          "type": "input_image",
          "image_url": "https://example.com/image2.jpg"
        }
      ]
    }
  ]
}
```

### Base64编码图像

```json  theme={null}
{
  "model": "gpt-5",
  "input": [
    {
      "role": "user",
      "content": [
        {
          "type": "input_text",
          "text": "分析这张图片"
        },
        {
          "type": "input_image",
          "image_url": "data:image/jpeg;base64,/9j/4AAQSkZJRg..."
        }
      ]
    }
  ]
}
```

### 使用文件搜索工具

```json  theme={null}
{
  "model": "gpt-5",
  "tools": [{"type": "file_search"}],
  "input": "根据已上传的文档，总结公司的季度业绩"
}
```

### 使用函数调用

```json  theme={null}
{
  "model": "gpt-5",
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_weather",
        "description": "获取指定城市的天气信息",
        "parameters": {
          "type": "object",
          "properties": {
            "city": {
              "type": "string",
              "description": "城市名称，例如：北京"
            },
            "unit": {
              "type": "string",
              "enum": ["celsius", "fahrenheit"],
              "description": "温度单位"
            }
          },
          "required": ["city"]
        }
      }
    }
  ],
  "input": "北京今天天气怎么样？"
}
```

### 使用远程MCP

```json  theme={null}
{
  "model": "gpt-5",
  "tools": [
    {
      "type": "remote_mcp",
      "remote_mcp": {
        "url": "https://mcp.example.com/api",
        "auth_token": "your_mcp_token"
      }
    }
  ],
  "input": "查询数据库中的用户信息"
}
```

### 组合使用多个工具

```json  theme={null}
{
  "model": "gpt-5",
  "tools": [
    {"type": "web_search"},
    {"type": "file_search"},
    {
      "type": "function",
      "function": {
        "name": "calculate",
        "description": "执行数学计算",
        "parameters": {
          "type": "object",
          "properties": {
            "expression": {
              "type": "string",
              "description": "数学表达式"
            }
          },
          "required": ["expression"]
        }
      }
    }
  ],
  "input": "搜索最新的比特币价格，并计算100个比特币的总价值"
}
```

## 内容类型说明

### input\_text

文本输入类型

**属性：**

* `type`: 固定为 `"input_text"`
* `text`: 文本内容（字符串）

### input\_image

图像输入类型

**属性：**

* `type`: 固定为 `"input_image"`
* `image_url`: 图像URL或Base64编码的数据URI

**支持的图像格式：**

* JPEG
* PNG
* GIF
* WebP

**图像大小限制：**

* 最大文件大小：20MB
* 推荐分辨率：不超过2048x2048像素

## 工具使用详解

### 网络搜索 (Web Search)

使用网络搜索工具可以让模型访问实时互联网信息。

**配置示例：**

```json  theme={null}
{
  "tools": [{"type": "web_search"}]
}
```

**适用场景：**

* 查询最新新闻和时事
* 获取实时数据（股票、天气、汇率等）
* 搜索最新的技术文档和资料
* 验证事实信息

### 文件搜索 (File Search)

文件搜索工具允许模型在已上传的文档中搜索相关信息。

**配置示例：**

```json  theme={null}
{
  "tools": [{"type": "file_search"}]
}
```

**适用场景：**

* 分析企业内部文档
* 搜索技术规范和手册
* 查询合同和法律文件
* 知识库问答系统

### 函数调用 (Function Calling)

定义自定义函数，让模型能够调用外部API或执行特定操作。

**完整配置示例：**

```json  theme={null}
{
  "tools": [
    {
      "type": "function",
      "function": {
        "name": "get_stock_price",
        "description": "获取股票的实时价格",
        "parameters": {
          "type": "object",
          "properties": {
            "symbol": {
              "type": "string",
              "description": "股票代码，例如：AAPL"
            },
            "currency": {
              "type": "string",
              "enum": ["USD", "CNY"],
              "description": "货币单位",
              "default": "USD"
            }
          },
          "required": ["symbol"]
        }
      }
    }
  ]
}
```

**参数说明：**

* `name`: 函数名称（必需）
* `description`: 函数功能描述（必需）
* `parameters`: 参数定义，使用JSON Schema格式
  * `type`: 参数类型
  * `properties`: 参数属性定义
  * `required`: 必需参数列表

**适用场景：**

* 调用第三方API
* 执行数据库查询
* 触发业务流程
* 与内部系统集成

### 远程MCP (Remote MCP)

连接到远程模型上下文协议（MCP）服务，扩展模型能力。

**配置示例：**

```json  theme={null}
{
  "tools": [
    {
      "type": "remote_mcp",
      "remote_mcp": {
        "url": "https://your-mcp-server.com/api",
        "auth_token": "your_auth_token",
        "timeout": 30
      }
    }
  ]
}
```

**参数说明：**

* `url`: MCP服务器地址（必需）
* `auth_token`: 认证令牌（可选）
* `timeout`: 超时时间（秒），默认30秒

**适用场景：**

* 连接企业级AI服务
* 使用专业领域模型
* 访问受保护的数据源
* 分布式AI系统集成

## 工具响应格式

当模型使用工具时，响应格式会包含工具调用信息：

```json  theme={null}
{
  "id": "resp-123456",
  "object": "response",
  "created": 1677652288,
  "model": "gpt-5",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": null,
        "tool_calls": [
          {
            "id": "call_abc123",
            "type": "function",
            "function": {
              "name": "get_weather",
              "arguments": "{\"city\": \"北京\"}"
            }
          }
        ]
      },
      "finish_reason": "tool_calls"
    }
  ]
}
```

**工具调用流程：**

1. 模型接收用户输入
2. 分析是否需要使用工具
3. 如需要，返回工具调用请求
4. 客户端执行工具调用
5. 将工具结果返回给模型
6. 模型生成最终响应

## 注意事项

1. **图像URL要求**：
   * 必须是公开可访问的URL
   * 或使用Base64编码的Data URI格式

2. **Token计费**：
   * 图像会根据其分辨率消耗相应的tokens
   * 高分辨率图像会自动调整大小以优化成本
   * 工具调用也会消耗额外的tokens

3. **内容顺序**：
   * content数组中的元素顺序会影响模型理解
   * 建议先放置文本指令，再放置图像

4. **多模态组合**：
   * 可以在一个请求中混合多个文本和图像
   * 支持多轮对话，保持上下文连贯性

5. **工具使用限制**：
   * 同时使用多个工具时，模型会智能选择最合适的工具
   * 函数调用需要明确的函数定义和参数说明
   * 网络搜索结果可能受地域和时间限制

6. **API兼容性**：
   * 完全兼容OpenAI Responses API格式
   * 可无缝迁移现有OpenAI代码
   * 支持所有OpenAI工具扩展功能

# Gemini 原生格式

>  - 使用 Google 原生 API 格式调用 Gemini 模型
- 同步处理模式，实时返回对话内容
- 最简化参数，快速上手 

<RequestExample>
  ```bash cURL theme={null}
  curl --request POST \
    --url https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
    "contents": [
      {
        "role": "user",
        "parts": [
          {
            "text": "你好，介绍一下自己"
          }
        ]
      }
    ]
  }'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent"

  payload = {
      "contents": [
          {
              "role": "user",
              "parts": [
                  {
                      "text": "你好，介绍一下自己"
                  }
              ]
          }
      ]
  }

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent";

  const payload = {
    contents: [
      {
        role: "user",
        parts: [
          {
            text: "你好，介绍一下自己"
          }
        ]
      }
    ]
  };

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent"

      payload := map[string]interface{}{
          "contents": []map[string]interface{}{
              {
                  "role": "user",
                  "parts": []map[string]interface{}{
                      {
                          "text": "你好，介绍一下自己",
                      },
                  },
              },
          },
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent";

          String payload = """
          {
            "contents": [
              {
                "role": "user",
                "parts": [
                  {
                    "text": "你好，介绍一下自己"
                  }
                ]
              }
            ]
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent";

  $payload = [
      "contents" => [
          [
              "role" => "user",
              "parts" => [
                  [
                      "text" => "你好，介绍一下自己"
                  ]
              ]
          ]
      ]
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1beta/models/gemini-2.5-pro:generateContent")

  payload = {
    contents: [
      {
        role: "user",
        parts: [
          {
            text: "你好，介绍一下自己"
          }
        ]
      }
    ]
  }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": {
      "candidates": [
        {
          "content": {
            "role": "model",
            "parts": [
              {
                "text": "你好！很高兴能向你介绍我自己。\n\n我是一个大型语言模型，由 Google 训练和开发..."
              }
            ]
          },
          "finishReason": "STOP",
          "index": 0,
          "safetyRatings": [
            {
              "category": "HARM_CATEGORY_HATE_SPEECH",
              "probability": "NEGLIGIBLE"
            }
          ]
        }
      ],
      "promptFeedback": {
        "safetyRatings": [
          {
            "category": "HARM_CATEGORY_HATE_SPEECH",
            "probability": "NEGLIGIBLE"
          }
        ]
      ]
    },
    "usageMetadata": {
      "promptTokenCount": 4,
      "candidatesTokenCount": 611,
      "totalTokenCount": 2422,
      "thoughtsTokenCount": 1807,
      "promptTokensDetails": [
        {
          "modality": "TEXT",
          "tokenCount": 4
        }
      ]
    }
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "无效的请求参数",
      "status": "INVALID_ARGUMENT"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "认证失败，请检查 API Key",
      "status": "UNAUTHENTICATED"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "余额不足，请充值",
      "status": "PAYMENT_REQUIRED"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "没有访问权限",
      "status": "PERMISSION_DENIED"
    }
  }
  ```

  ```json 404 theme={null}
  {
    "error": {
      "code": 404,
      "message": "找不到指定的模型",
      "status": "NOT_FOUND"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后重试",
      "status": "RESOURCE_EXHAUSTED"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误",
      "status": "INTERNAL"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务暂时不可用",
      "status": "BAD_GATEWAY"
    }
  }
  ```

  ```json 503 theme={null}
  {
    "error": {
      "code": 503,
      "message": "服务暂时不可用",
      "status": "UNAVAILABLE"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Path Parameters

<ParamField path="model" type="string" required>
  模型名称

  示例中使用 `gemini-2.5-pro`，您可以将其替换为其他支持的 Gemini 模型：

  * `gemini-2.5-flash` - Gemini 2.5 快速版
  * `gemini-2.5-pro` - Gemini 2.5 专业版
  * `gemini-2.5-flash-lite` - Gemini 2.5 超轻量版
  * `gemini-2.5-pro-thinking` - Gemini 2.5 Pro 深度思考版
</ParamField>

<ParamField path="method" type="enum<string>" required>
  生成方法（快速开始推荐使用 `generateContent`）：

  * `generateContent`: 等待完整响应后一次性返回
  * `streamGenerateContent`: 流式返回，逐块实时返回内容

  可选值：`generateContent`, `streamGenerateContent`
</ParamField>

## Body

<ParamField body="contents" type="array" required>
  对话内容列表

  最少需要1条消息

  <Expandable title="contents 对象结构">
    <ParamField body="role" type="string" required>
      角色类型：

      * `user`: 用户消息
      * `model`: 模型响应（对话历史中使用）
    </ParamField>

    <ParamField body="parts" type="array" required>
      消息内容部分

      <Expandable title="parts 对象结构">
        <ParamField body="text" type="string">
          文本内容
        </ParamField>

        <ParamField body="inlineData" type="object">
          内联数据（用于多模态输入）

          <Expandable title="inlineData 属性">
            <ParamField body="mimeType" type="string">
              MIME 类型，如 `image/jpeg`, `image/png`
            </ParamField>

            <ParamField body="data" type="string">
              Base64 编码的数据
            </ParamField>
          </Expandable>
        </ParamField>
      </Expandable>
    </ParamField>
  </Expandable>

  示例：

  ```json  theme={null}
  [
    {
      "role": "user",
      "parts": [{ "text": "你好，介绍一下自己" }]
    }
  ]
  ```
</ParamField>

<ParamField body="generationConfig" type="object">
  生成配置（可选）

  <Expandable title="generationConfig 属性">
    <ParamField body="temperature" type="number">
      控制输出随机性，范围 0.0-2.0

      * 较低的值使输出更确定
      * 较高的值使输出更随机

      默认值：1.0
    </ParamField>

    <ParamField body="maxOutputTokens" type="integer">
      生成的最大 token 数量

      不同模型有不同的最大值限制
    </ParamField>

    <ParamField body="topP" type="number">
      核采样参数，范围 0.0-1.0

      控制采样时考虑的概率质量
    </ParamField>

    <ParamField body="topK" type="integer">
      Top-K 采样参数

      每步只从概率最高的 K 个 token 中采样
    </ParamField>

    <ParamField body="stopSequences" type="array">
      停止序列列表

      遇到这些序列时停止生成
    </ParamField>
  </Expandable>
</ParamField>

<ParamField body="safetySettings" type="array">
  安全设置（可选）

  <Expandable title="safetySettings 对象结构">
    <ParamField body="category" type="string">
      安全类别：

      * `HARM_CATEGORY_HATE_SPEECH`: 仇恨言论
      * `HARM_CATEGORY_DANGEROUS_CONTENT`: 危险内容
      * `HARM_CATEGORY_HARASSMENT`: 骚扰
      * `HARM_CATEGORY_SEXUALLY_EXPLICIT`: 色情内容
    </ParamField>

    <ParamField body="threshold" type="string">
      阈值级别：

      * `BLOCK_NONE`: 不阻止
      * `BLOCK_ONLY_HIGH`: 仅阻止高风险
      * `BLOCK_MEDIUM_AND_ABOVE`: 阻止中等及以上风险
      * `BLOCK_LOW_AND_ABOVE`: 阻止低等及以上风险
    </ParamField>
  </Expandable>
</ParamField>

## Response

<ResponseField name="candidates" type="array">
  候选响应列表

  <Expandable title="candidates 对象结构">
    <ResponseField name="content" type="object">
      生成的内容

      <Expandable title="content 属性">
        <ResponseField name="role" type="string">
          角色，通常为 `model`
        </ResponseField>

        <ResponseField name="parts" type="array">
          内容部分列表

          <Expandable title="parts 对象">
            <ResponseField name="text" type="string">
              生成的文本内容
            </ResponseField>
          </Expandable>
        </ResponseField>
      </Expandable>
    </ResponseField>

    <ResponseField name="finishReason" type="string">
      完成原因：

      * `STOP`: 正常结束
      * `MAX_TOKENS`: 达到最大 token 限制
      * `SAFETY`: 因安全原因停止
      * `RECITATION`: 因重复内容停止
      * `OTHER`: 其他原因
    </ResponseField>

    <ResponseField name="index" type="integer">
      候选响应的索引
    </ResponseField>

    <ResponseField name="safetyRatings" type="array">
      安全评级列表

      <Expandable title="safetyRatings 对象">
        <ResponseField name="category" type="string">
          安全类别
        </ResponseField>

        <ResponseField name="probability" type="string">
          概率级别：`NEGLIGIBLE`, `LOW`, `MEDIUM`, `HIGH`
        </ResponseField>
      </Expandable>
    </ResponseField>
  </Expandable>
</ResponseField>

<ResponseField name="promptFeedback" type="object">
  提示词反馈信息

  <Expandable title="promptFeedback 属性">
    <ResponseField name="safetyRatings" type="array">
      提示词的安全评级
    </ResponseField>

    <ResponseField name="blockReason" type="string">
      阻止原因（如果提示词被阻止）
    </ResponseField>
  </Expandable>
</ResponseField>

<ResponseField name="usageMetadata" type="object">
  使用量统计

  <Expandable title="usageMetadata 属性">
    <ResponseField name="promptTokenCount" type="integer">
      提示词消耗的 token 数
    </ResponseField>

    <ResponseField name="candidatesTokenCount" type="integer">
      候选响应消耗的 token 数
    </ResponseField>

    <ResponseField name="totalTokenCount" type="integer">
      总消耗 token 数
    </ResponseField>

    <ResponseField name="thoughtsTokenCount" type="integer">
      思考过程消耗的 token 数（如适用）
    </ResponseField>

    <ResponseField name="promptTokensDetails" type="array">
      提示词 token 详情

      <Expandable title="promptTokensDetails 对象">
        <ResponseField name="modality" type="string">
          模态类型：`TEXT`, `IMAGE`, 等
        </ResponseField>

        <ResponseField name="tokenCount" type="integer">
          该模态的 token 数量
        </ResponseField>
      </Expandable>
    </ResponseField>
  </Expandable>
</ResponseField>

# Claude 消息接口

>  - 完全兼容 Claude Messages API 格式
- 支持多轮对话和单次查询
- 支持文本、图像等多模态内容 

<RequestExample>
  ```bash cURL theme={null}
  curl https://api.apimart.ai/v1/messages \
    -H "x-api-key: $API_KEY" \
    -H "anthropic-version: 2023-06-01" \
    -H "content-type: application/json" \
    -d '{
      "model": "claude-sonnet-4-5-20250929",
      "max_tokens": 1024,
      "messages": [
        {"role": "user", "content": "你好，世界"}
      ]
    }'
  ```

  ```python Python theme={null}
  import anthropic

  client = anthropic.Anthropic(
      api_key="YOUR_API_KEY",
      base_url="https://api.apimart.ai"
  )

  message = client.messages.create(
      model="claude-sonnet-4-5-20250929",
      max_tokens=1024,
      messages=[
          {"role": "user", "content": "你好，世界"}
      ]
  )

  print(message.content)
  ```

  ```javascript JavaScript theme={null}
  import Anthropic from '@anthropic-ai/sdk';

  const client = new Anthropic({
    apiKey: process.env.API_KEY,
    baseURL: 'https://api.apimart.ai'
  });

  const message = await client.messages.create({
    model: 'claude-sonnet-4-5-20250929',
    max_tokens: 1024,
    messages: [
      { role: 'user', content: '你好，世界' }
    ]
  });

  console.log(message.content);
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
      "os"
  )

  func main() {
      url := "https://api.apimart.ai/v1/messages"

      payload := map[string]interface{}{
          "model": "claude-sonnet-4-5-20250929",
          "max_tokens": 1024,
          "messages": []map[string]string{
              {
                  "role":    "user",
                  "content": "你好，世界",
              },
          },
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("x-api-key", os.Getenv("API_KEY"))
      req.Header.Set("anthropic-version", "2023-06-01")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/messages";
          String apiKey = System.getenv("API_KEY");

          String payload = """
          {
            "model": "claude-sonnet-4-5-20250929",
            "max_tokens": 1024,
            "messages": [
              {
                "role": "user",
                "content": "你好，世界"
              }
            ]
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("x-api-key", apiKey)
              .header("anthropic-version", "2023-06-01")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/messages";
  $apiKey = getenv('API_KEY');

  $payload = [
      "model" => "claude-sonnet-4-5-20250929",
      "max_tokens" => 1024,
      "messages" => [
          [
              "role" => "user",
              "content" => "你好，世界"
          ]
      ]
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "x-api-key: " . $apiKey,
      "anthropic-version: 2023-06-01",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/messages")
  api_key = ENV['API_KEY']

  payload = {
    model: "claude-sonnet-4-5-20250929",
    max_tokens: 1024,
    messages: [
      {
        role: "user",
        content: "你好，世界"
      }
    ]
  }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["x-api-key"] = api_key
  request["anthropic-version"] = "2023-06-01"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/messages")!
  let apiKey = ProcessInfo.processInfo.environment["API_KEY"] ?? ""

  let payload: [String: Any] = [
      "model": "claude-sonnet-4-5-20250929",
      "max_tokens": 1024,
      "messages": [
          [
              "role": "user",
              "content": "你好，世界"
          ]
      ]
  ]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue(apiKey, forHTTPHeaderField: "x-api-key")
  request.setValue("2023-06-01", forHTTPHeaderField: "anthropic-version")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/messages";
          var apiKey = Environment.GetEnvironmentVariable("API_KEY");

          var payload = @"{
              ""model"": ""claude-sonnet-4-5-20250929"",
              ""max_tokens"": 1024,
              ""messages"": [
                  {
                      ""role"": ""user"",
                      ""content"": ""你好，世界""
                  }
              ]
          }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("x-api-key", apiKey);
          client.DefaultRequestHeaders.Add("anthropic-version", "2023-06-01");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>
  #include <stdlib.h>

  int main(void) {
      CURL *curl;
      CURLcode res;
      const char *api_key = getenv("API_KEY");

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/messages";
          const char *payload = "{"
              "\"model\":\"claude-sonnet-4-5-20250929\","
              "\"max_tokens\":1024,"
              "\"messages\":[{\"role\":\"user\",\"content\":\"你好，世界\"}]"
          "}";

          char auth_header[256];
          snprintf(auth_header, sizeof(auth_header), "x-api-key: %s", api_key);

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, auth_header);
          headers = curl_slist_append(headers, "anthropic-version: 2023-06-01");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/messages"];
          NSString *apiKey = [NSProcessInfo processInfo].environment[@"API_KEY"];
          
          NSDictionary *payload = @{
              @"model": @"claude-sonnet-4-5-20250929",
              @"max_tokens": @1024,
              @"messages": @[
                  @{
                      @"role": @"user",
                      @"content": @"你好，世界"
                  }
              ]
          };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:apiKey forHTTPHeaderField:@"x-api-key"];
          [request setValue:@"2023-06-01" forHTTPHeaderField:@"anthropic-version"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/messages"
  let api_key = Sys.getenv "API_KEY"

  let payload = {|{
    "model": "claude-sonnet-4-5-20250929",
    "max_tokens": 1024,
    "messages": [
      {
        "role": "user",
        "content": "你好，世界"
      }
    ]
  }|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "x-api-key" api_key
      |> fun h -> Header.add h "anthropic-version" "2023-06-01"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'dart:io';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/messages');
    final apiKey = Platform.environment['API_KEY'];
    
    final payload = {
      'model': 'claude-sonnet-4-5-20250929',
      'max_tokens': 1024,
      'messages': [
        {
          'role': 'user',
          'content': '你好，世界'
        }
      ]
    };
    
    final response = await http.post(
      url,
      headers: {
        'x-api-key': apiKey!,
        'anthropic-version': '2023-06-01',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/messages"
  api_key <- Sys.getenv("API_KEY")

  payload <- list(
    model = "claude-sonnet-4-5-20250929",
    max_tokens = 1024,
    messages = list(
      list(
        role = "user",
        content = "你好，世界"
      )
    )
  )

  response <- POST(
    url,
    add_headers(
      `x-api-key` = api_key,
      `anthropic-version` = "2023-06-01",
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": {
      "id": "msg_013Zva2CMHLNnXjNJJKqJ2EF",
      "type": "message",
      "role": "assistant",
      "content": [
        {
          "type": "text",
          "text": "你好！我是Claude。很高兴见到你。"
        }
      ],
      "model": "claude-sonnet-4-5-20250929",
      "stop_reason": "end_turn",
      "stop_sequence": null,
      "usage": {
        "input_tokens": 12,
        "output_tokens": 18
      }
    }
  }
  ```

  ```json 400 theme={null}
  {
    "type": "error",
    "error": {
      "type": "invalid_request_error",
      "message": "请求参数无效"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "type": "error",
    "error": {
      "type": "authentication_error",
      "message": "无效的API密钥"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "type": "error",
    "error": {
      "type": "rate_limit_error",
      "message": "请求过于频繁"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "type": "error",
    "error": {
      "type": "api_error",
      "message": "服务器内部错误"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="x-api-key" type="string" required>
  API密钥用于身份验证

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  在请求头中添加：

  ```
  x-api-key: YOUR_API_KEY
  ```
</ParamField>

<ParamField header="anthropic-version" type="string" required>
  API版本号

  指定要使用的Claude API版本

  示例：`2023-06-01`
</ParamField>

## Body

<ParamField body="model" type="string" required>
  模型名称

  * `claude-haiku-4-5-20251001` - Claude 4.5 快速响应版本
  * `claude-sonnet-4-5-20250929` - Claude 4.5 平衡版本
  * `claude-opus-4-1-20250805` - 最强大的 Claude 4.1 旗舰模型
  * `claude-opus-4-1-20250805-thinking` - Claude 4.1 Opus 深度思考版
  * `claude-sonnet-4-5-20250929-thinking` - Claude 4.5 Sonnet 深度思考版
</ParamField>

<ParamField body="messages" type="array" required>
  消息列表

  消息数组，模型会基于这些消息生成下一条回复。支持交替的 `user` 和 `assistant` 角色。

  每条消息包含：

  * `role`: 角色（`user` 或 `assistant`）
  * `content`: 内容（字符串或内容块数组）

  **单条用户消息示例：**

  ```json  theme={null}
  [{"role": "user", "content": "你好，Claude"}]
  ```

  **多轮对话示例：**

  ```json  theme={null}
  [
    {"role": "user", "content": "你好"},
    {"role": "assistant", "content": "你好！我是Claude。"},
    {"role": "user", "content": "能解释一下AI吗？"}
  ]
  ```

  **预填充助手回复：**

  ```json  theme={null}
  [
    {"role": "user", "content": "太阳的希腊名称是？(A) Sol (B) Helios (C) Sun"},
    {"role": "assistant", "content": "答案是 ("}
  ]
  ```
</ParamField>

<ParamField body="max_tokens" type="integer" required>
  最大生成token数

  生成停止前的最大token数量。模型可能会在达到此限制前停止。

  不同模型有不同的最大值，请参考模型文档。

  最小值：1
</ParamField>

<ParamField body="system" type="string | array">
  系统提示词

  系统提示词用于设置Claude的角色、个性、目标和指令。

  **字符串格式：**

  ```json  theme={null}
  {
    "system": "你是一位专业的Python编程导师"
  }
  ```

  **结构化格式：**

  ```json  theme={null}
  {
    "system": [
      {
        "type": "text",
        "text": "你是一位专业的Python编程导师"
      }
    ]
  }
  ```
</ParamField>

<ParamField body="temperature" type="number">
  温度参数，范围 0-1

  控制输出的随机性：

  * 低值（如0.2）：更确定、更保守
  * 高值（如0.8）：更随机、更有创意

  默认值：1.0
</ParamField>

<ParamField body="top_p" type="number">
  核采样参数，范围 0-1

  使用nucleus sampling。建议使用 `temperature` 或 `top_p` 其中之一，不要同时使用。

  默认值：1.0
</ParamField>

<ParamField body="top_k" type="integer">
  Top-K采样

  只从概率最高的K个选项中采样，用于移除"长尾"低概率响应。

  建议仅在高级用例中使用。
</ParamField>

<ParamField body="stream" type="boolean">
  是否启用流式输出

  设置为 `true` 时，使用服务器发送事件（SSE）流式返回响应。

  默认值：false
</ParamField>

<ParamField body="stop_sequences" type="array">
  停止序列

  自定义文本序列，遇到这些序列时模型将停止生成。

  最多4个序列。

  示例：`["\n\nHuman:", "\n\nAssistant:"]`
</ParamField>

<ParamField body="metadata" type="object">
  元数据

  用于请求的元数据对象。

  包含：

  * `user_id`: 用户标识符
</ParamField>

<ParamField body="tools" type="array">
  工具定义

  工具列表，模型可以调用这些工具来完成任务。

  **函数工具示例：**

  ```json  theme={null}
  {
    "tools": [
      {
        "name": "get_weather",
        "description": "获取指定位置的当前天气",
        "input_schema": {
          "type": "object",
          "properties": {
            "location": {
              "type": "string",
              "description": "城市和省份，例如：北京"
            },
            "unit": {
              "type": "string",
              "enum": ["celsius", "fahrenheit"],
              "description": "温度单位"
            }
          },
          "required": ["location"]
        }
      }
    ]
  }
  ```

  支持的工具类型：

  * 自定义函数工具
  * 计算机使用工具（computer\_20241022）
  * 文本编辑器工具（text\_editor\_20241022）
  * Bash工具（bash\_20241022）
</ParamField>

<ParamField body="tool_choice" type="object">
  工具选择策略

  控制模型如何使用工具：

  * `{"type": "auto"}`: 自动决定（默认）
  * `{"type": "any"}`: 必须使用工具
  * `{"type": "tool", "name": "tool_name"}`: 使用指定工具
</ParamField>

## Response

<ResponseField name="id" type="string">
  唯一消息标识符

  示例：`"msg_013Zva2CMHLNnXjNJJKqJ2EF"`
</ResponseField>

<ResponseField name="type" type="string">
  对象类型

  固定为 `"message"`
</ResponseField>

<ResponseField name="role" type="string">
  角色

  固定为 `"assistant"`
</ResponseField>

<ResponseField name="content" type="array">
  内容块数组

  模型生成的内容，是一个内容块数组。

  **文本内容：**

  ```json  theme={null}
  [{"type": "text", "text": "你好！我是Claude。"}]
  ```

  **工具使用：**

  ```json  theme={null}
  [
    {
      "type": "tool_use",
      "id": "toolu_01A09q90qw90lq917835lq9",
      "name": "get_weather",
      "input": {"location": "北京", "unit": "celsius"}
    }
  ]
  ```

  内容类型：

  * `text`: 文本内容
  * `tool_use`: 工具调用
</ResponseField>

<ResponseField name="model" type="string">
  处理请求的模型

  示例：`"claude-sonnet-4-5-20250929"`
</ResponseField>

<ResponseField name="stop_reason" type="string">
  停止原因

  可能的值：

  * `end_turn`: 自然结束
  * `max_tokens`: 达到最大token数
  * `stop_sequence`: 遇到停止序列
  * `tool_use`: 调用了工具
</ResponseField>

<ResponseField name="stop_sequence" type="string | null">
  触发的停止序列

  如果因停止序列而停止，则为该序列内容；否则为 `null`
</ResponseField>

<ResponseField name="usage" type="object">
  Token使用统计

  <Expandable title="属性">
    <ResponseField name="input_tokens" type="integer">
      输入token数
    </ResponseField>

    <ResponseField name="output_tokens" type="integer">
      输出token数
    </ResponseField>
  </Expandable>
</ResponseField>

## 使用示例

### 基础对话

```python  theme={null}
import anthropic

client = anthropic.Anthropic(
    api_key="YOUR_API_KEY",
    base_url="https://api.apimart.ai"
)

message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=[
        {"role": "user", "content": "解释量子计算的基本原理"}
    ]
)

print(message.content[0].text)
```

### 多轮对话

```python  theme={null}
messages = [
    {"role": "user", "content": "什么是机器学习？"},
    {"role": "assistant", "content": "机器学习是人工智能的一个分支..."},
    {"role": "user", "content": "能举个实际应用的例子吗？"}
]

message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=messages
)
```

### 使用系统提示词

```python  theme={null}
message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    system="你是一位资深的Python开发专家，擅长代码审查和优化建议。",
    messages=[
        {"role": "user", "content": "如何优化这段代码？\n\n[代码]"}
    ]
)
```

### 流式响应

```python  theme={null}
with client.messages.stream(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=[
        {"role": "user", "content": "写一篇关于AI的短文"}
    ]
) as stream:
    for text in stream.text_stream:
        print(text, end="", flush=True)
```

### 工具使用

```python  theme={null}
tools = [
    {
        "name": "get_stock_price",
        "description": "获取股票的实时价格",
        "input_schema": {
            "type": "object",
            "properties": {
                "ticker": {
                    "type": "string",
                    "description": "股票代码，例如：AAPL"
                }
            },
            "required": ["ticker"]
        }
    }
]

message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    tools=tools,
    messages=[
        {"role": "user", "content": "特斯拉的股价是多少？"}
    ]
)

# 处理工具调用
if message.stop_reason == "tool_use":
    tool_use = next(block for block in message.content if block.type == "tool_use")
    print(f"调用工具: {tool_use.name}")
    print(f"参数: {tool_use.input}")
```

### 视觉理解

```python  theme={null}
message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=[
        {
            "role": "user",
            "content": [
                {
                    "type": "image",
                    "source": {
                        "type": "url",
                        "url": "https://example.com/image.jpg"
                    }
                },
                {
                    "type": "text",
                    "text": "描述这张图片"
                }
            ]
        }
    ]
)
```

### Base64图像

```python  theme={null}
import base64

with open("image.jpg", "rb") as image_file:
    image_data = base64.b64encode(image_file.read()).decode("utf-8")

message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=[
        {
            "role": "user",
            "content": [
                {
                    "type": "image",
                    "source": {
                        "type": "base64",
                        "media_type": "image/jpeg",
                        "data": image_data
                    }
                },
                {
                    "type": "text",
                    "text": "分析这张图片"
                }
            ]
        }
    ]
)
```

## 最佳实践

### 1. 提示词工程

**清晰的角色定义：**

```python  theme={null}
system = """你是一位经验丰富的数据科学家，专长包括：
- 统计分析和数据可视化
- 机器学习模型开发
- Python和R编程
请提供专业、准确的建议。"""
```

**结构化输出：**

```python  theme={null}
message = "请以JSON格式返回分析结果，包含summary、key_findings和recommendations字段。"
```

### 2. 错误处理

```python  theme={null}
from anthropic import APIError, RateLimitError

try:
    message = client.messages.create(
        model="claude-sonnet-4-5-20250929",
        max_tokens=1024,
        messages=[{"role": "user", "content": "你好"}]
    )
except RateLimitError:
    print("速率限制，请稍后重试")
except APIError as e:
    print(f"API错误: {e}")
```

### 3. Token优化

```python  theme={null}
# 使用更短的提示词
messages = [
    {"role": "user", "content": "总结要点：\n\n[长文本]"}
]

# 限制输出长度
message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=500,  # 限制输出
    messages=messages
)
```

### 4. 预填充响应

```python  theme={null}
# 引导模型以特定格式回复
messages = [
    {"role": "user", "content": "列出5个Python最佳实践"},
    {"role": "assistant", "content": "以下是5个Python最佳实践：\n\n1."}
]

message = client.messages.create(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=messages
)
```

## 流式响应处理

### Python流式示例

```python  theme={null}
import anthropic

client = anthropic.Anthropic(
    api_key="YOUR_API_KEY",
    base_url="https://api.apimart.ai"
)

with client.messages.stream(
    model="claude-sonnet-4-5-20250929",
    max_tokens=1024,
    messages=[
        {"role": "user", "content": "写一个Python装饰器示例"}
    ]
) as stream:
    for text in stream.text_stream:
        print(text, end="", flush=True)
```

### JavaScript流式示例

```javascript  theme={null}
import Anthropic from '@anthropic-ai/sdk';

const client = new Anthropic({
  apiKey: process.env.API_KEY,
  baseURL: 'https://api.apimart.ai'
});

const stream = await client.messages.stream({
  model: 'claude-sonnet-4-5-20250929',
  max_tokens: 1024,
  messages: [
    { role: 'user', content: '写一个React组件示例' }
  ]
});

for await (const chunk of stream) {
  if (chunk.type === 'content_block_delta' && 
      chunk.delta.type === 'text_delta') {
    process.stdout.write(chunk.delta.text);
  }
}
```

## 注意事项

1. **API密钥安全**：
   * 使用环境变量存储API密钥
   * 不要在代码中硬编码密钥
   * 定期轮换密钥

2. **速率限制**：
   * 注意API的速率限制
   * 实现重试机制
   * 使用指数退避策略

3. **Token管理**：
   * 监控token使用量
   * 优化提示词长度
   * 使用适当的max\_tokens值

4. **模型选择**：
   * Opus: 复杂任务、需要深度思考
   * Sonnet: 平衡性能和成本
   * Haiku: 快速响应、简单任务

5. **内容过滤**：
   * 验证用户输入
   * 过滤敏感信息
   * 实现内容审核机制

# 通用对话接口

>  - 统一的对话API接口，支持所有文本生成模型
- 通过 model 参数选择不同的AI模型
- 兼容 OpenAI Chat Completions API 格式 

<RequestExample>
  ```bash cURL theme={null}

  curl --request POST \
    --url https://api.apimart.ai/v1/chat/completions \
    --header 'Authorization: Bearer <token>' \
    --header 'Content-Type: application/json' \
    --data '{
      "model": "gpt-4o", # 可替换为任意支持的模型 ID
      "messages": [
        {
          "role": "system",
          "content": "你是一个专业的AI助手。"
        },
        {
          "role": "user",
          "content": "介绍一下人工智能的发展历史。"
        }
      ]
    }'
  ```

  ```python Python theme={null}
  import requests

  url = "https://api.apimart.ai/v1/chat/completions"

  payload = {
      "model": "gpt-4o",  # 可替换为任意支持的模型 ID
      "messages": [
          {
              "role": "system",
              "content": "你是一个专业的AI助手。"
          },
          {
              "role": "user",
              "content": "介绍一下人工智能的发展历史。"
          }
      ]
  }

  headers = {
      "Authorization": "Bearer <token>",
      "Content-Type": "application/json"
  }

  response = requests.post(url, json=payload, headers=headers)

  print(response.json())
  ```

  ```javascript JavaScript theme={null}
  const url = "https://api.apimart.ai/v1/chat/completions";

  const payload = {
    model: "gpt-4o",  // 可替换为任意支持的模型 ID
    messages: [
      {
        role: "system",
        content: "你是一个专业的AI助手。"
      },
      {
        role: "user",
        content: "介绍一下人工智能的发展历史。"
      }
    ]
  };

  const headers = {
    "Authorization": "Bearer <token>",
    "Content-Type": "application/json"
  };

  fetch(url, {
    method: "POST",
    headers: headers,
    body: JSON.stringify(payload)
  })
    .then(response => response.json())
    .then(data => console.log(data))
    .catch(error => console.error('Error:', error));
  ```

  ```go Go theme={null}
  package main

  import (
      "bytes"
      "encoding/json"
      "fmt"
      "io/ioutil"
      "net/http"
  )

  func main() {
      url := "https://api.apimart.ai/v1/chat/completions"

      payload := map[string]interface{}{
          "model": "gpt-4o",  // 可替换为任意支持的模型 ID
          "messages": []map[string]string{
              {
                  "role":    "system",
                  "content": "你是一个专业的AI助手。",
              },
              {
                  "role":    "user",
                  "content": "介绍一下人工智能的发展历史。",
              },
          },
      }

      jsonData, _ := json.Marshal(payload)

      req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
      req.Header.Set("Authorization", "Bearer <token>")
      req.Header.Set("Content-Type", "application/json")

      client := &http.Client{}
      resp, err := client.Do(req)
      if err != nil {
          panic(err)
      }
      defer resp.Body.Close()

      body, _ := ioutil.ReadAll(resp.Body)
      fmt.Println(string(body))
  }
  ```

  ```java Java theme={null}
  import java.net.http.HttpClient;
  import java.net.http.HttpRequest;
  import java.net.http.HttpResponse;
  import java.net.URI;

  public class Main {
      public static void main(String[] args) throws Exception {
          String url = "https://api.apimart.ai/v1/chat/completions";

          // 可替换为任意支持的模型 ID
          String payload = """
          {
            "model": "gpt-4o",
            "messages": [
              {
                "role": "system",
                "content": "你是一个专业的AI助手。"
              },
              {
                "role": "user",
                "content": "介绍一下人工智能的发展历史。"
              }
            ]
          }
          """;

          HttpClient client = HttpClient.newHttpClient();
          HttpRequest request = HttpRequest.newBuilder()
              .uri(URI.create(url))
              .header("Authorization", "Bearer <token>")
              .header("Content-Type", "application/json")
              .POST(HttpRequest.BodyPublishers.ofString(payload))
              .build();

          HttpResponse<String> response = client.send(request,
              HttpResponse.BodyHandlers.ofString());

          System.out.println(response.body());
      }
  }
  ```

  ```php PHP theme={null}
  <?php

  $url = "https://api.apimart.ai/v1/chat/completions";

  // 可替换为任意支持的模型 ID
  $payload = [
      "model" => "gpt-4o",
      "messages" => [
          [
              "role" => "system",
              "content" => "你是一个专业的AI助手。"
          ],
          [
              "role" => "user",
              "content" => "介绍一下人工智能的发展历史。"
          ]
      ]
  ];

  $ch = curl_init($url);
  curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
  curl_setopt($ch, CURLOPT_POST, true);
  curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
  curl_setopt($ch, CURLOPT_HTTPHEADER, [
      "Authorization: Bearer <token>",
      "Content-Type: application/json"
  ]);

  $response = curl_exec($ch);
  curl_close($ch);

  echo $response;
  ?>
  ```

  ```ruby Ruby theme={null}
  require 'net/http'
  require 'json'
  require 'uri'

  url = URI("https://api.apimart.ai/v1/chat/completions")

  # 可替换为任意支持的模型 ID
  payload = {
    model: "gpt-4o",
    messages: [
      {
        role: "system",
        content: "你是一个专业的AI助手。"
      },
      {
        role: "user",
        content: "介绍一下人工智能的发展历史。"
      }
    ]
  }

  http = Net::HTTP.new(url.host, url.port)
  http.use_ssl = true

  request = Net::HTTP::Post.new(url)
  request["Authorization"] = "Bearer <token>"
  request["Content-Type"] = "application/json"
  request.body = payload.to_json

  response = http.request(request)
  puts response.body
  ```

  ```swift Swift theme={null}
  import Foundation

  let url = URL(string: "https://api.apimart.ai/v1/chat/completions")!

  let payload: [String: Any] = [
      "model": "gpt-4o",  // 可替换为任意支持的模型 ID
      "messages": [
          [
              "role": "system",
              "content": "你是一个专业的AI助手。"
          ],
          [
              "role": "user",
              "content": "介绍一下人工智能的发展历史。"
          ]
      ]
  ]

  var request = URLRequest(url: url)
  request.httpMethod = "POST"
  request.setValue("Bearer <token>", forHTTPHeaderField: "Authorization")
  request.setValue("application/json", forHTTPHeaderField: "Content-Type")
  request.httpBody = try? JSONSerialization.data(withJSONObject: payload)

  let task = URLSession.shared.dataTask(with: request) { data, response, error in
      if let error = error {
          print("Error: \(error)")
          return
      }
      
      if let data = data, let responseString = String(data: data, encoding: .utf8) {
          print(responseString)
      }
  }

  task.resume()
  ```

  ```csharp C# theme={null}
  using System;
  using System.Net.Http;
  using System.Text;
  using System.Threading.Tasks;

  class Program
  {
      static async Task Main(string[] args)
      {
          var url = "https://api.apimart.ai/v1/chat/completions";

          // 可替换为任意支持的模型 ID
          var payload = @"{
              ""model"": ""gpt-4o"",
              ""messages"": [
                  {
                      ""role"": ""system"",
                      ""content"": ""你是一个专业的AI助手。""
                  },
                  {
                      ""role"": ""user"",
                      ""content"": ""介绍一下人工智能的发展历史。""
                  }
              ]
          }";

          using var client = new HttpClient();
          client.DefaultRequestHeaders.Add("Authorization", "Bearer <token>");

          var content = new StringContent(payload, Encoding.UTF8, "application/json");
          var response = await client.PostAsync(url, content);
          var result = await response.Content.ReadAsStringAsync();

          Console.WriteLine(result);
      }
  }
  ```

  ```c C theme={null}
  #include <stdio.h>
  #include <curl/curl.h>

  int main(void) {
      CURL *curl;
      CURLcode res;

      curl_global_init(CURL_GLOBAL_DEFAULT);
      curl = curl_easy_init();

      if(curl) {
          const char *url = "https://api.apimart.ai/v1/chat/completions";
          // 可替换为任意支持的模型 ID
          const char *payload = "{"
              "\"model\":\"gpt-4o\","
              "\"messages\":[{\"role\":\"system\",\"content\":\"你是一个专业的AI助手。\"},{\"role\":\"user\",\"content\":\"介绍一下人工智能的发展历史。\"}]"
          "}";

          struct curl_slist *headers = NULL;
          headers = curl_slist_append(headers, "Authorization: Bearer <token>");
          headers = curl_slist_append(headers, "Content-Type: application/json");

          curl_easy_setopt(curl, CURLOPT_URL, url);
          curl_easy_setopt(curl, CURLOPT_POSTFIELDS, payload);
          curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);

          res = curl_easy_perform(curl);

          if(res != CURLE_OK) {
              fprintf(stderr, "curl_easy_perform() failed: %s\n",
                      curl_easy_strerror(res));
          }

          curl_slist_free_all(headers);
          curl_easy_cleanup(curl);
      }

      curl_global_cleanup();
      return 0;
  }
  ```

  ```objectivec Objective-C theme={null}
  #import <Foundation/Foundation.h>

  int main(int argc, const char * argv[]) {
      @autoreleasepool {
          NSURL *url = [NSURL URLWithString:@"https://api.apimart.ai/v1/chat/completions"];
          
          // 可替换为任意支持的模型 ID
          NSDictionary *payload = @{
              @"model": @"gpt-4o",
              @"messages": @[
                  @{
                      @"role": @"system",
                      @"content": @"你是一个专业的AI助手。"
                  },
                  @{
                      @"role": @"user",
                      @"content": @"介绍一下人工智能的发展历史。"
                  }
              ]
          };
          
          NSError *error;
          NSData *jsonData = [NSJSONSerialization dataWithJSONObject:payload
                                                            options:0
                                                              error:&error];
          
          NSMutableURLRequest *request = [NSMutableURLRequest requestWithURL:url];
          [request setHTTPMethod:@"POST"];
          [request setValue:@"Bearer <token>" forHTTPHeaderField:@"Authorization"];
          [request setValue:@"application/json" forHTTPHeaderField:@"Content-Type"];
          [request setHTTPBody:jsonData];
          
          NSURLSessionDataTask *task = [[NSURLSession sharedSession] 
              dataTaskWithRequest:request
              completionHandler:^(NSData *data, NSURLResponse *response, NSError *error) {
                  if (error) {
                      NSLog(@"Error: %@", error);
                      return;
                  }
                  NSString *result = [[NSString alloc] initWithData:data 
                                                          encoding:NSUTF8StringEncoding];
                  NSLog(@"%@", result);
              }];
          
          [task resume];
          [[NSRunLoop mainRunLoop] run];
      }
      return 0;
  }
  ```

  ```ocaml OCaml theme={null}
  (* Requires cohttp and yojson libraries *)
  open Lwt
  open Cohttp
  open Cohttp_lwt_unix

  let url = "https://api.apimart.ai/v1/chat/completions"

  (* 可替换为任意支持的模型 ID *)
  let payload = {|{
    "model": "gpt-4o",
    "messages": [
      {
        "role": "system",
        "content": "你是一个专业的AI助手。"
      },
      {
        "role": "user",
        "content": "介绍一下人工智能的发展历史。"
      }
    ]
  }|}

  let () =
    let headers = Header.init ()
      |> fun h -> Header.add h "Authorization" "Bearer <token>"
      |> fun h -> Header.add h "Content-Type" "application/json"
    in
    let body = Cohttp_lwt.Body.of_string payload in
    
    let response = Client.post ~headers ~body (Uri.of_string url) >>= fun (resp, body) ->
      body |> Cohttp_lwt.Body.to_string >|= fun body_str ->
      print_endline body_str
    in
    Lwt_main.run response
  ```

  ```dart Dart theme={null}
  import 'dart:convert';
  import 'package:http/http.dart' as http;

  void main() async {
    final url = Uri.parse('https://api.apimart.ai/v1/chat/completions');
    
    // 可替换为任意支持的模型 ID
    final payload = {
      'model': 'gpt-4o',
      'messages': [
        {
          'role': 'system',
          'content': '你是一个专业的AI助手。'
        },
        {
          'role': 'user',
          'content': '介绍一下人工智能的发展历史。'
        }
      ]
    };
    
    final response = await http.post(
      url,
      headers: {
        'Authorization': 'Bearer <token>',
        'Content-Type': 'application/json',
      },
      body: jsonEncode(payload),
    );
    
    print(response.body);
  }
  ```

  ```r R theme={null}
  library(httr)
  library(jsonlite)

  url <- "https://api.apimart.ai/v1/chat/completions"

  # 可替换为任意支持的模型 ID
  payload <- list(
    model = "gpt-4o",
    messages = list(
      list(
        role = "system",
        content = "你是一个专业的AI助手。"
      ),
      list(
        role = "user",
        content = "介绍一下人工智能的发展历史。"
      )
    )
  )

  response <- POST(
    url,
    add_headers(
      Authorization = "Bearer <token>",
      `Content-Type` = "application/json"
    ),
    body = toJSON(payload, auto_unbox = TRUE),
    encode = "raw"
  )

  cat(content(response, "text"))
  ```
</RequestExample>

<ResponseExample>
  ```json 200 theme={null}
  {
    "code": 200,
    "data": {
      "id": "chatcmpl-9876543210",
      "object": "chat.completion",
      "created": 1677652288,
      "model": "gpt-4o",
      "choices": [
        {
          "index": 0,
          "message": {
            "role": "assistant",
            "content": "人工智能（AI）的发展历史可以追溯到20世纪50年代...\n\n1. **早期阶段（1950s-1960s）**：图灵测试的提出标志着AI研究的开始...\n\n2. **专家系统时代（1970s-1980s）**：基于规则的系统开始应用于医疗诊断、金融分析等领域...\n\n3. **机器学习兴起（1990s-2000s）**：统计学习方法逐渐成为主流...\n\n4. **深度学习革命（2010s-至今）**：神经网络技术的突破带来了AI的爆发式发展..."
          },
          "finish_reason": "stop"
        }
      ],
      "usage": {
        "prompt_tokens": 28,
        "completion_tokens": 320,
        "total_tokens": 348
      }
    }
  }
  ```

  ```json 400 theme={null}
  {
    "error": {
      "code": 400,
      "message": "请求参数无效",
      "type": "invalid_request_error"
    }
  }
  ```

  ```json 401 theme={null}
  {
    "error": {
      "code": 401,
      "message": "身份验证失败，请检查您的API密钥",
      "type": "authentication_error"
    }
  }
  ```

  ```json 402 theme={null}
  {
    "error": {
      "code": 402,
      "message": "账户余额不足，请充值后再试",
      "type": "payment_required"
    }
  }
  ```

  ```json 403 theme={null}
  {
    "error": {
      "code": 403,
      "message": "访问被禁止，您没有权限访问此资源",
      "type": "permission_error"
    }
  }
  ```

  ```json 429 theme={null}
  {
    "error": {
      "code": 429,
      "message": "请求过于频繁，请稍后再试",
      "type": "rate_limit_error"
    }
  }
  ```

  ```json 500 theme={null}
  {
    "error": {
      "code": 500,
      "message": "服务器内部错误，请稍后重试",
      "type": "server_error"
    }
  }
  ```

  ```json 502 theme={null}
  {
    "error": {
      "code": 502,
      "message": "网关错误，服务器暂时不可用",
      "type": "bad_gateway"
    }
  }
  ```
</ResponseExample>

## Authorizations

<ParamField header="Authorization" type="string" required>
  所有接口均需要使用Bearer Token进行认证

  获取 API Key：

  访问 [API Key 管理页面](https://api.apimart.ai/console/token) 获取您的 API Key

  使用时在请求头中添加：

  ```
  Authorization: Bearer YOUR_API_KEY
  ```
</ParamField>

## Body

<ParamField body="model" type="string" required>
  模型名称

  支持的模型包括：

  * **OpenAI**: `gpt-5`, `gpt-5-chat-latest`, `gpt-5-mini`, `gpt-5-nano`, `gpt-5-pro`
  * **Anthropic**: `claude-sonnet-4-5-20250929`, `claude-opus-4-1-20250805`, `claude-haiku-4-5-20251001`, `claude-opus-4-1-20250805-thinking`, `claude-sonnet-4-5-20250929-thinking`
  * **Google**: `gemini-2.5-pro`, `gemini-2.5-flash`, `gemini-2.5-pro-thinking`, `gemini-2.5-flash-lite`
  * **DeepSeek**: `deepseek-v3.1-250821`, `deepseek-v3.1-think-250821`, `deepseek-v3-0324`
  * **Doubao**: `doubao-seed-1-6-251015`, `doubao-seed-1-6-flash-250828`, `doubao-seed-1-6-thinking-250715`
  * 更多模型持续更新中...
</ParamField>

<ParamField body="messages" type="array" required>
  对话消息列表

  每条消息包含：

  * `role`: 角色类型（`system`、`user`、`assistant`）
  * `content`: 消息内容（字符串或多模态内容数组）
</ParamField>

<ParamField body="temperature" type="number">
  控制输出随机性，范围 0-2

  * 较低的值（如 0.2）使输出更确定
  * 较高的值（如 1.8）使输出更随机

  默认值：1.0
</ParamField>

<ParamField body="max_tokens" type="integer">
  生成的最大token数量

  不同模型有不同的最大值限制，请参考具体模型文档
</ParamField>

<ParamField body="stream" type="boolean">
  是否使用流式输出

  * `true`: 流式返回（SSE格式）
  * `false`: 一次性返回完整响应

  默认值：false
</ParamField>

<ParamField body="top_p" type="number">
  核采样参数，范围 0-1

  控制生成文本的多样性，建议与 temperature 二选一使用

  默认值：1.0
</ParamField>

<ParamField body="frequency_penalty" type="number">
  频率惩罚，范围 -2.0 到 2.0

  正值会降低重复使用相同词汇的可能性

  默认值：0
</ParamField>

<ParamField body="presence_penalty" type="number">
  存在惩罚，范围 -2.0 到 2.0

  正值会增加谈论新主题的可能性

  默认值：0
</ParamField>

<ParamField body="stop" type="string or array">
  停止序列

  最多4个序列，遇到这些序列时将停止生成
</ParamField>

<ParamField body="n" type="integer">
  生成的回复数量

  默认值：1
</ParamField>

## Response

<ResponseField name="id" type="string">
  响应的唯一标识符
</ResponseField>

<ResponseField name="object" type="string">
  对象类型，固定为 `chat.completion`
</ResponseField>

<ResponseField name="created" type="integer">
  创建时间戳
</ResponseField>

<ResponseField name="model" type="string">
  实际使用的模型名称
</ResponseField>

<ResponseField name="choices" type="array">
  生成的回复列表

  <Expandable title="属性">
    <ResponseField name="index" type="integer">
      选项索引
    </ResponseField>

    <ResponseField name="message" type="object">
      消息内容

      <Expandable title="属性">
        <ResponseField name="role" type="string">
          角色类型（assistant）
        </ResponseField>

        <ResponseField name="content" type="string">
          生成的文本内容
        </ResponseField>
      </Expandable>
    </ResponseField>

    <ResponseField name="finish_reason" type="string">
      结束原因

      可能的值：

      * `stop` - 自然结束
      * `length` - 达到最大长度
      * `content_filter` - 内容过滤
      * `function_call` - 函数调用
    </ResponseField>
  </Expandable>
</ResponseField>

<ResponseField name="usage" type="object">
  token使用统计

  <Expandable title="属性">
    <ResponseField name="prompt_tokens" type="integer">
      输入消息的token数
    </ResponseField>

    <ResponseField name="completion_tokens" type="integer">
      生成内容的token数
    </ResponseField>

    <ResponseField name="total_tokens" type="integer">
      总token数
    </ResponseField>
  </Expandable>
</ResponseField>

## 支持的模型列表

### OpenAI 系列

* `gpt-5` - GPT-5 基础模型
* `gpt-5-chat-latest` - GPT-5 最新对话版本
* `gpt-5-mini` - GPT-5 轻量级版本，性价比高
* `gpt-5-nano` - GPT-5 超轻量版本
* `gpt-5-pro` - GPT-5 专业增强版

### Anthropic 系列

* `claude-haiku-4-5-20251001` - Claude 4.5 快速响应版本
* `claude-sonnet-4-5-20250929` - Claude 4.5 平衡版本
* `claude-opus-4-1-20250805` - 最强大的 Claude 4.1 旗舰模型
* `claude-opus-4-1-20250805-thinking` - Claude 4.1 Opus 深度思考版
* `claude-sonnet-4-5-20250929-thinking` - Claude 4.5 Sonnet 深度思考版

### Google 系列

* `gemini-2.5-flash` - Gemini 2.5 快速版
* `gemini-2.5-pro` - Gemini 2.5 专业版
* `gemini-2.5-flash-lite` - Gemini 2.5 超轻量版
* `gemini-2.5-pro-thinking` - Gemini 2.5 Pro 深度思考版

### DeepSeek 系列

* `deepseek-v3.1-250821` - DeepSeek V3.1 基础版
* `deepseek-v3.1-think-250821` - DeepSeek V3.1 思考版
* `deepseek-v3-0324` - DeepSeek V3 标准版

### Doubao 系列

* `doubao-seed-1-6-flash-250828` - Doubao Seed 1.6 快速版
* `doubao-seed-1-6-thinking-250715` - Doubao Seed 1.6 思考版
* `doubao-seed-1-6-251015` - Doubao Seed 1.6 标准版

## 使用示例

### 基础对话

```json  theme={null}
{
  "model": "gpt-4o",
  "messages": [
    {"role": "user", "content": "你好"}
  ]
}
```

### 系统提示词

```json  theme={null}
{
  "model": "claude-3-5-sonnet",
  "messages": [
    {"role": "system", "content": "你是一位专业的Python编程导师"},
    {"role": "user", "content": "如何使用列表推导式？"}
  ]
}
```

### 多轮对话

```json  theme={null}
{
  "model": "gemini-2.0-flash",
  "messages": [
    {"role": "user", "content": "什么是机器学习？"},
    {"role": "assistant", "content": "机器学习是人工智能的一个分支..."},
    {"role": "user", "content": "能举个例子吗？"}
  ]
}
```

### 流式输出

```json  theme={null}
{
  "model": "gpt-4o",
  "messages": [
    {"role": "user", "content": "写一首关于春天的诗"}
  ],
  "stream": true
}
```
