'use client'

import { useEffect } from 'react'
import { CheckCircle2, AlertTriangle, X } from 'lucide-react'
import { useMenuStore } from '@/lib/store'
import { cn } from '@/lib/utils'

export function AppAlert() {
  const alert = useMenuStore((s) => s.alert)
  const clearAlert = useMenuStore((s) => s.clearAlert)

  useEffect(() => {
    if (!alert) return
    const timeout = setTimeout(clearAlert, 4000)
    return () => clearTimeout(timeout)
  }, [alert, clearAlert])

  if (!alert) return null

  const isSuccess = alert.type === 'success'

  return (
    <div
      role="alert"
      className={cn(
        'mx-auto mt-4 flex max-w-[800px] items-center gap-3 rounded-md px-4 py-3 text-sm font-medium animate-in fade-in slide-in-from-top-2',
        isSuccess
          ? 'bg-[#e6f4ea] text-[#288c2a]'
          : 'bg-[#fff4e6] text-[#c23e00]',
      )}
    >
      {isSuccess ? (
        <CheckCircle2 className="size-4 shrink-0" />
      ) : (
        <AlertTriangle className="size-4 shrink-0" />
      )}
      <span className="flex-1">{alert.message}</span>
      <button
        onClick={clearAlert}
        className="shrink-0 rounded-sm p-0.5 opacity-70 transition-opacity hover:opacity-100"
        aria-label="Dismiss alert"
      >
        <X className="size-3.5" />
      </button>
    </div>
  )
}
