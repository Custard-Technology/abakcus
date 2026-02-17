import type { Menu, MenuFormValues } from './types'

const BASE_URL = 'https://abakcus.onrender.com'
const BUSINESS_ID = 'business-123'

async function request<T>(
  path: string,
  options: RequestInit = {},
): Promise<T> {
  const res = await fetch(`${BASE_URL}${path}`, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      'X-Business-ID': BUSINESS_ID,
      ...options.headers,
    },
  })

  if (!res.ok) {
    const text = await res.text().catch(() => 'Unknown error')
    throw new Error(`API ${res.status}: ${text}`)
  }

  if (res.status === 204) return undefined as T

  return res.json() as Promise<T>
}

export async function listMenus(): Promise<Menu[]> {
  return request<Menu[]>('/menus')
}

export async function getMenu(id: string): Promise<Menu> {
  return request<Menu>(`/menus/${id}`)
}

export async function createMenu(data: MenuFormValues): Promise<Menu> {
  return request<Menu>('/menus', {
    method: 'POST',
    body: JSON.stringify(data),
  })
}

export async function updateMenu(
  id: string,
  data: MenuFormValues,
): Promise<Menu> {
  return request<Menu>(`/menus/${id}`, {
    method: 'PUT',
    body: JSON.stringify(data),
  })
}

export async function deleteMenu(id: string): Promise<void> {
  return request<void>(`/menus/${id}`, { method: 'DELETE' })
}
