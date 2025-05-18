import axiosClient from './axiosClient'

interface RestAdapterParams {
  url: string
  method?: 'get' | 'post' | 'put' | 'delete' | 'patch'
  data?: unknown
  params?: Record<string, never>
  headers?: Record<string, string>
}

export const restAdapter = async <T = unknown>({
  url,
  method = 'get',
  data,
  params,
  headers,
}: RestAdapterParams): Promise<T> => {
  const response = await axiosClient({
    url,
    method,
    data,
    params,
    headers,
  })

  return response.data
}
