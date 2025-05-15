import axiosClient from './axiosClient'

interface RestAdapterParams {
  url: string
  method?: 'get' | 'post' | 'put' | 'delete' | 'patch'
  data?: any
  params?: Record<string, any>
  headers?: Record<string, string>
}

export const restAdapter = async <T = any>({
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
