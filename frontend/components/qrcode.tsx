'use client'

import { cn } from '@/lib/utils'

interface QRCodeProps {
  value: string
  size?: number
  className?: string
}

export function QRCode({ value, size = 160, className }: QRCodeProps) {
  const url = `https://api.qrserver.com/v1/create-qr-code/?size=${size}x${size}&data=${encodeURIComponent(value)}`

  return (
    <div className={cn('flex flex-col items-center gap-2', className)}>
      {/* eslint-disable-next-line @next/next/no-img-element */}
      <img
        src={url}
        alt={`QR code for ${value}`}
        width={size}
        height={size}
        className="rounded-md border border-border bg-card"
        crossOrigin="anonymous"
      />
      <span className="max-w-[200px] truncate text-xs text-muted-foreground">
        {value}
      </span>
    </div>
  )
}
