'use client'

import { useEffect, useState } from 'react'
import { ArrowLeft } from 'lucide-react'
import { menuFormSchema, type MenuFormValues } from '@/lib/types'
import { createMenu, getMenu, updateMenu, listMenus } from '@/lib/api'
import { useMenuStore } from '@/lib/store'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Spinner } from '@/components/ui/spinner'
import { Skeleton } from '@/components/ui/skeleton'
import { cn } from '@/lib/utils'

interface MenuFormProps {
  menuId?: string
}

export function MenuForm({ menuId }: MenuFormProps) {
  const navigate = useMenuStore((s) => s.navigate)
  const showAlert = useMenuStore((s) => s.showAlert)
  const setMenus = useMenuStore((s) => s.setMenus)

  const isEdit = !!menuId

  const [name, setName] = useState('')
  const [description, setDescription] = useState('')
  const [isActive, setIsActive] = useState(true)
  const [errors, setErrors] = useState<Record<string, string>>({})
  const [submitting, setSubmitting] = useState(false)
  const [loadingMenu, setLoadingMenu] = useState(false)

  // Load existing menu for editing
  useEffect(() => {
    if (!menuId) return
    async function load() {
      setLoadingMenu(true)
      try {
        const menu = await getMenu(menuId!)
        setName(menu.name)
        setDescription(menu.description || '')
        setIsActive(menu.is_active)
      } catch {
        showAlert('error', 'Failed to load menu for editing.')
        navigate({ kind: 'list' })
      } finally {
        setLoadingMenu(false)
      }
    }
    load()
  }, [menuId, showAlert, navigate])

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault()
    setErrors({})

    const parsed = menuFormSchema.safeParse({ name, description, is_active: isActive })
    if (!parsed.success) {
      const fieldErrors: Record<string, string> = {}
      parsed.error.issues.forEach((issue) => {
        const key = issue.path[0]?.toString() || 'name'
        fieldErrors[key] = issue.message
      })
      setErrors(fieldErrors)
      return
    }

    setSubmitting(true)
    try {
      const payload: MenuFormValues = {
        name: parsed.data.name,
        description: parsed.data.description,
        is_active: parsed.data.is_active,
      }
      if (isEdit && menuId) {
        await updateMenu(menuId, payload)
        showAlert('success', `"${name}" updated successfully.`)
      } else {
        await createMenu(payload)
        showAlert('success', `"${name}" created successfully.`)
      }
      // Refresh the list
      const menus = await listMenus()
      setMenus(menus)
      navigate({ kind: 'list' })
    } catch {
      showAlert('error', `Failed to ${isEdit ? 'update' : 'create'} menu.`)
    } finally {
      setSubmitting(false)
    }
  }

  if (loadingMenu) {
    return (
      <div className="flex flex-col gap-6">
        <Skeleton className="h-8 w-48" />
        <Skeleton className="h-10 w-full" />
        <Skeleton className="h-10 w-full" />
        <Skeleton className="h-10 w-32" />
      </div>
    )
  }

  return (
    <div className="flex flex-col gap-6">
      {/* Back + title */}
      <div className="flex items-center gap-3">
        <Button
          variant="ghost"
          size="icon-sm"
          onClick={() => navigate({ kind: 'list' })}
          aria-label="Back to menu list"
        >
          <ArrowLeft className="size-4" />
        </Button>
        <h2 className="font-[family-name:var(--font-heading)] text-2xl font-bold text-foreground">
          {isEdit ? 'Edit Menu' : 'Create Menu'}
        </h2>
      </div>

      <form
        onSubmit={handleSubmit}
        noValidate
        className="flex flex-col gap-5 rounded-lg border border-border bg-card p-6 shadow-sm"
      >
        {/* Name */}
        <div className="flex flex-col gap-1.5">
          <Label
            htmlFor="menu-name"
            className="text-sm font-semibold text-foreground"
          >
            Menu Name <span className="text-destructive">*</span>
          </Label>
          <Input
            id="menu-name"
            value={name}
            onChange={(e) => {
              setName(e.target.value)
              if (errors.name) setErrors((prev) => ({ ...prev, name: '' }))
            }}
            placeholder="e.g. Lunch Menu"
            aria-invalid={!!errors.name}
            className={cn(
              'h-10 border-input bg-card text-foreground focus-visible:border-secondary focus-visible:ring-secondary/30',
              errors.name && 'border-destructive focus-visible:border-destructive focus-visible:ring-destructive/30',
            )}
          />
          {errors.name && (
            <p className="text-xs font-medium text-destructive">{errors.name}</p>
          )}
        </div>

        {/* Description */}
        <div className="flex flex-col gap-1.5">
          <Label
            htmlFor="menu-desc"
            className="text-sm font-semibold text-foreground"
          >
            Description
          </Label>
          <textarea
            id="menu-desc"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            rows={3}
            placeholder="Describe your menu..."
            className="w-full resize-none rounded-md border border-input bg-card px-3 py-2 text-sm text-foreground shadow-xs outline-none transition-colors placeholder:text-muted-foreground focus-visible:border-secondary focus-visible:ring-[3px] focus-visible:ring-secondary/30"
          />
        </div>

        {/* Active toggle (edit only) */}
        {isEdit && (
          <label className="flex cursor-pointer items-center gap-3">
            <input
              type="checkbox"
              checked={isActive}
              onChange={(e) => setIsActive(e.target.checked)}
              className="size-4 accent-primary"
            />
            <span className="text-sm font-medium text-foreground">Active</span>
          </label>
        )}

        {/* Actions */}
        <div className="flex items-center gap-3 pt-2">
          <Button
            type="submit"
            disabled={submitting}
            className="bg-primary text-primary-foreground hover:bg-primary/90"
          >
            {submitting && <Spinner className="size-4" />}
            {isEdit ? 'Save Changes' : 'Create Menu'}
          </Button>
          <Button
            type="button"
            variant="outline"
            onClick={() => navigate({ kind: 'list' })}
            className="border-border text-foreground hover:bg-muted"
          >
            Cancel
          </Button>
        </div>
      </form>
    </div>
  )
}
