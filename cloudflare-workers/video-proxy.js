/**
 * Cloudflare Workers - 视频URL代理
 * 
 * 功能：隐藏真实的视频供应商域名，避免泄露上游信息
 * 
 * 部署方式：
 * 1. 登录 Cloudflare Dashboard
 * 2. 进入 Workers & Pages
 * 3. 创建新的 Worker
 * 4. 复制此代码并部署
 * 5. 设置环境变量：
 *    - API_KEY: 与后端配置的 CFWorkerVideoKey 保持一致
 * 
 * 使用方式：
 * 访问方式：
 * - 新：GET https://your-worker.workers.dev/v/<encrypted_token>.mp4  （推荐，仅 Proxy 模式）
 * - 兼容：GET https://your-worker.workers.dev/video?token=<encrypted_token>
 */

addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
})

async function handleRequest(request) {
  // 只允许 GET 请求
  if (request.method !== 'GET') {
    return new Response('Method Not Allowed', { status: 405 })
  }

  const url = new URL(request.url)
  
  // 健康检查端点
  if (url.pathname === '/health') {
    return new Response(JSON.stringify({ 
      status: 'ok', 
      service: 'video-proxy',
      mode: 'proxy'
    }), {
      headers: { 'Content-Type': 'application/json' }
    })
  }

  // 兼容旧端点：/video?token=xxx
  if (url.pathname === '/video') {
    return handleVideoProxy(url.searchParams.get('token'))
  }

  // 新端点：/v/<id>.mp4 或 /v/<token>
  if (url.pathname.startsWith('/v/')) {
    const seg = url.pathname.substring('/v/'.length)
    // 若查询参数携带 token，则优先使用（支持短文件名场景）
    const qpToken = url.searchParams.get('token')
    const token = qpToken || seg.replace(/\.[^/.]+$/, '') // 去掉后缀（如 .mp4）
    return handleVideoProxy(token)
  }

  return new Response('Not Found', { status: 404 })
}

async function handleVideoProxy(token) {
  try {
    if (!token) {
      return new Response(JSON.stringify({ 
        error: 'Missing token parameter' 
      }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      })
    }

    // 解密token获取真实URL
    const payload = await decryptToken(token)
    if (!payload || !payload.url) {
      return new Response(JSON.stringify({ 
        error: 'Invalid token' 
      }), {
        status: 400,
        headers: { 'Content-Type': 'application/json' }
      })
    }

    // 检查过期时间
    if (payload.expires_at && Date.now() / 1000 > payload.expires_at) {
      return new Response(JSON.stringify({ 
        error: 'Token expired' 
      }), {
        status: 401,
        headers: { 'Content-Type': 'application/json' }
      })
    }

    const realURL = payload.url
    // 仅支持：流式代理（完全隐藏URL）
    return await proxyVideo(realURL)

  } catch (error) {
    console.error('Video proxy error:', error)
    return new Response(JSON.stringify({ 
      error: 'Internal server error',
      message: error.message 
    }), {
      status: 500,
      headers: { 'Content-Type': 'application/json' }
    })
  }
}

/**
 * 流式代理视频内容
 */
async function proxyVideo(realURL) {
  try {
    const response = await fetch(realURL, {
      cf: {
        // Cloudflare 缓存配置
        cacheTtl: 3600,
        cacheEverything: true
      }
    })

    if (!response.ok) {
      return new Response(`Upstream error: ${response.status}`, { 
        status: response.status 
      })
    }

    // 复制响应头
    const headers = new Headers(response.headers)
    
    // 添加CORS头（如果需要）
    headers.set('Access-Control-Allow-Origin', '*')
    headers.set('Access-Control-Allow-Methods', 'GET, HEAD, OPTIONS')
    
    // 移除可能泄露上游信息的头
    headers.delete('Server')
    headers.delete('X-Powered-By')
    headers.delete('Via')

    return new Response(response.body, {
      status: response.status,
      headers: headers
    })

  } catch (error) {
    console.error('Proxy video error:', error)
    return new Response('Failed to fetch video', { status: 502 })
  }
}

/**
 * 解密token（使用AES-256-GCM）
 */
async function decryptToken(encryptedToken) {
  try {
    const apiKey = API_KEY || ''
    if (!apiKey) {
      throw new Error('API_KEY not configured')
    }

    // Base64 URL解码
    const ciphertext = base64UrlDecode(encryptedToken)
    
    // 生成密钥（SHA-256）
    const keyData = await crypto.subtle.digest(
      'SHA-256',
      new TextEncoder().encode(apiKey)
    )

    const key = await crypto.subtle.importKey(
      'raw',
      keyData,
      { name: 'AES-GCM' },
      false,
      ['decrypt']
    )

    // 提取nonce（前12字节）
    const nonceSize = 12
    if (ciphertext.byteLength < nonceSize) {
      throw new Error('Ciphertext too short')
    }

    const nonce = ciphertext.slice(0, nonceSize)
    const data = ciphertext.slice(nonceSize)

    // 解密
    const decrypted = await crypto.subtle.decrypt(
      {
        name: 'AES-GCM',
        iv: nonce
      },
      key,
      data
    )

    // 解析JSON
    const jsonStr = new TextDecoder().decode(decrypted)
    return JSON.parse(jsonStr)

  } catch (error) {
    console.error('Decrypt error:', error)
    return null
  }
}

/**
 * Base64 URL 解码
 */
function base64UrlDecode(str) {
  // 替换URL安全字符
  str = str.replace(/-/g, '+').replace(/_/g, '/')
  
  // 补齐padding
  while (str.length % 4) {
    str += '='
  }

  // Base64解码
  const binaryString = atob(str)
  const bytes = new Uint8Array(binaryString.length)
  for (let i = 0; i < binaryString.length; i++) {
    bytes[i] = binaryString.charCodeAt(i)
  }
  
  return bytes.buffer
}
