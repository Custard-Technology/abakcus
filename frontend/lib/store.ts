import { create } from 'zustand'
import type { Menu, View } from './types'

interface AlertState {
  type: 'success' | 'error'
  message: string
}

interface MenuStore {
  /* data */
  menus: Menu[]
  setMenus: (menus: Menu[]) => void
  selectedMenu: Menu | null
  setSelectedMenu: (menu: Menu | null) => void

  /* navigation */
  view: View
  navigate: (view: View) => void

  /* loading */
  loading: boolean
  setLoading: (loading: boolean) => void

  /* alert */
  alert: AlertState | null
  showAlert: (type: 'success' | 'error', message: string) => void
  clearAlert: () => void
}

export const useMenuStore = create<MenuStore>((set) => ({
  menus: [],
  setMenus: (menus) => set({ menus }),
  selectedMenu: null,
  setSelectedMenu: (menu) => set({ selectedMenu: menu }),

  view: { kind: 'list' },
  navigate: (view) => set({ view, alert: null }),

  loading: false,
  setLoading: (loading) => set({ loading }),

  alert: null,
  showAlert: (type, message) => set({ alert: { type, message } }),
  clearAlert: () => set({ alert: null }),
}))
