'use client'

import { Navbar } from '@/components/navbar'
import { AppAlert } from '@/components/app-alert'
import { MenuList } from '@/components/menu-list'
import { MenuForm } from '@/components/menu-form'
import { MenuDetail } from '@/components/menu-detail'
import { useMenuStore } from '@/lib/store'

export default function Home() {
  const view = useMenuStore((s) => s.view)

  return (
    <div className="flex min-h-screen flex-col bg-background">
      <Navbar />
      <main className="mx-auto w-full max-w-[800px] flex-1 px-6 py-8">
        <AppAlert />
        <div className="mt-2">
          {view.kind === 'list' && <MenuList />}
          {view.kind === 'create' && <MenuForm />}
          {view.kind === 'edit' && <MenuForm menuId={view.menuId} />}
          {view.kind === 'detail' && <MenuDetail menuId={view.menuId} />}
        </div>
      </main>
      <footer className="border-t border-border bg-card py-4 text-center text-xs text-muted-foreground">
        Menu Manager &middot; made by MINDSGN STUDIO (PTY) LTD
      </footer>
    </div>
  )
}
