import { z } from 'zod'

/** Zod schema for creating / editing a menu */
export const menuFormSchema = z.object({
  name: z.string().min(1, 'Menu name is required'),
  description: z.string().optional(),
  is_active: z.boolean().optional(),
})

export type MenuFormValues = z.infer<typeof menuFormSchema>

/** Shape returned by the API */
export interface Menu {
  id: string
  business_id: string
  name: string
  description: string
  is_active: boolean
  created_at: string
  updated_at: string
}

/** Views used by conditional rendering */
export type View =
  | { kind: 'list' }
  | { kind: 'detail'; menuId: string }
  | { kind: 'create' }
  | { kind: 'edit'; menuId: string }
