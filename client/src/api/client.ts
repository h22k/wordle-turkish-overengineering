export interface Client {
  get<T>(url: string, params?: Record<string, unknown>): Promise<T>

  post<T, K>(url: string, data?: K): Promise<T>

  put<T, K>(url: string, data?: K): Promise<T>

  delete<T>(url: string): Promise<T>
}
