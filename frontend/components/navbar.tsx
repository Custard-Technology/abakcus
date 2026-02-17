'use client'

import { UtensilsCrossed } from 'lucide-react'
import { useMenuStore } from '@/lib/store'

export function Navbar() {
  const navigate = useMenuStore((s) => s.navigate)

  return (
    <header className="sticky top-0 z-50 border-b border-border bg-card">
      <div className="mx-auto flex max-w-[800px] items-center gap-3 px-6 py-4">
        <button
          onClick={() => navigate({ kind: 'list' })}
          className="flex items-center gap-2.5 transition-opacity hover:opacity-80"
          aria-label="Go to menu list"
        >
          <div className="flex size-9 items-center justify-center rounded-lg bg-primary">
            <UtensilsCrossed className="size-5 text-primary-foreground" />
          </div>
          <h1 className="font-[family-name:var(--font-heading)] text-xl font-bold text-foreground">
            Menus
          </h1>
        </button>
      </div>
    </header>
  )
}
