'use client'

import { useEffect, useState } from 'react'
import { ArrowLeft, Pencil, Trash2, Download } from 'lucide-react'
import { getMenu, deleteMenu, listMenus } from '@/lib/api'
import { useMenuStore } from '@/lib/store'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import { Spinner } from '@/components/ui/spinner'
import { QRCode } from '@/components/qrcode'
import { cn } from '@/lib/utils'
import type { Menu } from '@/lib/types'

interface MenuDetailProps {
  menuId: string
}

export function MenuDetail({ menuId }: MenuDetailProps) {
  const navigate = useMenuStore((s) => s.navigate)
  const showAlert = useMenuStore((s) => s.showAlert)
  const setMenus = useMenuStore((s) => s.setMenus)

  const [menu, setMenu] = useState<Menu | null>(null)
  const [loading, setLoading] = useState(true)
  const [deleting, setDeleting] = useState(false)

  const menuUrl = `https://abakcus.onrender.com/menus/${menuId}`

  useEffect(() => {
    async function load() {
      setLoading(true)
      try {
        const data = await getMenu(menuId)
        setMenu(data)
      } catch {
        showAlert('error', 'Failed to load menu details.')
        navigate({ kind: 'list' })
      } finally {
        setLoading(false)
      }
    }
    load()
  }, [menuId, showAlert, navigate])

  async function handleDelete() {
    if (!menu) return
    if (!window.confirm(`Delete "${menu.name}"? This cannot be undone.`)) return
    setDeleting(true)
    try {
      await deleteMenu(menuId)
      const menus = await listMenus()
      setMenus(menus)
      showAlert('success', `"${menu.name}" was deleted.`)
      navigate({ kind: 'list' })
    } catch {
      showAlert('error', `Failed to delete "${menu.name}".`)
      setDeleting(false)
    }
  }

  function handleDownloadQR() {
    const qrUrl = `https://api.qrserver.com/v1/create-qr-code/?size=400x400&data=${encodeURIComponent(menuUrl)}`
    const link = document.createElement('a')
    link.href = qrUrl
    link.download = `qr-${menu?.name?.replace(/\s+/g, '-').toLowerCase() || 'menu'}.png`
    link.target = '_blank'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }

  if (loading) {
    return (
      <div className="flex flex-col gap-6">
        <div className="flex items-center gap-3">
          <Skeleton className="size-8 rounded-md" />
          <Skeleton className="h-8 w-48" />
        </div>
        <div className="rounded-lg border border-border bg-card p-6">
          <Skeleton className="h-6 w-64" />
          <Skeleton className="mt-3 h-4 w-full" />
          <Skeleton className="mt-2 h-4 w-3/4" />
          <Skeleton className="mx-auto mt-6 size-40" />
        </div>
      </div>
    )
  }

  if (!menu) return null

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
          Menu Details
        </h2>
      </div>

      {/* Card */}
      <div className="rounded-lg border border-border bg-card p-6 shadow-sm">
        <div className="flex flex-col gap-5 sm:flex-row sm:gap-8">
          {/* Info */}
          <div className="flex min-w-0 flex-1 flex-col gap-3">
            <div className="flex items-center gap-2.5">
              <h3 className="truncate font-[family-name:var(--font-heading)] text-xl font-bold text-foreground">
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
              <p className="text-sm leading-relaxed text-muted-foreground">
                {menu.description}
              </p>
            )}

            <dl className="mt-1 flex flex-col gap-2 text-sm">
              <div className="flex gap-2">
                <dt className="font-medium text-foreground">Created:</dt>
                <dd className="text-muted-foreground">
                  {new Date(menu.created_at).toLocaleDateString()}
                </dd>
              </div>
              <div className="flex gap-2">
                <dt className="font-medium text-foreground">Updated:</dt>
                <dd className="text-muted-foreground">
                  {new Date(menu.updated_at).toLocaleDateString()}
                </dd>
              </div>
              <div className="flex gap-2">
                <dt className="font-medium text-foreground">ID:</dt>
                <dd className="truncate font-mono text-xs text-muted-foreground/70">
                  {menu.id}
                </dd>
              </div>
            </dl>
          </div>

          {/* QR Code */}
          <div className="flex flex-col items-center gap-3 rounded-md border border-border bg-muted/50 p-5">
            <p className="text-xs font-semibold uppercase tracking-wider text-muted-foreground">
              QR Code
            </p>
            <QRCode value={menuUrl} size={160} />
            <Button
              variant="outline"
              size="sm"
              onClick={handleDownloadQR}
              className="border-border text-foreground hover:bg-muted"
            >
              <Download className="size-3.5" />
              Download
            </Button>
          </div>
        </div>
      </div>

      {/* Actions */}
      <div className="flex items-center gap-3">
        <Button
          onClick={() => navigate({ kind: 'edit', menuId: menu.id })}
          className="bg-secondary text-secondary-foreground hover:bg-secondary/90"
        >
          <Pencil className="size-4" />
          Edit Menu
        </Button>
        <Button
          variant="outline"
          disabled={deleting}
          onClick={handleDelete}
          className="border-destructive/30 text-destructive hover:bg-destructive/10 hover:text-destructive"
        >
          {deleting ? <Spinner className="size-4" /> : <Trash2 className="size-4" />}
          Delete
        </Button>
      </div>
    </div>
  )
}
