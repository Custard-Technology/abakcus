
'use client'

import { useEffect } from 'react'
import { Plus, Eye, Pencil, Trash2, QrCode } from 'lucide-react'
import { listMenus, deleteMenu } from '@/lib/api'
import { useMenuStore } from '@/lib/store'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import { Spinner } from '@/components/ui/spinner'
import { QRCode } from '@/components/qrcode'
import { cn } from '@/lib/utils'
import { useState } from 'react'
import type { Menu } from '@/lib/types'

export function MenuList() {
  const menus = useMenuStore((s) => s.menus)
  const setMenus = useMenuStore((s) => s.setMenus)
  const loading = useMenuStore((s) => s.loading)
  const setLoading = useMenuStore((s) => s.setLoading)
  const navigate = useMenuStore((s) => s.navigate)
  const showAlert = useMenuStore((s) => s.showAlert)
  const [deletingId, setDeletingId] = useState<string | null>(null)
  const [qrOpenId, setQrOpenId] = useState<string | null>(null)

  useEffect(() => {
    async function load() {
      setLoading(true)
      try {
        const data = await listMenus()
        setMenus(data)
      } catch {
        showAlert('error', 'Failed to load menus. Please try again.')
      } finally {
        setLoading(false)
      }
    }
    load()
  }, [setMenus, setLoading, showAlert])

  async function handleDelete(menu: Menu) {
    if (!window.confirm(`Delete "${menu.name}"? This cannot be undone.`)) return
    setDeletingId(menu.id)
    try {
      await deleteMenu(menu.id)
      setMenus(menus.filter((m) => m.id !== menu.id))
      showAlert('success', `"${menu.name}" was deleted.`)
    } catch {
      showAlert('error', `Failed to delete "${menu.name}".`)
    } finally {
      setDeletingId(null)
    }
  }

  const menuUrl = (id: string) =>
    `https://abakcus.onrender.com/menus/${id}`

  if (loading) {
    return (
      <div className="flex flex-col gap-4">
        <div className="flex items-center justify-between">
          <Skeleton className="h-8 w-40" />
          <Skeleton className="h-10 w-32" />
        </div>
        {[1, 2, 3].map((i) => (
          <Skeleton key={i} className="h-28 w-full rounded-lg" />
        ))}
      </div>
    )
  }

  return (
    <div className="flex flex-col gap-6">
      {/* Header */}
      <div className="flex items-center justify-between">
        <div>
          <h2 className="font-[family-name:var(--font-heading)] text-2xl font-bold text-foreground">
            Your Menus
          </h2>
          <p className="mt-1 text-sm text-muted-foreground">
            {menus.length} menu{menus.length !== 1 ? 's' : ''} total
          </p>
        </div>
        <Button
          onClick={() => navigate({ kind: 'create' })}
          className="bg-primary text-primary-foreground hover:bg-primary/90"
        >
          <Plus className="size-4" />
          New Menu
        </Button>
      </div>

      {/* Empty state */}
      {menus.length === 0 && (
        <div className="flex flex-col items-center gap-4 rounded-lg border-2 border-dashed border-border bg-muted/50 py-16 text-center">
          <div className="flex size-14 items-center justify-center rounded-full bg-primary/10">
            <QrCode className="size-6 text-primary" />
          </div>
          <div>
            <p className="font-[family-name:var(--font-heading)] text-lg font-semibold text-foreground">
              No menus yet
            </p>
            <p className="mt-1 text-sm text-muted-foreground">
              Create your first menu to get started.
            </p>
          </div>
          <Button
            onClick={() => navigate({ kind: 'create' })}
            className="bg-primary text-primary-foreground hover:bg-primary/90"
          >
            <Plus className="size-4" />
            Create Menu
          </Button>
        </div>
      )}

      {/* Menu cards */}
      <div className="flex flex-col gap-4">
        {menus.map((menu) => (
          <div
            key={menu.id}
            className="group rounded-lg border border-border bg-card p-5 shadow-sm transition-shadow hover:shadow-md"
          >
            <div className="flex items-start justify-between gap-4">
              <div className="min-w-0 flex-1">
                <div className="flex items-center gap-2.5">
                  <h3 className="truncate font-[family-name:var(--font-heading)] text-lg font-semibold text-foreground">
                    {menu.name}
                  </h3>
                  <Badge
                    className={cn(
                      'shrink-0 text-[11px]',
                      menu.is_active
                        ? 'border-transparent bg-[#e6f4ea] text-[#288c2a]'
                        : 'border-transparent bg-muted text-muted-foreground',
                    )}
                  >
                    {menu.is_active ? 'Active' : 'Inactive'}
                  </Badge>
                </div>
                {menu.description && (
                  <p className="mt-1.5 line-clamp-2 text-sm leading-relaxed text-muted-foreground">
                    {menu.description}
                  </p>
                )}
                <p className="mt-2 text-xs text-muted-foreground/70">
                  Updated {new Date(menu.updated_at).toLocaleDateString()}
                </p>
              </div>
              <div className="flex shrink-0 items-center gap-1.5">
                <Button
                  variant="ghost"
                  size="icon-sm"
                  onClick={() => {
                    navigate({ kind: 'detail', menuId: menu.id })
                  }}
                  aria-label={`View ${menu.name}`}
                >
                  <Eye className="size-4" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon-sm"
                  onClick={() => navigate({ kind: 'edit', menuId: menu.id })}
                  aria-label={`Edit ${menu.name}`}
                >
                  <Pencil className="size-4" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon-sm"
                  onClick={() =>
                    setQrOpenId(qrOpenId === menu.id ? null : menu.id)
                  }
                  aria-label={`Show QR code for ${menu.name}`}
                  className={cn(
                    qrOpenId === menu.id && 'bg-secondary/10 text-secondary',
                  )}
                >
                  <QrCode className="size-4" />
                </Button>
                <Button
                  variant="ghost"
                  size="icon-sm"
                  disabled={deletingId === menu.id}
                  onClick={() => handleDelete(menu)}
                  aria-label={`Delete ${menu.name}`}
                  className="text-destructive hover:bg-destructive/10 hover:text-destructive"
                >
                  {deletingId === menu.id ? (
                    <Spinner className="size-4" />
                  ) : (
                    <Trash2 className="size-4" />
                  )}
                </Button>
              </div>
            </div>
            {/* QR code panel */}
            {qrOpenId === menu.id && (
              <div className="mt-4 flex items-center justify-center rounded-md border border-border bg-muted/50 py-5 animate-in fade-in slide-in-from-top-1">
                <QRCode value={menuUrl(menu.id)} size={140} />
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  )
}
